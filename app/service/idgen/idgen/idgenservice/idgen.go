/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package idgenservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/idgen.RPCIdgen/idgen.nextId": kitex.NewMethodInfo(
		nextIdHandler,
		newNextIdArgs,
		newNextIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/idgen.RPCIdgen/idgen.nextIds": kitex.NewMethodInfo(
		nextIdsHandler,
		newNextIdsArgs,
		newNextIdsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/idgen.RPCIdgen/idgen.getCurrentSeqId": kitex.NewMethodInfo(
		getCurrentSeqIdHandler,
		newGetCurrentSeqIdArgs,
		newGetCurrentSeqIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/idgen.RPCIdgen/idgen.setCurrentSeqId": kitex.NewMethodInfo(
		setCurrentSeqIdHandler,
		newSetCurrentSeqIdArgs,
		newSetCurrentSeqIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/idgen.RPCIdgen/idgen.getNextSeqId": kitex.NewMethodInfo(
		getNextSeqIdHandler,
		newGetNextSeqIdArgs,
		newGetNextSeqIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/idgen.RPCIdgen/idgen.getNextNSeqId": kitex.NewMethodInfo(
		getNextNSeqIdHandler,
		newGetNextNSeqIdArgs,
		newGetNextNSeqIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/idgen.RPCIdgen/idgen.getNextIdValList": kitex.NewMethodInfo(
		getNextIdValListHandler,
		newGetNextIdValListArgs,
		newGetNextIdValListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/idgen.RPCIdgen/idgen.getCurrentSeqIdList": kitex.NewMethodInfo(
		getCurrentSeqIdListHandler,
		newGetCurrentSeqIdListArgs,
		newGetCurrentSeqIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	idgenServiceServiceInfo                = NewServiceInfo()
	idgenServiceServiceInfoForClient       = NewServiceInfoForClient()
	idgenServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCIdgen", idgenServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCIdgen", idgenServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCIdgen", idgenServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return idgenServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return idgenServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return idgenServiceServiceInfoForClient
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
	serviceName := "RPCIdgen"
	handlerType := (*idgen.RPCIdgen)(nil)
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
		"PackageName": "idgen",
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

func nextIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*NextIdArgs)
	realResult := result.(*NextIdResult)
	success, err := handler.(idgen.RPCIdgen).IdgenNextId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newNextIdArgs() interface{} {
	return &NextIdArgs{}
}

func newNextIdResult() interface{} {
	return &NextIdResult{}
}

type NextIdArgs struct {
	Req *idgen.TLIdgenNextId
}

