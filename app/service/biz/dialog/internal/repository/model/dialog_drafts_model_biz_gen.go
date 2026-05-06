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
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizDialogDraftsModel interface {
	InsertOrUpdate(ctx context.Context, data *DialogDrafts) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (*DialogDrafts, error)
	SelectActiveByUser(ctx context.Context, userId int64) ([]DialogDrafts, error)
	SelectActiveByUserWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *DialogDrafts)) ([]DialogDrafts, error)
	ClearByUserPeerBeforeDate(ctx context.Context, entitiesPayload []byte, draftPayloadSchemaVersion int32, draftPayload []byte, date time.Time, userId int64, peerDialogId int64, clearBeforeDate string) (rowsAffected int64, err error)
}

type DialogDraftsTxModel interface {
	InsertOrUpdate(data *DialogDrafts) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(userId int64, peerType int32, peerId int64) (*DialogDrafts, error)
	SelectActiveByUser(userId int64) ([]DialogDrafts, error)
	ClearByUserPeerBeforeDate(entitiesPayload []byte, draftPayloadSchemaVersion int32, draftPayload []byte, date time.Time, userId int64, peerDialogId int64, clearBeforeDate string) (rowsAffected int64, err error)
}

type defaultDialogDraftsTxModel struct {
	tx *sqlx.Tx
}

func NewDialogDraftsTxModel(tx *sqlx.Tx) DialogDraftsTxModel {
	return &defaultDialogDraftsTxModel{tx: tx}
}

