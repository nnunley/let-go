package ir_passes_pipeline

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func unwrap_name(arg0 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var raw_name vm.Value
	var v22 vm.Value
	var v25 vm.Value
	var arg__32280 vm.Value
	var v14 bool
	var v17 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		raw_name = arg0
		goto b4
	} else {
		raw_name = arg0
		goto b5
	}
b1:
	;
	v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{raw_name})
	if callErr != nil {
		return nil, callErr
	}
	v25 = v22
	goto b3
b2:
	;
	v25 = raw_name
	goto b3
b3:
	;
	return v25, nil
b4:
	;
	arg__32280, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{raw_name})
	if callErr != nil {
		return nil, callErr
	}
	v14 = arg__32280 == vm.Symbol("with-meta")
	v17 = vm.Boolean(v14)
	goto b6
b5:
	;
	v17 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v17) {
		goto b1
	} else {
		goto b2
	}
}
func override_wrapper_fn_lit(arg0 vm.Value) (vm.Value, error) {
	var arity vm.Value
	var go_name vm.Value
	var fn_name vm.Value
	var needs_error_QMARK_ vm.Value
	var args_id vm.Value
	var arg__32306 vm.Value
	var arg__32309 vm.Value
	var len_call vm.Value
	var msg vm.Value
	var arg__32323 vm.Value
	var arg__32333 vm.Value
	var arg__32340 vm.Value
	var arg__32353 vm.Value
	var arg__32355 vm.Value
	var arg__32360 vm.Value
	var arg__32362 vm.Value
	var arg__32363 vm.Value
	var arg__32370 vm.Value
	var arg__32380 vm.Value
	var arg__32387 vm.Value
	var arg__32400 vm.Value
	var arg__32402 vm.Value
	var arg__32407 vm.Value
	var arg__32409 vm.Value
	var arg__32410 vm.Value
	var arg__32411 vm.Value
	var arity_mismatch vm.Value
	var arg__32425 vm.Value
	var arg__32445 vm.Value
	var arg__32446 vm.Value
	var arg__32449 vm.Value
	var len_check vm.Value
	var arg__32483 vm.Value
	var arg_exprs vm.Value
	var arg__32493 vm.Value
	var inner_call vm.Value
	var arg__32501 vm.Value
	var v190 vm.Value
	var arg__32507 vm.Value
	var arg__32515 vm.Value
	var arg__32516 vm.Value
	var v205 vm.Value
	var ret_stmt vm.Value
	var arg__32528 vm.Value
	var arg__32529 vm.Value
	var arg__32540 vm.Value
	var arg__32541 vm.Value
	var arg__32550 vm.Value
	var arg__32551 vm.Value
	var arg__32569 vm.Value
	var arg__32570 vm.Value
	var arg__32571 vm.Value
	var arg__32581 vm.Value
	var arg__32582 vm.Value
	var arg__32591 vm.Value
	var arg__32592 vm.Value
	var arg__32593 vm.Value
	var arg__32597 vm.Value
	var v297 vm.Value
	var callErr error
	arity, callErr = rt.InvokeValue(vm.Keyword("arity"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	go_name, callErr = rt.InvokeValue(vm.Keyword("go-name"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	fn_name, callErr = rt.InvokeValue(vm.Keyword("fn-name"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	needs_error_QMARK_, callErr = rt.InvokeValue(vm.Keyword("needs-error?"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	args_id, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("len")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{args_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__32306, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("len")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32309, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{args_id})
	if callErr != nil {
		return nil, callErr
	}
	len_call, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__32306, arg__32309})
	if callErr != nil {
		return nil, callErr
	}
	msg, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{fn_name, vm.String(": wrong number of arguments %d (expected "), arity, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32323, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fmt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32333, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fmt")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32333, vm.String("Errorf")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32340, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{msg})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32340, len_call})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fmt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32353, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fmt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32355, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32353, vm.String("Errorf")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32360, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{msg})
	if callErr != nil {
		return nil, callErr
	}
	arg__32362, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32360, len_call})
	if callErr != nil {
		return nil, callErr
	}
	arg__32363, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__32355, arg__32362})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32323, arg__32363})
	if callErr != nil {
		return nil, callErr
	}
	arg__32370, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fmt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32380, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fmt")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32380, vm.String("Errorf")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32387, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{msg})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32387, len_call})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fmt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32400, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fmt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32402, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32400, vm.String("Errorf")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32407, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{msg})
	if callErr != nil {
		return nil, callErr
	}
	arg__32409, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32407, len_call})
	if callErr != nil {
		return nil, callErr
	}
	arg__32410, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__32402, arg__32409})
	if callErr != nil {
		return nil, callErr
	}
	arg__32411, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32370, arg__32410})
	if callErr != nil {
		return nil, callErr
	}
	arity_mismatch, callErr = rt.InvokeValue(rt.LookupVar("gogen", "return-stmt").Deref(), []vm.Value{arg__32411})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{arity})
	if callErr != nil {
		return nil, callErr
	}
	arg__32425, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{arity})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "binary").Deref(), []vm.Value{vm.String("!="), len_call, arg__32425})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arity_mismatch})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{arity})
	if callErr != nil {
		return nil, callErr
	}
	arg__32445, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{arity})
	if callErr != nil {
		return nil, callErr
	}
	arg__32446, callErr = rt.InvokeValue(rt.LookupVar("gogen", "binary").Deref(), []vm.Value{vm.String("!="), len_call, arg__32445})
	if callErr != nil {
		return nil, callErr
	}
	arg__32449, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arity_mismatch})
	if callErr != nil {
		return nil, callErr
	}
	len_check, callErr = rt.InvokeValue(rt.LookupVar("gogen", "if-stmt").Deref(), []vm.Value{vm.NIL, arg__32446, arg__32449, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{arity})
	if callErr != nil {
		return nil, callErr
	}
	arg__32483, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{arity})
	if callErr != nil {
		return nil, callErr
	}
	arg_exprs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__32478 vm.Value
		var v7 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__32478, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v7, callErr = rt.InvokeValue(rt.LookupVar("gogen", "index").Deref(), []vm.Value{args_id, arg__32478})
		if callErr != nil {
			return nil, callErr
		}
		return v7, nil
	}), arg__32483})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{go_name})
	if callErr != nil {
		return nil, callErr
	}
	arg__32493, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{go_name})
	if callErr != nil {
		return nil, callErr
	}
	inner_call, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__32493, arg_exprs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(needs_error_QMARK_) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{inner_call})
	if callErr != nil {
		return nil, callErr
	}
	arg__32501, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{inner_call})
	if callErr != nil {
		return nil, callErr
	}
	v190, callErr = rt.InvokeValue(rt.LookupVar("gogen", "return-stmt").Deref(), []vm.Value{arg__32501})
	if callErr != nil {
		return nil, callErr
	}
	ret_stmt = v190
	goto b3
