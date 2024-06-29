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
	"context"
	"math/rand"
	"time"

	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// MessagesForwardMessages
// messages.forwardMessages#cc30290b flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true drop_author:flags.11?true drop_media_captions:flags.12?true noforwards:flags.14?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer schedule_date:flags.10?int send_as:flags.13?InputPeer = Updates;
func (c *MessagesCore) MessagesForwardMessages(in *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	var (
		fromPeer = mtproto.FromInputPeer2(c.MD.UserId, in.FromPeer)
		toPeer   = mtproto.FromInputPeer2(c.MD.UserId, in.ToPeer)
		err      error
		rUpdates *mtproto.Updates
		saved    = false
	)

	//if c.MD.IsBot {
	//	err := mtproto.ErrBotMethodInvalid
	//	c.Logger.Errorf("messages.forwardMessages - error: %v", err)
	//	return nil, err
	//}

	/*
	   ## android's from_peer maybe is empty
	   if (msgObj.messageOwner.to_id instanceof TLRPC.TL_peerChannel) {
	       TLRPC.Chat chat = MessagesController.getInstance(currentAccount).getChat(msgObj.messageOwner.to_id.channel_id);
	       req.from_peer = new TLRPC.TL_inputPeerChannel();
	       req.from_peer.channel_id = msgObj.messageOwner.to_id.channel_id;
	       if (chat != null) {
	           req.from_peer.access_hash = chat.access_hash;
	       }
	   } else {
	       req.from_peer = new TLRPC.TL_inputPeerEmpty();
	   }
	*/

	switch toPeer.PeerType {
	case mtproto.PEER_SELF:
		toPeer.PeerType = mtproto.PEER_USER
		saved = true
	case mtproto.PEER_USER:
		if toPeer.PeerId == c.MD.UserId {
			saved = true
		}
	case mtproto.PEER_CHAT:
	case mtproto.PEER_CHANNEL:
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.forwardMessages#708e0195 - error: %v", err)
		return nil, err
	}

	if len(in.Id) == 0 ||
		len(in.RandomId) == 0 ||
		len(in.Id) != len(in.RandomId) {

		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("invalid id or random_id")
		return nil, err
	}

	fwdOutboxList, err := c.makeForwardMessages(fromPeer, toPeer, saved, in)
	if err != nil {
		c.Logger.Errorf("messages.forwardMessages#708e0195 - error: %v", err)
		return nil, err
	}

	rUpdates, err = c.svcCtx.Dao.MsgClient.MsgSendMessageV2(
		c.ctx,
		&msgpb.TLMsgSendMessageV2{
			UserId:    c.MD.UserId,
			AuthKeyId: c.MD.PermAuthKeyId,
			PeerType:  toPeer.PeerType,
			PeerId:    toPeer.PeerId,
			Message:   fwdOutboxList,
		})
	if err != nil {
		c.Logger.Errorf("messages.forwardMessages - error: %v", err)
		return nil, err
	}

	return rUpdates, err
}

func (c *MessagesCore) checkForwardPrivacy(ctx context.Context, selfUserId, checkId int64) bool {
	rules, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
		UserId:  selfUserId,
		KeyType: mtproto.FORWARDS,
	})

	if len(rules.Datas) == 0 {
		return true
	}
	return mtproto.CheckPrivacyIsAllow(
		selfUserId,
		rules.Datas,
		checkId,
		func(id, checkId int64) bool {
			contact, _ := c.svcCtx.Dao.UserClient.UserCheckContact(c.ctx, &userpb.TLUserCheckContact{
				UserId: id,
				Id:     checkId,
			})
			return mtproto.FromBool(contact)
		},
		func(checkId int64, idList []int64) bool {
			chatIdList, _ := mtproto.SplitChatAndChannelIdList(idList)
			return c.svcCtx.Dao.ChatClient.CheckParticipantIsExist(c.ctx, checkId, chatIdList)
		})
}

