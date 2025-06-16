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
	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"updates.getState": kitex.NewMethodInfo(
		updatesGetStateHandler,
		newUpdatesGetStateArgs,
		newUpdatesGetStateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"updates.getDifference": kitex.NewMethodInfo(
		updatesGetDifferenceHandler,
		newUpdatesGetDifferenceArgs,
		newUpdatesGetDifferenceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"updates.getChannelDifference": kitex.NewMethodInfo(
		updatesGetChannelDifferenceHandler,
		newUpdatesGetChannelDifferenceArgs,
		newUpdatesGetChannelDifferenceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	updatesServiceServiceInfo                = NewServiceInfo()
	updatesServiceServiceInfoForClient       = NewServiceInfoForClient()
	updatesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

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
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RPCUpdates"
	handlerType := (*tg.RPCUpdates)(nil)
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

func updatesGetStateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatesGetStateArgs)
	realResult := result.(*UpdatesGetStateResult)
	success, err := handler.(tg.RPCUpdates).UpdatesGetState(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatesGetStateArgs() interface{} {
	return &UpdatesGetStateArgs{}
}

func newUpdatesGetStateResult() interface{} {
	return &UpdatesGetStateResult{}
}

type UpdatesGetStateArgs struct {
	Req *tg.TLUpdatesGetState
}

func (p *UpdatesGetStateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatesGetStateArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatesGetStateArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUpdatesGetState)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatesGetStateArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatesGetStateArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatesGetStateArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUpdatesGetState)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatesGetStateArgs_Req_DEFAULT *tg.TLUpdatesGetState

func (p *UpdatesGetStateArgs) GetReq() *tg.TLUpdatesGetState {
	if !p.IsSetReq() {
		return UpdatesGetStateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatesGetStateArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatesGetStateResult struct {
	Success *tg.UpdatesState
}

var UpdatesGetStateResult_Success_DEFAULT *tg.UpdatesState

func (p *UpdatesGetStateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatesGetStateResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatesGetStateResult) Unmarshal(in []byte) error {
	msg := new(tg.UpdatesState)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesGetStateResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatesGetStateResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatesGetStateResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UpdatesState)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesGetStateResult) GetSuccess() *tg.UpdatesState {
	if !p.IsSetSuccess() {
		return UpdatesGetStateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatesGetStateResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UpdatesState)
}

func (p *UpdatesGetStateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatesGetStateResult) GetResult() interface{} {
	return p.Success
}

func updatesGetDifferenceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatesGetDifferenceArgs)
	realResult := result.(*UpdatesGetDifferenceResult)
	success, err := handler.(tg.RPCUpdates).UpdatesGetDifference(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatesGetDifferenceArgs() interface{} {
	return &UpdatesGetDifferenceArgs{}
}

func newUpdatesGetDifferenceResult() interface{} {
	return &UpdatesGetDifferenceResult{}
}

type UpdatesGetDifferenceArgs struct {
	Req *tg.TLUpdatesGetDifference
}

func (p *UpdatesGetDifferenceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatesGetDifferenceArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatesGetDifferenceArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUpdatesGetDifference)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatesGetDifferenceArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatesGetDifferenceArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatesGetDifferenceArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUpdatesGetDifference)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatesGetDifferenceArgs_Req_DEFAULT *tg.TLUpdatesGetDifference

func (p *UpdatesGetDifferenceArgs) GetReq() *tg.TLUpdatesGetDifference {
	if !p.IsSetReq() {
		return UpdatesGetDifferenceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatesGetDifferenceArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatesGetDifferenceResult struct {
	Success *tg.UpdatesDifference
}

var UpdatesGetDifferenceResult_Success_DEFAULT *tg.UpdatesDifference

func (p *UpdatesGetDifferenceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatesGetDifferenceResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatesGetDifferenceResult) Unmarshal(in []byte) error {
	msg := new(tg.UpdatesDifference)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesGetDifferenceResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatesGetDifferenceResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatesGetDifferenceResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UpdatesDifference)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesGetDifferenceResult) GetSuccess() *tg.UpdatesDifference {
	if !p.IsSetSuccess() {
		return UpdatesGetDifferenceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatesGetDifferenceResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UpdatesDifference)
}

func (p *UpdatesGetDifferenceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatesGetDifferenceResult) GetResult() interface{} {
	return p.Success
}

func updatesGetChannelDifferenceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatesGetChannelDifferenceArgs)
	realResult := result.(*UpdatesGetChannelDifferenceResult)
	success, err := handler.(tg.RPCUpdates).UpdatesGetChannelDifference(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatesGetChannelDifferenceArgs() interface{} {
	return &UpdatesGetChannelDifferenceArgs{}
}

func newUpdatesGetChannelDifferenceResult() interface{} {
	return &UpdatesGetChannelDifferenceResult{}
}

type UpdatesGetChannelDifferenceArgs struct {
	Req *tg.TLUpdatesGetChannelDifference
}

func (p *UpdatesGetChannelDifferenceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatesGetChannelDifferenceArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatesGetChannelDifferenceArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUpdatesGetChannelDifference)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatesGetChannelDifferenceArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatesGetChannelDifferenceArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatesGetChannelDifferenceArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUpdatesGetChannelDifference)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatesGetChannelDifferenceArgs_Req_DEFAULT *tg.TLUpdatesGetChannelDifference

func (p *UpdatesGetChannelDifferenceArgs) GetReq() *tg.TLUpdatesGetChannelDifference {
	if !p.IsSetReq() {
		return UpdatesGetChannelDifferenceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatesGetChannelDifferenceArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatesGetChannelDifferenceResult struct {
	Success *tg.UpdatesChannelDifference
}

var UpdatesGetChannelDifferenceResult_Success_DEFAULT *tg.UpdatesChannelDifference

func (p *UpdatesGetChannelDifferenceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatesGetChannelDifferenceResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatesGetChannelDifferenceResult) Unmarshal(in []byte) error {
	msg := new(tg.UpdatesChannelDifference)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesGetChannelDifferenceResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatesGetChannelDifferenceResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatesGetChannelDifferenceResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UpdatesChannelDifference)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesGetChannelDifferenceResult) GetSuccess() *tg.UpdatesChannelDifference {
	if !p.IsSetSuccess() {
		return UpdatesGetChannelDifferenceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatesGetChannelDifferenceResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UpdatesChannelDifference)
}

func (p *UpdatesGetChannelDifferenceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatesGetChannelDifferenceResult) GetResult() interface{} {
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

func (p *kClient) UpdatesGetState(ctx context.Context, req *tg.TLUpdatesGetState) (r *tg.UpdatesState, err error) {
	var _args UpdatesGetStateArgs
	_args.Req = req
	var _result UpdatesGetStateResult
	if err = p.c.Call(ctx, "updates.getState", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatesGetDifference(ctx context.Context, req *tg.TLUpdatesGetDifference) (r *tg.UpdatesDifference, err error) {
	var _args UpdatesGetDifferenceArgs
	_args.Req = req
	var _result UpdatesGetDifferenceResult
	if err = p.c.Call(ctx, "updates.getDifference", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatesGetChannelDifference(ctx context.Context, req *tg.TLUpdatesGetChannelDifference) (r *tg.UpdatesChannelDifference, err error) {
	var _args UpdatesGetChannelDifferenceArgs
	_args.Req = req
	var _result UpdatesGetChannelDifferenceResult
	if err = p.c.Call(ctx, "updates.getChannelDifference", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
