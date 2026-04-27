package code

import "errors"

var (
	ErrPhoneCodeExpired = errors.New("code: phone code expired")
	ErrPhoneCodeInvalid = errors.New("code: phone code invalid")
	ErrCodeStorage      = errors.New("code: storage failure")
)
