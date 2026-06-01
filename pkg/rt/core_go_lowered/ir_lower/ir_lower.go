package ir_lower

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func add_patch_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v9 vm.Value
	var callErr error
	v9, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "update").Deref(), vm.Keyword("patches"), rt.LookupVar("clojure.core", "conj").Deref(), arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v9, nil
}
func args_at_top_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var n vm.Value
	var v12 vm.Value
	var l vm.Value
	var args vm.Value
	var v161 vm.Value
	var arg__10954 vm.Value
	var base vm.Value
	var v36 vm.Value
	var v156 vm.Value
	var v148 vm.Value
	var i int
	var v57 bool
	var arg__10975 vm.Value
	var pos vm.Value
	var or__x vm.Value
	var v141 vm.Value
	var v134 vm.Value
	var arg__10982 vm.Value
	var v98 vm.Value
	var v100 vm.Value
	var v123 int
	var callErr error
	n, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{n})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		goto b1
	} else {
		l = arg0
		args = arg1
		goto b2
	}
b1:
	;
	v161 = vm.Boolean(true)
	goto b3
b2:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v161, nil
b4:
	;
	arg__10954, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	base = rt.SubValue(arg__10954, n)
	v36, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "neg?").Deref(), []vm.Value{base})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v36) {
		goto b7
	} else {
		goto b8
	}
b5:
	;
	v156 = vm.NIL
	goto b6
b6:
	;
	v161 = v156
	goto b3
b7:
	;
	v148 = vm.Boolean(false)
	goto b9
b8:
	;
	i = 0
	goto b10
b9:
	;
	v156 = v148
	goto b6
b10:
	;
	v57 = rt.GeValue(vm.Int(i), n)
	if v57 {
		goto b11
	} else {
		goto b12
	}
b11:
	;
	v141 = vm.Boolean(true)
	goto b13
b12:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__10975, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	pos, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "value-pos-of").Deref(), []vm.Value{l, arg__10975})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{pos})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b17
	} else {
		goto b18
	}
b13:
	;
	v148 = v141
	goto b9
b14:
	;
	v134 = vm.Boolean(false)
	goto b16
b15:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b20
	} else {
		goto b21
	}
b16:
	;
	v141 = v134
	goto b13
b17:
	;
	v100 = or__x
	goto b19
b18:
	;
	arg__10982 = rt.AddValue(base, vm.Int(i))
	v98, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{pos, arg__10982})
	if callErr != nil {
		return nil, callErr
	}
	v100 = v98
	goto b19
b19:
	;
	if vm.IsTruthy(v100) {
		goto b14
	} else {
		goto b15
	}
b20:
	;
	v123 = i + 1
	i = v123
	goto b10
b21:
	;
	goto b22
b22:
	;
	v134 = vm.NIL
	goto b16
}
func block_ip(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__10993 vm.Value
	var arg__11001 vm.Value
	var arg__11002 vm.Value
	var v11 vm.Value
	var callErr error
	arg__10993, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("block-ips"), []vm.Value{arg__10993})
	if callErr != nil {
		return nil, callErr
	}
	arg__11001, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__11002, callErr = rt.InvokeValue(vm.Keyword("block-ips"), []vm.Value{arg__11001})
	if callErr != nil {
		return nil, callErr
	}
	v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__11002, arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v11, nil
}
func block_junk_of(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__11008 vm.Value
	var arg__11017 vm.Value
	var arg__11018 vm.Value
	var v13 vm.Value
	var callErr error
	arg__11008, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("block-junk"), []vm.Value{arg__11008})
	if callErr != nil {
		return nil, callErr
	}
	arg__11017, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__11018, callErr = rt.InvokeValue(vm.Keyword("block-junk"), []vm.Value{arg__11017})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__11018, arg1, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func bump_max_stack_BANG_(arg0 vm.Value) (vm.Value, error) {
	var c vm.Value
	var arg__11027 vm.Value
	var arg__11033 vm.Value
	var arg__11041 vm.Value
	var arg__11042 vm.Value
	var arg__11043 vm.Value
	var rt_sp vm.Value
	var arg__11048 vm.Value
	var v25 bool
	var v28 vm.Value
	var v32 vm.Value
	var callErr error
	c, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__11027, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__11033, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("current-block"), []vm.Value{arg__11033})
	if callErr != nil {
		return nil, callErr
	}
	arg__11041, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__11042, callErr = rt.InvokeValue(vm.Keyword("current-block"), []vm.Value{arg__11041})
	if callErr != nil {
		return nil, callErr
	}
	arg__11043, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "block-junk-of").Deref(), []vm.Value{arg0, arg__11042})
	if callErr != nil {
		return nil, callErr
	}
	rt_sp = rt.AddValue(arg__11027, arg__11043)
	arg__11048, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-max-stack").Deref(), []vm.Value{c})
	if callErr != nil {
		return nil, callErr
	}
	v25 = rt.GtValue(rt_sp, arg__11048)
	if v25 {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v28, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-set-max-stack!").Deref(), []vm.Value{c, rt_sp})
	if callErr != nil {
		return nil, callErr
	}
	v32 = v28
	goto b3
b2:
	;
	v32 = vm.NIL
	goto b3
b3:
	;
	return v32, nil
}
func bump_stack_sp_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v9 vm.Value
	var callErr error
	v9, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "update").Deref(), vm.Keyword("stack-sp"), rt.LookupVar("clojure.core", "+").Deref(), arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v9, nil
}
func check_cross_block_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var i int
	var uses vm.Value
	var f vm.Value
	var arg__11069 vm.Value
	var v16 bool
	var us vm.Value
	var v29 vm.Value
	var def_block vm.Value
	var v53 int
	var callErr error
	i = 0
	uses = arg1
	f = arg0
	goto b1
b1:
	;
	arg__11069, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{uses})
	if callErr != nil {
		return nil, callErr
	}
	v16 = rt.LtValue(vm.Int(i), arg__11069)
	if v16 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	us, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v29) {
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
	goto b7
b6:
	;
	def_block, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{vm.Int(i), f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-for-each").Deref(), []vm.Value{us, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var user_block vm.Value
		var and__x vm.Value
		var def_block_7 vm.Value
		var i_9 vm.Value
		var arg__11214 vm.Value
		var v72 vm.Value
		var v76 vm.Value
		var user_id vm.Value
		var def_block_19 vm.Value
		var f_20 vm.Value
		var i_21 vm.Value
		var arg__11177 vm.Value
		var v36 vm.Value
		var def_block_25 vm.Value
		var i_27 vm.Value
		var v39 vm.Value
		var def_block_41 vm.Value
		var i_43 vm.Value
		var callErr error
		user_block, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{arg0, f})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{user_block, def_block})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			user_id = arg0
			def_block_19 = def_block
			f_20 = f
			i_21 = vm.Int(i)
			goto b4
		} else {
			def_block_25 = def_block
			i_27 = vm.Int(i)
			goto b5
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/lower: cross-block use of value %"), i_9, vm.String(" via direct Refs (defined in block "), def_block_7, vm.String(", used in block "), user_block, vm.String(") — only branch-target args may cross blocks; "), vm.String("see follow-up issue D")})
		if callErr != nil {
			return nil, callErr
		}
		arg__11214, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/lower: cross-block use of value %"), i_9, vm.String(" via direct Refs (defined in block "), def_block_7, vm.String(", used in block "), user_block, vm.String(") — only branch-target args may cross blocks; "), vm.String("see follow-up issue D")})
		if callErr != nil {
			return nil, callErr
		}
		v72, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__11214})
		if callErr != nil {
			return nil, callErr
		}
		v76 = v72
		goto b3
	b2:
		;
		v76 = vm.NIL
		goto b3
	b3:
		;
		return v76, nil
	b4:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "is-terminator-branch-arg-use?").Deref(), []vm.Value{user_id, i_21, f_20})
		if callErr != nil {
			return nil, callErr
		}
		arg__11177, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "is-terminator-branch-arg-use?").Deref(), []vm.Value{user_id, i_21, f_20})
		if callErr != nil {
			return nil, callErr
		}
		v36, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__11177})
		if callErr != nil {
			return nil, callErr
		}
		v39 = v36
		def_block_41 = def_block_19
		i_43 = i_21
		goto b6
	b5:
		;
		v39 = and__x
		def_block_41 = def_block_25
		i_43 = i_27
		goto b6
	b6:
		;
		if vm.IsTruthy(v39) {
			def_block_7 = def_block_41
			i_9 = i_43
			goto b1
		} else {
			goto b2
		}
	})})
	if callErr != nil {
		return nil, callErr
	}
	goto b7
b7:
	;
	v53 = i + 1
	i = v53
	goto b1
}
func chunk_of(arg0 vm.Value) (vm.Value, error) {
	var arg__11221 vm.Value
	var v4 vm.Value
	var callErr error
	arg__11221, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("chunk"), []vm.Value{arg__11221})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func consume_refs_in_place_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var doseq_seq__11222 vm.Value
	var doseq_loop__11223 vm.Value
	var l vm.Value
	var r vm.Value
	var v22 vm.Value
	var callErr error
	doseq_seq__11222, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__11223 = doseq_seq__11222
	l = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__11223) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__11223})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "decrement-use!").Deref(), []vm.Value{l, r})
	if callErr != nil {
		return nil, callErr
	}
	v22, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__11223})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__11223 = v22
	goto b1
b3:
	;
	goto b4
b4:
	;
	return vm.NIL, nil
}
func contains_after_QMARK_(arg0 vm.Value, arg1 int, arg2 vm.Value) (vm.Value, error) {
	var v7 int
	var j int
	var refs vm.Value
	var target vm.Value
	var arg__11243 vm.Value
	var v19 bool
	var arg__11249 vm.Value
	var v33 bool
	var v63 vm.Value
	var v57 vm.Value
	var v47 int
	var callErr error
	v7 = arg1 + 1
	j = v7
	refs = arg0
	target = arg2
	goto b1
b1:
	;
	arg__11243, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v19 = rt.GeValue(vm.Int(j), arg__11243)
	if v19 {
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
	arg__11249, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(j)})
	if callErr != nil {
		return nil, callErr
	}
	v33 = arg__11249 == target
	if v33 {
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
	v47 = j + 1
	j = v47
	goto b1
b9:
	;
	goto b10
b10:
	;
	v57 = vm.NIL
	goto b7
}
func decrement_use_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var cur vm.Value
	var v12 vm.Value
	var l vm.Value
	var nid vm.Value
	var arg__11266 vm.Value
	var v23 vm.Value
	var v27 vm.Value
	var callErr error
	cur, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "use-count-of").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{cur})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		l = arg0
		nid = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	arg__11266 = rt.SubValue(cur, vm.Int(1))
	v23, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{l, rt.LookupVar("clojure.core", "update").Deref(), vm.Keyword("use-count"), rt.LookupVar("clojure.core", "assoc").Deref(), nid, arg__11266})
	if callErr != nil {
		return nil, callErr
	}
	v27 = v23
	goto b3
