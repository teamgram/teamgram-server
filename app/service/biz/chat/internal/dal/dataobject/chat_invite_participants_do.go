/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type ChatInviteParticipantsDO struct {
	Id      int64  `db:"id"`
	Link    string `db:"link"`
	UserId  int64  `db:"user_id"`
	Date2   int64  `db:"date2"`
	Deleted bool   `db:"deleted"`
}
