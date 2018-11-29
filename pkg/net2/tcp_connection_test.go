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

package net2

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type TestConnCodec struct {
	receiver chan interface{}
	sender   chan interface{}
	isClosed bool
	sync.RWMutex
}

func (c *TestConnCodec) Send(msg interface{}) error {
	c.RLock()
	defer c.RUnlock()

	if c.isClosed {
		return errors.New(`codec is closed`)
	}

	c.sender <- msg
	return nil
}

func (c *TestConnCodec) Receive() (interface{}, error) {
	if r, ok := <-c.receiver; ok {
		return r, nil
	}
	return nil, errors.New(`error`)
}

func (c *TestConnCodec) Close() error {
	c.Lock()
	defer c.Unlock()

	if c.receiver != nil {
		close(c.receiver)
	}
	if c.sender != nil {
		close(c.sender)
	}
	c.isClosed = true
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
type TestConnSimulation struct {
	name         string
	receivedChan chan interface{}
	sendChan     chan interface{}
	errChan      chan error
	connChanSize int
	numberOfMsg  int
}

func (tc *TestConnSimulation) simulateSend() (result []string, err error) {
	mockCodec := &TestConnCodec{sender: tc.sendChan}
	tc.errChan = make(chan error)

	conn := NewTcpConnection(tc.name, nil, tc.connChanSize, mockCodec, nil)

	for i := 0; i < tc.numberOfMsg; i++ {
		go func(n int) {
			if err = conn.Send(fmt.Sprintf(`msg%d`, n)); err != nil {
				tc.errChan <- err
			}
		}(i)
	}

	cnt := tc.numberOfMsg
	for {
		select {
		case m, _ := <-tc.sendChan:
			result = append(result, m.(string))
			cnt--
			if cnt == 0 {
				return
			}
		case m, _ := <-tc.errChan:
			return nil, m
		}
	}

	return
}

//////////////////////////////////////////////////////////////////////////////////////////
func TestSend(t *testing.T) {
	sendChan := make(chan interface{})

	simulation := &TestConnSimulation{
		name:         "",
		sendChan:     sendChan,
		numberOfMsg:  100,
		connChanSize: 100,
	}

	result, err := simulation.simulateSend()
	if err != nil {
		t.Error(err)
	}

	expected := make(map[string]bool)
	for i := 0; i < simulation.numberOfMsg; i++ {
		expected[fmt.Sprintf(`msg%d`, i)] = false
	}

	for _, v := range result {
		expected[v] = true
	}

	for k, v := range expected {
		if v == false {
			t.Errorf("message(%s) does not sent", k)
		}
	}
}

func TestConnectionError(t *testing.T) {
	sendChan := make(chan interface{})

	simulation := &TestConnSimulation{
		name:         "",
		sendChan:     sendChan,
		numberOfMsg:  2 * 1024,
		connChanSize: 1024,
	}

	_, err := simulation.simulateSend()

	if err != ConnectionClosedError && err != ConnectionBlockedError && err != nil {
		t.Errorf(`expected ConnectionClosedError or ConnectionBlockedError but get --> %s`, err.Error())
	}
}
