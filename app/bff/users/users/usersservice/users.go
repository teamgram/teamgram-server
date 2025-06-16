/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package usersservice

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
	"users.getUsers": kitex.NewMethodInfo(
		usersGetUsersHandler,
		newUsersGetUsersArgs,
		newUsersGetUsersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"users.getFullUser": kitex.NewMethodInfo(
		usersGetFullUserHandler,
		newUsersGetFullUserArgs,
		newUsersGetFullUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.resolvePhone": kitex.NewMethodInfo(
		contactsResolvePhoneHandler,
		newContactsResolvePhoneArgs,
		newContactsResolvePhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"users.getMe": kitex.NewMethodInfo(
		usersGetMeHandler,
		newUsersGetMeArgs,
		newUsersGetMeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	usersServiceServiceInfo                = NewServiceInfo()
	usersServiceServiceInfoForClient       = NewServiceInfoForClient()
	usersServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return usersServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return usersServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return usersServiceServiceInfoForClient
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
	serviceName := "RPCUsers"
	handlerType := (*tg.RPCUsers)(nil)
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
		"PackageName": "users",
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

func usersGetUsersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UsersGetUsersArgs)
	realResult := result.(*UsersGetUsersResult)
	success, err := handler.(tg.RPCUsers).UsersGetUsers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUsersGetUsersArgs() interface{} {
	return &UsersGetUsersArgs{}
}

func newUsersGetUsersResult() interface{} {
	return &UsersGetUsersResult{}
}

type UsersGetUsersArgs struct {
	Req *tg.TLUsersGetUsers
}

func (p *UsersGetUsersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UsersGetUsersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UsersGetUsersArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUsersGetUsers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UsersGetUsersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UsersGetUsersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UsersGetUsersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUsersGetUsers)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UsersGetUsersArgs_Req_DEFAULT *tg.TLUsersGetUsers

func (p *UsersGetUsersArgs) GetReq() *tg.TLUsersGetUsers {
	if !p.IsSetReq() {
		return UsersGetUsersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UsersGetUsersArgs) IsSetReq() bool {
	return p.Req != nil
}

type UsersGetUsersResult struct {
	Success *tg.VectorUser
}

var UsersGetUsersResult_Success_DEFAULT *tg.VectorUser

func (p *UsersGetUsersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UsersGetUsersResult")
	}
	return json.Marshal(p.Success)
}

func (p *UsersGetUsersResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersGetUsersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UsersGetUsersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UsersGetUsersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersGetUsersResult) GetSuccess() *tg.VectorUser {
	if !p.IsSetSuccess() {
		return UsersGetUsersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UsersGetUsersResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorUser)
}

func (p *UsersGetUsersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UsersGetUsersResult) GetResult() interface{} {
	return p.Success
}

func usersGetFullUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UsersGetFullUserArgs)
	realResult := result.(*UsersGetFullUserResult)
	success, err := handler.(tg.RPCUsers).UsersGetFullUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUsersGetFullUserArgs() interface{} {
	return &UsersGetFullUserArgs{}
}

func newUsersGetFullUserResult() interface{} {
	return &UsersGetFullUserResult{}
}

type UsersGetFullUserArgs struct {
	Req *tg.TLUsersGetFullUser
}

func (p *UsersGetFullUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UsersGetFullUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UsersGetFullUserArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUsersGetFullUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UsersGetFullUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UsersGetFullUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UsersGetFullUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUsersGetFullUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UsersGetFullUserArgs_Req_DEFAULT *tg.TLUsersGetFullUser

func (p *UsersGetFullUserArgs) GetReq() *tg.TLUsersGetFullUser {
	if !p.IsSetReq() {
		return UsersGetFullUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UsersGetFullUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type UsersGetFullUserResult struct {
	Success *tg.UsersUserFull
}

var UsersGetFullUserResult_Success_DEFAULT *tg.UsersUserFull

func (p *UsersGetFullUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UsersGetFullUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *UsersGetFullUserResult) Unmarshal(in []byte) error {
	msg := new(tg.UsersUserFull)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersGetFullUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UsersGetFullUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UsersGetFullUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UsersUserFull)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersGetFullUserResult) GetSuccess() *tg.UsersUserFull {
	if !p.IsSetSuccess() {
		return UsersGetFullUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UsersGetFullUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UsersUserFull)
}

func (p *UsersGetFullUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UsersGetFullUserResult) GetResult() interface{} {
	return p.Success
}

func contactsResolvePhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsResolvePhoneArgs)
	realResult := result.(*ContactsResolvePhoneResult)
	success, err := handler.(tg.RPCUsers).ContactsResolvePhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsResolvePhoneArgs() interface{} {
	return &ContactsResolvePhoneArgs{}
}

func newContactsResolvePhoneResult() interface{} {
	return &ContactsResolvePhoneResult{}
}

type ContactsResolvePhoneArgs struct {
	Req *tg.TLContactsResolvePhone
}

func (p *ContactsResolvePhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsResolvePhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsResolvePhoneArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsResolvePhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsResolvePhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsResolvePhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsResolvePhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsResolvePhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsResolvePhoneArgs_Req_DEFAULT *tg.TLContactsResolvePhone

func (p *ContactsResolvePhoneArgs) GetReq() *tg.TLContactsResolvePhone {
	if !p.IsSetReq() {
		return ContactsResolvePhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsResolvePhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsResolvePhoneResult struct {
	Success *tg.ContactsResolvedPeer
}

var ContactsResolvePhoneResult_Success_DEFAULT *tg.ContactsResolvedPeer

func (p *ContactsResolvePhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsResolvePhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsResolvePhoneResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsResolvedPeer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsResolvePhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsResolvePhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsResolvePhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsResolvedPeer)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsResolvePhoneResult) GetSuccess() *tg.ContactsResolvedPeer {
	if !p.IsSetSuccess() {
		return ContactsResolvePhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsResolvePhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsResolvedPeer)
}

func (p *ContactsResolvePhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsResolvePhoneResult) GetResult() interface{} {
	return p.Success
}

func usersGetMeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UsersGetMeArgs)
	realResult := result.(*UsersGetMeResult)
	success, err := handler.(tg.RPCUsers).UsersGetMe(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUsersGetMeArgs() interface{} {
	return &UsersGetMeArgs{}
}

func newUsersGetMeResult() interface{} {
	return &UsersGetMeResult{}
}

type UsersGetMeArgs struct {
	Req *tg.TLUsersGetMe
}

func (p *UsersGetMeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UsersGetMeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UsersGetMeArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUsersGetMe)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UsersGetMeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UsersGetMeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UsersGetMeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUsersGetMe)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UsersGetMeArgs_Req_DEFAULT *tg.TLUsersGetMe

func (p *UsersGetMeArgs) GetReq() *tg.TLUsersGetMe {
	if !p.IsSetReq() {
		return UsersGetMeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UsersGetMeArgs) IsSetReq() bool {
	return p.Req != nil
}

type UsersGetMeResult struct {
	Success *tg.User
}

var UsersGetMeResult_Success_DEFAULT *tg.User

func (p *UsersGetMeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UsersGetMeResult")
	}
	return json.Marshal(p.Success)
}

func (p *UsersGetMeResult) Unmarshal(in []byte) error {
	msg := new(tg.User)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersGetMeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UsersGetMeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UsersGetMeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.User)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UsersGetMeResult) GetSuccess() *tg.User {
	if !p.IsSetSuccess() {
		return UsersGetMeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UsersGetMeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.User)
}

func (p *UsersGetMeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UsersGetMeResult) GetResult() interface{} {
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

func (p *kClient) UsersGetUsers(ctx context.Context, req *tg.TLUsersGetUsers) (r *tg.VectorUser, err error) {
	var _args UsersGetUsersArgs
	_args.Req = req
	var _result UsersGetUsersResult
	if err = p.c.Call(ctx, "users.getUsers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UsersGetFullUser(ctx context.Context, req *tg.TLUsersGetFullUser) (r *tg.UsersUserFull, err error) {
	var _args UsersGetFullUserArgs
	_args.Req = req
	var _result UsersGetFullUserResult
	if err = p.c.Call(ctx, "users.getFullUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsResolvePhone(ctx context.Context, req *tg.TLContactsResolvePhone) (r *tg.ContactsResolvedPeer, err error) {
	var _args ContactsResolvePhoneArgs
	_args.Req = req
	var _result ContactsResolvePhoneResult
	if err = p.c.Call(ctx, "contacts.resolvePhone", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UsersGetMe(ctx context.Context, req *tg.TLUsersGetMe) (r *tg.User, err error) {
	var _args UsersGetMeArgs
	_args.Req = req
	var _result UsersGetMeResult
	if err = p.c.Call(ctx, "users.getMe", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
