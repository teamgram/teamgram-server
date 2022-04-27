/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
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

type EncryptedFilesDAO struct {
	db *sqlx.DB
}

func NewEncryptedFilesDAO(db *sqlx.DB) *EncryptedFilesDAO {
	return &EncryptedFilesDAO{db}
}

// Insert
// insert into encrypted_files(encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path) values (:encrypted_file_id, :access_hash, :dc_id, :file_size, :key_fingerprint, :md5_checksum, :file_path)
// TODO(@benqi): sqlmap
func (dao *EncryptedFilesDAO) Insert(ctx context.Context, do *dataobject.EncryptedFilesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into encrypted_files(encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path) values (:encrypted_file_id, :access_hash, :dc_id, :file_size, :key_fingerprint, :md5_checksum, :file_path)"
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
// insert into encrypted_files(encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path) values (:encrypted_file_id, :access_hash, :dc_id, :file_size, :key_fingerprint, :md5_checksum, :file_path)
// TODO(@benqi): sqlmap
func (dao *EncryptedFilesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.EncryptedFilesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into encrypted_files(encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path) values (:encrypted_file_id, :access_hash, :dc_id, :file_size, :key_fingerprint, :md5_checksum, :file_path)"
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
// select id, encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path from encrypted_files where dc_id = 2 and encrypted_file_id = :encrypted_file_id and access_hash = :access_hash
// TODO(@benqi): sqlmap
func (dao *EncryptedFilesDAO) SelectByFileLocation(ctx context.Context, encrypted_file_id int64, access_hash int64) (rValue *dataobject.EncryptedFilesDO, err error) {
	var (
		query = "select id, encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path from encrypted_files where dc_id = 2 and encrypted_file_id = ? and access_hash = ?"
		do    = &dataobject.EncryptedFilesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, encrypted_file_id, access_hash)

	if err != nil {
		if err != sqlx.ErrNotFound {
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

// SelectByIdList
// select id, encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path from encrypted_files where encrypted_file_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *EncryptedFilesDAO) SelectByIdList(ctx context.Context, idList []int64) (rList []dataobject.EncryptedFilesDO, err error) {
	var (
		query  = "select id, encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path from encrypted_files where encrypted_file_id in (?)"
		a      []interface{}
		values []dataobject.EncryptedFilesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.EncryptedFilesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path from encrypted_files where encrypted_file_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *EncryptedFilesDAO) SelectByIdListWithCB(ctx context.Context, idList []int64, cb func(i int, v *dataobject.EncryptedFilesDO)) (rList []dataobject.EncryptedFilesDO, err error) {
	var (
		query  = "select id, encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path from encrypted_files where encrypted_file_id in (?)"
		a      []interface{}
		values []dataobject.EncryptedFilesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.EncryptedFilesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}
