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
	"math"
)

/*
// messages.search#8614ef68 flags:#
	pee	r:InputPeer
	q:string
	from_id:flags.0?InputUser
	filter:MessagesFilter
	min_date:int
	max_date:int
	offset_id:int
	add_offset:int
	limit:int
	max_id:int
	min_id:int
	hash:int = messages.Messages;

{
    "constructor": "CRC32_messages_search_8614ef68",
    "peer": {
        "predicate_name": "inputPeerEmpty",
        "constructor": "CRC32_inputPeerEmpty"
    },
    "filter": {
        "predicate_name": "inputMessagesFilterPhoneCalls",
        "constructor": "CRC32_inputMessagesFilterPhoneCalls"
    },
    "limit": 50
}

	messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
	messages.messagesSlice#c8edce1e flags:# inexact:flags.1?true count:int next_rate:flags.0?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
	messages.channelMessages#99262e37 flags:# inexact:flags.1?true pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
	messages.messagesNotModified#74535f21 count:int = messages.Messages;
*/

/*
## 1. TL_messages_search
```
    public void loadMedia(final long uid, final int count, final int max_id, final int type, final int fromCache, final int classGuid) {
        final boolean isChannel = (int) uid < 0 && ChatObject.isChannel(-(int) uid, currentAccount);

        int lower_part = (int)uid;
        if (fromCache != 0 || lower_part == 0) {
            loadMediaDatabase(uid, count, max_id, type, classGuid, isChannel, fromCache);
        } else {
		   TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
		   req.limit = count;
		   req.offset_id = max_id;
		   if (type == MEDIA_PHOTOVIDEO) {
			   req.filter = new TLRPC.TL_inputMessagesFilterPhotoVideo();
		   } else if (type == MEDIA_FILE) {
			   req.filter = new TLRPC.TL_inputMessagesFilterDocument();
		   } else if (type == MEDIA_AUDIO) {
			   req.filter = new TLRPC.TL_inputMessagesFilterRoundVoice();
		   } else if (type == MEDIA_URL) {
			   req.filter = new TLRPC.TL_inputMessagesFilterUrl();
		   } else if (type == MEDIA_MUSIC) {
			   req.filter = new TLRPC.TL_inputMessagesFilterMusic();
		   }
```

## 2. TL_messages_search
```
                } else {
                    boolean missing = false;
                    for (int a = 0; a < counts.length; a++) {
                        if (counts[a] == -1 || old[a] == 1) {
                            final int type = a;

                            TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
                            req.limit = 1;
                            req.offset_id = 0;
                            if (a == MEDIA_PHOTOVIDEO) {
                                req.filter = new TLRPC.TL_inputMessagesFilterPhotoVideo();
                            } else if (a == MEDIA_FILE) {
                                req.filter = new TLRPC.TL_inputMessagesFilterDocument();
                            } else if (a == MEDIA_AUDIO) {
                                req.filter = new TLRPC.TL_inputMessagesFilterRoundVoice();
                            } else if (a == MEDIA_URL) {
                                req.filter = new TLRPC.TL_inputMessagesFilterUrl();
                            } else if (a == MEDIA_MUSIC) {
                                req.filter = new TLRPC.TL_inputMessagesFilterMusic();
                            }
```

## 3. TL_messages_search
```
                } else {
                    TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
                    req.filter = new TLRPC.TL_inputMessagesFilterChatPhotos();
                    req.limit = 80;
                    req.offset_id = 0;
                    req.q = "";
                    req.peer = MessagesController.getInstance(currentAccount).getInputPeer(id);
                    ConnectionsManager.getInstance(currentAccount).sendRequest(req, (response, error) -> onRequestComplete(locationKey, parentKey, response, true));
```

## 4. TL_messages_search
```
            } else if (did < 0) {
                TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
                req.filter = new TLRPC.TL_inputMessagesFilterChatPhotos();
                req.limit = count;
                req.offset_id = (int) max_id;
                req.q = "";
                req.peer = getInputPeer(did);
                int reqId = ConnectionsManager.getInstance(currentAccount).sendRequest(req, (response, error) -> {
                    if (error == null) {
                        TLRPC.messages_Messages messages = (TLRPC.messages_Messages) response;
                        TLRPC.TL_photos_photos res = new TLRPC.TL_photos_photos();
                        res.count = messages.count;
                        res.users.addAll(messages.users);
                        for (int a = 0; a < messages.messages.size(); a++) {
                            TLRPC.Message message = messages.messages.get(a);
                            if (message.action == null || message.action.photo == null) {
                                continue;
                            }
                            res.photos.add(message.action.photo);
                        }
                        processLoadedUserPhotos(res, did, count, max_id, false, classGuid);
                    }
                });
                ConnectionsManager.getInstance(currentAccount).bindRequestToGuid(reqId, classGuid);
```

## 5. TL_messages_search
```
            TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
            req.limit = 50;
            req.offset_id = max_id;
            if (currentType == 1) {
                req.filter = new TLRPC.TL_inputMessagesFilterDocument();
            } else if (currentType == 3) {
                req.filter = new TLRPC.TL_inputMessagesFilterUrl();
            } else if (currentType == 4) {
                req.filter = new TLRPC.TL_inputMessagesFilterMusic();
            }
            req.q = query;
            req.peer = MessagesController.getInstance(currentAccount).getInputPeer(uid);
            if (req.peer == null) {
                return;
            }
            final int currentReqId = ++lastReqId;
            searchesInProgress++;
            reqId = ConnectionsManager.getInstance(currentAccount).sendRequest(req, (response, error) -> {
                final ArrayList<MessageObject> messageObjects = new ArrayList<>();
                if (error == null) {

```

## 6. TL_messages_search
```
		TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
		req.limit = count;
		req.peer = new TLRPC.TL_inputPeerEmpty();
		req.filter = new TLRPC.TL_inputMessagesFilterPhoneCalls();
		req.q = "";
		req.offset_id = max_id;
		int reqId = ConnectionsManager.getInstance(currentAccount).sendRequest(req, (response, error) -> AndroidUtilities.runOnUIThread(() -> {
            if (error == null) {
                SparseArray<TLRPC.User> users = new SparseArray<>();
                TLRPC.messages_Messages msgs = (TLRPC.messages_Messages) response;
                endReached = msgs.messages.isEmpty();
                for (int a = 0; a < msgs.users.size(); a++) {
                    TLRPC.User user = msgs.users.get(a);
                    users.put(user.id, user);
                }
```
*/

