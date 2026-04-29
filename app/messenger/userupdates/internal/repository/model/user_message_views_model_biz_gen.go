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

type (
	bizUserMessageViewsModel interface {
		InsertOrUpdate(ctx context.Context, data *UserMessageViews) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UserMessageViews) (lastInsertId, rowsAffected int64, err error)

		SelectByUserCanonical(ctx context.Context, userId int64, canonicalMessageId int64) (*UserMessageViews, error)

		SelectByUserPeerSeq(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64) (*UserMessageViews, error)

		SelectPeerSeqRange(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64, maxPeerSeq int64, limit int32) ([]UserMessageViews, error)
		SelectPeerSeqRangeWithCB(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64, maxPeerSeq int64, limit int32, cb func(sz, i int, v *UserMessageViews)) ([]UserMessageViews, error)
	}
)

// InsertOrUpdate
// insert into user_message_views(user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at) values (:user_id, :peer_type, :peer_id, :peer_seq, :canonical_message_id, :from_user_id, :outgoing, :message_kind, :message_status, :edit_version, :date, :edit_date, :deleted_at, :view_schema_version, :view_payload, NOW(6), NOW(6)) on duplicate key update message_status = values(message_status), edit_version = values(edit_version), edit_date = values(edit_date), deleted_at = values(deleted_at), view_schema_version = values(view_schema_version), view_payload = values(view_payload), updated_at = NOW(6)
func (m *defaultUserMessageViewsModel) InsertOrUpdate(ctx context.Context, data *UserMessageViews) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_message_views(user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at) values (:user_id, :peer_type, :peer_id, :peer_seq, :canonical_message_id, :from_user_id, :outgoing, :message_kind, :message_status, :edit_version, :date, :edit_date, :deleted_at, :view_schema_version, :view_payload, NOW(6), NOW(6)) on duplicate key update message_status = values(message_status), edit_version = values(edit_version), edit_date = values(edit_date), deleted_at = values(deleted_at), view_schema_version = values(view_schema_version), view_payload = values(view_payload), updated_at = NOW(6)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_message_views.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_message_views.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_message_views.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into user_message_views(user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at) values (:user_id, :peer_type, :peer_id, :peer_seq, :canonical_message_id, :from_user_id, :outgoing, :message_kind, :message_status, :edit_version, :date, :edit_date, :deleted_at, :view_schema_version, :view_payload, NOW(6), NOW(6)) on duplicate key update message_status = values(message_status), edit_version = values(edit_version), edit_date = values(edit_date), deleted_at = values(deleted_at), view_schema_version = values(view_schema_version), view_payload = values(view_payload), updated_at = NOW(6)
func (m *defaultUserMessageViewsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UserMessageViews) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_message_views(user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at) values (:user_id, :peer_type, :peer_id, :peer_seq, :canonical_message_id, :from_user_id, :outgoing, :message_kind, :message_status, :edit_version, :date, :edit_date, :deleted_at, :view_schema_version, :view_payload, NOW(6), NOW(6)) on duplicate key update message_status = values(message_status), edit_version = values(edit_version), edit_date = values(edit_date), deleted_at = values(deleted_at), view_schema_version = values(view_schema_version), view_payload = values(view_payload), updated_at = NOW(6)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_message_views.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_message_views.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_message_views.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}

// SelectByUserCanonical
// select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at from user_message_views where user_id = :user_id and canonical_message_id = :canonical_message_id limit 1
func (m *defaultUserMessageViewsModel) SelectByUserCanonical(ctx context.Context, userId int64, canonicalMessageId int64) (rValue *UserMessageViews, err error) {

	var (
		query = "select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at from user_message_views where user_id = ? and canonical_message_id = ? limit 1"
		do    = &UserMessageViews{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_message_views",
				Key:      fmt.Sprintf("user_id=%v,canonical_message_id=%v", userId, canonicalMessageId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_message_views.SelectByUserCanonical: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserPeerSeq
// select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at from user_message_views where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and peer_seq = :peer_seq limit 1
func (m *defaultUserMessageViewsModel) SelectByUserPeerSeq(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64) (rValue *UserMessageViews, err error) {

	var (
		query = "select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at from user_message_views where user_id = ? and peer_type = ? and peer_id = ? and peer_seq = ? limit 1"
		do    = &UserMessageViews{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId, peerSeq)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_message_views",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v,peer_seq=%v", userId, peerType, peerId, peerSeq),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_message_views.SelectByUserPeerSeq: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectPeerSeqRange
// select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at from user_message_views where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and peer_seq > :peer_seq and peer_seq <= :maxPeerSeq order by peer_seq asc limit :limit
func (m *defaultUserMessageViewsModel) SelectPeerSeqRange(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64, maxPeerSeq int64, limit int32) (rList []UserMessageViews, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at from user_message_views where user_id = ? and peer_type = ? and peer_id = ? and peer_seq > ? and peer_seq <= ? order by peer_seq asc limit ?"
		values []UserMessageViews
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, peerType, peerId, peerSeq, maxPeerSeq, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserMessageViews{}
			err = nil
			return
		}
		err = fmt.Errorf("user_message_views.SelectPeerSeqRange: %w", err)
		return
	}

	rList = values

	return
}

// SelectPeerSeqRangeWithCB
// select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at from user_message_views where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and peer_seq > :peer_seq and peer_seq <= :maxPeerSeq order by peer_seq asc limit :limit
func (m *defaultUserMessageViewsModel) SelectPeerSeqRangeWithCB(ctx context.Context, userId int64, peerType int32, peerId int64, peerSeq int64, maxPeerSeq int64, limit int32, cb func(sz, i int, v *UserMessageViews)) (rList []UserMessageViews, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at from user_message_views where user_id = ? and peer_type = ? and peer_id = ? and peer_seq > ? and peer_seq <= ? order by peer_seq asc limit ?"
		values []UserMessageViews
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, peerType, peerId, peerSeq, maxPeerSeq, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserMessageViews{}
			err = nil
			return
		}
		err = fmt.Errorf("user_message_views.SelectPeerSeqRangeWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}
