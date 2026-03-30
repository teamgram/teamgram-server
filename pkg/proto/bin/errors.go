// Copyright 2024 Teamgooo Authors
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
	Type   string
	Length int
	Offset int
}

func (i *InvalidLengthError) Error() string {
	if i.Offset >= 0 {
		return fmt.Sprintf("invalid %s length: %d at offset %d", i.Type, i.Length, i.Offset)
	}
	return fmt.Sprintf("invalid %s length: %d", i.Type, i.Length)
}

// UnexpectedClazzIDErr means that unknown or unexpected type id was decoded.
type UnexpectedClazzIDError struct {
	Want   uint32
	Got    uint32
	Offset int
}

func (e *UnexpectedClazzIDError) Error() string {
	if e.Offset >= 0 {
		return fmt.Sprintf("unexpected clazzID got %#x want %#x at offset %d", e.Got, e.Want, e.Offset)
	}
	return fmt.Sprintf("unexpected clazzID got %#x want %#x", e.Got, e.Want)
}

// NewUnexpectedClazzID return new UnexpectedClazzIDErr.
func NewUnexpectedClazzID(want uint32, got uint32, offset int) error {
	return &UnexpectedClazzIDError{Want: want, Got: got, Offset: offset}
}
