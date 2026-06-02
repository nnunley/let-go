package ir_validate

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func check_no_cross_block_refs_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__31471 vm.Value
	var insts vm.Value
	var arg__31484 vm.Value
	var doseq_seq__31462 vm.Value
	var doseq_loop__31463 vm.Value
	var label vm.Value
	var vec__31466 vm.Value
	var i vm.Value
	var ins vm.Value
	var op vm.Value
	var v71 vm.Value
	var arg__31524 vm.Value
	var doseq_seq__31464 vm.Value
	var v285 vm.Value
	var doseq_loop__31465 vm.Value
	var r vm.Value
	var referent vm.Value
	var ref_op vm.Value
	var def_block vm.Value
	var use_block vm.Value
	var and__x vm.Value
	var v282 vm.Value
	var arg__31618 vm.Value
	var v262 vm.Value
	var v265 vm.Value
	var v201 vm.Value
	var v204 vm.Value
	var callErr error
	arg__31471, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__31471})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map-indexed").Deref(), []vm.Value{rt.LookupVar("clojure.core", "vector").Deref(), insts})
	if callErr != nil {
		return nil, callErr
	}
	arg__31484, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map-indexed").Deref(), []vm.Value{rt.LookupVar("clojure.core", "vector").Deref(), insts})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__31462, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__31484})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__31463 = doseq_seq__31462
	label = arg1
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__31463) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	vec__31466, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__31463})
	if callErr != nil {
		return nil, callErr
	}
	i, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__31466, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	ins, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{vec__31466, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{ins, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v71, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{op, vm.Keyword("invalid")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v71) {
		goto b5
	} else {
		goto b6
	}
b3:
	;
	goto b4
b4:
	;
	return vm.NIL, nil
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{ins, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__31524, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{ins, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__31464, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__31524})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__31465 = doseq_seq__31464
	goto b8
b6:
	;
	v285, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__31463})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__31463 = v285
	goto b1
b8:
	;
	if vm.IsTruthy(doseq_loop__31465) {
		goto b9
	} else {
		goto b10
	}
b9:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__31465})
	if callErr != nil {
		return nil, callErr
	}
	referent, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{insts, r})
	if callErr != nil {
		return nil, callErr
	}
	ref_op, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{referent, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	def_block, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{referent, vm.Int(3)})
	if callErr != nil {
		return nil, callErr
	}
	use_block, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{ins, vm.Int(3)})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{def_block, use_block})
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
	v282, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__31463})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__31463 = v282
	goto b1
b12:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": Inst #"), i, vm.String(" (op="), op, vm.String(", block="), use_block, vm.String(")"), vm.String(" has cross-block ref "), r, vm.String(" (defined in block "), def_block, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	arg__31618, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": Inst #"), i, vm.String(" (op="), op, vm.String(", block="), use_block, vm.String(")"), vm.String(" has cross-block ref "), r, vm.String(" (defined in block "), def_block, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__31618})
	if callErr != nil {
		return nil, callErr
	}
	v262, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__31465})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__31465 = v262
	goto b8
b13:
	;
	v265, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__31465})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__31465 = v265
	goto b8
b15:
	;
	v201, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("block-arg"), ref_op})
	if callErr != nil {
		return nil, callErr
	}
	v204 = v201
	goto b17
b16:
	;
	v204 = and__x
	goto b17
b17:
	;
	if vm.IsTruthy(v204) {
		goto b12
	} else {
		goto b13
	}
}
func check_branch_if_symmetric_args_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__31635 vm.Value
	var blocks vm.Value
	var arg__31640 vm.Value
	var insts vm.Value
	var i int
	var label vm.Value
	var arg__31645 vm.Value
	var v29 bool
	var arg__31652 vm.Value
	var term_id vm.Value
	var term vm.Value
	var op vm.Value
	var v98 bool
	var v211 int
	var v63 bool
	var and__x vm.Value
	var v66 vm.Value
	var aux vm.Value
	var arg__31680 vm.Value
	var t_args vm.Value
	var arg__31689 vm.Value
	var f_args vm.Value
	var v138 bool
	var arg__31740 vm.Value
	var callErr error
	arg__31635, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	blocks, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__31635})
	if callErr != nil {
		return nil, callErr
	}
	arg__31640, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__31640})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	label = arg1
	goto b1
