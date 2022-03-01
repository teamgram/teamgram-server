/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package photos_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type PhotosClient interface {
	PhotosUpdateProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error)
	PhotosUploadProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error)
	PhotosDeletePhotos(ctx context.Context, in *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error)
	PhotosGetUserPhotos(ctx context.Context, in *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error)
}

type defaultPhotosClient struct {
	cli zrpc.Client
}

func NewPhotosClient(cli zrpc.Client) PhotosClient {
	return &defaultPhotosClient{
		cli: cli,
	}
}

// PhotosUpdateProfilePhoto
// photos.updateProfilePhoto#72d4742c id:InputPhoto = photos.Photo;
func (m *defaultPhotosClient) PhotosUpdateProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error) {
	client := mtproto.NewRPCPhotosClient(m.cli.Conn())
	return client.PhotosUpdateProfilePhoto(ctx, in)
}

// PhotosUploadProfilePhoto
// photos.uploadProfilePhoto#89f30f69 flags:# file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = photos.Photo;
func (m *defaultPhotosClient) PhotosUploadProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	client := mtproto.NewRPCPhotosClient(m.cli.Conn())
	return client.PhotosUploadProfilePhoto(ctx, in)
}

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (m *defaultPhotosClient) PhotosDeletePhotos(ctx context.Context, in *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error) {
	client := mtproto.NewRPCPhotosClient(m.cli.Conn())
	return client.PhotosDeletePhotos(ctx, in)
}

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (m *defaultPhotosClient) PhotosGetUserPhotos(ctx context.Context, in *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	client := mtproto.NewRPCPhotosClient(m.cli.Conn())
	return client.PhotosGetUserPhotos(ctx, in)
}
