package core

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeInviteRepo struct {
	invite          *model.ChatInvites
	inviteErr       error
	usage           int32
	usageErr        error
	requesters      *chat.RecentChatInviteRequesters
	requestersErr   error
	recordCalls     int
	recordArg       repository.InviteParticipantArg
	hideCalls       int
	hideArg         repository.HideJoinRequestsArg
	editInvites     []tg.ExportedChatInviteClazz
	editArg         repository.EditExportedChatInviteArg
	createInvite    *tg.ExportedChatInvite
	exportedInvite  *tg.ExportedChatInvite
	exportedInvites []tg.ExportedChatInviteClazz
	importers       []tg.ChatInviteImporterClazz
	err             error
}

func (f *fakeInviteRepo) CreateExportedChatInvite(ctx context.Context, arg repository.ExportChatInviteArg) (*tg.ExportedChatInvite, error) {
	return f.createInvite, f.err
}

func (f *fakeInviteRepo) GetExportedChatInvite(ctx context.Context, chatID int64, link string) (*tg.ExportedChatInvite, error) {
	return f.exportedInvite, f.err
}

func (f *fakeInviteRepo) GetExportedChatInvites(ctx context.Context, chatID, adminID int64, revoked bool, offsetDate *int32, offsetLink *string, limit int32) ([]tg.ExportedChatInviteClazz, error) {
	return f.exportedInvites, f.err
}

func (f *fakeInviteRepo) EditExportedChatInvite(ctx context.Context, arg repository.EditExportedChatInviteArg) ([]tg.ExportedChatInviteClazz, error) {
	f.editArg = arg
	return f.editInvites, f.err
}

func (f *fakeInviteRepo) DeleteExportedChatInvite(ctx context.Context, chatID int64, link string) error {
	return f.err
}

func (f *fakeInviteRepo) DeleteRevokedExportedChatInvites(ctx context.Context, chatID, adminID int64) error {
	return f.err
}

func (f *fakeInviteRepo) GetChatInviteByLink(ctx context.Context, link string) (*model.ChatInvites, error) {
	return f.invite, f.inviteErr
}

func (f *fakeInviteRepo) GetAdminsWithInvites(ctx context.Context, chatID int64, adminIDs []int64) ([]tg.ChatAdminWithInvitesClazz, error) {
	return nil, f.err
}

func (f *fakeInviteRepo) CountChatInviteParticipants(ctx context.Context, link string, requested bool) (int32, error) {
	return f.usage, f.usageErr
}

func (f *fakeInviteRepo) RecordInviteParticipant(ctx context.Context, arg repository.InviteParticipantArg) error {
	f.recordCalls++
	f.recordArg = arg
	return f.err
}

func (f *fakeInviteRepo) GetChatInviteImporters(ctx context.Context, q repository.ChatInviteImporterQuery) ([]tg.ChatInviteImporterClazz, error) {
	return f.importers, f.err
}

func (f *fakeInviteRepo) GetRecentChatInviteRequesters(ctx context.Context, chatID int64) (*chat.RecentChatInviteRequesters, error) {
	return f.requesters, f.requestersErr
}

func (f *fakeInviteRepo) HideChatJoinRequest(ctx context.Context, arg repository.HideJoinRequestsArg) error {
	f.hideCalls++
	f.hideArg = arg
	return f.err
}

func newInviteTestCore(read *fakeReadRepo, write *fakeWriteRepo, invite *fakeInviteRepo) *ChatCore {
	return &ChatCore{
		ctx:        context.Background(),
		readRepo:   read,
		writeRepo:  write,
		inviteRepo: invite,
	}
}

func baseInviteRow(requestNeeded bool) *model.ChatInvites {
	return &model.ChatInvites{
		ChatId:        10,
		AdminId:       1,
		Link:          "hash",
		RequestNeeded: requestNeeded,
		UsageLimit:    3,
	}
}

