/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package configurationservice

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
	"help.getConfig": kitex.NewMethodInfo(
		helpGetConfigHandler,
		newHelpGetConfigArgs,
		newHelpGetConfigResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.getNearestDc": kitex.NewMethodInfo(
		helpGetNearestDcHandler,
		newHelpGetNearestDcArgs,
		newHelpGetNearestDcResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.getAppUpdate": kitex.NewMethodInfo(
		helpGetAppUpdateHandler,
		newHelpGetAppUpdateArgs,
		newHelpGetAppUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.getInviteText": kitex.NewMethodInfo(
		helpGetInviteTextHandler,
		newHelpGetInviteTextArgs,
		newHelpGetInviteTextResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.getSupport": kitex.NewMethodInfo(
		helpGetSupportHandler,
		newHelpGetSupportArgs,
		newHelpGetSupportResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.getAppConfig": kitex.NewMethodInfo(
		helpGetAppConfigHandler,
		newHelpGetAppConfigArgs,
		newHelpGetAppConfigResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.getSupportName": kitex.NewMethodInfo(
		helpGetSupportNameHandler,
		newHelpGetSupportNameArgs,
		newHelpGetSupportNameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.dismissSuggestion": kitex.NewMethodInfo(
		helpDismissSuggestionHandler,
		newHelpDismissSuggestionArgs,
		newHelpDismissSuggestionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"help.getCountriesList": kitex.NewMethodInfo(
		helpGetCountriesListHandler,
		newHelpGetCountriesListArgs,
		newHelpGetCountriesListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	configurationServiceServiceInfo                = NewServiceInfo()
	configurationServiceServiceInfoForClient       = NewServiceInfoForClient()
	configurationServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCConfiguration", configurationServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCConfiguration", configurationServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCConfiguration", configurationServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return configurationServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return configurationServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return configurationServiceServiceInfoForClient
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
	serviceName := "RPCConfiguration"
	handlerType := (*tg.RPCConfiguration)(nil)
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
		"PackageName": "configuration",
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

func helpGetConfigHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetConfigArgs)
	realResult := result.(*HelpGetConfigResult)
	success, err := handler.(tg.RPCConfiguration).HelpGetConfig(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetConfigArgs() interface{} {
	return &HelpGetConfigArgs{}
}

func newHelpGetConfigResult() interface{} {
	return &HelpGetConfigResult{}
}

type HelpGetConfigArgs struct {
	Req *tg.TLHelpGetConfig
}

func (p *HelpGetConfigArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetConfigArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetConfigArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetConfig)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetConfigArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetConfigArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetConfigArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetConfig)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetConfigArgs_Req_DEFAULT *tg.TLHelpGetConfig

func (p *HelpGetConfigArgs) GetReq() *tg.TLHelpGetConfig {
	if !p.IsSetReq() {
		return HelpGetConfigArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetConfigArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetConfigResult struct {
	Success *tg.Config
}

var HelpGetConfigResult_Success_DEFAULT *tg.Config

func (p *HelpGetConfigResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetConfigResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetConfigResult) Unmarshal(in []byte) error {
	msg := new(tg.Config)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetConfigResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetConfigResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetConfigResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Config)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetConfigResult) GetSuccess() *tg.Config {
	if !p.IsSetSuccess() {
		return HelpGetConfigResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetConfigResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Config)
}

func (p *HelpGetConfigResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetConfigResult) GetResult() interface{} {
	return p.Success
}

func helpGetNearestDcHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetNearestDcArgs)
	realResult := result.(*HelpGetNearestDcResult)
	success, err := handler.(tg.RPCConfiguration).HelpGetNearestDc(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetNearestDcArgs() interface{} {
	return &HelpGetNearestDcArgs{}
}

func newHelpGetNearestDcResult() interface{} {
	return &HelpGetNearestDcResult{}
}

type HelpGetNearestDcArgs struct {
	Req *tg.TLHelpGetNearestDc
}

func (p *HelpGetNearestDcArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetNearestDcArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetNearestDcArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetNearestDc)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetNearestDcArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetNearestDcArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetNearestDcArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetNearestDc)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetNearestDcArgs_Req_DEFAULT *tg.TLHelpGetNearestDc

func (p *HelpGetNearestDcArgs) GetReq() *tg.TLHelpGetNearestDc {
	if !p.IsSetReq() {
		return HelpGetNearestDcArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetNearestDcArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetNearestDcResult struct {
	Success *tg.NearestDc
}

var HelpGetNearestDcResult_Success_DEFAULT *tg.NearestDc

func (p *HelpGetNearestDcResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetNearestDcResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetNearestDcResult) Unmarshal(in []byte) error {
	msg := new(tg.NearestDc)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetNearestDcResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetNearestDcResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetNearestDcResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.NearestDc)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetNearestDcResult) GetSuccess() *tg.NearestDc {
	if !p.IsSetSuccess() {
		return HelpGetNearestDcResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetNearestDcResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.NearestDc)
}

func (p *HelpGetNearestDcResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetNearestDcResult) GetResult() interface{} {
	return p.Success
}

func helpGetAppUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetAppUpdateArgs)
	realResult := result.(*HelpGetAppUpdateResult)
	success, err := handler.(tg.RPCConfiguration).HelpGetAppUpdate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetAppUpdateArgs() interface{} {
	return &HelpGetAppUpdateArgs{}
}

func newHelpGetAppUpdateResult() interface{} {
	return &HelpGetAppUpdateResult{}
}

type HelpGetAppUpdateArgs struct {
	Req *tg.TLHelpGetAppUpdate
}

func (p *HelpGetAppUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetAppUpdateArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetAppUpdateArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetAppUpdate)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetAppUpdateArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetAppUpdateArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetAppUpdateArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetAppUpdate)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetAppUpdateArgs_Req_DEFAULT *tg.TLHelpGetAppUpdate

