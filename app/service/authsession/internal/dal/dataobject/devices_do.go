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

type DevicesDO struct {
	Id           int64  `db:"id"`
	AuthKeyId    int64  `db:"auth_key_id"`
	UserId       int64  `db:"user_id"`
	TokenType    int32  `db:"token_type"`
	Token        string `db:"token"`
	NoMuted      bool   `db:"no_muted"`
	LockedPeriod int32  `db:"locked_period"`
	AppSandbox   bool   `db:"app_sandbox"`
	Secret       string `db:"secret"`
	OtherUids    string `db:"other_uids"`
	State        bool   `db:"state"`
}
