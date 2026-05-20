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
	authUpdatePayloadsFieldNames          = builder.RawFieldNames(&AuthUpdatePayloads{})
	authUpdatePayloadsRows                = strings.Join(authUpdatePayloadsFieldNames, ",")
	authUpdatePayloadsRowsExpectAutoSet   = strings.Join(stringx.Remove(authUpdatePayloadsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authUpdatePayloadsRowsWithPlaceHolder = strings.Join(stringx.Remove(authUpdatePayloadsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	authUpdatePayloadsModel interface {
		Insert2(ctx context.Context, data *AuthUpdatePayloads) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AuthUpdatePayloads, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]AuthUpdatePayloads, error)
		Update2(ctx context.Context, data *AuthUpdatePayloads) error
		Delete2(ctx context.Context, id int64) error

		FindOneByPayloadId(ctx context.Context, payloadId string) (*AuthUpdatePayloads, error)
		FindListByPayloadIdList(ctx context.Context, payloadId ...string) ([]AuthUpdatePayloads, error)
	}

	defaultAuthUpdatePayloadsModel struct {
		db *sqlx.DB
	}

	AuthUpdatePayloads struct {
		Id          int64  `db:"id" json:"id"`
		PayloadId   string `db:"payload_id" json:"payload_id"`
		UpdateType  string `db:"update_type" json:"update_type"`
		Codec       int32  `db:"codec" json:"codec"`
		Layer       int32  `db:"layer" json:"layer"`
		TlBytes     []byte `db:"tl_bytes" json:"tl_bytes"`
		PayloadHash []byte `db:"payload_hash" json:"payload_hash"`
		ExpireAt    int64  `db:"expire_at" json:"expire_at"`
	}
)

func newAuthUpdatePayloadsModel(db *sqlx.DB) *defaultAuthUpdatePayloadsModel {
	return &defaultAuthUpdatePayloadsModel{
		db: db,
	}
}

func (m *defaultAuthUpdatePayloadsModel) Insert2(ctx context.Context, data *AuthUpdatePayloads) (sql.Result, error) {
	tableName := "auth_update_payloads"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?)", tableName, authUpdatePayloadsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.PayloadId, data.UpdateType, data.Codec, data.Layer, data.TlBytes, data.PayloadHash, data.ExpireAt)
	if err != nil {
		return nil, fmt.Errorf("auth_update_payloads.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultAuthUpdatePayloadsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "auth_update_payloads"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("auth_update_payloads.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthUpdatePayloadsModel) FindOne(ctx context.Context, id int64) (*AuthUpdatePayloads, error) {
	tableName := "auth_update_payloads"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", authUpdatePayloadsRows, tableName)
	var resp AuthUpdatePayloads

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_update_payloads",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("auth_update_payloads.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthUpdatePayloadsModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthUpdatePayloads, error) {
	if len(id) == 0 {
		return []AuthUpdatePayloads{}, nil
	}
	tableName := "auth_update_payloads"

	query := fmt.Sprintf("select %s from %s where id in (%s)", authUpdatePayloadsRows, tableName, sqlx.InInt64List(id))

	var resp []AuthUpdatePayloads
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []AuthUpdatePayloads{}, nil
		}
		return nil, fmt.Errorf("auth_update_payloads.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultAuthUpdatePayloadsModel) Update2(ctx context.Context, data *AuthUpdatePayloads) error {
	tableName := "auth_update_payloads"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, authUpdatePayloadsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.PayloadId, data.UpdateType, data.Codec, data.Layer, data.TlBytes, data.PayloadHash, data.ExpireAt, data.Id)
	if err != nil {
		return fmt.Errorf("auth_update_payloads.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthUpdatePayloadsModel) FindOneByPayloadId(ctx context.Context, payloadId string) (*AuthUpdatePayloads, error) {
	tableName := "auth_update_payloads"
	query := fmt.Sprintf("select %s from %s where payload_id = ? limit 1", authUpdatePayloadsRows, tableName)
	var resp AuthUpdatePayloads

	err := m.db.QueryRowPartial(ctx, &resp, query, payloadId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_update_payloads",
				Key:      fmt.Sprintf("payload_id=%v", payloadId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("auth_update_payloads.FindOneByPayloadId: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthUpdatePayloadsModel) FindListByPayloadIdList(ctx context.Context, payloadId ...string) ([]AuthUpdatePayloads, error) {
	if len(payloadId) == 0 {
		return []AuthUpdatePayloads{}, nil
	}
	tableName := "auth_update_payloads"

	query := fmt.Sprintf("select %s from %s where payload_id in (%s)", authUpdatePayloadsRows, tableName, sqlx.InStringList(payloadId))
	var resp []AuthUpdatePayloads
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []AuthUpdatePayloads{}, nil
		}
		return nil, fmt.Errorf("auth_update_payloads.FindListByPayloadIdList: %w", err)
	}

	return resp, nil
}
