/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

// "Are we fast yet?" — wall-clock benchmarks comparing baseline
// (source-compiler bytecode) vs IR-optimized bytecode (Build → ConstFold
// → CSE → DCE → Lower) across a small fixture suite.
//
// Run with:
//   go test -bench=. -benchmem -benchtime=2s ./pkg/ir/passes
//
// To compare baseline vs optimized side-by-side, run twice and use
// `benchstat`:
//   go test -bench='BenchmarkAWFY.*Baseline' -count=10 > /tmp/baseline.txt
//   go test -bench='BenchmarkAWFY.*Optimized' -count=10 > /tmp/optimized.txt
//   benchstat /tmp/baseline.txt /tmp/optimized.txt

package passes

import (
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/compiler"
	"github.com/nooga/let-go/pkg/ir"
	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

// compileAndOptimize returns (baselineChunk, optimizedChunk). The
// baseline is the source compiler's output; the optimized version runs
// through the full IR pipeline.
func compileAndOptimize(b *testing.B, src, fnName string, arity int) (*vm.CodeChunk, *vm.CodeChunk) {
	b.Helper()
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		b.Fatalf("compile: %v", err)
	}
	v := ns.Lookup(vm.Symbol(fnName))
	if v == nil {
		b.Fatalf("symbol %s not found", fnName)
	}
	fn := v.(*vm.Var).Deref().(*vm.Func)
	baseline := fn.Chunk()

	irFn, err := ir.Build(baseline, fnName, arity, false)
	if err != nil {
		b.Fatalf("Build: %v", err)
	}
	ConstFold(irFn)
	CSE(irFn)
	DCE(irFn)
	optimized, err := ir.Lower(irFn)
	if err != nil {
		b.Fatalf("Lower: %v", err)
	}
	return baseline, optimized
}

// runChunk invokes `chunk` with `args` for b.N iterations.
func runChunk(b *testing.B, chunk *vm.CodeChunk, args []vm.Value) {
	b.Helper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		frame := vm.NewFrame(chunk, args)
		_, err := frame.Run()
		vm.ReleaseFrame(frame)
		if err != nil {
			b.Fatalf("run: %v", err)
		}
	}
}

// --- Sumto: tight integer loop. -----------------------------------------

const sumtoSrc = `(defn sumto [n] (loop [i 0 acc 0] (if (< i n) (recur (inc i) (+ acc i)) acc)))`

func BenchmarkAWFY_Sumto10_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, sumtoSrc, "sumto", 1)
	runChunk(b, base, []vm.Value{vm.Int(10)})
}
func BenchmarkAWFY_Sumto10_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, sumtoSrc, "sumto", 1)
	runChunk(b, opt, []vm.Value{vm.Int(10)})
}
func BenchmarkAWFY_Sumto100_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, sumtoSrc, "sumto", 1)
	runChunk(b, base, []vm.Value{vm.Int(100)})
}
func BenchmarkAWFY_Sumto100_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, sumtoSrc, "sumto", 1)
	runChunk(b, opt, []vm.Value{vm.Int(100)})
}
func BenchmarkAWFY_Sumto1000_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, sumtoSrc, "sumto", 1)
	runChunk(b, base, []vm.Value{vm.Int(1000)})
}
func BenchmarkAWFY_Sumto1000_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, sumtoSrc, "sumto", 1)
	runChunk(b, opt, []vm.Value{vm.Int(1000)})
}

// --- Cadd / Cmul / Cmagsq: arithmetic-heavy helpers from mandelbrot. ----

const cmplxSrc = `
(def scale 1000)
(defn cadd [a b c d] (vector (+ a c) (+ b d)))
(defn cmul [a b c d]
  (vector (/ (- (* a c) (* b d)) scale)
          (/ (+ (* a d) (* b c)) scale)))
(defn cmagsq [a b]
  (/ (+ (* a a) (* b b)) scale))
`

