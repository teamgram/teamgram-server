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
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/internal/core"
)

// InboxEditUserMessageToInbox
// inbox.editUserMessageToInbox from_id:long peer_user_id:long message:Message = Void;
func (s *Service) InboxEditUserMessageToInbox(ctx context.Context, request *inbox.TLInboxEditUserMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.editUserMessageToInbox - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxEditUserMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.editUserMessageToInbox - reply: %s", r)
	return r, err
}

// InboxEditChatMessageToInbox
// inbox.editChatMessageToInbox from_id:long peer_chat_id:long message:Message = Void;
func (s *Service) InboxEditChatMessageToInbox(ctx context.Context, request *inbox.TLInboxEditChatMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.editChatMessageToInbox - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxEditChatMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.editChatMessageToInbox - reply: %s", r)
	return r, err
}

// InboxDeleteMessagesToInbox
// inbox.deleteMessagesToInbox from_id:long peer_type:int peer_id:long id:Vector<long> = Void;
func (s *Service) InboxDeleteMessagesToInbox(ctx context.Context, request *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.deleteMessagesToInbox - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxDeleteMessagesToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.deleteMessagesToInbox - reply: %s", r)
	return r, err
}

// InboxDeleteUserHistoryToInbox
// inbox.deleteUserHistoryToInbox flags:# from_id:long peer_user_id:long just_clear:flags.1?true max_id:int = Void;
func (s *Service) InboxDeleteUserHistoryToInbox(ctx context.Context, request *inbox.TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.deleteUserHistoryToInbox - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxDeleteUserHistoryToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.deleteUserHistoryToInbox - reply: %s", r)
	return r, err
}

// InboxDeleteChatHistoryToInbox
// inbox.deleteChatHistoryToInbox from_id:long peer_chat_id:long max_id:int = Void;
func (s *Service) InboxDeleteChatHistoryToInbox(ctx context.Context, request *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.deleteChatHistoryToInbox - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxDeleteChatHistoryToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.deleteChatHistoryToInbox - reply: %s", r)
	return r, err
}

// InboxReadUserMediaUnreadToInbox
// inbox.readUserMediaUnreadToInbox from_id:long peer_user_id:long id:Vector<InboxMessageId> = Void;
func (s *Service) InboxReadUserMediaUnreadToInbox(ctx context.Context, request *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.readUserMediaUnreadToInbox - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxReadUserMediaUnreadToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.readUserMediaUnreadToInbox - reply: %s", r)
	return r, err
}

// InboxReadChatMediaUnreadToInbox
// inbox.readChatMediaUnreadToInbox from_id:long peer_chat_id:long id:Vector<InboxMessageId> = Void;
func (s *Service) InboxReadChatMediaUnreadToInbox(ctx context.Context, request *inbox.TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.readChatMediaUnreadToInbox - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxReadChatMediaUnreadToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.readChatMediaUnreadToInbox - reply: %s", r)
	return r, err
}

// InboxUpdateHistoryReaded
// inbox.updateHistoryReaded from_id:long peer_type:int peer_id:long max_id:int sender:long = Void;
func (s *Service) InboxUpdateHistoryReaded(ctx context.Context, request *inbox.TLInboxUpdateHistoryReaded) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.updateHistoryReaded - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxUpdateHistoryReaded(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.updateHistoryReaded - reply: %s", r)
	return r, err
}

// InboxUpdatePinnedMessage
// inbox.updatePinnedMessage flags:# user_id:long unpin:flags.1?true peer_type:int peer_id:long id:int dialog_message_id:long = Void;
func (s *Service) InboxUpdatePinnedMessage(ctx context.Context, request *inbox.TLInboxUpdatePinnedMessage) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.updatePinnedMessage - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxUpdatePinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.updatePinnedMessage - reply: %s", r)
	return r, err
}

// InboxUnpinAllMessages
// inbox.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = Void;
func (s *Service) InboxUnpinAllMessages(ctx context.Context, request *inbox.TLInboxUnpinAllMessages) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.unpinAllMessages - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxUnpinAllMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.unpinAllMessages - reply: %s", r)
	return r, err
}

// InboxSendUserMessageToInboxV2
// inbox.sendUserMessageToInboxV2 flags:# user_id:long out:flags.0?true from_id:long peer_user_id:long inbox:MessageBox users:flags.1?Vector<ImmutableUser> = Void;
func (s *Service) InboxSendUserMessageToInboxV2(ctx context.Context, request *inbox.TLInboxSendUserMessageToInboxV2) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("inbox.sendUserMessageToInboxV2 - metadata: %s, request: %s", c.MD, request)

	r, err := c.InboxSendUserMessageToInboxV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("inbox.sendUserMessageToInboxV2 - reply: %s", r)
	return r, err
}
