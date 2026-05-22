/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir

import (
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/compiler"
	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

// TestDominators_StraightLine: a function with no branches has trivial
// dominance — each block's idom is its predecessor (or -1 for entry).
func TestDominators_StraightLine(t *testing.T) {
	consts := vm.NewConsts()
	f := NewFunction("straight", 0, false, consts)
	b0 := f.Entry
	b1 := f.AddBlock()
	b2 := f.AddBlock()
	// Add a sentinel node first so subsequent NodeIDs are non-zero.
	// (The IR uses InstId(0) as "no terminator" sentinel.)
	_ = f.AddNode(Inst{Op: OpInvalid, Block: b0})
	// b0 -> b1 -> b2
	br0 := f.AddNode(Inst{Op: OpBranch, Aux: &BranchTarget{Target: b1}, Block: b0})
	f.SetTerminator(b0, br0)
	f.AddPred(b1, b0)
	br1 := f.AddNode(Inst{Op: OpBranch, Aux: &BranchTarget{Target: b2}, Block: b1})
	f.SetTerminator(b1, br1)
	f.AddPred(b2, b1)
	ret := f.AddNode(Inst{Op: OpReturn, Block: b2})
	f.SetTerminator(b2, ret)

	idom := f.Dominators()
	if idom[b0] != -1 {
		t.Errorf("idom(b0)=%d, want -1", idom[b0])
	}
	if idom[b1] != b0 {
		t.Errorf("idom(b1)=%d, want b0=%d", idom[b1], b0)
	}
	if idom[b2] != b1 {
		t.Errorf("idom(b2)=%d, want b1=%d", idom[b2], b1)
	}
}

// TestDominators_IfElseJoin: diamond CFG. b0 branches to b1 or b2, both
// join at b3. b0 dominates all; b1 doesn't dominate b3 (b2 alt path).
func TestDominators_IfElseJoin(t *testing.T) {
	consts := vm.NewConsts()
	f := NewFunction("ifelse", 0, false, consts)
	b0 := f.Entry
	b1 := f.AddBlock()
	b2 := f.AddBlock()
	b3 := f.AddBlock()

	// Sentinel: InstId(0) reserved as "no terminator".
	_ = f.AddNode(Inst{Op: OpInvalid, Block: b0})
	// b0: BranchIf cond -> b1 : b2
	cond := f.AddNode(Inst{Op: OpConst, Aux: vm.Boolean(true), Block: b0})
	f.AppendToBlock(b0, cond)
	br0 := f.AddNode(Inst{Op: OpBranchIf, Refs: []InstId{cond}, Aux: &CondTarget{
		True:  &BranchTarget{Target: b1},
		False: &BranchTarget{Target: b2},
	}, Block: b0})
	f.SetTerminator(b0, br0)
	f.AddPred(b1, b0)
	f.AddPred(b2, b0)

	// b1 -> b3
	br1 := f.AddNode(Inst{Op: OpBranch, Aux: &BranchTarget{Target: b3}, Block: b1})
	f.SetTerminator(b1, br1)
	f.AddPred(b3, b1)

	// b2 -> b3
	br2 := f.AddNode(Inst{Op: OpBranch, Aux: &BranchTarget{Target: b3}, Block: b2})
	f.SetTerminator(b2, br2)
	f.AddPred(b3, b2)

	// b3: return
	ret := f.AddNode(Inst{Op: OpReturn, Block: b3})
	f.SetTerminator(b3, ret)

	idom := f.Dominators()
	if idom[b0] != -1 {
		t.Errorf("idom(b0)=%d, want -1", idom[b0])
	}
	if idom[b1] != b0 {
		t.Errorf("idom(b1)=%d, want b0", idom[b1])
	}
	if idom[b2] != b0 {
		t.Errorf("idom(b2)=%d, want b0", idom[b2])
	}
	if idom[b3] != b0 {
		t.Errorf("idom(b3)=%d, want b0 (join — b1 and b2 both reach b3)", idom[b3])
	}

	// Dominance relations.
	if !Dominates(idom, b0, b3) {
		t.Errorf("b0 should dominate b3")
	}
	if Dominates(idom, b1, b3) {
		t.Errorf("b1 should NOT dominate b3")
	}
}

// TestDominators_RealLoop: compile sumto and verify the loop header's
// dominance. The loop header should dominate the loop body (which
// contains the back-edge).
func TestDominators_RealLoop(t *testing.T) {
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	src := `(defn sumdom [n] (loop [i 0 acc 0] (if (< i n) (recur (inc i) (+ acc i)) acc)))`
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		t.Fatalf("compile: %v", err)
	}
	v := ns.Lookup(vm.Symbol("sumdom")).(*vm.Var).Deref()
	fn := v.(*vm.Func)
	irFn, err := Build(fn.Chunk(), "sumdom", 1, false)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}

	idom := irFn.Dominators()
	t.Logf("idom: %v", idom)
	t.Logf("IR:\n%s", Dump(irFn))

	// Verify: every reachable block (except Entry) has an idom that's
	// some block ID. Entry's idom is -1.
	if idom[irFn.Entry] != -1 {
		t.Errorf("idom(entry)=%d, want -1", idom[irFn.Entry])
	}
	for bi := range irFn.Blocks {
		bid := BlockID(bi)
		if bid == irFn.Entry {
			continue
		}
		if idom[bid] < 0 {
			t.Errorf("block %d has idom=%d (no dominator)", bid, idom[bid])
		}
		// idom must itself be a block in this function.
		if int(idom[bid]) >= len(irFn.Blocks) {
			t.Errorf("block %d idom=%d is out of range", bid, idom[bid])
		}
	}

	// The loop header (multi-pred block) should dominate any block on
	// the loop's back-edge path. Verify by finding the loop header and
	// checking dominance relations.
	var header BlockID = -1
	for bi, blk := range irFn.Blocks {
		if len(blk.Preds) >= 2 {
			header = BlockID(bi)
			break
		}
	}
	if header == -1 {
		t.Skip("no loop header found (multi-pred block); test fixture may have shifted")
		return
	}
	// The loop header should dominate at least one of its own predecessors
	// (the one carrying the back-edge).
	dominatesPred := false
	for _, p := range irFn.Blocks[header].Preds {
		if Dominates(idom, header, p) {
			dominatesPred = true
			break
		}
	}
	if !dominatesPred {
		t.Errorf("loop header %d doesn't dominate any of its predecessors %v — back-edge structure unusual", header, irFn.Blocks[header].Preds)
	}
}

// TestDominators_Unreachable: a block with no path from Entry should
// not be in the dominator tree (idom stays -1). Build prunes
// unreachable blocks, so this test constructs by hand.
func TestDominators_Unreachable(t *testing.T) {
	consts := vm.NewConsts()
	f := NewFunction("unreach", 0, false, consts)
	b0 := f.Entry
	b1 := f.AddBlock()
	b2 := f.AddBlock() // unreachable

	// Sentinel: InstId(0) reserved.
	_ = f.AddNode(Inst{Op: OpInvalid, Block: b0})
	// b0 -> b1 directly. b2 has no preds (unreachable).
	br := f.AddNode(Inst{Op: OpBranch, Aux: &BranchTarget{Target: b1}, Block: b0})
	f.SetTerminator(b0, br)
	f.AddPred(b1, b0)
	ret1 := f.AddNode(Inst{Op: OpReturn, Block: b1})
	f.SetTerminator(b1, ret1)
	ret2 := f.AddNode(Inst{Op: OpReturn, Block: b2})
	f.SetTerminator(b2, ret2)

	idom := f.Dominators()
	if idom[b2] != -1 {
		t.Errorf("unreachable b2 should have idom=-1, got %d", idom[b2])
	}
}