b1:
	;
	arg__31645, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	v29 = rt.LtValue(vm.Int(i), arg__31645)
	if v29 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	arg__31652, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{blocks, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	term_id, callErr = rt.InvokeValue(vm.Keyword("term"), []vm.Value{arg__31652})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(term_id) {
		goto b8
	} else {
		and__x = term_id
		goto b9
	}
b3:
	;
	goto b4
b4:
	;
	return vm.NIL, nil
b5:
	;
	term, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{insts, term_id})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{term, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v98 = op == vm.Keyword("branch-if")
	if v98 {
		goto b11
	} else {
		goto b12
	}
b6:
	;
	goto b7
b7:
	;
	v211 = i + 1
	i = v211
	goto b1
b8:
	;
	v63 = rt.GtValue(term_id, vm.Int(0))
	v66 = vm.Boolean(v63)
	goto b10
b9:
	;
	v66 = and__x
	goto b10
b10:
	;
	if vm.IsTruthy(v66) {
		goto b5
	} else {
		goto b6
	}
b11:
	;
	aux, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{term, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__31680, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	t_args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{arg__31680})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__31689, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	f_args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{arg__31689})
	if callErr != nil {
		return nil, callErr
	}
	v138 = t_args == f_args
	if v138 {
		goto b14
	} else {
		goto b15
	}
b12:
	;
	goto b13
b13:
	;
	goto b7
b14:
	;
	goto b16
b15:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), vm.Int(i), vm.String(" :branch-if has asymmetric"), vm.String(" branch-target args (true="), t_args, vm.String(", false="), f_args, vm.String(") — lower-block! requires"), vm.String(" true.args == false.args")})
	if callErr != nil {
		return nil, callErr
	}
	arg__31740, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), vm.Int(i), vm.String(" :branch-if has asymmetric"), vm.String(" branch-target args (true="), t_args, vm.String(", false="), f_args, vm.String(") — lower-block! requires"), vm.String(" true.args == false.args")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__31740})
	if callErr != nil {
		return nil, callErr
	}
	goto b16
b16:
	;
	goto b13
}
func check_inst_shapes_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var i int
	var label vm.Value
	var insts vm.Value
	var arg__31746 vm.Value
	var v15 bool
	var ins vm.Value
	var and__x vm.Value
	var arg__31788 vm.Value
	var v79 int
	var arg__31759 vm.Value
	var v43 bool
	var v46 vm.Value
	var callErr error
	i = 0
	label = arg1
	insts = arg0
	goto b1
b1:
	;
	arg__31746, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	v15 = rt.LtValue(vm.Int(i), arg__31746)
	if v15 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	ins, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{insts, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector?").Deref(), []vm.Value{ins})
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
	goto b4
b4:
	;
	return vm.NIL, nil
b5:
	;
	goto b7
b6:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": Inst #"), vm.Int(i), vm.String(" has bad shape: "), ins})
	if callErr != nil {
		return nil, callErr
	}
	arg__31788, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": Inst #"), vm.Int(i), vm.String(" has bad shape: "), ins})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__31788})
	if callErr != nil {
		return nil, callErr
	}
	goto b7
b7:
	;
	v79 = i + 1
	i = v79
	goto b1
b8:
	;
	arg__31759, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{ins})
	if callErr != nil {
		return nil, callErr
	}
	v43 = arg__31759 == vm.Int(6)
	v46 = vm.Boolean(v43)
	goto b10
b9:
	;
	v46 = and__x
	goto b10
b10:
	;
	if vm.IsTruthy(v46) {
		goto b5
	} else {
		goto b6
	}
}
func check_refs_in_range_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var n vm.Value
	var i int
	var label vm.Value
	var insts vm.Value
	var v18 bool
	var ins vm.Value
	var op vm.Value
	var refs vm.Value
	var v47 vm.Value
	var doseq_seq__31790 vm.Value
	var v252 int
	var doseq_loop__31791 vm.Value
	var r vm.Value
	var or__x_99 vm.Value
	var arg__31882 vm.Value
	var v227 vm.Value
	var or__x_110 vm.Value
	var or__x_125 bool
	var v165 vm.Value
	var or__x_136 bool
	var v150 bool
	var v152 bool
	var callErr error
	n, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	label = arg1
	insts = arg0
	goto b1
