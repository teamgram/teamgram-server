package presence

import "errors"

var (
	ErrPresenceInvalidArgument      = errors.New("presence: invalid argument")
	ErrPresencePermissionDenied     = errors.New("presence: permission denied")
	ErrPresenceStorage              = errors.New("presence: storage failure")
	ErrPresenceQuotaExceeded        = errors.New("presence: quota exceeded")
	ErrPresenceCorruptEntry         = errors.New("presence: corrupt entry")
	ErrPresenceMethodNotImplemented = errors.New("presence: method not implemented")
)
