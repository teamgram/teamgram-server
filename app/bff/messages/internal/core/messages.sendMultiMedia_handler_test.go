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
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesSendMultiMediaBuildsGroupedBatchAndReturnsFullUpdates(t *testing.T) {
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

	entity := tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{Offset: 0, Length: 3})
	result, err := c.MessagesSendMultiMedia(&tg.TLMessagesSendMultiMedia{
		Silent:      true,
		Background:  true,
		ClearDraft:  true,
		Noforwards:  true,
		InvertMedia: true,
		Peer:        inputPeerUser(300),
		ReplyTo: tg.MakeTLInputReplyToMessage(&tg.TLInputReplyToMessage{
			ReplyToMsgId: 7,
		}),
		MultiMedia: []tg.InputSingleMediaClazz{
			tg.MakeTLInputSingleMedia(&tg.TLInputSingleMedia{
				Media:    tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
				RandomId: 11,
				Message:  "one",
				Entities: []tg.MessageEntityClazz{
					entity,
				},
			}),
			tg.MakeTLInputSingleMedia(&tg.TLInputSingleMedia{
				Media:    tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
				RandomId: 12,
				Message:  "two",
			}),
		},
	})
	if err != nil {
		t.Fatalf("MessagesSendMultiMedia() error = %v", err)
	}
	if result != updates {
		t.Fatal("handler did not pass through tg.Updates")
	}
	if got == nil || len(got.Message) != 2 {
		t.Fatalf("MsgSendMessage input = %#v, want two outboxes", got)
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
	first := assertMultiMediaOutbox(t, got.Message[0], 11, "one", entity)
	second := assertMultiMediaOutbox(t, got.Message[1], 12, "two")
	if !got.Message[0].Background || !got.Message[1].Background {
		t.Fatalf("background flags = %t/%t, want true", got.Message[0].Background, got.Message[1].Background)
	}
	if first.GroupedId == nil || *first.GroupedId == 0 {
		t.Fatalf("first grouped_id = %v, want non-zero", first.GroupedId)
	}
	if second.GroupedId == nil || *second.GroupedId != *first.GroupedId {
		t.Fatalf("grouped ids = %v/%v, want same non-zero", first.GroupedId, second.GroupedId)
	}
	for i, message := range []*tg.TLMessage{first, second} {
		if !message.Out || !message.Silent || !message.Noforwards || !message.InvertMedia {
			t.Fatalf("message %d flags = out:%t silent:%t noforwards:%t invert:%t", i, message.Out, message.Silent, message.Noforwards, message.InvertMedia)
		}
		if _, ok := message.Media.(*tg.TLMessageMediaPhoto); !ok {
			t.Fatalf("message %d media = %T, want *tg.TLMessageMediaPhoto", i, message.Media)
		}
		if reply, ok := message.ReplyTo.(*tg.TLMessageReplyHeader); !ok || reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 7 {
			t.Fatalf("message %d ReplyTo = %#v, want reply header msg 7", i, message.ReplyTo)
		}
	}
}

func TestMessagesSendMultiMediaAllowsInputPeerChat(t *testing.T) {
	var got *msg.TLMsgSendMessage
	var checks []*chatpb.TLChatCheckMessageAction
	core := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: &messagesFakeChatClient{
			checkMessageAction: func(_ context.Context, in *chatpb.TLChatCheckMessageAction) (*chatpb.MessageActionCheckResult, error) {
				checks = append(checks, in)
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
	in := validSendMultiMediaRequest()
	in.Peer = inputPeerChat(55)

	_, err := core.MessagesSendMultiMedia(in)
	if err != nil {
		t.Fatalf("MessagesSendMultiMedia() error = %v", err)
	}
	if len(checks) != 2 ||
		checks[0].ChatId != 55 || checks[0].Action != chatpb.MessageActionSendAlbum || checks[0].MediaKind != "album" ||
		checks[1].ChatId != 55 || checks[1].Action != chatpb.MessageActionSendMediaPhoto || checks[1].MediaKind != "photo" {
		t.Fatalf("chat checks = %+v, want send_album then send_media_photo", checks)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 {
		t.Fatalf("msg request = %+v, want PeerTypeChat/chat 55", got)
	}
	first := assertMultiMediaOutbox(t, got.Message[0], 11, "one")
	second := assertMultiMediaOutbox(t, got.Message[1], 12, "two")
	if peer, ok := first.PeerId.(*tg.TLPeerChat); !ok || peer.ChatId != 55 {
		t.Fatalf("first peer = %#v, want peerChat 55", first.PeerId)
	}
	if peer, ok := second.PeerId.(*tg.TLPeerChat); !ok || peer.ChatId != 55 {
		t.Fatalf("second peer = %#v, want peerChat 55", second.PeerId)
	}
}

func TestMessagesSendMultiMediaEmptyCaptionAllowed(t *testing.T) {
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
	in := validSendMultiMediaRequest()
	in.MultiMedia[0].Message = ""
	_, err := c.MessagesSendMultiMedia(in)
	if err != nil {
		t.Fatalf("MessagesSendMultiMedia() error = %v", err)
	}
	message := assertMultiMediaOutbox(t, got.Message[0], 11, "")
	if message.Message != "" {
		t.Fatalf("Message = %q, want empty caption", message.Message)
	}
}

func TestMessagesSendMultiMediaCaptionTooLongRejected(t *testing.T) {
	called := false
	c := newMessagesCoreWithRepo(&repository.Repository{
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			called = true
			return nil, nil
		}},
	}, 100, 200)
	in := validSendMultiMediaRequest()
	in.MultiMedia[0].Message = strings.Repeat("a", 4097)
	_, err := c.MessagesSendMultiMedia(in)
	if err != tg.ErrMediaCaptionTooLong {
		t.Fatalf("error = %v, want MEDIA_CAPTION_TOO_LONG", err)
	}
	if called {
		t.Fatal("media service was called but should not have been")
	}
}

func TestMessagesSendMultiMediaPreflightsLaterCaptionBeforeMediaResolve(t *testing.T) {
	called := false
	c := newMessagesCoreWithRepo(&repository.Repository{
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			called = true
			return tg.MakeTLPhoto(&tg.TLPhoto{Id: 777}).ToPhoto(), nil
		}},
	}, 100, 200)
	in := validSendMultiMediaRequest()
	in.MultiMedia[1].Message = strings.Repeat("a", 4097)
	_, err := c.MessagesSendMultiMedia(in)
	if err != tg.ErrMediaCaptionTooLong {
		t.Fatalf("error = %v, want MEDIA_CAPTION_TOO_LONG", err)
	}
	if called {
		t.Fatal("media service was called before all captions were validated")
	}
}

