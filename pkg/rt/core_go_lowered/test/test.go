package test

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func testing_contexts_str() (vm.Value, error) {
	var arg__35336 vm.Value
	var arg__35353 vm.Value
	var arg__35354 vm.Value
	var v33 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "interpose").Deref(), []vm.Value{vm.String(" > "), rt.LookupVar("clojure.test", "*testing-contexts*").Deref()})
	if callErr != nil {
		return nil, callErr
	}
	arg__35336, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "interpose").Deref(), []vm.Value{vm.String(" > "), rt.LookupVar("clojure.test", "*testing-contexts*").Deref()})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__35336})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "interpose").Deref(), []vm.Value{vm.String(" > "), rt.LookupVar("clojure.test", "*testing-contexts*").Deref()})
	if callErr != nil {
		return nil, callErr
	}
	arg__35353, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "interpose").Deref(), []vm.Value{vm.String(" > "), rt.LookupVar("clojure.test", "*testing-contexts*").Deref()})
	if callErr != nil {
		return nil, callErr
	}
	arg__35354, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__35353})
	if callErr != nil {
		return nil, callErr
	}
	v33, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__35354})
	if callErr != nil {
		return nil, callErr
	}
	return v33, nil
}
func apply_template(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__35409 vm.Value
	var v8 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zipmap").Deref(), []vm.Value{arg0, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__35409, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zipmap").Deref(), []vm.Value{arg0, arg2})
	if callErr != nil {
		return nil, callErr
	}
	v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.walk", "postwalk-replace").Deref(), []vm.Value{arg__35409, arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v8, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("test", map[string]vm.Value{"testing-contexts-str": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 0 {
			return nil, fmt.Errorf("testing-contexts-str: wrong number of arguments %d (expected 0)", len(args))
		}
		return testing_contexts_str()
	}), "apply-template": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("apply-template: wrong number of arguments %d (expected 3)", len(args))
		}
		return apply_template(args[0], args[1], args[2])
	}),
	})
}
