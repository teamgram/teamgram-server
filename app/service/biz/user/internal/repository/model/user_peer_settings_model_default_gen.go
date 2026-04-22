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
	userPeerSettingsFieldNames          = builder.RawFieldNames(&UserPeerSettings{})
	userPeerSettingsRows                = strings.Join(userPeerSettingsFieldNames, ",")
	userPeerSettingsRowsExpectAutoSet   = strings.Join(stringx.Remove(userPeerSettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPeerSettingsRowsWithPlaceHolder = strings.Join(stringx.Remove(userPeerSettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTUserPeerSettingsIdPrefix = "cache:t:user_peer_settings:id:"

	cacheUserPeerSettingsIdPrefix = "cache#UserPeerSettings#id"

	cacheUserPeerSettingsUserIdPeerTypePeerIdPrefix = "cache#UserId#PeerType#PeerId"
)

type (
	userPeerSettingsModel interface {
		Insert2(ctx context.Context, data *UserPeerSettings) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserPeerSettings, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserPeerSettings, error)
		Update2(ctx context.Context, data *UserPeerSettings) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserPeerSettings, error)
	}

	defaultUserPeerSettingsModel struct {
		db *sqlx.DB
	}

	UserPeerSettings struct {
		Id                    int64 `db:"id" json:"id"`
		UserId                int64 `db:"user_id" json:"user_id"`
		PeerType              int32 `db:"peer_type" json:"peer_type"`
		PeerId                int64 `db:"peer_id" json:"peer_id"`
		Hide                  bool  `db:"hide" json:"hide"`
		ReportSpam            bool  `db:"report_spam" json:"report_spam"`
		AddContact            bool  `db:"add_contact" json:"add_contact"`
		BlockContact          bool  `db:"block_contact" json:"block_contact"`
		ShareContact          bool  `db:"share_contact" json:"share_contact"`
		NeedContactsException bool  `db:"need_contacts_exception" json:"need_contacts_exception"`
		ReportGeo             bool  `db:"report_geo" json:"report_geo"`
		Autoarchived          bool  `db:"autoarchived" json:"autoarchived"`
		InviteMembers         bool  `db:"invite_members" json:"invite_members"`
		GeoDistance           int32 `db:"geo_distance" json:"geo_distance"`
	}
)

func newUserPeerSettingsModel(db *sqlx.DB) *defaultUserPeerSettingsModel {
	return &defaultUserPeerSettingsModel{
		db: db,
	}
}

func (m *defaultUserPeerSettingsModel) Insert2(ctx context.Context, data *UserPeerSettings) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_peer_settings` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", userPeerSettingsRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.Hide, data.ReportSpam, data.AddContact, data.BlockContact, data.ShareContact, data.NeedContactsException, data.ReportGeo, data.Autoarchived, data.InviteMembers, data.GeoDistance)
}

func (m *defaultUserPeerSettingsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_peer_settings` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserPeerSettingsModel) FindOne(ctx context.Context, id int64) (*UserPeerSettings, error) {
	query := fmt.Sprintf("select %s from user_peer_settings where id = ? limit 1", userPeerSettingsRows)
	var resp UserPeerSettings
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserPeerSettingsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserPeerSettings, error) {
	if len(id) == 0 {
		return []UserPeerSettings{}, nil
	}

	query := fmt.Sprintf("select %s from user_peer_settings where id in (%s)", userPeerSettingsRows, sqlx.InInt64List(id))

	var resp []UserPeerSettings
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserPeerSettingsModel) Update2(ctx context.Context, data *UserPeerSettings) error {
	query := fmt.Sprintf("update `user_peer_settings` set %s where `id` = ?", userPeerSettingsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.Hide, data.ReportSpam, data.AddContact, data.BlockContact, data.ShareContact, data.NeedContactsException, data.ReportGeo, data.Autoarchived, data.InviteMembers, data.GeoDistance, data.Id)
	return err
}

func (m *defaultUserPeerSettingsModel) FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserPeerSettings, error) {
	query := fmt.Sprintf("select %s from user_peer_settings where user_id = ? AND peer_type = ? AND peer_id = ? limit 1", userPeerSettingsRows)
	var resp UserPeerSettings
	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserPeerSettingsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheUserPeerSettingsIdPrefix, primary)
}

func (m *defaultUserPeerSettingsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from user_peer_settings where id = ? limit 1", userPeerSettingsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
