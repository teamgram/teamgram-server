// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package alloc

import (
	"context"
	"errors"
	"testing"
	"time"
)

// fakeCache is a scriptable cache that returns a queue of pre-baked
// MallocResults and SetSeqStates, recording every interaction for assertion.
type fakeCache struct {
	mallocResults   []mallocStep
	setSeqResults   []setSeqStep
	invalidateErrs  []error
	mallocCalls     []mallocCall
	setSeqCalls     []setSeqCall
	invalidatedKeys []string
}

type mallocStep struct {
	res MallocResult
	err error
}

type setSeqStep struct {
	state SetSeqState
	err   error
}

type mallocCall struct {
	key  string
	size int64
}

type setSeqCall struct {
	key     string
	owner   string
	currSeq int64
	lastSeq int64
	mill    int64
}

func (f *fakeCache) Malloc(_ context.Context, key string, size int64) (MallocResult, error) {
	f.mallocCalls = append(f.mallocCalls, mallocCall{key: key, size: size})
	if len(f.mallocResults) == 0 {
		return MallocResult{}, errors.New("fakeCache: no malloc result scripted")
	}
	step := f.mallocResults[0]
	f.mallocResults = f.mallocResults[1:]
	return step.res, step.err
}

func (f *fakeCache) SetSeq(_ context.Context, key, owner string, curr, last, mill int64) (SetSeqState, error) {
	f.setSeqCalls = append(f.setSeqCalls, setSeqCall{
		key:     key,
		owner:   owner,
		currSeq: curr,
		lastSeq: last,
		mill:    mill,
	})
	if len(f.setSeqResults) == 0 {
		return SetSeqSuccess, nil
	}
	step := f.setSeqResults[0]
	f.setSeqResults = f.setSeqResults[1:]
	return step.state, step.err
}

func (f *fakeCache) Invalidate(_ context.Context, key string) error {
	f.invalidatedKeys = append(f.invalidatedKeys, key)
	if len(f.invalidateErrs) == 0 {
		return nil
	}
	err := f.invalidateErrs[0]
	f.invalidateErrs = f.invalidateErrs[1:]
	return err
}

type fakeSeqStore struct {
	mallocCalls    []mallocCall
	mallocSeqs     []int64
	mallocErrs     []error
	maxSeq         int64
	maxSeqErr      error
	setMaxSeqCalls []setMaxSeqCall
	setMaxSeqErr   error
}

type setMaxSeqCall struct {
	key string
	seq int64
}

func (f *fakeSeqStore) Malloc(_ context.Context, key string, size int64) (int64, error) {
	f.mallocCalls = append(f.mallocCalls, mallocCall{key: key, size: size})
	if len(f.mallocErrs) > 0 {
		err := f.mallocErrs[0]
		f.mallocErrs = f.mallocErrs[1:]
		if err != nil {
			return 0, err
		}
	}
	if len(f.mallocSeqs) == 0 {
		return 0, errors.New("fakeSeqStore: no malloc seq scripted")
	}
	seq := f.mallocSeqs[0]
	f.mallocSeqs = f.mallocSeqs[1:]
	return seq, nil
}

func (f *fakeSeqStore) GetMaxSeq(_ context.Context, _ string) (int64, error) {
	return f.maxSeq, f.maxSeqErr
}

func (f *fakeSeqStore) SetMaxSeq(_ context.Context, key string, seq int64) error {
	f.setMaxSeqCalls = append(f.setMaxSeqCalls, setMaxSeqCall{key: key, seq: seq})
	return f.setMaxSeqErr
}

func newAllocator(cache *fakeCache, store *fakeSeqStore) *Allocator {
	return NewAllocator(cache, store, WithWait(time.Microsecond), WithRetries(3))
}

// TestMallocSuccessHitsCacheOnly: Success path should not touch the store.
func TestMallocSuccessHitsCacheOnly(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocSuccess, CurrSeq: 10, LastSeq: 60, Mill: 1000}},
		},
	}
	store := &fakeSeqStore{}
	a := newAllocator(cache, store)

	seq, mill, err := a.MallocTime(context.Background(), "inbox:1", 5)
	if err != nil {
		t.Fatalf("MallocTime() err = %v", err)
	}
	if seq != 10 || mill != 1000 {
		t.Fatalf("MallocTime() = (%d, %d), want (10, 1000)", seq, mill)
	}
	if len(store.mallocCalls) != 0 {
		t.Fatalf("store.Malloc invoked %d times, want 0", len(store.mallocCalls))
	}
	if got := a.Stats().CacheHit; got != 1 {
		t.Fatalf("Stats.CacheHit = %d, want 1", got)
	}
}

