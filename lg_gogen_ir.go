//go:build gogen_ir

// gogen_ir build tag — pulls in the generated Go-native IR-stack packages
// from pkg/rt/core_go_lowered so each package's init() runs and registers
// its native overrides via rt.RegisterGoOverrides. The blank imports must
// live in package main (or any package higher than pkg/rt in the DAG)
// because the generated packages import pkg/rt themselves; importing them
// from inside pkg/rt would create a cycle.
//
// When this tag is set, ApplyGoOverrides drains each NS's pending map
// after the bytecode bundle has replayed, clobbering the Lisp vars with
// NativeFn wrappers that call directly into compiled Go. Untagged builds
// are completely unaffected — no overrides registered, hook is no-op.
//
// Directory layout mirrors the dotted ns: ir.passes.dce →
// pkg/rt/core_go_lowered/ir/passes/dce (package dce).

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
