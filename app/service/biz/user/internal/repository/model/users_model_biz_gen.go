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

type bizUsersModel interface {
	Insert(ctx context.Context, data *Users) (lastInsertId, rowsAffected int64, err error)
	InsertTestUser(ctx context.Context, data *Users) (lastInsertId, rowsAffected int64, err error)
	SelectByPhoneNumber(ctx context.Context, phone string) (*Users, error)
	SelectById(ctx context.Context, id int64) (*Users, error)
	SelectNextTestUserId(ctx context.Context, maxId int64) (*Users, error)
	SelectUsersByIdList(ctx context.Context, idList []int64) ([]Users, error)
	SelectUsersByIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *Users)) ([]Users, error)
	SelectUsersByPhoneList(ctx context.Context, phoneList []string) ([]Users, error)
	SelectUsersByPhoneListWithCB(ctx context.Context, phoneList []string, cb func(sz, i int, v *Users)) ([]Users, error)
	SearchByQueryString(ctx context.Context, q string, q2 string, idList []int64, limit int32) ([]int64, error)
	SearchByQueryStringWithCB(ctx context.Context, q string, q2 string, idList []int64, limit int32, cb func(sz, i int, v int64)) ([]int64, error)
	SearchByQueryNotIdList(ctx context.Context, q2 string, idList []int64, limit int32) ([]Users, error)
	SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, idList []int64, limit int32, cb func(sz, i int, v *Users)) ([]Users, error)
	Delete(ctx context.Context, phone string, deleteReason string, id int64) (rowsAffected int64, err error)
	UpdateUsername(ctx context.Context, username string, id int64) (rowsAffected int64, err error)
	UpdateFirstAndLastName(ctx context.Context, firstName string, lastName string, id int64) (rowsAffected int64, err error)
	UpdateAbout(ctx context.Context, about string, id int64) (rowsAffected int64, err error)
	UpdateProfile(ctx context.Context, firstName string, lastName string, about string, id int64) (rowsAffected int64, err error)
	SelectByUsername(ctx context.Context, username string) (*Users, error)
	SelectAccountDaysTTL(ctx context.Context, id int64) (*Users, error)
	UpdateAccountDaysTTL(ctx context.Context, accountDaysTtl int32, id int64) (rowsAffected int64, err error)
	SelectProfilePhoto(ctx context.Context, id int64) (int64, error)
	SelectCountryCode(ctx context.Context, id int64) (*Users, error)
	UpdateProfilePhoto(ctx context.Context, photoId int64, id int64) (rowsAffected int64, err error)
	UpdateUser(ctx context.Context, cMap map[string]interface{}, id int64) (rowsAffected int64, err error)
	UpdateEmojiStatus(ctx context.Context, emojiStatusDocumentId int64, emojiStatusUntil int32, id int64) (rowsAffected int64, err error)
	UpdateStoriesMaxId(ctx context.Context, storiesMaxId int32, id int64) (rowsAffected int64, err error)
	UpdateColor(ctx context.Context, color int32, colorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error)
	UpdateProfileColor(ctx context.Context, profileColor int32, profileColorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error)
	QueryChannelParticipants(ctx context.Context, channelId int64, q1 string, q2 string, q3 string) ([]Users, error)
	QueryChannelParticipantsWithCB(ctx context.Context, channelId int64, q1 string, q2 string, q3 string, cb func(sz, i int, v *Users)) ([]Users, error)
	SelectBots(ctx context.Context, idList []int64) ([]int64, error)
	SelectBotsWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v int64)) ([]int64, error)
	UpdateBirthday(ctx context.Context, birthday string, id int64) (rowsAffected int64, err error)
	UpdatePersonalChannelId(ctx context.Context, personalChannelId int64, id int64) (rowsAffected int64, err error)
	UpdateAuthorizationTTL(ctx context.Context, authorizationTtlDays int32, id int64) (rowsAffected int64, err error)
	SelectAuthorizationTTL(ctx context.Context, id int64) (*Users, error)
	UpdateSavedMusicId(ctx context.Context, savedMusicId int64, id int64) (rowsAffected int64, err error)
	UpdateMainTab(ctx context.Context, mainTab int32, id int64) (rowsAffected int64, err error)
}

