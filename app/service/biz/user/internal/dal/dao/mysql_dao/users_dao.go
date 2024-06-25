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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type UsersDAO struct {
	db *sqlx.DB
}

func NewUsersDAO(db *sqlx.DB) *UsersDAO {
	return &UsersDAO{
		db: db,
	}
}

// Insert
// insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
func (dao *UsersDAO) Insert(ctx context.Context, do *dataobject.UsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
func (dao *UsersDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTestUser
// insert into users(id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:id, :user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
func (dao *UsersDAO) InsertTestUser(ctx context.Context, do *dataobject.UsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:id, :user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertTestUser(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertTestUser(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertTestUser(%v)_error: %v", do, err)
	}

	return
}

// InsertTestUserTx
// insert into users(id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:id, :user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
func (dao *UsersDAO) InsertTestUserTx(tx *sqlx.Tx, do *dataobject.UsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:id, :user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertTestUser(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertTestUser(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertTestUser(%v)_error: %v", do, err)
	}

	return
}

// SelectByPhoneNumber
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where phone = :phone limit 1
func (dao *UsersDAO) SelectByPhoneNumber(ctx context.Context, phone string) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where phone = ? limit 1"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, phone)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByPhoneNumber(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectById
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id = :id limit 1
func (dao *UsersDAO) SelectById(ctx context.Context, id int64) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id = ? limit 1"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectById(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectNextTestUserId
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id < :maxId order by id desc limit 1
func (dao *UsersDAO) SelectNextTestUserId(ctx context.Context, maxId int64) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id < ? order by id desc limit 1"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, maxId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectNextTestUserId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectUsersByIdList
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id in (:id_list)
func (dao *UsersDAO) SelectUsersByIdList(ctx context.Context, idList []int64) (rList []dataobject.UsersDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.UsersDO
	)
	if len(idList) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUsersByIdListWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id in (:id_list)
func (dao *UsersDAO) SelectUsersByIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.UsersDO
	)
	if len(idList) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersByIdList(_), error: %v", err)
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
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where phone in (:phoneList)
func (dao *UsersDAO) SelectUsersByPhoneList(ctx context.Context, phoneList []string) (rList []dataobject.UsersDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where phone in (%s)", sqlx.InStringList(phoneList))
		values []dataobject.UsersDO
	)
	if len(phoneList) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersByPhoneList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUsersByPhoneListWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where phone in (:phoneList)
func (dao *UsersDAO) SelectUsersByPhoneListWithCB(ctx context.Context, phoneList []string, cb func(sz, i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where phone in (%s)", sqlx.InStringList(phoneList))
		values []dataobject.UsersDO
	)
	if len(phoneList) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersByPhoneList(_), error: %v", err)
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
// select id from users where (username like :q or first_name like :q2 or last_name like :q2) and id not in (:id_list) limit :limit
func (dao *UsersDAO) SearchByQueryString(ctx context.Context, q string, q2 string, idList []int64, limit int32) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where (username like ? or first_name like ? or last_name like ?) and id not in (%s) limit ?", sqlx.InInt64List(idList))
	)

	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &rList, query, q, q2, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SearchByQueryString(_), error: %v", err)
	}

	return
}

// SearchByQueryStringWithCB
// select id from users where (username like :q or first_name like :q2 or last_name like :q2) and id not in (:id_list) limit :limit
func (dao *UsersDAO) SearchByQueryStringWithCB(ctx context.Context, q string, q2 string, idList []int64, limit int32, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where (username like ? or first_name like ? or last_name like ?) and id not in (%s) limit ?", sqlx.InInt64List(idList))
	)

	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &rList, query, q, q2, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SearchByQueryString(_), error: %v", err)
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
// select id from users where username like :q2 and id not in (:id_list) limit :limit
func (dao *UsersDAO) SearchByQueryNotIdList(ctx context.Context, q2 string, idList []int64, limit int32) (rList []dataobject.UsersDO, err error) {
	var (
		query  = fmt.Sprintf("select id from users where username like ? and id not in (%s) limit ?", sqlx.InInt64List(idList))
		values []dataobject.UsersDO
	)

	if len(idList) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchByQueryNotIdListWithCB
// select id from users where username like :q2 and id not in (:id_list) limit :limit
func (dao *UsersDAO) SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, idList []int64, limit int32, cb func(sz, i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = fmt.Sprintf("select id from users where username like ? and id not in (%s) limit ?", sqlx.InInt64List(idList))
		values []dataobject.UsersDO
	)

	if len(idList) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
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
func (dao *UsersDAO) Delete(ctx context.Context, phone string, deleteReason string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set phone = ?, deleted = 1, delete_reason = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, phone, deleteReason, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// DeleteTx
// update users set phone = :phone, deleted = 1, delete_reason = :delete_reason where id = :id
func (dao *UsersDAO) DeleteTx(tx *sqlx.Tx, phone string, deleteReason string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set phone = ?, deleted = 1, delete_reason = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, phone, deleteReason, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// UpdateUsername
// update users set username = :username where id = :id
func (dao *UsersDAO) UpdateUsername(ctx context.Context, username string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set username = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, username, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateUsername(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateUsername(_), error: %v", err)
	}

	return
}

// UpdateUsernameTx
// update users set username = :username where id = :id
func (dao *UsersDAO) UpdateUsernameTx(tx *sqlx.Tx, username string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set username = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, username, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateUsername(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateUsername(_), error: %v", err)
	}

	return
}

// UpdateFirstAndLastName
// update users set first_name = :first_name, last_name = :last_name where id = :id
func (dao *UsersDAO) UpdateFirstAndLastName(ctx context.Context, firstName string, lastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, firstName, lastName, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateFirstAndLastName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateFirstAndLastName(_), error: %v", err)
	}

	return
}

// UpdateFirstAndLastNameTx
// update users set first_name = :first_name, last_name = :last_name where id = :id
func (dao *UsersDAO) UpdateFirstAndLastNameTx(tx *sqlx.Tx, firstName string, lastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, firstName, lastName, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateFirstAndLastName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateFirstAndLastName(_), error: %v", err)
	}

	return
}

// UpdateAbout
// update users set about = :about where id = :id
func (dao *UsersDAO) UpdateAbout(ctx context.Context, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set about = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, about, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

// UpdateAboutTx
// update users set about = :about where id = :id
func (dao *UsersDAO) UpdateAboutTx(tx *sqlx.Tx, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, about, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

// UpdateProfile
// update users set first_name = :first_name, last_name = :last_name, about = :about where id = :id
func (dao *UsersDAO) UpdateProfile(ctx context.Context, firstName string, lastName string, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, firstName, lastName, about, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateProfile(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateProfile(_), error: %v", err)
	}

	return
}

// UpdateProfileTx
// update users set first_name = :first_name, last_name = :last_name, about = :about where id = :id
func (dao *UsersDAO) UpdateProfileTx(tx *sqlx.Tx, firstName string, lastName string, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, firstName, lastName, about, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateProfile(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateProfile(_), error: %v", err)
	}

	return
}

// SelectByUsername
// select id from users where username = :username limit 1
func (dao *UsersDAO) SelectByUsername(ctx context.Context, username string) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id from users where username = ? limit 1"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, username)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByUsername(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectAccountDaysTTL
// select account_days_ttl from users where id = :id
func (dao *UsersDAO) SelectAccountDaysTTL(ctx context.Context, id int64) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select account_days_ttl from users where id = ?"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectAccountDaysTTL(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// UpdateAccountDaysTTL
// update users set account_days_ttl = :account_days_ttl where id = :id
func (dao *UsersDAO) UpdateAccountDaysTTL(ctx context.Context, accountDaysTtl int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set account_days_ttl = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, accountDaysTtl, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateAccountDaysTTL(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateAccountDaysTTL(_), error: %v", err)
	}

	return
}

// UpdateAccountDaysTTLTx
// update users set account_days_ttl = :account_days_ttl where id = :id
func (dao *UsersDAO) UpdateAccountDaysTTLTx(tx *sqlx.Tx, accountDaysTtl int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set account_days_ttl = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, accountDaysTtl, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateAccountDaysTTL(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateAccountDaysTTL(_), error: %v", err)
	}

	return
}

// SelectProfilePhoto
// select photo_id from users where id = :id limit 1
func (dao *UsersDAO) SelectProfilePhoto(ctx context.Context, id int64) (rValue int64, err error) {
	var query = "select photo_id from users where id = ? limit 1"
	err = dao.db.QueryRowPartial(ctx, &rValue, query, id)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("get in SelectProfilePhoto(_), error: %v", err)
			return
		} else {
			err = nil
		}
	}

	return
}

// SelectCountryCode
// select country_code from users where id = :id
func (dao *UsersDAO) SelectCountryCode(ctx context.Context, id int64) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select country_code from users where id = ?"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectCountryCode(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// UpdateProfilePhoto
// update users set photo_id = :photo_id where id = :id
func (dao *UsersDAO) UpdateProfilePhoto(ctx context.Context, photoId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set photo_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, photoId, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateProfilePhoto(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateProfilePhoto(_), error: %v", err)
	}

	return
}

// UpdateProfilePhotoTx
// update users set photo_id = :photo_id where id = :id
func (dao *UsersDAO) UpdateProfilePhotoTx(tx *sqlx.Tx, photoId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set photo_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, photoId, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateProfilePhoto(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateProfilePhoto(_), error: %v", err)
	}

	return
}

// UpdateUser
// update users set %s where id = :id
func (dao *UsersDAO) UpdateUser(ctx context.Context, cMap map[string]interface{}, id int64) (rowsAffected int64, err error) {
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

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateUser(_), error: %v", err)
	}

	return
}

// UpdateUserTx
// update users set %s where id = :id
func (dao *UsersDAO) UpdateUserTx(tx *sqlx.Tx, cMap map[string]interface{}, id int64) (rowsAffected int64, err error) {
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

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateUser(_), error: %v", err)
	}

	return
}

