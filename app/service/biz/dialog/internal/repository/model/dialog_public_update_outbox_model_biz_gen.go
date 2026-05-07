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
	MarkPublishing(ctx context.Context, status int32, leaseOwner string, leaseUntil int64, outboxId int64) (rowsAffected int64, err error)
	MarkPublishedPTS(ctx context.Context, status int32, publishedPts int64, publishedPtsCount int32, leaseUntil int64, outboxId int64) (rowsAffected int64, err error)
	MarkPublishedAuthSeq(ctx context.Context, status int32, publishedSeq int64, publishedDate int32, leaseUntil int64, outboxId int64) (rowsAffected int64, err error)
	MarkRetryable(ctx context.Context, status int32, attemptCount int32, nextRetryAt int64, leaseUntil int64, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error)
	MarkBlocked(ctx context.Context, status int32, leaseUntil int64, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error)
	ResetBlocked(ctx context.Context, status int32, nextRetryAt int64, leaseUntil int64, oldStatus int32, outboxId int64) (rowsAffected int64, err error)
}

type DialogPublicUpdateOutboxTxModel interface {
	Insert(data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error)
	InsertIgnore(data *DialogPublicUpdateOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectByTargetOperation(targetUserId int64, operationId string, deliveryPath string, publicUpdateType string) (*DialogPublicUpdateOutbox, error)
	MarkPublishing(status int32, leaseOwner string, leaseUntil int64, outboxId int64) (rowsAffected int64, err error)
	MarkPublishedPTS(status int32, publishedPts int64, publishedPtsCount int32, leaseUntil int64, outboxId int64) (rowsAffected int64, err error)
	MarkPublishedAuthSeq(status int32, publishedSeq int64, publishedDate int32, leaseUntil int64, outboxId int64) (rowsAffected int64, err error)
	MarkRetryable(status int32, attemptCount int32, nextRetryAt int64, leaseUntil int64, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error)
	MarkBlocked(status int32, leaseUntil int64, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error)
	ResetBlocked(status int32, nextRetryAt int64, leaseUntil int64, oldStatus int32, outboxId int64) (rowsAffected int64, err error)
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

// MarkPublishing
// update dialog_public_update_outbox set `status` = :status, lease_owner = :lease_owner, lease_until = :lease_until where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxModel) MarkPublishing(ctx context.Context, status int32, leaseOwner string, leaseUntil int64, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_public_update_outbox set `status` = ?, lease_owner = ?, lease_until = ? where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseOwner, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishing rows affected: %w", err)
		return
	}

	return
}

// MarkPublishing
// update dialog_public_update_outbox set `status` = :status, lease_owner = :lease_owner, lease_until = :lease_until where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxTxModel) MarkPublishing(status int32, leaseOwner string, leaseUntil int64, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_public_update_outbox set `status` = ?, lease_owner = ?, lease_until = ? where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseOwner, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishing rows affected: %w", err)
		return
	}

	return
}

// MarkPublishedPTS
// update dialog_public_update_outbox set `status` = :status, published_pts = :published_pts, published_pts_count = :published_pts_count, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxModel) MarkPublishedPTS(ctx context.Context, status int32, publishedPts int64, publishedPtsCount int32, leaseUntil int64, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_public_update_outbox set `status` = ?, published_pts = ?, published_pts_count = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, publishedPts, publishedPtsCount, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishedPTS exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishedPTS rows affected: %w", err)
		return
	}

	return
}

// MarkPublishedPTS
// update dialog_public_update_outbox set `status` = :status, published_pts = :published_pts, published_pts_count = :published_pts_count, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxTxModel) MarkPublishedPTS(status int32, publishedPts int64, publishedPtsCount int32, leaseUntil int64, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_public_update_outbox set `status` = ?, published_pts = ?, published_pts_count = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, publishedPts, publishedPtsCount, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishedPTS exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishedPTS rows affected: %w", err)
		return
	}

	return
}

