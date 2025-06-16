/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package premiumservice

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
	"help.getPremiumPromo": kitex.NewMethodInfo(
		helpGetPremiumPromoHandler,
		newHelpGetPremiumPromoArgs,
		newHelpGetPremiumPromoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"payments.assignAppStoreTransaction": kitex.NewMethodInfo(
		paymentsAssignAppStoreTransactionHandler,
		newPaymentsAssignAppStoreTransactionArgs,
		newPaymentsAssignAppStoreTransactionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"payments.assignPlayMarketTransaction": kitex.NewMethodInfo(
		paymentsAssignPlayMarketTransactionHandler,
		newPaymentsAssignPlayMarketTransactionArgs,
		newPaymentsAssignPlayMarketTransactionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"payments.canPurchasePremium": kitex.NewMethodInfo(
		paymentsCanPurchasePremiumHandler,
		newPaymentsCanPurchasePremiumArgs,
		newPaymentsCanPurchasePremiumResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	premiumServiceServiceInfo                = NewServiceInfo()
	premiumServiceServiceInfoForClient       = NewServiceInfoForClient()
	premiumServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return premiumServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return premiumServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return premiumServiceServiceInfoForClient
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
	serviceName := "RPCPremium"
	handlerType := (*tg.RPCPremium)(nil)
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
		"PackageName": "premium",
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

func helpGetPremiumPromoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*HelpGetPremiumPromoArgs)
	realResult := result.(*HelpGetPremiumPromoResult)
	success, err := handler.(tg.RPCPremium).HelpGetPremiumPromo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newHelpGetPremiumPromoArgs() interface{} {
	return &HelpGetPremiumPromoArgs{}
}

func newHelpGetPremiumPromoResult() interface{} {
	return &HelpGetPremiumPromoResult{}
}

type HelpGetPremiumPromoArgs struct {
	Req *tg.TLHelpGetPremiumPromo
}

func (p *HelpGetPremiumPromoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in HelpGetPremiumPromoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *HelpGetPremiumPromoArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLHelpGetPremiumPromo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *HelpGetPremiumPromoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in HelpGetPremiumPromoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *HelpGetPremiumPromoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLHelpGetPremiumPromo)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var HelpGetPremiumPromoArgs_Req_DEFAULT *tg.TLHelpGetPremiumPromo

func (p *HelpGetPremiumPromoArgs) GetReq() *tg.TLHelpGetPremiumPromo {
	if !p.IsSetReq() {
		return HelpGetPremiumPromoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelpGetPremiumPromoArgs) IsSetReq() bool {
	return p.Req != nil
}

type HelpGetPremiumPromoResult struct {
	Success *tg.HelpPremiumPromo
}

var HelpGetPremiumPromoResult_Success_DEFAULT *tg.HelpPremiumPromo

func (p *HelpGetPremiumPromoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in HelpGetPremiumPromoResult")
	}
	return json.Marshal(p.Success)
}

func (p *HelpGetPremiumPromoResult) Unmarshal(in []byte) error {
	msg := new(tg.HelpPremiumPromo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetPremiumPromoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in HelpGetPremiumPromoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *HelpGetPremiumPromoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.HelpPremiumPromo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelpGetPremiumPromoResult) GetSuccess() *tg.HelpPremiumPromo {
	if !p.IsSetSuccess() {
		return HelpGetPremiumPromoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelpGetPremiumPromoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.HelpPremiumPromo)
}

func (p *HelpGetPremiumPromoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelpGetPremiumPromoResult) GetResult() interface{} {
	return p.Success
}

func paymentsAssignAppStoreTransactionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PaymentsAssignAppStoreTransactionArgs)
	realResult := result.(*PaymentsAssignAppStoreTransactionResult)
	success, err := handler.(tg.RPCPremium).PaymentsAssignAppStoreTransaction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPaymentsAssignAppStoreTransactionArgs() interface{} {
	return &PaymentsAssignAppStoreTransactionArgs{}
}

func newPaymentsAssignAppStoreTransactionResult() interface{} {
	return &PaymentsAssignAppStoreTransactionResult{}
}

type PaymentsAssignAppStoreTransactionArgs struct {
	Req *tg.TLPaymentsAssignAppStoreTransaction
}

func (p *PaymentsAssignAppStoreTransactionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PaymentsAssignAppStoreTransactionArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PaymentsAssignAppStoreTransactionArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLPaymentsAssignAppStoreTransaction)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PaymentsAssignAppStoreTransactionArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PaymentsAssignAppStoreTransactionArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PaymentsAssignAppStoreTransactionArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLPaymentsAssignAppStoreTransaction)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PaymentsAssignAppStoreTransactionArgs_Req_DEFAULT *tg.TLPaymentsAssignAppStoreTransaction

func (p *PaymentsAssignAppStoreTransactionArgs) GetReq() *tg.TLPaymentsAssignAppStoreTransaction {
	if !p.IsSetReq() {
		return PaymentsAssignAppStoreTransactionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PaymentsAssignAppStoreTransactionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PaymentsAssignAppStoreTransactionResult struct {
	Success *tg.Updates
}

var PaymentsAssignAppStoreTransactionResult_Success_DEFAULT *tg.Updates

func (p *PaymentsAssignAppStoreTransactionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PaymentsAssignAppStoreTransactionResult")
	}
	return json.Marshal(p.Success)
}

func (p *PaymentsAssignAppStoreTransactionResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PaymentsAssignAppStoreTransactionResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PaymentsAssignAppStoreTransactionResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PaymentsAssignAppStoreTransactionResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PaymentsAssignAppStoreTransactionResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return PaymentsAssignAppStoreTransactionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PaymentsAssignAppStoreTransactionResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *PaymentsAssignAppStoreTransactionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentsAssignAppStoreTransactionResult) GetResult() interface{} {
	return p.Success
}

func paymentsAssignPlayMarketTransactionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PaymentsAssignPlayMarketTransactionArgs)
	realResult := result.(*PaymentsAssignPlayMarketTransactionResult)
	success, err := handler.(tg.RPCPremium).PaymentsAssignPlayMarketTransaction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPaymentsAssignPlayMarketTransactionArgs() interface{} {
	return &PaymentsAssignPlayMarketTransactionArgs{}
}

func newPaymentsAssignPlayMarketTransactionResult() interface{} {
	return &PaymentsAssignPlayMarketTransactionResult{}
}

type PaymentsAssignPlayMarketTransactionArgs struct {
	Req *tg.TLPaymentsAssignPlayMarketTransaction
}

func (p *PaymentsAssignPlayMarketTransactionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PaymentsAssignPlayMarketTransactionArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PaymentsAssignPlayMarketTransactionArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLPaymentsAssignPlayMarketTransaction)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PaymentsAssignPlayMarketTransactionArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PaymentsAssignPlayMarketTransactionArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PaymentsAssignPlayMarketTransactionArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLPaymentsAssignPlayMarketTransaction)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PaymentsAssignPlayMarketTransactionArgs_Req_DEFAULT *tg.TLPaymentsAssignPlayMarketTransaction

func (p *PaymentsAssignPlayMarketTransactionArgs) GetReq() *tg.TLPaymentsAssignPlayMarketTransaction {
	if !p.IsSetReq() {
		return PaymentsAssignPlayMarketTransactionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PaymentsAssignPlayMarketTransactionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PaymentsAssignPlayMarketTransactionResult struct {
	Success *tg.Updates
}

var PaymentsAssignPlayMarketTransactionResult_Success_DEFAULT *tg.Updates

func (p *PaymentsAssignPlayMarketTransactionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PaymentsAssignPlayMarketTransactionResult")
	}
	return json.Marshal(p.Success)
}

func (p *PaymentsAssignPlayMarketTransactionResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PaymentsAssignPlayMarketTransactionResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PaymentsAssignPlayMarketTransactionResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PaymentsAssignPlayMarketTransactionResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PaymentsAssignPlayMarketTransactionResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return PaymentsAssignPlayMarketTransactionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PaymentsAssignPlayMarketTransactionResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *PaymentsAssignPlayMarketTransactionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentsAssignPlayMarketTransactionResult) GetResult() interface{} {
	return p.Success
}

func paymentsCanPurchasePremiumHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*PaymentsCanPurchasePremiumArgs)
	realResult := result.(*PaymentsCanPurchasePremiumResult)
	success, err := handler.(tg.RPCPremium).PaymentsCanPurchasePremium(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newPaymentsCanPurchasePremiumArgs() interface{} {
	return &PaymentsCanPurchasePremiumArgs{}
}

func newPaymentsCanPurchasePremiumResult() interface{} {
	return &PaymentsCanPurchasePremiumResult{}
}

type PaymentsCanPurchasePremiumArgs struct {
	Req *tg.TLPaymentsCanPurchasePremium
}

func (p *PaymentsCanPurchasePremiumArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PaymentsCanPurchasePremiumArgs")
	}
	return json.Marshal(p.Req)
}

func (p *PaymentsCanPurchasePremiumArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLPaymentsCanPurchasePremium)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *PaymentsCanPurchasePremiumArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in PaymentsCanPurchasePremiumArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *PaymentsCanPurchasePremiumArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLPaymentsCanPurchasePremium)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var PaymentsCanPurchasePremiumArgs_Req_DEFAULT *tg.TLPaymentsCanPurchasePremium

