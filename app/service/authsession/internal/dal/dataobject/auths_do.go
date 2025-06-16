/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type AuthsDO struct {
	Id             int64  `db:"id" json:"id"`
	AuthKeyId      int64  `db:"auth_key_id" json:"auth_key_id"`
	Layer          int32  `db:"layer" json:"layer"`
	ApiId          int32  `db:"api_id" json:"api_id"`
	DeviceModel    string `db:"device_model" json:"device_model"`
	SystemVersion  string `db:"system_version" json:"system_version"`
	AppVersion     string `db:"app_version" json:"app_version"`
	SystemLangCode string `db:"system_lang_code" json:"system_lang_code"`
	LangPack       string `db:"lang_pack" json:"lang_pack"`
	LangCode       string `db:"lang_code" json:"lang_code"`
	SystemCode     string `db:"system_code" json:"system_code"`
	Proxy          string `db:"proxy" json:"proxy"`
	Params         string `db:"params" json:"params"`
	ClientIp       string `db:"client_ip" json:"client_ip"`
	DateActive     int64  `db:"date_active" json:"date_active"`
	Deleted        bool   `db:"deleted" json:"deleted"`
}
