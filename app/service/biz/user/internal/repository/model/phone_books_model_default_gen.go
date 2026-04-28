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
	phoneBooksFieldNames          = builder.RawFieldNames(&PhoneBooks{})
	phoneBooksRows                = strings.Join(phoneBooksFieldNames, ",")
	phoneBooksRowsExpectAutoSet   = strings.Join(stringx.Remove(phoneBooksFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	phoneBooksRowsWithPlaceHolder = strings.Join(stringx.Remove(phoneBooksFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
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
	tableName := "phone_books"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?)", tableName, phoneBooksRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.AuthKeyId, data.ClientId, data.Phone, data.FirstName, data.LastName)
	if err != nil {
		return nil, fmt.Errorf("phone_books.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultPhoneBooksModel) Delete2(ctx context.Context, id int64) error {
	tableName := "phone_books"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("phone_books.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultPhoneBooksModel) FindOne(ctx context.Context, id int64) (*PhoneBooks, error) {
	tableName := "phone_books"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", phoneBooksRows, tableName)
	var resp PhoneBooks

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "phone_books",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("phone_books.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultPhoneBooksModel) FindListByIdList(ctx context.Context, id ...int64) ([]PhoneBooks, error) {
	if len(id) == 0 {
		return []PhoneBooks{}, nil
	}
	tableName := "phone_books"

	query := fmt.Sprintf("select %s from %s where id in (%s)", phoneBooksRows, tableName, sqlx.InInt64List(id))

	var resp []PhoneBooks
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []PhoneBooks{}, nil
		}
		return nil, fmt.Errorf("phone_books.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultPhoneBooksModel) Update2(ctx context.Context, data *PhoneBooks) error {
	tableName := "phone_books"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, phoneBooksRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.AuthKeyId, data.ClientId, data.Phone, data.FirstName, data.LastName, data.Id)
	if err != nil {
		return fmt.Errorf("phone_books.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultPhoneBooksModel) FindOneByAuthKeyIdClientId(ctx context.Context, authKeyId int64, clientId int64) (*PhoneBooks, error) {
	tableName := "phone_books"
	query := fmt.Sprintf("select %s from %s where auth_key_id = ? AND client_id = ? limit 1", phoneBooksRows, tableName)
	var resp PhoneBooks

	err := m.db.QueryRowPartial(ctx, &resp, query, authKeyId, clientId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "phone_books",
				Key:      fmt.Sprintf("auth_key_id=%v,client_id=%v", authKeyId, clientId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("phone_books.FindOneByAuthKeyIdClientId: %w", err)
	}

	return &resp, nil
}
