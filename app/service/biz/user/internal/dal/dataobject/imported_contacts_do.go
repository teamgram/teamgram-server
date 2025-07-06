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

type ImportedContactsDO struct {
	Id             int64 `db:"id" json:"id"`
	UserId         int64 `db:"user_id" json:"user_id"`
	ImportedUserId int64 `db:"imported_user_id" json:"imported_user_id"`
	Deleted        bool  `db:"deleted" json:"deleted"`
}
