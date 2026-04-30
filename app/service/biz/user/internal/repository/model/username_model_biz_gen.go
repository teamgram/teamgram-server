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

type bizUsernameModel interface {
	Insert(ctx context.Context, data *Username) (lastInsertId, rowsAffected int64, err error)
	SelectList(ctx context.Context, nameList []string) ([]Username, error)
	SelectListWithCB(ctx context.Context, nameList []string, cb func(sz, i int, v *Username)) ([]Username, error)
	SelectByUsername(ctx context.Context, username string) (*Username, error)
	Update(ctx context.Context, cMap map[string]interface{}, username string) (rowsAffected int64, err error)
	Delete(ctx context.Context, username string) (rowsAffected int64, err error)
	DeleteByPeer(ctx context.Context, peerType int32, peerId int64) (rowsAffected int64, err error)
	DeleteByChannelId(ctx context.Context, peerId int64) (rowsAffected int64, err error)
	SelectByPeer(ctx context.Context, peerType int32, peerId int64) (*Username, error)
	SelectByUserId(ctx context.Context, peerId int64) (*Username, error)
	SelectListByUserId(ctx context.Context, peerId int64) ([]Username, error)
	SelectListByUserIdWithCB(ctx context.Context, peerId int64, cb func(sz, i int, v *Username)) ([]Username, error)
	SelectByChannelId(ctx context.Context, peerId int64) (*Username, error)
	SelectListByChannelId(ctx context.Context, peerId int64) ([]Username, error)
	SelectListByChannelIdWithCB(ctx context.Context, peerId int64, cb func(sz, i int, v *Username)) ([]Username, error)
	UpdateUsername(ctx context.Context, username string, peerType int32, peerId int64) (rowsAffected int64, err error)
	SearchByQueryNotIdList(ctx context.Context, q2 string, idList []int64, limit int32) ([]Username, error)
	SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, idList []int64, limit int32, cb func(sz, i int, v *Username)) ([]Username, error)
}

type UsernameTxModel interface {
	Insert(data *Username) (lastInsertId, rowsAffected int64, err error)
	SelectList(nameList []string) ([]Username, error)
	SelectByUsername(username string) (*Username, error)
	Update(cMap map[string]interface{}, username string) (rowsAffected int64, err error)
	Delete(username string) (rowsAffected int64, err error)
	DeleteByPeer(peerType int32, peerId int64) (rowsAffected int64, err error)
	DeleteByChannelId(peerId int64) (rowsAffected int64, err error)
	SelectByPeer(peerType int32, peerId int64) (*Username, error)
	SelectByUserId(peerId int64) (*Username, error)
	SelectListByUserId(peerId int64) ([]Username, error)
	SelectByChannelId(peerId int64) (*Username, error)
	SelectListByChannelId(peerId int64) ([]Username, error)
	UpdateUsername(username string, peerType int32, peerId int64) (rowsAffected int64, err error)
	SearchByQueryNotIdList(q2 string, idList []int64, limit int32) ([]Username, error)
}

type defaultUsernameTxModel struct {
	tx *sqlx.Tx
}

func NewUsernameTxModel(tx *sqlx.Tx) UsernameTxModel {
	return &defaultUsernameTxModel{tx: tx}
}

// Insert
// insert into username(username, peer_type, peer_id, editable, active, order2, deleted) values (:username, :peer_type, :peer_id, :editable, :active, :order2, 0)
func (m *defaultUsernameModel) Insert(ctx context.Context, data *Username) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into username(username, peer_type, peer_id, editable, active, order2, deleted) values (:username, :peer_type, :peer_id, :editable, :active, :order2, 0)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("username.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("username.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into username(username, peer_type, peer_id, editable, active, order2, deleted) values (:username, :peer_type, :peer_id, :editable, :active, :order2, 0)
func (m *defaultUsernameTxModel) Insert(data *Username) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into username(username, peer_type, peer_id, editable, active, order2, deleted) values (:username, :peer_type, :peer_id, :editable, :active, :order2, 0)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("username.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("username.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.Insert rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectList
// select username, peer_type, peer_id, editable, active, order2 from username where username in (:nameList) and editable = 1
func (m *defaultUsernameTxModel) SelectList(nameList []string) (rList []Username, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id, editable, active, order2 from username where username in (%s) and editable = 1", sqlx.InStringList(nameList))
		values []Username
	)
	if len(nameList) == 0 {
		rList = []Username{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("username=%v", username),
				Cause:    err,
			}
		}
		err = fmt.Errorf("username.SelectByUsername: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUsername
// select username, peer_type, peer_id, editable, active, order2, deleted from username where username = :username
func (m *defaultUsernameTxModel) SelectByUsername(username string) (rValue *Username, err error) {
	var (
		query = "select username, peer_type, peer_id, editable, active, order2, deleted from username where username = ?"
		do    = &Username{}
	)
	err = m.tx.QueryRowPartial(do, query, username)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("username=%v", username),
				Cause:    err,
			}
		}
		err = fmt.Errorf("username.SelectByUsername: %w", err)
		return
	}
	rValue = do

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
		err = fmt.Errorf("username.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.Update rows affected: %w", err)
		return
	}

	return
}

