package core

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	mediaclient "github.com/teamgram/teamgram-server/v2/app/service/media/client"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type messagesFakeMediaClient struct {
	mediaclient.MediaClient
	uploadPhotoFile func(ctx context.Context, in *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error)
}

func (f *messagesFakeMediaClient) MediaUploadPhotoFile(ctx context.Context, in *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
	return f.uploadPhotoFile(ctx, in)
}

func TestMessagesSendMediaBuildsOutboxAndReturnsFullUpdates(t *testing.T) {
	var got *msg.TLMsgSendMessage
	updates := testUpdates()
	resolved := tg.MakeTLPhoto(&tg.TLPhoto{Id: 777})
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return updates, nil
		}},
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, in *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			if in.OwnerId != 200 {
				t.Fatalf("OwnerId = %d, want auth key 200", in.OwnerId)
			}
			return resolved.ToPhoto(), nil
		}},
	}, 100, 200)

	entity := tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{Offset: 0, Length: 7})
	result, err := c.MessagesSendMedia(&tg.TLMessagesSendMedia{
		Silent:      true,
		Background:  true,
		ClearDraft:  true,
		Noforwards:  true,
		InvertMedia: true,
		Peer:        inputPeerUser(300),
		Media:       tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
		Message:     "caption",
		RandomId:    42,
		Entities:    []tg.MessageEntityClazz{entity},
		ReplyTo: tg.MakeTLInputReplyToMessage(&tg.TLInputReplyToMessage{
			ReplyToMsgId: 7,
		}),
	})
	if err != nil {
		t.Fatalf("MessagesSendMedia() error = %v", err)
	}
	if result != updates {
		t.Fatal("handler did not pass through tg.Updates")
	}
	if got == nil || len(got.Message) != 1 {
		t.Fatalf("MsgSendMessage input = %#v, want one outbox", got)
	}
	if !got.ClearDraft {
		t.Fatal("ClearDraft = false, want true")
	}
	if got.SourcePermAuthKeyId == nil || *got.SourcePermAuthKeyId != 200 {
		t.Fatalf("SourcePermAuthKeyId = %v, want 200", got.SourcePermAuthKeyId)
	}
	if got.ClearDraftBeforeDate == nil || *got.ClearDraftBeforeDate == 0 {
		t.Fatalf("ClearDraftBeforeDate = %v, want non-zero", got.ClearDraftBeforeDate)
	}
	outbox := got.Message[0]
	if outbox.RandomId != 42 {
		t.Fatalf("RandomId = %d, want 42", outbox.RandomId)
	}
	if !outbox.Background {
		t.Fatal("Background = false, want true")
	}
	message, ok := outbox.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message type = %T, want *tg.TLMessage", outbox.Message)
	}
	if !message.Out || !message.Silent || !message.Noforwards || !message.InvertMedia {
		t.Fatalf("message flags = out:%t silent:%t noforwards:%t invert:%t", message.Out, message.Silent, message.Noforwards, message.InvertMedia)
	}
	if message.Message != "caption" {
		t.Fatalf("Message = %q, want caption", message.Message)
	}
	if len(message.Entities) != 1 || message.Entities[0] != entity {
		t.Fatalf("Entities = %#v, want original entity", message.Entities)
	}
	if message.Media == nil {
		t.Fatal("Media is nil")
	}
	if _, ok := message.Media.(*tg.TLMessageMediaPhoto); !ok {
		t.Fatalf("Media = %T, want *tg.TLMessageMediaPhoto", message.Media)
	}
	if reply, ok := message.ReplyTo.(*tg.TLMessageReplyHeader); !ok || reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 7 {
		t.Fatalf("ReplyTo = %#v, want reply header msg 7", message.ReplyTo)
	}
}

