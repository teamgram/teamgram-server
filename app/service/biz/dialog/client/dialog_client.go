/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dialogclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog/dialogservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type DialogClient interface {
	DialogSaveDraftMessage(ctx context.Context, in *dialog.TLDialogSaveDraftMessage) (*tg.Bool, error)
	DialogClearDraftMessage(ctx context.Context, in *dialog.TLDialogClearDraftMessage) (*tg.Bool, error)
	DialogClearDraftAfterSend(ctx context.Context, in *dialog.TLDialogClearDraftAfterSend) (*tg.Bool, error)
	DialogGetAllDrafts(ctx context.Context, in *dialog.TLDialogGetAllDrafts) (*dialog.VectorPeerWithDraftMessage, error)
	DialogClearAllDrafts(ctx context.Context, in *dialog.TLDialogClearAllDrafts) (*dialog.VectorPeerWithDraftMessage, error)
	DialogMarkDialogUnread(ctx context.Context, in *dialog.TLDialogMarkDialogUnread) (*tg.Bool, error)
	DialogToggleDialogPin(ctx context.Context, in *dialog.TLDialogToggleDialogPin) (*tg.Int32, error)
	DialogGetDialogUnreadMarkList(ctx context.Context, in *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.VectorDialogPeer, error)
	DialogReorderPinnedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedDialogs) (*tg.Bool, error)
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
	DialogGetDialogsV2(ctx context.Context, in *dialog.TLDialogGetDialogsV2) (*dialog.DialogPage, error)
	DialogGetPeerDialogsV2(ctx context.Context, in *dialog.TLDialogGetPeerDialogsV2) (*dialog.VectorDialogExtV2, error)
	DialogGetPinnedDialogsV2(ctx context.Context, in *dialog.TLDialogGetPinnedDialogsV2) (*dialog.VectorDialogExtV2, error)
	DialogGetDialogByPeerV2(ctx context.Context, in *dialog.TLDialogGetDialogByPeerV2) (*dialog.DialogExtV2, error)
	DialogBatchGetDialogExtras(ctx context.Context, in *dialog.TLDialogBatchGetDialogExtras) (*dialog.VectorDialogExtras, error)
	DialogGetChannelMessageReadParticipants(ctx context.Context, in *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.VectorLong, error)
	DialogSetChatTheme(ctx context.Context, in *dialog.TLDialogSetChatTheme) (*tg.Bool, error)
	DialogSetHistoryTTL(ctx context.Context, in *dialog.TLDialogSetHistoryTTL) (*tg.Bool, error)
	DialogGetSavedDialogs(ctx context.Context, in *dialog.TLDialogGetSavedDialogs) (*dialog.SavedDialogList, error)
	DialogGetPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedSavedDialogs) (*dialog.SavedDialogList, error)
	DialogUpsertSavedDialogFromMessage(ctx context.Context, in *dialog.TLDialogUpsertSavedDialogFromMessage) (*tg.Bool, error)
	DialogToggleSavedDialogPin(ctx context.Context, in *dialog.TLDialogToggleSavedDialogPin) (*tg.Bool, error)
	DialogReorderPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedSavedDialogs) (*tg.Bool, error)
	DialogGetDialogFilter(ctx context.Context, in *dialog.TLDialogGetDialogFilter) (*dialog.DialogFilterExt, error)
	DialogGetDialogFilterBySlug(ctx context.Context, in *dialog.TLDialogGetDialogFilterBySlug) (*dialog.DialogFilterExt, error)
	DialogCreateDialogFilter(ctx context.Context, in *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error)
	DialogUpdateUnreadCount(ctx context.Context, in *dialog.TLDialogUpdateUnreadCount) (*tg.Bool, error)
	DialogToggleDialogFilterTags(ctx context.Context, in *dialog.TLDialogToggleDialogFilterTags) (*tg.Bool, error)
	DialogGetDialogFilterTags(ctx context.Context, in *dialog.TLDialogGetDialogFilterTags) (*tg.Bool, error)
	DialogSetChatWallpaper(ctx context.Context, in *dialog.TLDialogSetChatWallpaper) (*tg.Bool, error)
}

