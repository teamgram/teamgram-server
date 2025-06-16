/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package authorizationservice

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
	"auth.sendCode": kitex.NewMethodInfo(
		authSendCodeHandler,
		newAuthSendCodeArgs,
		newAuthSendCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.signUp": kitex.NewMethodInfo(
		authSignUpHandler,
		newAuthSignUpArgs,
		newAuthSignUpResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.signIn": kitex.NewMethodInfo(
		authSignInHandler,
		newAuthSignInArgs,
		newAuthSignInResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.logOut": kitex.NewMethodInfo(
		authLogOutHandler,
		newAuthLogOutArgs,
		newAuthLogOutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.resetAuthorizations": kitex.NewMethodInfo(
		authResetAuthorizationsHandler,
		newAuthResetAuthorizationsArgs,
		newAuthResetAuthorizationsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.exportAuthorization": kitex.NewMethodInfo(
		authExportAuthorizationHandler,
		newAuthExportAuthorizationArgs,
		newAuthExportAuthorizationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.importAuthorization": kitex.NewMethodInfo(
		authImportAuthorizationHandler,
		newAuthImportAuthorizationArgs,
		newAuthImportAuthorizationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.bindTempAuthKey": kitex.NewMethodInfo(
		authBindTempAuthKeyHandler,
		newAuthBindTempAuthKeyArgs,
		newAuthBindTempAuthKeyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.importBotAuthorization": kitex.NewMethodInfo(
		authImportBotAuthorizationHandler,
		newAuthImportBotAuthorizationArgs,
		newAuthImportBotAuthorizationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.checkPassword": kitex.NewMethodInfo(
		authCheckPasswordHandler,
		newAuthCheckPasswordArgs,
		newAuthCheckPasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.requestPasswordRecovery": kitex.NewMethodInfo(
		authRequestPasswordRecoveryHandler,
		newAuthRequestPasswordRecoveryArgs,
		newAuthRequestPasswordRecoveryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.recoverPassword": kitex.NewMethodInfo(
		authRecoverPasswordHandler,
		newAuthRecoverPasswordArgs,
		newAuthRecoverPasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.resendCode": kitex.NewMethodInfo(
		authResendCodeHandler,
		newAuthResendCodeArgs,
		newAuthResendCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.cancelCode": kitex.NewMethodInfo(
		authCancelCodeHandler,
		newAuthCancelCodeArgs,
		newAuthCancelCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.dropTempAuthKeys": kitex.NewMethodInfo(
		authDropTempAuthKeysHandler,
		newAuthDropTempAuthKeysArgs,
		newAuthDropTempAuthKeysResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.checkRecoveryPassword": kitex.NewMethodInfo(
		authCheckRecoveryPasswordHandler,
		newAuthCheckRecoveryPasswordArgs,
		newAuthCheckRecoveryPasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.importWebTokenAuthorization": kitex.NewMethodInfo(
		authImportWebTokenAuthorizationHandler,
		newAuthImportWebTokenAuthorizationArgs,
		newAuthImportWebTokenAuthorizationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.requestFirebaseSms": kitex.NewMethodInfo(
		authRequestFirebaseSmsHandler,
		newAuthRequestFirebaseSmsArgs,
		newAuthRequestFirebaseSmsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.resetLoginEmail": kitex.NewMethodInfo(
		authResetLoginEmailHandler,
		newAuthResetLoginEmailArgs,
		newAuthResetLoginEmailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.reportMissingCode": kitex.NewMethodInfo(
		authReportMissingCodeHandler,
		newAuthReportMissingCodeArgs,
		newAuthReportMissingCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.sendVerifyEmailCode": kitex.NewMethodInfo(
		accountSendVerifyEmailCodeHandler,
		newAccountSendVerifyEmailCodeArgs,
		newAccountSendVerifyEmailCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.verifyEmail": kitex.NewMethodInfo(
		accountVerifyEmailHandler,
		newAccountVerifyEmailArgs,
		newAccountVerifyEmailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.resetPassword": kitex.NewMethodInfo(
		accountResetPasswordHandler,
		newAccountResetPasswordArgs,
		newAccountResetPasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.setAuthorizationTTL": kitex.NewMethodInfo(
		accountSetAuthorizationTTLHandler,
		newAccountSetAuthorizationTTLArgs,
		newAccountSetAuthorizationTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.changeAuthorizationSettings": kitex.NewMethodInfo(
		accountChangeAuthorizationSettingsHandler,
		newAccountChangeAuthorizationSettingsArgs,
		newAccountChangeAuthorizationSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.invalidateSignInCodes": kitex.NewMethodInfo(
		accountInvalidateSignInCodesHandler,
		newAccountInvalidateSignInCodesArgs,
		newAccountInvalidateSignInCodesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"auth.toggleBan": kitex.NewMethodInfo(
		authToggleBanHandler,
		newAuthToggleBanArgs,
		newAuthToggleBanResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	authorizationServiceServiceInfo                = NewServiceInfo()
	authorizationServiceServiceInfoForClient       = NewServiceInfoForClient()
	authorizationServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return authorizationServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return authorizationServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return authorizationServiceServiceInfoForClient
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
	serviceName := "RPCAuthorization"
	handlerType := (*tg.RPCAuthorization)(nil)
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
		"PackageName": "authorization",
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

func authSendCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthSendCodeArgs)
	realResult := result.(*AuthSendCodeResult)
	success, err := handler.(tg.RPCAuthorization).AuthSendCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthSendCodeArgs() interface{} {
	return &AuthSendCodeArgs{}
}

func newAuthSendCodeResult() interface{} {
	return &AuthSendCodeResult{}
}

type AuthSendCodeArgs struct {
	Req *tg.TLAuthSendCode
}

func (p *AuthSendCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthSendCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthSendCodeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthSendCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthSendCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthSendCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthSendCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthSendCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthSendCodeArgs_Req_DEFAULT *tg.TLAuthSendCode

func (p *AuthSendCodeArgs) GetReq() *tg.TLAuthSendCode {
	if !p.IsSetReq() {
		return AuthSendCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthSendCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthSendCodeResult struct {
	Success *tg.AuthSentCode
}

var AuthSendCodeResult_Success_DEFAULT *tg.AuthSentCode

func (p *AuthSendCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthSendCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthSendCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthSentCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthSendCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthSendCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthSendCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthSentCode)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthSendCodeResult) GetSuccess() *tg.AuthSentCode {
	if !p.IsSetSuccess() {
		return AuthSendCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthSendCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthSentCode)
}

func (p *AuthSendCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthSendCodeResult) GetResult() interface{} {
	return p.Success
}

func authSignUpHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthSignUpArgs)
	realResult := result.(*AuthSignUpResult)
	success, err := handler.(tg.RPCAuthorization).AuthSignUp(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthSignUpArgs() interface{} {
	return &AuthSignUpArgs{}
}

func newAuthSignUpResult() interface{} {
	return &AuthSignUpResult{}
}

type AuthSignUpArgs struct {
	Req *tg.TLAuthSignUp
}

func (p *AuthSignUpArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthSignUpArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthSignUpArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthSignUp)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthSignUpArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthSignUpArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthSignUpArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthSignUp)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthSignUpArgs_Req_DEFAULT *tg.TLAuthSignUp

func (p *AuthSignUpArgs) GetReq() *tg.TLAuthSignUp {
	if !p.IsSetReq() {
		return AuthSignUpArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthSignUpArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthSignUpResult struct {
	Success *tg.AuthAuthorization
}

var AuthSignUpResult_Success_DEFAULT *tg.AuthAuthorization

func (p *AuthSignUpResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthSignUpResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthSignUpResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthSignUpResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthSignUpResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthSignUpResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthAuthorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthSignUpResult) GetSuccess() *tg.AuthAuthorization {
	if !p.IsSetSuccess() {
		return AuthSignUpResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthSignUpResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthAuthorization)
}

func (p *AuthSignUpResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthSignUpResult) GetResult() interface{} {
	return p.Success
}

func authSignInHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthSignInArgs)
	realResult := result.(*AuthSignInResult)
	success, err := handler.(tg.RPCAuthorization).AuthSignIn(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthSignInArgs() interface{} {
	return &AuthSignInArgs{}
}

func newAuthSignInResult() interface{} {
	return &AuthSignInResult{}
}

type AuthSignInArgs struct {
	Req *tg.TLAuthSignIn
}

func (p *AuthSignInArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthSignInArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthSignInArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthSignIn)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthSignInArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthSignInArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthSignInArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthSignIn)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthSignInArgs_Req_DEFAULT *tg.TLAuthSignIn

func (p *AuthSignInArgs) GetReq() *tg.TLAuthSignIn {
	if !p.IsSetReq() {
		return AuthSignInArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthSignInArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthSignInResult struct {
	Success *tg.AuthAuthorization
}

var AuthSignInResult_Success_DEFAULT *tg.AuthAuthorization

func (p *AuthSignInResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthSignInResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthSignInResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthSignInResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthSignInResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthSignInResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthAuthorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthSignInResult) GetSuccess() *tg.AuthAuthorization {
	if !p.IsSetSuccess() {
		return AuthSignInResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthSignInResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthAuthorization)
}

func (p *AuthSignInResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthSignInResult) GetResult() interface{} {
	return p.Success
}

func authLogOutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthLogOutArgs)
	realResult := result.(*AuthLogOutResult)
	success, err := handler.(tg.RPCAuthorization).AuthLogOut(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthLogOutArgs() interface{} {
	return &AuthLogOutArgs{}
}

func newAuthLogOutResult() interface{} {
	return &AuthLogOutResult{}
}

type AuthLogOutArgs struct {
	Req *tg.TLAuthLogOut
}

func (p *AuthLogOutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthLogOutArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthLogOutArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthLogOut)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthLogOutArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthLogOutArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthLogOutArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthLogOut)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthLogOutArgs_Req_DEFAULT *tg.TLAuthLogOut

func (p *AuthLogOutArgs) GetReq() *tg.TLAuthLogOut {
	if !p.IsSetReq() {
		return AuthLogOutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthLogOutArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthLogOutResult struct {
	Success *tg.AuthLoggedOut
}

var AuthLogOutResult_Success_DEFAULT *tg.AuthLoggedOut

func (p *AuthLogOutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthLogOutResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthLogOutResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthLoggedOut)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthLogOutResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthLogOutResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthLogOutResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthLoggedOut)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthLogOutResult) GetSuccess() *tg.AuthLoggedOut {
	if !p.IsSetSuccess() {
		return AuthLogOutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthLogOutResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthLoggedOut)
}

func (p *AuthLogOutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthLogOutResult) GetResult() interface{} {
	return p.Success
}

func authResetAuthorizationsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthResetAuthorizationsArgs)
	realResult := result.(*AuthResetAuthorizationsResult)
	success, err := handler.(tg.RPCAuthorization).AuthResetAuthorizations(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthResetAuthorizationsArgs() interface{} {
	return &AuthResetAuthorizationsArgs{}
}

func newAuthResetAuthorizationsResult() interface{} {
	return &AuthResetAuthorizationsResult{}
}

type AuthResetAuthorizationsArgs struct {
	Req *tg.TLAuthResetAuthorizations
}

func (p *AuthResetAuthorizationsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthResetAuthorizationsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthResetAuthorizationsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthResetAuthorizations)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthResetAuthorizationsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthResetAuthorizationsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthResetAuthorizationsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthResetAuthorizations)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthResetAuthorizationsArgs_Req_DEFAULT *tg.TLAuthResetAuthorizations

func (p *AuthResetAuthorizationsArgs) GetReq() *tg.TLAuthResetAuthorizations {
	if !p.IsSetReq() {
		return AuthResetAuthorizationsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthResetAuthorizationsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthResetAuthorizationsResult struct {
	Success *tg.Bool
}

var AuthResetAuthorizationsResult_Success_DEFAULT *tg.Bool

func (p *AuthResetAuthorizationsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthResetAuthorizationsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthResetAuthorizationsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthResetAuthorizationsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthResetAuthorizationsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthResetAuthorizationsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthResetAuthorizationsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AuthResetAuthorizationsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthResetAuthorizationsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AuthResetAuthorizationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthResetAuthorizationsResult) GetResult() interface{} {
	return p.Success
}

func authExportAuthorizationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthExportAuthorizationArgs)
	realResult := result.(*AuthExportAuthorizationResult)
	success, err := handler.(tg.RPCAuthorization).AuthExportAuthorization(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthExportAuthorizationArgs() interface{} {
	return &AuthExportAuthorizationArgs{}
}

func newAuthExportAuthorizationResult() interface{} {
	return &AuthExportAuthorizationResult{}
}

type AuthExportAuthorizationArgs struct {
	Req *tg.TLAuthExportAuthorization
}

func (p *AuthExportAuthorizationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthExportAuthorizationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthExportAuthorizationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthExportAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthExportAuthorizationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthExportAuthorizationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthExportAuthorizationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthExportAuthorization)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthExportAuthorizationArgs_Req_DEFAULT *tg.TLAuthExportAuthorization

func (p *AuthExportAuthorizationArgs) GetReq() *tg.TLAuthExportAuthorization {
	if !p.IsSetReq() {
		return AuthExportAuthorizationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthExportAuthorizationArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthExportAuthorizationResult struct {
	Success *tg.AuthExportedAuthorization
}

var AuthExportAuthorizationResult_Success_DEFAULT *tg.AuthExportedAuthorization

func (p *AuthExportAuthorizationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthExportAuthorizationResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthExportAuthorizationResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthExportedAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthExportAuthorizationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthExportAuthorizationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthExportAuthorizationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthExportedAuthorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthExportAuthorizationResult) GetSuccess() *tg.AuthExportedAuthorization {
	if !p.IsSetSuccess() {
		return AuthExportAuthorizationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthExportAuthorizationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthExportedAuthorization)
}

func (p *AuthExportAuthorizationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthExportAuthorizationResult) GetResult() interface{} {
	return p.Success
}

func authImportAuthorizationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthImportAuthorizationArgs)
	realResult := result.(*AuthImportAuthorizationResult)
	success, err := handler.(tg.RPCAuthorization).AuthImportAuthorization(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthImportAuthorizationArgs() interface{} {
	return &AuthImportAuthorizationArgs{}
}

func newAuthImportAuthorizationResult() interface{} {
	return &AuthImportAuthorizationResult{}
}

type AuthImportAuthorizationArgs struct {
	Req *tg.TLAuthImportAuthorization
}

func (p *AuthImportAuthorizationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthImportAuthorizationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthImportAuthorizationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthImportAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthImportAuthorizationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthImportAuthorizationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthImportAuthorizationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthImportAuthorization)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthImportAuthorizationArgs_Req_DEFAULT *tg.TLAuthImportAuthorization

func (p *AuthImportAuthorizationArgs) GetReq() *tg.TLAuthImportAuthorization {
	if !p.IsSetReq() {
		return AuthImportAuthorizationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthImportAuthorizationArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthImportAuthorizationResult struct {
	Success *tg.AuthAuthorization
}

var AuthImportAuthorizationResult_Success_DEFAULT *tg.AuthAuthorization

func (p *AuthImportAuthorizationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthImportAuthorizationResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthImportAuthorizationResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthImportAuthorizationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthImportAuthorizationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthImportAuthorizationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthAuthorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthImportAuthorizationResult) GetSuccess() *tg.AuthAuthorization {
	if !p.IsSetSuccess() {
		return AuthImportAuthorizationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthImportAuthorizationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthAuthorization)
}

func (p *AuthImportAuthorizationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthImportAuthorizationResult) GetResult() interface{} {
	return p.Success
}

func authBindTempAuthKeyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthBindTempAuthKeyArgs)
	realResult := result.(*AuthBindTempAuthKeyResult)
	success, err := handler.(tg.RPCAuthorization).AuthBindTempAuthKey(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthBindTempAuthKeyArgs() interface{} {
	return &AuthBindTempAuthKeyArgs{}
}

func newAuthBindTempAuthKeyResult() interface{} {
	return &AuthBindTempAuthKeyResult{}
}

type AuthBindTempAuthKeyArgs struct {
	Req *tg.TLAuthBindTempAuthKey
}

func (p *AuthBindTempAuthKeyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthBindTempAuthKeyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthBindTempAuthKeyArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthBindTempAuthKey)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthBindTempAuthKeyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthBindTempAuthKeyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthBindTempAuthKeyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthBindTempAuthKey)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthBindTempAuthKeyArgs_Req_DEFAULT *tg.TLAuthBindTempAuthKey

func (p *AuthBindTempAuthKeyArgs) GetReq() *tg.TLAuthBindTempAuthKey {
	if !p.IsSetReq() {
		return AuthBindTempAuthKeyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthBindTempAuthKeyArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthBindTempAuthKeyResult struct {
	Success *tg.Bool
}

var AuthBindTempAuthKeyResult_Success_DEFAULT *tg.Bool

func (p *AuthBindTempAuthKeyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthBindTempAuthKeyResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthBindTempAuthKeyResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthBindTempAuthKeyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthBindTempAuthKeyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthBindTempAuthKeyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthBindTempAuthKeyResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AuthBindTempAuthKeyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthBindTempAuthKeyResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AuthBindTempAuthKeyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthBindTempAuthKeyResult) GetResult() interface{} {
	return p.Success
}

func authImportBotAuthorizationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthImportBotAuthorizationArgs)
	realResult := result.(*AuthImportBotAuthorizationResult)
	success, err := handler.(tg.RPCAuthorization).AuthImportBotAuthorization(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthImportBotAuthorizationArgs() interface{} {
	return &AuthImportBotAuthorizationArgs{}
}

func newAuthImportBotAuthorizationResult() interface{} {
	return &AuthImportBotAuthorizationResult{}
}

type AuthImportBotAuthorizationArgs struct {
	Req *tg.TLAuthImportBotAuthorization
}

func (p *AuthImportBotAuthorizationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthImportBotAuthorizationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthImportBotAuthorizationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthImportBotAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthImportBotAuthorizationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthImportBotAuthorizationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthImportBotAuthorizationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthImportBotAuthorization)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthImportBotAuthorizationArgs_Req_DEFAULT *tg.TLAuthImportBotAuthorization

func (p *AuthImportBotAuthorizationArgs) GetReq() *tg.TLAuthImportBotAuthorization {
	if !p.IsSetReq() {
		return AuthImportBotAuthorizationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthImportBotAuthorizationArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthImportBotAuthorizationResult struct {
	Success *tg.AuthAuthorization
}

var AuthImportBotAuthorizationResult_Success_DEFAULT *tg.AuthAuthorization

func (p *AuthImportBotAuthorizationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthImportBotAuthorizationResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthImportBotAuthorizationResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthImportBotAuthorizationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthImportBotAuthorizationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthImportBotAuthorizationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthAuthorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthImportBotAuthorizationResult) GetSuccess() *tg.AuthAuthorization {
	if !p.IsSetSuccess() {
		return AuthImportBotAuthorizationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthImportBotAuthorizationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthAuthorization)
}

func (p *AuthImportBotAuthorizationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthImportBotAuthorizationResult) GetResult() interface{} {
	return p.Success
}

func authCheckPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthCheckPasswordArgs)
	realResult := result.(*AuthCheckPasswordResult)
	success, err := handler.(tg.RPCAuthorization).AuthCheckPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthCheckPasswordArgs() interface{} {
	return &AuthCheckPasswordArgs{}
}

func newAuthCheckPasswordResult() interface{} {
	return &AuthCheckPasswordResult{}
}

type AuthCheckPasswordArgs struct {
	Req *tg.TLAuthCheckPassword
}

func (p *AuthCheckPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthCheckPasswordArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthCheckPasswordArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthCheckPassword)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthCheckPasswordArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthCheckPasswordArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthCheckPasswordArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthCheckPassword)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthCheckPasswordArgs_Req_DEFAULT *tg.TLAuthCheckPassword

func (p *AuthCheckPasswordArgs) GetReq() *tg.TLAuthCheckPassword {
	if !p.IsSetReq() {
		return AuthCheckPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthCheckPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthCheckPasswordResult struct {
	Success *tg.AuthAuthorization
}

var AuthCheckPasswordResult_Success_DEFAULT *tg.AuthAuthorization

func (p *AuthCheckPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthCheckPasswordResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthCheckPasswordResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthCheckPasswordResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthCheckPasswordResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthCheckPasswordResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthAuthorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthCheckPasswordResult) GetSuccess() *tg.AuthAuthorization {
	if !p.IsSetSuccess() {
		return AuthCheckPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthCheckPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthAuthorization)
}

func (p *AuthCheckPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthCheckPasswordResult) GetResult() interface{} {
	return p.Success
}

func authRequestPasswordRecoveryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthRequestPasswordRecoveryArgs)
	realResult := result.(*AuthRequestPasswordRecoveryResult)
	success, err := handler.(tg.RPCAuthorization).AuthRequestPasswordRecovery(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthRequestPasswordRecoveryArgs() interface{} {
	return &AuthRequestPasswordRecoveryArgs{}
}

func newAuthRequestPasswordRecoveryResult() interface{} {
	return &AuthRequestPasswordRecoveryResult{}
}

type AuthRequestPasswordRecoveryArgs struct {
	Req *tg.TLAuthRequestPasswordRecovery
}

func (p *AuthRequestPasswordRecoveryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthRequestPasswordRecoveryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthRequestPasswordRecoveryArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthRequestPasswordRecovery)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthRequestPasswordRecoveryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthRequestPasswordRecoveryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthRequestPasswordRecoveryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthRequestPasswordRecovery)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthRequestPasswordRecoveryArgs_Req_DEFAULT *tg.TLAuthRequestPasswordRecovery

func (p *AuthRequestPasswordRecoveryArgs) GetReq() *tg.TLAuthRequestPasswordRecovery {
	if !p.IsSetReq() {
		return AuthRequestPasswordRecoveryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthRequestPasswordRecoveryArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthRequestPasswordRecoveryResult struct {
	Success *tg.AuthPasswordRecovery
}

var AuthRequestPasswordRecoveryResult_Success_DEFAULT *tg.AuthPasswordRecovery

func (p *AuthRequestPasswordRecoveryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthRequestPasswordRecoveryResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthRequestPasswordRecoveryResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthPasswordRecovery)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthRequestPasswordRecoveryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthRequestPasswordRecoveryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthRequestPasswordRecoveryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthPasswordRecovery)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthRequestPasswordRecoveryResult) GetSuccess() *tg.AuthPasswordRecovery {
	if !p.IsSetSuccess() {
		return AuthRequestPasswordRecoveryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthRequestPasswordRecoveryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthPasswordRecovery)
}

func (p *AuthRequestPasswordRecoveryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthRequestPasswordRecoveryResult) GetResult() interface{} {
	return p.Success
}

func authRecoverPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthRecoverPasswordArgs)
	realResult := result.(*AuthRecoverPasswordResult)
	success, err := handler.(tg.RPCAuthorization).AuthRecoverPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthRecoverPasswordArgs() interface{} {
	return &AuthRecoverPasswordArgs{}
}

func newAuthRecoverPasswordResult() interface{} {
	return &AuthRecoverPasswordResult{}
}

type AuthRecoverPasswordArgs struct {
	Req *tg.TLAuthRecoverPassword
}

func (p *AuthRecoverPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthRecoverPasswordArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthRecoverPasswordArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthRecoverPassword)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthRecoverPasswordArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthRecoverPasswordArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthRecoverPasswordArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthRecoverPassword)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthRecoverPasswordArgs_Req_DEFAULT *tg.TLAuthRecoverPassword

func (p *AuthRecoverPasswordArgs) GetReq() *tg.TLAuthRecoverPassword {
	if !p.IsSetReq() {
		return AuthRecoverPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthRecoverPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthRecoverPasswordResult struct {
	Success *tg.AuthAuthorization
}

var AuthRecoverPasswordResult_Success_DEFAULT *tg.AuthAuthorization

func (p *AuthRecoverPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthRecoverPasswordResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthRecoverPasswordResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthRecoverPasswordResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthRecoverPasswordResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthRecoverPasswordResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthAuthorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthRecoverPasswordResult) GetSuccess() *tg.AuthAuthorization {
	if !p.IsSetSuccess() {
		return AuthRecoverPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthRecoverPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthAuthorization)
}

func (p *AuthRecoverPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthRecoverPasswordResult) GetResult() interface{} {
	return p.Success
}

func authResendCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthResendCodeArgs)
	realResult := result.(*AuthResendCodeResult)
	success, err := handler.(tg.RPCAuthorization).AuthResendCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthResendCodeArgs() interface{} {
	return &AuthResendCodeArgs{}
}

func newAuthResendCodeResult() interface{} {
	return &AuthResendCodeResult{}
}

type AuthResendCodeArgs struct {
	Req *tg.TLAuthResendCode
}

func (p *AuthResendCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthResendCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthResendCodeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthResendCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthResendCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthResendCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthResendCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthResendCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthResendCodeArgs_Req_DEFAULT *tg.TLAuthResendCode

func (p *AuthResendCodeArgs) GetReq() *tg.TLAuthResendCode {
	if !p.IsSetReq() {
		return AuthResendCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthResendCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthResendCodeResult struct {
	Success *tg.AuthSentCode
}

var AuthResendCodeResult_Success_DEFAULT *tg.AuthSentCode

func (p *AuthResendCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthResendCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthResendCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthSentCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthResendCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthResendCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthResendCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthSentCode)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthResendCodeResult) GetSuccess() *tg.AuthSentCode {
	if !p.IsSetSuccess() {
		return AuthResendCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthResendCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthSentCode)
}

func (p *AuthResendCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthResendCodeResult) GetResult() interface{} {
	return p.Success
}

func authCancelCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthCancelCodeArgs)
	realResult := result.(*AuthCancelCodeResult)
	success, err := handler.(tg.RPCAuthorization).AuthCancelCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthCancelCodeArgs() interface{} {
	return &AuthCancelCodeArgs{}
}

func newAuthCancelCodeResult() interface{} {
	return &AuthCancelCodeResult{}
}

type AuthCancelCodeArgs struct {
	Req *tg.TLAuthCancelCode
}

func (p *AuthCancelCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthCancelCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthCancelCodeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthCancelCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthCancelCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthCancelCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthCancelCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthCancelCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthCancelCodeArgs_Req_DEFAULT *tg.TLAuthCancelCode

func (p *AuthCancelCodeArgs) GetReq() *tg.TLAuthCancelCode {
	if !p.IsSetReq() {
		return AuthCancelCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthCancelCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthCancelCodeResult struct {
	Success *tg.Bool
}

var AuthCancelCodeResult_Success_DEFAULT *tg.Bool

func (p *AuthCancelCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthCancelCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthCancelCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthCancelCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthCancelCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthCancelCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthCancelCodeResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AuthCancelCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthCancelCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AuthCancelCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthCancelCodeResult) GetResult() interface{} {
	return p.Success
}

func authDropTempAuthKeysHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthDropTempAuthKeysArgs)
	realResult := result.(*AuthDropTempAuthKeysResult)
	success, err := handler.(tg.RPCAuthorization).AuthDropTempAuthKeys(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthDropTempAuthKeysArgs() interface{} {
	return &AuthDropTempAuthKeysArgs{}
}

func newAuthDropTempAuthKeysResult() interface{} {
	return &AuthDropTempAuthKeysResult{}
}

type AuthDropTempAuthKeysArgs struct {
	Req *tg.TLAuthDropTempAuthKeys
}

func (p *AuthDropTempAuthKeysArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthDropTempAuthKeysArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthDropTempAuthKeysArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthDropTempAuthKeys)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthDropTempAuthKeysArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthDropTempAuthKeysArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthDropTempAuthKeysArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthDropTempAuthKeys)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthDropTempAuthKeysArgs_Req_DEFAULT *tg.TLAuthDropTempAuthKeys

func (p *AuthDropTempAuthKeysArgs) GetReq() *tg.TLAuthDropTempAuthKeys {
	if !p.IsSetReq() {
		return AuthDropTempAuthKeysArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthDropTempAuthKeysArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthDropTempAuthKeysResult struct {
	Success *tg.Bool
}

var AuthDropTempAuthKeysResult_Success_DEFAULT *tg.Bool

func (p *AuthDropTempAuthKeysResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthDropTempAuthKeysResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthDropTempAuthKeysResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthDropTempAuthKeysResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthDropTempAuthKeysResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthDropTempAuthKeysResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthDropTempAuthKeysResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AuthDropTempAuthKeysResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthDropTempAuthKeysResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AuthDropTempAuthKeysResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthDropTempAuthKeysResult) GetResult() interface{} {
	return p.Success
}

func authCheckRecoveryPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthCheckRecoveryPasswordArgs)
	realResult := result.(*AuthCheckRecoveryPasswordResult)
	success, err := handler.(tg.RPCAuthorization).AuthCheckRecoveryPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthCheckRecoveryPasswordArgs() interface{} {
	return &AuthCheckRecoveryPasswordArgs{}
}

func newAuthCheckRecoveryPasswordResult() interface{} {
	return &AuthCheckRecoveryPasswordResult{}
}

type AuthCheckRecoveryPasswordArgs struct {
	Req *tg.TLAuthCheckRecoveryPassword
}

func (p *AuthCheckRecoveryPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthCheckRecoveryPasswordArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthCheckRecoveryPasswordArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthCheckRecoveryPassword)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthCheckRecoveryPasswordArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthCheckRecoveryPasswordArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthCheckRecoveryPasswordArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthCheckRecoveryPassword)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthCheckRecoveryPasswordArgs_Req_DEFAULT *tg.TLAuthCheckRecoveryPassword

func (p *AuthCheckRecoveryPasswordArgs) GetReq() *tg.TLAuthCheckRecoveryPassword {
	if !p.IsSetReq() {
		return AuthCheckRecoveryPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthCheckRecoveryPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthCheckRecoveryPasswordResult struct {
	Success *tg.Bool
}

var AuthCheckRecoveryPasswordResult_Success_DEFAULT *tg.Bool

func (p *AuthCheckRecoveryPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthCheckRecoveryPasswordResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthCheckRecoveryPasswordResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthCheckRecoveryPasswordResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthCheckRecoveryPasswordResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthCheckRecoveryPasswordResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthCheckRecoveryPasswordResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AuthCheckRecoveryPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthCheckRecoveryPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AuthCheckRecoveryPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthCheckRecoveryPasswordResult) GetResult() interface{} {
	return p.Success
}

func authImportWebTokenAuthorizationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthImportWebTokenAuthorizationArgs)
	realResult := result.(*AuthImportWebTokenAuthorizationResult)
	success, err := handler.(tg.RPCAuthorization).AuthImportWebTokenAuthorization(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthImportWebTokenAuthorizationArgs() interface{} {
	return &AuthImportWebTokenAuthorizationArgs{}
}

func newAuthImportWebTokenAuthorizationResult() interface{} {
	return &AuthImportWebTokenAuthorizationResult{}
}

type AuthImportWebTokenAuthorizationArgs struct {
	Req *tg.TLAuthImportWebTokenAuthorization
}

func (p *AuthImportWebTokenAuthorizationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthImportWebTokenAuthorizationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthImportWebTokenAuthorizationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthImportWebTokenAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthImportWebTokenAuthorizationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthImportWebTokenAuthorizationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthImportWebTokenAuthorizationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthImportWebTokenAuthorization)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthImportWebTokenAuthorizationArgs_Req_DEFAULT *tg.TLAuthImportWebTokenAuthorization

func (p *AuthImportWebTokenAuthorizationArgs) GetReq() *tg.TLAuthImportWebTokenAuthorization {
	if !p.IsSetReq() {
		return AuthImportWebTokenAuthorizationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthImportWebTokenAuthorizationArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthImportWebTokenAuthorizationResult struct {
	Success *tg.AuthAuthorization
}

var AuthImportWebTokenAuthorizationResult_Success_DEFAULT *tg.AuthAuthorization

func (p *AuthImportWebTokenAuthorizationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthImportWebTokenAuthorizationResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthImportWebTokenAuthorizationResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthAuthorization)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthImportWebTokenAuthorizationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthImportWebTokenAuthorizationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthImportWebTokenAuthorizationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthAuthorization)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthImportWebTokenAuthorizationResult) GetSuccess() *tg.AuthAuthorization {
	if !p.IsSetSuccess() {
		return AuthImportWebTokenAuthorizationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthImportWebTokenAuthorizationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthAuthorization)
}

func (p *AuthImportWebTokenAuthorizationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthImportWebTokenAuthorizationResult) GetResult() interface{} {
	return p.Success
}

func authRequestFirebaseSmsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthRequestFirebaseSmsArgs)
	realResult := result.(*AuthRequestFirebaseSmsResult)
	success, err := handler.(tg.RPCAuthorization).AuthRequestFirebaseSms(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthRequestFirebaseSmsArgs() interface{} {
	return &AuthRequestFirebaseSmsArgs{}
}

func newAuthRequestFirebaseSmsResult() interface{} {
	return &AuthRequestFirebaseSmsResult{}
}

type AuthRequestFirebaseSmsArgs struct {
	Req *tg.TLAuthRequestFirebaseSms
}

func (p *AuthRequestFirebaseSmsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthRequestFirebaseSmsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthRequestFirebaseSmsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthRequestFirebaseSms)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthRequestFirebaseSmsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthRequestFirebaseSmsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthRequestFirebaseSmsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthRequestFirebaseSms)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthRequestFirebaseSmsArgs_Req_DEFAULT *tg.TLAuthRequestFirebaseSms

func (p *AuthRequestFirebaseSmsArgs) GetReq() *tg.TLAuthRequestFirebaseSms {
	if !p.IsSetReq() {
		return AuthRequestFirebaseSmsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthRequestFirebaseSmsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthRequestFirebaseSmsResult struct {
	Success *tg.Bool
}

var AuthRequestFirebaseSmsResult_Success_DEFAULT *tg.Bool

func (p *AuthRequestFirebaseSmsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthRequestFirebaseSmsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthRequestFirebaseSmsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthRequestFirebaseSmsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthRequestFirebaseSmsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthRequestFirebaseSmsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthRequestFirebaseSmsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AuthRequestFirebaseSmsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthRequestFirebaseSmsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AuthRequestFirebaseSmsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthRequestFirebaseSmsResult) GetResult() interface{} {
	return p.Success
}

func authResetLoginEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthResetLoginEmailArgs)
	realResult := result.(*AuthResetLoginEmailResult)
	success, err := handler.(tg.RPCAuthorization).AuthResetLoginEmail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthResetLoginEmailArgs() interface{} {
	return &AuthResetLoginEmailArgs{}
}

func newAuthResetLoginEmailResult() interface{} {
	return &AuthResetLoginEmailResult{}
}

type AuthResetLoginEmailArgs struct {
	Req *tg.TLAuthResetLoginEmail
}

func (p *AuthResetLoginEmailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthResetLoginEmailArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthResetLoginEmailArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthResetLoginEmail)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthResetLoginEmailArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthResetLoginEmailArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthResetLoginEmailArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthResetLoginEmail)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthResetLoginEmailArgs_Req_DEFAULT *tg.TLAuthResetLoginEmail

func (p *AuthResetLoginEmailArgs) GetReq() *tg.TLAuthResetLoginEmail {
	if !p.IsSetReq() {
		return AuthResetLoginEmailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthResetLoginEmailArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthResetLoginEmailResult struct {
	Success *tg.AuthSentCode
}

var AuthResetLoginEmailResult_Success_DEFAULT *tg.AuthSentCode

func (p *AuthResetLoginEmailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthResetLoginEmailResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthResetLoginEmailResult) Unmarshal(in []byte) error {
	msg := new(tg.AuthSentCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthResetLoginEmailResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthResetLoginEmailResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthResetLoginEmailResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AuthSentCode)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthResetLoginEmailResult) GetSuccess() *tg.AuthSentCode {
	if !p.IsSetSuccess() {
		return AuthResetLoginEmailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthResetLoginEmailResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AuthSentCode)
}

func (p *AuthResetLoginEmailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthResetLoginEmailResult) GetResult() interface{} {
	return p.Success
}

func authReportMissingCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthReportMissingCodeArgs)
	realResult := result.(*AuthReportMissingCodeResult)
	success, err := handler.(tg.RPCAuthorization).AuthReportMissingCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthReportMissingCodeArgs() interface{} {
	return &AuthReportMissingCodeArgs{}
}

func newAuthReportMissingCodeResult() interface{} {
	return &AuthReportMissingCodeResult{}
}

type AuthReportMissingCodeArgs struct {
	Req *tg.TLAuthReportMissingCode
}

func (p *AuthReportMissingCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthReportMissingCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthReportMissingCodeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthReportMissingCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthReportMissingCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthReportMissingCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthReportMissingCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthReportMissingCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthReportMissingCodeArgs_Req_DEFAULT *tg.TLAuthReportMissingCode

func (p *AuthReportMissingCodeArgs) GetReq() *tg.TLAuthReportMissingCode {
	if !p.IsSetReq() {
		return AuthReportMissingCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthReportMissingCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthReportMissingCodeResult struct {
	Success *tg.Bool
}

var AuthReportMissingCodeResult_Success_DEFAULT *tg.Bool

func (p *AuthReportMissingCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthReportMissingCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthReportMissingCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthReportMissingCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthReportMissingCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthReportMissingCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthReportMissingCodeResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AuthReportMissingCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthReportMissingCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AuthReportMissingCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthReportMissingCodeResult) GetResult() interface{} {
	return p.Success
}

func accountSendVerifyEmailCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSendVerifyEmailCodeArgs)
	realResult := result.(*AccountSendVerifyEmailCodeResult)
	success, err := handler.(tg.RPCAuthorization).AccountSendVerifyEmailCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSendVerifyEmailCodeArgs() interface{} {
	return &AccountSendVerifyEmailCodeArgs{}
}

func newAccountSendVerifyEmailCodeResult() interface{} {
	return &AccountSendVerifyEmailCodeResult{}
}

type AccountSendVerifyEmailCodeArgs struct {
	Req *tg.TLAccountSendVerifyEmailCode
}

func (p *AccountSendVerifyEmailCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSendVerifyEmailCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSendVerifyEmailCodeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSendVerifyEmailCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSendVerifyEmailCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSendVerifyEmailCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSendVerifyEmailCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSendVerifyEmailCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSendVerifyEmailCodeArgs_Req_DEFAULT *tg.TLAccountSendVerifyEmailCode

func (p *AccountSendVerifyEmailCodeArgs) GetReq() *tg.TLAccountSendVerifyEmailCode {
	if !p.IsSetReq() {
		return AccountSendVerifyEmailCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSendVerifyEmailCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSendVerifyEmailCodeResult struct {
	Success *tg.AccountSentEmailCode
}

var AccountSendVerifyEmailCodeResult_Success_DEFAULT *tg.AccountSentEmailCode

func (p *AccountSendVerifyEmailCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSendVerifyEmailCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSendVerifyEmailCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountSentEmailCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSendVerifyEmailCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSendVerifyEmailCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSendVerifyEmailCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountSentEmailCode)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSendVerifyEmailCodeResult) GetSuccess() *tg.AccountSentEmailCode {
	if !p.IsSetSuccess() {
		return AccountSendVerifyEmailCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSendVerifyEmailCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountSentEmailCode)
}

func (p *AccountSendVerifyEmailCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSendVerifyEmailCodeResult) GetResult() interface{} {
	return p.Success
}

func accountVerifyEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountVerifyEmailArgs)
	realResult := result.(*AccountVerifyEmailResult)
	success, err := handler.(tg.RPCAuthorization).AccountVerifyEmail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountVerifyEmailArgs() interface{} {
	return &AccountVerifyEmailArgs{}
}

func newAccountVerifyEmailResult() interface{} {
	return &AccountVerifyEmailResult{}
}

type AccountVerifyEmailArgs struct {
	Req *tg.TLAccountVerifyEmail
}

func (p *AccountVerifyEmailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountVerifyEmailArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountVerifyEmailArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountVerifyEmail)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountVerifyEmailArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountVerifyEmailArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountVerifyEmailArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountVerifyEmail)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountVerifyEmailArgs_Req_DEFAULT *tg.TLAccountVerifyEmail

func (p *AccountVerifyEmailArgs) GetReq() *tg.TLAccountVerifyEmail {
	if !p.IsSetReq() {
		return AccountVerifyEmailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountVerifyEmailArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountVerifyEmailResult struct {
	Success *tg.AccountEmailVerified
}

var AccountVerifyEmailResult_Success_DEFAULT *tg.AccountEmailVerified

func (p *AccountVerifyEmailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountVerifyEmailResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountVerifyEmailResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountEmailVerified)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountVerifyEmailResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountVerifyEmailResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountVerifyEmailResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountEmailVerified)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountVerifyEmailResult) GetSuccess() *tg.AccountEmailVerified {
	if !p.IsSetSuccess() {
		return AccountVerifyEmailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountVerifyEmailResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountEmailVerified)
}

func (p *AccountVerifyEmailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountVerifyEmailResult) GetResult() interface{} {
	return p.Success
}

func accountResetPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountResetPasswordArgs)
	realResult := result.(*AccountResetPasswordResult)
	success, err := handler.(tg.RPCAuthorization).AccountResetPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountResetPasswordArgs() interface{} {
	return &AccountResetPasswordArgs{}
}

func newAccountResetPasswordResult() interface{} {
	return &AccountResetPasswordResult{}
}

type AccountResetPasswordArgs struct {
	Req *tg.TLAccountResetPassword
}

func (p *AccountResetPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountResetPasswordArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountResetPasswordArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountResetPassword)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountResetPasswordArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountResetPasswordArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountResetPasswordArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountResetPassword)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountResetPasswordArgs_Req_DEFAULT *tg.TLAccountResetPassword

func (p *AccountResetPasswordArgs) GetReq() *tg.TLAccountResetPassword {
	if !p.IsSetReq() {
		return AccountResetPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountResetPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountResetPasswordResult struct {
	Success *tg.AccountResetPasswordResult
}

var AccountResetPasswordResult_Success_DEFAULT *tg.AccountResetPasswordResult

func (p *AccountResetPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountResetPasswordResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountResetPasswordResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountResetPasswordResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountResetPasswordResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountResetPasswordResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountResetPasswordResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountResetPasswordResult)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountResetPasswordResult) GetSuccess() *tg.AccountResetPasswordResult {
	if !p.IsSetSuccess() {
		return AccountResetPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountResetPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountResetPasswordResult)
}

func (p *AccountResetPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountResetPasswordResult) GetResult() interface{} {
	return p.Success
}

func accountSetAuthorizationTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSetAuthorizationTTLArgs)
	realResult := result.(*AccountSetAuthorizationTTLResult)
	success, err := handler.(tg.RPCAuthorization).AccountSetAuthorizationTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSetAuthorizationTTLArgs() interface{} {
	return &AccountSetAuthorizationTTLArgs{}
}

func newAccountSetAuthorizationTTLResult() interface{} {
	return &AccountSetAuthorizationTTLResult{}
}

type AccountSetAuthorizationTTLArgs struct {
	Req *tg.TLAccountSetAuthorizationTTL
}

func (p *AccountSetAuthorizationTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSetAuthorizationTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSetAuthorizationTTLArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSetAuthorizationTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSetAuthorizationTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSetAuthorizationTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSetAuthorizationTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSetAuthorizationTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSetAuthorizationTTLArgs_Req_DEFAULT *tg.TLAccountSetAuthorizationTTL

func (p *AccountSetAuthorizationTTLArgs) GetReq() *tg.TLAccountSetAuthorizationTTL {
	if !p.IsSetReq() {
		return AccountSetAuthorizationTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSetAuthorizationTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSetAuthorizationTTLResult struct {
	Success *tg.Bool
}

var AccountSetAuthorizationTTLResult_Success_DEFAULT *tg.Bool

func (p *AccountSetAuthorizationTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSetAuthorizationTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSetAuthorizationTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetAuthorizationTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSetAuthorizationTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSetAuthorizationTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetAuthorizationTTLResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountSetAuthorizationTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSetAuthorizationTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountSetAuthorizationTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSetAuthorizationTTLResult) GetResult() interface{} {
	return p.Success
}

func accountChangeAuthorizationSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountChangeAuthorizationSettingsArgs)
	realResult := result.(*AccountChangeAuthorizationSettingsResult)
	success, err := handler.(tg.RPCAuthorization).AccountChangeAuthorizationSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountChangeAuthorizationSettingsArgs() interface{} {
	return &AccountChangeAuthorizationSettingsArgs{}
}

func newAccountChangeAuthorizationSettingsResult() interface{} {
	return &AccountChangeAuthorizationSettingsResult{}
}

type AccountChangeAuthorizationSettingsArgs struct {
	Req *tg.TLAccountChangeAuthorizationSettings
}

func (p *AccountChangeAuthorizationSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountChangeAuthorizationSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountChangeAuthorizationSettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountChangeAuthorizationSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountChangeAuthorizationSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountChangeAuthorizationSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountChangeAuthorizationSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountChangeAuthorizationSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountChangeAuthorizationSettingsArgs_Req_DEFAULT *tg.TLAccountChangeAuthorizationSettings

func (p *AccountChangeAuthorizationSettingsArgs) GetReq() *tg.TLAccountChangeAuthorizationSettings {
	if !p.IsSetReq() {
		return AccountChangeAuthorizationSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountChangeAuthorizationSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountChangeAuthorizationSettingsResult struct {
	Success *tg.Bool
}

var AccountChangeAuthorizationSettingsResult_Success_DEFAULT *tg.Bool

func (p *AccountChangeAuthorizationSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountChangeAuthorizationSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountChangeAuthorizationSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountChangeAuthorizationSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountChangeAuthorizationSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountChangeAuthorizationSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountChangeAuthorizationSettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountChangeAuthorizationSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountChangeAuthorizationSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountChangeAuthorizationSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountChangeAuthorizationSettingsResult) GetResult() interface{} {
	return p.Success
}

func accountInvalidateSignInCodesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountInvalidateSignInCodesArgs)
	realResult := result.(*AccountInvalidateSignInCodesResult)
	success, err := handler.(tg.RPCAuthorization).AccountInvalidateSignInCodes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountInvalidateSignInCodesArgs() interface{} {
	return &AccountInvalidateSignInCodesArgs{}
}

func newAccountInvalidateSignInCodesResult() interface{} {
	return &AccountInvalidateSignInCodesResult{}
}

type AccountInvalidateSignInCodesArgs struct {
	Req *tg.TLAccountInvalidateSignInCodes
}

func (p *AccountInvalidateSignInCodesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountInvalidateSignInCodesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountInvalidateSignInCodesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountInvalidateSignInCodes)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountInvalidateSignInCodesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountInvalidateSignInCodesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountInvalidateSignInCodesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountInvalidateSignInCodes)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountInvalidateSignInCodesArgs_Req_DEFAULT *tg.TLAccountInvalidateSignInCodes

func (p *AccountInvalidateSignInCodesArgs) GetReq() *tg.TLAccountInvalidateSignInCodes {
	if !p.IsSetReq() {
		return AccountInvalidateSignInCodesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountInvalidateSignInCodesArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountInvalidateSignInCodesResult struct {
	Success *tg.Bool
}

var AccountInvalidateSignInCodesResult_Success_DEFAULT *tg.Bool

func (p *AccountInvalidateSignInCodesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountInvalidateSignInCodesResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountInvalidateSignInCodesResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountInvalidateSignInCodesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountInvalidateSignInCodesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountInvalidateSignInCodesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountInvalidateSignInCodesResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountInvalidateSignInCodesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountInvalidateSignInCodesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountInvalidateSignInCodesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountInvalidateSignInCodesResult) GetResult() interface{} {
	return p.Success
}

func authToggleBanHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AuthToggleBanArgs)
	realResult := result.(*AuthToggleBanResult)
	success, err := handler.(tg.RPCAuthorization).AuthToggleBan(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAuthToggleBanArgs() interface{} {
	return &AuthToggleBanArgs{}
}

func newAuthToggleBanResult() interface{} {
	return &AuthToggleBanResult{}
}

type AuthToggleBanArgs struct {
	Req *tg.TLAuthToggleBan
}

func (p *AuthToggleBanArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthToggleBanArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AuthToggleBanArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAuthToggleBan)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AuthToggleBanArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AuthToggleBanArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AuthToggleBanArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAuthToggleBan)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AuthToggleBanArgs_Req_DEFAULT *tg.TLAuthToggleBan

func (p *AuthToggleBanArgs) GetReq() *tg.TLAuthToggleBan {
	if !p.IsSetReq() {
		return AuthToggleBanArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthToggleBanArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthToggleBanResult struct {
	Success *tg.PredefinedUser
}

var AuthToggleBanResult_Success_DEFAULT *tg.PredefinedUser

func (p *AuthToggleBanResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthToggleBanResult")
	}
	return json.Marshal(p.Success)
}

func (p *AuthToggleBanResult) Unmarshal(in []byte) error {
	msg := new(tg.PredefinedUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthToggleBanResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AuthToggleBanResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AuthToggleBanResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.PredefinedUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthToggleBanResult) GetSuccess() *tg.PredefinedUser {
	if !p.IsSetSuccess() {
		return AuthToggleBanResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthToggleBanResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.PredefinedUser)
}

func (p *AuthToggleBanResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthToggleBanResult) GetResult() interface{} {
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

func (p *kClient) AuthSendCode(ctx context.Context, req *tg.TLAuthSendCode) (r *tg.AuthSentCode, err error) {
	var _args AuthSendCodeArgs
	_args.Req = req
	var _result AuthSendCodeResult
	if err = p.c.Call(ctx, "auth.sendCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthSignUp(ctx context.Context, req *tg.TLAuthSignUp) (r *tg.AuthAuthorization, err error) {
	var _args AuthSignUpArgs
	_args.Req = req
	var _result AuthSignUpResult
	if err = p.c.Call(ctx, "auth.signUp", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthSignIn(ctx context.Context, req *tg.TLAuthSignIn) (r *tg.AuthAuthorization, err error) {
	var _args AuthSignInArgs
	_args.Req = req
	var _result AuthSignInResult
	if err = p.c.Call(ctx, "auth.signIn", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthLogOut(ctx context.Context, req *tg.TLAuthLogOut) (r *tg.AuthLoggedOut, err error) {
	var _args AuthLogOutArgs
	_args.Req = req
	var _result AuthLogOutResult
	if err = p.c.Call(ctx, "auth.logOut", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthResetAuthorizations(ctx context.Context, req *tg.TLAuthResetAuthorizations) (r *tg.Bool, err error) {
	var _args AuthResetAuthorizationsArgs
	_args.Req = req
	var _result AuthResetAuthorizationsResult
	if err = p.c.Call(ctx, "auth.resetAuthorizations", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthExportAuthorization(ctx context.Context, req *tg.TLAuthExportAuthorization) (r *tg.AuthExportedAuthorization, err error) {
	var _args AuthExportAuthorizationArgs
	_args.Req = req
	var _result AuthExportAuthorizationResult
	if err = p.c.Call(ctx, "auth.exportAuthorization", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthImportAuthorization(ctx context.Context, req *tg.TLAuthImportAuthorization) (r *tg.AuthAuthorization, err error) {
	var _args AuthImportAuthorizationArgs
	_args.Req = req
	var _result AuthImportAuthorizationResult
	if err = p.c.Call(ctx, "auth.importAuthorization", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthBindTempAuthKey(ctx context.Context, req *tg.TLAuthBindTempAuthKey) (r *tg.Bool, err error) {
	var _args AuthBindTempAuthKeyArgs
	_args.Req = req
	var _result AuthBindTempAuthKeyResult
	if err = p.c.Call(ctx, "auth.bindTempAuthKey", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthImportBotAuthorization(ctx context.Context, req *tg.TLAuthImportBotAuthorization) (r *tg.AuthAuthorization, err error) {
	var _args AuthImportBotAuthorizationArgs
	_args.Req = req
	var _result AuthImportBotAuthorizationResult
	if err = p.c.Call(ctx, "auth.importBotAuthorization", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthCheckPassword(ctx context.Context, req *tg.TLAuthCheckPassword) (r *tg.AuthAuthorization, err error) {
	var _args AuthCheckPasswordArgs
	_args.Req = req
	var _result AuthCheckPasswordResult
	if err = p.c.Call(ctx, "auth.checkPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthRequestPasswordRecovery(ctx context.Context, req *tg.TLAuthRequestPasswordRecovery) (r *tg.AuthPasswordRecovery, err error) {
	var _args AuthRequestPasswordRecoveryArgs
	_args.Req = req
	var _result AuthRequestPasswordRecoveryResult
	if err = p.c.Call(ctx, "auth.requestPasswordRecovery", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthRecoverPassword(ctx context.Context, req *tg.TLAuthRecoverPassword) (r *tg.AuthAuthorization, err error) {
	var _args AuthRecoverPasswordArgs
	_args.Req = req
	var _result AuthRecoverPasswordResult
	if err = p.c.Call(ctx, "auth.recoverPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthResendCode(ctx context.Context, req *tg.TLAuthResendCode) (r *tg.AuthSentCode, err error) {
	var _args AuthResendCodeArgs
	_args.Req = req
	var _result AuthResendCodeResult
	if err = p.c.Call(ctx, "auth.resendCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthCancelCode(ctx context.Context, req *tg.TLAuthCancelCode) (r *tg.Bool, err error) {
	var _args AuthCancelCodeArgs
	_args.Req = req
	var _result AuthCancelCodeResult
	if err = p.c.Call(ctx, "auth.cancelCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthDropTempAuthKeys(ctx context.Context, req *tg.TLAuthDropTempAuthKeys) (r *tg.Bool, err error) {
	var _args AuthDropTempAuthKeysArgs
	_args.Req = req
	var _result AuthDropTempAuthKeysResult
	if err = p.c.Call(ctx, "auth.dropTempAuthKeys", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthCheckRecoveryPassword(ctx context.Context, req *tg.TLAuthCheckRecoveryPassword) (r *tg.Bool, err error) {
	var _args AuthCheckRecoveryPasswordArgs
	_args.Req = req
	var _result AuthCheckRecoveryPasswordResult
	if err = p.c.Call(ctx, "auth.checkRecoveryPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthImportWebTokenAuthorization(ctx context.Context, req *tg.TLAuthImportWebTokenAuthorization) (r *tg.AuthAuthorization, err error) {
	var _args AuthImportWebTokenAuthorizationArgs
	_args.Req = req
	var _result AuthImportWebTokenAuthorizationResult
	if err = p.c.Call(ctx, "auth.importWebTokenAuthorization", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthRequestFirebaseSms(ctx context.Context, req *tg.TLAuthRequestFirebaseSms) (r *tg.Bool, err error) {
	var _args AuthRequestFirebaseSmsArgs
	_args.Req = req
	var _result AuthRequestFirebaseSmsResult
	if err = p.c.Call(ctx, "auth.requestFirebaseSms", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthResetLoginEmail(ctx context.Context, req *tg.TLAuthResetLoginEmail) (r *tg.AuthSentCode, err error) {
	var _args AuthResetLoginEmailArgs
	_args.Req = req
	var _result AuthResetLoginEmailResult
	if err = p.c.Call(ctx, "auth.resetLoginEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthReportMissingCode(ctx context.Context, req *tg.TLAuthReportMissingCode) (r *tg.Bool, err error) {
	var _args AuthReportMissingCodeArgs
	_args.Req = req
	var _result AuthReportMissingCodeResult
	if err = p.c.Call(ctx, "auth.reportMissingCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountSendVerifyEmailCode(ctx context.Context, req *tg.TLAccountSendVerifyEmailCode) (r *tg.AccountSentEmailCode, err error) {
	var _args AccountSendVerifyEmailCodeArgs
	_args.Req = req
	var _result AccountSendVerifyEmailCodeResult
	if err = p.c.Call(ctx, "account.sendVerifyEmailCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountVerifyEmail(ctx context.Context, req *tg.TLAccountVerifyEmail) (r *tg.AccountEmailVerified, err error) {
	var _args AccountVerifyEmailArgs
	_args.Req = req
	var _result AccountVerifyEmailResult
	if err = p.c.Call(ctx, "account.verifyEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountResetPassword(ctx context.Context, req *tg.TLAccountResetPassword) (r *tg.AccountResetPasswordResult, err error) {
	var _args AccountResetPasswordArgs
	_args.Req = req
	var _result AccountResetPasswordResult
	if err = p.c.Call(ctx, "account.resetPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountSetAuthorizationTTL(ctx context.Context, req *tg.TLAccountSetAuthorizationTTL) (r *tg.Bool, err error) {
	var _args AccountSetAuthorizationTTLArgs
	_args.Req = req
	var _result AccountSetAuthorizationTTLResult
	if err = p.c.Call(ctx, "account.setAuthorizationTTL", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountChangeAuthorizationSettings(ctx context.Context, req *tg.TLAccountChangeAuthorizationSettings) (r *tg.Bool, err error) {
	var _args AccountChangeAuthorizationSettingsArgs
	_args.Req = req
	var _result AccountChangeAuthorizationSettingsResult
	if err = p.c.Call(ctx, "account.changeAuthorizationSettings", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountInvalidateSignInCodes(ctx context.Context, req *tg.TLAccountInvalidateSignInCodes) (r *tg.Bool, err error) {
	var _args AccountInvalidateSignInCodesArgs
	_args.Req = req
	var _result AccountInvalidateSignInCodesResult
	if err = p.c.Call(ctx, "account.invalidateSignInCodes", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AuthToggleBan(ctx context.Context, req *tg.TLAuthToggleBan) (r *tg.PredefinedUser, err error) {
	var _args AuthToggleBanArgs
	_args.Req = req
	var _result AuthToggleBanResult
	if err = p.c.Call(ctx, "auth.toggleBan", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
