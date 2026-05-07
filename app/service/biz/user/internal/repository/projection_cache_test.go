package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
)

func TestProjectionCacheKeysAreVersioned(t *testing.T) {
	tests := map[string]string{
		projectionFactsCacheKey(11):      "user:facts:v1:11",
		projectionPrivacyCacheKey(11):    "user:privacy:v1:11",
		projectionContactMapCacheKey(11): "user:contact-map:v1:11",
		projectionPresenceCacheKey(11):   "user:presence:v1:11",
	}
	for got, want := range tests {
		if got != want {
			t.Fatalf("key = %q, want %q", got, want)
		}
	}
}

func TestProjectionCacheDecodeRejectsStaleSchema(t *testing.T) {
	got, ok := decodeProjectionCache[map[string]int](`{"schema_version":0,"data":{"a":1}}`)
	if ok || got != nil {
		t.Fatalf("stale schema decoded: got=%v ok=%v", got, ok)
	}
}

func TestProjectionCacheDecodeStatusDistinguishesMissStaleAndCorrupt(t *testing.T) {
	if _, status := decodeProjectionCacheStatus[map[string]int](""); status != projectionCacheDecodeMiss {
		t.Fatalf("empty status = %v, want miss", status)
	}
	if _, status := decodeProjectionCacheStatus[map[string]int](`{"schema_version":0,"data":{"a":1}}`); status != projectionCacheDecodeStale {
		t.Fatalf("stale status = %v, want stale", status)
	}
	if _, status := decodeProjectionCacheStatus[map[string]int](`{"schema_version":`); status != projectionCacheDecodeCorrupt {
		t.Fatalf("corrupt status = %v, want corrupt", status)
	}
	if got, status := decodeProjectionCacheStatus[map[string]int](`{"schema_version":1,"data":{"a":1}}`); status != projectionCacheDecodeHit || got["a"] != 1 {
		t.Fatalf("hit decode = %v status=%v", got, status)
	}
}

func TestProjectionComponentCacheRoundTrip(t *testing.T) {
	cache := newFakeProjectionBatchCache()
	r := &Repository{CachedConn: sqlc.NewConnWithCache(nil, cache)}
	ctx := context.Background()
	key := projectionPresenceCacheKey(42)

	r.setProjectionComponentCache(ctx, key, projectionPresenceCacheDTO{
		UserID:      42,
		HasPresence: true,
		LastSeenAt:  100,
		Expires:     200,
	})

	var got projectionPresenceCacheDTO
	if !r.getProjectionComponentCache(ctx, key, &got) {
		t.Fatalf("cache miss after set")
	}
	if got.UserID != 42 || !got.HasPresence || got.LastSeenAt != 100 || got.Expires != 200 {
		t.Fatalf("cache value = %+v", got)
	}
	if cache.sets != 1 || cache.gets != 1 {
		t.Fatalf("cache calls: gets=%d sets=%d", cache.gets, cache.sets)
	}
}

func TestProjectionComponentCacheBulkRead(t *testing.T) {
	cache := newFakeProjectionBatchCache()
	r := &Repository{CachedConn: sqlc.NewConnWithCache(nil, cache)}
	ctx := context.Background()
	key1 := projectionPresenceCacheKey(42)
	key2 := projectionPresenceCacheKey(43)

	r.setProjectionComponentCache(ctx, key1, projectionPresenceCacheDTO{UserID: 42, HasPresence: true, LastSeenAt: 100, Expires: 200})
	r.setProjectionComponentCache(ctx, key2, projectionPresenceCacheDTO{UserID: 43})

	got := getProjectionComponentCaches[projectionPresenceCacheDTO](r, ctx, []string{key1, key2})
	if len(got) != 2 {
		t.Fatalf("bulk hits = %#v", got)
	}
	if got[key1].UserID != 42 || !got[key1].HasPresence || got[key2].UserID != 43 {
		t.Fatalf("bulk values = %#v", got)
	}
	if cache.takes != 1 {
		t.Fatalf("bulk cache calls = %d, want 1", cache.takes)
	}
}

func TestProjectionCacheIdentityMismatchDeletesKey(t *testing.T) {
	cache := newFakeProjectionBatchCache()
	r := &Repository{CachedConn: sqlc.NewConnWithCache(nil, cache)}
	ctx := context.Background()
	key := projectionPresenceCacheKey(42)

	r.setProjectionComponentCache(ctx, key, projectionPresenceCacheDTO{UserID: 43})
	r.logProjectionCacheIdentityMismatch(ctx, key, "presence", 42, 43)

	var got projectionPresenceCacheDTO
	if err := cache.GetCtx(ctx, key, &got); err != sql.ErrNoRows {
		t.Fatalf("cache key still present after mismatch delete: err=%v got=%+v", err, got)
	}
}

