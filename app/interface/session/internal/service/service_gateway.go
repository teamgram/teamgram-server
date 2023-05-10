// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	sessionpb "github.com/teamgram/teamgram-server/app/interface/session/session"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"

	"github.com/zeromicro/go-zero/core/logx"
)

func (s *Service) SessionCreateSession(ctx context.Context, r *sessionpb.TLSessionCreateSession) (*mtproto.Bool, error) {
	logx.WithContext(ctx).Debugf("session.createSession - request: %s", r.DebugString())

	var (
		sessList *authSessions
		ok       bool
		c        = r.GetClient()
		err      error
	)

	s.mu.RLock()
	sessList, ok = s.sessionsManager[c.GetAuthKeyId()]
	s.mu.RUnlock()
	if !ok {
		sessList, err = newAuthSessions(c.GetAuthKeyId(), s)
		if err != nil {
			logx.WithContext(ctx).Errorf("session.createSession - newAuthSessions error: %v", err)
			return nil, err
		}
		s.mu.Lock()
		s.sessionsManager[c.GetAuthKeyId()] = sessList
		s.mu.Unlock()
	}
	sessList.sessionClientNew(c.GetServerId(), c.GetSessionId())

	return mtproto.BoolTrue, nil
}

func (s *Service) SessionCloseSession(ctx context.Context, r *sessionpb.TLSessionCloseSession) (*mtproto.Bool, error) {
	logx.WithContext(ctx).Debugf("session.closeSession - request: %s", r.DebugString())

	var (
		sessList *authSessions
		ok       bool
		c        = r.GetClient()
	)

	s.mu.RLock()
	sessList, ok = s.sessionsManager[c.GetAuthKeyId()]
	s.mu.RUnlock()
	if !ok {
		logx.WithContext(ctx).Errorf("session.closeSession - not found sessList by keyId: %d", c.GetAuthKeyId())
	} else {
		sessList.sessionClientClosed(c.GetServerId(), c.GetSessionId())
	}

	return mtproto.BoolTrue, nil
}

func (s *Service) SessionSendDataToSession(ctx context.Context, r *sessionpb.TLSessionSendDataToSession) (res *mtproto.Bool, err error) {
	logx.WithContext(ctx).Debugf("session.sendDataToSession - request: {server_id: %s, conn_type: %d, auth_key_id: %d, session_id: %s, client_ip: %s, quick_ack: %d, salt: %d, payload: %d}",
		r.GetData().GetServerId(),
		r.GetData().GetConnType(),
		r.GetData().GetAuthKeyId(),
		r.GetData().GetSessionId(),
		r.GetData().GetClientIp(),
		r.GetData().GetQuickAck(),
		r.GetData().GetSalt(),
		len(r.GetData().GetPayload()))

	var (
		sessList *authSessions
		ok       bool
		data     = r.GetData()
	)

	s.mu.RLock()
	sessList, ok = s.sessionsManager[data.GetAuthKeyId()]
	s.mu.RUnlock()

	if !ok {
		sessList, err = newAuthSessions(data.GetAuthKeyId(), s)
		if err != nil {
			logx.WithContext(ctx).Errorf("session.createSession - newAuthSessions error: %v", err)
			return nil, err
		}

		s.mu.Lock()
		s.sessionsManager[data.GetAuthKeyId()] = sessList
		s.mu.Unlock()
	}

	sessList.sessionDataArrived(data.GetServerId(), data.GetClientIp(), data.GetSessionId(), data.GetSalt(), data.GetPayload())

	return mtproto.BoolTrue, nil
}

func (s *Service) SessionSendHttpDataToSession(ctx context.Context, r *sessionpb.TLSessionSendHttpDataToSession) (res *sessionpb.HttpSessionData, err error) {
	logx.WithContext(ctx).Errorf("not impl session.sendHttpDataToSession")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SessionQueryAuthKey(ctx context.Context, r *sessionpb.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	logx.WithContext(ctx).Debugf("session.queryAuthKey - request: %s", r.DebugString())

	key, err := s.Dao.AuthsessionClient.AuthsessionQueryAuthKey(ctx, &authsession.TLAuthsessionQueryAuthKey{
		AuthKeyId: r.AuthKeyId,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("session.queryAuthKey - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLAuthKeyInfo(&mtproto.AuthKeyInfo{
		AuthKeyId:          key.AuthKeyId,
		AuthKey:            key.AuthKey,
		AuthKeyType:        key.AuthKeyType,
		PermAuthKeyId:      key.PermAuthKeyId,
		TempAuthKeyId:      key.TempAuthKeyId,
		MediaTempAuthKeyId: key.MediaTempAuthKeyId,
	}).To_AuthKeyInfo(), nil
}

func (s *Service) SessionSetAuthKey(ctx context.Context, r *sessionpb.TLSessionSetAuthKey) (*mtproto.Bool, error) {
	logx.WithContext(ctx).Debugf("session.setAuthKey - request: %s", r.DebugString())

	if r.AuthKey == nil {
		logx.WithContext(ctx).Errorf("session.setAuthKey error: auth_key is nil")
		return nil, mtproto.ErrInputRequestInvalid
	}

	rV, err := s.Dao.AuthsessionClient.AuthsessionSetAuthKey(
		ctx,
		&authsession.TLAuthsessionSetAuthKey{
			AuthKey: mtproto.MakeTLAuthKeyInfo(&mtproto.AuthKeyInfo{
				AuthKeyId:          r.AuthKey.AuthKeyId,
				AuthKey:            r.AuthKey.AuthKey,
				AuthKeyType:        r.AuthKey.AuthKeyType,
				PermAuthKeyId:      r.AuthKey.PermAuthKeyId,
				TempAuthKeyId:      r.AuthKey.TempAuthKeyId,
				MediaTempAuthKeyId: r.AuthKey.MediaTempAuthKeyId,
			}).To_AuthKeyInfo(),
			FutureSalt: r.FutureSalt,
			ExpiresIn:  r.ExpiresIn,
		})
	if err != nil {
		logx.WithContext(ctx).Errorf("session.setAuthKey - error: %v", err)
		return nil, err
	}

	return rV, nil
}
