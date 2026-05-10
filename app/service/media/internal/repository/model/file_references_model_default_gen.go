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
	fileReferencesFieldNames          = builder.RawFieldNames(&FileReferences{})
	fileReferencesRows                = strings.Join(fileReferencesFieldNames, ",")
	fileReferencesRowsExpectAutoSet   = strings.Join(stringx.Remove(fileReferencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	fileReferencesRowsWithPlaceHolder = strings.Join(stringx.Remove(fileReferencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	fileReferencesModel interface {
		Insert2(ctx context.Context, data *FileReferences) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FileReferences, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]FileReferences, error)
		Update2(ctx context.Context, data *FileReferences) error
		Delete2(ctx context.Context, id int64) error
	}

	defaultFileReferencesModel struct {
		db *sqlx.DB
	}

	FileReferences struct {
		Id           int64  `db:"id" json:"id"`
		RefHash      []byte `db:"ref_hash" json:"ref_hash"`
		Domain       string `db:"domain" json:"domain"`
		MediaId      int64  `db:"media_id" json:"media_id"`
		AccessHash   int64  `db:"access_hash" json:"access_hash"`
		ObjectId     string `db:"object_id" json:"object_id"`
		OriginDomain string `db:"origin_domain" json:"origin_domain"`
		OriginId     int64  `db:"origin_id" json:"origin_id"`
		ExpireAt     int64  `db:"expire_at" json:"expire_at"`
		RevokedAt    int64  `db:"revoked_at" json:"revoked_at"`
	}
)

func newFileReferencesModel(db *sqlx.DB) *defaultFileReferencesModel {
	return &defaultFileReferencesModel{
		db: db,
	}
}

func (m *defaultFileReferencesModel) Insert2(ctx context.Context, data *FileReferences) (sql.Result, error) {
	tableName := "file_references"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, fileReferencesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.RefHash, data.Domain, data.MediaId, data.AccessHash, data.ObjectId, data.OriginDomain, data.OriginId, data.ExpireAt, data.RevokedAt)
	if err != nil {
		return nil, fmt.Errorf("file_references.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultFileReferencesModel) Delete2(ctx context.Context, id int64) error {
	tableName := "file_references"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("file_references.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultFileReferencesModel) FindOne(ctx context.Context, id int64) (*FileReferences, error) {
	tableName := "file_references"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", fileReferencesRows, tableName)
	var resp FileReferences

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "file_references",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("file_references.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultFileReferencesModel) FindListByIdList(ctx context.Context, id ...int64) ([]FileReferences, error) {
	if len(id) == 0 {
		return []FileReferences{}, nil
	}
	tableName := "file_references"

	query := fmt.Sprintf("select %s from %s where id in (%s)", fileReferencesRows, tableName, sqlx.InInt64List(id))

	var resp []FileReferences
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []FileReferences{}, nil
		}
		return nil, fmt.Errorf("file_references.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultFileReferencesModel) Update2(ctx context.Context, data *FileReferences) error {
	tableName := "file_references"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, fileReferencesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.RefHash, data.Domain, data.MediaId, data.AccessHash, data.ObjectId, data.OriginDomain, data.OriginId, data.ExpireAt, data.RevokedAt, data.Id)
	if err != nil {
		return fmt.Errorf("file_references.Update2 exec: %w", err)
	}

	return nil
}
