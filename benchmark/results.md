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

| | let-go | babashka | joker | go-joker | gloat | fennel | clojure JVM |
|---|---|---|---|---|---|---|---|
| **Version** | — | babashka v1.12.217 | joker v1.7.1 | go-joker v42.8.2 | gloat version 0.1.36 | Fennel 1.6.1 on PUC Lua 5.5 | Clojure CLI version 1.12.4.1618 |
| **Platform** | Go bytecode VM | GraalVM native | Go tree-walk interpreter | Go IR + WASM/wazero JIT | Go AOT (Clojure→Go) | Lua VM + cljlib | JVM (HotSpot) |
| **Binary/runtime size** | **10M** | 68M | 26M | 32M | 26M | 324K | 304M |

### Startup Time

let-go, babashka, and fennel are measured with 10 warmup + 100 timed runs and a
5% trimmed mean (top/bottom 5% dropped) to reject scheduling outliers. Other
runtimes use the default 3 warmup + 10 runs.

| Runtime | Time |
|---|---|
| **let-go** | **6.7ms ± 0.3ms** (1.00x) |
| joker | 12.3ms ± 1.6ms (1.85x) |
| go-joker | 12.5ms ± 0.6ms (1.88x) |
| gloat | 15.7ms ± 0.9ms (2.36x) |
| babashka | 18.0ms ± 0.9ms (2.70x) |
| fennel | 36.0ms ± 1.7ms (5.41x) |
| clojure JVM | 0.363s ± 0.007s (54.50x) |

### Peak Memory Usage (RSS)

| Workload | let-go | babashka | joker | go-joker | gloat | fennel | clojure JVM |
|---|---|---|---|---|---|---|---|
| startup (nil) | 13.5MB (1.0x) | 26.8MB (2.0x) | 21.6MB (1.6x) | 23.3MB (1.7x) | 22.9MB (1.7x) | **3.2MB** (0.2x) | 91.7MB (6.8x) |
| fib(35) | 14.4MB (1.0x) | 77.2MB (5.4x) | 33.1MB (2.3x) | 23.9MB (1.7x) | 32.8MB (2.3x) | **13.0MB** (0.9x) | 112.5MB (7.8x) |
| reduce 1M | **20.1MB** (1.0x) | 59.0MB (2.9x) | 33.2MB (1.7x) | 23.6MB (1.2x) | 25.7MB (1.3x) | 1137.0MB (56.6x) | 116.6MB (5.8x) |

### Performance

| Benchmark | let-go | babashka | joker | go-joker | gloat | fennel | clojure JVM |
|---|---|---|---|---|---|---|---|
| fib | 2.075s ± 0.012s (1.0x) | 1.997s ± 0.028s (1.0x) | 21.002s ± 0.076s (10.1x) | 1.472s ± 0.006s (0.7x) | 27.925s ± 0.508s (13.5x) | 2.004s ± 0.026s (1.0x) | **0.615s ± 0.013s** (0.3x) |
| loop-recur | 60.3ms ± 0.8ms (1.0x) | 68.0ms ± 1.2ms (1.1x) | 0.758s ± 0.026s (12.6x) | **14.3ms ± 0.5ms** (0.2x) | 1.047s ± 0.011s (17.4x) | 0.188s ± 0.001s (3.1x) | 0.520s ± 0.009s (8.6x) |
| map-filter | **7.9ms ± 0.5ms** (1.0x) | 21.5ms ± 0.8ms (2.7x) | 14.1ms ± 0.8ms (1.8x) | 13.3ms ± 1.3ms (1.7x) | 64.2ms ± 1.1ms (8.1x) | 1.105s ± 0.011s (140.2x) | 0.398s ± 0.004s (50.5x) |
| persistent-map | **20.8ms ± 0.6ms** (1.0x) | 23.7ms ± 0.8ms (1.1x) | 51.5ms ± 1.4ms (2.5x) | 21.3ms ± 0.8ms (1.0x) | 34.6ms ± 1.0ms (1.7x) | 3.773s ± 0.048s (181.6x) | 0.646s ± 0.188s (31.1x) |
| reduce | 80.7ms ± 2.3ms (1.0x) | 43.4ms ± 7.8ms (0.5x) | 2.620s ± 0.031s (32.5x) | **13.0ms ± 0.9ms** (0.2x) | 0.382s ± 0.003s (4.7x) | 8.875s ± 0.543s (110.0x) | 0.398s ± 0.005s (4.9x) |
| tak | 2.111s ± 0.004s (1.0x) | 1.973s ± 0.007s (0.9x) | — | 1.760s ± 0.130s (0.8x) | 22.604s ± 0.487s (10.7x) | 10.436s ± 0.108s (4.9x) | **0.525s ± 0.007s** (0.2x) |
| transducers | 45.9ms ± 4.2ms (1.00x) | 22.7ms ± 1.5ms (0.49x) | — | **15.4ms ± 0.8ms** (0.34x) | 0.198s ± 0.011s (4.32x) | 1.689s ± 0.073s (36.80x) | 0.379s ± 0.066s (8.26x) |

