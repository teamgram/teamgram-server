/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dfsservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/dfs.RPCDfs/dfs.writeFilePartData": kitex.NewMethodInfo(
		writeFilePartDataHandler,
		newWriteFilePartDataArgs,
		newWriteFilePartDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadPhotoFileV2": kitex.NewMethodInfo(
		uploadPhotoFileV2Handler,
		newUploadPhotoFileV2Args,
		newUploadPhotoFileV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadProfilePhotoFileV2": kitex.NewMethodInfo(
		uploadProfilePhotoFileV2Handler,
		newUploadProfilePhotoFileV2Args,
		newUploadProfilePhotoFileV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadEncryptedFileV2": kitex.NewMethodInfo(
		uploadEncryptedFileV2Handler,
		newUploadEncryptedFileV2Args,
		newUploadEncryptedFileV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.downloadFile": kitex.NewMethodInfo(
		downloadFileHandler,
		newDownloadFileArgs,
		newDownloadFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadDocumentFileV2": kitex.NewMethodInfo(
		uploadDocumentFileV2Handler,
		newUploadDocumentFileV2Args,
		newUploadDocumentFileV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadGifDocumentMedia": kitex.NewMethodInfo(
		uploadGifDocumentMediaHandler,
		newUploadGifDocumentMediaArgs,
		newUploadGifDocumentMediaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadMp4DocumentMedia": kitex.NewMethodInfo(
		uploadMp4DocumentMediaHandler,
		newUploadMp4DocumentMediaArgs,
		newUploadMp4DocumentMediaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadWallPaperFile": kitex.NewMethodInfo(
		uploadWallPaperFileHandler,
		newUploadWallPaperFileArgs,
		newUploadWallPaperFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadThemeFile": kitex.NewMethodInfo(
		uploadThemeFileHandler,
		newUploadThemeFileArgs,
		newUploadThemeFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadRingtoneFile": kitex.NewMethodInfo(
		uploadRingtoneFileHandler,
		newUploadRingtoneFileArgs,
		newUploadRingtoneFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/dfs.RPCDfs/dfs.uploadedProfilePhoto": kitex.NewMethodInfo(
		uploadedProfilePhotoHandler,
		newUploadedProfilePhotoArgs,
		newUploadedProfilePhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	dfsServiceServiceInfo                = NewServiceInfo()
	dfsServiceServiceInfoForClient       = NewServiceInfoForClient()
	dfsServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCDfs", dfsServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCDfs", dfsServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCDfs", dfsServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return dfsServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return dfsServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return dfsServiceServiceInfoForClient
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
	serviceName := "RPCDfs"
	handlerType := (*dfs.RPCDfs)(nil)
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
		"PackageName": "dfs",
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

func writeFilePartDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*WriteFilePartDataArgs)
	realResult := result.(*WriteFilePartDataResult)
	success, err := handler.(dfs.RPCDfs).DfsWriteFilePartData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newWriteFilePartDataArgs() interface{} {
	return &WriteFilePartDataArgs{}
}

func newWriteFilePartDataResult() interface{} {
	return &WriteFilePartDataResult{}
}

type WriteFilePartDataArgs struct {
	Req *dfs.TLDfsWriteFilePartData
}

func (p *WriteFilePartDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in WriteFilePartDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *WriteFilePartDataArgs) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsWriteFilePartData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *WriteFilePartDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in WriteFilePartDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *WriteFilePartDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dfs.TLDfsWriteFilePartData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var WriteFilePartDataArgs_Req_DEFAULT *dfs.TLDfsWriteFilePartData

