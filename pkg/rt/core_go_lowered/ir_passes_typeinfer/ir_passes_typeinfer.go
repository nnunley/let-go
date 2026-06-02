package ir_passes_typeinfer

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func const_type_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var t vm.Value
	var arg__28395 vm.Value
	var v11 bool
	var v14 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		t = arg0
		goto b1
	} else {
		goto b2
	}
b1:
	;
	arg__28395, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v11 = arg__28395 == vm.Keyword("const")
	v14 = vm.Boolean(v11)
	goto b3
b2:
	;
	v14 = and__x
	goto b3
b3:
	;
	return v14, nil
}
func sort_members(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var member vm.Value
		var v9 vm.Value
		var key vm.Value
		var v19 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "const-type?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v4) {
			member = arg0
			goto b1
		} else {
			member = arg0
			goto b2
		}
	b1:
		;
		v9, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{member, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		key = v9
		goto b3
	b2:
		;
		key = member
		goto b3
	b3:
		;
		v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.passes.typeinfer", "type-order").Deref(), key, vm.Int(100)})
		if callErr != nil {
			return nil, callErr
		}
		return v19, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func union_type_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var t vm.Value
	var arg__28438 vm.Value
	var v11 bool
	var v14 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		t = arg0
		goto b1
	} else {
		goto b2
	}
b1:
	;
	arg__28438, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v11 = arg__28438 == vm.Keyword("union")
	v14 = vm.Boolean(v11)
	goto b3
b2:
	;
	v14 = and__x
	goto b3
b3:
	;
	return v14, nil
}
func ti_inc_BANG_(arg0 vm.Value) (vm.Value, error) {
	var k vm.Value
	var arg__28457 vm.Value
	var v22 vm.Value
	var v26 vm.Value
	var callErr error
	if vm.IsTruthy(rt.LookupVar("ir.passes.typeinfer", "*ti-counters*").Deref()) {
		k = arg0
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "fnil").Deref(), []vm.Value{rt.LookupVar("clojure.core", "inc").Deref(), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__28457, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "fnil").Deref(), []vm.Value{rt.LookupVar("clojure.core", "inc").Deref(), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{rt.LookupVar("ir.passes.typeinfer", "*ti-counters*").Deref(), rt.LookupVar("clojure.core", "update").Deref(), k, arg__28457})
	if callErr != nil {
		return nil, callErr
	}
	v26 = v22
	goto b3
b2:
	;
	v26 = vm.NIL
	goto b3
b3:
	;
	return v26, nil
}
func refine_truthy(arg0 vm.Value) (vm.Value, error) {
	var t vm.Value
	var v6 bool
	var v13 bool
	var v99 vm.Value
	var v20 bool
	var v96 vm.Value
	var v27 vm.Value
	var v93 vm.Value
	var arg__28905 vm.Value
	var arg__28910 vm.Value
	var arg__28929 vm.Value
	var arg__28930 vm.Value
	var arg__28954 vm.Value
	var arg__28959 vm.Value
	var arg__28978 vm.Value
	var arg__28979 vm.Value
	var arg__28980 vm.Value
	var v78 vm.Value
	var v90 vm.Value
	var v87 vm.Value
	var callErr error
	t, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v6 = t == vm.Keyword("bool")
	if v6 {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v99 = vm.Keyword("true")
	goto b3
b2:
	;
	v13 = t == vm.Keyword("false")
	if v13 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v99, nil
b4:
	;
	v96 = vm.Keyword("bottom")
	goto b6
b5:
	;
	v20 = t == vm.Keyword("nil")
	if v20 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v99 = v96
	goto b3
b7:
	;
	v93 = vm.Keyword("bottom")
	goto b9
b8:
	;
	v27, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "union-type?").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v27) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v96 = v93
	goto b6
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__28905, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "remove").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var or__x vm.Value
		var m vm.Value
		var v10 vm.Value
		var v12 vm.Value
		or__x = vm.Boolean(arg0 == vm.Keyword("nil"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v10 = vm.Boolean(m == vm.Keyword("false"))
		v12 = v10
		goto b3
	b3:
		;
		return v12
	}), arg__28905})
	if callErr != nil {
		return nil, callErr
	}
	arg__28910, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__28929, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__28930, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "remove").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var or__x vm.Value
		var m vm.Value
		var v10 vm.Value
		var v12 vm.Value
		or__x = vm.Boolean(arg0 == vm.Keyword("nil"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v10 = vm.Boolean(m == vm.Keyword("false"))
		v12 = v10
		goto b3
	b3:
		;
		return v12
	}), arg__28929})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__28910, arg__28930})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__28954, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "remove").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var or__x vm.Value
		var m vm.Value
		var v10 vm.Value
		var v12 vm.Value
		or__x = vm.Boolean(arg0 == vm.Keyword("nil"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v10 = vm.Boolean(m == vm.Keyword("false"))
		v12 = v10
		goto b3
	b3:
		;
		return v12
	}), arg__28954})
	if callErr != nil {
		return nil, callErr
	}
	arg__28959, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__28978, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__28979, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "remove").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var or__x vm.Value
		var m vm.Value
		var v10 vm.Value
		var v12 vm.Value
		or__x = vm.Boolean(arg0 == vm.Keyword("nil"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v10 = vm.Boolean(m == vm.Keyword("false"))
		v12 = v10
		goto b3
	b3:
		;
		return v12
	}), arg__28978})
	if callErr != nil {
		return nil, callErr
	}
	arg__28980, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__28959, arg__28979})
	if callErr != nil {
		return nil, callErr
	}
	v78, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__28980})
	if callErr != nil {
		return nil, callErr
	}
	v90 = v78
	goto b12
b11:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v93 = v90
	goto b9
b13:
	;
	v87 = t
	goto b15
b14:
	;
	v87 = vm.NIL
	goto b15
b15:
	;
	v90 = v87
	goto b12
}
func truthy_type_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var t vm.Value
	var v6 bool
	var v13 bool
	var v126 vm.Value
	var v20 bool
	var v123 vm.Value
	var v27 bool
	var v120 vm.Value
	var v34 bool
	var v117 vm.Value
	var v41 bool
	var v114 vm.Value
	var v48 bool
	var v111 vm.Value
	var arg__29003 vm.Value
	var v59 bool
	var v108 vm.Value
	var arg__29009 vm.Value
	var v70 bool
	var v105 vm.Value
	var v77 vm.Value
	var v102 vm.Value
	var arg__29023 vm.Value
	var v86 vm.Value
	var v99 vm.Value
	var v96 vm.Value
	var callErr error
	t, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v6 = t == vm.Keyword("true")
	if v6 {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v126 = vm.Boolean(true)
	goto b3
b2:
	;
	v13 = t == vm.Keyword("false")
	if v13 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v126, nil
b4:
	;
	v123 = vm.Boolean(false)
	goto b6
b5:
	;
	v20 = t == vm.Keyword("nil")
	if v20 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v126 = v123
	goto b3
b7:
	;
	v120 = vm.Boolean(false)
	goto b9
b8:
	;
	v27 = t == vm.Keyword("int")
	if v27 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v123 = v120
	goto b6
b10:
	;
	v117 = vm.Boolean(true)
	goto b12
b11:
	;
	v34 = t == vm.Keyword("float")
	if v34 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v120 = v117
	goto b9
b13:
	;
	v114 = vm.Boolean(true)
	goto b15
b14:
	;
	v41 = t == vm.Keyword("number")
	if v41 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v117 = v114
	goto b12
b16:
	;
	v111 = vm.Boolean(true)
	goto b18
b17:
	;
	v48 = t == vm.Keyword("string")
	if v48 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v114 = v111
	goto b15
b19:
	;
	v108 = vm.Boolean(true)
	goto b21
b20:
	;
	arg__29003, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v59 = t == arg__29003
	if v59 {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v111 = v108
	goto b18
b22:
	;
	v105 = vm.Boolean(true)
	goto b24
b23:
	;
	arg__29009, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v70 = t == arg__29009
	if v70 {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v108 = v105
	goto b21
b25:
	;
	v102 = vm.Boolean(true)
	goto b27
b26:
	;
	v77, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "union-type?").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v77) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v105 = v102
	goto b24
b28:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__29023, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v86, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("ir.passes.typeinfer", "truthy-type?").Deref(), arg__29023})
	if callErr != nil {
		return nil, callErr
	}
	v99 = v86
	goto b30
b29:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b31
	} else {
		goto b32
	}
b30:
	;
	v102 = v99
	goto b27
b31:
	;
	v96 = vm.Boolean(false)
	goto b33
b32:
	;
	v96 = vm.NIL
	goto b33
b33:
	;
	v99 = v96
	goto b30
}
func refine_falsey(arg0 vm.Value) (vm.Value, error) {
	var t vm.Value
	var v6 bool
	var v13 bool
	var v99 vm.Value
	var v20 vm.Value
	var v96 vm.Value
	var v27 vm.Value
	var v93 vm.Value
	var arg__29058 vm.Value
	var arg__29063 vm.Value
	var arg__29082 vm.Value
	var arg__29083 vm.Value
	var arg__29107 vm.Value
	var arg__29112 vm.Value
	var arg__29131 vm.Value
	var arg__29132 vm.Value
	var arg__29133 vm.Value
	var v78 vm.Value
	var v90 vm.Value
	var v87 vm.Value
	var callErr error
	t, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v6 = t == vm.Keyword("bool")
	if v6 {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v99 = vm.Keyword("false")
	goto b3
b2:
	;
	v13 = t == vm.Keyword("true")
	if v13 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v99, nil
b4:
	;
	v96 = vm.Keyword("bottom")
	goto b6
b5:
	;
	v20, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "truthy-type?").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v20) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v99 = v96
	goto b3