type UsersTxModel interface {
	Insert(data *Users) (lastInsertId, rowsAffected int64, err error)
	InsertTestUser(data *Users) (lastInsertId, rowsAffected int64, err error)
	SelectByPhoneNumber(phone string) (*Users, error)
	SelectById(id int64) (*Users, error)
	SelectNextTestUserId(maxId int64) (*Users, error)
	SelectUsersByIdList(idList []int64) ([]Users, error)
	SelectUsersByPhoneList(phoneList []string) ([]Users, error)
	SearchByQueryString(q string, q2 string, idList []int64, limit int32) ([]int64, error)
	SearchByQueryNotIdList(q2 string, idList []int64, limit int32) ([]Users, error)
	Delete(phone string, deleteReason string, id int64) (rowsAffected int64, err error)
	UpdateUsername(username string, id int64) (rowsAffected int64, err error)
	UpdateFirstAndLastName(firstName string, lastName string, id int64) (rowsAffected int64, err error)
	UpdateAbout(about string, id int64) (rowsAffected int64, err error)
	UpdateProfile(firstName string, lastName string, about string, id int64) (rowsAffected int64, err error)
	SelectByUsername(username string) (*Users, error)
	SelectAccountDaysTTL(id int64) (*Users, error)
	UpdateAccountDaysTTL(accountDaysTtl int32, id int64) (rowsAffected int64, err error)
	SelectProfilePhoto(id int64) (int64, error)
	SelectCountryCode(id int64) (*Users, error)
	UpdateProfilePhoto(photoId int64, id int64) (rowsAffected int64, err error)
	UpdateUser(cMap map[string]interface{}, id int64) (rowsAffected int64, err error)
	UpdateEmojiStatus(emojiStatusDocumentId int64, emojiStatusUntil int32, id int64) (rowsAffected int64, err error)
	UpdateStoriesMaxId(storiesMaxId int32, id int64) (rowsAffected int64, err error)
	UpdateColor(color int32, colorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error)
	UpdateProfileColor(profileColor int32, profileColorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error)
	QueryChannelParticipants(channelId int64, q1 string, q2 string, q3 string) ([]Users, error)
	SelectBots(idList []int64) ([]int64, error)
	UpdateBirthday(birthday string, id int64) (rowsAffected int64, err error)
	UpdatePersonalChannelId(personalChannelId int64, id int64) (rowsAffected int64, err error)
	UpdateAuthorizationTTL(authorizationTtlDays int32, id int64) (rowsAffected int64, err error)
	SelectAuthorizationTTL(id int64) (*Users, error)
	UpdateSavedMusicId(savedMusicId int64, id int64) (rowsAffected int64, err error)
	UpdateMainTab(mainTab int32, id int64) (rowsAffected int64, err error)
}

type defaultUsersTxModel struct {
	tx *sqlx.Tx
}

func NewUsersTxModel(tx *sqlx.Tx) UsersTxModel {
	return &defaultUsersTxModel{tx: tx}
}

// Insert
// insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("users.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("users.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
func (m *defaultUsersTxModel) Insert(data *Users) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("users.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("users.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.Insert rows affected: %w", err)
	}

	return
}

// InsertTestUser
// insert into users(id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:id, :user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
func (m *defaultUsersModel) InsertTestUser(ctx context.Context, data *Users) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:id, :user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("users.InsertTestUser named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("users.InsertTestUser last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.InsertTestUser rows affected: %w", err)
	}

	return

}

// InsertTestUser
// insert into users(id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:id, :user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
func (m *defaultUsersTxModel) InsertTestUser(data *Users) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:id, :user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("users.InsertTestUser named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("users.InsertTestUser last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.InsertTestUser rows affected: %w", err)
	}

	return
}

// SelectByPhoneNumber
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone = :phone limit 1
func (m *defaultUsersModel) SelectByPhoneNumber(ctx context.Context, phone string) (rValue *Users, err error) {

	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone = ? limit 1"
		do    = &Users{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, phone)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("phone=%v", phone),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectByPhoneNumber: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByPhoneNumber
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone = :phone limit 1
func (m *defaultUsersTxModel) SelectByPhoneNumber(phone string) (rValue *Users, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone = ? limit 1"
		do    = &Users{}
	)
	err = m.tx.QueryRowPartial(do, query, phone)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("phone=%v", phone),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectByPhoneNumber: %w", err)
		return
	}
	rValue = do

	return
}

