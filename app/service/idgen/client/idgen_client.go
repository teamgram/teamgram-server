/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package idgen_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/idgen/idgen"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type IdgenClient interface {
	IdgenNextId(ctx context.Context, in *idgen.TLIdgenNextId) (*mtproto.Int64, error)
	IdgenNextIds(ctx context.Context, in *idgen.TLIdgenNextIds) (*idgen.Vector_Long, error)
	IdgenGetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqId) (*mtproto.Int64, error)
	IdgenSetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenSetCurrentSeqId) (*mtproto.Bool, error)
	IdgenGetNextSeqId(ctx context.Context, in *idgen.TLIdgenGetNextSeqId) (*mtproto.Int64, error)
	IdgenGetNextNSeqId(ctx context.Context, in *idgen.TLIdgenGetNextNSeqId) (*mtproto.Int64, error)
	IdgenGetNextIdValList(ctx context.Context, in *idgen.TLIdgenGetNextIdValList) (*idgen.Vector_IdVal, error)
	IdgenGetCurrentSeqIdList(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqIdList) (*idgen.Vector_IdVal, error)
}

type defaultIdgenClient struct {
	cli zrpc.Client
}

func NewIdgenClient(cli zrpc.Client) IdgenClient {
	return &defaultIdgenClient{
		cli: cli,
	}
}

// IdgenNextId
// idgen.nextId = Int64;
func (m *defaultIdgenClient) IdgenNextId(ctx context.Context, in *idgen.TLIdgenNextId) (*mtproto.Int64, error) {
	client := idgen.NewRPCIdgenClient(m.cli.Conn())
	return client.IdgenNextId(ctx, in)
}

// IdgenNextIds
// idgen.nextIds num:int = Vector<long>;
func (m *defaultIdgenClient) IdgenNextIds(ctx context.Context, in *idgen.TLIdgenNextIds) (*idgen.Vector_Long, error) {
	client := idgen.NewRPCIdgenClient(m.cli.Conn())
	return client.IdgenNextIds(ctx, in)
}

// IdgenGetCurrentSeqId
// idgen.getCurrentSeqId key:string = Int64;
func (m *defaultIdgenClient) IdgenGetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqId) (*mtproto.Int64, error) {
	client := idgen.NewRPCIdgenClient(m.cli.Conn())
	return client.IdgenGetCurrentSeqId(ctx, in)
}

// IdgenSetCurrentSeqId
// idgen.setCurrentSeqId key:string id:long = Bool;
func (m *defaultIdgenClient) IdgenSetCurrentSeqId(ctx context.Context, in *idgen.TLIdgenSetCurrentSeqId) (*mtproto.Bool, error) {
	client := idgen.NewRPCIdgenClient(m.cli.Conn())
	return client.IdgenSetCurrentSeqId(ctx, in)
}

// IdgenGetNextSeqId
// idgen.getNextSeqId key:string = Int64;
func (m *defaultIdgenClient) IdgenGetNextSeqId(ctx context.Context, in *idgen.TLIdgenGetNextSeqId) (*mtproto.Int64, error) {
	client := idgen.NewRPCIdgenClient(m.cli.Conn())
	return client.IdgenGetNextSeqId(ctx, in)
}

// IdgenGetNextNSeqId
// idgen.getNextNSeqId key:string n:int = Int64;
func (m *defaultIdgenClient) IdgenGetNextNSeqId(ctx context.Context, in *idgen.TLIdgenGetNextNSeqId) (*mtproto.Int64, error) {
	client := idgen.NewRPCIdgenClient(m.cli.Conn())
	return client.IdgenGetNextNSeqId(ctx, in)
}

// IdgenGetNextIdValList
// idgen.getNextIdValList id:Vector<InputId> = Vector<IdVal>;
func (m *defaultIdgenClient) IdgenGetNextIdValList(ctx context.Context, in *idgen.TLIdgenGetNextIdValList) (*idgen.Vector_IdVal, error) {
	client := idgen.NewRPCIdgenClient(m.cli.Conn())
	return client.IdgenGetNextIdValList(ctx, in)
}

// IdgenGetCurrentSeqIdList
// idgen.getCurrentSeqIdList id:Vector<InputId> = Vector<IdVal>;
func (m *defaultIdgenClient) IdgenGetCurrentSeqIdList(ctx context.Context, in *idgen.TLIdgenGetCurrentSeqIdList) (*idgen.Vector_IdVal, error) {
	client := idgen.NewRPCIdgenClient(m.cli.Conn())
	return client.IdgenGetCurrentSeqIdList(ctx, in)
}
