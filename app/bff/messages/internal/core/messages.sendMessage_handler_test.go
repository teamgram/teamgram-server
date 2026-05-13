package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/svc"
	msgclient "github.com/teamgram/teamgram-server/v2/app/messenger/msg/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	idgenclient "github.com/teamgram/teamgram-server/v2/app/service/idgen/client"
	idgenpb "github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type messagesFakeMsgClient struct {
	msgclient.MsgClient
	sendMessage         func(ctx context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error)
	getUserMessage      func(ctx context.Context, in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error)
	getUserMessageList  func(ctx context.Context, in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error)
	getHistory          func(ctx context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error)
	readHistoryV2       func(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error)
	updatePinnedMessage func(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error)
	unpinAllMessages    func(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error)
	deleteMessages      func(ctx context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error)
	deleteHistory       func(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error)
	editMessage         func(ctx context.Context, in *msg.TLMsgEditMessage) (*tg.Updates, error)
	searchHashtag       func(ctx context.Context, in *msg.TLMsgSearchHashtag) (*tg.MessagesMessages, error)
}

func (f *messagesFakeMsgClient) MsgSendMessage(ctx context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
	return f.sendMessage(ctx, in)
}

func (f *messagesFakeMsgClient) MsgGetUserMessage(ctx context.Context, in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error) {
	return f.getUserMessage(ctx, in)
}

func (f *messagesFakeMsgClient) MsgGetUserMessageList(ctx context.Context, in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
	return f.getUserMessageList(ctx, in)
}

func (f *messagesFakeMsgClient) MsgGetHistory(ctx context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
	return f.getHistory(ctx, in)
}

func (f *messagesFakeMsgClient) MsgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
	return f.readHistoryV2(ctx, in)
}

func (f *messagesFakeMsgClient) MsgUpdatePinnedMessage(ctx context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error) {
	return f.updatePinnedMessage(ctx, in)
}

func (f *messagesFakeMsgClient) MsgUnpinAllMessages(ctx context.Context, in *msg.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error) {
	return f.unpinAllMessages(ctx, in)
}

func (f *messagesFakeMsgClient) MsgDeleteMessages(ctx context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	return f.deleteMessages(ctx, in)
}

func (f *messagesFakeMsgClient) MsgDeleteHistory(ctx context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error) {
	return f.deleteHistory(ctx, in)
}

func (f *messagesFakeMsgClient) MsgEditMessage(ctx context.Context, in *msg.TLMsgEditMessage) (*tg.Updates, error) {
	return f.editMessage(ctx, in)
}

func (f *messagesFakeMsgClient) MsgSearchHashtag(ctx context.Context, in *msg.TLMsgSearchHashtag) (*tg.MessagesMessages, error) {
	return f.searchHashtag(ctx, in)
}

type messagesFakeUserupdatesClient struct {
	userupdatesclient.UserupdatesClient
	getOutboxReadDate func(ctx context.Context, in *userupdates.TLUserupdatesGetOutboxReadDate) (*tg.OutboxReadDate, error)
}

func (f *messagesFakeUserupdatesClient) UserupdatesGetOutboxReadDate(ctx context.Context, in *userupdates.TLUserupdatesGetOutboxReadDate) (*tg.OutboxReadDate, error) {
	return f.getOutboxReadDate(ctx, in)
}

type messagesFakeIdgenClient struct {
	idgenclient.IdgenClient
	nextID func(ctx context.Context, in *idgenpb.TLIdgenNextId) (*tg.Int64, error)
	seq    int64
}

func (f *messagesFakeIdgenClient) IdgenNextId(ctx context.Context, in *idgenpb.TLIdgenNextId) (*tg.Int64, error) {
	if f.nextID != nil {
		return f.nextID(ctx, in)
	}
	f.seq++
	return tg.MakeTLInt64(&tg.TLInt64{V: 900000 + f.seq}), nil
}

type messagesFakeUserClient struct {
	userclient.UserClient
	projectUsers     func(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
	getUserIDByPhone func(ctx context.Context, in *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error)
}

func (f *messagesFakeUserClient) UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	return f.projectUsers(ctx, in)
}

func (f *messagesFakeUserClient) UserGetUserIdByPhone(ctx context.Context, in *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error) {
	if f.getUserIDByPhone == nil {
		return nil, nil
	}
	return f.getUserIDByPhone(ctx, in)
}

