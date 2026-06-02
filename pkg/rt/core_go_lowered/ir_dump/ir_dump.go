package ir_dump

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func format_args(arg0 vm.Value) (vm.Value, error) {
	var arg__9566 vm.Value
	var v12 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("v"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9566, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("v"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(", "), arg__9566})
	if callErr != nil {
		return nil, callErr
	}
	return v12, nil
}
func scalar_type_display(arg0 vm.Value) vm.Value {
	var v6 bool
	var case__9567 vm.Value
	var v15 bool
	var v154 vm.Value
	var v24 bool
	var v150 vm.Value
	var v33 bool
	var v146 vm.Value
	var v42 bool
	var v142 vm.Value
	var v51 bool
	var v138 vm.Value
	var v60 bool
	var v134 vm.Value
	var v69 bool
	var v130 vm.Value
	var v78 bool
	var v126 vm.Value
	var v87 bool
	var v122 vm.Value
	var v96 bool
	var v118 vm.Value
	var v114 vm.Value
	var v110 vm.Value
	v6 = arg0 == vm.Keyword("unknown")
	if v6 {
		goto b1
	} else {
		case__9567 = arg0
		goto b2
	}
b1:
	;
	v154 = vm.String("unknown")
	goto b3
b2:
	;
	v15 = case__9567 == vm.Keyword("bottom")
	if v15 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v154
b4:
	;
	v150 = vm.String("bottom")
	goto b6
b5:
	;
	v24 = case__9567 == vm.Keyword("true")
	if v24 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v154 = v150
	goto b3
b7:
	;
	v146 = vm.String("true")
	goto b9
b8:
	;
	v33 = case__9567 == vm.Keyword("false")
	if v33 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v150 = v146
	goto b6
b10:
	;
	v142 = vm.String("false")
	goto b12
b11:
	;
	v42 = case__9567 == vm.Keyword("int")
	if v42 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v146 = v142
	goto b9
b13:
	;
	v138 = vm.String("int")
	goto b15
b14:
	;
	v51 = case__9567 == vm.Keyword("float")
	if v51 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v142 = v138
	goto b12
b16:
	;
	v134 = vm.String("float")
	goto b18
b17:
	;
	v60 = case__9567 == vm.Keyword("number")
	if v60 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v138 = v134
	goto b15
b19:
	;
	v130 = vm.String("number")
	goto b21
b20:
	;
	v69 = case__9567 == vm.Keyword("bool")
	if v69 {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v134 = v130
	goto b18
b22:
	;
	v126 = vm.String("bool")
	goto b24
b23:
	;
	v78 = case__9567 == vm.Keyword("nil")
	if v78 {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v130 = v126
	goto b21
b25:
	;
	v122 = vm.String("nil")
	goto b27
b26:
	;
	v87 = case__9567 == vm.Keyword("string")
	if v87 {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v126 = v122
	goto b24
b28:
	;
	v118 = vm.String("string")
	goto b30
b29:
	;
	v96 = case__9567 == vm.Keyword("any")
	if v96 {
		goto b31
	} else {
		goto b32
	}
b30:
	;
	v122 = v118
	goto b27
b31:
	;
	v114 = vm.String("any")
	goto b33
b32:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b34
	} else {
		goto b35
	}
b33:
	;
	v118 = v114
	goto b30
b34:
	;
	v110 = vm.String("??")
	goto b36
b35:
	;
	v110 = vm.NIL
	goto b36
b36:
	;
	v114 = v110
	goto b33
}
func type_display(arg0 vm.Value) (vm.Value, error) {
	var v5 vm.Value
	var t vm.Value
	var v8 vm.Value
	var and__x vm.Value
	var v236 vm.Value
	var arg__9632 vm.Value
	var arg__9660 vm.Value
	var arg__9661 vm.Value
	var arg__9690 vm.Value
	var arg__9718 vm.Value
	var arg__9719 vm.Value
	var arg__9720 vm.Value
	var arg__9751 vm.Value
	var arg__9779 vm.Value
	var arg__9780 vm.Value
	var arg__9809 vm.Value
	var arg__9837 vm.Value
	var arg__9838 vm.Value
	var arg__9839 vm.Value
	var arg__9840 vm.Value
	var v122 vm.Value
	var v233 vm.Value
	var arg__9604 vm.Value
	var v22 bool
	var v25 vm.Value
	var tag vm.Value
	var v vm.Value
	var v160 bool
	var v230 vm.Value
	var arg__9849 vm.Value
	var v136 bool
	var v139 vm.Value
	var v167 vm.Value
	var case__9590 vm.Value
	var v178 bool
	var v213 vm.Value
	var v185 vm.Value
	var v207 vm.Value
	var v201 vm.Value
	var v227 vm.Value
	var callErr error
	v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v5) {
		t = arg0
		goto b1
	} else {
		t = arg0
		goto b2
	}
b1:
	;
	v8, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "scalar-type-display").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v236 = v8
	goto b3
b2:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b7
	} else {
		goto b8
	}
