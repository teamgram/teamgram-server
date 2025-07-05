/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chatinvitesservice

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
	"messages.exportChatInvite": kitex.NewMethodInfo(
		messagesExportChatInviteHandler,
		newMessagesExportChatInviteArgs,
		newMessagesExportChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.checkChatInvite": kitex.NewMethodInfo(
		messagesCheckChatInviteHandler,
		newMessagesCheckChatInviteArgs,
		newMessagesCheckChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.importChatInvite": kitex.NewMethodInfo(
		messagesImportChatInviteHandler,
		newMessagesImportChatInviteArgs,
		newMessagesImportChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getExportedChatInvites": kitex.NewMethodInfo(
		messagesGetExportedChatInvitesHandler,
		newMessagesGetExportedChatInvitesArgs,
		newMessagesGetExportedChatInvitesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getExportedChatInvite": kitex.NewMethodInfo(
		messagesGetExportedChatInviteHandler,
		newMessagesGetExportedChatInviteArgs,
		newMessagesGetExportedChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.editExportedChatInvite": kitex.NewMethodInfo(
		messagesEditExportedChatInviteHandler,
		newMessagesEditExportedChatInviteArgs,
		newMessagesEditExportedChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.deleteRevokedExportedChatInvites": kitex.NewMethodInfo(
		messagesDeleteRevokedExportedChatInvitesHandler,
		newMessagesDeleteRevokedExportedChatInvitesArgs,
		newMessagesDeleteRevokedExportedChatInvitesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.deleteExportedChatInvite": kitex.NewMethodInfo(
		messagesDeleteExportedChatInviteHandler,
		newMessagesDeleteExportedChatInviteArgs,
		newMessagesDeleteExportedChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getAdminsWithInvites": kitex.NewMethodInfo(
		messagesGetAdminsWithInvitesHandler,
		newMessagesGetAdminsWithInvitesArgs,
		newMessagesGetAdminsWithInvitesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getChatInviteImporters": kitex.NewMethodInfo(
		messagesGetChatInviteImportersHandler,
		newMessagesGetChatInviteImportersArgs,
		newMessagesGetChatInviteImportersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.hideChatJoinRequest": kitex.NewMethodInfo(
		messagesHideChatJoinRequestHandler,
		newMessagesHideChatJoinRequestArgs,
		newMessagesHideChatJoinRequestResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.hideAllChatJoinRequests": kitex.NewMethodInfo(
		messagesHideAllChatJoinRequestsHandler,
		newMessagesHideAllChatJoinRequestsArgs,
		newMessagesHideAllChatJoinRequestsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.toggleJoinToSend": kitex.NewMethodInfo(
		channelsToggleJoinToSendHandler,
		newChannelsToggleJoinToSendArgs,
		newChannelsToggleJoinToSendResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.toggleJoinRequest": kitex.NewMethodInfo(
		channelsToggleJoinRequestHandler,
		newChannelsToggleJoinRequestArgs,
		newChannelsToggleJoinRequestResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	chatinvitesServiceServiceInfo                = NewServiceInfo()
	chatinvitesServiceServiceInfoForClient       = NewServiceInfoForClient()
	chatinvitesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCChatInvites", chatinvitesServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCChatInvites", chatinvitesServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCChatInvites", chatinvitesServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return chatinvitesServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return chatinvitesServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return chatinvitesServiceServiceInfoForClient
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
	serviceName := "RPCChatInvites"
	handlerType := (*tg.RPCChatInvites)(nil)
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
		"PackageName": "chatinvites",
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

func messagesExportChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesExportChatInviteArgs)
	realResult := result.(*MessagesExportChatInviteResult)
	success, err := handler.(tg.RPCChatInvites).MessagesExportChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesExportChatInviteArgs() interface{} {
	return &MessagesExportChatInviteArgs{}
}

func newMessagesExportChatInviteResult() interface{} {
	return &MessagesExportChatInviteResult{}
}

type MessagesExportChatInviteArgs struct {
	Req *tg.TLMessagesExportChatInvite
}

