package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeEditProjectionUserClient struct {
	userclient.UserClient
	in *userpb.TLUserGetUserProjectionBundle
}

func (f *fakeEditProjectionUserClient) UserGetUserProjectionBundle(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	f.in = in
	return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				tg.MakeTLUser(&tg.TLUser{Id: 1002}),
			}}),
		},
	}).ToUserProjectionBundle(), nil
}

func TestMessagesEditMessageProjectsUsers(t *testing.T) {
	userClient := &fakeEditProjectionUserClient{}
	core := New(context.Background(), &svc.ServiceContext{Repo: &repository.Repository{
		MsgClient: &messagesFakeMsgClient{
			editMessage: func(_ context.Context, _ *msg.TLMsgEditMessage) (*tg.Updates, error) {
				return tg.MakeTLUpdates(&tg.TLUpdates{
					Updates: []tg.UpdateClazz{tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
						Message: tg.MakeTLMessage(&tg.TLMessage{
							FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
							PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
						}),
					})},
					Users: []tg.UserClazz{},
					Chats: []tg.ChatClazz{},
					Date:  1_772_000_099,
					Seq:   0,
				}).ToUpdates(), nil
			},
		},
		UserClient: userClient,
	}})
	core.MD = &metadata.RpcMetadata{UserId: 1001, PermAuthKeyId: 9001}
	text := "edited"

	got, err := core.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:    inputPeerUser(1002),
		Id:      7,
		Message: &text,
	})
	if err != nil {
		t.Fatalf("MessagesEditMessage() error = %v", err)
	}
	updates, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("expected updates, got %s", got.ClazzName())
	}
	if updates.Date != 1_772_000_099 || len(updates.Updates) != 1 {
		t.Fatalf("updates = %+v", updates)
	}
	if len(updates.Users) != 2 || updates.Users[0].(*tg.TLUser).Id != 1001 || updates.Users[1].(*tg.TLUser).Id != 1002 {
		t.Fatalf("users = %#v", updates.Users)
	}
	if userClient.in == nil || len(userClient.in.ViewerUserIds) != 1 || userClient.in.ViewerUserIds[0] != 1001 || len(userClient.in.TargetUserIds) != 2 {
		t.Fatalf("projection request = %#v", userClient.in)
	}
}
