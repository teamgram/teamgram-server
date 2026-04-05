package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// --- fake dialog client ---

type fakeDialogQueryClient struct {
	req           *dialog.TLDialogGetDialogs
	dialogByIDReq *dialog.TLDialogGetDialogById
	dataReq       *dialog.TLDialogGetMyDialogsData
	result        *dialog.VectorDialogExt
	dataResult    *dialog.DialogsData
	dialogResult  *dialog.DialogExt
}

func (f *fakeDialogQueryClient) DialogGetDialogs(ctx context.Context, in *dialog.TLDialogGetDialogs) (*dialog.VectorDialogExt, error) {
	f.req = in
	return f.result, nil
}

func (f *fakeDialogQueryClient) DialogGetDialogById(ctx context.Context, in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	f.dialogByIDReq = in
	return f.dialogResult, nil
}

func (f *fakeDialogQueryClient) DialogGetMyDialogsData(ctx context.Context, in *dialog.TLDialogGetMyDialogsData) (*dialog.DialogsData, error) {
	f.dataReq = in
	return f.dataResult, nil
}

var _ svc.DialogQueryClient = (*fakeDialogQueryClient)(nil)

type fakeMessageQueryClient struct {
	req    *message.TLMessageGetHistoryMessages
	result *message.VectorMessageBox
}

func (f *fakeMessageQueryClient) MessageGetHistoryMessages(ctx context.Context, in *message.TLMessageGetHistoryMessages) (*message.VectorMessageBox, error) {
	f.req = in
	return f.result, nil
}

var _ svc.MessageQueryClient = (*fakeMessageQueryClient)(nil)

func TestMessagesGetDialogsReturnsSinglePlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 42, AccessHash: 0}),
		Limit:      20,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected dialogs result, got nil")
	}

	dialogsSlice, ok := result.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("expected messages.dialogsSlice placeholder, got %T", result.Clazz)
	}
	if dialogsSlice.Count != 1 {
		t.Fatalf("expected count=1, got %d", dialogsSlice.Count)
	}
	if len(dialogsSlice.Dialogs) != 1 || len(dialogsSlice.Messages) != 1 || len(dialogsSlice.Chats) != 0 || len(dialogsSlice.Users) != 1 {
		t.Fatalf("expected single placeholder dialog/message/user, got dialogs=%d messages=%d chats=%d users=%d",
			len(dialogsSlice.Dialogs), len(dialogsSlice.Messages), len(dialogsSlice.Chats), len(dialogsSlice.Users))
	}
	dialog, ok := dialogsSlice.Dialogs[0].(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected dialog placeholder, got %T", dialogsSlice.Dialogs[0])
	}
	if dialog.TopMessage != 10 {
		t.Fatalf("expected top_message=10, got %d", dialog.TopMessage)
	}
	user, ok := dialogsSlice.Users[0].(*tg.TLUserEmpty)
	if !ok {
		t.Fatalf("expected userEmpty placeholder, got %T", dialogsSlice.Users[0])
	}
	if user.Id != 42 {
		t.Fatalf("expected placeholder user id=42, got %d", user.Id)
	}
}

func TestMessagesGetDialogsReturnsEmptyDialogsSlicePlaceholderForEmptyPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Limit:      20,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected dialogs result, got nil")
	}

	dialogsSlice, ok := result.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("expected messages.dialogsSlice placeholder, got %T", result.Clazz)
	}
	if dialogsSlice.Count != 0 {
		t.Fatalf("expected count=0, got %d", dialogsSlice.Count)
	}
	if len(dialogsSlice.Dialogs) != 0 || len(dialogsSlice.Messages) != 0 || len(dialogsSlice.Chats) != 0 || len(dialogsSlice.Users) != 0 {
		t.Fatalf("expected empty placeholder lists, got dialogs=%d messages=%d chats=%d users=%d",
			len(dialogsSlice.Dialogs), len(dialogsSlice.Messages), len(dialogsSlice.Chats), len(dialogsSlice.Users))
	}
}

