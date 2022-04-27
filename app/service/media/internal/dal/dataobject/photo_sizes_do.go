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

type PhotoSizesDO struct {
	Id            int64  `db:"id"`
	PhotoSizeId   int64  `db:"photo_size_id"`
	SizeType      string `db:"size_type"`
	VolumeId      int64  `db:"volume_id"`
	LocalId       int32  `db:"local_id"`
	Secret        int64  `db:"secret"`
	Width         int32  `db:"width"`
	Height        int32  `db:"height"`
	FileSize      int32  `db:"file_size"`
	FilePath      string `db:"file_path"`
	HasStripped   bool   `db:"has_stripped"`
	StrippedBytes string `db:"stripped_bytes"`
	CachedType    int32  `db:"cached_type"`
	CachedBytes   string `db:"cached_bytes"`
}
