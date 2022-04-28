/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UsersDAO struct {
	db *sqlx.DB
}

func NewUsersDAO(db *sqlx.DB) *UsersDAO {
	return &UsersDAO{db}
}

// Insert
// insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot)
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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

// SelectByPhoneNumber
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where phone = :phone limit 1
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByPhoneNumber(ctx context.Context, phone string) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where phone = ? limit 1"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, phone)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id = :id limit 1
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectById(ctx context.Context, id int64) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id = ? limit 1"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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

// SelectUsersByIdList
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectUsersByIdList(ctx context.Context, id_list []int64) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id in (?)"
		a      []interface{}
		values []dataobject.UsersDO
	)
	if len(id_list) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectUsersByIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUsersByIdListWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectUsersByIdListWithCB(ctx context.Context, id_list []int64, cb func(i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id in (?)"
		a      []interface{}
		values []dataobject.UsersDO
	)
	if len(id_list) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectUsersByIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersByIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// SelectUsersByPhoneList
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where phone in (:phoneList)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectUsersByPhoneList(ctx context.Context, phoneList []string) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where phone in (?)"
		a      []interface{}
		values []dataobject.UsersDO
	)
	if len(phoneList) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectUsersByPhoneList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersByPhoneList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUsersByPhoneListWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where phone in (:phoneList)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectUsersByPhoneListWithCB(ctx context.Context, phoneList []string, cb func(i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where phone in (?)"
		a      []interface{}
		values []dataobject.UsersDO
	)
	if len(phoneList) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectUsersByPhoneList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUsersByPhoneList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// SelectByQueryString
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where username = :username or first_name = :first_name or last_name = :last_name or phone = :phone limit 20
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByQueryString(ctx context.Context, username string, first_name string, last_name string, phone string) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where username = ? or first_name = ? or last_name = ? or phone = ? limit 20"
		values []dataobject.UsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, username, first_name, last_name, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByQueryString(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByQueryStringWithCB
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where username = :username or first_name = :first_name or last_name = :last_name or phone = :phone limit 20
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByQueryStringWithCB(ctx context.Context, username string, first_name string, last_name string, phone string, cb func(i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where username = ? or first_name = ? or last_name = ? or phone = ? limit 20"
		values []dataobject.UsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, username, first_name, last_name, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByQueryString(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// SearchByQueryNotIdList
// select id from users where username like :q2 and id not in (:id_list) limit :limit
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SearchByQueryNotIdList(ctx context.Context, q2 string, id_list []int64, limit int32) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id from users where username like ? and id not in (?) limit ?"
		a      []interface{}
		values []dataobject.UsersDO
	)

	if len(id_list) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	query, a, err = sqlx.In(query, q2, id_list, limit)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SearchByQueryNotIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchByQueryNotIdListWithCB
// select id from users where username like :q2 and id not in (:id_list) limit :limit
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SearchByQueryNotIdListWithCB(ctx context.Context, q2 string, id_list []int64, limit int32, cb func(i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id from users where username like ? and id not in (?) limit ?"
		a      []interface{}
		values []dataobject.UsersDO
	)

	if len(id_list) == 0 {
		rList = []dataobject.UsersDO{}
		return
	}

	query, a, err = sqlx.In(query, q2, id_list, limit)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SearchByQueryNotIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// Delete
// update users set deleted = 1, delete_reason = :delete_reason where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) Delete(ctx context.Context, delete_reason string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set deleted = 1, delete_reason = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, delete_reason, id)

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

// update users set deleted = 1, delete_reason = :delete_reason where id = :id
// DeleteTx
// TODO(@benqi): sqlmap
func (dao *UsersDAO) DeleteTx(tx *sqlx.Tx, delete_reason string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set deleted = 1, delete_reason = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, delete_reason, id)

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
// TODO(@benqi): sqlmap
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

// update users set username = :username where id = :id
// UpdateUsernameTx
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateFirstAndLastName(ctx context.Context, first_name string, last_name string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, first_name, last_name, id)

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

// update users set first_name = :first_name, last_name = :last_name where id = :id
// UpdateFirstAndLastNameTx
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateFirstAndLastNameTx(tx *sqlx.Tx, first_name string, last_name string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, first_name, last_name, id)

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
// TODO(@benqi): sqlmap
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

// update users set about = :about where id = :id
// UpdateAboutTx
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateProfile(ctx context.Context, first_name string, last_name string, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, first_name, last_name, about, id)

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

// update users set first_name = :first_name, last_name = :last_name, about = :about where id = :id
// UpdateProfileTx
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateProfileTx(tx *sqlx.Tx, first_name string, last_name string, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, first_name, last_name, about, id)

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
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByUsername(ctx context.Context, username string) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id from users where username = ? limit 1"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, username)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectAccountDaysTTL(ctx context.Context, id int64) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select account_days_ttl from users where id = ?"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateAccountDaysTTL(ctx context.Context, account_days_ttl int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set account_days_ttl = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, account_days_ttl, id)

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

// update users set account_days_ttl = :account_days_ttl where id = :id
// UpdateAccountDaysTTLTx
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateAccountDaysTTLTx(tx *sqlx.Tx, account_days_ttl int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set account_days_ttl = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, account_days_ttl, id)

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
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectProfilePhoto(ctx context.Context, id int64) (rValue int64, err error) {
	var query = "select photo_id from users where id = ? limit 1"
	err = dao.db.QueryRowPartial(ctx, &rValue, query, id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectCountryCode(ctx context.Context, id int64) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select country_code from users where id = ?"
		do    = &dataobject.UsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateProfilePhoto(ctx context.Context, photo_id int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set photo_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, photo_id, id)

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

// update users set photo_id = :photo_id where id = :id
// UpdateProfilePhotoTx
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateProfilePhotoTx(tx *sqlx.Tx, photo_id int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update users set photo_id = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, photo_id, id)

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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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

// QueryChannelParticipants
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id in (select user_id from channel_participants where channel_id = :channelId and state = 0) and (first_name like :q1 or last_name like :q2 or username like :q3)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) QueryChannelParticipants(ctx context.Context, channelId int64, q1 string, q2 string, q3 string) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id in (select user_id from channel_participants where channel_id = ? and state = 0) and (first_name like ? or last_name like ? or username like ?)"
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
// select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id in (select user_id from channel_participants where channel_id = :channelId and state = 0) and (first_name like :q1 or last_name like :q2 or username like :q3)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) QueryChannelParticipantsWithCB(ctx context.Context, channelId int64, q1 string, q2 string, q3 string, cb func(i int, v *dataobject.UsersDO)) (rList []dataobject.UsersDO, err error) {
	var (
		query  = "select id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, photo_id, country_code, verified, about, is_bot, deleted from users where id in (select user_id from channel_participants where channel_id = ? and state = 0) and (first_name like ? or last_name like ? or username like ?)"
		values []dataobject.UsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, channelId, q1, q2, q3)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in QueryChannelParticipants(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}
