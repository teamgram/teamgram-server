/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package usernameservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/username/username"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/username.RPCUsername/username.getAccountUsername": kitex.NewMethodInfo(
		getAccountUsernameHandler,
		newGetAccountUsernameArgs,
		newGetAccountUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.checkAccountUsername": kitex.NewMethodInfo(
		checkAccountUsernameHandler,
		newCheckAccountUsernameArgs,
		newCheckAccountUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.getChannelUsername": kitex.NewMethodInfo(
		getChannelUsernameHandler,
		newGetChannelUsernameArgs,
		newGetChannelUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.checkChannelUsername": kitex.NewMethodInfo(
		checkChannelUsernameHandler,
		newCheckChannelUsernameArgs,
		newCheckChannelUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.updateUsernameByPeer": kitex.NewMethodInfo(
		updateUsernameByPeerHandler,
		newUpdateUsernameByPeerArgs,
		newUpdateUsernameByPeerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.checkUsername": kitex.NewMethodInfo(
		checkUsernameHandler,
		newCheckUsernameArgs,
		newCheckUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.updateUsername": kitex.NewMethodInfo(
		updateUsernameHandler,
		newUpdateUsernameArgs,
		newUpdateUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.deleteUsername": kitex.NewMethodInfo(
		deleteUsernameHandler,
		newDeleteUsernameArgs,
		newDeleteUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.resolveUsername": kitex.NewMethodInfo(
		resolveUsernameHandler,
		newResolveUsernameArgs,
		newResolveUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.getListByUsernameList": kitex.NewMethodInfo(
		getListByUsernameListHandler,
		newGetListByUsernameListArgs,
		newGetListByUsernameListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.deleteUsernameByPeer": kitex.NewMethodInfo(
		deleteUsernameByPeerHandler,
		newDeleteUsernameByPeerArgs,
		newDeleteUsernameByPeerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/username.RPCUsername/username.search": kitex.NewMethodInfo(
		searchHandler,
		newSearchArgs,
		newSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	usernameServiceServiceInfo                = NewServiceInfo()
	usernameServiceServiceInfoForClient       = NewServiceInfoForClient()
	usernameServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCUsername", usernameServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCUsername", usernameServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCUsername", usernameServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return usernameServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return usernameServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return usernameServiceServiceInfoForClient
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
	serviceName := "RPCUsername"
	handlerType := (*username.RPCUsername)(nil)
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
		"PackageName": "username",
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

func getAccountUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAccountUsernameArgs)
	realResult := result.(*GetAccountUsernameResult)
	success, err := handler.(username.RPCUsername).UsernameGetAccountUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAccountUsernameArgs() interface{} {
	return &GetAccountUsernameArgs{}
}

func newGetAccountUsernameResult() interface{} {
	return &GetAccountUsernameResult{}
}

type GetAccountUsernameArgs struct {
	Req *username.TLUsernameGetAccountUsername
}

