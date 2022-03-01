/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type DocumentsDAO struct {
	db *sqlx.DB
}

func NewDocumentsDAO(db *sqlx.DB) *DocumentsDAO {
	return &DocumentsDAO{db}
}

// Insert
// insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, date2) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :video_thumb_id, :attributes, :date2)
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
func (dao *DocumentsDAO) SelectByFileLocation(ctx context.Context, document_id int64, access_hash int64, version int32) (rValue *dataobject.DocumentsDO, err error) {
	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where dc_id = 2 and document_id = ? and access_hash = ? and version = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, document_id, access_hash, version)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByFileLocation(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.DocumentsDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByFileLocation(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectById
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id = :document_id
// TODO(@benqi): sqlmap
func (dao *DocumentsDAO) SelectById(ctx context.Context, document_id int64) (rValue *dataobject.DocumentsDO, err error) {
	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, document_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectById(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.DocumentsDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectById(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectByIdList
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *DocumentsDAO) SelectByIdList(ctx context.Context, idList []int64) (rList []dataobject.DocumentsDO, err error) {
	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.DocumentsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DocumentsDO
	for rows.Next() {
		v := dataobject.DocumentsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByIdListWithCB
// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *DocumentsDAO) SelectByIdListWithCB(ctx context.Context, idList []int64, cb func(i int, v *dataobject.DocumentsDO)) (rList []dataobject.DocumentsDO, err error) {
	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, video_thumb_id, attributes, version, date2 from documents where document_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.DocumentsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.DocumentsDO
	for rows.Next() {
		v := dataobject.DocumentsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
