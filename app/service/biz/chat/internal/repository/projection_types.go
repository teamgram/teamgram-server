package repository

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

type ChatProjectionBundle struct {
	ViewerChats    []ViewerChats
	MissingChatIds []int64
}

type ViewerChats struct {
	ViewerUserId int64
	Chats        []tg.ChatClazz
}
