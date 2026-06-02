package ir_passes_infer_arg_types

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func join_types(arg0 vm.Value, arg1 vm.Value) vm.Value {
	var v8 bool
	var b vm.Value
	var a vm.Value
	var v16 bool
	var v48 vm.Value
	var v23 bool
	var v44 vm.Value
	var v40 vm.Value
	var v36 vm.Value
	v8 = arg0 == vm.Keyword("unknown")
	if v8 {
		b = arg1
		goto b1
	} else {
		a = arg0
		b = arg1
		goto b2
	}
b1:
	;
	v48 = b
	goto b3
b2:
	;
	v16 = b == vm.Keyword("unknown")
	if v16 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v48
b4:
	;
	v44 = a
	goto b6
b5:
	;
	v23 = a == b
	if v23 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v48 = v44
	goto b3
b7:
	;
	v40 = a
	goto b9
b8:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v44 = v40
	goto b6
b10:
	;
	v36 = vm.Keyword("unknown")
	goto b12
b11:
	;
	v36 = vm.NIL
	goto b12
b12:
	;
	v40 = v36
	goto b9
}
func constraint_from_user(arg0 vm.Value) (vm.Value, error) {
	var v7 vm.Value
	var v22 vm.Value
	var v19 vm.Value
	var callErr error
	v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.passes.infer-arg-types", "int-constraint-ops").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v7) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v22 = vm.Keyword("int")
	goto b3
b2:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v22, nil
b4:
	;
	v19 = vm.Keyword("unknown")
	goto b6
b5:
	;
	v19 = vm.NIL
	goto b6
b6:
	;
	v22 = v19
	goto b3
}
func infer_one_load_arg_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var bs vm.Value
	var joined vm.Value
	var v21 vm.Value
	var f vm.Value
	var nid vm.Value
	var arg__21309 vm.Value
	var v52 vm.Value
	var v70 vm.Value
	var arg__21322 vm.Value
	var v59 vm.Value
	var v63 vm.Value
	var callErr error
	bs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	joined, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.Keyword("unknown")})
	if callErr != nil {
		return nil, callErr
	}
	v21, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v21) {
		goto b1
	} else {
		f = arg0
		nid = arg2
		goto b2
	}
b1:
	;
	v70 = vm.NIL
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-for-each").Deref(), []vm.Value{bs, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var user_op vm.Value
		var c vm.Value
		var v10 vm.Value
		var callErr error
		user_op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		c, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.infer-arg-types", "constraint-from-user").Deref(), []vm.Value{user_op})
		if callErr != nil {
			return nil, callErr
		}
		v10, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{joined, rt.LookupVar("ir.passes.infer-arg-types", "join-types").Deref(), c})
		if callErr != nil {
			return nil, callErr
		}
		return v10, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{joined})
	if callErr != nil {
		return nil, callErr
	}
	arg__21309, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{joined})
	if callErr != nil {
		return nil, callErr
	}
	v52, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("unknown"), arg__21309})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v52) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v70, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{joined})
	if callErr != nil {
		return nil, callErr
	}
	arg__21322, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{joined})
	if callErr != nil {
		return nil, callErr
	}
	v59, callErr = rt.InvokeValue(rt.LookupVar("ir", "set-type!").Deref(), []vm.Value{f, nid, arg__21322})
	if callErr != nil {
		return nil, callErr
	}
	v63 = v59
	goto b6
b5:
	;
	v63 = vm.NIL
	goto b6
b6:
	;
	v70 = v63
	goto b3
}
func infer_arg_types(arg0 vm.Value) (vm.Value, error) {
	var entry_bid vm.Value
	var insts vm.Value
	var use_index vm.Value
	var doseq_seq__21323 vm.Value
	var doseq_loop__21324 vm.Value
	var f vm.Value
	var nid vm.Value
	var arg__21348 vm.Value
	var v46 bool
	var v62 vm.Value
	var callErr error
	entry_bid, callErr = rt.InvokeValue(rt.LookupVar("ir", "entry-block").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{entry_bid, arg0})
	if callErr != nil {
		return nil, callErr
	}
	use_index, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__21323, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21324 = doseq_seq__21323
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__21324) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__21324})
	if callErr != nil {
		return nil, callErr
	}
	arg__21348, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v46 = arg__21348 == vm.Keyword("load-arg")
	if v46 {
		goto b5
	} else {
		goto b6
	}
b3:
	;
	goto b4
b4:
	;
	return f, nil
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.infer-arg-types", "infer-one-load-arg!").Deref(), []vm.Value{f, use_index, nid})
	if callErr != nil {
		return nil, callErr
	}
	goto b7
b6:
	;
	goto b7
b7:
	;
	v62, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__21324})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21324 = v62
	goto b1
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.infer-arg-types", map[string]vm.Value{"join-types": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("join-types: wrong number of arguments %d (expected 2)", len(args))
		}
		return join_types(args[0], args[1]), nil
	}), "constraint-from-user": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("constraint-from-user: wrong number of arguments %d (expected 1)", len(args))
		}
		return constraint_from_user(args[0])
	}), "infer-one-load-arg!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("infer-one-load-arg!: wrong number of arguments %d (expected 3)", len(args))
		}
		return infer_one_load_arg_BANG_(args[0], args[1], args[2])
	}), "infer-arg-types": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("infer-arg-types: wrong number of arguments %d (expected 1)", len(args))
		}
		return infer_arg_types(args[0])
	}),
	})
}
