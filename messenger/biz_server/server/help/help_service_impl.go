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

package help

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"encoding/json"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/util"
	"io/ioutil"
)

const (
	CONFIG_FILE = "/config.json"

	// date = 1509066502,    2017/10/27 09:08:22
	// expires = 1509070295, 2017/10/27 10:11:35
	EXPIRES_TIMEOUT = 3600 // 超时时间设置为3600秒

	// support user: @benqi
	SUPPORT_USER_ID = 2
)

var config mtproto.TLConfig

func init() {
	configFilePath := util.GetWorkingDirectory() + "/" + CONFIG_FILE
	configData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
		return
	}

	err = json.Unmarshal([]byte(configData), &config)
	if err != nil {
		panic(err)
		return
	}
}

type HelpServiceImpl struct {
}

func NewHelpServiceImpl(models []core.CoreModel) *HelpServiceImpl {
	impl := &HelpServiceImpl{}

	for _, m := range models {
		switch m.(type) {
		}
	}

	return impl
}
