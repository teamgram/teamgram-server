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
	usernameFieldNames          = builder.RawFieldNames(&Username{})
	usernameRows                = strings.Join(usernameFieldNames, ",")
	usernameRowsExpectAutoSet   = strings.Join(stringx.Remove(usernameFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usernameRowsWithPlaceHolder = strings.Join(stringx.Remove(usernameFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	usernameModel interface {
		Insert2(ctx context.Context, data *Username) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Username, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Username, error)
		Update2(ctx context.Context, data *Username) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUsername(ctx context.Context, username string) (*Username, error)
		FindListByUsernameList(ctx context.Context, username ...string) ([]Username, error)
	}

	defaultUsernameModel struct {
		db *sqlx.DB
	}

	Username struct {
		Id       int64  `db:"id" json:"id"`
		Username string `db:"username" json:"username"`
		PeerType int32  `db:"peer_type" json:"peer_type"`
		PeerId   int64  `db:"peer_id" json:"peer_id"`
		Editable bool   `db:"editable" json:"editable"`
		Active   bool   `db:"active" json:"active"`
		Order2   int64  `db:"order2" json:"order2"`
		Deleted  bool   `db:"deleted" json:"deleted"`
	}
)

func newUsernameModel(db *sqlx.DB) *defaultUsernameModel {
	return &defaultUsernameModel{
		db: db,
	}
}

func (m *defaultUsernameModel) Insert2(ctx context.Context, data *Username) (sql.Result, error) {
	query := fmt.Sprintf("insert into `username` (%s) values (?, ?, ?, ?, ?, ?, ?)", usernameRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.Username, data.PeerType, data.PeerId, data.Editable, data.Active, data.Order2, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("username.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUsernameModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `username` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("username.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUsernameModel) FindOne(ctx context.Context, id int64) (*Username, error) {
	query := fmt.Sprintf("select %s from username where id = ? limit 1", usernameRows)
	var resp Username

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("username.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUsernameModel) FindListByIdList(ctx context.Context, id ...int64) ([]Username, error) {
	if len(id) == 0 {
		return []Username{}, nil
	}

	query := fmt.Sprintf("select %s from username where id in (%s)", usernameRows, sqlx.InInt64List(id))

	var resp []Username
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Username{}, nil
		}
		return nil, fmt.Errorf("username.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUsernameModel) Update2(ctx context.Context, data *Username) error {
	query := fmt.Sprintf("update `username` set %s where `id` = ?", usernameRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.Username, data.PeerType, data.PeerId, data.Editable, data.Active, data.Order2, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("username.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUsernameModel) FindOneByUsername(ctx context.Context, username string) (*Username, error) {
	query := fmt.Sprintf("select %s from username where username = ? limit 1", usernameRows)
	var resp Username

	err := m.db.QueryRowPartial(ctx, &resp, query, username)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "username",
				Key:      fmt.Sprintf("username=%v", username),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("username.FindOneByUsername: %w", err)
	}

	return &resp, nil
}

func (m *defaultUsernameModel) FindListByUsernameList(ctx context.Context, username ...string) ([]Username, error) {
	if len(username) == 0 {
		return []Username{}, nil
	}

	query := fmt.Sprintf("select %s from username where username in (%s)", usernameRows, sqlx.InStringList(username))
	var resp []Username
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Username{}, nil
		}
		return nil, fmt.Errorf("username.FindListByUsernameList: %w", err)
	}

	return resp, nil
}
