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
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

const (
	// TODO(@benqi): add support user.
	kSupportUserID = int32(2)
)

// help.getSupport#9cdf08cd = help.Support;
func (s *HelpServiceImpl) HelpGetSupport(ctx context.Context, request *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("help.getSupport#9cdf08cd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpGetSupport logic
	reply := &mtproto.TLHelpSupport{Data2: &mtproto.Help_Support_Data{
		PhoneNumber: "+86 111 1111 1111",
		User:        &mtproto.User{Constructor: mtproto.TLConstructor_CRC32_userEmpty, Data2: &mtproto.User_Data{Id: kSupportUserID}},
	}}

	glog.Infof("help.getSupport#9cdf08cd - reply: {%v}\n", reply)
	return reply.To_Help_Support(), nil
}
