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

package server

import (
	"fmt"
	"github.com/nebula-chat/chatengine/pkg/sync2"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type MockEntry struct {
	id       int
	refcount sync2.AtomicInt32
}

func (e *MockEntry) AddRef() {
	e.refcount.Add(1)
}

func (e *MockEntry) Release() int32 {
	return e.refcount.Add(-1)
}

func (e *MockEntry) TimerCallback() {
	fmt.Println("TimerCallback - ", e)
}

func TestTimingWheel(t *testing.T) {
	wheel := NewTimingWheel(8)
	wheel.Start()

	entries := make([]MockEntry, 100)

	for i := 0; i < 10; i++ {
		entries[i].id = i
		// wheel.AddTimer(&entries[i], rand.Intn(8))
	}

	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			// entries[i].id = i
			wheel.AddTimer(&entries[i], rand.Intn(8))
		}
		time.Sleep(time.Second)
	}

	time.Sleep(10 * time.Second)
}