func BenchmarkAWFY_Cadd_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, cmplxSrc, "cadd", 4)
	runChunk(b, base, []vm.Value{vm.Int(100), vm.Int(200), vm.Int(50), vm.Int(75)})
}
func BenchmarkAWFY_Cadd_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, cmplxSrc, "cadd", 4)
	runChunk(b, opt, []vm.Value{vm.Int(100), vm.Int(200), vm.Int(50), vm.Int(75)})
}
func BenchmarkAWFY_Cmul_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, cmplxSrc, "cmul", 4)
	runChunk(b, base, []vm.Value{vm.Int(1500), vm.Int(2500), vm.Int(800), vm.Int(1200)})
}
func BenchmarkAWFY_Cmul_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, cmplxSrc, "cmul", 4)
	runChunk(b, opt, []vm.Value{vm.Int(1500), vm.Int(2500), vm.Int(800), vm.Int(1200)})
}
func BenchmarkAWFY_Cmagsq_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, cmplxSrc, "cmagsq", 2)
	runChunk(b, base, []vm.Value{vm.Int(1500), vm.Int(2500)})
}
func BenchmarkAWFY_Cmagsq_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, cmplxSrc, "cmagsq", 2)
	runChunk(b, opt, []vm.Value{vm.Int(1500), vm.Int(2500)})
}

// --- Const-folding wins: an expression that's entirely foldable. -------

// fullyFoldableSrc has a body that ConstFold can collapse to a single
// LOAD_CONST. The "optimized" benchmark should be dramatically faster
// than baseline since the runtime work goes to zero.
const fullyFoldableSrc = `
(defn fully-foldable []
  (+ (* 7 6)
     (- 100 50)
     (quot 1000 10)
     (bit-and 255 240)
     (bit-shift-left 1 4)))
`

func BenchmarkAWFY_FullyFoldable_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, fullyFoldableSrc, "fully-foldable", 0)
	runChunk(b, base, nil)
}
func BenchmarkAWFY_FullyFoldable_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, fullyFoldableSrc, "fully-foldable", 0)
	runChunk(b, opt, nil)
}

// --- Partially foldable: half the work is constant, half is runtime. ---

const partiallyFoldableSrc = `
(defn partial-fold [x]
  (+ (* 7 6)       ; foldable to 42
     (+ x 100)     ; partially foldable: x + 100 stays runtime
     (bit-and 255 (bit-shift-left 1 4)))) ; foldable to 16
`

func BenchmarkAWFY_PartialFold_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, partiallyFoldableSrc, "partial-fold", 1)
	runChunk(b, base, []vm.Value{vm.Int(50)})
}
func BenchmarkAWFY_PartialFold_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, partiallyFoldableSrc, "partial-fold", 1)
	runChunk(b, opt, []vm.Value{vm.Int(50)})
}

// --- Mandelbrot-escape: complex hot loop (uses recur). -----------------

const escapeSrc = `
(def scale 1000)
(defn cadd [a b c d] (vector (+ a c) (+ b d)))
(defn cmul [a b c d]
  (vector (quot (- (* a c) (* b d)) scale)
          (quot (+ (* a d) (* b c)) scale)))
(defn cmagsq [a b]
  (quot (+ (* a a) (* b b)) scale))
(defn mandelbrot-escape-flat [cr ci]
  (loop [zr 0 zi 0 n 0]
    (cond
      (>= n 50) 50
      (> (cmagsq zr zi) (* 4 scale)) n
      :else
        (let [sq (cmul zr zi zr zi)
              new-z (cadd (nth sq 0) (nth sq 1) cr ci)]
          (recur (nth new-z 0) (nth new-z 1) (inc n))))))
`

// We pick a point that escapes after a few iterations so the bench is
// bounded and fast: c=[2000 0] is well outside the mandelbrot set.
func BenchmarkAWFY_MandelbrotEscape_Baseline(b *testing.B) {
	base, _ := compileAndOptimize(b, escapeSrc, "mandelbrot-escape-flat", 2)
	runChunk(b, base, []vm.Value{vm.Int(2000), vm.Int(0)})
}
func BenchmarkAWFY_MandelbrotEscape_Optimized(b *testing.B) {
	_, opt := compileAndOptimize(b, escapeSrc, "mandelbrot-escape-flat", 2)
	runChunk(b, opt, []vm.Value{vm.Int(2000), vm.Int(0)})
}
