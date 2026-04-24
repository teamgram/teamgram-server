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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB

type (
	bizBotsModel interface {
		Select(ctx context.Context, botId int64) (*Bots, error)

		SelectByToken(ctx context.Context, token string) (int64, error)

		SelectByIdList(ctx context.Context, idList []int32) ([]Bots, error)
		SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *Bots)) ([]Bots, error)

		Update(ctx context.Context, cMap map[string]interface{}, botId int64) (rowsAffected int64, err error)
		UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, botId int64) (rowsAffected int64, err error)
	}
)

// Select
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id = :bot_id
func (m *defaultBotsModel) Select(ctx context.Context, botId int64) (rValue *Bots, err error) {

	var (
		query = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id = ?"
		do    = &Bots{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, botId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("bots.Select: %w", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByToken
// select bot_id from bots where token = :token
func (m *defaultBotsModel) SelectByToken(ctx context.Context, token string) (rValue int64, err error) {
	var query = "select bot_id from bots where token = ?"
	err = m.db.QueryRowPartial(ctx, &rValue, query, token)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("bots.SelectByToken: %w", err)
			return
		} else {
			err = nil
		}
	}

	return
}

// SelectByIdList
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (:id_list)
func (m *defaultBotsModel) SelectByIdList(ctx context.Context, idList []int32) (rList []Bots, err error) {
	var (
		query  = fmt.Sprintf("select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (%s)", sqlx.InInt32List(idList))
		values []Bots
	)
	if len(idList) == 0 {
		rList = []Bots{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		err = fmt.Errorf("bots.SelectByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (:id_list)
func (m *defaultBotsModel) SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *Bots)) (rList []Bots, err error) {
	var (
		query  = fmt.Sprintf("select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (%s)", sqlx.InInt32List(idList))
		values []Bots
	)
	if len(idList) == 0 {
		rList = []Bots{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		err = fmt.Errorf("bots.SelectByIdListWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// Update
// update bots set %s where bot_id = :bot_id
func (m *defaultBotsModel) Update(ctx context.Context, cMap map[string]interface{}, botId int64) (rowsAffected int64, err error) {

	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update bots set %s where bot_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, botId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		err = fmt.Errorf("bots.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("bots.Update rows affected: %w", err)
	}

	return
}

// UpdateTx
// update bots set %s where bot_id = :bot_id
func (m *defaultBotsModel) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, botId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update bots set %s where bot_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, botId)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		err = fmt.Errorf("bots.UpdateTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("bots.UpdateTx rows affected: %w", err)
	}

	return
}