func TestMessagesSendMediaAllowsInputPeerChat(t *testing.T) {
	var got *msg.TLMsgSendMessage
	var checked *chatpb.TLChatCheckMessageAction
	core := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: &messagesFakeChatClient{
			checkMessageAction: func(_ context.Context, in *chatpb.TLChatCheckMessageAction) (*chatpb.MessageActionCheckResult, error) {
				checked = in
				return chatpb.MakeTLMessageActionCheckResult(&chatpb.TLMessageActionCheckResult{
					SelfId: in.SelfId, ChatId: in.ChatId, Action: in.Action, MediaKind: in.MediaKind,
				}).ToMessageActionCheckResult(), nil
			},
		},
		MsgClient: &messagesFakeMsgClient{sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		}},
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			return tg.MakeTLPhoto(&tg.TLPhoto{Id: 777}).ToPhoto(), nil
		}},
	}, 1001, 9001)

	_, err := core.MessagesSendMedia(&tg.TLMessagesSendMedia{
		Peer:     inputPeerChat(55),
		Media:    tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
		Message:  "caption",
		RandomId: 42,
	})
	if err != nil {
		t.Fatalf("MessagesSendMedia() error = %v", err)
	}
	if checked == nil || checked.ChatId != 55 || checked.Action != chatpb.MessageActionSendMediaPhoto || checked.MediaKind != "photo" {
		t.Fatalf("chat check = %+v, want send_media_photo for chat 55", checked)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 {
		t.Fatalf("msg request = %+v, want PeerTypeChat/chat 55", got)
	}
	message, ok := got.Message[0].Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message type = %T, want *tg.TLMessage", got.Message[0].Message)
	}
	if peer, ok := message.PeerId.(*tg.TLPeerChat); !ok || peer.ChatId != 55 {
		t.Fatalf("outbox peer = %#v, want peerChat 55", message.PeerId)
	}
}

func TestMessagesSendMediaProjectsUsersIntoUpdates(t *testing.T) {
	updates := testUpdates()
	updates.Clazz.(*tg.TLUpdates).Updates = []tg.UpdateClazz{
		tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: 61, RandomId: 42}),
		tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{
				Id:      61,
				FromId:  tg.MakePeerUser(100),
				PeerId:  tg.MakePeerUser(300),
				Message: "caption",
			}),
			Pts:      31,
			PtsCount: 1,
		}),
	}
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return updates, nil
		}},
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			return tg.MakeTLPhoto(&tg.TLPhoto{Id: 777}).ToPhoto(), nil
		}},
		UserClient: &messagesFakeUserClient{
			projectUsers: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				if len(in.ViewerUserIds) != 1 || in.ViewerUserIds[0] != 100 {
					t.Fatalf("ViewerUserIds = %v, want [100]", in.ViewerUserIds)
				}
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 100, Self: true}),
							tg.MakeTLUser(&tg.TLUser{Id: 300, Contact: true}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
			},
		},
	}, 100, 200)

	got, err := c.MessagesSendMedia(validSendMediaRequest())
	if err != nil {
		t.Fatalf("MessagesSendMedia() error = %v", err)
	}
	full := got.Clazz.(*tg.TLUpdates)
	if len(full.Updates) != 2 {
		t.Fatalf("updates len = %d, want 2", len(full.Updates))
	}
	if len(full.Users) != 2 {
		t.Fatalf("users len = %d, want 2", len(full.Users))
	}
}