// SelectById
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id = :id limit 1
func (m *defaultUsersModel) SelectById(ctx context.Context, id int64) (rValue *Users, err error) {

	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id = ? limit 1"
		do    = &Users{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectById: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectById
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id = :id limit 1
func (m *defaultUsersTxModel) SelectById(id int64) (rValue *Users, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id = ? limit 1"
		do    = &Users{}
	)
	err = m.tx.QueryRowPartial(do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectById: %w", err)
		return
	}
	rValue = do

	return
}

// SelectNextTestUserId
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id < :maxId order by id desc limit 1
func (m *defaultUsersModel) SelectNextTestUserId(ctx context.Context, maxId int64) (rValue *Users, err error) {

	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id < ? order by id desc limit 1"
		do    = &Users{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, maxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("maxId=%v", maxId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectNextTestUserId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectNextTestUserId
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id < :maxId order by id desc limit 1
func (m *defaultUsersTxModel) SelectNextTestUserId(maxId int64) (rValue *Users, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id < ? order by id desc limit 1"
		do    = &Users{}
	)
	err = m.tx.QueryRowPartial(do, query, maxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("maxId=%v", maxId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectNextTestUserId: %w", err)
		return
	}
	rValue = do

	return
}

// SelectUsersByIdList
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (:id_list)
func (m *defaultUsersModel) SelectUsersByIdList(ctx context.Context, idList []int64) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (%s)", sqlx.InInt64List(idList))
		values []Users
	)
	if len(idList) == 0 {
		rList = []Users{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectUsersByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectUsersByIdList
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (:id_list)
func (m *defaultUsersTxModel) SelectUsersByIdList(idList []int64) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (%s)", sqlx.InInt64List(idList))
		values []Users
	)
	if len(idList) == 0 {
		rList = []Users{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectUsersByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectUsersByIdListWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (:id_list)
func (m *defaultUsersModel) SelectUsersByIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *Users)) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (%s)", sqlx.InInt64List(idList))
		values []Users
	)
	if len(idList) == 0 {
		rList = []Users{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectUsersByIdListWithCB: %w", err)
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

// SelectUsersByPhoneList
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone in (:phoneList)
func (m *defaultUsersModel) SelectUsersByPhoneList(ctx context.Context, phoneList []string) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone in (%s)", sqlx.InStringList(phoneList))
		values []Users
	)
	if len(phoneList) == 0 {
		rList = []Users{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectUsersByPhoneList: %w", err)
		return
	}

	rList = values

	return
}

// SelectUsersByPhoneList
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone in (:phoneList)
func (m *defaultUsersTxModel) SelectUsersByPhoneList(phoneList []string) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone in (%s)", sqlx.InStringList(phoneList))
		values []Users
	)
	if len(phoneList) == 0 {
		rList = []Users{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectUsersByPhoneList: %w", err)
		return
	}

	rList = values

	return
}

// SelectUsersByPhoneListWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone in (:phoneList)
func (m *defaultUsersModel) SelectUsersByPhoneListWithCB(ctx context.Context, phoneList []string, cb func(sz, i int, v *Users)) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where phone in (%s)", sqlx.InStringList(phoneList))
		values []Users
	)
	if len(phoneList) == 0 {
		rList = []Users{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectUsersByPhoneListWithCB: %w", err)
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

// SearchByQueryString
// select id from users where (username like :q or first_name like :q2 or last_name like :q2) and id not in (:id_list) and deleted = 0 limit :limit
func (m *defaultUsersModel) SearchByQueryString(ctx context.Context, q string, q2 string, idList []int64, limit int32) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where (username like ? or first_name like ? or last_name like ?) and id not in (%s) and deleted = 0 limit ?", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &rList, query, q, q2, q2, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SearchByQueryString: %w", err)
	}

	return
}

// SearchByQueryString
// select id from users where (username like :q or first_name like :q2 or last_name like :q2) and id not in (:id_list) and deleted = 0 limit :limit
func (m *defaultUsersTxModel) SearchByQueryString(q string, q2 string, idList []int64, limit int32) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where (username like ? or first_name like ? or last_name like ?) and id not in (%s) and deleted = 0 limit ?", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.tx.QueryRowsPartial(&rList, query, q, q2, q2, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SearchByQueryString: %w", err)
	}

	return
}

