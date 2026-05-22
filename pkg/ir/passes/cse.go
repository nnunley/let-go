/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package passes

import (
	"fmt"
	"strings"

	"github.com/nooga/let-go/pkg/ir"
	"github.com/nooga/let-go/pkg/vm"
)

// CSE — common subexpression elimination.
//
// Block-local for the spike: walks each block's body and merges
// duplicate pure expressions. A node is a duplicate of an earlier
// node if both have:
//   - The same Op
//   - The same Refs (after redirecting via previously-CSE'd nodes)
//   - The same Aux (for Const, that's the vm.Value equality)
//
// When a duplicate is detected:
//   - The duplicate node is marked dead (Op = OpInvalid)
//   - A redirect entry maps the dead node's ID to the original's ID
//
// After the per-block walk, every Refs[] in the function is rewritten
// to follow redirects. This canonicalizes the IR.
//
// Cross-block CSE would need dominance — defer to a later pass.
func CSE(f *ir.Function) (changed bool) {
	redirect := make(map[ir.InstId]ir.InstId, len(f.Insts))

	// Flow analysis: which vars are mutated anywhere in this function?
	// LoadVar nodes targeting stable (non-mutated) vars can be merged
	// despite OpLoadVar's defensive `pure=false` flag.
	mutatedVars, allMutated := computeMutatedVars(f)

	for bid := range f.Blocks {
		blk := &f.Blocks[bid]

		// key → first InstId that computed this value (in this block).
		// Key is a string built from Op + redirected Refs + Aux.
		seen := map[string]ir.InstId{}

		kept := blk.Insts[:0]
		for _, nid := range blk.Insts {
			n := f.Inst(nid)
			if !n.Op.IsPure() && !cseSafeImpureLoadVar(n, mutatedVars, allMutated) {
				kept = append(kept, nid)
				continue
			}
			// Build the canonical key. Follow redirects on operand refs
			// so two nodes with operands that became equivalent compare equal.
			key := nodeKey(n, redirect)
			if prev, ok := seen[key]; ok {
				// Duplicate. Inherit its spans into the survivor before marking dead.
				f.Insts[prev].SourceInfos = ir.MergeSourceInfo(f.Insts[prev].SourceInfos, f.Insts[nid].SourceInfos...)
				redirect[nid] = prev
				f.Insts[nid].Op = ir.OpInvalid
				changed = true
				continue
			}
			seen[key] = nid
			kept = append(kept, nid)
		}
		blk.Insts = kept
	}

	if !changed {
		return false
	}

	// Final pass: rewrite all Refs to follow redirects. Includes
	// terminator Refs and branch-target Args.
	for i := range f.Insts {
		n := &f.Insts[i]
		for j, r := range n.Refs {
			if final, ok := follow(redirect, r); ok {
				n.Refs[j] = final
			}
		}
		switch t := n.Aux.(type) {
		case *ir.BranchTarget:
			if t != nil {
				rewriteArgs(t.Args, redirect)
			}
		case *ir.CondTarget:
			if t != nil {
				if t.True != nil {
					rewriteArgs(t.True.Args, redirect)
				}
				if t.False != nil {
					rewriteArgs(t.False.Args, redirect)
				}
			}
		}
	}
	return true
}

// nodeKey computes a canonical key for n. Two nodes with equal keys
// compute the same value. Refs are normalized via the redirect table.
func nodeKey(n *ir.Inst, redirect map[ir.InstId]ir.InstId) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d|", n.Op)
	for _, r := range n.Refs {
		final, _ := follow(redirect, r)
		fmt.Fprintf(&sb, "%d,", final)
	}
	sb.WriteByte('|')
	// Aux serialization. For Const, hash by the vm.Value's string repr;
	// for LoadArg, by arg index; for LoadVar, by var identity.
	switch v := n.Aux.(type) {
	case vm.Value:
		// Two Const nodes are equal iff their boxed values' String() match.
		// For Int/Float/Bool/String/Keyword this is reliable; for compound
		// values, may over-collapse if values stringify identically but
		// aren't `=`. Conservative refinement (compare actual values via
		// vm.ValueEquals) deferred.
		sb.WriteString(v.String())
	case int:
		fmt.Fprintf(&sb, "%d", v)
	default:
		// For terminators and other ops, Aux often references mutable
		// state (BranchTarget pointers). Don't hash it — terminators
		// aren't CSE candidates anyway (impure).
		fmt.Fprintf(&sb, "%p", v)
	}
	return sb.String()
}

