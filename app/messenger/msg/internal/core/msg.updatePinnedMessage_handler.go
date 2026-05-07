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

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
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
	if c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}

	var canonical *repository.CanonicalMessage
	if !in.Unpin {
		var err error
		canonical, err = c.svcCtx.Repo.GetCanonicalMessageByPeerSeq(c.ctx, in.UserId, in.PeerType, in.PeerId, int64(in.Id))
		if err != nil {
			return nil, err
		}
	}
	body, hashBytes, err := buildPinnedMessageOperation(in, canonical)
	if err != nil {
		return nil, err
	}
	route := payload.RouteUser(in.UserId)
	authKeyID := in.AuthKeyId
	result, err := c.svcCtx.UserUpdates.UserupdatesProcessUserOperation(c.ctx, &userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               in.UserId,
			BucketId:             int32(route.BucketID),
			PartitionId:          int32(route.ReceiverPartitionID),
			OperationId:          updatePinnedOperationID(in.UserId, in.PeerId, in.Id, in.Unpin, in.AuthKeyId),
			OpType:               payload.OpTypeSendMessage,
			OpSource:             0,
			ActorUserId:          in.UserId,
			AuthKeyId:            &authKeyID,
			AuthKeyIdExclude:     &authKeyID,
			PeerType:             in.PeerType,
			PeerId:               in.PeerId,
			CanonicalPeerSeq:     int64Ptr(int64(in.Id)),
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         payload.PayloadCodecJSON,
			PayloadHash:          hashBytes,
			Payload:              body,
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", msg.ErrSenderSyncFailed, err)
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
	date, err := tg.DateInt32FromUnixSeconds(time.Now().UTC().Unix())
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

func buildPinnedMessageOperation(in *msg.TLMsgUpdatePinnedMessage, canonical *repository.CanonicalMessage) ([]byte, []byte, error) {
	date, err := tg.DateInt32FromUnixSeconds(time.Now().UTC().Unix())
	if err != nil {
		return nil, nil, err
	}
	var canonicalID int64
	var peerSeq int64
	var fromUserID int64
	var messageText string
	if canonical != nil {
		canonicalID = canonical.CanonicalMessageID
		peerSeq = canonical.PeerSeq
		fromUserID = canonical.FromUserID
		messageText = canonical.MessageText
		if canonical.MessageDate != 0 {
			date, err = tg.DateInt32FromUnixSeconds(canonical.MessageDate)
			if err != nil {
				return nil, nil, err
			}
		}
	}
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:            payload.MessageOperationSchemaVersion,
		OperationKind:            payload.OperationKindUpdatePinnedMessage,
		CanonicalMessageID:       canonicalID,
		PeerType:                 in.PeerType,
		PeerID:                   in.PeerId,
		PeerSeq:                  peerSeq,
		FromUserID:               fromUserID,
		ToUserID:                 in.UserId,
		Date:                     date,
		MessageText:              messageText,
		PinnedPeerSeq:            peerSeq,
		PinnedCanonicalMessageID: canonicalID,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: marshal update pinned operation user_id=%d peer_id=%d", msg.ErrMsgStorage, in.UserId, in.PeerId)
	}
	return body, payload.HashBytes(body), nil
}

func updatePinnedOperationID(userID int64, peerID int64, id int32, unpin bool, authKeyID int64) string {
	return fmt.Sprintf("v1:dialog:update_pinned:user:%d:peer:%d:id:%d:unpin:%t:auth:%d", userID, peerID, id, unpin, authKeyID)
}
