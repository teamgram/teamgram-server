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
	phoneBooksFieldNames          = builder.RawFieldNames(&PhoneBooks{})
	phoneBooksRows                = strings.Join(phoneBooksFieldNames, ",")
	phoneBooksRowsExpectAutoSet   = strings.Join(stringx.Remove(phoneBooksFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	phoneBooksRowsWithPlaceHolder = strings.Join(stringx.Remove(phoneBooksFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTPhoneBooksIdPrefix = "cache:t:phone_books:id:"

	cachePhoneBooksIdPrefix = "cache#PhoneBooks#id"

	cachePhoneBooksAuthKeyIdClientIdPrefix = "cache#AuthKeyId#ClientId"
)

type (
	phoneBooksModel interface {
		Insert2(ctx context.Context, data *PhoneBooks) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PhoneBooks, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]PhoneBooks, error)
		Update2(ctx context.Context, data *PhoneBooks) error
		Delete2(ctx context.Context, id int64) error

		FindOneByAuthKeyIdClientId(ctx context.Context, authKeyId int64, clientId int64) (*PhoneBooks, error)
	}

	defaultPhoneBooksModel struct {
		db *sqlx.DB
	}

	PhoneBooks struct {
		Id        int64  `db:"id" json:"id"`
		UserId    int64  `db:"user_id" json:"user_id"`
		AuthKeyId int64  `db:"auth_key_id" json:"auth_key_id"`
		ClientId  int64  `db:"client_id" json:"client_id"`
		Phone     string `db:"phone" json:"phone"`
		FirstName string `db:"first_name" json:"first_name"`
		LastName  string `db:"last_name" json:"last_name"`
	}
)

func newPhoneBooksModel(db *sqlx.DB) *defaultPhoneBooksModel {
	return &defaultPhoneBooksModel{
		db: db,
	}
}

func (m *defaultPhoneBooksModel) Insert2(ctx context.Context, data *PhoneBooks) (sql.Result, error) {
	query := fmt.Sprintf("insert into `phone_books` (%s) values (?, ?, ?, ?, ?, ?)", phoneBooksRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.AuthKeyId, data.ClientId, data.Phone, data.FirstName, data.LastName)
}

func (m *defaultPhoneBooksModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `phone_books` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultPhoneBooksModel) FindOne(ctx context.Context, id int64) (*PhoneBooks, error) {
	query := fmt.Sprintf("select %s from phone_books where id = ? limit 1", phoneBooksRows)
	var resp PhoneBooks
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultPhoneBooksModel) FindListByIdList(ctx context.Context, id ...int64) ([]PhoneBooks, error) {
	if len(id) == 0 {
		return []PhoneBooks{}, nil
	}

	query := fmt.Sprintf("select %s from phone_books where id in (%s)", phoneBooksRows, sqlx.InInt64List(id))

	var resp []PhoneBooks
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultPhoneBooksModel) Update2(ctx context.Context, data *PhoneBooks) error {
	query := fmt.Sprintf("update `phone_books` set %s where `id` = ?", phoneBooksRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.AuthKeyId, data.ClientId, data.Phone, data.FirstName, data.LastName, data.Id)
	return err
}

func (m *defaultPhoneBooksModel) FindOneByAuthKeyIdClientId(ctx context.Context, authKeyId int64, clientId int64) (*PhoneBooks, error) {
	query := fmt.Sprintf("select %s from phone_books where auth_key_id = ? AND client_id = ? limit 1", phoneBooksRows)
	var resp PhoneBooks
	err := m.db.QueryRowPartial(ctx, &resp, query, authKeyId, clientId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultPhoneBooksModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cachePhoneBooksIdPrefix, primary)
}

func (m *defaultPhoneBooksModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from phone_books where id = ? limit 1", phoneBooksRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