type defaultDialogClient struct {
	cli client.Client
	rpc dialogservice.Client
}

func NewDialogClient(cli client.Client) DialogClient {
	return &defaultDialogClient{
		cli: cli,
		rpc: dialogservice.NewRPCDialogClient(cli),
	}
}

// DialogSaveDraftMessage
// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage source_perm_auth_key_id:long operation_id:string outbox_id:long = Bool;
func (m *defaultDialogClient) DialogSaveDraftMessage(ctx context.Context, in *dialog.TLDialogSaveDraftMessage) (*tg.Bool, error) {
	return m.rpc.DialogSaveDraftMessage(ctx, in)
}

// DialogClearDraftMessage
// dialog.clearDraftMessage user_id:long peer_type:int peer_id:long source_perm_auth_key_id:long operation_id:string outbox_id:long = Bool;
func (m *defaultDialogClient) DialogClearDraftMessage(ctx context.Context, in *dialog.TLDialogClearDraftMessage) (*tg.Bool, error) {
	return m.rpc.DialogClearDraftMessage(ctx, in)
}

// DialogClearDraftAfterSend
// dialog.clearDraftAfterSend user_id:long peer_type:int peer_id:long clear_before_date:int source_perm_auth_key_id:long source_operation_id:string outbox_id:long = Bool;
func (m *defaultDialogClient) DialogClearDraftAfterSend(ctx context.Context, in *dialog.TLDialogClearDraftAfterSend) (*tg.Bool, error) {
	return m.rpc.DialogClearDraftAfterSend(ctx, in)
}

// DialogGetAllDrafts
// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (m *defaultDialogClient) DialogGetAllDrafts(ctx context.Context, in *dialog.TLDialogGetAllDrafts) (*dialog.VectorPeerWithDraftMessage, error) {
	return m.rpc.DialogGetAllDrafts(ctx, in)
}

// DialogClearAllDrafts
// dialog.clearAllDrafts user_id:long source_perm_auth_key_id:long operation_id:string outbox_ids:Vector<long> = Vector<PeerWithDraftMessage>;
func (m *defaultDialogClient) DialogClearAllDrafts(ctx context.Context, in *dialog.TLDialogClearAllDrafts) (*dialog.VectorPeerWithDraftMessage, error) {
	return m.rpc.DialogClearAllDrafts(ctx, in)
}

// DialogMarkDialogUnread
// dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;
func (m *defaultDialogClient) DialogMarkDialogUnread(ctx context.Context, in *dialog.TLDialogMarkDialogUnread) (*tg.Bool, error) {
	return m.rpc.DialogMarkDialogUnread(ctx, in)
}

// DialogToggleDialogPin
// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool source_perm_auth_key_id:long operation_id:string outbox_id:long = Int32;
func (m *defaultDialogClient) DialogToggleDialogPin(ctx context.Context, in *dialog.TLDialogToggleDialogPin) (*tg.Int32, error) {
	return m.rpc.DialogToggleDialogPin(ctx, in)
}

// DialogGetDialogUnreadMarkList
// dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;
func (m *defaultDialogClient) DialogGetDialogUnreadMarkList(ctx context.Context, in *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.VectorDialogPeer, error) {
	return m.rpc.DialogGetDialogUnreadMarkList(ctx, in)
}

// DialogReorderPinnedDialogs
// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> source_perm_auth_key_id:long operation_id:string outbox_id:long = Bool;
func (m *defaultDialogClient) DialogReorderPinnedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedDialogs) (*tg.Bool, error) {
	return m.rpc.DialogReorderPinnedDialogs(ctx, in)
}

// DialogGetTopMessage
// dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultDialogClient) DialogGetTopMessage(ctx context.Context, in *dialog.TLDialogGetTopMessage) (*tg.Int32, error) {
	return m.rpc.DialogGetTopMessage(ctx, in)
}

