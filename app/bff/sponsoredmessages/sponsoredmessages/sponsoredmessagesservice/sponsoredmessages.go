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
	"contacts.getSponsoredPeers": kitex.NewMethodInfo(
		contactsGetSponsoredPeersHandler,
		newContactsGetSponsoredPeersArgs,
		newContactsGetSponsoredPeersResult,
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

func contactsGetSponsoredPeersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetSponsoredPeersArgs)
	realResult := result.(*ContactsGetSponsoredPeersResult)
	success, err := handler.(tg.RPCSponsoredMessages).ContactsGetSponsoredPeers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetSponsoredPeersArgs() interface{} {
	return &ContactsGetSponsoredPeersArgs{}
}

func newContactsGetSponsoredPeersResult() interface{} {
	return &ContactsGetSponsoredPeersResult{}
}

type ContactsGetSponsoredPeersArgs struct {
	Req *tg.TLContactsGetSponsoredPeers
}

func (p *ContactsGetSponsoredPeersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetSponsoredPeersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetSponsoredPeersArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetSponsoredPeers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetSponsoredPeersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetSponsoredPeersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetSponsoredPeersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetSponsoredPeers)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetSponsoredPeersArgs_Req_DEFAULT *tg.TLContactsGetSponsoredPeers

func (p *ContactsGetSponsoredPeersArgs) GetReq() *tg.TLContactsGetSponsoredPeers {
	if !p.IsSetReq() {
		return ContactsGetSponsoredPeersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetSponsoredPeersArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetSponsoredPeersResult struct {
	Success *tg.ContactsSponsoredPeers
}

var ContactsGetSponsoredPeersResult_Success_DEFAULT *tg.ContactsSponsoredPeers

func (p *ContactsGetSponsoredPeersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetSponsoredPeersResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetSponsoredPeersResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsSponsoredPeers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetSponsoredPeersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetSponsoredPeersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetSponsoredPeersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsSponsoredPeers)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetSponsoredPeersResult) GetSuccess() *tg.ContactsSponsoredPeers {
	if !p.IsSetSuccess() {
		return ContactsGetSponsoredPeersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetSponsoredPeersResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsSponsoredPeers)
}

func (p *ContactsGetSponsoredPeersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetSponsoredPeersResult) GetResult() interface{} {
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

func (p *kClient) ContactsGetSponsoredPeers(ctx context.Context, req *tg.TLContactsGetSponsoredPeers) (r *tg.ContactsSponsoredPeers, err error) {
	var _args ContactsGetSponsoredPeersArgs
	_args.Req = req
	var _result ContactsGetSponsoredPeersResult
	if err = p.c.Call(ctx, "contacts.getSponsoredPeers", &_args, &_result); err != nil {
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