b3:
	;
	return v236, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9632, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v6 vm.Value
		var callErr error
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-order").Deref(), arg0, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__9632})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9660, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9661, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v6 vm.Value
		var callErr error
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-order").Deref(), arg0, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__9660})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-display").Deref(), arg__9661})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9690, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v6 vm.Value
		var callErr error
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-order").Deref(), arg0, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__9690})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9718, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9719, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v6 vm.Value
		var callErr error
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-order").Deref(), arg0, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__9718})
	if callErr != nil {
		return nil, callErr
	}
	arg__9720, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-display").Deref(), arg__9719})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(","), arg__9720})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9751, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v6 vm.Value
		var callErr error
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-order").Deref(), arg0, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__9751})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9779, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9780, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v6 vm.Value
		var callErr error
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-order").Deref(), arg0, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__9779})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-display").Deref(), arg__9780})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9809, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v6 vm.Value
		var callErr error
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-order").Deref(), arg0, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__9809})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9837, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9838, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v6 vm.Value
		var callErr error
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-order").Deref(), arg0, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__9837})
	if callErr != nil {
		return nil, callErr
	}
	arg__9839, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.dump", "type-display").Deref(), arg__9838})
	if callErr != nil {
		return nil, callErr
	}
	arg__9840, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(","), arg__9839})
	if callErr != nil {
		return nil, callErr
	}
	v122, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("union{"), arg__9840, vm.String("}")})
	if callErr != nil {
		return nil, callErr
	}
	v233 = v122
	goto b6
b5:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b13
	} else {
		goto b14
	}
b6:
	;
	v236 = v233
	goto b3
b7:
	;
	arg__9604, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v22 = arg__9604 == vm.Keyword("union")
	v25 = vm.Boolean(v22)
	goto b9
b8:
	;
	v25 = and__x
	goto b9
b9:
	;
	if vm.IsTruthy(v25) {
		goto b4
	} else {
		goto b5
	}
