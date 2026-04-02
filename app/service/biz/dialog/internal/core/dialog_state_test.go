package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestDialogGetDialogsReturnsSinglePlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogs(&dialog.TLDialogGetDialogs{
		UserId:        1,
		ExcludePinned: tg.BoolTrueClazz,
		FolderId:      0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected dialogs vector, got nil")
	}
	if len(result.Datas) != 1 {
		t.Fatalf("expected single placeholder dialog, got %d items", len(result.Datas))
	}
	dialogExt, ok := result.Datas[0].(*dialog.TLDialogExt)
	if !ok {
		t.Fatalf("expected dialogExt placeholder, got %T", result.Datas[0])
	}
	placeholderDialog, ok := dialogExt.Dialog.(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected embedded dialog placeholder, got %T", dialogExt.Dialog)
	}
	if placeholderDialog.TopMessage != 10 {
		t.Fatalf("expected top_message=10, got %d", placeholderDialog.TopMessage)
	}
}

func TestDialogGetMyDialogsDataReturnsUserPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetMyDialogsData(&dialog.TLDialogGetMyDialogsData{
		UserId: 1,
		User:   true,
		Chat:   true,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected dialogs data, got nil")
	}

	simpleData, ok := result.Clazz.(*dialog.TLSimpleDialogsData)
	if !ok {
		t.Fatalf("expected simpleDialogsData placeholder, got %T", result.Clazz)
	}
	if len(simpleData.Users) != 1 || simpleData.Users[0] != 1 {
		t.Fatalf("expected user placeholder id=1, got %#v", simpleData.Users)
	}
	if len(simpleData.Chats) != 0 || len(simpleData.Channels) != 0 {
		t.Fatalf("expected empty chats/channels, got chats=%d channels=%d",
			len(simpleData.Chats), len(simpleData.Channels))
	}
}

func TestDialogGetDialogByIdReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogById(&dialog.TLDialogGetDialogById{
		UserId:   1,
		PeerType: tg.PEER_CHAT,
		PeerId:   42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.Clazz == nil {
		t.Fatal("expected dialogExt placeholder, got nil")
	}
	dialogExt, ok := result.Clazz.(*dialog.TLDialogExt)
	if !ok {
		t.Fatalf("expected dialogExt, got %T", result.Clazz)
	}
	placeholderDialog, ok := dialogExt.Dialog.(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected embedded dialog placeholder, got %T", dialogExt.Dialog)
	}
	peer, ok := placeholderDialog.Peer.(*tg.TLPeerChat)
	if !ok || peer.ChatId != 42 {
		t.Fatalf("expected peerChat(42), got %#v", placeholderDialog.Peer)
	}
}

func TestDialogGetDialogsByIdListReturnsPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogsByIdList(&dialog.TLDialogGetDialogsByIdList{
		UserId: 1,
		IdList: []int64{11, 22},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected vector, got nil")
	}
	if len(result.Datas) != 2 {
		t.Fatalf("expected two placeholders, got %d", len(result.Datas))
	}
}

func TestDialogGetDialogsCountReturnsPlaceholderCount(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogsCount(&dialog.TLDialogGetDialogsCount{
		UserId:        1,
		ExcludePinned: tg.BoolFalseClazz,
		FolderId:      0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.V != 1 {
		t.Fatalf("expected placeholder count=1, got %#v", result)
	}
}

func TestDialogGetPinnedDialogsReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetPinnedDialogs(&dialog.TLDialogGetPinnedDialogs{
		UserId:   1,
		FolderId: 0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || len(result.Datas) != 1 {
		t.Fatalf("expected one pinned placeholder, got %#v", result)
	}
}

func TestDialogGetTopMessageReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetTopMessage(&dialog.TLDialogGetTopMessage{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.V != 10 {
		t.Fatalf("expected top message=10, got %#v", result)
	}
}

func TestDialogPinnedMessagePlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	getResult, err := c.DialogGetUserPinnedMessage(&dialog.TLDialogGetUserPinnedMessage{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if getResult == nil || getResult.V != 1 {
		t.Fatalf("expected pinned placeholder id=1, got %#v", getResult)
	}

	updateResult, err := c.DialogUpdateUserPinnedMessage(&dialog.TLDialogUpdateUserPinnedMessage{
		UserId:      1,
		PeerType:    tg.PEER_USER,
		PeerId:      2,
		PinnedMsgId: 9,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(updateResult) {
		t.Fatalf("expected boolTrue placeholder, got %#v", updateResult)
	}
}

func TestDialogWritePlaceholders(t *testing.T) {
	c := New(context.Background(), nil)
	unreadCount := int32(3)

	insertResult, err := c.DialogInsertOrUpdateDialog(&dialog.TLDialogInsertOrUpdateDialog{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(insertResult) {
		t.Fatalf("expected insert/update placeholder boolTrue, got %#v", insertResult)
	}

	deleteResult, err := c.DialogDeleteDialog(&dialog.TLDialogDeleteDialog{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(deleteResult) {
		t.Fatalf("expected delete placeholder boolTrue, got %#v", deleteResult)
	}

	unreadResult, err := c.DialogUpdateUnreadCount(&dialog.TLDialogUpdateUnreadCount{
		UserId:      1,
		PeerType:    tg.PEER_USER,
		PeerId:      2,
		UnreadCount: &unreadCount,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(unreadResult) {
		t.Fatalf("expected unread-count placeholder boolTrue, got %#v", unreadResult)
	}

	markResult, err := c.DialogMarkDialogUnread(&dialog.TLDialogMarkDialogUnread{
		UserId:     1,
		PeerType:   tg.PEER_USER,
		PeerId:     2,
		UnreadMark: tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(markResult) {
		t.Fatalf("expected unread-mark placeholder boolTrue, got %#v", markResult)
	}
}

func TestDialogGetDialogsByOffsetDateReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogsByOffsetDate(&dialog.TLDialogGetDialogsByOffsetDate{
		UserId:        1,
		ExcludePinned: tg.BoolFalseClazz,
		OffsetDate:    0,
		Limit:         10,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || len(result.Datas) != 1 {
		t.Fatalf("expected one offset-date placeholder, got %#v", result)
	}
}

func TestDialogGetDialogUnreadMarkListReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogUnreadMarkList(&dialog.TLDialogGetDialogUnreadMarkList{
		UserId: 1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || len(result.Datas) != 1 {
		t.Fatalf("expected one unread-mark placeholder, got %#v", result)
	}
	peer, ok := result.Datas[0].(*tg.TLDialogPeer)
	if !ok {
		t.Fatalf("expected dialogPeer placeholder, got %T", result.Datas[0])
	}
	userPeer, ok := peer.Peer.(*tg.TLPeerUser)
	if !ok || userPeer.UserId != 1 {
		t.Fatalf("expected peerUser(1), got %#v", peer.Peer)
	}
}