func (p *MessagesExportChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesExportChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesExportChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesExportChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesExportChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesExportChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesExportChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesExportChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesExportChatInviteArgs_Req_DEFAULT *tg.TLMessagesExportChatInvite

func (p *MessagesExportChatInviteArgs) GetReq() *tg.TLMessagesExportChatInvite {
	if !p.IsSetReq() {
		return MessagesExportChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesExportChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesExportChatInviteResult struct {
	Success *tg.ExportedChatInvite
}

var MessagesExportChatInviteResult_Success_DEFAULT *tg.ExportedChatInvite

func (p *MessagesExportChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesExportChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesExportChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.ExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesExportChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesExportChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesExportChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ExportedChatInvite)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesExportChatInviteResult) GetSuccess() *tg.ExportedChatInvite {
	if !p.IsSetSuccess() {
		return MessagesExportChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesExportChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ExportedChatInvite)
}

func (p *MessagesExportChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesExportChatInviteResult) GetResult() interface{} {
	return p.Success
}

func messagesCheckChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesCheckChatInviteArgs)
	realResult := result.(*MessagesCheckChatInviteResult)
	success, err := handler.(tg.RPCChatInvites).MessagesCheckChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesCheckChatInviteArgs() interface{} {
	return &MessagesCheckChatInviteArgs{}
}

func newMessagesCheckChatInviteResult() interface{} {
	return &MessagesCheckChatInviteResult{}
}

type MessagesCheckChatInviteArgs struct {
	Req *tg.TLMessagesCheckChatInvite
}

func (p *MessagesCheckChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesCheckChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesCheckChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesCheckChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesCheckChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesCheckChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesCheckChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesCheckChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesCheckChatInviteArgs_Req_DEFAULT *tg.TLMessagesCheckChatInvite

func (p *MessagesCheckChatInviteArgs) GetReq() *tg.TLMessagesCheckChatInvite {
	if !p.IsSetReq() {
		return MessagesCheckChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesCheckChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesCheckChatInviteResult struct {
	Success *tg.ChatInvite
}

var MessagesCheckChatInviteResult_Success_DEFAULT *tg.ChatInvite

func (p *MessagesCheckChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesCheckChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesCheckChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.ChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesCheckChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesCheckChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesCheckChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ChatInvite)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesCheckChatInviteResult) GetSuccess() *tg.ChatInvite {
	if !p.IsSetSuccess() {
		return MessagesCheckChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesCheckChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ChatInvite)
}

func (p *MessagesCheckChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesCheckChatInviteResult) GetResult() interface{} {
	return p.Success
}

func messagesImportChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesImportChatInviteArgs)
	realResult := result.(*MessagesImportChatInviteResult)
	success, err := handler.(tg.RPCChatInvites).MessagesImportChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesImportChatInviteArgs() interface{} {
	return &MessagesImportChatInviteArgs{}
}

func newMessagesImportChatInviteResult() interface{} {
	return &MessagesImportChatInviteResult{}
}

type MessagesImportChatInviteArgs struct {
	Req *tg.TLMessagesImportChatInvite
}

func (p *MessagesImportChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesImportChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesImportChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesImportChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesImportChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesImportChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesImportChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesImportChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesImportChatInviteArgs_Req_DEFAULT *tg.TLMessagesImportChatInvite

func (p *MessagesImportChatInviteArgs) GetReq() *tg.TLMessagesImportChatInvite {
	if !p.IsSetReq() {
		return MessagesImportChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesImportChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesImportChatInviteResult struct {
	Success *tg.Updates
}

var MessagesImportChatInviteResult_Success_DEFAULT *tg.Updates

func (p *MessagesImportChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesImportChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesImportChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesImportChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesImportChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesImportChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesImportChatInviteResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesImportChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesImportChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesImportChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesImportChatInviteResult) GetResult() interface{} {
	return p.Success
}

func messagesGetExportedChatInvitesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetExportedChatInvitesArgs)
	realResult := result.(*MessagesGetExportedChatInvitesResult)
	success, err := handler.(tg.RPCChatInvites).MessagesGetExportedChatInvites(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetExportedChatInvitesArgs() interface{} {
	return &MessagesGetExportedChatInvitesArgs{}
}

func newMessagesGetExportedChatInvitesResult() interface{} {
	return &MessagesGetExportedChatInvitesResult{}
}

type MessagesGetExportedChatInvitesArgs struct {
	Req *tg.TLMessagesGetExportedChatInvites
}

func (p *MessagesGetExportedChatInvitesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetExportedChatInvitesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetExportedChatInvitesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetExportedChatInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetExportedChatInvitesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetExportedChatInvitesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetExportedChatInvitesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetExportedChatInvites)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetExportedChatInvitesArgs_Req_DEFAULT *tg.TLMessagesGetExportedChatInvites

func (p *MessagesGetExportedChatInvitesArgs) GetReq() *tg.TLMessagesGetExportedChatInvites {
	if !p.IsSetReq() {
		return MessagesGetExportedChatInvitesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetExportedChatInvitesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetExportedChatInvitesResult struct {
	Success *tg.MessagesExportedChatInvites
}

var MessagesGetExportedChatInvitesResult_Success_DEFAULT *tg.MessagesExportedChatInvites

func (p *MessagesGetExportedChatInvitesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetExportedChatInvitesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetExportedChatInvitesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesExportedChatInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetExportedChatInvitesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetExportedChatInvitesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetExportedChatInvitesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesExportedChatInvites)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetExportedChatInvitesResult) GetSuccess() *tg.MessagesExportedChatInvites {
	if !p.IsSetSuccess() {
		return MessagesGetExportedChatInvitesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetExportedChatInvitesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesExportedChatInvites)
}

func (p *MessagesGetExportedChatInvitesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetExportedChatInvitesResult) GetResult() interface{} {
	return p.Success
}

func messagesGetExportedChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetExportedChatInviteArgs)
	realResult := result.(*MessagesGetExportedChatInviteResult)
	success, err := handler.(tg.RPCChatInvites).MessagesGetExportedChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetExportedChatInviteArgs() interface{} {
	return &MessagesGetExportedChatInviteArgs{}
}

func newMessagesGetExportedChatInviteResult() interface{} {
	return &MessagesGetExportedChatInviteResult{}
}

type MessagesGetExportedChatInviteArgs struct {
	Req *tg.TLMessagesGetExportedChatInvite
}

func (p *MessagesGetExportedChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetExportedChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetExportedChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetExportedChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetExportedChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetExportedChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetExportedChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetExportedChatInviteArgs_Req_DEFAULT *tg.TLMessagesGetExportedChatInvite

func (p *MessagesGetExportedChatInviteArgs) GetReq() *tg.TLMessagesGetExportedChatInvite {
	if !p.IsSetReq() {
		return MessagesGetExportedChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetExportedChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetExportedChatInviteResult struct {
	Success *tg.MessagesExportedChatInvite
}

var MessagesGetExportedChatInviteResult_Success_DEFAULT *tg.MessagesExportedChatInvite

func (p *MessagesGetExportedChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetExportedChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetExportedChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetExportedChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetExportedChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetExportedChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesExportedChatInvite)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetExportedChatInviteResult) GetSuccess() *tg.MessagesExportedChatInvite {
	if !p.IsSetSuccess() {
		return MessagesGetExportedChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetExportedChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesExportedChatInvite)
}

func (p *MessagesGetExportedChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetExportedChatInviteResult) GetResult() interface{} {
	return p.Success
}

func messagesEditExportedChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesEditExportedChatInviteArgs)
	realResult := result.(*MessagesEditExportedChatInviteResult)
	success, err := handler.(tg.RPCChatInvites).MessagesEditExportedChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesEditExportedChatInviteArgs() interface{} {
	return &MessagesEditExportedChatInviteArgs{}
}

func newMessagesEditExportedChatInviteResult() interface{} {
	return &MessagesEditExportedChatInviteResult{}
}

type MessagesEditExportedChatInviteArgs struct {
	Req *tg.TLMessagesEditExportedChatInvite
}

func (p *MessagesEditExportedChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesEditExportedChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesEditExportedChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesEditExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesEditExportedChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesEditExportedChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesEditExportedChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesEditExportedChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesEditExportedChatInviteArgs_Req_DEFAULT *tg.TLMessagesEditExportedChatInvite

func (p *MessagesEditExportedChatInviteArgs) GetReq() *tg.TLMessagesEditExportedChatInvite {
	if !p.IsSetReq() {
		return MessagesEditExportedChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesEditExportedChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesEditExportedChatInviteResult struct {
	Success *tg.MessagesExportedChatInvite
}

var MessagesEditExportedChatInviteResult_Success_DEFAULT *tg.MessagesExportedChatInvite

func (p *MessagesEditExportedChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesEditExportedChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesEditExportedChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditExportedChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesEditExportedChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesEditExportedChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesExportedChatInvite)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditExportedChatInviteResult) GetSuccess() *tg.MessagesExportedChatInvite {
	if !p.IsSetSuccess() {
		return MessagesEditExportedChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesEditExportedChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesExportedChatInvite)
}

func (p *MessagesEditExportedChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesEditExportedChatInviteResult) GetResult() interface{} {
	return p.Success
}

func messagesDeleteRevokedExportedChatInvitesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesDeleteRevokedExportedChatInvitesArgs)
	realResult := result.(*MessagesDeleteRevokedExportedChatInvitesResult)
	success, err := handler.(tg.RPCChatInvites).MessagesDeleteRevokedExportedChatInvites(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesDeleteRevokedExportedChatInvitesArgs() interface{} {
	return &MessagesDeleteRevokedExportedChatInvitesArgs{}
}

func newMessagesDeleteRevokedExportedChatInvitesResult() interface{} {
	return &MessagesDeleteRevokedExportedChatInvitesResult{}
}

type MessagesDeleteRevokedExportedChatInvitesArgs struct {
	Req *tg.TLMessagesDeleteRevokedExportedChatInvites
}

func (p *MessagesDeleteRevokedExportedChatInvitesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesDeleteRevokedExportedChatInvitesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesDeleteRevokedExportedChatInvitesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesDeleteRevokedExportedChatInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesDeleteRevokedExportedChatInvitesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesDeleteRevokedExportedChatInvitesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesDeleteRevokedExportedChatInvitesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesDeleteRevokedExportedChatInvites)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesDeleteRevokedExportedChatInvitesArgs_Req_DEFAULT *tg.TLMessagesDeleteRevokedExportedChatInvites

func (p *MessagesDeleteRevokedExportedChatInvitesArgs) GetReq() *tg.TLMessagesDeleteRevokedExportedChatInvites {
	if !p.IsSetReq() {
		return MessagesDeleteRevokedExportedChatInvitesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesDeleteRevokedExportedChatInvitesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesDeleteRevokedExportedChatInvitesResult struct {
	Success *tg.Bool
}

var MessagesDeleteRevokedExportedChatInvitesResult_Success_DEFAULT *tg.Bool

func (p *MessagesDeleteRevokedExportedChatInvitesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesDeleteRevokedExportedChatInvitesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesDeleteRevokedExportedChatInvitesResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteRevokedExportedChatInvitesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesDeleteRevokedExportedChatInvitesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesDeleteRevokedExportedChatInvitesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteRevokedExportedChatInvitesResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesDeleteRevokedExportedChatInvitesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesDeleteRevokedExportedChatInvitesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesDeleteRevokedExportedChatInvitesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesDeleteRevokedExportedChatInvitesResult) GetResult() interface{} {
	return p.Success
}

func messagesDeleteExportedChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesDeleteExportedChatInviteArgs)
	realResult := result.(*MessagesDeleteExportedChatInviteResult)
	success, err := handler.(tg.RPCChatInvites).MessagesDeleteExportedChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesDeleteExportedChatInviteArgs() interface{} {
	return &MessagesDeleteExportedChatInviteArgs{}
}

func newMessagesDeleteExportedChatInviteResult() interface{} {
	return &MessagesDeleteExportedChatInviteResult{}
}

type MessagesDeleteExportedChatInviteArgs struct {
	Req *tg.TLMessagesDeleteExportedChatInvite
}

func (p *MessagesDeleteExportedChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesDeleteExportedChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesDeleteExportedChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesDeleteExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesDeleteExportedChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesDeleteExportedChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesDeleteExportedChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesDeleteExportedChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesDeleteExportedChatInviteArgs_Req_DEFAULT *tg.TLMessagesDeleteExportedChatInvite

func (p *MessagesDeleteExportedChatInviteArgs) GetReq() *tg.TLMessagesDeleteExportedChatInvite {
	if !p.IsSetReq() {
		return MessagesDeleteExportedChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesDeleteExportedChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesDeleteExportedChatInviteResult struct {
	Success *tg.Bool
}

var MessagesDeleteExportedChatInviteResult_Success_DEFAULT *tg.Bool

func (p *MessagesDeleteExportedChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesDeleteExportedChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesDeleteExportedChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteExportedChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesDeleteExportedChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesDeleteExportedChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteExportedChatInviteResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesDeleteExportedChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesDeleteExportedChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesDeleteExportedChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesDeleteExportedChatInviteResult) GetResult() interface{} {
	return p.Success
}

func messagesGetAdminsWithInvitesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetAdminsWithInvitesArgs)
	realResult := result.(*MessagesGetAdminsWithInvitesResult)
	success, err := handler.(tg.RPCChatInvites).MessagesGetAdminsWithInvites(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetAdminsWithInvitesArgs() interface{} {
	return &MessagesGetAdminsWithInvitesArgs{}
}

func newMessagesGetAdminsWithInvitesResult() interface{} {
	return &MessagesGetAdminsWithInvitesResult{}
}

type MessagesGetAdminsWithInvitesArgs struct {
	Req *tg.TLMessagesGetAdminsWithInvites
}

func (p *MessagesGetAdminsWithInvitesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetAdminsWithInvitesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetAdminsWithInvitesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetAdminsWithInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetAdminsWithInvitesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetAdminsWithInvitesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetAdminsWithInvitesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetAdminsWithInvites)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetAdminsWithInvitesArgs_Req_DEFAULT *tg.TLMessagesGetAdminsWithInvites

func (p *MessagesGetAdminsWithInvitesArgs) GetReq() *tg.TLMessagesGetAdminsWithInvites {
	if !p.IsSetReq() {
		return MessagesGetAdminsWithInvitesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetAdminsWithInvitesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetAdminsWithInvitesResult struct {
	Success *tg.MessagesChatAdminsWithInvites
}

var MessagesGetAdminsWithInvitesResult_Success_DEFAULT *tg.MessagesChatAdminsWithInvites

func (p *MessagesGetAdminsWithInvitesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetAdminsWithInvitesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetAdminsWithInvitesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesChatAdminsWithInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetAdminsWithInvitesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetAdminsWithInvitesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetAdminsWithInvitesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesChatAdminsWithInvites)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetAdminsWithInvitesResult) GetSuccess() *tg.MessagesChatAdminsWithInvites {
	if !p.IsSetSuccess() {
		return MessagesGetAdminsWithInvitesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetAdminsWithInvitesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesChatAdminsWithInvites)
}

func (p *MessagesGetAdminsWithInvitesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetAdminsWithInvitesResult) GetResult() interface{} {
	return p.Success
}

func messagesGetChatInviteImportersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetChatInviteImportersArgs)
	realResult := result.(*MessagesGetChatInviteImportersResult)
	success, err := handler.(tg.RPCChatInvites).MessagesGetChatInviteImporters(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetChatInviteImportersArgs() interface{} {
	return &MessagesGetChatInviteImportersArgs{}
}

func newMessagesGetChatInviteImportersResult() interface{} {
	return &MessagesGetChatInviteImportersResult{}
}

type MessagesGetChatInviteImportersArgs struct {
	Req *tg.TLMessagesGetChatInviteImporters
}

func (p *MessagesGetChatInviteImportersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetChatInviteImportersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetChatInviteImportersArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetChatInviteImporters)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetChatInviteImportersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetChatInviteImportersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetChatInviteImportersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetChatInviteImporters)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetChatInviteImportersArgs_Req_DEFAULT *tg.TLMessagesGetChatInviteImporters

func (p *MessagesGetChatInviteImportersArgs) GetReq() *tg.TLMessagesGetChatInviteImporters {
	if !p.IsSetReq() {
		return MessagesGetChatInviteImportersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetChatInviteImportersArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetChatInviteImportersResult struct {
	Success *tg.MessagesChatInviteImporters
}

var MessagesGetChatInviteImportersResult_Success_DEFAULT *tg.MessagesChatInviteImporters

func (p *MessagesGetChatInviteImportersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetChatInviteImportersResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetChatInviteImportersResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesChatInviteImporters)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetChatInviteImportersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetChatInviteImportersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetChatInviteImportersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesChatInviteImporters)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetChatInviteImportersResult) GetSuccess() *tg.MessagesChatInviteImporters {
	if !p.IsSetSuccess() {
		return MessagesGetChatInviteImportersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetChatInviteImportersResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesChatInviteImporters)
}

func (p *MessagesGetChatInviteImportersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetChatInviteImportersResult) GetResult() interface{} {
	return p.Success
}

func messagesHideChatJoinRequestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesHideChatJoinRequestArgs)
	realResult := result.(*MessagesHideChatJoinRequestResult)
	success, err := handler.(tg.RPCChatInvites).MessagesHideChatJoinRequest(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesHideChatJoinRequestArgs() interface{} {
	return &MessagesHideChatJoinRequestArgs{}
}

func newMessagesHideChatJoinRequestResult() interface{} {
	return &MessagesHideChatJoinRequestResult{}
}

type MessagesHideChatJoinRequestArgs struct {
	Req *tg.TLMessagesHideChatJoinRequest
}

func (p *MessagesHideChatJoinRequestArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesHideChatJoinRequestArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesHideChatJoinRequestArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesHideChatJoinRequest)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesHideChatJoinRequestArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesHideChatJoinRequestArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesHideChatJoinRequestArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesHideChatJoinRequest)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesHideChatJoinRequestArgs_Req_DEFAULT *tg.TLMessagesHideChatJoinRequest

func (p *MessagesHideChatJoinRequestArgs) GetReq() *tg.TLMessagesHideChatJoinRequest {
	if !p.IsSetReq() {
		return MessagesHideChatJoinRequestArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesHideChatJoinRequestArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesHideChatJoinRequestResult struct {
	Success *tg.Updates
}

var MessagesHideChatJoinRequestResult_Success_DEFAULT *tg.Updates

func (p *MessagesHideChatJoinRequestResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesHideChatJoinRequestResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesHideChatJoinRequestResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesHideChatJoinRequestResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesHideChatJoinRequestResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesHideChatJoinRequestResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesHideChatJoinRequestResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesHideChatJoinRequestResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesHideChatJoinRequestResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesHideChatJoinRequestResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesHideChatJoinRequestResult) GetResult() interface{} {
	return p.Success
}

func messagesHideAllChatJoinRequestsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesHideAllChatJoinRequestsArgs)
	realResult := result.(*MessagesHideAllChatJoinRequestsResult)
	success, err := handler.(tg.RPCChatInvites).MessagesHideAllChatJoinRequests(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesHideAllChatJoinRequestsArgs() interface{} {
	return &MessagesHideAllChatJoinRequestsArgs{}
}

func newMessagesHideAllChatJoinRequestsResult() interface{} {
	return &MessagesHideAllChatJoinRequestsResult{}
}

type MessagesHideAllChatJoinRequestsArgs struct {
	Req *tg.TLMessagesHideAllChatJoinRequests
}

func (p *MessagesHideAllChatJoinRequestsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesHideAllChatJoinRequestsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesHideAllChatJoinRequestsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesHideAllChatJoinRequests)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesHideAllChatJoinRequestsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesHideAllChatJoinRequestsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesHideAllChatJoinRequestsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesHideAllChatJoinRequests)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesHideAllChatJoinRequestsArgs_Req_DEFAULT *tg.TLMessagesHideAllChatJoinRequests

func (p *MessagesHideAllChatJoinRequestsArgs) GetReq() *tg.TLMessagesHideAllChatJoinRequests {
	if !p.IsSetReq() {
		return MessagesHideAllChatJoinRequestsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesHideAllChatJoinRequestsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesHideAllChatJoinRequestsResult struct {
	Success *tg.Updates
}

var MessagesHideAllChatJoinRequestsResult_Success_DEFAULT *tg.Updates

func (p *MessagesHideAllChatJoinRequestsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesHideAllChatJoinRequestsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesHideAllChatJoinRequestsResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesHideAllChatJoinRequestsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesHideAllChatJoinRequestsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesHideAllChatJoinRequestsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesHideAllChatJoinRequestsResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesHideAllChatJoinRequestsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesHideAllChatJoinRequestsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesHideAllChatJoinRequestsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesHideAllChatJoinRequestsResult) GetResult() interface{} {
	return p.Success
}

func channelsToggleJoinToSendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsToggleJoinToSendArgs)
	realResult := result.(*ChannelsToggleJoinToSendResult)
	success, err := handler.(tg.RPCChatInvites).ChannelsToggleJoinToSend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsToggleJoinToSendArgs() interface{} {
	return &ChannelsToggleJoinToSendArgs{}
}

func newChannelsToggleJoinToSendResult() interface{} {
	return &ChannelsToggleJoinToSendResult{}
}

type ChannelsToggleJoinToSendArgs struct {
	Req *tg.TLChannelsToggleJoinToSend
}

func (p *ChannelsToggleJoinToSendArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsToggleJoinToSendArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsToggleJoinToSendArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsToggleJoinToSend)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsToggleJoinToSendArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsToggleJoinToSendArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsToggleJoinToSendArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsToggleJoinToSend)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsToggleJoinToSendArgs_Req_DEFAULT *tg.TLChannelsToggleJoinToSend

func (p *ChannelsToggleJoinToSendArgs) GetReq() *tg.TLChannelsToggleJoinToSend {
	if !p.IsSetReq() {
		return ChannelsToggleJoinToSendArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsToggleJoinToSendArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsToggleJoinToSendResult struct {
	Success *tg.Updates
}

var ChannelsToggleJoinToSendResult_Success_DEFAULT *tg.Updates

func (p *ChannelsToggleJoinToSendResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsToggleJoinToSendResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsToggleJoinToSendResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsToggleJoinToSendResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsToggleJoinToSendResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsToggleJoinToSendResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsToggleJoinToSendResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return ChannelsToggleJoinToSendResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsToggleJoinToSendResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *ChannelsToggleJoinToSendResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsToggleJoinToSendResult) GetResult() interface{} {
	return p.Success
}

func channelsToggleJoinRequestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsToggleJoinRequestArgs)
	realResult := result.(*ChannelsToggleJoinRequestResult)
	success, err := handler.(tg.RPCChatInvites).ChannelsToggleJoinRequest(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsToggleJoinRequestArgs() interface{} {
	return &ChannelsToggleJoinRequestArgs{}
}

func newChannelsToggleJoinRequestResult() interface{} {
	return &ChannelsToggleJoinRequestResult{}
}

type ChannelsToggleJoinRequestArgs struct {
	Req *tg.TLChannelsToggleJoinRequest
}

func (p *ChannelsToggleJoinRequestArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsToggleJoinRequestArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsToggleJoinRequestArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsToggleJoinRequest)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsToggleJoinRequestArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsToggleJoinRequestArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsToggleJoinRequestArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsToggleJoinRequest)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsToggleJoinRequestArgs_Req_DEFAULT *tg.TLChannelsToggleJoinRequest

func (p *ChannelsToggleJoinRequestArgs) GetReq() *tg.TLChannelsToggleJoinRequest {
	if !p.IsSetReq() {
		return ChannelsToggleJoinRequestArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsToggleJoinRequestArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsToggleJoinRequestResult struct {
	Success *tg.Updates
}

var ChannelsToggleJoinRequestResult_Success_DEFAULT *tg.Updates

func (p *ChannelsToggleJoinRequestResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsToggleJoinRequestResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsToggleJoinRequestResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsToggleJoinRequestResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsToggleJoinRequestResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsToggleJoinRequestResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsToggleJoinRequestResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return ChannelsToggleJoinRequestResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsToggleJoinRequestResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *ChannelsToggleJoinRequestResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsToggleJoinRequestResult) GetResult() interface{} {
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

func (p *kClient) MessagesExportChatInvite(ctx context.Context, req *tg.TLMessagesExportChatInvite) (r *tg.ExportedChatInvite, err error) {
	// var _args MessagesExportChatInviteArgs
	// _args.Req = req
	// var _result MessagesExportChatInviteResult

	_result := new(tg.ExportedChatInvite)
	if err = p.c.Call(ctx, "messages.exportChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesCheckChatInvite(ctx context.Context, req *tg.TLMessagesCheckChatInvite) (r *tg.ChatInvite, err error) {
	// var _args MessagesCheckChatInviteArgs
	// _args.Req = req
	// var _result MessagesCheckChatInviteResult

	_result := new(tg.ChatInvite)
	if err = p.c.Call(ctx, "messages.checkChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesImportChatInvite(ctx context.Context, req *tg.TLMessagesImportChatInvite) (r *tg.Updates, err error) {
	// var _args MessagesImportChatInviteArgs
	// _args.Req = req
	// var _result MessagesImportChatInviteResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "messages.importChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetExportedChatInvites(ctx context.Context, req *tg.TLMessagesGetExportedChatInvites) (r *tg.MessagesExportedChatInvites, err error) {
	// var _args MessagesGetExportedChatInvitesArgs
	// _args.Req = req
	// var _result MessagesGetExportedChatInvitesResult

	_result := new(tg.MessagesExportedChatInvites)
	if err = p.c.Call(ctx, "messages.getExportedChatInvites", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetExportedChatInvite(ctx context.Context, req *tg.TLMessagesGetExportedChatInvite) (r *tg.MessagesExportedChatInvite, err error) {
	// var _args MessagesGetExportedChatInviteArgs
	// _args.Req = req
	// var _result MessagesGetExportedChatInviteResult

	_result := new(tg.MessagesExportedChatInvite)
	if err = p.c.Call(ctx, "messages.getExportedChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesEditExportedChatInvite(ctx context.Context, req *tg.TLMessagesEditExportedChatInvite) (r *tg.MessagesExportedChatInvite, err error) {
	// var _args MessagesEditExportedChatInviteArgs
	// _args.Req = req
	// var _result MessagesEditExportedChatInviteResult

	_result := new(tg.MessagesExportedChatInvite)
	if err = p.c.Call(ctx, "messages.editExportedChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesDeleteRevokedExportedChatInvites(ctx context.Context, req *tg.TLMessagesDeleteRevokedExportedChatInvites) (r *tg.Bool, err error) {
	// var _args MessagesDeleteRevokedExportedChatInvitesArgs
	// _args.Req = req
	// var _result MessagesDeleteRevokedExportedChatInvitesResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "messages.deleteRevokedExportedChatInvites", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesDeleteExportedChatInvite(ctx context.Context, req *tg.TLMessagesDeleteExportedChatInvite) (r *tg.Bool, err error) {
	// var _args MessagesDeleteExportedChatInviteArgs
	// _args.Req = req
	// var _result MessagesDeleteExportedChatInviteResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "messages.deleteExportedChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetAdminsWithInvites(ctx context.Context, req *tg.TLMessagesGetAdminsWithInvites) (r *tg.MessagesChatAdminsWithInvites, err error) {
	// var _args MessagesGetAdminsWithInvitesArgs
	// _args.Req = req
	// var _result MessagesGetAdminsWithInvitesResult

	_result := new(tg.MessagesChatAdminsWithInvites)
	if err = p.c.Call(ctx, "messages.getAdminsWithInvites", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetChatInviteImporters(ctx context.Context, req *tg.TLMessagesGetChatInviteImporters) (r *tg.MessagesChatInviteImporters, err error) {
	// var _args MessagesGetChatInviteImportersArgs
	// _args.Req = req
	// var _result MessagesGetChatInviteImportersResult

	_result := new(tg.MessagesChatInviteImporters)
	if err = p.c.Call(ctx, "messages.getChatInviteImporters", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesHideChatJoinRequest(ctx context.Context, req *tg.TLMessagesHideChatJoinRequest) (r *tg.Updates, err error) {
	// var _args MessagesHideChatJoinRequestArgs
	// _args.Req = req
	// var _result MessagesHideChatJoinRequestResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "messages.hideChatJoinRequest", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesHideAllChatJoinRequests(ctx context.Context, req *tg.TLMessagesHideAllChatJoinRequests) (r *tg.Updates, err error) {
	// var _args MessagesHideAllChatJoinRequestsArgs
	// _args.Req = req
	// var _result MessagesHideAllChatJoinRequestsResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "messages.hideAllChatJoinRequests", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChannelsToggleJoinToSend(ctx context.Context, req *tg.TLChannelsToggleJoinToSend) (r *tg.Updates, err error) {
	// var _args ChannelsToggleJoinToSendArgs
	// _args.Req = req
	// var _result ChannelsToggleJoinToSendResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "channels.toggleJoinToSend", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChannelsToggleJoinRequest(ctx context.Context, req *tg.TLChannelsToggleJoinRequest) (r *tg.Updates, err error) {
	// var _args ChannelsToggleJoinRequestArgs
	// _args.Req = req
	// var _result ChannelsToggleJoinRequestResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "channels.toggleJoinRequest", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
