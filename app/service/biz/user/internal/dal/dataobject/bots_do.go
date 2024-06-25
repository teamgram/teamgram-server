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
	Id                   int64  `db:"id" json:"id"`
	BotId                int64  `db:"bot_id" json:"bot_id"`
	BotType              int32  `db:"bot_type" json:"bot_type"`
	CreatorUserId        int64  `db:"creator_user_id" json:"creator_user_id"`
	Token                string `db:"token" json:"token"`
	Description          string `db:"description" json:"description"`
	BotChatHistory       bool   `db:"bot_chat_history" json:"bot_chat_history"`
	BotNochats           bool   `db:"bot_nochats" json:"bot_nochats"`
	Verified             bool   `db:"verified" json:"verified"`
	BotInlineGeo         bool   `db:"bot_inline_geo" json:"bot_inline_geo"`
	BotInfoVersion       int32  `db:"bot_info_version" json:"bot_info_version"`
	BotInlinePlaceholder string `db:"bot_inline_placeholder" json:"bot_inline_placeholder"`
}
