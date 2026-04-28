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
	authSeqUpdatesFieldNames          = builder.RawFieldNames(&AuthSeqUpdates{})
	authSeqUpdatesRows                = strings.Join(authSeqUpdatesFieldNames, ",")
	authSeqUpdatesRowsExpectAutoSet   = strings.Join(stringx.Remove(authSeqUpdatesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authSeqUpdatesRowsWithPlaceHolder = strings.Join(stringx.Remove(authSeqUpdatesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	authSeqUpdatesModel interface {
		Insert2(ctx context.Context, data *AuthSeqUpdates) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AuthSeqUpdates, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]AuthSeqUpdates, error)
		Update2(ctx context.Context, data *AuthSeqUpdates) error
		Delete2(ctx context.Context, id int64) error

		FindOneByAuthIdUserIdSeq(ctx context.Context, authId int64, userId int64, seq int32) (*AuthSeqUpdates, error)
	}

	defaultAuthSeqUpdatesModel struct {
		db *sqlx.DB
	}

	AuthSeqUpdates struct {
		Id         int64  `db:"id" json:"id"`
		AuthId     int64  `db:"auth_id" json:"auth_id"`
		UserId     int64  `db:"user_id" json:"user_id"`
		Seq        int32  `db:"seq" json:"seq"`
		UpdateType int32  `db:"update_type" json:"update_type"`
		UpdateData string `db:"update_data" json:"update_data"`
		Date2      int64  `db:"date2" json:"date2"`
	}
)

func newAuthSeqUpdatesModel(db *sqlx.DB) *defaultAuthSeqUpdatesModel {
	return &defaultAuthSeqUpdatesModel{
		db: db,
	}
}

func (m *defaultAuthSeqUpdatesModel) Insert2(ctx context.Context, data *AuthSeqUpdates) (sql.Result, error) {
	query := fmt.Sprintf("insert into `auth_seq_updates` (%s) values (?, ?, ?, ?, ?, ?)", authSeqUpdatesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.AuthId, data.UserId, data.Seq, data.UpdateType, data.UpdateData, data.Date2)
	if err != nil {
		return nil, fmt.Errorf("auth_seq_updates.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultAuthSeqUpdatesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auth_seq_updates` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("auth_seq_updates.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthSeqUpdatesModel) FindOne(ctx context.Context, id int64) (*AuthSeqUpdates, error) {
	query := fmt.Sprintf("select %s from auth_seq_updates where id = ? limit 1", authSeqUpdatesRows)
	var resp AuthSeqUpdates

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_updates",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("auth_seq_updates.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthSeqUpdatesModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthSeqUpdates, error) {
	if len(id) == 0 {
		return []AuthSeqUpdates{}, nil
	}

	query := fmt.Sprintf("select %s from auth_seq_updates where id in (%s)", authSeqUpdatesRows, sqlx.InInt64List(id))

	var resp []AuthSeqUpdates
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []AuthSeqUpdates{}, nil
		}
		return nil, fmt.Errorf("auth_seq_updates.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultAuthSeqUpdatesModel) Update2(ctx context.Context, data *AuthSeqUpdates) error {
	query := fmt.Sprintf("update `auth_seq_updates` set %s where `id` = ?", authSeqUpdatesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.AuthId, data.UserId, data.Seq, data.UpdateType, data.UpdateData, data.Date2, data.Id)
	if err != nil {
		return fmt.Errorf("auth_seq_updates.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthSeqUpdatesModel) FindOneByAuthIdUserIdSeq(ctx context.Context, authId int64, userId int64, seq int32) (*AuthSeqUpdates, error) {
	query := fmt.Sprintf("select %s from auth_seq_updates where auth_id = ? AND user_id = ? AND seq = ? limit 1", authSeqUpdatesRows)
	var resp AuthSeqUpdates

	err := m.db.QueryRowPartial(ctx, &resp, query, authId, userId, seq)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_updates",
				Key:      fmt.Sprintf("auth_id=%v,user_id=%v,seq=%v", authId, userId, seq),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("auth_seq_updates.FindOneByAuthIdUserIdSeq: %w", err)
	}

	return &resp, nil
}
