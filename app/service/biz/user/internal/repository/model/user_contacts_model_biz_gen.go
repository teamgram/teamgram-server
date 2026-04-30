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

type bizUserContactsModel interface {
	InsertOrUpdate(ctx context.Context, data *UserContacts) (lastInsertId, rowsAffected int64, err error)
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
	UpdateContactName(ctx context.Context, contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
	UpdateMutual(ctx context.Context, mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
	DeleteContacts(ctx context.Context, ownerUserId int64, idList []int64) (rowsAffected int64, err error)
	UpdatePhoneByContactId(ctx context.Context, contactPhone string, contactUserId int64) (rowsAffected int64, err error)
	SelectUserReverseContactIdList(ctx context.Context, contactUserId int64) ([]int64, error)
	SelectUserReverseContactIdListWithCB(ctx context.Context, contactUserId int64, cb func(sz, i int, v int64)) ([]int64, error)
	SelectReverseListByIdList(ctx context.Context, contactUserId int64, idList []int64) ([]UserContacts, error)
	SelectReverseListByIdListWithCB(ctx context.Context, contactUserId int64, idList []int64, cb func(sz, i int, v *UserContacts)) ([]UserContacts, error)
	UpdateCloseFriend(ctx context.Context, closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error)
	UpdateStoriesHidden(ctx context.Context, storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
}

type UserContactsTxModel interface {
	InsertOrUpdate(data *UserContacts) (lastInsertId, rowsAffected int64, err error)
	SelectContact(ownerUserId int64, contactUserId int64) (*UserContacts, error)
	SelectByContactId(ownerUserId int64, contactUserId int64) (*UserContacts, error)
	SelectListByPhoneList(ownerUserId int64, phoneList []string) ([]UserContacts, error)
	SelectAllUserContacts(ownerUserId int64) ([]UserContacts, error)
	SelectUserContacts(ownerUserId int64) ([]UserContacts, error)
	SelectUserContactIdList(ownerUserId int64) ([]int64, error)
	SelectListByIdList(ownerUserId int64, idList []int64) ([]UserContacts, error)
	SelectListByOwnerListAndContactList(idList1 []int64, idList2 []int64) ([]UserContacts, error)
	UpdateContactNameById(contactFirstName string, contactLastName string, id int64) (rowsAffected int64, err error)
	UpdateContactName(contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
	UpdateMutual(mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
	DeleteContacts(ownerUserId int64, idList []int64) (rowsAffected int64, err error)
	UpdatePhoneByContactId(contactPhone string, contactUserId int64) (rowsAffected int64, err error)
	SelectUserReverseContactIdList(contactUserId int64) ([]int64, error)
	SelectReverseListByIdList(contactUserId int64, idList []int64) ([]UserContacts, error)
	UpdateCloseFriend(closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error)
	UpdateStoriesHidden(storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error)
}

type defaultUserContactsTxModel struct {
	tx *sqlx.Tx
}

func NewUserContactsTxModel(tx *sqlx.Tx) UserContactsTxModel {
	return &defaultUserContactsTxModel{tx: tx}
}

// InsertOrUpdate
// insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0
func (m *defaultUserContactsModel) InsertOrUpdate(ctx context.Context, data *UserContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_contacts.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_contacts.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0
func (m *defaultUserContactsTxModel) InsertOrUpdate(data *UserContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_contacts.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_contacts.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.InsertOrUpdate rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_contacts",
				Key:      fmt.Sprintf("owner_user_id=%v,contact_user_id=%v", ownerUserId, contactUserId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_contacts.SelectContact: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectContact
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
func (m *defaultUserContactsTxModel) SelectContact(ownerUserId int64, contactUserId int64) (rValue *UserContacts, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = ? and contact_user_id = ?"
		do    = &UserContacts{}
	)
	err = m.tx.QueryRowPartial(do, query, ownerUserId, contactUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_contacts",
				Key:      fmt.Sprintf("owner_user_id=%v,contact_user_id=%v", ownerUserId, contactUserId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_contacts.SelectContact: %w", err)
		return
	}
	rValue = do

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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_contacts",
				Key:      fmt.Sprintf("owner_user_id=%v,contact_user_id=%v", ownerUserId, contactUserId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_contacts.SelectByContactId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByContactId
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id and is_deleted = 0
func (m *defaultUserContactsTxModel) SelectByContactId(ownerUserId int64, contactUserId int64) (rValue *UserContacts, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_user_id = ? and is_deleted = 0"
		do    = &UserContacts{}
	)
	err = m.tx.QueryRowPartial(do, query, ownerUserId, contactUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_contacts",
				Key:      fmt.Sprintf("owner_user_id=%v,contact_user_id=%v", ownerUserId, contactUserId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_contacts.SelectByContactId: %w", err)
		return
	}
	rValue = do

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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByPhoneList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByPhoneList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_phone in (:phoneList) and is_deleted = 0
func (m *defaultUserContactsTxModel) SelectListByPhoneList(ownerUserId int64, phoneList []string) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_phone in (%s) and is_deleted = 0", sqlx.InStringList(phoneList))
		values []UserContacts
	)
	if len(phoneList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, ownerUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByPhoneList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByPhoneListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectAllUserContacts: %w", err)
		return
	}

	rList = values

	return
}

// SelectAllUserContacts
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0
func (m *defaultUserContactsTxModel) SelectAllUserContacts(ownerUserId int64) (rList []UserContacts, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0"
		values []UserContacts
	)
	err = m.tx.QueryRowsPartial(&values, query, ownerUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectAllUserContacts: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectAllUserContactsWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserContacts: %w", err)
		return
	}

	rList = values

	return
}

// SelectUserContacts
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsTxModel) SelectUserContacts(ownerUserId int64) (rList []UserContacts, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
		values []UserContacts
	)
	err = m.tx.QueryRowsPartial(&values, query, ownerUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserContacts: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserContactsWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserContactIdList: %w", err)
	}

	return
}

// SelectUserContactIdList
// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsTxModel) SelectUserContactIdList(ownerUserId int64) (rList []int64, err error) {
	var query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
	err = m.tx.QueryRowsPartial(&rList, query, ownerUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserContactIdList: %w", err)
	}

	return
}

// SelectUserContactIdListWithCB
// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsModel) SelectUserContactIdListWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, ownerUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserContactIdListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByIdList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id in (:id_list) and is_deleted = 0 order by contact_user_id asc
func (m *defaultUserContactsTxModel) SelectListByIdList(ownerUserId int64, idList []int64) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_user_id in (%s) and is_deleted = 0 order by contact_user_id asc", sqlx.InInt64List(idList))
		values []UserContacts
	)
	if len(idList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, ownerUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByIdListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByOwnerListAndContactList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByOwnerListAndContactList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (:idList1) and contact_user_id in (:idList2) and is_deleted = 0
func (m *defaultUserContactsTxModel) SelectListByOwnerListAndContactList(idList1 []int64, idList2 []int64) (rList []UserContacts, err error) {
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

	err = m.tx.QueryRowsPartial(&values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByOwnerListAndContactList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectListByOwnerListAndContactListWithCB: %w", err)
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
		err = fmt.Errorf("user_contacts.UpdateContactNameById exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateContactNameById rows affected: %w", err)
		return
	}

	return
}

// UpdateContactNameById
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where id = :id
func (m *defaultUserContactsTxModel) UpdateContactNameById(contactFirstName string, contactLastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, contactFirstName, contactLastName, id)

	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateContactNameById exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateContactNameById rows affected: %w", err)
		return
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
		err = fmt.Errorf("user_contacts.UpdateContactName exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateContactName rows affected: %w", err)
		return
	}

	return
}

// UpdateContactName
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (m *defaultUserContactsTxModel) UpdateContactName(contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, contactFirstName, contactLastName, ownerUserId, contactUserId)

	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateContactName exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateContactName rows affected: %w", err)
		return
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
		err = fmt.Errorf("user_contacts.UpdateMutual exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateMutual rows affected: %w", err)
		return
	}

	return
}

// UpdateMutual
// update user_contacts set mutual = :mutual where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (m *defaultUserContactsTxModel) UpdateMutual(mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, mutual, ownerUserId, contactUserId)

	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateMutual exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateMutual rows affected: %w", err)
		return
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
		err = fmt.Errorf("user_contacts.DeleteContacts exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.DeleteContacts rows affected: %w", err)
		return
	}

	return
}

