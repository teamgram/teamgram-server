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
	"github.com/teamgram/teamgram-server/app/service/idgen/idgen"
	"strconv"
)

// IdgenGetCurrentSeqIdList
// idgen.getCurrentSeqIdList id:Vector<InputId> = Vector<IdVal>;
func (c *IdgenCore) IdgenGetCurrentSeqIdList(in *idgen.TLIdgenGetCurrentSeqIdList) (*idgen.Vector_IdVal, error) {
	var (
		idList = make([]*idgen.IdVal, len(in.GetId()))
	)

	for i, id := range in.GetId() {
		switch id.GetPredicateName() {
		case idgen.Predicate_inputSeqId:
			sid, err := c.svcCtx.Dao.KV.GetCtx(c.ctx, id.Key)
			if err != nil {
				c.Logger.Errorf("idgen.getCurrentSeqIdList(%s) error: %v", id.Key, err)
				return nil, err
			}

			if sid == "" {
				idList[i] = idgen.MakeTLSeqIdVal(&idgen.IdVal{
					Id_INT64: 0,
				}).To_IdVal()
			} else {
				iV, _ := strconv.ParseInt(sid, 10, 64)
				idList[i] = idgen.MakeTLSeqIdVal(&idgen.IdVal{
					Id_INT64: iV,
				}).To_IdVal()
			}
		default:
			err := mtproto.ErrInputRequestInvalid
			c.Logger.Errorf("idgen.getCurrentSeqIdList - error: %v", err)
			return nil, err
		}
	}

	return &idgen.Vector_IdVal{
		Datas: idList,
	}, nil
}
