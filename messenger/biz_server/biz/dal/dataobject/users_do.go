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

package dataobject

type UsersDO struct {
	Id                int32  `db:"id"`
	UserType          int8   `db:"user_type"`
	AccessHash        int64  `db:"access_hash"`
	FirstName         string `db:"first_name"`
	LastName          string `db:"last_name"`
	Username          string `db:"username"`
	Phone             string `db:"phone"`
	CountryCode       string `db:"country_code"`
	Verified          int8   `db:"verified"`
	About             string `db:"about"`
	State             int32  `db:"state"`
	IsBot             int8   `db:"is_bot"`
	AccountDaysTtl    int32  `db:"account_days_ttl"`
	Photos            string `db:"photos"`
	Min               int8   `db:"min"`
	Restricted        int8   `db:"restricted"`
	RestrictionReason string `db:"restriction_reason"`
	Deleted           int8   `db:"deleted"`
	DeleteReason      string `db:"delete_reason"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
}
