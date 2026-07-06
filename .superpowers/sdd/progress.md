# SDD progress — native-yamlstar def-closure lowering (Phases A–C)

Plan: docs/superpowers/plans/2026-07-05-native-yamlstar-var-init-lowering.md
Base: uqxtovvv (#394 tip, commit 4876013) — each task `jj new` off previous.
VCS: jj (colocated). Acceptance: capability (InvokeValueEC→0, LookupVar→0,
check-generated + go test green, harness runs); perf observation-only.

- Task A1: complete (change nozrowpn, review clean)
- Task A2: complete (change ymlyuyyu, review clean; 5 test assertions updated for SetVarRoot/ApplyVarMetaV)
- Task B1: complete (change yrnzumyt, review clean)
- Task C1: acceptance PASSED (change wwtwllls; InvokeValueEC=0, LookupVar=0, RegisterGoVarInits=1, CachedVarDeref=440) — grammar lowered native. BLOCKED on pre-existing gogen bug: user fn `len` shadows Go builtin. Fix in progress (reserve Go predeclared idents).
- Task C2: pending

## Follow-up (planned, after hardening agent lands)
- Replace hardcoded go-predeclared-identifiers + keyword list in lower_go.lg with a gogen/reserved-ident? builtin sourced from go/token.IsKeyword + go/types.Universe (authoritative, version-tracking). Keep only ec/callErr as lowering-internal reserved.
