package dependencies

import (
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestCollectUpdates(t *testing.T) {
	tests := []struct {
		name    string
		updates []tg.UpdateClazz
		want    DependencySet
	}{
		{
			name: "updateNewMessage message from_id peer user",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: message(peerUser(101), nil),
				}),
			},
			want: DependencySet{UserIDs: []int64{101}},
		},
		{
			name: "updateNewMessage message peer_id peer user",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: message(nil, peerUser(102)),
				}),
			},
			want: DependencySet{UserIDs: []int64{102}},
		},
		{
			name: "updateNewMessage message peer_id peer chat",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: message(nil, peerChat(201)),
				}),
			},
			want: DependencySet{ChatIDs: []int64{201}},
		},
		{
			name: "messageActionChatCreate users",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: serviceMessage(peerChat(202), tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
						Users: []int64{104, 103, 104},
					})),
				}),
			},
			want: DependencySet{UserIDs: []int64{103, 104}, ChatIDs: []int64{202}},
		},
		{
			name: "messageActionChatAddUser users",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: serviceMessage(nil, tg.MakeTLMessageActionChatAddUser(&tg.TLMessageActionChatAddUser{
						Users: []int64{106, 105},
					})),
				}),
			},
			want: DependencySet{UserIDs: []int64{105, 106}},
		},
		{
			name: "messageActionChatDeleteUser user_id",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: serviceMessage(nil, tg.MakeTLMessageActionChatDeleteUser(&tg.TLMessageActionChatDeleteUser{
						UserId: 107,
					})),
				}),
			},
			want: DependencySet{UserIDs: []int64{107}},
		},
		{
			name: "updateEditMessage message",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
					Message: message(peerUser(108), peerChat(203)),
				}),
			},
			want: DependencySet{UserIDs: []int64{108}, ChatIDs: []int64{203}},
		},
		{
			name: "updateReadHistoryInbox peer",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateReadHistoryInbox(&tg.TLUpdateReadHistoryInbox{
					Peer: peerChat(204),
				}),
			},
			want: DependencySet{ChatIDs: []int64{204}},
		},
		{
			name: "updateReadHistoryOutbox peer",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateReadHistoryOutbox(&tg.TLUpdateReadHistoryOutbox{
					Peer: peerUser(109),
				}),
			},
			want: DependencySet{UserIDs: []int64{109}},
		},
		{
			name: "updateChat chat_id",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateChat(&tg.TLUpdateChat{ChatId: 205}),
			},
			want: DependencySet{ChatIDs: []int64{205}},
		},
		{
			name: "updateChatParticipants participants",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateChatParticipants(&tg.TLUpdateChatParticipants{
					Participants: tg.MakeTLChatParticipants(&tg.TLChatParticipants{
						ChatId: 206,
						Participants: []tg.ChatParticipantClazz{
							tg.MakeTLChatParticipant(&tg.TLChatParticipant{UserId: 110, InviterId: 111}),
							tg.MakeTLChatParticipantCreator(&tg.TLChatParticipantCreator{UserId: 112}),
							tg.MakeTLChatParticipantAdmin(&tg.TLChatParticipantAdmin{UserId: 113, InviterId: 114}),
						},
					}),
				}),
			},
			want: DependencySet{UserIDs: []int64{110, 111, 112, 113, 114}, ChatIDs: []int64{206}},
		},
		{
			name: "updateUserTyping user_id",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateUserTyping(&tg.TLUpdateUserTyping{UserId: 115}),
			},
			want: DependencySet{UserIDs: []int64{115}},
		},
		{
			name: "updateChatUserTyping chat_id and user_id",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateChatUserTyping(&tg.TLUpdateChatUserTyping{
					ChatId: 207,
					FromId: peerUser(116),
				}),
			},
			want: DependencySet{UserIDs: []int64{116}, ChatIDs: []int64{207}},
		},
		{
			name: "updatePeerSettings peer",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdatePeerSettings(&tg.TLUpdatePeerSettings{
					Peer: peerUser(117),
				}),
			},
			want: DependencySet{UserIDs: []int64{117}},
		},
		{
			name: "updateNotifySettings peer",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNotifySettings(&tg.TLUpdateNotifySettings{
					Peer: tg.MakeTLNotifyPeer(&tg.TLNotifyPeer{Peer: peerChat(208)}),
				}),
			},
			want: DependencySet{ChatIDs: []int64{208}},
		},
		{
			name: "channel peer cases",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewChannelMessage(&tg.TLUpdateNewChannelMessage{
					Message: message(peerUser(118), peerChannel(301)),
				}),
				tg.MakeTLUpdateReadChannelInbox(&tg.TLUpdateReadChannelInbox{ChannelId: 302}),
				tg.MakeTLUpdateReadChannelOutbox(&tg.TLUpdateReadChannelOutbox{ChannelId: 303}),
				tg.MakeTLUpdateChannelUserTyping(&tg.TLUpdateChannelUserTyping{
					ChannelId: 304,
					FromId:    peerUser(119),
				}),
			},
			want: DependencySet{UserIDs: []int64{118, 119}, ChannelIDs: []int64{301, 302, 303, 304}},
		},
		{
			name: "sorts and de-duplicates ids",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: message(peerUser(3), peerChat(2)),
				}),
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: message(peerUser(1), peerChannel(4)),
				}),
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: message(peerUser(3), peerChat(2)),
				}),
			},
			want: DependencySet{UserIDs: []int64{1, 3}, ChatIDs: []int64{2}, ChannelIDs: []int64{4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CollectUpdates(tt.updates); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("CollectUpdates() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestCollectUpdatesTypedNilValuesDoNotPanic(t *testing.T) {
	var nilUpdate *tg.TLUpdateNewMessage
	var nilMessage *tg.TLMessage
	var nilPeer *tg.TLPeerUser
	var nilNotifyPeer *tg.TLNotifyPeer
	var nilDialogPeer *tg.TLDialogPeer
	var nilChatParticipants *tg.TLChatParticipants
	var nilChatParticipant *tg.TLChatParticipant
	var nilReplyHeader *tg.TLMessageReplyHeader
	var nilMedia *tg.TLMessageMediaContact
	var nilAction *tg.TLMessageActionChatCreate

	tests := []struct {
		name    string
		updates []tg.UpdateClazz
	}{
		{
			name:    "typed nil UpdateClazz",
			updates: []tg.UpdateClazz{nilUpdate},
		},
		{
			name: "typed nil MessageClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{Message: nilMessage}),
			},
		},
		{
			name: "typed nil PeerClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: message(nilPeer, nil),
				}),
			},
		},
		{
			name: "typed nil NotifyPeerClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNotifySettings(&tg.TLUpdateNotifySettings{Peer: nilNotifyPeer}),
			},
		},
		{
			name: "typed nil DialogPeerClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateDialogPinned(&tg.TLUpdateDialogPinned{Peer: nilDialogPeer}),
			},
		},
		{
			name: "typed nil ChatParticipantsClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateChatParticipants(&tg.TLUpdateChatParticipants{Participants: nilChatParticipants}),
			},
		},
		{
			name: "typed nil ChatParticipantClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateChatParticipants(&tg.TLUpdateChatParticipants{
					Participants: tg.MakeTLChatParticipants(&tg.TLChatParticipants{
						Participants: []tg.ChatParticipantClazz{nilChatParticipant},
					}),
				}),
			},
		},
		{
			name: "typed nil MessageReplyHeaderClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: tg.MakeTLMessage(&tg.TLMessage{ReplyTo: nilReplyHeader}),
				}),
			},
		},
		{
			name: "typed nil MessageMediaClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: tg.MakeTLMessage(&tg.TLMessage{Media: nilMedia}),
				}),
			},
		},
		{
			name: "typed nil MessageActionClazz",
			updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message: serviceMessage(nil, nilAction),
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("CollectUpdates panicked: %v", r)
				}
			}()

			if got := CollectUpdates(tt.updates); !reflect.DeepEqual(got, DependencySet{}) {
				t.Fatalf("CollectUpdates() = %#v, want empty DependencySet", got)
			}
		})
	}
}

func message(fromID, peerID tg.PeerClazz) tg.MessageClazz {
	return tg.MakeTLMessage(&tg.TLMessage{
		FromId: fromID,
		PeerId: peerID,
	})
}

func serviceMessage(peerID tg.PeerClazz, action tg.MessageActionClazz) tg.MessageClazz {
	return tg.MakeTLMessageService(&tg.TLMessageService{
		PeerId: peerID,
		Action: action,
	})
}

func peerUser(id int64) tg.PeerClazz {
	return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: id})
}

func peerChat(id int64) tg.PeerClazz {
	return tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: id})
}

func peerChannel(id int64) tg.PeerClazz {
	return tg.MakeTLPeerChannel(&tg.TLPeerChannel{ChannelId: id})
}
