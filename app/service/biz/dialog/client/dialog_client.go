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

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type DialogClient interface {
	DialogSaveDraftMessage(ctx context.Context, in *dialog.TLDialogSaveDraftMessage) (*mtproto.Bool, error)
	DialogClearDraftMessage(ctx context.Context, in *dialog.TLDialogClearDraftMessage) (*mtproto.Bool, error)
	DialogGetAllDrafts(ctx context.Context, in *dialog.TLDialogGetAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error)
	DialogClearAllDrafts(ctx context.Context, in *dialog.TLDialogClearAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error)
	DialogMarkDialogUnread(ctx context.Context, in *dialog.TLDialogMarkDialogUnread) (*mtproto.Bool, error)
	DialogToggleDialogPin(ctx context.Context, in *dialog.TLDialogToggleDialogPin) (*mtproto.Int32, error)
	DialogGetDialogUnreadMarkList(ctx context.Context, in *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.Vector_DialogPeer, error)
	DialogGetDialogsByOffsetDate(ctx context.Context, in *dialog.TLDialogGetDialogsByOffsetDate) (*dialog.Vector_DialogExt, error)
	DialogGetDialogs(ctx context.Context, in *dialog.TLDialogGetDialogs) (*dialog.Vector_DialogExt, error)
	DialogGetDialogsByIdList(ctx context.Context, in *dialog.TLDialogGetDialogsByIdList) (*dialog.Vector_DialogExt, error)
	DialogGetDialogsCount(ctx context.Context, in *dialog.TLDialogGetDialogsCount) (*mtproto.Int32, error)
	DialogGetPinnedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedDialogs) (*dialog.Vector_DialogExt, error)
	DialogReorderPinnedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedDialogs) (*mtproto.Bool, error)
	DialogGetDialogById(ctx context.Context, in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error)
	DialogGetTopMessage(ctx context.Context, in *dialog.TLDialogGetTopMessage) (*mtproto.Int32, error)
	DialogUpdateReadInbox(ctx context.Context, in *dialog.TLDialogUpdateReadInbox) (*mtproto.Bool, error)
	DialogUpdateReadOutbox(ctx context.Context, in *dialog.TLDialogUpdateReadOutbox) (*mtproto.Bool, error)
	DialogInsertOrUpdateDialog(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialog) (*mtproto.Bool, error)
	DialogDeleteDialog(ctx context.Context, in *dialog.TLDialogDeleteDialog) (*mtproto.Bool, error)
	DialogGetUserPinnedMessage(ctx context.Context, in *dialog.TLDialogGetUserPinnedMessage) (*mtproto.Int32, error)
	DialogUpdateUserPinnedMessage(ctx context.Context, in *dialog.TLDialogUpdateUserPinnedMessage) (*mtproto.Bool, error)
	DialogInsertOrUpdateDialogFilter(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialogFilter) (*mtproto.Bool, error)
	DialogDeleteDialogFilter(ctx context.Context, in *dialog.TLDialogDeleteDialogFilter) (*mtproto.Bool, error)
	DialogUpdateDialogFiltersOrder(ctx context.Context, in *dialog.TLDialogUpdateDialogFiltersOrder) (*mtproto.Bool, error)
	DialogGetDialogFilters(ctx context.Context, in *dialog.TLDialogGetDialogFilters) (*dialog.Vector_DialogFilterExt, error)
	DialogGetDialogFolder(ctx context.Context, in *dialog.TLDialogGetDialogFolder) (*dialog.Vector_DialogExt, error)
	DialogEditPeerFolders(ctx context.Context, in *dialog.TLDialogEditPeerFolders) (*dialog.Vector_DialogPinnedExt, error)
	DialogGetChannelMessageReadParticipants(ctx context.Context, in *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.Vector_Long, error)
	DialogSetChatTheme(ctx context.Context, in *dialog.TLDialogSetChatTheme) (*mtproto.Bool, error)
	DialogSetHistoryTTL(ctx context.Context, in *dialog.TLDialogSetHistoryTTL) (*mtproto.Bool, error)
	DialogGetMyDialogsData(ctx context.Context, in *dialog.TLDialogGetMyDialogsData) (*dialog.DialogsData, error)
	DialogGetSavedDialogs(ctx context.Context, in *dialog.TLDialogGetSavedDialogs) (*dialog.SavedDialogList, error)
	DialogGetPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedSavedDialogs) (*dialog.SavedDialogList, error)
	DialogToggleSavedDialogPin(ctx context.Context, in *dialog.TLDialogToggleSavedDialogPin) (*mtproto.Bool, error)
	DialogReorderPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedSavedDialogs) (*mtproto.Bool, error)
	DialogGetDialogFilter(ctx context.Context, in *dialog.TLDialogGetDialogFilter) (*dialog.DialogFilterExt, error)
	DialogGetDialogFilterBySlug(ctx context.Context, in *dialog.TLDialogGetDialogFilterBySlug) (*dialog.DialogFilterExt, error)
	DialogCreateDialogFilter(ctx context.Context, in *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error)
	DialogUpdateUnreadCount(ctx context.Context, in *dialog.TLDialogUpdateUnreadCount) (*mtproto.Bool, error)
}

