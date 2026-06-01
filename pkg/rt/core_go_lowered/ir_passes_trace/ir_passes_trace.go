package ir_passes_trace

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func live_inst_count(arg0 vm.Value) (vm.Value, error) {
	var arg__34394 vm.Value
	var arg__34434 vm.Value
	var arg__34435 vm.Value
	var v29 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__34394, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__34389 vm.Value
		var v7 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__34389, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__34389})
		if callErr != nil {
			return nil, callErr
		}
		return v7, nil
	}), arg__34394})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__34434, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__34435, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__34429 vm.Value
		var v7 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__34429, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__34429})
		if callErr != nil {
			return nil, callErr
		}
		return v7, nil
	}), arg__34434})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.LookupVar("clojure.core", "+").Deref(), arg__34435})
	if callErr != nil {
		return nil, callErr
	}
	return v29, nil
}
func ns_now() (vm.Value, error) {
	var v1 vm.Value
	var callErr error
	v1, callErr = rt.InvokeValue(rt.LookupVar("System", "nanoTime").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	return v1, nil
}
func dump_pair(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var before_dump vm.Value
	var before_cnt vm.Value
	var t0 vm.Value
	var t1 vm.Value
	var after_cnt vm.Value
	var after_dump vm.Value
	var arg__34458 vm.Value
	var arg__34465 vm.Value
	var arg__34473 vm.Value
	var v30 vm.Value
	var callErr error
	before_dump, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "dump").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	before_cnt, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "live-inst-count").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	t0, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "ns-now").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(arg1, []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	t1, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "ns-now").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	after_cnt, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "live-inst-count").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	after_dump, callErr = rt.InvokeValue(rt.LookupVar("ir.dump", "dump").Deref(), []vm.Value{arg2})
	if callErr != nil {
		return nil, callErr
	}
	arg__34458 = rt.SubValue(t1, t0)
	arg__34465, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "/").Deref(), []vm.Value{arg__34458, vm.Float(1e+06)})
	if callErr != nil {
		return nil, callErr
	}
	arg__34473 = rt.SubValue(before_cnt, after_cnt)
	v30, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("ms"), arg__34465, vm.Keyword("pass"), arg0, vm.Keyword("after"), after_dump, vm.Keyword("delta"), arg__34473, vm.Keyword("before"), before_dump})
	if callErr != nil {
		return nil, callErr
	}
	return v30, nil
}
func optimize_fn_traced(arg0 vm.Value) (vm.Value, error) {
	var arg__34494 vm.Value
	var iter int
	var f vm.Value
	var before vm.Value
	var arg__34512 vm.Value
	var arg__34528 vm.Value
	var arg__34542 vm.Value
	var arg__34558 vm.Value
	var arg__34572 vm.Value
	var arg__34588 vm.Value
	var arg__34602 vm.Value
	var arg__34618 vm.Value
	var after vm.Value
	var v104 bool
	var v116 bool
	var arg__34657 vm.Value
	var arg__34642 vm.Value
	var v143 int
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{arg0, vm.String("build")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "typeinfer").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__34494, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "typeinfer").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "trace-pass").Deref(), []vm.Value{vm.String("typeinfer-pre"), vm.Int(-1), arg__34494, arg0})
	if callErr != nil {
		return nil, callErr
	}
	iter = 0
	f = arg0
	goto b1
