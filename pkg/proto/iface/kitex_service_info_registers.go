// Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
