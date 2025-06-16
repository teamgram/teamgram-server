/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package statusservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"status.setSessionOnline": kitex.NewMethodInfo(
		setSessionOnlineHandler,
		newSetSessionOnlineArgs,
		newSetSessionOnlineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.setSessionOffline": kitex.NewMethodInfo(
		setSessionOfflineHandler,
		newSetSessionOfflineArgs,
		newSetSessionOfflineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.getUserOnlineSessions": kitex.NewMethodInfo(
		getUserOnlineSessionsHandler,
		newGetUserOnlineSessionsArgs,
		newGetUserOnlineSessionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.getUsersOnlineSessionsList": kitex.NewMethodInfo(
		getUsersOnlineSessionsListHandler,
		newGetUsersOnlineSessionsListArgs,
		newGetUsersOnlineSessionsListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.getChannelOnlineUsers": kitex.NewMethodInfo(
		getChannelOnlineUsersHandler,
		newGetChannelOnlineUsersArgs,
		newGetChannelOnlineUsersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.setUserChannelsOnline": kitex.NewMethodInfo(
		setUserChannelsOnlineHandler,
		newSetUserChannelsOnlineArgs,
		newSetUserChannelsOnlineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.setUserChannelsOffline": kitex.NewMethodInfo(
		setUserChannelsOfflineHandler,
		newSetUserChannelsOfflineArgs,
		newSetUserChannelsOfflineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.setChannelUserOffline": kitex.NewMethodInfo(
		setChannelUserOfflineHandler,
		newSetChannelUserOfflineArgs,
		newSetChannelUserOfflineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.setChannelUsersOnline": kitex.NewMethodInfo(
		setChannelUsersOnlineHandler,
		newSetChannelUsersOnlineArgs,
		newSetChannelUsersOnlineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"status.setChannelOffline": kitex.NewMethodInfo(
		setChannelOfflineHandler,
		newSetChannelOfflineArgs,
		newSetChannelOfflineResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	statusServiceServiceInfo                = NewServiceInfo()
	statusServiceServiceInfoForClient       = NewServiceInfoForClient()
	statusServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return statusServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return statusServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return statusServiceServiceInfoForClient
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
	serviceName := "RPCStatus"
	handlerType := (*status.RPCStatus)(nil)
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
		"PackageName": "status",
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

func setSessionOnlineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetSessionOnlineArgs)
	realResult := result.(*SetSessionOnlineResult)
	success, err := handler.(status.RPCStatus).StatusSetSessionOnline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetSessionOnlineArgs() interface{} {
	return &SetSessionOnlineArgs{}
}

func newSetSessionOnlineResult() interface{} {
	return &SetSessionOnlineResult{}
}

type SetSessionOnlineArgs struct {
	Req *status.TLStatusSetSessionOnline
}

func (p *SetSessionOnlineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetSessionOnlineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetSessionOnlineArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusSetSessionOnline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetSessionOnlineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetSessionOnlineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetSessionOnlineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusSetSessionOnline)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetSessionOnlineArgs_Req_DEFAULT *status.TLStatusSetSessionOnline

func (p *SetSessionOnlineArgs) GetReq() *status.TLStatusSetSessionOnline {
	if !p.IsSetReq() {
		return SetSessionOnlineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetSessionOnlineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetSessionOnlineResult struct {
	Success *tg.Bool
}

var SetSessionOnlineResult_Success_DEFAULT *tg.Bool

func (p *SetSessionOnlineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetSessionOnlineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetSessionOnlineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetSessionOnlineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetSessionOnlineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetSessionOnlineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetSessionOnlineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetSessionOnlineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetSessionOnlineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetSessionOnlineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetSessionOnlineResult) GetResult() interface{} {
	return p.Success
}

func setSessionOfflineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetSessionOfflineArgs)
	realResult := result.(*SetSessionOfflineResult)
	success, err := handler.(status.RPCStatus).StatusSetSessionOffline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetSessionOfflineArgs() interface{} {
	return &SetSessionOfflineArgs{}
}

func newSetSessionOfflineResult() interface{} {
	return &SetSessionOfflineResult{}
}

type SetSessionOfflineArgs struct {
	Req *status.TLStatusSetSessionOffline
}

func (p *SetSessionOfflineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetSessionOfflineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetSessionOfflineArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusSetSessionOffline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetSessionOfflineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetSessionOfflineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetSessionOfflineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusSetSessionOffline)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetSessionOfflineArgs_Req_DEFAULT *status.TLStatusSetSessionOffline

func (p *SetSessionOfflineArgs) GetReq() *status.TLStatusSetSessionOffline {
	if !p.IsSetReq() {
		return SetSessionOfflineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetSessionOfflineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetSessionOfflineResult struct {
	Success *tg.Bool
}

var SetSessionOfflineResult_Success_DEFAULT *tg.Bool

func (p *SetSessionOfflineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetSessionOfflineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetSessionOfflineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetSessionOfflineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetSessionOfflineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetSessionOfflineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetSessionOfflineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetSessionOfflineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetSessionOfflineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetSessionOfflineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetSessionOfflineResult) GetResult() interface{} {
	return p.Success
}

func getUserOnlineSessionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserOnlineSessionsArgs)
	realResult := result.(*GetUserOnlineSessionsResult)
	success, err := handler.(status.RPCStatus).StatusGetUserOnlineSessions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserOnlineSessionsArgs() interface{} {
	return &GetUserOnlineSessionsArgs{}
}

func newGetUserOnlineSessionsResult() interface{} {
	return &GetUserOnlineSessionsResult{}
}

type GetUserOnlineSessionsArgs struct {
	Req *status.TLStatusGetUserOnlineSessions
}

func (p *GetUserOnlineSessionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserOnlineSessionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserOnlineSessionsArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusGetUserOnlineSessions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserOnlineSessionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserOnlineSessionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserOnlineSessionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusGetUserOnlineSessions)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserOnlineSessionsArgs_Req_DEFAULT *status.TLStatusGetUserOnlineSessions

func (p *GetUserOnlineSessionsArgs) GetReq() *status.TLStatusGetUserOnlineSessions {
	if !p.IsSetReq() {
		return GetUserOnlineSessionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserOnlineSessionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserOnlineSessionsResult struct {
	Success *status.UserSessionEntryList
}

var GetUserOnlineSessionsResult_Success_DEFAULT *status.UserSessionEntryList

func (p *GetUserOnlineSessionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserOnlineSessionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserOnlineSessionsResult) Unmarshal(in []byte) error {
	msg := new(status.UserSessionEntryList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserOnlineSessionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserOnlineSessionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserOnlineSessionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(status.UserSessionEntryList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserOnlineSessionsResult) GetSuccess() *status.UserSessionEntryList {
	if !p.IsSetSuccess() {
		return GetUserOnlineSessionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserOnlineSessionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*status.UserSessionEntryList)
}

func (p *GetUserOnlineSessionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserOnlineSessionsResult) GetResult() interface{} {
	return p.Success
}

func getUsersOnlineSessionsListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUsersOnlineSessionsListArgs)
	realResult := result.(*GetUsersOnlineSessionsListResult)
	success, err := handler.(status.RPCStatus).StatusGetUsersOnlineSessionsList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUsersOnlineSessionsListArgs() interface{} {
	return &GetUsersOnlineSessionsListArgs{}
}

func newGetUsersOnlineSessionsListResult() interface{} {
	return &GetUsersOnlineSessionsListResult{}
}

type GetUsersOnlineSessionsListArgs struct {
	Req *status.TLStatusGetUsersOnlineSessionsList
}

func (p *GetUsersOnlineSessionsListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUsersOnlineSessionsListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUsersOnlineSessionsListArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusGetUsersOnlineSessionsList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUsersOnlineSessionsListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUsersOnlineSessionsListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUsersOnlineSessionsListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusGetUsersOnlineSessionsList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUsersOnlineSessionsListArgs_Req_DEFAULT *status.TLStatusGetUsersOnlineSessionsList

func (p *GetUsersOnlineSessionsListArgs) GetReq() *status.TLStatusGetUsersOnlineSessionsList {
	if !p.IsSetReq() {
		return GetUsersOnlineSessionsListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUsersOnlineSessionsListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUsersOnlineSessionsListResult struct {
	Success *status.VectorUserSessionEntryList
}

var GetUsersOnlineSessionsListResult_Success_DEFAULT *status.VectorUserSessionEntryList

func (p *GetUsersOnlineSessionsListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUsersOnlineSessionsListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUsersOnlineSessionsListResult) Unmarshal(in []byte) error {
	msg := new(status.VectorUserSessionEntryList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUsersOnlineSessionsListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUsersOnlineSessionsListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUsersOnlineSessionsListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(status.VectorUserSessionEntryList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUsersOnlineSessionsListResult) GetSuccess() *status.VectorUserSessionEntryList {
	if !p.IsSetSuccess() {
		return GetUsersOnlineSessionsListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUsersOnlineSessionsListResult) SetSuccess(x interface{}) {
	p.Success = x.(*status.VectorUserSessionEntryList)
}

func (p *GetUsersOnlineSessionsListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUsersOnlineSessionsListResult) GetResult() interface{} {
	return p.Success
}

func getChannelOnlineUsersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetChannelOnlineUsersArgs)
	realResult := result.(*GetChannelOnlineUsersResult)
	success, err := handler.(status.RPCStatus).StatusGetChannelOnlineUsers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetChannelOnlineUsersArgs() interface{} {
	return &GetChannelOnlineUsersArgs{}
}

func newGetChannelOnlineUsersResult() interface{} {
	return &GetChannelOnlineUsersResult{}
}

type GetChannelOnlineUsersArgs struct {
	Req *status.TLStatusGetChannelOnlineUsers
}

func (p *GetChannelOnlineUsersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChannelOnlineUsersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetChannelOnlineUsersArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusGetChannelOnlineUsers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetChannelOnlineUsersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetChannelOnlineUsersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetChannelOnlineUsersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusGetChannelOnlineUsers)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetChannelOnlineUsersArgs_Req_DEFAULT *status.TLStatusGetChannelOnlineUsers

func (p *GetChannelOnlineUsersArgs) GetReq() *status.TLStatusGetChannelOnlineUsers {
	if !p.IsSetReq() {
		return GetChannelOnlineUsersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChannelOnlineUsersArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetChannelOnlineUsersResult struct {
	Success *status.VectorLong
}

var GetChannelOnlineUsersResult_Success_DEFAULT *status.VectorLong

func (p *GetChannelOnlineUsersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChannelOnlineUsersResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetChannelOnlineUsersResult) Unmarshal(in []byte) error {
	msg := new(status.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChannelOnlineUsersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetChannelOnlineUsersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetChannelOnlineUsersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(status.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChannelOnlineUsersResult) GetSuccess() *status.VectorLong {
	if !p.IsSetSuccess() {
		return GetChannelOnlineUsersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChannelOnlineUsersResult) SetSuccess(x interface{}) {
	p.Success = x.(*status.VectorLong)
}

func (p *GetChannelOnlineUsersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetChannelOnlineUsersResult) GetResult() interface{} {
	return p.Success
}

func setUserChannelsOnlineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetUserChannelsOnlineArgs)
	realResult := result.(*SetUserChannelsOnlineResult)
	success, err := handler.(status.RPCStatus).StatusSetUserChannelsOnline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetUserChannelsOnlineArgs() interface{} {
	return &SetUserChannelsOnlineArgs{}
}

func newSetUserChannelsOnlineResult() interface{} {
	return &SetUserChannelsOnlineResult{}
}

type SetUserChannelsOnlineArgs struct {
	Req *status.TLStatusSetUserChannelsOnline
}

func (p *SetUserChannelsOnlineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetUserChannelsOnlineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetUserChannelsOnlineArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusSetUserChannelsOnline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetUserChannelsOnlineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetUserChannelsOnlineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetUserChannelsOnlineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusSetUserChannelsOnline)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetUserChannelsOnlineArgs_Req_DEFAULT *status.TLStatusSetUserChannelsOnline

func (p *SetUserChannelsOnlineArgs) GetReq() *status.TLStatusSetUserChannelsOnline {
	if !p.IsSetReq() {
		return SetUserChannelsOnlineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetUserChannelsOnlineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetUserChannelsOnlineResult struct {
	Success *tg.Bool
}

var SetUserChannelsOnlineResult_Success_DEFAULT *tg.Bool

func (p *SetUserChannelsOnlineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetUserChannelsOnlineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetUserChannelsOnlineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetUserChannelsOnlineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetUserChannelsOnlineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetUserChannelsOnlineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetUserChannelsOnlineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetUserChannelsOnlineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetUserChannelsOnlineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetUserChannelsOnlineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetUserChannelsOnlineResult) GetResult() interface{} {
	return p.Success
}

func setUserChannelsOfflineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetUserChannelsOfflineArgs)
	realResult := result.(*SetUserChannelsOfflineResult)
	success, err := handler.(status.RPCStatus).StatusSetUserChannelsOffline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetUserChannelsOfflineArgs() interface{} {
	return &SetUserChannelsOfflineArgs{}
}

func newSetUserChannelsOfflineResult() interface{} {
	return &SetUserChannelsOfflineResult{}
}

type SetUserChannelsOfflineArgs struct {
	Req *status.TLStatusSetUserChannelsOffline
}

func (p *SetUserChannelsOfflineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetUserChannelsOfflineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetUserChannelsOfflineArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusSetUserChannelsOffline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetUserChannelsOfflineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetUserChannelsOfflineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetUserChannelsOfflineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusSetUserChannelsOffline)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetUserChannelsOfflineArgs_Req_DEFAULT *status.TLStatusSetUserChannelsOffline

func (p *SetUserChannelsOfflineArgs) GetReq() *status.TLStatusSetUserChannelsOffline {
	if !p.IsSetReq() {
		return SetUserChannelsOfflineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetUserChannelsOfflineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetUserChannelsOfflineResult struct {
	Success *tg.Bool
}

var SetUserChannelsOfflineResult_Success_DEFAULT *tg.Bool

func (p *SetUserChannelsOfflineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetUserChannelsOfflineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetUserChannelsOfflineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetUserChannelsOfflineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetUserChannelsOfflineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetUserChannelsOfflineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetUserChannelsOfflineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetUserChannelsOfflineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetUserChannelsOfflineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetUserChannelsOfflineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetUserChannelsOfflineResult) GetResult() interface{} {
	return p.Success
}

func setChannelUserOfflineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetChannelUserOfflineArgs)
	realResult := result.(*SetChannelUserOfflineResult)
	success, err := handler.(status.RPCStatus).StatusSetChannelUserOffline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetChannelUserOfflineArgs() interface{} {
	return &SetChannelUserOfflineArgs{}
}

func newSetChannelUserOfflineResult() interface{} {
	return &SetChannelUserOfflineResult{}
}

type SetChannelUserOfflineArgs struct {
	Req *status.TLStatusSetChannelUserOffline
}

func (p *SetChannelUserOfflineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetChannelUserOfflineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetChannelUserOfflineArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusSetChannelUserOffline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetChannelUserOfflineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetChannelUserOfflineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetChannelUserOfflineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusSetChannelUserOffline)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetChannelUserOfflineArgs_Req_DEFAULT *status.TLStatusSetChannelUserOffline

func (p *SetChannelUserOfflineArgs) GetReq() *status.TLStatusSetChannelUserOffline {
	if !p.IsSetReq() {
		return SetChannelUserOfflineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetChannelUserOfflineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetChannelUserOfflineResult struct {
	Success *tg.Bool
}

var SetChannelUserOfflineResult_Success_DEFAULT *tg.Bool

func (p *SetChannelUserOfflineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetChannelUserOfflineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetChannelUserOfflineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChannelUserOfflineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetChannelUserOfflineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetChannelUserOfflineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChannelUserOfflineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetChannelUserOfflineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetChannelUserOfflineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetChannelUserOfflineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetChannelUserOfflineResult) GetResult() interface{} {
	return p.Success
}

func setChannelUsersOnlineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetChannelUsersOnlineArgs)
	realResult := result.(*SetChannelUsersOnlineResult)
	success, err := handler.(status.RPCStatus).StatusSetChannelUsersOnline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetChannelUsersOnlineArgs() interface{} {
	return &SetChannelUsersOnlineArgs{}
}

func newSetChannelUsersOnlineResult() interface{} {
	return &SetChannelUsersOnlineResult{}
}

type SetChannelUsersOnlineArgs struct {
	Req *status.TLStatusSetChannelUsersOnline
}

func (p *SetChannelUsersOnlineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetChannelUsersOnlineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetChannelUsersOnlineArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusSetChannelUsersOnline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetChannelUsersOnlineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetChannelUsersOnlineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetChannelUsersOnlineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusSetChannelUsersOnline)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetChannelUsersOnlineArgs_Req_DEFAULT *status.TLStatusSetChannelUsersOnline

func (p *SetChannelUsersOnlineArgs) GetReq() *status.TLStatusSetChannelUsersOnline {
	if !p.IsSetReq() {
		return SetChannelUsersOnlineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetChannelUsersOnlineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetChannelUsersOnlineResult struct {
	Success *tg.Bool
}

var SetChannelUsersOnlineResult_Success_DEFAULT *tg.Bool

func (p *SetChannelUsersOnlineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetChannelUsersOnlineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetChannelUsersOnlineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChannelUsersOnlineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetChannelUsersOnlineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetChannelUsersOnlineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChannelUsersOnlineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetChannelUsersOnlineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetChannelUsersOnlineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetChannelUsersOnlineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetChannelUsersOnlineResult) GetResult() interface{} {
	return p.Success
}

func setChannelOfflineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetChannelOfflineArgs)
	realResult := result.(*SetChannelOfflineResult)
	success, err := handler.(status.RPCStatus).StatusSetChannelOffline(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetChannelOfflineArgs() interface{} {
	return &SetChannelOfflineArgs{}
}

func newSetChannelOfflineResult() interface{} {
	return &SetChannelOfflineResult{}
}

type SetChannelOfflineArgs struct {
	Req *status.TLStatusSetChannelOffline
}

func (p *SetChannelOfflineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetChannelOfflineArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetChannelOfflineArgs) Unmarshal(in []byte) error {
	msg := new(status.TLStatusSetChannelOffline)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetChannelOfflineArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetChannelOfflineArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetChannelOfflineArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(status.TLStatusSetChannelOffline)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetChannelOfflineArgs_Req_DEFAULT *status.TLStatusSetChannelOffline

func (p *SetChannelOfflineArgs) GetReq() *status.TLStatusSetChannelOffline {
	if !p.IsSetReq() {
		return SetChannelOfflineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetChannelOfflineArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetChannelOfflineResult struct {
	Success *tg.Bool
}

var SetChannelOfflineResult_Success_DEFAULT *tg.Bool

func (p *SetChannelOfflineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetChannelOfflineResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetChannelOfflineResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChannelOfflineResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetChannelOfflineResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetChannelOfflineResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChannelOfflineResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetChannelOfflineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetChannelOfflineResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetChannelOfflineResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetChannelOfflineResult) GetResult() interface{} {
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

func (p *kClient) StatusSetSessionOnline(ctx context.Context, req *status.TLStatusSetSessionOnline) (r *tg.Bool, err error) {
	var _args SetSessionOnlineArgs
	_args.Req = req
	var _result SetSessionOnlineResult
	if err = p.c.Call(ctx, "status.setSessionOnline", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusSetSessionOffline(ctx context.Context, req *status.TLStatusSetSessionOffline) (r *tg.Bool, err error) {
	var _args SetSessionOfflineArgs
	_args.Req = req
	var _result SetSessionOfflineResult
	if err = p.c.Call(ctx, "status.setSessionOffline", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusGetUserOnlineSessions(ctx context.Context, req *status.TLStatusGetUserOnlineSessions) (r *status.UserSessionEntryList, err error) {
	var _args GetUserOnlineSessionsArgs
	_args.Req = req
	var _result GetUserOnlineSessionsResult
	if err = p.c.Call(ctx, "status.getUserOnlineSessions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusGetUsersOnlineSessionsList(ctx context.Context, req *status.TLStatusGetUsersOnlineSessionsList) (r *status.VectorUserSessionEntryList, err error) {
	var _args GetUsersOnlineSessionsListArgs
	_args.Req = req
	var _result GetUsersOnlineSessionsListResult
	if err = p.c.Call(ctx, "status.getUsersOnlineSessionsList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusGetChannelOnlineUsers(ctx context.Context, req *status.TLStatusGetChannelOnlineUsers) (r *status.VectorLong, err error) {
	var _args GetChannelOnlineUsersArgs
	_args.Req = req
	var _result GetChannelOnlineUsersResult
	if err = p.c.Call(ctx, "status.getChannelOnlineUsers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusSetUserChannelsOnline(ctx context.Context, req *status.TLStatusSetUserChannelsOnline) (r *tg.Bool, err error) {
	var _args SetUserChannelsOnlineArgs
	_args.Req = req
	var _result SetUserChannelsOnlineResult
	if err = p.c.Call(ctx, "status.setUserChannelsOnline", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusSetUserChannelsOffline(ctx context.Context, req *status.TLStatusSetUserChannelsOffline) (r *tg.Bool, err error) {
	var _args SetUserChannelsOfflineArgs
	_args.Req = req
	var _result SetUserChannelsOfflineResult
	if err = p.c.Call(ctx, "status.setUserChannelsOffline", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusSetChannelUserOffline(ctx context.Context, req *status.TLStatusSetChannelUserOffline) (r *tg.Bool, err error) {
	var _args SetChannelUserOfflineArgs
	_args.Req = req
	var _result SetChannelUserOfflineResult
	if err = p.c.Call(ctx, "status.setChannelUserOffline", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusSetChannelUsersOnline(ctx context.Context, req *status.TLStatusSetChannelUsersOnline) (r *tg.Bool, err error) {
	var _args SetChannelUsersOnlineArgs
	_args.Req = req
	var _result SetChannelUsersOnlineResult
	if err = p.c.Call(ctx, "status.setChannelUsersOnline", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) StatusSetChannelOffline(ctx context.Context, req *status.TLStatusSetChannelOffline) (r *tg.Bool, err error) {
	var _args SetChannelOfflineArgs
	_args.Req = req
	var _result SetChannelOfflineResult
	if err = p.c.Call(ctx, "status.setChannelOffline", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
