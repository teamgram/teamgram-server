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

// account.resetAuthorization#df77f3bc hash:long = Bool;
func (s *AccountServiceImpl) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.resetAuthorization#df77f3bc - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	authKeyId := s.AccountModel.GetAuthKeyIdByHash(md.UserId, request.GetHash())
	if authKeyId == 0 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("account.resetAuthorization#df77f3bc - not found hash ", err)
		return nil, err
	}

	// TODO(@benqi): found session, kick off.
	s.AccountModel.DeleteAuthorization(authKeyId)

	glog.Infof("account.checkUsername#2714d86c - reply: {true}")
	return mtproto.ToBool(true), nil
}
