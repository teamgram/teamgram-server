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
	"fmt"

	"github.com/teamgram/proto/mtproto"
	sessionpb "github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

// SessionPushUpdatesData
// RPCPushClient is the client API for RPCPush service.
func (s *Service) SessionPushUpdatesData(ctx context.Context, r *sessionpb.TLSessionPushUpdatesData) (*mtproto.Bool, error) {
	logx.WithContext(ctx).Debugf("session.pushUpdatesData - request: %", r.DebugString())

	var (
		sessList *authSessions
		ok       bool
	)

	s.mu.RLock()
	sessList, ok = s.sessionsManager[r.AuthKeyId]
	s.mu.RUnlock()

	if !ok {
		err := fmt.Errorf("not found authKeyId(%d)", r.AuthKeyId)
		logx.WithContext(ctx).Errorf("session.pushUpdatesData - %v", err)
		return nil, err
	}
	sessList.syncDataArrived(r.Notification, &messageData{obj: r.Updates})

	return mtproto.BoolTrue, nil
}

// SessionPushSessionUpdatesData
// RPCPushClient is the client API for RPCPush service.
func (s *Service) SessionPushSessionUpdatesData(ctx context.Context, r *sessionpb.TLSessionPushSessionUpdatesData) (*mtproto.Bool, error) {
	logx.WithContext(ctx).Debugf("session.pushSessionUpdatesData - request: %", r.DebugString())

	var (
		sessList *authSessions
		ok       bool
	)

	s.mu.RLock()
	sessList, ok = s.sessionsManager[r.AuthKeyId]
	s.mu.RUnlock()

	if !ok {
		err := fmt.Errorf("not found authKeyId(%d)", r.AuthKeyId)
		logx.WithContext(ctx).Errorf("session.pushUpdatesData - %v", err)
		return nil, err
	}
	sessList.syncSessionDataArrived(r.SessionId, &messageData{obj: r.Updates})

	return mtproto.BoolTrue, nil
}

func (s *Service) SessionPushRpcResultData(ctx context.Context, r *sessionpb.TLSessionPushRpcResultData) (*mtproto.Bool, error) {
	logx.WithContext(ctx).Debugf("session.pushRpcResultData - request: %", r.DebugString())

	var (
		sessList *authSessions
		ok       bool
	)

	s.mu.RLock()
	sessList, ok = s.sessionsManager[r.AuthKeyId]
	s.mu.RUnlock()

	if !ok {
		err := fmt.Errorf("not found authKeyId(%d)", r.AuthKeyId)
		logx.WithContext(ctx).Errorf("session.pushRpcResultData - %v", err)
		return nil, err
	}
	sessList.syncRpcResultDataArrived(r.SessionId, r.ClientReqMsgId, r.RpcResultData)

	return mtproto.BoolTrue, nil
}