b7:
	;
	v93 = vm.Keyword("bottom")
	goto b9
b8:
	;
	v27, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "union-type?").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v27) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v96 = v93
	goto b6
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__29058, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var or__x vm.Value
		var m vm.Value
		var v10 vm.Value
		var v12 vm.Value
		or__x = vm.Boolean(arg0 == vm.Keyword("nil"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v10 = vm.Boolean(m == vm.Keyword("false"))
		v12 = v10
		goto b3
	b3:
		;
		return v12
	}), arg__29058})
	if callErr != nil {
		return nil, callErr
	}
	arg__29063, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__29082, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__29083, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var or__x vm.Value
		var m vm.Value
		var v10 vm.Value
		var v12 vm.Value
		or__x = vm.Boolean(arg0 == vm.Keyword("nil"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v10 = vm.Boolean(m == vm.Keyword("false"))
		v12 = v10
		goto b3
	b3:
		;
		return v12
	}), arg__29082})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__29063, arg__29083})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__29107, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var or__x vm.Value
		var m vm.Value
		var v10 vm.Value
		var v12 vm.Value
		or__x = vm.Boolean(arg0 == vm.Keyword("nil"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v10 = vm.Boolean(m == vm.Keyword("false"))
		v12 = v10
		goto b3
	b3:
		;
		return v12
	}), arg__29107})
	if callErr != nil {
		return nil, callErr
	}
	arg__29112, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__29131, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__29132, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var or__x vm.Value
		var m vm.Value
		var v10 vm.Value
		var v12 vm.Value
		or__x = vm.Boolean(arg0 == vm.Keyword("nil"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v10 = vm.Boolean(m == vm.Keyword("false"))
		v12 = v10
		goto b3
	b3:
		;
		return v12
	}), arg__29131})
	if callErr != nil {
		return nil, callErr
	}
	arg__29133, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__29112, arg__29132})
	if callErr != nil {
		return nil, callErr
	}
	v78, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__29133})
	if callErr != nil {
		return nil, callErr
	}
	v90 = v78
	goto b12
b11:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v93 = v90
	goto b9
b13:
	;
	v87 = t
	goto b15
b14:
	;
	v87 = vm.NIL
	goto b15
b15:
	;
	v90 = v87
	goto b12
}
func refine_edge_arg_type(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value, arg4 vm.Value) (vm.Value, error) {
	var term vm.Value
	var v20 vm.Value
	var arg_type vm.Value
	var target_bid vm.Value
	var arg_id vm.Value
	var f vm.Value
	var op vm.Value
	var refs vm.Value
	var aux vm.Value
	var or__x vm.Value
	var v304 vm.Value
	var cond_id vm.Value
	var tt vm.Value
	var ft vm.Value
	var arg__29178 vm.Value
	var true_edge_QMARK__97 bool
	var arg__29183 vm.Value
	var false_edge_QMARK__100 bool
	var v129 bool
	var v293 vm.Value
	var v74 vm.Value
	var v76 vm.Value
	var true_edge_QMARK__113 bool
	var false_edge_QMARK__114 bool
	var v277 vm.Value
	var v161 vm.Value
	var false_edge_QMARK__158 bool
	var v260 vm.Value
	var v193 vm.Value
	var v244 vm.Value
	var v228 vm.Value
	var callErr error
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{arg0, arg4})
	if callErr != nil {
		return nil, callErr
	}
	v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v20) {
		arg_type = arg3
		goto b1
	} else {
		target_bid = arg1
		arg_id = arg2
		arg_type = arg3
		f = arg4
		goto b2
	}
b1:
	;
	v304 = arg_type
	goto b3
b2:
	;
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{op, vm.Keyword("branch-if")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b7
	} else {
		goto b8
	}
b3:
	;
	return v304, nil
b4:
	;
	v293 = arg_type
	goto b6
b5:
	;
	cond_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	tt, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	ft, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__29178, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	true_edge_QMARK__97 = target_bid == arg__29178
	arg__29183, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	false_edge_QMARK__100 = target_bid == arg__29183
	v129 = arg_id == cond_id
	if v129 {
		true_edge_QMARK__113 = true_edge_QMARK__97
		false_edge_QMARK__114 = false_edge_QMARK__100
		goto b10
	} else {
		goto b11
	}
b6:
	;
	v304 = v293
	goto b3
b7:
	;
	v76 = or__x
	goto b9
b8:
	;
	v74, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v76 = v74
	goto b9
b9:
	;
	if vm.IsTruthy(v76) {
		goto b4
	} else {
		goto b5
	}
b10:
	;
	if true_edge_QMARK__113 {
		goto b13
	} else {
		false_edge_QMARK__158 = false_edge_QMARK__114
		goto b14
	}
b11:
	;
	v277 = arg_type
	goto b12
b12:
	;
	v293 = v277
	goto b6
b13:
	;
	v161, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "refine-truthy").Deref(), []vm.Value{arg_type})
	if callErr != nil {
		return nil, callErr
	}
	v260 = v161
	goto b15
b14:
	;
	if false_edge_QMARK__158 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v277 = v260
	goto b12
b16:
	;
	v193, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "refine-falsey").Deref(), []vm.Value{arg_type})
	if callErr != nil {
		return nil, callErr
	}
	v244 = v193
	goto b18
b17:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v260 = v244
	goto b15
b19:
	;
	v228 = arg_type
	goto b21
b20:
	;
	v228 = vm.NIL
	goto b21
b21:
	;
	v244 = v228
	goto b18
}
func target_arg_types(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var term vm.Value
	var v17 vm.Value
	var pred_bid vm.Value
	var target_bid vm.Value
	var param_pos vm.Value
	var f vm.Value
	var op vm.Value
	var aux vm.Value
	var v42 bool
	var v282 vm.Value
	var arg__29217 vm.Value
	var v62 bool
	var case__29192 vm.Value
	var v106 bool
	var v272 vm.Value
	var arg__29227 vm.Value
	var arg_id vm.Value
	var arg_type vm.Value
	var arg__29238 vm.Value
	var v75 vm.Value
	var v79 vm.Value
	var tt vm.Value
	var ft vm.Value
	var arg__29251 vm.Value
	var v134 bool
	var v262 vm.Value
	var arg__29261 vm.Value
	var arg__29293 vm.Value
	var arg__29294 vm.Value
	var v151 vm.Value
	var tt_types vm.Value
	var arg__29299 vm.Value
	var v190 bool
	var arg__29309 vm.Value
	var arg__29341 vm.Value
	var arg__29342 vm.Value
	var v207 vm.Value
	var ft_types vm.Value
	var arg__29355 vm.Value
	var v228 vm.Value
	var callErr error
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{arg0, arg3})
	if callErr != nil {
		return nil, callErr
	}
	v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v17) {
		goto b1
	} else {
		pred_bid = arg0
		target_bid = arg1
		param_pos = arg2
		f = arg3
		goto b2
	}
b1:
	;
	v282 = vm.NewArrayVector([]vm.Value{})
	goto b3
b2:
	;
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	v42 = op == vm.Keyword("branch")
	if v42 {
		goto b4
	} else {
		case__29192 = op
		goto b5
	}
b3:
	;
	return v282, nil
b4:
	;
	arg__29217, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v62 = target_bid == arg__29217
	if v62 {
		goto b7
	} else {
		goto b8
	}
b5:
	;
	v106 = case__29192 == vm.Keyword("branch-if")
	if v106 {
		goto b10
	} else {
		goto b11
	}
b6:
	;
	v282 = v272
	goto b3
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__29227, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__29227, param_pos})
	if callErr != nil {
		return nil, callErr
	}
	arg_type, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg_id, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__29238, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg_type})
	if callErr != nil {
		return nil, callErr
	}
	v75, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__29238})
	if callErr != nil {
		return nil, callErr
	}
	v79 = v75
	goto b9
b8:
	;
	v79 = vm.NewArrayVector([]vm.Value{})
	goto b9
b9:
	;
	v272 = v79
	goto b6
b10:
	;
	tt, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	ft, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__29251, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	v134 = target_bid == arg__29251
	if v134 {
		goto b13
	} else {
		goto b14
	}
b11:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b19
	} else {
		goto b20
	}
b12:
	;
	v272 = v262
	goto b6
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	arg__29261, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	arg_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__29261, param_pos})
	if callErr != nil {
		return nil, callErr
	}
	arg_type, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg_id, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "refine-edge-arg-type").Deref(), []vm.Value{pred_bid, target_bid, arg_id, arg_type, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__29293, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "refine-edge-arg-type").Deref(), []vm.Value{pred_bid, target_bid, arg_id, arg_type, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__29294, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__29293})
	if callErr != nil {
		return nil, callErr
	}
	v151, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__29294})
	if callErr != nil {
		return nil, callErr
	}
	tt_types = v151
	goto b15
