package xerr

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/xerr/stack"
)

const (
	ServerInternalError = 500 // Server internal error
)

func ErrPanic(r any) error {
	return ErrPanicMsg(r, ServerInternalError, "panic error", 9)
}

func ErrPanicMsg(r any, code int, msg string, skip int) error {
	if r == nil {
		return nil
	}
	err := &codeError{
		code:   code,
		msg:    msg,
		detail: fmt.Sprint(r),
	}
	return stack.New(err, skip)
}
