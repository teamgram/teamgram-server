package xerr

import (
	"fmt"
)

type ErrWrapper interface {
	Wrap() error
	Unwrap() error
	WrapMsg(msg string, kv ...any) error
	error
}

func NewErrorWrapper(err error, s string) ErrWrapper {
	return &errorWrapper{error: err, s: s}
}

type errorWrapper struct {
	error
	s string
}

func (e *errorWrapper) Error() string {
	return fmt.Sprintf("%s %s", e.error, e.s)
}

func (e *errorWrapper) Wrap() error {
	return Wrap(e)
}

func (e *errorWrapper) WrapMsg(msg string, kv ...any) error {
	return WrapMsg(e, msg, kv...)
}

func (e *errorWrapper) Unwrap() error {
	return e.error
}
