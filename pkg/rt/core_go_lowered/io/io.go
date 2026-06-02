package io

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func slurp_lines(arg0 vm.Value) (vm.Value, error) {
	var arg__305 vm.Value
	var v6 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("io", "read-lines").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__305, callErr = rt.InvokeValue(rt.LookupVar("io", "read-lines").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__305})
	if callErr != nil {
		return nil, callErr
	}
	return v6, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("io", map[string]vm.Value{"slurp-lines": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("slurp-lines: wrong number of arguments %d (expected 1)", len(args))
		}
		return slurp_lines(args[0])
	}),
	})
}
