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

type bizPredefinedUsersModel interface {
	Insert(ctx context.Context, data *PredefinedUsers) (lastInsertId, rowsAffected int64, err error)
	SelectByPhone(ctx context.Context, phone string) (*PredefinedUsers, error)
	SelectPredefinedUsersAll(ctx context.Context) ([]PredefinedUsers, error)
	SelectPredefinedUsersAllWithCB(ctx context.Context, cb func(sz, i int, v *PredefinedUsers)) ([]PredefinedUsers, error)
	Delete(ctx context.Context, phone string) (rowsAffected int64, err error)
	Update(ctx context.Context, cMap map[string]interface{}, phone string) (rowsAffected int64, err error)
}

type PredefinedUsersTxModel interface {
	Insert(data *PredefinedUsers) (lastInsertId, rowsAffected int64, err error)
	SelectByPhone(phone string) (*PredefinedUsers, error)
	SelectPredefinedUsersAll() ([]PredefinedUsers, error)
	Delete(phone string) (rowsAffected int64, err error)
	Update(cMap map[string]interface{}, phone string) (rowsAffected int64, err error)
}

type defaultPredefinedUsersTxModel struct {
	tx *sqlx.Tx
}

func NewPredefinedUsersTxModel(tx *sqlx.Tx) PredefinedUsersTxModel {
	return &defaultPredefinedUsersTxModel{tx: tx}
}

// Insert
// insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)
func (m *defaultPredefinedUsersModel) Insert(ctx context.Context, data *PredefinedUsers) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("predefined_users.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("predefined_users.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("predefined_users.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)
func (m *defaultPredefinedUsersTxModel) Insert(data *PredefinedUsers) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("predefined_users.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("predefined_users.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("predefined_users.Insert rows affected: %w", err)
	}

	return
}

// SelectByPhone
// select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where phone = :phone and deleted = 0 limit 1
func (m *defaultPredefinedUsersModel) SelectByPhone(ctx context.Context, phone string) (rValue *PredefinedUsers, err error) {

	var (
		query = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where phone = ? and deleted = 0 limit 1"
		do    = &PredefinedUsers{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, phone)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "predefined_users",
				Key:      fmt.Sprintf("phone=%v", phone),
				Cause:    err,
			}
		}
		err = fmt.Errorf("predefined_users.SelectByPhone: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByPhone
// select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where phone = :phone and deleted = 0 limit 1
func (m *defaultPredefinedUsersTxModel) SelectByPhone(phone string) (rValue *PredefinedUsers, err error) {
	var (
		query = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where phone = ? and deleted = 0 limit 1"
		do    = &PredefinedUsers{}
	)
	err = m.tx.QueryRowPartial(do, query, phone)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "predefined_users",
				Key:      fmt.Sprintf("phone=%v", phone),
				Cause:    err,
			}
		}
		err = fmt.Errorf("predefined_users.SelectByPhone: %w", err)
		return
	}
	rValue = do

	return
}

// SelectPredefinedUsersAll
// select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc
func (m *defaultPredefinedUsersModel) SelectPredefinedUsersAll(ctx context.Context) (rList []PredefinedUsers, err error) {
	var (
		query  = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc"
		values []PredefinedUsers
	)
	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PredefinedUsers{}
			err = nil
			return
		}
		err = fmt.Errorf("predefined_users.SelectPredefinedUsersAll: %w", err)
		return
	}

	rList = values

	return
}

// SelectPredefinedUsersAll
// select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc
func (m *defaultPredefinedUsersTxModel) SelectPredefinedUsersAll() (rList []PredefinedUsers, err error) {
	var (
		query  = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc"
		values []PredefinedUsers
	)
	err = m.tx.QueryRowsPartial(&values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PredefinedUsers{}
			err = nil
			return
		}
		err = fmt.Errorf("predefined_users.SelectPredefinedUsersAll: %w", err)
		return
	}

	rList = values

	return
}

// SelectPredefinedUsersAllWithCB
// select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc
func (m *defaultPredefinedUsersModel) SelectPredefinedUsersAllWithCB(ctx context.Context, cb func(sz, i int, v *PredefinedUsers)) (rList []PredefinedUsers, err error) {
	var (
		query  = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc"
		values []PredefinedUsers
	)
	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PredefinedUsers{}
			err = nil
			return
		}
		err = fmt.Errorf("predefined_users.SelectPredefinedUsersAllWithCB: %w", err)
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

// Delete
// update predefined_users set deleted = 0 where phone = :phone
func (m *defaultPredefinedUsersModel) Delete(ctx context.Context, phone string) (rowsAffected int64, err error) {

	var (
		query   = "update predefined_users set deleted = 0 where phone = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, phone)

	if err != nil {
		err = fmt.Errorf("predefined_users.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("predefined_users.Delete rows affected: %w", err)
		return
	}

	return
}

// Delete
// update predefined_users set deleted = 0 where phone = :phone
func (m *defaultPredefinedUsersTxModel) Delete(phone string) (rowsAffected int64, err error) {
	var (
		query   = "update predefined_users set deleted = 0 where phone = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, phone)

	if err != nil {
		err = fmt.Errorf("predefined_users.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("predefined_users.Delete rows affected: %w", err)
		return
	}

	return
}

// Update
// update predefined_users set %s where phone = :phone
func (m *defaultPredefinedUsersModel) Update(ctx context.Context, cMap map[string]interface{}, phone string) (rowsAffected int64, err error) {

	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update predefined_users set %s where phone = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, phone)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		err = fmt.Errorf("predefined_users.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("predefined_users.Update rows affected: %w", err)
		return
	}

	return
}

// Update
// update predefined_users set %s where phone = :phone
func (m *defaultPredefinedUsersTxModel) Update(cMap map[string]interface{}, phone string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update predefined_users set %s where phone = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, phone)

	rResult, err = m.tx.Exec(query, aValues...)

	if err != nil {
		err = fmt.Errorf("predefined_users.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("predefined_users.Update rows affected: %w", err)
		return
	}

	return
}
