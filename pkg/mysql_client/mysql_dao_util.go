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

package mysql_client

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"strings"
)

const (
	DB_OK				=  0
	DB_DUPLICATE 		= -1
	DB_IN_VALUES_EMPTY 	= -2
	DB_ERROR            = -3
)

type CommonDAO struct {
	db *sqlx.DB
}

func NewCommonDAO(db *sqlx.DB) *CommonDAO {
	return &CommonDAO{db}
}

// 检查是否存在
func (dao *CommonDAO) CheckExists(table string, params map[string]interface{}) bool {
	if len(params) == 0 {
		glog.Errorf("CheckExists - [%s] error: params empty!", table)
		return false
	}

	names := make([]string, 0, len(params))
	for k, _ := range params {
		names = append(names, k+" = :"+k)
		// glog.Info("k: ", k, ", v: ", v)
	}
	sql := fmt.Sprintf("SELECT 1 FROM %s WHERE %s LIMIT 1", table, strings.Join(names, " AND "))
	// glog.Info("checkExists - sql: ", sql, ", params: ", params)
	rows, err := dao.db.NamedQuery(sql, params)
	if err != nil {
		glog.Errorf("CheckExists - [%s] error: %s", table, err)
		return false
	}

	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}

func (dao *CommonDAO) CalcSize(table string, params map[string]interface{}) int {
	if len(params) == 0 {
		glog.Errorf("calcSize - [%s] error: params empty!", table)
		return -1
	}

	names := make([]string, 0, len(params))
	for k, _ := range params {
		names = append(names, k+" = :"+k)
		// glog.Info("k: ", k, ", v: ", v)
	}
	sql := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", table, strings.Join(names, " AND "))

	var count int
	err := dao.db.Get(&count, sql, params)
	if err != nil {
		glog.Errorf("calcSize - [%s] error: %s", sql, err)
		return -1
	}
	return count
}

////////////////////////////////////////////////////////////////////////////////////////////////////
//
//func (dao *CommonDAO) InsertOrUpdate(table string, params map[string]interface{}) bool {
//	return true
//}
//
//func (dao *CommonDAO) GetOrInsert(table string, params map[string]interface{}) bool {
//	return true
//}

////////////////////////////////////////////////////////////////////////////////////////////////////
func InsertIgonre() int {
	return 0
}

func InsertIgonreReturnLastInsertId() int64 {
	return 0
}

func InsertOrUpdate() int {
	return 0
}

func InsertOrUpdateReturnLastInsertId() int64 {
	return 0
}

func Insert() int {
	return 0
}

func InsertReturnLastInertId() int64 {
	return 0
}
