// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetAutoDownloadSettings
// account.getAutoDownloadSettings#56da0b3f = account.AutoDownloadSettings;
func (c *AutoDownloadCore) AccountGetAutoDownloadSettings(in *mtproto.TLAccountGetAutoDownloadSettings) (*mtproto.Account_AutoDownloadSettings, error) {
	makeAutoDownloadSettings := func(disabled, videoPreloadLarge, audioPreloadNext, phonecallsLessData bool, photoSizeMax, videoSizeMax, fileSizeMax int32) *mtproto.AutoDownloadSettings {
		return mtproto.MakeTLAutoDownloadSettings(&mtproto.AutoDownloadSettings{
			Disabled:           disabled,
			VideoPreloadLarge:  videoPreloadLarge,
			AudioPreloadNext:   audioPreloadNext,
			PhonecallsLessData: phonecallsLessData,
			PhotoSizeMax:       photoSizeMax,
			VideoSizeMax_INT32: videoSizeMax,
			VideoSizeMax_INT64: int64(videoSizeMax),
			FileSizeMax_INT32:  fileSizeMax,
			FileSizeMax_INT64:  int64(fileSizeMax),
		}).To_AutoDownloadSettings()
	}

	return mtproto.MakeTLAccountAutoDownloadSettings(&mtproto.Account_AutoDownloadSettings{
		Low: makeAutoDownloadSettings(
			false,
			true,
			true,
			true,
			1048576,
			512000,
			512000),
		Medium: makeAutoDownloadSettings(
			false,
			true,
			true,
			false,
			1048576,
			10485760,
			1048576),
		High: makeAutoDownloadSettings(
			false,
			true,
			true,
			false,
			1048576,
			15728640,
			3145728),
	}).To_Account_AutoDownloadSettings(), nil
}
