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

type ChatInvitesDO struct {
	Id            int64  `db:"id"`
	ChatId        int64  `db:"chat_id"`
	AdminId       int64  `db:"admin_id"`
	MigratedToId  int64  `db:"migrated_to_id"`
	Link          string `db:"link"`
	Permanent     bool   `db:"permanent"`
	Revoked       bool   `db:"revoked"`
	RequestNeeded bool   `db:"request_needed"`
	StartDate     int64  `db:"start_date"`
	ExpireDate    int64  `db:"expire_date"`
	UsageLimit    int32  `db:"usage_limit"`
	Usage2        int32  `db:"usage2"`
	Requested     int32  `db:"requested"`
	Title         string `db:"title"`
	Date2         int64  `db:"date2"`
	State         int32  `db:"state"`
}
