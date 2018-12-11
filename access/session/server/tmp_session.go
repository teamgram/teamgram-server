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
)

type tempSession struct {
	*session
}

func (c *tempSession) onMessageData(id ClientConnID, cntl *zrpc.ZRpcController, salt int64, msg *mtproto.TLMessage2) {
	c.session.processMessageData(id, cntl, salt, msg, func(sessMsg *mtproto.TLMessage2) {
	})

	if len(c.pendingMessages) > 0 {
		c.sendPendingMessagesToClient(id, cntl, c.pendingMessages)
		c.pendingMessages = []*pendingMessage{}
	}
}

