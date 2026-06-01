package ir_data_generated

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func type_of(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8644 vm.Value
	var arg__8652 vm.Value
	var arg__8653 vm.Value
	var arg__8655 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8644, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8644})
	if callErr != nil {
		return nil, callErr
	}
	arg__8652, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8653, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8652})
	if callErr != nil {
		return nil, callErr
	}
	arg__8655, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8653, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("type"), []vm.Value{arg__8655})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func block_set_preds_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__8671 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("preds")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8671, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("preds")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__8671, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func set_block_of_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__8688 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("block")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8688, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("block")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__8688, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func block_set_params_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__8705 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("params")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8705, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("params")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__8705, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func set_refs_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__8722 vm.Value
	var v26 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("refs")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8722, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("refs")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__8722, arg2})
	if callErr != nil {
		return nil, callErr
	}
	v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc").Deref(), vm.Keyword("uses-dirty?"), vm.Boolean(true), vm.Keyword("uses-cache"), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	return v26, nil
}
func refs(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8742 vm.Value
	var arg__8750 vm.Value
	var arg__8751 vm.Value
	var arg__8753 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8742, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8742})
	if callErr != nil {
		return nil, callErr
	}
	arg__8750, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8751, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8750})
	if callErr != nil {
		return nil, callErr
	}
	arg__8753, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8751, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("refs"), []vm.Value{arg__8753})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func fn_uses_cache(arg0 vm.Value) (vm.Value, error) {
	var arg__8758 vm.Value
	var v4 vm.Value
	var callErr error
	arg__8758, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("uses-cache"), []vm.Value{arg__8758})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func block_insts(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8764 vm.Value
	var arg__8772 vm.Value
	var arg__8773 vm.Value
	var arg__8775 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8764, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__8764})
	if callErr != nil {
		return nil, callErr
	}
	arg__8772, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8773, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__8772})
	if callErr != nil {
		return nil, callErr
	}
	arg__8775, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8773, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8775})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func set_aux_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__8791 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("aux")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8791, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("aux")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__8791, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func fn_consts(arg0 vm.Value) (vm.Value, error) {
	var arg__8797 vm.Value
	var v4 vm.Value
	var callErr error
	arg__8797, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("consts"), []vm.Value{arg__8797})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func block_id(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8803 vm.Value
	var arg__8811 vm.Value
	var arg__8812 vm.Value
	var arg__8814 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8803, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__8803})
	if callErr != nil {
		return nil, callErr
	}
	arg__8811, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8812, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__8811})
	if callErr != nil {
		return nil, callErr
	}
	arg__8814, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8812, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("id"), []vm.Value{arg__8814})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func op(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8820 vm.Value
	var arg__8828 vm.Value
	var arg__8829 vm.Value
	var arg__8831 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8820, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8820})
	if callErr != nil {
		return nil, callErr
	}
	arg__8828, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8829, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8828})
	if callErr != nil {
		return nil, callErr
	}
	arg__8831, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8829, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("op"), []vm.Value{arg__8831})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func fn_name(arg0 vm.Value) (vm.Value, error) {
	var arg__8836 vm.Value
	var v4 vm.Value
	var callErr error
	arg__8836, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("name"), []vm.Value{arg__8836})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func set_type_of_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__8852 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("type")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8852, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("type")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__8852, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func fn_entry(arg0 vm.Value) (vm.Value, error) {
	var arg__8858 vm.Value
	var v4 vm.Value
	var callErr error
	arg__8858, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("entry"), []vm.Value{arg__8858})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func fn_arity(arg0 vm.Value) (vm.Value, error) {
	var arg__8863 vm.Value
	var v4 vm.Value
	var callErr error
	arg__8863, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("arity"), []vm.Value{arg__8863})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func branch_target_target(arg0 vm.Value) (vm.Value, error) {
	var v2 vm.Value
	var callErr error
	v2, callErr = rt.InvokeValue(vm.Keyword("target"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v2, nil
}
func cond_target_false(arg0 vm.Value) (vm.Value, error) {
	var v2 vm.Value
	var callErr error
	v2, callErr = rt.InvokeValue(vm.Keyword("false-target"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v2, nil
}
func block_of(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8873 vm.Value
	var arg__8881 vm.Value
	var arg__8882 vm.Value
	var arg__8884 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8873, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8873})
	if callErr != nil {
		return nil, callErr
	}
	arg__8881, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8882, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8881})
	if callErr != nil {
		return nil, callErr
	}
	arg__8884, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8882, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("block"), []vm.Value{arg__8884})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func fn_insts(arg0 vm.Value) (vm.Value, error) {
	var arg__8889 vm.Value
	var v4 vm.Value
	var callErr error
	arg__8889, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8889})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func set_source_infos_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__8905 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("source-infos")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8905, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("source-infos")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__8905, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func set_op_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__8922 vm.Value
	var v26 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("op")})
	if callErr != nil {
		return nil, callErr
	}
	arg__8922, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("insts"), arg1, vm.Keyword("op")})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__8922, arg2})
	if callErr != nil {
		return nil, callErr
	}
	v26, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc").Deref(), vm.Keyword("uses-dirty?"), vm.Boolean(true), vm.Keyword("uses-cache"), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	return v26, nil
}
func block_preds(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8942 vm.Value
	var arg__8950 vm.Value
	var arg__8951 vm.Value
	var arg__8953 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8942, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__8942})
	if callErr != nil {
		return nil, callErr
	}
	arg__8950, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8951, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__8950})
	if callErr != nil {
		return nil, callErr
	}
	arg__8953, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8951, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("preds"), []vm.Value{arg__8953})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func source_infos(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8959 vm.Value
	var arg__8967 vm.Value
	var arg__8968 vm.Value
	var arg__8970 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8959, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8959})
	if callErr != nil {
		return nil, callErr
	}
	arg__8967, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8968, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8967})
	if callErr != nil {
		return nil, callErr
	}
	arg__8970, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8968, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("source-infos"), []vm.Value{arg__8970})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func aux(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__8976 vm.Value
	var arg__8984 vm.Value
	var arg__8985 vm.Value
	var arg__8987 vm.Value
	var v13 vm.Value
	var callErr error
	arg__8976, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8976})
	if callErr != nil {
		return nil, callErr
	}
	arg__8984, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__8985, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__8984})
	if callErr != nil {
		return nil, callErr
	}
	arg__8987, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__8985, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("aux"), []vm.Value{arg__8987})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func branch_target_args(arg0 vm.Value) (vm.Value, error) {
	var v2 vm.Value
	var callErr error
	v2, callErr = rt.InvokeValue(vm.Keyword("args"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v2, nil
}
func block_set_id_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__9005 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("id")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9005, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("id")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__9005, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func block_params(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__9012 vm.Value
	var arg__9020 vm.Value
	var arg__9021 vm.Value
	var arg__9023 vm.Value
	var v13 vm.Value
	var callErr error
	arg__9012, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__9012})
	if callErr != nil {
		return nil, callErr
	}
	arg__9020, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__9021, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__9020})
	if callErr != nil {
		return nil, callErr
	}
	arg__9023, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__9021, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("params"), []vm.Value{arg__9023})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func block_set_term_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__9039 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("term")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9039, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("term")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__9039, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func cond_target_true(arg0 vm.Value) (vm.Value, error) {
	var v2 vm.Value
	var callErr error
	v2, callErr = rt.InvokeValue(vm.Keyword("true-target"), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v2, nil
}
func fn_variadic_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var arg__9047 vm.Value
	var v4 vm.Value
	var callErr error
	arg__9047, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("variadic?"), []vm.Value{arg__9047})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func block_set_insts_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__9063 vm.Value
	var v14 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("insts")})
	if callErr != nil {
		return nil, callErr
	}
	arg__9063, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{vm.Keyword("blocks"), arg1, vm.Keyword("insts")})
	if callErr != nil {
		return nil, callErr
	}
	v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{arg0, rt.LookupVar("clojure.core", "assoc-in").Deref(), arg__9063, arg2})
	if callErr != nil {
		return nil, callErr
	}
	return v14, nil
}
func fn_uses_dirty_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var arg__9069 vm.Value
	var v4 vm.Value
	var callErr error
	arg__9069, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("uses-dirty?"), []vm.Value{arg__9069})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func fn_blocks(arg0 vm.Value) (vm.Value, error) {
	var arg__9074 vm.Value
	var v4 vm.Value
	var callErr error
	arg__9074, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	v4, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__9074})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func block_term(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var arg__9080 vm.Value
	var arg__9088 vm.Value
	var arg__9089 vm.Value
	var arg__9091 vm.Value
	var v13 vm.Value
	var callErr error
	arg__9080, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__9080})
	if callErr != nil {
		return nil, callErr
	}
	arg__9088, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{arg1})
	if callErr != nil {
		return nil, callErr
	}
	arg__9089, callErr = rt.InvokeValue(vm.Keyword("blocks"), []vm.Value{arg__9088})
	if callErr != nil {
		return nil, callErr
	}
	arg__9091, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__9089, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v13, callErr = rt.InvokeValue(vm.Keyword("term"), []vm.Value{arg__9091})
	if callErr != nil {
		return nil, callErr
	}
	return v13, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.data.generated", map[string]vm.Value{"type-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("type-of: wrong number of arguments %d (expected 2)", len(args))
		}
		return type_of(args[0], args[1])
	}), "block-set-preds!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("block-set-preds!: wrong number of arguments %d (expected 3)", len(args))
		}
		return block_set_preds_BANG_(args[0], args[1], args[2])
	}), "set-block-of!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-block-of!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_block_of_BANG_(args[0], args[1], args[2])
	}), "block-set-params!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("block-set-params!: wrong number of arguments %d (expected 3)", len(args))
		}
		return block_set_params_BANG_(args[0], args[1], args[2])
	}), "set-refs!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-refs!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_refs_BANG_(args[0], args[1], args[2])
	}), "refs": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("refs: wrong number of arguments %d (expected 2)", len(args))
		}
		return refs(args[0], args[1])
	}), "fn-uses-cache": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-uses-cache: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_uses_cache(args[0])
	}), "block-insts": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-insts: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_insts(args[0], args[1])
	}), "set-aux!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-aux!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_aux_BANG_(args[0], args[1], args[2])
	}), "fn-consts": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-consts: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_consts(args[0])
	}), "block-id": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-id: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_id(args[0], args[1])
	}), "op": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("op: wrong number of arguments %d (expected 2)", len(args))
		}
		return op(args[0], args[1])
	}), "fn-name": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-name: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_name(args[0])
	}), "set-type-of!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-type-of!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_type_of_BANG_(args[0], args[1], args[2])
	}), "fn-entry": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-entry: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_entry(args[0])
	}), "fn-arity": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-arity: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_arity(args[0])
	}), "branch-target-target": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("branch-target-target: wrong number of arguments %d (expected 1)", len(args))
		}
		return branch_target_target(args[0])
	}), "cond-target-false": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("cond-target-false: wrong number of arguments %d (expected 1)", len(args))
		}
		return cond_target_false(args[0])
	}), "block-of": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-of: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_of(args[0], args[1])
	}), "fn-insts": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-insts: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_insts(args[0])
	}), "set-source-infos!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-source-infos!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_source_infos_BANG_(args[0], args[1], args[2])
	}), "set-op!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set-op!: wrong number of arguments %d (expected 3)", len(args))
		}
		return set_op_BANG_(args[0], args[1], args[2])
	}), "block-preds": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-preds: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_preds(args[0], args[1])
	}), "source-infos": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("source-infos: wrong number of arguments %d (expected 2)", len(args))
		}
		return source_infos(args[0], args[1])
	}), "aux": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("aux: wrong number of arguments %d (expected 2)", len(args))
		}
		return aux(args[0], args[1])
	}), "branch-target-args": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("branch-target-args: wrong number of arguments %d (expected 1)", len(args))
		}
		return branch_target_args(args[0])
	}), "block-set-id!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("block-set-id!: wrong number of arguments %d (expected 3)", len(args))
		}
		return block_set_id_BANG_(args[0], args[1], args[2])
	}), "block-params": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-params: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_params(args[0], args[1])
	}), "block-set-term!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("block-set-term!: wrong number of arguments %d (expected 3)", len(args))
		}
		return block_set_term_BANG_(args[0], args[1], args[2])
	}), "cond-target-true": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("cond-target-true: wrong number of arguments %d (expected 1)", len(args))
		}
		return cond_target_true(args[0])
	}), "fn-variadic?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-variadic?: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_variadic_QMARK_(args[0])
	}), "block-set-insts!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("block-set-insts!: wrong number of arguments %d (expected 3)", len(args))
		}
		return block_set_insts_BANG_(args[0], args[1], args[2])
	}), "fn-uses-dirty?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-uses-dirty?: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_uses_dirty_QMARK_(args[0])
	}), "fn-blocks": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fn-blocks: wrong number of arguments %d (expected 1)", len(args))
		}
		return fn_blocks(args[0])
	}), "block-term": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("block-term: wrong number of arguments %d (expected 2)", len(args))
		}
		return block_term(args[0], args[1])
	}),
	})
}
