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

type bizHashTagsModel interface {
	InsertOrUpdate(ctx context.Context, data *HashTags) (lastInsertId, rowsAffected int64, err error)
}

type HashTagsTxModel interface {
	InsertOrUpdate(data *HashTags) (lastInsertId, rowsAffected int64, err error)
}

type defaultHashTagsTxModel struct {
	tx *sqlx.Tx
}

func NewHashTagsTxModel(tx *sqlx.Tx) HashTagsTxModel {
	return &defaultHashTagsTxModel{tx: tx}
}

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

// InsertOrUpdate
// insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0
func (m *defaultHashTagsTxModel) InsertOrUpdate(data *HashTags) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into hash_tags(user_id, peer_type, peer_id, hash_tag, hash_tag_message_id) values (:user_id, :peer_type, :peer_id, :hash_tag, :hash_tag_message_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
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
