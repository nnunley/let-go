---
status: active
last-verified: 2026-09-14
---

# Git Patch Parser Erosion Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Reduce `quality.diff/parse-patch-hunks` and corpus erosion while preserving every Git-diff record and error contract.

**Architecture:** Separate section scanning from name-status reconciliation. A line classifier owns patch metadata; a reconciliation helper owns ordered status/path checks and the special two-section type change. The public fn composes them without changing its signature.

**Tech Stack:** let-go `.lg`, Git patch format, `quality.metrics`, Go 1.26.5 test runner.

---

## Chunk 1: Characterize and extract the two parser phases

### Task 1: Preserve behavior and reduce measured complexity

**Files:**
- Modify: `test/quality_diff_test.lg` (characterization and score-bound tests)
- Modify: `scripts/quality/diff.lg` (`parse-patch-hunks` and private helpers)
- Spec: `docs/superpowers/specs/2026-09-14-patch-hunk-erosion-design.md`

- [ ] In `test/quality_diff_test.lg`, add `[quality.metrics :as metrics]` to the ns requirements and a structural red test:

```clojure
(deftest patch-hunk-parser-stays-below-erosion-threshold
  (let [defs (:definitions (metrics/file-metrics (slurp "../scripts/quality/diff.lg")))
        parser (first (filter #(= 'parse-patch-hunks (:name %)) defs))]
    (is (some? parser))
    (is (<= (:cc parser) 10))))
```

  Run `rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" go test ./test -run '^TestRunner/quality_diff_test.lg$' -count=1` from the dedicated worktree. Expected red: parser CC 50 exceeds 10. The full behavioral suite was green at branch start.

- [ ] Add focused characterization cases to the same file. Existing tests already cover quoted paths, zero-count anchors, rename, type-change pairing, symlinks, and dirty gitlinks. Add the following cases (adjust only formatting, not expected data):

```clojure
(defn- patch-error-data [patch paths]
  (try (d/parse-patch-hunks patch paths)
       nil
       (catch Throwable e (ex-data e))))

(deftest patch-errors-retain-diagnostic-data
  (let [a (first (d/parse-name-status-z "M\u0000a.lg\u0000"))
        two [a (first (d/parse-name-status-z "M\u0000b.lg\u0000"))]
        header "diff --git a/a.lg b/a.lg"
        wrong "diff --git a/wrong.lg b/wrong.lg"]
    (is (= {:sections 1 :statuses 2}
           (patch-error-data (str header "\n") two)))
    (is (= {:status a :header wrong}
           (patch-error-data (str wrong "\n") [a])))
    (is (= {:line "@@ malformed @@"}
           (patch-error-data (str header "\n@@ malformed @@\n") [a])))))

(deftest patch-metadata-flags-retain-file-types
  (let [paths (d/parse-name-status-z
               "M\u0000blob.bin\u0000M\u0000module\u0000M\u0000link\u0000")
        oid "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
        patch (str "diff --git a/blob.bin b/blob.bin\n"
                   "index 111..222 100644\nGIT binary patch\n"
                   "diff --git a/module b/module\n"
                   "index 111..222 160000\n"
                   "-Subproject commit " oid "\n+Subproject commit " oid "-dirty\n"
                   "diff --git a/link b/link\n"
                   "index 111..222 120000\n@@ -1 +1 @@\n-old\n+new\n")
        rows (d/parse-patch-hunks patch paths)]
    (is (= true (:binary? (first rows))))
    (is (= true (:gitlink? (second rows))))
    (is (= true (:symlink? (nth rows 2))))
    (is (= [{:start 1 :count 1}] (:new-spans (nth rows 2))))))
```

  Run the same focused test command before extraction. Expected: only the intentional CC-bound assertion fails; all characterization assertions pass. If they expose a real discrepancy, stop and distinguish existing behavior from a separate correction.

- [ ] Capture the **pre-refactor** quality report now, after test additions but before `scripts/quality/diff.lg` changes. Run `rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" QUALITY_TOP=20 make quality` as its own command and require exit 0. Record the printed `debt`, `erosion`, and `parse-patch-hunks` CC/mass lines in the task report. Do not pipe `make` to `rg` or treat an exit from a filter as proof the report succeeded.

- [ ] Extract the scanner from the first `let` binding of `parse-patch-hunks` into private `scan-patch-sections [patch]`. It still splits on `#"\n"`, starts a section only on `diff --git `, accumulates `:hunks`, and appends the final section. Extract a private line-classification/metadata helper so the scanner is not itself a long `cond`; preserve the existing regexes and exact flag keys (`:binary?`, `:blob?`, `:symlink?`, `:gitlink?`, `:gitlink-old?`, `:gitlink-new?`, `:deleted-file?`, `:new-file?`). Keep `hunk-coordinates` unchanged. **Define functions in file order `patch-mode-kind` → `advance-patch-section` → `scan-patch-sections`**; the scanner skeleton is displayed first only to show its interface, not as the source insertion order.

