package sessionstate

import (
	"sync"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const selectorQueueLimit = 64

type roleSessionKey struct {
	permAuthKeyId int64
	authKeyId     int64
	authKeyType   int32
	sessionId     int64
}

type selectorToken uint64

type selectorKind int

const (
	selectorGated selectorKind = iota
	selectorImmediate
)

type roleTransition struct {
	method string
	key    roleSessionKey
	token  selectorToken
	kind   selectorKind
}

type selectorCompletionResult struct {
	stale bool
	flush pendingUpdatesState
}

type roleRegistry struct {
	mu        sync.RWMutex
	nextToken selectorToken
	byPerm    map[int64]*rolePermState
	tokenPerm map[selectorToken]int64
}

type rolePermState struct {
	main      *roleSessionKey
	candidate *roleTransition
	pending   pendingUpdatesState
}

type pendingUpdatesState struct {
	payloads       []tg.UpdatesClazz
	tooLongPending bool
}

type roleSnapshot struct {
	main           *roleSessionKey
	candidateKey   roleSessionKey
	candidateToken selectorToken
	queued         int
	pendingTooLong bool
}

func newRoleRegistry() *roleRegistry {
	return &roleRegistry{
		byPerm:    make(map[int64]*rolePermState),
		tokenPerm: make(map[selectorToken]int64),
	}
}

func (r *roleRegistry) beginTransition(method string, key roleSessionKey) (roleTransition, bool) {
	if r == nil || !isNormalRoleKey(key) {
		return roleTransition{}, false
	}
	switch method {
	case tg.ClazzName_updates_getState, tg.ClazzName_updates_getDifference, tg.ClazzName_updates_getChannelDifference:
		return r.beginSelector(method, key), true
	case tg.ClazzName_account_updateStatus:
		return r.promoteImmediate(method, key), true
	default:
		return roleTransition{}, false
	}
}

func (r *roleRegistry) beginSelector(method string, key roleSessionKey) roleTransition {
	r.mu.Lock()
	defer r.mu.Unlock()
	state := r.stateLocked(key.permAuthKeyId)
	if state.candidate != nil {
		delete(r.tokenPerm, state.candidate.token)
	}
	r.nextToken++
	transition := roleTransition{method: method, key: key, token: r.nextToken, kind: selectorGated}
	state.candidate = &transition
	r.tokenPerm[transition.token] = key.permAuthKeyId
	return transition
}

func (r *roleRegistry) promoteImmediate(method string, key roleSessionKey) roleTransition {
	r.mu.Lock()
	defer r.mu.Unlock()
	state := r.stateLocked(key.permAuthKeyId)
	if state.candidate != nil {
		delete(r.tokenPerm, state.candidate.token)
		state.candidate = nil
	}
	main := key
	state.main = &main
	return roleTransition{method: method, key: key, kind: selectorImmediate}
}

func (r *roleRegistry) drainPending(permAuthKeyId int64) pendingUpdatesState {
	if r == nil || permAuthKeyId == 0 {
		return pendingUpdatesState{}
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	state := r.byPerm[permAuthKeyId]
	if state == nil {
		return pendingUpdatesState{}
	}
	return state.drainPending()
}

func (r *roleRegistry) markPendingTooLong(permAuthKeyId int64) {
	if r == nil || permAuthKeyId == 0 {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	state := r.stateLocked(permAuthKeyId)
	state.pending.payloads = nil
	state.pending.tooLongPending = true
}

func (r *roleRegistry) completeSelector(token selectorToken, success bool) selectorCompletionResult {
	if r == nil || token == 0 {
		return selectorCompletionResult{stale: true}
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	permAuthKeyID, ok := r.tokenPerm[token]
	if !ok {
		return selectorCompletionResult{stale: true}
	}
	state := r.byPerm[permAuthKeyID]
	if state == nil || state.candidate == nil || state.candidate.token != token {
		delete(r.tokenPerm, token)
		return selectorCompletionResult{stale: true}
	}
	transition := *state.candidate
	state.candidate = nil
	delete(r.tokenPerm, token)
	if success {
		main := transition.key
		state.main = &main
		return selectorCompletionResult{flush: state.drainPending()}
	}
	state.pending.payloads = nil
	state.pending.tooLongPending = true
	return selectorCompletionResult{}
}

func (r *roleRegistry) isMain(key roleSessionKey) bool {
	if r == nil {
		return false
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	state := r.byPerm[key.permAuthKeyId]
	return state != nil && state.main != nil && *state.main == key
}

func (r *roleRegistry) unregisterSession(key roleSessionKey) {
	if r == nil || key.permAuthKeyId == 0 {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	state := r.byPerm[key.permAuthKeyId]
	if state == nil {
		return
	}
	if state.main != nil && *state.main == key {
		state.main = nil
	}
	if state.candidate != nil && state.candidate.key == key {
		delete(r.tokenPerm, state.candidate.token)
		state.candidate = nil
		state.pending.payloads = nil
		state.pending.tooLongPending = true
	}
}

func (r *roleRegistry) enqueueSelectorUpdate(permAuthKeyId int64, updates ...tg.UpdatesClazz) {
	if r == nil || permAuthKeyId == 0 {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	state := r.stateLocked(permAuthKeyId)
	if state.candidate == nil || state.pending.tooLongPending {
		return
	}
	var update tg.UpdatesClazz
	if len(updates) > 0 {
		update = updates[0]
	}
	if len(state.pending.payloads) < selectorQueueLimit {
		state.pending.payloads = append(state.pending.payloads, update)
		return
	}
	state.pending.payloads = nil
	state.pending.tooLongPending = true
}

func (r *roleRegistry) shouldWriteGenericUpdate(permAuthKeyId int64, updates tg.UpdatesClazz) bool {
	if r == nil || permAuthKeyId == 0 || updates == nil {
		return false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	state := r.byPerm[permAuthKeyId]
	if state == nil {
		return false
	}
	if state.candidate != nil {
		if !state.pending.tooLongPending {
			if len(state.pending.payloads) < selectorQueueLimit {
				state.pending.payloads = append(state.pending.payloads, updates)
			} else {
				state.pending.payloads = nil
				state.pending.tooLongPending = true
			}
		}
		return false
	}
	return state.main != nil
}

func (r *roleRegistry) snapshot(permAuthKeyId int64) roleSnapshot {
	if r == nil || permAuthKeyId == 0 {
		return roleSnapshot{}
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	state := r.byPerm[permAuthKeyId]
	if state == nil {
		return roleSnapshot{}
	}
	snap := roleSnapshot{
		queued:         len(state.pending.payloads),
		pendingTooLong: state.pending.tooLongPending,
	}
	if state.main != nil {
		main := *state.main
		snap.main = &main
	}
	if state.candidate != nil {
		snap.candidateKey = state.candidate.key
		snap.candidateToken = state.candidate.token
	}
	return snap
}

func (r *roleRegistry) stateLocked(permAuthKeyId int64) *rolePermState {
	state := r.byPerm[permAuthKeyId]
	if state == nil {
		state = &rolePermState{}
		r.byPerm[permAuthKeyId] = state
	}
	return state
}

func (s *rolePermState) drainPending() pendingUpdatesState {
	if s == nil {
		return pendingUpdatesState{}
	}
	pending := pendingUpdatesState{
		payloads:       append([]tg.UpdatesClazz(nil), s.pending.payloads...),
		tooLongPending: s.pending.tooLongPending,
	}
	s.pending.payloads = nil
	s.pending.tooLongPending = false
	return pending
}

func isNormalRoleKey(key roleSessionKey) bool {
	return key.permAuthKeyId != 0 &&
		key.authKeyId != 0 &&
		(key.authKeyType == tg.AuthKeyTypePerm || key.authKeyType == tg.AuthKeyTypeTemp)
}
