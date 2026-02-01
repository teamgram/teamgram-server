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

type PredefinedUsersDO struct {
	Id               int64  `db:"id" json:"id"`
	Phone            string `db:"phone" json:"phone"`
	FirstName        string `db:"first_name" json:"first_name"`
	LastName         string `db:"last_name" json:"last_name"`
	Username         string `db:"username" json:"username"`
	Code             string `db:"code" json:"code"`
	Verified         bool   `db:"verified" json:"verified"`
	RegisteredUserId int64  `db:"registered_user_id" json:"registered_user_id"`
	Deleted          bool   `db:"deleted" json:"deleted"`
}
