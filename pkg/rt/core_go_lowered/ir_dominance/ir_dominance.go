package ir_dominance

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func intersect(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var b2 vm.Value
	var b1 vm.Value
	var rpo_idx vm.Value
	var idom vm.Value
	var v20 bool
	var arg__9099 vm.Value
	var arg__9105 vm.Value
	var v35 bool
	var v88 vm.Value
	var arg__9116 vm.Value
	var arg__9122 vm.Value
	var v52 bool
	var v82 vm.Value
	var v76 vm.Value
	var v70 vm.Value
	var callErr error
	b2 = arg3
	b1 = arg2
	rpo_idx = arg1
	idom = arg0
	goto b1
b1:
	;
	v20 = b1 == b2
	if v20 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v88 = b1
	goto b4
b3:
	;
	arg__9099, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{rpo_idx, b1})
	if callErr != nil {
		return nil, callErr
	}
	arg__9105, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{rpo_idx, b2})
	if callErr != nil {
		return nil, callErr
	}
	v35 = rt.GtValue(arg__9099, arg__9105)
	if v35 {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v88, nil
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{idom, b1})
	if callErr != nil {
		return nil, callErr
	}
	goto b1
b6:
	;
	arg__9116, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{rpo_idx, b2})
	if callErr != nil {
		return nil, callErr
	}
	arg__9122, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{rpo_idx, b1})
	if callErr != nil {
		return nil, callErr
	}
	v52 = rt.GtValue(arg__9116, arg__9122)
	if v52 {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	v88 = v82
	goto b4
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{idom, b2})
	if callErr != nil {
		return nil, callErr
	}
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
	v82 = v76
	goto b7
b11:
	;
	v70 = b1
	goto b13
b12:
	;
	v70 = vm.NIL
	goto b13
b13:
	;
	v76 = v70
	goto b10
}
func successors(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var term vm.Value
	var op vm.Value
	var aux vm.Value
	var v20 bool
	var arg__9149 vm.Value
	var v25 vm.Value
	var v38 bool
	var v81 vm.Value
	var t vm.Value
	var e vm.Value
	var arg__9162 vm.Value
	var arg__9166 vm.Value
	var v49 vm.Value
	var v74 vm.Value
	var callErr error
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, arg0})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{term, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v20 = op == vm.Keyword("branch")
	if v20 {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	arg__9149, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v25, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__9149})
	if callErr != nil {
		return nil, callErr
	}
	v81 = v25
	goto b3
b2:
	;
	v38 = op == vm.Keyword("branch-if")
	if v38 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v81, nil
b4:
	;
	t, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	e, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__9162, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__9166, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	v49, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__9162, arg__9166})
	if callErr != nil {
		return nil, callErr
	}
	v74 = v49
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
	v81 = v74
	goto b3
b7:
	;
	goto b9
b8:
	;
	goto b9
b9:
	;
	v74 = vm.NewArrayVector([]vm.Value{})
	goto b6
}
func dfs_postorder(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__9175 vm.Value
	var arg__9177 vm.Value
	var v13 vm.Value
	var v15 vm.Value
	var stack vm.Value
	var visited vm.Value
	var post vm.Value
	var f vm.Value
	var v29 vm.Value
	var top vm.Value
	var bid vm.Value
	var succs vm.Value
	var si vm.Value
	var arg__9206 vm.Value
	var v66 bool
	var v130 vm.Value
	var s vm.Value
	var arg__9222 vm.Value
	var arg__9229 vm.Value
	var arg__9230 vm.Value
	var arg__9236 vm.Value
	var stack_PRIME_ vm.Value
	var v107 vm.Value
	var v126 vm.Value
	var v128 vm.Value
	var arg__9250 vm.Value
	var arg__9262 vm.Value
	var arg__9264 vm.Value
	var v121 vm.Value
	var v123 vm.Value
	var callErr error
	arg__9175, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "successors").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__9177, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg1, arg__9175, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg__9177})
	if callErr != nil {
		return nil, callErr
	}
	v15, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	stack = v13
	visited = v15
	post = vm.NewArrayVector([]vm.Value{})
	f = arg0
	goto b1
b1:
	;
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{stack})
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
	v130 = post
	goto b4
b3:
	;
	top, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "peek").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	bid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{top, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	succs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{top, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	si, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{top, vm.Int(2)})
	if callErr != nil {
		return nil, callErr
	}
	arg__9206, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{succs})
	if callErr != nil {
		return nil, callErr
	}
	v66 = rt.LtValue(si, arg__9206)
	if v66 {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v130, nil
b5:
	;
	s, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{succs, si})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__9222 = rt.AddValue(si, vm.Int(1))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{bid, succs, arg__9222})
	if callErr != nil {
		return nil, callErr
	}
	arg__9229, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	arg__9230 = rt.SubValue(arg__9229, vm.Int(1))
	arg__9236, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{bid, succs, arg__9222})
	if callErr != nil {
		return nil, callErr
	}
	stack_PRIME_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{stack, arg__9230, arg__9236})
	if callErr != nil {
		return nil, callErr
	}
	v107, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{visited, s})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v107) {
		goto b8
	} else {
		goto b9
	}
