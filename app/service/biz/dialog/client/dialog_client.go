/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dialogclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog/dialogservice"

	"github.com/cloudwego/kitex/client"
)

type DialogClient interface {
	DialogSaveDraftMessage(ctx context.Context, in *dialog.TLDialogSaveDraftMessage) (*tg.Bool, error)
	DialogClearDraftMessage(ctx context.Context, in *dialog.TLDialogClearDraftMessage) (*tg.Bool, error)
	DialogGetAllDrafts(ctx context.Context, in *dialog.TLDialogGetAllDrafts) (*dialog.VectorPeerWithDraftMessage, error)
	DialogClearAllDrafts(ctx context.Context, in *dialog.TLDialogClearAllDrafts) (*dialog.VectorPeerWithDraftMessage, error)
	DialogMarkDialogUnread(ctx context.Context, in *dialog.TLDialogMarkDialogUnread) (*tg.Bool, error)
	DialogToggleDialogPin(ctx context.Context, in *dialog.TLDialogToggleDialogPin) (*tg.Int32, error)
	DialogGetDialogUnreadMarkList(ctx context.Context, in *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.VectorDialogPeer, error)
	DialogGetDialogsByOffsetDate(ctx context.Context, in *dialog.TLDialogGetDialogsByOffsetDate) (*dialog.VectorDialogExt, error)
	DialogGetDialogs(ctx context.Context, in *dialog.TLDialogGetDialogs) (*dialog.VectorDialogExt, error)
	DialogGetDialogsByIdList(ctx context.Context, in *dialog.TLDialogGetDialogsByIdList) (*dialog.VectorDialogExt, error)
	DialogGetDialogsCount(ctx context.Context, in *dialog.TLDialogGetDialogsCount) (*tg.Int32, error)
	DialogGetPinnedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedDialogs) (*dialog.VectorDialogExt, error)
	DialogReorderPinnedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedDialogs) (*tg.Bool, error)
	DialogGetDialogById(ctx context.Context, in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error)
	DialogGetTopMessage(ctx context.Context, in *dialog.TLDialogGetTopMessage) (*tg.Int32, error)
	DialogInsertOrUpdateDialog(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialog) (*tg.Bool, error)
	DialogDeleteDialog(ctx context.Context, in *dialog.TLDialogDeleteDialog) (*tg.Bool, error)
	DialogGetUserPinnedMessage(ctx context.Context, in *dialog.TLDialogGetUserPinnedMessage) (*tg.Int32, error)
	DialogUpdateUserPinnedMessage(ctx context.Context, in *dialog.TLDialogUpdateUserPinnedMessage) (*tg.Bool, error)
	DialogInsertOrUpdateDialogFilter(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialogFilter) (*tg.Bool, error)
	DialogDeleteDialogFilter(ctx context.Context, in *dialog.TLDialogDeleteDialogFilter) (*tg.Bool, error)
	DialogUpdateDialogFiltersOrder(ctx context.Context, in *dialog.TLDialogUpdateDialogFiltersOrder) (*tg.Bool, error)
	DialogGetDialogFilters(ctx context.Context, in *dialog.TLDialogGetDialogFilters) (*dialog.VectorDialogFilterExt, error)
	DialogGetDialogFolder(ctx context.Context, in *dialog.TLDialogGetDialogFolder) (*dialog.VectorDialogExt, error)
	DialogEditPeerFolders(ctx context.Context, in *dialog.TLDialogEditPeerFolders) (*dialog.VectorDialogPinnedExt, error)
	DialogGetChannelMessageReadParticipants(ctx context.Context, in *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.VectorLong, error)
	DialogSetChatTheme(ctx context.Context, in *dialog.TLDialogSetChatTheme) (*tg.Bool, error)
	DialogSetHistoryTTL(ctx context.Context, in *dialog.TLDialogSetHistoryTTL) (*tg.Bool, error)
	DialogGetMyDialogsData(ctx context.Context, in *dialog.TLDialogGetMyDialogsData) (*dialog.DialogsData, error)
	DialogGetSavedDialogs(ctx context.Context, in *dialog.TLDialogGetSavedDialogs) (*dialog.SavedDialogList, error)
	DialogGetPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedSavedDialogs) (*dialog.SavedDialogList, error)
	DialogToggleSavedDialogPin(ctx context.Context, in *dialog.TLDialogToggleSavedDialogPin) (*tg.Bool, error)
	DialogReorderPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedSavedDialogs) (*tg.Bool, error)
	DialogGetDialogFilter(ctx context.Context, in *dialog.TLDialogGetDialogFilter) (*dialog.DialogFilterExt, error)
	DialogGetDialogFilterBySlug(ctx context.Context, in *dialog.TLDialogGetDialogFilterBySlug) (*dialog.DialogFilterExt, error)
	DialogCreateDialogFilter(ctx context.Context, in *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error)
	DialogUpdateUnreadCount(ctx context.Context, in *dialog.TLDialogUpdateUnreadCount) (*tg.Bool, error)
}

