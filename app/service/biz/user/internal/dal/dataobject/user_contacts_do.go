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

type UserContactsDO struct {
	Id               int64  `db:"id" json:"id"`
	OwnerUserId      int64  `db:"owner_user_id" json:"owner_user_id"`
	ContactUserId    int64  `db:"contact_user_id" json:"contact_user_id"`
	ContactPhone     string `db:"contact_phone" json:"contact_phone"`
	ContactFirstName string `db:"contact_first_name" json:"contact_first_name"`
	ContactLastName  string `db:"contact_last_name" json:"contact_last_name"`
	Mutual           bool   `db:"mutual" json:"mutual"`
	CloseFriend      bool   `db:"close_friend" json:"close_friend"`
	StoriesHidden    bool   `db:"stories_hidden" json:"stories_hidden"`
	IsDeleted        bool   `db:"is_deleted" json:"is_deleted"`
	Date2            int64  `db:"date2" json:"date2"`
}
