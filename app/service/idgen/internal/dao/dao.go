// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package dao

import (
	"log"

	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/config"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

type Dao struct {
	*snowflake.Node
	KV kv.Store
}

func New(c config.Config) *Dao {
	var (
		err error
		d   = new(Dao)
	)

	d.Node, err = snowflake.NewNode(c.NodeId)
	if err != nil {
		log.Fatal("new snowflake node error: ", err)
	}
	d.KV = kv.NewStore(c.SeqIDGen)

	return d
}
