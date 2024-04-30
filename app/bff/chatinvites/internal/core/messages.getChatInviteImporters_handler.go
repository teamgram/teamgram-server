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
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// MessagesGetChatInviteImporters
// messages.getChatInviteImporters#df04dd4e flags:# requested:flags.0?true peer:InputPeer link:flags.1?string q:flags.2?string offset_date:int offset_user:InputUser limit:int = messages.ChatInviteImporters;
func (c *ChatInvitesCore) MessagesGetChatInviteImporters(in *mtproto.TLMessagesGetChatInviteImporters) (*mtproto.Messages_ChatInviteImporters, error) {
	var (
		peer       = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		offsetPeer = mtproto.FromInputUser(c.MD.UserId, in.OffsetUser)
		link       *wrapperspb.StringValue
		limit      = in.GetLimit()
	)

	if in.Link != nil {
		link = in.GetLink()
	} else {
		// link = mtproto.MakeFlagsString(in.Link)
	}

	if in.GetQ().GetValue() != "" && link.GetValue() != "" {
		// TODO: add ErrSearchWithLinkNotSupported
		err := mtproto.ErrSearchWithLinkNotSupported
		c.Logger.Errorf("messages.getChatInviteImporters - error: ", err)
		return nil, err
	}

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getChatInviteImporters - error: ", err)
		return nil, err
	}

	if limit == 0 {
		limit = 50
	}

	rValues := mtproto.MakeTLMessagesChatInviteImporters(&mtproto.Messages_ChatInviteImporters{
		Count:     0,
		Importers: []*mtproto.ChatInviteImporter{},
		Users:     []*mtproto.User{},
	}).To_Messages_ChatInviteImporters()

	inviteImporters, err := c.svcCtx.Dao.ChatClient.ChatGetChatInviteImporters(c.ctx, &chatpb.TLChatGetChatInviteImporters{
		SelfId:     c.MD.UserId,
		ChatId:     peer.PeerId,
		Requested:  in.Requested,
		Link:       link,
		Q:          in.Q,
		OffsetDate: in.OffsetDate,
		OffsetUser: offsetPeer.PeerId,
		Limit:      limit,
	})
	if err != nil {
		c.Logger.Errorf("messages.getChatInviteImporters - error: ", err)
		return nil, err
	}

	rValues.Importers = inviteImporters.Datas

	if len(rValues.Importers) == 0 {
		return rValues, nil
	}

	rValues.Count = int32(len(rValues.Importers))
	idHelper := mtproto.NewIDListHelper(c.MD.UserId)
	for _, a := range rValues.Importers {
		idHelper.AppendUsers(a.UserId)
	}

	idHelper.Visit(func(userIdList []int64) {
		users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: userIdList,
		})
		rValues.Users = users.GetUserListByIdList(c.MD.UserId, userIdList...)
	}, nil, nil)

	return rValues, nil
}
