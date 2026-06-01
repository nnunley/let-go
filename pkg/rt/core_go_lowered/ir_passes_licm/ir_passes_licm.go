package ir_passes_licm

import (
	"fmt"

	rt "github.com/nooga/let-go/pkg/rt"
	vm "github.com/nooga/let-go/pkg/vm"
)

func hoist_one_BANG_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var op vm.Value
	var refs vm.Value
	var aux vm.Value
	var clone vm.Value
	var from_block vm.Value
	var callErr error
	op, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	aux, callErr = rt.InvokeValue(rt.LookupVar("ir", "aux").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	clone, callErr = rt.InvokeValue(rt.LookupVar("ir", "add-inst").Deref(), []vm.Value{arg0, arg2, op, refs, aux})
	if callErr != nil {
		return nil, callErr
	}
	from_block, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "replace-all-uses!").Deref(), []vm.Value{arg0, arg1, clone})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "remove-inst!").Deref(), []vm.Value{arg0, from_block, arg1})
	if callErr != nil {
		return nil, callErr
	}
	return clone, nil
}
func back_edges(arg0 vm.Value) (vm.Value, error) {
	var for__a21484 vm.Value
	var for__iter21483 vm.Value
	var arg__24034 vm.Value
	var arg__26586 vm.Value
	var arg__26587 vm.Value
	var v45 vm.Value
	var callErr error
	for__a21484, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{for__a21484, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v14 vm.Value
		var callErr error
		v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "lazy-seq*").Deref(), []vm.Value{rt.BoxNativeFn(func() (vm.Value, error) {
			var tem__G__0 vm.Value
			var f_5 vm.Value
			var for__a21484_6 vm.Value
			var for__s_7 vm.Value
			var header vm.Value
			var for__a21482 vm.Value
			var for__iter21481 vm.Value
			var arg__23693 vm.Value
			var head__23702 vm.Value
			var arg__23706 vm.Value
			var arg__24006 vm.Value
			var arg__24007 vm.Value
			var head__24015 vm.Value
			var arg__24019 vm.Value
			var arg__24020 vm.Value
			var v81 vm.Value
			var v85 vm.Value
			var callErr error
			tem__G__0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg0})
			if callErr != nil {
				return nil, callErr
			}
			if vm.IsTruthy(tem__G__0) {
				f_5 = arg0
				for__a21484_6 = for__a21484
				for__s_7 = arg0
				goto b1
			} else {
				goto b2
			}
		b1:
			;
			header, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			for__a21482, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NIL})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{for__a21482, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
				var v18 vm.Value
				var callErr error
				v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "lazy-seq*").Deref(), []vm.Value{rt.BoxNativeFn(func() (vm.Value, error) {
					var tem__G__0 vm.Value
					var f_6 vm.Value
					var for__a21482_7 vm.Value
					var header_9 vm.Value
					var pred vm.Value
					var v34 vm.Value
					var v109 vm.Value
					var f_19 vm.Value
					var for__a21482_20 vm.Value
					var header_22 vm.Value
					var for__xs vm.Value
					var arg__23630 vm.Value
					var f_26 vm.Value
					var for__a21482_27 vm.Value
					var header_29 vm.Value
					var f_46 vm.Value
					var for__a21482_47 vm.Value
					var header_49 vm.Value
					var head__23639 vm.Value
					var arg__23643 vm.Value
					var v78 vm.Value
					var for__a21482_62 vm.Value
					var header_64 vm.Value
					var head__23645 vm.Value
					var arg__23661 vm.Value
					var v85 vm.Value
					var for__a21482_70 vm.Value
					var arg__23662 vm.Value
					var for__a21482_91 vm.Value
					var head__23670 vm.Value
					var arg__23674 vm.Value
					var arg__23675 vm.Value
					var v105 vm.Value
					var callErr error
					tem__G__0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{for__s_7})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(tem__G__0) {
						f_6 = f_5
						for__a21482_7 = for__a21482
						header_9 = header
						goto b1
					} else {
						goto b2
					}
				b1:
					;
					pred, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{tem__G__0})
					if callErr != nil {
						return nil, callErr
					}
					v34, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominates?").Deref(), []vm.Value{f_6, header_9, pred})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(v34) {
						f_19 = f_6
						for__a21482_20 = for__a21482_7
						header_22 = header_9
						for__xs = tem__G__0
						goto b4
					} else {
						f_26 = f_6
						for__a21482_27 = for__a21482_7
						header_29 = header_9
						for__xs = tem__G__0
						goto b5
					}
				b2:
					;
					v109 = vm.NIL
					goto b3
				b3:
					;
					return v109, nil
				b4:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_22})
					if callErr != nil {
						return nil, callErr
					}
					arg__23630, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_22})
					if callErr != nil {
						return nil, callErr
					}
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__23630})
					if callErr != nil {
						return nil, callErr
					}
					f_46 = f_19
					for__a21482_47 = for__a21482_20
					header_49 = header_22
					goto b6
				b5:
					;
					f_46 = f_26
					for__a21482_47 = for__a21482_27
					header_49 = header_29
					goto b6
				b6:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					head__23639, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482_47})
					if callErr != nil {
						return nil, callErr
					}
					arg__23643, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					_, callErr = rt.InvokeValue(head__23639, []vm.Value{arg__23643})
					if callErr != nil {
						return nil, callErr
					}
					v78, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominates?").Deref(), []vm.Value{f_46, header_49, pred})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(v78) {
						for__a21482_62 = for__a21482_47
						header_64 = header_49
						head__23645 = rt.LookupVar("clojure.core", "concat-list").Deref()
						goto b7
					} else {
						for__a21482_70 = for__a21482_47
						head__23645 = rt.LookupVar("clojure.core", "concat-list").Deref()
						goto b8
					}
				b7:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_64})
					if callErr != nil {
						return nil, callErr
					}
					arg__23661, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_64})
					if callErr != nil {
						return nil, callErr
					}
					v85, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__23661})
					if callErr != nil {
						return nil, callErr
					}
					arg__23662 = v85
					for__a21482_91 = for__a21482_62
					goto b9
				b8:
					;
					arg__23662 = vm.NIL
					for__a21482_91 = for__a21482_70
					goto b9
				b9:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					head__23670, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482_91})
					if callErr != nil {
						return nil, callErr
					}
					arg__23674, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					arg__23675, callErr = rt.InvokeValue(head__23670, []vm.Value{arg__23674})
					if callErr != nil {
						return nil, callErr
					}
					v105, callErr = rt.InvokeValue(head__23645, []vm.Value{arg__23662, arg__23675})
					if callErr != nil {
						return nil, callErr
					}
					v109 = v105
					goto b3
				})})
				if callErr != nil {
					return nil, callErr
				}
				return v18, nil
			})})
			if callErr != nil {
				return nil, callErr
			}
			for__iter21481, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{header, f_5})
			if callErr != nil {
				return nil, callErr
			}
			arg__23693, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{header, f_5})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(for__iter21481, []vm.Value{arg__23693})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			head__23702, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21484_6})
			if callErr != nil {
				return nil, callErr
			}
			arg__23706, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(head__23702, []vm.Value{arg__23706})
			if callErr != nil {
				return nil, callErr
			}
			for__a21482, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NIL})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{for__a21482, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
				var v18 vm.Value
				var callErr error
				v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "lazy-seq*").Deref(), []vm.Value{rt.BoxNativeFn(func() (vm.Value, error) {
					var tem__G__0 vm.Value
					var f_6 vm.Value
					var for__a21482_7 vm.Value
					var header_9 vm.Value
					var pred vm.Value
					var v34 vm.Value
					var v109 vm.Value
					var f_19 vm.Value
					var for__a21482_20 vm.Value
					var header_22 vm.Value
					var for__xs vm.Value
					var arg__23943 vm.Value
					var f_26 vm.Value
					var for__a21482_27 vm.Value
					var header_29 vm.Value
					var f_46 vm.Value
					var for__a21482_47 vm.Value
					var header_49 vm.Value
					var head__23952 vm.Value
					var arg__23956 vm.Value
					var v78 vm.Value
					var for__a21482_62 vm.Value
					var header_64 vm.Value
					var head__23958 vm.Value
					var arg__23974 vm.Value
					var v85 vm.Value
					var for__a21482_70 vm.Value
					var arg__23975 vm.Value
					var for__a21482_91 vm.Value
					var head__23983 vm.Value
					var arg__23987 vm.Value
					var arg__23988 vm.Value
					var v105 vm.Value
					var callErr error
					tem__G__0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{for__s_7})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(tem__G__0) {
						f_6 = f_5
						for__a21482_7 = for__a21482
						header_9 = header
						goto b1
					} else {
						goto b2
					}
				b1:
					;
					pred, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{tem__G__0})
					if callErr != nil {
						return nil, callErr
					}
					v34, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominates?").Deref(), []vm.Value{f_6, header_9, pred})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(v34) {
						f_19 = f_6
						for__a21482_20 = for__a21482_7
						header_22 = header_9
						for__xs = tem__G__0
						goto b4
					} else {
						f_26 = f_6
						for__a21482_27 = for__a21482_7
						header_29 = header_9
						for__xs = tem__G__0
						goto b5
					}
				b2:
					;
					v109 = vm.NIL
					goto b3
				b3:
					;
					return v109, nil
				b4:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_22})
					if callErr != nil {
						return nil, callErr
					}
					arg__23943, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_22})
					if callErr != nil {
						return nil, callErr
					}
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__23943})
					if callErr != nil {
						return nil, callErr
					}
					f_46 = f_19
					for__a21482_47 = for__a21482_20
					header_49 = header_22
					goto b6
				b5:
					;
					f_46 = f_26
					for__a21482_47 = for__a21482_27
					header_49 = header_29
					goto b6
				b6:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					head__23952, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482_47})
					if callErr != nil {
						return nil, callErr
					}
					arg__23956, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					_, callErr = rt.InvokeValue(head__23952, []vm.Value{arg__23956})
					if callErr != nil {
						return nil, callErr
					}
					v78, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominates?").Deref(), []vm.Value{f_46, header_49, pred})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(v78) {
						for__a21482_62 = for__a21482_47
						header_64 = header_49
						head__23958 = rt.LookupVar("clojure.core", "concat-list").Deref()
						goto b7
					} else {
						for__a21482_70 = for__a21482_47
						head__23958 = rt.LookupVar("clojure.core", "concat-list").Deref()
						goto b8
					}
				b7:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_64})
					if callErr != nil {
						return nil, callErr
					}
					arg__23974, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_64})
					if callErr != nil {
						return nil, callErr
					}
					v85, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__23974})
					if callErr != nil {
						return nil, callErr
					}
					arg__23975 = v85
					for__a21482_91 = for__a21482_62
					goto b9
				b8:
					;
					arg__23975 = vm.NIL
					for__a21482_91 = for__a21482_70
					goto b9
				b9:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					head__23983, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482_91})
					if callErr != nil {
						return nil, callErr
					}
					arg__23987, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					arg__23988, callErr = rt.InvokeValue(head__23983, []vm.Value{arg__23987})
					if callErr != nil {
						return nil, callErr
					}
					v105, callErr = rt.InvokeValue(head__23958, []vm.Value{arg__23975, arg__23988})
					if callErr != nil {
						return nil, callErr
					}
					v109 = v105
					goto b3
				})})
				if callErr != nil {
					return nil, callErr
				}
				return v18, nil
			})})
			if callErr != nil {
				return nil, callErr
			}
			for__iter21481, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{header, f_5})
			if callErr != nil {
				return nil, callErr
			}
			arg__24006, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{header, f_5})
			if callErr != nil {
				return nil, callErr
			}
			arg__24007, callErr = rt.InvokeValue(for__iter21481, []vm.Value{arg__24006})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			head__24015, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21484_6})
			if callErr != nil {
				return nil, callErr
			}
			arg__24019, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			arg__24020, callErr = rt.InvokeValue(head__24015, []vm.Value{arg__24019})
			if callErr != nil {
				return nil, callErr
			}
			v81, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat-list").Deref(), []vm.Value{arg__24007, arg__24020})
			if callErr != nil {
				return nil, callErr
			}
			v85 = v81
			goto b3
		b2:
			;
			v85 = vm.NIL
			goto b3
		b3:
			;
			return v85, nil
		})})
		if callErr != nil {
			return nil, callErr
		}
		return v14, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	for__iter21483, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21484})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__24034, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(for__iter21483, []vm.Value{arg__24034})
	if callErr != nil {
		return nil, callErr
	}
	for__a21484, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{for__a21484, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var v14 vm.Value
		var callErr error
		v14, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "lazy-seq*").Deref(), []vm.Value{rt.BoxNativeFn(func() (vm.Value, error) {
			var tem__G__0 vm.Value
			var f_5 vm.Value
			var for__a21484_6 vm.Value
			var for__s_7 vm.Value
			var header vm.Value
			var for__a21482 vm.Value
			var for__iter21481 vm.Value
			var arg__26245 vm.Value
			var head__26254 vm.Value
			var arg__26258 vm.Value
			var arg__26558 vm.Value
			var arg__26559 vm.Value
			var head__26567 vm.Value
			var arg__26571 vm.Value
			var arg__26572 vm.Value
			var v81 vm.Value
			var v85 vm.Value
			var callErr error
			tem__G__0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg0})
			if callErr != nil {
				return nil, callErr
			}
			if vm.IsTruthy(tem__G__0) {
				f_5 = arg0
				for__a21484_6 = for__a21484
				for__s_7 = arg0
				goto b1
			} else {
				goto b2
			}
		b1:
			;
			header, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			for__a21482, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NIL})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{for__a21482, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
				var v18 vm.Value
				var callErr error
				v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "lazy-seq*").Deref(), []vm.Value{rt.BoxNativeFn(func() (vm.Value, error) {
					var tem__G__0 vm.Value
					var f_6 vm.Value
					var for__a21482_7 vm.Value
					var header_9 vm.Value
					var pred vm.Value
					var v34 vm.Value
					var v109 vm.Value
					var f_19 vm.Value
					var for__a21482_20 vm.Value
					var header_22 vm.Value
					var for__xs vm.Value
					var arg__26182 vm.Value
					var f_26 vm.Value
					var for__a21482_27 vm.Value
					var header_29 vm.Value
					var f_46 vm.Value
					var for__a21482_47 vm.Value
					var header_49 vm.Value
					var head__26191 vm.Value
					var arg__26195 vm.Value
					var v78 vm.Value
					var for__a21482_62 vm.Value
					var header_64 vm.Value
					var head__26197 vm.Value
					var arg__26213 vm.Value
					var v85 vm.Value
					var for__a21482_70 vm.Value
					var arg__26214 vm.Value
					var for__a21482_91 vm.Value
					var head__26222 vm.Value
					var arg__26226 vm.Value
					var arg__26227 vm.Value
					var v105 vm.Value
					var callErr error
					tem__G__0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{for__s_7})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(tem__G__0) {
						f_6 = f_5
						for__a21482_7 = for__a21482
						header_9 = header
						goto b1
					} else {
						goto b2
					}
				b1:
					;
					pred, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{tem__G__0})
					if callErr != nil {
						return nil, callErr
					}
					v34, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominates?").Deref(), []vm.Value{f_6, header_9, pred})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(v34) {
						f_19 = f_6
						for__a21482_20 = for__a21482_7
						header_22 = header_9
						for__xs = tem__G__0
						goto b4
					} else {
						f_26 = f_6
						for__a21482_27 = for__a21482_7
						header_29 = header_9
						for__xs = tem__G__0
						goto b5
					}
				b2:
					;
					v109 = vm.NIL
					goto b3
				b3:
					;
					return v109, nil
				b4:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_22})
					if callErr != nil {
						return nil, callErr
					}
					arg__26182, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_22})
					if callErr != nil {
						return nil, callErr
					}
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__26182})
					if callErr != nil {
						return nil, callErr
					}
					f_46 = f_19
					for__a21482_47 = for__a21482_20
					header_49 = header_22
					goto b6
				b5:
					;
					f_46 = f_26
					for__a21482_47 = for__a21482_27
					header_49 = header_29
					goto b6
				b6:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					head__26191, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482_47})
					if callErr != nil {
						return nil, callErr
					}
					arg__26195, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					_, callErr = rt.InvokeValue(head__26191, []vm.Value{arg__26195})
					if callErr != nil {
						return nil, callErr
					}
					v78, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominates?").Deref(), []vm.Value{f_46, header_49, pred})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(v78) {
						for__a21482_62 = for__a21482_47
						header_64 = header_49
						head__26197 = rt.LookupVar("clojure.core", "concat-list").Deref()
						goto b7
					} else {
						for__a21482_70 = for__a21482_47
						head__26197 = rt.LookupVar("clojure.core", "concat-list").Deref()
						goto b8
					}
				b7:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_64})
					if callErr != nil {
						return nil, callErr
					}
					arg__26213, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_64})
					if callErr != nil {
						return nil, callErr
					}
					v85, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__26213})
					if callErr != nil {
						return nil, callErr
					}
					arg__26214 = v85
					for__a21482_91 = for__a21482_62
					goto b9
				b8:
					;
					arg__26214 = vm.NIL
					for__a21482_91 = for__a21482_70
					goto b9
				b9:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					head__26222, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482_91})
					if callErr != nil {
						return nil, callErr
					}
					arg__26226, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					arg__26227, callErr = rt.InvokeValue(head__26222, []vm.Value{arg__26226})
					if callErr != nil {
						return nil, callErr
					}
					v105, callErr = rt.InvokeValue(head__26197, []vm.Value{arg__26214, arg__26227})
					if callErr != nil {
						return nil, callErr
					}
					v109 = v105
					goto b3
				})})
				if callErr != nil {
					return nil, callErr
				}
				return v18, nil
			})})
			if callErr != nil {
				return nil, callErr
			}
			for__iter21481, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{header, f_5})
			if callErr != nil {
				return nil, callErr
			}
			arg__26245, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{header, f_5})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(for__iter21481, []vm.Value{arg__26245})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			head__26254, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21484_6})
			if callErr != nil {
				return nil, callErr
			}
			arg__26258, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(head__26254, []vm.Value{arg__26258})
			if callErr != nil {
				return nil, callErr
			}
			for__a21482, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NIL})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "reset!").Deref(), []vm.Value{for__a21482, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
				var v18 vm.Value
				var callErr error
				v18, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "lazy-seq*").Deref(), []vm.Value{rt.BoxNativeFn(func() (vm.Value, error) {
					var tem__G__0 vm.Value
					var f_6 vm.Value
					var for__a21482_7 vm.Value
					var header_9 vm.Value
					var pred vm.Value
					var v34 vm.Value
					var v109 vm.Value
					var f_19 vm.Value
					var for__a21482_20 vm.Value
					var header_22 vm.Value
					var for__xs vm.Value
					var arg__26495 vm.Value
					var f_26 vm.Value
					var for__a21482_27 vm.Value
					var header_29 vm.Value
					var f_46 vm.Value
					var for__a21482_47 vm.Value
					var header_49 vm.Value
					var head__26504 vm.Value
					var arg__26508 vm.Value
					var v78 vm.Value
					var for__a21482_62 vm.Value
					var header_64 vm.Value
					var head__26510 vm.Value
					var arg__26526 vm.Value
					var v85 vm.Value
					var for__a21482_70 vm.Value
					var arg__26527 vm.Value
					var for__a21482_91 vm.Value
					var head__26535 vm.Value
					var arg__26539 vm.Value
					var arg__26540 vm.Value
					var v105 vm.Value
					var callErr error
					tem__G__0, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{for__s_7})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(tem__G__0) {
						f_6 = f_5
						for__a21482_7 = for__a21482
						header_9 = header
						goto b1
					} else {
						goto b2
					}
				b1:
					;
					pred, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{tem__G__0})
					if callErr != nil {
						return nil, callErr
					}
					v34, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominates?").Deref(), []vm.Value{f_6, header_9, pred})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(v34) {
						f_19 = f_6
						for__a21482_20 = for__a21482_7
						header_22 = header_9
						for__xs = tem__G__0
						goto b4
					} else {
						f_26 = f_6
						for__a21482_27 = for__a21482_7
						header_29 = header_9
						for__xs = tem__G__0
						goto b5
					}
				b2:
					;
					v109 = vm.NIL
					goto b3
				b3:
					;
					return v109, nil
				b4:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_22})
					if callErr != nil {
						return nil, callErr
					}
					arg__26495, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_22})
					if callErr != nil {
						return nil, callErr
					}
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__26495})
					if callErr != nil {
						return nil, callErr
					}
					f_46 = f_19
					for__a21482_47 = for__a21482_20
					header_49 = header_22
					goto b6
				b5:
					;
					f_46 = f_26
					for__a21482_47 = for__a21482_27
					header_49 = header_29
					goto b6
				b6:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					head__26504, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482_47})
					if callErr != nil {
						return nil, callErr
					}
					arg__26508, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					_, callErr = rt.InvokeValue(head__26504, []vm.Value{arg__26508})
					if callErr != nil {
						return nil, callErr
					}
					v78, callErr = rt.InvokeValue(rt.LookupVar("ir.dominance", "dominates?").Deref(), []vm.Value{f_46, header_49, pred})
					if callErr != nil {
						return nil, callErr
					}
					if vm.IsTruthy(v78) {
						for__a21482_62 = for__a21482_47
						header_64 = header_49
						head__26510 = rt.LookupVar("clojure.core", "concat-list").Deref()
						goto b7
					} else {
						for__a21482_70 = for__a21482_47
						head__26510 = rt.LookupVar("clojure.core", "concat-list").Deref()
						goto b8
					}
				b7:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_64})
					if callErr != nil {
						return nil, callErr
					}
					arg__26526, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{pred, header_64})
					if callErr != nil {
						return nil, callErr
					}
					v85, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "list").Deref(), []vm.Value{arg__26526})
					if callErr != nil {
						return nil, callErr
					}
					arg__26527 = v85
					for__a21482_91 = for__a21482_62
					goto b9
				b8:
					;
					arg__26527 = vm.NIL
					for__a21482_91 = for__a21482_70
					goto b9
				b9:
					;
					_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					head__26535, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482_91})
					if callErr != nil {
						return nil, callErr
					}
					arg__26539, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{for__xs})
					if callErr != nil {
						return nil, callErr
					}
					arg__26540, callErr = rt.InvokeValue(head__26535, []vm.Value{arg__26539})
					if callErr != nil {
						return nil, callErr
					}
					v105, callErr = rt.InvokeValue(head__26510, []vm.Value{arg__26527, arg__26540})
					if callErr != nil {
						return nil, callErr
					}
					v109 = v105
					goto b3
				})})
				if callErr != nil {
					return nil, callErr
				}
				return v18, nil
			})})
			if callErr != nil {
				return nil, callErr
			}
			for__iter21481, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21482})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{header, f_5})
			if callErr != nil {
				return nil, callErr
			}
			arg__26558, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{header, f_5})
			if callErr != nil {
				return nil, callErr
			}
			arg__26559, callErr = rt.InvokeValue(for__iter21481, []vm.Value{arg__26558})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			head__26567, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21484_6})
			if callErr != nil {
				return nil, callErr
			}
			arg__26571, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "rest").Deref(), []vm.Value{tem__G__0})
			if callErr != nil {
				return nil, callErr
			}
			arg__26572, callErr = rt.InvokeValue(head__26567, []vm.Value{arg__26571})
			if callErr != nil {
				return nil, callErr
			}
			v81, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "concat-list").Deref(), []vm.Value{arg__26559, arg__26572})
			if callErr != nil {
				return nil, callErr
			}
			v85 = v81
			goto b3
		b2:
			;
			v85 = vm.NIL
			goto b3
		b3:
			;
			return v85, nil
		})})
		if callErr != nil {
			return nil, callErr
		}
		return v14, nil
	})})
	if callErr != nil {
		return nil, callErr
	}
	for__iter21483, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{for__a21484})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__26586, callErr = rt.InvokeValue(rt.LookupVar("ir", "blocks").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__26587, callErr = rt.InvokeValue(for__iter21483, []vm.Value{arg__26586})
	if callErr != nil {
		return nil, callErr
	}
	v45, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vec").Deref(), []vm.Value{arg__26587})
	if callErr != nil {
		return nil, callErr
	}
	return v45, nil
}
func pure_op_QMARK_(arg0 vm.Value) (vm.Value, error) {
	var v4 vm.Value
	var callErr error
	v4, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{rt.LookupVar("ir.passes.licm", "pure-ops").Deref(), arg0})
	if callErr != nil {
		return nil, callErr
	}
	return v4, nil
}
func unique_pre_header(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__26607 vm.Value
	var outs vm.Value
	var arg__26612 vm.Value
	var v21 bool
	var v24 vm.Value
	var v28 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__26607, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-preds").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	outs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "remove").Deref(), []vm.Value{arg2, arg__26607})
	if callErr != nil {
		return nil, callErr
	}
	arg__26612, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{outs})
	if callErr != nil {
		return nil, callErr
	}
	v21 = arg__26612 == vm.Int(1)
	if v21 {
		goto b1
	} else {
		goto b2
	}
