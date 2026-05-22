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

// TestCanonicalize_AddConstSecond: (+ 5 x) should be rewritten to (+ x 5)
// so that CSE can recognize (+ 5 x) and (+ x 5) as the same.
func TestCanonicalize_AddConstSecond(t *testing.T) {
	consts := vm.NewConsts()
	f := ir.NewFunction("c1", 1, false, consts)
	b0 := f.Entry
	arg := f.AddNode(ir.Inst{Op: ir.OpLoadArg, Aux: 0, Block: b0})
	f.AppendToBlock(b0, arg)
	c5 := f.AddNode(ir.Inst{Op: ir.OpConst, Aux: vm.Int(5), Block: b0})
	f.AppendToBlock(b0, c5)
	// (+ 5 x) — Const first, the form we want to canonicalize away from.
	add := f.AddNode(ir.Inst{Op: ir.OpAdd, Refs: []ir.InstId{c5, arg}, Block: b0})
	f.AppendToBlock(b0, add)
	ret := f.AddNode(ir.Inst{Op: ir.OpReturn, Refs: []ir.InstId{add}, Block: b0})
	f.SetTerminator(b0, ret)

	if !ConstFold(f) {
		t.Fatal("expected ConstFold to canonicalize the operand order")
	}
	addNode := f.Inst(add)
	if len(addNode.Refs) != 2 {
		t.Fatalf("add node ended up with %d refs", len(addNode.Refs))
	}
	if addNode.Refs[0] != arg || addNode.Refs[1] != c5 {
		t.Errorf("expected (+ x 5) after canonicalization, got refs %v (x=%d c5=%d)",
			addNode.Refs, arg, c5)
	}
}

// TestCanonicalize_CSEAcrossOrder: (+ 5 x) and (+ x 5) should CSE to
// the same value after canonicalization. Without canonicalization, CSE
// would treat them as distinct (different operand sequences).
func TestCanonicalize_CSEAcrossOrder(t *testing.T) {
	src := `(defn cse-add [x] (+ (+ x 5) (+ 5 x)))`
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		t.Fatalf("compile: %v", err)
	}
	v := ns.Lookup(vm.Symbol("cse-add")).(*vm.Var).Deref().(*vm.Func)
	irFn, err := ir.Build(v.Chunk(), "cse-add", 1, false)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	// Run ConstFold (canonicalizes) then CSE (merges duplicates).
	ConstFold(irFn)
	CSE(irFn)
	// After CSE, there should be exactly ONE OpAdd computing (+ x 5),
	// and a second OpAdd that adds the result to itself. Without
	// canonicalization there would be two distinct OpAdd nodes for
	// (+ x 5) and (+ 5 x).
	addCount := 0
	for _, n := range irFn.Insts {
		if n.Op == ir.OpAdd {
			addCount++
		}
	}
	if addCount > 2 {
		t.Errorf("expected ≤2 OpAdd nodes after canonicalize+CSE (one (+ x 5), one outer), got %d", addCount)
	}
}

// TestCanonicalize_MulConstSecond: same pattern for multiplication.
func TestCanonicalize_MulConstSecond(t *testing.T) {
	consts := vm.NewConsts()
	f := ir.NewFunction("c2", 1, false, consts)
	b0 := f.Entry
	arg := f.AddNode(ir.Inst{Op: ir.OpLoadArg, Aux: 0, Block: b0})
	f.AppendToBlock(b0, arg)
	c3 := f.AddNode(ir.Inst{Op: ir.OpConst, Aux: vm.Int(3), Block: b0})
	f.AppendToBlock(b0, c3)
	mul := f.AddNode(ir.Inst{Op: ir.OpMul, Refs: []ir.InstId{c3, arg}, Block: b0})
	f.AppendToBlock(b0, mul)
	ret := f.AddNode(ir.Inst{Op: ir.OpReturn, Refs: []ir.InstId{mul}, Block: b0})
	f.SetTerminator(b0, ret)

	if !ConstFold(f) {
		t.Fatal("expected canonicalization to swap (* 3 x) operands")
	}
	mulNode := f.Inst(mul)
	if mulNode.Refs[0] != arg || mulNode.Refs[1] != c3 {
		t.Errorf("expected (* x 3), got refs %v", mulNode.Refs)
	}
}

