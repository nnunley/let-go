/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package passes

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/compiler"
	"github.com/nooga/let-go/pkg/ir"
	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

// TestConstFold_AgainstClojure: differential test comparing ConstFold's
// compile-time evaluation of pure expressions against three references:
//
//  1. let-go's runtime evaluation of the same source (baseline).
//  2. JVM Clojure's evaluation via `clj -M -e`.
//  3. ConstFold's compile-time fold (the IR pipeline).
//
// All three must agree on every case. Disagreement means ConstFold has
// a semantic divergence from let-go runtime OR from JVM Clojure (which
// is the spec).
//
// Skipped if `clj` is not on PATH.
func TestConstFold_AgainstClojure(t *testing.T) {
	if _, err := exec.LookPath("clj"); err != nil {
		t.Skip("clj not on PATH; skipping JVM differential")
		return
	}

	// Each expression must be:
	//  - pure (no side effects)
	//  - composed of operations let-go supports
	//  - composed of constants (so ConstFold can reduce it)
	//
	// The corpus covers the full numeric tower (Int+Int, Int+Float,
	// Float+Float, BigInt, Ratio, mixed) and the function-call folding
	// path (/ quot rem mod bit-* shifts).
	corpus := []string{
		// Int arithmetic.
		"(+ 1 2)",
		"(+ 1 2 3 4)", // n-ary via repeated +
		"(- 10 3 2)",
		"(* 2 3 4)",
		"(* -2 3)",
		"(inc 41)",
		"(dec 100)",
		// Mixed Int+Float (Float-result expected).
		"(+ 1 2.5)",
		"(+ 2.5 1)",
		"(* 3 0.5)",
		"(- 1.0 1)",
		// Float+Float.
		"(+ 1.5 2.5)",
		"(* 0.1 0.2)",
		// Division semantics — Clojure /  makes ratios for Int/Int.
		"(/ 10 3)",
		"(/ 10 5)",
		"(/ 1.0 4)",
		// quot / rem / mod with various signs.
		"(quot 10 3)",
		"(quot -10 3)",
		"(quot 10 -3)",
		"(rem 10 3)",
		"(rem -10 3)",
		"(mod 10 3)",
		"(mod -7 3)",
		"(mod 7 -3)",
		// Comparison.
		"(< 1 2)",
		"(<= 2 2)",
		"(> 3 2)",
		"(>= 3 3)",
		"(= 5 5)",
		"(= 5 5.0)", // numeric equality across types
		// Bitops.
		"(bit-and 255 240)",
		"(bit-or 0 7)",
		"(bit-xor 5 3)",
		"(bit-not 0)",
		"(bit-shift-left 1 8)",
		"(bit-shift-right 1024 4)",
		"(unsigned-bit-shift-right -1 60)",
		// Nested expressions exercising primitive + call folding together.
		"(* (quot 10 3) (rem 10 3))",
		"(+ (bit-and 15 7) (bit-shift-left 1 2))",
		// Auto-promoting variants — these should produce BigInt at the
		// boundary. let-go and JVM Clojure both auto-promote with the
		// quoted operators.
		"(*' 9223372036854775807 2)",
		"(+' 9223372036854775807 1)",
		"(-' -9223372036854775808 1)",
		// BigInt literals (the `N` suffix tells the reader to use BigInt).
		"(* 9223372036854775808N 2)",
		"(+ 9223372036854775808N 0)",
		// Ratios: arithmetic that produces ratios via /.
		"(+ 1/2 1/3)",
		"(* 2/3 3/4)",
		"(- 1/2 1/4)",
		// Comparison across types (Int vs Ratio).
		"(= 1 1/1)",
		"(< 1/2 2/3)",
	}

	// Cases where evaluation throws on both sides (overflow with the
	// non-promoting operators). ConstFold should leave these as runtime
	// ops since folding would silently produce a wrong answer.
	errorCases := []string{
		"(* 9223372036854775807 2)",  // long overflow
		"(+ 9223372036854775807 1)",  // long overflow
		"(- -9223372036854775808 1)", // long overflow
		"(/ 1 0)",                    // divide by zero
		"(quot 1 0)",
		"(mod 1 0)",
	}

	jvmResults, err := evalClojureBatch(corpus)
	if err != nil {
		t.Fatalf("Clojure batch eval failed: %v", err)
	}
	if len(jvmResults) != len(corpus) {
		t.Fatalf("Clojure returned %d results for %d expressions", len(jvmResults), len(corpus))
	}

	for i, expr := range corpus {
		t.Run(expr, func(t *testing.T) {
			jvm := jvmResults[i]

			// Baseline: compile and run the expression in let-go's runtime.
			baseline := evalLetGoRuntime(t, expr)
			// Optimized: pipe through Build → ConstFold → DCE → Lower → Run.
			optimized := evalLetGoIROptimized(t, expr)

			if baseline != optimized {
				t.Errorf("baseline %q != optimized %q (JVM: %q)", baseline, optimized, jvm)
			}
			// JVM is the reference. Diverging from JVM is a let-go bug
			// (could be either runtime or fold). Report but don't fail
			// for known-divergent cases — initially we just want to see
			// what's different.
			if baseline != jvm {
				t.Errorf("let-go runtime %q diverges from JVM Clojure %q", baseline, jvm)
			}
		})
	}

	// Error cases: each expression should error in BOTH baseline and
	// IR-optimized let-go (matching what JVM Clojure does). The
	// important property is that ConstFold doesn't silently swallow an
	// error by folding to a wrong value.
	for _, expr := range errorCases {
		t.Run("error:"+expr, func(t *testing.T) {
			baseErr := errorFromLetGoRuntime(t, expr)
			optErr := errorFromLetGoIROptimized(t, expr)
			if baseErr == "" {
				t.Errorf("baseline %q expected error, got success", expr)
			}
			if optErr == "" {
				t.Errorf("optimized %q expected error, got success — ConstFold may have silently folded to wrong value", expr)
			}
		})
	}
}

