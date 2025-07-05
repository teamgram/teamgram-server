/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package savedmessagedialogsservice

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
	"messages.getSavedDialogs": kitex.NewMethodInfo(
		messagesGetSavedDialogsHandler,
		newMessagesGetSavedDialogsArgs,
		newMessagesGetSavedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getSavedHistory": kitex.NewMethodInfo(
		messagesGetSavedHistoryHandler,
		newMessagesGetSavedHistoryArgs,
		newMessagesGetSavedHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.deleteSavedHistory": kitex.NewMethodInfo(
		messagesDeleteSavedHistoryHandler,
		newMessagesDeleteSavedHistoryArgs,
		newMessagesDeleteSavedHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getPinnedSavedDialogs": kitex.NewMethodInfo(
		messagesGetPinnedSavedDialogsHandler,
		newMessagesGetPinnedSavedDialogsArgs,
		newMessagesGetPinnedSavedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.toggleSavedDialogPin": kitex.NewMethodInfo(
		messagesToggleSavedDialogPinHandler,
		newMessagesToggleSavedDialogPinArgs,
		newMessagesToggleSavedDialogPinResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.reorderPinnedSavedDialogs": kitex.NewMethodInfo(
		messagesReorderPinnedSavedDialogsHandler,
		newMessagesReorderPinnedSavedDialogsArgs,
		newMessagesReorderPinnedSavedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	savedmessagedialogsServiceServiceInfo                = NewServiceInfo()
	savedmessagedialogsServiceServiceInfoForClient       = NewServiceInfoForClient()
	savedmessagedialogsServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCSavedMessageDialogs", savedmessagedialogsServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCSavedMessageDialogs", savedmessagedialogsServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCSavedMessageDialogs", savedmessagedialogsServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return savedmessagedialogsServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return savedmessagedialogsServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return savedmessagedialogsServiceServiceInfoForClient
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
	serviceName := "RPCSavedMessageDialogs"
	handlerType := (*tg.RPCSavedMessageDialogs)(nil)
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
		"PackageName": "savedmessagedialogs",
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

func messagesGetSavedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetSavedDialogsArgs)
	realResult := result.(*MessagesGetSavedDialogsResult)
	success, err := handler.(tg.RPCSavedMessageDialogs).MessagesGetSavedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetSavedDialogsArgs() interface{} {
	return &MessagesGetSavedDialogsArgs{}
}

func newMessagesGetSavedDialogsResult() interface{} {
	return &MessagesGetSavedDialogsResult{}
}

type MessagesGetSavedDialogsArgs struct {
	Req *tg.TLMessagesGetSavedDialogs
}