type defaultDialogClient struct {
	cli client.Client
}

func NewDialogClient(cli client.Client) DialogClient {
	return &defaultDialogClient{
		cli: cli,
	}
}

// DialogSaveDraftMessage
// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage = Bool;
func (m *defaultDialogClient) DialogSaveDraftMessage(ctx context.Context, in *dialog.TLDialogSaveDraftMessage) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogSaveDraftMessage(ctx, in)
}

// DialogClearDraftMessage
// dialog.clearDraftMessage user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultDialogClient) DialogClearDraftMessage(ctx context.Context, in *dialog.TLDialogClearDraftMessage) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogClearDraftMessage(ctx, in)
}

// DialogGetAllDrafts
// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (m *defaultDialogClient) DialogGetAllDrafts(ctx context.Context, in *dialog.TLDialogGetAllDrafts) (*dialog.VectorPeerWithDraftMessage, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetAllDrafts(ctx, in)
}

// DialogClearAllDrafts
// dialog.clearAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (m *defaultDialogClient) DialogClearAllDrafts(ctx context.Context, in *dialog.TLDialogClearAllDrafts) (*dialog.VectorPeerWithDraftMessage, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogClearAllDrafts(ctx, in)
}

// DialogMarkDialogUnread
// dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;
func (m *defaultDialogClient) DialogMarkDialogUnread(ctx context.Context, in *dialog.TLDialogMarkDialogUnread) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogMarkDialogUnread(ctx, in)
}

// DialogToggleDialogPin
// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;
func (m *defaultDialogClient) DialogToggleDialogPin(ctx context.Context, in *dialog.TLDialogToggleDialogPin) (*tg.Int32, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogToggleDialogPin(ctx, in)
}

// DialogGetDialogUnreadMarkList
// dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;
func (m *defaultDialogClient) DialogGetDialogUnreadMarkList(ctx context.Context, in *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.VectorDialogPeer, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogUnreadMarkList(ctx, in)
}

// DialogGetDialogsByOffsetDate
// dialog.getDialogsByOffsetDate user_id:long exclude_pinned:Bool offset_date:int limit:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogsByOffsetDate(ctx context.Context, in *dialog.TLDialogGetDialogsByOffsetDate) (*dialog.VectorDialogExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogsByOffsetDate(ctx, in)
}

// DialogGetDialogs
// dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogs(ctx context.Context, in *dialog.TLDialogGetDialogs) (*dialog.VectorDialogExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogs(ctx, in)
}

// DialogGetDialogsByIdList
// dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogsByIdList(ctx context.Context, in *dialog.TLDialogGetDialogsByIdList) (*dialog.VectorDialogExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogsByIdList(ctx, in)
}

