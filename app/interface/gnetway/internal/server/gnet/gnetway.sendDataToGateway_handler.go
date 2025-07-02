// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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

package gnet

import (
	"context"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"

	"github.com/panjf2000/gnet/v2"
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

	for _, connId := range connIdList {
		s.eng.Trigger(connId, func(c gnet.Conn) {
			connCtx, _ := c.Context().(*connContext)
			if connCtx == nil {
				logx.WithContext(ctx).Errorf("invalid state - conn(%s) Context() is nil", c)
				return
			}

			if in.AuthKeyId != connCtx.getAuthKey().AuthKeyId() {
				logx.WithContext(ctx).Errorf("invalid state - conn(%s) c.keyId(%d) != in.keyId(%d) is nil", authKey.AuthKeyId(), in.AuthKeyId)
				return
			}

			err2 := UnThreadSafeWrite(c, msg)
			if err2 != nil {
				logx.WithContext(ctx).Errorf("sendToClient error: %v", err2)
			} else {
				logx.WithContext(ctx).Debugf("sendToConn: %v", connId)
			}
		})
	}

	return tg.BoolTrue, nil
}
