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
// photos.updateProfilePhoto#1c3d5956 flags:# fallback:flags.0?true id:InputPhoto = photos.Photo;
func (s *Service) PhotosUpdateProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("photos.updateProfilePhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosUpdateProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("photos.updateProfilePhoto - reply: %s", r.DebugString())
	return r, err
}

// PhotosUploadProfilePhoto
// photos.uploadProfilePhoto#93c9a51 flags:# fallback:flags.3?true file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = photos.Photo;
func (s *Service) PhotosUploadProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("photos.uploadProfilePhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosUploadProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("photos.uploadProfilePhoto - reply: %s", r.DebugString())
	return r, err
}

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (s *Service) PhotosDeletePhotos(ctx context.Context, request *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("photos.deletePhotos - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosDeletePhotos(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("photos.deletePhotos - reply: %s", r.DebugString())
	return r, err
}

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (s *Service) PhotosGetUserPhotos(ctx context.Context, request *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("photos.getUserPhotos - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosGetUserPhotos(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("photos.getUserPhotos - reply: %s", r.DebugString())
	return r, err
}

// PhotosUploadContactProfilePhoto
// photos.uploadContactProfilePhoto#e14c4a71 flags:# suggest:flags.3?true save:flags.4?true user_id:InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.5?VideoSize = photos.Photo;
func (s *Service) PhotosUploadContactProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUploadContactProfilePhoto) (*mtproto.Photos_Photo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("photos.uploadContactProfilePhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PhotosUploadContactProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("photos.uploadContactProfilePhoto - reply: %s", r.DebugString())
	return r, err
}
