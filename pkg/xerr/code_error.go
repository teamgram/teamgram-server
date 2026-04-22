package xerr

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/teamgram/teamgram-server/v2/pkg/xerr/stack"
)

const stackSkip = 4

var DefaultCodeRelation = newCodeRelation()

type CodeError interface {
	Code() int
	Msg() string
	Detail() string
	WithDetail(detail string) CodeError
	Error
}

func NewCodeError(code int, msg string) CodeError {
	return &codeError{
		code: code,
		msg:  msg,
	}
}

func NewCodeErrorf(code int, format string, a ...any) CodeError {
	return &codeError{
		code: code,
		msg:  fmt.Sprintf(format, a...),
	}
}

type codeError struct {
	code   int
	msg    string
	detail string
}

func (e *codeError) Code() int {
	return e.code
}

func (e *codeError) Msg() string {
	return e.msg
}

func (e *codeError) Detail() string {
	return e.detail
}

func (e *codeError) WithDetail(detail string) CodeError {
	var d string
	if e.detail == "" {
		d = detail
	} else {
		d = e.detail + ", " + detail
	}
	return &codeError{
		code:   e.code,
		msg:    e.msg,
		detail: d,
	}
}

func (e *codeError) Wrap() error {
	return stack.New(e, stackSkip)
}

func (e *codeError) clone() *codeError {
	return &codeError{
		code:   e.code,
		msg:    e.msg,
		detail: e.detail,
	}
}

func (e *codeError) WrapMsg(msg string, kv ...any) error {
	retErr := e.clone()
	if msg != "" || len(kv) > 0 {
		detail := toString(msg, kv)
		if retErr.detail == "" {
			retErr.detail = detail
		} else {
			retErr.detail += ", " + detail
		}
	}
	return stack.New(retErr, stackSkip)
}

func (e *codeError) Is(err error) bool {
	var codeErr CodeError
	ok := errors.As(Unwrap(err), &codeErr)
	if !ok {
		if err == nil && e == nil {
			return true
		}
		// target isn't CodeError type, tring convert it.
		for i := 0; i < len(Handlers); i++ {
			if codeErr = Handlers[i](err); codeErr != nil {
				break
			}
		}
		// convert failed.
		if codeErr == nil {
			return false
		}
	}
	if e == nil {
		return false
	}
	code := codeErr.Code()
	if e.code == code {
		return true
	}
	return DefaultCodeRelation.Is(e.code, code)
}

const initialCapacity = 3

func (e *codeError) Error() string {
	v := make([]string, 0, initialCapacity)
	v = append(v, strconv.Itoa(e.code), e.msg)

	if e.detail != "" {
		v = append(v, e.detail)
	}

	return strings.Join(v, " ")
}

func Unwrap(err error) error {
	for err != nil {
		unwrap, ok := err.(interface {
			error
			Unwrap() error
		})
		if !ok {
			break
		}
		err = unwrap.Unwrap()
		if err == nil {
			return unwrap
		}
	}
	return err
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return stack.New(err, stackSkip)
}

func WrapMsg(err error, msg string, kv ...any) error {
	if err == nil {
		return nil
	}
	err = NewErrorWrapper(err, toString(msg, kv))
	return stack.New(err, stackSkip)
}

type CodeRelation interface {
	Add(codes ...int) error
	Is(parent, child int) bool
}

func newCodeRelation() CodeRelation {
	return &codeRelation{m: make(map[int]map[int]struct{})}
}

type codeRelation struct {
	m map[int]map[int]struct{}
}

const minimumCodesLength = 2

func (r *codeRelation) Add(codes ...int) error {
	if len(codes) < minimumCodesLength {
		return New("codes length must be greater than 2", "codes", codes).Wrap()
	}
	for i := 1; i < len(codes); i++ {
		parent := codes[i-1]
		s, ok := r.m[parent]
		if !ok {
			s = make(map[int]struct{})
			r.m[parent] = s
		}
		for _, code := range codes[i:] {
			s[code] = struct{}{}
		}
	}
	return nil
}

func (r *codeRelation) Is(parent, child int) bool {
	if parent == child {
		return true
	}
	s, ok := r.m[parent]
	if !ok {
		return false
	}
	_, ok = s[child]
	return ok
}

var Handlers []func(err error) CodeError