b2:
	;
	v27 = vm.NIL
	goto b3
b3:
	;
	return v27, nil
}
func deferrable_branch_if_cond_QMARK_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var l vm.Value
	var term vm.Value
	var cond_ref vm.Value
	var arg__11281 vm.Value
	var and__x_14 bool
	var v116 vm.Value
	var arg__11286 vm.Value
	var uses vm.Value
	var arg__11291 vm.Value
	var v40 bool
	var and__x_22 bool
	var v108 vm.Value
	var v43 vm.Value
	var us vm.Value
	var arg__11305 vm.Value
	var and__x_71 vm.Value
	var and__x_59 vm.Value
	var v99 vm.Value
	var arg__11310 vm.Value
	var v87 bool
	var and__x_83 vm.Value
	var v90 vm.Value
	var callErr error
	if vm.IsTruthy(arg2) {
		l = arg0
		term = arg1
		cond_ref = arg2
		goto b1
	} else {
		goto b2
	}
b1:
	;
	arg__11281, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "use-count-of").Deref(), []vm.Value{l, cond_ref})
	if callErr != nil {
		return nil, callErr
	}
	and__x_14 = arg__11281 == vm.Int(1)
	if and__x_14 {
		goto b4
	} else {
		and__x_22 = and__x_14
		goto b5
	}
b2:
	;
	v116 = vm.NIL
	goto b3
b3:
	;
	return v116, nil
b4:
	;
	arg__11286, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	uses, callErr = rt.InvokeValue(vm.Keyword("uses"), []vm.Value{arg__11286})
	if callErr != nil {
		return nil, callErr
	}
	arg__11291, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{uses})
	if callErr != nil {
		return nil, callErr
	}
	v40 = rt.LtValue(cond_ref, arg__11291)
	if v40 {
		goto b7
	} else {
		goto b8
	}
b5:
	;
	v108 = vm.Boolean(and__x_22)
	goto b6
b6:
	;
	v116 = v108
	goto b3
b7:
	;
	v43, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, cond_ref})
	if callErr != nil {
		return nil, callErr
	}
	us = v43
	goto b9
b8:
	;
	us = vm.NIL
	goto b9
b9:
	;
	if vm.IsTruthy(us) {
		goto b10
	} else {
		and__x_59 = us
		goto b11
	}
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	arg__11305, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	and__x_71, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__11305})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x_71) {
		goto b13
	} else {
		and__x_83 = and__x_71
		goto b14
	}
b11:
	;
	v99 = and__x_59
	goto b12
b12:
	;
	v108 = v99
	goto b6
b13:
	;
	arg__11310, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-first").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	v87 = term == arg__11310
	v90 = vm.Boolean(v87)
	goto b15
b14:
	;
	v90 = and__x_83
	goto b15
b15:
	;
	v99 = v90
	goto b12
}
func emit_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var l vm.Value
	var inst_id vm.Value
	var op_kw vm.Value
	var arg__11329 vm.Value
	var arg__11334 vm.Value
	var v28 vm.Value
	var callErr error
	if vm.IsTruthy(arg1) {
		l = arg0
		inst_id = arg1
		op_kw = arg2
		goto b1
	} else {
		l = arg0
		op_kw = arg2
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "record-source-info!").Deref(), []vm.Value{l, inst_id})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11329, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11334, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v28, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-emit").Deref(), []vm.Value{arg__11329, op_kw, arg__11334})
	if callErr != nil {
		return nil, callErr
	}
	return v28, nil
}
func emit_inst_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var f vm.Value
	var op vm.Value
	var refs vm.Value
	var v20 vm.Value
	var l vm.Value
	var nid vm.Value
	var arg__11372 vm.Value
	var v49 bool
	var arg__11386 vm.Value
	var arg__11387 vm.Value
	var v58 vm.Value
	var v62 vm.Value
	var callErr error
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg1, f})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg1, f})
	if callErr != nil {
		return nil, callErr
	}
	v20, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "refs-at-top-last-use?").Deref(), []vm.Value{arg0, refs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v20) {
		l = arg0
		nid = arg1
		goto b1
	} else {
		l = arg0
		nid = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "consume-refs-in-place!").Deref(), []vm.Value{l, refs})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "materialize-refs!").Deref(), []vm.Value{l, refs})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "lower-node!").Deref(), []vm.Value{l, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__11372, callErr = rt.InvokeValue(rt.LookupVar("ir", "op-stack-out").Deref(), []vm.Value{op})
	if callErr != nil {
		return nil, callErr
	}
	v49 = arg__11372 == vm.Int(1)
	if v49 {
		goto b4
	} else {
		goto b5
	}
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11386, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11387 = rt.SubValue(arg__11386, vm.Int(1))
	v58, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "set-value-pos!").Deref(), []vm.Value{l, nid, arg__11387})
	if callErr != nil {
		return nil, callErr
	}
	v62 = v58
	goto b6
b5:
	;
	v62 = vm.NIL
	goto b6
b6:
	;
	return v62, nil
}
func emit_placeholder_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var l vm.Value
	var inst_id vm.Value
	var op_kw vm.Value
	var arg__11406 vm.Value
	var arg__11411 vm.Value
	var v29 vm.Value
	var callErr error
	if vm.IsTruthy(arg1) {
		l = arg0
		inst_id = arg1
		op_kw = arg2
		goto b1
	} else {
		l = arg0
		op_kw = arg2
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "record-source-info!").Deref(), []vm.Value{l, inst_id})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11406, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11411, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-emit-placeholder").Deref(), []vm.Value{arg__11406, op_kw, arg__11411})
	if callErr != nil {
		return nil, callErr
	}
	return v29, nil
}
func emit_with_arg_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value, arg3 vm.Value) (vm.Value, error) {
	var l vm.Value
	var inst_id vm.Value
	var op_kw vm.Value
	var arg vm.Value
	var arg__11431 vm.Value
	var arg__11436 vm.Value
	var v32 vm.Value
	var callErr error
	if vm.IsTruthy(arg1) {
		l = arg0
		inst_id = arg1
		op_kw = arg2
		arg = arg3
		goto b1
	} else {
		l = arg0
		op_kw = arg2
		arg = arg3
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "record-source-info!").Deref(), []vm.Value{l, inst_id})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11431, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11436, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v32, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-emit-with-arg").Deref(), []vm.Value{arg__11431, op_kw, arg__11436, arg})
	if callErr != nil {
		return nil, callErr
	}
	return v32, nil
}
func f_of(arg0 vm.Value) (vm.Value, error) {
	var arg__11442 vm.Value
	var v4 vm.Value
	var callErr error
	arg__11442, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("f"), []vm.Value{arg__11442})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func is_terminator_branch_arg_use_QMARK_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var op vm.Value
	var aux vm.Value
	var v19 bool
	var target_id vm.Value
	var arg__11469 vm.Value
	var v32 vm.Value
	var v45 bool
	var v134 vm.Value
	var t vm.Value
	var e vm.Value
	var arg__11492 vm.Value
	var or__x vm.Value
	var v127 vm.Value
	var arg__11507 vm.Value
	var v92 vm.Value
	var v94 vm.Value
	var v120 vm.Value
	var callErr error
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg2})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg0, arg2})
	if callErr != nil {
		return nil, callErr
	}
	v19 = op == vm.Keyword("branch")
	if v19 {
		target_id = arg1
		goto b1
	} else {
		target_id = arg1
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__11469, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v32, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var v2 vm.Value
		v2 = vm.Boolean(arg0 == target_id)
		return v2
	}), arg__11469})
	if callErr != nil {
		return nil, callErr
	}
	v134 = v32
	goto b3
b2:
	;
	v45 = op == vm.Keyword("branch-if")
	if v45 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v134, nil
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
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	arg__11492, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{t})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var v2 vm.Value
		v2 = vm.Boolean(arg0 == target_id)
		return v2
	}), arg__11492})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b7
	} else {
		goto b8
	}
b5:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b10
	} else {
		goto b11
	}
b6:
	;
	v134 = v127
	goto b3
b7:
	;
	v94 = or__x
	goto b9
b8:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	arg__11507, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{e})
	if callErr != nil {
		return nil, callErr
	}
	v92, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
		var v2 vm.Value
		v2 = vm.Boolean(arg0 == target_id)
		return v2
	}), arg__11507})
	if callErr != nil {
		return nil, callErr
	}
	v94 = v92
	goto b9
b9:
	;
	v127 = v94
	goto b6
b10:
	;
	v120 = vm.Boolean(false)
	goto b12
b11:
	;
	v120 = vm.NIL
	goto b12
b12:
	;
	v127 = v120
	goto b6
}
func lower(arg0 vm.Value) (vm.Value, error) {
	var v5 vm.Value
	var f vm.Value
	var l vm.Value
	var arg__11522 vm.Value
	var arg__11530 vm.Value
	var arg__11531 vm.Value
	var arg__11540 vm.Value
	var n_blocks vm.Value
	var bid int
	var v47 bool
	var arg__11553 vm.Value
	var arg__11566 vm.Value
	var arg__11567 vm.Value
	var v65 int
	var c vm.Value
	var arg__11588 vm.Value
	var acc vm.Value
	var v104 bool
	var arg__11603 vm.Value
	var p vm.Value
	var v128 bool
	var max_params vm.Value
	var arg__11619 vm.Value
	var arg__11621 vm.Value
	var v157 vm.Value
	var v132 vm.Value
	var callErr error
	v5, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v5) {
		f = arg0
		goto b1
	} else {
		f = arg0
		goto b2
	}
b1:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{vm.String("ir/lower: nil function")})
	if callErr != nil {
		return nil, callErr
	}
	goto b3
b2:
	;
	goto b3
b3:
	;
	l, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "new-lowerer").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__11522, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("uses"), []vm.Value{arg__11522})
	if callErr != nil {
		return nil, callErr
	}
	arg__11530, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11531, callErr = rt.InvokeValue(vm.Keyword("uses"), []vm.Value{arg__11530})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "check-cross-block!").Deref(), []vm.Value{f, arg__11531})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__11540, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	n_blocks, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__11540})
	if callErr != nil {
		return nil, callErr
	}
	bid = 0
	goto b4
b4:
	;
	v47 = rt.LtValue(vm.Int(bid), n_blocks)
	if v47 {
		goto b5
	} else {
		goto b6
	}
b5:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11553, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-length").Deref(), []vm.Value{arg__11553})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11566, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11567, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-length").Deref(), []vm.Value{arg__11566})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "set-block-ip!").Deref(), []vm.Value{l, vm.Int(bid), arg__11567})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "lower-block!").Deref(), []vm.Value{l, vm.Int(bid)})
	if callErr != nil {
		return nil, callErr
	}
	v65 = bid + 1
	bid = v65
	goto b4