func inviteMutableChat(normalSelf bool) *tg.MutableChat {
	state := chat.ChatMemberStateLeft
	if normalSelf {
		state = chat.ChatMemberStateNormal
	}
	return mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil),
		participantForMemberTests(10, 2, chat.ChatMemberNormal, state, nil))
}

func TestChatInviteInvalidHashReturnsSemanticError(t *testing.T) {
	core := newInviteTestCore(&fakeReadRepo{}, &fakeWriteRepo{}, &fakeInviteRepo{inviteErr: chat.ErrInviteHashInvalid})
	_, err := core.ChatCheckChatInvite(&chat.TLChatCheckChatInvite{SelfId: 2, Hash: "bad"})
	if !errors.Is(err, chat.ErrInviteHashInvalid) {
		t.Fatalf("ChatCheckChatInvite error = %v, want ErrInviteHashInvalid", err)
	}
}

func TestChatInviteExpiredBranchesReturnExpired(t *testing.T) {
	tests := []struct {
		name   string
		invite *model.ChatInvites
		usage  int32
	}{
		{name: "revoked", invite: func() *model.ChatInvites { v := baseInviteRow(false); v.Revoked = true; return v }()},
		{name: "expired date", invite: func() *model.ChatInvites { v := baseInviteRow(false); v.ExpireDate = 1; return v }()},
		{name: "usage limit", invite: baseInviteRow(false), usage: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := newInviteTestCore(&fakeReadRepo{mutableChat: inviteMutableChat(false)}, &fakeWriteRepo{}, &fakeInviteRepo{invite: tt.invite, usage: tt.usage})
			_, err := core.ChatCheckChatInvite(&chat.TLChatCheckChatInvite{SelfId: 2, Hash: "hash"})
			if !errors.Is(err, chat.ErrInviteHashExpired) {
				t.Fatalf("ChatCheckChatInvite error = %v, want ErrInviteHashExpired", err)
			}
		})
	}
}

func TestChatInviteMigratedChatReturnsSemanticError(t *testing.T) {
	mChat := inviteMutableChat(false)
	mChat.Chat.Deactivated = true
	mChat.Chat.MigratedTo = tg.MakeTLInputChannel(&tg.TLInputChannel{ChannelId: 99, AccessHash: 100})
	core := newInviteTestCore(&fakeReadRepo{mutableChat: mChat}, &fakeWriteRepo{}, &fakeInviteRepo{invite: baseInviteRow(false)})
	_, err := core.ChatImportChatInvite2(&chat.TLChatImportChatInvite2{SelfId: 2, Hash: "hash"})
	if !errors.Is(err, chat.ErrChatMigrated) {
		t.Fatalf("ChatImportChatInvite2 error = %v, want ErrChatMigrated", err)
	}
}

func TestChatInviteAlreadyParticipantReturnsSemanticError(t *testing.T) {
	core := newInviteTestCore(&fakeReadRepo{mutableChat: inviteMutableChat(true)}, &fakeWriteRepo{}, &fakeInviteRepo{invite: baseInviteRow(false)})
	_, err := core.ChatImportChatInvite2(&chat.TLChatImportChatInvite2{SelfId: 2, Hash: "hash"})
	if !errors.Is(err, chat.ErrUserAlreadyParticipant) {
		t.Fatalf("ChatImportChatInvite2 error = %v, want ErrUserAlreadyParticipant", err)
	}
}

func TestChatInviteRequestNeededImportReturnsImportedResult(t *testing.T) {
	requesters := chat.MakeTLRecentChatInviteRequesters(&chat.TLRecentChatInviteRequesters{
		RequestsPending:  1,
		RecentRequesters: []int64{2},
	}).ToRecentChatInviteRequesters()
	inviteRepo := &fakeInviteRepo{invite: baseInviteRow(true), requesters: requesters}
	core := newInviteTestCore(&fakeReadRepo{mutableChat: inviteMutableChat(false)}, &fakeWriteRepo{}, inviteRepo)
	got, err := core.ChatImportChatInvite2(&chat.TLChatImportChatInvite2{SelfId: 2, Hash: "hash"})
	if err != nil {
		t.Fatalf("ChatImportChatInvite2 error = %v", err)
	}
	if got == nil || got.Chat == nil || got.Requesters == nil {
		t.Fatalf("ChatImportChatInvite2 = %#v, want imported result with requesters", got)
	}
	if inviteRepo.recordCalls != 1 || !inviteRepo.recordArg.Requested {
		t.Fatalf("record invite arg = %#v, want requested join request", inviteRepo.recordArg)
	}
}

