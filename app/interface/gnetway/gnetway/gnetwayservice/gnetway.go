/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package gnetwayservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"gnetway.sendDataToGateway": kitex.NewMethodInfo(
		sendDataToGatewayHandler,
		newSendDataToGatewayArgs,
		newSendDataToGatewayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	gnetwayServiceServiceInfo                = NewServiceInfo()
	gnetwayServiceServiceInfoForClient       = NewServiceInfoForClient()
	gnetwayServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return gnetwayServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return gnetwayServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return gnetwayServiceServiceInfoForClient
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
	serviceName := "RPCGnetway"
	handlerType := (*gnetway.RPCGnetway)(nil)
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
		"PackageName": "gnetway",
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

func sendDataToGatewayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SendDataToGatewayArgs)
	realResult := result.(*SendDataToGatewayResult)
	success, err := handler.(gnetway.RPCGnetway).GnetwaySendDataToGateway(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSendDataToGatewayArgs() interface{} {
	return &SendDataToGatewayArgs{}
}

func newSendDataToGatewayResult() interface{} {
	return &SendDataToGatewayResult{}
}

type SendDataToGatewayArgs struct {
	Req *gnetway.TLGnetwaySendDataToGateway
}

func (p *SendDataToGatewayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendDataToGatewayArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SendDataToGatewayArgs) Unmarshal(in []byte) error {
	msg := new(gnetway.TLGnetwaySendDataToGateway)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SendDataToGatewayArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SendDataToGatewayArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SendDataToGatewayArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(gnetway.TLGnetwaySendDataToGateway)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SendDataToGatewayArgs_Req_DEFAULT *gnetway.TLGnetwaySendDataToGateway

func (p *SendDataToGatewayArgs) GetReq() *gnetway.TLGnetwaySendDataToGateway {
	if !p.IsSetReq() {
		return SendDataToGatewayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendDataToGatewayArgs) IsSetReq() bool {
	return p.Req != nil
}

type SendDataToGatewayResult struct {
	Success *tg.Bool
}

var SendDataToGatewayResult_Success_DEFAULT *tg.Bool

func (p *SendDataToGatewayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SendDataToGatewayResult")
	}
	return json.Marshal(p.Success)
}

func (p *SendDataToGatewayResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendDataToGatewayResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SendDataToGatewayResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SendDataToGatewayResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendDataToGatewayResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SendDataToGatewayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendDataToGatewayResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SendDataToGatewayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendDataToGatewayResult) GetResult() interface{} {
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

func (p *kClient) GnetwaySendDataToGateway(ctx context.Context, req *gnetway.TLGnetwaySendDataToGateway) (r *tg.Bool, err error) {
	var _args SendDataToGatewayArgs
	_args.Req = req
	var _result SendDataToGatewayResult
	if err = p.c.Call(ctx, "gnetway.sendDataToGateway", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
