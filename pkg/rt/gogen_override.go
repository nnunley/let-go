package rt

import (
	"github.com/nooga/let-go/pkg/vm"
)

// Go-native overrides for generated IR-stack namespaces.
//
// A Go-lowered package's init() registers its defns here. When the host
// resolver finishes loading a namespace (bytecode replay or source compile),
// it drains the pending overrides for that namespace, clobbering any vars
// the bytecode/source path produced with native Go implementations.
//
// Without any registrations (no gogen_ir build tag, no generated packages
// imported), the maps stay empty and the hook is a single map lookup.

var pendingGoOverrides = map[string]map[string]vm.Value{}

// RegisterGoOverrides queues a set of name → NativeFn bindings to apply once
// the namespace has been loaded. Safe to call from init() before the host
// runtime has materialized the namespace.
func RegisterGoOverrides(nsName string, defs map[string]vm.Value) {
	if defs == nil || len(defs) == 0 {
		return
	}
	existing := pendingGoOverrides[nsName]
	if existing == nil {
		existing = make(map[string]vm.Value, len(defs))
		pendingGoOverrides[nsName] = existing
	}
	for k, v := range defs {
		existing[k] = v
	}
}

// ApplyGoOverrides drains any pending overrides for ns and Defs them onto
// the namespace, replacing whatever bytecode or source replay produced.
// No-op if nothing is pending.
func ApplyGoOverrides(ns *vm.Namespace) {
	if ns == nil {
		return
	}
	defs := pendingGoOverrides[ns.Name()]
	if defs == nil {
		return
	}
	for name, fn := range defs {
		ns.Def(name, fn)
	}
	delete(pendingGoOverrides, ns.Name())
}
