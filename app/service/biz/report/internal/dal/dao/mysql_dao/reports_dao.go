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
	"github.com/teamgram/teamgram-server/app/service/biz/report/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type ReportsDAO struct {
	db *sqlx.DB
}

func NewReportsDAO(db *sqlx.DB) *ReportsDAO {
	return &ReportsDAO{db}
}

// Insert
// insert into reports(user_id, report_type, peer_type, peer_id, profile_photo_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :profile_photo_id, :message_sender_user_id, :message_id, :reason, :text)
// TODO(@benqi): sqlmap
func (dao *ReportsDAO) Insert(ctx context.Context, do *dataobject.ReportsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into reports(user_id, report_type, peer_type, peer_id, profile_photo_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :profile_photo_id, :message_sender_user_id, :message_id, :reason, :text)"
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
// insert into reports(user_id, report_type, peer_type, peer_id, profile_photo_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :profile_photo_id, :message_sender_user_id, :message_id, :reason, :text)
// TODO(@benqi): sqlmap
func (dao *ReportsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ReportsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into reports(user_id, report_type, peer_type, peer_id, profile_photo_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :profile_photo_id, :message_sender_user_id, :message_id, :reason, :text)"
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
// insert into reports(user_id, report_type, peer_type, peer_id, profile_photo_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :profile_photo_id, :message_sender_user_id, :message_id, :reason, :text)
// TODO(@benqi): sqlmap
func (dao *ReportsDAO) InsertBulk(ctx context.Context, doList []*dataobject.ReportsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into reports(user_id, report_type, peer_type, peer_id, profile_photo_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :profile_photo_id, :message_sender_user_id, :message_id, :reason, :text)"
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
// insert into reports(user_id, report_type, peer_type, peer_id, profile_photo_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :profile_photo_id, :message_sender_user_id, :message_id, :reason, :text)
// TODO(@benqi): sqlmap
func (dao *ReportsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.ReportsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into reports(user_id, report_type, peer_type, peer_id, profile_photo_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :profile_photo_id, :message_sender_user_id, :message_id, :reason, :text)"
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
