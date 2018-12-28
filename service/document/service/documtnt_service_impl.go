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

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/pkg/util"
	document2 "github.com/nebula-chat/chatengine/service/document/biz/core/document"
	photo2 "github.com/nebula-chat/chatengine/service/document/biz/core/photo"
	"github.com/nebula-chat/chatengine/service/nbfs/cachefs"
	"github.com/nebula-chat/chatengine/service/nbfs/nbfs"
	"time"
)

type DocumentServiceImpl struct {
	*document2.DocumentModel
	*photo2.PhotoModel
	nbfs_client.NbfsFacade
}

func NewDocumentServiceImpl(serverId int32, dataPath string, dbName, redisName string) *DocumentServiceImpl {
	if dataPath == "" {
		dataPath = "/opt/nbfs"
	}

	cachefs.InitCacheFS(dataPath)

	s := &DocumentServiceImpl{}
	s.NbfsFacade, _ = nbfs_client.NewNbfsFacade("local", util.Int32ToString(serverId))
	s.PhotoModel = photo2.NewPhotoModel(serverId, dbName, redisName)
	s.DocumentModel = document2.NewDocumentModel(serverId, dbName, redisName, s.PhotoModel)

	return s
}

// rpc
// rpc nbfs_uploadPhotoFile(UploadPhotoFileRequest) returns (PhotoDataRsp);
func (s *DocumentServiceImpl) NbfsUploadPhotoFile(ctx context.Context, request *mtproto.UploadPhotoFileRequest) (*mtproto.PhotoDataRsp, error) {
	glog.Infof("nbfs.uploadPhotoFile - request: %s", logger.JsonDebugData(request))

	var (
		reply     *mtproto.PhotoDataRsp
		inputFile = request.GetFile()
	)

	if request.GetFile() == nil {
		return nil, fmt.Errorf("bad request")
	}

	fileMDList, err := s.NbfsFacade.UploadPhotoFile(request.OwnerId, inputFile)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// glog.Info(fileMDList)
	photoId, accessHash, szList, err := s.PhotoModel.UploadPhotoFile2(fileMDList)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	reply = &mtproto.PhotoDataRsp{
		PhotoId:    photoId,
		AccessHash: accessHash,
		Date:       int32(time.Now().Unix()),
		SizeList:   szList,
	}

	// glog.Infof("nbfs.uploadPhotoFile - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

// rpc nbfs_getPhotoFileData(GetPhotoFileDataRequest) returns (PhotoDataRsp);
func (s *DocumentServiceImpl) NbfsGetPhotoFileData(ctx context.Context, request *mtproto.GetPhotoFileDataRequest) (*mtproto.PhotoDataRsp, error) {
	glog.Infof("nbfs.getPhotoFileData - request: %s", logger.JsonDebugData(request))

	var photoId = request.GetPhotoId()
	szList := s.PhotoModel.GetPhotoSizeList(photoId)
	reply := &mtproto.PhotoDataRsp{
		PhotoId:  photoId,
		SizeList: szList,
	}

	// glog.Infof("nbfs.getPhotoFileData - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

// inputMediaUploadedPhoto#2f37e231 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
func (s *DocumentServiceImpl) NbfsUploadedPhotoMedia(ctx context.Context, request *mtproto.NbfsUploadedPhotoMedia) (*mtproto.TLMessageMediaPhoto, error) {
	glog.Infof("nbfs.uploadedPhotoMedia - request: %s", logger.JsonDebugData(request))

	var (
		inputFile = request.GetMedia().GetFile()
	)

	fileMDList, err := s.NbfsFacade.UploadPhotoFile(request.OwnerId, inputFile)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	photoId, accessHash, szList, err := s.PhotoModel.UploadPhotoFile2(fileMDList)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	photo := &mtproto.TLPhotoLayer86{Data2: &mtproto.Photo_Data{
		Id:          photoId,
		HasStickers: false,
		AccessHash:  accessHash,
		Date:        int32(time.Now().Unix()),
		Sizes:       szList,
	}}

	// photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int
	var reply = &mtproto.TLMessageMediaPhoto{Data2: &mtproto.MessageMedia_Data{
		Photo_1: photo.To_Photo(),
		// Caption:    request.GetMedia().GetCaption(),
		TtlSeconds: request.GetMedia().GetTtlSeconds(),
	}}

	// glog.Info("nbfs.uploadedPhotoMedia - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

// inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
func (s *DocumentServiceImpl) NbfsUploadedDocumentMedia(ctx context.Context, request *mtproto.NbfsUploadedDocumentMedia) (*mtproto.TLMessageMediaDocument, error) {
	glog.Infof("nbfs.uploadedDocumentMedia - request: %s", logger.JsonDebugData(request))

	var (
		inputFile  = request.GetMedia().GetFile()
		inputThumb = request.GetMedia().GetThumb()
		media      = request.GetMedia()
		thumb      *mtproto.PhotoSize
		thumbId    int64
	)

	if media.GetThumb() != nil {
		thumbFileMDList, err := s.NbfsFacade.UploadPhotoFile(request.OwnerId, inputThumb)
		if err != nil {
			glog.Error(err)
			return nil, err
		}

		var szList []*mtproto.PhotoSize
		thumbId, _, szList, err = s.PhotoModel.UploadPhotoFile2(thumbFileMDList)
		if err != nil {
			glog.Error(err)
			return nil, err
		}

		thumb = &mtproto.PhotoSize{
			Constructor: mtproto.TLConstructor_CRC32_photoSize,
			Data2:       szList[0].GetData2(),
		}
		if thumb.Data2.Size == 0 {
			thumb.Data2.Size = int32(len(thumb.Data2.Bytes))
		}
	} else {
		thumb = &mtproto.PhotoSize{
			Constructor: mtproto.TLConstructor_CRC32_photoSizeEmpty,
			Data2: &mtproto.PhotoSize_Data{
				Type: "s",
			},
		}
	}

	fileMD, err := s.NbfsFacade.UploadDocumentFile(request.OwnerId, inputFile)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	fileMD.MimeType = request.GetMedia().GetMimeType()
	mediaAttributes, _ := json.Marshal(media.GetAttributes())
	data, err := s.DocumentModel.DoUploadedDocumentFile2(fileMD, thumbId, mediaAttributes)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	document := &mtproto.TLDocumentLayer86{Data2: &mtproto.Document_Data{
		Id:         data.DocumentId,
		AccessHash: data.AccessHash,
		Date:       int32(time.Now().Unix()),
		MimeType:   data.MimeType,
		Size:       int32(data.FileSize),
		Thumb:      thumb,
		DcId:       fileMD.DcId,
		// Version:    0,
		Attributes: media.GetAttributes(),
	}}

	// messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
	var reply = &mtproto.TLMessageMediaDocument{Data2: &mtproto.MessageMedia_Data{
		Document: document.To_Document(),
		// Caption:    request.GetMedia().GetCaption(),
		TtlSeconds: request.GetMedia().GetTtlSeconds(),
	}}

	// glog.Infof("nbfs.uploadedDocumentMedia - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

// rpc nbfs_getDocument(DocumentId) returns (PhotoDataRsp);
func (s *DocumentServiceImpl) NbfsGetDocument(ctx context.Context, request *mtproto.DocumentId) (*mtproto.Document, error) {
	glog.Infof("nbfs_getDocument - request: %s", logger.JsonDebugData(request))

	reply := s.DocumentModel.GetDocument(request.Id, request.AccessHash, request.Version)

	// glog.Infof("nbfs_getDocument - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

// rpc nbfs_getDocumentList(DocumentIdList) returns (DocumentList);
func (s *DocumentServiceImpl) NbfsGetDocumentList(ctx context.Context, request *mtproto.DocumentIdList) (*mtproto.DocumentList, error) {
	glog.Infof("nbfs_getDocumentList - request: %s", logger.JsonDebugData(request))

	documents := s.DocumentModel.GetDocumentList(request.IdList)

	// glog.Infof("nbfs_getDocumentList - reply: %s", logger.JsonDebugData(documents))
	return &mtproto.DocumentList{Documents: documents}, nil
}
