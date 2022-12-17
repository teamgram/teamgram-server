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
)

// IdgenGetNextIdValList
// idgen.getNextIdValList id:Vector<InputId> = Vector<IdVal>;
func (c *IdgenCore) IdgenGetNextIdValList(in *idgen.TLIdgenGetNextIdValList) (*idgen.Vector_IdVal, error) {
	var (
		idList = make([]*idgen.IdVal, len(in.GetId()))
	)

	for i, id := range in.GetId() {
		switch id.GetPredicateName() {
		case idgen.Predicate_inputId:
			idList[i] = idgen.MakeTLIdVal(&idgen.IdVal{
				Id_INT64: c.svcCtx.Dao.Node.Generate().Int64(),
			}).To_IdVal()
		case idgen.Predicate_inputIds:
			ids := make([]int64, id.Num)
			for j := int32(0); j < id.Num; j++ {
				// TODO: 库里提供ids方法，以减少Lock次数
				ids[j] = c.svcCtx.Node.Generate().Int64()
			}
			idList[i] = idgen.MakeTLIdVals(&idgen.IdVal{
				Id_VECTORINT64: ids,
			}).To_IdVal()
		case idgen.Predicate_inputSeqId:
			sid, err := c.svcCtx.Dao.KV.IncrbyCtx(c.ctx, id.Key, 1)
			if err != nil {
				c.Logger.Errorf("dgen.getNextIdValList(%s) error: %v", id.Key, err)
				return nil, err
			}
			idList[i] = idgen.MakeTLSeqIdVal(&idgen.IdVal{
				Id_INT64: sid,
			}).To_IdVal()
		case idgen.Predicate_inputNSeqId:
			sid, err := c.svcCtx.Dao.KV.IncrbyCtx(c.ctx, id.Key, int64(id.N))
			if err != nil {
				c.Logger.Errorf("dgen.getNextIdValList(%s, %d) error: %v", id.Key, id.N, err)
				return nil, err
			}
			idList[i] = idgen.MakeTLSeqIdVal(&idgen.IdVal{
				Id_INT64: sid,
			}).To_IdVal()
		default:
			err := mtproto.ErrInputRequestInvalid
			c.Logger.Errorf("idgen.getNextIdValList - error: %v", err)
			return nil, err
		}
	}

	return &idgen.Vector_IdVal{
		Datas: idList,
	}, nil
}
