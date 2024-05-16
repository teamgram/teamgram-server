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

type AuthUsersDO struct {
	Id            int64  `db:"id" json:"id"`
	AuthKeyId     int64  `db:"auth_key_id" json:"auth_key_id"`
	UserId        int64  `db:"user_id" json:"user_id"`
	Hash          int64  `db:"hash" json:"hash"`
	Layer         int32  `db:"layer" json:"layer"`
	DeviceModel   string `db:"device_model" json:"device_model"`
	Platform      string `db:"platform" json:"platform"`
	SystemVersion string `db:"system_version" json:"system_version"`
	ApiId         int32  `db:"api_id" json:"api_id"`
	AppName       string `db:"app_name" json:"app_name"`
	AppVersion    string `db:"app_version" json:"app_version"`
	DateCreated   int64  `db:"date_created" json:"date_created"`
	DateActived   int64  `db:"date_actived" json:"date_actived"`
	Ip            string `db:"ip" json:"ip"`
	Country       string `db:"country" json:"country"`
	Region        string `db:"region" json:"region"`
	State         int32  `db:"state" json:"state"`
	Deleted       bool   `db:"deleted" json:"deleted"`
}
