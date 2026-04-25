package stack

import (
	"errors"
	"runtime"
)

func callers(skip int) []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	return pcs[0:n]
}

func New(err error, skip int) error {
	return &stackError{
		err:   err,
		stack: callers(skip),
	}
}

type stackError struct {
	err   error
	stack []uintptr
}

func (e *stackError) Unwrap() error {
	return e.err
}

func (e *stackError) Cause() error {
	return e.err
}

func (e *stackError) Is(err error) bool {
	if e == nil {
		return err == nil
	}
	if err == nil {
		return false
	}
	if e == err || e.err == err {
		return true
	}
	return errors.Is(e.err, err)
}

func (e *stackError) Error() string {
	return e.err.Error()
}

func (e *stackError) String() string {
	return e.Error()
}
