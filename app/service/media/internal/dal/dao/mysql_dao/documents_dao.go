/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type DocumentsDAO struct {
	db *sqlx.DB
}

func NewDocumentsDAO(db *sqlx.DB) *DocumentsDAO {
	return &DocumentsDAO{
		db: db,
	}
}

// Insert
// insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)
func (dao *DocumentsDAO) Insert(ctx context.Context, do *dataobject.DocumentsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)
func (dao *DocumentsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.DocumentsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// SelectByFileLocation
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where dc_id = 2 and document_id = :document_id and access_hash = :access_hash and version = :version
func (dao *DocumentsDAO) SelectByFileLocation(ctx context.Context, documentId int64, accessHash int64, version int32) (rValue *dataobject.DocumentsDO, err error) {
	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where dc_id = 2 and document_id = ? and access_hash = ? and version = ?"
		do    = &dataobject.DocumentsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, documentId, accessHash, version)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByFileLocation(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByDocumentId
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id = :document_id
func (dao *DocumentsDAO) SelectByDocumentId(ctx context.Context, documentId int64) (rValue *dataobject.DocumentsDO, err error) {
	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id = ?"
		do    = &dataobject.DocumentsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, documentId)

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
func (dao *DocumentsDAO) SelectByDocumentIdList(ctx context.Context, idList []int64) (rList []dataobject.DocumentsDO, err error) {
	var (
		query  = fmt.Sprintf("select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.DocumentsDO
	)
	if len(idList) == 0 {
		rList = []dataobject.DocumentsDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByDocumentIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByDocumentIdListWithCB
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (:idList)
func (dao *DocumentsDAO) SelectByDocumentIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *dataobject.DocumentsDO)) (rList []dataobject.DocumentsDO, err error) {
	var (
		query  = fmt.Sprintf("select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.DocumentsDO
	)
	if len(idList) == 0 {
		rList = []dataobject.DocumentsDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

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
func (dao *DocumentsDAO) SelectByIdList(ctx context.Context, idList []int64) (rList []dataobject.DocumentsDO, err error) {
	var (
		query  = fmt.Sprintf("select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.DocumentsDO
	)
	if len(idList) == 0 {
		rList = []dataobject.DocumentsDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where id in (:idList)
func (dao *DocumentsDAO) SelectByIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *dataobject.DocumentsDO)) (rList []dataobject.DocumentsDO, err error) {
	var (
		query  = fmt.Sprintf("select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.DocumentsDO
	)
	if len(idList) == 0 {
		rList = []dataobject.DocumentsDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

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