func TestChatHideJoinRequestApproveUsesAddHelper(t *testing.T) {
	write := &fakeWriteRepo{participant: participantForMemberTests(10, 2, chat.ChatMemberNormal, chat.ChatMemberStateNormal, nil)}
	inviteRepo := &fakeInviteRepo{requesters: chat.MakeTLRecentChatInviteRequesters(&chat.TLRecentChatInviteRequesters{RecentRequesters: []int64{}}).ToRecentChatInviteRequesters()}
	core := newInviteTestCore(&fakeReadRepo{mutableChat: inviteMutableChat(false)}, write, inviteRepo)
	_, err := core.ChatHideChatJoinRequests(&chat.TLChatHideChatJoinRequests{SelfId: 1, ChatId: 10, Approved: true, UserId: int64Ptr(2)})
	if err != nil {
		t.Fatalf("ChatHideChatJoinRequests error = %v", err)
	}
	if write.addCalls != 1 {
		t.Fatalf("AddChatUser calls = %d, want 1", write.addCalls)
	}
	if inviteRepo.hideCalls != 1 || !inviteRepo.hideArg.Approve {
		t.Fatalf("hide arg = %#v, want approve", inviteRepo.hideArg)
	}
}

func TestChatEditExportedInviteReturnsOldAndNew(t *testing.T) {
	oldInvite := tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{Link: "old", AdminId: 1, Date: 1}).ToExportedChatInvite().Clazz
	newInvite := tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{Link: "new", AdminId: 1, Date: 2}).ToExportedChatInvite().Clazz
	inviteRepo := &fakeInviteRepo{editInvites: []tg.ExportedChatInviteClazz{oldInvite, newInvite}}
	core := newInviteTestCore(&fakeReadRepo{}, &fakeWriteRepo{}, inviteRepo)
	got, err := core.ChatEditExportedChatInvite(&chat.TLChatEditExportedChatInvite{SelfId: 1, ChatId: 10, Link: "old", Revoked: true})
	if err != nil {
		t.Fatalf("ChatEditExportedChatInvite error = %v", err)
	}
	if len(got.Datas) != 2 {
		t.Fatalf("returned invites = %d, want old and new invite", len(got.Datas))
	}
}

func TestTask7HandlersHaveNoPlaceholdersOrTGErrors(t *testing.T) {
	files := []string{
		"chat.exportChatInvite_handler.go",
		"chat.getAdminsWithInvites_handler.go",
		"chat.getExportedChatInvite_handler.go",
		"chat.getExportedChatInvites_handler.go",
		"chat.checkChatInvite_handler.go",
		"chat.importChatInvite_handler.go",
		"chat.importChatInvite2_handler.go",
		"chat.getChatInviteImporters_handler.go",
		"chat.deleteExportedChatInvite_handler.go",
		"chat.deleteRevokedExportedChatInvites_handler.go",
		"chat.editExportedChatInvite_handler.go",
		"chat.getRecentChatInviteRequesters_handler.go",
		"chat.hideChatJoinRequests_handler.go",
	}
	for _, file := range files {
		b, err := os.ReadFile(filepath.Join(".", file))
		if err != nil {
			t.Fatal(err)
		}
		s := string(b)
		if strings.Contains(s, "ErrMethodNotImpl") || strings.Contains(s, "not impl") || strings.Contains(s, "tg.Err") {
			t.Fatalf("%s still contains placeholder or tg.Err", file)
		}
	}
}

func int64Ptr(v int64) *int64 {
	return &v
}
