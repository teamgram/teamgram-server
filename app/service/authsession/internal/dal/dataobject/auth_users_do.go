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

type AuthUsersDO struct {
	Id            int64  `db:"id"`
	AuthKeyId     int64  `db:"auth_key_id"`
	UserId        int64  `db:"user_id"`
	Hash          int64  `db:"hash"`
	Layer         int32  `db:"layer"`
	DeviceModel   string `db:"device_model"`
	Platform      string `db:"platform"`
	SystemVersion string `db:"system_version"`
	ApiId         int32  `db:"api_id"`
	AppName       string `db:"app_name"`
	AppVersion    string `db:"app_version"`
	DateCreated   int64  `db:"date_created"`
	DateActived   int64  `db:"date_actived"`
	Ip            string `db:"ip"`
	Country       string `db:"country"`
	Region        string `db:"region"`
	Deleted       bool   `db:"deleted"`
}
