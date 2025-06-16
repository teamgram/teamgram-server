/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package authsessionservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"authsession.getAuthorizations": kitex.NewMethodInfo(
		getAuthorizationsHandler,
		newGetAuthorizationsArgs,
		newGetAuthorizationsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.resetAuthorization": kitex.NewMethodInfo(
		resetAuthorizationHandler,
		newResetAuthorizationArgs,
		newResetAuthorizationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getLayer": kitex.NewMethodInfo(
		getLayerHandler,
		newGetLayerArgs,
		newGetLayerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getLangPack": kitex.NewMethodInfo(
		getLangPackHandler,
		newGetLangPackArgs,
		newGetLangPackResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getClient": kitex.NewMethodInfo(
		getClientHandler,
		newGetClientArgs,
		newGetClientResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getLangCode": kitex.NewMethodInfo(
		getLangCodeHandler,
		newGetLangCodeArgs,
		newGetLangCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getUserId": kitex.NewMethodInfo(
		getUserIdHandler,
		newGetUserIdArgs,
		newGetUserIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getPushSessionId": kitex.NewMethodInfo(
		getPushSessionIdHandler,
		newGetPushSessionIdArgs,
		newGetPushSessionIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getFutureSalts": kitex.NewMethodInfo(
		getFutureSaltsHandler,
		newGetFutureSaltsArgs,
		newGetFutureSaltsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.queryAuthKey": kitex.NewMethodInfo(
		queryAuthKeyHandler,
		newQueryAuthKeyArgs,
		newQueryAuthKeyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.setAuthKey": kitex.NewMethodInfo(
		setAuthKeyHandler,
		newSetAuthKeyArgs,
		newSetAuthKeyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.bindAuthKeyUser": kitex.NewMethodInfo(
		bindAuthKeyUserHandler,
		newBindAuthKeyUserArgs,
		newBindAuthKeyUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.unbindAuthKeyUser": kitex.NewMethodInfo(
		unbindAuthKeyUserHandler,
		newUnbindAuthKeyUserArgs,
		newUnbindAuthKeyUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getPermAuthKeyId": kitex.NewMethodInfo(
		getPermAuthKeyIdHandler,
		newGetPermAuthKeyIdArgs,
		newGetPermAuthKeyIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.bindTempAuthKey": kitex.NewMethodInfo(
		bindTempAuthKeyHandler,
		newBindTempAuthKeyArgs,
		newBindTempAuthKeyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.setClientSessionInfo": kitex.NewMethodInfo(
		setClientSessionInfoHandler,
		newSetClientSessionInfoArgs,
		newSetClientSessionInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getAuthorization": kitex.NewMethodInfo(
		getAuthorizationHandler,
		newGetAuthorizationArgs,
		newGetAuthorizationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.getAuthStateData": kitex.NewMethodInfo(
		getAuthStateDataHandler,
		newGetAuthStateDataArgs,
		newGetAuthStateDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.setLayer": kitex.NewMethodInfo(
		setLayerHandler,
		newSetLayerArgs,
		newSetLayerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.setInitConnection": kitex.NewMethodInfo(
		setInitConnectionHandler,
		newSetInitConnectionArgs,
		newSetInitConnectionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"authsession.setAndroidPushSessionId": kitex.NewMethodInfo(
		setAndroidPushSessionIdHandler,
		newSetAndroidPushSessionIdArgs,
		newSetAndroidPushSessionIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	authsessionServiceServiceInfo                = NewServiceInfo()
	authsessionServiceServiceInfoForClient       = NewServiceInfoForClient()
	authsessionServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return authsessionServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return authsessionServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return authsessionServiceServiceInfoForClient
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
	serviceName := "RPCAuthsession"
	handlerType := (*authsession.RPCAuthsession)(nil)
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
		"PackageName": "authsession",
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

func getAuthorizationsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAuthorizationsArgs)
	realResult := result.(*GetAuthorizationsResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetAuthorizations(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAuthorizationsArgs() interface{} {
	return &GetAuthorizationsArgs{}
}

func newGetAuthorizationsResult() interface{} {
	return &GetAuthorizationsResult{}
}

type GetAuthorizationsArgs struct {
	Req *authsession.TLAuthsessionGetAuthorizations
}

func (p *GetAuthorizationsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAuthorizationsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAuthorizationsArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetAuthorizations)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAuthorizationsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAuthorizationsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAuthorizationsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetAuthorizations)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAuthorizationsArgs_Req_DEFAULT *authsession.TLAuthsessionGetAuthorizations

func (p *GetAuthorizationsArgs) GetReq() *authsession.TLAuthsessionGetAuthorizations {
	if !p.IsSetReq() {
		return GetAuthorizationsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAuthorizationsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAuthorizationsResult struct {
	Success *tg.AccountAuthorizations
}

var GetAuthorizationsResult_Success_DEFAULT *tg.AccountAuthorizations

func (p *GetAuthorizationsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAuthorizationsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAuthorizationsResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountAuthorizations)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthorizationsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAuthorizationsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAuthorizationsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountAuthorizations)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthorizationsResult) GetSuccess() *tg.AccountAuthorizations {
	if !p.IsSetSuccess() {
		return GetAuthorizationsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAuthorizationsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountAuthorizations)
}

func (p *GetAuthorizationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAuthorizationsResult) GetResult() interface{} {
	return p.Success
}

func resetAuthorizationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ResetAuthorizationArgs)
	realResult := result.(*ResetAuthorizationResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionResetAuthorization(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newResetAuthorizationArgs() interface{} {
	return &ResetAuthorizationArgs{}
}

func newResetAuthorizationResult() interface{} {
	return &ResetAuthorizationResult{}
}

type ResetAuthorizationArgs struct {
	Req *authsession.TLAuthsessionResetAuthorization
}

func (p *ResetAuthorizationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ResetAuthorizationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ResetAuthorizationArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionResetAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ResetAuthorizationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ResetAuthorizationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ResetAuthorizationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionResetAuthorization)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ResetAuthorizationArgs_Req_DEFAULT *authsession.TLAuthsessionResetAuthorization

func (p *ResetAuthorizationArgs) GetReq() *authsession.TLAuthsessionResetAuthorization {
	if !p.IsSetReq() {
		return ResetAuthorizationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ResetAuthorizationArgs) IsSetReq() bool {
	return p.Req != nil
}

type ResetAuthorizationResult struct {
	Success *authsession.VectorLong
}

var ResetAuthorizationResult_Success_DEFAULT *authsession.VectorLong

func (p *ResetAuthorizationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ResetAuthorizationResult")
	}
	return json.Marshal(p.Success)
}

func (p *ResetAuthorizationResult) Unmarshal(in []byte) error {
	msg := new(authsession.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ResetAuthorizationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ResetAuthorizationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ResetAuthorizationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ResetAuthorizationResult) GetSuccess() *authsession.VectorLong {
	if !p.IsSetSuccess() {
		return ResetAuthorizationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ResetAuthorizationResult) SetSuccess(x interface{}) {
	p.Success = x.(*authsession.VectorLong)
}

func (p *ResetAuthorizationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ResetAuthorizationResult) GetResult() interface{} {
	return p.Success
}

func getLayerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetLayerArgs)
	realResult := result.(*GetLayerResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetLayer(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetLayerArgs() interface{} {
	return &GetLayerArgs{}
}

func newGetLayerResult() interface{} {
	return &GetLayerResult{}
}

type GetLayerArgs struct {
	Req *authsession.TLAuthsessionGetLayer
}

func (p *GetLayerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLayerArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetLayerArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetLayer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetLayerArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetLayerArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetLayerArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetLayer)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetLayerArgs_Req_DEFAULT *authsession.TLAuthsessionGetLayer

func (p *GetLayerArgs) GetReq() *authsession.TLAuthsessionGetLayer {
	if !p.IsSetReq() {
		return GetLayerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLayerArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLayerResult struct {
	Success *tg.Int32
}

var GetLayerResult_Success_DEFAULT *tg.Int32

func (p *GetLayerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLayerResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetLayerResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLayerResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetLayerResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetLayerResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLayerResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetLayerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLayerResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetLayerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLayerResult) GetResult() interface{} {
	return p.Success
}

func getLangPackHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetLangPackArgs)
	realResult := result.(*GetLangPackResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetLangPack(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetLangPackArgs() interface{} {
	return &GetLangPackArgs{}
}

func newGetLangPackResult() interface{} {
	return &GetLangPackResult{}
}

type GetLangPackArgs struct {
	Req *authsession.TLAuthsessionGetLangPack
}

func (p *GetLangPackArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLangPackArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetLangPackArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetLangPack)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetLangPackArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetLangPackArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetLangPackArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetLangPack)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetLangPackArgs_Req_DEFAULT *authsession.TLAuthsessionGetLangPack

func (p *GetLangPackArgs) GetReq() *authsession.TLAuthsessionGetLangPack {
	if !p.IsSetReq() {
		return GetLangPackArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLangPackArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLangPackResult struct {
	Success *tg.String
}

var GetLangPackResult_Success_DEFAULT *tg.String

func (p *GetLangPackResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLangPackResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetLangPackResult) Unmarshal(in []byte) error {
	msg := new(tg.String)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLangPackResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetLangPackResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetLangPackResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.String)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLangPackResult) GetSuccess() *tg.String {
	if !p.IsSetSuccess() {
		return GetLangPackResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLangPackResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.String)
}

func (p *GetLangPackResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLangPackResult) GetResult() interface{} {
	return p.Success
}

func getClientHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetClientArgs)
	realResult := result.(*GetClientResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetClient(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetClientArgs() interface{} {
	return &GetClientArgs{}
}

func newGetClientResult() interface{} {
	return &GetClientResult{}
}

type GetClientArgs struct {
	Req *authsession.TLAuthsessionGetClient
}

func (p *GetClientArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetClientArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetClientArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetClient)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetClientArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetClientArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetClientArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetClient)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetClientArgs_Req_DEFAULT *authsession.TLAuthsessionGetClient

func (p *GetClientArgs) GetReq() *authsession.TLAuthsessionGetClient {
	if !p.IsSetReq() {
		return GetClientArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetClientArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetClientResult struct {
	Success *tg.String
}

var GetClientResult_Success_DEFAULT *tg.String

func (p *GetClientResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetClientResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetClientResult) Unmarshal(in []byte) error {
	msg := new(tg.String)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetClientResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetClientResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetClientResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.String)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetClientResult) GetSuccess() *tg.String {
	if !p.IsSetSuccess() {
		return GetClientResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetClientResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.String)
}

func (p *GetClientResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetClientResult) GetResult() interface{} {
	return p.Success
}

func getLangCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetLangCodeArgs)
	realResult := result.(*GetLangCodeResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetLangCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetLangCodeArgs() interface{} {
	return &GetLangCodeArgs{}
}

func newGetLangCodeResult() interface{} {
	return &GetLangCodeResult{}
}

type GetLangCodeArgs struct {
	Req *authsession.TLAuthsessionGetLangCode
}

func (p *GetLangCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLangCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetLangCodeArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetLangCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetLangCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetLangCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetLangCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetLangCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetLangCodeArgs_Req_DEFAULT *authsession.TLAuthsessionGetLangCode

func (p *GetLangCodeArgs) GetReq() *authsession.TLAuthsessionGetLangCode {
	if !p.IsSetReq() {
		return GetLangCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLangCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLangCodeResult struct {
	Success *tg.String
}

var GetLangCodeResult_Success_DEFAULT *tg.String

func (p *GetLangCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLangCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetLangCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.String)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLangCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetLangCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetLangCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.String)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLangCodeResult) GetSuccess() *tg.String {
	if !p.IsSetSuccess() {
		return GetLangCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLangCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.String)
}

func (p *GetLangCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLangCodeResult) GetResult() interface{} {
	return p.Success
}

func getUserIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserIdArgs)
	realResult := result.(*GetUserIdResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetUserId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserIdArgs() interface{} {
	return &GetUserIdArgs{}
}

func newGetUserIdResult() interface{} {
	return &GetUserIdResult{}
}

type GetUserIdArgs struct {
	Req *authsession.TLAuthsessionGetUserId
}

func (p *GetUserIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserIdArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetUserId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetUserId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserIdArgs_Req_DEFAULT *authsession.TLAuthsessionGetUserId

func (p *GetUserIdArgs) GetReq() *authsession.TLAuthsessionGetUserId {
	if !p.IsSetReq() {
		return GetUserIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserIdResult struct {
	Success *tg.Int64
}

var GetUserIdResult_Success_DEFAULT *tg.Int64

func (p *GetUserIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserIdResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return GetUserIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *GetUserIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserIdResult) GetResult() interface{} {
	return p.Success
}

func getPushSessionIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPushSessionIdArgs)
	realResult := result.(*GetPushSessionIdResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetPushSessionId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPushSessionIdArgs() interface{} {
	return &GetPushSessionIdArgs{}
}

func newGetPushSessionIdResult() interface{} {
	return &GetPushSessionIdResult{}
}

type GetPushSessionIdArgs struct {
	Req *authsession.TLAuthsessionGetPushSessionId
}

func (p *GetPushSessionIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPushSessionIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPushSessionIdArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetPushSessionId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPushSessionIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPushSessionIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPushSessionIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetPushSessionId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPushSessionIdArgs_Req_DEFAULT *authsession.TLAuthsessionGetPushSessionId

func (p *GetPushSessionIdArgs) GetReq() *authsession.TLAuthsessionGetPushSessionId {
	if !p.IsSetReq() {
		return GetPushSessionIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPushSessionIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPushSessionIdResult struct {
	Success *tg.Int64
}

var GetPushSessionIdResult_Success_DEFAULT *tg.Int64

func (p *GetPushSessionIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPushSessionIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPushSessionIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPushSessionIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPushSessionIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPushSessionIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPushSessionIdResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return GetPushSessionIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPushSessionIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *GetPushSessionIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPushSessionIdResult) GetResult() interface{} {
	return p.Success
}

func getFutureSaltsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetFutureSaltsArgs)
	realResult := result.(*GetFutureSaltsResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetFutureSalts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetFutureSaltsArgs() interface{} {
	return &GetFutureSaltsArgs{}
}

func newGetFutureSaltsResult() interface{} {
	return &GetFutureSaltsResult{}
}

type GetFutureSaltsArgs struct {
	Req *authsession.TLAuthsessionGetFutureSalts
}

func (p *GetFutureSaltsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFutureSaltsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetFutureSaltsArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetFutureSalts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetFutureSaltsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetFutureSaltsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetFutureSaltsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetFutureSalts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetFutureSaltsArgs_Req_DEFAULT *authsession.TLAuthsessionGetFutureSalts

func (p *GetFutureSaltsArgs) GetReq() *authsession.TLAuthsessionGetFutureSalts {
	if !p.IsSetReq() {
		return GetFutureSaltsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFutureSaltsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetFutureSaltsResult struct {
	Success *tg.FutureSalts
}

var GetFutureSaltsResult_Success_DEFAULT *tg.FutureSalts

func (p *GetFutureSaltsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFutureSaltsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetFutureSaltsResult) Unmarshal(in []byte) error {
	msg := new(tg.FutureSalts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFutureSaltsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetFutureSaltsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetFutureSaltsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.FutureSalts)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFutureSaltsResult) GetSuccess() *tg.FutureSalts {
	if !p.IsSetSuccess() {
		return GetFutureSaltsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFutureSaltsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.FutureSalts)
}

func (p *GetFutureSaltsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFutureSaltsResult) GetResult() interface{} {
	return p.Success
}

func queryAuthKeyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*QueryAuthKeyArgs)
	realResult := result.(*QueryAuthKeyResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionQueryAuthKey(ctx, realArg.Req)
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
	Req *authsession.TLAuthsessionQueryAuthKey
}

func (p *QueryAuthKeyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryAuthKeyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *QueryAuthKeyArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionQueryAuthKey)
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
	msg := new(authsession.TLAuthsessionQueryAuthKey)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var QueryAuthKeyArgs_Req_DEFAULT *authsession.TLAuthsessionQueryAuthKey

func (p *QueryAuthKeyArgs) GetReq() *authsession.TLAuthsessionQueryAuthKey {
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
	success, err := handler.(authsession.RPCAuthsession).AuthsessionSetAuthKey(ctx, realArg.Req)
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
	Req *authsession.TLAuthsessionSetAuthKey
}

func (p *SetAuthKeyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetAuthKeyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetAuthKeyArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionSetAuthKey)
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
	msg := new(authsession.TLAuthsessionSetAuthKey)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetAuthKeyArgs_Req_DEFAULT *authsession.TLAuthsessionSetAuthKey

func (p *SetAuthKeyArgs) GetReq() *authsession.TLAuthsessionSetAuthKey {
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

func bindAuthKeyUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*BindAuthKeyUserArgs)
	realResult := result.(*BindAuthKeyUserResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionBindAuthKeyUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newBindAuthKeyUserArgs() interface{} {
	return &BindAuthKeyUserArgs{}
}

func newBindAuthKeyUserResult() interface{} {
	return &BindAuthKeyUserResult{}
}

type BindAuthKeyUserArgs struct {
	Req *authsession.TLAuthsessionBindAuthKeyUser
}

func (p *BindAuthKeyUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in BindAuthKeyUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *BindAuthKeyUserArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionBindAuthKeyUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *BindAuthKeyUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in BindAuthKeyUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *BindAuthKeyUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionBindAuthKeyUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var BindAuthKeyUserArgs_Req_DEFAULT *authsession.TLAuthsessionBindAuthKeyUser

func (p *BindAuthKeyUserArgs) GetReq() *authsession.TLAuthsessionBindAuthKeyUser {
	if !p.IsSetReq() {
		return BindAuthKeyUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *BindAuthKeyUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type BindAuthKeyUserResult struct {
	Success *tg.Int64
}

var BindAuthKeyUserResult_Success_DEFAULT *tg.Int64

func (p *BindAuthKeyUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in BindAuthKeyUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *BindAuthKeyUserResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BindAuthKeyUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in BindAuthKeyUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *BindAuthKeyUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BindAuthKeyUserResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return BindAuthKeyUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *BindAuthKeyUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *BindAuthKeyUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BindAuthKeyUserResult) GetResult() interface{} {
	return p.Success
}

func unbindAuthKeyUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UnbindAuthKeyUserArgs)
	realResult := result.(*UnbindAuthKeyUserResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionUnbindAuthKeyUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUnbindAuthKeyUserArgs() interface{} {
	return &UnbindAuthKeyUserArgs{}
}

func newUnbindAuthKeyUserResult() interface{} {
	return &UnbindAuthKeyUserResult{}
}

type UnbindAuthKeyUserArgs struct {
	Req *authsession.TLAuthsessionUnbindAuthKeyUser
}

func (p *UnbindAuthKeyUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UnbindAuthKeyUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UnbindAuthKeyUserArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionUnbindAuthKeyUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UnbindAuthKeyUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UnbindAuthKeyUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UnbindAuthKeyUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionUnbindAuthKeyUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UnbindAuthKeyUserArgs_Req_DEFAULT *authsession.TLAuthsessionUnbindAuthKeyUser

func (p *UnbindAuthKeyUserArgs) GetReq() *authsession.TLAuthsessionUnbindAuthKeyUser {
	if !p.IsSetReq() {
		return UnbindAuthKeyUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UnbindAuthKeyUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type UnbindAuthKeyUserResult struct {
	Success *tg.Bool
}

var UnbindAuthKeyUserResult_Success_DEFAULT *tg.Bool

func (p *UnbindAuthKeyUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UnbindAuthKeyUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *UnbindAuthKeyUserResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnbindAuthKeyUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UnbindAuthKeyUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UnbindAuthKeyUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnbindAuthKeyUserResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UnbindAuthKeyUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UnbindAuthKeyUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UnbindAuthKeyUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UnbindAuthKeyUserResult) GetResult() interface{} {
	return p.Success
}

func getPermAuthKeyIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPermAuthKeyIdArgs)
	realResult := result.(*GetPermAuthKeyIdResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetPermAuthKeyId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPermAuthKeyIdArgs() interface{} {
	return &GetPermAuthKeyIdArgs{}
}

func newGetPermAuthKeyIdResult() interface{} {
	return &GetPermAuthKeyIdResult{}
}

type GetPermAuthKeyIdArgs struct {
	Req *authsession.TLAuthsessionGetPermAuthKeyId
}

func (p *GetPermAuthKeyIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPermAuthKeyIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPermAuthKeyIdArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetPermAuthKeyId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPermAuthKeyIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPermAuthKeyIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPermAuthKeyIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetPermAuthKeyId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPermAuthKeyIdArgs_Req_DEFAULT *authsession.TLAuthsessionGetPermAuthKeyId

func (p *GetPermAuthKeyIdArgs) GetReq() *authsession.TLAuthsessionGetPermAuthKeyId {
	if !p.IsSetReq() {
		return GetPermAuthKeyIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPermAuthKeyIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPermAuthKeyIdResult struct {
	Success *tg.Int64
}

var GetPermAuthKeyIdResult_Success_DEFAULT *tg.Int64

func (p *GetPermAuthKeyIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPermAuthKeyIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPermAuthKeyIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPermAuthKeyIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPermAuthKeyIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPermAuthKeyIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPermAuthKeyIdResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return GetPermAuthKeyIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPermAuthKeyIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *GetPermAuthKeyIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPermAuthKeyIdResult) GetResult() interface{} {
	return p.Success
}

func bindTempAuthKeyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*BindTempAuthKeyArgs)
	realResult := result.(*BindTempAuthKeyResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionBindTempAuthKey(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newBindTempAuthKeyArgs() interface{} {
	return &BindTempAuthKeyArgs{}
}

func newBindTempAuthKeyResult() interface{} {
	return &BindTempAuthKeyResult{}
}

type BindTempAuthKeyArgs struct {
	Req *authsession.TLAuthsessionBindTempAuthKey
}

func (p *BindTempAuthKeyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in BindTempAuthKeyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *BindTempAuthKeyArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionBindTempAuthKey)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *BindTempAuthKeyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in BindTempAuthKeyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *BindTempAuthKeyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionBindTempAuthKey)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var BindTempAuthKeyArgs_Req_DEFAULT *authsession.TLAuthsessionBindTempAuthKey

func (p *BindTempAuthKeyArgs) GetReq() *authsession.TLAuthsessionBindTempAuthKey {
	if !p.IsSetReq() {
		return BindTempAuthKeyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *BindTempAuthKeyArgs) IsSetReq() bool {
	return p.Req != nil
}

type BindTempAuthKeyResult struct {
	Success *tg.Bool
}

var BindTempAuthKeyResult_Success_DEFAULT *tg.Bool

func (p *BindTempAuthKeyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in BindTempAuthKeyResult")
	}
	return json.Marshal(p.Success)
}

func (p *BindTempAuthKeyResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BindTempAuthKeyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in BindTempAuthKeyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *BindTempAuthKeyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BindTempAuthKeyResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return BindTempAuthKeyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *BindTempAuthKeyResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *BindTempAuthKeyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BindTempAuthKeyResult) GetResult() interface{} {
	return p.Success
}

func setClientSessionInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetClientSessionInfoArgs)
	realResult := result.(*SetClientSessionInfoResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionSetClientSessionInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetClientSessionInfoArgs() interface{} {
	return &SetClientSessionInfoArgs{}
}

func newSetClientSessionInfoResult() interface{} {
	return &SetClientSessionInfoResult{}
}

type SetClientSessionInfoArgs struct {
	Req *authsession.TLAuthsessionSetClientSessionInfo
}

func (p *SetClientSessionInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetClientSessionInfoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetClientSessionInfoArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionSetClientSessionInfo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetClientSessionInfoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetClientSessionInfoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetClientSessionInfoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionSetClientSessionInfo)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetClientSessionInfoArgs_Req_DEFAULT *authsession.TLAuthsessionSetClientSessionInfo

func (p *SetClientSessionInfoArgs) GetReq() *authsession.TLAuthsessionSetClientSessionInfo {
	if !p.IsSetReq() {
		return SetClientSessionInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetClientSessionInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetClientSessionInfoResult struct {
	Success *tg.Bool
}

var SetClientSessionInfoResult_Success_DEFAULT *tg.Bool

func (p *SetClientSessionInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetClientSessionInfoResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetClientSessionInfoResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetClientSessionInfoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetClientSessionInfoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetClientSessionInfoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetClientSessionInfoResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetClientSessionInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetClientSessionInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetClientSessionInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetClientSessionInfoResult) GetResult() interface{} {
	return p.Success
}

func getAuthorizationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAuthorizationArgs)
	realResult := result.(*GetAuthorizationResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetAuthorization(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAuthorizationArgs() interface{} {
	return &GetAuthorizationArgs{}
}

func newGetAuthorizationResult() interface{} {
	return &GetAuthorizationResult{}
}

type GetAuthorizationArgs struct {
	Req *authsession.TLAuthsessionGetAuthorization
}

func (p *GetAuthorizationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAuthorizationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAuthorizationArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAuthorizationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAuthorizationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAuthorizationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetAuthorization)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAuthorizationArgs_Req_DEFAULT *authsession.TLAuthsessionGetAuthorization

func (p *GetAuthorizationArgs) GetReq() *authsession.TLAuthsessionGetAuthorization {
	if !p.IsSetReq() {
		return GetAuthorizationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAuthorizationArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAuthorizationResult struct {
	Success *tg.Authorization
}

var GetAuthorizationResult_Success_DEFAULT *tg.Authorization

func (p *GetAuthorizationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAuthorizationResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAuthorizationResult) Unmarshal(in []byte) error {
	msg := new(tg.Authorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthorizationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAuthorizationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAuthorizationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Authorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthorizationResult) GetSuccess() *tg.Authorization {
	if !p.IsSetSuccess() {
		return GetAuthorizationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAuthorizationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Authorization)
}

func (p *GetAuthorizationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAuthorizationResult) GetResult() interface{} {
	return p.Success
}

func getAuthStateDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAuthStateDataArgs)
	realResult := result.(*GetAuthStateDataResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionGetAuthStateData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAuthStateDataArgs() interface{} {
	return &GetAuthStateDataArgs{}
}

func newGetAuthStateDataResult() interface{} {
	return &GetAuthStateDataResult{}
}

type GetAuthStateDataArgs struct {
	Req *authsession.TLAuthsessionGetAuthStateData
}

func (p *GetAuthStateDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAuthStateDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAuthStateDataArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionGetAuthStateData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAuthStateDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAuthStateDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAuthStateDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionGetAuthStateData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAuthStateDataArgs_Req_DEFAULT *authsession.TLAuthsessionGetAuthStateData

func (p *GetAuthStateDataArgs) GetReq() *authsession.TLAuthsessionGetAuthStateData {
	if !p.IsSetReq() {
		return GetAuthStateDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAuthStateDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAuthStateDataResult struct {
	Success *authsession.AuthKeyStateData
}

var GetAuthStateDataResult_Success_DEFAULT *authsession.AuthKeyStateData

func (p *GetAuthStateDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAuthStateDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAuthStateDataResult) Unmarshal(in []byte) error {
	msg := new(authsession.AuthKeyStateData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthStateDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAuthStateDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAuthStateDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.AuthKeyStateData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthStateDataResult) GetSuccess() *authsession.AuthKeyStateData {
	if !p.IsSetSuccess() {
		return GetAuthStateDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAuthStateDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*authsession.AuthKeyStateData)
}

func (p *GetAuthStateDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAuthStateDataResult) GetResult() interface{} {
	return p.Success
}

func setLayerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetLayerArgs)
	realResult := result.(*SetLayerResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionSetLayer(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetLayerArgs() interface{} {
	return &SetLayerArgs{}
}

func newSetLayerResult() interface{} {
	return &SetLayerResult{}
}

type SetLayerArgs struct {
	Req *authsession.TLAuthsessionSetLayer
}

func (p *SetLayerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetLayerArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetLayerArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionSetLayer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetLayerArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetLayerArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetLayerArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionSetLayer)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetLayerArgs_Req_DEFAULT *authsession.TLAuthsessionSetLayer

func (p *SetLayerArgs) GetReq() *authsession.TLAuthsessionSetLayer {
	if !p.IsSetReq() {
		return SetLayerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetLayerArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetLayerResult struct {
	Success *tg.Bool
}

var SetLayerResult_Success_DEFAULT *tg.Bool

func (p *SetLayerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetLayerResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetLayerResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetLayerResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetLayerResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetLayerResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetLayerResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetLayerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetLayerResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetLayerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetLayerResult) GetResult() interface{} {
	return p.Success
}

func setInitConnectionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetInitConnectionArgs)
	realResult := result.(*SetInitConnectionResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionSetInitConnection(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetInitConnectionArgs() interface{} {
	return &SetInitConnectionArgs{}
}

func newSetInitConnectionResult() interface{} {
	return &SetInitConnectionResult{}
}

type SetInitConnectionArgs struct {
	Req *authsession.TLAuthsessionSetInitConnection
}

func (p *SetInitConnectionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetInitConnectionArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetInitConnectionArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionSetInitConnection)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetInitConnectionArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetInitConnectionArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetInitConnectionArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionSetInitConnection)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetInitConnectionArgs_Req_DEFAULT *authsession.TLAuthsessionSetInitConnection

func (p *SetInitConnectionArgs) GetReq() *authsession.TLAuthsessionSetInitConnection {
	if !p.IsSetReq() {
		return SetInitConnectionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetInitConnectionArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetInitConnectionResult struct {
	Success *tg.Bool
}

var SetInitConnectionResult_Success_DEFAULT *tg.Bool

func (p *SetInitConnectionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetInitConnectionResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetInitConnectionResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetInitConnectionResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetInitConnectionResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetInitConnectionResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetInitConnectionResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetInitConnectionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetInitConnectionResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetInitConnectionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetInitConnectionResult) GetResult() interface{} {
	return p.Success
}

func setAndroidPushSessionIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetAndroidPushSessionIdArgs)
	realResult := result.(*SetAndroidPushSessionIdResult)
	success, err := handler.(authsession.RPCAuthsession).AuthsessionSetAndroidPushSessionId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetAndroidPushSessionIdArgs() interface{} {
	return &SetAndroidPushSessionIdArgs{}
}

func newSetAndroidPushSessionIdResult() interface{} {
	return &SetAndroidPushSessionIdResult{}
}

type SetAndroidPushSessionIdArgs struct {
	Req *authsession.TLAuthsessionSetAndroidPushSessionId
}

func (p *SetAndroidPushSessionIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetAndroidPushSessionIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetAndroidPushSessionIdArgs) Unmarshal(in []byte) error {
	msg := new(authsession.TLAuthsessionSetAndroidPushSessionId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetAndroidPushSessionIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetAndroidPushSessionIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetAndroidPushSessionIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(authsession.TLAuthsessionSetAndroidPushSessionId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetAndroidPushSessionIdArgs_Req_DEFAULT *authsession.TLAuthsessionSetAndroidPushSessionId

func (p *SetAndroidPushSessionIdArgs) GetReq() *authsession.TLAuthsessionSetAndroidPushSessionId {
	if !p.IsSetReq() {
		return SetAndroidPushSessionIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetAndroidPushSessionIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetAndroidPushSessionIdResult struct {
	Success *tg.Bool
}

var SetAndroidPushSessionIdResult_Success_DEFAULT *tg.Bool

func (p *SetAndroidPushSessionIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetAndroidPushSessionIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetAndroidPushSessionIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetAndroidPushSessionIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetAndroidPushSessionIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetAndroidPushSessionIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetAndroidPushSessionIdResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetAndroidPushSessionIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetAndroidPushSessionIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetAndroidPushSessionIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetAndroidPushSessionIdResult) GetResult() interface{} {
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

func (p *kClient) AuthsessionGetAuthorizations(ctx context.Context, req *authsession.TLAuthsessionGetAuthorizations) (r *tg.AccountAuthorizations, err error) {
	var _args GetAuthorizationsArgs
	_args.Req = req
	var _result GetAuthorizationsResult
	if err = p.c.Call(ctx, "authsession.getAuthorizations", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionResetAuthorization(ctx context.Context, req *authsession.TLAuthsessionResetAuthorization) (r *authsession.VectorLong, err error) {
	var _args ResetAuthorizationArgs
	_args.Req = req
	var _result ResetAuthorizationResult
	if err = p.c.Call(ctx, "authsession.resetAuthorization", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetLayer(ctx context.Context, req *authsession.TLAuthsessionGetLayer) (r *tg.Int32, err error) {
	var _args GetLayerArgs
	_args.Req = req
	var _result GetLayerResult
	if err = p.c.Call(ctx, "authsession.getLayer", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetLangPack(ctx context.Context, req *authsession.TLAuthsessionGetLangPack) (r *tg.String, err error) {
	var _args GetLangPackArgs
	_args.Req = req
	var _result GetLangPackResult
	if err = p.c.Call(ctx, "authsession.getLangPack", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetClient(ctx context.Context, req *authsession.TLAuthsessionGetClient) (r *tg.String, err error) {
	var _args GetClientArgs
	_args.Req = req
	var _result GetClientResult
	if err = p.c.Call(ctx, "authsession.getClient", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetLangCode(ctx context.Context, req *authsession.TLAuthsessionGetLangCode) (r *tg.String, err error) {
	var _args GetLangCodeArgs
	_args.Req = req
	var _result GetLangCodeResult
	if err = p.c.Call(ctx, "authsession.getLangCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetUserId(ctx context.Context, req *authsession.TLAuthsessionGetUserId) (r *tg.Int64, err error) {
	var _args GetUserIdArgs
	_args.Req = req
	var _result GetUserIdResult
	if err = p.c.Call(ctx, "authsession.getUserId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetPushSessionId(ctx context.Context, req *authsession.TLAuthsessionGetPushSessionId) (r *tg.Int64, err error) {
	var _args GetPushSessionIdArgs
	_args.Req = req
	var _result GetPushSessionIdResult
	if err = p.c.Call(ctx, "authsession.getPushSessionId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetFutureSalts(ctx context.Context, req *authsession.TLAuthsessionGetFutureSalts) (r *tg.FutureSalts, err error) {
	var _args GetFutureSaltsArgs
	_args.Req = req
	var _result GetFutureSaltsResult
	if err = p.c.Call(ctx, "authsession.getFutureSalts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionQueryAuthKey(ctx context.Context, req *authsession.TLAuthsessionQueryAuthKey) (r *tg.AuthKeyInfo, err error) {
	var _args QueryAuthKeyArgs
	_args.Req = req
	var _result QueryAuthKeyResult
	if err = p.c.Call(ctx, "authsession.queryAuthKey", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionSetAuthKey(ctx context.Context, req *authsession.TLAuthsessionSetAuthKey) (r *tg.Bool, err error) {
	var _args SetAuthKeyArgs
	_args.Req = req
	var _result SetAuthKeyResult
	if err = p.c.Call(ctx, "authsession.setAuthKey", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionBindAuthKeyUser(ctx context.Context, req *authsession.TLAuthsessionBindAuthKeyUser) (r *tg.Int64, err error) {
	var _args BindAuthKeyUserArgs
	_args.Req = req
	var _result BindAuthKeyUserResult
	if err = p.c.Call(ctx, "authsession.bindAuthKeyUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionUnbindAuthKeyUser(ctx context.Context, req *authsession.TLAuthsessionUnbindAuthKeyUser) (r *tg.Bool, err error) {
	var _args UnbindAuthKeyUserArgs
	_args.Req = req
	var _result UnbindAuthKeyUserResult
	if err = p.c.Call(ctx, "authsession.unbindAuthKeyUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetPermAuthKeyId(ctx context.Context, req *authsession.TLAuthsessionGetPermAuthKeyId) (r *tg.Int64, err error) {
	var _args GetPermAuthKeyIdArgs
	_args.Req = req
	var _result GetPermAuthKeyIdResult
	if err = p.c.Call(ctx, "authsession.getPermAuthKeyId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionBindTempAuthKey(ctx context.Context, req *authsession.TLAuthsessionBindTempAuthKey) (r *tg.Bool, err error) {
	var _args BindTempAuthKeyArgs
	_args.Req = req
	var _result BindTempAuthKeyResult
	if err = p.c.Call(ctx, "authsession.bindTempAuthKey", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionSetClientSessionInfo(ctx context.Context, req *authsession.TLAuthsessionSetClientSessionInfo) (r *tg.Bool, err error) {
	var _args SetClientSessionInfoArgs
	_args.Req = req
	var _result SetClientSessionInfoResult
	if err = p.c.Call(ctx, "authsession.setClientSessionInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetAuthorization(ctx context.Context, req *authsession.TLAuthsessionGetAuthorization) (r *tg.Authorization, err error) {
	var _args GetAuthorizationArgs
	_args.Req = req
	var _result GetAuthorizationResult
	if err = p.c.Call(ctx, "authsession.getAuthorization", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionGetAuthStateData(ctx context.Context, req *authsession.TLAuthsessionGetAuthStateData) (r *authsession.AuthKeyStateData, err error) {
	var _args GetAuthStateDataArgs
	_args.Req = req
	var _result GetAuthStateDataResult
	if err = p.c.Call(ctx, "authsession.getAuthStateData", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionSetLayer(ctx context.Context, req *authsession.TLAuthsessionSetLayer) (r *tg.Bool, err error) {
	var _args SetLayerArgs
	_args.Req = req
	var _result SetLayerResult
	if err = p.c.Call(ctx, "authsession.setLayer", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionSetInitConnection(ctx context.Context, req *authsession.TLAuthsessionSetInitConnection) (r *tg.Bool, err error) {
	var _args SetInitConnectionArgs
	_args.Req = req
	var _result SetInitConnectionResult
	if err = p.c.Call(ctx, "authsession.setInitConnection", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthsessionSetAndroidPushSessionId(ctx context.Context, req *authsession.TLAuthsessionSetAndroidPushSessionId) (r *tg.Bool, err error) {
	var _args SetAndroidPushSessionIdArgs
	_args.Req = req
	var _result SetAndroidPushSessionIdResult
	if err = p.c.Call(ctx, "authsession.setAndroidPushSessionId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