func TestMessagesSendMultiMediaRejectsInvalidBatchSize(t *testing.T) {
	tests := []struct {
		name       string
		multimedia []tg.InputSingleMediaClazz
		want       error
	}{
		{name: "empty", multimedia: nil, want: tg.ErrMediaEmpty},
		{name: "too long", multimedia: make([]tg.InputSingleMediaClazz, maxSendMultiMediaItems+1), want: tg.ErrMultiMediaTooLong},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newMessagesCoreWithRepo(&repository.Repository{}, 100, 200)
			in := validSendMultiMediaRequest()
			in.MultiMedia = tt.multimedia
			for i := range in.MultiMedia {
				in.MultiMedia[i] = tg.MakeTLInputSingleMedia(&tg.TLInputSingleMedia{
					Media:    tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
					RandomId: int64(i + 1),
					Message:  "caption",
				})
			}
			_, err := c.MessagesSendMultiMedia(in)
			if err != tt.want {
				t.Fatalf("error = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestMessagesSendMultiMediaRejectsZeroAndDuplicateRandomID(t *testing.T) {
	tests := []struct {
		name  string
		patch func(*tg.TLMessagesSendMultiMedia)
		want  error
	}{
		{name: "zero", patch: func(in *tg.TLMessagesSendMultiMedia) { in.MultiMedia[0].RandomId = 0 }, want: tg.ErrRandomIdEmpty},
		{name: "duplicate", patch: func(in *tg.TLMessagesSendMultiMedia) { in.MultiMedia[1].RandomId = in.MultiMedia[0].RandomId }, want: tg.ErrRandomIdDuplicate},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newMessagesCoreWithRepo(&repository.Repository{}, 100, 200)
			in := validSendMultiMediaRequest()
			tt.patch(in)
			_, err := c.MessagesSendMultiMedia(in)
			if err != tt.want {
				t.Fatalf("error = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestMessagesSendMultiMediaRejectsUnsupportedFields(t *testing.T) {
	scheduleDate := int32(2000000)
	effect := int64(10)
	stars := int64(1)
	tests := []struct {
		name  string
		patch func(*tg.TLMessagesSendMultiMedia)
	}{
		{name: "schedule date", patch: func(in *tg.TLMessagesSendMultiMedia) { in.ScheduleDate = &scheduleDate }},
		{name: "update stickersets order", patch: func(in *tg.TLMessagesSendMultiMedia) { in.UpdateStickersetsOrder = true }},
		{name: "send as", patch: func(in *tg.TLMessagesSendMultiMedia) { in.SendAs = inputPeerUser(400) }},
		{name: "quick reply shortcut", patch: func(in *tg.TLMessagesSendMultiMedia) {
			in.QuickReplyShortcut = tg.MakeTLInputQuickReplyShortcut(&tg.TLInputQuickReplyShortcut{})
		}},
		{name: "effect", patch: func(in *tg.TLMessagesSendMultiMedia) { in.Effect = &effect }},
		{name: "allow paid floodskip", patch: func(in *tg.TLMessagesSendMultiMedia) { in.AllowPaidFloodskip = true }},
		{name: "allow paid stars", patch: func(in *tg.TLMessagesSendMultiMedia) { in.AllowPaidStars = &stars }},
		{name: "reply to story", patch: func(in *tg.TLMessagesSendMultiMedia) {
			in.ReplyTo = tg.MakeTLInputReplyToStory(&tg.TLInputReplyToStory{})
		}},
		{name: "reply to peer", patch: func(in *tg.TLMessagesSendMultiMedia) {
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
			in := validSendMultiMediaRequest()
			tt.patch(in)
			_, err := c.MessagesSendMultiMedia(in)
			if err != tg.ErrInputRequestInvalid && err != tg.ErrReplyToInvalid {
				t.Fatalf("error = %v, want unsupported field error", err)
			}
			if called {
				t.Fatal("downstream service was called but should not have been")
			}
		})
	}
}

func TestMessagesSendMultiMediaMapsInvalidMediaToMediaEmpty(t *testing.T) {
	c := newMessagesCoreWithRepo(&repository.Repository{
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			return nil, mediapb.ErrMediaInvalidArgument
		}},
	}, 100, 200)
	_, err := c.MessagesSendMultiMedia(validSendMultiMediaRequest())
	if !errors.Is(err, tg.ErrMediaEmpty) {
		t.Fatalf("error = %v, want MEDIA_EMPTY", err)
	}
}

func TestMessagesSendMultiMediaRejectsUnsupportedMediaBeforeMsgSend(t *testing.T) {
	called := false
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{sendMessage: func(context.Context, *msg.TLMsgSendMessage) (*tg.Updates, error) {
			called = true
			return nil, nil
		}},
	}, 100, 200)
	in := validSendMultiMediaRequest()
	in.MultiMedia[0].Media = tg.MakeTLInputMediaGeoPoint(&tg.TLInputMediaGeoPoint{
		GeoPoint: tg.MakeTLInputGeoPoint(&tg.TLInputGeoPoint{Lat: 1, Long: 2}),
	})
	_, err := c.MessagesSendMultiMedia(in)
	if !errors.Is(err, tg.ErrMediaEmpty) {
		t.Fatalf("error = %v, want MEDIA_EMPTY", err)
	}
	if called {
		t.Fatal("msg send was called for unsupported media")
	}
}

func TestMessagesSendMultiMediaTypedNilMediaRejected(t *testing.T) {
	called := false
	c := newMessagesCoreWithRepo(&repository.Repository{
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			called = true
			return nil, nil
		}},
	}, 100, 200)
	in := validSendMultiMediaRequest()
	var media *tg.TLInputMediaUploadedPhoto
	in.MultiMedia[0].Media = media
	_, err := c.MessagesSendMultiMedia(in)
	if !errors.Is(err, tg.ErrMediaEmpty) {
		t.Fatalf("error = %v, want MEDIA_EMPTY", err)
	}
	if called {
		t.Fatal("media service was called for typed-nil media")
	}
}