func newSendMsgCore(client msgclient.MsgClient, selfID, authKeyID int64) *MessagesCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{
			IdgenClient: &messagesFakeIdgenClient{},
			MsgClient:   client,
		},
	})
	c.MD = &metadata.RpcMetadata{
		UserId:        selfID,
		PermAuthKeyId: authKeyID,
	}
	return c
}

func newMessagesCoreWithRepo(repo *repository.Repository, selfID, authKeyID int64) *MessagesCore {
	if repo.IdgenClient == nil {
		repo.IdgenClient = &messagesFakeIdgenClient{}
	}
	c := New(context.Background(), &svc.ServiceContext{Repo: repo})
	c.MD = &metadata.RpcMetadata{
		UserId:        selfID,
		PermAuthKeyId: authKeyID,
	}
	return c
}

func testUpdates() *tg.Updates {
	return &tg.Updates{
		Clazz: tg.MakeTLUpdates(&tg.TLUpdates{
			Updates: []tg.UpdateClazz{},
			Users:   []tg.UserClazz{},
			Chats:   []tg.ChatClazz{},
			Date:    1000000,
			Seq:     0,
		}),
	}
}

func inputPeerUser(userID int64) *tg.TLInputPeerUser {
	return &tg.TLInputPeerUser{UserId: userID}
}

func inputPeerSelf() *tg.TLInputPeerSelf {
	return &tg.TLInputPeerSelf{}
}

func inputPeerChat(chatID int64) *tg.TLInputPeerChat {
	return &tg.TLInputPeerChat{ChatId: chatID}
}

// --- Success ---

func TestMessagesSendMessage_Success(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	r, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
	if got.UserId != 100 {
		t.Fatalf("UserId = %d, want 100", got.UserId)
	}
	if got.AuthKeyId != 200 {
		t.Fatalf("AuthKeyId = %d, want 200", got.AuthKeyId)
	}
	if got.PeerType != payload.PeerTypeUser {
		t.Fatalf("PeerType = %d, want %d", got.PeerType, payload.PeerTypeUser)
	}
	if got.PeerId != 300 {
		t.Fatalf("PeerId = %d, want 300", got.PeerId)
	}
	if len(got.Message) == 0 || got.Message[0] == nil {
		t.Fatal("OutboxMessage is nil or empty")
	}
	if r == nil {
		t.Fatal("result is nil")
	}
}

func TestMessagesSendMessageAllowsInputPeerChat(t *testing.T) {
	var got *msg.TLMsgSendMessage
	var checked *chatpb.TLChatCheckMessageAction
	core := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: &messagesFakeChatClient{
			checkMessageAction: func(_ context.Context, in *chatpb.TLChatCheckMessageAction) (*chatpb.MessageActionCheckResult, error) {
				checked = in
				return chatpb.MakeTLMessageActionCheckResult(&chatpb.TLMessageActionCheckResult{
					SelfId: in.SelfId, ChatId: in.ChatId, Action: in.Action, MediaKind: in.MediaKind,
				}).ToMessageActionCheckResult(), nil
			},
		},
		MsgClient: &messagesFakeMsgClient{sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		}},
	}, 1001, 9001)

	_, err := core.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerChat(55),
		ReplyTo:  tg.MakeTLInputReplyToMessage(&tg.TLInputReplyToMessage{ReplyToMsgId: 7}),
		Message:  "hello chat",
		RandomId: 12345,
	})
	if err != nil {
		t.Fatalf("MessagesSendMessage() error = %v", err)
	}
	if checked == nil || checked.ChatId != 55 || checked.SelfId != 1001 || checked.Action != chatpb.MessageActionSendText {
		t.Fatalf("chat check = %+v, want send_text for chat 55", checked)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 {
		t.Fatalf("msg request = %+v, want PeerTypeChat/chat 55", got)
	}
	message, ok := got.Message[0].Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message type = %T, want *tg.TLMessage", got.Message[0].Message)
	}
	if peer, ok := message.PeerId.(*tg.TLPeerChat); !ok || peer.ChatId != 55 {
		t.Fatalf("outbox peer = %#v, want peerChat 55", message.PeerId)
	}
	if reply, ok := message.ReplyTo.(*tg.TLMessageReplyHeader); !ok || reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 7 {
		t.Fatalf("reply header = %#v, want reply to chat message 7", message.ReplyTo)
	}
}

