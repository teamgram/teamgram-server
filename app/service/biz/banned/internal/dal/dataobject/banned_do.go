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

type BannedDO struct {
	Id           int64  `db:"id"`
	Phone        string `db:"phone"`
	BannedTime   int64  `db:"banned_time"`
	Expires      int64  `db:"expires"`
	BannedReason string `db:"banned_reason"`
	Log          string `db:"log"`
	State        int32  `db:"state"`
}
