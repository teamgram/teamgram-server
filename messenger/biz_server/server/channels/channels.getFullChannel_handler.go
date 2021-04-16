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
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// channels.getFullChannel#8736a09 channel:InputChannel = messages.ChatFull;
func (s *ChannelsServiceImpl) ChannelsGetFullChannel(ctx context.Context, request *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.getFullChannel#8736a09 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	if request.Channel.Constructor == mtproto.TLConstructor_CRC32_inputChannelEmpty {
		// TODO(@benqi): chatUser不能是inputUser和inputUserSelf
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("channels.exportInvite#c7560885 - error: ", err, "; InputPeer invalid")
		return nil, err
	}

	inputChannel := request.GetChannel().To_InputChannel()

	channelLogic, err := s.ChannelModel.NewChannelLogicById(inputChannel.GetChannelId())
	if err != nil {
		glog.Error("channels.getFullChannel#8736a09 - error: ", err)
		return nil, err
	}

	// idList := channelLogic.GetChannelParticipantIdList()
	messagesChatFull := &mtproto.TLMessagesChatFull{Data2: &mtproto.Messages_ChatFull_Data{
		FullChat: channelLogic.ToChannelFull(md.UserId),
		Chats:    []*mtproto.Chat{channelLogic.ToChannel(md.UserId)},
		Users:    []*mtproto.User{},
	}}

	glog.Infof("channels.getFullChannel#8736a09 - reply: %s", logger.JsonDebugData(messagesChatFull))
	return messagesChatFull.To_Messages_ChatFull(), nil
}
