package repository

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
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

func TestProjectionComponentCacheBulkRead(t *testing.T) {
	cache := newFakeProjectionRawCache()
	r := &Repository{projectionCache: cache}
	ctx := context.Background()
	key1 := projectionPresenceCacheKey(42)
	key2 := projectionPresenceCacheKey(43)

	r.setProjectionComponentCaches(ctx, map[string]interface{}{
		key1: projectionPresenceCacheDTO{UserID: 42, HasPresence: true, LastSeenAt: 100, Expires: 200},
		key2: projectionPresenceCacheDTO{UserID: 43},
	})

	got := getProjectionComponentCaches[projectionPresenceCacheDTO](r, ctx, []string{key1, key2})
	if len(got) != 2 {
		t.Fatalf("bulk hits = %#v", got)
	}
	if got[key1].UserID != 42 || !got[key1].HasPresence || got[key2].UserID != 43 {
		t.Fatalf("bulk values = %#v", got)
	}
	if cache.setManyCalls != 1 || cache.getManyCalls != 1 {
		t.Fatalf("cache calls: setMany=%d getMany=%d", cache.setManyCalls, cache.getManyCalls)
	}
}

func TestProjectionComponentCacheBulkReadDeletesInvalidEnvelope(t *testing.T) {
	cache := newFakeProjectionRawCache()
	r := &Repository{projectionCache: cache}
	ctx := context.Background()
	key := projectionPresenceCacheKey(42)
	cache.setRawEnvelope(key, projectionCacheEnvelope[projectionPresenceCacheDTO]{
		SchemaVersion: 0,
		Data:          projectionPresenceCacheDTO{UserID: 42},
	})

	got := getProjectionComponentCaches[projectionPresenceCacheDTO](r, ctx, []string{key})
	if len(got) != 0 {
		t.Fatalf("stale cache returned hit: %#v", got)
	}
	if _, ok := cache.values[key]; ok {
		t.Fatalf("stale cache key still present")
	}
	if cache.deleteCalls != 1 {
		t.Fatalf("delete calls = %d, want 1", cache.deleteCalls)
	}
}

func TestProjectionCacheIdentityMismatchDeletesKey(t *testing.T) {
	cache := newFakeProjectionRawCache()
	r := &Repository{projectionCache: cache}
	ctx := context.Background()
	key := projectionPresenceCacheKey(42)

	r.setProjectionComponentCaches(ctx, map[string]interface{}{key: projectionPresenceCacheDTO{UserID: 43}})
	r.logProjectionCacheIdentityMismatch(ctx, key, "presence", 42, 43)

	if _, ok := cache.values[key]; ok {
		t.Fatalf("cache key still present after mismatch delete")
	}
}

func TestProjectionBotFactsDoNotCarryToken(t *testing.T) {
	const token = "bot-secret-token"
	bot := botDataFromModel(&model.Bots{BotId: 42, Token: token, BotInfoVersion: 7}).ToBotData()
	if bot.Token != "" {
		t.Fatalf("botDataFromModel carried token: %+v", bot)
	}
	dto := botCacheDTOFromData(bot)
	raw, err := json.Marshal(dto)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), token) || strings.Contains(string(raw), "token") {
		t.Fatalf("bot cache dto carried token: %s", raw)
	}
	if got := botDataFromCacheDTO(dto).ToBotData(); got.Token != "" {
		t.Fatalf("botDataFromCacheDTO carried token: %+v", got)
	}
}

func TestProjectionContactMapCoverageIsExplicitForPartialMaps(t *testing.T) {
	dto := projectionContactMapCacheDTO{
		OwnerUserID:       1,
		Contacts:          map[int64]projectionContactFact{2: {FirstName: "A"}},
		CoveredContactIDs: []int64{2},
	}
	if !projectionContactMapCovers(dto, []int64{2}) {
		t.Fatalf("expected contact map to cover requested id")
	}
	if projectionContactMapCovers(dto, []int64{2, 3}) {
		t.Fatalf("partial contact map covered an unknown requested id")
	}
	dto.CoveredContactIDs = nil
	if projectionContactMapCovers(dto, []int64{2}) {
		t.Fatalf("contact map without explicit coverage covered requested id")
	}

	required := projectionRequiredContactIDs(1, int64Set([]int64{1}), int64Set([]int64{2}), []int64{1}, []int64{2})
	if len(required) != 1 || required[0] != 2 {
		t.Fatalf("required ids = %v, want [2]", required)
	}
}

func TestProjectionContactCacheFactsOnlyUseRequiredIds(t *testing.T) {
	facts := projectionFacts{Contacts: make(map[contactKey]*projectionContactFact)}
	addProjectionContactCacheFacts(1, []int64{2}, map[int64]projectionContactFact{
		2: {FirstName: "Covered"},
		3: {FirstName: "Stale"},
	}, facts)

	if facts.Contacts[contactKey{OwnerUserId: 1, ContactUserId: 2}] == nil {
		t.Fatalf("covered contact was not added")
	}
	if facts.Contacts[contactKey{OwnerUserId: 1, ContactUserId: 3}] != nil {
		t.Fatalf("non-required stale contact was added")
	}
}

type fakeProjectionRawCache struct {
	values       map[string]string
	getManyCalls int
	setManyCalls int
	deleteCalls  int
}

func newFakeProjectionRawCache() *fakeProjectionRawCache {
	return &fakeProjectionRawCache{values: make(map[string]string)}
}

func (c *fakeProjectionRawCache) getMany(_ context.Context, keys []string) (map[string]string, error) {
	c.getManyCalls++
	out := make(map[string]string, len(keys))
	for _, key := range keys {
		if value, ok := c.values[key]; ok {
			out[key] = value
		}
	}
	return out, nil
}

func (c *fakeProjectionRawCache) setMany(_ context.Context, values map[string]string) error {
	c.setManyCalls++
	for key, value := range values {
		c.values[key] = value
	}
	return nil
}

func (c *fakeProjectionRawCache) delete(_ context.Context, keys ...string) error {
	c.deleteCalls++
	for _, key := range keys {
		delete(c.values, key)
	}
	return nil
}

func (c *fakeProjectionRawCache) setRawEnvelope(key string, value interface{}) {
	raw, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	c.values[key] = string(raw)
}
