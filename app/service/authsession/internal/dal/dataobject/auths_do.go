/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type AuthsDO struct {
	Id             int64  `db:"id"`
	AuthKeyId      int64  `db:"auth_key_id"`
	Layer          int32  `db:"layer"`
	ApiId          int32  `db:"api_id"`
	DeviceModel    string `db:"device_model"`
	SystemVersion  string `db:"system_version"`
	AppVersion     string `db:"app_version"`
	SystemLangCode string `db:"system_lang_code"`
	LangPack       string `db:"lang_pack"`
	LangCode       string `db:"lang_code"`
	SystemCode     string `db:"system_code"`
	Proxy          string `db:"proxy"`
	Params         string `db:"params"`
	ClientIp       string `db:"client_ip"`
	DateActive     int64  `db:"date_active"`
	Deleted        bool   `db:"deleted"`
}
