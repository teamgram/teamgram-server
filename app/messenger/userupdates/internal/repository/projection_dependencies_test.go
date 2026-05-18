package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/envelope"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatprojection "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chatprojection"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	userprojection "github.com/teamgram/teamgram-server/v2/app/service/biz/user/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestEnvelopeBuildProjectsDependencyUsersAndChats(t *testing.T) {
	user111 := tg.MakeTLUser(&tg.TLUser{Id: 111})
	user222 := tg.MakeTLUser(&tg.TLUser{Id: 222})
	chat333 := tg.MakeTLChat(&tg.TLChat{Id: 333, Title: "chat"})
	projector := &fakePeerObjectProjector{
		users: []tg.UserClazz{user111, user222},
		chats: []tg.ChatClazz{chat333},
	}
	updates := []tg.UpdateClazz{
		tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message: tg.MakeTLMessageService(&tg.TLMessageService{
				Id:     10,
				FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 111}),
				PeerId: tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 333}),
				Action: tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
					Title: "chat",
					Users: []int64{222},
				}),
			}),
		}),
	}

	got, err := BuildUpdatesWithDependencies(context.Background(), projector, 111, envelope.Input{
		Mode:    envelope.ModeReceiverStream,
		Updates: updates,
	})
	if err != nil {
		t.Fatalf("BuildUpdatesWithDependencies() error = %v", err)
	}

	if !reflect.DeepEqual(projector.userCalls, [][]int64{{111, 222}}) {
		t.Fatalf("ProjectUsers calls = %#v, want [[111 222]]", projector.userCalls)
	}
	if !reflect.DeepEqual(projector.chatCalls, [][]int64{{333}}) {
		t.Fatalf("ProjectChats calls = %#v, want [[333]]", projector.chatCalls)
	}
	full, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
	if !reflect.DeepEqual(full.Users, []tg.UserClazz{user111, user222}) {
		t.Fatalf("users = %#v, want fake users", full.Users)
	}
	if !reflect.DeepEqual(full.Chats, []tg.ChatClazz{chat333}) {
		t.Fatalf("chats = %#v, want fake chats", full.Chats)
	}
}

func TestEnvelopeBuildRejectsEmptyProjectedDependencyUsers(t *testing.T) {
	projector := &fakePeerObjectProjector{}

	_, err := BuildUpdatesWithDependencies(context.Background(), projector, 111, envelope.Input{
		Mode: envelope.ModeReceiverStream,
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateUserTyping(&tg.TLUpdateUserTyping{UserId: 222}),
		},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("BuildUpdatesWithDependencies() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestEnvelopeBuildRejectsEmptyProjectedDependencyChats(t *testing.T) {
	projector := &fakePeerObjectProjector{}

	_, err := BuildUpdatesWithDependencies(context.Background(), projector, 111, envelope.Input{
		Mode: envelope.ModeReceiverStream,
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateChat(&tg.TLUpdateChat{ChatId: 333}),
		},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("BuildUpdatesWithDependencies() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestEnvelopeBuildMergesCallerSuppliedUsersAndChats(t *testing.T) {
	callerUser := tg.MakeTLUser(&tg.TLUser{Id: 111})
	projectedUser := tg.MakeTLUser(&tg.TLUser{Id: 222})
	projectedChat := tg.MakeTLChat(&tg.TLChat{Id: 333, Title: "projected"})
	callerChat := tg.MakeTLChat(&tg.TLChat{Id: 444, Title: "caller"})
	projector := &fakePeerObjectProjector{
		users: []tg.UserClazz{projectedUser},
		chats: []tg.ChatClazz{projectedChat},
	}

	got, err := BuildUpdatesWithDependencies(context.Background(), projector, 111, envelope.Input{
		Mode: envelope.ModeReceiverStream,
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
				Message: tg.MakeTLMessageService(&tg.TLMessageService{
					Id:     10,
					FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 222}),
					PeerId: tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 333}),
					Action: tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
						Title: "chat",
						Users: []int64{222},
					}),
				}),
			}),
		},
		Users: []tg.UserClazz{callerUser},
		Chats: []tg.ChatClazz{callerChat},
	})
	if err != nil {
		t.Fatalf("BuildUpdatesWithDependencies() error = %v", err)
	}

	full, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
	if !hasUserID(full.Users, 111) || !hasUserID(full.Users, 222) {
		t.Fatalf("users = %#v, want caller and projected users", full.Users)
	}
	if !hasChatID(full.Chats, 333) || !hasChatID(full.Chats, 444) {
		t.Fatalf("chats = %#v, want caller and projected chats", full.Chats)
	}
}

