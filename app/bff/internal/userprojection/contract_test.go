package userprojection

import (
	"context"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestInternalFillUpdatesUsersDelegatesToPublicSemantics(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				tg.MakeTLUser(&tg.TLUser{Id: 1002}),
			}}),
		},
	}).ToUserProjectionBundle()}
	updates := tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{
				FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
				PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
			}),
		})},
		Users: []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 9999})},
	})

	if err := FillUpdatesUsers(context.Background(), client, 1001, updates.ToUpdates(), MissingStoredReference); err != nil {
		t.Fatalf("FillUpdatesUsers() error = %v", err)
	}
	if len(updates.Users) != 2 || updates.Users[0].(*tg.TLUser).Id != 1001 || updates.Users[1].(*tg.TLUser).Id != 1002 {
		t.Fatalf("users = %#v", updates.Users)
	}
}
