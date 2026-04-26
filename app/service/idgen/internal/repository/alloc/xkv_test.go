package alloc

import (
	"context"
	"testing"
)

type fakeEvalStore struct {
	script string
	key    string
	args   []any
	res    any
}

func (f *fakeEvalStore) EvalCtx(ctx context.Context, script, key string, args ...any) (any, error) {
	f.script = script
	f.key = key
	f.args = args
	return f.res, nil
}

func TestXKVCacheMallocUsesEvalCtxAndParsesRedisArray(t *testing.T) {
	store := &fakeEvalStore{res: []any{int64(0), "10", int64(60), []byte("1700000000000")}}
	cache := NewXKVCacheWithStore(store)

	got, err := cache.Malloc(context.Background(), "idgen:malloc_seq:inbox:1001", 5)
	if err != nil {
		t.Fatalf("Malloc() error = %v", err)
	}
	want := []int64{0, 10, 60, 1700000000000}
	if len(got) != len(want) {
		t.Fatalf("Malloc() len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Malloc()[%d] = %d, want %d", i, got[i], want[i])
		}
	}
	if store.key != "idgen:malloc_seq:inbox:1001" {
		t.Fatalf("EvalCtx key = %q", store.key)
	}
	if len(store.args) != 4 || store.args[0] != int64(5) {
		t.Fatalf("EvalCtx args = %#v, want size as first arg", store.args)
	}
}

func TestMongoStoreFakeIsReservedOnly(t *testing.T) {
	store := NewMongoStoreFake()

	_, err := store.Malloc(context.Background(), "inbox:1001", 1)
	if err == nil {
		t.Fatal("Malloc() error = nil, want reserved fake error")
	}
}
