/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir

import (
	"errors"
	"fmt"

	"github.com/nooga/let-go/pkg/vm"
)

// Lower emits a bytecode chunk equivalent to the IR Function.
//
// Lowering strategy (local-slot, approach D):
//
//   - Most nodes are emitted in source order within their block. A
//     stable stack slot at the def site lets later uses DUP_NTH from
//     that slot.
//   - Cheap nodes (Const/LoadArg/LoadVar/LoadClosed) follow a hybrid
//     policy via shouldBodyEmitCheap: body-emit if multi-use OR if
//     they are the first ref of their (single) use; otherwise defer
//     to the use site where materialize re-emits the cheap load
//     inline (matching the source compiler's natural layout and
//     avoiding a wasted DUP_NTH when another operand is deeper on
//     the stack).
//   - Single-use values land naturally on top of the stack and are
//     consumed in-place by the next op (refsAtTopLastUse fires).
//   - Multi-use values are kept at their def-site slot and DUP_NTH'd
//     to the top at each non-last use site; last use is consumed in
//     place when already at the top.
//   - Branch terminators recognise two common arg layouts: args
//     already at positions 0..N-1 from the bottom (forward fall-into
//     case), and args already at the TOP of the stack in order (loop
//     back-edge case). Both skip per-arg materialization and let the
//     branch's existing cleanup mechanism (RECUR's k/ignore) preserve
//     the args in place.
//   - Block-args are stack-positioned by predecessors before branching;
//     the entry of a block records their absolute positions in
//     valueStackPos.
//
// For straight-line code with no joins, the result should be
// byte-equivalent to the original bytecode.
func Lower(f *Function) (*vm.CodeChunk, error) {
	if f == nil {
		return nil, errors.New("ir.Lower: nil function")
	}

	// Detect cross-block Refs uses (deferred to follow-up issue D).
	// A value defined in block A used by nodes in block B (other than as a
	// BlockArg) cannot be expressed by the spike's stack-based lowering.
	uses := ComputeUses(f)
	for nid, us := range uses {
		if len(us) == 0 {
			continue
		}
		defBlock := f.Insts[nid].Block
		for _, userID := range us {
			userBlock := f.Insts[userID].Block
			if userBlock == defBlock {
				continue
			}
			// Cross-block use is only legitimate as a branch-target arg
			// (BlockArgs are how SSA crosses blocks). Direct Refs uses
			// can't be expressed by the spike's stack-based lowering.
			if isTerminatorBranchArgUse(&f.Insts[userID], InstId(nid)) {
				continue
			}
			return nil, fmt.Errorf("ir.Lower: cross-block use of value %%%d via direct Refs (defined in block %d, used in block %d) — only branch-target args may cross blocks; see follow-up issue D", nid, defBlock, userBlock)
		}
	}

	l := &lowerer{
		f:             f,
		chunk:         vm.NewCodeChunk(f.SourceConsts),
		blockIPs:      make([]int, len(f.Blocks)),
		nodeSlot:      map[InstId]int{},
		useCount:      make(map[InstId]int),
		valueStackPos: make(map[InstId]int),
		uses:          uses,
	}
	// Seed useCount from the def→use index (reuse the already-computed uses).
	for id, us := range uses {
		if len(us) > 0 {
			l.useCount[InstId(id)] = len(us)
		}
	}

	// Phase 1: emit each block, recording its starting IP.
	// We emit blocks in ID order (matches construction order, which
	// follows the source bytecode's natural layout).
	for bid := range f.Blocks {
		l.blockIPs[bid] = l.chunk.Length()
		if err := l.lowerBlock(BlockID(bid)); err != nil {
			return nil, err
		}
	}

	// Phase 2: patch branch target offsets. Branches were emitted
	// with placeholder offsets (0); now resolve to absolute IPs.
	if err := l.patchBranches(); err != nil {
		return nil, err
	}

	return l.chunk, nil
}