b2:
	;
	arg__32507, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{inner_call, arg__32507})
	if callErr != nil {
		return nil, callErr
	}
	arg__32515, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32516, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{inner_call, arg__32515})
	if callErr != nil {
		return nil, callErr
	}
	v205, callErr = rt.InvokeValue(rt.LookupVar("gogen", "return-stmt").Deref(), []vm.Value{arg__32516})
	if callErr != nil {
		return nil, callErr
	}
	ret_stmt = v205
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32528, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32529, callErr = rt.InvokeValue(rt.LookupVar("gogen", "param").Deref(), []vm.Value{vm.String("args"), arg__32528})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32529})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32540, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32541, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__32540})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("error")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32550, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("error")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32551, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__32550})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32541, arg__32551})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{len_check, ret_stmt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32569, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32570, callErr = rt.InvokeValue(rt.LookupVar("gogen", "param").Deref(), []vm.Value{vm.String("args"), arg__32569})
	if callErr != nil {
		return nil, callErr
	}
	arg__32571, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32570})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32581, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32582, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__32581})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("error")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32591, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("error")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32592, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__32591})
	if callErr != nil {
		return nil, callErr
	}
	arg__32593, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32582, arg__32592})
	if callErr != nil {
		return nil, callErr
	}
	arg__32597, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{len_check, ret_stmt})
	if callErr != nil {
		return nil, callErr
	}
	v297, callErr = rt.InvokeValue(rt.LookupVar("gogen", "func-lit").Deref(), []vm.Value{arg__32571, arg__32593, arg__32597})
	if callErr != nil {
		return nil, callErr
	}
	return v297, nil
}
func override_init_decl(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var kv_pairs vm.Value
	var arg__32734 vm.Value
	var map_lit vm.Value
	var arg__32745 vm.Value
	var arg__32752 vm.Value
	var arg__32765 vm.Value
	var arg__32767 vm.Value
	var arg__32772 vm.Value
	var arg__32774 vm.Value
	var register_call vm.Value
	var arg__32782 vm.Value
	var arg__32792 vm.Value
	var arg__32793 vm.Value
	var v66 vm.Value
	var callErr error
	kv_pairs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__32668 vm.Value
		var arg__32678 vm.Value
		var arg__32684 vm.Value
		var arg__32689 vm.Value
		var arg__32690 vm.Value
		var arg__32699 vm.Value
		var arg__32700 vm.Value
		var arg__32709 vm.Value
		var arg__32715 vm.Value
		var arg__32720 vm.Value
		var arg__32721 vm.Value
		var arg__32722 vm.Value
		var v50 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(vm.Keyword("fn-name"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__32668, callErr = rt.InvokeValue(vm.Keyword("fn-name"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{arg__32668})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("__gogen_wrap")})
		if callErr != nil {
			return nil, callErr
		}
		arg__32678, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "override-wrapper-fn-lit").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32678})
		if callErr != nil {
			return nil, callErr
		}
		arg__32684, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("__gogen_wrap")})
		if callErr != nil {
			return nil, callErr
		}
		arg__32689, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "override-wrapper-fn-lit").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__32690, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32689})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__32684, arg__32690})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(vm.Keyword("fn-name"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__32699, callErr = rt.InvokeValue(vm.Keyword("fn-name"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__32700, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{arg__32699})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("__gogen_wrap")})
		if callErr != nil {
			return nil, callErr
		}
		arg__32709, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "override-wrapper-fn-lit").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32709})
		if callErr != nil {
			return nil, callErr
		}
		arg__32715, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("__gogen_wrap")})
		if callErr != nil {
			return nil, callErr
		}
		arg__32720, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "override-wrapper-fn-lit").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__32721, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32720})
		if callErr != nil {
			return nil, callErr
		}
		arg__32722, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__32715, arg__32721})
		if callErr != nil {
			return nil, callErr
		}
		v50, callErr = rt.InvokeValue(rt.LookupVar("gogen", "kv-expr").Deref(), []vm.Value{arg__32700, arg__32722})
		if callErr != nil {
			return nil, callErr
		}
		return v50, nil
	}), arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("map[string]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32734, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("map[string]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	map_lit, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit-multi").Deref(), []vm.Value{arg__32734, kv_pairs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32745, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32745, vm.String("RegisterGoOverrides")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32752, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32752, map_lit})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32765, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32767, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32765, vm.String("RegisterGoOverrides")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32772, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__32774, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32772, map_lit})
	if callErr != nil {
		return nil, callErr
	}
	register_call, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__32767, arg__32774})
	if callErr != nil {
		return nil, callErr
	}
	arg__32782, callErr = rt.InvokeValue(rt.LookupVar("gogen", "expr-stmt").Deref(), []vm.Value{register_call})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32782})
	if callErr != nil {
		return nil, callErr
	}
	arg__32792, callErr = rt.InvokeValue(rt.LookupVar("gogen", "expr-stmt").Deref(), []vm.Value{register_call})
	if callErr != nil {
		return nil, callErr
	}
	arg__32793, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32792})
	if callErr != nil {
		return nil, callErr
	}
	v66, callErr = rt.InvokeValue(rt.LookupVar("gogen", "func-decl").Deref(), []vm.Value{vm.String("init"), vm.NewArrayVector([]vm.Value{}), vm.NewArrayVector([]vm.Value{}), arg__32793})
	if callErr != nil {
		return nil, callErr
	}
	return v66, nil
}
func override_helper_decl() (vm.Value, error) {
	var arg__32803 vm.Value
	var arg__32817 vm.Value
	var arg__32819 vm.Value
	var arg__32826 vm.Value
	var arg__32838 vm.Value
	var arg__32852 vm.Value
	var arg__32854 vm.Value
	var arg__32856 vm.Value
	var arg__32861 vm.Value
	var arg__32862 vm.Value
	var wrap_call vm.Value
	var arg__32868 vm.Value
	var arg__32872 vm.Value
	var arg__32883 vm.Value
	var arg__32887 vm.Value
	var arg__32888 vm.Value
	var arg__32891 vm.Value
	var assign vm.Value
	var arg__32896 vm.Value
	var arg__32903 vm.Value
	var arg__32904 vm.Value
	var ret vm.Value
	var arg__32917 vm.Value
	var arg__32918 vm.Value
	var arg__32929 vm.Value
	var arg__32930 vm.Value
	var arg__32949 vm.Value
	var arg__32950 vm.Value
	var arg__32951 vm.Value
	var arg__32961 vm.Value
	var arg__32962 vm.Value
	var arg__32963 vm.Value
	var arg__32967 vm.Value
	var v172 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32803, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32803, vm.String("NativeFnType")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32817, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32819, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32817, vm.String("NativeFnType")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32819, vm.String("Wrap")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32826, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fn")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32826})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32838, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32838, vm.String("NativeFnType")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32852, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32854, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32852, vm.String("NativeFnType")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32856, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__32854, vm.String("Wrap")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32861, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("fn")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32862, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32861})
	if callErr != nil {
		return nil, callErr
	}
	wrap_call, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__32856, arg__32862})
	if callErr != nil {
		return nil, callErr
	}
	arg__32868, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("v")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32872, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("_")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32868, arg__32872})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{wrap_call})
	if callErr != nil {
		return nil, callErr
	}
	arg__32883, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("v")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32887, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("_")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32888, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32883, arg__32887})
	if callErr != nil {
		return nil, callErr
	}
	arg__32891, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{wrap_call})
	if callErr != nil {
		return nil, callErr
	}
	assign, callErr = rt.InvokeValue(rt.LookupVar("gogen", "multi-assign").Deref(), []vm.Value{vm.String(":="), arg__32888, arg__32891})
	if callErr != nil {
		return nil, callErr
	}
	arg__32896, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("v")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32896})
	if callErr != nil {
		return nil, callErr
	}
	arg__32903, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("v")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32904, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32903})
	if callErr != nil {
		return nil, callErr
	}
	ret, callErr = rt.InvokeValue(rt.LookupVar("gogen", "return-stmt").Deref(), []vm.Value{arg__32904})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("func(args []vm.Value) (vm.Value, error)")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32917, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("func(args []vm.Value) (vm.Value, error)")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32918, callErr = rt.InvokeValue(rt.LookupVar("gogen", "param").Deref(), []vm.Value{vm.String("fn"), arg__32917})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32918})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32929, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32930, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__32929})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32930})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{assign, ret})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("func(args []vm.Value) (vm.Value, error)")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32949, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("func(args []vm.Value) (vm.Value, error)")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32950, callErr = rt.InvokeValue(rt.LookupVar("gogen", "param").Deref(), []vm.Value{vm.String("fn"), arg__32949})
	if callErr != nil {
		return nil, callErr
	}
	arg__32951, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32950})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32961, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32962, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__32961})
	if callErr != nil {
		return nil, callErr
	}
	arg__32963, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__32962})
	if callErr != nil {
		return nil, callErr
	}
	arg__32967, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{assign, ret})
	if callErr != nil {
		return nil, callErr
	}
	v172, callErr = rt.InvokeValue(rt.LookupVar("gogen", "func-decl").Deref(), []vm.Value{vm.String("__gogen_wrap"), arg__32951, arg__32963, arg__32967})
	if callErr != nil {
		return nil, callErr
	}
	return v172, nil
}
func override_entries(arg0 vm.Value) (vm.Value, error) {
	var v5 vm.Value
	var callErr error
	v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filterv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var r vm.Value
		var v9 vm.Value
		var v12 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(vm.Keyword("fn-name"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			r = arg0
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v9, callErr = rt.InvokeValue(vm.Keyword("override-eligible?"), []vm.Value{r})
		if callErr != nil {
			return nil, callErr
		}
		v12 = v9
		goto b3
	b2:
		;
		v12 = and__x
		goto b3
	b3:
		;
		return v12, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v5, nil
}
func collect_call_targets(arg0 vm.Value) (vm.Value, error) {
	var arg__32992 vm.Value
	var acc vm.Value
	var letfn__32981 vm.Value
	var v26 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__32992, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	acc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{arg__32992})
	if callErr != nil {
		return nil, callErr
	}
	letfn__32981, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{letfn__32981, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v10 vm.Value
		var x vm.Value
		var acc_4 vm.Value
		var visit_5 vm.Value
		var head vm.Value
		var v23 vm.Value
		var visit_8 vm.Value
		var v86 vm.Value
		var acc_15 vm.Value
		var visit_16 vm.Value
		var arg__33117 vm.Value
		var visit_20 vm.Value
		var visit_39 vm.Value
		var arg__33126 vm.Value
		var doseq_seq__32982 vm.Value
		var doseq_loop__32983 vm.Value
		var visit_48 vm.Value
		var visit_55 vm.Value
		var el vm.Value
		var v67 vm.Value
		var visit_81 vm.Value
		var doseq_seq__32984 vm.Value
		var visit_84 vm.Value
		var v126 vm.Value
		var doseq_loop__32985 vm.Value
		var visit_91 vm.Value
		var visit_97 vm.Value
		var v108 vm.Value
		var visit_121 vm.Value
		var doseq_seq__32986 vm.Value
		var doseq_loop__32987 vm.Value
		var visit_131 vm.Value
		var visit_137 vm.Value
		var entry vm.Value
		var arg__33168 vm.Value
		var arg__33177 vm.Value
		var v157 vm.Value
		var callErr error
		v10, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq?").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v10) {
			x = arg0
			acc_4 = acc
			visit_5 = rt.BoxNativeFn(func(args ...vm.Value) (vm.Value, error) {
				var arg__33005 vm.Value
				var v7 vm.Value
				var callErr error
				_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{letfn__32981})
				if callErr != nil {
					return nil, callErr
				}
				arg__33005, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{letfn__32981})
				if callErr != nil {
					return nil, callErr
				}
				v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{arg__33005, rt.BoxRestArgs(args)})
				if callErr != nil {
					return nil, callErr
				}
				return v7, nil
			})
			goto b1
		} else {
			x = arg0
			visit_8 = rt.BoxNativeFn(func(args ...vm.Value) (vm.Value, error) {
				var arg__33005 vm.Value
				var v7 vm.Value
				var callErr error
				_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{letfn__32981})
				if callErr != nil {
					return nil, callErr
				}
				arg__33005, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{letfn__32981})
				if callErr != nil {
					return nil, callErr
				}
				v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{arg__33005, rt.BoxRestArgs(args)})
				if callErr != nil {
					return nil, callErr
				}
				return v7, nil
			})
			goto b2
		}
	b1:
		;
		head, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		v23, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{head})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v23) {
			acc_15 = acc_4
			visit_16 = visit_5
			goto b4
		} else {
			visit_20 = visit_5
			goto b5
		}
	b2:
		;
		v86, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v86) {
			visit_81 = visit_8
			goto b11
		} else {
			visit_84 = visit_8
			goto b12
		}
	b3:
		;
		return vm.NIL, nil
	b4:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{head})
		if callErr != nil {
			return nil, callErr
		}
		arg__33117, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{head})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{acc_15, rt.LookupVar("clojure.core", "conj").Deref(), arg__33117})
		if callErr != nil {
			return nil, callErr
		}
		visit_39 = visit_16
		goto b6
	b5:
		;
		visit_39 = visit_20
		goto b6
	b6:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		arg__33126, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		doseq_seq__32982, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__33126})
		if callErr != nil {
			return nil, callErr
		}
		doseq_loop__32983 = doseq_seq__32982
		visit_48 = visit_39
		goto b7
	b7:
		;
		if vm.IsTruthy(doseq_loop__32983) {
			visit_55 = visit_48
			goto b8
		} else {
			goto b9
		}
	b8:
		;
		el, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__32983})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(visit_55, []vm.Value{el})
		if callErr != nil {
			return nil, callErr
		}
		v67, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__32983})
		if callErr != nil {
			return nil, callErr
		}
		doseq_loop__32983 = v67
		visit_48 = visit_55
		goto b7
	b9:
		;
		goto b10
	b10:
		;
		goto b3
	b11:
		;
		doseq_seq__32984, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		doseq_loop__32985 = doseq_seq__32984
		visit_91 = visit_81
		goto b14
	b12:
		;
		v126, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v126) {
			visit_121 = visit_84
			goto b18
		} else {
			goto b19
		}
	b13:
		;
		goto b3
	b14:
		;
		if vm.IsTruthy(doseq_loop__32985) {
			visit_97 = visit_91
			goto b15
		} else {
			goto b16
		}
	b15:
		;
		el, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__32985})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(visit_97, []vm.Value{el})
		if callErr != nil {
			return nil, callErr
		}
		v108, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__32985})
		if callErr != nil {
			return nil, callErr
		}
		doseq_loop__32985 = v108
		visit_91 = visit_97
		goto b14
	b16:
		;
		goto b17
	b17:
		;
		goto b13
	b18:
		;
		doseq_seq__32986, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{x})
		if callErr != nil {
			return nil, callErr
		}
		doseq_loop__32987 = doseq_seq__32986
		visit_131 = visit_121
		goto b21
	b19:
		;
		if vm.IsTruthy(vm.Keyword("else")) {
			goto b25
		} else {
			goto b26
		}
	b20:
		;
		goto b13
	b21:
		;
		if vm.IsTruthy(doseq_loop__32987) {
			visit_137 = visit_131
			goto b22
		} else {
			goto b23
		}
	b22:
		;
		entry, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__32987})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{entry})
		if callErr != nil {
			return nil, callErr
		}
		arg__33168, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{entry})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(visit_137, []vm.Value{arg__33168})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{entry})
		if callErr != nil {
			return nil, callErr
		}
		arg__33177, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{entry})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(visit_137, []vm.Value{arg__33177})
		if callErr != nil {
			return nil, callErr
		}
		v157, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__32987})
		if callErr != nil {
			return nil, callErr
		}
		doseq_loop__32987 = v157
		visit_131 = visit_137
		goto b21
	b23:
		;
		goto b24
	b24:
		;
		goto b20
	b25:
		;
		goto b27
	b26:
		;
		goto b27
	b27:
		;
		goto b20
	})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(args ...vm.Value) (vm.Value, error) {
		var arg__33005 vm.Value
		var v7 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{letfn__32981})
		if callErr != nil {
			return nil, callErr
		}
		arg__33005, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{letfn__32981})
		if callErr != nil {
			return nil, callErr
		}
		v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{arg__33005, rt.BoxRestArgs(args)})
		if callErr != nil {
			return nil, callErr
		}
		return v7, nil
	}), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{acc})
	if callErr != nil {
		return nil, callErr
	}
	return v26, nil
}
func defn_key(arg0 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var form vm.Value
	var raw_name vm.Value
	var name_sym vm.Value
	var name_str vm.Value
	var maybe_doc vm.Value
	var has_doc_QMARK_ vm.Value
	var v179 vm.Value
	var arg__33194 vm.Value
	var v13 bool
	var v16 vm.Value
	var arg__33218 vm.Value
	var v62 bool
	var args_or_arity vm.Value
	var v102 vm.Value
	var v67 vm.Value
	var v71 vm.Value
	var arg__33233 vm.Value
	var v107 vm.Value
	var v124 vm.Value
	var v168 vm.Value
	var v128 vm.Value
	var v159 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		form = arg0
		goto b4
	} else {
		form = arg0
		goto b5
	}
