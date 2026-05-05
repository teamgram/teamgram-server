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

type bizDialogPeerPolicyModel interface {
	Upsert(ctx context.Context, data *DialogPeerPolicy) (lastInsertId, rowsAffected int64, err error)
	SelectByScope(ctx context.Context, scopeType string, scopeId string) (*DialogPeerPolicy, error)
}

type DialogPeerPolicyTxModel interface {
	Upsert(data *DialogPeerPolicy) (lastInsertId, rowsAffected int64, err error)
	SelectByScope(scopeType string, scopeId string) (*DialogPeerPolicy, error)
}

type defaultDialogPeerPolicyTxModel struct {
	tx *sqlx.Tx
}

func NewDialogPeerPolicyTxModel(tx *sqlx.Tx) DialogPeerPolicyTxModel {
	return &defaultDialogPeerPolicyTxModel{tx: tx}
}

// Upsert
// insert into dialog_peer_policy(scope_type, scope_id, peer_type, peer_id, ttl_period, theme_emoticon, policy_version) values (:scope_type, :scope_id, :peer_type, :peer_id, :ttl_period, :theme_emoticon, :policy_version) on duplicate key update peer_type = values(peer_type), peer_id = values(peer_id), ttl_period = values(ttl_period), theme_emoticon = values(theme_emoticon), policy_version = policy_version + 1
func (m *defaultDialogPeerPolicyModel) Upsert(ctx context.Context, data *DialogPeerPolicy) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_peer_policy(scope_type, scope_id, peer_type, peer_id, ttl_period, theme_emoticon, policy_version) values (:scope_type, :scope_id, :peer_type, :peer_id, :ttl_period, :theme_emoticon, :policy_version) on duplicate key update peer_type = values(peer_type), peer_id = values(peer_id), ttl_period = values(ttl_period), theme_emoticon = values(theme_emoticon), policy_version = policy_version + 1"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_peer_policy.Upsert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_peer_policy.Upsert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_peer_policy.Upsert rows affected: %w", err)
	}

	return

}

// Upsert
// insert into dialog_peer_policy(scope_type, scope_id, peer_type, peer_id, ttl_period, theme_emoticon, policy_version) values (:scope_type, :scope_id, :peer_type, :peer_id, :ttl_period, :theme_emoticon, :policy_version) on duplicate key update peer_type = values(peer_type), peer_id = values(peer_id), ttl_period = values(ttl_period), theme_emoticon = values(theme_emoticon), policy_version = policy_version + 1
func (m *defaultDialogPeerPolicyTxModel) Upsert(data *DialogPeerPolicy) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_peer_policy(scope_type, scope_id, peer_type, peer_id, ttl_period, theme_emoticon, policy_version) values (:scope_type, :scope_id, :peer_type, :peer_id, :ttl_period, :theme_emoticon, :policy_version) on duplicate key update peer_type = values(peer_type), peer_id = values(peer_id), ttl_period = values(ttl_period), theme_emoticon = values(theme_emoticon), policy_version = policy_version + 1"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_peer_policy.Upsert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_peer_policy.Upsert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_peer_policy.Upsert rows affected: %w", err)
	}

	return
}

// SelectByScope
// select scope_type, scope_id, peer_type, peer_id, ttl_period, theme_emoticon, policy_version from dialog_peer_policy where scope_type = :scope_type and scope_id = :scope_id limit 1
func (m *defaultDialogPeerPolicyModel) SelectByScope(ctx context.Context, scopeType string, scopeId string) (rValue *DialogPeerPolicy, err error) {

	var (
		query = "select scope_type, scope_id, peer_type, peer_id, ttl_period, theme_emoticon, policy_version from dialog_peer_policy where scope_type = ? and scope_id = ? limit 1"
		do    = &DialogPeerPolicy{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, scopeType, scopeId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_peer_policy",
				Key:      fmt.Sprintf("scope_type=%v,scope_id=%v", scopeType, scopeId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_peer_policy.SelectByScope: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByScope
// select scope_type, scope_id, peer_type, peer_id, ttl_period, theme_emoticon, policy_version from dialog_peer_policy where scope_type = :scope_type and scope_id = :scope_id limit 1
func (m *defaultDialogPeerPolicyTxModel) SelectByScope(scopeType string, scopeId string) (rValue *DialogPeerPolicy, err error) {
	var (
		query = "select scope_type, scope_id, peer_type, peer_id, ttl_period, theme_emoticon, policy_version from dialog_peer_policy where scope_type = ? and scope_id = ? limit 1"
		do    = &DialogPeerPolicy{}
	)
	err = m.tx.QueryRowPartial(do, query, scopeType, scopeId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_peer_policy",
				Key:      fmt.Sprintf("scope_type=%v,scope_id=%v", scopeType, scopeId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_peer_policy.SelectByScope: %w", err)
		return
	}
	rValue = do

	return
}
