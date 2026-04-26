// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Package alloc implements a per-key monotonic sequence allocator backed by an
// authoritative SeqStore (e.g. MySQL) and an optional Cache (e.g. Redis) that
// caches a pre-allocated segment in front of the store.
//
// The allocator guarantees that allocated ids are strictly increasing within a
// single key. It does NOT guarantee gapless ids: when the cache holding a
// pre-allocated segment is lost (eviction, restart, owner change), the
// remainder of that segment is wasted and the next allocation continues after
// the next store-backed range. The producer side is therefore expected to
// tolerate gaps in allocated ids; callers that need gapless ids should commit
// the consumed id together with the value it identifies (so an unused id is
// never durably observed).
package alloc

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// MallocState is the protocol-level result of a Cache.Malloc call.
type MallocState int

// MallocState values.
const (
	// MallocSuccess indicates the cache served the request entirely from the
	// in-memory segment. The caller owns [CurrSeq, CurrSeq+size).
	MallocSuccess MallocState = 0

	// MallocMiss indicates the cache had no segment for the key. The caller
	// holds the lock identified by Owner and must replenish from the store
	// then call SetSeq with that Owner.
	MallocMiss MallocState = 1

	// MallocLocked indicates another caller is currently replenishing the
	// segment. The caller should back off and retry.
	MallocLocked MallocState = 2

	// MallocExceed indicates the segment is partially exhausted: the caller
	// has been given the remainder [CurrSeq, LastSeq) and must replenish
	// from the store then call SetSeq with Owner.
	MallocExceed MallocState = 3
)

// SetSeqState is the protocol-level result of a Cache.SetSeq call.
type SetSeqState int

// SetSeqState values.
const (
	// SetSeqSuccess indicates the new segment was committed and the lock
	// released.
	SetSeqSuccess SetSeqState = 0

	// SetSeqLockLost indicates the lock identified by the caller's Owner is
	// no longer held (lock missing, taken over by someone else, or the cache
	// key was evicted). The caller's freshly fetched segment is wasted.
	SetSeqLockLost SetSeqState = 1
)

// MallocResult carries the outcome of a Cache.Malloc call.
type MallocResult struct {
	State   MallocState
	CurrSeq int64
	LastSeq int64
	Owner   string
	Mill    int64
}

// Cache is the read-through cache interface in front of SeqStore. The Lua
// script semantics are documented in scripts.go.
type Cache interface {
	// Malloc reserves up to `size` ids and returns a MallocResult describing
	// the outcome. size==0 means "peek the current cursor without consuming
	// the segment".
	Malloc(ctx context.Context, key string, size int64) (MallocResult, error)

	// SetSeq commits a freshly fetched segment [currSeq, lastSeq) under the
	// lock identified by owner.
	SetSeq(ctx context.Context, key, owner string, currSeq, lastSeq, mill int64) (SetSeqState, error)
}

// SeqStore is the authoritative store for max_seq.
type SeqStore interface {
	// Malloc atomically advances the persisted max_seq by `size` and returns
	// the start of the newly reserved range. size==0 means "read max_seq".
	Malloc(ctx context.Context, key string, size int64) (int64, error)

	// GetMaxSeq returns the current persisted max_seq.
	GetMaxSeq(ctx context.Context, key string) (int64, error)

	// SetMaxSeq forces max_seq to a specific value but never lets it go
	// backwards. Implementations must reject regressions.
	SetMaxSeq(ctx context.Context, key string, seq int64) error
}

// Errors returned by Allocator.
var (
	// ErrInvalidSize is returned when size is negative.
	ErrInvalidSize = errors.New("alloc: size must be greater than or equal to 0")

	// ErrLockTimeout is returned when the cache is locked and the configured
	// retry budget is exhausted.
	ErrLockTimeout = errors.New("alloc: waiting for cache lock timeout")

	// ErrInvalidState is returned when the cache layer returns an
	// inconsistent state machine result (treat as a bug or protocol drift).
	ErrInvalidState = errors.New("alloc: invalid cache state")
)