// DialogGetDialogsCount
// dialog.getDialogsCount user_id:long exclude_pinned:Bool folder_id:int = Int32;
func (m *defaultDialogClient) DialogGetDialogsCount(ctx context.Context, in *dialog.TLDialogGetDialogsCount) (*tg.Int32, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogsCount(ctx, in)
}

// DialogGetPinnedDialogs
// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetPinnedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedDialogs) (*dialog.VectorDialogExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetPinnedDialogs(ctx, in)
}

// DialogReorderPinnedDialogs
// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> = Bool;
func (m *defaultDialogClient) DialogReorderPinnedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedDialogs) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogReorderPinnedDialogs(ctx, in)
}

// DialogGetDialogById
// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
func (m *defaultDialogClient) DialogGetDialogById(ctx context.Context, in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogById(ctx, in)
}

// DialogGetTopMessage
// dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultDialogClient) DialogGetTopMessage(ctx context.Context, in *dialog.TLDialogGetTopMessage) (*tg.Int32, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetTopMessage(ctx, in)
}

// DialogInsertOrUpdateDialog
// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true date2:flags.5?long pinned_msg_id:flags.6?int = Bool;
func (m *defaultDialogClient) DialogInsertOrUpdateDialog(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialog) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogInsertOrUpdateDialog(ctx, in)
}

// DialogDeleteDialog
// dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultDialogClient) DialogDeleteDialog(ctx context.Context, in *dialog.TLDialogDeleteDialog) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogDeleteDialog(ctx, in)
}

// DialogGetUserPinnedMessage
// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultDialogClient) DialogGetUserPinnedMessage(ctx context.Context, in *dialog.TLDialogGetUserPinnedMessage) (*tg.Int32, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetUserPinnedMessage(ctx, in)
}

// DialogUpdateUserPinnedMessage
// dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;
func (m *defaultDialogClient) DialogUpdateUserPinnedMessage(ctx context.Context, in *dialog.TLDialogUpdateUserPinnedMessage) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogUpdateUserPinnedMessage(ctx, in)
}

// DialogInsertOrUpdateDialogFilter
// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
func (m *defaultDialogClient) DialogInsertOrUpdateDialogFilter(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialogFilter) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogInsertOrUpdateDialogFilter(ctx, in)
}

// DialogDeleteDialogFilter
// dialog.deleteDialogFilter user_id:long id:int = Bool;
func (m *defaultDialogClient) DialogDeleteDialogFilter(ctx context.Context, in *dialog.TLDialogDeleteDialogFilter) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogDeleteDialogFilter(ctx, in)
}

// DialogUpdateDialogFiltersOrder
// dialog.updateDialogFiltersOrder user_id:long order:Vector<int> = Bool;
func (m *defaultDialogClient) DialogUpdateDialogFiltersOrder(ctx context.Context, in *dialog.TLDialogUpdateDialogFiltersOrder) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogUpdateDialogFiltersOrder(ctx, in)
}

// DialogGetDialogFilters
// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;
func (m *defaultDialogClient) DialogGetDialogFilters(ctx context.Context, in *dialog.TLDialogGetDialogFilters) (*dialog.VectorDialogFilterExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogFilters(ctx, in)
}

// DialogGetDialogFolder
// dialog.getDialogFolder user_id:long folder_id:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogFolder(ctx context.Context, in *dialog.TLDialogGetDialogFolder) (*dialog.VectorDialogExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogFolder(ctx, in)
}

// DialogEditPeerFolders
// dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int = Vector<DialogPinnedExt>;
func (m *defaultDialogClient) DialogEditPeerFolders(ctx context.Context, in *dialog.TLDialogEditPeerFolders) (*dialog.VectorDialogPinnedExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogEditPeerFolders(ctx, in)
}

// DialogGetChannelMessageReadParticipants
// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;
func (m *defaultDialogClient) DialogGetChannelMessageReadParticipants(ctx context.Context, in *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.VectorLong, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetChannelMessageReadParticipants(ctx, in)
}

