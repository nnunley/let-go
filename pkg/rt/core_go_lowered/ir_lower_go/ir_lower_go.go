package ir_lower_go

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func any_fn_template_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var or__x vm.Value
	var aux vm.Value
	var v10 vm.Value
	var v12 vm.Value
	var callErr error
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "fn-template?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b1
	} else {
		aux = arg0
		goto b2
	}
b1:
	;
	v12 = or__x
	goto b3
b2:
	;
	v10, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "multi-fn-template?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v12 = v10
	goto b3
b3:
	;
	return v12, nil
}
func arg_local_decls(arg0 vm.Value) (vm.Value, error) {
	var arity vm.Value
	var variadic_QMARK_ vm.Value
	var v13 vm.Value
	var fixed_arity vm.Value
	var i int
	var decls vm.Value
	var assigns vm.Value
	var v43 bool
	var sym vm.Value
	var arg_sym vm.Value
	var v161 int
	var arg__13443 vm.Value
	var arg__13459 vm.Value
	var arg__13461 vm.Value
	var v187 vm.Value
	var arg__13477 vm.Value
	var arg__13481 vm.Value
	var arg__13499 vm.Value
	var arg__13503 vm.Value
	var arg__13504 vm.Value
	var v213 vm.Value
	var v215 vm.Value
	var final_decls vm.Value
	var final_assigns vm.Value
	var arg__13353 vm.Value
	var arg__13369 vm.Value
	var arg__13371 vm.Value
	var arg__13372 vm.Value
	var arg__13388 vm.Value
	var arg__13392 vm.Value
	var arg__13410 vm.Value
	var arg__13414 vm.Value
	var arg__13415 vm.Value
	var arg__13416 vm.Value
	var v137 vm.Value
	var v140 vm.Value
	var v142 vm.Value
	var callErr error
	arity, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arity").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	variadic_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(variadic_QMARK_) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v13 = rt.SubValue(arity, vm.Int(1))
	fixed_arity = v13
	goto b3
b2:
	;
	fixed_arity = arity
	goto b3
b3:
	;
	i = 0
	decls = vm.NewArrayVector([]vm.Value{})
	assigns = vm.NewArrayVector([]vm.Value{})
	goto b4
b4:
	;
	v43 = rt.GeValue(vm.Int(i), fixed_arity)
	if v43 {
		goto b5
	} else {
		goto b6
	}
b5:
	;
	if vm.IsTruthy(variadic_QMARK_) {
		final_decls = decls
		final_assigns = assigns
		goto b8
	} else {
		final_decls = decls
		final_assigns = assigns
		goto b9
	}
b6:
	;
	sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("a"), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg_sym, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("arg"), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v161 = i + 1
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13443, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "var-decl").Deref(), []vm.Value{sym, arg__13443, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13459, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13461, callErr = rt.InvokeValue(rt.LookupVar("gogen", "var-decl").Deref(), []vm.Value{sym, arg__13459, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v187, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{decls, arg__13461})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{arg_sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__13477, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__13481, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{arg_sym})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__13477, arg__13481})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{arg_sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__13499, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__13503, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{arg_sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__13504, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__13499, arg__13503})
	if callErr != nil {
		return nil, callErr
	}
	v213, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{assigns, arg__13504})
	if callErr != nil {
		return nil, callErr
	}
	i = v161
	decls = v187
	assigns = v213
	goto b4
b7:
	;
	return v215, nil
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13353, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "var-decl").Deref(), []vm.Value{vm.String("args0"), arg__13353, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13369, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13371, callErr = rt.InvokeValue(rt.LookupVar("gogen", "var-decl").Deref(), []vm.Value{vm.String("args0"), arg__13369, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	arg__13372, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{final_decls, arg__13371})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args0")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13388, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args0")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13392, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__13388, arg__13392})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args0")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13410, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args0")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13414, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13415, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__13410, arg__13414})
	if callErr != nil {
		return nil, callErr
	}
	arg__13416, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{final_assigns, arg__13415})
	if callErr != nil {
		return nil, callErr
	}
	v137, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13372, arg__13416})
	if callErr != nil {
		return nil, callErr
	}
	v142 = v137
	goto b10
b9:
	;
	v140, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{final_decls, final_assigns})
	if callErr != nil {
		return nil, callErr
	}
	v142 = v140
	goto b10
b10:
	;
	v215 = v142
	goto b7
}
func block_arg_sources(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__13509 vm.Value
	var tem__G__0 vm.Value
	var nid vm.Value
	var v14 vm.Value
	var f vm.Value
	var arg__13524 vm.Value
	var v21 vm.Value
	var v23 vm.Value
	var callErr error
	arg__13509, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	tem__G__0, callErr = rt.InvokeValue(vm.Keyword("param-sources"), []vm.Value{arg__13509})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(tem__G__0) {
		nid = arg1
		goto b1
	} else {
		f = arg0
		nid = arg1
		goto b2
	}
b1:
	;
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{tem__G__0, nid})
	if callErr != nil {
		return nil, callErr
	}
	v23 = v14
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "compute-param-sources").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__13524, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "compute-param-sources").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v21, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__13524, nid})
	if callErr != nil {
		return nil, callErr
	}
	v23 = v21
	goto b3
b3:
	;
	return v23, nil
}
func box_as_value(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var t vm.Value
	var spec vm.Value
	var expr vm.Value
	var v22 vm.Value
	var v39 bool
	var v249 vm.Value
	var v55 bool
	var v241 vm.Value
	var arg__13557 vm.Value
	var arg__13573 vm.Value
	var arg__13575 vm.Value
	var arg__13578 vm.Value
	var v86 vm.Value
	var v101 bool
	var v233 vm.Value
	var arg__13589 vm.Value
	var v110 vm.Value
	var v125 bool
	var v225 vm.Value
	var arg__13600 vm.Value
	var v134 vm.Value
	var v149 bool
	var v217 vm.Value
	var arg__13611 vm.Value
	var v158 vm.Value
	var v173 bool
	var v209 vm.Value
	var v201 vm.Value
	var v193 vm.Value
	var callErr error
	t, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg2, arg0})
	if callErr != nil {
		return nil, callErr
	}
	spec, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	expr, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-expr").Deref(), []vm.Value{arg0, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v22) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v249 = vm.NIL
	goto b3
b2:
	;
	v39 = spec == vm.String("vm.Value")
	if v39 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v249, nil
b4:
	;
	v241 = expr
	goto b6
b5:
	;
	v55 = spec == vm.String("bool")
	if v55 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v249 = v241
	goto b3
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13557, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__13557, vm.String("Boolean")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13573, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13575, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__13573, vm.String("Boolean")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13578, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	v86, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__13575, arg__13578})
	if callErr != nil {
		return nil, callErr
	}
	v233 = v86
	goto b9
b8:
	;
	v101 = spec == vm.String("int")
	if v101 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v241 = v233
	goto b6
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	arg__13589, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	v110, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("Int"), arg__13589})
	if callErr != nil {
		return nil, callErr
	}
	v225 = v110
	goto b12
b11:
	;
	v125 = spec == vm.String("float64")
	if v125 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v233 = v225
	goto b9
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	arg__13600, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	v134, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("Float"), arg__13600})
	if callErr != nil {
		return nil, callErr
	}
	v217 = v134
	goto b15
b14:
	;
	v149 = spec == vm.String("string")
	if v149 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v225 = v217
	goto b12
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	arg__13611, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	v158, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("String"), arg__13611})
	if callErr != nil {
		return nil, callErr
	}
	v209 = v158
	goto b18
b17:
	;
	v173 = spec == vm.String("vm.Char")
	if v173 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v217 = v209
	goto b15
b19:
	;
	v201 = expr
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
	v209 = v201
	goto b18
b22:
	;
	v193 = expr
	goto b24
b23:
	;
	v193 = vm.NIL
	goto b24
b24:
	;
	v201 = v193
	goto b21
}
func boxed_list_expr(arg0 vm.Value) (vm.Value, error) {
	var arg__13622 vm.Value
	var v8 vm.Value
	var v13 vm.Value
	var xs vm.Value
	var arg__13636 vm.Value
	var elems vm.Value
	var v30 vm.Value
	var v67 vm.Value
	var arg__13653 vm.Value
	var arg__13655 vm.Value
	var arg__13669 vm.Value
	var arg__13671 vm.Value
	var arg__13672 vm.Value
	var v59 vm.Value
	var v63 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__13622, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{arg__13622})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v8) {
		goto b1
	} else {
		xs = arg0
		goto b2
	}
b1:
	;
	v13, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-sel").Deref(), []vm.Value{vm.String("EmptyList")})
	if callErr != nil {
		return nil, callErr
	}
	v67 = v13
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{xs})
	if callErr != nil {
		return nil, callErr
	}
	arg__13636, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{xs})
	if callErr != nil {
		return nil, callErr
	}
	elems, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "boxed-value-expr").Deref(), arg__13636})
	if callErr != nil {
		return nil, callErr
	}
	v30, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("clojure.core", "identity").Deref(), elems})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v30) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v67, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13653, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13655, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__13653, elems})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13655})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13669, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13671, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__13669, elems})
	if callErr != nil {
		return nil, callErr
	}
	arg__13672, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13671})
	if callErr != nil {
		return nil, callErr
	}
	v59, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("NewList"), arg__13672})
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
	v67 = v63
	goto b3
}
func boxed_map_expr(arg0 vm.Value) (vm.Value, error) {
	var arg__13681 vm.Value
	var v8 vm.Value
	var v13 vm.Value
	var m vm.Value
	var entries vm.Value
	var v164 vm.Value
	var remaining vm.Value
	var out vm.Value
	var v30 vm.Value
	var entry vm.Value
	var arg__13702 vm.Value
	var kexpr vm.Value
	var arg__13711 vm.Value
	var vexpr vm.Value
	var or__x vm.Value
	var elems vm.Value
	var v97 vm.Value
	var v99 vm.Value
	var v82 vm.Value
	var v84 vm.Value
	var arg__13739 vm.Value
	var arg__13741 vm.Value
	var arg__13755 vm.Value
	var arg__13757 vm.Value
	var arg__13758 vm.Value
	var v153 vm.Value
	var v157 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__13681, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v8, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{arg__13681})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v8) {
		goto b1
	} else {
		m = arg0
		goto b2
	}
b1:
	;
	v13, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-sel").Deref(), []vm.Value{vm.String("EmptyPersistentMap")})
	if callErr != nil {
		return nil, callErr
	}
	v164 = v13
	goto b3
b2:
	;
	entries, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{m})
	if callErr != nil {
		return nil, callErr
	}
	remaining = entries
	out = vm.NewArrayVector([]vm.Value{})
	goto b4
b3:
	;
	return v164, nil
b4:
	;
	v30, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v30) {
		goto b5
	} else {
		goto b6
	}
b5:
	;
	elems = out
	goto b7
b6:
	;
	entry, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{entry})
	if callErr != nil {
		return nil, callErr
	}
	arg__13702, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{entry})
	if callErr != nil {
		return nil, callErr
	}
	kexpr, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "boxed-value-expr").Deref(), []vm.Value{arg__13702})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{entry})
	if callErr != nil {
		return nil, callErr
	}
	arg__13711, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "second").Deref(), []vm.Value{entry})
	if callErr != nil {
		return nil, callErr
	}
	vexpr, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "boxed-value-expr").Deref(), []vm.Value{arg__13711})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{kexpr})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b11
	} else {
		goto b12
	}
b7:
	;
	if vm.IsTruthy(elems) {
		goto b14
	} else {
		goto b15
	}
b8:
	;
	goto b10
b9:
	;
	v97, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v99, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, kexpr, vexpr})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v97
	out = v99
	goto b4
b10:
	;
	elems = vm.NIL
	goto b7
b11:
	;
	v84 = or__x
	goto b13
b12:
	;
	v82, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{vexpr})
	if callErr != nil {
		return nil, callErr
	}
	v84 = v82
	goto b13
b13:
	;
	if vm.IsTruthy(v84) {
		goto b8
	} else {
		goto b9
	}
b14:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13739, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13741, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__13739, elems})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13741})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13755, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13757, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__13755, elems})
	if callErr != nil {
		return nil, callErr
	}
	arg__13758, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13757})
	if callErr != nil {
		return nil, callErr
	}
	v153, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("NewPersistentMap"), arg__13758})
	if callErr != nil {
		return nil, callErr
	}
	v157 = v153
	goto b16
b15:
	;
	v157 = vm.NIL
	goto b16
b16:
	;
	v164 = v157
	goto b3
}
func boxed_value_expr(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var v9 vm.Value
	var v vm.Value
	var v14 bool
	var v247 vm.Value
	var v19 vm.Value
	var v24 bool
	var v244 vm.Value
	var v29 vm.Value
	var v34 vm.Value
	var v241 vm.Value
	var arg__13783 vm.Value
	var arg__13791 vm.Value
	var arg__13792 vm.Value
	var v47 vm.Value
	var v52 vm.Value
	var v238 vm.Value
	var arg__13801 vm.Value
	var arg__13809 vm.Value
	var arg__13810 vm.Value
	var v65 vm.Value
	var v70 vm.Value
	var v235 vm.Value
	var arg__13819 vm.Value
	var arg__13827 vm.Value
	var arg__13828 vm.Value
	var v83 vm.Value
	var v88 vm.Value
	var v232 vm.Value
	var arg__13837 vm.Value
	var arg__13845 vm.Value
	var arg__13846 vm.Value
	var v101 vm.Value
	var v106 vm.Value
	var v229 vm.Value
	var arg__13861 vm.Value
	var arg__13874 vm.Value
	var arg__13876 vm.Value
	var arg__13877 vm.Value
	var arg__13891 vm.Value
	var arg__13904 vm.Value
	var arg__13906 vm.Value
	var arg__13907 vm.Value
	var arg__13908 vm.Value
	var v151 vm.Value
	var v156 vm.Value
	var v226 vm.Value
	var arg__13922 vm.Value
	var arg__13923 vm.Value
	var arg__13936 vm.Value
	var arg__13937 vm.Value
	var arg__13938 vm.Value
	var v177 vm.Value
	var v182 vm.Value
	var v223 vm.Value
	var v185 vm.Value
	var v190 vm.Value
	var v220 vm.Value
	var v193 vm.Value
	var v198 vm.Value
	var v217 vm.Value
	var v201 vm.Value
	var v214 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v4) {
		goto b1
	} else {
		v = arg0
		goto b2
	}
b1:
	;
	v9, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-sel").Deref(), []vm.Value{vm.String("NIL")})
	if callErr != nil {
		return nil, callErr
	}
	v247 = v9
	goto b3
b2:
	;
	v14 = v == vm.Boolean(true)
	if v14 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v247, nil
b4:
	;
	v19, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-sel").Deref(), []vm.Value{vm.String("TRUE")})
	if callErr != nil {
		return nil, callErr
	}
	v244 = v19
	goto b6
b5:
	;
	v24 = v == vm.Boolean(false)
	if v24 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v247 = v244
	goto b3
b7:
	;
	v29, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-sel").Deref(), []vm.Value{vm.String("FALSE")})
	if callErr != nil {
		return nil, callErr
	}
	v241 = v29
	goto b9
b8:
	;
	v34, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "integer?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v34) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v244 = v241
	goto b6
b10:
	;
	arg__13783, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13783})
	if callErr != nil {
		return nil, callErr
	}
	arg__13791, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13792, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13791})
	if callErr != nil {
		return nil, callErr
	}
	v47, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("Int"), arg__13792})
	if callErr != nil {
		return nil, callErr
	}
	v238 = v47
	goto b12
b11:
	;
	v52, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "float?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v52) {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v241 = v238
	goto b9
b13:
	;
	arg__13801, callErr = rt.InvokeValue(rt.LookupVar("gogen", "float-lit").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13801})
	if callErr != nil {
		return nil, callErr
	}
	arg__13809, callErr = rt.InvokeValue(rt.LookupVar("gogen", "float-lit").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13810, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13809})
	if callErr != nil {
		return nil, callErr
	}
	v65, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("Float"), arg__13810})
	if callErr != nil {
		return nil, callErr
	}
	v235 = v65
	goto b15
b14:
	;
	v70, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v70) {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v238 = v235
	goto b12
b16:
	;
	arg__13819, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13819})
	if callErr != nil {
		return nil, callErr
	}
	arg__13827, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13828, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13827})
	if callErr != nil {
		return nil, callErr
	}
	v83, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("String"), arg__13828})
	if callErr != nil {
		return nil, callErr
	}
	v232 = v83
	goto b18
b17:
	;
	v88, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "char?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v88) {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v235 = v232
	goto b15
b19:
	;
	arg__13837, callErr = rt.InvokeValue(rt.LookupVar("gogen", "char-lit").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13837})
	if callErr != nil {
		return nil, callErr
	}
	arg__13845, callErr = rt.InvokeValue(rt.LookupVar("gogen", "char-lit").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13846, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13845})
	if callErr != nil {
		return nil, callErr
	}
	v101, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("Char"), arg__13846})
	if callErr != nil {
		return nil, callErr
	}
	v229 = v101
	goto b21
b20:
	;
	v106, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v106) {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v232 = v229
	goto b18
b22:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13861, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subs").Deref(), []vm.Value{arg__13861, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13874, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13876, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subs").Deref(), []vm.Value{arg__13874, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__13877, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{arg__13876})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13877})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13891, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subs").Deref(), []vm.Value{arg__13891, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13904, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13906, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subs").Deref(), []vm.Value{arg__13904, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__13907, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{arg__13906})
	if callErr != nil {
		return nil, callErr
	}
	arg__13908, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13907})
	if callErr != nil {
		return nil, callErr
	}
	v151, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("Keyword"), arg__13908})
	if callErr != nil {
		return nil, callErr
	}
	v226 = v151
	goto b24
b23:
	;
	v156, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "symbol?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v156) {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v229 = v226
	goto b21
b25:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13922, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13923, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{arg__13922})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13923})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13936, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	arg__13937, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{arg__13936})
	if callErr != nil {
		return nil, callErr
	}
	arg__13938, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13937})
	if callErr != nil {
		return nil, callErr
	}
	v177, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("Symbol"), arg__13938})
	if callErr != nil {
		return nil, callErr
	}
	v223 = v177
	goto b27
b26:
	;
	v182, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v182) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v226 = v223
	goto b24
b28:
	;
	v185, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "boxed-list-expr").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	v220 = v185
	goto b30
b29:
	;
	v190, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v190) {
		goto b31
	} else {
		goto b32
	}
b30:
	;
	v223 = v220
	goto b27
b31:
	;
	v193, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "boxed-vector-expr").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	v217 = v193
	goto b33
b32:
	;
	v198, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v198) {
		goto b34
	} else {
		goto b35
	}
b33:
	;
	v220 = v217
	goto b30
b34:
	;
	v201, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "boxed-map-expr").Deref(), []vm.Value{v})
	if callErr != nil {
		return nil, callErr
	}
	v214 = v201
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
	v217 = v214
	goto b33
b37:
	;
	goto b39
b38:
	;
	goto b39
b39:
	;
	v214 = vm.NIL
	goto b36
}
func boxed_vector_expr(arg0 vm.Value) (vm.Value, error) {
	var elems vm.Value
	var v12 vm.Value
	var arg__13978 vm.Value
	var arg__13980 vm.Value
	var arg__13994 vm.Value
	var arg__13996 vm.Value
	var arg__13997 vm.Value
	var v41 vm.Value
	var v45 vm.Value
	var callErr error
	elems, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "boxed-value-expr").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("clojure.core", "identity").Deref(), elems})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13978, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13980, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__13978, elems})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13980})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13994, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__13996, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__13994, elems})
	if callErr != nil {
		return nil, callErr
	}
	arg__13997, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__13996})
	if callErr != nil {
		return nil, callErr
	}
	v41, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-call").Deref(), []vm.Value{vm.String("NewArrayVector"), arg__13997})
	if callErr != nil {
		return nil, callErr
	}
	v45 = v41
	goto b3
b2:
	;
	v45 = vm.NIL
	goto b3
b3:
	;
	return v45, nil
}
func call_assign_stmts(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var refs vm.Value
	var arg__14016 vm.Value
	var callee vm.Value
	var arg__14041 vm.Value
	var arg_exprs vm.Value
	var arg__14051 vm.Value
	var args_slice vm.Value
	var arg__14062 vm.Value
	var arg__14079 vm.Value
	var arg__14081 vm.Value
	var arg__14085 vm.Value
	var invoke_expr vm.Value
	var err_id vm.Value
	var f vm.Value
	var nid vm.Value
	var v145 vm.Value
	var v463 vm.Value
	var v111 vm.Value
	var and__x vm.Value
	var v114 vm.Value
	var arg__14106 vm.Value
	var arg__14120 vm.Value
	var arg__14122 vm.Value
	var arg__14125 vm.Value
	var v162 vm.Value
	var arg__14131 vm.Value
	var arg__14143 vm.Value
	var arg__14145 vm.Value
	var arg__14148 vm.Value
	var v183 vm.Value
	var assign vm.Value
	var case__13998 vm.Value
	var v220 bool
	var v225 vm.Value
	var v250 bool
	var zero_expr vm.Value
	var arg__14188 vm.Value
	var arg__14199 vm.Value
	var arg__14200 vm.Value
	var arg__14217 vm.Value
	var arg__14218 vm.Value
	var arg__14228 vm.Value
	var arg__14229 vm.Value
	var arg__14230 vm.Value
	var err_check vm.Value
	var v459 vm.Value
	var v255 vm.Value
	var v280 bool
	var v387 vm.Value
	var v285 vm.Value
	var v310 bool
	var v374 vm.Value
	var v315 vm.Value
	var v361 vm.Value
	var v344 vm.Value
	var v348 vm.Value
	var callErr error
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg2, arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	arg__14016, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	callee, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "box-as-value").Deref(), []vm.Value{arg0, arg1, arg__14016})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	arg__14041, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	arg_exprs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "box-as-value").Deref(), []vm.Value{arg0, arg1, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), arg__14041})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14051, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	args_slice, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__14051, arg_exprs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14062, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__14062, vm.String("InvokeValue")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{callee, args_slice})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14079, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14081, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__14079, vm.String("InvokeValue")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14085, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{callee, args_slice})
	if callErr != nil {
		return nil, callErr
	}
	invoke_expr, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__14081, arg__14085})
	if callErr != nil {
		return nil, callErr
	}
	err_id, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("callErr")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(callee) {
		f = arg0
		nid = arg2
		goto b4
	} else {
		f = arg0
		nid = arg2
		and__x = callee
		goto b5
	}
b1:
	;
	v145, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v145) {
		goto b7
	} else {
		goto b8
	}
b2:
	;
	v463 = vm.NIL
	goto b3
b3:
	;
	return v463, nil
b4:
	;
	v111, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("clojure.core", "identity").Deref(), arg_exprs})
	if callErr != nil {
		return nil, callErr
	}
	v114 = v111
	goto b6
b5:
	;
	v114 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v114) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	arg__14106, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-ident").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__14106, err_id})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{invoke_expr})
	if callErr != nil {
		return nil, callErr
	}
	arg__14120, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-ident").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__14122, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__14120, err_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__14125, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{invoke_expr})
	if callErr != nil {
		return nil, callErr
	}
	v162, callErr = rt.InvokeValue(rt.LookupVar("gogen", "multi-assign").Deref(), []vm.Value{vm.String("="), arg__14122, arg__14125})
	if callErr != nil {
		return nil, callErr
	}
	assign = v162
	goto b9
b8:
	;
	arg__14131, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("_")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__14131, err_id})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{invoke_expr})
	if callErr != nil {
		return nil, callErr
	}
	arg__14143, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("_")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14145, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__14143, err_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__14148, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{invoke_expr})
	if callErr != nil {
		return nil, callErr
	}
	v183, callErr = rt.InvokeValue(rt.LookupVar("gogen", "multi-assign").Deref(), []vm.Value{vm.String("="), arg__14145, arg__14148})
	if callErr != nil {
		return nil, callErr
	}
	assign = v183
	goto b9
b9:
	;
	case__13998, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-return-spec").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v220 = case__13998 == vm.String("bool")
	if v220 {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	v225, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("false")})
	if callErr != nil {
		return nil, callErr
	}
	zero_expr = v225
	goto b12
