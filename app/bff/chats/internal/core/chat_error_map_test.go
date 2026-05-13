package core

import (
	"errors"
	"fmt"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface/ecode"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMapChatError(t *testing.T) {
	unknown := errors.New("unknown")

	tests := []struct {
		name string
		err  error
		want error
	}{
		{name: "storage maps to internal server error", err: fmt.Errorf("wrapped: %w", chatpb.ErrChatStorage), want: tg.ErrInternalServerError},
		{name: "chat not found maps to chat id invalid", err: fmt.Errorf("wrapped: %w", chatpb.ErrChatNotFound), want: tg.ErrChatIdInvalid},
		{name: "invite invalid maps to invite hash invalid", err: fmt.Errorf("wrapped: %w", chatpb.ErrInviteHashInvalid), want: tg.ErrInviteHashInvalid},
		{name: "already participant maps to user already participant", err: fmt.Errorf("wrapped: %w", chatpb.ErrUserAlreadyParticipant), want: tg.ErrUserAlreadyParticipant},
		{name: "chat admin required maps to 400 chat admin required", err: fmt.Errorf("wrapped: %w", chatpb.ErrChatAdminRequired), want: tg.Err400ChatAdminRequired},
		{name: "chat title empty maps to chat title empty", err: fmt.Errorf("wrapped: %w", chatpb.ErrChatTitleEmpty), want: tg.ErrChatTitleEmpty},
		{name: "chat not modified maps to chat not modified", err: fmt.Errorf("wrapped: %w", chatpb.ErrChatNotModified), want: tg.ErrChatNotModified},
		{name: "participant invalid maps to peer id invalid", err: fmt.Errorf("wrapped: %w", chatpb.ErrParticipantInvalid), want: tg.Err400PeerIdInvalid},
		{name: "input user deactivated maps to input user deactivated", err: fmt.Errorf("wrapped: %w", chatpb.ErrInputUserDeactivated), want: tg.ErrInputUserDeactivated},
		{name: "user not participant maps to 400 user not participant", err: fmt.Errorf("wrapped: %w", chatpb.ErrUserNotParticipant), want: tg.Err400UserNotParticipant},
		{name: "users too few maps to users too few", err: fmt.Errorf("wrapped: %w", chatpb.ErrUsersTooFew), want: tg.ErrUsersTooFew},
		{name: "users too much maps to users too much", err: fmt.Errorf("wrapped: %w", chatpb.ErrUsersTooMuch), want: tg.ErrUsersTooMuch},
		{name: "invite expired maps to 400 invite hash expired", err: fmt.Errorf("wrapped: %w", chatpb.ErrInviteHashExpired), want: tg.Err400InviteHashExpired},
		{name: "chat link exists maps to chat link exists", err: fmt.Errorf("wrapped: %w", chatpb.ErrChatLinkExists), want: tg.ErrChatLinkExists},
		{name: "unknown passthrough", err: unknown, want: unknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapChatError(tt.err); got != tt.want {
				t.Fatalf("mapChatError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapChatErrorMigratedAndDeactivatedMapToChatIDInvalid(t *testing.T) {
	for _, err := range []error{chatpb.ErrChatMigrated, chatpb.ErrChatDeactivated} {
		if got := mapChatError(fmt.Errorf("wrapped: %w", err)); got != tg.ErrChatIdInvalid {
			t.Fatalf("mapChatError(%v) = %v, want %v", err, got, tg.ErrChatIdInvalid)
		}
	}
}

func TestMapChatErrorFloodPreservesWaitSeconds(t *testing.T) {
	got := mapChatError(fmt.Errorf("wrapped: %w", chatpb.NewCreateChatFloodError(37)))

	gotCode, wantCode := mustCodeError(t, got), mustCodeError(t, tg.NewErrFloodWaitX(37))
	if gotCode.Code() != wantCode.Code() || gotCode.Msg() != wantCode.Msg() {
		t.Fatalf("mapChatError(flood) = %v, want %v", got, tg.NewErrFloodWaitX(37))
	}
}

func TestMapChatErrorPendingPreservesWaitSeconds(t *testing.T) {
	got := mapChatError(fmt.Errorf("wrapped: %w", chatpb.NewCreateChatOperationPendingError(23)))

	gotCode, wantCode := mustCodeError(t, got), mustCodeError(t, tg.NewErrFloodWaitX(23))
	if gotCode.Code() != wantCode.Code() || gotCode.Msg() != wantCode.Msg() {
		t.Fatalf("mapChatError(pending) = %v, want %v", got, tg.NewErrFloodWaitX(23))
	}
}

func mustCodeError(t *testing.T, err error) ecode.CodeError {
	t.Helper()

	var codeErr ecode.CodeError
	if !errors.As(err, &codeErr) {
		t.Fatalf("%v is not ecode.CodeError", err)
	}
	return codeErr
}
