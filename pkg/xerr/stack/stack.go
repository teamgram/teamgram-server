package stack

import (
	"errors"
	"path"
	"runtime"
	"strconv"
	"strings"
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
	if len(e.stack) == 0 {
		return e.err.Error()
	}
	var sb strings.Builder
	sb.WriteString("Error: ")
	sb.WriteString(e.err.Error())
	sb.WriteString(" |")
	for _, pc := range e.stack {
		fn := runtime.FuncForPC(pc - 1)
		if fn == nil {
			continue
		}
		name := path.Base(fn.Name())
		if strings.HasPrefix(name, "runtime.") {
			break
		}
		file, line := fn.FileLine(pc)
		sb.WriteString(" -> ")
		sb.WriteString(name)
		sb.WriteString("() ")
		sb.WriteString(file)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(line))
	}
	return sb.String()
}

func (e *stackError) String() string {
	return e.Error()
}
