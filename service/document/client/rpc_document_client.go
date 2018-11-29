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

package document_client

import (
	"context"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/mtproto"
)

type documentClient struct {
	client mtproto.RPCNbfsClient
}

var (
	nbfsInstance = &documentClient{}
)

func InstallNbfsClient(discovery *service_discovery.ServiceDiscoveryClientConfig) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	nbfsInstance.client = mtproto.NewRPCNbfsClient(conn)
}

func UploadPhotoFile(ownerId int64, file *mtproto.InputFile) (*mtproto.PhotoDataRsp, error) {
	// TODO(@benqi): Check nbfsInstance.client inited

	request := &mtproto.UploadPhotoFileRequest{
		OwnerId: ownerId,
		File:    file,
	}
	reply, err := nbfsInstance.client.NbfsUploadPhotoFile(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetPhotoSizeList(photoId int64) ([]*mtproto.PhotoSize, error) {
	// TODO(@benqi): Check nbfsInstance.client inited

	request := &mtproto.GetPhotoFileDataRequest{
		PhotoId: photoId,
	}
	reply, err := nbfsInstance.client.NbfsGetPhotoFileData(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply.SizeList, nil
}

func UploadedPhotoMedia(ownerId int64, media *mtproto.TLInputMediaUploadedPhoto) (*mtproto.TLMessageMediaPhoto, error) {
	// TODO(@benqi): Check nbfsInstance.client inited

	request := &mtproto.NbfsUploadedPhotoMedia{
		OwnerId: ownerId,
		Media:   media,
	}

	reply, err := nbfsInstance.client.NbfsUploadedPhotoMedia(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func UploadedDocumentMedia(ownerId int64, media *mtproto.TLInputMediaUploadedDocument) (*mtproto.TLMessageMediaDocument, error) {
	// TODO(@benqi): Check nbfsInstance.client inited

	request := &mtproto.NbfsUploadedDocumentMedia{
		OwnerId: ownerId,
		Media:   media,
	}

	reply, err := nbfsInstance.client.NbfsUploadedDocumentMedia(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetDocumentById(id, accessHash int64) (*mtproto.Document, error) {
	// TODO(@benqi): Check nbfsInstance.client inited

	request := &mtproto.DocumentId{
		Id:         id,
		AccessHash: accessHash,
		Version:    0,
	}
	reply, err := nbfsInstance.client.NbfsGetDocument(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetDocumentByIdList(idList []int64) ([]*mtproto.Document, error) {
	// TODO(@benqi): Check nbfsInstance.client inited
	reply, err := nbfsInstance.client.NbfsGetDocumentList(context.Background(), &mtproto.DocumentIdList{IdList: idList})
	if err != nil {
		return nil, err
	}

	return reply.Documents, nil
}
