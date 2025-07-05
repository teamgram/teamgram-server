/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package filesservice

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
	"/tg.RPCFiles/messages.getDocumentByHash": kitex.NewMethodInfo(
		messagesGetDocumentByHashHandler,
		newMessagesGetDocumentByHashArgs,
		newMessagesGetDocumentByHashResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/messages.uploadMedia": kitex.NewMethodInfo(
		messagesUploadMediaHandler,
		newMessagesUploadMediaArgs,
		newMessagesUploadMediaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/messages.uploadEncryptedFile": kitex.NewMethodInfo(
		messagesUploadEncryptedFileHandler,
		newMessagesUploadEncryptedFileArgs,
		newMessagesUploadEncryptedFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/upload.saveFilePart": kitex.NewMethodInfo(
		uploadSaveFilePartHandler,
		newUploadSaveFilePartArgs,
		newUploadSaveFilePartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/upload.getFile": kitex.NewMethodInfo(
		uploadGetFileHandler,
		newUploadGetFileArgs,
		newUploadGetFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/upload.saveBigFilePart": kitex.NewMethodInfo(
		uploadSaveBigFilePartHandler,
		newUploadSaveBigFilePartArgs,
		newUploadSaveBigFilePartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/upload.getWebFile": kitex.NewMethodInfo(
		uploadGetWebFileHandler,
		newUploadGetWebFileArgs,
		newUploadGetWebFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/upload.getCdnFile": kitex.NewMethodInfo(
		uploadGetCdnFileHandler,
		newUploadGetCdnFileArgs,
		newUploadGetCdnFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/upload.reuploadCdnFile": kitex.NewMethodInfo(
		uploadReuploadCdnFileHandler,
		newUploadReuploadCdnFileArgs,
		newUploadReuploadCdnFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/upload.getCdnFileHashes": kitex.NewMethodInfo(
		uploadGetCdnFileHashesHandler,
		newUploadGetCdnFileHashesArgs,
		newUploadGetCdnFileHashesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/upload.getFileHashes": kitex.NewMethodInfo(
		uploadGetFileHashesHandler,
		newUploadGetFileHashesArgs,
		newUploadGetFileHashesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/tg.RPCFiles/help.getCdnConfig": kitex.NewMethodInfo(
		helpGetCdnConfigHandler,
		newHelpGetCdnConfigArgs,
		newHelpGetCdnConfigResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	filesServiceServiceInfo                = NewServiceInfo()
	filesServiceServiceInfoForClient       = NewServiceInfoForClient()
	filesServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCFiles", filesServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCFiles", filesServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCFiles", filesServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return filesServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return filesServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return filesServiceServiceInfoForClient
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
	serviceName := "RPCFiles"
	handlerType := (*tg.RPCFiles)(nil)
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
		"PackageName": "files",
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

func messagesGetDocumentByHashHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesGetDocumentByHashArgs)
	realResult := result.(*MessagesGetDocumentByHashResult)
	success, err := handler.(tg.RPCFiles).MessagesGetDocumentByHash(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesGetDocumentByHashArgs() interface{} {
	return &MessagesGetDocumentByHashArgs{}
}

func newMessagesGetDocumentByHashResult() interface{} {
	return &MessagesGetDocumentByHashResult{}
}

type MessagesGetDocumentByHashArgs struct {
	Req *tg.TLMessagesGetDocumentByHash
}

func (p *MessagesGetDocumentByHashArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesGetDocumentByHashArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesGetDocumentByHashArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesGetDocumentByHash)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesGetDocumentByHashArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesGetDocumentByHashArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesGetDocumentByHashArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesGetDocumentByHash)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesGetDocumentByHashArgs_Req_DEFAULT *tg.TLMessagesGetDocumentByHash

