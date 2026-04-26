/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package geoipservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/geoip.RPCGeoip/geoip.getCountryAndRegionByIp": kitex.NewMethodInfo(
		getCountryAndRegionByIpHandler,
		newGetCountryAndRegionByIpArgs,
		newGetCountryAndRegionByIpResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	geoipServiceServiceInfo                = NewServiceInfo()
	geoipServiceServiceInfoForClient       = NewServiceInfoForClient()
	geoipServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCGeoip", geoipServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCGeoip", geoipServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCGeoip", geoipServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return geoipServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return geoipServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return geoipServiceServiceInfoForClient
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
	serviceName := "RPCGeoip"
	handlerType := (*geoip.RPCGeoip)(nil)
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
		"PackageName": "geoip",
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

func getCountryAndRegionByIpHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetCountryAndRegionByIpArgs)
	realResult := result.(*GetCountryAndRegionByIpResult)
	success, err := handler.(geoip.RPCGeoip).GeoipGetCountryAndRegionByIp(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetCountryAndRegionByIpArgs() interface{} {
	return &GetCountryAndRegionByIpArgs{}
}

func newGetCountryAndRegionByIpResult() interface{} {
	return &GetCountryAndRegionByIpResult{}
}

type GetCountryAndRegionByIpArgs struct {
	Req *geoip.TLGeoipGetCountryAndRegionByIp
}

func (p *GetCountryAndRegionByIpArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCountryAndRegionByIpArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetCountryAndRegionByIpArgs) Unmarshal(in []byte) error {
	msg := new(geoip.TLGeoipGetCountryAndRegionByIp)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetCountryAndRegionByIpArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetCountryAndRegionByIpArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetCountryAndRegionByIpArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(geoip.TLGeoipGetCountryAndRegionByIp)
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

var GetCountryAndRegionByIpArgs_Req_DEFAULT *geoip.TLGeoipGetCountryAndRegionByIp

func (p *GetCountryAndRegionByIpArgs) GetReq() *geoip.TLGeoipGetCountryAndRegionByIp {
	if !p.IsSetReq() {
		return GetCountryAndRegionByIpArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCountryAndRegionByIpArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCountryAndRegionByIpResult struct {
	Success *geoip.Region
}

var GetCountryAndRegionByIpResult_Success_DEFAULT *geoip.Region

func (p *GetCountryAndRegionByIpResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCountryAndRegionByIpResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetCountryAndRegionByIpResult) Unmarshal(in []byte) error {
	msg := new(geoip.Region)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCountryAndRegionByIpResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetCountryAndRegionByIpResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetCountryAndRegionByIpResult) Decode(d *bin.Decoder) (err error) {
	msg := new(geoip.Region)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCountryAndRegionByIpResult) GetSuccess() *geoip.Region {
	if !p.IsSetSuccess() {
		return GetCountryAndRegionByIpResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCountryAndRegionByIpResult) SetSuccess(x interface{}) {
	p.Success = x.(*geoip.Region)
}

func (p *GetCountryAndRegionByIpResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCountryAndRegionByIpResult) GetResult() interface{} {
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

func (p *kClient) GeoipGetCountryAndRegionByIp(ctx context.Context, req *geoip.TLGeoipGetCountryAndRegionByIp) (r *geoip.Region, err error) {
	// var _args GetCountryAndRegionByIpArgs
	// _args.Req = req
	// var _result GetCountryAndRegionByIpResult

	_result := new(geoip.Region)

	if err = p.c.Call(ctx, "/geoip.RPCGeoip/geoip.getCountryAndRegionByIp", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
