/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
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
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *logx.Logger

type (
	bizUserContactsModel interface {
		InsertOrUpdate(ctx context.Context, data *UserContacts) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UserContacts) (lastInsertId, rowsAffected int64, err error)

		SelectContact(ctx context.Context, ownerUserId int64, contactUserId int64) (*UserContacts, error)

		SelectByContactId(ctx context.Context, ownerUserId int64, contactUserId int64) (*UserContacts, error)

		SelectListByPhoneList(ctx context.Context, ownerUserId int64, phoneList []string) ([]UserContacts, error)
		SelectListByPhoneListWithCB(ctx context.Context, ownerUserId int64, phoneList []string, cb func(sz, i int, v *UserContacts)) ([]UserContacts, error)

		SelectAllUserContacts(ctx context.Context, ownerUserId int64) ([]UserContacts, error)
		SelectAllUserContactsWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v *UserContacts)) ([]UserContacts, error)

		SelectUserContacts(ctx context.Context, ownerUserId int64) ([]UserContacts, error)
		SelectUserContactsWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v *UserContacts)) ([]UserContacts, error)

		SelectUserContactIdList(ctx context.Context, ownerUserId int64) ([]int64, error)
		SelectUserContactIdListWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v int64)) ([]int64, error)

		SelectListByIdList(ctx context.Context, ownerUserId int64, idList []int64) ([]UserContacts, error)
		SelectListByIdListWithCB(ctx context.Context, ownerUserId int64, idList []int64, cb func(sz, i int, v *UserContacts)) ([]UserContacts, error)

		SelectListByOwnerListAndContactList(ctx context.Context, idList1 []int64, idList2 []int64) ([]UserContacts, error)
		SelectListByOwnerListAndContactListWithCB(ctx context.Context, idList1 []int64, idList2 []int64, cb func(sz, i int, v *UserContacts)) ([]UserContacts, error)

		UpdateContactNameById(ctx context.Context, contactFirstName string, contactLastName string, id int64) (rowsAffected int64, err error)
		UpdateContactNameByIdTx(tx *sqlx.Tx, contactFirstName string, contactLastName string, id int64) (rowsAffected int64, err error)

		UpdateContactName(ctx context.Context, contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
		UpdateContactNameTx(tx *sqlx.Tx, contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)

		UpdateMutual(ctx context.Context, mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
		UpdateMutualTx(tx *sqlx.Tx, mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)

		DeleteContacts(ctx context.Context, ownerUserId int64, idList []int64) (rowsAffected int64, err error)
		DeleteContactsTx(tx *sqlx.Tx, ownerUserId int64, idList []int64) (rowsAffected int64, err error)

		UpdatePhoneByContactId(ctx context.Context, contactPhone string, contactUserId int64) (rowsAffected int64, err error)
		UpdatePhoneByContactIdTx(tx *sqlx.Tx, contactPhone string, contactUserId int64) (rowsAffected int64, err error)

		SelectUserReverseContactIdList(ctx context.Context, contactUserId int64) ([]int64, error)
		SelectUserReverseContactIdListWithCB(ctx context.Context, contactUserId int64, cb func(sz, i int, v int64)) ([]int64, error)

		SelectReverseListByIdList(ctx context.Context, contactUserId int64, idList []int64) ([]UserContacts, error)
		SelectReverseListByIdListWithCB(ctx context.Context, contactUserId int64, idList []int64, cb func(sz, i int, v *UserContacts)) ([]UserContacts, error)

		UpdateCloseFriend(ctx context.Context, closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error)
		UpdateCloseFriendTx(tx *sqlx.Tx, closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error)

		UpdateStoriesHidden(ctx context.Context, storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
		UpdateStoriesHiddenTx(tx *sqlx.Tx, storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0
func (m *defaultUserContactsModel) InsertOrUpdate(ctx context.Context, data *UserContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// InsertOrUpdateTx
// insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0
func (m *defaultUserContactsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UserContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// SelectContact
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
func (m *defaultUserContactsModel) SelectContact(ctx context.Context, ownerUserId int64, contactUserId int64) (rValue *UserContacts, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = ? and contact_user_id = ?"
		do    = &UserContacts{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, ownerUserId, contactUserId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectContact(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByContactId
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id and is_deleted = 0
func (m *defaultUserContactsModel) SelectByContactId(ctx context.Context, ownerUserId int64, contactUserId int64) (rValue *UserContacts, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_user_id = ? and is_deleted = 0"
		do    = &UserContacts{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, ownerUserId, contactUserId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByContactId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectListByPhoneList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_phone in (:phoneList) and is_deleted = 0
func (m *defaultUserContactsModel) SelectListByPhoneList(ctx context.Context, ownerUserId int64, phoneList []string) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_phone in (%s) and is_deleted = 0", sqlx.InStringList(phoneList))
		values []UserContacts
	)
	if len(phoneList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhoneList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByPhoneListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_phone in (:phoneList) and is_deleted = 0
func (m *defaultUserContactsModel) SelectListByPhoneListWithCB(ctx context.Context, ownerUserId int64, phoneList []string, cb func(sz, i int, v *UserContacts)) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_phone in (%s) and is_deleted = 0", sqlx.InStringList(phoneList))
		values []UserContacts
	)
	if len(phoneList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhoneList(_), error: %v", err)
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

// SelectAllUserContacts
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0
func (m *defaultUserContactsModel) SelectAllUserContacts(ctx context.Context, ownerUserId int64) (rList []UserContacts, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0"
		values []UserContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllUserContacts(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllUserContactsWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0
func (m *defaultUserContactsModel) SelectAllUserContactsWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v *UserContacts)) (rList []UserContacts, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0"
		values []UserContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllUserContacts(_), error: %v", err)
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

// SelectUserContacts
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsModel) SelectUserContacts(ctx context.Context, ownerUserId int64) (rList []UserContacts, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
		values []UserContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUserContacts(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUserContactsWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsModel) SelectUserContactsWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v *UserContacts)) (rList []UserContacts, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
		values []UserContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUserContacts(_), error: %v", err)
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

// SelectUserContactIdList
// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsModel) SelectUserContactIdList(ctx context.Context, ownerUserId int64) (rList []int64, err error) {
	var query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserContactIdList(_), error: %v", err)
	}

	return
}

// SelectUserContactIdListWithCB
// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsModel) SelectUserContactIdListWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserContactIdList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectListByIdList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id in (:id_list) and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsModel) SelectListByIdList(ctx context.Context, ownerUserId int64, idList []int64) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_user_id in (%s) and is_deleted = 0 order by contact_user_id asc", sqlx.InInt64List(idList))
		values []UserContacts
	)
	if len(idList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByIdListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id in (:id_list) and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsModel) SelectListByIdListWithCB(ctx context.Context, ownerUserId int64, idList []int64, cb func(sz, i int, v *UserContacts)) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_user_id in (%s) and is_deleted = 0 order by contact_user_id asc", sqlx.InInt64List(idList))
		values []UserContacts
	)
	if len(idList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
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

// SelectListByOwnerListAndContactList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (:idList1) and contact_user_id in (:idList2) and is_deleted = 0
func (m *defaultUserContactsModel) SelectListByOwnerListAndContactList(ctx context.Context, idList1 []int64, idList2 []int64) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (%s) and contact_user_id in (%s) and is_deleted = 0", sqlx.InInt64List(idList1), sqlx.InInt64List(idList2))
		values []UserContacts
	)
	if len(idList1) == 0 {
		rList = []UserContacts{}
		return
	}
	if len(idList2) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByOwnerListAndContactList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByOwnerListAndContactListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (:idList1) and contact_user_id in (:idList2) and is_deleted = 0
func (m *defaultUserContactsModel) SelectListByOwnerListAndContactListWithCB(ctx context.Context, idList1 []int64, idList2 []int64, cb func(sz, i int, v *UserContacts)) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (%s) and contact_user_id in (%s) and is_deleted = 0", sqlx.InInt64List(idList1), sqlx.InInt64List(idList2))
		values []UserContacts
	)
	if len(idList1) == 0 {
		rList = []UserContacts{}
		return
	}
	if len(idList2) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByOwnerListAndContactList(_), error: %v", err)
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

// UpdateContactNameById
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where id = :id
func (m *defaultUserContactsModel) UpdateContactNameById(ctx context.Context, contactFirstName string, contactLastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, contactFirstName, contactLastName, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateContactNameById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateContactNameById(_), error: %v", err)
	}

	return
}

// UpdateContactNameByIdTx
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where id = :id
func (m *defaultUserContactsModel) UpdateContactNameByIdTx(tx *sqlx.Tx, contactFirstName string, contactLastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, contactFirstName, contactLastName, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateContactNameById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateContactNameById(_), error: %v", err)
	}

	return
}

// UpdateContactName
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (m *defaultUserContactsModel) UpdateContactName(ctx context.Context, contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, contactFirstName, contactLastName, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

// UpdateContactNameTx
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (m *defaultUserContactsModel) UpdateContactNameTx(tx *sqlx.Tx, contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, contactFirstName, contactLastName, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

// UpdateMutual
// update user_contacts set mutual = :mutual where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (m *defaultUserContactsModel) UpdateMutual(ctx context.Context, mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, mutual, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMutual(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMutual(_), error: %v", err)
	}

	return
}

// UpdateMutualTx
// update user_contacts set mutual = :mutual where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (m *defaultUserContactsModel) UpdateMutualTx(tx *sqlx.Tx, mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, mutual, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMutual(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMutual(_), error: %v", err)
	}

	return
}

// DeleteContacts
// update user_contacts set is_deleted = 1, mutual = 0, close_friend = 0, stories_hidden = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id in (:id_list))
func (m *defaultUserContactsModel) DeleteContacts(ctx context.Context, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update user_contacts set is_deleted = 1, mutual = 0, close_friend = 0, stories_hidden = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (%s))", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

// DeleteContactsTx
// update user_contacts set is_deleted = 1, mutual = 0, close_friend = 0, stories_hidden = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id in (:id_list))
func (m *defaultUserContactsModel) DeleteContactsTx(tx *sqlx.Tx, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update user_contacts set is_deleted = 1, mutual = 0, close_friend = 0, stories_hidden = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (%s))", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = tx.Exec(query, ownerUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

// UpdatePhoneByContactId
// update user_contacts set contact_phone = :contact_phone where contact_user_id = :contact_user_id
func (m *defaultUserContactsModel) UpdatePhoneByContactId(ctx context.Context, contactPhone string, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_phone = ? where contact_user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, contactPhone, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePhoneByContactId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePhoneByContactId(_), error: %v", err)
	}

	return
}

// UpdatePhoneByContactIdTx
// update user_contacts set contact_phone = :contact_phone where contact_user_id = :contact_user_id
func (m *defaultUserContactsModel) UpdatePhoneByContactIdTx(tx *sqlx.Tx, contactPhone string, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_phone = ? where contact_user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, contactPhone, contactUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePhoneByContactId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePhoneByContactId(_), error: %v", err)
	}

	return
}