// UpdateEmojiStatus
// update users set emoji_status_document_id = :emoji_status_document_id, emoji_status_until = :emoji_status_until where id = :id
func (dao *UsersDAO) UpdateEmojiStatus(ctx context.Context, emojiStatusDocumentId int64, emojiStatusUntil int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set emoji_status_document_id = ?, emoji_status_until = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, emojiStatusDocumentId, emojiStatusUntil, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateEmojiStatus(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateEmojiStatus(_), error: %v", err)
	}

	return
}

// UpdateEmojiStatusTx
// update users set emoji_status_document_id = :emoji_status_document_id, emoji_status_until = :emoji_status_until where id = :id
func (dao *UsersDAO) UpdateEmojiStatusTx(tx *sqlx.Tx, emojiStatusDocumentId int64, emojiStatusUntil int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set emoji_status_document_id = ?, emoji_status_until = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, emojiStatusDocumentId, emojiStatusUntil, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateEmojiStatus(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateEmojiStatus(_), error: %v", err)
	}

	return
}

// UpdateStoriesMaxId
// update users set stories_max_id = :stories_max_id where id = :id
func (dao *UsersDAO) UpdateStoriesMaxId(ctx context.Context, storiesMaxId int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set stories_max_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, storiesMaxId, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateStoriesMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateStoriesMaxId(_), error: %v", err)
	}

	return
}

