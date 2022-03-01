// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package model

import (
	"github.com/teamgram/proto/mtproto"
)

// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
//

// CheckFileParts
// FILE_PARTS_INVALID - Invalid number of parts. The value is not between 1..3000
func CheckFileParts(fileParts int32) (err error) {
	if fileParts < 1 || fileParts > 3000 {
		err = mtproto.ErrFilePartsInvalid
	}
	return
}

// CheckFilePart
// FILE_PART_INVALID: The file part number is invalid. The value is not between 0 and 2,999.
func CheckFilePart(filePart int32) (err error) {
	if filePart < 0 || filePart > 2900 {
		err = mtproto.ErrFilePartInvalid
	}
	return
}

// CheckFilePartSize
// FILE_PART_SIZE_INVALID - 512KB cannot be evenly divided by part_size
func CheckFilePartSize(partSize int32) (err error) {
	// part_size % 1024 = 0 (divisible by 1KB)
	// 524288 % part_size = 0 (512KB must be evenly divisible by part_size)
	if partSize%1024 != 0 {
		err = mtproto.ErrFilePartLengthInvalid
	} else if 524288%partSize != 0 {
		err = mtproto.ErrFilePartSizeInvalid
	} else if partSize > 524288 {
		// FILE_PART_TOO_BIG: The size limit (512 KB) for the content of the file part has been exceeded
		err = mtproto.ErrFilePartTooBig
	}

	return
}
