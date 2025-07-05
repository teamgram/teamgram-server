/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialogsservice

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
	"/tg.RPCDialogs/messages.getDialogs": kitex.NewMethodInfo(
		messagesGetDialogsHandler,
		newMessagesGetDialogsArgs,
		newMessagesGetDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.setTyping": kitex.NewMethodInfo(
		messagesSetTypingHandler,
		newMessagesSetTypingArgs,
		newMessagesSetTypingResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.getPeerSettings": kitex.NewMethodInfo(
		messagesGetPeerSettingsHandler,
		newMessagesGetPeerSettingsArgs,
		newMessagesGetPeerSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.getPeerDialogs": kitex.NewMethodInfo(
		messagesGetPeerDialogsHandler,
		newMessagesGetPeerDialogsArgs,
		newMessagesGetPeerDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.toggleDialogPin": kitex.NewMethodInfo(
		messagesToggleDialogPinHandler,
		newMessagesToggleDialogPinArgs,
		newMessagesToggleDialogPinResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.reorderPinnedDialogs": kitex.NewMethodInfo(
		messagesReorderPinnedDialogsHandler,
		newMessagesReorderPinnedDialogsArgs,
		newMessagesReorderPinnedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.getPinnedDialogs": kitex.NewMethodInfo(
		messagesGetPinnedDialogsHandler,
		newMessagesGetPinnedDialogsArgs,
		newMessagesGetPinnedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.sendScreenshotNotification": kitex.NewMethodInfo(
		messagesSendScreenshotNotificationHandler,
		newMessagesSendScreenshotNotificationArgs,
		newMessagesSendScreenshotNotificationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.markDialogUnread": kitex.NewMethodInfo(
		messagesMarkDialogUnreadHandler,
		newMessagesMarkDialogUnreadArgs,
		newMessagesMarkDialogUnreadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.getDialogUnreadMarks": kitex.NewMethodInfo(
		messagesGetDialogUnreadMarksHandler,
		newMessagesGetDialogUnreadMarksArgs,
		newMessagesGetDialogUnreadMarksResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.getOnlines": kitex.NewMethodInfo(
		messagesGetOnlinesHandler,
		newMessagesGetOnlinesArgs,
		newMessagesGetOnlinesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.hidePeerSettingsBar": kitex.NewMethodInfo(
		messagesHidePeerSettingsBarHandler,
		newMessagesHidePeerSettingsBarArgs,
		newMessagesHidePeerSettingsBarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCDialogs/messages.setHistoryTTL": kitex.NewMethodInfo(
		messagesSetHistoryTTLHandler,
		newMessagesSetHistoryTTLArgs,
		newMessagesSetHistoryTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	dialogsServiceServiceInfo                = NewServiceInfo()
	dialogsServiceServiceInfoForClient       = NewServiceInfoForClient()
	dialogsServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCDialogs", dialogsServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCDialogs", dialogsServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCDialogs", dialogsServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return dialogsServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return dialogsServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return dialogsServiceServiceInfoForClient
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
	serviceName := "RPCDialogs"
	handlerType := (*tg.RPCDialogs)(nil)
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
		"PackageName": "dialogs",
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

func messagesGetDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetDialogsArgs)
	realResult := result.(*MessagesGetDialogsResult)
	success, err := handler.(tg.RPCDialogs).MessagesGetDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetDialogsArgs() interface{} {
	return &MessagesGetDialogsArgs{}
}

func newMessagesGetDialogsResult() interface{} {
	return &MessagesGetDialogsResult{}
}

type MessagesGetDialogsArgs struct {
	Req *tg.TLMessagesGetDialogs
}

