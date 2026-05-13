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

type bizChatCreateOperationsModel interface {
	Insert(ctx context.Context, data *ChatCreateOperations) (lastInsertId, rowsAffected int64, err error)
	SelectByReplayKey(ctx context.Context, replayKey string) (*ChatCreateOperations, error)
	SelectByReplayKeyForUpdate(ctx context.Context, replayKey string) (*ChatCreateOperations, error)
	SelectLastCompletedByActor(ctx context.Context, actorUserId int64, status int32) (*ChatCreateOperations, error)
	EnsureActorLock(ctx context.Context, data *ChatCreateOperations) (lastInsertId, rowsAffected int64, err error)
	SelectActorLockForUpdate(ctx context.Context, operationId string) (*ChatCreateOperations, error)
	MarkChatCreated(ctx context.Context, chatId int64, participantsVersion int32, status int32, updatedAtSec int64, operationId string) (rowsAffected int64, err error)
	ResetForRetry(ctx context.Context, operationId string, actorUserId int64, clientMsgId int64, title string, inviteeIds string, ttlPeriod int32, participantsVersion int32, status int32, date int64, updatedAtSec int64, expiresAt int64, replayKey string) (rowsAffected int64, err error)
}

type ChatCreateOperationsTxModel interface {
	Insert(data *ChatCreateOperations) (lastInsertId, rowsAffected int64, err error)
	SelectByReplayKey(replayKey string) (*ChatCreateOperations, error)
	SelectByReplayKeyForUpdate(replayKey string) (*ChatCreateOperations, error)
	SelectLastCompletedByActor(actorUserId int64, status int32) (*ChatCreateOperations, error)
	EnsureActorLock(data *ChatCreateOperations) (lastInsertId, rowsAffected int64, err error)
	SelectActorLockForUpdate(operationId string) (*ChatCreateOperations, error)
	MarkChatCreated(chatId int64, participantsVersion int32, status int32, updatedAtSec int64, operationId string) (rowsAffected int64, err error)
	ResetForRetry(operationId string, actorUserId int64, clientMsgId int64, title string, inviteeIds string, ttlPeriod int32, participantsVersion int32, status int32, date int64, updatedAtSec int64, expiresAt int64, replayKey string) (rowsAffected int64, err error)
}

type defaultChatCreateOperationsTxModel struct {
	tx *sqlx.Tx
}

func NewChatCreateOperationsTxModel(tx *sqlx.Tx) ChatCreateOperationsTxModel {
	return &defaultChatCreateOperationsTxModel{tx: tx}
}

// Insert
// insert into chat_create_operations(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at) values (:operation_id, :replay_key, :actor_user_id, :client_msg_id, :title, :invitee_ids, :ttl_period, :chat_id, :participants_version, :status, :date, :updated_at_sec, :expires_at)
func (m *defaultChatCreateOperationsModel) Insert(ctx context.Context, data *ChatCreateOperations) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_create_operations(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at) values (:operation_id, :replay_key, :actor_user_id, :client_msg_id, :title, :invitee_ids, :ttl_period, :chat_id, :participants_version, :status, :date, :updated_at_sec, :expires_at)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("chat_create_operations.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into chat_create_operations(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at) values (:operation_id, :replay_key, :actor_user_id, :client_msg_id, :title, :invitee_ids, :ttl_period, :chat_id, :participants_version, :status, :date, :updated_at_sec, :expires_at)
func (m *defaultChatCreateOperationsTxModel) Insert(data *ChatCreateOperations) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_create_operations(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at) values (:operation_id, :replay_key, :actor_user_id, :client_msg_id, :title, :invitee_ids, :ttl_period, :chat_id, :participants_version, :status, :date, :updated_at_sec, :expires_at)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("chat_create_operations.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.Insert rows affected: %w", err)
	}

	return
}

