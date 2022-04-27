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

type UserProfilePhotosDO struct {
	Id      int64 `db:"id"`
	UserId  int64 `db:"user_id"`
	PhotoId int64 `db:"photo_id"`
	Date2   int64 `db:"date2"`
	Deleted bool  `db:"deleted"`
}
