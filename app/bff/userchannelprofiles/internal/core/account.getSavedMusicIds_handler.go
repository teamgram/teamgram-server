// Copyright 2025 Teamgram Authors
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

// AccountGetSavedMusicIds
// account.getSavedMusicIds#e09d5faf hash:long = account.SavedMusicIds;
func (c *UserChannelProfilesCore) AccountGetSavedMusicIds(in *mtproto.TLAccountGetSavedMusicIds) (*mtproto.Account_SavedMusicIds, error) {
	// TODO: not impl

	rV := mtproto.MakeTLAccountSavedMusicIds(&mtproto.Account_SavedMusicIds{
		Ids: []int64{},
	}).To_Account_SavedMusicIds()

	return rV, nil
}
