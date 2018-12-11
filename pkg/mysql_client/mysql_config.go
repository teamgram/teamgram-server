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
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"time"
)

type MySQLConfig struct {
	Name   string // for trace
	DSN    string // data source name
	Active int    // pool
	Idle   int    // pool
}

func NewSqlxDB(c *MySQLConfig) (db *sqlx.DB) {
	db, err := sqlx.Connect("mysql", c.DSN)
	if err != nil {
		glog.Errorf("Connect db error: %s", err)
	}

	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)

	// 避免使用服务器断开超时的空闲连接 mysql空闲连接超时时间默认为8小时
	// thanks @hns_space(space Mars)
	db.SetConnMaxLifetime(time.Minute * 10)
	return
}
