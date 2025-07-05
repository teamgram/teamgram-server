/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package notificationservice

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
	"account.registerDevice": kitex.NewMethodInfo(
		accountRegisterDeviceHandler,
		newAccountRegisterDeviceArgs,
		newAccountRegisterDeviceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.unregisterDevice": kitex.NewMethodInfo(
		accountUnregisterDeviceHandler,
		newAccountUnregisterDeviceArgs,
		newAccountUnregisterDeviceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.updateNotifySettings": kitex.NewMethodInfo(
		accountUpdateNotifySettingsHandler,
		newAccountUpdateNotifySettingsArgs,
		newAccountUpdateNotifySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.getNotifySettings": kitex.NewMethodInfo(
		accountGetNotifySettingsHandler,
		newAccountGetNotifySettingsArgs,
		newAccountGetNotifySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.resetNotifySettings": kitex.NewMethodInfo(
		accountResetNotifySettingsHandler,
		newAccountResetNotifySettingsArgs,
		newAccountResetNotifySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.updateDeviceLocked": kitex.NewMethodInfo(
		accountUpdateDeviceLockedHandler,
		newAccountUpdateDeviceLockedArgs,
		newAccountUpdateDeviceLockedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.getNotifyExceptions": kitex.NewMethodInfo(
		accountGetNotifyExceptionsHandler,
		newAccountGetNotifyExceptionsArgs,
		newAccountGetNotifyExceptionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	notificationServiceServiceInfo                = NewServiceInfo()
	notificationServiceServiceInfoForClient       = NewServiceInfoForClient()
	notificationServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCNotification", notificationServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCNotification", notificationServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCNotification", notificationServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return notificationServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return notificationServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return notificationServiceServiceInfoForClient
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
	serviceName := "RPCNotification"
	handlerType := (*tg.RPCNotification)(nil)
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
		"PackageName": "notification",
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

func accountRegisterDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountRegisterDeviceArgs)
	realResult := result.(*AccountRegisterDeviceResult)
	success, err := handler.(tg.RPCNotification).AccountRegisterDevice(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountRegisterDeviceArgs() interface{} {
	return &AccountRegisterDeviceArgs{}
}

func newAccountRegisterDeviceResult() interface{} {
	return &AccountRegisterDeviceResult{}
}

type AccountRegisterDeviceArgs struct {
	Req *tg.TLAccountRegisterDevice
}

func (p *AccountRegisterDeviceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountRegisterDeviceArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountRegisterDeviceArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountRegisterDevice)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountRegisterDeviceArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountRegisterDeviceArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountRegisterDeviceArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountRegisterDevice)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountRegisterDeviceArgs_Req_DEFAULT *tg.TLAccountRegisterDevice

