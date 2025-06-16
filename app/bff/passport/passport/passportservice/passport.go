/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package passportservice

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
	"account.getAuthorizations": kitex.NewMethodInfo(
		accountGetAuthorizationsHandler,
		newAccountGetAuthorizationsArgs,
		newAccountGetAuthorizationsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.getAllSecureValues": kitex.NewMethodInfo(
		accountGetAllSecureValuesHandler,
		newAccountGetAllSecureValuesArgs,
		newAccountGetAllSecureValuesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.getSecureValue": kitex.NewMethodInfo(
		accountGetSecureValueHandler,
		newAccountGetSecureValueArgs,
		newAccountGetSecureValueResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.saveSecureValue": kitex.NewMethodInfo(
		accountSaveSecureValueHandler,
		newAccountSaveSecureValueArgs,
		newAccountSaveSecureValueResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.deleteSecureValue": kitex.NewMethodInfo(
		accountDeleteSecureValueHandler,
		newAccountDeleteSecureValueArgs,
		newAccountDeleteSecureValueResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.getAuthorizationForm": kitex.NewMethodInfo(
		accountGetAuthorizationFormHandler,
		newAccountGetAuthorizationFormArgs,
		newAccountGetAuthorizationFormResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.acceptAuthorization": kitex.NewMethodInfo(
		accountAcceptAuthorizationHandler,
		newAccountAcceptAuthorizationArgs,
		newAccountAcceptAuthorizationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.sendVerifyPhoneCode": kitex.NewMethodInfo(
		accountSendVerifyPhoneCodeHandler,
		newAccountSendVerifyPhoneCodeArgs,
		newAccountSendVerifyPhoneCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.verifyPhone": kitex.NewMethodInfo(
		accountVerifyPhoneHandler,
		newAccountVerifyPhoneArgs,
		newAccountVerifyPhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"users.setSecureValueErrors": kitex.NewMethodInfo(
		usersSetSecureValueErrorsHandler,
		newUsersSetSecureValueErrorsArgs,
		newUsersSetSecureValueErrorsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.getPassportConfig": kitex.NewMethodInfo(
		helpGetPassportConfigHandler,
		newHelpGetPassportConfigArgs,
		newHelpGetPassportConfigResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	passportServiceServiceInfo                = NewServiceInfo()
	passportServiceServiceInfoForClient       = NewServiceInfoForClient()
	passportServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return passportServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return passportServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return passportServiceServiceInfoForClient
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
	serviceName := "RPCPassport"
	handlerType := (*tg.RPCPassport)(nil)
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
		"PackageName": "passport",
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

func accountGetAuthorizationsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetAuthorizationsArgs)
	realResult := result.(*AccountGetAuthorizationsResult)
	success, err := handler.(tg.RPCPassport).AccountGetAuthorizations(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetAuthorizationsArgs() interface{} {
	return &AccountGetAuthorizationsArgs{}
}

func newAccountGetAuthorizationsResult() interface{} {
	return &AccountGetAuthorizationsResult{}
}

type AccountGetAuthorizationsArgs struct {
	Req *tg.TLAccountGetAuthorizations
}

func (p *AccountGetAuthorizationsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetAuthorizationsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetAuthorizationsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetAuthorizations)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetAuthorizationsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetAuthorizationsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetAuthorizationsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetAuthorizations)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetAuthorizationsArgs_Req_DEFAULT *tg.TLAccountGetAuthorizations

func (p *AccountGetAuthorizationsArgs) GetReq() *tg.TLAccountGetAuthorizations {
	if !p.IsSetReq() {
		return AccountGetAuthorizationsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetAuthorizationsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetAuthorizationsResult struct {
	Success *tg.AccountAuthorizations
}

var AccountGetAuthorizationsResult_Success_DEFAULT *tg.AccountAuthorizations

func (p *AccountGetAuthorizationsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetAuthorizationsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetAuthorizationsResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountAuthorizations)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAuthorizationsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetAuthorizationsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetAuthorizationsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountAuthorizations)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAuthorizationsResult) GetSuccess() *tg.AccountAuthorizations {
	if !p.IsSetSuccess() {
		return AccountGetAuthorizationsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetAuthorizationsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountAuthorizations)
}

func (p *AccountGetAuthorizationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetAuthorizationsResult) GetResult() interface{} {
	return p.Success
}

func accountGetAllSecureValuesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetAllSecureValuesArgs)
	realResult := result.(*AccountGetAllSecureValuesResult)
	success, err := handler.(tg.RPCPassport).AccountGetAllSecureValues(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetAllSecureValuesArgs() interface{} {
	return &AccountGetAllSecureValuesArgs{}
}

func newAccountGetAllSecureValuesResult() interface{} {
	return &AccountGetAllSecureValuesResult{}
}

type AccountGetAllSecureValuesArgs struct {
	Req *tg.TLAccountGetAllSecureValues
}

func (p *AccountGetAllSecureValuesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetAllSecureValuesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetAllSecureValuesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetAllSecureValues)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetAllSecureValuesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetAllSecureValuesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetAllSecureValuesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetAllSecureValues)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetAllSecureValuesArgs_Req_DEFAULT *tg.TLAccountGetAllSecureValues

