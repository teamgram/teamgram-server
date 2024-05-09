// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package sess

import (
	"sync"

	"github.com/teamgram/teamgram-server/app/interface/session/internal/dao"
)

type AuthSessionsManager struct {
	mu      sync.Mutex
	authMgr map[int64]*AuthSessions
	*dao.Dao
}

func NewAuthSessionsManager(d *dao.Dao) *AuthSessionsManager {
	return &AuthSessionsManager{
		authMgr: make(map[int64]*AuthSessions),
		Dao:     d,
	}
}

func (s *AuthSessionsManager) GetOrCreateAuthSessions(authKeyId int64) (v *AuthSessions, isNew bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, isNew = s.authMgr[authKeyId]
	if !isNew {
		v = newAuthSessions(authKeyId, s)
		s.authMgr[authKeyId] = v
	}

	return v, !isNew
}

func (s *AuthSessionsManager) GetAuthSessions(authKeyId int64) *AuthSessions {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.authMgr[authKeyId]
	if ok {
		return v
	}

	return nil
}

func (s *AuthSessionsManager) DeleteByAuthKeyId(authKeyId int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sessList, ok := s.authMgr[authKeyId]; ok {
		sessList.Stop()
		delete(s.authMgr, authKeyId)
	}
}
