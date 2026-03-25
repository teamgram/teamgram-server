/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgooo Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package userchannelprofilesservice

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountUpdateProfile(ctx context.Context, req *tg.TLAccountUpdateProfile, callOptions ...callopt.Option) (r *tg.User, err error)
	AccountUpdateStatus(ctx context.Context, req *tg.TLAccountUpdateStatus, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountUpdateBirthday(ctx context.Context, req *tg.TLAccountUpdateBirthday, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountUpdatePersonalChannel(ctx context.Context, req *tg.TLAccountUpdatePersonalChannel, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountSetMainProfileTab(ctx context.Context, req *tg.TLAccountSetMainProfileTab, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountSaveMusic(ctx context.Context, req *tg.TLAccountSaveMusic, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountGetSavedMusicIds(ctx context.Context, req *tg.TLAccountGetSavedMusicIds, callOptions ...callopt.Option) (r *tg.AccountSavedMusicIds, err error)
	UsersGetSavedMusic(ctx context.Context, req *tg.TLUsersGetSavedMusic, callOptions ...callopt.Option) (r *tg.UsersSavedMusic, err error)
	UsersGetSavedMusicByID(ctx context.Context, req *tg.TLUsersGetSavedMusicByID, callOptions ...callopt.Option) (r *tg.UsersSavedMusic, err error)
	UsersSuggestBirthday(ctx context.Context, req *tg.TLUsersSuggestBirthday, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ContactsGetBirthdays(ctx context.Context, req *tg.TLContactsGetBirthdays, callOptions ...callopt.Option) (r *tg.ContactsContactBirthdays, err error)
	PhotosUpdateProfilePhoto(ctx context.Context, req *tg.TLPhotosUpdateProfilePhoto, callOptions ...callopt.Option) (r *tg.PhotosPhoto, err error)
	PhotosUploadProfilePhoto(ctx context.Context, req *tg.TLPhotosUploadProfilePhoto, callOptions ...callopt.Option) (r *tg.PhotosPhoto, err error)
	PhotosDeletePhotos(ctx context.Context, req *tg.TLPhotosDeletePhotos, callOptions ...callopt.Option) (r *tg.VectorLong, err error)
	PhotosGetUserPhotos(ctx context.Context, req *tg.TLPhotosGetUserPhotos, callOptions ...callopt.Option) (r *tg.PhotosPhotos, err error)
	PhotosUploadContactProfilePhoto(ctx context.Context, req *tg.TLPhotosUploadContactProfilePhoto, callOptions ...callopt.Option) (r *tg.PhotosPhoto, err error)
	ChannelsSetMainProfileTab(ctx context.Context, req *tg.TLChannelsSetMainProfileTab, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountUpdateVerified(ctx context.Context, req *tg.TLAccountUpdateVerified, callOptions ...callopt.Option) (r *tg.User, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserChannelProfilesClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserChannelProfilesClient struct {
	*kClient
}

func NewRPCUserChannelProfilesClient(cli client.Client) Client {
	return &kUserChannelProfilesClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kUserChannelProfilesClient) AccountUpdateProfile(ctx context.Context, req *tg.TLAccountUpdateProfile, callOptions ...callopt.Option) (r *tg.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUpdateProfile(ctx, req)
}

func (p *kUserChannelProfilesClient) AccountUpdateStatus(ctx context.Context, req *tg.TLAccountUpdateStatus, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUpdateStatus(ctx, req)
}

func (p *kUserChannelProfilesClient) AccountUpdateBirthday(ctx context.Context, req *tg.TLAccountUpdateBirthday, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUpdateBirthday(ctx, req)
}

func (p *kUserChannelProfilesClient) AccountUpdatePersonalChannel(ctx context.Context, req *tg.TLAccountUpdatePersonalChannel, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUpdatePersonalChannel(ctx, req)
}

func (p *kUserChannelProfilesClient) AccountSetMainProfileTab(ctx context.Context, req *tg.TLAccountSetMainProfileTab, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSetMainProfileTab(ctx, req)
}

func (p *kUserChannelProfilesClient) AccountSaveMusic(ctx context.Context, req *tg.TLAccountSaveMusic, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSaveMusic(ctx, req)
}

func (p *kUserChannelProfilesClient) AccountGetSavedMusicIds(ctx context.Context, req *tg.TLAccountGetSavedMusicIds, callOptions ...callopt.Option) (r *tg.AccountSavedMusicIds, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetSavedMusicIds(ctx, req)
}

func (p *kUserChannelProfilesClient) UsersGetSavedMusic(ctx context.Context, req *tg.TLUsersGetSavedMusic, callOptions ...callopt.Option) (r *tg.UsersSavedMusic, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsersGetSavedMusic(ctx, req)
}

func (p *kUserChannelProfilesClient) UsersGetSavedMusicByID(ctx context.Context, req *tg.TLUsersGetSavedMusicByID, callOptions ...callopt.Option) (r *tg.UsersSavedMusic, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsersGetSavedMusicByID(ctx, req)
}

func (p *kUserChannelProfilesClient) UsersSuggestBirthday(ctx context.Context, req *tg.TLUsersSuggestBirthday, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsersSuggestBirthday(ctx, req)
}

func (p *kUserChannelProfilesClient) ContactsGetBirthdays(ctx context.Context, req *tg.TLContactsGetBirthdays, callOptions ...callopt.Option) (r *tg.ContactsContactBirthdays, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsGetBirthdays(ctx, req)
}

func (p *kUserChannelProfilesClient) PhotosUpdateProfilePhoto(ctx context.Context, req *tg.TLPhotosUpdateProfilePhoto, callOptions ...callopt.Option) (r *tg.PhotosPhoto, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PhotosUpdateProfilePhoto(ctx, req)
}

func (p *kUserChannelProfilesClient) PhotosUploadProfilePhoto(ctx context.Context, req *tg.TLPhotosUploadProfilePhoto, callOptions ...callopt.Option) (r *tg.PhotosPhoto, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PhotosUploadProfilePhoto(ctx, req)
}

func (p *kUserChannelProfilesClient) PhotosDeletePhotos(ctx context.Context, req *tg.TLPhotosDeletePhotos, callOptions ...callopt.Option) (r *tg.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PhotosDeletePhotos(ctx, req)
}

func (p *kUserChannelProfilesClient) PhotosGetUserPhotos(ctx context.Context, req *tg.TLPhotosGetUserPhotos, callOptions ...callopt.Option) (r *tg.PhotosPhotos, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PhotosGetUserPhotos(ctx, req)
}

func (p *kUserChannelProfilesClient) PhotosUploadContactProfilePhoto(ctx context.Context, req *tg.TLPhotosUploadContactProfilePhoto, callOptions ...callopt.Option) (r *tg.PhotosPhoto, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PhotosUploadContactProfilePhoto(ctx, req)
}

func (p *kUserChannelProfilesClient) ChannelsSetMainProfileTab(ctx context.Context, req *tg.TLChannelsSetMainProfileTab, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsSetMainProfileTab(ctx, req)
}

func (p *kUserChannelProfilesClient) AccountUpdateVerified(ctx context.Context, req *tg.TLAccountUpdateVerified, callOptions ...callopt.Option) (r *tg.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUpdateVerified(ctx, req)
}
