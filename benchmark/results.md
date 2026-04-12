## Benchmark Results

### Methodology

All benchmarks use [hyperfine](https://github.com/sharkdp/hyperfine) with 3 warmup runs
and 10 timed runs per benchmark. Times shown are mean ± σ wall-clock time. Peak memory is
measured via `/usr/bin/time -l` (median of 3 runs).

Benchmark files are valid Clojure that runs unmodified on let-go, babashka, joker, glojure,
and Clojure JVM. Fennel uses equivalent implementations via
[fennel-cljlib](https://gitlab.com/andreyorst/fennel-cljlib) (lazy seqs, transducers,
persistent data structures). Gloat benchmarks are pre-compiled to native binaries via
[gloat](https://github.com/gloathub/gloat) AOT (Clojure→Go); compilation time is not
measured, only binary execution (analogous to how let-go is pre-built with `go build`).

Clojure JVM times include full JVM startup (~350-500ms) which dominates short benchmarks.
Joker is skipped for benchmarks that would exceed reasonable time limits or use unsupported
features (transducers). Binary sizes for gloat are averaged across all benchmark binaries.

**System:** Darwin arm64, Apple M1 Pro

**Runtimes:**

| | let-go | babashka | joker | gloat | fennel | clojure JVM |
|---|---|---|---|---|---|---|
| **Version** | — | babashka v1.12.217 | joker v1.7.1 | gloat version 0.1.23 | Fennel 1.6.1 on PUC Lua 5.5 | Clojure CLI version 1.12.4.1618 |
| **Platform** | Go bytecode VM | GraalVM native | Go tree-walk interpreter | Go AOT (Clojure→Go) | Lua VM + cljlib | JVM (HotSpot) |
| **Binary/runtime size** | **9.6M** | 68M | 26M | 26M | 324K | 304M |

### Startup Time

| Runtime | Time |
|---|---|
| **let-go** | **7.6ms ± 1.2ms** (1.0x) |
| babashka | 21.9ms ± 2.5ms (2.9x) |
| joker | 12.0ms ± 1.4ms (1.6x) |
| gloat | 16.4ms ± 2.7ms (2.2x) |
| fennel | 59.7ms ± 37.3ms (7.9x) |
| clojure JVM | 0.378s ± 0.041s (49.9x) |

### Peak Memory Usage (RSS)

| Workload | let-go | babashka | joker | gloat | fennel | clojure JVM |
|---|---|---|---|---|---|---|
| startup (nil) | 12.8MB (1.0x) | 26.7MB (2.1x) | 21.2MB (1.7x) | 22.9MB (1.8x) | **3.1MB** (0.2x) | 92.1MB (7.2x) |
| fib(35) | 13.7MB (1.0x) | 77.1MB (5.6x) | 33.3MB (2.4x) | 30.9MB (2.3x) | **12.6MB** (0.9x) | 111.6MB (8.1x) |
| reduce 1M | **20.0MB** (1.0x) | 58.9MB (2.9x) | 33.2MB (1.7x) | 25.8MB (1.3x) | 887.3MB (44.4x) | 116.5MB (5.8x) |

### Performance

| Benchmark | let-go | babashka | joker | gloat | fennel | clojure JVM |
|---|---|---|---|---|---|---|
| fib | 2.160s ± 0.207s (1.0x) | 2.118s ± 0.234s (1.0x) | 22.490s ± 1.008s (10.4x) | 27.445s ± 0.521s (12.7x) | 1.954s ± 0.104s (0.9x) | **0.529s ± 0.007s** (0.2x) |
| loop-recur | **56.6ms ± 0.6ms** (1.0x) | 64.2ms ± 2.1ms (1.1x) | 0.695s ± 0.013s (12.3x) | 1.015s ± 0.009s (18.0x) | 0.170s ± 0.004s (3.0x) | 0.442s ± 0.005s (7.8x) |
| map-filter | **6.4ms ± 0.2ms** (1.0x) | 18.9ms ± 1.5ms (2.9x) | 12.7ms ± 0.7ms (2.0x) | 61.4ms ± 1.3ms (9.5x) | 1.034s ± 0.027s (160.4x) | 0.340s ± 0.005s (52.7x) |
| persistent-map | **20.6ms ± 4.6ms** (1.0x) | 22.4ms ± 3.2ms (1.1x) | 48.2ms ± 1.0ms (2.3x) | 32.7ms ± 1.5ms (1.6x) | 3.577s ± 0.024s (173.9x) | 0.476s ± 0.018s (23.1x) |
| reduce | 71.1ms ± 1.0ms (1.0x) | **38.2ms ± 3.3ms** (0.5x) | 2.456s ± 0.033s (34.6x) | 0.377s ± 0.016s (5.3x) | 7.770s ± 0.152s (109.4x) | 0.338s ± 0.003s (4.8x) |
| tak | 2.033s ± 0.016s (1.0x) | 1.912s ± 0.018s (0.9x) | — | 21.762s ± 0.170s (10.7x) | 10.581s ± 0.092s (5.2x) | **0.552s ± 0.008s** (0.3x) |
| transducers | **6.4ms ± 0.3ms** (1.0x) | 19.9ms ± 3.0ms (3.1x) | — | 15.0ms ± 0.6ms (2.3x) | 1.022s ± 0.022s (160.0x) | 0.340s ± 0.004s (53.2x) |

