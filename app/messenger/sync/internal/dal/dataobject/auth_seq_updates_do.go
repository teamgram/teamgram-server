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

type AuthSeqUpdatesDO struct {
	Id         int64  `db:"id"`
	AuthId     int64  `db:"auth_id"`
	UserId     int64  `db:"user_id"`
	Seq        int32  `db:"seq"`
	UpdateType int32  `db:"update_type"`
	UpdateData string `db:"update_data"`
	Date2      int64  `db:"date2"`
}