// SearchByQueryStringWithCB
// select id from users where (username like :q or first_name like :q2 or last_name like :q2) and id not in (:id_list) and deleted = 0 limit :limit
func (m *defaultUsersModel) SearchByQueryStringWithCB(ctx context.Context, q string, q2 string, idList []int64, limit int32, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where (username like ? or first_name like ? or last_name like ?) and id not in (%s) and deleted = 0 limit ?", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &rList, query, q, q2, q2, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SearchByQueryStringWithCB: %w", err)
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

// SearchByQueryNotIdList
// select id from users where username like :q2 and id not in (:id_list) and deleted = 0 limit :limit
func (m *defaultUsersModel) SearchByQueryNotIdList(ctx context.Context, q2 string, idList []int64, limit int32) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id from users where username like ? and id not in (%s) and deleted = 0 limit ?", sqlx.InInt64List(idList))
		values []Users
	)
	if len(idList) == 0 {
		rList = []Users{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, q2, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SearchByQueryNotIdList: %w", err)
		return
	}

	rList = values

	return
}

// SearchByQueryNotIdList
// select id from users where username like :q2 and id not in (:id_list) and deleted = 0 limit :limit
func (m *defaultUsersTxModel) SearchByQueryNotIdList(q2 string, idList []int64, limit int32) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id from users where username like ? and id not in (%s) and deleted = 0 limit ?", sqlx.InInt64List(idList))
		values []Users
	)
	if len(idList) == 0 {
		rList = []Users{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, q2, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SearchByQueryNotIdList: %w", err)
		return
	}

	rList = values

	return
}

// SearchByQueryNotIdListWithCB
// select id from users where username like :q2 and id not in (:id_list) and deleted = 0 limit :limit
func (m *defaultUsersModel) SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, idList []int64, limit int32, cb func(sz, i int, v *Users)) (rList []Users, err error) {
	var (
		query  = fmt.Sprintf("select id from users where username like ? and id not in (%s) and deleted = 0 limit ?", sqlx.InInt64List(idList))
		values []Users
	)
	if len(idList) == 0 {
		rList = []Users{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, q2, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SearchByQueryNotIdListWithCB: %w", err)
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
// update users set phone = :phone, deleted = 1, delete_reason = :delete_reason where id = :id
func (m *defaultUsersModel) Delete(ctx context.Context, phone string, deleteReason string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set phone = ?, deleted = 1, delete_reason = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, phone, deleteReason, id)

	if err != nil {
		err = fmt.Errorf("users.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.Delete rows affected: %w", err)
		return
	}

	return
}

// Delete
// update users set phone = :phone, deleted = 1, delete_reason = :delete_reason where id = :id
func (m *defaultUsersTxModel) Delete(phone string, deleteReason string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set phone = ?, deleted = 1, delete_reason = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, phone, deleteReason, id)

	if err != nil {
		err = fmt.Errorf("users.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.Delete rows affected: %w", err)
		return
	}

	return
}

// UpdateUsername
// update users set username = :username where id = :id
func (m *defaultUsersModel) UpdateUsername(ctx context.Context, username string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set username = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, username, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateUsername exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateUsername rows affected: %w", err)
		return
	}

	return
}

// UpdateUsername
// update users set username = :username where id = :id
func (m *defaultUsersTxModel) UpdateUsername(username string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set username = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, username, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateUsername exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateUsername rows affected: %w", err)
		return
	}

	return
}

// UpdateFirstAndLastName
// update users set first_name = :first_name, last_name = :last_name where id = :id
func (m *defaultUsersModel) UpdateFirstAndLastName(ctx context.Context, firstName string, lastName string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set first_name = ?, last_name = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, firstName, lastName, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateFirstAndLastName exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateFirstAndLastName rows affected: %w", err)
		return
	}

	return
}

