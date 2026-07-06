//go:build yamlstar_native

// yamlstar-lg-native — AOT-native YamlStar driver running Ingy's bench.clj
// method (cold call, one load per input, adaptive bulk over "foo: 42"),
// directly comparable to the interpreted baseline. The yamlstar.* fns dispatch
// to the AOT-lowered native Go (blank imports register Go overrides via init();
// require'ing the namespaces + the resolver's ApplyGoOverrides/ApplyGoVarInits
// installs them) instead of the bytecode VM.
//
// Build (regenerate the lowered tree first — Task C1):
//
//	go build -tags "yamlstar_native gogen_ir" -o /tmp/yamlstar-lg-native ./examples/aot/yamlstar/native
//	YAMLSTAR_SRC=$HOME/development/yamlstar/core/src/yamlstar /tmp/yamlstar-lg-native
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nooga/let-go/pkg/api"

	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/composer"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/constructor"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/core"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/parser"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/parser/core"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/parser/grammar"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/parser/parser"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/parser/prelude"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/parser/receiver"
	_ "github.com/nooga/let-go/examples/aot/yamlstar/go/yamlstar/resolver"
)

// Ingy's bench.clj method, ported: cold call, one load per input, adaptive bulk.
const driver = `
(defn now-ns [] (System/nanoTime))
(defn fmt-time [ns]
  (let [ms (/ (double ns) 1000000.0)]
    (cond (>= ms 1000) (clojure.core/format "%8.2f s " (/ ms 1000))
          (< ms 1)     (clojure.core/format "%7.3f ms" ms)
          :else        (clojure.core/format "%7.1f ms" ms))))
(def inputs
  [["scalar"  "hello"]
   ["mapping" "foo: 42"]
   ["nested"  "root:\n  child1:\n    key: value\n  child2:\n  - item1\n  - item2\n  - item3"]
   ["types"   "string: hello\ninteger: 42\nfloat: 3.14\nbool: true\nnull_val: null"]])
(let [t0 (now-ns)]
  (yamlstar.core/load "warmup: true")
  (println (str "cold:    " (fmt-time (- (now-ns) t0)) "  (first call, includes ns init)")))
(println)
(println (clojure.core/format "%-12s %12s" "input" "time"))
(println (apply str (repeat 26 "-")))
(loop [remaining inputs]
  (when (seq remaining)
    (let [[label input] (first remaining)
          t0 (now-ns) _ (yamlstar.core/load input)
          elapsed (- (now-ns) t0)]
      (println (clojure.core/format "%-12s %s" label (fmt-time elapsed)))
      (recur (rest remaining)))))
(let [t0 (now-ns) _ (yamlstar.core/load "foo: 42")
      probe (/ (double (- (now-ns) t0)) 1000000.0)
      reps (cond (< probe 10) 100 (< probe 500) 10 :else 3)
      t0 (now-ns)]
  (dotimes [_ reps] (yamlstar.core/load "foo: 42"))
  (let [total (- (now-ns) t0)]
    (println)
    (println (clojure.core/format "%d x 'foo: 42':  %s total,  %s/call"
               reps (fmt-time total) (fmt-time (long (/ total reps)))))))
`

func main() {
	lg, err := api.NewLetGo("yamlstar-lg-native")
	if err != nil {
		die(err)
	}
	ys := os.Getenv("YAMLSTAR_SRC")
	if ys == "" {
		home, _ := os.UserHomeDir()
		ys = filepath.Join(home, "development", "yamlstar", "core", "src", "yamlstar")
	}
	if _, err := os.Stat(ys); err != nil {
		fmt.Fprintf(os.Stderr, "yamlstar-lg-native: sources not found at %q; set YAMLSTAR_SRC\n", ys)
		os.Exit(2)
	}
	lg.SetLoadPath([]string{ys, filepath.Dir(ys), "pkg/rt/gogen"})
	// Require first (the resolver drains ApplyGoOverrides + ApplyGoVarInits,
	// installing the native grammar), then run the ingy driver.
	if _, err := lg.Run("(require 'yamlstar.core 'yamlstar.parser 'yamlstar.composer 'yamlstar.resolver 'yamlstar.constructor)"); err != nil {
		die(err)
	}
	if _, err := lg.Run("(do " + driver + ")"); err != nil {
		die(err)
	}
}

func die(err error) {
	fmt.Fprintln(os.Stderr, "yamlstar-lg-native:", err)
	os.Exit(1)
}
