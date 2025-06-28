/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo2service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/echo2"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"echo2.echo": kitex.NewMethodInfo(
		echoHandler,
		newEchoArgs,
		newEchoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	echo2ServiceServiceInfo                = NewServiceInfo()
	echo2ServiceServiceInfoForClient       = NewServiceInfoForClient()
	echo2ServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCEcho2", echo2ServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCEcho2", echo2ServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCEcho2", echo2ServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return echo2ServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return echo2ServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return echo2ServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfoForClient creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}

// NewServiceInfoForStreamClient creates a new ServiceInfo containing streaming methods
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RPCEcho2"
	handlerType := (*echo2.RPCEcho2)(nil)
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
		"PackageName": "echo2",
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
	realArg := arg.(*EchoArgs)
	realResult := result.(*EchoResult)
	success, err := handler.(echo2.RPCEcho2).Echo2Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEchoArgs() interface{} {
	return &EchoArgs{}
}

func newEchoResult() interface{} {
	return &EchoResult{}
}

type EchoArgs struct {
	Req *echo2.TLEcho2Echo
}

func (p *EchoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EchoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EchoArgs) Unmarshal(in []byte) error {
	msg := new(echo2.TLEcho2Echo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EchoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EchoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EchoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(echo2.TLEcho2Echo)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EchoArgs_Req_DEFAULT *echo2.TLEcho2Echo

func (p *EchoArgs) GetReq() *echo2.TLEcho2Echo {
	if !p.IsSetReq() {
		return EchoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EchoArgs) IsSetReq() bool {
	return p.Req != nil
}

type EchoResult struct {
	Success *echo2.Echo
}

var EchoResult_Success_DEFAULT *echo2.Echo

func (p *EchoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EchoResult")
	}
	return json.Marshal(p.Success)
}

func (p *EchoResult) Unmarshal(in []byte) error {
	msg := new(echo2.Echo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EchoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EchoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EchoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(echo2.Echo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EchoResult) GetSuccess() *echo2.Echo {
	if !p.IsSetSuccess() {
		return EchoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EchoResult) SetSuccess(x interface{}) {
	p.Success = x.(*echo2.Echo)
}

func (p *EchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo2Echo(ctx context.Context, req *echo2.TLEcho2Echo) (r *echo2.Echo, err error) {
	// var _args EchoArgs
	// _args.Req = req
	// var _result EchoResult

	_result := new(echo2.Echo)

	if err = p.c.Call(ctx, "echo2.echo", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