func TestSendMessageClearDraftCarriesSourcePermAuthKey(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		ClearDraft: true,
		Peer:       inputPeerUser(300),
		Message:    "hello",
		RandomId:   42,
	})
	if err != nil {
		t.Fatalf("MessagesSendMessage error = %v", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
	if !got.ClearDraft {
		t.Fatal("ClearDraft = false, want true")
	}
	if got.SourcePermAuthKeyId == nil || *got.SourcePermAuthKeyId != 200 {
		t.Fatalf("SourcePermAuthKeyId = %v, want 200", got.SourcePermAuthKeyId)
	}
	if got.ClearDraftBeforeDate == nil || *got.ClearDraftBeforeDate == 0 {
		t.Fatalf("ClearDraftBeforeDate = %v, want non-zero", got.ClearDraftBeforeDate)
	}
}

func TestMessagesSendMessage_InputReplyToMessageSetsReplyHeader(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "reply body",
		RandomId: 42,
		ReplyTo: tg.MakeTLInputReplyToMessage(&tg.TLInputReplyToMessage{
			ReplyToMsgId: 7,
		}),
	})
	if err != nil {
		t.Fatalf("MessagesSendMessage error = %v", err)
	}
	if got == nil || len(got.Message) != 1 || got.Message[0] == nil {
		t.Fatalf("msg request missing outbox: %+v", got)
	}
	outboxMessage, ok := got.Message[0].Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message type = %T, want *tg.TLMessage", got.Message[0].Message)
	}
	reply, ok := outboxMessage.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok {
		t.Fatalf("ReplyTo = %T, want *tg.TLMessageReplyHeader", outboxMessage.ReplyTo)
	}
	if reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 7 {
		t.Fatalf("ReplyToMsgId = %v, want 7", reply.ReplyToMsgId)
	}
}

func TestMessagesGetHistory_UserPeerSuccess(t *testing.T) {
	var got *msg.TLMsgGetHistory
	var gotProjection *userpb.TLUserGetUserProjectionBundle
	reply := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{
			tg.MakeTLMessage(&tg.TLMessage{
				Id:      5,
				FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 100}),
				PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 300}),
				Message: "hello",
			}),
		},
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{},
	}).ToMessagesMessages()
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{
			getHistory: func(_ context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
				got = in
				return reply, nil
			},
		},
		UserClient: &messagesFakeUserClient{
			projectUsers: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotProjection = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 100, Self: true}),
							tg.MakeTLUser(&tg.TLUser{Id: 300, Contact: true}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
			},
		},
	}, 100, 200)

	r, err := c.MessagesGetHistory(&tg.TLMessagesGetHistory{
		Peer:       inputPeerUser(300),
		OffsetId:   7,
		OffsetDate: 8,
		AddOffset:  9,
		Limit:      10,
		MaxId:      11,
		MinId:      12,
		Hash:       13,
	})
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 100 ||
		len(gotProjection.TargetUserIds) != 2 || gotProjection.TargetUserIds[0] != 100 || gotProjection.TargetUserIds[1] != 300 {
		t.Fatalf("projection request = %+v, want viewer [100] target [100 300]", gotProjection)
	}
	if got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != payload.PeerTypeUser || got.PeerId != 300 {
		t.Fatalf("unexpected service identity/peer: %+v", got)
	}
	if got.OffsetId != 7 || got.OffsetDate != 8 || got.AddOffset != 9 || got.Limit != 10 || got.MaxId != 11 || got.MinId != 12 || got.Hash != 13 {
		t.Fatalf("unexpected paging input: %+v", got)
	}
	messages, ok := r.ToMessagesMessages()
	if !ok || len(messages.Users) != 2 {
		t.Fatalf("history users = %#v, ok=%v, want projected users", r, ok)
	}
}

