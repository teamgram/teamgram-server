// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetMessages
// messages.getMessages#63c66506 id:Vector<InputMessage> = messages.Messages;
func (c *MessagesCore) MessagesGetMessages(in *tg.TLMessagesGetMessages) (*tg.MessagesMessages, error) {
	ids := make([]int32, 0, len(in.Id_VECTORINPUTMESSAGE)+len(in.Id_VECTORINT32))
	for _, msg := range in.Id_VECTORINPUTMESSAGE {
		switch x := msg.(type) {
		case *tg.TLInputMessageID:
			ids = append(ids, x.Id)
		case *tg.TLInputMessageReplyTo:
			ids = append(ids, x.Id)
		case *tg.TLInputMessagePinned:
			ids = append(ids, 1)
		}
	}
	ids = append(ids, in.Id_VECTORINT32...)

	if c.svcCtx != nil && c.svcCtx.MessageClient != nil && c.MD != nil && c.MD.UserId != 0 {
		boxes, err := c.svcCtx.MessageClient.MessageGetUserMessageList(c.ctx, &message.TLMessageGetUserMessageList{
			UserId: c.MD.UserId,
			IdList: ids,
		})
		if err != nil {
			c.Logger.Errorf("messages.getMessages - MessageGetUserMessageList error: %v", err)
			return nil, err
		}
		if boxes != nil {
			return makeBffMessagesMessagesFromBoxes(boxes.Datas), nil
		}
	}

	return makeBffMessagesMessagesByIDs(tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 0}), ids, false), nil
}
