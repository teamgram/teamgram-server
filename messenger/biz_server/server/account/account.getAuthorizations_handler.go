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

package account

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

/*
	selfUser: hash = 0, flag = 1
	other:  hash and flag load from db
*/

// account.getAuthorizations#e320c158 = account.Authorizations;
func (s *AccountServiceImpl) AccountGetAuthorizations(ctx context.Context, request *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getAuthorizations#e320c158 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	sessionList := s.AccountModel.GetAuthorizationList(md.AuthId, md.UserId)
	authorizations := &mtproto.TLAccountAuthorizations{Data2: &mtproto.Account_Authorizations_Data{
		Authorizations: sessionList,
	}}

	glog.Infof("account.getAuthorizations#e320c158 - reply: {%s}", logger.JsonDebugData(authorizations))
	return authorizations.To_Account_Authorizations(), nil
}
