package core

import (
	"errors"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func mapChatError(err error) error {
	if err == nil {
		return nil
	}

	var floodErr *chatpb.CreateChatFloodError
	var pendingErr *chatpb.CreateChatOperationPendingError
	switch {
	case errors.Is(err, chatpb.ErrChatStorage):
		return tg.ErrInternalServerError
	case errors.As(err, &floodErr):
		return tg.NewErrFloodWaitX(floodErr.WaitSeconds)
	case errors.As(err, &pendingErr):
		return tg.NewErrFloodWaitX(pendingErr.WaitSeconds)
	case errors.Is(err, chatpb.ErrChatNotFound),
		errors.Is(err, chatpb.ErrChatMigrated),
		errors.Is(err, chatpb.ErrChatDeactivated):
		return tg.ErrChatIdInvalid
	case errors.Is(err, chatpb.ErrChatAdminRequired):
		return tg.Err400ChatAdminRequired
	case errors.Is(err, chatpb.ErrChatTitleEmpty):
		return tg.ErrChatTitleEmpty
	case errors.Is(err, chatpb.ErrChatNotModified):
		return tg.ErrChatNotModified
	case errors.Is(err, chatpb.ErrParticipantInvalid):
		return tg.Err400PeerIdInvalid
	case errors.Is(err, chatpb.ErrInputUserDeactivated):
		return tg.ErrInputUserDeactivated
	case errors.Is(err, chatpb.ErrUserAlreadyParticipant):
		return tg.ErrUserAlreadyParticipant
	case errors.Is(err, chatpb.ErrUserNotParticipant):
		return tg.Err400UserNotParticipant
	case errors.Is(err, chatpb.ErrUsersTooFew):
		return tg.ErrUsersTooFew
	case errors.Is(err, chatpb.ErrUsersTooMuch):
		return tg.ErrUsersTooMuch
	case errors.Is(err, chatpb.ErrInviteHashInvalid):
		return tg.ErrInviteHashInvalid
	case errors.Is(err, chatpb.ErrInviteHashExpired):
		return tg.Err400InviteHashExpired
	case errors.Is(err, chatpb.ErrChatLinkExists):
		return tg.ErrChatLinkExists
	default:
		return err
	}
}
