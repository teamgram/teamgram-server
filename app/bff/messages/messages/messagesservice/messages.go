/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package messagesservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/tg.RPCMessages/messages.composeMessageWithAI": kitex.NewMethodInfo(
		messagesComposeMessageWithAIHandler,
		newMessagesComposeMessageWithAIArgs,
		newMessagesComposeMessageWithAIResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCMessages/messages.reportReadMetrics": kitex.NewMethodInfo(
		messagesReportReadMetricsHandler,
		newMessagesReportReadMetricsArgs,
		newMessagesReportReadMetricsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCMessages/messages.reportMusicListen": kitex.NewMethodInfo(
		messagesReportMusicListenHandler,
		newMessagesReportMusicListenArgs,
		newMessagesReportMusicListenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCMessages/messages.addPollAnswer": kitex.NewMethodInfo(
		messagesAddPollAnswerHandler,
		newMessagesAddPollAnswerArgs,
		newMessagesAddPollAnswerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCMessages/messages.deletePollAnswer": kitex.NewMethodInfo(
		messagesDeletePollAnswerHandler,
		newMessagesDeletePollAnswerArgs,
		newMessagesDeletePollAnswerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCMessages/messages.getUnreadPollVotes": kitex.NewMethodInfo(
		messagesGetUnreadPollVotesHandler,
		newMessagesGetUnreadPollVotesArgs,
		newMessagesGetUnreadPollVotesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCMessages/messages.readPollVotes": kitex.NewMethodInfo(
		messagesReadPollVotesHandler,
		newMessagesReadPollVotesArgs,
		newMessagesReadPollVotesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	messagesServiceServiceInfo                = NewServiceInfo()
	messagesServiceServiceInfoForClient       = NewServiceInfoForClient()
	messagesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCMessages", messagesServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCMessages", messagesServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCMessages", messagesServiceServiceInfoForStreamClient)
}

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

// NewServiceInfoForStreamClient creates a new ServiceInfo containing all streaming methods
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

func messagesComposeMessageWithAIHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesComposeMessageWithAIArgs)
	realResult := result.(*MessagesComposeMessageWithAIResult)
	success, err := handler.(tg.RPCMessages).MessagesComposeMessageWithAI(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesComposeMessageWithAIArgs() interface{} {
	return &MessagesComposeMessageWithAIArgs{}
}

func newMessagesComposeMessageWithAIResult() interface{} {
	return &MessagesComposeMessageWithAIResult{}
}

type MessagesComposeMessageWithAIArgs struct {
	Req *tg.TLMessagesComposeMessageWithAI
}

func (p *MessagesComposeMessageWithAIArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesComposeMessageWithAIArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesComposeMessageWithAIArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesComposeMessageWithAI)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesComposeMessageWithAIArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesComposeMessageWithAIArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesComposeMessageWithAIArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesComposeMessageWithAI)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessagesComposeMessageWithAIArgs_Req_DEFAULT *tg.TLMessagesComposeMessageWithAI

func (p *MessagesComposeMessageWithAIArgs) GetReq() *tg.TLMessagesComposeMessageWithAI {
	if !p.IsSetReq() {
		return MessagesComposeMessageWithAIArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesComposeMessageWithAIArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesComposeMessageWithAIResult struct {
	Success *tg.MessagesComposedMessageWithAI
}

var MessagesComposeMessageWithAIResult_Success_DEFAULT *tg.MessagesComposedMessageWithAI

func (p *MessagesComposeMessageWithAIResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesComposeMessageWithAIResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesComposeMessageWithAIResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesComposedMessageWithAI)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesComposeMessageWithAIResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesComposeMessageWithAIResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesComposeMessageWithAIResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesComposedMessageWithAI)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesComposeMessageWithAIResult) GetSuccess() *tg.MessagesComposedMessageWithAI {
	if !p.IsSetSuccess() {
		return MessagesComposeMessageWithAIResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesComposeMessageWithAIResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesComposedMessageWithAI)
}

func (p *MessagesComposeMessageWithAIResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesComposeMessageWithAIResult) GetResult() interface{} {
	return p.Success
}

func messagesReportReadMetricsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReportReadMetricsArgs)
	realResult := result.(*MessagesReportReadMetricsResult)
	success, err := handler.(tg.RPCMessages).MessagesReportReadMetrics(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReportReadMetricsArgs() interface{} {
	return &MessagesReportReadMetricsArgs{}
}

func newMessagesReportReadMetricsResult() interface{} {
	return &MessagesReportReadMetricsResult{}
}

type MessagesReportReadMetricsArgs struct {
	Req *tg.TLMessagesReportReadMetrics
}

func (p *MessagesReportReadMetricsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReportReadMetricsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReportReadMetricsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReportReadMetrics)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReportReadMetricsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReportReadMetricsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReportReadMetricsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReportReadMetrics)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessagesReportReadMetricsArgs_Req_DEFAULT *tg.TLMessagesReportReadMetrics

func (p *MessagesReportReadMetricsArgs) GetReq() *tg.TLMessagesReportReadMetrics {
	if !p.IsSetReq() {
		return MessagesReportReadMetricsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReportReadMetricsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReportReadMetricsResult struct {
	Success *tg.Bool
}

var MessagesReportReadMetricsResult_Success_DEFAULT *tg.Bool

func (p *MessagesReportReadMetricsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReportReadMetricsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReportReadMetricsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReportReadMetricsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReportReadMetricsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReportReadMetricsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReportReadMetricsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesReportReadMetricsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReportReadMetricsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesReportReadMetricsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReportReadMetricsResult) GetResult() interface{} {
	return p.Success
}

func messagesReportMusicListenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReportMusicListenArgs)
	realResult := result.(*MessagesReportMusicListenResult)
	success, err := handler.(tg.RPCMessages).MessagesReportMusicListen(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReportMusicListenArgs() interface{} {
	return &MessagesReportMusicListenArgs{}
}

func newMessagesReportMusicListenResult() interface{} {
	return &MessagesReportMusicListenResult{}
}

type MessagesReportMusicListenArgs struct {
	Req *tg.TLMessagesReportMusicListen
}

func (p *MessagesReportMusicListenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReportMusicListenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReportMusicListenArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReportMusicListen)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReportMusicListenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReportMusicListenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReportMusicListenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReportMusicListen)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessagesReportMusicListenArgs_Req_DEFAULT *tg.TLMessagesReportMusicListen

func (p *MessagesReportMusicListenArgs) GetReq() *tg.TLMessagesReportMusicListen {
	if !p.IsSetReq() {
		return MessagesReportMusicListenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReportMusicListenArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReportMusicListenResult struct {
	Success *tg.Bool
}

var MessagesReportMusicListenResult_Success_DEFAULT *tg.Bool

func (p *MessagesReportMusicListenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReportMusicListenResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReportMusicListenResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReportMusicListenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReportMusicListenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReportMusicListenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReportMusicListenResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesReportMusicListenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReportMusicListenResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesReportMusicListenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReportMusicListenResult) GetResult() interface{} {
	return p.Success
}

func messagesAddPollAnswerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesAddPollAnswerArgs)
	realResult := result.(*MessagesAddPollAnswerResult)
	success, err := handler.(tg.RPCMessages).MessagesAddPollAnswer(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesAddPollAnswerArgs() interface{} {
	return &MessagesAddPollAnswerArgs{}
}

func newMessagesAddPollAnswerResult() interface{} {
	return &MessagesAddPollAnswerResult{}
}

type MessagesAddPollAnswerArgs struct {
	Req *tg.TLMessagesAddPollAnswer
}

func (p *MessagesAddPollAnswerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesAddPollAnswerArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesAddPollAnswerArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesAddPollAnswer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesAddPollAnswerArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesAddPollAnswerArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesAddPollAnswerArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesAddPollAnswer)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessagesAddPollAnswerArgs_Req_DEFAULT *tg.TLMessagesAddPollAnswer

func (p *MessagesAddPollAnswerArgs) GetReq() *tg.TLMessagesAddPollAnswer {
	if !p.IsSetReq() {
		return MessagesAddPollAnswerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesAddPollAnswerArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesAddPollAnswerResult struct {
	Success *tg.Updates
}

var MessagesAddPollAnswerResult_Success_DEFAULT *tg.Updates

func (p *MessagesAddPollAnswerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesAddPollAnswerResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesAddPollAnswerResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesAddPollAnswerResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesAddPollAnswerResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesAddPollAnswerResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesAddPollAnswerResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesAddPollAnswerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesAddPollAnswerResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesAddPollAnswerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesAddPollAnswerResult) GetResult() interface{} {
	return p.Success
}

func messagesDeletePollAnswerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesDeletePollAnswerArgs)
	realResult := result.(*MessagesDeletePollAnswerResult)
	success, err := handler.(tg.RPCMessages).MessagesDeletePollAnswer(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesDeletePollAnswerArgs() interface{} {
	return &MessagesDeletePollAnswerArgs{}
}

func newMessagesDeletePollAnswerResult() interface{} {
	return &MessagesDeletePollAnswerResult{}
}

type MessagesDeletePollAnswerArgs struct {
	Req *tg.TLMessagesDeletePollAnswer
}

func (p *MessagesDeletePollAnswerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesDeletePollAnswerArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesDeletePollAnswerArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesDeletePollAnswer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesDeletePollAnswerArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesDeletePollAnswerArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesDeletePollAnswerArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesDeletePollAnswer)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessagesDeletePollAnswerArgs_Req_DEFAULT *tg.TLMessagesDeletePollAnswer

func (p *MessagesDeletePollAnswerArgs) GetReq() *tg.TLMessagesDeletePollAnswer {
	if !p.IsSetReq() {
		return MessagesDeletePollAnswerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesDeletePollAnswerArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesDeletePollAnswerResult struct {
	Success *tg.Updates
}

var MessagesDeletePollAnswerResult_Success_DEFAULT *tg.Updates

func (p *MessagesDeletePollAnswerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesDeletePollAnswerResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesDeletePollAnswerResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeletePollAnswerResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesDeletePollAnswerResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesDeletePollAnswerResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeletePollAnswerResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesDeletePollAnswerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesDeletePollAnswerResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesDeletePollAnswerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesDeletePollAnswerResult) GetResult() interface{} {
	return p.Success
}

func messagesGetUnreadPollVotesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetUnreadPollVotesArgs)
	realResult := result.(*MessagesGetUnreadPollVotesResult)
	success, err := handler.(tg.RPCMessages).MessagesGetUnreadPollVotes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetUnreadPollVotesArgs() interface{} {
	return &MessagesGetUnreadPollVotesArgs{}
}

func newMessagesGetUnreadPollVotesResult() interface{} {
	return &MessagesGetUnreadPollVotesResult{}
}

type MessagesGetUnreadPollVotesArgs struct {
	Req *tg.TLMessagesGetUnreadPollVotes
}

func (p *MessagesGetUnreadPollVotesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetUnreadPollVotesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetUnreadPollVotesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetUnreadPollVotes)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetUnreadPollVotesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetUnreadPollVotesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetUnreadPollVotesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetUnreadPollVotes)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessagesGetUnreadPollVotesArgs_Req_DEFAULT *tg.TLMessagesGetUnreadPollVotes

func (p *MessagesGetUnreadPollVotesArgs) GetReq() *tg.TLMessagesGetUnreadPollVotes {
	if !p.IsSetReq() {
		return MessagesGetUnreadPollVotesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetUnreadPollVotesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetUnreadPollVotesResult struct {
	Success *tg.MessagesMessages
}

var MessagesGetUnreadPollVotesResult_Success_DEFAULT *tg.MessagesMessages

func (p *MessagesGetUnreadPollVotesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetUnreadPollVotesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetUnreadPollVotesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesMessages)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetUnreadPollVotesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetUnreadPollVotesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetUnreadPollVotesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesMessages)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetUnreadPollVotesResult) GetSuccess() *tg.MessagesMessages {
	if !p.IsSetSuccess() {
		return MessagesGetUnreadPollVotesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetUnreadPollVotesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesMessages)
}

func (p *MessagesGetUnreadPollVotesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetUnreadPollVotesResult) GetResult() interface{} {
	return p.Success
}

func messagesReadPollVotesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesReadPollVotesArgs)
	realResult := result.(*MessagesReadPollVotesResult)
	success, err := handler.(tg.RPCMessages).MessagesReadPollVotes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesReadPollVotesArgs() interface{} {
	return &MessagesReadPollVotesArgs{}
}

func newMessagesReadPollVotesResult() interface{} {
	return &MessagesReadPollVotesResult{}
}

type MessagesReadPollVotesArgs struct {
	Req *tg.TLMessagesReadPollVotes
}

func (p *MessagesReadPollVotesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesReadPollVotesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesReadPollVotesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesReadPollVotes)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesReadPollVotesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesReadPollVotesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesReadPollVotesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesReadPollVotes)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessagesReadPollVotesArgs_Req_DEFAULT *tg.TLMessagesReadPollVotes

func (p *MessagesReadPollVotesArgs) GetReq() *tg.TLMessagesReadPollVotes {
	if !p.IsSetReq() {
		return MessagesReadPollVotesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesReadPollVotesArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesReadPollVotesResult struct {
	Success *tg.MessagesAffectedHistory
}

var MessagesReadPollVotesResult_Success_DEFAULT *tg.MessagesAffectedHistory

func (p *MessagesReadPollVotesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesReadPollVotesResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesReadPollVotesResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesAffectedHistory)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReadPollVotesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesReadPollVotesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesReadPollVotesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesAffectedHistory)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesReadPollVotesResult) GetSuccess() *tg.MessagesAffectedHistory {
	if !p.IsSetSuccess() {
		return MessagesReadPollVotesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesReadPollVotesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesAffectedHistory)
}

func (p *MessagesReadPollVotesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesReadPollVotesResult) GetResult() interface{} {
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

func (p *kClient) MessagesComposeMessageWithAI(ctx context.Context, req *tg.TLMessagesComposeMessageWithAI) (r *tg.MessagesComposedMessageWithAI, err error) {
	// var _args MessagesComposeMessageWithAIArgs
	// _args.Req = req
	// var _result MessagesComposeMessageWithAIResult

	_result := new(tg.MessagesComposedMessageWithAI)
	if err = p.c.Call(ctx, "/tg.RPCMessages/messages.composeMessageWithAI", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesReportReadMetrics(ctx context.Context, req *tg.TLMessagesReportReadMetrics) (r *tg.Bool, err error) {
	// var _args MessagesReportReadMetricsArgs
	// _args.Req = req
	// var _result MessagesReportReadMetricsResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCMessages/messages.reportReadMetrics", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesReportMusicListen(ctx context.Context, req *tg.TLMessagesReportMusicListen) (r *tg.Bool, err error) {
	// var _args MessagesReportMusicListenArgs
	// _args.Req = req
	// var _result MessagesReportMusicListenResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCMessages/messages.reportMusicListen", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesAddPollAnswer(ctx context.Context, req *tg.TLMessagesAddPollAnswer) (r *tg.Updates, err error) {
	// var _args MessagesAddPollAnswerArgs
	// _args.Req = req
	// var _result MessagesAddPollAnswerResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "/tg.RPCMessages/messages.addPollAnswer", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesDeletePollAnswer(ctx context.Context, req *tg.TLMessagesDeletePollAnswer) (r *tg.Updates, err error) {
	// var _args MessagesDeletePollAnswerArgs
	// _args.Req = req
	// var _result MessagesDeletePollAnswerResult

	_result := new(tg.Updates)
	if err = p.c.Call(ctx, "/tg.RPCMessages/messages.deletePollAnswer", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesGetUnreadPollVotes(ctx context.Context, req *tg.TLMessagesGetUnreadPollVotes) (r *tg.MessagesMessages, err error) {
	// var _args MessagesGetUnreadPollVotesArgs
	// _args.Req = req
	// var _result MessagesGetUnreadPollVotesResult

	_result := new(tg.MessagesMessages)
	if err = p.c.Call(ctx, "/tg.RPCMessages/messages.getUnreadPollVotes", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesReadPollVotes(ctx context.Context, req *tg.TLMessagesReadPollVotes) (r *tg.MessagesAffectedHistory, err error) {
	// var _args MessagesReadPollVotesArgs
	// _args.Req = req
	// var _result MessagesReadPollVotesResult

	_result := new(tg.MessagesAffectedHistory)
	if err = p.c.Call(ctx, "/tg.RPCMessages/messages.readPollVotes", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
