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
	unregisteredContactsFieldNames          = builder.RawFieldNames(&UnregisteredContacts{})
	unregisteredContactsRows                = strings.Join(unregisteredContactsFieldNames, ",")
	unregisteredContactsRowsExpectAutoSet   = strings.Join(stringx.Remove(unregisteredContactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	unregisteredContactsRowsWithPlaceHolder = strings.Join(stringx.Remove(unregisteredContactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	unregisteredContactsModel interface {
		Insert2(ctx context.Context, data *UnregisteredContacts) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UnregisteredContacts, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UnregisteredContacts, error)
		Update2(ctx context.Context, data *UnregisteredContacts) error
		Delete2(ctx context.Context, id int64) error

		FindOneByPhoneImporterUserId(ctx context.Context, phone string, importerUserId int64) (*UnregisteredContacts, error)
	}

	defaultUnregisteredContactsModel struct {
		db *sqlx.DB
	}

	UnregisteredContacts struct {
		Id              int64  `db:"id" json:"id"`
		Phone           string `db:"phone" json:"phone"`
		ImporterUserId  int64  `db:"importer_user_id" json:"importer_user_id"`
		ImportFirstName string `db:"import_first_name" json:"import_first_name"`
		ImportLastName  string `db:"import_last_name" json:"import_last_name"`
		Imported        bool   `db:"imported" json:"imported"`
	}
)

func newUnregisteredContactsModel(db *sqlx.DB) *defaultUnregisteredContactsModel {
	return &defaultUnregisteredContactsModel{
		db: db,
	}
}

func (m *defaultUnregisteredContactsModel) Insert2(ctx context.Context, data *UnregisteredContacts) (sql.Result, error) {
	tableName := "unregistered_contacts"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?)", tableName, unregisteredContactsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.Phone, data.ImporterUserId, data.ImportFirstName, data.ImportLastName, data.Imported)
	if err != nil {
		return nil, fmt.Errorf("unregistered_contacts.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUnregisteredContactsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "unregistered_contacts"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("unregistered_contacts.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUnregisteredContactsModel) FindOne(ctx context.Context, id int64) (*UnregisteredContacts, error) {
	tableName := "unregistered_contacts"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", unregisteredContactsRows, tableName)
	var resp UnregisteredContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "unregistered_contacts",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("unregistered_contacts.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUnregisteredContactsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UnregisteredContacts, error) {
	if len(id) == 0 {
		return []UnregisteredContacts{}, nil
	}
	tableName := "unregistered_contacts"

	query := fmt.Sprintf("select %s from %s where id in (%s)", unregisteredContactsRows, tableName, sqlx.InInt64List(id))

	var resp []UnregisteredContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UnregisteredContacts{}, nil
		}
		return nil, fmt.Errorf("unregistered_contacts.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUnregisteredContactsModel) Update2(ctx context.Context, data *UnregisteredContacts) error {
	tableName := "unregistered_contacts"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, unregisteredContactsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.Phone, data.ImporterUserId, data.ImportFirstName, data.ImportLastName, data.Imported, data.Id)
	if err != nil {
		return fmt.Errorf("unregistered_contacts.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUnregisteredContactsModel) FindOneByPhoneImporterUserId(ctx context.Context, phone string, importerUserId int64) (*UnregisteredContacts, error) {
	tableName := "unregistered_contacts"
	query := fmt.Sprintf("select %s from %s where phone = ? AND importer_user_id = ? limit 1", unregisteredContactsRows, tableName)
	var resp UnregisteredContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, phone, importerUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "unregistered_contacts",
				Key:      fmt.Sprintf("phone=%v,importer_user_id=%v", phone, importerUserId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("unregistered_contacts.FindOneByPhoneImporterUserId: %w", err)
	}

	return &resp, nil
}
