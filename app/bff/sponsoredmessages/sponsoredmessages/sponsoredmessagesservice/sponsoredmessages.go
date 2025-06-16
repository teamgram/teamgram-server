/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package sponsoredmessagesservice

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
	"account.toggleSponsoredMessages": kitex.NewMethodInfo(
		accountToggleSponsoredMessagesHandler,
		newAccountToggleSponsoredMessagesArgs,
		newAccountToggleSponsoredMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.viewSponsoredMessage": kitex.NewMethodInfo(
		messagesViewSponsoredMessageHandler,
		newMessagesViewSponsoredMessageArgs,
		newMessagesViewSponsoredMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.clickSponsoredMessage": kitex.NewMethodInfo(
		messagesClickSponsoredMessageHandler,
		newMessagesClickSponsoredMessageArgs,
		newMessagesClickSponsoredMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.reportSponsoredMessage": kitex.NewMethodInfo(
		messagesReportSponsoredMessageHandler,
		newMessagesReportSponsoredMessageArgs,
		newMessagesReportSponsoredMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getSponsoredMessages": kitex.NewMethodInfo(
		messagesGetSponsoredMessagesHandler,
		newMessagesGetSponsoredMessagesArgs,
		newMessagesGetSponsoredMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.restrictSponsoredMessages": kitex.NewMethodInfo(
		channelsRestrictSponsoredMessagesHandler,
		newChannelsRestrictSponsoredMessagesArgs,
		newChannelsRestrictSponsoredMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.viewSponsoredMessage": kitex.NewMethodInfo(
		channelsViewSponsoredMessageHandler,
		newChannelsViewSponsoredMessageArgs,
		newChannelsViewSponsoredMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.getSponsoredMessages": kitex.NewMethodInfo(
		channelsGetSponsoredMessagesHandler,
		newChannelsGetSponsoredMessagesArgs,
		newChannelsGetSponsoredMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.clickSponsoredMessage": kitex.NewMethodInfo(
		channelsClickSponsoredMessageHandler,
		newChannelsClickSponsoredMessageArgs,
		newChannelsClickSponsoredMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.reportSponsoredMessage": kitex.NewMethodInfo(
		channelsReportSponsoredMessageHandler,
		newChannelsReportSponsoredMessageArgs,
		newChannelsReportSponsoredMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	sponsoredmessagesServiceServiceInfo                = NewServiceInfo()
	sponsoredmessagesServiceServiceInfoForClient       = NewServiceInfoForClient()
	sponsoredmessagesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return sponsoredmessagesServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return sponsoredmessagesServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return sponsoredmessagesServiceServiceInfoForClient
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
	serviceName := "RPCSponsoredMessages"
	handlerType := (*tg.RPCSponsoredMessages)(nil)
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
		"PackageName": "sponsoredmessages",
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

func accountToggleSponsoredMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountToggleSponsoredMessagesArgs)
	realResult := result.(*AccountToggleSponsoredMessagesResult)
	success, err := handler.(tg.RPCSponsoredMessages).AccountToggleSponsoredMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountToggleSponsoredMessagesArgs() interface{} {
	return &AccountToggleSponsoredMessagesArgs{}
}

func newAccountToggleSponsoredMessagesResult() interface{} {
	return &AccountToggleSponsoredMessagesResult{}
}

type AccountToggleSponsoredMessagesArgs struct {
	Req *tg.TLAccountToggleSponsoredMessages
}

