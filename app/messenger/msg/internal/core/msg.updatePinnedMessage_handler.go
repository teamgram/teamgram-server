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

// MsgUpdatePinnedMessage
// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (c *MsgCore) MsgUpdatePinnedMessage(in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error) {
	if in == nil {
		return nil, fmt.Errorf("%w: missing update pinned request", msg.ErrSendStateConflict)
	}
	if in.UserId <= 0 || in.PeerId <= 0 || in.Id < 0 {
		return nil, fmt.Errorf("%w: invalid update pinned request", msg.ErrSendStateConflict)
	}
	if in.PeerType != payload.PeerTypeUser {
		return nil, fmt.Errorf("%w: update pinned first slice only supports user peer", msg.ErrSendStateConflict)
	}
	var pinnedPeerSeq int64
	var pinnedCanonicalID int64
	var pinnedUserMessageID int64
	if !in.Unpin && in.Id > 0 {
		var err error
		resolved, err := c.svcCtx.Repo.ResolveMessageID(c.ctx, in.UserId, in.PeerType, in.PeerId, int64(in.Id))
		if err != nil {
			return nil, err
		}
		if resolved == nil {
			return nil, msg.ErrSendStateConflict
		}
		pinnedPeerSeq = resolved.PeerSeq
		pinnedCanonicalID = resolved.CanonicalMessageID
		pinnedUserMessageID = resolved.UserMessageID
	}
	body, hashBytes, err := buildPinnedMessageOperation(in, pinnedPeerSeq, pinnedCanonicalID, pinnedUserMessageID)
	if err != nil {
		return nil, err
	}
	authKeyID := in.AuthKeyId
	result, err := c.dispatchRequesterSync(OperationEnvelope{
		UserID:               in.UserId,
		OperationID:          updatePinnedOperationID(in.UserId, in.PeerId, in.Id, in.Unpin, in.AuthKeyId),
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindUpdatePinnedMessage,
		ActorUserID:          in.UserId,
		AuthKeyID:            &authKeyID,
		AuthKeyIDExclude:     &authKeyID,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		CanonicalMessageID:   int64Ptr(pinnedCanonicalID),
		CanonicalPeerSeq:     int64Ptr(pinnedPeerSeq),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		PayloadCodec:         payload.PayloadCodecJSON,
		PayloadHash:          hashBytes,
		Payload:              body,
		DeliveryPolicy:       DeliveryPolicyRequesterSync,
	}, nil)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	pts, err := int64ToInt32(result.Pts, "pts")
	if err != nil {
		return nil, err
	}
	messages := []int32(nil)
	if !in.Unpin {
		messages = []int32{in.Id}
	}
	date, err := msgDateInt32FromUnixSeconds(time.Now().UTC().Unix(), "update pinned date")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdatePinnedMessages(&tg.TLUpdatePinnedMessages{
			Pinned:   !in.Unpin,
			Peer:     tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: in.PeerId}),
			Messages: messages,
			Pts:      pts,
			PtsCount: result.PtsCount,
		}),
		Date: date,
	}).ToUpdates(), nil
}

func buildPinnedMessageOperation(in *msg.TLMsgUpdatePinnedMessage, pinnedPeerSeq int64, pinnedCanonicalID int64, pinnedUserMessageID int64) ([]byte, []byte, error) {
	date, err := msgDateInt32FromUnixSeconds(time.Now().UTC().Unix(), "pinned operation date")
	if err != nil {
		return nil, nil, err
	}
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:            payload.MessageOperationSchemaVersion,
		OperationKind:            payload.OperationKindUpdatePinnedMessage,
		CanonicalMessageID:       pinnedCanonicalID,
		PeerType:                 in.PeerType,
		PeerID:                   in.PeerId,
		PeerSeq:                  pinnedPeerSeq,
		ToUserID:                 in.UserId,
		Date:                     date,
		PinnedPeerSeq:            pinnedPeerSeq,
		PinnedUserMessageID:      pinnedUserMessageID,
		PinnedCanonicalMessageID: pinnedCanonicalID,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: marshal update pinned operation user_id=%d peer_id=%d", msg.ErrMsgStorage, in.UserId, in.PeerId)
	}
	return body, payload.HashBytes(body), nil
}

func updatePinnedOperationID(userID int64, peerID int64, id int32, unpin bool, authKeyID int64) string {
	return fmt.Sprintf("v1:dialog:update_pinned:user:%d:peer:%d:id:%d:unpin:%t:auth:%d", userID, peerID, id, unpin, authKeyID)
}
