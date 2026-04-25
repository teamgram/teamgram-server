package xerr

import (
	"bytes"
	"fmt"
)

type Error interface {
	Wrap() error
	WrapMsg(msg string, kv ...any) error
	error
}

func New(s string, kv ...any) Error {
	return &errorString{
		s: toString(s, kv),
	}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func (e *errorString) Wrap() error {
	return Wrap(e)
}

func (e *errorString) WrapMsg(msg string, kv ...any) error {
	return WrapMsg(e, msg, kv...)
}

func toString(s string, kv []any) string {
	if len(kv) == 0 {
		return s
	} else {
		var buf bytes.Buffer
		buf.WriteString(s)

		for i := 0; i < len(kv); i += 2 {
			if buf.Len() > 0 {
				buf.WriteString(", ")
			}

			key := fmt.Sprintf("%v", kv[i])
			buf.WriteString(key)
			buf.WriteString("=")

			if i+1 < len(kv) {
				value := fmt.Sprintf("%v", kv[i+1])
				buf.WriteString(value)
			} else {
				buf.WriteString("MISSING")
			}
		}
		return buf.String()
	}
}
