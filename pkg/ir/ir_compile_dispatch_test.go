//go:build gogen_ir

/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir_test

import (
	"testing"

	"github.com/nooga/let-go/pkg/rt"
)

// TestIRBenchDispatchesNativeUnderTag guards that BenchmarkIRCompile actually
// measures the NATIVE passes under -tags gogen_ir, not bytecode. It is the
// pkg/ir analogue of `make check-gogen-ir`: after the loader wires the IR
// namespaces, a known pass Var must resolve to a NativeFn override (installed
// by the generated wireup zz_gogen_ir_wire_test.go + the resolver's
// ApplyGoOverrides hook). If the wireup regresses — wrong package, missing
// blank import, dropped hook — this fails, instead of the benchmark silently
// degrading to the bytecode-vs-bytecode tautology.
func TestIRBenchDispatchesNativeUnderTag(t *testing.T) {
	ensureLoader()

	v := rt.LookupVar("ir.passes.dce", "dce")
	if v == nil {
		t.Fatal("ir.passes.dce/dce var not found — loader did not wire ir.passes.dce")
	}
	got := v.Deref()
	if name := got.Type().Name(); name != "let-go.lang.NativeFn" {
		t.Fatalf("ir.passes.dce/dce should be a NativeFn override under -tags gogen_ir; "+
			"got %s — the lowered tree is not dispatching (wireup or override hook regressed)", name)
	}
}
