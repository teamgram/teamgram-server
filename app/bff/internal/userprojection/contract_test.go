package userprojection

import (
	"context"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestProjectionContractSameViewerTargetAcrossBoundaries(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001, Self: true}),
				tg.MakeTLUser(&tg.TLUser{Id: 1002, Contact: true}),
			}}),
		},
	}).ToUserProjectionBundle()}

	editUpdates := tg.MakeTLUpdates(&tg.TLUpdates{Updates: []tg.UpdateClazz{
		tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{Message: tg.MakeTLMessage(&tg.TLMessage{
			FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
			PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
		})}),
	}})
	if err := FillUpdatesUsers(context.Background(), client, 1001, editUpdates.ToUpdates(), MissingStoredReference); err != nil {
		t.Fatalf("FillUpdatesUsers(edit) error = %v", err)
	}

	diff := tg.MakeTLUpdatesDifference(&tg.TLUpdatesDifference{
		NewMessages: []tg.MessageClazz{tg.MakeTLMessage(&tg.TLMessage{
			FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
			PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
		})},
	}).ToUpdatesDifference()
	if err := FillDifferenceUsers(context.Background(), client, 1001, diff, MissingStoredReference); err != nil {
		t.Fatalf("FillDifferenceUsers(diff) error = %v", err)
	}

	diffFull, ok := diff.ToUpdatesDifference()
	if !ok {
		t.Fatalf("difference = %s, want updates.difference", diff.ClazzName())
	}
	if len(editUpdates.Users) != 2 || len(diffFull.Users) != 2 {
		t.Fatalf("projection vector length drift edit=%#v diff=%#v", editUpdates.Users, diffFull.Users)
	}
	editSelf, ok := editUpdates.Users[0].(*tg.TLUser)
	if !ok || editSelf.Id != 1001 || !editSelf.Self {
		t.Fatalf("edit self projection = %#v", editUpdates.Users[0])
	}
	diffSelf, ok := diffFull.Users[0].(*tg.TLUser)
	if !ok || diffSelf.Id != 1001 || !diffSelf.Self {
		t.Fatalf("difference self projection = %#v", diffFull.Users[0])
	}
	editContact, ok := editUpdates.Users[1].(*tg.TLUser)
	if !ok || editContact.Id != 1002 || !editContact.Contact {
		t.Fatalf("edit contact projection = %#v", editUpdates.Users[1])
	}
	diffContact, ok := diffFull.Users[1].(*tg.TLUser)
	if !ok || diffContact.Id != 1002 || !diffContact.Contact {
		t.Fatalf("difference contact projection = %#v", diffFull.Users[1])
	}
}
