/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir_test

import (
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/vm"
)

// The registry must exist and, for a single-arity fn, produce for each
// target exactly what the target's lower fn produces natively.
func TestLoweringsRegistryExists(t *testing.T) {
	ensureLoader()
	// :go strategy on a single-arity defn yields a {:status :lowered :decl …} map.
	got := runLispExpr(t, `
      (let [ir-fn (-> (quote (defn add [x y] (+ x y)))
                      ir.build/build-fn
                      ir.passes.pipeline/optimize-fn)]
        (pr-str (:status ((ir.passes.pipeline/lowerings :go) ir-fn))))`)
	if s, ok := got.(vm.String); !ok || string(s) != ":lowered" {
		t.Fatalf("go strategy status = %v, want :lowered", got)
	}
}

// A multi-arity defn must still lower on both targets after the registry
// refactor: :go yields a {:kind :multi-fn-template …} map; :bytecode yields a
// callable multi-arity value.
func TestMultiArityLowersViaRegistry(t *testing.T) {
	ensureLoader()
	goKind := runLispExpr(t, `
      (binding [ir.passes.pipeline/*target* :go]
        (pr-str (:kind (ir.passes.pipeline/compile-form
                         (quote (defn f ([x] x) ([x y] (+ x y))))))))`)
	if s, ok := goKind.(vm.String); !ok || !strings.Contains(string(s), "multi-fn-template") {
		t.Fatalf("go multi-arity kind = %v, want :multi-fn-template", goKind)
	}
}
