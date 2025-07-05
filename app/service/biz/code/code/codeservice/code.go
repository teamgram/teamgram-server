/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package codeservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/code.RPCCode/code.createPhoneCode": kitex.NewMethodInfo(
		createPhoneCodeHandler,
		newCreatePhoneCodeArgs,
		newCreatePhoneCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/code.RPCCode/code.getPhoneCode": kitex.NewMethodInfo(
		getPhoneCodeHandler,
		newGetPhoneCodeArgs,
		newGetPhoneCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/code.RPCCode/code.deletePhoneCode": kitex.NewMethodInfo(
		deletePhoneCodeHandler,
		newDeletePhoneCodeArgs,
		newDeletePhoneCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/code.RPCCode/code.updatePhoneCodeData": kitex.NewMethodInfo(
		updatePhoneCodeDataHandler,
		newUpdatePhoneCodeDataArgs,
		newUpdatePhoneCodeDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	codeServiceServiceInfo                = NewServiceInfo()
	codeServiceServiceInfoForClient       = NewServiceInfoForClient()
	codeServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCCode", codeServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCCode", codeServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCCode", codeServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return codeServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return codeServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return codeServiceServiceInfoForClient
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
	serviceName := "RPCCode"
	handlerType := (*code.RPCCode)(nil)
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
		"PackageName": "code",
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

func createPhoneCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CreatePhoneCodeArgs)
	realResult := result.(*CreatePhoneCodeResult)
	success, err := handler.(code.RPCCode).CodeCreatePhoneCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCreatePhoneCodeArgs() interface{} {
	return &CreatePhoneCodeArgs{}
}

func newCreatePhoneCodeResult() interface{} {
	return &CreatePhoneCodeResult{}
}

type CreatePhoneCodeArgs struct {
	Req *code.TLCodeCreatePhoneCode
}

