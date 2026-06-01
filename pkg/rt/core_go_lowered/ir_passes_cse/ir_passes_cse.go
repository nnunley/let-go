package ir_passes_cse

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func inst_key(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__20972 vm.Value
	var arg__20985 vm.Value
	var arg__20986 vm.Value
	var arg__20992 vm.Value
	var v14 vm.Value
	var callErr error
	arg__20972, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__20985, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__20986, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__20985})
	if callErr != nil {
		return nil, callErr
	}
	arg__20992, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__20972, arg__20986, arg__20992})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func cse_eligible_QMARK_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var or__x vm.Value
	var op vm.Value
	var nid vm.Value
	var f vm.Value
	var var_facts vm.Value
	var and__x_21 bool
	var v51 vm.Value
	var arg__21014 vm.Value
	var v40 vm.Value
	var and__x_33 bool
	var v43 vm.Value
	var callErr error
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.passes.cse", "pure-cse-ops").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b1
	} else {
		op = arg0
		nid = arg1
		f = arg2
		var_facts = arg3
		goto b2
	}
b1:
	;
	v51 = or__x
	goto b3
b2:
	;
	and__x_21 = op == vm.Keyword("load-var")
	if and__x_21 {
		goto b4
	} else {
		and__x_33 = and__x_21
		goto b5
	}
b3:
	;
	return v51, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21014, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v40, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.mutability", "stable-load-var?").Deref(), []vm.Value{var_facts, arg__21014})
	if callErr != nil {
		return nil, callErr
	}
	v43 = v40
	goto b6
b5:
	;
	v43 = vm.Boolean(and__x_33)
	goto b6
b6:
	;
	v51 = v43
	goto b3
}
func cse(arg0 vm.Value) (vm.Value, error) {
	var var_facts vm.Value
	var arg__21030 vm.Value
	var doseq_seq__21015 vm.Value
	var doseq_loop__21016 vm.Value
	var f vm.Value
	var bid vm.Value
	var seen vm.Value
	var arg__21049 vm.Value
	var arg__21064 vm.Value
	var arg__21065 vm.Value
	var doseq_seq__21017 vm.Value
	var doseq_loop__21018 vm.Value
	var nid vm.Value
	var op vm.Value
	var v90 vm.Value
	var v161 vm.Value
	var k vm.Value
	var head__21092 vm.Value
	var tem__G__0 vm.Value
	var v147 vm.Value
	var v144 vm.Value
	var callErr error
	var_facts, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.mutability", "analyze-var-stability").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__21030, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__21015, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__21030})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21016 = doseq_seq__21015
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__21016) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__21016})
	if callErr != nil {
		return nil, callErr
	}
	seen, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.EmptyPersistentMap})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21049, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__21049})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21064, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{bid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__21065, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__21064})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__21017, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__21065})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21018 = doseq_seq__21017
	goto b5
b3:
	;
	goto b4
b4:
	;
	return f, nil
b5:
	;
	if vm.IsTruthy(doseq_loop__21018) {
		goto b6
	} else {
		goto b7
	}
b6:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__21018})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v90, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.cse", "cse-eligible?").Deref(), []vm.Value{op, nid, f, var_facts})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v90) {
		goto b9
	} else {
		goto b10
	}
b7:
	;
	goto b8
b8:
	;
	v161, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__21016})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21016 = v161
	goto b1
b9:
	;
	k, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.cse", "inst-key").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	head__21092, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{seen})
	if callErr != nil {
		return nil, callErr
	}
	tem__G__0, callErr = rt.InvokeValue(head__21092, []vm.Value{k})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(tem__G__0) {
		goto b12
	} else {
		goto b13
	}
b10:
	;
	v147, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__21018})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21018 = v147
	goto b5
b12:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "replace-all-uses!").Deref(), []vm.Value{f, nid, tem__G__0})
	if callErr != nil {
		return nil, callErr
	}
	goto b14
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{seen, rt.LookupVar("clojure.core", "assoc").Deref(), k, nid})
	if callErr != nil {
		return nil, callErr
	}
	goto b14
b14:
	;
	v144, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__21018})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__21018 = v144
	goto b5
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.cse", map[string]vm.Value{"inst-key": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("inst-key: wrong number of arguments %d (expected 2)", len(args))
		}
		return inst_key(args[0], args[1])
	}), "cse-eligible?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("cse-eligible?: wrong number of arguments %d (expected 4)", len(args))
		}
		return cse_eligible_QMARK_(args[0], args[1], args[2], args[3])
	}), "cse": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("cse: wrong number of arguments %d (expected 1)", len(args))
		}
		return cse(args[0])
	}),
	})
}
