/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/core"
)

// DialogSaveDraftMessage
// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage = Bool;
func (s *Service) DialogSaveDraftMessage(ctx context.Context, request *dialog.TLDialogSaveDraftMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.saveDraftMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogSaveDraftMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.saveDraftMessage - reply: {%s}", r)
	return r, err
}

// DialogClearDraftMessage
// dialog.clearDraftMessage user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) DialogClearDraftMessage(ctx context.Context, request *dialog.TLDialogClearDraftMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.clearDraftMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogClearDraftMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.clearDraftMessage - reply: {%s}", r)
	return r, err
}

// DialogGetAllDrafts
// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (s *Service) DialogGetAllDrafts(ctx context.Context, request *dialog.TLDialogGetAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getAllDrafts - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getAllDrafts - reply: {%s}", r)
	return r, err
}

// DialogClearAllDrafts
// dialog.clearAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (s *Service) DialogClearAllDrafts(ctx context.Context, request *dialog.TLDialogClearAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.clearAllDrafts - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogClearAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.clearAllDrafts - reply: {%s}", r)
	return r, err
}

// DialogMarkDialogUnread
// dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;
func (s *Service) DialogMarkDialogUnread(ctx context.Context, request *dialog.TLDialogMarkDialogUnread) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.markDialogUnread - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogMarkDialogUnread(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.markDialogUnread - reply: {%s}", r)
	return r, err
}

// DialogToggleDialogPin
// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;
func (s *Service) DialogToggleDialogPin(ctx context.Context, request *dialog.TLDialogToggleDialogPin) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.toggleDialogPin - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogToggleDialogPin(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.toggleDialogPin - reply: {%s}", r)
	return r, err
}

// DialogGetDialogUnreadMarkList
// dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;
func (s *Service) DialogGetDialogUnreadMarkList(ctx context.Context, request *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.Vector_DialogPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogUnreadMarkList - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogUnreadMarkList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogUnreadMarkList - reply: {%s}", r)
	return r, err
}

// DialogGetDialogsByOffsetDate
// dialog.getDialogsByOffsetDate user_id:long exclude_pinned:Bool offset_date:int limit:int = Vector<DialogExt>;
func (s *Service) DialogGetDialogsByOffsetDate(ctx context.Context, request *dialog.TLDialogGetDialogsByOffsetDate) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogsByOffsetDate - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogsByOffsetDate(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogsByOffsetDate - reply: {%s}", r)
	return r, err
}

// DialogGetDialogs
// dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;
func (s *Service) DialogGetDialogs(ctx context.Context, request *dialog.TLDialogGetDialogs) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogs - reply: {%s}", r)
	return r, err
}

// DialogGetDialogsByIdList
// dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;
func (s *Service) DialogGetDialogsByIdList(ctx context.Context, request *dialog.TLDialogGetDialogsByIdList) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogsByIdList - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogsByIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogsByIdList - reply: {%s}", r)
	return r, err
}

// DialogGetDialogsCount
// dialog.getDialogsCount user_id:long exclude_pinned:Bool folder_id:int = Int32;
func (s *Service) DialogGetDialogsCount(ctx context.Context, request *dialog.TLDialogGetDialogsCount) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogsCount - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogsCount(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogsCount - reply: {%s}", r)
	return r, err
}

// DialogGetPinnedDialogs
// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
func (s *Service) DialogGetPinnedDialogs(ctx context.Context, request *dialog.TLDialogGetPinnedDialogs) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getPinnedDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetPinnedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getPinnedDialogs - reply: {%s}", r)
	return r, err
}

// DialogReorderPinnedDialogs
// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> = Bool;
func (s *Service) DialogReorderPinnedDialogs(ctx context.Context, request *dialog.TLDialogReorderPinnedDialogs) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.reorderPinnedDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogReorderPinnedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.reorderPinnedDialogs - reply: {%s}", r)
	return r, err
}

// DialogGetDialogById
// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
func (s *Service) DialogGetDialogById(ctx context.Context, request *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogById - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogById(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogById - reply: {%s}", r)
	return r, err
}

// DialogGetTopMessage
// dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;
func (s *Service) DialogGetTopMessage(ctx context.Context, request *dialog.TLDialogGetTopMessage) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getTopMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetTopMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getTopMessage - reply: {%s}", r)
	return r, err
}

