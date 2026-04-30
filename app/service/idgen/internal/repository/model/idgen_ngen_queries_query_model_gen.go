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

type IdgenNgenRow struct {
	MaxSeq int64 `db:"max_seq"`
}

type IdgenNgenQueriesModel interface {
	GetMaxSeqQuery(ctx context.Context, id int64) (*IdgenNgenRow, error)
}

type IdgenNgenQueriesTxModel interface {
	GetMaxSeqQuery(id int64) (*IdgenNgenRow, error)
}

type defaultIdgenNgenQueriesModel struct {
	db *sqlx.DB
}

func NewIdgenNgenQueriesModel(db *sqlx.DB) IdgenNgenQueriesModel {
	return &defaultIdgenNgenQueriesModel{db: db}
}

type defaultIdgenNgenQueriesTxModel struct {
	tx *sqlx.Tx
}

func NewIdgenNgenQueriesTxModel(tx *sqlx.Tx) IdgenNgenQueriesTxModel {
	return &defaultIdgenNgenQueriesTxModel{tx: tx}
}

func (m *defaultIdgenNgenQueriesModel) GetMaxSeqQuery(ctx context.Context, id int64) (*IdgenNgenRow, error) {
	var rValue IdgenNgenRow
	query := "select max_seq from pts_updates_ngen where id = ? limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, id)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultIdgenNgenQueriesTxModel) GetMaxSeqQuery(id int64) (*IdgenNgenRow, error) {
	var rValue IdgenNgenRow
	query := "select max_seq from pts_updates_ngen where id = ? limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, id)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}
