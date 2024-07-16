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
	"google.golang.org/protobuf/proto"
	"time"

	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// MessagesEditMessage
// messages.editMessage#48f71778 flags:# no_webpage:flags.1?true peer:InputPeer id:int message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.15?int = Updates;
func (c *MessagesCore) MessagesEditMessage(in *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error) {
	var (
		peer         = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		editMessages *message.Vector_MessageBox
		dstMessage   *mtproto.MessageBox
		err          error
	)

	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		if peer.PeerType == mtproto.PEER_SELF {
			peer.PeerType = mtproto.PEER_USER
		}
		if peer.PeerType == mtproto.PEER_USER {
			//if !md.IsBot {
			//	hasBot = s.UserFacade.IsBot(ctx, peer.PeerId)
			//}
		}
		editMessages, _ = c.svcCtx.MessageClient.MessageGetUserMessageList(c.ctx, &message.TLMessageGetUserMessageList{
			UserId: c.MD.UserId,
			IdList: []int32{in.Id},
		})
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("messages.editMessage blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		c.Logger.Errorf("invalid peer: %v", in.Peer)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	if len(editMessages.GetDatas()) != 1 {
		err = mtproto.ErrMessageEmpty
		c.Logger.Errorf("messages.editMessage - emptyMessage(%d)", in.Id)
		return nil, err
	}

	// TODO(@benqi): check
	// editMessage := editMessages[0]
	if editMessages.Datas[0].SenderUserId != c.MD.UserId {
		err = mtproto.ErrMessageAuthorRequired
		c.Logger.Errorf("messages.editMessage - emptyMessage(%d)", in.Id)
		return nil, err
	}
	// ...
	dstMessage = proto.Clone(editMessages.Datas[0]).(*mtproto.MessageBox)

	outMessage := editMessages.Datas[0].Message
	// edit_date
	outMessage.EditDate = &wrapperspb.Int32Value{Value: int32(time.Now().Unix())}
	outMessage.EditHide = false

	// entities
	if in.Entities != nil {
		outMessage.Entities = in.Entities
	}

	// reply_markup
	if in.ReplyMarkup != nil {
		outMessage.ReplyMarkup = in.ReplyMarkup
	}

	if in.Media != nil {
		outMessage.Media, err = c.makeMediaByInputMedia(in.Media)
		if err != nil {
			c.Logger.Errorf("messages.editMessage - media error: %v", err)
			return nil, err
		}
	}
	// message
	if in.Message != nil {
		if in.Message.Value == "" {
			err = mtproto.ErrMessageEmpty
			c.Logger.Errorf("message empty: %v", err)
			return nil, err
		}
		outMessage.Message = in.Message.Value
		outMessage.Entities = nil
		//outMessage, _ = c.fixMessageEntities(c.MD.UserId, peer, in.NoWebpage, outMessage, func() bool {
		//	hasBot := c.MD.IsBot
		//	if !hasBot {
		//		//isBot, _ := c.svcCtx.Dao.UserClient.UserIsBot(c.ctx, &userpb.TLUserIsBot{
		//		//	Id: peer.PeerId,
		//		//})
		//		//hasBot = mtproto.FromBool(isBot)
		//	}
		//
		//	return hasBot
		//})
	}

	rUpdates, err := c.svcCtx.Dao.MsgClient.MsgEditMessageV2(c.ctx, &msgpb.TLMsgEditMessageV2{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.PermAuthKeyId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		EditType:  msgpb.EditTypeNormal,
		NewMessage: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    in.NoWebpage,
			Background:   false,
			RandomId:     0,
			Message:      outMessage,
			ScheduleDate: in.ScheduleDate,
		}).To_OutboxMessage(),
		DstMessage: dstMessage,
	})
	if err != nil {
		c.Logger.Errorf("messages.editMessage - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}
