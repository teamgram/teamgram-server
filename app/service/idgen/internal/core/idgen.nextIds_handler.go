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
	"fmt"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
)

var _ *tg.Bool

const (
	maxNextIdsNum = 100
)

// IdgenNextIds
// idgen.nextIds num:int = Vector<long>;
func (c *IdgenCore) IdgenNextIds(in *idgen.TLIdgenNextIds) (*idgen.VectorLong, error) {
	if in.Num > maxNextIdsNum || in.Num < 0 {
		c.Logger.Errorf("NextIds num can't be greater than %d or less than 0", maxNextIdsNum)
		return nil, fmt.Errorf("NextIds num: %d error", in.Num)
	}

	ids := make([]int64, in.Num)
	for i := int32(0); i < in.Num; i++ {
		// TODO: 库里提供ids方法，以减少Lock次数
		ids[i] = c.svcCtx.Node.Generate().Int64()
	}

	return &idgen.VectorLong{
		Datas: ids,
	}, nil
}
