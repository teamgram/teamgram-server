// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesMarkDialogUnread
// messages.markDialogUnread#8c5006f8 flags:# unread:flags.0?true parent_peer:flags.1?InputPeer peer:InputDialogPeer = Bool;
func (c *DialogsCore) MessagesMarkDialogUnread(in *tg.TLMessagesMarkDialogUnread) (*tg.Bool, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if c.MD.PermAuthKeyId == 0 {
		return nil, tg.ErrAuthKeyPermEmpty
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	peer, err := c.resolveInputDialogPeer(in.Peer)
	if err != nil {
		return nil, err
	}
	peerType, err := dialogFacadePeerType(peer.PeerType)
	if err != nil {
		return nil, err
	}
	if c.svcCtx.Repo.UserupdatesClient == nil {
		return nil, tg.ErrInternalServerError
	}
	unreadMark := in.Unread
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion: payload.MessageOperationSchemaVersion,
		OperationKind: payload.OperationKindMarkDialogUnread,
		PeerType:      peerType,
		PeerID:        peer.PeerId,
		FromUserID:    c.MD.UserId,
		ToUserID:      c.MD.UserId,
		Date:          int32(time.Now().Unix()),
		UnreadMark:    &unreadMark,
	})
	if err != nil {
		c.Logger.Errorf("messages.markDialogUnread - marshal operation failed: user_id: %d, peer_type: %d, peer_id: %d, err: %v", c.MD.UserId, peerType, peer.PeerId, err)
		return nil, tg.ErrInternalServerError
	}
	route := payload.RouteUser(c.MD.UserId)
	hashBytes := payload.HashBytes(body)
	authKeyID := c.MD.PermAuthKeyId
	operationID := fmt.Sprintf("v1:dialog:mark_unread:user:%d:peer:%d:%d:unread:%t:auth:%d", c.MD.UserId, peerType, peer.PeerId, unreadMark, authKeyID)
	if _, err := c.svcCtx.Repo.UserupdatesClient.UserupdatesProcessUserOperation(c.ctx, &userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               c.MD.UserId,
			BucketId:             int32(route.BucketID),
			PartitionId:          int32(route.ReceiverPartitionID),
			OperationId:          operationID,
			OpType:               payload.OpTypeSendMessage,
			OpSource:             0,
			ActorUserId:          c.MD.UserId,
			AuthKeyId:            &authKeyID,
			AuthKeyIdExclude:     &authKeyID,
			PeerType:             peerType,
			PeerId:               peer.PeerId,
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         payload.PayloadCodecJSON,
			PayloadHash:          hashBytes,
			Payload:              body,
		}),
	}); err != nil {
		c.Logger.Errorf("messages.markDialogUnread - userupdates.processUserOperation failed: user_id: %d, peer_type: %d, peer_id: %d, err: %v",
			c.MD.UserId, peerType, peer.PeerId, err)
		return nil, tg.ErrInternalServerError
	}
	return tg.BoolTrue, nil
}
