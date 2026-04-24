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
	popularContactsFieldNames          = builder.RawFieldNames(&PopularContacts{})
	popularContactsRows                = strings.Join(popularContactsFieldNames, ",")
	popularContactsRowsExpectAutoSet   = strings.Join(stringx.Remove(popularContactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	popularContactsRowsWithPlaceHolder = strings.Join(stringx.Remove(popularContactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	popularContactsModel interface {
		Insert2(ctx context.Context, data *PopularContacts) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PopularContacts, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]PopularContacts, error)
		Update2(ctx context.Context, data *PopularContacts) error
		Delete2(ctx context.Context, id int64) error

		FindOneByPhone(ctx context.Context, phone string) (*PopularContacts, error)
		FindListByPhoneList(ctx context.Context, phone ...string) ([]PopularContacts, error)
	}

	defaultPopularContactsModel struct {
		db *sqlx.DB
	}

	PopularContacts struct {
		Id        int64  `db:"id" json:"id"`
		Phone     string `db:"phone" json:"phone"`
		Importers int32  `db:"importers" json:"importers"`
		Deleted   bool   `db:"deleted" json:"deleted"`
		UpdateAt  string `db:"update_at" json:"update_at"`
	}
)

func newPopularContactsModel(db *sqlx.DB) *defaultPopularContactsModel {
	return &defaultPopularContactsModel{
		db: db,
	}
}

func (m *defaultPopularContactsModel) Insert2(ctx context.Context, data *PopularContacts) (sql.Result, error) {
	query := fmt.Sprintf("insert into `popular_contacts` (%s) values (?, ?, ?)", popularContactsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.Phone, data.Importers, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("popular_contacts.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultPopularContactsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `popular_contacts` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("popular_contacts.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultPopularContactsModel) FindOne(ctx context.Context, id int64) (*PopularContacts, error) {
	query := fmt.Sprintf("select %s from popular_contacts where id = ? limit 1", popularContactsRows)
	var resp PopularContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("popular_contacts.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultPopularContactsModel) FindListByIdList(ctx context.Context, id ...int64) ([]PopularContacts, error) {
	if len(id) == 0 {
		return []PopularContacts{}, nil
	}

	query := fmt.Sprintf("select %s from popular_contacts where id in (%s)", popularContactsRows, sqlx.InInt64List(id))

	var resp []PopularContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("popular_contacts.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultPopularContactsModel) Update2(ctx context.Context, data *PopularContacts) error {
	query := fmt.Sprintf("update `popular_contacts` set %s where `id` = ?", popularContactsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.Phone, data.Importers, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("popular_contacts.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultPopularContactsModel) FindOneByPhone(ctx context.Context, phone string) (*PopularContacts, error) {
	query := fmt.Sprintf("select %s from popular_contacts where phone = ? limit 1", popularContactsRows)
	var resp PopularContacts

	err := m.db.QueryRowPartial(ctx, &resp, query, phone)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("popular_contacts.FindOneByPhone: %w", err)
	}

	return &resp, nil
}

func (m *defaultPopularContactsModel) FindListByPhoneList(ctx context.Context, phone ...string) ([]PopularContacts, error) {
	if len(phone) == 0 {
		return []PopularContacts{}, nil
	}

	query := fmt.Sprintf("select %s from popular_contacts where phone in (%s)", popularContactsRows, sqlx.InStringList(phone))
	var resp []PopularContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("popular_contacts.FindListByPhoneList: %w", err)
	}

	return resp, nil
}
