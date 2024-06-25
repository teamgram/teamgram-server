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

type ChatInviteParticipantsDO struct {
	Id         int64  `db:"id" json:"id"`
	ChatId     int64  `db:"chat_id" json:"chat_id"`
	Link       string `db:"link" json:"link"`
	UserId     int64  `db:"user_id" json:"user_id"`
	Requested  bool   `db:"requested" json:"requested"`
	ApprovedBy int64  `db:"approved_by" json:"approved_by"`
	Date2      int64  `db:"date2" json:"date2"`
	Deleted    bool   `db:"deleted" json:"deleted"`
}
