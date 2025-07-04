// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package bffproxyclient

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/rpc/metadata"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/zeromicro/go-zero/core/logx"
)

//type proxyClient struct {
//	c kitex.Client
//}

//// NewClientByServiceInfoForClient returns a Client.
//func NewClientByServiceInfoForClient(c RpcClientConf, svc *serviceinfo.ServiceInfo) (Client, error) {
//	return NewClient(c, func(opts ...client.Option) (Client, error) {
//		return client.NewClient(svc, opts...)
//	})
//}

type BFFProxyClient2 struct {
	// zrpc.Client
	BFFClients map[string]kitex.Client
}

func NewBFFProxyClient2(cList []BFFProxyClientConf) *BFFProxyClient2 {
	var (
		clients   = make(map[string]kitex.Client)
		registers = iface.GetRPCContextRegisters()
	)

	for _, c := range cList {
		for _, serviceName := range c.ServiceNameList {
			c2 := c.RpcClientConf
			c2.ServiceName = serviceName

			cli := kitex.GetCachedKitexClient(c2)
			clients[serviceName] = cli
		}
	}

	bizClients := make(map[string]kitex.Client)

	for m, ctx := range registers {
		for k, v := range clients {
			if strings.HasPrefix(ctx.Method[4:], k) {
				bizClients[m] = v
				break
			}
		}
	}

	return &BFFProxyClient2{
		BFFClients: bizClients,
	}
}

func (c *BFFProxyClient2) GetRpcClientByRequest(t interface{}) (kitex.Client, error) {
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

func (c *BFFProxyClient2) Invoke(rpcMetaData *metadata.RpcMetadata, object iface.TLObject) (iface.TLObject, error) {
	return c.InvokeContext(context.Background(), rpcMetaData, object)
}

// InvokeContext 通用grpc转发器
func (c *BFFProxyClient2) InvokeContext(ctx context.Context, rpcMetaData *metadata.RpcMetadata, object iface.TLObject) (iface.TLObject, error) {
	logger := logx.WithContext(ctx)

	conn, err := c.GetRpcClientByRequest(object)
	if err != nil {
		if r, err2 := c.TryReturnFakeRpcResult(object); err2 != nil {
			return nil, tg.NewRpcError(err2)
		} else {
			return r, nil
		}
	}

	// hack: layer > 177, android's createChat method use old messages.createChat#34A818
	//if rpcMetaData.Client == "android" && rpcMetaData.Layer >= 177 {
	//	if createChat, ok := object.(*tg.TLMessagesCreateChat); ok {
	//		object = &tg.TLMessagesCreateChat{
	//			ClazzID:   tg.ClazzID_messages_createChat,
	//			Users:     createChat.Users,
	//			Title:     createChat.Title,
	//			TtlPeriod: createChat.TtlPeriod,
	//		}
	//	}
	//}

	t := iface.FindRPCContextTuple(object)
	if t == nil {
		err = fmt.Errorf("Invoke error: %v not regist!\n", object)
		logger.Error("FindRPCContextTuple error: %v", err)
		return nil, tg.NewRpcError(tg.ErrEnterpriseIsBlocked)
	}

	// logx.Infof("Invoke - method: {%s}", t.Method)
	r := t.NewReplyFunc()
	// logx.Infof("Invoke - NewReplyFunc: {%#v}, t: {%v}", r, reflect.TypeOf(r))

	var (
		// header, trailer metadata.MD
		ctxWithTimeout context.Context
	)

	// Fixed @LionPuChiPuChi, 2018-12-19
	// hack
	switch object.(type) {
	case *tg.TLMessagesSendMedia:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 60*time.Second)
	case *tg.TLMessagesSendMultiMedia:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 60*time.Second)
	case *tg.TLMessagesUploadMedia:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 60*time.Second)
	case *tg.TLMessagesEditMessage:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 60*time.Second)
	default:
		ctxWithTimeout, _ = context.WithTimeout(context.Background(), 5*time.Second)
	}

	// ctx2, _ := metadata.RpcMetadataToOutgoing(ctxWithTimeout, rpcMetaData)
	rt := time.Now()

	logger.Debugf("Invoke - NewReplyFunc: {%#v}", r)
	// err = conn.Conn().Invoke(ctx2, t.Method, object, r, grpc.Header(&header), grpc.Trailer(&trailer))
	err = conn.Call(ctxWithTimeout, t.Method, object, r)
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
		return nil, tg.NewRpcError(err)

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
		if reply, ok := r.(iface.TLObject); !ok {
			logger.Errorf("invalid reply type, maybe server side bug, %v", reply)
			return nil, tg.NewRpcError(tg.ErrInternalServerError)
		} else {
			return reply, nil
		}
	}
}
