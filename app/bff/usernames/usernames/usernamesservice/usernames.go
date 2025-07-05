/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package usernamesservice

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
	"/tg.RPCUsernames/account.checkUsername": kitex.NewMethodInfo(
		accountCheckUsernameHandler,
		newAccountCheckUsernameArgs,
		newAccountCheckUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCUsernames/account.updateUsername": kitex.NewMethodInfo(
		accountUpdateUsernameHandler,
		newAccountUpdateUsernameArgs,
		newAccountUpdateUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCUsernames/contacts.resolveUsername": kitex.NewMethodInfo(
		contactsResolveUsernameHandler,
		newContactsResolveUsernameArgs,
		newContactsResolveUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCUsernames/channels.checkUsername": kitex.NewMethodInfo(
		channelsCheckUsernameHandler,
		newChannelsCheckUsernameArgs,
		newChannelsCheckUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCUsernames/channels.updateUsername": kitex.NewMethodInfo(
		channelsUpdateUsernameHandler,
		newChannelsUpdateUsernameArgs,
		newChannelsUpdateUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	usernamesServiceServiceInfo                = NewServiceInfo()
	usernamesServiceServiceInfoForClient       = NewServiceInfoForClient()
	usernamesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCUsernames", usernamesServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCUsernames", usernamesServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCUsernames", usernamesServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return usernamesServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return usernamesServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return usernamesServiceServiceInfoForClient
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
	serviceName := "RPCUsernames"
	handlerType := (*tg.RPCUsernames)(nil)
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
		"PackageName": "usernames",
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

func accountCheckUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountCheckUsernameArgs)
	realResult := result.(*AccountCheckUsernameResult)
	success, err := handler.(tg.RPCUsernames).AccountCheckUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountCheckUsernameArgs() interface{} {
	return &AccountCheckUsernameArgs{}
}

func newAccountCheckUsernameResult() interface{} {
	return &AccountCheckUsernameResult{}
}

type AccountCheckUsernameArgs struct {
	Req *tg.TLAccountCheckUsername
}

