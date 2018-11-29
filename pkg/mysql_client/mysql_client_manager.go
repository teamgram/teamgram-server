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
	"sync"
)

type MysqlClientManager struct {
	mysqlClients sync.Map
}

var mysqlClients = &MysqlClientManager{}

func InstallMysqlClientManager(configs []MySQLConfig) {
	for _, config := range configs {
		client := NewSqlxDB(&config)
		if client == nil {
			err := fmt.Errorf("installMysqlClientManager - NewSqlxDB {%v} error", config)
			panic(err)
		}

		if config.Name == "" {
			err := fmt.Errorf("installMysqlClientManager - config error: config.Name is empty")
			panic(err)
		}
		if val, ok := mysqlClients.mysqlClients.Load(config.Name); ok {
			err := fmt.Errorf("installMysqlClientManager - config error: dublicated config.Name {%v}", val)
			panic(err)
		}
		mysqlClients.mysqlClients.Store(config.Name, client)
	}
}

func GetMysqlClient(dbName string) (client *sqlx.DB) {
	if val, ok := mysqlClients.mysqlClients.Load(dbName); ok {
		if client, ok = val.(*sqlx.DB); ok {
			return
		}
	}

	glog.Errorf("getMysqlClient - Not found client: %s", dbName)
	return
}

func GetMysqlClientManager() sync.Map {
	return mysqlClients.mysqlClients
}
