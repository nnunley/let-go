# SDD progress — native-yamlstar def-closure lowering (Phases A–C)

Plan: docs/superpowers/plans/2026-07-05-native-yamlstar-var-init-lowering.md
Base: uqxtovvv (#394 tip, commit 4876013) — each task `jj new` off previous.
VCS: jj (colocated). Acceptance: capability (InvokeValueEC→0, LookupVar→0,
check-generated + go test green, harness runs); perf observation-only.

- Task A1: complete (change nozrowpn, review clean)
- Task A2: complete (change ymlyuyyu, review clean; 5 test assertions updated for SetVarRoot/ApplyVarMetaV)
- Task B1: complete (change yrnzumyt, review clean)
- Task C1: complete (change wwtwllls, lg-compile env-gated inline). Acceptance PASSED: grammar lowers native, InvokeValueEC=0, LookupVar=0, CachedVarDeref=366.
- gogen fixes (unblock native compile): rwspprqt (reserve Go predeclared idents / len shadowing), nuowutqz (collision-proof __pkg_ import aliasing), yuyxyrts (source reserved names from go/token+go/types.Universe). Yamlstar tree compiles clean; check-generated green; tests green.
- Task C2: pending

## Follow-up (planned, after hardening agent lands)
- Replace hardcoded go-predeclared-identifiers + keyword list in lower_go.lg with a gogen/reserved-ident? builtin sourced from go/token.IsKeyword + go/types.Universe (authoritative, version-tracking). Keep only ec/callErr as lowering-internal reserved.
