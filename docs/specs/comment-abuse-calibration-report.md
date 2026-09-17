---
status: active
last-verified: 2026-09-16
authoritative-for:
  - lint-calibration
human-verified:
---

# Objective comment-abuse rules — implementation report

Date: 2026-09-12
Scope: `scripts/lint.lg`, `test/lint_lg_test.go` only (per brief).
Spec: `docs/superpowers/specs/2026-09-12-comment-abuse-rules-design.md`
(Copy of the original report; the working copy lives in the gitignored
`.superpowers/sdd/comment-abuse-report.md`. This tracked copy is what
`scripts/lint.lg`'s "see the report" references point at.)

## Status

All six rules (R1-R6) plus `--edn` and `--gate` are implemented and tested.
R1 for Go is explicitly **deferred**, not implemented weakly, per the
coordinator's mid-task correction (see "Scope correction" below).

## Commits (one per rule, in order)

| Rule | Commit | Summary |
|---|---|---|
| R1 | `7ca047d2` | commented-out-code, `.lg` only |
| R2 | `2ec621e7` | restatement |
| R3 | `196cd007` | duplicated-comment |
| R4 | `44213e1a` | comment-density-outlier |
| R5 + `--edn`/`--gate` | `178f9f50` | segregate heuristic, add machine output/gate modes |
| R6 | `7f408728` | comment-churn + `--churn` flag |

(These are the `jj` change commit ids as recorded at commit time; `jj log`
in this repo will show them under the working-copy history.)

## Test summary

`go test ./test/ -run 'TestLint' -v -count=1` — **7/7 PASS**:
`TestLintLgSkipMarkersAndPhraseMatching` (R5, pre-existing, unchanged),
`TestLintR1CommentedOutCode`, `TestLintR2Restatement`,
`TestLintR3DuplicatedComment`, `TestLintR4DensityOutlier`,
`TestLintEdnAndGate`, `TestLintR6CommentChurn`.

Every rule was driven test-first: a failing fixture-based test was run and
its actual RED output captured before implementation, then GREEN was
confirmed after. One rule (R3) caught a real bug in my own fixture during
the RED→GREEN cycle (consecutive comment lines merge into one block, so an
adjacent licence header and the intended duplicate text were one block and
both got excluded by the skip-marker check — fixed by separating them with
a blank line in the fixture, not in `lint.lg`). One infrastructure bug was
found and fixed while testing R6 manually: a globalally-configured
`diff.external` (difftastic) rewrites `git diff`'s output for **any**
subprocess, not just interactive use, so `os/sh "git" "diff" ...` was
returning difftastic's format instead of unified diff. Fixed by adding
`--no-ext-diff --no-color` to the invocation; this would have caused a
silent, permanent 0-finding R6 on any machine with a similar git config.

## Scope correction mid-task

R1 originally attempted a shared "code-shaped" structural test (delimiter
balance + reader/parser success) intended to cover both `.lg` and Go. The
coordinator corrected this: the architecture keeps `.lg` analysis in `.lg`
and puts all Go analysis in a separate Go tool being built by another
agent. Building even a partial Go parse from `.lg` (only `gogen/type`,
i.e. `go/parser.ParseExpr`, is reachable from `.lg`, and it only parses
*expressions*, not statements — confirmed by hand: `x := 5`, `return err`,
and `if err != nil {...}` all fail `gogen/type`, while `foo(bar)` and bare
literals succeed) would have been the wrong shape and thrown away. **R1 is
therefore implemented for `.lg` only.** `scripts/lint.lg`'s own output
says so explicitly on every run (`"N Go file(s) not covered yet"`), and a
Go file is never silently reported as "checked and clean" for R1.

Follow-up noted for whoever picks this up: a sibling task is building a Go
analysis command that will emit comment blocks with line ranges alongside
function records. When that exists, R1 for Go becomes a consumer of its
output, not new parsing work in `scripts/lint.lg`. I did not build that
command and did not coordinate with the agent building it, per instruction.

## Calibration — per-rule corpus counts and false-positive read

Corpus: `pkg scripts test` (903 files), the same default the tool scans.
For each rule I ran it, counted, and read real findings (not just the
first ten by file-sort order, in a couple of cases, to represent the
finding population more accurately than an artifact of directory-walk
order — noted below where relevant).

### R1 — commented-out-code (`.lg` only): 19 findings

First implementation (structural test = "has a balanced delimiter pair,
and `read-string` doesn't throw") produced **1199** findings, and the top
10 were **100% false positives** — doc comments quoting syntax examples or
containing an incidental parenthetical aside, e.g.
`pkg/ir/ir_bridge.lg:190 "...(e.g. ir/op vs op)..."`. Root cause: almost
any prose reads as a sequence of valid Lisp symbols without erroring
(`read-string` only reads and validates the *first* form, ignoring
trailing text), so "reads without error" barely discriminates anything —
the only real signal was ever the delimiter structure, and a parenthetical
*inside* a sentence isn't the same as *being* a form.

I tightened the rule (documented in `scripts/lint.lg` itself) to require
the **whole trimmed comment body** to be exactly one delimited group —
starts with an opener, ends with its matching closer — not merely to
*contain* a balanced parenthetical. This is still fully structural, no
lexicon. Re-run: **19 findings**.

Read all 19. **11 true positives, 8 false positives (42% FP rate
overall)**, but the false positives cluster at the front of the file-sorted
output (`pkg/` before `test/`), so **the literal "top ten" is 8 FP / 2 TP
(80% FP)**:
- False positives (8): all in `pkg/ir/defop_spec.lg`, `pkg/ir/ir_bridge.lg`,
  `pkg/rt/core/ir/lower_go.lg`, `pkg/rt/core/ir/passes/liveness.lg`,
  `pkg/rt/gogen/gogen.lg` — doc comments showing a **usage-syntax example**
  as a single parenthesized form, e.g. `;; (defmethod name dispatch-val
  [args] body)`, or terse math notation like `;; (live-in[t] − params[t])`.
  These are lexically identical to real leftover code; there is no
  structural way to tell "an example in a docstring" from "code someone
  forgot to delete" without semantics.
- True positives (11): genuinely commented-out calls, all in `test/`, e.g.
  `test/hello.lg:13 "; (run-tests basic)"`, five equivalent
  `; (run-tests ...)` lines across `test/ns/*.lg`, and two large commented
  test blocks in `test/simple.lg` (confirmed by reading the surrounding
  file — real dead test code, not documentation).

**Verdict: R1 is correctly implemented and objective, and it does find
real commented-out code with zero false negatives I could find, but it
also has a real, structural, un-fixable-without-semantics false-positive
class (syntax examples in doc comments). Per the spec's own suggestion
that R1 is "the rule most defensible as a blocking gate" — I disagree
based on this corpus: I would not gate CI on R1 as-is without first
excluding files that are mostly documentation-of-syntax (e.g.
`defop_spec.lg`, `gogen.lg`), which would mean an exclusion list — the
kind of hand-curated judgment the whole project is trying to avoid.
Advisory use only, for now.**

### R2 — restatement: 166 findings

Read the first 10 plus a few more. **0 of the ones read are wrong per the
rule's own definition** — every one I checked (`pkg/bytecode/decoder.go`
"`// Read extra map`" above `d.readMapValue()`; `pkg/ir/ir_bridge.lg`
"`;; --- emit a field getter ---`" above `(defn emit-getter ...)`;
`pkg/ir/lisp_fusion_test.go` "`// Validate IR.`" above an
`ir.validate/validate-fn!` call; several `pkg/bytecode_test.go` section
headers like "`// --- Func roundtrip ---`" above
`func TestFuncRoundtrip(...)`) genuinely restates the next line's
identifiers and adds nothing else.

**Caveat, not a defect:** a large share of the 166 are conventional
section-divider comments in test files ("`--- X roundtrip ---`" above
`TestXRoundtrip`), which is common, low-stakes style rather than the
narrative "devlog" abuse the spec is chasing. The rule is objectively
correct; whether "restates the next line" is the right proxy for "abuse"
in a codebase that likes section headers is a judgment call for whoever
sets the gate threshold. **Verdict: low false-positive rate against its
own definition; recommend advisory, worth a second look before gating
specifically because of the section-header pattern.**

### R3 — duplicated-comment: 251 findings

Read the frequency-ranked list (`awk`'d the quoted text out and counted
duplicates): the dominant pattern by far is **Go interface-implementation
one-liners** — `"Type implements Value"` (20x), `"Unbox implements
Value"` (14x), `"Unbox implements Unbox"` (12x), `"Meta implements
IMeta."` (11x), `"Nth implements Indexed: positional access by integer
index."` (7x), `"Next/More/First/Empty/Count/Cons implements
Seq/Collection"` (7x each), and similar — i.e., **idiomatic Go doc-comment
boilerplate repeated once per type implementing an interface method.**
This is not narrative slop; it is a widely accepted Go documentation
convention (godoc conventionally starts a method's comment with the
method name).

**Verdict: the mechanism is exactly correct per its literal definition
(same normalized text, 3+ occurrences, license headers excluded), but the
practical false-positive rate for "abuse" is high — I estimate the large
majority of the 251 are this idiomatic class, not agent-pasted narration.
I did not add an exclusion for it, because doing so (e.g., special-casing
an "X implements Y" pattern) would reintroduce exactly the
hand-curated-lexicon approach the whole design is trying to escape. I am
reporting this plainly: do not gate R3 without either accepting this
false-positive class or find a structural (not lexical) way to exclude
one-line "boilerplate" comments — e.g., only count duplicates whose
comment block also has >1 line, which I did not have time to try and
calibrate in this pass.**

### R4 — comment-density-outlier: 129 findings, p95 = 0.8889 over 2,871 definitions

Read the first 15+ across `pkg/api/api.go`, `pkg/api/ns.go`,
`pkg/bundle/bundle.go`, `pkg/bytecode/*.go`, `pkg/cli/*.go`,
`pkg/compiler/*.go`. Every one read is a **small, well-documented public
API function** — e.g. `pkg/api/api.go:WithStdout` and `:WithEmit`, each a
1-line-body function with 12-20 lines of careful godoc explaining
semantics and a concurrency caveat. This is high-quality documentation,
not abuse.

I tried raising the minimum-definition-size guard from 5 to 20 lines to
see if larger definitions would surface a more abuse-shaped tail (the
intuition: a big function that's still 90% comment is a stronger signal
than a small one with normal godoc). It changed the corpus but I did not
have a clean second calibration pass before the effort budget on this
task ran out, so I reverted to the spec's literal 5-line guard rather than
ship an untested threshold change.

**Verdict: mechanism correctly implemented (percentile self-calibrates
off the corpus, no magic constant), but on this corpus the 95th-percentile
tail is dominated by well-documented small public functions, not
comment-abuse. R4 is NOT ready to gate as calibrated. If this rule is to
be useful, the next step is almost certainly excluding exported
identifiers with godoc (a very different, much larger investigation than
this pass had budget for) or requiring a much larger minimum definition
size and re-calibrating — I did not do that work.**

### R5 — devlog-phrase (heuristic): 21 findings, behavior unchanged

Pre-existing test (`TestLintLgSkipMarkersAndPhraseMatching`) still passes
unmodified. Not gated, not part of the objective score, per spec.

### R6 — comment-churn: opt-in via `--churn <range>`, no corpus-wide count

R6 only runs against an explicit range, so there's no "finding count over
the corpus" the way R1-R4 have one. Sanity-checked against two real
ranges in this repo: `--churn HEAD~20..HEAD` → range-ratio 0.225,
baseline 0.079, multiple **2.86x**; a wider range gave multiple **3.88x**.
Both are plausible, non-degenerate numbers; I have no known-abusive commit
range in this repository to validate against a expected "should clearly
flag" case, so I can't report a false-positive rate the way I did for
R1-R4. The baseline itself is a documented simplification (current corpus
ratio as a proxy for "typical," not a full per-commit historical walk —
computing the latter would mean diffing every commit in the repo's
history on every lint invocation).

## Known implementation limitations (disclosed, not silently absorbed)

- **R1 covers `.lg` only.** Go is unimplemented by design (see "Scope
  correction" above), and the tool's own output says so on every run.
- **R1 and R4's delimiter/brace balancing does not special-case string
  literals** inside a comment body (a `"("` inside a quoted string would
  be counted as a real delimiter). Did not find this cause any actual
  false positive/negative in the calibration reads, but it's a known
  imprecision.
- **R2 only examines the single source line immediately following a
  comment block**, not the full extent of a multi-line form. Sufficient
  for every real case I found in this corpus (most restatement patterns
  are single-line function/call signatures), but a comment restating a
  form whose distinguishing identifiers appear only on a later line of a
  multi-line form would be missed (false negative, not false positive).
- **R4's definition-extent detection treats all three bracket kinds
  (`()[]{}`) as one undifferentiated depth counter** rather than validating
  type-matching, for both `.lg` and Go — a deliberate simplification
  consistent with R1.
- **R6's baseline is a corpus snapshot, not a historical walk** (see R6
  section above).
- Full-corpus run time: ~59s for `pkg scripts test` (903 files) with all
  rules including the O(n²)-ish R3/R4 corpus passes. Acceptable for a CI
  gate step, not for a pre-commit hook on every file save; did not
  optimize further given the effort budget.

## Recommendation on gating

Per spec, nothing gates by default; `--gate <kind>[,...]` is opt-in. Based
on calibration:
- **R1** (`commented-out-code`): the only one I'd consider gating soon,
  but only after excluding syntax-heavy doc files (`defop_spec.lg`,
  `gogen.lg`, etc.) from the gate path, or accepting an 80%-in-the-common-
  case false-positive rate on files like those.
- **R2** (`restatement`): reasonably trustworthy; the section-header
  caveat is a style judgment, not a bug.
- **R3** (`duplicated-comment`): NOT ready — dominated by idiomatic Go
  interface-boilerplate false positives.
- **R4** (`comment-density-outlier`): NOT ready — dominated by
  well-documented small public API false positives.
- **R6** (`comment-churn`): mechanism sound, insufficient data in this
  corpus to state a false-positive rate; recommend advisory-only until
  validated against a change known to be comment-heavy in practice.

## Integration note (not built, per instructions)

`--edn` prints the full finding vector (all rules, with `:line`/`:end`
ranges) machine-readable. This is intended to feed the quality tool's
verbosity numerator (flagged-lines ∪ clone-lines) per the design doc's
Integration section. I did not touch `scripts/quality.lg` or
`scripts/quality/**`, and did not wire the consumer side — another agent
owns that and is actively working there.

---

# Round 2 — false-positive fixes, code-verbosity, and magnitude

Commit: `9c191860` (fixes + code-verbosity catalog + tests, one combined
commit — the working copy mixed both in one file and jj has no clean
partial-file split without an interactive session, so I did not force an
artificial split).

Tests: `go test ./test/ -run 'TestLint' -v -count=1` — **10/10 PASS**
(the original 7 plus `TestLintR3ExcludesInterfaceBoilerplate`,
`TestLintCodeVerbosityCatalog`, `TestLintCodeVerbosityCatalogIsData`).

## False-positive fixes — before/after

### R1 commented-out-code: 19 (8 FP, 42%) -> 11 (3 FP, 27%)

Fix: exclude a comment whose parsed head symbol appears on the immediately
adjacent (next OR previous) source line — the "usage example directly
above/below its own declaration" shape every original false positive
actually had (e.g. `;; (defmethod name dispatch-val [args] body)` above
`(defmacro defmethod ...)`). I checked the two alternatives suggested
("exclude comments inside doc blocks", "require no prose precedes the
parsed content in the block") against the actual false positives first:
neither would have excluded any of them, since none had preceding prose in
the same block. This targets the mechanism the data actually showed instead.

Remaining 3 false positives (`pkg/rt/core/ir/lower_go.lg:2596`, `:2719`,
`pkg/rt/core/ir/passes/liveness.lg:117`) are prose-in-parens and math
notation with no adjacent-line self-reference to key off — a smaller,
harder residue. Top-10 (all 11 findings): 8 true positives (all in
`test/`, genuine dead `(run-tests ...)` / commented `deftest` calls,
manually confirmed by reading the surrounding file), 3 false positives —
27% overall, all in the tail rather than the front this time.

### R2 restatement: 166 -> 65 restatement + 102 comment-divider

Fix: `divider-shaped?` — trimmed text bracketed by a run of 2+ punctuation
characters (`-=*~#`) at start and/or end, with no sentence-ending
punctuation. Reads structurally, no lexicon. Re-read the restatement-only
findings: `pkg/bytecode/decoder.go:1173` ("Read extra map" /
`d.readMapValue()`), `pkg/ir/ir_bridge.lg:59` ("emit a field getter" /
`(defn emit-getter ...)`), several more — all still genuine restatements.
The divider findings (`--- Func roundtrip ---` etc.) are now visible under
their own kind and correctly excluded from `--gate restatement` and from
the "trustworthy" line count below.

### R3 duplicated-comment: 251 -> 135

Fix: exclude a comment whose first content word equals the following Go
declaration's method name (for any function) plus "implements" token, or
equals the method/receiver name outright (for a method specifically).
Verified this removes exactly the "Type implements Value" (20x), "Meta
implements IMeta." (11x), etc. cluster identified in round 1. Remaining
135 still include some idiomatic-but-different repetition, e.g. "Compile
via the stack VM pipeline" (7x in `pkg/ir/spike_rpnvm_test.go`) — a
repeated TEST-SECTION label, not an interface doc comment, so the fix
(scoped narrowly to the Go interface-method convention, per the brief) does
not touch it. This is a second, smaller idiomatic-repetition class I did
not fix, disclosed rather than silently left in.

### R4 comment-density-outlier: 129 (p95=0.889) -> 106 (p95=0.353)

Chose: exclude the leading doc-comment block from BOTH the numerator and
the denominator (interior comments only), over restricting the rule to
unexported declarations. Why: an unexported function can be just as padded
internally as an exported one, so restricting to unexported would miss the
actual failure mode (body padding) while halving the corpus available for
a meaningful percentile; excluding leading docs keeps every declaration
eligible and removes exactly the false-positive class found (well-
documented small public functions) without discarding signal.

Re-read the new top 10 (`pkg/ir/ir_ops.lg`, `pkg/rt/core/core.lg`,
`pkg/rt/core/ir/build.lg`): these are now definitions whose INTERIOR is
35-71% comment lines — plausibly real narration-in-body rather than
godoc. I did not have time in this pass to read each one's actual comment
text line-by-line the way I did for R1's calibration; this is a shallower
read than round 1's, disclosed as such.

## Code-verbosity: catalog, calibration, and the atoms-eliminable measure

Full corpus run (`pkg scripts test`, .lg only): **179 findings, 995 atoms
eliminable**. Breakdown and per-rule read:

| kind | count | read |
|---|---|---|
| `eta-expansion` | 52 | Spot-checked ~10 across `pkg/ir/defop_spec.lg`, `pkg/rt/core/core.lg`, `pkg/rt/core/ir/passes/constfold.lg`. All genuine: `(fn [a b] (+ a b))` -> `+`, etc. — real, mechanically-verifiable eta-reductions. No false positives found. |
| `composable-accessor` (ffirst/fnext/nfirst/nnext) | 38 | Spot-checked several in `pkg/rt/core/core.lg`. Structurally correct in every case read, but a few sit inside the core library itself at a point where the shorter named function may not yet be bootstrapped/available — a real caveat (not a wrongness of the finding, a reason a human might have avoided the rewrite) that a purely structural rule can't see. Disclosed, not fixed. |
| `redundant-eq-true` | 28 | Not individually re-read this pass; mechanism is exact (literal `true` match), no plausible false-positive shape for this pattern. |
| `redundant-not-eq` | 24 | Same — mechanically exact, spot-checked 5, all genuine `(not (= a b))`. |
| `do-single-form` | 12 | Not individually re-read; exact match. |
| `redundant-not-empty` | 8 | Not individually re-read; exact match. |
| `apply-str-interpose` | 7 | Not individually re-read; exact match. |
| `let-identity` | 6 | Not individually re-read; exact match. |
| `nested-if-chain` | 3 | Read all 3. Two are in `test/keyword_cond_switch_test.lg`, which is a fixture FOR an IR pass that optimizes keyword-dispatch cond/if chains — plausibly a deliberately-shaped test input, not verbosity to fix. Structurally correct regardless; flagged with this caveat disclosed. |
| `redundant-if-false-true` | 1 | Only one instance; not independently suspicious. |

Two catalog entries were built, unit-tested, and **dropped** per the
calibration duty ("a rule that fires zero times should be dropped"):
`redundant-if-true-false` (`(if ?c true false)` -> `?c`) and
`let-empty-bindings` (`(let [] ?e)` -> `?e`) — both confirmed working via
their own test fixtures, both zero real hits on `pkg scripts test`
(independently confirmed with a plain-text grep, not just the tool's own
say-so). They are documented as dropped, with rationale, directly in
`scripts/lint-code-rules.edn`.

Overall: for the mechanically-exact rules (anything matching a literal
`true`/`false`/`not`/`empty?`/head symbol with no semantic ambiguity — 8 of
the 10 shipped rules), I have high confidence there is no meaningful
false-positive class; the two rules with real structural judgment calls
(`composable-accessor`'s bootstrap-ordering caveat, `nested-if-chain`'s
test-fixture caveat) are disclosed above rather than silently trusted.

### One test proves the data-driven design

`TestLintCodeVerbosityCatalogIsData` invokes a COPY of `lint.lg` (zero code
changes) next to a fabricated `scripts/lint-code-rules.edn` containing one
rule not in the shipped catalog, and confirms it is found and applied —
proving a new pattern needs only a catalog edit.

### Form-delimited, not line-delimited (a deviation from the paper)

Every code-verbosity finding's `:line`/`:end` are the matched FORM's own
boundaries (from the position-tracking parser built for this), not an
enclosing top-level form or a line heuristic. This is strictly more
precise than SlopCodeBench's own line-delimited flagged regions, and the
per-finding line count is still exactly recoverable for the paper's line
ratio — but it is a real difference from the paper's methodology, noted
here as requested rather than left implicit.

## The two magnitude numbers

Corpus: `pkg scripts test`, 903 files, `.go` + `.lg`.

- **Total LOC** (all lines, including blank, in scanned files): **322,983**
- **Total comment lines** (lines matching `^\s*(//|;)`, a plain-text count
  independent of the tool, to cross-check): **23,778** (7.36% of LOC)

**Comment metric** (R1 + R2-restatement-only [excluding comment-divider] +
R3 + R4 — the four fixed, "trustworthy-per-their-own-definition" objective
rules; 317 findings total): line spans sum to **2,794 lines**
(sum of each finding's own `end - start + 1`, NOT de-duplicated where
findings overlap — a disclosed over-count, likely small since these rules
rarely fire on the same lines).

- 2,794 / 23,778 comment lines = **11.75%** of all comment lines are
  flagged by a trustworthy comment rule.
- 2,794 / 322,983 LOC = **0.87%** of total LOC.

**Code metric** (the 10 shipped code-verbosity rules; 179 findings): line
spans sum to **196 lines**.

- 196 / 322,983 LOC = **0.061%** of total LOC.
- **995 atoms eliminable** total across the corpus — this is the number I
  would use to size the "how much slop is actually removable" question the
  metric can't answer on its own; it is not comparable to a LOC fraction
  because atoms and lines are different units, deliberately (per your
  instruction, atoms are the ranking/impact unit, not a second verbosity
  formula).

Reading these together: comment padding (0.87% of LOC, 11.75% of comment
lines) is measurably larger than code-construct padding (0.061% of LOC) in
THIS corpus by line-fraction — but this is a mature, human-and-AI-mixed
codebase with an existing lint/review culture, not a fresh LLM-generated
corpus, so I would treat these as a floor/baseline rather than
representative of "how bad AI-generated code gets," and recommend
re-running both numbers against a corpus more representative of the
complaint (e.g. a freshly-generated PR) before finalizing a metric weight.

## Concerns carried into this round

- R3's remaining false-positive class (repeated test-section labels, not
  interface docs) is real and unfixed — scoped narrowly per instruction,
  disclosed rather than silently absorbed.
- R4's new top-10 was read more shallowly than R1's original calibration;
  flagging this gap rather than claiming parity of rigor.
- `composable-accessor` and `nested-if-chain` both carry a "structurally
  correct but context-dependent" caveat — not false positives by the
  rule's own definition, but a human reviewer might reasonably not apply
  the suggested rewrite in those specific spots.
- The comment-line and code-line magnitude sums are not de-duplicated
  across overlapping findings; given the low overlap I observed while
  spot-checking, I believe this inflates the true numbers only slightly,
  but I did not compute an exact de-duplicated union in this pass.
