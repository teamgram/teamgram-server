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
//
// Author: teamgramio (teamgram.io@gmail.com)

package model

import (
	"errors"
)

var ErrNotFound = errors.New("repository model: not found")

type NotFoundError struct {
	Resource string
	Key      string
	Cause    error
}

func (e *NotFoundError) Error() string {
	if e == nil {
		return ErrNotFound.Error()
	}
	if e.Cause == nil {
		return e.Resource + " not found: " + e.Key
	}
	return e.Resource + " not found: " + e.Key + ": " + e.Cause.Error()
}

func (e *NotFoundError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

func (e *NotFoundError) Is(target error) bool {
	return target == ErrNotFound
}
