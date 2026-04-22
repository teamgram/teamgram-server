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

package tg

import (
	"testing"
)

func TestNewRpcError(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		e := NewRpcError(nil)
		if e == nil {
			t.Fatal("expected rpc error, got nil")
		}
		if e.ErrorCode != ErrInternal {
			t.Fatalf("expected error code %d, got %d", ErrInternal, e.ErrorCode)
		}
		if e.ErrorMessage != "INTERNAL_SERVER_ERROR" {
			t.Fatalf("expected INTERNAL_SERVER_ERROR, got %q", e.ErrorMessage)
		}
	})

	t.Run("plain error", func(t *testing.T) {
		e := NewRpcError(assertErr("test error"))
		if e == nil {
			t.Fatal("expected rpc error, got nil")
		}
		if e.ErrorCode != ErrInternal {
			t.Fatalf("expected error code %d, got %d", ErrInternal, e.ErrorCode)
		}
		if e.ErrorMessage != "test error" {
			t.Fatalf("expected test error, got %q", e.ErrorMessage)
		}
	})

	t.Run("code error", func(t *testing.T) {
		e := NewRpcError(ErrInternalServerError)
		if e == nil {
			t.Fatal("expected rpc error, got nil")
		}
		if e.ErrorCode != ErrInternal {
			t.Fatalf("expected error code %d, got %d", ErrInternal, e.ErrorCode)
		}
		if e.ErrorMessage != "INTERNAL_SERVER_ERROR" {
			t.Fatalf("expected INTERNAL_SERVER_ERROR, got %q", e.ErrorMessage)
		}
	})

	t.Run("rpc error passthrough", func(t *testing.T) {
		src := &TLRpcError{
			ErrorCode:    ErrInternal,
			ErrorMessage: "INTERNAL_SERVER_ERROR",
		}
		e := NewRpcError(src)
		if e == nil {
			t.Fatal("expected rpc error, got nil")
		}
		if e.ErrorCode != src.ErrorCode || e.ErrorMessage != src.ErrorMessage {
			t.Fatalf("expected %+v, got %+v", src, e)
		}
	})
}

type assertErr string

func (e assertErr) Error() string {
	return string(e)
}
