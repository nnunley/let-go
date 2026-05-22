/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package passes

import (
	"github.com/nooga/let-go/pkg/ir"
	"github.com/nooga/let-go/pkg/vm"
)

// LICM hoists loop-invariant computations out of loops to their
// pre-headers, reducing redundant work inside hot iteration paths.
//
// Algorithm:
//
//  1. Compute dominators (f.Dominators()).
//  2. Find natural loops: a back-edge is an edge from block B to header
//     H where H dominates B. The loop body is all blocks dominated by H
//     that can reach B in the reverse CFG.
//  3. For each loop, identify a pre-header: the unique predecessor of
//     H that's NOT in the loop body. If there are multiple such
//     predecessors, skip this loop (creating a synthetic pre-header
//     is future work). If there's none, also skip (would need to
//     synthesize one, again future work).
//  4. For each pure node defined in the loop body, check if all its
//     operands are loop-invariant (defined outside the loop, OR are
//     themselves marked invariant earlier in this same pass). If so,
//     move it to the pre-header.
//
// Operates in-place; doesn't remove the original node's slot — moves
// the node by changing its Block and re-appending to the pre-header's
// body list. Original block's Insts list is rebuilt without it.
//
// Returns true if any node was hoisted.
//
// Limitations (filed as follow-ups):
//   - Doesn't synthesize pre-headers; skips loops without a clean one.
//   - Doesn't hoist across nested loops (each loop processed indepen-
//     dently; outer-loop-invariants from inner-loop bodies stay put).
//   - Doesn't handle multi-exit loops differently from single-exit.
//   - LoadVar of mutated vars is NOT invariant (uses the same mutability
//     analysis as CSE via computeMutatedVars).
func LICM(f *ir.Function) (changed bool) {
	idom := f.Dominators()
	if len(idom) == 0 {
		return false
	}

	loops := findLoops(f, idom)
	if len(loops) == 0 {
		return false
	}

	mutatedVars, allMutated := computeMutatedVars(f)

	for _, loop := range loops {
		if hoistInvariants(f, idom, loop, mutatedVars, allMutated) {
			changed = true
		}
	}
	return changed
}

// loopInfo describes a natural loop discovered via back-edge analysis.
type loopInfo struct {
	header    ir.BlockID
	body      map[ir.BlockID]bool // all blocks in the loop, including header
	preheader ir.BlockID          // the unique outside-loop predecessor of header, or -1 if none/multiple
}

// findLoops identifies natural loops in f. Back-edges define loops:
// an edge from B to H is a back-edge if H dominates B. The loop body
// is all blocks dominated by H that can reach B without leaving the
// region of H-dominated blocks.
//
// Returns one loopInfo per back-edge. (A header with multiple back-
// edges produces multiple loopInfo records sharing the same header —
// LICM treats them independently but in practice they collapse.)
func findLoops(f *ir.Function, idom []ir.BlockID) []loopInfo {
	var loops []loopInfo
	for bi := range f.Blocks {
		bid := ir.BlockID(bi)
		for _, succ := range f.Successors(bid) {
			if !ir.Dominates(idom, succ, bid) {
				continue
			}
			// Back-edge bid → succ. succ is the loop header.
			body := collectLoopBody(f, succ, bid)
			pre := findPreheader(f, succ, body)
			loops = append(loops, loopInfo{
				header:    succ,
				body:      body,
				preheader: pre,
			})
		}
	}
	return loops
}

// collectLoopBody returns the set of blocks in the natural loop whose
// header is `header` and that contains the back-edge from `tail`.
//
// Algorithm: start with {header, tail}; do reverse-CFG BFS from tail,
// stopping at header. Every block we visit is in the loop.
func collectLoopBody(f *ir.Function, header, tail ir.BlockID) map[ir.BlockID]bool {
	body := map[ir.BlockID]bool{header: true}
	if header == tail {
		return body
	}
	body[tail] = true
	worklist := []ir.BlockID{tail}
	for len(worklist) > 0 {
		b := worklist[0]
		worklist = worklist[1:]
		for _, p := range f.Blocks[b].Preds {
			if body[p] {
				continue
			}
			body[p] = true
			worklist = append(worklist, p)
		}
	}
	return body
}

// findPreheader returns the unique predecessor of header that's NOT in
// the loop body. Returns -1 if there are zero or multiple such preds.
// (When -1, LICM skips this loop — synthesizing a pre-header is future
// work.)
func findPreheader(f *ir.Function, header ir.BlockID, body map[ir.BlockID]bool) ir.BlockID {
	var pre ir.BlockID = -1
	count := 0
	for _, p := range f.Blocks[header].Preds {
		if body[p] {
			continue
		}
		count++
		pre = p
	}
	if count != 1 {
		return -1
	}
	return pre
}

