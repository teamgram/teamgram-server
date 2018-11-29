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

type AuthsDO struct {
	Id             int32  `db:"id"`
	AuthId         int64  `db:"auth_id"`
	ApiId          int32  `db:"api_id"`
	DeviceModel    string `db:"device_model"`
	SystemVersion  string `db:"system_version"`
	AppVersion     string `db:"app_version"`
	SystemLangCode string `db:"system_lang_code"`
	LangPack       string `db:"lang_pack"`
	LangCode       string `db:"lang_code"`
	ConnectionHash int64  `db:"connection_hash"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
	DeletedAt      string `db:"deleted_at"`
}
