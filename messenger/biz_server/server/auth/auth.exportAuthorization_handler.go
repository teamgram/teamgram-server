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


/*
	## Authorization Transfer
	The following methods can be used to eliminate the need for users to enter the code from a text message every time:

	```
	auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
	auth.authorization#f6b673a4 expires:int user:User = auth.Authorization;
	---functions---
	auth.importAuthorization#e3ef9613 id:int bytes:bytes = auth.Authorization;
	auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
	```

	auth.exportAuthorization must be executed in the current DC (the DC with which a connection has already been established),
	passing in dc_id as the value for the new DC.
	The method should return the user identifier and a long string of random data.
	An import operation can be performed at the new DC by sending it what was received.
	Queries requiring authorization can then be successfully executed in the new DC.
 */

// auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
func (s *AuthServiceImpl) AuthExportAuthorization(ctx context.Context, request *mtproto.TLAuthExportAuthorization) (*mtproto.Auth_ExportedAuthorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.exportAuthorization#e5bfffcd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	exported := &mtproto.TLAuthExportedAuthorization{Data2: &mtproto.Auth_ExportedAuthorization_Data{
		Id: request.GetDcId(),
		Bytes: []byte{1,2,3,4},
	}}

	glog.Infof("auth.exportAuthorization#e5bfffcd - reply: %s", logger.JsonDebugData(exported))
	return exported.To_Auth_ExportedAuthorization(), nil
}