b6:
	;
	goto b7
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "patch-branches!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	c, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__11588, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	n_blocks, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__11588})
	if callErr != nil {
		return nil, callErr
	}
	acc = vm.Int(0)
	goto b8
b8:
	;
	v104 = rt.GeValue(vm.Int(bid), n_blocks)
	if v104 {
		goto b9
	} else {
		goto b10
	}
b9:
	;
	max_params = acc
	goto b11
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{vm.Int(bid), f})
	if callErr != nil {
		return nil, callErr
	}
	arg__11603, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{vm.Int(bid), f})
	if callErr != nil {
		return nil, callErr
	}
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__11603})
	if callErr != nil {
		return nil, callErr
	}
	v128 = rt.GtValue(p, acc)
	if v128 {
		goto b12
	} else {
		goto b13
	}
b11:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-max-stack").Deref(), []vm.Value{c})
	if callErr != nil {
		return nil, callErr
	}
	arg__11619, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-max-stack").Deref(), []vm.Value{c})
	if callErr != nil {
		return nil, callErr
	}
	arg__11621 = rt.AddValue(arg__11619, max_params)
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-set-max-stack!").Deref(), []vm.Value{c, arg__11621})
	if callErr != nil {
		return nil, callErr
	}
	v157, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	return v157, nil
b12:
	;
	v132 = p
	goto b14
b13:
	;
	v132 = acc
	goto b14
b14:
	;
	acc = v132
	goto b8
}
func lower_block_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var f vm.Value
	var params vm.Value
	var insts vm.Value
	var term vm.Value
	var l vm.Value
	var v61 vm.Value
	var term_op vm.Value
	var v87 bool
	var arg__11664 vm.Value
	var v48 vm.Value
	var and__x vm.Value
	var v51 vm.Value
	var arg__11684 vm.Value
	var cond_ref vm.Value
	var v112 vm.Value
	var deferred_cond vm.Value
	var arg__11702 vm.Value
	var v117 vm.Value
	var i int
	var arg__11707 vm.Value
	var v168 bool
	var arg__11723 vm.Value
	var v176 int
	var doseq_seq__11625 vm.Value
	var doseq_loop__11626 vm.Value
	var nid vm.Value
	var op vm.Value
	var v251 bool
	var v282 bool
	var v663 vm.Value
	var v313 bool
	var arg__11755 vm.Value
	var doseq_seq__11627 vm.Value
	var doseq_loop__11628 vm.Value
	var r vm.Value
	var v360 vm.Value
	var v436 bool
	var v439 vm.Value
	var arg__11784 vm.Value
	var v519 vm.Value
	var v522 vm.Value
	var v764 bool
	var v1067 vm.Value
	var arg__11801 vm.Value
	var v727 vm.Value
	var v730 vm.Value
	var cond_aux vm.Value
	var t_bt vm.Value
	var args vm.Value
	var v1053 vm.Value
	var v919 vm.Value
	var term_refs vm.Value
	var v863 vm.Value
	var v968 vm.Value
	var v1041 vm.Value
	var v1011 bool
	var bt vm.Value
	var arg__11887 vm.Value
	var v1037 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc").Deref(), vm.Keyword("current-block"), arg1})
	if callErr != nil {
		return nil, callErr
	}
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	params, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{arg1, f})
	if callErr != nil {
		return nil, callErr
	}
	insts, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-insts").Deref(), []vm.Value{arg1, f})
	if callErr != nil {
		return nil, callErr
	}
	term, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-term").Deref(), []vm.Value{arg1, f})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(term) {
		l = arg0
		goto b4
	} else {
		l = arg0
		and__x = term
		goto b5
	}
b1:
	;
	v61, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	term_op = v61
	goto b3
b2:
	;
	term_op = vm.NIL
	goto b3
b3:
	;
	v87 = term_op == vm.Keyword("branch-if")
	if v87 {
		goto b7
	} else {
		goto b8
	}
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	arg__11664, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	v48, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__11664})
	if callErr != nil {
		return nil, callErr
	}
	v51 = v48
	goto b6
b5:
	;
	v51 = and__x
	goto b6
b6:
	;
	if vm.IsTruthy(v51) {
		goto b1
	} else {
		goto b2
	}
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__11684, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	cond_ref, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{arg__11684})
	if callErr != nil {
		return nil, callErr
	}
	v112, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "deferrable-branch-if-cond?").Deref(), []vm.Value{l, term, cond_ref})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v112) {
		goto b10
	} else {
		goto b11
	}
b8:
	;
	deferred_cond = vm.NIL
	goto b9
b9:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	arg__11702, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "set-stack-sp!").Deref(), []vm.Value{l, arg__11702})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	goto b13
b10:
	;
	v117 = cond_ref
	goto b12
b11:
	;
	v117 = vm.NIL
	goto b12
b12:
	;
	deferred_cond = v117
	goto b9
b13:
	;
	arg__11707, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{params})
	if callErr != nil {
		return nil, callErr
	}
	v168 = rt.LtValue(vm.Int(i), arg__11707)
	if v168 {
		goto b14
	} else {
		goto b15
	}
b14:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{params, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__11723, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{params, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "set-value-pos!").Deref(), []vm.Value{l, arg__11723, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v176 = i + 1
	i = v176
	goto b13
b15:
	;
	goto b16
b16:
	;
	doseq_seq__11625, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{insts})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__11626 = doseq_seq__11625
	goto b17
b17:
	;
	if vm.IsTruthy(doseq_loop__11626) {
		goto b18
	} else {
		goto b19
	}
b18:
	;
	nid, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__11626})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	v251 = op == vm.Keyword("invalid")
	if v251 {
		goto b21
	} else {
		goto b22
	}
b19:
	;
	goto b20
b20:
	;
	if vm.IsTruthy(term) {
		goto b52
	} else {
		and__x = term
		goto b53
	}
b21:
	;
	goto b23
b22:
	;
	v282 = op == vm.Keyword("block-arg")
	if v282 {
		goto b24
	} else {
		goto b25
	}
b23:
	;
	v663, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__11626})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__11626 = v663
	goto b17
b24:
	;
	goto b26
b25:
	;
	v313 = op == vm.Keyword("pop")
	if v313 {
		goto b27
	} else {
		goto b28
	}
b26:
	;
	goto b23
b27:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__11755, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__11627, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__11755})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__11628 = doseq_seq__11627
	goto b30
b28:
	;
	if vm.IsTruthy(deferred_cond) {
		goto b37
	} else {
		and__x = deferred_cond
		goto b38
	}
b29:
	;
	goto b26
b30:
	;
	if vm.IsTruthy(doseq_loop__11628) {
		goto b31
	} else {
		goto b32
	}
b31:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__11628})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "decrement-use!").Deref(), []vm.Value{l, r})
	if callErr != nil {
		return nil, callErr
	}
	v360, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__11628})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__11628 = v360
	goto b30
b32:
	;
	goto b33
b33:
	;
	goto b29
b34:
	;
	goto b36
b35:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("ir", "op-cheap-load?").Deref(), []vm.Value{op})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		goto b43
	} else {
		goto b44
	}
b36:
	;
	goto b29
b37:
	;
	v436 = nid == deferred_cond
	v439 = vm.Boolean(v436)
	goto b39
b38:
	;
	v439 = and__x
	goto b39
b39:
	;
	if vm.IsTruthy(v439) {
		goto b34
	} else {
		goto b35
	}
b40:
	;
	goto b42
b41:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b46
	} else {
		goto b47
	}
b42:
	;
	goto b36
b43:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "should-body-emit-cheap?").Deref(), []vm.Value{l, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__11784, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "should-body-emit-cheap?").Deref(), []vm.Value{l, nid})
	if callErr != nil {
		return nil, callErr
	}
	v519, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__11784})
	if callErr != nil {
		return nil, callErr
	}
	v522 = v519
	goto b45
b44:
	;
	v522 = and__x
	goto b45
b45:
	;
	if vm.IsTruthy(v522) {
		goto b40
	} else {
		goto b41
	}
b46:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-inst!").Deref(), []vm.Value{l, nid})
	if callErr != nil {
		return nil, callErr
	}
	goto b48
b47:
	;
	goto b48
b48:
	;
	goto b42
b49:
	;
	v764 = term_op == vm.Keyword("branch-if")
	if v764 {
		goto b55
	} else {
		goto b56
	}
b50:
	;
	v1067 = vm.NIL
	goto b51
b51:
	;
	return v1067, nil
b52:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	arg__11801, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{term})
	if callErr != nil {
		return nil, callErr
	}
	v727, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__11801})
	if callErr != nil {
		return nil, callErr
	}
	v730 = v727
	goto b54
b53:
	;
	v730 = and__x
	goto b54
b54:
	;
	if vm.IsTruthy(v730) {
		goto b49
	} else {
		goto b50
	}
b55:
	;
	cond_aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	t_bt, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{cond_aux})
	if callErr != nil {
		return nil, callErr
	}
	args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{t_bt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "materialize-branch-args!").Deref(), []vm.Value{l, args})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(deferred_cond) {
		goto b58
	} else {
		goto b59
	}
b56:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b67
	} else {
		goto b68
	}
b57:
	;
	v1067 = v1053
	goto b51
b58:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-inst!").Deref(), []vm.Value{l, deferred_cond})
	if callErr != nil {
		return nil, callErr
	}
	goto b60
b59:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b61
	} else {
		goto b62
	}
b60:
	;
	v919, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "lower-node!").Deref(), []vm.Value{l, term})
	if callErr != nil {
		return nil, callErr
	}
	v1053 = v919
	goto b57
b61:
	;
	term_refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	v863, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "refs-at-top-last-use?").Deref(), []vm.Value{l, term_refs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v863) {
		goto b64
	} else {
		goto b65
	}
b62:
	;
	goto b63
b63:
	;
	goto b60
b64:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "consume-refs-in-place!").Deref(), []vm.Value{l, term_refs})
	if callErr != nil {
		return nil, callErr
	}
	goto b66
b65:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "materialize-refs!").Deref(), []vm.Value{l, term_refs})
	if callErr != nil {
		return nil, callErr
	}
	goto b66
b66:
	;
	goto b63
b67:
	;
	term_refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	v968, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "refs-at-top-last-use?").Deref(), []vm.Value{l, term_refs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v968) {
		goto b70
	} else {
		goto b71
	}
b68:
	;
	v1041 = vm.NIL
	goto b69
b69:
	;
	v1053 = v1041
	goto b57
b70:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "consume-refs-in-place!").Deref(), []vm.Value{l, term_refs})
	if callErr != nil {
		return nil, callErr
	}
	goto b72
b71:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "materialize-refs!").Deref(), []vm.Value{l, term_refs})
	if callErr != nil {
		return nil, callErr
	}
	goto b72
