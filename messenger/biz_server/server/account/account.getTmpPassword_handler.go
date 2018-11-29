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
	"time"
)

// account.getTmpPassword#4a82327e password_hash:bytes period:int = account.TmpPassword;
func (s *AccountServiceImpl) AccountGetTmpPassword(ctx context.Context, request *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountGetTmpPassword - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Check password_hash invalid, android source code
	// byte[] hash = new byte[currentPassword.current_salt.length * 2 + passwordBytes.length];
	// System.arraycopy(currentPassword.current_salt, 0, hash, 0, currentPassword.current_salt.length);
	// System.arraycopy(passwordBytes, 0, hash, currentPassword.current_salt.length, passwordBytes.length);
	// System.arraycopy(currentPassword.current_salt, 0, hash, hash.length - currentPassword.current_salt.length, currentPassword.current_salt.length);

	// account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
	tmpPassword := mtproto.NewTLAccountTmpPassword()
	tmpPassword.SetTmpPassword([]byte("01234567899876543210"))
	tmpPassword.SetValidUntil(int32(time.Now().Unix()) + request.Period)

	glog.Infof("AccountServiceImpl - reply: %s", logger.JsonDebugData(tmpPassword))
	return tmpPassword.To_Account_TmpPassword(), nil
}
