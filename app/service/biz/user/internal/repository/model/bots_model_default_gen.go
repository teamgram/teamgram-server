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
	botsFieldNames          = builder.RawFieldNames(&Bots{})
	botsRows                = strings.Join(botsFieldNames, ",")
	botsRowsExpectAutoSet   = strings.Join(stringx.Remove(botsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	botsRowsWithPlaceHolder = strings.Join(stringx.Remove(botsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	botsModel interface {
		Insert2(ctx context.Context, data *Bots) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Bots, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Bots, error)
		Update2(ctx context.Context, data *Bots) error
		Delete2(ctx context.Context, id int64) error

		FindOneByBotId(ctx context.Context, botId int64) (*Bots, error)
		FindListByBotIdList(ctx context.Context, botId ...int64) ([]Bots, error)
	}

	defaultBotsModel struct {
		db *sqlx.DB
	}

	Bots struct {
		Id                    int64  `db:"id" json:"id"`
		BotId                 int64  `db:"bot_id" json:"bot_id"`
		BotType               int32  `db:"bot_type" json:"bot_type"`
		CreatorUserId         int64  `db:"creator_user_id" json:"creator_user_id"`
		Token                 string `db:"token" json:"token"`
		Description           string `db:"description" json:"description"`
		BotChatHistory        bool   `db:"bot_chat_history" json:"bot_chat_history"`
		BotNochats            bool   `db:"bot_nochats" json:"bot_nochats"`
		Verified              bool   `db:"verified" json:"verified"`
		BotInlineGeo          bool   `db:"bot_inline_geo" json:"bot_inline_geo"`
		BotInfoVersion        int32  `db:"bot_info_version" json:"bot_info_version"`
		BotInlinePlaceholder  string `db:"bot_inline_placeholder" json:"bot_inline_placeholder"`
		BotAttachMenu         bool   `db:"bot_attach_menu" json:"bot_attach_menu"`
		AttachMenuEnabled     bool   `db:"attach_menu_enabled" json:"attach_menu_enabled"`
		BotBusiness           bool   `db:"bot_business" json:"bot_business"`
		BotHasMainApp         bool   `db:"bot_has_main_app" json:"bot_has_main_app"`
		BotActiveUsers        int32  `db:"bot_active_users" json:"bot_active_users"`
		HasMenuButton         bool   `db:"has_menu_button" json:"has_menu_button"`
		MenuButtonText        string `db:"menu_button_text" json:"menu_button_text"`
		MenuButtonUrl         string `db:"menu_button_url" json:"menu_button_url"`
		BotCanEdit            bool   `db:"bot_can_edit" json:"bot_can_edit"`
		HasPreviewMedias      bool   `db:"has_preview_medias" json:"has_preview_medias"`
		DescriptionPhotoId    int64  `db:"description_photo_id" json:"description_photo_id"`
		DescriptionDocumentId int64  `db:"description_document_id" json:"description_document_id"`
		MainAppUrl            string `db:"main_app_url" json:"main_app_url"`
		HasAppSettings        bool   `db:"has_app_settings" json:"has_app_settings"`
		PlaceholderPath       string `db:"placeholder_path" json:"placeholder_path"`
		BackgroundColor       int32  `db:"background_color" json:"background_color"`
		BackgroundDarkColor   int32  `db:"background_dark_color" json:"background_dark_color"`
		HeaderColor           int32  `db:"header_color" json:"header_color"`
		HeaderDarkColor       int32  `db:"header_dark_color" json:"header_dark_color"`
		PrivacyPolicyUrl      string `db:"privacy_policy_url" json:"privacy_policy_url"`
		Mode                  int32  `db:"mode" json:"mode"`
	}
)

func newBotsModel(db *sqlx.DB) *defaultBotsModel {
	return &defaultBotsModel{
		db: db,
	}
}

func (m *defaultBotsModel) Insert2(ctx context.Context, data *Bots) (sql.Result, error) {
	tableName := "bots"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, botsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.BotId, data.BotType, data.CreatorUserId, data.Token, data.Description, data.BotChatHistory, data.BotNochats, data.Verified, data.BotInlineGeo, data.BotInfoVersion, data.BotInlinePlaceholder, data.BotAttachMenu, data.AttachMenuEnabled, data.BotBusiness, data.BotHasMainApp, data.BotActiveUsers, data.HasMenuButton, data.MenuButtonText, data.MenuButtonUrl, data.BotCanEdit, data.HasPreviewMedias, data.DescriptionPhotoId, data.DescriptionDocumentId, data.MainAppUrl, data.HasAppSettings, data.PlaceholderPath, data.BackgroundColor, data.BackgroundDarkColor, data.HeaderColor, data.HeaderDarkColor, data.PrivacyPolicyUrl, data.Mode)
	if err != nil {
		return nil, fmt.Errorf("bots.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultBotsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "bots"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("bots.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultBotsModel) FindOne(ctx context.Context, id int64) (*Bots, error) {
	tableName := "bots"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", botsRows, tableName)
	var resp Bots

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "bots",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("bots.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultBotsModel) FindListByIdList(ctx context.Context, id ...int64) ([]Bots, error) {
	if len(id) == 0 {
		return []Bots{}, nil
	}
	tableName := "bots"

	query := fmt.Sprintf("select %s from %s where id in (%s)", botsRows, tableName, sqlx.InInt64List(id))

	var resp []Bots
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Bots{}, nil
		}
		return nil, fmt.Errorf("bots.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultBotsModel) Update2(ctx context.Context, data *Bots) error {
	tableName := "bots"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, botsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.BotId, data.BotType, data.CreatorUserId, data.Token, data.Description, data.BotChatHistory, data.BotNochats, data.Verified, data.BotInlineGeo, data.BotInfoVersion, data.BotInlinePlaceholder, data.BotAttachMenu, data.AttachMenuEnabled, data.BotBusiness, data.BotHasMainApp, data.BotActiveUsers, data.HasMenuButton, data.MenuButtonText, data.MenuButtonUrl, data.BotCanEdit, data.HasPreviewMedias, data.DescriptionPhotoId, data.DescriptionDocumentId, data.MainAppUrl, data.HasAppSettings, data.PlaceholderPath, data.BackgroundColor, data.BackgroundDarkColor, data.HeaderColor, data.HeaderDarkColor, data.PrivacyPolicyUrl, data.Mode, data.Id)
	if err != nil {
		return fmt.Errorf("bots.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultBotsModel) FindOneByBotId(ctx context.Context, botId int64) (*Bots, error) {
	tableName := "bots"
	query := fmt.Sprintf("select %s from %s where bot_id = ? limit 1", botsRows, tableName)
	var resp Bots

	err := m.db.QueryRowPartial(ctx, &resp, query, botId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "bots",
				Key:      fmt.Sprintf("bot_id=%v", botId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("bots.FindOneByBotId: %w", err)
	}

	return &resp, nil
}

func (m *defaultBotsModel) FindListByBotIdList(ctx context.Context, botId ...int64) ([]Bots, error) {
	if len(botId) == 0 {
		return []Bots{}, nil
	}
	tableName := "bots"

	query := fmt.Sprintf("select %s from %s where bot_id in (%s)", botsRows, tableName, sqlx.InInt64List(botId))

	var resp []Bots
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Bots{}, nil
		}
		return nil, fmt.Errorf("bots.FindListByBotIdList: %w", err)
	}

	return resp, nil
}
