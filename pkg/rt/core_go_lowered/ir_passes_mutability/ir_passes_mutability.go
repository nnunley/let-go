package ir_passes_mutability

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func mutating_var_call_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__20811 vm.Value
	var v10 bool
	var nid vm.Value
	var f vm.Value
	var refs vm.Value
	var v21 vm.Value
	var v93 vm.Value
	var v24 vm.Value
	var callee vm.Value
	var arg__20829 vm.Value
	var and__x_46 bool
	var and__x_37 vm.Value
	var v84 vm.Value
	var arg__20843 vm.Value
	var arg__20859 vm.Value
	var arg__20860 vm.Value
	var v73 vm.Value
	var and__x_56 bool
	var v76 vm.Value
	var callErr error
	arg__20811, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v10 = arg__20811 == vm.Keyword("call")
	if v10 {
		nid = arg0
		f = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v21, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v21) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v93 = vm.NIL
	goto b3
b3:
	;
	return v93, nil
b4:
	;
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	callee = v24
	goto b6
b5:
	;
	callee = vm.NIL
	goto b6
b6:
	;
	if vm.IsTruthy(callee) {
		goto b7
	} else {
		and__x_37 = callee
		goto b8
	}
b7:
	;
	arg__20829, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{callee, f})
	if callErr != nil {
		return nil, callErr
	}
	and__x_46 = arg__20829 == vm.Keyword("load-var")
	if and__x_46 {
		goto b10
	} else {
		and__x_56 = and__x_46
		goto b11
	}
b8:
	;
	v84 = and__x_37
	goto b9
b9:
	;
	v93 = v84
	goto b3
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{callee, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__20843, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{callee, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__20843})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{callee, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__20859, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{callee, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__20860, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__20859})
	if callErr != nil {
		return nil, callErr
	}
	v73, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.passes.mutability", "known-var-mutating-builtins").Deref(), arg__20860})
	if callErr != nil {
		return nil, callErr
	}
	v76 = v73
	goto b12
b11:
	;
	v76 = vm.Boolean(and__x_56)
	goto b12
b12:
	;
	v84 = v76
	goto b9
}
func stable_load_var_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__20867 vm.Value
	var and__x vm.Value
	var facts vm.Value
	var var_value vm.Value
	var arg__20875 vm.Value
	var arg__20886 vm.Value
	var arg__20888 vm.Value
	var v29 vm.Value
	var v32 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(vm.Keyword("unknown-all?"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__20867, callErr = rt.InvokeValue(vm.Keyword("unknown-all?"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__20867})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		facts = arg0
		var_value = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("mutated-vars"), []vm.Value{facts})
	if callErr != nil {
		return nil, callErr
	}
	arg__20875, callErr = rt.InvokeValue(vm.Keyword("mutated-vars"), []vm.Value{facts})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg__20875, var_value})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("mutated-vars"), []vm.Value{facts})
	if callErr != nil {
		return nil, callErr
	}
	arg__20886, callErr = rt.InvokeValue(vm.Keyword("mutated-vars"), []vm.Value{facts})
	if callErr != nil {
		return nil, callErr
	}
	arg__20888, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg__20886, var_value})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__20888})
	if callErr != nil {
		return nil, callErr
	}
	v32 = v29
	goto b3
b2:
	;
	v32 = and__x
	goto b3
