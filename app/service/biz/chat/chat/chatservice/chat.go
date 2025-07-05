/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chatservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"chat.getMutableChat": kitex.NewMethodInfo(
		getMutableChatHandler,
		newGetMutableChatArgs,
		newGetMutableChatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getChatListByIdList": kitex.NewMethodInfo(
		getChatListByIdListHandler,
		newGetChatListByIdListArgs,
		newGetChatListByIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getChatBySelfId": kitex.NewMethodInfo(
		getChatBySelfIdHandler,
		newGetChatBySelfIdArgs,
		newGetChatBySelfIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.createChat2": kitex.NewMethodInfo(
		createChat2Handler,
		newCreateChat2Args,
		newCreateChat2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.deleteChat": kitex.NewMethodInfo(
		deleteChatHandler,
		newDeleteChatArgs,
		newDeleteChatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.deleteChatUser": kitex.NewMethodInfo(
		deleteChatUserHandler,
		newDeleteChatUserArgs,
		newDeleteChatUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.editChatTitle": kitex.NewMethodInfo(
		editChatTitleHandler,
		newEditChatTitleArgs,
		newEditChatTitleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.editChatAbout": kitex.NewMethodInfo(
		editChatAboutHandler,
		newEditChatAboutArgs,
		newEditChatAboutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.editChatPhoto": kitex.NewMethodInfo(
		editChatPhotoHandler,
		newEditChatPhotoArgs,
		newEditChatPhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.editChatAdmin": kitex.NewMethodInfo(
		editChatAdminHandler,
		newEditChatAdminArgs,
		newEditChatAdminResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.editChatDefaultBannedRights": kitex.NewMethodInfo(
		editChatDefaultBannedRightsHandler,
		newEditChatDefaultBannedRightsArgs,
		newEditChatDefaultBannedRightsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.addChatUser": kitex.NewMethodInfo(
		addChatUserHandler,
		newAddChatUserArgs,
		newAddChatUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getMutableChatByLink": kitex.NewMethodInfo(
		getMutableChatByLinkHandler,
		newGetMutableChatByLinkArgs,
		newGetMutableChatByLinkResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.toggleNoForwards": kitex.NewMethodInfo(
		toggleNoForwardsHandler,
		newToggleNoForwardsArgs,
		newToggleNoForwardsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.migratedToChannel": kitex.NewMethodInfo(
		migratedToChannelHandler,
		newMigratedToChannelArgs,
		newMigratedToChannelResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getChatParticipantIdList": kitex.NewMethodInfo(
		getChatParticipantIdListHandler,
		newGetChatParticipantIdListArgs,
		newGetChatParticipantIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getUsersChatIdList": kitex.NewMethodInfo(
		getUsersChatIdListHandler,
		newGetUsersChatIdListArgs,
		newGetUsersChatIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getMyChatList": kitex.NewMethodInfo(
		getMyChatListHandler,
		newGetMyChatListArgs,
		newGetMyChatListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.exportChatInvite": kitex.NewMethodInfo(
		exportChatInviteHandler,
		newExportChatInviteArgs,
		newExportChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getAdminsWithInvites": kitex.NewMethodInfo(
		getAdminsWithInvitesHandler,
		newGetAdminsWithInvitesArgs,
		newGetAdminsWithInvitesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getExportedChatInvite": kitex.NewMethodInfo(
		getExportedChatInviteHandler,
		newGetExportedChatInviteArgs,
		newGetExportedChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getExportedChatInvites": kitex.NewMethodInfo(
		getExportedChatInvitesHandler,
		newGetExportedChatInvitesArgs,
		newGetExportedChatInvitesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.checkChatInvite": kitex.NewMethodInfo(
		checkChatInviteHandler,
		newCheckChatInviteArgs,
		newCheckChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.importChatInvite": kitex.NewMethodInfo(
		importChatInviteHandler,
		newImportChatInviteArgs,
		newImportChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getChatInviteImporters": kitex.NewMethodInfo(
		getChatInviteImportersHandler,
		newGetChatInviteImportersArgs,
		newGetChatInviteImportersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.deleteExportedChatInvite": kitex.NewMethodInfo(
		deleteExportedChatInviteHandler,
		newDeleteExportedChatInviteArgs,
		newDeleteExportedChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.deleteRevokedExportedChatInvites": kitex.NewMethodInfo(
		deleteRevokedExportedChatInvitesHandler,
		newDeleteRevokedExportedChatInvitesArgs,
		newDeleteRevokedExportedChatInvitesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.editExportedChatInvite": kitex.NewMethodInfo(
		editExportedChatInviteHandler,
		newEditExportedChatInviteArgs,
		newEditExportedChatInviteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.setChatAvailableReactions": kitex.NewMethodInfo(
		setChatAvailableReactionsHandler,
		newSetChatAvailableReactionsArgs,
		newSetChatAvailableReactionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.setHistoryTTL": kitex.NewMethodInfo(
		setHistoryTTLHandler,
		newSetHistoryTTLArgs,
		newSetHistoryTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.search": kitex.NewMethodInfo(
		searchHandler,
		newSearchArgs,
		newSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.getRecentChatInviteRequesters": kitex.NewMethodInfo(
		getRecentChatInviteRequestersHandler,
		newGetRecentChatInviteRequestersArgs,
		newGetRecentChatInviteRequestersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.hideChatJoinRequests": kitex.NewMethodInfo(
		hideChatJoinRequestsHandler,
		newHideChatJoinRequestsArgs,
		newHideChatJoinRequestsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"chat.importChatInvite2": kitex.NewMethodInfo(
		importChatInvite2Handler,
		newImportChatInvite2Args,
		newImportChatInvite2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	chatServiceServiceInfo                = NewServiceInfo()
	chatServiceServiceInfoForClient       = NewServiceInfoForClient()
	chatServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCChat", chatServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCChat", chatServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCChat", chatServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return chatServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return chatServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return chatServiceServiceInfoForClient
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
	serviceName := "RPCChat"
	handlerType := (*chat.RPCChat)(nil)
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
		"PackageName": "chat",
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

func getMutableChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetMutableChatArgs)
	realResult := result.(*GetMutableChatResult)
	success, err := handler.(chat.RPCChat).ChatGetMutableChat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetMutableChatArgs() interface{} {
	return &GetMutableChatArgs{}
}

func newGetMutableChatResult() interface{} {
	return &GetMutableChatResult{}
}

type GetMutableChatArgs struct {
	Req *chat.TLChatGetMutableChat
}

func (p *GetMutableChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetMutableChatArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetMutableChatArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetMutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetMutableChatArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetMutableChatArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetMutableChatArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetMutableChat)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetMutableChatArgs_Req_DEFAULT *chat.TLChatGetMutableChat

func (p *GetMutableChatArgs) GetReq() *chat.TLChatGetMutableChat {
	if !p.IsSetReq() {
		return GetMutableChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMutableChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetMutableChatResult struct {
	Success *tg.MutableChat
}

var GetMutableChatResult_Success_DEFAULT *tg.MutableChat

func (p *GetMutableChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetMutableChatResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetMutableChatResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMutableChatResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetMutableChatResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetMutableChatResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMutableChatResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return GetMutableChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMutableChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *GetMutableChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMutableChatResult) GetResult() interface{} {
	return p.Success
}

func getChatListByIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetChatListByIdListArgs)
	realResult := result.(*GetChatListByIdListResult)
	success, err := handler.(chat.RPCChat).ChatGetChatListByIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetChatListByIdListArgs() interface{} {
	return &GetChatListByIdListArgs{}
}

func newGetChatListByIdListResult() interface{} {
	return &GetChatListByIdListResult{}
}

type GetChatListByIdListArgs struct {
	Req *chat.TLChatGetChatListByIdList
}

func (p *GetChatListByIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChatListByIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetChatListByIdListArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetChatListByIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetChatListByIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetChatListByIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetChatListByIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetChatListByIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetChatListByIdListArgs_Req_DEFAULT *chat.TLChatGetChatListByIdList

func (p *GetChatListByIdListArgs) GetReq() *chat.TLChatGetChatListByIdList {
	if !p.IsSetReq() {
		return GetChatListByIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChatListByIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetChatListByIdListResult struct {
	Success *chat.VectorMutableChat
}

var GetChatListByIdListResult_Success_DEFAULT *chat.VectorMutableChat

func (p *GetChatListByIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChatListByIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetChatListByIdListResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorMutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatListByIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetChatListByIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetChatListByIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.VectorMutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatListByIdListResult) GetSuccess() *chat.VectorMutableChat {
	if !p.IsSetSuccess() {
		return GetChatListByIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChatListByIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorMutableChat)
}

func (p *GetChatListByIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetChatListByIdListResult) GetResult() interface{} {
	return p.Success
}

func getChatBySelfIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetChatBySelfIdArgs)
	realResult := result.(*GetChatBySelfIdResult)
	success, err := handler.(chat.RPCChat).ChatGetChatBySelfId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetChatBySelfIdArgs() interface{} {
	return &GetChatBySelfIdArgs{}
}

func newGetChatBySelfIdResult() interface{} {
	return &GetChatBySelfIdResult{}
}

type GetChatBySelfIdArgs struct {
	Req *chat.TLChatGetChatBySelfId
}

func (p *GetChatBySelfIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChatBySelfIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetChatBySelfIdArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetChatBySelfId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetChatBySelfIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetChatBySelfIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetChatBySelfIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetChatBySelfId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetChatBySelfIdArgs_Req_DEFAULT *chat.TLChatGetChatBySelfId

func (p *GetChatBySelfIdArgs) GetReq() *chat.TLChatGetChatBySelfId {
	if !p.IsSetReq() {
		return GetChatBySelfIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChatBySelfIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetChatBySelfIdResult struct {
	Success *tg.MutableChat
}

var GetChatBySelfIdResult_Success_DEFAULT *tg.MutableChat

func (p *GetChatBySelfIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChatBySelfIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetChatBySelfIdResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatBySelfIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetChatBySelfIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetChatBySelfIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatBySelfIdResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return GetChatBySelfIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChatBySelfIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *GetChatBySelfIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetChatBySelfIdResult) GetResult() interface{} {
	return p.Success
}

func createChat2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CreateChat2Args)
	realResult := result.(*CreateChat2Result)
	success, err := handler.(chat.RPCChat).ChatCreateChat2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCreateChat2Args() interface{} {
	return &CreateChat2Args{}
}

func newCreateChat2Result() interface{} {
	return &CreateChat2Result{}
}

type CreateChat2Args struct {
	Req *chat.TLChatCreateChat2
}

func (p *CreateChat2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateChat2Args")
	}
	return json.Marshal(p.Req)
}

func (p *CreateChat2Args) Unmarshal(in []byte) error {
	msg := new(chat.TLChatCreateChat2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CreateChat2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CreateChat2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *CreateChat2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatCreateChat2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CreateChat2Args_Req_DEFAULT *chat.TLChatCreateChat2

func (p *CreateChat2Args) GetReq() *chat.TLChatCreateChat2 {
	if !p.IsSetReq() {
		return CreateChat2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateChat2Args) IsSetReq() bool {
	return p.Req != nil
}

type CreateChat2Result struct {
	Success *tg.MutableChat
}

var CreateChat2Result_Success_DEFAULT *tg.MutableChat

func (p *CreateChat2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateChat2Result")
	}
	return json.Marshal(p.Success)
}

func (p *CreateChat2Result) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateChat2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CreateChat2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *CreateChat2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateChat2Result) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return CreateChat2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateChat2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *CreateChat2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateChat2Result) GetResult() interface{} {
	return p.Success
}

func deleteChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteChatArgs)
	realResult := result.(*DeleteChatResult)
	success, err := handler.(chat.RPCChat).ChatDeleteChat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteChatArgs() interface{} {
	return &DeleteChatArgs{}
}

func newDeleteChatResult() interface{} {
	return &DeleteChatResult{}
}

type DeleteChatArgs struct {
	Req *chat.TLChatDeleteChat
}

func (p *DeleteChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteChatArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteChatArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatDeleteChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteChatArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteChatArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteChatArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatDeleteChat)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteChatArgs_Req_DEFAULT *chat.TLChatDeleteChat

func (p *DeleteChatArgs) GetReq() *chat.TLChatDeleteChat {
	if !p.IsSetReq() {
		return DeleteChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteChatResult struct {
	Success *tg.MutableChat
}

var DeleteChatResult_Success_DEFAULT *tg.MutableChat

func (p *DeleteChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteChatResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteChatResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteChatResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteChatResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteChatResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteChatResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return DeleteChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *DeleteChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteChatResult) GetResult() interface{} {
	return p.Success
}

func deleteChatUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteChatUserArgs)
	realResult := result.(*DeleteChatUserResult)
	success, err := handler.(chat.RPCChat).ChatDeleteChatUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteChatUserArgs() interface{} {
	return &DeleteChatUserArgs{}
}

func newDeleteChatUserResult() interface{} {
	return &DeleteChatUserResult{}
}

type DeleteChatUserArgs struct {
	Req *chat.TLChatDeleteChatUser
}

func (p *DeleteChatUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteChatUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteChatUserArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatDeleteChatUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteChatUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteChatUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteChatUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatDeleteChatUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteChatUserArgs_Req_DEFAULT *chat.TLChatDeleteChatUser

func (p *DeleteChatUserArgs) GetReq() *chat.TLChatDeleteChatUser {
	if !p.IsSetReq() {
		return DeleteChatUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteChatUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteChatUserResult struct {
	Success *tg.MutableChat
}

var DeleteChatUserResult_Success_DEFAULT *tg.MutableChat

func (p *DeleteChatUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteChatUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteChatUserResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteChatUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteChatUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteChatUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteChatUserResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return DeleteChatUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteChatUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *DeleteChatUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteChatUserResult) GetResult() interface{} {
	return p.Success
}

func editChatTitleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditChatTitleArgs)
	realResult := result.(*EditChatTitleResult)
	success, err := handler.(chat.RPCChat).ChatEditChatTitle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditChatTitleArgs() interface{} {
	return &EditChatTitleArgs{}
}

func newEditChatTitleResult() interface{} {
	return &EditChatTitleResult{}
}

type EditChatTitleArgs struct {
	Req *chat.TLChatEditChatTitle
}

func (p *EditChatTitleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditChatTitleArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditChatTitleArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatEditChatTitle)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditChatTitleArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditChatTitleArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditChatTitleArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatEditChatTitle)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditChatTitleArgs_Req_DEFAULT *chat.TLChatEditChatTitle

func (p *EditChatTitleArgs) GetReq() *chat.TLChatEditChatTitle {
	if !p.IsSetReq() {
		return EditChatTitleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditChatTitleArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditChatTitleResult struct {
	Success *tg.MutableChat
}

var EditChatTitleResult_Success_DEFAULT *tg.MutableChat

func (p *EditChatTitleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditChatTitleResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditChatTitleResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatTitleResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditChatTitleResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditChatTitleResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatTitleResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return EditChatTitleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditChatTitleResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *EditChatTitleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditChatTitleResult) GetResult() interface{} {
	return p.Success
}

func editChatAboutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditChatAboutArgs)
	realResult := result.(*EditChatAboutResult)
	success, err := handler.(chat.RPCChat).ChatEditChatAbout(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditChatAboutArgs() interface{} {
	return &EditChatAboutArgs{}
}

func newEditChatAboutResult() interface{} {
	return &EditChatAboutResult{}
}

type EditChatAboutArgs struct {
	Req *chat.TLChatEditChatAbout
}

func (p *EditChatAboutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditChatAboutArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditChatAboutArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatEditChatAbout)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditChatAboutArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditChatAboutArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditChatAboutArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatEditChatAbout)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditChatAboutArgs_Req_DEFAULT *chat.TLChatEditChatAbout

func (p *EditChatAboutArgs) GetReq() *chat.TLChatEditChatAbout {
	if !p.IsSetReq() {
		return EditChatAboutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditChatAboutArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditChatAboutResult struct {
	Success *tg.MutableChat
}

var EditChatAboutResult_Success_DEFAULT *tg.MutableChat

func (p *EditChatAboutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditChatAboutResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditChatAboutResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatAboutResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditChatAboutResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditChatAboutResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatAboutResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return EditChatAboutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditChatAboutResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *EditChatAboutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditChatAboutResult) GetResult() interface{} {
	return p.Success
}

func editChatPhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditChatPhotoArgs)
	realResult := result.(*EditChatPhotoResult)
	success, err := handler.(chat.RPCChat).ChatEditChatPhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditChatPhotoArgs() interface{} {
	return &EditChatPhotoArgs{}
}

func newEditChatPhotoResult() interface{} {
	return &EditChatPhotoResult{}
}

type EditChatPhotoArgs struct {
	Req *chat.TLChatEditChatPhoto
}

func (p *EditChatPhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditChatPhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditChatPhotoArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatEditChatPhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditChatPhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditChatPhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditChatPhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatEditChatPhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditChatPhotoArgs_Req_DEFAULT *chat.TLChatEditChatPhoto

func (p *EditChatPhotoArgs) GetReq() *chat.TLChatEditChatPhoto {
	if !p.IsSetReq() {
		return EditChatPhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditChatPhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditChatPhotoResult struct {
	Success *tg.MutableChat
}

var EditChatPhotoResult_Success_DEFAULT *tg.MutableChat

func (p *EditChatPhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditChatPhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditChatPhotoResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatPhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditChatPhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditChatPhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatPhotoResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return EditChatPhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditChatPhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *EditChatPhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditChatPhotoResult) GetResult() interface{} {
	return p.Success
}

func editChatAdminHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditChatAdminArgs)
	realResult := result.(*EditChatAdminResult)
	success, err := handler.(chat.RPCChat).ChatEditChatAdmin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditChatAdminArgs() interface{} {
	return &EditChatAdminArgs{}
}

func newEditChatAdminResult() interface{} {
	return &EditChatAdminResult{}
}

type EditChatAdminArgs struct {
	Req *chat.TLChatEditChatAdmin
}

func (p *EditChatAdminArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditChatAdminArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditChatAdminArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatEditChatAdmin)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditChatAdminArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditChatAdminArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditChatAdminArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatEditChatAdmin)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditChatAdminArgs_Req_DEFAULT *chat.TLChatEditChatAdmin

func (p *EditChatAdminArgs) GetReq() *chat.TLChatEditChatAdmin {
	if !p.IsSetReq() {
		return EditChatAdminArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditChatAdminArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditChatAdminResult struct {
	Success *tg.MutableChat
}

var EditChatAdminResult_Success_DEFAULT *tg.MutableChat

func (p *EditChatAdminResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditChatAdminResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditChatAdminResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatAdminResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditChatAdminResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditChatAdminResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatAdminResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return EditChatAdminResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditChatAdminResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *EditChatAdminResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditChatAdminResult) GetResult() interface{} {
	return p.Success
}

func editChatDefaultBannedRightsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditChatDefaultBannedRightsArgs)
	realResult := result.(*EditChatDefaultBannedRightsResult)
	success, err := handler.(chat.RPCChat).ChatEditChatDefaultBannedRights(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditChatDefaultBannedRightsArgs() interface{} {
	return &EditChatDefaultBannedRightsArgs{}
}

func newEditChatDefaultBannedRightsResult() interface{} {
	return &EditChatDefaultBannedRightsResult{}
}

type EditChatDefaultBannedRightsArgs struct {
	Req *chat.TLChatEditChatDefaultBannedRights
}

func (p *EditChatDefaultBannedRightsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditChatDefaultBannedRightsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditChatDefaultBannedRightsArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatEditChatDefaultBannedRights)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditChatDefaultBannedRightsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditChatDefaultBannedRightsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditChatDefaultBannedRightsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatEditChatDefaultBannedRights)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditChatDefaultBannedRightsArgs_Req_DEFAULT *chat.TLChatEditChatDefaultBannedRights

func (p *EditChatDefaultBannedRightsArgs) GetReq() *chat.TLChatEditChatDefaultBannedRights {
	if !p.IsSetReq() {
		return EditChatDefaultBannedRightsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditChatDefaultBannedRightsArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditChatDefaultBannedRightsResult struct {
	Success *tg.MutableChat
}

var EditChatDefaultBannedRightsResult_Success_DEFAULT *tg.MutableChat

func (p *EditChatDefaultBannedRightsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditChatDefaultBannedRightsResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditChatDefaultBannedRightsResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatDefaultBannedRightsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditChatDefaultBannedRightsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditChatDefaultBannedRightsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditChatDefaultBannedRightsResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return EditChatDefaultBannedRightsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditChatDefaultBannedRightsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *EditChatDefaultBannedRightsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditChatDefaultBannedRightsResult) GetResult() interface{} {
	return p.Success
}

func addChatUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AddChatUserArgs)
	realResult := result.(*AddChatUserResult)
	success, err := handler.(chat.RPCChat).ChatAddChatUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAddChatUserArgs() interface{} {
	return &AddChatUserArgs{}
}

func newAddChatUserResult() interface{} {
	return &AddChatUserResult{}
}

type AddChatUserArgs struct {
	Req *chat.TLChatAddChatUser
}

func (p *AddChatUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AddChatUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AddChatUserArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatAddChatUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AddChatUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AddChatUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AddChatUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatAddChatUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AddChatUserArgs_Req_DEFAULT *chat.TLChatAddChatUser

func (p *AddChatUserArgs) GetReq() *chat.TLChatAddChatUser {
	if !p.IsSetReq() {
		return AddChatUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddChatUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type AddChatUserResult struct {
	Success *tg.MutableChat
}

var AddChatUserResult_Success_DEFAULT *tg.MutableChat

func (p *AddChatUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AddChatUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *AddChatUserResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddChatUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AddChatUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AddChatUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddChatUserResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return AddChatUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddChatUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *AddChatUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddChatUserResult) GetResult() interface{} {
	return p.Success
}

func getMutableChatByLinkHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetMutableChatByLinkArgs)
	realResult := result.(*GetMutableChatByLinkResult)
	success, err := handler.(chat.RPCChat).ChatGetMutableChatByLink(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetMutableChatByLinkArgs() interface{} {
	return &GetMutableChatByLinkArgs{}
}

func newGetMutableChatByLinkResult() interface{} {
	return &GetMutableChatByLinkResult{}
}

type GetMutableChatByLinkArgs struct {
	Req *chat.TLChatGetMutableChatByLink
}

func (p *GetMutableChatByLinkArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetMutableChatByLinkArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetMutableChatByLinkArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetMutableChatByLink)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetMutableChatByLinkArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetMutableChatByLinkArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetMutableChatByLinkArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetMutableChatByLink)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetMutableChatByLinkArgs_Req_DEFAULT *chat.TLChatGetMutableChatByLink

func (p *GetMutableChatByLinkArgs) GetReq() *chat.TLChatGetMutableChatByLink {
	if !p.IsSetReq() {
		return GetMutableChatByLinkArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMutableChatByLinkArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetMutableChatByLinkResult struct {
	Success *tg.MutableChat
}

var GetMutableChatByLinkResult_Success_DEFAULT *tg.MutableChat

func (p *GetMutableChatByLinkResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetMutableChatByLinkResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetMutableChatByLinkResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMutableChatByLinkResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetMutableChatByLinkResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetMutableChatByLinkResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMutableChatByLinkResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return GetMutableChatByLinkResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMutableChatByLinkResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *GetMutableChatByLinkResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMutableChatByLinkResult) GetResult() interface{} {
	return p.Success
}

func toggleNoForwardsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ToggleNoForwardsArgs)
	realResult := result.(*ToggleNoForwardsResult)
	success, err := handler.(chat.RPCChat).ChatToggleNoForwards(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newToggleNoForwardsArgs() interface{} {
	return &ToggleNoForwardsArgs{}
}

func newToggleNoForwardsResult() interface{} {
	return &ToggleNoForwardsResult{}
}

type ToggleNoForwardsArgs struct {
	Req *chat.TLChatToggleNoForwards
}

func (p *ToggleNoForwardsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ToggleNoForwardsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ToggleNoForwardsArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatToggleNoForwards)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ToggleNoForwardsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ToggleNoForwardsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ToggleNoForwardsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatToggleNoForwards)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ToggleNoForwardsArgs_Req_DEFAULT *chat.TLChatToggleNoForwards

func (p *ToggleNoForwardsArgs) GetReq() *chat.TLChatToggleNoForwards {
	if !p.IsSetReq() {
		return ToggleNoForwardsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ToggleNoForwardsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ToggleNoForwardsResult struct {
	Success *tg.MutableChat
}

var ToggleNoForwardsResult_Success_DEFAULT *tg.MutableChat

func (p *ToggleNoForwardsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ToggleNoForwardsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ToggleNoForwardsResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ToggleNoForwardsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ToggleNoForwardsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ToggleNoForwardsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ToggleNoForwardsResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return ToggleNoForwardsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ToggleNoForwardsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *ToggleNoForwardsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ToggleNoForwardsResult) GetResult() interface{} {
	return p.Success
}

func migratedToChannelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MigratedToChannelArgs)
	realResult := result.(*MigratedToChannelResult)
	success, err := handler.(chat.RPCChat).ChatMigratedToChannel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMigratedToChannelArgs() interface{} {
	return &MigratedToChannelArgs{}
}

func newMigratedToChannelResult() interface{} {
	return &MigratedToChannelResult{}
}

type MigratedToChannelArgs struct {
	Req *chat.TLChatMigratedToChannel
}

func (p *MigratedToChannelArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MigratedToChannelArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MigratedToChannelArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatMigratedToChannel)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MigratedToChannelArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MigratedToChannelArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MigratedToChannelArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatMigratedToChannel)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MigratedToChannelArgs_Req_DEFAULT *chat.TLChatMigratedToChannel

func (p *MigratedToChannelArgs) GetReq() *chat.TLChatMigratedToChannel {
	if !p.IsSetReq() {
		return MigratedToChannelArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MigratedToChannelArgs) IsSetReq() bool {
	return p.Req != nil
}

type MigratedToChannelResult struct {
	Success *tg.Bool
}

var MigratedToChannelResult_Success_DEFAULT *tg.Bool

func (p *MigratedToChannelResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MigratedToChannelResult")
	}
	return json.Marshal(p.Success)
}

func (p *MigratedToChannelResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MigratedToChannelResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MigratedToChannelResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MigratedToChannelResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MigratedToChannelResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MigratedToChannelResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MigratedToChannelResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MigratedToChannelResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MigratedToChannelResult) GetResult() interface{} {
	return p.Success
}

func getChatParticipantIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetChatParticipantIdListArgs)
	realResult := result.(*GetChatParticipantIdListResult)
	success, err := handler.(chat.RPCChat).ChatGetChatParticipantIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetChatParticipantIdListArgs() interface{} {
	return &GetChatParticipantIdListArgs{}
}

func newGetChatParticipantIdListResult() interface{} {
	return &GetChatParticipantIdListResult{}
}

type GetChatParticipantIdListArgs struct {
	Req *chat.TLChatGetChatParticipantIdList
}

func (p *GetChatParticipantIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChatParticipantIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetChatParticipantIdListArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetChatParticipantIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetChatParticipantIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetChatParticipantIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetChatParticipantIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetChatParticipantIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetChatParticipantIdListArgs_Req_DEFAULT *chat.TLChatGetChatParticipantIdList

func (p *GetChatParticipantIdListArgs) GetReq() *chat.TLChatGetChatParticipantIdList {
	if !p.IsSetReq() {
		return GetChatParticipantIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChatParticipantIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetChatParticipantIdListResult struct {
	Success *chat.VectorLong
}

var GetChatParticipantIdListResult_Success_DEFAULT *chat.VectorLong

func (p *GetChatParticipantIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChatParticipantIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetChatParticipantIdListResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatParticipantIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetChatParticipantIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetChatParticipantIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatParticipantIdListResult) GetSuccess() *chat.VectorLong {
	if !p.IsSetSuccess() {
		return GetChatParticipantIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChatParticipantIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorLong)
}

func (p *GetChatParticipantIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetChatParticipantIdListResult) GetResult() interface{} {
	return p.Success
}

func getUsersChatIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUsersChatIdListArgs)
	realResult := result.(*GetUsersChatIdListResult)
	success, err := handler.(chat.RPCChat).ChatGetUsersChatIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUsersChatIdListArgs() interface{} {
	return &GetUsersChatIdListArgs{}
}

func newGetUsersChatIdListResult() interface{} {
	return &GetUsersChatIdListResult{}
}

type GetUsersChatIdListArgs struct {
	Req *chat.TLChatGetUsersChatIdList
}

func (p *GetUsersChatIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUsersChatIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUsersChatIdListArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetUsersChatIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUsersChatIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUsersChatIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUsersChatIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetUsersChatIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUsersChatIdListArgs_Req_DEFAULT *chat.TLChatGetUsersChatIdList

func (p *GetUsersChatIdListArgs) GetReq() *chat.TLChatGetUsersChatIdList {
	if !p.IsSetReq() {
		return GetUsersChatIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUsersChatIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUsersChatIdListResult struct {
	Success *chat.VectorUserChatIdList
}

var GetUsersChatIdListResult_Success_DEFAULT *chat.VectorUserChatIdList

func (p *GetUsersChatIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUsersChatIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUsersChatIdListResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorUserChatIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUsersChatIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUsersChatIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUsersChatIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.VectorUserChatIdList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUsersChatIdListResult) GetSuccess() *chat.VectorUserChatIdList {
	if !p.IsSetSuccess() {
		return GetUsersChatIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUsersChatIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorUserChatIdList)
}

func (p *GetUsersChatIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUsersChatIdListResult) GetResult() interface{} {
	return p.Success
}

func getMyChatListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetMyChatListArgs)
	realResult := result.(*GetMyChatListResult)
	success, err := handler.(chat.RPCChat).ChatGetMyChatList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetMyChatListArgs() interface{} {
	return &GetMyChatListArgs{}
}

func newGetMyChatListResult() interface{} {
	return &GetMyChatListResult{}
}

type GetMyChatListArgs struct {
	Req *chat.TLChatGetMyChatList
}

func (p *GetMyChatListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetMyChatListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetMyChatListArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetMyChatList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetMyChatListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetMyChatListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetMyChatListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetMyChatList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetMyChatListArgs_Req_DEFAULT *chat.TLChatGetMyChatList

func (p *GetMyChatListArgs) GetReq() *chat.TLChatGetMyChatList {
	if !p.IsSetReq() {
		return GetMyChatListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMyChatListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetMyChatListResult struct {
	Success *chat.VectorMutableChat
}

var GetMyChatListResult_Success_DEFAULT *chat.VectorMutableChat

func (p *GetMyChatListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetMyChatListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetMyChatListResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorMutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMyChatListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetMyChatListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetMyChatListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.VectorMutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMyChatListResult) GetSuccess() *chat.VectorMutableChat {
	if !p.IsSetSuccess() {
		return GetMyChatListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMyChatListResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorMutableChat)
}

func (p *GetMyChatListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMyChatListResult) GetResult() interface{} {
	return p.Success
}

func exportChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ExportChatInviteArgs)
	realResult := result.(*ExportChatInviteResult)
	success, err := handler.(chat.RPCChat).ChatExportChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newExportChatInviteArgs() interface{} {
	return &ExportChatInviteArgs{}
}

func newExportChatInviteResult() interface{} {
	return &ExportChatInviteResult{}
}

type ExportChatInviteArgs struct {
	Req *chat.TLChatExportChatInvite
}

func (p *ExportChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ExportChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ExportChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatExportChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ExportChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ExportChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ExportChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatExportChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ExportChatInviteArgs_Req_DEFAULT *chat.TLChatExportChatInvite

func (p *ExportChatInviteArgs) GetReq() *chat.TLChatExportChatInvite {
	if !p.IsSetReq() {
		return ExportChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ExportChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type ExportChatInviteResult struct {
	Success *tg.ExportedChatInvite
}

var ExportChatInviteResult_Success_DEFAULT *tg.ExportedChatInvite

func (p *ExportChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ExportChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *ExportChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.ExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ExportChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ExportChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ExportChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ExportedChatInvite)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ExportChatInviteResult) GetSuccess() *tg.ExportedChatInvite {
	if !p.IsSetSuccess() {
		return ExportChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ExportChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ExportedChatInvite)
}

func (p *ExportChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ExportChatInviteResult) GetResult() interface{} {
	return p.Success
}

func getAdminsWithInvitesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAdminsWithInvitesArgs)
	realResult := result.(*GetAdminsWithInvitesResult)
	success, err := handler.(chat.RPCChat).ChatGetAdminsWithInvites(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAdminsWithInvitesArgs() interface{} {
	return &GetAdminsWithInvitesArgs{}
}

func newGetAdminsWithInvitesResult() interface{} {
	return &GetAdminsWithInvitesResult{}
}

type GetAdminsWithInvitesArgs struct {
	Req *chat.TLChatGetAdminsWithInvites
}

func (p *GetAdminsWithInvitesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAdminsWithInvitesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAdminsWithInvitesArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetAdminsWithInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAdminsWithInvitesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAdminsWithInvitesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAdminsWithInvitesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetAdminsWithInvites)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAdminsWithInvitesArgs_Req_DEFAULT *chat.TLChatGetAdminsWithInvites

func (p *GetAdminsWithInvitesArgs) GetReq() *chat.TLChatGetAdminsWithInvites {
	if !p.IsSetReq() {
		return GetAdminsWithInvitesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAdminsWithInvitesArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAdminsWithInvitesResult struct {
	Success *chat.VectorChatAdminWithInvites
}

var GetAdminsWithInvitesResult_Success_DEFAULT *chat.VectorChatAdminWithInvites

func (p *GetAdminsWithInvitesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAdminsWithInvitesResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAdminsWithInvitesResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorChatAdminWithInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAdminsWithInvitesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAdminsWithInvitesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAdminsWithInvitesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.VectorChatAdminWithInvites)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAdminsWithInvitesResult) GetSuccess() *chat.VectorChatAdminWithInvites {
	if !p.IsSetSuccess() {
		return GetAdminsWithInvitesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAdminsWithInvitesResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorChatAdminWithInvites)
}

func (p *GetAdminsWithInvitesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAdminsWithInvitesResult) GetResult() interface{} {
	return p.Success
}

func getExportedChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetExportedChatInviteArgs)
	realResult := result.(*GetExportedChatInviteResult)
	success, err := handler.(chat.RPCChat).ChatGetExportedChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetExportedChatInviteArgs() interface{} {
	return &GetExportedChatInviteArgs{}
}

func newGetExportedChatInviteResult() interface{} {
	return &GetExportedChatInviteResult{}
}

type GetExportedChatInviteArgs struct {
	Req *chat.TLChatGetExportedChatInvite
}

func (p *GetExportedChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetExportedChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetExportedChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetExportedChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetExportedChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetExportedChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetExportedChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetExportedChatInviteArgs_Req_DEFAULT *chat.TLChatGetExportedChatInvite

func (p *GetExportedChatInviteArgs) GetReq() *chat.TLChatGetExportedChatInvite {
	if !p.IsSetReq() {
		return GetExportedChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetExportedChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetExportedChatInviteResult struct {
	Success *tg.ExportedChatInvite
}

var GetExportedChatInviteResult_Success_DEFAULT *tg.ExportedChatInvite

func (p *GetExportedChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetExportedChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetExportedChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.ExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetExportedChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetExportedChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetExportedChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ExportedChatInvite)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetExportedChatInviteResult) GetSuccess() *tg.ExportedChatInvite {
	if !p.IsSetSuccess() {
		return GetExportedChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetExportedChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ExportedChatInvite)
}

func (p *GetExportedChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetExportedChatInviteResult) GetResult() interface{} {
	return p.Success
}

func getExportedChatInvitesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetExportedChatInvitesArgs)
	realResult := result.(*GetExportedChatInvitesResult)
	success, err := handler.(chat.RPCChat).ChatGetExportedChatInvites(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetExportedChatInvitesArgs() interface{} {
	return &GetExportedChatInvitesArgs{}
}

func newGetExportedChatInvitesResult() interface{} {
	return &GetExportedChatInvitesResult{}
}

type GetExportedChatInvitesArgs struct {
	Req *chat.TLChatGetExportedChatInvites
}

func (p *GetExportedChatInvitesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetExportedChatInvitesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetExportedChatInvitesArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetExportedChatInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetExportedChatInvitesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetExportedChatInvitesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetExportedChatInvitesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetExportedChatInvites)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetExportedChatInvitesArgs_Req_DEFAULT *chat.TLChatGetExportedChatInvites

func (p *GetExportedChatInvitesArgs) GetReq() *chat.TLChatGetExportedChatInvites {
	if !p.IsSetReq() {
		return GetExportedChatInvitesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetExportedChatInvitesArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetExportedChatInvitesResult struct {
	Success *chat.VectorExportedChatInvite
}

var GetExportedChatInvitesResult_Success_DEFAULT *chat.VectorExportedChatInvite

func (p *GetExportedChatInvitesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetExportedChatInvitesResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetExportedChatInvitesResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetExportedChatInvitesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetExportedChatInvitesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetExportedChatInvitesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.VectorExportedChatInvite)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetExportedChatInvitesResult) GetSuccess() *chat.VectorExportedChatInvite {
	if !p.IsSetSuccess() {
		return GetExportedChatInvitesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetExportedChatInvitesResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorExportedChatInvite)
}

func (p *GetExportedChatInvitesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetExportedChatInvitesResult) GetResult() interface{} {
	return p.Success
}

func checkChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CheckChatInviteArgs)
	realResult := result.(*CheckChatInviteResult)
	success, err := handler.(chat.RPCChat).ChatCheckChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCheckChatInviteArgs() interface{} {
	return &CheckChatInviteArgs{}
}

func newCheckChatInviteResult() interface{} {
	return &CheckChatInviteResult{}
}

type CheckChatInviteArgs struct {
	Req *chat.TLChatCheckChatInvite
}

func (p *CheckChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CheckChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatCheckChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CheckChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CheckChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CheckChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatCheckChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CheckChatInviteArgs_Req_DEFAULT *chat.TLChatCheckChatInvite

func (p *CheckChatInviteArgs) GetReq() *chat.TLChatCheckChatInvite {
	if !p.IsSetReq() {
		return CheckChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckChatInviteResult struct {
	Success *chat.ChatInviteExt
}

var CheckChatInviteResult_Success_DEFAULT *chat.ChatInviteExt

func (p *CheckChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *CheckChatInviteResult) Unmarshal(in []byte) error {
	msg := new(chat.ChatInviteExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CheckChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CheckChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.ChatInviteExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckChatInviteResult) GetSuccess() *chat.ChatInviteExt {
	if !p.IsSetSuccess() {
		return CheckChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.ChatInviteExt)
}

func (p *CheckChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckChatInviteResult) GetResult() interface{} {
	return p.Success
}

func importChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ImportChatInviteArgs)
	realResult := result.(*ImportChatInviteResult)
	success, err := handler.(chat.RPCChat).ChatImportChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newImportChatInviteArgs() interface{} {
	return &ImportChatInviteArgs{}
}

func newImportChatInviteResult() interface{} {
	return &ImportChatInviteResult{}
}

type ImportChatInviteArgs struct {
	Req *chat.TLChatImportChatInvite
}

func (p *ImportChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ImportChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ImportChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatImportChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ImportChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ImportChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ImportChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatImportChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ImportChatInviteArgs_Req_DEFAULT *chat.TLChatImportChatInvite

func (p *ImportChatInviteArgs) GetReq() *chat.TLChatImportChatInvite {
	if !p.IsSetReq() {
		return ImportChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ImportChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type ImportChatInviteResult struct {
	Success *tg.MutableChat
}

var ImportChatInviteResult_Success_DEFAULT *tg.MutableChat

func (p *ImportChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ImportChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *ImportChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ImportChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ImportChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ImportChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ImportChatInviteResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return ImportChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ImportChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *ImportChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ImportChatInviteResult) GetResult() interface{} {
	return p.Success
}

func getChatInviteImportersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetChatInviteImportersArgs)
	realResult := result.(*GetChatInviteImportersResult)
	success, err := handler.(chat.RPCChat).ChatGetChatInviteImporters(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetChatInviteImportersArgs() interface{} {
	return &GetChatInviteImportersArgs{}
}

func newGetChatInviteImportersResult() interface{} {
	return &GetChatInviteImportersResult{}
}

type GetChatInviteImportersArgs struct {
	Req *chat.TLChatGetChatInviteImporters
}

func (p *GetChatInviteImportersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChatInviteImportersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetChatInviteImportersArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetChatInviteImporters)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetChatInviteImportersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetChatInviteImportersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetChatInviteImportersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetChatInviteImporters)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetChatInviteImportersArgs_Req_DEFAULT *chat.TLChatGetChatInviteImporters

func (p *GetChatInviteImportersArgs) GetReq() *chat.TLChatGetChatInviteImporters {
	if !p.IsSetReq() {
		return GetChatInviteImportersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChatInviteImportersArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetChatInviteImportersResult struct {
	Success *chat.VectorChatInviteImporter
}

var GetChatInviteImportersResult_Success_DEFAULT *chat.VectorChatInviteImporter

func (p *GetChatInviteImportersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChatInviteImportersResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetChatInviteImportersResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorChatInviteImporter)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatInviteImportersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetChatInviteImportersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetChatInviteImportersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.VectorChatInviteImporter)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatInviteImportersResult) GetSuccess() *chat.VectorChatInviteImporter {
	if !p.IsSetSuccess() {
		return GetChatInviteImportersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChatInviteImportersResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorChatInviteImporter)
}

func (p *GetChatInviteImportersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetChatInviteImportersResult) GetResult() interface{} {
	return p.Success
}

func deleteExportedChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteExportedChatInviteArgs)
	realResult := result.(*DeleteExportedChatInviteResult)
	success, err := handler.(chat.RPCChat).ChatDeleteExportedChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteExportedChatInviteArgs() interface{} {
	return &DeleteExportedChatInviteArgs{}
}

func newDeleteExportedChatInviteResult() interface{} {
	return &DeleteExportedChatInviteResult{}
}

type DeleteExportedChatInviteArgs struct {
	Req *chat.TLChatDeleteExportedChatInvite
}

func (p *DeleteExportedChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteExportedChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteExportedChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatDeleteExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteExportedChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteExportedChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteExportedChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatDeleteExportedChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteExportedChatInviteArgs_Req_DEFAULT *chat.TLChatDeleteExportedChatInvite

func (p *DeleteExportedChatInviteArgs) GetReq() *chat.TLChatDeleteExportedChatInvite {
	if !p.IsSetReq() {
		return DeleteExportedChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteExportedChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteExportedChatInviteResult struct {
	Success *tg.Bool
}

var DeleteExportedChatInviteResult_Success_DEFAULT *tg.Bool

func (p *DeleteExportedChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteExportedChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteExportedChatInviteResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteExportedChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteExportedChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteExportedChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteExportedChatInviteResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteExportedChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteExportedChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteExportedChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteExportedChatInviteResult) GetResult() interface{} {
	return p.Success
}

func deleteRevokedExportedChatInvitesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteRevokedExportedChatInvitesArgs)
	realResult := result.(*DeleteRevokedExportedChatInvitesResult)
	success, err := handler.(chat.RPCChat).ChatDeleteRevokedExportedChatInvites(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteRevokedExportedChatInvitesArgs() interface{} {
	return &DeleteRevokedExportedChatInvitesArgs{}
}

func newDeleteRevokedExportedChatInvitesResult() interface{} {
	return &DeleteRevokedExportedChatInvitesResult{}
}

type DeleteRevokedExportedChatInvitesArgs struct {
	Req *chat.TLChatDeleteRevokedExportedChatInvites
}

func (p *DeleteRevokedExportedChatInvitesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteRevokedExportedChatInvitesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteRevokedExportedChatInvitesArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatDeleteRevokedExportedChatInvites)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteRevokedExportedChatInvitesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteRevokedExportedChatInvitesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteRevokedExportedChatInvitesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatDeleteRevokedExportedChatInvites)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteRevokedExportedChatInvitesArgs_Req_DEFAULT *chat.TLChatDeleteRevokedExportedChatInvites

func (p *DeleteRevokedExportedChatInvitesArgs) GetReq() *chat.TLChatDeleteRevokedExportedChatInvites {
	if !p.IsSetReq() {
		return DeleteRevokedExportedChatInvitesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteRevokedExportedChatInvitesArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteRevokedExportedChatInvitesResult struct {
	Success *tg.Bool
}

var DeleteRevokedExportedChatInvitesResult_Success_DEFAULT *tg.Bool

func (p *DeleteRevokedExportedChatInvitesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteRevokedExportedChatInvitesResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteRevokedExportedChatInvitesResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteRevokedExportedChatInvitesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteRevokedExportedChatInvitesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteRevokedExportedChatInvitesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteRevokedExportedChatInvitesResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteRevokedExportedChatInvitesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteRevokedExportedChatInvitesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteRevokedExportedChatInvitesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteRevokedExportedChatInvitesResult) GetResult() interface{} {
	return p.Success
}

func editExportedChatInviteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditExportedChatInviteArgs)
	realResult := result.(*EditExportedChatInviteResult)
	success, err := handler.(chat.RPCChat).ChatEditExportedChatInvite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditExportedChatInviteArgs() interface{} {
	return &EditExportedChatInviteArgs{}
}

func newEditExportedChatInviteResult() interface{} {
	return &EditExportedChatInviteResult{}
}

type EditExportedChatInviteArgs struct {
	Req *chat.TLChatEditExportedChatInvite
}

func (p *EditExportedChatInviteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditExportedChatInviteArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditExportedChatInviteArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatEditExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditExportedChatInviteArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditExportedChatInviteArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditExportedChatInviteArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatEditExportedChatInvite)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditExportedChatInviteArgs_Req_DEFAULT *chat.TLChatEditExportedChatInvite

func (p *EditExportedChatInviteArgs) GetReq() *chat.TLChatEditExportedChatInvite {
	if !p.IsSetReq() {
		return EditExportedChatInviteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditExportedChatInviteArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditExportedChatInviteResult struct {
	Success *chat.VectorExportedChatInvite
}

var EditExportedChatInviteResult_Success_DEFAULT *chat.VectorExportedChatInvite

func (p *EditExportedChatInviteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditExportedChatInviteResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditExportedChatInviteResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorExportedChatInvite)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditExportedChatInviteResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditExportedChatInviteResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditExportedChatInviteResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.VectorExportedChatInvite)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditExportedChatInviteResult) GetSuccess() *chat.VectorExportedChatInvite {
	if !p.IsSetSuccess() {
		return EditExportedChatInviteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditExportedChatInviteResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorExportedChatInvite)
}

func (p *EditExportedChatInviteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditExportedChatInviteResult) GetResult() interface{} {
	return p.Success
}

func setChatAvailableReactionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetChatAvailableReactionsArgs)
	realResult := result.(*SetChatAvailableReactionsResult)
	success, err := handler.(chat.RPCChat).ChatSetChatAvailableReactions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetChatAvailableReactionsArgs() interface{} {
	return &SetChatAvailableReactionsArgs{}
}

func newSetChatAvailableReactionsResult() interface{} {
	return &SetChatAvailableReactionsResult{}
}

type SetChatAvailableReactionsArgs struct {
	Req *chat.TLChatSetChatAvailableReactions
}

func (p *SetChatAvailableReactionsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetChatAvailableReactionsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetChatAvailableReactionsArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatSetChatAvailableReactions)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetChatAvailableReactionsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetChatAvailableReactionsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetChatAvailableReactionsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatSetChatAvailableReactions)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetChatAvailableReactionsArgs_Req_DEFAULT *chat.TLChatSetChatAvailableReactions

func (p *SetChatAvailableReactionsArgs) GetReq() *chat.TLChatSetChatAvailableReactions {
	if !p.IsSetReq() {
		return SetChatAvailableReactionsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetChatAvailableReactionsArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetChatAvailableReactionsResult struct {
	Success *tg.MutableChat
}

var SetChatAvailableReactionsResult_Success_DEFAULT *tg.MutableChat

func (p *SetChatAvailableReactionsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetChatAvailableReactionsResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetChatAvailableReactionsResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChatAvailableReactionsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetChatAvailableReactionsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetChatAvailableReactionsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChatAvailableReactionsResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return SetChatAvailableReactionsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetChatAvailableReactionsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *SetChatAvailableReactionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetChatAvailableReactionsResult) GetResult() interface{} {
	return p.Success
}

func setHistoryTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetHistoryTTLArgs)
	realResult := result.(*SetHistoryTTLResult)
	success, err := handler.(chat.RPCChat).ChatSetHistoryTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetHistoryTTLArgs() interface{} {
	return &SetHistoryTTLArgs{}
}

func newSetHistoryTTLResult() interface{} {
	return &SetHistoryTTLResult{}
}

type SetHistoryTTLArgs struct {
	Req *chat.TLChatSetHistoryTTL
}

func (p *SetHistoryTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetHistoryTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetHistoryTTLArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatSetHistoryTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetHistoryTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetHistoryTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetHistoryTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatSetHistoryTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetHistoryTTLArgs_Req_DEFAULT *chat.TLChatSetHistoryTTL

func (p *SetHistoryTTLArgs) GetReq() *chat.TLChatSetHistoryTTL {
	if !p.IsSetReq() {
		return SetHistoryTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetHistoryTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetHistoryTTLResult struct {
	Success *tg.MutableChat
}

var SetHistoryTTLResult_Success_DEFAULT *tg.MutableChat

func (p *SetHistoryTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetHistoryTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetHistoryTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.MutableChat)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetHistoryTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetHistoryTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetHistoryTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetHistoryTTLResult) GetSuccess() *tg.MutableChat {
	if !p.IsSetSuccess() {
		return SetHistoryTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetHistoryTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableChat)
}

func (p *SetHistoryTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetHistoryTTLResult) GetResult() interface{} {
	return p.Success
}

func searchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SearchArgs)
	realResult := result.(*SearchResult)
	success, err := handler.(chat.RPCChat).ChatSearch(ctx, realArg.Req)
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
	Req *chat.TLChatSearch
}

func (p *SearchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SearchArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SearchArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatSearch)
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
	msg := new(chat.TLChatSearch)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SearchArgs_Req_DEFAULT *chat.TLChatSearch

func (p *SearchArgs) GetReq() *chat.TLChatSearch {
	if !p.IsSetReq() {
		return SearchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchArgs) IsSetReq() bool {
	return p.Req != nil
}

type SearchResult struct {
	Success *chat.VectorMutableChat
}

var SearchResult_Success_DEFAULT *chat.VectorMutableChat

func (p *SearchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SearchResult")
	}
	return json.Marshal(p.Success)
}

func (p *SearchResult) Unmarshal(in []byte) error {
	msg := new(chat.VectorMutableChat)
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
	msg := new(chat.VectorMutableChat)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchResult) GetSuccess() *chat.VectorMutableChat {
	if !p.IsSetSuccess() {
		return SearchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.VectorMutableChat)
}

func (p *SearchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchResult) GetResult() interface{} {
	return p.Success
}

func getRecentChatInviteRequestersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetRecentChatInviteRequestersArgs)
	realResult := result.(*GetRecentChatInviteRequestersResult)
	success, err := handler.(chat.RPCChat).ChatGetRecentChatInviteRequesters(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetRecentChatInviteRequestersArgs() interface{} {
	return &GetRecentChatInviteRequestersArgs{}
}

func newGetRecentChatInviteRequestersResult() interface{} {
	return &GetRecentChatInviteRequestersResult{}
}

type GetRecentChatInviteRequestersArgs struct {
	Req *chat.TLChatGetRecentChatInviteRequesters
}

func (p *GetRecentChatInviteRequestersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetRecentChatInviteRequestersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetRecentChatInviteRequestersArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatGetRecentChatInviteRequesters)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetRecentChatInviteRequestersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetRecentChatInviteRequestersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetRecentChatInviteRequestersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatGetRecentChatInviteRequesters)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetRecentChatInviteRequestersArgs_Req_DEFAULT *chat.TLChatGetRecentChatInviteRequesters

func (p *GetRecentChatInviteRequestersArgs) GetReq() *chat.TLChatGetRecentChatInviteRequesters {
	if !p.IsSetReq() {
		return GetRecentChatInviteRequestersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetRecentChatInviteRequestersArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetRecentChatInviteRequestersResult struct {
	Success *chat.RecentChatInviteRequesters
}

var GetRecentChatInviteRequestersResult_Success_DEFAULT *chat.RecentChatInviteRequesters

func (p *GetRecentChatInviteRequestersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetRecentChatInviteRequestersResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetRecentChatInviteRequestersResult) Unmarshal(in []byte) error {
	msg := new(chat.RecentChatInviteRequesters)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetRecentChatInviteRequestersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetRecentChatInviteRequestersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetRecentChatInviteRequestersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.RecentChatInviteRequesters)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetRecentChatInviteRequestersResult) GetSuccess() *chat.RecentChatInviteRequesters {
	if !p.IsSetSuccess() {
		return GetRecentChatInviteRequestersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetRecentChatInviteRequestersResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.RecentChatInviteRequesters)
}

func (p *GetRecentChatInviteRequestersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetRecentChatInviteRequestersResult) GetResult() interface{} {
	return p.Success
}

func hideChatJoinRequestsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HideChatJoinRequestsArgs)
	realResult := result.(*HideChatJoinRequestsResult)
	success, err := handler.(chat.RPCChat).ChatHideChatJoinRequests(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHideChatJoinRequestsArgs() interface{} {
	return &HideChatJoinRequestsArgs{}
}

func newHideChatJoinRequestsResult() interface{} {
	return &HideChatJoinRequestsResult{}
}

type HideChatJoinRequestsArgs struct {
	Req *chat.TLChatHideChatJoinRequests
}

func (p *HideChatJoinRequestsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HideChatJoinRequestsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HideChatJoinRequestsArgs) Unmarshal(in []byte) error {
	msg := new(chat.TLChatHideChatJoinRequests)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HideChatJoinRequestsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HideChatJoinRequestsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HideChatJoinRequestsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatHideChatJoinRequests)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HideChatJoinRequestsArgs_Req_DEFAULT *chat.TLChatHideChatJoinRequests

func (p *HideChatJoinRequestsArgs) GetReq() *chat.TLChatHideChatJoinRequests {
	if !p.IsSetReq() {
		return HideChatJoinRequestsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HideChatJoinRequestsArgs) IsSetReq() bool {
	return p.Req != nil
}

type HideChatJoinRequestsResult struct {
	Success *chat.RecentChatInviteRequesters
}

var HideChatJoinRequestsResult_Success_DEFAULT *chat.RecentChatInviteRequesters

func (p *HideChatJoinRequestsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HideChatJoinRequestsResult")
	}
	return json.Marshal(p.Success)
}

func (p *HideChatJoinRequestsResult) Unmarshal(in []byte) error {
	msg := new(chat.RecentChatInviteRequesters)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HideChatJoinRequestsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HideChatJoinRequestsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HideChatJoinRequestsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.RecentChatInviteRequesters)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HideChatJoinRequestsResult) GetSuccess() *chat.RecentChatInviteRequesters {
	if !p.IsSetSuccess() {
		return HideChatJoinRequestsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HideChatJoinRequestsResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.RecentChatInviteRequesters)
}

func (p *HideChatJoinRequestsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HideChatJoinRequestsResult) GetResult() interface{} {
	return p.Success
}

func importChatInvite2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ImportChatInvite2Args)
	realResult := result.(*ImportChatInvite2Result)
	success, err := handler.(chat.RPCChat).ChatImportChatInvite2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newImportChatInvite2Args() interface{} {
	return &ImportChatInvite2Args{}
}

func newImportChatInvite2Result() interface{} {
	return &ImportChatInvite2Result{}
}

type ImportChatInvite2Args struct {
	Req *chat.TLChatImportChatInvite2
}

func (p *ImportChatInvite2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ImportChatInvite2Args")
	}
	return json.Marshal(p.Req)
}

func (p *ImportChatInvite2Args) Unmarshal(in []byte) error {
	msg := new(chat.TLChatImportChatInvite2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ImportChatInvite2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ImportChatInvite2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *ImportChatInvite2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.TLChatImportChatInvite2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ImportChatInvite2Args_Req_DEFAULT *chat.TLChatImportChatInvite2

func (p *ImportChatInvite2Args) GetReq() *chat.TLChatImportChatInvite2 {
	if !p.IsSetReq() {
		return ImportChatInvite2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *ImportChatInvite2Args) IsSetReq() bool {
	return p.Req != nil
}

type ImportChatInvite2Result struct {
	Success *chat.ChatInviteImported
}

var ImportChatInvite2Result_Success_DEFAULT *chat.ChatInviteImported

func (p *ImportChatInvite2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ImportChatInvite2Result")
	}
	return json.Marshal(p.Success)
}

func (p *ImportChatInvite2Result) Unmarshal(in []byte) error {
	msg := new(chat.ChatInviteImported)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ImportChatInvite2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ImportChatInvite2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *ImportChatInvite2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(chat.ChatInviteImported)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ImportChatInvite2Result) GetSuccess() *chat.ChatInviteImported {
	if !p.IsSetSuccess() {
		return ImportChatInvite2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *ImportChatInvite2Result) SetSuccess(x interface{}) {
	p.Success = x.(*chat.ChatInviteImported)
}

func (p *ImportChatInvite2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ImportChatInvite2Result) GetResult() interface{} {
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

func (p *kClient) ChatGetMutableChat(ctx context.Context, req *chat.TLChatGetMutableChat) (r *tg.MutableChat, err error) {
	// var _args GetMutableChatArgs
	// _args.Req = req
	// var _result GetMutableChatResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.getMutableChat", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetChatListByIdList(ctx context.Context, req *chat.TLChatGetChatListByIdList) (r *chat.VectorMutableChat, err error) {
	// var _args GetChatListByIdListArgs
	// _args.Req = req
	// var _result GetChatListByIdListResult

	_result := new(chat.VectorMutableChat)

	if err = p.c.Call(ctx, "chat.getChatListByIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetChatBySelfId(ctx context.Context, req *chat.TLChatGetChatBySelfId) (r *tg.MutableChat, err error) {
	// var _args GetChatBySelfIdArgs
	// _args.Req = req
	// var _result GetChatBySelfIdResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.getChatBySelfId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatCreateChat2(ctx context.Context, req *chat.TLChatCreateChat2) (r *tg.MutableChat, err error) {
	// var _args CreateChat2Args
	// _args.Req = req
	// var _result CreateChat2Result

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.createChat2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatDeleteChat(ctx context.Context, req *chat.TLChatDeleteChat) (r *tg.MutableChat, err error) {
	// var _args DeleteChatArgs
	// _args.Req = req
	// var _result DeleteChatResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.deleteChat", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatDeleteChatUser(ctx context.Context, req *chat.TLChatDeleteChatUser) (r *tg.MutableChat, err error) {
	// var _args DeleteChatUserArgs
	// _args.Req = req
	// var _result DeleteChatUserResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.deleteChatUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatEditChatTitle(ctx context.Context, req *chat.TLChatEditChatTitle) (r *tg.MutableChat, err error) {
	// var _args EditChatTitleArgs
	// _args.Req = req
	// var _result EditChatTitleResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.editChatTitle", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatEditChatAbout(ctx context.Context, req *chat.TLChatEditChatAbout) (r *tg.MutableChat, err error) {
	// var _args EditChatAboutArgs
	// _args.Req = req
	// var _result EditChatAboutResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.editChatAbout", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatEditChatPhoto(ctx context.Context, req *chat.TLChatEditChatPhoto) (r *tg.MutableChat, err error) {
	// var _args EditChatPhotoArgs
	// _args.Req = req
	// var _result EditChatPhotoResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.editChatPhoto", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatEditChatAdmin(ctx context.Context, req *chat.TLChatEditChatAdmin) (r *tg.MutableChat, err error) {
	// var _args EditChatAdminArgs
	// _args.Req = req
	// var _result EditChatAdminResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.editChatAdmin", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatEditChatDefaultBannedRights(ctx context.Context, req *chat.TLChatEditChatDefaultBannedRights) (r *tg.MutableChat, err error) {
	// var _args EditChatDefaultBannedRightsArgs
	// _args.Req = req
	// var _result EditChatDefaultBannedRightsResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.editChatDefaultBannedRights", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatAddChatUser(ctx context.Context, req *chat.TLChatAddChatUser) (r *tg.MutableChat, err error) {
	// var _args AddChatUserArgs
	// _args.Req = req
	// var _result AddChatUserResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.addChatUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetMutableChatByLink(ctx context.Context, req *chat.TLChatGetMutableChatByLink) (r *tg.MutableChat, err error) {
	// var _args GetMutableChatByLinkArgs
	// _args.Req = req
	// var _result GetMutableChatByLinkResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.getMutableChatByLink", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatToggleNoForwards(ctx context.Context, req *chat.TLChatToggleNoForwards) (r *tg.MutableChat, err error) {
	// var _args ToggleNoForwardsArgs
	// _args.Req = req
	// var _result ToggleNoForwardsResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.toggleNoForwards", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatMigratedToChannel(ctx context.Context, req *chat.TLChatMigratedToChannel) (r *tg.Bool, err error) {
	// var _args MigratedToChannelArgs
	// _args.Req = req
	// var _result MigratedToChannelResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "chat.migratedToChannel", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetChatParticipantIdList(ctx context.Context, req *chat.TLChatGetChatParticipantIdList) (r *chat.VectorLong, err error) {
	// var _args GetChatParticipantIdListArgs
	// _args.Req = req
	// var _result GetChatParticipantIdListResult

	_result := new(chat.VectorLong)

	if err = p.c.Call(ctx, "chat.getChatParticipantIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetUsersChatIdList(ctx context.Context, req *chat.TLChatGetUsersChatIdList) (r *chat.VectorUserChatIdList, err error) {
	// var _args GetUsersChatIdListArgs
	// _args.Req = req
	// var _result GetUsersChatIdListResult

	_result := new(chat.VectorUserChatIdList)

	if err = p.c.Call(ctx, "chat.getUsersChatIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetMyChatList(ctx context.Context, req *chat.TLChatGetMyChatList) (r *chat.VectorMutableChat, err error) {
	// var _args GetMyChatListArgs
	// _args.Req = req
	// var _result GetMyChatListResult

	_result := new(chat.VectorMutableChat)

	if err = p.c.Call(ctx, "chat.getMyChatList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatExportChatInvite(ctx context.Context, req *chat.TLChatExportChatInvite) (r *tg.ExportedChatInvite, err error) {
	// var _args ExportChatInviteArgs
	// _args.Req = req
	// var _result ExportChatInviteResult

	_result := new(tg.ExportedChatInvite)

	if err = p.c.Call(ctx, "chat.exportChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetAdminsWithInvites(ctx context.Context, req *chat.TLChatGetAdminsWithInvites) (r *chat.VectorChatAdminWithInvites, err error) {
	// var _args GetAdminsWithInvitesArgs
	// _args.Req = req
	// var _result GetAdminsWithInvitesResult

	_result := new(chat.VectorChatAdminWithInvites)

	if err = p.c.Call(ctx, "chat.getAdminsWithInvites", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetExportedChatInvite(ctx context.Context, req *chat.TLChatGetExportedChatInvite) (r *tg.ExportedChatInvite, err error) {
	// var _args GetExportedChatInviteArgs
	// _args.Req = req
	// var _result GetExportedChatInviteResult

	_result := new(tg.ExportedChatInvite)

	if err = p.c.Call(ctx, "chat.getExportedChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetExportedChatInvites(ctx context.Context, req *chat.TLChatGetExportedChatInvites) (r *chat.VectorExportedChatInvite, err error) {
	// var _args GetExportedChatInvitesArgs
	// _args.Req = req
	// var _result GetExportedChatInvitesResult

	_result := new(chat.VectorExportedChatInvite)

	if err = p.c.Call(ctx, "chat.getExportedChatInvites", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatCheckChatInvite(ctx context.Context, req *chat.TLChatCheckChatInvite) (r *chat.ChatInviteExt, err error) {
	// var _args CheckChatInviteArgs
	// _args.Req = req
	// var _result CheckChatInviteResult

	_result := new(chat.ChatInviteExt)

	if err = p.c.Call(ctx, "chat.checkChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatImportChatInvite(ctx context.Context, req *chat.TLChatImportChatInvite) (r *tg.MutableChat, err error) {
	// var _args ImportChatInviteArgs
	// _args.Req = req
	// var _result ImportChatInviteResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.importChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetChatInviteImporters(ctx context.Context, req *chat.TLChatGetChatInviteImporters) (r *chat.VectorChatInviteImporter, err error) {
	// var _args GetChatInviteImportersArgs
	// _args.Req = req
	// var _result GetChatInviteImportersResult

	_result := new(chat.VectorChatInviteImporter)

	if err = p.c.Call(ctx, "chat.getChatInviteImporters", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatDeleteExportedChatInvite(ctx context.Context, req *chat.TLChatDeleteExportedChatInvite) (r *tg.Bool, err error) {
	// var _args DeleteExportedChatInviteArgs
	// _args.Req = req
	// var _result DeleteExportedChatInviteResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "chat.deleteExportedChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatDeleteRevokedExportedChatInvites(ctx context.Context, req *chat.TLChatDeleteRevokedExportedChatInvites) (r *tg.Bool, err error) {
	// var _args DeleteRevokedExportedChatInvitesArgs
	// _args.Req = req
	// var _result DeleteRevokedExportedChatInvitesResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "chat.deleteRevokedExportedChatInvites", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatEditExportedChatInvite(ctx context.Context, req *chat.TLChatEditExportedChatInvite) (r *chat.VectorExportedChatInvite, err error) {
	// var _args EditExportedChatInviteArgs
	// _args.Req = req
	// var _result EditExportedChatInviteResult

	_result := new(chat.VectorExportedChatInvite)

	if err = p.c.Call(ctx, "chat.editExportedChatInvite", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatSetChatAvailableReactions(ctx context.Context, req *chat.TLChatSetChatAvailableReactions) (r *tg.MutableChat, err error) {
	// var _args SetChatAvailableReactionsArgs
	// _args.Req = req
	// var _result SetChatAvailableReactionsResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.setChatAvailableReactions", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatSetHistoryTTL(ctx context.Context, req *chat.TLChatSetHistoryTTL) (r *tg.MutableChat, err error) {
	// var _args SetHistoryTTLArgs
	// _args.Req = req
	// var _result SetHistoryTTLResult

	_result := new(tg.MutableChat)

	if err = p.c.Call(ctx, "chat.setHistoryTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatSearch(ctx context.Context, req *chat.TLChatSearch) (r *chat.VectorMutableChat, err error) {
	// var _args SearchArgs
	// _args.Req = req
	// var _result SearchResult

	_result := new(chat.VectorMutableChat)

	if err = p.c.Call(ctx, "chat.search", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatGetRecentChatInviteRequesters(ctx context.Context, req *chat.TLChatGetRecentChatInviteRequesters) (r *chat.RecentChatInviteRequesters, err error) {
	// var _args GetRecentChatInviteRequestersArgs
	// _args.Req = req
	// var _result GetRecentChatInviteRequestersResult

	_result := new(chat.RecentChatInviteRequesters)

	if err = p.c.Call(ctx, "chat.getRecentChatInviteRequesters", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatHideChatJoinRequests(ctx context.Context, req *chat.TLChatHideChatJoinRequests) (r *chat.RecentChatInviteRequesters, err error) {
	// var _args HideChatJoinRequestsArgs
	// _args.Req = req
	// var _result HideChatJoinRequestsResult

	_result := new(chat.RecentChatInviteRequesters)

	if err = p.c.Call(ctx, "chat.hideChatJoinRequests", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) ChatImportChatInvite2(ctx context.Context, req *chat.TLChatImportChatInvite2) (r *chat.ChatInviteImported, err error) {
	// var _args ImportChatInvite2Args
	// _args.Req = req
	// var _result ImportChatInvite2Result

	_result := new(chat.ChatInviteImported)

	if err = p.c.Call(ctx, "chat.importChatInvite2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
