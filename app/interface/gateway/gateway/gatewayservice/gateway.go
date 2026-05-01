/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gatewayservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/gateway.RPCGateway/gateway.pushUpdatesData": kitex.NewMethodInfo(
		pushUpdatesDataHandler,
		newPushUpdatesDataArgs,
		newPushUpdatesDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/gateway.RPCGateway/gateway.pushSessionUpdatesData": kitex.NewMethodInfo(
		pushSessionUpdatesDataHandler,
		newPushSessionUpdatesDataArgs,
		newPushSessionUpdatesDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/gateway.RPCGateway/gateway.pushRpcResultData": kitex.NewMethodInfo(
		pushRpcResultDataHandler,
		newPushRpcResultDataArgs,
		newPushRpcResultDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	gatewayServiceServiceInfo                = NewServiceInfo()
	gatewayServiceServiceInfoForClient       = NewServiceInfoForClient()
	gatewayServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCGateway", gatewayServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCGateway", gatewayServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCGateway", gatewayServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return gatewayServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return gatewayServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return gatewayServiceServiceInfoForClient
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
	serviceName := "RPCGateway"
	handlerType := (*gateway.RPCGateway)(nil)
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
		"PackageName": "gateway",
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

func pushUpdatesDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushUpdatesDataArgs)
	realResult := result.(*PushUpdatesDataResult)
	success, err := handler.(gateway.RPCGateway).GatewayPushUpdatesData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPushUpdatesDataArgs() interface{} {
	return &PushUpdatesDataArgs{}
}

func newPushUpdatesDataResult() interface{} {
	return &PushUpdatesDataResult{}
}

type PushUpdatesDataArgs struct {
	Req *gateway.TLGatewayPushUpdatesData
}

func (p *PushUpdatesDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushUpdatesDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushUpdatesDataArgs) Unmarshal(in []byte) error {
	msg := new(gateway.TLGatewayPushUpdatesData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PushUpdatesDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PushUpdatesDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PushUpdatesDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(gateway.TLGatewayPushUpdatesData)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PushUpdatesDataArgs_Req_DEFAULT *gateway.TLGatewayPushUpdatesData

func (p *PushUpdatesDataArgs) GetReq() *gateway.TLGatewayPushUpdatesData {
	if !p.IsSetReq() {
		return PushUpdatesDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushUpdatesDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type PushUpdatesDataResult struct {
	Success *tg.Bool
}

var PushUpdatesDataResult_Success_DEFAULT *tg.Bool

func (p *PushUpdatesDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PushUpdatesDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *PushUpdatesDataResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushUpdatesDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PushUpdatesDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PushUpdatesDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushUpdatesDataResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return PushUpdatesDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushUpdatesDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *PushUpdatesDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushUpdatesDataResult) GetResult() interface{} {
	return p.Success
}

func pushSessionUpdatesDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushSessionUpdatesDataArgs)
	realResult := result.(*PushSessionUpdatesDataResult)
	success, err := handler.(gateway.RPCGateway).GatewayPushSessionUpdatesData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPushSessionUpdatesDataArgs() interface{} {
	return &PushSessionUpdatesDataArgs{}
}

func newPushSessionUpdatesDataResult() interface{} {
	return &PushSessionUpdatesDataResult{}
}

type PushSessionUpdatesDataArgs struct {
	Req *gateway.TLGatewayPushSessionUpdatesData
}

func (p *PushSessionUpdatesDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushSessionUpdatesDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushSessionUpdatesDataArgs) Unmarshal(in []byte) error {
	msg := new(gateway.TLGatewayPushSessionUpdatesData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PushSessionUpdatesDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PushSessionUpdatesDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PushSessionUpdatesDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(gateway.TLGatewayPushSessionUpdatesData)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PushSessionUpdatesDataArgs_Req_DEFAULT *gateway.TLGatewayPushSessionUpdatesData

func (p *PushSessionUpdatesDataArgs) GetReq() *gateway.TLGatewayPushSessionUpdatesData {
	if !p.IsSetReq() {
		return PushSessionUpdatesDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushSessionUpdatesDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type PushSessionUpdatesDataResult struct {
	Success *tg.Bool
}

var PushSessionUpdatesDataResult_Success_DEFAULT *tg.Bool

func (p *PushSessionUpdatesDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PushSessionUpdatesDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *PushSessionUpdatesDataResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushSessionUpdatesDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PushSessionUpdatesDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PushSessionUpdatesDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushSessionUpdatesDataResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return PushSessionUpdatesDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushSessionUpdatesDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *PushSessionUpdatesDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushSessionUpdatesDataResult) GetResult() interface{} {
	return p.Success
}

func pushRpcResultDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushRpcResultDataArgs)
	realResult := result.(*PushRpcResultDataResult)
	success, err := handler.(gateway.RPCGateway).GatewayPushRpcResultData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPushRpcResultDataArgs() interface{} {
	return &PushRpcResultDataArgs{}
}

func newPushRpcResultDataResult() interface{} {
	return &PushRpcResultDataResult{}
}

type PushRpcResultDataArgs struct {
	Req *gateway.TLGatewayPushRpcResultData
}

func (p *PushRpcResultDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushRpcResultDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushRpcResultDataArgs) Unmarshal(in []byte) error {
	msg := new(gateway.TLGatewayPushRpcResultData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PushRpcResultDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PushRpcResultDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PushRpcResultDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(gateway.TLGatewayPushRpcResultData)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PushRpcResultDataArgs_Req_DEFAULT *gateway.TLGatewayPushRpcResultData

func (p *PushRpcResultDataArgs) GetReq() *gateway.TLGatewayPushRpcResultData {
	if !p.IsSetReq() {
		return PushRpcResultDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushRpcResultDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type PushRpcResultDataResult struct {
	Success *tg.Bool
}

var PushRpcResultDataResult_Success_DEFAULT *tg.Bool

func (p *PushRpcResultDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PushRpcResultDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *PushRpcResultDataResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushRpcResultDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PushRpcResultDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PushRpcResultDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushRpcResultDataResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return PushRpcResultDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushRpcResultDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *PushRpcResultDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushRpcResultDataResult) GetResult() interface{} {
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

func (p *kClient) GatewayPushUpdatesData(ctx context.Context, req *gateway.TLGatewayPushUpdatesData) (r *tg.Bool, err error) {
	// var _args PushUpdatesDataArgs
	// _args.Req = req
	// var _result PushUpdatesDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/gateway.RPCGateway/gateway.pushUpdatesData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) GatewayPushSessionUpdatesData(ctx context.Context, req *gateway.TLGatewayPushSessionUpdatesData) (r *tg.Bool, err error) {
	// var _args PushSessionUpdatesDataArgs
	// _args.Req = req
	// var _result PushSessionUpdatesDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/gateway.RPCGateway/gateway.pushSessionUpdatesData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) GatewayPushRpcResultData(ctx context.Context, req *gateway.TLGatewayPushRpcResultData) (r *tg.Bool, err error) {
	// var _args PushRpcResultDataArgs
	// _args.Req = req
	// var _result PushRpcResultDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/gateway.RPCGateway/gateway.pushRpcResultData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
