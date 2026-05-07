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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizUserPrivaciesModel interface {
	InsertOrUpdate(ctx context.Context, data *UserPrivacies) (lastInsertId, rowsAffected int64, err error)
	InsertBulk(ctx context.Context, doList []*UserPrivacies) (lastInsertId, rowsAffected int64, err error)
	SelectPrivacy(ctx context.Context, userId int64, keyType int32) (*UserPrivacies, error)
	SelectPrivacyList(ctx context.Context, userId int64, keyList []int32) ([]UserPrivacies, error)
	SelectPrivacyListWithCB(ctx context.Context, userId int64, keyList []int32, cb func(sz, i int, v *UserPrivacies)) ([]UserPrivacies, error)
	SelectUsersPrivacyList(ctx context.Context, idList []int64, keyList []int32) ([]UserPrivacies, error)
	SelectUsersPrivacyListWithCB(ctx context.Context, idList []int64, keyList []int32, cb func(sz, i int, v *UserPrivacies)) ([]UserPrivacies, error)
	SelectPrivacyAll(ctx context.Context, userId int64) ([]UserPrivacies, error)
	SelectPrivacyAllWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *UserPrivacies)) ([]UserPrivacies, error)
}

type UserPrivaciesTxModel interface {
	InsertOrUpdate(data *UserPrivacies) (lastInsertId, rowsAffected int64, err error)
	InsertBulk(doList []*UserPrivacies) (lastInsertId, rowsAffected int64, err error)
	SelectPrivacy(userId int64, keyType int32) (*UserPrivacies, error)
	SelectPrivacyList(userId int64, keyList []int32) ([]UserPrivacies, error)
	SelectUsersPrivacyList(idList []int64, keyList []int32) ([]UserPrivacies, error)
	SelectPrivacyAll(userId int64) ([]UserPrivacies, error)
}

type defaultUserPrivaciesTxModel struct {
	tx *sqlx.Tx
}

func NewUserPrivaciesTxModel(tx *sqlx.Tx) UserPrivaciesTxModel {
	return &defaultUserPrivaciesTxModel{tx: tx}
}

// InsertOrUpdate
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)
func (m *defaultUserPrivaciesModel) InsertOrUpdate(ctx context.Context, data *UserPrivacies) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)
func (m *defaultUserPrivaciesTxModel) InsertOrUpdate(data *UserPrivacies) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertOrUpdate rows affected: %w", err)
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
		err = fmt.Errorf("user_privacies.InsertBulk named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertBulk last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertBulk rows affected: %w", err)
	}

	return
}

// InsertBulk
// insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)
func (m *defaultUserPrivaciesTxModel) InsertBulk(doList []*UserPrivacies) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = m.tx.NamedExec(query, doList)
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertBulk named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertBulk last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_privacies.InsertBulk rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_privacies",
				Key:      fmt.Sprintf("user_id=%v,key_type=%v", userId, keyType),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_privacies.SelectPrivacy: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectPrivacy
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type = :key_type
func (m *defaultUserPrivaciesTxModel) SelectPrivacy(userId int64, keyType int32) (rValue *UserPrivacies, err error) {
	var (
		query = "select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type = ?"
		do    = &UserPrivacies{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, keyType)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_privacies",
				Key:      fmt.Sprintf("user_id=%v,key_type=%v", userId, keyType),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_privacies.SelectPrivacy: %w", err)
		return
	}
	rValue = do

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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectPrivacyList: %w", err)
		return
	}

	rList = values

	return
}

// SelectPrivacyList
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id and key_type in (:keyList)
func (m *defaultUserPrivaciesTxModel) SelectPrivacyList(userId int64, keyList []int32) (rList []UserPrivacies, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type in (%s)", sqlx.InInt32List(keyList))
		values []UserPrivacies
	)
	if len(keyList) == 0 {
		rList = []UserPrivacies{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectPrivacyList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectPrivacyListWithCB: %w", err)
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
func (m *defaultUserPrivaciesModel) SelectUsersPrivacyList(ctx context.Context, idList []int64, keyList []int32) (rList []UserPrivacies, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id in (%s) and key_type in (%s)", sqlx.InInt64List(idList), sqlx.InInt32List(keyList))
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectUsersPrivacyList: %w", err)
		return
	}

	rList = values

	return
}

// SelectUsersPrivacyList
// select id, user_id, key_type, rules from user_privacies where user_id in (:idList) and key_type in (:keyList)
func (m *defaultUserPrivaciesTxModel) SelectUsersPrivacyList(idList []int64, keyList []int32) (rList []UserPrivacies, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id in (%s) and key_type in (%s)", sqlx.InInt64List(idList), sqlx.InInt32List(keyList))
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

	err = m.tx.QueryRowsPartial(&values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectUsersPrivacyList: %w", err)
		return
	}

	rList = values

	return
}

// SelectUsersPrivacyListWithCB
// select id, user_id, key_type, rules from user_privacies where user_id in (:idList) and key_type in (:keyList)
func (m *defaultUserPrivaciesModel) SelectUsersPrivacyListWithCB(ctx context.Context, idList []int64, keyList []int32, cb func(sz, i int, v *UserPrivacies)) (rList []UserPrivacies, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, key_type, rules from user_privacies where user_id in (%s) and key_type in (%s)", sqlx.InInt64List(idList), sqlx.InInt32List(keyList))
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectUsersPrivacyListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectPrivacyAll: %w", err)
		return
	}

	rList = values

	return
}

// SelectPrivacyAll
// select id, user_id, key_type, rules from user_privacies where user_id = :user_id
func (m *defaultUserPrivaciesTxModel) SelectPrivacyAll(userId int64) (rList []UserPrivacies, err error) {
	var (
		query  = "select id, user_id, key_type, rules from user_privacies where user_id = ?"
		values []UserPrivacies
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectPrivacyAll: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPrivacies{}
			err = nil
			return
		}
		err = fmt.Errorf("user_privacies.SelectPrivacyAllWithCB: %w", err)
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
