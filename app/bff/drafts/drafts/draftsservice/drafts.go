/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package draftsservice

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
	"messages.saveDraft": kitex.NewMethodInfo(
		messagesSaveDraftHandler,
		newMessagesSaveDraftArgs,
		newMessagesSaveDraftResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getAllDrafts": kitex.NewMethodInfo(
		messagesGetAllDraftsHandler,
		newMessagesGetAllDraftsArgs,
		newMessagesGetAllDraftsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.clearAllDrafts": kitex.NewMethodInfo(
		messagesClearAllDraftsHandler,
		newMessagesClearAllDraftsArgs,
		newMessagesClearAllDraftsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	draftsServiceServiceInfo                = NewServiceInfo()
	draftsServiceServiceInfoForClient       = NewServiceInfoForClient()
	draftsServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCDrafts", draftsServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCDrafts", draftsServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCDrafts", draftsServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return draftsServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return draftsServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return draftsServiceServiceInfoForClient
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
	serviceName := "RPCDrafts"
	handlerType := (*tg.RPCDrafts)(nil)
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
		"PackageName": "drafts",
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

func messagesSaveDraftHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSaveDraftArgs)
	realResult := result.(*MessagesSaveDraftResult)
	success, err := handler.(tg.RPCDrafts).MessagesSaveDraft(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSaveDraftArgs() interface{} {
	return &MessagesSaveDraftArgs{}
}

func newMessagesSaveDraftResult() interface{} {
	return &MessagesSaveDraftResult{}
}

type MessagesSaveDraftArgs struct {
	Req *tg.TLMessagesSaveDraft
}

func (p *MessagesSaveDraftArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSaveDraftArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSaveDraftArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSaveDraft)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSaveDraftArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSaveDraftArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSaveDraftArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSaveDraft)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSaveDraftArgs_Req_DEFAULT *tg.TLMessagesSaveDraft

func (p *MessagesSaveDraftArgs) GetReq() *tg.TLMessagesSaveDraft {
	if !p.IsSetReq() {
		return MessagesSaveDraftArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSaveDraftArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSaveDraftResult struct {
	Success *tg.Bool
}

var MessagesSaveDraftResult_Success_DEFAULT *tg.Bool

func (p *MessagesSaveDraftResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSaveDraftResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSaveDraftResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSaveDraftResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSaveDraftResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSaveDraftResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSaveDraftResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesSaveDraftResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSaveDraftResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesSaveDraftResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSaveDraftResult) GetResult() interface{} {
	return p.Success
}

func messagesGetAllDraftsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetAllDraftsArgs)
	realResult := result.(*MessagesGetAllDraftsResult)
	success, err := handler.(tg.RPCDrafts).MessagesGetAllDrafts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetAllDraftsArgs() interface{} {
	return &MessagesGetAllDraftsArgs{}
}

func newMessagesGetAllDraftsResult() interface{} {
	return &MessagesGetAllDraftsResult{}
}

type MessagesGetAllDraftsArgs struct {
	Req *tg.TLMessagesGetAllDrafts
}

func (p *MessagesGetAllDraftsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetAllDraftsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetAllDraftsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetAllDrafts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetAllDraftsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetAllDraftsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetAllDraftsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetAllDrafts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetAllDraftsArgs_Req_DEFAULT *tg.TLMessagesGetAllDrafts

func (p *MessagesGetAllDraftsArgs) GetReq() *tg.TLMessagesGetAllDrafts {
	if !p.IsSetReq() {
		return MessagesGetAllDraftsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetAllDraftsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetAllDraftsResult struct {
	Success *tg.Updates
}

var MessagesGetAllDraftsResult_Success_DEFAULT *tg.Updates

func (p *MessagesGetAllDraftsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetAllDraftsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetAllDraftsResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetAllDraftsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetAllDraftsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetAllDraftsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetAllDraftsResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesGetAllDraftsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetAllDraftsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesGetAllDraftsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetAllDraftsResult) GetResult() interface{} {
	return p.Success
}

func messagesClearAllDraftsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesClearAllDraftsArgs)
	realResult := result.(*MessagesClearAllDraftsResult)
	success, err := handler.(tg.RPCDrafts).MessagesClearAllDrafts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesClearAllDraftsArgs() interface{} {
	return &MessagesClearAllDraftsArgs{}
}

func newMessagesClearAllDraftsResult() interface{} {
	return &MessagesClearAllDraftsResult{}
}

type MessagesClearAllDraftsArgs struct {
	Req *tg.TLMessagesClearAllDrafts
}

func (p *MessagesClearAllDraftsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesClearAllDraftsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesClearAllDraftsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesClearAllDrafts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesClearAllDraftsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesClearAllDraftsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesClearAllDraftsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesClearAllDrafts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesClearAllDraftsArgs_Req_DEFAULT *tg.TLMessagesClearAllDrafts

func (p *MessagesClearAllDraftsArgs) GetReq() *tg.TLMessagesClearAllDrafts {
	if !p.IsSetReq() {
		return MessagesClearAllDraftsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesClearAllDraftsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesClearAllDraftsResult struct {
	Success *tg.Bool
}

var MessagesClearAllDraftsResult_Success_DEFAULT *tg.Bool

func (p *MessagesClearAllDraftsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesClearAllDraftsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesClearAllDraftsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesClearAllDraftsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesClearAllDraftsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesClearAllDraftsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesClearAllDraftsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesClearAllDraftsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesClearAllDraftsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesClearAllDraftsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesClearAllDraftsResult) GetResult() interface{} {
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

func (p *kClient) MessagesSaveDraft(ctx context.Context, req *tg.TLMessagesSaveDraft) (r *tg.Bool, err error) {
	// var _args MessagesSaveDraftArgs
	// _args.Req = req
	// var _result MessagesSaveDraftResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "messages.saveDraft", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetAllDrafts(ctx context.Context, req *tg.TLMessagesGetAllDrafts) (r *tg.Updates, err error) {
	// var _args MessagesGetAllDraftsArgs
	// _args.Req = req
	// var _result MessagesGetAllDraftsResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "messages.getAllDrafts", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesClearAllDrafts(ctx context.Context, req *tg.TLMessagesClearAllDrafts) (r *tg.Bool, err error) {
	// var _args MessagesClearAllDraftsArgs
	// _args.Req = req
	// var _result MessagesClearAllDraftsResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "messages.clearAllDrafts", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
