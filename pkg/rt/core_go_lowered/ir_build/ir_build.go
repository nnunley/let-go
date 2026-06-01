package ir_build

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func add_inst_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value, arg4 vm.Value) (vm.Value, error) {
	var f vm.Value
	var nid vm.Value
	var arg__329 vm.Value
	var si vm.Value
	var callErr error
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	nid, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{f, arg1, arg2, arg3, arg4})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__329, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	si, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__329, vm.Keyword("source-info")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(si) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-source-info!").Deref(), []vm.Value{f, nid, si})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	return nid, nil
}
func add_terminator_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value, arg4 vm.Value) (vm.Value, error) {
	var f vm.Value
	var nid vm.Value
	var arg__361 vm.Value
	var si vm.Value
	var callErr error
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	nid, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-terminator!").Deref(), []vm.Value{f, arg1, arg2, arg3, arg4})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__361, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	si, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__361, vm.Keyword("source-info")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(si) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-source-info!").Deref(), []vm.Value{f, nid, si})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	return nid, nil
}
func attach_name_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var ctx vm.Value
	var inst_id vm.Value
	var sym vm.Value
	var f vm.Value
	var arg__388 vm.Value
	var arg__401 vm.Value
	var arg__402 vm.Value
	var v46 vm.Value
	var v50 vm.Value
	var v22 bool
	var v25 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "integer?").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		ctx = arg0
		inst_id = arg1
		sym = arg2
		goto b4
	} else {
		ctx = arg0
		inst_id = arg1
		sym = arg2
		goto b5
	}
b1:
	;
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__388, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-named-source-info").Deref(), []vm.Value{arg__388})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__401, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__402, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-named-source-info").Deref(), []vm.Value{arg__401})
	if callErr != nil {
		return nil, callErr
	}
	v46, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-source-info!").Deref(), []vm.Value{f, inst_id, arg__402})
	if callErr != nil {
		return nil, callErr
	}
	v50 = v46
	goto b3
b2:
	;
	v50 = vm.NIL
	goto b3
b3:
	;
	return v50, nil
b4:
	;
	v22 = rt.GeValue(inst_id, vm.Int(0))
	v25 = vm.Boolean(v22)
	goto b6
b5:
	;
	v25 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v25) {
		goto b1
	} else {
		goto b2
	}
}
func bind_local_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var s vm.Value
	var stack vm.Value
	var arg__414 vm.Value
	var top_idx vm.Value
	var top vm.Value
	var arg__443 vm.Value
	var arg__468 vm.Value
	var arg__469 vm.Value
	var arg__495 vm.Value
	var arg__520 vm.Value
	var arg__521 vm.Value
	var arg__522 vm.Value
	var v49 vm.Value
	var callErr error
	s, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	stack, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__414, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	top_idx = rt.SubValue(arg__414, vm.Int(1))
	top, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{stack, top_idx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{top, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__443, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{top, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{stack, top_idx, arg__443})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{top, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__468, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{top, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__469, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{stack, top_idx, arg__468})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{s, vm.Keyword("locals"), arg__469})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{top, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__495, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{top, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{stack, top_idx, arg__495})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{top, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__520, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{top, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__521, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{stack, top_idx, arg__520})
	if callErr != nil {
		return nil, callErr
	}
	arg__522, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{s, vm.Keyword("locals"), arg__521})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{arg0, arg__522})
	if callErr != nil {
		return nil, callErr
	}
	v49, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "attach-name!").Deref(), []vm.Value{arg0, arg2, arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v49, nil
}
func binding_syms(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var sym vm.Value
	var v7 vm.Value
	var v12 vm.Value
	var v251 vm.Value
	var arg__551 vm.Value
	var arg__555 vm.Value
	var arg__566 vm.Value
	var arg__567 vm.Value
	var v35 vm.Value
	var v40 vm.Value
	var v248 vm.Value
	var tem__G__0 vm.Value
	var v245 vm.Value
	var arg__607 vm.Value
	var arg__637 vm.Value
	var v62 vm.Value
	var keys_syms vm.Value
	var arg__650 vm.Value
	var arg__656 vm.Value
	var v91 vm.Value
	var strs_syms vm.Value
	var v111 vm.Value
	var as_sym vm.Value
	var arg__689 vm.Value
	var arg__717 vm.Value
	var arg__718 vm.Value
	var arg__722 vm.Value
	var arg__748 vm.Value
	var arg__776 vm.Value
	var arg__777 vm.Value
	var arg__778 vm.Value
	var other_syms vm.Value
	var arg__793 vm.Value
	var arg__803 vm.Value
	var v231 vm.Value
	var v238 vm.Value
	var v242 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v4) {
		sym = arg0
		goto b1
	} else {
		sym = arg0
		goto b2
	}
b1:
	;
	v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	v251 = v7
	goto b3
b2:
	;
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v251, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "flatten").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__551, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "flatten").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.LookupVar("clojure.core", "symbol?").Deref(), arg__551})
	if callErr != nil {
		return nil, callErr
	}
	arg__555, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "flatten").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__566, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "flatten").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__567, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.LookupVar("clojure.core", "symbol?").Deref(), arg__566})
	if callErr != nil {
		return nil, callErr
	}
	v35, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__555, arg__567})
	if callErr != nil {
		return nil, callErr
	}
	v248 = v35
	goto b6
b5:
	;
	v40, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v40) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v251 = v248
	goto b3
b7:
	;
	tem__G__0, callErr = rt.InvokeValue(vm.Keyword("keys"), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(tem__G__0) {
		goto b10
	} else {
		goto b11
	}
b8:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b19
	} else {
		goto b20
	}
b9:
	;
	v248 = v245
	goto b6
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var x vm.Value
		var arg__601 vm.Value
		var v11 vm.Value
		var v14 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v4) {
			x = arg0
			goto b1
		} else {
			x = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		arg__601, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol").Deref(), []vm.Value{arg__601})
		if callErr != nil {
			return nil, callErr
		}
		v14 = v11
		goto b3
	b2:
		;
		v14 = x
		goto b3
	b3:
		;
		return v14, nil
	}), tem__G__0})
	if callErr != nil {
		return nil, callErr
	}
	arg__607, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__637, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var x vm.Value
		var arg__634 vm.Value
		var v11 vm.Value
		var v14 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v4) {
			x = arg0
			goto b1
		} else {
			x = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		arg__634, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol").Deref(), []vm.Value{arg__634})
		if callErr != nil {
			return nil, callErr
		}
		v14 = v11
		goto b3
	b2:
		;
		v14 = x
		goto b3
	b3:
		;
		return v14, nil
	}), tem__G__0})
	if callErr != nil {
		return nil, callErr
	}
	v62, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__607, arg__637})
	if callErr != nil {
		return nil, callErr
	}
	keys_syms = v62
	goto b12
b11:
	;
	keys_syms = vm.NIL
	goto b12
b12:
	;
	tem__G__0, callErr = rt.InvokeValue(vm.Keyword("strs"), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(tem__G__0) {
		goto b13
	} else {
		goto b14
	}
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("clojure.core", "symbol").Deref(), tem__G__0})
	if callErr != nil {
		return nil, callErr
	}
	arg__650, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__656, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("clojure.core", "symbol").Deref(), tem__G__0})
	if callErr != nil {
		return nil, callErr
	}
	v91, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__650, arg__656})
	if callErr != nil {
		return nil, callErr
	}
	strs_syms = v91
	goto b15
b14:
	;
	strs_syms = vm.NIL
	goto b15
b15:
	;
	tem__G__0, callErr = rt.InvokeValue(vm.Keyword("as"), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(tem__G__0) {
		goto b16
	} else {
		goto b17
	}
b16:
	;
	v111, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{tem__G__0})
	if callErr != nil {
		return nil, callErr
	}
	as_sym = v111
	goto b18
b17:
	;
	as_sym = vm.NIL
	goto b18
b18:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "dissoc").Deref(), []vm.Value{sym, vm.Keyword("keys"), vm.Keyword("strs"), vm.Keyword("as"), vm.Keyword("or")})
	if callErr != nil {
		return nil, callErr
	}
	arg__689, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "dissoc").Deref(), []vm.Value{sym, vm.Keyword("keys"), vm.Keyword("strs"), vm.Keyword("as"), vm.Keyword("or")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{arg__689})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "dissoc").Deref(), []vm.Value{sym, vm.Keyword("keys"), vm.Keyword("strs"), vm.Keyword("as"), vm.Keyword("or")})
	if callErr != nil {
		return nil, callErr
	}
	arg__717, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "dissoc").Deref(), []vm.Value{sym, vm.Keyword("keys"), vm.Keyword("strs"), vm.Keyword("as"), vm.Keyword("or")})
	if callErr != nil {
		return nil, callErr
	}
	arg__718, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{arg__717})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.LookupVar("clojure.core", "symbol?").Deref(), arg__718})
	if callErr != nil {
		return nil, callErr
	}
	arg__722, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "dissoc").Deref(), []vm.Value{sym, vm.Keyword("keys"), vm.Keyword("strs"), vm.Keyword("as"), vm.Keyword("or")})
	if callErr != nil {
		return nil, callErr
	}
	arg__748, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "dissoc").Deref(), []vm.Value{sym, vm.Keyword("keys"), vm.Keyword("strs"), vm.Keyword("as"), vm.Keyword("or")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{arg__748})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "dissoc").Deref(), []vm.Value{sym, vm.Keyword("keys"), vm.Keyword("strs"), vm.Keyword("as"), vm.Keyword("or")})
	if callErr != nil {
		return nil, callErr
	}
	arg__776, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "dissoc").Deref(), []vm.Value{sym, vm.Keyword("keys"), vm.Keyword("strs"), vm.Keyword("as"), vm.Keyword("or")})
	if callErr != nil {
		return nil, callErr
	}
	arg__777, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{arg__776})
	if callErr != nil {
		return nil, callErr
	}
	arg__778, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.LookupVar("clojure.core", "symbol?").Deref(), arg__777})
	if callErr != nil {
		return nil, callErr
	}
	other_syms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__722, arg__778})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{keys_syms, strs_syms, as_sym, other_syms})
	if callErr != nil {
		return nil, callErr
	}
	arg__793, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__803, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{keys_syms, strs_syms, as_sym, other_syms})
	if callErr != nil {
		return nil, callErr
	}
	v231, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__793, arg__803})
	if callErr != nil {
		return nil, callErr
	}
	v245 = v231
	goto b9
b19:
	;
	v238, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v242 = v238
	goto b21
b20:
	;
	v242 = vm.NIL
	goto b21
b21:
	;
	v245 = v242
	goto b9
}
func build_args(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var pre_locals vm.Value
	var syms vm.Value
	var results vm.Value
	var doseq_seq__805 vm.Value
	var doseq_loop__806 vm.Value
	var ctx vm.Value
	var a vm.Value
	var r vm.Value
	var v59 vm.Value
	var arg__901 vm.Value
	var arg__905 vm.Value
	var arg__932 vm.Value
	var arg__936 vm.Value
	var arg__937 vm.Value
	var threaded vm.Value
	var post_locals vm.Value
	var doseq_seq__807 vm.Value
	var sym vm.Value
	var v98 vm.Value
	var doseq_loop__808 vm.Value
	var vec__809 vm.Value
	var val vm.Value
	var and__x vm.Value
	var v290 vm.Value
	var arg__983 vm.Value
	var v251 vm.Value
	var v254 vm.Value
	var callErr error
	pre_locals, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "push-locals!").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	syms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NewArrayVector([]vm.Value{})})
	if callErr != nil {
		return nil, callErr
	}
	results, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NewArrayVector([]vm.Value{})})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__805, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__806 = doseq_seq__805
	ctx = arg1
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__806) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	a, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__806})
	if callErr != nil {
		return nil, callErr
	}
	r, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{a, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v59, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "terminated?").Deref(), []vm.Value{r})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v59) {
		goto b5
	} else {
		goto b6
	}
b3:
	;
	goto b4
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{syms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{results})
	if callErr != nil {
		return nil, callErr
	}
	arg__901, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{syms})
	if callErr != nil {
		return nil, callErr
	}
	arg__905, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{results})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var sym vm.Value
		var orig vm.Value
		var ctx_5 vm.Value
		var or__x vm.Value
		var v30 vm.Value
		var v23 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			sym = arg0
			orig = arg1
			ctx_5 = ctx
			goto b1
		} else {
			orig = arg1
			goto b2
		}
	b1:
		;
		or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx_5, sym})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(or__x) {
			goto b4
		} else {
			goto b5
		}
	b2:
		;
		v30 = orig
		goto b3
	b3:
		;
		return v30, nil
	b4:
		;
		v23 = or__x
		goto b6
	b5:
		;
		v23 = orig
		goto b6
	b6:
		;
		v30 = v23
		goto b3
	}), arg__901, arg__905})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{syms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{results})
	if callErr != nil {
		return nil, callErr
	}
	arg__932, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{syms})
	if callErr != nil {
		return nil, callErr
	}
	arg__936, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{results})
	if callErr != nil {
		return nil, callErr
	}
	arg__937, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var sym vm.Value
		var orig vm.Value
		var ctx_5 vm.Value
		var or__x vm.Value
		var v30 vm.Value
		var v23 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			sym = arg0
			orig = arg1
			ctx_5 = ctx
			goto b1
		} else {
			orig = arg1
			goto b2
		}
	b1:
		;
		or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx_5, sym})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(or__x) {
			goto b4
		} else {
			goto b5
		}
	b2:
		;
		v30 = orig
		goto b3
	b3:
		;
		return v30, nil
	b4:
		;
		v23 = or__x
		goto b6
	b5:
		;
		v23 = orig
		goto b6
	b6:
		;
		v30 = v23
		goto b3
	}), arg__932, arg__936})
	if callErr != nil {
		return nil, callErr
	}
	threaded, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__937})
	if callErr != nil {
		return nil, callErr
	}
	post_locals, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "pop-locals!").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__807, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{post_locals})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__808 = doseq_seq__807
	goto b8
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{syms, rt.LookupVar("clojure.core", "conj").Deref(), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{results, rt.LookupVar("clojure.core", "conj").Deref(), r})
	if callErr != nil {
		return nil, callErr
	}
	goto b7
b6:
	;
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "gensym").Deref(), []vm.Value{vm.String("arg__")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{ctx, sym, r})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{syms, rt.LookupVar("clojure.core", "conj").Deref(), sym})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{results, rt.LookupVar("clojure.core", "conj").Deref(), r})
	if callErr != nil {
		return nil, callErr
	}
	goto b7
b7:
	;
	v98, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__806})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__806 = v98
	goto b1
b8:
	;
	if vm.IsTruthy(doseq_loop__808) {
		goto b9
	} else {
		goto b10
	}
b9:
	;
	vec__809, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__808})
	if callErr != nil {
		return nil, callErr
	}
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__809, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	val, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__809, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{pre_locals, sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b15
	} else {
		goto b16
	}
b10:
	;
	goto b11
b11:
	;
	return threaded, nil
b12:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{ctx, sym, val})
	if callErr != nil {
		return nil, callErr
	}
	goto b14
b13:
	;
	goto b14
b14:
	;
	v290, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__808})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__808 = v290
	goto b8
b15:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{pre_locals, sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__983, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{pre_locals, sym})
	if callErr != nil {
		return nil, callErr
	}
	v251, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{val, arg__983})
	if callErr != nil {
		return nil, callErr
	}
	v254 = v251
	goto b17
b16:
	;
	v254 = and__x
	goto b17
b17:
	;
	if vm.IsTruthy(v254) {
		goto b12
	} else {
		goto b13
	}
}
func build_builtin_op(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__1003 vm.Value
	var args vm.Value
	var v20 vm.Value
	var op_kw vm.Value
	var form vm.Value
	var ctx vm.Value
	var arg__1014 vm.Value
	var v33 bool
	var arg__1061 vm.Value
	var v80 bool
	var v194 vm.Value
	var arg__1030 vm.Value
	var arg__1037 vm.Value
	var arg__1041 vm.Value
	var arg__1049 vm.Value
	var arg__1050 vm.Value
	var v58 vm.Value
	var v61 vm.Value
	var v63 vm.Value
	var v91 bool
	var arg__1123 vm.Value
	var v150 bool
	var v188 vm.Value
	var arg__1077 vm.Value
	var zero_id vm.Value
	var arg__1094 vm.Value
	var arg__1102 vm.Value
	var arg__1111 vm.Value
	var arg__1112 vm.Value
	var v126 vm.Value
	var v131 vm.Value
	var v133 vm.Value
	var arg__1137 vm.Value
	var v159 vm.Value
	var v182 vm.Value
	var v172 vm.Value
	var v176 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__1003, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	args, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-args").Deref(), []vm.Value{arg__1003, arg2})
	if callErr != nil {
		return nil, callErr
	}
	v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.build", "unary-only-ops").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v20) {
		op_kw = arg0
		form = arg1
		ctx = arg2
		goto b1
	} else {
		op_kw = arg0
		ctx = arg2
		goto b2
	}
b1:
	;
	arg__1014, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	v33 = arg__1014 == vm.Int(1)
	if v33 {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	arg__1061, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	v80 = arg__1061 == vm.Int(1)
	if v80 {
		goto b7
	} else {
		goto b8
	}
b3:
	;
	return v194, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1030, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__1030})
	if callErr != nil {
		return nil, callErr
	}
	arg__1037, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1041, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1049, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__1050, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__1049})
	if callErr != nil {
		return nil, callErr
	}
	v58, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{arg__1037, arg__1041, op_kw, arg__1050, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v63 = v58
	goto b6
b5:
	;
	v61, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-call").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v63 = v61
	goto b6
b6:
	;
	v194 = v63
	goto b3
b7:
	;
	v91 = op_kw == vm.Keyword("sub")
	if v91 {
		goto b10
	} else {
		goto b11
	}
b8:
	;
	arg__1123, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	v150 = arg__1123 == vm.Int(2)
	if v150 {
		goto b13
	} else {
		goto b14
	}
b9:
	;
	v194 = v188
	goto b3
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1077, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	zero_id, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__1077, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1094, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{zero_id, arg__1094})
	if callErr != nil {
		return nil, callErr
	}
	arg__1102, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1111, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__1112, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{zero_id, arg__1111})
	if callErr != nil {
		return nil, callErr
	}
	v126, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__1102, vm.Keyword("sub"), arg__1112, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v133 = v126
	goto b12
b11:
	;
	v131, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v133 = v131
	goto b12
b12:
	;
	v188 = v133
	goto b9
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1137, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v159, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__1137, op_kw, args, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v182 = v159
	goto b15
b14:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v188 = v182
	goto b9
b16:
	;
	v172, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "fold-binary-chain").Deref(), []vm.Value{op_kw, args, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v176 = v172
	goto b18
b17:
	;
	v176 = vm.NIL
	goto b18
b18:
	;
	v182 = v176
	goto b15
}
func build_call(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var head vm.Value
	var arg__1160 vm.Value
	var v19 vm.Value
	var form vm.Value
	var ctx vm.Value
	var v22 vm.Value
	var v25 vm.Value
	var fn_id vm.Value
	var arg__1186 vm.Value
	var v37 vm.Value
	var callErr error
	head, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__1160, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-args").Deref(), []vm.Value{arg__1160, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{head})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v19) {
		form = arg0
		ctx = arg1
		goto b1
	} else {
		form = arg0
		ctx = arg1
		goto b2
	}
b1:
	;
	v22, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-symbol").Deref(), []vm.Value{head, ctx})
	if callErr != nil {
		return nil, callErr
	}
	fn_id = v22
	goto b3
b2:
	;
	v25, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{head, ctx})
	if callErr != nil {
		return nil, callErr
	}
	fn_id = v25
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__1186, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v37, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-call-with-head").Deref(), []vm.Value{fn_id, arg__1186, ctx})
	if callErr != nil {
		return nil, callErr
	}
	return v37, nil
}
func build_call_with_head(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var head_sym vm.Value
	var arg_ids vm.Value
	var threaded vm.Value
	var arg__1232 vm.Value
	var arg__1243 vm.Value
	var arg__1257 vm.Value
	var arg__1258 vm.Value
	var arg__1262 vm.Value
	var v41 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "push-locals!").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	head_sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "gensym").Deref(), []vm.Value{vm.String("head__")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{arg2, head_sym, arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg_ids, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-args").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	threaded, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{arg2, head_sym})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "pop-locals!").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{threaded, arg_ids})
	if callErr != nil {
		return nil, callErr
	}
	arg__1232, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{threaded, arg_ids})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__1232})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg_ids})
	if callErr != nil {
		return nil, callErr
	}
	arg__1243, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{threaded, arg_ids})
	if callErr != nil {
		return nil, callErr
	}
	arg__1257, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{threaded, arg_ids})
	if callErr != nil {
		return nil, callErr
	}
	arg__1258, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__1257})
	if callErr != nil {
		return nil, callErr
	}
	arg__1262, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg_ids})
	if callErr != nil {
		return nil, callErr
	}
	v41, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{arg2, arg__1243, vm.Keyword("call"), arg__1258, arg__1262})
	if callErr != nil {
		return nil, callErr
	}
	return v41, nil
}
func build_do(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var forms vm.Value
	var fs vm.Value
	var last_id vm.Value
	var ctx vm.Value
	var v21 vm.Value
	var v24 vm.Value
	var arg__1281 vm.Value
	var v30 vm.Value
	var v33 vm.Value
	var callErr error
	forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	fs = forms
	last_id = vm.NIL
	ctx = arg1
	goto b1
b1:
	;
	v21, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v21) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	arg__1281, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	v30, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{arg__1281, ctx})
	if callErr != nil {
		return nil, callErr
	}
	fs = v24
	last_id = v30
	goto b1
