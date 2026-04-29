package dialog

import (
	"errors"
	"fmt"
)

var (
	ErrDialogNotFound = errors.New("dialog: not found")
	ErrDialogStorage  = errors.New("dialog: storage failure")
	ErrDialogInvalid  = errors.New("dialog: invalid request")
)

func WrapDialogStorage(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", ErrDialogStorage, op, err)
}
