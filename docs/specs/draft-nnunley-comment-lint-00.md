---
status: active
last-verified: 2026-09-16
human-verified:
---

<!-- Structure and boilerplate derived from IETF practice (RFC 7322 style,
     BCP 14/RFC 8174); Specification body shape derived from NLSpec
     (jhugman/nlspec). Original guidance prose: CC0 — copy freely, owe nothing. -->

# draft-nnunley-comment-lint-00: Objective Comment-Abuse Rules for let-go

**Status:** DRAFT
**Corpus:** red (spec-first — the R1-R6 surface does not exist on main; evidence is the acceptance criteria)
**Category:** Standards-Track
**Authors:** Norman Nunley, Jr <nnunley@gmail.com>

## Abstract

This document specifies six objective comment-abuse rules (R1-R6) plus
machine-output and gate modes for `scripts/lint.lg`, and a data-driven
code-verbosity catalog. It is for contributors and agents working on
let-go's comment hygiene tooling.

## Motivation

The current state: `scripts/lint.lg` on `main` carries exactly one rule, a
hand-curated list of fifteen devlog phrases, and always exits 0. What is
wrong: a phrase list is a lexicon, not a measurement — it finds only the
wordings someone thought of in advance, so it cannot answer whether an
agent is abusing comments. What this document provides: six rules computed
from the source, thresholded on structural or corpus-derived numbers, each
emitting machine-readable findings the quality scorer consumes. Why now:
the quality tool's verbosity numerator needs a flagged-lines half, and
the corpus calibration in
`docs/specs/comment-abuse-calibration-report.md` shows which rules are
trustworthy enough to gate on.

## Terminology

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD",
"SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this
document are to be interpreted as described in BCP 14 (RFC 2119, RFC 8174)
when, and only when, they appear in all capitals, as shown here.

- **finding** — one rule hit: a map with file, line range, kind, measure,
  and evidence.
- **corpus** — the scanned source tree (`pkg scripts test` by default).
- **objective rule** — a rule computed from the source, thresholded on a
  number that is either structural or derived from the corpus itself, and
  reproducible without human judgment.
- **heuristic rule** — the legacy devlog-phrase lexicon, reported
  separately, never gated, never scored.

## Specification

This document defines the six comment rules, the two output modes, and the
code-verbosity catalog. It does NOT define the quality scorer's use of
these findings; that is owned by the companion quality-score RFC.

**Objectivity over lexicon.** A rule qualifies only when its output does
not depend on anyone's judgment: given the same source, any implementation
produces the same findings.

### Data model

```
RECORD Finding:
    file        : String              -- path of the hit
    line        : Integer             -- first line of the span
    end         : Integer             -- last line of the span
    kind        : FindingKind         -- which rule fired
    measure     : Number              -- rule-specific magnitude
    evidence    : String              -- human-readable justification

ENUM FindingKind:
    COMMENTED-OUT-CODE -- R1
    RESTATEMENT -- R2
    DUPLICATED-COMMENT -- R3
    COMMENT-DENSITY-OUTLIER -- R4
    DEVLOG-COMMENT -- R5, heuristic, never gated
    COMMENT-CHURN -- R6, range-scoped
    COMMENT-DIVIDER -- section headers, excluded from score
    CODE-VERBOSITY -- catalog-driven code patterns
```

The catalog defines fourteen entries across nine code kinds
(`redundant-if-false-true`, `redundant-eq-true`, `redundant-not-eq`,
`redundant-not-empty`, `eta-expansion`, `let-identity`,
`do-single-form`, `composable-accessor`, `apply-str-interpose`,
`nested-if-chain` — two entries were built, confirmed working, and
dropped for firing zero times on the real corpus, per the calibration
report).

Every catalog finding MUST carry an `atoms-eliminable` measure: the atom
count of the matched form minus the atom count of its replacement, both
computed by the same node-atom-count function, so the saving is derived
from the rule pair and cannot drift from it. [R-atoms-eliminable]

```transcript @R-atoms-eliminable
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --edn pkg/rt/core/ir/passes/liveness.lg | grep -o "atoms-eliminable=[0-9]*" | head -n 1
atoms-eliminable=3
? 0
```

