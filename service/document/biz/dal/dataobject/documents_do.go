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

type DocumentsDO struct {
	Id               int64  `db:"id"`
	DocumentId       int64  `db:"document_id"`
	AccessHash       int64  `db:"access_hash"`
	DcId             int32  `db:"dc_id"`
	FilePath         string `db:"file_path"`
	FileSize         int32  `db:"file_size"`
	UploadedFileName string `db:"uploaded_file_name"`
	Ext              string `db:"ext"`
	MimeType         string `db:"mime_type"`
	ThumbId          int64  `db:"thumb_id"`
	Version          int32  `db:"version"`
	Attributes       string `db:"attributes"`
	FileId           int64  `db:"file_id"`
	State            int8   `db:"state"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