b6:
	;
	v126, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pop").Deref(), []vm.Value{stack})
	if callErr != nil {
		return nil, callErr
	}
	v128, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{post, bid})
	if callErr != nil {
		return nil, callErr
	}
	stack = v126
	post = v128
	goto b1
b8:
	;
	stack = stack_PRIME_
	goto b1
b9:
	;
	arg__9250, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "successors").Deref(), []vm.Value{f, s})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{s, arg__9250, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__9262, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "successors").Deref(), []vm.Value{f, s})
	if callErr != nil {
		return nil, callErr
	}
	arg__9264, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{s, arg__9262, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v121, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{stack_PRIME_, arg__9264})
	if callErr != nil {
		return nil, callErr
	}
	v123, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{visited, s})
	if callErr != nil {
		return nil, callErr
	}
	stack = v121
	visited = v123
	goto b1
}
func reverse_postorder(arg0 vm.Value) (vm.Value, error) {
	var arg__9288 vm.Value
	var arg__9301 vm.Value
	var arg__9302 vm.Value
	var arg__9315 vm.Value
	var arg__9328 vm.Value
	var arg__9329 vm.Value
	var arg__9330 vm.Value
	var v31 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9288, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dfs-postorder").Deref(), []vm.Value{arg0, arg__9288})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9301, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9302, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dfs-postorder").Deref(), []vm.Value{arg0, arg__9301})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reverse").Deref(), []vm.Value{arg__9302})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9315, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dfs-postorder").Deref(), []vm.Value{arg0, arg__9315})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9328, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9329, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dfs-postorder").Deref(), []vm.Value{arg0, arg__9328})
	if callErr != nil {
		return nil, callErr
	}
	arg__9330, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reverse").Deref(), []vm.Value{arg__9329})
	if callErr != nil {
		return nil, callErr
	}
	v31, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__9330})
	if callErr != nil {
		return nil, callErr
	}
	return v31, nil
}
func refine_idom(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var preds vm.Value
	var ps vm.Value
	var new_id vm.Value
	var rpo_idx vm.Value
	var idom vm.Value
	var v28 vm.Value
	var p vm.Value
	var arg__9348 vm.Value
	var v52 bool
	var v134 vm.Value
	var v55 vm.Value
	var v74 bool
	var v77 vm.Value
	var v98 vm.Value
	var v100 vm.Value
	var callErr error
	preds, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	ps = preds
	new_id = vm.Int(-1)
	rpo_idx = arg3
	idom = arg2
	goto b1
b1:
	;
	v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v28) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v134 = new_id
	goto b4
b3:
	;
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	arg__9348, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{idom, p})
	if callErr != nil {
		return nil, callErr
	}
	v52 = arg__9348 == vm.Int(-1)
	if v52 {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v134, nil
b5:
	;
	v55, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	ps = v55
	goto b1
b6:
	;
	v74 = new_id == vm.Int(-1)
	if v74 {
		goto b8
	} else {
		goto b9
	}
b7:
	;
	v134 = vm.NIL
	goto b4
b8:
	;
	v77, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	ps = v77
	new_id = p
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
	v98, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{ps})
	if callErr != nil {
		return nil, callErr
	}
	v100, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "intersect").Deref(), []vm.Value{idom, rpo_idx, p, new_id})
	if callErr != nil {
		return nil, callErr
	}
	ps = v98
	new_id = v100
	goto b1
b12:
	;
	goto b13
b13:
	;
	goto b10
}
func dominators(arg0 vm.Value) (vm.Value, error) {
	var arg__9377 vm.Value
	var n vm.Value
	var entry vm.Value
	var rpo vm.Value
	var arg__9396 vm.Value
	var v25 vm.Value
	var i int
	var idx vm.Value
	var arg__9401 vm.Value
	var v41 bool
	var f vm.Value
	var v44 int
	var arg__9418 vm.Value
	var v50 vm.Value
	var rpo_idx vm.Value
	var arg__9432 vm.Value
	var arg__9449 vm.Value
	var arg__9450 vm.Value
	var idom0 vm.Value
	var idom vm.Value
	var bs vm.Value
	var changed_QMARK_ vm.Value
	var v119 vm.Value
	var v122 vm.Value
	var b vm.Value
	var v150 bool
	var step vm.Value
	var v283 vm.Value
	var v153 vm.Value
	var ni vm.Value
	var or__x_184 bool
	var v236 vm.Value
	var v239 vm.Value
	var or__x_198 bool
	var arg__9483 vm.Value
	var v217 bool
	var v219 bool
	var v288 vm.Value
	var v293 vm.Value
	var final vm.Value
	var v311 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__9377, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	n, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__9377})
	if callErr != nil {
		return nil, callErr
	}
	entry, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-entry").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	rpo, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "reverse-postorder").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__9396, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	v25, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__9396})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	idx = v25
	goto b1