func TestEnvelopeBuildRejectsMissingProjectedDependencyUser(t *testing.T) {
	projector := &fakePeerObjectProjector{
		users: []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 111})},
	}

	_, err := BuildUpdatesWithDependencies(context.Background(), projector, 111, envelope.Input{
		Mode: envelope.ModeReceiverStream,
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
				Message: tg.MakeTLMessageService(&tg.TLMessageService{
					Id:     10,
					FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 111}),
					PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 222}),
				}),
			}),
		},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("BuildUpdatesWithDependencies() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestEnvelopeBuildRejectsMissingProjectedDependencyChat(t *testing.T) {
	projector := &fakePeerObjectProjector{
		chats: []tg.ChatClazz{tg.MakeTLChat(&tg.TLChat{Id: 333})},
	}

	_, err := BuildUpdatesWithDependencies(context.Background(), projector, 111, envelope.Input{
		Mode: envelope.ModeReceiverStream,
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateChat(&tg.TLUpdateChat{ChatId: 333}),
			tg.MakeTLUpdateChat(&tg.TLUpdateChat{ChatId: 444}),
		},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("BuildUpdatesWithDependencies() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestEnvelopeBuildRejectsChannelDependencies(t *testing.T) {
	projector := &fakePeerObjectProjector{
		users: []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 111})},
	}

	_, err := BuildUpdatesWithDependencies(context.Background(), projector, 111, envelope.Input{
		Mode: envelope.ModeReply,
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
				Message: tg.MakeTLMessage(&tg.TLMessage{
					Id:     10,
					FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 111}),
					PeerId: tg.MakeTLPeerChannel(&tg.TLPeerChannel{ChannelId: 444}),
				}),
			}),
		},
		MessageIDByID: map[int32]int64{10: 123456789},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("BuildUpdatesWithDependencies() error = %v, want ErrUserupdatesStorage", err)
	}
	if len(projector.userCalls) != 0 || len(projector.chatCalls) != 0 {
		t.Fatalf("projector calls = users:%#v chats:%#v, want no projection before channel support", projector.userCalls, projector.chatCalls)
	}
}

func TestRepositoryProjectionSkipsEmptyIDs(t *testing.T) {
	repo := NewForTest(nil, nil, "")

	users, err := repo.ProjectUsers(context.Background(), 111, nil)
	if err != nil {
		t.Fatalf("ProjectUsers(empty) error = %v", err)
	}
	if users != nil {
		t.Fatalf("ProjectUsers(empty) = %#v, want nil", users)
	}
	chats, err := repo.ProjectChats(context.Background(), 111, nil)
	if err != nil {
		t.Fatalf("ProjectChats(empty) error = %v", err)
	}
	if chats != nil {
		t.Fatalf("ProjectChats(empty) = %#v, want nil", chats)
	}
}

