/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package tosservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"help.getTermsOfServiceUpdate": kitex.NewMethodInfo(
		helpGetTermsOfServiceUpdateHandler,
		newHelpGetTermsOfServiceUpdateArgs,
		newHelpGetTermsOfServiceUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.acceptTermsOfService": kitex.NewMethodInfo(
		helpAcceptTermsOfServiceHandler,
		newHelpAcceptTermsOfServiceArgs,
		newHelpAcceptTermsOfServiceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	tosServiceServiceInfo                = NewServiceInfo()
	tosServiceServiceInfoForClient       = NewServiceInfoForClient()
	tosServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return tosServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return tosServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return tosServiceServiceInfoForClient
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
	serviceName := "RPCTos"
	handlerType := (*tg.RPCTos)(nil)
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
		"PackageName": "tos",
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

func helpGetTermsOfServiceUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetTermsOfServiceUpdateArgs)
	realResult := result.(*HelpGetTermsOfServiceUpdateResult)
	success, err := handler.(tg.RPCTos).HelpGetTermsOfServiceUpdate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetTermsOfServiceUpdateArgs() interface{} {
	return &HelpGetTermsOfServiceUpdateArgs{}
}

func newHelpGetTermsOfServiceUpdateResult() interface{} {
	return &HelpGetTermsOfServiceUpdateResult{}
}

type HelpGetTermsOfServiceUpdateArgs struct {
	Req *tg.TLHelpGetTermsOfServiceUpdate
}

func (p *HelpGetTermsOfServiceUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetTermsOfServiceUpdateArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetTermsOfServiceUpdateArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetTermsOfServiceUpdate)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetTermsOfServiceUpdateArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetTermsOfServiceUpdateArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetTermsOfServiceUpdateArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetTermsOfServiceUpdate)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetTermsOfServiceUpdateArgs_Req_DEFAULT *tg.TLHelpGetTermsOfServiceUpdate

func (p *HelpGetTermsOfServiceUpdateArgs) GetReq() *tg.TLHelpGetTermsOfServiceUpdate {
	if !p.IsSetReq() {
		return HelpGetTermsOfServiceUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetTermsOfServiceUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetTermsOfServiceUpdateResult struct {
	Success *tg.HelpTermsOfServiceUpdate
}

var HelpGetTermsOfServiceUpdateResult_Success_DEFAULT *tg.HelpTermsOfServiceUpdate

func (p *HelpGetTermsOfServiceUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetTermsOfServiceUpdateResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetTermsOfServiceUpdateResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpTermsOfServiceUpdate)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetTermsOfServiceUpdateResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetTermsOfServiceUpdateResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetTermsOfServiceUpdateResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpTermsOfServiceUpdate)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetTermsOfServiceUpdateResult) GetSuccess() *tg.HelpTermsOfServiceUpdate {
	if !p.IsSetSuccess() {
		return HelpGetTermsOfServiceUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetTermsOfServiceUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpTermsOfServiceUpdate)
}

func (p *HelpGetTermsOfServiceUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetTermsOfServiceUpdateResult) GetResult() interface{} {
	return p.Success
}

func helpAcceptTermsOfServiceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpAcceptTermsOfServiceArgs)
	realResult := result.(*HelpAcceptTermsOfServiceResult)
	success, err := handler.(tg.RPCTos).HelpAcceptTermsOfService(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpAcceptTermsOfServiceArgs() interface{} {
	return &HelpAcceptTermsOfServiceArgs{}
}

func newHelpAcceptTermsOfServiceResult() interface{} {
	return &HelpAcceptTermsOfServiceResult{}
}

type HelpAcceptTermsOfServiceArgs struct {
	Req *tg.TLHelpAcceptTermsOfService
}

func (p *HelpAcceptTermsOfServiceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpAcceptTermsOfServiceArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpAcceptTermsOfServiceArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpAcceptTermsOfService)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpAcceptTermsOfServiceArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpAcceptTermsOfServiceArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpAcceptTermsOfServiceArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpAcceptTermsOfService)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpAcceptTermsOfServiceArgs_Req_DEFAULT *tg.TLHelpAcceptTermsOfService

func (p *HelpAcceptTermsOfServiceArgs) GetReq() *tg.TLHelpAcceptTermsOfService {
	if !p.IsSetReq() {
		return HelpAcceptTermsOfServiceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpAcceptTermsOfServiceArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpAcceptTermsOfServiceResult struct {
	Success *tg.Bool
}

var HelpAcceptTermsOfServiceResult_Success_DEFAULT *tg.Bool

func (p *HelpAcceptTermsOfServiceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpAcceptTermsOfServiceResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpAcceptTermsOfServiceResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpAcceptTermsOfServiceResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpAcceptTermsOfServiceResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpAcceptTermsOfServiceResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpAcceptTermsOfServiceResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return HelpAcceptTermsOfServiceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpAcceptTermsOfServiceResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *HelpAcceptTermsOfServiceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpAcceptTermsOfServiceResult) GetResult() interface{} {
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

func (p *kClient) HelpGetTermsOfServiceUpdate(ctx context.Context, req *tg.TLHelpGetTermsOfServiceUpdate) (r *tg.HelpTermsOfServiceUpdate, err error) {
	var _args HelpGetTermsOfServiceUpdateArgs
	_args.Req = req
	var _result HelpGetTermsOfServiceUpdateResult
	if err = p.c.Call(ctx, "help.getTermsOfServiceUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) HelpAcceptTermsOfService(ctx context.Context, req *tg.TLHelpAcceptTermsOfService) (r *tg.Bool, err error) {
	var _args HelpAcceptTermsOfServiceArgs
	_args.Req = req
	var _result HelpAcceptTermsOfServiceResult
	if err = p.c.Call(ctx, "help.acceptTermsOfService", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