type defaultDialogClient struct {
	cli zrpc.Client
}

func NewDialogClient(cli zrpc.Client) DialogClient {
	return &defaultDialogClient{
		cli: cli,
	}
}

// DialogSaveDraftMessage
// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage = Bool;
func (m *defaultDialogClient) DialogSaveDraftMessage(ctx context.Context, in *dialog.TLDialogSaveDraftMessage) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogSaveDraftMessage(ctx, in)
}

// DialogClearDraftMessage
// dialog.clearDraftMessage user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultDialogClient) DialogClearDraftMessage(ctx context.Context, in *dialog.TLDialogClearDraftMessage) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogClearDraftMessage(ctx, in)
}

// DialogGetAllDrafts
// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (m *defaultDialogClient) DialogGetAllDrafts(ctx context.Context, in *dialog.TLDialogGetAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetAllDrafts(ctx, in)
}

// DialogClearAllDrafts
// dialog.clearAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (m *defaultDialogClient) DialogClearAllDrafts(ctx context.Context, in *dialog.TLDialogClearAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogClearAllDrafts(ctx, in)
}

// DialogMarkDialogUnread
// dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;
func (m *defaultDialogClient) DialogMarkDialogUnread(ctx context.Context, in *dialog.TLDialogMarkDialogUnread) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogMarkDialogUnread(ctx, in)
}

// DialogToggleDialogPin
// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;
func (m *defaultDialogClient) DialogToggleDialogPin(ctx context.Context, in *dialog.TLDialogToggleDialogPin) (*mtproto.Int32, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogToggleDialogPin(ctx, in)
}

// DialogGetDialogUnreadMarkList
// dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;
func (m *defaultDialogClient) DialogGetDialogUnreadMarkList(ctx context.Context, in *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.Vector_DialogPeer, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogUnreadMarkList(ctx, in)
}

// DialogGetDialogsByOffsetDate
// dialog.getDialogsByOffsetDate user_id:long exclude_pinned:Bool offset_date:int limit:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogsByOffsetDate(ctx context.Context, in *dialog.TLDialogGetDialogsByOffsetDate) (*dialog.Vector_DialogExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogsByOffsetDate(ctx, in)
}

// DialogGetDialogs
// dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogs(ctx context.Context, in *dialog.TLDialogGetDialogs) (*dialog.Vector_DialogExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogs(ctx, in)
}

// DialogGetDialogsByIdList
// dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogsByIdList(ctx context.Context, in *dialog.TLDialogGetDialogsByIdList) (*dialog.Vector_DialogExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogsByIdList(ctx, in)
}

// DialogGetDialogsCount
// dialog.getDialogsCount user_id:long exclude_pinned:Bool folder_id:int = Int32;
func (m *defaultDialogClient) DialogGetDialogsCount(ctx context.Context, in *dialog.TLDialogGetDialogsCount) (*mtproto.Int32, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogsCount(ctx, in)
}

// DialogGetPinnedDialogs
// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetPinnedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedDialogs) (*dialog.Vector_DialogExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetPinnedDialogs(ctx, in)
}

