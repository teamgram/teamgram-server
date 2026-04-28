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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	usersModel interface {
		Insert2(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Users, error)
		Update2(ctx context.Context, data *Users) error
		Delete2(ctx context.Context, id int64) error

		FindOneByPhone(ctx context.Context, phone string) (*Users, error)
		FindListByPhoneList(ctx context.Context, phone ...string) ([]Users, error)
	}

	defaultUsersModel struct {
		db *sqlx.DB
	}

	Users struct {
		Id                               int64  `db:"id" json:"id"`
		UserType                         int32  `db:"user_type" json:"user_type"`
		AccessHash                       int64  `db:"access_hash" json:"access_hash"`
		SecretKeyId                      int64  `db:"secret_key_id" json:"secret_key_id"`
		FirstName                        string `db:"first_name" json:"first_name"`
		LastName                         string `db:"last_name" json:"last_name"`
		Username                         string `db:"username" json:"username"`
		Phone                            string `db:"phone" json:"phone"`
		CountryCode                      string `db:"country_code" json:"country_code"`
		Verified                         bool   `db:"verified" json:"verified"`
		Support                          bool   `db:"support" json:"support"`
		Scam                             bool   `db:"scam" json:"scam"`
		Fake                             bool   `db:"fake" json:"fake"`
		Premium                          bool   `db:"premium" json:"premium"`
		PremiumExpireDate                int64  `db:"premium_expire_date" json:"premium_expire_date"`
		About                            string `db:"about" json:"about"`
		State                            int32  `db:"state" json:"state"`
		IsBot                            bool   `db:"is_bot" json:"is_bot"`
		AccountDaysTtl                   int32  `db:"account_days_ttl" json:"account_days_ttl"`
		PhotoId                          int64  `db:"photo_id" json:"photo_id"`
		Restricted                       bool   `db:"restricted" json:"restricted"`
		RestrictionReason                string `db:"restriction_reason" json:"restriction_reason"`
		ArchiveAndMuteNewNoncontactPeers bool   `db:"archive_and_mute_new_noncontact_peers" json:"archive_and_mute_new_noncontact_peers"`
		EmojiStatusDocumentId            int64  `db:"emoji_status_document_id" json:"emoji_status_document_id"`
		EmojiStatusUntil                 int32  `db:"emoji_status_until" json:"emoji_status_until"`
		StoriesMaxId                     int32  `db:"stories_max_id" json:"stories_max_id"`
		Color                            int32  `db:"color" json:"color"`
		ColorBackgroundEmojiId           int64  `db:"color_background_emoji_id" json:"color_background_emoji_id"`
		ProfileColor                     int32  `db:"profile_color" json:"profile_color"`
		ProfileColorBackgroundEmojiId    int64  `db:"profile_color_background_emoji_id" json:"profile_color_background_emoji_id"`
		Birthday                         string `db:"birthday" json:"birthday"`
		PersonalChannelId                int64  `db:"personal_channel_id" json:"personal_channel_id"`
		AuthorizationTtlDays             int32  `db:"authorization_ttl_days" json:"authorization_ttl_days"`
		SavedMusicId                     int64  `db:"saved_music_id" json:"saved_music_id"`
		MainTab                          int32  `db:"main_tab" json:"main_tab"`
		Deleted                          bool   `db:"deleted" json:"deleted"`
		DeleteReason                     string `db:"delete_reason" json:"delete_reason"`
	}
)

func newUsersModel(db *sqlx.DB) *defaultUsersModel {
	return &defaultUsersModel{
		db: db,
	}
}

func (m *defaultUsersModel) Insert2(ctx context.Context, data *Users) (sql.Result, error) {
	tableName := "users"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, usersRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserType, data.AccessHash, data.SecretKeyId, data.FirstName, data.LastName, data.Username, data.Phone, data.CountryCode, data.Verified, data.Support, data.Scam, data.Fake, data.Premium, data.PremiumExpireDate, data.About, data.State, data.IsBot, data.AccountDaysTtl, data.PhotoId, data.Restricted, data.RestrictionReason, data.ArchiveAndMuteNewNoncontactPeers, data.EmojiStatusDocumentId, data.EmojiStatusUntil, data.StoriesMaxId, data.Color, data.ColorBackgroundEmojiId, data.ProfileColor, data.ProfileColorBackgroundEmojiId, data.Birthday, data.PersonalChannelId, data.AuthorizationTtlDays, data.SavedMusicId, data.MainTab, data.Deleted, data.DeleteReason)
	if err != nil {
		return nil, fmt.Errorf("users.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUsersModel) Delete2(ctx context.Context, id int64) error {
	tableName := "users"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("users.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	tableName := "users"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", usersRows, tableName)
	var resp Users

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("users.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUsersModel) FindListByIdList(ctx context.Context, id ...int64) ([]Users, error) {
	if len(id) == 0 {
		return []Users{}, nil
	}
	tableName := "users"

	query := fmt.Sprintf("select %s from %s where id in (%s)", usersRows, tableName, sqlx.InInt64List(id))

	var resp []Users
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Users{}, nil
		}
		return nil, fmt.Errorf("users.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUsersModel) Update2(ctx context.Context, data *Users) error {
	tableName := "users"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, usersRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserType, data.AccessHash, data.SecretKeyId, data.FirstName, data.LastName, data.Username, data.Phone, data.CountryCode, data.Verified, data.Support, data.Scam, data.Fake, data.Premium, data.PremiumExpireDate, data.About, data.State, data.IsBot, data.AccountDaysTtl, data.PhotoId, data.Restricted, data.RestrictionReason, data.ArchiveAndMuteNewNoncontactPeers, data.EmojiStatusDocumentId, data.EmojiStatusUntil, data.StoriesMaxId, data.Color, data.ColorBackgroundEmojiId, data.ProfileColor, data.ProfileColorBackgroundEmojiId, data.Birthday, data.PersonalChannelId, data.AuthorizationTtlDays, data.SavedMusicId, data.MainTab, data.Deleted, data.DeleteReason, data.Id)
	if err != nil {
		return fmt.Errorf("users.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUsersModel) FindOneByPhone(ctx context.Context, phone string) (*Users, error) {
	tableName := "users"
	query := fmt.Sprintf("select %s from %s where phone = ? limit 1", usersRows, tableName)
	var resp Users

	err := m.db.QueryRowPartial(ctx, &resp, query, phone)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("phone=%v", phone),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("users.FindOneByPhone: %w", err)
	}

	return &resp, nil
}

func (m *defaultUsersModel) FindListByPhoneList(ctx context.Context, phone ...string) ([]Users, error) {
	if len(phone) == 0 {
		return []Users{}, nil
	}
	tableName := "users"

	query := fmt.Sprintf("select %s from %s where phone in (%s)", usersRows, tableName, sqlx.InStringList(phone))
	var resp []Users
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Users{}, nil
		}
		return nil, fmt.Errorf("users.FindListByPhoneList: %w", err)
	}

	return resp, nil
}
