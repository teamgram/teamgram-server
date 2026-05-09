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

type bizCanonicalMessagesModel interface {
	Insert(ctx context.Context, data *CanonicalMessages) (lastInsertId, rowsAffected int64, err error)
	SelectByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (*CanonicalMessages, error)
	SelectByPeerSeq(ctx context.Context, peerType int32, peerId int64, peerSeq int64) (*CanonicalMessages, error)
	UpdateMessageEdit(ctx context.Context, messageText string, editVersion int32, editDate int64, canonicalMessageId int64, oldEditVersion int32) (rowsAffected int64, err error)
}

type CanonicalMessagesTxModel interface {
	Insert(data *CanonicalMessages) (lastInsertId, rowsAffected int64, err error)
	SelectByCanonicalMessageId(canonicalMessageId int64) (*CanonicalMessages, error)
	SelectByPeerSeq(peerType int32, peerId int64, peerSeq int64) (*CanonicalMessages, error)
	UpdateMessageEdit(messageText string, editVersion int32, editDate int64, canonicalMessageId int64, oldEditVersion int32) (rowsAffected int64, err error)
}

type defaultCanonicalMessagesTxModel struct {
	tx *sqlx.Tx
}

func NewCanonicalMessagesTxModel(tx *sqlx.Tx) CanonicalMessagesTxModel {
	return &defaultCanonicalMessagesTxModel{tx: tx}
}

