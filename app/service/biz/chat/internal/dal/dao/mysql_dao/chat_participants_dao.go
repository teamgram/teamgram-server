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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type ChatParticipantsDAO struct {
	db *sqlx.DB
}

func NewChatParticipantsDAO(db *sqlx.DB) *ChatParticipantsDAO {
	return &ChatParticipantsDAO{db}
}

// Insert
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) Insert(ctx context.Context, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2)"
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
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2)"
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

// InsertBulk
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) InsertBulk(ctx context.Context, doList []*dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

// InsertBulkTx
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

// InsertOrUpdate
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), link = values(link), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0, date2 = values(date2)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), link = values(link), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0, date2 = values(date2)"
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
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), link = values(link), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0, date2 = values(date2)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), link = values(link), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0, date2 = values(date2)"
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

// SelectList
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = :chat_id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectList(ctx context.Context, chat_id int64) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, chat_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = :chat_id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectListWithCB(ctx context.Context, chat_id int64, cb func(i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, chat_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
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

// SelectByParticipantId
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = :chat_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectByParticipantId(ctx context.Context, chat_id int64, user_id int64) (rValue *dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = ? and user_id = ?"
		do    = &dataobject.ChatParticipantsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, chat_id, user_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in SelectByParticipantId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectListByParticipantIdList
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = :chat_id and user_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectListByParticipantIdList(ctx context.Context, chat_id int64, idList []int64) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = ? and user_id in (?)"
		a      []interface{}
		values []dataobject.ChatParticipantsDO
	)

	if len(idList) == 0 {
		rList = []dataobject.ChatParticipantsDO{}
		return
	}

	query, a, err = sqlx.In(query, chat_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByParticipantIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByParticipantIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByParticipantIdListWithCB
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = :chat_id and user_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectListByParticipantIdListWithCB(ctx context.Context, chat_id int64, idList []int64, cb func(i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = ? and user_id in (?)"
		a      []interface{}
		values []dataobject.ChatParticipantsDO
	)

	if len(idList) == 0 {
		rList = []dataobject.ChatParticipantsDO{}
		return
	}

	query, a, err = sqlx.In(query, chat_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByParticipantIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByParticipantIdList(_), error: %v", err)
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

// Update
// update chat_participants set participant_type = :participant_type, inviter_user_id = :inviter_user_id, invited_at = :invited_at, state = 0, kicked_at = 0, left_at = 0 where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) Update(ctx context.Context, participant_type int32, inviter_user_id int64, invited_at int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_type, inviter_user_id, invited_at, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

// update chat_participants set participant_type = :participant_type, inviter_user_id = :inviter_user_id, invited_at = :invited_at, state = 0, kicked_at = 0, left_at = 0 where id = :id
// UpdateTx
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateTx(tx *sqlx.Tx, participant_type int32, inviter_user_id int64, invited_at int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_type, inviter_user_id, invited_at, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

// UpdateKicked
// update chat_participants set kicked_at = :kicked_at, left_at = 0, state = 2 where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateKicked(ctx context.Context, kicked_at int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set kicked_at = ?, left_at = 0, state = 2 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, kicked_at, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateKicked(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateKicked(_), error: %v", err)
	}

	return
}

// update chat_participants set kicked_at = :kicked_at, left_at = 0, state = 2 where id = :id
// UpdateKickedTx
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateKickedTx(tx *sqlx.Tx, kicked_at int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set kicked_at = ?, left_at = 0, state = 2 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, kicked_at, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateKicked(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateKicked(_), error: %v", err)
	}

	return
}

// UpdateLeft
// update chat_participants set kicked_at = 0, left_at = :left_at, state = 1 where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateLeft(ctx context.Context, left_at int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set kicked_at = 0, left_at = ?, state = 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, left_at, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateLeft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateLeft(_), error: %v", err)
	}

	return
}

// update chat_participants set kicked_at = 0, left_at = :left_at, state = 1 where id = :id
// UpdateLeftTx
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateLeftTx(tx *sqlx.Tx, left_at int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set kicked_at = 0, left_at = ?, state = 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, left_at, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateLeft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateLeft(_), error: %v", err)
	}

	return
}

// UpdateParticipantType
// update chat_participants set participant_type = :participant_type where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateParticipantType(ctx context.Context, participant_type int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_type, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateParticipantType(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateParticipantType(_), error: %v", err)
	}

	return
}

