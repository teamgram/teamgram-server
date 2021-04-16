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

package channels

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
	update2 "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// channels.inviteToChannel#199f3a6c channel:InputChannel users:Vector<InputUser> = Updates;
func (s *ChannelsServiceImpl) ChannelsInviteToChannel(ctx context.Context, request *mtproto.TLChannelsInviteToChannel) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.inviteToChannel#199f3a6c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	if request.Channel.Constructor == mtproto.TLConstructor_CRC32_inputChannelEmpty {
		// TODO(@benqi): chatUser不能是inputUser和inputUserSelf
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("channels.exportInvite#c7560885 - error: ", err, "; InputPeer invalid")
		return nil, err
	}

	channelLogic, err := s.ChannelModel.NewChannelLogicById(request.GetChannel().GetData2().GetChannelId())
	if err != nil {
		glog.Error("channels.inviteToChannel#199f3a6c - error: ", err)
		return nil, err
	}

	updateChannel := &mtproto.TLUpdateChannel{Data2: &mtproto.Update_Data{
		ChannelId: channelLogic.GetChannelId(),
	}}

	addMemberMessage := channelLogic.MakeAddUserMessage(md.UserId, channelLogic.GetChannelId())
	peer := &base.PeerUtil{
		PeerType: base.PEER_CHANNEL,
		PeerId:   channelLogic.GetChannelId(),
	}
	pts := int32(core.NextChannelNPtsId(channelLogic.GetChannelId(), 2))
	ptsCount := int32(1)
	var boxList []*message.MessageBox2
	s.MessageModel.SendInternalMessage(md.UserId, peer, channelLogic.RandomId, false, addMemberMessage, func(i int32, box2 *message.MessageBox2) {
		// TopMessage
		channelLogic.SetTopMessage(box2.MessageId)
		boxList = append(boxList, box2)
	})
	for _, u := range request.Users {
		if u.GetConstructor() == mtproto.TLConstructor_CRC32_inputUserEmpty ||
			u.GetConstructor() == mtproto.TLConstructor_CRC32_inputUserSelf {
			// TODO(@benqi): handle inputUserSelf
			continue
		}
		channelLogic.InviteToChannel(md.UserId, u.GetData2().GetUserId())
		s.DialogModel.InsertOrChannelUpdateDialog(u.GetData2().GetUserId(), base.PEER_CHANNEL, channelLogic.GetChannelId(), channelLogic.TopMessage)

		pushUpdates := update2.NewUpdatesLogic(u.GetData2().GetUserId())
		pushUpdates.AddUpdate(updateChannel.To_Update())
		pushUpdates.AddUpdateNewChannelMessage(pts, ptsCount, boxList[0].ToMessage(u.Data2.GetUserId()))
		pushUpdates.AddChat(channelLogic.ToChannel(u.GetData2().GetUserId()))
		pushUpdates.AddUsers(s.UserModel.GetUserListByIdList(boxList[0].OwnerId, channelLogic.GetChannelParticipantIdList(md.UserId)))
		sync_client.GetSyncClient().PushChannelUpdates(channelLogic.GetChannelId(), u.GetData2().GetUserId(), pushUpdates.ToUpdates())
	}

	replyUpdates := update2.NewUpdatesLogic(md.UserId)
	replyUpdates.AddChat(channelLogic.ToChannel(md.UserId))

	reply := replyUpdates.ToUpdates()
	glog.Infof("channels.inviteToChannel#199f3a6c - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}
