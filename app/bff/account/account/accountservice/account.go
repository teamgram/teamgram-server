/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package accountservice

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
	"account.deleteAccount": kitex.NewMethodInfo(
		accountDeleteAccountHandler,
		newAccountDeleteAccountArgs,
		newAccountDeleteAccountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.getAccountTTL": kitex.NewMethodInfo(
		accountGetAccountTTLHandler,
		newAccountGetAccountTTLArgs,
		newAccountGetAccountTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.setAccountTTL": kitex.NewMethodInfo(
		accountSetAccountTTLHandler,
		newAccountSetAccountTTLArgs,
		newAccountSetAccountTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.sendChangePhoneCode": kitex.NewMethodInfo(
		accountSendChangePhoneCodeHandler,
		newAccountSendChangePhoneCodeArgs,
		newAccountSendChangePhoneCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.changePhone": kitex.NewMethodInfo(
		accountChangePhoneHandler,
		newAccountChangePhoneArgs,
		newAccountChangePhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.resetAuthorization": kitex.NewMethodInfo(
		accountResetAuthorizationHandler,
		newAccountResetAuthorizationArgs,
		newAccountResetAuthorizationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.sendConfirmPhoneCode": kitex.NewMethodInfo(
		accountSendConfirmPhoneCodeHandler,
		newAccountSendConfirmPhoneCodeArgs,
		newAccountSendConfirmPhoneCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.confirmPhone": kitex.NewMethodInfo(
		accountConfirmPhoneHandler,
		newAccountConfirmPhoneArgs,
		newAccountConfirmPhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	accountServiceServiceInfo                = NewServiceInfo()
	accountServiceServiceInfoForClient       = NewServiceInfoForClient()
	accountServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCAccount", accountServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCAccount", accountServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCAccount", accountServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return accountServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return accountServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return accountServiceServiceInfoForClient
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
	serviceName := "RPCAccount"
	handlerType := (*tg.RPCAccount)(nil)
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
		"PackageName": "account",
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

func accountDeleteAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountDeleteAccountArgs)
	realResult := result.(*AccountDeleteAccountResult)
	success, err := handler.(tg.RPCAccount).AccountDeleteAccount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountDeleteAccountArgs() interface{} {
	return &AccountDeleteAccountArgs{}
}

func newAccountDeleteAccountResult() interface{} {
	return &AccountDeleteAccountResult{}
}

type AccountDeleteAccountArgs struct {
	Req *tg.TLAccountDeleteAccount
}

func (p *AccountDeleteAccountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountDeleteAccountArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountDeleteAccountArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountDeleteAccount)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountDeleteAccountArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountDeleteAccountArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountDeleteAccountArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountDeleteAccount)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountDeleteAccountArgs_Req_DEFAULT *tg.TLAccountDeleteAccount