// DeleteContacts
// update user_contacts set is_deleted = 1, mutual = 0, close_friend = 0, stories_hidden = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id in (:id_list))
func (m *defaultUserContactsTxModel) DeleteContacts(ownerUserId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update user_contacts set is_deleted = 1, mutual = 0, close_friend = 0, stories_hidden = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (%s))", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, ownerUserId)

	if err != nil {
		err = fmt.Errorf("user_contacts.DeleteContacts exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.DeleteContacts rows affected: %w", err)
		return
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
		err = fmt.Errorf("user_contacts.UpdatePhoneByContactId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdatePhoneByContactId rows affected: %w", err)
		return
	}

	return
}

// UpdatePhoneByContactId
// update user_contacts set contact_phone = :contact_phone where contact_user_id = :contact_user_id
func (m *defaultUserContactsTxModel) UpdatePhoneByContactId(contactPhone string, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_phone = ? where contact_user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, contactPhone, contactUserId)

	if err != nil {
		err = fmt.Errorf("user_contacts.UpdatePhoneByContactId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdatePhoneByContactId rows affected: %w", err)
		return
	}

	return
}

// SelectUserReverseContactIdList
// select owner_user_id from user_contacts where contact_user_id = :contact_user_id and is_deleted = 0
func (m *defaultUserContactsModel) SelectUserReverseContactIdList(ctx context.Context, contactUserId int64) (rList []int64, err error) {
	var query = "select owner_user_id from user_contacts where contact_user_id = ? and is_deleted = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, contactUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserReverseContactIdList: %w", err)
	}

	return
}

