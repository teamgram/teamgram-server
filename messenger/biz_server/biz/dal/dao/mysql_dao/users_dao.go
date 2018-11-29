// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package mysql_dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

type UsersDAO struct {
	db *sqlx.DB
}

func NewUsersDAO(db *sqlx.DB) *UsersDAO {
	return &UsersDAO{db}
}

// insert into users(first_name, last_name, access_hash, username, phone, country_code, verified, about, is_bot) values (:first_name, :last_name, :access_hash, :username, :phone, :country_code, :verified, :about, :is_bot)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) Insert(do *dataobject.UsersDO) int64 {
	var query = "insert into users(first_name, last_name, access_hash, username, phone, country_code, verified, about, is_bot) values (:first_name, :last_name, :access_hash, :username, :phone, :country_code, :verified, :about, :is_bot)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in Insert(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in Insert(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}

// select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where phone = :phone limit 1
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByPhoneNumber(phone string) *dataobject.UsersDO {
	var query = "select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where phone = ? limit 1"
	rows, err := dao.db.Queryx(query, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByPhoneNumber(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByPhoneNumber(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByPhoneNumber(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where id = :id limit 1
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectById(id int32) *dataobject.UsersDO {
	var query = "select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where id = ? limit 1"
	rows, err := dao.db.Queryx(query, id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectById(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectById(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectById(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectUsersByIdList(id_list []int32) []dataobject.UsersDO {
	var q = "select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where id in (?)"
	query, a, err := sqlx.In(q, id_list)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectUsersByIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectUsersByIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectUsersByIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where phone in (:phoneList)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectUsersByPhoneList(phoneList []string) []dataobject.UsersDO {
	var q = "select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where phone in (?)"
	query, a, err := sqlx.In(q, phoneList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectUsersByPhoneList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectUsersByPhoneList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectUsersByPhoneList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where username = :username or first_name = :first_name or last_name = :last_name or phone = :phone limit 20
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByQueryString(username string, first_name string, last_name string, phone string) []dataobject.UsersDO {
	var query = "select id, access_hash, first_name, last_name, username, phone, photos, country_code, verified, about, is_bot from users where username = ? or first_name = ? or last_name = ? or phone = ? limit 20"
	rows, err := dao.db.Queryx(query, username, first_name, last_name, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByQueryString(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByQueryString(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByQueryString(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select id from users where username like :q2 and id not in (:id_list) limit :limit
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SearchByQueryNotIdList(q2 string, id_list []int32, limit int32) []dataobject.UsersDO {
	var q = "select id from users where username like ? and id not in (?) limit ?"
	query, a, err := sqlx.In(q, q2, id_list, limit)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SearchByQueryNotIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SearchByQueryNotIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SearchByQueryNotIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// update users set deleted = 1, delete_reason = :delete_reason where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) Delete(delete_reason string, id int32) int64 {
	var query = "update users set deleted = 1, delete_reason = ? where id = ?"
	r, err := dao.db.Exec(query, delete_reason, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in Delete(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in Delete(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update users set username = :username where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateUsername(username string, id int32) int64 {
	var query = "update users set username = ? where id = ?"
	r, err := dao.db.Exec(query, username, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateUsername(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateUsername(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update users set first_name = :first_name, last_name = :last_name where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateFirstAndLastName(first_name string, last_name string, id int32) int64 {
	var query = "update users set first_name = ?, last_name = ? where id = ?"
	r, err := dao.db.Exec(query, first_name, last_name, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateFirstAndLastName(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateFirstAndLastName(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update users set about = :about where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateAbout(about string, id int32) int64 {
	var query = "update users set about = ? where id = ?"
	r, err := dao.db.Exec(query, about, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateAbout(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateAbout(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update users set first_name = :first_name, last_name = :last_name, about = :about where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateProfile(first_name string, last_name string, about string, id int32) int64 {
	var query = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
	r, err := dao.db.Exec(query, first_name, last_name, about, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateProfile(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateProfile(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select id from users where username = :username limit 1
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByUsername(username string) *dataobject.UsersDO {
	var query = "select id from users where username = ? limit 1"
	rows, err := dao.db.Queryx(query, username)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByUsername(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByUsername(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByUsername(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select account_days_ttl from users where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectAccountDaysTTL(id int32) *dataobject.UsersDO {
	var query = "select account_days_ttl from users where id = ?"
	rows, err := dao.db.Queryx(query, id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectAccountDaysTTL(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectAccountDaysTTL(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectAccountDaysTTL(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update users set account_days_ttl = :account_days_ttl where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateAccountDaysTTL(account_days_ttl int32, id int32) int64 {
	var query = "update users set account_days_ttl = ? where id = ?"
	r, err := dao.db.Exec(query, account_days_ttl, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateAccountDaysTTL(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateAccountDaysTTL(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select photos from users where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectProfilePhotos(id int32) *dataobject.UsersDO {
	var query = "select photos from users where id = ?"
	rows, err := dao.db.Queryx(query, id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectProfilePhotos(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectProfilePhotos(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectProfilePhotos(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// select country_code from users where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectCountryCode(id int32) *dataobject.UsersDO {
	var query = "select country_code from users where id = ?"
	rows, err := dao.db.Queryx(query, id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectCountryCode(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectCountryCode(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectCountryCode(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

// update users set photos = :photos where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateProfilePhotos(photos string, id int32) int64 {
	var query = "update users set photos = ? where id = ?"
	r, err := dao.db.Exec(query, photos, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateProfilePhotos(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateProfilePhotos(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