func (p *MessagesGetDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetDialogsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetDialogsArgs_Req_DEFAULT *tg.TLMessagesGetDialogs

func (p *MessagesGetDialogsArgs) GetReq() *tg.TLMessagesGetDialogs {
	if !p.IsSetReq() {
		return MessagesGetDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetDialogsResult struct {
	Success *tg.MessagesDialogs
}

var MessagesGetDialogsResult_Success_DEFAULT *tg.MessagesDialogs

func (p *MessagesGetDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesDialogs)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetDialogsResult) GetSuccess() *tg.MessagesDialogs {
	if !p.IsSetSuccess() {
		return MessagesGetDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesDialogs)
}

func (p *MessagesGetDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetDialogsResult) GetResult() interface{} {
	return p.Success
}

func messagesSetTypingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSetTypingArgs)
	realResult := result.(*MessagesSetTypingResult)
	success, err := handler.(tg.RPCDialogs).MessagesSetTyping(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSetTypingArgs() interface{} {
	return &MessagesSetTypingArgs{}
}

func newMessagesSetTypingResult() interface{} {
	return &MessagesSetTypingResult{}
}

type MessagesSetTypingArgs struct {
	Req *tg.TLMessagesSetTyping
}

func (p *MessagesSetTypingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSetTypingArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSetTypingArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSetTyping)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSetTypingArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSetTypingArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSetTypingArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSetTyping)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSetTypingArgs_Req_DEFAULT *tg.TLMessagesSetTyping

func (p *MessagesSetTypingArgs) GetReq() *tg.TLMessagesSetTyping {
	if !p.IsSetReq() {
		return MessagesSetTypingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSetTypingArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSetTypingResult struct {
	Success *tg.Bool
}

var MessagesSetTypingResult_Success_DEFAULT *tg.Bool

func (p *MessagesSetTypingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSetTypingResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSetTypingResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSetTypingResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSetTypingResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSetTypingResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSetTypingResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesSetTypingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSetTypingResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesSetTypingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSetTypingResult) GetResult() interface{} {
	return p.Success
}

func messagesGetPeerSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetPeerSettingsArgs)
	realResult := result.(*MessagesGetPeerSettingsResult)
	success, err := handler.(tg.RPCDialogs).MessagesGetPeerSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetPeerSettingsArgs() interface{} {
	return &MessagesGetPeerSettingsArgs{}
}

func newMessagesGetPeerSettingsResult() interface{} {
	return &MessagesGetPeerSettingsResult{}
}

type MessagesGetPeerSettingsArgs struct {
	Req *tg.TLMessagesGetPeerSettings
}

func (p *MessagesGetPeerSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetPeerSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetPeerSettingsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetPeerSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetPeerSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetPeerSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetPeerSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetPeerSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetPeerSettingsArgs_Req_DEFAULT *tg.TLMessagesGetPeerSettings

func (p *MessagesGetPeerSettingsArgs) GetReq() *tg.TLMessagesGetPeerSettings {
	if !p.IsSetReq() {
		return MessagesGetPeerSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetPeerSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetPeerSettingsResult struct {
	Success *tg.MessagesPeerSettings
}

var MessagesGetPeerSettingsResult_Success_DEFAULT *tg.MessagesPeerSettings

func (p *MessagesGetPeerSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetPeerSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetPeerSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesPeerSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetPeerSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetPeerSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetPeerSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesPeerSettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetPeerSettingsResult) GetSuccess() *tg.MessagesPeerSettings {
	if !p.IsSetSuccess() {
		return MessagesGetPeerSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetPeerSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesPeerSettings)
}

func (p *MessagesGetPeerSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetPeerSettingsResult) GetResult() interface{} {
	return p.Success
}

func messagesGetPeerDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetPeerDialogsArgs)
	realResult := result.(*MessagesGetPeerDialogsResult)
	success, err := handler.(tg.RPCDialogs).MessagesGetPeerDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetPeerDialogsArgs() interface{} {
	return &MessagesGetPeerDialogsArgs{}
}

func newMessagesGetPeerDialogsResult() interface{} {
	return &MessagesGetPeerDialogsResult{}
}

type MessagesGetPeerDialogsArgs struct {
	Req *tg.TLMessagesGetPeerDialogs
}

func (p *MessagesGetPeerDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetPeerDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetPeerDialogsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetPeerDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetPeerDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetPeerDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetPeerDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetPeerDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetPeerDialogsArgs_Req_DEFAULT *tg.TLMessagesGetPeerDialogs

func (p *MessagesGetPeerDialogsArgs) GetReq() *tg.TLMessagesGetPeerDialogs {
	if !p.IsSetReq() {
		return MessagesGetPeerDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetPeerDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetPeerDialogsResult struct {
	Success *tg.MessagesPeerDialogs
}

var MessagesGetPeerDialogsResult_Success_DEFAULT *tg.MessagesPeerDialogs

func (p *MessagesGetPeerDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetPeerDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetPeerDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesPeerDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetPeerDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetPeerDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetPeerDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesPeerDialogs)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetPeerDialogsResult) GetSuccess() *tg.MessagesPeerDialogs {
	if !p.IsSetSuccess() {
		return MessagesGetPeerDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetPeerDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesPeerDialogs)
}

func (p *MessagesGetPeerDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetPeerDialogsResult) GetResult() interface{} {
	return p.Success
}

func messagesToggleDialogPinHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesToggleDialogPinArgs)
	realResult := result.(*MessagesToggleDialogPinResult)
	success, err := handler.(tg.RPCDialogs).MessagesToggleDialogPin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesToggleDialogPinArgs() interface{} {
	return &MessagesToggleDialogPinArgs{}
}

func newMessagesToggleDialogPinResult() interface{} {
	return &MessagesToggleDialogPinResult{}
}

type MessagesToggleDialogPinArgs struct {
	Req *tg.TLMessagesToggleDialogPin
}

func (p *MessagesToggleDialogPinArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesToggleDialogPinArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesToggleDialogPinArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesToggleDialogPin)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesToggleDialogPinArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesToggleDialogPinArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesToggleDialogPinArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesToggleDialogPin)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesToggleDialogPinArgs_Req_DEFAULT *tg.TLMessagesToggleDialogPin

func (p *MessagesToggleDialogPinArgs) GetReq() *tg.TLMessagesToggleDialogPin {
	if !p.IsSetReq() {
		return MessagesToggleDialogPinArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesToggleDialogPinArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesToggleDialogPinResult struct {
	Success *tg.Bool
}

var MessagesToggleDialogPinResult_Success_DEFAULT *tg.Bool

func (p *MessagesToggleDialogPinResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesToggleDialogPinResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesToggleDialogPinResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesToggleDialogPinResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesToggleDialogPinResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesToggleDialogPinResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesToggleDialogPinResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesToggleDialogPinResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesToggleDialogPinResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesToggleDialogPinResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesToggleDialogPinResult) GetResult() interface{} {
	return p.Success
}

func messagesReorderPinnedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReorderPinnedDialogsArgs)
	realResult := result.(*MessagesReorderPinnedDialogsResult)
	success, err := handler.(tg.RPCDialogs).MessagesReorderPinnedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReorderPinnedDialogsArgs() interface{} {
	return &MessagesReorderPinnedDialogsArgs{}
}

func newMessagesReorderPinnedDialogsResult() interface{} {
	return &MessagesReorderPinnedDialogsResult{}
}

type MessagesReorderPinnedDialogsArgs struct {
	Req *tg.TLMessagesReorderPinnedDialogs
}

func (p *MessagesReorderPinnedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReorderPinnedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReorderPinnedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReorderPinnedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReorderPinnedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReorderPinnedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReorderPinnedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReorderPinnedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesReorderPinnedDialogsArgs_Req_DEFAULT *tg.TLMessagesReorderPinnedDialogs

func (p *MessagesReorderPinnedDialogsArgs) GetReq() *tg.TLMessagesReorderPinnedDialogs {
	if !p.IsSetReq() {
		return MessagesReorderPinnedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReorderPinnedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReorderPinnedDialogsResult struct {
	Success *tg.Bool
}

var MessagesReorderPinnedDialogsResult_Success_DEFAULT *tg.Bool

func (p *MessagesReorderPinnedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReorderPinnedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReorderPinnedDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReorderPinnedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReorderPinnedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReorderPinnedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReorderPinnedDialogsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesReorderPinnedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReorderPinnedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesReorderPinnedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReorderPinnedDialogsResult) GetResult() interface{} {
	return p.Success
}

func messagesGetPinnedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetPinnedDialogsArgs)
	realResult := result.(*MessagesGetPinnedDialogsResult)
	success, err := handler.(tg.RPCDialogs).MessagesGetPinnedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetPinnedDialogsArgs() interface{} {
	return &MessagesGetPinnedDialogsArgs{}
}

func newMessagesGetPinnedDialogsResult() interface{} {
	return &MessagesGetPinnedDialogsResult{}
}

type MessagesGetPinnedDialogsArgs struct {
	Req *tg.TLMessagesGetPinnedDialogs
}

func (p *MessagesGetPinnedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetPinnedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetPinnedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetPinnedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetPinnedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetPinnedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetPinnedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetPinnedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetPinnedDialogsArgs_Req_DEFAULT *tg.TLMessagesGetPinnedDialogs

func (p *MessagesGetPinnedDialogsArgs) GetReq() *tg.TLMessagesGetPinnedDialogs {
	if !p.IsSetReq() {
		return MessagesGetPinnedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetPinnedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetPinnedDialogsResult struct {
	Success *tg.MessagesPeerDialogs
}

var MessagesGetPinnedDialogsResult_Success_DEFAULT *tg.MessagesPeerDialogs

func (p *MessagesGetPinnedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetPinnedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetPinnedDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesPeerDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetPinnedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetPinnedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetPinnedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesPeerDialogs)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetPinnedDialogsResult) GetSuccess() *tg.MessagesPeerDialogs {
	if !p.IsSetSuccess() {
		return MessagesGetPinnedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetPinnedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesPeerDialogs)
}

func (p *MessagesGetPinnedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetPinnedDialogsResult) GetResult() interface{} {
	return p.Success
}

func messagesSendScreenshotNotificationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSendScreenshotNotificationArgs)
	realResult := result.(*MessagesSendScreenshotNotificationResult)
	success, err := handler.(tg.RPCDialogs).MessagesSendScreenshotNotification(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSendScreenshotNotificationArgs() interface{} {
	return &MessagesSendScreenshotNotificationArgs{}
}

func newMessagesSendScreenshotNotificationResult() interface{} {
	return &MessagesSendScreenshotNotificationResult{}
}

type MessagesSendScreenshotNotificationArgs struct {
	Req *tg.TLMessagesSendScreenshotNotification
}

func (p *MessagesSendScreenshotNotificationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSendScreenshotNotificationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSendScreenshotNotificationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSendScreenshotNotification)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSendScreenshotNotificationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSendScreenshotNotificationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSendScreenshotNotificationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSendScreenshotNotification)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSendScreenshotNotificationArgs_Req_DEFAULT *tg.TLMessagesSendScreenshotNotification

func (p *MessagesSendScreenshotNotificationArgs) GetReq() *tg.TLMessagesSendScreenshotNotification {
	if !p.IsSetReq() {
		return MessagesSendScreenshotNotificationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSendScreenshotNotificationArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSendScreenshotNotificationResult struct {
	Success *tg.Updates
}

var MessagesSendScreenshotNotificationResult_Success_DEFAULT *tg.Updates

func (p *MessagesSendScreenshotNotificationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSendScreenshotNotificationResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSendScreenshotNotificationResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSendScreenshotNotificationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSendScreenshotNotificationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSendScreenshotNotificationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSendScreenshotNotificationResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesSendScreenshotNotificationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSendScreenshotNotificationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesSendScreenshotNotificationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSendScreenshotNotificationResult) GetResult() interface{} {
	return p.Success
}

func messagesMarkDialogUnreadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesMarkDialogUnreadArgs)
	realResult := result.(*MessagesMarkDialogUnreadResult)
	success, err := handler.(tg.RPCDialogs).MessagesMarkDialogUnread(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesMarkDialogUnreadArgs() interface{} {
	return &MessagesMarkDialogUnreadArgs{}
}

func newMessagesMarkDialogUnreadResult() interface{} {
	return &MessagesMarkDialogUnreadResult{}
}

type MessagesMarkDialogUnreadArgs struct {
	Req *tg.TLMessagesMarkDialogUnread
}

func (p *MessagesMarkDialogUnreadArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesMarkDialogUnreadArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesMarkDialogUnreadArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesMarkDialogUnread)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesMarkDialogUnreadArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesMarkDialogUnreadArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesMarkDialogUnreadArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesMarkDialogUnread)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesMarkDialogUnreadArgs_Req_DEFAULT *tg.TLMessagesMarkDialogUnread

func (p *MessagesMarkDialogUnreadArgs) GetReq() *tg.TLMessagesMarkDialogUnread {
	if !p.IsSetReq() {
		return MessagesMarkDialogUnreadArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesMarkDialogUnreadArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesMarkDialogUnreadResult struct {
	Success *tg.Bool
}

var MessagesMarkDialogUnreadResult_Success_DEFAULT *tg.Bool

func (p *MessagesMarkDialogUnreadResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesMarkDialogUnreadResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesMarkDialogUnreadResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesMarkDialogUnreadResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesMarkDialogUnreadResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesMarkDialogUnreadResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesMarkDialogUnreadResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesMarkDialogUnreadResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesMarkDialogUnreadResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesMarkDialogUnreadResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesMarkDialogUnreadResult) GetResult() interface{} {
	return p.Success
}

func messagesGetDialogUnreadMarksHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetDialogUnreadMarksArgs)
	realResult := result.(*MessagesGetDialogUnreadMarksResult)
	success, err := handler.(tg.RPCDialogs).MessagesGetDialogUnreadMarks(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetDialogUnreadMarksArgs() interface{} {
	return &MessagesGetDialogUnreadMarksArgs{}
}

func newMessagesGetDialogUnreadMarksResult() interface{} {
	return &MessagesGetDialogUnreadMarksResult{}
}

type MessagesGetDialogUnreadMarksArgs struct {
	Req *tg.TLMessagesGetDialogUnreadMarks
}

func (p *MessagesGetDialogUnreadMarksArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetDialogUnreadMarksArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetDialogUnreadMarksArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetDialogUnreadMarks)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetDialogUnreadMarksArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetDialogUnreadMarksArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetDialogUnreadMarksArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetDialogUnreadMarks)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetDialogUnreadMarksArgs_Req_DEFAULT *tg.TLMessagesGetDialogUnreadMarks

func (p *MessagesGetDialogUnreadMarksArgs) GetReq() *tg.TLMessagesGetDialogUnreadMarks {
	if !p.IsSetReq() {
		return MessagesGetDialogUnreadMarksArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetDialogUnreadMarksArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetDialogUnreadMarksResult struct {
	Success *tg.VectorDialogPeer
}

var MessagesGetDialogUnreadMarksResult_Success_DEFAULT *tg.VectorDialogPeer

func (p *MessagesGetDialogUnreadMarksResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetDialogUnreadMarksResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetDialogUnreadMarksResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorDialogPeer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetDialogUnreadMarksResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetDialogUnreadMarksResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetDialogUnreadMarksResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorDialogPeer)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetDialogUnreadMarksResult) GetSuccess() *tg.VectorDialogPeer {
	if !p.IsSetSuccess() {
		return MessagesGetDialogUnreadMarksResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetDialogUnreadMarksResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorDialogPeer)
}

func (p *MessagesGetDialogUnreadMarksResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetDialogUnreadMarksResult) GetResult() interface{} {
	return p.Success
}

func messagesGetOnlinesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetOnlinesArgs)
	realResult := result.(*MessagesGetOnlinesResult)
	success, err := handler.(tg.RPCDialogs).MessagesGetOnlines(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetOnlinesArgs() interface{} {
	return &MessagesGetOnlinesArgs{}
}

func newMessagesGetOnlinesResult() interface{} {
	return &MessagesGetOnlinesResult{}
}

type MessagesGetOnlinesArgs struct {
	Req *tg.TLMessagesGetOnlines
}

func (p *MessagesGetOnlinesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetOnlinesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetOnlinesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetOnlines)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetOnlinesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetOnlinesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetOnlinesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetOnlines)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetOnlinesArgs_Req_DEFAULT *tg.TLMessagesGetOnlines

func (p *MessagesGetOnlinesArgs) GetReq() *tg.TLMessagesGetOnlines {
	if !p.IsSetReq() {
		return MessagesGetOnlinesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetOnlinesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetOnlinesResult struct {
	Success *tg.ChatOnlines
}

var MessagesGetOnlinesResult_Success_DEFAULT *tg.ChatOnlines

func (p *MessagesGetOnlinesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetOnlinesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetOnlinesResult) Unmarshal(in []byte) error {
	msg := new(tg.ChatOnlines)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetOnlinesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetOnlinesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetOnlinesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ChatOnlines)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetOnlinesResult) GetSuccess() *tg.ChatOnlines {
	if !p.IsSetSuccess() {
		return MessagesGetOnlinesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetOnlinesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ChatOnlines)
}

func (p *MessagesGetOnlinesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetOnlinesResult) GetResult() interface{} {
	return p.Success
}

func messagesHidePeerSettingsBarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesHidePeerSettingsBarArgs)
	realResult := result.(*MessagesHidePeerSettingsBarResult)
	success, err := handler.(tg.RPCDialogs).MessagesHidePeerSettingsBar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesHidePeerSettingsBarArgs() interface{} {
	return &MessagesHidePeerSettingsBarArgs{}
}

func newMessagesHidePeerSettingsBarResult() interface{} {
	return &MessagesHidePeerSettingsBarResult{}
}

type MessagesHidePeerSettingsBarArgs struct {
	Req *tg.TLMessagesHidePeerSettingsBar
}

func (p *MessagesHidePeerSettingsBarArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesHidePeerSettingsBarArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesHidePeerSettingsBarArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesHidePeerSettingsBar)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesHidePeerSettingsBarArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesHidePeerSettingsBarArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesHidePeerSettingsBarArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesHidePeerSettingsBar)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesHidePeerSettingsBarArgs_Req_DEFAULT *tg.TLMessagesHidePeerSettingsBar

func (p *MessagesHidePeerSettingsBarArgs) GetReq() *tg.TLMessagesHidePeerSettingsBar {
	if !p.IsSetReq() {
		return MessagesHidePeerSettingsBarArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesHidePeerSettingsBarArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesHidePeerSettingsBarResult struct {
	Success *tg.Bool
}

var MessagesHidePeerSettingsBarResult_Success_DEFAULT *tg.Bool

func (p *MessagesHidePeerSettingsBarResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesHidePeerSettingsBarResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesHidePeerSettingsBarResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesHidePeerSettingsBarResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesHidePeerSettingsBarResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesHidePeerSettingsBarResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesHidePeerSettingsBarResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesHidePeerSettingsBarResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesHidePeerSettingsBarResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesHidePeerSettingsBarResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesHidePeerSettingsBarResult) GetResult() interface{} {
	return p.Success
}

func messagesSetHistoryTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSetHistoryTTLArgs)
	realResult := result.(*MessagesSetHistoryTTLResult)
	success, err := handler.(tg.RPCDialogs).MessagesSetHistoryTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSetHistoryTTLArgs() interface{} {
	return &MessagesSetHistoryTTLArgs{}
}

func newMessagesSetHistoryTTLResult() interface{} {
	return &MessagesSetHistoryTTLResult{}
}

type MessagesSetHistoryTTLArgs struct {
	Req *tg.TLMessagesSetHistoryTTL
}

func (p *MessagesSetHistoryTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSetHistoryTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSetHistoryTTLArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSetHistoryTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSetHistoryTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSetHistoryTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSetHistoryTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSetHistoryTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSetHistoryTTLArgs_Req_DEFAULT *tg.TLMessagesSetHistoryTTL

func (p *MessagesSetHistoryTTLArgs) GetReq() *tg.TLMessagesSetHistoryTTL {
	if !p.IsSetReq() {
		return MessagesSetHistoryTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSetHistoryTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSetHistoryTTLResult struct {
	Success *tg.Updates
}

var MessagesSetHistoryTTLResult_Success_DEFAULT *tg.Updates

func (p *MessagesSetHistoryTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSetHistoryTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSetHistoryTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSetHistoryTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSetHistoryTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSetHistoryTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSetHistoryTTLResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesSetHistoryTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSetHistoryTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesSetHistoryTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSetHistoryTTLResult) GetResult() interface{} {
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

func (p *kClient) MessagesGetDialogs(ctx context.Context, req *tg.TLMessagesGetDialogs) (r *tg.MessagesDialogs, err error) {
	// var _args MessagesGetDialogsArgs
	// _args.Req = req
	// var _result MessagesGetDialogsResult

	_result := new(tg.MessagesDialogs)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.getDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesSetTyping(ctx context.Context, req *tg.TLMessagesSetTyping) (r *tg.Bool, err error) {
	// var _args MessagesSetTypingArgs
	// _args.Req = req
	// var _result MessagesSetTypingResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.setTyping", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetPeerSettings(ctx context.Context, req *tg.TLMessagesGetPeerSettings) (r *tg.MessagesPeerSettings, err error) {
	// var _args MessagesGetPeerSettingsArgs
	// _args.Req = req
	// var _result MessagesGetPeerSettingsResult

	_result := new(tg.MessagesPeerSettings)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.getPeerSettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetPeerDialogs(ctx context.Context, req *tg.TLMessagesGetPeerDialogs) (r *tg.MessagesPeerDialogs, err error) {
	// var _args MessagesGetPeerDialogsArgs
	// _args.Req = req
	// var _result MessagesGetPeerDialogsResult

	_result := new(tg.MessagesPeerDialogs)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.getPeerDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesToggleDialogPin(ctx context.Context, req *tg.TLMessagesToggleDialogPin) (r *tg.Bool, err error) {
	// var _args MessagesToggleDialogPinArgs
	// _args.Req = req
	// var _result MessagesToggleDialogPinResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.toggleDialogPin", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesReorderPinnedDialogs(ctx context.Context, req *tg.TLMessagesReorderPinnedDialogs) (r *tg.Bool, err error) {
	// var _args MessagesReorderPinnedDialogsArgs
	// _args.Req = req
	// var _result MessagesReorderPinnedDialogsResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.reorderPinnedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetPinnedDialogs(ctx context.Context, req *tg.TLMessagesGetPinnedDialogs) (r *tg.MessagesPeerDialogs, err error) {
	// var _args MessagesGetPinnedDialogsArgs
	// _args.Req = req
	// var _result MessagesGetPinnedDialogsResult

	_result := new(tg.MessagesPeerDialogs)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.getPinnedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesSendScreenshotNotification(ctx context.Context, req *tg.TLMessagesSendScreenshotNotification) (r *tg.Updates, err error) {
	// var _args MessagesSendScreenshotNotificationArgs
	// _args.Req = req
	// var _result MessagesSendScreenshotNotificationResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.sendScreenshotNotification", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesMarkDialogUnread(ctx context.Context, req *tg.TLMessagesMarkDialogUnread) (r *tg.Bool, err error) {
	// var _args MessagesMarkDialogUnreadArgs
	// _args.Req = req
	// var _result MessagesMarkDialogUnreadResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.markDialogUnread", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetDialogUnreadMarks(ctx context.Context, req *tg.TLMessagesGetDialogUnreadMarks) (r *tg.VectorDialogPeer, err error) {
	// var _args MessagesGetDialogUnreadMarksArgs
	// _args.Req = req
	// var _result MessagesGetDialogUnreadMarksResult

	_result := new(tg.VectorDialogPeer)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.getDialogUnreadMarks", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetOnlines(ctx context.Context, req *tg.TLMessagesGetOnlines) (r *tg.ChatOnlines, err error) {
	// var _args MessagesGetOnlinesArgs
	// _args.Req = req
	// var _result MessagesGetOnlinesResult

	_result := new(tg.ChatOnlines)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.getOnlines", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesHidePeerSettingsBar(ctx context.Context, req *tg.TLMessagesHidePeerSettingsBar) (r *tg.Bool, err error) {
	// var _args MessagesHidePeerSettingsBarArgs
	// _args.Req = req
	// var _result MessagesHidePeerSettingsBarResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.hidePeerSettingsBar", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesSetHistoryTTL(ctx context.Context, req *tg.TLMessagesSetHistoryTTL) (r *tg.Updates, err error) {
	// var _args MessagesSetHistoryTTLArgs
	// _args.Req = req
	// var _result MessagesSetHistoryTTLResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "/tg.RPCDialogs/messages.setHistoryTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
