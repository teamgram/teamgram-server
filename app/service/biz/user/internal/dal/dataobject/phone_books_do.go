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

type PhoneBooksDO struct {
	Id        int64  `db:"id"`
	UserId    int64  `db:"user_id"`
	AuthKeyId int64  `db:"auth_key_id"`
	ClientId  int64  `db:"client_id"`
	Phone     string `db:"phone"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}
