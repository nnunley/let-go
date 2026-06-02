package ir_passes_dce

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func removable_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__21131 vm.Value
	var and__x vm.Value
	var nid vm.Value
	var f vm.Value
	var arg__21141 vm.Value
	var arg__21154 vm.Value
	var arg__21156 vm.Value
	var v28 vm.Value
	var v31 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__21131, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.dce", "pure-ops").Deref(), []vm.Value{arg__21131})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		nid = arg0
		f = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21141, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__21141, nid})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21154, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21156, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__21154, nid})
	if callErr != nil {
		return nil, callErr
	}
	v28, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{arg__21156})
	if callErr != nil {
		return nil, callErr
	}
	v31 = v28
	goto b3
b2:
	;
	v31 = and__x
	goto b3
b3:
	;
	return v31, nil
}
func one_pass(arg0 vm.Value) (vm.Value, error) {
	var removed vm.Value
	var arg__21174 vm.Value
	var doseq_seq__21159 vm.Value
	var doseq_loop__21160 vm.Value
	var f vm.Value
	var bid vm.Value
	var arg__21190 vm.Value
	var arg__21205 vm.Value
	var arg__21206 vm.Value
	var doseq_seq__21161 vm.Value
	var v115 vm.Value
	var doseq_loop__21162 vm.Value
	var nid vm.Value
	var v80 vm.Value
	var v105 vm.Value
	var v89 vm.Value
	var v92 vm.Value
	var callErr error
	removed, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.Boolean(false)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__21174, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__21159, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__21174})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21160 = doseq_seq__21159
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__21160) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__21160})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21190, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__21190})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21205, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21206, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__21205})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__21161, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__21206})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21162 = doseq_seq__21161
	goto b5
b3:
	;
	goto b4
b4:
	;
	v115, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{removed})
	if callErr != nil {
		return nil, callErr
	}
	return v115, nil
b5:
	;
	if vm.IsTruthy(doseq_loop__21162) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__21162})
	if callErr != nil {
		return nil, callErr
	}
	v80, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.dce", "removable?").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v80) {
		goto b9
	} else {
		goto b10
	}
b7:
	;
	goto b8
b8:
	;
	v105, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__21160})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21160 = v105
	goto b1
b9:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "remove-inst!").Deref(), []vm.Value{f, bid, nid})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{removed, vm.Boolean(true)})
	if callErr != nil {
		return nil, callErr
	}
	v89, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__21162})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21162 = v89
	goto b5
b10:
	;
	v92, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__21162})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21162 = v92
	goto b5
}
func dce(arg0 vm.Value) (vm.Value, error) {
	var f vm.Value
	var v7 vm.Value
	var callErr error
	f = arg0
	goto b1
b1:
	;
	v7, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.dce", "one-pass").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v7) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	goto b1
b3:
	;
	goto b4
b4:
	;
	return f, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.dce", map[string]vm.Value{"removable?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("removable?: wrong number of arguments %d (expected 2)", len(args))
		}
		return removable_QMARK_(args[0], args[1])
	}), "one-pass": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("one-pass: wrong number of arguments %d (expected 1)", len(args))
		}
		return one_pass(args[0])
	}), "dce": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("dce: wrong number of arguments %d (expected 1)", len(args))
		}
		return dce(args[0])
	}),
	})
}
