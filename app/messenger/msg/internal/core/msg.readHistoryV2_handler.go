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
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
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
	if in.PeerType != payload.PeerTypeUser {
		return nil, fmt.Errorf("%w: read history first slice only supports user peer", msg.ErrSendStateConflict)
	}
	if c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}

	date := int32(time.Now().Unix())
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:        payload.MessageOperationSchemaVersion,
		OperationKind:        payload.OperationKindReadHistory,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		PeerSeq:              int64(in.MaxId),
		FromUserID:           in.UserId,
		ToUserID:             in.UserId,
		Date:                 date,
		ReadInboxMaxPeerSeq:  int64(in.MaxId),
		ReadOutboxMaxPeerSeq: int64(in.MaxId),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: marshal read history operation user_id=%d peer_id=%d", msg.ErrMsgStorage, in.UserId, in.PeerId)
	}
	hashBytes := payload.HashBytes(body)
	route := payload.RouteUser(in.UserId)
	authKeyID := in.AuthKeyId
	operationID := readHistoryOperationID(in.UserId, in.PeerId, in.MaxId, in.AuthKeyId)
	result, err := c.svcCtx.UserUpdates.UserupdatesProcessUserOperation(c.ctx, &userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               in.UserId,
			BucketId:             int32(route.BucketID),
			PartitionId:          int32(route.ReceiverPartitionID),
			OperationId:          operationID,
			OpType:               payload.OpTypeSendMessage,
			OpSource:             0,
			ActorUserId:          in.UserId,
			AuthKeyId:            &authKeyID,
			AuthKeyIdExclude:     &authKeyID,
			PeerType:             in.PeerType,
			PeerId:               in.PeerId,
			CanonicalPeerSeq:     int64Ptr(int64(in.MaxId)),
			CanonicalDate:        int64Ptr(int64(date)),
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
	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      pts,
		PtsCount: result.PtsCount,
	}).ToMessagesAffectedMessages(), nil
}

func readHistoryOperationID(userID int64, peerID int64, maxID int32, authKeyID int64) string {
	return fmt.Sprintf("v1:dialog:read_history:user:%d:peer:%d:max:%d:auth:%d", userID, peerID, maxID, authKeyID)
}
