/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package rt

import (
	"testing"
	"time"

	"github.com/nooga/let-go/pkg/vm"
)

func asyncFn(t *testing.T, name string) vm.Fn {
	t.Helper()
	v := NS("async").Lookup(vm.Symbol(name))
	if v == nil {
		t.Fatalf("async/%s not found", name)
	}
	fn, ok := v.(*vm.Var).Deref().(vm.Fn)
	if !ok {
		t.Fatalf("async/%s is not an Fn", name)
	}
	return fn
}

func invoke(t *testing.T, fn vm.Fn, args ...vm.Value) vm.Value {
	t.Helper()
	v, err := fn.Invoke(args)
	if err != nil {
		t.Fatalf("invoke: %v", err)
	}
	return v
}

// drainChan reads all values from ch until it closes or timeout.
func drainChan(t *testing.T, ch vm.Chan, timeout time.Duration) []vm.Value {
	t.Helper()
	var out []vm.Value
	deadline := time.After(timeout)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return out
			}
			out = append(out, v)
		case <-deadline:
			t.Fatalf("drainChan timed out after %v (got %d values)", timeout, len(out))
			return out
		}
	}
}

// TestAsyncToChanPipe is a correctness regression guard: to-chan! emits a
// collection onto a channel and closes it; pipe forwards src→dst and
// closes dst when src closes.
func TestAsyncToChanPipe(t *testing.T) {
	toChan := asyncFn(t, "to-chan!")
	pipe := asyncFn(t, "pipe")

	src := invoke(t, toChan, vm.NewArrayVector([]vm.Value{vm.Int(1), vm.Int(2), vm.Int(3)})).(vm.Chan)
	dst := make(vm.Chan, 10)
	invoke(t, pipe, src, dst)

	got := drainChan(t, dst, 2*time.Second)
	if len(got) != 3 || got[0] != vm.Int(1) || got[2] != vm.Int(3) {
		t.Fatalf("expected [1 2 3] through pipe, got %v", got)
	}
}

// TestAsyncSplit is a correctness guard for split's routing + close.
func TestAsyncSplit(t *testing.T) {
	toChan := asyncFn(t, "to-chan!")
	split := asyncFn(t, "split")
	isEven, _ := vm.NativeFnType.Wrap(func(a []vm.Value) (vm.Value, error) {
		return vm.Boolean(int(a[0].(vm.Int))%2 == 0), nil
	})

	src := invoke(t, toChan, vm.NewArrayVector([]vm.Value{
		vm.Int(1), vm.Int(2), vm.Int(3), vm.Int(4)})).(vm.Chan)
	res := invoke(t, split, isEven, src)
	pair := res.(vm.ArrayVector)
	trueCh := pair.ValueAt(vm.Int(0)).(vm.Chan)
	falseCh := pair.ValueAt(vm.Int(1)).(vm.Chan)

	// split's channels are unbuffered and it sends interleaved, so read
	// both concurrently to avoid deadlock.
	type res2 struct{ evens, odds []vm.Value }
	ch := make(chan res2, 1)
	go func() {
		var r res2
		r.evens = drainChan(t, trueCh, 2*time.Second)
		ch <- r
	}()
	odds := drainChan(t, falseCh, 2*time.Second)
	evens := (<-ch).evens
	if len(evens) != 2 || len(odds) != 2 {
		t.Fatalf("expected 2 evens + 2 odds, got evens=%v odds=%v", evens, odds)
	}
}

// TestAsyncGoroutinesTrackedAndDrained pins the registry wiring: a
// channel pipeline that blocks forever (no reader on its output) must be
// (a) counted by the VM goroutine registry and (b) released by Drain via
// context cancellation — otherwise it leaks for the life of the process.
func TestAsyncGoroutinesTrackedAndDrained(t *testing.T) {
	// Drain anything left by other tests so the baseline is clean.
	vm.Goroutines.Drain(2 * time.Second)
	base := vm.Goroutines.Live()

	toChan := asyncFn(t, "to-chan!")
	pipe := asyncFn(t, "pipe")

	// Pipe into an UNBUFFERED dst that nobody reads → the pipe goroutine
	// blocks on `dst <- v` forever.
	src := invoke(t, toChan, vm.NewArrayVector([]vm.Value{vm.Int(1), vm.Int(2)})).(vm.Chan)
	dst := make(vm.Chan) // unbuffered, no reader
	invoke(t, pipe, src, dst)

	// Let the goroutines start and block.
	deadline := time.Now().Add(time.Second)
	for vm.Goroutines.Live() <= base && time.Now().Before(deadline) {
		time.Sleep(2 * time.Millisecond)
	}
	if vm.Goroutines.Live() <= base {
		t.Fatalf("expected async goroutines to be tracked (live=%d, base=%d)", vm.Goroutines.Live(), base)
	}

	if !vm.Goroutines.Drain(3 * time.Second) {
		t.Fatal("Drain did not release blocked async goroutines (not ctx-cancellable)")
	}
	if got := vm.Goroutines.Live(); got != 0 {
		t.Fatalf("expected 0 live after drain, got %d", got)
	}
}
