package xerr

import (
	"errors"
	"fmt"
)

type ErrWrapper interface {
	Is(err error) bool
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

func (e *errorWrapper) Is(err error) bool {
	if err == nil {
		return false
	}
	var t *errorWrapper
	ok := errors.As(err, &t)
	return ok && e.s == t.s
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
