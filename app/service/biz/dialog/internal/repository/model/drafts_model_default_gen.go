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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	draftsFieldNames          = builder.RawFieldNames(&Drafts{})
	draftsRows                = strings.Join(draftsFieldNames, ",")
	draftsRowsExpectAutoSet   = strings.Join(stringx.Remove(draftsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	draftsRowsWithPlaceHolder = strings.Join(stringx.Remove(draftsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTDraftsIdPrefix = "cache:t:drafts:id:"

	cacheDraftsIdPrefix = "cache#Drafts#id"

	cacheDraftsUserIdPeerDialogIdPrefix = "cache#UserId#PeerDialogId"
)

type (
	draftsModel interface {
		Insert2(ctx context.Context, data *Drafts) (sql.Result, error)
		FindOne(ctx context.Context, id int32) (*Drafts, error)
		FindListByIdList(ctx context.Context, id ...int32) ([]Drafts, error)
		Update2(ctx context.Context, data *Drafts) error
		Delete2(ctx context.Context, id int32) error

		FindOneByUserIdPeerDialogId(ctx context.Context, userId int32, peerDialogId int64) (*Drafts, error)
	}

	defaultDraftsModel struct {
		db *sqlx.DB
	}

	Drafts struct {
		Id               int32  `db:"id" json:"id"`
		UserId           int32  `db:"user_id" json:"user_id"`
		PeerDialogId     int64  `db:"peer_dialog_id" json:"peer_dialog_id"`
		DraftType        int32  `db:"draft_type" json:"draft_type"`
		DraftMessageData string `db:"draft_message_data" json:"draft_message_data"`
		Date2            int64  `db:"date2" json:"date2"`
	}
)

func newDraftsModel(db *sqlx.DB) *defaultDraftsModel {
	return &defaultDraftsModel{
		db: db,
	}
}

func (m *defaultDraftsModel) Insert2(ctx context.Context, data *Drafts) (sql.Result, error) {
	query := fmt.Sprintf("insert into `drafts` (%s) values (?, ?, ?, ?, ?)", draftsRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.PeerDialogId, data.DraftType, data.DraftMessageData, data.Date2)
}

func (m *defaultDraftsModel) Delete2(ctx context.Context, id int32) error {
	query := "delete from `drafts` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultDraftsModel) FindOne(ctx context.Context, id int32) (*Drafts, error) {
	query := fmt.Sprintf("select %s from drafts where id = ? limit 1", draftsRows)
	var resp Drafts
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultDraftsModel) FindListByIdList(ctx context.Context, id ...int32) ([]Drafts, error) {
	if len(id) == 0 {
		return []Drafts{}, nil
	}
	query := fmt.Sprintf("select %s from drafts where id in (%s)", draftsRows, sqlx.InInt32List(id))

	var resp []Drafts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultDraftsModel) Update2(ctx context.Context, data *Drafts) error {
	query := fmt.Sprintf("update `drafts` set %s where `id` = ?", draftsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerDialogId, data.DraftType, data.DraftMessageData, data.Date2, data.Id)
	return err
}

func (m *defaultDraftsModel) FindOneByUserIdPeerDialogId(ctx context.Context, userId int32, peerDialogId int64) (*Drafts, error) {
	query := fmt.Sprintf("select %s from drafts where user_id = ? AND peer_dialog_id = ? limit 1", draftsRows)
	var resp Drafts
	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerDialogId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultDraftsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheDraftsIdPrefix, primary)
}

func (m *defaultDraftsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from drafts where id = ? limit 1", draftsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