b14:
	;
	tt_types = vm.NewArrayVector([]vm.Value{})
	goto b15
b15:
	;
	arg__29299, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	v190 = target_bid == arg__29299
	if v190 {
		goto b16
	} else {
		goto b17
	}
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	arg__29309, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	arg_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__29309, param_pos})
	if callErr != nil {
		return nil, callErr
	}
	arg_type, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg_id, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "refine-edge-arg-type").Deref(), []vm.Value{pred_bid, target_bid, arg_id, arg_type, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__29341, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "refine-edge-arg-type").Deref(), []vm.Value{pred_bid, target_bid, arg_id, arg_type, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__29342, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__29341})
	if callErr != nil {
		return nil, callErr
	}
	v207, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__29342})
	if callErr != nil {
		return nil, callErr
	}
	ft_types = v207
	goto b18
b17:
	;
	ft_types = vm.NewArrayVector([]vm.Value{})
	goto b18
b18:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{tt_types, ft_types})
	if callErr != nil {
		return nil, callErr
	}
	arg__29355, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{tt_types, ft_types})
	if callErr != nil {
		return nil, callErr
	}
	v228, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__29355})
	if callErr != nil {
		return nil, callErr
	}
	v262 = v228
	goto b12
b19:
	;
	goto b21
b20:
	;
	goto b21
b21:
	;
	v262 = vm.NewArrayVector([]vm.Value{})
	goto b12
}
func param_has_ready_source_QMARK_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var arg__29402 vm.Value
	var arg__29406 vm.Value
	var arg__29455 vm.Value
	var arg__29459 vm.Value
	var arg__29460 vm.Value
	var v31 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(vm.Keyword("param-sources"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__29402, callErr = rt.InvokeValue(vm.Keyword("param-sources"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__29406, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__29402, arg__29406})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("param-sources"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__29455, callErr = rt.InvokeValue(vm.Keyword("param-sources"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__29459, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__29460, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__29455, arg__29459})
	if callErr != nil {
		return nil, callErr
	}
	v31, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__29423 vm.Value
		var arg__29441 vm.Value
		var arg__29443 vm.Value
		var v25 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(2)})
		if callErr != nil {
			return nil, callErr
		}
		arg__29423, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(2)})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg__29423, arg3})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(2)})
		if callErr != nil {
			return nil, callErr
		}
		arg__29441, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(2)})
		if callErr != nil {
			return nil, callErr
		}
		arg__29443, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg__29441, arg3})
		if callErr != nil {
			return nil, callErr
		}
		v25, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("bottom"), arg__29443})
		if callErr != nil {
			return nil, callErr
		}
		return v25, nil
	}), arg__29460})
	if callErr != nil {
		return nil, callErr
	}
	return v31, nil
}
func enqueue_entry(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var v14 vm.Value
	var queue vm.Value
	var queued vm.Value
	var v17 vm.Value
	var entry vm.Value
	var arg__29481 vm.Value
	var arg__29487 vm.Value
	var v28 vm.Value
	var v30 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "ti-inc!").Deref(), []vm.Value{vm.Keyword("enqueue-attempt")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v14) {
		queue = arg0
		queued = arg1
		goto b1
	} else {
		queue = arg0
		queued = arg1
		entry = arg2
		goto b2
	}
b1:
	;
	v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{queue, queued})
	if callErr != nil {
		return nil, callErr
	}
	v30 = v17
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "ti-inc!").Deref(), []vm.Value{vm.Keyword("enqueue")})
	if callErr != nil {
		return nil, callErr
	}
	arg__29481, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{queue, entry})
	if callErr != nil {
		return nil, callErr
	}
	arg__29487, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{queued, entry})
	if callErr != nil {
		return nil, callErr
	}
	v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__29481, arg__29487})
	if callErr != nil {
		return nil, callErr
	}
	v30 = v28
	goto b3
b3:
	;
	return v30, nil
}
func join_all(arg0 vm.Value) (vm.Value, error) {
	var v6 vm.Value
	var callErr error
	v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.LookupVar("ir.passes.typeinfer", "type-join").Deref(), vm.Keyword("bottom"), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v6, nil
}
func numeric_type_QMARK_(arg0 vm.Value) (bool, error) {
	var or__x_2 bool
	var or__x_4 bool
	var t vm.Value
	var or__x_10 bool
	var v56 bool
	var or__x_12 bool
	var or__x_18 bool
	var v52 bool
	var or__x_20 bool
	var arg__29506 vm.Value
	var or__x_30 bool
	var v48 bool
	var or__x_32 bool
	var arg__29512 vm.Value
	var v42 bool
	var v44 bool
	var callErr error
	or__x_2 = arg0 == vm.Keyword("int")
	if or__x_2 {
		or__x_4 = or__x_2
		goto b1
	} else {
		t = arg0
		goto b2
	}
b1:
	;
	v56 = or__x_4
	goto b3
b2:
	;
	or__x_10 = t == vm.Keyword("float")
	if or__x_10 {
		or__x_12 = or__x_10
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v56, nil
b4:
	;
	v52 = or__x_12
	goto b6
b5:
	;
	or__x_18 = t == vm.Keyword("number")
	if or__x_18 {
		or__x_20 = or__x_18
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v56 = v52
	goto b3
b7:
	;
	v48 = or__x_20
	goto b9
b8:
	;
	arg__29506, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return false, callErr
	}
	or__x_30 = t == arg__29506
	if or__x_30 {
		or__x_32 = or__x_30
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v52 = v48
	goto b6
b10:
	;
	v44 = or__x_32
	goto b12
b11:
	;
	arg__29512, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return false, callErr
	}
	v42 = t == arg__29512
	v44 = v42
	goto b12
b12:
	;
	v48 = v44
	goto b9
}
func type_join(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v10 bool
	var a vm.Value
	var b vm.Value
	var v18 bool
	var v675 vm.Value
	var v26 bool
	var v671 vm.Value
	var v34 bool
	var v667 vm.Value
	var v43 bool
	var v663 vm.Value
	var v52 bool
	var v659 vm.Value
	var v61 bool
	var v655 vm.Value
	var and__x_70 bool
	var v651 vm.Value
	var and__x_98 bool
	var v647 vm.Value
	var arg__29537 vm.Value
	var v83 bool
	var and__x_76 bool
	var v86 bool
	var and__x_126 bool
	var v643 vm.Value
	var arg__29545 vm.Value
	var v111 bool
	var and__x_104 bool
	var v114 bool
	var and__x_154 bool
	var v639 vm.Value
	var arg__29553 vm.Value
	var v139 bool
	var and__x_132 bool
	var v142 bool
	var and__x_182 bool
	var v635 vm.Value
	var arg__29561 vm.Value
	var v167 bool
	var and__x_160 bool
	var v170 bool
	var and__x_224 bool
	var v631 vm.Value
	var or__x_191 bool
	var and__x_188 bool
	var v212 bool
	var or__x_195 bool
	var v203 bool
	var v205 bool
	var v266 bool
	var v627 vm.Value
	var or__x_233 bool
	var and__x_230 bool
	var v254 bool
	var or__x_237 bool
	var v245 bool
	var v247 bool
	var or__x_273 bool
	var v363 bool
	var v623 vm.Value
	var arg__29604 vm.Value
	var v352 vm.Value
	var v354 vm.Value
	var or__x_276 bool
	var or__x_283 bool
	var v338 bool
	var or__x_286 bool
	var or__x_293 bool
	var v333 bool
	var or__x_296 bool
	var arg__29587 vm.Value
	var or__x_307 bool
	var v328 bool
	var or__x_310 bool
	var arg__29593 vm.Value
	var v321 bool
	var v323 bool
	var or__x_370 bool
	var or__x_460 bool
	var v619 vm.Value
	var arg__29635 vm.Value
	var v449 vm.Value
	var v451 vm.Value
	var or__x_373 bool
	var or__x_380 bool
	var v435 bool
	var or__x_383 bool
	var or__x_390 bool
	var v430 bool
	var or__x_393 bool
	var arg__29618 vm.Value
	var or__x_404 bool
	var v425 bool
	var or__x_407 bool
	var arg__29624 vm.Value
	var v418 bool
	var v420 bool
	var v615 vm.Value
	var or__x_463 bool
	var arg__29643 vm.Value
	var v474 bool
	var and__x_476 bool
	var or__x_488 bool
	var and__x_483 bool
	var or__x_513 bool
	var or__x_492 bool
	var arg__29651 vm.Value
	var v504 bool
	var v506 bool
	var or__x_517 bool
	var or__x_526 bool
	var v587 bool
	var or__x_529 bool
	var arg__29659 vm.Value
	var v540 bool
	var and__x_542 bool
	var or__x_556 bool
	var and__x_551 bool
	var v581 bool
	var or__x_560 bool
	var arg__29667 vm.Value
	var v572 bool
	var v574 bool
	var arg__29678 vm.Value
	var v607 vm.Value
	var v611 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "ti-inc!").Deref(), []vm.Value{vm.Keyword("type-join")})
	if callErr != nil {
		return nil, callErr
	}
	v10 = arg0 == arg1
	if v10 {
		a = arg0
		goto b1
	} else {
		a = arg0
		b = arg1
		goto b2
	}
b1:
	;
	v675 = a
	goto b3
b2:
	;
	v18 = a == vm.Keyword("bottom")
	if v18 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v675, nil
