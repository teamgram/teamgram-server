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

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UserupdatesGetDifference
// userupdates.getDifference flags:# user_id:long auth_key_id:long pts:long pts_total_limit:flags.0?int date:flags.1?long = UserDifference;
func (c *UserupdatesCore) UserupdatesGetDifference(in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
	// TODO: not impl
	c.Logger.Errorf("userupdates.getDifference - error: method UserupdatesGetDifference not impl")

	return nil, tg.ErrMethodNotImpl
}
