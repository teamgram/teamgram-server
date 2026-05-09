package core

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
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
	var got *msg.TLMsgSendMessageV2
	updates := testUpdates()
	resolved := tg.MakeTLPhoto(&tg.TLPhoto{Id: 777})
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
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
		t.Fatalf("MsgSendMessageV2 input = %#v, want one outbox", got)
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
				MsgClient: &messagesFakeMsgClient{sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
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
	var got *msg.TLMsgSendMessageV2
	c := newMessagesCoreWithRepo(&repository.Repository{
		MsgClient: &messagesFakeMsgClient{sendMessageV2: func(_ context.Context, in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
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
		sendMessageV2: func(_ context.Context, _ *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
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
