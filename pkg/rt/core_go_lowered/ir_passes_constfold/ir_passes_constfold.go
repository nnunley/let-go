package ir_passes_constfold

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func const_QMARK_(arg0 vm.Value, arg1 vm.Value) (bool, error) {
	var arg__20420 vm.Value
	var v5 bool
	var callErr error
	arg__20420, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return false, callErr
	}
	v5 = arg__20420 == vm.Keyword("const")
	return v5, nil
}
func const_val(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v7 vm.Value
	var nid vm.Value
	var f vm.Value
	var v10 vm.Value
	var v14 vm.Value
	var callErr error
	v7, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "const?").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v7) {
		nid = arg0
		f = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v10, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v14 = v10
	goto b3
b2:
	;
	v14 = vm.NIL
	goto b3
b3:
	;
	return v14, nil
}
func apply_action_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var v11 vm.Value
	var f vm.Value
	var nid vm.Value
	var action vm.Value
	var map__20558 vm.Value
	var op vm.Value
	var refs vm.Value
	var aux vm.Value
	var v32 vm.Value
	var v41 vm.Value
	var v78 vm.Value
	var arg__20611 vm.Value
	var v48 vm.Value
	var v57 vm.Value
	var v73 vm.Value
	var arg__20624 vm.Value
	var v64 vm.Value
	var v68 vm.Value
	var callErr error
	v11, callErr = rt.InvokeValue(vm.Keyword("replace-with"), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v11) {
		f = arg0
		nid = arg1
		action = arg2
		goto b1
	} else {
		f = arg0
		nid = arg1
		action = arg2
		goto b2
	}
b1:
	;
	map__20558, callErr = rt.InvokeValue(vm.Keyword("replace-with"), []vm.Value{action})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{map__20558, vm.Keyword("op")})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{map__20558, vm.Keyword("refs")})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{map__20558, vm.Keyword("aux")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "replace-op!").Deref(), []vm.Value{f, nid, op})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "set-refs!").Deref(), []vm.Value{f, nid, refs})
	if callErr != nil {
		return nil, callErr
	}
	v32, callErr = rt.InvokeValue(rt.LookupVar("ir", "set-aux!").Deref(), []vm.Value{f, nid, aux})
	if callErr != nil {
		return nil, callErr
	}
	v78 = v32
	goto b3
b2:
	;
	v41, callErr = rt.InvokeValue(vm.Keyword("replace-uses"), []vm.Value{action})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v41) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v78, nil
b4:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("replace-uses"), []vm.Value{action})
	if callErr != nil {
		return nil, callErr
	}
	arg__20611, callErr = rt.InvokeValue(vm.Keyword("replace-uses"), []vm.Value{action})
	if callErr != nil {
		return nil, callErr
	}
	v48, callErr = rt.InvokeValue(rt.LookupVar("ir", "replace-all-uses!").Deref(), []vm.Value{f, nid, arg__20611})
	if callErr != nil {
		return nil, callErr
	}
	v73 = v48
	goto b6
b5:
	;
	v57, callErr = rt.InvokeValue(vm.Keyword("swap-refs"), []vm.Value{action})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v57) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v78 = v73
	goto b3
b7:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("swap-refs"), []vm.Value{action})
	if callErr != nil {
		return nil, callErr
	}
	arg__20624, callErr = rt.InvokeValue(vm.Keyword("swap-refs"), []vm.Value{action})
	if callErr != nil {
		return nil, callErr
	}
	v64, callErr = rt.InvokeValue(rt.LookupVar("ir", "set-refs!").Deref(), []vm.Value{f, nid, arg__20624})
	if callErr != nil {
		return nil, callErr
	}
	v68 = v64
	goto b9
b8:
	;
	v68 = vm.NIL
	goto b9
