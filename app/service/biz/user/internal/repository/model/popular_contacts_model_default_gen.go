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
	popular_contactsFieldNames          = builder.RawFieldNames(&PopularContacts{})
	popular_contactsRows                = strings.Join(popular_contactsFieldNames, ",")
	popular_contactsRowsExpectAutoSet   = strings.Join(stringx.Remove(popular_contactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	popular_contactsRowsWithPlaceHolder = strings.Join(stringx.Remove(popular_contactsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTPopularContactsIdPrefix = "cache:t:popular_contacts:id:"

	cachePopularContactsIdPrefix = "cache#PopularContacts#id"

	cachePopularContactsPhonePrefix = "cache#Phone"
)

type (
	popular_contactsModel interface {
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
	query := fmt.Sprintf("insert into `popular_contacts` (%s) values (?, ?, ?)", popular_contactsRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.Phone, data.Importers, data.Deleted)
}

func (m *defaultPopularContactsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `popular_contacts` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultPopularContactsModel) FindOne(ctx context.Context, id int64) (*PopularContacts, error) {
	query := fmt.Sprintf("select %s from popular_contacts where id = ? limit 1", popular_contactsRows)
	var resp PopularContacts
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultPopularContactsModel) FindListByIdList(ctx context.Context, id ...int64) ([]PopularContacts, error) {
	if len(id) == 0 {
		return []PopularContacts{}, nil
	}

	query := fmt.Sprintf("select %s from popular_contacts where id in (%s)", popular_contactsRows, sqlx.InInt64List(id))

	var resp []PopularContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultPopularContactsModel) Update2(ctx context.Context, data *PopularContacts) error {
	query := fmt.Sprintf("update `popular_contacts` set %s where `id` = ?", popular_contactsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.Phone, data.Importers, data.Deleted, data.Id)
	return err
}

func (m *defaultPopularContactsModel) FindOneByPhone(ctx context.Context, phone string) (*PopularContacts, error) {
	query := fmt.Sprintf("select %s from popular_contacts where phone = ? limit 1", popular_contactsRows)
	var resp PopularContacts
	err := m.db.QueryRowPartial(ctx, &resp, query, phone)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultPopularContactsModel) FindListByPhoneList(ctx context.Context, phone ...string) ([]PopularContacts, error) {
	if len(phone) == 0 {
		return []PopularContacts{}, nil
	}

	query := fmt.Sprintf("select %s from popular_contacts where phone in (%s)", popular_contactsRows, sqlx.InStringList(phone))
	var resp []PopularContacts
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultPopularContactsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cachePopularContactsIdPrefix, primary)
}

func (m *defaultPopularContactsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from popular_contacts where id = ? limit 1", popular_contactsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