b4:
	;
	v671 = b
	goto b6
b5:
	;
	v26 = b == vm.Keyword("bottom")
	if v26 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v675 = v671
	goto b3
b7:
	;
	v667 = a
	goto b9
b8:
	;
	v34 = a == vm.Keyword("any")
	if v34 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v671 = v667
	goto b6
b10:
	;
	v663 = vm.Keyword("any")
	goto b12
b11:
	;
	v43 = b == vm.Keyword("any")
	if v43 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v667 = v663
	goto b9
b13:
	;
	v659 = vm.Keyword("any")
	goto b15
b14:
	;
	v52 = a == vm.Keyword("unknown")
	if v52 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v663 = v659
	goto b12
b16:
	;
	v655 = vm.Keyword("unknown")
	goto b18
b17:
	;
	v61 = b == vm.Keyword("unknown")
	if v61 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v659 = v655
	goto b15
b19:
	;
	v651 = vm.Keyword("unknown")
	goto b21
b20:
	;
	and__x_70 = a == vm.Keyword("int")
	if and__x_70 {
		goto b25
	} else {
		and__x_76 = and__x_70
		goto b26
	}
b21:
	;
	v655 = v651
	goto b18
b22:
	;
	v647 = vm.Keyword("int")
	goto b24
b23:
	;
	and__x_98 = b == vm.Keyword("int")
	if and__x_98 {
		goto b31
	} else {
		and__x_104 = and__x_98
		goto b32
	}
b24:
	;
	v651 = v647
	goto b21
b25:
	;
	arg__29537, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v83 = b == arg__29537
	v86 = v83
	goto b27
b26:
	;
	v86 = and__x_76
	goto b27
b27:
	;
	if v86 {
		goto b22
	} else {
		goto b23
	}
b28:
	;
	v643 = vm.Keyword("int")
	goto b30
b29:
	;
	and__x_126 = a == vm.Keyword("float")
	if and__x_126 {
		goto b37
	} else {
		and__x_132 = and__x_126
		goto b38
	}
b30:
	;
	v647 = v643
	goto b24
b31:
	;
	arg__29545, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v111 = a == arg__29545
	v114 = v111
	goto b33
b32:
	;
	v114 = and__x_104
	goto b33
b33:
	;
	if v114 {
		goto b28
	} else {
		goto b29
	}
b34:
	;
	v639 = vm.Keyword("float")
	goto b36
b35:
	;
	and__x_154 = b == vm.Keyword("float")
	if and__x_154 {
		goto b43
	} else {
		and__x_160 = and__x_154
		goto b44
	}
b36:
	;
	v643 = v639
	goto b30
b37:
	;
	arg__29553, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v139 = b == arg__29553
	v142 = v139
	goto b39
b38:
	;
	v142 = and__x_132
	goto b39
b39:
	;
	if v142 {
		goto b34
	} else {
		goto b35
	}
b40:
	;
	v635 = vm.Keyword("float")
	goto b42
b41:
	;
	and__x_182 = a == vm.Keyword("bool")
	if and__x_182 {
		goto b49
	} else {
		and__x_188 = and__x_182
		goto b50
	}
b42:
	;
	v639 = v635
	goto b36
b43:
	;
	arg__29561, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v167 = a == arg__29561
	v170 = v167
	goto b45
b44:
	;
	v170 = and__x_160
	goto b45
b45:
	;
	if v170 {
		goto b40
	} else {
		goto b41
	}
b46:
	;
	v631 = vm.Keyword("bool")
	goto b48
b47:
	;
	and__x_224 = b == vm.Keyword("bool")
	if and__x_224 {
		goto b58
	} else {
		and__x_230 = and__x_224
		goto b59
	}
b48:
	;
	v635 = v631
	goto b42
b49:
	;
	or__x_191 = b == vm.Keyword("true")
	if or__x_191 {
		or__x_195 = or__x_191
		goto b52
	} else {
		goto b53
	}
b50:
	;
	v212 = and__x_188
	goto b51
b51:
	;
	if v212 {
		goto b46
	} else {
		goto b47
	}
b52:
	;
	v205 = or__x_195
	goto b54
b53:
	;
	v203 = b == vm.Keyword("false")
	v205 = v203
	goto b54
b54:
	;
	v212 = v205
	goto b51
b55:
	;
	v627 = vm.Keyword("bool")
	goto b57
b56:
	;
	v266 = a == vm.Keyword("number")
	if v266 {
		goto b64
	} else {
		goto b65
	}
b57:
	;
	v631 = v627
	goto b48
b58:
	;
	or__x_233 = a == vm.Keyword("true")
	if or__x_233 {
		or__x_237 = or__x_233
		goto b61
	} else {
		goto b62
	}
b59:
	;
	v254 = and__x_230
	goto b60
b60:
	;
	if v254 {
		goto b55
	} else {
		goto b56
	}
b61:
	;
	v247 = or__x_237
	goto b63
b62:
	;
	v245 = a == vm.Keyword("false")
	v247 = v245
	goto b63
b63:
	;
	v254 = v247
	goto b60
b64:
	;
	or__x_273 = b == vm.Keyword("int")
	if or__x_273 {
		or__x_276 = or__x_273
		goto b70
	} else {
		goto b71
	}
b65:
	;
	v363 = b == vm.Keyword("number")
	if v363 {
		goto b82
	} else {
		goto b83
	}
b66:
	;
	v627 = v623
	goto b57
b67:
	;
	v354 = vm.Keyword("number")
	goto b69
b68:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union"), a, b})
	if callErr != nil {
		return nil, callErr
	}
	arg__29604, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union"), a, b})
	if callErr != nil {
		return nil, callErr
	}
	v352, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__29604})
	if callErr != nil {
		return nil, callErr
	}
	v354 = v352
	goto b69
b69:
	;
	v623 = v354
	goto b66
b70:
	;
	v338 = or__x_276
	goto b72
b71:
	;
	or__x_283 = b == vm.Keyword("float")
	if or__x_283 {
		or__x_286 = or__x_283
		goto b73
	} else {
		goto b74
	}
b72:
	;
	if v338 {
		goto b67
	} else {
		goto b68
	}
b73:
	;
	v333 = or__x_286
	goto b75
b74:
	;
	or__x_293 = b == vm.Keyword("number")
	if or__x_293 {
		or__x_296 = or__x_293
		goto b76
	} else {
		goto b77
	}
b75:
	;
	v338 = v333
	goto b72
b76:
	;
	v328 = or__x_296
	goto b78
b77:
	;
	arg__29587, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	or__x_307 = b == arg__29587
	if or__x_307 {
		or__x_310 = or__x_307
		goto b79
	} else {
		goto b80
	}
b78:
	;
	v333 = v328
	goto b75
b79:
	;
	v323 = or__x_310
	goto b81
b80:
	;
	arg__29593, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v321 = b == arg__29593
	v323 = v321
	goto b81
b81:
	;
	v328 = v323
	goto b78
b82:
	;
	or__x_370 = a == vm.Keyword("int")
	if or__x_370 {
		or__x_373 = or__x_370
		goto b88
	} else {
		goto b89
	}
b83:
	;
	or__x_460 = a == vm.Keyword("int")
	if or__x_460 {
		or__x_463 = or__x_460
		goto b103
	} else {
		goto b104
	}
b84:
	;
	v623 = v619
	goto b66
b85:
	;
	v451 = vm.Keyword("number")
	goto b87
b86:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union"), a, b})
	if callErr != nil {
		return nil, callErr
	}
	arg__29635, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union"), a, b})
	if callErr != nil {
		return nil, callErr
	}
	v449, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__29635})
	if callErr != nil {
		return nil, callErr
	}
	v451 = v449
	goto b87
b87:
	;
	v619 = v451
	goto b84
b88:
	;
	v435 = or__x_373
	goto b90
b89:
	;
	or__x_380 = a == vm.Keyword("float")
	if or__x_380 {
		or__x_383 = or__x_380
		goto b91
	} else {
		goto b92
	}
b90:
	;
	if v435 {
		goto b85
	} else {
		goto b86
	}
b91:
	;
	v430 = or__x_383
	goto b93
b92:
	;
	or__x_390 = a == vm.Keyword("number")
	if or__x_390 {
		or__x_393 = or__x_390
		goto b94
	} else {
		goto b95
	}
b93:
	;
	v435 = v430
	goto b90
b94:
	;
	v425 = or__x_393
	goto b96
b95:
	;
	arg__29618, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	or__x_404 = a == arg__29618
	if or__x_404 {
		or__x_407 = or__x_404
		goto b97
	} else {
		goto b98
	}
b96:
	;
	v430 = v425
	goto b93
b97:
	;
	v420 = or__x_407
	goto b99
b98:
	;
	arg__29624, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v418 = a == arg__29624
	v420 = v418
	goto b99
b99:
	;
	v425 = v420
	goto b96
b100:
	;
	v615 = vm.Keyword("number")
	goto b102
b101:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b124
	} else {
		goto b125
	}
b102:
	;
	v619 = v615
	goto b84
b103:
	;
	and__x_476 = or__x_463
	goto b105
b104:
	;
	arg__29643, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v474 = a == arg__29643
	and__x_476 = v474
	goto b105
b105:
	;
	if and__x_476 {
		goto b106
	} else {
		and__x_483 = and__x_476
		goto b107
	}
