// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// MsgPushUserMessage
// msg.pushUserMessage user_id:long auth_key_id:long peer_type:int peer_id:long push_type:int message:OutboxMessage = Bool;
func (c *MsgCore) MsgPushUserMessage(in *msg.TLMsgPushUserMessage) (*tg.Bool, error) {
	if in.PeerType != tg.PEER_USER && in.PeerType != tg.PEER_SELF {
		c.Logger.Errorf("msg.pushUserMessage - unsupported peer_type(%d), peer_id(%d)", in.PeerType, in.PeerId)
		return tg.BoolFalse, nil
	}

	// TODO: route user-directed messages through the real msg/inbox pipeline.
	return tg.BoolTrue, nil
}
