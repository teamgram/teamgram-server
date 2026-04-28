package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeWriteRepo struct {
	mutableChat       *tg.MutableChat
	err               error
	createArg         repository.CreateChatArg
	addArg            repository.AddChatUserArg
	deleteChatID      int64
	deleteChatUserArg repository.DeleteChatUserArg
	migratedArg       repository.MigratedToChannelArg
	createCalls       int
	addCalls          int
	deleteChatCalls   int
	deleteUserCalls   int
	migratedCalls     int
}

func (f *fakeWriteRepo) CreateChat(ctx context.Context, arg repository.CreateChatArg) (*tg.MutableChat, error) {
	f.createCalls++
	f.createArg = arg
	return f.mutableChat, f.err
}

func (f *fakeWriteRepo) DeleteChat(ctx context.Context, chatID int64) error {
	f.deleteChatCalls++
	f.deleteChatID = chatID
	return f.err
}

func (f *fakeWriteRepo) AddChatUser(ctx context.Context, arg repository.AddChatUserArg) (*tg.MutableChat, error) {
	f.addCalls++
	f.addArg = arg
	return f.mutableChat, f.err
}

func (f *fakeWriteRepo) DeleteChatUser(ctx context.Context, arg repository.DeleteChatUserArg) (*tg.MutableChat, error) {
	f.deleteUserCalls++
	f.deleteChatUserArg = arg
	return f.mutableChat, f.err
}

func (f *fakeWriteRepo) MigratedToChannel(ctx context.Context, arg repository.MigratedToChannelArg) (*tg.MutableChat, error) {
	f.migratedCalls++
	f.migratedArg = arg
	return f.mutableChat, f.err
}

func newWriteTestCore(read *fakeReadRepo, write *fakeWriteRepo) *ChatCore {
	return &ChatCore{
		ctx:       context.Background(),
		readRepo:  read,
		writeRepo: write,
	}
}

func mutableChatForMemberTests(chatID, creatorID int64, participants ...*tg.ImmutableChatParticipant) *tg.MutableChat {
	return tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{
			Id:                  chatID,
			Creator:             creatorID,
			Title:               "chat",
			ParticipantsCount:   int32(len(participants)),
			DefaultBannedRights: tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{}).ToChatBannedRights(),
		}).ToImmutableChat(),
		ChatParticipants: participants,
	}).ToMutableChat()
}

func participantForMemberTests(chatID, userID int64, typ, state int32, rights tg.ChatAdminRightsClazz) *tg.ImmutableChatParticipant {
	return tg.MakeTLImmutableChatParticipant(&tg.TLImmutableChatParticipant{
		Id:              userID,
		ChatId:          chatID,
		UserId:          userID,
		ParticipantType: typ,
		State:           state,
		AdminRights:     rights,
	}).ToImmutableChatParticipant()
}

func TestCreateChatReturnsTypedFloodError(t *testing.T) {
	flood := chat.NewCreateChatFloodError(17)
	write := &fakeWriteRepo{err: flood}
	core := newWriteTestCore(&fakeReadRepo{}, write)

	_, err := core.ChatCreateChat2(&chat.TLChatCreateChat2{CreatorId: 1, UserIdList: []int64{2}, Title: "team"})
	if !errors.Is(err, chat.ErrCreateChatFlood) {
		t.Fatalf("ChatCreateChat2 error = %v, want ErrCreateChatFlood", err)
	}
	var typed *chat.CreateChatFloodError
	if !errors.As(err, &typed) || typed.WaitSeconds != 17 {
		t.Fatalf("typed flood = %#v, want wait 17", typed)
	}
	if write.createCalls != 1 {
		t.Fatalf("CreateChat calls = %d, want 1", write.createCalls)
	}
}

func TestAddChatUserBranches(t *testing.T) {
	adminNoInvite := tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{BanUsers: true}).ToChatAdminRights()

	tests := []struct {
		name      string
		setup     func() *tg.MutableChat
		inviterID int64
		want      error
	}{
		{
			name: "already normal",
			setup: func() *tg.MutableChat {
				return mutableChatForMemberTests(10, 1,
					participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil),
					participantForMemberTests(10, 2, chat.ChatMemberNormal, chat.ChatMemberStateNormal, nil))
			},
			inviterID: 1,
			want:      chat.ErrUserAlreadyParticipant,
		},
		{
			name: "migrated",
			setup: func() *tg.MutableChat {
				m := mutableChatForMemberTests(10, 1,
					participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil))
				m.Chat.Deactivated = true
				m.Chat.MigratedTo = tg.MakeTLInputChannel(&tg.TLInputChannel{ChannelId: 99, AccessHash: 100})
				return m
			},
			inviterID: 3,
			want:      chat.ErrChatMigrated,
		},
		{
			name: "non-admin invite",
			setup: func() *tg.MutableChat {
				m := mutableChatForMemberTests(10, 1,
					participantForMemberTests(10, 3, chat.ChatMemberAdmin, chat.ChatMemberStateNormal, adminNoInvite))
				m.Chat.DefaultBannedRights = tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{InviteUsers: true}).ToChatBannedRights()
				return m
			},
			inviterID: 3,
			want:      chat.ErrChatAdminRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			write := &fakeWriteRepo{mutableChat: tt.setup()}
			core := newWriteTestCore(&fakeReadRepo{mutableChat: write.mutableChat}, write)
			_, err := core.addChatUser(context.Background(), addChatUserArg{chatID: 10, inviterID: tt.inviterID, userID: 2})
			if !errors.Is(err, tt.want) {
				t.Fatalf("addChatUser error = %v, want %v", err, tt.want)
			}
			if write.addCalls != 0 {
				t.Fatalf("AddChatUser calls = %d, want 0 for rejected branch", write.addCalls)
			}
		})
	}
}

