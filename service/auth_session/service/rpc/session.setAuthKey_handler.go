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

// session.setAuthKey auth_key:AuthKeyInfo = Bool;
func (s *SessionServiceImpl) SessionSetAuthKey(ctx context.Context, request *mtproto.TLSessionSetAuthKey) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("session.setAuthKey - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	keyInfo := request.GetAuthKey().To_AuthKeyInfo()
	err := s.AuthSessionModel.InsertAuthKey(keyInfo.GetAuthKeyId(), keyInfo.GetAuthKey(), keyInfo.GetFutureSalt().To_FutureSalt())
	if err != nil {
		glog.Error(err)
		return mtproto.ToBool(false), nil
	}

	glog.Info("session.setAuthKey - reply: {true}")
	return mtproto.ToBool(true), nil
}
