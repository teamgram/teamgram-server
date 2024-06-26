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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetImmutableUserV2
// user.getImmutableUserV2 flags:# id:long privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = ImmutableUser;
// user.getImmutableUserV2 flags:# id:long privacy:flags.0?true contacts:flags.1?Vector<long> reverse_contacts:flags.2?Vector<long> = ImmutableUser;
func (c *UserCore) UserGetImmutableUserV2(in *user.TLUserGetImmutableUserV2) (*mtproto.ImmutableUser, error) {
	rV, err := c.svcCtx.Dao.GetImmutableUserV2(
		c.ctx,
		in.Id,
		in.Privacy,
		in.HasTo,
		in.To)
	if err != nil {
		c.Logger.Errorf("user.getImmutableUserV2(%s) - error: %v", in, err)
		return nil, err
	}

	return rV, nil
}
