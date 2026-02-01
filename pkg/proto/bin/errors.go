// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bin

import (
	"fmt"
)

// InvalidLengthError is returned when decoder reads invalid length.
type InvalidLengthError struct {
	Length int
	Where  string
}

func (i *InvalidLengthError) Error() string {
	return fmt.Sprintf("invalid %s length: %d", i.Where, i.Length)
}

// UnexpectedClazzIDErr means that unknown or unexpected type id was decoded.
type UnexpectedClazzIDErr struct {
	ClazzID uint32
}

func (e *UnexpectedClazzIDErr) Error() string {
	return fmt.Sprintf("unexpected clazzID %#x", uint32(e.ClazzID))
}

// NewUnexpectedClazzID return new UnexpectedClazzIDErr.
func NewUnexpectedClazzID(id uint32) error {
	return &UnexpectedClazzIDErr{ClazzID: id}
}
