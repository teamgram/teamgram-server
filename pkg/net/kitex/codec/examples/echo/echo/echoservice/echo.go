/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echoservice

import (
	"context"
	"errors"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"echo.echo": kitex.NewMethodInfo(
		echoHandler,
		echo.NewTLEchoEchoArg,
		echo.NewEchoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	echoServiceServiceInfo                = NewServiceInfo()
	echoServiceServiceInfoForClient       = NewServiceInfoForClient()
	echoServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return echoServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return echoServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return echoServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfoForClient creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RPCEcho"
	handlerType := (*echo.RPCEcho)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "echo",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		KiteXGenVersion: "0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TLEchoEcho)
	realResult := result.(*echo.Echo)
	success, err := handler.(echo.RPCEcho).EchoEcho(ctx, realArg)
	if err != nil {
		return err
	}
	realResult.EchoClazz = success.EchoClazz
	return nil
}

//func newEchoArgs() interface{} {
//	return &EchoArgs{}
//}
//
//func newEchoResult() interface{} {
//	return &EchoResult{}
//}
//
//type EchoArgs struct {
//	Req *echo.TLEchoEcho
//}
//
//func (p *EchoArgs) Marshal(out []byte) ([]byte, error) {
//	if !p.IsSetReq() {
//		return out, fmt.Errorf("No req in EchoArgs")
//	}
//	return json.Marshal(p.Req)
//}
//
//func (p *EchoArgs) Unmarshal(in []byte) error {
//	msg := new(echo.TLEchoEcho)
//	if err := json.Unmarshal(in, msg); err != nil {
//		return err
//	}
//	p.Req = msg
//	return nil
//}
//
//func (p *EchoArgs) Encode(x *bin.Encoder, layer int32) error {
//	if !p.IsSetReq() {
//		return fmt.Errorf("No req in EchoArgs")
//	}
//
//	return p.Req.Encode(x, layer)
//}
//
//func (p *EchoArgs) Decode(d *bin.Decoder) (err error) {
//	msg := new(echo.TLEchoEcho)
//	msg.ClazzID, _ = d.ClazzID()
//	_ = msg.Decode(d)
//	p.Req = msg
//	return nil
//}
//
//var EchoArgs_Req_DEFAULT *echo.TLEchoEcho
//
//func (p *EchoArgs) GetReq() *echo.TLEchoEcho {
//	if !p.IsSetReq() {
//		return EchoArgs_Req_DEFAULT
//	}
//	return p.Req
//}
//
//func (p *EchoArgs) IsSetReq() bool {
//	return p.Req != nil
//}
//
//type EchoResult struct {
//	Success *echo.Echo
//}
//
//var EchoResult_Success_DEFAULT *echo.Echo
//
//func (p *EchoResult) Marshal(out []byte) ([]byte, error) {
//	if !p.IsSetSuccess() {
//		return out, fmt.Errorf("No req in EchoResult")
//	}
//	return json.Marshal(p.Success)
//}
//
//func (p *EchoResult) Unmarshal(in []byte) error {
//	msg := new(echo.Echo)
//	if err := json.Unmarshal(in, msg); err != nil {
//		return err
//	}
//	p.Success = msg
//	return nil
//}
//
//func (p *EchoResult) Encode(x *bin.Encoder, layer int32) error {
//	if !p.IsSetSuccess() {
//		return fmt.Errorf("No req in EchoResult")
//	}
//
//	return p.Success.Encode(x, layer)
//}
//
//func (p *EchoResult) Decode(d *bin.Decoder) (err error) {
//	msg := new(echo.Echo)
//	if err = msg.Decode(d); err != nil {
//		return err
//	}
//	p.Success = msg
//	return nil
//}
//
//func (p *EchoResult) GetSuccess() *echo.Echo {
//	if !p.IsSetSuccess() {
//		return EchoResult_Success_DEFAULT
//	}
//	return p.Success
//}
//
//func (p *EchoResult) SetSuccess(x interface{}) {
//	p.Success = x.(*echo.Echo)
//}
//
//func (p *EchoResult) IsSetSuccess() bool {
//	return p.Success != nil
//}
//
//func (p *EchoResult) GetResult() interface{} {
//	return p.Success
//}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EchoEcho(ctx context.Context, req *echo.TLEchoEcho) (r *echo.Echo, err error) {
	//var _args EchoArgs
	//_args.Req = req
	//var _result EchoResult
	r = echo.NewEchoResult().(*echo.Echo)
	if err = p.c.Call(ctx, "echo.echo", req, r); err != nil {
		return
	}
	return r, nil
}