b11:
	;
	v250 = case__13998 == vm.String("int")
	if v250 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14188, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "binary").Deref(), []vm.Value{vm.String("!="), err_id, arg__14188})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{zero_expr, err_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__14199, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{zero_expr, err_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__14200, callErr = rt.InvokeValue(rt.LookupVar("gogen", "return-stmt").Deref(), []vm.Value{arg__14199})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__14200})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14217, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	arg__14218, callErr = rt.InvokeValue(rt.LookupVar("gogen", "binary").Deref(), []vm.Value{vm.String("!="), err_id, arg__14217})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{zero_expr, err_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__14228, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{zero_expr, err_id})
	if callErr != nil {
		return nil, callErr
	}
	arg__14229, callErr = rt.InvokeValue(rt.LookupVar("gogen", "return-stmt").Deref(), []vm.Value{arg__14228})
	if callErr != nil {
		return nil, callErr
	}
	arg__14230, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__14229})
	if callErr != nil {
		return nil, callErr
	}
	err_check, callErr = rt.InvokeValue(rt.LookupVar("gogen", "if-stmt").Deref(), []vm.Value{vm.NIL, arg__14218, arg__14230, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v459, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{assign, err_check})
	if callErr != nil {
		return nil, callErr
	}
	v463 = v459
	goto b3
b13:
	;
	v255, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v387 = v255
	goto b15
b14:
	;
	v280 = case__13998 == vm.String("float64")
	if v280 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	zero_expr = v387
	goto b12
b16:
	;
	v285, callErr = rt.InvokeValue(rt.LookupVar("gogen", "float-lit").Deref(), []vm.Value{vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v374 = v285
	goto b18
b17:
	;
	v310 = case__13998 == vm.String("string")
	if v310 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v387 = v374
	goto b15
b19:
	;
	v315, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{vm.String("")})
	if callErr != nil {
		return nil, callErr
	}
	v361 = v315
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
	v374 = v361
	goto b18
b22:
	;
	v344, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	v348 = v344
	goto b24
b23:
	;
	v348 = vm.NIL
	goto b24
b24:
	;
	v361 = v348
	goto b21
}
func closure_expr(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var info vm.Value
	var f vm.Value
	var closed_exprs vm.Value
	var arg__14262 vm.Value
	var capture_exprs vm.Value
	var v41 vm.Value
	var v61 vm.Value
	var arg__14275 vm.Value
	var v48 vm.Value
	var v52 vm.Value
	var callErr error
	info, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-info").Deref(), []vm.Value{arg0, arg2})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(info) {
		f = arg0
		closed_exprs = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("captures"), []vm.Value{info})
	if callErr != nil {
		return nil, callErr
	}
	arg__14262, callErr = rt.InvokeValue(vm.Keyword("captures"), []vm.Value{info})
	if callErr != nil {
		return nil, callErr
	}
	capture_exprs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "box-as-value").Deref(), []vm.Value{f, closed_exprs, arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	}), arg__14262})
	if callErr != nil {
		return nil, callErr
	}
	v41, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("clojure.core", "identity").Deref(), capture_exprs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v41) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v61 = vm.NIL
	goto b3
b3:
	;
	return v61, nil
b4:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("template"), []vm.Value{info})
	if callErr != nil {
		return nil, callErr
	}
	arg__14275, callErr = rt.InvokeValue(vm.Keyword("template"), []vm.Value{info})
	if callErr != nil {
		return nil, callErr
	}
	v48, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower-template-closure-expr").Deref(), []vm.Value{arg__14275, capture_exprs})
	if callErr != nil {
		return nil, callErr
	}
	v52 = v48
	goto b6
b5:
	;
	v52 = vm.NIL
	goto b6
b6:
	;
	v61 = v52
	goto b3
}
func closure_info(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__14281 vm.Value
	var tem__G__0 vm.Value
	var f vm.Value
	var nid vm.Value
	var arg__14291 vm.Value
	var v26 vm.Value
	var arg__14330 vm.Value
	var v57 vm.Value
	var v59 vm.Value
	var memo vm.Value
	var arg__14302 vm.Value
	var v33 vm.Value
	var arg__14312 vm.Value
	var r vm.Value
	var v46 vm.Value
	var callErr error
	arg__14281, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	tem__G__0, callErr = rt.InvokeValue(vm.Keyword("closure-memo"), []vm.Value{arg__14281})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(tem__G__0) {
		f = arg0
		nid = arg1
		goto b1
	} else {
		f = arg0
		nid = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{tem__G__0})
	if callErr != nil {
		return nil, callErr
	}
	arg__14291, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{tem__G__0})
	if callErr != nil {
		return nil, callErr
	}
	v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg__14291, nid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v26) {
		memo = tem__G__0
		goto b4
	} else {
		memo = tem__G__0
		goto b5
	}
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__14330, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	v57, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-info*").Deref(), []vm.Value{f, nid, arg__14330})
	if callErr != nil {
		return nil, callErr
	}
	v59 = v57
	goto b3
b3:
	;
	return v59, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{memo})
	if callErr != nil {
		return nil, callErr
	}
	arg__14302, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{memo})
	if callErr != nil {
		return nil, callErr
	}
	v33, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__14302, nid})
	if callErr != nil {
		return nil, callErr
	}
	v46 = v33
	goto b6
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__14312, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	r, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-info*").Deref(), []vm.Value{f, nid, arg__14312})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{memo, rt.LookupVar("clojure.core", "assoc").Deref(), nid, r})
	if callErr != nil {
		return nil, callErr
	}
	v46 = r
	goto b6
b6:
	;
	v59 = v46
	goto b3
}
func closure_info_STAR_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var v10 vm.Value
	var f vm.Value
	var nid vm.Value
	var visited vm.Value
	var op vm.Value
	var aux vm.Value
	var refs vm.Value
	var and__x_33 bool
	var v322 vm.Value
	var v66 vm.Value
	var v81 bool
	var v314 vm.Value
	var v50 vm.Value
	var and__x_47 bool
	var v53 vm.Value
	var arg__14367 vm.Value
	var v98 bool
	var v134 bool
	var v306 vm.Value
	var arg__14383 vm.Value
	var v109 vm.Value
	var v113 vm.Value
	var arg__14391 vm.Value
	var v151 bool
	var v245 bool
	var v298 vm.Value
	var arg__14407 vm.Value
	var base vm.Value
	var v224 vm.Value
	var arg__14423 vm.Value
	var arg__14429 vm.Value
	var arg__14446 vm.Value
	var arg__14452 vm.Value
	var arg__14453 vm.Value
	var v209 vm.Value
	var v213 vm.Value
	var arg__14504 vm.Value
	var v262 vm.Value
	var v290 vm.Value
	var callErr error
	v10, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg2, arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v10) {
		goto b1
	} else {
		f = arg0
		nid = arg1
		visited = arg2
		goto b2
	}
b1:
	;
	v322 = vm.NIL
	goto b3
b2:
	;
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	and__x_33 = op == vm.Keyword("const")
	if and__x_33 {
		goto b7
	} else {
		and__x_47 = and__x_33
		goto b8
	}
b3:
	;
	return v322, nil
b4:
	;
	v66, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("template"), aux, vm.Keyword("captures"), vm.NewArrayVector([]vm.Value{})})
	if callErr != nil {
		return nil, callErr
	}
	v314 = v66
	goto b6
b5:
	;
	v81 = op == vm.Keyword("make-closure")
	if v81 {
		goto b10
	} else {
		goto b11
	}
b6:
	;
	v322 = v314
	goto b3
b7:
	;
	v50, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "any-fn-template?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v53 = v50
	goto b9
b8:
	;
	v53 = vm.Boolean(and__x_47)
	goto b9
b9:
	;
	if vm.IsTruthy(v53) {
		goto b4
	} else {
		goto b5
	}
b10:
	;
	arg__14367, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v98 = arg__14367 == vm.Int(1)
	if v98 {
		goto b13
	} else {
		goto b14
	}
b11:
	;
	v134 = op == vm.Keyword("push-closed")
	if v134 {
		goto b16
	} else {
		goto b17
	}
b12:
	;
	v314 = v306
	goto b6
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__14383, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v109, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-info*").Deref(), []vm.Value{f, arg__14383, visited})
	if callErr != nil {
		return nil, callErr
	}
	v113 = v109
	goto b15
b14:
	;
	v113 = vm.NIL
	goto b15
b15:
	;
	v306 = v113
	goto b12
b16:
	;
	arg__14391, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v151 = arg__14391 == vm.Int(2)
	if v151 {
		goto b19
	} else {
		goto b20
	}
b17:
	;
	v245 = op == vm.Keyword("block-arg")
	if v245 {
		goto b25
	} else {
		goto b26
	}
b18:
	;
	v306 = v298
	goto b12
b19:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__14407, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	base, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-info*").Deref(), []vm.Value{f, arg__14407, visited})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(base) {
		goto b22
	} else {
		goto b23
	}
b20:
	;
	v224 = vm.NIL
	goto b21
b21:
	;
	v298 = v224
	goto b18
b22:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("captures"), []vm.Value{base})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__14423, callErr = rt.InvokeValue(vm.Keyword("captures"), []vm.Value{base})
	if callErr != nil {
		return nil, callErr
	}
	arg__14429, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__14423, arg__14429})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("captures"), []vm.Value{base})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__14446, callErr = rt.InvokeValue(vm.Keyword("captures"), []vm.Value{base})
	if callErr != nil {
		return nil, callErr
	}
	arg__14452, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__14453, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__14446, arg__14452})
	if callErr != nil {
		return nil, callErr
	}
	v209, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{base, vm.Keyword("captures"), arg__14453})
	if callErr != nil {
		return nil, callErr
	}
	v213 = v209
	goto b24
b23:
	;
	v213 = vm.NIL
	goto b24
b24:
	;
	v224 = v213
	goto b21
b25:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "block-arg-sources").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__14504, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "block-arg-sources").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	v262, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__14497 vm.Value
		var v9 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{visited, nid})
		if callErr != nil {
			return nil, callErr
		}
		arg__14497, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{visited, nid})
		if callErr != nil {
			return nil, callErr
		}
		v9, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-info*").Deref(), []vm.Value{f, arg0, arg__14497})
		if callErr != nil {
			return nil, callErr
		}
		return v9, nil
	}), arg__14504})
	if callErr != nil {
		return nil, callErr
	}
	v290 = v262
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
	v298 = v290
	goto b18
b28:
	;
	goto b30
b29:
	;
	goto b30
b30:
	;
	v290 = vm.NIL
	goto b27
}
func closure_value_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__14517 vm.Value
	var v7 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-info").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__14517, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-info").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some?").Deref(), []vm.Value{arg__14517})
	if callErr != nil {
		return nil, callErr
	}
	return v7, nil
}
func coalesce_map(arg0 vm.Value) (vm.Value, error) {
	var arg__14530 vm.Value
	var arg__14534 vm.Value
	var unsafe vm.Value
	var nids vm.Value
	var name_types vm.Value
	var v36 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-unsafe-names").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-captured-names").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__14530, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-unsafe-names").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__14534, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-captured-names").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	unsafe, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{arg__14530, arg__14534})
	if callErr != nil {
		return nil, callErr
	}
	nids, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-all-nids").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	name_types, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var s vm.Value
		var and__x vm.Value
		var m vm.Value
		var nid vm.Value
		var f_7 vm.Value
		var arg__14691 vm.Value
		var arg__14705 vm.Value
		var arg__14716 vm.Value
		var arg__14717 vm.Value
		var arg__14730 vm.Value
		var arg__14731 vm.Value
		var arg__14744 vm.Value
		var arg__14758 vm.Value
		var arg__14769 vm.Value
		var arg__14770 vm.Value
		var arg__14783 vm.Value
		var arg__14784 vm.Value
		var arg__14785 vm.Value
		var v94 vm.Value
		var v97 vm.Value
		var f_17 vm.Value
		var arg__14680 vm.Value
		var v31 vm.Value
		var f_22 vm.Value
		var v34 vm.Value
		var f_37 vm.Value
		var callErr error
		s, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "source-name-of").Deref(), []vm.Value{arg0, arg1})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			m = arg0
			nid = arg1
			f_17 = arg0
			goto b4
		} else {
			m = arg0
			nid = arg1
			f_22 = arg0
			goto b5
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__14691, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{m, s, arg__14691})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_7})
		if callErr != nil {
			return nil, callErr
		}
		arg__14705, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_7})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__14705})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__14716, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__14717, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{m, s, arg__14716})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_7})
		if callErr != nil {
			return nil, callErr
		}
		arg__14730, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_7})
		if callErr != nil {
			return nil, callErr
		}
		arg__14731, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__14730})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__14717, arg__14731})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__14744, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{m, s, arg__14744})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_7})
		if callErr != nil {
			return nil, callErr
		}
		arg__14758, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_7})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__14758})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__14769, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
		if callErr != nil {
			return nil, callErr
		}
		arg__14770, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{m, s, arg__14769})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_7})
		if callErr != nil {
			return nil, callErr
		}
		arg__14783, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_7})
		if callErr != nil {
			return nil, callErr
		}
		arg__14784, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__14783})
		if callErr != nil {
			return nil, callErr
		}
		arg__14785, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__14770, arg__14784})
		if callErr != nil {
			return nil, callErr
		}
		v94, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{m, s, arg__14785})
		if callErr != nil {
			return nil, callErr
		}
		v97 = v94
		goto b3
	b2:
		;
		v97 = m
		goto b3
	b3:
		;
		return v97, nil
	b4:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		arg__14680, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		v31, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__14680})
		if callErr != nil {
			return nil, callErr
		}
		v34 = v31
		f_37 = f_17
		goto b6
	b5:
		;
		v34 = and__x
		f_37 = f_22
		goto b6
	b6:
		;
		if vm.IsTruthy(v34) {
			f_7 = f_37
			goto b1
		} else {
			goto b2
		}
	}), vm.EmptyPersistentMap, nids})
	if callErr != nil {
		return nil, callErr
	}
	v36, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var s vm.Value
		var and__x vm.Value
		var result vm.Value
		var nid vm.Value
		var gn vm.Value
		var arg__14918 vm.Value
		var v186 vm.Value
		var name_types_24 vm.Value
		var unsafe_25 vm.Value
		var arg__14878 vm.Value
		var v108 vm.Value
		var name_types_45 vm.Value
		var unsafe_46 vm.Value
		var arg__14891 vm.Value
		var v98 vm.Value
		var name_types_66 vm.Value
		var arg__14905 vm.Value
		var arg__14906 vm.Value
		var v85 vm.Value
		var v88 vm.Value
		var v173 vm.Value
		var v176 vm.Value
		var v159 vm.Value
		var v162 vm.Value
		var callErr error
		s, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "source-name-of").Deref(), []vm.Value{arg0, arg1})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			result = arg0
			nid = arg1
			name_types_24 = name_types
			unsafe_25 = unsafe
			goto b4
		} else {
			result = arg0
			nid = arg1
			goto b5
		}
	b1:
		;
		gn, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-name").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{gn})
		if callErr != nil {
			return nil, callErr
		}
		arg__14918, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{gn})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__14918})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			goto b16
		} else {
			goto b17
		}
	b2:
		;
		v186 = result
		goto b3
	b3:
		;
		return v186, nil
	b4:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		arg__14878, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{s})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__14878})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			name_types_45 = name_types_24
			unsafe_46 = unsafe_25
			goto b7
		} else {
			goto b8
		}
	b5:
		;
		v108 = and__x
		goto b6
	b6:
		;
		if vm.IsTruthy(v108) {
			goto b1
		} else {
			goto b2
		}
	b7:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{unsafe_46, s})
		if callErr != nil {
			return nil, callErr
		}
		arg__14891, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{unsafe_46, s})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__14891})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			name_types_66 = name_types_45
			goto b10
		} else {
			goto b11
		}
	b8:
		;
		v98 = and__x
		goto b9
	b9:
		;
		v108 = v98
		goto b6
	b10:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{name_types_66, s})
		if callErr != nil {
			return nil, callErr
		}
		arg__14905, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{name_types_66, s})
		if callErr != nil {
			return nil, callErr
		}
		arg__14906, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__14905})
		if callErr != nil {
			return nil, callErr
		}
		v85 = vm.Boolean(vm.Int(1) == arg__14906)
		v88 = v85
		goto b12
	b11:
		;
		v88 = and__x
		goto b12
	b12:
		;
		v98 = v88
		goto b9
	b13:
		;
		v173, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{result, nid, gn})
		if callErr != nil {
			return nil, callErr
		}
		v176 = v173
		goto b15
	b14:
		;
		v176 = result
		goto b15
	b15:
		;
		v186 = v176
		goto b3
	b16:
		;
		v159, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.String("_"), gn})
		if callErr != nil {
			return nil, callErr
		}
		v162 = v159
		goto b18
	b17:
		;
		v162 = and__x
		goto b18
	b18:
		;
		if vm.IsTruthy(v162) {
			goto b13
		} else {
			goto b14
		}
	}), vm.EmptyPersistentMap, nids})
	if callErr != nil {
		return nil, callErr
	}
	return v36, nil
}
func collect_local_ids(arg0 vm.Value) (vm.Value, error) {
	var block_ids vm.Value
	var bs vm.Value
	var out vm.Value
	var f vm.Value
	var v17 vm.Value
	var bid vm.Value
	var arg__15045 vm.Value
	var params vm.Value
	var arg__15148 vm.Value
	var insts vm.Value
	var v47 vm.Value
	var arg__15166 vm.Value
	var v53 vm.Value
	var v55 vm.Value
	var callErr error
	block_ids, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	bs = block_ids
	out = vm.NewArrayVector([]vm.Value{})
	f = arg0
	goto b1
b1:
	;
	v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v17) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v55 = out
	goto b4
b3:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__15045, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	params, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var and__x vm.Value
		var nid vm.Value
		var f_5 vm.Value
		var arg__15012 vm.Value
		var v57 vm.Value
		var f_18 vm.Value
		var arg__15025 vm.Value
		var v51 vm.Value
		var f_31 vm.Value
		var arg__15038 vm.Value
		var v42 vm.Value
		var v45 vm.Value
		var callErr error
		and__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f, arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			nid = arg0
			f_5 = f
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "const-param-of").Deref(), []vm.Value{f_5, nid})
		if callErr != nil {
			return nil, callErr
		}
		arg__15012, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "const-param-of").Deref(), []vm.Value{f_5, nid})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__15012})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_18 = f_5
			goto b4
		} else {
			goto b5
		}
	b2:
		;
		v57 = and__x
		goto b3
	b3:
		;
		return v57, nil
	b4:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-value?").Deref(), []vm.Value{f_18, nid})
		if callErr != nil {
			return nil, callErr
		}
		arg__15025, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-value?").Deref(), []vm.Value{f_18, nid})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__15025})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_31 = f_18
			goto b7
		} else {
			goto b8
		}
	b5:
		;
		v51 = and__x
		goto b6
	b6:
		;
		v57 = v51
		goto b3
	b7:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_31})
		if callErr != nil {
			return nil, callErr
		}
		arg__15038, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_31})
		if callErr != nil {
			return nil, callErr
		}
		v42, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "infer-go-type").Deref(), []vm.Value{arg__15038})
		if callErr != nil {
			return nil, callErr
		}
		v45 = v42
		goto b9
	b8:
		;
		v45 = and__x
		goto b9
	b9:
		;
		v51 = v45
		goto b6
	}), arg__15045})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__15148, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var op vm.Value
		var and__x vm.Value
		var nid vm.Value
		var f_7 vm.Value
		var v83 vm.Value
		var f_20 vm.Value
		var v76 vm.Value
		var f_31 vm.Value
		var arg__15128 vm.Value
		var v69 vm.Value
		var f_46 vm.Value
		var arg__15141 vm.Value
		var v59 vm.Value
		var v62 vm.Value
		var callErr error
		op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "local-carrying-op?").Deref(), []vm.Value{op})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			nid = arg0
			f_7 = f
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{op, vm.Keyword("block-arg")})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_20 = f_7
			goto b4
		} else {
			goto b5
		}
	b2:
		;
		v83 = and__x
		goto b3
	b3:
		;
		return v83, nil
	b4:
		;
		and__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f_20, nid})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_31 = f_20
			goto b7
		} else {
			goto b8
		}
	b5:
		;
		v76 = and__x
		goto b6
	b6:
		;
		v83 = v76
		goto b3
	b7:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-value?").Deref(), []vm.Value{f_31, nid})
		if callErr != nil {
			return nil, callErr
		}
		arg__15128, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-value?").Deref(), []vm.Value{f_31, nid})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__15128})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_46 = f_31
			goto b10
		} else {
			goto b11
		}
	b8:
		;
		v69 = and__x
		goto b9
	b9:
		;
		v76 = v69
		goto b6
	b10:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_46})
		if callErr != nil {
			return nil, callErr
		}
		arg__15141, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f_46})
		if callErr != nil {
			return nil, callErr
		}
		v59, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "infer-go-type").Deref(), []vm.Value{arg__15141})
		if callErr != nil {
			return nil, callErr
		}
		v62 = v59
		goto b12
	b11:
		;
		v62 = and__x
		goto b12
	b12:
		;
		v69 = v62
		goto b9
	}), arg__15148})
	if callErr != nil {
		return nil, callErr
	}
	v47, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{params, insts})
	if callErr != nil {
		return nil, callErr
	}
	arg__15166, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{params, insts})
	if callErr != nil {
		return nil, callErr
	}
	v53, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{out, arg__15166})
	if callErr != nil {
		return nil, callErr
	}
	bs = v47
	out = v53
	goto b1
b4:
	;
	return v55, nil
}
func compute_param_sources(arg0 vm.Value) (vm.Value, error) {
	var m vm.Value
	var arg__15184 vm.Value
	var doseq_seq__15167 vm.Value
	var doseq_loop__15168 vm.Value
	var f vm.Value
	var bid vm.Value
	var params vm.Value
	var np vm.Value
	var arg__15208 vm.Value
	var doseq_seq__15169 vm.Value
	var v457 vm.Value
	var doseq_loop__15170 vm.Value
	var p vm.Value
	var term vm.Value
	var v447 vm.Value
	var op vm.Value
	var aux vm.Value
	var arg__15239 vm.Value
	var doseq_seq__15171 vm.Value
	var v432 vm.Value
	var doseq_loop__15172 vm.Value
	var bt vm.Value
	var args vm.Value
	var v396 vm.Value
	var arg__15246 vm.Value
	var v207 bool
	var and__x vm.Value
	var v210 vm.Value
	var i int
	var v274 bool
	var arg__15257 vm.Value
	var v314 bool
	var pid vm.Value
	var a vm.Value
	var v352 int
	var callErr error
	m, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__15184, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__15167, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__15184})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15168 = doseq_seq__15167
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__15168) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__15168})
	if callErr != nil {
		return nil, callErr
	}
	params, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	np, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__15208, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__15169, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__15208})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15170 = doseq_seq__15169
	goto b5
b3:
	;
	goto b4
b4:
	;
	v457, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{m})
	if callErr != nil {
		return nil, callErr
	}
	return v457, nil
b5:
	;
	if vm.IsTruthy(doseq_loop__15170) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__15170})
	if callErr != nil {
		return nil, callErr
	}
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{p, f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(term) {
		goto b9
	} else {
		goto b10
	}
b7:
	;
	goto b8
b8:
	;
	v447, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__15168})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15168 = v447
	goto b1
b9:
	;
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "term-targets").Deref(), []vm.Value{op, aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__15239, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "term-targets").Deref(), []vm.Value{op, aux})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__15171, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__15239})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15172 = doseq_seq__15171
	goto b12
b10:
	;
	goto b11
b11:
	;
	v432, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__15170})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15170 = v432
	goto b5
b12:
	;
	if vm.IsTruthy(doseq_loop__15172) {
		goto b13
	} else {
		goto b14
	}
b13:
	;
	bt, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__15172})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(bt) {
		goto b19
	} else {
		and__x = bt
		goto b20
	}
b14:
	;
	goto b15
b15:
	;
	goto b11
b16:
	;
	args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	goto b22
b17:
	;
	goto b18
b18:
	;
	v396, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__15172})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15172 = v396
	goto b12
b19:
	;
	arg__15246, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	v207 = arg__15246 == bid
	v210 = vm.Boolean(v207)
	goto b21
b20:
	;
	v210 = and__x
	goto b21
b21:
	;
	if vm.IsTruthy(v210) {
		goto b16
	} else {
		goto b17
	}
b22:
	;
	v274 = rt.LtValue(vm.Int(i), np)
	if v274 {
		goto b23
	} else {
		goto b24
	}
b23:
	;
	arg__15257, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	v314 = rt.LtValue(vm.Int(i), arg__15257)
	if v314 {
		goto b26
	} else {
		goto b27
	}
b24:
	;
	goto b25
b25:
	;
	goto b18
