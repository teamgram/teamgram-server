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

type DevicesDO struct {
	Id           int64  `db:"id" json:"id"`
	AuthKeyId    int64  `db:"auth_key_id" json:"auth_key_id"`
	UserId       int64  `db:"user_id" json:"user_id"`
	TokenType    int32  `db:"token_type" json:"token_type"`
	Token        string `db:"token" json:"token"`
	NoMuted      bool   `db:"no_muted" json:"no_muted"`
	LockedPeriod int32  `db:"locked_period" json:"locked_period"`
	AppSandbox   bool   `db:"app_sandbox" json:"app_sandbox"`
	Secret       string `db:"secret" json:"secret"`
	OtherUids    string `db:"other_uids" json:"other_uids"`
	State        int32  `db:"state" json:"state"`
}