func TestAddChatUserPreservesCreatorParticipantType(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateLeft, nil))
	write := &fakeWriteRepo{mutableChat: m}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.addChatUser(context.Background(), addChatUserArg{chatID: 10, userID: 1})
	if err != nil {
		t.Fatalf("addChatUser error: %v", err)
	}
	if write.addCalls != 1 {
		t.Fatalf("AddChatUser calls = %d, want 1", write.addCalls)
	}
	if write.addArg.ParticipantType != chat.ChatMemberCreator {
		t.Fatalf("ParticipantType = %d, want creator", write.addArg.ParticipantType)
	}
}

func TestDeleteChatUserProtectsCreator(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil),
		participantForMemberTests(10, 3, chat.ChatMemberAdmin, chat.ChatMemberStateNormal,
			tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{BanUsers: true}).ToChatAdminRights()))
	write := &fakeWriteRepo{mutableChat: m}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.deleteChatUser(context.Background(), deleteChatUserArg{chatID: 10, operatorID: 3, deleteUserID: 1})
	if !errors.Is(err, chat.ErrChatAdminRequired) {
		t.Fatalf("deleteChatUser error = %v, want ErrChatAdminRequired", err)
	}
	if write.deleteUserCalls != 0 {
		t.Fatalf("DeleteChatUser calls = %d, want 0 for creator protection", write.deleteUserCalls)
	}
}

func TestDeleteChatUserProtectsAdminTarget(t *testing.T) {
	adminRights := tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{BanUsers: true}).ToChatAdminRights()
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil),
		participantForMemberTests(10, 3, chat.ChatMemberAdmin, chat.ChatMemberStateNormal, adminRights),
		participantForMemberTests(10, 4, chat.ChatMemberAdmin, chat.ChatMemberStateNormal, nil))
	write := &fakeWriteRepo{mutableChat: m}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.deleteChatUser(context.Background(), deleteChatUserArg{chatID: 10, operatorID: 3, deleteUserID: 4})
	if !errors.Is(err, chat.ErrChatAdminRequired) {
		t.Fatalf("deleteChatUser error = %v, want ErrChatAdminRequired", err)
	}
	if write.deleteUserCalls != 0 {
		t.Fatalf("DeleteChatUser calls = %d, want 0 for admin protection", write.deleteUserCalls)
	}
}

func TestDeleteChatRequiresCreator(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil),
		participantForMemberTests(10, 2, chat.ChatMemberNormal, chat.ChatMemberStateNormal, nil))
	write := &fakeWriteRepo{mutableChat: m}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatDeleteChat(&chat.TLChatDeleteChat{ChatId: 10, OperatorId: 2})
	if !errors.Is(err, chat.ErrChatAdminRequired) {
		t.Fatalf("ChatDeleteChat error = %v, want ErrChatAdminRequired", err)
	}
	if write.deleteChatCalls != 0 {
		t.Fatalf("DeleteChat calls = %d, want 0 for rejected branch", write.deleteChatCalls)
	}
}

func TestMigratedToChannelPassesChatIDToRepository(t *testing.T) {
	m := mutableChatForMemberTests(10, 1)
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{}, write)

	got, err := core.ChatMigratedToChannel(&chat.TLChatMigratedToChannel{Chat: m, Id: 99, AccessHash: 100})
	if err != nil {
		t.Fatalf("ChatMigratedToChannel error: %v", err)
	}
	if got != tg.BoolTrue {
		t.Fatalf("ChatMigratedToChannel = %v, want BoolTrue", got)
	}
	if write.migratedCalls != 1 || write.migratedArg.ChatID != 10 || write.migratedArg.ChannelID != 99 || write.migratedArg.AccessHash != 100 {
		t.Fatalf("migrated call = %#v calls=%d", write.migratedArg, write.migratedCalls)
	}
}
