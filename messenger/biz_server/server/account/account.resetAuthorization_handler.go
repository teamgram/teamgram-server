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
	"github.com/nebula-chat/chatengine/service/auth_session/client"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)

// account.resetAuthorization#df77f3bc hash:long = Bool;
func (s *AccountServiceImpl) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.resetAuthorization#df77f3bc - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	keyId := auth_session_client.ResetAuthorization(md.UserId, request.GetHash())

	// notify kill session
	upds := &mtproto.TLUpdateAccountResetAuthorization{Data2: &mtproto.Updates_Data{
		UserId:    md.UserId,
		AuthKeyId: keyId,
	}}
	sync_client.GetSyncClient().SyncUpdatesMe(md.UserId, keyId, 0, upds.To_Updates())

	glog.Infof("account.resetAuthorization#df77f3bc - reply: {%d}", keyId)
	return mtproto.ToBool(keyId != 0), nil
}
