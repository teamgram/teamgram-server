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
	var (
		sessList *authSessions
		ok       bool
		c        = r.GetClient()
		err      error
	)

	logx.WithContext(ctx).Infof("createSession - request: %s", r.DebugString())

	s.mu.Lock()
	defer s.mu.Unlock()

	if sessList, ok = s.sessionsManager[c.GetAuthKeyId()]; !ok {
		sessList, err = newAuthSessions(c.GetAuthKeyId(), s)
		if err != nil {
			return nil, err
		}
		s.sessionsManager[c.GetAuthKeyId()] = sessList
	}
	sessList.sessionClientNew(c.GetServerId(), c.GetSessionId())

	return mtproto.BoolTrue, nil
}

func (s *Service) SessionCloseSession(ctx context.Context, r *sessionpb.TLSessionCloseSession) (*mtproto.Bool, error) {
	var (
		sessList *authSessions
		ok       bool
		c        = r.GetClient()
		logger   = logx.WithContext(ctx)
	)

	logx.WithContext(ctx).Infof("closeSession - request: %s", r.DebugString())

	s.mu.RLock()
	defer s.mu.RUnlock()

	if sessList, ok = s.sessionsManager[c.GetAuthKeyId()]; !ok {
		logger.Errorf("not found sessList by keyId: %d", c.GetAuthKeyId())
	} else {
		sessList.sessionClientClosed(c.GetServerId(), c.GetSessionId())
	}

	return mtproto.BoolTrue, nil
}

func (s *Service) SessionSendDataToSession(ctx context.Context, r *sessionpb.TLSessionSendDataToSession) (res *mtproto.Bool, err error) {
	var (
		sessList *authSessions
		ok       bool
		data     = r.GetData()
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	if sessList, ok = s.sessionsManager[data.GetAuthKeyId()]; !ok {
		sessList, err = newAuthSessions(data.GetAuthKeyId(), s)
		if err != nil {
			return
		}

		s.sessionsManager[data.GetAuthKeyId()] = sessList
	}
	sessList.sessionDataArrived(data.GetServerId(), data.GetClientIp(), data.GetSessionId(), data.GetSalt(), data.GetPayload())
	res = mtproto.ToBool(true)
	return
}

func (s *Service) SessionSendHttpDataToSession(ctx context.Context, r *sessionpb.TLSessionSendHttpDataToSession) (res *sessionpb.HttpSessionData, err error) {
	//respChan := make(chan interface{}, 2)
	//
	//s.mu.Lock()
	//var (
	//	sessList *authSessions
	//	ok       bool
	//)
	//
	//if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
	//	sessList, err = newAuthSessions(r.AuthKeyId, s)
	//	if err != nil {
	//		s.mu.Unlock()
	//		return
	//	}
	//	s.sessionsManager[r.AuthKeyId] = sessList
	//}
	//sessList.sessionHttpDataArrived(r.ServerId, r.ClientIp, r.SessionId, r.Salt, r.Payload, respChan)
	//s.mu.Unlock()
	//
	//timer := time.NewTimer(60 * time.Second)
	//select {
	//case rc := <-respChan:
	//	if rc != nil {
	//		if payload, ok := rc.([]byte); ok {
	//			log.Debugf("recv http data: %d", len(payload))
	//			res = &_.SessionData{
	//				Payload: payload,
	//			}
	//			return
	//		}
	//	}
	//	err = mtproto.ErrInternelServerError
	//	log.Errorf("sendSyncSessionData - error: %v", err)
	//case <-timer.C:
	//	err = mtproto.ErrTimeOut503
	//	log.Errorf("sendSyncSessionData - error: %v", err)
	//}
	//

	logx.WithContext(ctx).Errorf("not impl session.sendHttpDataToSession")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SessionQueryAuthKey(ctx context.Context, r *sessionpb.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	key, err := s.Dao.AuthsessionClient.AuthsessionQueryAuthKey(ctx, &authsession.TLAuthsessionQueryAuthKey{
		AuthKeyId: r.AuthKeyId,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("queryAuthKey error: %v", err)
		return nil, err
	}

	return &mtproto.AuthKeyInfo{
		AuthKeyId:          key.AuthKeyId,
		AuthKey:            key.AuthKey,
		AuthKeyType:        key.AuthKeyType,
		PermAuthKeyId:      key.PermAuthKeyId,
		TempAuthKeyId:      key.TempAuthKeyId,
		MediaTempAuthKeyId: key.MediaTempAuthKeyId,
	}, nil
}

func (s *Service) SessionSetAuthKey(ctx context.Context, r *sessionpb.TLSessionSetAuthKey) (*mtproto.Bool, error) {
	if r.AuthKey == nil {
		logx.WithContext(ctx).Errorf("setAuthKey error: auth_key is nil")
		return nil, mtproto.ErrInputRequestInvalid
	}
	return s.Dao.AuthsessionClient.AuthsessionSetAuthKey(ctx, &authsession.TLAuthsessionSetAuthKey{
		AuthKey: &mtproto.AuthKeyInfo{
			AuthKeyId:          r.AuthKey.AuthKeyId,
			AuthKey:            r.AuthKey.AuthKey,
			AuthKeyType:        r.AuthKey.AuthKeyType,
			PermAuthKeyId:      r.AuthKey.PermAuthKeyId,
			TempAuthKeyId:      r.AuthKey.TempAuthKeyId,
			MediaTempAuthKeyId: r.AuthKey.MediaTempAuthKeyId,
		},
		FutureSalt: r.FutureSalt,
		ExpiresIn:  r.ExpiresIn,
	})
}