func TestRepositoryProjectionMissingClientWrapsStorageError(t *testing.T) {
	repo := NewForTest(nil, nil, "")

	if _, err := repo.ProjectUsers(context.Background(), 111, []int64{222}); !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectUsers() error = %v, want ErrUserupdatesStorage", err)
	}
	if _, err := repo.ProjectChats(context.Background(), 111, []int64{333}); !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectChats() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestRepositoryProjectionDownstreamErrorWrapsStorageError(t *testing.T) {
	repo := NewForTest(nil, nil, "")
	repo.SetPeerProjectionClients(
		&fakeUserProjectionClient{err: errors.New("user rpc down")},
		&fakeChatProjectionClient{err: errors.New("chat rpc down")},
	)

	if _, err := repo.ProjectUsers(context.Background(), 111, []int64{222}); !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectUsers() error = %v, want ErrUserupdatesStorage", err)
	}
	if _, err := repo.ProjectChats(context.Background(), 111, []int64{333}); !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectChats() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestRepositoryProjectUsersRejectsEmptyViewerUsers(t *testing.T) {
	userClient := &fakeUserProjectionClient{
		bundle: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
			ViewerUsers: []userpb.ViewerUsersClazz{
				userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 111}),
			},
		}),
	}
	repo := NewForTest(nil, nil, "")
	repo.SetPeerProjectionClients(userClient, nil)

	if _, err := repo.ProjectUsers(context.Background(), 111, []int64{222}); !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectUsers() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestRepositoryProjectUsersWrapsPublicProjectionSentinel(t *testing.T) {
	repo := NewForTest(nil, nil, "")
	userClient := &fakeUserProjectionClient{
		bundle: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
			ViewerUsers: []userpb.ViewerUsersClazz{
				userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 9999}),
			},
		}),
	}
	repo.SetPeerProjectionClients(userClient, nil)

	_, err := repo.ProjectUsers(context.Background(), 111, []int64{222})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectUsers() error = %v, want ErrUserupdatesStorage", err)
	}
	if !errors.Is(err, userprojection.ErrViewerProjectionMissing) {
		t.Fatalf("ProjectUsers() error = %v, want ErrViewerProjectionMissing preserved", err)
	}
}

func TestRepositoryProjectChatsRejectsEmptyDownstreamList(t *testing.T) {
	repo := NewForTest(nil, nil, "")
	repo.SetPeerProjectionClients(nil, &fakeChatProjectionClient{
		bundle: chatProjectionBundle(111),
	})

	if _, err := repo.ProjectChats(context.Background(), 111, []int64{333}); !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectChats() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestRepositoryProjectChatsWrapsDateOverflowSentinel(t *testing.T) {
	chatClient := &fakeChatProjectionClient{err: chatprojection.ErrChatDateOverflow}
	repo := NewForTest(nil, nil, "")
	repo.SetPeerProjectionClients(nil, chatClient)

	_, err := repo.ProjectChats(context.Background(), 111, []int64{333})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectChats() error = %v, want ErrUserupdatesStorage", err)
	}
	if !errors.Is(err, chatprojection.ErrChatDateOverflow) {
		t.Fatalf("ProjectChats() error = %v, want ErrChatDateOverflow preserved", err)
	}
}

func TestRepositoryProjectUsersCallsProjectionBundleForViewer(t *testing.T) {
	user222 := tg.MakeTLUser(&tg.TLUser{Id: 222})
	userClient := &fakeUserProjectionClient{
		bundle: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
			ViewerUsers: []userpb.ViewerUsersClazz{
				userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{
					ViewerUserId: 111,
					Users:        []tg.UserClazz{user222},
				}),
			},
		}),
	}
	repo := NewForTest(nil, nil, "")
	repo.SetPeerProjectionClients(userClient, nil)

	got, err := repo.ProjectUsers(context.Background(), 111, []int64{222})
	if err != nil {
		t.Fatalf("ProjectUsers() error = %v", err)
	}

	if !reflect.DeepEqual(userClient.viewerUserIDs, []int64{111}) {
		t.Fatalf("ViewerUserIds = %#v, want [111]", userClient.viewerUserIDs)
	}
	if !reflect.DeepEqual(userClient.targetUserIDs, []int64{222}) {
		t.Fatalf("TargetUserIds = %#v, want [222]", userClient.targetUserIDs)
	}
	if !reflect.DeepEqual(got, []tg.UserClazz{user222}) {
		t.Fatalf("ProjectUsers() = %#v, want projected user", got)
	}
}

