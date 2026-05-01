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

package svc

import (
	_ "github.com/teamgram/teamgram-server/v2/app/bff/authorization/authorization/authorizationservice"
	bffproxyclient "github.com/teamgram/teamgram-server/v2/app/bff/bff/client"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/configuration/configuration/configurationservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/qrcode/qrcode/qrcodeservice"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/repository"
)

type ServiceContext struct {
	Config config.Config
	Repo   *repository.Repository
	BFF    *bffproxyclient.BFFProxyClient2
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repository.NewRepository(c),
		BFF:    bffproxyclient.NewBFFProxyClient2(c.BffClient.Clients),
	}
}
func (s *ServiceContext) Close() error {
	if s == nil || s.Repo == nil {
		return nil
	}
	return s.Repo.Close()
}