func (p *AccountDeleteAccountArgs) GetReq() *tg.TLAccountDeleteAccount {
	if !p.IsSetReq() {
		return AccountDeleteAccountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountDeleteAccountArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountDeleteAccountResult struct {
	Success *tg.Bool
}

var AccountDeleteAccountResult_Success_DEFAULT *tg.Bool

func (p *AccountDeleteAccountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountDeleteAccountResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountDeleteAccountResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountDeleteAccountResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountDeleteAccountResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountDeleteAccountResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountDeleteAccountResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountDeleteAccountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountDeleteAccountResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountDeleteAccountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountDeleteAccountResult) GetResult() interface{} {
	return p.Success
}

func accountGetAccountTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetAccountTTLArgs)
	realResult := result.(*AccountGetAccountTTLResult)
	success, err := handler.(tg.RPCAccount).AccountGetAccountTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetAccountTTLArgs() interface{} {
	return &AccountGetAccountTTLArgs{}
}

func newAccountGetAccountTTLResult() interface{} {
	return &AccountGetAccountTTLResult{}
}

type AccountGetAccountTTLArgs struct {
	Req *tg.TLAccountGetAccountTTL
}

func (p *AccountGetAccountTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetAccountTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetAccountTTLArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetAccountTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetAccountTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetAccountTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetAccountTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetAccountTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetAccountTTLArgs_Req_DEFAULT *tg.TLAccountGetAccountTTL

func (p *AccountGetAccountTTLArgs) GetReq() *tg.TLAccountGetAccountTTL {
	if !p.IsSetReq() {
		return AccountGetAccountTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetAccountTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetAccountTTLResult struct {
	Success *tg.AccountDaysTTL
}

var AccountGetAccountTTLResult_Success_DEFAULT *tg.AccountDaysTTL

func (p *AccountGetAccountTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetAccountTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetAccountTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountDaysTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAccountTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetAccountTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetAccountTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountDaysTTL)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAccountTTLResult) GetSuccess() *tg.AccountDaysTTL {
	if !p.IsSetSuccess() {
		return AccountGetAccountTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetAccountTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountDaysTTL)
}

func (p *AccountGetAccountTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetAccountTTLResult) GetResult() interface{} {
	return p.Success
}

func accountSetAccountTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSetAccountTTLArgs)
	realResult := result.(*AccountSetAccountTTLResult)
	success, err := handler.(tg.RPCAccount).AccountSetAccountTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSetAccountTTLArgs() interface{} {
	return &AccountSetAccountTTLArgs{}
}

func newAccountSetAccountTTLResult() interface{} {
	return &AccountSetAccountTTLResult{}
}

type AccountSetAccountTTLArgs struct {
	Req *tg.TLAccountSetAccountTTL
}

func (p *AccountSetAccountTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSetAccountTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSetAccountTTLArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSetAccountTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSetAccountTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSetAccountTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSetAccountTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSetAccountTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSetAccountTTLArgs_Req_DEFAULT *tg.TLAccountSetAccountTTL

func (p *AccountSetAccountTTLArgs) GetReq() *tg.TLAccountSetAccountTTL {
	if !p.IsSetReq() {
		return AccountSetAccountTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSetAccountTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSetAccountTTLResult struct {
	Success *tg.Bool
}

var AccountSetAccountTTLResult_Success_DEFAULT *tg.Bool

func (p *AccountSetAccountTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSetAccountTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSetAccountTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetAccountTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSetAccountTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSetAccountTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetAccountTTLResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountSetAccountTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSetAccountTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountSetAccountTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSetAccountTTLResult) GetResult() interface{} {
	return p.Success
}

func accountSendChangePhoneCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSendChangePhoneCodeArgs)
	realResult := result.(*AccountSendChangePhoneCodeResult)
	success, err := handler.(tg.RPCAccount).AccountSendChangePhoneCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSendChangePhoneCodeArgs() interface{} {
	return &AccountSendChangePhoneCodeArgs{}
}

func newAccountSendChangePhoneCodeResult() interface{} {
	return &AccountSendChangePhoneCodeResult{}
}

type AccountSendChangePhoneCodeArgs struct {
	Req *tg.TLAccountSendChangePhoneCode
}

func (p *AccountSendChangePhoneCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSendChangePhoneCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSendChangePhoneCodeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSendChangePhoneCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSendChangePhoneCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSendChangePhoneCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSendChangePhoneCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSendChangePhoneCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSendChangePhoneCodeArgs_Req_DEFAULT *tg.TLAccountSendChangePhoneCode

func (p *AccountSendChangePhoneCodeArgs) GetReq() *tg.TLAccountSendChangePhoneCode {
	if !p.IsSetReq() {
		return AccountSendChangePhoneCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSendChangePhoneCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSendChangePhoneCodeResult struct {
	Success *tg.AuthSentCode
}

var AccountSendChangePhoneCodeResult_Success_DEFAULT *tg.AuthSentCode

func (p *AccountSendChangePhoneCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSendChangePhoneCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSendChangePhoneCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthSentCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSendChangePhoneCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSendChangePhoneCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSendChangePhoneCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthSentCode)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSendChangePhoneCodeResult) GetSuccess() *tg.AuthSentCode {
	if !p.IsSetSuccess() {
		return AccountSendChangePhoneCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSendChangePhoneCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthSentCode)
}

func (p *AccountSendChangePhoneCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSendChangePhoneCodeResult) GetResult() interface{} {
	return p.Success
}

func accountChangePhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountChangePhoneArgs)
	realResult := result.(*AccountChangePhoneResult)
	success, err := handler.(tg.RPCAccount).AccountChangePhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountChangePhoneArgs() interface{} {
	return &AccountChangePhoneArgs{}
}

func newAccountChangePhoneResult() interface{} {
	return &AccountChangePhoneResult{}
}

type AccountChangePhoneArgs struct {
	Req *tg.TLAccountChangePhone
}

func (p *AccountChangePhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountChangePhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountChangePhoneArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountChangePhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountChangePhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountChangePhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountChangePhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountChangePhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountChangePhoneArgs_Req_DEFAULT *tg.TLAccountChangePhone

func (p *AccountChangePhoneArgs) GetReq() *tg.TLAccountChangePhone {
	if !p.IsSetReq() {
		return AccountChangePhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountChangePhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountChangePhoneResult struct {
	Success *tg.User
}

var AccountChangePhoneResult_Success_DEFAULT *tg.User

func (p *AccountChangePhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountChangePhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountChangePhoneResult) Unmarshal(in []byte) error {
	msg := new(tg.User)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountChangePhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountChangePhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountChangePhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.User)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountChangePhoneResult) GetSuccess() *tg.User {
	if !p.IsSetSuccess() {
		return AccountChangePhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountChangePhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.User)
}

func (p *AccountChangePhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountChangePhoneResult) GetResult() interface{} {
	return p.Success
}

func accountResetAuthorizationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountResetAuthorizationArgs)
	realResult := result.(*AccountResetAuthorizationResult)
	success, err := handler.(tg.RPCAccount).AccountResetAuthorization(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountResetAuthorizationArgs() interface{} {
	return &AccountResetAuthorizationArgs{}
}

func newAccountResetAuthorizationResult() interface{} {
	return &AccountResetAuthorizationResult{}
}

type AccountResetAuthorizationArgs struct {
	Req *tg.TLAccountResetAuthorization
}

func (p *AccountResetAuthorizationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountResetAuthorizationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountResetAuthorizationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountResetAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountResetAuthorizationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountResetAuthorizationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountResetAuthorizationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountResetAuthorization)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountResetAuthorizationArgs_Req_DEFAULT *tg.TLAccountResetAuthorization

func (p *AccountResetAuthorizationArgs) GetReq() *tg.TLAccountResetAuthorization {
	if !p.IsSetReq() {
		return AccountResetAuthorizationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountResetAuthorizationArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountResetAuthorizationResult struct {
	Success *tg.Bool
}

var AccountResetAuthorizationResult_Success_DEFAULT *tg.Bool

func (p *AccountResetAuthorizationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountResetAuthorizationResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountResetAuthorizationResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountResetAuthorizationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountResetAuthorizationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountResetAuthorizationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountResetAuthorizationResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountResetAuthorizationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountResetAuthorizationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountResetAuthorizationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountResetAuthorizationResult) GetResult() interface{} {
	return p.Success
}

func accountSendConfirmPhoneCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSendConfirmPhoneCodeArgs)
	realResult := result.(*AccountSendConfirmPhoneCodeResult)
	success, err := handler.(tg.RPCAccount).AccountSendConfirmPhoneCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSendConfirmPhoneCodeArgs() interface{} {
	return &AccountSendConfirmPhoneCodeArgs{}
}

func newAccountSendConfirmPhoneCodeResult() interface{} {
	return &AccountSendConfirmPhoneCodeResult{}
}

type AccountSendConfirmPhoneCodeArgs struct {
	Req *tg.TLAccountSendConfirmPhoneCode
}

func (p *AccountSendConfirmPhoneCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSendConfirmPhoneCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSendConfirmPhoneCodeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSendConfirmPhoneCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSendConfirmPhoneCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSendConfirmPhoneCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSendConfirmPhoneCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSendConfirmPhoneCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSendConfirmPhoneCodeArgs_Req_DEFAULT *tg.TLAccountSendConfirmPhoneCode

func (p *AccountSendConfirmPhoneCodeArgs) GetReq() *tg.TLAccountSendConfirmPhoneCode {
	if !p.IsSetReq() {
		return AccountSendConfirmPhoneCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSendConfirmPhoneCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSendConfirmPhoneCodeResult struct {
	Success *tg.AuthSentCode
}

var AccountSendConfirmPhoneCodeResult_Success_DEFAULT *tg.AuthSentCode

func (p *AccountSendConfirmPhoneCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSendConfirmPhoneCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSendConfirmPhoneCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthSentCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSendConfirmPhoneCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSendConfirmPhoneCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSendConfirmPhoneCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthSentCode)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSendConfirmPhoneCodeResult) GetSuccess() *tg.AuthSentCode {
	if !p.IsSetSuccess() {
		return AccountSendConfirmPhoneCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSendConfirmPhoneCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthSentCode)
}

func (p *AccountSendConfirmPhoneCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSendConfirmPhoneCodeResult) GetResult() interface{} {
	return p.Success
}

func accountConfirmPhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountConfirmPhoneArgs)
	realResult := result.(*AccountConfirmPhoneResult)
	success, err := handler.(tg.RPCAccount).AccountConfirmPhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountConfirmPhoneArgs() interface{} {
	return &AccountConfirmPhoneArgs{}
}

func newAccountConfirmPhoneResult() interface{} {
	return &AccountConfirmPhoneResult{}
}

type AccountConfirmPhoneArgs struct {
	Req *tg.TLAccountConfirmPhone
}

func (p *AccountConfirmPhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountConfirmPhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountConfirmPhoneArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountConfirmPhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountConfirmPhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountConfirmPhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountConfirmPhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountConfirmPhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountConfirmPhoneArgs_Req_DEFAULT *tg.TLAccountConfirmPhone

func (p *AccountConfirmPhoneArgs) GetReq() *tg.TLAccountConfirmPhone {
	if !p.IsSetReq() {
		return AccountConfirmPhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountConfirmPhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountConfirmPhoneResult struct {
	Success *tg.Bool
}

var AccountConfirmPhoneResult_Success_DEFAULT *tg.Bool

func (p *AccountConfirmPhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountConfirmPhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountConfirmPhoneResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountConfirmPhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountConfirmPhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountConfirmPhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountConfirmPhoneResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountConfirmPhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountConfirmPhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountConfirmPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountConfirmPhoneResult) GetResult() interface{} {
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

func (p *kClient) AccountDeleteAccount(ctx context.Context, req *tg.TLAccountDeleteAccount) (r *tg.Bool, err error) {
	// var _args AccountDeleteAccountArgs
	// _args.Req = req
	// var _result AccountDeleteAccountResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.deleteAccount", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountGetAccountTTL(ctx context.Context, req *tg.TLAccountGetAccountTTL) (r *tg.AccountDaysTTL, err error) {
	// var _args AccountGetAccountTTLArgs
	// _args.Req = req
	// var _result AccountGetAccountTTLResult

	_result := new(tg.AccountDaysTTL)
	if err = p.c.Call(ctx, "account.getAccountTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountSetAccountTTL(ctx context.Context, req *tg.TLAccountSetAccountTTL) (r *tg.Bool, err error) {
	// var _args AccountSetAccountTTLArgs
	// _args.Req = req
	// var _result AccountSetAccountTTLResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.setAccountTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountSendChangePhoneCode(ctx context.Context, req *tg.TLAccountSendChangePhoneCode) (r *tg.AuthSentCode, err error) {
	// var _args AccountSendChangePhoneCodeArgs
	// _args.Req = req
	// var _result AccountSendChangePhoneCodeResult

	_result := new(tg.AuthSentCode)
	if err = p.c.Call(ctx, "account.sendChangePhoneCode", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountChangePhone(ctx context.Context, req *tg.TLAccountChangePhone) (r *tg.User, err error) {
	// var _args AccountChangePhoneArgs
	// _args.Req = req
	// var _result AccountChangePhoneResult

	_result := new(tg.User)
	if err = p.c.Call(ctx, "account.changePhone", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountResetAuthorization(ctx context.Context, req *tg.TLAccountResetAuthorization) (r *tg.Bool, err error) {
	// var _args AccountResetAuthorizationArgs
	// _args.Req = req
	// var _result AccountResetAuthorizationResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.resetAuthorization", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountSendConfirmPhoneCode(ctx context.Context, req *tg.TLAccountSendConfirmPhoneCode) (r *tg.AuthSentCode, err error) {
	// var _args AccountSendConfirmPhoneCodeArgs
	// _args.Req = req
	// var _result AccountSendConfirmPhoneCodeResult

	_result := new(tg.AuthSentCode)
	if err = p.c.Call(ctx, "account.sendConfirmPhoneCode", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountConfirmPhone(ctx context.Context, req *tg.TLAccountConfirmPhone) (r *tg.Bool, err error) {
	// var _args AccountConfirmPhoneArgs
	// _args.Req = req
	// var _result AccountConfirmPhoneResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.confirmPhone", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
