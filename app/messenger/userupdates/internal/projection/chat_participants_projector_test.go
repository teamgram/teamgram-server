package projection

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/envelope"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestProjectChatParticipantsChangedFactProjectsRoles(t *testing.T) {
	got, err := ProjectChatParticipantsChangedFact(payload.ChatParticipantsChangedFactV1{
		SchemaVersion: payload.MessageOperationSchemaVersionV4,
		ChatID:        55,
		ActorUserID:   1001,
		Version:       7,
		Participants: []payload.ChatParticipantFactV1{
			{UserID: 1001, Role: "creator", Date: 1_772_000_000},
			{UserID: 1002, Role: "admin", InviterUserID: 1001, Date: 1_772_000_001},
			{UserID: 1003, Role: "member", InviterUserID: 1001, Date: 1_772_000_002},
		},
	})
	if err != nil {
		t.Fatalf("ProjectChatParticipantsChangedFact() error = %v", err)
	}

	update, ok := got.Update.(*tg.TLUpdateChatParticipants)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateChatParticipants", got.Update)
	}
	participants, ok := update.Participants.(*tg.TLChatParticipants)
	if !ok {
		t.Fatalf("participants = %T, want *tg.TLChatParticipants", update.Participants)
	}
	if participants.ChatId != 55 || participants.Version != 7 || len(participants.Participants) != 3 {
		t.Fatalf("participants = %+v", participants)
	}
	if creator, ok := participants.Participants[0].(*tg.TLChatParticipantCreator); !ok || creator.UserId != 1001 {
		t.Fatalf("creator participant = %#v", participants.Participants[0])
	}
	if admin, ok := participants.Participants[1].(*tg.TLChatParticipantAdmin); !ok || admin.UserId != 1002 || admin.InviterId != 1001 || admin.Date != 1_772_000_001 {
		t.Fatalf("admin participant = %#v", participants.Participants[1])
	}
	if member, ok := participants.Participants[2].(*tg.TLChatParticipant); !ok || member.UserId != 1003 || member.InviterId != 1001 || member.Date != 1_772_000_002 {
		t.Fatalf("member participant = %#v", participants.Participants[2])
	}
}

func TestProjectCreateChatFactOrder(t *testing.T) {
	chatFact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, payload.ChatParticipantsChangedFactV1{
		SchemaVersion: payload.MessageOperationSchemaVersionV4,
		ChatID:        55,
		ActorUserID:   1001,
		Version:       1,
		Participants: []payload.ChatParticipantFactV1{
			{UserID: 1001, Role: "creator"},
			{UserID: 1002, Role: "member", InviterUserID: 1001, Date: 1_772_000_000},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	messageFact, err := payload.WrapFact(payload.FactKindNewMessage, payload.NewMessageFactV1{
		SchemaVersion:  payload.MessageOperationSchemaVersionV4,
		PeerType:       payload.PeerTypeChat,
		PeerID:         55,
		SenderUserID:   1001,
		ClientRandomID: 999,
		Date:           1_772_000_000,
		ServiceAction: &payload.ServiceActionRefV1{
			SchemaVersion: payload.ServiceActionSchemaVersionV1,
			Kind:          payload.ServiceActionKindChatCreate,
			Title:         "team",
			Users:         []int64{1001, 1002},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := ProjectFacts([]payload.UpdateFactV1{chatFact, messageFact}, ViewerContext{UserID: 1001}, envelope.ModeDifference, 21, 2, 101)
	if err != nil {
		t.Fatalf("ProjectFacts() error = %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("projected update count = %d, want 2", len(got))
	}
	if _, ok := got[0].Update.(*tg.TLUpdateChatParticipants); !ok {
		t.Fatalf("first update = %T, want *tg.TLUpdateChatParticipants", got[0].Update)
	}
	if _, ok := got[1].Update.(*tg.TLUpdateNewMessage); !ok {
		t.Fatalf("second update = %T, want *tg.TLUpdateNewMessage", got[1].Update)
	}
	service, ok := got[1].Update.(*tg.TLUpdateNewMessage).Message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("second message = %T, want *tg.TLMessageService", got[1].Update.(*tg.TLUpdateNewMessage).Message)
	}
	action, ok := service.Action.(*tg.TLMessageActionChatCreate)
	if !ok || action.Title != "team" || len(action.Users) != 2 {
		t.Fatalf("service action = %#v, want chat_create action", service.Action)
	}
}

func TestProjectChatParticipantsChangedFactRejectsInvalidUserID(t *testing.T) {
	_, err := ProjectChatParticipantsChangedFact(payload.ChatParticipantsChangedFactV1{
		SchemaVersion: payload.MessageOperationSchemaVersionV4,
		ChatID:        55,
		Version:       1,
		Participants: []payload.ChatParticipantFactV1{
			{UserID: 0, Role: "member", InviterUserID: 1001, Date: 1_772_000_000},
		},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestProjectChatParticipantsChangedFactRejectsInvalidDate(t *testing.T) {
	_, err := ProjectChatParticipantsChangedFact(payload.ChatParticipantsChangedFactV1{
		SchemaVersion: payload.MessageOperationSchemaVersionV4,
		ChatID:        55,
		Version:       1,
		Participants: []payload.ChatParticipantFactV1{
			{UserID: 1002, Role: "member", InviterUserID: 1001, Date: -1},
		},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("error = %v, want ErrUserupdatesStorage", err)
	}
}

func mustUpdateNewMessage(t *testing.T, projected ProjectedUpdate) *tg.TLUpdateNewMessage {
	t.Helper()
	update, ok := projected.Update.(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", projected.Update)
	}
	return update
}
