/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package qrcodeservice

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
	"/tg.RPCQrCode/auth.exportLoginToken": kitex.NewMethodInfo(
		authExportLoginTokenHandler,
		newAuthExportLoginTokenArgs,
		newAuthExportLoginTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCQrCode/auth.importLoginToken": kitex.NewMethodInfo(
		authImportLoginTokenHandler,
		newAuthImportLoginTokenArgs,
		newAuthImportLoginTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCQrCode/auth.acceptLoginToken": kitex.NewMethodInfo(
		authAcceptLoginTokenHandler,
		newAuthAcceptLoginTokenArgs,
		newAuthAcceptLoginTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	qrcodeServiceServiceInfo                = NewServiceInfo()
	qrcodeServiceServiceInfoForClient       = NewServiceInfoForClient()
	qrcodeServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCQrCode", qrcodeServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCQrCode", qrcodeServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCQrCode", qrcodeServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return qrcodeServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return qrcodeServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return qrcodeServiceServiceInfoForClient
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
	serviceName := "RPCQrCode"
	handlerType := (*tg.RPCQrCode)(nil)
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
		"PackageName": "qrcode",
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

func authExportLoginTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthExportLoginTokenArgs)
	realResult := result.(*AuthExportLoginTokenResult)
	success, err := handler.(tg.RPCQrCode).AuthExportLoginToken(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthExportLoginTokenArgs() interface{} {
	return &AuthExportLoginTokenArgs{}
}

func newAuthExportLoginTokenResult() interface{} {
	return &AuthExportLoginTokenResult{}
}

type AuthExportLoginTokenArgs struct {
	Req *tg.TLAuthExportLoginToken
}

func (p *AuthExportLoginTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthExportLoginTokenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthExportLoginTokenArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthExportLoginToken)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthExportLoginTokenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthExportLoginTokenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthExportLoginTokenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthExportLoginToken)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthExportLoginTokenArgs_Req_DEFAULT *tg.TLAuthExportLoginToken

func (p *AuthExportLoginTokenArgs) GetReq() *tg.TLAuthExportLoginToken {
	if !p.IsSetReq() {
		return AuthExportLoginTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthExportLoginTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthExportLoginTokenResult struct {
	Success *tg.AuthLoginToken
}

var AuthExportLoginTokenResult_Success_DEFAULT *tg.AuthLoginToken

func (p *AuthExportLoginTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthExportLoginTokenResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthExportLoginTokenResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthLoginToken)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthExportLoginTokenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthExportLoginTokenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthExportLoginTokenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthLoginToken)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthExportLoginTokenResult) GetSuccess() *tg.AuthLoginToken {
	if !p.IsSetSuccess() {
		return AuthExportLoginTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthExportLoginTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthLoginToken)
}

func (p *AuthExportLoginTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthExportLoginTokenResult) GetResult() interface{} {
	return p.Success
}

func authImportLoginTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthImportLoginTokenArgs)
	realResult := result.(*AuthImportLoginTokenResult)
	success, err := handler.(tg.RPCQrCode).AuthImportLoginToken(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthImportLoginTokenArgs() interface{} {
	return &AuthImportLoginTokenArgs{}
}

func newAuthImportLoginTokenResult() interface{} {
	return &AuthImportLoginTokenResult{}
}

type AuthImportLoginTokenArgs struct {
	Req *tg.TLAuthImportLoginToken
}

func (p *AuthImportLoginTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthImportLoginTokenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthImportLoginTokenArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthImportLoginToken)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthImportLoginTokenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthImportLoginTokenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthImportLoginTokenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthImportLoginToken)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthImportLoginTokenArgs_Req_DEFAULT *tg.TLAuthImportLoginToken

