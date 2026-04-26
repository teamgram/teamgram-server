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
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/teamgram/marmota/pkg/stores/kv"
)

// evalStore is the minimal Redis EVAL surface the Lua scripts need.
//
// It is satisfied by go-zero's kv.Store (and therefore by marmota's
// kv.ExtStore) but is kept as a private interface so tests can plug in a
// fake without depending on the full Redis stack.
type evalStore interface {
	EvalCtx(ctx context.Context, script, key string, args ...any) (any, error)
}

// Defaults for XKVCache. lockTTL is intentionally generous to absorb p99
// store latency; dataTTL is short enough to bound idle Redis memory while
// still keeping hot keys warm. Both are configurable via Options.
const (
	defaultLockTTL = 5 * time.Second
	defaultDataTTL = 7 * 24 * time.Hour
)

type xkvCache struct {
	kv       evalStore
	lockTTL  time.Duration
	dataTTL  time.Duration
	nowMilli func() int64

	// ownerSeq generates monotonic owner ids unique to this process. The
	// initial value is seeded from the wall clock at construction time so
	// owner ids do not collide across restarts of the same instance.
	ownerSeq atomic.Int64
}

// XKVOption configures an XKVCache.
type XKVOption func(*xkvCache)

// WithLockTTL sets the logical lock TTL. The lock is enforced inside the Lua
// script using nowMillis - LOCK_AT < lockTTL, so it does not require a
// separate Redis EXPIRE on the lock field.
func WithLockTTL(d time.Duration) XKVOption {
	return func(c *xkvCache) {
		if d > 0 {
			c.lockTTL = d
		}
	}
}

// WithDataTTL sets the data key TTL. Hot keys are kept alive on every
// successful Malloc/SetSeq; cold keys are evicted by Redis after this
// duration of inactivity and re-loaded from the store on next access.
func WithDataTTL(d time.Duration) XKVOption {
	return func(c *xkvCache) {
		if d > 0 {
			c.dataTTL = d
		}
	}
}

// WithNowMilli installs a custom millis clock. Intended for tests; the
// production default is time.Now().UnixMilli.
func WithNowMilli(fn func() int64) XKVOption {
	return func(c *xkvCache) {
		if fn != nil {
			c.nowMilli = fn
		}
	}
}

// NewXKVCache wraps a marmota kv.ExtStore as a Cache.
func NewXKVCache(store kv.ExtStore, opts ...XKVOption) Cache {
	return newXKVCacheWithStore(store, opts...)
}

// NewXKVCacheWithStore exposes the underlying eval-only contract for tests
// that don't want to construct a full kv.ExtStore.
func NewXKVCacheWithStore(store evalStore, opts ...XKVOption) Cache {
	return newXKVCacheWithStore(store, opts...)
}

func newXKVCacheWithStore(store evalStore, opts ...XKVOption) *xkvCache {
	c := &xkvCache{
		kv:       store,
		lockTTL:  defaultLockTTL,
		dataTTL:  defaultDataTTL,
		nowMilli: func() int64 { return time.Now().UnixMilli() },
	}
	// Seed ownerSeq from wall clock so concurrent processes do not produce
	// the same owner ids.
	c.ownerSeq.Store(time.Now().UnixNano())
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *xkvCache) nextOwner() string {
	return strconv.FormatInt(c.ownerSeq.Add(1), 10)
}

