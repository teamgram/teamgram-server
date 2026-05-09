package core

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesForwardMessagesReordersSourcesAndSendsBatch(t *testing.T) {
	var fetch *msg.TLMsgGetUserMessageList
	var got *msg.TLMsgSendMessageV2
	updates := testUpdates()
	entity := tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{Offset: 0, Length: 3})
	media := tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{Photo: tg.MakeTLPhoto(&tg.TLPhoto{Id: 777})})
	existingSavedMsgID := int32(91)
	c := newSendMsgCore(&messagesFakeMsgClient{
		getUserMessageList: func(_ context.Context, in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
			fetch = in
			return &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{
				tg.MakeTLMessageBox(&tg.TLMessageBox{
					MessageId: 2,
					UserId:    100,
					PeerType:  payload.PeerTypeUser,
					PeerId:    100,
					Message: tg.MakeTLMessage(&tg.TLMessage{
						Id:      2,
						FromId:  tg.MakePeerUser(302),
						PeerId:  tg.MakePeerUser(100),
						Date:    2002,
						Message: "two",
						Media:   media,
						FwdFrom: tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
							FromId:         tg.MakePeerUser(900),
							Date:           1900,
							SavedFromPeer:  tg.MakePeerUser(800),
							SavedFromMsgId: &existingSavedMsgID,
						}),
					}),
				}),
				tg.MakeTLMessageBox(&tg.TLMessageBox{
					MessageId: 1,
					UserId:    100,
					PeerType:  payload.PeerTypeUser,
					PeerId:    300,
					Message: tg.MakeTLMessage(&tg.TLMessage{
						Id:          1,
						FromId:      tg.MakePeerUser(301),
						PeerId:      tg.MakePeerUser(300),
						Date:        2001,
						Message:     "one",
						Entities:    []tg.MessageEntityClazz{entity},
						Reactions:   tg.MakeTLMessageReactions(&tg.TLMessageReactions{}),
						ReplyMarkup: tg.MakeTLReplyKeyboardHide(&tg.TLReplyKeyboardHide{}),
						ReplyTo:     tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &existingSavedMsgID}),
					}),
				}),
			}}, nil
		},
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return updates, nil
		},
	}, 100, 200)

	result, err := c.MessagesForwardMessages(&tg.TLMessagesForwardMessages{
		Silent:     true,
		Background: true,
		Noforwards: true,
		FromPeer:   inputPeerUser(300),
		ToPeer:     inputPeerUser(400),
		Id:         []int32{1, 2},
		RandomId:   []int64{11, 12},
	})
	if err != nil {
		t.Fatalf("MessagesForwardMessages() error = %v", err)
	}
	if result != updates {
		t.Fatal("handler did not pass through tg.Updates")
	}
	if fetch == nil || fetch.UserId != 100 || !reflect.DeepEqual(fetch.IdList, []int32{1, 2}) {
		t.Fatalf("fetch = %#v, want self user and requested ids", fetch)
	}
	if got == nil || len(got.Message) != 2 {
		t.Fatalf("send batch = %#v, want two outboxes", got)
	}
	if got.SourcePermAuthKeyId == nil || *got.SourcePermAuthKeyId != 200 {
		t.Fatalf("SourcePermAuthKeyId = %v, want 200", got.SourcePermAuthKeyId)
	}
	if got.PeerType != payload.PeerTypeUser || got.PeerId != 400 {
		t.Fatalf("target peer = %d/%d, want user 400", got.PeerType, got.PeerId)
	}
	first := assertForwardOutbox(t, got.Message[0], 11, "one")
	second := assertForwardOutbox(t, got.Message[1], 12, "two")
	if !got.Message[0].Background || !got.Message[1].Background {
		t.Fatalf("background flags = %t/%t, want true", got.Message[0].Background, got.Message[1].Background)
	}
	if !first.Silent || !first.Noforwards || !second.Silent || !second.Noforwards {
		t.Fatalf("message flags = first silent/noforwards %t/%t second %t/%t", first.Silent, first.Noforwards, second.Silent, second.Noforwards)
	}
	if first.FromId == nil || first.PeerId == nil || first.Date == 0 || !first.Out {
		t.Fatalf("first routing/date/out not set: %#v", first)
	}
	if first.Reactions != nil || first.ReplyMarkup != nil || first.ReplyTo != nil {
		t.Fatalf("first transient fields were not cleared: reactions=%T reply_markup=%T reply_to=%T", first.Reactions, first.ReplyMarkup, first.ReplyTo)
	}
	if len(first.Entities) != 1 || first.Entities[0] != entity {
		t.Fatalf("first entities = %#v, want source entities", first.Entities)
	}
	if first.FwdFrom == nil || first.FwdFrom.FromId == nil || first.FwdFrom.SavedFromPeer == nil || first.FwdFrom.SavedFromMsgId == nil || *first.FwdFrom.SavedFromMsgId != 1 {
		t.Fatalf("first fwd_from = %#v, want source identity for requested id 1", first.FwdFrom)
	}
	if second.Media != media {
		t.Fatalf("second media = %#v, want source media", second.Media)
	}
	if second.FwdFrom == nil || second.FwdFrom.FromId == nil || second.FwdFrom.SavedFromPeer == nil || second.FwdFrom.SavedFromMsgId == nil || *second.FwdFrom.SavedFromMsgId != 2 {
		t.Fatalf("second fwd_from = %#v, want source identity for requested id 2", second.FwdFrom)
	}
	if second.FwdFrom.Date != 1900 {
		t.Fatalf("second fwd date = %d, want preserved existing forward date", second.FwdFrom.Date)
	}
}

