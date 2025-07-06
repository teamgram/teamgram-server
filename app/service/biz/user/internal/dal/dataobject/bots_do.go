/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type BotsDO struct {
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
}