func TestMessagesGetHistoryAllowsInputPeerChat(t *testing.T) {
	var got *msg.TLMsgGetHistory
	var checked *chatpb.TLChatCheckChatAccess
	reply := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesMessages()
	c := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: &messagesFakeChatClient{
			checkChatAccess: func(_ context.Context, in *chatpb.TLChatCheckChatAccess) (*chatpb.ChatAccessCheckResult, error) {
				checked = in
				return chatpb.MakeTLChatAccessCheckResult(&chatpb.TLChatAccessCheckResult{
					SelfId: in.SelfId, ChatId: in.ChatId, AccessKind: in.AccessKind,
				}).ToChatAccessCheckResult(), nil
			},
		},
		MsgClient: &messagesFakeMsgClient{
			getHistory: func(_ context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
				got = in
				return reply, nil
			},
		},
	}, 1001, 9001)

	_, err := c.MessagesGetHistory(&tg.TLMessagesGetHistory{
		Peer:  inputPeerChat(55),
		Limit: 20,
	})
	if err != nil {
		t.Fatalf("MessagesGetHistory() error = %v", err)
	}
	if checked == nil || checked.ChatId != 55 || checked.SelfId != 1001 || checked.AccessKind != chatpb.ChatAccessGetHistory {
		t.Fatalf("chat access check = %+v, want get_history for chat 55", checked)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 || got.Limit != 20 {
		t.Fatalf("msg request = %+v, want PeerTypeChat/chat 55", got)
	}
}

func TestMessagesGetHistory_InputPeerSelfTargetsCurrentUser(t *testing.T) {
	var got *msg.TLMsgGetHistory
	reply := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesMessages()
	c := newSendMsgCore(&messagesFakeMsgClient{
		getHistory: func(_ context.Context, in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
			got = in
			return reply, nil
		},
	}, 100, 200)

	if _, err := c.MessagesGetHistory(&tg.TLMessagesGetHistory{
		Peer:  inputPeerSelf(),
		Limit: 30,
	}); err != nil {
		t.Fatalf("error = %v", err)
	}
	if got == nil || got.UserId != 100 || got.PeerId != 100 || got.Limit != 30 {
		t.Fatalf("unexpected history request: %+v", got)
	}
}

func TestMessagesSearchPinnedReturnsEmptyMessages(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{}, 100, 200)

	r, err := c.MessagesSearch(&tg.TLMessagesSearch{
		Peer:   inputPeerSelf(),
		Q:      "",
		Limit:  100,
		Filter: tg.MakeTLInputMessagesFilterPinned(&tg.TLInputMessagesFilterPinned{}),
	})
	if err != nil {
		t.Fatalf("MessagesSearch error = %v", err)
	}
	messages, ok := r.ToMessagesMessages()
	if !ok {
		t.Fatalf("MessagesSearch returned %s, want messages.messages", r.ClazzName())
	}
	if len(messages.Messages) != 0 || len(messages.Chats) != 0 || len(messages.Users) != 0 {
		t.Fatalf("MessagesSearch reply = %+v, want empty messages.messages", messages)
	}
}

func TestMessagesSearchHashtagRoutesToMsg(t *testing.T) {
	var got *msg.TLMsgSearchHashtag
	var gotProjection *userpb.TLUserGetUserProjectionBundle
	reply := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{
			tg.MakeTLMessage(&tg.TLMessage{
				Id:      7,
				FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 100}),
				PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 300}),
				Message: "#tag hello",
			}),
		},
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{},
	}).ToMessagesMessages()
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{
			searchHashtag: func(_ context.Context, in *msg.TLMsgSearchHashtag) (*tg.MessagesMessages, error) {
				got = in
				return reply, nil
			},
		},
		UserClient: &messagesFakeUserClient{
			projectUsers: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotProjection = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 100, Self: true}),
							tg.MakeTLUser(&tg.TLUser{Id: 300, Contact: true}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
			},
		},
	}, 100, 200)

	r, err := c.MessagesSearch(&tg.TLMessagesSearch{
		Peer:   inputPeerUser(300),
		Q:      "#tag",
		Limit:  20,
		Filter: tg.MakeTLInputMessagesFilterEmpty(&tg.TLInputMessagesFilterEmpty{}),
	})
	if err != nil {
		t.Fatalf("MessagesSearch error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if got == nil || got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != payload.PeerTypeUser || got.PeerId != 300 || got.HashTag != "tag" || got.Limit != 20 {
		t.Fatalf("MsgSearchHashtag request = %+v", got)
	}
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 100 ||
		len(gotProjection.TargetUserIds) != 2 || gotProjection.TargetUserIds[0] != 100 || gotProjection.TargetUserIds[1] != 300 {
		t.Fatalf("projection request = %+v, want viewer [100] target [100 300]", gotProjection)
	}
}

