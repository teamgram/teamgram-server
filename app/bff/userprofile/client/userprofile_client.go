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

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UserProfileClient interface {
	AccountUpdateProfile(ctx context.Context, in *mtproto.TLAccountUpdateProfile) (*mtproto.User, error)
	AccountUpdateStatus(ctx context.Context, in *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error)
	AccountUpdateBirthday(ctx context.Context, in *mtproto.TLAccountUpdateBirthday) (*mtproto.Bool, error)
	AccountUpdatePersonalChannel(ctx context.Context, in *mtproto.TLAccountUpdatePersonalChannel) (*mtproto.Bool, error)
	ContactsGetBirthdays(ctx context.Context, in *mtproto.TLContactsGetBirthdays) (*mtproto.Contacts_ContactBirthdays, error)
	PhotosUpdateProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error)
	PhotosUploadProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error)
	PhotosDeletePhotos(ctx context.Context, in *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error)
	PhotosGetUserPhotos(ctx context.Context, in *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error)
	PhotosUploadContactProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadContactProfilePhoto) (*mtproto.Photos_Photo, error)
	AccountUpdateVerified(ctx context.Context, in *mtproto.TLAccountUpdateVerified) (*mtproto.User, error)
}

type defaultUserProfileClient struct {
	cli zrpc.Client
}

func NewUserProfileClient(cli zrpc.Client) UserProfileClient {
	return &defaultUserProfileClient{
		cli: cli,
	}
}

// AccountUpdateProfile
// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (m *defaultUserProfileClient) AccountUpdateProfile(ctx context.Context, in *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.AccountUpdateProfile(ctx, in)
}

// AccountUpdateStatus
// account.updateStatus#6628562c offline:Bool = Bool;
func (m *defaultUserProfileClient) AccountUpdateStatus(ctx context.Context, in *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.AccountUpdateStatus(ctx, in)
}

// AccountUpdateBirthday
// account.updateBirthday#cc6e0c11 flags:# birthday:flags.0?Birthday = Bool;
func (m *defaultUserProfileClient) AccountUpdateBirthday(ctx context.Context, in *mtproto.TLAccountUpdateBirthday) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.AccountUpdateBirthday(ctx, in)
}

// AccountUpdatePersonalChannel
// account.updatePersonalChannel#d94305e0 channel:InputChannel = Bool;
func (m *defaultUserProfileClient) AccountUpdatePersonalChannel(ctx context.Context, in *mtproto.TLAccountUpdatePersonalChannel) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.AccountUpdatePersonalChannel(ctx, in)
}

// ContactsGetBirthdays
// contacts.getBirthdays#daeda864 = contacts.ContactBirthdays;
func (m *defaultUserProfileClient) ContactsGetBirthdays(ctx context.Context, in *mtproto.TLContactsGetBirthdays) (*mtproto.Contacts_ContactBirthdays, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.ContactsGetBirthdays(ctx, in)
}

// PhotosUpdateProfilePhoto
// photos.updateProfilePhoto#9e82039 flags:# fallback:flags.0?true bot:flags.1?InputUser id:InputPhoto = photos.Photo;
func (m *defaultUserProfileClient) PhotosUpdateProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.PhotosUpdateProfilePhoto(ctx, in)
}

// PhotosUploadProfilePhoto
// photos.uploadProfilePhoto#388a3b5 flags:# fallback:flags.3?true bot:flags.5?InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = photos.Photo;
func (m *defaultUserProfileClient) PhotosUploadProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.PhotosUploadProfilePhoto(ctx, in)
}

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (m *defaultUserProfileClient) PhotosDeletePhotos(ctx context.Context, in *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.PhotosDeletePhotos(ctx, in)
}

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (m *defaultUserProfileClient) PhotosGetUserPhotos(ctx context.Context, in *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.PhotosGetUserPhotos(ctx, in)
}

// PhotosUploadContactProfilePhoto
// photos.uploadContactProfilePhoto#e14c4a71 flags:# suggest:flags.3?true save:flags.4?true user_id:InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.5?VideoSize = photos.Photo;
func (m *defaultUserProfileClient) PhotosUploadContactProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadContactProfilePhoto) (*mtproto.Photos_Photo, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.PhotosUploadContactProfilePhoto(ctx, in)
}

// AccountUpdateVerified
// account.updateVerified flags:# id:long verified:flags.0?true = User;
func (m *defaultUserProfileClient) AccountUpdateVerified(ctx context.Context, in *mtproto.TLAccountUpdateVerified) (*mtproto.User, error) {
	client := mtproto.NewRPCUserProfileClient(m.cli.Conn())
	return client.AccountUpdateVerified(ctx, in)
}