// errorFromLetGoRuntime: like evalLetGoRuntime but returns the error
// string (empty if no error). Used for negative-test cases.
func errorFromLetGoRuntime(t *testing.T, expr string) string {
	t.Helper()
	src := "(defn __terr [] " + expr + ")"
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		return err.Error()
	}
	v := ns.Lookup(vm.Symbol("__terr"))
	if v == nil {
		return "symbol not found"
	}
	fn := v.(*vm.Var).Deref().(*vm.Func)
	_, err := fn.Invoke(nil)
	if err == nil {
		return ""
	}
	return err.Error()
}

// errorFromLetGoIROptimized: same shape, IR pipeline.
func errorFromLetGoIROptimized(t *testing.T, expr string) string {
	t.Helper()
	src := "(defn __oerr [] " + expr + ")"
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		return err.Error()
	}
	v := ns.Lookup(vm.Symbol("__oerr"))
	if v == nil {
		return "symbol not found"
	}
	fn := v.(*vm.Var).Deref().(*vm.Func)

	irFn, err := ir.Build(fn.Chunk(), "__oerr", 0, false)
	if err != nil {
		return err.Error()
	}
	ConstFold(irFn)
	DCE(irFn)
	chunk, err := ir.Lower(irFn)
	if err != nil {
		return err.Error()
	}
	frame := vm.NewFrame(chunk, nil)
	_, err = frame.Run()
	vm.ReleaseFrame(frame)
	if err == nil {
		return ""
	}
	return err.Error()
}

// evalLetGoRuntime compiles `(defn __test [] <expr>)` in let-go and
// invokes __test, returning the result's String().
func evalLetGoRuntime(t *testing.T, expr string) string {
	t.Helper()
	src := "(defn __test [] " + expr + ")"
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		t.Fatalf("compile %q: %v", expr, err)
	}
	v := ns.Lookup(vm.Symbol("__test"))
	if v == nil {
		t.Fatalf("__test symbol not found after compiling %q", expr)
	}
	fn := v.(*vm.Var).Deref().(*vm.Func)
	out, err := fn.Invoke(nil)
	if err != nil {
		t.Fatalf("run %q: %v", expr, err)
	}
	return out.String()
}

// evalLetGoIROptimized goes through Build → ConstFold → DCE → Lower and
// runs the optimized chunk. The expression should be fully foldable
// (all-constant), so the optimized chunk's body should be roughly a
// single LOAD_CONST + RETURN.
func evalLetGoIROptimized(t *testing.T, expr string) string {
	t.Helper()
	src := "(defn __opt [] " + expr + ")"
	consts := vm.NewConsts()
	ns := rt.NS(rt.NameCoreNS)
	ctx := compiler.NewCompiler(consts, ns)
	if _, _, err := ctx.CompileMultiple(strings.NewReader(src)); err != nil {
		t.Fatalf("compile %q: %v", expr, err)
	}
	v := ns.Lookup(vm.Symbol("__opt"))
	if v == nil {
		t.Fatalf("__opt symbol not found after compiling %q", expr)
	}
	fn := v.(*vm.Var).Deref().(*vm.Func)

	irFn, err := ir.Build(fn.Chunk(), "__opt", 0, false)
	if err != nil {
		t.Fatalf("Build %q: %v", expr, err)
	}
	ConstFold(irFn)
	DCE(irFn)
	chunk, err := ir.Lower(irFn)
	if err != nil {
		t.Fatalf("Lower %q: %v", expr, err)
	}
	frame := vm.NewFrame(chunk, nil)
	out, err := frame.Run()
	vm.ReleaseFrame(frame)
	if err != nil {
		t.Fatalf("run optimized %q: %v", expr, err)
	}
	return out.String()
}

// evalClojureBatch evaluates a slice of expressions in a single Clojure
// subprocess, returning one result per expression in pr-str form.
//
// Uses `clj -M -e` with a doseq that prints each result on its own line.
// The slice order is preserved.
func evalClojureBatch(exprs []string) ([]string, error) {
	// Build a Clojure expression that prints each input's result.
	// We use pr-str so the printed form is round-trippable (numbers
	// without surrounding quotes, ratios as "10/3", floats as "3.0").
	//
	// Using a vector of forms with eval keeps Clojure's reader happy
	// even for expressions containing top-level Clojure values.
	var sb strings.Builder
	sb.WriteString("(doseq [e (quote [")
	for _, e := range exprs {
		sb.WriteString(e)
		sb.WriteByte(' ')
	}
	sb.WriteString("])] (println (pr-str (eval e))))")

	cmd := exec.Command("clj", "-M", "-e", sb.String())
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("clj exec: %w (stderr: %s)", err, exitStderr(err))
	}
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// Skip "WARNING: Implicit use of clojure.main..." informational lines.
		if strings.HasPrefix(line, "WARNING:") {
			continue
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func exitStderr(err error) string {
	if ee, ok := err.(*exec.ExitError); ok {
		return string(ee.Stderr)
	}
	return ""
}
