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

type AuthOpLogsDO struct {
	Id        int64  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	Ip        string `db:"ip"`
	OpType    int32  `db:"op_type"`
	LogText   string `db:"log_text"`
}
