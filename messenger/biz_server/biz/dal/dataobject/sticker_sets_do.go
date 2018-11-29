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

type StickerSetsDO struct {
	Id           int32  `db:"id"`
	StickerSetId int64  `db:"sticker_set_id"`
	AccessHash   int64  `db:"access_hash"`
	Title        string `db:"title"`
	ShortName    string `db:"short_name"`
	Count        int32  `db:"count"`
	Hash         int32  `db:"hash"`
	Official     int8   `db:"official"`
	Mask         int8   `db:"mask"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}
