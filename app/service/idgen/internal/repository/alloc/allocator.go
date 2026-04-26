package alloc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	cacheMallocSuccess = int64(0)
	cacheMallocMiss    = int64(1)
	cacheMallocLocked  = int64(2)
	cacheMallocExceed  = int64(3)

	cacheSetSuccess         = int64(0)
	cacheSetLockExpired     = int64(1)
	cacheSetLockOwnerChange = int64(2)
)

var (
	ErrInvalidSize = errors.New("alloc: size must be greater than or equal to 0")
	ErrLockTimeout = errors.New("alloc: waiting for cache lock timeout")
)

type Cache interface {
	Malloc(ctx context.Context, key string, size int64) ([]int64, error)
	SetSeq(ctx context.Context, key string, owner, currSeq, lastSeq, mill int64) (int64, error)
}

type SeqStore interface {
	Malloc(ctx context.Context, key string, size int64) (int64, error)
	GetMaxSeq(ctx context.Context, key string) (int64, error)
	SetMaxSeq(ctx context.Context, key string, seq int64) error
}

type Allocator struct {
	cache     Cache
	store     SeqStore
	wait      time.Duration
	retries   int
	blockSize func(key string, size int64) int64
}

type Option func(*Allocator)

func NewAllocator(cache Cache, store SeqStore, opts ...Option) *Allocator {
	a := &Allocator{
		cache:     cache,
		store:     store,
		wait:      time.Second / 4,
		retries:   10,
		blockSize: defaultBlockSize,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func WithWait(wait time.Duration) Option {
	return func(a *Allocator) {
		a.wait = wait
	}
}

func WithRetries(retries int) Option {
	return func(a *Allocator) {
		a.retries = retries
	}
}

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

func (a *Allocator) Malloc(ctx context.Context, key string, size int64) (int64, error) {
	seq, _, err := a.MallocTime(ctx, key, size)
	return seq, err
}

func (a *Allocator) GetMaxSeq(ctx context.Context, key string) (int64, error) {
	return a.Malloc(ctx, key, 0)
}

func (a *Allocator) SetMaxSeq(ctx context.Context, key string, seq int64) error {
	return a.store.SetMaxSeq(ctx, key, seq)
}

func (a *Allocator) MallocTime(ctx context.Context, key string, size int64) (int64, int64, error) {
	if size < 0 {
		return 0, 0, ErrInvalidSize
	}
	if a.cache == nil {
		seq, err := a.store.Malloc(ctx, key, size)
		return seq, 0, err
	}

	ck := cacheKey(key)
	for i := 0; i < a.retries; i++ {
		states, err := a.cache.Malloc(ctx, ck, size)
		if err != nil {
			return 0, 0, err
		}
		if len(states) == 0 {
			return 0, 0, errors.New("alloc: empty cache malloc state")
		}

		switch states[0] {
		case cacheMallocSuccess:
			if len(states) < 4 {
				return 0, 0, fmt.Errorf("alloc: invalid success state %v", states)
			}
			return states[1], states[3], nil
		case cacheMallocMiss:
			if len(states) < 3 {
				return 0, 0, fmt.Errorf("alloc: invalid miss state %v", states)
			}
			mallocSize := a.blockSize(key, size)
			seq, err := a.store.Malloc(ctx, key, mallocSize)
			if err != nil {
				return 0, 0, err
			}
			a.setSeqRetry(ctx, ck, states[1], seq+size, seq+mallocSize, states[2])
			return seq, 0, nil
		case cacheMallocLocked:
			if err := a.waitLock(ctx); err != nil {
				return 0, 0, err
			}
		case cacheMallocExceed:
			if len(states) < 5 {
				return 0, 0, fmt.Errorf("alloc: invalid exceed state %v", states)
			}
			currSeq, lastSeq, owner, mill := states[1], states[2], states[3], states[4]
			mallocSize := a.blockSize(key, size)
			seq, err := a.store.Malloc(ctx, key, mallocSize)
			if err != nil {
				return 0, 0, err
			}
			if lastSeq == seq {
				a.setSeqRetry(ctx, ck, owner, currSeq+size, seq+mallocSize, mill)
				return currSeq, mill, nil
			}
			logx.WithContext(ctx).Infof("alloc: cache last seq mismatch: key=%s curr=%d last=%d store_seq=%d", key, currSeq, lastSeq, seq)
			a.setSeqRetry(ctx, ck, owner, seq+size, seq+mallocSize, mill)
			return seq, mill, nil
		default:
			return 0, 0, fmt.Errorf("alloc: unknown cache state %d", states[0])
		}
	}

	return 0, 0, ErrLockTimeout
}

func (a *Allocator) setSeqRetry(ctx context.Context, key string, owner, currSeq, lastSeq, mill int64) {
	for i := 0; i < a.retries; i++ {
		state, err := a.cache.SetSeq(ctx, key, owner, currSeq, lastSeq, mill)
		if err != nil {
			logx.WithContext(ctx).Errorf("alloc: set seq cache failed: key=%s owner=%d curr=%d last=%d count=%d err=%v", key, owner, currSeq, lastSeq, i+1, err)
			if err := a.waitLock(ctx); err != nil {
				return
			}
			continue
		}
		switch state {
		case cacheSetSuccess, cacheSetLockExpired:
		case cacheSetLockOwnerChange:
			logx.WithContext(ctx).Infof("alloc: set seq lock owned by another allocator: key=%s owner=%d curr=%d last=%d", key, owner, currSeq, lastSeq)
		default:
			logx.WithContext(ctx).Errorf("alloc: unknown set seq state: key=%s state=%d", key, state)
		}
		return
	}
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