// TestMallocMissReplenishesAndCommits: Miss should call store.Malloc with
// blockSize headroom and commit (curr=seq+size, last=seq+mallocSize).
func TestMallocMissReplenishesAndCommits(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocMiss, Owner: "owner-A", Mill: 9_000}},
		},
	}
	store := &fakeSeqStore{mallocSeqs: []int64{40}}
	a := newAllocator(cache, store)

	seq, mill, err := a.MallocTime(context.Background(), "inbox:1", 3)
	if err != nil {
		t.Fatalf("MallocTime() err = %v", err)
	}
	if seq != 40 || mill != 9_000 {
		t.Fatalf("MallocTime() = (%d, %d), want (40, 9000)", seq, mill)
	}
	if got, want := store.mallocCalls[0].size, int64(53); got != want {
		t.Fatalf("store.Malloc size = %d, want %d (size+50 headroom)", got, want)
	}
	if len(cache.setSeqCalls) != 1 {
		t.Fatalf("cache.SetSeq calls = %d, want 1", len(cache.setSeqCalls))
	}
	c := cache.setSeqCalls[0]
	if c.owner != "owner-A" || c.currSeq != 43 || c.lastSeq != 93 || c.mill != 9_000 {
		t.Fatalf("SetSeq = %+v, want curr=43 last=93 owner=owner-A mill=9000", c)
	}
	st := a.Stats()
	if st.CacheMiss != 1 || st.StoreMalloc != 1 {
		t.Fatalf("Stats = %+v, want miss=1 store=1", st)
	}
}

// TestMallocLockedRetriesUntilSuccess: the Allocator should back off on
// Locked and try again; final Success returns the cached result.
func TestMallocLockedRetriesUntilSuccess(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocLocked}},
			{res: MallocResult{State: MallocLocked}},
			{res: MallocResult{State: MallocSuccess, CurrSeq: 100, LastSeq: 150, Mill: 7_000}},
		},
	}
	a := newAllocator(cache, &fakeSeqStore{})

	seq, _, err := a.MallocTime(context.Background(), "inbox:1", 1)
	if err != nil {
		t.Fatalf("MallocTime() err = %v", err)
	}
	if seq != 100 {
		t.Fatalf("MallocTime() seq = %d, want 100", seq)
	}
	if got := a.Stats().CacheLocked; got != 2 {
		t.Fatalf("Stats.CacheLocked = %d, want 2", got)
	}
}

// TestMallocLockedExhaustsRetries: when Locked persists past the retry budget
// the Allocator returns ErrLockTimeout.
func TestMallocLockedExhaustsRetries(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocLocked}},
			{res: MallocResult{State: MallocLocked}},
			{res: MallocResult{State: MallocLocked}},
		},
	}
	a := newAllocator(cache, &fakeSeqStore{})

	_, _, err := a.MallocTime(context.Background(), "inbox:1", 1)
	if !errors.Is(err, ErrLockTimeout) {
		t.Fatalf("MallocTime() err = %v, want ErrLockTimeout", err)
	}
}