b72:
	;
	v1011 = term_op == vm.Keyword("branch")
	if v1011 {
		goto b73
	} else {
		goto b74
	}
b73:
	;
	bt, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{term, f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	arg__11887, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{bt})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "materialize-branch-args!").Deref(), []vm.Value{l, arg__11887})
	if callErr != nil {
		return nil, callErr
	}
	goto b75
b74:
	;
	goto b75
b75:
	;
	v1037, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "lower-node!").Deref(), []vm.Value{l, term})
	if callErr != nil {
		return nil, callErr
	}
	v1041 = v1037
	goto b69
}
func lower_node_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var f vm.Value
	var op vm.Value
	var aux vm.Value
	var v22 bool
	var l vm.Value
	var nid vm.Value
	var case__11893 vm.Value
	var v39 bool
	var v1057 vm.Value
	var arg__11923 vm.Value
	var arg__11927 vm.Value
	var idx vm.Value
	var v60 vm.Value
	var v75 bool
	var v1049 vm.Value
	var arg__11961 vm.Value
	var v90 vm.Value
	var v105 bool
	var v1041 vm.Value
	var arg__11981 vm.Value
	var v122 vm.Value
	var or__x_137 bool
	var v1033 vm.Value
	var v333 vm.Value
	var or__x_348 bool
	var v1025 vm.Value
	var or__x_144 bool
	var or__x_155 bool
	var v319 bool
	var or__x_162 bool
	var or__x_173 bool
	var v310 bool
	var or__x_180 bool
	var or__x_191 bool
	var v301 bool
	var or__x_198 bool
	var or__x_209 bool
	var v292 bool
	var or__x_216 bool
	var or__x_227 bool
	var v283 bool
	var or__x_234 bool
	var or__x_245 bool
	var v274 bool
	var or__x_252 bool
	var v263 bool
	var v265 bool
	var v378 vm.Value
	var v393 bool
	var v1017 vm.Value
	var or__x_355 bool
	var v366 bool
	var v368 bool
	var arg__12055 vm.Value
	var arg__12067 vm.Value
	var arg__12068 vm.Value
	var v412 vm.Value
	var v427 bool
	var v1009 vm.Value
	var v436 vm.Value
	var v451 bool
	var v1001 vm.Value
	var arg__12099 vm.Value
	var v460 vm.Value
	var v475 bool
	var v993 vm.Value
	var v480 vm.Value
	var v495 bool
	var v985 vm.Value
	var args vm.Value
	var target vm.Value
	var argc vm.Value
	var arg__12132 vm.Value
	var target_params vm.Value
	var cur_sp vm.Value
	var drop_count vm.Value
	var arg__12143 vm.Value
	var arg__12151 vm.Value
	var arg__12152 vm.Value
	var cur_junk vm.Value
	var v551 vm.Value
	var v689 bool
	var v977 vm.Value
	var ignore vm.Value
	var arg__12169 vm.Value
	var off_ip vm.Value
	var arg__12182 vm.Value
	var arg__12202 vm.Value
	var v583 vm.Value
	var v660 vm.Value
	var arg_ip vm.Value
	var arg__12226 vm.Value
	var arg__12246 vm.Value
	var v640 vm.Value
	var v644 vm.Value
	var ft vm.Value
	var tt vm.Value
	var ft_target vm.Value
	var tt_target vm.Value
	var arg__12273 vm.Value
	var arg__12281 vm.Value
	var arg__12282 vm.Value
	var arg__12286 vm.Value
	var arg__12288 vm.Value
	var arg__12302 vm.Value
	var arg__12303 vm.Value
	var v723 vm.Value
	var target_junk vm.Value
	var arg__12322 vm.Value
	var arg__12342 vm.Value
	var my_block vm.Value
	var next_block_id vm.Value
	var v789 vm.Value
	var v849 bool
	var v969 vm.Value
	var arg_ip2 vm.Value
	var arg__12370 vm.Value
	var arg__12390 vm.Value
	var v814 vm.Value
	var v818 vm.Value
	var arg__12407 vm.Value
	var v864 vm.Value
	var v879 bool
	var v961 vm.Value
	var v884 vm.Value
	var v899 bool
	var v953 vm.Value
	var v908 vm.Value
	var v945 vm.Value
	var arg__12451 vm.Value
	var v933 vm.Value
	var v937 vm.Value
	var callErr error
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg1, f})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg1, f})
	if callErr != nil {
		return nil, callErr
	}
	v22 = op == vm.Keyword("block-arg")
	if v22 {
		goto b1
	} else {
		l = arg0
		nid = arg1
		case__11893 = op
		goto b2
	}
b1:
	;
	v1057 = vm.NIL
	goto b3
b2:
	;
	v39 = case__11893 == vm.Keyword("const")
	if v39 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v1057, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "template-value").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__11923, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11927, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "template-value").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	idx, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-intern-const").Deref(), []vm.Value{arg__11923, arg__11927})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("const"), idx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v60, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v1049 = v60
	goto b6
b5:
	;
	v75 = case__11893 == vm.Keyword("load-arg")
	if v75 {
		goto b7
	} else {
		goto b8
	}
b6:
	;
	v1057 = v1049
	goto b3
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__11961, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("load-arg"), arg__11961})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v90, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v1041 = v90
	goto b9
b8:
	;
	v105 = case__11893 == vm.Keyword("load-var")
	if v105 {
		goto b10
	} else {
		goto b11
	}
b9:
	;
	v1049 = v1041
	goto b6
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__11981, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	idx, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-intern-const").Deref(), []vm.Value{arg__11981, aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("load-var"), idx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v122, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v1033 = v122
	goto b12
b11:
	;
	or__x_137 = case__11893 == vm.Keyword("add")
	if or__x_137 {
		or__x_144 = or__x_137
		goto b16
	} else {
		goto b17
	}
b12:
	;
	v1041 = v1033
	goto b9
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit!").Deref(), []vm.Value{l, nid, op})
	if callErr != nil {
		return nil, callErr
	}
	v333, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	v1025 = v333
	goto b15
b14:
	;
	or__x_348 = case__11893 == vm.Keyword("inc")
	if or__x_348 {
		or__x_355 = or__x_348
		goto b40
	} else {
		goto b41
	}
b15:
	;
	v1033 = v1025
	goto b12
b16:
	;
	v319 = or__x_144
	goto b18
b17:
	;
	or__x_155 = case__11893 == vm.Keyword("sub")
	if or__x_155 {
		or__x_162 = or__x_155
		goto b19
	} else {
		goto b20
	}
b18:
	;
	if v319 {
		goto b13
	} else {
		goto b14
	}
b19:
	;
	v310 = or__x_162
	goto b21
b20:
	;
	or__x_173 = case__11893 == vm.Keyword("mul")
	if or__x_173 {
		or__x_180 = or__x_173
		goto b22
	} else {
		goto b23
	}
b21:
	;
	v319 = v310
	goto b18
b22:
	;
	v301 = or__x_180
	goto b24
b23:
	;
	or__x_191 = case__11893 == vm.Keyword("lt")
	if or__x_191 {
		or__x_198 = or__x_191
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v310 = v301
	goto b21
b25:
	;
	v292 = or__x_198
	goto b27
b26:
	;
	or__x_209 = case__11893 == vm.Keyword("lte")
	if or__x_209 {
		or__x_216 = or__x_209
		goto b28
	} else {
		goto b29
	}
b27:
	;
	v301 = v292
	goto b24
b28:
	;
	v283 = or__x_216
	goto b30
b29:
	;
	or__x_227 = case__11893 == vm.Keyword("gt")
	if or__x_227 {
		or__x_234 = or__x_227
		goto b31
	} else {
		goto b32
	}
b30:
	;
	v292 = v283
	goto b27
b31:
	;
	v274 = or__x_234
	goto b33
b32:
	;
	or__x_245 = case__11893 == vm.Keyword("gte")
	if or__x_245 {
		or__x_252 = or__x_245
		goto b34
	} else {
		goto b35
	}
b33:
	;
	v283 = v274
	goto b30
b34:
	;
	v265 = or__x_252
	goto b36
b35:
	;
	v263 = case__11893 == vm.Keyword("eq")
	v265 = v263
	goto b36
b36:
	;
	v274 = v265
	goto b33
b37:
	;
	v378, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit!").Deref(), []vm.Value{l, nid, op})
	if callErr != nil {
		return nil, callErr
	}
	v1017 = v378
	goto b39
b38:
	;
	v393 = case__11893 == vm.Keyword("call")
	if v393 {
		goto b43
	} else {
		goto b44
	}
b39:
	;
	v1025 = v1017
	goto b15
b40:
	;
	v368 = or__x_355
	goto b42
b41:
	;
	v366 = case__11893 == vm.Keyword("dec")
	v368 = v366
	goto b42
b42:
	;
	if v368 {
		goto b37
	} else {
		goto b38
	}
b43:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__12055, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("call"), arg__12055})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__12067, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__12068 = rt.SubValue(vm.Int(0), arg__12067)
	v412, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, arg__12068})
	if callErr != nil {
		return nil, callErr
	}
	v1009 = v412
	goto b45
b44:
	;
	v427 = case__11893 == vm.Keyword("set-var")
	if v427 {
		goto b46
	} else {
		goto b47
	}
b45:
	;
	v1017 = v1009
	goto b39
b46:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit!").Deref(), []vm.Value{l, nid, vm.Keyword("set-var")})
	if callErr != nil {
		return nil, callErr
	}
	v436, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	v1001 = v436
	goto b48
b47:
	;
	v451 = case__11893 == vm.Keyword("tail-call")
	if v451 {
		goto b49
	} else {
		goto b50
	}
b48:
	;
	v1009 = v1001
	goto b45
b49:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__12099, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v460, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("tail-call"), arg__12099})
	if callErr != nil {
		return nil, callErr
	}
	v993 = v460
	goto b51
b50:
	;
	v475 = case__11893 == vm.Keyword("return")
	if v475 {
		goto b52
	} else {
		goto b53
	}
b51:
	;
	v1001 = v993
	goto b48
b52:
	;
	v480, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit!").Deref(), []vm.Value{l, nid, vm.Keyword("return")})
	if callErr != nil {
		return nil, callErr
	}
	v985 = v480
	goto b54
b53:
	;
	v495 = case__11893 == vm.Keyword("branch")
	if v495 {
		goto b55
	} else {
		goto b56
	}
b54:
	;
	v993 = v985
	goto b51
