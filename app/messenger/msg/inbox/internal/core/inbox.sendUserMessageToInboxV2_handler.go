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
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/inbox"
	synctypes "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// InboxSendUserMessageToInboxV2
// inbox.sendUserMessageToInboxV2 flags:# user_id:long out:flags.0?true from_id:long from_auth_keyId:long peer_type:int peer_id:long box_list:Vector<MessageBox> users:flags.1?Vector<User> chats:flags.2?Vector<Chat> layer:flags.3?int server_id:flags.4?string session_id:flags.5?long client_req_msg_id:flags.6?long auth_key_id:flags.7?long= Void;
func (c *InboxCore) InboxSendUserMessageToInboxV2(in *inbox.TLInboxSendUserMessageToInboxV2) (*tg.Void, error) {
	// Push updates to the recipient's other sessions via sync.
	if c.svcCtx != nil && c.svcCtx.SyncClient != nil && len(in.BoxList) > 0 {
		updates := tg.MakeTLUpdates(&tg.TLUpdates{
			Updates: buildUpdateNewMessages(in.BoxList),
			Users:   in.Users,
			Chats:   in.Chats,
			Date:    0,
			Seq:     0,
		})
		_, err := c.svcCtx.SyncClient.SyncPushUpdates(c.ctx, &synctypes.TLSyncPushUpdates{
			UserId:  in.UserId,
			Updates: updates,
		})
		if err != nil {
			c.Logger.Errorf("inbox.sendUserMessageToInboxV2 - sync push error: %v", err)
		}
	}
	return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
}

func buildUpdateNewMessages(boxes []tg.MessageBoxClazz) []tg.UpdateClazz {
	updates := make([]tg.UpdateClazz, 0, len(boxes))
	for _, box := range boxes {
		if box == nil {
			continue
		}
		var msg tg.MessageClazz
		if box.Message != nil {
			msg = box.Message
		}
		updates = append(updates, tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message:  msg,
			Pts:      box.Pts,
			PtsCount: box.PtsCount,
		}))
	}
	return updates
}
