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
	userPeerBlocksFieldNames          = builder.RawFieldNames(&UserPeerBlocks{})
	userPeerBlocksRows                = strings.Join(userPeerBlocksFieldNames, ",")
	userPeerBlocksRowsExpectAutoSet   = strings.Join(stringx.Remove(userPeerBlocksFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPeerBlocksRowsWithPlaceHolder = strings.Join(stringx.Remove(userPeerBlocksFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userPeerBlocksModel interface {
		Insert2(ctx context.Context, data *UserPeerBlocks) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserPeerBlocks, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserPeerBlocks, error)
		Update2(ctx context.Context, data *UserPeerBlocks) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserPeerBlocks, error)
	}

	defaultUserPeerBlocksModel struct {
		db *sqlx.DB
	}

	UserPeerBlocks struct {
		Id       int64 `db:"id" json:"id"`
		UserId   int64 `db:"user_id" json:"user_id"`
		PeerType int32 `db:"peer_type" json:"peer_type"`
		PeerId   int64 `db:"peer_id" json:"peer_id"`
		Date     int64 `db:"date" json:"date"`
		Deleted  bool  `db:"deleted" json:"deleted"`
	}
)

func newUserPeerBlocksModel(db *sqlx.DB) *defaultUserPeerBlocksModel {
	return &defaultUserPeerBlocksModel{
		db: db,
	}
}

func (m *defaultUserPeerBlocksModel) Insert2(ctx context.Context, data *UserPeerBlocks) (sql.Result, error) {
	tableName := "user_peer_blocks"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?)", tableName, userPeerBlocksRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.Date, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("user_peer_blocks.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserPeerBlocksModel) Delete2(ctx context.Context, id int64) error {
	tableName := "user_peer_blocks"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("user_peer_blocks.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserPeerBlocksModel) FindOne(ctx context.Context, id int64) (*UserPeerBlocks, error) {
	tableName := "user_peer_blocks"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", userPeerBlocksRows, tableName)
	var resp UserPeerBlocks

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_peer_blocks",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_peer_blocks.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserPeerBlocksModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserPeerBlocks, error) {
	if len(id) == 0 {
		return []UserPeerBlocks{}, nil
	}
	tableName := "user_peer_blocks"

	query := fmt.Sprintf("select %s from %s where id in (%s)", userPeerBlocksRows, tableName, sqlx.InInt64List(id))

	var resp []UserPeerBlocks
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserPeerBlocks{}, nil
		}
		return nil, fmt.Errorf("user_peer_blocks.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserPeerBlocksModel) Update2(ctx context.Context, data *UserPeerBlocks) error {
	tableName := "user_peer_blocks"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, userPeerBlocksRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.Date, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("user_peer_blocks.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserPeerBlocksModel) FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserPeerBlocks, error) {
	tableName := "user_peer_blocks"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND peer_type = ? AND peer_id = ? limit 1", userPeerBlocksRows, tableName)
	var resp UserPeerBlocks

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_peer_blocks",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_peer_blocks.FindOneByUserIdPeerTypePeerId: %w", err)
	}

	return &resp, nil
}