// DialogReorderPinnedDialogs
// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> = Bool;
func (m *defaultDialogClient) DialogReorderPinnedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedDialogs) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogReorderPinnedDialogs(ctx, in)
}

// DialogGetDialogById
// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
func (m *defaultDialogClient) DialogGetDialogById(ctx context.Context, in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogById(ctx, in)
}

// DialogGetTopMessage
// dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultDialogClient) DialogGetTopMessage(ctx context.Context, in *dialog.TLDialogGetTopMessage) (*mtproto.Int32, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetTopMessage(ctx, in)
}

// DialogUpdateReadInbox
// dialog.updateReadInbox user_id:long peer_type:int peer_id:long read_inbox_id:int = Bool;
func (m *defaultDialogClient) DialogUpdateReadInbox(ctx context.Context, in *dialog.TLDialogUpdateReadInbox) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogUpdateReadInbox(ctx, in)
}

// DialogUpdateReadOutbox
// dialog.updateReadOutbox user_id:long peer_type:int peer_id:long read_outbox_id:int = Bool;
func (m *defaultDialogClient) DialogUpdateReadOutbox(ctx context.Context, in *dialog.TLDialogUpdateReadOutbox) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogUpdateReadOutbox(ctx, in)
}

// DialogInsertOrUpdateDialog
// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true date2:flags.5?long = Bool;
func (m *defaultDialogClient) DialogInsertOrUpdateDialog(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialog) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogInsertOrUpdateDialog(ctx, in)
}

// DialogDeleteDialog
// dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultDialogClient) DialogDeleteDialog(ctx context.Context, in *dialog.TLDialogDeleteDialog) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogDeleteDialog(ctx, in)
}

// DialogGetUserPinnedMessage
// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultDialogClient) DialogGetUserPinnedMessage(ctx context.Context, in *dialog.TLDialogGetUserPinnedMessage) (*mtproto.Int32, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetUserPinnedMessage(ctx, in)
}

// DialogUpdateUserPinnedMessage
// dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;
func (m *defaultDialogClient) DialogUpdateUserPinnedMessage(ctx context.Context, in *dialog.TLDialogUpdateUserPinnedMessage) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogUpdateUserPinnedMessage(ctx, in)
}

// DialogInsertOrUpdateDialogFilter
// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
func (m *defaultDialogClient) DialogInsertOrUpdateDialogFilter(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialogFilter) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogInsertOrUpdateDialogFilter(ctx, in)
}

// DialogDeleteDialogFilter
// dialog.deleteDialogFilter user_id:long id:int = Bool;
func (m *defaultDialogClient) DialogDeleteDialogFilter(ctx context.Context, in *dialog.TLDialogDeleteDialogFilter) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogDeleteDialogFilter(ctx, in)
}

// DialogUpdateDialogFiltersOrder
// dialog.updateDialogFiltersOrder user_id:long order:Vector<int> = Bool;
func (m *defaultDialogClient) DialogUpdateDialogFiltersOrder(ctx context.Context, in *dialog.TLDialogUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogUpdateDialogFiltersOrder(ctx, in)
}

// DialogGetDialogFilters
// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;
func (m *defaultDialogClient) DialogGetDialogFilters(ctx context.Context, in *dialog.TLDialogGetDialogFilters) (*dialog.Vector_DialogFilterExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogFilters(ctx, in)
}

// DialogGetDialogFolder
// dialog.getDialogFolder user_id:long folder_id:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogFolder(ctx context.Context, in *dialog.TLDialogGetDialogFolder) (*dialog.Vector_DialogExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogFolder(ctx, in)
}

// DialogEditPeerFolders
// dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int = Vector<DialogPinnedExt>;
func (m *defaultDialogClient) DialogEditPeerFolders(ctx context.Context, in *dialog.TLDialogEditPeerFolders) (*dialog.Vector_DialogPinnedExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogEditPeerFolders(ctx, in)
}