func (p *MessagesGetSavedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetSavedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetSavedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetSavedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetSavedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetSavedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetSavedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetSavedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetSavedDialogsArgs_Req_DEFAULT *tg.TLMessagesGetSavedDialogs

func (p *MessagesGetSavedDialogsArgs) GetReq() *tg.TLMessagesGetSavedDialogs {
	if !p.IsSetReq() {
		return MessagesGetSavedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetSavedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetSavedDialogsResult struct {
	Success *tg.MessagesSavedDialogs
}

var MessagesGetSavedDialogsResult_Success_DEFAULT *tg.MessagesSavedDialogs

func (p *MessagesGetSavedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetSavedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetSavedDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesSavedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSavedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetSavedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetSavedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesSavedDialogs)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSavedDialogsResult) GetSuccess() *tg.MessagesSavedDialogs {
	if !p.IsSetSuccess() {
		return MessagesGetSavedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetSavedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesSavedDialogs)
}

func (p *MessagesGetSavedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetSavedDialogsResult) GetResult() interface{} {
	return p.Success
}

func messagesGetSavedHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetSavedHistoryArgs)
	realResult := result.(*MessagesGetSavedHistoryResult)
	success, err := handler.(tg.RPCSavedMessageDialogs).MessagesGetSavedHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetSavedHistoryArgs() interface{} {
	return &MessagesGetSavedHistoryArgs{}
}

func newMessagesGetSavedHistoryResult() interface{} {
	return &MessagesGetSavedHistoryResult{}
}

type MessagesGetSavedHistoryArgs struct {
	Req *tg.TLMessagesGetSavedHistory
}

func (p *MessagesGetSavedHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetSavedHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetSavedHistoryArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetSavedHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetSavedHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetSavedHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetSavedHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetSavedHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetSavedHistoryArgs_Req_DEFAULT *tg.TLMessagesGetSavedHistory

func (p *MessagesGetSavedHistoryArgs) GetReq() *tg.TLMessagesGetSavedHistory {
	if !p.IsSetReq() {
		return MessagesGetSavedHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetSavedHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetSavedHistoryResult struct {
	Success *tg.MessagesMessages
}

var MessagesGetSavedHistoryResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesGetSavedHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetSavedHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetSavedHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSavedHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetSavedHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetSavedHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSavedHistoryResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesGetSavedHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetSavedHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesGetSavedHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetSavedHistoryResult) GetResult() interface{} {
	return p.Success
}

func messagesDeleteSavedHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesDeleteSavedHistoryArgs)
	realResult := result.(*MessagesDeleteSavedHistoryResult)
	success, err := handler.(tg.RPCSavedMessageDialogs).MessagesDeleteSavedHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesDeleteSavedHistoryArgs() interface{} {
	return &MessagesDeleteSavedHistoryArgs{}
}

func newMessagesDeleteSavedHistoryResult() interface{} {
	return &MessagesDeleteSavedHistoryResult{}
}

type MessagesDeleteSavedHistoryArgs struct {
	Req *tg.TLMessagesDeleteSavedHistory
}

func (p *MessagesDeleteSavedHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesDeleteSavedHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesDeleteSavedHistoryArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesDeleteSavedHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesDeleteSavedHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesDeleteSavedHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesDeleteSavedHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesDeleteSavedHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesDeleteSavedHistoryArgs_Req_DEFAULT *tg.TLMessagesDeleteSavedHistory

func (p *MessagesDeleteSavedHistoryArgs) GetReq() *tg.TLMessagesDeleteSavedHistory {
	if !p.IsSetReq() {
		return MessagesDeleteSavedHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesDeleteSavedHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesDeleteSavedHistoryResult struct {
	Success *tg.MessagesAffectedHistory
}

var MessagesDeleteSavedHistoryResult_Success_DEFAULT *tg.MessagesAffectedHistory

func (p *MessagesDeleteSavedHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesDeleteSavedHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesDeleteSavedHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteSavedHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesDeleteSavedHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesDeleteSavedHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedHistory)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteSavedHistoryResult) GetSuccess() *tg.MessagesAffectedHistory {
	if !p.IsSetSuccess() {
		return MessagesDeleteSavedHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesDeleteSavedHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedHistory)
}

func (p *MessagesDeleteSavedHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesDeleteSavedHistoryResult) GetResult() interface{} {
	return p.Success
}

func messagesGetPinnedSavedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetPinnedSavedDialogsArgs)
	realResult := result.(*MessagesGetPinnedSavedDialogsResult)
	success, err := handler.(tg.RPCSavedMessageDialogs).MessagesGetPinnedSavedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetPinnedSavedDialogsArgs() interface{} {
	return &MessagesGetPinnedSavedDialogsArgs{}
}

func newMessagesGetPinnedSavedDialogsResult() interface{} {
	return &MessagesGetPinnedSavedDialogsResult{}
}

type MessagesGetPinnedSavedDialogsArgs struct {
	Req *tg.TLMessagesGetPinnedSavedDialogs
}

func (p *MessagesGetPinnedSavedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetPinnedSavedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetPinnedSavedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetPinnedSavedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetPinnedSavedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetPinnedSavedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetPinnedSavedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetPinnedSavedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetPinnedSavedDialogsArgs_Req_DEFAULT *tg.TLMessagesGetPinnedSavedDialogs

func (p *MessagesGetPinnedSavedDialogsArgs) GetReq() *tg.TLMessagesGetPinnedSavedDialogs {
	if !p.IsSetReq() {
		return MessagesGetPinnedSavedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetPinnedSavedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetPinnedSavedDialogsResult struct {
	Success *tg.MessagesSavedDialogs
}

var MessagesGetPinnedSavedDialogsResult_Success_DEFAULT *tg.MessagesSavedDialogs

func (p *MessagesGetPinnedSavedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetPinnedSavedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetPinnedSavedDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesSavedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetPinnedSavedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetPinnedSavedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetPinnedSavedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesSavedDialogs)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetPinnedSavedDialogsResult) GetSuccess() *tg.MessagesSavedDialogs {
	if !p.IsSetSuccess() {
		return MessagesGetPinnedSavedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetPinnedSavedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesSavedDialogs)
}

func (p *MessagesGetPinnedSavedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetPinnedSavedDialogsResult) GetResult() interface{} {
	return p.Success
}

func messagesToggleSavedDialogPinHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesToggleSavedDialogPinArgs)
	realResult := result.(*MessagesToggleSavedDialogPinResult)
	success, err := handler.(tg.RPCSavedMessageDialogs).MessagesToggleSavedDialogPin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesToggleSavedDialogPinArgs() interface{} {
	return &MessagesToggleSavedDialogPinArgs{}
}

func newMessagesToggleSavedDialogPinResult() interface{} {
	return &MessagesToggleSavedDialogPinResult{}
}

type MessagesToggleSavedDialogPinArgs struct {
	Req *tg.TLMessagesToggleSavedDialogPin
}

func (p *MessagesToggleSavedDialogPinArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesToggleSavedDialogPinArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesToggleSavedDialogPinArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesToggleSavedDialogPin)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesToggleSavedDialogPinArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesToggleSavedDialogPinArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesToggleSavedDialogPinArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesToggleSavedDialogPin)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesToggleSavedDialogPinArgs_Req_DEFAULT *tg.TLMessagesToggleSavedDialogPin

func (p *MessagesToggleSavedDialogPinArgs) GetReq() *tg.TLMessagesToggleSavedDialogPin {
	if !p.IsSetReq() {
		return MessagesToggleSavedDialogPinArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesToggleSavedDialogPinArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesToggleSavedDialogPinResult struct {
	Success *tg.Bool
}

var MessagesToggleSavedDialogPinResult_Success_DEFAULT *tg.Bool

func (p *MessagesToggleSavedDialogPinResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesToggleSavedDialogPinResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesToggleSavedDialogPinResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesToggleSavedDialogPinResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesToggleSavedDialogPinResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesToggleSavedDialogPinResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesToggleSavedDialogPinResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesToggleSavedDialogPinResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesToggleSavedDialogPinResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesToggleSavedDialogPinResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesToggleSavedDialogPinResult) GetResult() interface{} {
	return p.Success
}

func messagesReorderPinnedSavedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReorderPinnedSavedDialogsArgs)
	realResult := result.(*MessagesReorderPinnedSavedDialogsResult)
	success, err := handler.(tg.RPCSavedMessageDialogs).MessagesReorderPinnedSavedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReorderPinnedSavedDialogsArgs() interface{} {
	return &MessagesReorderPinnedSavedDialogsArgs{}
}

func newMessagesReorderPinnedSavedDialogsResult() interface{} {
	return &MessagesReorderPinnedSavedDialogsResult{}
}

type MessagesReorderPinnedSavedDialogsArgs struct {
	Req *tg.TLMessagesReorderPinnedSavedDialogs
}

func (p *MessagesReorderPinnedSavedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReorderPinnedSavedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReorderPinnedSavedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReorderPinnedSavedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReorderPinnedSavedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReorderPinnedSavedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReorderPinnedSavedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReorderPinnedSavedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesReorderPinnedSavedDialogsArgs_Req_DEFAULT *tg.TLMessagesReorderPinnedSavedDialogs

func (p *MessagesReorderPinnedSavedDialogsArgs) GetReq() *tg.TLMessagesReorderPinnedSavedDialogs {
	if !p.IsSetReq() {
		return MessagesReorderPinnedSavedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReorderPinnedSavedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReorderPinnedSavedDialogsResult struct {
	Success *tg.Bool
}

var MessagesReorderPinnedSavedDialogsResult_Success_DEFAULT *tg.Bool

func (p *MessagesReorderPinnedSavedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReorderPinnedSavedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReorderPinnedSavedDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReorderPinnedSavedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReorderPinnedSavedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReorderPinnedSavedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReorderPinnedSavedDialogsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesReorderPinnedSavedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReorderPinnedSavedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesReorderPinnedSavedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReorderPinnedSavedDialogsResult) GetResult() interface{} {
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

func (p *kClient) MessagesGetSavedDialogs(ctx context.Context, req *tg.TLMessagesGetSavedDialogs) (r *tg.MessagesSavedDialogs, err error) {
	// var _args MessagesGetSavedDialogsArgs
	// _args.Req = req
	// var _result MessagesGetSavedDialogsResult

	_result := new(tg.MessagesSavedDialogs)
	if err = p.c.Call(ctx, "messages.getSavedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetSavedHistory(ctx context.Context, req *tg.TLMessagesGetSavedHistory) (r *tg.MessagesMessages, err error) {
	// var _args MessagesGetSavedHistoryArgs
	// _args.Req = req
	// var _result MessagesGetSavedHistoryResult

	_result := new(tg.MessagesMessages)
	if err = p.c.Call(ctx, "messages.getSavedHistory", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesDeleteSavedHistory(ctx context.Context, req *tg.TLMessagesDeleteSavedHistory) (r *tg.MessagesAffectedHistory, err error) {
	// var _args MessagesDeleteSavedHistoryArgs
	// _args.Req = req
	// var _result MessagesDeleteSavedHistoryResult

	_result := new(tg.MessagesAffectedHistory)
	if err = p.c.Call(ctx, "messages.deleteSavedHistory", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetPinnedSavedDialogs(ctx context.Context, req *tg.TLMessagesGetPinnedSavedDialogs) (r *tg.MessagesSavedDialogs, err error) {
	// var _args MessagesGetPinnedSavedDialogsArgs
	// _args.Req = req
	// var _result MessagesGetPinnedSavedDialogsResult

	_result := new(tg.MessagesSavedDialogs)
	if err = p.c.Call(ctx, "messages.getPinnedSavedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesToggleSavedDialogPin(ctx context.Context, req *tg.TLMessagesToggleSavedDialogPin) (r *tg.Bool, err error) {
	// var _args MessagesToggleSavedDialogPinArgs
	// _args.Req = req
	// var _result MessagesToggleSavedDialogPinResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "messages.toggleSavedDialogPin", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesReorderPinnedSavedDialogs(ctx context.Context, req *tg.TLMessagesReorderPinnedSavedDialogs) (r *tg.Bool, err error) {
	// var _args MessagesReorderPinnedSavedDialogsArgs
	// _args.Req = req
	// var _result MessagesReorderPinnedSavedDialogsResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "messages.reorderPinnedSavedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
