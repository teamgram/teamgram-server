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
	VideoThumbId     int64  `db:"video_thumb_id"`
	Version          int32  `db:"version"`
	Attributes       string `db:"attributes"`
	Date2            int64  `db:"date2"`
}