// DialogUpdateReadInbox
// dialog.updateReadInbox user_id:long peer_type:int peer_id:long read_inbox_id:int = Bool;
func (s *Service) DialogUpdateReadInbox(ctx context.Context, request *dialog.TLDialogUpdateReadInbox) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.updateReadInbox - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogUpdateReadInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.updateReadInbox - reply: {%s}", r)
	return r, err
}

// DialogUpdateReadOutbox
// dialog.updateReadOutbox user_id:long peer_type:int peer_id:long read_outbox_id:int = Bool;
func (s *Service) DialogUpdateReadOutbox(ctx context.Context, request *dialog.TLDialogUpdateReadOutbox) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.updateReadOutbox - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogUpdateReadOutbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.updateReadOutbox - reply: {%s}", r)
	return r, err
}

// DialogInsertOrUpdateDialog
// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true date2:flags.5?long = Bool;
func (s *Service) DialogInsertOrUpdateDialog(ctx context.Context, request *dialog.TLDialogInsertOrUpdateDialog) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.insertOrUpdateDialog - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogInsertOrUpdateDialog(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.insertOrUpdateDialog - reply: {%s}", r)
	return r, err
}

// DialogDeleteDialog
// dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) DialogDeleteDialog(ctx context.Context, request *dialog.TLDialogDeleteDialog) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.deleteDialog - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogDeleteDialog(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.deleteDialog - reply: {%s}", r)
	return r, err
}

// DialogGetUserPinnedMessage
// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
func (s *Service) DialogGetUserPinnedMessage(ctx context.Context, request *dialog.TLDialogGetUserPinnedMessage) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getUserPinnedMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetUserPinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getUserPinnedMessage - reply: {%s}", r)
	return r, err
}

// DialogUpdateUserPinnedMessage
// dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;
func (s *Service) DialogUpdateUserPinnedMessage(ctx context.Context, request *dialog.TLDialogUpdateUserPinnedMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.updateUserPinnedMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogUpdateUserPinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.updateUserPinnedMessage - reply: {%s}", r)
	return r, err
}

// DialogInsertOrUpdateDialogFilter
// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
func (s *Service) DialogInsertOrUpdateDialogFilter(ctx context.Context, request *dialog.TLDialogInsertOrUpdateDialogFilter) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.insertOrUpdateDialogFilter - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogInsertOrUpdateDialogFilter(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.insertOrUpdateDialogFilter - reply: {%s}", r)
	return r, err
}

// DialogDeleteDialogFilter
// dialog.deleteDialogFilter user_id:long id:int = Bool;
func (s *Service) DialogDeleteDialogFilter(ctx context.Context, request *dialog.TLDialogDeleteDialogFilter) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.deleteDialogFilter - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogDeleteDialogFilter(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.deleteDialogFilter - reply: {%s}", r)
	return r, err
}

