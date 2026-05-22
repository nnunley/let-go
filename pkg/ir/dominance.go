/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package ir

// Dominators computes immediate dominators for each block in f.
//
// Block X dominates Block Y if every path from Entry to Y goes through
// X. The immediate dominator idom(Y) is the closest dominator other
// than Y itself. The entry block's idom is -1 (it has no dominator).
//
// Uses the iterative dataflow algorithm of Cooper, Harvey, and Kennedy
// ("A Simple, Fast Dominance Algorithm"). For typical function sizes
// (< 50 blocks) this runs in microseconds and is near-optimal.
//
// Returns idom[bid] for each block. For unreachable blocks, idom is -1
// (they aren't part of the dominator tree).
func (f *Function) Dominators() []BlockID {
	n := len(f.Blocks)
	idom := make([]BlockID, n)
	for i := range idom {
		idom[i] = -1
	}
	idom[f.Entry] = f.Entry

	// Reverse postorder of reachable blocks. Iterating in RPO maximizes
	// the chance that each block's predecessors have been processed when
	// we visit it, so the algorithm converges in few passes.
	rpo := f.reversePostorder()

	// Position-in-RPO lookup: for the intersect step we need to compare
	// "closeness to entry". RPO index does that — lower index = closer
	// to entry. We'll work with rpoIdx[bid] = position in rpo.
	rpoIdx := make([]int, n)
	for i := range rpoIdx {
		rpoIdx[i] = -1
	}
	for i, b := range rpo {
		rpoIdx[b] = i
	}

	changed := true
	for changed {
		changed = false
		// Skip entry (its idom is already set); visit others in RPO.
		for _, b := range rpo {
			if b == f.Entry {
				continue
			}
			// New idom is the intersection of all already-processed preds.
			var newIdom BlockID = -1
			for _, p := range f.Blocks[b].Preds {
				if idom[p] == -1 {
					continue // pred not yet processed
				}
				if newIdom == -1 {
					newIdom = p
				} else {
					newIdom = intersect(idom, rpoIdx, p, newIdom)
				}
			}
			if newIdom == -1 {
				// No reachable predecessor — unreachable block.
				continue
			}
			if idom[b] != newIdom {
				idom[b] = newIdom
				changed = true
			}
		}
	}

	// Restore entry's idom to -1 (it had idom[entry]=entry during the
	// algorithm as a sentinel, but consumers expect -1).
	idom[f.Entry] = -1
	return idom
}

// intersect walks up the dominator tree from b1 and b2 until they meet.
// rpoIdx provides ordering — the finger that's "deeper" (higher RPO
// index) moves up first.
func intersect(idom []BlockID, rpoIdx []int, b1, b2 BlockID) BlockID {
	for b1 != b2 {
		for rpoIdx[b1] > rpoIdx[b2] {
			b1 = idom[b1]
		}
		for rpoIdx[b2] > rpoIdx[b1] {
			b2 = idom[b2]
		}
	}
	return b1
}

// reversePostorder returns the reachable blocks in reverse postorder.
// Entry comes first, then blocks in an order such that each block
// appears AFTER one of its predecessors (for acyclic CFGs; with loops
// the back-edge predecessor obviously can't satisfy this — that's why
// the algorithm iterates).
func (f *Function) reversePostorder() []BlockID {
	n := len(f.Blocks)
	visited := make([]bool, n)
	var post []BlockID

	// Iterative DFS to avoid Go stack overflow on very deep CFGs.
	type frame struct {
		bid    BlockID
		succs  []BlockID
		nextSi int
	}
	stack := []frame{{bid: f.Entry, succs: f.Successors(f.Entry)}}
	visited[f.Entry] = true

	for len(stack) > 0 {
		top := &stack[len(stack)-1]
		if top.nextSi < len(top.succs) {
			s := top.succs[top.nextSi]
			top.nextSi++
			if int(s) < 0 || int(s) >= n {
				continue
			}
			if !visited[s] {
				visited[s] = true
				stack = append(stack, frame{bid: s, succs: f.Successors(s)})
			}
			continue
		}
		post = append(post, top.bid)
		stack = stack[:len(stack)-1]
	}

	// Reverse postorder.
	for i, j := 0, len(post)-1; i < j; i, j = i+1, j-1 {
		post[i], post[j] = post[j], post[i]
	}
	return post
}

// Successors returns the successor BlockIDs of bid, derived from the
// block's terminator. Order: for BranchIf, true-target first, then
// false. For unconditional Branch, single target. For Return/TailCall,
// empty (no in-function successors).
//
// Returns nil if the block has no terminator (shouldn't happen post-Build).
func (f *Function) Successors(bid BlockID) []BlockID {
	if int(bid) < 0 || int(bid) >= len(f.Blocks) {
		return nil
	}
	blk := &f.Blocks[bid]
	if blk.Term == 0 && bid != f.Entry {
		return nil
	}
	if blk.Term == 0 {
		return nil
	}
	term := &f.Insts[blk.Term]
	switch t := term.Aux.(type) {
	case *BranchTarget:
		if t == nil {
			return nil
		}
		return []BlockID{t.Target}
	case *CondTarget:
		if t == nil {
			return nil
		}
		var out []BlockID
		if t.True != nil {
			out = append(out, t.True.Target)
		}
		if t.False != nil {
			out = append(out, t.False.Target)
		}
		return out
	}
	return nil
}

// Dominates reports whether `a` dominates `b` (a == b is also true).
// idom is the array returned by f.Dominators().
//
// Walks up b's dominator chain until we hit a or fall off (idom==-1).
func Dominates(idom []BlockID, a, b BlockID) bool {
	for b != -1 {
		if a == b {
			return true
		}
		b = idom[b]
	}
	return false
}
