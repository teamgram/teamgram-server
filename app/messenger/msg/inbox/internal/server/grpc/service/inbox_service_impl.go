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
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/internal/core"
)

// InboxSendUserMessageToInbox
// inbox.sendUserMessageToInbox from_id:long peer_user_id:long message:InboxMessageData = Void;
func (s *Service) InboxSendUserMessageToInbox(ctx context.Context, request *inbox.TLInboxSendUserMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.sendUserMessageToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxSendUserMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.sendUserMessageToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxSendChatMessageToInbox
// inbox.sendChatMessageToInbox from_id:long peer_chat_id:long message:InboxMessageData = Void;
func (s *Service) InboxSendChatMessageToInbox(ctx context.Context, request *inbox.TLInboxSendChatMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.sendChatMessageToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxSendChatMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.sendChatMessageToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxSendUserMultiMessageToInbox
// inbox.sendUserMultiMessageToInbox from_id:long peer_user_id:long message:Vector<InboxMessageData> = Void;
func (s *Service) InboxSendUserMultiMessageToInbox(ctx context.Context, request *inbox.TLInboxSendUserMultiMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.sendUserMultiMessageToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxSendUserMultiMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.sendUserMultiMessageToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxSendChatMultiMessageToInbox
// inbox.sendChatMultiMessageToInbox from_id:long peer_chat_id:long message:Vector<InboxMessageData> = Void;
func (s *Service) InboxSendChatMultiMessageToInbox(ctx context.Context, request *inbox.TLInboxSendChatMultiMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.sendChatMultiMessageToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxSendChatMultiMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.sendChatMultiMessageToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxEditUserMessageToInbox
// inbox.editUserMessageToInbox from_id:long peer_user_id:long message:Message = Void;
func (s *Service) InboxEditUserMessageToInbox(ctx context.Context, request *inbox.TLInboxEditUserMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.editUserMessageToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxEditUserMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.editUserMessageToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxEditChatMessageToInbox
// inbox.editChatMessageToInbox from_id:long peer_chat_id:long message:Message = Void;
func (s *Service) InboxEditChatMessageToInbox(ctx context.Context, request *inbox.TLInboxEditChatMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.editChatMessageToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxEditChatMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.editChatMessageToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxDeleteMessagesToInbox
// inbox.deleteMessagesToInbox from_id:long id:Vector<long> = Void;
func (s *Service) InboxDeleteMessagesToInbox(ctx context.Context, request *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.deleteMessagesToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxDeleteMessagesToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.deleteMessagesToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxDeleteUserHistoryToInbox
// inbox.deleteUserHistoryToInbox flags:# from_id:long peer_user_id:long just_clear:flags.1?true max_id:int = Void;
func (s *Service) InboxDeleteUserHistoryToInbox(ctx context.Context, request *inbox.TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.deleteUserHistoryToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxDeleteUserHistoryToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.deleteUserHistoryToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxDeleteChatHistoryToInbox
// inbox.deleteChatHistoryToInbox from_id:long peer_chat_id:long max_id:int = Void;
func (s *Service) InboxDeleteChatHistoryToInbox(ctx context.Context, request *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.deleteChatHistoryToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxDeleteChatHistoryToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.deleteChatHistoryToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxReadUserMediaUnreadToInbox
// inbox.readUserMediaUnreadToInbox from_id:long id:Vector<int> = Void;
func (s *Service) InboxReadUserMediaUnreadToInbox(ctx context.Context, request *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.readUserMediaUnreadToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxReadUserMediaUnreadToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.readUserMediaUnreadToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxReadChatMediaUnreadToInbox
// inbox.readChatMediaUnreadToInbox from_id:long peer_chat_id:long id:Vector<int> = Void;
func (s *Service) InboxReadChatMediaUnreadToInbox(ctx context.Context, request *inbox.TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.readChatMediaUnreadToInbox - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxReadChatMediaUnreadToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.readChatMediaUnreadToInbox - reply: %s", r.DebugString())
	return r, err
}

// InboxUpdateHistoryReaded
// inbox.updateHistoryReaded from_id:long peer_type:int peer_id:long max_id:int sender:long = Void;
func (s *Service) InboxUpdateHistoryReaded(ctx context.Context, request *inbox.TLInboxUpdateHistoryReaded) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.updateHistoryReaded - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxUpdateHistoryReaded(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.updateHistoryReaded - reply: %s", r.DebugString())
	return r, err
}

// InboxUpdatePinnedMessage
// inbox.updatePinnedMessage flags:# user_id:long unpin:flags.1?true peer_type:int peer_id:long id:int dialog_message_id:long = Void;
func (s *Service) InboxUpdatePinnedMessage(ctx context.Context, request *inbox.TLInboxUpdatePinnedMessage) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.updatePinnedMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxUpdatePinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.updatePinnedMessage - reply: %s", r.DebugString())
	return r, err
}

// InboxUnpinAllMessages
// inbox.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = Void;
func (s *Service) InboxUnpinAllMessages(ctx context.Context, request *inbox.TLInboxUnpinAllMessages) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("inbox.unpinAllMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.InboxUnpinAllMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("inbox.unpinAllMessages - reply: %s", r.DebugString())
	return r, err
}