b3:
	;
	v33 = last_id
	goto b4
b4:
	;
	return v33, nil
}
func build_fn(arg0 vm.Value) (vm.Value, error) {
	var name_sym vm.Value
	var maybe_doc vm.Value
	var has_doc_QMARK_ vm.Value
	var defn_form vm.Value
	var v24 vm.Value
	var args_vec vm.Value
	var head__1305 vm.Value
	var arg__1306 int
	var body_forms vm.Value
	var multi_QMARK_ vm.Value
	var arg__1321 vm.Value
	var f vm.Value
	var ctx vm.Value
	var arities vm.Value
	var expanded_forms vm.Value
	var arg__1517 vm.Value
	var all_caps vm.Value
	var arg__1531 vm.Value
	var captures vm.Value
	var templates vm.Value
	var template vm.Value
	var closure_id vm.Value
	var final_blk vm.Value
	var arg__1617 vm.Value
	var expanded vm.Value
	var v691 vm.Value
	var v188 vm.Value
	var flat_args vm.Value
	var v220 vm.Value
	var flat_body vm.Value
	var v255 vm.Value
	var variadic_QMARK_ vm.Value
	var arity vm.Value
	var arg__1643 vm.Value
	var entry_blk vm.Value
	var arg__1666 vm.Value
	var arg__1671 vm.Value
	var arg__1688 vm.Value
	var arg__1693 vm.Value
	var arg__1694 vm.Value
	var i int
	var v348 bool
	var arg_id vm.Value
	var arg__1723 vm.Value
	var v362 int
	var fs vm.Value
	var last_id vm.Value
	var v425 vm.Value
	var v428 vm.Value
	var arg__1741 vm.Value
	var v434 vm.Value
	var last_val vm.Value
	var v499 vm.Value
	var v551 vm.Value
	var v634 vm.Value
	var arg__1760 vm.Value
	var arg__1761 vm.Value
	var head__1759 vm.Value
	var v639 vm.Value
	var arg__1768 vm.Value
	var callErr error
	name_sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	maybe_doc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	has_doc_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{maybe_doc})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(has_doc_QMARK_) {
		defn_form = arg0
		goto b1
	} else {
		defn_form = arg0
		goto b2
	}
b1:
	;
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{defn_form, vm.Int(3)})
	if callErr != nil {
		return nil, callErr
	}
	args_vec = v24
	goto b3
b2:
	;
	args_vec = maybe_doc
	goto b3
b3:
	;
	if vm.IsTruthy(has_doc_QMARK_) {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	goto b6
b5:
	;
	goto b6
b6:
	;
	if vm.IsTruthy(has_doc_QMARK_) {
		head__1305 = rt.LookupVar("clojure.core", "drop").Deref()
		goto b7
	} else {
		head__1305 = rt.LookupVar("clojure.core", "drop").Deref()
		goto b8
	}
b7:
	;
	arg__1306 = 4
	goto b9
b8:
	;
	arg__1306 = 3
	goto b9
b9:
	;
	body_forms, callErr = rt.InvokeValue(head__1305, []vm.Value{vm.Int(arg__1306), defn_form})
	if callErr != nil {
		return nil, callErr
	}
	multi_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{args_vec})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(multi_QMARK_) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{name_sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__1321, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{name_sym})
	if callErr != nil {
		return nil, callErr
	}
	f, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-fn").Deref(), []vm.Value{arg__1321, vm.Int(0), vm.Boolean(false)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "entry-block").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	ctx, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "new-context").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arities, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{args_vec, body_forms})
	if callErr != nil {
		return nil, callErr
	}
	expanded_forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg_vec vm.Value
		var body_forms vm.Value
		var e vm.Value
		var v22 vm.Value
		var v24 vm.Value
		var callErr error
		arg_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		body_forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		e, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-fn-args").Deref(), []vm.Value{arg_vec, body_forms})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(e) {
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v24 = e
		goto b3
	b2:
		;
		v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("variadic?"), vm.FALSE, vm.Keyword("flat-args"), arg_vec, vm.Keyword("body"), body_forms})
		if callErr != nil {
			return nil, callErr
		}
		v24 = v22
		goto b3
	b3:
		;
		return v24, nil
	}), arities})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__1517, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	all_caps, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var flat_args vm.Value
		var body vm.Value
		var arg__1463 vm.Value
		var arg_set vm.Value
		var arg__1478 vm.Value
		var frees vm.Value
		var arg__1514 vm.Value
		var v44 vm.Value
		var callErr error
		flat_args, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg1, vm.Keyword("flat-args")})
		if callErr != nil {
			return nil, callErr
		}
		body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg1, vm.Keyword("body")})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__1463, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg_set, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__1463, flat_args})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{vm.Symbol("do"), body})
		if callErr != nil {
			return nil, callErr
		}
		arg__1478, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{vm.Symbol("do"), body})
		if callErr != nil {
			return nil, callErr
		}
		frees, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg__1478, arg_set})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var v3 vm.Value
			var callErr error
			v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
			if callErr != nil {
				return nil, callErr
			}
			return v3, nil
		}), frees})
		if callErr != nil {
			return nil, callErr
		}
		arg__1514, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var v3 vm.Value
			var callErr error
			v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
			if callErr != nil {
				return nil, callErr
			}
			return v3, nil
		}), frees})
		if callErr != nil {
			return nil, callErr
		}
		v44, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg0, arg__1514})
		if callErr != nil {
			return nil, callErr
		}
		return v44, nil
	}), arg__1517, expanded_forms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), all_caps})
	if callErr != nil {
		return nil, callErr
	}
	arg__1531, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), all_caps})
	if callErr != nil {
		return nil, callErr
	}
	captures, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__1531})
	if callErr != nil {
		return nil, callErr
	}
	templates, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var flat_args vm.Value
		var body vm.Value
		var variadic_QMARK_ vm.Value
		var v16 vm.Value
		var callErr error
		flat_args, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg0, vm.Keyword("flat-args")})
		if callErr != nil {
			return nil, callErr
		}
		body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg0, vm.Keyword("body")})
		if callErr != nil {
			return nil, callErr
		}
		variadic_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg0, vm.Keyword("variadic?")})
		if callErr != nil {
			return nil, callErr
		}
		v16, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-inner-fn-template").Deref(), []vm.Value{name_sym, flat_args, body, captures, variadic_QMARK_})
		if callErr != nil {
			return nil, callErr
		}
		return v16, nil
	}), expanded_forms})
	if callErr != nil {
		return nil, callErr
	}
	template, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("fns"), templates, vm.Keyword("kind"), vm.Keyword("multi-fn-template")})
	if callErr != nil {
		return nil, callErr
	}
	closure_id, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "emit-template-closure").Deref(), []vm.Value{template, captures, ctx})
	if callErr != nil {
		return nil, callErr
	}
	final_blk, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{closure_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__1617, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{closure_id})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-terminator!").Deref(), []vm.Value{f, final_blk, vm.Keyword("return"), arg__1617, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v691 = f
	goto b12
b11:
	;
	expanded, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-fn-args").Deref(), []vm.Value{args_vec, body_forms})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(expanded) {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	return v691, nil
b13:
	;
	v188, callErr = rt.InvokeValue(vm.Keyword("flat-args"), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	flat_args = v188
	goto b15
b14:
	;
	flat_args = args_vec
	goto b15
b15:
	;
	if vm.IsTruthy(expanded) {
		goto b16
	} else {
		goto b17
	}
b16:
	;
	v220, callErr = rt.InvokeValue(vm.Keyword("body"), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	flat_body = v220
	goto b18
b17:
	;
	flat_body = body_forms
	goto b18
b18:
	;
	if vm.IsTruthy(expanded) {
		goto b19
	} else {
		goto b20
	}
b19:
	;
	v255, callErr = rt.InvokeValue(vm.Keyword("variadic?"), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	variadic_QMARK_ = v255
	goto b21
b20:
	;
	variadic_QMARK_ = vm.Boolean(false)
	goto b21
b21:
	;
	arity, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{flat_args})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{name_sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__1643, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{name_sym})
	if callErr != nil {
		return nil, callErr
	}
	f, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-fn").Deref(), []vm.Value{arg__1643, arity, variadic_QMARK_})
	if callErr != nil {
		return nil, callErr
	}
	entry_blk, callErr = rt.InvokeValue(rt.LookupVar("ir", "entry-block").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	ctx, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "new-context").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{flat_args})
	if callErr != nil {
		return nil, callErr
	}
	arg__1666, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1671, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{flat_args})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg__1666, vm.Keyword("fn-arg-syms"), arg__1671})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{flat_args})
	if callErr != nil {
		return nil, callErr
	}
	arg__1688, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__1693, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{flat_args})
	if callErr != nil {
		return nil, callErr
	}
	arg__1694, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg__1688, vm.Keyword("fn-arg-syms"), arg__1693})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{ctx, arg__1694})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	goto b22
b22:
	;
	v348 = rt.LtValue(vm.Int(i), arity)
	if v348 {
		goto b23
	} else {
		goto b24
	}
b23:
	;
	arg_id, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{f, entry_blk, vm.Keyword("load-arg"), vm.NewArrayVector([]vm.Value{}), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{flat_args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__1723, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{flat_args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{ctx, arg__1723, arg_id})
	if callErr != nil {
		return nil, callErr
	}
	v362 = i + 1
	i = v362
	goto b22
b24:
	;
	goto b25
b25:
	;
	fs = flat_body
	last_id = vm.NIL
	goto b26
b26:
	;
	v425, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v425) {
		goto b27
	} else {
		goto b28
	}
b27:
	;
	v428, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	arg__1741, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	v434, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{arg__1741, ctx})
	if callErr != nil {
		return nil, callErr
	}
	fs = v428
	last_id = v434
	goto b26
b28:
	;
	last_val = last_id
	goto b29
b29:
	;
	final_blk, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v499, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "terminated?").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v499) {
		goto b30
	} else {
		goto b31
	}
b30:
	;
	goto b32
b31:
	;
	v551, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v551) {
		goto b33
	} else {
		goto b34
	}
b32:
	;
	v691 = f
	goto b12
b33:
	;
	goto b35
b34:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	goto b35
b35:
	;
	v634, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v634) {
		arg__1760 = f
		arg__1761 = final_blk
		head__1759 = rt.LookupVar("ir", "add-terminator!").Deref()
		goto b36
	} else {
		arg__1760 = f
		arg__1761 = final_blk
		head__1759 = rt.LookupVar("ir", "add-terminator!").Deref()
		goto b37
	}
b36:
	;
	arg__1768 = vm.NewArrayVector([]vm.Value{})
	goto b38
b37:
	;
	v639, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	arg__1768 = v639
	goto b38
b38:
	;
	_, callErr = rt.InvokeValue(head__1759, []vm.Value{arg__1760, arg__1761, vm.Keyword("return"), arg__1768, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	goto b32
}
func build_fn_STAR_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var maybe_name vm.Value
	var raw_rest vm.Value
	var has_name_QMARK_ vm.Value
	var ctx vm.Value
	var name_sym vm.Value
	var v66 vm.Value
	var rest_forms vm.Value
	var and__x vm.Value
	var arg__1811 vm.Value
	var v105 vm.Value
	var multi_QMARK_ vm.Value
	var expanded_forms vm.Value
	var arg__1994 vm.Value
	var all_caps vm.Value
	var arg__2008 vm.Value
	var captures vm.Value
	var templates vm.Value
	var template vm.Value
	var v182 vm.Value
	var args_vec vm.Value
	var body_forms vm.Value
	var v189 vm.Value
	var v191 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	maybe_name, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	raw_rest, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), arg0})
	if callErr != nil {
		return nil, callErr
	}
	has_name_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{maybe_name})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(has_name_QMARK_) {
		ctx = arg1
		goto b1
	} else {
		ctx = arg1
		goto b2
	}
b1:
	;
	name_sym = maybe_name
	goto b3
b2:
	;
	name_sym = vm.String("fn*")
	goto b3
b3:
	;
	if vm.IsTruthy(has_name_QMARK_) {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	rest_forms = raw_rest
	goto b6
b5:
	;
	v66, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{maybe_name, raw_rest})
	if callErr != nil {
		return nil, callErr
	}
	rest_forms = v66
	goto b6
b6:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b7
	} else {
		goto b8
	}
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	arg__1811, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	v105, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{arg__1811})
	if callErr != nil {
		return nil, callErr
	}
	multi_QMARK_ = v105
	goto b9
b8:
	;
	multi_QMARK_ = and__x
	goto b9
b9:
	;
	if vm.IsTruthy(multi_QMARK_) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	expanded_forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var body_forms vm.Value
		var e vm.Value
		var v22 vm.Value
		var v24 vm.Value
		var callErr error
		args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		body_forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		e, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-fn-args").Deref(), []vm.Value{args_vec, body_forms})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(e) {
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v24 = e
		goto b3
	b2:
		;
		v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("variadic?"), vm.FALSE, vm.Keyword("flat-args"), args_vec, vm.Keyword("body"), body_forms})
		if callErr != nil {
			return nil, callErr
		}
		v24 = v22
		goto b3
	b3:
		;
		return v24, nil
	}), rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__1994, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	all_caps, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var flat_args vm.Value
		var body vm.Value
		var arg__1940 vm.Value
		var arg_set vm.Value
		var arg__1955 vm.Value
		var frees vm.Value
		var arg__1991 vm.Value
		var v44 vm.Value
		var callErr error
		flat_args, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg1, vm.Keyword("flat-args")})
		if callErr != nil {
			return nil, callErr
		}
		body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg1, vm.Keyword("body")})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__1940, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg_set, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__1940, flat_args})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{vm.Symbol("do"), body})
		if callErr != nil {
			return nil, callErr
		}
		arg__1955, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{vm.Symbol("do"), body})
		if callErr != nil {
			return nil, callErr
		}
		frees, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg__1955, arg_set})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var v3 vm.Value
			var callErr error
			v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
			if callErr != nil {
				return nil, callErr
			}
			return v3, nil
		}), frees})
		if callErr != nil {
			return nil, callErr
		}
		arg__1991, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var v3 vm.Value
			var callErr error
			v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
			if callErr != nil {
				return nil, callErr
			}
			return v3, nil
		}), frees})
		if callErr != nil {
			return nil, callErr
		}
		v44, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg0, arg__1991})
		if callErr != nil {
			return nil, callErr
		}
		return v44, nil
	}), arg__1994, expanded_forms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), all_caps})
	if callErr != nil {
		return nil, callErr
	}
	arg__2008, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), all_caps})
	if callErr != nil {
		return nil, callErr
	}
	captures, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__2008})
	if callErr != nil {
		return nil, callErr
	}
	templates, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var flat_args vm.Value
		var body vm.Value
		var variadic_QMARK_ vm.Value
		var v16 vm.Value
		var callErr error
		flat_args, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg0, vm.Keyword("flat-args")})
		if callErr != nil {
			return nil, callErr
		}
		body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg0, vm.Keyword("body")})
		if callErr != nil {
			return nil, callErr
		}
		variadic_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg0, vm.Keyword("variadic?")})
		if callErr != nil {
			return nil, callErr
		}
		v16, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-inner-fn-template").Deref(), []vm.Value{name_sym, flat_args, body, captures, variadic_QMARK_})
		if callErr != nil {
			return nil, callErr
		}
		return v16, nil
	}), expanded_forms})
	if callErr != nil {
		return nil, callErr
	}
	template, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("fns"), templates, vm.Keyword("kind"), vm.Keyword("multi-fn-template")})
	if callErr != nil {
		return nil, callErr
	}
	v182, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "emit-template-closure").Deref(), []vm.Value{template, captures, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v191 = v182
	goto b12
b11:
	;
	args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	body_forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	v189, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-single-fn*").Deref(), []vm.Value{name_sym, args_vec, body_forms, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v191 = v189
	goto b12
b12:
	;
	return v191, nil
}
func build_form(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var si vm.Value
	var arg__2105 vm.Value
	var old_si vm.Value
	var form vm.Value
	var ctx vm.Value
	var v47 vm.Value
	var arg__2132 vm.Value
	var v60 vm.Value
	var v73 vm.Value
	var nid vm.Value
	var arg__2152 vm.Value
	var v84 vm.Value
	var v97 vm.Value
	var v365 vm.Value
	var arg__2172 vm.Value
	var v108 vm.Value
	var v121 vm.Value
	var v358 vm.Value
	var arg__2192 vm.Value
	var v132 vm.Value
	var v145 vm.Value
	var v351 vm.Value
	var arg__2212 vm.Value
	var v156 vm.Value
	var v169 vm.Value
	var v344 vm.Value
	var arg__2232 vm.Value
	var v180 vm.Value
	var v193 vm.Value
	var v337 vm.Value
	var arg__2252 vm.Value
	var v204 vm.Value
	var v217 vm.Value
	var v330 vm.Value
	var v220 vm.Value
	var v233 vm.Value
	var v323 vm.Value
	var v236 vm.Value
	var v249 vm.Value
	var v316 vm.Value
	var v252 vm.Value
	var v265 vm.Value
	var v309 vm.Value
	var v268 vm.Value
	var v302 vm.Value
	var arg__2300 vm.Value
	var v291 vm.Value
	var v295 vm.Value
	var callErr error
	si, callErr = rt.InvokeValue(rt.LookupVar("ir", "form-source-info").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__2105, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	old_si, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__2105, vm.Keyword("source-info")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(si) {
		form = arg0
		ctx = arg1
		goto b1
	} else {
		form = arg0
		ctx = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{ctx, rt.LookupVar("clojure.core", "assoc").Deref(), vm.Keyword("source-info"), si})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	v47, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v47) {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__2132, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v60, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__2132, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	nid = v60
	goto b6
b5:
	;
	v73, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "integer?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v73) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	if vm.IsTruthy(si) {
		goto b40
	} else {
		goto b41
	}
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__2152, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v84, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__2152, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v365 = v84
	goto b9
b8:
	;
	v97, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "float?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v97) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	nid = v365
	goto b6
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__2172, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v108, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__2172, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v358 = v108
	goto b12
b11:
	;
	v121, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v121) {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v365 = v358
	goto b9
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__2192, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v132, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__2192, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v351 = v132
	goto b15
b14:
	;
	v145, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "char?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v145) {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v358 = v351
	goto b12
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__2212, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v156, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__2212, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v344 = v156
	goto b18
b17:
	;
	v169, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v169) {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v351 = v344
	goto b15
b19:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__2232, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v180, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__2232, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v337 = v180
	goto b21
b20:
	;
	v193, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "boolean?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v193) {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v344 = v337
	goto b18
b22:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__2252, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v204, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__2252, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v330 = v204
	goto b24
b23:
	;
	v217, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v217) {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v337 = v330
	goto b21
b25:
	;
	v220, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-symbol").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v323 = v220
	goto b27
b26:
	;
	v233, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v233) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v330 = v323
	goto b24
b28:
	;
	v236, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-vector").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v316 = v236
	goto b30
b29:
	;
	v249, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v249) {
		goto b31
	} else {
		goto b32
	}
b30:
	;
	v323 = v316
	goto b27
b31:
	;
	v252, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-map").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v309 = v252
	goto b33
b32:
	;
	v265, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v265) {
		goto b34
	} else {
		goto b35
	}
b33:
	;
	v316 = v309
	goto b30
b34:
	;
	v268, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-list").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v302 = v268
	goto b36
b35:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b37
	} else {
		goto b38
	}
b36:
	;
	v309 = v302
	goto b33
b37:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("build-form: unrecognized form "), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__2300, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("build-form: unrecognized form "), form})
	if callErr != nil {
		return nil, callErr
	}
	v291, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__2300})
	if callErr != nil {
		return nil, callErr
	}
	v295 = v291
	goto b39
b38:
	;
	v295 = vm.NIL
	goto b39
b39:
	;
	v302 = v295
	goto b36
b40:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{ctx, rt.LookupVar("clojure.core", "assoc").Deref(), vm.Keyword("source-info"), old_si})
	if callErr != nil {
		return nil, callErr
	}
	goto b42
b41:
	;
	goto b42