// Insert
// insert into canonical_messages(canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version) values (:canonical_message_id, :peer_type, :peer_id, :peer_seq, :from_user_id, :message_kind, :message_text, :entities_payload_schema_version, :entities_payload, :media_ref_schema_version, :media_ref_payload, :message_attrs_schema_version, :message_attrs_payload, :forward_ref_schema_version, :forward_ref_payload, :service_action_schema_version, :service_action_payload, :message_status, :edit_version, :date, :storage_schema_version)
func (m *defaultCanonicalMessagesModel) Insert(ctx context.Context, data *CanonicalMessages) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into canonical_messages(canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version) values (:canonical_message_id, :peer_type, :peer_id, :peer_seq, :from_user_id, :message_kind, :message_text, :entities_payload_schema_version, :entities_payload, :media_ref_schema_version, :media_ref_payload, :message_attrs_schema_version, :message_attrs_payload, :forward_ref_schema_version, :forward_ref_payload, :service_action_schema_version, :service_action_payload, :message_status, :edit_version, :date, :storage_schema_version)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("canonical_messages.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("canonical_messages.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("canonical_messages.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into canonical_messages(canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version) values (:canonical_message_id, :peer_type, :peer_id, :peer_seq, :from_user_id, :message_kind, :message_text, :entities_payload_schema_version, :entities_payload, :media_ref_schema_version, :media_ref_payload, :message_attrs_schema_version, :message_attrs_payload, :forward_ref_schema_version, :forward_ref_payload, :service_action_schema_version, :service_action_payload, :message_status, :edit_version, :date, :storage_schema_version)
func (m *defaultCanonicalMessagesTxModel) Insert(data *CanonicalMessages) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into canonical_messages(canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version) values (:canonical_message_id, :peer_type, :peer_id, :peer_seq, :from_user_id, :message_kind, :message_text, :entities_payload_schema_version, :entities_payload, :media_ref_schema_version, :media_ref_payload, :message_attrs_schema_version, :message_attrs_payload, :forward_ref_schema_version, :forward_ref_payload, :service_action_schema_version, :service_action_payload, :message_status, :edit_version, :date, :storage_schema_version)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("canonical_messages.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("canonical_messages.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("canonical_messages.Insert rows affected: %w", err)
	}

	return
}

// SelectByCanonicalMessageId
// select canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version from canonical_messages where canonical_message_id = :canonical_message_id limit 1
func (m *defaultCanonicalMessagesModel) SelectByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (rValue *CanonicalMessages, err error) {

	var (
		query = "select canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version from canonical_messages where canonical_message_id = ? limit 1"
		do    = &CanonicalMessages{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "canonical_messages",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("canonical_messages.SelectByCanonicalMessageId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByCanonicalMessageId
// select canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version from canonical_messages where canonical_message_id = :canonical_message_id limit 1
func (m *defaultCanonicalMessagesTxModel) SelectByCanonicalMessageId(canonicalMessageId int64) (rValue *CanonicalMessages, err error) {
	var (
		query = "select canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version from canonical_messages where canonical_message_id = ? limit 1"
		do    = &CanonicalMessages{}
	)
	err = m.tx.QueryRowPartial(do, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "canonical_messages",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("canonical_messages.SelectByCanonicalMessageId: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByPeerSeq
// select canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version from canonical_messages where peer_type = :peer_type and peer_id = :peer_id and peer_seq = :peer_seq limit 1
func (m *defaultCanonicalMessagesModel) SelectByPeerSeq(ctx context.Context, peerType int32, peerId int64, peerSeq int64) (rValue *CanonicalMessages, err error) {

	var (
		query = "select canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version from canonical_messages where peer_type = ? and peer_id = ? and peer_seq = ? limit 1"
		do    = &CanonicalMessages{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, peerType, peerId, peerSeq)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "canonical_messages",
				Key:      fmt.Sprintf("peer_type=%v,peer_id=%v,peer_seq=%v", peerType, peerId, peerSeq),
				Cause:    err,
			}
		}
		err = fmt.Errorf("canonical_messages.SelectByPeerSeq: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByPeerSeq
// select canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version from canonical_messages where peer_type = :peer_type and peer_id = :peer_id and peer_seq = :peer_seq limit 1
func (m *defaultCanonicalMessagesTxModel) SelectByPeerSeq(peerType int32, peerId int64, peerSeq int64) (rValue *CanonicalMessages, err error) {
	var (
		query = "select canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text, entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload, message_attrs_schema_version, message_attrs_payload, forward_ref_schema_version, forward_ref_payload, service_action_schema_version, service_action_payload, message_status, edit_version, `date`, storage_schema_version from canonical_messages where peer_type = ? and peer_id = ? and peer_seq = ? limit 1"
		do    = &CanonicalMessages{}
	)
	err = m.tx.QueryRowPartial(do, query, peerType, peerId, peerSeq)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "canonical_messages",
				Key:      fmt.Sprintf("peer_type=%v,peer_id=%v,peer_seq=%v", peerType, peerId, peerSeq),
				Cause:    err,
			}
		}
		err = fmt.Errorf("canonical_messages.SelectByPeerSeq: %w", err)
		return
	}
	rValue = do

	return
}

// UpdateMessageEdit
// update canonical_messages set message_text = :message_text, edit_version = :edit_version, edit_date = :edit_date where canonical_message_id = :canonical_message_id and edit_version = :old_edit_version
func (m *defaultCanonicalMessagesModel) UpdateMessageEdit(ctx context.Context, messageText string, editVersion int32, editDate int64, canonicalMessageId int64, oldEditVersion int32) (rowsAffected int64, err error) {

	var (
		query   = "update canonical_messages set message_text = ?, edit_version = ?, edit_date = ? where canonical_message_id = ? and edit_version = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, messageText, editVersion, editDate, canonicalMessageId, oldEditVersion)

	if err != nil {
		err = fmt.Errorf("canonical_messages.UpdateMessageEdit exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("canonical_messages.UpdateMessageEdit rows affected: %w", err)
		return
	}

	return
}

// UpdateMessageEdit
// update canonical_messages set message_text = :message_text, edit_version = :edit_version, edit_date = :edit_date where canonical_message_id = :canonical_message_id and edit_version = :old_edit_version
func (m *defaultCanonicalMessagesTxModel) UpdateMessageEdit(messageText string, editVersion int32, editDate int64, canonicalMessageId int64, oldEditVersion int32) (rowsAffected int64, err error) {
	var (
		query   = "update canonical_messages set message_text = ?, edit_version = ?, edit_date = ? where canonical_message_id = ? and edit_version = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, messageText, editVersion, editDate, canonicalMessageId, oldEditVersion)

	if err != nil {
		err = fmt.Errorf("canonical_messages.UpdateMessageEdit exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("canonical_messages.UpdateMessageEdit rows affected: %w", err)
		return
	}

	return
}