// update chat_participants set participant_type = :participant_type where id = :id
// UpdateParticipantTypeTx
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateParticipantTypeTx(tx *sqlx.Tx, participant_type int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_type, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateParticipantType(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateParticipantType(_), error: %v", err)
	}

	return
}

// SelectUsersChatIdList
// select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (:idList) and chats.id = chat_participants.chat_id and chats.deactivated = 0
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectUsersChatIdList(ctx context.Context, idList []int64) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (?) and chats.id = chat_participants.chat_id and chats.deactivated = 0"
		a      []interface{}
		values []dataobject.ChatParticipantsDO
	)
	if len(idList) == 0 {
		rList = []dataobject.ChatParticipantsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectUsersChatIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersChatIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUsersChatIdListWithCB
// select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (:idList) and chats.id = chat_participants.chat_id and chats.deactivated = 0
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectUsersChatIdListWithCB(ctx context.Context, idList []int64, cb func(i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (?) and chats.id = chat_participants.chat_id and chats.deactivated = 0"
		a      []interface{}
		values []dataobject.ChatParticipantsDO
	)
	if len(idList) == 0 {
		rList = []dataobject.ChatParticipantsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectUsersChatIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersChatIdList(_), error: %v", err)
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

// SelectMyAdminList
// select chat_id from chat_participants where user_id = :user_id and participant_type = 1 and state = 0
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectMyAdminList(ctx context.Context, user_id int64) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and participant_type = 1 and state = 0"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectMyAdminList(_), error: %v", err)
	}

	return
}

// SelectMyAdminListWithCB
// select chat_id from chat_participants where user_id = :user_id and participant_type = 1 and state = 0
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectMyAdminListWithCB(ctx context.Context, user_id int64, cb func(i int, v int64)) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and participant_type = 1 and state = 0"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectMyAdminList(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// SelectMyAllList
// select chat_id from chat_participants where user_id = :user_id and state = 0
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectMyAllList(ctx context.Context, user_id int64) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and state = 0"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectMyAllList(_), error: %v", err)
	}

	return
}

// SelectMyAllListWithCB
// select chat_id from chat_participants where user_id = :user_id and state = 0
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) SelectMyAllListWithCB(ctx context.Context, user_id int64, cb func(i int, v int64)) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and state = 0"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectMyAllList(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// UpdateStateByChatId
// update chat_participants set state = :state where chat_id = :chat_id and state = 0
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateStateByChatId(ctx context.Context, state int32, chat_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = ? where chat_id = ? and state = 0"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, state, chat_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateStateByChatId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateStateByChatId(_), error: %v", err)
	}

	return
}

// update chat_participants set state = :state where chat_id = :chat_id and state = 0
// UpdateStateByChatIdTx
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateStateByChatIdTx(tx *sqlx.Tx, state int32, chat_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = ? where chat_id = ? and state = 0"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, chat_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateStateByChatId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateStateByChatId(_), error: %v", err)
	}

	return
}

// UpdateLink
// update chat_participants set link = :link where chat_id = :chat_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateLink(ctx context.Context, link string, chat_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set link = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, link, chat_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateLink(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateLink(_), error: %v", err)
	}

	return
}

// update chat_participants set link = :link where chat_id = :chat_id and user_id = :user_id
// UpdateLinkTx
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateLinkTx(tx *sqlx.Tx, link string, chat_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set link = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, link, chat_id, user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateLink(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateLink(_), error: %v", err)
	}

	return
}

// UpdateLinkUsage
// update chat_participants set usage2 = :usage2 where chat_id = :chat_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateLinkUsage(ctx context.Context, usage2 int32, chat_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set usage2 = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, usage2, chat_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateLinkUsage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateLinkUsage(_), error: %v", err)
	}

	return
}

// update chat_participants set usage2 = :usage2 where chat_id = :chat_id and user_id = :user_id
// UpdateLinkUsageTx
// TODO(@benqi): sqlmap
func (dao *ChatParticipantsDAO) UpdateLinkUsageTx(tx *sqlx.Tx, usage2 int32, chat_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set usage2 = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, usage2, chat_id, user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateLinkUsage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateLinkUsage(_), error: %v", err)
	}

	return
}
