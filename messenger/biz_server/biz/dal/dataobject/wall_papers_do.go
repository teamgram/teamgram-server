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

type WallPapersDO struct {
	Id        int32  `db:"id"`
	Type      int8   `db:"type"`
	Title     string `db:"title"`
	Color     int32  `db:"color"`
	BgColor   int32  `db:"bg_color"`
	PhotoId   int64  `db:"photo_id"`
	CreatedAt string `db:"created_at"`
	DeletedAt int64  `db:"deleted_at"`
}
