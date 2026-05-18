package chatprojection

import (
	"context"
	"errors"
	"reflect"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeChatClient struct {
	in  *chatpb.TLChatGetChatProjectionBundle
	out *chatpb.ChatProjectionBundle
	err error
}

func (f *fakeChatClient) ChatGetChatProjectionBundle(_ context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error) {
	f.in = in
	return f.out, f.err
}

var _ MissingPolicy = MissingExplicitInput

func TestProjectChatsEmptyInputReturnsNil(t *testing.T) {
	client := &fakeChatClient{}

	got, err := ProjectChats(context.Background(), client, 1001, nil, Options{Missing: MissingExplicitInput})
	if err != nil {
		t.Fatalf("ProjectChats(empty) error = %v", err)
	}
	if got != nil {
		t.Fatalf("ProjectChats(empty) = %#v, want nil", got)
	}
	if client.in != nil {
		t.Fatalf("client called for empty input: %#v", client.in)
	}
}

func TestProjectChatsReturnsViewerVectorAndPreservesRequestOrder(t *testing.T) {
	first := tg.MakeTLChat(&tg.TLChat{Id: 3001, Title: "first"})
	second := tg.MakeTLChat(&tg.TLChat{Id: 3002, Title: "second"})
	client := &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{
				ViewerUserId: 1001,
				Chats:        []tg.ChatClazz{first, second},
			}),
		},
	}).ToChatProjectionBundle()}

	got, err := ProjectChats(context.Background(), client, 1001, []int64{3001, 3002}, Options{Missing: MissingExplicitInput})
	if err != nil {
		t.Fatalf("ProjectChats() error = %v", err)
	}
	if !reflect.DeepEqual(client.in.ViewerUserIds, []int64{1001}) {
		t.Fatalf("ViewerUserIds = %#v, want [1001]", client.in.ViewerUserIds)
	}
	if !reflect.DeepEqual(client.in.TargetChatIds, []int64{3001, 3002}) {
		t.Fatalf("TargetChatIds = %#v, want [3001 3002]", client.in.TargetChatIds)
	}
	if !reflect.DeepEqual(got, []tg.ChatClazz{first, second}) {
		t.Fatalf("ProjectChats() = %#v, want projected chats", got)
	}
}

func TestProjectChatsPreservesDuplicateTargetsForOwnerRPC(t *testing.T) {
	client := &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 1001}),
		},
	}).ToChatProjectionBundle()}

	_, err := ProjectChats(context.Background(), client, 1001, []int64{3001, 3001}, Options{Missing: MissingStoredReference})
	if err != nil {
		t.Fatalf("ProjectChats() error = %v", err)
	}
	if !reflect.DeepEqual(client.in.TargetChatIds, []int64{3001, 3001}) {
		t.Fatalf("TargetChatIds = %#v, want duplicate targets preserved", client.in.TargetChatIds)
	}
}

func TestProjectChatsMapsErrors(t *testing.T) {
	tests := []struct {
		name    string
		client  Client
		opts    Options
		wantErr error
	}{
		{
			name:    "nil client",
			client:  nil,
			opts:    Options{Missing: MissingExplicitInput},
			wantErr: ErrClientNotConfigured,
		},
		{
			name: "invalid request",
			client: &fakeChatClient{
				err: chatpb.ErrChatInvalidArgument,
			},
			opts:    Options{Missing: MissingExplicitInput},
			wantErr: ErrInvalidRequest,
		},
		{
			name:    "nil bundle",
			client:  &fakeChatClient{},
			opts:    Options{Missing: MissingExplicitInput},
			wantErr: ErrNilBundle,
		},
		{
			name: "explicit missing target",
			client: &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
				MissingChatIds: []int64{3001},
			}).ToChatProjectionBundle()},
			opts:    Options{Missing: MissingExplicitInput},
			wantErr: ErrExplicitChatMissing,
		},
		{
			name: "missing viewer",
			client: &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
				ViewerChats: []chatpb.ViewerChatsClazz{
					chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 9999}),
				},
			}).ToChatProjectionBundle()},
			opts:    Options{Missing: MissingStoredReference},
			wantErr: ErrViewerProjectionMissing,
		},
		{
			name: "empty viewer with require non empty",
			client: &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
				ViewerChats: []chatpb.ViewerChatsClazz{
					chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 1001}),
				},
			}).ToChatProjectionBundle()},
			opts:    Options{Missing: MissingStoredReference, RequireNonEmpty: true},
			wantErr: ErrViewerProjectionEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ProjectChats(context.Background(), tt.client, 1001, []int64{3001}, tt.opts)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ProjectChats() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectChatsStoredReferenceMissingReturnsProjectedSubset(t *testing.T) {
	projected := tg.MakeTLChat(&tg.TLChat{Id: 3001})
	client := &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		MissingChatIds: []int64{3002},
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{
				ViewerUserId: 1001,
				Chats:        []tg.ChatClazz{projected},
			}),
		},
	}).ToChatProjectionBundle()}

	got, err := ProjectChats(context.Background(), client, 1001, []int64{3001, 3002}, Options{Missing: MissingStoredReference})
	if err != nil {
		t.Fatalf("ProjectChats() error = %v", err)
	}
	if !reflect.DeepEqual(got, []tg.ChatClazz{projected}) {
		t.Fatalf("ProjectChats() = %#v, want projected subset", got)
	}
}

