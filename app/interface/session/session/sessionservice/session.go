/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package sessionservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"session.queryAuthKey": kitex.NewMethodInfo(
		queryAuthKeyHandler,
		newQueryAuthKeyArgs,
		newQueryAuthKeyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"session.setAuthKey": kitex.NewMethodInfo(
		setAuthKeyHandler,
		newSetAuthKeyArgs,
		newSetAuthKeyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"session.createSession": kitex.NewMethodInfo(
		createSessionHandler,
		newCreateSessionArgs,
		newCreateSessionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"session.sendDataToSession": kitex.NewMethodInfo(
		sendDataToSessionHandler,
		newSendDataToSessionArgs,
		newSendDataToSessionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"session.sendHttpDataToSession": kitex.NewMethodInfo(
		sendHttpDataToSessionHandler,
		newSendHttpDataToSessionArgs,
		newSendHttpDataToSessionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"session.closeSession": kitex.NewMethodInfo(
		closeSessionHandler,
		newCloseSessionArgs,
		newCloseSessionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"session.pushUpdatesData": kitex.NewMethodInfo(
		pushUpdatesDataHandler,
		newPushUpdatesDataArgs,
		newPushUpdatesDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"session.pushSessionUpdatesData": kitex.NewMethodInfo(
		pushSessionUpdatesDataHandler,
		newPushSessionUpdatesDataArgs,
		newPushSessionUpdatesDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"session.pushRpcResultData": kitex.NewMethodInfo(
		pushRpcResultDataHandler,
		newPushRpcResultDataArgs,
		newPushRpcResultDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	sessionServiceServiceInfo                = NewServiceInfo()
	sessionServiceServiceInfoForClient       = NewServiceInfoForClient()
	sessionServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCSession", sessionServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCSession", sessionServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCSession", sessionServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return sessionServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return sessionServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return sessionServiceServiceInfoForClient
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
	serviceName := "RPCSession"
	handlerType := (*session.RPCSession)(nil)
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
		"PackageName": "session",
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

func queryAuthKeyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*QueryAuthKeyArgs)
	realResult := result.(*QueryAuthKeyResult)
	success, err := handler.(session.RPCSession).SessionQueryAuthKey(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newQueryAuthKeyArgs() interface{} {
	return &QueryAuthKeyArgs{}
}

func newQueryAuthKeyResult() interface{} {
	return &QueryAuthKeyResult{}
}

type QueryAuthKeyArgs struct {
	Req *session.TLSessionQueryAuthKey
}

func (p *QueryAuthKeyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryAuthKeyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *QueryAuthKeyArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionQueryAuthKey)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *QueryAuthKeyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in QueryAuthKeyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *QueryAuthKeyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(session.TLSessionQueryAuthKey)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var QueryAuthKeyArgs_Req_DEFAULT *session.TLSessionQueryAuthKey

func (p *QueryAuthKeyArgs) GetReq() *session.TLSessionQueryAuthKey {
	if !p.IsSetReq() {
		return QueryAuthKeyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryAuthKeyArgs) IsSetReq() bool {
	return p.Req != nil
}

type QueryAuthKeyResult struct {
	Success *tg.AuthKeyInfo
}

var QueryAuthKeyResult_Success_DEFAULT *tg.AuthKeyInfo

func (p *QueryAuthKeyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryAuthKeyResult")
	}
	return json.Marshal(p.Success)
}

func (p *QueryAuthKeyResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthKeyInfo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryAuthKeyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in QueryAuthKeyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *QueryAuthKeyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthKeyInfo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryAuthKeyResult) GetSuccess() *tg.AuthKeyInfo {
	if !p.IsSetSuccess() {
		return QueryAuthKeyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryAuthKeyResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthKeyInfo)
}

func (p *QueryAuthKeyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryAuthKeyResult) GetResult() interface{} {
	return p.Success
}

func setAuthKeyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetAuthKeyArgs)
	realResult := result.(*SetAuthKeyResult)
	success, err := handler.(session.RPCSession).SessionSetAuthKey(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetAuthKeyArgs() interface{} {
	return &SetAuthKeyArgs{}
}

func newSetAuthKeyResult() interface{} {
	return &SetAuthKeyResult{}
}

type SetAuthKeyArgs struct {
	Req *session.TLSessionSetAuthKey
}

func (p *SetAuthKeyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetAuthKeyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetAuthKeyArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionSetAuthKey)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetAuthKeyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetAuthKeyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetAuthKeyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(session.TLSessionSetAuthKey)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetAuthKeyArgs_Req_DEFAULT *session.TLSessionSetAuthKey

func (p *SetAuthKeyArgs) GetReq() *session.TLSessionSetAuthKey {
	if !p.IsSetReq() {
		return SetAuthKeyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetAuthKeyArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetAuthKeyResult struct {
	Success *tg.Bool
}

var SetAuthKeyResult_Success_DEFAULT *tg.Bool

func (p *SetAuthKeyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetAuthKeyResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetAuthKeyResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetAuthKeyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetAuthKeyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetAuthKeyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetAuthKeyResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetAuthKeyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetAuthKeyResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetAuthKeyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetAuthKeyResult) GetResult() interface{} {
	return p.Success
}

func createSessionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CreateSessionArgs)
	realResult := result.(*CreateSessionResult)
	success, err := handler.(session.RPCSession).SessionCreateSession(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCreateSessionArgs() interface{} {
	return &CreateSessionArgs{}
}

func newCreateSessionResult() interface{} {
	return &CreateSessionResult{}
}

type CreateSessionArgs struct {
	Req *session.TLSessionCreateSession
}

func (p *CreateSessionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateSessionArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CreateSessionArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionCreateSession)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CreateSessionArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CreateSessionArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CreateSessionArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(session.TLSessionCreateSession)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CreateSessionArgs_Req_DEFAULT *session.TLSessionCreateSession

func (p *CreateSessionArgs) GetReq() *session.TLSessionCreateSession {
	if !p.IsSetReq() {
		return CreateSessionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateSessionArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateSessionResult struct {
	Success *tg.Bool
}

var CreateSessionResult_Success_DEFAULT *tg.Bool

func (p *CreateSessionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateSessionResult")
	}
	return json.Marshal(p.Success)
}

func (p *CreateSessionResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateSessionResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CreateSessionResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CreateSessionResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateSessionResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return CreateSessionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateSessionResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *CreateSessionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateSessionResult) GetResult() interface{} {
	return p.Success
}

func sendDataToSessionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SendDataToSessionArgs)
	realResult := result.(*SendDataToSessionResult)
	success, err := handler.(session.RPCSession).SessionSendDataToSession(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSendDataToSessionArgs() interface{} {
	return &SendDataToSessionArgs{}
}

func newSendDataToSessionResult() interface{} {
	return &SendDataToSessionResult{}
}

type SendDataToSessionArgs struct {
	Req *session.TLSessionSendDataToSession
}

func (p *SendDataToSessionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendDataToSessionArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SendDataToSessionArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionSendDataToSession)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SendDataToSessionArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SendDataToSessionArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SendDataToSessionArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(session.TLSessionSendDataToSession)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SendDataToSessionArgs_Req_DEFAULT *session.TLSessionSendDataToSession

func (p *SendDataToSessionArgs) GetReq() *session.TLSessionSendDataToSession {
	if !p.IsSetReq() {
		return SendDataToSessionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendDataToSessionArgs) IsSetReq() bool {
	return p.Req != nil
}

type SendDataToSessionResult struct {
	Success *tg.Bool
}

var SendDataToSessionResult_Success_DEFAULT *tg.Bool

func (p *SendDataToSessionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SendDataToSessionResult")
	}
	return json.Marshal(p.Success)
}

func (p *SendDataToSessionResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendDataToSessionResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SendDataToSessionResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SendDataToSessionResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendDataToSessionResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SendDataToSessionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendDataToSessionResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SendDataToSessionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendDataToSessionResult) GetResult() interface{} {
	return p.Success
}

func sendHttpDataToSessionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SendHttpDataToSessionArgs)
	realResult := result.(*SendHttpDataToSessionResult)
	success, err := handler.(session.RPCSession).SessionSendHttpDataToSession(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSendHttpDataToSessionArgs() interface{} {
	return &SendHttpDataToSessionArgs{}
}

func newSendHttpDataToSessionResult() interface{} {
	return &SendHttpDataToSessionResult{}
}

type SendHttpDataToSessionArgs struct {
	Req *session.TLSessionSendHttpDataToSession
}

func (p *SendHttpDataToSessionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendHttpDataToSessionArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SendHttpDataToSessionArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionSendHttpDataToSession)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SendHttpDataToSessionArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SendHttpDataToSessionArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SendHttpDataToSessionArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(session.TLSessionSendHttpDataToSession)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SendHttpDataToSessionArgs_Req_DEFAULT *session.TLSessionSendHttpDataToSession

func (p *SendHttpDataToSessionArgs) GetReq() *session.TLSessionSendHttpDataToSession {
	if !p.IsSetReq() {
		return SendHttpDataToSessionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendHttpDataToSessionArgs) IsSetReq() bool {
	return p.Req != nil
}

type SendHttpDataToSessionResult struct {
	Success *session.HttpSessionData
}

var SendHttpDataToSessionResult_Success_DEFAULT *session.HttpSessionData

func (p *SendHttpDataToSessionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SendHttpDataToSessionResult")
	}
	return json.Marshal(p.Success)
}

func (p *SendHttpDataToSessionResult) Unmarshal(in []byte) error {
	msg := new(session.HttpSessionData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendHttpDataToSessionResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SendHttpDataToSessionResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SendHttpDataToSessionResult) Decode(d *bin.Decoder) (err error) {
	msg := new(session.HttpSessionData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendHttpDataToSessionResult) GetSuccess() *session.HttpSessionData {
	if !p.IsSetSuccess() {
		return SendHttpDataToSessionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendHttpDataToSessionResult) SetSuccess(x interface{}) {
	p.Success = x.(*session.HttpSessionData)
}

func (p *SendHttpDataToSessionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendHttpDataToSessionResult) GetResult() interface{} {
	return p.Success
}

func closeSessionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CloseSessionArgs)
	realResult := result.(*CloseSessionResult)
	success, err := handler.(session.RPCSession).SessionCloseSession(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCloseSessionArgs() interface{} {
	return &CloseSessionArgs{}
}

func newCloseSessionResult() interface{} {
	return &CloseSessionResult{}
}

type CloseSessionArgs struct {
	Req *session.TLSessionCloseSession
}

func (p *CloseSessionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CloseSessionArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CloseSessionArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionCloseSession)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CloseSessionArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CloseSessionArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CloseSessionArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(session.TLSessionCloseSession)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CloseSessionArgs_Req_DEFAULT *session.TLSessionCloseSession

func (p *CloseSessionArgs) GetReq() *session.TLSessionCloseSession {
	if !p.IsSetReq() {
		return CloseSessionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CloseSessionArgs) IsSetReq() bool {
	return p.Req != nil
}

type CloseSessionResult struct {
	Success *tg.Bool
}

var CloseSessionResult_Success_DEFAULT *tg.Bool

func (p *CloseSessionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CloseSessionResult")
	}
	return json.Marshal(p.Success)
}

func (p *CloseSessionResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CloseSessionResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CloseSessionResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CloseSessionResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CloseSessionResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return CloseSessionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CloseSessionResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *CloseSessionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CloseSessionResult) GetResult() interface{} {
	return p.Success
}

func pushUpdatesDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushUpdatesDataArgs)
	realResult := result.(*PushUpdatesDataResult)
	success, err := handler.(session.RPCSession).SessionPushUpdatesData(ctx, realArg.Req)
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
	Req *session.TLSessionPushUpdatesData
}

func (p *PushUpdatesDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushUpdatesDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushUpdatesDataArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionPushUpdatesData)
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
	msg := new(session.TLSessionPushUpdatesData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PushUpdatesDataArgs_Req_DEFAULT *session.TLSessionPushUpdatesData

func (p *PushUpdatesDataArgs) GetReq() *session.TLSessionPushUpdatesData {
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
	success, err := handler.(session.RPCSession).SessionPushSessionUpdatesData(ctx, realArg.Req)
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
	Req *session.TLSessionPushSessionUpdatesData
}

func (p *PushSessionUpdatesDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushSessionUpdatesDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushSessionUpdatesDataArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionPushSessionUpdatesData)
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
	msg := new(session.TLSessionPushSessionUpdatesData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PushSessionUpdatesDataArgs_Req_DEFAULT *session.TLSessionPushSessionUpdatesData

func (p *PushSessionUpdatesDataArgs) GetReq() *session.TLSessionPushSessionUpdatesData {
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
	success, err := handler.(session.RPCSession).SessionPushRpcResultData(ctx, realArg.Req)
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
	Req *session.TLSessionPushRpcResultData
}

func (p *PushRpcResultDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushRpcResultDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushRpcResultDataArgs) Unmarshal(in []byte) error {
	msg := new(session.TLSessionPushRpcResultData)
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
	msg := new(session.TLSessionPushRpcResultData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PushRpcResultDataArgs_Req_DEFAULT *session.TLSessionPushRpcResultData

func (p *PushRpcResultDataArgs) GetReq() *session.TLSessionPushRpcResultData {
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

func (p *kClient) SessionQueryAuthKey(ctx context.Context, req *session.TLSessionQueryAuthKey) (r *tg.AuthKeyInfo, err error) {
	// var _args QueryAuthKeyArgs
	// _args.Req = req
	// var _result QueryAuthKeyResult

	_result := new(tg.AuthKeyInfo)

	if err = p.c.Call(ctx, "session.queryAuthKey", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SessionSetAuthKey(ctx context.Context, req *session.TLSessionSetAuthKey) (r *tg.Bool, err error) {
	// var _args SetAuthKeyArgs
	// _args.Req = req
	// var _result SetAuthKeyResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "session.setAuthKey", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SessionCreateSession(ctx context.Context, req *session.TLSessionCreateSession) (r *tg.Bool, err error) {
	// var _args CreateSessionArgs
	// _args.Req = req
	// var _result CreateSessionResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "session.createSession", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SessionSendDataToSession(ctx context.Context, req *session.TLSessionSendDataToSession) (r *tg.Bool, err error) {
	// var _args SendDataToSessionArgs
	// _args.Req = req
	// var _result SendDataToSessionResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "session.sendDataToSession", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SessionSendHttpDataToSession(ctx context.Context, req *session.TLSessionSendHttpDataToSession) (r *session.HttpSessionData, err error) {
	// var _args SendHttpDataToSessionArgs
	// _args.Req = req
	// var _result SendHttpDataToSessionResult

	_result := new(session.HttpSessionData)

	if err = p.c.Call(ctx, "session.sendHttpDataToSession", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SessionCloseSession(ctx context.Context, req *session.TLSessionCloseSession) (r *tg.Bool, err error) {
	// var _args CloseSessionArgs
	// _args.Req = req
	// var _result CloseSessionResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "session.closeSession", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SessionPushUpdatesData(ctx context.Context, req *session.TLSessionPushUpdatesData) (r *tg.Bool, err error) {
	// var _args PushUpdatesDataArgs
	// _args.Req = req
	// var _result PushUpdatesDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "session.pushUpdatesData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SessionPushSessionUpdatesData(ctx context.Context, req *session.TLSessionPushSessionUpdatesData) (r *tg.Bool, err error) {
	// var _args PushSessionUpdatesDataArgs
	// _args.Req = req
	// var _result PushSessionUpdatesDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "session.pushSessionUpdatesData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) SessionPushRpcResultData(ctx context.Context, req *session.TLSessionPushRpcResultData) (r *tg.Bool, err error) {
	// var _args PushRpcResultDataArgs
	// _args.Req = req
	// var _result PushRpcResultDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "session.pushRpcResultData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
