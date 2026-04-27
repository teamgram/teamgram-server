package user

import "errors"

var (
	ErrUserNotFound      = errors.New("user: not found")
	ErrUserStorage       = errors.New("user: storage failure")
	ErrUsernameNotFound  = errors.New("user: username not found")
	ErrUsernameInvalid   = errors.New("user: username invalid")
	ErrContactNotFound   = errors.New("user: contact not found")
	ErrBotNotFound       = errors.New("user: bot not found")
	ErrPrivacyKeyInvalid = errors.New("user: privacy key invalid")
)
