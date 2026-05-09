package msg

import "errors"

var (
	ErrRandomIdConflict      = errors.New("msg: random id conflict")
	ErrReplyToInvalid        = errors.New("msg: reply_to invalid")
	ErrMessageAuthorRequired = errors.New("msg: message author required")
	ErrMessageNotModified    = errors.New("msg: message not modified")
	ErrSendStateConflict     = errors.New("msg: send state conflict")
	ErrSenderSyncFailed      = errors.New("msg: sender sync failed")
	ErrReceiverBackpressure  = errors.New("msg: receiver backpressure")
	ErrMsgStorage            = errors.New("msg: storage failure")
	ErrMethodNotImplemented  = errors.New("msg: method not implemented")
	ErrMsgIdInvalid          = errors.New("msg: message id invalid")
)
