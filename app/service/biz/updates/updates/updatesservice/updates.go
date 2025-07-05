/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package updatesservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"updates.getStateV2": kitex.NewMethodInfo(
		getStateV2Handler,
		newGetStateV2Args,
		newGetStateV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"updates.getDifferenceV2": kitex.NewMethodInfo(
		getDifferenceV2Handler,
		newGetDifferenceV2Args,
		newGetDifferenceV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"updates.getChannelDifferenceV2": kitex.NewMethodInfo(
		getChannelDifferenceV2Handler,
		newGetChannelDifferenceV2Args,
		newGetChannelDifferenceV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	updatesServiceServiceInfo                = NewServiceInfo()
	updatesServiceServiceInfoForClient       = NewServiceInfoForClient()
	updatesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCUpdates", updatesServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCUpdates", updatesServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCUpdates", updatesServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return updatesServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return updatesServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return updatesServiceServiceInfoForClient
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
	serviceName := "RPCUpdates"
	handlerType := (*updates.RPCUpdates)(nil)
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
		"PackageName": "updates",
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

func getStateV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetStateV2Args)
	realResult := result.(*GetStateV2Result)
	success, err := handler.(updates.RPCUpdates).UpdatesGetStateV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetStateV2Args() interface{} {
	return &GetStateV2Args{}
}

func newGetStateV2Result() interface{} {
	return &GetStateV2Result{}
}

type GetStateV2Args struct {
	Req *updates.TLUpdatesGetStateV2
}

func (p *GetStateV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetStateV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *GetStateV2Args) Unmarshal(in []byte) error {
	msg := new(updates.TLUpdatesGetStateV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetStateV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetStateV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetStateV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(updates.TLUpdatesGetStateV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetStateV2Args_Req_DEFAULT *updates.TLUpdatesGetStateV2

func (p *GetStateV2Args) GetReq() *updates.TLUpdatesGetStateV2 {
	if !p.IsSetReq() {
		return GetStateV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *GetStateV2Args) IsSetReq() bool {
	return p.Req != nil
}

type GetStateV2Result struct {
	Success *tg.UpdatesState
}

var GetStateV2Result_Success_DEFAULT *tg.UpdatesState

func (p *GetStateV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetStateV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *GetStateV2Result) Unmarshal(in []byte) error {
	msg := new(tg.UpdatesState)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetStateV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetStateV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetStateV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UpdatesState)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetStateV2Result) GetSuccess() *tg.UpdatesState {
	if !p.IsSetSuccess() {
		return GetStateV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *GetStateV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UpdatesState)
}

func (p *GetStateV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetStateV2Result) GetResult() interface{} {
	return p.Success
}

func getDifferenceV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDifferenceV2Args)
	realResult := result.(*GetDifferenceV2Result)
	success, err := handler.(updates.RPCUpdates).UpdatesGetDifferenceV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDifferenceV2Args() interface{} {
	return &GetDifferenceV2Args{}
}

func newGetDifferenceV2Result() interface{} {
	return &GetDifferenceV2Result{}
}

type GetDifferenceV2Args struct {
	Req *updates.TLUpdatesGetDifferenceV2
}

func (p *GetDifferenceV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDifferenceV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *GetDifferenceV2Args) Unmarshal(in []byte) error {
	msg := new(updates.TLUpdatesGetDifferenceV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDifferenceV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDifferenceV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDifferenceV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(updates.TLUpdatesGetDifferenceV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDifferenceV2Args_Req_DEFAULT *updates.TLUpdatesGetDifferenceV2

func (p *GetDifferenceV2Args) GetReq() *updates.TLUpdatesGetDifferenceV2 {
	if !p.IsSetReq() {
		return GetDifferenceV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDifferenceV2Args) IsSetReq() bool {
	return p.Req != nil
}

type GetDifferenceV2Result struct {
	Success *updates.Difference
}

var GetDifferenceV2Result_Success_DEFAULT *updates.Difference

func (p *GetDifferenceV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDifferenceV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *GetDifferenceV2Result) Unmarshal(in []byte) error {
	msg := new(updates.Difference)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDifferenceV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDifferenceV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDifferenceV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(updates.Difference)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDifferenceV2Result) GetSuccess() *updates.Difference {
	if !p.IsSetSuccess() {
		return GetDifferenceV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDifferenceV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*updates.Difference)
}

func (p *GetDifferenceV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDifferenceV2Result) GetResult() interface{} {
	return p.Success
}

func getChannelDifferenceV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetChannelDifferenceV2Args)
	realResult := result.(*GetChannelDifferenceV2Result)
	success, err := handler.(updates.RPCUpdates).UpdatesGetChannelDifferenceV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetChannelDifferenceV2Args() interface{} {
	return &GetChannelDifferenceV2Args{}
}

func newGetChannelDifferenceV2Result() interface{} {
	return &GetChannelDifferenceV2Result{}
}

type GetChannelDifferenceV2Args struct {
	Req *updates.TLUpdatesGetChannelDifferenceV2
}

func (p *GetChannelDifferenceV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChannelDifferenceV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *GetChannelDifferenceV2Args) Unmarshal(in []byte) error {
	msg := new(updates.TLUpdatesGetChannelDifferenceV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetChannelDifferenceV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetChannelDifferenceV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetChannelDifferenceV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(updates.TLUpdatesGetChannelDifferenceV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetChannelDifferenceV2Args_Req_DEFAULT *updates.TLUpdatesGetChannelDifferenceV2

func (p *GetChannelDifferenceV2Args) GetReq() *updates.TLUpdatesGetChannelDifferenceV2 {
	if !p.IsSetReq() {
		return GetChannelDifferenceV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChannelDifferenceV2Args) IsSetReq() bool {
	return p.Req != nil
}

type GetChannelDifferenceV2Result struct {
	Success *updates.ChannelDifference
}

var GetChannelDifferenceV2Result_Success_DEFAULT *updates.ChannelDifference

func (p *GetChannelDifferenceV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChannelDifferenceV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *GetChannelDifferenceV2Result) Unmarshal(in []byte) error {
	msg := new(updates.ChannelDifference)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChannelDifferenceV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetChannelDifferenceV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetChannelDifferenceV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(updates.ChannelDifference)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChannelDifferenceV2Result) GetSuccess() *updates.ChannelDifference {
	if !p.IsSetSuccess() {
		return GetChannelDifferenceV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChannelDifferenceV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*updates.ChannelDifference)
}

func (p *GetChannelDifferenceV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetChannelDifferenceV2Result) GetResult() interface{} {
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

func (p *kClient) UpdatesGetStateV2(ctx context.Context, req *updates.TLUpdatesGetStateV2) (r *tg.UpdatesState, err error) {
	// var _args GetStateV2Args
	// _args.Req = req
	// var _result GetStateV2Result

	_result := new(tg.UpdatesState)

	if err = p.c.Call(ctx, "updates.getStateV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UpdatesGetDifferenceV2(ctx context.Context, req *updates.TLUpdatesGetDifferenceV2) (r *updates.Difference, err error) {
	// var _args GetDifferenceV2Args
	// _args.Req = req
	// var _result GetDifferenceV2Result

	_result := new(updates.Difference)

	if err = p.c.Call(ctx, "updates.getDifferenceV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UpdatesGetChannelDifferenceV2(ctx context.Context, req *updates.TLUpdatesGetChannelDifferenceV2) (r *updates.ChannelDifference, err error) {
	// var _args GetChannelDifferenceV2Args
	// _args.Req = req
	// var _result GetChannelDifferenceV2Result

	_result := new(updates.ChannelDifference)

	if err = p.c.Call(ctx, "updates.getChannelDifferenceV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
