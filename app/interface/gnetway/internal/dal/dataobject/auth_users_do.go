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

type AuthUsersDO struct {
	Id                   int64 `db:"id" json:"id"`
	AuthKeyId            int64 `db:"auth_key_id" json:"auth_key_id"`
	UserId               int64 `db:"user_id" json:"user_id"`
	Hash                 int64 `db:"hash" json:"hash"`
	DateCreated          int64 `db:"date_created" json:"date_created"`
	DateActive           int64 `db:"date_active" json:"date_active"`
	State                int32 `db:"state" json:"state"`
	AndroidPushSessionId int64 `db:"android_push_session_id" json:"android_push_session_id"`
	AuthorizationTtlDays int32 `db:"authorization_ttl_days" json:"authorization_ttl_days"`
	Deleted              bool  `db:"deleted" json:"deleted"`
}
