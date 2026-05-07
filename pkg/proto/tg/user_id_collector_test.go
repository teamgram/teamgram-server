package tg

import "testing"

func TestCollectUserIDsFromUpdatesKeepsOrderAndExcludesChannels(t *testing.T) {
	msg := MakeTLMessage(&TLMessage{
		FromId: MakeTLPeerUser(&TLPeerUser{UserId: 1001}),
		PeerId: MakeTLPeerUser(&TLPeerUser{UserId: 1002}),
	})
	updates := MakeTLUpdates(&TLUpdates{
		Updates: []UpdateClazz{
			MakeTLUpdateEditMessage(&TLUpdateEditMessage{Message: msg}),
			MakeTLUpdateNewChannelMessage(&TLUpdateNewChannelMessage{Message: MakeTLMessage(&TLMessage{
				FromId: MakeTLPeerUser(&TLPeerUser{UserId: 1003}),
				PeerId: MakeTLPeerChannel(&TLPeerChannel{ChannelId: 2001}),
			})}),
		},
	})

	got := CollectUserIDsFromUpdates(updates.ToUpdates())
	want := []int64{1001, 1002, 1003}
	if !sameInt64s(got, want) {
		t.Fatalf("ids = %v, want %v", got, want)
	}
}

func TestCollectUserIDsFromMessageIncludesFwdAndReplyPeerUsers(t *testing.T) {
	replyPeer := MakeTLPeerUser(&TLPeerUser{UserId: 1004})
	msg := MakeTLMessage(&TLMessage{
		FromId:   MakeTLPeerUser(&TLPeerUser{UserId: 1001}),
		PeerId:   MakeTLPeerUser(&TLPeerUser{UserId: 1002}),
		FwdFrom:  MakeTLMessageFwdHeader(&TLMessageFwdHeader{FromId: MakeTLPeerUser(&TLPeerUser{UserId: 1003})}),
		ReplyTo:  MakeTLMessageReplyHeader(&TLMessageReplyHeader{ReplyToPeerId: replyPeer}),
		ViaBotId: int64PtrForUserIDCollectorTest(1005),
	})

	got := CollectUserIDsFromMessage(msg)
	want := []int64{1001, 1002, 1003, 1005, 1004}
	if !sameInt64s(got, want) {
		t.Fatalf("ids = %v, want %v", got, want)
	}
}

func TestCollectUserIDsFromDifferenceCollectsMessagesAndOtherUpdates(t *testing.T) {
	diff := MakeTLUpdatesDifference(&TLUpdatesDifference{
		NewMessages: []MessageClazz{MakeTLMessage(&TLMessage{
			FromId: MakeTLPeerUser(&TLPeerUser{UserId: 1001}),
			PeerId: MakeTLPeerUser(&TLPeerUser{UserId: 1002}),
		})},
		OtherUpdates: []UpdateClazz{MakeTLUpdateEditMessage(&TLUpdateEditMessage{
			Message: MakeTLMessage(&TLMessage{
				FromId: MakeTLPeerUser(&TLPeerUser{UserId: 1002}),
				PeerId: MakeTLPeerUser(&TLPeerUser{UserId: 1003}),
			}),
		})},
	})

	got := CollectUserIDsFromDifference(diff.ToUpdatesDifference())
	want := []int64{1001, 1002, 1003}
	if !sameInt64s(got, want) {
		t.Fatalf("ids = %v, want %v", got, want)
	}
}

func TestCollectUserIDsFromMessagesMessagesCollectsMessages(t *testing.T) {
	messages := MakeTLMessagesMessages(&TLMessagesMessages{
		Messages: []MessageClazz{MakeTLMessage(&TLMessage{
			FromId: MakeTLPeerUser(&TLPeerUser{UserId: 1001}),
			PeerId: MakeTLPeerUser(&TLPeerUser{UserId: 1002}),
		})},
		Chats: []ChatClazz{},
		Users: []UserClazz{},
	}).ToMessagesMessages()

	got := CollectUserIDsFromMessagesMessages(messages)
	want := []int64{1001, 1002}
	if !sameInt64s(got, want) {
		t.Fatalf("ids = %v, want %v", got, want)
	}
}

func TestCollectUserIDsFromShortUpdatesIsEmpty(t *testing.T) {
	got := CollectUserIDsFromUpdates(MakeTLUpdateShortMessage(&TLUpdateShortMessage{UserId: 1001}).ToUpdates())
	if len(got) != 0 {
		t.Fatalf("short update ids = %v, want empty", got)
	}
}

func sameInt64s(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func int64PtrForUserIDCollectorTest(v int64) *int64 {
	return &v
}
