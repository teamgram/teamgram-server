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

package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type documentData struct {
	*dataobject.DocumentsDO
}

/*
document#1e87342b flags:#
	id:long
	access_hash:long
	file_reference:bytes
	date:int
	mime_type:string
	size:int
	thumbs:flags.0?Vector<PhotoSize>
	video_thumbs:flags.1?Vector<VideoSize>
	dc_id:int
	attributes:Vector<DocumentAttribute> = Document;
*/
func (m *Dao) makeDocumentByDO(ctx context.Context, id int64, do *dataobject.DocumentsDO) (document *mtproto.Document) {
	var (
		thumbs      []*mtproto.PhotoSize
		videoThumbs []*mtproto.VideoSize
	)

	if do == nil {
		document = mtproto.MakeTLDocumentEmpty(&mtproto.Document{
			Id: id,
		}).To_Document()
		return
	}

	// thumbs
	if do.ThumbId != 0 {
		thumbs = m.GetPhotoSizeListV2(ctx, do.ThumbId)
		logx.WithContext(ctx).Infof("sizeList = %#v", thumbs)
	}

	// video_thumbs
	if do.VideoThumbId > 0 {
		videoThumbs = m.GetVideoSizeList(ctx, do.VideoThumbId)
	}

	var attributes []*mtproto.DocumentAttribute
	err := json.Unmarshal([]byte(do.Attributes), &attributes)
	if err != nil {
		logx.WithContext(ctx).Error(err.Error())
		attributes = []*mtproto.DocumentAttribute{}
	}

	if do.Date2 == 0 {
		do.Date2 = time.Now().Unix()
	}
	// if do.Attributes
	document = mtproto.MakeTLDocument(&mtproto.Document{
		Id:            do.DocumentId,
		AccessHash:    do.AccessHash,
		FileReference: []byte{},
		Date:          int32(do.Date2),
		MimeType:      do.MimeType,
		Size2:         do.FileSize,
		Thumbs:        thumbs,
		VideoThumbs:   videoThumbs,
		DcId:          1,
		Attributes:    attributes,
	}).To_Document()

	return
}

// ???
func (m *Dao) GetDocument(ctx context.Context, id, accessHash int64, version int32) *mtproto.Document {
	do, _ := m.DocumentsDAO.SelectByFileLocation(ctx, id, accessHash, version)
	if do == nil {
		logx.WithContext(ctx).Infof("getDocument")
	}
	return m.makeDocumentByDO(ctx, id, do)
}

func (m *Dao) GetDocumentById(ctx context.Context, id int64) *mtproto.Document {
	do, _ := m.DocumentsDAO.SelectById(ctx, id)
	if do == nil {
		logx.WithContext(ctx).Infof("not found document by id: %d", id)
	}
	return m.makeDocumentByDO(ctx, id, do)
}

func (m *Dao) GetDocumentList(ctx context.Context, idList []int64) []*mtproto.Document {
	documentList := make([]*mtproto.Document, 0, len(idList))
	if len(idList) == 0 {
		return documentList
	}

	doList, _ := m.DocumentsDAO.SelectByIdList(ctx, idList)
	for i := 0; i < len(doList); i++ {
		documentList = append(documentList, m.makeDocumentByDO(ctx, 0, &doList[i]))
	}

	return documentList
}

func (m *Dao) SaveDocumentV2(ctx context.Context, fileName string, document *mtproto.Document) {
	var (
		aStr string
	)

	if document.GetAttributes() != nil {
		aBuf, _ := jsonx.Marshal(document.GetAttributes())
		aStr = hack.String(aBuf)
	}

	data := &dataobject.DocumentsDO{
		DocumentId:       document.Id,
		AccessHash:       document.AccessHash,
		DcId:             document.DcId,
		FilePath:         fmt.Sprintf("%d.dat", document.Id),
		FileSize:         document.Size2,
		UploadedFileName: fileName,
		Ext:              getFileExtName(fileName),
		MimeType:         document.MimeType,
		ThumbId:          0,
		VideoThumbId:     0,
		Version:          0,
		Attributes:       aStr,
		Date2:            int64(document.Date),
	}
	if len(document.GetThumbs()) > 0 {
		data.ThumbId = document.Id
	}
	if len(document.GetVideoThumbs()) > 0 {
		data.VideoThumbId = document.Id
	}

	data.Id, _, _ = m.DocumentsDAO.Insert(ctx, data)
	return
}
