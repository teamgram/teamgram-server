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
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
)

// IdgenGetCurrentSeqIdList
// idgen.getCurrentSeqIdList id:Vector<InputId> = Vector<IdVal>;
func (c *IdgenCore) IdgenGetCurrentSeqIdList(in *idgen.TLIdgenGetCurrentSeqIdList) (*idgen.VectorIdVal, error) {
	idList := make([]idgen.IdValClazz, len(in.Id))
	for i, input := range in.Id {
		id, ok := input.(*idgen.TLInputSeqId)
		if !ok {
			return nil, fmt.Errorf("%w: invalid current seq input id at index %d", idgen.ErrInvalidArgument, i)
		}
		seq, err := c.svcCtx.Repo.QueryCurrentSeqID(c.ctx, id.Key)
		if err != nil {
			return nil, err
		}
		idList[i] = idgen.MakeTLSeqIdVal(&idgen.TLSeqIdVal{Id: seq})
	}

	return &idgen.VectorIdVal{Datas: idList}, nil
}