b3:
	;
	return v32, nil
}
func analyze_var_stability(arg0 vm.Value) (vm.Value, error) {
	var arg__20893 vm.Value
	var arg__20900 vm.Value
	var arg__20901 vm.Value
	var arg__20908 vm.Value
	var arg__20915 vm.Value
	var arg__20916 vm.Value
	var arg__20917 vm.Value
	var v27 vm.Value
	var v29 vm.Value
	var ids vm.Value
	var mutated_vars vm.Value
	var unknown_all_QMARK__4 vm.Value
	var f vm.Value
	var unknown_all_QMARK__34 vm.Value
	var v65 vm.Value
	var unknown_all_QMARK__38 vm.Value
	var nid vm.Value
	var op vm.Value
	var v84 bool
	var v157 vm.Value
	var or__x_42 bool
	var unknown_all_QMARK__43 bool
	var unknown_all_QMARK__48 bool
	var v53 vm.Value
	var v55 vm.Value
	var unknown_all_QMARK__59 vm.Value
	var unknown_all_QMARK__73 vm.Value
	var v87 vm.Value
	var arg__20954 vm.Value
	var v93 vm.Value
	var unknown_all_QMARK__79 vm.Value
	var v108 vm.Value
	var v111 vm.Value
	var unknown_all_QMARK__103 vm.Value
	var unknown_all_QMARK__116 vm.Value
	var v129 vm.Value
	var callErr error
	arg__20893, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__20893})
	if callErr != nil {
		return nil, callErr
	}
	arg__20900, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__20901, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__20900})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__20901})
	if callErr != nil {
		return nil, callErr
	}
	arg__20908, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__20908})
	if callErr != nil {
		return nil, callErr
	}
	arg__20915, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__20916, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__20915})
	if callErr != nil {
		return nil, callErr
	}
	arg__20917, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__20916})
	if callErr != nil {
		return nil, callErr
	}
	v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{arg__20917})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	ids = v27
	mutated_vars = v29
	unknown_all_QMARK__4 = vm.Boolean(false)
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(unknown_all_QMARK__4) {
		or__x_42 = vm.IsTruthy(unknown_all_QMARK__4)
		unknown_all_QMARK__43 = vm.IsTruthy(unknown_all_QMARK__4)
		goto b5
	} else {
		unknown_all_QMARK__48 = vm.IsTruthy(unknown_all_QMARK__4)
		goto b6
	}
b2:
	;
	v65, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("mutated-vars"), mutated_vars, vm.Keyword("unknown-all?"), unknown_all_QMARK__34})
	if callErr != nil {
		return nil, callErr
	}
	v157 = v65
	goto b4
b3:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{ids})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v84 = op == vm.Keyword("set-var")
	if v84 {
		unknown_all_QMARK__73 = unknown_all_QMARK__38
		goto b8
	} else {
		unknown_all_QMARK__79 = unknown_all_QMARK__38
		goto b9
	}
b4:
	;
	return v157, nil
b5:
	;
	v55 = vm.Boolean(or__x_42)
	unknown_all_QMARK__59 = vm.Boolean(unknown_all_QMARK__43)
	goto b7
b6:
	;
	v53, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{ids})
	if callErr != nil {
		return nil, callErr
	}
	v55 = v53
	unknown_all_QMARK__59 = vm.Boolean(unknown_all_QMARK__48)
	goto b7
b7:
	;
	if vm.IsTruthy(v55) {
		unknown_all_QMARK__34 = unknown_all_QMARK__59
		goto b2
	} else {
		unknown_all_QMARK__38 = unknown_all_QMARK__59
		goto b3
	}
b8:
	;
	v87, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{ids})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__20954, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v93, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{mutated_vars, arg__20954})
	if callErr != nil {
		return nil, callErr
	}
	ids = v87
	mutated_vars = v93
	unknown_all_QMARK__4 = unknown_all_QMARK__73
	goto b1
b9:
	;
	v108, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.mutability", "mutating-var-call?").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v108) {
		goto b11
	} else {
		unknown_all_QMARK__103 = unknown_all_QMARK__79
		goto b12
	}
b10:
	;
	v157 = vm.NIL
	goto b4
b11:
	;
	v111, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{ids})
	if callErr != nil {
		return nil, callErr
	}
	ids = v111
	unknown_all_QMARK__4 = vm.Boolean(true)
	goto b1
b12:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		unknown_all_QMARK__116 = unknown_all_QMARK__103
		goto b14
	} else {
		goto b15
	}
b13:
	;
	goto b10
b14:
	;
	v129, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{ids})
	if callErr != nil {
		return nil, callErr
	}
	ids = v129
	unknown_all_QMARK__4 = unknown_all_QMARK__116
	goto b1
b15:
	;
	goto b16
b16:
	;
	goto b13
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.mutability", map[string]vm.Value{"mutating-var-call?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("mutating-var-call?: wrong number of arguments %d (expected 2)", len(args))
		}
		return mutating_var_call_QMARK_(args[0], args[1])
	}), "stable-load-var?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("stable-load-var?: wrong number of arguments %d (expected 2)", len(args))
		}
		return stable_load_var_QMARK_(args[0], args[1])
	}), "analyze-var-stability": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("analyze-var-stability: wrong number of arguments %d (expected 1)", len(args))
		}
		return analyze_var_stability(args[0])
	}),
	})
}
