package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestDialogGetDialogsReturnsEmptyPlaceholderList(t *testing.T) {
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
	if len(result.Datas) != 0 {
		t.Fatalf("expected empty dialog list, got %d items", len(result.Datas))
	}
}

func TestDialogGetMyDialogsDataReturnsEmptySimpleDataPlaceholder(t *testing.T) {
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
	if len(simpleData.Users) != 0 || len(simpleData.Chats) != 0 || len(simpleData.Channels) != 0 {
		t.Fatalf("expected empty ids, got users=%d chats=%d channels=%d",
			len(simpleData.Users), len(simpleData.Chats), len(simpleData.Channels))
	}
}