b9:
	;
	v73 = v68
	goto b6
}
func try_identity(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var op vm.Value
	var refs vm.Value
	var arg__20640 vm.Value
	var v18 bool
	var f vm.Value
	var r0 vm.Value
	var r1 vm.Value
	var v0 vm.Value
	var v1 vm.Value
	var v57 bool
	var v570 vm.Value
	var v81 vm.Value
	var case__20625 vm.Value
	var v205 bool
	var v555 vm.Value
	var v85 vm.Value
	var v109 vm.Value
	var v169 vm.Value
	var v113 vm.Value
	var v156 vm.Value
	var v229 vm.Value
	var v273 bool
	var v542 vm.Value
	var v233 vm.Value
	var v237 vm.Value
	var v297 vm.Value
	var v529 vm.Value
	var v301 vm.Value
	var v325 vm.Value
	var v475 vm.Value
	var v329 vm.Value
	var or__x vm.Value
	var v462 vm.Value
	var arg__20722 vm.Value
	var v406 vm.Value
	var v449 vm.Value
	var v380 vm.Value
	var v382 vm.Value
	var callErr error
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__20640, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v18 = arg__20640 == vm.Int(2)
	if v18 {
		f = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	r0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	r1, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v0, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "const-val").Deref(), []vm.Value{r0, f})
	if callErr != nil {
		return nil, callErr
	}
	v1, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "const-val").Deref(), []vm.Value{r1, f})
	if callErr != nil {
		return nil, callErr
	}
	v57 = op == vm.Keyword("add")
	if v57 {
		goto b4
	} else {
		case__20625 = op
		goto b5
	}
b2:
	;
	v570 = vm.NIL
	goto b3
b3:
	;
	return v570, nil
b4:
	;
	v81, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var v vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "number?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			v = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v9 = vm.Boolean(vm.Int(0) == v)
		v12 = v9
		goto b3
	b2:
		;
		v12 = and__x
		goto b3
	b3:
		;
		return v12, nil
	}), []vm.Value{v1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v81) {
		goto b7
	} else {
		goto b8
	}
b5:
	;
	v205 = case__20625 == vm.Keyword("sub")
	if v205 {
		goto b16
	} else {
		goto b17
	}
b6:
	;
	v570 = v555
	goto b3
b7:
	;
	v85, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("replace-uses"), r0})
	if callErr != nil {
		return nil, callErr
	}
	v169 = v85
	goto b9
b8:
	;
	v109, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var v vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "number?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			v = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v9 = vm.Boolean(vm.Int(0) == v)
		v12 = v9
		goto b3
	b2:
		;
		v12 = and__x
		goto b3
	b3:
		;
		return v12, nil
	}), []vm.Value{v0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v109) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v555 = v169
	goto b6
b10:
	;
	v113, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("replace-uses"), r1})
	if callErr != nil {
		return nil, callErr
	}
	v156 = v113
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
	v169 = v156
	goto b9
b13:
	;
	goto b15
b14:
	;
	goto b15
b15:
	;
	v156 = vm.NIL
	goto b12
b16:
	;
	v229, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var v vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "number?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			v = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v9 = vm.Boolean(vm.Int(0) == v)
		v12 = v9
		goto b3
	b2:
		;
		v12 = and__x
		goto b3
	b3:
		;
		return v12, nil
	}), []vm.Value{v1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v229) {
		goto b19
	} else {
		goto b20
	}
b17:
	;
	v273 = case__20625 == vm.Keyword("mul")
	if v273 {
		goto b22
	} else {
		goto b23
	}
b18:
	;
	v555 = v542
	goto b6
b19:
	;
	v233, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("replace-uses"), r0})
	if callErr != nil {
		return nil, callErr
	}
	v237 = v233
	goto b21
b20:
	;
	v237 = vm.NIL
	goto b21
b21:
	;
	v542 = v237
	goto b18
b22:
	;
	v297, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var v vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "number?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			v = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v9 = vm.Boolean(vm.Int(1) == v)
		v12 = v9
		goto b3
	b2:
		;
		v12 = and__x
		goto b3
	b3:
		;
		return v12, nil
	}), []vm.Value{v1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v297) {
		goto b25
	} else {
		goto b26
	}
b23:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b40
	} else {
		goto b41
	}
b24:
	;
	v542 = v529
	goto b18
b25:
	;
	v301, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("replace-uses"), r0})
	if callErr != nil {
		return nil, callErr
	}
	v475 = v301
	goto b27
b26:
	;
	v325, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var v vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "number?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			v = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v9 = vm.Boolean(vm.Int(1) == v)
		v12 = v9
		goto b3
	b2:
		;
		v12 = and__x
		goto b3
	b3:
		;
		return v12, nil
	}), []vm.Value{v0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v325) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v529 = v475
	goto b24
b28:
	;
	v329, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("replace-uses"), r1})
	if callErr != nil {
		return nil, callErr
	}
	v462 = v329
	goto b30
b29:
	;
	or__x, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var v vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "number?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			v = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v9 = vm.Boolean(vm.Int(0) == v)
		v12 = v9
		goto b3
	b2:
		;
		v12 = and__x
		goto b3
	b3:
		;
		return v12, nil
	}), []vm.Value{v0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b34
	} else {
		goto b35
	}
