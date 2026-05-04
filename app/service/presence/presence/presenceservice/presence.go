/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package presenceservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
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
	"/presence.RPCPresence/presence.setSessionOnline": kitex.NewMethodInfo(
		setSessionOnlineHandler,
		newSetSessionOnlineArgs,
		newSetSessionOnlineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/presence.RPCPresence/presence.setSessionOffline": kitex.NewMethodInfo(
		setSessionOfflineHandler,
		newSetSessionOfflineArgs,
		newSetSessionOfflineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/presence.RPCPresence/presence.getUserOnlineSessions": kitex.NewMethodInfo(
		getUserOnlineSessionsHandler,
		newGetUserOnlineSessionsArgs,
		newGetUserOnlineSessionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/presence.RPCPresence/presence.getUsersOnlineSessions": kitex.NewMethodInfo(
		getUsersOnlineSessionsHandler,
		newGetUsersOnlineSessionsArgs,
		newGetUsersOnlineSessionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/presence.RPCPresence/presence.getGatewaySessions": kitex.NewMethodInfo(
		getGatewaySessionsHandler,
		newGetGatewaySessionsArgs,
		newGetGatewaySessionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	presenceServiceServiceInfo                = NewServiceInfo()
	presenceServiceServiceInfoForClient       = NewServiceInfoForClient()
	presenceServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCPresence", presenceServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCPresence", presenceServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCPresence", presenceServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return presenceServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return presenceServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return presenceServiceServiceInfoForClient
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
	serviceName := "RPCPresence"
	handlerType := (*presence.RPCPresence)(nil)
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
		"PackageName": "presence",
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

func setSessionOnlineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetSessionOnlineArgs)
	realResult := result.(*SetSessionOnlineResult)
	success, err := handler.(presence.RPCPresence).PresenceSetSessionOnline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetSessionOnlineArgs() interface{} {
	return &SetSessionOnlineArgs{}
}

func newSetSessionOnlineResult() interface{} {
	return &SetSessionOnlineResult{}
}

type SetSessionOnlineArgs struct {
	Req *presence.TLPresenceSetSessionOnline
}