func TestMessagesGetPeerDialogsReturnsEmptyPeerDialogsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesGetPeerDialogs(&tg.TLMessagesGetPeerDialogs{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected peer dialogs result, got nil")
	}
	if len(result.Dialogs) != 0 || len(result.Messages) != 0 || len(result.Chats) != 0 || len(result.Users) != 0 {
		t.Fatalf("expected empty placeholder lists, got dialogs=%d messages=%d chats=%d users=%d",
			len(result.Dialogs), len(result.Messages), len(result.Chats), len(result.Users))
	}
	if result.State == nil {
		t.Fatal("expected placeholder updates state, got nil")
	}
}

func TestMessagesGetPeerDialogsReturnsSinglePlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesGetPeerDialogs(&tg.TLMessagesGetPeerDialogs{
		Peers: []tg.InputDialogPeerClazz{
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{
				Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 42, AccessHash: 0}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected peer dialogs result, got nil")
	}
	if len(result.Dialogs) != 1 {
		t.Fatalf("expected 1 dialog placeholder, got %d", len(result.Dialogs))
	}
	if len(result.Messages) != 1 {
		t.Fatalf("expected 1 message placeholder, got %d", len(result.Messages))
	}
	if len(result.Users) != 1 {
		t.Fatalf("expected 1 user placeholder, got %d", len(result.Users))
	}
	dialog, ok := result.Dialogs[0].(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected dialog placeholder, got %T", result.Dialogs[0])
	}
	if dialog.TopMessage != 10 {
		t.Fatalf("expected top_message=10, got %d", dialog.TopMessage)
	}
	user, ok := result.Users[0].(*tg.TLUserEmpty)
	if !ok {
		t.Fatalf("expected userEmpty placeholder, got %T", result.Users[0])
	}
	if user.Id != 42 {
		t.Fatalf("expected placeholder user id=42, got %d", user.Id)
	}
}

func TestDialogsPinUnreadAndTTLPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	toggleResult, err := c.MessagesToggleDialogPin(&tg.TLMessagesToggleDialogPin{
		Peer:   tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 1})}),
		Pinned: true,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(toggleResult) {
		t.Fatalf("expected toggle pin boolTrue, got %#v", toggleResult)
	}

	reorderResult, err := c.MessagesReorderPinnedDialogs(&tg.TLMessagesReorderPinnedDialogs{
		Force: true,
		Order: []tg.InputDialogPeerClazz{
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 1})}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(reorderResult) {
		t.Fatalf("expected reorder pin boolTrue, got %#v", reorderResult)
	}

	pinnedDialogs, err := c.MessagesGetPinnedDialogs(&tg.TLMessagesGetPinnedDialogs{FolderId: 0})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if pinnedDialogs == nil || len(pinnedDialogs.Dialogs) != 1 || len(pinnedDialogs.Messages) != 1 || len(pinnedDialogs.Users) != 1 {
		t.Fatalf("expected single pinned dialog placeholder, got %#v", pinnedDialogs)
	}

	unreadResult, err := c.MessagesMarkDialogUnread(&tg.TLMessagesMarkDialogUnread{
		Peer:   tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 1})}),
		Unread: true,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(unreadResult) {
		t.Fatalf("expected mark unread boolTrue, got %#v", unreadResult)
	}

	marks, err := c.MessagesGetDialogUnreadMarks(&tg.TLMessagesGetDialogUnreadMarks{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if marks == nil || len(marks.Datas) != 1 {
		t.Fatalf("expected one unread mark placeholder, got %#v", marks)
	}

	ttlUpdates, err := c.MessagesSetHistoryTTL(&tg.TLMessagesSetHistoryTTL{
		Peer:   tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 1}),
		Period: 86400,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if ttlUpdates == nil {
		t.Fatal("expected ttl updates placeholder, got nil")
	}
}

func TestDialogsPeerMetaPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)
	userPeer := tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 42, AccessHash: 0})

	peerSettings, err := c.MessagesGetPeerSettings(&tg.TLMessagesGetPeerSettings{Peer: userPeer})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if peerSettings == nil || peerSettings.Settings == nil {
		t.Fatalf("expected peer settings placeholder, got %#v", peerSettings)
	}

	hideBar, err := c.MessagesHidePeerSettingsBar(&tg.TLMessagesHidePeerSettingsBar{Peer: userPeer})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(hideBar) {
		t.Fatalf("expected hide bar boolTrue, got %#v", hideBar)
	}

	setTyping, err := c.MessagesSetTyping(&tg.TLMessagesSetTyping{
		Peer:   userPeer,
		Action: tg.MakeTLSendMessageTypingAction(&tg.TLSendMessageTypingAction{}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(setTyping) {
		t.Fatalf("expected set typing boolTrue, got %#v", setTyping)
	}

	onlines, err := c.MessagesGetOnlines(&tg.TLMessagesGetOnlines{Peer: userPeer})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if onlines == nil || onlines.Onlines != 1 {
		t.Fatalf("expected onlines=1 placeholder, got %#v", onlines)
	}

	screenshot, err := c.MessagesSendScreenshotNotification(&tg.TLMessagesSendScreenshotNotification{
		Peer:     userPeer,
		RandomId: 42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	short, ok := screenshot.ToUpdateShort()
	if !ok {
		t.Fatalf("expected updateShort placeholder, got %T", screenshot.Clazz)
	}
	updateNewMessage, ok := short.Update.(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("expected updateNewMessage placeholder, got %T", short.Update)
	}
	message, ok := updateNewMessage.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("expected tlMessage placeholder, got %T", updateNewMessage.Message)
	}
	if message.Id != 42 {
		t.Fatalf("expected placeholder message id=42, got %d", message.Id)
	}
}

func TestMessagesGetDialogsDelegatesToDialogClient(t *testing.T) {
	fakeDialogCli := &fakeDialogQueryClient{
		result: &dialog.VectorDialogExt{
			Datas: []dialog.DialogExtClazz{
				dialog.MakeTLDialogExt(&dialog.TLDialogExt{
					Order: 1,
					Dialog: tg.MakeTLDialog(&tg.TLDialog{
						Peer:            tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 99}),
						TopMessage:      5,
						ReadInboxMaxId:  5,
						ReadOutboxMaxId: 5,
						UnreadCount:     0,
						NotifySettings:  tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
					}),
					AvailableMinId: 1,
					Date:           10,
				}),
			},
		},
		dataResult: dialog.MakeTLSimpleDialogsData(&dialog.TLSimpleDialogsData{
			Users:    []int64{99},
			Chats:    []int64{7},
			Channels: []int64{8},
		}).ToDialogsData(),
	}
	fakeMessageCli := &fakeMessageQueryClient{
		result: &message.VectorMessageBox{
			Datas: []tg.MessageBoxClazz{
				tg.MakeTLMessageBox(&tg.TLMessageBox{
					UserId:    100,
					MessageId: 5,
					PeerType:  tg.PEER_USER,
					PeerId:    99,
					Message: tg.MakeTLMessage(&tg.TLMessage{
						Id:      5,
						Date:    10,
						Message: "from-service",
					}),
				}),
			},
		},
	}

	c := newWithMD(context.Background(), &svc.ServiceContext{
		DialogClient:  fakeDialogCli,
		MessageClient: fakeMessageCli,
	}, 100)

	result, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 42}),
		Limit:      20,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	if fakeDialogCli.req == nil {
		t.Fatal("expected DialogGetDialogs to be called")
	}
	if fakeDialogCli.req.UserId != 100 {
		t.Errorf("expected userId=100, got %d", fakeDialogCli.req.UserId)
	}
	if fakeDialogCli.dataReq == nil {
		t.Fatal("expected DialogGetMyDialogsData to be called")
	}
	if fakeMessageCli.req == nil {
		t.Fatal("expected MessageGetHistoryMessages to be called")
	}

	dialogsSlice, ok := result.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("expected messages.dialogsSlice, got %T", result.Clazz)
	}
	if dialogsSlice.Count != 1 {
		t.Fatalf("expected count=1, got %d", dialogsSlice.Count)
	}
	if len(dialogsSlice.Dialogs) != 1 {
		t.Fatalf("expected 1 dialog, got %d", len(dialogsSlice.Dialogs))
	}
	if len(dialogsSlice.Messages) != 1 {
		t.Fatalf("expected 1 delegated message, got %d", len(dialogsSlice.Messages))
	}
	if len(dialogsSlice.Users) != 1 {
		t.Fatalf("expected 1 user from dialogs data, got %d", len(dialogsSlice.Users))
	}
	if len(dialogsSlice.Chats) != 2 {
		t.Fatalf("expected 2 chats from dialogs data, got %d", len(dialogsSlice.Chats))
	}
	messageItem, ok := dialogsSlice.Messages[0].(*tg.TLMessage)
	if !ok || messageItem.Message != "from-service" {
		t.Fatalf("expected delegated service message, got %#v", dialogsSlice.Messages[0])
	}
}

