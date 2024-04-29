// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package core

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/interface/httpserver/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HttpserverCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD    *metadata.RpcMetadata
	Token string
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *HttpserverCore {
	return &HttpserverCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func New2(ctx context.Context, svcCtx *svc.ServiceContext, md *metadata.RpcMetadata) *HttpserverCore {
	return &HttpserverCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     md,
	}
}

func parseFromIncomingMessage(b []byte) (msgId int64, obj mtproto.TLObject, err error) {
	dBuf := mtproto.NewDecodeBuf(b)

	msgId = dBuf.Long()
	_ = dBuf.Int()
	obj = dBuf.Object()
	err = dBuf.GetError()

	return
}

func serializeToBuffer(x *mtproto.EncodeBuf, msgId int64, obj mtproto.TLObject) error {
	//obj.Encode(x, 0)
	// x := mtproto.NewEncodeBuf(8 + 4 + len(oBuf))
	x.Long(0)
	x.Long(msgId)
	offset := x.GetOffset()
	x.Int(0)
	err := obj.Encode(x, 0)
	if err != nil {
		return err
	}
	//x.Bytes(oBuf)
	x.IntOffset(offset, int32(x.GetOffset()-offset-4))
	return nil
}
