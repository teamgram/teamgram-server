/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/username/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type UsernameDAO struct {
	db *sqlx.DB
}

func NewUsernameDAO(db *sqlx.DB) *UsernameDAO {
	return &UsernameDAO{
		db: db,
	}
}

// Insert
// insert into username(peer_type, peer_id, username, deleted) values (:peer_type, :peer_id, :username, 0)
func (dao *UsernameDAO) Insert(ctx context.Context, do *dataobject.UsernameDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into username(peer_type, peer_id, username, deleted) values (:peer_type, :peer_id, :username, 0)"
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
// insert into username(peer_type, peer_id, username, deleted) values (:peer_type, :peer_id, :username, 0)
func (dao *UsernameDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UsernameDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into username(peer_type, peer_id, username, deleted) values (:peer_type, :peer_id, :username, 0)"
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

// SelectList
// select username, peer_type, peer_id from username where username in (:nameList)
func (dao *UsernameDAO) SelectList(ctx context.Context, nameList []string) (rList []dataobject.UsernameDO, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id from username where username in (%s)", sqlx.InStringList(nameList))
		values []dataobject.UsernameDO
	)
	if len(nameList) == 0 {
		rList = []dataobject.UsernameDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select username, peer_type, peer_id from username where username in (:nameList)
func (dao *UsernameDAO) SelectListWithCB(ctx context.Context, nameList []string, cb func(sz, i int, v *dataobject.UsernameDO)) (rList []dataobject.UsernameDO, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id from username where username in (%s)", sqlx.InStringList(nameList))
		values []dataobject.UsernameDO
	)
	if len(nameList) == 0 {
		rList = []dataobject.UsernameDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

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
// select username, peer_type, peer_id, deleted from username where username = :username
func (dao *UsernameDAO) SelectByUsername(ctx context.Context, username string) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select username, peer_type, peer_id, deleted from username where username = ?"
		do    = &dataobject.UsernameDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, username)

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
func (dao *UsernameDAO) Update(ctx context.Context, cMap map[string]interface{}, username string) (rowsAffected int64, err error) {
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
// update username set %s where username = :username
func (dao *UsernameDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, username string) (rowsAffected int64, err error) {
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
func (dao *UsernameDAO) Delete(ctx context.Context, username string) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where username = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, username)

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
func (dao *UsernameDAO) DeleteTx(tx *sqlx.Tx, username string) (rowsAffected int64, err error) {
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
// delete from username where peer_type = :peer_type and peer_id = :peer_id
func (dao *UsernameDAO) Delete2(ctx context.Context, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, peerType, peerId)

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
// delete from username where peer_type = :peer_type and peer_id = :peer_id
func (dao *UsernameDAO) Delete2Tx(tx *sqlx.Tx, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = ? and peer_id = ?"
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

// SelectByPeer
// select peer_type, peer_id, username from username where peer_type = :peer_type and peer_id = :peer_id
func (dao *UsernameDAO) SelectByPeer(ctx context.Context, peerType int32, peerId int64) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = ? and peer_id = ?"
		do    = &dataobject.UsernameDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, peerType, peerId)

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
// select peer_type, peer_id, username from username where peer_type = 2 and peer_id = :peer_id
func (dao *UsernameDAO) SelectByUserId(ctx context.Context, peerId int64) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = 2 and peer_id = ?"
		do    = &dataobject.UsernameDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, peerId)

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

// SelectByChannelId
// select peer_type, peer_id, username from username where peer_type = 4 and peer_id = :peer_id
func (dao *UsernameDAO) SelectByChannelId(ctx context.Context, peerId int64) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = 4 and peer_id = ?"
		do    = &dataobject.UsernameDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, peerId)

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

// UpdateUsername
// update username set username = :username where peer_type = :peer_type and peer_id = :peer_id
func (dao *UsernameDAO) UpdateUsername(ctx context.Context, username string, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, username, peerType, peerId)

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
// update username set username = :username where peer_type = :peer_type and peer_id = :peer_id
func (dao *UsernameDAO) UpdateUsernameTx(tx *sqlx.Tx, username string, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ?"
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
func (dao *UsernameDAO) SearchByQueryNotIdList(ctx context.Context, q2 string, idList []int64, limit int32) (rList []dataobject.UsernameDO, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id from username where username like ? and peer_id not in (%s) limit ?", sqlx.InInt64List(idList))
		values []dataobject.UsernameDO
	)

	if len(idList) == 0 {
		rList = []dataobject.UsernameDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchByQueryNotIdListWithCB
// select username, peer_type, peer_id from username where username like :q2 and peer_id not in (:id_list) limit :limit
func (dao *UsernameDAO) SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, idList []int64, limit int32, cb func(sz, i int, v *dataobject.UsernameDO)) (rList []dataobject.UsernameDO, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id from username where username like ? and peer_id not in (%s) limit ?", sqlx.InInt64List(idList))
		values []dataobject.UsernameDO
	)

	if len(idList) == 0 {
		rList = []dataobject.UsernameDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, q2, limit)

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
