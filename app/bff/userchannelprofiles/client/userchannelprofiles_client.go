/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2025 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userchannelprofilesclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UserChannelProfilesClient interface {
	AccountUpdateProfile(ctx context.Context, in *mtproto.TLAccountUpdateProfile) (*mtproto.User, error)
	AccountUpdateStatus(ctx context.Context, in *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error)
	AccountUpdateBirthday(ctx context.Context, in *mtproto.TLAccountUpdateBirthday) (*mtproto.Bool, error)
	AccountUpdatePersonalChannel(ctx context.Context, in *mtproto.TLAccountUpdatePersonalChannel) (*mtproto.Bool, error)
	AccountSetMainProfileTab(ctx context.Context, in *mtproto.TLAccountSetMainProfileTab) (*mtproto.Bool, error)
	AccountSaveMusic(ctx context.Context, in *mtproto.TLAccountSaveMusic) (*mtproto.Bool, error)
	AccountGetSavedMusicIds(ctx context.Context, in *mtproto.TLAccountGetSavedMusicIds) (*mtproto.Account_SavedMusicIds, error)
	UsersGetSavedMusic(ctx context.Context, in *mtproto.TLUsersGetSavedMusic) (*mtproto.Users_SavedMusic, error)
	UsersGetSavedMusicByID(ctx context.Context, in *mtproto.TLUsersGetSavedMusicByID) (*mtproto.Users_SavedMusic, error)
	UsersSuggestBirthday(ctx context.Context, in *mtproto.TLUsersSuggestBirthday) (*mtproto.Updates, error)
	ContactsGetBirthdays(ctx context.Context, in *mtproto.TLContactsGetBirthdays) (*mtproto.Contacts_ContactBirthdays, error)
	PhotosUpdateProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error)
	PhotosUploadProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error)
	PhotosDeletePhotos(ctx context.Context, in *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error)
	PhotosGetUserPhotos(ctx context.Context, in *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error)
	PhotosUploadContactProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadContactProfilePhoto) (*mtproto.Photos_Photo, error)
	ChannelsSetMainProfileTab(ctx context.Context, in *mtproto.TLChannelsSetMainProfileTab) (*mtproto.Bool, error)
	AccountUpdateVerified(ctx context.Context, in *mtproto.TLAccountUpdateVerified) (*mtproto.User, error)
}

type defaultUserChannelProfilesClient struct {
	cli zrpc.Client
}

func NewUserChannelProfilesClient(cli zrpc.Client) UserChannelProfilesClient {
	return &defaultUserChannelProfilesClient{
		cli: cli,
	}
}

// AccountUpdateProfile
// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (m *defaultUserChannelProfilesClient) AccountUpdateProfile(ctx context.Context, in *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.AccountUpdateProfile(ctx, in)
}

// AccountUpdateStatus
// account.updateStatus#6628562c offline:Bool = Bool;
func (m *defaultUserChannelProfilesClient) AccountUpdateStatus(ctx context.Context, in *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.AccountUpdateStatus(ctx, in)
}

// AccountUpdateBirthday
// account.updateBirthday#cc6e0c11 flags:# birthday:flags.0?Birthday = Bool;
func (m *defaultUserChannelProfilesClient) AccountUpdateBirthday(ctx context.Context, in *mtproto.TLAccountUpdateBirthday) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.AccountUpdateBirthday(ctx, in)
}

// AccountUpdatePersonalChannel
// account.updatePersonalChannel#d94305e0 channel:InputChannel = Bool;
func (m *defaultUserChannelProfilesClient) AccountUpdatePersonalChannel(ctx context.Context, in *mtproto.TLAccountUpdatePersonalChannel) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.AccountUpdatePersonalChannel(ctx, in)
}

// AccountSetMainProfileTab
// account.setMainProfileTab#5dee78b0 tab:ProfileTab = Bool;
func (m *defaultUserChannelProfilesClient) AccountSetMainProfileTab(ctx context.Context, in *mtproto.TLAccountSetMainProfileTab) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.AccountSetMainProfileTab(ctx, in)
}

// AccountSaveMusic
// account.saveMusic#b26732a9 flags:# unsave:flags.0?true id:InputDocument after_id:flags.1?InputDocument = Bool;
func (m *defaultUserChannelProfilesClient) AccountSaveMusic(ctx context.Context, in *mtproto.TLAccountSaveMusic) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.AccountSaveMusic(ctx, in)
}

