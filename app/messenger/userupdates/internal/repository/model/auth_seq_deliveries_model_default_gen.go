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
	authSeqDeliveriesFieldNames          = builder.RawFieldNames(&AuthSeqDeliveries{})
	authSeqDeliveriesRows                = strings.Join(authSeqDeliveriesFieldNames, ",")
	authSeqDeliveriesRowsExpectAutoSet   = strings.Join(stringx.Remove(authSeqDeliveriesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authSeqDeliveriesRowsWithPlaceHolder = strings.Join(stringx.Remove(authSeqDeliveriesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	authSeqDeliveriesModel interface {
		Insert2(ctx context.Context, data *AuthSeqDeliveries) (sql.Result, error)

		FindOneByUserIdPermAuthKeyIdOperationId(ctx context.Context, userId int64, permAuthKeyId int64, operationId string) (*AuthSeqDeliveries, error)
	}

	defaultAuthSeqDeliveriesModel struct {
		db *sqlx.DB
	}

	AuthSeqDeliveries struct {
		UserId              int64  `db:"user_id" json:"user_id"`
		PermAuthKeyId       int64  `db:"perm_auth_key_id" json:"perm_auth_key_id"`
		Seq                 int64  `db:"seq" json:"seq"`
		Date                int64  `db:"date" json:"date"`
		PayloadId           string `db:"payload_id" json:"payload_id"`
		ReplayPolicy        string `db:"replay_policy" json:"replay_policy"`
		SourcePermAuthKeyId int64  `db:"source_perm_auth_key_id" json:"source_perm_auth_key_id"`
		VisibilityPolicy    string `db:"visibility_policy" json:"visibility_policy"`
		OperationId         string `db:"operation_id" json:"operation_id"`
		ExpireAt            int64  `db:"expire_at" json:"expire_at"`
	}
)

func newAuthSeqDeliveriesModel(db *sqlx.DB) *defaultAuthSeqDeliveriesModel {
	return &defaultAuthSeqDeliveriesModel{
		db: db,
	}
}

func (m *defaultAuthSeqDeliveriesModel) Insert2(ctx context.Context, data *AuthSeqDeliveries) (sql.Result, error) {
	tableName := "auth_seq_deliveries"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, authSeqDeliveriesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PermAuthKeyId, data.Seq, data.Date, data.PayloadId, data.ReplayPolicy, data.SourcePermAuthKeyId, data.VisibilityPolicy, data.OperationId, data.ExpireAt)
	if err != nil {
		return nil, fmt.Errorf("auth_seq_deliveries.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultAuthSeqDeliveriesModel) FindOneByUserIdPermAuthKeyIdOperationId(ctx context.Context, userId int64, permAuthKeyId int64, operationId string) (*AuthSeqDeliveries, error) {
	tableName := "auth_seq_deliveries"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND perm_auth_key_id = ? AND operation_id = ? limit 1", authSeqDeliveriesRows, tableName)
	var resp AuthSeqDeliveries

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, permAuthKeyId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_seq_deliveries",
				Key:      fmt.Sprintf("user_id=%v,perm_auth_key_id=%v,operation_id=%v", userId, permAuthKeyId, operationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("auth_seq_deliveries.FindOneByUserIdPermAuthKeyIdOperationId: %w", err)
	}

	return &resp, nil
}
