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
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	startupThisDC        int32 = 1
	startupCountry             = "CN"
	startupConfigTTL           = int32(3600)
	startupAppConfigHash int32 = 0x54474143
	startupCountriesHash int32 = 0x5447434c
)

func makeStartupConfig(now int32) *tg.Config {
	return tg.MakeTLConfig(&tg.TLConfig{
		Date:                    now,
		Expires:                 now + startupConfigTTL,
		TestMode:                tg.BoolFalseClazz,
		ThisDc:                  startupThisDC,
		DcOptions:               makeStartupDcOptions(),
		DcTxtDomainName:         "localhost",
		ChatSizeMax:             200,
		MegagroupSizeMax:        10000,
		ForwardedCountMax:       100,
		OnlineUpdatePeriodMs:    210000,
		OfflineBlurTimeoutMs:    5000,
		OfflineIdleTimeoutMs:    30000,
		OnlineCloudTimeoutMs:    300000,
		NotifyCloudDelayMs:      30000,
		NotifyDefaultDelayMs:    1500,
		PushChatPeriodMs:        60000,
		PushChatLimit:           2,
		EditTimeLimit:           172800,
		RevokeTimeLimit:         172800,
		RevokePmTimeLimit:       172800,
		RatingEDecay:            2419200,
		StickersRecentLimit:     200,
		ChannelsReadMediaPeriod: 604800,
		CallReceiveTimeoutMs:    20000,
		CallRingTimeoutMs:       90000,
		CallConnectTimeoutMs:    30000,
		CallPacketTimeoutMs:     10000,
		MeUrlPrefix:             "https://t.me/",
		CaptionLengthMax:        1024,
		MessageLengthMax:        4096,
		WebfileDcId:             startupThisDC,
	}).ToConfig()
}

func makeStartupDcOptions() []tg.DcOptionClazz {
	return []tg.DcOptionClazz{
		tg.MakeTLDcOption(&tg.TLDcOption{
			Id:        startupThisDC,
			IpAddress: "127.0.0.1",
			Port:      443,
			Static:    true,
		}).ToDcOption(),
	}
}

func startupNow() int32 {
	return int32(time.Now().Unix())
}
