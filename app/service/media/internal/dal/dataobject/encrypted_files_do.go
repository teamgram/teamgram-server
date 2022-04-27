/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type EncryptedFilesDO struct {
	Id              int64  `db:"id"`
	EncryptedFileId int64  `db:"encrypted_file_id"`
	AccessHash      int64  `db:"access_hash"`
	DcId            int32  `db:"dc_id"`
	FileSize        int32  `db:"file_size"`
	KeyFingerprint  int32  `db:"key_fingerprint"`
	Md5Checksum     string `db:"md5_checksum"`
	FilePath        string `db:"file_path"`
}
