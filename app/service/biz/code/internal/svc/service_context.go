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

package svc

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/internal/repository"
)

// Repo is the interface CodeCore expects from the repository layer.
type Repo interface {
	GetCachePhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error)
	PutCachePhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error
	DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phone string) error
}

type ServiceContext struct {
	Config config.Config
	Repo   Repo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repository.NewRepository(c),
	}
}
