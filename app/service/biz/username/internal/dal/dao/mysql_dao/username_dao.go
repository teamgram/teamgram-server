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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/username/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UsernameDAO struct {
	db *sqlx.DB
}

func NewUsernameDAO(db *sqlx.DB) *UsernameDAO {
	return &UsernameDAO{db}
}

// Insert
// insert into username(peer_type, peer_id, username, deleted) values (:peer_type, :peer_id, :username, 0)
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) SelectList(ctx context.Context, nameList []string) (rList []dataobject.UsernameDO, err error) {
	var (
		query  = "select username, peer_type, peer_id from username where username in (?)"
		a      []interface{}
		values []dataobject.UsernameDO
	)
	if len(nameList) == 0 {
		rList = []dataobject.UsernameDO{}
		return
	}

	query, a, err = sqlx.In(query, nameList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select username, peer_type, peer_id from username where username in (:nameList)
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) SelectListWithCB(ctx context.Context, nameList []string, cb func(i int, v *dataobject.UsernameDO)) (rList []dataobject.UsernameDO, err error) {
	var (
		query  = "select username, peer_type, peer_id from username where username in (?)"
		a      []interface{}
		values []dataobject.UsernameDO
	)
	if len(nameList) == 0 {
		rList = []dataobject.UsernameDO{}
		return
	}

	query, a, err = sqlx.In(query, nameList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

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

// SelectByUsername
// select username, peer_type, peer_id, deleted from username where username = :username
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) SelectByUsername(ctx context.Context, username string) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select username, peer_type, peer_id, deleted from username where username = ?"
		do    = &dataobject.UsernameDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, username)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) Delete2(ctx context.Context, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, peer_type, peer_id)

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
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) Delete2Tx(tx *sqlx.Tx, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, peer_type, peer_id)

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
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) SelectByPeer(ctx context.Context, peer_type int32, peer_id int64) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = ? and peer_id = ?"
		do    = &dataobject.UsernameDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, peer_type, peer_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) SelectByUserId(ctx context.Context, peer_id int64) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = 2 and peer_id = ?"
		do    = &dataobject.UsernameDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, peer_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) SelectByChannelId(ctx context.Context, peer_id int64) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = 4 and peer_id = ?"
		do    = &dataobject.UsernameDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, peer_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) UpdateUsername(ctx context.Context, username string, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, username, peer_type, peer_id)

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

// update username set username = :username where peer_type = :peer_type and peer_id = :peer_id
// UpdateUsernameTx
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) UpdateUsernameTx(tx *sqlx.Tx, username string, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, username, peer_type, peer_id)

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
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) SearchByQueryNotIdList(ctx context.Context, q2 string, id_list []int64, limit int32) (rList []dataobject.UsernameDO, err error) {
	var (
		query  = "select username, peer_type, peer_id from username where username like ? and peer_id not in (?) limit ?"
		a      []interface{}
		values []dataobject.UsernameDO
	)

	if len(id_list) == 0 {
		rList = []dataobject.UsernameDO{}
		return
	}

	query, a, err = sqlx.In(query, q2, id_list, limit)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SearchByQueryNotIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchByQueryNotIdListWithCB
// select username, peer_type, peer_id from username where username like :q2 and peer_id not in (:id_list) limit :limit
// TODO(@benqi): sqlmap
func (dao *UsernameDAO) SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, id_list []int64, limit int32, cb func(i int, v *dataobject.UsernameDO)) (rList []dataobject.UsernameDO, err error) {
	var (
		query  = "select username, peer_type, peer_id from username where username like ? and peer_id not in (?) limit ?"
		a      []interface{}
		values []dataobject.UsernameDO
	)

	if len(id_list) == 0 {
		rList = []dataobject.UsernameDO{}
		return
	}

	query, a, err = sqlx.In(query, q2, id_list, limit)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SearchByQueryNotIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
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