// UpdateStoriesMaxIdTx
// update users set stories_max_id = :stories_max_id where id = :id
func (dao *UsersDAO) UpdateStoriesMaxIdTx(tx *sqlx.Tx, storiesMaxId int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set stories_max_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, storiesMaxId, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateStoriesMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateStoriesMaxId(_), error: %v", err)
	}

	return
}

// UpdateColor
// update users set color = :color, color_background_emoji_id = :color_background_emoji_id where id = :id
func (dao *UsersDAO) UpdateColor(ctx context.Context, color int32, colorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set color = ?, color_background_emoji_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, color, colorBackgroundEmojiId, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateColor(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateColor(_), error: %v", err)
	}

	return
}

// UpdateColorTx
// update users set color = :color, color_background_emoji_id = :color_background_emoji_id where id = :id
func (dao *UsersDAO) UpdateColorTx(tx *sqlx.Tx, color int32, colorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set color = ?, color_background_emoji_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, color, colorBackgroundEmojiId, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateColor(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateColor(_), error: %v", err)
	}

	return
}

// UpdateProfileColor
// update users set profile_color = :profile_color, profile_color_background_emoji_id = :profile_color_background_emoji_id where id = :id
func (dao *UsersDAO) UpdateProfileColor(ctx context.Context, profileColor int32, profileColorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set profile_color = ?, profile_color_background_emoji_id = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, profileColor, profileColorBackgroundEmojiId, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateProfileColor(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateProfileColor(_), error: %v", err)
	}

	return
}

