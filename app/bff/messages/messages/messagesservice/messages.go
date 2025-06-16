/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package messagesservice

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
	"messages.getMessages": kitex.NewMethodInfo(
		messagesGetMessagesHandler,
		newMessagesGetMessagesArgs,
		newMessagesGetMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getHistory": kitex.NewMethodInfo(
		messagesGetHistoryHandler,
		newMessagesGetHistoryArgs,
		newMessagesGetHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.search": kitex.NewMethodInfo(
		messagesSearchHandler,
		newMessagesSearchArgs,
		newMessagesSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.readHistory": kitex.NewMethodInfo(
		messagesReadHistoryHandler,
		newMessagesReadHistoryArgs,
		newMessagesReadHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.deleteHistory": kitex.NewMethodInfo(
		messagesDeleteHistoryHandler,
		newMessagesDeleteHistoryArgs,
		newMessagesDeleteHistoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.deleteMessages": kitex.NewMethodInfo(
		messagesDeleteMessagesHandler,
		newMessagesDeleteMessagesArgs,
		newMessagesDeleteMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.receivedMessages": kitex.NewMethodInfo(
		messagesReceivedMessagesHandler,
		newMessagesReceivedMessagesArgs,
		newMessagesReceivedMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.sendMessage": kitex.NewMethodInfo(
		messagesSendMessageHandler,
		newMessagesSendMessageArgs,
		newMessagesSendMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.sendMedia": kitex.NewMethodInfo(
		messagesSendMediaHandler,
		newMessagesSendMediaArgs,
		newMessagesSendMediaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.forwardMessages": kitex.NewMethodInfo(
		messagesForwardMessagesHandler,
		newMessagesForwardMessagesArgs,
		newMessagesForwardMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.readMessageContents": kitex.NewMethodInfo(
		messagesReadMessageContentsHandler,
		newMessagesReadMessageContentsArgs,
		newMessagesReadMessageContentsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getMessagesViews": kitex.NewMethodInfo(
		messagesGetMessagesViewsHandler,
		newMessagesGetMessagesViewsArgs,
		newMessagesGetMessagesViewsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.searchGlobal": kitex.NewMethodInfo(
		messagesSearchGlobalHandler,
		newMessagesSearchGlobalArgs,
		newMessagesSearchGlobalResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getMessageEditData": kitex.NewMethodInfo(
		messagesGetMessageEditDataHandler,
		newMessagesGetMessageEditDataArgs,
		newMessagesGetMessageEditDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.editMessage": kitex.NewMethodInfo(
		messagesEditMessageHandler,
		newMessagesEditMessageArgs,
		newMessagesEditMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getUnreadMentions": kitex.NewMethodInfo(
		messagesGetUnreadMentionsHandler,
		newMessagesGetUnreadMentionsArgs,
		newMessagesGetUnreadMentionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.readMentions": kitex.NewMethodInfo(
		messagesReadMentionsHandler,
		newMessagesReadMentionsArgs,
		newMessagesReadMentionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getRecentLocations": kitex.NewMethodInfo(
		messagesGetRecentLocationsHandler,
		newMessagesGetRecentLocationsArgs,
		newMessagesGetRecentLocationsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.sendMultiMedia": kitex.NewMethodInfo(
		messagesSendMultiMediaHandler,
		newMessagesSendMultiMediaArgs,
		newMessagesSendMultiMediaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.updatePinnedMessage": kitex.NewMethodInfo(
		messagesUpdatePinnedMessageHandler,
		newMessagesUpdatePinnedMessageArgs,
		newMessagesUpdatePinnedMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getSearchCounters": kitex.NewMethodInfo(
		messagesGetSearchCountersHandler,
		newMessagesGetSearchCountersArgs,
		newMessagesGetSearchCountersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.unpinAllMessages": kitex.NewMethodInfo(
		messagesUnpinAllMessagesHandler,
		newMessagesUnpinAllMessagesArgs,
		newMessagesUnpinAllMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getSearchResultsCalendar": kitex.NewMethodInfo(
		messagesGetSearchResultsCalendarHandler,
		newMessagesGetSearchResultsCalendarArgs,
		newMessagesGetSearchResultsCalendarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getSearchResultsPositions": kitex.NewMethodInfo(
		messagesGetSearchResultsPositionsHandler,
		newMessagesGetSearchResultsPositionsArgs,
		newMessagesGetSearchResultsPositionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.toggleNoForwards": kitex.NewMethodInfo(
		messagesToggleNoForwardsHandler,
		newMessagesToggleNoForwardsArgs,
		newMessagesToggleNoForwardsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.saveDefaultSendAs": kitex.NewMethodInfo(
		messagesSaveDefaultSendAsHandler,
		newMessagesSaveDefaultSendAsArgs,
		newMessagesSaveDefaultSendAsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.searchSentMedia": kitex.NewMethodInfo(
		messagesSearchSentMediaHandler,
		newMessagesSearchSentMediaArgs,
		newMessagesSearchSentMediaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getOutboxReadDate": kitex.NewMethodInfo(
		messagesGetOutboxReadDateHandler,
		newMessagesGetOutboxReadDateArgs,
		newMessagesGetOutboxReadDateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.getSendAs": kitex.NewMethodInfo(
		channelsGetSendAsHandler,
		newChannelsGetSendAsArgs,
		newChannelsGetSendAsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.searchPosts": kitex.NewMethodInfo(
		channelsSearchPostsHandler,
		newChannelsSearchPostsArgs,
		newChannelsSearchPostsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	messagesServiceServiceInfo                = NewServiceInfo()
	messagesServiceServiceInfoForClient       = NewServiceInfoForClient()
	messagesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return messagesServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return messagesServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return messagesServiceServiceInfoForClient
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
	serviceName := "RPCMessages"
	handlerType := (*tg.RPCMessages)(nil)
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
		"PackageName": "messages",
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

func messagesGetMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetMessagesArgs)
	realResult := result.(*MessagesGetMessagesResult)
	success, err := handler.(tg.RPCMessages).MessagesGetMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetMessagesArgs() interface{} {
	return &MessagesGetMessagesArgs{}
}

func newMessagesGetMessagesResult() interface{} {
	return &MessagesGetMessagesResult{}
}

type MessagesGetMessagesArgs struct {
	Req *tg.TLMessagesGetMessages
}

