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
	updates "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// channels.createChannel#f4893d7f flags:# broadcast:flags.0?true megagroup:flags.1?true title:string about:string = Updates;
func (s *ChannelsServiceImpl) ChannelsCreateChannel(ctx context.Context, request *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.createChannel - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// 1. 创建channel
	channel, _ := s.ChannelModel.NewChannelLogicByCreateChannel(md.UserId, request.GetBroadcast(), request.GetMegagroup(), request.GetTitle(), request.GetAbout(), md.ClientMsgId)

	peer := &base.PeerUtil{
		PeerType: base.PEER_CHANNEL,
		PeerId:   channel.GetChannelId(),
	}
	createChannelMessage := channel.MakeCreateChannelMessage(md.UserId)
	randomId := core.GetUUID()

	// 2. 创建channel createChannel message
	var boxList []*message.MessageBox2
	s.MessageModel.SendInternalMessage(md.UserId, peer, randomId, false, createChannelMessage, func(i int32, box2 *message.MessageBox2) {
		// TopMessage
		channel.SetTopMessage(box2.MessageId)
		s.DialogModel.InsertOrChannelUpdateDialog(md.UserId, base.PEER_CHANNEL, channel.GetChannelId(), box2.MessageId)
		boxList = append(boxList, box2)
	})

	// sync
	syncUpdates := updates.NewUpdatesLogic(md.UserId)
	pts := int32(core.NextChannelNPtsId(channel.GetChannelId(), 2))
	ptsCount := int32(1)

	updateChannel := &mtproto.TLUpdateChannel{Data2: &mtproto.Update_Data{
		ChannelId: channel.GetChannelId(),
	}}
	syncUpdates.AddUpdate(updateChannel.To_Update())
	updateReadChannelInbox := &mtproto.TLUpdateReadChannelInbox{Data2: &mtproto.Update_Data{
		ChannelId: channel.GetChannelId(),
		MaxId:     boxList[0].MessageId,
	}}
	syncUpdates.AddUpdate(updateReadChannelInbox.To_Update())

	syncUpdates.AddUpdateNewChannelMessage(pts, ptsCount, boxList[0].ToMessage(md.UserId))
	syncUpdates.AddChat(channel.ToChannel(md.UserId))

	sync_client.GetSyncClient().SyncChannelUpdatesNotMe(channel.GetChannelId(), md.UserId, md.AuthId, syncUpdates.ToUpdates())

	// reply
	resultUpdates := syncUpdates
	resultUpdates.AddUpdateMessageId(boxList[0].MessageId, randomId)
	// // Sorry: not impl ChannelsCreateChannel logic
	// glog.Warning("channels.createChannel blocked, License key from https://nebula.chat required to unlock enterprise features.")
	glog.Infof("channels.createChannel - reply: %v", resultUpdates.ToUpdates())
	return resultUpdates.ToUpdates(), nil
}
