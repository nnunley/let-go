---
status: active
last-verified: 2026-09-16
human-verified:
---

<!-- Structure and boilerplate derived from IETF practice (RFC 7322 style,
     BCP 14/RFC 8174); Specification body shape derived from NLSpec
     (jhugman/nlspec). Original guidance prose: CC0 — copy freely, owe nothing. -->

# draft-nnunley-quality-score-00: Bounded Debt Composite for let-go

**Status:** DRAFT
**Corpus:** red (spec-first — scripts/quality.lg does not exist on main; evidence is the acceptance criteria)
**Category:** Standards-Track
**Authors:** Norman Nunley, Jr <nnunley@gmail.com>

## Abstract

This document specifies the bounded [0,1] debt composite (`scripts/quality.lg`
plus `scripts/quality/*` and `cmd/go-callables`): eight corpus terms, a
per-file table, ranked issues, and a PR-delta mode. It is for contributors
and agents measuring or optimizing let-go code quality.

## Motivation

The current state: no quality scorer exists on `main` — there is nothing to
run, so "is this change making things worse" has no numeric answer. What is
wrong: without a deterministic, bounded composite, agents optimize against
vibes and reviewers cannot compare layers of a stack. What this document
provides: eight terms with fixed weights, per-file shares, impact-ranked
issues, and a signed PR delta, all reproducible from source at a commit.
Why now: the lint stack supplies the flagged-lines half of verbosity, and
the bounded-score design supersedes the abandoned unbounded debt model.

## Terminology

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD",
"SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this
document are to be interpreted as described in BCP 14 (RFC 2119, RFC 8174)
when, and only when, they appear in all capitals, as shown here.

- **debt** — a share in [0,1] where 0 is perfect and 1 is worst; lower is
  better, always.
- **term** — one of the eight corpus components (erosion, uncovered,
  verbosity, comment-debt, dead, defects, testing, cost).
- **mass** — `CC * sqrt(SLOC)`; complexity weighted by size.
- **PR delta** — base/head comparison with signed deltas; positive means
  worse.

## Specification

This document defines the eight terms, the composite, the per-file table,
issues, and PR-delta mode. It does NOT define the lint rules whose
findings feed verbosity and comment-debt; those are owned by the companion
comment-lint RFC.

**Bounded and unsaturating.** Every scored number is a share in [0,1] —
bounded by construction, strictly monotone, no clamping dead zone.

### Data model

```
RECORD CorpusTerms:
    erosion       : Number              -- mass share above CC 10
    uncovered     : Number              -- 1 minus coverage
    verbosity     : Number              -- flagged+clone lines over LOC
    comment-debt  : Number              -- flagged comment lines over comment lines
    dead          : Number              -- dead SLOC over all SLOC
    defects       : Number              -- squashed fix-commit rate
    testing       : Number              -- density+ratio shortfall squash
    cost          : Number              -- mass-weighted superlinear share

ENUM TermOmission:
    OMITTED       -- input missing; remaining weights renormalize
```

### Configuration

| Key | Type | Default | Description |
|---|---|---|---|
| `paths` | String list | per-language defaults | Roots measured |
| `top` | Integer | `20` | Rows in top-complexity table |
| `base` / `head` | Revisions | `(none)` | PR-delta endpoints; absent = snapshot |

The `thresholds` map in `quality.score` MUST be the single source of every
constant; nothing else hardcodes a number. [R-thresholds-single-source]

```transcript @R-thresholds-single-source
$ grep -rn "0\.30" scripts/quality/*.lg | grep -v ":weights"
? 1
```

### Behavior

Complexity MUST count branches per definition: 1 plus one per branch
point (`if`/`when`/`loop` family, extra `cond`/`case` clauses, extra
`and`/`or` operands, threading steps, `catch`). [R-complexity-branches]

```transcript @R-complexity-branches
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -c "cc 26"
2
? 0
```

Erosion MUST equal the mass share above CC 10 with strict inequality.
[R-erosion-strict]

```transcript @R-erosion-strict
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -c erosion
4
? 0
```

Verbosity MUST union flagged code lines with clone lines over LOC, and
MUST omit the flagged half loudly when lint cannot run. [R-verbosity-union]

```transcript @R-verbosity-union
$ QUALITY_LINT=/nonexistent LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -c "CLONE LINES ONLY"
1
? 0
```

