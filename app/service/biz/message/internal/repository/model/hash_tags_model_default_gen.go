/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
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
	hashTagsFieldNames          = builder.RawFieldNames(&HashTags{})
	hashTagsRows                = strings.Join(hashTagsFieldNames, ",")
	hashTagsRowsExpectAutoSet   = strings.Join(stringx.Remove(hashTagsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	hashTagsRowsWithPlaceHolder = strings.Join(stringx.Remove(hashTagsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTHashTagsIdPrefix = "cache:t:hash_tags:id:"

	cacheHashTagsIdPrefix = "cache#HashTags#id"

	cacheHashTagsUserIdHashTagHashTagMessageIdPrefix = "cache#UserId#HashTag#HashTagMessageId"
)

type (
	hashTagsModel interface {
		Insert2(ctx context.Context, data *HashTags) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HashTags, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]HashTags, error)
		Update2(ctx context.Context, data *HashTags) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdHashTagHashTagMessageId(ctx context.Context, userId int64, hashTag string, hashTagMessageId int32) (*HashTags, error)
	}

	defaultHashTagsModel struct {
		db *sqlx.DB
	}

	HashTags struct {
		Id               int64  `db:"id" json:"id"`
		UserId           int64  `db:"user_id" json:"user_id"`
		PeerType         int32  `db:"peer_type" json:"peer_type"`
		PeerId           int64  `db:"peer_id" json:"peer_id"`
		HashTag          string `db:"hash_tag" json:"hash_tag"`
		HashTagMessageId int32  `db:"hash_tag_message_id" json:"hash_tag_message_id"`
		Deleted          bool   `db:"deleted" json:"deleted"`
	}
)

func newHashTagsModel(db *sqlx.DB) *defaultHashTagsModel {
	return &defaultHashTagsModel{
		db: db,
	}
}

func (m *defaultHashTagsModel) Insert2(ctx context.Context, data *HashTags) (sql.Result, error) {
	query := fmt.Sprintf("insert into `hash_tags` (%s) values (?, ?, ?, ?, ?, ?)", hashTagsRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.HashTag, data.HashTagMessageId, data.Deleted)
}

func (m *defaultHashTagsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `hash_tags` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultHashTagsModel) FindOne(ctx context.Context, id int64) (*HashTags, error) {
	query := fmt.Sprintf("select %s from hash_tags where id = ? limit 1", hashTagsRows)
	var resp HashTags
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultHashTagsModel) FindListByIdList(ctx context.Context, id ...int64) ([]HashTags, error) {
	if len(id) == 0 {
		return []HashTags{}, nil
	}

	query := fmt.Sprintf("select %s from hash_tags where id in (%s)", hashTagsRows, sqlx.InInt64List(id))

	var resp []HashTags
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultHashTagsModel) Update2(ctx context.Context, data *HashTags) error {
	query := fmt.Sprintf("update `hash_tags` set %s where `id` = ?", hashTagsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.HashTag, data.HashTagMessageId, data.Deleted, data.Id)
	return err
}

func (m *defaultHashTagsModel) FindOneByUserIdHashTagHashTagMessageId(ctx context.Context, userId int64, hashTag string, hashTagMessageId int32) (*HashTags, error) {
	query := fmt.Sprintf("select %s from hash_tags where user_id = ? AND hash_tag = ? AND hash_tag_message_id = ? limit 1", hashTagsRows)
	var resp HashTags
	err := m.db.QueryRowPartial(ctx, &resp, query, userId, hashTag, hashTagMessageId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultHashTagsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheHashTagsIdPrefix, primary)
}

func (m *defaultHashTagsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from hash_tags where id = ? limit 1", hashTagsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
