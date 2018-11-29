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

package mysql_dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestReflectTLObject(t *testing.T) {
	mysqlDsn := "root:@/chatengine?charset=utf8"

	db, err := sqlx.Connect("mysql", mysqlDsn)
	if err != nil {
		glog.Fatalf("Connect mysql %s error: %s", mysqlDsn, err)
		return
	}

	userDialogsDAO := NewUserDialogsDAO(db)

	vals := userDialogsDAO.SelectPinnedDialogs(1)
	fmt.Println(vals)
}