func (p *AccountGetAllSecureValuesArgs) GetReq() *tg.TLAccountGetAllSecureValues {
	if !p.IsSetReq() {
		return AccountGetAllSecureValuesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetAllSecureValuesArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetAllSecureValuesResult struct {
	Success *tg.VectorSecureValue
}

var AccountGetAllSecureValuesResult_Success_DEFAULT *tg.VectorSecureValue

func (p *AccountGetAllSecureValuesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetAllSecureValuesResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetAllSecureValuesResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorSecureValue)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAllSecureValuesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetAllSecureValuesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetAllSecureValuesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorSecureValue)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAllSecureValuesResult) GetSuccess() *tg.VectorSecureValue {
	if !p.IsSetSuccess() {
		return AccountGetAllSecureValuesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetAllSecureValuesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorSecureValue)
}

func (p *AccountGetAllSecureValuesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetAllSecureValuesResult) GetResult() interface{} {
	return p.Success
}

func accountGetSecureValueHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetSecureValueArgs)
	realResult := result.(*AccountGetSecureValueResult)
	success, err := handler.(tg.RPCPassport).AccountGetSecureValue(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetSecureValueArgs() interface{} {
	return &AccountGetSecureValueArgs{}
}

func newAccountGetSecureValueResult() interface{} {
	return &AccountGetSecureValueResult{}
}

type AccountGetSecureValueArgs struct {
	Req *tg.TLAccountGetSecureValue
}

func (p *AccountGetSecureValueArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetSecureValueArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetSecureValueArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetSecureValue)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetSecureValueArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetSecureValueArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetSecureValueArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetSecureValue)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetSecureValueArgs_Req_DEFAULT *tg.TLAccountGetSecureValue

func (p *AccountGetSecureValueArgs) GetReq() *tg.TLAccountGetSecureValue {
	if !p.IsSetReq() {
		return AccountGetSecureValueArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetSecureValueArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetSecureValueResult struct {
	Success *tg.VectorSecureValue
}

var AccountGetSecureValueResult_Success_DEFAULT *tg.VectorSecureValue

func (p *AccountGetSecureValueResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetSecureValueResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetSecureValueResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorSecureValue)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetSecureValueResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetSecureValueResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetSecureValueResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorSecureValue)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetSecureValueResult) GetSuccess() *tg.VectorSecureValue {
	if !p.IsSetSuccess() {
		return AccountGetSecureValueResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetSecureValueResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorSecureValue)
}

func (p *AccountGetSecureValueResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetSecureValueResult) GetResult() interface{} {
	return p.Success
}

func accountSaveSecureValueHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSaveSecureValueArgs)
	realResult := result.(*AccountSaveSecureValueResult)
	success, err := handler.(tg.RPCPassport).AccountSaveSecureValue(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSaveSecureValueArgs() interface{} {
	return &AccountSaveSecureValueArgs{}
}

func newAccountSaveSecureValueResult() interface{} {
	return &AccountSaveSecureValueResult{}
}

type AccountSaveSecureValueArgs struct {
	Req *tg.TLAccountSaveSecureValue
}

func (p *AccountSaveSecureValueArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSaveSecureValueArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSaveSecureValueArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSaveSecureValue)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSaveSecureValueArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSaveSecureValueArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSaveSecureValueArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSaveSecureValue)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSaveSecureValueArgs_Req_DEFAULT *tg.TLAccountSaveSecureValue