// AccountGetSavedMusicIds
// account.getSavedMusicIds#e09d5faf hash:long = account.SavedMusicIds;
func (m *defaultUserChannelProfilesClient) AccountGetSavedMusicIds(ctx context.Context, in *mtproto.TLAccountGetSavedMusicIds) (*mtproto.Account_SavedMusicIds, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.AccountGetSavedMusicIds(ctx, in)
}

// UsersGetSavedMusic
// users.getSavedMusic#788d7fe3 id:InputUser offset:int limit:int hash:long = users.SavedMusic;
func (m *defaultUserChannelProfilesClient) UsersGetSavedMusic(ctx context.Context, in *mtproto.TLUsersGetSavedMusic) (*mtproto.Users_SavedMusic, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.UsersGetSavedMusic(ctx, in)
}

// UsersGetSavedMusicByID
// users.getSavedMusicByID#7573a4e9 id:InputUser documents:Vector<InputDocument> = users.SavedMusic;
func (m *defaultUserChannelProfilesClient) UsersGetSavedMusicByID(ctx context.Context, in *mtproto.TLUsersGetSavedMusicByID) (*mtproto.Users_SavedMusic, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.UsersGetSavedMusicByID(ctx, in)
}

// UsersSuggestBirthday
// users.suggestBirthday#fc533372 id:InputUser birthday:Birthday = Updates;
func (m *defaultUserChannelProfilesClient) UsersSuggestBirthday(ctx context.Context, in *mtproto.TLUsersSuggestBirthday) (*mtproto.Updates, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.UsersSuggestBirthday(ctx, in)
}

// ContactsGetBirthdays
// contacts.getBirthdays#daeda864 = contacts.ContactBirthdays;
func (m *defaultUserChannelProfilesClient) ContactsGetBirthdays(ctx context.Context, in *mtproto.TLContactsGetBirthdays) (*mtproto.Contacts_ContactBirthdays, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.ContactsGetBirthdays(ctx, in)
}

// PhotosUpdateProfilePhoto
// photos.updateProfilePhoto#9e82039 flags:# fallback:flags.0?true bot:flags.1?InputUser id:InputPhoto = photos.Photo;
func (m *defaultUserChannelProfilesClient) PhotosUpdateProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Photos_Photo, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.PhotosUpdateProfilePhoto(ctx, in)
}

// PhotosUploadProfilePhoto
// photos.uploadProfilePhoto#388a3b5 flags:# fallback:flags.3?true bot:flags.5?InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = photos.Photo;
func (m *defaultUserChannelProfilesClient) PhotosUploadProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.PhotosUploadProfilePhoto(ctx, in)
}

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (m *defaultUserChannelProfilesClient) PhotosDeletePhotos(ctx context.Context, in *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.PhotosDeletePhotos(ctx, in)
}

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (m *defaultUserChannelProfilesClient) PhotosGetUserPhotos(ctx context.Context, in *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.PhotosGetUserPhotos(ctx, in)
}

// PhotosUploadContactProfilePhoto
// photos.uploadContactProfilePhoto#e14c4a71 flags:# suggest:flags.3?true save:flags.4?true user_id:InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.5?VideoSize = photos.Photo;
func (m *defaultUserChannelProfilesClient) PhotosUploadContactProfilePhoto(ctx context.Context, in *mtproto.TLPhotosUploadContactProfilePhoto) (*mtproto.Photos_Photo, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.PhotosUploadContactProfilePhoto(ctx, in)
}

// ChannelsSetMainProfileTab
// channels.setMainProfileTab#3583fcb1 channel:InputChannel tab:ProfileTab = Bool;
func (m *defaultUserChannelProfilesClient) ChannelsSetMainProfileTab(ctx context.Context, in *mtproto.TLChannelsSetMainProfileTab) (*mtproto.Bool, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.ChannelsSetMainProfileTab(ctx, in)
}

// AccountUpdateVerified
// account.updateVerified flags:# id:long verified:flags.0?true = User;
func (m *defaultUserChannelProfilesClient) AccountUpdateVerified(ctx context.Context, in *mtproto.TLAccountUpdateVerified) (*mtproto.User, error) {
	client := mtproto.NewRPCUserChannelProfilesClient(m.cli.Conn())
	return client.AccountUpdateVerified(ctx, in)
}
