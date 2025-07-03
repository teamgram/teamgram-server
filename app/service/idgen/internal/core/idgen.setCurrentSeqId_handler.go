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
	"strconv"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
)

var _ *tg.Bool

// IdgenSetCurrentSeqId
// idgen.setCurrentSeqId key:string id:long = Bool;
func (c *IdgenCore) IdgenSetCurrentSeqId(in *idgen.TLIdgenSetCurrentSeqId) (*tg.Bool, error) {
	err := c.svcCtx.Dao.KV.SetCtx(c.ctx, in.Key, strconv.FormatInt(in.Id, 10))
	if err != nil {
		c.Logger.Errorf("idgen.setCurrentSeqId(%s, %d) error: %v", in.Key, in.Id, err)
		return nil, err
	}

	return tg.BoolTrue, nil
}
