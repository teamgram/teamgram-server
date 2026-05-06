package msg

import "errors"

var (
	ErrRandomIdConflict     = errors.New("msg: random id conflict")
	ErrReplyToInvalid       = errors.New("msg: reply_to invalid")
	ErrSendStateConflict    = errors.New("msg: send state conflict")
	ErrSenderSyncFailed     = errors.New("msg: sender sync failed")
	ErrReceiverBackpressure = errors.New("msg: receiver backpressure")
	ErrMsgStorage           = errors.New("msg: storage failure")
)
