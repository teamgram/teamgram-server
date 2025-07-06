/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package miscellaneousservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/tg.RPCMiscellaneous/help.saveAppLog": kitex.NewMethodInfo(
		helpSaveAppLogHandler,
		newHelpSaveAppLogArgs,
		newHelpSaveAppLogResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCMiscellaneous/help.test": kitex.NewMethodInfo(
		helpTestHandler,
		newHelpTestArgs,
		newHelpTestResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	miscellaneousServiceServiceInfo                = NewServiceInfo()
	miscellaneousServiceServiceInfoForClient       = NewServiceInfoForClient()
	miscellaneousServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCMiscellaneous", miscellaneousServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCMiscellaneous", miscellaneousServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCMiscellaneous", miscellaneousServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return miscellaneousServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return miscellaneousServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return miscellaneousServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfoForClient creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}

// NewServiceInfoForStreamClient creates a new ServiceInfo containing all streaming methods
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RPCMiscellaneous"
	handlerType := (*tg.RPCMiscellaneous)(nil)
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
		"PackageName": "miscellaneous",
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

func helpSaveAppLogHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpSaveAppLogArgs)
	realResult := result.(*HelpSaveAppLogResult)
	success, err := handler.(tg.RPCMiscellaneous).HelpSaveAppLog(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpSaveAppLogArgs() interface{} {
	return &HelpSaveAppLogArgs{}
}

func newHelpSaveAppLogResult() interface{} {
	return &HelpSaveAppLogResult{}
}

type HelpSaveAppLogArgs struct {
	Req *tg.TLHelpSaveAppLog
}

func (p *HelpSaveAppLogArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpSaveAppLogArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpSaveAppLogArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpSaveAppLog)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpSaveAppLogArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpSaveAppLogArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpSaveAppLogArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpSaveAppLog)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpSaveAppLogArgs_Req_DEFAULT *tg.TLHelpSaveAppLog

func (p *HelpSaveAppLogArgs) GetReq() *tg.TLHelpSaveAppLog {
	if !p.IsSetReq() {
		return HelpSaveAppLogArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpSaveAppLogArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpSaveAppLogResult struct {
	Success *tg.Bool
}

var HelpSaveAppLogResult_Success_DEFAULT *tg.Bool

func (p *HelpSaveAppLogResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpSaveAppLogResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpSaveAppLogResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpSaveAppLogResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpSaveAppLogResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpSaveAppLogResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpSaveAppLogResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return HelpSaveAppLogResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpSaveAppLogResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *HelpSaveAppLogResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpSaveAppLogResult) GetResult() interface{} {
	return p.Success
}

func helpTestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpTestArgs)
	realResult := result.(*HelpTestResult)
	success, err := handler.(tg.RPCMiscellaneous).HelpTest(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpTestArgs() interface{} {
	return &HelpTestArgs{}
}

func newHelpTestResult() interface{} {
	return &HelpTestResult{}
}

type HelpTestArgs struct {
	Req *tg.TLHelpTest
}

func (p *HelpTestArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpTestArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpTestArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpTest)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpTestArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpTestArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpTestArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpTest)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpTestArgs_Req_DEFAULT *tg.TLHelpTest

func (p *HelpTestArgs) GetReq() *tg.TLHelpTest {
	if !p.IsSetReq() {
		return HelpTestArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpTestArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpTestResult struct {
	Success *tg.Bool
}

var HelpTestResult_Success_DEFAULT *tg.Bool

func (p *HelpTestResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpTestResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpTestResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpTestResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpTestResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpTestResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpTestResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return HelpTestResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpTestResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *HelpTestResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpTestResult) GetResult() interface{} {
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

func (p *kClient) HelpSaveAppLog(ctx context.Context, req *tg.TLHelpSaveAppLog) (r *tg.Bool, err error) {
	// var _args HelpSaveAppLogArgs
	// _args.Req = req
	// var _result HelpSaveAppLogResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCMiscellaneous/help.saveAppLog", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpTest(ctx context.Context, req *tg.TLHelpTest) (r *tg.Bool, err error) {
	// var _args HelpTestArgs
	// _args.Req = req
	// var _result HelpTestResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCMiscellaneous/help.test", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
