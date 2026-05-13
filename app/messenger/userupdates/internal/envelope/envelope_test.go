package envelope

import (
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestReplyEnvelopeAddsUpdateMessageID(t *testing.T) {
	got, err := BuildUpdates(Input{
		Mode:          ModeReply,
		Updates:       []tg.UpdateClazz{newMessageUpdate(101)},
		MessageIDByID: map[int32]int64{101: 9001},
	})
	if err != nil {
		t.Fatalf("BuildUpdates() error = %v", err)
	}

	updates, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
	if len(updates.Updates) != 2 {
		t.Fatalf("len(updates) = %d, want 2", len(updates.Updates))
	}
	idUpdate, ok := updates.Updates[0].(*tg.TLUpdateMessageID)
	if !ok {
		t.Fatalf("first update = %T, want *tg.TLUpdateMessageID", updates.Updates[0])
	}
	if idUpdate.Id != 101 || idUpdate.RandomId != 9001 {
		t.Fatalf("updateMessageID = %+v, want id 101 random_id 9001", idUpdate)
	}
	newMessage, ok := updates.Updates[1].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("second update = %T, want *tg.TLUpdateNewMessage", updates.Updates[1])
	}
	message, ok := newMessage.Message.(*tg.TLMessage)
	if !ok || message.Id != 101 {
		t.Fatalf("second update message = %T/%+v, want id 101", newMessage.Message, newMessage.Message)
	}
}

func TestReplyEnvelopeRejectsMissingRandomID(t *testing.T) {
	_, err := BuildUpdates(Input{
		Mode:          ModeReply,
		Updates:       []tg.UpdateClazz{newMessageUpdate(101)},
		MessageIDByID: map[int32]int64{101: 0},
	})
	if err == nil {
		t.Fatal("BuildUpdates() error = nil, want error")
	}
}

func TestReplyEnvelopeRejectsMissingUpdateMessageIDMapping(t *testing.T) {
	_, err := BuildUpdates(Input{
		Mode:          ModeReply,
		Updates:       []tg.UpdateClazz{newMessageUpdate(101)},
		MessageIDByID: map[int32]int64{102: 9002},
	})
	if err == nil {
		t.Fatal("BuildUpdates() error = nil, want error")
	}
}

func TestStreamEnvelopeRejectsUpdateMessageID(t *testing.T) {
	_, err := BuildUpdates(Input{
		Mode:    ModeSenderStream,
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: 101, RandomId: 9001})},
	})
	if err == nil {
		t.Fatal("BuildUpdates() error = nil, want error")
	}
}

func TestEnvelopePreservesUsersChatsDateSeq(t *testing.T) {
	user := tg.MakeTLUser(&tg.TLUser{Id: 10})
	chat := tg.MakeTLChat(&tg.TLChat{Id: 20})

	got, err := BuildUpdates(Input{
		Mode:    ModeReceiverStream,
		Updates: []tg.UpdateClazz{newMessageUpdate(101)},
		Users:   []tg.UserClazz{user},
		Chats:   []tg.ChatClazz{chat},
		Date:    111,
		Seq:     222,
	})
	if err != nil {
		t.Fatalf("BuildUpdates() error = %v", err)
	}

	updates, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
	if updates.Date != 111 || updates.Seq != 222 {
		t.Fatalf("date/seq = %d/%d, want 111/222", updates.Date, updates.Seq)
	}
	if len(updates.Users) != 1 || updates.Users[0] != user {
		t.Fatalf("users = %+v, want preserved user", updates.Users)
	}
	if len(updates.Chats) != 1 || updates.Chats[0] != chat {
		t.Fatalf("chats = %+v, want preserved chat", updates.Chats)
	}
}

func TestReplyEnvelopeRejectsDuplicateMessageIDDifferentRandomID(t *testing.T) {
	_, err := BuildUpdates(Input{
		Mode: ModeReply,
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: 101, RandomId: 9001}),
			tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: 101, RandomId: 9002}),
			newMessageUpdate(101),
		},
		MessageIDByID: map[int32]int64{101: 9001},
	})
	if err == nil {
		t.Fatal("BuildUpdates() error = nil, want error")
	}
}

func TestReplyEnvelopeRejectsDuplicateUpdateNewMessageID(t *testing.T) {
	_, err := BuildUpdates(Input{
		Mode: ModeReply,
		Updates: []tg.UpdateClazz{
			newMessageUpdate(101),
			newMessageUpdate(101),
		},
		MessageIDByID: map[int32]int64{101: 9001},
	})
	if err == nil {
		t.Fatal("BuildUpdates() error = nil, want error")
	}
}

func TestReplyEnvelopeRejectsNilUpdateMessageID(t *testing.T) {
	var nilUpdateMessageID *tg.TLUpdateMessageID

	_, err := BuildUpdates(Input{
		Mode:    ModeReply,
		Updates: []tg.UpdateClazz{nilUpdateMessageID},
	})
	if err == nil {
		t.Fatal("BuildUpdates() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "nil updateMessageID") {
		t.Fatalf("BuildUpdates() error = %v, want nil updateMessageID", err)
	}
}

func TestReplyEnvelopeRejectsNilMessageValue(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		var nilMessage *tg.TLMessage

		_, err := BuildUpdates(Input{
			Mode: ModeReply,
			Updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{Message: nilMessage}),
			},
		})
		if err == nil {
			t.Fatal("BuildUpdates() error = nil, want error")
		}
	})

	t.Run("messageService", func(t *testing.T) {
		var nilMessageService *tg.TLMessageService

		_, err := BuildUpdates(Input{
			Mode: ModeReply,
			Updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{Message: nilMessageService}),
			},
		})
		if err == nil {
			t.Fatal("BuildUpdates() error = nil, want error")
		}
	})
}

func newMessageUpdate(id int32) tg.UpdateClazz {
	return tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
		Message: tg.MakeTLMessage(&tg.TLMessage{Id: id}),
	})
}
