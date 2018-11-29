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

package auth

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// auth.recoverPassword#4ea56e92 code:string = auth.Authorization;
func (s *AuthServiceImpl) AuthRecoverPassword(ctx context.Context, request *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.recoverPassword#4ea56e92 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		err error = nil
	)

	if request.Code == "" {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CODE_INVALID)
		glog.Error(err)
		return nil, err
	} else {
		err = s.AccountModel.CheckRecoverCode(md.UserId, request.Code)
		if err != nil {
			glog.Error(err)
			return nil, err
		}
	}

	user := s.UserModel.GetUserById(md.UserId, md.UserId)
	authAuthorization := &mtproto.TLAuthAuthorization{Data2: &mtproto.Auth_Authorization_Data{
		User: user,
	}}

	glog.Infof("auth.recoverPassword#4ea56e92 - reply: %s\n", logger.JsonDebugData(authAuthorization))
	return authAuthorization.To_Auth_Authorization(), nil
}
