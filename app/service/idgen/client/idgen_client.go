/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package idgenclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen/idgenservice"

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
}

type defaultIdgenClient struct {
	cli client.Client
}

func NewIdgenClient(cli client.Client) IdgenClient {
	return &defaultIdgenClient{
		cli: cli,
	}
}

// IdgenNextId
// idgen.nextId = Int64;
func (m *defaultIdgenClient) IdgenNextId(ctx context.Context, in *idgen.TLIdgenNextId) (*tg.Int64, error) {
	cli := idgenservice.NewRPCIdgenClient(m.cli)
	return cli.IdgenNextId(ctx, in)
}

// IdgenNextIds
// idgen.nextIds num:int = Vector<long>;
func (m *defaultIdgenClient) IdgenNextIds(ctx context.Context, in *idgen.TLIdgenNextIds) (*idgen.VectorLong, error) {
	cli := idgenservice.NewRPCIdgenClient(m.cli)
	return cli.IdgenNextIds(ctx, in)
}

// IdgenGetCurrentSeqId
// idgen.getCurrentSeqId key:string = Int64;
func (m *defaultIdgenClient) IdgenGetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqId) (*tg.Int64, error) {
	cli := idgenservice.NewRPCIdgenClient(m.cli)
	return cli.IdgenGetCurrentSeqId(ctx, in)
}

// IdgenSetCurrentSeqId
// idgen.setCurrentSeqId key:string id:long = Bool;
func (m *defaultIdgenClient) IdgenSetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenSetCurrentSeqId) (*tg.Bool, error) {
	cli := idgenservice.NewRPCIdgenClient(m.cli)
	return cli.IdgenSetCurrentSeqId(ctx, in)
}

// IdgenGetNextSeqId
// idgen.getNextSeqId key:string = Int64;
func (m *defaultIdgenClient) IdgenGetNextSeqId(ctx context.Context, in *idgen.TLIdgenGetNextSeqId) (*tg.Int64, error) {
	cli := idgenservice.NewRPCIdgenClient(m.cli)
	return cli.IdgenGetNextSeqId(ctx, in)
}

// IdgenGetNextNSeqId
// idgen.getNextNSeqId key:string n:int = Int64;
func (m *defaultIdgenClient) IdgenGetNextNSeqId(ctx context.Context, in *idgen.TLIdgenGetNextNSeqId) (*tg.Int64, error) {
	cli := idgenservice.NewRPCIdgenClient(m.cli)
	return cli.IdgenGetNextNSeqId(ctx, in)
}

// IdgenGetNextIdValList
// idgen.getNextIdValList id:Vector<InputId> = Vector<IdVal>;
func (m *defaultIdgenClient) IdgenGetNextIdValList(ctx context.Context, in *idgen.TLIdgenGetNextIdValList) (*idgen.VectorIdVal, error) {
	cli := idgenservice.NewRPCIdgenClient(m.cli)
	return cli.IdgenGetNextIdValList(ctx, in)
}

// IdgenGetCurrentSeqIdList
// idgen.getCurrentSeqIdList id:Vector<InputId> = Vector<IdVal>;
func (m *defaultIdgenClient) IdgenGetCurrentSeqIdList(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqIdList) (*idgen.VectorIdVal, error) {
	cli := idgenservice.NewRPCIdgenClient(m.cli)
	return cli.IdgenGetCurrentSeqIdList(ctx, in)
}
