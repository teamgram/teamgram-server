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

type ImportedContactsDO struct {
	Id             int64 `db:"id"`
	UserId         int64 `db:"user_id"`
	ImportedUserId int64 `db:"imported_user_id"`
	Deleted        bool  `db:"deleted"`
}