b1:
	;
	v18 = rt.LtValue(vm.Int(i), n)
	if v18 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	ins, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{insts, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{ins, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{ins, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v47, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{op, vm.Keyword("invalid")})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v47) {
		goto b5
	} else {
		goto b6
	}
b3:
	;
	goto b4
b4:
	;
	return vm.NIL, nil
b5:
	;
	doseq_seq__31790, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__31791 = doseq_seq__31790
	goto b8
b6:
	;
	goto b7
b7:
	;
	v252 = i + 1
	i = v252
	goto b1
b8:
	;
	if vm.IsTruthy(doseq_loop__31791) {
		goto b9
	} else {
		goto b10
	}
b9:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__31791})
	if callErr != nil {
		return nil, callErr
	}
	or__x_99, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{r})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x_99) {
		or__x_110 = or__x_99
		goto b15
	} else {
		goto b16
	}
b10:
	;
	goto b11
b11:
	;
	goto b7
b12:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": Inst #"), vm.Int(i), vm.String(" (op="), op, vm.String(")"), vm.String(" has out-of-range ref "), r, vm.String(" (insts count="), n, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	arg__31882, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": Inst #"), vm.Int(i), vm.String(" (op="), op, vm.String(")"), vm.String(" has out-of-range ref "), r, vm.String(" (insts count="), n, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__31882})
	if callErr != nil {
		return nil, callErr
	}
	goto b14
b13:
	;
	goto b14
b14:
	;
	v227, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__31791})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__31791 = v227
	goto b8
b15:
	;
	v165 = or__x_110
	goto b17
b16:
	;
	or__x_125 = rt.LtValue(r, vm.Int(0))
	if or__x_125 {
		or__x_136 = or__x_125
		goto b18
	} else {
		goto b19
	}
b17:
	;
	if vm.IsTruthy(v165) {
		goto b12
	} else {
		goto b13
	}
b18:
	;
	v152 = or__x_136
	goto b20
b19:
	;
	v150 = rt.GeValue(r, n)
	v152 = v150
	goto b20
b20:
	;
	v165 = vm.Boolean(v152)
	goto b17
}
func check_branch_arg_arity_for_target_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var target vm.Value
	var args vm.Value
	var tgt_blk vm.Value
	var params vm.Value
	var arg__31903 vm.Value
	var arg__31907 vm.Value
	var v33 bool
	var label vm.Value
	var src_block vm.Value
	var arg__31938 vm.Value
	var arg__31945 vm.Value
	var arg__31979 vm.Value
	var arg__31986 vm.Value
	var arg__31988 vm.Value
	var v86 vm.Value
	var v88 vm.Value
	var callErr error
	target, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	tgt_blk, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg3, target})
	if callErr != nil {
		return nil, callErr
	}
	params, callErr = rt.InvokeValue(vm.Keyword("params"), []vm.Value{tgt_blk})
	if callErr != nil {
		return nil, callErr
	}
	arg__31903, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	arg__31907, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	v33 = arg__31903 == arg__31907
	if v33 {
		goto b1
	} else {
		label = arg0
		src_block = arg1
		goto b2
	}
b1:
	;
	v88 = vm.NIL
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	arg__31938, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	arg__31945, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), src_block, vm.String(" branch to b"), target, vm.String(" passes "), arg__31938, vm.String(" args but b"), target, vm.String(" has "), arg__31945, vm.String(" params")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	arg__31979, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	arg__31986, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	arg__31988, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), src_block, vm.String(" branch to b"), target, vm.String(" passes "), arg__31979, vm.String(" args but b"), target, vm.String(" has "), arg__31986, vm.String(" params")})
	if callErr != nil {
		return nil, callErr
	}
	v86, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__31988})
	if callErr != nil {
		return nil, callErr
	}
	v88 = v86
	goto b3