b1:
	;
	arg__9401, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{rpo})
	if callErr != nil {
		return nil, callErr
	}
	v41 = rt.GeValue(vm.Int(i), arg__9401)
	if v41 {
		f = arg0
		goto b2
	} else {
		goto b3
	}
b2:
	;
	rpo_idx = idx
	goto b4
b3:
	;
	v44 = i + 1
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{rpo, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__9418, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{rpo, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v50, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{idx, arg__9418, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	i = v44
	idx = v50
	goto b1
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__9432, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__9432})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__9449, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	arg__9450, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__9449})
	if callErr != nil {
		return nil, callErr
	}
	idom0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg__9450, entry, entry})
	if callErr != nil {
		return nil, callErr
	}
	idom = idom0
	goto b5
b5:
	;
	bs = rpo
	changed_QMARK_ = vm.Boolean(false)
	goto b6
b6:
	;
	v119, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "empty?").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v119) {
		goto b7
	} else {
		goto b8
	}
b7:
	;
	v122, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{idom, changed_QMARK_})
	if callErr != nil {
		return nil, callErr
	}
	step = v122
	goto b9
b8:
	;
	b, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	v150 = b == entry
	if v150 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v283, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{step, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v283) {
		goto b19
	} else {
		goto b20
	}
b10:
	;
	v153, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	bs = v153
	goto b6
b11:
	;
	ni, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "refine-idom").Deref(), []vm.Value{f, b, idom, rpo_idx})
	if callErr != nil {
		return nil, callErr
	}
	or__x_184 = ni == vm.Int(-1)
	if or__x_184 {
		or__x_198 = or__x_184
		goto b16
	} else {
		goto b17
	}
b13:
	;
	v236, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	bs = v236
	goto b6
b14:
	;
	v239, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{bs})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{idom, b, ni})
	if callErr != nil {
		return nil, callErr
	}
	bs = v239
	changed_QMARK_ = vm.Boolean(true)
	goto b6
b16:
	;
	v219 = or__x_198
	goto b18
b17:
	;
	arg__9483, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{idom, b})
	if callErr != nil {
		return nil, callErr
	}
	v217 = arg__9483 == ni
	v219 = v217
	goto b18
b18:
	;
	if v219 {
		goto b13
	} else {
		goto b14
	}
b19:
	;
	v288, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{step, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	idom = v288
	goto b5
b20:
	;
	v293, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{step, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	final = v293
	goto b21
b21:
	;
	v311, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{final, entry, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	return v311, nil
}
func dominates_QMARK_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var idom vm.Value
	var b vm.Value
	var a vm.Value
	var v20 bool
	var v32 bool
	var v63 vm.Value
	var v57 vm.Value
	var callErr error
	idom, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominators").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	b = arg2
	a = arg1
	goto b1
b1:
	;
	v20 = b == vm.Int(-1)
	if v20 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v63 = vm.Boolean(false)
	goto b4
b3:
	;
	v32 = a == b
	if v32 {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	return v63, nil
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
	v63 = v57
	goto b4
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{idom, b})
	if callErr != nil {
		return nil, callErr
	}
	goto b1
b9:
	;
	goto b10
b10:
	;
	v57 = vm.NIL
	goto b7
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.dominance", map[string]vm.Value{"intersect": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("intersect: wrong number of arguments %d (expected 4)", len(args))
		}
		return intersect(args[0], args[1], args[2], args[3])
	}), "successors": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("successors: wrong number of arguments %d (expected 2)", len(args))
		}
		return successors(args[0], args[1])
	}), "dfs-postorder": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("dfs-postorder: wrong number of arguments %d (expected 2)", len(args))
		}
		return dfs_postorder(args[0], args[1])
	}), "reverse-postorder": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("reverse-postorder: wrong number of arguments %d (expected 1)", len(args))
		}
		return reverse_postorder(args[0])
	}), "refine-idom": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("refine-idom: wrong number of arguments %d (expected 4)", len(args))
		}
		return refine_idom(args[0], args[1], args[2], args[3])
	}), "dominators": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("dominators: wrong number of arguments %d (expected 1)", len(args))
		}
		return dominators(args[0])
	}), "dominates?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("dominates?: wrong number of arguments %d (expected 3)", len(args))
		}
		return dominates_QMARK_(args[0], args[1], args[2])
	}),
	})
}
