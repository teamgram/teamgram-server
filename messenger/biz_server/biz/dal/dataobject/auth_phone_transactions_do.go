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

type AuthPhoneTransactionsDO struct {
	Id               int64  `db:"id"`
	AuthKeyId        int64  `db:"auth_key_id"`
	PhoneNumber      string `db:"phone_number"`
	Code             string `db:"code"`
	CodeExpired      int32  `db:"code_expired"`
	CodeMsgId        string `db:"code_msg_id"`
	TransactionHash  string `db:"transaction_hash"`
	SentCodeType     int8   `db:"sent_code_type"`
	FlashCallPattern string `db:"flash_call_pattern"`
	NextCodeType     int8   `db:"next_code_type"`
	State            int8   `db:"state"`
	ApiId            int32  `db:"api_id"`
	ApiHash          string `db:"api_hash"`
	Attempts         int32  `db:"attempts"`
	CreatedTime      int64  `db:"created_time"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
	IsDeleted        int8   `db:"is_deleted"`
}
