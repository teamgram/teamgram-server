package idgen

import "errors"

var (
	ErrInvalidArgument         = errors.New("idgen: invalid argument")
	ErrSeqAllocatorUnavailable = errors.New("idgen: seq allocator unavailable")
	ErrSeqStorage              = errors.New("idgen: seq storage failure")
)