// UpdateFirstAndLastName
// update users set first_name = :first_name, last_name = :last_name where id = :id
func (m *defaultUsersTxModel) UpdateFirstAndLastName(firstName string, lastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, firstName, lastName, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateFirstAndLastName exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateFirstAndLastName rows affected: %w", err)
		return
	}

	return
}

// UpdateAbout
// update users set about = :about where id = :id
func (m *defaultUsersModel) UpdateAbout(ctx context.Context, about string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set about = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, about, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateAbout exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateAbout rows affected: %w", err)
		return
	}

	return
}

// UpdateAbout
// update users set about = :about where id = :id
func (m *defaultUsersTxModel) UpdateAbout(about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, about, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateAbout exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateAbout rows affected: %w", err)
		return
	}

	return
}

// UpdateProfile
// update users set first_name = :first_name, last_name = :last_name, about = :about where id = :id
func (m *defaultUsersModel) UpdateProfile(ctx context.Context, firstName string, lastName string, about string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, firstName, lastName, about, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateProfile exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateProfile rows affected: %w", err)
		return
	}

	return
}

// UpdateProfile
// update users set first_name = :first_name, last_name = :last_name, about = :about where id = :id
func (m *defaultUsersTxModel) UpdateProfile(firstName string, lastName string, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, firstName, lastName, about, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateProfile exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateProfile rows affected: %w", err)
		return
	}

	return
}

// SelectByUsername
// select id from users where username = :username limit 1
func (m *defaultUsersModel) SelectByUsername(ctx context.Context, username string) (rValue *Users, err error) {

	var (
		query = "select id from users where username = ? limit 1"
		do    = &Users{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, username)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("username=%v", username),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectByUsername: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUsername
// select id from users where username = :username limit 1
func (m *defaultUsersTxModel) SelectByUsername(username string) (rValue *Users, err error) {
	var (
		query = "select id from users where username = ? limit 1"
		do    = &Users{}
	)
	err = m.tx.QueryRowPartial(do, query, username)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("username=%v", username),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectByUsername: %w", err)
		return
	}
	rValue = do

	return
}

// SelectAccountDaysTTL
// select account_days_ttl from users where id = :id
func (m *defaultUsersModel) SelectAccountDaysTTL(ctx context.Context, id int64) (rValue *Users, err error) {

	var (
		query = "select account_days_ttl from users where id = ?"
		do    = &Users{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectAccountDaysTTL: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectAccountDaysTTL
// select account_days_ttl from users where id = :id
func (m *defaultUsersTxModel) SelectAccountDaysTTL(id int64) (rValue *Users, err error) {
	var (
		query = "select account_days_ttl from users where id = ?"
		do    = &Users{}
	)
	err = m.tx.QueryRowPartial(do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectAccountDaysTTL: %w", err)
		return
	}
	rValue = do

	return
}

// UpdateAccountDaysTTL
// update users set account_days_ttl = :account_days_ttl where id = :id
func (m *defaultUsersModel) UpdateAccountDaysTTL(ctx context.Context, accountDaysTtl int32, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set account_days_ttl = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, accountDaysTtl, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateAccountDaysTTL exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateAccountDaysTTL rows affected: %w", err)
		return
	}

	return
}

// UpdateAccountDaysTTL
// update users set account_days_ttl = :account_days_ttl where id = :id
func (m *defaultUsersTxModel) UpdateAccountDaysTTL(accountDaysTtl int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set account_days_ttl = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, accountDaysTtl, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateAccountDaysTTL exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateAccountDaysTTL rows affected: %w", err)
		return
	}

	return
}

// SelectProfilePhoto
// select photo_id from users where id = :id limit 1
func (m *defaultUsersModel) SelectProfilePhoto(ctx context.Context, id int64) (rValue int64, err error) {
	var query = "select photo_id from users where id = ? limit 1"
	err = m.db.QueryRowPartial(ctx, &rValue, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			err = &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
			return
		}
		err = fmt.Errorf("users.SelectProfilePhoto: %w", err)
		return
	}

	return
}

// SelectProfilePhoto
// select photo_id from users where id = :id limit 1
func (m *defaultUsersTxModel) SelectProfilePhoto(id int64) (rValue int64, err error) {
	var query = "select photo_id from users where id = ? limit 1"
	err = m.tx.QueryRowPartial(&rValue, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			err = &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
			return
		}
		err = fmt.Errorf("users.SelectProfilePhoto: %w", err)
		return
	}

	return
}

