package hash

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func xxh3_hasher() (vm.Value, error) {
	var arg__287 vm.Value
	var v5 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("xxh3", "New").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__287, callErr = rt.InvokeValue(rt.LookupVar("xxh3", "New").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v5, callErr = rt.InvokeValue(rt.LookupVar("hash", "hasher-map").Deref(), []vm.Value{arg__287})
	if callErr != nil {
		return nil, callErr
	}
	return v5, nil
}
func xxh3_hasher_seed(arg0 vm.Value) (vm.Value, error) {
	var arg__296 vm.Value
	var v6 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("xxh3", "NewSeed").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__296, callErr = rt.InvokeValue(rt.LookupVar("xxh3", "NewSeed").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v6, callErr = rt.InvokeValue(rt.LookupVar("hash", "hasher-map").Deref(), []vm.Value{arg__296})
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
	rt.RegisterGoOverrides("hash", map[string]vm.Value{"xxh3-hasher": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 0 {
			return nil, fmt.Errorf("xxh3-hasher: wrong number of arguments %d (expected 0)", len(args))
		}
		return xxh3_hasher()
	}), "xxh3-hasher-seed": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("xxh3-hasher-seed: wrong number of arguments %d (expected 1)", len(args))
		}
		return xxh3_hasher_seed(args[0])
	}),
	})
}
