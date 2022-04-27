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

type UserGlobalPrivacySettingsDAO struct {
	db *sqlx.DB
}

func NewUserGlobalPrivacySettingsDAO(db *sqlx.DB) *UserGlobalPrivacySettingsDAO {
	return &UserGlobalPrivacySettingsDAO{db}
}

// InsertOrUpdate
// insert into user_global_privacy_settings(user_id, archive_and_mute_new_noncontact_peers) values (:user_id, :archive_and_mute_new_noncontact_peers) on duplicate key update archive_and_mute_new_noncontact_peers = values(archive_and_mute_new_noncontact_peers)
// TODO(@benqi): sqlmap
func (dao *UserGlobalPrivacySettingsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserGlobalPrivacySettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_global_privacy_settings(user_id, archive_and_mute_new_noncontact_peers) values (:user_id, :archive_and_mute_new_noncontact_peers) on duplicate key update archive_and_mute_new_noncontact_peers = values(archive_and_mute_new_noncontact_peers)"
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
// insert into user_global_privacy_settings(user_id, archive_and_mute_new_noncontact_peers) values (:user_id, :archive_and_mute_new_noncontact_peers) on duplicate key update archive_and_mute_new_noncontact_peers = values(archive_and_mute_new_noncontact_peers)
// TODO(@benqi): sqlmap
func (dao *UserGlobalPrivacySettingsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserGlobalPrivacySettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_global_privacy_settings(user_id, archive_and_mute_new_noncontact_peers) values (:user_id, :archive_and_mute_new_noncontact_peers) on duplicate key update archive_and_mute_new_noncontact_peers = values(archive_and_mute_new_noncontact_peers)"
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
// select id, user_id, archive_and_mute_new_noncontact_peers from user_global_privacy_settings where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserGlobalPrivacySettingsDAO) Select(ctx context.Context, user_id int64) (rValue *dataobject.UserGlobalPrivacySettingsDO, err error) {
	var (
		query = "select id, user_id, archive_and_mute_new_noncontact_peers from user_global_privacy_settings where user_id = ?"
		do    = &dataobject.UserGlobalPrivacySettingsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, user_id)

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
