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

type (
	bizDocumentsModel interface {
		Insert(ctx context.Context, data *Documents) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *Documents) (lastInsertId, rowsAffected int64, err error)

		SelectByDocumentId(ctx context.Context, documentId int64) (*Documents, error)

		SelectByDocumentIdList(ctx context.Context, idList []int64) ([]Documents, error)
		SelectByDocumentIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *Documents)) ([]Documents, error)

		SelectByIdList(ctx context.Context, idList []int64) ([]Documents, error)
		SelectByIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *Documents)) ([]Documents, error)
	}
)

// Insert
// insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)
func (m *defaultDocumentsModel) Insert(ctx context.Context, data *Documents) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("documents.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("documents.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("documents.Insert rows affected: %w", err)
	}

	return

}

// InsertTx
// insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)
func (m *defaultDocumentsModel) InsertTx(tx *sqlx.Tx, data *Documents) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("documents.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("documents.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("documents.InsertTx rows affected: %w", err)
	}

	return
}

// SelectByDocumentId
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id = :document_id
func (m *defaultDocumentsModel) SelectByDocumentId(ctx context.Context, documentId int64) (rValue *Documents, err error) {

	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id = ?"
		do    = &Documents{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, documentId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "documents",
				Key:      fmt.Sprintf("document_id=%v", documentId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("documents.SelectByDocumentId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByDocumentIdList
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (:idList)
func (m *defaultDocumentsModel) SelectByDocumentIdList(ctx context.Context, idList []int64) (rList []Documents, err error) {
	var (
		query  = fmt.Sprintf("select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (%s)", sqlx.InInt64List(idList))
		values []Documents
	)
	if len(idList) == 0 {
		rList = []Documents{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Documents{}
			err = nil
			return
		}
		err = fmt.Errorf("documents.SelectByDocumentIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectByDocumentIdListWithCB
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (:idList)
func (m *defaultDocumentsModel) SelectByDocumentIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *Documents)) (rList []Documents, err error) {
	var (
		query  = fmt.Sprintf("select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (%s)", sqlx.InInt64List(idList))
		values []Documents
	)
	if len(idList) == 0 {
		rList = []Documents{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Documents{}
			err = nil
			return
		}
		err = fmt.Errorf("documents.SelectByDocumentIdListWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectByIdList
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where id in (:idList)
func (m *defaultDocumentsModel) SelectByIdList(ctx context.Context, idList []int64) (rList []Documents, err error) {
	var (
		query  = fmt.Sprintf("select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where id in (%s)", sqlx.InInt64List(idList))
		values []Documents
	)
	if len(idList) == 0 {
		rList = []Documents{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Documents{}
			err = nil
			return
		}
		err = fmt.Errorf("documents.SelectByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where id in (:idList)
func (m *defaultDocumentsModel) SelectByIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *Documents)) (rList []Documents, err error) {
	var (
		query  = fmt.Sprintf("select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where id in (%s)", sqlx.InInt64List(idList))
		values []Documents
	)
	if len(idList) == 0 {
		rList = []Documents{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Documents{}
			err = nil
			return
		}
		err = fmt.Errorf("documents.SelectByIdListWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}
