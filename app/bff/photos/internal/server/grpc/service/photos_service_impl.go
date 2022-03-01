/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/photos/internal/core"
)

// PhotosUpdateProfilePhoto
// photos.updateProfilePhoto#72d4742c id:InputPhoto = photos.Photo;
func (s *Service) PhotosUpdateProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("photos.updateProfilePhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosUpdateProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Infof("photos.updateProfilePhoto - reply: %s", r.DebugString())
	return r, err
}

// PhotosUploadProfilePhoto
// photos.uploadProfilePhoto#89f30f69 flags:# file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = photos.Photo;
func (s *Service) PhotosUploadProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("photos.uploadProfilePhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosUploadProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Infof("photos.uploadProfilePhoto - reply: %s", r.DebugString())
	return r, err
}

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (s *Service) PhotosDeletePhotos(ctx context.Context, request *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("photos.deletePhotos - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosDeletePhotos(request)
	if err != nil {
		return nil, err
	}

	c.Infof("photos.deletePhotos - reply: %s", r.DebugString())
	return r, err
}

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (s *Service) PhotosGetUserPhotos(ctx context.Context, request *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("photos.getUserPhotos - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosGetUserPhotos(request)
	if err != nil {
		return nil, err
	}

	c.Infof("photos.getUserPhotos - reply: %s", r.DebugString())
	return r, err
}
