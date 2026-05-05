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

type bizDialogPublicUpdateOutboxModel interface {
	Insert(ctx context.Context, data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error)
	InsertIgnore(ctx context.Context, data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectByTargetOperation(ctx context.Context, targetUserId int64, operationId string, deliveryPath string, publicUpdateType string) (*DialogPublicUpdateOutbox, error)
}

type DialogPublicUpdateOutboxTxModel interface {
	Insert(data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error)
	InsertIgnore(data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectByTargetOperation(targetUserId int64, operationId string, deliveryPath string, publicUpdateType string) (*DialogPublicUpdateOutbox, error)
}

type defaultDialogPublicUpdateOutboxTxModel struct {
	tx *sqlx.Tx
}

func NewDialogPublicUpdateOutboxTxModel(tx *sqlx.Tx) DialogPublicUpdateOutboxTxModel {
	return &defaultDialogPublicUpdateOutboxTxModel{tx: tx}
}

// Insert
// insert into dialog_public_update_outbox(outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :source_user_id, :source_perm_auth_key_id, :target_user_id, :target_auth_policy, :operation_id, :delivery_path, :public_update_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_pts, :published_pts_count, :published_seq, :published_date, :last_error_kind, :last_error_message)
func (m *defaultDialogPublicUpdateOutboxModel) Insert(ctx context.Context, data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_public_update_outbox(outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :source_user_id, :source_perm_auth_key_id, :target_user_id, :target_auth_policy, :operation_id, :delivery_path, :public_update_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_pts, :published_pts_count, :published_seq, :published_date, :last_error_kind, :last_error_message)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into dialog_public_update_outbox(outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :source_user_id, :source_perm_auth_key_id, :target_user_id, :target_auth_policy, :operation_id, :delivery_path, :public_update_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_pts, :published_pts_count, :published_seq, :published_date, :last_error_kind, :last_error_message)
func (m *defaultDialogPublicUpdateOutboxTxModel) Insert(data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_public_update_outbox(outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :source_user_id, :source_perm_auth_key_id, :target_user_id, :target_auth_policy, :operation_id, :delivery_path, :public_update_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_pts, :published_pts_count, :published_seq, :published_date, :last_error_kind, :last_error_message)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.Insert rows affected: %w", err)
	}

	return
}

// InsertIgnore
// insert ignore into dialog_public_update_outbox(outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :source_user_id, :source_perm_auth_key_id, :target_user_id, :target_auth_policy, :operation_id, :delivery_path, :public_update_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_pts, :published_pts_count, :published_seq, :published_date, :last_error_kind, :last_error_message)
func (m *defaultDialogPublicUpdateOutboxModel) InsertIgnore(ctx context.Context, data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialog_public_update_outbox(outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :source_user_id, :source_perm_auth_key_id, :target_user_id, :target_auth_policy, :operation_id, :delivery_path, :public_update_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_pts, :published_pts_count, :published_seq, :published_date, :last_error_kind, :last_error_message)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into dialog_public_update_outbox(outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :source_user_id, :source_perm_auth_key_id, :target_user_id, :target_auth_policy, :operation_id, :delivery_path, :public_update_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_pts, :published_pts_count, :published_seq, :published_date, :last_error_kind, :last_error_message)
func (m *defaultDialogPublicUpdateOutboxTxModel) InsertIgnore(data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialog_public_update_outbox(outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :source_user_id, :source_perm_auth_key_id, :target_user_id, :target_auth_policy, :operation_id, :delivery_path, :public_update_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_pts, :published_pts_count, :published_seq, :published_date, :last_error_kind, :last_error_message)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectByTargetOperation
// select outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message from dialog_public_update_outbox where target_user_id = :target_user_id and operation_id = :operation_id and delivery_path = :delivery_path and public_update_type = :public_update_type limit 1
func (m *defaultDialogPublicUpdateOutboxModel) SelectByTargetOperation(ctx context.Context, targetUserId int64, operationId string, deliveryPath string, publicUpdateType string) (rValue *DialogPublicUpdateOutbox, err error) {

	var (
		query = "select outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message from dialog_public_update_outbox where target_user_id = ? and operation_id = ? and delivery_path = ? and public_update_type = ? limit 1"
		do    = &DialogPublicUpdateOutbox{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, targetUserId, operationId, deliveryPath, publicUpdateType)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_public_update_outbox",
				Key:      fmt.Sprintf("target_user_id=%v,operation_id=%v,delivery_path=%v,public_update_type=%v", targetUserId, operationId, deliveryPath, publicUpdateType),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_public_update_outbox.SelectByTargetOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByTargetOperation
// select outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message from dialog_public_update_outbox where target_user_id = :target_user_id and operation_id = :operation_id and delivery_path = :delivery_path and public_update_type = :public_update_type limit 1
func (m *defaultDialogPublicUpdateOutboxTxModel) SelectByTargetOperation(targetUserId int64, operationId string, deliveryPath string, publicUpdateType string) (rValue *DialogPublicUpdateOutbox, err error) {
	var (
		query = "select outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message from dialog_public_update_outbox where target_user_id = ? and operation_id = ? and delivery_path = ? and public_update_type = ? limit 1"
		do    = &DialogPublicUpdateOutbox{}
	)
	err = m.tx.QueryRowPartial(do, query, targetUserId, operationId, deliveryPath, publicUpdateType)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_public_update_outbox",
				Key:      fmt.Sprintf("target_user_id=%v,operation_id=%v,delivery_path=%v,public_update_type=%v", targetUserId, operationId, deliveryPath, publicUpdateType),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_public_update_outbox.SelectByTargetOperation: %w", err)
		return
	}
	rValue = do

	return
}
