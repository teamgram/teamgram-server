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
	"errors"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
)

var _ *tg.Bool

// UserDeletePeerSettings
// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;
func (c *UserCore) UserDeletePeerSettings(in *user.TLUserDeletePeerSettings) (*tg.Bool, error) {
	// TODO: not impl
	// c.Logger.Errorf("user.deletePeerSettings blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, errors.New("user.deletePeerSettings not implemented")
}