b55:
	;
	args, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-args").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	target, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	argc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{target, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__12132, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{target, f})
	if callErr != nil {
		return nil, callErr
	}
	target_params, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__12132})
	if callErr != nil {
		return nil, callErr
	}
	cur_sp, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	drop_count = rt.SubValue(cur_sp, target_params)
	arg__12143, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("current-block"), []vm.Value{arg__12143})
	if callErr != nil {
		return nil, callErr
	}
	arg__12151, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12152, callErr = rt.InvokeValue(vm.Keyword("current-block"), []vm.Value{arg__12151})
	if callErr != nil {
		return nil, callErr
	}
	cur_junk, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "block-junk-of").Deref(), []vm.Value{l, arg__12152})
	if callErr != nil {
		return nil, callErr
	}
	v551, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "pos?").Deref(), []vm.Value{drop_count})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v551) {
		goto b58
	} else {
		goto b59
	}
b56:
	;
	v689 = case__11893 == vm.Keyword("branch-if")
	if v689 {
		goto b64
	} else {
		goto b65
	}
b57:
	;
	v985 = v977
	goto b54
b58:
	;
	ignore = rt.SubValue(drop_count, argc)
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12169, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	off_ip, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-emit-recur").Deref(), []vm.Value{arg__12169, cur_sp, argc, ignore})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "record-source-info!").Deref(), []vm.Value{l, nid})
	if callErr != nil {
		return nil, callErr
	}
	arg__12182 = rt.SubValue(off_ip, vm.Int(1))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("src-ip"), arg__12182, vm.Keyword("offset-slot"), vm.Int(1), vm.Keyword("negate?"), vm.Boolean(true), vm.Keyword("target-block"), target})
	if callErr != nil {
		return nil, callErr
	}
	arg__12202, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("src-ip"), arg__12182, vm.Keyword("offset-slot"), vm.Int(1), vm.Keyword("negate?"), vm.Boolean(true), vm.Keyword("target-block"), target})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "add-patch!").Deref(), []vm.Value{l, arg__12202})
	if callErr != nil {
		return nil, callErr
	}
	v583, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "set-stack-sp!").Deref(), []vm.Value{l, target_params})
	if callErr != nil {
		return nil, callErr
	}
	v660 = v583
	goto b60
b59:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b61
	} else {
		goto b62
	}
b60:
	;
	v977 = v660
	goto b57
b61:
	;
	arg_ip, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-placeholder!").Deref(), []vm.Value{l, nid, vm.Keyword("branch")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "record-block-junk!").Deref(), []vm.Value{l, target, cur_junk})
	if callErr != nil {
		return nil, callErr
	}
	arg__12226 = rt.SubValue(arg_ip, vm.Int(1))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("src-ip"), arg__12226, vm.Keyword("offset-slot"), vm.Int(1), vm.Keyword("negate?"), vm.Boolean(false), vm.Keyword("target-block"), target})
	if callErr != nil {
		return nil, callErr
	}
	arg__12246, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("src-ip"), arg__12226, vm.Keyword("offset-slot"), vm.Int(1), vm.Keyword("negate?"), vm.Boolean(false), vm.Keyword("target-block"), target})
	if callErr != nil {
		return nil, callErr
	}
	v640, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "add-patch!").Deref(), []vm.Value{l, arg__12246})
	if callErr != nil {
		return nil, callErr
	}
	v644 = v640
	goto b63
b62:
	;
	v644 = vm.NIL
	goto b63
b63:
	;
	v660 = v644
	goto b60
b64:
	;
	ft, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-false").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	tt, callErr = rt.InvokeValue(rt.LookupVar("ir", "cond-target-true").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	ft_target, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{ft})
	if callErr != nil {
		return nil, callErr
	}
	tt_target, callErr = rt.InvokeValue(rt.LookupVar("ir", "branch-target-target").Deref(), []vm.Value{tt})
	if callErr != nil {
		return nil, callErr
	}
	arg_ip, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-placeholder!").Deref(), []vm.Value{l, nid, vm.Keyword("branch-if")})
	if callErr != nil {
		return nil, callErr
	}
	arg__12273, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("current-block"), []vm.Value{arg__12273})
	if callErr != nil {
		return nil, callErr
	}
	arg__12281, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12282, callErr = rt.InvokeValue(vm.Keyword("current-block"), []vm.Value{arg__12281})
	if callErr != nil {
		return nil, callErr
	}
	cur_junk, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "block-junk-of").Deref(), []vm.Value{l, arg__12282})
	if callErr != nil {
		return nil, callErr
	}
	arg__12286, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12288 = rt.AddValue(arg__12286, cur_junk)
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{ft_target, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__12302, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-params").Deref(), []vm.Value{ft_target, f})
	if callErr != nil {
		return nil, callErr
	}
	arg__12303, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__12302})
	if callErr != nil {
		return nil, callErr
	}
	v723 = rt.SubValue(arg__12288, vm.Int(1))
	target_junk = rt.SubValue(v723, arg__12303)
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "record-block-junk!").Deref(), []vm.Value{l, ft_target, target_junk})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "record-block-junk!").Deref(), []vm.Value{l, tt_target, target_junk})
	if callErr != nil {
		return nil, callErr
	}
	arg__12322 = rt.SubValue(arg_ip, vm.Int(1))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("src-ip"), arg__12322, vm.Keyword("offset-slot"), vm.Int(1), vm.Keyword("negate?"), vm.Boolean(false), vm.Keyword("target-block"), ft_target})
	if callErr != nil {
		return nil, callErr
	}
	arg__12342, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("src-ip"), arg__12322, vm.Keyword("offset-slot"), vm.Int(1), vm.Keyword("negate?"), vm.Boolean(false), vm.Keyword("target-block"), ft_target})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "add-patch!").Deref(), []vm.Value{l, arg__12342})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	my_block, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{nid, f})
	if callErr != nil {
		return nil, callErr
	}
	next_block_id = rt.AddValue(my_block, vm.Int(1))
	v789, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{tt_target, next_block_id})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v789) {
		goto b67
	} else {
		goto b68
	}
b65:
	;
	v849 = case__11893 == vm.Keyword("load-closed")
	if v849 {
		goto b70
	} else {
		goto b71
	}
b66:
	;
	v977 = v969
	goto b57
b67:
	;
	arg_ip2, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-placeholder!").Deref(), []vm.Value{l, nid, vm.Keyword("branch")})
	if callErr != nil {
		return nil, callErr
	}
	arg__12370 = rt.SubValue(arg_ip2, vm.Int(1))
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("src-ip"), arg__12370, vm.Keyword("offset-slot"), vm.Int(1), vm.Keyword("negate?"), vm.Boolean(false), vm.Keyword("target-block"), tt_target})
	if callErr != nil {
		return nil, callErr
	}
	arg__12390, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("src-ip"), arg__12370, vm.Keyword("offset-slot"), vm.Int(1), vm.Keyword("negate?"), vm.Boolean(false), vm.Keyword("target-block"), tt_target})
	if callErr != nil {
		return nil, callErr
	}
	v814, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "add-patch!").Deref(), []vm.Value{l, arg__12390})
	if callErr != nil {
		return nil, callErr
	}
	v818 = v814
	goto b69
b68:
	;
	v818 = vm.NIL
	goto b69
b69:
	;
	v969 = v818
	goto b66
b70:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__12407, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("load-closed"), arg__12407})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v864, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v961 = v864
	goto b72
b71:
	;
	v879 = case__11893 == vm.Keyword("make-closure")
	if v879 {
		goto b73
	} else {
		goto b74
	}
b72:
	;
	v969 = v961
	goto b66
b73:
	;
	v884, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit!").Deref(), []vm.Value{l, nid, vm.Keyword("make-closure")})
	if callErr != nil {
		return nil, callErr
	}
	v953 = v884
	goto b75
b74:
	;
	v899 = case__11893 == vm.Keyword("push-closed")
	if v899 {
		goto b76
	} else {
		goto b77
	}
b75:
	;
	v961 = v953
	goto b72
b76:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit!").Deref(), []vm.Value{l, nid, vm.Keyword("push-closed")})
	if callErr != nil {
		return nil, callErr
	}
	v908, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(-1)})
	if callErr != nil {
		return nil, callErr
	}
	v945 = v908
	goto b78
b77:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b79
	} else {
		goto b80
	}
b78:
	;
	v953 = v945
	goto b75
b79:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/lower: unsupported op for lowering: "), op})
	if callErr != nil {
		return nil, callErr
	}
	arg__12451, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/lower: unsupported op for lowering: "), op})
	if callErr != nil {
		return nil, callErr
	}
	v933, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__12451})
	if callErr != nil {
		return nil, callErr
	}
	v937 = v933
	goto b81
b80:
	;
	v937 = vm.NIL
	goto b81
b81:
	;
	v945 = v937
	goto b78
}
func materialize_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var f vm.Value
	var pos vm.Value
	var l vm.Value
	var nid vm.Value
	var arg__12465 vm.Value
	var arg__12466 vm.Value
	var and__x_27 bool
	var op vm.Value
	var aux vm.Value
	var v100 bool
	var v303 vm.Value
	var arg__12477 vm.Value
	var arg__12478 vm.Value
	var nth_arg vm.Value
	var arg__12493 vm.Value
	var arg__12497 vm.Value
	var v73 vm.Value
	var v75 vm.Value
	var arg__12473 vm.Value
	var v42 bool
	var and__x_37 bool
	var v45 bool
	var arg__12531 vm.Value
	var arg__12535 vm.Value
	var idx vm.Value
	var v121 vm.Value
	var case__12452 vm.Value
	var v138 bool
	var v294 vm.Value
	var arg__12569 vm.Value
	var v153 vm.Value
	var v170 bool
	var v285 vm.Value
	var arg__12589 vm.Value
	var v187 vm.Value
	var v204 bool
	var v276 vm.Value
	var arg__12624 vm.Value
	var v219 vm.Value
	var v267 vm.Value
	var arg__12657 vm.Value
	var v254 vm.Value
	var v258 vm.Value
	var callErr error
	f, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	pos, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "value-pos-of").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(pos) {
		l = arg0
		nid = arg1
		goto b1
	} else {
		l = arg0
		nid = arg1
		goto b2
	}
b1:
	;
	arg__12465, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12466 = rt.SubValue(arg__12465, vm.Int(1))
	and__x_27 = pos == arg__12466
	if and__x_27 {
		goto b7
	} else {
		and__x_37 = and__x_27
		goto b8
	}
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
	v100 = op == vm.Keyword("const")
	if v100 {
		goto b10
	} else {
		case__12452 = op
		goto b11
	}
b3:
	;
	return v303, nil
b4:
	;
	v75 = vm.NIL
	goto b6
b5:
	;
	arg__12477, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12478 = rt.SubValue(arg__12477, vm.Int(1))
	nth_arg = rt.SubValue(arg__12478, pos)
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12493, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12497, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "stack-sp").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-emit-dup-nth").Deref(), []vm.Value{arg__12493, arg__12497, nth_arg})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v73, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v75 = v73
	goto b6
b6:
	;
	v303 = v75
	goto b3
