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
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
)

var _ *tg.Bool

// IdgenGetNextIdValList
// idgen.getNextIdValList id:Vector<InputId> = Vector<IdVal>;
func (c *IdgenCore) IdgenGetNextIdValList(in *idgen.TLIdgenGetNextIdValList) (*idgen.VectorIdVal, error) {
	var (
		idList = make([]*idgen.IdVal, len(in.Id))
	)

	for i, id := range in.Id {
		id.Match(
			func(id2 *idgen.TLInputId) interface{} {
				idList[i] = idgen.MakeIdVal(&idgen.TLIdVal{
					Id_INT64: c.svcCtx.Dao.Node.Generate().Int64(),
				})

				return nil
			},
			func(id2 *idgen.TLInputIds) interface{} {
				ids := make([]int64, id2.Num)
				for j := int32(0); j < id2.Num; j++ {
					// TODO: 库里提供ids方法，以减少Lock次数
					ids[j] = c.svcCtx.Node.Generate().Int64()
				}
				idList[i] = idgen.MakeIdVal(&idgen.TLIdVals{
					Id_VECTORINT64: ids,
				})

				return nil
			},
			func(id2 *idgen.TLInputNSeqId) interface{} {
				sid, err := c.svcCtx.Dao.KV.IncrbyCtx(c.ctx, id2.Key, 1)
				if err != nil {
					c.Logger.Errorf("dgen.getNextIdValList(%s) error: %v", id2.Key, err)
					return err
				}
				idList[i] = idgen.MakeIdVal(&idgen.TLSeqIdVal{
					Id_INT64: sid,
				})

				return nil
			},
			func(id2 *idgen.TLInputNSeqId) interface{} {
				sid, err := c.svcCtx.Dao.KV.IncrbyCtx(c.ctx, id2.Key, int64(id2.N))
				if err != nil {
					c.Logger.Errorf("dgen.getNextIdValList(%s, %d) error: %v", id2.Key, id2.N, err)
					return err
				}
				idList[i] = idgen.MakeIdVal(&idgen.TLSeqIdVal{
					Id_INT64: sid,
				})

				return nil
			})
	}

	return &idgen.VectorIdVal{
		Datas: idList,
	}, nil
}
