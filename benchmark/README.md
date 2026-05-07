# Benchmarks

Compares let-go against other Clojure-compatible runtimes.

## Prerequisites

Required:
- Go (for building let-go)
- [hyperfine](https://github.com/sharkdp/hyperfine) — benchmark runner
- python3 — results formatting

## Optional runtimes

Install any combination — the script auto-detects what's available:

| Runtime | Install |
|---|---|
| [babashka](https://babashka.org/) | `brew install babashka` |
| [joker](https://joker-lang.org/) | `brew install candid82/brew/joker` |
| [go-joker](https://github.com/rcarmo/go-joker) (optimized joker fork) | See [Go-joker setup](#go-joker-setup) |
| [Clojure JVM](https://clojure.org/) | `brew install clojure` |
| [fennel](https://fennel-lang.org/) + cljlib | See [Fennel setup](#fennel-setup) |
| [gloat](https://github.com/gloathub/gloat) (glojure AOT) | See [Gloat setup](#gloat-setup) |

## Running

```bash
bash benchmark/run.sh
```

Results are written to `benchmark/results.md`.

## Fennel setup

Fennel benchmarks use [fennel-cljlib](https://gitlab.com/andreyorst/fennel-cljlib) for
Clojure-compatible lazy seqs, transducers, and persistent data structures.

```bash
brew install fennel luarocks
luarocks install luabc
luarocks install luasocket

# Clone cljlib and fetch dependencies
cd benchmark/fennel
git clone https://gitlab.com/andreyorst/fennel-cljlib.git lib

# Install the deps.fnl package manager
curl -s https://gitlab.com/andreyorst/deps.fnl/-/raw/main/scripts/install | sh

# Fetch cljlib dependencies (may show errors for missing rocks — git deps still fetch fine)
cd lib
~/.local/bin/deps fetch || true
cd ../../..
```

Verify it works:

```bash
benchmark/fennel/run-fennel.sh -e '(local core (require :io.gitlab.andreyorst.cljlib.core)) (print (core.reduce core.+ 0 (core.range 10)))'
# Should print: 45
```

## Go-joker setup

[go-joker](https://github.com/rcarmo/go-joker) is an optimized fork of joker with an IR
bytecode interpreter and WASM/wazero JIT for hot numeric loops. Its binary is also called
`joker`, so it must be built into a separate location and pointed at via the `GOJOKER`
environment variable to avoid conflicting with upstream joker.

```bash
git clone https://github.com/rcarmo/go-joker /tmp/go-joker
cd /tmp/go-joker
go build -o go-joker .
export GOJOKER=/tmp/go-joker/go-joker
```

Set `GOJOKER` (in your shell or in front of `run.sh`) and the script auto-detects it:

```bash
GOJOKER=/tmp/go-joker/go-joker bash benchmark/run.sh
```

## Gloat setup

Gloat compiles Clojure to native Go binaries via glojure AOT. The run script
pre-compiles all benchmarks before timing them (only binary execution is measured).

```bash
git clone https://github.com/gloathub/gloat /tmp/gloat
source /tmp/gloat/.rc   # bootstraps glj, go, bb on first run
cp /tmp/gloat/bin/gloat ~/.local/bin/  # or anywhere on PATH
```

The first `run.sh` invocation with gloat will be slow as it compiles all benchmark
binaries and downloads Go dependencies. Subsequent runs reuse cached binaries
(recompiles only when source files change).
