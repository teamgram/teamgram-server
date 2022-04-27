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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UserNotifySettingsDAO struct {
	db *sqlx.DB
}

func NewUserNotifySettingsDAO(db *sqlx.DB) *UserNotifySettingsDAO {
	return &UserNotifySettingsDAO{db}
}

// InsertOrUpdate
// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserNotifySettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0"
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
// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserNotifySettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0"
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

// SelectAll
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) SelectAll(ctx context.Context, user_id int64) (rList []dataobject.UserNotifySettingsDO, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and deleted = 0"
		values []dataobject.UserNotifySettingsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllWithCB
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) SelectAllWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.UserNotifySettingsDO)) (rList []dataobject.UserNotifySettingsDO, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and deleted = 0"
		values []dataobject.UserNotifySettingsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAll(_), error: %v", err)
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

// Select
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) Select(ctx context.Context, user_id int64, peer_type int32, peer_id int64) (rValue *dataobject.UserNotifySettingsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		do    = &dataobject.UserNotifySettingsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, user_id, peer_type, peer_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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

// DeleteAll
// update user_notify_settings set deleted = 1 where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) DeleteAll(ctx context.Context, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_notify_settings set deleted = 1 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteAll(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteAll(_), error: %v", err)
	}

	return
}

// update user_notify_settings set deleted = 1 where user_id = :user_id
// DeleteAllTx
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) DeleteAllTx(tx *sqlx.Tx, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_notify_settings set deleted = 1 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteAll(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteAll(_), error: %v", err)
	}

	return
}
