/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir_test

// Slice A of typed-direct-call lowering: the IR type lattice gains concrete
// collection kinds (:vector/:map/:set/:list/:cons/:lazyseq) and derives vm
// interface satisfaction (Sequable/Counted/Seq) as a MUST query over the
// concrete kinds a value may be — kinds(x) ⊆ <iface>-kinds. Nullability is
// orthogonal: :nil is a separate lattice member, so a value is provably
// non-nil iff :nil is not among its members.
//
// These are behavior-neutral additions (nothing PRODUCES collection kinds yet;
// that is Slice B), so existing type-join is unchanged — collection kinds are
// non-maskable keyword types, like :ratio, and the bitset proptest is untouched.

import (
	"testing"
)

func TestLatticeCollectionKindsAndInterfaceSatisfaction(t *testing.T) {
	ensureLoader()
	got := runLispString(t, `(pr-str
		[;; collection kinds join like any other concrete kind (union)
		 (ir.lattice/type-join :vector :map)
		 (ir.lattice/normalize-type [:union :vector :vector])
		 ;; Sequable: every concrete kind must implement it
		 (ir.lattice/seqable-type? :vector)
		 (ir.lattice/seqable-type? :string)
		 (ir.lattice/seqable-type? :int)
		 (ir.lattice/seqable-type? [:union :vector :string])
		 (ir.lattice/seqable-type? [:union :vector :int])
		 ;; nullability orthogonal: vector|nil is STILL Sequable (nil ignored)
		 (ir.lattice/seqable-type? [:union :vector :nil])
		 ;; Counted
		 (ir.lattice/counted-type? :cons)
		 (ir.lattice/counted-type? :int)
		 ;; nullability is orthogonal — a separate member
		 (ir.lattice/non-nil-type? :vector)
		 (ir.lattice/non-nil-type? [:union :vector :nil])
		 (ir.lattice/non-nil-type? :nil)])`)
	want := `[[:union :vector :map] :vector true true false true false true true false true false false]`
	if got != want {
		t.Fatalf("lattice interface-satisfaction mismatch:\n got: %s\nwant: %s", got, want)
	}

	// Slice B source: classify-const types collection literals as their kind, so
	// a value flowing from a literal carries the interfaces it satisfies.
	cc := runLispString(t, `(pr-str [(ir.passes.typeinfer/classify-const [1 2 3])
	                                  (ir.passes.typeinfer/classify-const {:a 1})
	                                  (ir.passes.typeinfer/classify-const #{1 2})
	                                  (ir.passes.typeinfer/classify-const 7)])`)
	wantCC := `[:vector :map :set :int]`
	if cc != wantCC {
		t.Fatalf("classify-const mismatch:\n got: %s\nwant: %s", cc, wantCC)
	}
}
