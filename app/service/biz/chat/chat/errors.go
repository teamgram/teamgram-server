package chat

import (
	"errors"
	"fmt"
)

var (
	ErrChatNotFound           = errors.New("chat: chat not found")
	ErrChatMigrated           = errors.New("chat: chat migrated")
	ErrChatDeactivated        = errors.New("chat: chat deactivated")
	ErrChatAdminRequired      = errors.New("chat: chat admin required")
	ErrChatTitleEmpty         = errors.New("chat: title empty")
	ErrChatNotModified        = errors.New("chat: not modified")
	ErrCreateChatFlood        = errors.New("chat: create chat flood")
	ErrParticipantInvalid     = errors.New("chat: participant invalid")
	ErrInputUserDeactivated   = errors.New("chat: input user deactivated")
	ErrUserAlreadyParticipant = errors.New("chat: user already participant")
	ErrUserNotParticipant     = errors.New("chat: user not participant")
	ErrUsersTooFew            = errors.New("chat: users too few")
	ErrUsersTooMuch           = errors.New("chat: users too much")
	ErrInviteHashInvalid      = errors.New("chat: invite hash invalid")
	ErrInviteHashExpired      = errors.New("chat: invite hash expired")
	ErrChatLinkExists         = errors.New("chat: link exists")
	ErrChatStorage            = errors.New("chat: storage failure")
)

type CreateChatFloodError struct {
	WaitSeconds int32
}

func NewCreateChatFloodError(waitSeconds int32) error {
	return &CreateChatFloodError{WaitSeconds: waitSeconds}
}

func (e *CreateChatFloodError) Error() string {
	return fmt.Sprintf("%s: wait_seconds=%d", ErrCreateChatFlood, e.WaitSeconds)
}

func (e *CreateChatFloodError) Is(target error) bool {
	return target == ErrCreateChatFlood
}

func WrapChatStorage(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", ErrChatStorage, op, err)
}
