package core

import (
	"testing"
	"time"
)

func TestTypingLimiterCoalescesSenderPeer(t *testing.T) {
	limiter := newTypingLimiter(5 * time.Second)
	now := time.Unix(100, 0)
	if !limiter.Allow(1, 2, now) {
		t.Fatal("first Allow() = false, want true")
	}
	if limiter.Allow(1, 2, now.Add(4*time.Second)) {
		t.Fatal("second Allow() = true, want false")
	}
	if !limiter.Allow(1, 2, now.Add(5*time.Second)) {
		t.Fatal("third Allow() = false, want true")
	}
}
