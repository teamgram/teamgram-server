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

package server

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gateway/gateway"

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
	logger := logx.WithContext(ctx)
	logger.Debugf("ReceiveData - request: {kId: %d, sessionId: %d, payloadLen: %d}", in.AuthKeyId, in.SessionId, len(in.Payload))

	var (
		authKey *authKeyUtil
	)

	// TODO(@benqi): 并发问题
	authKey, connIdList := s.authSessionMgr.FoundSessionConnIdList(in.AuthKeyId, in.SessionId)
	if connIdList == nil {
		logger.Errorf("ReceiveData - not found connIdList - keyId: %d, sessionId: %d", in.AuthKeyId, in.SessionId)
		return mtproto.BoolFalse, nil
	}

	//msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(in.Payload)
	//x := mtproto.NewEncodeBuf(8 + len(msgKey) + len(mtpRawData))
	//x.Long(authKey.AuthKeyId())
	//x.Bytes(msgKey)
	//x.Bytes(mtpRawData)
	//msg := &mtproto.MTPRawMessage{Payload: x.GetBuf()}

	//for _, connId := range connIdList {
	//	s.svr.Trigger(connId, func(c gnet.Conn) {
	//		if err := c.UnThreadSafeWrite(msg); err != nil {
	//			logx.Errorf("sendToClient error: %v", err)
	//		}
	//	})
	//}

	for _, connId := range connIdList {
		logger.Debugf("[keyId: %d, sessionId: %d]: %v", in.AuthKeyId, in.SessionId, connId)
		conn2 := s.server.GetConnection(connId)
		if conn2 != nil {
			ctx2, _ := conn2.Context.(*connContext)
			authKey = ctx2.getAuthKey(in.AuthKeyId)
			if authKey == nil {
				logger.Errorf("invalid authKeyId, authKeyId = %d", in.AuthKeyId)
				continue
			}
			if ctx2.isHttp {
				// isHttp = true
				if !ctx2.canSend {
					continue
				}
			}
			// conn = conn2
			// break
			if err = s.SendToClient(conn2, authKey, in.Payload); err == nil {
				logger.Infof("ReceiveData -  result: {auth_key_id = %d, session_id = %d, conn = %s}",
					in.AuthKeyId,
					in.SessionId,
					conn2)

				if ctx2.isHttp {
					s.authSessionMgr.PushBackHttpData(in.AuthKeyId, in.SessionId, in.Payload)
				}
				return mtproto.ToBool(true), nil
			} else {
				logger.Errorf("ReceiveData - sendToClient error (%v), auth_key_id = %d, session_id = %d, conn_id_list = %v",
					err,
					in.AuthKeyId,
					in.SessionId,
					connIdList)
			}
		}
	}

	return mtproto.BoolTrue, nil
}