func (p *AccountToggleSponsoredMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountToggleSponsoredMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountToggleSponsoredMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountToggleSponsoredMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountToggleSponsoredMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountToggleSponsoredMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountToggleSponsoredMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountToggleSponsoredMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountToggleSponsoredMessagesArgs_Req_DEFAULT *tg.TLAccountToggleSponsoredMessages

func (p *AccountToggleSponsoredMessagesArgs) GetReq() *tg.TLAccountToggleSponsoredMessages {
	if !p.IsSetReq() {
		return AccountToggleSponsoredMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountToggleSponsoredMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountToggleSponsoredMessagesResult struct {
	Success *tg.Bool
}

var AccountToggleSponsoredMessagesResult_Success_DEFAULT *tg.Bool

func (p *AccountToggleSponsoredMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountToggleSponsoredMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountToggleSponsoredMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountToggleSponsoredMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountToggleSponsoredMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountToggleSponsoredMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountToggleSponsoredMessagesResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountToggleSponsoredMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountToggleSponsoredMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountToggleSponsoredMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountToggleSponsoredMessagesResult) GetResult() interface{} {
	return p.Success
}

func messagesViewSponsoredMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesViewSponsoredMessageArgs)
	realResult := result.(*MessagesViewSponsoredMessageResult)
	success, err := handler.(tg.RPCSponsoredMessages).MessagesViewSponsoredMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesViewSponsoredMessageArgs() interface{} {
	return &MessagesViewSponsoredMessageArgs{}
}

func newMessagesViewSponsoredMessageResult() interface{} {
	return &MessagesViewSponsoredMessageResult{}
}

type MessagesViewSponsoredMessageArgs struct {
	Req *tg.TLMessagesViewSponsoredMessage
}

func (p *MessagesViewSponsoredMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesViewSponsoredMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesViewSponsoredMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesViewSponsoredMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesViewSponsoredMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesViewSponsoredMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesViewSponsoredMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesViewSponsoredMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesViewSponsoredMessageArgs_Req_DEFAULT *tg.TLMessagesViewSponsoredMessage

func (p *MessagesViewSponsoredMessageArgs) GetReq() *tg.TLMessagesViewSponsoredMessage {
	if !p.IsSetReq() {
		return MessagesViewSponsoredMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesViewSponsoredMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesViewSponsoredMessageResult struct {
	Success *tg.Bool
}

var MessagesViewSponsoredMessageResult_Success_DEFAULT *tg.Bool

func (p *MessagesViewSponsoredMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesViewSponsoredMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesViewSponsoredMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesViewSponsoredMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesViewSponsoredMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesViewSponsoredMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesViewSponsoredMessageResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesViewSponsoredMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesViewSponsoredMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesViewSponsoredMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesViewSponsoredMessageResult) GetResult() interface{} {
	return p.Success
}

func messagesClickSponsoredMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesClickSponsoredMessageArgs)
	realResult := result.(*MessagesClickSponsoredMessageResult)
	success, err := handler.(tg.RPCSponsoredMessages).MessagesClickSponsoredMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesClickSponsoredMessageArgs() interface{} {
	return &MessagesClickSponsoredMessageArgs{}
}

func newMessagesClickSponsoredMessageResult() interface{} {
	return &MessagesClickSponsoredMessageResult{}
}

type MessagesClickSponsoredMessageArgs struct {
	Req *tg.TLMessagesClickSponsoredMessage
}

func (p *MessagesClickSponsoredMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesClickSponsoredMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesClickSponsoredMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesClickSponsoredMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesClickSponsoredMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesClickSponsoredMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesClickSponsoredMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesClickSponsoredMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesClickSponsoredMessageArgs_Req_DEFAULT *tg.TLMessagesClickSponsoredMessage

func (p *MessagesClickSponsoredMessageArgs) GetReq() *tg.TLMessagesClickSponsoredMessage {
	if !p.IsSetReq() {
		return MessagesClickSponsoredMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesClickSponsoredMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesClickSponsoredMessageResult struct {
	Success *tg.Bool
}

var MessagesClickSponsoredMessageResult_Success_DEFAULT *tg.Bool

func (p *MessagesClickSponsoredMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesClickSponsoredMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesClickSponsoredMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesClickSponsoredMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesClickSponsoredMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesClickSponsoredMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesClickSponsoredMessageResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesClickSponsoredMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesClickSponsoredMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesClickSponsoredMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesClickSponsoredMessageResult) GetResult() interface{} {
	return p.Success
}

func messagesReportSponsoredMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReportSponsoredMessageArgs)
	realResult := result.(*MessagesReportSponsoredMessageResult)
	success, err := handler.(tg.RPCSponsoredMessages).MessagesReportSponsoredMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReportSponsoredMessageArgs() interface{} {
	return &MessagesReportSponsoredMessageArgs{}
}

func newMessagesReportSponsoredMessageResult() interface{} {
	return &MessagesReportSponsoredMessageResult{}
}

type MessagesReportSponsoredMessageArgs struct {
	Req *tg.TLMessagesReportSponsoredMessage
}

func (p *MessagesReportSponsoredMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReportSponsoredMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReportSponsoredMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReportSponsoredMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReportSponsoredMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReportSponsoredMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReportSponsoredMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReportSponsoredMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesReportSponsoredMessageArgs_Req_DEFAULT *tg.TLMessagesReportSponsoredMessage

func (p *MessagesReportSponsoredMessageArgs) GetReq() *tg.TLMessagesReportSponsoredMessage {
	if !p.IsSetReq() {
		return MessagesReportSponsoredMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReportSponsoredMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReportSponsoredMessageResult struct {
	Success *tg.ChannelsSponsoredMessageReportResult
}

var MessagesReportSponsoredMessageResult_Success_DEFAULT *tg.ChannelsSponsoredMessageReportResult

func (p *MessagesReportSponsoredMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReportSponsoredMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReportSponsoredMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.ChannelsSponsoredMessageReportResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReportSponsoredMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReportSponsoredMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReportSponsoredMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ChannelsSponsoredMessageReportResult)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReportSponsoredMessageResult) GetSuccess() *tg.ChannelsSponsoredMessageReportResult {
	if !p.IsSetSuccess() {
		return MessagesReportSponsoredMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReportSponsoredMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ChannelsSponsoredMessageReportResult)
}

func (p *MessagesReportSponsoredMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReportSponsoredMessageResult) GetResult() interface{} {
	return p.Success
}

func messagesGetSponsoredMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetSponsoredMessagesArgs)
	realResult := result.(*MessagesGetSponsoredMessagesResult)
	success, err := handler.(tg.RPCSponsoredMessages).MessagesGetSponsoredMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetSponsoredMessagesArgs() interface{} {
	return &MessagesGetSponsoredMessagesArgs{}
}

func newMessagesGetSponsoredMessagesResult() interface{} {
	return &MessagesGetSponsoredMessagesResult{}
}

type MessagesGetSponsoredMessagesArgs struct {
	Req *tg.TLMessagesGetSponsoredMessages
}

func (p *MessagesGetSponsoredMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetSponsoredMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetSponsoredMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetSponsoredMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetSponsoredMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetSponsoredMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetSponsoredMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetSponsoredMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetSponsoredMessagesArgs_Req_DEFAULT *tg.TLMessagesGetSponsoredMessages

func (p *MessagesGetSponsoredMessagesArgs) GetReq() *tg.TLMessagesGetSponsoredMessages {
	if !p.IsSetReq() {
		return MessagesGetSponsoredMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetSponsoredMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetSponsoredMessagesResult struct {
	Success *tg.MessagesSponsoredMessages
}

var MessagesGetSponsoredMessagesResult_Success_DEFAULT *tg.MessagesSponsoredMessages

func (p *MessagesGetSponsoredMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetSponsoredMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetSponsoredMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesSponsoredMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSponsoredMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetSponsoredMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetSponsoredMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesSponsoredMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSponsoredMessagesResult) GetSuccess() *tg.MessagesSponsoredMessages {
	if !p.IsSetSuccess() {
		return MessagesGetSponsoredMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetSponsoredMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesSponsoredMessages)
}

func (p *MessagesGetSponsoredMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetSponsoredMessagesResult) GetResult() interface{} {
	return p.Success
}

func channelsRestrictSponsoredMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsRestrictSponsoredMessagesArgs)
	realResult := result.(*ChannelsRestrictSponsoredMessagesResult)
	success, err := handler.(tg.RPCSponsoredMessages).ChannelsRestrictSponsoredMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsRestrictSponsoredMessagesArgs() interface{} {
	return &ChannelsRestrictSponsoredMessagesArgs{}
}

func newChannelsRestrictSponsoredMessagesResult() interface{} {
	return &ChannelsRestrictSponsoredMessagesResult{}
}

type ChannelsRestrictSponsoredMessagesArgs struct {
	Req *tg.TLChannelsRestrictSponsoredMessages
}

func (p *ChannelsRestrictSponsoredMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsRestrictSponsoredMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsRestrictSponsoredMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsRestrictSponsoredMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsRestrictSponsoredMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsRestrictSponsoredMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsRestrictSponsoredMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsRestrictSponsoredMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsRestrictSponsoredMessagesArgs_Req_DEFAULT *tg.TLChannelsRestrictSponsoredMessages

func (p *ChannelsRestrictSponsoredMessagesArgs) GetReq() *tg.TLChannelsRestrictSponsoredMessages {
	if !p.IsSetReq() {
		return ChannelsRestrictSponsoredMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsRestrictSponsoredMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsRestrictSponsoredMessagesResult struct {
	Success *tg.Updates
}

var ChannelsRestrictSponsoredMessagesResult_Success_DEFAULT *tg.Updates

func (p *ChannelsRestrictSponsoredMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsRestrictSponsoredMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsRestrictSponsoredMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsRestrictSponsoredMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsRestrictSponsoredMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsRestrictSponsoredMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsRestrictSponsoredMessagesResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return ChannelsRestrictSponsoredMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsRestrictSponsoredMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *ChannelsRestrictSponsoredMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsRestrictSponsoredMessagesResult) GetResult() interface{} {
	return p.Success
}

func channelsViewSponsoredMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsViewSponsoredMessageArgs)
	realResult := result.(*ChannelsViewSponsoredMessageResult)
	success, err := handler.(tg.RPCSponsoredMessages).ChannelsViewSponsoredMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsViewSponsoredMessageArgs() interface{} {
	return &ChannelsViewSponsoredMessageArgs{}
}

func newChannelsViewSponsoredMessageResult() interface{} {
	return &ChannelsViewSponsoredMessageResult{}
}

type ChannelsViewSponsoredMessageArgs struct {
	Req *tg.TLChannelsViewSponsoredMessage
}

func (p *ChannelsViewSponsoredMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsViewSponsoredMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsViewSponsoredMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsViewSponsoredMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsViewSponsoredMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsViewSponsoredMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsViewSponsoredMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsViewSponsoredMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsViewSponsoredMessageArgs_Req_DEFAULT *tg.TLChannelsViewSponsoredMessage

func (p *ChannelsViewSponsoredMessageArgs) GetReq() *tg.TLChannelsViewSponsoredMessage {
	if !p.IsSetReq() {
		return ChannelsViewSponsoredMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsViewSponsoredMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsViewSponsoredMessageResult struct {
	Success *tg.Bool
}

var ChannelsViewSponsoredMessageResult_Success_DEFAULT *tg.Bool

func (p *ChannelsViewSponsoredMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsViewSponsoredMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsViewSponsoredMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsViewSponsoredMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsViewSponsoredMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsViewSponsoredMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsViewSponsoredMessageResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ChannelsViewSponsoredMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsViewSponsoredMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ChannelsViewSponsoredMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsViewSponsoredMessageResult) GetResult() interface{} {
	return p.Success
}

func channelsGetSponsoredMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsGetSponsoredMessagesArgs)
	realResult := result.(*ChannelsGetSponsoredMessagesResult)
	success, err := handler.(tg.RPCSponsoredMessages).ChannelsGetSponsoredMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsGetSponsoredMessagesArgs() interface{} {
	return &ChannelsGetSponsoredMessagesArgs{}
}

func newChannelsGetSponsoredMessagesResult() interface{} {
	return &ChannelsGetSponsoredMessagesResult{}
}

type ChannelsGetSponsoredMessagesArgs struct {
	Req *tg.TLChannelsGetSponsoredMessages
}

func (p *ChannelsGetSponsoredMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsGetSponsoredMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsGetSponsoredMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsGetSponsoredMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsGetSponsoredMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsGetSponsoredMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsGetSponsoredMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsGetSponsoredMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsGetSponsoredMessagesArgs_Req_DEFAULT *tg.TLChannelsGetSponsoredMessages

func (p *ChannelsGetSponsoredMessagesArgs) GetReq() *tg.TLChannelsGetSponsoredMessages {
	if !p.IsSetReq() {
		return ChannelsGetSponsoredMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsGetSponsoredMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsGetSponsoredMessagesResult struct {
	Success *tg.MessagesSponsoredMessages
}

var ChannelsGetSponsoredMessagesResult_Success_DEFAULT *tg.MessagesSponsoredMessages

func (p *ChannelsGetSponsoredMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsGetSponsoredMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsGetSponsoredMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesSponsoredMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsGetSponsoredMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsGetSponsoredMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsGetSponsoredMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesSponsoredMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsGetSponsoredMessagesResult) GetSuccess() *tg.MessagesSponsoredMessages {
	if !p.IsSetSuccess() {
		return ChannelsGetSponsoredMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsGetSponsoredMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesSponsoredMessages)
}

func (p *ChannelsGetSponsoredMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsGetSponsoredMessagesResult) GetResult() interface{} {
	return p.Success
}

func channelsClickSponsoredMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsClickSponsoredMessageArgs)
	realResult := result.(*ChannelsClickSponsoredMessageResult)
	success, err := handler.(tg.RPCSponsoredMessages).ChannelsClickSponsoredMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsClickSponsoredMessageArgs() interface{} {
	return &ChannelsClickSponsoredMessageArgs{}
}

func newChannelsClickSponsoredMessageResult() interface{} {
	return &ChannelsClickSponsoredMessageResult{}
}

type ChannelsClickSponsoredMessageArgs struct {
	Req *tg.TLChannelsClickSponsoredMessage
}

func (p *ChannelsClickSponsoredMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsClickSponsoredMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsClickSponsoredMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsClickSponsoredMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsClickSponsoredMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsClickSponsoredMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsClickSponsoredMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsClickSponsoredMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsClickSponsoredMessageArgs_Req_DEFAULT *tg.TLChannelsClickSponsoredMessage

func (p *ChannelsClickSponsoredMessageArgs) GetReq() *tg.TLChannelsClickSponsoredMessage {
	if !p.IsSetReq() {
		return ChannelsClickSponsoredMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsClickSponsoredMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsClickSponsoredMessageResult struct {
	Success *tg.Bool
}

var ChannelsClickSponsoredMessageResult_Success_DEFAULT *tg.Bool

func (p *ChannelsClickSponsoredMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsClickSponsoredMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsClickSponsoredMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsClickSponsoredMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsClickSponsoredMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsClickSponsoredMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsClickSponsoredMessageResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ChannelsClickSponsoredMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsClickSponsoredMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ChannelsClickSponsoredMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsClickSponsoredMessageResult) GetResult() interface{} {
	return p.Success
}

func channelsReportSponsoredMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsReportSponsoredMessageArgs)
	realResult := result.(*ChannelsReportSponsoredMessageResult)
	success, err := handler.(tg.RPCSponsoredMessages).ChannelsReportSponsoredMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsReportSponsoredMessageArgs() interface{} {
	return &ChannelsReportSponsoredMessageArgs{}
}

func newChannelsReportSponsoredMessageResult() interface{} {
	return &ChannelsReportSponsoredMessageResult{}
}

type ChannelsReportSponsoredMessageArgs struct {
	Req *tg.TLChannelsReportSponsoredMessage
}

func (p *ChannelsReportSponsoredMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsReportSponsoredMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsReportSponsoredMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsReportSponsoredMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsReportSponsoredMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsReportSponsoredMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsReportSponsoredMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsReportSponsoredMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsReportSponsoredMessageArgs_Req_DEFAULT *tg.TLChannelsReportSponsoredMessage

func (p *ChannelsReportSponsoredMessageArgs) GetReq() *tg.TLChannelsReportSponsoredMessage {
	if !p.IsSetReq() {
		return ChannelsReportSponsoredMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsReportSponsoredMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsReportSponsoredMessageResult struct {
	Success *tg.ChannelsSponsoredMessageReportResult
}

var ChannelsReportSponsoredMessageResult_Success_DEFAULT *tg.ChannelsSponsoredMessageReportResult

func (p *ChannelsReportSponsoredMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsReportSponsoredMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsReportSponsoredMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.ChannelsSponsoredMessageReportResult)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsReportSponsoredMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsReportSponsoredMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsReportSponsoredMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ChannelsSponsoredMessageReportResult)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsReportSponsoredMessageResult) GetSuccess() *tg.ChannelsSponsoredMessageReportResult {
	if !p.IsSetSuccess() {
		return ChannelsReportSponsoredMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsReportSponsoredMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ChannelsSponsoredMessageReportResult)
}

func (p *ChannelsReportSponsoredMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsReportSponsoredMessageResult) GetResult() interface{} {
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

func (p *kClient) AccountToggleSponsoredMessages(ctx context.Context, req *tg.TLAccountToggleSponsoredMessages) (r *tg.Bool, err error) {
	var _args AccountToggleSponsoredMessagesArgs
	_args.Req = req
	var _result AccountToggleSponsoredMessagesResult
	if err = p.c.Call(ctx, "account.toggleSponsoredMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesViewSponsoredMessage(ctx context.Context, req *tg.TLMessagesViewSponsoredMessage) (r *tg.Bool, err error) {
	var _args MessagesViewSponsoredMessageArgs
	_args.Req = req
	var _result MessagesViewSponsoredMessageResult
	if err = p.c.Call(ctx, "messages.viewSponsoredMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesClickSponsoredMessage(ctx context.Context, req *tg.TLMessagesClickSponsoredMessage) (r *tg.Bool, err error) {
	var _args MessagesClickSponsoredMessageArgs
	_args.Req = req
	var _result MessagesClickSponsoredMessageResult
	if err = p.c.Call(ctx, "messages.clickSponsoredMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesReportSponsoredMessage(ctx context.Context, req *tg.TLMessagesReportSponsoredMessage) (r *tg.ChannelsSponsoredMessageReportResult, err error) {
	var _args MessagesReportSponsoredMessageArgs
	_args.Req = req
	var _result MessagesReportSponsoredMessageResult
	if err = p.c.Call(ctx, "messages.reportSponsoredMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetSponsoredMessages(ctx context.Context, req *tg.TLMessagesGetSponsoredMessages) (r *tg.MessagesSponsoredMessages, err error) {
	var _args MessagesGetSponsoredMessagesArgs
	_args.Req = req
	var _result MessagesGetSponsoredMessagesResult
	if err = p.c.Call(ctx, "messages.getSponsoredMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsRestrictSponsoredMessages(ctx context.Context, req *tg.TLChannelsRestrictSponsoredMessages) (r *tg.Updates, err error) {
	var _args ChannelsRestrictSponsoredMessagesArgs
	_args.Req = req
	var _result ChannelsRestrictSponsoredMessagesResult
	if err = p.c.Call(ctx, "channels.restrictSponsoredMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsViewSponsoredMessage(ctx context.Context, req *tg.TLChannelsViewSponsoredMessage) (r *tg.Bool, err error) {
	var _args ChannelsViewSponsoredMessageArgs
	_args.Req = req
	var _result ChannelsViewSponsoredMessageResult
	if err = p.c.Call(ctx, "channels.viewSponsoredMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsGetSponsoredMessages(ctx context.Context, req *tg.TLChannelsGetSponsoredMessages) (r *tg.MessagesSponsoredMessages, err error) {
	var _args ChannelsGetSponsoredMessagesArgs
	_args.Req = req
	var _result ChannelsGetSponsoredMessagesResult
	if err = p.c.Call(ctx, "channels.getSponsoredMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsClickSponsoredMessage(ctx context.Context, req *tg.TLChannelsClickSponsoredMessage) (r *tg.Bool, err error) {
	var _args ChannelsClickSponsoredMessageArgs
	_args.Req = req
	var _result ChannelsClickSponsoredMessageResult
	if err = p.c.Call(ctx, "channels.clickSponsoredMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsReportSponsoredMessage(ctx context.Context, req *tg.TLChannelsReportSponsoredMessage) (r *tg.ChannelsSponsoredMessageReportResult, err error) {
	var _args ChannelsReportSponsoredMessageArgs
	_args.Req = req
	var _result ChannelsReportSponsoredMessageResult
	if err = p.c.Call(ctx, "channels.reportSponsoredMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
