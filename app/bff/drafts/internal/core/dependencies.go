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

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *DraftsCore) repository() (*repository.Repository, error) {
	if c.svcCtx == nil || c.svcCtx.Repo == nil {
		return nil, tg.ErrInternalServerError
	}
	return c.svcCtx.Repo, nil
}

func (c *DraftsCore) dialogClient() (repository.DialogClient, error) {
	repo, err := c.repository()
	if err != nil {
		return nil, err
	}
	if repo.DialogClient == nil {
		return nil, tg.ErrInternalServerError
	}
	return repo.DialogClient, nil
}
