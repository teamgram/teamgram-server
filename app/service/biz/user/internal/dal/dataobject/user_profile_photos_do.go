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

type UserProfilePhotosDO struct {
	Id      int64 `db:"id" json:"id"`
	UserId  int64 `db:"user_id" json:"user_id"`
	PhotoId int64 `db:"photo_id" json:"photo_id"`
	Date2   int64 `db:"date2" json:"date2"`
	Deleted bool  `db:"deleted" json:"deleted"`
}
