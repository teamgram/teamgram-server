/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package syncservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/sync.RPCSync/sync.updatesMe": kitex.NewMethodInfo(
		updatesMeHandler,
		newUpdatesMeArgs,
		newUpdatesMeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/sync.RPCSync/sync.updatesNotMe": kitex.NewMethodInfo(
		updatesNotMeHandler,
		newUpdatesNotMeArgs,
		newUpdatesNotMeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/sync.RPCSync/sync.pushUpdates": kitex.NewMethodInfo(
		pushUpdatesHandler,
		newPushUpdatesArgs,
		newPushUpdatesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/sync.RPCSync/sync.pushUpdatesIfNot": kitex.NewMethodInfo(
		pushUpdatesIfNotHandler,
		newPushUpdatesIfNotArgs,
		newPushUpdatesIfNotResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/sync.RPCSync/sync.pushBotUpdates": kitex.NewMethodInfo(
		pushBotUpdatesHandler,
		newPushBotUpdatesArgs,
		newPushBotUpdatesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/sync.RPCSync/sync.pushRpcResult": kitex.NewMethodInfo(
		pushRpcResultHandler,
		newPushRpcResultArgs,
		newPushRpcResultResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/sync.RPCSync/sync.broadcastUpdates": kitex.NewMethodInfo(
		broadcastUpdatesHandler,
		newBroadcastUpdatesArgs,
		newBroadcastUpdatesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	syncServiceServiceInfo                = NewServiceInfo()
	syncServiceServiceInfoForClient       = NewServiceInfoForClient()
	syncServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCSync", syncServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCSync", syncServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCSync", syncServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return syncServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return syncServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return syncServiceServiceInfoForClient
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
	serviceName := "RPCSync"
	handlerType := (*sync.RPCSync)(nil)
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
		"PackageName": "sync",
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

func updatesMeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatesMeArgs)
	realResult := result.(*UpdatesMeResult)
	success, err := handler.(sync.RPCSync).SyncUpdatesMe(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatesMeArgs() interface{} {
	return &UpdatesMeArgs{}
}

func newUpdatesMeResult() interface{} {
	return &UpdatesMeResult{}
}

type UpdatesMeArgs struct {
	Req *sync.TLSyncUpdatesMe
}

func (p *UpdatesMeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatesMeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatesMeArgs) Unmarshal(in []byte) error {
	msg := new(sync.TLSyncUpdatesMe)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatesMeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatesMeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatesMeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(sync.TLSyncUpdatesMe)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatesMeArgs_Req_DEFAULT *sync.TLSyncUpdatesMe

func (p *UpdatesMeArgs) GetReq() *sync.TLSyncUpdatesMe {
	if !p.IsSetReq() {
		return UpdatesMeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatesMeArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatesMeResult struct {
	Success *tg.Void
}

var UpdatesMeResult_Success_DEFAULT *tg.Void

func (p *UpdatesMeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatesMeResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatesMeResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesMeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatesMeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatesMeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesMeResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return UpdatesMeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatesMeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *UpdatesMeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatesMeResult) GetResult() interface{} {
	return p.Success
}

func updatesNotMeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatesNotMeArgs)
	realResult := result.(*UpdatesNotMeResult)
	success, err := handler.(sync.RPCSync).SyncUpdatesNotMe(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatesNotMeArgs() interface{} {
	return &UpdatesNotMeArgs{}
}

func newUpdatesNotMeResult() interface{} {
	return &UpdatesNotMeResult{}
}

type UpdatesNotMeArgs struct {
	Req *sync.TLSyncUpdatesNotMe
}

func (p *UpdatesNotMeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatesNotMeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatesNotMeArgs) Unmarshal(in []byte) error {
	msg := new(sync.TLSyncUpdatesNotMe)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatesNotMeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatesNotMeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatesNotMeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(sync.TLSyncUpdatesNotMe)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatesNotMeArgs_Req_DEFAULT *sync.TLSyncUpdatesNotMe

func (p *UpdatesNotMeArgs) GetReq() *sync.TLSyncUpdatesNotMe {
	if !p.IsSetReq() {
		return UpdatesNotMeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatesNotMeArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatesNotMeResult struct {
	Success *tg.Void
}

var UpdatesNotMeResult_Success_DEFAULT *tg.Void

func (p *UpdatesNotMeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatesNotMeResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatesNotMeResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesNotMeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatesNotMeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatesNotMeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatesNotMeResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return UpdatesNotMeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatesNotMeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *UpdatesNotMeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatesNotMeResult) GetResult() interface{} {
	return p.Success
}

func pushUpdatesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushUpdatesArgs)
	realResult := result.(*PushUpdatesResult)
	success, err := handler.(sync.RPCSync).SyncPushUpdates(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPushUpdatesArgs() interface{} {
	return &PushUpdatesArgs{}
}

func newPushUpdatesResult() interface{} {
	return &PushUpdatesResult{}
}

type PushUpdatesArgs struct {
	Req *sync.TLSyncPushUpdates
}

func (p *PushUpdatesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushUpdatesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushUpdatesArgs) Unmarshal(in []byte) error {
	msg := new(sync.TLSyncPushUpdates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PushUpdatesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PushUpdatesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PushUpdatesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(sync.TLSyncPushUpdates)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PushUpdatesArgs_Req_DEFAULT *sync.TLSyncPushUpdates

func (p *PushUpdatesArgs) GetReq() *sync.TLSyncPushUpdates {
	if !p.IsSetReq() {
		return PushUpdatesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushUpdatesArgs) IsSetReq() bool {
	return p.Req != nil
}

type PushUpdatesResult struct {
	Success *tg.Void
}

var PushUpdatesResult_Success_DEFAULT *tg.Void

func (p *PushUpdatesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PushUpdatesResult")
	}
	return json.Marshal(p.Success)
}

func (p *PushUpdatesResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushUpdatesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PushUpdatesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PushUpdatesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushUpdatesResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return PushUpdatesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushUpdatesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *PushUpdatesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushUpdatesResult) GetResult() interface{} {
	return p.Success
}

func pushUpdatesIfNotHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushUpdatesIfNotArgs)
	realResult := result.(*PushUpdatesIfNotResult)
	success, err := handler.(sync.RPCSync).SyncPushUpdatesIfNot(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPushUpdatesIfNotArgs() interface{} {
	return &PushUpdatesIfNotArgs{}
}

func newPushUpdatesIfNotResult() interface{} {
	return &PushUpdatesIfNotResult{}
}

type PushUpdatesIfNotArgs struct {
	Req *sync.TLSyncPushUpdatesIfNot
}

func (p *PushUpdatesIfNotArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushUpdatesIfNotArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushUpdatesIfNotArgs) Unmarshal(in []byte) error {
	msg := new(sync.TLSyncPushUpdatesIfNot)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PushUpdatesIfNotArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PushUpdatesIfNotArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PushUpdatesIfNotArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(sync.TLSyncPushUpdatesIfNot)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PushUpdatesIfNotArgs_Req_DEFAULT *sync.TLSyncPushUpdatesIfNot

func (p *PushUpdatesIfNotArgs) GetReq() *sync.TLSyncPushUpdatesIfNot {
	if !p.IsSetReq() {
		return PushUpdatesIfNotArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushUpdatesIfNotArgs) IsSetReq() bool {
	return p.Req != nil
}

type PushUpdatesIfNotResult struct {
	Success *tg.Void
}

var PushUpdatesIfNotResult_Success_DEFAULT *tg.Void

func (p *PushUpdatesIfNotResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PushUpdatesIfNotResult")
	}
	return json.Marshal(p.Success)
}

func (p *PushUpdatesIfNotResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushUpdatesIfNotResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PushUpdatesIfNotResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PushUpdatesIfNotResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushUpdatesIfNotResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return PushUpdatesIfNotResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushUpdatesIfNotResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *PushUpdatesIfNotResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushUpdatesIfNotResult) GetResult() interface{} {
	return p.Success
}

func pushBotUpdatesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushBotUpdatesArgs)
	realResult := result.(*PushBotUpdatesResult)
	success, err := handler.(sync.RPCSync).SyncPushBotUpdates(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPushBotUpdatesArgs() interface{} {
	return &PushBotUpdatesArgs{}
}

func newPushBotUpdatesResult() interface{} {
	return &PushBotUpdatesResult{}
}

type PushBotUpdatesArgs struct {
	Req *sync.TLSyncPushBotUpdates
}

func (p *PushBotUpdatesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushBotUpdatesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushBotUpdatesArgs) Unmarshal(in []byte) error {
	msg := new(sync.TLSyncPushBotUpdates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PushBotUpdatesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PushBotUpdatesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PushBotUpdatesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(sync.TLSyncPushBotUpdates)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PushBotUpdatesArgs_Req_DEFAULT *sync.TLSyncPushBotUpdates

func (p *PushBotUpdatesArgs) GetReq() *sync.TLSyncPushBotUpdates {
	if !p.IsSetReq() {
		return PushBotUpdatesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushBotUpdatesArgs) IsSetReq() bool {
	return p.Req != nil
}

type PushBotUpdatesResult struct {
	Success *tg.Void
}

var PushBotUpdatesResult_Success_DEFAULT *tg.Void

func (p *PushBotUpdatesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PushBotUpdatesResult")
	}
	return json.Marshal(p.Success)
}

func (p *PushBotUpdatesResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushBotUpdatesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PushBotUpdatesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PushBotUpdatesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushBotUpdatesResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return PushBotUpdatesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushBotUpdatesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *PushBotUpdatesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushBotUpdatesResult) GetResult() interface{} {
	return p.Success
}

func pushRpcResultHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushRpcResultArgs)
	realResult := result.(*PushRpcResultResult)
	success, err := handler.(sync.RPCSync).SyncPushRpcResult(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPushRpcResultArgs() interface{} {
	return &PushRpcResultArgs{}
}

func newPushRpcResultResult() interface{} {
	return &PushRpcResultResult{}
}

type PushRpcResultArgs struct {
	Req *sync.TLSyncPushRpcResult
}

func (p *PushRpcResultArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushRpcResultArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushRpcResultArgs) Unmarshal(in []byte) error {
	msg := new(sync.TLSyncPushRpcResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PushRpcResultArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PushRpcResultArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PushRpcResultArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(sync.TLSyncPushRpcResult)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PushRpcResultArgs_Req_DEFAULT *sync.TLSyncPushRpcResult

func (p *PushRpcResultArgs) GetReq() *sync.TLSyncPushRpcResult {
	if !p.IsSetReq() {
		return PushRpcResultArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushRpcResultArgs) IsSetReq() bool {
	return p.Req != nil
}

type PushRpcResultResult struct {
	Success *tg.Void
}

var PushRpcResultResult_Success_DEFAULT *tg.Void

func (p *PushRpcResultResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PushRpcResultResult")
	}
	return json.Marshal(p.Success)
}

func (p *PushRpcResultResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushRpcResultResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PushRpcResultResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PushRpcResultResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushRpcResultResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return PushRpcResultResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushRpcResultResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *PushRpcResultResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushRpcResultResult) GetResult() interface{} {
	return p.Success
}

func broadcastUpdatesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*BroadcastUpdatesArgs)
	realResult := result.(*BroadcastUpdatesResult)
	success, err := handler.(sync.RPCSync).SyncBroadcastUpdates(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newBroadcastUpdatesArgs() interface{} {
	return &BroadcastUpdatesArgs{}
}

func newBroadcastUpdatesResult() interface{} {
	return &BroadcastUpdatesResult{}
}

type BroadcastUpdatesArgs struct {
	Req *sync.TLSyncBroadcastUpdates
}

func (p *BroadcastUpdatesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in BroadcastUpdatesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *BroadcastUpdatesArgs) Unmarshal(in []byte) error {
	msg := new(sync.TLSyncBroadcastUpdates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *BroadcastUpdatesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in BroadcastUpdatesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *BroadcastUpdatesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(sync.TLSyncBroadcastUpdates)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var BroadcastUpdatesArgs_Req_DEFAULT *sync.TLSyncBroadcastUpdates

func (p *BroadcastUpdatesArgs) GetReq() *sync.TLSyncBroadcastUpdates {
	if !p.IsSetReq() {
		return BroadcastUpdatesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *BroadcastUpdatesArgs) IsSetReq() bool {
	return p.Req != nil
}

type BroadcastUpdatesResult struct {
	Success *tg.Void
}

var BroadcastUpdatesResult_Success_DEFAULT *tg.Void

func (p *BroadcastUpdatesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in BroadcastUpdatesResult")
	}
	return json.Marshal(p.Success)
}

func (p *BroadcastUpdatesResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BroadcastUpdatesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in BroadcastUpdatesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *BroadcastUpdatesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BroadcastUpdatesResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return BroadcastUpdatesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *BroadcastUpdatesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *BroadcastUpdatesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BroadcastUpdatesResult) GetResult() interface{} {
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

func (p *kClient) SyncUpdatesMe(ctx context.Context, req *sync.TLSyncUpdatesMe) (r *tg.Void, err error) {
	// var _args UpdatesMeArgs
	// _args.Req = req
	// var _result UpdatesMeResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/sync.RPCSync/sync.updatesMe", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SyncUpdatesNotMe(ctx context.Context, req *sync.TLSyncUpdatesNotMe) (r *tg.Void, err error) {
	// var _args UpdatesNotMeArgs
	// _args.Req = req
	// var _result UpdatesNotMeResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/sync.RPCSync/sync.updatesNotMe", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SyncPushUpdates(ctx context.Context, req *sync.TLSyncPushUpdates) (r *tg.Void, err error) {
	// var _args PushUpdatesArgs
	// _args.Req = req
	// var _result PushUpdatesResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/sync.RPCSync/sync.pushUpdates", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SyncPushUpdatesIfNot(ctx context.Context, req *sync.TLSyncPushUpdatesIfNot) (r *tg.Void, err error) {
	// var _args PushUpdatesIfNotArgs
	// _args.Req = req
	// var _result PushUpdatesIfNotResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/sync.RPCSync/sync.pushUpdatesIfNot", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SyncPushBotUpdates(ctx context.Context, req *sync.TLSyncPushBotUpdates) (r *tg.Void, err error) {
	// var _args PushBotUpdatesArgs
	// _args.Req = req
	// var _result PushBotUpdatesResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/sync.RPCSync/sync.pushBotUpdates", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SyncPushRpcResult(ctx context.Context, req *sync.TLSyncPushRpcResult) (r *tg.Void, err error) {
	// var _args PushRpcResultArgs
	// _args.Req = req
	// var _result PushRpcResultResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/sync.RPCSync/sync.pushRpcResult", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SyncBroadcastUpdates(ctx context.Context, req *sync.TLSyncBroadcastUpdates) (r *tg.Void, err error) {
	// var _args BroadcastUpdatesArgs
	// _args.Req = req
	// var _result BroadcastUpdatesResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/sync.RPCSync/sync.broadcastUpdates", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