b1:
	;
	raw_name, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	name_sym, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "unwrap-name").Deref(), []vm.Value{raw_name})
	if callErr != nil {
		return nil, callErr
	}
	name_str, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{name_sym})
	if callErr != nil {
		return nil, callErr
	}
	maybe_doc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	has_doc_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{maybe_doc})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(has_doc_QMARK_) {
		goto b7
	} else {
		goto b8
	}
b2:
	;
	v179 = vm.NIL
	goto b3
b3:
	;
	return v179, nil
b4:
	;
	arg__33194, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v13 = rt.GeValue(arg__33194, vm.Int(3))
	v16 = vm.Boolean(v13)
	goto b6
b5:
	;
	v16 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v16) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	arg__33218, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v62 = rt.GeValue(arg__33218, vm.Int(4))
	if v62 {
		goto b10
	} else {
		goto b11
	}
b8:
	;
	args_or_arity = maybe_doc
	goto b9
b9:
	;
	v102, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{args_or_arity})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v102) {
		goto b13
	} else {
		goto b14
	}
b10:
	;
	v67, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(3)})
	if callErr != nil {
		return nil, callErr
	}
	v71 = v67
	goto b12
b11:
	;
	v71 = vm.NIL
	goto b12
b12:
	;
	args_or_arity = v71
	goto b9
