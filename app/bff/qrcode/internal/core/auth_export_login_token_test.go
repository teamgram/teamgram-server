// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/qrcode/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/bff/qrcode/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthExportLoginTokenReturnsLoginToken(t *testing.T) {
	c := New(context.Background(), svc.NewServiceContext(config.Config{}))

	got, err := c.AuthExportLoginToken(&tg.TLAuthExportLoginToken{
		ApiId:     1,
		ApiHash:   "test-api-hash",
		ExceptIds: []int64{1, 2},
	})
	if err != nil {
		t.Fatalf("AuthExportLoginToken error = %v", err)
	}
	token, ok := got.ToAuthLoginToken()
	if !ok {
		t.Fatalf("AuthExportLoginToken returned %s, want auth.loginToken", got.ClazzName())
	}
	if token.Expires <= 0 {
		t.Fatalf("Expires = %d, want positive", token.Expires)
	}
	if len(token.Token) == 0 {
		t.Fatal("Token is empty")
	}
	if err := got.Validate(223); err != nil {
		t.Fatalf("AuthExportLoginToken result Validate(223) error = %v", err)
	}
}
