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
	bizUserPrivaciesModel interface {
		InsertOrUpdate(ctx context.Context, data *UserPrivacies) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UserPrivacies) (lastInsertId, rowsAffected int64, err error)

		InsertBulk(ctx context.Context, doList []*UserPrivacies) (lastInsertId, rowsAffected int64, err error)
		InsertBulkTx(tx *sqlx.Tx, doList []*UserPrivacies) (lastInsertId, rowsAffected int64, err error)

		SelectPrivacy(ctx context.Context, userId int64, keyType int32) (*UserPrivacies, error)

		SelectPrivacyList(ctx context.Context, userId int64, keyList []int32) ([]UserPrivacies, error)
		SelectPrivacyListWithCB(ctx context.Context, userId int64, keyList []int32, cb func(sz, i int, v *UserPrivacies)) ([]UserPrivacies, error)

		SelectUsersPrivacyList(ctx context.Context, idList []int32, keyList []int32) ([]UserPrivacies, error)
		SelectUsersPrivacyListWithCB(ctx context.Context, idList []int32, keyList []int32, cb func(sz, i int, v *UserPrivacies)) ([]UserPrivacies, error)

		SelectPrivacyAll(ctx context.Context, userId int64) ([]UserPrivacies, error)
		SelectPrivacyAllWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *UserPrivacies)) ([]UserPrivacies, error)
	}
)

// InsertOrUpdate
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)
func (m *defaultUserPrivaciesModel) InsertOrUpdate(ctx context.Context, data *UserPrivacies) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"
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
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)
func (m *defaultUserPrivaciesModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UserPrivacies) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"
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

// InsertBulk
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)
func (m *defaultUserPrivaciesModel) InsertBulk(ctx context.Context, doList []*UserPrivacies) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = m.db.NamedExec(ctx, query, doList)
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
func (m *defaultUserPrivaciesModel) InsertBulkTx(tx *sqlx.Tx, doList []*UserPrivacies) (lastInsertId, rowsAffected int64, err error) {
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
func (m *defaultUserPrivaciesModel) SelectPrivacy(ctx context.Context, userId int64, keyType int32) (rValue *UserPrivacies, err error) {
	var (
		query = "select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type = ?"
		do    = &UserPrivacies{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, keyType)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
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
func (m *defaultUserPrivaciesModel) SelectPrivacyList(ctx context.Context, userId int64, keyList []int32) (rList []UserPrivacies, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type in (%s)", sqlx.InInt32List(keyList))
		values []UserPrivacies
	)
	if len(keyList) == 0 {
		rList = []UserPrivacies{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPrivacyListWithCB
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type in (:keyList)
func (m *defaultUserPrivaciesModel) SelectPrivacyListWithCB(ctx context.Context, userId int64, keyList []int32, cb func(sz, i int, v *UserPrivacies)) (rList []UserPrivacies, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type in (%s)", sqlx.InInt32List(keyList))
		values []UserPrivacies
	)
	if len(keyList) == 0 {
		rList = []UserPrivacies{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyList(_), error: %v", err)
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

// SelectUsersPrivacyList
// select id, user_id, key_type, rules from user_privacies where user_id in (:idList) and key_type in (:keyList)
func (m *defaultUserPrivaciesModel) SelectUsersPrivacyList(ctx context.Context, idList []int32, keyList []int32) (rList []UserPrivacies, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id in (%s) and key_type in (%s)", sqlx.InInt32List(idList), sqlx.InInt32List(keyList))
		values []UserPrivacies
	)
	if len(idList) == 0 {
		rList = []UserPrivacies{}
		return
	}
	if len(keyList) == 0 {
		rList = []UserPrivacies{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersPrivacyList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUsersPrivacyListWithCB
// select id, user_id, key_type, rules from user_privacies where user_id in (:idList) and key_type in (:keyList)
func (m *defaultUserPrivaciesModel) SelectUsersPrivacyListWithCB(ctx context.Context, idList []int32, keyList []int32, cb func(sz, i int, v *UserPrivacies)) (rList []UserPrivacies, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id in (%s) and key_type in (%s)", sqlx.InInt32List(idList), sqlx.InInt32List(keyList))
		values []UserPrivacies
	)
	if len(idList) == 0 {
		rList = []UserPrivacies{}
		return
	}
	if len(keyList) == 0 {
		rList = []UserPrivacies{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersPrivacyList(_), error: %v", err)
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

// SelectPrivacyAll
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id
func (m *defaultUserPrivaciesModel) SelectPrivacyAll(ctx context.Context, userId int64) (rList []UserPrivacies, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id = ?"
		values []UserPrivacies
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyAll(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPrivacyAllWithCB
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id
func (m *defaultUserPrivaciesModel) SelectPrivacyAllWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *UserPrivacies)) (rList []UserPrivacies, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id = ?"
		values []UserPrivacies
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPrivacyAll(_), error: %v", err)
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