func (p *PaymentsCanPurchasePremiumArgs) GetReq() *tg.TLPaymentsCanPurchasePremium {
	if !p.IsSetReq() {
		return PaymentsCanPurchasePremiumArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PaymentsCanPurchasePremiumArgs) IsSetReq() bool {
	return p.Req != nil
}

type PaymentsCanPurchasePremiumResult struct {
	Success *tg.Bool
}

var PaymentsCanPurchasePremiumResult_Success_DEFAULT *tg.Bool

func (p *PaymentsCanPurchasePremiumResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PaymentsCanPurchasePremiumResult")
	}
	return json.Marshal(p.Success)
}

func (p *PaymentsCanPurchasePremiumResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PaymentsCanPurchasePremiumResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in PaymentsCanPurchasePremiumResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *PaymentsCanPurchasePremiumResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PaymentsCanPurchasePremiumResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return PaymentsCanPurchasePremiumResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PaymentsCanPurchasePremiumResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *PaymentsCanPurchasePremiumResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentsCanPurchasePremiumResult) GetResult() interface{} {
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

func (p *kClient) HelpGetPremiumPromo(ctx context.Context, req *tg.TLHelpGetPremiumPromo) (r *tg.HelpPremiumPromo, err error) {
	var _args HelpGetPremiumPromoArgs
	_args.Req = req
	var _result HelpGetPremiumPromoResult
	if err = p.c.Call(ctx, "help.getPremiumPromo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PaymentsAssignAppStoreTransaction(ctx context.Context, req *tg.TLPaymentsAssignAppStoreTransaction) (r *tg.Updates, err error) {
	var _args PaymentsAssignAppStoreTransactionArgs
	_args.Req = req
	var _result PaymentsAssignAppStoreTransactionResult
	if err = p.c.Call(ctx, "payments.assignAppStoreTransaction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PaymentsAssignPlayMarketTransaction(ctx context.Context, req *tg.TLPaymentsAssignPlayMarketTransaction) (r *tg.Updates, err error) {
	var _args PaymentsAssignPlayMarketTransactionArgs
	_args.Req = req
	var _result PaymentsAssignPlayMarketTransactionResult
	if err = p.c.Call(ctx, "payments.assignPlayMarketTransaction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PaymentsCanPurchasePremium(ctx context.Context, req *tg.TLPaymentsCanPurchasePremium) (r *tg.Bool, err error) {
	var _args PaymentsCanPurchasePremiumArgs
	_args.Req = req
	var _result PaymentsCanPurchasePremiumResult
	if err = p.c.Call(ctx, "payments.canPurchasePremium", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