b106:
	;
	or__x_488 = b == vm.Keyword("float")
	if or__x_488 {
		or__x_492 = or__x_488
		goto b109
	} else {
		goto b110
	}
b107:
	;
	or__x_513 = and__x_483
	goto b108
b108:
	;
	if or__x_513 {
		or__x_517 = or__x_513
		goto b112
	} else {
		goto b113
	}
b109:
	;
	v506 = or__x_492
	goto b111
b110:
	;
	arg__29651, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v504 = b == arg__29651
	v506 = v504
	goto b111
b111:
	;
	or__x_513 = v506
	goto b108
b112:
	;
	v587 = or__x_517
	goto b114
b113:
	;
	or__x_526 = a == vm.Keyword("float")
	if or__x_526 {
		or__x_529 = or__x_526
		goto b115
	} else {
		goto b116
	}
b114:
	;
	if v587 {
		goto b100
	} else {
		goto b101
	}
b115:
	;
	and__x_542 = or__x_529
	goto b117
b116:
	;
	arg__29659, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v540 = a == arg__29659
	and__x_542 = v540
	goto b117
b117:
	;
	if and__x_542 {
		goto b118
	} else {
		and__x_551 = and__x_542
		goto b119
	}
b118:
	;
	or__x_556 = b == vm.Keyword("int")
	if or__x_556 {
		or__x_560 = or__x_556
		goto b121
	} else {
		goto b122
	}
b119:
	;
	v581 = and__x_551
	goto b120
b120:
	;
	v587 = v581
	goto b114
b121:
	;
	v574 = or__x_560
	goto b123
b122:
	;
	arg__29667, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v572 = b == arg__29667
	v574 = v572
	goto b123
b123:
	;
	v581 = v574
	goto b120
b124:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union"), a, b})
	if callErr != nil {
		return nil, callErr
	}
	arg__29678, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("union"), a, b})
	if callErr != nil {
		return nil, callErr
	}
	v607, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__29678})
	if callErr != nil {
		return nil, callErr
	}
	v611 = v607
	goto b126
b125:
	;
	v611 = vm.NIL
	goto b126
b126:
	;
	v615 = v611
	goto b102
}
func set_type_if_changed_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var new_type vm.Value
	var old_type vm.Value
	var v16 bool
	var f vm.Value
	var inst vm.Value
	var v20 vm.Value
	var joined vm.Value
	var v37 bool
	var v45 vm.Value
	var callErr error
	new_type, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	old_type, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v16 = old_type == vm.Keyword("bottom")
	if v16 {
		f = arg0
		inst = arg1
		goto b1
	} else {
		f = arg0
		inst = arg1
		goto b2
	}
b1:
	;
	joined = new_type
	goto b3
b2:
	;
	v20, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "type-join").Deref(), []vm.Value{old_type, new_type})
	if callErr != nil {
		return nil, callErr
	}
	joined = v20
	goto b3
b3:
	;
	v37 = old_type == joined
	if v37 {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	v45 = vm.NIL
	goto b6
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "set-type!").Deref(), []vm.Value{f, inst, joined})
	if callErr != nil {
		return nil, callErr
	}
	v45 = vm.Boolean(true)
	goto b6
b6:
	;
	return v45, nil
}
func source_arg_type(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var pred_bid vm.Value
	var target_bid vm.Value
	var arg_id vm.Value
	var arg_type vm.Value
	var arg__29747 vm.Value
	var v21 vm.Value
	var callErr error
	pred_bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	target_bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	arg_type, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg_id, arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "refine-edge-arg-type").Deref(), []vm.Value{pred_bid, target_bid, arg_id, arg_type, arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__29747, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "refine-edge-arg-type").Deref(), []vm.Value{pred_bid, target_bid, arg_id, arg_type, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v21, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__29747})
	if callErr != nil {
		return nil, callErr
	}
	return v21, nil
}
func infer_block_param_from_deps(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var arg__29758 vm.Value
	var arg__29762 vm.Value
	var sources vm.Value
	var v25 vm.Value
	var f vm.Value
	var arg__29798 vm.Value
	var v44 vm.Value
	var v48 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(vm.Keyword("param-sources"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__29758, callErr = rt.InvokeValue(vm.Keyword("param-sources"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__29762, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	sources, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__29758, arg__29762})
	if callErr != nil {
		return nil, callErr
	}
	v25, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{sources})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v25) {
		f = arg3
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "source-arg-type").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), sources})
	if callErr != nil {
		return nil, callErr
	}
	arg__29798, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "source-arg-type").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), sources})
	if callErr != nil {
		return nil, callErr
	}
	v44, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "join-all").Deref(), []vm.Value{arg__29798})
	if callErr != nil {
		return nil, callErr
	}
	v48 = v44
	goto b3
b2:
	;
	v48 = vm.Keyword("bottom")
	goto b3
b3:
	;
	return v48, nil
}
func infer_block_param(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var preds vm.Value
	var v14 vm.Value
	var bid vm.Value
	var param_pos vm.Value
	var f vm.Value
	var arg__29855 vm.Value
	var v41 vm.Value
	var v45 vm.Value
	var callErr error
	preds, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{arg0, arg2})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{preds})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v14) {
		bid = arg0
		param_pos = arg1
		f = arg2
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v5 vm.Value
		var callErr error
		v5, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "target-arg-types").Deref(), []vm.Value{arg0, bid, param_pos, f})
		if callErr != nil {
			return nil, callErr
		}
		return v5, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	arg__29855, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v5 vm.Value
		var callErr error
		v5, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "target-arg-types").Deref(), []vm.Value{arg0, bid, param_pos, f})
		if callErr != nil {
			return nil, callErr
		}
		return v5, nil
	}), preds})
	if callErr != nil {
		return nil, callErr
	}
	v41, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "join-all").Deref(), []vm.Value{arg__29855})
	if callErr != nil {
		return nil, callErr
	}
	v45 = v41
	goto b3
b2:
	;
	v45 = vm.Keyword("bottom")
	goto b3
b3:
	;
	return v45, nil
}
func build_deps(arg0 vm.Value) (vm.Value, error) {
	var bu vm.Value
	var ps vm.Value
	var arg__29933 vm.Value
	var doseq_seq__29856 vm.Value
	var doseq_loop__29857 vm.Value
	var f vm.Value
	var bid vm.Value
	var term vm.Value
	var v63 vm.Value
	var arg__30030 vm.Value
	var arg__30035 vm.Value
	var arg__30040 vm.Value
	var v263 vm.Value
	var op vm.Value
	var aux vm.Value
	var v96 bool
	var v241 vm.Value
	var arg__29971 vm.Value
	var arg__29975 vm.Value
	var case__29858 vm.Value
	var v133 bool
	var tt vm.Value
	var ft vm.Value
	var arg__29998 vm.Value
	var arg__30002 vm.Value
	var arg__30017 vm.Value
	var arg__30021 vm.Value
	var callErr error
	bu, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	ps, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__29933, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__29856, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__29933})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__29857 = doseq_seq__29856
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__29857) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__29857})
	if callErr != nil {
		return nil, callErr
	}
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	v63, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v63) {
		goto b5
	} else {
		goto b6
	}
b3:
	;
	goto b4
b4:
	;
	arg__30030, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	arg__30035, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__30040, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{bu})
	if callErr != nil {
		return nil, callErr
	}
	v263, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("param-sources"), arg__30030, vm.Keyword("uses"), arg__30035, vm.Keyword("branch-arg-users"), arg__30040})
	if callErr != nil {
		return nil, callErr
	}
	return v263, nil
b5:
	;
	goto b7
b6:
	;
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	v96 = op == vm.Keyword("branch")
	if v96 {
		goto b8
	} else {
		case__29858 = op
		goto b9
	}
b7:
	;
	v241, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__29857})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__29857 = v241
	goto b1
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__29971, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__29975, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
		var as vm.Value
		var pos vm.Value
		var add_8 vm.Value
		var target vm.Value
		var ps_10 vm.Value
		var bu_11 vm.Value
		var pred vm.Value
		var v32 vm.Value
		var add_18 vm.Value
		var ps_20 vm.Value
		var bu_21 vm.Value
		var arg_id vm.Value
		var pp vm.Value
		var arg__29920 vm.Value
		var v45 vm.Value
		var v46 vm.Value
		var callErr error
		as = arg2
		pos = vm.Int(0)
		add_8 = rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
			var arg__29874 vm.Value
			var arg__29887 vm.Value
			var arg__29888 vm.Value
			var v22 vm.Value
			var callErr error
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29874, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "fnil").Deref(), []vm.Value{rt.LookupVar("clojure.core", "conj").Deref(), arg__29874})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29887, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29888, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "fnil").Deref(), []vm.Value{rt.LookupVar("clojure.core", "conj").Deref(), arg__29887})
			if callErr != nil {
				return nil, callErr
			}
			v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "update").Deref(), arg1, arg__29888, arg2})
			if callErr != nil {
				return nil, callErr
			}
			return v22, nil
		})
		target = arg1
		ps_10 = ps
		bu_11 = bu
		pred = arg0
		goto b1
	b1:
		;
		v32, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v32) {
			add_18 = add_8
			ps_20 = ps_10
			bu_21 = bu_11
			goto b2
		} else {
			goto b3
		}
	b2:
		;
		arg_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		pp, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{target, pos})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(add_18, []vm.Value{bu_21, arg_id, pp})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, target, arg_id})
		if callErr != nil {
			return nil, callErr
		}
		arg__29920, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, target, arg_id})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(add_18, []vm.Value{ps_20, pp, arg__29920})
		if callErr != nil {
			return nil, callErr
		}
		v45, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		v46 = rt.AddValue(pos, vm.Int(1))
		as = v45
		pos = v46
		add_8 = add_18
		ps_10 = ps_20
		bu_11 = bu_21
		goto b1
	b3:
		;
		goto b4
	b4:
		;
		return vm.NIL, nil
	}), []vm.Value{bid, arg__29971, arg__29975})
	if callErr != nil {
		return nil, callErr
	}
	goto b10