// computeMutatedVars scans the function for nodes that may mutate a
// var's root, returning a set of var-identity pointers.
//
// A var is potentially mutated by:
//
//   - OpSetVar n  — direct (set! var ...). The Aux holds *vm.Var.
//   - OpCall n where the callee is a known-mutating builtin
//     (alter-var-root!, with-redefs unwinds, etc.). We use the var's
//     name to identify these; if you call alter-var-root! on a var
//     whose identity we don't statically know, we conservatively
//     mark ALL vars as mutated.
//
// Returns nil if no mutation found (caller treats every var as stable).
// Returns a non-nil "all mutated" sentinel if we can't be precise.
//
// CSE uses this to decide whether to merge LoadVar nodes: a LoadVar
// whose target var is stable (not in the mutated set) can be CSE'd
// across the function without risk of a stale read.
func computeMutatedVars(f *ir.Function) (mutated map[*vm.Var]bool, allMutated bool) {
	mutated = make(map[*vm.Var]bool)
	for i := range f.Insts {
		n := &f.Insts[i]
		switch n.Op {
		case ir.OpSetVar:
			// Aux holds the target Var (set by Build).
			if vv, ok := n.Aux.(vm.Value); ok {
				if v, ok := vv.(*vm.Var); ok {
					mutated[v] = true
					continue
				}
			}
			// We saw a SetVar but couldn't identify its target.
			// Conservative: mark all vars as mutated.
			return nil, true
		case ir.OpCall:
			// Check if the callee is a known-mutating builtin.
			if len(n.Refs) == 0 {
				continue
			}
			fnNode := &f.Insts[n.Refs[0]]
			if fnNode.Op != ir.OpLoadVar {
				continue
			}
			fv, _ := fnNode.Aux.(vm.Value)
			fnVar, _ := fv.(*vm.Var)
			if fnVar == nil {
				continue
			}
			if knownMutatingBuiltins[fnVar.VarName()] {
				// We can't statically know which var this mutates
				// without deeper analysis (the var argument is at runtime).
				// Conservative: mark all vars as mutated.
				return nil, true
			}
		}
	}
	return mutated, false
}

// knownMutatingBuiltins lists core builtins that mutate var roots.
// When we see a call to one of these, we can't precisely know which
// var(s) it mutates from the static IR — so we conservatively mark
// all vars as potentially mutated.
//
// (A more precise analysis would examine the call's args: if the var
// arg is a constant *vm.Var, mark only that one. Deferred.)
var knownMutatingBuiltins = map[string]bool{
	"alter-var-root":  true,
	"alter-var-root!": true,
	"intern":          true,
	"with-redefs":     true,
}

// cseSafeImpureLoadVar reports whether n is a LoadVar whose target var
// is provably not mutated in this function. CSE can merge such LoadVars
// across the function — they always return the same value within a
// single invocation.
//
// The IR's general purity flag for LoadVar is `false` (defensive
// default), but a var that's never mutated anywhere is effectively
// constant from this function's perspective.
//
// Caller pre-computes the mutated set via computeMutatedVars.
func cseSafeImpureLoadVar(n *ir.Inst, mutated map[*vm.Var]bool, allMutated bool) bool {
	if allMutated {
		return false
	}
	if n.Op != ir.OpLoadVar {
		return false
	}
	v, ok := n.Aux.(vm.Value)
	if !ok {
		return false
	}
	vv, ok := v.(*vm.Var)
	if !ok {
		return false
	}
	return !mutated[vv]
}

// follow walks the redirect chain to its terminus. Returns the final
// InstId and whether any redirect was followed.
func follow(redirect map[ir.InstId]ir.InstId, id ir.InstId) (ir.InstId, bool) {
	final := id
	moved := false
	for {
		next, ok := redirect[final]
		if !ok {
			return final, moved
		}
		final = next
		moved = true
	}
}

func rewriteArgs(args []ir.InstId, redirect map[ir.InstId]ir.InstId) {
	for i, a := range args {
		if final, ok := follow(redirect, a); ok {
			args[i] = final
		}
	}
}