b26:
	;
	pid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{params, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	a, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{m, rt.LookupVar("clojure.core", "update").Deref(), pid, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v vm.Value
		var a_4 vm.Value
		var a_7 vm.Value
		var a_15 vm.Value
		var or__x vm.Value
		var a_19 vm.Value
		var head__15283 vm.Value
		var a_23 vm.Value
		var arg__15284 vm.Value
		var a_32 vm.Value
		var v34 vm.Value
		var callErr error
		if vm.IsTruthy(arg0) {
			v = arg0
			a_4 = a
			goto b1
		} else {
			v = arg0
			a_7 = a
			goto b2
		}
	b1:
		;
		a_15 = a_4
		goto b3
	b2:
		;
		a_15 = a_7
		goto b3
	b3:
		;
		if vm.IsTruthy(v) {
			or__x = v
			a_19 = a_15
			head__15283 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b4
		} else {
			a_23 = a_15
			head__15283 = rt.LookupVar("clojure.core", "conj").Deref()
			goto b5
		}
	b4:
		;
		arg__15284 = or__x
		a_32 = a_19
		goto b6
	b5:
		;
		arg__15284 = vm.NewArrayVector([]vm.Value{})
		a_32 = a_23
		goto b6
	b6:
		;
		v34, callErr = rt.InvokeValue(head__15283, []vm.Value{arg__15284, a_32})
		if callErr != nil {
			return nil, callErr
		}
		return v34, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	goto b28
b27:
	;
	goto b28
b28:
	;
	v352 = i + 1
	i = v352
	goto b22
}
func const_expr(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v7 vm.Value
	var v12 vm.Value
	var aux vm.Value
	var t vm.Value
	var v19 bool
	var v194 vm.Value
	var v24 vm.Value
	var v31 bool
	var v190 vm.Value
	var v36 vm.Value
	var v43 vm.Value
	var v186 vm.Value
	var v46 vm.Value
	var v53 vm.Value
	var v182 vm.Value
	var arg__15334 vm.Value
	var arg__15341 vm.Value
	var arg__15353 vm.Value
	var arg__15355 vm.Value
	var arg__15360 vm.Value
	var arg__15361 vm.Value
	var v88 vm.Value
	var v95 vm.Value
	var v178 vm.Value
	var or__x_104 bool
	var v174 vm.Value
	var v126 vm.Value
	var v133 vm.Value
	var v170 vm.Value
	var or__x_107 bool
	var arg__15372 vm.Value
	var v118 bool
	var v120 bool
	var v136 vm.Value
	var v143 vm.Value
	var v166 vm.Value
	var v146 vm.Value
	var v162 vm.Value
	var callErr error
	v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v7) {
		goto b1
	} else {
		aux = arg0
		t = arg1
		goto b2
	}
b1:
	;
	v12, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	v194 = v12
	goto b3
b2:
	;
	v19 = aux == vm.Boolean(true)
	if v19 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v194, nil
b4:
	;
	v24, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("true")})
	if callErr != nil {
		return nil, callErr
	}
	v190 = v24
	goto b6
b5:
	;
	v31 = aux == vm.Boolean(false)
	if v31 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v194 = v190
	goto b3
b7:
	;
	v36, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("false")})
	if callErr != nil {
		return nil, callErr
	}
	v186 = v36
	goto b9
b8:
	;
	v43, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v43) {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v190 = v186
	goto b6
b10:
	;
	v46, callErr = rt.InvokeValue(rt.LookupVar("gogen", "string-lit").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v182 = v46
	goto b12
b11:
	;
	v53, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "char?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v53) {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v186 = v182
	goto b9
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15334, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__15334, vm.String("Char")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15341, callErr = rt.InvokeValue(rt.LookupVar("gogen", "char-lit").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__15341})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15353, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15355, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__15353, vm.String("Char")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15360, callErr = rt.InvokeValue(rt.LookupVar("gogen", "char-lit").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__15361, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__15360})
	if callErr != nil {
		return nil, callErr
	}
	v88, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__15355, arg__15361})
	if callErr != nil {
		return nil, callErr
	}
	v178 = v88
	goto b15
b14:
	;
	v95, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "keyword?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v95) {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v182 = v178
	goto b12
b16:
	;
	v174 = vm.NIL
	goto b18
b17:
	;
	or__x_104 = t == vm.Keyword("float")
	if or__x_104 {
		or__x_107 = or__x_104
		goto b22
	} else {
		goto b23
	}
b18:
	;
	v178 = v174
	goto b15
b19:
	;
	v126, callErr = rt.InvokeValue(rt.LookupVar("gogen", "float-lit").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v170 = v126
	goto b21
b20:
	;
	v133, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "integer?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v133) {
		goto b25
	} else {
		goto b26
	}
b21:
	;
	v174 = v170
	goto b18
b22:
	;
	v120 = or__x_107
	goto b24
b23:
	;
	arg__15372, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v118 = t == arg__15372
	v120 = v118
	goto b24
b24:
	;
	if v120 {
		goto b19
	} else {
		goto b20
	}
b25:
	;
	v136, callErr = rt.InvokeValue(rt.LookupVar("gogen", "int-lit").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v166 = v136
	goto b27
b26:
	;
	v143, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "float?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v143) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v170 = v166
	goto b21
b28:
	;
	v146, callErr = rt.InvokeValue(rt.LookupVar("gogen", "float-lit").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v162 = v146
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
	v166 = v162
	goto b27
b31:
	;
	goto b33
b32:
	;
	goto b33
b33:
	;
	v162 = vm.NIL
	goto b30
}
func const_inst_val(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__15394 vm.Value
	var v9 bool
	var f vm.Value
	var nid vm.Value
	var aux vm.Value
	var arg__15408 vm.Value
	var v24 vm.Value
	var v38 vm.Value
	var v27 vm.Value
	var v31 vm.Value
	var callErr error
	arg__15394, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v9 = arg__15394 == vm.Keyword("const")
	if v9 {
		f = arg0
		nid = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "any-fn-template?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__15408, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "any-fn-template?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__15408})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v24) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v38 = vm.NIL
	goto b3
b3:
	;
	return v38, nil
b4:
	;
	v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v31 = v27
	goto b6
b5:
	;
	v31 = vm.NIL
	goto b6
b6:
	;
	v38 = v31
	goto b3
}
func const_param_map(arg0 vm.Value) (vm.Value, error) {
	var known vm.Value
	var arg__15430 vm.Value
	var doseq_seq__15413 vm.Value
	var doseq_loop__15414 vm.Value
	var f vm.Value
	var bid vm.Value
	var arg__15446 vm.Value
	var doseq_seq__15415 vm.Value
	var doseq_loop__15416 vm.Value
	var p vm.Value
	var v59 vm.Value
	var v72 vm.Value
	var changed vm.Value
	var arg__15476 vm.Value
	var doseq_seq__15417 vm.Value
	var doseq_loop__15418 vm.Value
	var params vm.Value
	var v320 vm.Value
	var i int
	var arg__15489 vm.Value
	var v144 bool
	var args vm.Value
	var v173 vm.Value
	var v298 vm.Value
	var arg__15581 vm.Value
	var v206 vm.Value
	var newv vm.Value
	var arg__15592 vm.Value
	var arg__15606 vm.Value
	var arg__15608 vm.Value
	var v257 vm.Value
	var v283 int
	var arg__15698 vm.Value
	var v340 vm.Value
	var callErr error
	known, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__15430, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__15413, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__15430})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15414 = doseq_seq__15413
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__15414) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__15414})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__15446, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__15415, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__15446})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15416 = doseq_seq__15415
	goto b5
b3:
	;
	goto b4
b4:
	;
	goto b9
b5:
	;
	if vm.IsTruthy(doseq_loop__15416) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__15416})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{known, rt.LookupVar("clojure.core", "assoc").Deref(), p, vm.Keyword("top")})
	if callErr != nil {
		return nil, callErr
	}
	v59, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__15416})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15416 = v59
	goto b5
b7:
	;
	goto b8
b8:
	;
	v72, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__15414})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15414 = v72
	goto b1
b9:
	;
	changed, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.Boolean(false)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__15476, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__15417, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__15476})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15418 = doseq_seq__15417
	goto b10
b10:
	;
	if vm.IsTruthy(doseq_loop__15418) {
		goto b11
	} else {
		goto b12
	}
b11:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__15418})
	if callErr != nil {
		return nil, callErr
	}
	params, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	goto b14
b12:
	;
	goto b13
b13:
	;
	v320, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{changed})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v320) {
		goto b24
	} else {
		goto b25
	}
b14:
	;
	arg__15489, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	v144 = rt.LtValue(vm.Int(i), arg__15489)
	if v144 {
		goto b15
	} else {
		goto b16
	}
b15:
	;
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{params, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	args, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-param-incoming-args").Deref(), []vm.Value{f, bid, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v173, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v173) {
		goto b18
	} else {
		goto b19
	}
b16:
	;
	goto b17
b17:
	;
	v298, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__15418})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__15418 = v298
	goto b10
b18:
	;
	newv = vm.Keyword("bot")
	goto b20
b19:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__15537 vm.Value
		var v9 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
		if callErr != nil {
			return nil, callErr
		}
		arg__15537, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
		if callErr != nil {
			return nil, callErr
		}
		v9, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "cp-value").Deref(), []vm.Value{f, arg__15537, arg0, p})
		if callErr != nil {
			return nil, callErr
		}
		return v9, nil
	}), args})
	if callErr != nil {
		return nil, callErr
	}
	arg__15581, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__15576 vm.Value
		var v9 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
		if callErr != nil {
			return nil, callErr
		}
		arg__15576, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
		if callErr != nil {
			return nil, callErr
		}
		v9, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "cp-value").Deref(), []vm.Value{f, arg__15576, arg0, p})
		if callErr != nil {
			return nil, callErr
		}
		return v9, nil
	}), args})
	if callErr != nil {
		return nil, callErr
	}
	v206, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "cp-meet").Deref(), vm.Keyword("top"), arg__15581})
	if callErr != nil {
		return nil, callErr
	}
	newv = v206
	goto b20
b20:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
	if callErr != nil {
		return nil, callErr
	}
	arg__15592, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__15592, p})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
	if callErr != nil {
		return nil, callErr
	}
	arg__15606, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
	if callErr != nil {
		return nil, callErr
	}
	arg__15608, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__15606, p})
	if callErr != nil {
		return nil, callErr
	}
	v257, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{newv, arg__15608})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v257) {
		goto b21
	} else {
		goto b22
	}
b21:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{known, rt.LookupVar("clojure.core", "assoc").Deref(), p, newv})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{changed, vm.Boolean(true)})
	if callErr != nil {
		return nil, callErr
	}
	goto b23
b22:
	;
	goto b23
b23:
	;
	v283 = i + 1
	i = v283
	goto b14
b24:
	;
	goto b9
b25:
	;
	goto b26
b26:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
	if callErr != nil {
		return nil, callErr
	}
	arg__15698, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{known})
	if callErr != nil {
		return nil, callErr
	}
	v340, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var arg__15673 vm.Value
		var v11 vm.Value
		var m vm.Value
		var kv vm.Value
		var arg__15688 vm.Value
		var arg__15692 vm.Value
		var v22 vm.Value
		var v25 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{arg1})
		if callErr != nil {
			return nil, callErr
		}
		arg__15673, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{arg1})
		if callErr != nil {
			return nil, callErr
		}
		v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{arg__15673})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v11) {
			m = arg0
			kv = arg1
			goto b1
		} else {
			m = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{kv})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{kv})
		if callErr != nil {
			return nil, callErr
		}
		arg__15688, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "key").Deref(), []vm.Value{kv})
		if callErr != nil {
			return nil, callErr
		}
		arg__15692, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "val").Deref(), []vm.Value{kv})
		if callErr != nil {
			return nil, callErr
		}
		v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{m, arg__15688, arg__15692})
		if callErr != nil {
			return nil, callErr
		}
		v25 = v22
		goto b3
	b2:
		;
		v25 = m
		goto b3
	b3:
		;
		return v25, nil
	}), vm.EmptyPersistentMap, arg__15698})
	if callErr != nil {
		return nil, callErr
	}
	return v340, nil
}
func const_param_of(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__15703 vm.Value
	var tem__G__0 vm.Value
	var nid vm.Value
	var v14 vm.Value
	var v18 vm.Value
	var callErr error
	arg__15703, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	tem__G__0, callErr = rt.InvokeValue(vm.Keyword("const-params"), []vm.Value{arg__15703})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(tem__G__0) {
		nid = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{tem__G__0, nid})
	if callErr != nil {
		return nil, callErr
	}
	v18 = v14
	goto b3
b2:
	;
	v18 = vm.NIL
	goto b3
b3:
	;
	return v18, nil
}
func cp_meet(arg0 vm.Value, arg1 vm.Value) vm.Value {
	var v7 bool
	var b vm.Value
	var a vm.Value
	var v15 bool
	var v73 vm.Value
	var v23 bool
	var v69 vm.Value
	var v32 bool
	var v65 vm.Value
	var v40 bool
	var v61 vm.Value
	var v57 vm.Value
	var v53 vm.Value
	v7 = arg0 == vm.Keyword("top")
	if v7 {
		b = arg1
		goto b1
	} else {
		a = arg0
		b = arg1
		goto b2
	}
b1:
	;
	v73 = b
	goto b3
b2:
	;
	v15 = b == vm.Keyword("top")
	if v15 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v73
b4:
	;
	v69 = a
	goto b6
b5:
	;
	v23 = a == vm.Keyword("bot")
	if v23 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v73 = v69
	goto b3
b7:
	;
	v65 = vm.Keyword("bot")
	goto b9
b8:
	;
	v32 = b == vm.Keyword("bot")
	if v32 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v69 = v65
	goto b6
b10:
	;
	v61 = vm.Keyword("bot")
	goto b12
b11:
	;
	v40 = a == b
	if v40 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v65 = v61
	goto b9
b13:
	;
	v57 = a
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
	v61 = v57
	goto b12
b16:
	;
	v53 = vm.Keyword("bot")
	goto b18
b17:
	;
	v53 = vm.NIL
	goto b18
b18:
	;
	v57 = v53
	goto b15
}
func cp_value(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var v12 bool
	var f vm.Value
	var known vm.Value
	var a vm.Value
	var or__x vm.Value
	var v42 vm.Value
	var v33 vm.Value
	var v35 vm.Value
	var callErr error
	v12 = arg2 == arg3
	if v12 {
		goto b1
	} else {
		f = arg0
		known = arg1
		a = arg2
		goto b2
	}
b1:
	;
	v42 = vm.Keyword("top")
	goto b3
b2:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "const-inst-val").Deref(), []vm.Value{f, a})
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
	return v42, nil
b4:
	;
	v35 = or__x
	goto b6
b5:
	;
	v33, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{known, a, vm.Keyword("bot")})
	if callErr != nil {
		return nil, callErr
	}
	v35 = v33
	goto b6
b6:
	;
	v42 = v35
	goto b3
}
func distinct_imports(arg0 vm.Value) (vm.Value, error) {
	var v5 vm.Value
	var remaining vm.Value
	var seen vm.Value
	var out vm.Value
	var v17 vm.Value
	var entry vm.Value
	var arg__15743 vm.Value
	var arg__15746 vm.Value
	var key vm.Value
	var v41 vm.Value
	var v53 vm.Value
	var v44 vm.Value
	var v47 vm.Value
	var v49 vm.Value
	var v51 vm.Value
	var callErr error
	v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	remaining = arg0
	seen = v5
	out = vm.NewArrayVector([]vm.Value{})
	goto b1
b1:
	;
	v17, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v17) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v53 = out
	goto b4
b3:
	;
	entry, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__15743, callErr = rt.InvokeValue(vm.Keyword("path"), []vm.Value{entry})
	if callErr != nil {
		return nil, callErr
	}
	arg__15746, callErr = rt.InvokeValue(vm.Keyword("alias"), []vm.Value{entry})
	if callErr != nil {
		return nil, callErr
	}
	key, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__15743, arg__15746})
	if callErr != nil {
		return nil, callErr
	}
	v41, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{seen, key})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v41) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v53, nil
b5:
	;
	v44, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v44
	goto b1
b6:
	;
	v47, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v49, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{seen, key})
	if callErr != nil {
		return nil, callErr
	}
	v51, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, entry})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v47
	seen = v49
	out = v51
	goto b1
}
func emit_assignments_for_target(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__15777 vm.Value
	var params vm.Value
	var args vm.Value
	var i int
	var pairs vm.Value
	var f vm.Value
	var closed_exprs vm.Value
	var arg__15786 vm.Value
	var v36 bool
	var v39 vm.Value
	var pid vm.Value
	var arg_nid vm.Value
	var arg__15814 vm.Value
	var or__x vm.Value
	var v541 vm.Value
	var v178 int
	var lhs vm.Value
	var arg__15848 vm.Value
	var lhs_spec vm.Value
	var arg__15861 vm.Value
	var arg_spec vm.Value
	var v219 bool
	var v166 vm.Value
	var v154 vm.Value
	var v140 vm.Value
	var v142 vm.Value
	var v222 vm.Value
	var and__x_249 bool
	var rhs vm.Value
	var v472 vm.Value
	var raw vm.Value
	var v418 vm.Value
	var v278 bool
	var and__x_275 bool
	var v281 bool
	var arg__15891 vm.Value
	var arg__15907 vm.Value
	var arg__15909 vm.Value
	var arg__15912 vm.Value
	var v354 vm.Value
	var v358 vm.Value
	var v400 vm.Value
	var v404 vm.Value
	var v503 vm.Value
	var v505 int
	var v507 int
	var arg__15940 vm.Value
	var v513 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__15777, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	params, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{arg__15777, arg0})
	if callErr != nil {
		return nil, callErr
	}
	args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	pairs = vm.NewArrayVector([]vm.Value{})
	f = arg0
	closed_exprs = arg1
	goto b1
b1:
	;
	arg__15786, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	v36 = rt.GeValue(vm.Int(i), arg__15786)
	if v36 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v39, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-sequence-copies").Deref(), []vm.Value{f, pairs})
	if callErr != nil {
		return nil, callErr
	}
	v541 = v39
	goto b4
b3:
	;
	pid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{params, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg_nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f, pid})
	if callErr != nil {
		return nil, callErr
	}
	arg__15814, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f, pid})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__15814})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b8
	} else {
		goto b9
	}
b4:
	;
	return v541, nil
b5:
	;
	v178 = i + 1
	i = v178
	goto b1
b6:
	;
	lhs, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-ident").Deref(), []vm.Value{f, pid})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{pid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__15848, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{pid, f})
	if callErr != nil {
		return nil, callErr
	}
	lhs_spec, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__15848})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg_nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__15861, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg_nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg_spec, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__15861})
	if callErr != nil {
		return nil, callErr
	}
	v219 = lhs_spec == vm.String("vm.Value")
	if v219 {
		goto b17
	} else {
		goto b18
	}
b7:
	;
	v541 = vm.NIL
	goto b4
b8:
	;
	v166 = or__x
	goto b10
b9:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "const-param-of").Deref(), []vm.Value{f, pid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b11
	} else {
		goto b12
	}
b10:
	;
	if vm.IsTruthy(v166) {
		goto b5
	} else {
		goto b6
	}
b11:
	;
	v154 = or__x
	goto b13
b12:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-value?").Deref(), []vm.Value{f, pid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b14
	} else {
		goto b15
	}
b13:
	;
	v166 = v154
	goto b10
b14:
	;
	v142 = or__x
	goto b16
b15:
	;
	v140, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "closure-value?").Deref(), []vm.Value{f, arg_nid})
	if callErr != nil {
		return nil, callErr
	}
	v142 = v140
	goto b16
b16:
	;
	v154 = v142
	goto b13
b17:
	;
	v222, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "box-as-value").Deref(), []vm.Value{f, closed_exprs, arg_nid})
	if callErr != nil {
		return nil, callErr
	}
	rhs = v222
	goto b19
b18:
	;
	and__x_249 = lhs_spec == vm.String("bool")
	if and__x_249 {
		goto b23
	} else {
		and__x_275 = and__x_249
		goto b24
	}
b19:
	;
	v472, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{rhs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v472) {
		goto b32
	} else {
		goto b33
	}
b20:
	;
	raw, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-expr").Deref(), []vm.Value{f, closed_exprs, arg_nid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(raw) {
		goto b26
	} else {
		goto b27
	}
b21:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b29
	} else {
		goto b30
	}
b22:
	;
	rhs = v418
	goto b19
b23:
	;
	v278 = arg_spec == vm.String("vm.Value")
	v281 = v278
	goto b25
b24:
	;
	v281 = and__x_275
	goto b25
b25:
	;
	if v281 {
		goto b20
	} else {
		goto b21
	}
b26:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15891, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__15891, vm.String("IsTruthy")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{raw})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15907, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15909, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__15907, vm.String("IsTruthy")})
	if callErr != nil {
		return nil, callErr
	}
	arg__15912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{raw})
	if callErr != nil {
		return nil, callErr
	}
	v354, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__15909, arg__15912})
	if callErr != nil {
		return nil, callErr
	}
	v358 = v354
	goto b28
b27:
	;
	v358 = vm.NIL
	goto b28
b28:
	;
	v418 = v358
	goto b22
b29:
	;
	v400, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-expr").Deref(), []vm.Value{f, closed_exprs, arg_nid})
	if callErr != nil {
		return nil, callErr
	}
	v404 = v400
	goto b31
b30:
	;
	v404 = vm.NIL
	goto b31
b31:
	;
	v418 = v404
	goto b22
b32:
	;
	goto b34
b33:
	;
	v503, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "same-ident?").Deref(), []vm.Value{lhs, rhs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v503) {
		goto b35
	} else {
		goto b36
	}
b34:
	;
	goto b7
b35:
	;
	v505 = i + 1
	i = v505
	goto b1
b36:
	;
	v507 = i + 1
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{lhs, rhs})
	if callErr != nil {
		return nil, callErr
	}
	arg__15940, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{lhs, rhs})
	if callErr != nil {
		return nil, callErr
	}
	v513, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{pairs, arg__15940})
	if callErr != nil {
		return nil, callErr
	}
	i = v507
	pairs = v513
	goto b1
}
func file(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var results vm.Value
	var arg__15964 vm.Value
	var imports vm.Value
	var decls vm.Value
	var import_nodes vm.Value
	var v23 vm.Value
	var callErr error
	results, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var entries vm.Value
		var v11 vm.Value
		var callErr error
		entries, callErr = rt.InvokeValue(vm.Keyword("imports"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(entries) {
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v11 = entries
		goto b3
	b2:
		;
		v11 = vm.NewArrayVector([]vm.Value{})
		goto b3
	b3:
		;
		return v11, nil
	}), results})
	if callErr != nil {
		return nil, callErr
	}
	arg__15964, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapcat").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var entries vm.Value
		var v11 vm.Value
		var callErr error
		entries, callErr = rt.InvokeValue(vm.Keyword("imports"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(entries) {
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		v11 = entries
		goto b3
	b2:
		;
		v11 = vm.NewArrayVector([]vm.Value{})
		goto b3
	b3:
		;
		return v11, nil
	}), results})
	if callErr != nil {
		return nil, callErr
	}
	imports, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "distinct-imports").Deref(), []vm.Value{arg__15964})
	if callErr != nil {
		return nil, callErr
	}
	decls, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v2 vm.Value
		var callErr error
		v2, callErr = rt.InvokeValue(vm.Keyword("decl"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v2, nil
	}), results})
	if callErr != nil {
		return nil, callErr
	}
	import_nodes, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "import-spec-node").Deref(), imports})
	if callErr != nil {
		return nil, callErr
	}
	v23, callErr = rt.InvokeValue(rt.LookupVar("gogen", "file").Deref(), []vm.Value{arg0, import_nodes, decls})
	if callErr != nil {
		return nil, callErr
	}
	return v23, nil
}
func fn_template_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var aux vm.Value
	var arg__15991 vm.Value
	var v11 bool
	var v14 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		aux = arg0
		goto b1
	} else {
		goto b2
	}
b1:
	;
	arg__15991, callErr = rt.InvokeValue(vm.Keyword("kind"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v11 = arg__15991 == vm.Keyword("fn-template")
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
func function_needs_error_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var blocks vm.Value
	var f vm.Value
	var v11 vm.Value
	var arg__16016 vm.Value
	var arg__16037 vm.Value
	var arg__16039 vm.Value
	var v38 vm.Value
	var v61 vm.Value
	var v57 vm.Value
	var v49 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v4
	f = arg0
	goto b1
b1:
	;
	v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v11) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v61 = vm.Boolean(false)
	goto b4
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__16016, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg__16016, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__16037, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__16039, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg__16037, f})
	if callErr != nil {
		return nil, callErr
	}
	v38, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__16026 vm.Value
		var v5 vm.Value
		var callErr error
		arg__16026, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		v5 = vm.Boolean(vm.Keyword("call") == arg__16026)
		return v5, nil
	}), arg__16039})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v38) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v61, nil
b5:
	;
	v57 = vm.Boolean(true)
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
	v61 = v57
	goto b4
b8:
	;
	v49, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v49
	goto b1
b9:
	;
	goto b10
b10:
	;
	v57 = vm.NIL
	goto b7
}
func function_needs_rt_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var blocks vm.Value
	var f vm.Value
	var v11 vm.Value
	var arg__16200 vm.Value
	var arg__16355 vm.Value
	var arg__16357 vm.Value
	var v38 vm.Value
	var local_QMARK__61 vm.Value
	var or__x_65 vm.Value
	var v57 vm.Value
	var v49 vm.Value
	var or__x_69 vm.Value
	var local_QMARK__70 vm.Value
	var v101 vm.Value
	var or__x_76 bool
	var arg__16374 vm.Value
	var v93 vm.Value
	var v95 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v4
	f = arg0
	goto b1