func (p *HelpGetAppUpdateArgs) GetReq() *tg.TLHelpGetAppUpdate {
	if !p.IsSetReq() {
		return HelpGetAppUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetAppUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetAppUpdateResult struct {
	Success *tg.HelpAppUpdate
}

var HelpGetAppUpdateResult_Success_DEFAULT *tg.HelpAppUpdate

func (p *HelpGetAppUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetAppUpdateResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetAppUpdateResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpAppUpdate)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetAppUpdateResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetAppUpdateResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetAppUpdateResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpAppUpdate)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetAppUpdateResult) GetSuccess() *tg.HelpAppUpdate {
	if !p.IsSetSuccess() {
		return HelpGetAppUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetAppUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpAppUpdate)
}

func (p *HelpGetAppUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetAppUpdateResult) GetResult() interface{} {
	return p.Success
}

func helpGetInviteTextHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetInviteTextArgs)
	realResult := result.(*HelpGetInviteTextResult)
	success, err := handler.(tg.RPCConfiguration).HelpGetInviteText(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetInviteTextArgs() interface{} {
	return &HelpGetInviteTextArgs{}
}

func newHelpGetInviteTextResult() interface{} {
	return &HelpGetInviteTextResult{}
}

type HelpGetInviteTextArgs struct {
	Req *tg.TLHelpGetInviteText
}

func (p *HelpGetInviteTextArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetInviteTextArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetInviteTextArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetInviteText)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetInviteTextArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetInviteTextArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetInviteTextArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetInviteText)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetInviteTextArgs_Req_DEFAULT *tg.TLHelpGetInviteText

func (p *HelpGetInviteTextArgs) GetReq() *tg.TLHelpGetInviteText {
	if !p.IsSetReq() {
		return HelpGetInviteTextArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetInviteTextArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetInviteTextResult struct {
	Success *tg.HelpInviteText
}

var HelpGetInviteTextResult_Success_DEFAULT *tg.HelpInviteText

func (p *HelpGetInviteTextResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetInviteTextResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetInviteTextResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpInviteText)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetInviteTextResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetInviteTextResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetInviteTextResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpInviteText)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetInviteTextResult) GetSuccess() *tg.HelpInviteText {
	if !p.IsSetSuccess() {
		return HelpGetInviteTextResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetInviteTextResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpInviteText)
}

func (p *HelpGetInviteTextResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetInviteTextResult) GetResult() interface{} {
	return p.Success
}

func helpGetSupportHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetSupportArgs)
	realResult := result.(*HelpGetSupportResult)
	success, err := handler.(tg.RPCConfiguration).HelpGetSupport(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetSupportArgs() interface{} {
	return &HelpGetSupportArgs{}
}

func newHelpGetSupportResult() interface{} {
	return &HelpGetSupportResult{}
}

type HelpGetSupportArgs struct {
	Req *tg.TLHelpGetSupport
}

func (p *HelpGetSupportArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetSupportArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetSupportArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetSupport)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetSupportArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetSupportArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetSupportArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetSupport)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetSupportArgs_Req_DEFAULT *tg.TLHelpGetSupport

func (p *HelpGetSupportArgs) GetReq() *tg.TLHelpGetSupport {
	if !p.IsSetReq() {
		return HelpGetSupportArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetSupportArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetSupportResult struct {
	Success *tg.HelpSupport
}

var HelpGetSupportResult_Success_DEFAULT *tg.HelpSupport

func (p *HelpGetSupportResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetSupportResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetSupportResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpSupport)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetSupportResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetSupportResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetSupportResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpSupport)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetSupportResult) GetSuccess() *tg.HelpSupport {
	if !p.IsSetSuccess() {
		return HelpGetSupportResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetSupportResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpSupport)
}

