/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialogservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DialogSaveDraftMessage(ctx context.Context, req *dialog.TLDialogSaveDraftMessage, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogClearDraftMessage(ctx context.Context, req *dialog.TLDialogClearDraftMessage, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogGetAllDrafts(ctx context.Context, req *dialog.TLDialogGetAllDrafts, callOptions ...callopt.Option) (r *dialog.VectorPeerWithDraftMessage, err error)
	DialogClearAllDrafts(ctx context.Context, req *dialog.TLDialogClearAllDrafts, callOptions ...callopt.Option) (r *dialog.VectorPeerWithDraftMessage, err error)
	DialogMarkDialogUnread(ctx context.Context, req *dialog.TLDialogMarkDialogUnread, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogToggleDialogPin(ctx context.Context, req *dialog.TLDialogToggleDialogPin, callOptions ...callopt.Option) (r *tg.Int32, err error)
	DialogGetDialogUnreadMarkList(ctx context.Context, req *dialog.TLDialogGetDialogUnreadMarkList, callOptions ...callopt.Option) (r *dialog.VectorDialogPeer, err error)
	DialogGetDialogsByOffsetDate(ctx context.Context, req *dialog.TLDialogGetDialogsByOffsetDate, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error)
	DialogGetDialogs(ctx context.Context, req *dialog.TLDialogGetDialogs, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error)
	DialogGetDialogsByIdList(ctx context.Context, req *dialog.TLDialogGetDialogsByIdList, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error)
	DialogGetDialogsCount(ctx context.Context, req *dialog.TLDialogGetDialogsCount, callOptions ...callopt.Option) (r *tg.Int32, err error)
	DialogGetPinnedDialogs(ctx context.Context, req *dialog.TLDialogGetPinnedDialogs, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error)
	DialogReorderPinnedDialogs(ctx context.Context, req *dialog.TLDialogReorderPinnedDialogs, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogGetDialogById(ctx context.Context, req *dialog.TLDialogGetDialogById, callOptions ...callopt.Option) (r *dialog.DialogExt, err error)
	DialogGetTopMessage(ctx context.Context, req *dialog.TLDialogGetTopMessage, callOptions ...callopt.Option) (r *tg.Int32, err error)
	DialogInsertOrUpdateDialog(ctx context.Context, req *dialog.TLDialogInsertOrUpdateDialog, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogDeleteDialog(ctx context.Context, req *dialog.TLDialogDeleteDialog, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogGetUserPinnedMessage(ctx context.Context, req *dialog.TLDialogGetUserPinnedMessage, callOptions ...callopt.Option) (r *tg.Int32, err error)
	DialogUpdateUserPinnedMessage(ctx context.Context, req *dialog.TLDialogUpdateUserPinnedMessage, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogInsertOrUpdateDialogFilter(ctx context.Context, req *dialog.TLDialogInsertOrUpdateDialogFilter, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogDeleteDialogFilter(ctx context.Context, req *dialog.TLDialogDeleteDialogFilter, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogUpdateDialogFiltersOrder(ctx context.Context, req *dialog.TLDialogUpdateDialogFiltersOrder, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogGetDialogFilters(ctx context.Context, req *dialog.TLDialogGetDialogFilters, callOptions ...callopt.Option) (r *dialog.VectorDialogFilterExt, err error)
	DialogGetDialogFolder(ctx context.Context, req *dialog.TLDialogGetDialogFolder, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error)
	DialogEditPeerFolders(ctx context.Context, req *dialog.TLDialogEditPeerFolders, callOptions ...callopt.Option) (r *dialog.VectorDialogPinnedExt, err error)
	DialogGetChannelMessageReadParticipants(ctx context.Context, req *dialog.TLDialogGetChannelMessageReadParticipants, callOptions ...callopt.Option) (r *dialog.VectorLong, err error)
	DialogSetChatTheme(ctx context.Context, req *dialog.TLDialogSetChatTheme, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogSetHistoryTTL(ctx context.Context, req *dialog.TLDialogSetHistoryTTL, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogGetMyDialogsData(ctx context.Context, req *dialog.TLDialogGetMyDialogsData, callOptions ...callopt.Option) (r *dialog.DialogsData, err error)
	DialogGetSavedDialogs(ctx context.Context, req *dialog.TLDialogGetSavedDialogs, callOptions ...callopt.Option) (r *dialog.SavedDialogList, err error)
	DialogGetPinnedSavedDialogs(ctx context.Context, req *dialog.TLDialogGetPinnedSavedDialogs, callOptions ...callopt.Option) (r *dialog.SavedDialogList, err error)
	DialogToggleSavedDialogPin(ctx context.Context, req *dialog.TLDialogToggleSavedDialogPin, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogReorderPinnedSavedDialogs(ctx context.Context, req *dialog.TLDialogReorderPinnedSavedDialogs, callOptions ...callopt.Option) (r *tg.Bool, err error)
	DialogGetDialogFilter(ctx context.Context, req *dialog.TLDialogGetDialogFilter, callOptions ...callopt.Option) (r *dialog.DialogFilterExt, err error)
	DialogGetDialogFilterBySlug(ctx context.Context, req *dialog.TLDialogGetDialogFilterBySlug, callOptions ...callopt.Option) (r *dialog.DialogFilterExt, err error)
	DialogCreateDialogFilter(ctx context.Context, req *dialog.TLDialogCreateDialogFilter, callOptions ...callopt.Option) (r *dialog.DialogFilterExt, err error)
	DialogUpdateUnreadCount(ctx context.Context, req *dialog.TLDialogUpdateUnreadCount, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kDialogClient{
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

type kDialogClient struct {
	*kClient
}

func NewRPCDialogClient(cli client.Client) Client {
	return &kDialogClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kDialogClient) DialogSaveDraftMessage(ctx context.Context, req *dialog.TLDialogSaveDraftMessage, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogSaveDraftMessage(ctx, req)
}

func (p *kDialogClient) DialogClearDraftMessage(ctx context.Context, req *dialog.TLDialogClearDraftMessage, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogClearDraftMessage(ctx, req)
}

func (p *kDialogClient) DialogGetAllDrafts(ctx context.Context, req *dialog.TLDialogGetAllDrafts, callOptions ...callopt.Option) (r *dialog.VectorPeerWithDraftMessage, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetAllDrafts(ctx, req)
}

func (p *kDialogClient) DialogClearAllDrafts(ctx context.Context, req *dialog.TLDialogClearAllDrafts, callOptions ...callopt.Option) (r *dialog.VectorPeerWithDraftMessage, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogClearAllDrafts(ctx, req)
}

func (p *kDialogClient) DialogMarkDialogUnread(ctx context.Context, req *dialog.TLDialogMarkDialogUnread, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogMarkDialogUnread(ctx, req)
}

func (p *kDialogClient) DialogToggleDialogPin(ctx context.Context, req *dialog.TLDialogToggleDialogPin, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogToggleDialogPin(ctx, req)
}

func (p *kDialogClient) DialogGetDialogUnreadMarkList(ctx context.Context, req *dialog.TLDialogGetDialogUnreadMarkList, callOptions ...callopt.Option) (r *dialog.VectorDialogPeer, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogUnreadMarkList(ctx, req)
}

func (p *kDialogClient) DialogGetDialogsByOffsetDate(ctx context.Context, req *dialog.TLDialogGetDialogsByOffsetDate, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogsByOffsetDate(ctx, req)
}

func (p *kDialogClient) DialogGetDialogs(ctx context.Context, req *dialog.TLDialogGetDialogs, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogs(ctx, req)
}

func (p *kDialogClient) DialogGetDialogsByIdList(ctx context.Context, req *dialog.TLDialogGetDialogsByIdList, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogsByIdList(ctx, req)
}

func (p *kDialogClient) DialogGetDialogsCount(ctx context.Context, req *dialog.TLDialogGetDialogsCount, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogsCount(ctx, req)
}

func (p *kDialogClient) DialogGetPinnedDialogs(ctx context.Context, req *dialog.TLDialogGetPinnedDialogs, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetPinnedDialogs(ctx, req)
}

func (p *kDialogClient) DialogReorderPinnedDialogs(ctx context.Context, req *dialog.TLDialogReorderPinnedDialogs, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogReorderPinnedDialogs(ctx, req)
}

func (p *kDialogClient) DialogGetDialogById(ctx context.Context, req *dialog.TLDialogGetDialogById, callOptions ...callopt.Option) (r *dialog.DialogExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogById(ctx, req)
}

func (p *kDialogClient) DialogGetTopMessage(ctx context.Context, req *dialog.TLDialogGetTopMessage, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetTopMessage(ctx, req)
}

func (p *kDialogClient) DialogInsertOrUpdateDialog(ctx context.Context, req *dialog.TLDialogInsertOrUpdateDialog, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogInsertOrUpdateDialog(ctx, req)
}

func (p *kDialogClient) DialogDeleteDialog(ctx context.Context, req *dialog.TLDialogDeleteDialog, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogDeleteDialog(ctx, req)
}

func (p *kDialogClient) DialogGetUserPinnedMessage(ctx context.Context, req *dialog.TLDialogGetUserPinnedMessage, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetUserPinnedMessage(ctx, req)
}

func (p *kDialogClient) DialogUpdateUserPinnedMessage(ctx context.Context, req *dialog.TLDialogUpdateUserPinnedMessage, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogUpdateUserPinnedMessage(ctx, req)
}

func (p *kDialogClient) DialogInsertOrUpdateDialogFilter(ctx context.Context, req *dialog.TLDialogInsertOrUpdateDialogFilter, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogInsertOrUpdateDialogFilter(ctx, req)
}

func (p *kDialogClient) DialogDeleteDialogFilter(ctx context.Context, req *dialog.TLDialogDeleteDialogFilter, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogDeleteDialogFilter(ctx, req)
}

func (p *kDialogClient) DialogUpdateDialogFiltersOrder(ctx context.Context, req *dialog.TLDialogUpdateDialogFiltersOrder, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogUpdateDialogFiltersOrder(ctx, req)
}

func (p *kDialogClient) DialogGetDialogFilters(ctx context.Context, req *dialog.TLDialogGetDialogFilters, callOptions ...callopt.Option) (r *dialog.VectorDialogFilterExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogFilters(ctx, req)
}

func (p *kDialogClient) DialogGetDialogFolder(ctx context.Context, req *dialog.TLDialogGetDialogFolder, callOptions ...callopt.Option) (r *dialog.VectorDialogExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogFolder(ctx, req)
}

func (p *kDialogClient) DialogEditPeerFolders(ctx context.Context, req *dialog.TLDialogEditPeerFolders, callOptions ...callopt.Option) (r *dialog.VectorDialogPinnedExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogEditPeerFolders(ctx, req)
}

func (p *kDialogClient) DialogGetChannelMessageReadParticipants(ctx context.Context, req *dialog.TLDialogGetChannelMessageReadParticipants, callOptions ...callopt.Option) (r *dialog.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetChannelMessageReadParticipants(ctx, req)
}

func (p *kDialogClient) DialogSetChatTheme(ctx context.Context, req *dialog.TLDialogSetChatTheme, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogSetChatTheme(ctx, req)
}

func (p *kDialogClient) DialogSetHistoryTTL(ctx context.Context, req *dialog.TLDialogSetHistoryTTL, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogSetHistoryTTL(ctx, req)
}

func (p *kDialogClient) DialogGetMyDialogsData(ctx context.Context, req *dialog.TLDialogGetMyDialogsData, callOptions ...callopt.Option) (r *dialog.DialogsData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetMyDialogsData(ctx, req)
}

func (p *kDialogClient) DialogGetSavedDialogs(ctx context.Context, req *dialog.TLDialogGetSavedDialogs, callOptions ...callopt.Option) (r *dialog.SavedDialogList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetSavedDialogs(ctx, req)
}

func (p *kDialogClient) DialogGetPinnedSavedDialogs(ctx context.Context, req *dialog.TLDialogGetPinnedSavedDialogs, callOptions ...callopt.Option) (r *dialog.SavedDialogList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetPinnedSavedDialogs(ctx, req)
}

func (p *kDialogClient) DialogToggleSavedDialogPin(ctx context.Context, req *dialog.TLDialogToggleSavedDialogPin, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogToggleSavedDialogPin(ctx, req)
}

func (p *kDialogClient) DialogReorderPinnedSavedDialogs(ctx context.Context, req *dialog.TLDialogReorderPinnedSavedDialogs, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogReorderPinnedSavedDialogs(ctx, req)
}

func (p *kDialogClient) DialogGetDialogFilter(ctx context.Context, req *dialog.TLDialogGetDialogFilter, callOptions ...callopt.Option) (r *dialog.DialogFilterExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogFilter(ctx, req)
}

func (p *kDialogClient) DialogGetDialogFilterBySlug(ctx context.Context, req *dialog.TLDialogGetDialogFilterBySlug, callOptions ...callopt.Option) (r *dialog.DialogFilterExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogGetDialogFilterBySlug(ctx, req)
}

func (p *kDialogClient) DialogCreateDialogFilter(ctx context.Context, req *dialog.TLDialogCreateDialogFilter, callOptions ...callopt.Option) (r *dialog.DialogFilterExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogCreateDialogFilter(ctx, req)
}

func (p *kDialogClient) DialogUpdateUnreadCount(ctx context.Context, req *dialog.TLDialogUpdateUnreadCount, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DialogUpdateUnreadCount(ctx, req)
}