func TestProjectMutableChatMapsFieldsIncludingPhoto(t *testing.T) {
	photo := tg.MakeTLPhoto(&tg.TLPhoto{Id: 7001, DcId: 4})
	mutable := tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{
			Id:                  3001,
			Creator:             1001,
			Title:               "team",
			Photo:               photo,
			Deactivated:         true,
			CallActive:          true,
			CallNotEmpty:        true,
			Noforwards:          true,
			ParticipantsCount:   2,
			Date:                1234,
			Version:             5,
			MigratedTo:          tg.MakeTLInputChannel(&tg.TLInputChannel{ChannelId: 9001}),
			DefaultBannedRights: tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendMessages: true}),
		}),
	})

	got, err := ProjectMutableChat(mutable, 1001)
	if err != nil {
		t.Fatalf("ProjectMutableChat() error = %v", err)
	}
	chat, ok := got.(*tg.TLChat)
	if !ok {
		t.Fatalf("ProjectMutableChat() = %T, want *tg.TLChat", got)
	}

	if !chat.Creator || chat.Left || !chat.Deactivated || !chat.CallActive || !chat.CallNotEmpty || !chat.Noforwards {
		t.Fatalf("flags not projected: %+v", chat)
	}
	if chat.Id != 3001 || chat.Title != "team" || chat.ParticipantsCount != 2 || chat.Date != 1234 || chat.Version != 5 {
		t.Fatalf("fields not projected: %+v", chat)
	}
	if chat.MigratedTo == nil || chat.DefaultBannedRights == nil {
		t.Fatalf("nested fields not projected: %+v", chat)
	}
	chatPhoto, ok := chat.Photo.(*tg.TLChatPhoto)
	if !ok || chatPhoto.PhotoId != 7001 || chatPhoto.DcId != 4 {
		t.Fatalf("Photo = %#v, want chatPhoto from TLPhoto", chat.Photo)
	}
}

func TestProjectMutableChatNoContentBehavior(t *testing.T) {
	mutable := tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{
			Id:    3001,
			Title: "team",
			Date:  1234,
		}),
	})

	got, err := ProjectMutableChat(mutable, 1001)
	if err != nil {
		t.Fatalf("ProjectMutableChat() error = %v", err)
	}
	if chat := got.(*tg.TLChat); chat.Left || chat.Deactivated {
		t.Fatalf("unexpected non-canonical flags projected: %+v", chat)
	}

	got, err = ProjectMutableChat(nil, 1001)
	if err != nil || got != nil {
		t.Fatalf("ProjectMutableChat(nil) = (%#v, %v), want nil nil", got, err)
	}
	got, err = ProjectMutableChat(tg.MakeTLMutableChat(&tg.TLMutableChat{}), 1001)
	if err != nil || got != nil {
		t.Fatalf("ProjectMutableChat(nil chat) = (%#v, %v), want nil nil", got, err)
	}
}

func TestProjectMutableChatListSkipsNoContent(t *testing.T) {
	projectable := tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{Id: 3001, Title: "team", Date: 1}),
	})

	got, err := ProjectMutableChatList([]tg.MutableChatClazz{
		nil,
		tg.MakeTLMutableChat(&tg.TLMutableChat{}),
		projectable,
	}, 1001)
	if err != nil {
		t.Fatalf("ProjectMutableChatList() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("ProjectMutableChatList() len = %d, want 1", len(got))
	}
	if chat := got[0].(*tg.TLChat); chat.Id != 3001 {
		t.Fatalf("ProjectMutableChatList()[0] = %+v, want chat 3001", chat)
	}
}

func TestInt32ChatIDOverflow(t *testing.T) {
	if got, err := Int32ChatID(3001); err != nil || got != 3001 {
		t.Fatalf("Int32ChatID(3001) = (%d, %v), want 3001 nil", got, err)
	}
	_, err := Int32ChatID(int64(maxInt32) + 1)
	if !errors.Is(err, ErrChatIDOverflow) {
		t.Fatalf("Int32ChatID(overflow) error = %v, want ErrChatIDOverflow", err)
	}
}

