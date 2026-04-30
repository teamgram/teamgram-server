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

type bizUserPeerBlocksModel interface {
	InsertOrUpdate(ctx context.Context, data *UserPeerBlocks) (lastInsertId, rowsAffected int64, err error)
	SelectList(ctx context.Context, userId int64, limit int32) ([]UserPeerBlocks, error)
	SelectListWithCB(ctx context.Context, userId int64, limit int32, cb func(sz, i int, v *UserPeerBlocks)) ([]UserPeerBlocks, error)
	SelectListByIdList(ctx context.Context, userId int64, idList []int64) ([]int64, error)
	SelectListByIdListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v int64)) ([]int64, error)
	Select(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserPeerBlocks, error)
	Delete(ctx context.Context, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type UserPeerBlocksTxModel interface {
	InsertOrUpdate(data *UserPeerBlocks) (lastInsertId, rowsAffected int64, err error)
	SelectList(userId int64, limit int32) ([]UserPeerBlocks, error)
	SelectListByIdList(userId int64, idList []int64) ([]int64, error)
	Select(userId int64, peerType int32, peerId int64) (*UserPeerBlocks, error)
	Delete(userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type defaultUserPeerBlocksTxModel struct {
	tx *sqlx.Tx
}

func NewUserPeerBlocksTxModel(tx *sqlx.Tx) UserPeerBlocksTxModel {
	return &defaultUserPeerBlocksTxModel{tx: tx}
}

// InsertOrUpdate
// insert into user_peer_blocks(user_id, peer_type, peer_id, `date`) values (:user_id, :peer_type, :peer_id, :date) on duplicate key update `date` = values(`date`), deleted = 0
func (m *defaultUserPeerBlocksModel) InsertOrUpdate(ctx context.Context, data *UserPeerBlocks) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_blocks(user_id, peer_type, peer_id, `date`) values (:user_id, :peer_type, :peer_id, :date) on duplicate key update `date` = values(`date`), deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_peer_blocks.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_peer_blocks.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_peer_blocks.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into user_peer_blocks(user_id, peer_type, peer_id, `date`) values (:user_id, :peer_type, :peer_id, :date) on duplicate key update `date` = values(`date`), deleted = 0
func (m *defaultUserPeerBlocksTxModel) InsertOrUpdate(data *UserPeerBlocks) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_blocks(user_id, peer_type, peer_id, `date`) values (:user_id, :peer_type, :peer_id, :date) on duplicate key update `date` = values(`date`), deleted = 0"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_peer_blocks.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_peer_blocks.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_peer_blocks.InsertOrUpdate rows affected: %w", err)
	}

	return
}

// SelectList
// select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = :user_id and deleted = 0 order by id asc limit :limit
func (m *defaultUserPeerBlocksModel) SelectList(ctx context.Context, userId int64, limit int32) (rList []UserPeerBlocks, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and deleted = 0 order by id asc limit ?"
		values []UserPeerBlocks
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPeerBlocks{}
			err = nil
			return
		}
		err = fmt.Errorf("user_peer_blocks.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectList
// select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = :user_id and deleted = 0 order by id asc limit :limit
func (m *defaultUserPeerBlocksTxModel) SelectList(userId int64, limit int32) (rList []UserPeerBlocks, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and deleted = 0 order by id asc limit ?"
		values []UserPeerBlocks
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPeerBlocks{}
			err = nil
			return
		}
		err = fmt.Errorf("user_peer_blocks.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = :user_id and deleted = 0 order by id asc limit :limit
func (m *defaultUserPeerBlocksModel) SelectListWithCB(ctx context.Context, userId int64, limit int32, cb func(sz, i int, v *UserPeerBlocks)) (rList []UserPeerBlocks, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and deleted = 0 order by id asc limit ?"
		values []UserPeerBlocks
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPeerBlocks{}
			err = nil
			return
		}
		err = fmt.Errorf("user_peer_blocks.SelectListWithCB: %w", err)
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

// SelectListByIdList
// select peer_id from user_peer_blocks where user_id = :user_id and peer_type = 2 and peer_id in (:idList) and deleted = 0
func (m *defaultUserPeerBlocksModel) SelectListByIdList(ctx context.Context, userId int64, idList []int64) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select peer_id from user_peer_blocks where user_id = ? and peer_type = 2 and peer_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_peer_blocks.SelectListByIdList: %w", err)
	}

	return
}

// SelectListByIdList
// select peer_id from user_peer_blocks where user_id = :user_id and peer_type = 2 and peer_id in (:idList) and deleted = 0
func (m *defaultUserPeerBlocksTxModel) SelectListByIdList(userId int64, idList []int64) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select peer_id from user_peer_blocks where user_id = ? and peer_type = 2 and peer_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.tx.QueryRowsPartial(&rList, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_peer_blocks.SelectListByIdList: %w", err)
	}

	return
}

// SelectListByIdListWithCB
// select peer_id from user_peer_blocks where user_id = :user_id and peer_type = 2 and peer_id in (:idList) and deleted = 0
func (m *defaultUserPeerBlocksModel) SelectListByIdListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select peer_id from user_peer_blocks where user_id = ? and peer_type = 2 and peer_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_peer_blocks.SelectListByIdListWithCB: %w", err)
		return
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// Select
// select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
func (m *defaultUserPeerBlocksModel) Select(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *UserPeerBlocks, err error) {

	var (
		query = "select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		do    = &UserPeerBlocks{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_peer_blocks",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_peer_blocks.Select: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// Select
// select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
func (m *defaultUserPeerBlocksTxModel) Select(userId int64, peerType int32, peerId int64) (rValue *UserPeerBlocks, err error) {
	var (
		query = "select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		do    = &UserPeerBlocks{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_peer_blocks",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_peer_blocks.Select: %w", err)
		return
	}
	rValue = do

	return
}

// Delete
// update user_peer_blocks set deleted = 1, `date` = 0 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserPeerBlocksModel) Delete(ctx context.Context, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_peer_blocks set deleted = 1, `date` = 0 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("user_peer_blocks.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_peer_blocks.Delete rows affected: %w", err)
		return
	}

	return
}

// Delete
// update user_peer_blocks set deleted = 1, `date` = 0 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserPeerBlocksTxModel) Delete(userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_blocks set deleted = 1, `date` = 0 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("user_peer_blocks.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_peer_blocks.Delete rows affected: %w", err)
		return
	}

	return
}