// Update
// update username set %s where username = :username
func (m *defaultUsernameTxModel) Update(cMap map[string]interface{}, username string) (rowsAffected int64, err error) {
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

	rResult, err = m.tx.Exec(query, aValues...)

	if err != nil {
		err = fmt.Errorf("username.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.Update rows affected: %w", err)
		return
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
		err = fmt.Errorf("username.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.Delete rows affected: %w", err)
		return
	}

	return
}

// Delete
// delete from username where username = :username
func (m *defaultUsernameTxModel) Delete(username string) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where username = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, username)

	if err != nil {
		err = fmt.Errorf("username.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.Delete rows affected: %w", err)
		return
	}

	return
}

// DeleteByPeer
// delete from username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameModel) DeleteByPeer(ctx context.Context, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "delete from username where peer_type = ? and peer_id = ? and editable = 1"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("username.DeleteByPeer exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.DeleteByPeer rows affected: %w", err)
		return
	}

	return
}

// DeleteByPeer
// delete from username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameTxModel) DeleteByPeer(peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = ? and peer_id = ? and editable = 1"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("username.DeleteByPeer exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.DeleteByPeer rows affected: %w", err)
		return
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
		err = fmt.Errorf("username.DeleteByChannelId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.DeleteByChannelId rows affected: %w", err)
		return
	}

	return
}

// DeleteByChannelId
// delete from username where peer_type = 2 and peer_id = :peer_id and editable = 0
func (m *defaultUsernameTxModel) DeleteByChannelId(peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from username where peer_type = 2 and peer_id = ? and editable = 0"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, peerId)

	if err != nil {
		err = fmt.Errorf("username.DeleteByChannelId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.DeleteByChannelId rows affected: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("peer_type=%v,peer_id=%v", peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("username.SelectByPeer: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByPeer
// select username, peer_type, peer_id, editable, active, order2 from username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameTxModel) SelectByPeer(peerType int32, peerId int64) (rValue *Username, err error) {
	var (
		query = "select username, peer_type, peer_id, editable, active, order2 from username where peer_type = ? and peer_id = ? and editable = 1"
		do    = &Username{}
	)
	err = m.tx.QueryRowPartial(do, query, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("peer_type=%v,peer_id=%v", peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("username.SelectByPeer: %w", err)
		return
	}
	rValue = do

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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("peer_id=%v", peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("username.SelectByUserId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserId
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = :peer_id and editable = 1
func (m *defaultUsernameTxModel) SelectByUserId(peerId int64) (rValue *Username, err error) {
	var (
		query = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = ? and editable = 1"
		do    = &Username{}
	)
	err = m.tx.QueryRowPartial(do, query, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("peer_id=%v", peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("username.SelectByUserId: %w", err)
		return
	}
	rValue = do

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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectListByUserId: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByUserId
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = :peer_id
func (m *defaultUsernameTxModel) SelectListByUserId(peerId int64) (rList []Username, err error) {
	var (
		query  = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id = ?"
		values []Username
	)
	err = m.tx.QueryRowsPartial(&values, query, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectListByUserId: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectListByUserIdWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("peer_id=%v", peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("username.SelectByChannelId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByChannelId
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = :peer_id and editable = 1
func (m *defaultUsernameTxModel) SelectByChannelId(peerId int64) (rValue *Username, err error) {
	var (
		query = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = ? and editable = 1"
		do    = &Username{}
	)
	err = m.tx.QueryRowPartial(do, query, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("peer_id=%v", peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("username.SelectByChannelId: %w", err)
		return
	}
	rValue = do

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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectListByChannelId: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByChannelId
// select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = :peer_id
func (m *defaultUsernameTxModel) SelectListByChannelId(peerId int64) (rList []Username, err error) {
	var (
		query  = "select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 4 and peer_id = ?"
		values []Username
	)
	err = m.tx.QueryRowsPartial(&values, query, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectListByChannelId: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SelectListByChannelIdWithCB: %w", err)
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
		err = fmt.Errorf("username.UpdateUsername exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.UpdateUsername rows affected: %w", err)
		return
	}

	return
}

// UpdateUsername
// update username set username = :username where peer_type = :peer_type and peer_id = :peer_id and editable = 1
func (m *defaultUsernameTxModel) UpdateUsername(username string, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ? and editable = 1"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, username, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("username.UpdateUsername exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("username.UpdateUsername rows affected: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SearchByQueryNotIdList: %w", err)
		return
	}

	rList = values

	return
}

// SearchByQueryNotIdList
// select username, peer_type, peer_id from username where username like :q2 and peer_id not in (:id_list) limit :limit
func (m *defaultUsernameTxModel) SearchByQueryNotIdList(q2 string, idList []int64, limit int32) (rList []Username, err error) {
	var (
		query  = fmt.Sprintf("select username, peer_type, peer_id from username where username like ? and peer_id not in (%s) limit ?", sqlx.InInt64List(idList))
		values []Username
	)
	if len(idList) == 0 {
		rList = []Username{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, q2, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SearchByQueryNotIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Username{}
			err = nil
			return
		}
		err = fmt.Errorf("username.SearchByQueryNotIdListWithCB: %w", err)
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
