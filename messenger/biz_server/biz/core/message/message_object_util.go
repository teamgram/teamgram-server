// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package message

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"regexp"
	"strings"
)

// updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
// message#44f9b43d flags:#
// 	out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string grouped_id:flags.17?long = Message;
// messageService#9e19a1f6 flags:#
// 	out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer reply_to_msg_id:flags.3?int date:int action:MessageAction = Message;
func MessageToUpdateShortMessage(message2 *mtproto.Message) (shortMessage *mtproto.TLUpdateShortMessage) {
	// TODO(@benqi): check message2.ToId
	var (
		userId int32
	)

	switch message2.GetConstructor() {
	case mtproto.TLConstructor_CRC32_message:
		message := message2.To_Message()
		if message.GetOut() {
			userId = message.GetToId().GetData2().GetUserId()
		} else {
			userId = message.GetFromId()
		}
		shortMessage = &mtproto.TLUpdateShortMessage{Data2: &mtproto.Updates_Data{
			Out:          message.GetOut(),
			Mentioned:    message.GetMentioned(),
			MediaUnread:  message.GetMediaUnread(),
			Silent:       message.GetSilent(),
			Id:           message.GetId(),
			UserId:       userId,
			Message:      message.GetMessage(),
			Date:         message.GetDate(),
			FwdFrom:      message.GetFwdFrom(),
			ViaBotId:     message.GetViaBotId(),
			ReplyToMsgId: message.GetReplyToMsgId(),
			Entities:     message.GetEntities(),
		}}
	case mtproto.TLConstructor_CRC32_messageService:
	default:
		// TODO(@benqi): error
	}
	return
}

// updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
func MessageToUpdateShortChatMessage(message2 *mtproto.Message) (shortMessage *mtproto.TLUpdateShortChatMessage) {
	// TODO(@benqi): check message2.ToId
	switch message2.GetConstructor() {
	case mtproto.TLConstructor_CRC32_message:
		message := message2.To_Message()
		//if message.GetOut() {
		//	userId = message.GetToId().GetData2().GetUserId()
		//} else {
		//	userId = message.GetFromId()
		//}
		shortMessage = &mtproto.TLUpdateShortChatMessage{Data2: &mtproto.Updates_Data{
			Out:         message.GetOut(),
			Mentioned:   message.GetMentioned(),
			MediaUnread: message.GetMediaUnread(),
			Silent:      message.GetSilent(),
			Id:          message.GetId(),
			FromId:      message.GetFromId(),
			ChatId:      message.GetToId().GetData2().GetChatId(),
			// UserId:       userId,
			Message:      message.GetMessage(),
			Date:         message.GetDate(),
			FwdFrom:      message.GetFwdFrom(),
			ViaBotId:     message.GetViaBotId(),
			ReplyToMsgId: message.GetReplyToMsgId(),
			Entities:     message.GetEntities(),
		}}
	case mtproto.TLConstructor_CRC32_messageService:
		//message := message2.To_MessageService()
		//shortMessage = &mtproto.TLUpdateShortChatMessage{Data2: &mtproto.Updates_Data{
		//	Out:          message.GetOut(),
		//	Mentioned:    message.GetMentioned(),
		//	MediaUnread:  message.GetMediaUnread(),
		//	Silent:       message.GetSilent(),
		//	Id:           message.GetId(),
		//	FromId:       message.GetFromId(),
		//	ChatId:       message.GetToId().GetData2().GetChatId(),
		//	// UserId:       userId,
		//	// Message:      message.GetMessage(),
		//	Date:         message.GetDate(),
		//	// FwdFrom:      message.GetFwdFrom(),
		//	// ViaBotId:     message.GetViaBotId(),
		//	ReplyToMsgId: message.GetReplyToMsgId(),
		//	// Entities:     message.GetEntities(),
		//
		//}}
	default:
		// TODO(@benqi): error
	}
	return
}

//// updateShortSentMessage#11f1331c flags:# out:flags.1?true id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;
func MessageToUpdateShortSentMessage(message2 *mtproto.Message) (sentMessage *mtproto.TLUpdateShortSentMessage) {
	switch message2.GetConstructor() {
	case mtproto.TLConstructor_CRC32_message:
		message := message2.To_Message()
		sentMessage = &mtproto.TLUpdateShortSentMessage{Data2: &mtproto.Updates_Data{
			Out: message.GetOut(),
			Id:  message.GetId(),
			// Pts:,
			// PtsCount,
			Date:     message.GetDate(),
			Media:    message.GetMedia(),
			Entities: message.GetEntities(),
		}}
	case mtproto.TLConstructor_CRC32_messageService:
	default:
		// TODO(@benqi): error
	}
	return
}

