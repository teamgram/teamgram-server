// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
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
	"os"
	"time"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

const (
	configFile = "./config.data"
	// date = 1509066502,    2017/10/27 09:08:22
	// expires = 1509070295, 2017/10/27 10:11:35
	expiresTimeout = 3600 // 超时时间设置为3600秒

	// support user: @benqi
	// SUPPORT_USER_ID = 2
)

var (
	config *tg.TLConfig
)

func init() {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	config2, err := iface.DecodeObject(bin.NewDecoder(configData))
	if err != nil {
		panic(err)
	}

	config, _ = config2.(*tg.TLConfig)
	if config == nil {
		panic("config is nil")
	}
}

// HelpGetConfig
// help.getConfig#c4f9186b = Config;
func (c *ConfigurationCore) HelpGetConfig(in *tg.TLHelpGetConfig) (*tg.Config, error) {
	_ = in

	c2 := &tg.TLConfig{}
	*c2 = *config

	now := int32(time.Now().Unix())
	c2.Date = now
	c2.Expires = now + expiresTimeout

	return c2.ToConfig(), nil
}
