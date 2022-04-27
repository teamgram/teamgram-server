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

type UnregisteredContactsDO struct {
	Id              int64  `db:"id"`
	Phone           string `db:"phone"`
	ImporterUserId  int64  `db:"importer_user_id"`
	ImportFirstName string `db:"import_first_name"`
	ImportLastName  string `db:"import_last_name"`
	Imported        bool   `db:"imported"`
}
