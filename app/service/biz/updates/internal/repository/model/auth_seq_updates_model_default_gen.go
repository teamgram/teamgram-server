/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	auth_seq_updatesFieldNames          = builder.RawFieldNames(&AuthSeqUpdates{})
	auth_seq_updatesRows                = strings.Join(auth_seq_updatesFieldNames, ",")
	auth_seq_updatesRowsExpectAutoSet   = strings.Join(stringx.Remove(auth_seq_updatesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	auth_seq_updatesRowsWithPlaceHolder = strings.Join(stringx.Remove(auth_seq_updatesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTAuthSeqUpdatesIdPrefix = "cache:t:auth_seq_updates:id:"

	cacheAuthSeqUpdatesIdPrefix = "cache#AuthSeqUpdates#id"

	cacheAuthSeqUpdatesAuthIdUserIdSeqPrefix = "cache#AuthId#UserId#Seq"
)

type (
	auth_seq_updatesModel interface {
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
	query := fmt.Sprintf("insert into `auth_seq_updates` (%s) values (?, ?, ?, ?, ?, ?)", auth_seq_updatesRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.AuthId, data.UserId, data.Seq, data.UpdateType, data.UpdateData, data.Date2)
}

func (m *defaultAuthSeqUpdatesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auth_seq_updates` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultAuthSeqUpdatesModel) FindOne(ctx context.Context, id int64) (*AuthSeqUpdates, error) {
	query := fmt.Sprintf("select %s from auth_seq_updates where id = ? limit 1", auth_seq_updatesRows)
	var resp AuthSeqUpdates
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAuthSeqUpdatesModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthSeqUpdates, error) {
	if len(id) == 0 {
		return []AuthSeqUpdates{}, nil
	}

	query := fmt.Sprintf("select %s from auth_seq_updates where id in (%s)", auth_seq_updatesRows, sqlx.InInt64List(id))

	var resp []AuthSeqUpdates
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAuthSeqUpdatesModel) Update2(ctx context.Context, data *AuthSeqUpdates) error {
	query := fmt.Sprintf("update `auth_seq_updates` set %s where `id` = ?", auth_seq_updatesRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.AuthId, data.UserId, data.Seq, data.UpdateType, data.UpdateData, data.Date2, data.Id)
	return err
}

func (m *defaultAuthSeqUpdatesModel) FindOneByAuthIdUserIdSeq(ctx context.Context, authId int64, userId int64, seq int32) (*AuthSeqUpdates, error) {
	query := fmt.Sprintf("select %s from auth_seq_updates where auth_id = ? AND user_id = ? AND seq = ? limit 1", auth_seq_updatesRows)
	var resp AuthSeqUpdates
	err := m.db.QueryRowPartial(ctx, &resp, query, authId, userId, seq)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAuthSeqUpdatesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheAuthSeqUpdatesIdPrefix, primary)
}

func (m *defaultAuthSeqUpdatesModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from auth_seq_updates where id = ? limit 1", auth_seq_updatesRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
