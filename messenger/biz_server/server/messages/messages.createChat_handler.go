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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
)

/*
   public static void showAddUserAlert(String error, final BaseFragment fragment, boolean isChannel) {
        if (error == null || fragment == null || fragment.getParentActivity() == null) {
            return;
        }
        AlertDialog.Builder builder = new AlertDialog.Builder(fragment.getParentActivity());
        builder.setTitle(LocaleController.getString("AppName", R.string.AppName));
        switch (error) {
            case "PEER_FLOOD":
                builder.setMessage(LocaleController.getString("NobodyLikesSpam2", R.string.NobodyLikesSpam2));
                builder.setNegativeButton(LocaleController.getString("MoreInfo", R.string.MoreInfo), (dialogInterface, i) -> MessagesController.getInstance(fragment.getCurrentAccount()).openByUserName("spambot", fragment, 1));
                break;
            case "USER_BLOCKED":
            case "USER_BOT":
            case "USER_ID_INVALID":
                if (isChannel) {
                    builder.setMessage(LocaleController.getString("ChannelUserCantAdd", R.string.ChannelUserCantAdd));
                } else {
                    builder.setMessage(LocaleController.getString("GroupUserCantAdd", R.string.GroupUserCantAdd));
                }
                break;
            case "USERS_TOO_MUCH":
                if (isChannel) {
                    builder.setMessage(LocaleController.getString("ChannelUserAddLimit", R.string.ChannelUserAddLimit));
                } else {
                    builder.setMessage(LocaleController.getString("GroupUserAddLimit", R.string.GroupUserAddLimit));
                }
                break;
            case "USER_NOT_MUTUAL_CONTACT":
                if (isChannel) {
                    builder.setMessage(LocaleController.getString("ChannelUserLeftError", R.string.ChannelUserLeftError));
                } else {
                    builder.setMessage(LocaleController.getString("GroupUserLeftError", R.string.GroupUserLeftError));
                }
                break;
            case "ADMINS_TOO_MUCH":
                if (isChannel) {
                    builder.setMessage(LocaleController.getString("ChannelUserCantAdmin", R.string.ChannelUserCantAdmin));
                } else {
                    builder.setMessage(LocaleController.getString("GroupUserCantAdmin", R.string.GroupUserCantAdmin));
                }
                break;
            case "BOTS_TOO_MUCH":
                if (isChannel) {
                    builder.setMessage(LocaleController.getString("ChannelUserCantBot", R.string.ChannelUserCantBot));
                } else {
                    builder.setMessage(LocaleController.getString("GroupUserCantBot", R.string.GroupUserCantBot));
                }
                break;
            case "USER_PRIVACY_RESTRICTED":
                if (isChannel) {
                    builder.setMessage(LocaleController.getString("InviteToChannelError", R.string.InviteToChannelError));
                } else {
                    builder.setMessage(LocaleController.getString("InviteToGroupError", R.string.InviteToGroupError));
                }
                break;
            case "USERS_TOO_FEW":
                builder.setMessage(LocaleController.getString("CreateGroupError", R.string.CreateGroupError));
                break;
            case "USER_RESTRICTED":
                builder.setMessage(LocaleController.getString("UserRestricted", R.string.UserRestricted));
                break;
            case "YOU_BLOCKED_USER":
                builder.setMessage(LocaleController.getString("YouBlockedUser", R.string.YouBlockedUser));
                break;
            case "CHAT_ADMIN_BAN_REQUIRED":
            case "USER_KICKED":
                builder.setMessage(LocaleController.getString("AddAdminErrorBlacklisted", R.string.AddAdminErrorBlacklisted));
                break;
            case "CHAT_ADMIN_INVITE_REQUIRED":
                builder.setMessage(LocaleController.getString("AddAdminErrorNotAMember", R.string.AddAdminErrorNotAMember));
                break;
            case "USER_ADMIN_INVALID":
                builder.setMessage(LocaleController.getString("AddBannedErrorAdmin", R.string.AddBannedErrorAdmin));
                break;
            default:
                builder.setMessage(LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error);
                break;
        }
        builder.setPositiveButton(LocaleController.getString("OK", R.string.OK), null);
        fragment.showDialog(builder.create(), true, null);
    }
 */