// SelectCountryCode
// select country_code from users where id = :id
func (m *defaultUsersModel) SelectCountryCode(ctx context.Context, id int64) (rValue *Users, err error) {

	var (
		query = "select country_code from users where id = ?"
		do    = &Users{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectCountryCode: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectCountryCode
// select country_code from users where id = :id
func (m *defaultUsersTxModel) SelectCountryCode(id int64) (rValue *Users, err error) {
	var (
		query = "select country_code from users where id = ?"
		do    = &Users{}
	)
	err = m.tx.QueryRowPartial(do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectCountryCode: %w", err)
		return
	}
	rValue = do

	return
}

// UpdateProfilePhoto
// update users set photo_id = :photo_id where id = :id
func (m *defaultUsersModel) UpdateProfilePhoto(ctx context.Context, photoId int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set photo_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, photoId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateProfilePhoto exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateProfilePhoto rows affected: %w", err)
		return
	}

	return
}

// UpdateProfilePhoto
// update users set photo_id = :photo_id where id = :id
func (m *defaultUsersTxModel) UpdateProfilePhoto(photoId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set photo_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, photoId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateProfilePhoto exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateProfilePhoto rows affected: %w", err)
		return
	}

	return
}

// UpdateUser
// update users set %s where id = :id
func (m *defaultUsersModel) UpdateUser(ctx context.Context, cMap map[string]interface{}, id int64) (rowsAffected int64, err error) {

	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update users set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		err = fmt.Errorf("users.UpdateUser exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateUser rows affected: %w", err)
		return
	}

	return
}

// UpdateUser
// update users set %s where id = :id
func (m *defaultUsersTxModel) UpdateUser(cMap map[string]interface{}, id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update users set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)

	rResult, err = m.tx.Exec(query, aValues...)

	if err != nil {
		err = fmt.Errorf("users.UpdateUser exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateUser rows affected: %w", err)
		return
	}

	return
}

// UpdateEmojiStatus
// update users set emoji_status_document_id = :emoji_status_document_id, emoji_status_until = :emoji_status_until where id = :id
func (m *defaultUsersModel) UpdateEmojiStatus(ctx context.Context, emojiStatusDocumentId int64, emojiStatusUntil int32, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set emoji_status_document_id = ?, emoji_status_until = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, emojiStatusDocumentId, emojiStatusUntil, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateEmojiStatus exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateEmojiStatus rows affected: %w", err)
		return
	}

	return
}

// UpdateEmojiStatus
// update users set emoji_status_document_id = :emoji_status_document_id, emoji_status_until = :emoji_status_until where id = :id
func (m *defaultUsersTxModel) UpdateEmojiStatus(emojiStatusDocumentId int64, emojiStatusUntil int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set emoji_status_document_id = ?, emoji_status_until = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, emojiStatusDocumentId, emojiStatusUntil, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateEmojiStatus exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateEmojiStatus rows affected: %w", err)
		return
	}

	return
}

// UpdateStoriesMaxId
// update users set stories_max_id = :stories_max_id where id = :id
func (m *defaultUsersModel) UpdateStoriesMaxId(ctx context.Context, storiesMaxId int32, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set stories_max_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, storiesMaxId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateStoriesMaxId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateStoriesMaxId rows affected: %w", err)
		return
	}

	return
}

// UpdateStoriesMaxId
// update users set stories_max_id = :stories_max_id where id = :id
func (m *defaultUsersTxModel) UpdateStoriesMaxId(storiesMaxId int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set stories_max_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, storiesMaxId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateStoriesMaxId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateStoriesMaxId rows affected: %w", err)
		return
	}

	return
}