// DialogUpdateDialogFiltersOrder
// dialog.updateDialogFiltersOrder user_id:long order:Vector<int> = Bool;
func (s *Service) DialogUpdateDialogFiltersOrder(ctx context.Context, request *dialog.TLDialogUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.updateDialogFiltersOrder - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogUpdateDialogFiltersOrder(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.updateDialogFiltersOrder - reply: {%s}", r)
	return r, err
}

// DialogGetDialogFilters
// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;
func (s *Service) DialogGetDialogFilters(ctx context.Context, request *dialog.TLDialogGetDialogFilters) (*dialog.Vector_DialogFilterExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogFilters - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogFilters(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogFilters - reply: {%s}", r)
	return r, err
}

// DialogGetDialogFolder
// dialog.getDialogFolder user_id:long folder_id:int = Vector<DialogExt>;
func (s *Service) DialogGetDialogFolder(ctx context.Context, request *dialog.TLDialogGetDialogFolder) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogFolder - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogFolder(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogFolder - reply: {%s}", r)
	return r, err
}

// DialogEditPeerFolders
// dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int = Vector<DialogPinnedExt>;
func (s *Service) DialogEditPeerFolders(ctx context.Context, request *dialog.TLDialogEditPeerFolders) (*dialog.Vector_DialogPinnedExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.editPeerFolders - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogEditPeerFolders(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.editPeerFolders - reply: {%s}", r)
	return r, err
}

// DialogGetChannelMessageReadParticipants
// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;
func (s *Service) DialogGetChannelMessageReadParticipants(ctx context.Context, request *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getChannelMessageReadParticipants - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetChannelMessageReadParticipants(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getChannelMessageReadParticipants - reply: {%s}", r)
	return r, err
}

// DialogSetChatTheme
// dialog.setChatTheme user_id:long peer_type:int peer_id:long theme_emoticon:string = Bool;
func (s *Service) DialogSetChatTheme(ctx context.Context, request *dialog.TLDialogSetChatTheme) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.setChatTheme - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogSetChatTheme(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.setChatTheme - reply: {%s}", r)
	return r, err
}

// DialogSetHistoryTTL
// dialog.setHistoryTTL user_id:long peer_type:int peer_id:long ttl_period:int = Bool;
func (s *Service) DialogSetHistoryTTL(ctx context.Context, request *dialog.TLDialogSetHistoryTTL) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.setHistoryTTL - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogSetHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.setHistoryTTL - reply: {%s}", r)
	return r, err
}

// DialogGetMyDialogsData
// dialog.getMyDialogsData flags:# user_id:long user:flags.0?true chat:flags.1?true channel:flags.2?true = DialogsData;
func (s *Service) DialogGetMyDialogsData(ctx context.Context, request *dialog.TLDialogGetMyDialogsData) (*dialog.DialogsData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getMyDialogsData - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetMyDialogsData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getMyDialogsData - reply: {%s}", r)
	return r, err
}

// DialogGetSavedDialogs
// dialog.getSavedDialogs user_id:long exclude_pinned:Bool offset_date:int offset_id:int offset_peer:PeerUtil limit:int = SavedDialogList;
func (s *Service) DialogGetSavedDialogs(ctx context.Context, request *dialog.TLDialogGetSavedDialogs) (*dialog.SavedDialogList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getSavedDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetSavedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getSavedDialogs - reply: {%s}", r)
	return r, err
}

// DialogGetPinnedSavedDialogs
// dialog.getPinnedSavedDialogs user_id:long = SavedDialogList;
func (s *Service) DialogGetPinnedSavedDialogs(ctx context.Context, request *dialog.TLDialogGetPinnedSavedDialogs) (*dialog.SavedDialogList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getPinnedSavedDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetPinnedSavedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getPinnedSavedDialogs - reply: {%s}", r)
	return r, err
}

// DialogToggleSavedDialogPin
// dialog.toggleSavedDialogPin user_id:long peer:PeerUtil pinned:Bool = Bool;
func (s *Service) DialogToggleSavedDialogPin(ctx context.Context, request *dialog.TLDialogToggleSavedDialogPin) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.toggleSavedDialogPin - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogToggleSavedDialogPin(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.toggleSavedDialogPin - reply: {%s}", r)
	return r, err
}

// DialogReorderPinnedSavedDialogs
// dialog.reorderPinnedSavedDialogs user_id:long force:Bool order:Vector<PeerUtil> = Bool;
func (s *Service) DialogReorderPinnedSavedDialogs(ctx context.Context, request *dialog.TLDialogReorderPinnedSavedDialogs) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.reorderPinnedSavedDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogReorderPinnedSavedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.reorderPinnedSavedDialogs - reply: {%s}", r)
	return r, err
}

// DialogGetDialogFilter
// dialog.getDialogFilter user_id:long id:int = DialogFilterExt;
func (s *Service) DialogGetDialogFilter(ctx context.Context, request *dialog.TLDialogGetDialogFilter) (*dialog.DialogFilterExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogFilter - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogFilter(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogFilter - reply: {%s}", r)
	return r, err
}

// DialogGetDialogFilterBySlug
// dialog.getDialogFilterBySlug user_id:long slug:string = DialogFilterExt;
func (s *Service) DialogGetDialogFilterBySlug(ctx context.Context, request *dialog.TLDialogGetDialogFilterBySlug) (*dialog.DialogFilterExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.getDialogFilterBySlug - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogGetDialogFilterBySlug(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.getDialogFilterBySlug - reply: {%s}", r)
	return r, err
}

// DialogCreateDialogFilter
// dialog.createDialogFilter user_id:long dialog_filter:DialogFilterExt = DialogFilterExt;
func (s *Service) DialogCreateDialogFilter(ctx context.Context, request *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.createDialogFilter - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogCreateDialogFilter(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.createDialogFilter - reply: {%s}", r)
	return r, err
}

// DialogUpdateUnreadCount
// dialog.updateUnreadCount user_id:long peer_type:int peer_id:long unread_count:flags.0?int unread_mentions_count:flags.1?int unread_reactions_count:flags.2?int = Bool;
func (s *Service) DialogUpdateUnreadCount(ctx context.Context, request *dialog.TLDialogUpdateUnreadCount) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("dialog.updateUnreadCount - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.DialogUpdateUnreadCount(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("dialog.updateUnreadCount - reply: {%s}", r)
	return r, err
}
