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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	imported_contactsFieldNames          = builder.RawFieldNames(&ImportedContacts{})
	imported_contactsRows                = strings.Join(imported_contactsFieldNames, ",")
	imported_contactsRowsExpectAutoSet   = strings.Join(stringx.Remove(imported_contactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	imported_contactsRowsWithPlaceHolder = strings.Join(stringx.Remove(imported_contactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTImportedContactsIdPrefix = "cache:t:imported_contacts:id:"

	cacheImportedContactsIdPrefix = "cache#ImportedContacts#id"

	cacheImportedContactsUserIdPrefix = "cache#UserId"

	cacheImportedContactsUserIdImportedUserIdPrefix = "cache#UserId#ImportedUserId"
)

type (
	imported_contactsModel interface {
		Insert2(ctx context.Context, data *ImportedContacts) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ImportedContacts, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]ImportedContacts, error)
		Update2(ctx context.Context, data *ImportedContacts) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserId(ctx context.Context, userId int64) (*ImportedContacts, error)
		FindListByUserIdList(ctx context.Context, userId ...int64) ([]ImportedContacts, error)

		FindOneByUserIdImportedUserId(ctx context.Context, userId int64, importedUserId int64) (*ImportedContacts, error)
	}

	defaultImportedContactsModel struct {
		db *sqlx.DB
	}

	ImportedContacts struct {
		Id             int64 `db:"id" json:"id"`
		UserId         int64 `db:"user_id" json:"user_id"`
		ImportedUserId int64 `db:"imported_user_id" json:"imported_user_id"`
		Deleted        bool  `db:"deleted" json:"deleted"`
	}
)

func newImportedContactsModel(db *sqlx.DB) *defaultImportedContactsModel {
	return &defaultImportedContactsModel{
		db: db,
	}
}

func (m *defaultImportedContactsModel) Insert2(ctx context.Context, data *ImportedContacts) (sql.Result, error) {
	query := fmt.Sprintf("insert into `imported_contacts` (%s) values (?, ?, ?)", imported_contactsRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.ImportedUserId, data.Deleted)
}

func (m *defaultImportedContactsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `imported_contacts` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultImportedContactsModel) FindOne(ctx context.Context, id int64) (*ImportedContacts, error) {
	query := fmt.Sprintf("select %s from imported_contacts where id = ? limit 1", imported_contactsRows)
	var resp ImportedContacts
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultImportedContactsModel) FindListByIdList(ctx context.Context, id ...int64) ([]ImportedContacts, error) {
	if len(id) == 0 {
		return []ImportedContacts{}, nil
	}

	query := fmt.Sprintf("select %s from imported_contacts where id in (%s)", imported_contactsRows, sqlx.InInt64List(id))

	var resp []ImportedContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultImportedContactsModel) Update2(ctx context.Context, data *ImportedContacts) error {
	query := fmt.Sprintf("update `imported_contacts` set %s where `id` = ?", imported_contactsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.ImportedUserId, data.Deleted, data.Id)
	return err
}

func (m *defaultImportedContactsModel) FindOneByUserId(ctx context.Context, userId int64) (*ImportedContacts, error) {
	query := fmt.Sprintf("select %s from imported_contacts where user_id = ? limit 1", imported_contactsRows)
	var resp ImportedContacts
	err := m.db.QueryRowPartial(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultImportedContactsModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]ImportedContacts, error) {
	if len(userId) == 0 {
		return []ImportedContacts{}, nil
	}

	query := fmt.Sprintf("select %s from imported_contacts where user_id in (%s)", imported_contactsRows, sqlx.InInt64List(userId))

	var resp []ImportedContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultImportedContactsModel) FindOneByUserIdImportedUserId(ctx context.Context, userId int64, importedUserId int64) (*ImportedContacts, error) {
	query := fmt.Sprintf("select %s from imported_contacts where user_id = ? AND imported_user_id = ? limit 1", imported_contactsRows)
	var resp ImportedContacts
	err := m.db.QueryRowPartial(ctx, &resp, query, userId, importedUserId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultImportedContactsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheImportedContactsIdPrefix, primary)
}

func (m *defaultImportedContactsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from imported_contacts where id = ? limit 1", imported_contactsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
