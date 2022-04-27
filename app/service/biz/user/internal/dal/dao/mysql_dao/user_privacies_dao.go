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

type UserPrivaciesDAO struct {
	db *sqlx.DB
}

func NewUserPrivaciesDAO(db *sqlx.DB) *UserPrivaciesDAO {
	return &UserPrivaciesDAO{db}
}

// InsertOrUpdate
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"
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
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"
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

// InsertBulk
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) InsertBulk(ctx context.Context, doList []*dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)"
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
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)"
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

// SelectPrivacy
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type = :key_type
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) SelectPrivacy(ctx context.Context, user_id int64, key_type int32) (rValue *dataobject.UserPrivaciesDO, err error) {
	var (
		query = "select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type = ?"
		do    = &dataobject.UserPrivaciesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, user_id, key_type)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in SelectPrivacy(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectPrivacyList
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type in (:keyList)
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) SelectPrivacyList(ctx context.Context, user_id int64, keyList []int32) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type in (?)"
		a      []interface{}
		values []dataobject.UserPrivaciesDO
	)

	if len(keyList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, keyList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectPrivacyList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPrivacyListWithCB
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type in (:keyList)
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) SelectPrivacyListWithCB(ctx context.Context, user_id int64, keyList []int32, cb func(i int, v *dataobject.UserPrivaciesDO)) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type in (?)"
		a      []interface{}
		values []dataobject.UserPrivaciesDO
	)

	if len(keyList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, keyList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectPrivacyList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyList(_), error: %v", err)
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

// SelectUsersPrivacyList
// select id, user_id, key_type, rules from user_privacies where user_id in (:idList) and key_type in (:keyList)
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) SelectUsersPrivacyList(ctx context.Context, idList []int32, keyList []int32) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id in (?) and key_type in (?)"
		a      []interface{}
		values []dataobject.UserPrivaciesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}
	if len(keyList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList, keyList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectUsersPrivacyList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersPrivacyList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUsersPrivacyListWithCB
// select id, user_id, key_type, rules from user_privacies where user_id in (:idList) and key_type in (:keyList)
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) SelectUsersPrivacyListWithCB(ctx context.Context, idList []int32, keyList []int32, cb func(i int, v *dataobject.UserPrivaciesDO)) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id in (?) and key_type in (?)"
		a      []interface{}
		values []dataobject.UserPrivaciesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}
	if len(keyList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList, keyList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectUsersPrivacyList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersPrivacyList(_), error: %v", err)
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

// SelectPrivacyAll
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) SelectPrivacyAll(ctx context.Context, user_id int64) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id = ?"
		values []dataobject.UserPrivaciesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyAll(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPrivacyAllWithCB
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserPrivaciesDAO) SelectPrivacyAllWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.UserPrivaciesDO)) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id = ?"
		values []dataobject.UserPrivaciesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyAll(_), error: %v", err)
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
