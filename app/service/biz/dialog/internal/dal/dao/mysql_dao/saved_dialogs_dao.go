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
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type SavedDialogsDAO struct {
	db *sqlx.DB
}

func NewSavedDialogsDAO(db *sqlx.DB) *SavedDialogsDAO {
	return &SavedDialogsDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into saved_dialogs(user_id, peer_type, peer_id, pinned, top_message) values (:user_id, :peer_type, :peer_id, 0, :top_message) on duplicate key update top_message = values(top_message), deleted = 0
func (dao *SavedDialogsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.SavedDialogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into saved_dialogs(user_id, peer_type, peer_id, pinned, top_message) values (:user_id, :peer_type, :peer_id, 0, :top_message) on duplicate key update top_message = values(top_message), deleted = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into saved_dialogs(user_id, peer_type, peer_id, pinned, top_message) values (:user_id, :peer_type, :peer_id, 0, :top_message) on duplicate key update top_message = values(top_message), deleted = 0
func (dao *SavedDialogsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.SavedDialogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into saved_dialogs(user_id, peer_type, peer_id, pinned, top_message) values (:user_id, :peer_type, :peer_id, 0, :top_message) on duplicate key update top_message = values(top_message), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// Select
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
func (dao *SavedDialogsDAO) Select(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *dataobject.SavedDialogsDO, err error) {
	var (
		query = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		do    = &dataobject.SavedDialogsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectPinnedDialogs
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and pinned > 0 and deleted = 0 order by pinned desc
func (dao *SavedDialogsDAO) SelectPinnedDialogs(ctx context.Context, userId int64) (rList []dataobject.SavedDialogsDO, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and pinned > 0 and deleted = 0 order by pinned desc"
		values []dataobject.SavedDialogsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPinnedDialogsWithCB
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and pinned > 0 and deleted = 0 order by pinned desc
func (dao *SavedDialogsDAO) SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *dataobject.SavedDialogsDO)) (rList []dataobject.SavedDialogsDO, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and pinned > 0 and deleted = 0 order by pinned desc"
		values []dataobject.SavedDialogsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
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

// SelectExcludePinnedDialogs
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and pinned = 0 and top_message < :top_message and deleted = 0 order by top_message desc limit :limit
func (dao *SavedDialogsDAO) SelectExcludePinnedDialogs(ctx context.Context, userId int64, topMessage int32, limit int32) (rList []dataobject.SavedDialogsDO, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and pinned = 0 and top_message < ? and deleted = 0 order by top_message desc limit ?"
		values []dataobject.SavedDialogsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, topMessage, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludePinnedDialogs(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectExcludePinnedDialogsWithCB
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and pinned = 0 and top_message < :top_message and deleted = 0 order by top_message desc limit :limit
func (dao *SavedDialogsDAO) SelectExcludePinnedDialogsWithCB(ctx context.Context, userId int64, topMessage int32, limit int32, cb func(sz, i int, v *dataobject.SavedDialogsDO)) (rList []dataobject.SavedDialogsDO, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and pinned = 0 and top_message < ? and deleted = 0 order by top_message desc limit ?"
		values []dataobject.SavedDialogsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, topMessage, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludePinnedDialogs(_), error: %v", err)
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

// SelectDialogs
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and top_message < :top_message and deleted = 0 order by top_message desc limit :limit
func (dao *SavedDialogsDAO) SelectDialogs(ctx context.Context, userId int64, topMessage int32, limit int32) (rList []dataobject.SavedDialogsDO, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and top_message < ? and deleted = 0 order by top_message desc limit ?"
		values []dataobject.SavedDialogsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, topMessage, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogs(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDialogsWithCB
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and top_message < :top_message and deleted = 0 order by top_message desc limit :limit
func (dao *SavedDialogsDAO) SelectDialogsWithCB(ctx context.Context, userId int64, topMessage int32, limit int32, cb func(sz, i int, v *dataobject.SavedDialogsDO)) (rList []dataobject.SavedDialogsDO, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and top_message < ? and deleted = 0 order by top_message desc limit ?"
		values []dataobject.SavedDialogsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, topMessage, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogs(_), error: %v", err)
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

// UpdateUserUnPinned
// update saved_dialogs set pinned = 0 where user_id = :user_id and pinned > 0 and deleted = 0
func (dao *SavedDialogsDAO) UpdateUserUnPinned(ctx context.Context, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = 0 where user_id = ? and pinned > 0 and deleted = 0"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateUserUnPinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateUserUnPinned(_), error: %v", err)
	}

	return
}

// UpdateUserUnPinnedTx
// update saved_dialogs set pinned = 0 where user_id = :user_id and pinned > 0 and deleted = 0
func (dao *SavedDialogsDAO) UpdateUserUnPinnedTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = 0 where user_id = ? and pinned > 0 and deleted = 0"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateUserUnPinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateUserUnPinned(_), error: %v", err)
	}

	return
}

// UpdateUserPeerPinned
// update saved_dialogs set pinned = :pinned where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *SavedDialogsDAO) UpdateUserPeerPinned(ctx context.Context, pinned int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, pinned, userId, peerType, peerId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateUserPeerPinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateUserPeerPinned(_), error: %v", err)
	}

	return
}

// UpdateUserPeerPinnedTx
// update saved_dialogs set pinned = :pinned where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *SavedDialogsDAO) UpdateUserPeerPinnedTx(tx *sqlx.Tx, pinned int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, pinned, userId, peerType, peerId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateUserPeerPinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateUserPeerPinned(_), error: %v", err)
	}

	return
}