b9:
	;
	v133 = case__29858 == vm.Keyword("branch-if")
	if v133 {
		goto b11
	} else {
		goto b12
	}
b10:
	;
	goto b7
b11:
	;
	tt, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	ft, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	arg__29998, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	arg__30002, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
		var as vm.Value
		var pos vm.Value
		var add_8 vm.Value
		var target vm.Value
		var ps_10 vm.Value
		var bu_11 vm.Value
		var pred vm.Value
		var v32 vm.Value
		var add_18 vm.Value
		var ps_20 vm.Value
		var bu_21 vm.Value
		var arg_id vm.Value
		var pp vm.Value
		var arg__29920 vm.Value
		var v45 vm.Value
		var v46 vm.Value
		var callErr error
		as = arg2
		pos = vm.Int(0)
		add_8 = rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
			var arg__29874 vm.Value
			var arg__29887 vm.Value
			var arg__29888 vm.Value
			var v22 vm.Value
			var callErr error
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29874, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "fnil").Deref(), []vm.Value{rt.LookupVar("clojure.core", "conj").Deref(), arg__29874})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29887, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29888, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "fnil").Deref(), []vm.Value{rt.LookupVar("clojure.core", "conj").Deref(), arg__29887})
			if callErr != nil {
				return nil, callErr
			}
			v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "update").Deref(), arg1, arg__29888, arg2})
			if callErr != nil {
				return nil, callErr
			}
			return v22, nil
		})
		target = arg1
		ps_10 = ps
		bu_11 = bu
		pred = arg0
		goto b1
	b1:
		;
		v32, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v32) {
			add_18 = add_8
			ps_20 = ps_10
			bu_21 = bu_11
			goto b2
		} else {
			goto b3
		}
	b2:
		;
		arg_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		pp, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{target, pos})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(add_18, []vm.Value{bu_21, arg_id, pp})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, target, arg_id})
		if callErr != nil {
			return nil, callErr
		}
		arg__29920, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, target, arg_id})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(add_18, []vm.Value{ps_20, pp, arg__29920})
		if callErr != nil {
			return nil, callErr
		}
		v45, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		v46 = rt.AddValue(pos, vm.Int(1))
		as = v45
		pos = v46
		add_8 = add_18
		ps_10 = ps_20
		bu_11 = bu_21
		goto b1
	b3:
		;
		goto b4
	b4:
		;
		return vm.NIL, nil
	}), []vm.Value{bid, arg__29998, arg__30002})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	arg__30017, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	arg__30021, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
		var as vm.Value
		var pos vm.Value
		var add_8 vm.Value
		var target vm.Value
		var ps_10 vm.Value
		var bu_11 vm.Value
		var pred vm.Value
		var v32 vm.Value
		var add_18 vm.Value
		var ps_20 vm.Value
		var bu_21 vm.Value
		var arg_id vm.Value
		var pp vm.Value
		var arg__29920 vm.Value
		var v45 vm.Value
		var v46 vm.Value
		var callErr error
		as = arg2
		pos = vm.Int(0)
		add_8 = rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
			var arg__29874 vm.Value
			var arg__29887 vm.Value
			var arg__29888 vm.Value
			var v22 vm.Value
			var callErr error
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29874, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "fnil").Deref(), []vm.Value{rt.LookupVar("clojure.core", "conj").Deref(), arg__29874})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29887, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__29888, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "fnil").Deref(), []vm.Value{rt.LookupVar("clojure.core", "conj").Deref(), arg__29887})
			if callErr != nil {
				return nil, callErr
			}
			v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "update").Deref(), arg1, arg__29888, arg2})
			if callErr != nil {
				return nil, callErr
			}
			return v22, nil
		})
		target = arg1
		ps_10 = ps
		bu_11 = bu
		pred = arg0
		goto b1
	b1:
		;
		v32, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v32) {
			add_18 = add_8
			ps_20 = ps_10
			bu_21 = bu_11
			goto b2
		} else {
			goto b3
		}
	b2:
		;
		arg_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		pp, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{target, pos})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(add_18, []vm.Value{bu_21, arg_id, pp})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, target, arg_id})
		if callErr != nil {
			return nil, callErr
		}
		arg__29920, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, target, arg_id})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(add_18, []vm.Value{ps_20, pp, arg__29920})
		if callErr != nil {
			return nil, callErr
		}
		v45, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{as})
		if callErr != nil {
			return nil, callErr
		}
		v46 = rt.AddValue(pos, vm.Int(1))
		as = v45
		pos = v46
		add_8 = add_18
		ps_10 = ps_20
		bu_11 = bu_21
		goto b1
	b3:
		;
		goto b4
	b4:
		;
		return vm.NIL, nil
	}), []vm.Value{bid, arg__30017, arg__30021})
	if callErr != nil {
		return nil, callErr
	}
	goto b13
b12:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b14
	} else {
		goto b15
	}
b13:
	;
	goto b10
b14:
	;
	goto b16
b15:
	;
	goto b16
b16:
	;
	goto b13
}
func classify_const(arg0 vm.Value) (vm.Value, error) {
	var v5 vm.Value
	var aux vm.Value
	var v12 bool
	var v132 vm.Value
	var v19 bool
	var v129 vm.Value
	var and__x vm.Value
	var v126 vm.Value
	var v44 vm.Value
	var v49 vm.Value
	var v123 vm.Value
	var v33 bool
	var v36 vm.Value
	var v120 vm.Value
	var v74 vm.Value
	var v79 vm.Value
	var v117 vm.Value
	var v63 bool
	var v66 vm.Value
	var v86 vm.Value
	var v114 vm.Value
	var v93 vm.Value
	var v111 vm.Value
	var v108 vm.Value
	var v105 vm.Value
	var callErr error
	v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v5) {
		goto b1
	} else {
		aux = arg0
		goto b2
	}
b1:
	;
	v132 = vm.Keyword("nil")
	goto b3
b2:
	;
	v12 = aux == vm.Boolean(true)
	if v12 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v132, nil
b4:
	;
	v129 = vm.Keyword("true")
	goto b6
b5:
	;
	v19 = aux == vm.Boolean(false)
	if v19 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v132 = v129
	goto b3
b7:
	;
	v126 = vm.Keyword("false")
	goto b9
b8:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "integer?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b13
	} else {
		goto b14
	}
b9:
	;
	v129 = v126
	goto b6
b10:
	;
	v44, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v123 = v44
	goto b12
b11:
	;
	v49, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "integer?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v49) {
		goto b16
	} else {
		goto b17
	}
b12:
	;
	v126 = v123
	goto b9
b13:
	;
	v33 = aux == vm.Int(0)
	v36 = vm.Boolean(v33)
	goto b15
b14:
	;
	v36 = and__x
	goto b15
b15:
	;
	if vm.IsTruthy(v36) {
		goto b10
	} else {
		goto b11
	}
b16:
	;
	v120 = vm.Keyword("int")
	goto b18
b17:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "float?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b22
	} else {
		goto b23
	}
b18:
	;
	v123 = v120
	goto b12
b19:
	;
	v74, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v117 = v74
	goto b21
b20:
	;
	v79, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "float?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v79) {
		goto b25
	} else {
		goto b26
	}
b21:
	;
	v120 = v117
	goto b18
b22:
	;
	v63 = aux == vm.Float(0)
	v66 = vm.Boolean(v63)
	goto b24
b23:
	;
	v66 = and__x
	goto b24
b24:
	;
	if vm.IsTruthy(v66) {
		goto b19
	} else {
		goto b20
	}
b25:
	;
	v114 = vm.Keyword("float")
	goto b27
b26:
	;
	v86, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v86) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v117 = v114
	goto b21
b28:
	;
	v111 = vm.Keyword("string")
	goto b30
b29:
	;
	v93, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "char?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v93) {
		goto b31
	} else {
		goto b32
	}
b30:
	;
	v114 = v111
	goto b27
b31:
	;
	v108 = vm.Keyword("char")
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
	v111 = v108
	goto b30
b34:
	;
	v105 = vm.Keyword("any")
	goto b36
b35:
	;
	v105 = vm.NIL
	goto b36