b13:
	;
	arg__33233, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args_or_arity})
	if callErr != nil {
		return nil, callErr
	}
	v107, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{name_str, arg__33233})
	if callErr != nil {
		return nil, callErr
	}
	v168 = v107
	goto b15
b14:
	;
	v124, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{args_or_arity})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v124) {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v179 = v168
	goto b3
b16:
	;
	v128, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{name_str, vm.Keyword("multi")})
	if callErr != nil {
		return nil, callErr
	}
	v159 = v128
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
	v168 = v159
	goto b15
b19:
	;
	goto b21
b20:
	;
	goto b21
b21:
	;
	v159 = vm.NIL
	goto b18
}
func forms_by_arity(arg0 vm.Value) (vm.Value, error) {
	var remaining vm.Value
	var acc vm.Value
	var v12 vm.Value
	var form vm.Value
	var k vm.Value
	var v20 vm.Value
	var v43 vm.Value
	var v33 vm.Value
	var v36 vm.Value
	var callErr error
	remaining = arg0
	acc = vm.EmptyPersistentMap
	goto b1
b1:
	;
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{remaining})
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
	v43 = acc
	goto b4
b3:
	;
	form, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	k, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(k) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v43, nil
b5:
	;
	v33, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{acc, k, form})
	if callErr != nil {
		return nil, callErr
	}
	v36 = v33
	goto b7
b6:
	;
	v36 = acc
	goto b7
