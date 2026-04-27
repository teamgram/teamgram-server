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

// IdgenGetNextIdValList
// idgen.getNextIdValList id:Vector<InputId> = Vector<IdVal>;
func (c *IdgenCore) IdgenGetNextIdValList(in *idgen.TLIdgenGetNextIdValList) (*idgen.VectorIdVal, error) {
	idList := make([]idgen.IdValClazz, len(in.Id))
	for i, input := range in.Id {
		switch id := input.(type) {
		case *idgen.TLInputId:
			idList[i] = idgen.MakeTLIdVal(&idgen.TLIdVal{Id: c.nextID()})

		case *idgen.TLInputIds:
			ids, err := c.nextIDs(id.Num)
			if err != nil {
				return nil, err
			}
			idList[i] = idgen.MakeTLIdVals(&idgen.TLIdVals{Id: ids})

		case *idgen.TLInputSeqId:
			seq, err := c.getNextSeqID(id.Key, 1)
			if err != nil {
				return nil, err
			}
			idList[i] = idgen.MakeTLSeqIdVal(&idgen.TLSeqIdVal{Id: seq})

		case *idgen.TLInputNSeqId:
			seq, err := c.getNextSeqID(id.Key, id.N)
			if err != nil {
				return nil, err
			}
			idList[i] = idgen.MakeTLSeqIdVal(&idgen.TLSeqIdVal{Id: seq})

		default:
			return nil, fmt.Errorf("%w: invalid input id at index %d", idgen.ErrInvalidArgument, i)
		}
	}

	return &idgen.VectorIdVal{Datas: idList}, nil
}
