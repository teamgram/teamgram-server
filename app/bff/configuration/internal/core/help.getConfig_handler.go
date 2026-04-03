// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
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
	config = fallbackConfig()
)

func loadConfig(path string) *tg.TLConfig {
	configData, err := os.ReadFile(path)
	if err != nil {
		return fallbackConfig()
	}

	config2, err := iface.DecodeObject(bin.NewDecoder(configData))
	if err != nil {
		return fallbackConfig()
	}

	config, _ := config2.(*tg.TLConfig)
	if config == nil {
		return fallbackConfig()
	}

	return config
}

func fallbackConfig() *tg.TLConfig {
	return tg.MakeTLConfig(&tg.TLConfig{
		TestMode:         tg.BoolFalseClazz,
		ThisDc:           1,
		DcOptions:        []tg.DcOptionClazz{},
		DcTxtDomainName:  "apv3.stel.com",
		ChatSizeMax:      200,
		MegagroupSizeMax: 10000,
		MessageLengthMax: 4096,
		CaptionLengthMax: 1024,
		MeUrlPrefix:      "https://t.me/",
	})
}

// HelpGetConfig
// help.getConfig#c4f9186b = Config;
func (c *ConfigurationCore) HelpGetConfig(in *tg.TLHelpGetConfig) (*tg.Config, error) {
	_ = in

	c2 := loadConfig(configFile)

	now := int32(time.Now().Unix())
	c2.Date = now
	c2.Expires = now + expiresTimeout

	return c2.ToConfig(), nil
}
