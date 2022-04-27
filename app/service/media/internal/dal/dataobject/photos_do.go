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

type PhotosDO struct {
	Id            int64  `db:"id"`
	PhotoId       int64  `db:"photo_id"`
	AccessHash    int64  `db:"access_hash"`
	HasStickers   bool   `db:"has_stickers"`
	DcId          int32  `db:"dc_id"`
	Date2         int64  `db:"date2"`
	HasVideo      bool   `db:"has_video"`
	SizeId        int64  `db:"size_id"`
	VideoSizeId   int64  `db:"video_size_id"`
	InputFileName string `db:"input_file_name"`
	Ext           string `db:"ext"`
}