func PickAllIDListByDialogs(dialogs []*mtproto.Dialog) (messageIdList, userIdList, chatIdList, channelIdList []int32) {
	if len(dialogs) == 0 {
		messageIdList = []int32{}
		userIdList = []int32{}
		chatIdList = []int32{}
		channelIdList = []int32{}
	} else {
		userIdList = make([]int32, 0, len(dialogs))
		chatIdList = make([]int32, 0, len(dialogs))
		channelIdList = make([]int32, 0, len(dialogs))

		for _, d := range dialogs {
			dialog := d.To_Dialog()
			messageIdList = append(messageIdList, dialog.GetTopMessage())

			p := dialog.GetPeer()
			// TODO(@benqi): 先假设只有PEER_USER
			switch p.GetConstructor() {
			case mtproto.TLConstructor_CRC32_peerUser:
				userIdList = append(userIdList, p.GetData2().GetUserId())
			case mtproto.TLConstructor_CRC32_peerChat:
				chatIdList = append(chatIdList, p.GetData2().GetChatId())
			case mtproto.TLConstructor_CRC32_peerChannel:
				channelIdList = append(channelIdList, p.GetData2().GetChannelId())
			}
		}
	}
	return
}

func PickAllIDListByMessages(messageList []*mtproto.Message) (userIdList, chatIdList, channelIdList []int32) {
	if len(messageList) == 0 {
		userIdList = []int32{}
		chatIdList = []int32{}
		channelIdList = []int32{}
	} else {
		userIdList = make([]int32, 0, len(messageList))
		chatIdList = make([]int32, 0, len(messageList))
		channelIdList = make([]int32, 0, len(messageList))

		AppendID := func (idList []int32, id int32) []int32 {
			for _, i := range idList {
				if i == id {
					return idList
				}
			}
			idList = append(idList, id)
			return idList
		}

		for _, m := range messageList {
			switch m.GetConstructor() {
			case mtproto.TLConstructor_CRC32_message:
				m2 := m.To_Message()
				userIdList = AppendID(userIdList, m2.GetFromId())

				p := m2.GetToId()
				switch p.GetConstructor() {
				case mtproto.TLConstructor_CRC32_peerUser:
					userIdList = AppendID(userIdList, p.GetData2().GetUserId())
				case mtproto.TLConstructor_CRC32_peerChat:
					chatIdList = AppendID(chatIdList, p.GetData2().GetChatId())
					if p.GetData2().GetChatId() == 0 {
						glog.Info("chat_id = 0: ", m)
						continue
					}
				case mtproto.TLConstructor_CRC32_peerChannel:
					channelIdList = AppendID(chatIdList, p.GetData2().GetChannelId())
				}
			case mtproto.TLConstructor_CRC32_messageService:
				m2 := m.To_MessageService()
				userIdList = AppendID(userIdList, m2.GetFromId())

				p := m2.GetToId()
				switch p.GetConstructor() {
				case mtproto.TLConstructor_CRC32_peerUser:
					userIdList = AppendID(userIdList, p.GetData2().GetUserId())
				case mtproto.TLConstructor_CRC32_peerChat:
					if p.GetData2().GetChatId() == 0 {
						glog.Info("chat_id = 0: ", m)
						continue
					}
					chatIdList = AppendID(chatIdList, p.GetData2().GetChatId())
				case mtproto.TLConstructor_CRC32_peerChannel:
					channelIdList = AppendID(chatIdList, p.GetData2().GetChannelId())
				}
			case mtproto.TLConstructor_CRC32_messageEmpty:
			}
		}
	}
	return
}

/*
  message: "@XXXXXXXX XXX @XXX 11111111" [STRING],
  ...
  entities: [ vector<0x0>
	{ inputMessageEntityMentionName
	  offset: 10 [INT],
	  length: 3 [INT],
	  user_id: { inputUser
		user_id: 607858518 [INT],
		access_hash: 17343137402047930393 [LONG],
	  },
	},
  ],
*/

func ParseMentions(message string) []*mtproto.MessageEntity {
	matches := regexp.MustCompile("@[0-9]+\\s*\\([^()@]+\\)").FindStringSubmatch(message)
	mentions := make([]*mtproto.MessageEntity, 0, len(matches))

	p := 0
	idx := 0
	for _, m := range matches {
		idx = strings.Index(message[p:], m)
		mentiton := &mtproto.TLMessageEntityMention{Data2: &mtproto.MessageEntity_Data{
			Offset: int32(p + idx),
			Length: int32(len(m)),
		}}
		mentions = append(mentions, mentiton.To_MessageEntity())
		p = idx
	}

	return mentions
}
