/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/app/service/idgen/idgen"
)

const (
	maxNextIdsNum = 100
)

// IdgenNextIds
// idgen.nextIds num:int = Vector<long>;
func (c *IdgenCore) IdgenNextIds(in *idgen.TLIdgenNextIds) (*idgen.Vector_Long, error) {
	if in.GetNum() > maxNextIdsNum || in.GetNum() < 0 {
		c.Logger.Errorf("NextIds num can't be greater than %d or less than 0", maxNextIdsNum)
		return nil, fmt.Errorf("NextIds num: %d error", in.GetNum())
	}

	ids := make([]int64, in.GetNum())
	for i := int32(0); i < in.GetNum(); i++ {
		// TODO: 库里提供ids方法，以减少Lock次数
		ids[i] = c.svcCtx.Node.Generate().Int64()
	}

	return &idgen.Vector_Long{
		Datas: ids,
	}, nil
}