func TestMessagesSendMediaRejectsUnsupportedFields(t *testing.T) {
	scheduleDate := int32(2000000)
	scheduleRepeat := int32(60)
	effect := int64(10)
	stars := int64(1)
	tests := []struct {
		name  string
		patch func(*tg.TLMessagesSendMedia)
	}{
		{name: "reply markup", patch: func(in *tg.TLMessagesSendMedia) {
			in.ReplyMarkup = tg.MakeTLReplyKeyboardMarkup(&tg.TLReplyKeyboardMarkup{})
		}},
		{name: "schedule date", patch: func(in *tg.TLMessagesSendMedia) { in.ScheduleDate = &scheduleDate }},
		{name: "schedule repeat", patch: func(in *tg.TLMessagesSendMedia) { in.ScheduleRepeatPeriod = &scheduleRepeat }},
		{name: "update stickersets order", patch: func(in *tg.TLMessagesSendMedia) { in.UpdateStickersetsOrder = true }},
		{name: "send as", patch: func(in *tg.TLMessagesSendMedia) { in.SendAs = inputPeerUser(400) }},
		{name: "quick reply shortcut", patch: func(in *tg.TLMessagesSendMedia) {
			in.QuickReplyShortcut = tg.MakeTLInputQuickReplyShortcut(&tg.TLInputQuickReplyShortcut{})
		}},
		{name: "effect", patch: func(in *tg.TLMessagesSendMedia) { in.Effect = &effect }},
		{name: "allow paid floodskip", patch: func(in *tg.TLMessagesSendMedia) { in.AllowPaidFloodskip = true }},
		{name: "allow paid stars", patch: func(in *tg.TLMessagesSendMedia) { in.AllowPaidStars = &stars }},
		{name: "suggested post", patch: func(in *tg.TLMessagesSendMedia) { in.SuggestedPost = tg.MakeTLSuggestedPost(&tg.TLSuggestedPost{}) }},
		{name: "reply to story", patch: func(in *tg.TLMessagesSendMedia) { in.ReplyTo = tg.MakeTLInputReplyToStory(&tg.TLInputReplyToStory{}) }},
		{name: "reply to peer", patch: func(in *tg.TLMessagesSendMedia) {
			in.ReplyTo = tg.MakeTLInputReplyToMessage(&tg.TLInputReplyToMessage{
				ReplyToMsgId:  1,
				ReplyToPeerId: inputPeerUser(400),
			})
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			c := newMessagesCoreWithRepo(&repository.Repository{
				MsgClient: &messagesFakeMsgClient{sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
					called = true
					return nil, nil
				}},
				MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
					called = true
					return nil, nil
				}},
			}, 100, 200)
			in := validSendMediaRequest()
			tt.patch(in)
			_, err := c.MessagesSendMedia(in)
			if err != tg.ErrInputRequestInvalid && err != tg.ErrReplyToInvalid {
				t.Fatalf("error = %v, want unsupported field error", err)
			}
			if called {
				t.Fatal("downstream service was called but should not have been")
			}
		})
	}
}

func TestMessagesSendMediaMapsInvalidMediaToMediaEmpty(t *testing.T) {
	c := newMessagesCoreWithRepo(&repository.Repository{
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			return nil, mediapb.ErrMediaInvalidArgument
		}},
	}, 100, 200)
	_, err := c.MessagesSendMedia(validSendMediaRequest())
	if !errors.Is(err, tg.ErrMediaEmpty) {
		t.Fatalf("error = %v, want MEDIA_EMPTY", err)
	}
}

func TestMessagesSendMediaRejectsUnsupportedMediaBeforeMsgSend(t *testing.T) {
	called := false
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{sendMessage: func(context.Context, *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		}},
	}, 100, 200)
	in := validSendMediaRequest()
	in.Media = tg.MakeTLInputMediaGeoPoint(&tg.TLInputMediaGeoPoint{
		GeoPoint: tg.MakeTLInputGeoPoint(&tg.TLInputGeoPoint{Lat: 1, Long: 2}),
	})
	_, err := c.MessagesSendMedia(in)
	if !errors.Is(err, tg.ErrMediaEmpty) {
		t.Fatalf("error = %v, want MEDIA_EMPTY", err)
	}
	if called {
		t.Fatal("msg send was called for unsupported media")
	}
}

