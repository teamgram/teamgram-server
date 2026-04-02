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
