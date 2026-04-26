// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package alloc

import (
	"context"
	"errors"
	"testing"
)

// fakeEvalStore is a scriptable EvalCtx implementation. It records the last
// invocation's script/key/args and returns a queue of pre-baked results so
// each test can control the sequence of script returns.
type fakeEvalStore struct {
	results []evalStep
	calls   []evalCall

	delErrs    []error
	delDeleted []int
	delCalls   []delCall
}

type delCall struct {
	keys []string
}

type evalStep struct {
	res any
	err error
}

type evalCall struct {
	script string
	key    string
	args   []any
}

func (f *fakeEvalStore) EvalCtx(_ context.Context, script, key string, args ...any) (any, error) {
	f.calls = append(f.calls, evalCall{script: script, key: key, args: append([]any(nil), args...)})
	if len(f.results) == 0 {
		return nil, errors.New("fakeEvalStore: no result scripted")
	}
	step := f.results[0]
	f.results = f.results[1:]
	return step.res, step.err
}

func (f *fakeEvalStore) DelCtx(_ context.Context, keys ...string) (int, error) {
	f.delCalls = append(f.delCalls, delCall{keys: append([]string(nil), keys...)})
	var deleted int
	if len(f.delDeleted) > 0 {
		deleted = f.delDeleted[0]
		f.delDeleted = f.delDeleted[1:]
	} else {
		deleted = len(keys)
	}
	if len(f.delErrs) > 0 {
		err := f.delErrs[0]
		f.delErrs = f.delErrs[1:]
		if err != nil {
			return 0, err
		}
	}
	return deleted, nil
}

// fixedClock returns a constant millis value, decoupling tests from wall time.
func fixedClock(t int64) func() int64 { return func() int64 { return t } }

func newCacheForTest(store evalStore) *xkvCache {
	c := newXKVCacheWithStore(store, WithNowMilli(fixedClock(1_700_000_000_000)))
	return c
}

// --- Malloc state parsing ---------------------------------------------------

func TestXKVMallocSuccessParsesArray(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{
		{res: []any{int64(0), int64(10), int64(60), []byte("1700000000000")}},
	}}
	c := newCacheForTest(store)

	got, err := c.Malloc(context.Background(), "k", 5)
	if err != nil {
		t.Fatalf("Malloc() err = %v", err)
	}
	want := MallocResult{State: MallocSuccess, CurrSeq: 10, LastSeq: 60, Mill: 1_700_000_000_000}
	if got != want {
		t.Fatalf("Malloc() = %+v, want %+v", got, want)
	}

	if len(store.calls) != 1 {
		t.Fatalf("EvalCtx invocations = %d, want 1", len(store.calls))
	}
	if store.calls[0].script != mallocScript {
		t.Fatal("EvalCtx script != mallocScript")
	}
	if store.calls[0].key != "k" {
		t.Fatalf("EvalCtx key = %q, want %q", store.calls[0].key, "k")
	}
	args := store.calls[0].args
	if len(args) != 5 {
		t.Fatalf("EvalCtx args arity = %d, want 5", len(args))
	}
	if args[0] != int64(5) {
		t.Fatalf("EvalCtx args[0] (size) = %#v, want int64(5)", args[0])
	}
	// args[1]=lockMillis, args[2]=dataSecond, args[3]=nowMillis, args[4]=owner
	if args[3] != int64(1_700_000_000_000) {
		t.Fatalf("EvalCtx args[3] (nowMillis) = %#v, want fixed clock", args[3])
	}
	if _, ok := args[4].(string); !ok {
		t.Fatalf("EvalCtx args[4] (owner) type = %T, want string", args[4])
	}
}

func TestXKVMallocMissParsesOwner(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{
		{res: []any{int64(1), "owner-1234", int64(1_700_000_000_000)}},
	}}
	c := newCacheForTest(store)

	got, err := c.Malloc(context.Background(), "k", 5)
	if err != nil {
		t.Fatalf("Malloc() err = %v", err)
	}
	if got.State != MallocMiss || got.Owner != "owner-1234" || got.Mill != 1_700_000_000_000 {
		t.Fatalf("Malloc() = %+v, want Miss owner=owner-1234", got)
	}
}

// TestXKVMallocColdPeekParsesEmptyOwner: the Lua script returns Miss with an
// empty owner string for a cold peek (size==0 on a missing key) so that no
// LOCK is left behind. The Go parser must accept that shape.
func TestXKVMallocColdPeekParsesEmptyOwner(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{
		{res: []any{int64(1), "", int64(1_700_000_000_000)}},
	}}
	c := newCacheForTest(store)

	got, err := c.Malloc(context.Background(), "k", 0)
	if err != nil {
		t.Fatalf("Malloc() err = %v", err)
	}
	if got.State != MallocMiss || got.Owner != "" {
		t.Fatalf("Malloc() = %+v, want Miss owner=\"\"", got)
	}
}

func TestXKVMallocLockedReturnsState(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{
		{res: []any{int64(2)}},
	}}
	c := newCacheForTest(store)

	got, err := c.Malloc(context.Background(), "k", 5)
	if err != nil {
		t.Fatalf("Malloc() err = %v", err)
	}
	if got.State != MallocLocked {
		t.Fatalf("Malloc() state = %v, want MallocLocked", got.State)
	}
}

func TestXKVMallocExceedParses5Tuple(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{
		{res: []any{int64(3), int64(43), int64(93), "owner-X", int64(1_700_000_000_000)}},
	}}
	c := newCacheForTest(store)

	got, err := c.Malloc(context.Background(), "k", 51)
	if err != nil {
		t.Fatalf("Malloc() err = %v", err)
	}
	want := MallocResult{
		State:   MallocExceed,
		CurrSeq: 43,
		LastSeq: 93,
		Owner:   "owner-X",
		Mill:    1_700_000_000_000,
	}
	if got != want {
		t.Fatalf("Malloc() = %+v, want %+v", got, want)
	}
}