// InsertOrUpdate
// insert into dialog_drafts(user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date`) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :draft_kind, :message, :entities_payload, :reply_to_peer_seq, :draft_payload_schema_version, :draft_payload, :date) on duplicate key update draft_kind = values(draft_kind), message = values(message), entities_payload = values(entities_payload), reply_to_peer_seq = values(reply_to_peer_seq), draft_payload_schema_version = values(draft_payload_schema_version), draft_payload = values(draft_payload), `date` = values(`date`)
func (m *defaultDialogDraftsModel) InsertOrUpdate(ctx context.Context, data *DialogDrafts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_drafts(user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date`) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :draft_kind, :message, :entities_payload, :reply_to_peer_seq, :draft_payload_schema_version, :draft_payload, :date) on duplicate key update draft_kind = values(draft_kind), message = values(message), entities_payload = values(entities_payload), reply_to_peer_seq = values(reply_to_peer_seq), draft_payload_schema_version = values(draft_payload_schema_version), draft_payload = values(draft_payload), `date` = values(`date`)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_drafts.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_drafts.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_drafts.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into dialog_drafts(user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date`) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :draft_kind, :message, :entities_payload, :reply_to_peer_seq, :draft_payload_schema_version, :draft_payload, :date) on duplicate key update draft_kind = values(draft_kind), message = values(message), entities_payload = values(entities_payload), reply_to_peer_seq = values(reply_to_peer_seq), draft_payload_schema_version = values(draft_payload_schema_version), draft_payload = values(draft_payload), `date` = values(`date`)
func (m *defaultDialogDraftsTxModel) InsertOrUpdate(data *DialogDrafts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_drafts(user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date`) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :draft_kind, :message, :entities_payload, :reply_to_peer_seq, :draft_payload_schema_version, :draft_payload, :date) on duplicate key update draft_kind = values(draft_kind), message = values(message), entities_payload = values(entities_payload), reply_to_peer_seq = values(reply_to_peer_seq), draft_payload_schema_version = values(draft_payload_schema_version), draft_payload = values(draft_payload), `date` = values(`date`)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_drafts.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_drafts.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_drafts.InsertOrUpdate rows affected: %w", err)
	}

	return
}

// SelectByUserPeer
// select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultDialogDraftsModel) SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *DialogDrafts, err error) {

	var (
		query = "select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = ? and peer_type = ? and peer_id = ? limit 1"
		do    = &DialogDrafts{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_drafts",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_drafts.SelectByUserPeer: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserPeer
// select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultDialogDraftsTxModel) SelectByUserPeer(userId int64, peerType int32, peerId int64) (rValue *DialogDrafts, err error) {
	var (
		query = "select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = ? and peer_type = ? and peer_id = ? limit 1"
		do    = &DialogDrafts{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_drafts",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_drafts.SelectByUserPeer: %w", err)
		return
	}
	rValue = do

	return
}

// SelectActiveByUser
// select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = :user_id and (draft_kind != 0 or message != ”)
func (m *defaultDialogDraftsModel) SelectActiveByUser(ctx context.Context, userId int64) (rList []DialogDrafts, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = ? and (draft_kind != 0 or message != '')"
		values []DialogDrafts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogDrafts{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_drafts.SelectActiveByUser: %w", err)
		return
	}

	rList = values

	return
}

// SelectActiveByUser
// select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = :user_id and (draft_kind != 0 or message != ”)
func (m *defaultDialogDraftsTxModel) SelectActiveByUser(userId int64) (rList []DialogDrafts, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = ? and (draft_kind != 0 or message != '')"
		values []DialogDrafts
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogDrafts{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_drafts.SelectActiveByUser: %w", err)
		return
	}

	rList = values

	return
}

// SelectActiveByUserWithCB
// select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = :user_id and (draft_kind != 0 or message != ”)
func (m *defaultDialogDraftsModel) SelectActiveByUserWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *DialogDrafts)) (rList []DialogDrafts, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, draft_kind, message, entities_payload, reply_to_peer_seq, draft_payload_schema_version, draft_payload, `date` from dialog_drafts where user_id = ? and (draft_kind != 0 or message != '')"
		values []DialogDrafts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogDrafts{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_drafts.SelectActiveByUserWithCB: %w", err)
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

// ClearByUserPeerBeforeDate
// update dialog_drafts set draft_kind = 0, message = ”, entities_payload = :entities_payload, reply_to_peer_seq = 0, draft_payload_schema_version = :draft_payload_schema_version, draft_payload = :draft_payload, `date` = :date where user_id = :user_id and peer_dialog_id = :peer_dialog_id and `date` <= :clear_before_date
func (m *defaultDialogDraftsModel) ClearByUserPeerBeforeDate(ctx context.Context, entitiesPayload []byte, draftPayloadSchemaVersion int32, draftPayload []byte, date time.Time, userId int64, peerDialogId int64, clearBeforeDate string) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_drafts set draft_kind = 0, message = '', entities_payload = ?, reply_to_peer_seq = 0, draft_payload_schema_version = ?, draft_payload = ?, `date` = ? where user_id = ? and peer_dialog_id = ? and `date` <= ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, entitiesPayload, draftPayloadSchemaVersion, draftPayload, date, userId, peerDialogId, clearBeforeDate)

	if err != nil {
		err = fmt.Errorf("dialog_drafts.ClearByUserPeerBeforeDate exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_drafts.ClearByUserPeerBeforeDate rows affected: %w", err)
		return
	}

	return
}

// ClearByUserPeerBeforeDate
// update dialog_drafts set draft_kind = 0, message = ”, entities_payload = :entities_payload, reply_to_peer_seq = 0, draft_payload_schema_version = :draft_payload_schema_version, draft_payload = :draft_payload, `date` = :date where user_id = :user_id and peer_dialog_id = :peer_dialog_id and `date` <= :clear_before_date
func (m *defaultDialogDraftsTxModel) ClearByUserPeerBeforeDate(entitiesPayload []byte, draftPayloadSchemaVersion int32, draftPayload []byte, date time.Time, userId int64, peerDialogId int64, clearBeforeDate string) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_drafts set draft_kind = 0, message = '', entities_payload = ?, reply_to_peer_seq = 0, draft_payload_schema_version = ?, draft_payload = ?, `date` = ? where user_id = ? and peer_dialog_id = ? and `date` <= ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, entitiesPayload, draftPayloadSchemaVersion, draftPayload, date, userId, peerDialogId, clearBeforeDate)

	if err != nil {
		err = fmt.Errorf("dialog_drafts.ClearByUserPeerBeforeDate exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_drafts.ClearByUserPeerBeforeDate rows affected: %w", err)
		return
	}

	return
}
