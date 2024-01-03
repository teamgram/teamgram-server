/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2023-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type CloseFriendsDO struct {
	Id            int64 `db:"id" json:"id"`
	UserId        int64 `db:"user_id" json:"user_id"`
	CloseFriendId int64 `db:"close_friend_id" json:"close_friend_id"`
	Date          int64 `db:"date" json:"date"`
	Deleted       bool  `db:"deleted" json:"deleted"`
}
