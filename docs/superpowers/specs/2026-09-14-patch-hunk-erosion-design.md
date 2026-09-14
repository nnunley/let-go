---
status: active
last-verified: 2026-09-14
---

# Reduce Git patch parser erosion without changing its contract

## Context

On published #871, `quality.diff/parse-patch-hunks` has CC 50, 85 SLOC,
and mass 461 (about 1.2% of corpus mass). The corpus erosion share is 0.414:
122 of 4095 callables exceed CC 10. This function combines two independent
tasks: scanning Git patch sections and reconciling those sections with
NUL-delimited name-status records.

## Design

Keep the public `parse-patch-hunks` signature, ordered return records, and
exception behavior. Extract a section scanner with a small line classifier for
hunks, binary markers, gitlink commit lines, blob/symlink/gitlink mode lines,
and new/deleted file markers. Keep metadata flags on section records rather
than introducing a new output format. Extract status reconciliation and row
construction: verify every section header against the authoritative NUL
status path, coalesce only the matched two-section file/symlink type change,
and preserve zero-count hunk anchors and path-only gitlinks. The public fn
becomes a short composition of these two phases.

The seams must represent domain operations, not one-line helpers invented to
lower a score. Helpers should have focused names and small branching surfaces.
No `jj` assumption, working-tree path convention, CLI output change, or
production benchmark change is in scope.

## Alternatives

- Start with `quality/measure`: larger mass (2.6%), but a wide orchestration
  fn with more contracts and a harder regression boundary.
- Start with compiler `inst-rhs`: 2.0% mass, but a runtime-critical lowering
  path; behavioral risk is much higher for a first erosion tranche.

## Evidence and success

Add characterization cases for section metadata and error paths not already
covered by `test/quality_diff_test.lg`, including `ex-info` data keys
`:sections/:statuses`, `:status/:header`, and malformed-hunk `:line`; run them
red before implementation if
they expose a missing contract, otherwise retain the existing passing
behavioral suite as the baseline. Run the focused quality-diff and quality
delta/touched suites, Go short suite, build/vet, and a before/after `make
quality` capture from the same branch/toolchain. Success requires unchanged
fixture behavior and a measured decrease in both `parse-patch-hunks` CC and
corpus erosion; if extraction merely shifts high CC to a helper or worsens the
composite score, revise it rather than claim a win. Keep this on a separate
branch based on #871, not in the published PR stack.
