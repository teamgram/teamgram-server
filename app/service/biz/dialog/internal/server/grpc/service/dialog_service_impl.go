/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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
	c.Infof("dialog.saveDraftMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogSaveDraftMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.saveDraftMessage - reply: %s", r.DebugString())
	return r, err
}

// DialogClearDraftMessage
// dialog.clearDraftMessage user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) DialogClearDraftMessage(ctx context.Context, request *dialog.TLDialogClearDraftMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.clearDraftMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogClearDraftMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.clearDraftMessage - reply: %s", r.DebugString())
	return r, err
}

// DialogGetAllDrafts
// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (s *Service) DialogGetAllDrafts(ctx context.Context, request *dialog.TLDialogGetAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getAllDrafts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getAllDrafts - reply: %s", r.DebugString())
	return r, err
}

// DialogClearAllDrafts
// dialog.clearAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (s *Service) DialogClearAllDrafts(ctx context.Context, request *dialog.TLDialogClearAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.clearAllDrafts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogClearAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.clearAllDrafts - reply: %s", r.DebugString())
	return r, err
}

// DialogMarkDialogUnread
// dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;
func (s *Service) DialogMarkDialogUnread(ctx context.Context, request *dialog.TLDialogMarkDialogUnread) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.markDialogUnread - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogMarkDialogUnread(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.markDialogUnread - reply: %s", r.DebugString())
	return r, err
}

// DialogToggleDialogPin
// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;
func (s *Service) DialogToggleDialogPin(ctx context.Context, request *dialog.TLDialogToggleDialogPin) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.toggleDialogPin - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogToggleDialogPin(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.toggleDialogPin - reply: %s", r.DebugString())
	return r, err
}

// DialogGetDialogUnreadMarkList
// dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;
func (s *Service) DialogGetDialogUnreadMarkList(ctx context.Context, request *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.Vector_DialogPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getDialogUnreadMarkList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetDialogUnreadMarkList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getDialogUnreadMarkList - reply: %s", r.DebugString())
	return r, err
}

// DialogGetDialogsByOffsetDate
// dialog.getDialogsByOffsetDate user_id:long exclude_pinned:Bool offset_date:int limit:int = Vector<DialogExt>;
func (s *Service) DialogGetDialogsByOffsetDate(ctx context.Context, request *dialog.TLDialogGetDialogsByOffsetDate) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getDialogsByOffsetDate - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetDialogsByOffsetDate(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getDialogsByOffsetDate - reply: %s", r.DebugString())
	return r, err
}

// DialogGetDialogs
// dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;
func (s *Service) DialogGetDialogs(ctx context.Context, request *dialog.TLDialogGetDialogs) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getDialogs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getDialogs - reply: %s", r.DebugString())
	return r, err
}

// DialogGetDialogsByIdList
// dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;
func (s *Service) DialogGetDialogsByIdList(ctx context.Context, request *dialog.TLDialogGetDialogsByIdList) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getDialogsByIdList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetDialogsByIdList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getDialogsByIdList - reply: %s", r.DebugString())
	return r, err
}

// DialogGetDialogsCount
// dialog.getDialogsCount user_id:long exclude_pinned:Bool folder_id:int = Int32;
func (s *Service) DialogGetDialogsCount(ctx context.Context, request *dialog.TLDialogGetDialogsCount) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getDialogsCount - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetDialogsCount(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getDialogsCount - reply: %s", r.DebugString())
	return r, err
}

// DialogGetPinnedDialogs
// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
func (s *Service) DialogGetPinnedDialogs(ctx context.Context, request *dialog.TLDialogGetPinnedDialogs) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getPinnedDialogs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetPinnedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getPinnedDialogs - reply: %s", r.DebugString())
	return r, err
}

// DialogReorderPinnedDialogs
// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> = Bool;
func (s *Service) DialogReorderPinnedDialogs(ctx context.Context, request *dialog.TLDialogReorderPinnedDialogs) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.reorderPinnedDialogs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogReorderPinnedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.reorderPinnedDialogs - reply: %s", r.DebugString())
	return r, err
}

// DialogGetDialogById
// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
func (s *Service) DialogGetDialogById(ctx context.Context, request *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getDialogById - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetDialogById(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getDialogById - reply: %s", r.DebugString())
	return r, err
}

// DialogGetTopMessage
// dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;
func (s *Service) DialogGetTopMessage(ctx context.Context, request *dialog.TLDialogGetTopMessage) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getTopMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetTopMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getTopMessage - reply: %s", r.DebugString())
	return r, err
}