func TestXKVMallocUnknownStateReturnsInvalidState(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{
		{res: []any{int64(99)}},
	}}
	c := newCacheForTest(store)

	_, err := c.Malloc(context.Background(), "k", 1)
	if !errors.Is(err, ErrInvalidState) {
		t.Fatalf("Malloc() err = %v, want ErrInvalidState", err)
	}
}

func TestXKVMallocPropagatesEvalError(t *testing.T) {
	wantErr := errors.New("redis down")
	store := &fakeEvalStore{results: []evalStep{{err: wantErr}}}
	c := newCacheForTest(store)

	_, err := c.Malloc(context.Background(), "k", 1)
	if !errors.Is(err, wantErr) {
		t.Fatalf("Malloc() err = %v, want wrapped %v", err, wantErr)
	}
}

// --- SetSeq state parsing ---------------------------------------------------

func TestXKVSetSeqSuccess(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{{res: int64(0)}}}
	c := newCacheForTest(store)

	state, err := c.SetSeq(context.Background(), "k", "owner-1", 100, 200, 1_700_000_000_000)
	if err != nil {
		t.Fatalf("SetSeq() err = %v", err)
	}
	if state != SetSeqSuccess {
		t.Fatalf("SetSeq() state = %v, want SetSeqSuccess", state)
	}
	args := store.calls[0].args
	if len(args) != 5 {
		t.Fatalf("EvalCtx args arity = %d, want 5", len(args))
	}
	if args[0] != "owner-1" {
		t.Fatalf("EvalCtx args[0] (owner) = %#v, want owner-1", args[0])
	}
	if args[2] != int64(100) || args[3] != int64(200) {
		t.Fatalf("EvalCtx args[2..3] (curr/last) = %v %v, want 100 200", args[2], args[3])
	}
}

func TestXKVSetSeqLockLost(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{{res: int64(1)}}}
	c := newCacheForTest(store)

	state, err := c.SetSeq(context.Background(), "k", "owner-1", 100, 200, 1)
	if err != nil {
		t.Fatalf("SetSeq() err = %v", err)
	}
	if state != SetSeqLockLost {
		t.Fatalf("SetSeq() state = %v, want SetSeqLockLost", state)
	}
}

func TestXKVSetSeqRejectsLastBelowCurr(t *testing.T) {
	c := newCacheForTest(&fakeEvalStore{})
	if _, err := c.SetSeq(context.Background(), "k", "owner-1", 200, 100, 0); err == nil {
		t.Fatal("SetSeq(last<curr) err = nil, want validation error")
	}
}

func TestXKVSetSeqRejectsEmptyOwner(t *testing.T) {
	c := newCacheForTest(&fakeEvalStore{})
	if _, err := c.SetSeq(context.Background(), "k", "", 100, 200, 0); err == nil {
		t.Fatal("SetSeq(owner=\"\") err = nil, want validation error")
	}
}

// --- Invalidate -------------------------------------------------------------

func TestXKVInvalidateDeletesKey(t *testing.T) {
	store := &fakeEvalStore{}
	c := newCacheForTest(store)

	if err := c.Invalidate(context.Background(), "k"); err != nil {
		t.Fatalf("Invalidate() err = %v", err)
	}
	if len(store.delCalls) != 1 {
		t.Fatalf("DelCtx invocations = %d, want 1", len(store.delCalls))
	}
	if got := store.delCalls[0].keys; len(got) != 1 || got[0] != "k" {
		t.Fatalf("DelCtx keys = %v, want [k]", got)
	}
}

func TestXKVInvalidatePropagatesError(t *testing.T) {
	wantErr := errors.New("redis down")
	store := &fakeEvalStore{delErrs: []error{wantErr}}
	c := newCacheForTest(store)

	err := c.Invalidate(context.Background(), "k")
	if !errors.Is(err, wantErr) {
		t.Fatalf("Invalidate() err = %v, want wrapped %v", err, wantErr)
	}
}

// --- Owner generation -------------------------------------------------------

func TestXKVNextOwnerIsMonotonic(t *testing.T) {
	c := newCacheForTest(&fakeEvalStore{})
	a := c.nextOwner()
	b := c.nextOwner()
	if a == b {
		t.Fatalf("nextOwner() returned duplicate %q", a)
	}
}

// --- Option wiring ----------------------------------------------------------

func TestXKVOptionsOverrideDefaults(t *testing.T) {
	store := &fakeEvalStore{results: []evalStep{
		{res: []any{int64(0), int64(1), int64(2), int64(3)}},
	}}
	c := newXKVCacheWithStore(
		store,
		WithLockTTL(123_000_000), // 123ms
		WithDataTTL(60_000_000_000),
		WithNowMilli(fixedClock(42)),
	)

	if _, err := c.Malloc(context.Background(), "k", 1); err != nil {
		t.Fatalf("Malloc() err = %v", err)
	}
	args := store.calls[0].args
	if args[1] != int64(123) { // 123ms in millis
		t.Fatalf("EvalCtx args[1] (lockMillis) = %v, want 123", args[1])
	}
	if args[2] != int64(60) { // 60s in seconds
		t.Fatalf("EvalCtx args[2] (dataSecond) = %v, want 60", args[2])
	}
	if args[3] != int64(42) {
		t.Fatalf("EvalCtx args[3] (nowMillis) = %v, want 42", args[3])
	}
}