func TestMessagesSearchInputPeerChatUsesChatPeerForHashtag(t *testing.T) {
	var got *msg.TLMsgSearchHashtag
	var checked *chatpb.TLChatCheckChatAccess
	reply := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesMessages()
	c := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: &messagesFakeChatClient{
			checkChatAccess: func(_ context.Context, in *chatpb.TLChatCheckChatAccess) (*chatpb.ChatAccessCheckResult, error) {
				checked = in
				return chatpb.MakeTLChatAccessCheckResult(&chatpb.TLChatAccessCheckResult{
					SelfId: in.SelfId, ChatId: in.ChatId, AccessKind: in.AccessKind,
				}).ToChatAccessCheckResult(), nil
			},
		},
		MsgClient: &messagesFakeMsgClient{
			searchHashtag: func(_ context.Context, in *msg.TLMsgSearchHashtag) (*tg.MessagesMessages, error) {
				got = in
				return reply, nil
			},
		},
	}, 1001, 9001)

	r, err := c.MessagesSearch(&tg.TLMessagesSearch{
		Peer:   inputPeerChat(55),
		Q:      "#topic",
		Limit:  20,
		Filter: tg.MakeTLInputMessagesFilterEmpty(&tg.TLInputMessagesFilterEmpty{}),
	})
	if err != nil {
		t.Fatalf("MessagesSearch error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if checked == nil || checked.ChatId != 55 || checked.SelfId != 1001 || checked.AccessKind != chatpb.ChatAccessSearch {
		t.Fatalf("chat access check = %+v, want search for chat 55", checked)
	}
	if got == nil || got.UserId != 1001 || got.AuthKeyId != 9001 || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 || got.HashTag != "topic" || got.Limit != 20 {
		t.Fatalf("MsgSearchHashtag request = %+v, want PeerTypeChat/chat 55", got)
	}
}

func TestMessagesSearchEmptyQueryRejectedForEmptyFilter(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{}, 100, 200)

	_, err := c.MessagesSearch(&tg.TLMessagesSearch{
		Peer:   inputPeerSelf(),
		Filter: tg.MakeTLInputMessagesFilterEmpty(&tg.TLInputMessagesFilterEmpty{}),
	})
	if err != tg.ErrSearchQueryEmpty {
		t.Fatalf("MessagesSearch error = %v, want %v", err, tg.ErrSearchQueryEmpty)
	}
}

func TestMessagesGetOutboxReadDateRoutesToUserupdates(t *testing.T) {
	var got *userupdates.TLUserupdatesGetOutboxReadDate
	reply := tg.MakeTLOutboxReadDate(&tg.TLOutboxReadDate{Date: 123456}).ToOutboxReadDate()
	c := newMessagesCoreWithRepo(&repository.Repository{
		UserupdatesClient: &messagesFakeUserupdatesClient{
			getOutboxReadDate: func(_ context.Context, in *userupdates.TLUserupdatesGetOutboxReadDate) (*tg.OutboxReadDate, error) {
				got = in
				return reply, nil
			},
		},
	}, 100, 200)

	r, err := c.MessagesGetOutboxReadDate(&tg.TLMessagesGetOutboxReadDate{
		Peer:  inputPeerUser(300),
		MsgId: 7,
	})
	if err != nil {
		t.Fatalf("MessagesGetOutboxReadDate error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if got == nil || got.UserId != 100 || got.PeerType != payload.PeerTypeUser || got.PeerId != 300 || got.MsgId != 7 {
		t.Fatalf("UserupdatesGetOutboxReadDate request = %+v", got)
	}
}

func TestMessagesReadHistory_InputPeerSelfSuccess(t *testing.T) {
	var got *msg.TLMsgReadHistoryV2
	reply := tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      3,
		PtsCount: 0,
	}).ToMessagesAffectedMessages()
	c := newSendMsgCore(&messagesFakeMsgClient{
		readHistoryV2: func(_ context.Context, in *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
			got = in
			return reply, nil
		},
	}, 100, 200)

	r, err := c.MessagesReadHistory(&tg.TLMessagesReadHistory{
		Peer:  inputPeerSelf(),
		MaxId: 2,
	})
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if got == nil || got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != payload.PeerTypeUser || got.PeerId != 100 || got.MaxId != 2 {
		t.Fatalf("unexpected read history request: %+v", got)
	}
}