b7:
	;
	arg__12473, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "use-count-of").Deref(), []vm.Value{l, nid})
	if callErr != nil {
		return nil, callErr
	}
	v42 = arg__12473 == vm.Int(1)
	v45 = v42
	goto b9
b8:
	;
	v45 = and__x_37
	goto b9
b9:
	;
	if v45 {
		goto b4
	} else {
		goto b5
	}
b10:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "template-value").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__12531, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12535, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "template-value").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	idx, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-intern-const").Deref(), []vm.Value{arg__12531, arg__12535})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("const"), idx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v121, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v294 = v121
	goto b12
b11:
	;
	v138 = case__12452 == vm.Keyword("load-arg")
	if v138 {
		goto b13
	} else {
		goto b14
	}
b12:
	;
	v303 = v294
	goto b3
b13:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__12569, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("load-arg"), arg__12569})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v153, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v285 = v153
	goto b15
b14:
	;
	v170 = case__12452 == vm.Keyword("load-var")
	if v170 {
		goto b16
	} else {
		goto b17
	}
b15:
	;
	v294 = v285
	goto b12
b16:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12589, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	idx, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-intern-const").Deref(), []vm.Value{arg__12589, aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("load-var"), idx})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v187, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v276 = v187
	goto b18
b17:
	;
	v204 = case__12452 == vm.Keyword("load-closed")
	if v204 {
		goto b19
	} else {
		goto b20
	}
b18:
	;
	v285 = v276
	goto b15
b19:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__12624, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "int").Deref(), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "emit-with-arg!").Deref(), []vm.Value{l, nid, vm.Keyword("load-closed"), arg__12624})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-stack-sp!").Deref(), []vm.Value{l, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	v219, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "bump-max-stack!").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	v267 = v219
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
	v276 = v267
	goto b18
b22:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/lower: value %"), nid, vm.String(" not on stack for materialize (op="), op, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	arg__12657, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "str").Deref(), []vm.Value{vm.String("ir/lower: value %"), nid, vm.String(" not on stack for materialize (op="), op, vm.String(")")})
	if callErr != nil {
		return nil, callErr
	}
	v254, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "throw").Deref(), []vm.Value{arg__12657})
	if callErr != nil {
		return nil, callErr
	}
	v258 = v254
	goto b24
b23:
	;
	v258 = vm.NIL
	goto b24
b24:
	;
	v267 = v258
	goto b21
}
func materialize_branch_args_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var argc vm.Value
	var k int
	var args vm.Value
	var l vm.Value
	var v19 bool
	var arg__12679 vm.Value
	var pos vm.Value
	var skip int
	var v87 bool
	var v62 int
	var v65 int
	var v51 bool
	var and__x vm.Value
	var v54 vm.Value
	var v90 vm.Value
	var v103 vm.Value
	var v226 vm.Value
	var v106 vm.Value
	var v219 vm.Value
	var arg__12709 vm.Value
	var arg__12722 vm.Value
	var arg__12724 vm.Value
	var doseq_seq__12658 vm.Value
	var doseq_loop__12659 vm.Value
	var a vm.Value
	var v159 vm.Value
	var v190 bool
	var arg__12757 vm.Value
	var callErr error
	argc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	k = 0
	args = arg1
	l = arg0
	goto b1
b1:
	;
	v19 = rt.GeValue(vm.Int(k), argc)
	if v19 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	skip = k
	goto b4
b3:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(k)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12679, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(k)})
	if callErr != nil {
		return nil, callErr
	}
	pos, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "value-pos-of").Deref(), []vm.Value{l, arg__12679})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(pos) {
		goto b8
	} else {
		and__x = pos
		goto b9
	}
b4:
	;
	v87 = vm.Int(skip) == argc
	if v87 {
		goto b11
	} else {
		goto b12
	}
b5:
	;
	v62 = k + 1
	k = v62
	goto b1
b6:
	;
	v65 = k
	goto b7
b7:
	;
	skip = v65
	goto b4
b8:
	;
	v51 = pos == vm.Int(k)
	v54 = vm.Boolean(v51)
	goto b10
b9:
	;
	v54 = and__x
	goto b10
b10:
	;
	if vm.IsTruthy(v54) {
		goto b5
	} else {
		goto b6
	}
b11:
	;
	v90, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "consume-refs-in-place!").Deref(), []vm.Value{l, args})
	if callErr != nil {
		return nil, callErr
	}
	v226 = v90
	goto b13
b12:
	;
	v103, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "args-at-top?").Deref(), []vm.Value{l, args})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v103) {
		goto b14
	} else {
		goto b15
	}
b13:
	;
	return v226, nil
b14:
	;
	v106, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "consume-refs-in-place!").Deref(), []vm.Value{l, args})
	if callErr != nil {
		return nil, callErr
	}
	v219 = v106
	goto b16
b15:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b17
	} else {
		goto b18
	}
b16:
	;
	v226 = v219
	goto b13
b17:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	arg__12709, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subvec").Deref(), []vm.Value{arg__12709, vm.Int(skip)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	arg__12722, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{args})
	if callErr != nil {
		return nil, callErr
	}
	arg__12724, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "subvec").Deref(), []vm.Value{arg__12722, vm.Int(skip)})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__12658, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__12724})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__12659 = doseq_seq__12658
	goto b20
b18:
	;
	goto b19
b19:
	;
	v219 = vm.NIL
	goto b16
b20:
	;
	if vm.IsTruthy(doseq_loop__12659) {
		goto b21
	} else {
		goto b22
	}
b21:
	;
	a, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__12659})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "materialize!").Deref(), []vm.Value{l, a})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "decrement-use!").Deref(), []vm.Value{l, a})
	if callErr != nil {
		return nil, callErr
	}
	v159, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__12659})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__12659 = v159
	goto b20
b22:
	;
	goto b23
b23:
	;
	goto b24
b24:
	;
	v190 = k < skip
	if v190 {
		goto b25
	} else {
		goto b26
	}
b25:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(k)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12757, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{args, vm.Int(k)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "decrement-use!").Deref(), []vm.Value{l, arg__12757})
	if callErr != nil {
		return nil, callErr
	}
	goto b24
b26:
	;
	goto b27
b27:
	;
	goto b19
}
func materialize_refs_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var doseq_seq__12759 vm.Value
	var doseq_loop__12760 vm.Value
	var l vm.Value
	var r vm.Value
	var v24 vm.Value
	var callErr error
	doseq_seq__12759, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__12760 = doseq_seq__12759
	l = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__12760) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	r, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__12760})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "materialize!").Deref(), []vm.Value{l, r})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "decrement-use!").Deref(), []vm.Value{l, r})
	if callErr != nil {
		return nil, callErr
	}
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__12760})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__12760 = v24
	goto b1
b3:
	;
	goto b4
b4:
	;
	return vm.NIL, nil
}
func new_lowerer(arg0 vm.Value) (vm.Value, error) {
	var uses vm.Value
	var arg__12791 vm.Value
	var n_blocks vm.Value
	var i int
	var acc vm.Value
	var arg__12796 vm.Value
	var v28 bool
	var f vm.Value
	var us vm.Value
	var v33 int
	var v47 vm.Value
	var use_count vm.Value
	var arg__12831 vm.Value
	var arg__12832 vm.Value
	var arg__12850 vm.Value
	var arg__12851 vm.Value
	var arg__12865 vm.Value
	var arg__12866 vm.Value
	var arg__12889 vm.Value
	var arg__12890 vm.Value
	var arg__12908 vm.Value
	var arg__12909 vm.Value
	var arg__12923 vm.Value
	var arg__12924 vm.Value
	var arg__12933 vm.Value
	var v156 vm.Value
	var arg__12818 vm.Value
	var v55 vm.Value
	var v57 vm.Value
	var callErr error
	uses, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__12791, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	n_blocks, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__12791})
	if callErr != nil {
		return nil, callErr
	}
	i = 0
	acc = vm.EmptyPersistentMap
	goto b1
b1:
	;
	arg__12796, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{uses})
	if callErr != nil {
		return nil, callErr
	}
	v28 = rt.GeValue(vm.Int(i), arg__12796)
	if v28 {
		f = arg0
		goto b2
	} else {
		goto b3
	}
b2:
	;
	use_count = acc
	goto b4
b3:
	;
	us, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v33 = i + 1
	v47, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v47) {
		goto b5
	} else {
		goto b6
	}
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-consts").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__12831, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-consts").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__12832, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-chunk").Deref(), []vm.Value{arg__12831})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n_blocks, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12850, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n_blocks, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12851, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__12850})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n_blocks, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12865, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n_blocks, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12866, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__12865})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("value-stack-pos"), vm.EmptyPersistentMap, vm.Keyword("chunk"), arg__12832, vm.Keyword("uses"), uses, vm.Keyword("f"), f, vm.Keyword("block-junk"), arg__12851, vm.Keyword("block-ips"), arg__12866, vm.Keyword("patches"), vm.NewArrayVector([]vm.Value{}), vm.Keyword("use-count"), use_count, vm.Keyword("current-block"), vm.Int(0), vm.Keyword("stack-sp"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-consts").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__12889, callErr = rt.InvokeValue(rt.LookupVar("ir", "fn-consts").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__12890, callErr = rt.InvokeValue(rt.LookupVar("ir", "new-chunk").Deref(), []vm.Value{arg__12889})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n_blocks, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12908, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n_blocks, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12909, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__12908})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n_blocks, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12923, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "repeat").Deref(), []vm.Value{n_blocks, vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	arg__12924, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__12923})
	if callErr != nil {
		return nil, callErr
	}
	arg__12933, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("value-stack-pos"), vm.EmptyPersistentMap, vm.Keyword("chunk"), arg__12890, vm.Keyword("uses"), uses, vm.Keyword("f"), f, vm.Keyword("block-junk"), arg__12909, vm.Keyword("block-ips"), arg__12924, vm.Keyword("patches"), vm.NewArrayVector([]vm.Value{}), vm.Keyword("use-count"), use_count, vm.Keyword("current-block"), vm.Int(0), vm.Keyword("stack-sp"), vm.Int(0)})
	if callErr != nil {
		return nil, callErr
	}
	v156, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{arg__12933})
	if callErr != nil {
		return nil, callErr
	}
	return v156, nil
b5:
	;
	v57 = acc
	goto b7
b6:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-count").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	arg__12818, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-count").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	v55, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{acc, vm.Int(i), arg__12818})
	if callErr != nil {
		return nil, callErr
	}
	v57 = v55
	goto b7
