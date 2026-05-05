package core

import (
	"context"
	"errors"
	"testing"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/svc"
)

func TestDialogDeprecatedHandlersReturnDeprecatedMethod(t *testing.T) {
	core := New(context.Background(), &svc.ServiceContext{})

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "get my dialogs data",
			call: func() error {
				_, err := core.DialogGetMyDialogsData(&dialogpb.TLDialogGetMyDialogsData{})
				return err
			},
		},
		{
			name: "get dialogs by offset date",
			call: func() error {
				_, err := core.DialogGetDialogsByOffsetDate(&dialogpb.TLDialogGetDialogsByOffsetDate{})
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); !errors.Is(err, dialogpb.ErrDeprecatedMethod) {
				t.Fatalf("error = %v, want ErrDeprecatedMethod", err)
			}
		})
	}
}

func TestDialogOwnerBoundaryHandlersReturnWrongOwner(t *testing.T) {
	core := New(context.Background(), &svc.ServiceContext{})

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "insert or update dialog",
			call: func() error {
				_, err := core.DialogInsertOrUpdateDialog(&dialogpb.TLDialogInsertOrUpdateDialog{})
				return err
			},
		},
		{
			name: "update unread count",
			call: func() error {
				_, err := core.DialogUpdateUnreadCount(&dialogpb.TLDialogUpdateUnreadCount{})
				return err
			},
		},
		{
			name: "update user pinned message",
			call: func() error {
				_, err := core.DialogUpdateUserPinnedMessage(&dialogpb.TLDialogUpdateUserPinnedMessage{})
				return err
			},
		},
		{
			name: "get top message",
			call: func() error {
				_, err := core.DialogGetTopMessage(&dialogpb.TLDialogGetTopMessage{})
				return err
			},
		},
		{
			name: "get user pinned message",
			call: func() error {
				_, err := core.DialogGetUserPinnedMessage(&dialogpb.TLDialogGetUserPinnedMessage{})
				return err
			},
		},
		{
			name: "get channel read participants",
			call: func() error {
				_, err := core.DialogGetChannelMessageReadParticipants(&dialogpb.TLDialogGetChannelMessageReadParticipants{})
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); !errors.Is(err, dialogpb.ErrWrongOwner) {
				t.Fatalf("error = %v, want ErrWrongOwner", err)
			}
		})
	}
}
