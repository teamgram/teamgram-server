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

package util

import (
	"flag"
	"github.com/golang/glog"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var GAppInstance AppInstance

func init() {
	flag.Parse()
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type AppInstance interface {
	Initialize() error
	RunLoop()
	Destroy()
}

var ch = make(chan os.Signal, 1)

func DoMainAppInstance(instance AppInstance) {
	rand.Seed(time.Now().UnixNano())

	if instance == nil {
		// panic("instance is nil!!!!")
		glog.Errorf("instance is nil, will exit.")
		return
	}

	// global
	GAppInstance = instance

	glog.Info("instance initialize...")
	err := instance.Initialize()
	if err != nil {
		glog.Infof("instance initialize error: {%v}", err)
		return
	}

	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	glog.Info("instance run_loop...")
	go instance.RunLoop()

	// fmt.Printf("%d", os.Getpid())
	glog.Info("wait quit...")

	s2 := <-ch
	if i, ok := s2.(syscall.Signal); ok {
		glog.Infof("instance recv os.Exit(%d) signal...", i)
	} else {
		glog.Info("instance exit...", i)
	}

	instance.Destroy()
	glog.Info("instance quited!")
}

func QuitAppInstance() {
	/*notifier := make(chan os.Signal, 1)
	signal.Stop(notifier)*/
	ch <- syscall.SIGQUIT
}
