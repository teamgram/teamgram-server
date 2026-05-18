package chatprojection

import (
	"context"
	"errors"
	"reflect"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	bizchatproj "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chatprojection"
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

func TestMissingPolicyMapsToPublicProjectionPolicy(t *testing.T) {
	if publicMissingPolicy(MissingExplicitInput) != bizchatproj.MissingExplicitInput {
		t.Fatalf("MissingExplicitInput did not map to public explicit policy")
	}
	if publicMissingPolicy(MissingStoredReference) != bizchatproj.MissingStoredReference {
		t.Fatalf("MissingStoredReference did not map to public stored-reference policy")
	}
}

func TestProjectChatsEmptyInputPreservesBFFEmptySlice(t *testing.T) {
	client := &fakeChatClient{}
	got, err := ProjectChats(context.Background(), client, 1001, nil, MissingExplicitInput)
	if err != nil {
		t.Fatalf("ProjectChats(empty) error = %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("ProjectChats(empty) = %#v, want empty non-nil slice", got)
	}
	if client.in != nil {
		t.Fatalf("client called for empty input: %#v", client.in)
	}
}

func TestProjectChatsMapsExplicitMissingToChatIdInvalid(t *testing.T) {
	client := &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		MissingChatIds: []int64{3001},
	}).ToChatProjectionBundle()}
	_, err := ProjectChats(context.Background(), client, 1001, []int64{3001}, MissingExplicitInput)
	if !errors.Is(err, tg.ErrChatIdInvalid) {
		t.Fatalf("ProjectChats() error = %v, want ErrChatIdInvalid", err)
	}
}

func TestProjectChatsWrapsInvalidRequest(t *testing.T) {
	client := &fakeChatClient{err: chatpb.ErrChatInvalidArgument}

	_, err := ProjectChats(context.Background(), client, 1001, []int64{3001}, MissingExplicitInput)
	if !errors.Is(err, bizchatproj.ErrInvalidRequest) {
		t.Fatalf("ProjectChats() error = %v, want wrapped invalid request", err)
	}
}

func TestProjectChatsMissingViewerPreservesBFFEmptySlice(t *testing.T) {
	client := &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 9999}),
		},
	}).ToChatProjectionBundle()}
	got, err := ProjectChats(context.Background(), client, 1001, []int64{3001}, MissingStoredReference)
	if err != nil {
		t.Fatalf("ProjectChats(missing viewer) error = %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("ProjectChats(missing viewer) = %#v, want empty non-nil slice", got)
	}
}

func TestFillUpdatesChatsReplacesGroupCallPeerChat(t *testing.T) {
	projected := tg.MakeTLChat(&tg.TLChat{Id: 7001})
	client := &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{
				ViewerUserId: 1001,
				Chats:        []tg.ChatClazz{projected},
			}),
		},
	}).ToChatProjectionBundle()}
	updates := tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateGroupCall(&tg.TLUpdateGroupCall{
				Peer: tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 7001}),
			}),
		},
		Chats: []tg.ChatClazz{tg.MakeTLChat(&tg.TLChat{Id: 9999})},
	})

	err := FillUpdatesChats(context.Background(), client, 1001, updates.ToUpdates(), MissingStoredReference)
	if err != nil {
		t.Fatalf("FillUpdatesChats() error = %v", err)
	}
	if !reflect.DeepEqual(client.in.TargetChatIds, []int64{7001}) {
		t.Fatalf("TargetChatIds = %#v, want [7001]", client.in.TargetChatIds)
	}
	if len(updates.Chats) != 1 || updates.Chats[0] != projected {
		t.Fatalf("updates.Chats = %#v, want projected chat", updates.Chats)
	}
}

func TestFillUpdatesChatsReplacesCombinedChats(t *testing.T) {
	projected := tg.MakeTLChat(&tg.TLChat{Id: 7002})
	client := &fakeChatClient{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{
				ViewerUserId: 1001,
				Chats:        []tg.ChatClazz{projected},
			}),
		},
	}).ToChatProjectionBundle()}
	updates := tg.MakeTLUpdatesCombined(&tg.TLUpdatesCombined{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateGroupCallParticipants(&tg.TLUpdateGroupCallParticipants{
				Participants: []tg.GroupCallParticipantClazz{
					tg.MakeTLGroupCallParticipant(&tg.TLGroupCallParticipant{
						Peer: tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 7002}),
					}),
				},
			}),
		},
		Chats: []tg.ChatClazz{tg.MakeTLChat(&tg.TLChat{Id: 9999})},
	})

	err := FillUpdatesChats(context.Background(), client, 1001, updates.ToUpdates(), MissingStoredReference)
	if err != nil {
		t.Fatalf("FillUpdatesChats(combined) error = %v", err)
	}
	if !reflect.DeepEqual(client.in.TargetChatIds, []int64{7002}) {
		t.Fatalf("TargetChatIds = %#v, want [7002]", client.in.TargetChatIds)
	}
	if len(updates.Chats) != 1 || updates.Chats[0] != projected {
		t.Fatalf("updatesCombined.Chats = %#v, want projected chat", updates.Chats)
	}
}
