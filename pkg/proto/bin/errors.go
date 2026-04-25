// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func NewInvalidLengthError(typ string, length int, offset int) error {
	return &InvalidLengthError{Type: typ, Length: length, Offset: offset}
}

func (i *InvalidLengthError) Error() string {
	if i.Offset >= 0 {
		return fmt.Sprintf("invalid %s length: %d at offset %d", i.Type, i.Length, i.Offset)
	}
	return fmt.Sprintf("invalid %s length: %d", i.Type, i.Length)
}

// UnexpectedClazzIDError means that unknown or unexpected type id was decoded.
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
