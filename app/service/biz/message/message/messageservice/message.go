/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package messageservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"message.getUserMessage": kitex.NewMethodInfo(
		getUserMessageHandler,
		newGetUserMessageArgs,
		newGetUserMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getUserMessageList": kitex.NewMethodInfo(
		getUserMessageListHandler,
		newGetUserMessageListArgs,
		newGetUserMessageListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getUserMessageListByDataIdList": kitex.NewMethodInfo(
		getUserMessageListByDataIdListHandler,
		newGetUserMessageListByDataIdListArgs,
		newGetUserMessageListByDataIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getUserMessageListByDataIdUserIdList": kitex.NewMethodInfo(
		getUserMessageListByDataIdUserIdListHandler,
		newGetUserMessageListByDataIdUserIdListArgs,
		newGetUserMessageListByDataIdUserIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getHistoryMessages": kitex.NewMethodInfo(
		getHistoryMessagesHandler,
		newGetHistoryMessagesArgs,
		newGetHistoryMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getHistoryMessagesCount": kitex.NewMethodInfo(
		getHistoryMessagesCountHandler,
		newGetHistoryMessagesCountArgs,
		newGetHistoryMessagesCountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getPeerUserMessageId": kitex.NewMethodInfo(
		getPeerUserMessageIdHandler,
		newGetPeerUserMessageIdArgs,
		newGetPeerUserMessageIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getPeerUserMessage": kitex.NewMethodInfo(
		getPeerUserMessageHandler,
		newGetPeerUserMessageArgs,
		newGetPeerUserMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.searchByMediaType": kitex.NewMethodInfo(
		searchByMediaTypeHandler,
		newSearchByMediaTypeArgs,
		newSearchByMediaTypeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.search": kitex.NewMethodInfo(
		searchHandler,
		newSearchArgs,
		newSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.searchGlobal": kitex.NewMethodInfo(
		searchGlobalHandler,
		newSearchGlobalArgs,
		newSearchGlobalResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.searchByPinned": kitex.NewMethodInfo(
		searchByPinnedHandler,
		newSearchByPinnedArgs,
		newSearchByPinnedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getSearchCounter": kitex.NewMethodInfo(
		getSearchCounterHandler,
		newGetSearchCounterArgs,
		newGetSearchCounterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.searchV2": kitex.NewMethodInfo(
		searchV2Handler,
		newSearchV2Args,
		newSearchV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getLastTwoPinnedMessageId": kitex.NewMethodInfo(
		getLastTwoPinnedMessageIdHandler,
		newGetLastTwoPinnedMessageIdArgs,
		newGetLastTwoPinnedMessageIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.updatePinnedMessageId": kitex.NewMethodInfo(
		updatePinnedMessageIdHandler,
		newUpdatePinnedMessageIdArgs,
		newUpdatePinnedMessageIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getPinnedMessageIdList": kitex.NewMethodInfo(
		getPinnedMessageIdListHandler,
		newGetPinnedMessageIdListArgs,
		newGetPinnedMessageIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.unPinAllMessages": kitex.NewMethodInfo(
		unPinAllMessagesHandler,
		newUnPinAllMessagesArgs,
		newUnPinAllMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getUnreadMentions": kitex.NewMethodInfo(
		getUnreadMentionsHandler,
		newGetUnreadMentionsArgs,
		newGetUnreadMentionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"message.getUnreadMentionsCount": kitex.NewMethodInfo(
		getUnreadMentionsCountHandler,
		newGetUnreadMentionsCountArgs,
		newGetUnreadMentionsCountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	messageServiceServiceInfo                = NewServiceInfo()
	messageServiceServiceInfoForClient       = NewServiceInfoForClient()
	messageServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCMessage", messageServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCMessage", messageServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCMessage", messageServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return messageServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return messageServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return messageServiceServiceInfoForClient
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
	serviceName := "RPCMessage"
	handlerType := (*message.RPCMessage)(nil)
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
		"PackageName": "message",
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

func getUserMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserMessageArgs)
	realResult := result.(*GetUserMessageResult)
	success, err := handler.(message.RPCMessage).MessageGetUserMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserMessageArgs() interface{} {
	return &GetUserMessageArgs{}
}

func newGetUserMessageResult() interface{} {
	return &GetUserMessageResult{}
}

type GetUserMessageArgs struct {
	Req *message.TLMessageGetUserMessage
}

func (p *GetUserMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserMessageArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetUserMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetUserMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserMessageArgs_Req_DEFAULT *message.TLMessageGetUserMessage

func (p *GetUserMessageArgs) GetReq() *message.TLMessageGetUserMessage {
	if !p.IsSetReq() {
		return GetUserMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserMessageResult struct {
	Success *tg.MessageBox
}

var GetUserMessageResult_Success_DEFAULT *tg.MessageBox

func (p *GetUserMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.MessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserMessageResult) GetSuccess() *tg.MessageBox {
	if !p.IsSetSuccess() {
		return GetUserMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessageBox)
}

func (p *GetUserMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserMessageResult) GetResult() interface{} {
	return p.Success
}

func getUserMessageListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserMessageListArgs)
	realResult := result.(*GetUserMessageListResult)
	success, err := handler.(message.RPCMessage).MessageGetUserMessageList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserMessageListArgs() interface{} {
	return &GetUserMessageListArgs{}
}

func newGetUserMessageListResult() interface{} {
	return &GetUserMessageListResult{}
}

type GetUserMessageListArgs struct {
	Req *message.TLMessageGetUserMessageList
}

func (p *GetUserMessageListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserMessageListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserMessageListArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetUserMessageList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserMessageListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserMessageListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserMessageListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetUserMessageList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserMessageListArgs_Req_DEFAULT *message.TLMessageGetUserMessageList

func (p *GetUserMessageListArgs) GetReq() *message.TLMessageGetUserMessageList {
	if !p.IsSetReq() {
		return GetUserMessageListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserMessageListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserMessageListResult struct {
	Success *message.VectorMessageBox
}

var GetUserMessageListResult_Success_DEFAULT *message.VectorMessageBox

func (p *GetUserMessageListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserMessageListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserMessageListResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserMessageListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserMessageListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserMessageListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserMessageListResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return GetUserMessageListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserMessageListResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *GetUserMessageListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserMessageListResult) GetResult() interface{} {
	return p.Success
}

func getUserMessageListByDataIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserMessageListByDataIdListArgs)
	realResult := result.(*GetUserMessageListByDataIdListResult)
	success, err := handler.(message.RPCMessage).MessageGetUserMessageListByDataIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserMessageListByDataIdListArgs() interface{} {
	return &GetUserMessageListByDataIdListArgs{}
}

func newGetUserMessageListByDataIdListResult() interface{} {
	return &GetUserMessageListByDataIdListResult{}
}

type GetUserMessageListByDataIdListArgs struct {
	Req *message.TLMessageGetUserMessageListByDataIdList
}

func (p *GetUserMessageListByDataIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserMessageListByDataIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserMessageListByDataIdListArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetUserMessageListByDataIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserMessageListByDataIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserMessageListByDataIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserMessageListByDataIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetUserMessageListByDataIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserMessageListByDataIdListArgs_Req_DEFAULT *message.TLMessageGetUserMessageListByDataIdList

func (p *GetUserMessageListByDataIdListArgs) GetReq() *message.TLMessageGetUserMessageListByDataIdList {
	if !p.IsSetReq() {
		return GetUserMessageListByDataIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserMessageListByDataIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserMessageListByDataIdListResult struct {
	Success *message.VectorMessageBox
}

var GetUserMessageListByDataIdListResult_Success_DEFAULT *message.VectorMessageBox

func (p *GetUserMessageListByDataIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserMessageListByDataIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserMessageListByDataIdListResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserMessageListByDataIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserMessageListByDataIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserMessageListByDataIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserMessageListByDataIdListResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return GetUserMessageListByDataIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserMessageListByDataIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *GetUserMessageListByDataIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserMessageListByDataIdListResult) GetResult() interface{} {
	return p.Success
}

func getUserMessageListByDataIdUserIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserMessageListByDataIdUserIdListArgs)
	realResult := result.(*GetUserMessageListByDataIdUserIdListResult)
	success, err := handler.(message.RPCMessage).MessageGetUserMessageListByDataIdUserIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserMessageListByDataIdUserIdListArgs() interface{} {
	return &GetUserMessageListByDataIdUserIdListArgs{}
}

func newGetUserMessageListByDataIdUserIdListResult() interface{} {
	return &GetUserMessageListByDataIdUserIdListResult{}
}

type GetUserMessageListByDataIdUserIdListArgs struct {
	Req *message.TLMessageGetUserMessageListByDataIdUserIdList
}

func (p *GetUserMessageListByDataIdUserIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserMessageListByDataIdUserIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserMessageListByDataIdUserIdListArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetUserMessageListByDataIdUserIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserMessageListByDataIdUserIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserMessageListByDataIdUserIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserMessageListByDataIdUserIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetUserMessageListByDataIdUserIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserMessageListByDataIdUserIdListArgs_Req_DEFAULT *message.TLMessageGetUserMessageListByDataIdUserIdList

func (p *GetUserMessageListByDataIdUserIdListArgs) GetReq() *message.TLMessageGetUserMessageListByDataIdUserIdList {
	if !p.IsSetReq() {
		return GetUserMessageListByDataIdUserIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserMessageListByDataIdUserIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserMessageListByDataIdUserIdListResult struct {
	Success *message.VectorMessageBox
}

var GetUserMessageListByDataIdUserIdListResult_Success_DEFAULT *message.VectorMessageBox

func (p *GetUserMessageListByDataIdUserIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserMessageListByDataIdUserIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserMessageListByDataIdUserIdListResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserMessageListByDataIdUserIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserMessageListByDataIdUserIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserMessageListByDataIdUserIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserMessageListByDataIdUserIdListResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return GetUserMessageListByDataIdUserIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserMessageListByDataIdUserIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *GetUserMessageListByDataIdUserIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserMessageListByDataIdUserIdListResult) GetResult() interface{} {
	return p.Success
}

func getHistoryMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetHistoryMessagesArgs)
	realResult := result.(*GetHistoryMessagesResult)
	success, err := handler.(message.RPCMessage).MessageGetHistoryMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetHistoryMessagesArgs() interface{} {
	return &GetHistoryMessagesArgs{}
}

func newGetHistoryMessagesResult() interface{} {
	return &GetHistoryMessagesResult{}
}

type GetHistoryMessagesArgs struct {
	Req *message.TLMessageGetHistoryMessages
}

func (p *GetHistoryMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetHistoryMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetHistoryMessagesArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetHistoryMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetHistoryMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetHistoryMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetHistoryMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetHistoryMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetHistoryMessagesArgs_Req_DEFAULT *message.TLMessageGetHistoryMessages

func (p *GetHistoryMessagesArgs) GetReq() *message.TLMessageGetHistoryMessages {
	if !p.IsSetReq() {
		return GetHistoryMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetHistoryMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetHistoryMessagesResult struct {
	Success *message.VectorMessageBox
}

var GetHistoryMessagesResult_Success_DEFAULT *message.VectorMessageBox

func (p *GetHistoryMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetHistoryMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetHistoryMessagesResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetHistoryMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetHistoryMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetHistoryMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetHistoryMessagesResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return GetHistoryMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetHistoryMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *GetHistoryMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetHistoryMessagesResult) GetResult() interface{} {
	return p.Success
}

func getHistoryMessagesCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetHistoryMessagesCountArgs)
	realResult := result.(*GetHistoryMessagesCountResult)
	success, err := handler.(message.RPCMessage).MessageGetHistoryMessagesCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetHistoryMessagesCountArgs() interface{} {
	return &GetHistoryMessagesCountArgs{}
}

func newGetHistoryMessagesCountResult() interface{} {
	return &GetHistoryMessagesCountResult{}
}

type GetHistoryMessagesCountArgs struct {
	Req *message.TLMessageGetHistoryMessagesCount
}

func (p *GetHistoryMessagesCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetHistoryMessagesCountArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetHistoryMessagesCountArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetHistoryMessagesCount)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetHistoryMessagesCountArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetHistoryMessagesCountArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetHistoryMessagesCountArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetHistoryMessagesCount)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetHistoryMessagesCountArgs_Req_DEFAULT *message.TLMessageGetHistoryMessagesCount

func (p *GetHistoryMessagesCountArgs) GetReq() *message.TLMessageGetHistoryMessagesCount {
	if !p.IsSetReq() {
		return GetHistoryMessagesCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetHistoryMessagesCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetHistoryMessagesCountResult struct {
	Success *tg.Int32
}

var GetHistoryMessagesCountResult_Success_DEFAULT *tg.Int32

func (p *GetHistoryMessagesCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetHistoryMessagesCountResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetHistoryMessagesCountResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetHistoryMessagesCountResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetHistoryMessagesCountResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetHistoryMessagesCountResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetHistoryMessagesCountResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetHistoryMessagesCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetHistoryMessagesCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetHistoryMessagesCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetHistoryMessagesCountResult) GetResult() interface{} {
	return p.Success
}

func getPeerUserMessageIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPeerUserMessageIdArgs)
	realResult := result.(*GetPeerUserMessageIdResult)
	success, err := handler.(message.RPCMessage).MessageGetPeerUserMessageId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPeerUserMessageIdArgs() interface{} {
	return &GetPeerUserMessageIdArgs{}
}

func newGetPeerUserMessageIdResult() interface{} {
	return &GetPeerUserMessageIdResult{}
}

type GetPeerUserMessageIdArgs struct {
	Req *message.TLMessageGetPeerUserMessageId
}

func (p *GetPeerUserMessageIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPeerUserMessageIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPeerUserMessageIdArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetPeerUserMessageId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPeerUserMessageIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPeerUserMessageIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPeerUserMessageIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetPeerUserMessageId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPeerUserMessageIdArgs_Req_DEFAULT *message.TLMessageGetPeerUserMessageId

func (p *GetPeerUserMessageIdArgs) GetReq() *message.TLMessageGetPeerUserMessageId {
	if !p.IsSetReq() {
		return GetPeerUserMessageIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPeerUserMessageIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPeerUserMessageIdResult struct {
	Success *tg.Int32
}

var GetPeerUserMessageIdResult_Success_DEFAULT *tg.Int32

func (p *GetPeerUserMessageIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPeerUserMessageIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPeerUserMessageIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPeerUserMessageIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPeerUserMessageIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPeerUserMessageIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPeerUserMessageIdResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetPeerUserMessageIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPeerUserMessageIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetPeerUserMessageIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPeerUserMessageIdResult) GetResult() interface{} {
	return p.Success
}

func getPeerUserMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPeerUserMessageArgs)
	realResult := result.(*GetPeerUserMessageResult)
	success, err := handler.(message.RPCMessage).MessageGetPeerUserMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPeerUserMessageArgs() interface{} {
	return &GetPeerUserMessageArgs{}
}

func newGetPeerUserMessageResult() interface{} {
	return &GetPeerUserMessageResult{}
}

type GetPeerUserMessageArgs struct {
	Req *message.TLMessageGetPeerUserMessage
}

func (p *GetPeerUserMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPeerUserMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPeerUserMessageArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetPeerUserMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPeerUserMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPeerUserMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPeerUserMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetPeerUserMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPeerUserMessageArgs_Req_DEFAULT *message.TLMessageGetPeerUserMessage

func (p *GetPeerUserMessageArgs) GetReq() *message.TLMessageGetPeerUserMessage {
	if !p.IsSetReq() {
		return GetPeerUserMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPeerUserMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPeerUserMessageResult struct {
	Success *tg.MessageBox
}

var GetPeerUserMessageResult_Success_DEFAULT *tg.MessageBox

func (p *GetPeerUserMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPeerUserMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPeerUserMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.MessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPeerUserMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPeerUserMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPeerUserMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPeerUserMessageResult) GetSuccess() *tg.MessageBox {
	if !p.IsSetSuccess() {
		return GetPeerUserMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPeerUserMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessageBox)
}

func (p *GetPeerUserMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPeerUserMessageResult) GetResult() interface{} {
	return p.Success
}

func searchByMediaTypeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SearchByMediaTypeArgs)
	realResult := result.(*SearchByMediaTypeResult)
	success, err := handler.(message.RPCMessage).MessageSearchByMediaType(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSearchByMediaTypeArgs() interface{} {
	return &SearchByMediaTypeArgs{}
}

func newSearchByMediaTypeResult() interface{} {
	return &SearchByMediaTypeResult{}
}

type SearchByMediaTypeArgs struct {
	Req *message.TLMessageSearchByMediaType
}

func (p *SearchByMediaTypeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SearchByMediaTypeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SearchByMediaTypeArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageSearchByMediaType)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SearchByMediaTypeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SearchByMediaTypeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SearchByMediaTypeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageSearchByMediaType)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SearchByMediaTypeArgs_Req_DEFAULT *message.TLMessageSearchByMediaType

func (p *SearchByMediaTypeArgs) GetReq() *message.TLMessageSearchByMediaType {
	if !p.IsSetReq() {
		return SearchByMediaTypeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchByMediaTypeArgs) IsSetReq() bool {
	return p.Req != nil
}

type SearchByMediaTypeResult struct {
	Success *message.VectorMessageBox
}

var SearchByMediaTypeResult_Success_DEFAULT *message.VectorMessageBox

func (p *SearchByMediaTypeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SearchByMediaTypeResult")
	}
	return json.Marshal(p.Success)
}

func (p *SearchByMediaTypeResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchByMediaTypeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SearchByMediaTypeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SearchByMediaTypeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchByMediaTypeResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return SearchByMediaTypeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchByMediaTypeResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *SearchByMediaTypeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchByMediaTypeResult) GetResult() interface{} {
	return p.Success
}

func searchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SearchArgs)
	realResult := result.(*SearchResult)
	success, err := handler.(message.RPCMessage).MessageSearch(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSearchArgs() interface{} {
	return &SearchArgs{}
}

func newSearchResult() interface{} {
	return &SearchResult{}
}

type SearchArgs struct {
	Req *message.TLMessageSearch
}

func (p *SearchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SearchArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SearchArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageSearch)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SearchArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SearchArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SearchArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageSearch)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SearchArgs_Req_DEFAULT *message.TLMessageSearch

func (p *SearchArgs) GetReq() *message.TLMessageSearch {
	if !p.IsSetReq() {
		return SearchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchArgs) IsSetReq() bool {
	return p.Req != nil
}

type SearchResult struct {
	Success *message.VectorMessageBox
}

var SearchResult_Success_DEFAULT *message.VectorMessageBox

func (p *SearchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SearchResult")
	}
	return json.Marshal(p.Success)
}

func (p *SearchResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SearchResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SearchResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return SearchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *SearchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchResult) GetResult() interface{} {
	return p.Success
}

func searchGlobalHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SearchGlobalArgs)
	realResult := result.(*SearchGlobalResult)
	success, err := handler.(message.RPCMessage).MessageSearchGlobal(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSearchGlobalArgs() interface{} {
	return &SearchGlobalArgs{}
}

func newSearchGlobalResult() interface{} {
	return &SearchGlobalResult{}
}

type SearchGlobalArgs struct {
	Req *message.TLMessageSearchGlobal
}

func (p *SearchGlobalArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SearchGlobalArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SearchGlobalArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageSearchGlobal)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SearchGlobalArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SearchGlobalArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SearchGlobalArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageSearchGlobal)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SearchGlobalArgs_Req_DEFAULT *message.TLMessageSearchGlobal

func (p *SearchGlobalArgs) GetReq() *message.TLMessageSearchGlobal {
	if !p.IsSetReq() {
		return SearchGlobalArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchGlobalArgs) IsSetReq() bool {
	return p.Req != nil
}

type SearchGlobalResult struct {
	Success *message.VectorMessageBox
}

var SearchGlobalResult_Success_DEFAULT *message.VectorMessageBox

func (p *SearchGlobalResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SearchGlobalResult")
	}
	return json.Marshal(p.Success)
}

func (p *SearchGlobalResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchGlobalResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SearchGlobalResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SearchGlobalResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchGlobalResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return SearchGlobalResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchGlobalResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *SearchGlobalResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchGlobalResult) GetResult() interface{} {
	return p.Success
}

func searchByPinnedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SearchByPinnedArgs)
	realResult := result.(*SearchByPinnedResult)
	success, err := handler.(message.RPCMessage).MessageSearchByPinned(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSearchByPinnedArgs() interface{} {
	return &SearchByPinnedArgs{}
}

func newSearchByPinnedResult() interface{} {
	return &SearchByPinnedResult{}
}

type SearchByPinnedArgs struct {
	Req *message.TLMessageSearchByPinned
}

func (p *SearchByPinnedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SearchByPinnedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SearchByPinnedArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageSearchByPinned)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SearchByPinnedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SearchByPinnedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SearchByPinnedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageSearchByPinned)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SearchByPinnedArgs_Req_DEFAULT *message.TLMessageSearchByPinned

func (p *SearchByPinnedArgs) GetReq() *message.TLMessageSearchByPinned {
	if !p.IsSetReq() {
		return SearchByPinnedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchByPinnedArgs) IsSetReq() bool {
	return p.Req != nil
}

type SearchByPinnedResult struct {
	Success *message.VectorMessageBox
}

var SearchByPinnedResult_Success_DEFAULT *message.VectorMessageBox

func (p *SearchByPinnedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SearchByPinnedResult")
	}
	return json.Marshal(p.Success)
}

func (p *SearchByPinnedResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchByPinnedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SearchByPinnedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SearchByPinnedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchByPinnedResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return SearchByPinnedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchByPinnedResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *SearchByPinnedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchByPinnedResult) GetResult() interface{} {
	return p.Success
}

func getSearchCounterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetSearchCounterArgs)
	realResult := result.(*GetSearchCounterResult)
	success, err := handler.(message.RPCMessage).MessageGetSearchCounter(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetSearchCounterArgs() interface{} {
	return &GetSearchCounterArgs{}
}

func newGetSearchCounterResult() interface{} {
	return &GetSearchCounterResult{}
}

type GetSearchCounterArgs struct {
	Req *message.TLMessageGetSearchCounter
}

func (p *GetSearchCounterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetSearchCounterArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetSearchCounterArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetSearchCounter)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetSearchCounterArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetSearchCounterArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetSearchCounterArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetSearchCounter)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetSearchCounterArgs_Req_DEFAULT *message.TLMessageGetSearchCounter

func (p *GetSearchCounterArgs) GetReq() *message.TLMessageGetSearchCounter {
	if !p.IsSetReq() {
		return GetSearchCounterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetSearchCounterArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetSearchCounterResult struct {
	Success *tg.Int32
}

var GetSearchCounterResult_Success_DEFAULT *tg.Int32

func (p *GetSearchCounterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetSearchCounterResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetSearchCounterResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetSearchCounterResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetSearchCounterResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetSearchCounterResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetSearchCounterResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetSearchCounterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetSearchCounterResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetSearchCounterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetSearchCounterResult) GetResult() interface{} {
	return p.Success
}

func searchV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SearchV2Args)
	realResult := result.(*SearchV2Result)
	success, err := handler.(message.RPCMessage).MessageSearchV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSearchV2Args() interface{} {
	return &SearchV2Args{}
}

func newSearchV2Result() interface{} {
	return &SearchV2Result{}
}

type SearchV2Args struct {
	Req *message.TLMessageSearchV2
}

func (p *SearchV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SearchV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *SearchV2Args) Unmarshal(in []byte) error {
	msg := new(message.TLMessageSearchV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SearchV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SearchV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *SearchV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageSearchV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SearchV2Args_Req_DEFAULT *message.TLMessageSearchV2

func (p *SearchV2Args) GetReq() *message.TLMessageSearchV2 {
	if !p.IsSetReq() {
		return SearchV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchV2Args) IsSetReq() bool {
	return p.Req != nil
}

type SearchV2Result struct {
	Success *message.VectorMessageBox
}

var SearchV2Result_Success_DEFAULT *message.VectorMessageBox

func (p *SearchV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SearchV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *SearchV2Result) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SearchV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *SearchV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchV2Result) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return SearchV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *SearchV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchV2Result) GetResult() interface{} {
	return p.Success
}

func getLastTwoPinnedMessageIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetLastTwoPinnedMessageIdArgs)
	realResult := result.(*GetLastTwoPinnedMessageIdResult)
	success, err := handler.(message.RPCMessage).MessageGetLastTwoPinnedMessageId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetLastTwoPinnedMessageIdArgs() interface{} {
	return &GetLastTwoPinnedMessageIdArgs{}
}

func newGetLastTwoPinnedMessageIdResult() interface{} {
	return &GetLastTwoPinnedMessageIdResult{}
}

type GetLastTwoPinnedMessageIdArgs struct {
	Req *message.TLMessageGetLastTwoPinnedMessageId
}

func (p *GetLastTwoPinnedMessageIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLastTwoPinnedMessageIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetLastTwoPinnedMessageIdArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetLastTwoPinnedMessageId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetLastTwoPinnedMessageIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetLastTwoPinnedMessageIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetLastTwoPinnedMessageIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetLastTwoPinnedMessageId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetLastTwoPinnedMessageIdArgs_Req_DEFAULT *message.TLMessageGetLastTwoPinnedMessageId

func (p *GetLastTwoPinnedMessageIdArgs) GetReq() *message.TLMessageGetLastTwoPinnedMessageId {
	if !p.IsSetReq() {
		return GetLastTwoPinnedMessageIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLastTwoPinnedMessageIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLastTwoPinnedMessageIdResult struct {
	Success *message.VectorInt
}

var GetLastTwoPinnedMessageIdResult_Success_DEFAULT *message.VectorInt

func (p *GetLastTwoPinnedMessageIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLastTwoPinnedMessageIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetLastTwoPinnedMessageIdResult) Unmarshal(in []byte) error {
	msg := new(message.VectorInt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLastTwoPinnedMessageIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetLastTwoPinnedMessageIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetLastTwoPinnedMessageIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorInt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLastTwoPinnedMessageIdResult) GetSuccess() *message.VectorInt {
	if !p.IsSetSuccess() {
		return GetLastTwoPinnedMessageIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLastTwoPinnedMessageIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorInt)
}

func (p *GetLastTwoPinnedMessageIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLastTwoPinnedMessageIdResult) GetResult() interface{} {
	return p.Success
}

func updatePinnedMessageIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatePinnedMessageIdArgs)
	realResult := result.(*UpdatePinnedMessageIdResult)
	success, err := handler.(message.RPCMessage).MessageUpdatePinnedMessageId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatePinnedMessageIdArgs() interface{} {
	return &UpdatePinnedMessageIdArgs{}
}

func newUpdatePinnedMessageIdResult() interface{} {
	return &UpdatePinnedMessageIdResult{}
}

type UpdatePinnedMessageIdArgs struct {
	Req *message.TLMessageUpdatePinnedMessageId
}

func (p *UpdatePinnedMessageIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatePinnedMessageIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatePinnedMessageIdArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageUpdatePinnedMessageId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatePinnedMessageIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatePinnedMessageIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatePinnedMessageIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageUpdatePinnedMessageId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatePinnedMessageIdArgs_Req_DEFAULT *message.TLMessageUpdatePinnedMessageId

func (p *UpdatePinnedMessageIdArgs) GetReq() *message.TLMessageUpdatePinnedMessageId {
	if !p.IsSetReq() {
		return UpdatePinnedMessageIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePinnedMessageIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatePinnedMessageIdResult struct {
	Success *tg.Bool
}

var UpdatePinnedMessageIdResult_Success_DEFAULT *tg.Bool

func (p *UpdatePinnedMessageIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatePinnedMessageIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatePinnedMessageIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePinnedMessageIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatePinnedMessageIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatePinnedMessageIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePinnedMessageIdResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdatePinnedMessageIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePinnedMessageIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdatePinnedMessageIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePinnedMessageIdResult) GetResult() interface{} {
	return p.Success
}

func getPinnedMessageIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPinnedMessageIdListArgs)
	realResult := result.(*GetPinnedMessageIdListResult)
	success, err := handler.(message.RPCMessage).MessageGetPinnedMessageIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPinnedMessageIdListArgs() interface{} {
	return &GetPinnedMessageIdListArgs{}
}

func newGetPinnedMessageIdListResult() interface{} {
	return &GetPinnedMessageIdListResult{}
}

type GetPinnedMessageIdListArgs struct {
	Req *message.TLMessageGetPinnedMessageIdList
}

func (p *GetPinnedMessageIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPinnedMessageIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPinnedMessageIdListArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetPinnedMessageIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPinnedMessageIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPinnedMessageIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPinnedMessageIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetPinnedMessageIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPinnedMessageIdListArgs_Req_DEFAULT *message.TLMessageGetPinnedMessageIdList

func (p *GetPinnedMessageIdListArgs) GetReq() *message.TLMessageGetPinnedMessageIdList {
	if !p.IsSetReq() {
		return GetPinnedMessageIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPinnedMessageIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPinnedMessageIdListResult struct {
	Success *message.VectorInt
}

var GetPinnedMessageIdListResult_Success_DEFAULT *message.VectorInt

func (p *GetPinnedMessageIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPinnedMessageIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPinnedMessageIdListResult) Unmarshal(in []byte) error {
	msg := new(message.VectorInt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPinnedMessageIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPinnedMessageIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPinnedMessageIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorInt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPinnedMessageIdListResult) GetSuccess() *message.VectorInt {
	if !p.IsSetSuccess() {
		return GetPinnedMessageIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPinnedMessageIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorInt)
}

func (p *GetPinnedMessageIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPinnedMessageIdListResult) GetResult() interface{} {
	return p.Success
}

func unPinAllMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UnPinAllMessagesArgs)
	realResult := result.(*UnPinAllMessagesResult)
	success, err := handler.(message.RPCMessage).MessageUnPinAllMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUnPinAllMessagesArgs() interface{} {
	return &UnPinAllMessagesArgs{}
}

func newUnPinAllMessagesResult() interface{} {
	return &UnPinAllMessagesResult{}
}

type UnPinAllMessagesArgs struct {
	Req *message.TLMessageUnPinAllMessages
}

func (p *UnPinAllMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UnPinAllMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UnPinAllMessagesArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageUnPinAllMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UnPinAllMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UnPinAllMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UnPinAllMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageUnPinAllMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UnPinAllMessagesArgs_Req_DEFAULT *message.TLMessageUnPinAllMessages

func (p *UnPinAllMessagesArgs) GetReq() *message.TLMessageUnPinAllMessages {
	if !p.IsSetReq() {
		return UnPinAllMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UnPinAllMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type UnPinAllMessagesResult struct {
	Success *message.VectorInt
}

var UnPinAllMessagesResult_Success_DEFAULT *message.VectorInt

func (p *UnPinAllMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UnPinAllMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *UnPinAllMessagesResult) Unmarshal(in []byte) error {
	msg := new(message.VectorInt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnPinAllMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UnPinAllMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UnPinAllMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorInt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnPinAllMessagesResult) GetSuccess() *message.VectorInt {
	if !p.IsSetSuccess() {
		return UnPinAllMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UnPinAllMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorInt)
}

func (p *UnPinAllMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UnPinAllMessagesResult) GetResult() interface{} {
	return p.Success
}

func getUnreadMentionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUnreadMentionsArgs)
	realResult := result.(*GetUnreadMentionsResult)
	success, err := handler.(message.RPCMessage).MessageGetUnreadMentions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUnreadMentionsArgs() interface{} {
	return &GetUnreadMentionsArgs{}
}

func newGetUnreadMentionsResult() interface{} {
	return &GetUnreadMentionsResult{}
}

type GetUnreadMentionsArgs struct {
	Req *message.TLMessageGetUnreadMentions
}

func (p *GetUnreadMentionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUnreadMentionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUnreadMentionsArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetUnreadMentions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUnreadMentionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUnreadMentionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUnreadMentionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetUnreadMentions)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUnreadMentionsArgs_Req_DEFAULT *message.TLMessageGetUnreadMentions

func (p *GetUnreadMentionsArgs) GetReq() *message.TLMessageGetUnreadMentions {
	if !p.IsSetReq() {
		return GetUnreadMentionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUnreadMentionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUnreadMentionsResult struct {
	Success *message.VectorMessageBox
}

var GetUnreadMentionsResult_Success_DEFAULT *message.VectorMessageBox

func (p *GetUnreadMentionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUnreadMentionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUnreadMentionsResult) Unmarshal(in []byte) error {
	msg := new(message.VectorMessageBox)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUnreadMentionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUnreadMentionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUnreadMentionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(message.VectorMessageBox)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUnreadMentionsResult) GetSuccess() *message.VectorMessageBox {
	if !p.IsSetSuccess() {
		return GetUnreadMentionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUnreadMentionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.VectorMessageBox)
}

func (p *GetUnreadMentionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUnreadMentionsResult) GetResult() interface{} {
	return p.Success
}

func getUnreadMentionsCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUnreadMentionsCountArgs)
	realResult := result.(*GetUnreadMentionsCountResult)
	success, err := handler.(message.RPCMessage).MessageGetUnreadMentionsCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUnreadMentionsCountArgs() interface{} {
	return &GetUnreadMentionsCountArgs{}
}

func newGetUnreadMentionsCountResult() interface{} {
	return &GetUnreadMentionsCountResult{}
}

type GetUnreadMentionsCountArgs struct {
	Req *message.TLMessageGetUnreadMentionsCount
}

func (p *GetUnreadMentionsCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUnreadMentionsCountArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUnreadMentionsCountArgs) Unmarshal(in []byte) error {
	msg := new(message.TLMessageGetUnreadMentionsCount)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUnreadMentionsCountArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUnreadMentionsCountArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUnreadMentionsCountArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(message.TLMessageGetUnreadMentionsCount)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUnreadMentionsCountArgs_Req_DEFAULT *message.TLMessageGetUnreadMentionsCount

func (p *GetUnreadMentionsCountArgs) GetReq() *message.TLMessageGetUnreadMentionsCount {
	if !p.IsSetReq() {
		return GetUnreadMentionsCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUnreadMentionsCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUnreadMentionsCountResult struct {
	Success *tg.Int32
}

var GetUnreadMentionsCountResult_Success_DEFAULT *tg.Int32

func (p *GetUnreadMentionsCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUnreadMentionsCountResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUnreadMentionsCountResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUnreadMentionsCountResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUnreadMentionsCountResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUnreadMentionsCountResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUnreadMentionsCountResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetUnreadMentionsCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUnreadMentionsCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetUnreadMentionsCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUnreadMentionsCountResult) GetResult() interface{} {
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

func (p *kClient) MessageGetUserMessage(ctx context.Context, req *message.TLMessageGetUserMessage) (r *tg.MessageBox, err error) {
	// var _args GetUserMessageArgs
	// _args.Req = req
	// var _result GetUserMessageResult

	_result := new(tg.MessageBox)

	if err = p.c.Call(ctx, "message.getUserMessage", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetUserMessageList(ctx context.Context, req *message.TLMessageGetUserMessageList) (r *message.VectorMessageBox, err error) {
	// var _args GetUserMessageListArgs
	// _args.Req = req
	// var _result GetUserMessageListResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.getUserMessageList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetUserMessageListByDataIdList(ctx context.Context, req *message.TLMessageGetUserMessageListByDataIdList) (r *message.VectorMessageBox, err error) {
	// var _args GetUserMessageListByDataIdListArgs
	// _args.Req = req
	// var _result GetUserMessageListByDataIdListResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.getUserMessageListByDataIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, req *message.TLMessageGetUserMessageListByDataIdUserIdList) (r *message.VectorMessageBox, err error) {
	// var _args GetUserMessageListByDataIdUserIdListArgs
	// _args.Req = req
	// var _result GetUserMessageListByDataIdUserIdListResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.getUserMessageListByDataIdUserIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetHistoryMessages(ctx context.Context, req *message.TLMessageGetHistoryMessages) (r *message.VectorMessageBox, err error) {
	// var _args GetHistoryMessagesArgs
	// _args.Req = req
	// var _result GetHistoryMessagesResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.getHistoryMessages", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetHistoryMessagesCount(ctx context.Context, req *message.TLMessageGetHistoryMessagesCount) (r *tg.Int32, err error) {
	// var _args GetHistoryMessagesCountArgs
	// _args.Req = req
	// var _result GetHistoryMessagesCountResult

	_result := new(tg.Int32)

	if err = p.c.Call(ctx, "message.getHistoryMessagesCount", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetPeerUserMessageId(ctx context.Context, req *message.TLMessageGetPeerUserMessageId) (r *tg.Int32, err error) {
	// var _args GetPeerUserMessageIdArgs
	// _args.Req = req
	// var _result GetPeerUserMessageIdResult

	_result := new(tg.Int32)

	if err = p.c.Call(ctx, "message.getPeerUserMessageId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetPeerUserMessage(ctx context.Context, req *message.TLMessageGetPeerUserMessage) (r *tg.MessageBox, err error) {
	// var _args GetPeerUserMessageArgs
	// _args.Req = req
	// var _result GetPeerUserMessageResult

	_result := new(tg.MessageBox)

	if err = p.c.Call(ctx, "message.getPeerUserMessage", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageSearchByMediaType(ctx context.Context, req *message.TLMessageSearchByMediaType) (r *message.VectorMessageBox, err error) {
	// var _args SearchByMediaTypeArgs
	// _args.Req = req
	// var _result SearchByMediaTypeResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.searchByMediaType", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageSearch(ctx context.Context, req *message.TLMessageSearch) (r *message.VectorMessageBox, err error) {
	// var _args SearchArgs
	// _args.Req = req
	// var _result SearchResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.search", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageSearchGlobal(ctx context.Context, req *message.TLMessageSearchGlobal) (r *message.VectorMessageBox, err error) {
	// var _args SearchGlobalArgs
	// _args.Req = req
	// var _result SearchGlobalResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.searchGlobal", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageSearchByPinned(ctx context.Context, req *message.TLMessageSearchByPinned) (r *message.VectorMessageBox, err error) {
	// var _args SearchByPinnedArgs
	// _args.Req = req
	// var _result SearchByPinnedResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.searchByPinned", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetSearchCounter(ctx context.Context, req *message.TLMessageGetSearchCounter) (r *tg.Int32, err error) {
	// var _args GetSearchCounterArgs
	// _args.Req = req
	// var _result GetSearchCounterResult

	_result := new(tg.Int32)

	if err = p.c.Call(ctx, "message.getSearchCounter", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageSearchV2(ctx context.Context, req *message.TLMessageSearchV2) (r *message.VectorMessageBox, err error) {
	// var _args SearchV2Args
	// _args.Req = req
	// var _result SearchV2Result

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.searchV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetLastTwoPinnedMessageId(ctx context.Context, req *message.TLMessageGetLastTwoPinnedMessageId) (r *message.VectorInt, err error) {
	// var _args GetLastTwoPinnedMessageIdArgs
	// _args.Req = req
	// var _result GetLastTwoPinnedMessageIdResult

	_result := new(message.VectorInt)

	if err = p.c.Call(ctx, "message.getLastTwoPinnedMessageId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageUpdatePinnedMessageId(ctx context.Context, req *message.TLMessageUpdatePinnedMessageId) (r *tg.Bool, err error) {
	// var _args UpdatePinnedMessageIdArgs
	// _args.Req = req
	// var _result UpdatePinnedMessageIdResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "message.updatePinnedMessageId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetPinnedMessageIdList(ctx context.Context, req *message.TLMessageGetPinnedMessageIdList) (r *message.VectorInt, err error) {
	// var _args GetPinnedMessageIdListArgs
	// _args.Req = req
	// var _result GetPinnedMessageIdListResult

	_result := new(message.VectorInt)

	if err = p.c.Call(ctx, "message.getPinnedMessageIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageUnPinAllMessages(ctx context.Context, req *message.TLMessageUnPinAllMessages) (r *message.VectorInt, err error) {
	// var _args UnPinAllMessagesArgs
	// _args.Req = req
	// var _result UnPinAllMessagesResult

	_result := new(message.VectorInt)

	if err = p.c.Call(ctx, "message.unPinAllMessages", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetUnreadMentions(ctx context.Context, req *message.TLMessageGetUnreadMentions) (r *message.VectorMessageBox, err error) {
	// var _args GetUnreadMentionsArgs
	// _args.Req = req
	// var _result GetUnreadMentionsResult

	_result := new(message.VectorMessageBox)

	if err = p.c.Call(ctx, "message.getUnreadMentions", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessageGetUnreadMentionsCount(ctx context.Context, req *message.TLMessageGetUnreadMentionsCount) (r *tg.Int32, err error) {
	// var _args GetUnreadMentionsCountArgs
	// _args.Req = req
	// var _result GetUnreadMentionsCountResult

	_result := new(tg.Int32)

	if err = p.c.Call(ctx, "message.getUnreadMentionsCount", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