b30:
	;
	v475 = v462
	goto b27
b31:
	;
	arg__20722, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("op"), vm.Keyword("const"), vm.Keyword("aux"), vm.Int(0), vm.Keyword("refs"), vm.NewArrayVector([]vm.Value{})})
	if callErr != nil {
		return nil, callErr
	}
	v406, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("replace-with"), arg__20722})
	if callErr != nil {
		return nil, callErr
	}
	v449 = v406
	goto b33
b32:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b37
	} else {
		goto b38
	}
b33:
	;
	v462 = v449
	goto b30
b34:
	;
	v382 = or__x
	goto b36
b35:
	;
	v380, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var v vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "number?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			v = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v9 = vm.Boolean(vm.Int(0) == v)
		v12 = v9
		goto b3
	b2:
		;
		v12 = and__x
		goto b3
	b3:
		;
		return v12, nil
	}), []vm.Value{v1})
	if callErr != nil {
		return nil, callErr
	}
	v382 = v380
	goto b36
b36:
	;
	if vm.IsTruthy(v382) {
		goto b31
	} else {
		goto b32
	}
b37:
	;
	goto b39
b38:
	;
	goto b39
b39:
	;
	v449 = vm.NIL
	goto b33
b40:
	;
	goto b42
b41:
	;
	goto b42
b42:
	;
	v529 = vm.NIL
	goto b24
}
func try_canonicalize(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var op vm.Value
	var refs vm.Value
	var and__x_16 vm.Value
	var arg__20797 vm.Value
	var arg__20803 vm.Value
	var arg__20804 vm.Value
	var v123 vm.Value
	var v127 vm.Value
	var f vm.Value
	var arg__20740 vm.Value
	var and__x_31 bool
	var and__x_26 vm.Value
	var v104 vm.Value
	var arg__20754 vm.Value
	var and__x_52 vm.Value
	var and__x_41 bool
	var v96 vm.Value
	var arg__20769 vm.Value
	var arg__20786 vm.Value
	var arg__20788 vm.Value
	var v85 vm.Value
	var and__x_62 vm.Value
	var v88 vm.Value
	var callErr error
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	and__x_16, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "commutative").Deref(), []vm.Value{op})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x_16) {
		f = arg1
		goto b4
	} else {
		and__x_26 = and__x_16
		goto b5
	}
b1:
	;
	arg__20797, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__20803, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__20804, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__20797, arg__20803})
	if callErr != nil {
		return nil, callErr
	}
	v123, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("swap-refs"), arg__20804})
	if callErr != nil {
		return nil, callErr
	}
	v127 = v123
	goto b3
b2:
	;
	v127 = vm.NIL
	goto b3
b3:
	;
	return v127, nil
b4:
	;
	arg__20740, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	and__x_31 = arg__20740 == vm.Int(2)
	if and__x_31 {
		goto b7
	} else {
		and__x_41 = and__x_31
		goto b8
	}
b5:
	;
	v104 = and__x_26
	goto b6
b6:
	;
	if vm.IsTruthy(v104) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__20754, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	and__x_52, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "const?").Deref(), []vm.Value{arg__20754, f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x_52) {
		goto b10
	} else {
		and__x_62 = and__x_52
		goto b11
	}
b8:
	;
	v96 = vm.Boolean(and__x_41)
	goto b9
b9:
	;
	v104 = v96
	goto b6
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__20769, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "const?").Deref(), []vm.Value{arg__20769, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__20786, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__20788, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "const?").Deref(), []vm.Value{arg__20786, f})
	if callErr != nil {
		return nil, callErr
	}
	v85, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__20788})
	if callErr != nil {
		return nil, callErr
	}
	v88 = v85
	goto b12
b11:
	;
	v88 = and__x_62
	goto b12
b12:
	;
	v96 = v88
	goto b9
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.constfold", map[string]vm.Value{"const-val": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("const-val: wrong number of arguments %d (expected 2)", len(args))
		}
		return const_val(args[0], args[1])
	}), "apply-action!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("apply-action!: wrong number of arguments %d (expected 3)", len(args))
		}
		return apply_action_BANG_(args[0], args[1], args[2])
	}), "try-identity": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("try-identity: wrong number of arguments %d (expected 2)", len(args))
		}
		return try_identity(args[0], args[1])
	}), "try-canonicalize": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("try-canonicalize: wrong number of arguments %d (expected 2)", len(args))
		}
		return try_canonicalize(args[0], args[1])
	}),
	})
}
