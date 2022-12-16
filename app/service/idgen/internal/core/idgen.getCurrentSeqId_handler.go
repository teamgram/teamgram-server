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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/idgen/idgen"
	"strconv"
)

// IdgenGetCurrentSeqId
// idgen.getCurrentSeqId key:string = Int64;
func (c *IdgenCore) IdgenGetCurrentSeqId(in *idgen.TLIdgenGetCurrentSeqId) (*mtproto.Int64, error) {
	id, err := c.svcCtx.Dao.KV.GetCtx(c.ctx, in.GetKey())
	if err != nil {
		c.Logger.Errorf("dgen.getCurrentSeqId(%s) error: %v", in.GetKey(), err)
		return nil, err
	}

	if id == "" {
		return &mtproto.Int64{
			V: 0,
		}, nil
	}

	intValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Logger.Errorf("dgen.getCurrentSeqId(%s) error: %v", in.GetKey(), err)
		return nil, fmt.Errorf("the value %q cannot parsed as int", err)
	}

	return &mtproto.Int64{
		V: intValue,
	}, nil
}