// DialogInsertOrUpdateDialog
// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true date2:flags.5?long pinned_msg_id:flags.6?int = Bool;
func (m *defaultDialogClient) DialogInsertOrUpdateDialog(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialog) (*tg.Bool, error) {
	return m.rpc.DialogInsertOrUpdateDialog(ctx, in)
}

// DialogDeleteDialog
// dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultDialogClient) DialogDeleteDialog(ctx context.Context, in *dialog.TLDialogDeleteDialog) (*tg.Bool, error) {
	return m.rpc.DialogDeleteDialog(ctx, in)
}

// DialogGetUserPinnedMessage
// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultDialogClient) DialogGetUserPinnedMessage(ctx context.Context, in *dialog.TLDialogGetUserPinnedMessage) (*tg.Int32, error) {
	return m.rpc.DialogGetUserPinnedMessage(ctx, in)
}

// DialogUpdateUserPinnedMessage
// dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;
func (m *defaultDialogClient) DialogUpdateUserPinnedMessage(ctx context.Context, in *dialog.TLDialogUpdateUserPinnedMessage) (*tg.Bool, error) {
	return m.rpc.DialogUpdateUserPinnedMessage(ctx, in)
}

// DialogInsertOrUpdateDialogFilter
// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
func (m *defaultDialogClient) DialogInsertOrUpdateDialogFilter(ctx context.Context, in *dialog.TLDialogInsertOrUpdateDialogFilter) (*tg.Bool, error) {
	return m.rpc.DialogInsertOrUpdateDialogFilter(ctx, in)
}

// DialogDeleteDialogFilter
// dialog.deleteDialogFilter user_id:long id:int = Bool;
func (m *defaultDialogClient) DialogDeleteDialogFilter(ctx context.Context, in *dialog.TLDialogDeleteDialogFilter) (*tg.Bool, error) {
	return m.rpc.DialogDeleteDialogFilter(ctx, in)
}

// DialogUpdateDialogFiltersOrder
// dialog.updateDialogFiltersOrder user_id:long order:Vector<int> = Bool;
func (m *defaultDialogClient) DialogUpdateDialogFiltersOrder(ctx context.Context, in *dialog.TLDialogUpdateDialogFiltersOrder) (*tg.Bool, error) {
	return m.rpc.DialogUpdateDialogFiltersOrder(ctx, in)
}

// DialogGetDialogFilters
// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;
func (m *defaultDialogClient) DialogGetDialogFilters(ctx context.Context, in *dialog.TLDialogGetDialogFilters) (*dialog.VectorDialogFilterExt, error) {
	return m.rpc.DialogGetDialogFilters(ctx, in)
}

// DialogGetDialogFolder
// dialog.getDialogFolder user_id:long folder_id:int = Vector<DialogExt>;
func (m *defaultDialogClient) DialogGetDialogFolder(ctx context.Context, in *dialog.TLDialogGetDialogFolder) (*dialog.VectorDialogExt, error) {
	return m.rpc.DialogGetDialogFolder(ctx, in)
}

// DialogEditPeerFolders
// dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int source_perm_auth_key_id:long operation_id:string outbox_ids:Vector<long> = Vector<DialogPinnedExt>;
func (m *defaultDialogClient) DialogEditPeerFolders(ctx context.Context, in *dialog.TLDialogEditPeerFolders) (*dialog.VectorDialogPinnedExt, error) {
	return m.rpc.DialogEditPeerFolders(ctx, in)
}

// DialogGetDialogsV2
// dialog.getDialogsV2 user_id:long cursor:DialogCursor exclude_pinned:Bool limit:int = DialogPage;
func (m *defaultDialogClient) DialogGetDialogsV2(ctx context.Context, in *dialog.TLDialogGetDialogsV2) (*dialog.DialogPage, error) {
	return m.rpc.DialogGetDialogsV2(ctx, in)
}