type lowerer struct {
	f        *Function
	chunk    *vm.CodeChunk
	blockIPs []int          // block ID → starting IP in emitted chunk
	nodeSlot map[InstId]int // (unused; reserved for future local-slot allocator)
	patches  []branchPatch  // pending offset patches
	stackSP  int            // current stack depth (logical, for SP encoding)

	useCount      map[InstId]int // total use count per node; decremented by the materializer as consumers are emitted
	valueStackPos map[InstId]int // absolute stack-depth index (0 = bottom) of values currently on stack
	uses          Uses           // def→use index (set once at Lower-time)
}

type branchPatch struct {
	srcIP       int // IP of the branch opcode (where offset is written at srcIP+1, or +2 for cond)
	targetBlock BlockID
	offsetSlot  int  // word offset from srcIP where the patched value goes (1 for unconditional, ...)
	negate      bool // true for OP_RECUR (ip -= offset; offset = srcIP - targetIP)
}

// recordSourceInfo records all the node's source spans at the current
// write position. Call BEFORE emitting the bytecode word the spans
// describe — SourceMap.Add associates with c.length, which is the IP
// about to be written. Multiple spans at the same IP are stored as
// separate entries; SourceMap.Lookup returns the most recent (the last
// span added), which gives a deterministic answer.
func (l *lowerer) recordSourceInfo(n *Inst) {
	if n == nil {
		return
	}
	for _, si := range n.SourceInfos {
		l.chunk.AddSourceInfo(si)
	}
}

func (l *lowerer) emit(op int32) {
	l.chunk.Append(op | int32(l.stackSP<<16))
}

func (l *lowerer) emitWithArg(op int32, arg int) {
	l.chunk.Append(op | int32(l.stackSP<<16))
	l.chunk.Append32(arg)
}

func (l *lowerer) emitPlaceholder(op int32) int {
	// Returns the IP of the argument slot (caller patches via Update32).
	l.chunk.Append(op | int32(l.stackSP<<16))
	argIP := l.chunk.Length()
	l.chunk.Append32(0) // placeholder
	return argIP
}

