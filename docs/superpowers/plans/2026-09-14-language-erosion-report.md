---
status: active
last-verified: 2026-09-14
---

# Language-specific erosion reporting Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Report `.lg` and Go erosion distinctly, while withholding combined debt and PR deltas if requested Go analysis fails.

**Architecture:** `quality/measure` owns completeness and the EDN result contract; `quality.report` only renders that contract. A missing Go analyzer sets the combined erosion debt input and top-level debt to nil, but retains `.lg` erosion. `compare-results` rejects incomplete inputs before subtracting. Successful analysis retains the existing combined debt arithmetic.

**Tech Stack:** let-go `.lg`, Go 1.26.5 test runner, Git fixture corpus, EDN.

**Spec:** `docs/superpowers/specs/2026-09-14-language-erosion-report-design.md`.

---

## Chunk 1: Two language scores and fail-closed report

### Task 1: Characterize result states before changing production

**Files:**
- Modify: `test/quality_cli_test.lg`
- Modify: `scripts/quality.lg`
- Modify: `scripts/quality/report.lg`
- Create: `test/fixtures/quality-go-empty/empty.go` (isolated tracked fixture)

- [ ] Create `test/fixtures/quality-go-empty/empty.go` with only `package qualitygoempty`; stage only it with `rtk proxy git add -f test/fixtures/quality-go-empty/empty.go` before discovery tests. Keep it outside `fixtures/quality` so existing fixture scores stay unchanged. It is test data, not a dependency.
- [ ] Capture the pre-edit successful mixed-corpus numeric baseline using existing `cli-produces-deterministic-report-over-fixtures` arguments (`--ci-seconds 300 --go-cover fixtures/quality/cover.out --since 2000-01-01 --source-root fixtures/quality fixtures/quality`). The focused baseline run already wrote `/tmp/quality-cli-test.edn`: top-level `:debt` is `0.7563004013065922` and `[:terms :erosion]` is `0.6594225937101561` under Go 1.26.5 with `go run`. Re-run the baseline test before production edits if this artifact has been replaced; record the exact two numbers from its EDN. In the completed test assert both values with a tolerance of `1e-9`, not only self-consistency.
- [ ] Add an unavailable-Go test using tracked `fixtures/quality` (`.lg` and `.go`). Save `(os/getenv "QUALITY_GO_CALLABLES")`, set `"/nonexistent/go-callables"`, call `q/main ["--since" "2000-01-01" "--edn" edn-path "--source-root" fixtures fixtures]`, and restore with `(os/setenv "QUALITY_GO_CALLABLES" (or saved ""))` in `finally`. Empty and absent are behaviorally equivalent to `go-analysis`; do not claim byte-identical process environment if initially absent. Use unique `(str (os/temp-dir) "/quality-go-unavailable.edn")`. Assert:

```clojure
(is (= :incomplete-go-analysis (:score-status r)))
(is (nil? (:debt r)))
(is (nil? (get-in r [:terms :erosion])))
(is (number? (get-in r [:erosion :lg])))
(is (nil? (get-in r [:erosion :go])))
(is (= "unavailable" (get-in r [:erosion :go-source])))
(is (nil? (get-in r [:erosion :callables])))
(is (nil? (get-in r [:erosion :high])))
(is (not (contains? (:erosion r) :value)))
(is (str/includes? out "debt UNAVAILABLE"))
(is (str/includes? out "go=UNAVAILABLE"))
(is (str/includes? out "Go callables excluded"))
```

  Complete test skeleton (keep the assertions above inside its inner `let`):

```clojure
(deftest unavailable-go-analysis-never-lowers-debt
  (let [saved (os/getenv "QUALITY_GO_CALLABLES")
        edn-path (str (os/temp-dir) "/quality-go-unavailable.edn")]
    (try
      (os/setenv "QUALITY_GO_CALLABLES" "/nonexistent/go-callables")
      (let [out (with-out-str
                  (q/main ["--since" "2000-01-01" "--edn" edn-path
                           "--source-root" fixtures fixtures]))
            r (read-string (slurp edn-path))]
        ;; Place the exact assertions above here.
        (is (= :incomplete-go-analysis (:score-status r)))
        (is (nil? (:debt r)))
        (is (str/includes? out "go=UNAVAILABLE")))
      (finally (os/setenv "QUALITY_GO_CALLABLES" (or saved ""))))))
```

