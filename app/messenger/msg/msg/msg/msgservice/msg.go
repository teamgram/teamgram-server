/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msgservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"msg.pushUserMessage": kitex.NewMethodInfo(
		pushUserMessageHandler,
		newPushUserMessageArgs,
		newPushUserMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.readMessageContents": kitex.NewMethodInfo(
		readMessageContentsHandler,
		newReadMessageContentsArgs,
		newReadMessageContentsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.sendMessageV2": kitex.NewMethodInfo(
		sendMessageV2Handler,
		newSendMessageV2Args,
		newSendMessageV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.editMessageV2": kitex.NewMethodInfo(
		editMessageV2Handler,
		newEditMessageV2Args,
		newEditMessageV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.deleteMessages": kitex.NewMethodInfo(
		deleteMessagesHandler,
		newDeleteMessagesArgs,
		newDeleteMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.deleteHistory": kitex.NewMethodInfo(
		deleteHistoryHandler,
		newDeleteHistoryArgs,
		newDeleteHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.deletePhoneCallHistory": kitex.NewMethodInfo(
		deletePhoneCallHistoryHandler,
		newDeletePhoneCallHistoryArgs,
		newDeletePhoneCallHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.deleteChatHistory": kitex.NewMethodInfo(
		deleteChatHistoryHandler,
		newDeleteChatHistoryArgs,
		newDeleteChatHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.readHistory": kitex.NewMethodInfo(
		readHistoryHandler,
		newReadHistoryArgs,
		newReadHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.readHistoryV2": kitex.NewMethodInfo(
		readHistoryV2Handler,
		newReadHistoryV2Args,
		newReadHistoryV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.updatePinnedMessage": kitex.NewMethodInfo(
		updatePinnedMessageHandler,
		newUpdatePinnedMessageArgs,
		newUpdatePinnedMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"msg.unpinAllMessages": kitex.NewMethodInfo(
		unpinAllMessagesHandler,
		newUnpinAllMessagesArgs,
		newUnpinAllMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	msgServiceServiceInfo                = NewServiceInfo()
	msgServiceServiceInfoForClient       = NewServiceInfoForClient()
	msgServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return msgServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return msgServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return msgServiceServiceInfoForClient
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
	serviceName := "RPCMsg"
	handlerType := (*msg.RPCMsg)(nil)
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
		"PackageName": "msg",
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

func pushUserMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PushUserMessageArgs)
	realResult := result.(*PushUserMessageResult)
	success, err := handler.(msg.RPCMsg).MsgPushUserMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPushUserMessageArgs() interface{} {
	return &PushUserMessageArgs{}
}

func newPushUserMessageResult() interface{} {
	return &PushUserMessageResult{}
}

type PushUserMessageArgs struct {
	Req *msg.TLMsgPushUserMessage
}

func (p *PushUserMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PushUserMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PushUserMessageArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgPushUserMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PushUserMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PushUserMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PushUserMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgPushUserMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PushUserMessageArgs_Req_DEFAULT *msg.TLMsgPushUserMessage

func (p *PushUserMessageArgs) GetReq() *msg.TLMsgPushUserMessage {
	if !p.IsSetReq() {
		return PushUserMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushUserMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type PushUserMessageResult struct {
	Success *tg.Bool
}

var PushUserMessageResult_Success_DEFAULT *tg.Bool

func (p *PushUserMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PushUserMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *PushUserMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushUserMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PushUserMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PushUserMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushUserMessageResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return PushUserMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushUserMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *PushUserMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushUserMessageResult) GetResult() interface{} {
	return p.Success
}

func readMessageContentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReadMessageContentsArgs)
	realResult := result.(*ReadMessageContentsResult)
	success, err := handler.(msg.RPCMsg).MsgReadMessageContents(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReadMessageContentsArgs() interface{} {
	return &ReadMessageContentsArgs{}
}

func newReadMessageContentsResult() interface{} {
	return &ReadMessageContentsResult{}
}

type ReadMessageContentsArgs struct {
	Req *msg.TLMsgReadMessageContents
}

func (p *ReadMessageContentsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReadMessageContentsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ReadMessageContentsArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgReadMessageContents)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReadMessageContentsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReadMessageContentsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReadMessageContentsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgReadMessageContents)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReadMessageContentsArgs_Req_DEFAULT *msg.TLMsgReadMessageContents

func (p *ReadMessageContentsArgs) GetReq() *msg.TLMsgReadMessageContents {
	if !p.IsSetReq() {
		return ReadMessageContentsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReadMessageContentsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ReadMessageContentsResult struct {
	Success *tg.MessagesAffectedMessages
}

var ReadMessageContentsResult_Success_DEFAULT *tg.MessagesAffectedMessages

func (p *ReadMessageContentsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReadMessageContentsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ReadMessageContentsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadMessageContentsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReadMessageContentsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReadMessageContentsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadMessageContentsResult) GetSuccess() *tg.MessagesAffectedMessages {
	if !p.IsSetSuccess() {
		return ReadMessageContentsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReadMessageContentsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedMessages)
}

func (p *ReadMessageContentsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReadMessageContentsResult) GetResult() interface{} {
	return p.Success
}

func sendMessageV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SendMessageV2Args)
	realResult := result.(*SendMessageV2Result)
	success, err := handler.(msg.RPCMsg).MsgSendMessageV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSendMessageV2Args() interface{} {
	return &SendMessageV2Args{}
}

func newSendMessageV2Result() interface{} {
	return &SendMessageV2Result{}
}

type SendMessageV2Args struct {
	Req *msg.TLMsgSendMessageV2
}

func (p *SendMessageV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendMessageV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *SendMessageV2Args) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgSendMessageV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SendMessageV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SendMessageV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *SendMessageV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgSendMessageV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SendMessageV2Args_Req_DEFAULT *msg.TLMsgSendMessageV2

func (p *SendMessageV2Args) GetReq() *msg.TLMsgSendMessageV2 {
	if !p.IsSetReq() {
		return SendMessageV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *SendMessageV2Args) IsSetReq() bool {
	return p.Req != nil
}

type SendMessageV2Result struct {
	Success *tg.Updates
}

var SendMessageV2Result_Success_DEFAULT *tg.Updates

func (p *SendMessageV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SendMessageV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *SendMessageV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendMessageV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SendMessageV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *SendMessageV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendMessageV2Result) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return SendMessageV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *SendMessageV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *SendMessageV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendMessageV2Result) GetResult() interface{} {
	return p.Success
}

func editMessageV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditMessageV2Args)
	realResult := result.(*EditMessageV2Result)
	success, err := handler.(msg.RPCMsg).MsgEditMessageV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditMessageV2Args() interface{} {
	return &EditMessageV2Args{}
}

func newEditMessageV2Result() interface{} {
	return &EditMessageV2Result{}
}

type EditMessageV2Args struct {
	Req *msg.TLMsgEditMessageV2
}

func (p *EditMessageV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditMessageV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *EditMessageV2Args) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgEditMessageV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditMessageV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditMessageV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditMessageV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgEditMessageV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditMessageV2Args_Req_DEFAULT *msg.TLMsgEditMessageV2

func (p *EditMessageV2Args) GetReq() *msg.TLMsgEditMessageV2 {
	if !p.IsSetReq() {
		return EditMessageV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *EditMessageV2Args) IsSetReq() bool {
	return p.Req != nil
}

type EditMessageV2Result struct {
	Success *tg.Updates
}

var EditMessageV2Result_Success_DEFAULT *tg.Updates

func (p *EditMessageV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditMessageV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *EditMessageV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditMessageV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditMessageV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditMessageV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditMessageV2Result) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return EditMessageV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *EditMessageV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *EditMessageV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditMessageV2Result) GetResult() interface{} {
	return p.Success
}

func deleteMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteMessagesArgs)
	realResult := result.(*DeleteMessagesResult)
	success, err := handler.(msg.RPCMsg).MsgDeleteMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteMessagesArgs() interface{} {
	return &DeleteMessagesArgs{}
}

func newDeleteMessagesResult() interface{} {
	return &DeleteMessagesResult{}
}

type DeleteMessagesArgs struct {
	Req *msg.TLMsgDeleteMessages
}

func (p *DeleteMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteMessagesArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgDeleteMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgDeleteMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteMessagesArgs_Req_DEFAULT *msg.TLMsgDeleteMessages

func (p *DeleteMessagesArgs) GetReq() *msg.TLMsgDeleteMessages {
	if !p.IsSetReq() {
		return DeleteMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteMessagesResult struct {
	Success *tg.MessagesAffectedMessages
}

var DeleteMessagesResult_Success_DEFAULT *tg.MessagesAffectedMessages

func (p *DeleteMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteMessagesResult) GetSuccess() *tg.MessagesAffectedMessages {
	if !p.IsSetSuccess() {
		return DeleteMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedMessages)
}

func (p *DeleteMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteMessagesResult) GetResult() interface{} {
	return p.Success
}

func deleteHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteHistoryArgs)
	realResult := result.(*DeleteHistoryResult)
	success, err := handler.(msg.RPCMsg).MsgDeleteHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteHistoryArgs() interface{} {
	return &DeleteHistoryArgs{}
}

func newDeleteHistoryResult() interface{} {
	return &DeleteHistoryResult{}
}

type DeleteHistoryArgs struct {
	Req *msg.TLMsgDeleteHistory
}

func (p *DeleteHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteHistoryArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgDeleteHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgDeleteHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteHistoryArgs_Req_DEFAULT *msg.TLMsgDeleteHistory

func (p *DeleteHistoryArgs) GetReq() *msg.TLMsgDeleteHistory {
	if !p.IsSetReq() {
		return DeleteHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteHistoryResult struct {
	Success *tg.MessagesAffectedHistory
}

var DeleteHistoryResult_Success_DEFAULT *tg.MessagesAffectedHistory

func (p *DeleteHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedHistory)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteHistoryResult) GetSuccess() *tg.MessagesAffectedHistory {
	if !p.IsSetSuccess() {
		return DeleteHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedHistory)
}

func (p *DeleteHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteHistoryResult) GetResult() interface{} {
	return p.Success
}

func deletePhoneCallHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeletePhoneCallHistoryArgs)
	realResult := result.(*DeletePhoneCallHistoryResult)
	success, err := handler.(msg.RPCMsg).MsgDeletePhoneCallHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeletePhoneCallHistoryArgs() interface{} {
	return &DeletePhoneCallHistoryArgs{}
}

func newDeletePhoneCallHistoryResult() interface{} {
	return &DeletePhoneCallHistoryResult{}
}

type DeletePhoneCallHistoryArgs struct {
	Req *msg.TLMsgDeletePhoneCallHistory
}

func (p *DeletePhoneCallHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeletePhoneCallHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeletePhoneCallHistoryArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgDeletePhoneCallHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeletePhoneCallHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeletePhoneCallHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeletePhoneCallHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgDeletePhoneCallHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeletePhoneCallHistoryArgs_Req_DEFAULT *msg.TLMsgDeletePhoneCallHistory

func (p *DeletePhoneCallHistoryArgs) GetReq() *msg.TLMsgDeletePhoneCallHistory {
	if !p.IsSetReq() {
		return DeletePhoneCallHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePhoneCallHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeletePhoneCallHistoryResult struct {
	Success *tg.MessagesAffectedFoundMessages
}

var DeletePhoneCallHistoryResult_Success_DEFAULT *tg.MessagesAffectedFoundMessages

func (p *DeletePhoneCallHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeletePhoneCallHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeletePhoneCallHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedFoundMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePhoneCallHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeletePhoneCallHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeletePhoneCallHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedFoundMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePhoneCallHistoryResult) GetSuccess() *tg.MessagesAffectedFoundMessages {
	if !p.IsSetSuccess() {
		return DeletePhoneCallHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePhoneCallHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedFoundMessages)
}

func (p *DeletePhoneCallHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeletePhoneCallHistoryResult) GetResult() interface{} {
	return p.Success
}

func deleteChatHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteChatHistoryArgs)
	realResult := result.(*DeleteChatHistoryResult)
	success, err := handler.(msg.RPCMsg).MsgDeleteChatHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteChatHistoryArgs() interface{} {
	return &DeleteChatHistoryArgs{}
}

func newDeleteChatHistoryResult() interface{} {
	return &DeleteChatHistoryResult{}
}

type DeleteChatHistoryArgs struct {
	Req *msg.TLMsgDeleteChatHistory
}

func (p *DeleteChatHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteChatHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteChatHistoryArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgDeleteChatHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteChatHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteChatHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteChatHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgDeleteChatHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteChatHistoryArgs_Req_DEFAULT *msg.TLMsgDeleteChatHistory

func (p *DeleteChatHistoryArgs) GetReq() *msg.TLMsgDeleteChatHistory {
	if !p.IsSetReq() {
		return DeleteChatHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteChatHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteChatHistoryResult struct {
	Success *tg.Bool
}

var DeleteChatHistoryResult_Success_DEFAULT *tg.Bool

func (p *DeleteChatHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteChatHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteChatHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteChatHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteChatHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteChatHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteChatHistoryResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteChatHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteChatHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteChatHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteChatHistoryResult) GetResult() interface{} {
	return p.Success
}

func readHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReadHistoryArgs)
	realResult := result.(*ReadHistoryResult)
	success, err := handler.(msg.RPCMsg).MsgReadHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReadHistoryArgs() interface{} {
	return &ReadHistoryArgs{}
}

func newReadHistoryResult() interface{} {
	return &ReadHistoryResult{}
}

type ReadHistoryArgs struct {
	Req *msg.TLMsgReadHistory
}

func (p *ReadHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReadHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ReadHistoryArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgReadHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReadHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReadHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReadHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgReadHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReadHistoryArgs_Req_DEFAULT *msg.TLMsgReadHistory

func (p *ReadHistoryArgs) GetReq() *msg.TLMsgReadHistory {
	if !p.IsSetReq() {
		return ReadHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReadHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type ReadHistoryResult struct {
	Success *tg.MessagesAffectedMessages
}

var ReadHistoryResult_Success_DEFAULT *tg.MessagesAffectedMessages

func (p *ReadHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReadHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *ReadHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReadHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReadHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadHistoryResult) GetSuccess() *tg.MessagesAffectedMessages {
	if !p.IsSetSuccess() {
		return ReadHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReadHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedMessages)
}

func (p *ReadHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReadHistoryResult) GetResult() interface{} {
	return p.Success
}

func readHistoryV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReadHistoryV2Args)
	realResult := result.(*ReadHistoryV2Result)
	success, err := handler.(msg.RPCMsg).MsgReadHistoryV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReadHistoryV2Args() interface{} {
	return &ReadHistoryV2Args{}
}

func newReadHistoryV2Result() interface{} {
	return &ReadHistoryV2Result{}
}

type ReadHistoryV2Args struct {
	Req *msg.TLMsgReadHistoryV2
}

func (p *ReadHistoryV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReadHistoryV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *ReadHistoryV2Args) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgReadHistoryV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReadHistoryV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReadHistoryV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReadHistoryV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(msg.TLMsgReadHistoryV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReadHistoryV2Args_Req_DEFAULT *msg.TLMsgReadHistoryV2

func (p *ReadHistoryV2Args) GetReq() *msg.TLMsgReadHistoryV2 {
	if !p.IsSetReq() {
		return ReadHistoryV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *ReadHistoryV2Args) IsSetReq() bool {
	return p.Req != nil
}

type ReadHistoryV2Result struct {
	Success *tg.MessagesAffectedMessages
}

var ReadHistoryV2Result_Success_DEFAULT *tg.MessagesAffectedMessages

func (p *ReadHistoryV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReadHistoryV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *ReadHistoryV2Result) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadHistoryV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReadHistoryV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReadHistoryV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReadHistoryV2Result) GetSuccess() *tg.MessagesAffectedMessages {
	if !p.IsSetSuccess() {
		return ReadHistoryV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *ReadHistoryV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedMessages)
}

func (p *ReadHistoryV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReadHistoryV2Result) GetResult() interface{} {
	return p.Success
}

func updatePinnedMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatePinnedMessageArgs)
	realResult := result.(*UpdatePinnedMessageResult)
	success, err := handler.(msg.RPCMsg).MsgUpdatePinnedMessage(ctx, realArg.Req)
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
	Req *msg.TLMsgUpdatePinnedMessage
}

func (p *UpdatePinnedMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatePinnedMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatePinnedMessageArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgUpdatePinnedMessage)
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
	msg := new(msg.TLMsgUpdatePinnedMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatePinnedMessageArgs_Req_DEFAULT *msg.TLMsgUpdatePinnedMessage

func (p *UpdatePinnedMessageArgs) GetReq() *msg.TLMsgUpdatePinnedMessage {
	if !p.IsSetReq() {
		return UpdatePinnedMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePinnedMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatePinnedMessageResult struct {
	Success *tg.Updates
}

var UpdatePinnedMessageResult_Success_DEFAULT *tg.Updates

func (p *UpdatePinnedMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatePinnedMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatePinnedMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
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
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePinnedMessageResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return UpdatePinnedMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePinnedMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
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
	success, err := handler.(msg.RPCMsg).MsgUnpinAllMessages(ctx, realArg.Req)
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
	Req *msg.TLMsgUnpinAllMessages
}

func (p *UnpinAllMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UnpinAllMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UnpinAllMessagesArgs) Unmarshal(in []byte) error {
	msg := new(msg.TLMsgUnpinAllMessages)
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
	msg := new(msg.TLMsgUnpinAllMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UnpinAllMessagesArgs_Req_DEFAULT *msg.TLMsgUnpinAllMessages

func (p *UnpinAllMessagesArgs) GetReq() *msg.TLMsgUnpinAllMessages {
	if !p.IsSetReq() {
		return UnpinAllMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UnpinAllMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type UnpinAllMessagesResult struct {
	Success *tg.MessagesAffectedHistory
}

var UnpinAllMessagesResult_Success_DEFAULT *tg.MessagesAffectedHistory

func (p *UnpinAllMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UnpinAllMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *UnpinAllMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedHistory)
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
	msg := new(tg.MessagesAffectedHistory)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnpinAllMessagesResult) GetSuccess() *tg.MessagesAffectedHistory {
	if !p.IsSetSuccess() {
		return UnpinAllMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UnpinAllMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedHistory)
}

func (p *UnpinAllMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UnpinAllMessagesResult) GetResult() interface{} {
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

func (p *kClient) MsgPushUserMessage(ctx context.Context, req *msg.TLMsgPushUserMessage) (r *tg.Bool, err error) {
	var _args PushUserMessageArgs
	_args.Req = req
	var _result PushUserMessageResult
	if err = p.c.Call(ctx, "msg.pushUserMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgReadMessageContents(ctx context.Context, req *msg.TLMsgReadMessageContents) (r *tg.MessagesAffectedMessages, err error) {
	var _args ReadMessageContentsArgs
	_args.Req = req
	var _result ReadMessageContentsResult
	if err = p.c.Call(ctx, "msg.readMessageContents", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgSendMessageV2(ctx context.Context, req *msg.TLMsgSendMessageV2) (r *tg.Updates, err error) {
	var _args SendMessageV2Args
	_args.Req = req
	var _result SendMessageV2Result
	if err = p.c.Call(ctx, "msg.sendMessageV2", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgEditMessageV2(ctx context.Context, req *msg.TLMsgEditMessageV2) (r *tg.Updates, err error) {
	var _args EditMessageV2Args
	_args.Req = req
	var _result EditMessageV2Result
	if err = p.c.Call(ctx, "msg.editMessageV2", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgDeleteMessages(ctx context.Context, req *msg.TLMsgDeleteMessages) (r *tg.MessagesAffectedMessages, err error) {
	var _args DeleteMessagesArgs
	_args.Req = req
	var _result DeleteMessagesResult
	if err = p.c.Call(ctx, "msg.deleteMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgDeleteHistory(ctx context.Context, req *msg.TLMsgDeleteHistory) (r *tg.MessagesAffectedHistory, err error) {
	var _args DeleteHistoryArgs
	_args.Req = req
	var _result DeleteHistoryResult
	if err = p.c.Call(ctx, "msg.deleteHistory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgDeletePhoneCallHistory(ctx context.Context, req *msg.TLMsgDeletePhoneCallHistory) (r *tg.MessagesAffectedFoundMessages, err error) {
	var _args DeletePhoneCallHistoryArgs
	_args.Req = req
	var _result DeletePhoneCallHistoryResult
	if err = p.c.Call(ctx, "msg.deletePhoneCallHistory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgDeleteChatHistory(ctx context.Context, req *msg.TLMsgDeleteChatHistory) (r *tg.Bool, err error) {
	var _args DeleteChatHistoryArgs
	_args.Req = req
	var _result DeleteChatHistoryResult
	if err = p.c.Call(ctx, "msg.deleteChatHistory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgReadHistory(ctx context.Context, req *msg.TLMsgReadHistory) (r *tg.MessagesAffectedMessages, err error) {
	var _args ReadHistoryArgs
	_args.Req = req
	var _result ReadHistoryResult
	if err = p.c.Call(ctx, "msg.readHistory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgReadHistoryV2(ctx context.Context, req *msg.TLMsgReadHistoryV2) (r *tg.MessagesAffectedMessages, err error) {
	var _args ReadHistoryV2Args
	_args.Req = req
	var _result ReadHistoryV2Result
	if err = p.c.Call(ctx, "msg.readHistoryV2", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgUpdatePinnedMessage(ctx context.Context, req *msg.TLMsgUpdatePinnedMessage) (r *tg.Updates, err error) {
	var _args UpdatePinnedMessageArgs
	_args.Req = req
	var _result UpdatePinnedMessageResult
	if err = p.c.Call(ctx, "msg.updatePinnedMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MsgUnpinAllMessages(ctx context.Context, req *msg.TLMsgUnpinAllMessages) (r *tg.MessagesAffectedHistory, err error) {
	var _args UnpinAllMessagesArgs
	_args.Req = req
	var _result UnpinAllMessagesResult
	if err = p.c.Call(ctx, "msg.unpinAllMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
