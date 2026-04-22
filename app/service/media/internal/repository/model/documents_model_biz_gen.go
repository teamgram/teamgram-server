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
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *logx.Logger

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
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
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
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
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
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByDocumentId(_), error: %v", err)
			return
		} else {
			err = nil
		}
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
		logx.WithContext(ctx).Errorf("queryx in SelectByDocumentIdList(_), error: %v", err)
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
		logx.WithContext(ctx).Errorf("queryx in SelectByDocumentIdList(_), error: %v", err)
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
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
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
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
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
