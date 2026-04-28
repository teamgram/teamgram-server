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
	invite             *model.ChatInvites
	inviteErr          error
	usage              int32
	usageErr           error
	requesters         *chat.RecentChatInviteRequesters
	requestersErr      error
	pendingRequests    []repository.JoinRequest
	recordCalls        int
	recordArg          repository.InviteParticipantArg
	hideCalls          int
	hideArg            repository.HideJoinRequestsArg
	createCalls        int
	deleteRevokedCalls int
	editInvites        []tg.ExportedChatInviteClazz
	editArg            repository.EditExportedChatInviteArg
	createInvite       *tg.ExportedChatInvite
	exportedInvite     *tg.ExportedChatInvite
	exportedInvites    []tg.ExportedChatInviteClazz
	importers          []tg.ChatInviteImporterClazz
	err                error
}

func (f *fakeInviteRepo) CreateExportedChatInvite(ctx context.Context, arg repository.ExportChatInviteArg) (*tg.ExportedChatInvite, error) {
	f.createCalls++
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
	f.deleteRevokedCalls++
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

func (f *fakeInviteRepo) GetPendingJoinRequests(ctx context.Context, chatID int64, link *string) ([]repository.JoinRequest, error) {
	return f.pendingRequests, f.err
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

func TestChatHideJoinRequestApproveUsesAtomicAddMutation(t *testing.T) {
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
	if !write.addArg.ApproveJoinRequest || write.addArg.ApprovedBy != 1 {
		t.Fatalf("AddChatUser arg = %#v, want atomic join request approval", write.addArg)
	}
	if inviteRepo.hideCalls != 0 {
		t.Fatalf("HideChatJoinRequest calls = %d, want 0 on approve path", inviteRepo.hideCalls)
	}
}

func TestChatHideJoinRequestsAllApproveUsesAtomicAddMutation(t *testing.T) {
	write := &fakeWriteRepo{participant: participantForMemberTests(10, 2, chat.ChatMemberNormal, chat.ChatMemberStateNormal, nil)}
	inviteRepo := &fakeInviteRepo{
		pendingRequests: []repository.JoinRequest{
			{ChatID: 10, Link: "hash", UserID: 2},
			{ChatID: 10, Link: "hash", UserID: 3},
		},
		requesters: chat.MakeTLRecentChatInviteRequesters(&chat.TLRecentChatInviteRequesters{RecentRequesters: []int64{}}).ToRecentChatInviteRequesters(),
	}
	core := newInviteTestCore(&fakeReadRepo{mutableChat: inviteMutableChat(false)}, write, inviteRepo)
	_, err := core.ChatHideChatJoinRequests(&chat.TLChatHideChatJoinRequests{SelfId: 1, ChatId: 10, Approved: true})
	if err != nil {
		t.Fatalf("ChatHideChatJoinRequests all approve error = %v", err)
	}
	if write.addCalls != 2 {
		t.Fatalf("AddChatUser calls = %d, want 2", write.addCalls)
	}
	if !write.addArg.ApproveJoinRequest || write.addArg.ApprovedBy != 1 {
		t.Fatalf("AddChatUser arg = %#v, want atomic join request approval", write.addArg)
	}
	if inviteRepo.hideCalls != 0 {
		t.Fatalf("HideChatJoinRequest calls = %d, want 0 on all approve path", inviteRepo.hideCalls)
	}
}

func TestChatHideJoinRequestsAllDeleteSkipsAddHelper(t *testing.T) {
	write := &fakeWriteRepo{}
	inviteRepo := &fakeInviteRepo{
		pendingRequests: []repository.JoinRequest{{ChatID: 10, Link: "hash", UserID: 2}},
		requesters:      chat.MakeTLRecentChatInviteRequesters(&chat.TLRecentChatInviteRequesters{RecentRequesters: []int64{}}).ToRecentChatInviteRequesters(),
	}
	core := newInviteTestCore(&fakeReadRepo{mutableChat: inviteMutableChat(false)}, write, inviteRepo)
	_, err := core.ChatHideChatJoinRequests(&chat.TLChatHideChatJoinRequests{SelfId: 1, ChatId: 10})
	if err != nil {
		t.Fatalf("ChatHideChatJoinRequests all delete error = %v", err)
	}
	if write.addCalls != 0 {
		t.Fatalf("AddChatUser calls = %d, want 0", write.addCalls)
	}
	if inviteRepo.hideCalls != 1 || inviteRepo.hideArg.Approve {
		t.Fatalf("hide calls=%d arg=%#v, want one delete", inviteRepo.hideCalls, inviteRepo.hideArg)
	}
}

func TestChatImportInviteRecordsParticipantAtomically(t *testing.T) {
	write := &fakeWriteRepo{participant: participantForMemberTests(10, 2, chat.ChatMemberNormal, chat.ChatMemberStateNormal, nil)}
	inviteRepo := &fakeInviteRepo{invite: baseInviteRow(false)}
	core := newInviteTestCore(&fakeReadRepo{mutableChat: inviteMutableChat(false)}, write, inviteRepo)
	_, err := core.ChatImportChatInvite(&chat.TLChatImportChatInvite{SelfId: 2, Hash: "hash"})
	if err != nil {
		t.Fatalf("ChatImportChatInvite error = %v", err)
	}
	if write.addCalls != 1 {
		t.Fatalf("AddChatUser calls = %d, want 1", write.addCalls)
	}
	if !write.addArg.RecordInviteParticipant || write.addArg.InviteLink != "hash" {
		t.Fatalf("AddChatUser arg = %#v, want invite participant record in transaction", write.addArg)
	}
	if inviteRepo.recordCalls != 0 {
		t.Fatalf("RecordInviteParticipant calls = %d, want 0 after add", inviteRepo.recordCalls)
	}
}

func TestChatExportInviteRequiresInvitePermission(t *testing.T) {
	inviteRepo := &fakeInviteRepo{}
	core := newInviteTestCore(&fakeReadRepo{mutableChat: mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 2, chat.ChatMemberNormal, chat.ChatMemberStateNormal, nil),
	)}, &fakeWriteRepo{}, inviteRepo)
	_, err := core.ChatExportChatInvite(&chat.TLChatExportChatInvite{ChatId: 10, AdminId: 2})
	if !errors.Is(err, chat.ErrChatAdminRequired) {
		t.Fatalf("ChatExportChatInvite error = %v, want ErrChatAdminRequired", err)
	}
	if inviteRepo.createCalls != 0 {
		t.Fatalf("CreateExportedChatInvite calls = %d, want 0", inviteRepo.createCalls)
	}
}

