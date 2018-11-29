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
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/naming"
	"sync"
)

var DefaultSelector = NewRandomSelector()

type AddrInfo struct {
	addr      grpc.Address
	weight    int    //load weigth
	load      uint64 //current number of requests
	connected bool
}

type balancer struct {
	r        naming.Resolver
	w        naming.Watcher
	selector Selector
	mu       sync.Mutex
	addrCh   chan []grpc.Address // the channel to notify gRPC internals the list of addresses the client should connect to.
	waitCh   chan struct{}       // the channel to block when there is no connected address available
	done     bool                // The Balancer is closed.
}

func NewBalancer(r naming.Resolver, selector Selector) grpc.Balancer {
	if selector == nil {
		selector = DefaultSelector
	}
	return &balancer{r: r, selector: selector}
}

func (b *balancer) watchAddrUpdates() error {
	updates, err := b.w.Next()
	if err != nil {
		grpclog.Printf("grpc: the naming watcher stops working due to %v.\n", err)
		return err
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, update := range updates {
		addr := grpc.Address{
			Addr:     update.Addr,
			Metadata: update.Metadata,
		}
		switch update.Op {
		case naming.Add:
			b.selector.Add(addr)
		case naming.Delete:
			b.selector.Delete(addr)
		default:
			grpclog.Println("Unknown update.Op ", update.Op)
		}
	}
	if b.done {
		return grpc.ErrClientConnClosing
	}
	select {
	case <-b.addrCh:
	default:
	}

	addrs := b.selector.AddrList()
	b.addrCh <- addrs
	return nil
}

func (b *balancer) Start(target string, config grpc.BalancerConfig) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.done {
		return grpc.ErrClientConnClosing
	}
	if b.r == nil {
		return nil
	}
	w, err := b.r.Resolve(target)
	if err != nil {
		return err
	}
	b.w = w
	b.addrCh = make(chan []grpc.Address, 1)
	go func() {
		for {
			if err := b.watchAddrUpdates(); err != nil {
				return
			}
		}
	}()
	return nil
}

// Up sets the connected state of addr and sends notification if there are pending
// Get() calls.
func (b *balancer) Up(addr grpc.Address) func(error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	cnt, connected := b.selector.Up(addr)
	if connected {
		return func(err error) {
			b.down(addr, err)
		}
	}

	// addr is only one which is connected. Notify the Get() callers who are blocking.
	if cnt == 1 && b.waitCh != nil {
		close(b.waitCh)
		b.waitCh = nil
	}
	return func(err error) {
		b.down(addr, err)
	}
}

// down unsets the connected state of addr.
func (b *balancer) down(addr grpc.Address, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.selector.Down(addr)
}

// Get returns the next addr in the rotation.
func (b *balancer) Get(ctx context.Context, opts grpc.BalancerGetOptions) (addr grpc.Address, put func(), err error) {
	var ch chan struct{}
	b.mu.Lock()
	if b.done {
		b.mu.Unlock()
		err = grpc.ErrClientConnClosing
		return
	}

	addr, err = b.selector.Get(ctx)
	if err == nil {
		b.mu.Unlock()

		put = func() {
			b.selector.Put(addr.Addr)
		}

		return
	}

	// Wait on b.waitCh for non-failfast RPCs.
	if b.waitCh == nil {
		ch = make(chan struct{})
		b.waitCh = ch
	} else {
		ch = b.waitCh
	}
	b.mu.Unlock()
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		case <-ch:
			b.mu.Lock()
			if b.done {
				b.mu.Unlock()
				err = grpc.ErrClientConnClosing
				return
			}

			addr, err = b.selector.Get(ctx)
			if err == nil {
				put = func() {
					b.selector.Put(addr.Addr)
				}
				b.mu.Unlock()
				return
			}

			// The newly added addr got removed by Down() again.
			if b.waitCh == nil {
				ch = make(chan struct{})
				b.waitCh = ch
			} else {
				ch = b.waitCh
			}
			b.mu.Unlock()
		}
	}
}

func (b *balancer) Notify() <-chan []grpc.Address {
	return b.addrCh
}

func (b *balancer) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.done = true
	if b.w != nil {
		b.w.Close()
	}
	if b.waitCh != nil {
		close(b.waitCh)
		b.waitCh = nil
	}
	if b.addrCh != nil {
		close(b.addrCh)
	}
	return nil
}