```clojure
(defn- scan-patch-sections [patch]
  (loop [lines (str/split patch #"\n") current nil out []]
    (if (empty? lines)
      (if current (conj out current) out)
      (let [line (first lines)]
        (if (str/starts-with? line "diff --git ")
          (recur (rest lines) {:header line :hunks []}
                 (if current (conj out current) out))
          (recur (rest lines) (if current (advance-patch-section current line) nil)
                 out))))))
```

  Use the following classifier and step; the regex accepts the exact four
  modes and two metadata formats accepted by the original branches.

```clojure
(defn- patch-mode-kind [line]
  (let [[_ mode]
        (re-matches
          #"^(?:(?:new file mode|deleted file mode|old mode|new mode) |index [0-9a-fA-F]+\.\.[0-9a-fA-F]+ )(100644|100755|120000|160000)$"
          line)]
    (case mode
      "100644" :blob "100755" :blob
      "120000" :symlink "160000" :gitlink
      nil)))

(defn- advance-patch-section [section line]
  (cond
    (str/starts-with? line "@@ ")
    (update section :hunks conj (hunk-coordinates line))

    (re-matches #"^[+-]Subproject commit [0-9a-fA-F]{40,64}(?:-dirty)?$" line)
    (assoc section (if (str/starts-with? line "-") :gitlink-old? :gitlink-new?) true)

    (or (str/starts-with? line "Binary files ") (= line "GIT binary patch"))
    (assoc section :binary? true)

    :else
    (let [kind (patch-mode-kind line)
          key (get {:blob :blob? :symlink :symlink? :gitlink :gitlink?} kind)]
      (cond-> (if key (assoc section key true) section)
        (str/starts-with? line "deleted file mode ") (assoc :deleted-file? true)
        (str/starts-with? line "new file mode ") (assoc :new-file? true)))))
```

  Lines outside sections remain ignored. In particular, a regular blob's
  `Subproject commit` text does not set the output `:gitlink?` flag.

- [ ] Extract the status loop into the following domain helpers, moving the
  current row logic verbatim. Keep `header-paths` and `hunk-coordinates`
  unchanged. The only public-fn body change is the final composition.

```clojure
(defn- type-pair? [path section next-section]
  (and (= :type-changed (:status path))
       (:deleted-file? section) (:new-file? next-section)
       (= (:header section) (:header next-section))))

(defn- section-row [path section]
  (let [row (assoc path
                   :old-spans (if (= :type-changed (:status path)) []
                                  (vec (map first (:hunks section))))
                   :new-spans (if (= :type-changed (:status path)) []
                                  (vec (map second (:hunks section)))))]
    (cond-> row
      (:binary? section) (assoc :binary? true)
      (or (:gitlink? section)
          (and (not (:blob? section))
               (:gitlink-old? section) (:gitlink-new? section)))
      (assoc :gitlink? true)
      (:symlink? section) (assoc :symlink? true))))

(defn- reconcile-patch-sections [sections paths]
  (loop [remaining sections statuses paths out []]
    (if (empty? statuses)
      (if (empty? remaining) out
          (throw (ex-info "Git patch section count differs from name-status"
                          {:sections (count sections) :statuses (count paths)})))
      (let [path (first statuses)
            section (first remaining)
            next-section (second remaining)]
        (when-not section
          (throw (ex-info "Git patch section count differs from name-status"
                          {:sections (count sections) :statuses (count paths)})))
        (header-paths (:header section) path)
        (let [pair? (type-pair? path section next-section)]
          (when pair? (header-paths (:header next-section) path))
          (recur (drop (if pair? 2 1) remaining)
                 (rest statuses)
                 (conj out (section-row path section))))))))

(defn parse-patch-hunks [patch paths]
  (reconcile-patch-sections (scan-patch-sections patch) paths))
```

- [ ] Run these exact verification commands from the dedicated worktree. Each
  must exit 0; `diff --check` must print nothing.

```sh
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" go test ./test -run '^TestRunner/(quality_diff_test.lg|quality_delta_test.lg|quality_touched_test.lg)$' -count=1
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" go test -short ./...
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" go build ./...
rtk proxy env -u GOROOT PATH="/Users/ndn/.local/share/mise/installs/go/1.26.5/bin:$PATH" go vet ./...
rtk git diff --check
```

- [ ] Run the same standalone `make quality` command after extraction. Require
  exit 0, parser CC <=10, corpus erosion below the recorded pre-refactor
  value (about 0.414), composite debt no higher than the pre-refactor value
  (about 0.356), and no new helper above CC 10 that merely inherits the
  parser's branching. Record both sets of printed lines in the task report.
  If the goals fail, revise the decomposition and rerun tests/score.

- [ ] Self-review behavior and score evidence, then commit only `scripts/quality/diff.lg` and `test/quality_diff_test.lg` on the separate erosion branch. Do not push or modify the published #871 branch.
