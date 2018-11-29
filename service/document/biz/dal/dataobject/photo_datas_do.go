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

type PhotoDatasDO struct {
	Id         int32  `db:"id"`
	PhotoId    int64  `db:"photo_id"`
	PhotoType  int8   `db:"photo_type"`
	DcId       int32  `db:"dc_id"`
	VolumeId   int64  `db:"volume_id"`
	LocalId    int32  `db:"local_id"`
	AccessHash int64  `db:"access_hash"`
	Width      int32  `db:"width"`
	Height     int32  `db:"height"`
	FileSize   int32  `db:"file_size"`
	FilePath   string `db:"file_path"`
	Ext        string `db:"ext"`
	FileId     int64  `db:"file_id"`
	State      int8   `db:"state"`
	CreatedAt  string `db:"created_at"`
}