func TestChatDeleteRevokedInvitesRequiresInvitePermission(t *testing.T) {
	inviteRepo := &fakeInviteRepo{}
	core := newInviteTestCore(&fakeReadRepo{mutableChat: mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 2, chat.ChatMemberNormal, chat.ChatMemberStateNormal, nil),
	)}, &fakeWriteRepo{}, inviteRepo)
	_, err := core.ChatDeleteRevokedExportedChatInvites(&chat.TLChatDeleteRevokedExportedChatInvites{SelfId: 2, ChatId: 10, AdminId: 2})
	if !errors.Is(err, chat.ErrChatAdminRequired) {
		t.Fatalf("ChatDeleteRevokedExportedChatInvites error = %v, want ErrChatAdminRequired", err)
	}
	if inviteRepo.deleteRevokedCalls != 0 {
		t.Fatalf("DeleteRevokedExportedChatInvites calls = %d, want 0", inviteRepo.deleteRevokedCalls)
	}
}

func TestChatEditExportedInviteReturnsOldAndNew(t *testing.T) {
	oldInvite := tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{Link: "old", AdminId: 1, Date: 1}).ToExportedChatInvite().Clazz
	newInvite := tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{Link: "new", AdminId: 1, Date: 2}).ToExportedChatInvite().Clazz
	inviteRepo := &fakeInviteRepo{editInvites: []tg.ExportedChatInviteClazz{oldInvite, newInvite}}
	core := newInviteTestCore(&fakeReadRepo{mutableChat: mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil),
	)}, &fakeWriteRepo{}, inviteRepo)
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
