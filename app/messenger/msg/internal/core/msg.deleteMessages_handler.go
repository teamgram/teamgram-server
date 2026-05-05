// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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
	"encoding/json"
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgDeleteMessages
// msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;
func (c *MsgCore) MsgDeleteMessages(in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	if in == nil || in.UserId <= 0 || in.PeerId <= 0 || len(in.Id) == 0 {
		return nil, fmt.Errorf("%w: invalid delete messages request", msg.ErrSendStateConflict)
	}
	if in.PeerType != payload.PeerTypeUser {
		return nil, fmt.Errorf("%w: delete messages first slice only supports user peer", msg.ErrSendStateConflict)
	}
	if c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	deletePeerSeqs := make([]int64, 0, len(in.Id))
	for _, id := range in.Id {
		if id <= 0 {
			return nil, fmt.Errorf("%w: invalid message id", msg.ErrSendStateConflict)
		}
		deletePeerSeqs = append(deletePeerSeqs, int64(id))
	}
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:  payload.MessageOperationSchemaVersion,
		OperationKind:  payload.OperationKindDeleteMessages,
		PeerType:       in.PeerType,
		PeerID:         in.PeerId,
		FromUserID:     in.UserId,
		ToUserID:       in.UserId,
		Date:           int32(time.Now().Unix()),
		DeletePeerSeqs: deletePeerSeqs,
		Revoke:         in.Revoke,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: marshal delete messages operation user_id=%d peer_id=%d", msg.ErrMsgStorage, in.UserId, in.PeerId)
	}
	result, err := c.processUserDialogOperation(in.UserId, in.AuthKeyId, in.PeerType, in.PeerId, deleteMessagesOperationID(in.UserId, in.PeerId, in.Id, in.Revoke, in.AuthKeyId), body)
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(result.Pts, "pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{Pts: pts, PtsCount: result.PtsCount}).ToMessagesAffectedMessages(), nil
}

func deleteMessagesOperationID(userID int64, peerID int64, ids []int32, revoke bool, authKeyID int64) string {
	return fmt.Sprintf("v1:dialog:delete_messages:user:%d:peer:%d:ids:%v:revoke:%t:auth:%d", userID, peerID, ids, revoke, authKeyID)
}
