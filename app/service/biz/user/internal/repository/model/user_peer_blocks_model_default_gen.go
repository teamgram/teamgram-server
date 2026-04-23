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
	query := fmt.Sprintf("insert into `user_peer_blocks` (%s) values (?, ?, ?, ?, ?)", userPeerBlocksRowsExpectAutoSet)

	return m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.Date, data.Deleted)

}

func (m *defaultUserPeerBlocksModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_peer_blocks` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserPeerBlocksModel) FindOne(ctx context.Context, id int64) (*UserPeerBlocks, error) {
	query := fmt.Sprintf("select %s from user_peer_blocks where id = ? limit 1", userPeerBlocksRows)
	var resp UserPeerBlocks

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserPeerBlocksModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserPeerBlocks, error) {
	if len(id) == 0 {
		return []UserPeerBlocks{}, nil
	}

	query := fmt.Sprintf("select %s from user_peer_blocks where id in (%s)", userPeerBlocksRows, sqlx.InInt64List(id))

	var resp []UserPeerBlocks
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserPeerBlocksModel) Update2(ctx context.Context, data *UserPeerBlocks) error {
	query := fmt.Sprintf("update `user_peer_blocks` set %s where `id` = ?", userPeerBlocksRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.Date, data.Deleted, data.Id)
	return err
}

func (m *defaultUserPeerBlocksModel) FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserPeerBlocks, error) {
	query := fmt.Sprintf("select %s from user_peer_blocks where user_id = ? AND peer_type = ? AND peer_id = ? limit 1", userPeerBlocksRows)
	var resp UserPeerBlocks

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}
