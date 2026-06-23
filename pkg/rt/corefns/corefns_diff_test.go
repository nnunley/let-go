/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package corefns_test

// Differential tests: each lifted corefns export MUST behave identically to
// the canonical clojure.core builtin it mirrors. The lift duplicates the
// closure body from pkg/rt/lang.go (forced by package structure — corefns
// imports rt, so rt cannot import corefns), so this test is the guard against
// the copies drifting from the originals.

import (
	"testing"

	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/rt/corefns"
	"github.com/nooga/let-go/pkg/vm"
)

func coreFnInputs(t *testing.T) []vm.Value {
	t.Helper()
	return []vm.Value{
		vm.NIL,
		vm.EmptyList,
		vm.NewArrayVector([]vm.Value{}),
		vm.NewArrayVector([]vm.Value{vm.Int(1)}),
		vm.NewArrayVector([]vm.Value{vm.Int(1), vm.Int(2), vm.Int(3)}),
		vm.String(""),
		vm.String("abc"),
		vm.NewCons(vm.Int(9), vm.EmptyList),
	}
}

func TestCorefnsArity1MatchBuiltins(t *testing.T) {
	cases := []struct {
		name string
		fn   func(vm.Value) (vm.Value, error)
	}{
		{"first", corefns.First},
		{"count", corefns.Count},
		{"rest", corefns.Rest},
		{"second", corefns.Second},
		{"next", corefns.Next},
	}
	ns := rt.NS(rt.NameCoreNS)
	// Canonical equality: the clojure.core/= builtin itself.
	eqVar := ns.LookupLocal(vm.Symbol("="))
	if eqVar == nil {
		t.Fatal("clojure.core/= not defined")
	}
	equal := func(a, b vm.Value) bool {
		r, err := eqVar.Invoke([]vm.Value{a, b})
		return err == nil && r == vm.Boolean(true)
	}
	for _, c := range cases {
		v := ns.LookupLocal(vm.Symbol(c.name))
		if v == nil {
			t.Fatalf("clojure.core/%s not defined", c.name)
		}
		for _, in := range coreFnInputs(t) {
			want, wantErr := v.Invoke([]vm.Value{in})
			got, gotErr := c.fn(in)
			if (wantErr == nil) != (gotErr == nil) {
				t.Errorf("%s(%s): error mismatch: builtin err=%v, corefns err=%v",
					c.name, in.String(), wantErr, gotErr)
				continue
			}
			if wantErr != nil {
				continue // both errored — good enough; messages may differ
			}
			if !equal(want, got) {
				t.Errorf("%s(%s): builtin=%s, corefns=%s",
					c.name, in.String(), want.String(), got.String())
			}
		}
	}
}