func (l *lowerer) lowerBlock(bid BlockID) error {
	blk := &l.f.Blocks[bid]
	// Block-args are assumed to be on the stack already (placed there
	// by predecessors before branching). The deepest one is on the
	// bottom of the block's "fresh" stack. The entry block has no
	// Params (function args live in f.args[] and are accessed via
	// OpLoadArg, not the value stack), so this falls through to
	// stackSP = 0 naturally.
	l.stackSP = len(blk.Params)
	for i, pid := range blk.Params {
		l.valueStackPos[pid] = i
	}

	// Lower each body node in source order. Every node — including
	// cheap ones (Const/LoadArg/LoadVar/LoadClosed) — gets a stable
	// stack slot at its definition site. Multi-use values are reused
	// from that slot via DUP_NTH; single-use values feed the next op
	// directly.
	for _, nid := range blk.Insts {
		n := l.f.Inst(nid)
		// Skip dead nodes (left over from CSE/DCE).
		if n.Op == OpInvalid {
			continue
		}
		// BlockArgs aren't emitted (they're stack-positioned by predecessors).
		if n.Op == OpBlockArg {
			continue
		}
		// OpPop is constructed by translateOp as a marker but doesn't
		// emit any bytecode; skip.
		if n.Op == OpPop {
			// Decrement the operand's use count, since the consume was virtual.
			for _, r := range n.Refs {
				if c, ok := l.useCount[r]; ok && c > 0 {
					l.useCount[r] = c - 1
				}
			}
			continue
		}
		// Cheap-node body-emit policy:
		//
		//   - Multi-use cheap nodes are body-emitted so subsequent
		//     uses can DUP_NTH from a stable slot (cheaper than
		//     re-LOADing).
		//   - Single-use cheap nodes are body-emitted ONLY if they're
		//     the first ref of their use. In that position, the
		//     body-walk's natural ordering will leave them at the
		//     bottom of the consumer's operand slice — exactly where
		//     the consumer expects them — and refsAtTopLastUse will
		//     fire, eliminating any materialize cost.
		//   - Otherwise, defer: the use site's materialize will
		//     re-emit them inline via the cheap fallback path, which
		//     mirrors the source compiler's natural layout and avoids
		//     wasting an opcode when another operand is deeper on the
		//     stack.
		if isCheapMaterializeOp(n.Op) && !l.shouldBodyEmitCheap(nid) {
			continue
		}
		// Fast path: if the node's refs are already laid out at the
		// top of the stack in correct order AND each is at its last
		// use, just consume them in place. This is the common
		// straight-line case after body-walk emission: a binop's two
		// operands were just emitted into positions (sp-2, sp-1).
		if l.refsAtTopLastUse(n.Refs) {
			for _, r := range n.Refs {
				if c, ok := l.useCount[r]; ok && c > 0 {
					l.useCount[r] = c - 1
				}
			}
		} else {
			// Materialize each operand in order.
			for _, r := range n.Refs {
				if err := l.materialize(r); err != nil {
					return fmt.Errorf("block %d node %d: %w", bid, nid, err)
				}
				if c, ok := l.useCount[r]; ok && c > 0 {
					l.useCount[r] = c - 1
				}
			}
		}
		// Emit the opcode for this node (consumes operands, may push result).
		if err := l.lowerNode(nid); err != nil {
			return fmt.Errorf("block %d node %d: %w", bid, nid, err)
		}
		// If this op pushes a result, record its stack position.
		if n.Op.StackOut() == 1 {
			l.valueStackPos[nid] = l.stackSP - 1
		}
	}

	// Lower the terminator.
	if blk.Term == 0 && bid != l.f.Entry {
		// Unset terminator on a non-entry block — likely a construction bug;
		// caller should have set one.
		return fmt.Errorf("block %d has no terminator", bid)
	}
	if blk.Term != 0 {
		// Materialize terminator operands.
		term := l.f.Inst(blk.Term)
		// Direct Refs (e.g., OpReturn's value, OpBranchIf's cond,
		// OpTailCall's fn + args, OpCall's fn + args). The same
		// "already at top in correct order" fast path applies as for
		// body nodes.
		if l.refsAtTopLastUse(term.Refs) {
			for _, r := range term.Refs {
				if c, ok := l.useCount[r]; ok && c > 0 {
					l.useCount[r] = c - 1
				}
			}
		} else {
			for _, r := range term.Refs {
				if err := l.materialize(r); err != nil {
					return fmt.Errorf("block %d terminator refs: %w", bid, err)
				}
				if c, ok := l.useCount[r]; ok && c > 0 {
					l.useCount[r] = c - 1
				}
			}
		}
		// Branch-target Args.
		switch t := term.Aux.(type) {
		case *BranchTarget:
			if t != nil {
				args := t.Args
				// Common case 1: args are already in their target
				// positions (positions 0..len(args)-1 from bottom).
				// This is the "fall-through with no stack to clean"
				// case (the entry block's first jump into a loop
				// header, for example).
				skip := 0
				for skip < len(args) {
					if pos, ok := l.valueStackPos[args[skip]]; ok && pos == skip {
						skip++
						continue
					}
					break
				}
				// Common case 2: args are already at the TOP of the
				// stack in correct order. This is the back-edge of a
				// loop after body-walk emission: the stack ends in
				// [..., arg0, arg1, ..., argN-1] and we want RECUR to
				// drop the prefix while preserving the top args.
				if skip < len(args) && l.argsAtTop(args) {
					for _, a := range args {
						if c, ok := l.useCount[a]; ok && c > 0 {
							l.useCount[a] = c - 1
						}
					}
					break
				}
				for _, a := range args[skip:] {
					if err := l.materialize(a); err != nil {
						return fmt.Errorf("block %d branch args: %w", bid, err)
					}
					if c, ok := l.useCount[a]; ok && c > 0 {
						l.useCount[a] = c - 1
					}
				}
				// Decrement use counts for the skipped args (they're still
				// consumed by the branch, just without re-emit).
				for _, a := range args[:skip] {
					if c, ok := l.useCount[a]; ok && c > 0 {
						l.useCount[a] = c - 1
					}
				}
			}
		case *CondTarget:
			_ = t
			// BranchIf's cond is in Refs (materialized above). True/False
			// branch args for BranchIf are passed via runtime semantics that
			// don't yet pass args in the spike; leave as no-op.
		}
		if err := l.lowerNode(blk.Term); err != nil {
			return fmt.Errorf("block %d terminator: %w", bid, err)
		}
	}
	return nil
}