func (p *WriteFilePartDataArgs) GetReq() *dfs.TLDfsWriteFilePartData {
	if !p.IsSetReq() {
		return WriteFilePartDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *WriteFilePartDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type WriteFilePartDataResult struct {
	Success *tg.Bool
}

var WriteFilePartDataResult_Success_DEFAULT *tg.Bool

func (p *WriteFilePartDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in WriteFilePartDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *WriteFilePartDataResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *WriteFilePartDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in WriteFilePartDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *WriteFilePartDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *WriteFilePartDataResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return WriteFilePartDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *WriteFilePartDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *WriteFilePartDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *WriteFilePartDataResult) GetResult() interface{} {
	return p.Success
}

func uploadPhotoFileV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadPhotoFileV2Args)
	realResult := result.(*UploadPhotoFileV2Result)
	success, err := handler.(dfs.RPCDfs).DfsUploadPhotoFileV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadPhotoFileV2Args() interface{} {
	return &UploadPhotoFileV2Args{}
}

func newUploadPhotoFileV2Result() interface{} {
	return &UploadPhotoFileV2Result{}
}

type UploadPhotoFileV2Args struct {
	Req *dfs.TLDfsUploadPhotoFileV2
}

func (p *UploadPhotoFileV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadPhotoFileV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *UploadPhotoFileV2Args) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadPhotoFileV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadPhotoFileV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadPhotoFileV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadPhotoFileV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(dfs.TLDfsUploadPhotoFileV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadPhotoFileV2Args_Req_DEFAULT *dfs.TLDfsUploadPhotoFileV2

func (p *UploadPhotoFileV2Args) GetReq() *dfs.TLDfsUploadPhotoFileV2 {
	if !p.IsSetReq() {
		return UploadPhotoFileV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadPhotoFileV2Args) IsSetReq() bool {
	return p.Req != nil
}

type UploadPhotoFileV2Result struct {
	Success *tg.Photo
}

var UploadPhotoFileV2Result_Success_DEFAULT *tg.Photo

func (p *UploadPhotoFileV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadPhotoFileV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *UploadPhotoFileV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Photo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadPhotoFileV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadPhotoFileV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadPhotoFileV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Photo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadPhotoFileV2Result) GetSuccess() *tg.Photo {
	if !p.IsSetSuccess() {
		return UploadPhotoFileV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadPhotoFileV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Photo)
}

func (p *UploadPhotoFileV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadPhotoFileV2Result) GetResult() interface{} {
	return p.Success
}

func uploadProfilePhotoFileV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadProfilePhotoFileV2Args)
	realResult := result.(*UploadProfilePhotoFileV2Result)
	success, err := handler.(dfs.RPCDfs).DfsUploadProfilePhotoFileV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadProfilePhotoFileV2Args() interface{} {
	return &UploadProfilePhotoFileV2Args{}
}

func newUploadProfilePhotoFileV2Result() interface{} {
	return &UploadProfilePhotoFileV2Result{}
}

type UploadProfilePhotoFileV2Args struct {
	Req *dfs.TLDfsUploadProfilePhotoFileV2
}

func (p *UploadProfilePhotoFileV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadProfilePhotoFileV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *UploadProfilePhotoFileV2Args) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadProfilePhotoFileV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadProfilePhotoFileV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadProfilePhotoFileV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadProfilePhotoFileV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(dfs.TLDfsUploadProfilePhotoFileV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadProfilePhotoFileV2Args_Req_DEFAULT *dfs.TLDfsUploadProfilePhotoFileV2

func (p *UploadProfilePhotoFileV2Args) GetReq() *dfs.TLDfsUploadProfilePhotoFileV2 {
	if !p.IsSetReq() {
		return UploadProfilePhotoFileV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadProfilePhotoFileV2Args) IsSetReq() bool {
	return p.Req != nil
}

type UploadProfilePhotoFileV2Result struct {
	Success *tg.Photo
}

var UploadProfilePhotoFileV2Result_Success_DEFAULT *tg.Photo

func (p *UploadProfilePhotoFileV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadProfilePhotoFileV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *UploadProfilePhotoFileV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Photo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadProfilePhotoFileV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadProfilePhotoFileV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadProfilePhotoFileV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Photo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadProfilePhotoFileV2Result) GetSuccess() *tg.Photo {
	if !p.IsSetSuccess() {
		return UploadProfilePhotoFileV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadProfilePhotoFileV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Photo)
}

func (p *UploadProfilePhotoFileV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadProfilePhotoFileV2Result) GetResult() interface{} {
	return p.Success
}

func uploadEncryptedFileV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadEncryptedFileV2Args)
	realResult := result.(*UploadEncryptedFileV2Result)
	success, err := handler.(dfs.RPCDfs).DfsUploadEncryptedFileV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadEncryptedFileV2Args() interface{} {
	return &UploadEncryptedFileV2Args{}
}

func newUploadEncryptedFileV2Result() interface{} {
	return &UploadEncryptedFileV2Result{}
}

type UploadEncryptedFileV2Args struct {
	Req *dfs.TLDfsUploadEncryptedFileV2
}

func (p *UploadEncryptedFileV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadEncryptedFileV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *UploadEncryptedFileV2Args) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadEncryptedFileV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadEncryptedFileV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadEncryptedFileV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadEncryptedFileV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(dfs.TLDfsUploadEncryptedFileV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadEncryptedFileV2Args_Req_DEFAULT *dfs.TLDfsUploadEncryptedFileV2

func (p *UploadEncryptedFileV2Args) GetReq() *dfs.TLDfsUploadEncryptedFileV2 {
	if !p.IsSetReq() {
		return UploadEncryptedFileV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadEncryptedFileV2Args) IsSetReq() bool {
	return p.Req != nil
}

type UploadEncryptedFileV2Result struct {
	Success *tg.EncryptedFile
}

var UploadEncryptedFileV2Result_Success_DEFAULT *tg.EncryptedFile

func (p *UploadEncryptedFileV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadEncryptedFileV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *UploadEncryptedFileV2Result) Unmarshal(in []byte) error {
	msg := new(tg.EncryptedFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadEncryptedFileV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadEncryptedFileV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadEncryptedFileV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.EncryptedFile)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadEncryptedFileV2Result) GetSuccess() *tg.EncryptedFile {
	if !p.IsSetSuccess() {
		return UploadEncryptedFileV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadEncryptedFileV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.EncryptedFile)
}

func (p *UploadEncryptedFileV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadEncryptedFileV2Result) GetResult() interface{} {
	return p.Success
}

func downloadFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DownloadFileArgs)
	realResult := result.(*DownloadFileResult)
	success, err := handler.(dfs.RPCDfs).DfsDownloadFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDownloadFileArgs() interface{} {
	return &DownloadFileArgs{}
}

func newDownloadFileResult() interface{} {
	return &DownloadFileResult{}
}

type DownloadFileArgs struct {
	Req *dfs.TLDfsDownloadFile
}

func (p *DownloadFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DownloadFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DownloadFileArgs) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsDownloadFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DownloadFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DownloadFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DownloadFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dfs.TLDfsDownloadFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DownloadFileArgs_Req_DEFAULT *dfs.TLDfsDownloadFile

func (p *DownloadFileArgs) GetReq() *dfs.TLDfsDownloadFile {
	if !p.IsSetReq() {
		return DownloadFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DownloadFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type DownloadFileResult struct {
	Success *tg.UploadFile
}

var DownloadFileResult_Success_DEFAULT *tg.UploadFile

func (p *DownloadFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DownloadFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *DownloadFileResult) Unmarshal(in []byte) error {
	msg := new(tg.UploadFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DownloadFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DownloadFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DownloadFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UploadFile)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DownloadFileResult) GetSuccess() *tg.UploadFile {
	if !p.IsSetSuccess() {
		return DownloadFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DownloadFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UploadFile)
}

func (p *DownloadFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DownloadFileResult) GetResult() interface{} {
	return p.Success
}

func uploadDocumentFileV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadDocumentFileV2Args)
	realResult := result.(*UploadDocumentFileV2Result)
	success, err := handler.(dfs.RPCDfs).DfsUploadDocumentFileV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadDocumentFileV2Args() interface{} {
	return &UploadDocumentFileV2Args{}
}

func newUploadDocumentFileV2Result() interface{} {
	return &UploadDocumentFileV2Result{}
}

type UploadDocumentFileV2Args struct {
	Req *dfs.TLDfsUploadDocumentFileV2
}

func (p *UploadDocumentFileV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadDocumentFileV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *UploadDocumentFileV2Args) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadDocumentFileV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadDocumentFileV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadDocumentFileV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadDocumentFileV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(dfs.TLDfsUploadDocumentFileV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadDocumentFileV2Args_Req_DEFAULT *dfs.TLDfsUploadDocumentFileV2

func (p *UploadDocumentFileV2Args) GetReq() *dfs.TLDfsUploadDocumentFileV2 {
	if !p.IsSetReq() {
		return UploadDocumentFileV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadDocumentFileV2Args) IsSetReq() bool {
	return p.Req != nil
}

type UploadDocumentFileV2Result struct {
	Success *tg.Document
}

var UploadDocumentFileV2Result_Success_DEFAULT *tg.Document

func (p *UploadDocumentFileV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadDocumentFileV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *UploadDocumentFileV2Result) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadDocumentFileV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadDocumentFileV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadDocumentFileV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadDocumentFileV2Result) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return UploadDocumentFileV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadDocumentFileV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *UploadDocumentFileV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadDocumentFileV2Result) GetResult() interface{} {
	return p.Success
}

func uploadGifDocumentMediaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadGifDocumentMediaArgs)
	realResult := result.(*UploadGifDocumentMediaResult)
	success, err := handler.(dfs.RPCDfs).DfsUploadGifDocumentMedia(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadGifDocumentMediaArgs() interface{} {
	return &UploadGifDocumentMediaArgs{}
}

func newUploadGifDocumentMediaResult() interface{} {
	return &UploadGifDocumentMediaResult{}
}

type UploadGifDocumentMediaArgs struct {
	Req *dfs.TLDfsUploadGifDocumentMedia
}

func (p *UploadGifDocumentMediaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadGifDocumentMediaArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadGifDocumentMediaArgs) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadGifDocumentMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadGifDocumentMediaArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadGifDocumentMediaArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadGifDocumentMediaArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dfs.TLDfsUploadGifDocumentMedia)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadGifDocumentMediaArgs_Req_DEFAULT *dfs.TLDfsUploadGifDocumentMedia

func (p *UploadGifDocumentMediaArgs) GetReq() *dfs.TLDfsUploadGifDocumentMedia {
	if !p.IsSetReq() {
		return UploadGifDocumentMediaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadGifDocumentMediaArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadGifDocumentMediaResult struct {
	Success *tg.Document
}

var UploadGifDocumentMediaResult_Success_DEFAULT *tg.Document

func (p *UploadGifDocumentMediaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadGifDocumentMediaResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadGifDocumentMediaResult) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGifDocumentMediaResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadGifDocumentMediaResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadGifDocumentMediaResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGifDocumentMediaResult) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return UploadGifDocumentMediaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadGifDocumentMediaResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *UploadGifDocumentMediaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadGifDocumentMediaResult) GetResult() interface{} {
	return p.Success
}

func uploadMp4DocumentMediaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadMp4DocumentMediaArgs)
	realResult := result.(*UploadMp4DocumentMediaResult)
	success, err := handler.(dfs.RPCDfs).DfsUploadMp4DocumentMedia(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadMp4DocumentMediaArgs() interface{} {
	return &UploadMp4DocumentMediaArgs{}
}

func newUploadMp4DocumentMediaResult() interface{} {
	return &UploadMp4DocumentMediaResult{}
}

type UploadMp4DocumentMediaArgs struct {
	Req *dfs.TLDfsUploadMp4DocumentMedia
}

func (p *UploadMp4DocumentMediaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadMp4DocumentMediaArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadMp4DocumentMediaArgs) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadMp4DocumentMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadMp4DocumentMediaArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadMp4DocumentMediaArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadMp4DocumentMediaArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(dfs.TLDfsUploadMp4DocumentMedia)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadMp4DocumentMediaArgs_Req_DEFAULT *dfs.TLDfsUploadMp4DocumentMedia

func (p *UploadMp4DocumentMediaArgs) GetReq() *dfs.TLDfsUploadMp4DocumentMedia {
	if !p.IsSetReq() {
		return UploadMp4DocumentMediaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadMp4DocumentMediaArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadMp4DocumentMediaResult struct {
	Success *tg.Document
}

var UploadMp4DocumentMediaResult_Success_DEFAULT *tg.Document

func (p *UploadMp4DocumentMediaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadMp4DocumentMediaResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadMp4DocumentMediaResult) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadMp4DocumentMediaResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadMp4DocumentMediaResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadMp4DocumentMediaResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadMp4DocumentMediaResult) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return UploadMp4DocumentMediaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadMp4DocumentMediaResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *UploadMp4DocumentMediaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadMp4DocumentMediaResult) GetResult() interface{} {
	return p.Success
}

func uploadWallPaperFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadWallPaperFileArgs)
	realResult := result.(*UploadWallPaperFileResult)
	success, err := handler.(dfs.RPCDfs).DfsUploadWallPaperFile(ctx, realArg.Req)
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
	Req *dfs.TLDfsUploadWallPaperFile
}

func (p *UploadWallPaperFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadWallPaperFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadWallPaperFileArgs) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadWallPaperFile)
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
	msg := new(dfs.TLDfsUploadWallPaperFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadWallPaperFileArgs_Req_DEFAULT *dfs.TLDfsUploadWallPaperFile

func (p *UploadWallPaperFileArgs) GetReq() *dfs.TLDfsUploadWallPaperFile {
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
	success, err := handler.(dfs.RPCDfs).DfsUploadThemeFile(ctx, realArg.Req)
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
	Req *dfs.TLDfsUploadThemeFile
}

func (p *UploadThemeFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadThemeFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadThemeFileArgs) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadThemeFile)
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
	msg := new(dfs.TLDfsUploadThemeFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadThemeFileArgs_Req_DEFAULT *dfs.TLDfsUploadThemeFile

func (p *UploadThemeFileArgs) GetReq() *dfs.TLDfsUploadThemeFile {
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

func uploadRingtoneFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadRingtoneFileArgs)
	realResult := result.(*UploadRingtoneFileResult)
	success, err := handler.(dfs.RPCDfs).DfsUploadRingtoneFile(ctx, realArg.Req)
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
	Req *dfs.TLDfsUploadRingtoneFile
}

func (p *UploadRingtoneFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadRingtoneFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadRingtoneFileArgs) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadRingtoneFile)
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
	msg := new(dfs.TLDfsUploadRingtoneFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadRingtoneFileArgs_Req_DEFAULT *dfs.TLDfsUploadRingtoneFile

func (p *UploadRingtoneFileArgs) GetReq() *dfs.TLDfsUploadRingtoneFile {
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
	success, err := handler.(dfs.RPCDfs).DfsUploadedProfilePhoto(ctx, realArg.Req)
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
	Req *dfs.TLDfsUploadedProfilePhoto
}

func (p *UploadedProfilePhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadedProfilePhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadedProfilePhotoArgs) Unmarshal(in []byte) error {
	msg := new(dfs.TLDfsUploadedProfilePhoto)
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
	msg := new(dfs.TLDfsUploadedProfilePhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadedProfilePhotoArgs_Req_DEFAULT *dfs.TLDfsUploadedProfilePhoto

func (p *UploadedProfilePhotoArgs) GetReq() *dfs.TLDfsUploadedProfilePhoto {
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

func (p *kClient) DfsWriteFilePartData(ctx context.Context, req *dfs.TLDfsWriteFilePartData) (r *tg.Bool, err error) {
	// var _args WriteFilePartDataArgs
	// _args.Req = req
	// var _result WriteFilePartDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.writeFilePartData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadPhotoFileV2(ctx context.Context, req *dfs.TLDfsUploadPhotoFileV2) (r *tg.Photo, err error) {
	// var _args UploadPhotoFileV2Args
	// _args.Req = req
	// var _result UploadPhotoFileV2Result

	_result := new(tg.Photo)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadPhotoFileV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadProfilePhotoFileV2(ctx context.Context, req *dfs.TLDfsUploadProfilePhotoFileV2) (r *tg.Photo, err error) {
	// var _args UploadProfilePhotoFileV2Args
	// _args.Req = req
	// var _result UploadProfilePhotoFileV2Result

	_result := new(tg.Photo)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadProfilePhotoFileV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadEncryptedFileV2(ctx context.Context, req *dfs.TLDfsUploadEncryptedFileV2) (r *tg.EncryptedFile, err error) {
	// var _args UploadEncryptedFileV2Args
	// _args.Req = req
	// var _result UploadEncryptedFileV2Result

	_result := new(tg.EncryptedFile)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadEncryptedFileV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsDownloadFile(ctx context.Context, req *dfs.TLDfsDownloadFile) (r *tg.UploadFile, err error) {
	// var _args DownloadFileArgs
	// _args.Req = req
	// var _result DownloadFileResult

	_result := new(tg.UploadFile)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.downloadFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadDocumentFileV2(ctx context.Context, req *dfs.TLDfsUploadDocumentFileV2) (r *tg.Document, err error) {
	// var _args UploadDocumentFileV2Args
	// _args.Req = req
	// var _result UploadDocumentFileV2Result

	_result := new(tg.Document)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadDocumentFileV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadGifDocumentMedia(ctx context.Context, req *dfs.TLDfsUploadGifDocumentMedia) (r *tg.Document, err error) {
	// var _args UploadGifDocumentMediaArgs
	// _args.Req = req
	// var _result UploadGifDocumentMediaResult

	_result := new(tg.Document)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadGifDocumentMedia", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadMp4DocumentMedia(ctx context.Context, req *dfs.TLDfsUploadMp4DocumentMedia) (r *tg.Document, err error) {
	// var _args UploadMp4DocumentMediaArgs
	// _args.Req = req
	// var _result UploadMp4DocumentMediaResult

	_result := new(tg.Document)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadMp4DocumentMedia", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadWallPaperFile(ctx context.Context, req *dfs.TLDfsUploadWallPaperFile) (r *tg.Document, err error) {
	// var _args UploadWallPaperFileArgs
	// _args.Req = req
	// var _result UploadWallPaperFileResult

	_result := new(tg.Document)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadWallPaperFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadThemeFile(ctx context.Context, req *dfs.TLDfsUploadThemeFile) (r *tg.Document, err error) {
	// var _args UploadThemeFileArgs
	// _args.Req = req
	// var _result UploadThemeFileResult

	_result := new(tg.Document)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadThemeFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadRingtoneFile(ctx context.Context, req *dfs.TLDfsUploadRingtoneFile) (r *tg.Document, err error) {
	// var _args UploadRingtoneFileArgs
	// _args.Req = req
	// var _result UploadRingtoneFileResult

	_result := new(tg.Document)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadRingtoneFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) DfsUploadedProfilePhoto(ctx context.Context, req *dfs.TLDfsUploadedProfilePhoto) (r *tg.Photo, err error) {
	// var _args UploadedProfilePhotoArgs
	// _args.Req = req
	// var _result UploadedProfilePhotoResult

	_result := new(tg.Photo)

	if err = p.c.Call(ctx, "/dfs.RPCDfs/dfs.uploadedProfilePhoto", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
