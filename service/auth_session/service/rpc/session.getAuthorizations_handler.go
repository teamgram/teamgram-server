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

package rpc

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// session.getAuthorizations user_id:int = account.Authorizations;
func (s *SessionServiceImpl) SessionGetAuthorizations(ctx context.Context, request *mtproto.TLSessionGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("session.getAuthorizations - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	authorizationList := s.AuthSessionModel.GetAuthorizations(request.GetUserId())
    reply := &mtproto.TLAccountAuthorizations{Data2: &mtproto.Account_Authorizations_Data{
        Authorizations: authorizationList,
    }}

	glog.Infof("session.getAuthorizations - reply: {%s}", logger.JsonDebugData(reply))
	return reply.To_Account_Authorizations(), nil
}