b42:
	;
	return nid, nil
}
func build_inner_fn_template(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value, arg4 vm.Value) (vm.Value, error) {
	var arg__3360 vm.Value
	var arg__3364 vm.Value
	var inner vm.Value
	var entry vm.Value
	var inner_ctx vm.Value
	var arg__3386 vm.Value
	var arg__3391 vm.Value
	var arg__3408 vm.Value
	var arg__3413 vm.Value
	var arg__3414 vm.Value
	var i int
	var args_vec vm.Value
	var arg__3419 vm.Value
	var v71 bool
	var arg__3446 vm.Value
	var arg__3458 vm.Value
	var v91 int
	var body_forms vm.Value
	var capture_syms vm.Value
	var variadic_QMARK_ vm.Value
	var arg__3464 vm.Value
	var v133 bool
	var arg__3491 vm.Value
	var arg__3503 vm.Value
	var fs vm.Value
	var last_id vm.Value
	var v195 vm.Value
	var v198 vm.Value
	var arg__3520 vm.Value
	var v204 vm.Value
	var last_val vm.Value
	var final_blk vm.Value
	var v248 vm.Value
	var v286 vm.Value
	var arg__3558 vm.Value
	var v398 vm.Value
	var v348 vm.Value
	var arg__3539 vm.Value
	var arg__3540 vm.Value
	var head__3538 vm.Value
	var v353 vm.Value
	var arg__3547 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3360, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__3364, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	inner, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-fn").Deref(), []vm.Value{arg__3360, arg__3364, arg4})
	if callErr != nil {
		return nil, callErr
	}
	entry, callErr = rt.InvokeValue(rt.LookupVar("ir", "entry-block").Deref(), []vm.Value{inner})
	if callErr != nil {
		return nil, callErr
	}
	inner_ctx, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "new-context").Deref(), []vm.Value{inner})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{inner_ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3386, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{inner_ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__3391, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg__3386, vm.Keyword("fn-arg-syms"), arg__3391})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{inner_ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3408, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{inner_ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__3413, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3414, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg__3408, vm.Keyword("fn-arg-syms"), arg__3413})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{inner_ctx, arg__3414})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	args_vec = arg1
	goto b1
b1:
	;
	arg__3419, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args_vec})
	if callErr != nil {
		return nil, callErr
	}
	v71 = rt.LtValue(vm.Int(i), arg__3419)
	if v71 {
		goto b2
	} else {
		body_forms = arg2
		capture_syms = arg3
		variadic_QMARK_ = arg4
		goto b3
	}
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args_vec, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{inner_ctx, entry, vm.Keyword("load-arg"), vm.NewArrayVector([]vm.Value{}), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__3446, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args_vec, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__3458, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{inner_ctx, entry, vm.Keyword("load-arg"), vm.NewArrayVector([]vm.Value{}), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{inner_ctx, arg__3446, arg__3458})
	if callErr != nil {
		return nil, callErr
	}
	v91 = i + 1
	i = v91
	goto b1
b3:
	;
	goto b4
b4:
	;
	goto b5
b5:
	;
	arg__3464, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{capture_syms})
	if callErr != nil {
		return nil, callErr
	}
	v133 = rt.LtValue(vm.Int(i), arg__3464)
	if v133 {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{capture_syms, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{inner, entry, vm.Keyword("load-closed"), vm.NewArrayVector([]vm.Value{}), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__3491, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{capture_syms, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__3503, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{inner, entry, vm.Keyword("load-closed"), vm.NewArrayVector([]vm.Value{}), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{inner_ctx, arg__3491, arg__3503})
	if callErr != nil {
		return nil, callErr
	}
	goto b5
b7:
	;
	goto b8
b8:
	;
	fs = body_forms
	last_id = vm.NIL
	goto b9
b9:
	;
	v195, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v195) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	v198, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	arg__3520, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	v204, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{arg__3520, inner_ctx})
	if callErr != nil {
		return nil, callErr
	}
	fs = v198
	last_id = v204
	goto b9
b11:
	;
	last_val = last_id
	goto b12
b12:
	;
	final_blk, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{inner_ctx})
	if callErr != nil {
		return nil, callErr
	}
	v248, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "terminated?").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v248) {
		goto b13
	} else {
		goto b14
	}
b13:
	;
	goto b15
b14:
	;
	v286, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v286) {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	arg__3558, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args_vec})
	if callErr != nil {
		return nil, callErr
	}
	v398, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("fn"), inner, vm.Keyword("variadic?"), variadic_QMARK_, vm.Keyword("arity"), arg__3558, vm.Keyword("kind"), vm.Keyword("fn-template")})
	if callErr != nil {
		return nil, callErr
	}
	return v398, nil
b16:
	;
	goto b18
b17:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	goto b18
b18:
	;
	v348, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v348) {
		arg__3539 = inner_ctx
		arg__3540 = final_blk
		head__3538 = rt.LookupVar("ir.build", "add-terminator!").Deref()
		goto b19
	} else {
		arg__3539 = inner_ctx
		arg__3540 = final_blk
		head__3538 = rt.LookupVar("ir.build", "add-terminator!").Deref()
		goto b20
	}
b19:
	;
	arg__3547 = vm.NewArrayVector([]vm.Value{})
	goto b21
b20:
	;
	v353, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{last_val})
	if callErr != nil {
		return nil, callErr
	}
	arg__3547 = v353
	goto b21
b21:
	;
	_, callErr = rt.InvokeValue(head__3538, []vm.Value{arg__3539, arg__3540, vm.Keyword("return"), arg__3547, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	goto b15
}
func build_let(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var bindings vm.Value
	var body vm.Value
	var i int
	var ctx vm.Value
	var arg__3591 vm.Value
	var v42 bool
	var sym vm.Value
	var arg__3599 int
	var expr vm.Value
	var expr_id vm.Value
	var v55 int
	var fs vm.Value
	var last_id vm.Value
	var v91 vm.Value
	var v94 vm.Value
	var arg__3633 vm.Value
	var v100 vm.Value
	var result vm.Value
	var arg__3654 vm.Value
	var arg__3675 vm.Value
	var arg__3677 vm.Value
	var arg__3681 vm.Value
	var arg__3699 vm.Value
	var arg__3720 vm.Value
	var arg__3722 vm.Value
	var arg__3723 vm.Value
	var let_syms vm.Value
	var post_locals vm.Value
	var doseq_seq__3562 vm.Value
	var doseq_loop__3563 vm.Value
	var vec__3564 vm.Value
	var val vm.Value
	var arg__3758 vm.Value
	var and__x vm.Value
	var v355 vm.Value
	var arg__3773 vm.Value
	var v308 vm.Value
	var v311 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	bindings, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "push-locals!").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	ctx = arg1
	goto b1
b1:
	;
	arg__3591, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	v42 = rt.LtValue(vm.Int(i), arg__3591)
	if v42 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__3599 = i + 1
	expr, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, vm.Int(arg__3599)})
	if callErr != nil {
		return nil, callErr
	}
	expr_id, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{expr, ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{ctx, sym, expr_id})
	if callErr != nil {
		return nil, callErr
	}
	v55 = i + 2
	i = v55
	goto b1
b3:
	;
	goto b4
b4:
	;
	fs = body
	last_id = vm.NIL
	goto b5
b5:
	;
	v91, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v91) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	v94, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	arg__3633, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	v100, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{arg__3633, ctx})
	if callErr != nil {
		return nil, callErr
	}
	fs = v94
	last_id = v100
	goto b5
b7:
	;
	result = last_id
	goto b8
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__3654, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{vm.Int(0), arg__3654, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__3675, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__3677, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{vm.Int(0), arg__3675, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), arg__3677})
	if callErr != nil {
		return nil, callErr
	}
	arg__3681, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__3699, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{vm.Int(0), arg__3699, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__3720, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__3722, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{vm.Int(0), arg__3720, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	arg__3723, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), arg__3722})
	if callErr != nil {
		return nil, callErr
	}
	let_syms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__3681, arg__3723})
	if callErr != nil {
		return nil, callErr
	}
	post_locals, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "pop-locals!").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__3562, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{post_locals})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__3563 = doseq_seq__3562
	goto b9
b9:
	;
	if vm.IsTruthy(doseq_loop__3563) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	vec__3564, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__3563})
	if callErr != nil {
		return nil, callErr
	}
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__3564, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	val, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__3564, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(let_syms, []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__3758, callErr = rt.InvokeValue(let_syms, []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__3758})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b16
	} else {
		goto b17
	}
b11:
	;
	goto b12
b12:
	;
	return result, nil
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{ctx, sym, val})
	if callErr != nil {
		return nil, callErr
	}
	goto b15
b14:
	;
	goto b15
b15:
	;
	v355, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__3563})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__3563 = v355
	goto b9
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__3773, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, sym})
	if callErr != nil {
		return nil, callErr
	}
	v308, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{val, arg__3773})
	if callErr != nil {
		return nil, callErr
	}
	v311 = v308
	goto b18
b17:
	;
	v311 = and__x
	goto b18
b18:
	;
	if vm.IsTruthy(v311) {
		goto b13
	} else {
		goto b14
	}
}
func build_list(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var head vm.Value
	var builtin vm.Value
	var v17 bool
	var form vm.Value
	var ctx vm.Value
	var v20 vm.Value
	var v31 bool
	var v329 vm.Value
	var v34 vm.Value
	var v45 bool
	var v323 vm.Value
	var v48 vm.Value
	var v59 bool
	var v317 vm.Value
	var v62 vm.Value
	var v73 bool
	var v311 vm.Value
	var v76 vm.Value
	var v87 bool
	var v305 vm.Value
	var v90 vm.Value
	var v101 bool
	var v299 vm.Value
	var v104 vm.Value
	var v115 bool
	var v293 vm.Value
	var v118 vm.Value
	var v129 bool
	var v287 vm.Value
	var v132 vm.Value
	var v143 bool
	var v281 vm.Value
	var v146 vm.Value
	var v157 bool
	var v275 vm.Value
	var v160 vm.Value
	var v269 vm.Value
	var v172 vm.Value
	var or__x vm.Value
	var v263 vm.Value
	var fn_id vm.Value
	var arg__3901 vm.Value
	var v234 vm.Value
	var v257 vm.Value
	var v220 vm.Value
	var v211 vm.Value
	var v213 vm.Value
	var v247 vm.Value
	var v251 vm.Value
	var callErr error
	head, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	builtin, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.build", "builtin-ops").Deref(), head})
	if callErr != nil {
		return nil, callErr
	}
	v17 = head == vm.Symbol("if")
	if v17 {
		form = arg0
		ctx = arg1
		goto b1
	} else {
		form = arg0
		ctx = arg1
		goto b2
	}
b1:
	;
	v20, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-if").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v329 = v20
	goto b3
b2:
	;
	v31 = head == vm.Symbol("let")
	if v31 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v329, nil
b4:
	;
	v34, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-let").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v323 = v34
	goto b6
b5:
	;
	v45 = head == vm.Symbol("let*")
	if v45 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v329 = v323
	goto b3
b7:
	;
	v48, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-let").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v317 = v48
	goto b9
b8:
	;
	v59 = head == vm.Symbol("do")
	if v59 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v323 = v317
	goto b6
b10:
	;
	v62, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-do").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v311 = v62
	goto b12
b11:
	;
	v73 = head == vm.Symbol("quote")
	if v73 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v317 = v311
	goto b9
b13:
	;
	v76, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-quote").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v305 = v76
	goto b15
b14:
	;
	v87 = head == vm.Symbol("var")
	if v87 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v311 = v305
	goto b12
b16:
	;
	v90, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-var").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v299 = v90
	goto b18
b17:
	;
	v101 = head == vm.Symbol("set!")
	if v101 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v305 = v299
	goto b15
b19:
	;
	v104, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-set!").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v293 = v104
	goto b21
b20:
	;
	v115 = head == vm.Symbol("loop")
	if v115 {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v299 = v293
	goto b18
b22:
	;
	v118, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-loop").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v287 = v118
	goto b24
b23:
	;
	v129 = head == vm.Symbol("loop*")
	if v129 {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v293 = v287
	goto b21
b25:
	;
	v132, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-loop").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v281 = v132
	goto b27
b26:
	;
	v143 = head == vm.Symbol("recur")
	if v143 {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v287 = v281
	goto b24
b28:
	;
	v146, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-recur").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v275 = v146
	goto b30
b29:
	;
	v157 = head == vm.Symbol("fn*")
	if v157 {
		goto b31
	} else {
		goto b32
	}
b30:
	;
	v281 = v275
	goto b27
b31:
	;
	v160, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn*").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v269 = v160
	goto b33
b32:
	;
	if vm.IsTruthy(builtin) {
		goto b34
	} else {
		goto b35
	}
b33:
	;
	v275 = v269
	goto b30
b34:
	;
	v172, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-builtin-op").Deref(), []vm.Value{builtin, form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v263 = v172
	goto b36
b35:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{head})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b40
	} else {
		goto b41
	}
b36:
	;
	v269 = v263
	goto b33
b37:
	;
	fn_id, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{head, ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__3901, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v234, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-call-with-head").Deref(), []vm.Value{fn_id, arg__3901, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v257 = v234
	goto b39
b38:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b46
	} else {
		goto b47
	}
b39:
	;
	v263 = v257
	goto b36
b40:
	;
	v220 = or__x
	goto b42
b41:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "number?").Deref(), []vm.Value{head})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b43
	} else {
		goto b44
	}
b42:
	;
	if vm.IsTruthy(v220) {
		goto b37
	} else {
		goto b38
	}
b43:
	;
	v213 = or__x
	goto b45
b44:
	;
	v211, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{head})
	if callErr != nil {
		return nil, callErr
	}
	v213 = v211
	goto b45
b45:
	;
	v220 = v213
	goto b42
b46:
	;
	v247, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-call").Deref(), []vm.Value{form, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v251 = v247
	goto b48
b47:
	;
	v251 = vm.NIL
	goto b48
b48:
	;
	v257 = v251
	goto b39
}
func build_loop(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var bindings vm.Value
	var body vm.Value
	var f vm.Value
	var arg__3948 vm.Value
	var n_slots vm.Value
	var arg__3959 vm.Value
	var arg__3970 vm.Value
	var arg__3980 vm.Value
	var arg__3991 vm.Value
	var arg__3992 vm.Value
	var known vm.Value
	var arg__4030 vm.Value
	var arg__4071 vm.Value
	var arg__4073 vm.Value
	var body_caps vm.Value
	var n_caps vm.Value
	var header vm.Value
	var arg__4320 vm.Value
	var loop_param_ids vm.Value
	var arg__4505 vm.Value
	var cap_param_ids vm.Value
	var arg__4570 vm.Value
	var init_vals vm.Value
	var cap_vals vm.Value
	var arg__4599 vm.Value
	var arg__4615 vm.Value
	var arg__4616 vm.Value
	var bt vm.Value
	var entry_end vm.Value
	var arg__4654 vm.Value
	var doseq_seq__3909 vm.Value
	var doseq_loop__3910 vm.Value
	var ctx vm.Value
	var i vm.Value
	var arg__4662 vm.Value
	var arg__4686 vm.Value
	var arg__4692 vm.Value
	var v254 vm.Value
	var arg__4712 vm.Value
	var doseq_seq__3911 vm.Value
	var doseq_loop__3912 vm.Value
	var vec__3913 vm.Value
	var sym vm.Value
	var pid vm.Value
	var v358 vm.Value
	var arg__4756 vm.Value
	var arg__4788 vm.Value
	var arg__4796 vm.Value
	var arg__4822 vm.Value
	var arg__4854 vm.Value
	var arg__4862 vm.Value
	var arg__4870 vm.Value
	var pre_locals vm.Value
	var fs vm.Value
	var last_val vm.Value
	var v522 vm.Value
	var v525 vm.Value
	var arg__4889 vm.Value
	var v531 vm.Value
	var result vm.Value
	var post_locals vm.Value
	var doseq_seq__3914 vm.Value
	var doseq_loop__3915 vm.Value
	var vec__3916 vm.Value
	var val vm.Value
	var and__x vm.Value
	var arg__4961 vm.Value
	var arg__4987 vm.Value
	var arg__4993 vm.Value
	var arg__5015 vm.Value
	var arg__5041 vm.Value
	var arg__5047 vm.Value
	var arg__5053 vm.Value
	var v858 vm.Value
	var arg__4936 vm.Value
	var v781 vm.Value
	var v784 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	bindings, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), arg0})
	if callErr != nil {
		return nil, callErr
	}
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__3948, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{bindings})
	if callErr != nil {
		return nil, callErr
	}
	n_slots, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "/").Deref(), []vm.Value{arg__3948, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3959, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__3959, vm.Keyword("loop-header")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3970, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__3970, vm.Keyword("loop-capture-syms")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3980, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{arg__3980})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3991, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__3992, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{arg__3991})
	if callErr != nil {
		return nil, callErr
	}
	known, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "set").Deref(), []vm.Value{arg__3992})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__4030, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var arg__4027 vm.Value
		var v8 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg1, known})
		if callErr != nil {
			return nil, callErr
		}
		arg__4027, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg1, known})
		if callErr != nil {
			return nil, callErr
		}
		v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg0, arg__4027})
		if callErr != nil {
			return nil, callErr
		}
		return v8, nil
	}), arg__4030, body})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__4071, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__4073, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var arg__4068 vm.Value
		var v8 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg1, known})
		if callErr != nil {
			return nil, callErr
		}
		arg__4068, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg1, known})
		if callErr != nil {
			return nil, callErr
		}
		v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg0, arg__4068})
		if callErr != nil {
			return nil, callErr
		}
		return v8, nil
	}), arg__4071, body})
	if callErr != nil {
		return nil, callErr
	}
	body_caps, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__4073})
	if callErr != nil {
		return nil, callErr
	}
	n_caps, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{body_caps})
	if callErr != nil {
		return nil, callErr
	}
	header, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-block").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{n_slots})
	if callErr != nil {
		return nil, callErr
	}
	arg__4320, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{n_slots})
	if callErr != nil {
		return nil, callErr
	}
	loop_param_ids, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var pid vm.Value
		var arg__4229 vm.Value
		var arg__4240 vm.Value
		var arg__4241 vm.Value
		var arg__4252 vm.Value
		var arg__4263 vm.Value
		var arg__4264 vm.Value
		var arg__4265 vm.Value
		var arg__4278 vm.Value
		var arg__4289 vm.Value
		var arg__4290 vm.Value
		var arg__4301 vm.Value
		var arg__4312 vm.Value
		var arg__4313 vm.Value
		var arg__4314 vm.Value
		var arg__4315 vm.Value
		var callErr error
		pid, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{arg1, header, vm.Keyword("block-arg"), vm.NewArrayVector([]vm.Value{}), arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-block-param!").Deref(), []vm.Value{f, header, pid})
		if callErr != nil {
			return nil, callErr
		}
		arg__4229 = rt.MulValue(arg0, vm.Int(2))
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4229})
		if callErr != nil {
			return nil, callErr
		}
		arg__4240 = rt.MulValue(arg0, vm.Int(2))
		arg__4241, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4240})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__4241})
		if callErr != nil {
			return nil, callErr
		}
		arg__4252 = rt.MulValue(arg0, vm.Int(2))
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4252})
		if callErr != nil {
			return nil, callErr
		}
		arg__4263 = rt.MulValue(arg0, vm.Int(2))
		arg__4264, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4263})
		if callErr != nil {
			return nil, callErr
		}
		arg__4265, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__4264})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-named-source-info").Deref(), []vm.Value{arg__4265})
		if callErr != nil {
			return nil, callErr
		}
		arg__4278 = rt.MulValue(arg0, vm.Int(2))
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4278})
		if callErr != nil {
			return nil, callErr
		}
		arg__4289 = rt.MulValue(arg0, vm.Int(2))
		arg__4290, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4289})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__4290})
		if callErr != nil {
			return nil, callErr
		}
		arg__4301 = rt.MulValue(arg0, vm.Int(2))
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4301})
		if callErr != nil {
			return nil, callErr
		}
		arg__4312 = rt.MulValue(arg0, vm.Int(2))
		arg__4313, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4312})
		if callErr != nil {
			return nil, callErr
		}
		arg__4314, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__4313})
		if callErr != nil {
			return nil, callErr
		}
		arg__4315, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-named-source-info").Deref(), []vm.Value{arg__4314})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-source-info!").Deref(), []vm.Value{f, pid, arg__4315})
		if callErr != nil {
			return nil, callErr
		}
		return pid, nil
	}), arg__4320})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{n_caps})
	if callErr != nil {
		return nil, callErr
	}
	arg__4505, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{n_caps})
	if callErr != nil {
		return nil, callErr
	}
	cap_param_ids, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__4428 vm.Value
		var pid vm.Value
		var arg__4450 vm.Value
		var arg__4465 vm.Value
		var arg__4466 vm.Value
		var arg__4483 vm.Value
		var arg__4498 vm.Value
		var arg__4499 vm.Value
		var arg__4500 vm.Value
		var callErr error
		arg__4428 = rt.AddValue(n_slots, arg0)
		pid, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{arg1, header, vm.Keyword("block-arg"), vm.NewArrayVector([]vm.Value{}), arg__4428})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-block-param!").Deref(), []vm.Value{f, header, pid})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{body_caps, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__4450, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{body_caps, arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__4450})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{body_caps, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__4465, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{body_caps, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__4466, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__4465})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-named-source-info").Deref(), []vm.Value{arg__4466})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{body_caps, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__4483, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{body_caps, arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__4483})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{body_caps, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__4498, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{body_caps, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__4499, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__4498})
		if callErr != nil {
			return nil, callErr
		}
		arg__4500, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-named-source-info").Deref(), []vm.Value{arg__4499})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-source-info!").Deref(), []vm.Value{f, pid, arg__4500})
		if callErr != nil {
			return nil, callErr
		}
		return pid, nil
	}), arg__4505})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{n_slots})
	if callErr != nil {
		return nil, callErr
	}
	arg__4570, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{n_slots})
	if callErr != nil {
		return nil, callErr
	}
	init_vals, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__4548 vm.Value
		var arg__4549 vm.Value
		var arg__4562 vm.Value
		var arg__4563 vm.Value
		var arg__4564 vm.Value
		var v20 vm.Value
		var callErr error
		arg__4548 = rt.MulValue(arg0, vm.Int(2))
		arg__4549 = rt.AddValue(arg__4548, vm.Int(1))
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4549})
		if callErr != nil {
			return nil, callErr
		}
		arg__4562 = rt.MulValue(arg0, vm.Int(2))
		arg__4563 = rt.AddValue(arg__4562, vm.Int(1))
		arg__4564, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4563})
		if callErr != nil {
			return nil, callErr
		}
		v20, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{arg__4564, arg1})
		if callErr != nil {
			return nil, callErr
		}
		return v20, nil
	}), arg__4570})
	if callErr != nil {
		return nil, callErr
	}
	cap_vals, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{arg1, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body_caps})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{init_vals, cap_vals})
	if callErr != nil {
		return nil, callErr
	}
	arg__4599, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{init_vals, cap_vals})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__4599})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{init_vals, cap_vals})
	if callErr != nil {
		return nil, callErr
	}
	arg__4615, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{init_vals, cap_vals})
	if callErr != nil {
		return nil, callErr
	}
	arg__4616, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__4615})
	if callErr != nil {
		return nil, callErr
	}
	bt, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-branch-target").Deref(), []vm.Value{header, arg__4616})
	if callErr != nil {
		return nil, callErr
	}
	entry_end, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-terminator!").Deref(), []vm.Value{arg1, entry_end, vm.Keyword("branch"), vm.NewArrayVector([]vm.Value{}), bt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-pred!").Deref(), []vm.Value{f, header, entry_end})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-set-block!").Deref(), []vm.Value{arg1, header})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "push-locals!").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{n_slots})
	if callErr != nil {
		return nil, callErr
	}
	arg__4654, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{n_slots})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__3909, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__4654})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__3910 = doseq_seq__3909
	ctx = arg1
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__3910) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	i, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__3910})
	if callErr != nil {
		return nil, callErr
	}
	arg__4662 = rt.MulValue(i, vm.Int(2))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4662})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{loop_param_ids, i})
	if callErr != nil {
		return nil, callErr
	}
	arg__4686, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{bindings, arg__4662})
	if callErr != nil {
		return nil, callErr
	}
	arg__4692, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{loop_param_ids, i})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{ctx, arg__4686, arg__4692})
	if callErr != nil {
		return nil, callErr
	}
	v254, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__3910})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__3910 = v254
	goto b1
