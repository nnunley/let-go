###
# Auto install `go` into ./.cache/local/ if not available.
# Also if `make ... GO-VERSION=1.x.y` is used.

ifneq (,$(or $(if $(shell which go),,1),$(GO-VERSION)))
R := https://github.com/makeplus/makes
M := .cache/makes
$(shell [ -d '$M' ] || git clone -q $R '$M')
include $M/init.mk
# override default Go version with: `make ... GO-VERSION=1.x.y`
GO-VERSION ?= 1.26.3
include $M/go.mk
include $M/shell.mk
endif

# Prefer - to _ for make var names (won't conflict with env vars):
LG := lg
GOLANGCI-LINT := github.com/golangci/golangci-lint/cmd/golangci-lint
REPORT-SCRIPT := scripts/clojure_compat_report.sh

# Resource caps for test invocations. GOMEMLIMIT bounds the Go heap
# (soft cap — runtime aggressively GCs to stay under). GO-TEST-TIMEOUT
# matches CI's per-package timeout so runaway tests die locally too.
# Override on the command line: make test GOMEMLIMIT=4GiB.
GOMEMLIMIT ?= 2GiB
GO-TEST-TIMEOUT ?= 60s

# Standard flags + env for `go test`. Use as: $(GO-TEST-ENV) go test $(GO-TEST-FLAGS) ./...
GO-TEST-ENV := GOMEMLIMIT=$(GOMEMLIMIT)
GO-TEST-FLAGS := -timeout $(GO-TEST-TIMEOUT)


# Start repl by default:
default:: run

run: $(LG)
	./$<

build: $(LG)

# Bundle target. The runtime loads compiled bytecode for the core
# namespaces from this file, NOT from the .lg sources. Anyone editing
# a .lg under pkg/rt/core/ must regenerate the bundle or runtime
# behavior silently diverges from source. This prereq rule makes the
# regeneration automatic — `make test`, `make build`, etc. now keep
# the bundle in lockstep with the .lg sources.
CORE-LG-FILES := $(shell find pkg/rt/core -name '*.lg' -type f 2>/dev/null)
LGBGEN-SOURCES := $(shell find cmd/lgbgen -name '*.go' -type f 2>/dev/null)
pkg/rt/core_compiled.lgb: $(CORE-LG-FILES) $(LGBGEN-SOURCES)
	go run -tags bootstrap ./cmd/lgbgen

# Loud check used by CI / pre-commit to detect a bundle that was
# committed in a stale state (e.g. someone ran a manual `go test`
# without going through make and committed the half-stale state).
# Exits non-zero with a clear remediation when any .lg under
# pkg/rt/core is newer than the committed bundle.
check-bundle-fresh:
	@stale=$$(find pkg/rt/core -name '*.lg' -newer pkg/rt/core_compiled.lgb 2>/dev/null); \
	if [ -n "$$stale" ]; then \
		echo "ERROR: pkg/rt/core_compiled.lgb is stale relative to:"; \
		echo "$$stale" | sed 's/^/  /'; \
		echo "Run 'make generate' to regenerate the bundle."; \
		exit 1; \
	fi

generate: $(GO) generate-ir-ops generate-ir-bridge generate-ir-data pkg/rt/core_compiled.lgb

# Regenerate pkg/ir/op_generated.go from examples/go-gen/ir_ops.lg.
# Requires ./lg to exist (built by `make build`).
generate-ir-ops: build
	./scripts/generate-ir-ops.sh

# Regenerate pkg/rt/ir_bridge_generated.go from examples/go-gen/ir_bridge.lg.
# Requires ./lg to exist (built by `make build`).
generate-ir-bridge: build
	./scripts/generate-ir-bridge.sh

# Regenerate pkg/rt/core/ir/data_generated.lg from examples/go-gen/ir_data.lg.
# Lisp output (mechanical accessor surface for the IR data types).
# Requires ./lg to exist (built by `make build`).
generate-ir-data: build
	./scripts/generate-ir-data.sh

$(LG): $(GO) lg.go pkg/**/* pkg/rt/core_compiled.lgb
	which go
	go build -ldflags="-s -w" -o $@ .

test: pkg/**/* pkg/rt/core_compiled.lgb $(GO)
	$(GO-TEST-ENV) go test $(GO-TEST-FLAGS) -count=1 -v ./test

clojure-compat-report: $(GO)
	@$(REPORT-SCRIPT)

clean:
	$(RM) $(LG)

distclean: clean
ifneq (,$(wildcard .cache))
	chmod -R +w .cache
	$(RM) -r .cache
endif

lint: install-golangci-lint
	golangci-lint run

install-golangci-lint: $(GO)
	which golangci-lint || \
	  GO111MODULE=off go get -u $(GO111MODULE-LINT)

# PHONY targets are for ones that have conflicting files/dirs present:
.PHONY: test check-bundle-fresh
