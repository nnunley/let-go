/*
 * Copyright (c) 2026 Norman Nunley, Jr <nnunley@gmail.com>
 * Part of the let-go project; see CONTRIBUTORS for full list of authors.
 * SPDX-License-Identifier: MIT
 */

package vm

import (
	"sync"
	"testing"
)

func TestGoIDIsNonZeroAndDistinct(t *testing.T) {
	a := goID()
	if a == 0 {
		t.Fatal("goID returned 0 for the running goroutine")
	}
	var b int64
	done := make(chan struct{})
	go func() { b = goID(); close(done) }()
	<-done
	if b == 0 || a == b {
		t.Fatalf("expected distinct non-zero goids, got a=%d b=%d", a, b)
	}
}

func TestCurrentScopeDefaultsToRootWhenUnscoped(t *testing.T) {
	if scopedLive.Load() != 0 {
		t.Skipf("scopedLive not clean: %d", scopedLive.Load())
	}
	if CurrentScope() != Goroutines {
		t.Fatal("unscoped goroutine should see the root scope")
	}
	if CurrentContext() != Goroutines.Context() {
		t.Fatal("unscoped CurrentContext should be the root context")
	}
}

func TestBindRestoreNesting(t *testing.T) {
	c1 := Goroutines.Child()
	c2 := Goroutines.Child()
	r1 := c1.bind()
	if CurrentScope() != c1 {
		t.Fatal("after binding c1, current should be c1")
	}
	r2 := c2.bind()
	if CurrentScope() != c2 {
		t.Fatal("after binding c2, current should be c2")
	}
	r2()
	if CurrentScope() != c1 {
		t.Fatal("after restoring c2, current should be c1 again")
	}
	r1()
	if CurrentScope() != Goroutines {
		t.Fatal("after restoring c1, current should be root")
	}
	if scopedLive.Load() != 0 {
		t.Fatalf("scopedLive should return to 0, got %d", scopedLive.Load())
	}
}

func TestConcurrentGoIDNoRace(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); _ = goID() }()
	}
	wg.Wait()
}