- [ ] Extend the successful mixed-corpus test: assert `:score-status :complete`, numeric `.lg` and Go erosion, and `(:debt r) = (:debt (quality.score/corpus-score (:terms r)))` (add the score require) to lock unchanged arithmetic without a brittle literal. Assert `erosion lg=`/`go=` text and no old unlabeled aggregate summary line. Replace the old `#{"go run" "none"}` source assertion with a success check permitting a configured binary path.
- [ ] Add a `.lg`-only request for tracked `fixtures/quality/cli_src.lg`: assert `:go-source "none"`, `:go nil`, complete status, numeric debt, `go=NOT REQUESTED`. Add Go-only request for tracked `fixtures/quality/dup_a.go`: assert `lg=N/A`, `:lg nil`, Go numeric, complete. Add staged `fixtures/quality-go-empty/empty.go` request: assert `lg=N/A go=N/A`, both language scores nil, non-unavailable `:go-source`, complete. Each request uses a unique temp EDN basename.

  Use this helper and exact inputs; add the listed assertions for each `r`/`out` pair:

```clojure
(defn- fixture-report [edn-name source-root path]
  (let [edn-path (str (os/temp-dir) "/" edn-name)
        out (with-out-str
              (q/main ["--since" "2000-01-01" "--edn" edn-path
                       "--source-root" source-root path]))]
    {:out out :r (read-string (slurp edn-path))}))

(deftest language-erosion-empty-boundaries
  (let [{lg-out :out lg-r :r}
        (fixture-report "quality-lg-only.edn" fixtures
                        (str fixtures "/cli_src.lg"))
        {go-out :out go-r :r}
        (fixture-report "quality-go-only.edn" fixtures
                        (str fixtures "/dup_a.go"))
        {empty-out :out empty-r :r}
        (fixture-report "quality-go-empty.edn" "fixtures/quality-go-empty"
                        "fixtures/quality-go-empty/empty.go")]
    (is (= "none" (get-in lg-r [:erosion :go-source])))
    (is (nil? (get-in lg-r [:erosion :go])))
    (is (= :complete (:score-status lg-r)))
    (is (number? (:debt lg-r)))
    (is (str/includes? lg-out "go=NOT REQUESTED"))
    (is (nil? (get-in go-r [:erosion :lg])))
    (is (number? (get-in go-r [:erosion :go])))
    (is (= :complete (:score-status go-r)))
    (is (str/includes? go-out "lg=N/A"))
    (is (nil? (get-in empty-r [:erosion :lg])))
    (is (nil? (get-in empty-r [:erosion :go])))
    (is (not= "unavailable" (get-in empty-r [:erosion :go-source])))
    (is (= :complete (:score-status empty-r)))
    (is (str/includes? empty-out "lg=N/A go=N/A"))))
```
- [ ] Run focused tests and capture the **expected red** assertion failures on current production:

```sh
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" TMPDIR=/tmp go test ./test -run '^TestRunner/quality_cli_test.lg$' -count=1
```

Expected: new assertions FAIL on missing status, numeric debt after unavailable Go, and old aggregate erosion text/EDN shape. A Go-cache sandbox denial is not a red test; retry the identical command with approval.

- [ ] Implement minimal `measure` contract. Compute `go-unavailable?` from `(:source go-analysis)`, not from `:go` being nil (zero functions is valid). For unavailable Go, set `corpus-terms :erosion` to nil before `score/corpus-score`; then overwrite its renormalized `:debt` with nil and add `:score-status :incomplete-go-analysis`. For other states use `:score-status :complete`. Set the language `:erosion :go` to nil on unavailable or not-requested, and remove `:erosion :value`; leave `:terms :erosion` numeric on complete reports. Set combined `:callables/:high` nil only when Go analysis failed.

  Core binding shape (adapt names to existing `let` without recomputing analyses):

```clojure
go-unavailable? (= "unavailable" (:source go-analysis))
corpus-terms {:erosion (when-not go-unavailable? (terms/erosion callables))
              ;; all other existing terms unchanged
              }
raw-corpus (score/corpus-score corpus-terms)
corpus (assoc raw-corpus
              :score-status (if go-unavailable? :incomplete-go-analysis :complete)
              :debt (when-not go-unavailable? (:debt raw-corpus)))
```

  In the result map, retain `:erosion :lg` and `:go-source`; set `:go` to nil if source is `"unavailable"` or `"none"`, remove `:value`, and set `:callables/:high` to nil only if `go-unavailable?`.
