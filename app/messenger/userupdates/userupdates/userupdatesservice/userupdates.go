/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userupdatesservice

import (
	"context"
	"encoding/json"
	"fmt"
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/userupdates.RPCUserupdates/userupdates.processUserOperation": kitex.NewMethodInfo(
		processUserOperationHandler,
		newProcessUserOperationArgs,
		newProcessUserOperationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/userupdates.RPCUserupdates/userupdates.getOperationResult": kitex.NewMethodInfo(
		getOperationResultHandler,
		newGetOperationResultArgs,
		newGetOperationResultResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/userupdates.RPCUserupdates/userupdates.getState": kitex.NewMethodInfo(
		getStateHandler,
		newGetStateArgs,
		newGetStateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/userupdates.RPCUserupdates/userupdates.getDifference": kitex.NewMethodInfo(
		getDifferenceHandler,
		newGetDifferenceArgs,
		newGetDifferenceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	
}

var (
    userupdatesServiceServiceInfo                = NewServiceInfo()
    userupdatesServiceServiceInfoForClient       = NewServiceInfoForClient()
    userupdatesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
    iface.RegisterKitexServiceInfo("RPCUserupdates", userupdatesServiceServiceInfo)
    iface.RegisterKitexServiceInfoForClient("RPCUserupdates", userupdatesServiceServiceInfoForClient)
    iface.RegisterKitexServiceInfoForStreamClient("RPCUserupdates", userupdatesServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userupdatesServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userupdatesServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userupdatesServiceServiceInfoForClient
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
	serviceName := "RPCUserupdates"
	handlerType := (*userupdates.RPCUserupdates)(nil)
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
		"PackageName": "userupdates",
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


func processUserOperationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ProcessUserOperationArgs)
	realResult := result.(*ProcessUserOperationResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesProcessUserOperation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newProcessUserOperationArgs() interface{} {
	return &ProcessUserOperationArgs{}
}

func newProcessUserOperationResult() interface{} {
	return &ProcessUserOperationResult{}
}

type ProcessUserOperationArgs struct {
	Req *userupdates.TLUserupdatesProcessUserOperation
}

func (p *ProcessUserOperationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ProcessUserOperationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ProcessUserOperationArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesProcessUserOperation)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ProcessUserOperationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ProcessUserOperationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ProcessUserOperationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesProcessUserOperation)
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

var ProcessUserOperationArgs_Req_DEFAULT *userupdates.TLUserupdatesProcessUserOperation

func (p *ProcessUserOperationArgs) GetReq() *userupdates.TLUserupdatesProcessUserOperation {
	if !p.IsSetReq() {
		return ProcessUserOperationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProcessUserOperationArgs) IsSetReq() bool {
	return p.Req != nil
}

type ProcessUserOperationResult struct {
	Success *userupdates.UserOperationResult
}

var ProcessUserOperationResult_Success_DEFAULT *userupdates.UserOperationResult

func (p *ProcessUserOperationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ProcessUserOperationResult")
	}
	return json.Marshal(p.Success)
}

func (p *ProcessUserOperationResult) Unmarshal(in []byte) error {
	msg := new(userupdates.UserOperationResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessUserOperationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ProcessUserOperationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ProcessUserOperationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserOperationResult)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessUserOperationResult) GetSuccess() *userupdates.UserOperationResult {
	if !p.IsSetSuccess() {
		return ProcessUserOperationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProcessUserOperationResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.UserOperationResult)
}


func (p *ProcessUserOperationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProcessUserOperationResult) GetResult() interface{} {
	return p.Success
}


func getOperationResultHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetOperationResultArgs)
	realResult := result.(*GetOperationResultResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesGetOperationResult(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetOperationResultArgs() interface{} {
	return &GetOperationResultArgs{}
}

func newGetOperationResultResult() interface{} {
	return &GetOperationResultResult{}
}

type GetOperationResultArgs struct {
	Req *userupdates.TLUserupdatesGetOperationResult
}

func (p *GetOperationResultArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetOperationResultArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetOperationResultArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesGetOperationResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetOperationResultArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetOperationResultArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetOperationResultArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesGetOperationResult)
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

var GetOperationResultArgs_Req_DEFAULT *userupdates.TLUserupdatesGetOperationResult

func (p *GetOperationResultArgs) GetReq() *userupdates.TLUserupdatesGetOperationResult {
	if !p.IsSetReq() {
		return GetOperationResultArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetOperationResultArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetOperationResultResult struct {
	Success *userupdates.UserOperationResult
}

var GetOperationResultResult_Success_DEFAULT *userupdates.UserOperationResult

func (p *GetOperationResultResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetOperationResultResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetOperationResultResult) Unmarshal(in []byte) error {
	msg := new(userupdates.UserOperationResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetOperationResultResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetOperationResultResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetOperationResultResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserOperationResult)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetOperationResultResult) GetSuccess() *userupdates.UserOperationResult {
	if !p.IsSetSuccess() {
		return GetOperationResultResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetOperationResultResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.UserOperationResult)
}


func (p *GetOperationResultResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetOperationResultResult) GetResult() interface{} {
	return p.Success
}


func getStateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetStateArgs)
	realResult := result.(*GetStateResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesGetState(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetStateArgs() interface{} {
	return &GetStateArgs{}
}

func newGetStateResult() interface{} {
	return &GetStateResult{}
}

type GetStateArgs struct {
	Req *userupdates.TLUserupdatesGetState
}

func (p *GetStateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetStateArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetStateArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesGetState)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetStateArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetStateArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetStateArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesGetState)
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

var GetStateArgs_Req_DEFAULT *userupdates.TLUserupdatesGetState

func (p *GetStateArgs) GetReq() *userupdates.TLUserupdatesGetState {
	if !p.IsSetReq() {
		return GetStateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetStateArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetStateResult struct {
	Success *userupdates.UserState
}

var GetStateResult_Success_DEFAULT *userupdates.UserState

func (p *GetStateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetStateResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetStateResult) Unmarshal(in []byte) error {
	msg := new(userupdates.UserState)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetStateResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetStateResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetStateResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserState)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetStateResult) GetSuccess() *userupdates.UserState {
	if !p.IsSetSuccess() {
		return GetStateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetStateResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.UserState)
}


func (p *GetStateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetStateResult) GetResult() interface{} {
	return p.Success
}


func getDifferenceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDifferenceArgs)
	realResult := result.(*GetDifferenceResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesGetDifference(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDifferenceArgs() interface{} {
	return &GetDifferenceArgs{}
}

func newGetDifferenceResult() interface{} {
	return &GetDifferenceResult{}
}

type GetDifferenceArgs struct {
	Req *userupdates.TLUserupdatesGetDifference
}

func (p *GetDifferenceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDifferenceArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDifferenceArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesGetDifference)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDifferenceArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDifferenceArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDifferenceArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesGetDifference)
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

var GetDifferenceArgs_Req_DEFAULT *userupdates.TLUserupdatesGetDifference

func (p *GetDifferenceArgs) GetReq() *userupdates.TLUserupdatesGetDifference {
	if !p.IsSetReq() {
		return GetDifferenceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDifferenceArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDifferenceResult struct {
	Success *userupdates.UserDifference
}

var GetDifferenceResult_Success_DEFAULT *userupdates.UserDifference

func (p *GetDifferenceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDifferenceResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDifferenceResult) Unmarshal(in []byte) error {
	msg := new(userupdates.UserDifference)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDifferenceResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDifferenceResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDifferenceResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserDifference)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDifferenceResult) GetSuccess() *userupdates.UserDifference {
	if !p.IsSetSuccess() {
		return GetDifferenceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDifferenceResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.UserDifference)
}


func (p *GetDifferenceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDifferenceResult) GetResult() interface{} {
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

func (p *kClient) UserupdatesProcessUserOperation(ctx context.Context, req *userupdates.TLUserupdatesProcessUserOperation) (r *userupdates.UserOperationResult, err error) {
	// var _args ProcessUserOperationArgs
	// _args.Req = req
	// var _result ProcessUserOperationResult

	_result := new(userupdates.UserOperationResult)

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.processUserOperation", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserupdatesGetOperationResult(ctx context.Context, req *userupdates.TLUserupdatesGetOperationResult) (r *userupdates.UserOperationResult, err error) {
	// var _args GetOperationResultArgs
	// _args.Req = req
	// var _result GetOperationResultResult

	_result := new(userupdates.UserOperationResult)

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getOperationResult", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserupdatesGetState(ctx context.Context, req *userupdates.TLUserupdatesGetState) (r *userupdates.UserState, err error) {
	// var _args GetStateArgs
	// _args.Req = req
	// var _result GetStateResult

	_result := new(userupdates.UserState)

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getState", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserupdatesGetDifference(ctx context.Context, req *userupdates.TLUserupdatesGetDifference) (r *userupdates.UserDifference, err error) {
	// var _args GetDifferenceArgs
	// _args.Req = req
	// var _result GetDifferenceResult

	_result := new(userupdates.UserDifference)

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getDifference", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}