b7:
	;
	i = v33
	acc = v57
	goto b1
}
func patch_branches_BANG_(arg0 vm.Value) (vm.Value, error) {
	var arg__12940 vm.Value
	var arg__12947 vm.Value
	var arg__12948 vm.Value
	var doseq_seq__12934 vm.Value
	var doseq_loop__12935 vm.Value
	var l vm.Value
	var p vm.Value
	var arg__12960 vm.Value
	var target_ip vm.Value
	var src_ip vm.Value
	var v44 vm.Value
	var v46 vm.Value
	var v48 vm.Value
	var offset vm.Value
	var arg__12983 vm.Value
	var arg__12987 vm.Value
	var arg__12988 vm.Value
	var v70 vm.Value
	var callErr error
	arg__12940, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("patches"), []vm.Value{arg__12940})
	if callErr != nil {
		return nil, callErr
	}
	arg__12947, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__12948, callErr = rt.InvokeValue(vm.Keyword("patches"), []vm.Value{arg__12947})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__12934, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__12948})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__12935 = doseq_seq__12934
	l = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__12935) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	p, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__12935})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("target-block"), []vm.Value{p})
	if callErr != nil {
		return nil, callErr
	}
	arg__12960, callErr = rt.InvokeValue(vm.Keyword("target-block"), []vm.Value{p})
	if callErr != nil {
		return nil, callErr
	}
	target_ip, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "block-ip").Deref(), []vm.Value{l, arg__12960})
	if callErr != nil {
		return nil, callErr
	}
	src_ip, callErr = rt.InvokeValue(vm.Keyword("src-ip"), []vm.Value{p})
	if callErr != nil {
		return nil, callErr
	}
	v44, callErr = rt.InvokeValue(vm.Keyword("negate?"), []vm.Value{p})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v44) {
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
	v46 = rt.SubValue(src_ip, target_ip)
	offset = v46
	goto b7
b6:
	;
	v48 = rt.SubValue(target_ip, src_ip)
	offset = v48
	goto b7
b7:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("offset-slot"), []vm.Value{p})
	if callErr != nil {
		return nil, callErr
	}
	arg__12983, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__12987, callErr = rt.InvokeValue(vm.Keyword("offset-slot"), []vm.Value{p})
	if callErr != nil {
		return nil, callErr
	}
	arg__12988 = rt.AddValue(src_ip, arg__12987)
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-update!").Deref(), []vm.Value{arg__12983, arg__12988, offset})
	if callErr != nil {
		return nil, callErr
	}
	v70, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__12935})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__12935 = v70
	goto b1
}
func record_block_junk_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 int) (vm.Value, error) {
	var cur vm.Value
	var v14 bool
	var l vm.Value
	var bid vm.Value
	var junk int
	var v23 vm.Value
	var v27 vm.Value
	var callErr error
	cur, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "block-junk-of").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v14 = rt.GtValue(vm.Int(arg2), cur)
	if v14 {
		l = arg0
		bid = arg1
		junk = arg2
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v23, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{l, rt.LookupVar("clojure.core", "update").Deref(), vm.Keyword("block-junk"), rt.LookupVar("clojure.core", "assoc").Deref(), bid, vm.Int(junk)})
	if callErr != nil {
		return nil, callErr
	}
	v27 = v23
	goto b3
b2:
	;
	v27 = vm.NIL
	goto b3
b3:
	;
	return v27, nil
}
func record_source_info_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var l vm.Value
	var inst_id vm.Value
	var c vm.Value
	var arg__13028 vm.Value
	var arg__13041 vm.Value
	var arg__13042 vm.Value
	var doseq_seq__13013 vm.Value
	var doseq_loop__13014 vm.Value
	var si vm.Value
	var v42 vm.Value
	var callErr error
	if vm.IsTruthy(arg1) {
		l = arg0
		inst_id = arg1
		goto b1
	} else {
		goto b2
	}
b1:
	;
	c, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "chunk-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__13028, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "source-infos").Deref(), []vm.Value{inst_id, arg__13028})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__13041, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__13042, callErr = rt.InvokeValue(rt.LookupVar("ir", "source-infos").Deref(), []vm.Value{inst_id, arg__13041})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__13013, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__13042})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__13014 = doseq_seq__13013
	goto b4
b2:
	;
	goto b3
b3:
	;
	return vm.NIL, nil
b4:
	;
	if vm.IsTruthy(doseq_loop__13014) {
		goto b5
	} else {
		goto b6
	}
b5:
	;
	si, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__13014})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "chunk-add-source-info!").Deref(), []vm.Value{c, si})
	if callErr != nil {
		return nil, callErr
	}
	v42, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__13014})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__13014 = v42
	goto b4
b6:
	;
	goto b7
b7:
	;
	goto b3
}
func refs_at_top_last_use_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__13062 vm.Value
	var v12 vm.Value
	var l vm.Value
	var refs vm.Value
	var arg__13075 vm.Value
	var v25 vm.Value
	var v134 vm.Value
	var v130 vm.Value
	var v126 vm.Value
	var i int
	var arg__13080 vm.Value
	var v48 bool
	var arg__13096 vm.Value
	var arg__13114 vm.Value
	var arg__13115 vm.Value
	var v73 vm.Value
	var v119 vm.Value
	var arg__13132 vm.Value
	var v88 vm.Value
	var v114 vm.Value
	var v109 vm.Value
	var v100 int
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__13062, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{arg__13062})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		goto b1
	} else {
		l = arg0
		refs = arg1
		goto b2
	}
b1:
	;
	v134 = vm.Boolean(true)
	goto b3
b2:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "args-at-top?").Deref(), []vm.Value{l, refs})
	if callErr != nil {
		return nil, callErr
	}
	arg__13075, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "args-at-top?").Deref(), []vm.Value{l, refs})
	if callErr != nil {
		return nil, callErr
	}
	v25, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__13075})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v25) {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v134, nil
b4:
	;
	v130 = vm.Boolean(false)
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
	v134 = v130
	goto b3
b7:
	;
	i = 0
	goto b10
b8:
	;
	v126 = vm.NIL
	goto b9
b9:
	;
	v130 = v126
	goto b6
b10:
	;
	arg__13080, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{refs})
	if callErr != nil {
		return nil, callErr
	}
	v48 = rt.GeValue(vm.Int(i), arg__13080)
	if v48 {
		goto b11
	} else {
		goto b12
	}
b11:
	;
	v119 = vm.Boolean(true)
	goto b13
b12:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__13096, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "use-count-of").Deref(), []vm.Value{l, arg__13096})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__13114, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__13115, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "use-count-of").Deref(), []vm.Value{l, arg__13114})
	if callErr != nil {
		return nil, callErr
	}
	v73, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Int(1), arg__13115})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v73) {
		goto b14
	} else {
		goto b15
	}
b13:
	;
	v126 = v119
	goto b9
b14:
	;
	v114 = vm.Boolean(false)
	goto b16
b15:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	arg__13132, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{refs, vm.Int(i)})
	if callErr != nil {
		return nil, callErr
	}
	v88, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "contains-after?").Deref(), []vm.Value{refs, vm.Int(i), arg__13132})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v88) {
		goto b17
	} else {
		goto b18
	}
b16:
	;
	v119 = v114
	goto b13
b17:
	;
	v109 = vm.Boolean(false)
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
	v114 = v109
	goto b16
b20:
	;
	v100 = i + 1
	i = v100
	goto b10
b21:
	;
	goto b22
b22:
	;
	v109 = vm.NIL
	goto b19
}
func set_block_ip_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var v16 vm.Value
	var callErr error
	v16, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "update").Deref(), vm.Keyword("block-ips"), rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v4 vm.Value
		var callErr error
		v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "assoc").Deref(), []vm.Value{arg0, arg1, arg2})
		if callErr != nil {
			return nil, callErr
		}
		return v4, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	return v16, nil
}
func set_stack_sp_BANG_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v7 vm.Value
	var callErr error
	v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc").Deref(), vm.Keyword("stack-sp"), arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v7, nil
}
func set_value_pos_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var v10 vm.Value
	var callErr error
	v10, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "update").Deref(), vm.Keyword("value-stack-pos"), rt.LookupVar("clojure.core", "assoc").Deref(), arg1, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v10, nil
}
func should_body_emit_cheap_QMARK_(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var uc vm.Value
	var v12 vm.Value
	var l vm.Value
	var nid vm.Value
	var v23 bool
	var v211 vm.Value
	var v206 vm.Value
	var arg__13193 vm.Value
	var uses vm.Value
	var arg__13198 vm.Value
	var v49 bool
	var v201 vm.Value
	var v52 vm.Value
	var us vm.Value
	var or__x vm.Value
	var v192 vm.Value
	var v88 vm.Value
	var v90 vm.Value
	var user_id vm.Value
	var arg__13223 vm.Value
	var user_refs vm.Value
	var arg__13232 vm.Value
	var v139 vm.Value
	var v185 vm.Value
	var v174 vm.Value
	var arg__13236 vm.Value
	var v161 bool
	var v165 vm.Value
	var callErr error
	uc, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "use-count-of").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	v12, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{uc})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v12) {
		goto b1
	} else {
		l = arg0
		nid = arg1
		goto b2
	}
b1:
	;
	v211 = vm.Boolean(false)
	goto b3
b2:
	;
	v23 = rt.GtValue(uc, vm.Int(1))
	if v23 {
		goto b4
	} else {
		goto b5
	}
b3:
	;
	return v211, nil
b4:
	;
	v206 = vm.Boolean(false)
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
	v211 = v206
	goto b3
b7:
	;
	arg__13193, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	uses, callErr = rt.InvokeValue(vm.Keyword("uses"), []vm.Value{arg__13193})
	if callErr != nil {
		return nil, callErr
	}
	arg__13198, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{uses})
	if callErr != nil {
		return nil, callErr
	}
	v49 = rt.LtValue(nid, arg__13198)
	if v49 {
		goto b10
	} else {
		goto b11
	}
b8:
	;
	v201 = vm.NIL
	goto b9
b9:
	;
	v206 = v201
	goto b6
b10:
	;
	v52, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{uses, nid})
	if callErr != nil {
		return nil, callErr
	}
	us = v52
	goto b12
b11:
	;
	us = vm.NIL
	goto b12
b12:
	;
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nil?").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b16
	} else {
		goto b17
	}
b13:
	;
	v192 = vm.Boolean(false)
	goto b15
b14:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b19
	} else {
		goto b20
	}
b15:
	;
	v201 = v192
	goto b9
b16:
	;
	v90 = or__x
	goto b18
b17:
	;
	v88, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-empty?").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	v90 = v88
	goto b18
b18:
	;
	if vm.IsTruthy(v90) {
		goto b13
	} else {
		goto b14
	}
b19:
	;
	user_id, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-first").Deref(), []vm.Value{us})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	arg__13223, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "f-of").Deref(), []vm.Value{l})
	if callErr != nil {
		return nil, callErr
	}
	user_refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{user_id, arg__13223})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{user_refs})
	if callErr != nil {
		return nil, callErr
	}
	arg__13232, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{user_refs})
	if callErr != nil {
		return nil, callErr
	}
	v139, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "zero?").Deref(), []vm.Value{arg__13232})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v139) {
		goto b22
	} else {
		goto b23
	}
