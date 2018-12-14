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

// auth.importAuthorization#e3ef9613 id:int bytes:bytes = auth.Authorization;
func (s *AuthServiceImpl) AuthImportAuthorization(ctx context.Context, request *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.importAuthorization#e3ef9613 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AuthExportAuthorization logic
	authorization := &mtproto.TLAuthAuthorization{Data2: &mtproto.Auth_Authorization_Data{
		TmpSessions: request.GetId(),
		User:        mtproto.NewTLUserEmpty().To_User(),
	}}

	glog.Infof("auth.importAuthorization#e3ef9613- reply: %s", logger.JsonDebugData(authorization))
	return authorization.To_Auth_Authorization(), nil
}
