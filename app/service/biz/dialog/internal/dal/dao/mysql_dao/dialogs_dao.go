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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type DialogsDAO struct {
	db *sqlx.DB
}

func NewDialogsDAO(db *sqlx.DB) *DialogsDAO {
	return &DialogsDAO{db}
}

// InsertIgnore
// insert ignore into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :read_inbox_max_id, :read_outbox_max_id, :unread_count, :unread_mark, :draft_message_data, :date2)
func (dao *DialogsDAO) InsertIgnore(ctx context.Context, do *dataobject.DialogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :read_inbox_max_id, :read_outbox_max_id, :unread_count, :unread_mark, :draft_message_data, :date2)"
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
// insert ignore into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :read_inbox_max_id, :read_outbox_max_id, :unread_count, :unread_mark, :draft_message_data, :date2)
func (dao *DialogsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.DialogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :read_inbox_max_id, :read_outbox_max_id, :unread_count, :unread_mark, :draft_message_data, :date2)"
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

// InsertOrUpdate
// insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, unread_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :unread_count, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), date2 = values(date2)
func (dao *DialogsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.DialogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, unread_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :unread_count, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), date2 = values(date2)"
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
// insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, unread_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :unread_count, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), date2 = values(date2)
func (dao *DialogsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DialogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, unread_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :unread_count, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), date2 = values(date2)"
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

// InsertOrUpdateDialog
// insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, peer_dialog_id, read_inbox_max_id, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :peer_dialog_id, :read_inbox_max_id, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), read_inbox_max_id = values(read_inbox_max_id), draft_message_data = values(draft_message_data), date2 = values(date2), deleted = 0
func (dao *DialogsDAO) InsertOrUpdateDialog(ctx context.Context, do *dataobject.DialogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, peer_dialog_id, read_inbox_max_id, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :peer_dialog_id, :read_inbox_max_id, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), read_inbox_max_id = values(read_inbox_max_id), draft_message_data = values(draft_message_data), date2 = values(date2), deleted = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdateDialog(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdateDialog(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdateDialog(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateDialogTx
// insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, peer_dialog_id, read_inbox_max_id, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :peer_dialog_id, :read_inbox_max_id, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), read_inbox_max_id = values(read_inbox_max_id), draft_message_data = values(draft_message_data), date2 = values(date2), deleted = 0
func (dao *DialogsDAO) InsertOrUpdateDialogTx(tx *sqlx.Tx, do *dataobject.DialogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, peer_dialog_id, read_inbox_max_id, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :peer_dialog_id, :read_inbox_max_id, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), read_inbox_max_id = values(read_inbox_max_id), draft_message_data = values(draft_message_data), date2 = values(date2), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdateDialog(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdateDialog(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdateDialog(%v)_error: %v", do, err)
	}

	return
}

// UpdateOutboxDialog
// update dialogs set unread_count = 0, deleted = 0, top_message = :top_message, date2 = :date2, unread_mark = 0, draft_message_data = 'null' where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) UpdateOutboxDialog(ctx context.Context, top_message int32, date2 int64, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set unread_count = 0, deleted = 0, top_message = ?, date2 = ?, unread_mark = 0, draft_message_data = 'null' where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, top_message, date2, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateOutboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateOutboxDialog(_), error: %v", err)
	}

	return
}

// UpdateOutboxDialogTx
// update dialogs set unread_count = 0, deleted = 0, top_message = :top_message, date2 = :date2, unread_mark = 0, draft_message_data = 'null' where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) UpdateOutboxDialogTx(tx *sqlx.Tx, top_message int32, date2 int64, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set unread_count = 0, deleted = 0, top_message = ?, date2 = ?, unread_mark = 0, draft_message_data = 'null' where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, top_message, date2, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateOutboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateOutboxDialog(_), error: %v", err)
	}

	return
}

// UpdateInboxDialog
// update dialogs set unread_count = unread_count + 1, deleted = 0, %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) UpdateInboxDialog(ctx context.Context, cMap map[string]interface{}, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update dialogs set unread_count = unread_count + 1, deleted = 0, %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_type)
	aValues = append(aValues, peer_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateInboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateInboxDialog(_), error: %v", err)
	}

	return
}

// UpdateInboxDialogTx
// update dialogs set unread_count = unread_count + 1, deleted = 0, %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) UpdateInboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update dialogs set unread_count = unread_count + 1, deleted = 0, %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_type)
	aValues = append(aValues, peer_id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateInboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateInboxDialog(_), error: %v", err)
	}

	return
}

// SelectPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned > 0 and deleted = 0
func (dao *DialogsDAO) SelectPinnedDialogs(ctx context.Context, user_id int64) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = 0 and pinned > 0 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPinnedDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned > 0 and deleted = 0
func (dao *DialogsDAO) SelectPinnedDialogsWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.DialogsDO)) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = 0 and pinned > 0 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
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

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectFolderPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned > 0 and deleted = 0
func (dao *DialogsDAO) SelectFolderPinnedDialogs(ctx context.Context, user_id int64) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned > 0 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectFolderPinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectFolderPinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectFolderPinnedDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned > 0 and deleted = 0
func (dao *DialogsDAO) SelectFolderPinnedDialogsWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.DialogsDO)) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned > 0 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectFolderPinnedDialogs(_), error: %v", err)
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

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectFolderPinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPeerDialogList
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and peer_dialog_id in (:idList) and deleted = 0
func (dao *DialogsDAO) SelectPeerDialogList(ctx context.Context, user_id int64, idList []int64) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and peer_dialog_id in (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)

	if len(idList) == 0 {
		rList = []dataobject.DialogsDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectPeerDialogList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPeerDialogList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPeerDialogList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPeerDialogListWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and peer_dialog_id in (:idList) and deleted = 0
func (dao *DialogsDAO) SelectPeerDialogListWithCB(ctx context.Context, user_id int64, idList []int64, cb func(i int, v *dataobject.DialogsDO)) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and peer_dialog_id in (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)

	if len(idList) == 0 {
		rList = []dataobject.DialogsDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectPeerDialogList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPeerDialogList(_), error: %v", err)
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

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPeerDialogList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialog
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
func (dao *DialogsDAO) SelectDialog(ctx context.Context, user_id int64, peer_type int32, peer_id int64) (rValue *dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialog(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.DialogsDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialog(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = :folder_id and deleted = 0
func (dao *DialogsDAO) SelectDialogs(ctx context.Context, user_id int64, folder_id int32) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, folder_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = :folder_id and deleted = 0
func (dao *DialogsDAO) SelectDialogsWithCB(ctx context.Context, user_id int64, folder_id int32, cb func(i int, v *dataobject.DialogsDO)) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, folder_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogs(_), error: %v", err)
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

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectExcludePinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned = 0 and deleted = 0
func (dao *DialogsDAO) SelectExcludePinnedDialogs(ctx context.Context, user_id int64) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = 0 and pinned = 0 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludePinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectExcludePinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectExcludePinnedDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned = 0 and deleted = 0
func (dao *DialogsDAO) SelectExcludePinnedDialogsWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.DialogsDO)) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = 0 and pinned = 0 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludePinnedDialogs(_), error: %v", err)
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

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectExcludePinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectExcludeFolderPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned = 0 and deleted = 0
func (dao *DialogsDAO) SelectExcludeFolderPinnedDialogs(ctx context.Context, user_id int64) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned = 0 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludeFolderPinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectExcludeFolderPinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectExcludeFolderPinnedDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned = 0 and deleted = 0
func (dao *DialogsDAO) SelectExcludeFolderPinnedDialogsWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.DialogsDO)) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned = 0 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludeFolderPinnedDialogs(_), error: %v", err)
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

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectExcludeFolderPinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// UpdateReadInboxMaxId
// update dialogs set unread_count = 0, unread_mark = 0, read_inbox_max_id = :read_inbox_max_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DialogsDAO) UpdateReadInboxMaxId(ctx context.Context, read_inbox_max_id int32, user_id int64, peer_dialog_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_inbox_max_id, user_id, peer_dialog_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateReadInboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateReadInboxMaxId(_), error: %v", err)
	}

	return
}

// UpdateReadInboxMaxIdTx
// update dialogs set unread_count = 0, unread_mark = 0, read_inbox_max_id = :read_inbox_max_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DialogsDAO) UpdateReadInboxMaxIdTx(tx *sqlx.Tx, read_inbox_max_id int32, user_id int64, peer_dialog_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_inbox_max_id, user_id, peer_dialog_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateReadInboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateReadInboxMaxId(_), error: %v", err)
	}

	return
}

// UpdateReadOutboxMaxId
// update dialogs set read_outbox_max_id = :read_outbox_max_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DialogsDAO) UpdateReadOutboxMaxId(ctx context.Context, read_outbox_max_id int32, user_id int64, peer_dialog_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set read_outbox_max_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_outbox_max_id, user_id, peer_dialog_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateReadOutboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateReadOutboxMaxId(_), error: %v", err)
	}

	return
}

// UpdateReadOutboxMaxIdTx
// update dialogs set read_outbox_max_id = :read_outbox_max_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DialogsDAO) UpdateReadOutboxMaxIdTx(tx *sqlx.Tx, read_outbox_max_id int32, user_id int64, peer_dialog_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set read_outbox_max_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_outbox_max_id, user_id, peer_dialog_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateReadOutboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateReadOutboxMaxId(_), error: %v", err)
	}

	return
}

