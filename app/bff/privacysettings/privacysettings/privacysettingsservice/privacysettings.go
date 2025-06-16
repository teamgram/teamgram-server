/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package privacysettingsservice

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
	"account.getPrivacy": kitex.NewMethodInfo(
		accountGetPrivacyHandler,
		newAccountGetPrivacyArgs,
		newAccountGetPrivacyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.setPrivacy": kitex.NewMethodInfo(
		accountSetPrivacyHandler,
		newAccountSetPrivacyArgs,
		newAccountSetPrivacyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.getGlobalPrivacySettings": kitex.NewMethodInfo(
		accountGetGlobalPrivacySettingsHandler,
		newAccountGetGlobalPrivacySettingsArgs,
		newAccountGetGlobalPrivacySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.setGlobalPrivacySettings": kitex.NewMethodInfo(
		accountSetGlobalPrivacySettingsHandler,
		newAccountSetGlobalPrivacySettingsArgs,
		newAccountSetGlobalPrivacySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"users.getIsPremiumRequiredToContact": kitex.NewMethodInfo(
		usersGetIsPremiumRequiredToContactHandler,
		newUsersGetIsPremiumRequiredToContactArgs,
		newUsersGetIsPremiumRequiredToContactResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.setDefaultHistoryTTL": kitex.NewMethodInfo(
		messagesSetDefaultHistoryTTLHandler,
		newMessagesSetDefaultHistoryTTLArgs,
		newMessagesSetDefaultHistoryTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getDefaultHistoryTTL": kitex.NewMethodInfo(
		messagesGetDefaultHistoryTTLHandler,
		newMessagesGetDefaultHistoryTTLArgs,
		newMessagesGetDefaultHistoryTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	privacysettingsServiceServiceInfo                = NewServiceInfo()
	privacysettingsServiceServiceInfoForClient       = NewServiceInfoForClient()
	privacysettingsServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return privacysettingsServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return privacysettingsServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return privacysettingsServiceServiceInfoForClient
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
	serviceName := "RPCPrivacySettings"
	handlerType := (*tg.RPCPrivacySettings)(nil)
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
		"PackageName": "privacysettings",
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

func accountGetPrivacyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetPrivacyArgs)
	realResult := result.(*AccountGetPrivacyResult)
	success, err := handler.(tg.RPCPrivacySettings).AccountGetPrivacy(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetPrivacyArgs() interface{} {
	return &AccountGetPrivacyArgs{}
}

func newAccountGetPrivacyResult() interface{} {
	return &AccountGetPrivacyResult{}
}

type AccountGetPrivacyArgs struct {
	Req *tg.TLAccountGetPrivacy
}

func (p *AccountGetPrivacyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetPrivacyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetPrivacyArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetPrivacy)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetPrivacyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetPrivacyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetPrivacyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetPrivacy)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetPrivacyArgs_Req_DEFAULT *tg.TLAccountGetPrivacy

func (p *AccountGetPrivacyArgs) GetReq() *tg.TLAccountGetPrivacy {
	if !p.IsSetReq() {
		return AccountGetPrivacyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetPrivacyArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetPrivacyResult struct {
	Success *tg.AccountPrivacyRules
}

var AccountGetPrivacyResult_Success_DEFAULT *tg.AccountPrivacyRules

func (p *AccountGetPrivacyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetPrivacyResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetPrivacyResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountPrivacyRules)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetPrivacyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetPrivacyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetPrivacyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountPrivacyRules)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetPrivacyResult) GetSuccess() *tg.AccountPrivacyRules {
	if !p.IsSetSuccess() {
		return AccountGetPrivacyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetPrivacyResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountPrivacyRules)
}

func (p *AccountGetPrivacyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetPrivacyResult) GetResult() interface{} {
	return p.Success
}

func accountSetPrivacyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSetPrivacyArgs)
	realResult := result.(*AccountSetPrivacyResult)
	success, err := handler.(tg.RPCPrivacySettings).AccountSetPrivacy(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSetPrivacyArgs() interface{} {
	return &AccountSetPrivacyArgs{}
}

func newAccountSetPrivacyResult() interface{} {
	return &AccountSetPrivacyResult{}
}

type AccountSetPrivacyArgs struct {
	Req *tg.TLAccountSetPrivacy
}

func (p *AccountSetPrivacyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSetPrivacyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSetPrivacyArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSetPrivacy)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSetPrivacyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSetPrivacyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSetPrivacyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSetPrivacy)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSetPrivacyArgs_Req_DEFAULT *tg.TLAccountSetPrivacy