// SelectUserReverseContactIdList
// select owner_user_id from user_contacts where contact_user_id = :contact_user_id and is_deleted = 0
func (m *defaultUserContactsTxModel) SelectUserReverseContactIdList(contactUserId int64) (rList []int64, err error) {
	var query = "select owner_user_id from user_contacts where contact_user_id = ? and is_deleted = 0"
	err = m.tx.QueryRowsPartial(&rList, query, contactUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserReverseContactIdList: %w", err)
	}

	return
}

// SelectUserReverseContactIdListWithCB
// select owner_user_id from user_contacts where contact_user_id = :contact_user_id and is_deleted = 0
func (m *defaultUserContactsModel) SelectUserReverseContactIdListWithCB(ctx context.Context, contactUserId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select owner_user_id from user_contacts where contact_user_id = ? and is_deleted = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, contactUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectUserReverseContactIdListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectReverseListByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectReverseListByIdList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = :contact_user_id and owner_user_id in (:id_list) and is_deleted = 0
func (m *defaultUserContactsTxModel) SelectReverseListByIdList(contactUserId int64, idList []int64) (rList []UserContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = ? and owner_user_id in (%s) and is_deleted = 0", sqlx.InInt64List(idList))
		values []UserContacts
	)
	if len(idList) == 0 {
		rList = []UserContacts{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, contactUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectReverseListByIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("user_contacts.SelectReverseListByIdListWithCB: %w", err)
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
		err = fmt.Errorf("user_contacts.UpdateCloseFriend exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateCloseFriend rows affected: %w", err)
		return
	}

	return
}

// UpdateCloseFriend
// update user_contacts set close_friend = :close_friend where owner_user_id = :owner_user_id and contact_user_id in (:idList)
func (m *defaultUserContactsTxModel) UpdateCloseFriend(closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update user_contacts set close_friend = ? where owner_user_id = ? and contact_user_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, closeFriend, ownerUserId)

	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateCloseFriend exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateCloseFriend rows affected: %w", err)
		return
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
		err = fmt.Errorf("user_contacts.UpdateStoriesHidden exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateStoriesHidden rows affected: %w", err)
		return
	}

	return
}

// UpdateStoriesHidden
// update user_contacts set stories_hidden = :stories_hidden where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
func (m *defaultUserContactsTxModel) UpdateStoriesHidden(storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set stories_hidden = ? where owner_user_id = ? and contact_user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, storiesHidden, ownerUserId, contactUserId)

	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateStoriesHidden exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_contacts.UpdateStoriesHidden rows affected: %w", err)
		return
	}

	return
}
