/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/dialogs/internal/core"
)

// MessagesGetDialogs
// messages.getDialogs#a0f4cb4f flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.Dialogs;
func (s *Service) MessagesGetDialogs(ctx context.Context, request *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getDialogs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getDialogs - reply: %s", r.DebugString())
	return r, err
}

// MessagesSetTyping
// messages.setTyping#58943ee2 flags:# peer:InputPeer top_msg_id:flags.0?int action:SendMessageAction = Bool;
func (s *Service) MessagesSetTyping(ctx context.Context, request *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.setTyping - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSetTyping(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.setTyping - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetPeerSettingsEFD9A6A2
// messages.getPeerSettings#efd9a6a2 peer:InputPeer = messages.PeerSettings;
func (s *Service) MessagesGetPeerSettingsEFD9A6A2(ctx context.Context, request *mtproto.TLMessagesGetPeerSettingsEFD9A6A2) (*mtproto.Messages_PeerSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getPeerSettingsEFD9A6A2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetPeerSettingsEFD9A6A2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getPeerSettingsEFD9A6A2 - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetPeerDialogs
// messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
func (s *Service) MessagesGetPeerDialogs(ctx context.Context, request *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getPeerDialogs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetPeerDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getPeerDialogs - reply: %s", r.DebugString())
	return r, err
}

// MessagesToggleDialogPin
// messages.toggleDialogPin#a731e257 flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (s *Service) MessagesToggleDialogPin(ctx context.Context, request *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.toggleDialogPin - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesToggleDialogPin(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.toggleDialogPin - reply: %s", r.DebugString())
	return r, err
}

// MessagesReorderPinnedDialogs
// messages.reorderPinnedDialogs#3b1adf37 flags:# force:flags.0?true folder_id:int order:Vector<InputDialogPeer> = Bool;
func (s *Service) MessagesReorderPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.reorderPinnedDialogs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReorderPinnedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.reorderPinnedDialogs - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetPinnedDialogs
// messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
func (s *Service) MessagesGetPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getPinnedDialogs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetPinnedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getPinnedDialogs - reply: %s", r.DebugString())
	return r, err
}

// MessagesSendScreenshotNotification
// messages.sendScreenshotNotification#c97df020 peer:InputPeer reply_to_msg_id:int random_id:long = Updates;
func (s *Service) MessagesSendScreenshotNotification(ctx context.Context, request *mtproto.TLMessagesSendScreenshotNotification) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendScreenshotNotification - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendScreenshotNotification(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendScreenshotNotification - reply: %s", r.DebugString())
	return r, err
}

// MessagesMarkDialogUnread
// messages.markDialogUnread#c286d98f flags:# unread:flags.0?true peer:InputDialogPeer = Bool;
func (s *Service) MessagesMarkDialogUnread(ctx context.Context, request *mtproto.TLMessagesMarkDialogUnread) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.markDialogUnread - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesMarkDialogUnread(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.markDialogUnread - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetDialogUnreadMarks
// messages.getDialogUnreadMarks#22e24e22 = Vector<DialogPeer>;
func (s *Service) MessagesGetDialogUnreadMarks(ctx context.Context, request *mtproto.TLMessagesGetDialogUnreadMarks) (*mtproto.Vector_DialogPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getDialogUnreadMarks - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetDialogUnreadMarks(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getDialogUnreadMarks - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetOnlines
// messages.getOnlines#6e2be050 peer:InputPeer = ChatOnlines;
func (s *Service) MessagesGetOnlines(ctx context.Context, request *mtproto.TLMessagesGetOnlines) (*mtproto.ChatOnlines, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getOnlines - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetOnlines(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getOnlines - reply: %s", r.DebugString())
	return r, err
}

// MessagesHidePeerSettingsBar
// messages.hidePeerSettingsBar#4facb138 peer:InputPeer = Bool;
func (s *Service) MessagesHidePeerSettingsBar(ctx context.Context, request *mtproto.TLMessagesHidePeerSettingsBar) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.hidePeerSettingsBar - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesHidePeerSettingsBar(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.hidePeerSettingsBar - reply: %s", r.DebugString())
	return r, err
}

// MessagesSetHistoryTTL
// messages.setHistoryTTL#b80e5fe4 peer:InputPeer period:int = Updates;
func (s *Service) MessagesSetHistoryTTL(ctx context.Context, request *mtproto.TLMessagesSetHistoryTTL) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.setHistoryTTL - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSetHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.setHistoryTTL - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetPeerSettings3672E09C
// messages.getPeerSettings#3672e09c peer:InputPeer = PeerSettings;
func (s *Service) MessagesGetPeerSettings3672E09C(ctx context.Context, request *mtproto.TLMessagesGetPeerSettings3672E09C) (*mtproto.PeerSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getPeerSettings3672E09C - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetPeerSettings3672E09C(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getPeerSettings3672E09C - reply: %s", r.DebugString())
	return r, err
}