b3:
	;
	goto b4
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("clojure.core", "vector").Deref(), body_caps, cap_param_ids})
	if callErr != nil {
		return nil, callErr
	}
	arg__4712, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("clojure.core", "vector").Deref(), body_caps, cap_param_ids})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__3911, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__4712})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__3912 = doseq_seq__3911
	goto b5
b5:
	;
	if vm.IsTruthy(doseq_loop__3912) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	vec__3913, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__3912})
	if callErr != nil {
		return nil, callErr
	}
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__3913, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	pid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__3913, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "rebind-local!").Deref(), []vm.Value{ctx, sym, pid})
	if callErr != nil {
		return nil, callErr
	}
	v358, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__3912})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__3912 = v358
	goto b5
b7:
	;
	goto b8
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4756, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4756, vm.Keyword("loop-headers"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var header_4 vm.Value
		var header_7 vm.Value
		var header_15 vm.Value
		var or__x vm.Value
		var header_19 vm.Value
		var head__4760 vm.Value
		var header_23 vm.Value
		var arg__4761 vm.Value
		var header_32 vm.Value
		var v34 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			header_4 = header
			goto b1
		} else {
			s = arg0
			header_7 = header
			goto b2
		}
	b1:
		;
		header_15 = header_4
		goto b3
	b2:
		;
		header_15 = header_7
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			header_19 = header_15
			head__4760 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b4
		} else {
			header_23 = header_15
			head__4760 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b5
		}
	b4:
		;
		arg__4761 = or__x
		header_32 = header_19
		goto b6
	b5:
		;
		arg__4761 = vm.NewArrayVector([]vm.Value{})
		header_32 = header_23
		goto b6
	b6:
		;
		v34, callErr = rt.InvokeValue(head__4760, []vm.Value{arg__4761, header_32})
		if callErr != nil {
			return nil, callErr
		}
		return v34, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4788, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4796, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4788, vm.Keyword("loop-headers"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var header_4 vm.Value
		var header_7 vm.Value
		var header_15 vm.Value
		var or__x vm.Value
		var header_19 vm.Value
		var head__4792 vm.Value
		var header_23 vm.Value
		var arg__4793 vm.Value
		var header_32 vm.Value
		var v34 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			header_4 = header
			goto b1
		} else {
			s = arg0
			header_7 = header
			goto b2
		}
	b1:
		;
		header_15 = header_4
		goto b3
	b2:
		;
		header_15 = header_7
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			header_19 = header_15
			head__4792 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b4
		} else {
			header_23 = header_15
			head__4792 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b5
		}
	b4:
		;
		arg__4793 = or__x
		header_32 = header_19
		goto b6
	b5:
		;
		arg__4793 = vm.NewArrayVector([]vm.Value{})
		header_32 = header_23
		goto b6
	b6:
		;
		v34, callErr = rt.InvokeValue(head__4792, []vm.Value{arg__4793, header_32})
		if callErr != nil {
			return nil, callErr
		}
		return v34, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4796, vm.Keyword("loop-capture-syms-stack"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var body_caps_4 vm.Value
		var body_caps_7 vm.Value
		var body_caps_15 vm.Value
		var or__x vm.Value
		var body_caps_19 vm.Value
		var head__4800 vm.Value
		var body_caps_23 vm.Value
		var arg__4801 vm.Value
		var body_caps_32 vm.Value
		var v34 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			body_caps_4 = body_caps
			goto b1
		} else {
			s = arg0
			body_caps_7 = body_caps
			goto b2
		}
	b1:
		;
		body_caps_15 = body_caps_4
		goto b3
	b2:
		;
		body_caps_15 = body_caps_7
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			body_caps_19 = body_caps_15
			head__4800 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b4
		} else {
			body_caps_23 = body_caps_15
			head__4800 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b5
		}
	b4:
		;
		arg__4801 = or__x
		body_caps_32 = body_caps_19
		goto b6
	b5:
		;
		arg__4801 = vm.NewArrayVector([]vm.Value{})
		body_caps_32 = body_caps_23
		goto b6
	b6:
		;
		v34, callErr = rt.InvokeValue(head__4800, []vm.Value{arg__4801, body_caps_32})
		if callErr != nil {
			return nil, callErr
		}
		return v34, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4822, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4822, vm.Keyword("loop-headers"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var header_4 vm.Value
		var header_7 vm.Value
		var header_15 vm.Value
		var or__x vm.Value
		var header_19 vm.Value
		var head__4826 vm.Value
		var header_23 vm.Value
		var arg__4827 vm.Value
		var header_32 vm.Value
		var v34 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			header_4 = header
			goto b1
		} else {
			s = arg0
			header_7 = header
			goto b2
		}
	b1:
		;
		header_15 = header_4
		goto b3
	b2:
		;
		header_15 = header_7
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			header_19 = header_15
			head__4826 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b4
		} else {
			header_23 = header_15
			head__4826 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b5
		}
	b4:
		;
		arg__4827 = or__x
		header_32 = header_19
		goto b6
	b5:
		;
		arg__4827 = vm.NewArrayVector([]vm.Value{})
		header_32 = header_23
		goto b6
	b6:
		;
		v34, callErr = rt.InvokeValue(head__4826, []vm.Value{arg__4827, header_32})
		if callErr != nil {
			return nil, callErr
		}
		return v34, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4854, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4862, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4854, vm.Keyword("loop-headers"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var header_4 vm.Value
		var header_7 vm.Value
		var header_15 vm.Value
		var or__x vm.Value
		var header_19 vm.Value
		var head__4858 vm.Value
		var header_23 vm.Value
		var arg__4859 vm.Value
		var header_32 vm.Value
		var v34 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			header_4 = header
			goto b1
		} else {
			s = arg0
			header_7 = header
			goto b2
		}
	b1:
		;
		header_15 = header_4
		goto b3
	b2:
		;
		header_15 = header_7
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			header_19 = header_15
			head__4858 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b4
		} else {
			header_23 = header_15
			head__4858 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b5
		}
	b4:
		;
		arg__4859 = or__x
		header_32 = header_19
		goto b6
	b5:
		;
		arg__4859 = vm.NewArrayVector([]vm.Value{})
		header_32 = header_23
		goto b6
	b6:
		;
		v34, callErr = rt.InvokeValue(head__4858, []vm.Value{arg__4859, header_32})
		if callErr != nil {
			return nil, callErr
		}
		return v34, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	arg__4870, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4862, vm.Keyword("loop-capture-syms-stack"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var body_caps_4 vm.Value
		var body_caps_7 vm.Value
		var body_caps_15 vm.Value
		var or__x vm.Value
		var body_caps_19 vm.Value
		var head__4866 vm.Value
		var body_caps_23 vm.Value
		var arg__4867 vm.Value
		var body_caps_32 vm.Value
		var v34 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			body_caps_4 = body_caps
			goto b1
		} else {
			s = arg0
			body_caps_7 = body_caps
			goto b2
		}
	b1:
		;
		body_caps_15 = body_caps_4
		goto b3
	b2:
		;
		body_caps_15 = body_caps_7
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			body_caps_19 = body_caps_15
			head__4866 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b4
		} else {
			body_caps_23 = body_caps_15
			head__4866 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b5
		}
	b4:
		;
		arg__4867 = or__x
		body_caps_32 = body_caps_19
		goto b6
	b5:
		;
		arg__4867 = vm.NewArrayVector([]vm.Value{})
		body_caps_32 = body_caps_23
		goto b6
	b6:
		;
		v34, callErr = rt.InvokeValue(head__4866, []vm.Value{arg__4867, body_caps_32})
		if callErr != nil {
			return nil, callErr
		}
		return v34, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{ctx, arg__4870})
	if callErr != nil {
		return nil, callErr
	}
	pre_locals, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	fs = body
	last_val = vm.NIL
	goto b9
b9:
	;
	v522, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v522) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	v525, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	arg__4889, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{fs})
	if callErr != nil {
		return nil, callErr
	}
	v531, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{arg__4889, ctx})
	if callErr != nil {
		return nil, callErr
	}
	fs = v525
	last_val = v531
	goto b9
b11:
	;
	result = last_val
	goto b12
b12:
	;
	post_locals, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "current-locals-flat").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "pop-locals!").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__3914, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{post_locals})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__3915 = doseq_seq__3914
	goto b13
b13:
	;
	if vm.IsTruthy(doseq_loop__3915) {
		goto b14
	} else {
		goto b15
	}
b14:
	;
	vec__3916, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__3915})
	if callErr != nil {
		return nil, callErr
	}
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__3916, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	val, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__3916, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{pre_locals, sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b20
	} else {
		goto b21
	}
b15:
	;
	goto b16
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4961, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4961, vm.Keyword("loop-headers"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var or__x vm.Value
		var head__4964 vm.Value
		var arg__4965 vm.Value
		var v27 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			goto b1
		} else {
			s = arg0
			goto b2
		}
	b1:
		;
		goto b3
	b2:
		;
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			head__4964 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b4
		} else {
			head__4964 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b5
		}
	b4:
		;
		arg__4965 = or__x
		goto b6
	b5:
		;
		arg__4965 = vm.NewArrayVector([]vm.Value{})
		goto b6
	b6:
		;
		v27, callErr = rt.InvokeValue(head__4964, []vm.Value{arg__4965})
		if callErr != nil {
			return nil, callErr
		}
		return v27, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4987, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__4993, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4987, vm.Keyword("loop-headers"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var or__x vm.Value
		var head__4990 vm.Value
		var arg__4991 vm.Value
		var v27 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			goto b1
		} else {
			s = arg0
			goto b2
		}
	b1:
		;
		goto b3
	b2:
		;
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			head__4990 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b4
		} else {
			head__4990 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b5
		}
	b4:
		;
		arg__4991 = or__x
		goto b6
	b5:
		;
		arg__4991 = vm.NewArrayVector([]vm.Value{})
		goto b6
	b6:
		;
		v27, callErr = rt.InvokeValue(head__4990, []vm.Value{arg__4991})
		if callErr != nil {
			return nil, callErr
		}
		return v27, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__4993, vm.Keyword("loop-capture-syms-stack"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var or__x vm.Value
		var head__4996 vm.Value
		var arg__4997 vm.Value
		var v27 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			goto b1
		} else {
			s = arg0
			goto b2
		}
	b1:
		;
		goto b3
	b2:
		;
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			head__4996 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b4
		} else {
			head__4996 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b5
		}
	b4:
		;
		arg__4997 = or__x
		goto b6
	b5:
		;
		arg__4997 = vm.NewArrayVector([]vm.Value{})
		goto b6
	b6:
		;
		v27, callErr = rt.InvokeValue(head__4996, []vm.Value{arg__4997})
		if callErr != nil {
			return nil, callErr
		}
		return v27, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5015, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__5015, vm.Keyword("loop-headers"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var or__x vm.Value
		var head__5018 vm.Value
		var arg__5019 vm.Value
		var v27 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			goto b1
		} else {
			s = arg0
			goto b2
		}
	b1:
		;
		goto b3
	b2:
		;
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			head__5018 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b4
		} else {
			head__5018 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b5
		}
	b4:
		;
		arg__5019 = or__x
		goto b6
	b5:
		;
		arg__5019 = vm.NewArrayVector([]vm.Value{})
		goto b6
	b6:
		;
		v27, callErr = rt.InvokeValue(head__5018, []vm.Value{arg__5019})
		if callErr != nil {
			return nil, callErr
		}
		return v27, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5041, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5047, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__5041, vm.Keyword("loop-headers"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var or__x vm.Value
		var head__5044 vm.Value
		var arg__5045 vm.Value
		var v27 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			goto b1
		} else {
			s = arg0
			goto b2
		}
	b1:
		;
		goto b3
	b2:
		;
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			head__5044 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b4
		} else {
			head__5044 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b5
		}
	b4:
		;
		arg__5045 = or__x
		goto b6
	b5:
		;
		arg__5045 = vm.NewArrayVector([]vm.Value{})
		goto b6
	b6:
		;
		v27, callErr = rt.InvokeValue(head__5044, []vm.Value{arg__5045})
		if callErr != nil {
			return nil, callErr
		}
		return v27, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	arg__5053, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "update").Deref(), []vm.Value{arg__5047, vm.Keyword("loop-capture-syms-stack"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var s vm.Value
		var or__x vm.Value
		var head__5050 vm.Value
		var arg__5051 vm.Value
		var v27 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			s = arg0
			goto b1
		} else {
			s = arg0
			goto b2
		}
	b1:
		;
		goto b3
	b2:
		;
		goto b3
	b3:
		;
		if vm.IsTruthy(s) {
			or__x = s
			head__5050 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b4
		} else {
			head__5050 = rt.LookupVar("clojure.core", "pop").Deref()
			goto b5
		}
	b4:
		;
		arg__5051 = or__x
		goto b6
	b5:
		;
		arg__5051 = vm.NewArrayVector([]vm.Value{})
		goto b6
	b6:
		;
		v27, callErr = rt.InvokeValue(head__5050, []vm.Value{arg__5051})
		if callErr != nil {
			return nil, callErr
		}
		return v27, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{ctx, arg__5053})
	if callErr != nil {
		return nil, callErr
	}
	return result, nil
b17:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{ctx, sym, val})
	if callErr != nil {
		return nil, callErr
	}
	goto b19
b18:
	;
	goto b19
b19:
	;
	v858, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__3915})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__3915 = v858
	goto b13
b20:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{pre_locals, sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__4936, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{pre_locals, sym})
	if callErr != nil {
		return nil, callErr
	}
	v781, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{val, arg__4936})
	if callErr != nil {
		return nil, callErr
	}
	v784 = v781
	goto b22
b21:
	;
	v784 = and__x
	goto b22
b22:
	;
	if vm.IsTruthy(v784) {
		goto b17
	} else {
		goto b18
	}
}
func build_map(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__5062 vm.Value
	var v12 vm.Value
	var form vm.Value
	var ctx vm.Value
	var arg__5076 vm.Value
	var v23 vm.Value
	var arg__5126 vm.Value
	var all_const_QMARK_ vm.Value
	var arg__5155 vm.Value
	var arg__5186 vm.Value
	var arg__5187 vm.Value
	var pairs vm.Value
	var v98 vm.Value
	var arg__5201 vm.Value
	var v69 vm.Value
	var arg__5221 vm.Value
	var arg__5227 vm.Value
	var fn_id vm.Value
	var v90 vm.Value
	var v92 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__5062, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{arg__5062})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		form = arg0
		ctx = arg1
		goto b1
	} else {
		form = arg0
		ctx = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5076, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v23, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5076, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v98 = v23
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__5126, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	all_const_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__5112 vm.Value
		var and__x vm.Value
		var e vm.Value
		var arg__5121 vm.Value
		var v17 vm.Value
		var v20 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__5112, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "is-literal?").Deref(), []vm.Value{arg__5112})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			e = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{e})
		if callErr != nil {
			return nil, callErr
		}
		arg__5121, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{e})
		if callErr != nil {
			return nil, callErr
		}
		v17, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "is-literal?").Deref(), []vm.Value{arg__5121})
		if callErr != nil {
			return nil, callErr
		}
		v20 = v17
		goto b3
	b2:
		;
		v20 = and__x
		goto b3
	b3:
		;
		return v20, nil
	}), arg__5126})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__5155, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__5146 vm.Value
		var arg__5150 vm.Value
		var v6 vm.Value
		var callErr error
		arg__5146, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__5150, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__5146, arg__5150})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__5155})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__5186, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__5187, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__5177 vm.Value
		var arg__5181 vm.Value
		var v6 vm.Value
		var callErr error
		arg__5177, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__5181, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__5177, arg__5181})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__5186})
	if callErr != nil {
		return nil, callErr
	}
	pairs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__5187})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(all_const_QMARK_) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v98, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5201, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v69, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5201, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v92 = v69
	goto b6
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "resolve").Deref(), []vm.Value{vm.Symbol("array-map")})
	if callErr != nil {
		return nil, callErr
	}
	arg__5221, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5227, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "resolve").Deref(), []vm.Value{vm.Symbol("array-map")})
	if callErr != nil {
		return nil, callErr
	}
	fn_id, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5221, vm.Keyword("load-var"), vm.NewArrayVector([]vm.Value{}), arg__5227})
	if callErr != nil {
		return nil, callErr
	}
	v90, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-call-with-head").Deref(), []vm.Value{fn_id, pairs, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v92 = v90
	goto b6
b6:
	;
	v98 = v92
	goto b3
}
func build_quote(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var x vm.Value
	var arg__5263 vm.Value
	var v24 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__5263, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	v24, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{arg1, arg__5263, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), x})
	if callErr != nil {
		return nil, callErr
	}
	return v24, nil
}
func build_recur(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__5276 vm.Value
	var headers vm.Value
	var f vm.Value
	var form vm.Value
	var ctx vm.Value
	var header vm.Value
	var arg__5296 vm.Value
	var cap_syms_stack vm.Value
	var or__x vm.Value
	var arg__5385 vm.Value
	var arg_syms vm.Value
	var v176 vm.Value
	var v33 vm.Value
	var and__x vm.Value
	var v36 vm.Value
	var cap_syms vm.Value
	var cap_vals vm.Value
	var arg__5336 vm.Value
	var loop_arg_ids vm.Value
	var arg__5349 vm.Value
	var all_args vm.Value
	var bt vm.Value
	var cur vm.Value
	var arg__5407 vm.Value
	var arg_ids vm.Value
	var arg__5427 vm.Value
	var v167 vm.Value
	var v169 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__5276, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	headers, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__5276, vm.Keyword("loop-headers")})
	if callErr != nil {
		return nil, callErr
	}
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(headers) {
		form = arg0
		ctx = arg1
		goto b4
	} else {
		form = arg0
		ctx = arg1
		and__x = headers
		goto b5
	}
b1:
	;
	header, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "peek").Deref(), []vm.Value{headers})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5296, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	cap_syms_stack, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__5296, vm.Keyword("loop-capture-syms-stack")})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "peek").Deref(), []vm.Value{cap_syms_stack})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b7
	} else {
		goto b8
	}
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5385, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg_syms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__5385, vm.Keyword("fn-arg-syms")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(arg_syms) {
		goto b10
	} else {
		goto b11
	}
b3:
	;
	return v176, nil
b4:
	;
	v33, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{headers})
	if callErr != nil {
		return nil, callErr
	}
	v36 = v33
	goto b6
b5:
	;
	v36 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v36) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	cap_syms = or__x
	goto b9
b8:
	;
	cap_syms = vm.NewArrayVector([]vm.Value{})
	goto b9
