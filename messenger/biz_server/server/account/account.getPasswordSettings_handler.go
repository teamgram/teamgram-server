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

// account.getPasswordSettings#bc8d11bb current_password_hash:bytes = account.PasswordSettings;
func (s *AccountServiceImpl) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getPasswordSettings#bc8d11bb - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//passwordLogic, err := s.AccountModel.MakePasswordData(md.UserId)
	//if err != nil {
	//	glog.Error("account.getPassword#548a30f5 - error: ", err)
	//	return nil, err
	//}
	//
	//settings, err := passwordLogic.GetPasswordSetting(request.GetCurrentPasswordHash())
	//if err != nil {
	//	glog.Error("account.getPassword#548a30f5 - error: ", err)
	//	return nil, err
	//}
	//
	//glog.Infof("account.getPasswordSettings#bc8d11bb - reply: %s", logger.JsonDebugData(settings))
	return nil, nil
}