func (p *NextIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in NextIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *NextIdArgs) Unmarshal(in []byte) error {
	msg := new(idgen.TLIdgenNextId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *NextIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in NextIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *NextIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.TLIdgenNextId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var NextIdArgs_Req_DEFAULT *idgen.TLIdgenNextId

func (p *NextIdArgs) GetReq() *idgen.TLIdgenNextId {
	if !p.IsSetReq() {
		return NextIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NextIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type NextIdResult struct {
	Success *tg.Int64
}

var NextIdResult_Success_DEFAULT *tg.Int64

func (p *NextIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in NextIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *NextIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NextIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in NextIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *NextIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NextIdResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return NextIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NextIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *NextIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NextIdResult) GetResult() interface{} {
	return p.Success
}

func nextIdsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*NextIdsArgs)
	realResult := result.(*NextIdsResult)
	success, err := handler.(idgen.RPCIdgen).IdgenNextIds(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newNextIdsArgs() interface{} {
	return &NextIdsArgs{}
}

func newNextIdsResult() interface{} {
	return &NextIdsResult{}
}

type NextIdsArgs struct {
	Req *idgen.TLIdgenNextIds
}

func (p *NextIdsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in NextIdsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *NextIdsArgs) Unmarshal(in []byte) error {
	msg := new(idgen.TLIdgenNextIds)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *NextIdsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in NextIdsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *NextIdsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.TLIdgenNextIds)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var NextIdsArgs_Req_DEFAULT *idgen.TLIdgenNextIds

func (p *NextIdsArgs) GetReq() *idgen.TLIdgenNextIds {
	if !p.IsSetReq() {
		return NextIdsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NextIdsArgs) IsSetReq() bool {
	return p.Req != nil
}

type NextIdsResult struct {
	Success *idgen.VectorLong
}

var NextIdsResult_Success_DEFAULT *idgen.VectorLong

func (p *NextIdsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in NextIdsResult")
	}
	return json.Marshal(p.Success)
}

func (p *NextIdsResult) Unmarshal(in []byte) error {
	msg := new(idgen.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NextIdsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in NextIdsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *NextIdsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NextIdsResult) GetSuccess() *idgen.VectorLong {
	if !p.IsSetSuccess() {
		return NextIdsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NextIdsResult) SetSuccess(x interface{}) {
	p.Success = x.(*idgen.VectorLong)
}

func (p *NextIdsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NextIdsResult) GetResult() interface{} {
	return p.Success
}

func getCurrentSeqIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetCurrentSeqIdArgs)
	realResult := result.(*GetCurrentSeqIdResult)
	success, err := handler.(idgen.RPCIdgen).IdgenGetCurrentSeqId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetCurrentSeqIdArgs() interface{} {
	return &GetCurrentSeqIdArgs{}
}

func newGetCurrentSeqIdResult() interface{} {
	return &GetCurrentSeqIdResult{}
}

type GetCurrentSeqIdArgs struct {
	Req *idgen.TLIdgenGetCurrentSeqId
}

func (p *GetCurrentSeqIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCurrentSeqIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetCurrentSeqIdArgs) Unmarshal(in []byte) error {
	msg := new(idgen.TLIdgenGetCurrentSeqId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetCurrentSeqIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetCurrentSeqIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetCurrentSeqIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.TLIdgenGetCurrentSeqId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetCurrentSeqIdArgs_Req_DEFAULT *idgen.TLIdgenGetCurrentSeqId

func (p *GetCurrentSeqIdArgs) GetReq() *idgen.TLIdgenGetCurrentSeqId {
	if !p.IsSetReq() {
		return GetCurrentSeqIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCurrentSeqIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCurrentSeqIdResult struct {
	Success *tg.Int64
}

var GetCurrentSeqIdResult_Success_DEFAULT *tg.Int64

func (p *GetCurrentSeqIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCurrentSeqIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetCurrentSeqIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCurrentSeqIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetCurrentSeqIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetCurrentSeqIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCurrentSeqIdResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return GetCurrentSeqIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCurrentSeqIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *GetCurrentSeqIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCurrentSeqIdResult) GetResult() interface{} {
	return p.Success
}

func setCurrentSeqIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetCurrentSeqIdArgs)
	realResult := result.(*SetCurrentSeqIdResult)
	success, err := handler.(idgen.RPCIdgen).IdgenSetCurrentSeqId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetCurrentSeqIdArgs() interface{} {
	return &SetCurrentSeqIdArgs{}
}

func newSetCurrentSeqIdResult() interface{} {
	return &SetCurrentSeqIdResult{}
}

type SetCurrentSeqIdArgs struct {
	Req *idgen.TLIdgenSetCurrentSeqId
}

func (p *SetCurrentSeqIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetCurrentSeqIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetCurrentSeqIdArgs) Unmarshal(in []byte) error {
	msg := new(idgen.TLIdgenSetCurrentSeqId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetCurrentSeqIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetCurrentSeqIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetCurrentSeqIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.TLIdgenSetCurrentSeqId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetCurrentSeqIdArgs_Req_DEFAULT *idgen.TLIdgenSetCurrentSeqId

func (p *SetCurrentSeqIdArgs) GetReq() *idgen.TLIdgenSetCurrentSeqId {
	if !p.IsSetReq() {
		return SetCurrentSeqIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetCurrentSeqIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetCurrentSeqIdResult struct {
	Success *tg.Bool
}

var SetCurrentSeqIdResult_Success_DEFAULT *tg.Bool

func (p *SetCurrentSeqIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetCurrentSeqIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetCurrentSeqIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetCurrentSeqIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetCurrentSeqIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetCurrentSeqIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetCurrentSeqIdResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetCurrentSeqIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetCurrentSeqIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetCurrentSeqIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetCurrentSeqIdResult) GetResult() interface{} {
	return p.Success
}

func getNextSeqIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetNextSeqIdArgs)
	realResult := result.(*GetNextSeqIdResult)
	success, err := handler.(idgen.RPCIdgen).IdgenGetNextSeqId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetNextSeqIdArgs() interface{} {
	return &GetNextSeqIdArgs{}
}

func newGetNextSeqIdResult() interface{} {
	return &GetNextSeqIdResult{}
}

type GetNextSeqIdArgs struct {
	Req *idgen.TLIdgenGetNextSeqId
}

func (p *GetNextSeqIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetNextSeqIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetNextSeqIdArgs) Unmarshal(in []byte) error {
	msg := new(idgen.TLIdgenGetNextSeqId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetNextSeqIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetNextSeqIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetNextSeqIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.TLIdgenGetNextSeqId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetNextSeqIdArgs_Req_DEFAULT *idgen.TLIdgenGetNextSeqId

func (p *GetNextSeqIdArgs) GetReq() *idgen.TLIdgenGetNextSeqId {
	if !p.IsSetReq() {
		return GetNextSeqIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetNextSeqIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetNextSeqIdResult struct {
	Success *tg.Int64
}

var GetNextSeqIdResult_Success_DEFAULT *tg.Int64

func (p *GetNextSeqIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetNextSeqIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetNextSeqIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNextSeqIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetNextSeqIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetNextSeqIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNextSeqIdResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return GetNextSeqIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetNextSeqIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *GetNextSeqIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetNextSeqIdResult) GetResult() interface{} {
	return p.Success
}

func getNextNSeqIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetNextNSeqIdArgs)
	realResult := result.(*GetNextNSeqIdResult)
	success, err := handler.(idgen.RPCIdgen).IdgenGetNextNSeqId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetNextNSeqIdArgs() interface{} {
	return &GetNextNSeqIdArgs{}
}

func newGetNextNSeqIdResult() interface{} {
	return &GetNextNSeqIdResult{}
}

type GetNextNSeqIdArgs struct {
	Req *idgen.TLIdgenGetNextNSeqId
}

func (p *GetNextNSeqIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetNextNSeqIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetNextNSeqIdArgs) Unmarshal(in []byte) error {
	msg := new(idgen.TLIdgenGetNextNSeqId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetNextNSeqIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetNextNSeqIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetNextNSeqIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.TLIdgenGetNextNSeqId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetNextNSeqIdArgs_Req_DEFAULT *idgen.TLIdgenGetNextNSeqId

func (p *GetNextNSeqIdArgs) GetReq() *idgen.TLIdgenGetNextNSeqId {
	if !p.IsSetReq() {
		return GetNextNSeqIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetNextNSeqIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetNextNSeqIdResult struct {
	Success *tg.Int64
}

var GetNextNSeqIdResult_Success_DEFAULT *tg.Int64

func (p *GetNextNSeqIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetNextNSeqIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetNextNSeqIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNextNSeqIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetNextNSeqIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetNextNSeqIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNextNSeqIdResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return GetNextNSeqIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetNextNSeqIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *GetNextNSeqIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetNextNSeqIdResult) GetResult() interface{} {
	return p.Success
}

func getNextIdValListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetNextIdValListArgs)
	realResult := result.(*GetNextIdValListResult)
	success, err := handler.(idgen.RPCIdgen).IdgenGetNextIdValList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetNextIdValListArgs() interface{} {
	return &GetNextIdValListArgs{}
}

func newGetNextIdValListResult() interface{} {
	return &GetNextIdValListResult{}
}

type GetNextIdValListArgs struct {
	Req *idgen.TLIdgenGetNextIdValList
}

func (p *GetNextIdValListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetNextIdValListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetNextIdValListArgs) Unmarshal(in []byte) error {
	msg := new(idgen.TLIdgenGetNextIdValList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetNextIdValListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetNextIdValListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetNextIdValListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.TLIdgenGetNextIdValList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetNextIdValListArgs_Req_DEFAULT *idgen.TLIdgenGetNextIdValList

func (p *GetNextIdValListArgs) GetReq() *idgen.TLIdgenGetNextIdValList {
	if !p.IsSetReq() {
		return GetNextIdValListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetNextIdValListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetNextIdValListResult struct {
	Success *idgen.VectorIdVal
}

var GetNextIdValListResult_Success_DEFAULT *idgen.VectorIdVal

func (p *GetNextIdValListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetNextIdValListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetNextIdValListResult) Unmarshal(in []byte) error {
	msg := new(idgen.VectorIdVal)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNextIdValListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetNextIdValListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetNextIdValListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.VectorIdVal)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNextIdValListResult) GetSuccess() *idgen.VectorIdVal {
	if !p.IsSetSuccess() {
		return GetNextIdValListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetNextIdValListResult) SetSuccess(x interface{}) {
	p.Success = x.(*idgen.VectorIdVal)
}

func (p *GetNextIdValListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetNextIdValListResult) GetResult() interface{} {
	return p.Success
}

func getCurrentSeqIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetCurrentSeqIdListArgs)
	realResult := result.(*GetCurrentSeqIdListResult)
	success, err := handler.(idgen.RPCIdgen).IdgenGetCurrentSeqIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetCurrentSeqIdListArgs() interface{} {
	return &GetCurrentSeqIdListArgs{}
}

func newGetCurrentSeqIdListResult() interface{} {
	return &GetCurrentSeqIdListResult{}
}

type GetCurrentSeqIdListArgs struct {
	Req *idgen.TLIdgenGetCurrentSeqIdList
}

func (p *GetCurrentSeqIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCurrentSeqIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetCurrentSeqIdListArgs) Unmarshal(in []byte) error {
	msg := new(idgen.TLIdgenGetCurrentSeqIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetCurrentSeqIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetCurrentSeqIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetCurrentSeqIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.TLIdgenGetCurrentSeqIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetCurrentSeqIdListArgs_Req_DEFAULT *idgen.TLIdgenGetCurrentSeqIdList

func (p *GetCurrentSeqIdListArgs) GetReq() *idgen.TLIdgenGetCurrentSeqIdList {
	if !p.IsSetReq() {
		return GetCurrentSeqIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCurrentSeqIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCurrentSeqIdListResult struct {
	Success *idgen.VectorIdVal
}

var GetCurrentSeqIdListResult_Success_DEFAULT *idgen.VectorIdVal

func (p *GetCurrentSeqIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCurrentSeqIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetCurrentSeqIdListResult) Unmarshal(in []byte) error {
	msg := new(idgen.VectorIdVal)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCurrentSeqIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetCurrentSeqIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetCurrentSeqIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(idgen.VectorIdVal)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCurrentSeqIdListResult) GetSuccess() *idgen.VectorIdVal {
	if !p.IsSetSuccess() {
		return GetCurrentSeqIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCurrentSeqIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*idgen.VectorIdVal)
}

func (p *GetCurrentSeqIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCurrentSeqIdListResult) GetResult() interface{} {
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

func (p *kClient) IdgenNextId(ctx context.Context, req *idgen.TLIdgenNextId) (r *tg.Int64, err error) {
	// var _args NextIdArgs
	// _args.Req = req
	// var _result NextIdResult

	_result := new(tg.Int64)

	if err = p.c.Call(ctx, "/idgen.RPCIdgen/idgen.nextId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) IdgenNextIds(ctx context.Context, req *idgen.TLIdgenNextIds) (r *idgen.VectorLong, err error) {
	// var _args NextIdsArgs
	// _args.Req = req
	// var _result NextIdsResult

	_result := new(idgen.VectorLong)

	if err = p.c.Call(ctx, "/idgen.RPCIdgen/idgen.nextIds", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) IdgenGetCurrentSeqId(ctx context.Context, req *idgen.TLIdgenGetCurrentSeqId) (r *tg.Int64, err error) {
	// var _args GetCurrentSeqIdArgs
	// _args.Req = req
	// var _result GetCurrentSeqIdResult

	_result := new(tg.Int64)

	if err = p.c.Call(ctx, "/idgen.RPCIdgen/idgen.getCurrentSeqId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) IdgenSetCurrentSeqId(ctx context.Context, req *idgen.TLIdgenSetCurrentSeqId) (r *tg.Bool, err error) {
	// var _args SetCurrentSeqIdArgs
	// _args.Req = req
	// var _result SetCurrentSeqIdResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/idgen.RPCIdgen/idgen.setCurrentSeqId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) IdgenGetNextSeqId(ctx context.Context, req *idgen.TLIdgenGetNextSeqId) (r *tg.Int64, err error) {
	// var _args GetNextSeqIdArgs
	// _args.Req = req
	// var _result GetNextSeqIdResult

	_result := new(tg.Int64)

	if err = p.c.Call(ctx, "/idgen.RPCIdgen/idgen.getNextSeqId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) IdgenGetNextNSeqId(ctx context.Context, req *idgen.TLIdgenGetNextNSeqId) (r *tg.Int64, err error) {
	// var _args GetNextNSeqIdArgs
	// _args.Req = req
	// var _result GetNextNSeqIdResult

	_result := new(tg.Int64)

	if err = p.c.Call(ctx, "/idgen.RPCIdgen/idgen.getNextNSeqId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) IdgenGetNextIdValList(ctx context.Context, req *idgen.TLIdgenGetNextIdValList) (r *idgen.VectorIdVal, err error) {
	// var _args GetNextIdValListArgs
	// _args.Req = req
	// var _result GetNextIdValListResult

	_result := new(idgen.VectorIdVal)

	if err = p.c.Call(ctx, "/idgen.RPCIdgen/idgen.getNextIdValList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) IdgenGetCurrentSeqIdList(ctx context.Context, req *idgen.TLIdgenGetCurrentSeqIdList) (r *idgen.VectorIdVal, err error) {
	// var _args GetCurrentSeqIdListArgs
	// _args.Req = req
	// var _result GetCurrentSeqIdListResult

	_result := new(idgen.VectorIdVal)

	if err = p.c.Call(ctx, "/idgen.RPCIdgen/idgen.getCurrentSeqIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