func (p *AccountRegisterDeviceArgs) GetReq() *tg.TLAccountRegisterDevice {
	if !p.IsSetReq() {
		return AccountRegisterDeviceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountRegisterDeviceArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountRegisterDeviceResult struct {
	Success *tg.Bool
}

var AccountRegisterDeviceResult_Success_DEFAULT *tg.Bool

func (p *AccountRegisterDeviceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountRegisterDeviceResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountRegisterDeviceResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountRegisterDeviceResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountRegisterDeviceResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountRegisterDeviceResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountRegisterDeviceResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountRegisterDeviceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountRegisterDeviceResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountRegisterDeviceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountRegisterDeviceResult) GetResult() interface{} {
	return p.Success
}

func accountUnregisterDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUnregisterDeviceArgs)
	realResult := result.(*AccountUnregisterDeviceResult)
	success, err := handler.(tg.RPCNotification).AccountUnregisterDevice(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUnregisterDeviceArgs() interface{} {
	return &AccountUnregisterDeviceArgs{}
}

func newAccountUnregisterDeviceResult() interface{} {
	return &AccountUnregisterDeviceResult{}
}

type AccountUnregisterDeviceArgs struct {
	Req *tg.TLAccountUnregisterDevice
}

func (p *AccountUnregisterDeviceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUnregisterDeviceArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUnregisterDeviceArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUnregisterDevice)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUnregisterDeviceArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUnregisterDeviceArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUnregisterDeviceArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUnregisterDevice)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUnregisterDeviceArgs_Req_DEFAULT *tg.TLAccountUnregisterDevice

func (p *AccountUnregisterDeviceArgs) GetReq() *tg.TLAccountUnregisterDevice {
	if !p.IsSetReq() {
		return AccountUnregisterDeviceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUnregisterDeviceArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUnregisterDeviceResult struct {
	Success *tg.Bool
}

var AccountUnregisterDeviceResult_Success_DEFAULT *tg.Bool

func (p *AccountUnregisterDeviceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUnregisterDeviceResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUnregisterDeviceResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUnregisterDeviceResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUnregisterDeviceResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUnregisterDeviceResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUnregisterDeviceResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountUnregisterDeviceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUnregisterDeviceResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountUnregisterDeviceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUnregisterDeviceResult) GetResult() interface{} {
	return p.Success
}

func accountUpdateNotifySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUpdateNotifySettingsArgs)
	realResult := result.(*AccountUpdateNotifySettingsResult)
	success, err := handler.(tg.RPCNotification).AccountUpdateNotifySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUpdateNotifySettingsArgs() interface{} {
	return &AccountUpdateNotifySettingsArgs{}
}

func newAccountUpdateNotifySettingsResult() interface{} {
	return &AccountUpdateNotifySettingsResult{}
}

type AccountUpdateNotifySettingsArgs struct {
	Req *tg.TLAccountUpdateNotifySettings
}

func (p *AccountUpdateNotifySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUpdateNotifySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUpdateNotifySettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUpdateNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUpdateNotifySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUpdateNotifySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUpdateNotifySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUpdateNotifySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUpdateNotifySettingsArgs_Req_DEFAULT *tg.TLAccountUpdateNotifySettings

func (p *AccountUpdateNotifySettingsArgs) GetReq() *tg.TLAccountUpdateNotifySettings {
	if !p.IsSetReq() {
		return AccountUpdateNotifySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUpdateNotifySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUpdateNotifySettingsResult struct {
	Success *tg.Bool
}

var AccountUpdateNotifySettingsResult_Success_DEFAULT *tg.Bool

func (p *AccountUpdateNotifySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUpdateNotifySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUpdateNotifySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateNotifySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUpdateNotifySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUpdateNotifySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateNotifySettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountUpdateNotifySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUpdateNotifySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountUpdateNotifySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUpdateNotifySettingsResult) GetResult() interface{} {
	return p.Success
}

func accountGetNotifySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetNotifySettingsArgs)
	realResult := result.(*AccountGetNotifySettingsResult)
	success, err := handler.(tg.RPCNotification).AccountGetNotifySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetNotifySettingsArgs() interface{} {
	return &AccountGetNotifySettingsArgs{}
}

func newAccountGetNotifySettingsResult() interface{} {
	return &AccountGetNotifySettingsResult{}
}

type AccountGetNotifySettingsArgs struct {
	Req *tg.TLAccountGetNotifySettings
}

func (p *AccountGetNotifySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetNotifySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetNotifySettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetNotifySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetNotifySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetNotifySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetNotifySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetNotifySettingsArgs_Req_DEFAULT *tg.TLAccountGetNotifySettings

func (p *AccountGetNotifySettingsArgs) GetReq() *tg.TLAccountGetNotifySettings {
	if !p.IsSetReq() {
		return AccountGetNotifySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetNotifySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetNotifySettingsResult struct {
	Success *tg.PeerNotifySettings
}

var AccountGetNotifySettingsResult_Success_DEFAULT *tg.PeerNotifySettings

func (p *AccountGetNotifySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetNotifySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetNotifySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.PeerNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetNotifySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetNotifySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetNotifySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.PeerNotifySettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetNotifySettingsResult) GetSuccess() *tg.PeerNotifySettings {
	if !p.IsSetSuccess() {
		return AccountGetNotifySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetNotifySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.PeerNotifySettings)
}

func (p *AccountGetNotifySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetNotifySettingsResult) GetResult() interface{} {
	return p.Success
}

func accountResetNotifySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountResetNotifySettingsArgs)
	realResult := result.(*AccountResetNotifySettingsResult)
	success, err := handler.(tg.RPCNotification).AccountResetNotifySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountResetNotifySettingsArgs() interface{} {
	return &AccountResetNotifySettingsArgs{}
}

func newAccountResetNotifySettingsResult() interface{} {
	return &AccountResetNotifySettingsResult{}
}

type AccountResetNotifySettingsArgs struct {
	Req *tg.TLAccountResetNotifySettings
}

func (p *AccountResetNotifySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountResetNotifySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountResetNotifySettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountResetNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountResetNotifySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountResetNotifySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountResetNotifySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountResetNotifySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountResetNotifySettingsArgs_Req_DEFAULT *tg.TLAccountResetNotifySettings

func (p *AccountResetNotifySettingsArgs) GetReq() *tg.TLAccountResetNotifySettings {
	if !p.IsSetReq() {
		return AccountResetNotifySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountResetNotifySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountResetNotifySettingsResult struct {
	Success *tg.Bool
}

var AccountResetNotifySettingsResult_Success_DEFAULT *tg.Bool

func (p *AccountResetNotifySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountResetNotifySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountResetNotifySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountResetNotifySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountResetNotifySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountResetNotifySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountResetNotifySettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountResetNotifySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountResetNotifySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountResetNotifySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountResetNotifySettingsResult) GetResult() interface{} {
	return p.Success
}

func accountUpdateDeviceLockedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUpdateDeviceLockedArgs)
	realResult := result.(*AccountUpdateDeviceLockedResult)
	success, err := handler.(tg.RPCNotification).AccountUpdateDeviceLocked(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUpdateDeviceLockedArgs() interface{} {
	return &AccountUpdateDeviceLockedArgs{}
}

func newAccountUpdateDeviceLockedResult() interface{} {
	return &AccountUpdateDeviceLockedResult{}
}

type AccountUpdateDeviceLockedArgs struct {
	Req *tg.TLAccountUpdateDeviceLocked
}

func (p *AccountUpdateDeviceLockedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUpdateDeviceLockedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUpdateDeviceLockedArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUpdateDeviceLocked)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUpdateDeviceLockedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUpdateDeviceLockedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUpdateDeviceLockedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUpdateDeviceLocked)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUpdateDeviceLockedArgs_Req_DEFAULT *tg.TLAccountUpdateDeviceLocked

func (p *AccountUpdateDeviceLockedArgs) GetReq() *tg.TLAccountUpdateDeviceLocked {
	if !p.IsSetReq() {
		return AccountUpdateDeviceLockedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUpdateDeviceLockedArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUpdateDeviceLockedResult struct {
	Success *tg.Bool
}

var AccountUpdateDeviceLockedResult_Success_DEFAULT *tg.Bool

func (p *AccountUpdateDeviceLockedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUpdateDeviceLockedResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUpdateDeviceLockedResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateDeviceLockedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUpdateDeviceLockedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUpdateDeviceLockedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateDeviceLockedResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountUpdateDeviceLockedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUpdateDeviceLockedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountUpdateDeviceLockedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUpdateDeviceLockedResult) GetResult() interface{} {
	return p.Success
}

func accountGetNotifyExceptionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetNotifyExceptionsArgs)
	realResult := result.(*AccountGetNotifyExceptionsResult)
	success, err := handler.(tg.RPCNotification).AccountGetNotifyExceptions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetNotifyExceptionsArgs() interface{} {
	return &AccountGetNotifyExceptionsArgs{}
}

func newAccountGetNotifyExceptionsResult() interface{} {
	return &AccountGetNotifyExceptionsResult{}
}

type AccountGetNotifyExceptionsArgs struct {
	Req *tg.TLAccountGetNotifyExceptions
}

func (p *AccountGetNotifyExceptionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetNotifyExceptionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetNotifyExceptionsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetNotifyExceptions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetNotifyExceptionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetNotifyExceptionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetNotifyExceptionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetNotifyExceptions)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetNotifyExceptionsArgs_Req_DEFAULT *tg.TLAccountGetNotifyExceptions

func (p *AccountGetNotifyExceptionsArgs) GetReq() *tg.TLAccountGetNotifyExceptions {
	if !p.IsSetReq() {
		return AccountGetNotifyExceptionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetNotifyExceptionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetNotifyExceptionsResult struct {
	Success *tg.Updates
}

var AccountGetNotifyExceptionsResult_Success_DEFAULT *tg.Updates

func (p *AccountGetNotifyExceptionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetNotifyExceptionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetNotifyExceptionsResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetNotifyExceptionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetNotifyExceptionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetNotifyExceptionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetNotifyExceptionsResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return AccountGetNotifyExceptionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetNotifyExceptionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *AccountGetNotifyExceptionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetNotifyExceptionsResult) GetResult() interface{} {
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

func (p *kClient) AccountRegisterDevice(ctx context.Context, req *tg.TLAccountRegisterDevice) (r *tg.Bool, err error) {
	// var _args AccountRegisterDeviceArgs
	// _args.Req = req
	// var _result AccountRegisterDeviceResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.registerDevice", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountUnregisterDevice(ctx context.Context, req *tg.TLAccountUnregisterDevice) (r *tg.Bool, err error) {
	// var _args AccountUnregisterDeviceArgs
	// _args.Req = req
	// var _result AccountUnregisterDeviceResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.unregisterDevice", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountUpdateNotifySettings(ctx context.Context, req *tg.TLAccountUpdateNotifySettings) (r *tg.Bool, err error) {
	// var _args AccountUpdateNotifySettingsArgs
	// _args.Req = req
	// var _result AccountUpdateNotifySettingsResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.updateNotifySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountGetNotifySettings(ctx context.Context, req *tg.TLAccountGetNotifySettings) (r *tg.PeerNotifySettings, err error) {
	// var _args AccountGetNotifySettingsArgs
	// _args.Req = req
	// var _result AccountGetNotifySettingsResult

	_result := new(tg.PeerNotifySettings)
	if err = p.c.Call(ctx, "account.getNotifySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountResetNotifySettings(ctx context.Context, req *tg.TLAccountResetNotifySettings) (r *tg.Bool, err error) {
	// var _args AccountResetNotifySettingsArgs
	// _args.Req = req
	// var _result AccountResetNotifySettingsResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.resetNotifySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountUpdateDeviceLocked(ctx context.Context, req *tg.TLAccountUpdateDeviceLocked) (r *tg.Bool, err error) {
	// var _args AccountUpdateDeviceLockedArgs
	// _args.Req = req
	// var _result AccountUpdateDeviceLockedResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.updateDeviceLocked", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountGetNotifyExceptions(ctx context.Context, req *tg.TLAccountGetNotifyExceptions) (r *tg.Updates, err error) {
	// var _args AccountGetNotifyExceptionsArgs
	// _args.Req = req
	// var _result AccountGetNotifyExceptionsResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "account.getNotifyExceptions", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