// UpdateTopMessage
// update dialogs set top_message = :top_message where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DialogsDAO) UpdateTopMessage(ctx context.Context, top_message int32, user_id int64, peer_dialog_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set top_message = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, top_message, user_id, peer_dialog_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateTopMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateTopMessage(_), error: %v", err)
	}

	return
}

// UpdateTopMessageTx
// update dialogs set top_message = :top_message where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DialogsDAO) UpdateTopMessageTx(tx *sqlx.Tx, top_message int32, user_id int64, peer_dialog_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set top_message = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, top_message, user_id, peer_dialog_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateTopMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateTopMessage(_), error: %v", err)
	}

	return
}

// UpdatePinnedMsgId
// update dialogs set pinned_msg_id = :pinned_msg_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DialogsDAO) UpdatePinnedMsgId(ctx context.Context, pinned_msg_id int32, user_id int64, peer_dialog_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set pinned_msg_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, pinned_msg_id, user_id, peer_dialog_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePinnedMsgId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePinnedMsgId(_), error: %v", err)
	}

	return
}

// UpdatePinnedMsgIdTx
// update dialogs set pinned_msg_id = :pinned_msg_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (dao *DialogsDAO) UpdatePinnedMsgIdTx(tx *sqlx.Tx, pinned_msg_id int32, user_id int64, peer_dialog_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set pinned_msg_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, pinned_msg_id, user_id, peer_dialog_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePinnedMsgId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePinnedMsgId(_), error: %v", err)
	}

	return
}

// Delete
// delete from dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) Delete(ctx context.Context, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from dialogs where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, peer_type, peer_id)

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
// delete from dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) DeleteTx(tx *sqlx.Tx, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from dialogs where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, peer_type, peer_id)

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

// SelectDialogsByGTReadInboxMaxId
// select user_id from dialogs where peer_type = :peer_type and peer_id = :peer_id and read_inbox_max_id >= :read_inbox_max_id and user_id != :user_id
func (dao *DialogsDAO) SelectDialogsByGTReadInboxMaxId(ctx context.Context, peer_type int32, peer_id int64, read_inbox_max_id int32, user_id int64) (rList []int64, err error) {
	var query = "select user_id from dialogs where peer_type = ? and peer_id = ? and read_inbox_max_id >= ? and user_id != ?"
	err = dao.db.Select(ctx, &rList, query, peer_type, peer_id, read_inbox_max_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectDialogsByGTReadInboxMaxId(_), error: %v", err)
	}

	return
}

// SelectDialogsByGTReadInboxMaxIdWithCB
// select user_id from dialogs where peer_type = :peer_type and peer_id = :peer_id and read_inbox_max_id >= :read_inbox_max_id and user_id != :user_id
func (dao *DialogsDAO) SelectDialogsByGTReadInboxMaxIdWithCB(ctx context.Context, peer_type int32, peer_id int64, read_inbox_max_id int32, user_id int64, cb func(i int, v int64)) (rList []int64, err error) {
	var query = "select user_id from dialogs where peer_type = ? and peer_id = ? and read_inbox_max_id >= ? and user_id != ?"
	err = dao.db.Select(ctx, &rList, query, peer_type, peer_id, read_inbox_max_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectDialogsByGTReadInboxMaxId(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// UpdateCustomMap
// update dialogs set %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update dialogs set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_type)
	aValues = append(aValues, peer_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}

// UpdateCustomMapTx
// update dialogs set %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update dialogs set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_type)
	aValues = append(aValues, peer_id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}

// SaveDraft
// update dialogs set draft_type = :draft_type, draft_message_data = :draft_message_data where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) SaveDraft(ctx context.Context, draft_type int32, draft_message_data string, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set draft_type = ?, draft_message_data = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, draft_type, draft_message_data, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in SaveDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in SaveDraft(_), error: %v", err)
	}

	return
}

// SaveDraftTx
// update dialogs set draft_type = :draft_type, draft_message_data = :draft_message_data where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) SaveDraftTx(tx *sqlx.Tx, draft_type int32, draft_message_data string, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set draft_type = ?, draft_message_data = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, draft_type, draft_message_data, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in SaveDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in SaveDraft(_), error: %v", err)
	}

	return
}

// SelectAllDrafts
// select user_id, peer_id, draft_message_data from dialogs where user_id = :user_id and draft_type > 0
func (dao *DialogsDAO) SelectAllDrafts(ctx context.Context, user_id int64) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select user_id, peer_id, draft_message_data from dialogs where user_id = ? and draft_type > 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllDrafts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectAllDrafts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectAllDraftsWithCB
// select user_id, peer_id, draft_message_data from dialogs where user_id = :user_id and draft_type > 0
func (dao *DialogsDAO) SelectAllDraftsWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.DialogsDO)) (rList []dataobject.DialogsDO, err error) {
	var (
		query = "select user_id, peer_id, draft_message_data from dialogs where user_id = ? and draft_type > 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllDrafts(_), error: %v", err)
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

	var values []dataobject.DialogsDO
	for rows.Next() {
		v := dataobject.DialogsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectAllDrafts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// ClearAllDrafts
// update dialogs set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and draft_type = 2
func (dao *DialogsDAO) ClearAllDrafts(ctx context.Context, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set draft_type = 0, draft_message_data = 'null' where user_id = ? and draft_type = 2"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in ClearAllDrafts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in ClearAllDrafts(_), error: %v", err)
	}

	return
}