// TestCanonicalize_NonCommutativeUnchanged: (- 5 x) should NOT be
// canonicalized — subtraction isn't commutative.
func TestCanonicalize_NonCommutativeUnchanged(t *testing.T) {
	consts := vm.NewConsts()
	f := ir.NewFunction("c3", 1, false, consts)
	b0 := f.Entry
	arg := f.AddNode(ir.Inst{Op: ir.OpLoadArg, Aux: 0, Block: b0})
	f.AppendToBlock(b0, arg)
	c5 := f.AddNode(ir.Inst{Op: ir.OpConst, Aux: vm.Int(5), Block: b0})
	f.AppendToBlock(b0, c5)
	sub := f.AddNode(ir.Inst{Op: ir.OpSub, Refs: []ir.InstId{c5, arg}, Block: b0})
	f.AppendToBlock(b0, sub)
	ret := f.AddNode(ir.Inst{Op: ir.OpReturn, Refs: []ir.InstId{sub}, Block: b0})
	f.SetTerminator(b0, ret)

	ConstFold(f)
	subNode := f.Inst(sub)
	if subNode.Refs[0] != c5 || subNode.Refs[1] != arg {
		t.Errorf("expected (- 5 x) preserved, got refs %v", subNode.Refs)
	}
}

// TestCanonicalize_BothConstSkipped: (+ 5 3) has both operands as
// Const; primitive fold should collapse it. Canonicalization should
// not interfere.
func TestCanonicalize_BothConstSkipped(t *testing.T) {
	consts := vm.NewConsts()
	f := ir.NewFunction("c4", 0, false, consts)
	b0 := f.Entry
	c5 := f.AddNode(ir.Inst{Op: ir.OpConst, Aux: vm.Int(5), Block: b0})
	f.AppendToBlock(b0, c5)
	c3 := f.AddNode(ir.Inst{Op: ir.OpConst, Aux: vm.Int(3), Block: b0})
	f.AppendToBlock(b0, c3)
	add := f.AddNode(ir.Inst{Op: ir.OpAdd, Refs: []ir.InstId{c5, c3}, Block: b0})
	f.AppendToBlock(b0, add)
	ret := f.AddNode(ir.Inst{Op: ir.OpReturn, Refs: []ir.InstId{add}, Block: b0})
	f.SetTerminator(b0, ret)

	ConstFold(f)
	addNode := f.Inst(add)
	if addNode.Op != ir.OpConst {
		t.Errorf("expected (+ 5 3) folded to OpConst, got %s", addNode.Op)
	}
	if v, _ := addNode.Aux.(vm.Value); v.String() != "8" {
		t.Errorf("expected folded value 8, got %v", addNode.Aux)
	}
}

// TestCanonicalize_BitAndCommutative: (bit-and 0xFF x) → (bit-and x 0xFF).
func TestCanonicalize_BitAndCommutative(t *testing.T) {
	src := `(defn bit-test [x] (bit-and 255 x))`
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		t.Fatalf("compile: %v", err)
	}
	v := ns.Lookup(vm.Symbol("bit-test")).(*vm.Var).Deref().(*vm.Func)
	irFn, err := ir.Build(v.Chunk(), "bit-test", 1, false)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	ConstFold(irFn)
	// Find the OpCall to bit-and.
	for _, n := range irFn.Insts {
		if n.Op != ir.OpCall {
			continue
		}
		fnNode := irFn.Inst(n.Refs[0])
		if fnNode.Op != ir.OpLoadVar {
			continue
		}
		varVal, _ := fnNode.Aux.(vm.Value)
		variable, _ := varVal.(*vm.Var)
		if variable == nil || variable.VarName() != "bit-and" {
			continue
		}
		// Args are at Refs[1] and Refs[2]. After canonicalization
		// Refs[1] should be the non-Const (x), Refs[2] should be Const (255).
		lhsOp := irFn.Inst(n.Refs[1]).Op
		rhsOp := irFn.Inst(n.Refs[2]).Op
		if lhsOp == ir.OpConst && rhsOp != ir.OpConst {
			t.Errorf("bit-and not canonicalized: still has (bit-and Const Var)")
		}
		return
	}
	t.Skip("bit-and call not found in IR; compiler may have folded it differently")
}
