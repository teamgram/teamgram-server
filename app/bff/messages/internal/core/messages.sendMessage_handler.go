// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/threading"
	"time"

	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesSendMessage
// messages.sendMessage#d9d75a4 flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer = Updates;
func (c *MessagesCore) MessagesSendMessage(in *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error) {
	// peer
	var (
		hasBot = c.MD.IsBot
		peer   = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	)

	if !peer.IsChatOrUser() {
		c.Logger.Errorf("invalid peer: %v", in.Peer)
		err := mtproto.ErrPeerIdInvalid
		return nil, err
	}

	if peer.IsUser() {
		if peer.IsSelfUser(c.MD.UserId) {
			peer.PeerType = mtproto.PEER_USER
		} else {
			if !c.MD.IsBot {
				mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
					Id: []int64{c.MD.UserId, peer.PeerId},
				})
				peerUser, _ := mUsers.GetImmutableUser(peer.PeerId)
				if peerUser == nil {
					err := mtproto.ErrPeerIdInvalid
					return nil, err
				}
				hasBot = peerUser.IsBot()
			}
		}
	}

	if in.Message == "" {
		err := mtproto.ErrMessageEmpty
		c.Logger.Errorf("message empty: %v", err)
		return nil, err
	}
	// TODO(@benqi): calc utf16len(message)
	//else if len(request.Message) > 4000 {
	//	err = mtproto.ErrMessageTooLong
	//	c.Logger.Errorf("messages.sendMessage: %v", err)
	//	return
	//}

	outMessage := mtproto.MakeTLMessage(&mtproto.Message{
		Out:               true,
		Mentioned:         false,
		MediaUnread:       false,
		Silent:            in.GetSilent(),
		Post:              false,
		FromScheduled:     false,
		Legacy:            false,
		EditHide:          false,
		Pinned:            false,
		Noforwards:        in.GetNoforwards(),
		Id:                0,
		FromId:            mtproto.MakePeerUser(c.MD.UserId),
		PeerId:            peer.ToPeer(),
		FwdFrom:           nil,
		ViaBotId:          nil,
		ReplyTo:           nil,
		Date:              int32(time.Now().Unix()),
		Message:           in.Message,
		Media:             nil,
		ReplyMarkup:       in.ReplyMarkup,
		Entities:          in.Entities,
		Views:             nil,
		Forwards:          nil,
		Replies:           nil,
		EditDate:          nil,
		PostAuthor:        nil,
		GroupedId:         nil,
		Reactions:         nil,
		RestrictionReason: nil,
		TtlPeriod:         nil,
	}).To_Message()

	// Fix ReplyToMsgId
	if in.GetReplyToMsgId() != nil {
		outMessage.ReplyTo = mtproto.MakeTLMessageReplyHeader(&mtproto.MessageReplyHeader{
			ReplyToMsgId:  in.GetReplyToMsgId().GetValue(),
			ReplyToPeerId: nil,
			ReplyToTopId:  nil,
		}).To_MessageReplyHeader()
	}

	outMessage, _ = c.fixMessageEntities(c.MD.UserId, peer, in.NoWebpage, outMessage, hasBot)
	rUpdate, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    in.NoWebpage,
			Background:   in.Background,
			RandomId:     in.RandomId,
			Message:      outMessage,
			ScheduleDate: in.ScheduleDate,
		}).To_OutboxMessage(),
	})

	if err != nil {
		c.Logger.Errorf("messages.sendMessage#fa88427a - error: %v", err)
		return nil, err
	}

	if in.ClearDraft {
		ctx := contextx.ValueOnlyFrom(c.ctx)
		threading.GoSafe(func() {
			c.doClearDraft(ctx, c.MD.UserId, c.MD.AuthId, peer)
		})
	}

	return rUpdate, nil
}
