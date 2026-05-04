package svc

import (
	"sync"
	"time"
)

type TypingLimiter interface {
	Allow(senderUserID, peerUserID int64, now time.Time) bool
}

type typingLimiter struct {
	mu       sync.Mutex
	interval time.Duration
	last     map[typingKey]time.Time
}

type typingKey struct {
	sender int64
	peer   int64
}

func NewTypingLimiter(interval time.Duration) TypingLimiter {
	return &typingLimiter{interval: interval, last: make(map[typingKey]time.Time)}
}

func (l *typingLimiter) Allow(senderUserID, peerUserID int64, now time.Time) bool {
	if l == nil || l.interval <= 0 {
		return true
	}
	key := typingKey{sender: senderUserID, peer: peerUserID}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.cleanupLocked(now)
	last, ok := l.last[key]
	if ok && now.Sub(last) < l.interval {
		return false
	}
	l.last[key] = now
	return true
}

func (l *typingLimiter) cleanupLocked(now time.Time) {
	for key, last := range l.last {
		if now.Sub(last) >= l.interval {
			delete(l.last, key)
		}
	}
}
