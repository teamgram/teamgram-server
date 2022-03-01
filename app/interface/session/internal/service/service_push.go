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

	"github.com/zeromicro/go-zero/core/logx"
)

// SessionPushUpdatesData
// RPCPushClient is the client API for RPCPush service.
func (s *Service) SessionPushUpdatesData(ctx context.Context, r *sessionpb.TLSessionPushUpdatesData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		logx.WithContext(ctx).Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncDataArrived(r.Notification, &messageData{obj: r.Updates})
	return mtproto.ToBool(true), nil
}

// SessionPushSessionUpdatesData
// RPCPushClient is the client API for RPCPush service.
func (s *Service) SessionPushSessionUpdatesData(ctx context.Context, r *sessionpb.TLSessionPushSessionUpdatesData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		logx.WithContext(ctx).Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncSessionDataArrived(r.SessionId, &messageData{obj: r.Updates})
	return mtproto.ToBool(true), nil
}

func (s *Service) SessionPushRpcResultData(ctx context.Context, r *sessionpb.TLSessionPushRpcResultData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		logx.WithContext(ctx).Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncRpcResultDataArrived(r.SessionId, r.ClientReqMsgId, r.RpcResultData)
	return mtproto.ToBool(true), nil
}