func TestMessagesSendMediaSupportsContactMedia(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		}},
		UserClient: &messagesFakeUserClient{
			getUserIDByPhone: func(_ context.Context, in *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error) {
				if in.Phone != "8613000000001" {
					t.Fatalf("phone lookup = %q, want normalized contact phone", in.Phone)
				}
				return tg.MakeTLInt64(&tg.TLInt64{V: 1571266964}), nil
			},
		},
	}, 100, 200)
	in := validSendMediaRequest()
	in.Media = tg.MakeTLInputMediaContact(&tg.TLInputMediaContact{
		PhoneNumber: "8613000000001",
		FirstName:   "13000000001",
		LastName:    "t2",
		Vcard:       "",
	})

	_, err := c.MessagesSendMedia(in)
	if err != nil {
		t.Fatalf("MessagesSendMedia() error = %v", err)
	}
	if got == nil || len(got.Message) != 1 {
		t.Fatalf("MsgSendMessage input = %#v, want one outbox", got)
	}
	message, ok := got.Message[0].Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message type = %T, want *tg.TLMessage", got.Message[0].Message)
	}
	contact, ok := message.Media.(*tg.TLMessageMediaContact)
	if !ok {
		t.Fatalf("Media = %T, want *tg.TLMessageMediaContact", message.Media)
	}
	if contact.PhoneNumber != "8613000000001" || contact.FirstName != "13000000001" || contact.LastName != "t2" || contact.UserId != 1571266964 {
		t.Fatalf("contact media = %#v, want shared contact fields and user id", contact)
	}
}

func TestMessagesSendMediaMapsMediaInfraFailureToInternal(t *testing.T) {
	c := newMessagesCoreWithRepo(&repository.Repository{
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			return nil, mediapb.ErrMediaStorage
		}},
	}, 100, 200)
	_, err := c.MessagesSendMedia(validSendMediaRequest())
	if !errors.Is(err, tg.ErrInternalServerError) {
		t.Fatalf("error = %v, want INTERNAL", err)
	}
	if errors.Is(err, mediapb.ErrMediaStorage) {
		t.Fatalf("error leaked media domain error: %v", err)
	}
}

func TestMessagesSendMediaEmptyCaptionAllowed(t *testing.T) {
	var got *msg.TLMsgSendMessage
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{sendMessage: func(_ context.Context, in *msg.TLMsgSendMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		}},
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			return tg.MakeTLPhoto(&tg.TLPhoto{Id: 777}).ToPhoto(), nil
		}},
	}, 100, 200)
	in := validSendMediaRequest()
	in.Message = ""
	_, err := c.MessagesSendMedia(in)
	if err != nil {
		t.Fatalf("MessagesSendMedia() error = %v", err)
	}
	message, ok := got.Message[0].Message.(*tg.TLMessage)
	if !ok || message.Message != "" {
		t.Fatalf("outbox message = %#v, want empty caption", got.Message[0].Message)
	}
}

func TestMessagesSendMediaCaptionTooLongRejected(t *testing.T) {
	called := false
	c := newMessagesCoreWithRepo(&repository.Repository{
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			called = true
			return nil, nil
		}},
	}, 100, 200)
	in := validSendMediaRequest()
	in.Message = strings.Repeat("a", 4097)
	_, err := c.MessagesSendMedia(in)
	if err != tg.ErrMediaCaptionTooLong {
		t.Fatalf("error = %v, want MEDIA_CAPTION_TOO_LONG", err)
	}
	if called {
		t.Fatal("media service was called but should not have been")
	}
}

func TestMessagesSendMessage_MsgIdInvalidMapped(t *testing.T) {
	c := newSendMsgCore(&messagesFakeMsgClient{
		sendMessage: func(_ context.Context, _ *msg.TLMsgSendMessage) (*tg.Updates, error) {
			return nil, msg.ErrMsgIdInvalid
		},
	}, 100, 200)

	_, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     inputPeerUser(300),
		Message:  "hello",
		RandomId: 42,
	})
	if err != tg.ErrMessageIdInvalid {
		t.Fatalf("error = %v, want %v", err, tg.ErrMessageIdInvalid)
	}
}

func validSendMediaRequest() *tg.TLMessagesSendMedia {
	return &tg.TLMessagesSendMedia{
		Peer:     inputPeerUser(300),
		Media:    tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
		Message:  "caption",
		RandomId: 42,
	}
}
