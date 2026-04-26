package status

import "errors"

var (
	ErrStatusInvalidArgument = errors.New("status: invalid argument")
	ErrStatusStorage         = errors.New("status: storage failure")
)
