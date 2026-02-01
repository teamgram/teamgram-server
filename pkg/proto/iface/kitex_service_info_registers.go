// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package iface

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

var (
	kitexServiceInfoRegisters                = make(map[string]*serviceinfo.ServiceInfo)
	kitexServiceInfoForClientRegisters       = make(map[string]*serviceinfo.ServiceInfo)
	kitexServiceInfoForStreamClientRegisters = make(map[string]*serviceinfo.ServiceInfo)
)

func RegisterKitexServiceInfo(svcName string, svc *serviceinfo.ServiceInfo) {
	if _, ok := kitexServiceInfoRegisters[svcName]; ok {
		panic(fmt.Sprintf("kitexServiceInfoRegister service %s already exists", svcName))
	}

	kitexServiceInfoRegisters[svcName] = svc
}

func RegisterKitexServiceInfoForClient(svcName string, svc *serviceinfo.ServiceInfo) {
	if _, ok := kitexServiceInfoForClientRegisters[svcName]; ok {
		panic(fmt.Sprintf("kitexServiceInfoForClientRegister service %s already exists", svcName))
	}

	kitexServiceInfoForClientRegisters[svcName] = svc
}

func RegisterKitexServiceInfoForStreamClient(svcName string, svc *serviceinfo.ServiceInfo) {
	if _, ok := kitexServiceInfoForStreamClientRegisters[svcName]; ok {
		panic(fmt.Sprintf("kitexServiceInfoForStreamClientRegister service %s already exists", svcName))
	}

	kitexServiceInfoForStreamClientRegisters[svcName] = svc
}

func GetKitexServiceInfo(svcName string) *serviceinfo.ServiceInfo {
	svc, ok := kitexServiceInfoRegisters[svcName]
	if !ok {
		panic(fmt.Sprintf("kitexServiceInfoRegister service %s not found", svcName))
	}

	return svc
}

func GetKitexServiceInfoForClient(svcName string) *serviceinfo.ServiceInfo {
	svc, ok := kitexServiceInfoForClientRegisters[svcName]
	if !ok {
		panic(fmt.Sprintf("kitexServiceInfoForClientRegister service %s not found", svcName))
	}

	return svc
}

func GetKitexServiceInfoForStreamClient(svcName string) *serviceinfo.ServiceInfo {
	svc, ok := kitexServiceInfoForStreamClientRegisters[svcName]
	if !ok {
		panic(fmt.Sprintf("kitexServiceInfoForStreamClientRegister service %s not found", svcName))
	}

	return svc
}