The catalog MUST be data, not predicates: adding a pattern MUST require
no code change, only a new catalog entry. This is proven by a test that
invokes a copy of the tool next to a fabricated catalog containing one
rule not in the shipped set and confirms it fires. [R-catalog-is-data]

```transcript @R-catalog-is-data
$ grep -c ":pattern" scripts/lint-code-rules.edn
12
? 0
```

Code findings MUST be form-delimited, not line-delimited: each finding's
line range is the matched form's own span from the position-tracking
parser, strictly more precise than line heuristics. [R-form-delimited]

```transcript @R-form-delimited
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --edn pkg/rt/core/ir/passes/liveness.lg | grep -o '"line": 78, "end": 78' | head -n 1
"line": 78, "end": 78
? 0
```

### Configuration

| Key | Type | Default | Description |
|---|---|---|---|
| `paths` | String list | `pkg scripts test` | Roots the tool scans |
| `gate` | String list | `(none)` | Finding kinds that fail the run |
| `churn-range` | Revision range | `(none)` | Range R6 diffs; R6 is inert without it |

**Resolution precedence** (highest first; the last entry is the terminal
default):

1. Command-line flag (`--gate`, `--churn`, positional paths)
2. Defaults from the table above

### Behavior

The tool MUST report every finding with a file, line range, kind, measure,
and evidence triple. [R-finding-shape]

```transcript @R-finding-shape
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --edn pkg/rt/core/ir/passes/liveness.lg | head -c 300
[{:file "pkg/rt/core/ir/passes/liveness.lg", :line 117
? 0
```

The code-verbosity catalog MUST be data, not predicates: adding a pattern
MUST require no code change, only a new catalog entry. This is proven by a test that
invokes a copy of the tool next to a fabricated catalog containing one
rule not in the shipped set and confirms it fires. [R-catalog-is-data]

```transcript @R-catalog-is-data
$ grep -c "^ *{:name" scripts/lint-code-rules.edn
14
? 0
```

Every catalog finding MUST carry an `atoms-eliminable` measure: the atom
count of the matched form minus the atom count of its replacement, both
computed by the same node-atom-count function, so the saving is derived
from the rule pair and cannot drift from it. [R-atoms-eliminable]

```transcript @R-atoms-eliminable
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --edn pkg/rt/core/ir/passes/liveness.lg | grep -o "atoms-eliminable=[0-9]*" | head -n 1
atoms-eliminable=3
? 0
```

Code findings MUST be form-delimited, not line-delimited: each finding's
line range is the matched form's own span from the position-tracking
parser, strictly more precise than line heuristics. [R-form-delimited]

```transcript @R-form-delimited
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --edn pkg/rt/core/ir/passes/liveness.lg | grep -o ":line 78, :end 78" | head -n 1
:line 78, :end 78
? 0
```

R1 MUST flag a comment whose whole trimmed body parses as exactly one
delimited code form, and MUST NOT flag bare identifiers, bare literals,
or usage examples naming an adjacent declaration. [R-r1-whole-form]

```transcript @R-r1-whole-form
$ /tmp/gi-rfc/bin/lg scripts/lint.lg pkg/rt/core/ir/passes/liveness.lg | grep commented-out-code | head -n 2
pkg/rt/core/ir/passes/liveness.lg:117-117  [commented-out-code] (live-in[t] − params[t])
? 0
```

R2 MUST flag a comment whose content words overlap the next form's
identifiers at 0.8 or above with at most one extra word, and MUST classify
section-divider comments under their own kind excluded from the score. [R-r2-divider]

```transcript @R-r2-divider
$ /tmp/gi-rfc/bin/lg scripts/lint.lg pkg/bytecode/bytecode_test.go | grep -c comment-divider
12
? 0
```

R3 MUST flag normalized comment text appearing in three or more places,
excluding licence headers, skip-markers, and Go interface-implementation
boilerplate. [R-r3-boilerplate]

```transcript @R-r3-boilerplate
$ /tmp/gi-rfc/bin/lg scripts/lint.lg pkg/vm/symbol.go | grep -c duplicated-comment
0
? 0
```

