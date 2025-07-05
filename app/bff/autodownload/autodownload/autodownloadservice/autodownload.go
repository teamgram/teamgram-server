/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package autodownloadservice

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
	"account.getAutoDownloadSettings": kitex.NewMethodInfo(
		accountGetAutoDownloadSettingsHandler,
		newAccountGetAutoDownloadSettingsArgs,
		newAccountGetAutoDownloadSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.saveAutoDownloadSettings": kitex.NewMethodInfo(
		accountSaveAutoDownloadSettingsHandler,
		newAccountSaveAutoDownloadSettingsArgs,
		newAccountSaveAutoDownloadSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	autodownloadServiceServiceInfo                = NewServiceInfo()
	autodownloadServiceServiceInfoForClient       = NewServiceInfoForClient()
	autodownloadServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCAutoDownload", autodownloadServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCAutoDownload", autodownloadServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCAutoDownload", autodownloadServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return autodownloadServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return autodownloadServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return autodownloadServiceServiceInfoForClient
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
	serviceName := "RPCAutoDownload"
	handlerType := (*tg.RPCAutoDownload)(nil)
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
		"PackageName": "autodownload",
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

func accountGetAutoDownloadSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetAutoDownloadSettingsArgs)
	realResult := result.(*AccountGetAutoDownloadSettingsResult)
	success, err := handler.(tg.RPCAutoDownload).AccountGetAutoDownloadSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetAutoDownloadSettingsArgs() interface{} {
	return &AccountGetAutoDownloadSettingsArgs{}
}

func newAccountGetAutoDownloadSettingsResult() interface{} {
	return &AccountGetAutoDownloadSettingsResult{}
}

type AccountGetAutoDownloadSettingsArgs struct {
	Req *tg.TLAccountGetAutoDownloadSettings
}

func (p *AccountGetAutoDownloadSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetAutoDownloadSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetAutoDownloadSettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetAutoDownloadSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetAutoDownloadSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetAutoDownloadSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetAutoDownloadSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetAutoDownloadSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetAutoDownloadSettingsArgs_Req_DEFAULT *tg.TLAccountGetAutoDownloadSettings

func (p *AccountGetAutoDownloadSettingsArgs) GetReq() *tg.TLAccountGetAutoDownloadSettings {
	if !p.IsSetReq() {
		return AccountGetAutoDownloadSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetAutoDownloadSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetAutoDownloadSettingsResult struct {
	Success *tg.AccountAutoDownloadSettings
}

var AccountGetAutoDownloadSettingsResult_Success_DEFAULT *tg.AccountAutoDownloadSettings

func (p *AccountGetAutoDownloadSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetAutoDownloadSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetAutoDownloadSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountAutoDownloadSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAutoDownloadSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetAutoDownloadSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetAutoDownloadSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountAutoDownloadSettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetAutoDownloadSettingsResult) GetSuccess() *tg.AccountAutoDownloadSettings {
	if !p.IsSetSuccess() {
		return AccountGetAutoDownloadSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetAutoDownloadSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountAutoDownloadSettings)
}

func (p *AccountGetAutoDownloadSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetAutoDownloadSettingsResult) GetResult() interface{} {
	return p.Success
}

func accountSaveAutoDownloadSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSaveAutoDownloadSettingsArgs)
	realResult := result.(*AccountSaveAutoDownloadSettingsResult)
	success, err := handler.(tg.RPCAutoDownload).AccountSaveAutoDownloadSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSaveAutoDownloadSettingsArgs() interface{} {
	return &AccountSaveAutoDownloadSettingsArgs{}
}

func newAccountSaveAutoDownloadSettingsResult() interface{} {
	return &AccountSaveAutoDownloadSettingsResult{}
}

type AccountSaveAutoDownloadSettingsArgs struct {
	Req *tg.TLAccountSaveAutoDownloadSettings
}

func (p *AccountSaveAutoDownloadSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSaveAutoDownloadSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSaveAutoDownloadSettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSaveAutoDownloadSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSaveAutoDownloadSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSaveAutoDownloadSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSaveAutoDownloadSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSaveAutoDownloadSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSaveAutoDownloadSettingsArgs_Req_DEFAULT *tg.TLAccountSaveAutoDownloadSettings

func (p *AccountSaveAutoDownloadSettingsArgs) GetReq() *tg.TLAccountSaveAutoDownloadSettings {
	if !p.IsSetReq() {
		return AccountSaveAutoDownloadSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSaveAutoDownloadSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSaveAutoDownloadSettingsResult struct {
	Success *tg.Bool
}

var AccountSaveAutoDownloadSettingsResult_Success_DEFAULT *tg.Bool

func (p *AccountSaveAutoDownloadSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSaveAutoDownloadSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSaveAutoDownloadSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSaveAutoDownloadSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSaveAutoDownloadSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSaveAutoDownloadSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSaveAutoDownloadSettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountSaveAutoDownloadSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSaveAutoDownloadSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountSaveAutoDownloadSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSaveAutoDownloadSettingsResult) GetResult() interface{} {
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

func (p *kClient) AccountGetAutoDownloadSettings(ctx context.Context, req *tg.TLAccountGetAutoDownloadSettings) (r *tg.AccountAutoDownloadSettings, err error) {
	// var _args AccountGetAutoDownloadSettingsArgs
	// _args.Req = req
	// var _result AccountGetAutoDownloadSettingsResult

	_result := new(tg.AccountAutoDownloadSettings)
	if err = p.c.Call(ctx, "account.getAutoDownloadSettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountSaveAutoDownloadSettings(ctx context.Context, req *tg.TLAccountSaveAutoDownloadSettings) (r *tg.Bool, err error) {
	// var _args AccountSaveAutoDownloadSettingsArgs
	// _args.Req = req
	// var _result AccountSaveAutoDownloadSettingsResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.saveAutoDownloadSettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
