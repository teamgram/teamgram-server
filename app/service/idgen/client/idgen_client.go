/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package idgenclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen/idgenservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type IdgenClient interface {
	IdgenNextId(ctx context.Context, in *idgen.TLIdgenNextId) (*tg.Int64, error)
	IdgenNextIds(ctx context.Context, in *idgen.TLIdgenNextIds) (*idgen.VectorLong, error)
	IdgenGetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqId) (*tg.Int64, error)
	IdgenSetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenSetCurrentSeqId) (*tg.Bool, error)
	IdgenGetNextSeqId(ctx context.Context, in *idgen.TLIdgenGetNextSeqId) (*tg.Int64, error)
	IdgenGetNextNSeqId(ctx context.Context, in *idgen.TLIdgenGetNextNSeqId) (*tg.Int64, error)
	IdgenGetNextIdValList(ctx context.Context, in *idgen.TLIdgenGetNextIdValList) (*idgen.VectorIdVal, error)
	IdgenGetCurrentSeqIdList(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqIdList) (*idgen.VectorIdVal, error)
	Close() error
}

type defaultIdgenClient struct {
	cli client.Client
	rpc idgenservice.Client
}

func NewIdgenClient(cli client.Client) IdgenClient {
	return &defaultIdgenClient{
		cli: cli,
		rpc: idgenservice.NewRPCIdgenClient(cli),
	}
}

func (m *defaultIdgenClient) Close() error {
	if closer, ok := any(m.cli).(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

// IdgenNextId
// idgen.nextId = Int64;
func (m *defaultIdgenClient) IdgenNextId(ctx context.Context, in *idgen.TLIdgenNextId) (*tg.Int64, error) {
	return m.rpc.IdgenNextId(ctx, in)
}

// IdgenNextIds
// idgen.nextIds num:int = Vector<long>;
func (m *defaultIdgenClient) IdgenNextIds(ctx context.Context, in *idgen.TLIdgenNextIds) (*idgen.VectorLong, error) {
	return m.rpc.IdgenNextIds(ctx, in)
}

// IdgenGetCurrentSeqId
// idgen.getCurrentSeqId key:string = Int64;
func (m *defaultIdgenClient) IdgenGetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqId) (*tg.Int64, error) {
	return m.rpc.IdgenGetCurrentSeqId(ctx, in)
}

// IdgenSetCurrentSeqId
// idgen.setCurrentSeqId key:string id:long = Bool;
func (m *defaultIdgenClient) IdgenSetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenSetCurrentSeqId) (*tg.Bool, error) {
	return m.rpc.IdgenSetCurrentSeqId(ctx, in)
}

// IdgenGetNextSeqId
// idgen.getNextSeqId key:string = Int64;
func (m *defaultIdgenClient) IdgenGetNextSeqId(ctx context.Context, in *idgen.TLIdgenGetNextSeqId) (*tg.Int64, error) {
	return m.rpc.IdgenGetNextSeqId(ctx, in)
}

// IdgenGetNextNSeqId
// idgen.getNextNSeqId key:string n:int = Int64;
func (m *defaultIdgenClient) IdgenGetNextNSeqId(ctx context.Context, in *idgen.TLIdgenGetNextNSeqId) (*tg.Int64, error) {
	return m.rpc.IdgenGetNextNSeqId(ctx, in)
}

// IdgenGetNextIdValList
// idgen.getNextIdValList id:Vector<InputId> = Vector<IdVal>;
func (m *defaultIdgenClient) IdgenGetNextIdValList(ctx context.Context, in *idgen.TLIdgenGetNextIdValList) (*idgen.VectorIdVal, error) {
	return m.rpc.IdgenGetNextIdValList(ctx, in)
}

// IdgenGetCurrentSeqIdList
// idgen.getCurrentSeqIdList id:Vector<InputId> = Vector<IdVal>;
func (m *defaultIdgenClient) IdgenGetCurrentSeqIdList(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqIdList) (*idgen.VectorIdVal, error) {
	return m.rpc.IdgenGetCurrentSeqIdList(ctx, in)
}
