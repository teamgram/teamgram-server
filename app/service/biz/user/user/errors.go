package user

import "errors"

var (
	ErrUserNotFound      = errors.New("user: not found")
	ErrUserStorage       = errors.New("user: storage failure")
	ErrUsernameNotFound  = errors.New("user: username not found")
	ErrUsernameInvalid   = errors.New("user: username invalid")
	ErrPhoneNumberInUse  = errors.New("user: phone number in use")
	ErrContactNotFound   = errors.New("user: contact not found")
	ErrBotNotFound       = errors.New("user: bot not found")
	ErrPrivacyKeyInvalid = errors.New("user: privacy key invalid")
)
