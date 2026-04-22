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
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *logx.Logger

type (
	bizUsernameModel interface {
		Insert(ctx context.Context, data *Username) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *Username) (lastInsertId, rowsAffected int64, err error)

		SelectList(ctx context.Context, nameList []string) ([]Username, error)
		SelectListWithCB(ctx context.Context, nameList []string, cb func(sz, i int, v *Username)) ([]Username, error)

		SelectByUsername(ctx context.Context, username string) (*Username, error)

		Update(ctx context.Context, cMap map[string]interface{}, username string) (rowsAffected int64, err error)
		UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, username string) (rowsAffected int64, err error)

		Delete(ctx context.Context, username string) (rowsAffected int64, err error)
		DeleteTx(tx *sqlx.Tx, username string) (rowsAffected int64, err error)

		Delete2(ctx context.Context, peerType int32, peerId int64) (rowsAffected int64, err error)
		Delete2Tx(tx *sqlx.Tx, peerType int32, peerId int64) (rowsAffected int64, err error)

		DeleteByChannelId(ctx context.Context, peerId int64) (rowsAffected int64, err error)
		DeleteByChannelIdTx(tx *sqlx.Tx, peerId int64) (rowsAffected int64, err error)

		SelectByPeer(ctx context.Context, peerType int32, peerId int64) (*Username, error)

		SelectByUserId(ctx context.Context, peerId int64) (*Username, error)

		SelectListByUserId(ctx context.Context, peerId int64) ([]Username, error)
		SelectListByUserIdWithCB(ctx context.Context, peerId int64, cb func(sz, i int, v *Username)) ([]Username, error)

		SelectByChannelId(ctx context.Context, peerId int64) (*Username, error)

		SelectListByChannelId(ctx context.Context, peerId int64) ([]Username, error)
		SelectListByChannelIdWithCB(ctx context.Context, peerId int64, cb func(sz, i int, v *Username)) ([]Username, error)

		UpdateUsername(ctx context.Context, username string, peerType int32, peerId int64) (rowsAffected int64, err error)
		UpdateUsernameTx(tx *sqlx.Tx, username string, peerType int32, peerId int64) (rowsAffected int64, err error)

		SearchByQueryNotIdList(ctx context.Context, q2 string, idList []int64, limit int32) ([]Username, error)
		SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, idList []int64, limit int32, cb func(sz, i int, v *Username)) ([]Username, error)
	}
)

// Insert
// insert into username(username, peer_type, peer_id, editable, active, order2, deleted) values (:username, :peer_type, :peer_id, :editable, :active, :order2, 0)
func (m *defaultUsernameModel) Insert(ctx context.Context, data *Username) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into username(username, peer_type, peer_id, editable, active, order2, deleted) values (:username, :peer_type, :peer_id, :editable, :active, :order2, 0)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// InsertTx
// insert into username(username, peer_type, peer_id, editable, active, order2, deleted) values (:username, :peer_type, :peer_id, :editable, :active, :order2, 0)
func (m *defaultUsernameModel) InsertTx(tx *sqlx.Tx, data *Username) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into username(username, peer_type, peer_id, editable, active, order2, deleted) values (:username, :peer_type, :peer_id, :editable, :active, :order2, 0)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// SelectList
// select username, peer_type, peer_id, editable, active, order2 from username where username in (:nameList) and editable = 1
func (m *defaultUsernameModel) SelectList(ctx context.Context, nameList []string) (rList []Username, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id, editable, active, order2 from username where username in (%s) and editable = 1", sqlx.InStringList(nameList))
		values []Username
	)
	if len(nameList) == 0 {
		rList = []Username{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select username, peer_type, peer_id, editable, active, order2 from username where username in (:nameList) and editable = 1
func (m *defaultUsernameModel) SelectListWithCB(ctx context.Context, nameList []string, cb func(sz, i int, v *Username)) (rList []Username, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id, editable, active, order2 from username where username in (%s) and editable = 1", sqlx.InStringList(nameList))
		values []Username
	)
	if len(nameList) == 0 {
		rList = []Username{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
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

// SelectByUsername
// select username, peer_type, peer_id, editable, active, order2, deleted from username where username = :username
func (m *defaultUsernameModel) SelectByUsername(ctx context.Context, username string) (rValue *Username, err error) {
	var (
		query = "select username, peer_type, peer_id, editable, active, order2, deleted from username where username = ?"
		do    = &Username{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, username)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByUsername(_), error: %v", err)
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
// update username set %s where username = :username
func (m *defaultUsernameModel) Update(ctx context.Context, cMap map[string]interface{}, username string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update username set %s where username = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, username)

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
// update username set %s where username = :username
func (m *defaultUsernameModel) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, username string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update username set %s where username = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, username)

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
// delete from username where username = :username
func (m *defaultUsernameModel) Delete(ctx context.Context, username string) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where username = ?"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, username)

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
// delete from username where username = :username
func (m *defaultUsernameModel) DeleteTx(tx *sqlx.Tx, username string) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where username = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, username)

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

// Delete2
// delete from username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameModel) Delete2(ctx context.Context, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = ? and peer_id = ? and editable = 1"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, peerType, peerId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Delete2(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Delete2(_), error: %v", err)
	}

	return
}

// Delete2Tx
// delete from username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameModel) Delete2Tx(tx *sqlx.Tx, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = ? and peer_id = ? and editable = 1"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, peerType, peerId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Delete2(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Delete2(_), error: %v", err)
	}

	return
}

