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
	importedContactsFieldNames          = builder.RawFieldNames(&ImportedContacts{})
	importedContactsRows                = strings.Join(importedContactsFieldNames, ",")
	importedContactsRowsExpectAutoSet   = strings.Join(stringx.Remove(importedContactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	importedContactsRowsWithPlaceHolder = strings.Join(stringx.Remove(importedContactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	importedContactsModel interface {
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
	query := fmt.Sprintf("insert into `imported_contacts` (%s) values (?, ?, ?)", importedContactsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.ImportedUserId, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("imported_contacts.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultImportedContactsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `imported_contacts` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("imported_contacts.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultImportedContactsModel) FindOne(ctx context.Context, id int64) (*ImportedContacts, error) {
	query := fmt.Sprintf("select %s from imported_contacts where id = ? limit 1", importedContactsRows)
	var resp ImportedContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "imported_contacts",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("imported_contacts.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultImportedContactsModel) FindListByIdList(ctx context.Context, id ...int64) ([]ImportedContacts, error) {
	if len(id) == 0 {
		return []ImportedContacts{}, nil
	}

	query := fmt.Sprintf("select %s from imported_contacts where id in (%s)", importedContactsRows, sqlx.InInt64List(id))

	var resp []ImportedContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []ImportedContacts{}, nil
		}
		return nil, fmt.Errorf("imported_contacts.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultImportedContactsModel) Update2(ctx context.Context, data *ImportedContacts) error {
	query := fmt.Sprintf("update `imported_contacts` set %s where `id` = ?", importedContactsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.ImportedUserId, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("imported_contacts.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultImportedContactsModel) FindOneByUserId(ctx context.Context, userId int64) (*ImportedContacts, error) {
	query := fmt.Sprintf("select %s from imported_contacts where user_id = ? limit 1", importedContactsRows)
	var resp ImportedContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "imported_contacts",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("imported_contacts.FindOneByUserId: %w", err)
	}

	return &resp, nil
}

func (m *defaultImportedContactsModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]ImportedContacts, error) {
	if len(userId) == 0 {
		return []ImportedContacts{}, nil
	}

	query := fmt.Sprintf("select %s from imported_contacts where user_id in (%s)", importedContactsRows, sqlx.InInt64List(userId))

	var resp []ImportedContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []ImportedContacts{}, nil
		}
		return nil, fmt.Errorf("imported_contacts.FindListByUserIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultImportedContactsModel) FindOneByUserIdImportedUserId(ctx context.Context, userId int64, importedUserId int64) (*ImportedContacts, error) {
	query := fmt.Sprintf("select %s from imported_contacts where user_id = ? AND imported_user_id = ? limit 1", importedContactsRows)
	var resp ImportedContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, importedUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "imported_contacts",
				Key:      fmt.Sprintf("user_id=%v,imported_user_id=%v", userId, importedUserId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("imported_contacts.FindOneByUserIdImportedUserId: %w", err)
	}

	return &resp, nil
}
