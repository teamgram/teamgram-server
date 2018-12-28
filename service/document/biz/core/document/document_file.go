// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package document

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/document/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/service/nbfs/proto"
	"time"
)

type documentData struct {
	*dataobject.DocumentsDO
}

func (m *DocumentModel) DoUploadedDocumentFile2(fileMD *nbfs.DocumentFileMetadata, thumbId int64, attributes []byte) (*documentData, error) {
	data := &dataobject.DocumentsDO{
		DocumentId:       fileMD.DocumentId,
		AccessHash:       fileMD.AccessHash,
		DcId:             fileMD.DcId,
		FilePath:         fileMD.FilePath,
		FileSize:         fileMD.FileSize,
		UploadedFileName: fileMD.UploadedFileName,
		Ext:              fileMD.Ext,
		MimeType:         fileMD.MimeType,
		ThumbId:          thumbId,
		Attributes:       string(attributes),
		Version:          0,
	}
	data.Id = m.dao.DocumentsDAO.Insert(data)
	return &documentData{DocumentsDO: data}, nil
}

func (m *DocumentModel) makeDocumentByDO(do *dataobject.DocumentsDO) *mtproto.Document {
	var (
		thumb    *mtproto.PhotoSize
		document *mtproto.Document
	)

	if do == nil {
		document = mtproto.NewTLDocumentEmpty().To_Document()
	} else {
		if do.ThumbId != 0 {
			sizeList := m.cb.GetPhotoSizeList(do.ThumbId)
			if len(sizeList) > 0 {
				thumb = sizeList[0]
			}
		}
		if thumb == nil {
			thumb = mtproto.NewTLPhotoSizeEmpty().To_PhotoSize()
		}

		var attributes []*mtproto.DocumentAttribute
		err := json.Unmarshal([]byte(do.Attributes), &attributes)
		if err != nil {
			glog.Error(err)
			attributes = []*mtproto.DocumentAttribute{}
		}

		// if do.Attributes
		document = &mtproto.Document{
			Constructor: mtproto.TLConstructor_CRC32_documentLayer86,
			Data2: &mtproto.Document_Data{
				Id:         do.DocumentId,
				AccessHash: do.AccessHash,
				Date:       int32(time.Now().Unix()),
				MimeType:   do.MimeType,
				Size:       do.FileSize,
				Thumb:      thumb,
				DcId:       2,
				// Version:    do.Version,
				Attributes: attributes,
			},
		}
	}

	return document
}

func (m *DocumentModel) GetDocument(id, accessHash int64, version int32) *mtproto.Document {
	do := m.dao.DocumentsDAO.SelectByFileLocation(id, accessHash, version)
	if do == nil {
		glog.Warning("")
	}
	return m.makeDocumentByDO(do)
}

func (m *DocumentModel) GetDocumentList(idList []int64) []*mtproto.Document {
	doList := m.dao.DocumentsDAO.SelectByIdList(idList)
	documetList := make([]*mtproto.Document, len(doList))
	for i := 0; i < len(doList); i++ {
		documetList[i] = m.makeDocumentByDO(&doList[i])
	}
	return documetList
}