- [ ] Update `report/debt-block`: print `debt UNAVAILABLE` for incomplete status; render `erosion lg=<n|N/A> go=<n|N/A|UNAVAILABLE|NOT REQUESTED>` and an explicit exclusion warning on failure. Use `:go-source` to distinguish failed, absent, and successful empty Go analysis. Do not render the old unlabeled aggregate `erosion <n>` summary. Keep score weights and non-erosion terms unchanged.

  Required rendering shape:

```clojure
(str "debt " (if (= :incomplete-go-analysis score-status)
               "UNAVAILABLE" (n3 total)) "   (0 is perfect, 1 is worst)")
(str "  erosion lg=" (if (nil? (:lg erosion)) "N/A" (n3 (:lg erosion)))
     " go=" (case (:go-source erosion)
               "unavailable" "UNAVAILABLE"
               "none" "NOT REQUESTED"
               (if (nil? (:go erosion)) "N/A" (n3 (:go erosion)))))
```

  Add `"Go callables excluded; combined debt unavailable"` only on incomplete status. Any combined callable count displayed on that path must be labeled `.lg-only` or omitted; the EDN combined counts are nil.
- [ ] Re-run the focused test; expect green. Run `rtk proxy git diff --check`; inspect the exact diff. Commit only these source/test files (and any necessary tracked zero-function fixture) with a conventional message explaining why missing analysis must not look like improvement. Do not touch `Makefile`.

## Chunk 2: Refuse incomplete PR deltas and verify

### Task 2: Delta guard

**Files:**
- Modify: `test/quality_delta_test.lg`
- Modify: `scripts/quality.lg`

- [ ] Add pure `compare-results` tests for incomplete base and head. Use `{:debt nil :score-status :incomplete-go-analysis :terms {:erosion nil} :files []}` as the incomplete map and `{:debt 0.3 :score-status :complete :terms {:erosion 0.4} :files []}` as complete. In `try/catch Throwable`, assert `(get (ex-data error) :side)` is `:base` when the first map is incomplete and `:head` when the second is incomplete. Assert two complete maps still yield numeric `:debt-delta`. Run the **red** command before production edit:

```sh
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" TMPDIR=/tmp go test ./test -run '^TestRunner/quality_delta_test.lg$' -count=1
```

Expected: current code attempts nil subtraction or lacks side-specific `ex-data`, so assertions FAIL.
- [ ] At the start of `compare-results`, reject a base or head whose `:score-status` is `:incomplete-go-analysis` or whose debt is nil, before term/file subtraction. Keep complete deltas unchanged and include base/head side in `ex-data`.

```clojure
(when (or (= :incomplete-go-analysis (:score-status base-res))
          (nil? (:debt base-res)))
  (throw (ex-info "Cannot compare incomplete quality report"
                  {:side :base :score-status (:score-status base-res)})))
(when (or (= :incomplete-go-analysis (:score-status head-res))
          (nil? (:debt head-res)))
  (throw (ex-info "Cannot compare incomplete quality report"
                  {:side :head :score-status (:score-status head-res)})))
```
- [ ] Re-run focused and full checks. Each must exit 0, and diff check prints nothing. Do not run jj-dependent tests. If sandbox denies Go cache or local listeners, retry the identical command with approval rather than treating denial as a code failure.

```sh
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" TMPDIR=/tmp go test ./test -run '^TestRunner/(quality_cli_test.lg|quality_score_test.lg|quality_delta_test.lg)$' -count=1
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" TMPDIR=/tmp go test -p 1 -short ./...
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" go build ./...
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" go vet ./...
rtk proxy python3 scripts/docs_frontmatter_hook.py --check docs/superpowers/specs/2026-09-14-language-erosion-report-design.md docs/superpowers/plans/2026-09-14-language-erosion-report.md
rtk proxy git diff --check
```
- [ ] Confirm `go-analysis` failure text and EDN do not contain a numeric top-level debt or combined erosion score; confirm successful mixed-corpus debt is unchanged from pre-edit baseline. Commit only the delta source/test files. Report exact verification and branch status. Do not push without a separate request.
