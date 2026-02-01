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

type PopularContactsDO struct {
	Id        int64  `db:"id" json:"id"`
	Phone     string `db:"phone" json:"phone"`
	Importers int32  `db:"importers" json:"importers"`
	Deleted   bool   `db:"deleted" json:"deleted"`
	UpdateAt  string `db:"update_at" json:"update_at"`
}
