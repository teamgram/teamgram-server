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

type SavedDialogNextPinOrderRow struct {
	NextPinOrder int64 `db:"next_pin_order"`
}

type DialogRepositoryQueriesModel interface {
	SelectSavedDialogNextPinOrder(ctx context.Context, userId int64) (*SavedDialogNextPinOrderRow, error)
}

type DialogRepositoryQueriesTxModel interface {
	SelectSavedDialogNextPinOrder(userId int64) (*SavedDialogNextPinOrderRow, error)
}

type defaultDialogRepositoryQueriesModel struct {
	db *sqlx.DB
}

func NewDialogRepositoryQueriesModel(db *sqlx.DB) DialogRepositoryQueriesModel {
	return &defaultDialogRepositoryQueriesModel{db: db}
}

type defaultDialogRepositoryQueriesTxModel struct {
	tx *sqlx.Tx
}

func NewDialogRepositoryQueriesTxModel(tx *sqlx.Tx) DialogRepositoryQueriesTxModel {
	return &defaultDialogRepositoryQueriesTxModel{tx: tx}
}

func (m *defaultDialogRepositoryQueriesModel) SelectSavedDialogNextPinOrder(ctx context.Context, userId int64) (*SavedDialogNextPinOrderRow, error) {
	var rValue SavedDialogNextPinOrderRow
	query := "select COALESCE(MAX(pin_order), 0) + 1 as next_pin_order from saved_dialogs where user_id = ? and pinned = 1 and deleted = 0 limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, userId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultDialogRepositoryQueriesTxModel) SelectSavedDialogNextPinOrder(userId int64) (*SavedDialogNextPinOrderRow, error) {
	var rValue SavedDialogNextPinOrderRow
	query := "select COALESCE(MAX(pin_order), 0) + 1 as next_pin_order from saved_dialogs where user_id = ? and pinned = 1 and deleted = 0 limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, userId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}