b10:
	;
	tag, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{t, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{t, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	v160 = tag == vm.Keyword("int")
	if v160 {
		goto b16
	} else {
		case__9590 = tag
		goto b17
	}
b11:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b25
	} else {
		goto b26
	}
b12:
	;
	v233 = v230
	goto b6
b13:
	;
	arg__9849, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v136 = arg__9849 == vm.Keyword("const")
	v139 = vm.Boolean(v136)
	goto b15
b14:
	;
	v139 = and__x
	goto b15
b15:
	;
	if vm.IsTruthy(v139) {
		goto b10
	} else {
		goto b11
	}
b16:
	;
	v167, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("int("), v, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	v213 = v167
	goto b18
b17:
	;
	v178 = case__9590 == vm.Keyword("float")
	if v178 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v230 = v213
	goto b12
b19:
	;
	v185, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("float("), v, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	v207 = v185
	goto b21
b20:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v213 = v207
	goto b18
b22:
	;
	v201 = vm.String("??")
	goto b24
b23:
	;
	v201 = vm.NIL
	goto b24
b24:
	;
	v207 = v201
	goto b21
b25:
	;
	v227 = vm.String("??")
	goto b27
b26:
	;
	v227 = vm.NIL
	goto b27
b27:
	;
	v230 = v227
	goto b12
}
func op_display_name(arg0 vm.Value) (vm.Value, error) {
	var arg__9892 vm.Value
	var arg__9896 vm.Value
	var arg__9912 vm.Value
	var arg__9916 vm.Value
	var arg__9917 vm.Value
	var arg__9934 vm.Value
	var arg__9938 vm.Value
	var arg__9954 vm.Value
	var arg__9958 vm.Value
	var arg__9959 vm.Value
	var arg__9960 vm.Value
	var v69 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "re-pattern").Deref(), []vm.Value{vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9892, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9896, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "re-pattern").Deref(), []vm.Value{vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "split").Deref(), []vm.Value{arg__9892, arg__9896})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "re-pattern").Deref(), []vm.Value{vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9916, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "re-pattern").Deref(), []vm.Value{vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9917, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "split").Deref(), []vm.Value{arg__9912, arg__9916})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("clojure.string", "capitalize").Deref(), arg__9917})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "re-pattern").Deref(), []vm.Value{vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9934, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9938, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "re-pattern").Deref(), []vm.Value{vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "split").Deref(), []vm.Value{arg__9934, arg__9938})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "re-pattern").Deref(), []vm.Value{vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9954, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9958, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "re-pattern").Deref(), []vm.Value{vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9959, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "split").Deref(), []vm.Value{arg__9954, arg__9958})
	if callErr != nil {
		return nil, callErr
	}
	arg__9960, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("clojure.string", "capitalize").Deref(), arg__9959})
	if callErr != nil {
		return nil, callErr
	}
	v69, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(""), arg__9960})
	if callErr != nil {
		return nil, callErr
	}
	return v69, nil
}
func format_refs(arg0 vm.Value) (vm.Value, error) {
	var arg__9995 vm.Value
	var v12 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String(" v"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9995, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String(" v"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(""), arg__9995})
	if callErr != nil {
		return nil, callErr
	}
	return v12, nil
}
func format_target(arg0 vm.Value) (vm.Value, error) {
	var arg__10010 vm.Value
	var arg__10018 vm.Value
	var arg__10028 vm.Value
	var arg__10029 vm.Value
	var v24 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10010, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-args").Deref(), []vm.Value{arg__10010})
	if callErr != nil {
		return nil, callErr
	}
	arg__10018, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10028, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10029, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-args").Deref(), []vm.Value{arg__10028})
	if callErr != nil {
		return nil, callErr
	}
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg__10018, vm.String("("), arg__10029, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	return v24, nil
}
func terminator_targets_str(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v7 bool
	var aux vm.Value
	var arg__10043 vm.Value
	var v16 vm.Value
	var op vm.Value
	var v23 bool
	var v74 vm.Value
	var arg__10055 vm.Value
	var arg__10066 vm.Value
	var arg__10078 vm.Value
	var arg__10079 vm.Value
	var arg__10089 vm.Value
	var arg__10090 vm.Value
	var v54 vm.Value
	var v70 vm.Value
	var v66 vm.Value
	var callErr error
	v7 = arg0 == vm.Keyword("branch")
	if v7 {
		aux = arg1
		goto b1
	} else {
		op = arg0
		aux = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-target").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10043, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-target").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v16, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String(" -> "), arg__10043})
	if callErr != nil {
		return nil, callErr
	}
	v74 = v16
	goto b3
b2:
	;
	v23 = op == vm.Keyword("branch-if")
	if v23 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v74, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10055, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-target").Deref(), []vm.Value{arg__10055})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10066, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-target").Deref(), []vm.Value{arg__10066})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10078, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10079, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-target").Deref(), []vm.Value{arg__10078})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10089, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10090, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-target").Deref(), []vm.Value{arg__10089})
	if callErr != nil {
		return nil, callErr
	}
	v54, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String(" -> "), arg__10079, vm.String(" : "), arg__10090})
	if callErr != nil {
		return nil, callErr
	}
	v70 = v54
	goto b6
b5:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v74 = v70
	goto b3
b7:
	;
	v66 = vm.String("")
	goto b9
b8:
	;
	v66 = vm.NIL
	goto b9