b7:
	;
	remaining = v20
	acc = v36
	goto b1
}
func order_defn_forms(arg0 vm.Value) (vm.Value, error) {
	var keyed vm.Value
	var arg__33274 vm.Value
	var arg__33287 vm.Value
	var arg__33288 vm.Value
	var own_names vm.Value
	var arg__33352 vm.Value
	var arg__33365 vm.Value
	var arg__33366 vm.Value
	var sorted_ks vm.Value
	var sorted_forms vm.Value
	var topo vm.Value
	var forms vm.Value
	var or__x vm.Value
	var ordered vm.Value
	var arg__33481 vm.Value
	var non_defn vm.Value
	var v87 vm.Value
	var callErr error
	keyed, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "forms-by-arity").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{keyed})
	if callErr != nil {
		return nil, callErr
	}
	arg__33274, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{keyed})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("clojure.core", "first").Deref(), arg__33274})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{keyed})
	if callErr != nil {
		return nil, callErr
	}
	arg__33287, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{keyed})
	if callErr != nil {
		return nil, callErr
	}
	arg__33288, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("clojure.core", "first").Deref(), arg__33287})
	if callErr != nil {
		return nil, callErr
	}
	own_names, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "set").Deref(), []vm.Value{arg__33288})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{keyed})
	if callErr != nil {
		return nil, callErr
	}
	arg__33352, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{keyed})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var na vm.Value
		var aa vm.Value
		var nb vm.Value
		var ab vm.Value
		var c vm.Value
		var v43 vm.Value
		var arg__33337 vm.Value
		var arg__33341 vm.Value
		var v54 vm.Value
		var v57 vm.Value
		var callErr error
		na, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		aa, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		nb, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		ab, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		c, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "compare").Deref(), []vm.Value{na, nb})
		if callErr != nil {
			return nil, callErr
		}
		v43, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{c})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v43) {
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{aa})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{ab})
		if callErr != nil {
			return nil, callErr
		}
		arg__33337, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{aa})
		if callErr != nil {
			return nil, callErr
		}
		arg__33341, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{ab})
		if callErr != nil {
			return nil, callErr
		}
		v54, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "compare").Deref(), []vm.Value{arg__33337, arg__33341})
		if callErr != nil {
			return nil, callErr
		}
		v57 = v54
		goto b3
	b2:
		;
		v57 = c
		goto b3
	b3:
		;
		return v57, nil
	}), arg__33352})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{keyed})
	if callErr != nil {
		return nil, callErr
	}
	arg__33365, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keys").Deref(), []vm.Value{keyed})
	if callErr != nil {
		return nil, callErr
	}
	arg__33366, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var na vm.Value
		var aa vm.Value
		var nb vm.Value
		var ab vm.Value
		var c vm.Value
		var v43 vm.Value
		var arg__33337 vm.Value
		var arg__33341 vm.Value
		var v54 vm.Value
		var v57 vm.Value
		var callErr error
		na, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		aa, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		nb, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		ab, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		c, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "compare").Deref(), []vm.Value{na, nb})
		if callErr != nil {
			return nil, callErr
		}
		v43, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{c})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v43) {
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{aa})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{ab})
		if callErr != nil {
			return nil, callErr
		}
		arg__33337, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{aa})
		if callErr != nil {
			return nil, callErr
		}
		arg__33341, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{ab})
		if callErr != nil {
			return nil, callErr
		}
		v54, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "compare").Deref(), []vm.Value{arg__33337, arg__33341})
		if callErr != nil {
			return nil, callErr
		}
		v57 = v54
		goto b3
	b2:
		;
		v57 = c
		goto b3
	b3:
		;
		return v57, nil
	}), arg__33365})
	if callErr != nil {
		return nil, callErr
	}
	sorted_ks, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__33366})
	if callErr != nil {
		return nil, callErr
	}
	sorted_forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{keyed, sorted_ks})
	if callErr != nil {
		return nil, callErr
	}
	topo, callErr = rt.InvokeValue(rt.LookupVar("graph", "toposort-by").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__33429 vm.Value
		var v6 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__33429, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg__33429})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), sorted_forms, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__33380 vm.Value
		var self_name vm.Value
		var arg__33407 vm.Value
		var v21 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__33380, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		self_name, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg__33380})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "collect-call-targets").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__33407, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "collect-call-targets").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v21, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filterv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var and__x vm.Value
			var n vm.Value
			var own_names_6 vm.Value
			var v14 vm.Value
			var v17 vm.Value
			var callErr error
			and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{arg0, self_name})
			if callErr != nil {
				return nil, callErr
			}
			if vm.IsTruthy(and__x) {
				n = arg0
				own_names_6 = own_names
				goto b1
			} else {
				goto b2
			}
		b1:
			;
			v14, callErr = rt.InvokeValue(own_names_6, []vm.Value{n})
			if callErr != nil {
				return nil, callErr
			}
			v17 = v14
			goto b3
		b2:
			;
			v17 = and__x
			goto b3
		b3:
			;
			return v17, nil
		}), arg__33407})
		if callErr != nil {
			return nil, callErr
		}
		return v21, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(topo) {
		forms = arg0
		or__x = topo
		goto b1
	} else {
		forms = arg0
		goto b2
	}
b1:
	;
	ordered = or__x
	goto b3
b2:
	;
	ordered = sorted_forms
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "remove").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__33453 vm.Value
		var v6 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__33453, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some?").Deref(), []vm.Value{arg__33453})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), forms})
	if callErr != nil {
		return nil, callErr
	}
	arg__33481, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "remove").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__33478 vm.Value
		var v6 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__33478, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "defn-key").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some?").Deref(), []vm.Value{arg__33478})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), forms})
	if callErr != nil {
		return nil, callErr
	}
	non_defn, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__33481})
	if callErr != nil {
		return nil, callErr
	}
	v87, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{ordered, non_defn})
	if callErr != nil {
		return nil, callErr
	}
	return v87, nil
}
func compile_form_STAR_(arg0 vm.Value) (vm.Value, error) {
	var arg__33499 vm.Value
	var name_sym vm.Value
	var maybe_doc vm.Value
	var has_doc_QMARK_ vm.Value
	var form vm.Value
	var v30 vm.Value
	var first_form vm.Value
	var go_target_QMARK__40 bool
	var v54 vm.Value
	var go_target_QMARK__46 bool
	var go_target_QMARK__52 bool
	var v359 vm.Value
	var args_vec vm.Value
	var go_target_QMARK__63 bool
	var go_target_QMARK__71 bool
	var go_target_QMARK__85 bool
	var go_target_QMARK__94 bool
	var head__33521 vm.Value
	var go_target_QMARK__103 bool
	var arg__33522 int
	var go_target_QMARK__118 bool
	var go_target_QMARK__130 bool
	var head__33525 vm.Value
	var arg__33526 vm.Value
	var go_target_QMARK__139 bool
	var go_target_QMARK__154 bool
	var head__33529 vm.Value
	var arg__33530 int
	var arg__33532 vm.Value
	var body_forms vm.Value
	var expanded vm.Value
	var arg__33552 vm.Value
	var ir_fn vm.Value
	var v232 vm.Value
	var v235 vm.Value
	var v237 vm.Value
	var go_target_QMARK__254 bool
	var go_target_QMARK__260 bool
	var go_target_QMARK__272 bool
	var go_target_QMARK__279 bool
	var head__33563 vm.Value
	var go_target_QMARK__286 bool
	var arg__33564 int
	var go_target_QMARK__299 bool
	var arities vm.Value
	var fn_templates vm.Value
	var v329 vm.Value
	var fn_vals vm.Value
	var arg__33807 vm.Value
	var v348 vm.Value
	var v350 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__33499, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	name_sym, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "unwrap-name").Deref(), []vm.Value{arg__33499})
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
		form = arg0
		goto b1
	} else {
		form = arg0
		goto b2
	}