func (c *MessagesCore) makeForwardMessages(
	fromPeer, toPeer *mtproto.PeerUtil,
	saved bool,
	request *mtproto.TLMessagesForwardMessages) ([]*msgpb.OutboxMessage, error) {

	var (
		idList  = request.Id
		ridList = request.RandomId
		now     = int32(time.Now().Unix())
		// messageList []*mtproto.Message
		// err error
	)

	// TODO(@benqi): sorted map
	findRandomIdById := func(id int32) int64 {
		for i := 0; i < len(idList); i++ {
			if id == idList[i] {
				return ridList[i]
			}
		}
		return 0
	}

	var (
		messageList *message.Vector_MessageBox
	)

	switch fromPeer.PeerType {
	case mtproto.PEER_CHANNEL:
		// TODO: not impl
		c.Logger.Errorf("messages.forwardMessages blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		messageList, _ = c.svcCtx.Dao.MessageClient.MessageGetUserMessageList(c.ctx, &message.TLMessageGetUserMessageList{
			UserId: c.MD.UserId,
			IdList: idList,
		})
		if messageList.Length() > 0 {
			msgBox0 := messageList.Datas[0]
			if msgBox0.PeerType == mtproto.PEER_CHAT {
				chat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
					ChatId: msgBox0.PeerId,
				})
				if err != nil {
					c.Logger.Errorf("messages.forwardMessages - error: %v", err)
					return nil, err
				}

				if chat.Noforwards() {
					err = mtproto.ErrChatForwardsRestricted
					c.Logger.Errorf("messages.forwardMessages - error: %v", err)
					return nil, err
				}
			}
		}
	}

	fwdOutboxList := make([]*msgpb.OutboxMessage, 0, int(messageList.Length()))
	groupedIds := make(map[int64]int64)
	for _, box := range messageList.Datas {
		m := box.Message
		// TODO(@benqi): rid is 0

		if m.GetGroupedId() != nil {
			groupedId := m.GetGroupedId().GetValue()
			if _, ok := groupedIds[groupedId]; !ok {
				groupedIds[groupedId] = rand.Int63()
			}
		}
		if mtproto.IsMusicMessage(m) {
			m.FwdFrom = nil
		} else {
			if m.FwdFrom == nil {
				fwdFrom := mtproto.MakeTLMessageFwdHeader(&mtproto.MessageFwdHeader{
					Imported:       false,
					SavedOut:       false,
					FromId:         nil,
					FromName:       nil,
					Date:           m.GetDate(),
					ChannelPost:    nil,
					PostAuthor:     nil,
					SavedFromPeer:  nil,
					SavedFromMsgId: nil,
					SavedFromId:    nil,
					SavedFromName:  nil,
					SavedDate:      nil,
					PsaType:        nil,
				}).To_MessageFwdHeader()

				//fwdFrom := mtproto.MakeTLMessageFwdHeader(&mtproto.MessageFwdHeader{
				//	// FromId: m.GetFromId(),
				//	Date: m.GetDate(),
				//}).To_MessageFwdHeader()

				if m.Views != nil {
					// Broadcast
					// fwdFrom.ChannelId = &wrapperspb.Int32Value{Value: fromPeer.PeerId}
					fwdFrom.ChannelPost = &wrapperspb.Int32Value{Value: m.Id}
					fwdFrom.PostAuthor = m.PostAuthor
					fwdFrom.FromId = mtproto.MakePeerChannel(fromPeer.PeerId)
					// TODO(@benqi): saved_from_peer and saved_from_msg_id??
				} else {
					fromId := box.SenderUserId
					if c.checkForwardPrivacy(c.ctx, fromId, c.MD.UserId) {
						fwdFrom.FromId = mtproto.MakePeerUser(fromId)
					} else {
						uname, _ := c.svcCtx.Dao.UsernameClient.UsernameGetAccountUsername(c.ctx, &username.TLUsernameGetAccountUsername{
							UserId: fromId,
						})
						fwdFrom.FromName = &wrapperspb.StringValue{Value: uname.GetUsername()}
					}
					m.Post = false
					m.PostAuthor = nil
				}

				if saved {
					if m.Views != nil {
						// fwdFrom
						fwdFrom.SavedFromPeer = box.Message.GetPeerId()
					} else {
						fwdFrom.SavedFromPeer = mtproto.MakePeerUser(box.SenderUserId)
					}
					fwdFrom.SavedFromMsgId = &wrapperspb.Int32Value{Value: m.Id}
					m.SavedPeerId = fwdFrom.SavedFromPeer
				} else {
					m.SavedPeerId = nil
				}
				m.FwdFrom = fwdFrom
			} else {
				if saved {
					if m.Views != nil {
						// fwdFrom
						m.FwdFrom.SavedFromPeer = box.Message.GetFromId()
					} else {
						m.FwdFrom.SavedFromPeer = mtproto.MakePeerUser(box.SenderUserId)
					}
					m.FwdFrom.SavedFromMsgId = &wrapperspb.Int32Value{Value: m.Id}
					m.SavedPeerId = m.FwdFrom.SavedFromPeer
				} else {
					m.FwdFrom.SavedFromPeer = nil
					m.FwdFrom.SavedFromMsgId = nil
					m.SavedPeerId = nil
				}
			}
		}

		// TODO(@benqi): make message, ref sendMessage
		m.PeerId = toPeer.ToPeer()
		m.FromId = mtproto.MakePeerUser(c.MD.UserId)
		m.Date = now
		m.Silent = request.Silent
		m.Post = false
		if m.GetGroupedId() != nil {
			groupedId := groupedIds[m.GetGroupedId().GetValue()]
			m.GroupedId = mtproto.MakeFlagsInt64(groupedId)
		} else {
			m.GroupedId = nil
		}
		m.ReplyTo = nil
		m.Reactions = nil
		if m.ReplyMarkup != nil {
			m.ReplyMarkup = m.ReplyMarkup.ToForwardMessage()
		}

		fwdOutboxList = append(fwdOutboxList, &msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     findRandomIdById(m.GetId()),
			Message:      m,
			ScheduleDate: request.GetScheduleDate(),
		})
	}

	return fwdOutboxList, nil
}
