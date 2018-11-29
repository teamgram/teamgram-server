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

// author: @yumcoder(https://github.com/yumcoder-platform)
// @benqi copy code from telegram group [Telegramd](https://t.me/joinchat/D8b0DRJiuH8EcIHNZQmCxQ)
//

package mysql_client

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestMysqlPool(t *testing.T) {
	mySQLConfig := &MySQLConfig{
		DSN:    "root:1@/nebulaim?charset=utf8",
		Active: 200,
		Idle:   100,
	}
	mysql := NewSqlxDB(mySQLConfig)
	defer mysql.Close()
	var query = "select id,message_data from messages"

	r, err := mysql.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)

	var wg sync.WaitGroup
	var (
		id           int
		message_data string
	)

	start := time.Now()
	for i := 0; i < mySQLConfig.Active-1; i++ {
		wg.Add(1)
		go func(n int) {
			rows, err := mysql.Query(query)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&id, &message_data)
				if err != nil {
					log.Fatal(err)
				}
				//log.Println(id, message_data)
			}
			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("time: %s", elapsed)
}