// DialogUpdateReadInbox
// dialog.updateReadInbox user_id:long peer_type:int peer_id:long read_inbox_id:int = Bool;
func (s *Service) DialogUpdateReadInbox(ctx context.Context, request *dialog.TLDialogUpdateReadInbox) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.updateReadInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogUpdateReadInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.updateReadInbox - reply: %s", r.DebugString())
	return r, err
}

// DialogUpdateReadOutbox
// dialog.updateReadOutbox user_id:long peer_type:int peer_id:long read_outbox_id:int = Bool;
func (s *Service) DialogUpdateReadOutbox(ctx context.Context, request *dialog.TLDialogUpdateReadOutbox) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.updateReadOutbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogUpdateReadOutbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.updateReadOutbox - reply: %s", r.DebugString())
	return r, err
}

// DialogInsertOrUpdateDialog
// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true = Bool;
func (s *Service) DialogInsertOrUpdateDialog(ctx context.Context, request *dialog.TLDialogInsertOrUpdateDialog) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.insertOrUpdateDialog - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogInsertOrUpdateDialog(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.insertOrUpdateDialog - reply: %s", r.DebugString())
	return r, err
}

// DialogDeleteDialog
// dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) DialogDeleteDialog(ctx context.Context, request *dialog.TLDialogDeleteDialog) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.deleteDialog - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogDeleteDialog(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.deleteDialog - reply: %s", r.DebugString())
	return r, err
}

// DialogGetUserPinnedMessage
// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
func (s *Service) DialogGetUserPinnedMessage(ctx context.Context, request *dialog.TLDialogGetUserPinnedMessage) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getUserPinnedMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetUserPinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getUserPinnedMessage - reply: %s", r.DebugString())
	return r, err
}

// DialogUpdateUserPinnedMessage
// dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;
func (s *Service) DialogUpdateUserPinnedMessage(ctx context.Context, request *dialog.TLDialogUpdateUserPinnedMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.updateUserPinnedMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogUpdateUserPinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.updateUserPinnedMessage - reply: %s", r.DebugString())
	return r, err
}

// DialogInsertOrUpdateDialogFilter
// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
func (s *Service) DialogInsertOrUpdateDialogFilter(ctx context.Context, request *dialog.TLDialogInsertOrUpdateDialogFilter) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.insertOrUpdateDialogFilter - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogInsertOrUpdateDialogFilter(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.insertOrUpdateDialogFilter - reply: %s", r.DebugString())
	return r, err
}

// DialogDeleteDialogFilter
// dialog.deleteDialogFilter user_id:long id:int = Bool;
func (s *Service) DialogDeleteDialogFilter(ctx context.Context, request *dialog.TLDialogDeleteDialogFilter) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.deleteDialogFilter - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogDeleteDialogFilter(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.deleteDialogFilter - reply: %s", r.DebugString())
	return r, err
}

// DialogUpdateDialogFiltersOrder
// dialog.updateDialogFiltersOrder user_id:long order:Vector<int> = Bool;
func (s *Service) DialogUpdateDialogFiltersOrder(ctx context.Context, request *dialog.TLDialogUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.updateDialogFiltersOrder - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogUpdateDialogFiltersOrder(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.updateDialogFiltersOrder - reply: %s", r.DebugString())
	return r, err
}

// DialogGetDialogFilters
// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;
func (s *Service) DialogGetDialogFilters(ctx context.Context, request *dialog.TLDialogGetDialogFilters) (*dialog.Vector_DialogFilterExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getDialogFilters - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetDialogFilters(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getDialogFilters - reply: %s", r.DebugString())
	return r, err
}

// DialogGetDialogFolder
// dialog.getDialogFolder user_id:long folder_id:int = Vector<DialogExt>;
func (s *Service) DialogGetDialogFolder(ctx context.Context, request *dialog.TLDialogGetDialogFolder) (*dialog.Vector_DialogExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getDialogFolder - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetDialogFolder(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getDialogFolder - reply: %s", r.DebugString())
	return r, err
}

// DialogEditPeerFolders
// dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int = Vector<DialogPinnedExt>;
func (s *Service) DialogEditPeerFolders(ctx context.Context, request *dialog.TLDialogEditPeerFolders) (*dialog.Vector_DialogPinnedExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.editPeerFolders - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogEditPeerFolders(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.editPeerFolders - reply: %s", r.DebugString())
	return r, err
}

// DialogGetChannelMessageReadParticipants
// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;
func (s *Service) DialogGetChannelMessageReadParticipants(ctx context.Context, request *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("dialog.getChannelMessageReadParticipants - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.DialogGetChannelMessageReadParticipants(request)
	if err != nil {
		return nil, err
	}

	c.Infof("dialog.getChannelMessageReadParticipants - reply: %s", r.DebugString())
	return r, err
}
