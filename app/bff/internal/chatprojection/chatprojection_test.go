package chatprojection

import (
	"context"
	"errors"
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
