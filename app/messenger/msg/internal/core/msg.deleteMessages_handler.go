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

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const maxInt32 = int32(^uint32(0) >> 1)

// MsgDeleteMessages
// msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;
func (c *MsgCore) MsgDeleteMessages(in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	if in == nil || in.UserId <= 0 || len(in.Id) == 0 {
		return nil, fmt.Errorf("%w: invalid delete messages request", msg.ErrSendStateConflict)
	}
	if in.PeerType != 0 && in.PeerType != payload.PeerTypeUser {
		return nil, fmt.Errorf("%w: delete messages first slice only supports user peer", msg.ErrSendStateConflict)
	}
	if in.PeerType == 0 && in.PeerId != 0 {
		return nil, fmt.Errorf("%w: invalid delete messages peer", msg.ErrSendStateConflict)
	}
	if in.PeerType != 0 && in.PeerId <= 0 {
		return nil, fmt.Errorf("%w: invalid delete messages peer", msg.ErrSendStateConflict)
	}
	if c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	userMessageIDs := make([]int64, 0, len(in.Id))
	for _, id := range in.Id {
		if id <= 0 {
			return nil, fmt.Errorf("%w: invalid message id", msg.ErrSendStateConflict)
		}
		userMessageIDs = append(userMessageIDs, int64(id))
	}
	resolved, err := c.svcCtx.Repo.ResolveMessageIDsForDelete(c.ctx, in.UserId, userMessageIDs)
	if err != nil {
		return nil, err
	}
	groups := groupDeleteMessageIDs(resolved, in.PeerType, in.PeerId)
	if len(groups) == 0 {
		return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{}).ToMessagesAffectedMessages(), nil
	}
	var finalPTS int64
	var ptsCount int32
	for _, group := range groups {
		body, hashBytes, err := buildDeleteMessagesPayload(in.UserId, in.UserId, group.peerType, group.peerID, group.messageDate, group.peerSeqs, group.userMessageIDs, in.Revoke)
		if err != nil {
			return nil, err
		}
		authKeyID := in.AuthKeyId
		var effects []OperationEnvelope
		if in.Revoke && group.peerType == payload.PeerTypeUser && group.peerID != in.UserId {
			peerBody, peerHash, err := buildDeleteMessagesPayload(in.UserId, group.peerID, group.peerType, in.UserId, group.messageDate, group.peerSeqs, nil, in.Revoke)
			if err != nil {
				return nil, err
			}
			effects = append(effects, OperationEnvelope{
				UserID:               group.peerID,
				OperationID:          deleteMessagesPeerSeqOperationID(group.peerID, in.UserId, group.peerSeqs, in.Revoke),
				OpType:               payload.OpTypeSendMessage,
				OperationKind:        payload.OperationKindDeleteMessages,
				ActorUserID:          in.UserId,
				PeerType:             group.peerType,
				PeerID:               in.UserId,
				PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
				PayloadCodec:         payload.PayloadCodecJSON,
				PayloadHash:          peerHash,
				Payload:              peerBody,
				DeliveryPolicy:       DeliveryPolicyDurableAsync,
			})
		}
		result, err := c.dispatchRequesterSync(OperationEnvelope{
			UserID:               in.UserId,
			OperationID:          deleteMessagesOperationID(in.UserId, group.peerID, int64SliceToInt32(group.userMessageIDs), in.Revoke, in.AuthKeyId),
			OpType:               payload.OpTypeSendMessage,
			OperationKind:        payload.OperationKindDeleteMessages,
			ActorUserID:          in.UserId,
			AuthKeyID:            &authKeyID,
			AuthKeyIDExclude:     &authKeyID,
			PeerType:             group.peerType,
			PeerID:               group.peerID,
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         payload.PayloadCodecJSON,
			PayloadHash:          hashBytes,
			Payload:              body,
			DeliveryPolicy:       DeliveryPolicyRequesterSync,
		}, effects)
		if err != nil {
			return nil, err
		}
		finalPTS = result.Pts
		ptsCount += result.PtsCount
	}
	pts, err := int64ToInt32(finalPTS, "pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{Pts: pts, PtsCount: ptsCount}).ToMessagesAffectedMessages(), nil
}

func deleteMessagesOperationID(userID int64, peerID int64, ids []int32, revoke bool, authKeyID int64) string {
	return fmt.Sprintf("v2:dialog:delete_messages:user:%d:peer:%d:ids:%v:revoke:%t:auth:%d", userID, peerID, ids, revoke, authKeyID)
}

func deleteMessagesPeerSeqOperationID(userID int64, peerID int64, peerSeqs []int64, revoke bool) string {
	return fmt.Sprintf("v2:dialog:delete_messages:user:%d:peer:%d:peer_seqs:%v:revoke:%t", userID, peerID, peerSeqs, revoke)
}

type deleteMessageGroup struct {
	peerType       int32
	peerID         int64
	messageDate    int32
	peerSeqs       []int64
	userMessageIDs []int64
}

func groupDeleteMessageIDs(items []repository.ResolvedMessageID, peerType int32, peerID int64) []deleteMessageGroup {
	groups := make([]deleteMessageGroup, 0, len(items))
	index := map[struct {
		peerType int32
		peerID   int64
	}]int{}
	for _, item := range items {
		if item.PeerSeq <= 0 || item.UserMessageID <= 0 {
			continue
		}
		if peerType != 0 && (item.PeerType != peerType || item.PeerID != peerID) {
			continue
		}
		key := struct {
			peerType int32
			peerID   int64
		}{peerType: item.PeerType, peerID: item.PeerID}
		groupIndex, ok := index[key]
		if !ok {
			groupIndex = len(groups)
			index[key] = groupIndex
			groups = append(groups, deleteMessageGroup{peerType: item.PeerType, peerID: item.PeerID})
		}
		if date := stableDeleteMessageDate(item.MessageDate); date > groups[groupIndex].messageDate {
			groups[groupIndex].messageDate = date
		}
		groups[groupIndex].peerSeqs = append(groups[groupIndex].peerSeqs, item.PeerSeq)
		groups[groupIndex].userMessageIDs = append(groups[groupIndex].userMessageIDs, item.UserMessageID)
	}
	return groups
}

func buildDeleteMessagesPayload(fromUserID int64, toUserID int64, peerType int32, peerID int64, date int32, peerSeqs []int64, userMessageIDs []int64, revoke bool) ([]byte, []byte, error) {
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:        payload.MessageOperationSchemaVersion,
		OperationKind:        payload.OperationKindDeleteMessages,
		PeerType:             peerType,
		PeerID:               peerID,
		FromUserID:           fromUserID,
		ToUserID:             toUserID,
		Date:                 date,
		DeletePeerSeqs:       peerSeqs,
		DeleteUserMessageIDs: userMessageIDs,
		Revoke:               revoke,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: marshal delete messages operation user_id=%d peer_id=%d", msg.ErrMsgStorage, fromUserID, peerID)
	}
	return body, payload.HashBytes(body), nil
}

func stableDeleteMessageDate(messageDate int64) int32 {
	switch {
	case messageDate <= 0:
		return 1
	case messageDate > int64(maxInt32):
		return maxInt32
	default:
		return int32(messageDate)
	}
}

func int64SliceToInt32(values []int64) []int32 {
	out := make([]int32, 0, len(values))
	for _, value := range values {
		out = append(out, int32(value))
	}
	return out
}