Comment-debt MUST divide by comment lines, not LOC. [R-comment-debt-denominator]

```transcript @R-comment-debt-denominator
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -c comment-debt
2
? 0
```

Dead code MUST be SLOC-weighted with production roots and MUST report
test-only reachability separately. [R-dead-roots]

```transcript @R-dead-roots
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -c "reachable only from tests"
1
? 0
```

Untested findings MUST print qualified names, since bare names are
ambiguous across namespaces. [R-qname-hints]

```transcript @R-qname-hints
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -c "no test reaches"
109
? 0
```

The composite MUST be the weighted mean of present terms with weights
erosion 0.30, uncovered 0.20, verbosity 0.10, comment-debt 0.10, dead
0.10, defects 0.10, testing 0.10, cost 0.0, renormalized over what is
present. [R-composite-weights]

```transcript @R-composite-weights
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -E "^  erosion +0"
  erosion       0.598  weight 0.30
? 0
```

Cost MUST be reported at weight 0 until loop-source tracing lands; the
structural degree MUST NOT be confused with a loop-nesting count.
[R-cost-parked]

```transcript @R-cost-parked
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -E "^  cost +0"
  cost          0.036  weight 0.00
? 0
```

Issues MUST rank by estimated composite impact, worst first, so the first
line is the best next fix. [R-issues-ranked]

```transcript @R-issues-ranked
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -c "best targets"
1
? 0
```

PR deltas MUST be signed with positive meaning worse, and MUST name
threshold crossings as the actionable half. [R-delta-signed]

```transcript @R-delta-signed
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --help 2>/dev/null | grep -c "base REV"
1
? 0
```

Maintainability index MUST stay a per-file diagnostic and MUST NOT enter
the composite: the SEI formula double-counts what erosion already
carries. [R-mi-diagnostic]

```transcript @R-mi-diagnostic
$ LG_SOURCE_PATHS=scripts ./bin/lg scripts/quality.lg --source-root test/fixtures/quality test/fixtures/quality 2>/dev/null | grep -c "mi(raw)"
1
? 0
```

### Errors

| Error | Example | Recovery |
|---|---|---|
| `E-NO-SOURCE` | paths match nothing tracked | Error; exit 2, never a perfect score |
| `E-NO-LG-BINARY` | lint needs `lg`, none found | Omit flagged half; report loudly |
| `E-NO-COVER` | no `--go-cover` | Omit Go dynamic term; renormalize |

## Out of Scope

**Efficiency term.** Dividing coverage by CI seconds made the composite
move with machine load; removed, reported as points only. No extension
point — the decision is final.

**Dynamic `.lg` coverage.** Static reachability is an upper bound until a
VM hit-count hook lands. Extension point: the coverage term's lg-static
input.

**Cost promotion.** Weight 0 until iteration sources trace to parameters.
Extension point: `quality.cost` loop-source rule.

## Alternatives Considered

**Why not the unbounded 0–100 debt model?** Clamping made complexity 30
and 200 score identically, so fixing the worst file registered nothing.
Shares fix the defect the shape had, not boundedness.

**Why not score MI in the composite?** The SEI formula double-counts what
erosion carries; MI stays a per-file diagnostic.

**Why not fold comments into verbosity?** The paper's rules are code
constructs; comment slop on LOC vanishes into rounding error. Comment-debt
gets its own denominator instead.

## Security Considerations

The tool shells to `git`/`jj`, the Go toolchain, and the linter — all
local subprocesses with pinned inputs echoed in the header. No network
(except an explicit `--ci-run` `gh api` lookup), no secrets. The `os/sh`
re-parenting gap (killed parents orphan children) is documented in the
tool's own header; it cannot corrupt a score, only leak a process.

## Compatibility

New tool (`scripts/quality.lg`, `scripts/quality/*`, `cmd/go-callables`,
`make quality`); no existing behavior changes. EDN keeps every raw field
for trend tracking. At the time of writing (September 2026), `main`
carries none of this; the RFC is spec-first red against that base.

## References

- `docs/superpowers/specs/2026-09-12-quality-bounded-score-design.md` —
  the bounded design this RFC formalizes (gitignored working copy).
- `docs/superpowers/specs/2026-09-11-code-quality-score-design.md` — the
  original tool design both RFCs build on (gitignored working copy).
- SlopCodeBench, arXiv 2603.24755 §2.3 — erosion/verbosity definitions.