// Stats is a snapshot of Allocator counters. Counter values are cumulative
// since the Allocator was constructed.
type Stats struct {
	CacheHit       int64 // MallocSuccess on the cache fast path
	CacheMiss      int64 // MallocMiss: cold key, replenished from store
	CacheLocked    int64 // MallocLocked: backed off and retried
	CacheExceed    int64 // MallocExceed: segment partially exhausted
	StoreMalloc    int64 // store.Malloc invocations (excluding GetMaxSeq peeks)
	StoreErr       int64 // store.Malloc errors
	SetSeqLockLost int64 // setSeq saw the lock taken over by another owner
	WastedSegments int64 // segments dropped due to splice mismatch / commit failure / lock lost
}

// Allocator is the per-key monotonic sequence allocator.
type Allocator struct {
	cache     Cache
	store     SeqStore
	wait      time.Duration
	retries   int
	blockSize func(key string, size int64) int64

	cacheHit       atomic.Int64
	cacheMiss      atomic.Int64
	cacheLocked    atomic.Int64
	cacheExceed    atomic.Int64
	storeMalloc    atomic.Int64
	storeErr       atomic.Int64
	setSeqLockLost atomic.Int64
	wastedSegments atomic.Int64
}

// Option configures an Allocator.
type Option func(*Allocator)

