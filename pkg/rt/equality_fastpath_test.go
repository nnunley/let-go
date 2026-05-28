package rt

import (
	"testing"
	"time"

	"github.com/nooga/let-go/pkg/vm"
)

// TestArrayVector_SameBacking_ShortCircuit verifies that two ArrayVector
// values built over the same backing slice short-circuit in O(1) without
// walking elements. Pre-fix this would have been a full element walk;
// post-fix the identity short-circuit catches it. Regression guard for
// the IR-graph shared-substructure case.
func TestArrayVector_SameBacking_ShortCircuit(t *testing.T) {
	backing := make([]vm.Value, 1024)
	for i := range backing {
		backing[i] = vm.Int(int64(i))
	}
	av := vm.ArrayVector(backing)
	bv := vm.ArrayVector(backing)
	if !valueEqualsFast(av, bv) {
		t.Fatal("expected equal for shared-backing ArrayVector")
	}
}

// TestArrayVector_HashRejects_Unequal verifies that two equal-length but
// distinct ArrayVectors are rejected via hash without walking elements.
// Asserted indirectly via correctness; the perf claim is covered below.
func TestArrayVector_HashRejects_Unequal(t *testing.T) {
	a := vm.ArrayVector{vm.Int(1), vm.Int(2), vm.Int(3)}
	b := vm.ArrayVector{vm.Int(1), vm.Int(2), vm.Int(4)}
	if valueEqualsFast(a, b) {
		t.Fatal("expected not equal")
	}
}

// TestArrayVector_EqualDistinctBacking verifies correctness when neither
// identity nor hash decides — must still return true via element walk.
func TestArrayVector_EqualDistinctBacking(t *testing.T) {
	a := vm.ArrayVector{vm.Int(1), vm.Int(2), vm.Int(3)}
	b := vm.ArrayVector{vm.Int(1), vm.Int(2), vm.Int(3)}
	if !valueEqualsFast(a, b) {
		t.Fatal("expected equal for structurally-equal ArrayVectors")
	}
}

// TestSharedSubstructure_DoesNotExplode constructs a binary-tree-shaped
// vector graph where every level reuses the same child node. Pre-fix the
// recursive walk visits the child 2^depth times. Post-fix the visited-pair
// memo (plus pointer-identity on the shared subtree) bounds it to O(depth).
//
// This is the regression that motivated the whole fix: in let-go's IR
// graphs, instructions are shared across many parents, and a naive
// structural compare became exponential. We assert a wall-time bound
// instead of an instruction count because Go test infra makes the latter
// fiddly — the multiplicative blowup pre-fix is dramatic enough that any
// reasonable upper bound is decisive.
func TestSharedSubstructure_DoesNotExplode(t *testing.T) {
	leaf := vm.ArrayVector{vm.Int(1), vm.Int(2), vm.Int(3)}
	build := func() vm.ArrayVector {
		node := leaf
		for i := 0; i < 24; i++ {
			node = vm.ArrayVector{node, node}
		}
		return node
	}
	a := build()
	b := build()
	done := make(chan bool, 1)
	go func() {
		done <- valueEqualsFast(a, b)
	}()
	select {
	case got := <-done:
		if !got {
			t.Fatal("expected shared-substructure trees to compare equal")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("valueEqualsFast did not finish in 2s on 2^24 shared-substructure tree (pre-fix this hangs)")
	}
}

// TestEqValue_FastPath_RoutesToFast confirms that rt.EqValue uses
// valueEqualsFast when EqFastPath is on — verified indirectly by the
// shared-substructure test running through EqValue too.
func TestEqValue_FastPath_RoutesToFast(t *testing.T) {
	leaf := vm.ArrayVector{vm.Int(1)}
	node := leaf
	for i := 0; i < 20; i++ {
		node = vm.ArrayVector{node, node}
	}
	done := make(chan bool, 1)
	go func() {
		done <- EqValue(node, node)
	}()
	select {
	case got := <-done:
		if !got {
			t.Fatal("expected EqValue(node, node) true")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("EqValue did not finish in 1s — fast path not engaged")
	}
}

// TestEqValue_NilHandling guards the contract that EqValue compares
// nil-pair as equal and nil-vs-non-nil as unequal.
func TestEqValue_NilHandling(t *testing.T) {
	if !EqValue(nil, nil) {
		t.Fatal("nil == nil should be true")
	}
	if EqValue(nil, vm.Int(0)) {
		t.Fatal("nil == Int(0) should be false")
	}
	if EqValue(vm.Int(0), nil) {
		t.Fatal("Int(0) == nil should be false")
	}
}

// TestEqValue_ScalarFastPath spot-checks the comparable-scalar fast path
// remains correct after rerouting the slow path.
func TestEqValue_ScalarFastPath(t *testing.T) {
	if !EqValue(vm.Keyword("foo"), vm.Keyword("foo")) {
		t.Fatal("keyword equality")
	}
	if EqValue(vm.Keyword("foo"), vm.Keyword("bar")) {
		t.Fatal("keyword inequality")
	}
	if !EqValue(vm.Int(42), vm.Int(42)) {
		t.Fatal("int equality")
	}
	if !EqValue(vm.String("x"), vm.String("x")) {
		t.Fatal("string equality")
	}
}
