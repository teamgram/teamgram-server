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
	bizUserPeerSettingsModel interface {
		InsertOrUpdate(ctx context.Context, data *UserPeerSettings) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UserPeerSettings) (lastInsertId, rowsAffected int64, err error)

		Select(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserPeerSettings, error)

		Update(ctx context.Context, cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
		UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)

		Delete(ctx context.Context, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
		DeleteTx(tx *sqlx.Tx, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance) on duplicate key update report_spam = values(report_spam), add_contact = values(add_contact), block_contact = values(block_contact), share_contact = values(share_contact), need_contacts_exception = values(need_contacts_exception), report_geo = values(report_geo), autoarchived = values(autoarchived), geo_distance = values(geo_distance), hide = 0
func (m *defaultUserPeerSettingsModel) InsertOrUpdate(ctx context.Context, data *UserPeerSettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance) on duplicate key update report_spam = values(report_spam), add_contact = values(add_contact), block_contact = values(block_contact), share_contact = values(share_contact), need_contacts_exception = values(need_contacts_exception), report_geo = values(report_geo), autoarchived = values(autoarchived), geo_distance = values(geo_distance), hide = 0"
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
// insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance) on duplicate key update report_spam = values(report_spam), add_contact = values(add_contact), block_contact = values(block_contact), share_contact = values(share_contact), need_contacts_exception = values(need_contacts_exception), report_geo = values(report_geo), autoarchived = values(autoarchived), geo_distance = values(geo_distance), hide = 0
func (m *defaultUserPeerSettingsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UserPeerSettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance) on duplicate key update report_spam = values(report_spam), add_contact = values(add_contact), block_contact = values(block_contact), share_contact = values(share_contact), need_contacts_exception = values(need_contacts_exception), report_geo = values(report_geo), autoarchived = values(autoarchived), geo_distance = values(geo_distance), hide = 0"
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
// select user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance from user_peer_settings where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and hide = 0
func (m *defaultUserPeerSettingsModel) Select(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *UserPeerSettings, err error) {
	var (
		query = "select user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance from user_peer_settings where user_id = ? and peer_type = ? and peer_id = ? and hide = 0"
		do    = &UserPeerSettings{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

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

// Update
// update user_peer_settings set %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserPeerSettingsModel) Update(ctx context.Context, cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
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

	aValues = append(aValues, userId)
	aValues = append(aValues, peerType)
	aValues = append(aValues, peerId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

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
func (m *defaultUserPeerSettingsModel) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
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

	aValues = append(aValues, userId)
	aValues = append(aValues, peerType)
	aValues = append(aValues, peerId)

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
func (m *defaultUserPeerSettingsModel) Delete(ctx context.Context, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_settings set hide = 1 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, peerType, peerId)

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
// update user_peer_settings set hide = 1 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserPeerSettingsModel) DeleteTx(tx *sqlx.Tx, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_settings set hide = 1 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, peerType, peerId)

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