b1:
	;
	v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v11) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	local_QMARK__61 = vm.Boolean(false)
	goto b4
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__16200, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg__16200, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__16355, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__16357, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg__16355, f})
	if callErr != nil {
		return nil, callErr
	}
	v38, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var op vm.Value
		var aux vm.Value
		var refs vm.Value
		var or__x vm.Value
		var f_17 vm.Value
		var v301 vm.Value
		var f_33 vm.Value
		var v293 vm.Value
		var f_49 vm.Value
		var and__x vm.Value
		var v285 vm.Value
		var f_61 vm.Value
		var f_68 vm.Value
		var f_181 vm.Value
		var f_80 vm.Value
		var arg__16250 vm.Value
		var arg__16267 vm.Value
		var arg__16269 vm.Value
		var arg__16270 vm.Value
		var f_87 vm.Value
		var v169 vm.Value
		var f_171 vm.Value
		var f_119 vm.Value
		var f_126 vm.Value
		var arg__16285 vm.Value
		var arg__16302 vm.Value
		var arg__16304 vm.Value
		var arg__16305 vm.Value
		var v157 vm.Value
		var v159 vm.Value
		var f_161 vm.Value
		var f_195 vm.Value
		var v277 vm.Value
		var f_204 vm.Value
		var f_210 vm.Value
		var v218 vm.Value
		var f_222 vm.Value
		var f_230 vm.Value
		var arg__16324 vm.Value
		var arg__16341 vm.Value
		var arg__16343 vm.Value
		var arg__16344 vm.Value
		var v265 vm.Value
		var v268 vm.Value
		var callErr error
		op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		or__x = vm.Boolean(op == vm.Keyword("call"))
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			f_17 = f
			goto b2
		}
	b1:
		;
		v301 = or__x
		goto b3
	b2:
		;
		or__x = vm.Boolean(op == vm.Keyword("load-var"))
		if vm.IsTruthy(or__x) {
			goto b4
		} else {
			f_33 = f_17
			goto b5
		}
	b3:
		;
		return v301, nil
	b4:
		;
		v293 = or__x
		goto b6
	b5:
		;
		or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "any-fn-template?").Deref(), []vm.Value{aux})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(or__x) {
			goto b7
		} else {
			f_49 = f_33
			goto b8
		}
	b6:
		;
		v301 = v293
		goto b3
	b7:
		;
		v285 = or__x
		goto b9
	b8:
		;
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "binary-op").Deref(), op})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_61 = f_49
			goto b10
		} else {
			f_68 = f_49
			goto b11
		}
	b9:
		;
		v293 = v285
		goto b6
	b10:
		;
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{op, vm.Keyword("eq")})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_80 = f_61
			goto b13
		} else {
			f_87 = f_61
			goto b14
		}
	b11:
		;
		or__x = and__x
		f_181 = f_68
		goto b12
	b12:
		;
		if vm.IsTruthy(or__x) {
			goto b19
		} else {
			f_195 = f_181
			goto b20
		}
	b13:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16250, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg__16250, f_80})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16267, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16269, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg__16267, f_80})
		if callErr != nil {
			return nil, callErr
		}
		arg__16270, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__16269})
		if callErr != nil {
			return nil, callErr
		}
		or__x = vm.Boolean(vm.String("vm.Value") == arg__16270)
		if vm.IsTruthy(or__x) {
			f_119 = f_80
			goto b16
		} else {
			f_126 = f_80
			goto b17
		}
	b14:
		;
		v169 = and__x
		f_171 = f_87
		goto b15
	b15:
		;
		or__x = v169
		f_181 = f_171
		goto b12
	b16:
		;
		v159 = or__x
		f_161 = f_119
		goto b18
	b17:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16285, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg__16285, f_126})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16302, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16304, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg__16302, f_126})
		if callErr != nil {
			return nil, callErr
		}
		arg__16305, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__16304})
		if callErr != nil {
			return nil, callErr
		}
		v157 = vm.Boolean(vm.String("vm.Value") == arg__16305)
		v159 = v157
		f_161 = f_126
		goto b18
	b18:
		;
		v169 = v159
		f_171 = f_161
		goto b15
	b19:
		;
		v277 = or__x
		goto b21
	b20:
		;
		or__x = vm.Boolean(op == vm.Keyword("inc"))
		if vm.IsTruthy(or__x) {
			f_204 = f_195
			goto b22
		} else {
			f_210 = f_195
			goto b23
		}
	b21:
		;
		v285 = v277
		goto b9
	b22:
		;
		and__x = or__x
		f_222 = f_204
		goto b24
	b23:
		;
		v218 = vm.Boolean(op == vm.Keyword("dec"))
		and__x = v218
		f_222 = f_210
		goto b24
	b24:
		;
		if vm.IsTruthy(and__x) {
			f_230 = f_222
			goto b25
		} else {
			goto b26
		}
	b25:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16324, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg__16324, f_230})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16341, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		arg__16343, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg__16341, f_230})
		if callErr != nil {
			return nil, callErr
		}
		arg__16344, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__16343})
		if callErr != nil {
			return nil, callErr
		}
		v265 = vm.Boolean(vm.String("vm.Value") == arg__16344)
		v268 = v265
		goto b27
	b26:
		;
		v268 = and__x
		goto b27
	b27:
		;
		v277 = v268
		goto b21
	}), arg__16357})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v38) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	or__x_65, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x_65) {
		or__x_69 = or__x_65
		goto b11
	} else {
		local_QMARK__70 = local_QMARK__61
		goto b12
	}
b5:
	;
	v57 = vm.Boolean(true)
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
	local_QMARK__61 = v57
	goto b4
b8:
	;
	v49, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v49
	goto b1
b9:
	;
	goto b10
b10:
	;
	v57 = vm.NIL
	goto b7
b11:
	;
	v101 = or__x_69
	goto b13
b12:
	;
	if vm.IsTruthy(local_QMARK__70) {
		or__x_76 = vm.IsTruthy(local_QMARK__70)
		goto b14
	} else {
		goto b15
	}
b13:
	;
	return v101, nil
b14:
	;
	v95 = vm.Boolean(or__x_76)
	goto b16
b15:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "nested-template-fns").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__16374, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "nested-template-fns").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v93, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "function-needs-rt?").Deref(), arg__16374})
	if callErr != nil {
		return nil, callErr
	}
	v95 = v93
	goto b16
b16:
	;
	v101 = v95
	goto b13
}
func function_needs_tail_call_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var blocks vm.Value
	var f vm.Value
	var v11 vm.Value
	var v86 vm.Value
	var arg__16390 vm.Value
	var term vm.Value
	var v34 vm.Value
	var v82 vm.Value
	var v37 vm.Value
	var arg__16404 vm.Value
	var v48 bool
	var v75 vm.Value
	var v70 vm.Value
	var v61 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v4
	f = arg0
	goto b1
b1:
	;
	v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v11) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v86 = vm.Boolean(false)
	goto b4
b3:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v86, nil
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__16390, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{arg__16390, f})
	if callErr != nil {
		return nil, callErr
	}
	v34, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v34) {
		goto b8
	} else {
		goto b9
	}
b6:
	;
	v82 = vm.NIL
	goto b7
b7:
	;
	v86 = v82
	goto b4
b8:
	;
	v37, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v37
	goto b1
b9:
	;
	arg__16404, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	v48 = arg__16404 == vm.Keyword("tail-call")
	if v48 {
		goto b11
	} else {
		goto b12
	}
b10:
	;
	v82 = v75
	goto b7
b11:
	;
	v70 = vm.Boolean(true)
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
	v75 = v70
	goto b10
b14:
	;
	v61, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v61
	goto b1
b15:
	;
	goto b16
b16:
	;
	v70 = vm.NIL
	goto b13
}
func function_needs_vm_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var or__x vm.Value
	var f vm.Value
	var arg_types vm.Value
	var ret_ids vm.Value
	var v113 vm.Value
	var v107 vm.Value
	var v101 vm.Value
	var arg__16480 vm.Value
	var v93 vm.Value
	var v95 vm.Value
	var callErr error
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arg-types").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		f = arg0
		goto b1
	} else {
		f = arg0
		goto b2
	}
b1:
	;
	arg_types = or__x
	goto b3
b2:
	;
	arg_types = vm.NewArrayVector([]vm.Value{})
	goto b3
b3:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "return-ref-ids").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	ret_ids = or__x
	goto b6
b5:
	;
	ret_ids = vm.NewArrayVector([]vm.Value{})
	goto b6
b6:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b7
	} else {
		goto b8
	}
b7:
	;
	v113 = or__x
	goto b9
b8:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-needs-rt?").Deref(), []vm.Value{f})
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
	return v113, nil
b10:
	;
	v107 = or__x
	goto b12
b11:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__16432 vm.Value
		var v4 vm.Value
		var callErr error
		arg__16432, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v4 = vm.Boolean(vm.String("vm.Value") == arg__16432)
		return v4, nil
	}), arg_types})
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
	v113 = v107
	goto b9
b13:
	;
	v101 = or__x
	goto b15
b14:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__16466 vm.Value
		var arg__16467 vm.Value
		var v9 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__16466, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__16467, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__16466})
		if callErr != nil {
			return nil, callErr
		}
		v9 = vm.Boolean(vm.String("vm.Value") == arg__16467)
		return v9, nil
	}), ret_ids})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v107 = v101
	goto b12
b16:
	;
	v95 = or__x
	goto b18
b17:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "nested-template-fns").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__16480, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "nested-template-fns").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v93, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "function-needs-vm?").Deref(), arg__16480})
	if callErr != nil {
		return nil, callErr
	}
	v95 = v93
	goto b18
b18:
	;
	v101 = v95
	goto b15
}
func function_param_specs(arg0 vm.Value) (vm.Value, error) {
	var arity vm.Value
	var variadic_QMARK_ vm.Value
	var f vm.Value
	var v12 vm.Value
	var fixed_arity vm.Value
	var or__x vm.Value
	var arg_types vm.Value
	var i int
	var out vm.Value
	var v63 bool
	var load_arg_t vm.Value
	var arg__16502 vm.Value
	var v86 bool
	var v145 vm.Value
	var v89 vm.Value
	var meta_t vm.Value
	var v123 vm.Value
	var t vm.Value
	var v137 int
	var arg__16524 vm.Value
	var v143 vm.Value
	var callErr error
	arity, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arity").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	variadic_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(variadic_QMARK_) {
		f = arg0
		goto b1
	} else {
		f = arg0
		goto b2
	}
b1:
	;
	v12 = rt.SubValue(arity, vm.Int(1))
	fixed_arity = v12
	goto b3
b2:
	;
	fixed_arity = arity
	goto b3
b3:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arg-types").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	arg_types = or__x
	goto b6
b5:
	;
	arg_types = vm.NewArrayVector([]vm.Value{})
	goto b6
b6:
	;
	i = 0
	out = vm.NewArrayVector([]vm.Value{})
	goto b7
b7:
	;
	v63 = rt.GeValue(vm.Int(i), fixed_arity)
	if v63 {
		goto b8
	} else {
		goto b9
	}
b8:
	;
	v145 = out
	goto b10
b9:
	;
	load_arg_t, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-load-arg-type").Deref(), []vm.Value{f, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__16502, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg_types})
	if callErr != nil {
		return nil, callErr
	}
	v86 = rt.LtValue(vm.Int(i), arg__16502)
	if v86 {
		goto b11
	} else {
		goto b12
	}
b10:
	;
	return v145, nil
b11:
	;
	v89, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg_types, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	meta_t = v89
	goto b13
b12:
	;
	meta_t = vm.Keyword("unknown")
	goto b13
b13:
	;
	v123, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("unknown"), load_arg_t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v123) {
		goto b14
	} else {
		goto b15
	}
b14:
	;
	t = load_arg_t
	goto b16
b15:
	;
	t = meta_t
	goto b16
b16:
	;
	v137 = i + 1
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__16524, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v143, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, arg__16524})
	if callErr != nil {
		return nil, callErr
	}
	i = v137
	out = v143
	goto b7
}
func function_return_spec(arg0 vm.Value) (vm.Value, error) {
	var ret_ids vm.Value
	var f vm.Value
	var specs vm.Value
	var and__x vm.Value
	var v82 vm.Value
	var arg__16536 vm.Value
	var v19 vm.Value
	var v22 vm.Value
	var v71 vm.Value
	var v75 vm.Value
	var arg__16582 vm.Value
	var arg__16583 vm.Value
	var v61 bool
	var v64 vm.Value
	var callErr error
	ret_ids, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "return-ref-ids").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(ret_ids) {
		f = arg0
		goto b4
	} else {
		f = arg0
		and__x = ret_ids
		goto b5
	}
b1:
	;
	specs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__16565 vm.Value
		var v7 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__16565, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		v7, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__16565})
		if callErr != nil {
			return nil, callErr
		}
		return v7, nil
	}), ret_ids})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("clojure.core", "identity").Deref(), specs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b10
	} else {
		goto b11
	}
b2:
	;
	v82 = vm.NIL
	goto b3
b3:
	;
	return v82, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{ret_ids})
	if callErr != nil {
		return nil, callErr
	}
	arg__16536, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{ret_ids})
	if callErr != nil {
		return nil, callErr
	}
	v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__16536})
	if callErr != nil {
		return nil, callErr
	}
	v22 = v19
	goto b6
b5:
	;
	v22 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v22) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	v71, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	v75 = v71
	goto b9
b8:
	;
	v75 = vm.NIL
	goto b9
b9:
	;
	v82 = v75
	goto b3
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "distinct").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	arg__16582, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "distinct").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	arg__16583, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__16582})
	if callErr != nil {
		return nil, callErr
	}
	v61 = arg__16583 == vm.Int(1)
	v64 = vm.Boolean(v61)
	goto b12
b11:
	;
	v64 = and__x
	goto b12
b12:
	;
	if vm.IsTruthy(v64) {
		goto b7
	} else {
		goto b8
	}
}
func go_name(arg0 vm.Value) (vm.Value, error) {
	var arg__16621 vm.Value
	var munged vm.Value
	var v20 vm.Value
	var v25 vm.Value
	var v28 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var or__x vm.Value
		var ch vm.Value
		var v12 vm.Value
		var callErr error
		or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "go-name-munge-map").Deref(), arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			ch = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v12 = ch
		goto b3
	b3:
		;
		return v12, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__16621, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var or__x vm.Value
		var ch vm.Value
		var v12 vm.Value
		var callErr error
		or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "go-name-munge-map").Deref(), arg0})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(or__x) {
			goto b1
		} else {
			ch = arg0
			goto b2
		}
	b1:
		;
		v12 = or__x
		goto b3
	b2:
		;
		v12 = ch
		goto b3
	b3:
		;
		return v12, nil
	}), arg0})
	if callErr != nil {
		return nil, callErr
	}
	munged, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__16621})
	if callErr != nil {
		return nil, callErr
	}
	v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "go-reserved-words").Deref(), munged})
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
	v25, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{munged, vm.String("_")})
	if callErr != nil {
		return nil, callErr
	}
	v28 = v25
	goto b3
b2:
	;
	v28 = munged
	goto b3
b3:
	;
	return v28, nil
}
func go_type_spec(arg0 vm.Value) (vm.Value, error) {
	var or__x_4 bool
	var t vm.Value
	var or__x_27 bool
	var v173 vm.Value
	var or__x_6 bool
	var arg__16639 vm.Value
	var v16 bool
	var v18 bool
	var or__x_50 bool
	var v170 vm.Value
	var or__x_29 bool
	var arg__16647 vm.Value
	var v39 bool
	var v41 bool
	var v81 bool
	var v167 vm.Value
	var or__x_52 bool
	var or__x_58 bool
	var v72 bool
	var or__x_60 bool
	var v66 bool
	var v68 bool
	var v88 bool
	var v164 vm.Value
	var or__x_95 bool
	var v161 vm.Value
	var and__x vm.Value
	var v158 vm.Value
	var or__x_97 bool
	var or__x_103 bool
	var v117 bool
	var or__x_105 bool
	var v111 bool
	var v113 bool
	var v155 vm.Value
	var arg__16670 vm.Value
	var v135 bool
	var v138 vm.Value
	var callErr error
	or__x_4 = arg0 == vm.Keyword("int")
	if or__x_4 {
		t = arg0
		or__x_6 = or__x_4
		goto b4
	} else {
		t = arg0
		goto b5
	}
b1:
	;
	v173 = vm.String("int")
	goto b3
b2:
	;
	or__x_27 = t == vm.Keyword("float")
	if or__x_27 {
		or__x_29 = or__x_27
		goto b10
	} else {
		goto b11
	}
b3:
	;
	return v173, nil
b4:
	;
	v18 = or__x_6
	goto b6
b5:
	;
	arg__16639, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("int"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v16 = t == arg__16639
	v18 = v16
	goto b6
b6:
	;
	if v18 {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	v170 = vm.String("float64")
	goto b9
b8:
	;
	or__x_50 = t == vm.Keyword("bool")
	if or__x_50 {
		or__x_52 = or__x_50
		goto b16
	} else {
		goto b17
	}
b9:
	;
	v173 = v170
	goto b3
b10:
	;
	v41 = or__x_29
	goto b12
b11:
	;
	arg__16647, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("const"), vm.Keyword("float"), vm.Float(0)})
	if callErr != nil {
		return nil, callErr
	}
	v39 = t == arg__16647
	v41 = v39
	goto b12
b12:
	;
	if v41 {
		goto b7
	} else {
		goto b8
	}
b13:
	;
	v167 = vm.String("bool")
	goto b15
b14:
	;
	v81 = t == vm.Keyword("string")
	if v81 {
		goto b22
	} else {
		goto b23
	}
b15:
	;
	v170 = v167
	goto b9
b16:
	;
	v72 = or__x_52
	goto b18
b17:
	;
	or__x_58 = t == vm.Keyword("true")
	if or__x_58 {
		or__x_60 = or__x_58
		goto b19
	} else {
		goto b20
	}
b18:
	;
	if v72 {
		goto b13
	} else {
		goto b14
	}
b19:
	;
	v68 = or__x_60
	goto b21
b20:
	;
	v66 = t == vm.Keyword("false")
	v68 = v66
	goto b21
b21:
	;
	v72 = v68
	goto b18
b22:
	;
	v164 = vm.String("string")
	goto b24
b23:
	;
	v88 = t == vm.Keyword("char")
	if v88 {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v167 = v164
	goto b15
b25:
	;
	v161 = vm.String("vm.Char")
	goto b27
b26:
	;
	or__x_95 = t == vm.Keyword("unknown")
	if or__x_95 {
		or__x_97 = or__x_95
		goto b31
	} else {
		goto b32
	}
b27:
	;
	v164 = v161
	goto b24
b28:
	;
	v158 = vm.String("vm.Value")
	goto b30
b29:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b40
	} else {
		goto b41
	}
b30:
	;
	v161 = v158
	goto b27
b31:
	;
	v117 = or__x_97
	goto b33
b32:
	;
	or__x_103 = t == vm.Keyword("any")
	if or__x_103 {
		or__x_105 = or__x_103
		goto b34
	} else {
		goto b35
	}
b33:
	;
	if v117 {
		goto b28
	} else {
		goto b29
	}
b34:
	;
	v113 = or__x_105
	goto b36
b35:
	;
	v111 = t == vm.Keyword("nil")
	v113 = v111
	goto b36
b36:
	;
	v117 = v113
	goto b33
b37:
	;
	v155 = vm.String("vm.Value")
	goto b39
b38:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b43
	} else {
		goto b44
	}
b39:
	;
	v158 = v155
	goto b30
b40:
	;
	arg__16670, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v135 = arg__16670 == vm.Keyword("union")
	v138 = vm.Boolean(v135)
	goto b42
b41:
	;
	v138 = and__x
	goto b42
b42:
	;
	if vm.IsTruthy(v138) {
		goto b37
	} else {
		goto b38
	}
b43:
	;
	goto b45
b44:
	;
	goto b45
b45:
	;
	v155 = vm.NIL
	goto b39
}
func import_entry(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v5 vm.Value
	var callErr error
	v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("path"), arg0, vm.Keyword("alias"), arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v5, nil
}
func import_spec_node(arg0 vm.Value) (vm.Value, error) {
	var path vm.Value
	var alias vm.Value
	var v13 vm.Value
	var v16 vm.Value
	var v18 vm.Value
	var callErr error
	path, callErr = rt.InvokeValue(vm.Keyword("path"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	alias, callErr = rt.InvokeValue(vm.Keyword("alias"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(alias) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v13, callErr = rt.InvokeValue(rt.LookupVar("gogen", "import-spec").Deref(), []vm.Value{path, alias})
	if callErr != nil {
		return nil, callErr
	}
	v18 = v13
	goto b3
b2:
	;
	v16, callErr = rt.InvokeValue(rt.LookupVar("gogen", "import-spec").Deref(), []vm.Value{path})
	if callErr != nil {
		return nil, callErr
	}
	v18 = v16
	goto b3
b3:
	;
	return v18, nil
}
func infer_go_type(arg0 vm.Value) (vm.Value, error) {
	var spec vm.Value
	var v9 vm.Value
	var v13 vm.Value
	var callErr error
	spec, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(spec) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v9, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{spec})
	if callErr != nil {
		return nil, callErr
	}
	v13 = v9
	goto b3
b2:
	;
	v13 = vm.NIL
	goto b3
b3:
	;
	return v13, nil
}
func label_name(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("b"), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func live_nids(arg0 vm.Value) (vm.Value, error) {
	var arg__17212 vm.Value
	var live vm.Value
	var arg__17228 vm.Value
	var doseq_seq__17192 vm.Value
	var doseq_loop__17193 vm.Value
	var f vm.Value
	var bid vm.Value
	var term vm.Value
	var op vm.Value
	var or__x_69 bool
	var arg__17282 vm.Value
	var doseq_seq__17196 vm.Value
	var arg__17260 vm.Value
	var doseq_seq__17194 vm.Value
	var or__x_78 bool
	var or__x_91 bool
	var v126 bool
	var or__x_100 bool
	var v113 bool
	var v115 bool
	var doseq_loop__17195 vm.Value
	var r vm.Value
	var v171 vm.Value
	var doseq_loop__17197 vm.Value
	var nid vm.Value
	var arg__17292 vm.Value
	var v259 bool
	var v341 vm.Value
	var arg__17305 vm.Value
	var doseq_seq__17198 vm.Value
	var v327 vm.Value
	var doseq_loop__17199 vm.Value
	var v297 vm.Value
	var arg__17329 vm.Value
	var before vm.Value
	var arg__17338 vm.Value
	var doseq_seq__17200 vm.Value
	var doseq_loop__17201 vm.Value
	var arg__17354 vm.Value
	var doseq_seq__17202 vm.Value
	var arg__17485 vm.Value
	var arg__17486 vm.Value
	var v870 bool
	var doseq_loop__17203 vm.Value
	var arg__17367 vm.Value
	var v450 vm.Value
	var arg__17381 vm.Value
	var doseq_seq__17204 vm.Value
	var v526 vm.Value
	var doseq_loop__17205 vm.Value
	var v492 vm.Value
	var arg__17411 vm.Value
	var doseq_seq__17206 vm.Value
	var v840 vm.Value
	var doseq_loop__17207 vm.Value
	var bt vm.Value
	var tgt vm.Value
	var params vm.Value
	var args vm.Value
	var i int
	var arg__17430 vm.Value
	var v651 bool
	var arg__17435 vm.Value
	var and__x_689 bool
	var v808 vm.Value
	var arg__17469 vm.Value
	var v785 int
	var arg__17450 vm.Value
	var arg__17456 vm.Value
	var v736 vm.Value
	var and__x_725 bool
	var v739 vm.Value
	var v883 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__17212, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	live, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{arg__17212})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__17228, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__17192, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__17228})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17193 = doseq_seq__17192
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__17193) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__17193})
	if callErr != nil {
		return nil, callErr
	}
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(term) {
		goto b5
	} else {
		goto b6
	}
b3:
	;
	goto b4
b4:
	;
	goto b32
b5:
	;
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	or__x_69 = op == vm.Keyword("return")
	if or__x_69 {
		or__x_78 = or__x_69
		goto b11
	} else {
		goto b12
	}
b6:
	;
	goto b7
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17282, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__17196, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__17282})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17197 = doseq_seq__17196
	goto b21
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17260, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__17194, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__17260})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17195 = doseq_seq__17194
	goto b17
b9:
	;
	goto b10
b10:
	;
	goto b7
b11:
	;
	v126 = or__x_78
	goto b13
b12:
	;
	or__x_91 = op == vm.Keyword("branch-if")
	if or__x_91 {
		or__x_100 = or__x_91
		goto b14
	} else {
		goto b15
	}
