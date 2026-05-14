package sessionstate

import "sync"

type runtimeSessionKey struct {
	authKeyId   int64
	authKeyType int32
	sessionId   int64
}

type runtimeState struct {
	mu             sync.RWMutex
	destroyed      map[runtimeSessionKey]struct{}
	outboundUnacks map[runtimeSessionKey]map[int64]struct{}
}

func newRuntimeState() *runtimeState {
	return &runtimeState{
		destroyed:      make(map[runtimeSessionKey]struct{}),
		outboundUnacks: make(map[runtimeSessionKey]map[int64]struct{}),
	}
}

func (r *runtimeState) recordOutbound(key runtimeSessionKey, msgID int64) {
	if r == nil || key.authKeyId == 0 || key.sessionId == 0 || msgID == 0 {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	ids := r.outboundUnacks[key]
	if ids == nil {
		ids = make(map[int64]struct{})
		r.outboundUnacks[key] = ids
	}
	ids[msgID] = struct{}{}
}

func (r *runtimeState) ackOutbound(key runtimeSessionKey, msgIDs []int64) {
	if r == nil || len(msgIDs) == 0 {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	ids := r.outboundUnacks[key]
	if ids == nil {
		return
	}
	for _, msgID := range msgIDs {
		delete(ids, msgID)
	}
	if len(ids) == 0 {
		delete(r.outboundUnacks, key)
	}
}

func (r *runtimeState) hasOutboundUnacked(key runtimeSessionKey, msgID int64) bool {
	if r == nil {
		return false
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	ids := r.outboundUnacks[key]
	if ids == nil {
		return false
	}
	_, ok := ids[msgID]
	return ok
}

func (r *runtimeState) destroySession(key runtimeSessionKey) bool {
	if r == nil || key.authKeyId == 0 || key.sessionId == 0 {
		return false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.destroyed[key]; ok {
		return false
	}
	r.destroyed[key] = struct{}{}
	delete(r.outboundUnacks, key)
	return true
}

func (r *runtimeState) isDestroyed(key runtimeSessionKey) bool {
	if r == nil {
		return false
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, ok := r.destroyed[key]
	return ok
}
