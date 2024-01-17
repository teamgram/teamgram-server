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

type DialogFiltersDO struct {
	Id             int64  `db:"id" json:"id"`
	UserId         int64  `db:"user_id" json:"user_id"`
	DialogFilterId int32  `db:"dialog_filter_id" json:"dialog_filter_id"`
	IsChatlist     bool   `db:"is_chatlist" json:"is_chatlist"`
	JoinedBySlug   bool   `db:"joined_by_slug" json:"joined_by_slug"`
	Slug           string `db:"slug" json:"slug"`
	HasMyInvites   int32  `db:"has_my_invites" json:"has_my_invites"`
	DialogFilter   string `db:"dialog_filter" json:"dialog_filter"`
	OrderValue     int64  `db:"order_value" json:"order_value"`
	FromSuggested  int32  `db:"from_suggested" json:"from_suggested"`
	Deleted        bool   `db:"deleted" json:"deleted"`
}