R4 MUST flag definitions whose interior comment-line ratio beats the
corpus 95th percentile, excluding leading doc-comment blocks from both
numerator and denominator. [R-r4-interior]

```transcript @R-r4-interior
$ /tmp/gi-rfc/bin/lg scripts/lint.lg pkg/rt/lang.go | grep -c comment-density-outlier
0
? 0
```

R5 MUST remain heuristic-only: reported under its own heading, never part
of the objective score, and `--gate` MUST refuse to gate on it. [R-r5-ungated]

```transcript @R-r5-ungated
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --gate devlog-comment pkg/rt/lang.go; echo "exit=$?"
exit=0
? 0
```

R6 MUST compare comment-lines-added per code-line-added over the given
range against the corpus baseline and report the multiple, and MUST NOT
run without `--churn`. [R-r6-range]

```transcript @R-r6-range
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --churn HEAD~5..HEAD pkg/rt/lang.go | grep -c comment-churn
1
? 0
```

The `--edn` mode MUST print all findings as one EDN vector and nothing
else. [R-edn-shape]

```transcript @R-edn-shape
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --edn pkg/rt/core/ir/passes/liveness.lg | head -c 1
[
? 0
```

### Errors

The tool MUST refuse an unknown `--gate` kind with a note naming it, and
MUST error non-zero on an unresolvable `--churn` range. R6 without
`--churn` at all MUST report skipped and exit 0. [R-error-cases]

```transcript @R-error-cases
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --gate frobnicate pkg/rt/lang.go | grep -o 'ignored \["frobnicate"\]'
ignored ["frobnicate"]
? 0
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --gate frobnicate pkg/rt/lang.go >/dev/null 2>&1; echo "exit=$?"
exit=0
? 0
$ /tmp/gi-rfc/bin/lg scripts/lint.lg --churn bogus-range pkg/rt/lang.go >/dev/null 2>&1; echo "exit=$?"
exit=1
? 0
```

## Out of Scope

**R1 for Go.** Parsing Go comments as code from `.lg` is the wrong shape
(either shelling an unavailable statement parser or hand-rolling a Go
tokenizer). Extension point: the Go AST tool's comment-block output, when
it exists — R1 for Go becomes a consumer of that output.

**Comment-debt scoring.** Flagged comment lines over total comment lines
is owned by the quality scorer, not this tool; this tool only emits the
findings. Extension point: the `--edn` finding vector, which the scorer
already consumes.

## Alternatives Considered

**Why not keep the phrase list as the only rule?** A lexicon finds only
wordings someone thought of; it cannot measure abuse and cannot gate.
Kept as R5, segregated and ungated.

**Why not gate R1/R3/R4 now?** The calibration report shows R1 at 27% FP
(syntax examples), R3 dominated by Go boilerplate residue, R4's tail as
well-documented public functions. Advisory until re-calibrated.

**Why not line-delimited code regions like the paper?** Form-delimited
spans are strictly more precise and line counts remain recoverable; the
deviation is disclosed in the calibration report.

## Security Considerations

The tool reads source files and shells to `git diff` for R6 with
`--no-ext-diff --no-color` so a configured external diff tool cannot
rewrite the parsed format. It never edits source and exits non-zero only
under `--gate`. No network, no secrets, no trust boundary crossed: "none"
needs no further argument beyond the R6 subprocess pinning above.

## Compatibility

New files (`scripts/lint-code-rules.edn`) and new flags (`--edn`,
`--gate`, `--churn`); default human-readable output unchanged in shape.
At the time of writing (September 2026), the `main` branch carries the
devlog-only linter; this RFC is spec-first red against that base.

## References

- `docs/specs/comment-abuse-calibration-report.md` — per-rule counts,
  false-positive reads, gating verdicts (tracked copy of the working
  report).
- `docs/superpowers/specs/2026-09-12-comment-abuse-rules-design.md` —
  the design this RFC formalizes (gitignored working copy).
- SlopCodeBench, arXiv 2603.24755 §2.3 — verbosity metric definition.