func (p *HelpGetSupportResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetSupportResult) GetResult() interface{} {
	return p.Success
}

func helpGetAppConfigHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetAppConfigArgs)
	realResult := result.(*HelpGetAppConfigResult)
	success, err := handler.(tg.RPCConfiguration).HelpGetAppConfig(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetAppConfigArgs() interface{} {
	return &HelpGetAppConfigArgs{}
}

func newHelpGetAppConfigResult() interface{} {
	return &HelpGetAppConfigResult{}
}

type HelpGetAppConfigArgs struct {
	Req *tg.TLHelpGetAppConfig
}

func (p *HelpGetAppConfigArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetAppConfigArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetAppConfigArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetAppConfig)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetAppConfigArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetAppConfigArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetAppConfigArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetAppConfig)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetAppConfigArgs_Req_DEFAULT *tg.TLHelpGetAppConfig

func (p *HelpGetAppConfigArgs) GetReq() *tg.TLHelpGetAppConfig {
	if !p.IsSetReq() {
		return HelpGetAppConfigArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetAppConfigArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetAppConfigResult struct {
	Success *tg.HelpAppConfig
}

var HelpGetAppConfigResult_Success_DEFAULT *tg.HelpAppConfig

func (p *HelpGetAppConfigResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetAppConfigResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetAppConfigResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpAppConfig)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetAppConfigResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetAppConfigResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetAppConfigResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpAppConfig)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetAppConfigResult) GetSuccess() *tg.HelpAppConfig {
	if !p.IsSetSuccess() {
		return HelpGetAppConfigResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetAppConfigResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpAppConfig)
}

func (p *HelpGetAppConfigResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetAppConfigResult) GetResult() interface{} {
	return p.Success
}

func helpGetSupportNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetSupportNameArgs)
	realResult := result.(*HelpGetSupportNameResult)
	success, err := handler.(tg.RPCConfiguration).HelpGetSupportName(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetSupportNameArgs() interface{} {
	return &HelpGetSupportNameArgs{}
}

func newHelpGetSupportNameResult() interface{} {
	return &HelpGetSupportNameResult{}
}

type HelpGetSupportNameArgs struct {
	Req *tg.TLHelpGetSupportName
}

func (p *HelpGetSupportNameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetSupportNameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetSupportNameArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetSupportName)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetSupportNameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetSupportNameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetSupportNameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetSupportName)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetSupportNameArgs_Req_DEFAULT *tg.TLHelpGetSupportName

func (p *HelpGetSupportNameArgs) GetReq() *tg.TLHelpGetSupportName {
	if !p.IsSetReq() {
		return HelpGetSupportNameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetSupportNameArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetSupportNameResult struct {
	Success *tg.HelpSupportName
}

var HelpGetSupportNameResult_Success_DEFAULT *tg.HelpSupportName

func (p *HelpGetSupportNameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetSupportNameResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetSupportNameResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpSupportName)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetSupportNameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetSupportNameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetSupportNameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpSupportName)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetSupportNameResult) GetSuccess() *tg.HelpSupportName {
	if !p.IsSetSuccess() {
		return HelpGetSupportNameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetSupportNameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpSupportName)
}

func (p *HelpGetSupportNameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetSupportNameResult) GetResult() interface{} {
	return p.Success
}

func helpDismissSuggestionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpDismissSuggestionArgs)
	realResult := result.(*HelpDismissSuggestionResult)
	success, err := handler.(tg.RPCConfiguration).HelpDismissSuggestion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpDismissSuggestionArgs() interface{} {
	return &HelpDismissSuggestionArgs{}
}

func newHelpDismissSuggestionResult() interface{} {
	return &HelpDismissSuggestionResult{}
}

type HelpDismissSuggestionArgs struct {
	Req *tg.TLHelpDismissSuggestion
}

func (p *HelpDismissSuggestionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpDismissSuggestionArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpDismissSuggestionArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpDismissSuggestion)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpDismissSuggestionArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpDismissSuggestionArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpDismissSuggestionArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpDismissSuggestion)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpDismissSuggestionArgs_Req_DEFAULT *tg.TLHelpDismissSuggestion

