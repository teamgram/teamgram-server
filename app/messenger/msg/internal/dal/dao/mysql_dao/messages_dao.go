/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	message_helper "github.com/teamgram/teamgram-server/app/service/biz/message"
)

type (
	MessagesDAO = message_helper.MessagesDAO
)

var (
	NewMessagesDAO = message_helper.NewMessagesDAO
)