// UpdateColor
// update users set color = :color, color_background_emoji_id = :color_background_emoji_id where id = :id
func (m *defaultUsersModel) UpdateColor(ctx context.Context, color int32, colorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set color = ?, color_background_emoji_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, color, colorBackgroundEmojiId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateColor exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateColor rows affected: %w", err)
		return
	}

	return
}

// UpdateColor
// update users set color = :color, color_background_emoji_id = :color_background_emoji_id where id = :id
func (m *defaultUsersTxModel) UpdateColor(color int32, colorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set color = ?, color_background_emoji_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, color, colorBackgroundEmojiId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateColor exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateColor rows affected: %w", err)
		return
	}

	return
}

// UpdateProfileColor
// update users set profile_color = :profile_color, profile_color_background_emoji_id = :profile_color_background_emoji_id where id = :id
func (m *defaultUsersModel) UpdateProfileColor(ctx context.Context, profileColor int32, profileColorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set profile_color = ?, profile_color_background_emoji_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, profileColor, profileColorBackgroundEmojiId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateProfileColor exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateProfileColor rows affected: %w", err)
		return
	}

	return
}

// UpdateProfileColor
// update users set profile_color = :profile_color, profile_color_background_emoji_id = :profile_color_background_emoji_id where id = :id
func (m *defaultUsersTxModel) UpdateProfileColor(profileColor int32, profileColorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set profile_color = ?, profile_color_background_emoji_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, profileColor, profileColorBackgroundEmojiId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateProfileColor exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateProfileColor rows affected: %w", err)
		return
	}

	return
}

// QueryChannelParticipants
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = :channelId and state = 0) and (first_name like :q1 or last_name like :q2 or username like :q3)
func (m *defaultUsersModel) QueryChannelParticipants(ctx context.Context, channelId int64, q1 string, q2 string, q3 string) (rList []Users, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = ? and state = 0) and (first_name like ? or last_name like ? or username like ?)"
		values []Users
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, channelId, q1, q2, q3)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.QueryChannelParticipants: %w", err)
		return
	}

	rList = values

	return
}

// QueryChannelParticipants
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = :channelId and state = 0) and (first_name like :q1 or last_name like :q2 or username like :q3)
func (m *defaultUsersTxModel) QueryChannelParticipants(channelId int64, q1 string, q2 string, q3 string) (rList []Users, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = ? and state = 0) and (first_name like ? or last_name like ? or username like ?)"
		values []Users
	)
	err = m.tx.QueryRowsPartial(&values, query, channelId, q1, q2, q3)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.QueryChannelParticipants: %w", err)
		return
	}

	rList = values

	return
}

// QueryChannelParticipantsWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = :channelId and state = 0) and (first_name like :q1 or last_name like :q2 or username like :q3)
func (m *defaultUsersModel) QueryChannelParticipantsWithCB(ctx context.Context, channelId int64, q1 string, q2 string, q3 string, cb func(sz, i int, v *Users)) (rList []Users, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, premium_expire_date, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, personal_channel_id, saved_music_id, main_tab, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = ? and state = 0) and (first_name like ? or last_name like ? or username like ?)"
		values []Users
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, channelId, q1, q2, q3)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Users{}
			err = nil
			return
		}
		err = fmt.Errorf("users.QueryChannelParticipantsWithCB: %w", err)
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

// SelectBots
// select id from users where id in (:id_list) and is_bot = 1
func (m *defaultUsersModel) SelectBots(ctx context.Context, idList []int64) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where id in (%s) and is_bot = 1", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &rList, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectBots: %w", err)
	}

	return
}

// SelectBots
// select id from users where id in (:id_list) and is_bot = 1
func (m *defaultUsersTxModel) SelectBots(idList []int64) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where id in (%s) and is_bot = 1", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.tx.QueryRowsPartial(&rList, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectBots: %w", err)
	}

	return
}

// SelectBotsWithCB
// select id from users where id in (:id_list) and is_bot = 1
func (m *defaultUsersModel) SelectBotsWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where id in (%s) and is_bot = 1", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &rList, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("users.SelectBotsWithCB: %w", err)
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

