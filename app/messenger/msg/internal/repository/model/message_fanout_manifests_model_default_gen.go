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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	messageFanoutManifestsFieldNames          = builder.RawFieldNames(&MessageFanoutManifests{})
	messageFanoutManifestsRows                = strings.Join(messageFanoutManifestsFieldNames, ",")
	messageFanoutManifestsRowsExpectAutoSet   = strings.Join(stringx.Remove(messageFanoutManifestsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messageFanoutManifestsRowsWithPlaceHolder = strings.Join(stringx.Remove(messageFanoutManifestsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	messageFanoutManifestsModel interface {
		Insert2(ctx context.Context, data *MessageFanoutManifests) (sql.Result, error)
		FindOne(ctx context.Context, manifestId int64) (*MessageFanoutManifests, error)
		FindListByManifestIdList(ctx context.Context, manifestId ...int64) ([]MessageFanoutManifests, error)
		Update2(ctx context.Context, data *MessageFanoutManifests) error
		Delete2(ctx context.Context, manifestId int64) error

		FindOneByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (*MessageFanoutManifests, error)
		FindListByCanonicalMessageIdList(ctx context.Context, canonicalMessageId ...int64) ([]MessageFanoutManifests, error)
	}

	defaultMessageFanoutManifestsModel struct {
		db *sqlx.DB
	}

	MessageFanoutManifests struct {
		ManifestId         int64 `db:"manifest_id" json:"manifest_id"`
		CanonicalMessageId int64 `db:"canonical_message_id" json:"canonical_message_id"`
		PeerType           int32 `db:"peer_type" json:"peer_type"`
		PeerId             int64 `db:"peer_id" json:"peer_id"`
		PeerSeq            int64 `db:"peer_seq" json:"peer_seq"`
		ActorUserId        int64 `db:"actor_user_id" json:"actor_user_id"`
		AffectedUserCount  int32 `db:"affected_user_count" json:"affected_user_count"`
		Status             int32 `db:"status" json:"status"`
		CompletedAt        int64 `db:"completed_at" json:"completed_at"`
	}
)

func newMessageFanoutManifestsModel(db *sqlx.DB) *defaultMessageFanoutManifestsModel {
	return &defaultMessageFanoutManifestsModel{
		db: db,
	}
}

func (m *defaultMessageFanoutManifestsModel) Insert2(ctx context.Context, data *MessageFanoutManifests) (sql.Result, error) {
	tableName := "message_fanout_manifests"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", tableName, messageFanoutManifestsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.CanonicalMessageId, data.PeerType, data.PeerId, data.PeerSeq, data.ActorUserId, data.AffectedUserCount, data.Status, data.CompletedAt)
	if err != nil {
		return nil, fmt.Errorf("message_fanout_manifests.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultMessageFanoutManifestsModel) Delete2(ctx context.Context, manifestId int64) error {
	tableName := "message_fanout_manifests"
	query := fmt.Sprintf("delete from `%s` where `manifest_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, manifestId)
	if err != nil {
		return fmt.Errorf("message_fanout_manifests.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessageFanoutManifestsModel) FindOne(ctx context.Context, manifestId int64) (*MessageFanoutManifests, error) {
	tableName := "message_fanout_manifests"
	query := fmt.Sprintf("select %s from %s where manifest_id = ? limit 1", messageFanoutManifestsRows, tableName)
	var resp MessageFanoutManifests

	err := m.db.QueryRowPartial(ctx, &resp, query, manifestId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_manifests",
				Key:      fmt.Sprintf("manifest_id=%v", manifestId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_fanout_manifests.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessageFanoutManifestsModel) FindListByManifestIdList(ctx context.Context, manifestId ...int64) ([]MessageFanoutManifests, error) {
	if len(manifestId) == 0 {
		return []MessageFanoutManifests{}, nil
	}
	tableName := "message_fanout_manifests"

	query := fmt.Sprintf("select %s from %s where manifest_id in (%s)", messageFanoutManifestsRows, tableName, sqlx.InInt64List(manifestId))

	var resp []MessageFanoutManifests
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []MessageFanoutManifests{}, nil
		}
		return nil, fmt.Errorf("message_fanout_manifests.FindListByManifestIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultMessageFanoutManifestsModel) Update2(ctx context.Context, data *MessageFanoutManifests) error {
	tableName := "message_fanout_manifests"
	query := fmt.Sprintf("update `%s` set %s where `manifest_id` = ?", tableName, messageFanoutManifestsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.CanonicalMessageId, data.PeerType, data.PeerId, data.PeerSeq, data.ActorUserId, data.AffectedUserCount, data.Status, data.CompletedAt, data.ManifestId)
	if err != nil {
		return fmt.Errorf("message_fanout_manifests.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessageFanoutManifestsModel) FindOneByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (*MessageFanoutManifests, error) {
	tableName := "message_fanout_manifests"
	query := fmt.Sprintf("select %s from %s where canonical_message_id = ? limit 1", messageFanoutManifestsRows, tableName)
	var resp MessageFanoutManifests

	err := m.db.QueryRowPartial(ctx, &resp, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_manifests",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_fanout_manifests.FindOneByCanonicalMessageId: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessageFanoutManifestsModel) FindListByCanonicalMessageIdList(ctx context.Context, canonicalMessageId ...int64) ([]MessageFanoutManifests, error) {
	if len(canonicalMessageId) == 0 {
		return []MessageFanoutManifests{}, nil
	}
	tableName := "message_fanout_manifests"

	query := fmt.Sprintf("select %s from %s where canonical_message_id in (%s)", messageFanoutManifestsRows, tableName, sqlx.InInt64List(canonicalMessageId))

	var resp []MessageFanoutManifests
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []MessageFanoutManifests{}, nil
		}
		return nil, fmt.Errorf("message_fanout_manifests.FindListByCanonicalMessageIdList: %w", err)
	}

	return resp, nil
}