b13:
	;
	if v126 {
		goto b8
	} else {
		goto b9
	}
b14:
	;
	v115 = or__x_100
	goto b16
b15:
	;
	v113 = op == vm.Keyword("tail-call")
	v115 = v113
	goto b16
b16:
	;
	v126 = v115
	goto b13
b17:
	;
	if vm.IsTruthy(doseq_loop__17195) {
		goto b18
	} else {
		goto b19
	}
b18:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__17195})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v5 vm.Value
		var callErr error
		v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{live, rt.LookupVar("clojure.core", "conj").Deref(), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v5, nil
	}), []vm.Value{r})
	if callErr != nil {
		return nil, callErr
	}
	v171, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__17195})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17195 = v171
	goto b17
b19:
	;
	goto b20
b20:
	;
	goto b10
b21:
	;
	if vm.IsTruthy(doseq_loop__17197) {
		goto b22
	} else {
		goto b23
	}
b22:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__17197})
	if callErr != nil {
		return nil, callErr
	}
	arg__17292, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v259 = arg__17292 == vm.Keyword("call")
	if v259 {
		goto b25
	} else {
		goto b26
	}
b23:
	;
	goto b24
b24:
	;
	v341, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__17193})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17193 = v341
	goto b1
b25:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17305, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__17198, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__17305})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17199 = doseq_seq__17198
	goto b28
b26:
	;
	goto b27
b27:
	;
	v327, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__17197})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17197 = v327
	goto b21
b28:
	;
	if vm.IsTruthy(doseq_loop__17199) {
		goto b29
	} else {
		goto b30
	}
b29:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__17199})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v5 vm.Value
		var callErr error
		v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{live, rt.LookupVar("clojure.core", "conj").Deref(), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v5, nil
	}), []vm.Value{r})
	if callErr != nil {
		return nil, callErr
	}
	v297, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__17199})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17199 = v297
	goto b28
b30:
	;
	goto b31
b31:
	;
	goto b27
b32:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	arg__17329, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	before, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__17329})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17338, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__17200, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__17338})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17201 = doseq_seq__17200
	goto b33
b33:
	;
	if vm.IsTruthy(doseq_loop__17201) {
		goto b34
	} else {
		goto b35
	}
b34:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__17201})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17354, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__17202, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__17354})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17203 = doseq_seq__17202
	goto b37
b35:
	;
	goto b36
b36:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	arg__17485, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	arg__17486, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__17485})
	if callErr != nil {
		return nil, callErr
	}
	v870 = rt.GtValue(arg__17486, before)
	if v870 {
		goto b65
	} else {
		goto b66
	}
b37:
	;
	if vm.IsTruthy(doseq_loop__17203) {
		goto b38
	} else {
		goto b39
	}
b38:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__17203})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	arg__17367, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	v450, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg__17367, nid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v450) {
		goto b41
	} else {
		goto b42
	}
b39:
	;
	goto b40
b40:
	;
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(term) {
		goto b48
	} else {
		goto b49
	}
b41:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17381, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__17204, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__17381})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17205 = doseq_seq__17204
	goto b44
b42:
	;
	goto b43
b43:
	;
	v526, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__17203})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17203 = v526
	goto b37
b44:
	;
	if vm.IsTruthy(doseq_loop__17205) {
		goto b45
	} else {
		goto b46
	}
b45:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__17205})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v5 vm.Value
		var callErr error
		v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{live, rt.LookupVar("clojure.core", "conj").Deref(), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v5, nil
	}), []vm.Value{r})
	if callErr != nil {
		return nil, callErr
	}
	v492, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__17205})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17205 = v492
	goto b44
b46:
	;
	goto b47
b47:
	;
	goto b43
b48:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-branch-targets").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17411, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-branch-targets").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__17206, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__17411})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17207 = doseq_seq__17206
	goto b51
b49:
	;
	goto b50
b50:
	;
	v840, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__17201})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17201 = v840
	goto b33
b51:
	;
	if vm.IsTruthy(doseq_loop__17207) {
		goto b52
	} else {
		goto b53
	}
b52:
	;
	bt, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__17207})
	if callErr != nil {
		return nil, callErr
	}
	tgt, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	params, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{tgt, f})
	if callErr != nil {
		return nil, callErr
	}
	args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	goto b55
b53:
	;
	goto b54
b54:
	;
	goto b50
b55:
	;
	arg__17430, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	v651 = rt.LtValue(vm.Int(i), arg__17430)
	if v651 {
		goto b56
	} else {
		goto b57
	}
b56:
	;
	arg__17435, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	and__x_689 = rt.LtValue(vm.Int(i), arg__17435)
	if and__x_689 {
		goto b62
	} else {
		and__x_725 = and__x_689
		goto b63
	}
b57:
	;
	goto b58
b58:
	;
	v808, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__17207})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__17207 = v808
	goto b51
b59:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__17469, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v5 vm.Value
		var callErr error
		v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{live, rt.LookupVar("clojure.core", "conj").Deref(), arg0})
		if callErr != nil {
			return nil, callErr
		}
		return v5, nil
	}), []vm.Value{arg__17469})
	if callErr != nil {
		return nil, callErr
	}
	goto b61
b60:
	;
	goto b61
b61:
	;
	v785 = i + 1
	i = v785
	goto b55
b62:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{params, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__17450, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	arg__17456, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{params, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v736, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg__17450, arg__17456})
	if callErr != nil {
		return nil, callErr
	}
	v739 = v736
	goto b64
b63:
	;
	v739 = vm.Boolean(and__x_725)
	goto b64
b64:
	;
	if vm.IsTruthy(v739) {
		goto b59
	} else {
		goto b60
	}
b65:
	;
	goto b32
b66:
	;
	goto b67
b67:
	;
	v883, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{live})
	if callErr != nil {
		return nil, callErr
	}
	return v883, nil
}
func live_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__17495 vm.Value
	var ls vm.Value
	var nid vm.Value
	var v14 vm.Value
	var f vm.Value
	var v17 vm.Value
	var v19 vm.Value
	var callErr error
	arg__17495, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	ls, callErr = rt.InvokeValue(vm.Keyword("live"), []vm.Value{arg__17495})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(ls) {
		nid = arg1
		goto b1
	} else {
		f = arg0
		nid = arg1
		goto b2
	}
b1:
	;
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{ls, nid})
	if callErr != nil {
		return nil, callErr
	}
	v19 = v14
	goto b3
b2:
	;
	v17, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "used?").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	v19 = v17
	goto b3
b3:
	;
	return v19, nil
}
func local_carrying_op_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var or__x_2 bool
	var or__x_4 bool
	var op vm.Value
	var or__x_10 bool
	var v50 vm.Value
	var or__x_12 bool
	var or__x_18 bool
	var v46 vm.Value
	var or__x_20 bool
	var or__x_26 bool
	var v42 vm.Value
	var or__x_28 bool
	var v36 vm.Value
	var v38 vm.Value
	var callErr error
	or__x_2 = arg0 == vm.Keyword("block-arg")
	if or__x_2 {
		or__x_4 = or__x_2
		goto b1
	} else {
		op = arg0
		goto b2
	}
b1:
	;
	v50 = vm.Boolean(or__x_4)
	goto b3
b2:
	;
	or__x_10 = op == vm.Keyword("call")
	if or__x_10 {
		or__x_12 = or__x_10
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v50, nil
b4:
	;
	v46 = vm.Boolean(or__x_12)
	goto b6
b5:
	;
	or__x_18 = op == vm.Keyword("inc")
	if or__x_18 {
		or__x_20 = or__x_18
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v50 = v46
	goto b3
b7:
	;
	v42 = vm.Boolean(or__x_20)
	goto b9
b8:
	;
	or__x_26 = op == vm.Keyword("dec")
	if or__x_26 {
		or__x_28 = or__x_26
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v46 = v42
	goto b6
b10:
	;
	v38 = vm.Boolean(or__x_28)
	goto b12
b11:
	;
	v36, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "binary-op").Deref(), op})
	if callErr != nil {
		return nil, callErr
	}
	v38 = v36
	goto b12
b12:
	;
	v42 = v38
	goto b9
}
func local_decls(arg0 vm.Value) (vm.Value, error) {
	var arg__17527 vm.Value
	var ids vm.Value
	var arg__17532 vm.Value
	var seen vm.Value
	var remaining vm.Value
	var out vm.Value
	var f vm.Value
	var v30 vm.Value
	var nid vm.Value
	var arg__17551 vm.Value
	var go_type vm.Value
	var nm vm.Value
	var v60 vm.Value
	var v157 vm.Value
	var arg__17569 vm.Value
	var v85 vm.Value
	var v88 vm.Value
	var v113 vm.Value
	var arg__17602 vm.Value
	var v123 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "collect-local-ids").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__17527, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "collect-local-ids").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	ids, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "distinct").Deref(), []vm.Value{arg__17527})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__17532, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	seen, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{arg__17532})
	if callErr != nil {
		return nil, callErr
	}
	remaining = ids
	out = vm.NewArrayVector([]vm.Value{})
	f = arg0
	goto b1
b1:
	;
	v30, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v30) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v157 = out
	goto b4
b3:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17551, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	go_type, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "infer-go-type").Deref(), []vm.Value{arg__17551})
	if callErr != nil {
		return nil, callErr
	}
	nm, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-name").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	v60, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{go_type})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v60) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v157, nil
b5:
	;
	goto b7
b6:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{seen})
	if callErr != nil {
		return nil, callErr
	}
	arg__17569, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{seen})
	if callErr != nil {
		return nil, callErr
	}
	v85, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg__17569, nm})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v85) {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	v157 = vm.NIL
	goto b4
b8:
	;
	v88, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v88
	goto b1
b9:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b11
	} else {
		goto b12
	}
b10:
	;
	goto b7
b11:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{seen, rt.LookupVar("clojure.core", "conj").Deref(), nm})
	if callErr != nil {
		return nil, callErr
	}
	v113, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "var-decl").Deref(), []vm.Value{nm, go_type, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	arg__17602, callErr = rt.InvokeValue(rt.LookupVar("gogen", "var-decl").Deref(), []vm.Value{nm, go_type, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v123, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, arg__17602})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v113
	out = v123
	goto b1
b12:
	;
	goto b13
b13:
	;
	goto b10
}
func lower_block_stmts(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__17619 vm.Value
	var v11 bool
	var f vm.Value
	var closed_exprs vm.Value
	var bid vm.Value
	var arg__17630 vm.Value
	var arg__17632 vm.Value
	var v24 vm.Value
	var label_stmts vm.Value
	var insts vm.Value
	var remaining vm.Value
	var out vm.Value
	var v53 vm.Value
	var arg__17653 vm.Value
	var stmts vm.Value
	var v79 vm.Value
	var inst_stmts vm.Value
	var term_stmts vm.Value
	var arg__17666 vm.Value
	var op vm.Value
	var and__x vm.Value
	var v370 vm.Value
	var v372 vm.Value
	var v354 vm.Value
	var or__x_131 bool
	var v341 vm.Value
	var or__x_142 bool
	var or__x_157 bool
	var v327 vm.Value
	var or__x_168 bool
	var or__x_183 bool
	var v314 vm.Value
	var or__x_194 bool
	var or__x_209 bool
	var v301 vm.Value
	var or__x_220 bool
	var or__x_235 bool
	var v288 vm.Value
	var or__x_246 bool
	var arg__17693 vm.Value
	var arg__17706 vm.Value
	var arg__17707 vm.Value
	var v273 vm.Value
	var v275 vm.Value
	var arg__17742 vm.Value
	var v452 vm.Value
	var v456 vm.Value
	var v435 vm.Value
	var callErr error
	arg__17619, callErr = rt.InvokeValue(rt.LookupVar("ir", "entry-block").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v11 = arg2 == arg__17619
	if v11 {
		f = arg0
		closed_exprs = arg1
		bid = arg2
		goto b1
	} else {
		f = arg0
		closed_exprs = arg1
		bid = arg2
		goto b2
	}
b1:
	;
	label_stmts = vm.NewArrayVector([]vm.Value{})
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "label-name").Deref(), []vm.Value{bid})
	if callErr != nil {
		return nil, callErr
	}
	arg__17630, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "label-name").Deref(), []vm.Value{bid})
	if callErr != nil {
		return nil, callErr
	}
	arg__17632, callErr = rt.InvokeValue(rt.LookupVar("gogen", "label-stmt").Deref(), []vm.Value{arg__17630, vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17632})
	if callErr != nil {
		return nil, callErr
	}
	label_stmts = v24
	goto b3
b3:
	;
	insts, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	remaining = insts
	out = vm.NewArrayVector([]vm.Value{})
	goto b4
b4:
	;
	v53, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v53) {
		goto b5
	} else {
		goto b6
	}
b5:
	;
	inst_stmts = out
	goto b7
b6:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__17653, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	stmts, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower-inst-stmts").Deref(), []vm.Value{f, closed_exprs, arg__17653})
	if callErr != nil {
		return nil, callErr
	}
	v79, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{stmts})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v79) {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	term_stmts, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower-terminator").Deref(), []vm.Value{f, closed_exprs, bid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(inst_stmts) {
		goto b35
	} else {
		and__x = inst_stmts
		goto b36
	}
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__17666, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg__17666, f})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{op, vm.Keyword("call")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b14
	} else {
		goto b15
	}
b9:
	;
	v370, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v372, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{out, stmts})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v370
	out = v372
	goto b4
b10:
	;
	inst_stmts = vm.NIL
	goto b7
b11:
	;
	v354, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v354
	goto b4
b12:
	;
	goto b13
b13:
	;
	goto b10
b14:
	;
	or__x_131 = op == vm.Keyword("load-arg")
	if or__x_131 {
		or__x_142 = or__x_131
		goto b17
	} else {
		goto b18
	}
b15:
	;
	v341 = and__x
	goto b16
b16:
	;
	if vm.IsTruthy(v341) {
		goto b11
	} else {
		goto b12
	}
b17:
	;
	v327 = vm.Boolean(or__x_142)
	goto b19
b18:
	;
	or__x_157 = op == vm.Keyword("load-var")
	if or__x_157 {
		or__x_168 = or__x_157
		goto b20
	} else {
		goto b21
	}
b19:
	;
	v341 = v327
	goto b16
b20:
	;
	v314 = vm.Boolean(or__x_168)
	goto b22
b21:
	;
	or__x_183 = op == vm.Keyword("load-closed")
	if or__x_183 {
		or__x_194 = or__x_183
		goto b23
	} else {
		goto b24
	}
b22:
	;
	v327 = v314
	goto b19
b23:
	;
	v301 = vm.Boolean(or__x_194)
	goto b25
b24:
	;
	or__x_209 = op == vm.Keyword("const")
	if or__x_209 {
		or__x_220 = or__x_209
		goto b26
	} else {
		goto b27
	}
b25:
	;
	v314 = v301
	goto b22
b26:
	;
	v288 = vm.Boolean(or__x_220)
	goto b28
b27:
	;
	or__x_235 = op == vm.Keyword("block-arg")
	if or__x_235 {
		or__x_246 = or__x_235
		goto b29
	} else {
		goto b30
	}
b28:
	;
	v301 = v288
	goto b25
b29:
	;
	v275 = vm.Boolean(or__x_246)
	goto b31
b30:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__17693, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f, arg__17693})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__17706, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__17707, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f, arg__17706})
	if callErr != nil {
		return nil, callErr
	}
	v273, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__17707})
	if callErr != nil {
		return nil, callErr
	}
	v275 = v273
	goto b31
b31:
	;
	v288 = v275
	goto b28
b32:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{label_stmts, inst_stmts, term_stmts})
	if callErr != nil {
		return nil, callErr
	}
	arg__17742, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{label_stmts, inst_stmts, term_stmts})
	if callErr != nil {
		return nil, callErr
	}
	v452, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__17742})
	if callErr != nil {
		return nil, callErr
	}
	v456 = v452
	goto b34
b33:
	;
	v456 = vm.NIL
	goto b34
b34:
	;
	return v456, nil
b35:
	;
	v435 = term_stmts
	goto b37
b36:
	;
	v435 = and__x
	goto b37
b37:
	;
	if vm.IsTruthy(v435) {
		goto b32
	} else {
		goto b33
	}
}
func lower_fn_STAR_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var params vm.Value
	var results vm.Value
	var body vm.Value
	var needs_rt_QMARK_ vm.Value
	var needs_vm_QMARK_ vm.Value
	var needs_error_QMARK_ vm.Value
	var f vm.Value
	var mode vm.Value
	var arg__17767 vm.Value
	var arg__17775 vm.Value
	var head__17777 vm.Value
	var arg__17784 vm.Value
	var v118 vm.Value
	var arg__17785 vm.Value
	var arg__17792 vm.Value
	var v160 vm.Value
	var arg__17793 vm.Value
	var head__17795 vm.Value
	var arg__17802 vm.Value
	var arg__17810 vm.Value
	var head__17812 vm.Value
	var arg__17819 vm.Value
	var v289 vm.Value
	var arg__17820 vm.Value
	var arg__17827 vm.Value
	var v334 vm.Value
	var arg__17828 vm.Value
	var arg__17829 vm.Value
	var imports vm.Value
	var v371 vm.Value
	var v376 vm.Value
	var v397 vm.Value
	var v527 vm.Value
	var v402 vm.Value
	var v423 vm.Value
	var v516 vm.Value
	var v428 vm.Value
	var v505 vm.Value
	var arg__17870 vm.Value
	var arg__17884 vm.Value
	var arg__17885 vm.Value
	var arg__17889 vm.Value
	var arg__17899 vm.Value
	var arg__17900 vm.Value
	var arg__17905 vm.Value
	var arg__17910 vm.Value
	var arg__17920 vm.Value
	var arg__17921 vm.Value
	var v490 vm.Value
	var v494 vm.Value
	var callErr error
	params, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "params-for").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	results, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "result-node").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower-body").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	needs_rt_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-needs-rt?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	needs_vm_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-needs-vm?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	needs_error_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-needs-error?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(needs_rt_QMARK_) {
		f = arg0
		mode = arg1
		goto b1
	} else {
		f = arg0
		mode = arg1
		goto b2
	}
b1:
	;
	arg__17767, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "import-entry").Deref(), []vm.Value{vm.String("github.com/nooga/let-go/pkg/rt"), vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17767})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	if vm.IsTruthy(needs_vm_QMARK_) {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	arg__17775, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "import-entry").Deref(), []vm.Value{vm.String("github.com/nooga/let-go/pkg/vm"), vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17775})
	if callErr != nil {
		return nil, callErr
	}
	goto b6
b5:
	;
	goto b6
b6:
	;
	if vm.IsTruthy(needs_rt_QMARK_) {
		head__17777 = rt.LookupVar("clojure.core", "concat").Deref()
		goto b7
	} else {
		head__17777 = rt.LookupVar("clojure.core", "concat").Deref()
		goto b8
	}
b7:
	;
	arg__17784, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "import-entry").Deref(), []vm.Value{vm.String("github.com/nooga/let-go/pkg/rt"), vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	v118, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17784})
	if callErr != nil {
		return nil, callErr
	}
	arg__17785 = v118
	goto b9
b8:
	;
	arg__17785 = vm.NewArrayVector([]vm.Value{})
	goto b9
b9:
	;
	if vm.IsTruthy(needs_vm_QMARK_) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	arg__17792, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "import-entry").Deref(), []vm.Value{vm.String("github.com/nooga/let-go/pkg/vm"), vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	v160, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17792})
	if callErr != nil {
		return nil, callErr
	}
	arg__17793 = v160
	goto b12
b11:
	;
	arg__17793 = vm.NewArrayVector([]vm.Value{})
	goto b12
b12:
	;
	_, callErr = rt.InvokeValue(head__17777, []vm.Value{arg__17785, arg__17793})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(needs_rt_QMARK_) {
		head__17795 = rt.LookupVar("clojure.core", "vec").Deref()
		goto b13
	} else {
		head__17795 = rt.LookupVar("clojure.core", "vec").Deref()
		goto b14
	}
b13:
	;
	arg__17802, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "import-entry").Deref(), []vm.Value{vm.String("github.com/nooga/let-go/pkg/rt"), vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17802})
	if callErr != nil {
		return nil, callErr
	}
	goto b15
b14:
	;
	goto b15
b15:
	;
	if vm.IsTruthy(needs_vm_QMARK_) {
		goto b16
	} else {
		goto b17
	}
b16:
	;
	arg__17810, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "import-entry").Deref(), []vm.Value{vm.String("github.com/nooga/let-go/pkg/vm"), vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17810})
	if callErr != nil {
		return nil, callErr
	}
	goto b18
b17:
	;
	goto b18
b18:
	;
	if vm.IsTruthy(needs_rt_QMARK_) {
		head__17812 = rt.LookupVar("clojure.core", "concat").Deref()
		goto b19
	} else {
		head__17812 = rt.LookupVar("clojure.core", "concat").Deref()
		goto b20
	}
b19:
	;
	arg__17819, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "import-entry").Deref(), []vm.Value{vm.String("github.com/nooga/let-go/pkg/rt"), vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	v289, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17819})
	if callErr != nil {
		return nil, callErr
	}
	arg__17820 = v289
	goto b21
b20:
	;
	arg__17820 = vm.NewArrayVector([]vm.Value{})
	goto b21
b21:
	;
	if vm.IsTruthy(needs_vm_QMARK_) {
		goto b22
	} else {
		goto b23
	}
b22:
	;
	arg__17827, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "import-entry").Deref(), []vm.Value{vm.String("github.com/nooga/let-go/pkg/vm"), vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	v334, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__17827})
	if callErr != nil {
		return nil, callErr
	}
	arg__17828 = v334
	goto b24
b23:
	;
	arg__17828 = vm.NewArrayVector([]vm.Value{})
	goto b24
b24:
	;
	arg__17829, callErr = rt.InvokeValue(head__17812, []vm.Value{arg__17820, arg__17828})
	if callErr != nil {
		return nil, callErr
	}
	imports, callErr = rt.InvokeValue(head__17795, []vm.Value{arg__17829})
	if callErr != nil {
		return nil, callErr
	}
	v371, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v371) {
		goto b25
	} else {
		goto b26
	}
b25:
	;
	v376, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "unsupported").Deref(), []vm.Value{mode, vm.String("unsupported parameter types")})
	if callErr != nil {
		return nil, callErr
	}
	v527 = v376
	goto b27
b26:
	;
	v397, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{results})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v397) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	return v527, nil
b28:
	;
	v402, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "unsupported").Deref(), []vm.Value{mode, vm.String("unsupported result type")})
	if callErr != nil {
		return nil, callErr
	}
	v516 = v402
	goto b30
b29:
	;
	v423, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{body})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v423) {
		goto b31
	} else {
		goto b32
	}
b30:
	;
	v527 = v516
	goto b27
b31:
	;
	v428, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "unsupported").Deref(), []vm.Value{mode, vm.String("unsupported function body shape")})
	if callErr != nil {
		return nil, callErr
	}
	v505 = v428
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
	v516 = v505
	goto b30
b34:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17870, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-name").Deref(), []vm.Value{arg__17870})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17884, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17885, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-name").Deref(), []vm.Value{arg__17884})
	if callErr != nil {
		return nil, callErr
	}
	arg__17889, callErr = rt.InvokeValue(rt.LookupVar("gogen", "func-decl").Deref(), []vm.Value{arg__17885, params, results, body})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17899, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17900, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{arg__17899})
	if callErr != nil {
		return nil, callErr
	}
	arg__17905, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "override-uniform-value?").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17910, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arity").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17920, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-name").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__17921, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-name").Deref(), []vm.Value{arg__17920})
	if callErr != nil {
		return nil, callErr
	}
	v490, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("imports"), imports, vm.Keyword("status"), vm.Keyword("lowered"), vm.Keyword("needs-error?"), needs_error_QMARK_, vm.Keyword("decl"), arg__17889, vm.Keyword("fn-name"), arg__17900, vm.Keyword("override-eligible?"), arg__17905, vm.Keyword("arity"), arg__17910, vm.Keyword("go-name"), arg__17921})
	if callErr != nil {
		return nil, callErr
	}
	v494 = v490
	goto b36
b35:
	;
	v494 = vm.NIL
	goto b36
b36:
	;
	v505 = v494
	goto b33
}
func lower_fn_lit(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var params vm.Value
	var results vm.Value
	var body vm.Value
	var v64 vm.Value
	var v68 vm.Value
	var and__x vm.Value
	var v55 vm.Value
	var v46 vm.Value
	var callErr error
	params, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "params-for").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	results, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "result-node").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	body, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower-body").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(params) {
		goto b4
	} else {
		and__x = params
		goto b5
	}
b1:
	;
	v64, callErr = rt.InvokeValue(rt.LookupVar("gogen", "func-lit").Deref(), []vm.Value{params, results, body})
	if callErr != nil {
		return nil, callErr
	}
	v68 = v64
	goto b3