b3:
	;
	return v88, nil
}
func check_blocks_terminated_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__31993 vm.Value
	var blocks vm.Value
	var arg__31998 vm.Value
	var insts vm.Value
	var i int
	var label vm.Value
	var arg__32003 vm.Value
	var v28 bool
	var blk vm.Value
	var preds vm.Value
	var term_id vm.Value
	var or__x_53 bool
	var v105 vm.Value
	var v313 int
	var or__x_62 bool
	var v75 vm.Value
	var v77 vm.Value
	var arg__32045 vm.Value
	var or__x_154 bool
	var arg__32076 vm.Value
	var arg__32103 vm.Value
	var arg__32105 vm.Value
	var term_inst vm.Value
	var term_op vm.Value
	var v263 vm.Value
	var or__x_163 bool
	var arg__32052 vm.Value
	var v177 bool
	var v179 bool
	var arg__32155 vm.Value
	var callErr error
	arg__31993, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	blocks, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__31993})
	if callErr != nil {
		return nil, callErr
	}
	arg__31998, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__31998})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	label = arg1
	goto b1
b1:
	;
	arg__32003, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	v28 = rt.LtValue(vm.Int(i), arg__32003)
	if v28 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	blk, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{blocks, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	preds, callErr = rt.InvokeValue(vm.Keyword("preds"), []vm.Value{blk})
	if callErr != nil {
		return nil, callErr
	}
	term_id, callErr = rt.InvokeValue(vm.Keyword("term"), []vm.Value{blk})
	if callErr != nil {
		return nil, callErr
	}
	or__x_53 = i == 0
	if or__x_53 {
		or__x_62 = or__x_53
		goto b8
	} else {
		goto b9
	}
b3:
	;
	goto b4
b4:
	;
	return vm.NIL, nil
b5:
	;
	v105, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{term_id})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v105) {
		goto b11
	} else {
		goto b12
	}
b6:
	;
	goto b7
b7:
	;
	v313 = i + 1
	i = v313
	goto b1
b8:
	;
	v77 = vm.Boolean(or__x_62)
	goto b10
b9:
	;
	v75, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{preds})
	if callErr != nil {
		return nil, callErr
	}
	v77 = v75
	goto b10
b10:
	;
	if vm.IsTruthy(v77) {
		goto b5
	} else {
		goto b6
	}
b11:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), vm.Int(i), vm.String(" has no :term field")})
	if callErr != nil {
		return nil, callErr
	}
	arg__32045, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), vm.Int(i), vm.String(" has no :term field")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__32045})
	if callErr != nil {
		return nil, callErr
	}
	goto b13
b12:
	;
	goto b13
b13:
	;
	or__x_154 = rt.LtValue(term_id, vm.Int(0))
	if or__x_154 {
		or__x_163 = or__x_154
		goto b17
	} else {
		goto b18
	}
b14:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	arg__32076, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), vm.Int(i), vm.String(" :term="), term_id, vm.String(" out of insts range [0,"), arg__32076, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	arg__32103, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	arg__32105, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), vm.Int(i), vm.String(" :term="), term_id, vm.String(" out of insts range [0,"), arg__32103, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__32105})
	if callErr != nil {
		return nil, callErr
	}
	goto b16
b15:
	;
	goto b16
b16:
	;
	term_inst, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{insts, term_id})
	if callErr != nil {
		return nil, callErr
	}
	term_op, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{term_inst, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v263, callErr = rt.InvokeValue(rt.LookupVar("ir", "op-terminator?").Deref(), []vm.Value{term_op})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v263) {
		goto b20
	} else {
		goto b21
	}
b17:
	;
	v179 = or__x_163
	goto b19
b18:
	;
	arg__32052, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	v177 = rt.GeValue(term_id, arg__32052)
	v179 = v177
	goto b19
b19:
	;
	if v179 {
		goto b14
	} else {
		goto b15
	}
b20:
	;
	goto b22
b21:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), vm.Int(i), vm.String(" :term Inst #"), term_id, vm.String(" has non-terminator op "), term_op})
	if callErr != nil {
		return nil, callErr
	}
	arg__32155, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("validate after "), label, vm.String(": block #"), vm.Int(i), vm.String(" :term Inst #"), term_id, vm.String(" has non-terminator op "), term_op})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__32155})
	if callErr != nil {
		return nil, callErr
	}
	goto b22
