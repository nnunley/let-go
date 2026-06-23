---
status: active
last-verified: 2026-06-23
human-verified:
---

# gogen_ir native-lowering perf gate (compute-bound microbench)

`compute-bound.lg` is a clean, startup-isolated, output-free microbenchmark for
answering one question: **does `-tags gogen_ir` native lowering actually beat the
bytecode VM on compute?**

It deliberately lives outside `test/` so the language test runner's
`filepath.Walk` does not execute it on every `go test`.

## Harness

```sh
# Build both engines from the same sources:
go build -o /tmp/lg-bc  .              # bytecode VM
go build -tags gogen_ir -o /tmp/lg-aot .   # native-lowered core

# Warm runs (first run includes one-time costs; take the steady-state):
for i in 1 2 3; do /usr/bin/time -p /tmp/lg-bc  docs/perf/microbench/compute-bound.lg; done
for i in 1 2 3; do /usr/bin/time -p /tmp/lg-aot docs/perf/microbench/compute-bound.lg; done

# Startup baseline (subtract from the above to isolate compute):
printf '(println 1)\n' > /tmp/noop.lg && /usr/bin/time -p /tmp/lg-bc /tmp/noop.lg
```

## Baseline (2026-06-23, Apple M3)

| Workload | bytecode | native (gogen_ir) |
|---|---|---|
| startup (noop) | 0.05s | — |
| `compute-bound.lg` | 1.52 / 1.51 / 1.53s | 1.54 / 1.51 / 1.52s |
| jank suite (`BenchmarkClojureTestSuite`) | 379 ms | 394 ms (3× compile) |

**Finding (the "red flag"):** native lowering delivers **no compute speedup**
(marginally slower), at ~3× compile cost.

## Root cause

The lowered Go is dynamic-dispatch-bound, not a fallback issue (`core_go_lowered/`
is populated; ~496 fns are natively overridden). It emits ~1083
`ec.Invoke(rt.CachedVarFn(&var, ns, name), []vm.Value{args})` dynamic-dispatch
sites vs only ~107 direct calls, with all values boxed as `vm.Value` and ~1120
`goto`s. Every inter-fn call (e.g. `clojure.core/map`, `reduce`, `+`) pays a var
lookup + `IFn.Invoke` dispatch + per-call slice allocation + boxing — the same
per-op work the bytecode VM does. Native only removes the bytecode fetch-decode
loop, which is not the bottleneck.

## Fix levers (Milestone A)

1. **Direct-call lowering** (dominant): rewrite `ec.Invoke(CachedVarFn(…), …)`
   → direct `pkg.Fn(ec, args)` for known lowered/native callees, especially
   cross-ns calls into `clojure.core` (the #168 registry exists but reaches only
   ~107/1190 sites).
2. **Unbox** typed int/float locals.
3. **Structural control flow** (`structurize`) to remove goto-CFG and unlock Go
   inlining.

A win shows up as `lg-aot` time on `compute-bound.lg` dropping below `lg-bc`.