b36:
	;
	v108 = v105
	goto b33
}
func arg_seed(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var stored vm.Value
	var arg_types vm.Value
	var idx vm.Value
	var and__x vm.Value
	var v50 vm.Value
	var v137 vm.Value
	var v38 vm.Value
	var v41 vm.Value
	var arg__30253 vm.Value
	var arg__30266 vm.Value
	var arg__30267 vm.Value
	var v105 vm.Value
	var v130 vm.Value
	var arg__30238 vm.Value
	var arg__30242 vm.Value
	var v81 bool
	var v84 vm.Value
	var v123 vm.Value
	var callErr error
	stored, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg_types, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arg-types").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	idx, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{stored, vm.Keyword("unknown")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b4
	} else {
		goto b5
	}
b1:
	;
	v50, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{stored})
	if callErr != nil {
		return nil, callErr
	}
	v137 = v50
	goto b3
b2:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{arg_types})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b10
	} else {
		goto b11
	}
b3:
	;
	return v137, nil
b4:
	;
	v38, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{stored, vm.Keyword("bottom")})
	if callErr != nil {
		return nil, callErr
	}
	v41 = v38
	goto b6
b5:
	;
	v41 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v41) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{idx})
	if callErr != nil {
		return nil, callErr
	}
	arg__30253, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{idx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg_types, arg__30253})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{idx})
	if callErr != nil {
		return nil, callErr
	}
	arg__30266, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{idx})
	if callErr != nil {
		return nil, callErr
	}
	arg__30267, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg_types, arg__30266})
	if callErr != nil {
		return nil, callErr
	}
	v105, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__30267})
	if callErr != nil {
		return nil, callErr
	}
	v130 = v105
	goto b9
b8:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b13
	} else {
		goto b14
	}
b9:
	;
	v137 = v130
	goto b3
b10:
	;
	arg__30238, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{idx})
	if callErr != nil {
		return nil, callErr
	}
	arg__30242, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg_types})
	if callErr != nil {
		return nil, callErr
	}
	v81 = rt.LtValue(arg__30238, arg__30242)
	v84 = vm.Boolean(v81)
	goto b12
b11:
	;
	v84 = and__x
	goto b12
b12:
	;
	if vm.IsTruthy(v84) {
		goto b7
	} else {
		goto b8
	}
b13:
	;
	v123 = vm.Keyword("unknown")
	goto b15
b14:
	;
	v123 = vm.NIL
	goto b15
b15:
	;
	v130 = v123
	goto b9
}
func propagate_changed_BANG_(arg0 int, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var uses vm.Value
	var arg__30277 vm.Value
	var v18 bool
	var v int
	var deps vm.Value
	var queue vm.Value
	var queued vm.Value
	var v21 vm.Value
	var us vm.Value
	var arg__30291 vm.Value
	var state vm.Value
	var vec__30268 vm.Value
	var arg__30426 vm.Value
	var v127 vm.Value
	var arg__30300 vm.Value
	var v73 vm.Value
	var and__x vm.Value
	var v76 vm.Value
	var params vm.Value
	var q vm.Value
	var s vm.Value
	var v146 vm.Value
	var arg__30438 vm.Value
	var arg__30448 vm.Value
	var arg__30449 vm.Value
	var vec__30270 vm.Value
	var q2 vm.Value
	var s2 vm.Value
	var v173 vm.Value
	var v176 vm.Value
	var v178 vm.Value
	var callErr error
	uses, callErr = rt.InvokeValue(vm.Keyword("uses"), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__30277, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{uses})
	if callErr != nil {
		return nil, callErr
	}
	v18 = rt.LtValue(vm.Int(arg0), arg__30277)
	if v18 {
		v = arg0
		deps = arg1
		queue = arg2
		queued = arg3
		goto b1
	} else {
		v = arg0
		deps = arg1
		queue = arg2
		queued = arg3
		goto b2
	}
b1:
	;
	v21, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, vm.Int(v)})
	if callErr != nil {
		return nil, callErr
	}
	us = v21
	goto b3
b2:
	;
	us = vm.NIL
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{queue, queued})
	if callErr != nil {
		return nil, callErr
	}
	arg__30291, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{queue, queued})
	if callErr != nil {
		return nil, callErr
	}
	state, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{arg__30291})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(us) {
		goto b7
	} else {
		and__x = us
		goto b8
	}
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-for-each").Deref(), []vm.Value{us, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var vec__30269 vm.Value
		var q vm.Value
		var s vm.Value
		var arg__30383 vm.Value
		var arg__30399 vm.Value
		var arg__30400 vm.Value
		var v33 vm.Value
		var callErr error
		vec__30269, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{state})
		if callErr != nil {
			return nil, callErr
		}
		q, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__30269, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		s, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__30269, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("inst"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__30383, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("inst"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "enqueue-entry").Deref(), []vm.Value{q, s, arg__30383})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("inst"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__30399, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("inst"), arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__30400, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "enqueue-entry").Deref(), []vm.Value{q, s, arg__30399})
		if callErr != nil {
			return nil, callErr
		}
		v33, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{state, arg__30400})
		if callErr != nil {
			return nil, callErr
		}
		return v33, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	goto b6
b5:
	;
	goto b6
b6:
	;
	vec__30268, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{state})
	if callErr != nil {
		return nil, callErr
	}
	queue, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__30268, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	queued, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__30268, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("branch-arg-users"), []vm.Value{deps})
	if callErr != nil {
		return nil, callErr
	}
	arg__30426, callErr = rt.InvokeValue(vm.Keyword("branch-arg-users"), []vm.Value{deps})
	if callErr != nil {
		return nil, callErr
	}
	v127, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__30426, vm.Int(v)})
	if callErr != nil {
		return nil, callErr
	}
	params = v127
	q = queue
	s = queued
	goto b10
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	arg__30300, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	v73, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__30300})
	if callErr != nil {
		return nil, callErr
	}
	v76 = v73
	goto b9
b8:
	;
	v76 = and__x
	goto b9
b9:
	;
	if vm.IsTruthy(v76) {
		goto b4
	} else {
		goto b5
	}
b10:
	;
	v146, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v146) {
		goto b11
	} else {
		goto b12
	}
b11:
	;
	arg__30438, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("param"), arg__30438})
	if callErr != nil {
		return nil, callErr
	}
	arg__30448, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	arg__30449, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("param"), arg__30448})
	if callErr != nil {
		return nil, callErr
	}
	vec__30270, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "enqueue-entry").Deref(), []vm.Value{q, s, arg__30449})
	if callErr != nil {
		return nil, callErr
	}
	q2, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__30270, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	s2, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__30270, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v173, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	params = v173
	q = q2
	s = s2
	goto b10
b12:
	;
	v176, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{q, s})
	if callErr != nil {
		return nil, callErr
	}
	v178 = v176
	goto b13
b13:
	;
	return v178, nil
}
func infer_one(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var op vm.Value
	var refs vm.Value
	var v15 bool
	var inst vm.Value
	var f vm.Value
	var arg__30494 vm.Value
	var v22 vm.Value
	var v33 bool
	var v225 vm.Value
	var v36 vm.Value
	var v47 bool
	var v219 vm.Value
	var v60 bool
	var v213 vm.Value
	var v73 bool
	var v207 vm.Value
	var v86 bool
	var v201 vm.Value
	var arg__30522 vm.Value
	var v93 vm.Value
	var v106 vm.Value
	var v195 vm.Value
	var v121 vm.Value
	var v189 vm.Value
	var v136 vm.Value
	var v183 vm.Value
	var arg__30570 vm.Value
	var v155 vm.Value
	var v177 vm.Value
	var v171 vm.Value
	var callErr error
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v15 = op == vm.Keyword("const")
	if v15 {
		inst = arg0
		f = arg1
		goto b1
	} else {
		inst = arg0
		f = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{inst, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__30494, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{inst, f})
	if callErr != nil {
		return nil, callErr
	}
	v22, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "classify-const").Deref(), []vm.Value{arg__30494})
	if callErr != nil {
		return nil, callErr
	}
	v225 = v22
	goto b3
b2:
	;
	v33 = op == vm.Keyword("load-arg")
	if v33 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v225, nil
b4:
	;
	v36, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "arg-seed").Deref(), []vm.Value{inst, f})
	if callErr != nil {
		return nil, callErr
	}
	v219 = v36
	goto b6
b5:
	;
	v47 = op == vm.Keyword("load-var")
	if v47 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v225 = v219
	goto b3
b7:
	;
	v213 = vm.Keyword("unknown")
	goto b9
b8:
	;
	v60 = op == vm.Keyword("load-closed")
	if v60 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v219 = v213
	goto b6
b10:
	;
	v207 = vm.Keyword("unknown")
	goto b12
b11:
	;
	v73 = op == vm.Keyword("call")
	if v73 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v213 = v207
	goto b9
b13:
	;
	v201 = vm.Keyword("unknown")
	goto b15
b14:
	;
	v86 = op == vm.Keyword("block-arg")
	if v86 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v207 = v201
	goto b12
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{inst, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__30522, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{inst, f})
	if callErr != nil {
		return nil, callErr
	}
	v93, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg__30522})
	if callErr != nil {
		return nil, callErr
	}
	v195 = v93
	goto b18
b17:
	;
	v106, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.passes.typeinfer", "terminator-ops").Deref(), op})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v106) {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v201 = v195
	goto b15
b19:
	;
	v189 = vm.Keyword("nil")
	goto b21
b20:
	;
	v121, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.passes.typeinfer", "comparison-ops").Deref(), op})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v121) {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v195 = v189
	goto b18
b22:
	;
	v183 = vm.Keyword("bool")
	goto b24
