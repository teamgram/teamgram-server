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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserSetDefaultHistoryTTL
// user.setDefaultHistoryTTL user_id:long ttl:int = Bool;
func (c *UserCore) UserSetDefaultHistoryTTL(in *user.TLUserSetDefaultHistoryTTL) (*mtproto.Bool, error) {
	_, _, _ = c.svcCtx.Dao.DefaultHistoryTtlDAO.InsertOrUpdate(c.ctx, &dataobject.DefaultHistoryTtlDO{
		UserId: in.GetUserId(),
		Period: in.GetTtl(),
	})

	return mtproto.BoolTrue, nil
}
