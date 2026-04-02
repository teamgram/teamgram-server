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

var _ *tg.Bool

// MessageGetUserMessage
// message.getUserMessage user_id:long id:int = MessageBox;
func (c *MessageCore) MessageGetUserMessage(in *message.TLMessageGetUserMessage) (*tg.MessageBox, error) {
	return makePlaceholderMessageBox(in.UserId, tg.PEER_USER, in.UserId, in.Id), nil
}

func makePlaceholderMessageBox(userID int64, peerType int32, peerID int64, messageID int32) *tg.MessageBox {
	return tg.MakeTLMessageBox(&tg.TLMessageBox{
		UserId:       userID,
		MessageId:    messageID,
		SenderUserId: userID,
		PeerType:     peerType,
		PeerId:       peerID,
		Message: tg.MakeTLMessage(&tg.TLMessage{
			Out:     true,
			Id:      messageID,
			Date:    int32(messageID),
			Message: "placeholder",
		}),
		Pts:      1,
		PtsCount: 1,
		Reaction: "",
	}).ToMessageBox()
}
