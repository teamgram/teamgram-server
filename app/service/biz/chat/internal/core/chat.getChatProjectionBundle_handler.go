package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
)

// ChatGetChatProjectionBundle
// chat.getChatProjectionBundle viewer_user_ids:Vector<long> target_chat_ids:Vector<long> = ChatProjectionBundle;
func (c *ChatCore) ChatGetChatProjectionBundle(in *chat.TLChatGetChatProjectionBundle) (*chat.ChatProjectionBundle, error) {
	if in == nil {
		return nil, chat.ErrChatInvalidArgument
	}
	if len(in.ViewerUserIds) == 0 || len(in.TargetChatIds) == 0 {
		return emptyChatProjectionBundle(), nil
	}

	bundle, err := c.repo().GetChatProjectionBundle(c.ctx, in.ViewerUserIds, in.TargetChatIds)
	if err != nil {
		return nil, err
	}
	if bundle == nil {
		return nil, fmt.Errorf("%w: nil projection bundle", chat.ErrChatStorage)
	}

	return chat.MakeTLChatProjectionBundle(&chat.TLChatProjectionBundle{
		ViewerChats:    repositoryViewerChatsToTL(bundle.ViewerChats),
		MissingChatIds: bundle.MissingChatIds,
	}).ToChatProjectionBundle(), nil
}

func emptyChatProjectionBundle() *chat.ChatProjectionBundle {
	return chat.MakeTLChatProjectionBundle(&chat.TLChatProjectionBundle{}).ToChatProjectionBundle()
}

func repositoryViewerChatsToTL(in []repository.ViewerChats) []chat.ViewerChatsClazz {
	out := make([]chat.ViewerChatsClazz, 0, len(in))
	for _, item := range in {
		out = append(out, chat.MakeTLViewerChats(&chat.TLViewerChats{
			ViewerUserId: item.ViewerUserId,
			Chats:        item.Chats,
		}).ToViewerChats())
	}
	return out
}