// hoistInvariants moves loop-invariant nodes from the loop body to the
// pre-header. Returns true if any node was moved.
func hoistInvariants(
	f *ir.Function,
	idom []ir.BlockID,
	loop loopInfo,
	mutatedVars map[*vm.Var]bool,
	allMutated bool,
) (changed bool) {
	if loop.preheader < 0 {
		return false // can't hoist without a pre-header
	}

	// Mark nodes as invariant in fixed-point fashion. Walk loop body
	// nodes; a node is invariant if all its operands are invariant.
	// "Invariant" for a node means: defined outside the loop body, OR
	// pure/constant with all-invariant operands.
	invariant := make(map[ir.InstId]bool)

	// Seed: any node defined OUTSIDE the loop body is invariant.
	for i := range f.Insts {
		n := &f.Insts[i]
		if n.Op == ir.OpInvalid {
			continue
		}
		if !loop.body[n.Block] {
			invariant[ir.InstId(i)] = true
		}
	}

	// Iterate to fixed point: a body node becomes invariant when all
	// its refs are invariant AND it's pure (or a safe non-pure case
	// like LoadVar of a non-mutated var).
	for progress := true; progress; {
		progress = false
		for bid := range loop.body {
			blk := &f.Blocks[bid]
			for _, nid := range blk.Insts {
				if invariant[nid] {
					continue
				}
				n := f.Inst(nid)
				if !nodeHoistable(n, mutatedVars, allMutated) {
					continue
				}
				allRefsInvariant := true
				for _, r := range n.Refs {
					if !invariant[r] {
						allRefsInvariant = false
						break
					}
				}
				if allRefsInvariant {
					invariant[nid] = true
					progress = true
				}
			}
		}
	}

	// Now move the hoistable body nodes to the pre-header. Walk each
	// body block, filter its Insts list, and append hoisted ones to
	// the pre-header (in the order they were encountered to preserve
	// def-before-use within the pre-header).
	var hoisted []ir.InstId
	for bid := range loop.body {
		blk := &f.Blocks[bid]
		kept := blk.Insts[:0]
		for _, nid := range blk.Insts {
			if !invariant[nid] || f.Insts[nid].Block == loop.preheader {
				kept = append(kept, nid)
				continue
			}
			// Confirm this node was defined inside the loop body — we
			// only hoist body-defined nodes, not nodes that were seeded
			// as invariant because they're already outside.
			if loop.body[f.Insts[nid].Block] {
				hoisted = append(hoisted, nid)
				continue
			}
			kept = append(kept, nid)
		}
		blk.Insts = kept
	}
	if len(hoisted) == 0 {
		return false
	}

	// Re-block hoisted nodes to the pre-header. Insert them BEFORE the
	// pre-header's terminator (since we don't want to break terminator
	// position). Since blk.Insts excludes the terminator, we append.
	preBlock := &f.Blocks[loop.preheader]
	for _, nid := range hoisted {
		f.Insts[nid].Block = loop.preheader
		preBlock.Insts = append(preBlock.Insts, nid)
	}

	// CRITICAL: hoisted values are now defined in the pre-header but
	// may be used inside the loop body. The Lower pass rejects cross-
	// block uses via direct Refs — values must flow through BlockArgs.
	//
	// For each hoisted value used in the loop, add a new block-param
	// to the loop header, rewire uses inside the loop to that param,
	// and update all branches entering the header to pass the hoisted
	// value as the new arg.
	threadHoistedThroughHeader(f, loop, hoisted)

	return true
}

