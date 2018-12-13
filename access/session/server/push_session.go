// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package server

import (
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/golang/glog"
	"reflect"
)

type pushSession struct {
	*session
}

var tooLong = mtproto.NewTLUpdatesTooLong()

func (c *pushSession) onMessageData(id ClientConnID, cntl *zrpc.ZRpcController, salt int64, msg *mtproto.TLMessage2) {
	glog.Infof("onMessageData - pushSession: {id: %s, cntl: %s, obj: %s}",
		id,
		cntl,
		reflect.TypeOf(msg.Object))

	c.session.processMessageData(id, cntl, salt, msg, func(sessMsg *mtproto.TLMessage2) {
	})

	if len(c.pendingMessages) > 0 {
		c.sendPendingMessagesToClient(id, cntl, c.pendingMessages)
		c.pendingMessages = []*pendingMessage{}
	}
}

func (c *pushSession) onSyncData(cntl *zrpc.ZRpcController) {
	glog.Infof("onSyncData - sendPending {sess: {%s}, pushObj: {updatesTooLong}}", c)

	if c.sessionOnline() {
		syncMessage := &pendingMessage{
			messageId: mtproto.GenerateMessageId(),
			confirm:   true,
			tl:        tooLong,
		}
		c.syncMessages = append(c.syncMessages, syncMessage)
		c.sendPendingMessagesToClient(c.connId, cntl, c.syncMessages)
		c.syncMessages = []*pendingMessage{}
	}
}