func TestMessagesGetPeerDialogsDelegatesToDialogClient(t *testing.T) {
	fakeDialogCli := &fakeDialogQueryClient{
		dialogResult: dialog.MakeTLDialogExt(&dialog.TLDialogExt{
			Order: 1,
			Dialog: tg.MakeTLDialog(&tg.TLDialog{
				Peer:            tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 42}),
				TopMessage:      77,
				ReadInboxMaxId:  77,
				ReadOutboxMaxId: 77,
				UnreadCount:     0,
				NotifySettings:  tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
			}),
			AvailableMinId: 1,
			Date:           10,
		}),
	}
	fakeMessageCli := &fakeMessageQueryClient{
		result: &message.VectorMessageBox{
			Datas: []tg.MessageBoxClazz{
				tg.MakeTLMessageBox(&tg.TLMessageBox{
					UserId:    100,
					MessageId: 77,
					PeerType:  tg.PEER_USER,
					PeerId:    42,
					Message: tg.MakeTLMessage(&tg.TLMessage{
						Id:      77,
						Date:    10,
						Message: "peer-dialog-message",
					}),
				}),
			},
		},
	}

	c := newWithMD(context.Background(), &svc.ServiceContext{
		DialogClient:  fakeDialogCli,
		MessageClient: fakeMessageCli,
	}, 100)

	result, err := c.MessagesGetPeerDialogs(&tg.TLMessagesGetPeerDialogs{
		Peers: []tg.InputDialogPeerClazz{
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{
				Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 42}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if fakeDialogCli.dialogByIDReq == nil {
		t.Fatal("expected DialogGetDialogById to be called")
	}
	if fakeMessageCli.req == nil {
		t.Fatal("expected MessageGetHistoryMessages to be called")
	}
	if fakeDialogCli.dialogByIDReq.UserId != 100 {
		t.Fatalf("expected dialog lookup user_id=100, got %d", fakeDialogCli.dialogByIDReq.UserId)
	}
	if len(result.Dialogs) != 1 || len(result.Messages) != 1 || len(result.Users) != 1 {
		t.Fatalf("expected one peer dialog envelope item, got dialogs=%d messages=%d users=%d",
			len(result.Dialogs), len(result.Messages), len(result.Users))
	}
	dialogItem, ok := result.Dialogs[0].(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected delegated dialog, got %T", result.Dialogs[0])
	}
	if dialogItem.TopMessage != 77 {
		t.Fatalf("expected delegated top_message=77, got %d", dialogItem.TopMessage)
	}
	messageItem, ok := result.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("expected delegated message, got %T", result.Messages[0])
	}
	if messageItem.Id != 77 || messageItem.Message != "peer-dialog-message" {
		t.Fatalf("expected delegated message id=77, got %#v", messageItem)
	}
}

func newWithMD(ctx context.Context, svcCtx *svc.ServiceContext, userId int64) *DialogsCore {
	c := New(ctx, svcCtx)
	c.MD = &metadata.RpcMetadata{UserId: userId}
	return c
}