// SelectUserReverseContactIdList
// select owner_user_id from user_contacts where contact_user_id = :contact_user_id and is_deleted = 0
func (m *defaultUserContactsModel) SelectUserReverseContactIdList(ctx context.Context, contactUserId int64) (rList []int64, err error) {
	var query = "select owner_user_id from user_contacts where contact_user_id = ? and is_deleted = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserReverseContactIdList(_), error: %v", err)
	}

	return
}

// SelectUserReverseContactIdListWithCB
// select owner_user_id from user_contacts where contact_user_id = :contact_user_id and is_deleted = 0
func (m *defaultUserContactsModel) SelectUserReverseContactIdListWithCB(ctx context.Context, contactUserId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select owner_user_id from user_contacts where contact_user_id = ? and is_deleted = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserReverseContactIdList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectReverseListByIdList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = :contact_user_id and owner_user_id in (:id_list) and is_deleted = 0
func (m *defaultUserContactsModel) SelectReverseListByIdList(ctx context.Context, contactUserId int64, idList []int64) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = ? and owner_user_id in (%s) and is_deleted = 0", sqlx.InInt64List(idList))
		values []UserContacts
	)
	if len(idList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectReverseListByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectReverseListByIdListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = :contact_user_id and owner_user_id in (:id_list) and is_deleted = 0
func (m *defaultUserContactsModel) SelectReverseListByIdListWithCB(ctx context.Context, contactUserId int64, idList []int64, cb func(sz, i int, v *UserContacts)) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = ? and owner_user_id in (%s) and is_deleted = 0", sqlx.InInt64List(idList))
		values []UserContacts
	)
	if len(idList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectReverseListByIdList(_), error: %v", err)
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

// UpdateCloseFriend
// update user_contacts set close_friend = :close_friend where owner_user_id = :owner_user_id and contact_user_id in (:idList)
func (m *defaultUserContactsModel) UpdateCloseFriend(ctx context.Context, closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update user_contacts set close_friend = ? where owner_user_id = ? and contact_user_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, closeFriend, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateCloseFriend(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateCloseFriend(_), error: %v", err)
	}

	return
}

// UpdateCloseFriendTx
// update user_contacts set close_friend = :close_friend where owner_user_id = :owner_user_id and contact_user_id in (:idList)
func (m *defaultUserContactsModel) UpdateCloseFriendTx(tx *sqlx.Tx, closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update user_contacts set close_friend = ? where owner_user_id = ? and contact_user_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = tx.Exec(query, closeFriend, ownerUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateCloseFriend(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateCloseFriend(_), error: %v", err)
	}

	return
}

// UpdateStoriesHidden
// update user_contacts set stories_hidden = :stories_hidden where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
func (m *defaultUserContactsModel) UpdateStoriesHidden(ctx context.Context, storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set stories_hidden = ? where owner_user_id = ? and contact_user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, storiesHidden, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateStoriesHidden(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateStoriesHidden(_), error: %v", err)
	}

	return
}

// UpdateStoriesHiddenTx
// update user_contacts set stories_hidden = :stories_hidden where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
func (m *defaultUserContactsModel) UpdateStoriesHiddenTx(tx *sqlx.Tx, storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set stories_hidden = ? where owner_user_id = ? and contact_user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, storiesHidden, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateStoriesHidden(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateStoriesHidden(_), error: %v", err)
	}

	return
}