func (p *CreatePhoneCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreatePhoneCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CreatePhoneCodeArgs) Unmarshal(in []byte) error {
	msg := new(code.TLCodeCreatePhoneCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CreatePhoneCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CreatePhoneCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CreatePhoneCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(code.TLCodeCreatePhoneCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CreatePhoneCodeArgs_Req_DEFAULT *code.TLCodeCreatePhoneCode

func (p *CreatePhoneCodeArgs) GetReq() *code.TLCodeCreatePhoneCode {
	if !p.IsSetReq() {
		return CreatePhoneCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreatePhoneCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreatePhoneCodeResult struct {
	Success *code.PhoneCodeTransaction
}

var CreatePhoneCodeResult_Success_DEFAULT *code.PhoneCodeTransaction

func (p *CreatePhoneCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreatePhoneCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *CreatePhoneCodeResult) Unmarshal(in []byte) error {
	msg := new(code.PhoneCodeTransaction)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreatePhoneCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CreatePhoneCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CreatePhoneCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(code.PhoneCodeTransaction)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreatePhoneCodeResult) GetSuccess() *code.PhoneCodeTransaction {
	if !p.IsSetSuccess() {
		return CreatePhoneCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreatePhoneCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*code.PhoneCodeTransaction)
}

func (p *CreatePhoneCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreatePhoneCodeResult) GetResult() interface{} {
	return p.Success
}

func getPhoneCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPhoneCodeArgs)
	realResult := result.(*GetPhoneCodeResult)
	success, err := handler.(code.RPCCode).CodeGetPhoneCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPhoneCodeArgs() interface{} {
	return &GetPhoneCodeArgs{}
}

func newGetPhoneCodeResult() interface{} {
	return &GetPhoneCodeResult{}
}

type GetPhoneCodeArgs struct {
	Req *code.TLCodeGetPhoneCode
}

func (p *GetPhoneCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPhoneCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPhoneCodeArgs) Unmarshal(in []byte) error {
	msg := new(code.TLCodeGetPhoneCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPhoneCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPhoneCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPhoneCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(code.TLCodeGetPhoneCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPhoneCodeArgs_Req_DEFAULT *code.TLCodeGetPhoneCode

func (p *GetPhoneCodeArgs) GetReq() *code.TLCodeGetPhoneCode {
	if !p.IsSetReq() {
		return GetPhoneCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPhoneCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPhoneCodeResult struct {
	Success *code.PhoneCodeTransaction
}

var GetPhoneCodeResult_Success_DEFAULT *code.PhoneCodeTransaction

func (p *GetPhoneCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPhoneCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPhoneCodeResult) Unmarshal(in []byte) error {
	msg := new(code.PhoneCodeTransaction)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhoneCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPhoneCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPhoneCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(code.PhoneCodeTransaction)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhoneCodeResult) GetSuccess() *code.PhoneCodeTransaction {
	if !p.IsSetSuccess() {
		return GetPhoneCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPhoneCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*code.PhoneCodeTransaction)
}

func (p *GetPhoneCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPhoneCodeResult) GetResult() interface{} {
	return p.Success
}

func deletePhoneCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeletePhoneCodeArgs)
	realResult := result.(*DeletePhoneCodeResult)
	success, err := handler.(code.RPCCode).CodeDeletePhoneCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeletePhoneCodeArgs() interface{} {
	return &DeletePhoneCodeArgs{}
}

func newDeletePhoneCodeResult() interface{} {
	return &DeletePhoneCodeResult{}
}

type DeletePhoneCodeArgs struct {
	Req *code.TLCodeDeletePhoneCode
}

func (p *DeletePhoneCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeletePhoneCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeletePhoneCodeArgs) Unmarshal(in []byte) error {
	msg := new(code.TLCodeDeletePhoneCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeletePhoneCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeletePhoneCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeletePhoneCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(code.TLCodeDeletePhoneCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeletePhoneCodeArgs_Req_DEFAULT *code.TLCodeDeletePhoneCode

func (p *DeletePhoneCodeArgs) GetReq() *code.TLCodeDeletePhoneCode {
	if !p.IsSetReq() {
		return DeletePhoneCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePhoneCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeletePhoneCodeResult struct {
	Success *tg.Bool
}

var DeletePhoneCodeResult_Success_DEFAULT *tg.Bool

func (p *DeletePhoneCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeletePhoneCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeletePhoneCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePhoneCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeletePhoneCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeletePhoneCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePhoneCodeResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeletePhoneCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePhoneCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeletePhoneCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeletePhoneCodeResult) GetResult() interface{} {
	return p.Success
}

func updatePhoneCodeDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatePhoneCodeDataArgs)
	realResult := result.(*UpdatePhoneCodeDataResult)
	success, err := handler.(code.RPCCode).CodeUpdatePhoneCodeData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatePhoneCodeDataArgs() interface{} {
	return &UpdatePhoneCodeDataArgs{}
}

func newUpdatePhoneCodeDataResult() interface{} {
	return &UpdatePhoneCodeDataResult{}
}

type UpdatePhoneCodeDataArgs struct {
	Req *code.TLCodeUpdatePhoneCodeData
}

func (p *UpdatePhoneCodeDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatePhoneCodeDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatePhoneCodeDataArgs) Unmarshal(in []byte) error {
	msg := new(code.TLCodeUpdatePhoneCodeData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatePhoneCodeDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatePhoneCodeDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatePhoneCodeDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(code.TLCodeUpdatePhoneCodeData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatePhoneCodeDataArgs_Req_DEFAULT *code.TLCodeUpdatePhoneCodeData

func (p *UpdatePhoneCodeDataArgs) GetReq() *code.TLCodeUpdatePhoneCodeData {
	if !p.IsSetReq() {
		return UpdatePhoneCodeDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePhoneCodeDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatePhoneCodeDataResult struct {
	Success *tg.Bool
}

var UpdatePhoneCodeDataResult_Success_DEFAULT *tg.Bool

func (p *UpdatePhoneCodeDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatePhoneCodeDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatePhoneCodeDataResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePhoneCodeDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatePhoneCodeDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatePhoneCodeDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePhoneCodeDataResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdatePhoneCodeDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePhoneCodeDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdatePhoneCodeDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePhoneCodeDataResult) GetResult() interface{} {
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

func (p *kClient) CodeCreatePhoneCode(ctx context.Context, req *code.TLCodeCreatePhoneCode) (r *code.PhoneCodeTransaction, err error) {
	// var _args CreatePhoneCodeArgs
	// _args.Req = req
	// var _result CreatePhoneCodeResult

	_result := new(code.PhoneCodeTransaction)

	if err = p.c.Call(ctx, "/code.RPCCode/code.createPhoneCode", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) CodeGetPhoneCode(ctx context.Context, req *code.TLCodeGetPhoneCode) (r *code.PhoneCodeTransaction, err error) {
	// var _args GetPhoneCodeArgs
	// _args.Req = req
	// var _result GetPhoneCodeResult

	_result := new(code.PhoneCodeTransaction)

	if err = p.c.Call(ctx, "/code.RPCCode/code.getPhoneCode", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) CodeDeletePhoneCode(ctx context.Context, req *code.TLCodeDeletePhoneCode) (r *tg.Bool, err error) {
	// var _args DeletePhoneCodeArgs
	// _args.Req = req
	// var _result DeletePhoneCodeResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/code.RPCCode/code.deletePhoneCode", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) CodeUpdatePhoneCodeData(ctx context.Context, req *code.TLCodeUpdatePhoneCodeData) (r *tg.Bool, err error) {
	// var _args UpdatePhoneCodeDataArgs
	// _args.Req = req
	// var _result UpdatePhoneCodeDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/code.RPCCode/code.updatePhoneCodeData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
