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

// MsgReadHistoryV2
// msg.readHistoryV2 user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (c *MsgCore) MsgReadHistoryV2(in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
	if in == nil {
		return nil, fmt.Errorf("%w: missing read history request", msg.ErrSendStateConflict)
	}
	if in.UserId <= 0 || in.PeerId <= 0 || in.MaxId < 0 {
		return nil, fmt.Errorf("%w: invalid read history request", msg.ErrSendStateConflict)
	}
	if in.PeerType != payload.PeerTypeUser && in.PeerType != payload.PeerTypeChat {
		return nil, fmt.Errorf("%w: unsupported read history peer type=%d", msg.ErrSendStateConflict, in.PeerType)
	}
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	if in.PeerType == payload.PeerTypeChat {
		if err := c.checkChatAccess(in.UserId, in.PeerId, chatpb.ChatAccessReadHistory); err != nil {
			return nil, err
		}
	}

	maxPeerSeq := int64(0)
	maxUserMessageID := int64(0)
	if in.MaxId > 0 {
		resolved, err := c.svcCtx.Repo.ResolveMessageID(c.ctx, in.UserId, in.PeerType, in.PeerId, int64(in.MaxId))
		if err != nil {
			return nil, err
		}
		if resolved == nil {
			return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{}).ToMessagesAffectedMessages(), nil
		}
		maxPeerSeq = resolved.PeerSeq
		maxUserMessageID = resolved.UserMessageID
	}

	date := int32(time.Now().Unix())
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:        payload.MessageOperationSchemaVersion,
		OperationKind:        payload.OperationKindReadHistory,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		PeerSeq:              maxPeerSeq,
		FromUserID:           in.UserId,
		ToUserID:             in.UserId,
		Date:                 date,
		ReadMaxUserMessageID: maxUserMessageID,
		ReadInboxMaxPeerSeq:  maxPeerSeq,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: marshal read history operation user_id=%d peer_id=%d", msg.ErrMsgStorage, in.UserId, in.PeerId)
	}
	hashBytes := payload.HashBytes(body)
	authKeyID := in.AuthKeyId
	operationID := readHistoryOperationID(in.UserId, in.PeerId, in.MaxId, in.AuthKeyId)

	requester := OperationEnvelope{
		UserID:               in.UserId,
		OperationID:          operationID,
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindReadHistory,
		ActorUserID:          in.UserId,
		AuthKeyID:            &authKeyID,
		AuthKeyIDExclude:     &authKeyID,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		CanonicalPeerSeq:     int64Ptr(maxPeerSeq),
		CanonicalDate:        int64Ptr(int64(date)),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		PayloadCodec:         payload.PayloadCodecJSON,
		PayloadHash:          hashBytes,
		Payload:              body,
		DeliveryPolicy:       DeliveryPolicyRequesterSync,
	}

	effects, err := c.readHistoryOutboxEffects(in, maxPeerSeq, date)
	if err != nil {
		return nil, err
	}

	result, err := c.dispatchRequesterSync(requester, effects)
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(result.Pts, "pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      pts,
		PtsCount: result.PtsCount,
	}).ToMessagesAffectedMessages(), nil
}

func (c *MsgCore) readHistoryOutboxEffects(in *msg.TLMsgReadHistoryV2, maxPeerSeq int64, date int32) ([]OperationEnvelope, error) {
	if maxPeerSeq <= 0 {
		return nil, nil
	}
	switch in.PeerType {
	case payload.PeerTypeUser:
		if in.UserId == in.PeerId {
			return nil, nil
		}
		effect, err := readHistoryOutboxEffect(in.PeerId, in.UserId, in.UserId, in.PeerId, in.PeerType, maxPeerSeq, date)
		if err != nil {
			return nil, err
		}
		return []OperationEnvelope{effect}, nil
	case payload.PeerTypeChat:
		memberIDs, err := c.activeChatReceiverIDs(in.PeerId, in.UserId)
		if err != nil {
			return nil, err
		}
		effects := make([]OperationEnvelope, 0, len(memberIDs))
		for _, memberID := range memberIDs {
			effect, err := readHistoryOutboxEffect(memberID, in.PeerId, in.UserId, memberID, in.PeerType, maxPeerSeq, date)
			if err != nil {
				return nil, err
			}
			effects = append(effects, effect)
		}
		return effects, nil
	default:
		return nil, nil
	}
}

func readHistoryOutboxEffect(userID int64, peerID int64, readerUserID int64, targetUserID int64, peerType int32, maxPeerSeq int64, date int32) (OperationEnvelope, error) {
	peerBody, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:        payload.MessageOperationSchemaVersion,
		OperationKind:        payload.OperationKindReadHistory,
		PeerType:             peerType,
		PeerID:               peerID,
		PeerSeq:              maxPeerSeq,
		FromUserID:           readerUserID,
		ToUserID:             targetUserID,
		Date:                 date,
		Out:                  true,
		ReadOutboxMaxPeerSeq: maxPeerSeq,
	})
	if err != nil {
		return OperationEnvelope{}, fmt.Errorf("%w: marshal read outbox operation user_id=%d peer_id=%d", msg.ErrMsgStorage, userID, peerID)
	}
	return OperationEnvelope{
		UserID:               userID,
		OperationID:          readHistoryOutboxOperationID(userID, peerID, readerUserID, maxPeerSeq),
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindReadHistory,
		ActorUserID:          readerUserID,
		PeerType:             peerType,
		PeerID:               peerID,
		CanonicalPeerSeq:     int64Ptr(maxPeerSeq),
		CanonicalDate:        int64Ptr(int64(date)),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		PayloadCodec:         payload.PayloadCodecJSON,
		PayloadHash:          payload.HashBytes(peerBody),
		Payload:              peerBody,
		DeliveryPolicy:       DeliveryPolicyDurableAsync,
	}, nil
}

func readHistoryOperationID(userID int64, peerID int64, maxID int32, authKeyID int64) string {
	return fmt.Sprintf("v2:dialog:read_history:user:%d:peer:%d:max_user:%d:auth:%d", userID, peerID, maxID, authKeyID)
}

func readHistoryOutboxOperationID(userID int64, peerID int64, readerUserID int64, maxPeerSeq int64) string {
	return fmt.Sprintf("v2:dialog:read_history_outbox:user:%d:peer:%d:reader:%d:max_peer_seq:%d", userID, peerID, readerUserID, maxPeerSeq)
}
