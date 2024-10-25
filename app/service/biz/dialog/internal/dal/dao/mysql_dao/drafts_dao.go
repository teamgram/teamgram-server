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

type DraftsDAO struct {
	db *sqlx.DB
}

func NewDraftsDAO(db *sqlx.DB) *DraftsDAO {
	return &DraftsDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)
func (dao *DraftsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.DraftsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)"
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
// insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)
func (dao *DraftsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DraftsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)"
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
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DraftsDAO) Select(ctx context.Context, userId int32, peerDialogId int64) (rValue *dataobject.DraftsDO, err error) {
	var (
		query = "select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id = ?"
		do    = &dataobject.DraftsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, userId, peerDialogId)

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

// SelectIdList
// select peer_dialog_id from drafts where user_id = :user_id
func (dao *DraftsDAO) SelectIdList(ctx context.Context, userId int32) (rList []int64, err error) {
	var query = "select peer_dialog_id from drafts where user_id = ?"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectIdList(_), error: %v", err)
	}

	return
}

// SelectIdListWithCB
// select peer_dialog_id from drafts where user_id = :user_id
func (dao *DraftsDAO) SelectIdListWithCB(ctx context.Context, userId int32, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select peer_dialog_id from drafts where user_id = ?"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectIdList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectByIdList
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id in (:idList)
func (dao *DraftsDAO) SelectByIdList(ctx context.Context, userId int32, idList []int64) (rValue *dataobject.DraftsDO, err error) {
	var (
		query = fmt.Sprintf("select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		do    = &dataobject.DraftsDO{}
	)

	if len(idList) == 0 {
		return
	}

	err = dao.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// ClearByIdList
// update drafts set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and peer_dialog_id in (:idList)
func (dao *DraftsDAO) ClearByIdList(ctx context.Context, userId int32, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update drafts set draft_type = 0, draft_message_data = 'null' where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = dao.db.Exec(ctx, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in ClearByIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in ClearByIdList(_), error: %v", err)
	}

	return
}

// ClearByIdListTx
// update drafts set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and peer_dialog_id in (:idList)
func (dao *DraftsDAO) ClearByIdListTx(tx *sqlx.Tx, userId int32, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update drafts set draft_type = 0, draft_message_data = 'null' where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = tx.Exec(query, userId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in ClearByIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in ClearByIdList(_), error: %v", err)
	}

	return
}