func (p *MessagesGetMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetMessagesArgs_Req_DEFAULT *tg.TLMessagesGetMessages

func (p *MessagesGetMessagesArgs) GetReq() *tg.TLMessagesGetMessages {
	if !p.IsSetReq() {
		return MessagesGetMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetMessagesResult struct {
	Success *tg.MessagesMessages
}

var MessagesGetMessagesResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesGetMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetMessagesResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesGetMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesGetMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetMessagesResult) GetResult() interface{} {
	return p.Success
}

func messagesGetHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetHistoryArgs)
	realResult := result.(*MessagesGetHistoryResult)
	success, err := handler.(tg.RPCMessages).MessagesGetHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetHistoryArgs() interface{} {
	return &MessagesGetHistoryArgs{}
}

func newMessagesGetHistoryResult() interface{} {
	return &MessagesGetHistoryResult{}
}

type MessagesGetHistoryArgs struct {
	Req *tg.TLMessagesGetHistory
}

func (p *MessagesGetHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetHistoryArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetHistoryArgs_Req_DEFAULT *tg.TLMessagesGetHistory

func (p *MessagesGetHistoryArgs) GetReq() *tg.TLMessagesGetHistory {
	if !p.IsSetReq() {
		return MessagesGetHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetHistoryResult struct {
	Success *tg.MessagesMessages
}

var MessagesGetHistoryResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesGetHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetHistoryResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesGetHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesGetHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetHistoryResult) GetResult() interface{} {
	return p.Success
}

func messagesSearchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSearchArgs)
	realResult := result.(*MessagesSearchResult)
	success, err := handler.(tg.RPCMessages).MessagesSearch(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSearchArgs() interface{} {
	return &MessagesSearchArgs{}
}

func newMessagesSearchResult() interface{} {
	return &MessagesSearchResult{}
}

type MessagesSearchArgs struct {
	Req *tg.TLMessagesSearch
}

func (p *MessagesSearchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSearchArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSearchArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSearch)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSearchArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSearchArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSearchArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSearch)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSearchArgs_Req_DEFAULT *tg.TLMessagesSearch

func (p *MessagesSearchArgs) GetReq() *tg.TLMessagesSearch {
	if !p.IsSetReq() {
		return MessagesSearchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSearchArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSearchResult struct {
	Success *tg.MessagesMessages
}

var MessagesSearchResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesSearchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSearchResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSearchResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSearchResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSearchResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSearchResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSearchResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesSearchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSearchResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesSearchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSearchResult) GetResult() interface{} {
	return p.Success
}

func messagesReadHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReadHistoryArgs)
	realResult := result.(*MessagesReadHistoryResult)
	success, err := handler.(tg.RPCMessages).MessagesReadHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReadHistoryArgs() interface{} {
	return &MessagesReadHistoryArgs{}
}

func newMessagesReadHistoryResult() interface{} {
	return &MessagesReadHistoryResult{}
}

type MessagesReadHistoryArgs struct {
	Req *tg.TLMessagesReadHistory
}

func (p *MessagesReadHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReadHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReadHistoryArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReadHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReadHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReadHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReadHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReadHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesReadHistoryArgs_Req_DEFAULT *tg.TLMessagesReadHistory

func (p *MessagesReadHistoryArgs) GetReq() *tg.TLMessagesReadHistory {
	if !p.IsSetReq() {
		return MessagesReadHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReadHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReadHistoryResult struct {
	Success *tg.MessagesAffectedMessages
}

var MessagesReadHistoryResult_Success_DEFAULT *tg.MessagesAffectedMessages

func (p *MessagesReadHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReadHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReadHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReadHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReadHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReadHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReadHistoryResult) GetSuccess() *tg.MessagesAffectedMessages {
	if !p.IsSetSuccess() {
		return MessagesReadHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReadHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedMessages)
}

func (p *MessagesReadHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReadHistoryResult) GetResult() interface{} {
	return p.Success
}

func messagesDeleteHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesDeleteHistoryArgs)
	realResult := result.(*MessagesDeleteHistoryResult)
	success, err := handler.(tg.RPCMessages).MessagesDeleteHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesDeleteHistoryArgs() interface{} {
	return &MessagesDeleteHistoryArgs{}
}

func newMessagesDeleteHistoryResult() interface{} {
	return &MessagesDeleteHistoryResult{}
}

type MessagesDeleteHistoryArgs struct {
	Req *tg.TLMessagesDeleteHistory
}

func (p *MessagesDeleteHistoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesDeleteHistoryArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesDeleteHistoryArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesDeleteHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesDeleteHistoryArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesDeleteHistoryArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesDeleteHistoryArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesDeleteHistory)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesDeleteHistoryArgs_Req_DEFAULT *tg.TLMessagesDeleteHistory

func (p *MessagesDeleteHistoryArgs) GetReq() *tg.TLMessagesDeleteHistory {
	if !p.IsSetReq() {
		return MessagesDeleteHistoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesDeleteHistoryArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesDeleteHistoryResult struct {
	Success *tg.MessagesAffectedHistory
}

var MessagesDeleteHistoryResult_Success_DEFAULT *tg.MessagesAffectedHistory

func (p *MessagesDeleteHistoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesDeleteHistoryResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesDeleteHistoryResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteHistoryResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesDeleteHistoryResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesDeleteHistoryResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedHistory)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteHistoryResult) GetSuccess() *tg.MessagesAffectedHistory {
	if !p.IsSetSuccess() {
		return MessagesDeleteHistoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesDeleteHistoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedHistory)
}

func (p *MessagesDeleteHistoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesDeleteHistoryResult) GetResult() interface{} {
	return p.Success
}

func messagesDeleteMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesDeleteMessagesArgs)
	realResult := result.(*MessagesDeleteMessagesResult)
	success, err := handler.(tg.RPCMessages).MessagesDeleteMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesDeleteMessagesArgs() interface{} {
	return &MessagesDeleteMessagesArgs{}
}

func newMessagesDeleteMessagesResult() interface{} {
	return &MessagesDeleteMessagesResult{}
}

type MessagesDeleteMessagesArgs struct {
	Req *tg.TLMessagesDeleteMessages
}

func (p *MessagesDeleteMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesDeleteMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesDeleteMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesDeleteMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesDeleteMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesDeleteMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesDeleteMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesDeleteMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesDeleteMessagesArgs_Req_DEFAULT *tg.TLMessagesDeleteMessages

func (p *MessagesDeleteMessagesArgs) GetReq() *tg.TLMessagesDeleteMessages {
	if !p.IsSetReq() {
		return MessagesDeleteMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesDeleteMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesDeleteMessagesResult struct {
	Success *tg.MessagesAffectedMessages
}

var MessagesDeleteMessagesResult_Success_DEFAULT *tg.MessagesAffectedMessages

func (p *MessagesDeleteMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesDeleteMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesDeleteMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesDeleteMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesDeleteMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteMessagesResult) GetSuccess() *tg.MessagesAffectedMessages {
	if !p.IsSetSuccess() {
		return MessagesDeleteMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesDeleteMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedMessages)
}

func (p *MessagesDeleteMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesDeleteMessagesResult) GetResult() interface{} {
	return p.Success
}

func messagesReceivedMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReceivedMessagesArgs)
	realResult := result.(*MessagesReceivedMessagesResult)
	success, err := handler.(tg.RPCMessages).MessagesReceivedMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReceivedMessagesArgs() interface{} {
	return &MessagesReceivedMessagesArgs{}
}

func newMessagesReceivedMessagesResult() interface{} {
	return &MessagesReceivedMessagesResult{}
}

type MessagesReceivedMessagesArgs struct {
	Req *tg.TLMessagesReceivedMessages
}

func (p *MessagesReceivedMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReceivedMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReceivedMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReceivedMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReceivedMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReceivedMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReceivedMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReceivedMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesReceivedMessagesArgs_Req_DEFAULT *tg.TLMessagesReceivedMessages

func (p *MessagesReceivedMessagesArgs) GetReq() *tg.TLMessagesReceivedMessages {
	if !p.IsSetReq() {
		return MessagesReceivedMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReceivedMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReceivedMessagesResult struct {
	Success *tg.VectorReceivedNotifyMessage
}

var MessagesReceivedMessagesResult_Success_DEFAULT *tg.VectorReceivedNotifyMessage

func (p *MessagesReceivedMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReceivedMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReceivedMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorReceivedNotifyMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReceivedMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReceivedMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReceivedMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorReceivedNotifyMessage)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReceivedMessagesResult) GetSuccess() *tg.VectorReceivedNotifyMessage {
	if !p.IsSetSuccess() {
		return MessagesReceivedMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReceivedMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorReceivedNotifyMessage)
}

func (p *MessagesReceivedMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReceivedMessagesResult) GetResult() interface{} {
	return p.Success
}

func messagesSendMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSendMessageArgs)
	realResult := result.(*MessagesSendMessageResult)
	success, err := handler.(tg.RPCMessages).MessagesSendMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSendMessageArgs() interface{} {
	return &MessagesSendMessageArgs{}
}

func newMessagesSendMessageResult() interface{} {
	return &MessagesSendMessageResult{}
}

type MessagesSendMessageArgs struct {
	Req *tg.TLMessagesSendMessage
}

func (p *MessagesSendMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSendMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSendMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSendMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSendMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSendMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSendMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSendMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSendMessageArgs_Req_DEFAULT *tg.TLMessagesSendMessage

func (p *MessagesSendMessageArgs) GetReq() *tg.TLMessagesSendMessage {
	if !p.IsSetReq() {
		return MessagesSendMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSendMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSendMessageResult struct {
	Success *tg.Updates
}

var MessagesSendMessageResult_Success_DEFAULT *tg.Updates

func (p *MessagesSendMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSendMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSendMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSendMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSendMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSendMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSendMessageResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesSendMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSendMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesSendMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSendMessageResult) GetResult() interface{} {
	return p.Success
}

func messagesSendMediaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSendMediaArgs)
	realResult := result.(*MessagesSendMediaResult)
	success, err := handler.(tg.RPCMessages).MessagesSendMedia(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSendMediaArgs() interface{} {
	return &MessagesSendMediaArgs{}
}

func newMessagesSendMediaResult() interface{} {
	return &MessagesSendMediaResult{}
}

type MessagesSendMediaArgs struct {
	Req *tg.TLMessagesSendMedia
}

func (p *MessagesSendMediaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSendMediaArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSendMediaArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSendMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSendMediaArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSendMediaArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSendMediaArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSendMedia)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSendMediaArgs_Req_DEFAULT *tg.TLMessagesSendMedia

func (p *MessagesSendMediaArgs) GetReq() *tg.TLMessagesSendMedia {
	if !p.IsSetReq() {
		return MessagesSendMediaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSendMediaArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSendMediaResult struct {
	Success *tg.Updates
}

var MessagesSendMediaResult_Success_DEFAULT *tg.Updates

func (p *MessagesSendMediaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSendMediaResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSendMediaResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSendMediaResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSendMediaResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSendMediaResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSendMediaResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesSendMediaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSendMediaResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesSendMediaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSendMediaResult) GetResult() interface{} {
	return p.Success
}

func messagesForwardMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesForwardMessagesArgs)
	realResult := result.(*MessagesForwardMessagesResult)
	success, err := handler.(tg.RPCMessages).MessagesForwardMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesForwardMessagesArgs() interface{} {
	return &MessagesForwardMessagesArgs{}
}

func newMessagesForwardMessagesResult() interface{} {
	return &MessagesForwardMessagesResult{}
}

type MessagesForwardMessagesArgs struct {
	Req *tg.TLMessagesForwardMessages
}

func (p *MessagesForwardMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesForwardMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesForwardMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesForwardMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesForwardMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesForwardMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesForwardMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesForwardMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesForwardMessagesArgs_Req_DEFAULT *tg.TLMessagesForwardMessages

func (p *MessagesForwardMessagesArgs) GetReq() *tg.TLMessagesForwardMessages {
	if !p.IsSetReq() {
		return MessagesForwardMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesForwardMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesForwardMessagesResult struct {
	Success *tg.Updates
}

var MessagesForwardMessagesResult_Success_DEFAULT *tg.Updates

func (p *MessagesForwardMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesForwardMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesForwardMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesForwardMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesForwardMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesForwardMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesForwardMessagesResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesForwardMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesForwardMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesForwardMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesForwardMessagesResult) GetResult() interface{} {
	return p.Success
}

func messagesReadMessageContentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReadMessageContentsArgs)
	realResult := result.(*MessagesReadMessageContentsResult)
	success, err := handler.(tg.RPCMessages).MessagesReadMessageContents(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReadMessageContentsArgs() interface{} {
	return &MessagesReadMessageContentsArgs{}
}

func newMessagesReadMessageContentsResult() interface{} {
	return &MessagesReadMessageContentsResult{}
}

type MessagesReadMessageContentsArgs struct {
	Req *tg.TLMessagesReadMessageContents
}

func (p *MessagesReadMessageContentsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReadMessageContentsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReadMessageContentsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReadMessageContents)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReadMessageContentsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReadMessageContentsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReadMessageContentsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReadMessageContents)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesReadMessageContentsArgs_Req_DEFAULT *tg.TLMessagesReadMessageContents

func (p *MessagesReadMessageContentsArgs) GetReq() *tg.TLMessagesReadMessageContents {
	if !p.IsSetReq() {
		return MessagesReadMessageContentsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReadMessageContentsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReadMessageContentsResult struct {
	Success *tg.MessagesAffectedMessages
}

var MessagesReadMessageContentsResult_Success_DEFAULT *tg.MessagesAffectedMessages

func (p *MessagesReadMessageContentsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReadMessageContentsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReadMessageContentsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReadMessageContentsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReadMessageContentsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReadMessageContentsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReadMessageContentsResult) GetSuccess() *tg.MessagesAffectedMessages {
	if !p.IsSetSuccess() {
		return MessagesReadMessageContentsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReadMessageContentsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedMessages)
}

func (p *MessagesReadMessageContentsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReadMessageContentsResult) GetResult() interface{} {
	return p.Success
}

func messagesGetMessagesViewsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetMessagesViewsArgs)
	realResult := result.(*MessagesGetMessagesViewsResult)
	success, err := handler.(tg.RPCMessages).MessagesGetMessagesViews(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetMessagesViewsArgs() interface{} {
	return &MessagesGetMessagesViewsArgs{}
}

func newMessagesGetMessagesViewsResult() interface{} {
	return &MessagesGetMessagesViewsResult{}
}

type MessagesGetMessagesViewsArgs struct {
	Req *tg.TLMessagesGetMessagesViews
}

func (p *MessagesGetMessagesViewsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetMessagesViewsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetMessagesViewsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetMessagesViews)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetMessagesViewsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetMessagesViewsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetMessagesViewsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetMessagesViews)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetMessagesViewsArgs_Req_DEFAULT *tg.TLMessagesGetMessagesViews

func (p *MessagesGetMessagesViewsArgs) GetReq() *tg.TLMessagesGetMessagesViews {
	if !p.IsSetReq() {
		return MessagesGetMessagesViewsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetMessagesViewsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetMessagesViewsResult struct {
	Success *tg.MessagesMessageViews
}

var MessagesGetMessagesViewsResult_Success_DEFAULT *tg.MessagesMessageViews

func (p *MessagesGetMessagesViewsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetMessagesViewsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetMessagesViewsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessageViews)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetMessagesViewsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetMessagesViewsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetMessagesViewsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessageViews)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetMessagesViewsResult) GetSuccess() *tg.MessagesMessageViews {
	if !p.IsSetSuccess() {
		return MessagesGetMessagesViewsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetMessagesViewsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessageViews)
}

func (p *MessagesGetMessagesViewsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetMessagesViewsResult) GetResult() interface{} {
	return p.Success
}

func messagesSearchGlobalHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSearchGlobalArgs)
	realResult := result.(*MessagesSearchGlobalResult)
	success, err := handler.(tg.RPCMessages).MessagesSearchGlobal(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSearchGlobalArgs() interface{} {
	return &MessagesSearchGlobalArgs{}
}

func newMessagesSearchGlobalResult() interface{} {
	return &MessagesSearchGlobalResult{}
}

type MessagesSearchGlobalArgs struct {
	Req *tg.TLMessagesSearchGlobal
}

func (p *MessagesSearchGlobalArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSearchGlobalArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSearchGlobalArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSearchGlobal)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSearchGlobalArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSearchGlobalArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSearchGlobalArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSearchGlobal)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSearchGlobalArgs_Req_DEFAULT *tg.TLMessagesSearchGlobal

func (p *MessagesSearchGlobalArgs) GetReq() *tg.TLMessagesSearchGlobal {
	if !p.IsSetReq() {
		return MessagesSearchGlobalArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSearchGlobalArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSearchGlobalResult struct {
	Success *tg.MessagesMessages
}

var MessagesSearchGlobalResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesSearchGlobalResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSearchGlobalResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSearchGlobalResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSearchGlobalResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSearchGlobalResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSearchGlobalResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSearchGlobalResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesSearchGlobalResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSearchGlobalResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesSearchGlobalResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSearchGlobalResult) GetResult() interface{} {
	return p.Success
}

func messagesGetMessageEditDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetMessageEditDataArgs)
	realResult := result.(*MessagesGetMessageEditDataResult)
	success, err := handler.(tg.RPCMessages).MessagesGetMessageEditData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetMessageEditDataArgs() interface{} {
	return &MessagesGetMessageEditDataArgs{}
}

func newMessagesGetMessageEditDataResult() interface{} {
	return &MessagesGetMessageEditDataResult{}
}

type MessagesGetMessageEditDataArgs struct {
	Req *tg.TLMessagesGetMessageEditData
}

func (p *MessagesGetMessageEditDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetMessageEditDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetMessageEditDataArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetMessageEditData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetMessageEditDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetMessageEditDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetMessageEditDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetMessageEditData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetMessageEditDataArgs_Req_DEFAULT *tg.TLMessagesGetMessageEditData

func (p *MessagesGetMessageEditDataArgs) GetReq() *tg.TLMessagesGetMessageEditData {
	if !p.IsSetReq() {
		return MessagesGetMessageEditDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetMessageEditDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetMessageEditDataResult struct {
	Success *tg.MessagesMessageEditData
}

var MessagesGetMessageEditDataResult_Success_DEFAULT *tg.MessagesMessageEditData

func (p *MessagesGetMessageEditDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetMessageEditDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetMessageEditDataResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessageEditData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetMessageEditDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetMessageEditDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetMessageEditDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessageEditData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetMessageEditDataResult) GetSuccess() *tg.MessagesMessageEditData {
	if !p.IsSetSuccess() {
		return MessagesGetMessageEditDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetMessageEditDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessageEditData)
}

func (p *MessagesGetMessageEditDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetMessageEditDataResult) GetResult() interface{} {
	return p.Success
}

func messagesEditMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesEditMessageArgs)
	realResult := result.(*MessagesEditMessageResult)
	success, err := handler.(tg.RPCMessages).MessagesEditMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesEditMessageArgs() interface{} {
	return &MessagesEditMessageArgs{}
}

func newMessagesEditMessageResult() interface{} {
	return &MessagesEditMessageResult{}
}

type MessagesEditMessageArgs struct {
	Req *tg.TLMessagesEditMessage
}

func (p *MessagesEditMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesEditMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesEditMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesEditMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesEditMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesEditMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesEditMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesEditMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesEditMessageArgs_Req_DEFAULT *tg.TLMessagesEditMessage

func (p *MessagesEditMessageArgs) GetReq() *tg.TLMessagesEditMessage {
	if !p.IsSetReq() {
		return MessagesEditMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesEditMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesEditMessageResult struct {
	Success *tg.Updates
}

var MessagesEditMessageResult_Success_DEFAULT *tg.Updates

func (p *MessagesEditMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesEditMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesEditMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesEditMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesEditMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditMessageResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesEditMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesEditMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesEditMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesEditMessageResult) GetResult() interface{} {
	return p.Success
}

func messagesGetUnreadMentionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetUnreadMentionsArgs)
	realResult := result.(*MessagesGetUnreadMentionsResult)
	success, err := handler.(tg.RPCMessages).MessagesGetUnreadMentions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetUnreadMentionsArgs() interface{} {
	return &MessagesGetUnreadMentionsArgs{}
}

func newMessagesGetUnreadMentionsResult() interface{} {
	return &MessagesGetUnreadMentionsResult{}
}

type MessagesGetUnreadMentionsArgs struct {
	Req *tg.TLMessagesGetUnreadMentions
}

func (p *MessagesGetUnreadMentionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetUnreadMentionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetUnreadMentionsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetUnreadMentions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetUnreadMentionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetUnreadMentionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetUnreadMentionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetUnreadMentions)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetUnreadMentionsArgs_Req_DEFAULT *tg.TLMessagesGetUnreadMentions

func (p *MessagesGetUnreadMentionsArgs) GetReq() *tg.TLMessagesGetUnreadMentions {
	if !p.IsSetReq() {
		return MessagesGetUnreadMentionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetUnreadMentionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetUnreadMentionsResult struct {
	Success *tg.MessagesMessages
}

var MessagesGetUnreadMentionsResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesGetUnreadMentionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetUnreadMentionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetUnreadMentionsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetUnreadMentionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetUnreadMentionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetUnreadMentionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetUnreadMentionsResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesGetUnreadMentionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetUnreadMentionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesGetUnreadMentionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetUnreadMentionsResult) GetResult() interface{} {
	return p.Success
}

func messagesReadMentionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReadMentionsArgs)
	realResult := result.(*MessagesReadMentionsResult)
	success, err := handler.(tg.RPCMessages).MessagesReadMentions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReadMentionsArgs() interface{} {
	return &MessagesReadMentionsArgs{}
}

func newMessagesReadMentionsResult() interface{} {
	return &MessagesReadMentionsResult{}
}

type MessagesReadMentionsArgs struct {
	Req *tg.TLMessagesReadMentions
}

func (p *MessagesReadMentionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReadMentionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReadMentionsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReadMentions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReadMentionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReadMentionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReadMentionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReadMentions)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesReadMentionsArgs_Req_DEFAULT *tg.TLMessagesReadMentions

func (p *MessagesReadMentionsArgs) GetReq() *tg.TLMessagesReadMentions {
	if !p.IsSetReq() {
		return MessagesReadMentionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReadMentionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReadMentionsResult struct {
	Success *tg.MessagesAffectedHistory
}

var MessagesReadMentionsResult_Success_DEFAULT *tg.MessagesAffectedHistory

func (p *MessagesReadMentionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReadMentionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReadMentionsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReadMentionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReadMentionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReadMentionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedHistory)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReadMentionsResult) GetSuccess() *tg.MessagesAffectedHistory {
	if !p.IsSetSuccess() {
		return MessagesReadMentionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReadMentionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedHistory)
}

func (p *MessagesReadMentionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReadMentionsResult) GetResult() interface{} {
	return p.Success
}

func messagesGetRecentLocationsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetRecentLocationsArgs)
	realResult := result.(*MessagesGetRecentLocationsResult)
	success, err := handler.(tg.RPCMessages).MessagesGetRecentLocations(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetRecentLocationsArgs() interface{} {
	return &MessagesGetRecentLocationsArgs{}
}

func newMessagesGetRecentLocationsResult() interface{} {
	return &MessagesGetRecentLocationsResult{}
}

type MessagesGetRecentLocationsArgs struct {
	Req *tg.TLMessagesGetRecentLocations
}

func (p *MessagesGetRecentLocationsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetRecentLocationsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetRecentLocationsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetRecentLocations)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetRecentLocationsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetRecentLocationsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetRecentLocationsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetRecentLocations)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetRecentLocationsArgs_Req_DEFAULT *tg.TLMessagesGetRecentLocations

func (p *MessagesGetRecentLocationsArgs) GetReq() *tg.TLMessagesGetRecentLocations {
	if !p.IsSetReq() {
		return MessagesGetRecentLocationsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetRecentLocationsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetRecentLocationsResult struct {
	Success *tg.MessagesMessages
}

var MessagesGetRecentLocationsResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesGetRecentLocationsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetRecentLocationsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetRecentLocationsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetRecentLocationsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetRecentLocationsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetRecentLocationsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetRecentLocationsResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesGetRecentLocationsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetRecentLocationsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesGetRecentLocationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetRecentLocationsResult) GetResult() interface{} {
	return p.Success
}

func messagesSendMultiMediaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSendMultiMediaArgs)
	realResult := result.(*MessagesSendMultiMediaResult)
	success, err := handler.(tg.RPCMessages).MessagesSendMultiMedia(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSendMultiMediaArgs() interface{} {
	return &MessagesSendMultiMediaArgs{}
}

func newMessagesSendMultiMediaResult() interface{} {
	return &MessagesSendMultiMediaResult{}
}

type MessagesSendMultiMediaArgs struct {
	Req *tg.TLMessagesSendMultiMedia
}

func (p *MessagesSendMultiMediaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSendMultiMediaArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSendMultiMediaArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSendMultiMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSendMultiMediaArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSendMultiMediaArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSendMultiMediaArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSendMultiMedia)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSendMultiMediaArgs_Req_DEFAULT *tg.TLMessagesSendMultiMedia

func (p *MessagesSendMultiMediaArgs) GetReq() *tg.TLMessagesSendMultiMedia {
	if !p.IsSetReq() {
		return MessagesSendMultiMediaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSendMultiMediaArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSendMultiMediaResult struct {
	Success *tg.Updates
}

var MessagesSendMultiMediaResult_Success_DEFAULT *tg.Updates

func (p *MessagesSendMultiMediaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSendMultiMediaResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSendMultiMediaResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSendMultiMediaResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSendMultiMediaResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSendMultiMediaResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSendMultiMediaResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesSendMultiMediaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSendMultiMediaResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesSendMultiMediaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSendMultiMediaResult) GetResult() interface{} {
	return p.Success
}

func messagesUpdatePinnedMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesUpdatePinnedMessageArgs)
	realResult := result.(*MessagesUpdatePinnedMessageResult)
	success, err := handler.(tg.RPCMessages).MessagesUpdatePinnedMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesUpdatePinnedMessageArgs() interface{} {
	return &MessagesUpdatePinnedMessageArgs{}
}

func newMessagesUpdatePinnedMessageResult() interface{} {
	return &MessagesUpdatePinnedMessageResult{}
}

type MessagesUpdatePinnedMessageArgs struct {
	Req *tg.TLMessagesUpdatePinnedMessage
}

func (p *MessagesUpdatePinnedMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesUpdatePinnedMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesUpdatePinnedMessageArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesUpdatePinnedMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesUpdatePinnedMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesUpdatePinnedMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesUpdatePinnedMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesUpdatePinnedMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesUpdatePinnedMessageArgs_Req_DEFAULT *tg.TLMessagesUpdatePinnedMessage

func (p *MessagesUpdatePinnedMessageArgs) GetReq() *tg.TLMessagesUpdatePinnedMessage {
	if !p.IsSetReq() {
		return MessagesUpdatePinnedMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesUpdatePinnedMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesUpdatePinnedMessageResult struct {
	Success *tg.Updates
}

var MessagesUpdatePinnedMessageResult_Success_DEFAULT *tg.Updates

func (p *MessagesUpdatePinnedMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesUpdatePinnedMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesUpdatePinnedMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesUpdatePinnedMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesUpdatePinnedMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesUpdatePinnedMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesUpdatePinnedMessageResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesUpdatePinnedMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesUpdatePinnedMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesUpdatePinnedMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesUpdatePinnedMessageResult) GetResult() interface{} {
	return p.Success
}

func messagesGetSearchCountersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetSearchCountersArgs)
	realResult := result.(*MessagesGetSearchCountersResult)
	success, err := handler.(tg.RPCMessages).MessagesGetSearchCounters(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetSearchCountersArgs() interface{} {
	return &MessagesGetSearchCountersArgs{}
}

func newMessagesGetSearchCountersResult() interface{} {
	return &MessagesGetSearchCountersResult{}
}

type MessagesGetSearchCountersArgs struct {
	Req *tg.TLMessagesGetSearchCounters
}

func (p *MessagesGetSearchCountersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetSearchCountersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetSearchCountersArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetSearchCounters)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetSearchCountersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetSearchCountersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetSearchCountersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetSearchCounters)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetSearchCountersArgs_Req_DEFAULT *tg.TLMessagesGetSearchCounters

func (p *MessagesGetSearchCountersArgs) GetReq() *tg.TLMessagesGetSearchCounters {
	if !p.IsSetReq() {
		return MessagesGetSearchCountersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetSearchCountersArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetSearchCountersResult struct {
	Success *tg.VectorMessagesSearchCounter
}

var MessagesGetSearchCountersResult_Success_DEFAULT *tg.VectorMessagesSearchCounter

func (p *MessagesGetSearchCountersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetSearchCountersResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetSearchCountersResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorMessagesSearchCounter)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSearchCountersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetSearchCountersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetSearchCountersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorMessagesSearchCounter)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSearchCountersResult) GetSuccess() *tg.VectorMessagesSearchCounter {
	if !p.IsSetSuccess() {
		return MessagesGetSearchCountersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetSearchCountersResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorMessagesSearchCounter)
}

func (p *MessagesGetSearchCountersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetSearchCountersResult) GetResult() interface{} {
	return p.Success
}

func messagesUnpinAllMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesUnpinAllMessagesArgs)
	realResult := result.(*MessagesUnpinAllMessagesResult)
	success, err := handler.(tg.RPCMessages).MessagesUnpinAllMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesUnpinAllMessagesArgs() interface{} {
	return &MessagesUnpinAllMessagesArgs{}
}

func newMessagesUnpinAllMessagesResult() interface{} {
	return &MessagesUnpinAllMessagesResult{}
}

type MessagesUnpinAllMessagesArgs struct {
	Req *tg.TLMessagesUnpinAllMessages
}

func (p *MessagesUnpinAllMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesUnpinAllMessagesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesUnpinAllMessagesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesUnpinAllMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesUnpinAllMessagesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesUnpinAllMessagesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesUnpinAllMessagesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesUnpinAllMessages)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesUnpinAllMessagesArgs_Req_DEFAULT *tg.TLMessagesUnpinAllMessages

func (p *MessagesUnpinAllMessagesArgs) GetReq() *tg.TLMessagesUnpinAllMessages {
	if !p.IsSetReq() {
		return MessagesUnpinAllMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesUnpinAllMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesUnpinAllMessagesResult struct {
	Success *tg.MessagesAffectedHistory
}

var MessagesUnpinAllMessagesResult_Success_DEFAULT *tg.MessagesAffectedHistory

func (p *MessagesUnpinAllMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesUnpinAllMessagesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesUnpinAllMessagesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesUnpinAllMessagesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesUnpinAllMessagesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesUnpinAllMessagesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedHistory)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesUnpinAllMessagesResult) GetSuccess() *tg.MessagesAffectedHistory {
	if !p.IsSetSuccess() {
		return MessagesUnpinAllMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesUnpinAllMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedHistory)
}

func (p *MessagesUnpinAllMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesUnpinAllMessagesResult) GetResult() interface{} {
	return p.Success
}

func messagesGetSearchResultsCalendarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetSearchResultsCalendarArgs)
	realResult := result.(*MessagesGetSearchResultsCalendarResult)
	success, err := handler.(tg.RPCMessages).MessagesGetSearchResultsCalendar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetSearchResultsCalendarArgs() interface{} {
	return &MessagesGetSearchResultsCalendarArgs{}
}

func newMessagesGetSearchResultsCalendarResult() interface{} {
	return &MessagesGetSearchResultsCalendarResult{}
}

type MessagesGetSearchResultsCalendarArgs struct {
	Req *tg.TLMessagesGetSearchResultsCalendar
}

func (p *MessagesGetSearchResultsCalendarArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetSearchResultsCalendarArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetSearchResultsCalendarArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetSearchResultsCalendar)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetSearchResultsCalendarArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetSearchResultsCalendarArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetSearchResultsCalendarArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetSearchResultsCalendar)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetSearchResultsCalendarArgs_Req_DEFAULT *tg.TLMessagesGetSearchResultsCalendar

func (p *MessagesGetSearchResultsCalendarArgs) GetReq() *tg.TLMessagesGetSearchResultsCalendar {
	if !p.IsSetReq() {
		return MessagesGetSearchResultsCalendarArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetSearchResultsCalendarArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetSearchResultsCalendarResult struct {
	Success *tg.MessagesSearchResultsCalendar
}

var MessagesGetSearchResultsCalendarResult_Success_DEFAULT *tg.MessagesSearchResultsCalendar

func (p *MessagesGetSearchResultsCalendarResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetSearchResultsCalendarResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetSearchResultsCalendarResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesSearchResultsCalendar)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSearchResultsCalendarResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetSearchResultsCalendarResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetSearchResultsCalendarResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesSearchResultsCalendar)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSearchResultsCalendarResult) GetSuccess() *tg.MessagesSearchResultsCalendar {
	if !p.IsSetSuccess() {
		return MessagesGetSearchResultsCalendarResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetSearchResultsCalendarResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesSearchResultsCalendar)
}

func (p *MessagesGetSearchResultsCalendarResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetSearchResultsCalendarResult) GetResult() interface{} {
	return p.Success
}

func messagesGetSearchResultsPositionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetSearchResultsPositionsArgs)
	realResult := result.(*MessagesGetSearchResultsPositionsResult)
	success, err := handler.(tg.RPCMessages).MessagesGetSearchResultsPositions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetSearchResultsPositionsArgs() interface{} {
	return &MessagesGetSearchResultsPositionsArgs{}
}

func newMessagesGetSearchResultsPositionsResult() interface{} {
	return &MessagesGetSearchResultsPositionsResult{}
}

type MessagesGetSearchResultsPositionsArgs struct {
	Req *tg.TLMessagesGetSearchResultsPositions
}

func (p *MessagesGetSearchResultsPositionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetSearchResultsPositionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetSearchResultsPositionsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetSearchResultsPositions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetSearchResultsPositionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetSearchResultsPositionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetSearchResultsPositionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetSearchResultsPositions)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetSearchResultsPositionsArgs_Req_DEFAULT *tg.TLMessagesGetSearchResultsPositions

func (p *MessagesGetSearchResultsPositionsArgs) GetReq() *tg.TLMessagesGetSearchResultsPositions {
	if !p.IsSetReq() {
		return MessagesGetSearchResultsPositionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetSearchResultsPositionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetSearchResultsPositionsResult struct {
	Success *tg.MessagesSearchResultsPositions
}

var MessagesGetSearchResultsPositionsResult_Success_DEFAULT *tg.MessagesSearchResultsPositions

func (p *MessagesGetSearchResultsPositionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetSearchResultsPositionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetSearchResultsPositionsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesSearchResultsPositions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSearchResultsPositionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetSearchResultsPositionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetSearchResultsPositionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesSearchResultsPositions)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetSearchResultsPositionsResult) GetSuccess() *tg.MessagesSearchResultsPositions {
	if !p.IsSetSuccess() {
		return MessagesGetSearchResultsPositionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetSearchResultsPositionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesSearchResultsPositions)
}

func (p *MessagesGetSearchResultsPositionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetSearchResultsPositionsResult) GetResult() interface{} {
	return p.Success
}

func messagesToggleNoForwardsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesToggleNoForwardsArgs)
	realResult := result.(*MessagesToggleNoForwardsResult)
	success, err := handler.(tg.RPCMessages).MessagesToggleNoForwards(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesToggleNoForwardsArgs() interface{} {
	return &MessagesToggleNoForwardsArgs{}
}

func newMessagesToggleNoForwardsResult() interface{} {
	return &MessagesToggleNoForwardsResult{}
}

type MessagesToggleNoForwardsArgs struct {
	Req *tg.TLMessagesToggleNoForwards
}

func (p *MessagesToggleNoForwardsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesToggleNoForwardsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesToggleNoForwardsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesToggleNoForwards)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesToggleNoForwardsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesToggleNoForwardsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesToggleNoForwardsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesToggleNoForwards)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesToggleNoForwardsArgs_Req_DEFAULT *tg.TLMessagesToggleNoForwards

func (p *MessagesToggleNoForwardsArgs) GetReq() *tg.TLMessagesToggleNoForwards {
	if !p.IsSetReq() {
		return MessagesToggleNoForwardsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesToggleNoForwardsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesToggleNoForwardsResult struct {
	Success *tg.Updates
}

var MessagesToggleNoForwardsResult_Success_DEFAULT *tg.Updates

func (p *MessagesToggleNoForwardsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesToggleNoForwardsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesToggleNoForwardsResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesToggleNoForwardsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesToggleNoForwardsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesToggleNoForwardsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesToggleNoForwardsResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesToggleNoForwardsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesToggleNoForwardsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesToggleNoForwardsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesToggleNoForwardsResult) GetResult() interface{} {
	return p.Success
}

func messagesSaveDefaultSendAsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSaveDefaultSendAsArgs)
	realResult := result.(*MessagesSaveDefaultSendAsResult)
	success, err := handler.(tg.RPCMessages).MessagesSaveDefaultSendAs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSaveDefaultSendAsArgs() interface{} {
	return &MessagesSaveDefaultSendAsArgs{}
}

func newMessagesSaveDefaultSendAsResult() interface{} {
	return &MessagesSaveDefaultSendAsResult{}
}

type MessagesSaveDefaultSendAsArgs struct {
	Req *tg.TLMessagesSaveDefaultSendAs
}

func (p *MessagesSaveDefaultSendAsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSaveDefaultSendAsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSaveDefaultSendAsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSaveDefaultSendAs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSaveDefaultSendAsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSaveDefaultSendAsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSaveDefaultSendAsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSaveDefaultSendAs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSaveDefaultSendAsArgs_Req_DEFAULT *tg.TLMessagesSaveDefaultSendAs

func (p *MessagesSaveDefaultSendAsArgs) GetReq() *tg.TLMessagesSaveDefaultSendAs {
	if !p.IsSetReq() {
		return MessagesSaveDefaultSendAsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSaveDefaultSendAsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSaveDefaultSendAsResult struct {
	Success *tg.Bool
}

var MessagesSaveDefaultSendAsResult_Success_DEFAULT *tg.Bool

func (p *MessagesSaveDefaultSendAsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSaveDefaultSendAsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSaveDefaultSendAsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSaveDefaultSendAsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSaveDefaultSendAsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSaveDefaultSendAsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSaveDefaultSendAsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesSaveDefaultSendAsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSaveDefaultSendAsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesSaveDefaultSendAsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSaveDefaultSendAsResult) GetResult() interface{} {
	return p.Success
}

func messagesSearchSentMediaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesSearchSentMediaArgs)
	realResult := result.(*MessagesSearchSentMediaResult)
	success, err := handler.(tg.RPCMessages).MessagesSearchSentMedia(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesSearchSentMediaArgs() interface{} {
	return &MessagesSearchSentMediaArgs{}
}

func newMessagesSearchSentMediaResult() interface{} {
	return &MessagesSearchSentMediaResult{}
}

type MessagesSearchSentMediaArgs struct {
	Req *tg.TLMessagesSearchSentMedia
}

func (p *MessagesSearchSentMediaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesSearchSentMediaArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesSearchSentMediaArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesSearchSentMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesSearchSentMediaArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesSearchSentMediaArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesSearchSentMediaArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesSearchSentMedia)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesSearchSentMediaArgs_Req_DEFAULT *tg.TLMessagesSearchSentMedia

func (p *MessagesSearchSentMediaArgs) GetReq() *tg.TLMessagesSearchSentMedia {
	if !p.IsSetReq() {
		return MessagesSearchSentMediaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesSearchSentMediaArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesSearchSentMediaResult struct {
	Success *tg.MessagesMessages
}

var MessagesSearchSentMediaResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesSearchSentMediaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesSearchSentMediaResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesSearchSentMediaResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSearchSentMediaResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesSearchSentMediaResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesSearchSentMediaResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesSearchSentMediaResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesSearchSentMediaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesSearchSentMediaResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesSearchSentMediaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesSearchSentMediaResult) GetResult() interface{} {
	return p.Success
}

func messagesGetOutboxReadDateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetOutboxReadDateArgs)
	realResult := result.(*MessagesGetOutboxReadDateResult)
	success, err := handler.(tg.RPCMessages).MessagesGetOutboxReadDate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetOutboxReadDateArgs() interface{} {
	return &MessagesGetOutboxReadDateArgs{}
}

func newMessagesGetOutboxReadDateResult() interface{} {
	return &MessagesGetOutboxReadDateResult{}
}

type MessagesGetOutboxReadDateArgs struct {
	Req *tg.TLMessagesGetOutboxReadDate
}

func (p *MessagesGetOutboxReadDateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetOutboxReadDateArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetOutboxReadDateArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetOutboxReadDate)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetOutboxReadDateArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetOutboxReadDateArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetOutboxReadDateArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetOutboxReadDate)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetOutboxReadDateArgs_Req_DEFAULT *tg.TLMessagesGetOutboxReadDate

func (p *MessagesGetOutboxReadDateArgs) GetReq() *tg.TLMessagesGetOutboxReadDate {
	if !p.IsSetReq() {
		return MessagesGetOutboxReadDateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetOutboxReadDateArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetOutboxReadDateResult struct {
	Success *tg.OutboxReadDate
}

var MessagesGetOutboxReadDateResult_Success_DEFAULT *tg.OutboxReadDate

func (p *MessagesGetOutboxReadDateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetOutboxReadDateResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetOutboxReadDateResult) Unmarshal(in []byte) error {
	msg := new(tg.OutboxReadDate)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetOutboxReadDateResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetOutboxReadDateResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetOutboxReadDateResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.OutboxReadDate)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetOutboxReadDateResult) GetSuccess() *tg.OutboxReadDate {
	if !p.IsSetSuccess() {
		return MessagesGetOutboxReadDateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetOutboxReadDateResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.OutboxReadDate)
}

func (p *MessagesGetOutboxReadDateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetOutboxReadDateResult) GetResult() interface{} {
	return p.Success
}

func channelsGetSendAsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsGetSendAsArgs)
	realResult := result.(*ChannelsGetSendAsResult)
	success, err := handler.(tg.RPCMessages).ChannelsGetSendAs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsGetSendAsArgs() interface{} {
	return &ChannelsGetSendAsArgs{}
}

func newChannelsGetSendAsResult() interface{} {
	return &ChannelsGetSendAsResult{}
}

type ChannelsGetSendAsArgs struct {
	Req *tg.TLChannelsGetSendAs
}

func (p *ChannelsGetSendAsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsGetSendAsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsGetSendAsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsGetSendAs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsGetSendAsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsGetSendAsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsGetSendAsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsGetSendAs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsGetSendAsArgs_Req_DEFAULT *tg.TLChannelsGetSendAs

func (p *ChannelsGetSendAsArgs) GetReq() *tg.TLChannelsGetSendAs {
	if !p.IsSetReq() {
		return ChannelsGetSendAsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsGetSendAsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsGetSendAsResult struct {
	Success *tg.ChannelsSendAsPeers
}

var ChannelsGetSendAsResult_Success_DEFAULT *tg.ChannelsSendAsPeers

func (p *ChannelsGetSendAsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsGetSendAsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsGetSendAsResult) Unmarshal(in []byte) error {
	msg := new(tg.ChannelsSendAsPeers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsGetSendAsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsGetSendAsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsGetSendAsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ChannelsSendAsPeers)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsGetSendAsResult) GetSuccess() *tg.ChannelsSendAsPeers {
	if !p.IsSetSuccess() {
		return ChannelsGetSendAsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsGetSendAsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ChannelsSendAsPeers)
}

func (p *ChannelsGetSendAsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsGetSendAsResult) GetResult() interface{} {
	return p.Success
}

func channelsSearchPostsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsSearchPostsArgs)
	realResult := result.(*ChannelsSearchPostsResult)
	success, err := handler.(tg.RPCMessages).ChannelsSearchPosts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsSearchPostsArgs() interface{} {
	return &ChannelsSearchPostsArgs{}
}

func newChannelsSearchPostsResult() interface{} {
	return &ChannelsSearchPostsResult{}
}

type ChannelsSearchPostsArgs struct {
	Req *tg.TLChannelsSearchPosts
}

func (p *ChannelsSearchPostsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsSearchPostsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsSearchPostsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsSearchPosts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsSearchPostsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsSearchPostsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsSearchPostsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsSearchPosts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsSearchPostsArgs_Req_DEFAULT *tg.TLChannelsSearchPosts

func (p *ChannelsSearchPostsArgs) GetReq() *tg.TLChannelsSearchPosts {
	if !p.IsSetReq() {
		return ChannelsSearchPostsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsSearchPostsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsSearchPostsResult struct {
	Success *tg.MessagesMessages
}

var ChannelsSearchPostsResult_Success_DEFAULT *tg.MessagesMessages

func (p *ChannelsSearchPostsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsSearchPostsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsSearchPostsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsSearchPostsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsSearchPostsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsSearchPostsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsSearchPostsResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return ChannelsSearchPostsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsSearchPostsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *ChannelsSearchPostsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsSearchPostsResult) GetResult() interface{} {
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

func (p *kClient) MessagesGetMessages(ctx context.Context, req *tg.TLMessagesGetMessages) (r *tg.MessagesMessages, err error) {
	var _args MessagesGetMessagesArgs
	_args.Req = req
	var _result MessagesGetMessagesResult
	if err = p.c.Call(ctx, "messages.getMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetHistory(ctx context.Context, req *tg.TLMessagesGetHistory) (r *tg.MessagesMessages, err error) {
	var _args MessagesGetHistoryArgs
	_args.Req = req
	var _result MessagesGetHistoryResult
	if err = p.c.Call(ctx, "messages.getHistory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesSearch(ctx context.Context, req *tg.TLMessagesSearch) (r *tg.MessagesMessages, err error) {
	var _args MessagesSearchArgs
	_args.Req = req
	var _result MessagesSearchResult
	if err = p.c.Call(ctx, "messages.search", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesReadHistory(ctx context.Context, req *tg.TLMessagesReadHistory) (r *tg.MessagesAffectedMessages, err error) {
	var _args MessagesReadHistoryArgs
	_args.Req = req
	var _result MessagesReadHistoryResult
	if err = p.c.Call(ctx, "messages.readHistory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesDeleteHistory(ctx context.Context, req *tg.TLMessagesDeleteHistory) (r *tg.MessagesAffectedHistory, err error) {
	var _args MessagesDeleteHistoryArgs
	_args.Req = req
	var _result MessagesDeleteHistoryResult
	if err = p.c.Call(ctx, "messages.deleteHistory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesDeleteMessages(ctx context.Context, req *tg.TLMessagesDeleteMessages) (r *tg.MessagesAffectedMessages, err error) {
	var _args MessagesDeleteMessagesArgs
	_args.Req = req
	var _result MessagesDeleteMessagesResult
	if err = p.c.Call(ctx, "messages.deleteMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesReceivedMessages(ctx context.Context, req *tg.TLMessagesReceivedMessages) (r *tg.VectorReceivedNotifyMessage, err error) {
	var _args MessagesReceivedMessagesArgs
	_args.Req = req
	var _result MessagesReceivedMessagesResult
	if err = p.c.Call(ctx, "messages.receivedMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesSendMessage(ctx context.Context, req *tg.TLMessagesSendMessage) (r *tg.Updates, err error) {
	var _args MessagesSendMessageArgs
	_args.Req = req
	var _result MessagesSendMessageResult
	if err = p.c.Call(ctx, "messages.sendMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesSendMedia(ctx context.Context, req *tg.TLMessagesSendMedia) (r *tg.Updates, err error) {
	var _args MessagesSendMediaArgs
	_args.Req = req
	var _result MessagesSendMediaResult
	if err = p.c.Call(ctx, "messages.sendMedia", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesForwardMessages(ctx context.Context, req *tg.TLMessagesForwardMessages) (r *tg.Updates, err error) {
	var _args MessagesForwardMessagesArgs
	_args.Req = req
	var _result MessagesForwardMessagesResult
	if err = p.c.Call(ctx, "messages.forwardMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesReadMessageContents(ctx context.Context, req *tg.TLMessagesReadMessageContents) (r *tg.MessagesAffectedMessages, err error) {
	var _args MessagesReadMessageContentsArgs
	_args.Req = req
	var _result MessagesReadMessageContentsResult
	if err = p.c.Call(ctx, "messages.readMessageContents", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetMessagesViews(ctx context.Context, req *tg.TLMessagesGetMessagesViews) (r *tg.MessagesMessageViews, err error) {
	var _args MessagesGetMessagesViewsArgs
	_args.Req = req
	var _result MessagesGetMessagesViewsResult
	if err = p.c.Call(ctx, "messages.getMessagesViews", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesSearchGlobal(ctx context.Context, req *tg.TLMessagesSearchGlobal) (r *tg.MessagesMessages, err error) {
	var _args MessagesSearchGlobalArgs
	_args.Req = req
	var _result MessagesSearchGlobalResult
	if err = p.c.Call(ctx, "messages.searchGlobal", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetMessageEditData(ctx context.Context, req *tg.TLMessagesGetMessageEditData) (r *tg.MessagesMessageEditData, err error) {
	var _args MessagesGetMessageEditDataArgs
	_args.Req = req
	var _result MessagesGetMessageEditDataResult
	if err = p.c.Call(ctx, "messages.getMessageEditData", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesEditMessage(ctx context.Context, req *tg.TLMessagesEditMessage) (r *tg.Updates, err error) {
	var _args MessagesEditMessageArgs
	_args.Req = req
	var _result MessagesEditMessageResult
	if err = p.c.Call(ctx, "messages.editMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetUnreadMentions(ctx context.Context, req *tg.TLMessagesGetUnreadMentions) (r *tg.MessagesMessages, err error) {
	var _args MessagesGetUnreadMentionsArgs
	_args.Req = req
	var _result MessagesGetUnreadMentionsResult
	if err = p.c.Call(ctx, "messages.getUnreadMentions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesReadMentions(ctx context.Context, req *tg.TLMessagesReadMentions) (r *tg.MessagesAffectedHistory, err error) {
	var _args MessagesReadMentionsArgs
	_args.Req = req
	var _result MessagesReadMentionsResult
	if err = p.c.Call(ctx, "messages.readMentions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetRecentLocations(ctx context.Context, req *tg.TLMessagesGetRecentLocations) (r *tg.MessagesMessages, err error) {
	var _args MessagesGetRecentLocationsArgs
	_args.Req = req
	var _result MessagesGetRecentLocationsResult
	if err = p.c.Call(ctx, "messages.getRecentLocations", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesSendMultiMedia(ctx context.Context, req *tg.TLMessagesSendMultiMedia) (r *tg.Updates, err error) {
	var _args MessagesSendMultiMediaArgs
	_args.Req = req
	var _result MessagesSendMultiMediaResult
	if err = p.c.Call(ctx, "messages.sendMultiMedia", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesUpdatePinnedMessage(ctx context.Context, req *tg.TLMessagesUpdatePinnedMessage) (r *tg.Updates, err error) {
	var _args MessagesUpdatePinnedMessageArgs
	_args.Req = req
	var _result MessagesUpdatePinnedMessageResult
	if err = p.c.Call(ctx, "messages.updatePinnedMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetSearchCounters(ctx context.Context, req *tg.TLMessagesGetSearchCounters) (r *tg.VectorMessagesSearchCounter, err error) {
	var _args MessagesGetSearchCountersArgs
	_args.Req = req
	var _result MessagesGetSearchCountersResult
	if err = p.c.Call(ctx, "messages.getSearchCounters", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesUnpinAllMessages(ctx context.Context, req *tg.TLMessagesUnpinAllMessages) (r *tg.MessagesAffectedHistory, err error) {
	var _args MessagesUnpinAllMessagesArgs
	_args.Req = req
	var _result MessagesUnpinAllMessagesResult
	if err = p.c.Call(ctx, "messages.unpinAllMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetSearchResultsCalendar(ctx context.Context, req *tg.TLMessagesGetSearchResultsCalendar) (r *tg.MessagesSearchResultsCalendar, err error) {
	var _args MessagesGetSearchResultsCalendarArgs
	_args.Req = req
	var _result MessagesGetSearchResultsCalendarResult
	if err = p.c.Call(ctx, "messages.getSearchResultsCalendar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetSearchResultsPositions(ctx context.Context, req *tg.TLMessagesGetSearchResultsPositions) (r *tg.MessagesSearchResultsPositions, err error) {
	var _args MessagesGetSearchResultsPositionsArgs
	_args.Req = req
	var _result MessagesGetSearchResultsPositionsResult
	if err = p.c.Call(ctx, "messages.getSearchResultsPositions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesToggleNoForwards(ctx context.Context, req *tg.TLMessagesToggleNoForwards) (r *tg.Updates, err error) {
	var _args MessagesToggleNoForwardsArgs
	_args.Req = req
	var _result MessagesToggleNoForwardsResult
	if err = p.c.Call(ctx, "messages.toggleNoForwards", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesSaveDefaultSendAs(ctx context.Context, req *tg.TLMessagesSaveDefaultSendAs) (r *tg.Bool, err error) {
	var _args MessagesSaveDefaultSendAsArgs
	_args.Req = req
	var _result MessagesSaveDefaultSendAsResult
	if err = p.c.Call(ctx, "messages.saveDefaultSendAs", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesSearchSentMedia(ctx context.Context, req *tg.TLMessagesSearchSentMedia) (r *tg.MessagesMessages, err error) {
	var _args MessagesSearchSentMediaArgs
	_args.Req = req
	var _result MessagesSearchSentMediaResult
	if err = p.c.Call(ctx, "messages.searchSentMedia", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetOutboxReadDate(ctx context.Context, req *tg.TLMessagesGetOutboxReadDate) (r *tg.OutboxReadDate, err error) {
	var _args MessagesGetOutboxReadDateArgs
	_args.Req = req
	var _result MessagesGetOutboxReadDateResult
	if err = p.c.Call(ctx, "messages.getOutboxReadDate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsGetSendAs(ctx context.Context, req *tg.TLChannelsGetSendAs) (r *tg.ChannelsSendAsPeers, err error) {
	var _args ChannelsGetSendAsArgs
	_args.Req = req
	var _result ChannelsGetSendAsResult
	if err = p.c.Call(ctx, "channels.getSendAs", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsSearchPosts(ctx context.Context, req *tg.TLChannelsSearchPosts) (r *tg.MessagesMessages, err error) {
	var _args ChannelsSearchPostsArgs
	_args.Req = req
	var _result ChannelsSearchPostsResult
	if err = p.c.Call(ctx, "channels.searchPosts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
