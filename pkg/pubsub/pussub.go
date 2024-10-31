// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package pubsub

import (
	"sync"
)

type PubSub struct {
	channels []chan struct{}
	lock     *sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		channels: make([]chan struct{}, 0),
		lock:     new(sync.RWMutex),
	}
}

func (p *PubSub) Subscribe() (<-chan struct{}, func()) {
	p.lock.Lock()
	defer p.lock.Unlock()

	c := make(chan struct{}, 1)
	p.channels = append(p.channels, c)
	return c, func() {
		p.lock.Lock()
		defer p.lock.Unlock()

		for i, channel := range p.channels {
			if channel == c {
				p.channels = append(p.channels[:i], p.channels[i+1:]...)
				close(c)
				return
			}
		}
	}
}

func (p *PubSub) Publish() {
	p.lock.RLock()
	defer p.lock.RUnlock()

	for _, channel := range p.channels {
		channel <- struct{}{}
	}
}
