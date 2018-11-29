// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package load_balancer

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strings"
)

type KetamaSelector struct {
	baseSelector
	hash      *Ketama
	ketamaKey string
}

var (
	DefaultKetamaKey  = "grpc-lb-ketama-key"
	KetamaKeyEmptyErr = errors.New("ketama key is empty")
)

func NewKetamaSelector(ketamaKey string) Selector {
	if ketamaKey == "" {
		ketamaKey = DefaultKetamaKey
	}
	return &KetamaSelector{
		hash:         NewKetama(10, nil),
		ketamaKey:    ketamaKey,
		baseSelector: baseSelector{addrMap: make(map[string]*AddrInfo)},
	}
}

func (s *KetamaSelector) wrapAddr(addr string, idx int) string {
	return fmt.Sprintf("%s-%d", addr, idx)
}

func (s *KetamaSelector) upWrapAddr(addr string) string {
	ss := strings.Split(addr, "-")
	return ss[0]
}

func (s *KetamaSelector) Add(addr grpc.Address) error {
	fmt.Println("add", addr.Addr)
	err := s.baseSelector.Add(addr)
	if err == nil {
		a, _ := s.addrMap[addr.Addr]
		for i := 0; i < a.weight; i++ {
			s.hash.Add(s.wrapAddr(addr.Addr, i))
		}
	}
	return err
}

func (s *KetamaSelector) Delete(addr grpc.Address) error {
	err := s.baseSelector.Delete(addr)
	if err == nil {
		a, ok := s.addrMap[addr.Addr]
		if ok {
			for i := 0; i < a.weight; i++ {
				s.hash.Remove(s.wrapAddr(addr.Addr, i))
			}
		}
	}
	return err
}

func (s *KetamaSelector) Get(ctx context.Context) (addr grpc.Address, err error) {
	if len(s.addrs) == 0 {
		err = AddrListEmptyErr
		return
	}
	key, ok := ctx.Value(s.ketamaKey).(string)
	if ok {
		targetAddr, ok := s.hash.Get(key)
		if ok {
			targetAddr = s.upWrapAddr(targetAddr)
			for _, v := range s.addrs {
				if v == targetAddr {
					if addrInfo, ok := s.addrMap[v]; ok {
						if addrInfo.connected {
							addrInfo.load++
							return addrInfo.addr, nil
						}
					}
				}
			}
		} else {
			err = AddrDoseNotExistErr
		}
	} else {
		err = KetamaKeyEmptyErr
	}

	return addr, NoAvailableAddressErr
}
