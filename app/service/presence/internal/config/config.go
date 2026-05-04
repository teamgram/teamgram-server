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

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type Config struct {
	kitex.RpcServerConf
	KV kv.KvConf

	SessionExpiresSeconds                  int  `json:",default=60"`
	HashKeyTTLSeconds                      int  `json:",default=600"`
	CleanupOnWriteIntervalSeconds          int  `json:",default=60"`
	PresenceQueryDefaultQPSPerCaller       int  `json:",default=50"`
	PresenceGatewayDiagnosticsQPSPerCaller int  `json:",default=1"`
	RequireCallerIdentity                  bool `json:",default=false"`
	GatewayCallers                         []string
	SyncCallers                            []string
	AdminCallers                           []string
	DebugCallers                           []string
}

func (c *Config) Validate() error {
	if len(c.KV) == 0 {
		return fmt.Errorf("KV must contain at least one node")
	}
	totalWeight := 0
	for i, node := range c.KV {
		if node.Host == "" {
			return fmt.Errorf("KV[%d].Host must not be empty", i)
		}
		totalWeight += node.Weight
	}
	if totalWeight <= 0 {
		return fmt.Errorf("KV total node weight must be positive, got %d", totalWeight)
	}
	if c.SessionExpiresSeconds <= 0 {
		return fmt.Errorf("SessionExpiresSeconds must be positive, got %d", c.SessionExpiresSeconds)
	}
	if c.HashKeyTTLSeconds < c.SessionExpiresSeconds {
		return fmt.Errorf("HashKeyTTLSeconds must be >= SessionExpiresSeconds, got %d < %d", c.HashKeyTTLSeconds, c.SessionExpiresSeconds)
	}
	if c.CleanupOnWriteIntervalSeconds <= 0 {
		return fmt.Errorf("CleanupOnWriteIntervalSeconds must be positive, got %d", c.CleanupOnWriteIntervalSeconds)
	}
	if c.PresenceQueryDefaultQPSPerCaller <= 0 {
		return fmt.Errorf("PresenceQueryDefaultQPSPerCaller must be positive, got %d", c.PresenceQueryDefaultQPSPerCaller)
	}
	if c.PresenceGatewayDiagnosticsQPSPerCaller <= 0 {
		return fmt.Errorf("PresenceGatewayDiagnosticsQPSPerCaller must be positive, got %d", c.PresenceGatewayDiagnosticsQPSPerCaller)
	}
	return nil
}
