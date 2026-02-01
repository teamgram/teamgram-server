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

type AuthKeysDO struct {
	Id        int64  `db:"id" json:"id"`
	AuthKeyId int64  `db:"auth_key_id" json:"auth_key_id"`
	Body      string `db:"body" json:"body"`
	Deleted   bool   `db:"deleted" json:"deleted"`
}
