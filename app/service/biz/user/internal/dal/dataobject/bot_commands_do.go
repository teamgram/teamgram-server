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

type BotCommandsDO struct {
	Id          int64  `db:"id"`
	BotId       int64  `db:"bot_id"`
	Command     string `db:"command"`
	Description string `db:"description"`
}