// TestMallocExceedSpliceOK: when store-returned start equals the in-flight
// LAST, the caller's user-visible id is the original CURR; cache is committed
// with newCurr = origCurr + size.
func TestMallocExceedSpliceOK(t *testing.T) {
	// Cache state: tail [43, 93), tailUsed = 50.
	// Caller asks for 51 ids. freshNeeded = 1. blockSize(1) = 51.
	// store returns 93 (matches LastSeq) -> splice OK.
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{
				State:   MallocExceed,
				CurrSeq: 43,
				LastSeq: 93,
				Owner:   "owner-X",
				Mill:    5_000,
			}},
		},
	}
	store := &fakeSeqStore{mallocSeqs: []int64{93}}
	a := newAllocator(cache, store)

	seq, mill, err := a.MallocTime(context.Background(), "inbox:1", 51)
	if err != nil {
		t.Fatalf("MallocTime() err = %v", err)
	}
	if seq != 43 || mill != 5_000 {
		t.Fatalf("MallocTime() = (%d, %d), want (43, 5000)", seq, mill)
	}
	if got, want := store.mallocCalls[0].size, int64(51); got != want {
		t.Fatalf("store.Malloc size = %d, want %d (freshNeeded=1, +50 headroom)", got, want)
	}
	c := cache.setSeqCalls[0]
	wantCurr := int64(43 + 51) // origCurr + size
	wantLast := int64(93 + 51) // seq + mallocSize
	if c.currSeq != wantCurr || c.lastSeq != wantLast || c.owner != "owner-X" {
		t.Fatalf("SetSeq = %+v, want curr=%d last=%d owner=owner-X", c, wantCurr, wantLast)
	}
	if got := a.Stats().WastedSegments; got != 0 {
		t.Fatalf("Stats.WastedSegments = %d, want 0 on splice OK", got)
	}
}

// TestMallocExceedSpliceMismatchWastesTail: when store returns a seq that
// does not equal the in-flight LAST, the in-flight tail is treated as
// wasted; caller's id starts at the new seq.
func TestMallocExceedSpliceMismatchWastesTail(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{
				State:   MallocExceed,
				CurrSeq: 43,
				LastSeq: 93,
				Owner:   "owner-Y",
				Mill:    6_000,
			}},
		},
	}
	store := &fakeSeqStore{mallocSeqs: []int64{200}} // not == LastSeq
	a := newAllocator(cache, store)

	seq, _, err := a.MallocTime(context.Background(), "inbox:1", 51)
	if err != nil {
		t.Fatalf("MallocTime() err = %v", err)
	}
	if seq != 200 {
		t.Fatalf("MallocTime() seq = %d, want 200 (fresh start)", seq)
	}
	c := cache.setSeqCalls[0]
	if c.currSeq != 200+51 {
		t.Fatalf("SetSeq.currSeq = %d, want %d", c.currSeq, 200+51)
	}
	if got := a.Stats().WastedSegments; got != 1 {
		t.Fatalf("Stats.WastedSegments = %d, want 1", got)
	}
}

// TestMallocStoreErrorPropagates: store errors should bubble up and be
// counted; cache is not committed.
func TestMallocStoreErrorPropagates(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocMiss, Owner: "owner-A"}},
		},
	}
	storeErr := errors.New("db down")
	store := &fakeSeqStore{mallocErrs: []error{storeErr}}
	a := newAllocator(cache, store)

	_, _, err := a.MallocTime(context.Background(), "inbox:1", 1)
	if !errors.Is(err, storeErr) {
		t.Fatalf("MallocTime() err = %v, want %v", err, storeErr)
	}
	if len(cache.setSeqCalls) != 0 {
		t.Fatalf("cache.SetSeq invoked despite store error")
	}
	if got := a.Stats().StoreErr; got != 1 {
		t.Fatalf("Stats.StoreErr = %d, want 1", got)
	}
}

// TestMallocSetSeqLockLostMarksWaste: when the cache responds with
// SetSeqLockLost the caller's freshly fetched segment is wasted but the call
// itself still succeeds (the ids returned are still valid because the
// store's max_seq advance is durable).
func TestMallocSetSeqLockLostMarksWaste(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocMiss, Owner: "owner-A"}},
		},
		setSeqResults: []setSeqStep{{state: SetSeqLockLost}},
	}
	store := &fakeSeqStore{mallocSeqs: []int64{500}}
	a := newAllocator(cache, store)

	seq, _, err := a.MallocTime(context.Background(), "inbox:1", 3)
	if err != nil {
		t.Fatalf("MallocTime() err = %v", err)
	}
	if seq != 500 {
		t.Fatalf("MallocTime() seq = %d, want 500", seq)
	}
	st := a.Stats()
	if st.SetSeqLockLost != 1 || st.WastedSegments != 1 {
		t.Fatalf("Stats = %+v, want lockLost=1 wasted=1", st)
	}
}

