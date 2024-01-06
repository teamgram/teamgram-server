/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package idgen_client

import (
	"context"
	"strconv"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/idgen/idgen"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type IDGenClient2 struct {
	cli IdgenClient
}

func NewIDGenClient2(cli zrpc.Client) IDGenClient2 {
	return IDGenClient2{
		cli: NewIdgenClient(cli),
	}
}

func (m *IDGenClient2) NextId(ctx context.Context) (id int64) {
	rVal, err := m.cli.IdgenNextId(ctx, &idgen.TLIdgenNextId{})
	if err != nil {
		logx.WithContext(ctx).Errorf("idgen.nextId - error: %v", err)
	} else {
		id = rVal.V
	}

	return
}

func (m *IDGenClient2) NextIds(ctx context.Context, num int) (idList []int64) {
	rVal, err := m.cli.IdgenNextIds(ctx, &idgen.TLIdgenNextIds{
		Num: int32(num),
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("idgen.nextIds - error: %v", err)
	} else {
		idList = rVal.Datas
	}

	return

}

func (m *IDGenClient2) getCurrentSeqId(ctx context.Context, k string) (id int64) {
	rVal, err := m.cli.IdgenGetCurrentSeqId(ctx, &idgen.TLIdgenGetCurrentSeqId{
		Key: k,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("idgen.getCurrentSeqId - error: %v", err)
	} else {
		id = rVal.V
	}

	return
}

func (m *IDGenClient2) setCurrentSeqId(ctx context.Context, k string, v int64) (err error) {
	_, err = m.cli.IdgenSetCurrentSeqId(ctx, &idgen.TLIdgenSetCurrentSeqId{
		Key: k,
		Id:  v,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("idgen.setCurrentSeqId - error: %v", err)
	}

	return
}

func (m *IDGenClient2) getNextSeqId(ctx context.Context, k string) (id int64) {
	rVal, err := m.cli.IdgenGetNextSeqId(ctx, &idgen.TLIdgenGetNextSeqId{
		Key: k,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("idgen.setCurrentSeqId - error: %v", err)
	} else {
		id = rVal.V
	}

	return
}

func (m *IDGenClient2) getNextNSeqId(ctx context.Context, k string, n int) (id int64) {
	rVal, err := m.cli.IdgenGetNextNSeqId(ctx, &idgen.TLIdgenGetNextNSeqId{
		Key: k,
		N:   int32(n),
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("idgen.setCurrentSeqId - error: %v", err)
	} else {
		id = rVal.V
	}

	return
}

func (m *IDGenClient2) NextMessageBoxId(ctx context.Context, key int64) (seq int32) {
	return int32(m.getNextSeqId(ctx, messageBoxUpdatesNgenId+strconv.FormatInt(key, 10)))
}

func (m *IDGenClient2) CurrentMessageBoxId(ctx context.Context, key int64) (seq int64) {
	seq = m.getCurrentSeqId(ctx, messageBoxUpdatesNgenId+strconv.FormatInt(key, 10))
	return
}

func (m *IDGenClient2) SetCurrentMessageBoxId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, messageBoxUpdatesNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextMessageDataId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, messageDataNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentMessageDataId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, messageDataNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextChannelMessageBoxId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, channelMessageBoxNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) CurrentChannelMessageBoxId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getCurrentSeqId(ctx, channelMessageBoxNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentChannelMessageBoxId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, channelMessageBoxNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextSeqId(ctx context.Context, key int64) (seq int64) {
	seq = m.getNextSeqId(ctx, seqUpdatesNgenId+strconv.FormatInt(key, 10))
	return
}

func (m *IDGenClient2) CurrentSeqId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getCurrentSeqId(ctx, seqUpdatesNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentSeqId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, seqUpdatesNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextPtsId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, ptsUpdatesNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) NextNPtsId(ctx context.Context, key int64, n int) (seq int32) {
	seq = int32(m.getNextNSeqId(ctx, ptsUpdatesNgenId+strconv.FormatInt(key, 10), n))
	return
}

func (m *IDGenClient2) CurrentPtsId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getCurrentSeqId(ctx, ptsUpdatesNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentPtsId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, ptsUpdatesNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextQtsId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, qtsUpdatesNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) CurrentQtsId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getCurrentSeqId(ctx, qtsUpdatesNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentQtsId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, qtsUpdatesNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextChannelPtsId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, channelPtsUpdatesNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) NextChannelNPtsId(ctx context.Context, key int64, n int) (seq int32) {
	seq = int32(m.getNextNSeqId(ctx, channelPtsUpdatesNgenId+strconv.FormatInt(key, 10), n))
	return
}

func (m *IDGenClient2) CurrentChannelPtsId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getCurrentSeqId(ctx, channelPtsUpdatesNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentChannelPtsId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, channelPtsUpdatesNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextScheduledMessageBoxId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, scheduledMessageNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentScheduledMessageBoxId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, scheduledMessageNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextBotUpdateId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, botUpdatesNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentBotUpdateId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, botUpdatesNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextStoryId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, storyNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentStoryId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, storyNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) NextChannelStoryId(ctx context.Context, key int64) (seq int32) {
	seq = int32(m.getNextSeqId(ctx, channelStoryNgenId+strconv.FormatInt(key, 10)))
	return
}

func (m *IDGenClient2) SetCurrentChannelStoryId(ctx context.Context, key int64, v int32) {
	m.setCurrentSeqId(ctx, channelStoryNgenId+strconv.FormatInt(key, 10), int64(v))
}

func (m *IDGenClient2) GetNextIdList(ctx context.Context, idList ...IDTypeNgen) []IDValue {
	if len(idList) == 0 {
		return nil
	}

	ids := make([]*idgen.InputId, len(idList))
	for i, id := range idList {
		ids[i] = id.ToInputId()
	}
	idValList, _ := m.cli.IdgenGetNextIdValList(ctx, &idgen.TLIdgenGetNextIdValList{
		Id: ids,
	})
	if len(idValList.GetDatas()) != len(idList) {
		return nil
	}

	rIdValList := make([]IDValue, len(idList))
	for i, id := range idValList.GetDatas() {
		rIdValList[i].IDType = idList[i].IDType
		rIdValList[i].Id = id.Id_INT64
		rIdValList[i].IdN = id.Id_VECTORINT64
	}

	return rIdValList
}

func (m *IDGenClient2) GetCurrentSeqIdList(ctx context.Context, idList ...IDTypeNgen) []IDValue {
	if len(idList) == 0 {
		return nil
	}

	ids := make([]*idgen.InputId, len(idList))
	for i, id := range idList {
		ids[i] = id.ToInputId()
	}
	idValList, _ := m.cli.IdgenGetCurrentSeqIdList(ctx, &idgen.TLIdgenGetCurrentSeqIdList{
		Id: ids,
	})
	if len(idValList.GetDatas()) != len(idList) {
		return nil
	}

	rIdValList := make([]IDValue, len(idList))
	for i, id := range idValList.GetDatas() {
		rIdValList[i].IDType = idList[i].IDType
		rIdValList[i].Id = id.Id_INT64
		rIdValList[i].IdN = id.Id_VECTORINT64
	}

	return rIdValList
}