b1:
	;
	v24, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{outs})
	if callErr != nil {
		return nil, callErr
	}
	v28 = v24
	goto b3
b2:
	;
	v28 = vm.NIL
	goto b3
b3:
	;
	return v28, nil
}
func collect_hoistable(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var v7 vm.Value
	var invariant vm.Value
	var loop_set vm.Value
	var f vm.Value
	var arg__26767 vm.Value
	var arg__26774 vm.Value
	var arg__26775 vm.Value
	var arg__26782 vm.Value
	var arg__26789 vm.Value
	var arg__26790 vm.Value
	var arg__26791 vm.Value
	var arg__26943 vm.Value
	var arg__26950 vm.Value
	var arg__26951 vm.Value
	var arg__26958 vm.Value
	var arg__26965 vm.Value
	var arg__26966 vm.Value
	var arg__26967 vm.Value
	var arg__26968 vm.Value
	var arg__27121 vm.Value
	var arg__27128 vm.Value
	var arg__27129 vm.Value
	var arg__27136 vm.Value
	var arg__27143 vm.Value
	var arg__27144 vm.Value
	var arg__27145 vm.Value
	var arg__27297 vm.Value
	var arg__27304 vm.Value
	var arg__27305 vm.Value
	var arg__27312 vm.Value
	var arg__27319 vm.Value
	var arg__27320 vm.Value
	var arg__27321 vm.Value
	var arg__27322 vm.Value
	var arg__27323 vm.Value
	var new_invariant vm.Value
	var v131 bool
	var v135 vm.Value
	var callErr error
	v7, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "hash-set").Deref(), []vm.Value{})
	if callErr != nil {
		return nil, callErr
	}
	invariant = v7
	loop_set = arg1
	f = arg0
	goto b1