func TestMessagesForwardMessagesRejectsInvalidRequests(t *testing.T) {
	topMsgID := int32(1)
	scheduleDate := int32(2000000)
	scheduleRepeat := int32(60)
	effect := int64(10)
	videoTimestamp := int32(2)
	stars := int64(1)
	tests := []struct {
		name  string
		in    *tg.TLMessagesForwardMessages
		patch func(*tg.TLMessagesForwardMessages)
		want  error
	}{
		{name: "nil", in: nil, want: tg.ErrInputRequestInvalid},
		{name: "mismatched lengths", patch: func(in *tg.TLMessagesForwardMessages) { in.RandomId = []int64{11} }, want: tg.ErrInputRequestInvalid},
		{name: "empty", patch: func(in *tg.TLMessagesForwardMessages) { in.Id = nil; in.RandomId = nil }, want: tg.ErrMessageIdInvalid},
		{name: "too many", patch: func(in *tg.TLMessagesForwardMessages) {
			in.Id = make([]int32, maxForwardMessages+1)
			in.RandomId = make([]int64, maxForwardMessages+1)
			for i := range in.Id {
				in.Id[i] = int32(i + 1)
				in.RandomId[i] = int64(i + 1)
			}
		}, want: tg.ErrInputRequestInvalid},
		{name: "zero id", patch: func(in *tg.TLMessagesForwardMessages) { in.Id[0] = 0 }, want: tg.ErrMessageIdInvalid},
		{name: "zero random id", patch: func(in *tg.TLMessagesForwardMessages) { in.RandomId[0] = 0 }, want: tg.ErrRandomIdEmpty},
		{name: "duplicate random id", patch: func(in *tg.TLMessagesForwardMessages) { in.RandomId[1] = in.RandomId[0] }, want: tg.ErrRandomIdDuplicate},
		{name: "unsupported source peer", patch: func(in *tg.TLMessagesForwardMessages) { in.FromPeer = inputPeerChat(5) }, want: tg.Err400PeerIdInvalid},
		{name: "unsupported target peer", patch: func(in *tg.TLMessagesForwardMessages) { in.ToPeer = inputPeerChat(5) }, want: tg.Err400PeerIdInvalid},
		{name: "with my score", patch: func(in *tg.TLMessagesForwardMessages) { in.WithMyScore = true }, want: tg.ErrInputRequestInvalid},
		{name: "drop author", patch: func(in *tg.TLMessagesForwardMessages) { in.DropAuthor = true }, want: tg.ErrInputRequestInvalid},
		{name: "drop media captions", patch: func(in *tg.TLMessagesForwardMessages) { in.DropMediaCaptions = true }, want: tg.ErrInputRequestInvalid},
		{name: "allow paid floodskip", patch: func(in *tg.TLMessagesForwardMessages) { in.AllowPaidFloodskip = true }, want: tg.ErrInputRequestInvalid},
		{name: "top msg id", patch: func(in *tg.TLMessagesForwardMessages) { in.TopMsgId = &topMsgID }, want: tg.ErrInputRequestInvalid},
		{name: "reply to", patch: func(in *tg.TLMessagesForwardMessages) {
			in.ReplyTo = tg.MakeTLInputReplyToMessage(&tg.TLInputReplyToMessage{ReplyToMsgId: 1})
		}, want: tg.ErrInputRequestInvalid},
		{name: "schedule date", patch: func(in *tg.TLMessagesForwardMessages) { in.ScheduleDate = &scheduleDate }, want: tg.ErrInputRequestInvalid},
		{name: "schedule repeat", patch: func(in *tg.TLMessagesForwardMessages) { in.ScheduleRepeatPeriod = &scheduleRepeat }, want: tg.ErrInputRequestInvalid},
		{name: "send as", patch: func(in *tg.TLMessagesForwardMessages) { in.SendAs = inputPeerUser(401) }, want: tg.ErrInputRequestInvalid},
		{name: "quick reply shortcut", patch: func(in *tg.TLMessagesForwardMessages) {
			in.QuickReplyShortcut = tg.MakeTLInputQuickReplyShortcut(&tg.TLInputQuickReplyShortcut{})
		}, want: tg.ErrInputRequestInvalid},
		{name: "effect", patch: func(in *tg.TLMessagesForwardMessages) { in.Effect = &effect }, want: tg.ErrInputRequestInvalid},
		{name: "video timestamp", patch: func(in *tg.TLMessagesForwardMessages) { in.VideoTimestamp = &videoTimestamp }, want: tg.ErrInputRequestInvalid},
		{name: "allow paid stars", patch: func(in *tg.TLMessagesForwardMessages) { in.AllowPaidStars = &stars }, want: tg.ErrInputRequestInvalid},
		{name: "suggested post", patch: func(in *tg.TLMessagesForwardMessages) {
			in.SuggestedPost = tg.MakeTLSuggestedPost(&tg.TLSuggestedPost{})
		}, want: tg.ErrInputRequestInvalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			c := newSendMsgCore(&messagesFakeMsgClient{
				getUserMessageList: func(context.Context, *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
					called = true
					return nil, nil
				},
				sendMessageV2: func(context.Context, *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
					called = true
					return nil, nil
				},
			}, 100, 200)
			in := tt.in
			if in == nil && tt.name != "nil" {
				in = validForwardMessagesRequest()
			}
			if tt.patch != nil {
				tt.patch(in)
			}
			_, err := c.MessagesForwardMessages(in)
			if !errors.Is(err, tt.want) {
				t.Fatalf("error = %v, want %v", err, tt.want)
			}
			if called {
				t.Fatal("downstream service was called but should not have been")
			}
		})
	}
}

