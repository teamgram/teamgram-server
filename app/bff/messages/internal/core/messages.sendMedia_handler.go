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
	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"time"
)

// MessagesSendMedia
// messages.sendMedia#e25ff8e0 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer = Updates;
func (c *MessagesCore) MessagesSendMedia(in *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	// peer
	var (
		hasBot     = c.MD.IsBot
		peer       *mtproto.PeerUtil
		linkChatId int64
		err        error
	)

	peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	switch peer.PeerType {
	case mtproto.PEER_SELF:
		peer.PeerType = mtproto.PEER_USER
	case mtproto.PEER_USER:
		if !c.MD.IsBot {
			// hasBot = s.UserFacade.IsBot(ctx, peer.PeerId)
		}
	case mtproto.PEER_CHAT:
	case mtproto.PEER_CHANNEL:
		//channel, _ := s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId)
		//if channel != nil && channel.Channel.LinkedChatId > 0 {
		//	linkChatId = channel.Channel.LinkedChatId
		//}
	default:
		c.Logger.Errorf("invalid peer: %v", in.Peer)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	if len(in.Message) > 4000 {
		err = mtproto.ErrMediaCaptionTooLong
		c.Logger.Errorf("messages.sendMedia: %v", err)
		return nil, err
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// 发件箱

	/*
		Name	Type	Description
		flags	#	Flags, see TL conditional fields
		silent	flags.5?true	Send message silently (no notification should be triggered)
		background	flags.6?true	Send message in background
		clear_draft	flags.7?true	Clear the draft
		peer	InputPeer	Destination
		reply_to_msg_id	flags.0?int	Message ID to which this message should reply to
		media	InputMedia	Attached media
		message	string	Caption
		random_id	long	Random ID to avoid resending the same message
		reply_markup	flags.2?ReplyMarkup	Reply markup for bot keyboards
		entities	flags.3?Vector<MessageEntity>	Message entities for styled text
		schedule_date	flags.10?int	Scheduled message date for scheduled messages
	*/
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
		Id:                0,
		FromId:            mtproto.MakePeerUser(c.MD.UserId),
		PeerId:            peer.ToPeer(),
		FwdFrom:           nil,
		ViaBotId:          nil,
		ReplyTo:           nil,
		Date:              int32(time.Now().Unix()),
		Media:             nil,
		Message:           in.Message,
		ReplyMarkup:       in.ReplyMarkup,
		Entities:          in.Entities,
		Views:             nil,
		Forwards:          nil,
		Replies:           nil,
		EditDate:          nil,
		PostAuthor:        nil,
		GroupedId:         nil,
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

	if linkChatId > 0 {
		outMessage.Replies = mtproto.MakeTLMessageReplies(&mtproto.MessageReplies{
			Comments:       true,
			Replies:        0,
			RepliesPts:     0,
			RecentRepliers: nil,
			ChannelId:      mtproto.MakeFlagsInt64(linkChatId),
			MaxId:          nil,
			ReadMaxId:      nil,
		}).To_MessageReplies()
	}

	outMessage.Media, err = c.makeMediaByInputMedia(in.Media)
	if err != nil {
		c.Logger.Errorf("messages.sendMedia - error: %v", err)
		return nil, err
	}

	//// poll
	//poll, _ := mtproto.GetPollByMessage(outMessage.Media)
	//if poll != nil {
	//	var correctAnswers []int32
	//	for _, v := range in.Media.CorrectAnswers {
	//		if iV, err := strconv.ParseInt(hack.String(v), 10, 64); err == nil {
	//			correctAnswers = append(correctAnswers, int32(iV))
	//		}
	//	}
	//
	//	// TODO(@benqi): 112
	//	// solution:flags.1?string solution_entities:flags.1?Vector<MessageEntity>
	//
	//	// outMessage.Media,
	//	pollMedia, err := c.svcCtx.Dao.PollClient.PollCreateMediaPoll(c.ctx, &pollpb.TLPollCreateMediaPoll{
	//		UserId:         c.MD.UserId,
	//		CorrectAnswers: correctAnswers,
	//		Poll:           poll,
	//	})
	//	if err != nil {
	//		c.Logger.Errorf("createMediaPoll error - %v", err)
	//		return nil, err
	//	}
	//	outMessage.Media = pollMedia
	//
	//	// TODO(@benqi): poll2.0 disable, debug it.
	//	// outMessage.Media.Results.Results = nil
	//}

	outMessage, _ = c.fixMessageEntities(c.MD.UserId, peer, true, outMessage, hasBot)
	rUpdate, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   in.Background,
			RandomId:     in.RandomId,
			Message:      outMessage,
			ScheduleDate: in.ScheduleDate,
		}).To_OutboxMessage(),
	})

	if err != nil {
		c.Logger.Errorf("messages.sendMedia#c8f16791 - error: %v", err)
		return nil, err
	} else {
		//if peer.PeerType == model.PEER_CHANNEL {
		//	updateChannelInbox := mtproto.MakeTLUpdateReadChannelInbox(&mtproto.Update{
		//		FolderId:         nil,
		//		ChannelId:        peer.PeerId,
		//		MaxId:            0,
		//		StillUnreadCount: 0,
		//		Pts_INT32:        0,
		//	}).To_Update()
		//
		//	model.VisitUpdates(md.UserId, reply, map[string]model.UpdateVisitedFunc{
		//		// DC_UPDATE
		//		mtproto.Predicate_updateNewChannelMessage: func(
		//			userId int32,
		//			update *mtproto.Update,
		//			users []*mtproto.User,
		//			chats []*mtproto.Chat,
		//			date int32) {
		//			updateChannelInbox.MaxId = update.Message_MESSAGE.Id
		//			updateChannelInbox.Pts_INT32 = update.Pts_INT32
		//		},
		//	})
		//
		//	reply.Updates = append(reply.Updates, updateChannelInbox)
		//}

		if in.ClearDraft {
			c.doClearDraft(c.MD.UserId, c.MD.AuthId, peer)
		}

		c.svcCtx.Dao.UserClient.UserUpdateLastSeen(c.ctx, &userpb.TLUserUpdateLastSeen{
			Id:         c.MD.UserId,
			LastSeenAt: time.Now().Unix(),
		})

		//sync_client.PushUpdates(ctx, md.UserId, model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadChannelInbox(&mtproto.Update{
		//	FolderId:         nil,
		//	ChannelId:        peer.PeerId,
		//	MaxId:            readInboxId,
		//	StillUnreadCount: 0,
		//	Pts_INT32:        0,
		//}).To_Update()))
	}

	return rUpdate, nil
}
