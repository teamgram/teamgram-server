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
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

/*
## android client
- void setMessageObjectInternal(MessageObject messageObject)
```
    private void setMessageObjectInternal(MessageObject messageObject) {
        if (((messageObject.messageOwner.flags & TLRPC.MESSAGE_FLAG_HAS_VIEWS) != 0 || messageObject.messageOwner.replies != null) && !currentMessageObject.scheduled && !currentMessageObject.isSponsored()) {
            if (!currentMessageObject.viewsReloaded) {
                MessagesController.getInstance(currentAccount).addToViewsQueue(currentMessageObject);
                currentMessageObject.viewsReloaded = true;
            }
        }
```

- didReceivedNotification
```
        } else if (id == NotificationCenter.removeAllMessagesFromDialog) {
            long did = (Long) args[0];
            if (dialog_id == did) {
                if (threadMessageId != 0) {
                    if (forwardEndReached[0]) {
                        forwardEndReached[0] = false;
                        chatAdapter.notifyItemInserted(0);
                    }
                    getMessagesController().addToViewsQueue(threadMessageObject);
```
*/

// MessagesGetMessagesViews
// messages.getMessagesViews#5784d3e1 peer:InputPeer id:Vector<int> increment:Bool = messages.MessageViews;
func (c *MessagesCore) MessagesGetMessagesViews(in *mtproto.TLMessagesGetMessagesViews) (*mtproto.Messages_MessageViews, error) {
	type msgIdPeerIdPair struct {
		msgId  int32
		peerId int64
	}
	var (
		boxMsgList    *message.Vector_MessageBox
		channelIds    = make(map[int64][]int32, 0)
		reqIndexViews = make(map[int32]msgIdPeerIdPair, 0)
		views         = make(map[int64][]int32, 0)
		// increment     = mtproto.FromBool(in.GetIncrement())
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	)

	if len(in.Id) == 0 {
		err := mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("messages.getMessagesViews#5784d3e - error: %v, invalid request, len(id) == 0", err)
		return nil, err
	}

	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		switch peer.PeerType {
		case mtproto.PEER_USER:
			mutableUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(
				c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: []int64{c.MD.UserId, peer.PeerId},
				})
			if err != nil {
				c.Logger.Errorf("messages.getMessagesViews#5784d3e - error: %v, not found peer(%v) == 0", err, peer)
				return nil, err
			}
			if mutableUsers.Length() != 2 {
				err = mtproto.ErrPeerIdInvalid
				c.Logger.Errorf("messages.getMessagesViews#5784d3e - error: %v, not found peer(%v) == 0", err, peer)
				return nil, err
			}
		case mtproto.PEER_CHAT:
			mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(
				c.ctx,
				&chatpb.TLChatGetMutableChat{
					ChatId: peer.PeerId,
				})
			if err != nil {
				c.Logger.Errorf("messages.getMessagesViews#5784d3e - error: not found chat by id(%v) type", peer)
				return nil, mtproto.ErrChatIdInvalid
			}
			if _, ok := mutableChat.GetImmutableChatParticipant(c.MD.UserId); !ok {
				c.Logger.Errorf("messages.getMessagesViews#5784d3e - error: not in chat(%v) type", in.Id)
				return nil, mtproto.ErrChatIdInvalid
			}
		}

		var (
			err error
		)
		boxMsgList, err = c.svcCtx.Dao.MessageClient.MessageGetUserMessageList(
			c.ctx,
			&message.TLMessageGetUserMessageList{
				UserId: c.MD.UserId,
				IdList: in.Id,
			})
		if err != nil {
			c.Logger.Errorf("messages.getMessagesViews - error: %v", err)
			return nil, err
		}
		//// md.UserId, request.Id)
		//if len(request.Id) != len(boxMsgList) {
		//	log.Errorf("messages.getMessagesViews#5784d3e - error: boxMsgList empty by id(%v) type", request.Id)
		//	return nil, mtproto.ErrMsgIdInvalid
		//}
		// log.Debugf("user boxMsgList: %s", boxMsgList)

		boxMsgList.Walk(func(idx int, v *mtproto.MessageBox) {
			fwdFrom := v.Message.GetFwdFrom()
			if fwdFrom == nil {
				// TODO:
				// c.Logger.Errorf("messages.getMessagesViews#5784d3e - error: boxMsgList empty by id(%v) type", in.Id)
				// err = mtproto.ErrMsgIdInvalid
				// return
			} else {
				// TODO
				//if idList, ok := channelIds[fwdFrom.GetChannelId().GetValue()]; !ok {
				//	channelIds[fwdFrom.GetChannelId().GetValue()] = []int32{fwdFrom.GetChannelPost().GetValue()}
				//} else {
				//	idList = append(idList, fwdFrom.GetChannelPost().GetValue())
				//	channelIds[fwdFrom.GetChannelId().GetValue()] = idList
				//}
				//reqIndexViews[box.MessageId] = int64(fwdFrom.GetChannelId().GetValue())<<32 | int64(fwdFrom.GetChannelPost().GetValue())
			}
		})
	case mtproto.PEER_CHANNEL:
		//channel, err := c.svcCtx.Dao.ChannelClient.ChannelGetMutableChannel(
		//	c.ctx,
		//	&channelpb.TLChannelGetMutableChannel{
		//		ChannelId: peer.PeerId,
		//		Id:        []int64{c.MD.UserId},
		//	})
		//if err != nil {
		//	c.Logger.Errorf("messages.getMessagesViews#5784d3e - error: not found chat by id(%v) type", peer)
		//	return nil, mtproto.ErrChannelInvalid
		//}
		//
		///*{"user_id":1,"send_user_id":2,"message_id":70,"dialog_id":-1073741834,"dialog_message_id":0,"message_data_id":4611686061377060934,"random_id":-4236442189977411877,"pts":0,"pts_count":0,"message_filter_type":0,"message_box_type":2,"message_type":2,"message":{"predicate_name":"message","id":70,"out":true,"post":true,"to_id":{"predicate_name":"peerChannel","channel_id":1073741834},"date":1583378362,"message":"2233","views":{"value":1},"post_author":{"value":"璧君"}},"views":0,"reply_owner_id":0}*/
		//boxMsgList, err = c.svcCtx.Dao.MessageClient.MessageGetChannelMessageList(
		//	c.ctx,
		//	&message.TLMessageGetChannelMessageList{
		//		UserId:    c.MD.UserId,
		//		ChannelId: peer.PeerId,
		//		IdList:    in.Id,
		//	})
		//if err != nil {
		//	c.Logger.Errorf("messages.getMessagesViews - error: %v", err)
		//	return nil, err
		//}
		//
		////log.Debugf("messages.getMessagesViews#5784d3e - channel boxMsgList: %s", boxMsgList)
		//
		//if channel.Megagroup() {
		//	boxMsgList.Walk(func(idx int, v *mtproto.MessageBox) {
		//		fwdFrom := v.Message.GetFwdFrom()
		//		if fwdFrom == nil {
		//			// TODO:
		//			// c.Logger.Errorf("messages.getMessagesViews#5784d3e - error: boxMsgList empty by id(%v) type", in.Id)
		//			// err = mtproto.ErrMsgIdInvalid
		//			return
		//		} else {
		//			// TODO
		//			//if idList, ok := channelIds[fwdFrom.GetChannelId().GetValue()]; !ok {
		//			//	channelIds[fwdFrom.GetChannelId().GetValue()] = []int32{fwdFrom.GetChannelPost().GetValue()}
		//			//} else {
		//			//	idList = append(idList, box.Message.GetFwdFrom().GetChannelPost().GetValue())
		//			//	channelIds[fwdFrom.GetChannelId().GetValue()] = idList
		//			//}
		//			//reqIndexViews[box.MessageId] = int64(fwdFrom.GetChannelId().GetValue())<<32 | int64(fwdFrom.GetChannelPost().GetValue())
		//		}
		//	})
		//} else if channel.Broadcast() {
		//	boxMsgList.Walk(func(idx int, v *mtproto.MessageBox) {
		//		fwdFrom := v.Message.GetFwdFrom()
		//		if fwdFrom == nil {
		//			if idList, ok := channelIds[channel.Id()]; !ok {
		//				channelIds[channel.Id()] = []int32{v.MessageId}
		//				// log.Debugf("add boxId: (%d, %d, %v)", channel.Channel.Id, box.MessageId, channelIds)
		//			} else {
		//				idList = append(idList, v.MessageId)
		//				channelIds[channel.Channel.Id] = idList
		//				// log.Debugf("add boxId: (%d, %d, %v)", channel.Channel.Id, box.MessageId, channelIds)
		//			}
		//			reqIndexViews[v.MessageId] = msgIdPeerIdPair{msgId: v.MessageId, peerId: channel.Id()}
		//		} else {
		//			// TODO
		//			//if idList, ok := channelIds[fwdFrom.GetChannelId().GetValue()]; !ok {
		//			//	channelIds[fwdFrom.GetChannelId().GetValue()] = []int32{fwdFrom.GetChannelPost().GetValue()}
		//			//	// log.Debugf("add boxId: (%d, %d, %v)", channel.Channel.Id, box.MessageId, channelIds)
		//			//} else {
		//			//	idList = append(idList, box.Message.GetFwdFrom().GetChannelPost().GetValue())
		//			//	channelIds[fwdFrom.GetChannelId().GetValue()] = idList
		//			//	// log.Debugf("add boxId: (%d, %d, %v)", channel.Channel.Id, box.MessageId, channelIds)
		//			//}
		//			//reqIndexViews[box.MessageId] = int64(fwdFrom.GetChannelId().GetValue())<<32 | int64(fwdFrom.GetChannelPost().GetValue())
		//		}
		//	})
		//} else {
		//	c.Logger.Errorf("messages.getMessagesViews#5784d3e -  error: invalid peer(%v) type", peer)
		//	return nil, mtproto.ErrMsgIdInvalid
		//}
		c.Logger.Errorf("messages.getRecentLocations blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		c.Logger.Errorf("messages.getMessagesViews#5784d3e -  error: invalid peer(%v) type", peer)
		return nil, mtproto.ErrInputRequestInvalid
	}

	//for k, v := range channelIds {
	//	if increment {
	//		rViews, _ := c.svcCtx.Dao.ChannelClient.ChannelIncrementChannelMessagesViews(
	//			c.ctx,
	//			&channelpb.TLChannelIncrementChannelMessagesViews{
	//				ChannelId: k,
	//				IdList:    v,
	//			})
	//		if rViews.GetDatas() != nil {
	//			views[k] = rViews.GetDatas()
	//		}
	//	} else {
	//		rViews, _ := c.svcCtx.Dao.ChannelClient.ChannelGetChannelMessagesViews(
	//			c.ctx,
	//			&channelpb.TLChannelGetChannelMessagesViews{
	//				ChannelId: k,
	//				IdList:    v,
	//			})
	//		if rViews.GetDatas() != nil {
	//			views[k] = rViews.GetDatas()
	//		}
	//	}
	//}

	rViews := mtproto.MakeTLMessagesMessageViews(&mtproto.Messages_MessageViews{
		Views: []*mtproto.MessageViews{},
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Messages_MessageViews()

	for _, id := range in.Id {
		if indexV, ok := reqIndexViews[id]; !ok {
			rViews.Views = append(rViews.Views, mtproto.MakeTLMessageViews(&mtproto.MessageViews{
				Views:    nil,
				Forwards: nil,
				Replies:  nil,
			}).To_MessageViews())
		} else {
			// rcId := int32(indexV >> 32)
			// rmId := int32(indexV & 0xffffffff)
			mId := int32(0)
			for i, id2 := range channelIds[indexV.peerId] {
				if id2 == indexV.msgId {
					mId = views[indexV.peerId][i]
					break
				}
			}
			if mId == 0 {
				rViews.Views = append(rViews.Views, mtproto.MakeTLMessageViews(&mtproto.MessageViews{
					Views:    nil,
					Forwards: nil,
					Replies:  nil,
				}).To_MessageViews())
			} else {
				rViews.Views = append(rViews.Views, mtproto.MakeTLMessageViews(&mtproto.MessageViews{
					Views:    mtproto.MakeFlagsInt32(mId),
					Forwards: nil,
					Replies:  nil,
					//Replies: mtproto.MakeTLMessageReplies(&mtproto.MessageReplies{
					//	Comments:       true,
					//	Replies:        0,
					//	RepliesPts:     0,
					//	RecentRepliers: nil,
					//	ChannelId:      mtproto.MakeFlagsInt32(rcId),
					//	MaxId:          nil,
					//	ReadMaxId:      nil,
					//}).To_MessageReplies(),
				}).To_MessageViews())
			}
		}
	}

	boxMsgList.Walk(func(idx int, v *mtproto.MessageBox) {
		for _, v2 := range rViews.GetViews() {
			if v2.GetReplies() != nil {
				v2.Replies = v2.GetReplies()
			}
		}
	})

	return rViews, nil
}