// argsAtTop reports whether the given args occupy the top
// len(args) stack slots in order: args[0] at stackSP-len(args),
// args[len-1] at stackSP-1.
func (l *lowerer) argsAtTop(args []InstId) bool {
	n := len(args)
	if n == 0 {
		return true
	}
	base := l.stackSP - n
	if base < 0 {
		return false
	}
	for i, a := range args {
		pos, ok := l.valueStackPos[a]
		if !ok || pos != base+i {
			return false
		}
	}
	return true
}

// shouldBodyEmitCheap reports whether the given cheap node should be
// emitted at its definition site (body walk) rather than deferred to
// its use site. Body-emit when:
//
//   - The node has multiple uses: a stable def-site slot lets each
//     subsequent use DUP_NTH instead of re-LOAD.
//   - The node has a single use AND it's the first ref of that use
//     ("fn position" of a Call, "left operand" of a binop, etc.): the
//     body-walk will leave it at the bottom of the consumer's operand
//     slice, where refsAtTopLastUse can fire without DUP_NTH cost.
//
// Otherwise defer: the consumer's materialize will re-LOAD inline,
// which is exactly what the source compiler emits for the same expr.
func (l *lowerer) shouldBodyEmitCheap(nid InstId) bool {
	uc := l.useCount[nid]
	if uc == 0 {
		// Dead; don't bother emitting.
		return false
	}
	// For multi-use cheap nodes, the body-emit-then-DUP_NTH strategy
	// costs MORE than just re-emitting the cheap load at each use site.
	// Profiling (cmul / cmagsq with N=2 uses each) showed: body-emit + N
	// DUP_NTH = 2 + 2N bytes vs source compiler's N LOAD_ARGs = 2N bytes.
	// Even at the break-even point body-emit loses on dispatch count.
	// So: for cheap multi-use, re-emit at the use site (cheap path of
	// materialize handles this with a single LOAD_ARG / LOAD_CONST etc).
	//
	// Body-emit is reserved for the SINGLE-USE-AS-FIRST-REF case where
	// the body-walk's natural ordering leaves the value at the consumer's
	// expected position (refsAtTopLastUse fires, zero cost).
	if uc > 1 {
		return false
	}
	// Single use: find it and check if this node is the first ref.
	if int(nid) >= len(l.uses) {
		return false
	}
	us := l.uses[nid]
	if len(us) != 1 {
		return false
	}
	user := l.f.Inst(us[0])
	if len(user.Refs) == 0 {
		// Used only as a branch-target arg (no direct Refs); body
		// emit so the branch's arg-materialize sees a stable position.
		return true
	}
	return user.Refs[0] == nid
}

// refsAtTopLastUse reports whether the given refs occupy the top
// len(refs) stack slots in order AND each is at its last use. When
// true, the consuming op can pop them in place without any
// materialization.
func (l *lowerer) refsAtTopLastUse(refs []InstId) bool {
	n := len(refs)
	if n == 0 {
		return true
	}
	if !l.argsAtTop(refs) {
		return false
	}
	// Each ref must appear exactly once in refs (no duplicate ref →
	// same value used twice in the same op) AND be at its last use.
	// Use an O(n²) scan instead of a map — refs lists are short
	// (typically 1-3 items), so this avoids per-call map allocation.
	for i, r := range refs {
		if l.useCount[r] != 1 {
			return false
		}
		for j := i + 1; j < n; j++ {
			if refs[j] == r {
				return false
			}
		}
	}
	return true
}

