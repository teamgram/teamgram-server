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

type UserSettingsDO struct {
	Id      int64  `db:"id" json:"id"`
	UserId  int64  `db:"user_id" json:"user_id"`
	Key2    string `db:"key2" json:"key2"`
	Value   string `db:"value" json:"value"`
	Deleted bool   `db:"deleted" json:"deleted"`
}