b9:
	;
	cap_vals, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), cap_syms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__5336, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	loop_arg_ids, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{arg0, ctx})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), arg__5336})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{loop_arg_ids, cap_vals})
	if callErr != nil {
		return nil, callErr
	}
	arg__5349, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{loop_arg_ids, cap_vals})
	if callErr != nil {
		return nil, callErr
	}
	all_args, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__5349})
	if callErr != nil {
		return nil, callErr
	}
	bt, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-branch-target").Deref(), []vm.Value{header, all_args})
	if callErr != nil {
		return nil, callErr
	}
	cur, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-terminator!").Deref(), []vm.Value{f, cur, vm.Keyword("branch"), vm.NewArrayVector([]vm.Value{}), bt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-pred!").Deref(), []vm.Value{f, header, cur})
	if callErr != nil {
		return nil, callErr
	}
	v176 = rt.LookupVar("ir.build", "TERMINATED").Deref()
	goto b3
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__5407, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg_ids, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{arg0, ctx})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), arg__5407})
	if callErr != nil {
		return nil, callErr
	}
	cur, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg_ids})
	if callErr != nil {
		return nil, callErr
	}
	arg__5427, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg_ids})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-terminator!").Deref(), []vm.Value{f, cur, vm.Keyword("tail-call"), arg_ids, arg__5427})
	if callErr != nil {
		return nil, callErr
	}
	v169 = rt.LookupVar("ir.build", "TERMINATED").Deref()
	goto b12
b11:
	;
	v167, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{vm.String("recur outside of loop or function")})
	if callErr != nil {
		return nil, callErr
	}
	v169 = v167
	goto b12
b12:
	;
	v176 = v169
	goto b3
}
func build_set_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var sym vm.Value
	var val vm.Value
	var v vm.Value
	var v38 vm.Value
	var ctx vm.Value
	var arg__5471 vm.Value
	var arg__5485 vm.Value
	var var_nid vm.Value
	var val_nid vm.Value
	var arg__5510 vm.Value
	var arg__5515 vm.Value
	var v86 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	val, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(2), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "resolve").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	v38, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v38) {
		ctx = arg1
		goto b1
	} else {
		ctx = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/build: set! can't resolve "), sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__5471, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/build: set! can't resolve "), sym})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__5471})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5485, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	var_nid, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5485, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), v})
	if callErr != nil {
		return nil, callErr
	}
	val_nid, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-form").Deref(), []vm.Value{val, ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{var_nid, val_nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__5510, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5515, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{var_nid, val_nid})
	if callErr != nil {
		return nil, callErr
	}
	v86, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5510, vm.Keyword("set-var"), arg__5515, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	return v86, nil
}
func build_single_fn_STAR_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var expanded vm.Value
	var name_sym vm.Value
	var body_forms vm.Value
	var ctx vm.Value
	var v19 vm.Value
	var args_vec vm.Value
	var flat_args vm.Value
	var v42 vm.Value
	var flat_body vm.Value
	var v68 vm.Value
	var variadic_QMARK_ vm.Value
	var arg__5533 vm.Value
	var arg_set vm.Value
	var arg__5548 vm.Value
	var frees vm.Value
	var arg__5584 vm.Value
	var arg__5621 vm.Value
	var arg__5622 vm.Value
	var captures vm.Value
	var template vm.Value
	var v141 vm.Value
	var callErr error
	expanded, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-fn-args").Deref(), []vm.Value{arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(expanded) {
		name_sym = arg0
		body_forms = arg2
		ctx = arg3
		goto b1
	} else {
		name_sym = arg0
		args_vec = arg1
		body_forms = arg2
		ctx = arg3
		goto b2
	}
b1:
	;
	v19, callErr = rt.InvokeValue(vm.Keyword("flat-args"), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	flat_args = v19
	goto b3
b2:
	;
	flat_args = args_vec
	goto b3
b3:
	;
	if vm.IsTruthy(expanded) {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	v42, callErr = rt.InvokeValue(vm.Keyword("body"), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	flat_body = v42
	goto b6
b5:
	;
	flat_body = body_forms
	goto b6
b6:
	;
	if vm.IsTruthy(expanded) {
		goto b7
	} else {
		goto b8
	}
b7:
	;
	v68, callErr = rt.InvokeValue(vm.Keyword("variadic?"), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	variadic_QMARK_ = v68
	goto b9
b8:
	;
	variadic_QMARK_ = vm.Boolean(false)
	goto b9
b9:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__5533, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg_set, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5533, flat_args})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{vm.Symbol("do"), flat_body})
	if callErr != nil {
		return nil, callErr
	}
	arg__5548, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{vm.Symbol("do"), flat_body})
	if callErr != nil {
		return nil, callErr
	}
	frees, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg__5548, arg_set})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), frees})
	if callErr != nil {
		return nil, callErr
	}
	arg__5584, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), frees})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__5584})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), frees})
	if callErr != nil {
		return nil, callErr
	}
	arg__5621, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{ctx, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), frees})
	if callErr != nil {
		return nil, callErr
	}
	arg__5622, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort-by").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__5621})
	if callErr != nil {
		return nil, callErr
	}
	captures, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__5622})
	if callErr != nil {
		return nil, callErr
	}
	template, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-inner-fn-template").Deref(), []vm.Value{name_sym, flat_args, flat_body, captures, variadic_QMARK_})
	if callErr != nil {
		return nil, callErr
	}
	v141, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "emit-template-closure").Deref(), []vm.Value{template, captures, ctx})
	if callErr != nil {
		return nil, callErr
	}
	return v141, nil
}
func build_symbol(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var local vm.Value
	var sym vm.Value
	var ctx vm.Value
	var v vm.Value
	var v23 vm.Value
	var v53 vm.Value
	var arg__5664 vm.Value
	var v34 vm.Value
	var arg__5678 vm.Value
	var v45 vm.Value
	var v47 vm.Value
	var callErr error
	local, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "lookup-local").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(local) {
		goto b1
	} else {
		sym = arg0
		ctx = arg1
		goto b2
	}
b1:
	;
	v53 = local
	goto b3
b2:
	;
	v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "resolve").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	v23, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v23) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v53, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/build: unresolved symbol "), sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__5664, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/build: unresolved symbol "), sym})
	if callErr != nil {
		return nil, callErr
	}
	v34, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__5664})
	if callErr != nil {
		return nil, callErr
	}
	v47 = v34
	goto b6
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5678, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v45, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5678, vm.Keyword("load-var"), vm.NewArrayVector([]vm.Value{}), v})
	if callErr != nil {
		return nil, callErr
	}
	v47 = v45
	goto b6
b6:
	;
	v53 = v47
	goto b3
}
func build_var(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var sym vm.Value
	var v vm.Value
	var v30 vm.Value
	var arg__5715 vm.Value
	var v41 vm.Value
	var ctx vm.Value
	var arg__5729 vm.Value
	var v52 vm.Value
	var v54 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "resolve").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	v30, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v30) {
		goto b1
	} else {
		ctx = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/build: can't resolve var "), sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__5715, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/build: can't resolve var "), sym})
	if callErr != nil {
		return nil, callErr
	}
	v41, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__5715})
	if callErr != nil {
		return nil, callErr
	}
	v54 = v41
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5729, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v52, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5729, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), v})
	if callErr != nil {
		return nil, callErr
	}
	v54 = v52
	goto b3
b3:
	;
	return v54, nil
}
func build_vector(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__5741 vm.Value
	var v12 vm.Value
	var form vm.Value
	var ctx vm.Value
	var arg__5755 vm.Value
	var v23 vm.Value
	var arg__5775 vm.Value
	var arg__5781 vm.Value
	var fn_id vm.Value
	var arg__5793 vm.Value
	var v48 vm.Value
	var v50 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__5741, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{arg__5741})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		form = arg0
		ctx = arg1
		goto b1
	} else {
		form = arg0
		ctx = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5755, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	v23, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5755, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), form})
	if callErr != nil {
		return nil, callErr
	}
	v50 = v23
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "resolve").Deref(), []vm.Value{vm.Symbol("vector")})
	if callErr != nil {
		return nil, callErr
	}
	arg__5775, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__5781, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "resolve").Deref(), []vm.Value{vm.Symbol("vector")})
	if callErr != nil {
		return nil, callErr
	}
	fn_id, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__5775, vm.Keyword("load-var"), vm.NewArrayVector([]vm.Value{}), arg__5781})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__5793, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v48, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-call-with-head").Deref(), []vm.Value{fn_id, arg__5793, ctx})
	if callErr != nil {
		return nil, callErr
	}
	v50 = v48
	goto b3
b3:
	;
	return v50, nil
}
func captures_of(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v8 vm.Value
	var form vm.Value
	var known_locals vm.Value
	var v14 vm.Value
	var or__x_31 vm.Value
	var v600 vm.Value
	var v17 vm.Value
	var v20 vm.Value
	var v22 vm.Value
	var head vm.Value
	var v57 bool
	var v552 vm.Value
	var v596 vm.Value
	var or__x_34 vm.Value
	var v41 vm.Value
	var v43 vm.Value
	var v60 vm.Value
	var v69 bool
	var v542 vm.Value
	var v72 vm.Value
	var v81 bool
	var v537 vm.Value
	var val vm.Value
	var v102 vm.Value
	var v111 bool
	var v532 vm.Value
	var maybe_name vm.Value
	var raw_rest vm.Value
	var has_name_QMARK_ vm.Value
	var or__x_345 bool
	var v527 vm.Value
	var name_sym vm.Value
	var v181 vm.Value
	var rest_forms vm.Value
	var and__x vm.Value
	var arg__5894 vm.Value
	var v223 vm.Value
	var multi_QMARK_ vm.Value
	var arg__5995 vm.Value
	var arg__6065 vm.Value
	var v291 vm.Value
	var args_vec vm.Value
	var body vm.Value
	var arg__6082 vm.Value
	var inner_known vm.Value
	var arg__6103 vm.Value
	var arg__6119 vm.Value
	var v322 vm.Value
	var v324 vm.Value
	var bindings vm.Value
	var pairs vm.Value
	var arg__6217 vm.Value
	var arg__6219 vm.Value
	var arg__6288 vm.Value
	var arg__6290 vm.Value
	var arg__6291 vm.Value
	var vec__5798 vm.Value
	var used vm.Value
	var let_bound_set vm.Value
	var new_locals vm.Value
	var arg__6332 vm.Value
	var arg__6348 vm.Value
	var body_captures vm.Value
	var arg__6363 vm.Value
	var v482 vm.Value
	var v522 vm.Value
	var or__x_349 bool
	var or__x_357 bool
	var v395 bool
	var or__x_361 bool
	var or__x_369 bool
	var v389 bool
	var or__x_373 bool
	var v381 bool
	var v383 bool
	var arg__6384 vm.Value
	var arg__6400 vm.Value
	var v513 vm.Value
	var v517 vm.Value
	var arg__6424 vm.Value
	var arg__6440 vm.Value
	var v575 vm.Value
	var v592 vm.Value
	var v584 vm.Value
	var v588 vm.Value
	var callErr error
	v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v8) {
		form = arg0
		known_locals = arg1
		goto b1
	} else {
		form = arg0
		known_locals = arg1
		goto b2
	}
b1:
	;
	v14, callErr = rt.InvokeValue(known_locals, []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v14) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	or__x_31, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x_31) {
		or__x_34 = or__x_31
		goto b10
	} else {
		goto b11
	}
b3:
	;
	return v600, nil
b4:
	;
	v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v22 = v17
	goto b6
b5:
	;
	v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v22 = v20
	goto b6
b6:
	;
	v600 = v22
	goto b3
b7:
	;
	head, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v57 = head == vm.Symbol("quote")
	if v57 {
		goto b13
	} else {
		goto b14
	}
b8:
	;
	v552, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v552) {
		goto b52
	} else {
		goto b53
	}
b9:
	;
	v600 = v596
	goto b3
b10:
	;
	v43 = or__x_34
	goto b12
b11:
	;
	v41, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v43 = v41
	goto b12
b12:
	;
	if vm.IsTruthy(v43) {
		goto b7
	} else {
		goto b8
	}
b13:
	;
	v60, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v542 = v60
	goto b15
b14:
	;
	v69 = head == vm.Symbol("var")
	if v69 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v596 = v542
	goto b9
b16:
	;
	v72, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v537 = v72
	goto b18
b17:
	;
	v81 = head == vm.Symbol("set!")
	if v81 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v542 = v537
	goto b15
b19:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	val, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(2), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v102, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{val, known_locals})
	if callErr != nil {
		return nil, callErr
	}
	v532 = v102
	goto b21
b20:
	;
	v111 = head == vm.Symbol("fn*")
	if v111 {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v537 = v532
	goto b18
b22:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	maybe_name, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	raw_rest, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), form})
	if callErr != nil {
		return nil, callErr
	}
	has_name_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{maybe_name})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(has_name_QMARK_) {
		goto b25
	} else {
		goto b26
	}
b23:
	;
	or__x_345 = head == vm.Symbol("let")
	if or__x_345 {
		or__x_349 = or__x_345
		goto b40
	} else {
		goto b41
	}
b24:
	;
	v532 = v527
	goto b21
b25:
	;
	name_sym = maybe_name
	goto b27
b26:
	;
	name_sym = vm.NIL
	goto b27
b27:
	;
	if vm.IsTruthy(has_name_QMARK_) {
		goto b28
	} else {
		goto b29
	}
b28:
	;
	rest_forms = raw_rest
	goto b30
b29:
	;
	v181, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{maybe_name, raw_rest})
	if callErr != nil {
		return nil, callErr
	}
	rest_forms = v181
	goto b30
b30:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b31
	} else {
		goto b32
	}
b31:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	arg__5894, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	v223, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{arg__5894})
	if callErr != nil {
		return nil, callErr
	}
	multi_QMARK_ = v223
	goto b33
b32:
	;
	multi_QMARK_ = and__x
	goto b33
b33:
	;
	if vm.IsTruthy(multi_QMARK_) {
		goto b34
	} else {
		goto b35
	}
b34:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var body vm.Value
		var arg__5974 vm.Value
		var inner_known vm.Value
		var v18 vm.Value
		var callErr error
		args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var args_vec vm.Value
			var name_sym_3 vm.Value
			var arg__5900 vm.Value
			var arg__5910 vm.Value
			var arg__5912 vm.Value
			var v20 vm.Value
			var arg__5919 vm.Value
			var v27 vm.Value
			var v29 vm.Value
			var callErr error
			if vm.IsTruthy(name_sym) {
				args_vec = arg0
				name_sym_3 = name_sym
				goto b1
			} else {
				args_vec = arg0
				goto b2
			}
		b1:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5900, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5900, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5910, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5910, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__5912, name_sym_3})
			if callErr != nil {
				return nil, callErr
			}
			v29 = v20
			goto b3
		b2:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5919, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5919, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			v29 = v27
			goto b3
		b3:
			;
			return v29, nil
		}), []vm.Value{args_vec})
		if callErr != nil {
			return nil, callErr
		}
		arg__5974, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var args_vec vm.Value
			var name_sym_3 vm.Value
			var arg__5900 vm.Value
			var arg__5910 vm.Value
			var arg__5912 vm.Value
			var v20 vm.Value
			var arg__5919 vm.Value
			var v27 vm.Value
			var v29 vm.Value
			var callErr error
			if vm.IsTruthy(name_sym) {
				args_vec = arg0
				name_sym_3 = name_sym
				goto b1
			} else {
				args_vec = arg0
				goto b2
			}
		b1:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5900, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5900, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5910, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5910, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__5912, name_sym_3})
			if callErr != nil {
				return nil, callErr
			}
			v29 = v20
			goto b3
		b2:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5919, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5919, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			v29 = v27
			goto b3
		b3:
			;
			return v29, nil
		}), []vm.Value{args_vec})
		if callErr != nil {
			return nil, callErr
		}
		inner_known, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "difference").Deref(), []vm.Value{known_locals, arg__5974})
		if callErr != nil {
			return nil, callErr
		}
		v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var v3 vm.Value
			var callErr error
			v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, inner_known})
			if callErr != nil {
				return nil, callErr
			}
			return v3, nil
		}), body})
		if callErr != nil {
			return nil, callErr
		}
		return v18, nil
	}), rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	arg__5995, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__6065, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var body vm.Value
		var arg__6047 vm.Value
		var inner_known vm.Value
		var v18 vm.Value
		var callErr error
		args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var args_vec vm.Value
			var name_sym_3 vm.Value
			var arg__5900 vm.Value
			var arg__5910 vm.Value
			var arg__5912 vm.Value
			var v20 vm.Value
			var arg__5919 vm.Value
			var v27 vm.Value
			var v29 vm.Value
			var callErr error
			if vm.IsTruthy(name_sym) {
				args_vec = arg0
				name_sym_3 = name_sym
				goto b1
			} else {
				args_vec = arg0
				goto b2
			}
		b1:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5900, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5900, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5910, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5910, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__5912, name_sym_3})
			if callErr != nil {
				return nil, callErr
			}
			v29 = v20
			goto b3
		b2:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5919, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5919, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			v29 = v27
			goto b3
		b3:
			;
			return v29, nil
		}), []vm.Value{args_vec})
		if callErr != nil {
			return nil, callErr
		}
		arg__6047, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var args_vec vm.Value
			var name_sym_3 vm.Value
			var arg__5900 vm.Value
			var arg__5910 vm.Value
			var arg__5912 vm.Value
			var v20 vm.Value
			var arg__5919 vm.Value
			var v27 vm.Value
			var v29 vm.Value
			var callErr error
			if vm.IsTruthy(name_sym) {
				args_vec = arg0
				name_sym_3 = name_sym
				goto b1
			} else {
				args_vec = arg0
				goto b2
			}
		b1:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5900, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5900, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5910, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5910, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__5912, name_sym_3})
			if callErr != nil {
				return nil, callErr
			}
			v29 = v20
			goto b3
		b2:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			arg__5919, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
			if callErr != nil {
				return nil, callErr
			}
			v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5919, args_vec})
			if callErr != nil {
				return nil, callErr
			}
			v29 = v27
			goto b3
		b3:
			;
			return v29, nil
		}), []vm.Value{args_vec})
		if callErr != nil {
			return nil, callErr
		}
		inner_known, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "difference").Deref(), []vm.Value{known_locals, arg__6047})
		if callErr != nil {
			return nil, callErr
		}
		v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var v3 vm.Value
			var callErr error
			v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, inner_known})
			if callErr != nil {
				return nil, callErr
			}
			return v3, nil
		}), body})
		if callErr != nil {
			return nil, callErr
		}
		return v18, nil
	}), rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	v291, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5995, arg__6065})
	if callErr != nil {
		return nil, callErr
	}
	v324 = v291
	goto b36
b35:
	;
	args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var name_sym_3 vm.Value
		var arg__5900 vm.Value
		var arg__5910 vm.Value
		var arg__5912 vm.Value
		var v20 vm.Value
		var arg__5919 vm.Value
		var v27 vm.Value
		var v29 vm.Value
		var callErr error
		if vm.IsTruthy(name_sym) {
			args_vec = arg0
			name_sym_3 = name_sym
			goto b1
		} else {
			args_vec = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__5900, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5900, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__5910, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__5912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5910, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__5912, name_sym_3})
		if callErr != nil {
			return nil, callErr
		}
		v29 = v20
		goto b3
	b2:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__5919, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5919, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		v29 = v27
		goto b3
	b3:
		;
		return v29, nil
	}), []vm.Value{args_vec})
	if callErr != nil {
		return nil, callErr
	}
	arg__6082, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var name_sym_3 vm.Value
		var arg__5900 vm.Value
		var arg__5910 vm.Value
		var arg__5912 vm.Value
		var v20 vm.Value
		var arg__5919 vm.Value
		var v27 vm.Value
		var v29 vm.Value
		var callErr error
		if vm.IsTruthy(name_sym) {
			args_vec = arg0
			name_sym_3 = name_sym
			goto b1
		} else {
			args_vec = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__5900, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5900, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__5910, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__5912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5910, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__5912, name_sym_3})
		if callErr != nil {
			return nil, callErr
		}
		v29 = v20
		goto b3
	b2:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__5919, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__5919, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		v29 = v27
		goto b3
	b3:
		;
		return v29, nil
	}), []vm.Value{args_vec})
	if callErr != nil {
		return nil, callErr
	}
	inner_known, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "difference").Deref(), []vm.Value{known_locals, arg__6082})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, inner_known})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body})
	if callErr != nil {
		return nil, callErr
	}
	arg__6103, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__6119, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, inner_known})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body})
	if callErr != nil {
		return nil, callErr
	}
	v322, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__6103, arg__6119})
	if callErr != nil {
		return nil, callErr
	}
	v324 = v322
	goto b36
b36:
	;
	v527 = v324
	goto b24
b37:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	bindings, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), form})
	if callErr != nil {
		return nil, callErr
	}
	pairs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "partition").Deref(), []vm.Value{vm.Int(2), bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__6217, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__6219, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__6217, arg__6219})
	if callErr != nil {
		return nil, callErr
	}
	arg__6288, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__6290, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__6291, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__6288, arg__6290})
	if callErr != nil {
		return nil, callErr
	}
	vec__5798, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var caps vm.Value
		var shadowed vm.Value
		var sym vm.Value
		var init vm.Value
		var locs vm.Value
		var arg__6271 vm.Value
		var arg__6272 vm.Value
		var arg__6283 vm.Value
		var arg__6284 vm.Value
		var v42 vm.Value
		var callErr error
		caps, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		shadowed, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		init, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		locs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "difference").Deref(), []vm.Value{known_locals, shadowed})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{init, locs})
		if callErr != nil {
			return nil, callErr
		}
		arg__6271, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{init, locs})
		if callErr != nil {
			return nil, callErr
		}
		arg__6272, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{caps, arg__6271})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "binding-syms").Deref(), []vm.Value{sym})
		if callErr != nil {
			return nil, callErr
		}
		arg__6283, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "binding-syms").Deref(), []vm.Value{sym})
		if callErr != nil {
			return nil, callErr
		}
		arg__6284, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{shadowed, arg__6283})
		if callErr != nil {
			return nil, callErr
		}
		v42, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__6272, arg__6284})
		if callErr != nil {
			return nil, callErr
		}
		return v42, nil
	}), arg__6291, pairs})
	if callErr != nil {
		return nil, callErr
	}
	used, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__5798, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	let_bound_set, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__5798, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	new_locals, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{known_locals, let_bound_set})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, new_locals})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body})
	if callErr != nil {
		return nil, callErr
	}
	arg__6332, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__6348, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, new_locals})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body})
	if callErr != nil {
		return nil, callErr
	}
	body_captures, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__6332, arg__6348})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "difference").Deref(), []vm.Value{body_captures, let_bound_set})
	if callErr != nil {
		return nil, callErr
	}
	arg__6363, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "difference").Deref(), []vm.Value{body_captures, let_bound_set})
	if callErr != nil {
		return nil, callErr
	}
	v482, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{used, arg__6363})
	if callErr != nil {
		return nil, callErr
	}
	v522 = v482
	goto b39