// MessagesSearch
// messages.search#a0fda762 flags:# peer:InputPeer q:string from_id:flags.0?InputPeer top_msg_id:flags.1?int filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (c *MessagesCore) MessagesSearch(in *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if c.MD.IsBot {
		err := mtproto.ErrBotMethodInvalid
		c.Logger.Errorf("messages.search - error: %v", err)
		return nil, err
	}

	var (
		rValues  *mtproto.Messages_Messages
		offsetId = in.OffsetId
		limit    = in.Limit
		boxList  *message.Vector_MessageBox
		err      error
	)

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	if limit > 50 {
		limit = 50
	}

	peer := mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	switch peer.PeerType {
	case mtproto.PEER_EMPTY, mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		// TODO(@benqi): Not impl
		rValues = mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
		}).To_Messages_Messages()
	case mtproto.PEER_CHANNEL:
		rValues = mtproto.MakeTLMessagesChannelMessages(&mtproto.Messages_Messages{
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
		}).To_Messages_Messages()
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.search - error: %v", err)
		return nil, err
	}

	//TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
	//req.limit = count;
	//req.offset_id = max_id;
	//if (type == MEDIA_PHOTOVIDEO) {
	//	req.filter = new TLRPC.TL_inputMessagesFilterPhotoVideo();
	//} else if (type == MEDIA_FILE) {
	//	req.filter = new TLRPC.TL_inputMessagesFilterDocument();
	//} else if (type == MEDIA_AUDIO) {
	//	req.filter = new TLRPC.TL_inputMessagesFilterRoundVoice();
	//} else if (type == MEDIA_URL) {
	//	req.filter = new TLRPC.TL_inputMessagesFilterUrl();
	//} else if (type == MEDIA_MUSIC) {
	//	req.filter = new TLRPC.TL_inputMessagesFilterMusic();
	//}

	/*
		- 未使用

		inputMessagesFilterMyMentions#c1f8e69a = MessagesFilter;
		inputMessagesFilterGeo#e7026d0d = MessagesFilter;
		inputMessagesFilterContacts#e062db83 = MessagesFilter;
	*/
	filterType := mtproto.FromMessagesFilter(in.Filter)
	switch filterType {
	case mtproto.FilterPhotos:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterVideo:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterPhotoVideo:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_PHOTOVIDEO,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterDocument:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_FILE,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterUrl:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_URL,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterGif:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterVoice:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterMusic:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_MUSIC,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterChatPhotos:
		// TODO
	case mtproto.FilterPhoneCalls:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_PHONE_CALL,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterRoundVoice:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_AUDIO,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterRoundVideo:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterMyMentions:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterGeo:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterContacts:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterPinned:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByPinned(c.ctx, &message.TLMessageSearchByPinned{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterEmpty:
		if in.Q == "" {
			err = mtproto.ErrSearchQueryEmpty
			c.Logger.Errorf("messages.search - error: %v", err)
			return nil, err
		}

		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearch(c.ctx, &message.TLMessageSearch{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
			Q:        in.Q,
			Offset:   offsetId,
			Limit:    limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	default:
		// TODO
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	}

	//
	if peer.PeerType == mtproto.PEER_CHANNEL {
		rValues.Count = boxList.Length()
		//channelLogic, err := s.ChannelCore.NewChannelLogicById(ctx, peer.PeerId)
		//if err != nil {
		//	messages.Pts = channelLogic.Pts
		//}
	} else {

	}

	boxList.Visit(c.MD.UserId,
		func(messageList []*mtproto.Message) {
			rValues.Messages = messageList
		},
		func(userIdList []int64) {
			mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			rValues.Users = append(rValues.Users, mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			mChats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rValues.Chats = append(rValues.Chats, mChats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			//mChannels, _ := c.svcCtx.Dao.ChannelClient.ChannelGetChannelListByIdList(c.ctx,
			//	&channelpb.TLChannelGetChannelListByIdList{
			//		SelfUserId: c.MD.UserId,
			//		Id:         channelIdList,
			//	})
			//if len(mChannels.GetDatas()) > 0 {
			//	rValues.Chats = append(rValues.Chats, mChannels.GetDatas()...)
			//}
		})

	return rValues, nil
}