// ClearAllDraftsTx
// update dialogs set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and draft_type = 2
func (dao *DialogsDAO) ClearAllDraftsTx(tx *sqlx.Tx, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set draft_type = 0, draft_message_data = 'null' where user_id = ? and draft_type = 2"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in ClearAllDrafts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in ClearAllDrafts(_), error: %v", err)
	}

	return
}

// UpdatePeerFolderId
// update dialogs set folder_id = :folder_id where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) UpdatePeerFolderId(ctx context.Context, folder_id int32, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set folder_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, folder_id, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePeerFolderId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePeerFolderId(_), error: %v", err)
	}

	return
}

// UpdatePeerFolderIdTx
// update dialogs set folder_id = :folder_id where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (dao *DialogsDAO) UpdatePeerFolderIdTx(tx *sqlx.Tx, folder_id int32, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set folder_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, folder_id, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePeerFolderId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePeerFolderId(_), error: %v", err)
	}

	return
}

// UpdatePeerDialogListFolderId
// update dialogs set folder_id = :folder_id where user_id = :user_id and peer_dialog_id in (:idList)
func (dao *DialogsDAO) UpdatePeerDialogListFolderId(ctx context.Context, folder_id int32, user_id int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set folder_id = ? where user_id = ? and peer_dialog_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, folder_id, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in UpdatePeerDialogListFolderId(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePeerDialogListFolderId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePeerDialogListFolderId(_), error: %v", err)
	}

	return
}

// UpdatePeerDialogListFolderIdTx
// update dialogs set folder_id = :folder_id where user_id = :user_id and peer_dialog_id in (:idList)
func (dao *DialogsDAO) UpdatePeerDialogListFolderIdTx(tx *sqlx.Tx, folder_id int32, user_id int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set folder_id = ? where user_id = ? and peer_dialog_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, folder_id, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in UpdatePeerDialogListFolderId(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePeerDialogListFolderId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePeerDialogListFolderId(_), error: %v", err)
	}

	return
}

// UpdatePeerDialogListPinned
// update dialogs set pinned = :pinned where user_id = :user_id and folder_id = 0 and peer_dialog_id in (:idList)
func (dao *DialogsDAO) UpdatePeerDialogListPinned(ctx context.Context, pinned int64, user_id int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set pinned = ? where user_id = ? and folder_id = 0 and peer_dialog_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, pinned, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in UpdatePeerDialogListPinned(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePeerDialogListPinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePeerDialogListPinned(_), error: %v", err)
	}

	return
}

// UpdatePeerDialogListPinnedTx
// update dialogs set pinned = :pinned where user_id = :user_id and folder_id = 0 and peer_dialog_id in (:idList)
func (dao *DialogsDAO) UpdatePeerDialogListPinnedTx(tx *sqlx.Tx, pinned int64, user_id int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set pinned = ? where user_id = ? and folder_id = 0 and peer_dialog_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, pinned, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in UpdatePeerDialogListPinned(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePeerDialogListPinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePeerDialogListPinned(_), error: %v", err)
	}

	return
}

// UpdateFolderPeerDialogListPinned
// update dialogs set folder_pinned = :folder_pinned where user_id = :user_id and folder_id = 1 and peer_dialog_id in (:idList)
func (dao *DialogsDAO) UpdateFolderPeerDialogListPinned(ctx context.Context, folder_pinned int64, user_id int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set folder_pinned = ? where user_id = ? and folder_id = 1 and peer_dialog_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, folder_pinned, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in UpdateFolderPeerDialogListPinned(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateFolderPeerDialogListPinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateFolderPeerDialogListPinned(_), error: %v", err)
	}

	return
}

// UpdateFolderPeerDialogListPinnedTx
// update dialogs set folder_pinned = :folder_pinned where user_id = :user_id and folder_id = 1 and peer_dialog_id in (:idList)
func (dao *DialogsDAO) UpdateFolderPeerDialogListPinnedTx(tx *sqlx.Tx, folder_pinned int64, user_id int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set folder_pinned = ? where user_id = ? and folder_id = 1 and peer_dialog_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, folder_pinned, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in UpdateFolderPeerDialogListPinned(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateFolderPeerDialogListPinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateFolderPeerDialogListPinned(_), error: %v", err)
	}

	return
}
