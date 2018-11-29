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

package messages

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
)

// messages.addChatUser#f9a0aa09 chat_id:int user_id:InputUser fwd_limit:int = Updates;
func (s *MessagesServiceImpl) MessagesAddChatUser(ctx context.Context, request *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.addChatUser#f9a0aa09 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		err error
		addChatUserId int32
	)

	if request.GetUserId().GetConstructor() == mtproto.TLConstructor_CRC32_inputUserEmpty {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("messages.sendMessage#fa88427a - invalid peer", err)
		return nil, err
	}

	switch request.GetUserId().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserEmpty:
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		addChatUserId = md.UserId

	case mtproto.TLConstructor_CRC32_inputUser:
		addChatUserId = request.GetUserId().GetData2().GetUserId()
	}

	chatLogic, _ := s.ChatModel.NewChatLogicById(request.GetChatId())
	chatLogic.AddChatUser(md.UserId, addChatUserId)

	peer := &base.PeerUtil{
		PeerType: base.PEER_CHAT,
		PeerId:   chatLogic.GetChatId(),
	}

	addUserMessage := chatLogic.MakeAddUserMessage(md.UserId, addChatUserId)
	randomId := core.GetUUID()

	resultCB := func(pts, ptsCount int32, outBox *message.MessageBox2) (*mtproto.Updates, error) {
		syncUpdates := updates.NewUpdatesLogic(md.UserId)

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
		}}
		syncUpdates.AddUpdate(updateChatParticipants.To_Update())
		syncUpdates.AddUpdateNewMessage(pts, ptsCount, outBox.ToMessage(outBox.OwnerId))
		syncUpdates.AddUsers(s.UserModel.GetUserListByIdList(md.UserId, chatLogic.GetChatParticipantIdList()))
		syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

		syncUpdates.AddUpdateMessageId(outBox.MessageId, outBox.RandomId)

		return syncUpdates.ToUpdates(), nil
	}

	syncNotMeCB := func(pts, ptsCount int32, outBox *message.MessageBox2) (int64, *mtproto.Updates, error) {
		syncUpdates := updates.NewUpdatesLogic(md.UserId)

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
		}}
		syncUpdates.AddUpdate(updateChatParticipants.To_Update())
		syncUpdates.AddUpdateNewMessage(pts, ptsCount, outBox.ToMessage(outBox.OwnerId))
		syncUpdates.AddUsers(s.UserModel.GetUserListByIdList(md.UserId, chatLogic.GetChatParticipantIdList()))
		syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

		return md.AuthId, syncUpdates.ToUpdates(), nil
	}

	pushCB := func(pts, ptsCount int32, inBox *message.MessageBox2) (*mtproto.Updates, error) {
		pushUpdates := updates.NewUpdatesLogic(md.UserId)

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
		}}
		pushUpdates.AddUpdate(updateChatParticipants.To_Update())
		pushUpdates.AddUpdateNewMessage(pts, ptsCount, inBox.ToMessage(inBox.OwnerId))
		pushUpdates.AddUsers(s.UserModel.GetUserListByIdList(inBox.OwnerId, chatLogic.GetChatParticipantIdList()))
		pushUpdates.AddChat(chatLogic.ToChat(inBox.OwnerId))

		return pushUpdates.ToUpdates(), nil
	}

	replyUpdates, _ := s.MessageModel.SendMessage(
		md.UserId,
		peer,
		randomId,
		addUserMessage,
		resultCB,
		syncNotMeCB,
		pushCB)

	glog.Infof("messages.addChatUser#f9a0aa09 - reply: {%v}", replyUpdates)
	return replyUpdates, nil
}