// DeleteByChannelId
// delete from username where peer_type = 2 and peer_id = :peer_id and editable = 0
func (m *defaultUsernameModel) DeleteByChannelId(ctx context.Context, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = 2 and peer_id = ? and editable = 0"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, peerId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteByChannelId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteByChannelId(_), error: %v", err)
	}

	return
}

// DeleteByChannelIdTx
// delete from username where peer_type = 2 and peer_id = :peer_id and editable = 0
func (m *defaultUsernameModel) DeleteByChannelIdTx(tx *sqlx.Tx, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = 2 and peer_id = ? and editable = 0"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, peerId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteByChannelId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteByChannelId(_), error: %v", err)
	}

	return
}

// SelectByPeer
// select username, peer_type, peer_id, editable, active, order2 from username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameModel) SelectByPeer(ctx context.Context, peerType int32, peerId int64) (rValue *Username, err error) {
	var (
		query = "select username, peer_type, peer_id, editable, active, order2 from username where peer_type = ? and peer_id = ? and editable = 1"
		do    = &Username{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, peerType, peerId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByPeer(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByUserId
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = :peer_id and editable = 1
func (m *defaultUsernameModel) SelectByUserId(ctx context.Context, peerId int64) (rValue *Username, err error) {
	var (
		query = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = ? and editable = 1"
		do    = &Username{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, peerId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByUserId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectListByUserId
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = :peer_id
func (m *defaultUsernameModel) SelectListByUserId(ctx context.Context, peerId int64) (rList []Username, err error) {
	var (
		query  = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = ?"
		values []Username
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, peerId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByUserId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByUserIdWithCB
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = :peer_id
func (m *defaultUsernameModel) SelectListByUserIdWithCB(ctx context.Context, peerId int64, cb func(sz, i int, v *Username)) (rList []Username, err error) {
	var (
		query  = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = ?"
		values []Username
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, peerId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByUserId(_), error: %v", err)
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

// SelectByChannelId
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = :peer_id and editable = 1
func (m *defaultUsernameModel) SelectByChannelId(ctx context.Context, peerId int64) (rValue *Username, err error) {
	var (
		query = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = ? and editable = 1"
		do    = &Username{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, peerId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByChannelId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectListByChannelId
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = :peer_id
func (m *defaultUsernameModel) SelectListByChannelId(ctx context.Context, peerId int64) (rList []Username, err error) {
	var (
		query  = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = ?"
		values []Username
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, peerId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChannelId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByChannelIdWithCB
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = :peer_id
func (m *defaultUsernameModel) SelectListByChannelIdWithCB(ctx context.Context, peerId int64, cb func(sz, i int, v *Username)) (rList []Username, err error) {
	var (
		query  = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = ?"
		values []Username
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, peerId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChannelId(_), error: %v", err)
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

// UpdateUsername
// update username set username = :username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameModel) UpdateUsername(ctx context.Context, username string, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ? and editable = 1"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, username, peerType, peerId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateUsername(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateUsername(_), error: %v", err)
	}

	return
}

// UpdateUsernameTx
// update username set username = :username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameModel) UpdateUsernameTx(tx *sqlx.Tx, username string, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ? and editable = 1"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, username, peerType, peerId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateUsername(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateUsername(_), error: %v", err)
	}

	return
}

// SearchByQueryNotIdList
// select username, peer_type, peer_id from username where username like :q2 and peer_id not in (:id_list) limit :limit
func (m *defaultUsernameModel) SearchByQueryNotIdList(ctx context.Context, q2 string, idList []int64, limit int32) (rList []Username, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id from username where username like ? and peer_id not in (%s) limit ?", sqlx.InInt64List(idList))
		values []Username
	)
	if len(idList) == 0 {
		rList = []Username{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchByQueryNotIdListWithCB
// select username, peer_type, peer_id from username where username like :q2 and peer_id not in (:id_list) limit :limit
func (m *defaultUsernameModel) SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, idList []int64, limit int32, cb func(sz, i int, v *Username)) (rList []Username, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id from username where username like ? and peer_id not in (%s) limit ?", sqlx.InInt64List(idList))
		values []Username
	)
	if len(idList) == 0 {
		rList = []Username{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
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