// DialogGetPeerDialogsV2
// dialog.getPeerDialogsV2 user_id:long peers:Vector<DialogPeer> = Vector<DialogExtV2>;
func (m *defaultDialogClient) DialogGetPeerDialogsV2(ctx context.Context, in *dialog.TLDialogGetPeerDialogsV2) (*dialog.VectorDialogExtV2, error) {
	return m.rpc.DialogGetPeerDialogsV2(ctx, in)
}

// DialogGetPinnedDialogsV2
// dialog.getPinnedDialogsV2 user_id:long folder_id:int limit:int = Vector<DialogExtV2>;
func (m *defaultDialogClient) DialogGetPinnedDialogsV2(ctx context.Context, in *dialog.TLDialogGetPinnedDialogsV2) (*dialog.VectorDialogExtV2, error) {
	return m.rpc.DialogGetPinnedDialogsV2(ctx, in)
}

// DialogGetDialogByPeerV2
// dialog.getDialogByPeerV2 user_id:long peer:DialogPeer = DialogExtV2;
func (m *defaultDialogClient) DialogGetDialogByPeerV2(ctx context.Context, in *dialog.TLDialogGetDialogByPeerV2) (*dialog.DialogExtV2, error) {
	return m.rpc.DialogGetDialogByPeerV2(ctx, in)
}

// DialogBatchGetDialogExtras
// dialog.batchGetDialogExtras user_id:long peers:Vector<DialogPeer> = Vector<DialogExtras>;
func (m *defaultDialogClient) DialogBatchGetDialogExtras(ctx context.Context, in *dialog.TLDialogBatchGetDialogExtras) (*dialog.VectorDialogExtras, error) {
	return m.rpc.DialogBatchGetDialogExtras(ctx, in)
}

// DialogGetChannelMessageReadParticipants
// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;
func (m *defaultDialogClient) DialogGetChannelMessageReadParticipants(ctx context.Context, in *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.VectorLong, error) {
	return m.rpc.DialogGetChannelMessageReadParticipants(ctx, in)
}

// DialogSetChatTheme
// dialog.setChatTheme user_id:long peer_type:int peer_id:long theme_emoticon:string = Bool;
func (m *defaultDialogClient) DialogSetChatTheme(ctx context.Context, in *dialog.TLDialogSetChatTheme) (*tg.Bool, error) {
	return m.rpc.DialogSetChatTheme(ctx, in)
}

// DialogSetHistoryTTL
// dialog.setHistoryTTL user_id:long peer_type:int peer_id:long ttl_period:int = Bool;
func (m *defaultDialogClient) DialogSetHistoryTTL(ctx context.Context, in *dialog.TLDialogSetHistoryTTL) (*tg.Bool, error) {
	return m.rpc.DialogSetHistoryTTL(ctx, in)
}

// DialogGetSavedDialogs
// dialog.getSavedDialogs user_id:long exclude_pinned:Bool offset_date:int offset_id:int offset_peer:PeerUtil limit:int = SavedDialogList;
func (m *defaultDialogClient) DialogGetSavedDialogs(ctx context.Context, in *dialog.TLDialogGetSavedDialogs) (*dialog.SavedDialogList, error) {
	return m.rpc.DialogGetSavedDialogs(ctx, in)
}

// DialogGetPinnedSavedDialogs
// dialog.getPinnedSavedDialogs user_id:long = SavedDialogList;
func (m *defaultDialogClient) DialogGetPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogGetPinnedSavedDialogs) (*dialog.SavedDialogList, error) {
	return m.rpc.DialogGetPinnedSavedDialogs(ctx, in)
}

