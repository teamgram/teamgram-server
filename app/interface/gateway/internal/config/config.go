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

package config

import (
	"fmt"
	"strings"

	bffproxyclient "github.com/teamgram/teamgram-server/v2/app/bff/bff/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type Config struct {
	kitex.RpcServerConf
	Transport                                     TransportConf
	GatewayId                                     string
	AuthsessionClient                             kitex.RpcClientConf
	PresenceClient                                kitex.RpcClientConf
	BffClient                                     bffproxyclient.BFFProxyClientListConf
	AdvertiseRpcAddr                              string
	PresenceRefreshMinIntervalSeconds             int `json:",default=20"`
	PresenceRefreshScanIntervalSeconds            int `json:",default=10"`
	PresenceRefreshRetryMinIntervalSeconds        int `json:",default=5"`
	GatewayShutdownPresenceOfflineDeadlineSeconds int `json:",default=5"`
	GatewayShutdownPresenceOfflineMaxSessions     int `json:",default=10000"`
}

type TransportConf struct {
	TCPListenOn string
	// Reserved for future HTTP/WebSocket MTProto transport. Phase 3 only serves TCP.
	HTTPListenOn string `json:",optional"`
}

func (c Config) Validate() error {
	if c.GatewayId == "" {
		return fmt.Errorf("gateway config: GatewayId is required")
	}
	if c.AdvertiseRpcAddr == "" {
		return fmt.Errorf("gateway config: AdvertiseRpcAddr is required")
	}
	if strings.HasPrefix(c.AdvertiseRpcAddr, "0.0.0.0:") {
		return fmt.Errorf("gateway config: AdvertiseRpcAddr must not use wildcard host")
	}
	if c.PresenceRefreshMinIntervalSeconds <= 0 {
		return fmt.Errorf("gateway config: PresenceRefreshMinIntervalSeconds must be positive")
	}
	if c.PresenceRefreshScanIntervalSeconds <= 0 {
		return fmt.Errorf("gateway config: PresenceRefreshScanIntervalSeconds must be positive")
	}
	if c.PresenceRefreshRetryMinIntervalSeconds <= 0 {
		return fmt.Errorf("gateway config: PresenceRefreshRetryMinIntervalSeconds must be positive")
	}
	if c.GatewayShutdownPresenceOfflineDeadlineSeconds <= 0 {
		return fmt.Errorf("gateway config: GatewayShutdownPresenceOfflineDeadlineSeconds must be positive")
	}
	if c.GatewayShutdownPresenceOfflineMaxSessions <= 0 {
		return fmt.Errorf("gateway config: GatewayShutdownPresenceOfflineMaxSessions must be positive")
	}
	return nil
}
