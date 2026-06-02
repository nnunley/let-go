/*
 * Copyright (c) 2026 let-go contributors
 * SPDX-License-Identifier: MIT
 */

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

// RegisterGoOverrides queues a set of name → NativeFn bindings. If the
// target namespace already exists in the registry, the defs are applied
// immediately (the host has already finished bundle replay); otherwise
// they sit in pendingGoOverrides until ApplyGoOverrides drains them.
//
// Init-order hazard this guards against: pkg/compiler's init() replays
// the core bundle and calls ApplyGoOverrides(coreNS) BEFORE the blank-
// imported lowered packages in main get to run their own init(). Without
// the immediate-apply path here, every override registered after that
// point would sit in the queue forever — the queue would be drained
// once when empty, then filled and never re-drained. Symptom is the
// dispatch counters reporting zero and `(reduce + (range 1000))` going
// through bytecode dispatch even with -tags gogen_ir + all the blank
// imports wired up.
func RegisterGoOverrides(nsName string, defs map[string]vm.Value) {
	if len(defs) == 0 {
		return
	}
	if ns := LookupNS(nsName); ns != nil {
		for name, fn := range defs {
			ns.Def(name, fn)
		}
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