func TestRepositoryProjectChatsCallsProjectionBundle(t *testing.T) {
	projected := tg.MakeTLChat(&tg.TLChat{Id: 333, Title: "chat"})
	chatClient := &fakeChatProjectionClient{
		bundle: chatProjectionBundle(111, projected),
	}
	repo := NewForTest(nil, nil, "")
	repo.SetPeerProjectionClients(nil, chatClient)

	got, err := repo.ProjectChats(context.Background(), 111, []int64{333})
	if err != nil {
		t.Fatalf("ProjectChats() error = %v", err)
	}

	if !reflect.DeepEqual(chatClient.viewerUserIDs, []int64{111}) {
		t.Fatalf("ViewerUserIds = %#v, want [111]", chatClient.viewerUserIDs)
	}
	if !reflect.DeepEqual(chatClient.targetChatIDs, []int64{333}) {
		t.Fatalf("TargetChatIds = %#v, want [333]", chatClient.targetChatIDs)
	}
	if !reflect.DeepEqual(got, []tg.ChatClazz{projected}) {
		t.Fatalf("ProjectChats() = %#v, want projected chat", got)
	}
}

func TestRepositoryProjectChatsWrapsPublicProjectionSentinel(t *testing.T) {
	repo := NewForTest(nil, nil, "")
	chatClient := &fakeChatProjectionClient{
		bundle: chatProjectionBundle(9999),
	}
	repo.SetPeerProjectionClients(nil, chatClient)

	_, err := repo.ProjectChats(context.Background(), 111, []int64{333})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectChats() error = %v, want ErrUserupdatesStorage", err)
	}
	if !errors.Is(err, chatprojection.ErrViewerProjectionMissing) {
		t.Fatalf("ProjectChats() error = %v, want ErrViewerProjectionMissing preserved", err)
	}
}

type fakePeerObjectProjector struct {
	userCalls [][]int64
	chatCalls [][]int64
	users     []tg.UserClazz
	chats     []tg.ChatClazz
}

func (p *fakePeerObjectProjector) ProjectUsers(_ context.Context, _ int64, ids []int64) ([]tg.UserClazz, error) {
	p.userCalls = append(p.userCalls, append([]int64(nil), ids...))
	return p.users, nil
}

func (p *fakePeerObjectProjector) ProjectChats(_ context.Context, _ int64, ids []int64) ([]tg.ChatClazz, error) {
	p.chatCalls = append(p.chatCalls, append([]int64(nil), ids...))
	return p.chats, nil
}

type fakeUserProjectionClient struct {
	viewerUserIDs []int64
	targetUserIDs []int64
	bundle        *userpb.UserProjectionBundle
	err           error
}

func (c *fakeUserProjectionClient) UserGetUserProjectionBundle(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	if in != nil {
		c.viewerUserIDs = append([]int64(nil), in.ViewerUserIds...)
		c.targetUserIDs = append([]int64(nil), in.TargetUserIds...)
	}
	return c.bundle, c.err
}

type fakeChatProjectionClient struct {
	viewerUserIDs []int64
	targetChatIDs []int64
	bundle        *chatpb.ChatProjectionBundle
	err           error
}

func (c *fakeChatProjectionClient) ChatGetChatProjectionBundle(_ context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error) {
	if in != nil {
		c.viewerUserIDs = append([]int64(nil), in.ViewerUserIds...)
		c.targetChatIDs = append([]int64(nil), in.TargetChatIds...)
	}
	return c.bundle, c.err
}

func chatProjectionBundle(viewerUserID int64, chats ...tg.ChatClazz) *chatpb.ChatProjectionBundle {
	return chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{
				ViewerUserId: viewerUserID,
				Chats:        chats,
			}),
		},
	}).ToChatProjectionBundle()
}

func hasUserID(users []tg.UserClazz, id int64) bool {
	for _, user := range users {
		if userID(user) == id {
			return true
		}
	}
	return false
}

func hasChatID(chats []tg.ChatClazz, id int64) bool {
	for _, chat := range chats {
		if chatID(chat) == id {
			return true
		}
	}
	return false
}