b2:
	;
	v68 = vm.NIL
	goto b3
b3:
	;
	return v68, nil
b4:
	;
	if vm.IsTruthy(results) {
		goto b7
	} else {
		and__x = results
		goto b8
	}
b5:
	;
	v55 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v55) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	v46 = body
	goto b9
b8:
	;
	v46 = and__x
	goto b9
b9:
	;
	v55 = v46
	goto b6
}
func lower_inst_stmts(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var op vm.Value
	var aux vm.Value
	var v18 bool
	var f vm.Value
	var closed_exprs vm.Value
	var nid vm.Value
	var v33 bool
	var v428 vm.Value
	var v48 bool
	var v421 vm.Value
	var and__x_63 bool
	var v414 vm.Value
	var v102 bool
	var v407 vm.Value
	var v78 vm.Value
	var and__x_75 bool
	var v81 vm.Value
	var v117 bool
	var v400 vm.Value
	var v132 bool
	var v393 vm.Value
	var v147 bool
	var v386 vm.Value
	var v164 vm.Value
	var v379 vm.Value
	var v177 vm.Value
	var or__x_235 bool
	var v372 vm.Value
	var rhs vm.Value
	var v217 vm.Value
	var arg__18002 vm.Value
	var arg__18004 vm.Value
	var v203 vm.Value
	var v207 vm.Value
	var v272 vm.Value
	var v330 bool
	var v365 vm.Value
	var or__x_241 bool
	var v251 bool
	var v253 bool
	var v312 vm.Value
	var arg__18037 vm.Value
	var arg__18039 vm.Value
	var v298 vm.Value
	var v302 vm.Value
	var v333 vm.Value
	var v358 vm.Value
	var callErr error
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg2, arg0})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg2, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v18 = op == vm.Keyword("load-arg")
	if v18 {
		goto b1
	} else {
		f = arg0
		closed_exprs = arg1
		nid = arg2
		goto b2
	}
b1:
	;
	v428 = vm.NewArrayVector([]vm.Value{})
	goto b3
b2:
	;
	v33 = op == vm.Keyword("load-var")
	if v33 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v428, nil
b4:
	;
	v421 = vm.NewArrayVector([]vm.Value{})
	goto b6
b5:
	;
	v48 = op == vm.Keyword("load-closed")
	if v48 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v428 = v421
	goto b3
b7:
	;
	v414 = vm.NewArrayVector([]vm.Value{})
	goto b9
b8:
	;
	and__x_63 = op == vm.Keyword("const")
	if and__x_63 {
		goto b13
	} else {
		and__x_75 = and__x_63
		goto b14
	}
b9:
	;
	v421 = v414
	goto b6
b10:
	;
	v407 = vm.NewArrayVector([]vm.Value{})
	goto b12
b11:
	;
	v102 = op == vm.Keyword("const")
	if v102 {
		goto b16
	} else {
		goto b17
	}
b12:
	;
	v414 = v407
	goto b9
b13:
	;
	v78, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "any-fn-template?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v81 = v78
	goto b15
b14:
	;
	v81 = vm.Boolean(and__x_75)
	goto b15
b15:
	;
	if vm.IsTruthy(v81) {
		goto b10
	} else {
		goto b11
	}
b16:
	;
	v400 = vm.NewArrayVector([]vm.Value{})
	goto b18
b17:
	;
	v117 = op == vm.Keyword("make-closure")
	if v117 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v407 = v400
	goto b12
b19:
	;
	v393 = vm.NewArrayVector([]vm.Value{})
	goto b21
b20:
	;
	v132 = op == vm.Keyword("push-closed")
	if v132 {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v400 = v393
	goto b18
b22:
	;
	v386 = vm.NewArrayVector([]vm.Value{})
	goto b24
b23:
	;
	v147 = op == vm.Keyword("block-arg")
	if v147 {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v393 = v386
	goto b21
b25:
	;
	v379 = vm.NewArrayVector([]vm.Value{})
	goto b27
b26:
	;
	v164, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.lower-go", "binary-op").Deref(), op})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v164) {
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v386 = v379
	goto b24
b28:
	;
	v177, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v177) {
		goto b31
	} else {
		goto b32
	}
b29:
	;
	or__x_235 = op == vm.Keyword("inc")
	if or__x_235 {
		or__x_241 = or__x_235
		goto b40
	} else {
		goto b41
	}
b30:
	;
	v379 = v372
	goto b27
b31:
	;
	rhs, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "inst-rhs").Deref(), []vm.Value{f, closed_exprs, nid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(rhs) {
		goto b34
	} else {
		goto b35
	}
b32:
	;
	v217 = vm.NewArrayVector([]vm.Value{})
	goto b33
b33:
	;
	v372 = v217
	goto b30
b34:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-ident").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__18002, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-ident").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__18004, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__18002, rhs})
	if callErr != nil {
		return nil, callErr
	}
	v203, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__18004})
	if callErr != nil {
		return nil, callErr
	}
	v207 = v203
	goto b36
b35:
	;
	v207 = vm.NIL
	goto b36
b36:
	;
	v217 = v207
	goto b33
b37:
	;
	v272, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "live?").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v272) {
		goto b43
	} else {
		goto b44
	}
b38:
	;
	v330 = op == vm.Keyword("call")
	if v330 {
		goto b49
	} else {
		goto b50
	}
b39:
	;
	v372 = v365
	goto b30
b40:
	;
	v253 = or__x_241
	goto b42
b41:
	;
	v251 = op == vm.Keyword("dec")
	v253 = v251
	goto b42
b42:
	;
	if v253 {
		goto b37
	} else {
		goto b38
	}
b43:
	;
	rhs, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "inst-rhs").Deref(), []vm.Value{f, closed_exprs, nid})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(rhs) {
		goto b46
	} else {
		goto b47
	}
b44:
	;
	v312 = vm.NewArrayVector([]vm.Value{})
	goto b45
b45:
	;
	v365 = v312
	goto b39
b46:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-ident").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__18037, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-ident").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__18039, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__18037, rhs})
	if callErr != nil {
		return nil, callErr
	}
	v298, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__18039})
	if callErr != nil {
		return nil, callErr
	}
	v302 = v298
	goto b48
b47:
	;
	v302 = vm.NIL
	goto b48
b48:
	;
	v312 = v302
	goto b45
b49:
	;
	v333, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "call-assign-stmts").Deref(), []vm.Value{f, closed_exprs, nid})
	if callErr != nil {
		return nil, callErr
	}
	v358 = v333
	goto b51
b50:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b52
	} else {
		goto b53
	}
b51:
	;
	v365 = v358
	goto b39
b52:
	;
	goto b54
b53:
	;
	goto b54
b54:
	;
	v358 = vm.NIL
	goto b51
}
func lower_template_closure_expr(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v7 vm.Value
	var template vm.Value
	var capture_exprs vm.Value
	var arg__18059 vm.Value
	var inner_STAR_ vm.Value
	var v65 vm.Value
	var v170 vm.Value
	var arg__18070 vm.Value
	var arg__18086 vm.Value
	var arg__18088 vm.Value
	var arg__18091 vm.Value
	var v51 vm.Value
	var v55 vm.Value
	var arg__18121 vm.Value
	var inners vm.Value
	var v88 vm.Value
	var v166 vm.Value
	var boxed vm.Value
	var arg__18203 vm.Value
	var arg__18216 vm.Value
	var arg__18218 vm.Value
	var arg__18230 vm.Value
	var arg__18232 vm.Value
	var arg__18243 vm.Value
	var arg__18245 vm.Value
	var arg__18246 vm.Value
	var v143 vm.Value
	var v147 vm.Value
	var callErr error
	v7, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "fn-template?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v7) {
		template = arg0
		capture_exprs = arg1
		goto b1
	} else {
		template = arg0
		capture_exprs = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("fn"), []vm.Value{template})
	if callErr != nil {
		return nil, callErr
	}
	arg__18059, callErr = rt.InvokeValue(vm.Keyword("fn"), []vm.Value{template})
	if callErr != nil {
		return nil, callErr
	}
	inner_STAR_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower-fn-lit").Deref(), []vm.Value{arg__18059, capture_exprs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(inner_STAR_) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v65, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "multi-fn-template?").Deref(), []vm.Value{template})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v65) {
		goto b7
	} else {
		goto b8
	}
b3:
	;
	return v170, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18070, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__18070, vm.String("BoxNativeFn")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{inner_STAR_})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18086, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18088, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__18086, vm.String("BoxNativeFn")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18091, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{inner_STAR_})
	if callErr != nil {
		return nil, callErr
	}
	v51, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__18088, arg__18091})
	if callErr != nil {
		return nil, callErr
	}
	v55 = v51
	goto b6
b5:
	;
	v55 = vm.NIL
	goto b6
b6:
	;
	v170 = v55
	goto b3
b7:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("fns"), []vm.Value{template})
	if callErr != nil {
		return nil, callErr
	}
	arg__18121, callErr = rt.InvokeValue(vm.Keyword("fns"), []vm.Value{template})
	if callErr != nil {
		return nil, callErr
	}
	inners, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__18116 vm.Value
		var v7 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(vm.Keyword("fn"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__18116, callErr = rt.InvokeValue(vm.Keyword("fn"), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v7, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "lower-fn-lit").Deref(), []vm.Value{arg__18116, capture_exprs})
		if callErr != nil {
			return nil, callErr
		}
		return v7, nil
	}), arg__18121})
	if callErr != nil {
		return nil, callErr
	}
	v88, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("clojure.core", "identity").Deref(), inners})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v88) {
		goto b10
	} else {
		goto b11
	}
b8:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b13
	} else {
		goto b14
	}
b9:
	;
	v170 = v166
	goto b3
b10:
	;
	boxed, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__18170 vm.Value
		var arg__18186 vm.Value
		var arg__18188 vm.Value
		var arg__18191 vm.Value
		var v30 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
		if callErr != nil {
			return nil, callErr
		}
		arg__18170, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__18170, vm.String("BoxNativeFn")})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
		if callErr != nil {
			return nil, callErr
		}
		arg__18186, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
		if callErr != nil {
			return nil, callErr
		}
		arg__18188, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__18186, vm.String("BoxNativeFn")})
		if callErr != nil {
			return nil, callErr
		}
		arg__18191, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		v30, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__18188, arg__18191})
		if callErr != nil {
			return nil, callErr
		}
		return v30, nil
	}), inners})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18203, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__18203, vm.String("MakeNativeMultiArity")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18216, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18218, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__18216, boxed})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__18218})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18230, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("rt")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18232, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__18230, vm.String("MakeNativeMultiArity")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18243, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("[]vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18245, callErr = rt.InvokeValue(rt.LookupVar("gogen", "composite-lit").Deref(), []vm.Value{arg__18243, boxed})
	if callErr != nil {
		return nil, callErr
	}
	arg__18246, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__18245})
	if callErr != nil {
		return nil, callErr
	}
	v143, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__18232, arg__18246})
	if callErr != nil {
		return nil, callErr
	}
	v147 = v143
	goto b12
b11:
	;
	v147 = vm.NIL
	goto b12
b12:
	;
	v166 = v147
	goto b9
b13:
	;
	goto b15
b14:
	;
	goto b15
b15:
	;
	v166 = vm.NIL
	goto b9
}
func lower_terminator(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var term vm.Value
	var op vm.Value
	var refs vm.Value
	var aux vm.Value
	var needs_error_QMARK_ vm.Value
	var v30 bool
	var f vm.Value
	var closed_exprs vm.Value
	var arg__18276 vm.Value
	var v51 bool
	var v270 bool
	var v1075 vm.Value
	var ret_spec vm.Value
	var v74 bool
	var v243 vm.Value
	var arg__18298 vm.Value
	var v85 vm.Value
	var arg__18315 vm.Value
	var v96 vm.Value
	var expr vm.Value
	var v229 vm.Value
	var head__18316 vm.Value
	var arg__18322 vm.Value
	var head__18326 vm.Value
	var arg__18332 vm.Value
	var v206 vm.Value
	var v209 vm.Value
	var arg__18335 vm.Value
	var arg__18336 vm.Value
	var v225 vm.Value
	var v273 vm.Value
	var v292 bool
	var v1065 vm.Value
	var arg__18352 vm.Value
	var v313 bool
	var v627 bool
	var v1055 vm.Value
	var cond_nid vm.Value
	var arg__18370 vm.Value
	var cond_spec vm.Value
	var raw_cond vm.Value
	var v350 bool
	var v600 vm.Value
	var cond_expr vm.Value
	var arg__18423 vm.Value
	var true_stmts vm.Value
	var arg__18436 vm.Value
	var false_stmts vm.Value
	var arg__18389 vm.Value
	var arg__18405 vm.Value
	var arg__18407 vm.Value
	var arg__18410 vm.Value
	var v405 vm.Value
	var v409 vm.Value
	var arg__18447 vm.Value
	var v578 vm.Value
	var v582 vm.Value
	var and__x vm.Value
	var v556 vm.Value
	var v538 vm.Value
	var v648 vm.Value
	var v1045 vm.Value
	var tail_args vm.Value
	var arg__18456 vm.Value
	var v652 vm.Value
	var v655 vm.Value
	var fixed_arity vm.Value
	var i int
	var remaining vm.Value
	var out vm.Value
	var v703 vm.Value
	var arg__18475 vm.Value
	var val_expr vm.Value
	var assigns vm.Value
	var v796 vm.Value
	var arg__18496 vm.Value
	var v807 vm.Value
	var lhs vm.Value
	var v855 vm.Value
	var v773 bool
	var v776 vm.Value
	var v859 int
	var v861 vm.Value
	var v893 vm.Value
	var arg__18527 vm.Value
	var v905 vm.Value
	var v907 vm.Value
	var arg__18538 vm.Value
	var v993 vm.Value
	var v997 vm.Value
	var callErr error
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{arg2, arg0})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, arg0})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, arg0})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{term, arg0})
	if callErr != nil {
		return nil, callErr
	}
	needs_error_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-needs-error?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v30 = op == vm.Keyword("return")
	if v30 {
		f = arg0
		closed_exprs = arg1
		goto b1
	} else {
		f = arg0
		closed_exprs = arg1
		goto b2
	}
b1:
	;
	arg__18276, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v51 = arg__18276 == vm.Int(1)
	if v51 {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v270 = op == vm.Keyword("branch")
	if v270 {
		goto b19
	} else {
		goto b20
	}
b3:
	;
	return v1075, nil
b4:
	;
	ret_spec, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-return-spec").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v74 = ret_spec == vm.String("vm.Value")
	if v74 {
		goto b7
	} else {
		goto b8
	}
b5:
	;
	v243 = vm.NIL
	goto b6
b6:
	;
	v1075 = v243
	goto b3
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__18298, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v85, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "box-as-value").Deref(), []vm.Value{f, closed_exprs, arg__18298})
	if callErr != nil {
		return nil, callErr
	}
	expr = v85
	goto b9
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__18315, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v96, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-expr").Deref(), []vm.Value{f, closed_exprs, arg__18315})
	if callErr != nil {
		return nil, callErr
	}
	expr = v96
	goto b9
b9:
	;
	if vm.IsTruthy(expr) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	if vm.IsTruthy(needs_error_QMARK_) {
		head__18316 = rt.LookupVar("clojure.core", "vector").Deref()
		goto b13
	} else {
		head__18316 = rt.LookupVar("clojure.core", "vector").Deref()
		goto b14
	}
b11:
	;
	v229 = vm.NIL
	goto b12
b12:
	;
	v243 = v229
	goto b6
b13:
	;
	arg__18322, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr, arg__18322})
	if callErr != nil {
		return nil, callErr
	}
	goto b15
b14:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	goto b15
b15:
	;
	if vm.IsTruthy(needs_error_QMARK_) {
		head__18326 = rt.LookupVar("gogen", "return-stmt").Deref()
		goto b16
	} else {
		head__18326 = rt.LookupVar("gogen", "return-stmt").Deref()
		goto b17
	}
b16:
	;
	arg__18332, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("nil")})
	if callErr != nil {
		return nil, callErr
	}
	v206, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr, arg__18332})
	if callErr != nil {
		return nil, callErr
	}
	arg__18335 = v206
	goto b18
b17:
	;
	v209, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{expr})
	if callErr != nil {
		return nil, callErr
	}
	arg__18335 = v209
	goto b18
b18:
	;
	arg__18336, callErr = rt.InvokeValue(head__18326, []vm.Value{arg__18335})
	if callErr != nil {
		return nil, callErr
	}
	v225, callErr = rt.InvokeValue(head__18316, []vm.Value{arg__18336})
	if callErr != nil {
		return nil, callErr
	}
	v229 = v225
	goto b12
b19:
	;
	v273, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "transfer-stmts").Deref(), []vm.Value{f, closed_exprs, aux})
	if callErr != nil {
		return nil, callErr
	}
	v1065 = v273
	goto b21
b20:
	;
	v292 = op == vm.Keyword("branch-if")
	if v292 {
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v1075 = v1065
	goto b3
b22:
	;
	arg__18352, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v313 = arg__18352 == vm.Int(1)
	if v313 {
		goto b25
	} else {
		goto b26
	}
b23:
	;
	v627 = op == vm.Keyword("tail-call")
	if v627 {
		goto b43
	} else {
		goto b44
	}
b24:
	;
	v1065 = v1055
	goto b21
b25:
	;
	cond_nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{cond_nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__18370, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{cond_nid, f})
	if callErr != nil {
		return nil, callErr
	}
	cond_spec, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__18370})
	if callErr != nil {
		return nil, callErr
	}
	raw_cond, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "value-expr").Deref(), []vm.Value{f, closed_exprs, cond_nid})
	if callErr != nil {
		return nil, callErr
	}
	v350 = cond_spec == vm.String("bool")
	if v350 {
		goto b28
	} else {
		goto b29
	}
b26:
	;
	v600 = vm.NIL
	goto b27
b27:
	;
	v1055 = v600
	goto b24
b28:
	;
	cond_expr = raw_cond
	goto b30
b29:
	;
	if vm.IsTruthy(raw_cond) {
		goto b31
	} else {
		goto b32
	}
b30:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__18423, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	true_stmts, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "transfer-stmts").Deref(), []vm.Value{f, closed_exprs, arg__18423})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__18436, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	false_stmts, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "transfer-stmts").Deref(), []vm.Value{f, closed_exprs, arg__18436})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(cond_expr) {
		goto b37
	} else {
		and__x = cond_expr
		goto b38
	}
b31:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18389, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__18389, vm.String("IsTruthy")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{raw_cond})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18405, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18407, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__18405, vm.String("IsTruthy")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18410, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{raw_cond})
	if callErr != nil {
		return nil, callErr
	}
	v405, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__18407, arg__18410})
	if callErr != nil {
		return nil, callErr
	}
	v409 = v405
	goto b33
b32:
	;
	v409 = vm.NIL
	goto b33
b33:
	;
	cond_expr = v409
	goto b30
b34:
	;
	arg__18447, callErr = rt.InvokeValue(rt.LookupVar("gogen", "if-stmt").Deref(), []vm.Value{vm.NIL, cond_expr, true_stmts, false_stmts})
	if callErr != nil {
		return nil, callErr
	}
	v578, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__18447})
	if callErr != nil {
		return nil, callErr
	}
	v582 = v578
	goto b36
b35:
	;
	v582 = vm.NIL
	goto b36
b36:
	;
	v600 = v582
	goto b27
b37:
	;
	if vm.IsTruthy(true_stmts) {
		goto b40
	} else {
		and__x = true_stmts
		goto b41
	}
b38:
	;
	v556 = and__x
	goto b39
b39:
	;
	if vm.IsTruthy(v556) {
		goto b34
	} else {
		goto b35
	}
b40:
	;
	v538 = false_stmts
	goto b42
b41:
	;
	v538 = and__x
	goto b42
b42:
	;
	v556 = v538
	goto b39
b43:
	;
	v648, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v648) {
		tail_args = refs
		goto b46
	} else {
		tail_args = refs
		goto b47
	}
b44:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b68
	} else {
		goto b69
	}
b45:
	;
	v1055 = v1045
	goto b24
b46:
	;
	arg__18456, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arity").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v652 = rt.SubValue(arg__18456, vm.Int(1))
	fixed_arity = v652
	goto b48
b47:
	;
	v655, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arity").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	fixed_arity = v655
	goto b48
b48:
	;
	i = 0
	remaining = tail_args
	out = vm.NewArrayVector([]vm.Value{})
	goto b49
b49:
	;
	v703, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v703) {
		goto b50
	} else {
		goto b51
	}
b50:
	;
	assigns = out
	goto b52
b51:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__18475, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	val_expr, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "box-as-value").Deref(), []vm.Value{f, closed_exprs, arg__18475})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b56
	} else {
		goto b57
	}
b52:
	;
	if vm.IsTruthy(assigns) {
		goto b65
	} else {
		goto b66
	}
b53:
	;
	v796, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("args0")})
	if callErr != nil {
		return nil, callErr
	}
	lhs = v796
	goto b55
b54:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("a"), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__18496, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("a"), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v807, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{arg__18496})
	if callErr != nil {
		return nil, callErr
	}
	lhs = v807
	goto b55
b55:
	;
	v855, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{val_expr})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v855) {
		goto b59
	} else {
		goto b60
	}
b56:
	;
	v773 = vm.Int(i) == fixed_arity
	v776 = vm.Boolean(v773)
	goto b58
b57:
	;
	v776 = and__x
	goto b58
b58:
	;
	if vm.IsTruthy(v776) {
		goto b53
	} else {
		goto b54
	}
b59:
	;
	goto b61
b60:
	;
	v859 = i + 1
	v861, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	v893, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "same-ident?").Deref(), []vm.Value{lhs, val_expr})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v893) {
		goto b62
	} else {
		goto b63
	}
b61:
	;
	assigns = vm.NIL
	goto b52
b62:
	;
	v907 = out
	goto b64
b63:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), lhs, val_expr})
	if callErr != nil {
		return nil, callErr
	}
	arg__18527, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), lhs, val_expr})
	if callErr != nil {
		return nil, callErr
	}
	v905, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, arg__18527})
	if callErr != nil {
		return nil, callErr
	}
	v907 = v905
	goto b64
b64:
	;
	i = v859
	remaining = v861
	out = v907
	goto b49
b65:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "goto-stmt").Deref(), []vm.Value{vm.String("func_entry")})
	if callErr != nil {
		return nil, callErr
	}
	arg__18538, callErr = rt.InvokeValue(rt.LookupVar("gogen", "goto-stmt").Deref(), []vm.Value{vm.String("func_entry")})
	if callErr != nil {
		return nil, callErr
	}
	v993, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{assigns, arg__18538})
	if callErr != nil {
		return nil, callErr
	}
	v997 = v993
	goto b67
b66:
	;
	v997 = vm.NIL
	goto b67
b67:
	;
	v1045 = v997
	goto b45
b68:
	;
	goto b70
b69:
	;
	goto b70
b70:
	;
	v1045 = vm.NIL
	goto b45
}
func multi_fn_template_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var aux vm.Value
	var arg__18544 vm.Value
	var v11 bool
	var v14 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		aux = arg0
		goto b1
	} else {
		goto b2
	}
b1:
	;
	arg__18544, callErr = rt.InvokeValue(vm.Keyword("kind"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v11 = arg__18544 == vm.Keyword("multi-fn-template")
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
func nested_template_fns(arg0 vm.Value) (vm.Value, error) {
	var v5 vm.Value
	var blocks vm.Value
	var out vm.Value
	var f vm.Value
	var v15 vm.Value
	var arg__18561 vm.Value
	var insts vm.Value
	var found vm.Value
	var v35 vm.Value
	var v37 vm.Value
	var v39 vm.Value
	var callErr error
	v5, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v5
	out = vm.NewArrayVector([]vm.Value{})
	f = arg0
	goto b1
b1:
	;
	v15, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v15) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v39 = out
	goto b4
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__18561, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg__18561, f})
	if callErr != nil {
		return nil, callErr
	}
	found, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var aux vm.Value
		var v14 vm.Value
		var acc vm.Value
		var arg__18604 vm.Value
		var v21 vm.Value
		var v24 vm.Value
		var callErr error
		aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg1, f})
		if callErr != nil {
			return nil, callErr
		}
		v14, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "any-fn-template?").Deref(), []vm.Value{aux})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(v14) {
			acc = arg0
			goto b1
		} else {
			acc = arg0
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "template-fns").Deref(), []vm.Value{aux})
		if callErr != nil {
			return nil, callErr
		}
		arg__18604, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "template-fns").Deref(), []vm.Value{aux})
		if callErr != nil {
			return nil, callErr
		}
		v21, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{acc, arg__18604})
		if callErr != nil {
			return nil, callErr
		}
		v24 = v21
		goto b3
	b2:
		;
		v24 = acc
		goto b3
	b3:
		;
		return v24, nil
	}), vm.NewArrayVector([]vm.Value{}), insts})
	if callErr != nil {
		return nil, callErr
	}
	v35, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	v37, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{out, found})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v35
	out = v37
	goto b1
