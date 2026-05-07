package repository

import "testing"

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