b1:
	;
	v30, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{form, vm.Int(3)})
	if callErr != nil {
		return nil, callErr
	}
	first_form = v30
	goto b3
b2:
	;
	first_form = maybe_doc
	goto b3
b3:
	;
	go_target_QMARK__40 = rt.LookupVar("ir.passes.pipeline", "*target*").Deref() == vm.Keyword("go")
	v54, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{first_form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v54) {
		go_target_QMARK__46 = go_target_QMARK__40
		goto b4
	} else {
		go_target_QMARK__52 = go_target_QMARK__40
		goto b5
	}
b4:
	;
	if vm.IsTruthy(has_doc_QMARK_) {
		args_vec = first_form
		go_target_QMARK__63 = go_target_QMARK__46
		goto b7
	} else {
		args_vec = first_form
		go_target_QMARK__71 = go_target_QMARK__46
		goto b8
	}
b5:
	;
	if vm.IsTruthy(has_doc_QMARK_) {
		go_target_QMARK__254 = go_target_QMARK__52
		goto b22
	} else {
		go_target_QMARK__260 = go_target_QMARK__52
		goto b23
	}
b6:
	;
	return v359, nil
b7:
	;
	go_target_QMARK__85 = go_target_QMARK__63
	goto b9
b8:
	;
	go_target_QMARK__85 = go_target_QMARK__71
	goto b9
b9:
	;
	if vm.IsTruthy(has_doc_QMARK_) {
		go_target_QMARK__94 = go_target_QMARK__85
		head__33521 = rt.LookupVar("clojure.core", "drop").Deref()
		goto b10
	} else {
		go_target_QMARK__103 = go_target_QMARK__85
		head__33521 = rt.LookupVar("clojure.core", "drop").Deref()
		goto b11
	}
b10:
	;
	arg__33522 = 4
	go_target_QMARK__118 = go_target_QMARK__94
	goto b12
b11:
	;
	arg__33522 = 3
	go_target_QMARK__118 = go_target_QMARK__103
	goto b12
b12:
	;
	_, callErr = rt.InvokeValue(head__33521, []vm.Value{vm.Int(arg__33522), form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(has_doc_QMARK_) {
		go_target_QMARK__130 = go_target_QMARK__118
		head__33525 = rt.LookupVar("clojure.core", "map").Deref()
		arg__33526 = rt.LookupVar("ir.passes.pipeline", "expand-all").Deref()
		goto b13
	} else {
		go_target_QMARK__139 = go_target_QMARK__118
		head__33525 = rt.LookupVar("clojure.core", "map").Deref()
		arg__33526 = rt.LookupVar("ir.passes.pipeline", "expand-all").Deref()
		goto b14
	}
b13:
	;
	go_target_QMARK__154 = go_target_QMARK__130
	goto b15
b14:
	;
	go_target_QMARK__154 = go_target_QMARK__139
	goto b15
b15:
	;
	if vm.IsTruthy(has_doc_QMARK_) {
		head__33529 = rt.LookupVar("clojure.core", "drop").Deref()
		goto b16
	} else {
		head__33529 = rt.LookupVar("clojure.core", "drop").Deref()
		goto b17
	}
b16:
	;
	arg__33530 = 4
	goto b18
b17:
	;
	arg__33530 = 3
	goto b18
b18:
	;
	arg__33532, callErr = rt.InvokeValue(head__33529, []vm.Value{vm.Int(arg__33530), form})
	if callErr != nil {
		return nil, callErr
	}
	body_forms, callErr = rt.InvokeValue(head__33525, []vm.Value{arg__33526, arg__33532})
	if callErr != nil {
		return nil, callErr
	}
	expanded, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), vm.Symbol("defn"), name_sym, args_vec, body_forms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	arg__33552, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	ir_fn, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "optimize-fn").Deref(), []vm.Value{arg__33552})
	if callErr != nil {
		return nil, callErr
	}
	if go_target_QMARK__154 {
		goto b19
	} else {
		goto b20
	}
b19:
	;
	v232, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower").Deref(), []vm.Value{ir_fn, vm.Keyword("bridge")})
	if callErr != nil {
		return nil, callErr
	}
	v237 = v232
	goto b21
b20:
	;
	v235, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "lower").Deref(), []vm.Value{ir_fn})
	if callErr != nil {
		return nil, callErr
	}
	v237 = v235
	goto b21
b21:
	;
	v359 = v237
	goto b6
b22:
	;
	go_target_QMARK__272 = go_target_QMARK__254
	goto b24
b23:
	;
	go_target_QMARK__272 = go_target_QMARK__260
	goto b24
b24:
	;
	if vm.IsTruthy(has_doc_QMARK_) {
		go_target_QMARK__279 = go_target_QMARK__272
		head__33563 = rt.LookupVar("clojure.core", "drop").Deref()
		goto b25
	} else {
		go_target_QMARK__286 = go_target_QMARK__272
		head__33563 = rt.LookupVar("clojure.core", "drop").Deref()
		goto b26
	}
b25:
	;
	arg__33564 = 3
	go_target_QMARK__299 = go_target_QMARK__279
	goto b27
b26:
	;
	arg__33564 = 2
	go_target_QMARK__299 = go_target_QMARK__286
	goto b27
b27:
	;
	arities, callErr = rt.InvokeValue(head__33563, []vm.Value{vm.Int(arg__33564), form})
	if callErr != nil {
		return nil, callErr
	}
	if go_target_QMARK__299 {
		goto b28
	} else {
		goto b29
	}