func TestMessagesReadHistory_UnsupportedPeerRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		readHistoryV2: func(context.Context, *msg.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesReadHistory(&tg.TLMessagesReadHistory{
		Peer:  tg.MakeTLInputPeerChannel(&tg.TLInputPeerChannel{ChannelId: 300}),
		MaxId: 2,
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesUpdatePinnedMessage_UserPeerSuccess(t *testing.T) {
	var got *msg.TLMsgUpdatePinnedMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		updatePinnedMessage: func(_ context.Context, in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	r, err := c.MessagesUpdatePinnedMessage(&tg.TLMessagesUpdatePinnedMessage{
		Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 300}),
		Id:   7,
	})
	if err != nil {
		t.Fatalf("MessagesUpdatePinnedMessage() error = %v", err)
	}
	if r == nil {
		t.Fatal("result is nil")
	}
	if got == nil || got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != payload.PeerTypeUser || got.PeerId != 300 || got.Id != 7 {
		t.Fatalf("MsgUpdatePinnedMessage request = %+v", got)
	}
}

func TestMessagesDeleteHistory_UserPeerSuccess(t *testing.T) {
	var got *msg.TLMsgDeleteHistory
	reply := tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{Pts: 9, PtsCount: 1}).ToMessagesAffectedHistory()
	c := newSendMsgCore(&messagesFakeMsgClient{
		deleteHistory: func(_ context.Context, in *msg.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error) {
			got = in
			return reply, nil
		},
	}, 100, 200)

	r, err := c.MessagesDeleteHistory(&tg.TLMessagesDeleteHistory{
		Peer:  tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 300}),
		MaxId: 7,
	})
	if err != nil {
		t.Fatalf("MessagesDeleteHistory() error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if got == nil || got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != payload.PeerTypeUser || got.PeerId != 300 || got.MaxId != 7 {
		t.Fatalf("MsgDeleteHistory request = %+v", got)
	}
}

func TestMessagesDeleteMessagesRoutesGlobalPublicIDs(t *testing.T) {
	var got *msg.TLMsgDeleteMessages
	reply := tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{Pts: 10, PtsCount: 1}).ToMessagesAffectedMessages()
	c := newSendMsgCore(&messagesFakeMsgClient{
		deleteMessages: func(_ context.Context, in *msg.TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error) {
			got = in
			return reply, nil
		},
	}, 100, 200)

	r, err := c.MessagesDeleteMessages(&tg.TLMessagesDeleteMessages{Revoke: true, Id: []int32{7, 8}})
	if err != nil {
		t.Fatalf("MessagesDeleteMessages() error = %v", err)
	}
	if r != reply {
		t.Fatalf("reply mismatch: got %p want %p", r, reply)
	}
	if got == nil || got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != 0 || got.PeerId != 0 || !got.Revoke || len(got.Id) != 2 {
		t.Fatalf("MsgDeleteMessages request = %+v", got)
	}
}

func TestMessagesEditMessage_TextUserPeerRoutesToMsg(t *testing.T) {
	var got *msg.TLMsgEditMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		editMessage: func(_ context.Context, in *msg.TLMsgEditMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)
	text := "edited"

	r, err := c.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:      inputPeerUser(300),
		Id:        7,
		Message:   &text,
		NoWebpage: true,
		Entities:  []tg.MessageEntityClazz{tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{})},
	})
	if err != nil {
		t.Fatalf("MessagesEditMessage error = %v", err)
	}
	if r == nil {
		t.Fatal("result is nil")
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
	if got.UserId != 100 || got.AuthKeyId != 200 || got.PeerType != payload.PeerTypeUser || got.PeerId != 300 {
		t.Fatalf("unexpected edit request identity/peer: %+v", got)
	}
	if got.DstMessage == nil || got.DstMessage.MessageId != 7 || got.DstMessage.SenderUserId != 100 {
		t.Fatalf("unexpected dst message: %+v", got.DstMessage)
	}
	if got.NewMessage == nil {
		t.Fatal("new outbox message is nil")
	}
	newMessage, ok := got.NewMessage.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("new message type = %T, want *tg.TLMessage", got.NewMessage.Message)
	}
	if newMessage.Message != text || newMessage.EditDate == nil || newMessage.EditHide {
		t.Fatalf("unexpected new message edit fields: %+v", newMessage)
	}
	if len(newMessage.Entities) != 1 {
		t.Fatalf("entities len = %d, want 1", len(newMessage.Entities))
	}
	if !got.NewMessage.NoWebpage {
		t.Fatal("NoWebpage = false, want true")
	}
}

