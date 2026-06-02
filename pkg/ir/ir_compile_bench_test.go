/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/compiler"
	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

// BenchmarkIRCompile measures IR-pipeline compile throughput: building IR
// from source and running the full optimizer (ir.passes.pipeline/optimize-fn
// → dce, constfold, cse, licm, typeinfer, …) over a representative corpus.
//
// This is the yardstick for "is the lowered-to-Go IR usable?". Under the
// bench-ratchet fast gate it runs in two variants:
//   - bytecode  (untagged): the passes execute as replayed Lisp bytecode.
//   - gogen_ir  (-tags gogen_ir): the SAME passes dispatch to their native
//     Go overrides (see pkg/rt/gogen_override.go). The wireup that links the
//     lowered tree into this test binary is the generated, gitignored
//     zz_gogen_ir_wire_test.go (lgbgen --target=go). Without it, -tags
//     gogen_ir matches no files here and the two variants would be identical.
//
// The corpus intentionally exercises several passes: const arithmetic
// (constfold), loop-carried accumulation (licm + typeinfer over block
// params), nested conditionals, repeated sub-expressions (cse), and dead
// lets (dce).
var irCompileCorpus = []string{
	`(defn b-arith [] (+ 1 (* 2 3) (- 10 4)))`,
	`(defn b-poly [x y] (+ (* x x) (* 2 (* x y)) (* y y)))`,
	`(defn b-loop [n] (loop [i 0 acc 0] (if (< i n) (recur (inc i) (+ acc i)) acc)))`,
	`(defn b-cond [x] (cond (< x 0) :neg (= x 0) :zero (< x 10) :small :else :big))`,
	`(defn b-cse [a b] (let [s (+ a b)] (* (+ a b) (+ a b) s)))`,
	`(defn b-dead [x] (let [unused (* x x x) y (+ x 1)] (* y 2)))`,
	`(defn b-nest [x] (if (> x 0) (if (> x 100) :huge :pos) (if (< x -100) :tiny :neg)))`,
}

func BenchmarkIRCompile(b *testing.B) {
	ensureLoader()

	// Drive the same path as `ir-stress ir-compile`: compile each defn under
	// (binding [*ir-compile* true]), which routes the compiler through the
	// IR-optimizing pipeline (build → optimize-fn → bytecode). This is the
	// realistic workload — and the one whose passes the gogen_ir overrides
	// replace. (Calling optimize-fn in isolation skips compiler setup the
	// native passes rely on; the full compile path is what we measure.)
	exprs := make([]string, len(irCompileCorpus))
	for i, src := range irCompileCorpus {
		exprs[i] = fmt.Sprintf(`(binding [*ir-compile* true] (eval (quote %s)))`, src)
	}
	consts := vm.NewConsts()
	coreNS := rt.NS(rt.NameCoreNS)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, expr := range exprs {
			c := compiler.NewCompiler(consts, coreNS)
			c.SetSource("ir-compile-bench")
			if _, _, err := c.CompileMultiple(strings.NewReader(expr)); err != nil {
				b.Fatalf("compile: %v", err)
			}
		}
	}
}