func TestMessagesSendMultiMediaMapsMediaInfraFailureToInternal(t *testing.T) {
	c := newMessagesCoreWithRepo(&repository.Repository{
		MediaClient: &messagesFakeMediaClient{uploadPhotoFile: func(_ context.Context, _ *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
			return nil, mediapb.ErrMediaStorage
		}},
	}, 100, 200)
	_, err := c.MessagesSendMultiMedia(validSendMultiMediaRequest())
	if !errors.Is(err, tg.ErrInternalServerError) {
		t.Fatalf("error = %v, want INTERNAL", err)
	}
	if errors.Is(err, mediapb.ErrMediaStorage) {
		t.Fatalf("error leaked media domain error: %v", err)
	}
}

func validSendMultiMediaRequest() *tg.TLMessagesSendMultiMedia {
	return &tg.TLMessagesSendMultiMedia{
		Peer: inputPeerUser(300),
		MultiMedia: []tg.InputSingleMediaClazz{
			tg.MakeTLInputSingleMedia(&tg.TLInputSingleMedia{
				Media:    tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
				RandomId: 11,
				Message:  "one",
			}),
			tg.MakeTLInputSingleMedia(&tg.TLInputSingleMedia{
				Media:    tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
				RandomId: 12,
				Message:  "two",
			}),
		},
	}
}

func assertMultiMediaOutbox(t *testing.T, outbox *msg.TLOutboxMessage, randomID int64, caption string, entities ...tg.MessageEntityClazz) *tg.TLMessage {
	t.Helper()
	if outbox.RandomId != randomID {
		t.Fatalf("RandomId = %d, want %d", outbox.RandomId, randomID)
	}
	message, ok := outbox.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("outbox message type = %T, want *tg.TLMessage", outbox.Message)
	}
	if message.Message != caption {
		t.Fatalf("Message = %q, want %q", message.Message, caption)
	}
	if len(message.Entities) != len(entities) {
		t.Fatalf("Entities = %#v, want %#v", message.Entities, entities)
	}
	for i := range entities {
		if message.Entities[i] != entities[i] {
			t.Fatalf("Entities[%d] = %#v, want %#v", i, message.Entities[i], entities[i])
		}
	}
	return message
}
