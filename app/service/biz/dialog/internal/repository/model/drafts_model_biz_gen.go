/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
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
	bizDraftsModel interface {
		InsertOrUpdate(ctx context.Context, data *Drafts) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *Drafts) (lastInsertId, rowsAffected int64, err error)

		Select(ctx context.Context, userId int32, peerDialogId int64) (*Drafts, error)

		SelectIdList(ctx context.Context, userId int32) ([]int64, error)
		SelectIdListWithCB(ctx context.Context, userId int32, cb func(sz, i int, v int64)) ([]int64, error)

		SelectByIdList(ctx context.Context, userId int32, idList []int64) ([]Drafts, error)
		SelectByIdListWithCB(ctx context.Context, userId int32, idList []int64, cb func(sz, i int, v *Drafts)) ([]Drafts, error)

		ClearByIdList(ctx context.Context, userId int32, idList []int64) (rowsAffected int64, err error)
		ClearByIdListTx(tx *sqlx.Tx, userId int32, idList []int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)
func (m *defaultDraftsModel) InsertOrUpdate(ctx context.Context, data *Drafts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// InsertOrUpdateTx
// insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)
func (m *defaultDraftsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *Drafts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// Select
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDraftsModel) Select(ctx context.Context, userId int32, peerDialogId int64) (rValue *Drafts, err error) {
	var (
		query = "select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id = ?"
		do    = &Drafts{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerDialogId)

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
func (m *defaultDraftsModel) SelectIdList(ctx context.Context, userId int32) (rList []int64, err error) {
	var query = "select peer_dialog_id from drafts where user_id = ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectIdList(_), error: %v", err)
	}

	return
}

// SelectIdListWithCB
// select peer_dialog_id from drafts where user_id = :user_id
func (m *defaultDraftsModel) SelectIdListWithCB(ctx context.Context, userId int32, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select peer_dialog_id from drafts where user_id = ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

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
func (m *defaultDraftsModel) SelectByIdList(ctx context.Context, userId int32, idList []int64) (rList []Drafts, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		values []Drafts
	)
	if len(idList) == 0 {
		rList = []Drafts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id in (:idList)
func (m *defaultDraftsModel) SelectByIdListWithCB(ctx context.Context, userId int32, idList []int64, cb func(sz, i int, v *Drafts)) (rList []Drafts, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		values []Drafts
	)
	if len(idList) == 0 {
		rList = []Drafts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

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

// ClearByIdList
// update drafts set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and peer_dialog_id in (:idList)
func (m *defaultDraftsModel) ClearByIdList(ctx context.Context, userId int32, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update drafts set draft_type = 0, draft_message_data = 'null' where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, userId)

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
func (m *defaultDraftsModel) ClearByIdListTx(tx *sqlx.Tx, userId int32, idList []int64) (rowsAffected int64, err error) {
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
