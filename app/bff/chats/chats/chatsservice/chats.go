/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chatsservice

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
	"messages.getChats": kitex.NewMethodInfo(
		messagesGetChatsHandler,
		newMessagesGetChatsArgs,
		newMessagesGetChatsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getFullChat": kitex.NewMethodInfo(
		messagesGetFullChatHandler,
		newMessagesGetFullChatArgs,
		newMessagesGetFullChatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.editChatTitle": kitex.NewMethodInfo(
		messagesEditChatTitleHandler,
		newMessagesEditChatTitleArgs,
		newMessagesEditChatTitleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.editChatPhoto": kitex.NewMethodInfo(
		messagesEditChatPhotoHandler,
		newMessagesEditChatPhotoArgs,
		newMessagesEditChatPhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.addChatUser": kitex.NewMethodInfo(
		messagesAddChatUserHandler,
		newMessagesAddChatUserArgs,
		newMessagesAddChatUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.deleteChatUser": kitex.NewMethodInfo(
		messagesDeleteChatUserHandler,
		newMessagesDeleteChatUserArgs,
		newMessagesDeleteChatUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.createChat": kitex.NewMethodInfo(
		messagesCreateChatHandler,
		newMessagesCreateChatArgs,
		newMessagesCreateChatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.editChatAdmin": kitex.NewMethodInfo(
		messagesEditChatAdminHandler,
		newMessagesEditChatAdminArgs,
		newMessagesEditChatAdminResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.migrateChat": kitex.NewMethodInfo(
		messagesMigrateChatHandler,
		newMessagesMigrateChatArgs,
		newMessagesMigrateChatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getCommonChats": kitex.NewMethodInfo(
		messagesGetCommonChatsHandler,
		newMessagesGetCommonChatsArgs,
		newMessagesGetCommonChatsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.editChatAbout": kitex.NewMethodInfo(
		messagesEditChatAboutHandler,
		newMessagesEditChatAboutArgs,
		newMessagesEditChatAboutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.editChatDefaultBannedRights": kitex.NewMethodInfo(
		messagesEditChatDefaultBannedRightsHandler,
		newMessagesEditChatDefaultBannedRightsArgs,
		newMessagesEditChatDefaultBannedRightsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.deleteChat": kitex.NewMethodInfo(
		messagesDeleteChatHandler,
		newMessagesDeleteChatArgs,
		newMessagesDeleteChatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"messages.getMessageReadParticipants": kitex.NewMethodInfo(
		messagesGetMessageReadParticipantsHandler,
		newMessagesGetMessageReadParticipantsArgs,
		newMessagesGetMessageReadParticipantsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.convertToGigagroup": kitex.NewMethodInfo(
		channelsConvertToGigagroupHandler,
		newChannelsConvertToGigagroupArgs,
		newChannelsConvertToGigagroupResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"channels.setEmojiStickers": kitex.NewMethodInfo(
		channelsSetEmojiStickersHandler,
		newChannelsSetEmojiStickersArgs,
		newChannelsSetEmojiStickersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	chatsServiceServiceInfo                = NewServiceInfo()
	chatsServiceServiceInfoForClient       = NewServiceInfoForClient()
	chatsServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return chatsServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return chatsServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return chatsServiceServiceInfoForClient
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
	serviceName := "RPCChats"
	handlerType := (*tg.RPCChats)(nil)
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
		"PackageName": "chats",
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

func messagesGetChatsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetChatsArgs)
	realResult := result.(*MessagesGetChatsResult)
	success, err := handler.(tg.RPCChats).MessagesGetChats(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetChatsArgs() interface{} {
	return &MessagesGetChatsArgs{}
}

func newMessagesGetChatsResult() interface{} {
	return &MessagesGetChatsResult{}
}

type MessagesGetChatsArgs struct {
	Req *tg.TLMessagesGetChats
}

func (p *MessagesGetChatsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetChatsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetChatsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetChats)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetChatsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetChatsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetChatsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetChats)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetChatsArgs_Req_DEFAULT *tg.TLMessagesGetChats

func (p *MessagesGetChatsArgs) GetReq() *tg.TLMessagesGetChats {
	if !p.IsSetReq() {
		return MessagesGetChatsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetChatsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetChatsResult struct {
	Success *tg.MessagesChats
}

var MessagesGetChatsResult_Success_DEFAULT *tg.MessagesChats

func (p *MessagesGetChatsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetChatsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetChatsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesChats)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetChatsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetChatsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetChatsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesChats)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetChatsResult) GetSuccess() *tg.MessagesChats {
	if !p.IsSetSuccess() {
		return MessagesGetChatsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetChatsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesChats)
}

func (p *MessagesGetChatsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetChatsResult) GetResult() interface{} {
	return p.Success
}

func messagesGetFullChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetFullChatArgs)
	realResult := result.(*MessagesGetFullChatResult)
	success, err := handler.(tg.RPCChats).MessagesGetFullChat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetFullChatArgs() interface{} {
	return &MessagesGetFullChatArgs{}
}

func newMessagesGetFullChatResult() interface{} {
	return &MessagesGetFullChatResult{}
}

type MessagesGetFullChatArgs struct {
	Req *tg.TLMessagesGetFullChat
}

func (p *MessagesGetFullChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetFullChatArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetFullChatArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetFullChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetFullChatArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetFullChatArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetFullChatArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetFullChat)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetFullChatArgs_Req_DEFAULT *tg.TLMessagesGetFullChat

func (p *MessagesGetFullChatArgs) GetReq() *tg.TLMessagesGetFullChat {
	if !p.IsSetReq() {
		return MessagesGetFullChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetFullChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetFullChatResult struct {
	Success *tg.MessagesChatFull
}

var MessagesGetFullChatResult_Success_DEFAULT *tg.MessagesChatFull

func (p *MessagesGetFullChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetFullChatResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetFullChatResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesChatFull)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetFullChatResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetFullChatResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetFullChatResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesChatFull)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetFullChatResult) GetSuccess() *tg.MessagesChatFull {
	if !p.IsSetSuccess() {
		return MessagesGetFullChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetFullChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesChatFull)
}

func (p *MessagesGetFullChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetFullChatResult) GetResult() interface{} {
	return p.Success
}

func messagesEditChatTitleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesEditChatTitleArgs)
	realResult := result.(*MessagesEditChatTitleResult)
	success, err := handler.(tg.RPCChats).MessagesEditChatTitle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesEditChatTitleArgs() interface{} {
	return &MessagesEditChatTitleArgs{}
}

