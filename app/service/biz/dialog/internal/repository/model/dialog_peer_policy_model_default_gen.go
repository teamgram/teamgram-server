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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	dialogPeerPolicyFieldNames          = builder.RawFieldNames(&DialogPeerPolicy{})
	dialogPeerPolicyRows                = strings.Join(dialogPeerPolicyFieldNames, ",")
	dialogPeerPolicyRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogPeerPolicyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogPeerPolicyRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogPeerPolicyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogPeerPolicyModel interface {
		Insert2(ctx context.Context, data *DialogPeerPolicy) (sql.Result, error)
	}

	defaultDialogPeerPolicyModel struct {
		db *sqlx.DB
	}

	DialogPeerPolicy struct {
		ScopeType     string `db:"scope_type" json:"scope_type"`
		ScopeId       string `db:"scope_id" json:"scope_id"`
		PeerType      int32  `db:"peer_type" json:"peer_type"`
		PeerId        int64  `db:"peer_id" json:"peer_id"`
		TtlPeriod     int32  `db:"ttl_period" json:"ttl_period"`
		ThemeEmoticon string `db:"theme_emoticon" json:"theme_emoticon"`
		PolicyVersion int64  `db:"policy_version" json:"policy_version"`
	}
)

func newDialogPeerPolicyModel(db *sqlx.DB) *defaultDialogPeerPolicyModel {
	return &defaultDialogPeerPolicyModel{
		db: db,
	}
}

func (m *defaultDialogPeerPolicyModel) Insert2(ctx context.Context, data *DialogPeerPolicy) (sql.Result, error) {
	tableName := "dialog_peer_policy"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?)", tableName, dialogPeerPolicyRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.ScopeType, data.ScopeId, data.PeerType, data.PeerId, data.TtlPeriod, data.ThemeEmoticon, data.PolicyVersion)
	if err != nil {
		return nil, fmt.Errorf("dialog_peer_policy.Insert2 exec: %w", err)
	}

	return r, nil
}
