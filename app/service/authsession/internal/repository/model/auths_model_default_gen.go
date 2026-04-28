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
	authsFieldNames          = builder.RawFieldNames(&Auths{})
	authsRows                = strings.Join(authsFieldNames, ",")
	authsRowsExpectAutoSet   = strings.Join(stringx.Remove(authsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authsRowsWithPlaceHolder = strings.Join(stringx.Remove(authsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	authsModel interface {
		Insert2(ctx context.Context, data *Auths) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Auths, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Auths, error)
		Update2(ctx context.Context, data *Auths) error
		Delete2(ctx context.Context, id int64) error

		FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*Auths, error)
		FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]Auths, error)
	}

	defaultAuthsModel struct {
		db *sqlx.DB
	}

	Auths struct {
		Id             int64  `db:"id" json:"id"`
		AuthKeyId      int64  `db:"auth_key_id" json:"auth_key_id"`
		Layer          int32  `db:"layer" json:"layer"`
		ApiId          int32  `db:"api_id" json:"api_id"`
		DeviceModel    string `db:"device_model" json:"device_model"`
		SystemVersion  string `db:"system_version" json:"system_version"`
		AppVersion     string `db:"app_version" json:"app_version"`
		SystemLangCode string `db:"system_lang_code" json:"system_lang_code"`
		LangPack       string `db:"lang_pack" json:"lang_pack"`
		LangCode       string `db:"lang_code" json:"lang_code"`
		SystemCode     string `db:"system_code" json:"system_code"`
		Proxy          string `db:"proxy" json:"proxy"`
		Params         string `db:"params" json:"params"`
		ClientIp       string `db:"client_ip" json:"client_ip"`
		DateActive     int64  `db:"date_active" json:"date_active"`
		Deleted        bool   `db:"deleted" json:"deleted"`
	}
)

func newAuthsModel(db *sqlx.DB) *defaultAuthsModel {
	return &defaultAuthsModel{
		db: db,
	}
}

func (m *defaultAuthsModel) Insert2(ctx context.Context, data *Auths) (sql.Result, error) {
	query := fmt.Sprintf("insert into `auths` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", authsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.AuthKeyId, data.Layer, data.ApiId, data.DeviceModel, data.SystemVersion, data.AppVersion, data.SystemLangCode, data.LangPack, data.LangCode, data.SystemCode, data.Proxy, data.Params, data.ClientIp, data.DateActive, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("auths.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultAuthsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auths` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("auths.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthsModel) FindOne(ctx context.Context, id int64) (*Auths, error) {
	query := fmt.Sprintf("select %s from auths where id = ? limit 1", authsRows)
	var resp Auths

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auths",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("auths.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthsModel) FindListByIdList(ctx context.Context, id ...int64) ([]Auths, error) {
	if len(id) == 0 {
		return []Auths{}, nil
	}

	query := fmt.Sprintf("select %s from auths where id in (%s)", authsRows, sqlx.InInt64List(id))

	var resp []Auths
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Auths{}, nil
		}
		return nil, fmt.Errorf("auths.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultAuthsModel) Update2(ctx context.Context, data *Auths) error {
	query := fmt.Sprintf("update `auths` set %s where `id` = ?", authsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.AuthKeyId, data.Layer, data.ApiId, data.DeviceModel, data.SystemVersion, data.AppVersion, data.SystemLangCode, data.LangPack, data.LangCode, data.SystemCode, data.Proxy, data.Params, data.ClientIp, data.DateActive, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("auths.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthsModel) FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*Auths, error) {
	query := fmt.Sprintf("select %s from auths where auth_key_id = ? limit 1", authsRows)
	var resp Auths

	err := m.db.QueryRowPartial(ctx, &resp, query, authKeyId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auths",
				Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("auths.FindOneByAuthKeyId: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthsModel) FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]Auths, error) {
	if len(authKeyId) == 0 {
		return []Auths{}, nil
	}

	query := fmt.Sprintf("select %s from auths where auth_key_id in (%s)", authsRows, sqlx.InInt64List(authKeyId))

	var resp []Auths
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Auths{}, nil
		}
		return nil, fmt.Errorf("auths.FindListByAuthKeyIdList: %w", err)
	}

	return resp, nil
}
