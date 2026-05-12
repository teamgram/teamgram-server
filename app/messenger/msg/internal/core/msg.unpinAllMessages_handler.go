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
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgUnpinAllMessages
// msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (c *MsgCore) MsgUnpinAllMessages(in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error) {
	if in == nil || in.UserId <= 0 || in.PeerId <= 0 {
		return nil, fmt.Errorf("%w: invalid unpin all request", msg.ErrSendStateConflict)
	}
	if in.PeerType != payload.PeerTypeUser && in.PeerType != payload.PeerTypeChat {
		return nil, fmt.Errorf("%w: unsupported unpin all peer type=%d", msg.ErrSendStateConflict, in.PeerType)
	}
	if in.PeerType == payload.PeerTypeChat {
		if err := c.checkChatAction(in.UserId, in.PeerId, chatpb.MessageActionUnpinAll, ""); err != nil {
			return nil, err
		}
	}
	pin := &msg.TLMsgUpdatePinnedMessage{
		UserId:    in.UserId,
		AuthKeyId: in.AuthKeyId,
		Unpin:     true,
		PeerType:  in.PeerType,
		PeerId:    in.PeerId,
		Id:        0,
	}
	body, hashBytes, err := buildPinnedMessageOperation(pin, in.UserId, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	effects, err := c.buildUpdatePinnedChatReceiverEffects(pin, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	authKeyID := in.AuthKeyId
	result, err := c.dispatchRequesterSync(OperationEnvelope{
		UserID:               in.UserId,
		OperationID:          updatePinnedOperationID(in.UserId, in.PeerId, 0, true, in.AuthKeyId),
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindUpdatePinnedMessage,
		ActorUserID:          in.UserId,
		AuthKeyID:            &authKeyID,
		AuthKeyIDExclude:     &authKeyID,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		PayloadCodec:         payload.PayloadCodecJSON,
		PayloadHash:          hashBytes,
		Payload:              body,
		DeliveryPolicy:       DeliveryPolicyRequesterSync,
	}, effects)
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(result.Pts, "pts")
	if err != nil {
		return nil, err
	}

	return tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{
		Pts:      pts,
		PtsCount: result.PtsCount,
	}).ToMessagesAffectedHistory(), nil
}
