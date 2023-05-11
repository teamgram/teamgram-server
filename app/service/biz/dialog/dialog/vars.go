// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dialog

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"
)

const (
	dialogKeyPrefix = "dialog"
)

const (
	dialogFiltersKeyPrefix = "dialog_filters"
)

const (
	conversationsKeyPrefix = "user_conversations"
	chatsKeyPrefix         = "user_chats"
	channelsKeyPrefix      = "user_channels"
)

func GenDialogCacheKey(userId, peerDialogId int64) string {
	return fmt.Sprintf("%s#%d_%d", dialogKeyPrefix, userId, peerDialogId)
}

func GenDialogCacheKeyByPeer(userId int64, peerType int32, peerId int64) string {
	return GenDialogCacheKey(userId, mtproto.MakePeerDialogId(peerType, peerId))
}

func GenDialogFilterCacheKey(userId int64) string {
	return fmt.Sprintf("%s#%d", dialogFiltersKeyPrefix, userId)
}

func GenConversationsCacheKey(userId int64) string {
	return fmt.Sprintf("%s#%d", conversationsKeyPrefix, userId)
}

func GenChatsCacheKey(userId int64) string {
	return fmt.Sprintf("%s#%d", chatsKeyPrefix, userId)
}

func GenChannelsCacheKey(userId int64) string {
	return fmt.Sprintf("%s#%d", channelsKeyPrefix, userId)
}

func GenCacheKeyByPeerType(userId int64, peerType int32) string {
	switch peerType {
	case mtproto.PEER_USER:
		return GenConversationsCacheKey(userId)
	case mtproto.PEER_CHAT:
		return GenChatsCacheKey(userId)
	case mtproto.PEER_CHANNEL:
		return GenChannelsCacheKey(userId)
	}

	return ""
}
