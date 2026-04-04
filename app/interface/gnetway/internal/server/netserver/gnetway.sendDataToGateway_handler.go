// Copyright (c) 2021-present,  Teamgooo Studio (https://teamgram.io).
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

package netserver

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/logx"
)

// GnetwaySendDataToGateway
// gnetway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (s *Server) GnetwaySendDataToGateway(ctx context.Context, in *gnetway.TLGnetwaySendDataToGateway) (reply *tg.Bool, err error) {
	logx.WithContext(ctx).Infof("ReceiveData - request: {kId: %d, sessionId: %d, payloadLen: %d}", in.AuthKeyId, in.SessionId, len(in.Payload))

	authKey, connIdList := s.authSessionMgr.FoundSessionConnId(in.AuthKeyId, in.SessionId)
	if len(connIdList) == 0 {
		logx.WithContext(ctx).Errorf("ReceiveData - not found connId - keyId: %d, sessionId: %d", in.AuthKeyId, in.SessionId)
		return tg.BoolFalse, nil
	} else {
		logx.WithContext(ctx).Debugf("found: {k: %v, idList: %v}", in.AuthKeyId, connIdList)
	}

	ctx = contextx.ValueOnlyFrom(ctx)
	msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(in.Payload)
	x := bin.NewEncoder()
	defer x.End()

	// x := mtproto.NewEncodeBuf(8 + len(msgKey) + len(mtpRawData))
	x.PutInt64(authKey.AuthKeyId())
	x.Put(msgKey)
	x.Put(mtpRawData)
	msg := x.Bytes()

	delivered := false
	for _, connId := range connIdList {
		// Direct synchronous connection access
		c, ok := s.connMgr.get(connId)
		if !ok || c == nil {
			logx.WithContext(ctx).Errorf("invalid state - conn(%d) is nil or not found", connId)
			continue
		}

		connAuthKey := c.getAuthKey()
		if connAuthKey == nil {
			logx.WithContext(ctx).Errorf("invalid state - conn(%d) auth key is nil", connId)
			continue
		}

		if in.AuthKeyId != connAuthKey.AuthKeyId() {
			logx.WithContext(ctx).Errorf("invalid state - conn(%d) c.keyId(%d) != in.keyId(%d)", connId, connAuthKey.AuthKeyId(), in.AuthKeyId)
			continue
		}

		err2 := s.writeToConnection(c, msg)
		if err2 != nil {
			logx.WithContext(ctx).Errorf("sendToClient error: %v", err2)
		} else {
			delivered = true
			logx.WithContext(ctx).Debugf("sendToConn: %v", connId)
		}
	}

	return tg.ToBool(delivered), nil
}
