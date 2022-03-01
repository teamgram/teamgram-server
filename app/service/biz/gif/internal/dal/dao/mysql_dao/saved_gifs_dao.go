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
	"github.com/teamgram/teamgram-server/app/service/biz/gif/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type SavedGifsDAO struct {
	db *sqlx.DB
}

func NewSavedGifsDAO(db *sqlx.DB) *SavedGifsDAO {
	return &SavedGifsDAO{db}
}

// InsertIgnore
// insert ignore into saved_gifs(user_id, gif_id) values (:user_id, :gif_id)
// TODO(@benqi): sqlmap
func (dao *SavedGifsDAO) InsertIgnore(ctx context.Context, do *dataobject.SavedGifsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into saved_gifs(user_id, gif_id) values (:user_id, :gif_id)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertIgnore(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertIgnore(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertIgnore(%v)_error: %v", do, err)
	}

	return
}

// InsertIgnoreTx
// insert ignore into saved_gifs(user_id, gif_id) values (:user_id, :gif_id)
// TODO(@benqi): sqlmap
func (dao *SavedGifsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.SavedGifsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into saved_gifs(user_id, gif_id) values (:user_id, :gif_id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertIgnore(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertIgnore(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertIgnore(%v)_error: %v", do, err)
	}

	return
}

// SelectAll
// select gif_id from saved_gifs where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *SavedGifsDAO) SelectAll(ctx context.Context, user_id int64) (rList []int64, err error) {
	var query = "select gif_id from saved_gifs where user_id = ?"
	err = dao.db.Select(ctx, &rList, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectAll(_), error: %v", err)
	}

	return
}

// SelectAllWithCB
// select gif_id from saved_gifs where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *SavedGifsDAO) SelectAllWithCB(ctx context.Context, user_id int64, cb func(i int, v int64)) (rList []int64, err error) {
	var query = "select gif_id from saved_gifs where user_id = ?"
	err = dao.db.Select(ctx, &rList, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectAll(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// Delete
// delete from saved_gifs where user_id = :user_id and gif_id = :gif_id
// TODO(@benqi): sqlmap
func (dao *SavedGifsDAO) Delete(ctx context.Context, user_id int64, gif_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from saved_gifs where user_id = ? and gif_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, gif_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// DeleteTx
// delete from saved_gifs where user_id = :user_id and gif_id = :gif_id
// TODO(@benqi): sqlmap
func (dao *SavedGifsDAO) DeleteTx(tx *sqlx.Tx, user_id int64, gif_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from saved_gifs where user_id = ? and gif_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, gif_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}