func (p *GetAccountUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAccountUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAccountUsernameArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameGetAccountUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAccountUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAccountUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAccountUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameGetAccountUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAccountUsernameArgs_Req_DEFAULT *username.TLUsernameGetAccountUsername

func (p *GetAccountUsernameArgs) GetReq() *username.TLUsernameGetAccountUsername {
	if !p.IsSetReq() {
		return GetAccountUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAccountUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAccountUsernameResult struct {
	Success *username.UsernameData
}

var GetAccountUsernameResult_Success_DEFAULT *username.UsernameData

func (p *GetAccountUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAccountUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAccountUsernameResult) Unmarshal(in []byte) error {
	msg := new(username.UsernameData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAccountUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAccountUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAccountUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(username.UsernameData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAccountUsernameResult) GetSuccess() *username.UsernameData {
	if !p.IsSetSuccess() {
		return GetAccountUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAccountUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*username.UsernameData)
}

func (p *GetAccountUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAccountUsernameResult) GetResult() interface{} {
	return p.Success
}

func checkAccountUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CheckAccountUsernameArgs)
	realResult := result.(*CheckAccountUsernameResult)
	success, err := handler.(username.RPCUsername).UsernameCheckAccountUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCheckAccountUsernameArgs() interface{} {
	return &CheckAccountUsernameArgs{}
}

func newCheckAccountUsernameResult() interface{} {
	return &CheckAccountUsernameResult{}
}

type CheckAccountUsernameArgs struct {
	Req *username.TLUsernameCheckAccountUsername
}

func (p *CheckAccountUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckAccountUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CheckAccountUsernameArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameCheckAccountUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CheckAccountUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CheckAccountUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CheckAccountUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameCheckAccountUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CheckAccountUsernameArgs_Req_DEFAULT *username.TLUsernameCheckAccountUsername

func (p *CheckAccountUsernameArgs) GetReq() *username.TLUsernameCheckAccountUsername {
	if !p.IsSetReq() {
		return CheckAccountUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckAccountUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckAccountUsernameResult struct {
	Success *username.UsernameExist
}

var CheckAccountUsernameResult_Success_DEFAULT *username.UsernameExist

func (p *CheckAccountUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckAccountUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *CheckAccountUsernameResult) Unmarshal(in []byte) error {
	msg := new(username.UsernameExist)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckAccountUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CheckAccountUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CheckAccountUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(username.UsernameExist)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckAccountUsernameResult) GetSuccess() *username.UsernameExist {
	if !p.IsSetSuccess() {
		return CheckAccountUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckAccountUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*username.UsernameExist)
}

func (p *CheckAccountUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckAccountUsernameResult) GetResult() interface{} {
	return p.Success
}

func getChannelUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetChannelUsernameArgs)
	realResult := result.(*GetChannelUsernameResult)
	success, err := handler.(username.RPCUsername).UsernameGetChannelUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetChannelUsernameArgs() interface{} {
	return &GetChannelUsernameArgs{}
}

func newGetChannelUsernameResult() interface{} {
	return &GetChannelUsernameResult{}
}

type GetChannelUsernameArgs struct {
	Req *username.TLUsernameGetChannelUsername
}

func (p *GetChannelUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChannelUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetChannelUsernameArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameGetChannelUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetChannelUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetChannelUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetChannelUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameGetChannelUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetChannelUsernameArgs_Req_DEFAULT *username.TLUsernameGetChannelUsername

func (p *GetChannelUsernameArgs) GetReq() *username.TLUsernameGetChannelUsername {
	if !p.IsSetReq() {
		return GetChannelUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChannelUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetChannelUsernameResult struct {
	Success *username.UsernameData
}

var GetChannelUsernameResult_Success_DEFAULT *username.UsernameData

func (p *GetChannelUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChannelUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetChannelUsernameResult) Unmarshal(in []byte) error {
	msg := new(username.UsernameData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChannelUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetChannelUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetChannelUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(username.UsernameData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChannelUsernameResult) GetSuccess() *username.UsernameData {
	if !p.IsSetSuccess() {
		return GetChannelUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChannelUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*username.UsernameData)
}

func (p *GetChannelUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetChannelUsernameResult) GetResult() interface{} {
	return p.Success
}

func checkChannelUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CheckChannelUsernameArgs)
	realResult := result.(*CheckChannelUsernameResult)
	success, err := handler.(username.RPCUsername).UsernameCheckChannelUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCheckChannelUsernameArgs() interface{} {
	return &CheckChannelUsernameArgs{}
}

func newCheckChannelUsernameResult() interface{} {
	return &CheckChannelUsernameResult{}
}

type CheckChannelUsernameArgs struct {
	Req *username.TLUsernameCheckChannelUsername
}

func (p *CheckChannelUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckChannelUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CheckChannelUsernameArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameCheckChannelUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CheckChannelUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CheckChannelUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CheckChannelUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameCheckChannelUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CheckChannelUsernameArgs_Req_DEFAULT *username.TLUsernameCheckChannelUsername

func (p *CheckChannelUsernameArgs) GetReq() *username.TLUsernameCheckChannelUsername {
	if !p.IsSetReq() {
		return CheckChannelUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckChannelUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckChannelUsernameResult struct {
	Success *username.UsernameExist
}

var CheckChannelUsernameResult_Success_DEFAULT *username.UsernameExist

func (p *CheckChannelUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckChannelUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *CheckChannelUsernameResult) Unmarshal(in []byte) error {
	msg := new(username.UsernameExist)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckChannelUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CheckChannelUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CheckChannelUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(username.UsernameExist)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckChannelUsernameResult) GetSuccess() *username.UsernameExist {
	if !p.IsSetSuccess() {
		return CheckChannelUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckChannelUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*username.UsernameExist)
}

func (p *CheckChannelUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckChannelUsernameResult) GetResult() interface{} {
	return p.Success
}

func updateUsernameByPeerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateUsernameByPeerArgs)
	realResult := result.(*UpdateUsernameByPeerResult)
	success, err := handler.(username.RPCUsername).UsernameUpdateUsernameByPeer(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateUsernameByPeerArgs() interface{} {
	return &UpdateUsernameByPeerArgs{}
}

func newUpdateUsernameByPeerResult() interface{} {
	return &UpdateUsernameByPeerResult{}
}

type UpdateUsernameByPeerArgs struct {
	Req *username.TLUsernameUpdateUsernameByPeer
}

func (p *UpdateUsernameByPeerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateUsernameByPeerArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateUsernameByPeerArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameUpdateUsernameByPeer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateUsernameByPeerArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateUsernameByPeerArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateUsernameByPeerArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameUpdateUsernameByPeer)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateUsernameByPeerArgs_Req_DEFAULT *username.TLUsernameUpdateUsernameByPeer

func (p *UpdateUsernameByPeerArgs) GetReq() *username.TLUsernameUpdateUsernameByPeer {
	if !p.IsSetReq() {
		return UpdateUsernameByPeerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUsernameByPeerArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateUsernameByPeerResult struct {
	Success *tg.Bool
}

var UpdateUsernameByPeerResult_Success_DEFAULT *tg.Bool

func (p *UpdateUsernameByPeerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateUsernameByPeerResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateUsernameByPeerResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUsernameByPeerResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateUsernameByPeerResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateUsernameByPeerResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUsernameByPeerResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateUsernameByPeerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUsernameByPeerResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateUsernameByPeerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUsernameByPeerResult) GetResult() interface{} {
	return p.Success
}

func checkUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CheckUsernameArgs)
	realResult := result.(*CheckUsernameResult)
	success, err := handler.(username.RPCUsername).UsernameCheckUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCheckUsernameArgs() interface{} {
	return &CheckUsernameArgs{}
}

func newCheckUsernameResult() interface{} {
	return &CheckUsernameResult{}
}

type CheckUsernameArgs struct {
	Req *username.TLUsernameCheckUsername
}

func (p *CheckUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CheckUsernameArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameCheckUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CheckUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CheckUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CheckUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameCheckUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CheckUsernameArgs_Req_DEFAULT *username.TLUsernameCheckUsername

func (p *CheckUsernameArgs) GetReq() *username.TLUsernameCheckUsername {
	if !p.IsSetReq() {
		return CheckUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckUsernameResult struct {
	Success *username.UsernameExist
}

var CheckUsernameResult_Success_DEFAULT *username.UsernameExist

func (p *CheckUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *CheckUsernameResult) Unmarshal(in []byte) error {
	msg := new(username.UsernameExist)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CheckUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CheckUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(username.UsernameExist)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUsernameResult) GetSuccess() *username.UsernameExist {
	if !p.IsSetSuccess() {
		return CheckUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*username.UsernameExist)
}

func (p *CheckUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckUsernameResult) GetResult() interface{} {
	return p.Success
}

func updateUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateUsernameArgs)
	realResult := result.(*UpdateUsernameResult)
	success, err := handler.(username.RPCUsername).UsernameUpdateUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateUsernameArgs() interface{} {
	return &UpdateUsernameArgs{}
}

func newUpdateUsernameResult() interface{} {
	return &UpdateUsernameResult{}
}

type UpdateUsernameArgs struct {
	Req *username.TLUsernameUpdateUsername
}

func (p *UpdateUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateUsernameArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameUpdateUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameUpdateUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateUsernameArgs_Req_DEFAULT *username.TLUsernameUpdateUsername

func (p *UpdateUsernameArgs) GetReq() *username.TLUsernameUpdateUsername {
	if !p.IsSetReq() {
		return UpdateUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateUsernameResult struct {
	Success *tg.Bool
}

var UpdateUsernameResult_Success_DEFAULT *tg.Bool

func (p *UpdateUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUsernameResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUsernameResult) GetResult() interface{} {
	return p.Success
}

func deleteUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteUsernameArgs)
	realResult := result.(*DeleteUsernameResult)
	success, err := handler.(username.RPCUsername).UsernameDeleteUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteUsernameArgs() interface{} {
	return &DeleteUsernameArgs{}
}

func newDeleteUsernameResult() interface{} {
	return &DeleteUsernameResult{}
}

type DeleteUsernameArgs struct {
	Req *username.TLUsernameDeleteUsername
}

func (p *DeleteUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteUsernameArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameDeleteUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameDeleteUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteUsernameArgs_Req_DEFAULT *username.TLUsernameDeleteUsername

func (p *DeleteUsernameArgs) GetReq() *username.TLUsernameDeleteUsername {
	if !p.IsSetReq() {
		return DeleteUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteUsernameResult struct {
	Success *tg.Bool
}

var DeleteUsernameResult_Success_DEFAULT *tg.Bool

func (p *DeleteUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUsernameResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteUsernameResult) GetResult() interface{} {
	return p.Success
}

func resolveUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ResolveUsernameArgs)
	realResult := result.(*ResolveUsernameResult)
	success, err := handler.(username.RPCUsername).UsernameResolveUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newResolveUsernameArgs() interface{} {
	return &ResolveUsernameArgs{}
}

func newResolveUsernameResult() interface{} {
	return &ResolveUsernameResult{}
}

type ResolveUsernameArgs struct {
	Req *username.TLUsernameResolveUsername
}

func (p *ResolveUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ResolveUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ResolveUsernameArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameResolveUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ResolveUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ResolveUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ResolveUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameResolveUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ResolveUsernameArgs_Req_DEFAULT *username.TLUsernameResolveUsername

func (p *ResolveUsernameArgs) GetReq() *username.TLUsernameResolveUsername {
	if !p.IsSetReq() {
		return ResolveUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ResolveUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type ResolveUsernameResult struct {
	Success *tg.Peer
}

var ResolveUsernameResult_Success_DEFAULT *tg.Peer

func (p *ResolveUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ResolveUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *ResolveUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.Peer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ResolveUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ResolveUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ResolveUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Peer)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ResolveUsernameResult) GetSuccess() *tg.Peer {
	if !p.IsSetSuccess() {
		return ResolveUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ResolveUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Peer)
}

func (p *ResolveUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ResolveUsernameResult) GetResult() interface{} {
	return p.Success
}

func getListByUsernameListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetListByUsernameListArgs)
	realResult := result.(*GetListByUsernameListResult)
	success, err := handler.(username.RPCUsername).UsernameGetListByUsernameList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetListByUsernameListArgs() interface{} {
	return &GetListByUsernameListArgs{}
}

func newGetListByUsernameListResult() interface{} {
	return &GetListByUsernameListResult{}
}

type GetListByUsernameListArgs struct {
	Req *username.TLUsernameGetListByUsernameList
}

func (p *GetListByUsernameListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetListByUsernameListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetListByUsernameListArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameGetListByUsernameList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetListByUsernameListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetListByUsernameListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetListByUsernameListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameGetListByUsernameList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetListByUsernameListArgs_Req_DEFAULT *username.TLUsernameGetListByUsernameList

func (p *GetListByUsernameListArgs) GetReq() *username.TLUsernameGetListByUsernameList {
	if !p.IsSetReq() {
		return GetListByUsernameListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetListByUsernameListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetListByUsernameListResult struct {
	Success *username.VectorUsernameData
}

var GetListByUsernameListResult_Success_DEFAULT *username.VectorUsernameData

func (p *GetListByUsernameListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetListByUsernameListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetListByUsernameListResult) Unmarshal(in []byte) error {
	msg := new(username.VectorUsernameData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetListByUsernameListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetListByUsernameListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetListByUsernameListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(username.VectorUsernameData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetListByUsernameListResult) GetSuccess() *username.VectorUsernameData {
	if !p.IsSetSuccess() {
		return GetListByUsernameListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetListByUsernameListResult) SetSuccess(x interface{}) {
	p.Success = x.(*username.VectorUsernameData)
}

func (p *GetListByUsernameListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetListByUsernameListResult) GetResult() interface{} {
	return p.Success
}

func deleteUsernameByPeerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteUsernameByPeerArgs)
	realResult := result.(*DeleteUsernameByPeerResult)
	success, err := handler.(username.RPCUsername).UsernameDeleteUsernameByPeer(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteUsernameByPeerArgs() interface{} {
	return &DeleteUsernameByPeerArgs{}
}

func newDeleteUsernameByPeerResult() interface{} {
	return &DeleteUsernameByPeerResult{}
}

type DeleteUsernameByPeerArgs struct {
	Req *username.TLUsernameDeleteUsernameByPeer
}

func (p *DeleteUsernameByPeerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteUsernameByPeerArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteUsernameByPeerArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameDeleteUsernameByPeer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteUsernameByPeerArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteUsernameByPeerArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteUsernameByPeerArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(username.TLUsernameDeleteUsernameByPeer)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteUsernameByPeerArgs_Req_DEFAULT *username.TLUsernameDeleteUsernameByPeer

func (p *DeleteUsernameByPeerArgs) GetReq() *username.TLUsernameDeleteUsernameByPeer {
	if !p.IsSetReq() {
		return DeleteUsernameByPeerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteUsernameByPeerArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteUsernameByPeerResult struct {
	Success *tg.Bool
}

var DeleteUsernameByPeerResult_Success_DEFAULT *tg.Bool

func (p *DeleteUsernameByPeerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteUsernameByPeerResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteUsernameByPeerResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUsernameByPeerResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteUsernameByPeerResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteUsernameByPeerResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUsernameByPeerResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteUsernameByPeerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteUsernameByPeerResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteUsernameByPeerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteUsernameByPeerResult) GetResult() interface{} {
	return p.Success
}

func searchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SearchArgs)
	realResult := result.(*SearchResult)
	success, err := handler.(username.RPCUsername).UsernameSearch(ctx, realArg.Req)
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
	Req *username.TLUsernameSearch
}

func (p *SearchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SearchArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SearchArgs) Unmarshal(in []byte) error {
	msg := new(username.TLUsernameSearch)
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
	msg := new(username.TLUsernameSearch)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SearchArgs_Req_DEFAULT *username.TLUsernameSearch

func (p *SearchArgs) GetReq() *username.TLUsernameSearch {
	if !p.IsSetReq() {
		return SearchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchArgs) IsSetReq() bool {
	return p.Req != nil
}

type SearchResult struct {
	Success *username.VectorUsernameData
}

var SearchResult_Success_DEFAULT *username.VectorUsernameData

func (p *SearchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SearchResult")
	}
	return json.Marshal(p.Success)
}

func (p *SearchResult) Unmarshal(in []byte) error {
	msg := new(username.VectorUsernameData)
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
	msg := new(username.VectorUsernameData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchResult) GetSuccess() *username.VectorUsernameData {
	if !p.IsSetSuccess() {
		return SearchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchResult) SetSuccess(x interface{}) {
	p.Success = x.(*username.VectorUsernameData)
}

func (p *SearchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchResult) GetResult() interface{} {
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

func (p *kClient) UsernameGetAccountUsername(ctx context.Context, req *username.TLUsernameGetAccountUsername) (r *username.UsernameData, err error) {
	// var _args GetAccountUsernameArgs
	// _args.Req = req
	// var _result GetAccountUsernameResult

	_result := new(username.UsernameData)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.getAccountUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameCheckAccountUsername(ctx context.Context, req *username.TLUsernameCheckAccountUsername) (r *username.UsernameExist, err error) {
	// var _args CheckAccountUsernameArgs
	// _args.Req = req
	// var _result CheckAccountUsernameResult

	_result := new(username.UsernameExist)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.checkAccountUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameGetChannelUsername(ctx context.Context, req *username.TLUsernameGetChannelUsername) (r *username.UsernameData, err error) {
	// var _args GetChannelUsernameArgs
	// _args.Req = req
	// var _result GetChannelUsernameResult

	_result := new(username.UsernameData)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.getChannelUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameCheckChannelUsername(ctx context.Context, req *username.TLUsernameCheckChannelUsername) (r *username.UsernameExist, err error) {
	// var _args CheckChannelUsernameArgs
	// _args.Req = req
	// var _result CheckChannelUsernameResult

	_result := new(username.UsernameExist)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.checkChannelUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameUpdateUsernameByPeer(ctx context.Context, req *username.TLUsernameUpdateUsernameByPeer) (r *tg.Bool, err error) {
	// var _args UpdateUsernameByPeerArgs
	// _args.Req = req
	// var _result UpdateUsernameByPeerResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.updateUsernameByPeer", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameCheckUsername(ctx context.Context, req *username.TLUsernameCheckUsername) (r *username.UsernameExist, err error) {
	// var _args CheckUsernameArgs
	// _args.Req = req
	// var _result CheckUsernameResult

	_result := new(username.UsernameExist)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.checkUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameUpdateUsername(ctx context.Context, req *username.TLUsernameUpdateUsername) (r *tg.Bool, err error) {
	// var _args UpdateUsernameArgs
	// _args.Req = req
	// var _result UpdateUsernameResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.updateUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameDeleteUsername(ctx context.Context, req *username.TLUsernameDeleteUsername) (r *tg.Bool, err error) {
	// var _args DeleteUsernameArgs
	// _args.Req = req
	// var _result DeleteUsernameResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.deleteUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameResolveUsername(ctx context.Context, req *username.TLUsernameResolveUsername) (r *tg.Peer, err error) {
	// var _args ResolveUsernameArgs
	// _args.Req = req
	// var _result ResolveUsernameResult

	_result := new(tg.Peer)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.resolveUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameGetListByUsernameList(ctx context.Context, req *username.TLUsernameGetListByUsernameList) (r *username.VectorUsernameData, err error) {
	// var _args GetListByUsernameListArgs
	// _args.Req = req
	// var _result GetListByUsernameListResult

	_result := new(username.VectorUsernameData)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.getListByUsernameList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameDeleteUsernameByPeer(ctx context.Context, req *username.TLUsernameDeleteUsernameByPeer) (r *tg.Bool, err error) {
	// var _args DeleteUsernameByPeerArgs
	// _args.Req = req
	// var _result DeleteUsernameByPeerResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.deleteUsernameByPeer", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UsernameSearch(ctx context.Context, req *username.TLUsernameSearch) (r *username.VectorUsernameData, err error) {
	// var _args SearchArgs
	// _args.Req = req
	// var _result SearchResult

	_result := new(username.VectorUsernameData)

	if err = p.c.Call(ctx, "/username.RPCUsername/username.search", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
