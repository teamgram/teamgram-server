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

type bizAuthUpdatePayloadsModel interface {
	InsertIgnore(ctx context.Context, data *AuthUpdatePayloads) (lastInsertId, rowsAffected int64, err error)
	SelectByPayloadId(ctx context.Context, payloadId string) (*AuthUpdatePayloads, error)
}

type AuthUpdatePayloadsTxModel interface {
	InsertIgnore(data *AuthUpdatePayloads) (lastInsertId, rowsAffected int64, err error)
	SelectByPayloadId(payloadId string) (*AuthUpdatePayloads, error)
}

type defaultAuthUpdatePayloadsTxModel struct {
	tx *sqlx.Tx
}

func NewAuthUpdatePayloadsTxModel(tx *sqlx.Tx) AuthUpdatePayloadsTxModel {
	return &defaultAuthUpdatePayloadsTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into auth_update_payloads(payload_id, update_type, codec, layer, tl_bytes, payload_hash, expire_at) values (:payload_id, :update_type, :codec, :layer, :tl_bytes, :payload_hash, :expire_at)
func (m *defaultAuthUpdatePayloadsModel) InsertIgnore(ctx context.Context, data *AuthUpdatePayloads) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into auth_update_payloads(payload_id, update_type, codec, layer, tl_bytes, payload_hash, expire_at) values (:payload_id, :update_type, :codec, :layer, :tl_bytes, :payload_hash, :expire_at)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("auth_update_payloads.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_update_payloads.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_update_payloads.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into auth_update_payloads(payload_id, update_type, codec, layer, tl_bytes, payload_hash, expire_at) values (:payload_id, :update_type, :codec, :layer, :tl_bytes, :payload_hash, :expire_at)
func (m *defaultAuthUpdatePayloadsTxModel) InsertIgnore(data *AuthUpdatePayloads) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into auth_update_payloads(payload_id, update_type, codec, layer, tl_bytes, payload_hash, expire_at) values (:payload_id, :update_type, :codec, :layer, :tl_bytes, :payload_hash, :expire_at)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("auth_update_payloads.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_update_payloads.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_update_payloads.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectByPayloadId
// select payload_id, update_type, codec, layer, tl_bytes, payload_hash, expire_at from auth_update_payloads where payload_id = :payload_id limit 1
func (m *defaultAuthUpdatePayloadsModel) SelectByPayloadId(ctx context.Context, payloadId string) (rValue *AuthUpdatePayloads, err error) {

	var (
		query = "select payload_id, update_type, codec, layer, tl_bytes, payload_hash, expire_at from auth_update_payloads where payload_id = ? limit 1"
		do    = &AuthUpdatePayloads{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, payloadId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_update_payloads",
				Key:      fmt.Sprintf("payload_id=%v", payloadId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_update_payloads.SelectByPayloadId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByPayloadId
// select payload_id, update_type, codec, layer, tl_bytes, payload_hash, expire_at from auth_update_payloads where payload_id = :payload_id limit 1
func (m *defaultAuthUpdatePayloadsTxModel) SelectByPayloadId(payloadId string) (rValue *AuthUpdatePayloads, err error) {
	var (
		query = "select payload_id, update_type, codec, layer, tl_bytes, payload_hash, expire_at from auth_update_payloads where payload_id = ? limit 1"
		do    = &AuthUpdatePayloads{}
	)
	err = m.tx.QueryRowPartial(do, query, payloadId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_update_payloads",
				Key:      fmt.Sprintf("payload_id=%v", payloadId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_update_payloads.SelectByPayloadId: %w", err)
		return
	}
	rValue = do

	return
}