b1:
	;
	arg__26767, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__26767})
	if callErr != nil {
		return nil, callErr
	}
	arg__26774, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__26775, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__26774})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__26775})
	if callErr != nil {
		return nil, callErr
	}
	arg__26782, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__26782})
	if callErr != nil {
		return nil, callErr
	}
	arg__26789, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__26790, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__26789})
	if callErr != nil {
		return nil, callErr
	}
	arg__26791, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__26790})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{arg__26791})
	if callErr != nil {
		return nil, callErr
	}
	arg__26943, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__26943})
	if callErr != nil {
		return nil, callErr
	}
	arg__26950, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__26951, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__26950})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__26951})
	if callErr != nil {
		return nil, callErr
	}
	arg__26958, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__26958})
	if callErr != nil {
		return nil, callErr
	}
	arg__26965, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__26966, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__26965})
	if callErr != nil {
		return nil, callErr
	}
	arg__26967, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__26966})
	if callErr != nil {
		return nil, callErr
	}
	arg__26968, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{arg__26967})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__26802 vm.Value
		var and__x vm.Value
		var nid vm.Value
		var f_9 vm.Value
		var invariant_10 vm.Value
		var loop_set_11 vm.Value
		var arg__26817 vm.Value
		var v140 vm.Value
		var f_28 vm.Value
		var invariant_29 vm.Value
		var loop_set_30 vm.Value
		var arg__26830 vm.Value
		var v132 vm.Value
		var f_44 vm.Value
		var invariant_45 vm.Value
		var loop_set_46 vm.Value
		var arg__26843 vm.Value
		var v124 vm.Value
		var f_61 vm.Value
		var invariant_62 vm.Value
		var loop_set_63 vm.Value
		var arg__26858 vm.Value
		var v116 vm.Value
		var f_80 vm.Value
		var invariant_81 vm.Value
		var loop_set_82 vm.Value
		var arg__26937 vm.Value
		var v105 vm.Value
		var v108 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(invariant, []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__26802, callErr = rt.InvokeValue(invariant, []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__26802})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			nid = arg0
			f_9 = f
			invariant_10 = invariant
			loop_set_11 = loop_set
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_9})
		if callErr != nil {
			return nil, callErr
		}
		arg__26817, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_9})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("invalid"), arg__26817})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_28 = f_9
			invariant_29 = invariant_10
			loop_set_30 = loop_set_11
			goto b4
		} else {
			goto b5
		}
	b2:
		;
		v140 = and__x
		goto b3
	b3:
		;
		return v140, nil
	b4:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{nid, f_28})
		if callErr != nil {
			return nil, callErr
		}
		arg__26830, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{nid, f_28})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(loop_set_30, []vm.Value{arg__26830})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_44 = f_28
			invariant_45 = invariant_29
			loop_set_46 = loop_set_30
			goto b7
		} else {
			goto b8
		}
	b5:
		;
		v132 = and__x
		goto b6
	b6:
		;
		v140 = v132
		goto b3
	b7:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_44})
		if callErr != nil {
			return nil, callErr
		}
		arg__26843, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_44})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "pure-op?").Deref(), []vm.Value{arg__26843})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_61 = f_44
			invariant_62 = invariant_45
			loop_set_63 = loop_set_46
			goto b10
		} else {
			goto b11
		}
	b8:
		;
		v124 = and__x
		goto b9
	b9:
		;
		v132 = v124
		goto b6
	b10:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_61})
		if callErr != nil {
			return nil, callErr
		}
		arg__26858, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_61})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("block-arg"), arg__26858})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_80 = f_61
			invariant_81 = invariant_62
			loop_set_82 = loop_set_63
			goto b13
		} else {
			goto b14
		}
	b11:
		;
		v116 = and__x
		goto b12
	b12:
		;
		v124 = v116
		goto b9
	b13:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f_80})
		if callErr != nil {
			return nil, callErr
		}
		arg__26937, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f_80})
		if callErr != nil {
			return nil, callErr
		}
		v105, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var or__x vm.Value
			var r vm.Value
			var f_11 vm.Value
			var loop_set_13 vm.Value
			var arg__26914 vm.Value
			var arg__26929 vm.Value
			var arg__26930 vm.Value
			var v28 vm.Value
			var v30 vm.Value
			var callErr error
			or__x, callErr = rt.InvokeValue(invariant_81, []vm.Value{arg0})
			if callErr != nil {
				return nil, callErr
			}
			if vm.IsTruthy(or__x) {
				goto b1
			} else {
				r = arg0
				f_11 = f_80
				loop_set_13 = loop_set_82
				goto b2
			}
		b1:
			;
			v30 = or__x
			goto b3
		b2:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{r, f_11})
			if callErr != nil {
				return nil, callErr
			}
			arg__26914, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{r, f_11})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(loop_set_13, []vm.Value{arg__26914})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{r, f_11})
			if callErr != nil {
				return nil, callErr
			}
			arg__26929, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{r, f_11})
			if callErr != nil {
				return nil, callErr
			}
			arg__26930, callErr = rt.InvokeValue(loop_set_13, []vm.Value{arg__26929})
			if callErr != nil {
				return nil, callErr
			}
			v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__26930})
			if callErr != nil {
				return nil, callErr
			}
			v30 = v28
			goto b3
		b3:
			;
			return v30, nil
		}), arg__26937})
		if callErr != nil {
			return nil, callErr
		}
		v108 = v105
		goto b15
	b14:
		;
		v108 = and__x
		goto b15
	b15:
		;
		v116 = v108
		goto b12
	}), arg__26968})
	if callErr != nil {
		return nil, callErr
	}
	arg__27121, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__27121})
	if callErr != nil {
		return nil, callErr
	}
	arg__27128, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__27129, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__27128})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__27129})
	if callErr != nil {
		return nil, callErr
	}
	arg__27136, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__27136})
	if callErr != nil {
		return nil, callErr
	}
	arg__27143, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__27144, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__27143})
	if callErr != nil {
		return nil, callErr
	}
	arg__27145, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__27144})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{arg__27145})
	if callErr != nil {
		return nil, callErr
	}
	arg__27297, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__27297})
	if callErr != nil {
		return nil, callErr
	}
	arg__27304, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__27305, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__27304})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__27305})
	if callErr != nil {
		return nil, callErr
	}
	arg__27312, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__27312})
	if callErr != nil {
		return nil, callErr
	}
	arg__27319, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{f})
	if callErr != nil {
		return nil, callErr
	}
	arg__27320, callErr = rt.InvokeValue(vm.Keyword("insts"), []vm.Value{arg__27319})
	if callErr != nil {
		return nil, callErr
	}
	arg__27321, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "count").Deref(), []vm.Value{arg__27320})
	if callErr != nil {
		return nil, callErr
	}
	arg__27322, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "range").Deref(), []vm.Value{arg__27321})
	if callErr != nil {
		return nil, callErr
	}
	arg__27323, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "filter").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__27156 vm.Value
		var and__x vm.Value
		var nid vm.Value
		var f_9 vm.Value
		var invariant_10 vm.Value
		var loop_set_11 vm.Value
		var arg__27171 vm.Value
		var v140 vm.Value
		var f_28 vm.Value
		var invariant_29 vm.Value
		var loop_set_30 vm.Value
		var arg__27184 vm.Value
		var v132 vm.Value
		var f_44 vm.Value
		var invariant_45 vm.Value
		var loop_set_46 vm.Value
		var arg__27197 vm.Value
		var v124 vm.Value
		var f_61 vm.Value
		var invariant_62 vm.Value
		var loop_set_63 vm.Value
		var arg__27212 vm.Value
		var v116 vm.Value
		var f_80 vm.Value
		var invariant_81 vm.Value
		var loop_set_82 vm.Value
		var arg__27291 vm.Value
		var v105 vm.Value
		var v108 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(invariant, []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__27156, callErr = rt.InvokeValue(invariant, []vm.Value{arg0})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__27156})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			nid = arg0
			f_9 = f
			invariant_10 = invariant
			loop_set_11 = loop_set
			goto b1
		} else {
			goto b2
		}
	b1:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_9})
		if callErr != nil {
			return nil, callErr
		}
		arg__27171, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_9})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("invalid"), arg__27171})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_28 = f_9
			invariant_29 = invariant_10
			loop_set_30 = loop_set_11
			goto b4
		} else {
			goto b5
		}
	b2:
		;
		v140 = and__x
		goto b3
	b3:
		;
		return v140, nil
	b4:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{nid, f_28})
		if callErr != nil {
			return nil, callErr
		}
		arg__27184, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{nid, f_28})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(loop_set_30, []vm.Value{arg__27184})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_44 = f_28
			invariant_45 = invariant_29
			loop_set_46 = loop_set_30
			goto b7
		} else {
			goto b8
		}
	b5:
		;
		v132 = and__x
		goto b6
	b6:
		;
		v140 = v132
		goto b3
	b7:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_44})
		if callErr != nil {
			return nil, callErr
		}
		arg__27197, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_44})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "pure-op?").Deref(), []vm.Value{arg__27197})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_61 = f_44
			invariant_62 = invariant_45
			loop_set_63 = loop_set_46
			goto b10
		} else {
			goto b11
		}
	b8:
		;
		v124 = and__x
		goto b9
	b9:
		;
		v132 = v124
		goto b6
	b10:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_61})
		if callErr != nil {
			return nil, callErr
		}
		arg__27212, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{nid, f_61})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("block-arg"), arg__27212})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			f_80 = f_61
			invariant_81 = invariant_62
			loop_set_82 = loop_set_63
			goto b13
		} else {
			goto b14
		}
	b11:
		;
		v116 = and__x
		goto b12
	b12:
		;
		v124 = v116
		goto b9
	b13:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f_80})
		if callErr != nil {
			return nil, callErr
		}
		arg__27291, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f_80})
		if callErr != nil {
			return nil, callErr
		}
		v105, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
			var or__x vm.Value
			var r vm.Value
			var f_11 vm.Value
			var loop_set_13 vm.Value
			var arg__27268 vm.Value
			var arg__27283 vm.Value
			var arg__27284 vm.Value
			var v28 vm.Value
			var v30 vm.Value
			var callErr error
			or__x, callErr = rt.InvokeValue(invariant_81, []vm.Value{arg0})
			if callErr != nil {
				return nil, callErr
			}
			if vm.IsTruthy(or__x) {
				goto b1
			} else {
				r = arg0
				f_11 = f_80
				loop_set_13 = loop_set_82
				goto b2
			}
		b1:
			;
			v30 = or__x
			goto b3
		b2:
			;
			_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{r, f_11})
			if callErr != nil {
				return nil, callErr
			}
			arg__27268, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{r, f_11})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(loop_set_13, []vm.Value{arg__27268})
			if callErr != nil {
				return nil, callErr
			}
			_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{r, f_11})
			if callErr != nil {
				return nil, callErr
			}
			arg__27283, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{r, f_11})
			if callErr != nil {
				return nil, callErr
			}
			arg__27284, callErr = rt.InvokeValue(loop_set_13, []vm.Value{arg__27283})
			if callErr != nil {
				return nil, callErr
			}
			v28, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__27284})
			if callErr != nil {
				return nil, callErr
			}
			v30 = v28
			goto b3
		b3:
			;
			return v30, nil
		}), arg__27291})
		if callErr != nil {
			return nil, callErr
		}
		v108 = v105
		goto b15
	b14:
		;
		v108 = and__x
		goto b15
	b15:
		;
		v116 = v108
		goto b12
	}), arg__27322})
	if callErr != nil {
		return nil, callErr
	}
	new_invariant, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "into").Deref(), []vm.Value{invariant, arg__27323})
	if callErr != nil {
		return nil, callErr
	}
	v131 = invariant == new_invariant
	if v131 {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	v135 = invariant
	goto b4
b3:
	;
	invariant = new_invariant
	goto b1
b4:
	;
	return v135, nil
}
func loop_blocks(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var pred vm.Value
	var header vm.Value
	var arg__27358 vm.Value
	var v20 vm.Value
	var callErr error
	pred, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(0), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	header, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1), vm.NIL})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "reachable-without").Deref(), []vm.Value{arg0, pred, header})
	if callErr != nil {
		return nil, callErr
	}
	arg__27358, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "reachable-without").Deref(), []vm.Value{arg0, pred, header})
	if callErr != nil {
		return nil, callErr
	}
	v20, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "conj").Deref(), []vm.Value{arg__27358, header})
	if callErr != nil {
		return nil, callErr
	}
	return v20, nil
}
func find_in_loop_users(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var arg__27369 vm.Value
	var clone_users vm.Value
	var acc vm.Value
	var v29 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__27369, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	clone_users, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg__27369, arg1})
	if callErr != nil {
		return nil, callErr
	}
	acc, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "atom").Deref(), []vm.Value{vm.NewArrayVector([]vm.Value{})})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir", "uses-for-each").Deref(), []vm.Value{clone_users, rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__27446 vm.Value
		var and__x vm.Value
		var nid vm.Value
		var acc_6 vm.Value
		var v88 vm.Value
		var v92 vm.Value
		var acc_24 vm.Value
		var clone_25 vm.Value
		var f_26 vm.Value
		var loop_set_27 vm.Value
		var arg__27459 vm.Value
		var acc_30 vm.Value
		var v77 vm.Value
		var acc_79 vm.Value
		var acc_42 vm.Value
		var clone_43 vm.Value
		var f_44 vm.Value
		var arg__27478 vm.Value
		var v65 vm.Value
		var acc_48 vm.Value
		var v68 vm.Value
		var acc_70 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__27446, callErr = rt.InvokeValue(rt.LookupVar("ir", "op").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not=").Deref(), []vm.Value{vm.Keyword("invalid"), arg__27446})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			nid = arg0
			acc_24 = acc
			clone_25 = arg1
			f_26 = arg0
			loop_set_27 = arg2
			goto b4
		} else {
			nid = arg0
			acc_30 = acc
			goto b5
		}
	b1:
		;
		v88, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "swap!").Deref(), []vm.Value{acc_6, rt.LookupVar("clojure.core", "conj").Deref(), nid})
		if callErr != nil {
			return nil, callErr
		}
		v92 = v88
		goto b3
	b2:
		;
		v92 = vm.NIL
		goto b3
	b3:
		;
		return v92, nil
	b4:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{nid, f_26})
		if callErr != nil {
			return nil, callErr
		}
		arg__27459, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{nid, f_26})
		if callErr != nil {
			return nil, callErr
		}
		and__x, callErr = rt.InvokeValue(loop_set_27, []vm.Value{arg__27459})
		if callErr != nil {
			return nil, callErr
		}
		if vm.IsTruthy(and__x) {
			acc_42 = acc_24
			clone_43 = clone_25
			f_44 = f_26
			goto b7
		} else {
			acc_48 = acc_24
			goto b8
		}
	b5:
		;
		v77 = and__x
		acc_79 = acc_30
		goto b6
	b6:
		;
		if vm.IsTruthy(v77) {
			acc_6 = acc_79
			goto b1
		} else {
			goto b2
		}
	b7:
		;
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f_44})
		if callErr != nil {
			return nil, callErr
		}
		arg__27478, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{nid, f_44})
		if callErr != nil {
			return nil, callErr
		}
		v65, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "some").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) vm.Value {
			var v2 vm.Value
			v2 = vm.Boolean(arg0 == clone_43)
			return v2
		}), arg__27478})
		if callErr != nil {
			return nil, callErr
		}
		v68 = v65
		acc_70 = acc_42
		goto b9
	b8:
		;
		v68 = and__x
		acc_70 = acc_48
		goto b9
	b9:
		;
		v77 = v68
		acc_79 = acc_70
		goto b6
	})})
	if callErr != nil {
		return nil, callErr
	}
	v29, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "deref").Deref(), []vm.Value{acc})
	if callErr != nil {
		return nil, callErr
	}
	return v29, nil
}
func licm_one_loop(arg0 vm.Value, arg1 vm.Value) (vm.Value, error) {
	var header vm.Value
	var loop_set vm.Value
	var pre_header vm.Value
	var f vm.Value
	var hoistable vm.Value
	var arg__28265 vm.Value
	var hoisted_pairs vm.Value
	var v53 vm.Value
	var v81 vm.Value
	var arg__28288 vm.Value
	var v66 vm.Value
	var v70 vm.Value
	var callErr error
	header, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "nth").Deref(), []vm.Value{arg1, vm.Int(1)})
	if callErr != nil {
		return nil, callErr
	}
	loop_set, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "loop-blocks").Deref(), []vm.Value{arg0, arg1})
	if callErr != nil {
		return nil, callErr
	}
	pre_header, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "unique-pre-header").Deref(), []vm.Value{arg0, header, loop_set})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(pre_header) {
		f = arg0
		goto b1
	} else {
		goto b2
	}