b28:
	;
	fn_templates, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var arg__33630 vm.Value
		var body_forms vm.Value
		var expanded vm.Value
		var arg__33650 vm.Value
		var ir_fn vm.Value
		var result vm.Value
		var arg__33659 vm.Value
		var v45 vm.Value
		var arg__33664 vm.Value
		var v51 vm.Value
		var v54 vm.Value
		var callErr error
		args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__33630, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		body_forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__33630})
		if callErr != nil {
			return nil, callErr
		}
		expanded, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), vm.Symbol("defn"), name_sym, args_vec, body_forms})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn").Deref(), []vm.Value{expanded})
		if callErr != nil {
			return nil, callErr
		}
		arg__33650, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn").Deref(), []vm.Value{expanded})
		if callErr != nil {
			return nil, callErr
		}
		ir_fn, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "optimize-fn").Deref(), []vm.Value{arg__33650})
		if callErr != nil {
			return nil, callErr
		}
		result, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower").Deref(), []vm.Value{ir_fn, vm.Keyword("bridge")})
		if callErr != nil {
			return nil, callErr
		}
		arg__33659, callErr = rt.InvokeValue(vm.Keyword("status"), []vm.Value{result})
		if callErr != nil {
			return nil, callErr
		}
		v45 = vm.Boolean(vm.Keyword("lowered") == arg__33659)
		if vm.IsTruthy(v45) {
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		arg__33664, callErr = rt.InvokeValue(vm.Keyword("decl"), []vm.Value{result})
		if callErr != nil {
			return nil, callErr
		}
		v51, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("fn"), arg__33664})
		if callErr != nil {
			return nil, callErr
		}
		v54 = v51
		goto b3
	b2:
		;
		v54 = result
		goto b3
	b3:
		;
		return v54, nil
	}), arities})
	if callErr != nil {
		return nil, callErr
	}
	v329, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("fns"), fn_templates, vm.Keyword("kind"), vm.Keyword("multi-fn-template")})
	if callErr != nil {
		return nil, callErr
	}
	v350 = v329
	goto b30
b29:
	;
	fn_vals, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var args_vec vm.Value
		var arg__33747 vm.Value
		var body_forms vm.Value
		var expanded vm.Value
		var arg__33767 vm.Value
		var arg__33778 vm.Value
		var arg__33779 vm.Value
		var chunk vm.Value
		var arg__33790 vm.Value
		var v39 vm.Value
		var callErr error
		args_vec, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__33747, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		body_forms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__33747})
		if callErr != nil {
			return nil, callErr
		}
		expanded, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), vm.Symbol("defn"), name_sym, args_vec, body_forms})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn").Deref(), []vm.Value{expanded})
		if callErr != nil {
			return nil, callErr
		}
		arg__33767, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn").Deref(), []vm.Value{expanded})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "optimize-fn").Deref(), []vm.Value{arg__33767})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn").Deref(), []vm.Value{expanded})
		if callErr != nil {
			return nil, callErr
		}
		arg__33778, callErr = rt.InvokeValue(rt.LookupVar("ir.build", "build-fn").Deref(), []vm.Value{expanded})
		if callErr != nil {
			return nil, callErr
		}
		arg__33779, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "optimize-fn").Deref(), []vm.Value{arg__33778})
		if callErr != nil {
			return nil, callErr
		}
		chunk, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "lower").Deref(), []vm.Value{arg__33779})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args_vec})
		if callErr != nil {
			return nil, callErr
		}
		arg__33790, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args_vec})
		if callErr != nil {
			return nil, callErr
		}
		v39, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "chunk->fn").Deref(), []vm.Value{arg__33790, vm.FALSE, chunk})
		if callErr != nil {
			return nil, callErr
		}
		return v39, nil
	}), arities})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), fn_vals})
	if callErr != nil {
		return nil, callErr
	}
	arg__33807, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), fn_vals})
	if callErr != nil {
		return nil, callErr
	}
	v348, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "make-multi-arity").Deref(), []vm.Value{arg__33807})
	if callErr != nil {
		return nil, callErr
	}
	v350 = v348
	goto b30
b30:
	;
	v359 = v350
	goto b6
}
func expand_fully(arg0 vm.Value) (vm.Value, error) {
	var f vm.Value
	var e vm.Value
	var v11 bool
	var v15 vm.Value
	var callErr error
	f = arg0
	goto b1
b1:
	;
	e, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "macroexpand").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v11 = e == f
	if v11 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v15 = e
	goto b4
b3:
	;
	f = e
	goto b1
b4:
	;
	return v15, nil
}
func go_name_for(arg0 vm.Value) (vm.Value, error) {
	var s vm.Value
	var arg__33893 vm.Value
	var v15 vm.Value
	var callErr error
	s, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var v4 vm.Value
		var c vm.Value
		var v9 vm.Value
		v4 = vm.Boolean(arg0 == vm.Char('-'))
		if vm.IsTruthy(v4) {
			goto b1
		} else {
			c = arg0
			goto b2
		}
	b1:
		;
		v9 = vm.Char('_')
		goto b3
	b2:
		;
		v9 = c
		goto b3
	b3:
		;
		return v9
	}), s})
	if callErr != nil {
		return nil, callErr
	}
	arg__33893, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var v4 vm.Value
		var c vm.Value
		var v9 vm.Value
		v4 = vm.Boolean(arg0 == vm.Char('-'))
		if vm.IsTruthy(v4) {
			goto b1
		} else {
			c = arg0
			goto b2
		}
	b1:
		;
		v9 = vm.Char('_')
		goto b3
	b2:
		;
		v9 = c
		goto b3
	b3:
		;
		return v9
	}), s})
	if callErr != nil {
		return nil, callErr
	}
	v15, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__33893})
	if callErr != nil {
		return nil, callErr
	}
	return v15, nil
}
func optimize_fn(arg0 vm.Value) (vm.Value, error) {
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("build")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.infer-arg-types", "infer-arg-types").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("infer-arg-types")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "typeinfer").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "constfold").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("constfold-1")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.cse", "cse").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("cse")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "constfold").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("constfold-2")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "licm").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("licm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "constfold").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("constfold-3")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.dce", "dce").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("dce")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "typeinfer").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	return arg0, nil
}
func expand_all(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var form vm.Value
	var head vm.Value
	var v15 vm.Value
	var v146 vm.Value
	var v217 vm.Value
	var arg__33985 vm.Value
	var arg__33998 vm.Value
	var arg__33999 vm.Value
	var arg__34014 vm.Value
	var arg__34027 vm.Value
	var arg__34028 vm.Value
	var arg__34029 vm.Value
	var v56 vm.Value
	var expanded vm.Value
	var v67 vm.Value
	var v139 vm.Value
	var arg__34040 vm.Value
	var v78 bool
	var v132 vm.Value
	var v134 vm.Value
	var arg__34056 vm.Value
	var arg__34069 vm.Value
	var arg__34070 vm.Value
	var arg__34077 vm.Value
	var arg__34088 vm.Value
	var arg__34101 vm.Value
	var arg__34102 vm.Value
	var arg__34103 vm.Value
	var v124 vm.Value
	var v126 vm.Value
	var arg__34122 vm.Value
	var arg__34137 vm.Value
	var arg__34138 vm.Value
	var v169 vm.Value
	var v174 vm.Value
	var v214 vm.Value
	var arg__34247 vm.Value
	var arg__34355 vm.Value
	var arg__34356 vm.Value
	var v199 vm.Value
	var v211 vm.Value
	var v208 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v4) {
		form = arg0
		goto b1
	} else {
		form = arg0
		goto b2
	}
b1:
	;
	head, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v15, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "build-known-heads").Deref(), head})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v15) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v146, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v146) {
		goto b13
	} else {
		goto b14
	}