b9:
	;
	v70 = v66
	goto b6
}
func write_node(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var op vm.Value
	var refs vm.Value
	var aux vm.Value
	var v21 vm.Value
	var arg__10132 vm.Value
	var arg__10136 vm.Value
	var arg__10142 vm.Value
	var v40 vm.Value
	var f vm.Value
	var id vm.Value
	var t vm.Value
	var arg__10168 vm.Value
	var v77 vm.Value
	var v269 vm.Value
	var v125 vm.Value
	var arg__10190 vm.Value
	var arg__10200 vm.Value
	var arg__10204 vm.Value
	var arg__10213 vm.Value
	var v188 vm.Value
	var arg__10195 vm.Value
	var head__10193 vm.Value
	var v193 vm.Value
	var arg__10219 vm.Value
	var v239 vm.Value
	var arg__10235 vm.Value
	var v248 vm.Value
	var arg__10236 vm.Value
	var v267 vm.Value
	var callErr error
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v21, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.dump", "terminator-ops").Deref(), op})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v21) {
		goto b1
	} else {
		f = arg0
		id = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "op-display-name").Deref(), []vm.Value{op})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-refs").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "terminator-targets-str").Deref(), []vm.Value{op, aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10132, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "op-display-name").Deref(), []vm.Value{op})
	if callErr != nil {
		return nil, callErr
	}
	arg__10136, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-refs").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	arg__10142, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "terminator-targets-str").Deref(), []vm.Value{op, aux})
	if callErr != nil {
		return nil, callErr
	}
	v40, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("    "), arg__10132, arg__10136, arg__10142, vm.String("\n")})
	if callErr != nil {
		return nil, callErr
	}
	v269 = v40
	goto b3
b2:
	;
	t, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{id, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "op-display-name").Deref(), []vm.Value{op})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-refs").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10168, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v77, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__10168})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v77) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v269, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String(" ; "), aux})
	if callErr != nil {
		return nil, callErr
	}
	goto b6
b5:
	;
	goto b6
b6:
	;
	v125, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{t, vm.Keyword("unknown")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v125) {
		goto b7
	} else {
		goto b8
	}
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__10190, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String(" : "), arg__10190})
	if callErr != nil {
		return nil, callErr
	}
	goto b9
b8:
	;
	goto b9
b9:
	;
	arg__10200, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "op-display-name").Deref(), []vm.Value{op})
	if callErr != nil {
		return nil, callErr
	}
	arg__10204, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "format-refs").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10213, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v188, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__10213})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v188) {
		arg__10195 = id
		head__10193 = rt.LookupVar("clojure.core", "str").Deref()
		goto b10
	} else {
		arg__10195 = id
		head__10193 = rt.LookupVar("clojure.core", "str").Deref()
		goto b11
	}
b10:
	;
	v193, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String(" ; "), aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__10219 = v193
	goto b12
b11:
	;
	arg__10219 = vm.String("")
	goto b12
b12:
	;
	v239, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{t, vm.Keyword("unknown")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v239) {
		goto b13
	} else {
		goto b14
	}
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__10235, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v248, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String(" : "), arg__10235})
	if callErr != nil {
		return nil, callErr
	}
	arg__10236 = v248
	goto b15
b14:
	;
	arg__10236 = vm.String("")
	goto b15
