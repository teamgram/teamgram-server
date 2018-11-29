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

package idgen

import (
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/golang/glog"
)

type SnowflakeUUIDGen struct {
	idgen *snowflake.Node
}

func snowflakeUUIDGenInstance() UUIDGen {
	cli := &SnowflakeUUIDGen{}
	return cli
}

func NewSnowflakeUUIDGen(serverId int) *SnowflakeUUIDGen {
	idgen, err := snowflake.NewNode(int64(serverId))
	if err != nil {
		glog.Fatal(err)
	}
	return &SnowflakeUUIDGen{idgen}
}

func (id *SnowflakeUUIDGen) Initialize(config string) error {
	serverId, err := util.StringToInt64(config)
	if err != nil {
		glog.Error("start id generator error: ", err)
		return err
	}

	// TODO(@benqi): parse node from config
	id.idgen, err = snowflake.NewNode(serverId)
	if err != nil {
		glog.Error("start id generator error: ", err)
	}
	return err
}

func (id *SnowflakeUUIDGen) GetUUID() (int64, error) {
	var err error
	if id.idgen == nil {
		err = errors.New("idgen not init")
		glog.Error(err)
		return 0, err
	}
	return id.idgen.Generate().Int64(), nil
}

func init() {
	UUIDGenRegister("snowflake", snowflakeUUIDGenInstance)
}
