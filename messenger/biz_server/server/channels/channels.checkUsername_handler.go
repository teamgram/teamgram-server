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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/username"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (s *ChannelsServiceImpl) ChannelsCheckUsername(ctx context.Context, request *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.checkUsername#10e6bd2c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// var checked bool
	if request.GetChannel().GetConstructor() == mtproto.TLConstructor_CRC32_inputChannelEmpty {
		glog.Infof("channels.checkUsername#10e6bd2c - reply: {false}")
		return mtproto.ToBool(false), nil
	}

	checked := s.UsernameModel.CheckChannelUsername(request.GetChannel().GetData2().GetChannelId(), request.GetUsername())
	if checked == username.USERNAME_EXISTED_NOTME {
		// err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_NOT_MODIFIED)
		// glog.Error(err)
		// return nil, err
		return mtproto.ToBool(false), nil
	}

	glog.Infof("channels.checkUsername#10e6bd2c - reply: {true}")
	return mtproto.ToBool(true), nil
}