func NotSelfByInputUser(selfUserId int32, user *mtproto.InputUser) bool {
	if user.GetConstructor() == mtproto.TLConstructor_CRC32_inputUser {
		return selfUserId != user.GetData2().GetUserId()
	}
	return false
}

// messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;
func (s *MessagesServiceImpl) MessagesCreateChat(ctx context.Context, request *mtproto.TLMessagesCreateChat) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.createChat#9cb126e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//randomId := md.ClientMsgId

	var (
		err error
		chatUserIdList []int32
		chatTitle = request.GetTitle()
	)

	// check chat title
	if chatTitle == "" {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_CHAT_TITLE)
		glog.Error("messages.createChat#9cb126e - error: ", err, "; chat title empty")
		return nil, err
	}

	// check user too much
	if len(request.GetUsers()) > 200 {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERS_TOO_MUCH)
		glog.Error("messages.createChat#9cb126e - error: ", err)
		return nil, err
	}

	// check len(users)
	chatUserIdList = make([]int32, 0, len(request.GetUsers()))
	for _, u := range request.GetUsers() {
		if NotSelfByInputUser(md.UserId, u) {
			// TODO(@benqi): check access_hash
			chatUserIdList = append(chatUserIdList, u.GetData2().GetUserId())
		} else {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			glog.Error("messages.createChat#9cb126e - error: ", err, "; invalid inputPeer")
			return nil, err
		}
	}

	chat, _ := s.ChatModel.CreateChat(md.UserId, chatUserIdList, chatTitle)

	peer := &base.PeerUtil{
		PeerType: base.PEER_CHAT,
		PeerId:   chat.GetChatId(),
	}

	createChatMessage := chat.MakeCreateChatMessage(md.UserId)
	randomId := core.GetUUID()

	resultCB := func(pts, ptsCount int32, outBox *message.MessageBox2) (*mtproto.Updates, error) {
		syncUpdates := updates.NewUpdatesLogic(md.UserId)

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chat.GetChatParticipants().To_ChatParticipants(),
		}}
		syncUpdates.AddUpdate(updateChatParticipants.To_Update())
		syncUpdates.AddUpdateNewMessage(pts, ptsCount, outBox.ToMessage(md.UserId))
		syncUpdates.AddUsers(s.UserModel.GetUserListByIdList(md.UserId, chat.GetChatParticipantIdList()))
		syncUpdates.AddChat(chat.ToChat(md.UserId))
		syncUpdates.AddUpdateMessageId(outBox.MessageId, outBox.RandomId)

		return syncUpdates.ToUpdates(), nil
	}

	syncNotMeCB := func(pts, ptsCount int32, outBox *message.MessageBox2) (int64, *mtproto.Updates, error) {
		syncUpdates := updates.NewUpdatesLogic(md.UserId)

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chat.GetChatParticipants().To_ChatParticipants(),
		}}
		syncUpdates.AddUpdate(updateChatParticipants.To_Update())
		syncUpdates.AddUpdateNewMessage(pts, ptsCount, outBox.ToMessage(md.UserId))
		syncUpdates.AddUsers(s.UserModel.GetUserListByIdList(md.UserId, chat.GetChatParticipantIdList()))
		syncUpdates.AddChat(chat.ToChat(md.UserId))

		return md.AuthId, syncUpdates.ToUpdates(), nil
	}

	pushCB := func(pts, ptsCount int32, inBox *message.MessageBox2) (*mtproto.Updates, error) {
		pushUpdates := updates.NewUpdatesLogic(md.UserId)
		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chat.GetChatParticipants().To_ChatParticipants(),
		}}
		pushUpdates.AddUpdate(updateChatParticipants.To_Update())
		pushUpdates.AddUpdateNewMessage(pts, ptsCount, inBox.ToMessage(inBox.OwnerId))
		pushUpdates.AddUsers(s.UserModel.GetUserListByIdList(inBox.OwnerId, chat.GetChatParticipantIdList()))
		pushUpdates.AddChat(chat.ToChat(inBox.OwnerId))

		return pushUpdates.ToUpdates(), nil
	}

	replyUpdates, _ := s.MessageModel.SendMessage(
		md.UserId,
		peer,
		randomId,
		createChatMessage,
		resultCB,
		syncNotMeCB,
		pushCB)

	glog.Infof("messages.createChat#9cb126e - reply: {%v}", replyUpdates)
	return replyUpdates, nil
}