// TestGetMaxSeqDoesNotPoisonCacheOnMiss: GetMaxSeq must fall through to the
// store on cache miss without committing an empty segment. The Lua script
// signals a lock-less peek-Miss with an empty Owner so the Allocator must
// tolerate that too.
func TestGetMaxSeqDoesNotPoisonCacheOnMiss(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocMiss, Owner: ""}},
		},
	}
	store := &fakeSeqStore{maxSeq: 12345}
	a := newAllocator(cache, store)

	got, err := a.GetMaxSeq(context.Background(), "inbox:1")
	if err != nil {
		t.Fatalf("GetMaxSeq() err = %v", err)
	}
	if got != 12345 {
		t.Fatalf("GetMaxSeq() = %d, want 12345", got)
	}
	if len(cache.setSeqCalls) != 0 {
		t.Fatalf("cache.SetSeq invoked %d times, must be 0", len(cache.setSeqCalls))
	}
	if len(store.mallocCalls) != 0 {
		t.Fatalf("store.Malloc invoked %d times, must be 0 (used GetMaxSeq read path)", len(store.mallocCalls))
	}
}

// TestGetMaxSeqHitsCacheWhenWarm: when Malloc(size=0) returns Success the
// cached CURR is reported directly without touching the store.
func TestGetMaxSeqHitsCacheWhenWarm(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocSuccess, CurrSeq: 777, LastSeq: 800, Mill: 1}},
		},
	}
	store := &fakeSeqStore{maxSeq: 9_999}
	a := newAllocator(cache, store)

	got, err := a.GetMaxSeq(context.Background(), "inbox:1")
	if err != nil {
		t.Fatalf("GetMaxSeq() err = %v", err)
	}
	if got != 777 {
		t.Fatalf("GetMaxSeq() = %d, want 777 (cache hit)", got)
	}
}

// TestRejectsNegativeSize: negative size is a programming error.
func TestRejectsNegativeSize(t *testing.T) {
	a := newAllocator(&fakeCache{}, &fakeSeqStore{})
	if _, err := a.Malloc(context.Background(), "inbox:1", -1); !errors.Is(err, ErrInvalidSize) {
		t.Fatalf("Malloc(-1) err = %v, want ErrInvalidSize", err)
	}
}

// TestNilCacheGoesDirectToStore: when no cache is configured the Allocator
// must still satisfy the contract by routing every call through the store.
func TestNilCacheGoesDirectToStore(t *testing.T) {
	store := &fakeSeqStore{mallocSeqs: []int64{42}, maxSeq: 99}
	a := NewAllocator(nil, store)

	seq, err := a.Malloc(context.Background(), "inbox:1", 5)
	if err != nil {
		t.Fatalf("Malloc() err = %v", err)
	}
	if seq != 42 {
		t.Fatalf("Malloc() seq = %d, want 42", seq)
	}

	got, err := a.GetMaxSeq(context.Background(), "inbox:1")
	if err != nil {
		t.Fatalf("GetMaxSeq() err = %v", err)
	}
	if got != 99 {
		t.Fatalf("GetMaxSeq() = %d, want 99", got)
	}
}

// TestSetMaxSeqInvalidatesCache: after a successful store write the cached
// segment must be dropped so subsequent Mallocs cannot serve ids below the
// new max_seq.
func TestSetMaxSeqInvalidatesCache(t *testing.T) {
	cache := &fakeCache{}
	store := &fakeSeqStore{}
	a := newAllocator(cache, store)

	if err := a.SetMaxSeq(context.Background(), "inbox:1", 999); err != nil {
		t.Fatalf("SetMaxSeq() err = %v", err)
	}
	if len(store.setMaxSeqCalls) != 1 || store.setMaxSeqCalls[0].seq != 999 {
		t.Fatalf("store.SetMaxSeq calls = %+v, want one call with seq=999", store.setMaxSeqCalls)
	}
	wantKey := cacheKey("inbox:1")
	if len(cache.invalidatedKeys) != 1 || cache.invalidatedKeys[0] != wantKey {
		t.Fatalf("cache.Invalidate keys = %+v, want [%s]", cache.invalidatedKeys, wantKey)
	}
}

