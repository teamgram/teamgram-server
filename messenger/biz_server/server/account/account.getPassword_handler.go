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

// TODO(@benqi): hard code
var (
	newAlgoSalt1 = []byte{0x51, 0x20, 0xBA, 0xD7, 0xD2, 0x1B, 0xE2, 0x5B}
	newAlgoSalt2 = []byte{0x5E, 0x19, 0x37, 0xE6, 0x2B, 0xD2, 0xD1, 0xAC, 0x8E, 0x59, 0xEF, 0x72, 0x51, 0xE3, 0xD9, 0x5E}
	newAlgoG = int32(2)
	newAlgoP = []byte{
		0xC7, 0x60, 0x85, 0x31, 0xC4, 0xE0, 0x98, 0x3C,
		0xA9, 0xDE, 0xC7, 0x30, 0x25, 0x75, 0xB7, 0xF9,
		0xE3, 0x17, 0x9B, 0x66, 0x8A, 0x72, 0x7B, 0x49,
		0x40, 0x82, 0x99, 0x8E, 0xD0, 0x05, 0x0E, 0xF3}
	newSecureAlgoSalt = []byte{0x7D, 0x04, 0xB3, 0x4B, 0x94, 0x82, 0x8C, 0x3D}
)

// account.password#ad2641f8 flags:# has_recovery:flags.0?true has_secure_values:flags.1?true has_password:flags.2?true current_algo:flags.2?PasswordKdfAlgo srp_B:flags.2?bytes srp_id:flags.2?long hint:flags.3?string email_unconfirmed_pattern:flags.4?string new_algo:PasswordKdfAlgo new_secure_algo:SecurePasswordKdfAlgo secure_random:bytes = account.Password;
// account.getPassword#548a30f5 = account.Password;
func (s *AccountServiceImpl) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getPassword#548a30f5 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	newAlgo := &mtproto.TLPasswordKdfAlgoModPow{Data2: &mtproto.PasswordKdfAlgo_Data{
		Salt1: newAlgoSalt1,
		Salt2: newAlgoSalt2,
		G:     newAlgoG,
		P:     newAlgoP,
	}}
	newSecureAlgo := &mtproto.TLSecurePasswordKdfAlgoPBKDF2{Data2: &mtproto.SecurePasswordKdfAlgo_Data{
		Salt: newSecureAlgoSalt,
	}}
	password := &mtproto.TLAccountPassword{Data2: &mtproto.Account_Password_Data{
		NewAlgo:       newAlgo.To_PasswordKdfAlgo(),
		NewSecureAlgo: newSecureAlgo.To_SecurePasswordKdfAlgo(),
	}}

	glog.Infof("account.getPassword#548a30f5 - reply: %s", logger.JsonDebugData(password))
	return password.To_Account_Password(), nil
}
