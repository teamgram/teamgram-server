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

type UserPtsUpdatesDO struct {
	Id         int64  `db:"id" json:"id"`
	UserId     int64  `db:"user_id" json:"user_id"`
	Pts        int32  `db:"pts" json:"pts"`
	PtsCount   int32  `db:"pts_count" json:"pts_count"`
	UpdateType int32  `db:"update_type" json:"update_type"`
	UpdateData string `db:"update_data" json:"update_data"`
	Date2      int64  `db:"date2" json:"date2"`
}
