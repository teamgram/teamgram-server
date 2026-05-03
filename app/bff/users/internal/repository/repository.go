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

package repository

import (
	"github.com/teamgram/teamgram-server/v2/app/bff/users/internal/config"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

// Repository is the dependency container for BFF users business logic.
type Repository struct {
	UserClient userclient.UserClient
}

// NewRepository creates a new Repository. Clients are created only when their
// config section is populated, which keeps unit tests lightweight.
func NewRepository(c config.Config) *Repository {
	r := &Repository{}
	if hasRPCClientConfig(c.UserClient) {
		r.UserClient = userclient.NewUserClient(userclient.MustNewKitexClient(c.UserClient))
	}
	return r
}

func hasRPCClientConfig(c kitex.RpcClientConf) bool {
	return len(c.Endpoints) > 0 || len(c.Target) > 0 || c.HasEtcd()
}
