/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package nsfwservice

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
	"account.setContentSettings": kitex.NewMethodInfo(
		accountSetContentSettingsHandler,
		newAccountSetContentSettingsArgs,
		newAccountSetContentSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.getContentSettings": kitex.NewMethodInfo(
		accountGetContentSettingsHandler,
		newAccountGetContentSettingsArgs,
		newAccountGetContentSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	nsfwServiceServiceInfo                = NewServiceInfo()
	nsfwServiceServiceInfoForClient       = NewServiceInfoForClient()
	nsfwServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return nsfwServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return nsfwServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return nsfwServiceServiceInfoForClient
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
	serviceName := "RPCNsfw"
	handlerType := (*tg.RPCNsfw)(nil)
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
		"PackageName": "nsfw",
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

func accountSetContentSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSetContentSettingsArgs)
	realResult := result.(*AccountSetContentSettingsResult)
	success, err := handler.(tg.RPCNsfw).AccountSetContentSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSetContentSettingsArgs() interface{} {
	return &AccountSetContentSettingsArgs{}
}

func newAccountSetContentSettingsResult() interface{} {
	return &AccountSetContentSettingsResult{}
}

type AccountSetContentSettingsArgs struct {
	Req *tg.TLAccountSetContentSettings
}

func (p *AccountSetContentSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSetContentSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSetContentSettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSetContentSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSetContentSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSetContentSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSetContentSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSetContentSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSetContentSettingsArgs_Req_DEFAULT *tg.TLAccountSetContentSettings

func (p *AccountSetContentSettingsArgs) GetReq() *tg.TLAccountSetContentSettings {
	if !p.IsSetReq() {
		return AccountSetContentSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSetContentSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSetContentSettingsResult struct {
	Success *tg.Bool
}

var AccountSetContentSettingsResult_Success_DEFAULT *tg.Bool

func (p *AccountSetContentSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSetContentSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSetContentSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetContentSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSetContentSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSetContentSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetContentSettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountSetContentSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSetContentSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountSetContentSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSetContentSettingsResult) GetResult() interface{} {
	return p.Success
}

func accountGetContentSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetContentSettingsArgs)
	realResult := result.(*AccountGetContentSettingsResult)
	success, err := handler.(tg.RPCNsfw).AccountGetContentSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetContentSettingsArgs() interface{} {
	return &AccountGetContentSettingsArgs{}
}

func newAccountGetContentSettingsResult() interface{} {
	return &AccountGetContentSettingsResult{}
}

type AccountGetContentSettingsArgs struct {
	Req *tg.TLAccountGetContentSettings
}

func (p *AccountGetContentSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetContentSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetContentSettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetContentSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetContentSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetContentSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetContentSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetContentSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetContentSettingsArgs_Req_DEFAULT *tg.TLAccountGetContentSettings

func (p *AccountGetContentSettingsArgs) GetReq() *tg.TLAccountGetContentSettings {
	if !p.IsSetReq() {
		return AccountGetContentSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetContentSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetContentSettingsResult struct {
	Success *tg.AccountContentSettings
}

var AccountGetContentSettingsResult_Success_DEFAULT *tg.AccountContentSettings

func (p *AccountGetContentSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetContentSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetContentSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountContentSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetContentSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetContentSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetContentSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountContentSettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetContentSettingsResult) GetSuccess() *tg.AccountContentSettings {
	if !p.IsSetSuccess() {
		return AccountGetContentSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetContentSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountContentSettings)
}

func (p *AccountGetContentSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetContentSettingsResult) GetResult() interface{} {
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

func (p *kClient) AccountSetContentSettings(ctx context.Context, req *tg.TLAccountSetContentSettings) (r *tg.Bool, err error) {
	var _args AccountSetContentSettingsArgs
	_args.Req = req
	var _result AccountSetContentSettingsResult
	if err = p.c.Call(ctx, "account.setContentSettings", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountGetContentSettings(ctx context.Context, req *tg.TLAccountGetContentSettings) (r *tg.AccountContentSettings, err error) {
	var _args AccountGetContentSettingsArgs
	_args.Req = req
	var _result AccountGetContentSettingsResult
	if err = p.c.Call(ctx, "account.getContentSettings", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
