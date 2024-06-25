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

type ChatInvitesDO struct {
	Id            int64  `db:"id" json:"id"`
	ChatId        int64  `db:"chat_id" json:"chat_id"`
	AdminId       int64  `db:"admin_id" json:"admin_id"`
	MigratedToId  int64  `db:"migrated_to_id" json:"migrated_to_id"`
	Link          string `db:"link" json:"link"`
	Permanent     bool   `db:"permanent" json:"permanent"`
	Revoked       bool   `db:"revoked" json:"revoked"`
	RequestNeeded bool   `db:"request_needed" json:"request_needed"`
	StartDate     int64  `db:"start_date" json:"start_date"`
	ExpireDate    int64  `db:"expire_date" json:"expire_date"`
	UsageLimit    int32  `db:"usage_limit" json:"usage_limit"`
	Usage2        int32  `db:"usage2" json:"usage2"`
	Requested     int32  `db:"requested" json:"requested"`
	Title         string `db:"title" json:"title"`
	Date2         int64  `db:"date2" json:"date2"`
	State         int32  `db:"state" json:"state"`
}
