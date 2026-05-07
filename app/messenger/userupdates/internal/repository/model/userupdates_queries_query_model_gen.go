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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sqlx.DB
var _ *sqlx.Tx

type VisibleDialogCountRow struct {
	Count int32 `db:"count"`
}

type UnreadDialogCountRow struct {
	Count int32 `db:"count"`
}

type UserupdatesQueriesModel interface {
	CountVisibleDialogs(ctx context.Context, userId int64) (*VisibleDialogCountRow, error)
	SumUnreadDialogs(ctx context.Context, userId int64) (*UnreadDialogCountRow, error)
}

type UserupdatesQueriesTxModel interface {
	CountVisibleDialogs(userId int64) (*VisibleDialogCountRow, error)
	SumUnreadDialogs(userId int64) (*UnreadDialogCountRow, error)
}

type defaultUserupdatesQueriesModel struct {
	db *sqlx.DB
}

func NewUserupdatesQueriesModel(db *sqlx.DB) UserupdatesQueriesModel {
	return &defaultUserupdatesQueriesModel{db: db}
}

type defaultUserupdatesQueriesTxModel struct {
	tx *sqlx.Tx
}

func NewUserupdatesQueriesTxModel(tx *sqlx.Tx) UserupdatesQueriesTxModel {
	return &defaultUserupdatesQueriesTxModel{tx: tx}
}

func (m *defaultUserupdatesQueriesModel) CountVisibleDialogs(ctx context.Context, userId int64) (*VisibleDialogCountRow, error) {
	var rValue VisibleDialogCountRow
	query := "select COUNT(*) as count from user_dialogs where user_id = ? and hidden = 0"

	err := m.db.QueryRowPartial(ctx, &rValue, query, userId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultUserupdatesQueriesTxModel) CountVisibleDialogs(userId int64) (*VisibleDialogCountRow, error) {
	var rValue VisibleDialogCountRow
	query := "select COUNT(*) as count from user_dialogs where user_id = ? and hidden = 0"

	err := m.tx.QueryRowPartial(&rValue, query, userId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultUserupdatesQueriesModel) SumUnreadDialogs(ctx context.Context, userId int64) (*UnreadDialogCountRow, error) {
	var rValue UnreadDialogCountRow
	query := "select COALESCE(SUM(unread_count), 0) as count from user_dialogs where user_id = ? and hidden = 0"

	err := m.db.QueryRowPartial(ctx, &rValue, query, userId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultUserupdatesQueriesTxModel) SumUnreadDialogs(userId int64) (*UnreadDialogCountRow, error) {
	var rValue UnreadDialogCountRow
	query := "select COALESCE(SUM(unread_count), 0) as count from user_dialogs where user_id = ? and hidden = 0"

	err := m.tx.QueryRowPartial(&rValue, query, userId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}