func (p *SetSessionOnlineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in SetSessionOnlineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetSessionOnlineArgs) Unmarshal(in []byte) error {
	msg := new(presence.TLPresenceSetSessionOnline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetSessionOnlineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in SetSessionOnlineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetSessionOnlineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(presence.TLPresenceSetSessionOnline)
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

var SetSessionOnlineArgs_Req_DEFAULT *presence.TLPresenceSetSessionOnline

func (p *SetSessionOnlineArgs) GetReq() *presence.TLPresenceSetSessionOnline {
	if !p.IsSetReq() {
		return SetSessionOnlineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetSessionOnlineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetSessionOnlineResult struct {
	Success *tg.Bool
}

var SetSessionOnlineResult_Success_DEFAULT *tg.Bool

func (p *SetSessionOnlineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in SetSessionOnlineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetSessionOnlineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetSessionOnlineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in SetSessionOnlineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetSessionOnlineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetSessionOnlineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetSessionOnlineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetSessionOnlineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetSessionOnlineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetSessionOnlineResult) GetResult() interface{} {
	return p.Success
}

func setSessionOfflineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetSessionOfflineArgs)
	realResult := result.(*SetSessionOfflineResult)
	success, err := handler.(presence.RPCPresence).PresenceSetSessionOffline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetSessionOfflineArgs() interface{} {
	return &SetSessionOfflineArgs{}
}

func newSetSessionOfflineResult() interface{} {
	return &SetSessionOfflineResult{}
}

type SetSessionOfflineArgs struct {
	Req *presence.TLPresenceSetSessionOffline
}

func (p *SetSessionOfflineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in SetSessionOfflineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetSessionOfflineArgs) Unmarshal(in []byte) error {
	msg := new(presence.TLPresenceSetSessionOffline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetSessionOfflineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in SetSessionOfflineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetSessionOfflineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(presence.TLPresenceSetSessionOffline)
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

var SetSessionOfflineArgs_Req_DEFAULT *presence.TLPresenceSetSessionOffline

func (p *SetSessionOfflineArgs) GetReq() *presence.TLPresenceSetSessionOffline {
	if !p.IsSetReq() {
		return SetSessionOfflineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetSessionOfflineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetSessionOfflineResult struct {
	Success *tg.Bool
}

var SetSessionOfflineResult_Success_DEFAULT *tg.Bool

func (p *SetSessionOfflineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in SetSessionOfflineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetSessionOfflineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetSessionOfflineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in SetSessionOfflineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetSessionOfflineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetSessionOfflineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetSessionOfflineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetSessionOfflineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetSessionOfflineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetSessionOfflineResult) GetResult() interface{} {
	return p.Success
}

func getUserOnlineSessionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserOnlineSessionsArgs)
	realResult := result.(*GetUserOnlineSessionsResult)
	success, err := handler.(presence.RPCPresence).PresenceGetUserOnlineSessions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserOnlineSessionsArgs() interface{} {
	return &GetUserOnlineSessionsArgs{}
}

func newGetUserOnlineSessionsResult() interface{} {
	return &GetUserOnlineSessionsResult{}
}

type GetUserOnlineSessionsArgs struct {
	Req *presence.TLPresenceGetUserOnlineSessions
}

func (p *GetUserOnlineSessionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in GetUserOnlineSessionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserOnlineSessionsArgs) Unmarshal(in []byte) error {
	msg := new(presence.TLPresenceGetUserOnlineSessions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserOnlineSessionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in GetUserOnlineSessionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserOnlineSessionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(presence.TLPresenceGetUserOnlineSessions)
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

var GetUserOnlineSessionsArgs_Req_DEFAULT *presence.TLPresenceGetUserOnlineSessions

func (p *GetUserOnlineSessionsArgs) GetReq() *presence.TLPresenceGetUserOnlineSessions {
	if !p.IsSetReq() {
		return GetUserOnlineSessionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserOnlineSessionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserOnlineSessionsResult struct {
	Success *presence.UserOnlineSessions
}

var GetUserOnlineSessionsResult_Success_DEFAULT *presence.UserOnlineSessions

func (p *GetUserOnlineSessionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in GetUserOnlineSessionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserOnlineSessionsResult) Unmarshal(in []byte) error {
	msg := new(presence.UserOnlineSessions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserOnlineSessionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in GetUserOnlineSessionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserOnlineSessionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(presence.UserOnlineSessions)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserOnlineSessionsResult) GetSuccess() *presence.UserOnlineSessions {
	if !p.IsSetSuccess() {
		return GetUserOnlineSessionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserOnlineSessionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*presence.UserOnlineSessions)
}

func (p *GetUserOnlineSessionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserOnlineSessionsResult) GetResult() interface{} {
	return p.Success
}

func getUsersOnlineSessionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUsersOnlineSessionsArgs)
	realResult := result.(*GetUsersOnlineSessionsResult)
	success, err := handler.(presence.RPCPresence).PresenceGetUsersOnlineSessions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUsersOnlineSessionsArgs() interface{} {
	return &GetUsersOnlineSessionsArgs{}
}

func newGetUsersOnlineSessionsResult() interface{} {
	return &GetUsersOnlineSessionsResult{}
}

type GetUsersOnlineSessionsArgs struct {
	Req *presence.TLPresenceGetUsersOnlineSessions
}

func (p *GetUsersOnlineSessionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in GetUsersOnlineSessionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUsersOnlineSessionsArgs) Unmarshal(in []byte) error {
	msg := new(presence.TLPresenceGetUsersOnlineSessions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUsersOnlineSessionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in GetUsersOnlineSessionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUsersOnlineSessionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(presence.TLPresenceGetUsersOnlineSessions)
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

var GetUsersOnlineSessionsArgs_Req_DEFAULT *presence.TLPresenceGetUsersOnlineSessions

func (p *GetUsersOnlineSessionsArgs) GetReq() *presence.TLPresenceGetUsersOnlineSessions {
	if !p.IsSetReq() {
		return GetUsersOnlineSessionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUsersOnlineSessionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUsersOnlineSessionsResult struct {
	Success *presence.VectorUserOnlineSessions
}

var GetUsersOnlineSessionsResult_Success_DEFAULT *presence.VectorUserOnlineSessions

func (p *GetUsersOnlineSessionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in GetUsersOnlineSessionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUsersOnlineSessionsResult) Unmarshal(in []byte) error {
	msg := new(presence.VectorUserOnlineSessions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUsersOnlineSessionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in GetUsersOnlineSessionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUsersOnlineSessionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(presence.VectorUserOnlineSessions)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUsersOnlineSessionsResult) GetSuccess() *presence.VectorUserOnlineSessions {
	if !p.IsSetSuccess() {
		return GetUsersOnlineSessionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUsersOnlineSessionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*presence.VectorUserOnlineSessions)
}

func (p *GetUsersOnlineSessionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUsersOnlineSessionsResult) GetResult() interface{} {
	return p.Success
}

func getGatewaySessionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetGatewaySessionsArgs)
	realResult := result.(*GetGatewaySessionsResult)
	success, err := handler.(presence.RPCPresence).PresenceGetGatewaySessions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetGatewaySessionsArgs() interface{} {
	return &GetGatewaySessionsArgs{}
}

func newGetGatewaySessionsResult() interface{} {
	return &GetGatewaySessionsResult{}
}

type GetGatewaySessionsArgs struct {
	Req *presence.TLPresenceGetGatewaySessions
}

func (p *GetGatewaySessionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in GetGatewaySessionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetGatewaySessionsArgs) Unmarshal(in []byte) error {
	msg := new(presence.TLPresenceGetGatewaySessions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetGatewaySessionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in GetGatewaySessionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetGatewaySessionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(presence.TLPresenceGetGatewaySessions)
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

var GetGatewaySessionsArgs_Req_DEFAULT *presence.TLPresenceGetGatewaySessions

func (p *GetGatewaySessionsArgs) GetReq() *presence.TLPresenceGetGatewaySessions {
	if !p.IsSetReq() {
		return GetGatewaySessionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetGatewaySessionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetGatewaySessionsResult struct {
	Success *presence.VectorOnlineSession
}

var GetGatewaySessionsResult_Success_DEFAULT *presence.VectorOnlineSession

func (p *GetGatewaySessionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in GetGatewaySessionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetGatewaySessionsResult) Unmarshal(in []byte) error {
	msg := new(presence.VectorOnlineSession)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetGatewaySessionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in GetGatewaySessionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetGatewaySessionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(presence.VectorOnlineSession)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetGatewaySessionsResult) GetSuccess() *presence.VectorOnlineSession {
	if !p.IsSetSuccess() {
		return GetGatewaySessionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetGatewaySessionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*presence.VectorOnlineSession)
}

func (p *GetGatewaySessionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetGatewaySessionsResult) GetResult() interface{} {
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

func (p *kClient) PresenceSetSessionOnline(ctx context.Context, req *presence.TLPresenceSetSessionOnline) (r *tg.Bool, err error) {
	// var _args SetSessionOnlineArgs
	// _args.Req = req
	var _result SetSessionOnlineResult

	if err = p.c.Call(ctx, "/presence.RPCPresence/presence.setSessionOnline", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) PresenceSetSessionOffline(ctx context.Context, req *presence.TLPresenceSetSessionOffline) (r *tg.Bool, err error) {
	// var _args SetSessionOfflineArgs
	// _args.Req = req
	var _result SetSessionOfflineResult

	if err = p.c.Call(ctx, "/presence.RPCPresence/presence.setSessionOffline", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) PresenceGetUserOnlineSessions(ctx context.Context, req *presence.TLPresenceGetUserOnlineSessions) (r *presence.UserOnlineSessions, err error) {
	// var _args GetUserOnlineSessionsArgs
	// _args.Req = req
	var _result GetUserOnlineSessionsResult

	if err = p.c.Call(ctx, "/presence.RPCPresence/presence.getUserOnlineSessions", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) PresenceGetUsersOnlineSessions(ctx context.Context, req *presence.TLPresenceGetUsersOnlineSessions) (r *presence.VectorUserOnlineSessions, err error) {
	// var _args GetUsersOnlineSessionsArgs
	// _args.Req = req
	var _result GetUsersOnlineSessionsResult

	if err = p.c.Call(ctx, "/presence.RPCPresence/presence.getUsersOnlineSessions", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) PresenceGetGatewaySessions(ctx context.Context, req *presence.TLPresenceGetGatewaySessions) (r *presence.VectorOnlineSession, err error) {
	// var _args GetGatewaySessionsArgs
	// _args.Req = req
	var _result GetGatewaySessionsResult

	if err = p.c.Call(ctx, "/presence.RPCPresence/presence.getGatewaySessions", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}
