package ir_passes

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func replace_refs_BANG_(arg0 vm.Value) (vm.Value, error) {
	var v6 vm.Value
	var callErr error
	v6, callErr = rt.InvokeValue(rt.LookupVar("ir", "set-refs!").Deref(), []vm.Value{rt.LookupVar("ir.passes", "*current-fn*").Deref(), rt.LookupVar("ir.passes", "*current-inst*").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v6, nil
}
func replace_aux_BANG_(arg0 vm.Value) (vm.Value, error) {
	var v6 vm.Value
	var callErr error
	v6, callErr = rt.InvokeValue(rt.LookupVar("ir", "set-aux!").Deref(), []vm.Value{rt.LookupVar("ir.passes", "*current-fn*").Deref(), rt.LookupVar("ir.passes", "*current-inst*").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v6, nil
}
func remove_BANG_() (vm.Value, error) {
	var arg__20412 vm.Value
	var v13 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.zipper", "current-block").Deref(), []vm.Value{rt.LookupVar("ir.passes", "*current-zip*").Deref()})
	if callErr != nil {
		return nil, callErr
	}
	arg__20412, callErr = rt.InvokeValue(rt.LookupVar("ir.zipper", "current-block").Deref(), []vm.Value{rt.LookupVar("ir.passes", "*current-zip*").Deref()})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(rt.LookupVar("ir", "remove-inst!").Deref(), []vm.Value{rt.LookupVar("ir.passes", "*current-fn*").Deref(), arg__20412, rt.LookupVar("ir.passes", "*current-inst*").Deref()})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes", map[string]vm.Value{"replace-refs!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("replace-refs!: wrong number of arguments %d (expected 1)", len(args))
		}
		return replace_refs_BANG_(args[0])
	}), "replace-aux!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("replace-aux!: wrong number of arguments %d (expected 1)", len(args))
		}
		return replace_aux_BANG_(args[0])
	}), "remove!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 0 {
			return nil, fmt.Errorf("remove!: wrong number of arguments %d (expected 0)", len(args))
		}
		return remove_BANG_()
	}),
	})
}