type fakeProjectionBatchCache struct {
	values map[string]interface{}
	gets   int
	sets   int
	takes  int
}

func newFakeProjectionBatchCache() *fakeProjectionBatchCache {
	return &fakeProjectionBatchCache{values: make(map[string]interface{})}
}

func (c *fakeProjectionBatchCache) Del(keys ...string) error {
	return c.DelCtx(context.Background(), keys...)
}

func (c *fakeProjectionBatchCache) DelCtx(_ context.Context, keys ...string) error {
	for _, key := range keys {
		delete(c.values, key)
	}
	return nil
}

func (c *fakeProjectionBatchCache) Get(key string, val any) error {
	return c.GetCtx(context.Background(), key, val)
}

func (c *fakeProjectionBatchCache) GetCtx(_ context.Context, key string, val any) error {
	c.gets++
	raw, ok := c.values[key]
	if !ok {
		return sql.ErrNoRows
	}
	b, err := json.Marshal(raw)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, val)
}

func (c *fakeProjectionBatchCache) IsNotFound(err error) bool {
	return err == sql.ErrNoRows
}

func (c *fakeProjectionBatchCache) Set(key string, val any) error {
	return c.SetCtx(context.Background(), key, val)
}

func (c *fakeProjectionBatchCache) SetCtx(_ context.Context, key string, val any) error {
	c.sets++
	c.values[key] = val
	return nil
}

func (c *fakeProjectionBatchCache) SetWithExpire(key string, val any, _ time.Duration) error {
	return c.SetCtx(context.Background(), key, val)
}

func (c *fakeProjectionBatchCache) SetWithExpireCtx(ctx context.Context, key string, val any, _ time.Duration) error {
	return c.SetCtx(ctx, key, val)
}

func (c *fakeProjectionBatchCache) Take(val any, key string, query func(val any) error) error {
	return c.TakeCtx(context.Background(), val, key, query)
}

func (c *fakeProjectionBatchCache) TakeCtx(ctx context.Context, val any, key string, query func(val any) error) error {
	if err := c.GetCtx(ctx, key, val); err == nil {
		return nil
	}
	if err := query(val); err != nil {
		return err
	}
	return c.SetCtx(ctx, key, val)
}

func (c *fakeProjectionBatchCache) TakeWithExpire(val any, key string, query func(val any, expire time.Duration) error) error {
	return c.TakeWithExpireCtx(context.Background(), val, key, query)
}

func (c *fakeProjectionBatchCache) TakeWithExpireCtx(ctx context.Context, val any, key string, query func(val any, expire time.Duration) error) error {
	if err := c.GetCtx(ctx, key, val); err == nil {
		return nil
	}
	if err := query(val, time.Minute); err != nil {
		return err
	}
	return c.SetCtx(ctx, key, val)
}

func (c *fakeProjectionBatchCache) Takes(query func(keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	return c.TakesCtx(context.Background(), query, cacheF, keys...)
}

func (c *fakeProjectionBatchCache) TakesCtx(ctx context.Context, query func(keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	c.takes++
	missKeys := make([]string, 0, len(keys))
	for _, key := range keys {
		raw, ok := c.values[key]
		if !ok {
			missKeys = append(missKeys, key)
			continue
		}
		b, err := json.Marshal(raw)
		if err != nil {
			missKeys = append(missKeys, key)
			continue
		}
		if _, err := cacheF(key, string(b)); err != nil {
			missKeys = append(missKeys, key)
		}
	}
	values, err := query(missKeys...)
	if err != nil {
		return err
	}
	for key, value := range values {
		converted, err := cacheF(key, "")
		if err == nil && converted != nil {
			value = converted
		}
		_ = c.SetCtx(ctx, key, value)
	}
	return nil
}

func (c *fakeProjectionBatchCache) TakesWithExpire(query func(expire time.Duration, keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	return c.TakesWithExpireCtx(context.Background(), query, cacheF, keys...)
}

func (c *fakeProjectionBatchCache) TakesWithExpireCtx(ctx context.Context, query func(expire time.Duration, keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	values, err := query(time.Minute, keys...)
	if err != nil {
		return err
	}
	for key, value := range values {
		converted, err := cacheF(key, "")
		if err == nil && converted != nil {
			value = converted
		}
		_ = c.SetCtx(ctx, key, value)
	}
	return nil
}
