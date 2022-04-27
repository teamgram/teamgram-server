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

type DialogFiltersDO struct {
	Id             int64  `db:"id"`
	UserId         int64  `db:"user_id"`
	DialogFilterId int32  `db:"dialog_filter_id"`
	DialogFilter   string `db:"dialog_filter"`
	OrderValue     int64  `db:"order_value"`
	Deleted        bool   `db:"deleted"`
}
