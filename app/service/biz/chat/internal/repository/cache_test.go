package repository

import "testing"

func TestCacheKeys(t *testing.T) {
	if got := chatAggregateCacheKey(10); got != "chat:aggregate:10" {
		t.Fatalf("aggregate key = %q", got)
	}
	if got := chatParticipantCacheKey(10, 20); got != "chat:participant:10:20" {
		t.Fatalf("participant key = %q", got)
	}
	if got := createChatFloodKey(30); got != "chat:create:flood:30" {
		t.Fatalf("flood key = %q", got)
	}
}