// TestSetMaxSeqStoreErrorSkipsInvalidate: when the store write fails the
// cache must be left untouched (the prior segment is still authoritative).
func TestSetMaxSeqStoreErrorSkipsInvalidate(t *testing.T) {
	storeErr := errors.New("db down")
	cache := &fakeCache{}
	store := &fakeSeqStore{setMaxSeqErr: storeErr}
	a := newAllocator(cache, store)

	if err := a.SetMaxSeq(context.Background(), "inbox:1", 999); !errors.Is(err, storeErr) {
		t.Fatalf("SetMaxSeq() err = %v, want %v", err, storeErr)
	}
	if len(cache.invalidatedKeys) != 0 {
		t.Fatalf("cache.Invalidate invoked despite store error: %+v", cache.invalidatedKeys)
	}
}

// TestSetMaxSeqInvalidateFailureSurfacesError: Invalidate failure must be
// surfaced to the caller — returning nil here would silently allow the
// cached segment to keep serving ids below the new max_seq, which is the
// exact correctness gap SetMaxSeq is supposed to close. The store write
// has already happened so the caller can simply retry (both store and
// cache operations are idempotent on the same seq).
func TestSetMaxSeqInvalidateFailureSurfacesError(t *testing.T) {
	wantErr := errors.New("redis down")
	cache := &fakeCache{invalidateErrs: []error{wantErr}}
	store := &fakeSeqStore{}
	a := newAllocator(cache, store)

	err := a.SetMaxSeq(context.Background(), "inbox:1", 42)
	if !errors.Is(err, wantErr) {
		t.Fatalf("SetMaxSeq() err = %v, want wrapped %v", err, wantErr)
	}
	if len(store.setMaxSeqCalls) != 1 || store.setMaxSeqCalls[0].seq != 42 {
		t.Fatalf("store.SetMaxSeq calls = %+v, want one call with seq=42", store.setMaxSeqCalls)
	}
	if len(cache.invalidatedKeys) != 1 {
		t.Fatalf("cache.Invalidate calls = %d, want 1", len(cache.invalidatedKeys))
	}
}

// TestSetMaxSeqInvalidateRetryIsIdempotent: after a transient invalidation
// failure the caller is expected to retry; the second attempt must succeed
// without double-counting or rejecting the unchanged seq value.
func TestSetMaxSeqInvalidateRetryIsIdempotent(t *testing.T) {
	cache := &fakeCache{invalidateErrs: []error{errors.New("redis down"), nil}}
	store := &fakeSeqStore{}
	a := newAllocator(cache, store)

	if err := a.SetMaxSeq(context.Background(), "inbox:1", 42); err == nil {
		t.Fatal("first SetMaxSeq() err = nil, want error")
	}
	if err := a.SetMaxSeq(context.Background(), "inbox:1", 42); err != nil {
		t.Fatalf("retry SetMaxSeq() err = %v, want nil", err)
	}
	if len(store.setMaxSeqCalls) != 2 {
		t.Fatalf("store.SetMaxSeq calls = %d, want 2 (idempotent retry)", len(store.setMaxSeqCalls))
	}
	if len(cache.invalidatedKeys) != 2 {
		t.Fatalf("cache.Invalidate calls = %d, want 2", len(cache.invalidatedKeys))
	}
}

// TestSetMaxSeqWithoutCacheNoOpInvalidate: in DB-direct mode there is no
// cache to invalidate; SetMaxSeq must not panic.
func TestSetMaxSeqWithoutCacheNoOpInvalidate(t *testing.T) {
	store := &fakeSeqStore{}
	a := NewAllocator(nil, store)

	if err := a.SetMaxSeq(context.Background(), "inbox:1", 42); err != nil {
		t.Fatalf("SetMaxSeq() err = %v", err)
	}
	if len(store.setMaxSeqCalls) != 1 {
		t.Fatalf("store.SetMaxSeq calls = %d, want 1", len(store.setMaxSeqCalls))
	}
}

// TestMallocCancelledContextStopsRetries: a cancelled context during the
// back-off should propagate immediately.
func TestMallocCancelledContextStopsRetries(t *testing.T) {
	cache := &fakeCache{
		mallocResults: []mallocStep{
			{res: MallocResult{State: MallocLocked}},
		},
	}
	a := NewAllocator(cache, &fakeSeqStore{}, WithWait(time.Hour), WithRetries(5))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, _, err := a.MallocTime(ctx, "inbox:1", 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("MallocTime() err = %v, want context.Canceled", err)
	}
}
