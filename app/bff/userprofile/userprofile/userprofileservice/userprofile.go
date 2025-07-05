/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package userprofileservice

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
	"account.updateProfile": kitex.NewMethodInfo(
		accountUpdateProfileHandler,
		newAccountUpdateProfileArgs,
		newAccountUpdateProfileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.updateStatus": kitex.NewMethodInfo(
		accountUpdateStatusHandler,
		newAccountUpdateStatusArgs,
		newAccountUpdateStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.updateBirthday": kitex.NewMethodInfo(
		accountUpdateBirthdayHandler,
		newAccountUpdateBirthdayArgs,
		newAccountUpdateBirthdayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.updatePersonalChannel": kitex.NewMethodInfo(
		accountUpdatePersonalChannelHandler,
		newAccountUpdatePersonalChannelArgs,
		newAccountUpdatePersonalChannelResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.getBirthdays": kitex.NewMethodInfo(
		contactsGetBirthdaysHandler,
		newContactsGetBirthdaysArgs,
		newContactsGetBirthdaysResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"photos.updateProfilePhoto": kitex.NewMethodInfo(
		photosUpdateProfilePhotoHandler,
		newPhotosUpdateProfilePhotoArgs,
		newPhotosUpdateProfilePhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"photos.uploadProfilePhoto": kitex.NewMethodInfo(
		photosUploadProfilePhotoHandler,
		newPhotosUploadProfilePhotoArgs,
		newPhotosUploadProfilePhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"photos.deletePhotos": kitex.NewMethodInfo(
		photosDeletePhotosHandler,
		newPhotosDeletePhotosArgs,
		newPhotosDeletePhotosResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"photos.getUserPhotos": kitex.NewMethodInfo(
		photosGetUserPhotosHandler,
		newPhotosGetUserPhotosArgs,
		newPhotosGetUserPhotosResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"photos.uploadContactProfilePhoto": kitex.NewMethodInfo(
		photosUploadContactProfilePhotoHandler,
		newPhotosUploadContactProfilePhotoArgs,
		newPhotosUploadContactProfilePhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.updateVerified": kitex.NewMethodInfo(
		accountUpdateVerifiedHandler,
		newAccountUpdateVerifiedArgs,
		newAccountUpdateVerifiedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userprofileServiceServiceInfo                = NewServiceInfo()
	userprofileServiceServiceInfoForClient       = NewServiceInfoForClient()
	userprofileServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCUserProfile", userprofileServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCUserProfile", userprofileServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCUserProfile", userprofileServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userprofileServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userprofileServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userprofileServiceServiceInfoForClient
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
	serviceName := "RPCUserProfile"
	handlerType := (*tg.RPCUserProfile)(nil)
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
		"PackageName": "userprofile",
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

func accountUpdateProfileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUpdateProfileArgs)
	realResult := result.(*AccountUpdateProfileResult)
	success, err := handler.(tg.RPCUserProfile).AccountUpdateProfile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUpdateProfileArgs() interface{} {
	return &AccountUpdateProfileArgs{}
}

func newAccountUpdateProfileResult() interface{} {
	return &AccountUpdateProfileResult{}
}

type AccountUpdateProfileArgs struct {
	Req *tg.TLAccountUpdateProfile
}

func (p *AccountUpdateProfileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUpdateProfileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUpdateProfileArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUpdateProfile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUpdateProfileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUpdateProfileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUpdateProfileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUpdateProfile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUpdateProfileArgs_Req_DEFAULT *tg.TLAccountUpdateProfile

func (p *AccountUpdateProfileArgs) GetReq() *tg.TLAccountUpdateProfile {
	if !p.IsSetReq() {
		return AccountUpdateProfileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUpdateProfileArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUpdateProfileResult struct {
	Success *tg.User
}

var AccountUpdateProfileResult_Success_DEFAULT *tg.User

func (p *AccountUpdateProfileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUpdateProfileResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUpdateProfileResult) Unmarshal(in []byte) error {
	msg := new(tg.User)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateProfileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUpdateProfileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUpdateProfileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.User)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateProfileResult) GetSuccess() *tg.User {
	if !p.IsSetSuccess() {
		return AccountUpdateProfileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUpdateProfileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.User)
}

func (p *AccountUpdateProfileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUpdateProfileResult) GetResult() interface{} {
	return p.Success
}

func accountUpdateStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUpdateStatusArgs)
	realResult := result.(*AccountUpdateStatusResult)
	success, err := handler.(tg.RPCUserProfile).AccountUpdateStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUpdateStatusArgs() interface{} {
	return &AccountUpdateStatusArgs{}
}

func newAccountUpdateStatusResult() interface{} {
	return &AccountUpdateStatusResult{}
}

type AccountUpdateStatusArgs struct {
	Req *tg.TLAccountUpdateStatus
}

func (p *AccountUpdateStatusArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUpdateStatusArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUpdateStatusArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUpdateStatus)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUpdateStatusArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUpdateStatusArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUpdateStatusArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUpdateStatus)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUpdateStatusArgs_Req_DEFAULT *tg.TLAccountUpdateStatus

func (p *AccountUpdateStatusArgs) GetReq() *tg.TLAccountUpdateStatus {
	if !p.IsSetReq() {
		return AccountUpdateStatusArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUpdateStatusArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUpdateStatusResult struct {
	Success *tg.Bool
}

var AccountUpdateStatusResult_Success_DEFAULT *tg.Bool

func (p *AccountUpdateStatusResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUpdateStatusResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUpdateStatusResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateStatusResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUpdateStatusResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUpdateStatusResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateStatusResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountUpdateStatusResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUpdateStatusResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountUpdateStatusResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUpdateStatusResult) GetResult() interface{} {
	return p.Success
}

func accountUpdateBirthdayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUpdateBirthdayArgs)
	realResult := result.(*AccountUpdateBirthdayResult)
	success, err := handler.(tg.RPCUserProfile).AccountUpdateBirthday(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUpdateBirthdayArgs() interface{} {
	return &AccountUpdateBirthdayArgs{}
}

func newAccountUpdateBirthdayResult() interface{} {
	return &AccountUpdateBirthdayResult{}
}

type AccountUpdateBirthdayArgs struct {
	Req *tg.TLAccountUpdateBirthday
}

func (p *AccountUpdateBirthdayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUpdateBirthdayArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUpdateBirthdayArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUpdateBirthday)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUpdateBirthdayArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUpdateBirthdayArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUpdateBirthdayArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUpdateBirthday)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUpdateBirthdayArgs_Req_DEFAULT *tg.TLAccountUpdateBirthday

func (p *AccountUpdateBirthdayArgs) GetReq() *tg.TLAccountUpdateBirthday {
	if !p.IsSetReq() {
		return AccountUpdateBirthdayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUpdateBirthdayArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUpdateBirthdayResult struct {
	Success *tg.Bool
}

var AccountUpdateBirthdayResult_Success_DEFAULT *tg.Bool

func (p *AccountUpdateBirthdayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUpdateBirthdayResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUpdateBirthdayResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateBirthdayResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUpdateBirthdayResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUpdateBirthdayResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateBirthdayResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountUpdateBirthdayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUpdateBirthdayResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountUpdateBirthdayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUpdateBirthdayResult) GetResult() interface{} {
	return p.Success
}

func accountUpdatePersonalChannelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUpdatePersonalChannelArgs)
	realResult := result.(*AccountUpdatePersonalChannelResult)
	success, err := handler.(tg.RPCUserProfile).AccountUpdatePersonalChannel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUpdatePersonalChannelArgs() interface{} {
	return &AccountUpdatePersonalChannelArgs{}
}

func newAccountUpdatePersonalChannelResult() interface{} {
	return &AccountUpdatePersonalChannelResult{}
}

type AccountUpdatePersonalChannelArgs struct {
	Req *tg.TLAccountUpdatePersonalChannel
}

func (p *AccountUpdatePersonalChannelArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUpdatePersonalChannelArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUpdatePersonalChannelArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUpdatePersonalChannel)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUpdatePersonalChannelArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUpdatePersonalChannelArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUpdatePersonalChannelArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUpdatePersonalChannel)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUpdatePersonalChannelArgs_Req_DEFAULT *tg.TLAccountUpdatePersonalChannel

func (p *AccountUpdatePersonalChannelArgs) GetReq() *tg.TLAccountUpdatePersonalChannel {
	if !p.IsSetReq() {
		return AccountUpdatePersonalChannelArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUpdatePersonalChannelArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUpdatePersonalChannelResult struct {
	Success *tg.Bool
}

var AccountUpdatePersonalChannelResult_Success_DEFAULT *tg.Bool

func (p *AccountUpdatePersonalChannelResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUpdatePersonalChannelResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUpdatePersonalChannelResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdatePersonalChannelResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUpdatePersonalChannelResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUpdatePersonalChannelResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdatePersonalChannelResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountUpdatePersonalChannelResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUpdatePersonalChannelResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountUpdatePersonalChannelResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUpdatePersonalChannelResult) GetResult() interface{} {
	return p.Success
}

func contactsGetBirthdaysHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetBirthdaysArgs)
	realResult := result.(*ContactsGetBirthdaysResult)
	success, err := handler.(tg.RPCUserProfile).ContactsGetBirthdays(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetBirthdaysArgs() interface{} {
	return &ContactsGetBirthdaysArgs{}
}

func newContactsGetBirthdaysResult() interface{} {
	return &ContactsGetBirthdaysResult{}
}

type ContactsGetBirthdaysArgs struct {
	Req *tg.TLContactsGetBirthdays
}

func (p *ContactsGetBirthdaysArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetBirthdaysArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetBirthdaysArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetBirthdays)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetBirthdaysArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetBirthdaysArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetBirthdaysArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetBirthdays)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetBirthdaysArgs_Req_DEFAULT *tg.TLContactsGetBirthdays

func (p *ContactsGetBirthdaysArgs) GetReq() *tg.TLContactsGetBirthdays {
	if !p.IsSetReq() {
		return ContactsGetBirthdaysArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetBirthdaysArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetBirthdaysResult struct {
	Success *tg.ContactsContactBirthdays
}

var ContactsGetBirthdaysResult_Success_DEFAULT *tg.ContactsContactBirthdays

func (p *ContactsGetBirthdaysResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetBirthdaysResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetBirthdaysResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsContactBirthdays)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetBirthdaysResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetBirthdaysResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetBirthdaysResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsContactBirthdays)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetBirthdaysResult) GetSuccess() *tg.ContactsContactBirthdays {
	if !p.IsSetSuccess() {
		return ContactsGetBirthdaysResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetBirthdaysResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsContactBirthdays)
}

func (p *ContactsGetBirthdaysResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetBirthdaysResult) GetResult() interface{} {
	return p.Success
}

func photosUpdateProfilePhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PhotosUpdateProfilePhotoArgs)
	realResult := result.(*PhotosUpdateProfilePhotoResult)
	success, err := handler.(tg.RPCUserProfile).PhotosUpdateProfilePhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPhotosUpdateProfilePhotoArgs() interface{} {
	return &PhotosUpdateProfilePhotoArgs{}
}

func newPhotosUpdateProfilePhotoResult() interface{} {
	return &PhotosUpdateProfilePhotoResult{}
}

type PhotosUpdateProfilePhotoArgs struct {
	Req *tg.TLPhotosUpdateProfilePhoto
}

func (p *PhotosUpdateProfilePhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PhotosUpdateProfilePhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PhotosUpdateProfilePhotoArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLPhotosUpdateProfilePhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PhotosUpdateProfilePhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PhotosUpdateProfilePhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PhotosUpdateProfilePhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLPhotosUpdateProfilePhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PhotosUpdateProfilePhotoArgs_Req_DEFAULT *tg.TLPhotosUpdateProfilePhoto

func (p *PhotosUpdateProfilePhotoArgs) GetReq() *tg.TLPhotosUpdateProfilePhoto {
	if !p.IsSetReq() {
		return PhotosUpdateProfilePhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PhotosUpdateProfilePhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type PhotosUpdateProfilePhotoResult struct {
	Success *tg.PhotosPhoto
}

var PhotosUpdateProfilePhotoResult_Success_DEFAULT *tg.PhotosPhoto

func (p *PhotosUpdateProfilePhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PhotosUpdateProfilePhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *PhotosUpdateProfilePhotoResult) Unmarshal(in []byte) error {
	msg := new(tg.PhotosPhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosUpdateProfilePhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PhotosUpdateProfilePhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PhotosUpdateProfilePhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.PhotosPhoto)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosUpdateProfilePhotoResult) GetSuccess() *tg.PhotosPhoto {
	if !p.IsSetSuccess() {
		return PhotosUpdateProfilePhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PhotosUpdateProfilePhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.PhotosPhoto)
}

func (p *PhotosUpdateProfilePhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PhotosUpdateProfilePhotoResult) GetResult() interface{} {
	return p.Success
}

func photosUploadProfilePhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PhotosUploadProfilePhotoArgs)
	realResult := result.(*PhotosUploadProfilePhotoResult)
	success, err := handler.(tg.RPCUserProfile).PhotosUploadProfilePhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPhotosUploadProfilePhotoArgs() interface{} {
	return &PhotosUploadProfilePhotoArgs{}
}

func newPhotosUploadProfilePhotoResult() interface{} {
	return &PhotosUploadProfilePhotoResult{}
}

type PhotosUploadProfilePhotoArgs struct {
	Req *tg.TLPhotosUploadProfilePhoto
}

func (p *PhotosUploadProfilePhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PhotosUploadProfilePhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PhotosUploadProfilePhotoArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLPhotosUploadProfilePhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PhotosUploadProfilePhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PhotosUploadProfilePhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PhotosUploadProfilePhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLPhotosUploadProfilePhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PhotosUploadProfilePhotoArgs_Req_DEFAULT *tg.TLPhotosUploadProfilePhoto

func (p *PhotosUploadProfilePhotoArgs) GetReq() *tg.TLPhotosUploadProfilePhoto {
	if !p.IsSetReq() {
		return PhotosUploadProfilePhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PhotosUploadProfilePhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type PhotosUploadProfilePhotoResult struct {
	Success *tg.PhotosPhoto
}

var PhotosUploadProfilePhotoResult_Success_DEFAULT *tg.PhotosPhoto

func (p *PhotosUploadProfilePhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PhotosUploadProfilePhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *PhotosUploadProfilePhotoResult) Unmarshal(in []byte) error {
	msg := new(tg.PhotosPhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosUploadProfilePhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PhotosUploadProfilePhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PhotosUploadProfilePhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.PhotosPhoto)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosUploadProfilePhotoResult) GetSuccess() *tg.PhotosPhoto {
	if !p.IsSetSuccess() {
		return PhotosUploadProfilePhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PhotosUploadProfilePhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.PhotosPhoto)
}

func (p *PhotosUploadProfilePhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PhotosUploadProfilePhotoResult) GetResult() interface{} {
	return p.Success
}

func photosDeletePhotosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PhotosDeletePhotosArgs)
	realResult := result.(*PhotosDeletePhotosResult)
	success, err := handler.(tg.RPCUserProfile).PhotosDeletePhotos(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPhotosDeletePhotosArgs() interface{} {
	return &PhotosDeletePhotosArgs{}
}

func newPhotosDeletePhotosResult() interface{} {
	return &PhotosDeletePhotosResult{}
}

type PhotosDeletePhotosArgs struct {
	Req *tg.TLPhotosDeletePhotos
}

func (p *PhotosDeletePhotosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PhotosDeletePhotosArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PhotosDeletePhotosArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLPhotosDeletePhotos)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PhotosDeletePhotosArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PhotosDeletePhotosArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PhotosDeletePhotosArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLPhotosDeletePhotos)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PhotosDeletePhotosArgs_Req_DEFAULT *tg.TLPhotosDeletePhotos

func (p *PhotosDeletePhotosArgs) GetReq() *tg.TLPhotosDeletePhotos {
	if !p.IsSetReq() {
		return PhotosDeletePhotosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PhotosDeletePhotosArgs) IsSetReq() bool {
	return p.Req != nil
}

type PhotosDeletePhotosResult struct {
	Success *tg.VectorLong
}

var PhotosDeletePhotosResult_Success_DEFAULT *tg.VectorLong

func (p *PhotosDeletePhotosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PhotosDeletePhotosResult")
	}
	return json.Marshal(p.Success)
}

func (p *PhotosDeletePhotosResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosDeletePhotosResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PhotosDeletePhotosResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PhotosDeletePhotosResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosDeletePhotosResult) GetSuccess() *tg.VectorLong {
	if !p.IsSetSuccess() {
		return PhotosDeletePhotosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PhotosDeletePhotosResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorLong)
}

func (p *PhotosDeletePhotosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PhotosDeletePhotosResult) GetResult() interface{} {
	return p.Success
}

func photosGetUserPhotosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PhotosGetUserPhotosArgs)
	realResult := result.(*PhotosGetUserPhotosResult)
	success, err := handler.(tg.RPCUserProfile).PhotosGetUserPhotos(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPhotosGetUserPhotosArgs() interface{} {
	return &PhotosGetUserPhotosArgs{}
}

func newPhotosGetUserPhotosResult() interface{} {
	return &PhotosGetUserPhotosResult{}
}

type PhotosGetUserPhotosArgs struct {
	Req *tg.TLPhotosGetUserPhotos
}

func (p *PhotosGetUserPhotosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PhotosGetUserPhotosArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PhotosGetUserPhotosArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLPhotosGetUserPhotos)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PhotosGetUserPhotosArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PhotosGetUserPhotosArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PhotosGetUserPhotosArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLPhotosGetUserPhotos)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PhotosGetUserPhotosArgs_Req_DEFAULT *tg.TLPhotosGetUserPhotos

func (p *PhotosGetUserPhotosArgs) GetReq() *tg.TLPhotosGetUserPhotos {
	if !p.IsSetReq() {
		return PhotosGetUserPhotosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PhotosGetUserPhotosArgs) IsSetReq() bool {
	return p.Req != nil
}

type PhotosGetUserPhotosResult struct {
	Success *tg.PhotosPhotos
}

var PhotosGetUserPhotosResult_Success_DEFAULT *tg.PhotosPhotos

func (p *PhotosGetUserPhotosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PhotosGetUserPhotosResult")
	}
	return json.Marshal(p.Success)
}

func (p *PhotosGetUserPhotosResult) Unmarshal(in []byte) error {
	msg := new(tg.PhotosPhotos)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosGetUserPhotosResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PhotosGetUserPhotosResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PhotosGetUserPhotosResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.PhotosPhotos)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosGetUserPhotosResult) GetSuccess() *tg.PhotosPhotos {
	if !p.IsSetSuccess() {
		return PhotosGetUserPhotosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PhotosGetUserPhotosResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.PhotosPhotos)
}

func (p *PhotosGetUserPhotosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PhotosGetUserPhotosResult) GetResult() interface{} {
	return p.Success
}

func photosUploadContactProfilePhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PhotosUploadContactProfilePhotoArgs)
	realResult := result.(*PhotosUploadContactProfilePhotoResult)
	success, err := handler.(tg.RPCUserProfile).PhotosUploadContactProfilePhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPhotosUploadContactProfilePhotoArgs() interface{} {
	return &PhotosUploadContactProfilePhotoArgs{}
}

func newPhotosUploadContactProfilePhotoResult() interface{} {
	return &PhotosUploadContactProfilePhotoResult{}
}

type PhotosUploadContactProfilePhotoArgs struct {
	Req *tg.TLPhotosUploadContactProfilePhoto
}

func (p *PhotosUploadContactProfilePhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PhotosUploadContactProfilePhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PhotosUploadContactProfilePhotoArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLPhotosUploadContactProfilePhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PhotosUploadContactProfilePhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PhotosUploadContactProfilePhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PhotosUploadContactProfilePhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLPhotosUploadContactProfilePhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PhotosUploadContactProfilePhotoArgs_Req_DEFAULT *tg.TLPhotosUploadContactProfilePhoto

func (p *PhotosUploadContactProfilePhotoArgs) GetReq() *tg.TLPhotosUploadContactProfilePhoto {
	if !p.IsSetReq() {
		return PhotosUploadContactProfilePhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PhotosUploadContactProfilePhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type PhotosUploadContactProfilePhotoResult struct {
	Success *tg.PhotosPhoto
}

var PhotosUploadContactProfilePhotoResult_Success_DEFAULT *tg.PhotosPhoto

func (p *PhotosUploadContactProfilePhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PhotosUploadContactProfilePhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *PhotosUploadContactProfilePhotoResult) Unmarshal(in []byte) error {
	msg := new(tg.PhotosPhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosUploadContactProfilePhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PhotosUploadContactProfilePhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PhotosUploadContactProfilePhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.PhotosPhoto)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PhotosUploadContactProfilePhotoResult) GetSuccess() *tg.PhotosPhoto {
	if !p.IsSetSuccess() {
		return PhotosUploadContactProfilePhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PhotosUploadContactProfilePhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.PhotosPhoto)
}

func (p *PhotosUploadContactProfilePhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PhotosUploadContactProfilePhotoResult) GetResult() interface{} {
	return p.Success
}

func accountUpdateVerifiedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountUpdateVerifiedArgs)
	realResult := result.(*AccountUpdateVerifiedResult)
	success, err := handler.(tg.RPCUserProfile).AccountUpdateVerified(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountUpdateVerifiedArgs() interface{} {
	return &AccountUpdateVerifiedArgs{}
}

func newAccountUpdateVerifiedResult() interface{} {
	return &AccountUpdateVerifiedResult{}
}

type AccountUpdateVerifiedArgs struct {
	Req *tg.TLAccountUpdateVerified
}

func (p *AccountUpdateVerifiedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountUpdateVerifiedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountUpdateVerifiedArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountUpdateVerified)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountUpdateVerifiedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountUpdateVerifiedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountUpdateVerifiedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountUpdateVerified)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountUpdateVerifiedArgs_Req_DEFAULT *tg.TLAccountUpdateVerified

func (p *AccountUpdateVerifiedArgs) GetReq() *tg.TLAccountUpdateVerified {
	if !p.IsSetReq() {
		return AccountUpdateVerifiedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountUpdateVerifiedArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountUpdateVerifiedResult struct {
	Success *tg.User
}

var AccountUpdateVerifiedResult_Success_DEFAULT *tg.User

func (p *AccountUpdateVerifiedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountUpdateVerifiedResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountUpdateVerifiedResult) Unmarshal(in []byte) error {
	msg := new(tg.User)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateVerifiedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountUpdateVerifiedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountUpdateVerifiedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.User)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountUpdateVerifiedResult) GetSuccess() *tg.User {
	if !p.IsSetSuccess() {
		return AccountUpdateVerifiedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountUpdateVerifiedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.User)
}

func (p *AccountUpdateVerifiedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountUpdateVerifiedResult) GetResult() interface{} {
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

func (p *kClient) AccountUpdateProfile(ctx context.Context, req *tg.TLAccountUpdateProfile) (r *tg.User, err error) {
	// var _args AccountUpdateProfileArgs
	// _args.Req = req
	// var _result AccountUpdateProfileResult

	_result := new(tg.User)
	if err = p.c.Call(ctx, "account.updateProfile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountUpdateStatus(ctx context.Context, req *tg.TLAccountUpdateStatus) (r *tg.Bool, err error) {
	// var _args AccountUpdateStatusArgs
	// _args.Req = req
	// var _result AccountUpdateStatusResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.updateStatus", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountUpdateBirthday(ctx context.Context, req *tg.TLAccountUpdateBirthday) (r *tg.Bool, err error) {
	// var _args AccountUpdateBirthdayArgs
	// _args.Req = req
	// var _result AccountUpdateBirthdayResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.updateBirthday", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountUpdatePersonalChannel(ctx context.Context, req *tg.TLAccountUpdatePersonalChannel) (r *tg.Bool, err error) {
	// var _args AccountUpdatePersonalChannelArgs
	// _args.Req = req
	// var _result AccountUpdatePersonalChannelResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "account.updatePersonalChannel", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ContactsGetBirthdays(ctx context.Context, req *tg.TLContactsGetBirthdays) (r *tg.ContactsContactBirthdays, err error) {
	// var _args ContactsGetBirthdaysArgs
	// _args.Req = req
	// var _result ContactsGetBirthdaysResult

	_result := new(tg.ContactsContactBirthdays)
	if err = p.c.Call(ctx, "contacts.getBirthdays", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) PhotosUpdateProfilePhoto(ctx context.Context, req *tg.TLPhotosUpdateProfilePhoto) (r *tg.PhotosPhoto, err error) {
	// var _args PhotosUpdateProfilePhotoArgs
	// _args.Req = req
	// var _result PhotosUpdateProfilePhotoResult

	_result := new(tg.PhotosPhoto)
	if err = p.c.Call(ctx, "photos.updateProfilePhoto", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) PhotosUploadProfilePhoto(ctx context.Context, req *tg.TLPhotosUploadProfilePhoto) (r *tg.PhotosPhoto, err error) {
	// var _args PhotosUploadProfilePhotoArgs
	// _args.Req = req
	// var _result PhotosUploadProfilePhotoResult

	_result := new(tg.PhotosPhoto)
	if err = p.c.Call(ctx, "photos.uploadProfilePhoto", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) PhotosDeletePhotos(ctx context.Context, req *tg.TLPhotosDeletePhotos) (r *tg.VectorLong, err error) {
	// var _args PhotosDeletePhotosArgs
	// _args.Req = req
	// var _result PhotosDeletePhotosResult

	_result := new(tg.VectorLong)
	if err = p.c.Call(ctx, "photos.deletePhotos", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) PhotosGetUserPhotos(ctx context.Context, req *tg.TLPhotosGetUserPhotos) (r *tg.PhotosPhotos, err error) {
	// var _args PhotosGetUserPhotosArgs
	// _args.Req = req
	// var _result PhotosGetUserPhotosResult

	_result := new(tg.PhotosPhotos)
	if err = p.c.Call(ctx, "photos.getUserPhotos", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) PhotosUploadContactProfilePhoto(ctx context.Context, req *tg.TLPhotosUploadContactProfilePhoto) (r *tg.PhotosPhoto, err error) {
	// var _args PhotosUploadContactProfilePhotoArgs
	// _args.Req = req
	// var _result PhotosUploadContactProfilePhotoResult

	_result := new(tg.PhotosPhoto)
	if err = p.c.Call(ctx, "photos.uploadContactProfilePhoto", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) AccountUpdateVerified(ctx context.Context, req *tg.TLAccountUpdateVerified) (r *tg.User, err error) {
	// var _args AccountUpdateVerifiedArgs
	// _args.Req = req
	// var _result AccountUpdateVerifiedResult

	_result := new(tg.User)
	if err = p.c.Call(ctx, "account.updateVerified", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
