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
	bizDefaultHistoryTtlModel interface {
		InsertOrUpdate(ctx context.Context, data *DefaultHistoryTtl) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *DefaultHistoryTtl) (lastInsertId, rowsAffected int64, err error)

		Select(ctx context.Context, userId int64) (*DefaultHistoryTtl, error)
	}
)

// InsertOrUpdate
// insert into default_history_ttl(user_id, period) values (:user_id, :period) on duplicate key update period = values(period)
func (m *defaultDefaultHistoryTtlModel) InsertOrUpdate(ctx context.Context, data *DefaultHistoryTtl) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into default_history_ttl(user_id, period) values (:user_id, :period) on duplicate key update period = values(period)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("default_history_ttl.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("default_history_ttl.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("default_history_ttl.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into default_history_ttl(user_id, period) values (:user_id, :period) on duplicate key update period = values(period)
func (m *defaultDefaultHistoryTtlModel) InsertOrUpdateTx(tx *sqlx.Tx, data *DefaultHistoryTtl) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into default_history_ttl(user_id, period) values (:user_id, :period) on duplicate key update period = values(period)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("default_history_ttl.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("default_history_ttl.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("default_history_ttl.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}

// Select
// select id, user_id, period from default_history_ttl where user_id = :user_id
func (m *defaultDefaultHistoryTtlModel) Select(ctx context.Context, userId int64) (rValue *DefaultHistoryTtl, err error) {

	var (
		query = "select id, user_id, period from default_history_ttl where user_id = ?"
		do    = &DefaultHistoryTtl{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("default_history_ttl.Select: %w", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}