// UpdateBirthday
// update users set birthday = :birthday where id = :id
func (m *defaultUsersModel) UpdateBirthday(ctx context.Context, birthday string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set birthday = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, birthday, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateBirthday exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateBirthday rows affected: %w", err)
		return
	}

	return
}

// UpdateBirthday
// update users set birthday = :birthday where id = :id
func (m *defaultUsersTxModel) UpdateBirthday(birthday string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set birthday = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, birthday, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateBirthday exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateBirthday rows affected: %w", err)
		return
	}

	return
}

// UpdatePersonalChannelId
// update users set personal_channel_id = :personal_channel_id where id = :id
func (m *defaultUsersModel) UpdatePersonalChannelId(ctx context.Context, personalChannelId int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set personal_channel_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, personalChannelId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdatePersonalChannelId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdatePersonalChannelId rows affected: %w", err)
		return
	}

	return
}

// UpdatePersonalChannelId
// update users set personal_channel_id = :personal_channel_id where id = :id
func (m *defaultUsersTxModel) UpdatePersonalChannelId(personalChannelId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set personal_channel_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, personalChannelId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdatePersonalChannelId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdatePersonalChannelId rows affected: %w", err)
		return
	}

	return
}

// UpdateAuthorizationTTL
// update users set authorization_ttl_days = :authorization_ttl_days where id = :id
func (m *defaultUsersModel) UpdateAuthorizationTTL(ctx context.Context, authorizationTtlDays int32, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set authorization_ttl_days = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, authorizationTtlDays, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateAuthorizationTTL exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateAuthorizationTTL rows affected: %w", err)
		return
	}

	return
}

// UpdateAuthorizationTTL
// update users set authorization_ttl_days = :authorization_ttl_days where id = :id
func (m *defaultUsersTxModel) UpdateAuthorizationTTL(authorizationTtlDays int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set authorization_ttl_days = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, authorizationTtlDays, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateAuthorizationTTL exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateAuthorizationTTL rows affected: %w", err)
		return
	}

	return
}

// SelectAuthorizationTTL
// select authorization_ttl_days from users where id = :id
func (m *defaultUsersModel) SelectAuthorizationTTL(ctx context.Context, id int64) (rValue *Users, err error) {

	var (
		query = "select authorization_ttl_days from users where id = ?"
		do    = &Users{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectAuthorizationTTL: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectAuthorizationTTL
// select authorization_ttl_days from users where id = :id
func (m *defaultUsersTxModel) SelectAuthorizationTTL(id int64) (rValue *Users, err error) {
	var (
		query = "select authorization_ttl_days from users where id = ?"
		do    = &Users{}
	)
	err = m.tx.QueryRowPartial(do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("users.SelectAuthorizationTTL: %w", err)
		return
	}
	rValue = do

	return
}

// UpdateSavedMusicId
// update users set saved_music_id = :saved_music_id where id = :id
func (m *defaultUsersModel) UpdateSavedMusicId(ctx context.Context, savedMusicId int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set saved_music_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, savedMusicId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateSavedMusicId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateSavedMusicId rows affected: %w", err)
		return
	}

	return
}

// UpdateSavedMusicId
// update users set saved_music_id = :saved_music_id where id = :id
func (m *defaultUsersTxModel) UpdateSavedMusicId(savedMusicId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set saved_music_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, savedMusicId, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateSavedMusicId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateSavedMusicId rows affected: %w", err)
		return
	}

	return
}

// UpdateMainTab
// update users set main_tab = :main_tab where id = :id
func (m *defaultUsersModel) UpdateMainTab(ctx context.Context, mainTab int32, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update users set main_tab = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, mainTab, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateMainTab exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateMainTab rows affected: %w", err)
		return
	}

	return
}

// UpdateMainTab
// update users set main_tab = :main_tab where id = :id
func (m *defaultUsersTxModel) UpdateMainTab(mainTab int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set main_tab = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, mainTab, id)

	if err != nil {
		err = fmt.Errorf("users.UpdateMainTab exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("users.UpdateMainTab rows affected: %w", err)
		return
	}

	return
}
