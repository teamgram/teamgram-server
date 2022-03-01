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

package server

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/app/interface/gateway/internal/server/codec"

	"github.com/panjf2000/gnet"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
   bool error = false;
   if (length <= 24 + 32) {
       int32_t code = data->readInt32(&error);
       if (code == 0) {
           if (LOGS_ENABLED) DEBUG_D("mtproto noop");
       } else if (code == -1) {
           int32_t ackId = data->readInt32(&error);
           if (!error) {
               onConnectionQuickAckReceived(connection, ackId & (~(1 << 31)));
           }
       } else {
           Datacenter *datacenter = connection->getDatacenter();
           if (LOGS_ENABLED) DEBUG_W("mtproto error = %d", code);
           if (code == -444 && connection->getConnectionType() == ConnectionTypeGeneric && !proxyAddress.empty() && !proxySecret.empty()) {
               if (delegate != nullptr) {
                   delegate->onProxyError(instanceNum);
               }
           } else if (code == -404 && (datacenter->isCdnDatacenter || PFS_ENABLED)) {
               if (!datacenter->isHandshaking(connection->isMediaConnection)) {
                   datacenter->clearAuthKey(connection->isMediaConnection ? HandshakeTypeMediaTemp : HandshakeTypeTemp);
                   datacenter->beginHandshake(connection->isMediaConnection ? HandshakeTypeMediaTemp : HandshakeTypeTemp, true);
                   if (LOGS_ENABLED) DEBUG_D("connection(%p, account%u, dc%u, type %d) reset auth key due to -404 error", connection, instanceNum, datacenter->getDatacenterId(), connection->getConnectionType());
               }
           } else {
               connection->reconnect();
           }
       }
       return;
   }
*/

// GatewaySendDataToGateway
// gateway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (s *Server) GatewaySendDataToGateway(ctx context.Context, in *gateway.TLGatewaySendDataToGateway) (reply *mtproto.Bool, err error) {
	logx.Infof("ReceiveData - request: {kId: %d, sessionId: %d, payloadLen: %d}", in.AuthKeyId, in.SessionId, len(in.Payload))

	var (
		authKey *authKeyUtil
	)

	// TODO(@benqi): 并发问题
	authKey, connIdList := s.authSessionMgr.FoundSessionConnIdList(in.AuthKeyId, in.SessionId)
	if connIdList == nil {
		logx.Errorf("ReceiveData - not found connIdList - keyId: %d, sessionId: %d", in.AuthKeyId, in.SessionId)
		return mtproto.BoolFalse, nil
	}

	msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(in.Payload)
	x := mtproto.NewEncodeBuf(8 + len(msgKey) + len(mtpRawData))
	x.Long(authKey.AuthKeyId())
	x.Bytes(msgKey)
	x.Bytes(mtpRawData)
	msg := &codec.MTPRawMessage{Payload: x.GetBuf()}

	for _, connId := range connIdList {
		s.svr.Trigger(connId, func(c gnet.Conn) {
			if err := c.UnThreadSafeWrite(msg); err != nil {
				logx.Errorf("sendToClient error: %v", err)
			}
		})
	}
	return mtproto.BoolTrue, nil
}