b1:
	;
	before, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "live-inst-count").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "constfold").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__34512, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.constfold", "constfold").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "trace-pass").Deref(), []vm.Value{vm.String("constfold"), vm.Int(iter), arg__34512, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("constfold/"), vm.Int(iter)})
	if callErr != nil {
		return nil, callErr
	}
	arg__34528, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("constfold/"), vm.Int(iter)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{f, arg__34528})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.cse", "cse").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__34542, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.cse", "cse").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "trace-pass").Deref(), []vm.Value{vm.String("cse"), vm.Int(iter), arg__34542, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("cse/"), vm.Int(iter)})
	if callErr != nil {
		return nil, callErr
	}
	arg__34558, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("cse/"), vm.Int(iter)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{f, arg__34558})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "licm").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__34572, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "licm").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "trace-pass").Deref(), []vm.Value{vm.String("licm"), vm.Int(iter), arg__34572, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("licm/"), vm.Int(iter)})
	if callErr != nil {
		return nil, callErr
	}
	arg__34588, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("licm/"), vm.Int(iter)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{f, arg__34588})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.dce", "dce").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__34602, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.dce", "dce").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "trace-pass").Deref(), []vm.Value{vm.String("dce"), vm.Int(iter), arg__34602, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("dce/"), vm.Int(iter)})
	if callErr != nil {
		return nil, callErr
	}
	arg__34618, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("dce/"), vm.Int(iter)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.validate", "validate-fn!").Deref(), []vm.Value{f, arg__34618})
	if callErr != nil {
		return nil, callErr
	}
	after, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "live-inst-count").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	v104 = before == after
	if v104 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	goto b4
b3:
	;
	v116 = iter >= 15
	if v116 {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "typeinfer").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__34657, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.typeinfer", "typeinfer").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.trace", "trace-pass").Deref(), []vm.Value{vm.String("typeinfer-post"), vm.Int(-1), arg__34657, f})
	if callErr != nil {
		return nil, callErr
	}
	return f, nil
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("warn: optimize-fn-traced max iters (16) reached, "), before, vm.String(" insts after 16 cycles")})
	if callErr != nil {
		return nil, callErr
	}
	arg__34642, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("warn: optimize-fn-traced max iters (16) reached, "), before, vm.String(" insts after 16 cycles")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "println").Deref(), []vm.Value{arg__34642})
	if callErr != nil {
		return nil, callErr
	}
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
	goto b4
b8:
	;
	v143 = iter + 1
	iter = v143
	goto b1
b9:
	;
	goto b10
b10:
	;
	goto b7
}
func print_trace(arg0 vm.Value) (vm.Value, error) {
	var arg__34693 vm.Value
	var arg__34708 vm.Value
	var arg__34725 vm.Value
	var arg__34726 vm.Value
	var doseq_seq__34659 vm.Value
	var doseq_loop__34660 vm.Value
	var e vm.Value
	var arg__34756 vm.Value
	var arg__34759 vm.Value
	var arg__34762 vm.Value
	var arg__34765 vm.Value
	var arg__34768 vm.Value
	var arg__34771 vm.Value
	var arg__34797 vm.Value
	var arg__34800 vm.Value
	var arg__34803 vm.Value
	var arg__34806 vm.Value
	var arg__34809 vm.Value
	var arg__34812 vm.Value
	var arg__34813 vm.Value
	var v142 vm.Value
	var trace vm.Value
	var arg__34831 vm.Value
	var total_ms vm.Value
	var arg__34846 vm.Value
	var total_removed vm.Value
	var arg__34861 vm.Value
	var arg__34878 vm.Value
	var arg__34879 vm.Value
	var arg__34892 vm.Value
	var arg__34909 vm.Value
	var arg__34912 vm.Value
	var v225 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "format").Deref(), []vm.Value{vm.String("%4s  %-18s  %6s  %6s  %6s  %7s"), vm.String("iter"), vm.String("pass"), vm.String("before"), vm.String("after"), vm.String("delta"), vm.String("ms")})
	if callErr != nil {
		return nil, callErr
	}
	arg__34693, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "format").Deref(), []vm.Value{vm.String("%4s  %-18s  %6s  %6s  %6s  %7s"), vm.String("iter"), vm.String("pass"), vm.String("before"), vm.String("after"), vm.String("delta"), vm.String("ms")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "println").Deref(), []vm.Value{arg__34693})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{vm.Int(55), vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__34708, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{vm.Int(55), vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__34708})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{vm.Int(55), vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__34725, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{vm.Int(55), vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__34726, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__34725})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "println").Deref(), []vm.Value{arg__34726})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__34659, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__34660 = doseq_seq__34659
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__34660) {
		goto b2
	} else {
		trace = arg0
		goto b3
	}