// Malloc implements Cache.
func (c *xkvCache) Malloc(ctx context.Context, key string, size int64) (MallocResult, error) {
	owner := c.nextOwner()
	res, err := c.kv.EvalCtx(
		ctx,
		mallocScript,
		key,
		size,
		int64(c.lockTTL/time.Millisecond),
		int64(c.dataTTL/time.Second),
		c.nowMilli(),
		owner,
	)
	if err != nil {
		return MallocResult{}, fmt.Errorf("alloc: xkv malloc eval: %w", err)
	}
	arr, err := toAnySlice(res)
	if err != nil {
		return MallocResult{}, fmt.Errorf("alloc: xkv malloc parse: %w", err)
	}
	if len(arr) == 0 {
		return MallocResult{}, fmt.Errorf("alloc: xkv malloc empty result")
	}

	state, err := toInt64(arr[0])
	if err != nil {
		return MallocResult{}, fmt.Errorf("alloc: xkv malloc state: %w", err)
	}

	switch MallocState(state) {
	case MallocSuccess:
		if len(arr) < 4 {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc success arity=%d", len(arr))
		}
		curr, err := toInt64(arr[1])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc curr: %w", err)
		}
		last, err := toInt64(arr[2])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc last: %w", err)
		}
		mill, err := toInt64(arr[3])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc mill: %w", err)
		}
		return MallocResult{
			State:   MallocSuccess,
			CurrSeq: curr,
			LastSeq: last,
			Mill:    mill,
		}, nil

	case MallocMiss:
		if len(arr) < 3 {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc miss arity=%d", len(arr))
		}
		ownerStr, err := toString(arr[1])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc miss owner: %w", err)
		}
		mill, err := toInt64(arr[2])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc miss mill: %w", err)
		}
		return MallocResult{
			State: MallocMiss,
			Owner: ownerStr,
			Mill:  mill,
		}, nil

	case MallocLocked:
		return MallocResult{State: MallocLocked}, nil

	case MallocExceed:
		if len(arr) < 5 {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc exceed arity=%d", len(arr))
		}
		curr, err := toInt64(arr[1])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc exceed curr: %w", err)
		}
		last, err := toInt64(arr[2])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc exceed last: %w", err)
		}
		ownerStr, err := toString(arr[3])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc exceed owner: %w", err)
		}
		mill, err := toInt64(arr[4])
		if err != nil {
			return MallocResult{}, fmt.Errorf("alloc: xkv malloc exceed mill: %w", err)
		}
		return MallocResult{
			State:   MallocExceed,
			CurrSeq: curr,
			LastSeq: last,
			Owner:   ownerStr,
			Mill:    mill,
		}, nil

	default:
		return MallocResult{}, fmt.Errorf("%w: xkv malloc state=%d", ErrInvalidState, state)
	}
}

// SetSeq implements Cache.
func (c *xkvCache) SetSeq(
	ctx context.Context,
	key, owner string,
	currSeq, lastSeq, mill int64,
) (SetSeqState, error) {
	if owner == "" {
		return 0, fmt.Errorf("alloc: xkv setSeq owner is empty")
	}
	if lastSeq < currSeq {
		return 0, fmt.Errorf("alloc: xkv setSeq lastSeq %d < currSeq %d", lastSeq, currSeq)
	}
	res, err := c.kv.EvalCtx(
		ctx,
		setSeqScript,
		key,
		owner,
		int64(c.dataTTL/time.Second),
		currSeq,
		lastSeq,
		mill,
	)
	if err != nil {
		return 0, fmt.Errorf("alloc: xkv setSeq eval: %w", err)
	}
	state, err := toInt64(res)
	if err != nil {
		return 0, fmt.Errorf("alloc: xkv setSeq parse: %w", err)
	}
	switch SetSeqState(state) {
	case SetSeqSuccess, SetSeqLockLost:
		return SetSeqState(state), nil
	default:
		return 0, fmt.Errorf("%w: xkv setSeq state=%d", ErrInvalidState, state)
	}
}

func toAnySlice(v any) ([]any, error) {
	switch x := v.(type) {
	case []any:
		return x, nil
	case []int64:
		out := make([]any, len(x))
		for i, n := range x {
			out[i] = n
		}
		return out, nil
	default:
		return nil, fmt.Errorf("alloc: invalid array type %T", v)
	}
}

func toInt64(v any) (int64, error) {
	switch x := v.(type) {
	case int:
		return int64(x), nil
	case int8:
		return int64(x), nil
	case int16:
		return int64(x), nil
	case int32:
		return int64(x), nil
	case int64:
		return x, nil
	case uint:
		return int64(x), nil
	case uint8:
		return int64(x), nil
	case uint16:
		return int64(x), nil
	case uint32:
		return int64(x), nil
	case uint64:
		return int64(x), nil
	case string:
		return strconv.ParseInt(x, 10, 64)
	case []byte:
		return strconv.ParseInt(string(x), 10, 64)
	default:
		return 0, fmt.Errorf("alloc: invalid int64 type %T", v)
	}
}

func toString(v any) (string, error) {
	switch x := v.(type) {
	case string:
		return x, nil
	case []byte:
		return string(x), nil
	case int:
		return strconv.FormatInt(int64(x), 10), nil
	case int64:
		return strconv.FormatInt(x, 10), nil
	default:
		return "", fmt.Errorf("alloc: invalid string type %T", v)
	}
}
