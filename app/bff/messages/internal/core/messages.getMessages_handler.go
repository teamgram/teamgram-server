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
	"errors"

	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetMessages
// messages.getMessages#63c66506 id:Vector<InputMessage> = messages.Messages;
func (c *MessagesCore) MessagesGetMessages(in *tg.TLMessagesGetMessages) (*tg.MessagesMessages, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	idList, err := getMessagesIDList(in)
	if err != nil {
		return nil, err
	}

	r := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: make([]tg.MessageClazz, 0, len(idList)),
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesMessages()
	if len(idList) == 0 {
		return r, nil
	}

	messages, indexes := getMessagesEmptyMessages(idList)
	var getMessagesClient sendMessageClient = c.svcCtx.Repo.MsgClient
	boxes, err := getMessagesClient.MsgGetUserMessageList(c.ctx, &msg.TLMsgGetUserMessageList{
		UserId: md.UserId,
		IdList: idList,
	})
	if err != nil {
		if !errors.Is(err, msg.ErrMsgIdInvalid) {
			c.Logger.Errorf("messages.getMessages - msg error: self_user_id: %d, id_list: %v, err: %v",
				md.UserId, idList, err)
			return nil, mapMsgSendError(err)
		}
		if err = c.fillGetMessagesIndividually(getMessagesClient, md.UserId, idList, messages, indexes); err != nil {
			return nil, err
		}
	} else {
		fillGetMessagesFromBoxes(messages, indexes, boxes)
	}

	if full, ok := r.ToMessagesMessages(); ok {
		full.Messages = messages
	}
	if err = userprojection.FillMessagesMessagesUsers(c.ctx, c.svcCtx.Repo.UserClient, md.UserId, r, userprojection.MissingStoredReference); err != nil {
		return nil, err
	}
	return r, nil
}

func getMessagesIDList(in *tg.TLMessagesGetMessages) ([]int32, error) {
	idList := make([]int32, 0, len(in.Id_VECTORINPUTMESSAGE)+len(in.Id_VECTORINT32))
	for _, id := range in.Id_VECTORINPUTMESSAGE {
		switch id := id.(type) {
		case *tg.TLInputMessageID:
			idList = append(idList, id.Id)
		case *tg.TLInputMessageReplyTo:
			idList = append(idList, id.Id)
		case *tg.TLInputMessagePinned:
		case *tg.TLInputMessageCallbackQuery:
		default:
			return nil, tg.ErrInputConstructorInvalid
		}
	}
	idList = append(idList, in.Id_VECTORINT32...)
	return idList, nil
}

func getMessagesEmptyMessages(idList []int32) ([]tg.MessageClazz, map[int32][]int) {
	messages := make([]tg.MessageClazz, len(idList))
	indexes := make(map[int32][]int, len(idList))
	for i, id := range idList {
		messages[i] = tg.MakeTLMessageEmpty(&tg.TLMessageEmpty{Id: id})
		indexes[id] = append(indexes[id], i)
	}
	return messages, indexes
}

func (c *MessagesCore) fillGetMessagesIndividually(client sendMessageClient, userID int64, idList []int32, messages []tg.MessageClazz, indexes map[int32][]int) error {
	for _, id := range idList {
		box, err := client.MsgGetUserMessage(c.ctx, &msg.TLMsgGetUserMessage{
			UserId: userID,
			Id:     id,
		})
		if err != nil {
			if errors.Is(err, msg.ErrMsgIdInvalid) {
				continue
			}
			c.Logger.Errorf("messages.getMessages - msg error: self_user_id: %d, id: %d, err: %v",
				userID, id, err)
			return mapMsgSendError(err)
		}
		fillGetMessagesFromBox(messages, indexes, box)
	}
	return nil
}

func fillGetMessagesFromBoxes(messages []tg.MessageClazz, indexes map[int32][]int, boxes *msg.VectorMessageBox) {
	if boxes == nil {
		return
	}
	for _, box := range boxes.Datas {
		fillGetMessagesFromBox(messages, indexes, box)
	}
}

func fillGetMessagesFromBox(messages []tg.MessageClazz, indexes map[int32][]int, box *tg.MessageBox) {
	if box == nil || box.Message == nil {
		return
	}
	for _, i := range indexes[box.MessageId] {
		messages[i] = box.Message
	}
}
