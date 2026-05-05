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
	"errors"
	"fmt"
	"reflect"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

func decodeConstructorIfPresent(d *bin.Decoder, msg interface{}) error {
	v := reflect.ValueOf(msg)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}

	f := v.FieldByName("ClazzID")
	if !f.IsValid() || !f.CanSet() || f.Kind() != reflect.Uint32 {
		return nil
	}

	clazzID, err := d.ClazzID()
	if err != nil {
		return err
	}
	f.SetUint(uint64(clazzID))
	return nil
}

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
	"/userupdates.RPCUserupdates/userupdates.listDialogs": kitex.NewMethodInfo(
		listDialogsHandler,
		newListDialogsArgs,
		newListDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/userupdates.RPCUserupdates/userupdates.getDialogsByPeers": kitex.NewMethodInfo(
		getDialogsByPeersHandler,
		newGetDialogsByPeersArgs,
		newGetDialogsByPeersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/userupdates.RPCUserupdates/userupdates.getDialogCount": kitex.NewMethodInfo(
		getDialogCountHandler,
		newGetDialogCountArgs,
		newGetDialogCountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/userupdates.RPCUserupdates/userupdates.getMessageViewsByPeerSeqs": kitex.NewMethodInfo(
		getMessageViewsByPeerSeqsHandler,
		newGetMessageViewsByPeerSeqsArgs,
		newGetMessageViewsByPeerSeqsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/userupdates.RPCUserupdates/userupdates.appendDialogAuthSeqSideEffect": kitex.NewMethodInfo(
		appendDialogAuthSeqSideEffectHandler,
		newAppendDialogAuthSeqSideEffectArgs,
		newAppendDialogAuthSeqSideEffectResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/userupdates.RPCUserupdates/userupdates.appendDialogPtsSideEffect": kitex.NewMethodInfo(
		appendDialogPtsSideEffectHandler,
		newAppendDialogPtsSideEffectArgs,
		newAppendDialogPtsSideEffectResult,
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
		return out, fmt.Errorf("no req in ProcessUserOperationArgs")
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
		return fmt.Errorf("no req in ProcessUserOperationArgs")
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
		return out, fmt.Errorf("no req in ProcessUserOperationResult")
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
		return fmt.Errorf("no req in ProcessUserOperationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ProcessUserOperationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserOperationResult)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
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
		return out, fmt.Errorf("no req in GetOperationResultArgs")
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
		return fmt.Errorf("no req in GetOperationResultArgs")
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
		return out, fmt.Errorf("no req in GetOperationResultResult")
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
		return fmt.Errorf("no req in GetOperationResultResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetOperationResultResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserOperationResult)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
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
		return out, fmt.Errorf("no req in GetStateArgs")
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
		return fmt.Errorf("no req in GetStateArgs")
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
		return out, fmt.Errorf("no req in GetStateResult")
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
		return fmt.Errorf("no req in GetStateResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetStateResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserState)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
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
		return out, fmt.Errorf("no req in GetDifferenceArgs")
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
		return fmt.Errorf("no req in GetDifferenceArgs")
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
		return out, fmt.Errorf("no req in GetDifferenceResult")
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
		return fmt.Errorf("no req in GetDifferenceResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDifferenceResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserDifference)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
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

func listDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ListDialogsArgs)
	realResult := result.(*ListDialogsResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesListDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newListDialogsArgs() interface{} {
	return &ListDialogsArgs{}
}

func newListDialogsResult() interface{} {
	return &ListDialogsResult{}
}

type ListDialogsArgs struct {
	Req *userupdates.TLUserupdatesListDialogs
}

func (p *ListDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in ListDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ListDialogsArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesListDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ListDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in ListDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ListDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesListDialogs)
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

var ListDialogsArgs_Req_DEFAULT *userupdates.TLUserupdatesListDialogs

func (p *ListDialogsArgs) GetReq() *userupdates.TLUserupdatesListDialogs {
	if !p.IsSetReq() {
		return ListDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ListDialogsResult struct {
	Success *userupdates.DialogProjectionList
}

var ListDialogsResult_Success_DEFAULT *userupdates.DialogProjectionList

func (p *ListDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in ListDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ListDialogsResult) Unmarshal(in []byte) error {
	msg := new(userupdates.DialogProjectionList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in ListDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ListDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.DialogProjectionList)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListDialogsResult) GetSuccess() *userupdates.DialogProjectionList {
	if !p.IsSetSuccess() {
		return ListDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.DialogProjectionList)
}

func (p *ListDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListDialogsResult) GetResult() interface{} {
	return p.Success
}

func getDialogsByPeersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogsByPeersArgs)
	realResult := result.(*GetDialogsByPeersResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesGetDialogsByPeers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogsByPeersArgs() interface{} {
	return &GetDialogsByPeersArgs{}
}

func newGetDialogsByPeersResult() interface{} {
	return &GetDialogsByPeersResult{}
}

type GetDialogsByPeersArgs struct {
	Req *userupdates.TLUserupdatesGetDialogsByPeers
}

func (p *GetDialogsByPeersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in GetDialogsByPeersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogsByPeersArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesGetDialogsByPeers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogsByPeersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in GetDialogsByPeersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogsByPeersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesGetDialogsByPeers)
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

var GetDialogsByPeersArgs_Req_DEFAULT *userupdates.TLUserupdatesGetDialogsByPeers

func (p *GetDialogsByPeersArgs) GetReq() *userupdates.TLUserupdatesGetDialogsByPeers {
	if !p.IsSetReq() {
		return GetDialogsByPeersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogsByPeersArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogsByPeersResult struct {
	Success *userupdates.VectorDialogProjection
}

var GetDialogsByPeersResult_Success_DEFAULT *userupdates.VectorDialogProjection

func (p *GetDialogsByPeersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in GetDialogsByPeersResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogsByPeersResult) Unmarshal(in []byte) error {
	msg := new(userupdates.VectorDialogProjection)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsByPeersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in GetDialogsByPeersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogsByPeersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.VectorDialogProjection)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsByPeersResult) GetSuccess() *userupdates.VectorDialogProjection {
	if !p.IsSetSuccess() {
		return GetDialogsByPeersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogsByPeersResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.VectorDialogProjection)
}

func (p *GetDialogsByPeersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogsByPeersResult) GetResult() interface{} {
	return p.Success
}

func getDialogCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogCountArgs)
	realResult := result.(*GetDialogCountResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesGetDialogCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogCountArgs() interface{} {
	return &GetDialogCountArgs{}
}

func newGetDialogCountResult() interface{} {
	return &GetDialogCountResult{}
}

type GetDialogCountArgs struct {
	Req *userupdates.TLUserupdatesGetDialogCount
}

func (p *GetDialogCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in GetDialogCountArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogCountArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesGetDialogCount)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogCountArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in GetDialogCountArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogCountArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesGetDialogCount)
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

var GetDialogCountArgs_Req_DEFAULT *userupdates.TLUserupdatesGetDialogCount

func (p *GetDialogCountArgs) GetReq() *userupdates.TLUserupdatesGetDialogCount {
	if !p.IsSetReq() {
		return GetDialogCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogCountResult struct {
	Success *tg.Int32
}

var GetDialogCountResult_Success_DEFAULT *tg.Int32

func (p *GetDialogCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in GetDialogCountResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogCountResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogCountResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in GetDialogCountResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogCountResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogCountResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetDialogCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetDialogCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogCountResult) GetResult() interface{} {
	return p.Success
}

func getMessageViewsByPeerSeqsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetMessageViewsByPeerSeqsArgs)
	realResult := result.(*GetMessageViewsByPeerSeqsResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesGetMessageViewsByPeerSeqs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetMessageViewsByPeerSeqsArgs() interface{} {
	return &GetMessageViewsByPeerSeqsArgs{}
}

func newGetMessageViewsByPeerSeqsResult() interface{} {
	return &GetMessageViewsByPeerSeqsResult{}
}

type GetMessageViewsByPeerSeqsArgs struct {
	Req *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs
}

func (p *GetMessageViewsByPeerSeqsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in GetMessageViewsByPeerSeqsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetMessageViewsByPeerSeqsArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesGetMessageViewsByPeerSeqs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetMessageViewsByPeerSeqsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in GetMessageViewsByPeerSeqsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetMessageViewsByPeerSeqsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesGetMessageViewsByPeerSeqs)
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

var GetMessageViewsByPeerSeqsArgs_Req_DEFAULT *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs

func (p *GetMessageViewsByPeerSeqsArgs) GetReq() *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs {
	if !p.IsSetReq() {
		return GetMessageViewsByPeerSeqsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMessageViewsByPeerSeqsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetMessageViewsByPeerSeqsResult struct {
	Success *userupdates.MessageViewList
}

var GetMessageViewsByPeerSeqsResult_Success_DEFAULT *userupdates.MessageViewList

func (p *GetMessageViewsByPeerSeqsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in GetMessageViewsByPeerSeqsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetMessageViewsByPeerSeqsResult) Unmarshal(in []byte) error {
	msg := new(userupdates.MessageViewList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMessageViewsByPeerSeqsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in GetMessageViewsByPeerSeqsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetMessageViewsByPeerSeqsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.MessageViewList)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMessageViewsByPeerSeqsResult) GetSuccess() *userupdates.MessageViewList {
	if !p.IsSetSuccess() {
		return GetMessageViewsByPeerSeqsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMessageViewsByPeerSeqsResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.MessageViewList)
}

func (p *GetMessageViewsByPeerSeqsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMessageViewsByPeerSeqsResult) GetResult() interface{} {
	return p.Success
}

func appendDialogAuthSeqSideEffectHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AppendDialogAuthSeqSideEffectArgs)
	realResult := result.(*AppendDialogAuthSeqSideEffectResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesAppendDialogAuthSeqSideEffect(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAppendDialogAuthSeqSideEffectArgs() interface{} {
	return &AppendDialogAuthSeqSideEffectArgs{}
}

func newAppendDialogAuthSeqSideEffectResult() interface{} {
	return &AppendDialogAuthSeqSideEffectResult{}
}

type AppendDialogAuthSeqSideEffectArgs struct {
	Req *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect
}

func (p *AppendDialogAuthSeqSideEffectArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in AppendDialogAuthSeqSideEffectArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AppendDialogAuthSeqSideEffectArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AppendDialogAuthSeqSideEffectArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in AppendDialogAuthSeqSideEffectArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AppendDialogAuthSeqSideEffectArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect)
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

var AppendDialogAuthSeqSideEffectArgs_Req_DEFAULT *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect

func (p *AppendDialogAuthSeqSideEffectArgs) GetReq() *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect {
	if !p.IsSetReq() {
		return AppendDialogAuthSeqSideEffectArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AppendDialogAuthSeqSideEffectArgs) IsSetReq() bool {
	return p.Req != nil
}

type AppendDialogAuthSeqSideEffectResult struct {
	Success *userupdates.UserAuthSeqAppendResult
}

var AppendDialogAuthSeqSideEffectResult_Success_DEFAULT *userupdates.UserAuthSeqAppendResult

func (p *AppendDialogAuthSeqSideEffectResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in AppendDialogAuthSeqSideEffectResult")
	}
	return json.Marshal(p.Success)
}

func (p *AppendDialogAuthSeqSideEffectResult) Unmarshal(in []byte) error {
	msg := new(userupdates.UserAuthSeqAppendResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AppendDialogAuthSeqSideEffectResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in AppendDialogAuthSeqSideEffectResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AppendDialogAuthSeqSideEffectResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserAuthSeqAppendResult)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AppendDialogAuthSeqSideEffectResult) GetSuccess() *userupdates.UserAuthSeqAppendResult {
	if !p.IsSetSuccess() {
		return AppendDialogAuthSeqSideEffectResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AppendDialogAuthSeqSideEffectResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.UserAuthSeqAppendResult)
}

func (p *AppendDialogAuthSeqSideEffectResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AppendDialogAuthSeqSideEffectResult) GetResult() interface{} {
	return p.Success
}

func appendDialogPtsSideEffectHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AppendDialogPtsSideEffectArgs)
	realResult := result.(*AppendDialogPtsSideEffectResult)
	success, err := handler.(userupdates.RPCUserupdates).UserupdatesAppendDialogPtsSideEffect(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAppendDialogPtsSideEffectArgs() interface{} {
	return &AppendDialogPtsSideEffectArgs{}
}

func newAppendDialogPtsSideEffectResult() interface{} {
	return &AppendDialogPtsSideEffectResult{}
}

type AppendDialogPtsSideEffectArgs struct {
	Req *userupdates.TLUserupdatesAppendDialogPtsSideEffect
}

func (p *AppendDialogPtsSideEffectArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in AppendDialogPtsSideEffectArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AppendDialogPtsSideEffectArgs) Unmarshal(in []byte) error {
	msg := new(userupdates.TLUserupdatesAppendDialogPtsSideEffect)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AppendDialogPtsSideEffectArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in AppendDialogPtsSideEffectArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AppendDialogPtsSideEffectArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.TLUserupdatesAppendDialogPtsSideEffect)
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

var AppendDialogPtsSideEffectArgs_Req_DEFAULT *userupdates.TLUserupdatesAppendDialogPtsSideEffect

func (p *AppendDialogPtsSideEffectArgs) GetReq() *userupdates.TLUserupdatesAppendDialogPtsSideEffect {
	if !p.IsSetReq() {
		return AppendDialogPtsSideEffectArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AppendDialogPtsSideEffectArgs) IsSetReq() bool {
	return p.Req != nil
}

type AppendDialogPtsSideEffectResult struct {
	Success *userupdates.UserPtsAppendResult
}

var AppendDialogPtsSideEffectResult_Success_DEFAULT *userupdates.UserPtsAppendResult

func (p *AppendDialogPtsSideEffectResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in AppendDialogPtsSideEffectResult")
	}
	return json.Marshal(p.Success)
}

func (p *AppendDialogPtsSideEffectResult) Unmarshal(in []byte) error {
	msg := new(userupdates.UserPtsAppendResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AppendDialogPtsSideEffectResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in AppendDialogPtsSideEffectResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AppendDialogPtsSideEffectResult) Decode(d *bin.Decoder) (err error) {
	msg := new(userupdates.UserPtsAppendResult)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AppendDialogPtsSideEffectResult) GetSuccess() *userupdates.UserPtsAppendResult {
	if !p.IsSetSuccess() {
		return AppendDialogPtsSideEffectResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AppendDialogPtsSideEffectResult) SetSuccess(x interface{}) {
	p.Success = x.(*userupdates.UserPtsAppendResult)
}

func (p *AppendDialogPtsSideEffectResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AppendDialogPtsSideEffectResult) GetResult() interface{} {
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
	var _result ProcessUserOperationResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.processUserOperation", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesGetOperationResult(ctx context.Context, req *userupdates.TLUserupdatesGetOperationResult) (r *userupdates.UserOperationResult, err error) {
	// var _args GetOperationResultArgs
	// _args.Req = req
	var _result GetOperationResultResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getOperationResult", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesGetState(ctx context.Context, req *userupdates.TLUserupdatesGetState) (r *userupdates.UserState, err error) {
	// var _args GetStateArgs
	// _args.Req = req
	var _result GetStateResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getState", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesGetDifference(ctx context.Context, req *userupdates.TLUserupdatesGetDifference) (r *userupdates.UserDifference, err error) {
	// var _args GetDifferenceArgs
	// _args.Req = req
	var _result GetDifferenceResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getDifference", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesListDialogs(ctx context.Context, req *userupdates.TLUserupdatesListDialogs) (r *userupdates.DialogProjectionList, err error) {
	// var _args ListDialogsArgs
	// _args.Req = req
	var _result ListDialogsResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.listDialogs", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesGetDialogsByPeers(ctx context.Context, req *userupdates.TLUserupdatesGetDialogsByPeers) (r *userupdates.VectorDialogProjection, err error) {
	// var _args GetDialogsByPeersArgs
	// _args.Req = req
	var _result GetDialogsByPeersResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getDialogsByPeers", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesGetDialogCount(ctx context.Context, req *userupdates.TLUserupdatesGetDialogCount) (r *tg.Int32, err error) {
	// var _args GetDialogCountArgs
	// _args.Req = req
	var _result GetDialogCountResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getDialogCount", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesGetMessageViewsByPeerSeqs(ctx context.Context, req *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (r *userupdates.MessageViewList, err error) {
	// var _args GetMessageViewsByPeerSeqsArgs
	// _args.Req = req
	var _result GetMessageViewsByPeerSeqsResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.getMessageViewsByPeerSeqs", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, req *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect) (r *userupdates.UserAuthSeqAppendResult, err error) {
	// var _args AppendDialogAuthSeqSideEffectArgs
	// _args.Req = req
	var _result AppendDialogAuthSeqSideEffectResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.appendDialogAuthSeqSideEffect", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) UserupdatesAppendDialogPtsSideEffect(ctx context.Context, req *userupdates.TLUserupdatesAppendDialogPtsSideEffect) (r *userupdates.UserPtsAppendResult, err error) {
	// var _args AppendDialogPtsSideEffectArgs
	// _args.Req = req
	var _result AppendDialogPtsSideEffectResult

	if err = p.c.Call(ctx, "/userupdates.RPCUserupdates/userupdates.appendDialogPtsSideEffect", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}