b22:
	;
	goto b7
}
func check_branch_arg_arities_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__32161 vm.Value
	var blocks vm.Value
	var arg__32166 vm.Value
	var insts vm.Value
	var i int
	var label vm.Value
	var arg__32171 vm.Value
	var v29 bool
	var blk vm.Value
	var term_id vm.Value
	var term vm.Value
	var op vm.Value
	var aux vm.Value
	var v81 bool
	var v158 int
	var v107 bool
	var arg__32220 vm.Value
	var arg__32235 vm.Value
	var callErr error
	arg__32161, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	blocks, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__32161})
	if callErr != nil {
		return nil, callErr
	}
	arg__32166, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__32166})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	label = arg1
	goto b1
b1:
	;
	arg__32171, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{blocks})
	if callErr != nil {
		return nil, callErr
	}
	v29 = rt.LtValue(vm.Int(i), arg__32171)
	if v29 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	blk, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{blocks, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	term_id, callErr = rt.InvokeValue(vm.Keyword("term"), []vm.Value{blk})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(term_id) {
		goto b5
	} else {
		goto b6
	}
b3:
	;
	goto b4
b4:
	;
	return vm.NIL, nil
b5:
	;
	term, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{insts, term_id})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{term, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{term, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	v81 = op == vm.Keyword("branch")
	if v81 {
		goto b8
	} else {
		goto b9
	}
b6:
	;
	goto b7
b7:
	;
	v158 = i + 1
	i = v158
	goto b1
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-branch-arg-arity-for-target!").Deref(), []vm.Value{label, vm.Int(i), aux, blocks})
	if callErr != nil {
		return nil, callErr
	}
	goto b10
b9:
	;
	v107 = op == vm.Keyword("branch-if")
	if v107 {
		goto b11
	} else {
		goto b12
	}
b10:
	;
	goto b7
b11:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__32220, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-branch-arg-arity-for-target!").Deref(), []vm.Value{label, vm.Int(i), arg__32220, blocks})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__32235, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-branch-arg-arity-for-target!").Deref(), []vm.Value{label, vm.Int(i), arg__32235, blocks})
	if callErr != nil {
		return nil, callErr
	}
	goto b13
b12:
	;
	goto b13
b13:
	;
	goto b10
}
func validate_fn_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__32242 vm.Value
	var insts vm.Value
	var callErr error
	arg__32242, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__32242})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-inst-shapes!").Deref(), []vm.Value{insts, arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-refs-in-range!").Deref(), []vm.Value{insts, arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-blocks-terminated!").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-no-cross-block-refs!").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-branch-arg-arities!").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "check-branch-if-symmetric-args!").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	return arg0, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.validate", map[string]vm.Value{"check-no-cross-block-refs!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("check-no-cross-block-refs!: wrong number of arguments %d (expected 2)", len(args))
		}
		return check_no_cross_block_refs_BANG_(args[0], args[1])
	}), "check-branch-if-symmetric-args!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("check-branch-if-symmetric-args!: wrong number of arguments %d (expected 2)", len(args))
		}
		return check_branch_if_symmetric_args_BANG_(args[0], args[1])
	}), "check-inst-shapes!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("check-inst-shapes!: wrong number of arguments %d (expected 2)", len(args))
		}
		return check_inst_shapes_BANG_(args[0], args[1])
	}), "check-refs-in-range!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("check-refs-in-range!: wrong number of arguments %d (expected 2)", len(args))
		}
		return check_refs_in_range_BANG_(args[0], args[1])
	}), "check-branch-arg-arity-for-target!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("check-branch-arg-arity-for-target!: wrong number of arguments %d (expected 4)", len(args))
		}
		return check_branch_arg_arity_for_target_BANG_(args[0], args[1], args[2], args[3])
	}), "check-blocks-terminated!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("check-blocks-terminated!: wrong number of arguments %d (expected 2)", len(args))
		}
		return check_blocks_terminated_BANG_(args[0], args[1])
	}), "check-branch-arg-arities!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("check-branch-arg-arities!: wrong number of arguments %d (expected 2)", len(args))
		}
		return check_branch_arg_arities_BANG_(args[0], args[1])
	}), "validate-fn!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("validate-fn!: wrong number of arguments %d (expected 2)", len(args))
		}
		return validate_fn_BANG_(args[0], args[1])
	}),
	})
}
