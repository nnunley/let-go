/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir_test

// Tests for cross-namespace direct calls into native modules (lg-dq58).
//
// main already has the typed direct-call emission (emit-typed-direct-call)
// and cross-ns resolution (resolve-call-entry handles :native? entries +
// records the import). The missing piece is seeding *lowered-registry* with
// the native modules exposed via (rt/native-modules), so a call to a native
// clojure.core fn (e.g. seq -> corefns.Seq) resolves to a direct Go call
// instead of an rt.CachedVarFn / rt.InvokeValue trampoline.

import (
	"fmt"
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/rt"
	// Link corefns so its init() registers clojure.core native modules
	// (e.g. seq -> corefns.Seq) into (rt/native-modules). The real lowering
	// path (lgbgen) links this transitively; the unit test must do so too.
	_ "github.com/nooga/let-go/pkg/rt/corefns"
	"github.com/nooga/let-go/pkg/vm"
)

// lowerWithNativeRegistry lowers the IR fn bound to varName with
// *lowered-registry* seeded from (rt/native-modules), mirroring what
// pipeline.lg does for a real package lowering.
func lowerWithNativeRegistry(t *testing.T, varName string) *vm.PersistentMap {
	t.Helper()
	v := runLispExpr(t, fmt.Sprintf(
		`(binding [ir.lower-go/*lowered-registry* (ir.lower-go/native-registry (rt/native-modules))]
		   (ir.lower-go/lower %s :strict))`, varName))
	m, ok := v.(*vm.PersistentMap)
	if !ok {
		t.Fatalf("expected lower to return a map, got %T", v)
	}
	return m
}

func TestNativeModuleSeedsCrossNsDirectCall(t *testing.T) {
	ensureLoader()
	fn := buildLispIR(t, `(defn use-seq [x] (seq x))`)
	optimizeLispIR(t, fn)
	passVarCounter++
	varName := fmt.Sprintf("*native-dc-fn-%d*", passVarCounter)
	rt.NS(rt.NameCoreNS).Def(varName, fn)

	result := lowerWithNativeRegistry(t, varName)
	rendered := bindAndRenderGoDecl(t, result)

	if !strings.Contains(rendered, "corefns.Seq(") {
		t.Fatalf("expected cross-ns native direct call corefns.Seq(...):\n--- go ---\n%s", rendered)
	}
	if strings.Contains(rendered, "InvokeValue") || strings.Contains(rendered, "CachedVarFn") {
		t.Fatalf("expected a direct call, but found a trampoline:\n--- go ---\n%s", rendered)
	}
}
