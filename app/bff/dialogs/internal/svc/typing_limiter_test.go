package svc

import (
	"testing"
	"time"
)

func TestTypingLimiterCleansStaleSenderPeerEntries(t *testing.T) {
	limiter := NewTypingLimiter(5 * time.Second).(*typingLimiter)
	now := time.Unix(100, 0)

	if !limiter.Allow(1, 2, now) {
		t.Fatal("Allow(1, 2) = false, want true")
	}
	if !limiter.Allow(3, 4, now) {
		t.Fatal("Allow(3, 4) = false, want true")
	}
	if got := len(limiter.last); got != 2 {
		t.Fatalf("len(last) = %d, want 2", got)
	}

	if !limiter.Allow(5, 6, now.Add(6*time.Second)) {
		t.Fatal("Allow(5, 6) after stale interval = false, want true")
	}
	if got := len(limiter.last); got != 1 {
		t.Fatalf("len(last) after cleanup = %d, want only the fresh key", got)
	}
	if _, ok := limiter.last[typingKey{sender: 5, peer: 6}]; !ok {
		t.Fatal("fresh sender/peer key was not retained")
	}
}