func (p *AccountSaveSecureValueArgs) GetReq() *tg.TLAccountSaveSecureValue {
	if !p.IsSetReq() {
		return AccountSaveSecureValueArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSaveSecureValueArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSaveSecureValueResult struct {
	Success *tg.SecureValue
}

var AccountSaveSecureValueResult_Success_DEFAULT *tg.SecureValue

func (p *AccountSaveSecureValueResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSaveSecureValueResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSaveSecureValueResult) Unmarshal(in []byte) error {
	msg := new(tg.SecureValue)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSaveSecureValueResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSaveSecureValueResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSaveSecureValueResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.SecureValue)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSaveSecureValueResult) GetSuccess() *tg.SecureValue {
	if !p.IsSetSuccess() {
		return AccountSaveSecureValueResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSaveSecureValueResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.SecureValue)
}

func (p *AccountSaveSecureValueResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSaveSecureValueResult) GetResult() interface{} {
	return p.Success
}

func accountDeleteSecureValueHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountDeleteSecureValueArgs)
	realResult := result.(*AccountDeleteSecureValueResult)
	success, err := handler.(tg.RPCPassport).AccountDeleteSecureValue(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountDeleteSecureValueArgs() interface{} {
	return &AccountDeleteSecureValueArgs{}
}

func newAccountDeleteSecureValueResult() interface{} {
	return &AccountDeleteSecureValueResult{}
}

type AccountDeleteSecureValueArgs struct {
	Req *tg.TLAccountDeleteSecureValue
}

func (p *AccountDeleteSecureValueArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountDeleteSecureValueArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountDeleteSecureValueArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountDeleteSecureValue)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountDeleteSecureValueArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountDeleteSecureValueArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountDeleteSecureValueArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountDeleteSecureValue)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountDeleteSecureValueArgs_Req_DEFAULT *tg.TLAccountDeleteSecureValue