func TestMessagesForwardMessagesMapsMissingOrInvalidSourceToMessageIdInvalid(t *testing.T) {
	tests := []struct {
		name string
		list *msg.VectorMessageBox
		err  error
		want error
	}{
		{name: "fetch error", err: msg.ErrMsgIdInvalid, want: tg.ErrMessageIdInvalid},
		{name: "missing", list: &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{messageBox(2, "two")}}, want: tg.ErrMessageIdInvalid},
		{name: "nil box", list: &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{nil}}, want: tg.ErrMessageIdInvalid},
		{name: "nil message", list: &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{tg.MakeTLMessageBox(&tg.TLMessageBox{MessageId: 1})}}, want: tg.ErrMessageIdInvalid},
		{name: "non TL message", list: &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{
			tg.MakeTLMessageBox(&tg.TLMessageBox{MessageId: 1, Message: tg.MakeTLMessageService(&tg.TLMessageService{Id: 1})}),
		}}, want: tg.ErrMessageIdInvalid},
		{name: "missing message id", list: &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{
			tg.MakeTLMessageBox(&tg.TLMessageBox{MessageId: 1, Message: tg.MakeTLMessage(&tg.TLMessage{Message: "one"})}),
		}}, want: tg.ErrMessageIdInvalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newSendMsgCore(&messagesFakeMsgClient{
				getUserMessageList: func(context.Context, *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
					return tt.list, tt.err
				},
			}, 100, 200)
			_, err := c.MessagesForwardMessages(validForwardMessagesRequest())
			if !errors.Is(err, tt.want) {
				t.Fatalf("error = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestMessagesForwardMessagesRejectsUnsupportedSourceMediaBeforeMsgSend(t *testing.T) {
	called := false
	c := newSendMsgCore(&messagesFakeMsgClient{
		getUserMessageList: func(context.Context, *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
			return &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{
				tg.MakeTLMessageBox(&tg.TLMessageBox{MessageId: 1, PeerType: payload.PeerTypeUser, PeerId: 300, Message: tg.MakeTLMessage(&tg.TLMessage{Id: 1, FromId: tg.MakePeerUser(300), PeerId: tg.MakePeerUser(100), Date: 1, Message: "contact", Media: tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{PhoneNumber: "1", FirstName: "a"})})}),
				messageBox(2, "plain"),
			}}, nil
		},
		sendMessageV2: func(context.Context, *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			called = true
			return nil, nil
		},
	}, 100, 200)
	_, err := c.MessagesForwardMessages(validForwardMessagesRequest())
	if !errors.Is(err, tg.ErrMessageIdInvalid) {
		t.Fatalf("error = %v, want MESSAGE_ID_INVALID", err)
	}
	if called {
		t.Fatal("msg send was called for unsupported source media")
	}
}

