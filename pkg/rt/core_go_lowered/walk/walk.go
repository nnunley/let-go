package walk

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func walk(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var v10 vm.Value
	var inner vm.Value
	var outer vm.Value
	var form vm.Value
	var arg__34930 vm.Value
	var arg__34947 vm.Value
	var arg__34948 vm.Value
	var v28 vm.Value
	var v37 vm.Value
	var v163 vm.Value
	var arg__34960 vm.Value
	var arg__34970 vm.Value
	var arg__34981 vm.Value
	var arg__34982 vm.Value
	var arg__34991 vm.Value
	var arg__34992 vm.Value
	var arg__35003 vm.Value
	var arg__35013 vm.Value
	var arg__35024 vm.Value
	var arg__35025 vm.Value
	var arg__35034 vm.Value
	var arg__35035 vm.Value
	var arg__35036 vm.Value
	var v83 vm.Value
	var v92 vm.Value
	var v158 vm.Value
	var arg__35052 vm.Value
	var v98 vm.Value
	var v107 vm.Value
	var v153 vm.Value
	var arg__35070 vm.Value
	var arg__35076 vm.Value
	var arg__35093 vm.Value
	var arg__35099 vm.Value
	var arg__35100 vm.Value
	var v129 vm.Value
	var v148 vm.Value
	var v139 vm.Value
	var v143 vm.Value
	var callErr error
	v10, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v10) {
		inner = arg0
		outer = arg1
		form = arg2
		goto b1
	} else {
		inner = arg0
		outer = arg1
		form = arg2
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34930, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), arg__34930})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34947, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34948, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), arg__34947})
	if callErr != nil {
		return nil, callErr
	}
	v28, callErr = rt.InvokeValue(outer, []vm.Value{arg__34948})
	if callErr != nil {
		return nil, callErr
	}
	v163 = v28
	goto b3
b2:
	;
	v37, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map-entry?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v37) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v163, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34960, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(inner, []vm.Value{arg__34960})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34970, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(inner, []vm.Value{arg__34970})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34981, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34982, callErr = rt.InvokeValue(inner, []vm.Value{arg__34981})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34991, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34992, callErr = rt.InvokeValue(inner, []vm.Value{arg__34991})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.lang.MapEntry", "create").Deref(), []vm.Value{arg__34982, arg__34992})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35003, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(inner, []vm.Value{arg__35003})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35013, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(inner, []vm.Value{arg__35013})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35024, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35025, callErr = rt.InvokeValue(inner, []vm.Value{arg__35024})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35034, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35035, callErr = rt.InvokeValue(inner, []vm.Value{arg__35034})
	if callErr != nil {
		return nil, callErr
	}
	arg__35036, callErr = rt.InvokeValue(rt.LookupVar("clojure.lang.MapEntry", "create").Deref(), []vm.Value{arg__35025, arg__35035})
	if callErr != nil {
		return nil, callErr
	}
	v83, callErr = rt.InvokeValue(outer, []vm.Value{arg__35036})
	if callErr != nil {
		return nil, callErr
	}
	v158 = v83
	goto b6
b5:
	;
	v92, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v92) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v163 = v158
	goto b3
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35052, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	v98, callErr = rt.InvokeValue(outer, []vm.Value{arg__35052})
	if callErr != nil {
		return nil, callErr
	}
	v153 = v98
	goto b9
b8:
	;
	v107, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "coll?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v107) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v158 = v153
	goto b6
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35070, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35076, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__35070, arg__35076})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35093, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35099, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{inner, form})
	if callErr != nil {
		return nil, callErr
	}
	arg__35100, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__35093, arg__35099})
	if callErr != nil {
		return nil, callErr
	}
	v129, callErr = rt.InvokeValue(outer, []vm.Value{arg__35100})
	if callErr != nil {
		return nil, callErr
	}
	v148 = v129
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
	v153 = v148
	goto b9