func (p *AccountSetPrivacyArgs) GetReq() *tg.TLAccountSetPrivacy {
	if !p.IsSetReq() {
		return AccountSetPrivacyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSetPrivacyArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSetPrivacyResult struct {
	Success *tg.AccountPrivacyRules
}

var AccountSetPrivacyResult_Success_DEFAULT *tg.AccountPrivacyRules

func (p *AccountSetPrivacyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSetPrivacyResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSetPrivacyResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountPrivacyRules)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetPrivacyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSetPrivacyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSetPrivacyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountPrivacyRules)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetPrivacyResult) GetSuccess() *tg.AccountPrivacyRules {
	if !p.IsSetSuccess() {
		return AccountSetPrivacyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSetPrivacyResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountPrivacyRules)
}

func (p *AccountSetPrivacyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSetPrivacyResult) GetResult() interface{} {
	return p.Success
}

func accountGetGlobalPrivacySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetGlobalPrivacySettingsArgs)
	realResult := result.(*AccountGetGlobalPrivacySettingsResult)
	success, err := handler.(tg.RPCPrivacySettings).AccountGetGlobalPrivacySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetGlobalPrivacySettingsArgs() interface{} {
	return &AccountGetGlobalPrivacySettingsArgs{}
}

func newAccountGetGlobalPrivacySettingsResult() interface{} {
	return &AccountGetGlobalPrivacySettingsResult{}
}

type AccountGetGlobalPrivacySettingsArgs struct {
	Req *tg.TLAccountGetGlobalPrivacySettings
}

func (p *AccountGetGlobalPrivacySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetGlobalPrivacySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetGlobalPrivacySettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetGlobalPrivacySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetGlobalPrivacySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetGlobalPrivacySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetGlobalPrivacySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetGlobalPrivacySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetGlobalPrivacySettingsArgs_Req_DEFAULT *tg.TLAccountGetGlobalPrivacySettings

func (p *AccountGetGlobalPrivacySettingsArgs) GetReq() *tg.TLAccountGetGlobalPrivacySettings {
	if !p.IsSetReq() {
		return AccountGetGlobalPrivacySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetGlobalPrivacySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetGlobalPrivacySettingsResult struct {
	Success *tg.GlobalPrivacySettings
}

var AccountGetGlobalPrivacySettingsResult_Success_DEFAULT *tg.GlobalPrivacySettings

func (p *AccountGetGlobalPrivacySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetGlobalPrivacySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetGlobalPrivacySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.GlobalPrivacySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetGlobalPrivacySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetGlobalPrivacySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetGlobalPrivacySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.GlobalPrivacySettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetGlobalPrivacySettingsResult) GetSuccess() *tg.GlobalPrivacySettings {
	if !p.IsSetSuccess() {
		return AccountGetGlobalPrivacySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetGlobalPrivacySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.GlobalPrivacySettings)
}

func (p *AccountGetGlobalPrivacySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetGlobalPrivacySettingsResult) GetResult() interface{} {
	return p.Success
}

func accountSetGlobalPrivacySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSetGlobalPrivacySettingsArgs)
	realResult := result.(*AccountSetGlobalPrivacySettingsResult)
	success, err := handler.(tg.RPCPrivacySettings).AccountSetGlobalPrivacySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSetGlobalPrivacySettingsArgs() interface{} {
	return &AccountSetGlobalPrivacySettingsArgs{}
}

func newAccountSetGlobalPrivacySettingsResult() interface{} {
	return &AccountSetGlobalPrivacySettingsResult{}
}

type AccountSetGlobalPrivacySettingsArgs struct {
	Req *tg.TLAccountSetGlobalPrivacySettings
}

func (p *AccountSetGlobalPrivacySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSetGlobalPrivacySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSetGlobalPrivacySettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSetGlobalPrivacySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSetGlobalPrivacySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSetGlobalPrivacySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSetGlobalPrivacySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSetGlobalPrivacySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSetGlobalPrivacySettingsArgs_Req_DEFAULT *tg.TLAccountSetGlobalPrivacySettings

func (p *AccountSetGlobalPrivacySettingsArgs) GetReq() *tg.TLAccountSetGlobalPrivacySettings {
	if !p.IsSetReq() {
		return AccountSetGlobalPrivacySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSetGlobalPrivacySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSetGlobalPrivacySettingsResult struct {
	Success *tg.GlobalPrivacySettings
}

var AccountSetGlobalPrivacySettingsResult_Success_DEFAULT *tg.GlobalPrivacySettings

func (p *AccountSetGlobalPrivacySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSetGlobalPrivacySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSetGlobalPrivacySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.GlobalPrivacySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetGlobalPrivacySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSetGlobalPrivacySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSetGlobalPrivacySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.GlobalPrivacySettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetGlobalPrivacySettingsResult) GetSuccess() *tg.GlobalPrivacySettings {
	if !p.IsSetSuccess() {
		return AccountSetGlobalPrivacySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSetGlobalPrivacySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.GlobalPrivacySettings)
}

func (p *AccountSetGlobalPrivacySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSetGlobalPrivacySettingsResult) GetResult() interface{} {
	return p.Success
}

func usersGetIsPremiumRequiredToContactHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UsersGetIsPremiumRequiredToContactArgs)
	realResult := result.(*UsersGetIsPremiumRequiredToContactResult)
	success, err := handler.(tg.RPCPrivacySettings).UsersGetIsPremiumRequiredToContact(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUsersGetIsPremiumRequiredToContactArgs() interface{} {
	return &UsersGetIsPremiumRequiredToContactArgs{}
}

func newUsersGetIsPremiumRequiredToContactResult() interface{} {
	return &UsersGetIsPremiumRequiredToContactResult{}
}

type UsersGetIsPremiumRequiredToContactArgs struct {
	Req *tg.TLUsersGetIsPremiumRequiredToContact
}

func (p *UsersGetIsPremiumRequiredToContactArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UsersGetIsPremiumRequiredToContactArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UsersGetIsPremiumRequiredToContactArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUsersGetIsPremiumRequiredToContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UsersGetIsPremiumRequiredToContactArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UsersGetIsPremiumRequiredToContactArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UsersGetIsPremiumRequiredToContactArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUsersGetIsPremiumRequiredToContact)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UsersGetIsPremiumRequiredToContactArgs_Req_DEFAULT *tg.TLUsersGetIsPremiumRequiredToContact

func (p *UsersGetIsPremiumRequiredToContactArgs) GetReq() *tg.TLUsersGetIsPremiumRequiredToContact {
	if !p.IsSetReq() {
		return UsersGetIsPremiumRequiredToContactArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UsersGetIsPremiumRequiredToContactArgs) IsSetReq() bool {
	return p.Req != nil
}

type UsersGetIsPremiumRequiredToContactResult struct {
	Success *tg.VectorBool
}

var UsersGetIsPremiumRequiredToContactResult_Success_DEFAULT *tg.VectorBool

func (p *UsersGetIsPremiumRequiredToContactResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UsersGetIsPremiumRequiredToContactResult")
	}
	return json.Marshal(p.Success)
}

func (p *UsersGetIsPremiumRequiredToContactResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorBool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersGetIsPremiumRequiredToContactResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UsersGetIsPremiumRequiredToContactResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UsersGetIsPremiumRequiredToContactResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorBool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersGetIsPremiumRequiredToContactResult) GetSuccess() *tg.VectorBool {
	if !p.IsSetSuccess() {
		return UsersGetIsPremiumRequiredToContactResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UsersGetIsPremiumRequiredToContactResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorBool)
}

func (p *UsersGetIsPremiumRequiredToContactResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UsersGetIsPremiumRequiredToContactResult) GetResult() interface{} {
	return p.Success
}

func messagesSetDefaultHistoryTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSetDefaultHistoryTTLArgs)
	realResult := result.(*MessagesSetDefaultHistoryTTLResult)
	success, err := handler.(tg.RPCPrivacySettings).MessagesSetDefaultHistoryTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSetDefaultHistoryTTLArgs() interface{} {
	return &MessagesSetDefaultHistoryTTLArgs{}
}

func newMessagesSetDefaultHistoryTTLResult() interface{} {
	return &MessagesSetDefaultHistoryTTLResult{}
}

type MessagesSetDefaultHistoryTTLArgs struct {
	Req *tg.TLMessagesSetDefaultHistoryTTL
}

func (p *MessagesSetDefaultHistoryTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSetDefaultHistoryTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSetDefaultHistoryTTLArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSetDefaultHistoryTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSetDefaultHistoryTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSetDefaultHistoryTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSetDefaultHistoryTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSetDefaultHistoryTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSetDefaultHistoryTTLArgs_Req_DEFAULT *tg.TLMessagesSetDefaultHistoryTTL

func (p *MessagesSetDefaultHistoryTTLArgs) GetReq() *tg.TLMessagesSetDefaultHistoryTTL {
	if !p.IsSetReq() {
		return MessagesSetDefaultHistoryTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSetDefaultHistoryTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSetDefaultHistoryTTLResult struct {
	Success *tg.Bool
}

var MessagesSetDefaultHistoryTTLResult_Success_DEFAULT *tg.Bool

func (p *MessagesSetDefaultHistoryTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSetDefaultHistoryTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSetDefaultHistoryTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSetDefaultHistoryTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSetDefaultHistoryTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSetDefaultHistoryTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSetDefaultHistoryTTLResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesSetDefaultHistoryTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSetDefaultHistoryTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesSetDefaultHistoryTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSetDefaultHistoryTTLResult) GetResult() interface{} {
	return p.Success
}

func messagesGetDefaultHistoryTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetDefaultHistoryTTLArgs)
	realResult := result.(*MessagesGetDefaultHistoryTTLResult)
	success, err := handler.(tg.RPCPrivacySettings).MessagesGetDefaultHistoryTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetDefaultHistoryTTLArgs() interface{} {
	return &MessagesGetDefaultHistoryTTLArgs{}
}

func newMessagesGetDefaultHistoryTTLResult() interface{} {
	return &MessagesGetDefaultHistoryTTLResult{}
}

type MessagesGetDefaultHistoryTTLArgs struct {
	Req *tg.TLMessagesGetDefaultHistoryTTL
}

func (p *MessagesGetDefaultHistoryTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetDefaultHistoryTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetDefaultHistoryTTLArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetDefaultHistoryTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetDefaultHistoryTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetDefaultHistoryTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetDefaultHistoryTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetDefaultHistoryTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetDefaultHistoryTTLArgs_Req_DEFAULT *tg.TLMessagesGetDefaultHistoryTTL

func (p *MessagesGetDefaultHistoryTTLArgs) GetReq() *tg.TLMessagesGetDefaultHistoryTTL {
	if !p.IsSetReq() {
		return MessagesGetDefaultHistoryTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetDefaultHistoryTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetDefaultHistoryTTLResult struct {
	Success *tg.DefaultHistoryTTL
}

var MessagesGetDefaultHistoryTTLResult_Success_DEFAULT *tg.DefaultHistoryTTL

func (p *MessagesGetDefaultHistoryTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetDefaultHistoryTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetDefaultHistoryTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.DefaultHistoryTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetDefaultHistoryTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetDefaultHistoryTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetDefaultHistoryTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.DefaultHistoryTTL)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetDefaultHistoryTTLResult) GetSuccess() *tg.DefaultHistoryTTL {
	if !p.IsSetSuccess() {
		return MessagesGetDefaultHistoryTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetDefaultHistoryTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.DefaultHistoryTTL)
}

func (p *MessagesGetDefaultHistoryTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetDefaultHistoryTTLResult) GetResult() interface{} {
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

func (p *kClient) AccountGetPrivacy(ctx context.Context, req *tg.TLAccountGetPrivacy) (r *tg.AccountPrivacyRules, err error) {
	var _args AccountGetPrivacyArgs
	_args.Req = req
	var _result AccountGetPrivacyResult
	if err = p.c.Call(ctx, "account.getPrivacy", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountSetPrivacy(ctx context.Context, req *tg.TLAccountSetPrivacy) (r *tg.AccountPrivacyRules, err error) {
	var _args AccountSetPrivacyArgs
	_args.Req = req
	var _result AccountSetPrivacyResult
	if err = p.c.Call(ctx, "account.setPrivacy", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountGetGlobalPrivacySettings(ctx context.Context, req *tg.TLAccountGetGlobalPrivacySettings) (r *tg.GlobalPrivacySettings, err error) {
	var _args AccountGetGlobalPrivacySettingsArgs
	_args.Req = req
	var _result AccountGetGlobalPrivacySettingsResult
	if err = p.c.Call(ctx, "account.getGlobalPrivacySettings", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountSetGlobalPrivacySettings(ctx context.Context, req *tg.TLAccountSetGlobalPrivacySettings) (r *tg.GlobalPrivacySettings, err error) {
	var _args AccountSetGlobalPrivacySettingsArgs
	_args.Req = req
	var _result AccountSetGlobalPrivacySettingsResult
	if err = p.c.Call(ctx, "account.setGlobalPrivacySettings", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UsersGetIsPremiumRequiredToContact(ctx context.Context, req *tg.TLUsersGetIsPremiumRequiredToContact) (r *tg.VectorBool, err error) {
	var _args UsersGetIsPremiumRequiredToContactArgs
	_args.Req = req
	var _result UsersGetIsPremiumRequiredToContactResult
	if err = p.c.Call(ctx, "users.getIsPremiumRequiredToContact", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesSetDefaultHistoryTTL(ctx context.Context, req *tg.TLMessagesSetDefaultHistoryTTL) (r *tg.Bool, err error) {
	var _args MessagesSetDefaultHistoryTTLArgs
	_args.Req = req
	var _result MessagesSetDefaultHistoryTTLResult
	if err = p.c.Call(ctx, "messages.setDefaultHistoryTTL", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetDefaultHistoryTTL(ctx context.Context, req *tg.TLMessagesGetDefaultHistoryTTL) (r *tg.DefaultHistoryTTL, err error) {
	var _args MessagesGetDefaultHistoryTTLArgs
	_args.Req = req
	var _result MessagesGetDefaultHistoryTTLResult
	if err = p.c.Call(ctx, "messages.getDefaultHistoryTTL", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