b1:
	;
	hoistable, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "collect-hoistable").Deref(), []vm.Value{f, loop_set})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort").Deref(), []vm.Value{hoistable})
	if callErr != nil {
		return nil, callErr
	}
	arg__28265, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "sort").Deref(), []vm.Value{hoistable})
	if callErr != nil {
		return nil, callErr
	}
	hoisted_pairs, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "mapv").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__28260 vm.Value
		var v6 vm.Value
		var callErr error
		arg__28260, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "hoist-one!").Deref(), []vm.Value{f, arg0, pre_header})
		if callErr != nil {
			return nil, callErr
		}
		v6, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "vector").Deref(), []vm.Value{arg0, arg__28260})
		if callErr != nil {
			return nil, callErr
		}
		return v6, nil
	}), arg__28265})
	if callErr != nil {
		return nil, callErr
	}
	v53, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{hoisted_pairs})
	if callErr != nil {
		return nil, callErr
	}
	if vm.IsTruthy(v53) {
		goto b4
	} else {
		goto b5
	}
b2:
	;
	v81 = vm.NIL
	goto b3
b3:
	;
	return v81, nil
b4:
	;
	_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("body"), loop_set, vm.Keyword("preheader"), pre_header, vm.Keyword("header"), header})
	if callErr != nil {
		return nil, callErr
	}
	arg__28288, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "array-map").Deref(), []vm.Value{vm.Keyword("body"), loop_set, vm.Keyword("preheader"), pre_header, vm.Keyword("header"), header})
	if callErr != nil {
		return nil, callErr
	}
	v66, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "thread-hoisted-through-header!").Deref(), []vm.Value{f, arg__28288, hoisted_pairs})
	if callErr != nil {
		return nil, callErr
	}
	v70 = v66
	goto b6
