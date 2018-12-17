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

// Copy from https://github.com/liyue201/grpc-lb
//

package load_balancer

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

type RandomSelector struct {
	baseSelector
	r *rand.Rand
}

func NewRandomSelector() Selector {
	return &RandomSelector{
		r:            rand.New(rand.NewSource(time.Now().UnixNano())),
		baseSelector: baseSelector{addrMap: make(map[string]*AddrInfo)},
	}
}

func (r *RandomSelector) Get(ctx context.Context) (addr grpc.Address, err error) {
	if len(r.addrs) == 0 {
		return addr, errors.New("addr list is emtpy")
	}

	size := len(r.addrs)
	idx := r.r.Int() % size

	for i := 0; i < size; i++ {
		addr := r.addrs[(idx+i)%size]
		if addrInfo, ok := r.addrMap[addr]; ok {
			if addrInfo.connected {
				addrInfo.load++
				return addrInfo.addr, nil
			}
		}
	}
	return addr, NoAvailableAddressErr
}
