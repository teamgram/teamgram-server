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
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgDeleteHistory
// msg.deleteHistory flags:# user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:flags.0?true revoke:flags.1?true max_id:int = messages.AffectedHistory;
func (c *MsgCore) MsgDeleteHistory(in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error) {
	if in == nil || in.UserId <= 0 || in.PeerId <= 0 || in.MaxId < 0 {
		return nil, fmt.Errorf("%w: invalid delete history request", msg.ErrSendStateConflict)
	}
	if in.PeerType != payload.PeerTypeUser && in.PeerType != payload.PeerTypeChat {
		return nil, fmt.Errorf("%w: unsupported delete history peer type=%d", msg.ErrSendStateConflict, in.PeerType)
	}
	if in.PeerType == payload.PeerTypeChat {
		action := chatpb.MessageActionDeleteLocal
		if in.Revoke {
			action = chatpb.MessageActionDeleteRevoke
		}
		if err := c.checkChatAction(in.UserId, in.PeerId, action, ""); err != nil {
			return nil, err
		}
	}
	maxPeerSeq := int64(0)
	if in.MaxId > 0 {
		resolved, err := c.svcCtx.Repo.ResolveMessageID(c.ctx, in.UserId, in.PeerType, in.PeerId, int64(in.MaxId))
		if err != nil {
			return nil, err
		}
		if resolved == nil {
			return tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{}).ToMessagesAffectedHistory(), nil
		}
		maxPeerSeq = resolved.PeerSeq
	}
	date := int32(time.Now().Unix())
	body, err := buildDeleteHistoryPayload(in.UserId, in.UserId, in.PeerType, in.PeerId, maxPeerSeq, date, in.JustClear, in.Revoke)
	if err != nil {
		return nil, err
	}
	effects, err := c.buildDeleteHistoryChatReceiverEffects(in, maxPeerSeq, date)
	if err != nil {
		return nil, err
	}
	result, err := c.processUserDialogOperation(in.UserId, in.AuthKeyId, in.PeerType, in.PeerId, deleteHistoryOperationID(in.UserId, in.PeerId, maxPeerSeq, in.JustClear, in.Revoke, in.AuthKeyId), body, effects)
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(result.Pts, "pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{Pts: pts, PtsCount: result.PtsCount, Offset: 0}).ToMessagesAffectedHistory(), nil
}

func deleteHistoryOperationID(userID int64, peerID int64, maxPeerSeq int64, justClear bool, revoke bool, authKeyID int64) string {
	return fmt.Sprintf("v2:dialog:delete_history:user:%d:peer:%d:max_peer_seq:%d:clear:%t:revoke:%t:auth:%d", userID, peerID, maxPeerSeq, justClear, revoke, authKeyID)
}

func buildDeleteHistoryPayload(fromUserID int64, toUserID int64, peerType int32, peerID int64, maxPeerSeq int64, date int32, justClear bool, revoke bool) ([]byte, error) {
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:    payload.MessageOperationSchemaVersion,
		OperationKind:    payload.OperationKindDeleteHistory,
		PeerType:         peerType,
		PeerID:           peerID,
		PeerSeq:          maxPeerSeq,
		FromUserID:       fromUserID,
		ToUserID:         toUserID,
		Date:             date,
		DeleteMaxPeerSeq: maxPeerSeq,
		JustClear:        justClear,
		Revoke:           revoke,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: marshal delete history operation user_id=%d peer_id=%d", msg.ErrMsgStorage, fromUserID, peerID)
	}
	return body, nil
}

func (c *MsgCore) buildDeleteHistoryChatReceiverEffects(in *msg.TLMsgDeleteHistory, maxPeerSeq int64, date int32) ([]OperationEnvelope, error) {
	if in.PeerType != payload.PeerTypeChat || !in.Revoke {
		return nil, nil
	}
	receiverIDs, err := c.activeChatReceiverIDs(in.PeerId, in.UserId)
	if err != nil {
		return nil, err
	}
	effects := make([]OperationEnvelope, 0, len(receiverIDs))
	for _, receiverUserID := range receiverIDs {
		peerBody, err := buildDeleteHistoryPayload(in.UserId, receiverUserID, payload.PeerTypeChat, in.PeerId, maxPeerSeq, date, in.JustClear, in.Revoke)
		if err != nil {
			return nil, err
		}
		effects = append(effects, OperationEnvelope{
			UserID:               receiverUserID,
			OperationID:          deleteHistoryOperationID(receiverUserID, in.PeerId, maxPeerSeq, in.JustClear, in.Revoke, 0),
			OpType:               payload.OpTypeSendMessage,
			OperationKind:        payload.OperationKindDeleteHistory,
			ActorUserID:          in.UserId,
			PeerType:             payload.PeerTypeChat,
			PeerID:               in.PeerId,
			CanonicalPeerSeq:     int64Ptr(maxPeerSeq),
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         payload.PayloadCodecJSON,
			PayloadHash:          payload.HashBytes(peerBody),
			Payload:              peerBody,
			DeliveryPolicy:       DeliveryPolicyDurableAsync,
		})
	}
	return effects, nil
}