b4:
	;
	return v39, nil
}
func oos_all_nids(arg0 vm.Value) (vm.Value, error) {
	var acc vm.Value
	var arg__18633 vm.Value
	var doseq_seq__18616 vm.Value
	var doseq_loop__18617 vm.Value
	var f vm.Value
	var bid vm.Value
	var arg__18649 vm.Value
	var doseq_seq__18618 vm.Value
	var arg__18700 vm.Value
	var v130 vm.Value
	var doseq_loop__18619 vm.Value
	var p vm.Value
	var v57 vm.Value
	var arg__18675 vm.Value
	var doseq_seq__18620 vm.Value
	var doseq_loop__18621 vm.Value
	var nid vm.Value
	var v102 vm.Value
	var v116 vm.Value
	var callErr error
	acc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NewArrayVector([]vm.Value{})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__18633, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__18616, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__18633})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18617 = doseq_seq__18616
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__18617) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__18617})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__18649, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__18618, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__18649})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18619 = doseq_seq__18618
	goto b5
b3:
	;
	goto b4
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{acc})
	if callErr != nil {
		return nil, callErr
	}
	arg__18700, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{acc})
	if callErr != nil {
		return nil, callErr
	}
	v130, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "distinct").Deref(), []vm.Value{arg__18700})
	if callErr != nil {
		return nil, callErr
	}
	return v130, nil
b5:
	;
	if vm.IsTruthy(doseq_loop__18619) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__18619})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{acc, rt.LookupVar("clojure.core", "conj").Deref(), p})
	if callErr != nil {
		return nil, callErr
	}
	v57, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__18619})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18619 = v57
	goto b5
b7:
	;
	goto b8
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__18675, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__18620, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__18675})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18621 = doseq_seq__18620
	goto b9
b9:
	;
	if vm.IsTruthy(doseq_loop__18621) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__18621})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{acc, rt.LookupVar("clojure.core", "conj").Deref(), nid})
	if callErr != nil {
		return nil, callErr
	}
	v102, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__18621})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18621 = v102
	goto b9
b11:
	;
	goto b12
b12:
	;
	v116, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__18617})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18617 = v116
	goto b1
}
func oos_branch_targets(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var op vm.Value
	var aux vm.Value
	var v15 bool
	var v18 vm.Value
	var v29 bool
	var v64 vm.Value
	var arg__18721 vm.Value
	var arg__18725 vm.Value
	var v36 vm.Value
	var v58 vm.Value
	var callErr error
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v15 = op == vm.Keyword("branch")
	if v15 {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v64 = v18
	goto b3
b2:
	;
	v29 = op == vm.Keyword("branch-if")
	if v29 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v64, nil
b4:
	;
	arg__18721, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__18725, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v36, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__18721, arg__18725})
	if callErr != nil {
		return nil, callErr
	}
	v58 = v36
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
	v64 = v58
	goto b3
b7:
	;
	goto b9
b8:
	;
	goto b9
b9:
	;
	v58 = vm.NewArrayVector([]vm.Value{})
	goto b6
}
func oos_captured_names(arg0 vm.Value) (vm.Value, error) {
	var arg__18736 vm.Value
	var acc vm.Value
	var arg__18745 vm.Value
	var doseq_seq__18728 vm.Value
	var doseq_loop__18729 vm.Value
	var f vm.Value
	var bid vm.Value
	var arg__18761 vm.Value
	var doseq_seq__18730 vm.Value
	var v192 vm.Value
	var doseq_loop__18731 vm.Value
	var nid vm.Value
	var arg__18771 vm.Value
	var v74 bool
	var v182 vm.Value
	var s vm.Value
	var and__x vm.Value
	var v169 vm.Value
	var arg__18788 vm.Value
	var v124 vm.Value
	var v127 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__18736, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	acc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{arg__18736})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__18745, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__18728, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__18745})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18729 = doseq_seq__18728
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__18729) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__18729})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__18761, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__18730, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__18761})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18731 = doseq_seq__18730
	goto b5
b3:
	;
	goto b4
b4:
	;
	v192, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{acc})
	if callErr != nil {
		return nil, callErr
	}
	return v192, nil
b5:
	;
	if vm.IsTruthy(doseq_loop__18731) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__18731})
	if callErr != nil {
		return nil, callErr
	}
	arg__18771, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v74 = arg__18771 == vm.Keyword("load-closed")
	if v74 {
		goto b9
	} else {
		goto b10
	}
b7:
	;
	goto b8
b8:
	;
	v182, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__18729})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18729 = v182
	goto b1
b9:
	;
	s, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "source-name-of").Deref(), []vm.Value{f, nid})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{s})
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
	v169, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__18731})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18731 = v169
	goto b5
b12:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{acc, rt.LookupVar("clojure.core", "conj").Deref(), s})
	if callErr != nil {
		return nil, callErr
	}
	goto b14
b13:
	;
	goto b14
b14:
	;
	goto b11
b15:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{s})
	if callErr != nil {
		return nil, callErr
	}
	arg__18788, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{s})
	if callErr != nil {
		return nil, callErr
	}
	v124, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__18788})
	if callErr != nil {
		return nil, callErr
	}
	v127 = v124
	goto b17
b16:
	;
	v127 = and__x
	goto b17
b17:
	;
	if vm.IsTruthy(v127) {
		goto b12
	} else {
		goto b13
	}
}
func oos_next_tmp_BANG_(arg0 vm.Value) (vm.Value, error) {
	var arg__18809 vm.Value
	var or__x vm.Value
	var f vm.Value
	var n vm.Value
	var arg__18814 vm.Value
	var v27 vm.Value
	var callErr error
	arg__18809, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(vm.Keyword("pc-tmp"), []vm.Value{arg__18809})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		f = arg0
		goto b1
	} else {
		f = arg0
		goto b2
	}
b1:
	;
	n = or__x
	goto b3
b2:
	;
	n = vm.Int(0)
	goto b3
b3:
	;
	arg__18814 = rt.AddValue(n, vm.Int(1))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{f, rt.LookupVar("clojure.core", "assoc").Deref(), vm.Keyword("pc-tmp"), arg__18814})
	if callErr != nil {
		return nil, callErr
	}
	v27, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("_pc"), n})
	if callErr != nil {
		return nil, callErr
	}
	return v27, nil
}
func oos_param_incoming_args(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var acc vm.Value
	var arg__18845 vm.Value
	var doseq_seq__18826 vm.Value
	var doseq_loop__18827 vm.Value
	var i vm.Value
	var b vm.Value
	var f vm.Value
	var bp vm.Value
	var term vm.Value
	var v223 vm.Value
	var arg__18866 vm.Value
	var doseq_seq__18828 vm.Value
	var v211 vm.Value
	var doseq_loop__18829 vm.Value
	var bt vm.Value
	var arg__18873 vm.Value
	var v111 bool
	var args vm.Value
	var arg__18882 vm.Value
	var v141 bool
	var v183 vm.Value
	var arg__18899 vm.Value
	var callErr error
	acc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NewArrayVector([]vm.Value{})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__18845, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__18826, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__18845})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18827 = doseq_seq__18826
	i = arg2
	b = arg1
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__18827) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bp, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__18827})
	if callErr != nil {
		return nil, callErr
	}
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{bp, f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(term) {
		goto b5
	} else {
		goto b6
	}
b3:
	;
	goto b4
b4:
	;
	v223, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{acc})
	if callErr != nil {
		return nil, callErr
	}
	return v223, nil
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-branch-targets").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__18866, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-branch-targets").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__18828, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__18866})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18829 = doseq_seq__18828
	goto b8
b6:
	;
	goto b7
b7:
	;
	v211, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__18827})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18827 = v211
	goto b1
b8:
	;
	if vm.IsTruthy(doseq_loop__18829) {
		goto b9
	} else {
		goto b10
	}
b9:
	;
	bt, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__18829})
	if callErr != nil {
		return nil, callErr
	}
	arg__18873, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	v111 = arg__18873 == b
	if v111 {
		goto b12
	} else {
		goto b13
	}
b10:
	;
	goto b11
b11:
	;
	goto b7
b12:
	;
	args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	arg__18882, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	v141 = rt.LtValue(i, arg__18882)
	if v141 {
		goto b15
	} else {
		goto b16
	}
b13:
	;
	goto b14
b14:
	;
	v183, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__18829})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__18829 = v183
	goto b8
b15:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, i})
	if callErr != nil {
		return nil, callErr
	}
	arg__18899, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, i})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{acc, rt.LookupVar("clojure.core", "conj").Deref(), arg__18899})
	if callErr != nil {
		return nil, callErr
	}
	goto b17
b16:
	;
	goto b17
b17:
	;
	goto b14
}
func oos_rhs_read_name(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var rhs vm.Value
	var v7 vm.Value
	var v11 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v4) {
		rhs = arg0
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v7, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident-name").Deref(), []vm.Value{rhs})
	if callErr != nil {
		return nil, callErr
	}
	v11 = v7
	goto b3
b2:
	;
	v11 = vm.NIL
	goto b3
b3:
	;
	return v11, nil
}
func oos_sequence_copies(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__18918 vm.Value
	var v9 bool
	var pairs vm.Value
	var v14 vm.Value
	var f vm.Value
	var v244 vm.Value
	var pending vm.Value
	var out vm.Value
	var v30 vm.Value
	var arg__19025 vm.Value
	var reads vm.Value
	var v238 vm.Value
	var i int
	var arg__19031 vm.Value
	var v60 bool
	var arg__19046 vm.Value
	var arg__19063 vm.Value
	var arg__19065 vm.Value
	var arg__19082 vm.Value
	var arg__19099 vm.Value
	var arg__19101 vm.Value
	var arg__19102 vm.Value
	var v111 vm.Value
	var ready int
	var v148 bool
	var v113 int
	var v116 int
	var p vm.Value
	var v153 vm.Value
	var arg__19137 vm.Value
	var arg__19143 vm.Value
	var arg__19167 vm.Value
	var arg__19173 vm.Value
	var arg__19174 vm.Value
	var v195 vm.Value
	var arg__19192 vm.Value
	var dname vm.Value
	var tname vm.Value
	var arg__19210 vm.Value
	var arg__19214 vm.Value
	var save vm.Value
	var pending_STAR_ vm.Value
	var v236 vm.Value
	var callErr error
	arg__18918, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	v9 = rt.LeValue(arg__18918, vm.Int(1))
	if v9 {
		pairs = arg1
		goto b1
	} else {
		f = arg0
		pairs = arg1
		goto b2
	}
b1:
	;
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__18970 vm.Value
		var arg__18976 vm.Value
		var v20 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		arg__18970, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		arg__18976, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		v20, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__18970, arg__18976})
		if callErr != nil {
			return nil, callErr
		}
		return v20, nil
	}), pairs})
	if callErr != nil {
		return nil, callErr
	}
	v244 = v14
	goto b3
b2:
	;
	pending = pairs
	out = vm.NewArrayVector([]vm.Value{})
	goto b4
b3:
	;
	return v244, nil
b4:
	;
	v30, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{pending})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v30) {
		goto b5
	} else {
		goto b6
	}
b5:
	;
	v238 = out
	goto b7
b6:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__19025, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	reads, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
		var arg__19017 vm.Value
		var rd vm.Value
		var s vm.Value
		var v20 vm.Value
		var v23 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		arg__19017, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		rd, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-rhs-read-name").Deref(), []vm.Value{arg__19017})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(rd) {
			s = arg0
			goto b1
		} else {
			s = arg0
			goto b2
		}
	b1:
		;
		v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{s, rd})
		if callErr != nil {
			return nil, callErr
		}
		v23 = v20
		goto b3
	b2:
		;
		v23 = s
		goto b3
	b3:
		;
		return v23, nil
	}), arg__19025, pending})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	goto b8
b7:
	;
	v244 = v238
	goto b3
b8:
	;
	arg__19031, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{pending})
	if callErr != nil {
		return nil, callErr
	}
	v60 = rt.LtValue(vm.Int(i), arg__19031)
	if v60 {
		goto b9
	} else {
		goto b10
	}
b9:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19046, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__19046, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19063, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19065, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__19063, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident-name").Deref(), []vm.Value{arg__19065})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19082, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__19082, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19099, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19101, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__19099, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19102, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident-name").Deref(), []vm.Value{arg__19101})
	if callErr != nil {
		return nil, callErr
	}
	v111, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{reads, arg__19102})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v111) {
		goto b12
	} else {
		goto b13
	}
b10:
	;
	ready = -1
	goto b11
b11:
	;
	v148 = ready >= 0
	if v148 {
		goto b15
	} else {
		goto b16
	}
b12:
	;
	v113 = i + 1
	i = v113
	goto b8
b13:
	;
	v116 = i
	goto b14
b14:
	;
	ready = v116
	goto b11
b15:
	;
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(ready)})
	if callErr != nil {
		return nil, callErr
	}
	v153, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-vec-remove").Deref(), []vm.Value{pending, vm.Int(ready)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19137, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19143, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__19137, arg__19143})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19167, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19173, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19174, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String("="), arg__19167, arg__19173})
	if callErr != nil {
		return nil, callErr
	}
	v195, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, arg__19174})
	if callErr != nil {
		return nil, callErr
	}
	pending = v153
	out = v195
	goto b4
b16:
	;
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{pending, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19192, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{p, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	dname, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident-name").Deref(), []vm.Value{arg__19192})
	if callErr != nil {
		return nil, callErr
	}
	tname, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "oos-next-tmp!").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{tname})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{dname})
	if callErr != nil {
		return nil, callErr
	}
	arg__19210, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{tname})
	if callErr != nil {
		return nil, callErr
	}
	arg__19214, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{dname})
	if callErr != nil {
		return nil, callErr
	}
	save, callErr = rt.InvokeValue(rt.LookupVar("gogen", "assign").Deref(), []vm.Value{vm.String(":="), arg__19210, arg__19214})
	if callErr != nil {
		return nil, callErr
	}
	pending_STAR_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var r vm.Value
		var and__x vm.Value
		var q vm.Value
		var tname_9 vm.Value
		var arg__19261 vm.Value
		var arg__19265 vm.Value
		var v47 vm.Value
		var v50 vm.Value
		var dname_18 vm.Value
		var tname_19 vm.Value
		var arg__19254 vm.Value
		var v30 vm.Value
		var tname_24 vm.Value
		var v33 vm.Value
		var tname_36 vm.Value
		var callErr error
		r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg0, vm.Int(1)})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident?").Deref(), []vm.Value{r})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			q = arg0
			dname_18 = dname
			tname_19 = tname
			goto b4
		} else {
			q = arg0
			tname_24 = tname
			goto b5
		}
	b1:
		;
		arg__19261, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{q, vm.Int(0)})
		if callErr != nil {
			return nil, callErr
		}
		arg__19265, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{tname_9})
		if callErr != nil {
			return nil, callErr
		}
		v47, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__19261, arg__19265})
		if callErr != nil {
			return nil, callErr
		}
		v50 = v47
		goto b3
	b2:
		;
		v50 = q
		goto b3
	b3:
		;
		return v50, nil
	b4:
		;
		arg__19254, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident-name").Deref(), []vm.Value{r})
		if callErr != nil {
			return nil, callErr
		}
		v30 = vm.Boolean(dname_18 == arg__19254)
		v33 = v30
		tname_36 = tname_19
		goto b6
	b5:
		;
		v33 = and__x
		tname_36 = tname_24
		goto b6
	b6:
		;
		if vm.IsTruthy(v33) {
			tname_9 = tname_36
			goto b1
		} else {
			goto b2
		}
	}), pending})
	if callErr != nil {
		return nil, callErr
	}
	v236, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, save})
	if callErr != nil {
		return nil, callErr
	}
	pending = pending_STAR_
	out = v236
	goto b4
}
func oos_unsafe_names(arg0 vm.Value) (vm.Value, error) {
	var arg__19285 vm.Value
	var unsafe vm.Value
	var arg__19294 vm.Value
	var doseq_seq__19275 vm.Value
	var doseq_loop__19276 vm.Value
	var f vm.Value
	var bid vm.Value
	var arg__19310 vm.Value
	var doseq_seq__19277 vm.Value
	var v348 vm.Value
	var doseq_loop__19278 vm.Value
	var nid vm.Value
	var seen vm.Value
	var arg__19329 vm.Value
	var doseq_seq__19279 vm.Value
	var v338 vm.Value
	var doseq_loop__19280 vm.Value
	var r vm.Value
	var s vm.Value
	var and__x vm.Value
	var v325 vm.Value
	var arg__19359 vm.Value
	var tem__G__0 vm.Value
	var v308 vm.Value
	var arg__19349 vm.Value
	var v159 vm.Value
	var v162 vm.Value
	var v244 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	arg__19285, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	unsafe, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{arg__19285})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19294, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__19275, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__19294})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__19276 = doseq_seq__19275
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__19276) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__19276})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__19310, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__19277, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__19310})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__19278 = doseq_seq__19277
	goto b5
b3:
	;
	goto b4
b4:
	;
	v348, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{unsafe})
	if callErr != nil {
		return nil, callErr
	}
	return v348, nil
b5:
	;
	if vm.IsTruthy(doseq_loop__19278) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__19278})
	if callErr != nil {
		return nil, callErr
	}
	seen, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__19329, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__19279, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__19329})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__19280 = doseq_seq__19279
	goto b9
b7:
	;
	goto b8
b8:
	;
	v338, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__19276})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__19276 = v338
	goto b1
b9:
	;
	if vm.IsTruthy(doseq_loop__19280) {
		goto b10
	} else {
		goto b11
	}
b10:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__19280})
	if callErr != nil {
		return nil, callErr
	}
	s, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "source-name-of").Deref(), []vm.Value{f, r})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{s})
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
	v325, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__19278})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__19278 = v325
	goto b5
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{seen})
	if callErr != nil {
		return nil, callErr
	}
	arg__19359, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{seen})
	if callErr != nil {
		return nil, callErr
	}
	tem__G__0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__19359, s})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(tem__G__0) {
		goto b19
	} else {
		goto b20
	}
b14:
	;
	goto b15
b15:
	;
	v308, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__19280})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__19280 = v308
	goto b9
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{s})
	if callErr != nil {
		return nil, callErr
	}
	arg__19349, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{s})
	if callErr != nil {
		return nil, callErr
	}
	v159, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__19349})
	if callErr != nil {
		return nil, callErr
	}
	v162 = v159
	goto b18
b17:
	;
	v162 = and__x
	goto b18
b18:
	;
	if vm.IsTruthy(v162) {
		goto b13
	} else {
		goto b14
	}
b19:
	;
	v244, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{tem__G__0, r})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v244) {
		goto b22
	} else {
		goto b23
	}
b20:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{seen, rt.LookupVar("clojure.core", "assoc").Deref(), s, r})
	if callErr != nil {
		return nil, callErr
	}
	goto b21
b21:
	;
	goto b15
b22:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{unsafe, rt.LookupVar("clojure.core", "conj").Deref(), s})
	if callErr != nil {
		return nil, callErr
	}
	goto b24
b23:
	;
	goto b24
b24:
	;
	goto b21
}
func oos_vec_remove(arg0 vm.Value, arg1 int) (vm.Value, error) {
	var arg__19401 int
	var arg__19414 vm.Value
	var arg__19422 vm.Value
	var arg__19445 vm.Value
	var arg__19453 vm.Value
	var arg__19454 vm.Value
	var v31 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "take").Deref(), []vm.Value{vm.Int(arg1), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19401 = arg1 + 1
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(arg__19401), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19414, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "take").Deref(), []vm.Value{vm.Int(arg1), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19422, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(arg__19401), arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{arg__19414, arg__19422})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "take").Deref(), []vm.Value{vm.Int(arg1), arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(arg__19401), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19445, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "take").Deref(), []vm.Value{vm.Int(arg1), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19453, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "drop").Deref(), []vm.Value{vm.Int(arg__19401), arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19454, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat").Deref(), []vm.Value{arg__19445, arg__19453})
	if callErr != nil {
		return nil, callErr
	}
	v31, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__19454})
	if callErr != nil {
		return nil, callErr
	}
	return v31, nil
}
func override_uniform_value_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var arg__19463 vm.Value
	var and__x_6 vm.Value
	var f vm.Value
	var arg__19468 vm.Value
	var and__x_15 bool
	var and__x_10 vm.Value
	var v36 vm.Value
	var arg__19483 vm.Value
	var v28 vm.Value
	var and__x_19 bool
	var v31 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19463, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	and__x_6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__19463})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x_6) {
		f = arg0
		goto b1
	} else {
		and__x_10 = and__x_6
		goto b2
	}
b1:
	;
	arg__19468, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-return-spec").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	and__x_15 = arg__19468 == vm.String("vm.Value")
	if and__x_15 {
		goto b4
	} else {
		and__x_19 = and__x_15
		goto b5
	}
b2:
	;
	v36 = and__x_10
	goto b3
b3:
	;
	return v36, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-param-specs").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__19483, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-param-specs").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var v2 vm.Value
		v2 = vm.Boolean(vm.String("vm.Value") == arg0)
		return v2
	}), arg__19483})
	if callErr != nil {
		return nil, callErr
	}
	v31 = v28
	goto b6
b5:
	;
	v31 = vm.Boolean(and__x_19)
	goto b6
b6:
	;
	v36 = v31
	goto b3
}
func params_for(arg0 vm.Value) (vm.Value, error) {
	var arity vm.Value
	var variadic_QMARK_ vm.Value
	var f vm.Value
	var v12 vm.Value
	var fixed_arity vm.Value
	var or__x vm.Value
	var arg_types vm.Value
	var i int
	var out vm.Value
	var v64 bool
	var load_arg_t vm.Value
	var arg__19532 vm.Value
	var v138 bool
	var v288 vm.Value
	var arg__19507 vm.Value
	var arg__19521 vm.Value
	var arg__19522 vm.Value
	var v106 vm.Value
	var v109 vm.Value
	var v141 vm.Value
	var meta_t vm.Value
	var v175 vm.Value
	var t vm.Value
	var go_type vm.Value
	var v247 vm.Value
	var v201 vm.Value
	var v251 int
	var arg__19564 vm.Value
	var arg__19582 vm.Value
	var arg__19584 vm.Value
	var v273 vm.Value
	var callErr error
	arity, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arity").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	variadic_QMARK_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-variadic?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(variadic_QMARK_) {
		f = arg0
		goto b1
	} else {
		f = arg0
		goto b2
	}
b1:
	;
	v12 = rt.SubValue(arity, vm.Int(1))
	fixed_arity = v12
	goto b3
b2:
	;
	fixed_arity = arity
	goto b3
b3:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-arg-types").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	arg_types = or__x
	goto b6
b5:
	;
	arg_types = vm.NewArrayVector([]vm.Value{})
	goto b6
b6:
	;
	i = 0
	out = vm.NewArrayVector([]vm.Value{})
	goto b7
b7:
	;
	v64 = rt.GeValue(vm.Int(i), fixed_arity)
	if v64 {
		goto b8
	} else {
		goto b9
	}
b8:
	;
	if vm.IsTruthy(variadic_QMARK_) {
		goto b11
	} else {
		goto b12
	}
b9:
	;
	load_arg_t, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-load-arg-type").Deref(), []vm.Value{f, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19532, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg_types})
	if callErr != nil {
		return nil, callErr
	}
	v138 = rt.LtValue(vm.Int(i), arg__19532)
	if v138 {
		goto b14
	} else {
		goto b15
	}
b10:
	;
	return v288, nil
b11:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__19507, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "variadic-param").Deref(), []vm.Value{vm.String("args"), arg__19507})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__19521, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("vm.Value")})
	if callErr != nil {
		return nil, callErr
	}
	arg__19522, callErr = rt.InvokeValue(rt.LookupVar("gogen", "variadic-param").Deref(), []vm.Value{vm.String("args"), arg__19521})
	if callErr != nil {
		return nil, callErr
	}
	v106, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, arg__19522})
	if callErr != nil {
		return nil, callErr
	}
	v109 = v106
	goto b13
b12:
	;
	v109 = out
	goto b13
b13:
	;
	v288 = v109
	goto b10
b14:
	;
	v141, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg_types, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	meta_t = v141
	goto b16
b15:
	;
	meta_t = vm.Keyword("unknown")
	goto b16
b16:
	;
	v175, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("unknown"), load_arg_t})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v175) {
		goto b17
	} else {
		goto b18
	}
b17:
	;
	t = load_arg_t
	goto b19
b18:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b20
	} else {
		goto b21
	}
b19:
	;
	go_type, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "infer-go-type").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	v247, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{go_type})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v247) {
		goto b23
	} else {
		goto b24
	}
b20:
	;
	v201 = meta_t
	goto b22
b21:
	;
	v201 = vm.NIL
	goto b22
b22:
	;
	t = v201
	goto b19
b23:
	;
	goto b25
