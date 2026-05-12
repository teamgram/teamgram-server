package core

import (
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func mapChatMessageError(err error) error {
	if mapped, ok := mapKnownChatMessageError(err); ok {
		return mapped
	}
	return err
}

func mapKnownChatMessageError(err error) (error, bool) {
	if err == nil {
		return nil, false
	}
	switch {
	case isChatMessageServiceError(err, chatpb.ErrChatStorage):
		return tg.ErrInternalServerError, true
	case isChatMessageServiceError(err, chatpb.ErrChatNotFound),
		isChatMessageServiceError(err, chatpb.ErrChatMigrated),
		isChatMessageServiceError(err, chatpb.ErrChatDeactivated),
		isChatMessageServiceError(err, chatpb.ErrParticipantInvalid),
		isChatMessageServiceError(err, chatpb.ErrUserNotParticipant):
		return tg.Err400PeerIdInvalid, true
	case isChatMessageServiceError(err, chatpb.ErrChatAdminRequired):
		return tg.Err400ChatAdminRequired, true
	case isChatMessageServiceError(err, chatpb.ErrChatWriteForbidden):
		return tg.ErrChatWriteForbidden, true
	case isChatMessageServiceError(err, chatpb.ErrMessageActionUnsupported):
		return tg.ErrMethodNotImpl, true
	default:
		return nil, false
	}
}
