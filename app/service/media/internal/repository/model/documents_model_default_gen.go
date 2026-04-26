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
	documentsFieldNames          = builder.RawFieldNames(&Documents{})
	documentsRows                = strings.Join(documentsFieldNames, ",")
	documentsRowsExpectAutoSet   = strings.Join(stringx.Remove(documentsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	documentsRowsWithPlaceHolder = strings.Join(stringx.Remove(documentsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	documentsModel interface {
		Insert2(ctx context.Context, data *Documents) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Documents, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Documents, error)
		Update2(ctx context.Context, data *Documents) error
		Delete2(ctx context.Context, id int64) error

		FindOneByDocumentId(ctx context.Context, documentId int64) (*Documents, error)
		FindListByDocumentIdList(ctx context.Context, documentId ...int64) ([]Documents, error)
	}

	defaultDocumentsModel struct {
		db *sqlx.DB
	}

	Documents struct {
		Id               int64  `db:"id" json:"id"`
		DocumentId       int64  `db:"document_id" json:"document_id"`
		AccessHash       int64  `db:"access_hash" json:"access_hash"`
		DcId             int32  `db:"dc_id" json:"dc_id"`
		FilePath         string `db:"file_path" json:"file_path"`
		FileSize         int64  `db:"file_size" json:"file_size"`
		UploadedFileName string `db:"uploaded_file_name" json:"uploaded_file_name"`
		Ext              string `db:"ext" json:"ext"`
		MimeType         string `db:"mime_type" json:"mime_type"`
		ThumbId          int64  `db:"thumb_id" json:"thumb_id"`
		VideoThumbId     int64  `db:"video_thumb_id" json:"video_thumb_id"`
		Version          int32  `db:"version" json:"version"`
		Attributes       string `db:"attributes" json:"attributes"`
		Date2            int64  `db:"date2" json:"date2"`
		ImportDocumentId int64  `db:"import_document_id" json:"import_document_id"`
		Deleted          bool   `db:"deleted" json:"deleted"`
	}
)

func newDocumentsModel(db *sqlx.DB) *defaultDocumentsModel {
	return &defaultDocumentsModel{
		db: db,
	}
}

func (m *defaultDocumentsModel) Insert2(ctx context.Context, data *Documents) (sql.Result, error) {
	query := fmt.Sprintf("insert into `documents` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", documentsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.DocumentId, data.AccessHash, data.DcId, data.FilePath, data.FileSize, data.UploadedFileName, data.Ext, data.MimeType, data.ThumbId, data.VideoThumbId, data.Version, data.Attributes, data.Date2, data.ImportDocumentId, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("documents.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDocumentsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `documents` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("documents.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultDocumentsModel) FindOne(ctx context.Context, id int64) (*Documents, error) {
	query := fmt.Sprintf("select %s from documents where id = ? limit 1", documentsRows)
	var resp Documents

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "documents",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("documents.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultDocumentsModel) FindListByIdList(ctx context.Context, id ...int64) ([]Documents, error) {
	if len(id) == 0 {
		return []Documents{}, nil
	}

	query := fmt.Sprintf("select %s from documents where id in (%s)", documentsRows, sqlx.InInt64List(id))

	var resp []Documents
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("documents.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultDocumentsModel) Update2(ctx context.Context, data *Documents) error {
	query := fmt.Sprintf("update `documents` set %s where `id` = ?", documentsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.DocumentId, data.AccessHash, data.DcId, data.FilePath, data.FileSize, data.UploadedFileName, data.Ext, data.MimeType, data.ThumbId, data.VideoThumbId, data.Version, data.Attributes, data.Date2, data.ImportDocumentId, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("documents.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultDocumentsModel) FindOneByDocumentId(ctx context.Context, documentId int64) (*Documents, error) {
	query := fmt.Sprintf("select %s from documents where document_id = ? limit 1", documentsRows)
	var resp Documents

	err := m.db.QueryRowPartial(ctx, &resp, query, documentId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "documents",
				Key:      fmt.Sprintf("document_id=%v", documentId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("documents.FindOneByDocumentId: %w", err)
	}

	return &resp, nil
}

func (m *defaultDocumentsModel) FindListByDocumentIdList(ctx context.Context, documentId ...int64) ([]Documents, error) {
	if len(documentId) == 0 {
		return []Documents{}, nil
	}

	query := fmt.Sprintf("select %s from documents where document_id in (%s)", documentsRows, sqlx.InInt64List(documentId))

	var resp []Documents
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("documents.FindListByDocumentIdList: %w", err)
	}

	return resp, nil
}
