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
	"github.com/nebula-chat/chatengine/service/document/biz/dal/dao/mysql_dao"
	"sync"
)

const (
	DB_MASTER = "immaster"
	DB_SLAVE  = "imslave"
)

type MysqlDAOList struct {
	// FilePartsDAO  *mysql_dao.FilePartsDAO
	// FilesDAO      *mysql_dao.FilesDAO
	PhotoDatasDAO *mysql_dao.PhotoDatasDAO
	DocumentsDAO  *mysql_dao.DocumentsDAO
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

		daoList := &MysqlDAOList{}

		// daoList.FilePartsDAO = mysql_dao.NewFilePartsDAO(v)
		// daoList.FilesDAO = mysql_dao.NewFilesDAO(v)
		daoList.PhotoDatasDAO = mysql_dao.NewPhotoDatasDAO(v)
		daoList.DocumentsDAO = mysql_dao.NewDocumentsDAO(v)

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

/*
func GetFilePartsDAO(dbName string) (dao *mysql_dao.FilePartsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.FilePartsDAO
	}
	return
}

func GetFilesDAO(dbName string) (dao *mysql_dao.FilesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.FilesDAO
	}
	return
}
*/

func GetPhotoDatasDAO(dbName string) (dao *mysql_dao.PhotoDatasDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.PhotoDatasDAO
	}
	return
}

func GetDocumentsDAO(dbName string) (dao *mysql_dao.DocumentsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.DocumentsDAO
	}
	return
}
