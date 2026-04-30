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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	ptsUpdatesNgenFieldNames          = builder.RawFieldNames(&PtsUpdatesNgen{})
	ptsUpdatesNgenRows                = strings.Join(ptsUpdatesNgenFieldNames, ",")
	ptsUpdatesNgenRowsExpectAutoSet   = strings.Join(stringx.Remove(ptsUpdatesNgenFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	ptsUpdatesNgenRowsWithPlaceHolder = strings.Join(stringx.Remove(ptsUpdatesNgenFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	ptsUpdatesNgenModel interface {
		Insert2(ctx context.Context, data *PtsUpdatesNgen) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PtsUpdatesNgen, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]PtsUpdatesNgen, error)
		Update2(ctx context.Context, data *PtsUpdatesNgen) error
		Delete2(ctx context.Context, id int64) error
	}

	defaultPtsUpdatesNgenModel struct {
		db *sqlx.DB
	}

	PtsUpdatesNgen struct {
		Id     int64 `db:"id" json:"id"`
		MinSeq int64 `db:"min_seq" json:"min_seq"`
		MaxSeq int64 `db:"max_seq" json:"max_seq"`
	}
)

func newPtsUpdatesNgenModel(db *sqlx.DB) *defaultPtsUpdatesNgenModel {
	return &defaultPtsUpdatesNgenModel{
		db: db,
	}
}

func (m *defaultPtsUpdatesNgenModel) Insert2(ctx context.Context, data *PtsUpdatesNgen) (sql.Result, error) {
	tableName := "pts_updates_ngen"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?)", tableName, ptsUpdatesNgenRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.MinSeq, data.MaxSeq)
	if err != nil {
		return nil, fmt.Errorf("pts_updates_ngen.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultPtsUpdatesNgenModel) Delete2(ctx context.Context, id int64) error {
	tableName := "pts_updates_ngen"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("pts_updates_ngen.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultPtsUpdatesNgenModel) FindOne(ctx context.Context, id int64) (*PtsUpdatesNgen, error) {
	tableName := "pts_updates_ngen"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", ptsUpdatesNgenRows, tableName)
	var resp PtsUpdatesNgen

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "pts_updates_ngen",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("pts_updates_ngen.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultPtsUpdatesNgenModel) FindListByIdList(ctx context.Context, id ...int64) ([]PtsUpdatesNgen, error) {
	if len(id) == 0 {
		return []PtsUpdatesNgen{}, nil
	}
	tableName := "pts_updates_ngen"

	query := fmt.Sprintf("select %s from %s where id in (%s)", ptsUpdatesNgenRows, tableName, sqlx.InInt64List(id))

	var resp []PtsUpdatesNgen
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []PtsUpdatesNgen{}, nil
		}
		return nil, fmt.Errorf("pts_updates_ngen.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultPtsUpdatesNgenModel) Update2(ctx context.Context, data *PtsUpdatesNgen) error {
	tableName := "pts_updates_ngen"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, ptsUpdatesNgenRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.MinSeq, data.MaxSeq, data.Id)
	if err != nil {
		return fmt.Errorf("pts_updates_ngen.Update2 exec: %w", err)
	}

	return nil
}
