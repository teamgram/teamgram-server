/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package inboxservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/inbox"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/inbox.RPCInbox/inbox.editUserMessageToInbox": kitex.NewMethodInfo(
		editUserMessageToInboxHandler,
		newEditUserMessageToInboxArgs,
		newEditUserMessageToInboxResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.editChatMessageToInbox": kitex.NewMethodInfo(
		editChatMessageToInboxHandler,
		newEditChatMessageToInboxArgs,
		newEditChatMessageToInboxResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.deleteMessagesToInbox": kitex.NewMethodInfo(
		deleteMessagesToInboxHandler,
		newDeleteMessagesToInboxArgs,
		newDeleteMessagesToInboxResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.deleteUserHistoryToInbox": kitex.NewMethodInfo(
		deleteUserHistoryToInboxHandler,
		newDeleteUserHistoryToInboxArgs,
		newDeleteUserHistoryToInboxResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.deleteChatHistoryToInbox": kitex.NewMethodInfo(
		deleteChatHistoryToInboxHandler,
		newDeleteChatHistoryToInboxArgs,
		newDeleteChatHistoryToInboxResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.readUserMediaUnreadToInbox": kitex.NewMethodInfo(
		readUserMediaUnreadToInboxHandler,
		newReadUserMediaUnreadToInboxArgs,
		newReadUserMediaUnreadToInboxResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.readChatMediaUnreadToInbox": kitex.NewMethodInfo(
		readChatMediaUnreadToInboxHandler,
		newReadChatMediaUnreadToInboxArgs,
		newReadChatMediaUnreadToInboxResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.updateHistoryReaded": kitex.NewMethodInfo(
		updateHistoryReadedHandler,
		newUpdateHistoryReadedArgs,
		newUpdateHistoryReadedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.updatePinnedMessage": kitex.NewMethodInfo(
		updatePinnedMessageHandler,
		newUpdatePinnedMessageArgs,
		newUpdatePinnedMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.unpinAllMessages": kitex.NewMethodInfo(
		unpinAllMessagesHandler,
		newUnpinAllMessagesArgs,
		newUnpinAllMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.sendUserMessageToInboxV2": kitex.NewMethodInfo(
		sendUserMessageToInboxV2Handler,
		newSendUserMessageToInboxV2Args,
		newSendUserMessageToInboxV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.editMessageToInboxV2": kitex.NewMethodInfo(
		editMessageToInboxV2Handler,
		newEditMessageToInboxV2Args,
		newEditMessageToInboxV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.readInboxHistory": kitex.NewMethodInfo(
		readInboxHistoryHandler,
		newReadInboxHistoryArgs,
		newReadInboxHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.readOutboxHistory": kitex.NewMethodInfo(
		readOutboxHistoryHandler,
		newReadOutboxHistoryArgs,
		newReadOutboxHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.readMediaUnreadToInboxV2": kitex.NewMethodInfo(
		readMediaUnreadToInboxV2Handler,
		newReadMediaUnreadToInboxV2Args,
		newReadMediaUnreadToInboxV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/inbox.RPCInbox/inbox.updatePinnedMessageV2": kitex.NewMethodInfo(
		updatePinnedMessageV2Handler,
		newUpdatePinnedMessageV2Args,
		newUpdatePinnedMessageV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	inboxServiceServiceInfo                = NewServiceInfo()
	inboxServiceServiceInfoForClient       = NewServiceInfoForClient()
	inboxServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCInbox", inboxServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCInbox", inboxServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCInbox", inboxServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return inboxServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return inboxServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return inboxServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfoForClient creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}

// NewServiceInfoForStreamClient creates a new ServiceInfo containing streaming methods
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RPCInbox"
	handlerType := (*inbox.RPCInbox)(nil)
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
		"PackageName": "inbox",
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

func editUserMessageToInboxHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditUserMessageToInboxArgs)
	realResult := result.(*EditUserMessageToInboxResult)
	success, err := handler.(inbox.RPCInbox).InboxEditUserMessageToInbox(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditUserMessageToInboxArgs() interface{} {
	return &EditUserMessageToInboxArgs{}
}

func newEditUserMessageToInboxResult() interface{} {
	return &EditUserMessageToInboxResult{}
}

type EditUserMessageToInboxArgs struct {
	Req *inbox.TLInboxEditUserMessageToInbox
}

func (p *EditUserMessageToInboxArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditUserMessageToInboxArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditUserMessageToInboxArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxEditUserMessageToInbox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditUserMessageToInboxArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditUserMessageToInboxArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditUserMessageToInboxArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxEditUserMessageToInbox)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditUserMessageToInboxArgs_Req_DEFAULT *inbox.TLInboxEditUserMessageToInbox

func (p *EditUserMessageToInboxArgs) GetReq() *inbox.TLInboxEditUserMessageToInbox {
	if !p.IsSetReq() {
		return EditUserMessageToInboxArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditUserMessageToInboxArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditUserMessageToInboxResult struct {
	Success *tg.Void
}

var EditUserMessageToInboxResult_Success_DEFAULT *tg.Void

func (p *EditUserMessageToInboxResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditUserMessageToInboxResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditUserMessageToInboxResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditUserMessageToInboxResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditUserMessageToInboxResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditUserMessageToInboxResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditUserMessageToInboxResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return EditUserMessageToInboxResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditUserMessageToInboxResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *EditUserMessageToInboxResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditUserMessageToInboxResult) GetResult() interface{} {
	return p.Success
}

func editChatMessageToInboxHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditChatMessageToInboxArgs)
	realResult := result.(*EditChatMessageToInboxResult)
	success, err := handler.(inbox.RPCInbox).InboxEditChatMessageToInbox(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditChatMessageToInboxArgs() interface{} {
	return &EditChatMessageToInboxArgs{}
}

func newEditChatMessageToInboxResult() interface{} {
	return &EditChatMessageToInboxResult{}
}

type EditChatMessageToInboxArgs struct {
	Req *inbox.TLInboxEditChatMessageToInbox
}

func (p *EditChatMessageToInboxArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditChatMessageToInboxArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditChatMessageToInboxArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxEditChatMessageToInbox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditChatMessageToInboxArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditChatMessageToInboxArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditChatMessageToInboxArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxEditChatMessageToInbox)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditChatMessageToInboxArgs_Req_DEFAULT *inbox.TLInboxEditChatMessageToInbox

func (p *EditChatMessageToInboxArgs) GetReq() *inbox.TLInboxEditChatMessageToInbox {
	if !p.IsSetReq() {
		return EditChatMessageToInboxArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditChatMessageToInboxArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditChatMessageToInboxResult struct {
	Success *tg.Void
}

var EditChatMessageToInboxResult_Success_DEFAULT *tg.Void

func (p *EditChatMessageToInboxResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditChatMessageToInboxResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditChatMessageToInboxResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatMessageToInboxResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditChatMessageToInboxResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditChatMessageToInboxResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatMessageToInboxResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return EditChatMessageToInboxResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditChatMessageToInboxResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *EditChatMessageToInboxResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditChatMessageToInboxResult) GetResult() interface{} {
	return p.Success
}

func deleteMessagesToInboxHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteMessagesToInboxArgs)
	realResult := result.(*DeleteMessagesToInboxResult)
	success, err := handler.(inbox.RPCInbox).InboxDeleteMessagesToInbox(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteMessagesToInboxArgs() interface{} {
	return &DeleteMessagesToInboxArgs{}
}

func newDeleteMessagesToInboxResult() interface{} {
	return &DeleteMessagesToInboxResult{}
}

type DeleteMessagesToInboxArgs struct {
	Req *inbox.TLInboxDeleteMessagesToInbox
}

func (p *DeleteMessagesToInboxArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteMessagesToInboxArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteMessagesToInboxArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxDeleteMessagesToInbox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteMessagesToInboxArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteMessagesToInboxArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteMessagesToInboxArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxDeleteMessagesToInbox)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteMessagesToInboxArgs_Req_DEFAULT *inbox.TLInboxDeleteMessagesToInbox

func (p *DeleteMessagesToInboxArgs) GetReq() *inbox.TLInboxDeleteMessagesToInbox {
	if !p.IsSetReq() {
		return DeleteMessagesToInboxArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteMessagesToInboxArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteMessagesToInboxResult struct {
	Success *tg.Void
}

var DeleteMessagesToInboxResult_Success_DEFAULT *tg.Void

func (p *DeleteMessagesToInboxResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteMessagesToInboxResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteMessagesToInboxResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteMessagesToInboxResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteMessagesToInboxResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteMessagesToInboxResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteMessagesToInboxResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return DeleteMessagesToInboxResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteMessagesToInboxResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *DeleteMessagesToInboxResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteMessagesToInboxResult) GetResult() interface{} {
	return p.Success
}

func deleteUserHistoryToInboxHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteUserHistoryToInboxArgs)
	realResult := result.(*DeleteUserHistoryToInboxResult)
	success, err := handler.(inbox.RPCInbox).InboxDeleteUserHistoryToInbox(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteUserHistoryToInboxArgs() interface{} {
	return &DeleteUserHistoryToInboxArgs{}
}

func newDeleteUserHistoryToInboxResult() interface{} {
	return &DeleteUserHistoryToInboxResult{}
}

type DeleteUserHistoryToInboxArgs struct {
	Req *inbox.TLInboxDeleteUserHistoryToInbox
}

func (p *DeleteUserHistoryToInboxArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteUserHistoryToInboxArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteUserHistoryToInboxArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxDeleteUserHistoryToInbox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteUserHistoryToInboxArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteUserHistoryToInboxArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteUserHistoryToInboxArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxDeleteUserHistoryToInbox)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteUserHistoryToInboxArgs_Req_DEFAULT *inbox.TLInboxDeleteUserHistoryToInbox

func (p *DeleteUserHistoryToInboxArgs) GetReq() *inbox.TLInboxDeleteUserHistoryToInbox {
	if !p.IsSetReq() {
		return DeleteUserHistoryToInboxArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteUserHistoryToInboxArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteUserHistoryToInboxResult struct {
	Success *tg.Void
}

var DeleteUserHistoryToInboxResult_Success_DEFAULT *tg.Void

func (p *DeleteUserHistoryToInboxResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteUserHistoryToInboxResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteUserHistoryToInboxResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUserHistoryToInboxResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteUserHistoryToInboxResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteUserHistoryToInboxResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUserHistoryToInboxResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return DeleteUserHistoryToInboxResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteUserHistoryToInboxResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *DeleteUserHistoryToInboxResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteUserHistoryToInboxResult) GetResult() interface{} {
	return p.Success
}

func deleteChatHistoryToInboxHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteChatHistoryToInboxArgs)
	realResult := result.(*DeleteChatHistoryToInboxResult)
	success, err := handler.(inbox.RPCInbox).InboxDeleteChatHistoryToInbox(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteChatHistoryToInboxArgs() interface{} {
	return &DeleteChatHistoryToInboxArgs{}
}

func newDeleteChatHistoryToInboxResult() interface{} {
	return &DeleteChatHistoryToInboxResult{}
}

type DeleteChatHistoryToInboxArgs struct {
	Req *inbox.TLInboxDeleteChatHistoryToInbox
}

func (p *DeleteChatHistoryToInboxArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteChatHistoryToInboxArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteChatHistoryToInboxArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxDeleteChatHistoryToInbox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteChatHistoryToInboxArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteChatHistoryToInboxArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteChatHistoryToInboxArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxDeleteChatHistoryToInbox)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteChatHistoryToInboxArgs_Req_DEFAULT *inbox.TLInboxDeleteChatHistoryToInbox

func (p *DeleteChatHistoryToInboxArgs) GetReq() *inbox.TLInboxDeleteChatHistoryToInbox {
	if !p.IsSetReq() {
		return DeleteChatHistoryToInboxArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteChatHistoryToInboxArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteChatHistoryToInboxResult struct {
	Success *tg.Void
}

var DeleteChatHistoryToInboxResult_Success_DEFAULT *tg.Void

func (p *DeleteChatHistoryToInboxResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteChatHistoryToInboxResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteChatHistoryToInboxResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteChatHistoryToInboxResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteChatHistoryToInboxResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteChatHistoryToInboxResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteChatHistoryToInboxResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return DeleteChatHistoryToInboxResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteChatHistoryToInboxResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *DeleteChatHistoryToInboxResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteChatHistoryToInboxResult) GetResult() interface{} {
	return p.Success
}

func readUserMediaUnreadToInboxHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReadUserMediaUnreadToInboxArgs)
	realResult := result.(*ReadUserMediaUnreadToInboxResult)
	success, err := handler.(inbox.RPCInbox).InboxReadUserMediaUnreadToInbox(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReadUserMediaUnreadToInboxArgs() interface{} {
	return &ReadUserMediaUnreadToInboxArgs{}
}

func newReadUserMediaUnreadToInboxResult() interface{} {
	return &ReadUserMediaUnreadToInboxResult{}
}

type ReadUserMediaUnreadToInboxArgs struct {
	Req *inbox.TLInboxReadUserMediaUnreadToInbox
}

func (p *ReadUserMediaUnreadToInboxArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReadUserMediaUnreadToInboxArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ReadUserMediaUnreadToInboxArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxReadUserMediaUnreadToInbox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReadUserMediaUnreadToInboxArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReadUserMediaUnreadToInboxArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReadUserMediaUnreadToInboxArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxReadUserMediaUnreadToInbox)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReadUserMediaUnreadToInboxArgs_Req_DEFAULT *inbox.TLInboxReadUserMediaUnreadToInbox

func (p *ReadUserMediaUnreadToInboxArgs) GetReq() *inbox.TLInboxReadUserMediaUnreadToInbox {
	if !p.IsSetReq() {
		return ReadUserMediaUnreadToInboxArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReadUserMediaUnreadToInboxArgs) IsSetReq() bool {
	return p.Req != nil
}

type ReadUserMediaUnreadToInboxResult struct {
	Success *tg.Void
}

var ReadUserMediaUnreadToInboxResult_Success_DEFAULT *tg.Void

func (p *ReadUserMediaUnreadToInboxResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReadUserMediaUnreadToInboxResult")
	}
	return json.Marshal(p.Success)
}

func (p *ReadUserMediaUnreadToInboxResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadUserMediaUnreadToInboxResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReadUserMediaUnreadToInboxResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReadUserMediaUnreadToInboxResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadUserMediaUnreadToInboxResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return ReadUserMediaUnreadToInboxResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReadUserMediaUnreadToInboxResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *ReadUserMediaUnreadToInboxResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReadUserMediaUnreadToInboxResult) GetResult() interface{} {
	return p.Success
}

func readChatMediaUnreadToInboxHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReadChatMediaUnreadToInboxArgs)
	realResult := result.(*ReadChatMediaUnreadToInboxResult)
	success, err := handler.(inbox.RPCInbox).InboxReadChatMediaUnreadToInbox(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReadChatMediaUnreadToInboxArgs() interface{} {
	return &ReadChatMediaUnreadToInboxArgs{}
}

func newReadChatMediaUnreadToInboxResult() interface{} {
	return &ReadChatMediaUnreadToInboxResult{}
}

type ReadChatMediaUnreadToInboxArgs struct {
	Req *inbox.TLInboxReadChatMediaUnreadToInbox
}

func (p *ReadChatMediaUnreadToInboxArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReadChatMediaUnreadToInboxArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ReadChatMediaUnreadToInboxArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxReadChatMediaUnreadToInbox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReadChatMediaUnreadToInboxArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReadChatMediaUnreadToInboxArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReadChatMediaUnreadToInboxArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxReadChatMediaUnreadToInbox)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReadChatMediaUnreadToInboxArgs_Req_DEFAULT *inbox.TLInboxReadChatMediaUnreadToInbox

func (p *ReadChatMediaUnreadToInboxArgs) GetReq() *inbox.TLInboxReadChatMediaUnreadToInbox {
	if !p.IsSetReq() {
		return ReadChatMediaUnreadToInboxArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReadChatMediaUnreadToInboxArgs) IsSetReq() bool {
	return p.Req != nil
}

type ReadChatMediaUnreadToInboxResult struct {
	Success *tg.Void
}

var ReadChatMediaUnreadToInboxResult_Success_DEFAULT *tg.Void

func (p *ReadChatMediaUnreadToInboxResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReadChatMediaUnreadToInboxResult")
	}
	return json.Marshal(p.Success)
}

func (p *ReadChatMediaUnreadToInboxResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadChatMediaUnreadToInboxResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReadChatMediaUnreadToInboxResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReadChatMediaUnreadToInboxResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadChatMediaUnreadToInboxResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return ReadChatMediaUnreadToInboxResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReadChatMediaUnreadToInboxResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *ReadChatMediaUnreadToInboxResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReadChatMediaUnreadToInboxResult) GetResult() interface{} {
	return p.Success
}

func updateHistoryReadedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateHistoryReadedArgs)
	realResult := result.(*UpdateHistoryReadedResult)
	success, err := handler.(inbox.RPCInbox).InboxUpdateHistoryReaded(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateHistoryReadedArgs() interface{} {
	return &UpdateHistoryReadedArgs{}
}

func newUpdateHistoryReadedResult() interface{} {
	return &UpdateHistoryReadedResult{}
}

type UpdateHistoryReadedArgs struct {
	Req *inbox.TLInboxUpdateHistoryReaded
}

func (p *UpdateHistoryReadedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateHistoryReadedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateHistoryReadedArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxUpdateHistoryReaded)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateHistoryReadedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateHistoryReadedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateHistoryReadedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxUpdateHistoryReaded)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateHistoryReadedArgs_Req_DEFAULT *inbox.TLInboxUpdateHistoryReaded

func (p *UpdateHistoryReadedArgs) GetReq() *inbox.TLInboxUpdateHistoryReaded {
	if !p.IsSetReq() {
		return UpdateHistoryReadedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateHistoryReadedArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateHistoryReadedResult struct {
	Success *tg.Void
}

var UpdateHistoryReadedResult_Success_DEFAULT *tg.Void

func (p *UpdateHistoryReadedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateHistoryReadedResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateHistoryReadedResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateHistoryReadedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateHistoryReadedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateHistoryReadedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateHistoryReadedResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return UpdateHistoryReadedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateHistoryReadedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *UpdateHistoryReadedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateHistoryReadedResult) GetResult() interface{} {
	return p.Success
}

func updatePinnedMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatePinnedMessageArgs)
	realResult := result.(*UpdatePinnedMessageResult)
	success, err := handler.(inbox.RPCInbox).InboxUpdatePinnedMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatePinnedMessageArgs() interface{} {
	return &UpdatePinnedMessageArgs{}
}

func newUpdatePinnedMessageResult() interface{} {
	return &UpdatePinnedMessageResult{}
}

type UpdatePinnedMessageArgs struct {
	Req *inbox.TLInboxUpdatePinnedMessage
}

func (p *UpdatePinnedMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatePinnedMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatePinnedMessageArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxUpdatePinnedMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatePinnedMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatePinnedMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatePinnedMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxUpdatePinnedMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatePinnedMessageArgs_Req_DEFAULT *inbox.TLInboxUpdatePinnedMessage

func (p *UpdatePinnedMessageArgs) GetReq() *inbox.TLInboxUpdatePinnedMessage {
	if !p.IsSetReq() {
		return UpdatePinnedMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePinnedMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatePinnedMessageResult struct {
	Success *tg.Void
}

var UpdatePinnedMessageResult_Success_DEFAULT *tg.Void

func (p *UpdatePinnedMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatePinnedMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatePinnedMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePinnedMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatePinnedMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatePinnedMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePinnedMessageResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return UpdatePinnedMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePinnedMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *UpdatePinnedMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePinnedMessageResult) GetResult() interface{} {
	return p.Success
}

func unpinAllMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UnpinAllMessagesArgs)
	realResult := result.(*UnpinAllMessagesResult)
	success, err := handler.(inbox.RPCInbox).InboxUnpinAllMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUnpinAllMessagesArgs() interface{} {
	return &UnpinAllMessagesArgs{}
}

func newUnpinAllMessagesResult() interface{} {
	return &UnpinAllMessagesResult{}
}

type UnpinAllMessagesArgs struct {
	Req *inbox.TLInboxUnpinAllMessages
}

func (p *UnpinAllMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UnpinAllMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UnpinAllMessagesArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxUnpinAllMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UnpinAllMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UnpinAllMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UnpinAllMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxUnpinAllMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UnpinAllMessagesArgs_Req_DEFAULT *inbox.TLInboxUnpinAllMessages

func (p *UnpinAllMessagesArgs) GetReq() *inbox.TLInboxUnpinAllMessages {
	if !p.IsSetReq() {
		return UnpinAllMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UnpinAllMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type UnpinAllMessagesResult struct {
	Success *tg.Void
}

var UnpinAllMessagesResult_Success_DEFAULT *tg.Void

func (p *UnpinAllMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UnpinAllMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *UnpinAllMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnpinAllMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UnpinAllMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UnpinAllMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnpinAllMessagesResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return UnpinAllMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UnpinAllMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *UnpinAllMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UnpinAllMessagesResult) GetResult() interface{} {
	return p.Success
}

func sendUserMessageToInboxV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SendUserMessageToInboxV2Args)
	realResult := result.(*SendUserMessageToInboxV2Result)
	success, err := handler.(inbox.RPCInbox).InboxSendUserMessageToInboxV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSendUserMessageToInboxV2Args() interface{} {
	return &SendUserMessageToInboxV2Args{}
}

func newSendUserMessageToInboxV2Result() interface{} {
	return &SendUserMessageToInboxV2Result{}
}

type SendUserMessageToInboxV2Args struct {
	Req *inbox.TLInboxSendUserMessageToInboxV2
}

func (p *SendUserMessageToInboxV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendUserMessageToInboxV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *SendUserMessageToInboxV2Args) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxSendUserMessageToInboxV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SendUserMessageToInboxV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SendUserMessageToInboxV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *SendUserMessageToInboxV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxSendUserMessageToInboxV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SendUserMessageToInboxV2Args_Req_DEFAULT *inbox.TLInboxSendUserMessageToInboxV2

func (p *SendUserMessageToInboxV2Args) GetReq() *inbox.TLInboxSendUserMessageToInboxV2 {
	if !p.IsSetReq() {
		return SendUserMessageToInboxV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *SendUserMessageToInboxV2Args) IsSetReq() bool {
	return p.Req != nil
}

type SendUserMessageToInboxV2Result struct {
	Success *tg.Void
}

var SendUserMessageToInboxV2Result_Success_DEFAULT *tg.Void

func (p *SendUserMessageToInboxV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SendUserMessageToInboxV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *SendUserMessageToInboxV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendUserMessageToInboxV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SendUserMessageToInboxV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *SendUserMessageToInboxV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendUserMessageToInboxV2Result) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return SendUserMessageToInboxV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *SendUserMessageToInboxV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *SendUserMessageToInboxV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendUserMessageToInboxV2Result) GetResult() interface{} {
	return p.Success
}

func editMessageToInboxV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditMessageToInboxV2Args)
	realResult := result.(*EditMessageToInboxV2Result)
	success, err := handler.(inbox.RPCInbox).InboxEditMessageToInboxV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditMessageToInboxV2Args() interface{} {
	return &EditMessageToInboxV2Args{}
}

func newEditMessageToInboxV2Result() interface{} {
	return &EditMessageToInboxV2Result{}
}

type EditMessageToInboxV2Args struct {
	Req *inbox.TLInboxEditMessageToInboxV2
}

func (p *EditMessageToInboxV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditMessageToInboxV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *EditMessageToInboxV2Args) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxEditMessageToInboxV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditMessageToInboxV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditMessageToInboxV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditMessageToInboxV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxEditMessageToInboxV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditMessageToInboxV2Args_Req_DEFAULT *inbox.TLInboxEditMessageToInboxV2

func (p *EditMessageToInboxV2Args) GetReq() *inbox.TLInboxEditMessageToInboxV2 {
	if !p.IsSetReq() {
		return EditMessageToInboxV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *EditMessageToInboxV2Args) IsSetReq() bool {
	return p.Req != nil
}

type EditMessageToInboxV2Result struct {
	Success *tg.Void
}

var EditMessageToInboxV2Result_Success_DEFAULT *tg.Void

func (p *EditMessageToInboxV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditMessageToInboxV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *EditMessageToInboxV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditMessageToInboxV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditMessageToInboxV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditMessageToInboxV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditMessageToInboxV2Result) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return EditMessageToInboxV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *EditMessageToInboxV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *EditMessageToInboxV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditMessageToInboxV2Result) GetResult() interface{} {
	return p.Success
}

func readInboxHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReadInboxHistoryArgs)
	realResult := result.(*ReadInboxHistoryResult)
	success, err := handler.(inbox.RPCInbox).InboxReadInboxHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReadInboxHistoryArgs() interface{} {
	return &ReadInboxHistoryArgs{}
}

func newReadInboxHistoryResult() interface{} {
	return &ReadInboxHistoryResult{}
}

type ReadInboxHistoryArgs struct {
	Req *inbox.TLInboxReadInboxHistory
}

func (p *ReadInboxHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReadInboxHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ReadInboxHistoryArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxReadInboxHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReadInboxHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReadInboxHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReadInboxHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxReadInboxHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReadInboxHistoryArgs_Req_DEFAULT *inbox.TLInboxReadInboxHistory

func (p *ReadInboxHistoryArgs) GetReq() *inbox.TLInboxReadInboxHistory {
	if !p.IsSetReq() {
		return ReadInboxHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReadInboxHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type ReadInboxHistoryResult struct {
	Success *tg.Void
}

var ReadInboxHistoryResult_Success_DEFAULT *tg.Void

func (p *ReadInboxHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReadInboxHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *ReadInboxHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadInboxHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReadInboxHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReadInboxHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadInboxHistoryResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return ReadInboxHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReadInboxHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *ReadInboxHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReadInboxHistoryResult) GetResult() interface{} {
	return p.Success
}

func readOutboxHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReadOutboxHistoryArgs)
	realResult := result.(*ReadOutboxHistoryResult)
	success, err := handler.(inbox.RPCInbox).InboxReadOutboxHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReadOutboxHistoryArgs() interface{} {
	return &ReadOutboxHistoryArgs{}
}

func newReadOutboxHistoryResult() interface{} {
	return &ReadOutboxHistoryResult{}
}

type ReadOutboxHistoryArgs struct {
	Req *inbox.TLInboxReadOutboxHistory
}

func (p *ReadOutboxHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReadOutboxHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ReadOutboxHistoryArgs) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxReadOutboxHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReadOutboxHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReadOutboxHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReadOutboxHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxReadOutboxHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReadOutboxHistoryArgs_Req_DEFAULT *inbox.TLInboxReadOutboxHistory

func (p *ReadOutboxHistoryArgs) GetReq() *inbox.TLInboxReadOutboxHistory {
	if !p.IsSetReq() {
		return ReadOutboxHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReadOutboxHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type ReadOutboxHistoryResult struct {
	Success *tg.Void
}

var ReadOutboxHistoryResult_Success_DEFAULT *tg.Void

func (p *ReadOutboxHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReadOutboxHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *ReadOutboxHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadOutboxHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReadOutboxHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReadOutboxHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadOutboxHistoryResult) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return ReadOutboxHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReadOutboxHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *ReadOutboxHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReadOutboxHistoryResult) GetResult() interface{} {
	return p.Success
}

func readMediaUnreadToInboxV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReadMediaUnreadToInboxV2Args)
	realResult := result.(*ReadMediaUnreadToInboxV2Result)
	success, err := handler.(inbox.RPCInbox).InboxReadMediaUnreadToInboxV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReadMediaUnreadToInboxV2Args() interface{} {
	return &ReadMediaUnreadToInboxV2Args{}
}

func newReadMediaUnreadToInboxV2Result() interface{} {
	return &ReadMediaUnreadToInboxV2Result{}
}

type ReadMediaUnreadToInboxV2Args struct {
	Req *inbox.TLInboxReadMediaUnreadToInboxV2
}

func (p *ReadMediaUnreadToInboxV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReadMediaUnreadToInboxV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *ReadMediaUnreadToInboxV2Args) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxReadMediaUnreadToInboxV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReadMediaUnreadToInboxV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReadMediaUnreadToInboxV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReadMediaUnreadToInboxV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxReadMediaUnreadToInboxV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReadMediaUnreadToInboxV2Args_Req_DEFAULT *inbox.TLInboxReadMediaUnreadToInboxV2

func (p *ReadMediaUnreadToInboxV2Args) GetReq() *inbox.TLInboxReadMediaUnreadToInboxV2 {
	if !p.IsSetReq() {
		return ReadMediaUnreadToInboxV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *ReadMediaUnreadToInboxV2Args) IsSetReq() bool {
	return p.Req != nil
}

type ReadMediaUnreadToInboxV2Result struct {
	Success *tg.Void
}

var ReadMediaUnreadToInboxV2Result_Success_DEFAULT *tg.Void

func (p *ReadMediaUnreadToInboxV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReadMediaUnreadToInboxV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *ReadMediaUnreadToInboxV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadMediaUnreadToInboxV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReadMediaUnreadToInboxV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReadMediaUnreadToInboxV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadMediaUnreadToInboxV2Result) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return ReadMediaUnreadToInboxV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *ReadMediaUnreadToInboxV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *ReadMediaUnreadToInboxV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReadMediaUnreadToInboxV2Result) GetResult() interface{} {
	return p.Success
}

func updatePinnedMessageV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatePinnedMessageV2Args)
	realResult := result.(*UpdatePinnedMessageV2Result)
	success, err := handler.(inbox.RPCInbox).InboxUpdatePinnedMessageV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatePinnedMessageV2Args() interface{} {
	return &UpdatePinnedMessageV2Args{}
}

func newUpdatePinnedMessageV2Result() interface{} {
	return &UpdatePinnedMessageV2Result{}
}

type UpdatePinnedMessageV2Args struct {
	Req *inbox.TLInboxUpdatePinnedMessageV2
}

func (p *UpdatePinnedMessageV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatePinnedMessageV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatePinnedMessageV2Args) Unmarshal(in []byte) error {
	msg := new(inbox.TLInboxUpdatePinnedMessageV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatePinnedMessageV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatePinnedMessageV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatePinnedMessageV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(inbox.TLInboxUpdatePinnedMessageV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatePinnedMessageV2Args_Req_DEFAULT *inbox.TLInboxUpdatePinnedMessageV2

func (p *UpdatePinnedMessageV2Args) GetReq() *inbox.TLInboxUpdatePinnedMessageV2 {
	if !p.IsSetReq() {
		return UpdatePinnedMessageV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePinnedMessageV2Args) IsSetReq() bool {
	return p.Req != nil
}

type UpdatePinnedMessageV2Result struct {
	Success *tg.Void
}

var UpdatePinnedMessageV2Result_Success_DEFAULT *tg.Void

func (p *UpdatePinnedMessageV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatePinnedMessageV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatePinnedMessageV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Void)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePinnedMessageV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatePinnedMessageV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatePinnedMessageV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Void)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePinnedMessageV2Result) GetSuccess() *tg.Void {
	if !p.IsSetSuccess() {
		return UpdatePinnedMessageV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePinnedMessageV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Void)
}

func (p *UpdatePinnedMessageV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePinnedMessageV2Result) GetResult() interface{} {
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

func (p *kClient) InboxEditUserMessageToInbox(ctx context.Context, req *inbox.TLInboxEditUserMessageToInbox) (r *tg.Void, err error) {
	// var _args EditUserMessageToInboxArgs
	// _args.Req = req
	// var _result EditUserMessageToInboxResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.editUserMessageToInbox", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxEditChatMessageToInbox(ctx context.Context, req *inbox.TLInboxEditChatMessageToInbox) (r *tg.Void, err error) {
	// var _args EditChatMessageToInboxArgs
	// _args.Req = req
	// var _result EditChatMessageToInboxResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.editChatMessageToInbox", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxDeleteMessagesToInbox(ctx context.Context, req *inbox.TLInboxDeleteMessagesToInbox) (r *tg.Void, err error) {
	// var _args DeleteMessagesToInboxArgs
	// _args.Req = req
	// var _result DeleteMessagesToInboxResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.deleteMessagesToInbox", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxDeleteUserHistoryToInbox(ctx context.Context, req *inbox.TLInboxDeleteUserHistoryToInbox) (r *tg.Void, err error) {
	// var _args DeleteUserHistoryToInboxArgs
	// _args.Req = req
	// var _result DeleteUserHistoryToInboxResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.deleteUserHistoryToInbox", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxDeleteChatHistoryToInbox(ctx context.Context, req *inbox.TLInboxDeleteChatHistoryToInbox) (r *tg.Void, err error) {
	// var _args DeleteChatHistoryToInboxArgs
	// _args.Req = req
	// var _result DeleteChatHistoryToInboxResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.deleteChatHistoryToInbox", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxReadUserMediaUnreadToInbox(ctx context.Context, req *inbox.TLInboxReadUserMediaUnreadToInbox) (r *tg.Void, err error) {
	// var _args ReadUserMediaUnreadToInboxArgs
	// _args.Req = req
	// var _result ReadUserMediaUnreadToInboxResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.readUserMediaUnreadToInbox", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxReadChatMediaUnreadToInbox(ctx context.Context, req *inbox.TLInboxReadChatMediaUnreadToInbox) (r *tg.Void, err error) {
	// var _args ReadChatMediaUnreadToInboxArgs
	// _args.Req = req
	// var _result ReadChatMediaUnreadToInboxResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.readChatMediaUnreadToInbox", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxUpdateHistoryReaded(ctx context.Context, req *inbox.TLInboxUpdateHistoryReaded) (r *tg.Void, err error) {
	// var _args UpdateHistoryReadedArgs
	// _args.Req = req
	// var _result UpdateHistoryReadedResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.updateHistoryReaded", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxUpdatePinnedMessage(ctx context.Context, req *inbox.TLInboxUpdatePinnedMessage) (r *tg.Void, err error) {
	// var _args UpdatePinnedMessageArgs
	// _args.Req = req
	// var _result UpdatePinnedMessageResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.updatePinnedMessage", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxUnpinAllMessages(ctx context.Context, req *inbox.TLInboxUnpinAllMessages) (r *tg.Void, err error) {
	// var _args UnpinAllMessagesArgs
	// _args.Req = req
	// var _result UnpinAllMessagesResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.unpinAllMessages", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxSendUserMessageToInboxV2(ctx context.Context, req *inbox.TLInboxSendUserMessageToInboxV2) (r *tg.Void, err error) {
	// var _args SendUserMessageToInboxV2Args
	// _args.Req = req
	// var _result SendUserMessageToInboxV2Result

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.sendUserMessageToInboxV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxEditMessageToInboxV2(ctx context.Context, req *inbox.TLInboxEditMessageToInboxV2) (r *tg.Void, err error) {
	// var _args EditMessageToInboxV2Args
	// _args.Req = req
	// var _result EditMessageToInboxV2Result

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.editMessageToInboxV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxReadInboxHistory(ctx context.Context, req *inbox.TLInboxReadInboxHistory) (r *tg.Void, err error) {
	// var _args ReadInboxHistoryArgs
	// _args.Req = req
	// var _result ReadInboxHistoryResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.readInboxHistory", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxReadOutboxHistory(ctx context.Context, req *inbox.TLInboxReadOutboxHistory) (r *tg.Void, err error) {
	// var _args ReadOutboxHistoryArgs
	// _args.Req = req
	// var _result ReadOutboxHistoryResult

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.readOutboxHistory", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxReadMediaUnreadToInboxV2(ctx context.Context, req *inbox.TLInboxReadMediaUnreadToInboxV2) (r *tg.Void, err error) {
	// var _args ReadMediaUnreadToInboxV2Args
	// _args.Req = req
	// var _result ReadMediaUnreadToInboxV2Result

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.readMediaUnreadToInboxV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) InboxUpdatePinnedMessageV2(ctx context.Context, req *inbox.TLInboxUpdatePinnedMessageV2) (r *tg.Void, err error) {
	// var _args UpdatePinnedMessageV2Args
	// _args.Req = req
	// var _result UpdatePinnedMessageV2Result

	_result := new(tg.Void)

	if err = p.c.Call(ctx, "/inbox.RPCInbox/inbox.updatePinnedMessageV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
