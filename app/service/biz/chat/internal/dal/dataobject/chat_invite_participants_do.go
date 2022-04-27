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

type ChatInviteParticipantsDO struct {
	Id         int64  `db:"id"`
	ChatId     int64  `db:"chat_id"`
	Link       string `db:"link"`
	UserId     int64  `db:"user_id"`
	Requested  bool   `db:"requested"`
	ApprovedBy int64  `db:"approved_by"`
	Date2      int64  `db:"date2"`
	Deleted    bool   `db:"deleted"`
}
