/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package vm

// ExecContext is the single object that resolves an execution's dynamic state —
// dynamic-var bindings now, the structured-concurrency Scope later (see
// docs/design/exec-context-threading.md). It is the *implementer* of
// invocation and of binding resolution: the eval loop, builtins, and goroutine
// spawn all go through an ExecContext rather than looking state up by goroutine
// id.
//
// There is always an ExecContext in play. RootExecContext is the process
// default — its binding stack is the shared global one, so code that never asks
// for isolation behaves exactly as it always has. Per-goroutine isolation is
// opt-in: a spawn hands the child a fresh ExecContext seeded from a snapshot of
// the parent's (ec.Child()), so the two cannot interleave each other's
// bindings, with no goroutine-id lookup and no reuse hazard.
type ExecContext struct {
	bindings *bindingStack
}

// globalBindingStack backs RootExecContext (and the package-level Var binding
// API used by host/lowered code). Sharing it is what makes the root context
// behave like the historical process-global binding store.
var globalBindingStack = newBindingStack()

// RootExecContext is the always-available default context. Frames with no more
// specific context resolve against it.
var RootExecContext = &ExecContext{bindings: globalBindingStack}

// NewExecContext returns a context with a fresh, empty binding stack.
func NewExecContext() *ExecContext {
	return &ExecContext{bindings: newBindingStack()}
}

// Child returns a fresh ExecContext seeded from a snapshot of this one's
// bindings — the explicit-propagation primitive used at goroutine boundaries.
// A nil receiver seeds from the root context.
func (ec *ExecContext) Child() *ExecContext {
	src := ec
	if src == nil {
		src = RootExecContext
	}
	c := NewExecContext()
	c.bindings.installSnapshot(src.bindings.snapshot())
	return c
}

// orRoot normalises a possibly-nil context to the root.
func (ec *ExecContext) orRoot() *ExecContext {
	if ec == nil {
		return RootExecContext
	}
	return ec
}

// --- dynamic-var resolution (ec owns the binding state) ---------------------

// deref returns v's current value in this context: the top dynamic binding if
// any, else v's root.
func (ec *ExecContext) deref(v *Var) Value {
	ec = ec.orRoot()
	if v.isDynamic {
		if val, ok := ec.bindings.current(v); ok {
			return val
		}
	}
	return v.Root()
}

func (ec *ExecContext) pushBinding(v *Var, val Value) {
	v.isDynamic = true
	ec.orRoot().bindings.push(v, val)
}

func (ec *ExecContext) popBinding(v *Var) {
	ec.orRoot().bindings.pop(v)
}

func (ec *ExecContext) hasBinding(v *Var) bool {
	return ec.orRoot().bindings.hasBinding(v)
}

// Exported entry points for runtime builtins (pkg/rt) that resolve dynamic
// vars against an ExecContext handed to them by ec.Invoke.
func (ec *ExecContext) PushBinding(v *Var, val Value) { ec.pushBinding(v, val) }
func (ec *ExecContext) PopBinding(v *Var)             { ec.popBinding(v) }
func (ec *ExecContext) Deref(v *Var) Value            { return ec.deref(v) }

// --- invocation (ec is the implementer) -------------------------------------

// Invoke runs fn with args in this context. Closures propagate the context to
// their child frame; context-aware natives receive it; pure functions are
// invoked unchanged. This is the single dispatch site the eval loop uses.
func (ec *ExecContext) Invoke(fn Fn, args []Value) (Value, error) {
	ec = ec.orRoot()
	switch f := fn.(type) {
	case *Func:
		return f.invokeIn(ec, args)
	case *NativeFn:
		if f.ctxProxy != nil {
			return f.invokeCtx(ec, args)
		}
		return f.Invoke(args)
	default:
		return fn.Invoke(args)
	}
}
