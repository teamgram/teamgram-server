/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userprofileclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/userprofile/userprofile/userprofileservice"

	"github.com/cloudwego/kitex/client"
)

type UserProfileClient interface {
	AccountUpdateProfile(ctx context.Context, in *tg.TLAccountUpdateProfile) (*tg.User, error)
	AccountUpdateStatus(ctx context.Context, in *tg.TLAccountUpdateStatus) (*tg.Bool, error)
	AccountUpdateBirthday(ctx context.Context, in *tg.TLAccountUpdateBirthday) (*tg.Bool, error)
	AccountUpdatePersonalChannel(ctx context.Context, in *tg.TLAccountUpdatePersonalChannel) (*tg.Bool, error)
	ContactsGetBirthdays(ctx context.Context, in *tg.TLContactsGetBirthdays) (*tg.ContactsContactBirthdays, error)
	PhotosUpdateProfilePhoto(ctx context.Context, in *tg.TLPhotosUpdateProfilePhoto) (*tg.PhotosPhoto, error)
	PhotosUploadProfilePhoto(ctx context.Context, in *tg.TLPhotosUploadProfilePhoto) (*tg.PhotosPhoto, error)
	PhotosDeletePhotos(ctx context.Context, in *tg.TLPhotosDeletePhotos) (*tg.VectorLong, error)
	PhotosGetUserPhotos(ctx context.Context, in *tg.TLPhotosGetUserPhotos) (*tg.PhotosPhotos, error)
	PhotosUploadContactProfilePhoto(ctx context.Context, in *tg.TLPhotosUploadContactProfilePhoto) (*tg.PhotosPhoto, error)
	AccountUpdateVerified(ctx context.Context, in *tg.TLAccountUpdateVerified) (*tg.User, error)
}

type defaultUserProfileClient struct {
	cli client.Client
}

func NewUserProfileClient(cli client.Client) UserProfileClient {
	return &defaultUserProfileClient{
		cli: cli,
	}
}

// AccountUpdateProfile
// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (m *defaultUserProfileClient) AccountUpdateProfile(ctx context.Context, in *tg.TLAccountUpdateProfile) (*tg.User, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.AccountUpdateProfile(ctx, in)
}

// AccountUpdateStatus
// account.updateStatus#6628562c offline:Bool = Bool;
func (m *defaultUserProfileClient) AccountUpdateStatus(ctx context.Context, in *tg.TLAccountUpdateStatus) (*tg.Bool, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.AccountUpdateStatus(ctx, in)
}

// AccountUpdateBirthday
// account.updateBirthday#cc6e0c11 flags:# birthday:flags.0?Birthday = Bool;
func (m *defaultUserProfileClient) AccountUpdateBirthday(ctx context.Context, in *tg.TLAccountUpdateBirthday) (*tg.Bool, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.AccountUpdateBirthday(ctx, in)
}

// AccountUpdatePersonalChannel
// account.updatePersonalChannel#d94305e0 channel:InputChannel = Bool;
func (m *defaultUserProfileClient) AccountUpdatePersonalChannel(ctx context.Context, in *tg.TLAccountUpdatePersonalChannel) (*tg.Bool, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.AccountUpdatePersonalChannel(ctx, in)
}

// ContactsGetBirthdays
// contacts.getBirthdays#daeda864 = contacts.ContactBirthdays;
func (m *defaultUserProfileClient) ContactsGetBirthdays(ctx context.Context, in *tg.TLContactsGetBirthdays) (*tg.ContactsContactBirthdays, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.ContactsGetBirthdays(ctx, in)
}

// PhotosUpdateProfilePhoto
// photos.updateProfilePhoto#9e82039 flags:# fallback:flags.0?true bot:flags.1?InputUser id:InputPhoto = photos.Photo;
func (m *defaultUserProfileClient) PhotosUpdateProfilePhoto(ctx context.Context, in *tg.TLPhotosUpdateProfilePhoto) (*tg.PhotosPhoto, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.PhotosUpdateProfilePhoto(ctx, in)
}

// PhotosUploadProfilePhoto
// photos.uploadProfilePhoto#388a3b5 flags:# fallback:flags.3?true bot:flags.5?InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = photos.Photo;
func (m *defaultUserProfileClient) PhotosUploadProfilePhoto(ctx context.Context, in *tg.TLPhotosUploadProfilePhoto) (*tg.PhotosPhoto, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.PhotosUploadProfilePhoto(ctx, in)
}

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (m *defaultUserProfileClient) PhotosDeletePhotos(ctx context.Context, in *tg.TLPhotosDeletePhotos) (*tg.VectorLong, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.PhotosDeletePhotos(ctx, in)
}

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (m *defaultUserProfileClient) PhotosGetUserPhotos(ctx context.Context, in *tg.TLPhotosGetUserPhotos) (*tg.PhotosPhotos, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.PhotosGetUserPhotos(ctx, in)
}

// PhotosUploadContactProfilePhoto
// photos.uploadContactProfilePhoto#e14c4a71 flags:# suggest:flags.3?true save:flags.4?true user_id:InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.5?VideoSize = photos.Photo;
func (m *defaultUserProfileClient) PhotosUploadContactProfilePhoto(ctx context.Context, in *tg.TLPhotosUploadContactProfilePhoto) (*tg.PhotosPhoto, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.PhotosUploadContactProfilePhoto(ctx, in)
}

// AccountUpdateVerified
// account.updateVerified flags:# id:long verified:flags.0?true = User;
func (m *defaultUserProfileClient) AccountUpdateVerified(ctx context.Context, in *tg.TLAccountUpdateVerified) (*tg.User, error) {
	cli := userprofileservice.NewRPCUserProfileClient(m.cli)
	return cli.AccountUpdateVerified(ctx, in)
}