b5:
	;
	v70 = vm.NIL
	goto b6
b6:
	;
	v81 = v70
	goto b3
}
func operand_defined_outside_QMARK_(arg0 vm.Value, arg1 vm.Value, arg2 vm.Value) (vm.Value, error) {
	var refs vm.Value
	var v15 vm.Value
	var callErr error
	refs, callErr = rt.InvokeValue(rt.LookupVar("ir", "refs").Deref(), []vm.Value{arg1, arg0})
	if callErr != nil {
		return nil, callErr
	}
	v15, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "every?").Deref(), []vm.Value{rt.BoxNativeFn(func(arg0 vm.Value) (vm.Value, error) {
		var arg__28345 vm.Value
		var arg__28362 vm.Value
		var arg__28363 vm.Value
		var v16 vm.Value
		var callErr error
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__28345, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg2, arg__28345})
		if callErr != nil {
			return nil, callErr
		}
		_, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__28362, callErr = rt.InvokeValue(rt.LookupVar("ir", "block-of").Deref(), []vm.Value{arg0, arg0})
		if callErr != nil {
			return nil, callErr
		}
		arg__28363, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "contains?").Deref(), []vm.Value{arg2, arg__28362})
		if callErr != nil {
			return nil, callErr
		}
		v16, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "not").Deref(), []vm.Value{arg__28363})
		if callErr != nil {
			return nil, callErr
		}
		return v16, nil
	}), refs})
	if callErr != nil {
		return nil, callErr
	}
	return v15, nil
}
func licm(arg0 vm.Value) (vm.Value, error) {
	var arg__28376 vm.Value
	var doseq_seq__28366 vm.Value
	var doseq_loop__28367 vm.Value
	var f vm.Value
	var be vm.Value
	var v23 vm.Value
	var callErr error
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "back-edges").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	arg__28376, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "back-edges").Deref(), []vm.Value{arg0})
	if callErr != nil {
		return nil, callErr
	}
	doseq_seq__28366, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "seq").Deref(), []vm.Value{arg__28376})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__28367 = doseq_seq__28366
	f = arg0
	goto b1
