/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialogservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"dialog.saveDraftMessage": kitex.NewMethodInfo(
		saveDraftMessageHandler,
		newSaveDraftMessageArgs,
		newSaveDraftMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.clearDraftMessage": kitex.NewMethodInfo(
		clearDraftMessageHandler,
		newClearDraftMessageArgs,
		newClearDraftMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getAllDrafts": kitex.NewMethodInfo(
		getAllDraftsHandler,
		newGetAllDraftsArgs,
		newGetAllDraftsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.clearAllDrafts": kitex.NewMethodInfo(
		clearAllDraftsHandler,
		newClearAllDraftsArgs,
		newClearAllDraftsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.markDialogUnread": kitex.NewMethodInfo(
		markDialogUnreadHandler,
		newMarkDialogUnreadArgs,
		newMarkDialogUnreadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.toggleDialogPin": kitex.NewMethodInfo(
		toggleDialogPinHandler,
		newToggleDialogPinArgs,
		newToggleDialogPinResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogUnreadMarkList": kitex.NewMethodInfo(
		getDialogUnreadMarkListHandler,
		newGetDialogUnreadMarkListArgs,
		newGetDialogUnreadMarkListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogsByOffsetDate": kitex.NewMethodInfo(
		getDialogsByOffsetDateHandler,
		newGetDialogsByOffsetDateArgs,
		newGetDialogsByOffsetDateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogs": kitex.NewMethodInfo(
		getDialogsHandler,
		newGetDialogsArgs,
		newGetDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogsByIdList": kitex.NewMethodInfo(
		getDialogsByIdListHandler,
		newGetDialogsByIdListArgs,
		newGetDialogsByIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogsCount": kitex.NewMethodInfo(
		getDialogsCountHandler,
		newGetDialogsCountArgs,
		newGetDialogsCountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getPinnedDialogs": kitex.NewMethodInfo(
		getPinnedDialogsHandler,
		newGetPinnedDialogsArgs,
		newGetPinnedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.reorderPinnedDialogs": kitex.NewMethodInfo(
		reorderPinnedDialogsHandler,
		newReorderPinnedDialogsArgs,
		newReorderPinnedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogById": kitex.NewMethodInfo(
		getDialogByIdHandler,
		newGetDialogByIdArgs,
		newGetDialogByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getTopMessage": kitex.NewMethodInfo(
		getTopMessageHandler,
		newGetTopMessageArgs,
		newGetTopMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.insertOrUpdateDialog": kitex.NewMethodInfo(
		insertOrUpdateDialogHandler,
		newInsertOrUpdateDialogArgs,
		newInsertOrUpdateDialogResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.deleteDialog": kitex.NewMethodInfo(
		deleteDialogHandler,
		newDeleteDialogArgs,
		newDeleteDialogResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getUserPinnedMessage": kitex.NewMethodInfo(
		getUserPinnedMessageHandler,
		newGetUserPinnedMessageArgs,
		newGetUserPinnedMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.updateUserPinnedMessage": kitex.NewMethodInfo(
		updateUserPinnedMessageHandler,
		newUpdateUserPinnedMessageArgs,
		newUpdateUserPinnedMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.insertOrUpdateDialogFilter": kitex.NewMethodInfo(
		insertOrUpdateDialogFilterHandler,
		newInsertOrUpdateDialogFilterArgs,
		newInsertOrUpdateDialogFilterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.deleteDialogFilter": kitex.NewMethodInfo(
		deleteDialogFilterHandler,
		newDeleteDialogFilterArgs,
		newDeleteDialogFilterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.updateDialogFiltersOrder": kitex.NewMethodInfo(
		updateDialogFiltersOrderHandler,
		newUpdateDialogFiltersOrderArgs,
		newUpdateDialogFiltersOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogFilters": kitex.NewMethodInfo(
		getDialogFiltersHandler,
		newGetDialogFiltersArgs,
		newGetDialogFiltersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogFolder": kitex.NewMethodInfo(
		getDialogFolderHandler,
		newGetDialogFolderArgs,
		newGetDialogFolderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.editPeerFolders": kitex.NewMethodInfo(
		editPeerFoldersHandler,
		newEditPeerFoldersArgs,
		newEditPeerFoldersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getChannelMessageReadParticipants": kitex.NewMethodInfo(
		getChannelMessageReadParticipantsHandler,
		newGetChannelMessageReadParticipantsArgs,
		newGetChannelMessageReadParticipantsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.setChatTheme": kitex.NewMethodInfo(
		setChatThemeHandler,
		newSetChatThemeArgs,
		newSetChatThemeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.setHistoryTTL": kitex.NewMethodInfo(
		setHistoryTTLHandler,
		newSetHistoryTTLArgs,
		newSetHistoryTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getMyDialogsData": kitex.NewMethodInfo(
		getMyDialogsDataHandler,
		newGetMyDialogsDataArgs,
		newGetMyDialogsDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getSavedDialogs": kitex.NewMethodInfo(
		getSavedDialogsHandler,
		newGetSavedDialogsArgs,
		newGetSavedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getPinnedSavedDialogs": kitex.NewMethodInfo(
		getPinnedSavedDialogsHandler,
		newGetPinnedSavedDialogsArgs,
		newGetPinnedSavedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.toggleSavedDialogPin": kitex.NewMethodInfo(
		toggleSavedDialogPinHandler,
		newToggleSavedDialogPinArgs,
		newToggleSavedDialogPinResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.reorderPinnedSavedDialogs": kitex.NewMethodInfo(
		reorderPinnedSavedDialogsHandler,
		newReorderPinnedSavedDialogsArgs,
		newReorderPinnedSavedDialogsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogFilter": kitex.NewMethodInfo(
		getDialogFilterHandler,
		newGetDialogFilterArgs,
		newGetDialogFilterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogFilterBySlug": kitex.NewMethodInfo(
		getDialogFilterBySlugHandler,
		newGetDialogFilterBySlugArgs,
		newGetDialogFilterBySlugResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.createDialogFilter": kitex.NewMethodInfo(
		createDialogFilterHandler,
		newCreateDialogFilterArgs,
		newCreateDialogFilterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.updateUnreadCount": kitex.NewMethodInfo(
		updateUnreadCountHandler,
		newUpdateUnreadCountArgs,
		newUpdateUnreadCountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.toggleDialogFilterTags": kitex.NewMethodInfo(
		toggleDialogFilterTagsHandler,
		newToggleDialogFilterTagsArgs,
		newToggleDialogFilterTagsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.getDialogFilterTags": kitex.NewMethodInfo(
		getDialogFilterTagsHandler,
		newGetDialogFilterTagsArgs,
		newGetDialogFilterTagsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"dialog.setChatWallpaper": kitex.NewMethodInfo(
		setChatWallpaperHandler,
		newSetChatWallpaperArgs,
		newSetChatWallpaperResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	dialogServiceServiceInfo                = NewServiceInfo()
	dialogServiceServiceInfoForClient       = NewServiceInfoForClient()
	dialogServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCDialog", dialogServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCDialog", dialogServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCDialog", dialogServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return dialogServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return dialogServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return dialogServiceServiceInfoForClient
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
	serviceName := "RPCDialog"
	handlerType := (*dialog.RPCDialog)(nil)
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
		"PackageName": "dialog",
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

func saveDraftMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SaveDraftMessageArgs)
	realResult := result.(*SaveDraftMessageResult)
	success, err := handler.(dialog.RPCDialog).DialogSaveDraftMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSaveDraftMessageArgs() interface{} {
	return &SaveDraftMessageArgs{}
}

func newSaveDraftMessageResult() interface{} {
	return &SaveDraftMessageResult{}
}

type SaveDraftMessageArgs struct {
	Req *dialog.TLDialogSaveDraftMessage
}

func (p *SaveDraftMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SaveDraftMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SaveDraftMessageArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogSaveDraftMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SaveDraftMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SaveDraftMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SaveDraftMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogSaveDraftMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SaveDraftMessageArgs_Req_DEFAULT *dialog.TLDialogSaveDraftMessage

func (p *SaveDraftMessageArgs) GetReq() *dialog.TLDialogSaveDraftMessage {
	if !p.IsSetReq() {
		return SaveDraftMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SaveDraftMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type SaveDraftMessageResult struct {
	Success *tg.Bool
}

var SaveDraftMessageResult_Success_DEFAULT *tg.Bool

func (p *SaveDraftMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SaveDraftMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *SaveDraftMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SaveDraftMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SaveDraftMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SaveDraftMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SaveDraftMessageResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SaveDraftMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SaveDraftMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SaveDraftMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SaveDraftMessageResult) GetResult() interface{} {
	return p.Success
}

func clearDraftMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ClearDraftMessageArgs)
	realResult := result.(*ClearDraftMessageResult)
	success, err := handler.(dialog.RPCDialog).DialogClearDraftMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newClearDraftMessageArgs() interface{} {
	return &ClearDraftMessageArgs{}
}

func newClearDraftMessageResult() interface{} {
	return &ClearDraftMessageResult{}
}

type ClearDraftMessageArgs struct {
	Req *dialog.TLDialogClearDraftMessage
}

func (p *ClearDraftMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ClearDraftMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ClearDraftMessageArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogClearDraftMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ClearDraftMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ClearDraftMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ClearDraftMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogClearDraftMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ClearDraftMessageArgs_Req_DEFAULT *dialog.TLDialogClearDraftMessage

func (p *ClearDraftMessageArgs) GetReq() *dialog.TLDialogClearDraftMessage {
	if !p.IsSetReq() {
		return ClearDraftMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ClearDraftMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type ClearDraftMessageResult struct {
	Success *tg.Bool
}

var ClearDraftMessageResult_Success_DEFAULT *tg.Bool

func (p *ClearDraftMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ClearDraftMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *ClearDraftMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ClearDraftMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ClearDraftMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ClearDraftMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ClearDraftMessageResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ClearDraftMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ClearDraftMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ClearDraftMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ClearDraftMessageResult) GetResult() interface{} {
	return p.Success
}

func getAllDraftsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAllDraftsArgs)
	realResult := result.(*GetAllDraftsResult)
	success, err := handler.(dialog.RPCDialog).DialogGetAllDrafts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAllDraftsArgs() interface{} {
	return &GetAllDraftsArgs{}
}

func newGetAllDraftsResult() interface{} {
	return &GetAllDraftsResult{}
}

type GetAllDraftsArgs struct {
	Req *dialog.TLDialogGetAllDrafts
}

func (p *GetAllDraftsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAllDraftsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAllDraftsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetAllDrafts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAllDraftsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAllDraftsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAllDraftsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetAllDrafts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAllDraftsArgs_Req_DEFAULT *dialog.TLDialogGetAllDrafts

func (p *GetAllDraftsArgs) GetReq() *dialog.TLDialogGetAllDrafts {
	if !p.IsSetReq() {
		return GetAllDraftsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAllDraftsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAllDraftsResult struct {
	Success *dialog.VectorPeerWithDraftMessage
}

var GetAllDraftsResult_Success_DEFAULT *dialog.VectorPeerWithDraftMessage

func (p *GetAllDraftsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAllDraftsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAllDraftsResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorPeerWithDraftMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAllDraftsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAllDraftsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAllDraftsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorPeerWithDraftMessage)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAllDraftsResult) GetSuccess() *dialog.VectorPeerWithDraftMessage {
	if !p.IsSetSuccess() {
		return GetAllDraftsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAllDraftsResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorPeerWithDraftMessage)
}

func (p *GetAllDraftsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAllDraftsResult) GetResult() interface{} {
	return p.Success
}

func clearAllDraftsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ClearAllDraftsArgs)
	realResult := result.(*ClearAllDraftsResult)
	success, err := handler.(dialog.RPCDialog).DialogClearAllDrafts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newClearAllDraftsArgs() interface{} {
	return &ClearAllDraftsArgs{}
}

func newClearAllDraftsResult() interface{} {
	return &ClearAllDraftsResult{}
}

type ClearAllDraftsArgs struct {
	Req *dialog.TLDialogClearAllDrafts
}

func (p *ClearAllDraftsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ClearAllDraftsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ClearAllDraftsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogClearAllDrafts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ClearAllDraftsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ClearAllDraftsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ClearAllDraftsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogClearAllDrafts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ClearAllDraftsArgs_Req_DEFAULT *dialog.TLDialogClearAllDrafts

func (p *ClearAllDraftsArgs) GetReq() *dialog.TLDialogClearAllDrafts {
	if !p.IsSetReq() {
		return ClearAllDraftsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ClearAllDraftsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ClearAllDraftsResult struct {
	Success *dialog.VectorPeerWithDraftMessage
}

var ClearAllDraftsResult_Success_DEFAULT *dialog.VectorPeerWithDraftMessage

func (p *ClearAllDraftsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ClearAllDraftsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ClearAllDraftsResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorPeerWithDraftMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ClearAllDraftsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ClearAllDraftsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ClearAllDraftsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorPeerWithDraftMessage)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ClearAllDraftsResult) GetSuccess() *dialog.VectorPeerWithDraftMessage {
	if !p.IsSetSuccess() {
		return ClearAllDraftsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ClearAllDraftsResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorPeerWithDraftMessage)
}

func (p *ClearAllDraftsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ClearAllDraftsResult) GetResult() interface{} {
	return p.Success
}

func markDialogUnreadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MarkDialogUnreadArgs)
	realResult := result.(*MarkDialogUnreadResult)
	success, err := handler.(dialog.RPCDialog).DialogMarkDialogUnread(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMarkDialogUnreadArgs() interface{} {
	return &MarkDialogUnreadArgs{}
}

func newMarkDialogUnreadResult() interface{} {
	return &MarkDialogUnreadResult{}
}

type MarkDialogUnreadArgs struct {
	Req *dialog.TLDialogMarkDialogUnread
}

func (p *MarkDialogUnreadArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MarkDialogUnreadArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MarkDialogUnreadArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogMarkDialogUnread)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MarkDialogUnreadArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MarkDialogUnreadArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MarkDialogUnreadArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogMarkDialogUnread)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MarkDialogUnreadArgs_Req_DEFAULT *dialog.TLDialogMarkDialogUnread

func (p *MarkDialogUnreadArgs) GetReq() *dialog.TLDialogMarkDialogUnread {
	if !p.IsSetReq() {
		return MarkDialogUnreadArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MarkDialogUnreadArgs) IsSetReq() bool {
	return p.Req != nil
}

type MarkDialogUnreadResult struct {
	Success *tg.Bool
}

var MarkDialogUnreadResult_Success_DEFAULT *tg.Bool

func (p *MarkDialogUnreadResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MarkDialogUnreadResult")
	}
	return json.Marshal(p.Success)
}

func (p *MarkDialogUnreadResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MarkDialogUnreadResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MarkDialogUnreadResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MarkDialogUnreadResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MarkDialogUnreadResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return MarkDialogUnreadResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MarkDialogUnreadResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *MarkDialogUnreadResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MarkDialogUnreadResult) GetResult() interface{} {
	return p.Success
}

func toggleDialogPinHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ToggleDialogPinArgs)
	realResult := result.(*ToggleDialogPinResult)
	success, err := handler.(dialog.RPCDialog).DialogToggleDialogPin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newToggleDialogPinArgs() interface{} {
	return &ToggleDialogPinArgs{}
}

func newToggleDialogPinResult() interface{} {
	return &ToggleDialogPinResult{}
}

type ToggleDialogPinArgs struct {
	Req *dialog.TLDialogToggleDialogPin
}

func (p *ToggleDialogPinArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ToggleDialogPinArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ToggleDialogPinArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogToggleDialogPin)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ToggleDialogPinArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ToggleDialogPinArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ToggleDialogPinArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogToggleDialogPin)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ToggleDialogPinArgs_Req_DEFAULT *dialog.TLDialogToggleDialogPin

func (p *ToggleDialogPinArgs) GetReq() *dialog.TLDialogToggleDialogPin {
	if !p.IsSetReq() {
		return ToggleDialogPinArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ToggleDialogPinArgs) IsSetReq() bool {
	return p.Req != nil
}

type ToggleDialogPinResult struct {
	Success *tg.Int32
}

var ToggleDialogPinResult_Success_DEFAULT *tg.Int32

func (p *ToggleDialogPinResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ToggleDialogPinResult")
	}
	return json.Marshal(p.Success)
}

func (p *ToggleDialogPinResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ToggleDialogPinResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ToggleDialogPinResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ToggleDialogPinResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ToggleDialogPinResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return ToggleDialogPinResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ToggleDialogPinResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *ToggleDialogPinResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ToggleDialogPinResult) GetResult() interface{} {
	return p.Success
}

func getDialogUnreadMarkListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogUnreadMarkListArgs)
	realResult := result.(*GetDialogUnreadMarkListResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogUnreadMarkList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogUnreadMarkListArgs() interface{} {
	return &GetDialogUnreadMarkListArgs{}
}

func newGetDialogUnreadMarkListResult() interface{} {
	return &GetDialogUnreadMarkListResult{}
}

type GetDialogUnreadMarkListArgs struct {
	Req *dialog.TLDialogGetDialogUnreadMarkList
}

func (p *GetDialogUnreadMarkListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogUnreadMarkListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogUnreadMarkListArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogUnreadMarkList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogUnreadMarkListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogUnreadMarkListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogUnreadMarkListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogUnreadMarkList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogUnreadMarkListArgs_Req_DEFAULT *dialog.TLDialogGetDialogUnreadMarkList

func (p *GetDialogUnreadMarkListArgs) GetReq() *dialog.TLDialogGetDialogUnreadMarkList {
	if !p.IsSetReq() {
		return GetDialogUnreadMarkListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogUnreadMarkListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogUnreadMarkListResult struct {
	Success *dialog.VectorDialogPeer
}

var GetDialogUnreadMarkListResult_Success_DEFAULT *dialog.VectorDialogPeer

func (p *GetDialogUnreadMarkListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogUnreadMarkListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogUnreadMarkListResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorDialogPeer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogUnreadMarkListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogUnreadMarkListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogUnreadMarkListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorDialogPeer)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogUnreadMarkListResult) GetSuccess() *dialog.VectorDialogPeer {
	if !p.IsSetSuccess() {
		return GetDialogUnreadMarkListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogUnreadMarkListResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorDialogPeer)
}

func (p *GetDialogUnreadMarkListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogUnreadMarkListResult) GetResult() interface{} {
	return p.Success
}

func getDialogsByOffsetDateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogsByOffsetDateArgs)
	realResult := result.(*GetDialogsByOffsetDateResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogsByOffsetDate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogsByOffsetDateArgs() interface{} {
	return &GetDialogsByOffsetDateArgs{}
}

func newGetDialogsByOffsetDateResult() interface{} {
	return &GetDialogsByOffsetDateResult{}
}

type GetDialogsByOffsetDateArgs struct {
	Req *dialog.TLDialogGetDialogsByOffsetDate
}

func (p *GetDialogsByOffsetDateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogsByOffsetDateArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogsByOffsetDateArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogsByOffsetDate)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogsByOffsetDateArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogsByOffsetDateArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogsByOffsetDateArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogsByOffsetDate)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogsByOffsetDateArgs_Req_DEFAULT *dialog.TLDialogGetDialogsByOffsetDate

func (p *GetDialogsByOffsetDateArgs) GetReq() *dialog.TLDialogGetDialogsByOffsetDate {
	if !p.IsSetReq() {
		return GetDialogsByOffsetDateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogsByOffsetDateArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogsByOffsetDateResult struct {
	Success *dialog.VectorDialogExt
}

var GetDialogsByOffsetDateResult_Success_DEFAULT *dialog.VectorDialogExt

func (p *GetDialogsByOffsetDateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogsByOffsetDateResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogsByOffsetDateResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorDialogExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsByOffsetDateResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogsByOffsetDateResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogsByOffsetDateResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorDialogExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsByOffsetDateResult) GetSuccess() *dialog.VectorDialogExt {
	if !p.IsSetSuccess() {
		return GetDialogsByOffsetDateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogsByOffsetDateResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorDialogExt)
}

func (p *GetDialogsByOffsetDateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogsByOffsetDateResult) GetResult() interface{} {
	return p.Success
}

func getDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogsArgs)
	realResult := result.(*GetDialogsResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogsArgs() interface{} {
	return &GetDialogsArgs{}
}

func newGetDialogsResult() interface{} {
	return &GetDialogsResult{}
}

type GetDialogsArgs struct {
	Req *dialog.TLDialogGetDialogs
}

func (p *GetDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogsArgs_Req_DEFAULT *dialog.TLDialogGetDialogs

func (p *GetDialogsArgs) GetReq() *dialog.TLDialogGetDialogs {
	if !p.IsSetReq() {
		return GetDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogsResult struct {
	Success *dialog.VectorDialogExt
}

var GetDialogsResult_Success_DEFAULT *dialog.VectorDialogExt

func (p *GetDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogsResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorDialogExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorDialogExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsResult) GetSuccess() *dialog.VectorDialogExt {
	if !p.IsSetSuccess() {
		return GetDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorDialogExt)
}

func (p *GetDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogsResult) GetResult() interface{} {
	return p.Success
}

func getDialogsByIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogsByIdListArgs)
	realResult := result.(*GetDialogsByIdListResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogsByIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogsByIdListArgs() interface{} {
	return &GetDialogsByIdListArgs{}
}

func newGetDialogsByIdListResult() interface{} {
	return &GetDialogsByIdListResult{}
}

type GetDialogsByIdListArgs struct {
	Req *dialog.TLDialogGetDialogsByIdList
}

func (p *GetDialogsByIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogsByIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogsByIdListArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogsByIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogsByIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogsByIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogsByIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogsByIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogsByIdListArgs_Req_DEFAULT *dialog.TLDialogGetDialogsByIdList

func (p *GetDialogsByIdListArgs) GetReq() *dialog.TLDialogGetDialogsByIdList {
	if !p.IsSetReq() {
		return GetDialogsByIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogsByIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogsByIdListResult struct {
	Success *dialog.VectorDialogExt
}

var GetDialogsByIdListResult_Success_DEFAULT *dialog.VectorDialogExt

func (p *GetDialogsByIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogsByIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogsByIdListResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorDialogExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsByIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogsByIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogsByIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorDialogExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsByIdListResult) GetSuccess() *dialog.VectorDialogExt {
	if !p.IsSetSuccess() {
		return GetDialogsByIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogsByIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorDialogExt)
}

func (p *GetDialogsByIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogsByIdListResult) GetResult() interface{} {
	return p.Success
}

func getDialogsCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogsCountArgs)
	realResult := result.(*GetDialogsCountResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogsCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogsCountArgs() interface{} {
	return &GetDialogsCountArgs{}
}

func newGetDialogsCountResult() interface{} {
	return &GetDialogsCountResult{}
}

type GetDialogsCountArgs struct {
	Req *dialog.TLDialogGetDialogsCount
}

func (p *GetDialogsCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogsCountArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogsCountArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogsCount)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogsCountArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogsCountArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogsCountArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogsCount)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogsCountArgs_Req_DEFAULT *dialog.TLDialogGetDialogsCount

func (p *GetDialogsCountArgs) GetReq() *dialog.TLDialogGetDialogsCount {
	if !p.IsSetReq() {
		return GetDialogsCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogsCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogsCountResult struct {
	Success *tg.Int32
}

var GetDialogsCountResult_Success_DEFAULT *tg.Int32

func (p *GetDialogsCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogsCountResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogsCountResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsCountResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogsCountResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogsCountResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogsCountResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetDialogsCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogsCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetDialogsCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogsCountResult) GetResult() interface{} {
	return p.Success
}

func getPinnedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPinnedDialogsArgs)
	realResult := result.(*GetPinnedDialogsResult)
	success, err := handler.(dialog.RPCDialog).DialogGetPinnedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPinnedDialogsArgs() interface{} {
	return &GetPinnedDialogsArgs{}
}

func newGetPinnedDialogsResult() interface{} {
	return &GetPinnedDialogsResult{}
}

type GetPinnedDialogsArgs struct {
	Req *dialog.TLDialogGetPinnedDialogs
}

func (p *GetPinnedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPinnedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPinnedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetPinnedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPinnedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPinnedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPinnedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetPinnedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPinnedDialogsArgs_Req_DEFAULT *dialog.TLDialogGetPinnedDialogs

func (p *GetPinnedDialogsArgs) GetReq() *dialog.TLDialogGetPinnedDialogs {
	if !p.IsSetReq() {
		return GetPinnedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPinnedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPinnedDialogsResult struct {
	Success *dialog.VectorDialogExt
}

var GetPinnedDialogsResult_Success_DEFAULT *dialog.VectorDialogExt

func (p *GetPinnedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPinnedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPinnedDialogsResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorDialogExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPinnedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPinnedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPinnedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorDialogExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPinnedDialogsResult) GetSuccess() *dialog.VectorDialogExt {
	if !p.IsSetSuccess() {
		return GetPinnedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPinnedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorDialogExt)
}

func (p *GetPinnedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPinnedDialogsResult) GetResult() interface{} {
	return p.Success
}

func reorderPinnedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReorderPinnedDialogsArgs)
	realResult := result.(*ReorderPinnedDialogsResult)
	success, err := handler.(dialog.RPCDialog).DialogReorderPinnedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReorderPinnedDialogsArgs() interface{} {
	return &ReorderPinnedDialogsArgs{}
}

func newReorderPinnedDialogsResult() interface{} {
	return &ReorderPinnedDialogsResult{}
}

type ReorderPinnedDialogsArgs struct {
	Req *dialog.TLDialogReorderPinnedDialogs
}

func (p *ReorderPinnedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReorderPinnedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ReorderPinnedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogReorderPinnedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReorderPinnedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReorderPinnedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReorderPinnedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogReorderPinnedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReorderPinnedDialogsArgs_Req_DEFAULT *dialog.TLDialogReorderPinnedDialogs

func (p *ReorderPinnedDialogsArgs) GetReq() *dialog.TLDialogReorderPinnedDialogs {
	if !p.IsSetReq() {
		return ReorderPinnedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReorderPinnedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ReorderPinnedDialogsResult struct {
	Success *tg.Bool
}

var ReorderPinnedDialogsResult_Success_DEFAULT *tg.Bool

func (p *ReorderPinnedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReorderPinnedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ReorderPinnedDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReorderPinnedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReorderPinnedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReorderPinnedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReorderPinnedDialogsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ReorderPinnedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReorderPinnedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ReorderPinnedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReorderPinnedDialogsResult) GetResult() interface{} {
	return p.Success
}

func getDialogByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogByIdArgs)
	realResult := result.(*GetDialogByIdResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogById(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogByIdArgs() interface{} {
	return &GetDialogByIdArgs{}
}

func newGetDialogByIdResult() interface{} {
	return &GetDialogByIdResult{}
}

type GetDialogByIdArgs struct {
	Req *dialog.TLDialogGetDialogById
}

func (p *GetDialogByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogByIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogByIdArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogById)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogByIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogByIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogByIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogById)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogByIdArgs_Req_DEFAULT *dialog.TLDialogGetDialogById

func (p *GetDialogByIdArgs) GetReq() *dialog.TLDialogGetDialogById {
	if !p.IsSetReq() {
		return GetDialogByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogByIdResult struct {
	Success *dialog.DialogExt
}

var GetDialogByIdResult_Success_DEFAULT *dialog.DialogExt

func (p *GetDialogByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogByIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogByIdResult) Unmarshal(in []byte) error {
	msg := new(dialog.DialogExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogByIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogByIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogByIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.DialogExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogByIdResult) GetSuccess() *dialog.DialogExt {
	if !p.IsSetSuccess() {
		return GetDialogByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.DialogExt)
}

func (p *GetDialogByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogByIdResult) GetResult() interface{} {
	return p.Success
}

func getTopMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetTopMessageArgs)
	realResult := result.(*GetTopMessageResult)
	success, err := handler.(dialog.RPCDialog).DialogGetTopMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetTopMessageArgs() interface{} {
	return &GetTopMessageArgs{}
}

func newGetTopMessageResult() interface{} {
	return &GetTopMessageResult{}
}

type GetTopMessageArgs struct {
	Req *dialog.TLDialogGetTopMessage
}

func (p *GetTopMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetTopMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetTopMessageArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetTopMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetTopMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetTopMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetTopMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetTopMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetTopMessageArgs_Req_DEFAULT *dialog.TLDialogGetTopMessage

func (p *GetTopMessageArgs) GetReq() *dialog.TLDialogGetTopMessage {
	if !p.IsSetReq() {
		return GetTopMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetTopMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetTopMessageResult struct {
	Success *tg.Int32
}

var GetTopMessageResult_Success_DEFAULT *tg.Int32

func (p *GetTopMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetTopMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetTopMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetTopMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetTopMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetTopMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetTopMessageResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetTopMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetTopMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetTopMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetTopMessageResult) GetResult() interface{} {
	return p.Success
}

func insertOrUpdateDialogHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*InsertOrUpdateDialogArgs)
	realResult := result.(*InsertOrUpdateDialogResult)
	success, err := handler.(dialog.RPCDialog).DialogInsertOrUpdateDialog(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newInsertOrUpdateDialogArgs() interface{} {
	return &InsertOrUpdateDialogArgs{}
}

func newInsertOrUpdateDialogResult() interface{} {
	return &InsertOrUpdateDialogResult{}
}

type InsertOrUpdateDialogArgs struct {
	Req *dialog.TLDialogInsertOrUpdateDialog
}

func (p *InsertOrUpdateDialogArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in InsertOrUpdateDialogArgs")
	}
	return json.Marshal(p.Req)
}

func (p *InsertOrUpdateDialogArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogInsertOrUpdateDialog)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *InsertOrUpdateDialogArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in InsertOrUpdateDialogArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *InsertOrUpdateDialogArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogInsertOrUpdateDialog)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var InsertOrUpdateDialogArgs_Req_DEFAULT *dialog.TLDialogInsertOrUpdateDialog

func (p *InsertOrUpdateDialogArgs) GetReq() *dialog.TLDialogInsertOrUpdateDialog {
	if !p.IsSetReq() {
		return InsertOrUpdateDialogArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *InsertOrUpdateDialogArgs) IsSetReq() bool {
	return p.Req != nil
}

type InsertOrUpdateDialogResult struct {
	Success *tg.Bool
}

var InsertOrUpdateDialogResult_Success_DEFAULT *tg.Bool

func (p *InsertOrUpdateDialogResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in InsertOrUpdateDialogResult")
	}
	return json.Marshal(p.Success)
}

func (p *InsertOrUpdateDialogResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *InsertOrUpdateDialogResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in InsertOrUpdateDialogResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *InsertOrUpdateDialogResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *InsertOrUpdateDialogResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return InsertOrUpdateDialogResult_Success_DEFAULT
	}
	return p.Success
}

func (p *InsertOrUpdateDialogResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *InsertOrUpdateDialogResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *InsertOrUpdateDialogResult) GetResult() interface{} {
	return p.Success
}

func deleteDialogHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteDialogArgs)
	realResult := result.(*DeleteDialogResult)
	success, err := handler.(dialog.RPCDialog).DialogDeleteDialog(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteDialogArgs() interface{} {
	return &DeleteDialogArgs{}
}

func newDeleteDialogResult() interface{} {
	return &DeleteDialogResult{}
}

type DeleteDialogArgs struct {
	Req *dialog.TLDialogDeleteDialog
}

func (p *DeleteDialogArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteDialogArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteDialogArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogDeleteDialog)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteDialogArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteDialogArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteDialogArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogDeleteDialog)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteDialogArgs_Req_DEFAULT *dialog.TLDialogDeleteDialog

func (p *DeleteDialogArgs) GetReq() *dialog.TLDialogDeleteDialog {
	if !p.IsSetReq() {
		return DeleteDialogArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteDialogArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteDialogResult struct {
	Success *tg.Bool
}

var DeleteDialogResult_Success_DEFAULT *tg.Bool

func (p *DeleteDialogResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteDialogResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteDialogResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteDialogResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteDialogResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteDialogResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteDialogResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteDialogResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteDialogResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteDialogResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteDialogResult) GetResult() interface{} {
	return p.Success
}

func getUserPinnedMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserPinnedMessageArgs)
	realResult := result.(*GetUserPinnedMessageResult)
	success, err := handler.(dialog.RPCDialog).DialogGetUserPinnedMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserPinnedMessageArgs() interface{} {
	return &GetUserPinnedMessageArgs{}
}

func newGetUserPinnedMessageResult() interface{} {
	return &GetUserPinnedMessageResult{}
}

type GetUserPinnedMessageArgs struct {
	Req *dialog.TLDialogGetUserPinnedMessage
}

func (p *GetUserPinnedMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserPinnedMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserPinnedMessageArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetUserPinnedMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserPinnedMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserPinnedMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserPinnedMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetUserPinnedMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserPinnedMessageArgs_Req_DEFAULT *dialog.TLDialogGetUserPinnedMessage

func (p *GetUserPinnedMessageArgs) GetReq() *dialog.TLDialogGetUserPinnedMessage {
	if !p.IsSetReq() {
		return GetUserPinnedMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserPinnedMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserPinnedMessageResult struct {
	Success *tg.Int32
}

var GetUserPinnedMessageResult_Success_DEFAULT *tg.Int32

func (p *GetUserPinnedMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserPinnedMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserPinnedMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Int32)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserPinnedMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserPinnedMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserPinnedMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int32)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserPinnedMessageResult) GetSuccess() *tg.Int32 {
	if !p.IsSetSuccess() {
		return GetUserPinnedMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserPinnedMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int32)
}

func (p *GetUserPinnedMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserPinnedMessageResult) GetResult() interface{} {
	return p.Success
}

func updateUserPinnedMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateUserPinnedMessageArgs)
	realResult := result.(*UpdateUserPinnedMessageResult)
	success, err := handler.(dialog.RPCDialog).DialogUpdateUserPinnedMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateUserPinnedMessageArgs() interface{} {
	return &UpdateUserPinnedMessageArgs{}
}

func newUpdateUserPinnedMessageResult() interface{} {
	return &UpdateUserPinnedMessageResult{}
}

type UpdateUserPinnedMessageArgs struct {
	Req *dialog.TLDialogUpdateUserPinnedMessage
}

func (p *UpdateUserPinnedMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateUserPinnedMessageArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateUserPinnedMessageArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogUpdateUserPinnedMessage)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateUserPinnedMessageArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateUserPinnedMessageArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateUserPinnedMessageArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogUpdateUserPinnedMessage)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateUserPinnedMessageArgs_Req_DEFAULT *dialog.TLDialogUpdateUserPinnedMessage

func (p *UpdateUserPinnedMessageArgs) GetReq() *dialog.TLDialogUpdateUserPinnedMessage {
	if !p.IsSetReq() {
		return UpdateUserPinnedMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserPinnedMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateUserPinnedMessageResult struct {
	Success *tg.Bool
}

var UpdateUserPinnedMessageResult_Success_DEFAULT *tg.Bool

func (p *UpdateUserPinnedMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateUserPinnedMessageResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateUserPinnedMessageResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserPinnedMessageResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateUserPinnedMessageResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateUserPinnedMessageResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserPinnedMessageResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateUserPinnedMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserPinnedMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateUserPinnedMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserPinnedMessageResult) GetResult() interface{} {
	return p.Success
}

func insertOrUpdateDialogFilterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*InsertOrUpdateDialogFilterArgs)
	realResult := result.(*InsertOrUpdateDialogFilterResult)
	success, err := handler.(dialog.RPCDialog).DialogInsertOrUpdateDialogFilter(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newInsertOrUpdateDialogFilterArgs() interface{} {
	return &InsertOrUpdateDialogFilterArgs{}
}

func newInsertOrUpdateDialogFilterResult() interface{} {
	return &InsertOrUpdateDialogFilterResult{}
}

type InsertOrUpdateDialogFilterArgs struct {
	Req *dialog.TLDialogInsertOrUpdateDialogFilter
}

func (p *InsertOrUpdateDialogFilterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in InsertOrUpdateDialogFilterArgs")
	}
	return json.Marshal(p.Req)
}

func (p *InsertOrUpdateDialogFilterArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogInsertOrUpdateDialogFilter)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *InsertOrUpdateDialogFilterArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in InsertOrUpdateDialogFilterArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *InsertOrUpdateDialogFilterArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogInsertOrUpdateDialogFilter)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var InsertOrUpdateDialogFilterArgs_Req_DEFAULT *dialog.TLDialogInsertOrUpdateDialogFilter

func (p *InsertOrUpdateDialogFilterArgs) GetReq() *dialog.TLDialogInsertOrUpdateDialogFilter {
	if !p.IsSetReq() {
		return InsertOrUpdateDialogFilterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *InsertOrUpdateDialogFilterArgs) IsSetReq() bool {
	return p.Req != nil
}

type InsertOrUpdateDialogFilterResult struct {
	Success *tg.Bool
}

var InsertOrUpdateDialogFilterResult_Success_DEFAULT *tg.Bool

func (p *InsertOrUpdateDialogFilterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in InsertOrUpdateDialogFilterResult")
	}
	return json.Marshal(p.Success)
}

func (p *InsertOrUpdateDialogFilterResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *InsertOrUpdateDialogFilterResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in InsertOrUpdateDialogFilterResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *InsertOrUpdateDialogFilterResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *InsertOrUpdateDialogFilterResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return InsertOrUpdateDialogFilterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *InsertOrUpdateDialogFilterResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *InsertOrUpdateDialogFilterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *InsertOrUpdateDialogFilterResult) GetResult() interface{} {
	return p.Success
}

func deleteDialogFilterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteDialogFilterArgs)
	realResult := result.(*DeleteDialogFilterResult)
	success, err := handler.(dialog.RPCDialog).DialogDeleteDialogFilter(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteDialogFilterArgs() interface{} {
	return &DeleteDialogFilterArgs{}
}

func newDeleteDialogFilterResult() interface{} {
	return &DeleteDialogFilterResult{}
}

type DeleteDialogFilterArgs struct {
	Req *dialog.TLDialogDeleteDialogFilter
}

func (p *DeleteDialogFilterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteDialogFilterArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteDialogFilterArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogDeleteDialogFilter)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteDialogFilterArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteDialogFilterArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteDialogFilterArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogDeleteDialogFilter)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteDialogFilterArgs_Req_DEFAULT *dialog.TLDialogDeleteDialogFilter

func (p *DeleteDialogFilterArgs) GetReq() *dialog.TLDialogDeleteDialogFilter {
	if !p.IsSetReq() {
		return DeleteDialogFilterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteDialogFilterArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteDialogFilterResult struct {
	Success *tg.Bool
}

var DeleteDialogFilterResult_Success_DEFAULT *tg.Bool

func (p *DeleteDialogFilterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteDialogFilterResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteDialogFilterResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteDialogFilterResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteDialogFilterResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteDialogFilterResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteDialogFilterResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteDialogFilterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteDialogFilterResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteDialogFilterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteDialogFilterResult) GetResult() interface{} {
	return p.Success
}

func updateDialogFiltersOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateDialogFiltersOrderArgs)
	realResult := result.(*UpdateDialogFiltersOrderResult)
	success, err := handler.(dialog.RPCDialog).DialogUpdateDialogFiltersOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateDialogFiltersOrderArgs() interface{} {
	return &UpdateDialogFiltersOrderArgs{}
}

func newUpdateDialogFiltersOrderResult() interface{} {
	return &UpdateDialogFiltersOrderResult{}
}

type UpdateDialogFiltersOrderArgs struct {
	Req *dialog.TLDialogUpdateDialogFiltersOrder
}

func (p *UpdateDialogFiltersOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateDialogFiltersOrderArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateDialogFiltersOrderArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogUpdateDialogFiltersOrder)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateDialogFiltersOrderArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateDialogFiltersOrderArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateDialogFiltersOrderArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogUpdateDialogFiltersOrder)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateDialogFiltersOrderArgs_Req_DEFAULT *dialog.TLDialogUpdateDialogFiltersOrder

func (p *UpdateDialogFiltersOrderArgs) GetReq() *dialog.TLDialogUpdateDialogFiltersOrder {
	if !p.IsSetReq() {
		return UpdateDialogFiltersOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateDialogFiltersOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateDialogFiltersOrderResult struct {
	Success *tg.Bool
}

var UpdateDialogFiltersOrderResult_Success_DEFAULT *tg.Bool

func (p *UpdateDialogFiltersOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateDialogFiltersOrderResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateDialogFiltersOrderResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateDialogFiltersOrderResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateDialogFiltersOrderResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateDialogFiltersOrderResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateDialogFiltersOrderResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateDialogFiltersOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateDialogFiltersOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateDialogFiltersOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateDialogFiltersOrderResult) GetResult() interface{} {
	return p.Success
}

func getDialogFiltersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogFiltersArgs)
	realResult := result.(*GetDialogFiltersResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogFilters(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogFiltersArgs() interface{} {
	return &GetDialogFiltersArgs{}
}

func newGetDialogFiltersResult() interface{} {
	return &GetDialogFiltersResult{}
}

type GetDialogFiltersArgs struct {
	Req *dialog.TLDialogGetDialogFilters
}

func (p *GetDialogFiltersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogFiltersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogFiltersArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogFilters)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogFiltersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogFiltersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogFiltersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogFilters)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogFiltersArgs_Req_DEFAULT *dialog.TLDialogGetDialogFilters

func (p *GetDialogFiltersArgs) GetReq() *dialog.TLDialogGetDialogFilters {
	if !p.IsSetReq() {
		return GetDialogFiltersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogFiltersArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogFiltersResult struct {
	Success *dialog.VectorDialogFilterExt
}

var GetDialogFiltersResult_Success_DEFAULT *dialog.VectorDialogFilterExt

func (p *GetDialogFiltersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogFiltersResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogFiltersResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorDialogFilterExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFiltersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogFiltersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogFiltersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorDialogFilterExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFiltersResult) GetSuccess() *dialog.VectorDialogFilterExt {
	if !p.IsSetSuccess() {
		return GetDialogFiltersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogFiltersResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorDialogFilterExt)
}

func (p *GetDialogFiltersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogFiltersResult) GetResult() interface{} {
	return p.Success
}

func getDialogFolderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogFolderArgs)
	realResult := result.(*GetDialogFolderResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogFolder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogFolderArgs() interface{} {
	return &GetDialogFolderArgs{}
}

func newGetDialogFolderResult() interface{} {
	return &GetDialogFolderResult{}
}

type GetDialogFolderArgs struct {
	Req *dialog.TLDialogGetDialogFolder
}

func (p *GetDialogFolderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogFolderArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogFolderArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogFolder)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogFolderArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogFolderArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogFolderArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogFolder)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogFolderArgs_Req_DEFAULT *dialog.TLDialogGetDialogFolder

func (p *GetDialogFolderArgs) GetReq() *dialog.TLDialogGetDialogFolder {
	if !p.IsSetReq() {
		return GetDialogFolderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogFolderArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogFolderResult struct {
	Success *dialog.VectorDialogExt
}

var GetDialogFolderResult_Success_DEFAULT *dialog.VectorDialogExt

func (p *GetDialogFolderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogFolderResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogFolderResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorDialogExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFolderResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogFolderResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogFolderResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorDialogExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFolderResult) GetSuccess() *dialog.VectorDialogExt {
	if !p.IsSetSuccess() {
		return GetDialogFolderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogFolderResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorDialogExt)
}

func (p *GetDialogFolderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogFolderResult) GetResult() interface{} {
	return p.Success
}

func editPeerFoldersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditPeerFoldersArgs)
	realResult := result.(*EditPeerFoldersResult)
	success, err := handler.(dialog.RPCDialog).DialogEditPeerFolders(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditPeerFoldersArgs() interface{} {
	return &EditPeerFoldersArgs{}
}

func newEditPeerFoldersResult() interface{} {
	return &EditPeerFoldersResult{}
}

type EditPeerFoldersArgs struct {
	Req *dialog.TLDialogEditPeerFolders
}

func (p *EditPeerFoldersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditPeerFoldersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditPeerFoldersArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogEditPeerFolders)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditPeerFoldersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditPeerFoldersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditPeerFoldersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogEditPeerFolders)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditPeerFoldersArgs_Req_DEFAULT *dialog.TLDialogEditPeerFolders

func (p *EditPeerFoldersArgs) GetReq() *dialog.TLDialogEditPeerFolders {
	if !p.IsSetReq() {
		return EditPeerFoldersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditPeerFoldersArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditPeerFoldersResult struct {
	Success *dialog.VectorDialogPinnedExt
}

var EditPeerFoldersResult_Success_DEFAULT *dialog.VectorDialogPinnedExt

func (p *EditPeerFoldersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditPeerFoldersResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditPeerFoldersResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorDialogPinnedExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditPeerFoldersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditPeerFoldersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditPeerFoldersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorDialogPinnedExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditPeerFoldersResult) GetSuccess() *dialog.VectorDialogPinnedExt {
	if !p.IsSetSuccess() {
		return EditPeerFoldersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditPeerFoldersResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorDialogPinnedExt)
}

func (p *EditPeerFoldersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditPeerFoldersResult) GetResult() interface{} {
	return p.Success
}

func getChannelMessageReadParticipantsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetChannelMessageReadParticipantsArgs)
	realResult := result.(*GetChannelMessageReadParticipantsResult)
	success, err := handler.(dialog.RPCDialog).DialogGetChannelMessageReadParticipants(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetChannelMessageReadParticipantsArgs() interface{} {
	return &GetChannelMessageReadParticipantsArgs{}
}

func newGetChannelMessageReadParticipantsResult() interface{} {
	return &GetChannelMessageReadParticipantsResult{}
}

type GetChannelMessageReadParticipantsArgs struct {
	Req *dialog.TLDialogGetChannelMessageReadParticipants
}

func (p *GetChannelMessageReadParticipantsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChannelMessageReadParticipantsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetChannelMessageReadParticipantsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetChannelMessageReadParticipants)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetChannelMessageReadParticipantsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetChannelMessageReadParticipantsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetChannelMessageReadParticipantsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetChannelMessageReadParticipants)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetChannelMessageReadParticipantsArgs_Req_DEFAULT *dialog.TLDialogGetChannelMessageReadParticipants

func (p *GetChannelMessageReadParticipantsArgs) GetReq() *dialog.TLDialogGetChannelMessageReadParticipants {
	if !p.IsSetReq() {
		return GetChannelMessageReadParticipantsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChannelMessageReadParticipantsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetChannelMessageReadParticipantsResult struct {
	Success *dialog.VectorLong
}

var GetChannelMessageReadParticipantsResult_Success_DEFAULT *dialog.VectorLong

func (p *GetChannelMessageReadParticipantsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChannelMessageReadParticipantsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetChannelMessageReadParticipantsResult) Unmarshal(in []byte) error {
	msg := new(dialog.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChannelMessageReadParticipantsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetChannelMessageReadParticipantsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetChannelMessageReadParticipantsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChannelMessageReadParticipantsResult) GetSuccess() *dialog.VectorLong {
	if !p.IsSetSuccess() {
		return GetChannelMessageReadParticipantsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChannelMessageReadParticipantsResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.VectorLong)
}

func (p *GetChannelMessageReadParticipantsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetChannelMessageReadParticipantsResult) GetResult() interface{} {
	return p.Success
}

func setChatThemeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetChatThemeArgs)
	realResult := result.(*SetChatThemeResult)
	success, err := handler.(dialog.RPCDialog).DialogSetChatTheme(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetChatThemeArgs() interface{} {
	return &SetChatThemeArgs{}
}

func newSetChatThemeResult() interface{} {
	return &SetChatThemeResult{}
}

type SetChatThemeArgs struct {
	Req *dialog.TLDialogSetChatTheme
}

func (p *SetChatThemeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetChatThemeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetChatThemeArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogSetChatTheme)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetChatThemeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetChatThemeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetChatThemeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogSetChatTheme)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetChatThemeArgs_Req_DEFAULT *dialog.TLDialogSetChatTheme

func (p *SetChatThemeArgs) GetReq() *dialog.TLDialogSetChatTheme {
	if !p.IsSetReq() {
		return SetChatThemeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetChatThemeArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetChatThemeResult struct {
	Success *tg.Bool
}

var SetChatThemeResult_Success_DEFAULT *tg.Bool

func (p *SetChatThemeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetChatThemeResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetChatThemeResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChatThemeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetChatThemeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetChatThemeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChatThemeResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetChatThemeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetChatThemeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetChatThemeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetChatThemeResult) GetResult() interface{} {
	return p.Success
}

func setHistoryTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetHistoryTTLArgs)
	realResult := result.(*SetHistoryTTLResult)
	success, err := handler.(dialog.RPCDialog).DialogSetHistoryTTL(ctx, realArg.Req)
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
	Req *dialog.TLDialogSetHistoryTTL
}

func (p *SetHistoryTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetHistoryTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetHistoryTTLArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogSetHistoryTTL)
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
	msg := new(dialog.TLDialogSetHistoryTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetHistoryTTLArgs_Req_DEFAULT *dialog.TLDialogSetHistoryTTL

func (p *SetHistoryTTLArgs) GetReq() *dialog.TLDialogSetHistoryTTL {
	if !p.IsSetReq() {
		return SetHistoryTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetHistoryTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetHistoryTTLResult struct {
	Success *tg.Bool
}

var SetHistoryTTLResult_Success_DEFAULT *tg.Bool

func (p *SetHistoryTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetHistoryTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetHistoryTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
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
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetHistoryTTLResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetHistoryTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetHistoryTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetHistoryTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetHistoryTTLResult) GetResult() interface{} {
	return p.Success
}

func getMyDialogsDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetMyDialogsDataArgs)
	realResult := result.(*GetMyDialogsDataResult)
	success, err := handler.(dialog.RPCDialog).DialogGetMyDialogsData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetMyDialogsDataArgs() interface{} {
	return &GetMyDialogsDataArgs{}
}

func newGetMyDialogsDataResult() interface{} {
	return &GetMyDialogsDataResult{}
}

type GetMyDialogsDataArgs struct {
	Req *dialog.TLDialogGetMyDialogsData
}

func (p *GetMyDialogsDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetMyDialogsDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetMyDialogsDataArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetMyDialogsData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetMyDialogsDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetMyDialogsDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetMyDialogsDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetMyDialogsData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetMyDialogsDataArgs_Req_DEFAULT *dialog.TLDialogGetMyDialogsData

func (p *GetMyDialogsDataArgs) GetReq() *dialog.TLDialogGetMyDialogsData {
	if !p.IsSetReq() {
		return GetMyDialogsDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMyDialogsDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetMyDialogsDataResult struct {
	Success *dialog.DialogsData
}

var GetMyDialogsDataResult_Success_DEFAULT *dialog.DialogsData

func (p *GetMyDialogsDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetMyDialogsDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetMyDialogsDataResult) Unmarshal(in []byte) error {
	msg := new(dialog.DialogsData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMyDialogsDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetMyDialogsDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetMyDialogsDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.DialogsData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMyDialogsDataResult) GetSuccess() *dialog.DialogsData {
	if !p.IsSetSuccess() {
		return GetMyDialogsDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMyDialogsDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.DialogsData)
}

func (p *GetMyDialogsDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMyDialogsDataResult) GetResult() interface{} {
	return p.Success
}

func getSavedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetSavedDialogsArgs)
	realResult := result.(*GetSavedDialogsResult)
	success, err := handler.(dialog.RPCDialog).DialogGetSavedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetSavedDialogsArgs() interface{} {
	return &GetSavedDialogsArgs{}
}

func newGetSavedDialogsResult() interface{} {
	return &GetSavedDialogsResult{}
}

type GetSavedDialogsArgs struct {
	Req *dialog.TLDialogGetSavedDialogs
}

func (p *GetSavedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetSavedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetSavedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetSavedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetSavedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetSavedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetSavedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetSavedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetSavedDialogsArgs_Req_DEFAULT *dialog.TLDialogGetSavedDialogs

func (p *GetSavedDialogsArgs) GetReq() *dialog.TLDialogGetSavedDialogs {
	if !p.IsSetReq() {
		return GetSavedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetSavedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetSavedDialogsResult struct {
	Success *dialog.SavedDialogList
}

var GetSavedDialogsResult_Success_DEFAULT *dialog.SavedDialogList

func (p *GetSavedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetSavedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetSavedDialogsResult) Unmarshal(in []byte) error {
	msg := new(dialog.SavedDialogList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetSavedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetSavedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetSavedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.SavedDialogList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetSavedDialogsResult) GetSuccess() *dialog.SavedDialogList {
	if !p.IsSetSuccess() {
		return GetSavedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetSavedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.SavedDialogList)
}

func (p *GetSavedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetSavedDialogsResult) GetResult() interface{} {
	return p.Success
}

func getPinnedSavedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPinnedSavedDialogsArgs)
	realResult := result.(*GetPinnedSavedDialogsResult)
	success, err := handler.(dialog.RPCDialog).DialogGetPinnedSavedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPinnedSavedDialogsArgs() interface{} {
	return &GetPinnedSavedDialogsArgs{}
}

func newGetPinnedSavedDialogsResult() interface{} {
	return &GetPinnedSavedDialogsResult{}
}

type GetPinnedSavedDialogsArgs struct {
	Req *dialog.TLDialogGetPinnedSavedDialogs
}

func (p *GetPinnedSavedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPinnedSavedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPinnedSavedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetPinnedSavedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPinnedSavedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPinnedSavedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPinnedSavedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetPinnedSavedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPinnedSavedDialogsArgs_Req_DEFAULT *dialog.TLDialogGetPinnedSavedDialogs

func (p *GetPinnedSavedDialogsArgs) GetReq() *dialog.TLDialogGetPinnedSavedDialogs {
	if !p.IsSetReq() {
		return GetPinnedSavedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPinnedSavedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPinnedSavedDialogsResult struct {
	Success *dialog.SavedDialogList
}

var GetPinnedSavedDialogsResult_Success_DEFAULT *dialog.SavedDialogList

func (p *GetPinnedSavedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPinnedSavedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPinnedSavedDialogsResult) Unmarshal(in []byte) error {
	msg := new(dialog.SavedDialogList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPinnedSavedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPinnedSavedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPinnedSavedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.SavedDialogList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPinnedSavedDialogsResult) GetSuccess() *dialog.SavedDialogList {
	if !p.IsSetSuccess() {
		return GetPinnedSavedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPinnedSavedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.SavedDialogList)
}

func (p *GetPinnedSavedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPinnedSavedDialogsResult) GetResult() interface{} {
	return p.Success
}

func toggleSavedDialogPinHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ToggleSavedDialogPinArgs)
	realResult := result.(*ToggleSavedDialogPinResult)
	success, err := handler.(dialog.RPCDialog).DialogToggleSavedDialogPin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newToggleSavedDialogPinArgs() interface{} {
	return &ToggleSavedDialogPinArgs{}
}

func newToggleSavedDialogPinResult() interface{} {
	return &ToggleSavedDialogPinResult{}
}

type ToggleSavedDialogPinArgs struct {
	Req *dialog.TLDialogToggleSavedDialogPin
}

func (p *ToggleSavedDialogPinArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ToggleSavedDialogPinArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ToggleSavedDialogPinArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogToggleSavedDialogPin)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ToggleSavedDialogPinArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ToggleSavedDialogPinArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ToggleSavedDialogPinArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogToggleSavedDialogPin)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ToggleSavedDialogPinArgs_Req_DEFAULT *dialog.TLDialogToggleSavedDialogPin

func (p *ToggleSavedDialogPinArgs) GetReq() *dialog.TLDialogToggleSavedDialogPin {
	if !p.IsSetReq() {
		return ToggleSavedDialogPinArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ToggleSavedDialogPinArgs) IsSetReq() bool {
	return p.Req != nil
}

type ToggleSavedDialogPinResult struct {
	Success *tg.Bool
}

var ToggleSavedDialogPinResult_Success_DEFAULT *tg.Bool

func (p *ToggleSavedDialogPinResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ToggleSavedDialogPinResult")
	}
	return json.Marshal(p.Success)
}

func (p *ToggleSavedDialogPinResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ToggleSavedDialogPinResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ToggleSavedDialogPinResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ToggleSavedDialogPinResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ToggleSavedDialogPinResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ToggleSavedDialogPinResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ToggleSavedDialogPinResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ToggleSavedDialogPinResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ToggleSavedDialogPinResult) GetResult() interface{} {
	return p.Success
}

func reorderPinnedSavedDialogsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ReorderPinnedSavedDialogsArgs)
	realResult := result.(*ReorderPinnedSavedDialogsResult)
	success, err := handler.(dialog.RPCDialog).DialogReorderPinnedSavedDialogs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newReorderPinnedSavedDialogsArgs() interface{} {
	return &ReorderPinnedSavedDialogsArgs{}
}

func newReorderPinnedSavedDialogsResult() interface{} {
	return &ReorderPinnedSavedDialogsResult{}
}

type ReorderPinnedSavedDialogsArgs struct {
	Req *dialog.TLDialogReorderPinnedSavedDialogs
}

func (p *ReorderPinnedSavedDialogsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ReorderPinnedSavedDialogsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ReorderPinnedSavedDialogsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogReorderPinnedSavedDialogs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ReorderPinnedSavedDialogsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ReorderPinnedSavedDialogsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ReorderPinnedSavedDialogsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogReorderPinnedSavedDialogs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ReorderPinnedSavedDialogsArgs_Req_DEFAULT *dialog.TLDialogReorderPinnedSavedDialogs

func (p *ReorderPinnedSavedDialogsArgs) GetReq() *dialog.TLDialogReorderPinnedSavedDialogs {
	if !p.IsSetReq() {
		return ReorderPinnedSavedDialogsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReorderPinnedSavedDialogsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ReorderPinnedSavedDialogsResult struct {
	Success *tg.Bool
}

var ReorderPinnedSavedDialogsResult_Success_DEFAULT *tg.Bool

func (p *ReorderPinnedSavedDialogsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ReorderPinnedSavedDialogsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ReorderPinnedSavedDialogsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReorderPinnedSavedDialogsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ReorderPinnedSavedDialogsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ReorderPinnedSavedDialogsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReorderPinnedSavedDialogsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ReorderPinnedSavedDialogsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReorderPinnedSavedDialogsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ReorderPinnedSavedDialogsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReorderPinnedSavedDialogsResult) GetResult() interface{} {
	return p.Success
}

func getDialogFilterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogFilterArgs)
	realResult := result.(*GetDialogFilterResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogFilter(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogFilterArgs() interface{} {
	return &GetDialogFilterArgs{}
}

func newGetDialogFilterResult() interface{} {
	return &GetDialogFilterResult{}
}

type GetDialogFilterArgs struct {
	Req *dialog.TLDialogGetDialogFilter
}

func (p *GetDialogFilterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogFilterArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogFilterArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogFilter)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogFilterArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogFilterArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogFilterArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogFilter)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogFilterArgs_Req_DEFAULT *dialog.TLDialogGetDialogFilter

func (p *GetDialogFilterArgs) GetReq() *dialog.TLDialogGetDialogFilter {
	if !p.IsSetReq() {
		return GetDialogFilterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogFilterArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogFilterResult struct {
	Success *dialog.DialogFilterExt
}

var GetDialogFilterResult_Success_DEFAULT *dialog.DialogFilterExt

func (p *GetDialogFilterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogFilterResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogFilterResult) Unmarshal(in []byte) error {
	msg := new(dialog.DialogFilterExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFilterResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogFilterResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogFilterResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.DialogFilterExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFilterResult) GetSuccess() *dialog.DialogFilterExt {
	if !p.IsSetSuccess() {
		return GetDialogFilterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogFilterResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.DialogFilterExt)
}

func (p *GetDialogFilterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogFilterResult) GetResult() interface{} {
	return p.Success
}

func getDialogFilterBySlugHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogFilterBySlugArgs)
	realResult := result.(*GetDialogFilterBySlugResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogFilterBySlug(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogFilterBySlugArgs() interface{} {
	return &GetDialogFilterBySlugArgs{}
}

func newGetDialogFilterBySlugResult() interface{} {
	return &GetDialogFilterBySlugResult{}
}

type GetDialogFilterBySlugArgs struct {
	Req *dialog.TLDialogGetDialogFilterBySlug
}

func (p *GetDialogFilterBySlugArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogFilterBySlugArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogFilterBySlugArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogFilterBySlug)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogFilterBySlugArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogFilterBySlugArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogFilterBySlugArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogFilterBySlug)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogFilterBySlugArgs_Req_DEFAULT *dialog.TLDialogGetDialogFilterBySlug

func (p *GetDialogFilterBySlugArgs) GetReq() *dialog.TLDialogGetDialogFilterBySlug {
	if !p.IsSetReq() {
		return GetDialogFilterBySlugArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogFilterBySlugArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogFilterBySlugResult struct {
	Success *dialog.DialogFilterExt
}

var GetDialogFilterBySlugResult_Success_DEFAULT *dialog.DialogFilterExt

func (p *GetDialogFilterBySlugResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogFilterBySlugResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogFilterBySlugResult) Unmarshal(in []byte) error {
	msg := new(dialog.DialogFilterExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFilterBySlugResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogFilterBySlugResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogFilterBySlugResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.DialogFilterExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFilterBySlugResult) GetSuccess() *dialog.DialogFilterExt {
	if !p.IsSetSuccess() {
		return GetDialogFilterBySlugResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogFilterBySlugResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.DialogFilterExt)
}

func (p *GetDialogFilterBySlugResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogFilterBySlugResult) GetResult() interface{} {
	return p.Success
}

func createDialogFilterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CreateDialogFilterArgs)
	realResult := result.(*CreateDialogFilterResult)
	success, err := handler.(dialog.RPCDialog).DialogCreateDialogFilter(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCreateDialogFilterArgs() interface{} {
	return &CreateDialogFilterArgs{}
}

func newCreateDialogFilterResult() interface{} {
	return &CreateDialogFilterResult{}
}

type CreateDialogFilterArgs struct {
	Req *dialog.TLDialogCreateDialogFilter
}

func (p *CreateDialogFilterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateDialogFilterArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CreateDialogFilterArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogCreateDialogFilter)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CreateDialogFilterArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CreateDialogFilterArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CreateDialogFilterArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogCreateDialogFilter)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CreateDialogFilterArgs_Req_DEFAULT *dialog.TLDialogCreateDialogFilter

func (p *CreateDialogFilterArgs) GetReq() *dialog.TLDialogCreateDialogFilter {
	if !p.IsSetReq() {
		return CreateDialogFilterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateDialogFilterArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateDialogFilterResult struct {
	Success *dialog.DialogFilterExt
}

var CreateDialogFilterResult_Success_DEFAULT *dialog.DialogFilterExt

func (p *CreateDialogFilterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateDialogFilterResult")
	}
	return json.Marshal(p.Success)
}

func (p *CreateDialogFilterResult) Unmarshal(in []byte) error {
	msg := new(dialog.DialogFilterExt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateDialogFilterResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CreateDialogFilterResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CreateDialogFilterResult) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.DialogFilterExt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateDialogFilterResult) GetSuccess() *dialog.DialogFilterExt {
	if !p.IsSetSuccess() {
		return CreateDialogFilterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateDialogFilterResult) SetSuccess(x interface{}) {
	p.Success = x.(*dialog.DialogFilterExt)
}

func (p *CreateDialogFilterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateDialogFilterResult) GetResult() interface{} {
	return p.Success
}

func updateUnreadCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateUnreadCountArgs)
	realResult := result.(*UpdateUnreadCountResult)
	success, err := handler.(dialog.RPCDialog).DialogUpdateUnreadCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateUnreadCountArgs() interface{} {
	return &UpdateUnreadCountArgs{}
}

func newUpdateUnreadCountResult() interface{} {
	return &UpdateUnreadCountResult{}
}

type UpdateUnreadCountArgs struct {
	Req *dialog.TLDialogUpdateUnreadCount
}

func (p *UpdateUnreadCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateUnreadCountArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateUnreadCountArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogUpdateUnreadCount)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateUnreadCountArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateUnreadCountArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateUnreadCountArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogUpdateUnreadCount)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateUnreadCountArgs_Req_DEFAULT *dialog.TLDialogUpdateUnreadCount

func (p *UpdateUnreadCountArgs) GetReq() *dialog.TLDialogUpdateUnreadCount {
	if !p.IsSetReq() {
		return UpdateUnreadCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUnreadCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateUnreadCountResult struct {
	Success *tg.Bool
}

var UpdateUnreadCountResult_Success_DEFAULT *tg.Bool

func (p *UpdateUnreadCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateUnreadCountResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateUnreadCountResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUnreadCountResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateUnreadCountResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateUnreadCountResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUnreadCountResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateUnreadCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUnreadCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateUnreadCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUnreadCountResult) GetResult() interface{} {
	return p.Success
}

func toggleDialogFilterTagsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ToggleDialogFilterTagsArgs)
	realResult := result.(*ToggleDialogFilterTagsResult)
	success, err := handler.(dialog.RPCDialog).DialogToggleDialogFilterTags(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newToggleDialogFilterTagsArgs() interface{} {
	return &ToggleDialogFilterTagsArgs{}
}

func newToggleDialogFilterTagsResult() interface{} {
	return &ToggleDialogFilterTagsResult{}
}

type ToggleDialogFilterTagsArgs struct {
	Req *dialog.TLDialogToggleDialogFilterTags
}

func (p *ToggleDialogFilterTagsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ToggleDialogFilterTagsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ToggleDialogFilterTagsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogToggleDialogFilterTags)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ToggleDialogFilterTagsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ToggleDialogFilterTagsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ToggleDialogFilterTagsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogToggleDialogFilterTags)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ToggleDialogFilterTagsArgs_Req_DEFAULT *dialog.TLDialogToggleDialogFilterTags

func (p *ToggleDialogFilterTagsArgs) GetReq() *dialog.TLDialogToggleDialogFilterTags {
	if !p.IsSetReq() {
		return ToggleDialogFilterTagsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ToggleDialogFilterTagsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ToggleDialogFilterTagsResult struct {
	Success *tg.Bool
}

var ToggleDialogFilterTagsResult_Success_DEFAULT *tg.Bool

func (p *ToggleDialogFilterTagsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ToggleDialogFilterTagsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ToggleDialogFilterTagsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ToggleDialogFilterTagsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ToggleDialogFilterTagsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ToggleDialogFilterTagsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ToggleDialogFilterTagsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ToggleDialogFilterTagsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ToggleDialogFilterTagsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ToggleDialogFilterTagsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ToggleDialogFilterTagsResult) GetResult() interface{} {
	return p.Success
}

func getDialogFilterTagsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDialogFilterTagsArgs)
	realResult := result.(*GetDialogFilterTagsResult)
	success, err := handler.(dialog.RPCDialog).DialogGetDialogFilterTags(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDialogFilterTagsArgs() interface{} {
	return &GetDialogFilterTagsArgs{}
}

func newGetDialogFilterTagsResult() interface{} {
	return &GetDialogFilterTagsResult{}
}

type GetDialogFilterTagsArgs struct {
	Req *dialog.TLDialogGetDialogFilterTags
}

func (p *GetDialogFilterTagsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDialogFilterTagsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDialogFilterTagsArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogGetDialogFilterTags)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDialogFilterTagsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDialogFilterTagsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDialogFilterTagsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogGetDialogFilterTags)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDialogFilterTagsArgs_Req_DEFAULT *dialog.TLDialogGetDialogFilterTags

func (p *GetDialogFilterTagsArgs) GetReq() *dialog.TLDialogGetDialogFilterTags {
	if !p.IsSetReq() {
		return GetDialogFilterTagsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDialogFilterTagsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDialogFilterTagsResult struct {
	Success *tg.Bool
}

var GetDialogFilterTagsResult_Success_DEFAULT *tg.Bool

func (p *GetDialogFilterTagsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDialogFilterTagsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDialogFilterTagsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFilterTagsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDialogFilterTagsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDialogFilterTagsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDialogFilterTagsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return GetDialogFilterTagsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDialogFilterTagsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *GetDialogFilterTagsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDialogFilterTagsResult) GetResult() interface{} {
	return p.Success
}

func setChatWallpaperHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetChatWallpaperArgs)
	realResult := result.(*SetChatWallpaperResult)
	success, err := handler.(dialog.RPCDialog).DialogSetChatWallpaper(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetChatWallpaperArgs() interface{} {
	return &SetChatWallpaperArgs{}
}

func newSetChatWallpaperResult() interface{} {
	return &SetChatWallpaperResult{}
}

type SetChatWallpaperArgs struct {
	Req *dialog.TLDialogSetChatWallpaper
}

func (p *SetChatWallpaperArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetChatWallpaperArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetChatWallpaperArgs) Unmarshal(in []byte) error {
	msg := new(dialog.TLDialogSetChatWallpaper)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetChatWallpaperArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetChatWallpaperArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetChatWallpaperArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dialog.TLDialogSetChatWallpaper)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetChatWallpaperArgs_Req_DEFAULT *dialog.TLDialogSetChatWallpaper

func (p *SetChatWallpaperArgs) GetReq() *dialog.TLDialogSetChatWallpaper {
	if !p.IsSetReq() {
		return SetChatWallpaperArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetChatWallpaperArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetChatWallpaperResult struct {
	Success *tg.Bool
}

var SetChatWallpaperResult_Success_DEFAULT *tg.Bool

func (p *SetChatWallpaperResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetChatWallpaperResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetChatWallpaperResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChatWallpaperResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetChatWallpaperResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetChatWallpaperResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetChatWallpaperResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetChatWallpaperResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetChatWallpaperResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetChatWallpaperResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetChatWallpaperResult) GetResult() interface{} {
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

func (p *kClient) DialogSaveDraftMessage(ctx context.Context, req *dialog.TLDialogSaveDraftMessage) (r *tg.Bool, err error) {
	// var _args SaveDraftMessageArgs
	// _args.Req = req
	// var _result SaveDraftMessageResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.saveDraftMessage", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogClearDraftMessage(ctx context.Context, req *dialog.TLDialogClearDraftMessage) (r *tg.Bool, err error) {
	// var _args ClearDraftMessageArgs
	// _args.Req = req
	// var _result ClearDraftMessageResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.clearDraftMessage", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetAllDrafts(ctx context.Context, req *dialog.TLDialogGetAllDrafts) (r *dialog.VectorPeerWithDraftMessage, err error) {
	// var _args GetAllDraftsArgs
	// _args.Req = req
	// var _result GetAllDraftsResult

	_result := new(dialog.VectorPeerWithDraftMessage)

	if err = p.c.Call(ctx, "dialog.getAllDrafts", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogClearAllDrafts(ctx context.Context, req *dialog.TLDialogClearAllDrafts) (r *dialog.VectorPeerWithDraftMessage, err error) {
	// var _args ClearAllDraftsArgs
	// _args.Req = req
	// var _result ClearAllDraftsResult

	_result := new(dialog.VectorPeerWithDraftMessage)

	if err = p.c.Call(ctx, "dialog.clearAllDrafts", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogMarkDialogUnread(ctx context.Context, req *dialog.TLDialogMarkDialogUnread) (r *tg.Bool, err error) {
	// var _args MarkDialogUnreadArgs
	// _args.Req = req
	// var _result MarkDialogUnreadResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.markDialogUnread", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogToggleDialogPin(ctx context.Context, req *dialog.TLDialogToggleDialogPin) (r *tg.Int32, err error) {
	// var _args ToggleDialogPinArgs
	// _args.Req = req
	// var _result ToggleDialogPinResult

	_result := new(tg.Int32)

	if err = p.c.Call(ctx, "dialog.toggleDialogPin", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogUnreadMarkList(ctx context.Context, req *dialog.TLDialogGetDialogUnreadMarkList) (r *dialog.VectorDialogPeer, err error) {
	// var _args GetDialogUnreadMarkListArgs
	// _args.Req = req
	// var _result GetDialogUnreadMarkListResult

	_result := new(dialog.VectorDialogPeer)

	if err = p.c.Call(ctx, "dialog.getDialogUnreadMarkList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogsByOffsetDate(ctx context.Context, req *dialog.TLDialogGetDialogsByOffsetDate) (r *dialog.VectorDialogExt, err error) {
	// var _args GetDialogsByOffsetDateArgs
	// _args.Req = req
	// var _result GetDialogsByOffsetDateResult

	_result := new(dialog.VectorDialogExt)

	if err = p.c.Call(ctx, "dialog.getDialogsByOffsetDate", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogs(ctx context.Context, req *dialog.TLDialogGetDialogs) (r *dialog.VectorDialogExt, err error) {
	// var _args GetDialogsArgs
	// _args.Req = req
	// var _result GetDialogsResult

	_result := new(dialog.VectorDialogExt)

	if err = p.c.Call(ctx, "dialog.getDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogsByIdList(ctx context.Context, req *dialog.TLDialogGetDialogsByIdList) (r *dialog.VectorDialogExt, err error) {
	// var _args GetDialogsByIdListArgs
	// _args.Req = req
	// var _result GetDialogsByIdListResult

	_result := new(dialog.VectorDialogExt)

	if err = p.c.Call(ctx, "dialog.getDialogsByIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogsCount(ctx context.Context, req *dialog.TLDialogGetDialogsCount) (r *tg.Int32, err error) {
	// var _args GetDialogsCountArgs
	// _args.Req = req
	// var _result GetDialogsCountResult

	_result := new(tg.Int32)

	if err = p.c.Call(ctx, "dialog.getDialogsCount", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetPinnedDialogs(ctx context.Context, req *dialog.TLDialogGetPinnedDialogs) (r *dialog.VectorDialogExt, err error) {
	// var _args GetPinnedDialogsArgs
	// _args.Req = req
	// var _result GetPinnedDialogsResult

	_result := new(dialog.VectorDialogExt)

	if err = p.c.Call(ctx, "dialog.getPinnedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogReorderPinnedDialogs(ctx context.Context, req *dialog.TLDialogReorderPinnedDialogs) (r *tg.Bool, err error) {
	// var _args ReorderPinnedDialogsArgs
	// _args.Req = req
	// var _result ReorderPinnedDialogsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.reorderPinnedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogById(ctx context.Context, req *dialog.TLDialogGetDialogById) (r *dialog.DialogExt, err error) {
	// var _args GetDialogByIdArgs
	// _args.Req = req
	// var _result GetDialogByIdResult

	_result := new(dialog.DialogExt)

	if err = p.c.Call(ctx, "dialog.getDialogById", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetTopMessage(ctx context.Context, req *dialog.TLDialogGetTopMessage) (r *tg.Int32, err error) {
	// var _args GetTopMessageArgs
	// _args.Req = req
	// var _result GetTopMessageResult

	_result := new(tg.Int32)

	if err = p.c.Call(ctx, "dialog.getTopMessage", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogInsertOrUpdateDialog(ctx context.Context, req *dialog.TLDialogInsertOrUpdateDialog) (r *tg.Bool, err error) {
	// var _args InsertOrUpdateDialogArgs
	// _args.Req = req
	// var _result InsertOrUpdateDialogResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.insertOrUpdateDialog", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogDeleteDialog(ctx context.Context, req *dialog.TLDialogDeleteDialog) (r *tg.Bool, err error) {
	// var _args DeleteDialogArgs
	// _args.Req = req
	// var _result DeleteDialogResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.deleteDialog", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetUserPinnedMessage(ctx context.Context, req *dialog.TLDialogGetUserPinnedMessage) (r *tg.Int32, err error) {
	// var _args GetUserPinnedMessageArgs
	// _args.Req = req
	// var _result GetUserPinnedMessageResult

	_result := new(tg.Int32)

	if err = p.c.Call(ctx, "dialog.getUserPinnedMessage", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogUpdateUserPinnedMessage(ctx context.Context, req *dialog.TLDialogUpdateUserPinnedMessage) (r *tg.Bool, err error) {
	// var _args UpdateUserPinnedMessageArgs
	// _args.Req = req
	// var _result UpdateUserPinnedMessageResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.updateUserPinnedMessage", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogInsertOrUpdateDialogFilter(ctx context.Context, req *dialog.TLDialogInsertOrUpdateDialogFilter) (r *tg.Bool, err error) {
	// var _args InsertOrUpdateDialogFilterArgs
	// _args.Req = req
	// var _result InsertOrUpdateDialogFilterResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.insertOrUpdateDialogFilter", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogDeleteDialogFilter(ctx context.Context, req *dialog.TLDialogDeleteDialogFilter) (r *tg.Bool, err error) {
	// var _args DeleteDialogFilterArgs
	// _args.Req = req
	// var _result DeleteDialogFilterResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.deleteDialogFilter", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogUpdateDialogFiltersOrder(ctx context.Context, req *dialog.TLDialogUpdateDialogFiltersOrder) (r *tg.Bool, err error) {
	// var _args UpdateDialogFiltersOrderArgs
	// _args.Req = req
	// var _result UpdateDialogFiltersOrderResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.updateDialogFiltersOrder", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogFilters(ctx context.Context, req *dialog.TLDialogGetDialogFilters) (r *dialog.VectorDialogFilterExt, err error) {
	// var _args GetDialogFiltersArgs
	// _args.Req = req
	// var _result GetDialogFiltersResult

	_result := new(dialog.VectorDialogFilterExt)

	if err = p.c.Call(ctx, "dialog.getDialogFilters", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogFolder(ctx context.Context, req *dialog.TLDialogGetDialogFolder) (r *dialog.VectorDialogExt, err error) {
	// var _args GetDialogFolderArgs
	// _args.Req = req
	// var _result GetDialogFolderResult

	_result := new(dialog.VectorDialogExt)

	if err = p.c.Call(ctx, "dialog.getDialogFolder", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogEditPeerFolders(ctx context.Context, req *dialog.TLDialogEditPeerFolders) (r *dialog.VectorDialogPinnedExt, err error) {
	// var _args EditPeerFoldersArgs
	// _args.Req = req
	// var _result EditPeerFoldersResult

	_result := new(dialog.VectorDialogPinnedExt)

	if err = p.c.Call(ctx, "dialog.editPeerFolders", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetChannelMessageReadParticipants(ctx context.Context, req *dialog.TLDialogGetChannelMessageReadParticipants) (r *dialog.VectorLong, err error) {
	// var _args GetChannelMessageReadParticipantsArgs
	// _args.Req = req
	// var _result GetChannelMessageReadParticipantsResult

	_result := new(dialog.VectorLong)

	if err = p.c.Call(ctx, "dialog.getChannelMessageReadParticipants", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogSetChatTheme(ctx context.Context, req *dialog.TLDialogSetChatTheme) (r *tg.Bool, err error) {
	// var _args SetChatThemeArgs
	// _args.Req = req
	// var _result SetChatThemeResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.setChatTheme", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogSetHistoryTTL(ctx context.Context, req *dialog.TLDialogSetHistoryTTL) (r *tg.Bool, err error) {
	// var _args SetHistoryTTLArgs
	// _args.Req = req
	// var _result SetHistoryTTLResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.setHistoryTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetMyDialogsData(ctx context.Context, req *dialog.TLDialogGetMyDialogsData) (r *dialog.DialogsData, err error) {
	// var _args GetMyDialogsDataArgs
	// _args.Req = req
	// var _result GetMyDialogsDataResult

	_result := new(dialog.DialogsData)

	if err = p.c.Call(ctx, "dialog.getMyDialogsData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetSavedDialogs(ctx context.Context, req *dialog.TLDialogGetSavedDialogs) (r *dialog.SavedDialogList, err error) {
	// var _args GetSavedDialogsArgs
	// _args.Req = req
	// var _result GetSavedDialogsResult

	_result := new(dialog.SavedDialogList)

	if err = p.c.Call(ctx, "dialog.getSavedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetPinnedSavedDialogs(ctx context.Context, req *dialog.TLDialogGetPinnedSavedDialogs) (r *dialog.SavedDialogList, err error) {
	// var _args GetPinnedSavedDialogsArgs
	// _args.Req = req
	// var _result GetPinnedSavedDialogsResult

	_result := new(dialog.SavedDialogList)

	if err = p.c.Call(ctx, "dialog.getPinnedSavedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogToggleSavedDialogPin(ctx context.Context, req *dialog.TLDialogToggleSavedDialogPin) (r *tg.Bool, err error) {
	// var _args ToggleSavedDialogPinArgs
	// _args.Req = req
	// var _result ToggleSavedDialogPinResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.toggleSavedDialogPin", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogReorderPinnedSavedDialogs(ctx context.Context, req *dialog.TLDialogReorderPinnedSavedDialogs) (r *tg.Bool, err error) {
	// var _args ReorderPinnedSavedDialogsArgs
	// _args.Req = req
	// var _result ReorderPinnedSavedDialogsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.reorderPinnedSavedDialogs", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogFilter(ctx context.Context, req *dialog.TLDialogGetDialogFilter) (r *dialog.DialogFilterExt, err error) {
	// var _args GetDialogFilterArgs
	// _args.Req = req
	// var _result GetDialogFilterResult

	_result := new(dialog.DialogFilterExt)

	if err = p.c.Call(ctx, "dialog.getDialogFilter", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogFilterBySlug(ctx context.Context, req *dialog.TLDialogGetDialogFilterBySlug) (r *dialog.DialogFilterExt, err error) {
	// var _args GetDialogFilterBySlugArgs
	// _args.Req = req
	// var _result GetDialogFilterBySlugResult

	_result := new(dialog.DialogFilterExt)

	if err = p.c.Call(ctx, "dialog.getDialogFilterBySlug", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogCreateDialogFilter(ctx context.Context, req *dialog.TLDialogCreateDialogFilter) (r *dialog.DialogFilterExt, err error) {
	// var _args CreateDialogFilterArgs
	// _args.Req = req
	// var _result CreateDialogFilterResult

	_result := new(dialog.DialogFilterExt)

	if err = p.c.Call(ctx, "dialog.createDialogFilter", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogUpdateUnreadCount(ctx context.Context, req *dialog.TLDialogUpdateUnreadCount) (r *tg.Bool, err error) {
	// var _args UpdateUnreadCountArgs
	// _args.Req = req
	// var _result UpdateUnreadCountResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.updateUnreadCount", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogToggleDialogFilterTags(ctx context.Context, req *dialog.TLDialogToggleDialogFilterTags) (r *tg.Bool, err error) {
	// var _args ToggleDialogFilterTagsArgs
	// _args.Req = req
	// var _result ToggleDialogFilterTagsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.toggleDialogFilterTags", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogGetDialogFilterTags(ctx context.Context, req *dialog.TLDialogGetDialogFilterTags) (r *tg.Bool, err error) {
	// var _args GetDialogFilterTagsArgs
	// _args.Req = req
	// var _result GetDialogFilterTagsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.getDialogFilterTags", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DialogSetChatWallpaper(ctx context.Context, req *dialog.TLDialogSetChatWallpaper) (r *tg.Bool, err error) {
	// var _args SetChatWallpaperArgs
	// _args.Req = req
	// var _result SetChatWallpaperResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "dialog.setChatWallpaper", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
