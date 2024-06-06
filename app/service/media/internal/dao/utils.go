// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"path"
	"strings"

	"github.com/teamgram/proto/mtproto"
)

var (
	emptyPhoto = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
)

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	if ext == "" {
		ext = ".partial"
	}
	return strings.ToLower(ext)
}

func makePhotoEmpty(id int64) *mtproto.Photo {
	return mtproto.MakeTLPhotoEmpty(&mtproto.Photo{
		Id: id,
	}).To_Photo()
}
