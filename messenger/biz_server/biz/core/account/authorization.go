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
	"github.com/nebula-chat/chatengine/mtproto"
	"time"
)

func (m *AccountModel) GetAuthorizationList(selfAuthKeyId int64, userId int32) []*mtproto.Authorization {
	doList := m.dao.AuthUsersDAO.SelectListByUserId(userId)
	sessList := make([]*mtproto.Authorization, 0, len(doList))
	var (
		hash  int64
		Flags int32
	)
	for _, do := range doList {
		if selfAuthKeyId == do.AuthId {
			hash = 0
			Flags = 1
		} else {
			hash = do.Hash
			Flags = 0
		}
		sess := &mtproto.TLAuthorization{Data2: &mtproto.Authorization_Data{
			Hash:          hash,
			Flags:         Flags,
			DeviceModel:   do.DeviceModel,
			Platform:      do.Platform,
			SystemVersion: do.SystemVersion,
			ApiId:         do.ApiId,
			AppName:       do.AppName,
			AppVersion:    do.AppVersion,
			DateCreated:   do.DateCreated,
			DateActive:    do.DateActive,
			Ip:            do.Ip,
			Country:       do.Country,
			Region:        do.Region,
		}}
		sessList = append(sessList, sess.To_Authorization())
	}

	return sessList
}

func (m *AccountModel) GetAuthKeyIdByHash(userId int32, hash int64) int64 {
	do := m.dao.AuthUsersDAO.SelectByHash(userId, hash)
	if do == nil {
		return 0
	}
	return do.AuthId
}

func (m *AccountModel) DeleteAuthorization(authKeyId int64) {
	m.dao.AuthUsersDAO.Delete(time.Now().Unix(), authKeyId)
}