b38:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b49
	} else {
		goto b50
	}
b39:
	;
	v527 = v522
	goto b24
b40:
	;
	v395 = or__x_349
	goto b42
b41:
	;
	or__x_357 = head == vm.Symbol("let*")
	if or__x_357 {
		or__x_361 = or__x_357
		goto b43
	} else {
		goto b44
	}
b42:
	;
	if v395 {
		goto b37
	} else {
		goto b38
	}
b43:
	;
	v389 = or__x_361
	goto b45
b44:
	;
	or__x_369 = head == vm.Symbol("loop")
	if or__x_369 {
		or__x_373 = or__x_369
		goto b46
	} else {
		goto b47
	}
b45:
	;
	v395 = v389
	goto b42
b46:
	;
	v383 = or__x_373
	goto b48
b47:
	;
	v381 = head == vm.Symbol("loop*")
	v383 = v381
	goto b48
b48:
	;
	v389 = v383
	goto b45
b49:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, known_locals})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__6384, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__6400, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, known_locals})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	v513, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__6384, arg__6400})
	if callErr != nil {
		return nil, callErr
	}
	v517 = v513
	goto b51
b50:
	;
	v517 = vm.NIL
	goto b51
b51:
	;
	v522 = v517
	goto b39
b52:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, known_locals})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__6424, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__6440, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "captures-of").Deref(), []vm.Value{arg0, known_locals})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	v575, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__6424, arg__6440})
	if callErr != nil {
		return nil, callErr
	}
	v592 = v575
	goto b54
b53:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b55
	} else {
		goto b56
	}
b54:
	;
	v596 = v592
	goto b9
b55:
	;
	v584, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v588 = v584
	goto b57
b56:
	;
	v588 = vm.NIL
	goto b57
b57:
	;
	v592 = v588
	goto b54
}
func ctx_block(arg0 vm.Value) (vm.Value, error) {
	var arg__6451 vm.Value
	var v8 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__6451, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__6451, vm.Keyword("current-block")})
	if callErr != nil {
		return nil, callErr
	}
	return v8, nil
}
func ctx_fn(arg0 vm.Value) (vm.Value, error) {
	var arg__6462 vm.Value
	var v8 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__6462, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__6462, vm.Keyword("fn")})
	if callErr != nil {
		return nil, callErr
	}
	return v8, nil
}
func ctx_set_block_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__6475 vm.Value
	var arg__6491 vm.Value
	var arg__6494 vm.Value
	var v19 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__6475, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg__6475, vm.Keyword("current-block"), arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__6491, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__6494, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg__6491, vm.Keyword("current-block"), arg1})
	if callErr != nil {
		return nil, callErr
	}
	v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{arg0, arg__6494})
	if callErr != nil {
		return nil, callErr
	}
	return v19, nil
}
func current_locals_flat(arg0 vm.Value) (vm.Value, error) {
	var arg__6506 vm.Value
	var arg__6521 vm.Value
	var arg__6523 vm.Value
	var v23 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__6506, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__6506, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__6521, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__6523, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__6521, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	v23, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.LookupVar("clojure.core", "merge").Deref(), vm.EmptyPersistentMap, arg__6523})
	if callErr != nil {
		return nil, callErr
	}
	return v23, nil
}
func emit_template_closure(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__6539 vm.Value
	var arg__6543 vm.Value
	var const_id vm.Value
	var v26 vm.Value
	var capture_syms vm.Value
	var ctx vm.Value
	var arg__6567 vm.Value
	var arg__6571 vm.Value
	var arg__6575 vm.Value
	var v45 vm.Value
	var closure_id vm.Value
	var cls vm.Value
	var caps vm.Value
	var v72 vm.Value
	var cap_sym vm.Value
	var cap_val vm.Value
	var arg__6606 vm.Value
	var arg__6610 vm.Value
	var arg__6615 vm.Value
	var push_id vm.Value
	var v97 vm.Value
	var v100 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__6539, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__6543, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	const_id, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{arg__6539, arg__6543, vm.Keyword("const"), vm.NewArrayVector([]vm.Value{}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v26) {
		capture_syms = arg1
		ctx = arg2
		goto b1
	} else {
		capture_syms = arg1
		ctx = arg2
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{const_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__6567, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__6571, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__6575, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{const_id})
	if callErr != nil {
		return nil, callErr
	}
	v45, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{arg__6567, arg__6571, vm.Keyword("make-closure"), arg__6575, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	closure_id = v45
	goto b3
b2:
	;
	closure_id = const_id
	goto b3
b3:
	;
	cls = closure_id
	caps = capture_syms
	goto b4
b4:
	;
	v72, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{caps})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v72) {
		goto b5
	} else {
		goto b6
	}
b5:
	;
	cap_sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{caps})
	if callErr != nil {
		return nil, callErr
	}
	cap_val, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-symbol").Deref(), []vm.Value{cap_sym, ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{cls, cap_val})
	if callErr != nil {
		return nil, callErr
	}
	arg__6606, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-fn").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__6610, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__6615, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{cls, cap_val})
	if callErr != nil {
		return nil, callErr
	}
	push_id, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{arg__6606, arg__6610, vm.Keyword("push-closed"), arg__6615, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v97, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{caps})
	if callErr != nil {
		return nil, callErr
	}
	cls = push_id
	caps = v97
	goto b4
b6:
	;
	v100 = cls
	goto b7
b7:
	;
	return v100, nil
}
func expand_binding(arg0 vm.Value) (vm.Value, error) {
	var b vm.Value
	var out vm.Value
	var v12 vm.Value
	var n vm.Value
	var v vm.Value
	var v30 vm.Value
	var v150 vm.Value
	var gs vm.Value
	var v39 vm.Value
	var arg__6662 vm.Value
	var arg__6668 vm.Value
	var v49 vm.Value
	var v62 vm.Value
	var v71 vm.Value
	var arg__6702 vm.Value
	var arg__6708 vm.Value
	var v81 vm.Value
	var v94 vm.Value
	var v99 vm.Value
	var v116 vm.Value
	var v118 vm.Value
	var callErr error
	b = arg0
	out = vm.NewArrayVector([]vm.Value{})
	goto b1
b1:
	;
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{b})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v150 = out
	goto b4
b3:
	;
	n, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{b})
	if callErr != nil {
		return nil, callErr
	}
	v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{b})
	if callErr != nil {
		return nil, callErr
	}
	v30, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{n})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v30) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v150, nil
b5:
	;
	gs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "gensym").Deref(), []vm.Value{vm.String("vec__")})
	if callErr != nil {
		return nil, callErr
	}
	v39, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), b})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, gs, v})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-vector-pattern").Deref(), []vm.Value{gs, n})
	if callErr != nil {
		return nil, callErr
	}
	arg__6662, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, gs, v})
	if callErr != nil {
		return nil, callErr
	}
	arg__6668, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-vector-pattern").Deref(), []vm.Value{gs, n})
	if callErr != nil {
		return nil, callErr
	}
	v49, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__6662, arg__6668})
	if callErr != nil {
		return nil, callErr
	}
	b = v39
	out = v49
	goto b1
b6:
	;
	v62, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{n})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v62) {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	v150 = vm.NIL
	goto b4
b8:
	;
	gs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "gensym").Deref(), []vm.Value{vm.String("map__")})
	if callErr != nil {
		return nil, callErr
	}
	v71, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), b})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, gs, v})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-map-pattern").Deref(), []vm.Value{gs, n})
	if callErr != nil {
		return nil, callErr
	}
	arg__6702, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, gs, v})
	if callErr != nil {
		return nil, callErr
	}
	arg__6708, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-map-pattern").Deref(), []vm.Value{gs, n})
	if callErr != nil {
		return nil, callErr
	}
	v81, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__6702, arg__6708})
	if callErr != nil {
		return nil, callErr
	}
	b = v71
	out = v81
	goto b1
b9:
	;
	v94, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{n})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v94) {
		goto b11
	} else {
		goto b12
	}
b10:
	;
	goto b7
b11:
	;
	v99, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), b})
	if callErr != nil {
		return nil, callErr
	}
	b = v99
	goto b1
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
	v116, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), b})
	if callErr != nil {
		return nil, callErr
	}
	v118, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, n, v})
	if callErr != nil {
		return nil, callErr
	}
	b = v116
	out = v118
	goto b1
b15:
	;
	goto b16
b16:
	;
	goto b13
}
func expand_fn_args(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var i int
	var remaining vm.Value
	var v16 vm.Value
	var args_vec vm.Value
	var body_forms vm.Value
	var arg__6735 vm.Value
	var v31 bool
	var amp_pos int
	var variadic_QMARK_ vm.Value
	var v34 int
	var v36 vm.Value
	var v38 int
	var v67 vm.Value
	var fixed_args vm.Value
	var arg__6753 int
	var v97 vm.Value
	var rest_sym vm.Value
	var has_destructure_QMARK_ vm.Value
	var v560 vm.Value
	var or__x vm.Value
	var v154 vm.Value
	var flat_args vm.Value
	var let_binds vm.Value
	var v196 vm.Value
	var v201 vm.Value
	var x vm.Value
	var v230 vm.Value
	var result vm.Value
	var v235 vm.Value
	var gs vm.Value
	var v268 vm.Value
	var v271 vm.Value
	var v300 vm.Value
	var binds vm.Value
	var v401 vm.Value
	var v403 vm.Value
	var v303 vm.Value
	var v369 vm.Value
	var arg__6834 vm.Value
	var arg__6847 vm.Value
	var arg__6848 vm.Value
	var v350 vm.Value
	var v354 vm.Value
	var arg__6869 vm.Value
	var v447 vm.Value
	var v450 vm.Value
	var final_flat_args vm.Value
	var arg__6879 vm.Value
	var v496 vm.Value
	var arg__6888 vm.Value
	var arg__6900 vm.Value
	var arg__6901 vm.Value
	var arg__6913 vm.Value
	var arg__6925 vm.Value
	var arg__6926 vm.Value
	var arg__6928 vm.Value
	var v535 vm.Value
	var body vm.Value
	var v556 vm.Value
	var callErr error
	i = 0
	remaining = arg0
	goto b1
b1:
	;
	v16, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v16) {
		args_vec = arg0
		body_forms = arg1
		goto b2
	} else {
		args_vec = arg0
		body_forms = arg1
		goto b3
	}
b2:
	;
	amp_pos = -1
	goto b4
b3:
	;
	arg__6735, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v31 = arg__6735 == vm.Symbol("&")
	if v31 {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	variadic_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Int(amp_pos), vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(variadic_QMARK_) {
		goto b8
	} else {
		goto b9
	}
b5:
	;
	v38 = i
	goto b7
b6:
	;
	v34 = i + 1
	v36, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	i = v34
	remaining = v36
	goto b1
b7:
	;
	amp_pos = v38
	goto b4
b8:
	;
	v67, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "take").Deref(), []vm.Value{vm.Int(amp_pos), args_vec})
	if callErr != nil {
		return nil, callErr
	}
	fixed_args = v67
	goto b10
b9:
	;
	fixed_args = args_vec
	goto b10
b10:
	;
	if vm.IsTruthy(variadic_QMARK_) {
		goto b11
	} else {
		goto b12
	}
b11:
	;
	arg__6753 = amp_pos + 1
	v97, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args_vec, vm.Int(arg__6753), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	rest_sym = v97
	goto b13
b12:
	;
	rest_sym = vm.NIL
	goto b13
b13:
	;
	has_destructure_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__6780 vm.Value
		var v6 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__6780, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__6780})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), fixed_args})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(variadic_QMARK_) {
		or__x = variadic_QMARK_
		goto b17
	} else {
		goto b18
	}
b14:
	;
	flat_args = vm.NewArrayVector([]vm.Value{})
	let_binds = vm.NewArrayVector([]vm.Value{})
	goto b20
b15:
	;
	v560 = vm.NIL
	goto b16
b16:
	;
	return v560, nil
b17:
	;
	v154 = or__x
	goto b19
b18:
	;
	v154 = has_destructure_QMARK_
	goto b19
b19:
	;
	if vm.IsTruthy(v154) {
		goto b14
	} else {
		goto b15
	}
b20:
	;
	v196, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v196) {
		goto b21
	} else {
		goto b22
	}
b21:
	;
	v201, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("flat-args"), flat_args, vm.Keyword("let-binds"), let_binds})
	if callErr != nil {
		return nil, callErr
	}
	result = v201
	goto b23
b22:
	;
	x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v230, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v230) {
		goto b24
	} else {
		goto b25
	}
b23:
	;
	if vm.IsTruthy(variadic_QMARK_) {
		goto b36
	} else {
		goto b37
	}
b24:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v235, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{flat_args, x})
	if callErr != nil {
		return nil, callErr
	}
	flat_args = v235
	goto b20
b25:
	;
	gs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "gensym").Deref(), []vm.Value{vm.String("p__")})
	if callErr != nil {
		return nil, callErr
	}
	v268, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v268) {
		goto b27
	} else {
		goto b28
	}
b27:
	;
	v271, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-vector-pattern").Deref(), []vm.Value{gs, x})
	if callErr != nil {
		return nil, callErr
	}
	binds = v271
	goto b29
b28:
	;
	v300, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v300) {
		goto b30
	} else {
		goto b31
	}
b29:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v401, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{flat_args, gs})
	if callErr != nil {
		return nil, callErr
	}
	v403, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{let_binds, binds})
	if callErr != nil {
		return nil, callErr
	}
	flat_args = v401
	let_binds = v403
	goto b20
b30:
	;
	v303, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-map-pattern").Deref(), []vm.Value{gs, x})
	if callErr != nil {
		return nil, callErr
	}
	v369 = v303
	goto b32
b31:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b33
	} else {
		goto b34
	}
b32:
	;
	binds = v369
	goto b29
b33:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pr-str").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	arg__6834, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pr-str").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("unsupported arg pattern: "), arg__6834})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pr-str").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	arg__6847, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pr-str").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	arg__6848, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("unsupported arg pattern: "), arg__6847})
	if callErr != nil {
		return nil, callErr
	}
	v350, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__6848})
	if callErr != nil {
		return nil, callErr
	}
	v354 = v350
	goto b35
b34:
	;
	v354 = vm.NIL
	goto b35
b35:
	;
	v369 = v354
	goto b32
b36:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("flat-args"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	arg__6869, callErr = rt.InvokeValue(vm.Keyword("flat-args"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	v447, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__6869, rest_sym})
	if callErr != nil {
		return nil, callErr
	}
	final_flat_args = v447
	goto b38
b37:
	;
	v450, callErr = rt.InvokeValue(vm.Keyword("flat-args"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	final_flat_args = v450
	goto b38
b38:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	arg__6879, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	v496, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__6879})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v496) {
		goto b39
	} else {
		goto b40
	}
b39:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	arg__6888, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__6888})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	arg__6900, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	arg__6901, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__6900})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), vm.Symbol("let*"), arg__6901, body_forms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	arg__6913, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__6913})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	arg__6925, callErr = rt.InvokeValue(vm.Keyword("let-binds"), []vm.Value{result})
	if callErr != nil {
		return nil, callErr
	}
	arg__6926, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__6925})
	if callErr != nil {
		return nil, callErr
	}
	arg__6928, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), vm.Symbol("let*"), arg__6926, body_forms})
	if callErr != nil {
		return nil, callErr
	}
	v535, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__6928})
	if callErr != nil {
		return nil, callErr
	}
	body = v535
	goto b41
b40:
	;
	body = body_forms
	goto b41
b41:
	;
	v556, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("variadic?"), variadic_QMARK_, vm.Keyword("flat-args"), final_flat_args, vm.Keyword("body"), body})
	if callErr != nil {
		return nil, callErr
	}
	v560 = v556
	goto b16
}
func expand_map_pattern(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var defaults vm.Value
	var as_sym vm.Value
	var keys_STAR_ vm.Value
	var strs_STAR_ vm.Value
	var ks vm.Value
	var out vm.Value
	var v67 vm.Value
	var sym vm.Value
	var or__x vm.Value
	var v37 vm.Value
	var k vm.Value
	var v93 vm.Value
	var binds vm.Value
	var arg__6989 vm.Value
	var v100 vm.Value
	var kn vm.Value
	var v115 vm.Value
	var arg__7009 vm.Value
	var arg__7028 vm.Value
	var arg__7029 vm.Value
	var v127 vm.Value
	var v169 vm.Value
	var v205 vm.Value
	var v220 vm.Value
	var callErr error
	defaults, callErr = rt.InvokeValue(vm.Keyword("or"), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	as_sym, callErr = rt.InvokeValue(vm.Keyword("as"), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	keys_STAR_, callErr = rt.InvokeValue(vm.Keyword("keys"), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	strs_STAR_, callErr = rt.InvokeValue(vm.Keyword("strs"), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(keys_STAR_) {
		sym = arg0
		or__x = keys_STAR_
		goto b2
	} else {
		sym = arg0
		goto b3
	}
b1:
	;
	v67, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{ks})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v67) {
		goto b5
	} else {
		goto b6
	}
b2:
	;
	v37 = or__x
	goto b4
b3:
	;
	v37 = vm.NewArrayVector([]vm.Value{})
	goto b4
b4:
	;
	ks = v37
	out = vm.NewArrayVector([]vm.Value{})
	goto b1
b5:
	;
	binds = out
	goto b7
b6:
	;
	k, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{ks})
	if callErr != nil {
		return nil, callErr
	}
	v93, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{k})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v93) {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	if vm.IsTruthy(strs_STAR_) {
		goto b11
	} else {
		goto b12
	}
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{k})
	if callErr != nil {
		return nil, callErr
	}
	arg__6989, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{k})
	if callErr != nil {
		return nil, callErr
	}
	v100, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol").Deref(), []vm.Value{arg__6989})
	if callErr != nil {
		return nil, callErr
	}
	kn = v100
	goto b10
b9:
	;
	kn = k
	goto b10
b10:
	;
	v115, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{ks})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{defaults, kn})
	if callErr != nil {
		return nil, callErr
	}
	arg__7009, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{defaults, kn})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var kn vm.Value
		var default_ vm.Value
		var sym_5 vm.Value
		var arg__6957 vm.Value
		var v17 vm.Value
		var sym_8 vm.Value
		var arg__6971 vm.Value
		var v26 vm.Value
		var v28 vm.Value
		var callErr error
		if vm.IsTruthy(arg1) {
			kn = arg0
			default_ = arg1
			sym_5 = arg0
			goto b1
		} else {
			kn = arg0
			sym_8 = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
		if callErr != nil {
			return nil, callErr
		}
		arg__6957, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
		if callErr != nil {
			return nil, callErr
		}
		v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("get"), sym_5, arg__6957, default_})
		if callErr != nil {
			return nil, callErr
		}
		v28 = v17
		goto b3
	b2:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
		if callErr != nil {
			return nil, callErr
		}
		arg__6971, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
		if callErr != nil {
			return nil, callErr
		}
		v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("get"), sym_8, arg__6971})
		if callErr != nil {
			return nil, callErr
		}
		v28 = v26
		goto b3
	b3:
		;
		return v28, nil
	}), []vm.Value{kn, arg__7009})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{defaults, kn})
	if callErr != nil {
		return nil, callErr
	}
	arg__7028, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{defaults, kn})
	if callErr != nil {
		return nil, callErr
	}
	arg__7029, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var kn vm.Value
		var default_ vm.Value
		var sym_5 vm.Value
		var arg__6957 vm.Value
		var v17 vm.Value
		var sym_8 vm.Value
		var arg__6971 vm.Value
		var v26 vm.Value
		var v28 vm.Value
		var callErr error
		if vm.IsTruthy(arg1) {
			kn = arg0
			default_ = arg1
			sym_5 = arg0
			goto b1
		} else {
			kn = arg0
			sym_8 = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
		if callErr != nil {
			return nil, callErr
		}
		arg__6957, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
		if callErr != nil {
			return nil, callErr
		}
		v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("get"), sym_5, arg__6957, default_})
		if callErr != nil {
			return nil, callErr
		}
		v28 = v17
		goto b3
	b2:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
		if callErr != nil {
			return nil, callErr
		}
		arg__6971, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
		if callErr != nil {
			return nil, callErr
		}
		v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("get"), sym_8, arg__6971})
		if callErr != nil {
			return nil, callErr
		}
		v28 = v26
		goto b3
	b3:
		;
		return v28, nil
	}), []vm.Value{kn, arg__7028})
	if callErr != nil {
		return nil, callErr
	}
	v127, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, kn, arg__7029})
	if callErr != nil {
		return nil, callErr
	}
	ks = v115
	out = v127
	goto b1
