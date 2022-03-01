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

type PopularContactsDO struct {
	Id        int64  `db:"id"`
	Phone     string `db:"phone"`
	Importers int32  `db:"importers"`
	Deleted   bool   `db:"deleted"`
	UpdateAt  string `db:"update_at"`
}
