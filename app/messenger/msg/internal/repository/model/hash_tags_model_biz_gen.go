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

type (
	bizHashTagsModel interface {
		InsertOrUpdate(ctx context.Context, data *HashTags) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *HashTags) (lastInsertId, rowsAffected int64, err error)

		SelectPeerHashTagList(ctx context.Context, userId int64, peerType int32, peerId int64, hashTag string) ([]int32, error)
		SelectPeerHashTagListTx(tx *sqlx.Tx, userId int64, peerType int32, peerId int64, hashTag string) ([]int32, error)
		SelectPeerHashTagListWithCB(ctx context.Context, userId int64, peerType int32, peerId int64, hashTag string, cb func(sz, i int, v int32)) ([]int32, error)

		DeleteHashTagMessageId(ctx context.Context, userId int64, hashTagMessageId int32) (rowsAffected int64, err error)
		DeleteHashTagMessageIdTx(tx *sqlx.Tx, userId int64, hashTagMessageId int32) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0
func (m *defaultHashTagsModel) InsertOrUpdate(ctx context.Context, data *HashTags) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("hash_tags.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("hash_tags.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("hash_tags.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0
func (m *defaultHashTagsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *HashTags) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("hash_tags.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("hash_tags.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("hash_tags.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}

// SelectPeerHashTagList
// select hash_tag_message_id from hash_tags where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and hash_tag = :hash_tag and deleted = 0
func (m *defaultHashTagsModel) SelectPeerHashTagList(ctx context.Context, userId int64, peerType int32, peerId int64, hashTag string) (rList []int32, err error) {
	var query = "select hash_tag_message_id from hash_tags where user_id = ? and peer_type = ? and peer_id = ? and hash_tag = ? and deleted = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, peerType, peerId, hashTag)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int32{}
			err = nil
			return
		}
		err = fmt.Errorf("hash_tags.SelectPeerHashTagList: %w", err)
	}

	return
}

// SelectPeerHashTagListTx
// select hash_tag_message_id from hash_tags where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and hash_tag = :hash_tag and deleted = 0
func (m *defaultHashTagsModel) SelectPeerHashTagListTx(tx *sqlx.Tx, userId int64, peerType int32, peerId int64, hashTag string) (rList []int32, err error) {
	var query = "select hash_tag_message_id from hash_tags where user_id = ? and peer_type = ? and peer_id = ? and hash_tag = ? and deleted = 0"
	err = tx.QueryRowsPartial(&rList, query, userId, peerType, peerId, hashTag)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int32{}
			err = nil
			return
		}
		err = fmt.Errorf("hash_tags.SelectPeerHashTagListTx: %w", err)
	}

	return
}

// SelectPeerHashTagListWithCB
// select hash_tag_message_id from hash_tags where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and hash_tag = :hash_tag and deleted = 0
func (m *defaultHashTagsModel) SelectPeerHashTagListWithCB(ctx context.Context, userId int64, peerType int32, peerId int64, hashTag string, cb func(sz, i int, v int32)) (rList []int32, err error) {
	var query = "select hash_tag_message_id from hash_tags where user_id = ? and peer_type = ? and peer_id = ? and hash_tag = ? and deleted = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, peerType, peerId, hashTag)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int32{}
			err = nil
			return
		}
		err = fmt.Errorf("hash_tags.SelectPeerHashTagListWithCB: %w", err)
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

// DeleteHashTagMessageId
// update hash_tags set deleted = 1 where user_id = :user_id and hash_tag_message_id = :hash_tag_message_id
func (m *defaultHashTagsModel) DeleteHashTagMessageId(ctx context.Context, userId int64, hashTagMessageId int32) (rowsAffected int64, err error) {

	var (
		query   = "update hash_tags set deleted = 1 where user_id = ? and hash_tag_message_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, hashTagMessageId)

	if err != nil {
		err = fmt.Errorf("hash_tags.DeleteHashTagMessageId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("hash_tags.DeleteHashTagMessageId rows affected: %w", err)
		return
	}

	if rowsAffected == 0 {
		err = &NotFoundError{
			Resource: "hash_tags",
			Key:      fmt.Sprintf("user_id=%v,hash_tag_message_id=%v", userId, hashTagMessageId),
			Cause:    ErrNotFound,
		}
	}

	return
}

// DeleteHashTagMessageIdTx
// update hash_tags set deleted = 1 where user_id = :user_id and hash_tag_message_id = :hash_tag_message_id
func (m *defaultHashTagsModel) DeleteHashTagMessageIdTx(tx *sqlx.Tx, userId int64, hashTagMessageId int32) (rowsAffected int64, err error) {
	var (
		query   = "update hash_tags set deleted = 1 where user_id = ? and hash_tag_message_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, hashTagMessageId)

	if err != nil {
		err = fmt.Errorf("hash_tags.DeleteHashTagMessageIdTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("hash_tags.DeleteHashTagMessageIdTx rows affected: %w", err)
		return
	}

	if rowsAffected == 0 {
		err = &NotFoundError{
			Resource: "hash_tags",
			Key:      fmt.Sprintf("user_id=%v,hash_tag_message_id=%v", userId, hashTagMessageId),
			Cause:    ErrNotFound,
		}
	}

	return
}
