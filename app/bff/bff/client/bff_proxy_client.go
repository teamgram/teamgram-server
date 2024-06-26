// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package bff_proxy_client

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type BFFProxyClient struct {
	// zrpc.Client
	BFFClients map[string]zrpc.Client
}

func NewBFFProxyClients(cList []zrpc.RpcClientConf, idMap map[string]string) *BFFProxyClient {
	var (
		clients   = make(map[string]zrpc.Client)
		registers = mtproto.GetRPCContextRegisters()
	)

	for _, c := range cList {
		cli := rpcx.GetCachedRpcClient(c)
		for k, v := range idMap {
			if v == c.Etcd.Key {
				clients[k] = cli
			}
		}
	}

	bizClients := make(map[string]zrpc.Client)
	for m, ctx := range registers {
		for k, _ := range idMap {
			if strings.HasPrefix(ctx.Method, k) {
				bizClients[m] = clients[k]
				break
			}
		}
	}

	return &BFFProxyClient{
		BFFClients: bizClients,
	}
}

func (c *BFFProxyClient) GetRpcClientByRequest(t interface{}) (zrpc.Client, error) {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if c2, ok := c.BFFClients[rt.Name()]; ok {
		return c2, nil
	} else {
		// logx.Errorf("not found method: %s", rt.Name())
		// logx.Errorf("%s blocked, License key from https://teamgram.net required to unlock enterprise features.", rt.Name())
	}

	// TODO(@benqi):
	// err := mtproto.ErrMethodNotImpl
	return nil, fmt.Errorf("not found method: %s", rt.Name())
}

func (c *BFFProxyClient) Invoke(rpcMetaData *metadata.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	return c.InvokeContext(context.Background(), rpcMetaData, object)
}

// InvokeContext 通用grpc转发器
func (c *BFFProxyClient) InvokeContext(ctx context.Context, rpcMetaData *metadata.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	logger := logx.WithContext(ctx)

	conn, err := c.GetRpcClientByRequest(object)
	if err != nil {
		if r, err2 := c.TryReturnFakeRpcResult(object); err2 != nil {
			return nil, mtproto.NewRpcError(err2)
		} else {
			return r, nil
		}
	}

	// hack: layer > 177, android's createChat method use old messages.createChat#34A818
	if rpcMetaData.Client == "android" && rpcMetaData.Layer >= 177 {
		if createChat, ok := object.(*mtproto.TLMessagesCreateChat34A818); ok {
			object = &mtproto.TLMessagesCreateChat92CEDDD4{
				Constructor: mtproto.CRC32_messages_createChat92CEDDD4,
				Users:       createChat.Users,
				Title:       createChat.Title,
				TtlPeriod:   createChat.TtlPeriod,
			}
		}
	}

	t := mtproto.FindRPCContextTuple(object)
	if t == nil {
		err = fmt.Errorf("Invoke error: %v not regist!\n", object)
		logger.Error("FindRPCContextTuple error: %v", err)
		return nil, mtproto.NewRpcError(mtproto.ErrEnterpriseIsBlocked)
	}

	// logx.Infof("Invoke - method: {%s}", t.Method)
	r := t.NewReplyFunc()
	// logx.Infof("Invoke - NewReplyFunc: {%#v}, t: {%v}", r, reflect.TypeOf(r))

	var (
		header, trailer metadata.MD
		ctxWithTimeout  context.Context
	)

	// Fixed @LionPuChiPuChi, 2018-12-19
	// hack
	switch object.(type) {
	case *mtproto.TLMessagesSendMedia:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 60*time.Second)
	case *mtproto.TLMessagesSendMultiMedia:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 60*time.Second)
	case *mtproto.TLMessagesUploadMedia:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 60*time.Second)
	case *mtproto.TLMessagesEditMessage:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 60*time.Second)
	default:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 5*time.Second)
	}

	ctx2, _ := metadata.RpcMetadataToOutgoing(ctxWithTimeout, rpcMetaData)
	rt := time.Now()

	logger.Debugf("Invoke - NewReplyFunc: {%#v}", r)
	err = conn.Conn().Invoke(ctx2, t.Method, object, r, grpc.Header(&header), grpc.Trailer(&trailer))

	logger.Debugf("rpc Invoke: {method: %s, metadata: %s, result: {%s}, error: {%s}}, cost = %v",
		t.Method,
		rpcMetaData,
		reflect.TypeOf(r),
		err,
		time.Since(rt))

	// TODO(@benqi): process header from serverF
	// grpc.Header(&header)
	// log.Debugf("Invoke - error: {%v}", err)

	if err != nil {
		logger.Errorf("RPC Invoke error: {method: %s, metadata: %s, error: %s}", t.Method, rpcMetaData, err)
		return nil, mtproto.NewRpcError(err)

		//case *mysql.MySQLError:
		//if rpcErr, ok := err.(*mtproto.TLRpcError); ok {
		//	return nil, rpcErr
		//} else {
		//	return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		//}
		//
		//// TODO(@benqi): 哪些情况需要断开客户端连接
		//if s, ok := status.FromError(err); ok {
		//	//switch s.Code() {
		//	//// TODO(@benqi): Rpc error, trailer has rpc_error metadata
		//	//case codes.Unknown:
		//	//	return nil, grpc_util.RpcErrorFromMD(trailer)
		//	//}
		//	return nil, mtproto.FromGRPCStatus(s)
		//} else {
		//	return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		//}
	} else {
		if reply, ok := r.(mtproto.TLObject); !ok {
			logger.Errorf("invalid reply type, maybe server side bug, %v", reply)
			return nil, mtproto.NewRpcError(mtproto.ErrInternalServerError)
		} else {
			return reply, nil
		}
	}
}