func (p *HelpDismissSuggestionArgs) GetReq() *tg.TLHelpDismissSuggestion {
	if !p.IsSetReq() {
		return HelpDismissSuggestionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpDismissSuggestionArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpDismissSuggestionResult struct {
	Success *tg.Bool
}

var HelpDismissSuggestionResult_Success_DEFAULT *tg.Bool

func (p *HelpDismissSuggestionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpDismissSuggestionResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpDismissSuggestionResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpDismissSuggestionResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpDismissSuggestionResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpDismissSuggestionResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpDismissSuggestionResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return HelpDismissSuggestionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpDismissSuggestionResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *HelpDismissSuggestionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpDismissSuggestionResult) GetResult() interface{} {
	return p.Success
}

func helpGetCountriesListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetCountriesListArgs)
	realResult := result.(*HelpGetCountriesListResult)
	success, err := handler.(tg.RPCConfiguration).HelpGetCountriesList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetCountriesListArgs() interface{} {
	return &HelpGetCountriesListArgs{}
}

func newHelpGetCountriesListResult() interface{} {
	return &HelpGetCountriesListResult{}
}

type HelpGetCountriesListArgs struct {
	Req *tg.TLHelpGetCountriesList
}

func (p *HelpGetCountriesListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetCountriesListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetCountriesListArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetCountriesList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetCountriesListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetCountriesListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetCountriesListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetCountriesList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetCountriesListArgs_Req_DEFAULT *tg.TLHelpGetCountriesList

func (p *HelpGetCountriesListArgs) GetReq() *tg.TLHelpGetCountriesList {
	if !p.IsSetReq() {
		return HelpGetCountriesListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetCountriesListArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetCountriesListResult struct {
	Success *tg.HelpCountriesList
}

var HelpGetCountriesListResult_Success_DEFAULT *tg.HelpCountriesList

func (p *HelpGetCountriesListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetCountriesListResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetCountriesListResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpCountriesList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetCountriesListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetCountriesListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetCountriesListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpCountriesList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetCountriesListResult) GetSuccess() *tg.HelpCountriesList {
	if !p.IsSetSuccess() {
		return HelpGetCountriesListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetCountriesListResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpCountriesList)
}

func (p *HelpGetCountriesListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetCountriesListResult) GetResult() interface{} {
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

func (p *kClient) HelpGetConfig(ctx context.Context, req *tg.TLHelpGetConfig) (r *tg.Config, err error) {
	// var _args HelpGetConfigArgs
	// _args.Req = req
	// var _result HelpGetConfigResult

	_result := new(tg.Config)
	if err = p.c.Call(ctx, "help.getConfig", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpGetNearestDc(ctx context.Context, req *tg.TLHelpGetNearestDc) (r *tg.NearestDc, err error) {
	// var _args HelpGetNearestDcArgs
	// _args.Req = req
	// var _result HelpGetNearestDcResult

	_result := new(tg.NearestDc)
	if err = p.c.Call(ctx, "help.getNearestDc", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpGetAppUpdate(ctx context.Context, req *tg.TLHelpGetAppUpdate) (r *tg.HelpAppUpdate, err error) {
	// var _args HelpGetAppUpdateArgs
	// _args.Req = req
	// var _result HelpGetAppUpdateResult

	_result := new(tg.HelpAppUpdate)
	if err = p.c.Call(ctx, "help.getAppUpdate", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpGetInviteText(ctx context.Context, req *tg.TLHelpGetInviteText) (r *tg.HelpInviteText, err error) {
	// var _args HelpGetInviteTextArgs
	// _args.Req = req
	// var _result HelpGetInviteTextResult

	_result := new(tg.HelpInviteText)
	if err = p.c.Call(ctx, "help.getInviteText", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpGetSupport(ctx context.Context, req *tg.TLHelpGetSupport) (r *tg.HelpSupport, err error) {
	// var _args HelpGetSupportArgs
	// _args.Req = req
	// var _result HelpGetSupportResult

	_result := new(tg.HelpSupport)
	if err = p.c.Call(ctx, "help.getSupport", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpGetAppConfig(ctx context.Context, req *tg.TLHelpGetAppConfig) (r *tg.HelpAppConfig, err error) {
	// var _args HelpGetAppConfigArgs
	// _args.Req = req
	// var _result HelpGetAppConfigResult

	_result := new(tg.HelpAppConfig)
	if err = p.c.Call(ctx, "help.getAppConfig", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpGetSupportName(ctx context.Context, req *tg.TLHelpGetSupportName) (r *tg.HelpSupportName, err error) {
	// var _args HelpGetSupportNameArgs
	// _args.Req = req
	// var _result HelpGetSupportNameResult

	_result := new(tg.HelpSupportName)
	if err = p.c.Call(ctx, "help.getSupportName", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpDismissSuggestion(ctx context.Context, req *tg.TLHelpDismissSuggestion) (r *tg.Bool, err error) {
	// var _args HelpDismissSuggestionArgs
	// _args.Req = req
	// var _result HelpDismissSuggestionResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "help.dismissSuggestion", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpGetCountriesList(ctx context.Context, req *tg.TLHelpGetCountriesList) (r *tg.HelpCountriesList, err error) {
	// var _args HelpGetCountriesListArgs
	// _args.Req = req
	// var _result HelpGetCountriesListResult

	_result := new(tg.HelpCountriesList)
	if err = p.c.Call(ctx, "help.getCountriesList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
