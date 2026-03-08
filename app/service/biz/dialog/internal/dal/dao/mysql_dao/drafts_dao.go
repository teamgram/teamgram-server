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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

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
		query string
		r     sql.Result
	)
	query = "insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)
func (dao *DraftsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DraftsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// Select
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DraftsDAO) Select(ctx context.Context, userId int32, peerDialogId int64) (rValue *dataobject.DraftsDO, err error) {
	var (
		query string
		do    = &dataobject.DraftsDO{}
	)
	query = "select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id = ?"

	err = dao.db.QueryRowPartial(ctx, do, query, userId, peerDialogId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
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
	var query string
	query = "select peer_dialog_id from drafts where user_id = ?"

	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectIdList(_), error: %v", err)
	}

	return
}

// SelectIdListWithCB
// select peer_dialog_id from drafts where user_id = :user_id
func (dao *DraftsDAO) SelectIdListWithCB(ctx context.Context, userId int32, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query string
	query = "select peer_dialog_id from drafts where user_id = ?"

	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectIdList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectByIdList
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id in (:idList)
func (dao *DraftsDAO) SelectByIdList(ctx context.Context, userId int32, idList []int64) (rList []dataobject.DraftsDO, err error) {

	if len(idList) == 0 {
		rList = []dataobject.DraftsDO{}
		return
	}

	var (
		query  string
		values []dataobject.DraftsDO
	)
	query = fmt.Sprintf("select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id in (:idList)
func (dao *DraftsDAO) SelectByIdListWithCB(ctx context.Context, userId int32, idList []int64, cb func(sz, i int, v *dataobject.DraftsDO)) (rList []dataobject.DraftsDO, err error) {

	if len(idList) == 0 {
		rList = []dataobject.DraftsDO{}
		return
	}

	var (
		query  string
		values []dataobject.DraftsDO
	)
	query = fmt.Sprintf("select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// ClearByIdList
// update drafts set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and peer_dialog_id in (:idList)
func (dao *DraftsDAO) ClearByIdList(ctx context.Context, userId int32, idList []int64) (rowsAffected int64, err error) {

	if len(idList) == 0 {
		return
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update drafts set draft_type = 0, draft_message_data = 'null' where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))

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

	if len(idList) == 0 {
		return
	}
	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update drafts set draft_type = 0, draft_message_data = 'null' where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))

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
