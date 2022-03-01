// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package model

import (
	"strings"

	"github.com/teamgram/marmota/pkg/random2"
	"github.com/teamgram/marmota/pkg/utils"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/pkg/env2"
)

func GetChatTypeByInviteHash(hash string) int {
	if len(hash) != 20 {
		return mtproto.PEER_UNKNOWN
	}

	if utils.IsLetter(hash[0]) {
		return mtproto.PEER_CHANNEL
	} else if utils.IsNumber(hash[0]) {
		return mtproto.PEER_CHAT
	} else {
		return mtproto.PEER_UNKNOWN
	}
}

func GenChatInviteHash() string {
	return random2.RandomNumeric(1) + random2.RandomAlphanumeric(19)
}

func GenChannelInviteHash() string {
	return random2.RandomAlphabetic(1) + random2.RandomAlphanumeric(19)
}

func GetInviteHashByLink(link string) string {
	if strings.HasPrefix(link, "https://"+env2.TDotMe+"/+") {
		link = link[len("https://"+env2.TDotMe+"/+"):]
	}
	return link
}