// DialogSetChatTheme
// dialog.setChatTheme user_id:long peer_type:int peer_id:long theme_emoticon:string = Bool;
func (m *defaultDialogClient) DialogSetChatTheme(ctx context.Context, in *dialog.TLDialogSetChatTheme) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogSetChatTheme(ctx, in)
}

// DialogSetHistoryTTL
// dialog.setHistoryTTL user_id:long peer_type:int peer_id:long ttl_period:int = Bool;
func (m *defaultDialogClient) DialogSetHistoryTTL(ctx context.Context, in *dialog.TLDialogSetHistoryTTL) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogSetHistoryTTL(ctx, in)
}

// DialogGetMyDialogsData
// dialog.getMyDialogsData flags:# user_id:long user:flags.0?true chat:flags.1?true channel:flags.2?true = DialogsData;
func (m *defaultDialogClient) DialogGetMyDialogsData(ctx context.Context, in *dialog.TLDialogGetMyDialogsData) (*dialog.DialogsData, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetMyDialogsData(ctx, in)
}

// DialogGetSavedDialogs
// dialog.getSavedDialogs user_id:long exclude_pinned:Bool offset_date:int offset_id:int offset_peer:PeerUtil limit:int = SavedDialogList;
func (m *defaultDialogClient) DialogGetSavedDialogs(ctx context.Context, in *dialog.TLDialogGetSavedDialogs) (*dialog.SavedDialogList, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetSavedDialogs(ctx, in)
}

// DialogGetPinnedSavedDialogs
// dialog.getPinnedSavedDialogs user_id:long = SavedDialogList;
func (m *defaultDialogClient) DialogGetPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedSavedDialogs) (*dialog.SavedDialogList, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetPinnedSavedDialogs(ctx, in)
}

// DialogToggleSavedDialogPin
// dialog.toggleSavedDialogPin user_id:long peer:PeerUtil pinned:Bool = Bool;
func (m *defaultDialogClient) DialogToggleSavedDialogPin(ctx context.Context, in *dialog.TLDialogToggleSavedDialogPin) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogToggleSavedDialogPin(ctx, in)
}

// DialogReorderPinnedSavedDialogs
// dialog.reorderPinnedSavedDialogs user_id:long force:Bool order:Vector<PeerUtil> = Bool;
func (m *defaultDialogClient) DialogReorderPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedSavedDialogs) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogReorderPinnedSavedDialogs(ctx, in)
}

// DialogGetDialogFilter
// dialog.getDialogFilter user_id:long id:int = DialogFilterExt;
func (m *defaultDialogClient) DialogGetDialogFilter(ctx context.Context, in *dialog.TLDialogGetDialogFilter) (*dialog.DialogFilterExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogFilter(ctx, in)
}

// DialogGetDialogFilterBySlug
// dialog.getDialogFilterBySlug user_id:long slug:string = DialogFilterExt;
func (m *defaultDialogClient) DialogGetDialogFilterBySlug(ctx context.Context, in *dialog.TLDialogGetDialogFilterBySlug) (*dialog.DialogFilterExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogGetDialogFilterBySlug(ctx, in)
}

// DialogCreateDialogFilter
// dialog.createDialogFilter user_id:long dialog_filter:DialogFilterExt = DialogFilterExt;
func (m *defaultDialogClient) DialogCreateDialogFilter(ctx context.Context, in *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogCreateDialogFilter(ctx, in)
}

// DialogUpdateUnreadCount
// dialog.updateUnreadCount flags:# user_id:long peer_type:int peer_id:long unread_count:flags.0?int unread_mentions_count:flags.1?int unread_reactions_count:flags.2?int = Bool;
func (m *defaultDialogClient) DialogUpdateUnreadCount(ctx context.Context, in *dialog.TLDialogUpdateUnreadCount) (*tg.Bool, error) {
	cli := dialogservice.NewRPCDialogClient(m.cli)
	return cli.DialogUpdateUnreadCount(ctx, in)
}