// MarkPublishedAuthSeq
// update dialog_public_update_outbox set `status` = :status, published_seq = :published_seq, published_date = :published_date, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxModel) MarkPublishedAuthSeq(ctx context.Context, status int32, publishedSeq int64, publishedDate int32, leaseUntil int64, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_public_update_outbox set `status` = ?, published_seq = ?, published_date = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, publishedSeq, publishedDate, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishedAuthSeq exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishedAuthSeq rows affected: %w", err)
		return
	}

	return
}

// MarkPublishedAuthSeq
// update dialog_public_update_outbox set `status` = :status, published_seq = :published_seq, published_date = :published_date, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxTxModel) MarkPublishedAuthSeq(status int32, publishedSeq int64, publishedDate int32, leaseUntil int64, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_public_update_outbox set `status` = ?, published_seq = ?, published_date = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, publishedSeq, publishedDate, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishedAuthSeq exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkPublishedAuthSeq rows affected: %w", err)
		return
	}

	return
}

// MarkRetryable
// update dialog_public_update_outbox set `status` = :status, attempt_count = :attempt_count, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_kind = :last_error_kind, last_error_message = :last_error_message where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxModel) MarkRetryable(ctx context.Context, status int32, attemptCount int32, nextRetryAt int64, leaseUntil int64, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_public_update_outbox set `status` = ?, attempt_count = ?, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ? where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, attemptCount, nextRetryAt, leaseUntil, lastErrorKind, lastErrorMessage, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkRetryable exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkRetryable rows affected: %w", err)
		return
	}

	return
}

// MarkRetryable
// update dialog_public_update_outbox set `status` = :status, attempt_count = :attempt_count, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_kind = :last_error_kind, last_error_message = :last_error_message where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxTxModel) MarkRetryable(status int32, attemptCount int32, nextRetryAt int64, leaseUntil int64, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_public_update_outbox set `status` = ?, attempt_count = ?, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ? where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, attemptCount, nextRetryAt, leaseUntil, lastErrorKind, lastErrorMessage, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkRetryable exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkRetryable rows affected: %w", err)
		return
	}

	return
}

// MarkBlocked
// update dialog_public_update_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_kind = :last_error_kind, last_error_message = :last_error_message where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxModel) MarkBlocked(ctx context.Context, status int32, leaseUntil int64, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_public_update_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ? where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseUntil, lastErrorKind, lastErrorMessage, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkBlocked rows affected: %w", err)
		return
	}

	return
}

// MarkBlocked
// update dialog_public_update_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_kind = :last_error_kind, last_error_message = :last_error_message where outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxTxModel) MarkBlocked(status int32, leaseUntil int64, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_public_update_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ? where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseUntil, lastErrorKind, lastErrorMessage, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.MarkBlocked rows affected: %w", err)
		return
	}

	return
}

// ResetBlocked
// update dialog_public_update_outbox set `status` = :status, attempt_count = 0, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where `status` = :old_status and outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxModel) ResetBlocked(ctx context.Context, status int32, nextRetryAt int64, leaseUntil int64, oldStatus int32, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_public_update_outbox set `status` = ?, attempt_count = 0, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where `status` = ? and outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, nextRetryAt, leaseUntil, oldStatus, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.ResetBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.ResetBlocked rows affected: %w", err)
		return
	}

	return
}

// ResetBlocked
// update dialog_public_update_outbox set `status` = :status, attempt_count = 0, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where `status` = :old_status and outbox_id = :outbox_id
func (m *defaultDialogPublicUpdateOutboxTxModel) ResetBlocked(status int32, nextRetryAt int64, leaseUntil int64, oldStatus int32, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_public_update_outbox set `status` = ?, attempt_count = 0, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where `status` = ? and outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, nextRetryAt, leaseUntil, oldStatus, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.ResetBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_public_update_outbox.ResetBlocked rows affected: %w", err)
		return
	}

	return
}