b20:
	;
	v185 = vm.NIL
	goto b21
b21:
	;
	v192 = v185
	goto b15
b22:
	;
	v174 = vm.Boolean(true)
	goto b24
b23:
	;
	if vm.IsTruthy(vm.Keyword("else")) {
		goto b25
	} else {
		goto b26
	}
b24:
	;
	v185 = v174
	goto b21
b25:
	;
	arg__13236, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{user_refs})
	if callErr != nil {
		return nil, callErr
	}
	v161 = arg__13236 == nid
	v165 = vm.Boolean(v161)
	goto b27
b26:
	;
	v165 = vm.NIL
	goto b27
b27:
	;
	v174 = v165
	goto b24
}
func stack_sp(arg0 vm.Value) (vm.Value, error) {
	var arg__13242 vm.Value
	var v4 vm.Value
	var callErr error
	arg__13242, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("stack-sp"), []vm.Value{arg__13242})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func template_value(arg0 vm.Value) (vm.Value, error) {
	var and__x vm.Value
	var aux vm.Value
	var inner vm.Value
	var arity vm.Value
	var variadic_QMARK_ vm.Value
	var chunk vm.Value
	var v29 vm.Value
	var v82 vm.Value
	var arg__13248 vm.Value
	var v13 bool
	var v16 vm.Value
	var arg__13281 vm.Value
	var fn_vals vm.Value
	var arg__13294 vm.Value
	var v67 vm.Value
	var v79 vm.Value
	var arg__13271 vm.Value
	var v43 bool
	var v46 vm.Value
	var v76 vm.Value
	var callErr error
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(and__x) {
		aux = arg0
		goto b4
	} else {
		aux = arg0
		goto b5
	}
b1:
	;
	inner, callErr = rt.InvokeValue(vm.Keyword("fn"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arity, callErr = rt.InvokeValue(vm.Keyword("arity"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	variadic_QMARK_, callErr = rt.InvokeValue(vm.Keyword("variadic?"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	chunk, callErr = rt.InvokeValue(rt.LookupVar("ir.lower", "lower").Deref(), []vm.Value{inner})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "chunk->fn").Deref(), []vm.Value{arity, variadic_QMARK_, chunk})
	if callErr != nil {
		return nil, callErr
	}
	v82 = v29
	goto b3
b2:
	;
	and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "map?").Deref(), []vm.Value{aux})
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
	return v82, nil
b4:
	;
	arg__13248, callErr = rt.InvokeValue(vm.Keyword("kind"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v13 = arg__13248 == vm.Keyword("fn-template")
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
	_, callErr = rt.InvokeValue(vm.Keyword("fns"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	arg__13281, callErr = rt.InvokeValue(vm.Keyword("fns"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	fn_vals, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.LookupVar("ir.lower", "template-value").Deref(), arg__13281})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), fn_vals})
	if callErr != nil {
		return nil, callErr
	}
	arg__13294, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "apply").Deref(), []vm.Value{rt.LookupVar("clojure.core", "list").Deref(), fn_vals})
	if callErr != nil {
		return nil, callErr
	}
	v67, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "make-multi-arity").Deref(), []vm.Value{arg__13294})
	if callErr != nil {
		return nil, callErr
	}
	v79 = v67
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
	v82 = v79
	goto b3
b10:
	;
	arg__13271, callErr = rt.InvokeValue(vm.Keyword("kind"), []vm.Value{aux})
	if callErr != nil {
		return nil, callErr
	}
	v43 = arg__13271 == vm.Keyword("multi-fn-template")
	v46 = vm.Boolean(v43)
	goto b12
b11:
	;
	v46 = and__x
	goto b12
b12:
	;
	if vm.IsTruthy(v46) {
		goto b7
	} else {
		goto b8
	}
b13:
	;
	v76 = aux
	goto b15
b14:
	;
	v76 = vm.NIL
	goto b15
b15:
	;
	v79 = v76
	goto b9
}
func use_count_of(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__13299 vm.Value
	var arg__13307 vm.Value
	var arg__13308 vm.Value
	var or__x vm.Value
	var v22 vm.Value
	var callErr error
	arg__13299, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("use-count"), []vm.Value{arg__13299})
	if callErr != nil {
		return nil, callErr
	}
	arg__13307, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__13308, callErr = rt.InvokeValue(vm.Keyword("use-count"), []vm.Value{arg__13307})
	if callErr != nil {
		return nil, callErr
	}
	or__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__13308, arg1})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(or__x) {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v22 = or__x
	goto b3
b2:
	;
	v22 = vm.Int(0)
	goto b3
b3:
	;
	return v22, nil
}
func value_pos_of(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__13314 vm.Value
	var arg__13322 vm.Value
	var arg__13323 vm.Value
	var v11 vm.Value
	var callErr error
	arg__13314, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("value-stack-pos"), []vm.Value{arg__13314})
	if callErr != nil {
		return nil, callErr
	}
	arg__13322, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__13323, callErr = rt.InvokeValue(vm.Keyword("value-stack-pos"), []vm.Value{arg__13322})
	if callErr != nil {
		return nil, callErr
	}
	v11, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "get").Deref(), []vm.Value{arg__13323, arg1})
	if callErr != nil {
		return nil, callErr
	}
	return v11, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.lower", map[string]vm.Value{"add-patch!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("add-patch!: wrong number of arguments %d (expected 2)", len(args))
		}
		return add_patch_BANG_(args[0], args[1])
	}), "args-at-top?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("args-at-top?: wrong number of arguments %d (expected 2)", len(args))
		}
		return args_at_top_QMARK_(args[0], args[1])
	}), "block-ip": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-ip: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_ip(args[0], args[1])
	}), "block-junk-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-junk-of: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_junk_of(args[0], args[1])
	}), "bump-max-stack!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("bump-max-stack!: wrong number of arguments %d (expected 1)", len(args))
		}
		return bump_max_stack_BANG_(args[0])
	}), "bump-stack-sp!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("bump-stack-sp!: wrong number of arguments %d (expected 2)", len(args))
		}
		return bump_stack_sp_BANG_(args[0], args[1])
	}), "check-cross-block!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("check-cross-block!: wrong number of arguments %d (expected 2)", len(args))
		}
		return check_cross_block_BANG_(args[0], args[1])
	}), "chunk-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("chunk-of: wrong number of arguments %d (expected 1)", len(args))
		}
		return chunk_of(args[0])
	}), "consume-refs-in-place!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("consume-refs-in-place!: wrong number of arguments %d (expected 2)", len(args))
		}
		return consume_refs_in_place_BANG_(args[0], args[1])
	}), "decrement-use!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("decrement-use!: wrong number of arguments %d (expected 2)", len(args))
		}
		return decrement_use_BANG_(args[0], args[1])
	}), "deferrable-branch-if-cond?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("deferrable-branch-if-cond?: wrong number of arguments %d (expected 3)", len(args))
		}
		return deferrable_branch_if_cond_QMARK_(args[0], args[1], args[2])
	}), "emit!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("emit!: wrong number of arguments %d (expected 3)", len(args))
		}
		return emit_BANG_(args[0], args[1], args[2])
	}), "emit-inst!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("emit-inst!: wrong number of arguments %d (expected 2)", len(args))
		}
		return emit_inst_BANG_(args[0], args[1])
	}), "emit-placeholder!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("emit-placeholder!: wrong number of arguments %d (expected 3)", len(args))
		}
		return emit_placeholder_BANG_(args[0], args[1], args[2])
	}), "emit-with-arg!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 4 {
			return nil, fmt.Errorf("emit-with-arg!: wrong number of arguments %d (expected 4)", len(args))
		}
		return emit_with_arg_BANG_(args[0], args[1], args[2], args[3])
	}), "f-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("f-of: wrong number of arguments %d (expected 1)", len(args))
		}
		return f_of(args[0])
	}), "is-terminator-branch-arg-use?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("is-terminator-branch-arg-use?: wrong number of arguments %d (expected 3)", len(args))
		}
		return is_terminator_branch_arg_use_QMARK_(args[0], args[1], args[2])
	}), "lower": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("lower: wrong number of arguments %d (expected 1)", len(args))
		}
		return lower(args[0])
	}), "lower-block!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("lower-block!: wrong number of arguments %d (expected 2)", len(args))
		}
		return lower_block_BANG_(args[0], args[1])
	}), "lower-node!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("lower-node!: wrong number of arguments %d (expected 2)", len(args))
		}
		return lower_node_BANG_(args[0], args[1])
	}), "materialize!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("materialize!: wrong number of arguments %d (expected 2)", len(args))
		}
		return materialize_BANG_(args[0], args[1])
	}), "materialize-branch-args!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("materialize-branch-args!: wrong number of arguments %d (expected 2)", len(args))
		}
		return materialize_branch_args_BANG_(args[0], args[1])
	}), "materialize-refs!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("materialize-refs!: wrong number of arguments %d (expected 2)", len(args))
		}
		return materialize_refs_BANG_(args[0], args[1])
	}), "new-lowerer": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("new-lowerer: wrong number of arguments %d (expected 1)", len(args))
		}
		return new_lowerer(args[0])
	}), "patch-branches!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("patch-branches!: wrong number of arguments %d (expected 1)", len(args))
		}
		return patch_branches_BANG_(args[0])
	}), "record-source-info!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("record-source-info!: wrong number of arguments %d (expected 2)", len(args))
		}
		return record_source_info_BANG_(args[0], args[1])
	}), "refs-at-top-last-use?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("refs-at-top-last-use?: wrong number of arguments %d (expected 2)", len(args))
		}
		return refs_at_top_last_use_QMARK_(args[0], args[1])
	}), "set-block-ip!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-block-ip!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_block_ip_BANG_(args[0], args[1], args[2])
	}), "set-stack-sp!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("set-stack-sp!: wrong number of arguments %d (expected 2)", len(args))
		}
		return set_stack_sp_BANG_(args[0], args[1])
	}), "set-value-pos!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-value-pos!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_value_pos_BANG_(args[0], args[1], args[2])
	}), "should-body-emit-cheap?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("should-body-emit-cheap?: wrong number of arguments %d (expected 2)", len(args))
		}
		return should_body_emit_cheap_QMARK_(args[0], args[1])
	}), "stack-sp": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("stack-sp: wrong number of arguments %d (expected 1)", len(args))
		}
		return stack_sp(args[0])
	}), "template-value": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("template-value: wrong number of arguments %d (expected 1)", len(args))
		}
		return template_value(args[0])
	}), "use-count-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("use-count-of: wrong number of arguments %d (expected 2)", len(args))
		}
		return use_count_of(args[0], args[1])
	}), "value-pos-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("value-pos-of: wrong number of arguments %d (expected 2)", len(args))
		}
		return value_pos_of(args[0], args[1])
	}),
	})
}