func TestProjectMutableChatPhotoFallbackAndDateOverflow(t *testing.T) {
	chatWithEmptyPhoto := tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{Id: 3001, Photo: tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{}), Date: 10}),
	})
	got, err := ProjectMutableChat(chatWithEmptyPhoto, 1001)
	if err != nil {
		t.Fatalf("ProjectMutableChat(photo empty) error = %v", err)
	}
	if _, ok := got.(*tg.TLChat).Photo.(*tg.TLChatPhotoEmpty); !ok {
		t.Fatalf("Photo = %T, want *tg.TLChatPhotoEmpty", got.(*tg.TLChat).Photo)
	}

	overflow := tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{Id: 3002, Date: int64(maxInt32) + 1}),
	})
	_, err = ProjectMutableChat(overflow, 1001)
	if !errors.Is(err, ErrChatDateOverflow) {
		t.Fatalf("ProjectMutableChat(overflow) error = %v, want ErrChatDateOverflow", err)
	}
}

func TestProjectMutableChatForViewersOrderingAndShape(t *testing.T) {
	first := tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{Id: 3001, Title: "first", Date: 1}),
	})
	second := tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{Id: 3002, Title: "second", Date: 2}),
	})

	got, err := ProjectMutableChatForViewers([]tg.MutableChatClazz{first, second}, []int64{1001, 1002})
	if err != nil {
		t.Fatalf("ProjectMutableChatForViewers() error = %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("ProjectMutableChatForViewers() len = %d, want 2", len(got))
	}

	firstViewer := got[0]
	secondViewer := got[1]
	if firstViewer.ViewerUserId != 1001 || secondViewer.ViewerUserId != 1002 {
		t.Fatalf("viewer order = [%d %d], want [1001 1002]", firstViewer.ViewerUserId, secondViewer.ViewerUserId)
	}
	if !reflect.DeepEqual(chatIDs(firstViewer.Chats), []int64{3001, 3002}) {
		t.Fatalf("first viewer chats = %#v, want [3001 3002]", chatIDs(firstViewer.Chats))
	}
	if firstViewer.Chats[0].(*tg.TLChat).Left || firstViewer.Chats[1].(*tg.TLChat).Left {
		t.Fatalf("first viewer left flags = [%t %t], want canonical zero values", firstViewer.Chats[0].(*tg.TLChat).Left, firstViewer.Chats[1].(*tg.TLChat).Left)
	}
	if secondViewer.Chats[0].(*tg.TLChat).Left || secondViewer.Chats[1].(*tg.TLChat).Left {
		t.Fatalf("second viewer left flags = [%t %t], want canonical zero values", secondViewer.Chats[0].(*tg.TLChat).Left, secondViewer.Chats[1].(*tg.TLChat).Left)
	}
}

func TestFillDifferenceChatsReplacesLegacyChats(t *testing.T) {
	projected := tg.MakeTLChat(&tg.TLChat{Id: 3001})
	client := &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{
				ViewerUserId: 1001,
				Chats:        []tg.ChatClazz{projected},
			}),
		},
	}).ToChatProjectionBundle()}
	diff := tg.MakeTLUpdatesDifference(&tg.TLUpdatesDifference{
		NewMessages: []tg.MessageClazz{tg.MakeTLMessage(&tg.TLMessage{
			PeerId: tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 3001}),
		})},
		Chats: []tg.ChatClazz{tg.MakeTLChat(&tg.TLChat{Id: 9999})},
	}).ToUpdatesDifference()

	err := FillDifferenceChats(context.Background(), client, 1001, diff, Options{Missing: MissingStoredReference})
	if err != nil {
		t.Fatalf("FillDifferenceChats() error = %v", err)
	}
	if !reflect.DeepEqual(client.in.TargetChatIds, []int64{3001}) {
		t.Fatalf("TargetChatIds = %#v, want [3001]", client.in.TargetChatIds)
	}
	full, ok := diff.ToUpdatesDifference()
	if !ok {
		t.Fatalf("diff = %s, want updates.difference", diff.ClazzName())
	}
	if len(full.Chats) != 1 || full.Chats[0] != projected {
		t.Fatalf("diff.Chats = %#v, want projected chat", full.Chats)
	}
}

func chatIDs(chats []tg.ChatClazz) []int64 {
	ids := make([]int64, 0, len(chats))
	for _, chat := range chats {
		if c, ok := chat.(*tg.TLChat); ok {
			ids = append(ids, c.Id)
		}
	}
	return ids
}
