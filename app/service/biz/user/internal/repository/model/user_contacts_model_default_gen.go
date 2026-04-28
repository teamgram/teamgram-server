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
	tableName := "user_contacts"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, userContactsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.OwnerUserId, data.ContactUserId, data.ContactPhone, data.ContactFirstName, data.ContactLastName, data.Mutual, data.CloseFriend, data.StoriesHidden, data.IsDeleted, data.Date2)
	if err != nil {
		return nil, fmt.Errorf("user_contacts.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserContactsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "user_contacts"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("user_contacts.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserContactsModel) FindOne(ctx context.Context, id int64) (*UserContacts, error) {
	tableName := "user_contacts"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", userContactsRows, tableName)
	var resp UserContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_contacts",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_contacts.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserContactsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserContacts, error) {
	if len(id) == 0 {
		return []UserContacts{}, nil
	}
	tableName := "user_contacts"

	query := fmt.Sprintf("select %s from %s where id in (%s)", userContactsRows, tableName, sqlx.InInt64List(id))

	var resp []UserContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserContacts{}, nil
		}
		return nil, fmt.Errorf("user_contacts.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserContactsModel) Update2(ctx context.Context, data *UserContacts) error {
	tableName := "user_contacts"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, userContactsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.OwnerUserId, data.ContactUserId, data.ContactPhone, data.ContactFirstName, data.ContactLastName, data.Mutual, data.CloseFriend, data.StoriesHidden, data.IsDeleted, data.Date2, data.Id)
	if err != nil {
		return fmt.Errorf("user_contacts.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserContactsModel) FindOneByOwnerUserIdContactUserId(ctx context.Context, ownerUserId int64, contactUserId int64) (*UserContacts, error) {
	tableName := "user_contacts"
	query := fmt.Sprintf("select %s from %s where owner_user_id = ? AND contact_user_id = ? limit 1", userContactsRows, tableName)
	var resp UserContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, ownerUserId, contactUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_contacts",
				Key:      fmt.Sprintf("owner_user_id=%v,contact_user_id=%v", ownerUserId, contactUserId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_contacts.FindOneByOwnerUserIdContactUserId: %w", err)
	}

	return &resp, nil
}
