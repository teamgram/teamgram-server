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

// channels.exportInvite#c7560885 channel:InputChannel = ExportedChatInvite;
func (s *ChannelsServiceImpl) ChannelsExportInvite(ctx context.Context, request *mtproto.TLChannelsExportInvite) (*mtproto.ExportedChatInvite, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.exportInvite#c7560885 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	if request.Channel.Constructor == mtproto.TLConstructor_CRC32_inputChannelEmpty {
		// TODO(@benqi): chatUser不能是inputUser和inputUserSelf
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("channels.exportInvite#c7560885 - error: ", err, "; InputPeer invalid")
		return nil, err
	}

	channelLogic, err := s.ChannelModel.NewChannelLogicById(request.GetChannel().GetData2().GetChannelId())
	if err != nil {

	}

	exportedChatInvite := &mtproto.TLChatInviteExported{Data2: &mtproto.ExportedChatInvite_Data{
		Link: channelLogic.ExportedChatInvite(),
	}}

	glog.Infof("channels.exportInvite#c7560885 - reply: {%v}", exportedChatInvite)
	return exportedChatInvite.To_ExportedChatInvite(), nil
}