// SelectByReplayKey
// select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where replay_key = :replay_key
func (m *defaultChatCreateOperationsModel) SelectByReplayKey(ctx context.Context, replayKey string) (rValue *ChatCreateOperations, err error) {

	var (
		query = "select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where replay_key = ?"
		do    = &ChatCreateOperations{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, replayKey)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("replay_key=%v", replayKey),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chat_create_operations.SelectByReplayKey: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByReplayKey
// select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where replay_key = :replay_key
func (m *defaultChatCreateOperationsTxModel) SelectByReplayKey(replayKey string) (rValue *ChatCreateOperations, err error) {
	var (
		query = "select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where replay_key = ?"
		do    = &ChatCreateOperations{}
	)
	err = m.tx.QueryRowPartial(do, query, replayKey)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("replay_key=%v", replayKey),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chat_create_operations.SelectByReplayKey: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByReplayKeyForUpdate
// select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where replay_key = :replay_key for update
func (m *defaultChatCreateOperationsModel) SelectByReplayKeyForUpdate(ctx context.Context, replayKey string) (rValue *ChatCreateOperations, err error) {

	var (
		query = "select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where replay_key = ? for update"
		do    = &ChatCreateOperations{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, replayKey)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("replay_key=%v", replayKey),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chat_create_operations.SelectByReplayKeyForUpdate: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByReplayKeyForUpdate
// select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where replay_key = :replay_key for update
func (m *defaultChatCreateOperationsTxModel) SelectByReplayKeyForUpdate(replayKey string) (rValue *ChatCreateOperations, err error) {
	var (
		query = "select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where replay_key = ? for update"
		do    = &ChatCreateOperations{}
	)
	err = m.tx.QueryRowPartial(do, query, replayKey)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("replay_key=%v", replayKey),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chat_create_operations.SelectByReplayKeyForUpdate: %w", err)
		return
	}
	rValue = do

	return
}

// SelectLastCompletedByActor
// select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where actor_user_id = :actor_user_id and `status` = :status order by updated_at_sec desc limit 1
func (m *defaultChatCreateOperationsModel) SelectLastCompletedByActor(ctx context.Context, actorUserId int64, status int32) (rValue *ChatCreateOperations, err error) {

	var (
		query = "select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where actor_user_id = ? and `status` = ? order by updated_at_sec desc limit 1"
		do    = &ChatCreateOperations{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, actorUserId, status)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("actor_user_id=%v,status=%v", actorUserId, status),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chat_create_operations.SelectLastCompletedByActor: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectLastCompletedByActor
// select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where actor_user_id = :actor_user_id and `status` = :status order by updated_at_sec desc limit 1
func (m *defaultChatCreateOperationsTxModel) SelectLastCompletedByActor(actorUserId int64, status int32) (rValue *ChatCreateOperations, err error) {
	var (
		query = "select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where actor_user_id = ? and `status` = ? order by updated_at_sec desc limit 1"
		do    = &ChatCreateOperations{}
	)
	err = m.tx.QueryRowPartial(do, query, actorUserId, status)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("actor_user_id=%v,status=%v", actorUserId, status),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chat_create_operations.SelectLastCompletedByActor: %w", err)
		return
	}
	rValue = do

	return
}

// EnsureActorLock
// insert ignore into chat_create_operations(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at) values (:operation_id, :replay_key, :actor_user_id, 0, ”, ”, 0, 0, 0, :status, :date, :updated_at_sec, :expires_at)
func (m *defaultChatCreateOperationsModel) EnsureActorLock(ctx context.Context, data *ChatCreateOperations) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into chat_create_operations(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at) values (:operation_id, :replay_key, :actor_user_id, 0, '', '', 0, 0, 0, :status, :date, :updated_at_sec, :expires_at)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("chat_create_operations.EnsureActorLock named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.EnsureActorLock last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.EnsureActorLock rows affected: %w", err)
	}

	return

}

// EnsureActorLock
// insert ignore into chat_create_operations(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at) values (:operation_id, :replay_key, :actor_user_id, 0, ”, ”, 0, 0, 0, :status, :date, :updated_at_sec, :expires_at)
func (m *defaultChatCreateOperationsTxModel) EnsureActorLock(data *ChatCreateOperations) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into chat_create_operations(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at) values (:operation_id, :replay_key, :actor_user_id, 0, '', '', 0, 0, 0, :status, :date, :updated_at_sec, :expires_at)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("chat_create_operations.EnsureActorLock named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.EnsureActorLock last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.EnsureActorLock rows affected: %w", err)
	}

	return
}

// SelectActorLockForUpdate
// select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where operation_id = :operation_id for update
func (m *defaultChatCreateOperationsModel) SelectActorLockForUpdate(ctx context.Context, operationId string) (rValue *ChatCreateOperations, err error) {

	var (
		query = "select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where operation_id = ? for update"
		do    = &ChatCreateOperations{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("operation_id=%v", operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chat_create_operations.SelectActorLockForUpdate: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectActorLockForUpdate
// select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where operation_id = :operation_id for update
func (m *defaultChatCreateOperationsTxModel) SelectActorLockForUpdate(operationId string) (rValue *ChatCreateOperations, err error) {
	var (
		query = "select id, operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, `status`, `date`, updated_at_sec, expires_at from chat_create_operations where operation_id = ? for update"
		do    = &ChatCreateOperations{}
	)
	err = m.tx.QueryRowPartial(do, query, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chat_create_operations",
				Key:      fmt.Sprintf("operation_id=%v", operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chat_create_operations.SelectActorLockForUpdate: %w", err)
		return
	}
	rValue = do

	return
}

// MarkChatCreated
// update chat_create_operations set chat_id = :chat_id, participants_version = :participants_version, `status` = :status, updated_at_sec = :updated_at_sec where operation_id = :operation_id
func (m *defaultChatCreateOperationsModel) MarkChatCreated(ctx context.Context, chatId int64, participantsVersion int32, status int32, updatedAtSec int64, operationId string) (rowsAffected int64, err error) {

	var (
		query   = "update chat_create_operations set chat_id = ?, participants_version = ?, `status` = ?, updated_at_sec = ? where operation_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, chatId, participantsVersion, status, updatedAtSec, operationId)

	if err != nil {
		err = fmt.Errorf("chat_create_operations.MarkChatCreated exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.MarkChatCreated rows affected: %w", err)
		return
	}

	return
}

// MarkChatCreated
// update chat_create_operations set chat_id = :chat_id, participants_version = :participants_version, `status` = :status, updated_at_sec = :updated_at_sec where operation_id = :operation_id
func (m *defaultChatCreateOperationsTxModel) MarkChatCreated(chatId int64, participantsVersion int32, status int32, updatedAtSec int64, operationId string) (rowsAffected int64, err error) {
	var (
		query   = "update chat_create_operations set chat_id = ?, participants_version = ?, `status` = ?, updated_at_sec = ? where operation_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, chatId, participantsVersion, status, updatedAtSec, operationId)

	if err != nil {
		err = fmt.Errorf("chat_create_operations.MarkChatCreated exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.MarkChatCreated rows affected: %w", err)
		return
	}

	return
}

// ResetForRetry
// update chat_create_operations set operation_id = :operation_id, actor_user_id = :actor_user_id, client_msg_id = :client_msg_id, title = :title, invitee_ids = :invitee_ids, ttl_period = :ttl_period, chat_id = 0, participants_version = :participants_version, `status` = :status, `date` = :date, updated_at_sec = :updated_at_sec, expires_at = :expires_at where replay_key = :replay_key
func (m *defaultChatCreateOperationsModel) ResetForRetry(ctx context.Context, operationId string, actorUserId int64, clientMsgId int64, title string, inviteeIds string, ttlPeriod int32, participantsVersion int32, status int32, date int64, updatedAtSec int64, expiresAt int64, replayKey string) (rowsAffected int64, err error) {

	var (
		query   = "update chat_create_operations set operation_id = ?, actor_user_id = ?, client_msg_id = ?, title = ?, invitee_ids = ?, ttl_period = ?, chat_id = 0, participants_version = ?, `status` = ?, `date` = ?, updated_at_sec = ?, expires_at = ? where replay_key = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, operationId, actorUserId, clientMsgId, title, inviteeIds, ttlPeriod, participantsVersion, status, date, updatedAtSec, expiresAt, replayKey)

	if err != nil {
		err = fmt.Errorf("chat_create_operations.ResetForRetry exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.ResetForRetry rows affected: %w", err)
		return
	}

	return
}

// ResetForRetry
// update chat_create_operations set operation_id = :operation_id, actor_user_id = :actor_user_id, client_msg_id = :client_msg_id, title = :title, invitee_ids = :invitee_ids, ttl_period = :ttl_period, chat_id = 0, participants_version = :participants_version, `status` = :status, `date` = :date, updated_at_sec = :updated_at_sec, expires_at = :expires_at where replay_key = :replay_key
func (m *defaultChatCreateOperationsTxModel) ResetForRetry(operationId string, actorUserId int64, clientMsgId int64, title string, inviteeIds string, ttlPeriod int32, participantsVersion int32, status int32, date int64, updatedAtSec int64, expiresAt int64, replayKey string) (rowsAffected int64, err error) {
	var (
		query   = "update chat_create_operations set operation_id = ?, actor_user_id = ?, client_msg_id = ?, title = ?, invitee_ids = ?, ttl_period = ?, chat_id = 0, participants_version = ?, `status` = ?, `date` = ?, updated_at_sec = ?, expires_at = ? where replay_key = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, operationId, actorUserId, clientMsgId, title, inviteeIds, ttlPeriod, participantsVersion, status, date, updatedAtSec, expiresAt, replayKey)

	if err != nil {
		err = fmt.Errorf("chat_create_operations.ResetForRetry exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_create_operations.ResetForRetry rows affected: %w", err)
		return
	}

	return
}
