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

package photo

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/mysql_client"
	"github.com/nebula-chat/chatengine/service/document/biz/dal/dao/mysql_dao"
	"github.com/nebula-chat/chatengine/service/idgen/client"
)

type photosDAO struct {
	*mysql_dao.PhotoDatasDAO
	// *mysql_dao.FilePartsDAO
	idgen.UUIDGen
	//idgen.SeqIDGen
}

type PhotoModel struct {
	dao *photosDAO
}

func NewPhotoModel(serverId int32, dbName, redisName string) *PhotoModel {
	m := &PhotoModel{dao: &photosDAO{}}
	db := mysql_client.GetMysqlClient(dbName)
	if db == nil {
		glog.Fatal("not found db: ", dbName)
	}

	m.dao.PhotoDatasDAO = mysql_dao.NewPhotoDatasDAO(db)
	// m.dao.FilePartsDAO = mysql_dao.NewFilePartsDAO(db)

	var err error
	m.dao.UUIDGen, err = idgen.NewUUIDGen("snowflake", util.Int32ToString(serverId))
	if err != nil {
		glog.Fatal("uuidgen init error: ", err)
	}
	//m.dao.SeqIDGen, _ = idgen.NewSeqIDGen("redis", redisName)
	//if err != nil {
	//	glog.Fatal("seqidgen init error: ", err)
	//}
	return m
}