b3:
	;
	return v217, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__33985, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__33985})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__33998, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__33999, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__33998})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "doall").Deref(), []vm.Value{arg__33999})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34014, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__34014})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34027, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34028, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__34027})
	if callErr != nil {
		return nil, callErr
	}
	arg__34029, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "doall").Deref(), []vm.Value{arg__34028})
	if callErr != nil {
		return nil, callErr
	}
	v56, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), head, arg__34029})
	if callErr != nil {
		return nil, callErr
	}
	v139 = v56
	goto b6
b5:
	;
	expanded, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-fully").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	v67, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq?").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v67) {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v217 = v139
	goto b3
b7:
	;
	arg__34040, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	v78 = arg__34040 == vm.Symbol("quote")
	if v78 {
		goto b10
	} else {
		goto b11
	}
b8:
	;
	v132, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	v134 = v132
	goto b9
b9:
	;
	v139 = v134
	goto b6
b10:
	;
	v126 = expanded
	goto b12
b11:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	arg__34056, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__34056})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	arg__34069, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	arg__34070, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__34069})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "doall").Deref(), []vm.Value{arg__34070})
	if callErr != nil {
		return nil, callErr
	}
	arg__34077, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	arg__34088, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__34088})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	arg__34101, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{expanded})
	if callErr != nil {
		return nil, callErr
	}
	arg__34102, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), arg__34101})
	if callErr != nil {
		return nil, callErr
	}
	arg__34103, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "doall").Deref(), []vm.Value{arg__34102})
	if callErr != nil {
		return nil, callErr
	}
	v124, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), arg__34077, arg__34103})
	if callErr != nil {
		return nil, callErr
	}
	v126 = v124
	goto b12
b12:
	;
	v134 = v126
	goto b9
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34122, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "doall").Deref(), []vm.Value{arg__34122})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34137, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34138, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "doall").Deref(), []vm.Value{arg__34137})
	if callErr != nil {
		return nil, callErr
	}
	v169, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__34138})
	if callErr != nil {
		return nil, callErr
	}
	v214 = v169
	goto b15
b14:
	;
	v174, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{form})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v174) {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v217 = v214
	goto b3
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var k vm.Value
		var v vm.Value
		var arg__34187 vm.Value
		var arg__34191 vm.Value
		var v18 vm.Value
		var callErr error
		k, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		arg__34187, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{k})
		if callErr != nil {
			return nil, callErr
		}
		arg__34191, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{v})
		if callErr != nil {
			return nil, callErr
		}
		v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__34187, arg__34191})
		if callErr != nil {
			return nil, callErr
		}
		return v18, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34247, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var k vm.Value
		var v vm.Value
		var arg__34240 vm.Value
		var arg__34244 vm.Value
		var v18 vm.Value
		var callErr error
		k, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		arg__34240, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{k})
		if callErr != nil {
			return nil, callErr
		}
		arg__34244, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{v})
		if callErr != nil {
			return nil, callErr
		}
		v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__34240, arg__34244})
		if callErr != nil {
			return nil, callErr
		}
		return v18, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "doall").Deref(), []vm.Value{arg__34247})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var k vm.Value
		var v vm.Value
		var arg__34295 vm.Value
		var arg__34299 vm.Value
		var v18 vm.Value
		var callErr error
		k, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		arg__34295, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{k})
		if callErr != nil {
			return nil, callErr
		}
		arg__34299, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{v})
		if callErr != nil {
			return nil, callErr
		}
		v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__34295, arg__34299})
		if callErr != nil {
			return nil, callErr
		}
		return v18, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34355, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var k vm.Value
		var v vm.Value
		var arg__34348 vm.Value
		var arg__34352 vm.Value
		var v18 vm.Value
		var callErr error
		k, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		v, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1), vm.NIL})
		if callErr != nil {
			return nil, callErr
		}
		arg__34348, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{k})
		if callErr != nil {
			return nil, callErr
		}
		arg__34352, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.pipeline", "expand-all").Deref(), []vm.Value{v})
		if callErr != nil {
			return nil, callErr
		}
		v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__34348, arg__34352})
		if callErr != nil {
			return nil, callErr
		}
		return v18, nil
	}), form})
	if callErr != nil {
		return nil, callErr
	}
	arg__34356, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "doall").Deref(), []vm.Value{arg__34355})
	if callErr != nil {
		return nil, callErr
	}
	v199, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{vm.EmptyPersistentMap, arg__34356})
	if callErr != nil {
		return nil, callErr
	}
	v211 = v199
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
	v214 = v211
	goto b15
b19:
	;
	v208 = form
	goto b21
b20:
	;
	v208 = vm.NIL
	goto b21
b21:
	;
	v211 = v208
	goto b18
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.pipeline", map[string]vm.Value{"unwrap-name": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("unwrap-name: wrong number of arguments %d (expected 1)", len(args))
		}
		return unwrap_name(args[0])
	}), "override-wrapper-fn-lit": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("override-wrapper-fn-lit: wrong number of arguments %d (expected 1)", len(args))
		}
		return override_wrapper_fn_lit(args[0])
	}), "override-init-decl": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("override-init-decl: wrong number of arguments %d (expected 2)", len(args))
		}
		return override_init_decl(args[0], args[1])
	}), "override-helper-decl": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 0 {
			return nil, fmt.Errorf("override-helper-decl: wrong number of arguments %d (expected 0)", len(args))
		}
		return override_helper_decl()
	}), "override-entries": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("override-entries: wrong number of arguments %d (expected 1)", len(args))
		}
		return override_entries(args[0])
	}), "collect-call-targets": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("collect-call-targets: wrong number of arguments %d (expected 1)", len(args))
		}
		return collect_call_targets(args[0])
	}), "defn-key": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("defn-key: wrong number of arguments %d (expected 1)", len(args))
		}
		return defn_key(args[0])
	}), "forms-by-arity": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("forms-by-arity: wrong number of arguments %d (expected 1)", len(args))
		}
		return forms_by_arity(args[0])
	}), "order-defn-forms": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("order-defn-forms: wrong number of arguments %d (expected 1)", len(args))
		}
		return order_defn_forms(args[0])
	}), "compile-form*": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("compile-form*: wrong number of arguments %d (expected 1)", len(args))
		}
		return compile_form_STAR_(args[0])
	}), "expand-fully": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expand-fully: wrong number of arguments %d (expected 1)", len(args))
		}
		return expand_fully(args[0])
	}), "go-name-for": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("go-name-for: wrong number of arguments %d (expected 1)", len(args))
		}
		return go_name_for(args[0])
	}), "optimize-fn": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("optimize-fn: wrong number of arguments %d (expected 1)", len(args))
		}
		return optimize_fn(args[0])
	}), "expand-all": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expand-all: wrong number of arguments %d (expected 1)", len(args))
		}
		return expand_all(args[0])
	}),
	})
}
