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
	"time"

	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/threading"
)

// MessagesSendMessage
// messages.sendMessage#d9d75a4 flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer = Updates;
func (c *MessagesCore) MessagesSendMessage(in *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	)

	if !peer.IsChatOrUser() {
		c.Logger.Errorf("invalid peer: %v", in.Peer)
		err := mtproto.ErrEnterpriseIsBlocked
		return nil, err
	}

	if peer.IsUser() && peer.IsSelfUser(c.MD.UserId) {
		peer.PeerType = mtproto.PEER_USER
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
		Silent:            in.Silent,
		Post:              false,
		FromScheduled:     false,
		Legacy:            false,
		EditHide:          false,
		Pinned:            false,
		Noforwards:        in.Noforwards,
		InvertMedia:       in.InvertMedia,
		Id:                0,
		FromId:            mtproto.MakePeerUser(c.MD.UserId),
		PeerId:            peer.ToPeer(),
		SavedPeerId:       nil,
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

	// Fix SavedPeerId
	if peer.IsSelfUser(c.MD.UserId) {
		outMessage.SavedPeerId = peer.ToPeer()
	}

	// Fix ReplyToMsgId
	if in.GetReplyToMsgId() != nil {
		outMessage.ReplyTo = mtproto.MakeTLMessageReplyHeader(&mtproto.MessageReplyHeader{
			ReplyToScheduled:       false,
			ForumTopic:             false,
			Quote:                  false,
			ReplyToMsgId:           in.GetReplyToMsgId().GetValue(),
			ReplyToMsgId_INT32:     in.GetReplyToMsgId().GetValue(),
			ReplyToMsgId_FLAGINT32: in.GetReplyToMsgId(),
			ReplyToPeerId:          nil,
			ReplyFrom:              nil,
			ReplyMedia:             nil,
			ReplyToTopId:           nil,
			QuoteText:              nil,
			QuoteEntities:          nil,
			QuoteOffset:            nil,
		}).To_MessageReplyHeader()
	} else if in.GetReplyTo() != nil {
		replyTo := in.GetReplyTo()
		switch in.ReplyTo.PredicateName {
		case mtproto.Predicate_inputReplyToMessage:
			outMessage.ReplyTo = mtproto.MakeTLMessageReplyHeader(&mtproto.MessageReplyHeader{
				ReplyToScheduled:       false,
				ForumTopic:             false,
				Quote:                  false,
				ReplyToMsgId:           replyTo.GetReplyToMsgId(),
				ReplyToMsgId_INT32:     replyTo.GetReplyToMsgId(),
				ReplyToMsgId_FLAGINT32: mtproto.MakeFlagsInt32(replyTo.GetReplyToMsgId()),
				ReplyToPeerId:          nil,
				ReplyFrom:              nil,
				ReplyMedia:             nil,
				ReplyToTopId:           nil,
				QuoteText:              nil,
				QuoteEntities:          nil,
				QuoteOffset:            nil,
			}).To_MessageReplyHeader()
			if replyTo.GetQuoteText() != nil {
				outMessage.ReplyTo.Quote = true
				outMessage.ReplyTo.QuoteText = replyTo.GetQuoteText()
				outMessage.ReplyTo.QuoteEntities = replyTo.GetQuoteEntities()
				outMessage.ReplyTo.QuoteOffset = replyTo.GetQuoteOffset()
			}

			// disable replyToPeerId
			// TODO enable replyToPeerId
			if replyTo.ReplyToPeerId != nil {
				outMessage.ReplyTo = nil
			}

		case mtproto.Predicate_inputReplyToStory:
			// TODO:
			var (
				rPeer  *mtproto.PeerUtil
				userId int64
			)

			if replyTo.GetUserId() != nil {
				rPeer = mtproto.FromInputUser(c.MD.UserId, replyTo.GetUserId())
				userId = rPeer.PeerId
			} else if replyTo.GetPeer() != nil {
				rPeer = mtproto.FromInputPeer2(c.MD.UserId, replyTo.GetPeer())
				if rPeer.IsUser() {
					userId = peer.PeerId
				}
			}

			if rPeer != nil {
				outMessage.ReplyTo = mtproto.MakeTLMessageReplyStoryHeader(&mtproto.MessageReplyHeader{
					UserId:  userId,
					Peer:    rPeer.ToPeer(),
					StoryId: replyTo.GetStoryId(),
				}).To_MessageReplyHeader()
			}
		}
	}

	//outMessage, _ = c.fixMessageEntities(c.MD.UserId, peer, in.NoWebpage, outMessage, func() bool {
	//	hasBot := c.MD.IsBot
	//	if !hasBot {
	//		isBot, _ := c.svcCtx.Dao.UserClient.UserIsBot(c.ctx, &userpb.TLUserIsBot{
	//			Id: peer.PeerId,
	//		})
	//		hasBot = mtproto.FromBool(isBot)
	//	}
	//
	//	return hasBot
	//})
	rUpdate, err := c.svcCtx.Dao.MsgClient.MsgSendMessageV2(c.ctx, &msgpb.TLMsgSendMessageV2{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.PermAuthKeyId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		Message: []*msgpb.OutboxMessage{
			msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
				NoWebpage:    in.NoWebpage,
				Background:   in.Background,
				RandomId:     in.RandomId,
				Message:      outMessage,
				ScheduleDate: in.ScheduleDate,
			}).To_OutboxMessage(),
		},
	})

	if err != nil {
		c.Logger.Errorf("messages.sendMessage#fa88427a - error: %v", err)
		return nil, err
	}

	if in.ClearDraft {
		ctx := contextx.ValueOnlyFrom(c.ctx)
		threading.GoSafe(func() {
			c.doClearDraft(ctx, c.MD.UserId, c.MD.PermAuthKeyId, peer)
		})
	}

	return rUpdate, nil
}
