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
	userContactsFieldNames          = builder.RawFieldNames(&UserContacts{})
	userContactsRows                = strings.Join(userContactsFieldNames, ",")
	userContactsRowsExpectAutoSet   = strings.Join(stringx.Remove(userContactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userContactsRowsWithPlaceHolder = strings.Join(stringx.Remove(userContactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userContactsModel interface {
		Insert2(ctx context.Context, data *UserContacts) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserContacts, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserContacts, error)
		Update2(ctx context.Context, data *UserContacts) error
		Delete2(ctx context.Context, id int64) error

		FindOneByOwnerUserIdContactUserId(ctx context.Context, ownerUserId int64, contactUserId int64) (*UserContacts, error)
	}

	defaultUserContactsModel struct {
		db *sqlx.DB
	}

	UserContacts struct {
		Id               int64  `db:"id" json:"id"`
		OwnerUserId      int64  `db:"owner_user_id" json:"owner_user_id"`
		ContactUserId    int64  `db:"contact_user_id" json:"contact_user_id"`
		ContactPhone     string `db:"contact_phone" json:"contact_phone"`
		ContactFirstName string `db:"contact_first_name" json:"contact_first_name"`
		ContactLastName  string `db:"contact_last_name" json:"contact_last_name"`
		Mutual           bool   `db:"mutual" json:"mutual"`
		CloseFriend      bool   `db:"close_friend" json:"close_friend"`
		StoriesHidden    bool   `db:"stories_hidden" json:"stories_hidden"`
		IsDeleted        bool   `db:"is_deleted" json:"is_deleted"`
		Date2            int64  `db:"date2" json:"date2"`
	}
)

func newUserContactsModel(db *sqlx.DB) *defaultUserContactsModel {
	return &defaultUserContactsModel{
		db: db,
	}
}

func (m *defaultUserContactsModel) Insert2(ctx context.Context, data *UserContacts) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_contacts` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", userContactsRowsExpectAutoSet)

	return m.db.Exec(ctx, query, data.OwnerUserId, data.ContactUserId, data.ContactPhone, data.ContactFirstName, data.ContactLastName, data.Mutual, data.CloseFriend, data.StoriesHidden, data.IsDeleted, data.Date2)

}

func (m *defaultUserContactsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_contacts` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserContactsModel) FindOne(ctx context.Context, id int64) (*UserContacts, error) {
	query := fmt.Sprintf("select %s from user_contacts where id = ? limit 1", userContactsRows)
	var resp UserContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserContactsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserContacts, error) {
	if len(id) == 0 {
		return []UserContacts{}, nil
	}

	query := fmt.Sprintf("select %s from user_contacts where id in (%s)", userContactsRows, sqlx.InInt64List(id))

	var resp []UserContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserContactsModel) Update2(ctx context.Context, data *UserContacts) error {
	query := fmt.Sprintf("update `user_contacts` set %s where `id` = ?", userContactsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.OwnerUserId, data.ContactUserId, data.ContactPhone, data.ContactFirstName, data.ContactLastName, data.Mutual, data.CloseFriend, data.StoriesHidden, data.IsDeleted, data.Date2, data.Id)
	return err
}

func (m *defaultUserContactsModel) FindOneByOwnerUserIdContactUserId(ctx context.Context, ownerUserId int64, contactUserId int64) (*UserContacts, error) {
	query := fmt.Sprintf("select %s from user_contacts where owner_user_id = ? AND contact_user_id = ? limit 1", userContactsRows)
	var resp UserContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, ownerUserId, contactUserId)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}