func newMessagesEditChatTitleResult() interface{} {
	return &MessagesEditChatTitleResult{}
}

type MessagesEditChatTitleArgs struct {
	Req *tg.TLMessagesEditChatTitle
}

func (p *MessagesEditChatTitleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesEditChatTitleArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesEditChatTitleArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesEditChatTitle)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesEditChatTitleArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesEditChatTitleArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesEditChatTitleArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesEditChatTitle)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesEditChatTitleArgs_Req_DEFAULT *tg.TLMessagesEditChatTitle

func (p *MessagesEditChatTitleArgs) GetReq() *tg.TLMessagesEditChatTitle {
	if !p.IsSetReq() {
		return MessagesEditChatTitleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesEditChatTitleArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesEditChatTitleResult struct {
	Success *tg.Updates
}

var MessagesEditChatTitleResult_Success_DEFAULT *tg.Updates

func (p *MessagesEditChatTitleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesEditChatTitleResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesEditChatTitleResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatTitleResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesEditChatTitleResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesEditChatTitleResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatTitleResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesEditChatTitleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesEditChatTitleResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesEditChatTitleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesEditChatTitleResult) GetResult() interface{} {
	return p.Success
}

func messagesEditChatPhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesEditChatPhotoArgs)
	realResult := result.(*MessagesEditChatPhotoResult)
	success, err := handler.(tg.RPCChats).MessagesEditChatPhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesEditChatPhotoArgs() interface{} {
	return &MessagesEditChatPhotoArgs{}
}

func newMessagesEditChatPhotoResult() interface{} {
	return &MessagesEditChatPhotoResult{}
}

type MessagesEditChatPhotoArgs struct {
	Req *tg.TLMessagesEditChatPhoto
}

func (p *MessagesEditChatPhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesEditChatPhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesEditChatPhotoArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesEditChatPhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesEditChatPhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesEditChatPhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesEditChatPhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesEditChatPhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesEditChatPhotoArgs_Req_DEFAULT *tg.TLMessagesEditChatPhoto

func (p *MessagesEditChatPhotoArgs) GetReq() *tg.TLMessagesEditChatPhoto {
	if !p.IsSetReq() {
		return MessagesEditChatPhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesEditChatPhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesEditChatPhotoResult struct {
	Success *tg.Updates
}

var MessagesEditChatPhotoResult_Success_DEFAULT *tg.Updates

func (p *MessagesEditChatPhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesEditChatPhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesEditChatPhotoResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatPhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesEditChatPhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesEditChatPhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatPhotoResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesEditChatPhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesEditChatPhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesEditChatPhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesEditChatPhotoResult) GetResult() interface{} {
	return p.Success
}

func messagesAddChatUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesAddChatUserArgs)
	realResult := result.(*MessagesAddChatUserResult)
	success, err := handler.(tg.RPCChats).MessagesAddChatUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesAddChatUserArgs() interface{} {
	return &MessagesAddChatUserArgs{}
}

func newMessagesAddChatUserResult() interface{} {
	return &MessagesAddChatUserResult{}
}

type MessagesAddChatUserArgs struct {
	Req *tg.TLMessagesAddChatUser
}

func (p *MessagesAddChatUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesAddChatUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesAddChatUserArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesAddChatUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesAddChatUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesAddChatUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesAddChatUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesAddChatUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesAddChatUserArgs_Req_DEFAULT *tg.TLMessagesAddChatUser

func (p *MessagesAddChatUserArgs) GetReq() *tg.TLMessagesAddChatUser {
	if !p.IsSetReq() {
		return MessagesAddChatUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesAddChatUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesAddChatUserResult struct {
	Success *tg.MessagesInvitedUsers
}

var MessagesAddChatUserResult_Success_DEFAULT *tg.MessagesInvitedUsers

func (p *MessagesAddChatUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesAddChatUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesAddChatUserResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesInvitedUsers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesAddChatUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesAddChatUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesAddChatUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesInvitedUsers)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesAddChatUserResult) GetSuccess() *tg.MessagesInvitedUsers {
	if !p.IsSetSuccess() {
		return MessagesAddChatUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesAddChatUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesInvitedUsers)
}

func (p *MessagesAddChatUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesAddChatUserResult) GetResult() interface{} {
	return p.Success
}

func messagesDeleteChatUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesDeleteChatUserArgs)
	realResult := result.(*MessagesDeleteChatUserResult)
	success, err := handler.(tg.RPCChats).MessagesDeleteChatUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesDeleteChatUserArgs() interface{} {
	return &MessagesDeleteChatUserArgs{}
}

func newMessagesDeleteChatUserResult() interface{} {
	return &MessagesDeleteChatUserResult{}
}

type MessagesDeleteChatUserArgs struct {
	Req *tg.TLMessagesDeleteChatUser
}

func (p *MessagesDeleteChatUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesDeleteChatUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesDeleteChatUserArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesDeleteChatUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesDeleteChatUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesDeleteChatUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesDeleteChatUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesDeleteChatUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesDeleteChatUserArgs_Req_DEFAULT *tg.TLMessagesDeleteChatUser

func (p *MessagesDeleteChatUserArgs) GetReq() *tg.TLMessagesDeleteChatUser {
	if !p.IsSetReq() {
		return MessagesDeleteChatUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesDeleteChatUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesDeleteChatUserResult struct {
	Success *tg.Updates
}

var MessagesDeleteChatUserResult_Success_DEFAULT *tg.Updates

func (p *MessagesDeleteChatUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesDeleteChatUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesDeleteChatUserResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteChatUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesDeleteChatUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesDeleteChatUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteChatUserResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesDeleteChatUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesDeleteChatUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesDeleteChatUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesDeleteChatUserResult) GetResult() interface{} {
	return p.Success
}

func messagesCreateChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesCreateChatArgs)
	realResult := result.(*MessagesCreateChatResult)
	success, err := handler.(tg.RPCChats).MessagesCreateChat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesCreateChatArgs() interface{} {
	return &MessagesCreateChatArgs{}
}

func newMessagesCreateChatResult() interface{} {
	return &MessagesCreateChatResult{}
}

type MessagesCreateChatArgs struct {
	Req *tg.TLMessagesCreateChat
}

func (p *MessagesCreateChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesCreateChatArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesCreateChatArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesCreateChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesCreateChatArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesCreateChatArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesCreateChatArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesCreateChat)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesCreateChatArgs_Req_DEFAULT *tg.TLMessagesCreateChat

func (p *MessagesCreateChatArgs) GetReq() *tg.TLMessagesCreateChat {
	if !p.IsSetReq() {
		return MessagesCreateChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesCreateChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesCreateChatResult struct {
	Success *tg.MessagesInvitedUsers
}

var MessagesCreateChatResult_Success_DEFAULT *tg.MessagesInvitedUsers

func (p *MessagesCreateChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesCreateChatResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesCreateChatResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesInvitedUsers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesCreateChatResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesCreateChatResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesCreateChatResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesInvitedUsers)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesCreateChatResult) GetSuccess() *tg.MessagesInvitedUsers {
	if !p.IsSetSuccess() {
		return MessagesCreateChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesCreateChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesInvitedUsers)
}

func (p *MessagesCreateChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesCreateChatResult) GetResult() interface{} {
	return p.Success
}

func messagesEditChatAdminHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesEditChatAdminArgs)
	realResult := result.(*MessagesEditChatAdminResult)
	success, err := handler.(tg.RPCChats).MessagesEditChatAdmin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesEditChatAdminArgs() interface{} {
	return &MessagesEditChatAdminArgs{}
}

func newMessagesEditChatAdminResult() interface{} {
	return &MessagesEditChatAdminResult{}
}

type MessagesEditChatAdminArgs struct {
	Req *tg.TLMessagesEditChatAdmin
}

func (p *MessagesEditChatAdminArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesEditChatAdminArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesEditChatAdminArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesEditChatAdmin)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesEditChatAdminArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesEditChatAdminArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesEditChatAdminArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesEditChatAdmin)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesEditChatAdminArgs_Req_DEFAULT *tg.TLMessagesEditChatAdmin

func (p *MessagesEditChatAdminArgs) GetReq() *tg.TLMessagesEditChatAdmin {
	if !p.IsSetReq() {
		return MessagesEditChatAdminArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesEditChatAdminArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesEditChatAdminResult struct {
	Success *tg.Bool
}

var MessagesEditChatAdminResult_Success_DEFAULT *tg.Bool

func (p *MessagesEditChatAdminResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesEditChatAdminResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesEditChatAdminResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatAdminResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesEditChatAdminResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesEditChatAdminResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatAdminResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesEditChatAdminResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesEditChatAdminResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesEditChatAdminResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesEditChatAdminResult) GetResult() interface{} {
	return p.Success
}

func messagesMigrateChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesMigrateChatArgs)
	realResult := result.(*MessagesMigrateChatResult)
	success, err := handler.(tg.RPCChats).MessagesMigrateChat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesMigrateChatArgs() interface{} {
	return &MessagesMigrateChatArgs{}
}

func newMessagesMigrateChatResult() interface{} {
	return &MessagesMigrateChatResult{}
}

type MessagesMigrateChatArgs struct {
	Req *tg.TLMessagesMigrateChat
}

func (p *MessagesMigrateChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesMigrateChatArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesMigrateChatArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesMigrateChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesMigrateChatArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesMigrateChatArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesMigrateChatArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesMigrateChat)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesMigrateChatArgs_Req_DEFAULT *tg.TLMessagesMigrateChat

func (p *MessagesMigrateChatArgs) GetReq() *tg.TLMessagesMigrateChat {
	if !p.IsSetReq() {
		return MessagesMigrateChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesMigrateChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesMigrateChatResult struct {
	Success *tg.Updates
}

var MessagesMigrateChatResult_Success_DEFAULT *tg.Updates

func (p *MessagesMigrateChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesMigrateChatResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesMigrateChatResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesMigrateChatResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesMigrateChatResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesMigrateChatResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesMigrateChatResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesMigrateChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesMigrateChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesMigrateChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesMigrateChatResult) GetResult() interface{} {
	return p.Success
}

func messagesGetCommonChatsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetCommonChatsArgs)
	realResult := result.(*MessagesGetCommonChatsResult)
	success, err := handler.(tg.RPCChats).MessagesGetCommonChats(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetCommonChatsArgs() interface{} {
	return &MessagesGetCommonChatsArgs{}
}

func newMessagesGetCommonChatsResult() interface{} {
	return &MessagesGetCommonChatsResult{}
}

type MessagesGetCommonChatsArgs struct {
	Req *tg.TLMessagesGetCommonChats
}

func (p *MessagesGetCommonChatsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetCommonChatsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetCommonChatsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetCommonChats)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetCommonChatsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetCommonChatsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetCommonChatsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetCommonChats)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetCommonChatsArgs_Req_DEFAULT *tg.TLMessagesGetCommonChats

func (p *MessagesGetCommonChatsArgs) GetReq() *tg.TLMessagesGetCommonChats {
	if !p.IsSetReq() {
		return MessagesGetCommonChatsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetCommonChatsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetCommonChatsResult struct {
	Success *tg.MessagesChats
}

var MessagesGetCommonChatsResult_Success_DEFAULT *tg.MessagesChats

func (p *MessagesGetCommonChatsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetCommonChatsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetCommonChatsResult) Unmarshal(in []byte) error {
	msg := new(tg.MessagesChats)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetCommonChatsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetCommonChatsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetCommonChatsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessagesChats)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetCommonChatsResult) GetSuccess() *tg.MessagesChats {
	if !p.IsSetSuccess() {
		return MessagesGetCommonChatsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetCommonChatsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessagesChats)
}

func (p *MessagesGetCommonChatsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetCommonChatsResult) GetResult() interface{} {
	return p.Success
}

func messagesEditChatAboutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesEditChatAboutArgs)
	realResult := result.(*MessagesEditChatAboutResult)
	success, err := handler.(tg.RPCChats).MessagesEditChatAbout(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesEditChatAboutArgs() interface{} {
	return &MessagesEditChatAboutArgs{}
}

func newMessagesEditChatAboutResult() interface{} {
	return &MessagesEditChatAboutResult{}
}

type MessagesEditChatAboutArgs struct {
	Req *tg.TLMessagesEditChatAbout
}

func (p *MessagesEditChatAboutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesEditChatAboutArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesEditChatAboutArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesEditChatAbout)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesEditChatAboutArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesEditChatAboutArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesEditChatAboutArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesEditChatAbout)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesEditChatAboutArgs_Req_DEFAULT *tg.TLMessagesEditChatAbout

func (p *MessagesEditChatAboutArgs) GetReq() *tg.TLMessagesEditChatAbout {
	if !p.IsSetReq() {
		return MessagesEditChatAboutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesEditChatAboutArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesEditChatAboutResult struct {
	Success *tg.Bool
}

var MessagesEditChatAboutResult_Success_DEFAULT *tg.Bool

func (p *MessagesEditChatAboutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesEditChatAboutResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesEditChatAboutResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatAboutResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesEditChatAboutResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesEditChatAboutResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatAboutResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesEditChatAboutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesEditChatAboutResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesEditChatAboutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesEditChatAboutResult) GetResult() interface{} {
	return p.Success
}

func messagesEditChatDefaultBannedRightsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesEditChatDefaultBannedRightsArgs)
	realResult := result.(*MessagesEditChatDefaultBannedRightsResult)
	success, err := handler.(tg.RPCChats).MessagesEditChatDefaultBannedRights(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesEditChatDefaultBannedRightsArgs() interface{} {
	return &MessagesEditChatDefaultBannedRightsArgs{}
}

func newMessagesEditChatDefaultBannedRightsResult() interface{} {
	return &MessagesEditChatDefaultBannedRightsResult{}
}

type MessagesEditChatDefaultBannedRightsArgs struct {
	Req *tg.TLMessagesEditChatDefaultBannedRights
}

func (p *MessagesEditChatDefaultBannedRightsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesEditChatDefaultBannedRightsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesEditChatDefaultBannedRightsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesEditChatDefaultBannedRights)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesEditChatDefaultBannedRightsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesEditChatDefaultBannedRightsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesEditChatDefaultBannedRightsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesEditChatDefaultBannedRights)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesEditChatDefaultBannedRightsArgs_Req_DEFAULT *tg.TLMessagesEditChatDefaultBannedRights

func (p *MessagesEditChatDefaultBannedRightsArgs) GetReq() *tg.TLMessagesEditChatDefaultBannedRights {
	if !p.IsSetReq() {
		return MessagesEditChatDefaultBannedRightsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesEditChatDefaultBannedRightsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesEditChatDefaultBannedRightsResult struct {
	Success *tg.Updates
}

var MessagesEditChatDefaultBannedRightsResult_Success_DEFAULT *tg.Updates

func (p *MessagesEditChatDefaultBannedRightsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesEditChatDefaultBannedRightsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesEditChatDefaultBannedRightsResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatDefaultBannedRightsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesEditChatDefaultBannedRightsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesEditChatDefaultBannedRightsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesEditChatDefaultBannedRightsResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return MessagesEditChatDefaultBannedRightsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesEditChatDefaultBannedRightsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *MessagesEditChatDefaultBannedRightsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesEditChatDefaultBannedRightsResult) GetResult() interface{} {
	return p.Success
}

func messagesDeleteChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesDeleteChatArgs)
	realResult := result.(*MessagesDeleteChatResult)
	success, err := handler.(tg.RPCChats).MessagesDeleteChat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesDeleteChatArgs() interface{} {
	return &MessagesDeleteChatArgs{}
}

func newMessagesDeleteChatResult() interface{} {
	return &MessagesDeleteChatResult{}
}

type MessagesDeleteChatArgs struct {
	Req *tg.TLMessagesDeleteChat
}

func (p *MessagesDeleteChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesDeleteChatArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesDeleteChatArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesDeleteChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesDeleteChatArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesDeleteChatArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesDeleteChatArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesDeleteChat)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesDeleteChatArgs_Req_DEFAULT *tg.TLMessagesDeleteChat

func (p *MessagesDeleteChatArgs) GetReq() *tg.TLMessagesDeleteChat {
	if !p.IsSetReq() {
		return MessagesDeleteChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesDeleteChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesDeleteChatResult struct {
	Success *tg.Bool
}

var MessagesDeleteChatResult_Success_DEFAULT *tg.Bool

func (p *MessagesDeleteChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesDeleteChatResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesDeleteChatResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteChatResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesDeleteChatResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesDeleteChatResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesDeleteChatResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MessagesDeleteChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesDeleteChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MessagesDeleteChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesDeleteChatResult) GetResult() interface{} {
	return p.Success
}

func messagesGetMessageReadParticipantsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetMessageReadParticipantsArgs)
	realResult := result.(*MessagesGetMessageReadParticipantsResult)
	success, err := handler.(tg.RPCChats).MessagesGetMessageReadParticipants(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetMessageReadParticipantsArgs() interface{} {
	return &MessagesGetMessageReadParticipantsArgs{}
}

func newMessagesGetMessageReadParticipantsResult() interface{} {
	return &MessagesGetMessageReadParticipantsResult{}
}

type MessagesGetMessageReadParticipantsArgs struct {
	Req *tg.TLMessagesGetMessageReadParticipants
}

func (p *MessagesGetMessageReadParticipantsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetMessageReadParticipantsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetMessageReadParticipantsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetMessageReadParticipants)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetMessageReadParticipantsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetMessageReadParticipantsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetMessageReadParticipantsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetMessageReadParticipants)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetMessageReadParticipantsArgs_Req_DEFAULT *tg.TLMessagesGetMessageReadParticipants

func (p *MessagesGetMessageReadParticipantsArgs) GetReq() *tg.TLMessagesGetMessageReadParticipants {
	if !p.IsSetReq() {
		return MessagesGetMessageReadParticipantsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetMessageReadParticipantsArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetMessageReadParticipantsResult struct {
	Success *tg.VectorReadParticipantDate
}

var MessagesGetMessageReadParticipantsResult_Success_DEFAULT *tg.VectorReadParticipantDate

func (p *MessagesGetMessageReadParticipantsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetMessageReadParticipantsResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetMessageReadParticipantsResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorReadParticipantDate)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetMessageReadParticipantsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetMessageReadParticipantsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetMessageReadParticipantsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorReadParticipantDate)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetMessageReadParticipantsResult) GetSuccess() *tg.VectorReadParticipantDate {
	if !p.IsSetSuccess() {
		return MessagesGetMessageReadParticipantsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetMessageReadParticipantsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorReadParticipantDate)
}

func (p *MessagesGetMessageReadParticipantsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetMessageReadParticipantsResult) GetResult() interface{} {
	return p.Success
}

func channelsConvertToGigagroupHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsConvertToGigagroupArgs)
	realResult := result.(*ChannelsConvertToGigagroupResult)
	success, err := handler.(tg.RPCChats).ChannelsConvertToGigagroup(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsConvertToGigagroupArgs() interface{} {
	return &ChannelsConvertToGigagroupArgs{}
}

func newChannelsConvertToGigagroupResult() interface{} {
	return &ChannelsConvertToGigagroupResult{}
}

type ChannelsConvertToGigagroupArgs struct {
	Req *tg.TLChannelsConvertToGigagroup
}

func (p *ChannelsConvertToGigagroupArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsConvertToGigagroupArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsConvertToGigagroupArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsConvertToGigagroup)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsConvertToGigagroupArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsConvertToGigagroupArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsConvertToGigagroupArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsConvertToGigagroup)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsConvertToGigagroupArgs_Req_DEFAULT *tg.TLChannelsConvertToGigagroup

func (p *ChannelsConvertToGigagroupArgs) GetReq() *tg.TLChannelsConvertToGigagroup {
	if !p.IsSetReq() {
		return ChannelsConvertToGigagroupArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsConvertToGigagroupArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsConvertToGigagroupResult struct {
	Success *tg.Updates
}

var ChannelsConvertToGigagroupResult_Success_DEFAULT *tg.Updates

func (p *ChannelsConvertToGigagroupResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsConvertToGigagroupResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsConvertToGigagroupResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsConvertToGigagroupResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsConvertToGigagroupResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsConvertToGigagroupResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsConvertToGigagroupResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return ChannelsConvertToGigagroupResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsConvertToGigagroupResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *ChannelsConvertToGigagroupResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsConvertToGigagroupResult) GetResult() interface{} {
	return p.Success
}

func channelsSetEmojiStickersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChannelsSetEmojiStickersArgs)
	realResult := result.(*ChannelsSetEmojiStickersResult)
	success, err := handler.(tg.RPCChats).ChannelsSetEmojiStickers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChannelsSetEmojiStickersArgs() interface{} {
	return &ChannelsSetEmojiStickersArgs{}
}

func newChannelsSetEmojiStickersResult() interface{} {
	return &ChannelsSetEmojiStickersResult{}
}

type ChannelsSetEmojiStickersArgs struct {
	Req *tg.TLChannelsSetEmojiStickers
}

func (p *ChannelsSetEmojiStickersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChannelsSetEmojiStickersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChannelsSetEmojiStickersArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLChannelsSetEmojiStickers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChannelsSetEmojiStickersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChannelsSetEmojiStickersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChannelsSetEmojiStickersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLChannelsSetEmojiStickers)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChannelsSetEmojiStickersArgs_Req_DEFAULT *tg.TLChannelsSetEmojiStickers

func (p *ChannelsSetEmojiStickersArgs) GetReq() *tg.TLChannelsSetEmojiStickers {
	if !p.IsSetReq() {
		return ChannelsSetEmojiStickersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChannelsSetEmojiStickersArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChannelsSetEmojiStickersResult struct {
	Success *tg.Bool
}

var ChannelsSetEmojiStickersResult_Success_DEFAULT *tg.Bool

func (p *ChannelsSetEmojiStickersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChannelsSetEmojiStickersResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChannelsSetEmojiStickersResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsSetEmojiStickersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChannelsSetEmojiStickersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChannelsSetEmojiStickersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChannelsSetEmojiStickersResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ChannelsSetEmojiStickersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChannelsSetEmojiStickersResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ChannelsSetEmojiStickersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChannelsSetEmojiStickersResult) GetResult() interface{} {
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

func (p *kClient) MessagesGetChats(ctx context.Context, req *tg.TLMessagesGetChats) (r *tg.MessagesChats, err error) {
	var _args MessagesGetChatsArgs
	_args.Req = req
	var _result MessagesGetChatsResult
	if err = p.c.Call(ctx, "messages.getChats", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetFullChat(ctx context.Context, req *tg.TLMessagesGetFullChat) (r *tg.MessagesChatFull, err error) {
	var _args MessagesGetFullChatArgs
	_args.Req = req
	var _result MessagesGetFullChatResult
	if err = p.c.Call(ctx, "messages.getFullChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesEditChatTitle(ctx context.Context, req *tg.TLMessagesEditChatTitle) (r *tg.Updates, err error) {
	var _args MessagesEditChatTitleArgs
	_args.Req = req
	var _result MessagesEditChatTitleResult
	if err = p.c.Call(ctx, "messages.editChatTitle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesEditChatPhoto(ctx context.Context, req *tg.TLMessagesEditChatPhoto) (r *tg.Updates, err error) {
	var _args MessagesEditChatPhotoArgs
	_args.Req = req
	var _result MessagesEditChatPhotoResult
	if err = p.c.Call(ctx, "messages.editChatPhoto", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesAddChatUser(ctx context.Context, req *tg.TLMessagesAddChatUser) (r *tg.MessagesInvitedUsers, err error) {
	var _args MessagesAddChatUserArgs
	_args.Req = req
	var _result MessagesAddChatUserResult
	if err = p.c.Call(ctx, "messages.addChatUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesDeleteChatUser(ctx context.Context, req *tg.TLMessagesDeleteChatUser) (r *tg.Updates, err error) {
	var _args MessagesDeleteChatUserArgs
	_args.Req = req
	var _result MessagesDeleteChatUserResult
	if err = p.c.Call(ctx, "messages.deleteChatUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesCreateChat(ctx context.Context, req *tg.TLMessagesCreateChat) (r *tg.MessagesInvitedUsers, err error) {
	var _args MessagesCreateChatArgs
	_args.Req = req
	var _result MessagesCreateChatResult
	if err = p.c.Call(ctx, "messages.createChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesEditChatAdmin(ctx context.Context, req *tg.TLMessagesEditChatAdmin) (r *tg.Bool, err error) {
	var _args MessagesEditChatAdminArgs
	_args.Req = req
	var _result MessagesEditChatAdminResult
	if err = p.c.Call(ctx, "messages.editChatAdmin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesMigrateChat(ctx context.Context, req *tg.TLMessagesMigrateChat) (r *tg.Updates, err error) {
	var _args MessagesMigrateChatArgs
	_args.Req = req
	var _result MessagesMigrateChatResult
	if err = p.c.Call(ctx, "messages.migrateChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetCommonChats(ctx context.Context, req *tg.TLMessagesGetCommonChats) (r *tg.MessagesChats, err error) {
	var _args MessagesGetCommonChatsArgs
	_args.Req = req
	var _result MessagesGetCommonChatsResult
	if err = p.c.Call(ctx, "messages.getCommonChats", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesEditChatAbout(ctx context.Context, req *tg.TLMessagesEditChatAbout) (r *tg.Bool, err error) {
	var _args MessagesEditChatAboutArgs
	_args.Req = req
	var _result MessagesEditChatAboutResult
	if err = p.c.Call(ctx, "messages.editChatAbout", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesEditChatDefaultBannedRights(ctx context.Context, req *tg.TLMessagesEditChatDefaultBannedRights) (r *tg.Updates, err error) {
	var _args MessagesEditChatDefaultBannedRightsArgs
	_args.Req = req
	var _result MessagesEditChatDefaultBannedRightsResult
	if err = p.c.Call(ctx, "messages.editChatDefaultBannedRights", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesDeleteChat(ctx context.Context, req *tg.TLMessagesDeleteChat) (r *tg.Bool, err error) {
	var _args MessagesDeleteChatArgs
	_args.Req = req
	var _result MessagesDeleteChatResult
	if err = p.c.Call(ctx, "messages.deleteChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessagesGetMessageReadParticipants(ctx context.Context, req *tg.TLMessagesGetMessageReadParticipants) (r *tg.VectorReadParticipantDate, err error) {
	var _args MessagesGetMessageReadParticipantsArgs
	_args.Req = req
	var _result MessagesGetMessageReadParticipantsResult
	if err = p.c.Call(ctx, "messages.getMessageReadParticipants", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsConvertToGigagroup(ctx context.Context, req *tg.TLChannelsConvertToGigagroup) (r *tg.Updates, err error) {
	var _args ChannelsConvertToGigagroupArgs
	_args.Req = req
	var _result ChannelsConvertToGigagroupResult
	if err = p.c.Call(ctx, "channels.convertToGigagroup", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChannelsSetEmojiStickers(ctx context.Context, req *tg.TLChannelsSetEmojiStickers) (r *tg.Bool, err error) {
	var _args ChannelsSetEmojiStickersArgs
	_args.Req = req
	var _result ChannelsSetEmojiStickersResult
	if err = p.c.Call(ctx, "channels.setEmojiStickers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
