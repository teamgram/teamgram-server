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

package help

import (
	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
	"time"
)

// help.getConfig#c4f9186b = Config;
func (s *HelpServiceImpl) HelpGetConfig(ctx context.Context, request *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("help.getConfig#c4f9186b - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): 设置Reply对象累死人了, 得想个办法实现model和mtproto的自动转换
	helpConfig, _ := proto.Clone(&config).(*mtproto.TLConfig)
	now := int32(time.Now().Unix())
	helpConfig.SetDate(now)
	helpConfig.SetExpires(now + EXPIRES_TIMEOUT)

	reply := helpConfig.To_Config()
	glog.Infof("help.getConfig#c4f9186b - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}
