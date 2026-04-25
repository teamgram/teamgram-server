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
	if msg == "" {
		msg = "panic error"
	}
	_ = code
	return stack.New(fmt.Errorf("%s: %v", msg, r), skip)
}
