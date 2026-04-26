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
	geoipclient "github.com/teamgram/teamgram-server/v2/app/infra/geoip/client"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/xkv"
)

// Type aliases keep the dependency types reachable from a single import.
type (
	AuthKeysModelType         = model.AuthKeysModel
	AuthUsersModelType        = model.AuthUsersModel
	AuthsModelType            = model.AuthsModel
	FutureSaltsModelType      = xkv.FutureSaltsModel
	AuthKeyLifecycleModelType = xkv.AuthKeyLifecycleModel
	GeoipClientType           = geoipclient.GeoipClient
)
