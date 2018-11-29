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

package util

import (
	"github.com/golang/glog"
	"sync"
	"testing"
)

type NullInstance struct {
	state int
	m     func()
}

func (e *NullInstance) Initialize() error {
	glog.Info("null instance initialize...")
	e.state = 1
	return nil
}

func (e *NullInstance) RunLoop() {
	glog.Info("null run_loop...")
	e.state = 2
	e.m()
}

func (e *NullInstance) Destroy() {
	glog.Info("null destroy...")
	e.state = 3
}

func TestRun1(t *testing.T) {
	instance := &NullInstance{}
	instance.m = func() {
		QuitAppInstance()
	}

	DoMainAppInstance(instance)

	result := instance.state
	expect := 3

	if result != expect {
		t.Error(`expect:`, expect, `result:`, result)
	}
}

func TestRun2(t *testing.T) {
	instance := &NullInstance{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	instance.m = func() {
		glog.Info("done...")
		wg.Done()
	}

	go DoMainAppInstance(instance)
	wg.Wait()
	result := instance.state
	expect := 2

	if result != expect {
		t.Error(`expect:`, expect, `result:`, result)
	}
}
