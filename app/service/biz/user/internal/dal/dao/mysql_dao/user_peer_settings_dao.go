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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UserPeerSettingsDAO struct {
	db *sqlx.DB
}

func NewUserPeerSettingsDAO(db *sqlx.DB) *UserPeerSettingsDAO {
	return &UserPeerSettingsDAO{db}
}

// InsertIgnore
// insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance)
// TODO(@benqi): sqlmap
func (dao *UserPeerSettingsDAO) InsertIgnore(ctx context.Context, do *dataobject.UserPeerSettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance)"
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
// insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance)
// TODO(@benqi): sqlmap
func (dao *UserPeerSettingsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.UserPeerSettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance)"
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

// Select
// select user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance from user_peer_settings where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and hide = 0
// TODO(@benqi): sqlmap
func (dao *UserPeerSettingsDAO) Select(ctx context.Context, user_id int64, peer_type int32, peer_id int64) (rValue *dataobject.UserPeerSettingsDO, err error) {
	var (
		query = "select user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance from user_peer_settings where user_id = ? and peer_type = ? and peer_id = ? and hide = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserPeerSettingsDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// Update
// update user_peer_settings set %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserPeerSettingsDAO) Update(ctx context.Context, cMap map[string]interface{}, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update user_peer_settings set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_type)
	aValues = append(aValues, peer_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

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

// UpdateTx
// update user_peer_settings set %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserPeerSettingsDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update user_peer_settings set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_type)
	aValues = append(aValues, peer_id)

	rResult, err = tx.Exec(query, aValues...)

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

// Delete
// update user_peer_settings set hide = 1 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserPeerSettingsDAO) Delete(ctx context.Context, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_settings set hide = 1 where user_id = ? and peer_type = ? and peer_id = ?"
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

// update user_peer_settings set hide = 1 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// DeleteTx
// TODO(@benqi): sqlmap
func (dao *UserPeerSettingsDAO) DeleteTx(tx *sqlx.Tx, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_settings set hide = 1 where user_id = ? and peer_type = ? and peer_id = ?"
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