b11:
	;
	v169, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var v13 vm.Value
		var out vm.Value
		var s vm.Value
		var defaults_6 vm.Value
		var get_expr_7 vm.Value
		var arg__7094 vm.Value
		var v20 vm.Value
		var defaults_10 vm.Value
		var get_expr_11 vm.Value
		var kn vm.Value
		var defaults_26 vm.Value
		var get_expr_27 vm.Value
		var arg__7111 vm.Value
		var arg__7130 vm.Value
		var arg__7131 vm.Value
		var v39 vm.Value
		var callErr error
		v13, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{arg1})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v13) {
			out = arg0
			s = arg1
			defaults_6 = defaults
			get_expr_7 = rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
				var kn vm.Value
				var default_ vm.Value
				var sym_5 vm.Value
				var arg__6957 vm.Value
				var v17 vm.Value
				var sym_8 vm.Value
				var arg__6971 vm.Value
				var v26 vm.Value
				var v28 vm.Value
				var callErr error
				if vm.IsTruthy(arg1) {
					kn = arg0
					default_ = arg1
					sym_5 = arg0
					goto b1
				} else {
					kn = arg0
					sym_8 = arg0
					goto b2
				}
			b1:
				;
				_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
				if callErr != nil {
					return nil, callErr
				}
				arg__6957, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
				if callErr != nil {
					return nil, callErr
				}
				v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("get"), sym_5, arg__6957, default_})
				if callErr != nil {
					return nil, callErr
				}
				v28 = v17
				goto b3
			b2:
				;
				_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
				if callErr != nil {
					return nil, callErr
				}
				arg__6971, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
				if callErr != nil {
					return nil, callErr
				}
				v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("get"), sym_8, arg__6971})
				if callErr != nil {
					return nil, callErr
				}
				v28 = v26
				goto b3
			b3:
				;
				return v28, nil
			})
			goto b1
		} else {
			out = arg0
			s = arg1
			defaults_10 = defaults
			get_expr_11 = rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
				var kn vm.Value
				var default_ vm.Value
				var sym_5 vm.Value
				var arg__6957 vm.Value
				var v17 vm.Value
				var sym_8 vm.Value
				var arg__6971 vm.Value
				var v26 vm.Value
				var v28 vm.Value
				var callErr error
				if vm.IsTruthy(arg1) {
					kn = arg0
					default_ = arg1
					sym_5 = arg0
					goto b1
				} else {
					kn = arg0
					sym_8 = arg0
					goto b2
				}
			b1:
				;
				_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
				if callErr != nil {
					return nil, callErr
				}
				arg__6957, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
				if callErr != nil {
					return nil, callErr
				}
				v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("get"), sym_5, arg__6957, default_})
				if callErr != nil {
					return nil, callErr
				}
				v28 = v17
				goto b3
			b2:
				;
				_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
				if callErr != nil {
					return nil, callErr
				}
				arg__6971, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword").Deref(), []vm.Value{kn})
				if callErr != nil {
					return nil, callErr
				}
				v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("get"), sym_8, arg__6971})
				if callErr != nil {
					return nil, callErr
				}
				v28 = v26
				goto b3
			b3:
				;
				return v28, nil
			})
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		arg__7094, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "name").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol").Deref(), []vm.Value{arg__7094})
		if callErr != nil {
			return nil, callErr
		}
		kn = v20
		defaults_26 = defaults_6
		get_expr_27 = get_expr_7
		goto b3
	b2:
		;
		kn = s
		defaults_26 = defaults_10
		get_expr_27 = get_expr_11
		goto b3
	b3:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{defaults_26, kn})
		if callErr != nil {
			return nil, callErr
		}
		arg__7111, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{defaults_26, kn})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(get_expr_27, []vm.Value{kn, arg__7111})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{defaults_26, kn})
		if callErr != nil {
			return nil, callErr
		}
		arg__7130, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{defaults_26, kn})
		if callErr != nil {
			return nil, callErr
		}
		arg__7131, callErr = rt.InvokeValue(get_expr_27, []vm.Value{kn, arg__7130})
		if callErr != nil {
			return nil, callErr
		}
		v39, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, kn, arg__7131})
		if callErr != nil {
			return nil, callErr
		}
		return v39, nil
	}), binds, strs_STAR_})
	if callErr != nil {
		return nil, callErr
	}
	binds = v169
	goto b13
b12:
	;
	goto b13
b13:
	;
	if vm.IsTruthy(as_sym) {
		goto b14
	} else {
		goto b15
	}
b14:
	;
	v205, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{binds, as_sym, sym})
	if callErr != nil {
		return nil, callErr
	}
	binds = v205
	goto b16
b15:
	;
	goto b16
b16:
	;
	v220, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-binding").Deref(), []vm.Value{binds})
	if callErr != nil {
		return nil, callErr
	}
	return v220, nil
}
func expand_vector_pattern(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var i int
	var remaining vm.Value
	var out vm.Value
	var sym vm.Value
	var v20 vm.Value
	var v23 vm.Value
	var x vm.Value
	var v40 bool
	var v444 vm.Value
	var rest_pat vm.Value
	var rem_STAR_ vm.Value
	var rest_expr vm.Value
	var v71 vm.Value
	var v198 bool
	var v436 vm.Value
	var v74 vm.Value
	var gs vm.Value
	var arg__7200 vm.Value
	var arg__7210 vm.Value
	var arg__7219 vm.Value
	var arg__7220 vm.Value
	var v97 vm.Value
	var out_STAR_ vm.Value
	var and__x vm.Value
	var v176 vm.Value
	var arg__7245 vm.Value
	var v182 vm.Value
	var arg__7227 vm.Value
	var v157 bool
	var v160 vm.Value
	var v203 vm.Value
	var arg__7265 vm.Value
	var v209 vm.Value
	var v224 vm.Value
	var v428 vm.Value
	var v226 int
	var v228 vm.Value
	var arg__7298 vm.Value
	var v242 vm.Value
	var or__x vm.Value
	var v420 vm.Value
	var v305 vm.Value
	var v412 vm.Value
	var v275 vm.Value
	var v277 vm.Value
	var v308 vm.Value
	var v325 vm.Value
	var nested vm.Value
	var arg__7348 vm.Value
	var v363 int
	var v365 vm.Value
	var v367 vm.Value
	var v328 vm.Value
	var v332 vm.Value
	var arg__7368 vm.Value
	var arg__7381 vm.Value
	var arg__7382 vm.Value
	var v400 vm.Value
	var v404 vm.Value
	var callErr error
	i = 0
	remaining = arg1
	out = vm.NewArrayVector([]vm.Value{})
	sym = arg0
	goto b1
b1:
	;
	v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v20) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v23, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-binding").Deref(), []vm.Value{out})
	if callErr != nil {
		return nil, callErr
	}
	v444 = v23
	goto b4
b3:
	;
	x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v40 = x == vm.Symbol("&")
	if v40 {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v444, nil
b5:
	;
	rest_pat, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	rem_STAR_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), remaining})
	if callErr != nil {
		return nil, callErr
	}
	rest_expr, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("drop"), vm.Int(i), sym})
	if callErr != nil {
		return nil, callErr
	}
	v71, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{rest_pat})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v71) {
		goto b8
	} else {
		goto b9
	}
b6:
	;
	v198 = x == vm.Keyword("as")
	if v198 {
		goto b17
	} else {
		goto b18
	}
b7:
	;
	v444 = v436
	goto b4
b8:
	;
	v74, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, rest_pat, rest_expr})
	if callErr != nil {
		return nil, callErr
	}
	out_STAR_ = v74
	goto b10
b9:
	;
	gs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "gensym").Deref(), []vm.Value{vm.String("rest__")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, gs, rest_expr})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{gs, rest_pat})
	if callErr != nil {
		return nil, callErr
	}
	arg__7200, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{gs, rest_pat})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-binding").Deref(), []vm.Value{arg__7200})
	if callErr != nil {
		return nil, callErr
	}
	arg__7210, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, gs, rest_expr})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{gs, rest_pat})
	if callErr != nil {
		return nil, callErr
	}
	arg__7219, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{gs, rest_pat})
	if callErr != nil {
		return nil, callErr
	}
	arg__7220, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-binding").Deref(), []vm.Value{arg__7219})
	if callErr != nil {
		return nil, callErr
	}
	v97, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7210, arg__7220})
	if callErr != nil {
		return nil, callErr
	}
	out_STAR_ = v97
	goto b10
b10:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{rem_STAR_})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b14
	} else {
		goto b15
	}
b11:
	;
	v176, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), rem_STAR_})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{rem_STAR_})
	if callErr != nil {
		return nil, callErr
	}
	arg__7245, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{rem_STAR_})
	if callErr != nil {
		return nil, callErr
	}
	v182, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out_STAR_, arg__7245, sym})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v176
	out = v182
	goto b1
b12:
	;
	remaining = rem_STAR_
	out = out_STAR_
	goto b1
b14:
	;
	arg__7227, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rem_STAR_})
	if callErr != nil {
		return nil, callErr
	}
	v157 = arg__7227 == vm.Keyword("as")
	v160 = vm.Boolean(v157)
	goto b16
b15:
	;
	v160 = and__x
	goto b16
b16:
	;
	if vm.IsTruthy(v160) {
		goto b11
	} else {
		goto b12
	}
b17:
	;
	v203, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), remaining})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__7265, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v209, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, arg__7265, sym})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v203
	out = v209
	goto b1
b18:
	;
	v224, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v224) {
		goto b20
	} else {
		goto b21
	}
b19:
	;
	v436 = v428
	goto b7
b20:
	;
	v226 = i + 1
	v228, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("nth"), sym, vm.Int(i), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	arg__7298, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("nth"), sym, vm.Int(i), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v242, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, x, arg__7298})
	if callErr != nil {
		return nil, callErr
	}
	i = v226
	remaining = v228
	out = v242
	goto b1
b21:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b26
	} else {
		goto b27
	}
b22:
	;
	v428 = v420
	goto b19
b23:
	;
	gs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "gensym").Deref(), []vm.Value{vm.String("v__")})
	if callErr != nil {
		return nil, callErr
	}
	v305, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v305) {
		goto b29
	} else {
		goto b30
	}
b24:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b35
	} else {
		goto b36
	}
b25:
	;
	v420 = v412
	goto b22
b26:
	;
	v277 = or__x
	goto b28
b27:
	;
	v275, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	v277 = v275
	goto b28
b28:
	;
	if vm.IsTruthy(v277) {
		goto b23
	} else {
		goto b24
	}
b29:
	;
	v308, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-vector-pattern").Deref(), []vm.Value{gs, x})
	if callErr != nil {
		return nil, callErr
	}
	nested = v308
	goto b31
b30:
	;
	v325, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v325) {
		goto b32
	} else {
		goto b33
	}
b31:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("nth"), sym, vm.Int(i), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	arg__7348, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{vm.Symbol("nth"), sym, vm.Int(i), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	out_STAR_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, gs, arg__7348})
	if callErr != nil {
		return nil, callErr
	}
	v363 = i + 1
	v365, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v367, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{out_STAR_, nested})
	if callErr != nil {
		return nil, callErr
	}
	i = v363
	remaining = v365
	out = v367
	goto b1
b32:
	;
	v328, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "expand-map-pattern").Deref(), []vm.Value{gs, x})
	if callErr != nil {
		return nil, callErr
	}
	v332 = v328
	goto b34
b33:
	;
	v332 = vm.NIL
	goto b34
b34:
	;
	nested = v332
	goto b31
b35:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pr-str").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	arg__7368, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pr-str").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("unsupported destructuring pattern: "), arg__7368})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pr-str").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	arg__7381, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pr-str").Deref(), []vm.Value{x})
	if callErr != nil {
		return nil, callErr
	}
	arg__7382, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("unsupported destructuring pattern: "), arg__7381})
	if callErr != nil {
		return nil, callErr
	}
	v400, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__7382})
	if callErr != nil {
		return nil, callErr
	}
	v404 = v400
	goto b37
b36:
	;
	v404 = vm.NIL
	goto b37
b37:
	;
	v412 = v404
	goto b25
}
func fold_binary_chain(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__7395 vm.Value
	var arg__7401 vm.Value
	var arg__7409 vm.Value
	var arg__7417 vm.Value
	var arg__7423 vm.Value
	var arg__7424 vm.Value
	var v36 vm.Value
	var acc vm.Value
	var i int
	var op_kw vm.Value
	var ctx vm.Value
	var args vm.Value
	var arg__7430 vm.Value
	var v51 bool
	var arg__7444 vm.Value
	var arg__7452 vm.Value
	var arg__7461 vm.Value
	var arg__7462 vm.Value
	var v69 vm.Value
	var v70 int
	var v72 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__7395, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__7401, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__7395, arg__7401})
	if callErr != nil {
		return nil, callErr
	}
	arg__7409, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__7417, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__7423, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__7424, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__7417, arg__7423})
	if callErr != nil {
		return nil, callErr
	}
	v36, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{arg2, arg__7409, arg0, arg__7424, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	acc = v36
	i = 2
	op_kw = arg0
	ctx = arg2
	args = arg1
	goto b1
b1:
	;
	arg__7430, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	v51 = rt.GeValue(vm.Int(i), arg__7430)
	if v51 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v72 = acc
	goto b4
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__7444, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{acc, arg__7444})
	if callErr != nil {
		return nil, callErr
	}
	arg__7452, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "ctx-block").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__7461, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__7462, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{acc, arg__7461})
	if callErr != nil {
		return nil, callErr
	}
	v69, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "add-inst!").Deref(), []vm.Value{ctx, arg__7452, op_kw, arg__7462, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v70 = i + 1
	acc = v69
	i = v70
	goto b1
b4:
	;
	return v72, nil
}
func free_vars(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v8 vm.Value
	var form vm.Value
	var bound vm.Value
	var v14 vm.Value
	var or__x_31 vm.Value
	var v629 vm.Value
	var v17 vm.Value
	var v20 vm.Value
	var v22 vm.Value
	var head vm.Value
	var v57 bool
	var v581 vm.Value
	var v625 vm.Value
	var or__x_34 vm.Value
	var v41 vm.Value
	var v43 vm.Value
	var v60 vm.Value
	var v69 bool
	var v571 vm.Value
	var v72 vm.Value
	var v81 bool
	var v566 vm.Value
	var val vm.Value
	var v102 vm.Value
	var v111 bool
	var v561 vm.Value
	var maybe_name vm.Value
	var raw_rest vm.Value
	var has_name_QMARK_ vm.Value
	var or__x_394 bool
	var v556 vm.Value
	var name_sym vm.Value
	var v181 vm.Value
	var rest_forms vm.Value
	var and__x vm.Value
	var arg__7564 vm.Value
	var v223 vm.Value
	var multi_QMARK_ vm.Value
	var arg__7669 vm.Value
	var arg__7769 vm.Value
	var v282 vm.Value
	var args_vec vm.Value
	var body vm.Value
	var v374 vm.Value
	var arg__7781 vm.Value
	var arg__7791 vm.Value
	var arg__7793 vm.Value
	var v328 vm.Value
	var arg__7800 vm.Value
	var v335 vm.Value
	var arg_set vm.Value
	var arg__7822 vm.Value
	var arg__7838 vm.Value
	var v372 vm.Value
	var bindings vm.Value
	var pairs vm.Value
	var arg__7931 vm.Value
	var arg__7996 vm.Value
	var arg__7998 vm.Value
	var vec__7468 vm.Value
	var used vm.Value
	var new_bound vm.Value
	var arg__8048 vm.Value
	var v511 vm.Value
	var v551 vm.Value
	var or__x_398 bool
	var or__x_406 bool
	var v444 bool
	var or__x_410 bool
	var or__x_418 bool
	var v438 bool
	var or__x_422 bool
	var v430 bool
	var v432 bool
	var arg__8069 vm.Value
	var arg__8085 vm.Value
	var v542 vm.Value
	var v546 vm.Value
	var arg__8109 vm.Value
	var arg__8125 vm.Value
	var v604 vm.Value
	var v621 vm.Value
	var v613 vm.Value
	var v617 vm.Value
	var callErr error
	v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v8) {
		form = arg0
		bound = arg1
		goto b1
	} else {
		form = arg0
		bound = arg1
		goto b2
	}
b1:
	;
	v14, callErr = rt.InvokeValue(bound, []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v14) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	or__x_31, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x_31) {
		or__x_34 = or__x_31
		goto b10
	} else {
		goto b11
	}
b3:
	;
	return v629, nil
b4:
	;
	v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v22 = v17
	goto b6
b5:
	;
	v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v22 = v20
	goto b6
b6:
	;
	v629 = v22
	goto b3
b7:
	;
	head, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v57 = head == vm.Symbol("quote")
	if v57 {
		goto b13
	} else {
		goto b14
	}
b8:
	;
	v581, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v581) {
		goto b55
	} else {
		goto b56
	}
b9:
	;
	v629 = v625
	goto b3
b10:
	;
	v43 = or__x_34
	goto b12
b11:
	;
	v41, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v43 = v41
	goto b12
b12:
	;
	if vm.IsTruthy(v43) {
		goto b7
	} else {
		goto b8
	}
b13:
	;
	v60, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v571 = v60
	goto b15
b14:
	;
	v69 = head == vm.Symbol("var")
	if v69 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v625 = v571
	goto b9
b16:
	;
	v72, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v566 = v72
	goto b18
b17:
	;
	v81 = head == vm.Symbol("set!")
	if v81 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v571 = v566
	goto b15
b19:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	val, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(2), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v102, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{val, bound})
	if callErr != nil {
		return nil, callErr
	}
	v561 = v102
	goto b21
b20:
	;
	v111 = head == vm.Symbol("fn*")
	if v111 {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v566 = v561
	goto b18
b22:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	maybe_name, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	raw_rest, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), form})
	if callErr != nil {
		return nil, callErr
	}
	has_name_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{maybe_name})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(has_name_QMARK_) {
		goto b25
	} else {
		goto b26
	}
b23:
	;
	or__x_394 = head == vm.Symbol("let")
	if or__x_394 {
		or__x_398 = or__x_394
		goto b43
	} else {
		goto b44
	}
b24:
	;
	v561 = v556
	goto b21
b25:
	;
	name_sym = maybe_name
	goto b27
b26:
	;
	name_sym = vm.NIL
	goto b27
b27:
	;
	if vm.IsTruthy(has_name_QMARK_) {
		goto b28
	} else {
		goto b29
	}
b28:
	;
	rest_forms = raw_rest
	goto b30
b29:
	;
	v181, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "cons").Deref(), []vm.Value{maybe_name, raw_rest})
	if callErr != nil {
		return nil, callErr
	}
	rest_forms = v181
	goto b30
b30:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b31
	} else {
		goto b32
	}
b31:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	arg__7564, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	v223, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{arg__7564})
	if callErr != nil {
		return nil, callErr
	}
	multi_QMARK_ = v223
	goto b33
b32:
	;
	multi_QMARK_ = and__x
	goto b33
b33:
	;
	if vm.IsTruthy(multi_QMARK_) {
		goto b34
	} else {
		goto b35
	}
b34:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var body vm.Value
		var name_sym_7 vm.Value
		var arg__7628 vm.Value
		var arg__7638 vm.Value
		var arg__7640 vm.Value
		var v28 vm.Value
		var arg__7647 vm.Value
		var v35 vm.Value
		var arg_set vm.Value
		var v49 vm.Value
		var callErr error
		args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(name_sym) {
			name_sym_7 = name_sym
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__7628, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7628, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__7638, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__7640, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7638, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__7640, name_sym_7})
		if callErr != nil {
			return nil, callErr
		}
		arg_set = v28
		goto b3
	b2:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__7647, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		v35, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7647, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		arg_set = v35
		goto b3
	b3:
		;
		v49, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var v3 vm.Value
			var callErr error
			v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, arg_set})
			if callErr != nil {
				return nil, callErr
			}
			return v3, nil
		}), body})
		if callErr != nil {
			return nil, callErr
		}
		return v49, nil
	}), rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	arg__7669, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__7769, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var body vm.Value
		var name_sym_7 vm.Value
		var arg__7731 vm.Value
		var arg__7741 vm.Value
		var arg__7743 vm.Value
		var v28 vm.Value
		var arg__7750 vm.Value
		var v35 vm.Value
		var arg_set vm.Value
		var v49 vm.Value
		var callErr error
		args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(name_sym) {
			name_sym_7 = name_sym
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__7731, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7731, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__7741, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__7743, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7741, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__7743, name_sym_7})
		if callErr != nil {
			return nil, callErr
		}
		arg_set = v28
		goto b3
	b2:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__7750, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		v35, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7750, args_vec})
		if callErr != nil {
			return nil, callErr
		}
		arg_set = v35
		goto b3
	b3:
		;
		v49, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var v3 vm.Value
			var callErr error
			v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, arg_set})
			if callErr != nil {
				return nil, callErr
			}
			return v3, nil
		}), body})
		if callErr != nil {
			return nil, callErr
		}
		return v49, nil
	}), rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	v282, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7669, arg__7769})
	if callErr != nil {
		return nil, callErr
	}
	v374 = v282
	goto b36