// UpdateProfileColorTx
// update users set profile_color = :profile_color, profile_color_background_emoji_id = :profile_color_background_emoji_id where id = :id
func (dao *UsersDAO) UpdateProfileColorTx(tx *sqlx.Tx, profileColor int32, profileColorBackgroundEmojiId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set profile_color = ?, profile_color_background_emoji_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, profileColor, profileColorBackgroundEmojiId, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateProfileColor(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateProfileColor(_), error: %v", err)
	}

	return
}

// QueryChannelParticipants
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = :channelId and state = 0) and (first_name like :q1 or last_name like :q2 or username like :q3)
func (dao *UsersDAO) QueryChannelParticipants(ctx context.Context, channelId int64, q1 string, q2 string, q3 string) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = ? and state = 0) and (first_name like ? or last_name like ? or username like ?)"
		values []dataobject.UsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, channelId, q1, q2, q3)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in QueryChannelParticipants(_), error: %v", err)
		return
	}

	rList = values

	return
}

// QueryChannelParticipantsWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = :channelId and state = 0) and (first_name like :q1 or last_name like :q2 or username like :q3)
func (dao *UsersDAO) QueryChannelParticipantsWithCB(ctx context.Context, channelId int64, q1 string, q2 string, q3 string, cb func(sz, i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, stories_max_id, color, color_background_emoji_id, profile_color, profile_color_background_emoji_id, birthday, deleted, delete_reason from users where id in (select user_id from channel_participants where channel_id = ? and state = 0) and (first_name like ? or last_name like ? or username like ?)"
		values []dataobject.UsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, channelId, q1, q2, q3)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in QueryChannelParticipants(_), error: %v", err)
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
func (dao *UsersDAO) SelectBots(ctx context.Context, idList []int64) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where id in (%s) and is_bot = 1", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &rList, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectBots(_), error: %v", err)
	}

	return
}

// SelectBotsWithCB
// select id from users where id in (:id_list) and is_bot = 1
func (dao *UsersDAO) SelectBotsWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var (
		query = fmt.Sprintf("select id from users where id in (%s) and is_bot = 1", sqlx.InInt64List(idList))
	)
	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &rList, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectBots(_), error: %v", err)
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
func (dao *UsersDAO) UpdateBirthday(ctx context.Context, birthday string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set birthday = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, birthday, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateBirthday(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateBirthday(_), error: %v", err)
	}

	return
}

// UpdateBirthdayTx
// update users set birthday = :birthday where id = :id
func (dao *UsersDAO) UpdateBirthdayTx(tx *sqlx.Tx, birthday string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set birthday = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, birthday, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateBirthday(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateBirthday(_), error: %v", err)
	}

	return
}
