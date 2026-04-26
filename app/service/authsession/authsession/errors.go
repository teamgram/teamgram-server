package authsession

import "errors"

var (
	ErrAuthKeyNotFound         = errors.New("authsession: auth key not found")
	ErrAuthKeyInvalid          = errors.New("authsession: auth key invalid")
	ErrPermAuthKeyEmpty        = errors.New("authsession: permanent auth key empty")
	ErrClientSessionEmpty      = errors.New("authsession: client session empty")
	ErrEncryptedMessageInvalid = errors.New("authsession: encrypted message invalid")
	ErrAuthorizationNotFound   = errors.New("authsession: authorization not found")
	ErrAuthSessionStorage      = errors.New("authsession: storage failure")
)