func (p *MessagesGetDocumentByHashArgs) GetReq() *tg.TLMessagesGetDocumentByHash {
	if !p.IsSetReq() {
		return MessagesGetDocumentByHashArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesGetDocumentByHashArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesGetDocumentByHashResult struct {
	Success *tg.Document
}

var MessagesGetDocumentByHashResult_Success_DEFAULT *tg.Document

func (p *MessagesGetDocumentByHashResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesGetDocumentByHashResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesGetDocumentByHashResult) Unmarshal(in []byte) error {
	msg := new(tg.Document)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetDocumentByHashResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesGetDocumentByHashResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesGetDocumentByHashResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Document)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesGetDocumentByHashResult) GetSuccess() *tg.Document {
	if !p.IsSetSuccess() {
		return MessagesGetDocumentByHashResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesGetDocumentByHashResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Document)
}

func (p *MessagesGetDocumentByHashResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesGetDocumentByHashResult) GetResult() interface{} {
	return p.Success
}

func messagesUploadMediaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesUploadMediaArgs)
	realResult := result.(*MessagesUploadMediaResult)
	success, err := handler.(tg.RPCFiles).MessagesUploadMedia(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesUploadMediaArgs() interface{} {
	return &MessagesUploadMediaArgs{}
}

func newMessagesUploadMediaResult() interface{} {
	return &MessagesUploadMediaResult{}
}

type MessagesUploadMediaArgs struct {
	Req *tg.TLMessagesUploadMedia
}

func (p *MessagesUploadMediaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesUploadMediaArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesUploadMediaArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesUploadMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesUploadMediaArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesUploadMediaArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesUploadMediaArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesUploadMedia)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesUploadMediaArgs_Req_DEFAULT *tg.TLMessagesUploadMedia

func (p *MessagesUploadMediaArgs) GetReq() *tg.TLMessagesUploadMedia {
	if !p.IsSetReq() {
		return MessagesUploadMediaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesUploadMediaArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesUploadMediaResult struct {
	Success *tg.MessageMedia
}

var MessagesUploadMediaResult_Success_DEFAULT *tg.MessageMedia

func (p *MessagesUploadMediaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesUploadMediaResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesUploadMediaResult) Unmarshal(in []byte) error {
	msg := new(tg.MessageMedia)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesUploadMediaResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesUploadMediaResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesUploadMediaResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MessageMedia)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesUploadMediaResult) GetSuccess() *tg.MessageMedia {
	if !p.IsSetSuccess() {
		return MessagesUploadMediaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesUploadMediaResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MessageMedia)
}

func (p *MessagesUploadMediaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesUploadMediaResult) GetResult() interface{} {
	return p.Success
}

func messagesUploadEncryptedFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*MessagesUploadEncryptedFileArgs)
	realResult := result.(*MessagesUploadEncryptedFileResult)
	success, err := handler.(tg.RPCFiles).MessagesUploadEncryptedFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newMessagesUploadEncryptedFileArgs() interface{} {
	return &MessagesUploadEncryptedFileArgs{}
}

func newMessagesUploadEncryptedFileResult() interface{} {
	return &MessagesUploadEncryptedFileResult{}
}

type MessagesUploadEncryptedFileArgs struct {
	Req *tg.TLMessagesUploadEncryptedFile
}

func (p *MessagesUploadEncryptedFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessagesUploadEncryptedFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *MessagesUploadEncryptedFileArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLMessagesUploadEncryptedFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *MessagesUploadEncryptedFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in MessagesUploadEncryptedFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *MessagesUploadEncryptedFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLMessagesUploadEncryptedFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var MessagesUploadEncryptedFileArgs_Req_DEFAULT *tg.TLMessagesUploadEncryptedFile

func (p *MessagesUploadEncryptedFileArgs) GetReq() *tg.TLMessagesUploadEncryptedFile {
	if !p.IsSetReq() {
		return MessagesUploadEncryptedFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessagesUploadEncryptedFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessagesUploadEncryptedFileResult struct {
	Success *tg.EncryptedFile
}

var MessagesUploadEncryptedFileResult_Success_DEFAULT *tg.EncryptedFile

func (p *MessagesUploadEncryptedFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessagesUploadEncryptedFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *MessagesUploadEncryptedFileResult) Unmarshal(in []byte) error {
	msg := new(tg.EncryptedFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesUploadEncryptedFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in MessagesUploadEncryptedFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *MessagesUploadEncryptedFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.EncryptedFile)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessagesUploadEncryptedFileResult) GetSuccess() *tg.EncryptedFile {
	if !p.IsSetSuccess() {
		return MessagesUploadEncryptedFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessagesUploadEncryptedFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.EncryptedFile)
}

func (p *MessagesUploadEncryptedFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessagesUploadEncryptedFileResult) GetResult() interface{} {
	return p.Success
}

func uploadSaveFilePartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadSaveFilePartArgs)
	realResult := result.(*UploadSaveFilePartResult)
	success, err := handler.(tg.RPCFiles).UploadSaveFilePart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadSaveFilePartArgs() interface{} {
	return &UploadSaveFilePartArgs{}
}

func newUploadSaveFilePartResult() interface{} {
	return &UploadSaveFilePartResult{}
}

type UploadSaveFilePartArgs struct {
	Req *tg.TLUploadSaveFilePart
}

func (p *UploadSaveFilePartArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadSaveFilePartArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadSaveFilePartArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUploadSaveFilePart)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadSaveFilePartArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadSaveFilePartArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadSaveFilePartArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUploadSaveFilePart)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadSaveFilePartArgs_Req_DEFAULT *tg.TLUploadSaveFilePart

func (p *UploadSaveFilePartArgs) GetReq() *tg.TLUploadSaveFilePart {
	if !p.IsSetReq() {
		return UploadSaveFilePartArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadSaveFilePartArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadSaveFilePartResult struct {
	Success *tg.Bool
}

var UploadSaveFilePartResult_Success_DEFAULT *tg.Bool

func (p *UploadSaveFilePartResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadSaveFilePartResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadSaveFilePartResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadSaveFilePartResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadSaveFilePartResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadSaveFilePartResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadSaveFilePartResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UploadSaveFilePartResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadSaveFilePartResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UploadSaveFilePartResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadSaveFilePartResult) GetResult() interface{} {
	return p.Success
}

func uploadGetFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadGetFileArgs)
	realResult := result.(*UploadGetFileResult)
	success, err := handler.(tg.RPCFiles).UploadGetFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadGetFileArgs() interface{} {
	return &UploadGetFileArgs{}
}

func newUploadGetFileResult() interface{} {
	return &UploadGetFileResult{}
}

type UploadGetFileArgs struct {
	Req *tg.TLUploadGetFile
}

func (p *UploadGetFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadGetFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadGetFileArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUploadGetFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadGetFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadGetFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadGetFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUploadGetFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadGetFileArgs_Req_DEFAULT *tg.TLUploadGetFile

func (p *UploadGetFileArgs) GetReq() *tg.TLUploadGetFile {
	if !p.IsSetReq() {
		return UploadGetFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadGetFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadGetFileResult struct {
	Success *tg.UploadFile
}

var UploadGetFileResult_Success_DEFAULT *tg.UploadFile

func (p *UploadGetFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadGetFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadGetFileResult) Unmarshal(in []byte) error {
	msg := new(tg.UploadFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadGetFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadGetFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UploadFile)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetFileResult) GetSuccess() *tg.UploadFile {
	if !p.IsSetSuccess() {
		return UploadGetFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadGetFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UploadFile)
}

func (p *UploadGetFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadGetFileResult) GetResult() interface{} {
	return p.Success
}

func uploadSaveBigFilePartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadSaveBigFilePartArgs)
	realResult := result.(*UploadSaveBigFilePartResult)
	success, err := handler.(tg.RPCFiles).UploadSaveBigFilePart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadSaveBigFilePartArgs() interface{} {
	return &UploadSaveBigFilePartArgs{}
}

func newUploadSaveBigFilePartResult() interface{} {
	return &UploadSaveBigFilePartResult{}
}

type UploadSaveBigFilePartArgs struct {
	Req *tg.TLUploadSaveBigFilePart
}

func (p *UploadSaveBigFilePartArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadSaveBigFilePartArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadSaveBigFilePartArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUploadSaveBigFilePart)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadSaveBigFilePartArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadSaveBigFilePartArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadSaveBigFilePartArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUploadSaveBigFilePart)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadSaveBigFilePartArgs_Req_DEFAULT *tg.TLUploadSaveBigFilePart

func (p *UploadSaveBigFilePartArgs) GetReq() *tg.TLUploadSaveBigFilePart {
	if !p.IsSetReq() {
		return UploadSaveBigFilePartArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadSaveBigFilePartArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadSaveBigFilePartResult struct {
	Success *tg.Bool
}

var UploadSaveBigFilePartResult_Success_DEFAULT *tg.Bool

func (p *UploadSaveBigFilePartResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadSaveBigFilePartResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadSaveBigFilePartResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadSaveBigFilePartResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadSaveBigFilePartResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadSaveBigFilePartResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadSaveBigFilePartResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UploadSaveBigFilePartResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadSaveBigFilePartResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UploadSaveBigFilePartResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadSaveBigFilePartResult) GetResult() interface{} {
	return p.Success
}

func uploadGetWebFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadGetWebFileArgs)
	realResult := result.(*UploadGetWebFileResult)
	success, err := handler.(tg.RPCFiles).UploadGetWebFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadGetWebFileArgs() interface{} {
	return &UploadGetWebFileArgs{}
}

func newUploadGetWebFileResult() interface{} {
	return &UploadGetWebFileResult{}
}

type UploadGetWebFileArgs struct {
	Req *tg.TLUploadGetWebFile
}

func (p *UploadGetWebFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadGetWebFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadGetWebFileArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUploadGetWebFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadGetWebFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadGetWebFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadGetWebFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUploadGetWebFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadGetWebFileArgs_Req_DEFAULT *tg.TLUploadGetWebFile

func (p *UploadGetWebFileArgs) GetReq() *tg.TLUploadGetWebFile {
	if !p.IsSetReq() {
		return UploadGetWebFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadGetWebFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadGetWebFileResult struct {
	Success *tg.UploadWebFile
}

var UploadGetWebFileResult_Success_DEFAULT *tg.UploadWebFile

func (p *UploadGetWebFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadGetWebFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadGetWebFileResult) Unmarshal(in []byte) error {
	msg := new(tg.UploadWebFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetWebFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadGetWebFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadGetWebFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UploadWebFile)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetWebFileResult) GetSuccess() *tg.UploadWebFile {
	if !p.IsSetSuccess() {
		return UploadGetWebFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadGetWebFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UploadWebFile)
}

func (p *UploadGetWebFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadGetWebFileResult) GetResult() interface{} {
	return p.Success
}

func uploadGetCdnFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadGetCdnFileArgs)
	realResult := result.(*UploadGetCdnFileResult)
	success, err := handler.(tg.RPCFiles).UploadGetCdnFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadGetCdnFileArgs() interface{} {
	return &UploadGetCdnFileArgs{}
}

func newUploadGetCdnFileResult() interface{} {
	return &UploadGetCdnFileResult{}
}

type UploadGetCdnFileArgs struct {
	Req *tg.TLUploadGetCdnFile
}

func (p *UploadGetCdnFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadGetCdnFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadGetCdnFileArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUploadGetCdnFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadGetCdnFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadGetCdnFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadGetCdnFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUploadGetCdnFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadGetCdnFileArgs_Req_DEFAULT *tg.TLUploadGetCdnFile

func (p *UploadGetCdnFileArgs) GetReq() *tg.TLUploadGetCdnFile {
	if !p.IsSetReq() {
		return UploadGetCdnFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadGetCdnFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadGetCdnFileResult struct {
	Success *tg.UploadCdnFile
}

var UploadGetCdnFileResult_Success_DEFAULT *tg.UploadCdnFile

func (p *UploadGetCdnFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadGetCdnFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadGetCdnFileResult) Unmarshal(in []byte) error {
	msg := new(tg.UploadCdnFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetCdnFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadGetCdnFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadGetCdnFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UploadCdnFile)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetCdnFileResult) GetSuccess() *tg.UploadCdnFile {
	if !p.IsSetSuccess() {
		return UploadGetCdnFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadGetCdnFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UploadCdnFile)
}

func (p *UploadGetCdnFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadGetCdnFileResult) GetResult() interface{} {
	return p.Success
}

func uploadReuploadCdnFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadReuploadCdnFileArgs)
	realResult := result.(*UploadReuploadCdnFileResult)
	success, err := handler.(tg.RPCFiles).UploadReuploadCdnFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadReuploadCdnFileArgs() interface{} {
	return &UploadReuploadCdnFileArgs{}
}

func newUploadReuploadCdnFileResult() interface{} {
	return &UploadReuploadCdnFileResult{}
}

type UploadReuploadCdnFileArgs struct {
	Req *tg.TLUploadReuploadCdnFile
}

func (p *UploadReuploadCdnFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadReuploadCdnFileArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadReuploadCdnFileArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUploadReuploadCdnFile)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadReuploadCdnFileArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadReuploadCdnFileArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadReuploadCdnFileArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUploadReuploadCdnFile)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadReuploadCdnFileArgs_Req_DEFAULT *tg.TLUploadReuploadCdnFile

func (p *UploadReuploadCdnFileArgs) GetReq() *tg.TLUploadReuploadCdnFile {
	if !p.IsSetReq() {
		return UploadReuploadCdnFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadReuploadCdnFileArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadReuploadCdnFileResult struct {
	Success *tg.VectorFileHash
}

var UploadReuploadCdnFileResult_Success_DEFAULT *tg.VectorFileHash

func (p *UploadReuploadCdnFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadReuploadCdnFileResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadReuploadCdnFileResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorFileHash)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadReuploadCdnFileResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadReuploadCdnFileResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadReuploadCdnFileResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorFileHash)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadReuploadCdnFileResult) GetSuccess() *tg.VectorFileHash {
	if !p.IsSetSuccess() {
		return UploadReuploadCdnFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadReuploadCdnFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorFileHash)
}

func (p *UploadReuploadCdnFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadReuploadCdnFileResult) GetResult() interface{} {
	return p.Success
}

func uploadGetCdnFileHashesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadGetCdnFileHashesArgs)
	realResult := result.(*UploadGetCdnFileHashesResult)
	success, err := handler.(tg.RPCFiles).UploadGetCdnFileHashes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadGetCdnFileHashesArgs() interface{} {
	return &UploadGetCdnFileHashesArgs{}
}

func newUploadGetCdnFileHashesResult() interface{} {
	return &UploadGetCdnFileHashesResult{}
}

type UploadGetCdnFileHashesArgs struct {
	Req *tg.TLUploadGetCdnFileHashes
}

func (p *UploadGetCdnFileHashesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadGetCdnFileHashesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadGetCdnFileHashesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUploadGetCdnFileHashes)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadGetCdnFileHashesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadGetCdnFileHashesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadGetCdnFileHashesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUploadGetCdnFileHashes)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadGetCdnFileHashesArgs_Req_DEFAULT *tg.TLUploadGetCdnFileHashes

func (p *UploadGetCdnFileHashesArgs) GetReq() *tg.TLUploadGetCdnFileHashes {
	if !p.IsSetReq() {
		return UploadGetCdnFileHashesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadGetCdnFileHashesArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadGetCdnFileHashesResult struct {
	Success *tg.VectorFileHash
}

var UploadGetCdnFileHashesResult_Success_DEFAULT *tg.VectorFileHash

func (p *UploadGetCdnFileHashesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadGetCdnFileHashesResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadGetCdnFileHashesResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorFileHash)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetCdnFileHashesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadGetCdnFileHashesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadGetCdnFileHashesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorFileHash)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetCdnFileHashesResult) GetSuccess() *tg.VectorFileHash {
	if !p.IsSetSuccess() {
		return UploadGetCdnFileHashesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadGetCdnFileHashesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorFileHash)
}

func (p *UploadGetCdnFileHashesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadGetCdnFileHashesResult) GetResult() interface{} {
	return p.Success
}

func uploadGetFileHashesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UploadGetFileHashesArgs)
	realResult := result.(*UploadGetFileHashesResult)
	success, err := handler.(tg.RPCFiles).UploadGetFileHashes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUploadGetFileHashesArgs() interface{} {
	return &UploadGetFileHashesArgs{}
}

func newUploadGetFileHashesResult() interface{} {
	return &UploadGetFileHashesResult{}
}

type UploadGetFileHashesArgs struct {
	Req *tg.TLUploadGetFileHashes
}

func (p *UploadGetFileHashesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadGetFileHashesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UploadGetFileHashesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLUploadGetFileHashes)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UploadGetFileHashesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UploadGetFileHashesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UploadGetFileHashesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLUploadGetFileHashes)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UploadGetFileHashesArgs_Req_DEFAULT *tg.TLUploadGetFileHashes

func (p *UploadGetFileHashesArgs) GetReq() *tg.TLUploadGetFileHashes {
	if !p.IsSetReq() {
		return UploadGetFileHashesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadGetFileHashesArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadGetFileHashesResult struct {
	Success *tg.VectorFileHash
}

var UploadGetFileHashesResult_Success_DEFAULT *tg.VectorFileHash

func (p *UploadGetFileHashesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadGetFileHashesResult")
	}
	return json.Marshal(p.Success)
}

func (p *UploadGetFileHashesResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorFileHash)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetFileHashesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UploadGetFileHashesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UploadGetFileHashesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorFileHash)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadGetFileHashesResult) GetSuccess() *tg.VectorFileHash {
	if !p.IsSetSuccess() {
		return UploadGetFileHashesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadGetFileHashesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorFileHash)
}

func (p *UploadGetFileHashesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadGetFileHashesResult) GetResult() interface{} {
	return p.Success
}

func helpGetCdnConfigHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetCdnConfigArgs)
	realResult := result.(*HelpGetCdnConfigResult)
	success, err := handler.(tg.RPCFiles).HelpGetCdnConfig(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetCdnConfigArgs() interface{} {
	return &HelpGetCdnConfigArgs{}
}

func newHelpGetCdnConfigResult() interface{} {
	return &HelpGetCdnConfigResult{}
}

type HelpGetCdnConfigArgs struct {
	Req *tg.TLHelpGetCdnConfig
}

func (p *HelpGetCdnConfigArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetCdnConfigArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetCdnConfigArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetCdnConfig)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetCdnConfigArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetCdnConfigArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetCdnConfigArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetCdnConfig)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetCdnConfigArgs_Req_DEFAULT *tg.TLHelpGetCdnConfig

func (p *HelpGetCdnConfigArgs) GetReq() *tg.TLHelpGetCdnConfig {
	if !p.IsSetReq() {
		return HelpGetCdnConfigArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetCdnConfigArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetCdnConfigResult struct {
	Success *tg.CdnConfig
}

var HelpGetCdnConfigResult_Success_DEFAULT *tg.CdnConfig

func (p *HelpGetCdnConfigResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetCdnConfigResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetCdnConfigResult) Unmarshal(in []byte) error {
	msg := new(tg.CdnConfig)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetCdnConfigResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetCdnConfigResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetCdnConfigResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.CdnConfig)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetCdnConfigResult) GetSuccess() *tg.CdnConfig {
	if !p.IsSetSuccess() {
		return HelpGetCdnConfigResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetCdnConfigResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.CdnConfig)
}

func (p *HelpGetCdnConfigResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetCdnConfigResult) GetResult() interface{} {
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

func (p *kClient) MessagesGetDocumentByHash(ctx context.Context, req *tg.TLMessagesGetDocumentByHash) (r *tg.Document, err error) {
	// var _args MessagesGetDocumentByHashArgs
	// _args.Req = req
	// var _result MessagesGetDocumentByHashResult

	_result := new(tg.Document)
	if err = p.c.Call(ctx, "/tg.RPCFiles/messages.getDocumentByHash", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesUploadMedia(ctx context.Context, req *tg.TLMessagesUploadMedia) (r *tg.MessageMedia, err error) {
	// var _args MessagesUploadMediaArgs
	// _args.Req = req
	// var _result MessagesUploadMediaResult

	_result := new(tg.MessageMedia)
	if err = p.c.Call(ctx, "/tg.RPCFiles/messages.uploadMedia", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) MessagesUploadEncryptedFile(ctx context.Context, req *tg.TLMessagesUploadEncryptedFile) (r *tg.EncryptedFile, err error) {
	// var _args MessagesUploadEncryptedFileArgs
	// _args.Req = req
	// var _result MessagesUploadEncryptedFileResult

	_result := new(tg.EncryptedFile)
	if err = p.c.Call(ctx, "/tg.RPCFiles/messages.uploadEncryptedFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UploadSaveFilePart(ctx context.Context, req *tg.TLUploadSaveFilePart) (r *tg.Bool, err error) {
	// var _args UploadSaveFilePartArgs
	// _args.Req = req
	// var _result UploadSaveFilePartResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCFiles/upload.saveFilePart", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UploadGetFile(ctx context.Context, req *tg.TLUploadGetFile) (r *tg.UploadFile, err error) {
	// var _args UploadGetFileArgs
	// _args.Req = req
	// var _result UploadGetFileResult

	_result := new(tg.UploadFile)
	if err = p.c.Call(ctx, "/tg.RPCFiles/upload.getFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UploadSaveBigFilePart(ctx context.Context, req *tg.TLUploadSaveBigFilePart) (r *tg.Bool, err error) {
	// var _args UploadSaveBigFilePartArgs
	// _args.Req = req
	// var _result UploadSaveBigFilePartResult

	_result := new(tg.Bool)
	if err = p.c.Call(ctx, "/tg.RPCFiles/upload.saveBigFilePart", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UploadGetWebFile(ctx context.Context, req *tg.TLUploadGetWebFile) (r *tg.UploadWebFile, err error) {
	// var _args UploadGetWebFileArgs
	// _args.Req = req
	// var _result UploadGetWebFileResult

	_result := new(tg.UploadWebFile)
	if err = p.c.Call(ctx, "/tg.RPCFiles/upload.getWebFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UploadGetCdnFile(ctx context.Context, req *tg.TLUploadGetCdnFile) (r *tg.UploadCdnFile, err error) {
	// var _args UploadGetCdnFileArgs
	// _args.Req = req
	// var _result UploadGetCdnFileResult

	_result := new(tg.UploadCdnFile)
	if err = p.c.Call(ctx, "/tg.RPCFiles/upload.getCdnFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UploadReuploadCdnFile(ctx context.Context, req *tg.TLUploadReuploadCdnFile) (r *tg.VectorFileHash, err error) {
	// var _args UploadReuploadCdnFileArgs
	// _args.Req = req
	// var _result UploadReuploadCdnFileResult

	_result := new(tg.VectorFileHash)
	if err = p.c.Call(ctx, "/tg.RPCFiles/upload.reuploadCdnFile", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UploadGetCdnFileHashes(ctx context.Context, req *tg.TLUploadGetCdnFileHashes) (r *tg.VectorFileHash, err error) {
	// var _args UploadGetCdnFileHashesArgs
	// _args.Req = req
	// var _result UploadGetCdnFileHashesResult

	_result := new(tg.VectorFileHash)
	if err = p.c.Call(ctx, "/tg.RPCFiles/upload.getCdnFileHashes", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UploadGetFileHashes(ctx context.Context, req *tg.TLUploadGetFileHashes) (r *tg.VectorFileHash, err error) {
	// var _args UploadGetFileHashesArgs
	// _args.Req = req
	// var _result UploadGetFileHashesResult

	_result := new(tg.VectorFileHash)
	if err = p.c.Call(ctx, "/tg.RPCFiles/upload.getFileHashes", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) HelpGetCdnConfig(ctx context.Context, req *tg.TLHelpGetCdnConfig) (r *tg.CdnConfig, err error) {
	// var _args HelpGetCdnConfigArgs
	// _args.Req = req
	// var _result HelpGetCdnConfigResult

	_result := new(tg.CdnConfig)
	if err = p.c.Call(ctx, "/tg.RPCFiles/help.getCdnConfig", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
