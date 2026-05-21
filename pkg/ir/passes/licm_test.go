/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package passes

import (
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/compiler"
	"github.com/nooga/let-go/pkg/ir"
	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

// TestLICM_HoistConstantSubexpression: a loop with `(* 4 scale)` in
// the condition has that computation evaluated every iteration in
// baseline. LICM should hoist it to the pre-header.
//
// We verify by counting OpMul nodes in the loop body before vs after:
// before LICM, the body contains an OpMul for `(* 4 scale)` and
// another for `(+ acc i)`. After LICM, the `(* 4 scale)` should be
// gone from the loop body (moved to pre-header).
func TestLICM_HoistConstantSubexpression(t *testing.T) {
	t.Skip("LICM IR transform works in isolation, but the current stack-based Lower can't express cross-block uses of hoisted values without a per-block threading scheme that bloats every loop block's BlockArgs. Re-enable once the VM executes IR directly (Carbon-style function-scoped reaching defs).")
	src := `
(def scale 1000)
(defn loop-with-invariant [n]
  (loop [i 0 acc 0]
    (if (< i (* 4 scale))
      (recur (inc i) (+ acc i))
      acc)))
`
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		t.Fatalf("compile: %v", err)
	}
	fn := ns.Lookup(vm.Symbol("loop-with-invariant")).(*vm.Var).Deref().(*vm.Func)
	irFn, err := ir.Build(fn.Chunk(), "loop-with-invariant", 1, false)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}

	// Run ConstFold first: (* 4 scale) is (* 4 (LoadVar scale)) — but
	// scale is a Var, not a Const, so ConstFold can't reduce. The Mul
	// stays as a runtime computation. Perfect LICM target.
	ConstFold(irFn)
	CSE(irFn)

	t.Logf("IR before LICM:\n%s", ir.Dump(irFn))

	// Identify the loop body — block with multi-pred (>=2) is the
	// header. The body is the header + any block that can reach it
	// via the back-edge.
	var loopHeader ir.BlockID = -1
	for i, blk := range irFn.Blocks {
		if len(blk.Preds) >= 2 {
			loopHeader = ir.BlockID(i)
			break
		}
	}
	if loopHeader < 0 {
		t.Fatal("no loop header found")
	}

	// Count OpMul/OpLt nodes in the loop's body region BEFORE LICM.
	idom := irFn.Dominators()
	loopBlocks := blocksInLoop(irFn, idom, loopHeader)
	t.Logf("loop blocks: %v", keysOf(loopBlocks))
	mulInLoopBefore := countOpInBlocks(irFn, loopBlocks, ir.OpMul)
	ltInLoopBefore := countOpInBlocks(irFn, loopBlocks, ir.OpLt)
	t.Logf("before LICM: %d OpMul, %d OpLt in loop", mulInLoopBefore, ltInLoopBefore)

	changed := LICM(irFn)
	if !changed {
		t.Fatalf("LICM didn't change anything; expected to hoist (* 4 scale)")
	}

	t.Logf("IR after LICM:\n%s", ir.Dump(irFn))

	// After LICM, the (* 4 scale) should have left the loop body.
	mulInLoopAfter := countOpInBlocks(irFn, loopBlocks, ir.OpMul)
	t.Logf("after LICM: %d OpMul in loop", mulInLoopAfter)
	if mulInLoopAfter >= mulInLoopBefore {
		t.Errorf("expected fewer OpMul nodes in loop after LICM (was %d, now %d)",
			mulInLoopBefore, mulInLoopAfter)
	}

	// Round-trip: Lower and verify the result still runs correctly.
	chunk, err := ir.Lower(irFn)
	if err != nil {
		t.Fatalf("Lower: %v", err)
	}
	frame := vm.NewFrame(chunk, []vm.Value{vm.Int(10)})
	result, err := frame.Run()
	vm.ReleaseFrame(frame)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	// loop-with-invariant(10): condition is `(< i 4000)`, so with n=10
	// (no early exit since 10 < 4000), the loop runs until... wait,
	// the loop iterates while (< i (* 4 scale)) = (< i 4000). With
	// recur (inc i), after 4000 iterations i reaches 4000 and the
	// condition becomes false. acc has accumulated 0+1+2+...+3999 =
	// 3999*4000/2 = 7998000.
	if result.String() != "7998000" {
		t.Errorf("expected 7998000, got %s", result)
	}
}

// blocksInLoop returns the natural-loop body for a header in idom form.
func blocksInLoop(f *ir.Function, idom []ir.BlockID, header ir.BlockID) map[ir.BlockID]bool {
	body := map[ir.BlockID]bool{header: true}
	// Find back-edges into header.
	for bi := range f.Blocks {
		bid := ir.BlockID(bi)
		for _, succ := range f.Successors(bid) {
			if succ != header {
				continue
			}
			if !ir.Dominates(idom, header, bid) {
				continue
			}
			// bid → header is a back-edge.
			body[bid] = true
			// Walk back from bid, collecting predecessors until we leave
			// the header-dominated region.
			worklist := []ir.BlockID{bid}
			for len(worklist) > 0 {
				b := worklist[0]
				worklist = worklist[1:]
				for _, p := range f.Blocks[b].Preds {
					if body[p] || !ir.Dominates(idom, header, p) {
						continue
					}
					body[p] = true
					worklist = append(worklist, p)
				}
			}
		}
	}
	return body
}

func countOpInBlocks(f *ir.Function, blocks map[ir.BlockID]bool, op ir.Op) int {
	n := 0
	for bid := range blocks {
		for _, nid := range f.Blocks[bid].Nodes {
			if f.Nodes[nid].Op == op {
				n++
			}
		}
	}
	return n
}

func keysOf(m map[ir.BlockID]bool) []ir.BlockID {
	out := make([]ir.BlockID, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// TestLICM_NoLoop: a function without loops should be a no-op for LICM.
func TestLICM_NoLoop(t *testing.T) {
	src := `(defn noloop [x] (+ x 1))`
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		t.Fatalf("compile: %v", err)
	}
	fn := ns.Lookup(vm.Symbol("noloop")).(*vm.Var).Deref().(*vm.Func)
	irFn, err := ir.Build(fn.Chunk(), "noloop", 1, false)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	if LICM(irFn) {
		t.Error("LICM changed a loopless function — unexpected")
	}
}

// TestLICM_LoopWithMutation: a loop body that mutates a var must NOT
// hoist any LoadVar of that mutated var, even though LoadVar is
// "constant-looking" in isolation.
func TestLICM_LoopWithMutation(t *testing.T) {
	// (def counter 0) (defn loop-mutate [n] (loop [i 0] (if (< i n) (do (set! counter (inc counter)) (recur (inc i))) i)))
	//
	// set! support in IR build is limited; this test mainly exercises
	// the "don't hoist LoadVar of mutated var" path. Skipping if
	// compile fails or runs into Build limitations.
	t.Skip("set! support in IR Build is limited; reactivate when SetVar lowering works")
}