// threadHoistedThroughHeader adds new BlockArg parameters to the loop
// header for each hoisted value that's used inside the loop body, then
// rewires consumers inside the loop to reference the new params. All
// branches entering the header — both the pre-header→header edge and
// any back-edges from within the loop — must also pass the hoisted
// value (it's loop-invariant so it's the same value on every iteration).
func threadHoistedThroughHeader(f *ir.Function, loop loopInfo, hoisted []ir.InstId) {
	if len(hoisted) == 0 {
		return
	}
	// Per hoisted value: only thread it if at least one user is in the
	// loop body. (If no loop-body user, it's been hoisted and could've
	// just been DCE'd; but it's still legal to leave it.)
	header := &f.Blocks[loop.header]
	type threading struct {
		hoistedID ir.InstId // original hoisted node, in pre-header
		paramID   ir.InstId // new BlockArg in header that receives hoistedID
	}
	var threads []threading

	for _, hID := range hoisted {
		usedInLoop := false
		for i := range f.Insts {
			n := &f.Insts[i]
			if n.Op == ir.OpInvalid {
				continue
			}
			if !loop.body[n.Block] {
				continue
			}
			for _, r := range n.Refs {
				if r == hID {
					usedInLoop = true
					break
				}
			}
			if usedInLoop {
				break
			}
		}
		if !usedInLoop {
			continue
		}
		// Create a new BlockArg in the header.
		paramID := f.AddNode(ir.Inst{
			Op:    ir.OpBlockArg,
			Aux:   len(header.Params),
			Block: loop.header,
		})
		header.Params = append(header.Params, paramID)
		threads = append(threads, threading{hoistedID: hID, paramID: paramID})
	}

	if len(threads) == 0 {
		return
	}

	// Rewire all loop-body uses of each hoisted value to the new param.
	// Also update branch-target Args within the loop body — when a
	// branch leaves the loop, if it referenced a hoisted value as
	// arg, it should now reference the corresponding param (which
	// holds the same value as the hoisted def).
	for _, t := range threads {
		for i := range f.Insts {
			n := &f.Insts[i]
			if n.Op == ir.OpInvalid {
				continue
			}
			// Skip the hoisted node itself.
			if ir.InstId(i) == t.hoistedID {
				continue
			}
			// Only rewrite uses inside the loop.
			if !loop.body[n.Block] {
				continue
			}
			for j, r := range n.Refs {
				if r == t.hoistedID {
					n.Refs[j] = t.paramID
				}
			}
			// Branch-target args of terminators in the loop body might
			// also reference the hoisted value.
			switch ta := n.Aux.(type) {
			case *ir.BranchTarget:
				if ta != nil {
					for j, a := range ta.Args {
						if a == t.hoistedID {
							ta.Args[j] = t.paramID
						}
					}
				}
			case *ir.CondTarget:
				if ta != nil {
					if ta.True != nil {
						for j, a := range ta.True.Args {
							if a == t.hoistedID {
								ta.True.Args[j] = t.paramID
							}
						}
					}
					if ta.False != nil {
						for j, a := range ta.False.Args {
							if a == t.hoistedID {
								ta.False.Args[j] = t.paramID
							}
						}
					}
				}
			}
		}
	}

	// Every branch entering the header must now pass the hoisted value
	// as an additional arg. This includes the pre-header→header edge
	// AND any back-edges from within the loop body. For pre-header→
	// header, the arg is the hoistedID (defined in pre-header). For
	// back-edges (defined inside loop), the value flowing back is the
	// header's own param — it's loop-invariant.
	for _, p := range header.Preds {
		predTerm := &f.Insts[f.Blocks[p].Term]
		isFromLoop := loop.body[p]
		switch ta := predTerm.Aux.(type) {
		case *ir.BranchTarget:
			if ta != nil && ta.Target == loop.header {
				for _, t := range threads {
					if isFromLoop {
						ta.Args = append(ta.Args, t.paramID)
					} else {
						ta.Args = append(ta.Args, t.hoistedID)
					}
				}
			}
		case *ir.CondTarget:
			if ta != nil {
				if ta.True != nil && ta.True.Target == loop.header {
					for _, t := range threads {
						if isFromLoop {
							ta.True.Args = append(ta.True.Args, t.paramID)
						} else {
							ta.True.Args = append(ta.True.Args, t.hoistedID)
						}
					}
				}
				if ta.False != nil && ta.False.Target == loop.header {
					for _, t := range threads {
						if isFromLoop {
							ta.False.Args = append(ta.False.Args, t.paramID)
						} else {
							ta.False.Args = append(ta.False.Args, t.hoistedID)
						}
					}
				}
			}
		}
	}
}

// nodeHoistable reports whether n's effect is reproducible across
// invocations (no side effects observable to the program). Pure ops
// are always hoistable. LoadVar is hoistable only when the target var
// isn't mutated anywhere in the function.
//
// OpCall, OpTailCall, OpSetVar, and terminators are never hoisted.
func nodeHoistable(n *ir.Inst, mutatedVars map[*vm.Var]bool, allMutated bool) bool {
	if n.Op.IsTerminator() {
		return false
	}
	if n.Op.IsPure() {
		return true
	}
	// LoadVar of a stable var is effectively invariant.
	if n.Op == ir.OpLoadVar && !allMutated {
		if v, ok := n.Aux.(vm.Value); ok {
			if vv, ok := v.(*vm.Var); ok {
				return !mutatedVars[vv]
			}
		}
	}
	return false
}