func TestMessagesEditMessage_InputPeerSelfTargetsCurrentUser(t *testing.T) {
	var got *msg.TLMsgEditMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		editMessage: func(_ context.Context, in *msg.TLMsgEditMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)
	text := "edited"

	_, err := c.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:    inputPeerSelf(),
		Id:      7,
		Message: &text,
	})
	if err != nil {
		t.Fatalf("MessagesEditMessage error = %v", err)
	}
	if got == nil || got.PeerId != 100 {
		t.Fatalf("unexpected edit request: %+v", got)
	}
}

func TestMessagesEditMessage_EmptyTextRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		editMessage: func(_ context.Context, _ *msg.TLMsgEditMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)
	text := ""

	_, err := c.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:    inputPeerUser(300),
		Id:      7,
		Message: &text,
	})
	if err != tg.ErrMessageEmpty {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageEmpty)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesEditMessage_UnsupportedPeerRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		editMessage: func(_ context.Context, _ *msg.TLMsgEditMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)
	text := "edited"

	_, err := c.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:    tg.MakeTLInputPeerChannel(&tg.TLInputPeerChannel{ChannelId: 300}),
		Id:      7,
		Message: &text,
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesEditMessage_MediaRejectedUntilSupported(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		editMessage: func(_ context.Context, _ *msg.TLMsgEditMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)
	text := "edited"

	_, err := c.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:    inputPeerUser(300),
		Id:      7,
		Message: &text,
		Media:   tg.MakeTLInputMediaEmpty(&tg.TLInputMediaEmpty{}),
	})
	if err != tg.ErrMediaInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrMediaInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesEditMessage_MsgStateConflictMappedToMsgIdInvalid(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		editMessage: func(_ context.Context, _ *msg.TLMsgEditMessage) (*tg.Updates, error) {
			return nil, msg.ErrSendStateConflict
		},
	}, 100, 200)
	text := "edited"

	_, err := c.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:    inputPeerUser(300),
		Id:      7,
		Message: &text,
	})
	if err != tg.ErrMsgIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrMsgIdInvalid)
	}
}

func TestMessagesEditMessage_RemoteMsgNotModifiedMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		editMessage: func(_ context.Context, _ *msg.TLMsgEditMessage) (*tg.Updates, error) {
			return nil, errors.New("remote or network error: biz error: msg: message not modified")
		},
	}, 100, 200)
	text := "edited"

	_, err := c.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:    inputPeerUser(300),
		Id:      7,
		Message: &text,
	})
	if err != tg.ErrMessageNotModified {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageNotModified)
	}
}

// --- Input validation (must NOT call msg) ---

func TestMessagesSendMessage_InputPeerSelfTargetsCurrentUser(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	if _, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerSelf(),
		Message:  "hello",
		RandomId: 42,
	}); err != nil {
		t.Fatalf("error = %v", err)
	}
	if got == nil || got.UserId != 100 || got.PeerId != 100 {
		t.Fatalf("unexpected msg request: %+v", got)
	}
}

func TestMessagesSendMessage_NilPeerRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_UnsupportedPeerRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerChannel(&tg.TLInputPeerChannel{ChannelId: 42}),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_EmptyMessageRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "",
		RandomId: 42,
	})
	if err != tg.ErrMessageEmpty {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageEmpty)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_WhitespaceMessageRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "   ",
		RandomId: 42,
	})
	if err != tg.ErrMessageEmpty {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageEmpty)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_MessageTooLongRejected(t *testing.T) {
	text := ""
	for i := 0; i < 4097; i++ {
		text += "a"
	}
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  text,
		RandomId: 42,
	})
	if err != tg.ErrMessageTooLong {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageTooLong)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_Message4096CodeUnitsAccepted(t *testing.T) {
	text := ""
	for i := 0; i < 4096; i++ {
		text += "a"
	}
	var got *msg.TLMsgSendMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  text,
		RandomId: 42,
	})
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
}

