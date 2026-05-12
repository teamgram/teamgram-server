package core

import (
	"errors"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func mapChatMessageError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, chatpb.ErrChatStorage):
		return tg.ErrInternalServerError
	case errors.Is(err, chatpb.ErrChatNotFound),
		errors.Is(err, chatpb.ErrChatMigrated),
		errors.Is(err, chatpb.ErrChatDeactivated),
		errors.Is(err, chatpb.ErrParticipantInvalid),
		errors.Is(err, chatpb.ErrUserNotParticipant):
		return tg.Err400PeerIdInvalid
	case errors.Is(err, chatpb.ErrChatAdminRequired):
		return tg.Err400ChatAdminRequired
	case errors.Is(err, chatpb.ErrChatWriteForbidden):
		return tg.ErrChatWriteForbidden
	case errors.Is(err, chatpb.ErrMessageActionUnsupported):
		return tg.ErrMethodNotImpl
	default:
		return err
	}
}
