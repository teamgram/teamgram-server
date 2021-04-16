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

package updates

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (s *UpdatesServiceImpl) UpdatesGetChannelDifference(ctx context.Context, request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("updates.getChannelDifference#3173d78 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		lastPts  = request.GetPts()
		userList []*mtproto.User
		chatList []*mtproto.Chat
	)

	// var difference *mtproto.Updates_ChannelDifference

	channelId := request.GetChannel().GetData2().ChannelId
	channelLogic, _ := s.ChannelModel.NewChannelLogicById(channelId)
	participant := channelLogic.GetChannelParticipant(md.UserId)
	switch participant.GetConstructor() {
	case mtproto.TLConstructor_CRC32_channelParticipantsBanned:
		// TODO(@benqi):
		//banned := channel.MakeChannelBannedRights(participant.GetData2().GetBannedRights().To_ChannelBannedRights())
		//if banned.IsForbidden() {
		//	return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHANNEL_PRIVATE)
		//}
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHANNEL_PRIVATE)

	}

	channelDifference, err := sync_client.GetSyncClient().SyncGetChannelDifference(md.GetAuthId(), md.GetUserId(), lastPts, request.GetChannel())
	if err != nil {
		glog.Error("updates.getChannelDifference#3173d78 error - ", err)
		return nil, err
	}
	switch channelDifference.GetConstructor() {
	case mtproto.TLConstructor_CRC32_updates_differenceEmpty:

	case mtproto.TLConstructor_CRC32_updates_channelDifference:
		diff := channelDifference.To_UpdatesChannelDifference()

		userIdList, chatIdList, _ := message.PickAllIDListByMessages(diff.GetNewMessages())
		userList = s.UserModel.GetUserListByIdList(md.UserId, userIdList)
		chatList = s.ChatModel.GetChatListBySelfAndIDList(md.UserId, chatIdList)
		diff.SetUsers(userList)
		diff.SetChats(chatList)
	default:
		err = fmt.Errorf("not impl")
		return nil, err
	}

	glog.Infof("updates.getChannelDifference#3173d78 - reply: %s", logger.JsonDebugData(channelDifference))
	return channelDifference, nil
}
