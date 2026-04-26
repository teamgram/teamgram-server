package alloc

import (
	"context"
	"errors"
	"testing"
	"time"
)

type fakeCache struct {
	mallocResults [][]int64
	setSeqCalls   []setSeqCall
}

type setSeqCall struct {
	key     string
	owner   int64
	currSeq int64
	lastSeq int64
	mill    int64
}

func (f *fakeCache) Malloc(ctx context.Context, key string, size int64) ([]int64, error) {
	if len(f.mallocResults) == 0 {
		return nil, errors.New("no malloc result")
	}
	res := f.mallocResults[0]
	f.mallocResults = f.mallocResults[1:]
	return res, nil
}

func (f *fakeCache) SetSeq(ctx context.Context, key string, owner, currSeq, lastSeq, mill int64) (int64, error) {
	f.setSeqCalls = append(f.setSeqCalls, setSeqCall{
		key:     key,
		owner:   owner,
		currSeq: currSeq,
		lastSeq: lastSeq,
		mill:    mill,
	})
	return cacheSetSuccess, nil
}

type fakeSeqStore struct {
	mallocCalls []mallocCall
	nextSeq     int64
}

type mallocCall struct {
	key  string
	size int64
}

func (f *fakeSeqStore) Malloc(ctx context.Context, key string, size int64) (int64, error) {
	f.mallocCalls = append(f.mallocCalls, mallocCall{key: key, size: size})
	return f.nextSeq, nil
}

func (f *fakeSeqStore) GetMaxSeq(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

func (f *fakeSeqStore) SetMaxSeq(ctx context.Context, key string, seq int64) error {
	return nil
}

func TestAllocatorMallocMissAllocatesBackingRangeAndReturnsFirstSeq(t *testing.T) {
	cache := &fakeCache{
		mallocResults: [][]int64{{cacheMallocMiss, 1234, 1700000000000}},
	}
	store := &fakeSeqStore{nextSeq: 40}
	a := NewAllocator(cache, store, WithWait(time.Nanosecond))

	seq, err := a.Malloc(context.Background(), "inbox:1001", 3)
	if err != nil {
		t.Fatalf("Malloc() error = %v", err)
	}
	if seq != 40 {
		t.Fatalf("Malloc() seq = %d, want 40", seq)
	}
	if len(store.mallocCalls) != 1 {
		t.Fatalf("store malloc calls = %d, want 1", len(store.mallocCalls))
	}
	if got, want := store.mallocCalls[0].size, int64(53); got != want {
		t.Fatalf("store malloc size = %d, want %d", got, want)
	}
	if len(cache.setSeqCalls) != 1 {
		t.Fatalf("cache setSeq calls = %d, want 1", len(cache.setSeqCalls))
	}
	call := cache.setSeqCalls[0]
	if call.currSeq != 43 || call.lastSeq != 93 || call.owner != 1234 {
		t.Fatalf("cache setSeq call = %+v, want curr=43 last=93 owner=1234", call)
	}
}

func TestAllocatorMallocCacheHitDoesNotTouchBackingStore(t *testing.T) {
	cache := &fakeCache{
		mallocResults: [][]int64{{cacheMallocSuccess, 10, 60, 1700000001000}},
	}
	store := &fakeSeqStore{}
	a := NewAllocator(cache, store, WithWait(time.Nanosecond))

	seq, err := a.Malloc(context.Background(), "inbox:1001", 5)
	if err != nil {
		t.Fatalf("Malloc() error = %v", err)
	}
	if seq != 10 {
		t.Fatalf("Malloc() seq = %d, want 10", seq)
	}
	if len(store.mallocCalls) != 0 {
		t.Fatalf("store malloc calls = %d, want 0", len(store.mallocCalls))
	}
}

func TestAllocatorRejectsNegativeSize(t *testing.T) {
	a := NewAllocator(&fakeCache{}, &fakeSeqStore{}, WithWait(time.Nanosecond))

	_, err := a.Malloc(context.Background(), "inbox:1001", -1)
	if err == nil {
		t.Fatal("Malloc() error = nil, want error")
	}
}
