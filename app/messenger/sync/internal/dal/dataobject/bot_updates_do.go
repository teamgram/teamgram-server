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

type BotUpdatesDO struct {
	Id         int64  `db:"id"`
	BotId      int64  `db:"bot_id"`
	UpdateId   int32  `db:"update_id"`
	UpdateType int32  `db:"update_type"`
	UpdateData string `db:"update_data"`
	Date2      int64  `db:"date2"`
}
