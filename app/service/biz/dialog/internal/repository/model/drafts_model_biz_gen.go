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
var _ *sqlx.Tx

type bizDraftsModel interface {
	InsertOrUpdate(ctx context.Context, data *Drafts) (lastInsertId, rowsAffected int64, err error)
	Select(ctx context.Context, userId int32, peerDialogId int64) (*Drafts, error)
	SelectIdList(ctx context.Context, userId int32) ([]int64, error)
	SelectIdListWithCB(ctx context.Context, userId int32, cb func(sz, i int, v int64)) ([]int64, error)
	SelectByIdList(ctx context.Context, userId int32, idList []int64) ([]Drafts, error)
	SelectByIdListWithCB(ctx context.Context, userId int32, idList []int64, cb func(sz, i int, v *Drafts)) ([]Drafts, error)
	ClearByIdList(ctx context.Context, userId int32, idList []int64) (rowsAffected int64, err error)
}

type DraftsTxModel interface {
	InsertOrUpdate(data *Drafts) (lastInsertId, rowsAffected int64, err error)
	Select(userId int32, peerDialogId int64) (*Drafts, error)
	SelectIdList(userId int32) ([]int64, error)
	SelectByIdList(userId int32, idList []int64) ([]Drafts, error)
	ClearByIdList(userId int32, idList []int64) (rowsAffected int64, err error)
}

type defaultDraftsTxModel struct {
	tx *sqlx.Tx
}

func NewDraftsTxModel(tx *sqlx.Tx) DraftsTxModel {
	return &defaultDraftsTxModel{tx: tx}
}

// InsertOrUpdate
// insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)
func (m *defaultDraftsModel) InsertOrUpdate(ctx context.Context, data *Drafts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("drafts.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("drafts.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("drafts.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)
func (m *defaultDraftsTxModel) InsertOrUpdate(data *Drafts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into drafts(user_id, peer_dialog_id, draft_type, draft_message_data, date2) values (:user_id, :peer_dialog_id, :draft_type, :draft_message_data, :date2) on duplicate key update draft_type = values(draft_type), draft_message_data = values(draft_message_data), date2 = values(date2)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("drafts.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("drafts.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("drafts.InsertOrUpdate rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "drafts",
				Key:      fmt.Sprintf("user_id=%v,peer_dialog_id=%v", userId, peerDialogId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("drafts.Select: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// Select
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDraftsTxModel) Select(userId int32, peerDialogId int64) (rValue *Drafts, err error) {
	var (
		query = "select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id = ?"
		do    = &Drafts{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerDialogId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "drafts",
				Key:      fmt.Sprintf("user_id=%v,peer_dialog_id=%v", userId, peerDialogId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("drafts.Select: %w", err)
		return
	}
	rValue = do

	return
}

// SelectIdList
// select peer_dialog_id from drafts where user_id = :user_id
func (m *defaultDraftsModel) SelectIdList(ctx context.Context, userId int32) (rList []int64, err error) {
	var query = "select peer_dialog_id from drafts where user_id = ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("drafts.SelectIdList: %w", err)
	}

	return
}

// SelectIdList
// select peer_dialog_id from drafts where user_id = :user_id
func (m *defaultDraftsTxModel) SelectIdList(userId int32) (rList []int64, err error) {
	var query = "select peer_dialog_id from drafts where user_id = ?"
	err = m.tx.QueryRowsPartial(&rList, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("drafts.SelectIdList: %w", err)
	}

	return
}

// SelectIdListWithCB
// select peer_dialog_id from drafts where user_id = :user_id
func (m *defaultDraftsModel) SelectIdListWithCB(ctx context.Context, userId int32, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select peer_dialog_id from drafts where user_id = ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("drafts.SelectIdListWithCB: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Drafts{}
			err = nil
			return
		}
		err = fmt.Errorf("drafts.SelectByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectByIdList
// select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = :user_id and peer_dialog_id in (:idList)
func (m *defaultDraftsTxModel) SelectByIdList(userId int32, idList []int64) (rList []Drafts, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_dialog_id, draft_type, draft_message_data, date2 from drafts where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		values []Drafts
	)
	if len(idList) == 0 {
		rList = []Drafts{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Drafts{}
			err = nil
			return
		}
		err = fmt.Errorf("drafts.SelectByIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Drafts{}
			err = nil
			return
		}
		err = fmt.Errorf("drafts.SelectByIdListWithCB: %w", err)
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
		err = fmt.Errorf("drafts.ClearByIdList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("drafts.ClearByIdList rows affected: %w", err)
		return
	}

	return
}

// ClearByIdList
// update drafts set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and peer_dialog_id in (:idList)
func (m *defaultDraftsTxModel) ClearByIdList(userId int32, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update drafts set draft_type = 0, draft_message_data = 'null' where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("drafts.ClearByIdList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("drafts.ClearByIdList rows affected: %w", err)
		return
	}

	return
}
