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

type bizDialogPreferenceVersionsModel interface {
	Increment(ctx context.Context, data *DialogPreferenceVersions) (lastInsertId, rowsAffected int64, err error)
	SelectOne(ctx context.Context, userId int64, scopeKind string, folderId int32) (*DialogPreferenceVersions, error)
}

type DialogPreferenceVersionsTxModel interface {
	Increment(data *DialogPreferenceVersions) (lastInsertId, rowsAffected int64, err error)
	SelectOne(userId int64, scopeKind string, folderId int32) (*DialogPreferenceVersions, error)
}

type defaultDialogPreferenceVersionsTxModel struct {
	tx *sqlx.Tx
}

func NewDialogPreferenceVersionsTxModel(tx *sqlx.Tx) DialogPreferenceVersionsTxModel {
	return &defaultDialogPreferenceVersionsTxModel{tx: tx}
}

// Increment
// insert into dialog_preference_versions(user_id, scope_kind, folder_id, aggregate_version) values (:user_id, :scope_kind, :folder_id, 1) on duplicate key update aggregate_version = aggregate_version + 1
func (m *defaultDialogPreferenceVersionsModel) Increment(ctx context.Context, data *DialogPreferenceVersions) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preference_versions(user_id, scope_kind, folder_id, aggregate_version) values (:user_id, :scope_kind, :folder_id, 1) on duplicate key update aggregate_version = aggregate_version + 1"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preference_versions.Increment named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preference_versions.Increment last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preference_versions.Increment rows affected: %w", err)
	}

	return

}

// Increment
// insert into dialog_preference_versions(user_id, scope_kind, folder_id, aggregate_version) values (:user_id, :scope_kind, :folder_id, 1) on duplicate key update aggregate_version = aggregate_version + 1
func (m *defaultDialogPreferenceVersionsTxModel) Increment(data *DialogPreferenceVersions) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preference_versions(user_id, scope_kind, folder_id, aggregate_version) values (:user_id, :scope_kind, :folder_id, 1) on duplicate key update aggregate_version = aggregate_version + 1"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preference_versions.Increment named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preference_versions.Increment last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preference_versions.Increment rows affected: %w", err)
	}

	return
}

// SelectOne
// select user_id, scope_kind, folder_id, aggregate_version from dialog_preference_versions where user_id = :user_id and scope_kind = :scope_kind and folder_id = :folder_id limit 1
func (m *defaultDialogPreferenceVersionsModel) SelectOne(ctx context.Context, userId int64, scopeKind string, folderId int32) (rValue *DialogPreferenceVersions, err error) {

	var (
		query = "select user_id, scope_kind, folder_id, aggregate_version from dialog_preference_versions where user_id = ? and scope_kind = ? and folder_id = ? limit 1"
		do    = &DialogPreferenceVersions{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, scopeKind, folderId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_preference_versions",
				Key:      fmt.Sprintf("user_id=%v,scope_kind=%v,folder_id=%v", userId, scopeKind, folderId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_preference_versions.SelectOne: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectOne
// select user_id, scope_kind, folder_id, aggregate_version from dialog_preference_versions where user_id = :user_id and scope_kind = :scope_kind and folder_id = :folder_id limit 1
func (m *defaultDialogPreferenceVersionsTxModel) SelectOne(userId int64, scopeKind string, folderId int32) (rValue *DialogPreferenceVersions, err error) {
	var (
		query = "select user_id, scope_kind, folder_id, aggregate_version from dialog_preference_versions where user_id = ? and scope_kind = ? and folder_id = ? limit 1"
		do    = &DialogPreferenceVersions{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, scopeKind, folderId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_preference_versions",
				Key:      fmt.Sprintf("user_id=%v,scope_kind=%v,folder_id=%v", userId, scopeKind, folderId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_preference_versions.SelectOne: %w", err)
		return
	}
	rValue = do

	return
}
