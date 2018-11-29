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

// help.getInviteText#4d392343 = help.InviteText;
func (s *HelpServiceImpl) HelpGetInviteText(ctx context.Context, request *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("help.getInviteText#4d392343 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	inviteText := &mtproto.TLHelpInviteText{Data2: &mtproto.Help_InviteText_Data{
		Message: "Invited by @benqi",
	}}

	glog.Infof("help.getInviteText#4d392343 - reply: %s", logger.JsonDebugData(inviteText))
	return inviteText.To_Help_InviteText(), nil
}
