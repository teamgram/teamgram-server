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
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type HashTagsDAO struct {
	db *sqlx.DB
}

func NewHashTagsDAO(db *sqlx.DB) *HashTagsDAO {
	return &HashTagsDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0
func (dao *HashTagsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.HashTagsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0"
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
// insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0
func (dao *HashTagsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.HashTagsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0"
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

// SelectPeerHashTagList
// select hash_tag_message_id from hash_tags where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and hash_tag = :hash_tag and deleted = 0
func (dao *HashTagsDAO) SelectPeerHashTagList(ctx context.Context, userId int64, peerType int32, peerId int64, hashTag string) (rList []int32, err error) {
	var query = "select hash_tag_message_id from hash_tags where user_id = ? and peer_type = ? and peer_id = ? and hash_tag = ? and deleted = 0"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId, peerType, peerId, hashTag)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPeerHashTagList(_), error: %v", err)
	}

	return
}

// SelectPeerHashTagListWithCB
// select hash_tag_message_id from hash_tags where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and hash_tag = :hash_tag and deleted = 0
func (dao *HashTagsDAO) SelectPeerHashTagListWithCB(ctx context.Context, userId int64, peerType int32, peerId int64, hashTag string, cb func(sz, i int, v int32)) (rList []int32, err error) {
	var query = "select hash_tag_message_id from hash_tags where user_id = ? and peer_type = ? and peer_id = ? and hash_tag = ? and deleted = 0"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId, peerType, peerId, hashTag)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPeerHashTagList(_), error: %v", err)
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
func (dao *HashTagsDAO) DeleteHashTagMessageId(ctx context.Context, userId int64, hashTagMessageId int32) (rowsAffected int64, err error) {
	var (
		query   = "update hash_tags set deleted = 1 where user_id = ? and hash_tag_message_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, hashTagMessageId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteHashTagMessageId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteHashTagMessageId(_), error: %v", err)
	}

	return
}

// DeleteHashTagMessageIdTx
// update hash_tags set deleted = 1 where user_id = :user_id and hash_tag_message_id = :hash_tag_message_id
func (dao *HashTagsDAO) DeleteHashTagMessageIdTx(tx *sqlx.Tx, userId int64, hashTagMessageId int32) (rowsAffected int64, err error) {
	var (
		query   = "update hash_tags set deleted = 1 where user_id = ? and hash_tag_message_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, hashTagMessageId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteHashTagMessageId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteHashTagMessageId(_), error: %v", err)
	}

	return
}