b1:
	;
	if vm.IsTruthy(doseq_loop__28367) {
		goto b2
	} else {
		goto b3
	}
b2:
	;
	be, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "first").Deref(), []vm.Value{doseq_loop__28367})
	if callErr != nil {
		return nil, callErr
	}
	_, callErr = rt.InvokeValue(rt.LookupVar("ir.passes.licm", "licm-one-loop").Deref(), []vm.Value{f, be})
	if callErr != nil {
		return nil, callErr
	}
	v23, callErr = rt.InvokeValue(rt.LookupVar("clojure.core", "next").Deref(), []vm.Value{doseq_loop__28367})
	if callErr != nil {
		return nil, callErr
	}
	doseq_loop__28367 = v23
	goto b1
b3:
	;
	goto b4
b4:
	;
	return f, nil
}
func __gogen_wrap(fn func(args []vm.Value) (vm.Value, error)) vm.Value {
	v, _ := vm.NativeFnType.Wrap(fn)
	return v
}
func init() {
	rt.RegisterGoOverrides("ir.passes.licm", map[string]vm.Value{"hoist-one!": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("hoist-one!: wrong number of arguments %d (expected 3)", len(args))
		}
		return hoist_one_BANG_(args[0], args[1], args[2])
	}), "back-edges": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("back-edges: wrong number of arguments %d (expected 1)", len(args))
		}
		return back_edges(args[0])
	}), "pure-op?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("pure-op?: wrong number of arguments %d (expected 1)", len(args))
		}
		return pure_op_QMARK_(args[0])
	}), "unique-pre-header": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("unique-pre-header: wrong number of arguments %d (expected 3)", len(args))
		}
		return unique_pre_header(args[0], args[1], args[2])
	}), "collect-hoistable": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("collect-hoistable: wrong number of arguments %d (expected 2)", len(args))
		}
		return collect_hoistable(args[0], args[1])
	}), "loop-blocks": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("loop-blocks: wrong number of arguments %d (expected 2)", len(args))
		}
		return loop_blocks(args[0], args[1])
	}), "find-in-loop-users": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("find-in-loop-users: wrong number of arguments %d (expected 3)", len(args))
		}
		return find_in_loop_users(args[0], args[1], args[2])
	}), "licm-one-loop": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("licm-one-loop: wrong number of arguments %d (expected 2)", len(args))
		}
		return licm_one_loop(args[0], args[1])
	}), "operand-defined-outside?": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("operand-defined-outside?: wrong number of arguments %d (expected 3)", len(args))
		}
		return operand_defined_outside_QMARK_(args[0], args[1], args[2])
	}), "licm": __gogen_wrap(func(args []vm.Value) (vm.Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("licm: wrong number of arguments %d (expected 1)", len(args))
		}
		return licm(args[0])
	}),
	})
}
