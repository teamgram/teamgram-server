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

package auth_session

import (
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/auth_session/biz/dal/dataobject"
)

func getCountryAndRegionByIp(ip string) (string, string) {
	return "Hangzhou, China", ""
}

func getAppNameByAppId(appId int32) string {
	return "tdesktop"
}

func (m *AuthSessionModel) GetAuthorizations(userId int32) (authorizations []*mtproto.Authorization) {
	var (
		selfAuthKeyId int64
		hash  int64
		flags int32
	)

	authUsersDOList := m.dao.AuthUsersDAO.SelectAuthKeyIds(userId)
	authorizations = make([]*mtproto.Authorization, 0, len(authUsersDOList))
	idList := make([]int64, 0, len(authUsersDOList))

	for i := 0; i < len(authUsersDOList); i++ {
		if userId == authUsersDOList[i].UserId {
			selfAuthKeyId = authUsersDOList[i].AuthKeyId
		}
		idList = append(idList, authUsersDOList[i].AuthKeyId)
	}
	if len(idList) == 0 {
		// ??
		return
	}

	getAuthUsersDO := func(authKeyId int64) *dataobject.AuthUsersDO {
		for i := 0; i < len(authUsersDOList); i++ {
			if authKeyId == authUsersDOList[i].AuthKeyId {
				return &authUsersDOList[i]
			}
		}
		return nil
	}

	authsDOList := m.dao.AuthsDAO.SelectSessions(idList)
	for i := 0; i < len(authsDOList); i++ {
		if selfAuthKeyId == authsDOList[i].AuthKeyId {
			hash = 0
			flags = 1
		} else {
			// TODO(@benqi): hash
			hash = authsDOList[i].AuthKeyId
			flags = 0
		}

		authUsersDO := getAuthUsersDO(authsDOList[i].AuthKeyId)
		country, region := getCountryAndRegionByIp(authsDOList[i].ClientIp)
		// TODO(@benqi): fill plat_form, app_name, (country, region)
		authorization := &mtproto.TLAuthorization{Data2: &mtproto.Authorization_Data{
			Hash:          hash,
			Flags:         flags,
			DeviceModel:   authsDOList[i].DeviceModel,
			Platform:      "",
			SystemVersion: authsDOList[i].SystemVersion,
			ApiId:         authsDOList[i].ApiId,
			AppName:       getAppNameByAppId(authsDOList[i].ApiId),
			AppVersion:    authsDOList[i].AppVersion,
			DateCreated:   authUsersDO.DateCreated,
			DateActive:    authUsersDO.DateActived,
			Ip:            authsDOList[i].ClientIp,
			Country:       country,
			Region:        region,
		}}

		authorizations = append(authorizations, authorization.To_Authorization())
	}
	return
}

func (m *AuthSessionModel) ResetAuthorization(hash int64) int64 {
	do := m.dao.AuthUsersDAO.SelectByHash(hash)
	if do == nil {
		return 0
	}

	m.dao.AuthUsersDAO.DeleteByHash(do.Id)
	return do.AuthKeyId
}