func TestMessagesSendMessage_RandomIdZeroRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 0,
	})
	if err != tg.ErrRandomIdEmpty {
		t.Fatalf("error = %v, want %v", err, tg.ErrRandomIdEmpty)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

// --- Supported semantic fields ---

func TestMessagesSendMessage_EntitiesPassedToMsg(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	bold := tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{Offset: 14, Length: 12})
	spoiler := tg.MakeTLMessageEntitySpoiler(&tg.TLMessageEntitySpoiler{Offset: 27, Length: 15})
	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:       inputPeerUser(300),
		Message:    "是的方法反反复复房贷首付分\nfrdddccccdde\nrwerwerwerwerwe\nererewrwe",
		RandomId:   42,
		ClearDraft: true,
		Entities:   []tg.MessageEntityClazz{bold, spoiler},
	})
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
	if got == nil || len(got.Message) != 1 {
		t.Fatalf("msg request = %#v, want one outbox", got)
	}
	message, ok := got.Message[0].Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message = %T, want *tg.TLMessage", got.Message[0].Message)
	}
	if len(message.Entities) != 2 || message.Entities[0] != bold || message.Entities[1] != spoiler {
		t.Fatalf("entities = %#v, want original bold and spoiler entities", message.Entities)
	}
	if !got.ClearDraft {
		t.Fatal("ClearDraft = false, want true")
	}
}

// --- Unsupported field rejection ---

func TestMessagesSendMessage_SilentTrueRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
		Silent:   true,
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_NoforwardsTrueRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:       inputPeerUser(300),
		Message:    "hello",
		RandomId:   42,
		Noforwards: true,
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_ScheduleDateRejected(t *testing.T) {
	called := false
	sched := int32(2000000)
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:         inputPeerUser(300),
		Message:      "hello",
		RandomId:     42,
		ScheduleDate: &sched,
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

func TestMessagesSendMessage_ReplyMarkupRejected(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:        inputPeerUser(300),
		Message:     "hello",
		RandomId:    42,
		ReplyMarkup: tg.MakeTLReplyKeyboardMarkup(&tg.TLReplyKeyboardMarkup{}),
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
	if called {
		t.Fatal("msg service was called but should not have been")
	}
}

// --- Default/empty fields accepted ---

func TestMessagesSendMessage_EmptyEntitiesAccepted(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
		Entities: []tg.MessageEntityClazz{},
	})
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
}

func TestMessagesSendMessage_SilentFalseAccepted(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
		Silent:   false,
	})
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
	if got == nil {
		t.Fatal("msg service was not called")
	}
}

// --- Error mapping ---

func TestMessagesSendMessage_RandomIdConflictMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, msg.ErrRandomIdConflict
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrRandomIdDuplicate {
		t.Fatalf("error = %v, want %v", err, tg.ErrRandomIdDuplicate)
	}
}

func TestMessagesSendMessage_ReceiverBackpressureMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, msg.ErrReceiverBackpressure
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_SenderSyncFailedMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, msg.ErrSenderSyncFailed
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_MsgStorageMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, msg.ErrMsgStorage
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_SendStateConflictMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, msg.ErrSendStateConflict
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_ContextDeadlineMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, context.DeadlineExceeded
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrTimeout {
		t.Fatalf("error = %v, want %v", err, tg.ErrTimeout)
	}
}

func TestMessagesSendMessage_UnknownErrorMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, errors.New("some unknown transport error")
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSendMessage_TgErrorPassThrough(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, tg.ErrChatIdInvalid
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrChatIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrChatIdInvalid)
	}
}

// --- Metadata validation ---

func TestMessagesSendMessage_MissingMetadataRejected(t *testing.T) {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{MsgClient: &messagesFakeMsgClient{}},
	})

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
}

func TestMessagesSendMessage_UserIdZeroRejected(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{}, 0, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
}
