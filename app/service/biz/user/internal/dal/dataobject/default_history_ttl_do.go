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

type DefaultHistoryTtlDO struct {
	Id     int64 `db:"id" json:"id"`
	UserId int64 `db:"user_id" json:"user_id"`
	Period int32 `db:"period" json:"period"`
}