// DialogUpsertSavedDialogFromMessage
// dialog.upsertSavedDialogFromMessage user_id:long peer_type:int peer_id:long top_peer_seq:long top_canonical_message_id:long top_message_date:int payload:bytes = Bool;
func (m *defaultDialogClient) DialogUpsertSavedDialogFromMessage(ctx context.Context, in *dialog.TLDialogUpsertSavedDialogFromMessage) (*tg.Bool, error) {
	return m.rpc.DialogUpsertSavedDialogFromMessage(ctx, in)
}

// DialogToggleSavedDialogPin
// dialog.toggleSavedDialogPin user_id:long peer:PeerUtil pinned:Bool = Bool;
func (m *defaultDialogClient) DialogToggleSavedDialogPin(ctx context.Context, in *dialog.TLDialogToggleSavedDialogPin) (*tg.Bool, error) {
	return m.rpc.DialogToggleSavedDialogPin(ctx, in)
}

// DialogReorderPinnedSavedDialogs
// dialog.reorderPinnedSavedDialogs user_id:long force:Bool order:Vector<PeerUtil> = Bool;
func (m *defaultDialogClient) DialogReorderPinnedSavedDialogs(ctx context.Context, in *dialog.TLDialogReorderPinnedSavedDialogs) (*tg.Bool, error) {
	return m.rpc.DialogReorderPinnedSavedDialogs(ctx, in)
}

// DialogGetDialogFilter
// dialog.getDialogFilter user_id:long id:int = DialogFilterExt;
func (m *defaultDialogClient) DialogGetDialogFilter(ctx context.Context, in *dialog.TLDialogGetDialogFilter) (*dialog.DialogFilterExt, error) {
	return m.rpc.DialogGetDialogFilter(ctx, in)
}

// DialogGetDialogFilterBySlug
// dialog.getDialogFilterBySlug user_id:long slug:string = DialogFilterExt;
func (m *defaultDialogClient) DialogGetDialogFilterBySlug(ctx context.Context, in *dialog.TLDialogGetDialogFilterBySlug) (*dialog.DialogFilterExt, error) {
	return m.rpc.DialogGetDialogFilterBySlug(ctx, in)
}

// DialogCreateDialogFilter
// dialog.createDialogFilter user_id:long dialog_filter:DialogFilterExt = DialogFilterExt;
func (m *defaultDialogClient) DialogCreateDialogFilter(ctx context.Context, in *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error) {
	return m.rpc.DialogCreateDialogFilter(ctx, in)
}

// DialogUpdateUnreadCount
// dialog.updateUnreadCount flags:# user_id:long peer_type:int peer_id:long unread_count:flags.0?int unread_mentions_count:flags.1?int unread_reactions_count:flags.2?int = Bool;
func (m *defaultDialogClient) DialogUpdateUnreadCount(ctx context.Context, in *dialog.TLDialogUpdateUnreadCount) (*tg.Bool, error) {
	return m.rpc.DialogUpdateUnreadCount(ctx, in)
}

// DialogToggleDialogFilterTags
// dialog.toggleDialogFilterTags user_id:long enabled:Bool = Bool;
func (m *defaultDialogClient) DialogToggleDialogFilterTags(ctx context.Context, in *dialog.TLDialogToggleDialogFilterTags) (*tg.Bool, error) {
	return m.rpc.DialogToggleDialogFilterTags(ctx, in)
}

// DialogGetDialogFilterTags
// dialog.getDialogFilterTags user_id:long = Bool;
func (m *defaultDialogClient) DialogGetDialogFilterTags(ctx context.Context, in *dialog.TLDialogGetDialogFilterTags) (*tg.Bool, error) {
	return m.rpc.DialogGetDialogFilterTags(ctx, in)
}

// DialogSetChatWallpaper
// dialog.setChatWallpaper flags:# user_id:long peer_type:int peer_id:long wallpaper_id:long wallpaper_overridden:flags.0?true = Bool;
func (m *defaultDialogClient) DialogSetChatWallpaper(ctx context.Context, in *dialog.TLDialogSetChatWallpaper) (*tg.Bool, error) {
	return m.rpc.DialogSetChatWallpaper(ctx, in)
}
