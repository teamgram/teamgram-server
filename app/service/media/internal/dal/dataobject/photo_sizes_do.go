/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type PhotoSizesDO struct {
	Id            int64  `db:"id" json:"id"`
	PhotoSizeId   int64  `db:"photo_size_id" json:"photo_size_id"`
	SizeType      string `db:"size_type" json:"size_type"`
	VolumeId      int64  `db:"volume_id" json:"volume_id"`
	LocalId       int32  `db:"local_id" json:"local_id"`
	Secret        int64  `db:"secret" json:"secret"`
	Width         int32  `db:"width" json:"width"`
	Height        int32  `db:"height" json:"height"`
	FileSize      int32  `db:"file_size" json:"file_size"`
	FilePath      string `db:"file_path" json:"file_path"`
	HasStripped   bool   `db:"has_stripped" json:"has_stripped"`
	StrippedBytes string `db:"stripped_bytes" json:"stripped_bytes"`
	CachedType    int32  `db:"cached_type" json:"cached_type"`
	CachedBytes   string `db:"cached_bytes" json:"cached_bytes"`
}
