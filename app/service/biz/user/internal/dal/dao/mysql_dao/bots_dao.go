/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type BotsDAO struct {
	db *sqlx.DB
}

func NewBotsDAO(db *sqlx.DB) *BotsDAO {
	return &BotsDAO{
		db: db,
	}
}

// Select
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id = :bot_id
func (dao *BotsDAO) Select(ctx context.Context, botId int64) (rValue *dataobject.BotsDO, err error) {
	var (
		query string
		do    = &dataobject.BotsDO{}
	)
	query = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id = ?"

	err = dao.db.QueryRowPartial(ctx, do, query, botId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByToken
// select bot_id from bots where token = :token
func (dao *BotsDAO) SelectByToken(ctx context.Context, token string) (rValue int64, err error) {
	var query string
	query = "select bot_id from bots where token = ?"

	err = dao.db.QueryRowPartial(ctx, &rValue, query, token)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("get in SelectByToken(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	}

	return
}

// SelectByIdList
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (:id_list)
func (dao *BotsDAO) SelectByIdList(ctx context.Context, idList []int32) (rList []dataobject.BotsDO, err error) {
	if len(idList) == 0 {
		rList = []dataobject.BotsDO{}
		return
	}

	var (
		query  string
		values []dataobject.BotsDO
	)
	query = fmt.Sprintf("select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (%s)", sqlx.InInt32List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (:id_list)
func (dao *BotsDAO) SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *dataobject.BotsDO)) (rList []dataobject.BotsDO, err error) {
	if len(idList) == 0 {
		rList = []dataobject.BotsDO{}
		return
	}

	var (
		query  string
		values []dataobject.BotsDO
	)
	query = fmt.Sprintf("select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (%s)", sqlx.InInt32List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// Update
// update bots set %s where bot_id = :bot_id
func (dao *BotsDAO) Update(ctx context.Context, cMap map[string]interface{}, botId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update bots set %s where bot_id = ?", strings.Join(names, ", "))

	aValues = append(aValues, botId)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

// UpdateTx
// update bots set %s where bot_id = :bot_id
func (dao *BotsDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, botId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update bots set %s where bot_id = ?", strings.Join(names, ", "))

	aValues = append(aValues, botId)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}
