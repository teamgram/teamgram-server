/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type BotsDO struct {
	Id                   int64  `db:"id"`
	BotId                int64  `db:"bot_id"`
	BotType              int32  `db:"bot_type"`
	CreatorUserId        int64  `db:"creator_user_id"`
	Token                string `db:"token"`
	Description          string `db:"description"`
	BotChatHistory       bool   `db:"bot_chat_history"`
	BotNochats           bool   `db:"bot_nochats"`
	Verified             bool   `db:"verified"`
	BotInlineGeo         bool   `db:"bot_inline_geo"`
	BotInfoVersion       int32  `db:"bot_info_version"`
	BotInlinePlaceholder string `db:"bot_inline_placeholder"`
}