b13:
	;
	v139, callErr = rt.InvokeValue(outer, []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v143 = v139
	goto b15
b14:
	;
	v143 = vm.NIL
	goto b15
b15:
	;
	v148 = v143
	goto b12
}
func postwalk(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__35118 vm.Value
	var v11 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "partial").Deref(), []vm.Value{rt.LookupVar("clojure.walk", "postwalk").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__35118, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "partial").Deref(), []vm.Value{rt.LookupVar("clojure.walk", "postwalk").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.walk", "walk").Deref(), []vm.Value{arg__35118, arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v11, nil
}
func keywordize_keys(arg0 vm.Value) (vm.Value, error) {
	var v9 vm.Value
	var callErr error
	v9, callErr = rt.InvokeValue(rt.LookupVar("clojure.walk", "postwalk").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v7 vm.Value
		var x vm.Value
		var f_3 vm.Value
		var arg__35186 vm.Value
		var v16 vm.Value
		var v19 vm.Value
		var callErr error
		v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v7) {
			x = arg0
			f_3 = rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
				var k vm.Value
				var v vm.Value
				var v20 vm.Value
				var arg__35143 vm.Value
				var v25 vm.Value
				var v28 vm.Value
				var v30 vm.Value
				var callErr error
				k, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
				if callErr != nil {
					return nil, callErr
				}
				v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
				if callErr != nil {
					return nil, callErr
				}
				v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{k})
				if callErr != nil {
					return nil, callErr
				}
				if vm.IsTruthy(v20) {
					goto b1
				} else {
					goto b2
				}
			b1:
				;
				arg__35143, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{k})
				if callErr != nil {
					return nil, callErr
				}
				v25, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__35143, v})
				if callErr != nil {
					return nil, callErr
				}
				v30 = v25
				goto b3
			b2:
				;
				v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{k, v})
				if callErr != nil {
					return nil, callErr
				}
				v30 = v28
				goto b3
			b3:
				;
				return v30, nil
			})
			goto b1
		} else {
			x = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{f_3, x})
		if callErr != nil {
			return nil, callErr
		}
		arg__35186, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{f_3, x})
		if callErr != nil {
			return nil, callErr
		}
		v16, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{vm.EmptyPersistentMap, arg__35186})
		if callErr != nil {
			return nil, callErr
		}
		v19 = v16
		goto b3
	b2:
		;
		v19 = x
		goto b3
	b3:
		;
		return v19, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v9, nil
}
func prewalk(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__35206 vm.Value
	var arg__35211 vm.Value
	var v15 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "partial").Deref(), []vm.Value{rt.LookupVar("clojure.walk", "prewalk").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(arg0, []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__35206, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "partial").Deref(), []vm.Value{rt.LookupVar("clojure.walk", "prewalk").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__35211, callErr = rt.InvokeValue(arg0, []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	v15, callErr = rt.InvokeValue(rt.LookupVar("clojure.walk", "walk").Deref(), []vm.Value{arg__35206, rt.LookupVar("clojure.core", "identity").Deref(), arg__35211})
	if callErr != nil {
		return nil, callErr
	}
	return v15, nil
}
func prewalk_replace(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v9 vm.Value
	var callErr error
	v9, callErr = rt.InvokeValue(rt.LookupVar("clojure.walk", "prewalk").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v7 vm.Value
		var x vm.Value
		var smap_3 vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v7) {
			x = arg0
			smap_3 = arg0
			goto b1
		} else {
			x = arg0
			goto b2
		}
	b1:
		;
		v9, callErr = rt.InvokeValue(smap_3, []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		v12 = v9
		goto b3
	b2:
		;
		v12 = x
		goto b3
	b3:
		;
		return v12, nil
	}), arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v9, nil
}
func stringify_keys(arg0 vm.Value) (vm.Value, error) {
	var v9 vm.Value
	var callErr error
	v9, callErr = rt.InvokeValue(rt.LookupVar("clojure.walk", "postwalk").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v7 vm.Value
		var x vm.Value
		var f_3 vm.Value
		var arg__35298 vm.Value
		var v16 vm.Value
		var v19 vm.Value
		var callErr error
		v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v7) {
			x = arg0
			f_3 = rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
				var k vm.Value
				var v vm.Value
				var v20 vm.Value
				var arg__35255 vm.Value
				var v25 vm.Value
				var v28 vm.Value
				var v30 vm.Value
				var callErr error
				k, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
				if callErr != nil {
					return nil, callErr
				}
				v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
				if callErr != nil {
					return nil, callErr
				}
				v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{k})
				if callErr != nil {
					return nil, callErr
				}
				if vm.IsTruthy(v20) {
					goto b1
				} else {
					goto b2
				}
			b1:
				;
				arg__35255, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{k})
				if callErr != nil {
					return nil, callErr
				}
				v25, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__35255, v})
				if callErr != nil {
					return nil, callErr
				}
				v30 = v25
				goto b3
			b2:
				;
				v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{k, v})
				if callErr != nil {
					return nil, callErr
				}
				v30 = v28
				goto b3
			b3:
				;
				return v30, nil
			})
			goto b1
		} else {
			x = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{f_3, x})
		if callErr != nil {
			return nil, callErr
		}
		arg__35298, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{f_3, x})
		if callErr != nil {
			return nil, callErr
		}
		v16, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{vm.EmptyPersistentMap, arg__35298})
		if callErr != nil {
			return nil, callErr
		}
		v19 = v16
		goto b3
	b2:
		;
		v19 = x
		goto b3
	b3:
		;
		return v19, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v9, nil
}
func postwalk_replace(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v9 vm.Value
	var callErr error
	v9, callErr = rt.InvokeValue(rt.LookupVar("clojure.walk", "postwalk").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v7 vm.Value
		var x vm.Value
		var smap_3 vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v7) {
			x = arg0
			smap_3 = arg0
			goto b1
		} else {
			x = arg0
			goto b2
		}
	b1:
		;
		v9, callErr = rt.InvokeValue(smap_3, []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		v12 = v9
		goto b3
	b2:
		;
		v12 = x
		goto b3
	b3:
		;
		return v12, nil
	}), arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v9, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("walk", map[string]vm.Value{"walk": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("walk: wrong number of arguments %d (expected 3)", len(args))
		}
		return walk(args[0], args[1], args[2])
	}), "postwalk": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("postwalk: wrong number of arguments %d (expected 2)", len(args))
		}
		return postwalk(args[0], args[1])
	}), "keywordize-keys": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("keywordize-keys: wrong number of arguments %d (expected 1)", len(args))
		}
		return keywordize_keys(args[0])
	}), "prewalk": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("prewalk: wrong number of arguments %d (expected 2)", len(args))
		}
		return prewalk(args[0], args[1])
	}), "prewalk-replace": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("prewalk-replace: wrong number of arguments %d (expected 2)", len(args))
		}
		return prewalk_replace(args[0], args[1])
	}), "stringify-keys": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("stringify-keys: wrong number of arguments %d (expected 1)", len(args))
		}
		return stringify_keys(args[0])
	}), "postwalk-replace": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("postwalk-replace: wrong number of arguments %d (expected 2)", len(args))
		}
		return postwalk_replace(args[0], args[1])
	}),
	})
}
