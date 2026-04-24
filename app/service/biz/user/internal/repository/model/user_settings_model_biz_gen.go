/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB

type (
	bizUserSettingsModel interface {
		InsertOrUpdate(ctx context.Context, data *UserSettings) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UserSettings) (lastInsertId, rowsAffected int64, err error)

		SelectByKey(ctx context.Context, userId int64, key2 string) (*UserSettings, error)

		Update(ctx context.Context, value string, userId int64, key2 string) (rowsAffected int64, err error)
		UpdateTx(tx *sqlx.Tx, value string, userId int64, key2 string) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)
func (m *defaultUserSettingsModel) InsertOrUpdate(ctx context.Context, data *UserSettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_settings.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_settings.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_settings.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)
func (m *defaultUserSettingsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UserSettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_settings.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_settings.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_settings.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}

// SelectByKey
// select id, user_id, key2, value from user_settings where user_id = :user_id and key2 = :key2 and deleted = 0 limit 1
func (m *defaultUserSettingsModel) SelectByKey(ctx context.Context, userId int64, key2 string) (rValue *UserSettings, err error) {

	var (
		query = "select id, user_id, key2, value from user_settings where user_id = ? and key2 = ? and deleted = 0 limit 1"
		do    = &UserSettings{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, key2)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("user_settings.SelectByKey: %w", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// Update
// update user_settings set value = :value, deleted = 0 where user_id = :user_id and key2 = :key2
func (m *defaultUserSettingsModel) Update(ctx context.Context, value string, userId int64, key2 string) (rowsAffected int64, err error) {

	var (
		query   = "update user_settings set value = ?, deleted = 0 where user_id = ? and key2 = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, value, userId, key2)

	if err != nil {
		err = fmt.Errorf("user_settings.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_settings.Update rows affected: %w", err)
	}

	return
}

// UpdateTx
// update user_settings set value = :value, deleted = 0 where user_id = :user_id and key2 = :key2
func (m *defaultUserSettingsModel) UpdateTx(tx *sqlx.Tx, value string, userId int64, key2 string) (rowsAffected int64, err error) {
	var (
		query   = "update user_settings set value = ?, deleted = 0 where user_id = ? and key2 = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, value, userId, key2)

	if err != nil {
		err = fmt.Errorf("user_settings.UpdateTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_settings.UpdateTx rows affected: %w", err)
	}

	return
}
