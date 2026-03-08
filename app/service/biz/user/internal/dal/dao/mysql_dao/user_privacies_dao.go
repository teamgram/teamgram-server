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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPrivaciesDAO struct {
	db *sqlx.DB
}

func NewUserPrivaciesDAO(db *sqlx.DB) *UserPrivaciesDAO {
	return &UserPrivaciesDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)
func (dao *UserPrivaciesDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)
func (dao *UserPrivaciesDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// InsertBulk
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)
func (dao *UserPrivaciesDAO) InsertBulk(ctx context.Context, doList []*dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	if len(doList) == 0 {
		return
	}

	var (
		query string
		r     sql.Result
	)
	query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)"

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertBulk(%v), error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertBulk(%v), error: %v", doList, err)
	}

	return
}

// InsertBulkTx
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)
func (dao *UserPrivaciesDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	if len(doList) == 0 {
		return
	}

	var (
		query string
		r     sql.Result
	)
	query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)"

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertBulk(%v), error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertBulk(%v), error: %v", doList, err)
	}

	return
}

// SelectPrivacy
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type = :key_type
func (dao *UserPrivaciesDAO) SelectPrivacy(ctx context.Context, userId int64, keyType int32) (rValue *dataobject.UserPrivaciesDO, err error) {
	var (
		query string
		do    = &dataobject.UserPrivaciesDO{}
	)
	query = "select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type = ?"

	err = dao.db.QueryRowPartial(ctx, do, query, userId, keyType)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectPrivacy(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectPrivacyList
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type in (:keyList)
func (dao *UserPrivaciesDAO) SelectPrivacyList(ctx context.Context, userId int64, keyList []int32) (rList []dataobject.UserPrivaciesDO, err error) {

	if len(keyList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserPrivaciesDO
	)
	query = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type in (%s)", sqlx.InInt32List(keyList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPrivacyListWithCB
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type in (:keyList)
func (dao *UserPrivaciesDAO) SelectPrivacyListWithCB(ctx context.Context, userId int64, keyList []int32, cb func(sz, i int, v *dataobject.UserPrivaciesDO)) (rList []dataobject.UserPrivaciesDO, err error) {

	if len(keyList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserPrivaciesDO
	)
	query = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type in (%s)", sqlx.InInt32List(keyList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectUsersPrivacyList
// select id, user_id, key_type, rules from user_privacies where user_id in (:idList) and key_type in (:keyList)
func (dao *UserPrivaciesDAO) SelectUsersPrivacyList(ctx context.Context, idList []int32, keyList []int32) (rList []dataobject.UserPrivaciesDO, err error) {
	if len(idList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	if len(keyList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserPrivaciesDO
	)
	query = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id in (%s) and key_type in (%s)", sqlx.InInt32List(idList), sqlx.InInt32List(keyList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersPrivacyList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUsersPrivacyListWithCB
// select id, user_id, key_type, rules from user_privacies where user_id in (:idList) and key_type in (:keyList)
func (dao *UserPrivaciesDAO) SelectUsersPrivacyListWithCB(ctx context.Context, idList []int32, keyList []int32, cb func(sz, i int, v *dataobject.UserPrivaciesDO)) (rList []dataobject.UserPrivaciesDO, err error) {
	if len(idList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	if len(keyList) == 0 {
		rList = []dataobject.UserPrivaciesDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserPrivaciesDO
	)
	query = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id in (%s) and key_type in (%s)", sqlx.InInt32List(idList), sqlx.InInt32List(keyList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersPrivacyList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectPrivacyAll
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id
func (dao *UserPrivaciesDAO) SelectPrivacyAll(ctx context.Context, userId int64) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query  string
		values []dataobject.UserPrivaciesDO
	)
	query = "select id, user_id, key_type, rules from user_privacies where user_id = ?"

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyAll(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPrivacyAllWithCB
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id
func (dao *UserPrivaciesDAO) SelectPrivacyAllWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *dataobject.UserPrivaciesDO)) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query  string
		values []dataobject.UserPrivaciesDO
	)
	query = "select id, user_id, key_type, rules from user_privacies where user_id = ?"

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyAll(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}
