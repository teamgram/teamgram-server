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

type UserPrivaciesDO struct {
	Id      int64  `db:"id"`
	UserId  int64  `db:"user_id"`
	KeyType int32  `db:"key_type"`
	Rules   string `db:"rules"`
}
