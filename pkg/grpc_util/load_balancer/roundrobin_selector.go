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
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type RoundRobinSelector struct {
	baseSelector
	next int
}

func NewRoundRobinSelector() Selector {
	return &RoundRobinSelector{
		next:         0,
		baseSelector: baseSelector{addrMap: make(map[string]*AddrInfo)},
	}
}

func (r *RoundRobinSelector) Get(ctx context.Context) (addr grpc.Address, err error) {
	if len(r.addrs) == 0 {
		err = AddrListEmptyErr
		return
	}

	if r.next >= len(r.addrs) {
		r.next = 0
	}
	next := r.next
	for {
		a := r.addrs[next]
		next = (next + 1) % len(r.addrs)

		if addrInfo, ok := r.addrMap[a]; ok {
			if addrInfo.connected {
				addr = addrInfo.addr
				addrInfo.load++
				r.next = next
				return
			}
			if next == r.next {
				// Has iterated all the possible address but none is connected.
				addr = addrInfo.addr
				addrInfo.load++
				r.next = next
				return
			}
		}
	}
}
