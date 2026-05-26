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
var _ *sqlx.Tx

type bizDialogFilterTagsModel interface {
	InsertOrUpdate(ctx context.Context, data *DialogFilterTags) (lastInsertId, rowsAffected int64, err error)
	Select(ctx context.Context, userId int64) (*DialogFilterTags, error)
	Update(ctx context.Context, enabled bool, userId int64) (rowsAffected int64, err error)
}

type DialogFilterTagsTxModel interface {
	InsertOrUpdate(data *DialogFilterTags) (lastInsertId, rowsAffected int64, err error)
	Select(userId int64) (*DialogFilterTags, error)
	Update(enabled bool, userId int64) (rowsAffected int64, err error)
}

type defaultDialogFilterTagsTxModel struct {
	tx *sqlx.Tx
}

func NewDialogFilterTagsTxModel(tx *sqlx.Tx) DialogFilterTagsTxModel {
	return &defaultDialogFilterTagsTxModel{tx: tx}
}

// InsertOrUpdate
// insert into dialog_filter_tags(user_id, enabled) values (:user_id, :enabled) on duplicate key update enabled = values(enabled)
func (m *defaultDialogFilterTagsModel) InsertOrUpdate(ctx context.Context, data *DialogFilterTags) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filter_tags(user_id, enabled) values (:user_id, :enabled) on duplicate key update enabled = values(enabled)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into dialog_filter_tags(user_id, enabled) values (:user_id, :enabled) on duplicate key update enabled = values(enabled)
func (m *defaultDialogFilterTagsTxModel) InsertOrUpdate(data *DialogFilterTags) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filter_tags(user_id, enabled) values (:user_id, :enabled) on duplicate key update enabled = values(enabled)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.InsertOrUpdate rows affected: %w", err)
	}

	return
}

// Select
// select id, user_id, enabled from dialog_filter_tags where user_id = :user_id
func (m *defaultDialogFilterTagsModel) Select(ctx context.Context, userId int64) (rValue *DialogFilterTags, err error) {

	var (
		query = "select id, user_id, enabled from dialog_filter_tags where user_id = ?"
		do    = &DialogFilterTags{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_filter_tags",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_filter_tags.Select: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// Select
// select id, user_id, enabled from dialog_filter_tags where user_id = :user_id
func (m *defaultDialogFilterTagsTxModel) Select(userId int64) (rValue *DialogFilterTags, err error) {
	var (
		query = "select id, user_id, enabled from dialog_filter_tags where user_id = ?"
		do    = &DialogFilterTags{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_filter_tags",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_filter_tags.Select: %w", err)
		return
	}
	rValue = do

	return
}

// Update
// update dialog_filter_tags set enabled = :enabled where user_id = :user_id
func (m *defaultDialogFilterTagsModel) Update(ctx context.Context, enabled bool, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_filter_tags set enabled = ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, enabled, userId)

	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.Update rows affected: %w", err)
		return
	}

	return
}

// Update
// update dialog_filter_tags set enabled = :enabled where user_id = :user_id
func (m *defaultDialogFilterTagsTxModel) Update(enabled bool, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filter_tags set enabled = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, enabled, userId)

	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filter_tags.Update rows affected: %w", err)
		return
	}

	return
}
