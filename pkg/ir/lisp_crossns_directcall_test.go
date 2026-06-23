/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir_test

// Phase 1 of cross-ns typed direct calls: the EMIT machinery. A *lowered-registry*
// entry for a fn in ANOTHER lowered namespace (:go-pkg set, :native? false) lowers
// a call to `pkg.Fn(ec, args)` — package-qualified like a native module call, but
// ec-passing like an intra-ns lowered call — and records the callee package for
// import. (Phase 2 wires the corpus-wide registry that produces such entries.)

import (
	"fmt"
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

func TestCrossNsLoweredDirectCallEmit(t *testing.T) {
	ensureLoader()
	runLispExpr(t, `(do (create-ns (quote otherns)) (intern (quote otherns) (quote callee)))`)

	fn := buildLispIR(t, `(defn xcaller [x] (otherns/callee x))`)
	optimizeLispIR(t, fn)
	passVarCounter++
	varName := fmt.Sprintf("*xns-dc-fn-%d*", passVarCounter)
	rt.NS(rt.NameCoreNS).Def(varName, fn)

	// Hand-seed a cross-ns lowered entry (what Phase 2's corpus-wide registry
	// will build). Key = [internal-ns-sym("otherns") "callee" 1] = ['otherns ...].
	v := runLispExpr(t, fmt.Sprintf(`
	  (let [reg {(ir.lower-go/registry-key (quote otherns) "callee" 1)
	             {:go-name "Callee" :arity 1 :needs-error? false
	              :param-specs ["vm.Value"] :result-spec "vm.Value"
	              :native? false
	              :go-pkg "github.com/nooga/let-go/pkg/rt/core_go_lowered/otherns"}}]
	    (binding [ir.lower-go/*lowered-registry* reg
	              ir.lower-go/*native-imports-used* (atom #{})]
	      (ir.lower-go/lower %s :strict)))`, varName))
	m, ok := v.(*vm.PersistentMap)
	if !ok {
		t.Fatalf("expected lower to return a map, got %T", v)
	}
	rendered := bindAndRenderGoDecl(t, m)

	if !strings.Contains(rendered, "otherns.Callee(ec,") {
		t.Fatalf("expected cross-ns lowered direct call otherns.Callee(ec, ...):\n--- go ---\n%s", rendered)
	}
	if strings.Contains(rendered, "InvokeValue") || strings.Contains(rendered, "CachedVarFn") {
		t.Fatalf("expected a direct call, but found a trampoline:\n--- go ---\n%s", rendered)
	}
}