// DialogGetChannelMessageReadParticipants
// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;
func (m *defaultDialogClient) DialogGetChannelMessageReadParticipants(ctx context.Context, in *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.Vector_Long, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetChannelMessageReadParticipants(ctx, in)
}

// DialogSetChatTheme
// dialog.setChatTheme user_id:long peer_type:int peer_id:long theme_emoticon:string = Bool;
func (m *defaultDialogClient) DialogSetChatTheme(ctx context.Context, in *dialog.TLDialogSetChatTheme) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogSetChatTheme(ctx, in)
}

// DialogSetHistoryTTL
// dialog.setHistoryTTL user_id:long peer_type:int peer_id:long ttl_period:int = Bool;
func (m *defaultDialogClient) DialogSetHistoryTTL(ctx context.Context, in *dialog.TLDialogSetHistoryTTL) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogSetHistoryTTL(ctx, in)
}

// DialogGetMyDialogsData
// dialog.getMyDialogsData flags:# user_id:long user:flags.0?true chat:flags.1?true channel:flags.2?true = DialogsData;
func (m *defaultDialogClient) DialogGetMyDialogsData(ctx context.Context, in *dialog.TLDialogGetMyDialogsData) (*dialog.DialogsData, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetMyDialogsData(ctx, in)
}

// DialogGetSavedDialogs
// dialog.getSavedDialogs user_id:long exclude_pinned:Bool offset_date:int offset_id:int offset_peer:PeerUtil limit:int = SavedDialogList;
func (m *defaultDialogClient) DialogGetSavedDialogs(ctx context.Context, in *dialog.TLDialogGetSavedDialogs) (*dialog.SavedDialogList, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetSavedDialogs(ctx, in)
}

// DialogGetPinnedSavedDialogs
// dialog.getPinnedSavedDialogs user_id:long = SavedDialogList;
func (m *defaultDialogClient) DialogGetPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedSavedDialogs) (*dialog.SavedDialogList, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetPinnedSavedDialogs(ctx, in)
}

// DialogToggleSavedDialogPin
// dialog.toggleSavedDialogPin user_id:long peer:PeerUtil pinned:Bool = Bool;
func (m *defaultDialogClient) DialogToggleSavedDialogPin(ctx context.Context, in *dialog.TLDialogToggleSavedDialogPin) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogToggleSavedDialogPin(ctx, in)
}

// DialogReorderPinnedSavedDialogs
// dialog.reorderPinnedSavedDialogs user_id:long force:Bool order:Vector<PeerUtil> = Bool;
func (m *defaultDialogClient) DialogReorderPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedSavedDialogs) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogReorderPinnedSavedDialogs(ctx, in)
}

// DialogGetDialogFilter
// dialog.getDialogFilter user_id:long id:int = DialogFilterExt;
func (m *defaultDialogClient) DialogGetDialogFilter(ctx context.Context, in *dialog.TLDialogGetDialogFilter) (*dialog.DialogFilterExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogFilter(ctx, in)
}

// DialogGetDialogFilterBySlug
// dialog.getDialogFilterBySlug user_id:long slug:string = DialogFilterExt;
func (m *defaultDialogClient) DialogGetDialogFilterBySlug(ctx context.Context, in *dialog.TLDialogGetDialogFilterBySlug) (*dialog.DialogFilterExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogGetDialogFilterBySlug(ctx, in)
}

// DialogCreateDialogFilter
// dialog.createDialogFilter user_id:long dialog_filter:DialogFilterExt = DialogFilterExt;
func (m *defaultDialogClient) DialogCreateDialogFilter(ctx context.Context, in *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogCreateDialogFilter(ctx, in)
}

// DialogUpdateUnreadCount
// dialog.updateUnreadCount user_id:long peer_type:int peer_id:long unread_count:flags.0?int unread_mentions_count:flags.1?int unread_reactions_count:flags.2?int = Bool;
func (m *defaultDialogClient) DialogUpdateUnreadCount(ctx context.Context, in *dialog.TLDialogUpdateUnreadCount) (*mtproto.Bool, error) {
	client := dialog.NewRPCDialogClient(m.cli.Conn())
	return client.DialogUpdateUnreadCount(ctx, in)
}