b15:
	;
	v267, callErr = rt.InvokeValue(head__10193, []vm.Value{vm.String("    v"), arg__10195, vm.String(" = "), arg__10200, arg__10204, arg__10219, arg__10236, vm.String("\n")})
	if callErr != nil {
		return nil, callErr
	}
	v269 = v267
	goto b3
}
func write_block(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var params vm.Value
	var preds vm.Value
	var insts vm.Value
	var term vm.Value
	var arg__10262 vm.Value
	var entry_QMARK__12 bool
	var f vm.Value
	var bid vm.Value
	var entry_QMARK__20 bool
	var entry_QMARK__28 bool
	var entry_QMARK__42 vm.Value
	var arg__10422 vm.Value
	var v96 vm.Value
	var entry_QMARK__75 vm.Value
	var arg__10463 vm.Value
	var arg__10501 vm.Value
	var arg__10502 vm.Value
	var entry_QMARK__89 vm.Value
	var entry_QMARK__138 vm.Value
	var head__10505 vm.Value
	var arg__10507 string
	var arg__10665 vm.Value
	var arg__10666 vm.Value
	var v234 vm.Value
	var arg__10509 vm.Value
	var arg__10706 vm.Value
	var arg__10744 vm.Value
	var arg__10745 vm.Value
	var v263 vm.Value
	var arg__10746 vm.Value
	var header vm.Value
	var arg__10782 vm.Value
	var body vm.Value
	var or__x vm.Value
	var v365 vm.Value
	var term_line vm.Value
	var v382 vm.Value
	var v350 vm.Value
	var v352 vm.Value
	var callErr error
	params, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	preds, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10262, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	entry_QMARK__12 = arg1 == arg__10262
	if entry_QMARK__12 {
		f = arg0
		bid = arg1
		entry_QMARK__20 = entry_QMARK__12
		goto b1
	} else {
		f = arg0
		bid = arg1
		entry_QMARK__28 = entry_QMARK__12
		goto b2
	}
b1:
	;
	entry_QMARK__42 = vm.Boolean(entry_QMARK__20)
	goto b3
b2:
	;
	entry_QMARK__42 = vm.Boolean(entry_QMARK__28)
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__10322 vm.Value
		var arg__10340 vm.Value
		var arg__10341 vm.Value
		var v19 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10322, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{arg__10322})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10340, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10341, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{arg__10340})
		if callErr != nil {
			return nil, callErr
		}
		v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("v"), arg0, vm.String(": "), arg__10341})
		if callErr != nil {
			return nil, callErr
		}
		return v19, nil
	}), params})
	if callErr != nil {
		return nil, callErr
	}
	arg__10422, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__10400 vm.Value
		var arg__10418 vm.Value
		var arg__10419 vm.Value
		var v19 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10400, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{arg__10400})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10418, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10419, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{arg__10418})
		if callErr != nil {
			return nil, callErr
		}
		v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("v"), arg0, vm.String(": "), arg__10419})
		if callErr != nil {
			return nil, callErr
		}
		return v19, nil
	}), params})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(", "), arg__10422})
	if callErr != nil {
		return nil, callErr
	}
	v96, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{preds})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v96) {
		entry_QMARK__75 = entry_QMARK__42
		goto b4
	} else {
		entry_QMARK__89 = entry_QMARK__42
		goto b5
	}
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	arg__10463, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(", "), arg__10463})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	arg__10501, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	arg__10502, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(", "), arg__10501})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("    ; preds: "), arg__10502})
	if callErr != nil {
		return nil, callErr
	}
	entry_QMARK__138 = entry_QMARK__75
	goto b6
b5:
	;
	entry_QMARK__138 = entry_QMARK__89
	goto b6
b6:
	;
	if vm.IsTruthy(entry_QMARK__138) {
		head__10505 = rt.LookupVar("clojure.core", "str").Deref()
		goto b7
	} else {
		head__10505 = rt.LookupVar("clojure.core", "str").Deref()
		goto b8
	}
b7:
	;
	arg__10507 = "entry "
	goto b9
b8:
	;
	arg__10507 = ""
	goto b9
b9:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__10565 vm.Value
		var arg__10583 vm.Value
		var arg__10584 vm.Value
		var v19 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10565, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{arg__10565})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10583, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10584, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{arg__10583})
		if callErr != nil {
			return nil, callErr
		}
		v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("v"), arg0, vm.String(": "), arg__10584})
		if callErr != nil {
			return nil, callErr
		}
		return v19, nil
	}), params})
	if callErr != nil {
		return nil, callErr
	}
	arg__10665, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__10643 vm.Value
		var arg__10661 vm.Value
		var arg__10662 vm.Value
		var v19 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10643, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{arg__10643})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10661, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__10662, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "type-display").Deref(), []vm.Value{arg__10661})
		if callErr != nil {
			return nil, callErr
		}
		v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("v"), arg0, vm.String(": "), arg__10662})
		if callErr != nil {
			return nil, callErr
		}
		return v19, nil
	}), params})
	if callErr != nil {
		return nil, callErr
	}
	arg__10666, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(", "), arg__10665})
	if callErr != nil {
		return nil, callErr
	}
	v234, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{preds})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v234) {
		arg__10509 = bid
		goto b10
	} else {
		arg__10509 = bid
		goto b11
	}
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	arg__10706, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(", "), arg__10706})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	arg__10744, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	arg__10745, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(", "), arg__10744})
	if callErr != nil {
		return nil, callErr
	}
	v263, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("    ; preds: "), arg__10745})
	if callErr != nil {
		return nil, callErr
	}
	arg__10746 = v263
	goto b12
