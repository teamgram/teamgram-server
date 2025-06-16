/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package mediaservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"media.uploadPhotoFile": kitex.NewMethodInfo(
		uploadPhotoFileHandler,
		newUploadPhotoFileArgs,
		newUploadPhotoFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.uploadProfilePhotoFile": kitex.NewMethodInfo(
		uploadProfilePhotoFileHandler,
		newUploadProfilePhotoFileArgs,
		newUploadProfilePhotoFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.getPhoto": kitex.NewMethodInfo(
		getPhotoHandler,
		newGetPhotoArgs,
		newGetPhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.getPhotoSizeList": kitex.NewMethodInfo(
		getPhotoSizeListHandler,
		newGetPhotoSizeListArgs,
		newGetPhotoSizeListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.getPhotoSizeListList": kitex.NewMethodInfo(
		getPhotoSizeListListHandler,
		newGetPhotoSizeListListArgs,
		newGetPhotoSizeListListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.getVideoSizeList": kitex.NewMethodInfo(
		getVideoSizeListHandler,
		newGetVideoSizeListArgs,
		newGetVideoSizeListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.uploadedDocumentMedia": kitex.NewMethodInfo(
		uploadedDocumentMediaHandler,
		newUploadedDocumentMediaArgs,
		newUploadedDocumentMediaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.getDocument": kitex.NewMethodInfo(
		getDocumentHandler,
		newGetDocumentArgs,
		newGetDocumentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.getDocumentList": kitex.NewMethodInfo(
		getDocumentListHandler,
		newGetDocumentListArgs,
		newGetDocumentListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.uploadEncryptedFile": kitex.NewMethodInfo(
		uploadEncryptedFileHandler,
		newUploadEncryptedFileArgs,
		newUploadEncryptedFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.getEncryptedFile": kitex.NewMethodInfo(
		getEncryptedFileHandler,
		newGetEncryptedFileArgs,
		newGetEncryptedFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.uploadWallPaperFile": kitex.NewMethodInfo(
		uploadWallPaperFileHandler,
		newUploadWallPaperFileArgs,
		newUploadWallPaperFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.uploadThemeFile": kitex.NewMethodInfo(
		uploadThemeFileHandler,
		newUploadThemeFileArgs,
		newUploadThemeFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.uploadStickerFile": kitex.NewMethodInfo(
		uploadStickerFileHandler,
		newUploadStickerFileArgs,
		newUploadStickerFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.uploadRingtoneFile": kitex.NewMethodInfo(
		uploadRingtoneFileHandler,
		newUploadRingtoneFileArgs,
		newUploadRingtoneFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"media.uploadedProfilePhoto": kitex.NewMethodInfo(
		uploadedProfilePhotoHandler,
		newUploadedProfilePhotoArgs,
		newUploadedProfilePhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	mediaServiceServiceInfo                = NewServiceInfo()
	mediaServiceServiceInfoForClient       = NewServiceInfoForClient()
	mediaServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return mediaServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return mediaServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return mediaServiceServiceInfoForClient
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
	serviceName := "RPCMedia"
	handlerType := (*media.RPCMedia)(nil)
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
		"PackageName": "media",
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

func uploadPhotoFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadPhotoFileArgs)
	realResult := result.(*UploadPhotoFileResult)
	success, err := handler.(media.RPCMedia).MediaUploadPhotoFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadPhotoFileArgs() interface{} {
	return &UploadPhotoFileArgs{}
}

func newUploadPhotoFileResult() interface{} {
	return &UploadPhotoFileResult{}
}

type UploadPhotoFileArgs struct {
	Req *media.TLMediaUploadPhotoFile
}

func (p *UploadPhotoFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadPhotoFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadPhotoFileArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadPhotoFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadPhotoFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadPhotoFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadPhotoFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadPhotoFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadPhotoFileArgs_Req_DEFAULT *media.TLMediaUploadPhotoFile

func (p *UploadPhotoFileArgs) GetReq() *media.TLMediaUploadPhotoFile {
	if !p.IsSetReq() {
		return UploadPhotoFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadPhotoFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadPhotoFileResult struct {
	Success *tg.Photo
}

var UploadPhotoFileResult_Success_DEFAULT *tg.Photo

func (p *UploadPhotoFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadPhotoFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadPhotoFileResult) Unmarshal(in []byte) error {
	msg := new(tg.Photo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadPhotoFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadPhotoFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadPhotoFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Photo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadPhotoFileResult) GetSuccess() *tg.Photo {
	if !p.IsSetSuccess() {
		return UploadPhotoFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadPhotoFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Photo)
}

func (p *UploadPhotoFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadPhotoFileResult) GetResult() interface{} {
	return p.Success
}

func uploadProfilePhotoFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadProfilePhotoFileArgs)
	realResult := result.(*UploadProfilePhotoFileResult)
	success, err := handler.(media.RPCMedia).MediaUploadProfilePhotoFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadProfilePhotoFileArgs() interface{} {
	return &UploadProfilePhotoFileArgs{}
}

func newUploadProfilePhotoFileResult() interface{} {
	return &UploadProfilePhotoFileResult{}
}

type UploadProfilePhotoFileArgs struct {
	Req *media.TLMediaUploadProfilePhotoFile
}

func (p *UploadProfilePhotoFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadProfilePhotoFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadProfilePhotoFileArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadProfilePhotoFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadProfilePhotoFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadProfilePhotoFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadProfilePhotoFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadProfilePhotoFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadProfilePhotoFileArgs_Req_DEFAULT *media.TLMediaUploadProfilePhotoFile

func (p *UploadProfilePhotoFileArgs) GetReq() *media.TLMediaUploadProfilePhotoFile {
	if !p.IsSetReq() {
		return UploadProfilePhotoFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadProfilePhotoFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadProfilePhotoFileResult struct {
	Success *tg.Photo
}

var UploadProfilePhotoFileResult_Success_DEFAULT *tg.Photo

func (p *UploadProfilePhotoFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadProfilePhotoFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadProfilePhotoFileResult) Unmarshal(in []byte) error {
	msg := new(tg.Photo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadProfilePhotoFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadProfilePhotoFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadProfilePhotoFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Photo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadProfilePhotoFileResult) GetSuccess() *tg.Photo {
	if !p.IsSetSuccess() {
		return UploadProfilePhotoFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadProfilePhotoFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Photo)
}

func (p *UploadProfilePhotoFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadProfilePhotoFileResult) GetResult() interface{} {
	return p.Success
}

func getPhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPhotoArgs)
	realResult := result.(*GetPhotoResult)
	success, err := handler.(media.RPCMedia).MediaGetPhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPhotoArgs() interface{} {
	return &GetPhotoArgs{}
}

func newGetPhotoResult() interface{} {
	return &GetPhotoResult{}
}

type GetPhotoArgs struct {
	Req *media.TLMediaGetPhoto
}

func (p *GetPhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPhotoArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaGetPhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaGetPhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPhotoArgs_Req_DEFAULT *media.TLMediaGetPhoto

func (p *GetPhotoArgs) GetReq() *media.TLMediaGetPhoto {
	if !p.IsSetReq() {
		return GetPhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPhotoResult struct {
	Success *tg.Photo
}

var GetPhotoResult_Success_DEFAULT *tg.Photo

func (p *GetPhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPhotoResult) Unmarshal(in []byte) error {
	msg := new(tg.Photo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Photo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhotoResult) GetSuccess() *tg.Photo {
	if !p.IsSetSuccess() {
		return GetPhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Photo)
}

func (p *GetPhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPhotoResult) GetResult() interface{} {
	return p.Success
}

func getPhotoSizeListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPhotoSizeListArgs)
	realResult := result.(*GetPhotoSizeListResult)
	success, err := handler.(media.RPCMedia).MediaGetPhotoSizeList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPhotoSizeListArgs() interface{} {
	return &GetPhotoSizeListArgs{}
}

func newGetPhotoSizeListResult() interface{} {
	return &GetPhotoSizeListResult{}
}

type GetPhotoSizeListArgs struct {
	Req *media.TLMediaGetPhotoSizeList
}

func (p *GetPhotoSizeListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPhotoSizeListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPhotoSizeListArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaGetPhotoSizeList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPhotoSizeListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPhotoSizeListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPhotoSizeListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaGetPhotoSizeList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPhotoSizeListArgs_Req_DEFAULT *media.TLMediaGetPhotoSizeList

func (p *GetPhotoSizeListArgs) GetReq() *media.TLMediaGetPhotoSizeList {
	if !p.IsSetReq() {
		return GetPhotoSizeListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPhotoSizeListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPhotoSizeListResult struct {
	Success *media.PhotoSizeList
}

var GetPhotoSizeListResult_Success_DEFAULT *media.PhotoSizeList

func (p *GetPhotoSizeListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPhotoSizeListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPhotoSizeListResult) Unmarshal(in []byte) error {
	msg := new(media.PhotoSizeList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhotoSizeListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPhotoSizeListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPhotoSizeListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(media.PhotoSizeList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhotoSizeListResult) GetSuccess() *media.PhotoSizeList {
	if !p.IsSetSuccess() {
		return GetPhotoSizeListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPhotoSizeListResult) SetSuccess(x interface{}) {
	p.Success = x.(*media.PhotoSizeList)
}

func (p *GetPhotoSizeListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPhotoSizeListResult) GetResult() interface{} {
	return p.Success
}

func getPhotoSizeListListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPhotoSizeListListArgs)
	realResult := result.(*GetPhotoSizeListListResult)
	success, err := handler.(media.RPCMedia).MediaGetPhotoSizeListList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPhotoSizeListListArgs() interface{} {
	return &GetPhotoSizeListListArgs{}
}

func newGetPhotoSizeListListResult() interface{} {
	return &GetPhotoSizeListListResult{}
}

type GetPhotoSizeListListArgs struct {
	Req *media.TLMediaGetPhotoSizeListList
}

func (p *GetPhotoSizeListListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPhotoSizeListListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPhotoSizeListListArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaGetPhotoSizeListList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPhotoSizeListListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPhotoSizeListListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPhotoSizeListListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaGetPhotoSizeListList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPhotoSizeListListArgs_Req_DEFAULT *media.TLMediaGetPhotoSizeListList

func (p *GetPhotoSizeListListArgs) GetReq() *media.TLMediaGetPhotoSizeListList {
	if !p.IsSetReq() {
		return GetPhotoSizeListListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPhotoSizeListListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPhotoSizeListListResult struct {
	Success *media.VectorPhotoSizeList
}

var GetPhotoSizeListListResult_Success_DEFAULT *media.VectorPhotoSizeList

func (p *GetPhotoSizeListListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPhotoSizeListListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPhotoSizeListListResult) Unmarshal(in []byte) error {
	msg := new(media.VectorPhotoSizeList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhotoSizeListListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPhotoSizeListListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPhotoSizeListListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(media.VectorPhotoSizeList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhotoSizeListListResult) GetSuccess() *media.VectorPhotoSizeList {
	if !p.IsSetSuccess() {
		return GetPhotoSizeListListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPhotoSizeListListResult) SetSuccess(x interface{}) {
	p.Success = x.(*media.VectorPhotoSizeList)
}

func (p *GetPhotoSizeListListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPhotoSizeListListResult) GetResult() interface{} {
	return p.Success
}

func getVideoSizeListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetVideoSizeListArgs)
	realResult := result.(*GetVideoSizeListResult)
	success, err := handler.(media.RPCMedia).MediaGetVideoSizeList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetVideoSizeListArgs() interface{} {
	return &GetVideoSizeListArgs{}
}

func newGetVideoSizeListResult() interface{} {
	return &GetVideoSizeListResult{}
}

type GetVideoSizeListArgs struct {
	Req *media.TLMediaGetVideoSizeList
}

func (p *GetVideoSizeListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetVideoSizeListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetVideoSizeListArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaGetVideoSizeList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetVideoSizeListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetVideoSizeListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetVideoSizeListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaGetVideoSizeList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetVideoSizeListArgs_Req_DEFAULT *media.TLMediaGetVideoSizeList

func (p *GetVideoSizeListArgs) GetReq() *media.TLMediaGetVideoSizeList {
	if !p.IsSetReq() {
		return GetVideoSizeListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetVideoSizeListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetVideoSizeListResult struct {
	Success *media.VideoSizeList
}

var GetVideoSizeListResult_Success_DEFAULT *media.VideoSizeList

func (p *GetVideoSizeListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetVideoSizeListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetVideoSizeListResult) Unmarshal(in []byte) error {
	msg := new(media.VideoSizeList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideoSizeListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetVideoSizeListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetVideoSizeListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(media.VideoSizeList)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideoSizeListResult) GetSuccess() *media.VideoSizeList {
	if !p.IsSetSuccess() {
		return GetVideoSizeListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetVideoSizeListResult) SetSuccess(x interface{}) {
	p.Success = x.(*media.VideoSizeList)
}

func (p *GetVideoSizeListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetVideoSizeListResult) GetResult() interface{} {
	return p.Success
}

func uploadedDocumentMediaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadedDocumentMediaArgs)
	realResult := result.(*UploadedDocumentMediaResult)
	success, err := handler.(media.RPCMedia).MediaUploadedDocumentMedia(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadedDocumentMediaArgs() interface{} {
	return &UploadedDocumentMediaArgs{}
}

func newUploadedDocumentMediaResult() interface{} {
	return &UploadedDocumentMediaResult{}
}

type UploadedDocumentMediaArgs struct {
	Req *media.TLMediaUploadedDocumentMedia
}

func (p *UploadedDocumentMediaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadedDocumentMediaArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadedDocumentMediaArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadedDocumentMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadedDocumentMediaArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadedDocumentMediaArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadedDocumentMediaArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadedDocumentMedia)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadedDocumentMediaArgs_Req_DEFAULT *media.TLMediaUploadedDocumentMedia

func (p *UploadedDocumentMediaArgs) GetReq() *media.TLMediaUploadedDocumentMedia {
	if !p.IsSetReq() {
		return UploadedDocumentMediaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadedDocumentMediaArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadedDocumentMediaResult struct {
	Success *tg.MessageMedia
}

var UploadedDocumentMediaResult_Success_DEFAULT *tg.MessageMedia

func (p *UploadedDocumentMediaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadedDocumentMediaResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadedDocumentMediaResult) Unmarshal(in []byte) error {
	msg := new(tg.MessageMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadedDocumentMediaResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadedDocumentMediaResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadedDocumentMediaResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessageMedia)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadedDocumentMediaResult) GetSuccess() *tg.MessageMedia {
	if !p.IsSetSuccess() {
		return UploadedDocumentMediaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadedDocumentMediaResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessageMedia)
}

func (p *UploadedDocumentMediaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadedDocumentMediaResult) GetResult() interface{} {
	return p.Success
}

func getDocumentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDocumentArgs)
	realResult := result.(*GetDocumentResult)
	success, err := handler.(media.RPCMedia).MediaGetDocument(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDocumentArgs() interface{} {
	return &GetDocumentArgs{}
}

func newGetDocumentResult() interface{} {
	return &GetDocumentResult{}
}

type GetDocumentArgs struct {
	Req *media.TLMediaGetDocument
}

func (p *GetDocumentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDocumentArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDocumentArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaGetDocument)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDocumentArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDocumentArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDocumentArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaGetDocument)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDocumentArgs_Req_DEFAULT *media.TLMediaGetDocument

func (p *GetDocumentArgs) GetReq() *media.TLMediaGetDocument {
	if !p.IsSetReq() {
		return GetDocumentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDocumentArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDocumentResult struct {
	Success *tg.Document
}

var GetDocumentResult_Success_DEFAULT *tg.Document

func (p *GetDocumentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDocumentResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDocumentResult) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDocumentResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDocumentResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDocumentResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDocumentResult) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return GetDocumentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDocumentResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *GetDocumentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDocumentResult) GetResult() interface{} {
	return p.Success
}

func getDocumentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetDocumentListArgs)
	realResult := result.(*GetDocumentListResult)
	success, err := handler.(media.RPCMedia).MediaGetDocumentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetDocumentListArgs() interface{} {
	return &GetDocumentListArgs{}
}

func newGetDocumentListResult() interface{} {
	return &GetDocumentListResult{}
}

type GetDocumentListArgs struct {
	Req *media.TLMediaGetDocumentList
}

func (p *GetDocumentListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDocumentListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetDocumentListArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaGetDocumentList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetDocumentListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetDocumentListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetDocumentListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaGetDocumentList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetDocumentListArgs_Req_DEFAULT *media.TLMediaGetDocumentList

func (p *GetDocumentListArgs) GetReq() *media.TLMediaGetDocumentList {
	if !p.IsSetReq() {
		return GetDocumentListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDocumentListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDocumentListResult struct {
	Success *media.VectorDocument
}

var GetDocumentListResult_Success_DEFAULT *media.VectorDocument

func (p *GetDocumentListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDocumentListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetDocumentListResult) Unmarshal(in []byte) error {
	msg := new(media.VectorDocument)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDocumentListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetDocumentListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetDocumentListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(media.VectorDocument)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDocumentListResult) GetSuccess() *media.VectorDocument {
	if !p.IsSetSuccess() {
		return GetDocumentListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDocumentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*media.VectorDocument)
}

func (p *GetDocumentListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDocumentListResult) GetResult() interface{} {
	return p.Success
}

func uploadEncryptedFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadEncryptedFileArgs)
	realResult := result.(*UploadEncryptedFileResult)
	success, err := handler.(media.RPCMedia).MediaUploadEncryptedFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadEncryptedFileArgs() interface{} {
	return &UploadEncryptedFileArgs{}
}

func newUploadEncryptedFileResult() interface{} {
	return &UploadEncryptedFileResult{}
}

type UploadEncryptedFileArgs struct {
	Req *media.TLMediaUploadEncryptedFile
}

func (p *UploadEncryptedFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadEncryptedFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadEncryptedFileArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadEncryptedFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadEncryptedFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadEncryptedFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadEncryptedFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadEncryptedFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadEncryptedFileArgs_Req_DEFAULT *media.TLMediaUploadEncryptedFile

func (p *UploadEncryptedFileArgs) GetReq() *media.TLMediaUploadEncryptedFile {
	if !p.IsSetReq() {
		return UploadEncryptedFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadEncryptedFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadEncryptedFileResult struct {
	Success *tg.EncryptedFile
}

var UploadEncryptedFileResult_Success_DEFAULT *tg.EncryptedFile

func (p *UploadEncryptedFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadEncryptedFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadEncryptedFileResult) Unmarshal(in []byte) error {
	msg := new(tg.EncryptedFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadEncryptedFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadEncryptedFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadEncryptedFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.EncryptedFile)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadEncryptedFileResult) GetSuccess() *tg.EncryptedFile {
	if !p.IsSetSuccess() {
		return UploadEncryptedFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadEncryptedFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.EncryptedFile)
}

func (p *UploadEncryptedFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadEncryptedFileResult) GetResult() interface{} {
	return p.Success
}

func getEncryptedFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetEncryptedFileArgs)
	realResult := result.(*GetEncryptedFileResult)
	success, err := handler.(media.RPCMedia).MediaGetEncryptedFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetEncryptedFileArgs() interface{} {
	return &GetEncryptedFileArgs{}
}

func newGetEncryptedFileResult() interface{} {
	return &GetEncryptedFileResult{}
}

type GetEncryptedFileArgs struct {
	Req *media.TLMediaGetEncryptedFile
}

func (p *GetEncryptedFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetEncryptedFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetEncryptedFileArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaGetEncryptedFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetEncryptedFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetEncryptedFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetEncryptedFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaGetEncryptedFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetEncryptedFileArgs_Req_DEFAULT *media.TLMediaGetEncryptedFile

func (p *GetEncryptedFileArgs) GetReq() *media.TLMediaGetEncryptedFile {
	if !p.IsSetReq() {
		return GetEncryptedFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetEncryptedFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetEncryptedFileResult struct {
	Success *tg.EncryptedFile
}

var GetEncryptedFileResult_Success_DEFAULT *tg.EncryptedFile

func (p *GetEncryptedFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetEncryptedFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetEncryptedFileResult) Unmarshal(in []byte) error {
	msg := new(tg.EncryptedFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetEncryptedFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetEncryptedFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetEncryptedFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.EncryptedFile)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetEncryptedFileResult) GetSuccess() *tg.EncryptedFile {
	if !p.IsSetSuccess() {
		return GetEncryptedFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetEncryptedFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.EncryptedFile)
}

func (p *GetEncryptedFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetEncryptedFileResult) GetResult() interface{} {
	return p.Success
}

func uploadWallPaperFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadWallPaperFileArgs)
	realResult := result.(*UploadWallPaperFileResult)
	success, err := handler.(media.RPCMedia).MediaUploadWallPaperFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadWallPaperFileArgs() interface{} {
	return &UploadWallPaperFileArgs{}
}

func newUploadWallPaperFileResult() interface{} {
	return &UploadWallPaperFileResult{}
}

type UploadWallPaperFileArgs struct {
	Req *media.TLMediaUploadWallPaperFile
}

func (p *UploadWallPaperFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadWallPaperFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadWallPaperFileArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadWallPaperFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadWallPaperFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadWallPaperFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadWallPaperFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadWallPaperFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadWallPaperFileArgs_Req_DEFAULT *media.TLMediaUploadWallPaperFile

func (p *UploadWallPaperFileArgs) GetReq() *media.TLMediaUploadWallPaperFile {
	if !p.IsSetReq() {
		return UploadWallPaperFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadWallPaperFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadWallPaperFileResult struct {
	Success *tg.Document
}

var UploadWallPaperFileResult_Success_DEFAULT *tg.Document

func (p *UploadWallPaperFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadWallPaperFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadWallPaperFileResult) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadWallPaperFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadWallPaperFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadWallPaperFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadWallPaperFileResult) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return UploadWallPaperFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadWallPaperFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *UploadWallPaperFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadWallPaperFileResult) GetResult() interface{} {
	return p.Success
}

func uploadThemeFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadThemeFileArgs)
	realResult := result.(*UploadThemeFileResult)
	success, err := handler.(media.RPCMedia).MediaUploadThemeFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadThemeFileArgs() interface{} {
	return &UploadThemeFileArgs{}
}

func newUploadThemeFileResult() interface{} {
	return &UploadThemeFileResult{}
}

type UploadThemeFileArgs struct {
	Req *media.TLMediaUploadThemeFile
}

func (p *UploadThemeFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadThemeFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadThemeFileArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadThemeFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadThemeFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadThemeFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadThemeFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadThemeFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadThemeFileArgs_Req_DEFAULT *media.TLMediaUploadThemeFile

func (p *UploadThemeFileArgs) GetReq() *media.TLMediaUploadThemeFile {
	if !p.IsSetReq() {
		return UploadThemeFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadThemeFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadThemeFileResult struct {
	Success *tg.Document
}

var UploadThemeFileResult_Success_DEFAULT *tg.Document

func (p *UploadThemeFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadThemeFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadThemeFileResult) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadThemeFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadThemeFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadThemeFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadThemeFileResult) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return UploadThemeFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadThemeFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *UploadThemeFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadThemeFileResult) GetResult() interface{} {
	return p.Success
}

func uploadStickerFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadStickerFileArgs)
	realResult := result.(*UploadStickerFileResult)
	success, err := handler.(media.RPCMedia).MediaUploadStickerFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadStickerFileArgs() interface{} {
	return &UploadStickerFileArgs{}
}

func newUploadStickerFileResult() interface{} {
	return &UploadStickerFileResult{}
}

type UploadStickerFileArgs struct {
	Req *media.TLMediaUploadStickerFile
}

func (p *UploadStickerFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadStickerFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadStickerFileArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadStickerFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadStickerFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadStickerFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadStickerFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadStickerFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadStickerFileArgs_Req_DEFAULT *media.TLMediaUploadStickerFile

func (p *UploadStickerFileArgs) GetReq() *media.TLMediaUploadStickerFile {
	if !p.IsSetReq() {
		return UploadStickerFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadStickerFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadStickerFileResult struct {
	Success *tg.Document
}

var UploadStickerFileResult_Success_DEFAULT *tg.Document

func (p *UploadStickerFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadStickerFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadStickerFileResult) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadStickerFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadStickerFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadStickerFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadStickerFileResult) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return UploadStickerFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadStickerFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *UploadStickerFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadStickerFileResult) GetResult() interface{} {
	return p.Success
}

func uploadRingtoneFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadRingtoneFileArgs)
	realResult := result.(*UploadRingtoneFileResult)
	success, err := handler.(media.RPCMedia).MediaUploadRingtoneFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadRingtoneFileArgs() interface{} {
	return &UploadRingtoneFileArgs{}
}

func newUploadRingtoneFileResult() interface{} {
	return &UploadRingtoneFileResult{}
}

type UploadRingtoneFileArgs struct {
	Req *media.TLMediaUploadRingtoneFile
}

func (p *UploadRingtoneFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadRingtoneFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadRingtoneFileArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadRingtoneFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadRingtoneFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadRingtoneFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadRingtoneFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadRingtoneFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadRingtoneFileArgs_Req_DEFAULT *media.TLMediaUploadRingtoneFile

func (p *UploadRingtoneFileArgs) GetReq() *media.TLMediaUploadRingtoneFile {
	if !p.IsSetReq() {
		return UploadRingtoneFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadRingtoneFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadRingtoneFileResult struct {
	Success *tg.Document
}

var UploadRingtoneFileResult_Success_DEFAULT *tg.Document

func (p *UploadRingtoneFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadRingtoneFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadRingtoneFileResult) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadRingtoneFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadRingtoneFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadRingtoneFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadRingtoneFileResult) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return UploadRingtoneFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadRingtoneFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *UploadRingtoneFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadRingtoneFileResult) GetResult() interface{} {
	return p.Success
}

func uploadedProfilePhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadedProfilePhotoArgs)
	realResult := result.(*UploadedProfilePhotoResult)
	success, err := handler.(media.RPCMedia).MediaUploadedProfilePhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadedProfilePhotoArgs() interface{} {
	return &UploadedProfilePhotoArgs{}
}

func newUploadedProfilePhotoResult() interface{} {
	return &UploadedProfilePhotoResult{}
}

type UploadedProfilePhotoArgs struct {
	Req *media.TLMediaUploadedProfilePhoto
}

func (p *UploadedProfilePhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadedProfilePhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadedProfilePhotoArgs) Unmarshal(in []byte) error {
	msg := new(media.TLMediaUploadedProfilePhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadedProfilePhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadedProfilePhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadedProfilePhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(media.TLMediaUploadedProfilePhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadedProfilePhotoArgs_Req_DEFAULT *media.TLMediaUploadedProfilePhoto

func (p *UploadedProfilePhotoArgs) GetReq() *media.TLMediaUploadedProfilePhoto {
	if !p.IsSetReq() {
		return UploadedProfilePhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadedProfilePhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadedProfilePhotoResult struct {
	Success *tg.Photo
}

var UploadedProfilePhotoResult_Success_DEFAULT *tg.Photo

func (p *UploadedProfilePhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadedProfilePhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadedProfilePhotoResult) Unmarshal(in []byte) error {
	msg := new(tg.Photo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadedProfilePhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadedProfilePhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadedProfilePhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Photo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadedProfilePhotoResult) GetSuccess() *tg.Photo {
	if !p.IsSetSuccess() {
		return UploadedProfilePhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadedProfilePhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Photo)
}

func (p *UploadedProfilePhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadedProfilePhotoResult) GetResult() interface{} {
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

func (p *kClient) MediaUploadPhotoFile(ctx context.Context, req *media.TLMediaUploadPhotoFile) (r *tg.Photo, err error) {
	var _args UploadPhotoFileArgs
	_args.Req = req
	var _result UploadPhotoFileResult
	if err = p.c.Call(ctx, "media.uploadPhotoFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaUploadProfilePhotoFile(ctx context.Context, req *media.TLMediaUploadProfilePhotoFile) (r *tg.Photo, err error) {
	var _args UploadProfilePhotoFileArgs
	_args.Req = req
	var _result UploadProfilePhotoFileResult
	if err = p.c.Call(ctx, "media.uploadProfilePhotoFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaGetPhoto(ctx context.Context, req *media.TLMediaGetPhoto) (r *tg.Photo, err error) {
	var _args GetPhotoArgs
	_args.Req = req
	var _result GetPhotoResult
	if err = p.c.Call(ctx, "media.getPhoto", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaGetPhotoSizeList(ctx context.Context, req *media.TLMediaGetPhotoSizeList) (r *media.PhotoSizeList, err error) {
	var _args GetPhotoSizeListArgs
	_args.Req = req
	var _result GetPhotoSizeListResult
	if err = p.c.Call(ctx, "media.getPhotoSizeList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaGetPhotoSizeListList(ctx context.Context, req *media.TLMediaGetPhotoSizeListList) (r *media.VectorPhotoSizeList, err error) {
	var _args GetPhotoSizeListListArgs
	_args.Req = req
	var _result GetPhotoSizeListListResult
	if err = p.c.Call(ctx, "media.getPhotoSizeListList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaGetVideoSizeList(ctx context.Context, req *media.TLMediaGetVideoSizeList) (r *media.VideoSizeList, err error) {
	var _args GetVideoSizeListArgs
	_args.Req = req
	var _result GetVideoSizeListResult
	if err = p.c.Call(ctx, "media.getVideoSizeList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaUploadedDocumentMedia(ctx context.Context, req *media.TLMediaUploadedDocumentMedia) (r *tg.MessageMedia, err error) {
	var _args UploadedDocumentMediaArgs
	_args.Req = req
	var _result UploadedDocumentMediaResult
	if err = p.c.Call(ctx, "media.uploadedDocumentMedia", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaGetDocument(ctx context.Context, req *media.TLMediaGetDocument) (r *tg.Document, err error) {
	var _args GetDocumentArgs
	_args.Req = req
	var _result GetDocumentResult
	if err = p.c.Call(ctx, "media.getDocument", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaGetDocumentList(ctx context.Context, req *media.TLMediaGetDocumentList) (r *media.VectorDocument, err error) {
	var _args GetDocumentListArgs
	_args.Req = req
	var _result GetDocumentListResult
	if err = p.c.Call(ctx, "media.getDocumentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaUploadEncryptedFile(ctx context.Context, req *media.TLMediaUploadEncryptedFile) (r *tg.EncryptedFile, err error) {
	var _args UploadEncryptedFileArgs
	_args.Req = req
	var _result UploadEncryptedFileResult
	if err = p.c.Call(ctx, "media.uploadEncryptedFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaGetEncryptedFile(ctx context.Context, req *media.TLMediaGetEncryptedFile) (r *tg.EncryptedFile, err error) {
	var _args GetEncryptedFileArgs
	_args.Req = req
	var _result GetEncryptedFileResult
	if err = p.c.Call(ctx, "media.getEncryptedFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaUploadWallPaperFile(ctx context.Context, req *media.TLMediaUploadWallPaperFile) (r *tg.Document, err error) {
	var _args UploadWallPaperFileArgs
	_args.Req = req
	var _result UploadWallPaperFileResult
	if err = p.c.Call(ctx, "media.uploadWallPaperFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaUploadThemeFile(ctx context.Context, req *media.TLMediaUploadThemeFile) (r *tg.Document, err error) {
	var _args UploadThemeFileArgs
	_args.Req = req
	var _result UploadThemeFileResult
	if err = p.c.Call(ctx, "media.uploadThemeFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaUploadStickerFile(ctx context.Context, req *media.TLMediaUploadStickerFile) (r *tg.Document, err error) {
	var _args UploadStickerFileArgs
	_args.Req = req
	var _result UploadStickerFileResult
	if err = p.c.Call(ctx, "media.uploadStickerFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaUploadRingtoneFile(ctx context.Context, req *media.TLMediaUploadRingtoneFile) (r *tg.Document, err error) {
	var _args UploadRingtoneFileArgs
	_args.Req = req
	var _result UploadRingtoneFileResult
	if err = p.c.Call(ctx, "media.uploadRingtoneFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MediaUploadedProfilePhoto(ctx context.Context, req *media.TLMediaUploadedProfilePhoto) (r *tg.Photo, err error) {
	var _args UploadedProfilePhotoArgs
	_args.Req = req
	var _result UploadedProfilePhotoResult
	if err = p.c.Call(ctx, "media.uploadedProfilePhoto", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
