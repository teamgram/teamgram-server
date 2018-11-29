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

package dao

import (
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebula-chat/chatengine/messenger/sync/biz/dal/dao/mysql_dao"
	"sync"
)

const (
	DB_MASTER = "immaster"
	DB_SLAVE  = "imslave"
)

type MysqlDAOList struct {
	AuthSeqUpdatesDAO    *mysql_dao.AuthSeqUpdatesDAO
	AuthUpdatesStateDAO  *mysql_dao.AuthUpdatesStateDAO
	ChannelPtsUpdatesDAO *mysql_dao.ChannelPtsUpdatesDAO
	UserPtsUpdatesDAO    *mysql_dao.UserPtsUpdatesDAO
	UserQtsUpdatesDAO    *mysql_dao.UserQtsUpdatesDAO
}

// TODO(@benqi): 一主多从
type MysqlDAOManager struct {
	daoListMap map[string]*MysqlDAOList
}

var mysqlDAOManager = &MysqlDAOManager{make(map[string]*MysqlDAOList)}

func InstallMysqlDAOManager(clients sync.Map /*map[string]*sqlx.DB*/) {
	clients.Range(func(key, value interface{}) bool {
		k, _ := key.(string)
		v, _ := value.(*sqlx.DB)

		daoList := &MysqlDAOList{
			AuthSeqUpdatesDAO:    mysql_dao.NewAuthSeqUpdatesDAO(v),
			AuthUpdatesStateDAO:  mysql_dao.NewAuthUpdatesStateDAO(v),
			ChannelPtsUpdatesDAO: mysql_dao.NewChannelPtsUpdatesDAO(v),
			UserPtsUpdatesDAO:    mysql_dao.NewUserPtsUpdatesDAO(v),
			UserQtsUpdatesDAO:    mysql_dao.NewUserQtsUpdatesDAO(v),
		}

		mysqlDAOManager.daoListMap[k] = daoList
		return true
	})
}

func GetMysqlDAOListMap() map[string]*MysqlDAOList {
	return mysqlDAOManager.daoListMap
}

func GetMysqlDAOList(dbName string) (daoList *MysqlDAOList) {
	daoList, ok := mysqlDAOManager.daoListMap[dbName]
	if !ok {
		glog.Errorf("GetMysqlDAOList - Not found daoList: %s", dbName)
	}
	return
}

func GetAuthSeqUpdatesDAO(dbName string) (dao *mysql_dao.AuthSeqUpdatesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthSeqUpdatesDAO
	}
	return
}

func GetAuthUpdatesStateDAO(dbName string) (dao *mysql_dao.AuthUpdatesStateDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthUpdatesStateDAO
	}
	return
}

func GetChannelPtsUpdatesDAO(dbName string) (dao *mysql_dao.ChannelPtsUpdatesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.ChannelPtsUpdatesDAO
	}
	return
}

func GetUserPtsUpdatesDAO(dbName string) (dao *mysql_dao.UserPtsUpdatesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserPtsUpdatesDAO
	}
	return
}

func GetUserQtsUpdatesDAO(dbName string) (dao *mysql_dao.UserQtsUpdatesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.UserQtsUpdatesDAO
	}
	return
}
