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

type DialogAuthSeqOutboxRow struct {
	OutboxId             int64  `db:"outbox_id"`
	UserId               int64  `db:"user_id"`
	SourcePermAuthKeyId  int64  `db:"source_perm_auth_key_id"`
	TargetAuthPolicy     string `db:"target_auth_policy"`
	OperationId          string `db:"operation_id"`
	EventType            string `db:"event_type"`
	PeerType             int32  `db:"peer_type"`
	PeerId               int64  `db:"peer_id"`
	PayloadSchemaVersion int32  `db:"payload_schema_version"`
	Payload              []byte `db:"payload"`
	PayloadHash          []byte `db:"payload_hash"`
	Status               int32  `db:"status"`
	AttemptCount         int32  `db:"attempt_count"`
	NextRetryAt          int64  `db:"next_retry_at"`
	LeaseOwner           string `db:"lease_owner"`
	LeaseUntil           int64  `db:"lease_until"`
	PublishedSeq         int64  `db:"published_seq"`
	PublishedDate        int32  `db:"published_date"`
	LastErrorKind        string `db:"last_error_kind"`
	LastErrorMessage     string `db:"last_error_message"`
}

type DialogPublicUpdateOutboxRow struct {
	OutboxId             int64  `db:"outbox_id"`
	SourceUserId         int64  `db:"source_user_id"`
	SourcePermAuthKeyId  int64  `db:"source_perm_auth_key_id"`
	TargetUserId         int64  `db:"target_user_id"`
	TargetAuthPolicy     string `db:"target_auth_policy"`
	OperationId          string `db:"operation_id"`
	DeliveryPath         string `db:"delivery_path"`
	PublicUpdateType     string `db:"public_update_type"`
	PeerType             int32  `db:"peer_type"`
	PeerId               int64  `db:"peer_id"`
	PayloadSchemaVersion int32  `db:"payload_schema_version"`
	Payload              []byte `db:"payload"`
	PayloadHash          []byte `db:"payload_hash"`
	Status               int32  `db:"status"`
	AttemptCount         int32  `db:"attempt_count"`
	NextRetryAt          int64  `db:"next_retry_at"`
	LeaseOwner           string `db:"lease_owner"`
	LeaseUntil           int64  `db:"lease_until"`
	PublishedPts         int64  `db:"published_pts"`
	PublishedPtsCount    int32  `db:"published_pts_count"`
	PublishedSeq         int64  `db:"published_seq"`
	PublishedDate        int32  `db:"published_date"`
	LastErrorKind        string `db:"last_error_kind"`
	LastErrorMessage     string `db:"last_error_message"`
}

type DialogRepositoryQueriesModel interface {
	SelectSavedDialogNextPinOrder(ctx context.Context, userId int64) (*SavedDialogNextPinOrderRow, error)
	SelectAuthSeqOutboxClaimCandidates(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, readyBefore int64, publishingStatus int32, leaseExpiredBefore int64, limit int32) ([]DialogAuthSeqOutboxRow, error)
	SelectPublicUpdateOutboxClaimCandidates(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, readyBefore int64, publishingStatus int32, leaseExpiredBefore int64, limit int32) ([]DialogPublicUpdateOutboxRow, error)
}

type DialogRepositoryQueriesTxModel interface {
	SelectSavedDialogNextPinOrder(userId int64) (*SavedDialogNextPinOrderRow, error)
	SelectAuthSeqOutboxClaimCandidates(pendingStatus int32, failedRetryableStatus int32, readyBefore int64, publishingStatus int32, leaseExpiredBefore int64, limit int32) ([]DialogAuthSeqOutboxRow, error)
	SelectPublicUpdateOutboxClaimCandidates(pendingStatus int32, failedRetryableStatus int32, readyBefore int64, publishingStatus int32, leaseExpiredBefore int64, limit int32) ([]DialogPublicUpdateOutboxRow, error)
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

func (m *defaultDialogRepositoryQueriesModel) SelectAuthSeqOutboxClaimCandidates(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, readyBefore int64, publishingStatus int32, leaseExpiredBefore int64, limit int32) ([]DialogAuthSeqOutboxRow, error) {
	var rList []DialogAuthSeqOutboxRow
	query := "select outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message from dialog_auth_seq_outbox where (`status` in (?, ?) and next_retry_at > 0 and next_retry_at <= ?) or (`status` = ? and (lease_until = 0 or lease_until <= ?)) order by next_retry_at asc, outbox_id asc limit ? for update"

	err := m.db.QueryRowsPartial(ctx, &rList, query, pendingStatus, failedRetryableStatus, readyBefore, publishingStatus, leaseExpiredBefore, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultDialogRepositoryQueriesTxModel) SelectAuthSeqOutboxClaimCandidates(pendingStatus int32, failedRetryableStatus int32, readyBefore int64, publishingStatus int32, leaseExpiredBefore int64, limit int32) ([]DialogAuthSeqOutboxRow, error) {
	var rList []DialogAuthSeqOutboxRow
	query := "select outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message from dialog_auth_seq_outbox where (`status` in (?, ?) and next_retry_at > 0 and next_retry_at <= ?) or (`status` = ? and (lease_until = 0 or lease_until <= ?)) order by next_retry_at asc, outbox_id asc limit ? for update"

	err := m.tx.QueryRowsPartial(&rList, query, pendingStatus, failedRetryableStatus, readyBefore, publishingStatus, leaseExpiredBefore, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultDialogRepositoryQueriesModel) SelectPublicUpdateOutboxClaimCandidates(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, readyBefore int64, publishingStatus int32, leaseExpiredBefore int64, limit int32) ([]DialogPublicUpdateOutboxRow, error) {
	var rList []DialogPublicUpdateOutboxRow
	query := "select outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message from dialog_public_update_outbox where (`status` in (?, ?) and next_retry_at > 0 and next_retry_at <= ?) or (`status` = ? and (lease_until = 0 or lease_until <= ?)) order by next_retry_at asc, outbox_id asc limit ? for update"

	err := m.db.QueryRowsPartial(ctx, &rList, query, pendingStatus, failedRetryableStatus, readyBefore, publishingStatus, leaseExpiredBefore, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}

func (m *defaultDialogRepositoryQueriesTxModel) SelectPublicUpdateOutboxClaimCandidates(pendingStatus int32, failedRetryableStatus int32, readyBefore int64, publishingStatus int32, leaseExpiredBefore int64, limit int32) ([]DialogPublicUpdateOutboxRow, error) {
	var rList []DialogPublicUpdateOutboxRow
	query := "select outbox_id, source_user_id, source_perm_auth_key_id, target_user_id, target_auth_policy, operation_id, delivery_path, public_update_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_pts, published_pts_count, published_seq, published_date, last_error_kind, last_error_message from dialog_public_update_outbox where (`status` in (?, ?) and next_retry_at > 0 and next_retry_at <= ?) or (`status` = ? and (lease_until = 0 or lease_until <= ?)) order by next_retry_at asc, outbox_id asc limit ? for update"

	err := m.tx.QueryRowsPartial(&rList, query, pendingStatus, failedRetryableStatus, readyBefore, publishingStatus, leaseExpiredBefore, limit)
	if err != nil {
		return nil, err
	}
	return rList, nil
}
