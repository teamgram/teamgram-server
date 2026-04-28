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
	hashTagsFieldNames          = builder.RawFieldNames(&HashTags{})
	hashTagsRows                = strings.Join(hashTagsFieldNames, ",")
	hashTagsRowsExpectAutoSet   = strings.Join(stringx.Remove(hashTagsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	hashTagsRowsWithPlaceHolder = strings.Join(stringx.Remove(hashTagsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
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

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.HashTag, data.HashTagMessageId, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("hash_tags.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultHashTagsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `hash_tags` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("hash_tags.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultHashTagsModel) FindOne(ctx context.Context, id int64) (*HashTags, error) {
	query := fmt.Sprintf("select %s from hash_tags where id = ? limit 1", hashTagsRows)
	var resp HashTags

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "hash_tags",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("hash_tags.FindOne: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return []HashTags{}, nil
		}
		return nil, fmt.Errorf("hash_tags.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultHashTagsModel) Update2(ctx context.Context, data *HashTags) error {
	query := fmt.Sprintf("update `hash_tags` set %s where `id` = ?", hashTagsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.HashTag, data.HashTagMessageId, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("hash_tags.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultHashTagsModel) FindOneByUserIdHashTagHashTagMessageId(ctx context.Context, userId int64, hashTag string, hashTagMessageId int32) (*HashTags, error) {
	query := fmt.Sprintf("select %s from hash_tags where user_id = ? AND hash_tag = ? AND hash_tag_message_id = ? limit 1", hashTagsRows)
	var resp HashTags

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, hashTag, hashTagMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "hash_tags",
				Key:      fmt.Sprintf("user_id=%v,hash_tag=%v,hash_tag_message_id=%v", userId, hashTag, hashTagMessageId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("hash_tags.FindOneByUserIdHashTagHashTagMessageId: %w", err)
	}

	return &resp, nil
}