b24:
	;
	v251 = i + 1
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("arg"), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19564, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("arg"), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "param").Deref(), []vm.Value{arg__19564, go_type})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("arg"), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19582, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("arg"), vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19584, callErr = rt.InvokeValue(rt.LookupVar("gogen", "param").Deref(), []vm.Value{arg__19582, go_type})
	if callErr != nil {
		return nil, callErr
	}
	v273, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, arg__19584})
	if callErr != nil {
		return nil, callErr
	}
	i = v251
	out = v273
	goto b7
b25:
	;
	v288 = vm.NIL
	goto b10
}
func result_node(arg0 vm.Value) (vm.Value, error) {
	var ret_ids vm.Value
	var f vm.Value
	var specs vm.Value
	var and__x vm.Value
	var v136 vm.Value
	var arg__19596 vm.Value
	var v19 vm.Value
	var v22 vm.Value
	var arg__19653 vm.Value
	var arg__19664 vm.Value
	var arg__19665 vm.Value
	var arg__19666 vm.Value
	var results vm.Value
	var v95 vm.Value
	var v129 vm.Value
	var arg__19642 vm.Value
	var arg__19643 vm.Value
	var v61 bool
	var v64 vm.Value
	var arg__19679 vm.Value
	var arg__19691 vm.Value
	var arg__19692 vm.Value
	var v118 vm.Value
	var v121 vm.Value
	var callErr error
	ret_ids, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "return-ref-ids").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(ret_ids) {
		f = arg0
		goto b4
	} else {
		f = arg0
		and__x = ret_ids
		goto b5
	}
b1:
	;
	specs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__19625 vm.Value
		var v7 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		arg__19625, callErr = rt.InvokeValue(rt.LookupVar("ir", "type-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		v7, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "go-type-spec").Deref(), []vm.Value{arg__19625})
		if callErr != nil {
			return nil, callErr
		}
		return v7, nil
	}), ret_ids})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.LookupVar("clojure.core", "identity").Deref(), specs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b10
	} else {
		goto b11
	}
b2:
	;
	v136 = vm.NIL
	goto b3
b3:
	;
	return v136, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{ret_ids})
	if callErr != nil {
		return nil, callErr
	}
	arg__19596, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{ret_ids})
	if callErr != nil {
		return nil, callErr
	}
	v19, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__19596})
	if callErr != nil {
		return nil, callErr
	}
	v22 = v19
	goto b6
b5:
	;
	v22 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v22) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	arg__19653, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{arg__19653})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	arg__19664, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	arg__19665, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{arg__19664})
	if callErr != nil {
		return nil, callErr
	}
	arg__19666, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__19665})
	if callErr != nil {
		return nil, callErr
	}
	results, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__19666})
	if callErr != nil {
		return nil, callErr
	}
	v95, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "function-needs-error?").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v95) {
		goto b13
	} else {
		goto b14
	}
b8:
	;
	v129 = vm.NIL
	goto b9
b9:
	;
	v136 = v129
	goto b3
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "distinct").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	arg__19642, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "distinct").Deref(), []vm.Value{specs})
	if callErr != nil {
		return nil, callErr
	}
	arg__19643, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__19642})
	if callErr != nil {
		return nil, callErr
	}
	v61 = arg__19643 == vm.Int(1)
	v64 = vm.Boolean(v61)
	goto b12
b11:
	;
	v64 = and__x
	goto b12
b12:
	;
	if vm.IsTruthy(v64) {
		goto b7
	} else {
		goto b8
	}
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("error")})
	if callErr != nil {
		return nil, callErr
	}
	arg__19679, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("error")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__19679})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("error")})
	if callErr != nil {
		return nil, callErr
	}
	arg__19691, callErr = rt.InvokeValue(rt.LookupVar("gogen", "type").Deref(), []vm.Value{vm.String("error")})
	if callErr != nil {
		return nil, callErr
	}
	arg__19692, callErr = rt.InvokeValue(rt.LookupVar("gogen", "result").Deref(), []vm.Value{arg__19691})
	if callErr != nil {
		return nil, callErr
	}
	v118, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{results, arg__19692})
	if callErr != nil {
		return nil, callErr
	}
	v121 = v118
	goto b15
b14:
	;
	v121 = results
	goto b15
b15:
	;
	v129 = v121
	goto b9
}
func return_ref_ids(arg0 vm.Value) (vm.Value, error) {
	var v5 vm.Value
	var blocks vm.Value
	var out vm.Value
	var f vm.Value
	var v15 vm.Value
	var arg__19708 vm.Value
	var term vm.Value
	var v33 vm.Value
	var v123 vm.Value
	var v36 vm.Value
	var arg__19721 vm.Value
	var v49 bool
	var refs vm.Value
	var arg__19732 vm.Value
	var v66 bool
	var v69 vm.Value
	var arg__19750 vm.Value
	var v79 vm.Value
	var v101 vm.Value
	var callErr error
	v5, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v5
	out = vm.NewArrayVector([]vm.Value{})
	f = arg0
	goto b1
b1:
	;
	v15, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v15) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v123 = out
	goto b4
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	arg__19708, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{arg__19708, f})
	if callErr != nil {
		return nil, callErr
	}
	v33, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v33) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v123, nil
b5:
	;
	v36, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v36
	goto b1
b6:
	;
	arg__19721, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	v49 = arg__19721 == vm.Keyword("return")
	if v49 {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	v123 = vm.NIL
	goto b4
b8:
	;
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__19732, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v66 = arg__19732 == vm.Int(1)
	if v66 {
		goto b11
	} else {
		goto b12
	}
b9:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b14
	} else {
		goto b15
	}
b10:
	;
	goto b7
b11:
	;
	v69, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19750, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v79, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{out, arg__19750})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v69
	out = v79
	goto b1
b12:
	;
	goto b13
b13:
	;
	goto b10
b14:
	;
	v101, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	blocks = v101
	goto b1
b15:
	;
	goto b16
b16:
	;
	goto b10
}
func same_ident_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var a vm.Value
	var b vm.Value
	var v33 vm.Value
	var arg__19763 vm.Value
	var arg__19767 vm.Value
	var v24 bool
	var v27 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		a = arg0
		b = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident?").Deref(), []vm.Value{b})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v33 = and__x
	goto b3
b3:
	;
	return v33, nil
b4:
	;
	arg__19763, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident-name").Deref(), []vm.Value{a})
	if callErr != nil {
		return nil, callErr
	}
	arg__19767, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident-name").Deref(), []vm.Value{b})
	if callErr != nil {
		return nil, callErr
	}
	v24 = arg__19763 == arg__19767
	v27 = vm.Boolean(v24)
	goto b6
b5:
	;
	v27 = and__x
	goto b6
b6:
	;
	v33 = v27
	goto b3
}
func source_name_of(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var sis vm.Value
	var remaining vm.Value
	var v16 vm.Value
	var arg__19784 vm.Value
	var sym vm.Value
	var and__x vm.Value
	var v78 vm.Value
	var v67 vm.Value
	var v69 vm.Value
	var arg__19796 vm.Value
	var v54 vm.Value
	var v57 vm.Value
	var callErr error
	sis, callErr = rt.InvokeValue(rt.LookupVar("ir", "source-infos").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	remaining = sis
	goto b1
b1:
	;
	v16, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v16) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	arg__19784, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	sym, callErr = rt.InvokeValue(rt.LookupVar("ir", "source-info-symbol").Deref(), []vm.Value{arg__19784})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "string?").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b8
	} else {
		goto b9
	}
b3:
	;
	v78 = vm.NIL
	goto b4
b4:
	;
	return v78, nil
b5:
	;
	v69 = sym
	goto b7
b6:
	;
	v67, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{remaining})
	if callErr != nil {
		return nil, callErr
	}
	remaining = v67
	goto b1
b7:
	;
	v78 = v69
	goto b4
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	arg__19796, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{sym})
	if callErr != nil {
		return nil, callErr
	}
	v54, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__19796})
	if callErr != nil {
		return nil, callErr
	}
	v57 = v54
	goto b10
b9:
	;
	v57 = and__x
	goto b10
b10:
	;
	if vm.IsTruthy(v57) {
		goto b5
	} else {
		goto b6
	}
}
func template_fns(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var aux vm.Value
	var arg__19806 vm.Value
	var v9 vm.Value
	var v14 vm.Value
	var v39 vm.Value
	var arg__19818 vm.Value
	var v23 vm.Value
	var v36 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "fn-template?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v4) {
		aux = arg0
		goto b1
	} else {
		aux = arg0
		goto b2
	}
b1:
	;
	arg__19806, callErr = rt.InvokeValue(vm.Keyword("fn"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v9, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__19806})
	if callErr != nil {
		return nil, callErr
	}
	v39 = v9
	goto b3
b2:
	;
	v14, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "multi-fn-template?").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v14) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v39, nil
b4:
	;
	_, callErr = rt.InvokeValue(vm.Keyword("fns"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__19818, callErr = rt.InvokeValue(vm.Keyword("fns"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v23, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{vm.Keyword("fn"), arg__19818})
	if callErr != nil {
		return nil, callErr
	}
	v36 = v23
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
	v39 = v36
	goto b3
b7:
	;
	goto b9
b8:
	;
	goto b9
b9:
	;
	v36 = vm.NewArrayVector([]vm.Value{})
	goto b6
}
func term_targets(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v7 bool
	var aux vm.Value
	var v10 vm.Value
	var op vm.Value
	var v17 bool
	var v44 vm.Value
	var arg__19829 vm.Value
	var arg__19833 vm.Value
	var v24 vm.Value
	var v40 vm.Value
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
	v10, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v44 = v10
	goto b3
b2:
	;
	v17 = op == vm.Keyword("branch-if")
	if v17 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v44, nil
b4:
	;
	arg__19829, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__19833, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__19829, arg__19833})
	if callErr != nil {
		return nil, callErr
	}
	v40 = v24
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
	v44 = v40
	goto b3
b7:
	;
	goto b9
b8:
	;
	goto b9
b9:
	;
	v40 = vm.NewArrayVector([]vm.Value{})
	goto b6
}
func transfer_stmts(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var assigns vm.Value
	var bt vm.Value
	var arg__19850 vm.Value
	var arg__19861 vm.Value
	var arg__19862 vm.Value
	var arg__19874 vm.Value
	var arg__19885 vm.Value
	var arg__19886 vm.Value
	var arg__19887 vm.Value
	var v43 vm.Value
	var v47 vm.Value
	var callErr error
	assigns, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "emit-assignments-for-target").Deref(), []vm.Value{arg0, arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(assigns) {
		bt = arg2
		goto b1
	} else {
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	arg__19850, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "label-name").Deref(), []vm.Value{arg__19850})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	arg__19861, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	arg__19862, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "label-name").Deref(), []vm.Value{arg__19861})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "goto-stmt").Deref(), []vm.Value{arg__19862})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	arg__19874, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "label-name").Deref(), []vm.Value{arg__19874})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	arg__19885, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	arg__19886, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "label-name").Deref(), []vm.Value{arg__19885})
	if callErr != nil {
		return nil, callErr
	}
	arg__19887, callErr = rt.InvokeValue(rt.LookupVar("gogen", "goto-stmt").Deref(), []vm.Value{arg__19886})
	if callErr != nil {
		return nil, callErr
	}
	v43, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{assigns, arg__19887})
	if callErr != nil {
		return nil, callErr
	}
	v47 = v43
	goto b3
b2:
	;
	v47 = vm.NIL
	goto b3
b3:
	;
	return v47, nil
}
func unsupported(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v7 bool
	var msg vm.Value
	var v15 vm.Value
	var arg__19909 vm.Value
	var v26 vm.Value
	var v28 vm.Value
	var callErr error
	v7 = arg0 == vm.Keyword("bridge")
	if v7 {
		msg = arg1
		goto b1
	} else {
		msg = arg1
		goto b2
	}
b1:
	;
	v15, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("status"), vm.Keyword("fallback"), vm.Keyword("decl"), vm.NIL, vm.Keyword("reason"), msg})
	if callErr != nil {
		return nil, callErr
	}
	v28 = v15
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir.lower-go: "), msg})
	if callErr != nil {
		return nil, callErr
	}
	arg__19909, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir.lower-go: "), msg})
	if callErr != nil {
		return nil, callErr
	}
	v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__19909})
	if callErr != nil {
		return nil, callErr
	}
	v28 = v26
	goto b3
b3:
	;
	return v28, nil
}
func used_QMARK_(arg0 vm.Value, arg1 int) (vm.Value, error) {
	var uses vm.Value
	var arg__19917 vm.Value
	var and__x_6 bool
	var nid int
	var arg__19930 vm.Value
	var arg__19945 vm.Value
	var arg__19946 vm.Value
	var v29 vm.Value
	var and__x_14 bool
	var v32 vm.Value
	var callErr error
	uses, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__19917, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{uses})
	if callErr != nil {
		return nil, callErr
	}
	and__x_6 = rt.LtValue(vm.Int(arg1), arg__19917)
	if and__x_6 {
		nid = arg1
		goto b1
	} else {
		and__x_14 = and__x_6
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, vm.Int(nid)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19930, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, vm.Int(nid)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__19930})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, vm.Int(nid)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19945, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, vm.Int(nid)})
	if callErr != nil {
		return nil, callErr
	}
	arg__19946, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__19945})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{arg__19946})
	if callErr != nil {
		return nil, callErr
	}
	v32 = v29
	goto b3
b2:
	;
	v32 = vm.Boolean(and__x_14)
	goto b3
b3:
	;
	return v32, nil
}
func vm_call(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__20374 vm.Value
	var v7 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-sel").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__20374, callErr = rt.InvokeValue(rt.LookupVar("ir.lower-go", "vm-sel").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v7, callErr = rt.InvokeValue(rt.LookupVar("gogen", "call").Deref(), []vm.Value{arg__20374, arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v7, nil
}
func vm_sel(arg0 vm.Value) (vm.Value, error) {
	var arg__20385 vm.Value
	var v10 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	arg__20385, callErr = rt.InvokeValue(rt.LookupVar("gogen", "ident").Deref(), []vm.Value{vm.String("vm")})
	if callErr != nil {
		return nil, callErr
	}
	v10, callErr = rt.InvokeValue(rt.LookupVar("gogen", "field-sel").Deref(), []vm.Value{arg__20385, arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v10, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.lower-go", map[string]vm.Value{"any-fn-template?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("any-fn-template?: wrong number of arguments %d (expected 1)", len(args))
		}
		return any_fn_template_QMARK_(args[0])
	}), "arg-local-decls": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("arg-local-decls: wrong number of arguments %d (expected 1)", len(args))
		}
		return arg_local_decls(args[0])
	}), "block-arg-sources": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-arg-sources: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_arg_sources(args[0], args[1])
	}), "box-as-value": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("box-as-value: wrong number of arguments %d (expected 3)", len(args))
		}
		return box_as_value(args[0], args[1], args[2])
	}), "boxed-list-expr": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("boxed-list-expr: wrong number of arguments %d (expected 1)", len(args))
		}
		return boxed_list_expr(args[0])
	}), "boxed-map-expr": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("boxed-map-expr: wrong number of arguments %d (expected 1)", len(args))
		}
		return boxed_map_expr(args[0])
	}), "boxed-value-expr": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("boxed-value-expr: wrong number of arguments %d (expected 1)", len(args))
		}
		return boxed_value_expr(args[0])
	}), "boxed-vector-expr": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("boxed-vector-expr: wrong number of arguments %d (expected 1)", len(args))
		}
		return boxed_vector_expr(args[0])
	}), "call-assign-stmts": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("call-assign-stmts: wrong number of arguments %d (expected 3)", len(args))
		}
		return call_assign_stmts(args[0], args[1], args[2])
	}), "closure-expr": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("closure-expr: wrong number of arguments %d (expected 3)", len(args))
		}
		return closure_expr(args[0], args[1], args[2])
	}), "closure-info": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("closure-info: wrong number of arguments %d (expected 2)", len(args))
		}
		return closure_info(args[0], args[1])
	}), "closure-info*": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("closure-info*: wrong number of arguments %d (expected 3)", len(args))
		}
		return closure_info_STAR_(args[0], args[1], args[2])
	}), "closure-value?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("closure-value?: wrong number of arguments %d (expected 2)", len(args))
		}
		return closure_value_QMARK_(args[0], args[1])
	}), "coalesce-map": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("coalesce-map: wrong number of arguments %d (expected 1)", len(args))
		}
		return coalesce_map(args[0])
	}), "collect-local-ids": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("collect-local-ids: wrong number of arguments %d (expected 1)", len(args))
		}
		return collect_local_ids(args[0])
	}), "compute-param-sources": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("compute-param-sources: wrong number of arguments %d (expected 1)", len(args))
		}
		return compute_param_sources(args[0])
	}), "const-expr": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("const-expr: wrong number of arguments %d (expected 2)", len(args))
		}
		return const_expr(args[0], args[1])
	}), "const-inst-val": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("const-inst-val: wrong number of arguments %d (expected 2)", len(args))
		}
		return const_inst_val(args[0], args[1])
	}), "const-param-map": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("const-param-map: wrong number of arguments %d (expected 1)", len(args))
		}
		return const_param_map(args[0])
	}), "const-param-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("const-param-of: wrong number of arguments %d (expected 2)", len(args))
		}
		return const_param_of(args[0], args[1])
	}), "cp-meet": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("cp-meet: wrong number of arguments %d (expected 2)", len(args))
		}
		return cp_meet(args[0], args[1]), nil
	}), "cp-value": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("cp-value: wrong number of arguments %d (expected 4)", len(args))
		}
		return cp_value(args[0], args[1], args[2], args[3])
	}), "distinct-imports": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("distinct-imports: wrong number of arguments %d (expected 1)", len(args))
		}
		return distinct_imports(args[0])
	}), "emit-assignments-for-target": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("emit-assignments-for-target: wrong number of arguments %d (expected 3)", len(args))
		}
		return emit_assignments_for_target(args[0], args[1], args[2])
	}), "file": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("file: wrong number of arguments %d (expected 2)", len(args))
		}
		return file(args[0], args[1])
	}), "fn-template?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-template?: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_template_QMARK_(args[0])
	}), "function-needs-error?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("function-needs-error?: wrong number of arguments %d (expected 1)", len(args))
		}
		return function_needs_error_QMARK_(args[0])
	}), "function-needs-rt?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("function-needs-rt?: wrong number of arguments %d (expected 1)", len(args))
		}
		return function_needs_rt_QMARK_(args[0])
	}), "function-needs-tail-call?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("function-needs-tail-call?: wrong number of arguments %d (expected 1)", len(args))
		}
		return function_needs_tail_call_QMARK_(args[0])
	}), "function-needs-vm?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("function-needs-vm?: wrong number of arguments %d (expected 1)", len(args))
		}
		return function_needs_vm_QMARK_(args[0])
	}), "function-param-specs": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("function-param-specs: wrong number of arguments %d (expected 1)", len(args))
		}
		return function_param_specs(args[0])
	}), "function-return-spec": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("function-return-spec: wrong number of arguments %d (expected 1)", len(args))
		}
		return function_return_spec(args[0])
	}), "go-name": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("go-name: wrong number of arguments %d (expected 1)", len(args))
		}
		return go_name(args[0])
	}), "go-type-spec": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("go-type-spec: wrong number of arguments %d (expected 1)", len(args))
		}
		return go_type_spec(args[0])
	}), "import-entry": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("import-entry: wrong number of arguments %d (expected 2)", len(args))
		}
		return import_entry(args[0], args[1])
	}), "import-spec-node": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("import-spec-node: wrong number of arguments %d (expected 1)", len(args))
		}
		return import_spec_node(args[0])
	}), "infer-go-type": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("infer-go-type: wrong number of arguments %d (expected 1)", len(args))
		}
		return infer_go_type(args[0])
	}), "label-name": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("label-name: wrong number of arguments %d (expected 1)", len(args))
		}
		return label_name(args[0])
	}), "live-nids": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("live-nids: wrong number of arguments %d (expected 1)", len(args))
		}
		return live_nids(args[0])
	}), "live?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("live?: wrong number of arguments %d (expected 2)", len(args))
		}
		return live_QMARK_(args[0], args[1])
	}), "local-carrying-op?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("local-carrying-op?: wrong number of arguments %d (expected 1)", len(args))
		}
		return local_carrying_op_QMARK_(args[0])
	}), "local-decls": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("local-decls: wrong number of arguments %d (expected 1)", len(args))
		}
		return local_decls(args[0])
	}), "lower-block-stmts": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("lower-block-stmts: wrong number of arguments %d (expected 3)", len(args))
		}
		return lower_block_stmts(args[0], args[1], args[2])
	}), "lower-fn*": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("lower-fn*: wrong number of arguments %d (expected 2)", len(args))
		}
		return lower_fn_STAR_(args[0], args[1])
	}), "lower-fn-lit": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("lower-fn-lit: wrong number of arguments %d (expected 2)", len(args))
		}
		return lower_fn_lit(args[0], args[1])
	}), "lower-inst-stmts": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("lower-inst-stmts: wrong number of arguments %d (expected 3)", len(args))
		}
		return lower_inst_stmts(args[0], args[1], args[2])
	}), "lower-template-closure-expr": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("lower-template-closure-expr: wrong number of arguments %d (expected 2)", len(args))
		}
		return lower_template_closure_expr(args[0], args[1])
	}), "lower-terminator": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("lower-terminator: wrong number of arguments %d (expected 3)", len(args))
		}
		return lower_terminator(args[0], args[1], args[2])
	}), "multi-fn-template?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("multi-fn-template?: wrong number of arguments %d (expected 1)", len(args))
		}
		return multi_fn_template_QMARK_(args[0])
	}), "nested-template-fns": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("nested-template-fns: wrong number of arguments %d (expected 1)", len(args))
		}
		return nested_template_fns(args[0])
	}), "oos-all-nids": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("oos-all-nids: wrong number of arguments %d (expected 1)", len(args))
		}
		return oos_all_nids(args[0])
	}), "oos-branch-targets": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("oos-branch-targets: wrong number of arguments %d (expected 2)", len(args))
		}
		return oos_branch_targets(args[0], args[1])
	}), "oos-captured-names": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("oos-captured-names: wrong number of arguments %d (expected 1)", len(args))
		}
		return oos_captured_names(args[0])
	}), "oos-next-tmp!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("oos-next-tmp!: wrong number of arguments %d (expected 1)", len(args))
		}
		return oos_next_tmp_BANG_(args[0])
	}), "oos-param-incoming-args": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("oos-param-incoming-args: wrong number of arguments %d (expected 3)", len(args))
		}
		return oos_param_incoming_args(args[0], args[1], args[2])
	}), "oos-rhs-read-name": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("oos-rhs-read-name: wrong number of arguments %d (expected 1)", len(args))
		}
		return oos_rhs_read_name(args[0])
	}), "oos-sequence-copies": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("oos-sequence-copies: wrong number of arguments %d (expected 2)", len(args))
		}
		return oos_sequence_copies(args[0], args[1])
	}), "oos-unsafe-names": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("oos-unsafe-names: wrong number of arguments %d (expected 1)", len(args))
		}
		return oos_unsafe_names(args[0])
	}), "override-uniform-value?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("override-uniform-value?: wrong number of arguments %d (expected 1)", len(args))
		}
		return override_uniform_value_QMARK_(args[0])
	}), "params-for": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("params-for: wrong number of arguments %d (expected 1)", len(args))
		}
		return params_for(args[0])
	}), "result-node": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("result-node: wrong number of arguments %d (expected 1)", len(args))
		}
		return result_node(args[0])
	}), "return-ref-ids": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("return-ref-ids: wrong number of arguments %d (expected 1)", len(args))
		}
		return return_ref_ids(args[0])
	}), "same-ident?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("same-ident?: wrong number of arguments %d (expected 2)", len(args))
		}
		return same_ident_QMARK_(args[0], args[1])
	}), "source-name-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("source-name-of: wrong number of arguments %d (expected 2)", len(args))
		}
		return source_name_of(args[0], args[1])
	}), "template-fns": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("template-fns: wrong number of arguments %d (expected 1)", len(args))
		}
		return template_fns(args[0])
	}), "term-targets": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("term-targets: wrong number of arguments %d (expected 2)", len(args))
		}
		return term_targets(args[0], args[1])
	}), "transfer-stmts": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("transfer-stmts: wrong number of arguments %d (expected 3)", len(args))
		}
		return transfer_stmts(args[0], args[1], args[2])
	}), "unsupported": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("unsupported: wrong number of arguments %d (expected 2)", len(args))
		}
		return unsupported(args[0], args[1])
	}), "vm-call": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("vm-call: wrong number of arguments %d (expected 2)", len(args))
		}
		return vm_call(args[0], args[1])
	}), "vm-sel": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("vm-sel: wrong number of arguments %d (expected 1)", len(args))
		}
		return vm_sel(args[0])
	}),
	})
}
