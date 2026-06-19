package bytecode_test

import (
	"bytes"
	"testing"

	"github.com/nooga/let-go/pkg/bytecode"
	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

// TestNoInitDefLGBParity guards def/declare "promise" semantics across the LGB
// decode path. The compiler emits only OP_LOAD_CONST for a no-init (def x), so
// the var that ends up in the namespace is the one the decode resolver interns
// for the var constant — never written by bytecode. That interned var must be
// UNBOUND, matching the source-compiled path (LookupOrAdd). If the resolver's
// stub installs a NIL root instead, (bound? (var x)) wrongly reports true for
// precompiled/LGB-loaded code, diverging from freshly compiled code.
func TestNoInitDefLGBParity(t *testing.T) {
	const nsName = "no-init-lgb-decode-parity"
	const varName = "x"

	consts := vm.NewConsts()
	chunk := vm.NewCodeChunk(consts)
	idx := consts.Intern(vm.NewVar(nil, nsName, varName))
	chunk.Append(vm.OP_LOAD_CONST)
	chunk.Append32(idx)
	chunk.Append(vm.OP_RETURN)

	var buf bytes.Buffer
	if err := bytecode.EncodeCompilation(&buf, consts, chunk); err != nil {
		t.Fatalf("encode failed: %v", err)
	}

	// Mirror the bundle-decode resolver in pkg/compiler/eval.go: minimal
	// namespace, look up locally, otherwise intern a stub for the var ref.
	resolve := func(ns, name string) *vm.Var {
		n := rt.DefNSBare(ns)
		if v := n.LookupLocal(vm.Symbol(name)); v != nil {
			return v
		}
		return n.DefStub(name)
	}
	if _, err := bytecode.DecodeToExecUnit(&buf, resolve); err != nil {
		t.Fatalf("decode failed: %v", err)
	}

	resolved := rt.NS(nsName).LookupLocal(vm.Symbol(varName))
	if resolved == nil {
		t.Fatalf("var %q was not interned during decode", varName)
	}
	if resolved.IsBound() {
		t.Fatalf("decoded var constant is bound before bytecode execution")
	}
}