// NewAllocator creates an Allocator. cache may be nil, in which case every
// call goes straight to store (suitable for low-QPS keys or local testing).
func NewAllocator(cache Cache, store SeqStore, opts ...Option) *Allocator {
	a := &Allocator{
		cache:     cache,
		store:     store,
		wait:      250 * time.Millisecond,
		retries:   10,
		blockSize: defaultBlockSize,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// WithWait sets the back-off duration between retries on MallocLocked.
func WithWait(wait time.Duration) Option {
	return func(a *Allocator) {
		if wait > 0 {
			a.wait = wait
		}
	}
}

// WithRetries sets the maximum number of MallocLocked retries before giving
// up with ErrLockTimeout.
func WithRetries(retries int) Option {
	return func(a *Allocator) {
		if retries > 0 {
			a.retries = retries
		}
	}
}

// WithBlockSize installs a custom strategy for computing the segment size to
// fetch from the store. The default keeps a 50-id headroom on top of the
// caller's request size; adaptive strategies (e.g. proportional to recent
// consumption rate) can be plugged in here.
func WithBlockSize(fn func(key string, size int64) int64) Option {
	return func(a *Allocator) {
		if fn != nil {
			a.blockSize = fn
		}
	}
}

func defaultBlockSize(_ string, size int64) int64 {
	if size == 0 {
		return 0
	}
	return size + 50
}

func cacheKey(key string) string {
	return fmt.Sprintf("idgen:malloc_seq:%s", key)
}

// Stats returns a snapshot of Allocator counters. Useful for exposing under
// /metrics or periodic logging.
func (a *Allocator) Stats() Stats {
	return Stats{
		CacheHit:       a.cacheHit.Load(),
		CacheMiss:      a.cacheMiss.Load(),
		CacheLocked:    a.cacheLocked.Load(),
		CacheExceed:    a.cacheExceed.Load(),
		StoreMalloc:    a.storeMalloc.Load(),
		StoreErr:       a.storeErr.Load(),
		SetSeqLockLost: a.setSeqLockLost.Load(),
		WastedSegments: a.wastedSegments.Load(),
	}
}

// Malloc reserves `size` ids and returns the first one allocated.
func (a *Allocator) Malloc(ctx context.Context, key string, size int64) (int64, error) {
	seq, _, err := a.MallocTime(ctx, key, size)
	return seq, err
}

// GetMaxSeq returns the current next-to-be-allocated id without consuming it.
//
// When backed by a Cache, GetMaxSeq does NOT poison the cache state on miss:
// it falls through to the store directly so a peek stays a true read-only
// operation.
func (a *Allocator) GetMaxSeq(ctx context.Context, key string) (int64, error) {
	if a.cache == nil {
		return a.store.GetMaxSeq(ctx, key)
	}
	res, err := a.cache.Malloc(ctx, cacheKey(key), 0)
	if err != nil {
		return 0, err
	}
	switch res.State {
	case MallocSuccess:
		a.cacheHit.Add(1)
		return res.CurrSeq, nil
	case MallocMiss, MallocLocked, MallocExceed:
		// Read-through to the store without writing back to the cache.
		return a.store.GetMaxSeq(ctx, key)
	default:
		return 0, fmt.Errorf("%w: %d", ErrInvalidState, res.State)
	}
}

// SetMaxSeq forces the persisted max_seq to a value. The store must reject
// regressions; see SeqStore.SetMaxSeq.
func (a *Allocator) SetMaxSeq(ctx context.Context, key string, seq int64) error {
	return a.store.SetMaxSeq(ctx, key, seq)
}

// MallocTime reserves `size` ids and returns the first allocated id together
// with the cache's monotonic millis clock at allocation time. The clock is
// useful for relating the id with surrounding events (e.g. message_date in
// IM scenarios).
//
// When size==0 MallocTime degenerates into a peek (see GetMaxSeq) and returns
// (currMaxSeq, 0, nil) without consuming any id.
func (a *Allocator) MallocTime(ctx context.Context, key string, size int64) (int64, int64, error) {
	if size < 0 {
		return 0, 0, ErrInvalidSize
	}
	if a.cache == nil {
		a.storeMalloc.Add(1)
		seq, err := a.store.Malloc(ctx, key, size)
		if err != nil {
			a.storeErr.Add(1)
			return 0, 0, err
		}
		return seq, 0, nil
	}

	if size == 0 {
		seq, err := a.GetMaxSeq(ctx, key)
		return seq, 0, err
	}

	ck := cacheKey(key)
	for i := 0; i < a.retries; i++ {
		res, err := a.cache.Malloc(ctx, ck, size)
		if err != nil {
			return 0, 0, err
		}

		switch res.State {
		case MallocSuccess:
			a.cacheHit.Add(1)
			return res.CurrSeq, res.Mill, nil

		case MallocMiss:
			a.cacheMiss.Add(1)
			seq, err := a.handleMiss(ctx, ck, key, size, res.Owner, res.Mill)
			return seq, res.Mill, err

		case MallocLocked:
			a.cacheLocked.Add(1)
			if err := a.waitLock(ctx); err != nil {
				return 0, 0, err
			}

		case MallocExceed:
			a.cacheExceed.Add(1)
			seq, err := a.handleExceed(ctx, ck, key, size, res)
			return seq, res.Mill, err

		default:
			return 0, 0, fmt.Errorf("%w: %d", ErrInvalidState, res.State)
		}
	}

	return 0, 0, ErrLockTimeout
}

// handleMiss replenishes a cold key. Caller's range is [seq, seq+size).
func (a *Allocator) handleMiss(
	ctx context.Context,
	ck, key string,
	size int64,
	owner string,
	mill int64,
) (int64, error) {
	mallocSize := a.blockSize(key, size)
	if mallocSize < size {
		mallocSize = size
	}
	seq, err := a.storeMalloc1(ctx, key, mallocSize)
	if err != nil {
		return 0, err
	}
	a.commitSegment(ctx, ck, owner, seq+size, seq+mallocSize, mill)
	return seq, nil
}

// handleExceed replenishes after a partial-exhaustion result. The caller
// consumed [res.CurrSeq, res.LastSeq) from the in-flight tail and needs
// (size - tailUsed) more ids from a fresh segment.
//
// If the store-returned start equals res.LastSeq the splice is exact: the
// caller's user-visible range is [res.CurrSeq, res.CurrSeq+size). Otherwise
// the in-flight tail is wasted (gap) and the caller starts fresh at the
// store-returned seq.
func (a *Allocator) handleExceed(
	ctx context.Context,
	ck, key string,
	size int64,
	res MallocResult,
) (int64, error) {
	tailUsed := res.LastSeq - res.CurrSeq
	freshNeeded := size - tailUsed
	mallocSize := a.blockSize(key, freshNeeded)
	if mallocSize < freshNeeded {
		mallocSize = freshNeeded
	}
	seq, err := a.storeMalloc1(ctx, key, mallocSize)
	if err != nil {
		return 0, err
	}

	if seq == res.LastSeq {
		// Splice OK: caller-visible range = [res.CurrSeq, res.CurrSeq+size).
		// CURR after commit = res.CurrSeq + size (which equals seq + freshNeeded).
		newCurr := res.CurrSeq + size
		a.commitSegment(ctx, ck, res.Owner, newCurr, seq+mallocSize, res.Mill)
		return res.CurrSeq, nil
	}

	// Splice mismatch: drop the in-flight tail, surface the new start.
	a.wastedSegments.Add(1)
	logx.WithContext(ctx).Infof(
		"alloc: cache last seq mismatch (tail wasted): key=%s tail=[%d,%d) got=%d",
		key, res.CurrSeq, res.LastSeq, seq,
	)
	a.commitSegment(ctx, ck, res.Owner, seq+size, seq+mallocSize, res.Mill)
	return seq, nil
}

// storeMalloc1 wraps store.Malloc with metric counters.
func (a *Allocator) storeMalloc1(ctx context.Context, key string, mallocSize int64) (int64, error) {
	a.storeMalloc.Add(1)
	seq, err := a.store.Malloc(ctx, key, mallocSize)
	if err != nil {
		a.storeErr.Add(1)
		return 0, err
	}
	return seq, nil
}

// commitSegment commits the (curr, last) pair to the cache under owner. On
// commit failure or lock-lost, the freshly fetched segment past the caller's
// `size` ids is wasted (a producer-visible gap), but the ids already returned
// to the caller stay valid because the store's max_seq advance is durable.
func (a *Allocator) commitSegment(
	ctx context.Context,
	ck, owner string,
	currSeq, lastSeq, mill int64,
) {
	state, err := a.setSeqRetry(ctx, ck, owner, currSeq, lastSeq, mill)
	if err != nil {
		a.wastedSegments.Add(1)
		logx.WithContext(ctx).Errorf(
			"alloc: setSeq failed (segment wasted): key=%s owner=%s curr=%d last=%d err=%v",
			ck, owner, currSeq, lastSeq, err,
		)
		return
	}
	if state == SetSeqLockLost {
		a.setSeqLockLost.Add(1)
		a.wastedSegments.Add(1)
		logx.WithContext(ctx).Infof(
			"alloc: setSeq lock lost (segment wasted): key=%s owner=%s curr=%d last=%d",
			ck, owner, currSeq, lastSeq,
		)
	}
}

func (a *Allocator) setSeqRetry(
	ctx context.Context,
	key, owner string,
	currSeq, lastSeq, mill int64,
) (SetSeqState, error) {
	var lastErr error
	for i := 0; i < a.retries; i++ {
		state, err := a.cache.SetSeq(ctx, key, owner, currSeq, lastSeq, mill)
		if err != nil {
			lastErr = err
			logx.WithContext(ctx).Errorf(
				"alloc: set seq cache failed: key=%s owner=%s curr=%d last=%d attempt=%d err=%v",
				key, owner, currSeq, lastSeq, i+1, err,
			)
			if waitErr := a.waitLock(ctx); waitErr != nil {
				return 0, waitErr
			}
			continue
		}
		return state, nil
	}
	if lastErr != nil {
		return 0, lastErr
	}
	return 0, ErrLockTimeout
}

func (a *Allocator) waitLock(ctx context.Context) error {
	timer := time.NewTimer(a.wait)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