b23:
	;
	v136, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.passes.typeinfer", "numeric-ops").Deref(), op})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v136) {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v189 = v183
	goto b21
b25:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), refs})
	if callErr != nil {
		return nil, callErr
	}
	arg__30570, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), refs})
	if callErr != nil {
		return nil, callErr
	}
	v155, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "numeric-op-type").Deref(), []vm.Value{arg__30570})
	if callErr != nil {
		return nil, callErr
	}
	v177 = v155
	goto b27
b26:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v183 = v177
	goto b24
b28:
	;
	v171 = vm.Keyword("unknown")
	goto b30
b29:
	;
	v171 = vm.NIL
	goto b30
b30:
	;
	v177 = v171
	goto b27
}
func analyze_block_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var params vm.Value
	var ps vm.Value
	var idx int
	var changed_QMARK_ vm.Value
	var f vm.Value
	var bid vm.Value
	var v29 vm.Value
	var param_id vm.Value
	var param_type vm.Value
	var param_changed_QMARK_ vm.Value
	var v38 vm.Value
	var v39 int
	var changed_param_QMARK_ vm.Value
	var v88 vm.Value
	var or__x vm.Value
	var v63 vm.Value
	var insts vm.Value
	var v108 vm.Value
	var inst vm.Value
	var arg__31410 vm.Value
	var v119 vm.Value
	var changed_inst_QMARK_ vm.Value
	var term vm.Value
	var v191 vm.Value
	var arg__31438 vm.Value
	var v200 vm.Value
	var changed_term_QMARK_ vm.Value
	var v280 vm.Value
	var v266 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "ti-inc!").Deref(), []vm.Value{vm.Keyword("analyze-block")})
	if callErr != nil {
		return nil, callErr
	}
	params, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	ps = params
	idx = 0
	changed_QMARK_ = vm.Boolean(false)
	f = arg1
	bid = arg0
	goto b1
b1:
	;
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v29) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	param_id, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	param_type, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "infer-block-param").Deref(), []vm.Value{bid, vm.Int(idx), f})
	if callErr != nil {
		return nil, callErr
	}
	param_changed_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "set-type-if-changed!").Deref(), []vm.Value{f, param_id, param_type})
	if callErr != nil {
		return nil, callErr
	}
	v38, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	v39 = idx + 1
	if vm.IsTruthy(changed_QMARK_) {
		or__x = changed_QMARK_
		goto b5
	} else {
		goto b6
	}
b3:
	;
	changed_param_QMARK_ = changed_QMARK_
	goto b4
b4:
	;
	v88, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	insts = v88
	goto b8
b5:
	;
	v63 = or__x
	goto b7
b6:
	;
	v63 = param_changed_QMARK_
	goto b7
b7:
	;
	ps = v38
	idx = v39
	changed_QMARK_ = v63
	goto b1
b8:
	;
	v108, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v108) {
		goto b9
	} else {
		goto b10
	}
b9:
	;
	inst, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "infer-one").Deref(), []vm.Value{inst, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__31410, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "infer-one").Deref(), []vm.Value{inst, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "set-type-if-changed!").Deref(), []vm.Value{f, inst, arg__31410})
	if callErr != nil {
		return nil, callErr
	}
	v119, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(changed_QMARK_) {
		goto b12
	} else {
		goto b13
	}
b10:
	;
	changed_inst_QMARK_ = changed_QMARK_
	goto b11
b11:
	;
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	v191, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v191) {
		goto b15
	} else {
		goto b16
	}
b12:
	;
	goto b14
b13:
	;
	goto b14
b14:
	;
	insts = v119
	goto b8
b15:
	;
	changed_term_QMARK_ = vm.Boolean(false)
	goto b17
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "infer-one").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__31438, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "infer-one").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	v200, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "set-type-if-changed!").Deref(), []vm.Value{f, term, arg__31438})
	if callErr != nil {
		return nil, callErr
	}
	changed_term_QMARK_ = v200
	goto b17
b17:
	;
	if vm.IsTruthy(changed_param_QMARK_) {
		or__x = changed_param_QMARK_
		goto b18
	} else {
		goto b19
	}
b18:
	;
	v280 = or__x
	goto b20
b19:
	;
	if vm.IsTruthy(changed_inst_QMARK_) {
		or__x = changed_inst_QMARK_
		goto b21
	} else {
		goto b22
	}
b20:
	;
	return v280, nil
b21:
	;
	v266 = or__x
	goto b23
b22:
	;
	v266 = changed_term_QMARK_
	goto b23
b23:
	;
	v280 = v266
	goto b20
}
func falsey_type_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var t vm.Value
	var v6 bool
	var v13 bool
	var v48 vm.Value
	var v20 vm.Value
	var v45 vm.Value
	var arg__31459 vm.Value
	var v29 vm.Value
	var v42 vm.Value
	var v39 vm.Value
	var callErr error
	t, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "normalize-type").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v6 = t == vm.Keyword("nil")
	if v6 {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v48 = vm.Boolean(true)
	goto b3
b2:
	;
	v13 = t == vm.Keyword("false")
	if v13 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v48, nil
b4:
	;
	v45 = vm.Boolean(true)
	goto b6
b5:
	;
	v20, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "union-type?").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v20) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v48 = v45
	goto b3
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__31459, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("ir.passes.typeinfer", "falsey-type?").Deref(), arg__31459})
	if callErr != nil {
		return nil, callErr
	}
	v42 = v29
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
	v45 = v42
	goto b6
b10:
	;
	v39 = vm.Boolean(false)
	goto b12
b11:
	;
	v39 = vm.NIL
	goto b12
b12:
	;
	v42 = v39
	goto b9
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.typeinfer", map[string]vm.Value{"const-type?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("const-type?: wrong number of arguments %d (expected 1)", len(args))
		}
		return const_type_QMARK_(args[0])
	}), "sort-members": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("sort-members: wrong number of arguments %d (expected 1)", len(args))
		}
		return sort_members(args[0])
	}), "union-type?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("union-type?: wrong number of arguments %d (expected 1)", len(args))
		}
		return union_type_QMARK_(args[0])
	}), "ti-inc!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("ti-inc!: wrong number of arguments %d (expected 1)", len(args))
		}
		return ti_inc_BANG_(args[0])
	}), "refine-truthy": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("refine-truthy: wrong number of arguments %d (expected 1)", len(args))
		}
		return refine_truthy(args[0])
	}), "truthy-type?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("truthy-type?: wrong number of arguments %d (expected 1)", len(args))
		}
		return truthy_type_QMARK_(args[0])
	}), "refine-falsey": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("refine-falsey: wrong number of arguments %d (expected 1)", len(args))
		}
		return refine_falsey(args[0])
	}), "refine-edge-arg-type": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 5 {
			return nil, fmt.Errorf("refine-edge-arg-type: wrong number of arguments %d (expected 5)", len(args))
		}
		return refine_edge_arg_type(args[0], args[1], args[2], args[3], args[4])
	}), "target-arg-types": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("target-arg-types: wrong number of arguments %d (expected 4)", len(args))
		}
		return target_arg_types(args[0], args[1], args[2], args[3])
	}), "param-has-ready-source?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("param-has-ready-source?: wrong number of arguments %d (expected 4)", len(args))
		}
		return param_has_ready_source_QMARK_(args[0], args[1], args[2], args[3])
	}), "enqueue-entry": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("enqueue-entry: wrong number of arguments %d (expected 3)", len(args))
		}
		return enqueue_entry(args[0], args[1], args[2])
	}), "join-all": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("join-all: wrong number of arguments %d (expected 1)", len(args))
		}
		return join_all(args[0])
	}), "type-join": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("type-join: wrong number of arguments %d (expected 2)", len(args))
		}
		return type_join(args[0], args[1])
	}), "set-type-if-changed!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-type-if-changed!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_type_if_changed_BANG_(args[0], args[1], args[2])
	}), "source-arg-type": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("source-arg-type: wrong number of arguments %d (expected 2)", len(args))
		}
		return source_arg_type(args[0], args[1])
	}), "infer-block-param-from-deps": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("infer-block-param-from-deps: wrong number of arguments %d (expected 4)", len(args))
		}
		return infer_block_param_from_deps(args[0], args[1], args[2], args[3])
	}), "infer-block-param": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("infer-block-param: wrong number of arguments %d (expected 3)", len(args))
		}
		return infer_block_param(args[0], args[1], args[2])
	}), "build-deps": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("build-deps: wrong number of arguments %d (expected 1)", len(args))
		}
		return build_deps(args[0])
	}), "classify-const": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("classify-const: wrong number of arguments %d (expected 1)", len(args))
		}
		return classify_const(args[0])
	}), "arg-seed": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("arg-seed: wrong number of arguments %d (expected 2)", len(args))
		}
		return arg_seed(args[0], args[1])
	}), "infer-one": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("infer-one: wrong number of arguments %d (expected 2)", len(args))
		}
		return infer_one(args[0], args[1])
	}), "analyze-block!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("analyze-block!: wrong number of arguments %d (expected 2)", len(args))
		}
		return analyze_block_BANG_(args[0], args[1])
	}), "falsey-type?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("falsey-type?: wrong number of arguments %d (expected 1)", len(args))
		}
		return falsey_type_QMARK_(args[0])
	}),
	})
}