func TestMessagesForwardMessagesAssignsNewGroupedIDForGroupedMedia(t *testing.T) {
	groupedID := int64(99)
	var got *msg.TLMsgSendMessageV2
	c := newSendMsgCore(&messagesFakeMsgClient{
		getUserMessageList: func(context.Context, *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
			return &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{
				tg.MakeTLMessageBox(&tg.TLMessageBox{MessageId: 1, PeerType: payload.PeerTypeUser, PeerId: 300, Message: tg.MakeTLMessage(&tg.TLMessage{Id: 1, FromId: tg.MakePeerUser(300), PeerId: tg.MakePeerUser(100), Date: 1, Message: "one", GroupedId: &groupedID})}),
				tg.MakeTLMessageBox(&tg.TLMessageBox{MessageId: 2, PeerType: payload.PeerTypeUser, PeerId: 300, Message: tg.MakeTLMessage(&tg.TLMessage{Id: 2, FromId: tg.MakePeerUser(300), PeerId: tg.MakePeerUser(100), Date: 2, Message: "two", GroupedId: &groupedID})}),
			}}, nil
		},
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesForwardMessages(validForwardMessagesRequest())
	if err != nil {
		t.Fatalf("MessagesForwardMessages() error = %v", err)
	}
	first := assertForwardOutbox(t, got.Message[0], 11, "one")
	second := assertForwardOutbox(t, got.Message[1], 12, "two")
	if first.GroupedId == nil || second.GroupedId == nil || *first.GroupedId == 0 || *second.GroupedId != *first.GroupedId || *first.GroupedId == groupedID {
		t.Fatalf("grouped ids = %v/%v source=%d, want same new non-zero id", first.GroupedId, second.GroupedId, groupedID)
	}
}

func TestMessagesForwardMessagesDoesNotGroupPlainMessagesAfterGroupedMedia(t *testing.T) {
	groupedID := int64(99)
	var got *msg.TLMsgSendMessageV2
	c := newSendMsgCore(&messagesFakeMsgClient{
		getUserMessageList: func(context.Context, *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
			return &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{
				tg.MakeTLMessageBox(&tg.TLMessageBox{MessageId: 1, PeerType: payload.PeerTypeUser, PeerId: 300, Message: tg.MakeTLMessage(&tg.TLMessage{Id: 1, FromId: tg.MakePeerUser(300), PeerId: tg.MakePeerUser(100), Date: 1, Message: "grouped", GroupedId: &groupedID})}),
				messageBox(2, "plain"),
			}}, nil
		},
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesForwardMessages(validForwardMessagesRequest())
	if err != nil {
		t.Fatalf("MessagesForwardMessages() error = %v", err)
	}
	first := assertForwardOutbox(t, got.Message[0], 11, "grouped")
	second := assertForwardOutbox(t, got.Message[1], 12, "plain")
	if first.GroupedId == nil || *first.GroupedId == 0 || *first.GroupedId == groupedID {
		t.Fatalf("first grouped_id = %v source=%d, want new non-zero id", first.GroupedId, groupedID)
	}
	if second.GroupedId != nil {
		t.Fatalf("second grouped_id = %v, want nil for non-grouped source", second.GroupedId)
	}
}

func TestMessagesForwardMessagesMarksSavedPeerWhenForwardingToSelf(t *testing.T) {
	var got *msg.TLMsgSendMessageV2
	c := newSendMsgCore(&messagesFakeMsgClient{
		getUserMessageList: func(context.Context, *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
			return &msg.VectorMessageBox{Datas: []tg.MessageBoxClazz{
				tg.MakeTLMessageBox(&tg.TLMessageBox{
					MessageId: 1,
					PeerType:  payload.PeerTypeUser,
					PeerId:    300,
					Message: tg.MakeTLMessage(&tg.TLMessage{
						Id:      1,
						FromId:  tg.MakePeerUser(100),
						PeerId:  tg.MakePeerUser(300),
						Date:    1,
						Message: "saved outgoing",
					}),
				}),
			}}, nil
		},
		sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		},
	}, 100, 200)

	_, err := c.MessagesForwardMessages(&tg.TLMessagesForwardMessages{
		FromPeer: inputPeerUser(300),
		ToPeer:   inputPeerSelf(),
		Id:       []int32{1},
		RandomId: []int64{11},
	})
	if err != nil {
		t.Fatalf("MessagesForwardMessages() error = %v", err)
	}
	if got == nil || got.PeerId != 100 {
		t.Fatalf("send target = %#v, want self user 100", got)
	}
	message := assertForwardOutbox(t, got.Message[0], 11, "saved outgoing")
	if peerUserID(message.SavedPeerId) != 100 {
		t.Fatalf("SavedPeerId = %#v, want source sender peerUser(100)", message.SavedPeerId)
	}

	got = nil
	_, err = c.MessagesForwardMessages(&tg.TLMessagesForwardMessages{
		FromPeer: inputPeerUser(300),
		ToPeer:   inputPeerUser(400),
		Id:       []int32{1},
		RandomId: []int64{12},
	})
	if err != nil {
		t.Fatalf("MessagesForwardMessages(non-saved) error = %v", err)
	}
	message = assertForwardOutbox(t, got.Message[0], 12, "saved outgoing")
	if message.SavedPeerId != nil {
		t.Fatalf("SavedPeerId = %#v, want nil when target is not self", message.SavedPeerId)
	}
}

func validForwardMessagesRequest() *tg.TLMessagesForwardMessages {
	return &tg.TLMessagesForwardMessages{
		FromPeer: inputPeerUser(300),
		ToPeer:   inputPeerUser(400),
		Id:       []int32{1, 2},
		RandomId: []int64{11, 12},
	}
}

func messageBox(id int32, text string) tg.MessageBoxClazz {
	return tg.MakeTLMessageBox(&tg.TLMessageBox{
		MessageId: id,
		PeerType:  payload.PeerTypeUser,
		PeerId:    300,
		Message: tg.MakeTLMessage(&tg.TLMessage{
			Id:      id,
			FromId:  tg.MakePeerUser(300),
			PeerId:  tg.MakePeerUser(100),
			Date:    1700000000 + id,
			Message: text,
		}),
	})
}

func peerUserID(peer tg.PeerClazz) int64 {
	if p, ok := peer.(*tg.TLPeerUser); ok && p != nil {
		return p.UserId
	}
	return 0
}

func assertForwardOutbox(t *testing.T, outbox msg.OutboxMessageClazz, randomID int64, text string) *tg.TLMessage {
	t.Helper()
	if outbox == nil {
		t.Fatal("outbox is nil")
	}
	if outbox.RandomId != randomID {
		t.Fatalf("RandomId = %d, want %d", outbox.RandomId, randomID)
	}
	message, ok := outbox.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message type = %T, want *tg.TLMessage", outbox.Message)
	}
	if message.Message != text {
		t.Fatalf("Message = %q, want %q", message.Message, text)
	}
	return message
}
