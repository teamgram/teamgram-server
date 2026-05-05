package dialog

import (
	"errors"
	"fmt"
)

var (
	ErrDialogStorage         = errors.New("dialog: storage failure")
	ErrDialogNotFound        = errors.New("dialog: dialog not found")
	ErrInvalidPeer           = errors.New("dialog: invalid peer")
	ErrDialogInvalid         = errors.New("dialog: invalid request")
	ErrDeprecatedMethod      = errors.New("dialog: deprecated method")
	ErrWrongOwner            = errors.New("dialog: wrong owner")
	ErrSourceAuthKeyRequired = errors.New("dialog: source auth key required")
	ErrPayloadConflict       = errors.New("dialog: payload conflict")
	ErrOutboxUnavailable     = errors.New("dialog: outbox unavailable")
)

func WrapDialogStorage(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", ErrDialogStorage, op, err)
}
