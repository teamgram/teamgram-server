// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"net"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

func (d *Dao) getCountryAndRegionByIp(ip string) (string, string) {
	if d.MMDB == nil {
		return "UNKNOWN", ""
	} else {
		r, err := d.MMDB.City(net.ParseIP(ip))
		if err != nil {
			logx.Errorf("getCountryAndRegionByIp - error: %v", err)
			return "UNKNOWN", ""
		}

		return r.City.Names["en"] + ", " + r.Country.Names["en"], r.Country.IsoCode
	}
}

func getAppNameByAppId(appId int32) string {
	return "tdesktop"
}

func (d *Dao) GetAuthorization(ctx context.Context, authKeyId int64) (*mtproto.Authorization, error) {
	authUsersDO, err := d.AuthsDAO.SelectByAuthKeyId(ctx, authKeyId)
	if err != nil || authUsersDO == nil {
		return nil, err
	}

	country, region := d.getCountryAndRegionByIp(authUsersDO.ClientIp)
	// TODO(@benqi): fill plat_form, app_name, (country, region)
	return mtproto.MakeTLAuthorization(&mtproto.Authorization{
		Current:         false,
		OfficialApp:     true,
		Hash:            0,
		PasswordPending: false,
		DeviceModel:     authUsersDO.DeviceModel,
		Platform:        "",
		SystemVersion:   authUsersDO.SystemVersion,
		ApiId:           authUsersDO.ApiId,
		AppName:         authUsersDO.LangPack,
		AppVersion:      authUsersDO.AppVersion,
		DateCreated:     0,
		DateActive:      0,
		Ip:              authUsersDO.ClientIp,
		Country:         country,
		Region:          region,
	}).To_Authorization(), nil

}

func (d *Dao) GetAuthorizations(ctx context.Context, userId int64, excludeAuthKeyId int64) (authorizations []*mtproto.Authorization) {
	//var (
	//	// selfAuthKeyId int64
	//	hash  int64
	//	flags int32
	//)

	authUsersDOList, _ := d.AuthUsersDAO.SelectAuthKeyIds(ctx, userId)
	authorizations = make([]*mtproto.Authorization, 0, len(authUsersDOList))
	idList := make([]int64, 0, len(authUsersDOList))

	for i := 0; i < len(authUsersDOList); i++ {
		//if userId == authUsersDOList[i].UserId {
		//	selfAuthKeyId = authUsersDOList[i].AuthKeyId
		//}
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

	myIdx := -1
	authsDOList, _ := d.AuthsDAO.SelectSessions(ctx, idList)
	for i := 0; i < len(authsDOList); i++ {
		authUsersDO := getAuthUsersDO(authsDOList[i].AuthKeyId)
		if excludeAuthKeyId == authsDOList[i].AuthKeyId {
			//hash = 0
			//flags = 1
			myIdx = i
			continue
		}
		//else {
		//	// TODO(@benqi): hash
		//	hash = authUsersDO.Hash
		//	// authsDOList[i].AuthKeyId
		//	flags = 0
		//}

		// authorization#ad01d61d flags:#
		//	current:flags.0?true
		//	official_app:flags.1?true
		//	password_pending:flags.2?true
		//	hash:long
		//	device_model:string
		//	platform:string
		//	system_version:string
		//	api_id:int
		//	app_name:string
		//	app_version:string
		//	date_created:int
		//	date_active:int
		//	ip:string
		//	country:string
		//	region:string = Authorization;
		country, region := d.getCountryAndRegionByIp(authsDOList[i].ClientIp)
		// TODO(@benqi): fill plat_form, app_name, (country, region)
		authorization := mtproto.MakeTLAuthorization(&mtproto.Authorization{
			Current:         false,
			OfficialApp:     true,
			PasswordPending: false,
			Hash:            authUsersDO.Hash,
			DeviceModel:     authsDOList[i].DeviceModel,
			Platform:        "",
			SystemVersion:   authsDOList[i].SystemVersion,
			ApiId:           authsDOList[i].ApiId,
			AppName:         authsDOList[i].LangPack,
			AppVersion:      authsDOList[i].AppVersion,
			DateCreated:     int32(authUsersDO.DateCreated),
			DateActive:      int32(authsDOList[i].DateActive),
			Ip:              authsDOList[i].ClientIp,
			Country:         country,
			Region:          region,
		}).To_Authorization()

		// log.Debugf("%d - %s", i, authorization.DebugString())
		authorizations = append(authorizations, authorization)
	}

	if myIdx != -1 {
		// log.Debugf("excludeAuthKeyId - %d", excludeAuthKeyId)
		authUsersDO := getAuthUsersDO(excludeAuthKeyId)
		country, region := d.getCountryAndRegionByIp(authsDOList[myIdx].ClientIp)
		authorizations = append([]*mtproto.Authorization{mtproto.MakeTLAuthorization(&mtproto.Authorization{
			Current:         true,
			OfficialApp:     true,
			PasswordPending: false,
			Hash:            0,
			DeviceModel:     authsDOList[myIdx].DeviceModel,
			Platform:        "",
			SystemVersion:   authsDOList[myIdx].SystemVersion,
			ApiId:           authsDOList[myIdx].ApiId,
			AppName:         authsDOList[myIdx].LangPack,
			AppVersion:      authsDOList[myIdx].AppVersion,
			DateCreated:     int32(authUsersDO.DateCreated),
			DateActive:      int32(authsDOList[myIdx].DateActive),
			Ip:              authsDOList[myIdx].ClientIp,
			Country:         country,
			Region:          region,
		}).To_Authorization()}, authorizations...)
	}

	return
}

func (d *Dao) ResetAuthorization(ctx context.Context, userId int64, authKeyId, hash int64) []int64 {
	doList, _ := d.AuthUsersDAO.SelectListByUserId(ctx, userId)
	if len(doList) == 0 {
		return []int64{}
	}

	var (
		keyIdList []int64
		idList    []int64
	)
	if hash == 0 {
		for i := 0; i < len(doList); i++ {
			if doList[i].AuthKeyId != authKeyId {
				idList = append(idList, doList[i].Id)
				keyIdList = append(keyIdList, doList[i].AuthKeyId)
			}
		}
	} else {
		for i := 0; i < len(doList); i++ {
			if doList[i].Hash == hash && doList[i].AuthKeyId != authKeyId {
				idList = append(idList, doList[i].Id)
				keyIdList = append(keyIdList, doList[i].AuthKeyId)
			}
		}
	}

	if len(idList) > 0 {
		d.AuthUsersDAO.DeleteByHashList(ctx, idList)
	} else {
		keyIdList = []int64{}
	}
	return keyIdList
}