b2:
	;
	e, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__34660})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("iter"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("pass"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("before"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("after"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("delta"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("ms"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34756, callErr = rt.InvokeValue(vm.Keyword("iter"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34759, callErr = rt.InvokeValue(vm.Keyword("pass"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34762, callErr = rt.InvokeValue(vm.Keyword("before"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34765, callErr = rt.InvokeValue(vm.Keyword("after"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34768, callErr = rt.InvokeValue(vm.Keyword("delta"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34771, callErr = rt.InvokeValue(vm.Keyword("ms"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "format").Deref(), []vm.Value{vm.String("%4d  %-18s  %6d  %6d  %+6d  %7.2f"), arg__34756, arg__34759, arg__34762, arg__34765, arg__34768, arg__34771})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("iter"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("pass"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("before"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("after"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("delta"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("ms"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34797, callErr = rt.InvokeValue(vm.Keyword("iter"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34800, callErr = rt.InvokeValue(vm.Keyword("pass"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34803, callErr = rt.InvokeValue(vm.Keyword("before"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34806, callErr = rt.InvokeValue(vm.Keyword("after"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34809, callErr = rt.InvokeValue(vm.Keyword("delta"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34812, callErr = rt.InvokeValue(vm.Keyword("ms"), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__34813, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "format").Deref(), []vm.Value{vm.String("%4d  %-18s  %6d  %6d  %+6d  %7.2f"), arg__34797, arg__34800, arg__34803, arg__34806, arg__34809, arg__34812})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "println").Deref(), []vm.Value{arg__34813})
	if callErr != nil {
		return nil, callErr
	}
	v142, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__34660})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__34660 = v142
	goto b1
b3:
	;
	goto b4
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{vm.Keyword("ms"), trace})
	if callErr != nil {
		return nil, callErr
	}
	arg__34831, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{vm.Keyword("ms"), trace})
	if callErr != nil {
		return nil, callErr
	}
	total_ms, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.LookupVar("clojure.core", "+").Deref(), arg__34831})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{vm.Keyword("delta"), trace})
	if callErr != nil {
		return nil, callErr
	}
	arg__34846, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map").Deref(), []vm.Value{vm.Keyword("delta"), trace})
	if callErr != nil {
		return nil, callErr
	}
	total_removed, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reduce").Deref(), []vm.Value{rt.LookupVar("clojure.core", "+").Deref(), arg__34846})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{vm.Int(55), vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__34861, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{vm.Int(55), vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__34861})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{vm.Int(55), vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__34878, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{vm.Int(55), vm.String("-")})
	if callErr != nil {
		return nil, callErr
	}
	arg__34879, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "str").Deref(), arg__34878})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "println").Deref(), []vm.Value{arg__34879})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{trace})
	if callErr != nil {
		return nil, callErr
	}
	arg__34892, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{trace})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "format").Deref(), []vm.Value{vm.String("  %d passes, %d insts removed, %.2f ms total"), arg__34892, total_removed, total_ms})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{trace})
	if callErr != nil {
		return nil, callErr
	}
	arg__34909, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{trace})
	if callErr != nil {
		return nil, callErr
	}
	arg__34912, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "format").Deref(), []vm.Value{vm.String("  %d passes, %d insts removed, %.2f ms total"), arg__34909, total_removed, total_ms})
	if callErr != nil {
		return nil, callErr
	}
	v225, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "println").Deref(), []vm.Value{arg__34912})
	if callErr != nil {
		return nil, callErr
	}
	return v225, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.trace", map[string]vm.Value{"live-inst-count": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("live-inst-count: wrong number of arguments %d (expected 1)", len(args))
		}
		return live_inst_count(args[0])
	}), "ns-now": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 0 {
			return nil, fmt.Errorf("ns-now: wrong number of arguments %d (expected 0)", len(args))
		}
		return ns_now()
	}), "dump-pair": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("dump-pair: wrong number of arguments %d (expected 3)", len(args))
		}
		return dump_pair(args[0], args[1], args[2])
	}), "optimize-fn-traced": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("optimize-fn-traced: wrong number of arguments %d (expected 1)", len(args))
		}
		return optimize_fn_traced(args[0])
	}), "print-trace": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("print-trace: wrong number of arguments %d (expected 1)", len(args))
		}
		return print_trace(args[0])
	}),
	})
}