func (p *AuthImportLoginTokenArgs) GetReq() *tg.TLAuthImportLoginToken {
	if !p.IsSetReq() {
		return AuthImportLoginTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthImportLoginTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthImportLoginTokenResult struct {
	Success *tg.AuthLoginToken
}

var AuthImportLoginTokenResult_Success_DEFAULT *tg.AuthLoginToken

func (p *AuthImportLoginTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthImportLoginTokenResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthImportLoginTokenResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthLoginToken)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthImportLoginTokenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthImportLoginTokenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthImportLoginTokenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthLoginToken)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthImportLoginTokenResult) GetSuccess() *tg.AuthLoginToken {
	if !p.IsSetSuccess() {
		return AuthImportLoginTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthImportLoginTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthLoginToken)
}

func (p *AuthImportLoginTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthImportLoginTokenResult) GetResult() interface{} {
	return p.Success
}

func authAcceptLoginTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthAcceptLoginTokenArgs)
	realResult := result.(*AuthAcceptLoginTokenResult)
	success, err := handler.(tg.RPCQrCode).AuthAcceptLoginToken(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthAcceptLoginTokenArgs() interface{} {
	return &AuthAcceptLoginTokenArgs{}
}

func newAuthAcceptLoginTokenResult() interface{} {
	return &AuthAcceptLoginTokenResult{}
}

type AuthAcceptLoginTokenArgs struct {
	Req *tg.TLAuthAcceptLoginToken
}

func (p *AuthAcceptLoginTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthAcceptLoginTokenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthAcceptLoginTokenArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthAcceptLoginToken)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthAcceptLoginTokenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthAcceptLoginTokenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthAcceptLoginTokenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthAcceptLoginToken)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthAcceptLoginTokenArgs_Req_DEFAULT *tg.TLAuthAcceptLoginToken

func (p *AuthAcceptLoginTokenArgs) GetReq() *tg.TLAuthAcceptLoginToken {
	if !p.IsSetReq() {
		return AuthAcceptLoginTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthAcceptLoginTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthAcceptLoginTokenResult struct {
	Success *tg.Authorization
}

var AuthAcceptLoginTokenResult_Success_DEFAULT *tg.Authorization

func (p *AuthAcceptLoginTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthAcceptLoginTokenResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthAcceptLoginTokenResult) Unmarshal(in []byte) error {
	msg := new(tg.Authorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthAcceptLoginTokenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthAcceptLoginTokenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthAcceptLoginTokenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Authorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthAcceptLoginTokenResult) GetSuccess() *tg.Authorization {
	if !p.IsSetSuccess() {
		return AuthAcceptLoginTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthAcceptLoginTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Authorization)
}

func (p *AuthAcceptLoginTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthAcceptLoginTokenResult) GetResult() interface{} {
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

func (p *kClient) AuthExportLoginToken(ctx context.Context, req *tg.TLAuthExportLoginToken) (r *tg.AuthLoginToken, err error) {
	// var _args AuthExportLoginTokenArgs
	// _args.Req = req
	// var _result AuthExportLoginTokenResult

	_result := new(tg.AuthLoginToken)
	if err = p.c.Call(ctx, "/tg.RPCQrCode/auth.exportLoginToken", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AuthImportLoginToken(ctx context.Context, req *tg.TLAuthImportLoginToken) (r *tg.AuthLoginToken, err error) {
	// var _args AuthImportLoginTokenArgs
	// _args.Req = req
	// var _result AuthImportLoginTokenResult

	_result := new(tg.AuthLoginToken)
	if err = p.c.Call(ctx, "/tg.RPCQrCode/auth.importLoginToken", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AuthAcceptLoginToken(ctx context.Context, req *tg.TLAuthAcceptLoginToken) (r *tg.Authorization, err error) {
	// var _args AuthAcceptLoginTokenArgs
	// _args.Req = req
	// var _result AuthAcceptLoginTokenResult

	_result := new(tg.Authorization)
	if err = p.c.Call(ctx, "/tg.RPCQrCode/auth.acceptLoginToken", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