b11:
	;
	arg__10746 = vm.String("")
	goto b12
b12:
	;
	header, callErr = rt.InvokeValue(head__10505, []vm.Value{vm.String("  "), vm.String(arg__10507), vm.String("b"), arg__10509, vm.String("("), arg__10666, vm.String("):"), arg__10746, vm.String("\n")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "write-node").Deref(), []vm.Value{f, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), insts})
	if callErr != nil {
		return nil, callErr
	}
	arg__10782, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "write-node").Deref(), []vm.Value{f, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), insts})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(""), arg__10782})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{term, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b16
	} else {
		goto b17
	}
b13:
	;
	v365, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "write-node").Deref(), []vm.Value{f, term})
	if callErr != nil {
		return nil, callErr
	}
	term_line = v365
	goto b15
b14:
	;
	term_line = vm.String("")
	goto b15
b15:
	;
	v382, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{header, body, term_line, vm.String("\n")})
	if callErr != nil {
		return nil, callErr
	}
	return v382, nil
b16:
	;
	v352 = or__x
	goto b18
b17:
	;
	v350, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	v352 = v350
	goto b18
b18:
	;
	if vm.IsTruthy(v352) {
		goto b13
	} else {
		goto b14
	}
}
func dump(arg0 vm.Value) (vm.Value, error) {
	var arg__10842 vm.Value
	var arg__10866 vm.Value
	var arg__10867 vm.Value
	var arg__10874 vm.Value
	var arg__10879 vm.Value
	var arg__10884 vm.Value
	var arg__10907 vm.Value
	var arg__10931 vm.Value
	var arg__10932 vm.Value
	var arg__10933 vm.Value
	var v79 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arity").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10842, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "write-block").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), arg__10842})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10866, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10867, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "write-block").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), arg__10866})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(""), arg__10867})
	if callErr != nil {
		return nil, callErr
	}
	arg__10874, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10879, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arity").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10884, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10907, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "write-block").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), arg__10907})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10931, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__10932, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "write-block").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), arg__10931})
	if callErr != nil {
		return nil, callErr
	}
	arg__10933, callErr = rt.InvokeValue(rt.LookupVar("clojure.string", "join").Deref(), []vm.Value{vm.String(""), arg__10932})
	if callErr != nil {
		return nil, callErr
	}
	v79, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("fn "), arg__10874, vm.String("(arity="), arg__10879, vm.String(", variadic="), arg__10884, vm.String("):\n"), arg__10933})
	if callErr != nil {
		return nil, callErr
	}
	return v79, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.dump", map[string]vm.Value{"format-args": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("format-args: wrong number of arguments %d (expected 1)", len(args))
		}
		return format_args(args[0])
	}), "scalar-type-display": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("scalar-type-display: wrong number of arguments %d (expected 1)", len(args))
		}
		return scalar_type_display(args[0]), nil
	}), "type-display": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("type-display: wrong number of arguments %d (expected 1)", len(args))
		}
		return type_display(args[0])
	}), "op-display-name": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("op-display-name: wrong number of arguments %d (expected 1)", len(args))
		}
		return op_display_name(args[0])
	}), "format-refs": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("format-refs: wrong number of arguments %d (expected 1)", len(args))
		}
		return format_refs(args[0])
	}), "format-target": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("format-target: wrong number of arguments %d (expected 1)", len(args))
		}
		return format_target(args[0])
	}), "terminator-targets-str": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("terminator-targets-str: wrong number of arguments %d (expected 2)", len(args))
		}
		return terminator_targets_str(args[0], args[1])
	}), "write-node": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("write-node: wrong number of arguments %d (expected 2)", len(args))
		}
		return write_node(args[0], args[1])
	}), "write-block": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("write-block: wrong number of arguments %d (expected 2)", len(args))
		}
		return write_block(args[0], args[1])
	}), "dump": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("dump: wrong number of arguments %d (expected 1)", len(args))
		}
		return dump(args[0])
	}),
	})
}