func (l *lowerer) lowerNode(nid InstId) error {
	n := l.f.Inst(nid)
	switch n.Op {

	case OpBlockArg:
		// Block-args are pre-placed on the stack by predecessors;
		// no opcode emission needed.
		return nil

	case OpConst:
		idx := l.f.SourceConsts.Intern(n.Aux.(vm.Value))
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_LOAD_CONST, idx)
		l.stackSP++
		l.bumpMaxStack()

	case OpLoadArg:
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_LOAD_ARG, n.Aux.(int))
		l.stackSP++
		l.bumpMaxStack()

	case OpLoadVar:
		idx := l.f.SourceConsts.Intern(n.Aux.(vm.Value))
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_LOAD_VAR, idx)
		l.stackSP++
		l.bumpMaxStack()

	case OpAdd, OpSub, OpMul, OpLt, OpLte, OpGt, OpGte, OpEq:
		// Operands are expected on top of the stack already, in order.
		// For the spike, we trust the source-order walk to have placed
		// them via prior LoadArg/LoadConst/etc. emissions.
		l.recordSourceInfo(n)
		l.emit(irOpToBytecode(n.Op))
		l.stackSP-- // 2 -> 1

	case OpInc, OpDec:
		l.recordSourceInfo(n)
		l.emit(irOpToBytecode(n.Op))
		// 1 -> 1; no SP change

	case OpCall:
		arity := n.Aux.(int)
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_INVOKE, arity)
		l.stackSP -= arity // pops fn + arity args, pushes result; net -arity

	case OpTailCall:
		arity := n.Aux.(int)
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_TAIL_CALL, arity)
		// TAIL_CALL is a true terminator in our model: for *Func callees
		// the VM replaces the frame; for builtin callees the VM returns
		// the result directly from Frame.Run. No trailing RETURN needed.

	case OpReturn:
		l.recordSourceInfo(n)
		l.emit(vm.OP_RETURN)

	case OpBranch:
		bt := n.Aux.(*BranchTarget)
		// After materializing branch args, the top `argc` items on the
		// stack are the args we promised the target. The target block
		// expects sp == len(target.Params) on entry; if the current sp
		// exceeds that, we have "dead" values below the args (e.g.,
		// loop-bindings from a previous iteration) that must be cleaned
		// up. We use OP_RECUR's atomic "save top, drop, restore" for
		// multi-arg cleanup since there's no swap/store primitive.
		argc := len(bt.Args)
		targetParams := len(l.f.Blocks[bt.Target].Params)
		curSP := l.stackSP
		dropCount := curSP - targetParams
		l.recordSourceInfo(n)
		if dropCount > 0 {
			// We want to end at sp = targetParams. Two layouts to handle:
			//
			//   (a) Back-edge (recur): stack = [...old-K-params..., ...new-K-args...]
			//       where K = argc = targetParams. curSP = K + K = 2K (or more,
			//       with intermediate non-tracked junk between). dropCount = K
			//       (or more). OP_RECUR(k=argc, ignore=dropCount-argc) is the
			//       right mechanism: pops argc args saved, drops dropCount-argc
			//       extra items, then drops the saved argc items, pushes argc
			//       back. Net stack change = -(argc + ignore) = -dropCount.
			//
			//   (b) Forward branch with extra junk (e.g., LICM pre-header):
			//       stack = [..., extra-junk..., ...K-args...]. argc > dropCount.
			//       The args are NOT replacing stale slots; they're fresh.
			//       OP_RECUR doesn't fit (its semantics drop argc*2+ignore from
			//       top, which would consume the args twice). Instead: save
			//       args, drop dropCount junk items, restore args. Same shape
			//       as RECUR but with ignore = dropCount and the runtime
			//       drop count = argc + ignore = argc + dropCount.
			//
			// The runtime's RECUR handler drops `argc*2 + ignore` items, which
			// implements case (a). For case (b) we'd need a different
			// "save-drop-restore" primitive. As a stopgap, we can simulate
			// case (b) using an UNUSED form: a sequence of DUP_NTH pulls the
			// non-arg dead values into junk-only positions and then we POP_N
			// them. Or even simpler: emit POP_N to drop just the non-arg dead
			// (NOT the args), but POP_N drops contiguous items from top —
			// which would eat the args.
			//
			// For now: only the back-edge case (dropCount >= argc) works.
			// The forward-with-junk case errors out — caller (LICM) should
			// arrange that no junk sits below newly-materialized args.
			if dropCount < argc {
				return fmt.Errorf("ir.Lower: block %d branch to %d has dropCount=%d < argc=%d (curSP=%d targetParams=%d) — pre-header→header with junk-below-args not yet handled; LICM may need to thread values differently", n.Block, bt.Target, dropCount, argc, curSP, targetParams)
			}
			ignore := dropCount - argc
			l.chunk.Append(vm.OP_RECUR | int32(l.stackSP<<16))
			offIP := l.chunk.Length()
			l.chunk.Append32(0) // placeholder offset
			l.chunk.Append32(argc)
			l.chunk.Append32(ignore)
			l.patches = append(l.patches, branchPatch{
				srcIP:       offIP - 1,
				targetBlock: bt.Target,
				offsetSlot:  1,
				negate:      true,
			})
			l.stackSP = targetParams
			return nil
		}
		argIP := l.emitPlaceholder(vm.OP_JUMP)
		l.patches = append(l.patches, branchPatch{
			srcIP:       argIP - 1,
			targetBlock: bt.Target,
			offsetSlot:  1,
		})

	case OpBranchIf:
		ct := n.Aux.(*CondTarget)
		// Pop the cond. We emit BRANCH_FALSE to the FALSE target;
		// fall-through goes to the TRUE target (which we'll emit a
		// JUMP for if it's not the next block).
		l.recordSourceInfo(n)
		argIP := l.emitPlaceholder(vm.OP_BRANCH_FALSE)
		l.patches = append(l.patches, branchPatch{
			srcIP:       argIP - 1,
			targetBlock: ct.False.Target,
			offsetSlot:  1,
		})
		l.stackSP-- // pops cond

		// If true-target is the immediately-following block, no JUMP
		// needed (fall-through). Otherwise emit an unconditional jump.
		nextBlockID := BlockID(int(n.Block) + 1)
		if ct.True.Target != nextBlockID {
			l.recordSourceInfo(n)
			argIP2 := l.emitPlaceholder(vm.OP_JUMP)
			l.patches = append(l.patches, branchPatch{
				srcIP:       argIP2 - 1,
				targetBlock: ct.True.Target,
				offsetSlot:  1,
			})
		}

	case OpLoadClosed:
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_LOAD_CLOSEDOVER, n.Aux.(int))
		l.stackSP++
		l.bumpMaxStack()

	case OpMakeClosure:
		// Refs[0] = func value, already on top of stack (materialized
		// by the body walk). OP_MAKE_CLOSURE pops the *Func and pushes
		// a *Closure in its place — net stack change 0.
		l.recordSourceInfo(n)
		l.emit(vm.OP_MAKE_CLOSURE)
		// stack: 1 in, 1 out → no change

	case OpPushClosed:
		// Refs[0] = closure, Refs[1] = value. After materialization
		// they sit as [closure, value] on top of stack. OP_PUSH_CLOSEDOVER
		// pops the value and mutates the closure beneath, leaving the
		// closure on top — net -1.
		l.recordSourceInfo(n)
		l.emit(vm.OP_PUSH_CLOSEDOVER)
		l.stackSP-- // 2 in, 1 out

	default:
		return fmt.Errorf("unsupported op for lowering: %s", n.Op)
	}
	return nil
}