b35:
	;
	args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{rest_forms})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(name_sym) {
		goto b37
	} else {
		goto b38
	}
b36:
	;
	v556 = v374
	goto b24
b37:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__7781, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7781, args_vec})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__7791, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__7793, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7791, args_vec})
	if callErr != nil {
		return nil, callErr
	}
	v328, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__7793, name_sym})
	if callErr != nil {
		return nil, callErr
	}
	arg_set = v328
	goto b39
b38:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__7800, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v335, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7800, args_vec})
	if callErr != nil {
		return nil, callErr
	}
	arg_set = v335
	goto b39
b39:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, arg_set})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body})
	if callErr != nil {
		return nil, callErr
	}
	arg__7822, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__7838, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, arg_set})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body})
	if callErr != nil {
		return nil, callErr
	}
	v372, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__7822, arg__7838})
	if callErr != nil {
		return nil, callErr
	}
	v374 = v372
	goto b36
b40:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	bindings, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(2), form})
	if callErr != nil {
		return nil, callErr
	}
	pairs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "partition").Deref(), []vm.Value{vm.Int(2), bindings})
	if callErr != nil {
		return nil, callErr
	}
	arg__7931, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__7931, bound})
	if callErr != nil {
		return nil, callErr
	}
	arg__7996, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__7998, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__7996, bound})
	if callErr != nil {
		return nil, callErr
	}
	vec__7468, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var caps vm.Value
		var b vm.Value
		var sym vm.Value
		var init vm.Value
		var arg__7979 vm.Value
		var arg__7980 vm.Value
		var arg__7991 vm.Value
		var arg__7992 vm.Value
		var v39 vm.Value
		var callErr error
		caps, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		b, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		init, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{init, b})
		if callErr != nil {
			return nil, callErr
		}
		arg__7979, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{init, b})
		if callErr != nil {
			return nil, callErr
		}
		arg__7980, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{caps, arg__7979})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "binding-syms").Deref(), []vm.Value{sym})
		if callErr != nil {
			return nil, callErr
		}
		arg__7991, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "binding-syms").Deref(), []vm.Value{sym})
		if callErr != nil {
			return nil, callErr
		}
		arg__7992, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{b, arg__7991})
		if callErr != nil {
			return nil, callErr
		}
		v39, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__7980, arg__7992})
		if callErr != nil {
			return nil, callErr
		}
		return v39, nil
	}), arg__7998, pairs})
	if callErr != nil {
		return nil, callErr
	}
	used, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__7468, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	new_bound, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__7468, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, new_bound})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body})
	if callErr != nil {
		return nil, callErr
	}
	arg__8048, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, new_bound})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), body})
	if callErr != nil {
		return nil, callErr
	}
	v511, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{used, arg__8048})
	if callErr != nil {
		return nil, callErr
	}
	v551 = v511
	goto b42
b41:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b52
	} else {
		goto b53
	}
b42:
	;
	v556 = v551
	goto b24
b43:
	;
	v444 = or__x_398
	goto b45
b44:
	;
	or__x_406 = head == vm.Symbol("let*")
	if or__x_406 {
		or__x_410 = or__x_406
		goto b46
	} else {
		goto b47
	}
b45:
	;
	if v444 {
		goto b40
	} else {
		goto b41
	}
b46:
	;
	v438 = or__x_410
	goto b48
b47:
	;
	or__x_418 = head == vm.Symbol("loop")
	if or__x_418 {
		or__x_422 = or__x_418
		goto b49
	} else {
		goto b50
	}
b48:
	;
	v444 = v438
	goto b45
b49:
	;
	v432 = or__x_422
	goto b51
b50:
	;
	v430 = head == vm.Symbol("loop*")
	v432 = v430
	goto b51
b51:
	;
	v438 = v432
	goto b48
b52:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, bound})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__8069, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__8085, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, bound})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	v542, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__8069, arg__8085})
	if callErr != nil {
		return nil, callErr
	}
	v546 = v542
	goto b54
b53:
	;
	v546 = vm.NIL
	goto b54
b54:
	;
	v551 = v546
	goto b42
b55:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, bound})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__8109, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__8125, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v3 vm.Value
		var callErr error
		v3, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "free-vars").Deref(), []vm.Value{arg0, bound})
		if callErr != nil {
			return nil, callErr
		}
		return v3, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	v604, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__8109, arg__8125})
	if callErr != nil {
		return nil, callErr
	}
	v621 = v604
	goto b57
b56:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b58
	} else {
		goto b59
	}
b57:
	;
	v625 = v621
	goto b9
b58:
	;
	v613, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v617 = v613
	goto b60
b59:
	;
	v617 = vm.NIL
	goto b60
b60:
	;
	v621 = v617
	goto b57
}
func is_literal_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var or__x vm.Value
	var v vm.Value
	var v61 vm.Value
	var v57 vm.Value
	var v53 vm.Value
	var v49 vm.Value
	var v43 vm.Value
	var v45 vm.Value
	var callErr error
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b1
	} else {
		v = arg0
		goto b2
	}
b1:
	;
	v61 = or__x
	goto b3
b2:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "integer?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v61, nil
b4:
	;
	v57 = or__x
	goto b6
b5:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "float?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v61 = v57
	goto b3
b7:
	;
	v53 = or__x
	goto b9
b8:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v57 = v53
	goto b6
b10:
	;
	v49 = or__x
	goto b12
b11:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v53 = v49
	goto b9
b13:
	;
	v45 = or__x
	goto b15
b14:
	;
	v43, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "boolean?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	v45 = v43
	goto b15
b15:
	;
	v49 = v45
	goto b12
}
func lookup_local(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8154 vm.Value
	var arg__8167 vm.Value
	var arg__8169 vm.Value
	var arg__8170 vm.Value
	var v23 vm.Value
	var i vm.Value
	var sym vm.Value
	var ctx vm.Value
	var v32 bool
	var arg__8182 vm.Value
	var arg__8196 vm.Value
	var arg__8198 vm.Value
	var arg__8212 vm.Value
	var arg__8226 vm.Value
	var arg__8228 vm.Value
	var arg__8230 vm.Value
	var v79 vm.Value
	var v142 vm.Value
	var arg__8241 vm.Value
	var arg__8255 vm.Value
	var arg__8257 vm.Value
	var arg__8271 vm.Value
	var arg__8285 vm.Value
	var arg__8287 vm.Value
	var arg__8289 vm.Value
	var v118 vm.Value
	var v137 vm.Value
	var v128 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__8154, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8154, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__8167, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__8169, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8167, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8170, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__8169})
	if callErr != nil {
		return nil, callErr
	}
	v23 = rt.SubValue(arg__8170, vm.Int(1))
	i = v23
	sym = arg1
	ctx = arg0
	goto b1
b1:
	;
	v32 = rt.LtValue(i, vm.Int(0))
	if v32 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v142 = vm.NIL
	goto b4
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8182, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8182, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8196, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8198, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8196, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8198, i})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8212, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8212, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8226, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8228, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8226, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8230, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8228, i})
	if callErr != nil {
		return nil, callErr
	}
	v79, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg__8230, sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v79) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v142, nil
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8241, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8241, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8255, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8257, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8255, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8257, i})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8271, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8271, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8285, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{ctx})
	if callErr != nil {
		return nil, callErr
	}
	arg__8287, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8285, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8289, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8287, i})
	if callErr != nil {
		return nil, callErr
	}
	v118, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__8289, sym})
	if callErr != nil {
		return nil, callErr
	}
	v137 = v118
	goto b7
b6:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	v142 = v137
	goto b4
b8:
	;
	v128 = rt.SubValue(i, vm.Int(1))
	i = v128
	goto b1
b9:
	;
	goto b10
b10:
	;
	v137 = vm.NIL
	goto b7
}
func new_context(arg0 vm.Value) (vm.Value, error) {
	var arg__8299 vm.Value
	var arg__8303 vm.Value
	var arg__8313 vm.Value
	var arg__8317 vm.Value
	var arg__8318 vm.Value
	var v22 vm.Value
	var callErr error
	arg__8299, callErr = rt.InvokeValue(rt.LookupVar("ir", "entry-block").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__8303, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("fn"), arg0, vm.Keyword("current-block"), arg__8299, vm.Keyword("locals"), arg__8303})
	if callErr != nil {
		return nil, callErr
	}
	arg__8313, callErr = rt.InvokeValue(rt.LookupVar("ir", "entry-block").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__8317, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	arg__8318, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("fn"), arg0, vm.Keyword("current-block"), arg__8313, vm.Keyword("locals"), arg__8317})
	if callErr != nil {
		return nil, callErr
	}
	v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{arg__8318})
	if callErr != nil {
		return nil, callErr
	}
	return v22, nil
}
func parse_defn_args_body(arg0 vm.Value) (vm.Value, error) {
	var x2 vm.Value
	var x3 vm.Value
	var x4 vm.Value
	var v28 vm.Value
	var form vm.Value
	var v39 vm.Value
	var v70 vm.Value
	var v110 vm.Value
	var arg__8353 vm.Value
	var v46 vm.Value
	var arg__8361 vm.Value
	var v53 vm.Value
	var v55 vm.Value
	var arg__8372 vm.Value
	var v77 vm.Value
	var v104 vm.Value
	var arg__8380 vm.Value
	var v94 vm.Value
	var v98 vm.Value
	var callErr error
	x2, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(2), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	x3, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(3), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	x4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(4), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{x2})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v28) {
		form = arg0
		goto b1
	} else {
		form = arg0
		goto b2
	}
b1:
	;
	v39, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{x3})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v39) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v70, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{x2})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v70) {
		goto b7
	} else {
		goto b8
	}
b3:
	;
	return v110, nil
b4:
	;
	arg__8353, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(5), form})
	if callErr != nil {
		return nil, callErr
	}
	v46, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{x4, arg__8353})
	if callErr != nil {
		return nil, callErr
	}
	v55 = v46
	goto b6
b5:
	;
	arg__8361, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(4), form})
	if callErr != nil {
		return nil, callErr
	}
	v53, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{x3, arg__8361})
	if callErr != nil {
		return nil, callErr
	}
	v55 = v53
	goto b6
b6:
	;
	v110 = v55
	goto b3
b7:
	;
	arg__8372, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(4), form})
	if callErr != nil {
		return nil, callErr
	}
	v77, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{x3, arg__8372})
	if callErr != nil {
		return nil, callErr
	}
	v104 = v77
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
	v110 = v104
	goto b3
b10:
	;
	arg__8380, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(3), form})
	if callErr != nil {
		return nil, callErr
	}
	v94, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{x2, arg__8380})
	if callErr != nil {
		return nil, callErr
	}
	v98 = v94
	goto b12
b11:
	;
	v98 = vm.NIL
	goto b12
b12:
	;
	v104 = v98
	goto b9
}
func pop_locals_BANG_(arg0 vm.Value) (vm.Value, error) {
	var s vm.Value
	var stack vm.Value
	var arg__8405 vm.Value
	var arg__8406 vm.Value
	var arg__8424 vm.Value
	var arg__8425 vm.Value
	var arg__8426 vm.Value
	var arg__8445 vm.Value
	var arg__8446 vm.Value
	var arg__8464 vm.Value
	var arg__8465 vm.Value
	var arg__8466 vm.Value
	var arg__8467 vm.Value
	var v56 vm.Value
	var callErr error
	s, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	stack, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8405, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8406 = rt.SubValue(arg__8405, vm.Int(1))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subvec").Deref(), []vm.Value{stack, vm.Int(0), arg__8406})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8424, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8425 = rt.SubValue(arg__8424, vm.Int(1))
	arg__8426, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subvec").Deref(), []vm.Value{stack, vm.Int(0), arg__8425})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{s, vm.Keyword("locals"), arg__8426})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8445, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8446 = rt.SubValue(arg__8445, vm.Int(1))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subvec").Deref(), []vm.Value{stack, vm.Int(0), arg__8446})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8464, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8465 = rt.SubValue(arg__8464, vm.Int(1))
	arg__8466, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subvec").Deref(), []vm.Value{stack, vm.Int(0), arg__8465})
	if callErr != nil {
		return nil, callErr
	}
	arg__8467, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{s, vm.Keyword("locals"), arg__8466})
	if callErr != nil {
		return nil, callErr
	}
	v56, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{arg0, arg__8467})
	if callErr != nil {
		return nil, callErr
	}
	return v56, nil
}
func push_locals_BANG_(arg0 vm.Value) (vm.Value, error) {
	var s vm.Value
	var arg__8487 vm.Value
	var arg__8506 vm.Value
	var arg__8508 vm.Value
	var arg__8527 vm.Value
	var arg__8546 vm.Value
	var arg__8548 vm.Value
	var arg__8549 vm.Value
	var v60 vm.Value
	var callErr error
	s, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8487, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__8487, vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8506, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8508, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__8506, vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{s, vm.Keyword("locals"), arg__8508})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8527, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__8527, vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8546, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8548, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__8546, vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	arg__8549, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{s, vm.Keyword("locals"), arg__8548})
	if callErr != nil {
		return nil, callErr
	}
	v60, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{arg0, arg__8549})
	if callErr != nil {
		return nil, callErr
	}
	return v60, nil
}
func rebind_local_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var s vm.Value
	var stack vm.Value
	var arg__8561 vm.Value
	var v18 vm.Value
	var i vm.Value
	var sym vm.Value
	var ctx vm.Value
	var inst_id vm.Value
	var v33 bool
	var v36 vm.Value
	var arg__8584 vm.Value
	var v55 vm.Value
	var v112 vm.Value
	var arg__8600 vm.Value
	var updated_frame vm.Value
	var new_stack vm.Value
	var arg__8628 vm.Value
	var v76 vm.Value
	var v104 vm.Value
	var v92 vm.Value
	var callErr error
	s, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	stack, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{s, vm.Keyword("locals")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8561, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	v18 = rt.SubValue(arg__8561, vm.Int(1))
	i = v18
	sym = arg1
	ctx = arg0
	inst_id = arg2
	goto b1
b1:
	;
	v33 = rt.LtValue(i, vm.Int(0))
	if v33 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v36, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "bind-local!").Deref(), []vm.Value{ctx, sym, inst_id})
	if callErr != nil {
		return nil, callErr
	}
	v112 = v36
	goto b4
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{stack, i})
	if callErr != nil {
		return nil, callErr
	}
	arg__8584, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{stack, i})
	if callErr != nil {
		return nil, callErr
	}
	v55, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg__8584, sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v55) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v112, nil
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{stack, i})
	if callErr != nil {
		return nil, callErr
	}
	arg__8600, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{stack, i})
	if callErr != nil {
		return nil, callErr
	}
	updated_frame, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg__8600, sym, inst_id})
	if callErr != nil {
		return nil, callErr
	}
	new_stack, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{stack, i, updated_frame})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{s, vm.Keyword("locals"), new_stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__8628, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{s, vm.Keyword("locals"), new_stack})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{ctx, arg__8628})
	if callErr != nil {
		return nil, callErr
	}
	v76, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "attach-name!").Deref(), []vm.Value{ctx, inst_id, sym})
	if callErr != nil {
		return nil, callErr
	}
	v104 = v76
	goto b7
b6:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	v112 = v104
	goto b4
b8:
	;
	v92 = rt.SubValue(i, vm.Int(1))
	i = v92
	goto b1
b9:
	;
	goto b10
b10:
	;
	v104 = vm.NIL
	goto b7
}
func terminated_QMARK_(arg0 vm.Value) bool {
	var v2 bool
	v2 = arg0 == rt.LookupVar("ir.build", "TERMINATED").Deref()
	return v2
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.build", map[string]vm.Value{"add-inst!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 5 {
			return nil, fmt.Errorf("add-inst!: wrong number of arguments %d (expected 5)", len(args))
		}
		return add_inst_BANG_(args[0], args[1], args[2], args[3], args[4])
	}), "add-terminator!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 5 {
			return nil, fmt.Errorf("add-terminator!: wrong number of arguments %d (expected 5)", len(args))
		}
		return add_terminator_BANG_(args[0], args[1], args[2], args[3], args[4])
	}), "attach-name!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("attach-name!: wrong number of arguments %d (expected 3)", len(args))
		}
		return attach_name_BANG_(args[0], args[1], args[2])
	}), "bind-local!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("bind-local!: wrong number of arguments %d (expected 3)", len(args))
		}
		return bind_local_BANG_(args[0], args[1], args[2])
	}), "binding-syms": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("binding-syms: wrong number of arguments %d (expected 1)", len(args))
		}
		return binding_syms(args[0])
	}), "build-args": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-args: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_args(args[0], args[1])
	}), "build-builtin-op": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("build-builtin-op: wrong number of arguments %d (expected 3)", len(args))
		}
		return build_builtin_op(args[0], args[1], args[2])
	}), "build-call": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-call: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_call(args[0], args[1])
	}), "build-call-with-head": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("build-call-with-head: wrong number of arguments %d (expected 3)", len(args))
		}
		return build_call_with_head(args[0], args[1], args[2])
	}), "build-do": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-do: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_do(args[0], args[1])
	}), "build-fn": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("build-fn: wrong number of arguments %d (expected 1)", len(args))
		}
		return build_fn(args[0])
	}), "build-fn*": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-fn*: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_fn_STAR_(args[0], args[1])
	}), "build-form": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-form: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_form(args[0], args[1])
	}), "build-inner-fn-template": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 5 {
			return nil, fmt.Errorf("build-inner-fn-template: wrong number of arguments %d (expected 5)", len(args))
		}
		return build_inner_fn_template(args[0], args[1], args[2], args[3], args[4])
	}), "build-let": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-let: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_let(args[0], args[1])
	}), "build-list": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-list: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_list(args[0], args[1])
	}), "build-loop": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-loop: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_loop(args[0], args[1])
	}), "build-map": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-map: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_map(args[0], args[1])
	}), "build-quote": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-quote: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_quote(args[0], args[1])
	}), "build-recur": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-recur: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_recur(args[0], args[1])
	}), "build-set!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-set!: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_set_BANG_(args[0], args[1])
	}), "build-single-fn*": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("build-single-fn*: wrong number of arguments %d (expected 4)", len(args))
		}
		return build_single_fn_STAR_(args[0], args[1], args[2], args[3])
	}), "build-symbol": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-symbol: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_symbol(args[0], args[1])
	}), "build-var": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-var: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_var(args[0], args[1])
	}), "build-vector": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("build-vector: wrong number of arguments %d (expected 2)", len(args))
		}
		return build_vector(args[0], args[1])
	}), "captures-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("captures-of: wrong number of arguments %d (expected 2)", len(args))
		}
		return captures_of(args[0], args[1])
	}), "ctx-block": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("ctx-block: wrong number of arguments %d (expected 1)", len(args))
		}
		return ctx_block(args[0])
	}), "ctx-fn": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("ctx-fn: wrong number of arguments %d (expected 1)", len(args))
		}
		return ctx_fn(args[0])
	}), "ctx-set-block!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("ctx-set-block!: wrong number of arguments %d (expected 2)", len(args))
		}
		return ctx_set_block_BANG_(args[0], args[1])
	}), "current-locals-flat": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("current-locals-flat: wrong number of arguments %d (expected 1)", len(args))
		}
		return current_locals_flat(args[0])
	}), "emit-template-closure": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("emit-template-closure: wrong number of arguments %d (expected 3)", len(args))
		}
		return emit_template_closure(args[0], args[1], args[2])
	}), "expand-binding": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expand-binding: wrong number of arguments %d (expected 1)", len(args))
		}
		return expand_binding(args[0])
	}), "expand-fn-args": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("expand-fn-args: wrong number of arguments %d (expected 2)", len(args))
		}
		return expand_fn_args(args[0], args[1])
	}), "expand-map-pattern": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("expand-map-pattern: wrong number of arguments %d (expected 2)", len(args))
		}
		return expand_map_pattern(args[0], args[1])
	}), "expand-vector-pattern": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("expand-vector-pattern: wrong number of arguments %d (expected 2)", len(args))
		}
		return expand_vector_pattern(args[0], args[1])
	}), "fold-binary-chain": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("fold-binary-chain: wrong number of arguments %d (expected 3)", len(args))
		}
		return fold_binary_chain(args[0], args[1], args[2])
	}), "free-vars": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("free-vars: wrong number of arguments %d (expected 2)", len(args))
		}
		return free_vars(args[0], args[1])
	}), "is-literal?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("is-literal?: wrong number of arguments %d (expected 1)", len(args))
		}
		return is_literal_QMARK_(args[0])
	}), "lookup-local": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("lookup-local: wrong number of arguments %d (expected 2)", len(args))
		}
		return lookup_local(args[0], args[1])
	}), "new-context": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("new-context: wrong number of arguments %d (expected 1)", len(args))
		}
		return new_context(args[0])
	}), "parse-defn-args-body": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("parse-defn-args-body: wrong number of arguments %d (expected 1)", len(args))
		}
		return parse_defn_args_body(args[0])
	}), "pop-locals!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("pop-locals!: wrong number of arguments %d (expected 1)", len(args))
		}
		return pop_locals_BANG_(args[0])
	}), "push-locals!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("push-locals!: wrong number of arguments %d (expected 1)", len(args))
		}
		return push_locals_BANG_(args[0])
	}), "rebind-local!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("rebind-local!: wrong number of arguments %d (expected 3)", len(args))
		}
		return rebind_local_BANG_(args[0], args[1], args[2])
	}),
	})
}
