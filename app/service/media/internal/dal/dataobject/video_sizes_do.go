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

type VideoSizesDO struct {
	Id           int64   `db:"id" json:"id"`
	VideoSizeId  int64   `db:"video_size_id" json:"video_size_id"`
	SizeType     string  `db:"size_type" json:"size_type"`
	Width        int32   `db:"width" json:"width"`
	Height       int32   `db:"height" json:"height"`
	FileSize     int32   `db:"file_size" json:"file_size"`
	VideoStartTs float64 `db:"video_start_ts" json:"video_start_ts"`
	FilePath     string  `db:"file_path" json:"file_path"`
}