func (l *lowerer) patchBranches() error {
	for _, p := range l.patches {
		targetIP := l.blockIPs[p.targetBlock]
		var offset int
		if p.negate {
			// OP_RECUR uses f.ip -= offset, so we encode positive for backward.
			offset = p.srcIP - targetIP
		} else {
			offset = targetIP - p.srcIP
		}
		l.chunk.Update32(p.srcIP+p.offsetSlot, int32(offset))
	}
	return nil
}

// isTerminatorBranchArgUse reports whether a terminator node references
// target via a branch-target Args slot (the legitimate SSA cross-block
// channel). Direct Refs uses are not cross-block-safe.
func isTerminatorBranchArgUse(term *Inst, target InstId) bool {
	if !term.Op.IsTerminator() {
		return false
	}
	switch t := term.Aux.(type) {
	case *BranchTarget:
		if t != nil {
			for _, a := range t.Args {
				if a == target {
					return true
				}
			}
		}
	case *CondTarget:
		if t != nil {
			if t.True != nil {
				for _, a := range t.True.Args {
					if a == target {
						return true
					}
				}
			}
			if t.False != nil {
				for _, a := range t.False.Args {
					if a == target {
						return true
					}
				}
			}
		}
	}
	return false
}

// materialize ensures the value produced by nid is available on top of
// the stack at the point of call.
//
// Since every node is body-emitted in source order (including cheap
// ops), the value is always already on the stack at some tracked
// position. The two paths are:
//
//  1. Already on top AND this is the last remaining use → no emission
//     needed; the consumer will pop it in-place.
//  2. Otherwise → emit DUP_NTH to copy from the tracked position to
//     the top.
//
// A cheap-re-emit fallback is kept for safety (e.g., a value whose
// body emission was skipped by some future path); it should not
// trigger in normal flow.
//
// Caller is responsible for decrementing useCount[nid] after the
// consuming op finishes.
func (l *lowerer) materialize(nid InstId) error {
	n := l.f.Inst(nid)
	// Last-use-at-top optimization: if the value is already on top of
	// the stack and this is its last remaining use, just leave it there.
	// The upcoming consuming op will pop it directly.
	if pos, ok := l.valueStackPos[nid]; ok {
		if pos == l.stackSP-1 && l.useCount[nid] == 1 {
			return nil
		}
		// Default: DUP_NTH from the tracked position.
		nth := l.stackSP - 1 - pos
		l.emitWithArg(vm.OP_DUP_NTH, nth)
		l.stackSP++
		l.bumpMaxStack()
		return nil
	}
	// Safety net: if the value isn't tracked on the stack, fall back to
	// cheap re-materialization. With the current body-walk emitting all
	// nodes in source order, this should not occur.
	switch n.Op {
	case OpConst:
		idx := l.f.SourceConsts.Intern(n.Aux.(vm.Value))
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_LOAD_CONST, idx)
		l.stackSP++
		l.bumpMaxStack()
		return nil
	case OpLoadArg:
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_LOAD_ARG, n.Aux.(int))
		l.stackSP++
		l.bumpMaxStack()
		return nil
	case OpLoadVar:
		idx := l.f.SourceConsts.Intern(n.Aux.(vm.Value))
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_LOAD_VAR, idx)
		l.stackSP++
		l.bumpMaxStack()
		return nil
	case OpLoadClosed:
		l.recordSourceInfo(n)
		l.emitWithArg(vm.OP_LOAD_CLOSEDOVER, n.Aux.(int))
		l.stackSP++
		l.bumpMaxStack()
		return nil
	}
	return fmt.Errorf("ir.Lower: value %%%d not on stack for materialize (op=%s)", nid, n.Op)
}

// bumpMaxStack updates the chunk's recorded peak depth so the VM can
// size its frame stack correctly.
func (l *lowerer) bumpMaxStack() {
	if l.stackSP > l.chunk.MaxStack() {
		l.chunk.SetMaxStack(l.stackSP)
	}
}
