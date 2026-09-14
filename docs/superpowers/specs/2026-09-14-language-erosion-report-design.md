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

- The human-facing erosion summary is `erosion lg=<number|N/A>
  go=<number|N/A|UNAVAILABLE|NOT REQUESTED>`, never the old unlabeled `erosion
  <number>` prefix. The EDN breakdown retains `:lg`, `:go`, and `:go-source`;
  it does not expose a separate `:erosion :value`. The existing
  `:terms :erosion` remains the auditable combined *debt input* on complete
  reports, not a third language score. The top-level EDN field
  `:score-status` is `:complete` or `:incomplete-go-analysis`.
- When Go paths were requested and the Go analyzer fails, retain the existing
  `.lg` erosion value (numeric when it has callable mass, otherwise `nil`/`N/A`),
  set Go erosion to `nil`, set combined erosion term and top-level debt
  to `nil`, and set `:score-status :incomplete-go-analysis`. Text says `debt
  UNAVAILABLE` and explicitly says Go callables were excluded. The old
  combined `:erosion :callables` and `:high` counts are `nil` on this path;
  text identifies any shown counts as `.lg`-only. EDN carries the same nils
  and status; it never substitutes zero, a renormalized debt, or an invented
  bound.
- “Go not requested” means no non-generated Go file enters the filtered,
  tokenized measured corpus. That is the existing `go-analysis` `:source
  "none"` path, even if default Go roots were configured. Report
  `go=NOT REQUESTED`, `:go nil`, `:go-source "none"`, and
  `:score-status :complete`; a valid `.lg`-only debt remains numeric. When Go
  analysis succeeds, combined debt and term values remain byte-for-byte
  numerically equivalent to the current calculation, with
  `:score-status :complete`.
- `N/A` means zero callable mass for that language among measured files,
  including a Go-only corpus with no `.lg` files; then
  `quality.terms/erosion` returned `nil`. Its EDN language score is
  `nil`, yet `:score-status` stays `:complete` when the analyzer succeeded.
  This is distinct from `NOT REQUESTED` (no measured Go files) and
  `UNAVAILABLE` (a requested analyzer failed). A Go-only corpus may similarly
  show `lg=N/A` without making the result incomplete.
- `compare-results` rejects a base or head with unavailable debt before
  subtracting or reporting a signed delta, with an error identifying the
  incomplete side. No misleading PR delta is emitted.
- The separate `make quality` prerequisite bug that can leave
  `build/go-callables` missing is not changed here. This report contract makes
  that failure visible rather than masking it.

## Verification

Use a fixture containing both `.lg` and Go files with an intentionally missing
`QUALITY_GO_CALLABLES` path. Assert text labels, partial-count labels, EDN
nil/status fields, and a rejected delta. Assert the filtered corpus with no Go
files remains numeric and reports Go not requested; assert Go-only and
successful-Go-with-no-functions cases distinguish `N/A` from failure; assert a
successful Go run retains the prior combined debt. Run focused quality CLI, score, and
delta tests under `TMPDIR=/tmp`, then the short Go suite, build, vet, and
frontmatter check. No jj-dependent test is required.