func (p *AccountCheckUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountCheckUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountCheckUsernameArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountCheckUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountCheckUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountCheckUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountCheckUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountCheckUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountCheckUsernameArgs_Req_DEFAULT *tg.TLAccountCheckUsername

func (p *AccountCheckUsernameArgs) GetReq() *tg.TLAccountCheckUsername {
	if !p.IsSetReq() {
		return AccountCheckUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountCheckUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountCheckUsernameResult struct {
	Success *tg.Bool
}

var AccountCheckUsernameResult_Success_DEFAULT *tg.Bool

func (p *AccountCheckUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountCheckUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountCheckUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountCheckUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountCheckUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountCheckUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountCheckUsernameResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountCheckUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountCheckUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountCheckUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountCheckUsernameResult) GetResult() interface{} {
	return p.Success
}

func accountUpdateUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUpdateUsernameArgs)
	realResult := result.(*AccountUpdateUsernameResult)
	success, err := handler.(tg.RPCUsernames).AccountUpdateUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUpdateUsernameArgs() interface{} {
	return &AccountUpdateUsernameArgs{}
}

func newAccountUpdateUsernameResult() interface{} {
	return &AccountUpdateUsernameResult{}
}

type AccountUpdateUsernameArgs struct {
	Req *tg.TLAccountUpdateUsername
}

func (p *AccountUpdateUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUpdateUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUpdateUsernameArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUpdateUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUpdateUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUpdateUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUpdateUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUpdateUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUpdateUsernameArgs_Req_DEFAULT *tg.TLAccountUpdateUsername

func (p *AccountUpdateUsernameArgs) GetReq() *tg.TLAccountUpdateUsername {
	if !p.IsSetReq() {
		return AccountUpdateUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUpdateUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUpdateUsernameResult struct {
	Success *tg.User
}

var AccountUpdateUsernameResult_Success_DEFAULT *tg.User

func (p *AccountUpdateUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUpdateUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUpdateUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.User)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUpdateUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUpdateUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.User)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateUsernameResult) GetSuccess() *tg.User {
	if !p.IsSetSuccess() {
		return AccountUpdateUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUpdateUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.User)
}

func (p *AccountUpdateUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUpdateUsernameResult) GetResult() interface{} {
	return p.Success
}

func contactsResolveUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsResolveUsernameArgs)
	realResult := result.(*ContactsResolveUsernameResult)
	success, err := handler.(tg.RPCUsernames).ContactsResolveUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsResolveUsernameArgs() interface{} {
	return &ContactsResolveUsernameArgs{}
}

func newContactsResolveUsernameResult() interface{} {
	return &ContactsResolveUsernameResult{}
}

type ContactsResolveUsernameArgs struct {
	Req *tg.TLContactsResolveUsername
}

func (p *ContactsResolveUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsResolveUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsResolveUsernameArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsResolveUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsResolveUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsResolveUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsResolveUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsResolveUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsResolveUsernameArgs_Req_DEFAULT *tg.TLContactsResolveUsername

func (p *ContactsResolveUsernameArgs) GetReq() *tg.TLContactsResolveUsername {
	if !p.IsSetReq() {
		return ContactsResolveUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsResolveUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsResolveUsernameResult struct {
	Success *tg.ContactsResolvedPeer
}

var ContactsResolveUsernameResult_Success_DEFAULT *tg.ContactsResolvedPeer

func (p *ContactsResolveUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsResolveUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsResolveUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsResolvedPeer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsResolveUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsResolveUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsResolveUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsResolvedPeer)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsResolveUsernameResult) GetSuccess() *tg.ContactsResolvedPeer {
	if !p.IsSetSuccess() {
		return ContactsResolveUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsResolveUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsResolvedPeer)
}

func (p *ContactsResolveUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsResolveUsernameResult) GetResult() interface{} {
	return p.Success
}

func channelsCheckUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsCheckUsernameArgs)
	realResult := result.(*ChannelsCheckUsernameResult)
	success, err := handler.(tg.RPCUsernames).ChannelsCheckUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsCheckUsernameArgs() interface{} {
	return &ChannelsCheckUsernameArgs{}
}

func newChannelsCheckUsernameResult() interface{} {
	return &ChannelsCheckUsernameResult{}
}

type ChannelsCheckUsernameArgs struct {
	Req *tg.TLChannelsCheckUsername
}

func (p *ChannelsCheckUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsCheckUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsCheckUsernameArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsCheckUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsCheckUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsCheckUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsCheckUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsCheckUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsCheckUsernameArgs_Req_DEFAULT *tg.TLChannelsCheckUsername

func (p *ChannelsCheckUsernameArgs) GetReq() *tg.TLChannelsCheckUsername {
	if !p.IsSetReq() {
		return ChannelsCheckUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsCheckUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsCheckUsernameResult struct {
	Success *tg.Bool
}

var ChannelsCheckUsernameResult_Success_DEFAULT *tg.Bool

func (p *ChannelsCheckUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsCheckUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsCheckUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsCheckUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsCheckUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsCheckUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsCheckUsernameResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ChannelsCheckUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsCheckUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ChannelsCheckUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsCheckUsernameResult) GetResult() interface{} {
	return p.Success
}

func channelsUpdateUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsUpdateUsernameArgs)
	realResult := result.(*ChannelsUpdateUsernameResult)
	success, err := handler.(tg.RPCUsernames).ChannelsUpdateUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsUpdateUsernameArgs() interface{} {
	return &ChannelsUpdateUsernameArgs{}
}

func newChannelsUpdateUsernameResult() interface{} {
	return &ChannelsUpdateUsernameResult{}
}

type ChannelsUpdateUsernameArgs struct {
	Req *tg.TLChannelsUpdateUsername
}

func (p *ChannelsUpdateUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsUpdateUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsUpdateUsernameArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsUpdateUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsUpdateUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsUpdateUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsUpdateUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsUpdateUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsUpdateUsernameArgs_Req_DEFAULT *tg.TLChannelsUpdateUsername

func (p *ChannelsUpdateUsernameArgs) GetReq() *tg.TLChannelsUpdateUsername {
	if !p.IsSetReq() {
		return ChannelsUpdateUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsUpdateUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsUpdateUsernameResult struct {
	Success *tg.Bool
}

var ChannelsUpdateUsernameResult_Success_DEFAULT *tg.Bool

func (p *ChannelsUpdateUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsUpdateUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsUpdateUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsUpdateUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsUpdateUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsUpdateUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsUpdateUsernameResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ChannelsUpdateUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsUpdateUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ChannelsUpdateUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsUpdateUsernameResult) GetResult() interface{} {
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

func (p *kClient) AccountCheckUsername(ctx context.Context, req *tg.TLAccountCheckUsername) (r *tg.Bool, err error) {
	// var _args AccountCheckUsernameArgs
	// _args.Req = req
	// var _result AccountCheckUsernameResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCUsernames/account.checkUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountUpdateUsername(ctx context.Context, req *tg.TLAccountUpdateUsername) (r *tg.User, err error) {
	// var _args AccountUpdateUsernameArgs
	// _args.Req = req
	// var _result AccountUpdateUsernameResult

	_result := new(tg.User)
	if err = p.c.Call(ctx, "/tg.RPCUsernames/account.updateUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ContactsResolveUsername(ctx context.Context, req *tg.TLContactsResolveUsername) (r *tg.ContactsResolvedPeer, err error) {
	// var _args ContactsResolveUsernameArgs
	// _args.Req = req
	// var _result ContactsResolveUsernameResult

	_result := new(tg.ContactsResolvedPeer)
	if err = p.c.Call(ctx, "/tg.RPCUsernames/contacts.resolveUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChannelsCheckUsername(ctx context.Context, req *tg.TLChannelsCheckUsername) (r *tg.Bool, err error) {
	// var _args ChannelsCheckUsernameArgs
	// _args.Req = req
	// var _result ChannelsCheckUsernameResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCUsernames/channels.checkUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChannelsUpdateUsername(ctx context.Context, req *tg.TLChannelsUpdateUsername) (r *tg.Bool, err error) {
	// var _args ChannelsUpdateUsernameArgs
	// _args.Req = req
	// var _result ChannelsUpdateUsernameResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCUsernames/channels.updateUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
