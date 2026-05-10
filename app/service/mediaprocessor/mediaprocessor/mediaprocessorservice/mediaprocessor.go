/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mediaprocessorservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

func decodeConstructorIfPresent(d *bin.Decoder, msg interface{}) error {
	v := reflect.ValueOf(msg)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}

	f := v.FieldByName("ClazzID")
	if !f.IsValid() || !f.CanSet() || f.Kind() != reflect.Uint32 {
		return nil
	}

	clazzID, err := d.ClazzID()
	if err != nil {
		return err
	}
	f.SetUint(uint64(clazzID))
	return nil
}

var serviceMethods = map[string]kitex.MethodInfo{
	"/mediaprocessor.RPCMediaProcessor/mediaProcessor.processPhoto": kitex.NewMethodInfo(
		processPhotoHandler,
		newProcessPhotoArgs,
		newProcessPhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/mediaprocessor.RPCMediaProcessor/mediaProcessor.processGif": kitex.NewMethodInfo(
		processGifHandler,
		newProcessGifArgs,
		newProcessGifResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/mediaprocessor.RPCMediaProcessor/mediaProcessor.processMp4": kitex.NewMethodInfo(
		processMp4Handler,
		newProcessMp4Args,
		newProcessMp4Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	mediaprocessorServiceServiceInfo                = NewServiceInfo()
	mediaprocessorServiceServiceInfoForClient       = NewServiceInfoForClient()
	mediaprocessorServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCMediaProcessor", mediaprocessorServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCMediaProcessor", mediaprocessorServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCMediaProcessor", mediaprocessorServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return mediaprocessorServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return mediaprocessorServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return mediaprocessorServiceServiceInfoForClient
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
	serviceName := "RPCMediaProcessor"
	handlerType := (*mediaprocessor.RPCMediaProcessor)(nil)
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
		"PackageName": "mediaprocessor",
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

func processPhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ProcessPhotoArgs)
	realResult := result.(*ProcessPhotoResult)
	success, err := handler.(mediaprocessor.RPCMediaProcessor).MediaProcessorProcessPhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newProcessPhotoArgs() interface{} {
	return &ProcessPhotoArgs{}
}

func newProcessPhotoResult() interface{} {
	return &ProcessPhotoResult{}
}

type ProcessPhotoArgs struct {
	Req *mediaprocessor.TLMediaProcessorProcessPhoto
}

func (p *ProcessPhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in ProcessPhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ProcessPhotoArgs) Unmarshal(in []byte) error {
	msg := new(mediaprocessor.TLMediaProcessorProcessPhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ProcessPhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in ProcessPhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ProcessPhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(mediaprocessor.TLMediaProcessorProcessPhoto)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProcessPhotoArgs_Req_DEFAULT *mediaprocessor.TLMediaProcessorProcessPhoto

func (p *ProcessPhotoArgs) GetReq() *mediaprocessor.TLMediaProcessorProcessPhoto {
	if !p.IsSetReq() {
		return ProcessPhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProcessPhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type ProcessPhotoResult struct {
	Success *mediaprocessor.ProcessedPhoto
}

var ProcessPhotoResult_Success_DEFAULT *mediaprocessor.ProcessedPhoto

func (p *ProcessPhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in ProcessPhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *ProcessPhotoResult) Unmarshal(in []byte) error {
	msg := new(mediaprocessor.ProcessedPhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessPhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in ProcessPhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ProcessPhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(mediaprocessor.ProcessedPhoto)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessPhotoResult) GetSuccess() *mediaprocessor.ProcessedPhoto {
	if !p.IsSetSuccess() {
		return ProcessPhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProcessPhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*mediaprocessor.ProcessedPhoto)
}

func (p *ProcessPhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProcessPhotoResult) GetResult() interface{} {
	return p.Success
}

func processGifHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ProcessGifArgs)
	realResult := result.(*ProcessGifResult)
	success, err := handler.(mediaprocessor.RPCMediaProcessor).MediaProcessorProcessGif(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newProcessGifArgs() interface{} {
	return &ProcessGifArgs{}
}

func newProcessGifResult() interface{} {
	return &ProcessGifResult{}
}

type ProcessGifArgs struct {
	Req *mediaprocessor.TLMediaProcessorProcessGif
}

func (p *ProcessGifArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in ProcessGifArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ProcessGifArgs) Unmarshal(in []byte) error {
	msg := new(mediaprocessor.TLMediaProcessorProcessGif)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ProcessGifArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in ProcessGifArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ProcessGifArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(mediaprocessor.TLMediaProcessorProcessGif)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProcessGifArgs_Req_DEFAULT *mediaprocessor.TLMediaProcessorProcessGif

func (p *ProcessGifArgs) GetReq() *mediaprocessor.TLMediaProcessorProcessGif {
	if !p.IsSetReq() {
		return ProcessGifArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProcessGifArgs) IsSetReq() bool {
	return p.Req != nil
}

type ProcessGifResult struct {
	Success *mediaprocessor.ProcessedDocument
}

var ProcessGifResult_Success_DEFAULT *mediaprocessor.ProcessedDocument

func (p *ProcessGifResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in ProcessGifResult")
	}
	return json.Marshal(p.Success)
}

func (p *ProcessGifResult) Unmarshal(in []byte) error {
	msg := new(mediaprocessor.ProcessedDocument)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessGifResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in ProcessGifResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ProcessGifResult) Decode(d *bin.Decoder) (err error) {
	msg := new(mediaprocessor.ProcessedDocument)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessGifResult) GetSuccess() *mediaprocessor.ProcessedDocument {
	if !p.IsSetSuccess() {
		return ProcessGifResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProcessGifResult) SetSuccess(x interface{}) {
	p.Success = x.(*mediaprocessor.ProcessedDocument)
}

func (p *ProcessGifResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProcessGifResult) GetResult() interface{} {
	return p.Success
}

func processMp4Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ProcessMp4Args)
	realResult := result.(*ProcessMp4Result)
	success, err := handler.(mediaprocessor.RPCMediaProcessor).MediaProcessorProcessMp4(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newProcessMp4Args() interface{} {
	return &ProcessMp4Args{}
}

func newProcessMp4Result() interface{} {
	return &ProcessMp4Result{}
}

type ProcessMp4Args struct {
	Req *mediaprocessor.TLMediaProcessorProcessMp4
}

func (p *ProcessMp4Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("no req in ProcessMp4Args")
	}
	return json.Marshal(p.Req)
}

func (p *ProcessMp4Args) Unmarshal(in []byte) error {
	msg := new(mediaprocessor.TLMediaProcessorProcessMp4)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ProcessMp4Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("no req in ProcessMp4Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *ProcessMp4Args) Decode(d *bin.Decoder) (err error) {
	msg := new(mediaprocessor.TLMediaProcessorProcessMp4)
	msg.ClazzID, err = d.ClazzID()
	if err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProcessMp4Args_Req_DEFAULT *mediaprocessor.TLMediaProcessorProcessMp4

func (p *ProcessMp4Args) GetReq() *mediaprocessor.TLMediaProcessorProcessMp4 {
	if !p.IsSetReq() {
		return ProcessMp4Args_Req_DEFAULT
	}
	return p.Req
}

func (p *ProcessMp4Args) IsSetReq() bool {
	return p.Req != nil
}

type ProcessMp4Result struct {
	Success *mediaprocessor.ProcessedDocument
}

var ProcessMp4Result_Success_DEFAULT *mediaprocessor.ProcessedDocument

func (p *ProcessMp4Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("no req in ProcessMp4Result")
	}
	return json.Marshal(p.Success)
}

func (p *ProcessMp4Result) Unmarshal(in []byte) error {
	msg := new(mediaprocessor.ProcessedDocument)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessMp4Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("no req in ProcessMp4Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *ProcessMp4Result) Decode(d *bin.Decoder) (err error) {
	msg := new(mediaprocessor.ProcessedDocument)
	if err = decodeConstructorIfPresent(d, msg); err != nil {
		return err
	}
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessMp4Result) GetSuccess() *mediaprocessor.ProcessedDocument {
	if !p.IsSetSuccess() {
		return ProcessMp4Result_Success_DEFAULT
	}
	return p.Success
}

func (p *ProcessMp4Result) SetSuccess(x interface{}) {
	p.Success = x.(*mediaprocessor.ProcessedDocument)
}

func (p *ProcessMp4Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProcessMp4Result) GetResult() interface{} {
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

func (p *kClient) MediaProcessorProcessPhoto(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessPhoto) (r *mediaprocessor.ProcessedPhoto, err error) {
	// var _args ProcessPhotoArgs
	// _args.Req = req
	var _result ProcessPhotoResult

	if err = p.c.Call(ctx, "/mediaprocessor.RPCMediaProcessor/mediaProcessor.processPhoto", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) MediaProcessorProcessGif(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessGif) (r *mediaprocessor.ProcessedDocument, err error) {
	// var _args ProcessGifArgs
	// _args.Req = req
	var _result ProcessGifResult

	if err = p.c.Call(ctx, "/mediaprocessor.RPCMediaProcessor/mediaProcessor.processGif", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}

func (p *kClient) MediaProcessorProcessMp4(ctx context.Context, req *mediaprocessor.TLMediaProcessorProcessMp4) (r *mediaprocessor.ProcessedDocument, err error) {
	// var _args ProcessMp4Args
	// _args.Req = req
	var _result ProcessMp4Result

	if err = p.c.Call(ctx, "/mediaprocessor.RPCMediaProcessor/mediaProcessor.processMp4", req, &_result); err != nil {
		return
	}

	return _result.GetSuccess(), nil
}
