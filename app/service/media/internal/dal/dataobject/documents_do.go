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

type DocumentsDO struct {
	Id               int64  `db:"id" json:"id"`
	DocumentId       int64  `db:"document_id" json:"document_id"`
	AccessHash       int64  `db:"access_hash" json:"access_hash"`
	DcId             int32  `db:"dc_id" json:"dc_id"`
	FilePath         string `db:"file_path" json:"file_path"`
	FileSize         int64  `db:"file_size" json:"file_size"`
	UploadedFileName string `db:"uploaded_file_name" json:"uploaded_file_name"`
	Ext              string `db:"ext" json:"ext"`
	MimeType         string `db:"mime_type" json:"mime_type"`
	ThumbId          int64  `db:"thumb_id" json:"thumb_id"`
	VideoThumbId     int64  `db:"video_thumb_id" json:"video_thumb_id"`
	Version          int32  `db:"version" json:"version"`
	Attributes       string `db:"attributes" json:"attributes"`
	Date2            int64  `db:"date2" json:"date2"`
	ImportDocumentId int64  `db:"import_document_id" json:"import_document_id"`
	Deleted          bool   `db:"deleted" json:"deleted"`
}
