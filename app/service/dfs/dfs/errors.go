package dfs

import (
	"errors"
	"fmt"
)

var (
	ErrDfsInvalidArgument    = errors.New("dfs: invalid argument")
	ErrDfsInvalidFilePart    = errors.New("dfs: invalid file part")
	ErrDfsFileNotFound       = errors.New("dfs: file not found")
	ErrDfsChecksumInvalid    = errors.New("dfs: checksum invalid")
	ErrDfsImageProcessFailed = errors.New("dfs: image process failed")
	ErrDfsVideoProcessFailed = errors.New("dfs: video process failed")
	ErrDfsStorage            = errors.New("dfs: storage failure")
	ErrDfsDownstream         = errors.New("dfs: downstream failure")
)

func WrapDfsStorage(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", ErrDfsStorage, op, err)
}

func WrapDfsDownstream(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", ErrDfsDownstream, op, err)
}
