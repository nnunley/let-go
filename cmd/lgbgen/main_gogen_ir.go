//go:build bootstrap && gogen_ir

// Meta-circular bootstrap: when lgbgen is built with both -tags bootstrap
// AND -tags gogen_ir, link in the generated Go-native IR-stack packages.
// Each package's init() registers its NativeFn overrides via
// rt.RegisterGoOverrides; the resolver's post-load hook drains those
// onto the host VM as namespaces load. The bytecode bundle still drives
// initial namespace replay (some defns can't be lowered to Go), so the
// generated overrides clobber the Lisp vars after replay — passes that
// lowered cleanly run native, the rest fall back to bytecode.
//
// Without this file, `go run -tags 'bootstrap gogen_ir' ./cmd/lgbgen`
// would silently not pick up the overrides because cmd/lgbgen is its
// own package main — the repo-root lg_gogen_ir.go applies to the lg
// binary only.

package main

import (
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/graph"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/build"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/dominance"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/dump"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/lower"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/lower_go"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/constfold"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/cse"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/dce"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/infer_arg_types"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/licm"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/mutability"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/pipeline"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/trace"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/typeinfer"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/validate"
	_ "github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/zipper"
)
