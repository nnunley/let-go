---
status: active
last-verified: 2026-09-14
---

# Language-specific erosion reporting

## Decision

Report erosion separately for let-go (`.lg`) and Go. A missing Go analysis is
not a low Go score and must never turn the combined quality debt or a PR delta
into an apparent improvement. The existing combined erosion arithmetic remains
an input to quality debt when all requested analyses succeed; this change does
not create two new composite-debt systems or change weights.

## Current failure

`quality/go-analysis` returns `:source "unavailable"` and no functions when its
command fails. `measure` currently computes aggregate erosion from the remaining
`.lg` callables and still emits a numeric `:debt`. On the same corpus, that made
erosion appear to fall from about 0.596 to 0.414 solely because Go data was
missing. The text prints `go=UNAVAILABLE`, but neither the aggregate erosion nor
debt is marked incomplete. `compare-results` then subtracts debts without
checking completeness.

## Result contract

- The human-facing erosion summary is `lg=<number> go=<number|UNAVAILABLE|NOT
  REQUESTED>`, not an unlabeled aggregate erosion number. The EDN breakdown
  retains `:lg`, `:go`, and `:go-source`; it does not expose a separate
  `:erosion :value`. The existing `:terms :erosion` remains the auditable
  combined *debt input* on complete reports, not a third language score.
- When Go paths were requested and the Go analyzer fails, retain numeric `.lg`
  erosion, set Go erosion to `nil`, set combined erosion term and top-level debt
  to `nil`, and mark the result incomplete. Text says `debt UNAVAILABLE` and
  explicitly says Go callables were excluded. EDN carries the same nils and
  status; it never substitutes zero, a renormalized debt, or an invented bound.
- When no Go paths were requested, Go is `NOT REQUESTED`/`nil`; a valid
  `.lg`-only debt remains numeric. When Go analysis succeeds, combined debt and
  term values remain byte-for-byte numerically equivalent to the current
  calculation.
- `compare-results` rejects a base or head with unavailable debt before
  subtracting or reporting a signed delta, with an error identifying the
  incomplete side. No misleading PR delta is emitted.
- The separate `make quality` prerequisite bug that can leave
  `build/go-callables` missing is not changed here. This report contract makes
  that failure visible rather than masking it.

## Verification

Use a fixture containing both `.lg` and Go files with an intentionally missing
`QUALITY_GO_CALLABLES` path. Assert text labels, EDN nil/status fields, and a
rejected delta. Assert an `.lg`-only request remains numeric and a successful
Go run retains the prior combined debt. Run focused quality CLI, score, and
delta tests under `TMPDIR=/tmp`, then the short Go suite, build, vet, and
frontmatter check. No jj-dependent test is required.
