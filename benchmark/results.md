## Benchmark Results

### Methodology

All benchmarks use [hyperfine](https://github.com/sharkdp/hyperfine) with 3 warmup runs
and 10 timed runs per benchmark. Times shown are mean ± σ wall-clock time. Peak memory is
measured via `/usr/bin/time -l` (median of 3 runs).

Benchmark files are valid Clojure that runs unmodified on let-go, babashka, joker, go-joker,
glojure, and Clojure JVM. Fennel uses equivalent implementations via
[fennel-cljlib](https://gitlab.com/andreyorst/fennel-cljlib) (lazy seqs, transducers,
persistent data structures). Gloat benchmarks are pre-compiled to native binaries via
[gloat](https://github.com/gloathub/gloat) AOT (Clojure→Go); compilation time is not
measured, only binary execution (analogous to how let-go is pre-built with `go build`).

Clojure JVM times include full JVM startup (~350-500ms) which dominates short benchmarks.
Joker is skipped for benchmarks that would exceed reasonable time limits or use unsupported
features (transducers). Binary sizes for gloat are averaged across all benchmark binaries.

**System:** Darwin arm64, Apple M1 Pro

**Runtimes:**

| | let-go | babashka | joker | go-joker | fennel | clojure JVM |
|---|---|---|---|---|---|---|
| **Version** | — | babashka v1.12.217 | joker v1.7.1 | go-joker v1.7.1 | Fennel 1.6.1 on PUC Lua 5.5 | Clojure CLI version 1.12.4.1618 |
| **Platform** | Go bytecode VM | GraalVM native | Go tree-walk interpreter | Go IR + WASM/wazero JIT | Lua VM + cljlib | JVM (HotSpot) |
| **Binary/runtime size** | **10M** | 68M | 26M | 31M | 324K | 304M |

### Startup Time

| Runtime | Time |
|---|---|
| **let-go** | **7.8ms ± 0.7ms** (1.0x) |
| babashka | 21.7ms ± 2.0ms (2.8x) |
| joker | 11.8ms ± 0.4ms (1.5x) |
| go-joker | 12.5ms ± 0.7ms (1.6x) |
| fennel | 58.4ms ± 12.1ms (7.5x) |
| clojure JVM | 0.388s ± 0.012s (49.6x) |

### Peak Memory Usage (RSS)

| Workload | let-go | babashka | joker | go-joker | fennel | clojure JVM |
|---|---|---|---|---|---|---|
| startup (nil) | 13.5MB (1.0x) | 26.8MB (2.0x) | 21.3MB (1.6x) | 22.7MB (1.7x) | **3.2MB** (0.2x) | 92.2MB (6.8x) |
| fib(35) | 14.2MB (1.0x) | 77.2MB (5.4x) | 33.2MB (2.3x) | —MB | **13.1MB** (0.9x) | 112.2MB (7.9x) |
| reduce 1M | **20.0MB** (1.0x) | 59.0MB (3.0x) | 33.1MB (1.7x) | 33.8MB (1.7x) | 851.1MB (42.6x) | 116.5MB (5.8x) |

### Performance

| Benchmark | let-go | babashka | joker | go-joker | fennel | clojure JVM |
|---|---|---|---|---|---|---|
| fib | 2.218s ± 0.163s (1.0x) | 2.020s ± 0.036s (0.9x) | 21.577s ± 0.321s (9.7x) | — | 2.058s ± 0.056s (0.9x) | **0.656s ± 0.025s** (0.3x) |
| loop-recur | 63.1ms ± 3.8ms (1.0x) | 68.4ms ± 1.8ms (1.1x) | 0.774s ± 0.029s (12.3x) | **14.3ms ± 1.3ms** (0.2x) | 0.193s ± 0.004s (3.1x) | 0.546s ± 0.011s (8.6x) |
| map-filter | **9.1ms ± 1.1ms** (1.0x) | 21.5ms ± 0.9ms (2.4x) | 14.2ms ± 0.6ms (1.6x) | 15.3ms ± 1.1ms (1.7x) | 1.169s ± 0.042s (129.0x) | 0.487s ± 0.075s (53.7x) |
| persistent-map | **21.8ms ± 1.0ms** (1.0x) | 26.5ms ± 2.8ms (1.2x) | 53.3ms ± 1.2ms (2.4x) | 52.4ms ± 1.7ms (2.4x) | 3.980s ± 0.132s (182.3x) | 0.563s ± 0.006s (25.8x) |
| reduce | 79.8ms ± 2.5ms (1.0x) | **38.1ms ± 1.0ms** (0.5x) | 2.667s ± 0.049s (33.4x) | 2.790s ± 0.211s (35.0x) | 8.552s ± 0.148s (107.1x) | 0.350s ± 0.009s (4.4x) |
| tak | 2.067s ± 0.023s (1.0x) | 1.932s ± 0.023s (0.9x) | — | — | 10.686s ± 0.206s (5.2x) | **0.562s ± 0.020s** (0.3x) |
| transducers | **6.9ms ± 0.5ms** (1.0x) | 19.2ms ± 0.9ms (2.8x) | — | — | 1.041s ± 0.021s (151.5x) | 0.352s ± 0.005s (51.2x) |