func (p *AccountDeleteSecureValueArgs) GetReq() *tg.TLAccountDeleteSecureValue {
	if !p.IsSetReq() {
		return AccountDeleteSecureValueArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountDeleteSecureValueArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountDeleteSecureValueResult struct {
	Success *tg.Bool
}

var AccountDeleteSecureValueResult_Success_DEFAULT *tg.Bool

func (p *AccountDeleteSecureValueResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountDeleteSecureValueResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountDeleteSecureValueResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountDeleteSecureValueResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountDeleteSecureValueResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountDeleteSecureValueResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountDeleteSecureValueResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountDeleteSecureValueResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountDeleteSecureValueResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountDeleteSecureValueResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountDeleteSecureValueResult) GetResult() interface{} {
	return p.Success
}

func accountGetAuthorizationFormHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetAuthorizationFormArgs)
	realResult := result.(*AccountGetAuthorizationFormResult)
	success, err := handler.(tg.RPCPassport).AccountGetAuthorizationForm(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetAuthorizationFormArgs() interface{} {
	return &AccountGetAuthorizationFormArgs{}
}

func newAccountGetAuthorizationFormResult() interface{} {
	return &AccountGetAuthorizationFormResult{}
}

type AccountGetAuthorizationFormArgs struct {
	Req *tg.TLAccountGetAuthorizationForm
}

func (p *AccountGetAuthorizationFormArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetAuthorizationFormArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetAuthorizationFormArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetAuthorizationForm)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetAuthorizationFormArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetAuthorizationFormArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetAuthorizationFormArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetAuthorizationForm)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetAuthorizationFormArgs_Req_DEFAULT *tg.TLAccountGetAuthorizationForm

func (p *AccountGetAuthorizationFormArgs) GetReq() *tg.TLAccountGetAuthorizationForm {
	if !p.IsSetReq() {
		return AccountGetAuthorizationFormArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetAuthorizationFormArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetAuthorizationFormResult struct {
	Success *tg.AccountAuthorizationForm
}

var AccountGetAuthorizationFormResult_Success_DEFAULT *tg.AccountAuthorizationForm

func (p *AccountGetAuthorizationFormResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetAuthorizationFormResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetAuthorizationFormResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountAuthorizationForm)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAuthorizationFormResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetAuthorizationFormResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetAuthorizationFormResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountAuthorizationForm)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAuthorizationFormResult) GetSuccess() *tg.AccountAuthorizationForm {
	if !p.IsSetSuccess() {
		return AccountGetAuthorizationFormResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetAuthorizationFormResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountAuthorizationForm)
}

func (p *AccountGetAuthorizationFormResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetAuthorizationFormResult) GetResult() interface{} {
	return p.Success
}

func accountAcceptAuthorizationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountAcceptAuthorizationArgs)
	realResult := result.(*AccountAcceptAuthorizationResult)
	success, err := handler.(tg.RPCPassport).AccountAcceptAuthorization(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountAcceptAuthorizationArgs() interface{} {
	return &AccountAcceptAuthorizationArgs{}
}

func newAccountAcceptAuthorizationResult() interface{} {
	return &AccountAcceptAuthorizationResult{}
}

type AccountAcceptAuthorizationArgs struct {
	Req *tg.TLAccountAcceptAuthorization
}

func (p *AccountAcceptAuthorizationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountAcceptAuthorizationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountAcceptAuthorizationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountAcceptAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountAcceptAuthorizationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountAcceptAuthorizationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountAcceptAuthorizationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountAcceptAuthorization)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountAcceptAuthorizationArgs_Req_DEFAULT *tg.TLAccountAcceptAuthorization

func (p *AccountAcceptAuthorizationArgs) GetReq() *tg.TLAccountAcceptAuthorization {
	if !p.IsSetReq() {
		return AccountAcceptAuthorizationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountAcceptAuthorizationArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountAcceptAuthorizationResult struct {
	Success *tg.Bool
}

var AccountAcceptAuthorizationResult_Success_DEFAULT *tg.Bool

func (p *AccountAcceptAuthorizationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountAcceptAuthorizationResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountAcceptAuthorizationResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountAcceptAuthorizationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountAcceptAuthorizationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountAcceptAuthorizationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountAcceptAuthorizationResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountAcceptAuthorizationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountAcceptAuthorizationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountAcceptAuthorizationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountAcceptAuthorizationResult) GetResult() interface{} {
	return p.Success
}

func accountSendVerifyPhoneCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSendVerifyPhoneCodeArgs)
	realResult := result.(*AccountSendVerifyPhoneCodeResult)
	success, err := handler.(tg.RPCPassport).AccountSendVerifyPhoneCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSendVerifyPhoneCodeArgs() interface{} {
	return &AccountSendVerifyPhoneCodeArgs{}
}

func newAccountSendVerifyPhoneCodeResult() interface{} {
	return &AccountSendVerifyPhoneCodeResult{}
}

type AccountSendVerifyPhoneCodeArgs struct {
	Req *tg.TLAccountSendVerifyPhoneCode
}

func (p *AccountSendVerifyPhoneCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSendVerifyPhoneCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSendVerifyPhoneCodeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSendVerifyPhoneCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSendVerifyPhoneCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSendVerifyPhoneCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSendVerifyPhoneCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSendVerifyPhoneCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSendVerifyPhoneCodeArgs_Req_DEFAULT *tg.TLAccountSendVerifyPhoneCode

func (p *AccountSendVerifyPhoneCodeArgs) GetReq() *tg.TLAccountSendVerifyPhoneCode {
	if !p.IsSetReq() {
		return AccountSendVerifyPhoneCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSendVerifyPhoneCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSendVerifyPhoneCodeResult struct {
	Success *tg.AuthSentCode
}

var AccountSendVerifyPhoneCodeResult_Success_DEFAULT *tg.AuthSentCode

func (p *AccountSendVerifyPhoneCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSendVerifyPhoneCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSendVerifyPhoneCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthSentCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSendVerifyPhoneCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSendVerifyPhoneCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSendVerifyPhoneCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthSentCode)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSendVerifyPhoneCodeResult) GetSuccess() *tg.AuthSentCode {
	if !p.IsSetSuccess() {
		return AccountSendVerifyPhoneCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSendVerifyPhoneCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthSentCode)
}

func (p *AccountSendVerifyPhoneCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSendVerifyPhoneCodeResult) GetResult() interface{} {
	return p.Success
}

func accountVerifyPhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountVerifyPhoneArgs)
	realResult := result.(*AccountVerifyPhoneResult)
	success, err := handler.(tg.RPCPassport).AccountVerifyPhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountVerifyPhoneArgs() interface{} {
	return &AccountVerifyPhoneArgs{}
}

func newAccountVerifyPhoneResult() interface{} {
	return &AccountVerifyPhoneResult{}
}

type AccountVerifyPhoneArgs struct {
	Req *tg.TLAccountVerifyPhone
}

func (p *AccountVerifyPhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountVerifyPhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountVerifyPhoneArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountVerifyPhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountVerifyPhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountVerifyPhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountVerifyPhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountVerifyPhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountVerifyPhoneArgs_Req_DEFAULT *tg.TLAccountVerifyPhone

func (p *AccountVerifyPhoneArgs) GetReq() *tg.TLAccountVerifyPhone {
	if !p.IsSetReq() {
		return AccountVerifyPhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountVerifyPhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountVerifyPhoneResult struct {
	Success *tg.Bool
}

var AccountVerifyPhoneResult_Success_DEFAULT *tg.Bool

func (p *AccountVerifyPhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountVerifyPhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountVerifyPhoneResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountVerifyPhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountVerifyPhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountVerifyPhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountVerifyPhoneResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountVerifyPhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountVerifyPhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountVerifyPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountVerifyPhoneResult) GetResult() interface{} {
	return p.Success
}

func usersSetSecureValueErrorsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UsersSetSecureValueErrorsArgs)
	realResult := result.(*UsersSetSecureValueErrorsResult)
	success, err := handler.(tg.RPCPassport).UsersSetSecureValueErrors(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUsersSetSecureValueErrorsArgs() interface{} {
	return &UsersSetSecureValueErrorsArgs{}
}

func newUsersSetSecureValueErrorsResult() interface{} {
	return &UsersSetSecureValueErrorsResult{}
}

type UsersSetSecureValueErrorsArgs struct {
	Req *tg.TLUsersSetSecureValueErrors
}

func (p *UsersSetSecureValueErrorsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UsersSetSecureValueErrorsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UsersSetSecureValueErrorsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUsersSetSecureValueErrors)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UsersSetSecureValueErrorsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UsersSetSecureValueErrorsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UsersSetSecureValueErrorsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUsersSetSecureValueErrors)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UsersSetSecureValueErrorsArgs_Req_DEFAULT *tg.TLUsersSetSecureValueErrors

func (p *UsersSetSecureValueErrorsArgs) GetReq() *tg.TLUsersSetSecureValueErrors {
	if !p.IsSetReq() {
		return UsersSetSecureValueErrorsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UsersSetSecureValueErrorsArgs) IsSetReq() bool {
	return p.Req != nil
}

type UsersSetSecureValueErrorsResult struct {
	Success *tg.Bool
}

var UsersSetSecureValueErrorsResult_Success_DEFAULT *tg.Bool

func (p *UsersSetSecureValueErrorsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UsersSetSecureValueErrorsResult")
	}
	return json.Marshal(p.Success)
}

func (p *UsersSetSecureValueErrorsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersSetSecureValueErrorsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UsersSetSecureValueErrorsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UsersSetSecureValueErrorsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersSetSecureValueErrorsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UsersSetSecureValueErrorsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UsersSetSecureValueErrorsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UsersSetSecureValueErrorsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UsersSetSecureValueErrorsResult) GetResult() interface{} {
	return p.Success
}

func helpGetPassportConfigHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetPassportConfigArgs)
	realResult := result.(*HelpGetPassportConfigResult)
	success, err := handler.(tg.RPCPassport).HelpGetPassportConfig(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetPassportConfigArgs() interface{} {
	return &HelpGetPassportConfigArgs{}
}

func newHelpGetPassportConfigResult() interface{} {
	return &HelpGetPassportConfigResult{}
}

type HelpGetPassportConfigArgs struct {
	Req *tg.TLHelpGetPassportConfig
}

func (p *HelpGetPassportConfigArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetPassportConfigArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetPassportConfigArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetPassportConfig)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetPassportConfigArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetPassportConfigArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetPassportConfigArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetPassportConfig)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetPassportConfigArgs_Req_DEFAULT *tg.TLHelpGetPassportConfig

func (p *HelpGetPassportConfigArgs) GetReq() *tg.TLHelpGetPassportConfig {
	if !p.IsSetReq() {
		return HelpGetPassportConfigArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetPassportConfigArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetPassportConfigResult struct {
	Success *tg.HelpPassportConfig
}

var HelpGetPassportConfigResult_Success_DEFAULT *tg.HelpPassportConfig

func (p *HelpGetPassportConfigResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetPassportConfigResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetPassportConfigResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpPassportConfig)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetPassportConfigResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetPassportConfigResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetPassportConfigResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpPassportConfig)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetPassportConfigResult) GetSuccess() *tg.HelpPassportConfig {
	if !p.IsSetSuccess() {
		return HelpGetPassportConfigResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetPassportConfigResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpPassportConfig)
}

func (p *HelpGetPassportConfigResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetPassportConfigResult) GetResult() interface{} {
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

func (p *kClient) AccountGetAuthorizations(ctx context.Context, req *tg.TLAccountGetAuthorizations) (r *tg.AccountAuthorizations, err error) {
	var _args AccountGetAuthorizationsArgs
	_args.Req = req
	var _result AccountGetAuthorizationsResult
	if err = p.c.Call(ctx, "account.getAuthorizations", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountGetAllSecureValues(ctx context.Context, req *tg.TLAccountGetAllSecureValues) (r *tg.VectorSecureValue, err error) {
	var _args AccountGetAllSecureValuesArgs
	_args.Req = req
	var _result AccountGetAllSecureValuesResult
	if err = p.c.Call(ctx, "account.getAllSecureValues", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountGetSecureValue(ctx context.Context, req *tg.TLAccountGetSecureValue) (r *tg.VectorSecureValue, err error) {
	var _args AccountGetSecureValueArgs
	_args.Req = req
	var _result AccountGetSecureValueResult
	if err = p.c.Call(ctx, "account.getSecureValue", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountSaveSecureValue(ctx context.Context, req *tg.TLAccountSaveSecureValue) (r *tg.SecureValue, err error) {
	var _args AccountSaveSecureValueArgs
	_args.Req = req
	var _result AccountSaveSecureValueResult
	if err = p.c.Call(ctx, "account.saveSecureValue", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountDeleteSecureValue(ctx context.Context, req *tg.TLAccountDeleteSecureValue) (r *tg.Bool, err error) {
	var _args AccountDeleteSecureValueArgs
	_args.Req = req
	var _result AccountDeleteSecureValueResult
	if err = p.c.Call(ctx, "account.deleteSecureValue", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountGetAuthorizationForm(ctx context.Context, req *tg.TLAccountGetAuthorizationForm) (r *tg.AccountAuthorizationForm, err error) {
	var _args AccountGetAuthorizationFormArgs
	_args.Req = req
	var _result AccountGetAuthorizationFormResult
	if err = p.c.Call(ctx, "account.getAuthorizationForm", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountAcceptAuthorization(ctx context.Context, req *tg.TLAccountAcceptAuthorization) (r *tg.Bool, err error) {
	var _args AccountAcceptAuthorizationArgs
	_args.Req = req
	var _result AccountAcceptAuthorizationResult
	if err = p.c.Call(ctx, "account.acceptAuthorization", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountSendVerifyPhoneCode(ctx context.Context, req *tg.TLAccountSendVerifyPhoneCode) (r *tg.AuthSentCode, err error) {
	var _args AccountSendVerifyPhoneCodeArgs
	_args.Req = req
	var _result AccountSendVerifyPhoneCodeResult
	if err = p.c.Call(ctx, "account.sendVerifyPhoneCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountVerifyPhone(ctx context.Context, req *tg.TLAccountVerifyPhone) (r *tg.Bool, err error) {
	var _args AccountVerifyPhoneArgs
	_args.Req = req
	var _result AccountVerifyPhoneResult
	if err = p.c.Call(ctx, "account.verifyPhone", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UsersSetSecureValueErrors(ctx context.Context, req *tg.TLUsersSetSecureValueErrors) (r *tg.Bool, err error) {
	var _args UsersSetSecureValueErrorsArgs
	_args.Req = req
	var _result UsersSetSecureValueErrorsResult
	if err = p.c.Call(ctx, "users.setSecureValueErrors", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) HelpGetPassportConfig(ctx context.Context, req *tg.TLHelpGetPassportConfig) (r *tg.HelpPassportConfig, err error) {
	var _args HelpGetPassportConfigArgs
	_args.Req = req
	var _result HelpGetPassportConfigResult
	if err = p.c.Call(ctx, "help.getPassportConfig", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
