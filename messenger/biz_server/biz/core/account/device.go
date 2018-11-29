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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/util"
	"encoding/base64"
)

//const (
//	TOKEN_TYPE_APNS = 1
//	TOKEN_TYPE_GCM = 2
//	TOKEN_TYPE_MPNS = 3
//	TOKEN_TYPE_SIMPLE_PUSH = 4
//	TOKEN_TYPE_UBUNTU_PHONE = 5
//	TOKEN_TYPE_BLACKBERRY = 6
//	// Android里使用
//	TOKEN_TYPE_INTERNAL_PUSH = 7
//)
//

func (m *AccountModel) RegisterDevice(authKeyId int64, userId int32, device *mtproto.TLAccountRegisterDevice) bool {
	do := &dataobject.DevicesDO{
		AuthKeyId:  authKeyId,
		UserId:     userId,
		TokenType:  int8(device.TokenType),
		Token:      device.Token,
		AppSandbox: util.BoolToInt8(mtproto.FromBool(device.AppSandbox)),
		Secret:     base64.RawStdEncoding.EncodeToString(device.Secret),
		OtherUids:  util.JoinInt32List(device.OtherUids, ":"),
	}

	m.dao.DevicesDAO.InsertOrUpdate(do)
	return true
}

func (m *AccountModel) UnRegisterDevice(authKeyId int64, userId int32) bool {
	m.dao.DevicesDAO.UpdateState(1, authKeyId, userId)
	return true
}
