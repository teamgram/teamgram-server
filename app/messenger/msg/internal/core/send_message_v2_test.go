package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/svc"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/pagination"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestNormalizeOutboxMessageUsesAttrsGroupedID(t *testing.T) {
	groupedID := int64(12345)
	message := tg.MakeTLMessage(&tg.TLMessage{
		Message:   "caption",
		GroupedId: &groupedID,
		Media:     tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{}),
	})
	outbox := msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{RandomId: 99, Message: message})
	got, err := normalizeOutboxMessage(normalizeOutboxInput{
		SenderUserID: 100,
		PeerType:     payload.PeerTypeUser,
		PeerID:       200,
		Outbox:       outbox,
	})
	if err != nil {
		t.Fatalf("normalizeOutboxMessage() error = %v", err)
	}
	if got.Attrs.GroupedID != 12345 {
		t.Fatalf("GroupedID = %d, want 12345", got.Attrs.GroupedID)
	}
}

func TestNormalizeOutboxMessageSupportsChatCreateServiceAction(t *testing.T) {
	message := tg.MakeTLMessageService(&tg.TLMessageService{
		Out:    true,
		FromId: tg.MakePeerUser(1001),
		PeerId: tg.MakePeerChat(55),
		Date:   1_778_648_899,
		Action: tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
			Title: "new chat",
			Users: []int64{1002, 1003},
		}),
	})
	outbox := msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{RandomId: 99, Message: message})
	got, err := normalizeOutboxMessage(normalizeOutboxInput{
		SenderUserID: 1001,
		PeerType:     payload.PeerTypeChat,
		PeerID:       55,
		Outbox:       outbox,
	})
	if err != nil {
		t.Fatalf("normalizeOutboxMessage() error = %v", err)
	}
	if got.ServiceAction == nil || got.ServiceAction.Kind != payload.ServiceActionKindChatCreate {
		t.Fatalf("ServiceAction = %+v, want chat create", got.ServiceAction)
	}
	if got.ServiceAction.Title != "new chat" || len(got.ServiceAction.Users) != 2 || got.ServiceAction.Users[0] != 1002 {
		t.Fatalf("ServiceAction = %+v, want title/users preserved", got.ServiceAction)
	}
}

func TestNormalizeOutboxMessageSupportsChatAddUserServiceAction(t *testing.T) {
	message := tg.MakeTLMessageService(&tg.TLMessageService{
		Out:    true,
		FromId: tg.MakePeerUser(1001),
		PeerId: tg.MakePeerChat(55),
		Date:   1_778_648_899,
		Action: tg.MakeTLMessageActionChatAddUser(&tg.TLMessageActionChatAddUser{
			Users: []int64{1002},
		}),
	})
	outbox := msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{RandomId: 99, Message: message})
	got, err := normalizeOutboxMessage(normalizeOutboxInput{
		SenderUserID: 1001,
		PeerType:     payload.PeerTypeChat,
		PeerID:       55,
		Outbox:       outbox,
	})
	if err != nil {
		t.Fatalf("normalizeOutboxMessage() error = %v", err)
	}
	if got.ServiceAction == nil || got.ServiceAction.Kind != payload.ServiceActionKindChatAddUser {
		t.Fatalf("ServiceAction = %+v, want chat add user", got.ServiceAction)
	}
	if len(got.ServiceAction.Users) != 1 || got.ServiceAction.Users[0] != 1002 {
		t.Fatalf("ServiceAction users = %+v, want [1002]", got.ServiceAction.Users)
	}
}

func TestNormalizeMediaRefDocumentPreservesV2DocumentPayload(t *testing.T) {
	videoStartTs := 1.25
	videoTimestamp := int32(7)
	media := tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
		Round:          true,
		Spoiler:        true,
		VideoTimestamp: &videoTimestamp,
		VideoCover: tg.MakeTLPhoto(&tg.TLPhoto{
			Id:            777,
			AccessHash:    888,
			FileReference: []byte("cover-ref"),
			Date:          1_700_000_001,
			DcId:          5,
			Sizes: []tg.PhotoSizeClazz{
				tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "x", W: 640, H: 360, Size2: 4321}),
			},
			VideoSizes: []tg.VideoSizeClazz{
				tg.MakeTLVideoSizeEmojiMarkup(&tg.TLVideoSizeEmojiMarkup{EmojiId: 9009, BackgroundColors: []int32{1, 2}}),
			},
		}),
		Document: tg.MakeTLDocument(&tg.TLDocument{
			Id:            555,
			AccessHash:    666,
			FileReference: []byte("doc-ref"),
			Date:          1_700_000_000,
			DcId:          4,
			MimeType:      "video/mp4",
			Size2:         98765,
			Thumbs: []tg.PhotoSizeClazz{
				tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 200, Size2: 1234}),
				tg.MakeTLPhotoPathSize(&tg.TLPhotoPathSize{Type: "j", Bytes: []byte{9, 8, 7}}),
			},
			VideoThumbs: []tg.VideoSizeClazz{
				tg.MakeTLVideoSize(&tg.TLVideoSize{Type: "v", W: 320, H: 200, Size2: 4567, VideoStartTs: &videoStartTs}),
				tg.MakeTLVideoSizeStickerMarkup(&tg.TLVideoSizeStickerMarkup{
					Stickerset:       tg.MakeTLInputStickerSetDice(&tg.TLInputStickerSetDice{Emoticon: "🎲"}),
					StickerId:        8080,
					BackgroundColors: []int32{3, 4},
				}),
			},
			Attributes: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeSticker(&tg.TLDocumentAttributeSticker{
					Alt:        ":)",
					Mask:       true,
					Stickerset: tg.MakeTLInputStickerSetID(&tg.TLInputStickerSetID{Id: 1001, AccessHash: 2002}),
					MaskCoords: tg.MakeTLMaskCoords(&tg.TLMaskCoords{N: 1, X: 0.5, Y: 0.25, Zoom: 1.5}),
				}),
				tg.MakeTLDocumentAttributeCustomEmoji(&tg.TLDocumentAttributeCustomEmoji{
					Alt:        ":)",
					Free:       true,
					TextColor:  true,
					Stickerset: tg.MakeTLInputStickerSetDice(&tg.TLInputStickerSetDice{Emoticon: "🎯"}),
				}),
				tg.MakeTLDocumentAttributeHasStickers(&tg.TLDocumentAttributeHasStickers{}),
			},
		}),
	})
	got, err := normalizeMediaRef(media)
	if err != nil {
		t.Fatalf("normalizeMediaRef() error = %v", err)
	}
	if got.SchemaVersion != payload.MediaRefSchemaVersionV2 || got.Kind != "document" || got.ID != 555 ||
		got.AccessHash != 666 || string(got.FileReference) != "doc-ref" || got.MimeType != "video/mp4" || got.Size != 98765 {
		t.Fatalf("media ref identity = %+v, want full V2 document identity", got)
	}
	if got.DocumentMediaFlags == nil || got.DocumentMediaFlags.Video || !got.DocumentMediaFlags.Round || got.DocumentMediaFlags.Voice || !got.DocumentMediaFlags.Spoiler {
		t.Fatalf("DocumentMediaFlags = %+v, want explicit false/true values", got.DocumentMediaFlags)
	}
	if got.VideoTimestamp == nil || *got.VideoTimestamp != videoTimestamp {
		t.Fatalf("VideoTimestamp = %v, want %d", got.VideoTimestamp, videoTimestamp)
	}
	if got.VideoCover == nil || got.VideoCover.ID != 777 || len(got.VideoCover.VideoSizes) != 1 || got.VideoCover.VideoSizes[0].EmojiID != 9009 {
		t.Fatalf("VideoCover = %+v, want full photo ref with video sizes", got.VideoCover)
	}
	if len(got.DocumentThumbs) != 2 || got.DocumentThumbs[0].Type != "m" {
		t.Fatalf("DocumentThumbs = %+v, want photo thumb m plus path", got.DocumentThumbs)
	}
	if got.DocumentThumbs[1].Kind != "path" || got.DocumentThumbs[1].Type != "j" || string(got.DocumentThumbs[1].Bytes) != string([]byte{9, 8, 7}) {
		t.Fatalf("DocumentThumbs[1] = %+v, want path j bytes", got.DocumentThumbs[1])
	}
	if len(got.DocumentVideoThumbs) != 2 || got.DocumentVideoThumbs[0].VideoStartTs == nil || *got.DocumentVideoThumbs[0].VideoStartTs != videoStartTs {
		t.Fatalf("DocumentVideoThumbs = %+v, want video thumb with start ts", got.DocumentVideoThumbs)
	}
	if got.DocumentVideoThumbs[1].StickerSet == nil || got.DocumentVideoThumbs[1].StickerSet.Kind != "dice" ||
		got.DocumentVideoThumbs[1].StickerSet.Emoticon != "🎲" || got.DocumentVideoThumbs[1].StickerID != 8080 {
		t.Fatalf("DocumentVideoThumbs[1] = %+v, want sticker markup", got.DocumentVideoThumbs[1])
	}
	if len(got.DocumentAttributes) != 3 {
		t.Fatalf("DocumentAttributes len = %d, want 3", len(got.DocumentAttributes))
	}
	sticker := got.DocumentAttributes[0]
	if sticker.Kind != "sticker" || sticker.Alt != ":)" || sticker.StickerSet == nil || sticker.StickerSet.ID != 1001 || !sticker.Mask || sticker.MaskCoords == nil {
		t.Fatalf("sticker attr = %+v, want full sticker metadata", sticker)
	}
	customEmoji := got.DocumentAttributes[1]
	if customEmoji.Kind != "custom_emoji" || customEmoji.StickerSet == nil ||
		customEmoji.StickerSet.Kind != "dice" || customEmoji.StickerSet.Emoticon != "🎯" || !customEmoji.Free || !customEmoji.TextColor {
		t.Fatalf("custom emoji attr = %+v, want full custom emoji metadata", customEmoji)
	}
	if got.DocumentAttributes[2].Kind != "has_stickers" {
		t.Fatalf("DocumentAttributes[2] = %+v, want has_stickers", got.DocumentAttributes[2])
	}
}

func TestNormalizeOutboxMessageRejectsUnsupportedEntity(t *testing.T) {
	message := tg.MakeTLMessage(&tg.TLMessage{
		Message: "emoji",
		Entities: []tg.MessageEntityClazz{
			tg.MakeTLMessageEntityCustomEmoji(&tg.TLMessageEntityCustomEmoji{
				Offset:     0,
				Length:     5,
				DocumentId: 12345,
			}),
		},
	})
	outbox := msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{RandomId: 99, Message: message})
	_, err := normalizeOutboxMessage(normalizeOutboxInput{
		SenderUserID: 100,
		PeerType:     payload.PeerTypeUser,
		PeerID:       200,
		Outbox:       outbox,
	})
	if !errors.Is(err, msgpb.ErrSendStateConflict) {
		t.Fatalf("normalizeOutboxMessage() error = %v, want ErrSendStateConflict", err)
	}
}

func TestNormalizeOutboxMessageSupportsMentionNameEntity(t *testing.T) {
	message := tg.MakeTLMessage(&tg.TLMessage{
		Message: "user",
		Entities: []tg.MessageEntityClazz{
			tg.MakeTLMessageEntityMentionName(&tg.TLMessageEntityMentionName{
				Offset: 0,
				Length: 4,
				UserId: 300,
			}),
		},
	})
	outbox := msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{RandomId: 99, Message: message})
	normalized, err := normalizeOutboxMessage(normalizeOutboxInput{
		SenderUserID: 100,
		PeerType:     payload.PeerTypeUser,
		PeerID:       200,
		Outbox:       outbox,
	})
	if err != nil {
		t.Fatalf("normalizeOutboxMessage() error = %v", err)
	}
	if len(normalized.Entities) != 1 {
		t.Fatalf("entities len = %d, want 1", len(normalized.Entities))
	}
	if normalized.Entities[0].Kind != "mention_name" || normalized.Entities[0].UserID != 300 {
		t.Fatalf("entity = %+v, want mention_name user 300", normalized.Entities[0])
	}
	roundTrip := sentMessageEntities(normalized.Entities)
	mentionName, ok := roundTrip[0].(*tg.TLMessageEntityMentionName)
	if !ok {
		t.Fatalf("roundTrip[0] = %T, want *tg.TLMessageEntityMentionName", roundTrip[0])
	}
	if mentionName.UserId != 300 || mentionName.Offset != 0 || mentionName.Length != 4 {
		t.Fatalf("mentionName = %#v, want user 300 range 0:4", mentionName)
	}
}

func TestMarshalSendRequestV3IncludesMediaAttrsForward(t *testing.T) {
	msg1 := normalizedOutboxMessage{
		SchemaVersion: NormalizedOutboxSchemaVersionV1,
		RandomID:      11,
		FromUserID:    100,
		PeerType:      payload.PeerTypeUser,
		PeerID:        200,
		MessageText:   "caption",
		MediaRef:      &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 300},
		Attrs:         payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 400},
		ForwardRef:    &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 500, Date: 1700000000},
	}
	msg2 := msg1
	msg2.MediaRef = &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 301}
	_, hash1, err := marshalSendRequestV3(msg1, batchSideEffects{})
	if err != nil {
		t.Fatalf("marshalSendRequestV3(msg1) error = %v", err)
	}
	body1, _, err := marshalSendRequestV3(msg1, batchSideEffects{})
	if err != nil {
		t.Fatalf("marshalSendRequestV3(body1) error = %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(body1, &raw); err != nil {
		t.Fatalf("decode request v3: %v", err)
	}
	if _, ok := raw["grouped_id"]; ok {
		t.Fatalf("request payload has top-level grouped_id: %s", string(body1))
	}
	attrs, ok := raw["attrs"].(map[string]any)
	if !ok || attrs["grouped_id"] != float64(400) {
		t.Fatalf("request payload attrs grouped_id missing: %s", string(body1))
	}
	_, hash2, err := marshalSendRequestV3(msg2, batchSideEffects{})
	if err != nil {
		t.Fatalf("marshalSendRequestV3(msg2) error = %v", err)
	}
	if bytes.Equal(hash1, hash2) {
		t.Fatal("request hash did not change when media ref changed")
	}

	msg3 := msg1
	msg3.Attrs.Silent = true
	_, hash3, err := marshalSendRequestV3(msg3, batchSideEffects{})
	if err != nil {
		t.Fatalf("marshalSendRequestV3(msg3) error = %v", err)
	}
	if bytes.Equal(hash1, hash3) {
		t.Fatal("request hash did not change when attrs changed")
	}

	msg4 := msg1
	msg4.ForwardRef = &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 501, Date: 1700000000}
	_, hash4, err := marshalSendRequestV3(msg4, batchSideEffects{})
	if err != nil {
		t.Fatalf("marshalSendRequestV3(msg4) error = %v", err)
	}
	if bytes.Equal(hash1, hash4) {
		t.Fatal("request hash did not change when forward ref changed")
	}

	_, sideEffectHash, err := marshalSendRequestV3(msg1, batchSideEffects{ClearDraft: true, SourcePermAuthKeyID: 9001, ClearDraftBeforeDate: 1700000002})
	if err != nil {
		t.Fatalf("marshalSendRequestV3(side effects) error = %v", err)
	}
	if bytes.Equal(hash1, sideEffectHash) {
		t.Fatal("request hash did not change when batch side effects changed")
	}
}

func TestMsgSendMessageSingleUserPublishesReceiverOperation(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":11,"pts_count":1,"event_type":"new_message","user_message_id":42}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 2001,
			PeerSeq:            5,
			MessageDate:        1_772_000_000,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(2001, 1001),
			Status:              1,
			Pts:                 11,
			PtsCount:            1,
			CurrentPts:          11,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgSendMessage(sendMessageRequest(1001, 1002, 9001, "hello"))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	short, ok := got.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if short.Id != 42 || short.Pts != 11 || short.PtsCount != 1 || short.Date != 1_772_000_000 {
		t.Fatalf("unexpected short sent message: %+v", short)
	}
	if repo.markCanonicalCalls != 1 || repo.markSenderCalls != 1 || repo.markReceiverAckedCalls != 1 || repo.markCompletedCalls != 1 {
		t.Fatalf("unexpected repo call counts: %+v", repo)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != payload.ReceiverOperationID(2001, 1002) {
		t.Fatalf("unexpected receiver operation: %+v", publisher.published)
	}
	receiverOp := mustDecodeMessageOperationV4(t, publisher.published.Payload)
	if receiverOp.MessageFact.SenderUserID != 1001 || receiverOp.MessageFact.ToUserID != 1002 || receiverOp.MessageFact.PeerID != 1001 {
		t.Fatalf("unexpected receiver payload: %+v", receiverOp)
	}
	if updatesClient.processed == nil || updatesClient.processed.UserId != 1001 {
		t.Fatalf("sender operation was not sent to userupdates: %+v", updatesClient.processed)
	}
	if updatesClient.processed.AuthKeyIdExclude == nil || *updatesClient.processed.AuthKeyIdExclude != 9001 {
		t.Fatalf("sender operation auth_key_id_exclude = %v, want 9001", updatesClient.processed.AuthKeyIdExclude)
	}
	senderOp := mustDecodeMessageOperationV4(t, updatesClient.processed.Payload)
	if senderOp.MessageFact.PeerID != 1002 || senderOp.MessageFact.SenderUserID != 1001 || senderOp.MessageFact.ToUserID != 1002 || senderOp.MessageFact.MessageText != "hello" {
		t.Fatalf("unexpected sender payload: %+v", senderOp)
	}
}

func TestMsgSendMessageBuildsV4PayloadWithClientRandomID(t *testing.T) {
	responsePayload := mustOperationResponseV3Payload(t, payload.SenderOperationID(2001, 1001), 11, 1, 42, 77, replyUpdatesForText(t, 42, 77, 11, 1))
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 2001,
			PeerSeq:            5,
			MessageDate:        1_772_000_000,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: testUserOperationResult(1001, payload.SenderOperationID(2001, 1001), 11, responsePayload),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	if _, err := core.MsgSendMessage(sendMessageRequest(1001, 1002, 9001, "hello")); err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if updatesClient.processed == nil {
		t.Fatal("sender operation was not sent to userupdates")
	}
	var senderOp payload.MessageOperationV4
	if err := json.Unmarshal(updatesClient.processed.Payload, &senderOp); err != nil {
		t.Fatalf("decode sender v4 payload: %v", err)
	}
	if senderOp.SchemaVersion != payload.MessageOperationSchemaVersionV4 || senderOp.OperationKind != payload.OperationKindSendMessage {
		t.Fatalf("sender v4 identity = %+v", senderOp)
	}
	if senderOp.MessageFact.ClientRandomID != 77 || senderOp.MessageFact.SenderUserID != 1001 ||
		senderOp.MessageFact.PeerType != payload.PeerTypeUser || senderOp.MessageFact.PeerID != 1002 ||
		senderOp.MessageFact.MessageText != "hello" {
		t.Fatalf("sender message fact = %+v, want random/user/peer/text preserved", senderOp.MessageFact)
	}
	if !bytes.Equal(updatesClient.processed.PayloadHash, payload.HashBytes(updatesClient.processed.Payload)) {
		t.Fatalf("payload hash = %x, want hash of exact payload bytes", updatesClient.processed.PayloadHash)
	}
}

func TestMsgSendMessagePeerTypeChatUsesAffectedOutboxFanout(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":21,"pts_count":1,"event_type":"new_message","user_message_id":42}`)
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(7001, 1001),
			Status:              1,
			Pts:                 21,
			PtsCount:            1,
			CurrentPts:          21,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: payload.HashBytes(responsePayload),
		}),
	}
	chatClient := &fakeMsgChatClient{
		memberIDs: []int64{1001, 1002, 1003},
	}
	repo := newFakeSendMessageRepositoryForChat(7001, 1)
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})

	got, err := core.MsgSendMessage(&msgpb.TLMsgSendMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeChat,
		PeerId:    55,
		Message: []msgpb.OutboxMessageClazz{msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			RandomId: 12345,
			Message: tg.MakeTLMessage(&tg.TLMessage{
				Out:     true,
				FromId:  tg.MakePeerUser(1001),
				PeerId:  tg.MakePeerChat(55),
				Message: "hello chat",
				Date:    1778584910,
			}),
		})},
	})
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if got == nil {
		t.Fatal("MsgSendMessage() returned nil updates")
	}
	if len(chatClient.actions) != 1 || chatClient.actions[0].Action != chatpb.MessageActionSendText || chatClient.actions[0].ChatId != 55 {
		t.Fatalf("chat action checks = %+v, want send_text for chat 55", chatClient.actions)
	}
	if updatesClient.processWithEffects == nil {
		t.Fatal("UserupdatesProcessUserOperationWithEffects was not called")
	}
	if len(updatesClient.processWithEffects.AffectedEffects) != 2 {
		t.Fatalf("affected effects = %d, want 2 receivers", len(updatesClient.processWithEffects.AffectedEffects))
	}
	if repo.markReceiverAckedCalls != 1 || repo.markCompletedCalls != 1 {
		t.Fatalf("unexpected completion calls: repo=%+v", repo)
	}
	for _, effect := range updatesClient.processWithEffects.AffectedEffects {
		if effect.DeliveryPolicy != int32(DeliveryPolicyDurableAsync) {
			t.Fatalf("effect policy = %d, want durable async", effect.DeliveryPolicy)
		}
		op := effect.Operation
		if op.PeerType != payload.PeerTypeChat || op.PeerId != 55 {
			t.Fatalf("receiver peer = type:%d id:%d, want chat 55", op.PeerType, op.PeerId)
		}
		if op.UserId == 1001 {
			t.Fatalf("sender was included as receiver effect: %+v", op)
		}
		body := mustDecodeMessageOperationV4(t, op.Payload)
		if body.MessageFact.PeerType != payload.PeerTypeChat || body.MessageFact.PeerID != 55 || body.MessageFact.SenderUserID != 1001 || body.MessageFact.ToUserID != op.UserId {
			t.Fatalf("receiver payload = %+v, want chat incoming from sender", body)
		}
	}
}

func TestMsgSendMessagePropagatesAttachFactsToChatReceivers(t *testing.T) {
	responsePayload := mustOperationResponseV3Payload(t, payload.SenderOperationID(7101, 1001), 22, 1, 43, 12346, replyUpdatesForChatService(t, 43, 12346, 22, 1))
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: testUserOperationResult(1001, payload.SenderOperationID(7101, 1001), 22, responsePayload),
	}
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002, 1003}}
	repo := newFakeSendMessageRepositoryForChat(7101, 4)
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})
	attachFact := chatParticipantsChangedTLFact(t, 55, 1001, []int64{1001, 1002, 1003})
	req := chatCreateServiceSendRequest(1001, 55, 9001, 12346)
	req.AttachFacts = []msgpb.UpdateFactClazz{attachFact}

	if _, err := core.MsgSendMessage(req); err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if updatesClient.processWithEffects == nil {
		t.Fatal("UserupdatesProcessUserOperationWithEffects was not called")
	}
	var senderOp payload.MessageOperationV4
	if err := json.Unmarshal(updatesClient.processWithEffects.Operation.Payload, &senderOp); err != nil {
		t.Fatalf("decode sender payload: %v", err)
	}
	if len(senderOp.AttachFacts) != 1 || senderOp.AttachFacts[0].Kind != payload.FactKindChatParticipantsChanged {
		t.Fatalf("sender attach facts = %+v, want one chat participants fact", senderOp.AttachFacts)
	}
	if len(updatesClient.processWithEffects.AffectedEffects) != 2 {
		t.Fatalf("affected effects = %d, want 2 receivers", len(updatesClient.processWithEffects.AffectedEffects))
	}
	for _, effect := range updatesClient.processWithEffects.AffectedEffects {
		var receiverOp payload.MessageOperationV4
		if err := json.Unmarshal(effect.Operation.Payload, &receiverOp); err != nil {
			t.Fatalf("decode receiver payload: %v", err)
		}
		if len(receiverOp.AttachFacts) != 1 || receiverOp.AttachFacts[0].Kind != payload.FactKindChatParticipantsChanged {
			t.Fatalf("receiver attach facts = %+v, want one chat participants fact", receiverOp.AttachFacts)
		}
		if receiverOp.MessageFact.PeerID != 55 || receiverOp.MessageFact.SenderUserID != 1001 || receiverOp.MessageFact.ClientRandomID != 12346 {
			t.Fatalf("receiver message fact = %+v, want chat peer/sender/random", receiverOp.MessageFact)
		}
	}
}

func TestMsgSendMessageReplyUsesUserupdatesEnvelope(t *testing.T) {
	responsePayload := mustOperationResponseV3Payload(t, payload.SenderOperationID(7201, 1001), 33, 1, 44, 12347, replyUpdatesForChatService(t, 44, 12347, 33, 1))
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: testUserOperationResult(1001, payload.SenderOperationID(7201, 1001), 33, responsePayload),
	}
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002}}
	repo := newFakeSendMessageRepositoryForChat(7201, 5)
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})

	got, err := core.MsgSendMessage(chatCreateServiceSendRequest(1001, 55, 9001, 12347))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	updates, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("updates = %T, want TLUpdates envelope from userupdates", got.Clazz)
	}
	if len(updates.Users) == 0 || len(updates.Chats) == 0 {
		t.Fatalf("reply envelope dependencies users=%d chats=%d, want non-empty", len(updates.Users), len(updates.Chats))
	}
	if len(updates.Updates) < 2 {
		t.Fatalf("reply updates len = %d, want updateMessageID and updateNewMessage", len(updates.Updates))
	}
	if _, ok := updates.Updates[0].(*tg.TLUpdateMessageID); !ok {
		t.Fatalf("updates[0] = %T, want updateMessageID from userupdates envelope", updates.Updates[0])
	}
	if _, ok := updates.Updates[1].(*tg.TLUpdateNewMessage); !ok {
		t.Fatalf("updates[1] = %T, want updateNewMessage from userupdates envelope", updates.Updates[1])
	}
	if updatesClient.getOperationResultCalls != 0 {
		t.Fatalf("get operation result calls = %d, want no extra userupdates RPC", updatesClient.getOperationResultCalls)
	}
}

func TestMsgSendMessageReplyRejectsLegacyOrEmptyUserupdatesEnvelope(t *testing.T) {
	tests := []struct {
		name string
		body []byte
	}{
		{
			name: "v2 response",
			body: []byte(`{"schema_version":2,"pts":11,"pts_count":1,"event_type":"new_message","user_message_id":42}`),
		},
		{
			name: "empty v3 reply envelope",
			body: mustMarshalOperationResponseV3(t, payload.OperationResponseV3{
				SchemaVersion:       payload.OperationResponseSchemaVersionV3,
				OperationID:         payload.SenderOperationID(2001, 1001),
				Pts:                 11,
				PtsCount:            1,
				EventType:           payload.EventKindNewMessage,
				UserMessageID:       42,
				ClientRandomID:      77,
				ReplyEnvelopeCodec:  payload.ReplyEnvelopeCodecTLBinary,
				ReplyEnvelopeSchema: payload.ReplyEnvelopeSchemaV1,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sentMessageUpdatesFromUserupdatesEnvelope(testUserOperationResult(1001, payload.SenderOperationID(2001, 1001), 11, tt.body), "sender")
			if !errors.Is(err, msgpb.ErrSenderSyncFailed) {
				t.Fatalf("sentMessageUpdatesFromUserupdatesEnvelope() error = %v, want ErrSenderSyncFailed", err)
			}
		})
	}
}

func TestMsgSendMessagePeerTypeChatCreateServiceActionFansOut(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":22,"pts_count":1,"event_type":"new_message","user_message_id":43}`)
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: testUserOperationResult(1001, payload.SenderOperationID(7004, 1001), 22, responsePayload),
	}
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002, 1003}}
	repo := newFakeSendMessageRepositoryForChat(7004, 4)
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})

	got, err := core.MsgSendMessage(&msgpb.TLMsgSendMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeChat,
		PeerId:    55,
		Message: []msgpb.OutboxMessageClazz{msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			RandomId: 12346,
			Message: tg.MakeTLMessageService(&tg.TLMessageService{
				Out:    true,
				FromId: tg.MakePeerUser(1001),
				PeerId: tg.MakePeerChat(55),
				Date:   1778648899,
				Action: tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
					Title: "new chat",
					Users: []int64{1002, 1003},
				}),
			}),
		})},
	})
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	updates, ok := got.ToUpdates()
	if !ok || len(updates.Updates) != 2 {
		t.Fatalf("updates = %T %+v, want updateMessageID + updateNewMessage", got, got)
	}
	newMessage, ok := updates.Updates[1].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("updates[1] = %T, want updateNewMessage", updates.Updates[1])
	}
	service, ok := newMessage.Message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("sent message = %T, want messageService", newMessage.Message)
	}
	action, ok := service.Action.(*tg.TLMessageActionChatCreate)
	if !ok || action.Title != "new chat" || len(action.Users) != 2 {
		t.Fatalf("service action = %T %+v, want chat create title/users", service.Action, service.Action)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 2 {
		t.Fatalf("affected effects = %+v, want 2 receivers", updatesClient.processWithEffects)
	}
	senderOp := mustDecodeMessageOperationV4(t, updatesClient.processWithEffects.Operation.Payload)
	if senderOp.MessageFact.ServiceAction == nil || senderOp.MessageFact.ServiceAction.Kind != payload.ServiceActionKindChatCreate {
		t.Fatalf("sender service action = %+v, want chat create", senderOp.MessageFact.ServiceAction)
	}
	for _, effect := range updatesClient.processWithEffects.AffectedEffects {
		receiverOp := mustDecodeMessageOperationV4(t, effect.Operation.Payload)
		if receiverOp.MessageFact.ServiceAction == nil || receiverOp.MessageFact.ServiceAction.Kind != payload.ServiceActionKindChatCreate || receiverOp.MessageFact.PeerID != 55 {
			t.Fatalf("receiver payload = %+v, want chat create in peer 55", receiverOp)
		}
	}
}

func TestMsgSendMessagePeerTypeChatMediaReturnsFullUpdatesAndFansOut(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":31,"pts_count":1,"event_type":"new_message","user_message_id":61}`)
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: testUserOperationResult(1001, payload.SenderOperationID(7002, 1001), 31, responsePayload),
	}
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002}}
	repo := newFakeSendMessageRepositoryForChat(7002, 2)
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})
	req := buildMediaSendRequestForTest(1001, 55, 9001, 11, "caption")
	req.PeerType = payload.PeerTypeChat
	req.PeerId = 55
	req.Message[0].Message.(*tg.TLMessage).PeerId = tg.MakePeerChat(55)

	got, err := core.MsgSendMessage(req)
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if len(chatClient.actions) != 1 || chatClient.actions[0].Action != chatpb.MessageActionSendMediaPhoto || chatClient.actions[0].MediaKind != "photo" {
		t.Fatalf("chat action checks = %+v, want send_media_photo/photo", chatClient.actions)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 1 {
		t.Fatalf("affected effects = %+v, want one chat receiver", updatesClient.processWithEffects)
	}
	message := firstSentUpdateMessage(t, got)
	if peer, ok := message.PeerId.(*tg.TLPeerChat); !ok || peer.ChatId != 55 {
		t.Fatalf("sent peer = %#v, want peerChat 55", message.PeerId)
	}
	if _, ok := message.Media.(*tg.TLMessageMediaPhoto); !ok {
		t.Fatalf("sent media = %T, want photo", message.Media)
	}
}

func TestMsgSendMessagePeerTypeChatForwardUsesForwardAction(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":32,"pts_count":1,"event_type":"new_message","user_message_id":62}`)
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: testUserOperationResult(1001, payload.SenderOperationID(7003, 1001), 32, responsePayload),
	}
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002}}
	repo := newFakeSendMessageRepositoryForChat(7003, 3)
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})
	req := buildForwardSendRequestForTest(1001, 55, 9001, 12, 7)
	req.PeerType = payload.PeerTypeChat
	req.PeerId = 55
	req.Message[0].Message.(*tg.TLMessage).PeerId = tg.MakePeerChat(55)

	got, err := core.MsgSendMessage(req)
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if len(chatClient.actions) != 1 || chatClient.actions[0].Action != chatpb.MessageActionForwardToChat {
		t.Fatalf("chat action checks = %+v, want forward_to_chat", chatClient.actions)
	}
	message := firstSentUpdateMessage(t, got)
	if peer, ok := message.PeerId.(*tg.TLPeerChat); !ok || peer.ChatId != 55 {
		t.Fatalf("sent peer = %#v, want peerChat 55", message.PeerId)
	}
	if message.FwdFrom == nil {
		t.Fatal("sent message missing forward header")
	}
}

func TestMsgSendMessagePeerTypeChatAlbumUsesBatchCanonicalAndAffectedEffects(t *testing.T) {
	groupedID := int64(9901)
	responsePayload1 := []byte(`{"schema_version":2,"pts":41,"pts_count":1,"event_type":"new_message","user_message_id":71}`)
	responsePayload2 := []byte(`{"schema_version":2,"pts":42,"pts_count":1,"event_type":"new_message","user_message_id":72}`)
	repo := &fakeMsgRepository{
		forwardVisible: true,
		batchResult: &repository.CanonicalBatchResult{Items: []repository.CanonicalMessageResult{
			{SendStateID: 601, CanonicalMessageID: 8001, PeerSeq: 1, MessageDate: 1_772_000_000, RequestPayloadHash: payload.HashBytes([]byte("h1")), CreatedNew: true},
			{SendStateID: 602, CanonicalMessageID: 8002, PeerSeq: 2, MessageDate: 1_772_000_001, RequestPayloadHash: payload.HashBytes([]byte("h2")), CreatedNew: true},
		}},
	}
	updatesClient := &fakeUserUpdatesClient{processResults: []*userupdates.UserOperationResult{
		testUserOperationResult(1001, payload.SenderOperationID(8001, 1001), 41, responsePayload1),
		testUserOperationResult(1001, payload.SenderOperationID(8002, 1001), 42, responsePayload2),
	}}
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002}}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})
	req := buildBatchSendRequestForTest(1001, 55, 9001, []int64{11, 12})
	req.PeerType = payload.PeerTypeChat
	req.PeerId = 55
	for _, outbox := range req.Message {
		message := outbox.Message.(*tg.TLMessage)
		message.PeerId = tg.MakePeerChat(55)
		message.GroupedId = &groupedID
		message.Media = tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{Photo: tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: 333})})
	}

	got, err := core.MsgSendMessage(req)
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if repo.batchCreateCalls != 1 {
		t.Fatalf("batchCreateCalls = %d, want 1", repo.batchCreateCalls)
	}
	if len(chatClient.actions) != 2 ||
		chatClient.actions[0].Action != chatpb.MessageActionSendAlbum || chatClient.actions[0].MediaKind != "album" ||
		chatClient.actions[1].Action != chatpb.MessageActionSendMediaPhoto || chatClient.actions[1].MediaKind != "photo" {
		t.Fatalf("chat action checks = %+v, want send_album then send_media_photo", chatClient.actions)
	}
	if updatesClient.batchApplyCalls != 0 {
		t.Fatalf("batchApplyCalls = %d, want per-sender with-effects dispatch", updatesClient.batchApplyCalls)
	}
	if len(updatesClient.processWithEffectsList) != 2 || len(updatesClient.processedOperations) != 2 {
		t.Fatalf("with-effects calls=%d processed=%d, want 2/2", len(updatesClient.processWithEffectsList), len(updatesClient.processedOperations))
	}
	for i, call := range updatesClient.processWithEffectsList {
		if len(call.AffectedEffects) != 1 {
			t.Fatalf("call %d affected effects = %d, want one receiver", i, len(call.AffectedEffects))
		}
	}
	if repo.markSenderCalls != 2 || repo.markReceiverAckedCalls != 2 || repo.markCompletedCalls != 2 {
		t.Fatalf("unexpected repo calls: %+v", repo)
	}
	message := firstSentUpdateMessage(t, got)
	if peer, ok := message.PeerId.(*tg.TLPeerChat); !ok || peer.ChatId != 55 {
		t.Fatalf("sent peer = %#v, want peerChat 55", message.PeerId)
	}
}

func TestMsgSendMessageBatchReturnsFullUpdates(t *testing.T) {
	core, fakeRepo, fakeUpdates := newBatchMsgCoreForTest(t)
	in := buildBatchSendRequestForTest(100, 200, 9001, []int64{11, 12})

	got, err := core.MsgSendMessage(in)
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if fakeRepo.batchCreateCalls != 1 {
		t.Fatalf("batchCreateCalls = %d, want 1", fakeRepo.batchCreateCalls)
	}
	if fakeRepo.markSenderCalls != 2 {
		t.Fatalf("markSenderCalls = %d, want 2", fakeRepo.markSenderCalls)
	}
	if fakeUpdates.batchApplyCalls != 1 {
		t.Fatalf("batchApplyCalls = %d, want 1", fakeUpdates.batchApplyCalls)
	}
	if fakeUpdates.batchApplyLen != 2 {
		t.Fatalf("batchApplyLen = %d, want 2", fakeUpdates.batchApplyLen)
	}
	if _, ok := got.Clazz.(*tg.TLUpdates); !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
}

func TestMsgSendMessageBatchCombinesUserupdatesReplyEnvelopesInItemOrder(t *testing.T) {
	core, _, fakeUpdates := newBatchMsgCoreForTest(t)
	fakeUpdates.processResults = []*userupdates.UserOperationResult{
		testUserOperationResult(100, payload.SenderOperationID(9001, 100), 31,
			mustOperationResponseV3Payload(t, payload.SenderOperationID(9001, 100), 31, 1, 101, 11,
				replyUpdatesForBatchEnvelope(t, 101, 910011, 31, 1, "from envelope first", 3001, 901))),
		testUserOperationResult(100, payload.SenderOperationID(9002, 100), 32,
			mustOperationResponseV3Payload(t, payload.SenderOperationID(9002, 100), 32, 1, 102, 12,
				replyUpdatesForBatchEnvelope(t, 102, 910012, 32, 1, "from envelope second", 3002, 902))),
	}

	got, err := core.MsgSendMessage(buildBatchSendRequestForTest(100, 200, 9001, []int64{11, 12}))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	updates, ok := got.Clazz.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
	if len(updates.Updates) != 4 {
		t.Fatalf("updates len = %d, want 4", len(updates.Updates))
	}
	firstID, ok := updates.Updates[0].(*tg.TLUpdateMessageID)
	if !ok || firstID.RandomId != 910011 {
		t.Fatalf("first update = %#v, want envelope updateMessageID random 910011", updates.Updates[0])
	}
	firstNew, ok := updates.Updates[1].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("second update = %T, want updateNewMessage", updates.Updates[1])
	}
	firstMessage, ok := firstNew.Message.(*tg.TLMessage)
	if !ok || firstMessage.Message != "from envelope first" {
		t.Fatalf("first message = %#v, want envelope message text", firstNew.Message)
	}
	secondID, ok := updates.Updates[2].(*tg.TLUpdateMessageID)
	if !ok || secondID.RandomId != 910012 {
		t.Fatalf("third update = %#v, want envelope updateMessageID random 910012", updates.Updates[2])
	}
	secondNew, ok := updates.Updates[3].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("fourth update = %T, want updateNewMessage", updates.Updates[3])
	}
	secondMessage, ok := secondNew.Message.(*tg.TLMessage)
	if !ok || secondMessage.Message != "from envelope second" {
		t.Fatalf("second message = %#v, want envelope message text", secondNew.Message)
	}
	if len(updates.Users) != 2 {
		t.Fatalf("users len = %d, want 2", len(updates.Users))
	}
	firstUser, firstUserOK := updates.Users[0].(*tg.TLUserEmpty)
	secondUser, secondUserOK := updates.Users[1].(*tg.TLUserEmpty)
	if !firstUserOK || firstUser.Id != 3001 || !secondUserOK || secondUser.Id != 3002 {
		t.Fatalf("users = %#v, want envelope users in item order", updates.Users)
	}
	if len(updates.Chats) != 2 {
		t.Fatalf("chats len = %d, want 2", len(updates.Chats))
	}
	firstChat, firstChatOK := updates.Chats[0].(*tg.TLChatEmpty)
	secondChat, secondChatOK := updates.Chats[1].(*tg.TLChatEmpty)
	if !firstChatOK || firstChat.Id != 901 || !secondChatOK || secondChat.Id != 902 {
		t.Fatalf("chats = %#v, want envelope chats in item order", updates.Chats)
	}
}

func TestMsgGetUserMessageReadsViewerMessageID(t *testing.T) {
	core, fakeRepo := newMsgReadCoreForTest(t)
	fakeRepo.userMessages = map[int64]*repository.UserMessageBox{
		7: {
			UserID:             100,
			UserMessageID:      7,
			CanonicalMessageID: 900,
			PeerType:           payload.PeerTypeUser,
			PeerID:             200,
			PeerSeq:            3,
			FromUserID:         200,
			MessageText:        "source",
			MessageDate:        1_772_000_300,
		},
	}

	got, err := core.MsgGetUserMessage(&msgpb.TLMsgGetUserMessage{UserId: 100, Id: 7})
	if err != nil {
		t.Fatalf("MsgGetUserMessage() error = %v", err)
	}
	if got == nil || got.Message == nil {
		t.Fatalf("MessageBox = %#v, want message", got)
	}
	if got.MessageId != 7 {
		t.Fatalf("MessageId = %d, want viewer public id 7", got.MessageId)
	}
	message, ok := got.Message.(*tg.TLMessage)
	if !ok || message.Id != 7 || message.Message != "source" {
		t.Fatalf("Message = %+v ok=%v, want public id text message", got.Message, ok)
	}
	if fakeRepo.getUserMessageInput.UserID != 100 || fakeRepo.getUserMessageInput.UserMessageID != 7 {
		t.Fatalf("GetUserMessage input = %+v", fakeRepo.getUserMessageInput)
	}
}

func TestMsgGetUserMessageListRejectsInvalidID(t *testing.T) {
	core, fakeRepo := newMsgReadCoreForTest(t)
	fakeRepo.userMessages = map[int64]*repository.UserMessageBox{
		7: {UserID: 100, UserMessageID: 7, MessageText: "source", MessageDate: 1_772_000_300},
	}

	_, err := core.MsgGetUserMessageList(&msgpb.TLMsgGetUserMessageList{UserId: 100, IdList: []int32{7, 0}})
	if !errors.Is(err, msgpb.ErrMsgIdInvalid) {
		t.Fatalf("MsgGetUserMessageList() error = %v, want ErrMsgIdInvalid", err)
	}
	if len(fakeRepo.getUserMessageListInput.IDs) != 2 || fakeRepo.getUserMessageListInput.IDs[1] != 0 {
		t.Fatalf("GetUserMessageList ids = %+v, want invalid id preserved for repository validation", fakeRepo.getUserMessageListInput.IDs)
	}
}

func TestMsgGetUserMessageListRejectsMissingID(t *testing.T) {
	core, fakeRepo := newMsgReadCoreForTest(t)
	fakeRepo.userMessages = map[int64]*repository.UserMessageBox{
		7: {UserID: 100, UserMessageID: 7, MessageText: "source", MessageDate: 1_772_000_300},
	}

	_, err := core.MsgGetUserMessageList(&msgpb.TLMsgGetUserMessageList{UserId: 100, IdList: []int32{7, 99}})
	if !errors.Is(err, msgpb.ErrMsgIdInvalid) {
		t.Fatalf("MsgGetUserMessageList() error = %v, want ErrMsgIdInvalid", err)
	}
}

func TestMsgGetUserMessagePreservesV3ViewPayloadShape(t *testing.T) {
	core, fakeRepo := newMsgReadCoreForTest(t)
	eventPayload := mustMarshalMsgMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:        payload.MessageEventSchemaVersionV3,
		EventKind:            payload.EventKindNewMessage,
		CanonicalMessageID:   900,
		PeerSeq:              3,
		MessageID:            7,
		PeerType:             payload.PeerTypeUser,
		PeerID:               200,
		FromUserID:           200,
		ToUserID:             100,
		Date:                 1_772_000_300,
		Out:                  false,
		MessageText:          "caption",
		ReplyToUserMessageID: 5,
		MediaRef:             &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "document", ID: 333},
		Attrs:                &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444, Noforwards: true},
		ForwardRef:           &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 300, Date: 1_772_000_001},
	})
	fakeRepo.userMessages = map[int64]*repository.UserMessageBox{
		7: {
			UserID:             100,
			UserMessageID:      7,
			CanonicalMessageID: 900,
			PeerType:           payload.PeerTypeUser,
			PeerID:             200,
			PeerSeq:            3,
			FromUserID:         200,
			MessageText:        "plain fallback must not win",
			MessageDate:        1_772_000_300,
			ViewPayload:        eventPayload,
		},
	}

	got, err := core.MsgGetUserMessage(&msgpb.TLMsgGetUserMessage{UserId: 100, Id: 7})
	if err != nil {
		t.Fatalf("MsgGetUserMessage() error = %v", err)
	}
	message, ok := got.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("Message = %T, want *tg.TLMessage", got.Message)
	}
	if message.Message != "caption" || message.Noforwards != true {
		t.Fatalf("message text/noforwards = %q/%t, want persisted V3 shape", message.Message, message.Noforwards)
	}
	if _, ok := message.Media.(*tg.TLMessageMediaDocument); !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", message.Media)
	}
	if message.GroupedId == nil || *message.GroupedId != 444 || message.FwdFrom == nil {
		t.Fatalf("grouped/forward = grouped:%v fwd:%T", message.GroupedId, message.FwdFrom)
	}
	reply, ok := message.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok || reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 5 {
		t.Fatalf("reply = %#v, want reply_to_msg_id 5", message.ReplyTo)
	}
}

func TestSentMessageDocumentMediaProjectsFullUploadedDocumentContract(t *testing.T) {
	videoStartTs := 1.25
	videoTimestamp := int32(7)
	media := sentMessageMedia(&payload.MediaRefV1{
		SchemaVersion: payload.MediaRefSchemaVersionV2,
		Kind:          "document",
		ID:            555,
		AccessHash:    666,
		FileReference: []byte("doc-ref"),
		Date:          1_700_000_000,
		DcID:          4,
		MimeType:      "video/mp4",
		Size:          98765,
		DocumentThumbs: []payload.PhotoSizeRefV1{
			{Kind: "size", Type: "m", W: 320, H: 200, Size: 1234},
		},
		DocumentVideoThumbs: []payload.VideoSizeRefV1{
			{Kind: "size", Type: "v", W: 320, H: 200, Size: 4567, VideoStartTs: &videoStartTs},
		},
		DocumentAttributes: []payload.DocumentAttributeRefV1{
			{Kind: "filename", FileName: "clip.mp4"},
			{Kind: "video", W: 1280, H: 720, DurationFloat: 3.5, SupportsStreaming: true, VideoStartTs: &videoStartTs},
			{Kind: "sticker", Alt: ":)", StickerSet: &payload.StickerSetRefV1{Kind: "id", ID: 1001, AccessHash: 2002}, Mask: true, MaskCoords: &payload.MaskCoordsRefV1{N: 1, X: 0.5, Y: 0.25, Zoom: 1.5}},
			{Kind: "custom_emoji", Alt: ":)", StickerSet: &payload.StickerSetRefV1{Kind: "id", ID: 3003, AccessHash: 4004}, Free: true, TextColor: true},
			{Kind: "has_stickers"},
		},
		DocumentMediaFlags: &payload.DocumentMediaFlagsV1{Video: true, Spoiler: true},
		VideoCover: &payload.PhotoRefV1{
			ID:            777,
			AccessHash:    888,
			FileReference: []byte("cover-ref"),
			Date:          1_700_000_001,
			DcID:          5,
			Sizes: []payload.PhotoSizeRefV1{
				{Kind: "size", Type: "x", W: 640, H: 360, Size: 4321},
			},
		},
		VideoTimestamp: &videoTimestamp,
	})

	documentMedia, ok := media.(*tg.TLMessageMediaDocument)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", media)
	}
	if !documentMedia.Video || !documentMedia.Spoiler {
		t.Fatalf("messageMediaDocument flags video=%v spoiler=%v, want true/true", documentMedia.Video, documentMedia.Spoiler)
	}
	document, ok := documentMedia.Document.(*tg.TLDocument)
	if !ok {
		t.Fatalf("document = %T, want *tg.TLDocument", documentMedia.Document)
	}
	assertSentDocumentIdentity(t, document)
	if len(document.VideoThumbs) != 1 {
		t.Fatalf("VideoThumbs len = %d, want 1", len(document.VideoThumbs))
	}
	videoThumb, ok := document.VideoThumbs[0].(*tg.TLVideoSize)
	if !ok {
		t.Fatalf("VideoThumbs[0] = %T, want *tg.TLVideoSize", document.VideoThumbs[0])
	}
	if videoThumb.Type != "v" || videoThumb.W != 320 || videoThumb.H != 200 || videoThumb.Size2 != 4567 {
		t.Fatalf("VideoThumbs[0] = %#v, want videoSize v 320x200/4567", videoThumb)
	}
	if videoThumb.VideoStartTs == nil || *videoThumb.VideoStartTs != videoStartTs {
		t.Fatalf("VideoThumbs[0].VideoStartTs = %v, want %v", videoThumb.VideoStartTs, videoStartTs)
	}
	assertSentDocumentAttributes(t, document.Attributes, videoStartTs)
	if documentMedia.VideoTimestamp == nil || *documentMedia.VideoTimestamp != videoTimestamp {
		t.Fatalf("VideoTimestamp = %v, want %d", documentMedia.VideoTimestamp, videoTimestamp)
	}
	videoCover, ok := documentMedia.VideoCover.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("VideoCover = %T, want *tg.TLPhoto", documentMedia.VideoCover)
	}
	assertSentVideoCover(t, videoCover)
}

func TestSentMessageDocumentMediaInfersLegacyFlagsFromAttributes(t *testing.T) {
	media := sentMessageMedia(&payload.MediaRefV1{
		Kind:     "document",
		MimeType: "video/mp4",
		DocumentAttributes: []payload.DocumentAttributeRefV1{
			{Kind: "video", RoundMessage: true},
			{Kind: "audio", Voice: true},
		},
	})
	documentMedia, ok := media.(*tg.TLMessageMediaDocument)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", media)
	}
	if !documentMedia.Video || !documentMedia.Round || !documentMedia.Voice {
		t.Fatalf("document media flags = video:%v round:%v voice:%v, want all inferred", documentMedia.Video, documentMedia.Round, documentMedia.Voice)
	}
}

func TestSentMessageDocumentMediaDoesNotInferVideoForWebMStickerOrCustomEmoji(t *testing.T) {
	for _, tt := range []struct {
		name string
		attr payload.DocumentAttributeRefV1
	}{
		{name: "sticker", attr: payload.DocumentAttributeRefV1{Kind: "sticker"}},
		{name: "custom_emoji", attr: payload.DocumentAttributeRefV1{Kind: "custom_emoji"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			media := sentMessageMedia(&payload.MediaRefV1{
				Kind:     "document",
				MimeType: "video/webm",
				DocumentAttributes: []payload.DocumentAttributeRefV1{
					{Kind: "video"},
					tt.attr,
				},
			})
			documentMedia, ok := media.(*tg.TLMessageMediaDocument)
			if !ok {
				t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", media)
			}
			if documentMedia.Video {
				t.Fatalf("document media Video = true, want false for video/webm %s", tt.name)
			}
		})
	}
}

func assertSentDocumentIdentity(t *testing.T, document *tg.TLDocument) {
	t.Helper()
	if document.Id != 555 ||
		document.AccessHash != 666 ||
		string(document.FileReference) != "doc-ref" ||
		document.Date != 1_700_000_000 ||
		document.DcId != 4 ||
		document.MimeType != "video/mp4" ||
		document.Size2 != 98765 {
		t.Fatalf("document identity = %#v, want full uploaded document identity", document)
	}
}

func TestForwardRevalidationRejectsInvisibleSource(t *testing.T) {
	core, fakeRepo, _ := newBatchMsgCoreForTest(t)
	fakeRepo.forwardVisible = false

	_, err := core.MsgSendMessage(buildForwardSendRequestForTest(100, 200, 9001, 11, 7))
	if !errors.Is(err, msgpb.ErrMsgIdInvalid) {
		t.Fatalf("error = %v, want ErrMsgIdInvalid", err)
	}
	if fakeRepo.batchCreateCalls != 0 {
		t.Fatal("canonical rows were created before forward revalidation")
	}
}

func assertSentDocumentAttributes(t *testing.T, attrs []tg.DocumentAttributeClazz, videoStartTs float64) {
	t.Helper()
	filename, hasFilename := findSentDocumentAttribute[*tg.TLDocumentAttributeFilename](attrs)
	video, hasVideo := findSentDocumentAttribute[*tg.TLDocumentAttributeVideo](attrs)
	sticker, hasSticker := findSentDocumentAttribute[*tg.TLDocumentAttributeSticker](attrs)
	customEmoji, hasCustomEmoji := findSentDocumentAttribute[*tg.TLDocumentAttributeCustomEmoji](attrs)
	_, hasStickers := findSentDocumentAttribute[*tg.TLDocumentAttributeHasStickers](attrs)
	if !hasFilename || !hasVideo || !hasSticker || !hasCustomEmoji || !hasStickers {
		t.Fatalf("document attrs = %#v, want filename/video/sticker/custom_emoji/has_stickers", attrs)
	}
	if filename.FileName != "clip.mp4" {
		t.Fatalf("filename attr FileName = %q, want clip.mp4", filename.FileName)
	}
	if video.Duration != 3.5 || video.W != 1280 || video.H != 720 || !video.SupportsStreaming {
		t.Fatalf("video attr = %#v, want duration/w/h/supports_streaming preserved", video)
	}
	if video.VideoStartTs == nil || *video.VideoStartTs != videoStartTs {
		t.Fatalf("video attr VideoStartTs = %v, want %v", video.VideoStartTs, videoStartTs)
	}
	stickerSet, ok := sticker.Stickerset.(*tg.TLInputStickerSetID)
	if !ok || stickerSet.Id != 1001 || stickerSet.AccessHash != 2002 {
		t.Fatalf("sticker stickerset = %#v, want inputStickerSetID 1001/2002", sticker.Stickerset)
	}
	maskCoords := sticker.MaskCoords
	if maskCoords == nil || maskCoords.N != 1 || maskCoords.X != 0.5 || maskCoords.Y != 0.25 || maskCoords.Zoom != 1.5 {
		t.Fatalf("sticker mask coords = %#v, want exact TLMaskCoords", sticker.MaskCoords)
	}
	if sticker.Alt != ":)" || !sticker.Mask {
		t.Fatalf("sticker attr = %#v, want alt and mask preserved", sticker)
	}
	customStickerSet, ok := customEmoji.Stickerset.(*tg.TLInputStickerSetID)
	if !ok || customStickerSet.Id != 3003 || customStickerSet.AccessHash != 4004 {
		t.Fatalf("custom emoji stickerset = %#v, want inputStickerSetID 3003/4004", customEmoji.Stickerset)
	}
	if customEmoji.Alt != ":)" || !customEmoji.Free || !customEmoji.TextColor {
		t.Fatalf("custom emoji attr = %#v, want alt/free/text_color preserved", customEmoji)
	}
}

func assertSentVideoCover(t *testing.T, cover *tg.TLPhoto) {
	t.Helper()
	if cover.Id != 777 || cover.AccessHash != 888 || string(cover.FileReference) != "cover-ref" || cover.Date != 1_700_000_001 || cover.DcId != 5 {
		t.Fatalf("VideoCover = %#v, want full photo 777", cover)
	}
	if len(cover.Sizes) != 1 {
		t.Fatalf("VideoCover.Sizes len = %d, want 1", len(cover.Sizes))
	}
	size, ok := cover.Sizes[0].(*tg.TLPhotoSize)
	if !ok {
		t.Fatalf("VideoCover.Sizes[0] = %T, want *tg.TLPhotoSize", cover.Sizes[0])
	}
	if size.Type != "x" || size.W != 640 || size.H != 360 || size.Size2 != 4321 {
		t.Fatalf("VideoCover.Sizes[0] = %#v, want photoSize x 640x360/4321", size)
	}
}

func findSentDocumentAttribute[T tg.DocumentAttributeClazz](attrs []tg.DocumentAttributeClazz) (T, bool) {
	for _, attr := range attrs {
		if got, ok := attr.(T); ok {
			return got, true
		}
	}
	var zero T
	return zero, false
}

func TestMsgSendMessageBatchReceiverAckFailureIsRetryable(t *testing.T) {
	core, fakeRepo, fakeUpdates := newBatchMsgCoreForTest(t)
	fakeUpdates.receiverErr = errors.New("broker unavailable")

	_, err := core.MsgSendMessage(buildBatchSendRequestForTest(100, 200, 9001, []int64{11, 12}))
	if !errors.Is(err, msgpb.ErrSenderSyncFailed) {
		t.Fatalf("error = %v, want retryable msg error", err)
	}
	if fakeRepo.markRetryableCalls != 2 {
		t.Fatalf("MarkRetryableFailure calls = %d, want 2", fakeRepo.markRetryableCalls)
	}
}

func TestMsgSendMessageBatchRetrySkipsCommittedSenderItems(t *testing.T) {
	core, fakeRepo, fakeUpdates := newBatchMsgCoreForTest(t)
	committedPayload := mustOperationResponseV3Payload(t, payload.SenderOperationID(9001, 100), 21, 1, 51, 11, replyUpdatesForText(t, 51, 11, 21, 1))
	fakeRepo.batchResult.Items[0].SendStateStatus = repository.SendStateStatusSenderCommitted
	fakeRepo.batchResult.Items[0].SenderOperationID = payload.SenderOperationID(fakeRepo.batchResult.Items[0].CanonicalMessageID, 100)
	fakeRepo.batchResult.Items[0].SenderPTS = 21
	fakeRepo.batchResult.Items[0].SenderPTSCount = 1
	fakeRepo.batchResult.Items[0].SenderUpdatePayload = committedPayload
	fakeRepo.batchResult.Items[0].SenderUpdatePayloadHash = payload.HashBytes(committedPayload)
	fakeRepo.batchResult.Items[1].SendStateStatus = repository.SendStateStatusCanonical

	got, err := core.MsgSendMessage(buildBatchSendRequestForTest(100, 200, 9001, []int64{11, 12}))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if fakeUpdates.batchApplyCalls != 1 {
		t.Fatalf("batchApplyCalls = %d, want 1", fakeUpdates.batchApplyCalls)
	}
	if fakeUpdates.batchApplyLen != 1 {
		t.Fatalf("batchApplyLen = %d, want 1", fakeUpdates.batchApplyLen)
	}
	if len(fakeUpdates.processedOperations) != 1 || fakeUpdates.processedOperations[0].OperationId != payload.SenderOperationID(9002, 100) {
		t.Fatalf("processed operations = %+v, want only second sender operation", fakeUpdates.processedOperations)
	}
	if fakeRepo.markSenderCalls != 1 {
		t.Fatalf("markSenderCalls = %d, want 1", fakeRepo.markSenderCalls)
	}
	if _, ok := got.Clazz.(*tg.TLUpdates); !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
}

func TestMsgSendMessageBatchRetryResumesReceiverAckWithoutSenderDuplicate(t *testing.T) {
	core, fakeRepo, fakeUpdates := newBatchMsgCoreForTest(t)
	for i := range fakeRepo.batchResult.Items {
		item := &fakeRepo.batchResult.Items[i]
		pts := int64(21 + i)
		userMessageID := int64(51 + i)
		randomID := int64(11 + i)
		responsePayload := mustOperationResponseV3Payload(t, payload.SenderOperationID(item.CanonicalMessageID, 100), pts, 1, userMessageID, randomID, replyUpdatesForText(t, userMessageID, randomID, pts, 1))
		item.SendStateStatus = repository.SendStateStatusFailedRetryable
		item.SenderOperationID = payload.SenderOperationID(item.CanonicalMessageID, 100)
		item.SenderPTS = pts
		item.SenderPTSCount = 1
		item.SenderUpdatePayload = responsePayload
		item.SenderUpdatePayloadHash = payload.HashBytes(responsePayload)
	}

	got, err := core.MsgSendMessage(buildBatchSendRequestForTest(100, 200, 9001, []int64{11, 12}))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if fakeUpdates.batchApplyCalls != 0 || fakeRepo.markSenderCalls != 0 {
		t.Fatalf("sender projection duplicated: batchApplyCalls=%d markSenderCalls=%d", fakeUpdates.batchApplyCalls, fakeRepo.markSenderCalls)
	}
	if len(fakeUpdates.receiverPublished) != 1 {
		t.Fatalf("receiver published len = %d, want 1", len(fakeUpdates.receiverPublished))
	}
	if fakeRepo.markReceiverAckedCalls != 2 || fakeRepo.markCompletedCalls != 2 {
		t.Fatalf("unexpected retry completion calls: repo=%+v", fakeRepo)
	}
	if _, ok := got.Clazz.(*tg.TLUpdates); !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
}

func TestMsgSendMessageBatchPublishesOneReceiverBatchOperation(t *testing.T) {
	core, _, fakeUpdates := newBatchMsgCoreForTest(t)

	if _, err := core.MsgSendMessage(buildBatchSendRequestForTest(100, 200, 9001, []int64{11, 12})); err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if len(fakeUpdates.receiverPublished) != 1 {
		t.Fatalf("receiver published len = %d, want 1", len(fakeUpdates.receiverPublished))
	}
	receiver := fakeUpdates.receiverPublished[0]
	if receiver.OperationID != payload.ReceiverBatchOperationID(200, []int64{9001, 9002}) {
		t.Fatalf("receiver operation_id = %q, want batch id", receiver.OperationID)
	}
	batch := mustDecodeMessageOperationBatchV1(t, receiver.Payload)
	if len(batch.Messages) != 2 {
		t.Fatalf("batch messages len = %d, want 2", len(batch.Messages))
	}
	if batch.Messages[0].CanonicalMessageID != 9001 || batch.Messages[0].ClientRandomID != 11 || batch.Messages[0].MessageText != "a" {
		t.Fatalf("first batch message = %+v, want canonical 9001 random 11 text a", batch.Messages[0])
	}
	if batch.Messages[1].CanonicalMessageID != 9002 || batch.Messages[1].ClientRandomID != 12 || batch.Messages[1].MessageText != "b" {
		t.Fatalf("second batch message = %+v, want canonical 9002 random 12 text b", batch.Messages[1])
	}
}

func TestMsgSendMessageBatchReceiverPayloadIncludesRichFields(t *testing.T) {
	core, fakeRepo, fakeUpdates := newBatchMsgCoreForTest(t)
	fakeRepo.resolvedMessageID = &repository.ResolvedMessageID{
		UserID:             100,
		PeerType:           payload.PeerTypeUser,
		PeerID:             200,
		UserMessageID:      77,
		PeerSeq:            7,
		CanonicalMessageID: 7001,
	}
	req := buildBatchSendRequestForTest(100, 200, 9001, []int64{11, 12})
	groupedID := int64(7001)
	forwardDate := int32(1_772_000_010)
	replyToMsgID := int32(77)
	req.Message[0] = msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
		RandomId:        11,
		ForwardSourceId: &replyToMsgID,
		Message: tg.MakeTLMessage(&tg.TLMessage{
			Message: "caption",
			Media: tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
				Photo: tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: 333}),
			}),
			GroupedId: &groupedID,
			FwdFrom: tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
				FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 300}),
				Date:   forwardDate,
			}),
			Entities: []tg.MessageEntityClazz{
				tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{Offset: 0, Length: 7}),
			},
			ReplyTo: tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &replyToMsgID}),
		}),
	})

	got, err := core.MsgSendMessage(req)
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if _, ok := got.Clazz.(*tg.TLUpdates); !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
	if len(fakeUpdates.receiverPublished) != 1 {
		t.Fatalf("receiver published len = %d, want 1", len(fakeUpdates.receiverPublished))
	}
	receiverBatch := mustDecodeMessageOperationBatchV1(t, fakeUpdates.receiverPublished[0].Payload)
	if len(receiverBatch.Messages) != 2 {
		t.Fatalf("receiver batch len = %d, want 2", len(receiverBatch.Messages))
	}
	receiverMessage := receiverBatch.Messages[0]
	if receiverMessage.PeerID != 100 || receiverMessage.ToUserID != 200 || receiverMessage.MessageText != "caption" {
		t.Fatalf("unexpected receiver viewer fields: %+v", receiverMessage)
	}
	if receiverMessage.MediaRef == nil || receiverMessage.MediaRef.Kind != "photo" || receiverMessage.MediaRef.ID != 333 {
		t.Fatalf("receiver media ref = %+v, want photo 333", receiverMessage.MediaRef)
	}
	if receiverMessage.Attrs == nil || receiverMessage.Attrs.GroupedID != groupedID {
		t.Fatalf("receiver attrs = %+v, want grouped_id %d", receiverMessage.Attrs, groupedID)
	}
	if receiverMessage.ForwardRef == nil || receiverMessage.ForwardRef.FromUserID != 300 {
		t.Fatalf("receiver forward ref = %+v, want from user 300", receiverMessage.ForwardRef)
	}
	if receiverMessage.ForwardRef.SourceMessageID != 0 {
		t.Fatalf("receiver source_message_id = %d, want empty for non-channel user forward", receiverMessage.ForwardRef.SourceMessageID)
	}
	if receiverMessage.ForwardRef.SavedFromPeerType != 0 || receiverMessage.ForwardRef.SavedFromPeerID != 0 || receiverMessage.ForwardRef.SavedFromMessageID != 0 {
		t.Fatalf("receiver saved forward fields = %+v, want empty for non-saved forward", receiverMessage.ForwardRef)
	}
	if len(receiverMessage.Entities) != 1 || receiverMessage.Entities[0].Kind != "bold" {
		t.Fatalf("receiver entities = %+v, want bold entity", receiverMessage.Entities)
	}
	if receiverMessage.ReplyToCanonicalMessageID != 7001 {
		t.Fatalf("reply_to_canonical_message_id = %d, want 7001", receiverMessage.ReplyToCanonicalMessageID)
	}
	tlUpdates, ok := got.Clazz.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
	var sentMessage *tg.TLMessage
	for _, update := range tlUpdates.Updates {
		if updateNewMessage, ok := update.(*tg.TLUpdateNewMessage); ok {
			sentMessage, _ = updateNewMessage.Message.(*tg.TLMessage)
			break
		}
	}
	if sentMessage == nil || sentMessage.FwdFrom == nil {
		t.Fatalf("sent forward message missing: %+v", got)
	}
	if sentMessage.FwdFrom.SavedFromPeer != nil || sentMessage.FwdFrom.SavedFromMsgId != nil {
		t.Fatalf("sent fwd_from saved fields = %+v, want nil for non-saved forward", sentMessage.FwdFrom)
	}
	if sentMessage.FwdFrom.ChannelPost != nil {
		t.Fatalf("sent fwd_from channel_post = %v, want nil for non-channel user forward", sentMessage.FwdFrom.ChannelPost)
	}
}

func TestNormalizeOutboxMessageKeepsSavedForwardFieldsOnlyForSavedMessages(t *testing.T) {
	sourceID := int32(77)
	repo := &fakeMsgRepository{forwardVisible: true}
	got, err := normalizeOutboxMessage(normalizeOutboxInput{
		Ctx:          context.Background(),
		SenderUserID: 100,
		PeerType:     payload.PeerTypeUser,
		PeerID:       100,
		Repo:         repo,
		Outbox: msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			RandomId:        11,
			ForwardSourceId: &sourceID,
			Message: tg.MakeTLMessage(&tg.TLMessage{
				Message:     "saved forward",
				SavedPeerId: tg.MakePeerUser(200),
				FwdFrom: tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
					FromId:         tg.MakePeerUser(300),
					Date:           1_772_000_010,
					SavedFromPeer:  tg.MakePeerUser(200),
					SavedFromMsgId: &sourceID,
				}),
			}),
		}),
	})
	if err != nil {
		t.Fatalf("normalizeOutboxMessage() error = %v", err)
	}
	if repo.resolveForwardInput.SourcePeerID != 200 || repo.resolveForwardInput.SourceUserMessageID != 77 {
		t.Fatalf("forward resolve input = %+v, want source peer 200 message 77", repo.resolveForwardInput)
	}
	if got.ForwardRef == nil ||
		got.ForwardRef.SourcePeerID != 200 ||
		got.ForwardRef.SourceMessageID != 0 ||
		got.ForwardRef.SavedFromPeerID != 200 ||
		got.ForwardRef.SavedFromMessageID != 77 {
		t.Fatalf("forward ref = %+v, want saved and source fields", got.ForwardRef)
	}
	if peer, ok := sentMessageForwardPeer(got.ForwardRef).(*tg.TLPeerUser); !ok || peer.UserId != 300 {
		t.Fatalf("sent forward peer = %#v, want original author peerUser 300", sentMessageForwardPeer(got.ForwardRef))
	}
}

func TestNormalizeOutboxMessagePreservesChatForwardSourcePeer(t *testing.T) {
	sourceID := int32(77)
	repo := &fakeMsgRepository{forwardVisible: true}
	got, err := normalizeOutboxMessage(normalizeOutboxInput{
		Ctx:          context.Background(),
		SenderUserID: 1001,
		PeerType:     payload.PeerTypeChat,
		PeerID:       55,
		Repo:         repo,
		Outbox: msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			RandomId:        11,
			ForwardSourceId: &sourceID,
			Message: tg.MakeTLMessage(&tg.TLMessage{
				Message: "chat forward",
				FwdFrom: tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
					FromId:        tg.MakePeerUser(2002),
					Date:          1_772_000_010,
					SavedFromPeer: tg.MakePeerChat(44),
				}),
			}),
		}),
	})
	if err != nil {
		t.Fatalf("normalizeOutboxMessage() error = %v", err)
	}
	if repo.resolveForwardInput.SourcePeerType != payload.PeerTypeChat || repo.resolveForwardInput.SourcePeerID != 44 || repo.resolveForwardInput.SourceUserMessageID != 77 {
		t.Fatalf("forward resolve input = %+v, want chat 44 message 77", repo.resolveForwardInput)
	}
	if got.ForwardRef == nil ||
		got.ForwardRef.FromUserID != 2002 ||
		got.ForwardRef.SourcePeerType != payload.PeerTypeChat ||
		got.ForwardRef.SourcePeerID != 44 ||
		got.ForwardRef.SavedFromPeerType != 0 ||
		got.ForwardRef.SavedFromPeerID != 0 ||
		got.ForwardRef.SavedFromMessageID != 0 {
		t.Fatalf("forward ref = %+v, want sender user and non-saved chat source", got.ForwardRef)
	}
	if peer, ok := sentMessageForwardPeer(got.ForwardRef).(*tg.TLPeerChat); !ok || peer.ChatId != 44 {
		t.Fatalf("sent forward peer = %#v, want source peerChat 44", sentMessageForwardPeer(got.ForwardRef))
	}
}

func TestNormalizeOutboxMessageResolvesReplyInChatScope(t *testing.T) {
	replyToMsgID := int32(42)
	repo := &fakeMsgRepository{resolvedMessageID: &repository.ResolvedMessageID{
		UserID:             1001,
		PeerType:           payload.PeerTypeChat,
		PeerID:             55,
		UserMessageID:      42,
		CanonicalMessageID: 7001,
	}}
	got, err := normalizeOutboxMessage(normalizeOutboxInput{
		Ctx:          context.Background(),
		SenderUserID: 1001,
		PeerType:     payload.PeerTypeChat,
		PeerID:       55,
		Repo:         repo,
		Outbox: msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			RandomId: 99,
			Message: tg.MakeTLMessage(&tg.TLMessage{
				Message: "reply",
				ReplyTo: tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{
					ReplyToMsgId: &replyToMsgID,
				}),
			}),
		}),
	})
	if err != nil {
		t.Fatalf("normalizeOutboxMessage() error = %v", err)
	}
	if repo.resolveInput.UserID != 1001 || repo.resolveInput.PeerType != payload.PeerTypeChat || repo.resolveInput.PeerID != 55 || repo.resolveInput.UserMessageID != 42 {
		t.Fatalf("reply resolver input = %+v, want chat-scoped lookup", repo.resolveInput)
	}
	if got.ReplyToCanonicalMessageID != 7001 || got.ReplyToUserMessageID != 42 {
		t.Fatalf("reply ids = canonical:%d user:%d, want 7001/42", got.ReplyToCanonicalMessageID, got.ReplyToUserMessageID)
	}
}

func TestMsgSendMessageMediaReturnsTLUpdates(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":31,"pts_count":1,"event_type":"new_message","user_message_id":61}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 9101,
			PeerSeq:            31,
			MessageDate:        1_772_000_200,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              100,
			OperationId:         payload.SenderOperationID(9101, 100),
			Status:              1,
			Pts:                 31,
			PtsCount:            1,
			CurrentPts:          31,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: &fakeReceiverPublisher{},
	})

	got, err := core.MsgSendMessage(buildMediaSendRequestForTest(100, 200, 9001, 11, "caption"))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	updates, ok := got.Clazz.(*tg.TLUpdates)
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
	if idUpdate.Id != 61 || idUpdate.RandomId != 11 {
		t.Fatalf("updateMessageID = %+v, want id 61 random_id 11", idUpdate)
	}
	update, ok := updates.Updates[1].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("second update = %T, want *tg.TLUpdateNewMessage", updates.Updates[1])
	}
	message, ok := update.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", update.Message)
	}
	if message.Media == nil || message.Message != "caption" {
		t.Fatalf("message = %+v, want media caption", message)
	}
	media, ok := message.Media.(*tg.TLMessageMediaPhoto)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaPhoto", message.Media)
	}
	photo, ok := media.Photo.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("photo = %T, want *tg.TLPhoto", media.Photo)
	}
	if photo.Id != 333 || photo.AccessHash != 444 || len(photo.FileReference) != 25 || photo.DcId != 2 || len(photo.Sizes) != 2 {
		t.Fatalf("photo = %+v, want displayable photo metadata", photo)
	}
	stripped, ok := photo.Sizes[0].(*tg.TLPhotoStrippedSize)
	if !ok {
		t.Fatalf("projected size = %T, want TLPhotoStrippedSize", photo.Sizes[0])
	}
	if !bytes.Equal(stripped.Bytes, []byte{0x01, 0x16, 0x28, 0xaa}) {
		t.Fatalf("stripped bytes = %x, want telegram stripped payload", stripped.Bytes)
	}
	progressive, ok := photo.Sizes[1].(*tg.TLPhotoSizeProgressive)
	if !ok {
		t.Fatalf("projected size = %T, want TLPhotoSizeProgressive", photo.Sizes[1])
	}
	if progressive.Sizes[2] != 300 {
		t.Fatalf("progressive sizes = %#v, want final offset 300", progressive.Sizes)
	}
}

func TestMsgSendMessageContactMediaReturnsTLUpdates(t *testing.T) {
	core := newSingleSendMsgCoreForTest(t, 100, 9101, 61)
	req := sendMessageRequest(100, 200, 9001, "")
	req.Message[0].Message.(*tg.TLMessage).Media = tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{
		PhoneNumber: "8613000000001",
		FirstName:   "13000000001",
		LastName:    "t2",
		UserId:      1571266964,
	})

	got, err := core.MsgSendMessage(req)
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	message := firstSentUpdateMessage(t, got)
	media, ok := message.Media.(*tg.TLMessageMediaContact)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaContact", message.Media)
	}
	if media.PhoneNumber != "8613000000001" || media.FirstName != "13000000001" || media.LastName != "t2" || media.UserId != 1571266964 {
		t.Fatalf("contact media = %+v, want shared contact fields", media)
	}
}

func TestMsgSendMessageGroupedAndForwardReturnTLUpdates(t *testing.T) {
	t.Run("grouped", func(t *testing.T) {
		core := newSingleSendMsgCoreForTest(t, 100, 9102, 62)
		req := sendMessageRequest(100, 200, 9001, "album caption")
		groupedID := int64(7002)
		req.Message[0].Message.(*tg.TLMessage).GroupedId = &groupedID

		got, err := core.MsgSendMessage(req)
		if err != nil {
			t.Fatalf("MsgSendMessage() error = %v", err)
		}
		message := firstSentUpdateMessage(t, got)
		if message.GroupedId == nil || *message.GroupedId != groupedID {
			t.Fatalf("grouped_id = %v, want %d", message.GroupedId, groupedID)
		}
	})

	t.Run("forward", func(t *testing.T) {
		core := newSingleSendMsgCoreForTest(t, 100, 9103, 63)
		req := sendMessageRequest(100, 200, 9001, "forwarded")
		sourceID := int32(7)
		req.Message[0].Message.(*tg.TLMessage).FwdFrom = tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
			FromId:         tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 300}),
			Date:           1_772_000_011,
			SavedFromPeer:  tg.MakePeerUser(200),
			SavedFromMsgId: &sourceID,
		})

		got, err := core.MsgSendMessage(req)
		if err != nil {
			t.Fatalf("MsgSendMessage() error = %v", err)
		}
		message := firstSentUpdateMessage(t, got)
		if message.FwdFrom == nil {
			t.Fatalf("fwd_from = nil, want forward header")
		}
	})
}

func TestMsgSendMessageBatchTooLargeFailsBeforeRepositoryBatchCreate(t *testing.T) {
	core, fakeRepo, _ := newBatchMsgCoreForTest(t)
	_, err := core.MsgSendMessage(buildBatchSendRequestForTest(100, 200, 9001, makeSequentialRandomIDsForTest(101)))
	if !errors.Is(err, msgpb.ErrBatchTooLarge) {
		t.Fatalf("MsgSendMessage() error = %v, want %v", err, msgpb.ErrBatchTooLarge)
	}
	if fakeRepo.batchCreateCalls != 0 {
		t.Fatalf("batchCreateCalls = %d, want 0", fakeRepo.batchCreateCalls)
	}
}

func TestMsgSendMessageClearDraftWritesSenderOperationPayload(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":15,"pts_count":1,"event_type":"new_message","user_message_id":45}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 6001,
			PeerSeq:            9,
			MessageDate:        1_772_000_050,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(6001, 1001),
			Status:              1,
			Pts:                 15,
			PtsCount:            1,
			CurrentPts:          15,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: &fakeReceiverPublisher{},
	})

	sourceAuth := int64(9001)
	clearBefore := int32(1_772_000_049)
	req := sendMessageRequest(1001, 1002, 9001, "hello")
	req.ClearDraft = true
	req.SourcePermAuthKeyId = &sourceAuth
	req.ClearDraftBeforeDate = &clearBefore

	if _, err := core.MsgSendMessage(req); err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	senderOp := mustDecodeMessageOperationV4(t, updatesClient.processed.Payload)
	if !senderOp.MessageFact.ClearDraft || senderOp.MessageFact.SourcePermAuthKeyID != sourceAuth || senderOp.MessageFact.ClearDraftBeforeDate != clearBefore {
		t.Fatalf("unexpected clear draft payload: %+v", senderOp)
	}
}

func TestMsgSendMessagePeerTypeChatClearDraftSideEffect(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":22,"pts_count":1,"event_type":"new_message","user_message_id":43}`)
	responseHash := mustHashBytes(t, responsePayload)
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(7001, 1001),
			Status:              1,
			Pts:                 22,
			PtsCount:            1,
			CurrentPts:          22,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002}}
	repo := newFakeSendMessageRepositoryForChat(7001, 1)
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})

	sourceAuth := int64(9001)
	clearBefore := int32(1_772_000_049)
	req := &msgpb.TLMsgSendMessage{
		UserId:               1001,
		AuthKeyId:            9001,
		PeerType:             payload.PeerTypeChat,
		PeerId:               55,
		ClearDraft:           true,
		SourcePermAuthKeyId:  &sourceAuth,
		ClearDraftBeforeDate: &clearBefore,
		Message: []msgpb.OutboxMessageClazz{msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			RandomId: 12345,
			Message: tg.MakeTLMessage(&tg.TLMessage{
				Out:     true,
				FromId:  tg.MakePeerUser(1001),
				PeerId:  tg.MakePeerChat(55),
				Message: "clear chat draft",
				Date:    1778584910,
			}),
		})},
	}

	if _, err := core.MsgSendMessage(req); err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if updatesClient.processWithEffects == nil {
		t.Fatal("UserupdatesProcessUserOperationWithEffects was not called")
	}
	senderOp := mustDecodeMessageOperationV4(t, updatesClient.processWithEffects.Operation.Payload)
	if !senderOp.MessageFact.ClearDraft || senderOp.MessageFact.SourcePermAuthKeyID != sourceAuth || senderOp.MessageFact.ClearDraftBeforeDate != clearBefore {
		t.Fatalf("sender clear draft payload = %+v, want clear draft side effect", senderOp)
	}
	if len(updatesClient.processWithEffects.AffectedEffects) != 1 {
		t.Fatalf("affected effects = %d, want one receiver", len(updatesClient.processWithEffects.AffectedEffects))
	}
	receiverOp := mustDecodeMessageOperationV4(t, updatesClient.processWithEffects.AffectedEffects[0].Operation.Payload)
	if receiverOp.MessageFact.ClearDraft || receiverOp.MessageFact.SourcePermAuthKeyID != 0 || receiverOp.MessageFact.ClearDraftBeforeDate != 0 {
		t.Fatalf("receiver payload carried sender draft side effect: %+v", receiverOp)
	}
}

func TestMsgSendMessageReceiverDispatchUsesBrokerDurableAck(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":17,"pts_count":1,"event_type":"new_message","user_message_id":47}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 9001,
			PeerSeq:            11,
			MessageDate:        1_772_000_070,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(9001, 1001),
			Status:              1,
			Pts:                 17,
			PtsCount:            1,
			CurrentPts:          17,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgSendMessage(sendMessageRequest(1001, 1002, 9001, "broker ack"))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if _, ok := got.ToUpdateShortSentMessage(); !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != payload.ReceiverOperationID(9001, 1002) {
		t.Fatalf("unexpected receiver operation: %+v", publisher.published)
	}
	if publisher.published.PeerID != 1001 || publisher.published.PayloadCodec != payload.PayloadCodecJSON {
		t.Fatalf("unexpected receiver route payload metadata: %+v", publisher.published)
	}
	if len(updatesClient.processedList) != 1 || updatesClient.processWithEffects != nil {
		t.Fatalf("sender path should use requester sync only, processed=%d with_effects=%+v", len(updatesClient.processedList), updatesClient.processWithEffects)
	}
	if repo.markReceiverAckedCalls != 1 {
		t.Fatalf("mark receiver acked calls = %d, want 1", repo.markReceiverAckedCalls)
	}

	publishErr := errors.New("broker unavailable")
	repo = &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 2, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        2,
			CanonicalMessageID: 9002,
			PeerSeq:            12,
			MessageDate:        1_772_000_071,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient = &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(9002, 1001),
			Status:              1,
			Pts:                 18,
			PtsCount:            1,
			CurrentPts:          18,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher = &fakeReceiverPublisher{publishErr: publishErr}
	core = New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err = core.MsgSendMessage(sendMessageRequest(1001, 1002, 9001, "broker fail"))
	if err == nil {
		t.Fatalf("MsgSendMessage() error = nil, got=%+v", got)
	}
	if !errors.Is(err, msgpb.ErrReceiverBackpressure) {
		t.Fatalf("MsgSendMessage() error = %v, want ErrReceiverBackpressure", err)
	}
	if !errors.Is(err, publishErr) {
		t.Fatalf("MsgSendMessage() error = %v, want upstream publish error", err)
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
	if repo.markReceiverAckedCalls != 0 {
		t.Fatalf("mark receiver acked calls = %d, want 0 after publish failure", repo.markReceiverAckedCalls)
	}
}

func TestMarshalSendRequestHashIgnoresClearDraftBeforeDate(t *testing.T) {
	_, firstHash, err := marshalSendRequest(1001, payload.PeerTypeUser, 1001, 77, "hello", 0, true, 9001, 1_778_160_035)
	if err != nil {
		t.Fatalf("marshalSendRequest(first) error = %v", err)
	}
	_, retryHash, err := marshalSendRequest(1001, payload.PeerTypeUser, 1001, 77, "hello", 0, true, 9001, 1_778_160_066)
	if err != nil {
		t.Fatalf("marshalSendRequest(retry) error = %v", err)
	}
	if string(firstHash) != string(retryHash) {
		t.Fatalf("request hash changed when only clear_draft_before_date changed: first=%x retry=%x", firstHash, retryHash)
	}

	_, changedTextHash, err := marshalSendRequest(1001, payload.PeerTypeUser, 1001, 77, "changed", 0, true, 9001, 1_778_160_066)
	if err != nil {
		t.Fatalf("marshalSendRequest(changed text) error = %v", err)
	}
	if string(firstHash) == string(changedTextHash) {
		t.Fatalf("request hash did not change when message text changed: hash=%x", firstHash)
	}
}

func TestMsgSendMessageReplyToPayloadUsesPublicIDResolution(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":16,"pts_count":1,"event_type":"new_message","user_message_id":43}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 8001,
			PeerSeq:            10,
			MessageDate:        1_772_000_060,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
		resolvedMessageID: &repository.ResolvedMessageID{
			UserID:             1001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             1002,
			UserMessageID:      42,
			PeerSeq:            7,
			CanonicalMessageID: 7001,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(8001, 1001),
			Status:              1,
			Pts:                 16,
			PtsCount:            1,
			CurrentPts:          16,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	req := sendMessageRequest(1001, 1002, 9001, "reply body")
	replyToMsgID := int32(42)
	req.Message[0].Message.(*tg.TLMessage).ReplyTo = tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{
		ReplyToMsgId: &replyToMsgID,
	})

	if _, err := core.MsgSendMessage(req); err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if repo.resolveInput.UserMessageID != 42 || repo.resolveInput.PeerSeq != 0 {
		t.Fatalf("reply resolver input = %+v, want public user_message_id 42", repo.resolveInput)
	}
	for name, body := range map[string][]byte{
		"sender":   updatesClient.processed.Payload,
		"receiver": publisher.published.Payload,
	} {
		op := mustDecodeMessageOperationV4(t, body)
		if op.MessageFact.ReplyToCanonicalMessageID != 7001 {
			t.Fatalf("%s reply_to_canonical_message_id = %v, want 7001; payload=%s", name, op.MessageFact.ReplyToCanonicalMessageID, string(body))
		}
		if op.MessageFact.ReplyToUserMessageID != 0 {
			t.Fatalf("%s payload must not include reply_to_user_message_id for retry hash compatibility: %s", name, string(body))
		}
	}
}

func TestMsgSendMessageRecoversSenderCommitFromUserUpdatesResult(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":12,"pts_count":1,"event_type":"new_message","user_message_id":44}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 3001,
			PeerSeq:            6,
			MessageDate:        1_772_000_010,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
		markSenderErrs: []error{errors.New("temporary mysql failure"), nil},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(3001, 1001),
			Status:              1,
			Pts:                 12,
			PtsCount:            1,
			CurrentPts:          12,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
		getResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(3001, 1001),
			Status:              1,
			Pts:                 12,
			PtsCount:            1,
			CurrentPts:          12,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgSendMessage(sendMessageRequest(1001, 1002, 9001, "recover"))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if _, ok := got.ToUpdateShortSentMessage(); !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if repo.markSenderCalls != 2 {
		t.Fatalf("mark sender calls = %d, want 2", repo.markSenderCalls)
	}
	if updatesClient.getOperationResultCalls != 1 {
		t.Fatalf("get operation result calls = %d, want 1", updatesClient.getOperationResultCalls)
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
}

func TestMsgSendMessageRetrySkipsCanonicalMarkWhenAlreadyCanonical(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":13,"pts_count":1,"event_type":"new_message","user_message_id":46}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{
			SendStateID:        1,
			Status:             repository.SendStateStatusCanonical,
			CanonicalMessageID: 4001,
			PeerSeq:            7,
		},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 4001,
			PeerSeq:            7,
			MessageDate:        1_772_000_030,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         false,
		},
		markCanonicalErr: errors.New("canonical mark should not be retried"),
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(4001, 1001),
			Status:              1,
			Pts:                 13,
			PtsCount:            1,
			CurrentPts:          13,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgSendMessage(sendMessageRequest(1001, 1002, 9001, "retry"))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	if _, ok := got.ToUpdateShortSentMessage(); !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if repo.markCanonicalCalls != 0 {
		t.Fatalf("mark canonical calls = %d, want 0", repo.markCanonicalCalls)
	}
	if repo.markSenderCalls != 1 || publisher.calls != 1 || repo.markCompletedCalls != 1 {
		t.Fatalf("unexpected retry call counts: repo=%+v publisher_calls=%d", repo, publisher.calls)
	}
}

func TestMsgSendMessageRetryableWithoutSenderResultReprocessesSender(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":15,"pts_count":1,"event_type":"new_message","user_message_id":50}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{
			SendStateID:        1,
			Status:             repository.SendStateStatusFailedRetryable,
			CanonicalMessageID: 6001,
			PeerSeq:            9,
		},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 6001,
			PeerSeq:            9,
			MessageDate:        1_772_000_050,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         false,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(6001, 1001),
			Status:              1,
			Pts:                 15,
			PtsCount:            1,
			CurrentPts:          15,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: &fakeReceiverPublisher{},
	})

	got, err := core.MsgSendMessage(sendMessageRequest(1001, 1001, 9001, "retryable"))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	short, ok := got.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if short.Pts != 15 || short.PtsCount != 1 || short.Id != 50 {
		t.Fatalf("unexpected short sent message: %+v", short)
	}
	if updatesClient.processed == nil {
		t.Fatal("UserupdatesProcessUserOperation was not called")
	}
	if repo.markSenderCalls != 1 {
		t.Fatalf("mark sender calls = %d, want 1", repo.markSenderCalls)
	}
}

func TestMsgSendMessageSelfRetryUsesCommittedSenderState(t *testing.T) {
	responsePayload := mustOperationResponseV3Payload(t, payload.SenderOperationID(5001, 1001), 14, 1, 48, 77, replyShortSentMessageForText(t, 48, 14, 1))
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{
			SendStateID:             1,
			Status:                  repository.SendStateStatusSenderCommitted,
			CanonicalMessageID:      5001,
			PeerSeq:                 8,
			SenderOperationID:       payload.SenderOperationID(5001, 1001),
			SenderPTS:               14,
			SenderPTSCount:          1,
			SenderUpdatePayload:     responsePayload,
			SenderUpdatePayloadHash: responseHash,
		},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 5001,
			PeerSeq:            8,
			MessageDate:        1_772_000_040,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         false,
		},
	}
	updatesClient := &fakeUserUpdatesClient{}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgSendMessage(sendMessageRequest(1001, 1001, 9001, "self retry"))
	if err != nil {
		t.Fatalf("MsgSendMessage() error = %v", err)
	}
	short, ok := got.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if short.Pts != 14 || short.PtsCount != 1 || short.Id != 48 {
		t.Fatalf("unexpected short sent message: %+v", short)
	}
	if updatesClient.processed != nil || repo.markSenderCalls != 0 || publisher.calls != 0 {
		t.Fatalf("unexpected retry side effects: processed=%+v repo=%+v publisher_calls=%d", updatesClient.processed, repo, publisher.calls)
	}
	if repo.markReceiverAckedCalls != 1 || repo.markCompletedCalls != 1 {
		t.Fatalf("unexpected completion calls: repo=%+v", repo)
	}
}

func TestMsgGetHistoryReturnsCanonicalTextMessages(t *testing.T) {
	repo := &fakeMsgRepository{
		historyCursorBounds: repository.HistoryCursorBounds{
			OffsetPeerSeq: 13,
			MaxPeerSeq:    12,
			MinPeerSeq:    11,
		},
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID:   9001,
				PeerSeq:              2,
				UserMessageID:        102,
				ReplyToUserMessageID: 101,
				FromUserID:           1001,
				Outgoing:             true,
				PeerType:             payload.PeerTypeUser,
				PeerID:               1002,
				MessageKind:          repository.MessageKindText,
				MessageText:          "second",
				MessageDate:          1_772_000_020,
			},
			{
				CanonicalMessageID: 9000,
				PeerSeq:            1,
				UserMessageID:      101,
				FromUserID:         1001,
				Outgoing:           true,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				MessageKind:        repository.MessageKindText,
				MessageText:        "first",
				MessageDate:        1_772_000_010,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		OffsetId:  103,
		AddOffset: -2,
		MaxId:     102,
		MinId:     101,
		Limit:     20,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("expected messages.messages, got %s", got.ClazzName())
	}
	if len(messages.Messages) != 2 {
		t.Fatalf("messages len = %d, want 2", len(messages.Messages))
	}
	newest, ok := messages.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("message[0] = %T, want *tg.TLMessage", messages.Messages[0])
	}
	if newest.Id != 102 || newest.Message != "second" || newest.Date != 1_772_000_020 || !newest.Out {
		t.Fatalf("unexpected newest message: %+v", newest)
	}
	reply, ok := newest.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok {
		t.Fatalf("newest ReplyTo = %T, want *tg.TLMessageReplyHeader", newest.ReplyTo)
	}
	if reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 101 {
		t.Fatalf("reply_to_msg_id = %v, want 101", reply.ReplyToMsgId)
	}
	if repo.historyInput.PeerType != payload.PeerTypeUser ||
		repo.historyInput.UserID != 1001 ||
		repo.historyInput.PeerID != 1002 ||
		repo.historyInput.AddOffset != -2 ||
		repo.historyInput.Limit != 20 ||
		!repo.historyInput.CursorsResolved ||
		repo.historyInput.ResolvedCursorBounds.OffsetPeerSeq != 13 ||
		repo.historyInput.ResolvedCursorBounds.MaxPeerSeq != 12 ||
		repo.historyInput.ResolvedCursorBounds.MinPeerSeq != 11 {
		t.Fatalf("unexpected history input: %+v", repo.historyInput)
	}
	if repo.resolveHistoryInput.offsetID != 103 || repo.resolveHistoryInput.maxID != 102 || repo.resolveHistoryInput.minID != 101 {
		t.Fatalf("history cursors were not resolved from public ids: %+v", repo.resolveHistoryInput)
	}
}

func TestMsgGetHistoryReturnsMediaFromViewPayload(t *testing.T) {
	viewPayload, err := json.Marshal(payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 9002,
		PeerSeq:            3,
		MessageID:          103,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1_772_000_030,
		Out:                true,
		MessageText:        "photo caption",
		MediaRef: &payload.MediaRefV1{
			SchemaVersion: payload.MediaRefSchemaVersionV1,
			Kind:          "photo",
			ID:            7001,
			AccessHash:    8001,
			FileReference: []byte("1234567890123456789012345"),
			Date:          1_772_000_030,
			DcID:          2,
			PhotoSizes: []payload.PhotoSizeRefV1{
				{Kind: "stripped", Type: "i", Bytes: []byte{0x01, 0x16, 0x28, 0xaa}},
				{Kind: "progressive", Type: "y", W: 1280, H: 394, Sizes: []int32{100, 200, 300}},
			},
		},
	})
	if err != nil {
		t.Fatalf("marshal view payload: %v", err)
	}
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9002,
				PeerSeq:            3,
				UserMessageID:      103,
				FromUserID:         1001,
				Outgoing:           true,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				MessageKind:        repository.MessageKindText,
				MessageText:        "photo caption",
				MessageDate:        1_772_000_030,
				ViewPayload:        viewPayload,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("expected messages.messages, got %s", got.ClazzName())
	}
	if len(messages.Messages) != 1 {
		t.Fatalf("messages len = %d, want 1", len(messages.Messages))
	}
	message, ok := messages.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("message[0] = %T, want *tg.TLMessage", messages.Messages[0])
	}
	if message.Message != "photo caption" || message.Media == nil {
		t.Fatalf("history message missing caption/media: %+v", message)
	}
	media, ok := message.Media.(*tg.TLMessageMediaPhoto)
	if !ok {
		t.Fatalf("message media = %T, want *tg.TLMessageMediaPhoto", message.Media)
	}
	photo, ok := media.Photo.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("media photo = %T, want *tg.TLPhoto", media.Photo)
	}
	if photo.Id != 7001 || photo.AccessHash != 8001 || photo.DcId != 2 || len(photo.Sizes) != 2 {
		t.Fatalf("unexpected photo projection: %+v", photo)
	}
	stripped, ok := photo.Sizes[0].(*tg.TLPhotoStrippedSize)
	if !ok {
		t.Fatalf("projected size = %T, want TLPhotoStrippedSize", photo.Sizes[0])
	}
	if !bytes.Equal(stripped.Bytes, []byte{0x01, 0x16, 0x28, 0xaa}) {
		t.Fatalf("stripped bytes = %#v, want telegram stripped preview bytes", stripped.Bytes)
	}
	progressive, ok := photo.Sizes[1].(*tg.TLPhotoSizeProgressive)
	if !ok {
		t.Fatalf("projected size = %T, want TLPhotoSizeProgressive", photo.Sizes[1])
	}
	if progressive.Sizes[2] != 300 {
		t.Fatalf("progressive sizes = %#v, want final offset 300", progressive.Sizes)
	}
}

func TestMsgGetHistoryReturnsV4MessageFromViewPayload(t *testing.T) {
	viewPayload, err := json.Marshal(payload.MessageEventV4{
		SchemaVersion: payload.MessageEventSchemaVersionV4,
		EventKind:     payload.EventKindNewMessage,
		MessageFact: payload.NewMessageFactV1{
			SchemaVersion:      payload.MessageOperationSchemaVersionV4,
			CanonicalMessageID: 9004,
			PeerSeq:            4,
			PeerType:           payload.PeerTypeChat,
			PeerID:             6,
			SenderUserID:       1571266964,
			Date:               1_772_000_040,
			MessageText:        "v4 history",
		},
		MessageID: 104,
		Pts:       41,
		PtsCount:  1,
	})
	if err != nil {
		t.Fatalf("marshal view payload: %v", err)
	}
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9004,
				PeerSeq:            4,
				UserMessageID:      104,
				FromUserID:         1571266964,
				Outgoing:           false,
				PeerType:           payload.PeerTypeChat,
				PeerID:             6,
				MessageKind:        repository.MessageKindText,
				MessageText:        "v4 history",
				MessageDate:        1_772_000_040,
				ViewPayload:        viewPayload,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:   1571266963,
		PeerType: payload.PeerTypeChat,
		PeerId:   6,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("expected messages.messages, got %s", got.ClazzName())
	}
	if len(messages.Messages) != 1 {
		t.Fatalf("messages len = %d, want 1", len(messages.Messages))
	}
	message, ok := messages.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("message[0] = %T, want *tg.TLMessage", messages.Messages[0])
	}
	if message.Id != 104 || message.Message != "v4 history" || message.Out {
		t.Fatalf("history message = %+v", message)
	}
	peer, ok := message.PeerId.(*tg.TLPeerChat)
	if !ok || peer.ChatId != 6 {
		t.Fatalf("peer = %#v, want peerChat(6)", message.PeerId)
	}
	from, ok := message.FromId.(*tg.TLPeerUser)
	if !ok || from.UserId != 1571266964 {
		t.Fatalf("from = %#v, want peerUser(1571266964)", message.FromId)
	}
}

func TestMsgGetHistoryUsesViewerScopedOutgoingFlag(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9201,
				PeerSeq:            3,
				UserMessageID:      103,
				FromUserID:         1001,
				Outgoing:           false,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1001,
				MessageKind:        repository.MessageKindText,
				MessageText:        "saved from self as incoming projection",
				MessageDate:        1_772_000_030,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:   1002,
		PeerType: payload.PeerTypeUser,
		PeerId:   1001,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("expected messages.messages, got %s", got.ClazzName())
	}
	if len(messages.Messages) != 1 {
		t.Fatalf("messages len = %d, want 1", len(messages.Messages))
	}
	message, ok := messages.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("message[0] = %T, want *tg.TLMessage", messages.Messages[0])
	}
	if message.Out {
		t.Fatalf("message.Out = true, want false from viewer-scoped projection: %+v", message)
	}
	if message.FromId != nil {
		t.Fatalf("message.FromId = %#v, want nil for incoming private chat projection", message.FromId)
	}
	peer, ok := message.PeerId.(*tg.TLPeerUser)
	if !ok || peer.UserId != 1001 {
		t.Fatalf("message.PeerId = %#v, want peerUser(1001)", message.PeerId)
	}
}

func TestMsgGetHistoryMissingPublicCursorReturnsEmpty(t *testing.T) {
	repo := &fakeMsgRepository{
		historyCursorBounds: repository.HistoryCursorBounds{NoMatch: true},
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9202,
				PeerSeq:            4,
				UserMessageID:      104,
				FromUserID:         1002,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1001,
				MessageKind:        repository.MessageKindText,
				MessageText:        "must not leak across unresolved cursor",
				MessageDate:        1_772_000_040,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		MaxId:    999,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("expected messages.messages, got %s", got.ClazzName())
	}
	if len(messages.Messages) != 0 {
		t.Fatalf("messages len = %d, want empty for unresolved positive public cursor", len(messages.Messages))
	}
	if !repo.historyInput.ResolvedCursorBounds.NoMatch || repo.resolveHistoryInput.maxID != 999 {
		t.Fatalf("history no-match bounds not propagated: input=%+v resolved=%+v", repo.historyInput, repo.resolveHistoryInput)
	}
}

func TestMsgGetHistoryReturnsNotModifiedForMatchingHash(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9101,
				PeerSeq:            2,
				UserMessageID:      102,
				FromUserID:         1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				MessageKind:        repository.MessageKindText,
				MessageText:        "second",
				MessageDate:        1_772_000_020,
			},
			{
				CanonicalMessageID: 9100,
				PeerSeq:            1,
				UserMessageID:      101,
				FromUserID:         1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				MessageKind:        repository.MessageKindText,
				MessageText:        "first",
				MessageDate:        1_772_000_010,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		Limit:    20,
		Hash:     pagination.HashInt64IDs([]int64{102, 101}),
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	if _, ok := got.ToMessagesMessagesNotModified(); !ok {
		t.Fatalf("MsgGetHistory() = %s, want messages.messagesNotModified", got.ClazzName())
	}
}

func TestMsgGetHistoryPassesViewerUserID(t *testing.T) {
	repo := &fakeMsgRepository{}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:    1003,
		AuthKeyId: 9003,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Limit:     30,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	if repo.historyInput.UserID != 1003 {
		t.Fatalf("history input user_id = %d, want viewer user_id 1003", repo.historyInput.UserID)
	}
}

func TestMsgReadHistoryV2ReturnsAffectedMessagesAck(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: readHistoryOperationID(1001, 1002, 102, 9001),
			Status:      1,
			Pts:         15,
			PtsCount:    1,
			CurrentPts:  15,
		}),
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 102}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      102,
				PeerSeq:            2,
				CanonicalMessageID: 7002,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     102,
	})
	if err != nil {
		t.Fatalf("MsgReadHistoryV2() error = %v", err)
	}
	if got == nil {
		t.Fatal("MsgReadHistoryV2() returned nil")
	}
	if got.Pts != 15 || got.PtsCount != 1 {
		t.Fatalf("affected messages = %+v, want pts=15 pts_count=1", got)
	}
	if len(updatesClient.processedList) != 0 {
		t.Fatalf("direct processed operations = %d, want 0", len(updatesClient.processedList))
	}
	if updatesClient.processWithEffects == nil {
		t.Fatal("UserupdatesProcessUserOperationWithEffects was not called")
	}
	readerOperation := updatesClient.processWithEffects.Operation
	if readerOperation == nil {
		t.Fatal("with-effects requester operation is nil")
	}
	if readerOperation.OperationId != readHistoryOperationID(1001, 1002, 102, 9001) {
		t.Fatalf("reader operation_id = %q", readerOperation.OperationId)
	}
	if readerOperation.AuthKeyIdExclude == nil || *readerOperation.AuthKeyIdExclude != 9001 {
		t.Fatalf("auth_key_id_exclude = %v, want 9001", readerOperation.AuthKeyIdExclude)
	}
	var readerOp payload.MessageOperationV1
	if err := json.Unmarshal(readerOperation.Payload, &readerOp); err != nil {
		t.Fatalf("decode read history payload: %v", err)
	}
	if readerOp.OperationKind != payload.OperationKindReadHistory || readerOp.PeerID != 1002 || readerOp.ReadInboxMaxPeerSeq != 2 || readerOp.ReadMaxUserMessageID != 102 || readerOp.ReadOutboxMaxPeerSeq != 0 || readerOp.Out {
		t.Fatalf("unexpected reader read history payload: %+v", readerOp)
	}
	if len(updatesClient.processWithEffects.AffectedEffects) != 1 {
		t.Fatalf("affected effects = %d, want 1", len(updatesClient.processWithEffects.AffectedEffects))
	}
	affected := updatesClient.processWithEffects.AffectedEffects[0]
	if affected.RequesterUserId != 1001 || affected.DeliveryPolicy != int32(DeliveryPolicyDurableAsync) || affected.OperationKind != payload.OperationKindReadHistory {
		t.Fatalf("unexpected affected metadata: %+v", affected)
	}
	peerOperation := affected.Operation
	if peerOperation == nil {
		t.Fatal("affected peer operation is nil")
	}
	if peerOperation.UserId != 1002 || peerOperation.PeerId != 1001 {
		t.Fatalf("unexpected peer operation routing: %+v", peerOperation)
	}
	if peerOperation.OperationId != readHistoryOutboxOperationID(1002, 1001, 1001, 2) {
		t.Fatalf("peer read outbox operation_id = %q, want peer-seq scoped id", peerOperation.OperationId)
	}
	if peerOperation.OperationId == readHistoryOperationID(1002, 1001, 102, 0) {
		t.Fatalf("peer read outbox operation_id reused requester public max_id: %q", peerOperation.OperationId)
	}
	if peerOperation.AuthKeyIdExclude != nil {
		t.Fatalf("peer operation auth_key_id_exclude = %v, want nil", peerOperation.AuthKeyIdExclude)
	}
	var peerOp payload.MessageOperationV1
	if err := json.Unmarshal(peerOperation.Payload, &peerOp); err != nil {
		t.Fatalf("decode peer read outbox payload: %v", err)
	}
	if peerOp.OperationKind != payload.OperationKindReadHistory || peerOp.PeerID != 1001 || peerOp.ReadInboxMaxPeerSeq != 0 || peerOp.ReadOutboxMaxPeerSeq != 2 || peerOp.ReadMaxUserMessageID != 0 || !peerOp.Out {
		t.Fatalf("unexpected peer read outbox payload: %+v", peerOp)
	}
}

func TestMsgReadHistoryV2MissingPublicMaxIDNoopsWithoutPTS(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     102,
	})
	if err != nil {
		t.Fatalf("MsgReadHistoryV2() error = %v", err)
	}
	if got == nil || got.Pts != 0 || got.PtsCount != 0 {
		t.Fatalf("affected messages = %+v, want no-op pts=0 pts_count=0", got)
	}
	if updatesClient.processed != nil || updatesClient.processWithEffects != nil || len(updatesClient.processedList) != 0 {
		t.Fatalf("missing public max_id dispatched side effects: processed=%+v with_effects=%+v list=%d", updatesClient.processed, updatesClient.processWithEffects, len(updatesClient.processedList))
	}
	if repo.resolveInput.UserID != 1001 || repo.resolveInput.PeerType != payload.PeerTypeUser || repo.resolveInput.PeerID != 1002 || repo.resolveInput.UserMessageID != 102 {
		t.Fatalf("resolve input = %+v, want requester scoped public id", repo.resolveInput)
	}
}

func TestMsgReadHistoryV2SkipsSelfAffectedOutbox(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: readHistoryOperationID(1001, 1001, 2, 9001),
			Status:      1,
			Pts:         16,
			PtsCount:    1,
			CurrentPts:  16,
		}),
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1001, userMessageID: 2}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1001,
				UserMessageID:      2,
				PeerSeq:            2,
				CanonicalMessageID: 7002,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1001,
		MaxId:     2,
	})
	if err != nil {
		t.Fatalf("MsgReadHistoryV2() error = %v", err)
	}
	if got == nil || got.Pts != 16 || got.PtsCount != 1 {
		t.Fatalf("affected messages = %+v, want pts=16 pts_count=1", got)
	}
	if updatesClient.processWithEffects != nil {
		t.Fatalf("with-effects call = %+v, want nil for self read history", updatesClient.processWithEffects)
	}
	if len(updatesClient.processedList) != 1 {
		t.Fatalf("direct processed operations = %d, want 1", len(updatesClient.processedList))
	}
}

func TestMsgReadHistoryV2ReturnsErrorWhenDurableEffectAcceptFails(t *testing.T) {
	acceptErr := errors.New("affected effect accept failed")
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsErr: acceptErr,
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 2}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      2,
				PeerSeq:            2,
				CanonicalMessageID: 7002,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     2,
	})
	if err == nil {
		t.Fatalf("MsgReadHistoryV2() error = nil, got = %+v", got)
	}
	if !errors.Is(err, msgpb.ErrSenderSyncFailed) {
		t.Fatalf("MsgReadHistoryV2() error = %v, want ErrSenderSyncFailed", err)
	}
	if !errors.Is(err, acceptErr) {
		t.Fatalf("MsgReadHistoryV2() error = %v, want upstream accept error", err)
	}
	if updatesClient.processWithEffects == nil {
		t.Fatal("UserupdatesProcessUserOperationWithEffects was not called")
	}
	if len(updatesClient.processedList) != 0 {
		t.Fatalf("direct processed operations = %d, want 0", len(updatesClient.processedList))
	}
}

func TestMsgReadHistoryV2NilServiceContextReturnsSenderSyncFailed(t *testing.T) {
	core := New(context.Background(), nil)
	_, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     2,
	})
	if !errors.Is(err, msgpb.ErrSenderSyncFailed) {
		t.Fatalf("MsgReadHistoryV2() error = %v, want ErrSenderSyncFailed", err)
	}
}

func TestMsgUpdatePinnedMessageRoutesProjectionOperation(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: updatePinnedOperationID(1001, 1002, 107, false, 9001),
			Status:      1,
			Pts:         21,
			PtsCount:    1,
			CurrentPts:  21,
		}),
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 107}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      107,
				PeerSeq:            7,
				CanonicalMessageID: 7001,
				MessageDate:        1_772_000_123,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgUpdatePinnedMessage(&msgpb.TLMsgUpdatePinnedMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        107,
	})
	if err != nil {
		t.Fatalf("MsgUpdatePinnedMessage() error = %v", err)
	}
	short, ok := got.ToUpdateShort()
	if !ok {
		t.Fatalf("MsgUpdatePinnedMessage() = %s, want updateShort", got.ClazzName())
	}
	pinned, ok := (&tg.Update{Clazz: short.Update}).ToUpdatePinnedMessages()
	if !ok || !pinned.Pinned || len(pinned.Messages) != 1 || pinned.Messages[0] != 107 || pinned.Pts != 21 {
		t.Fatalf("pinned update = %+v ok=%v", pinned, ok)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &op); err != nil {
		t.Fatalf("decode update pinned payload: %v", err)
	}
	if op.OperationKind != payload.OperationKindUpdatePinnedMessage || op.PinnedPeerSeq != 7 || op.PinnedUserMessageID != 107 || op.PinnedCanonicalMessageID != 7001 {
		t.Fatalf("unexpected update pinned payload: %+v", op)
	}
	if updatesClient.processed.CanonicalPeerSeq == nil || *updatesClient.processed.CanonicalPeerSeq != 7 {
		t.Fatalf("canonical peer seq = %v, want resolved peer_seq 7", updatesClient.processed.CanonicalPeerSeq)
	}
}

func TestMsgUpdatePinnedMessageRejectsZeroIDForPin(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgUpdatePinnedMessage(&msgpb.TLMsgUpdatePinnedMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        0,
		Unpin:     false,
	})
	if !errors.Is(err, msgpb.ErrSendStateConflict) {
		t.Fatalf("MsgUpdatePinnedMessage() error = %v, want ErrSendStateConflict, got=%+v", err, got)
	}
	if updatesClient.processed != nil || len(updatesClient.processedList) != 0 || updatesClient.processWithEffects != nil {
		t.Fatalf("zero pin id dispatched operation: processed=%+v list=%d with_effects=%+v", updatesClient.processed, len(updatesClient.processedList), updatesClient.processWithEffects)
	}
	if repo.resolveInput.UserMessageID != 0 {
		t.Fatalf("zero pin id resolved message unexpectedly: %+v", repo.resolveInput)
	}
}

func TestMsgUpdatePinnedMessageAllowsZeroIDForUnpin(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: updatePinnedOperationID(1001, 1002, 0, true, 9001),
			Status:      1,
			Pts:         22,
			PtsCount:    1,
			CurrentPts:  22,
		}),
	}
	repo := &fakeMsgRepository{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgUpdatePinnedMessage(&msgpb.TLMsgUpdatePinnedMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		Unpin:     true,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        0,
	})
	if err != nil {
		t.Fatalf("MsgUpdatePinnedMessage() error = %v", err)
	}
	short, ok := got.ToUpdateShort()
	if !ok {
		t.Fatalf("MsgUpdatePinnedMessage() = %s, want updateShort", got.ClazzName())
	}
	pinned, ok := (&tg.Update{Clazz: short.Update}).ToUpdatePinnedMessages()
	if !ok || pinned.Pinned || len(pinned.Messages) != 0 || pinned.Pts != 22 {
		t.Fatalf("unpin update = %+v ok=%v", pinned, ok)
	}
	if repo.resolveInput.UserMessageID != 0 {
		t.Fatalf("zero unpin id resolved message unexpectedly: %+v", repo.resolveInput)
	}
}

func TestMsgResolveDialogCursorTopMessageResolvesGlobalPublicID(t *testing.T) {
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			42: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             2002,
				UserMessageID:      42,
				PeerSeq:            7,
				CanonicalMessageID: 7001,
				MessageDate:        1_772_000_000,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgResolveDialogCursorTopMessage(&msgpb.TLMsgResolveDialogCursorTopMessage{
		UserId:       1001,
		TopMessageId: 42,
	})
	if err != nil {
		t.Fatalf("MsgResolveDialogCursorTopMessage error = %v", err)
	}
	if got.Found != tg.BoolTrueClazz || got.PeerType != payload.PeerTypeUser || got.PeerId != 2002 ||
		got.PeerSeq != 7 || got.MessageDate != 1_772_000_000 {
		t.Fatalf("resolved cursor = %+v, want public id mapped to internal cursor", got)
	}
	if repo.resolveInput.UserID != 1001 || repo.resolveInput.UserMessageID != 42 {
		t.Fatalf("resolver input = %+v, want user/global public id", repo.resolveInput)
	}
}

func TestMsgResolveDialogCursorTopMessageReturnsNotFoundForUnknownPositiveID(t *testing.T) {
	repo := &fakeMsgRepository{resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{}}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgResolveDialogCursorTopMessage(&msgpb.TLMsgResolveDialogCursorTopMessage{
		UserId:       1001,
		TopMessageId: 404,
	})
	if err != nil {
		t.Fatalf("MsgResolveDialogCursorTopMessage error = %v", err)
	}
	if got.Found != tg.BoolFalseClazz || got.PeerSeq != 0 || got.PeerId != 0 {
		t.Fatalf("resolved cursor = %+v, want explicit not found", got)
	}
}

func TestMsgDeleteMessagesRoutesGlobalPublicIDsByRealPeer(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{107, 108}, false, 9001),
			Status:      1,
			Pts:         31,
			PtsCount:    1,
			CurrentPts:  31,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
			108: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 108, PeerSeq: 8, CanonicalMessageID: 7008},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  0,
		PeerId:    0,
		Id:        []int32{107, 108},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 31 || got.PtsCount != 1 {
		t.Fatalf("affected = %+v", got)
	}
	if updatesClient.processed == nil {
		t.Fatalf("delete operation was not dispatched")
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &op); err != nil {
		t.Fatalf("decode delete payload: %v", err)
	}
	if op.PeerID != 1002 || len(op.DeletePeerSeqs) != 2 || op.DeletePeerSeqs[0] != 7 || op.DeletePeerSeqs[1] != 8 {
		t.Fatalf("unexpected delete payload: %+v", op)
	}
	if len(op.DeleteUserMessageIDs) != 2 || op.DeleteUserMessageIDs[0] != 107 || op.DeleteUserMessageIDs[1] != 108 {
		t.Fatalf("unexpected public delete ids: %+v", op)
	}
}

func TestMsgDeleteMessagesRetryResolvesDeletedViewAndReturnsStoredResult(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, false, 9001),
			Status:      1,
			Pts:         51,
			PtsCount:    1,
			CurrentPts:  51,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{},
		resolveDeleteByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        []int32{107},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 51 || got.PtsCount != 1 {
		t.Fatalf("affected = %+v, want stored operation result", got)
	}
	if repo.resolveDeleteCalls != 1 {
		t.Fatalf("delete resolver calls = %d, want 1", repo.resolveDeleteCalls)
	}
	if updatesClient.processed == nil {
		t.Fatal("retry over deleted view should still dispatch requester operation")
	}
	if updatesClient.processed.OperationId != deleteMessagesOperationID(1001, 1002, []int32{107}, false, 9001) {
		t.Fatalf("operation id = %s", updatesClient.processed.OperationId)
	}
}

func TestBuildDeleteMessagesPayloadUsesStableDateAndHash(t *testing.T) {
	const deleteDate = int32(1_772_000_123)
	body1, hash1, err := buildDeleteMessagesPayload(1001, 1001, payload.PeerTypeUser, 1002, deleteDate, []int64{7}, []int64{107}, false)
	if err != nil {
		t.Fatalf("build payload 1: %v", err)
	}
	body2, hash2, err := buildDeleteMessagesPayload(1001, 1001, payload.PeerTypeUser, 1002, deleteDate, []int64{7}, []int64{107}, false)
	if err != nil {
		t.Fatalf("build payload 2: %v", err)
	}
	if string(body1) != string(body2) {
		t.Fatalf("delete payload changed across repeated builds:\n%s\n%s", body1, body2)
	}
	if string(hash1) != string(hash2) {
		t.Fatalf("delete payload hash changed: %x != %x", hash1, hash2)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(body1, &op); err != nil {
		t.Fatalf("decode payload: %v", err)
	}
	if op.Date != deleteDate {
		t.Fatalf("delete payload date = %d, want stable date %d", op.Date, deleteDate)
	}
}

func TestMsgDeleteMessagesGroupsGlobalIDsByPeer(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResults: []*userupdates.UserOperationResult{
			userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
				UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, false, 9001), Status: 1, Pts: 31, PtsCount: 1, CurrentPts: 31,
			}),
			userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
				UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1003, []int32{208}, false, 9001), Status: 1, Pts: 32, PtsCount: 1, CurrentPts: 32,
			}),
		},
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
			208: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1003, UserMessageID: 208, PeerSeq: 2, CanonicalMessageID: 8002},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		Id:        []int32{107, 208},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 32 || got.PtsCount != 2 {
		t.Fatalf("affected = %+v, want final pts 32 and two requester operations", got)
	}
	if len(updatesClient.processedOperations) != 2 {
		t.Fatalf("processed operations = %d, want 2", len(updatesClient.processedOperations))
	}
}

func TestMsgDeleteMessagesRevokeEnqueuesPeerEffect(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, true, 9001),
			Status:      1,
			Pts:         33,
			PtsCount:    1,
			CurrentPts:  33,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	_, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		Revoke:    true,
		Id:        []int32{107},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 1 {
		t.Fatalf("affected effects = %+v", updatesClient.processWithEffects)
	}
	effect := updatesClient.processWithEffects.AffectedEffects[0].ToAffectedUserOperation()
	if effect == nil || effect.Operation == nil || effect.Operation.UserId != 1002 || effect.Operation.PeerId != 1001 {
		t.Fatalf("peer effect = %+v", effect)
	}
	var peerOp payload.MessageOperationV1
	if err := json.Unmarshal(effect.Operation.Payload, &peerOp); err != nil {
		t.Fatalf("decode peer payload: %v", err)
	}
	if len(peerOp.DeleteUserMessageIDs) != 0 || len(peerOp.DeletePeerSeqs) != 1 || peerOp.DeletePeerSeqs[0] != 7 {
		t.Fatalf("peer delete payload = %+v", peerOp)
	}
}

func TestMsgDeleteMessagesRevokeSelfDoesNotEnqueuePeerEffect(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1001, []int32{3}, true, 9001), Status: 1, Pts: 34, PtsCount: 1, CurrentPts: 34,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			3: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1001, UserMessageID: 3, PeerSeq: 3, CanonicalMessageID: 7003},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})
	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{UserId: 1001, AuthKeyId: 9001, Revoke: true, Id: []int32{3}})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 34 || got.PtsCount != 1 {
		t.Fatalf("affected = %+v, want requester pts 34 and pts_count 1", got)
	}
	if updatesClient.processed == nil {
		t.Fatal("requester delete operation was not dispatched")
	}
	if updatesClient.processed.OperationId != deleteMessagesOperationID(1001, 1001, []int32{3}, true, 9001) {
		t.Fatalf("requester operation id = %s", updatesClient.processed.OperationId)
	}
	if updatesClient.processWithEffects != nil {
		t.Fatalf("self delete should not use with-effects dispatch: %+v", updatesClient.processWithEffects)
	}
}

func TestMsgDeleteMessagesRejectsPeerEmptyWithNonZeroPeerID(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, false, 9001),
			Status:      1,
			Pts:         31,
			PtsCount:    1,
			CurrentPts:  31,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  0,
		PeerId:    1002,
		Id:        []int32{107},
	})
	if err == nil {
		t.Fatalf("MsgDeleteMessages() error = nil, got = %+v", got)
	}
	if !errors.Is(err, msgpb.ErrSendStateConflict) {
		t.Fatalf("MsgDeleteMessages() error = %v, want ErrSendStateConflict", err)
	}
}

func TestMsgDeleteMessagesRevokeGroupsGlobalIDsByPeerUsesSequentialResults(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResults: []*userupdates.UserOperationResult{
			userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
				UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, true, 9001), Status: 1, Pts: 41, PtsCount: 1, CurrentPts: 41,
			}),
			userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
				UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1003, []int32{208}, true, 9001), Status: 1, Pts: 42, PtsCount: 1, CurrentPts: 42,
			}),
		},
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
			208: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1003, UserMessageID: 208, PeerSeq: 2, CanonicalMessageID: 8002},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		Revoke:    true,
		Id:        []int32{107, 208},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 42 || got.PtsCount != 2 {
		t.Fatalf("affected = %+v, want final pts 42 and two requester operations", got)
	}
	if len(updatesClient.processedOperations) != 2 {
		t.Fatalf("processed operations = %d, want 2", len(updatesClient.processedOperations))
	}
	if updatesClient.processedOperations[0].PeerId != 1002 || updatesClient.processedOperations[1].PeerId != 1003 {
		t.Fatalf("processed operation peers = %d,%d; want 1002,1003", updatesClient.processedOperations[0].PeerId, updatesClient.processedOperations[1].PeerId)
	}
}

func TestMsgDeleteMessagesIgnoresMissingPublicIDs(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{}}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        []int32{404},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 0 || got.PtsCount != 0 {
		t.Fatalf("affected = %+v, want empty no-op", got)
	}
	if updatesClient.processed != nil || updatesClient.processWithEffects != nil {
		t.Fatalf("missing public ids should not dispatch operations: processed=%+v with_effects=%+v", updatesClient.processed, updatesClient.processWithEffects)
	}
}

func TestMsgDeleteMessagesIgnoresResolvedIDsFromOtherPeer(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			207: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 2002, UserMessageID: 207, PeerSeq: 17, CanonicalMessageID: 7207},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        []int32{207},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 0 || got.PtsCount != 0 {
		t.Fatalf("affected = %+v, want no-op for other peer id", got)
	}
	if updatesClient.processed != nil || updatesClient.processWithEffects != nil {
		t.Fatalf("other peer public id dispatched delete: processed=%+v with_effects=%+v", updatesClient.processed, updatesClient.processWithEffects)
	}
}

func TestMsgDeleteHistoryRoutesProjectionOperation(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteHistoryOperationID(1001, 1002, 9, true, false, 9001),
			Status:      1,
			Pts:         32,
			PtsCount:    1,
			CurrentPts:  32,
		}),
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 9}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      9,
				PeerSeq:            99,
				CanonicalMessageID: 7099,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteHistory(&msgpb.TLMsgDeleteHistory{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		JustClear: true,
		MaxId:     9,
	})
	if err != nil {
		t.Fatalf("MsgDeleteHistory() error = %v", err)
	}
	if got.Pts != 32 || got.PtsCount != 1 {
		t.Fatalf("affected history = %+v", got)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &op); err != nil {
		t.Fatalf("decode delete history payload: %v", err)
	}
	if op.OperationKind != payload.OperationKindDeleteHistory || op.DeleteMaxPeerSeq != 99 || op.PeerSeq != 99 || !op.JustClear {
		t.Fatalf("unexpected delete history payload: %+v", op)
	}
	if updatesClient.processed.OperationId != deleteHistoryOperationID(1001, 1002, 99, true, false, 9001) {
		t.Fatalf("delete history operation_id = %q, want resolved peer-seq key", updatesClient.processed.OperationId)
	}
}

func TestMsgDeleteHistoryMissingPositiveMaxIDIsNoOp(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{}}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteHistory(&msgpb.TLMsgDeleteHistory{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     404,
	})
	if err != nil {
		t.Fatalf("MsgDeleteHistory() error = %v", err)
	}
	if got.Pts != 0 || got.PtsCount != 0 || got.Offset != 0 {
		t.Fatalf("affected history = %+v, want empty no-op", got)
	}
	if updatesClient.processed != nil || updatesClient.processWithEffects != nil {
		t.Fatalf("missing max_id should not dispatch operations: processed=%+v with_effects=%+v", updatesClient.processed, updatesClient.processWithEffects)
	}
}

func TestMsgSearchHashtagDelegatesPublicOffsetToRepository(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{{
			UserMessageID: 707,
			PeerSeq:       7,
			FromUserID:    1001,
			PeerType:      payload.PeerTypeUser,
			PeerID:        1002,
			MessageKind:   repository.MessageKindText,
			MessageText:   "#tag result",
			MessageDate:   1_772_000_001,
		}},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgSearchHashtag(&msgpb.TLMsgSearchHashtag{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		HashTag:  "#tag",
		OffsetId: 909,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgSearchHashtag() error = %v", err)
	}
	if repo.searchInput.OffsetID != 909 {
		t.Fatalf("search offset id = %d, want public offset id 909 delegated to repository", repo.searchInput.OffsetID)
	}
	if repo.resolveInput.UserMessageID != 0 {
		t.Fatalf("handler resolved offset id directly: %+v", repo.resolveInput)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok || len(messages.Messages) != 1 {
		t.Fatalf("messages = %+v ok=%v", got, ok)
	}
	msg, ok := messages.Messages[0].(*tg.TLMessage)
	if !ok || msg.Id != 707 {
		t.Fatalf("search result message = %+v ok=%v, want public id 707", messages.Messages[0], ok)
	}
}

func TestMsgSearchHashtagMissingPublicOffsetReturnsEmpty(t *testing.T) {
	repo := &fakeMsgRepository{
		searchNoMatch: true,
		history: []repository.HistoryMessage{{
			UserMessageID: 707,
			PeerSeq:       7,
			FromUserID:    1001,
			PeerType:      payload.PeerTypeUser,
			PeerID:        1002,
			MessageKind:   repository.MessageKindText,
			MessageText:   "#tag first page",
			MessageDate:   1_772_000_001,
		}},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgSearchHashtag(&msgpb.TLMsgSearchHashtag{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		HashTag:  "#tag",
		OffsetId: 909,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgSearchHashtag() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("messages = %s, want messages.messages", got.ClazzName())
	}
	if len(messages.Messages) != 0 {
		t.Fatalf("messages len = %d, want empty for missing positive offset", len(messages.Messages))
	}
}

func TestTask6DialogOperationIDsUseV2ResolvedIdentity(t *testing.T) {
	if got := readHistoryOperationID(1001, 1002, 102, 9001); got != "v2:dialog:read_history:user:1001:peer:1002:max_user:102:auth:9001" {
		t.Fatalf("readHistoryOperationID() = %q", got)
	}
	if got := readHistoryOutboxOperationID(1002, 1001, 1001, 2); got != "v2:dialog:read_history_outbox:user:1002:peer:1001:reader:1001:max_peer_seq:2" {
		t.Fatalf("readHistoryOutboxOperationID() = %q", got)
	}
	if got := deleteMessagesOperationID(1001, 1002, []int32{107, 108}, true, 9001); got != "v2:dialog:delete_messages:user:1001:peer:1002:ids:[107 108]:revoke:true:auth:9001" {
		t.Fatalf("deleteMessagesOperationID() = %q", got)
	}
	if got := deleteMessagesPeerSeqOperationID(1002, 1001, []int64{7, 8}, true); got != "v2:dialog:delete_messages:user:1002:peer:1001:peer_seqs:[7 8]:revoke:true" {
		t.Fatalf("deleteMessagesPeerSeqOperationID() = %q", got)
	}
	if got := deleteHistoryOperationID(1001, 1002, 99, true, false, 9001); got != "v2:dialog:delete_history:user:1001:peer:1002:max_peer_seq:99:clear:true:revoke:false:auth:9001" {
		t.Fatalf("deleteHistoryOperationID() = %q", got)
	}
	if got := updatePinnedOperationID(1001, 1002, 107, false, 9001); got != "v2:dialog:update_pinned:user:1001:peer:1002:id:107:unpin:false:auth:9001" {
		t.Fatalf("updatePinnedOperationID() = %q", got)
	}
}

func TestMsgEditMessageUpdatesCanonicalAndRoutesOperations(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":41,"pts_count":1,"user_message_id":107}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 107}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      107,
				PeerSeq:            7,
				CanonicalMessageID: 7001,
				MessageDate:        1_772_000_010,
			},
		},
		editResult: &repository.EditMessageResult{
			CanonicalMessageID: 7001,
			PeerSeq:            7,
			FromUserID:         1001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             4294968298,
			MessageText:        "edited",
			MessageDate:        1_772_000_010,
			EditDate:           1_772_000_100,
			EditVersion:        1,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         editMessageOperationID(7001, 1, 1001),
			Status:              1,
			Pts:                 41,
			PtsCount:            1,
			CurrentPts:          41,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgEditMessage(editMessageRequest(1001, 1002, 9001, 107, "edited"))
	if err != nil {
		t.Fatalf("MsgEditMessage() error = %v", err)
	}
	updates, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("expected updates, got %s", got.ClazzName())
	}
	if updates.Date != 1_772_000_099 || updates.Seq != 0 || len(updates.Updates) != 1 || len(updates.Users) != 0 {
		t.Fatalf("unexpected edit updates envelope: %+v", updates)
	}
	if updates.Users == nil {
		t.Fatalf("msg edit response must use an empty users vector, not nil, so BFF replacement is deterministic")
	}
	edit, ok := updates.Updates[0].(*tg.TLUpdateEditMessage)
	if !ok {
		t.Fatalf("expected updateEditMessage, got %T", updates.Updates[0])
	}
	if edit.Pts != 41 || edit.PtsCount != 1 {
		t.Fatalf("unexpected edit update pts: %+v", edit)
	}
	editMessage, ok := edit.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("edit message type = %T, want *tg.TLMessage", edit.Message)
	}
	if peer, ok := editMessage.PeerId.(*tg.TLPeerUser); !ok || peer.UserId != 1002 {
		t.Fatalf("edit response peer_id = %#v, want peerUser(1002)", editMessage.PeerId)
	}
	if editMessage.Id != 107 {
		t.Fatalf("edit response id = %d, want public user_message_id 107", editMessage.Id)
	}
	if repo.editInput.NewMessageText != "edited" || repo.editInput.ActorUserID != 1001 || repo.editInput.PeerSeq != 7 {
		t.Fatalf("unexpected edit input: %+v", repo.editInput)
	}
	if updatesClient.processed == nil || updatesClient.processed.OperationId != editMessageOperationID(7001, 1, 1001) {
		t.Fatalf("sender edit operation was not sent to userupdates: %+v", updatesClient.processed)
	}
	if updatesClient.processed.OperationId != "v2:msg:7001:edit:1:1001" {
		t.Fatalf("sender edit operation_id = %q, want v2 id", updatesClient.processed.OperationId)
	}
	var senderOp payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &senderOp); err != nil {
		t.Fatalf("decode sender edit payload: %v", err)
	}
	if senderOp.OperationKind != payload.OperationKindEditMessage || senderOp.MessageText != "edited" || senderOp.PeerSeq != 7 || senderOp.UserMessageID != 107 || senderOp.Date != 1_772_000_010 || senderOp.EditDate != 1_772_000_100 || senderOp.EditVersion != 1 {
		t.Fatalf("unexpected sender edit payload: %+v", senderOp)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != editMessageOperationID(7001, 1, 1002) {
		t.Fatalf("unexpected receiver edit operation: %+v", publisher.published)
	}
	if publisher.published.OperationID != "v2:msg:7001:edit:1:1002" {
		t.Fatalf("receiver edit operation_id = %q, want v2 id", publisher.published.OperationID)
	}
}

func TestMsgEditMessageOperationIDIncludesEditVersion(t *testing.T) {
	first := editMessageOperationID(7001, 1, 1001)
	second := editMessageOperationID(7001, 2, 1001)
	if first == second {
		t.Fatalf("edit operation ids must differ across edit versions: %q", first)
	}
	if first != "v2:msg:7001:edit:1:1001" || second != "v2:msg:7001:edit:2:1001" {
		t.Fatalf("unexpected edit operation ids: first=%q second=%q", first, second)
	}
}

func TestMsgEditMessageReceiverDispatchUsesBrokerDurableAck(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":42,"pts_count":1,"user_message_id":108}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 108}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      108,
				PeerSeq:            8,
				CanonicalMessageID: 7101,
			},
		},
		editResult: &repository.EditMessageResult{
			CanonicalMessageID: 7101,
			PeerSeq:            8,
			FromUserID:         1001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             1002,
			MessageText:        "edited by broker ack",
			MessageDate:        1_772_000_020,
			EditDate:           1_772_000_120,
			EditVersion:        2,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         editMessageOperationID(7101, 2, 1001),
			Status:              1,
			Pts:                 42,
			PtsCount:            1,
			CurrentPts:          42,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgEditMessage(editMessageRequest(1001, 1002, 9001, 108, "edited by broker ack"))
	if err != nil {
		t.Fatalf("MsgEditMessage() error = %v", err)
	}
	if _, ok := got.ToUpdates(); !ok {
		t.Fatalf("expected updates, got %s", got.ClazzName())
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != editMessageOperationID(7101, 2, 1002) {
		t.Fatalf("unexpected receiver edit operation: %+v", publisher.published)
	}
	if publisher.published.PeerID != 1001 || publisher.published.PayloadCodec != payload.PayloadCodecJSON {
		t.Fatalf("unexpected receiver edit metadata: %+v", publisher.published)
	}
	if len(updatesClient.processedList) != 1 || updatesClient.processWithEffects != nil {
		t.Fatalf("edit sender path should use requester sync only, processed=%d with_effects=%+v", len(updatesClient.processedList), updatesClient.processWithEffects)
	}

	publishErr := errors.New("broker unavailable")
	repo = &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 109}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      109,
				PeerSeq:            9,
				CanonicalMessageID: 7201,
			},
		},
		editResult: &repository.EditMessageResult{
			CanonicalMessageID: 7201,
			PeerSeq:            9,
			FromUserID:         1001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             1002,
			MessageText:        "edited fail",
			MessageDate:        1_772_000_021,
			EditDate:           1_772_000_121,
			EditVersion:        3,
		},
	}
	updatesClient = &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         editMessageOperationID(7201, 3, 1001),
			Status:              1,
			Pts:                 43,
			PtsCount:            1,
			CurrentPts:          43,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher = &fakeReceiverPublisher{publishErr: publishErr}
	core = New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err = core.MsgEditMessage(editMessageRequest(1001, 1002, 9001, 109, "edited fail"))
	if err == nil {
		t.Fatalf("MsgEditMessage() error = nil, got=%+v", got)
	}
	if !errors.Is(err, msgpb.ErrReceiverBackpressure) {
		t.Fatalf("MsgEditMessage() error = %v, want ErrReceiverBackpressure", err)
	}
	if !errors.Is(err, publishErr) {
		t.Fatalf("MsgEditMessage() error = %v, want upstream publish error", err)
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
}

func TestMsgEditMessageRejectsNonAuthor(t *testing.T) {
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 107}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      107,
				PeerSeq:            7,
				CanonicalMessageID: 7001,
			},
		},
		editErr: msgpb.ErrMessageAuthorRequired,
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: &fakeUserUpdatesClient{}, ReceiverPublisher: &fakeReceiverPublisher{}})

	_, err := core.MsgEditMessage(editMessageRequest(1001, 1002, 9001, 107, "edited"))
	if !errors.Is(err, msgpb.ErrMessageAuthorRequired) {
		t.Fatalf("MsgEditMessage error = %v, want %v", err, msgpb.ErrMessageAuthorRequired)
	}
}

type fakeMsgRepository struct {
	sendState           *repository.SendState
	canonical           *repository.CanonicalMessageResult
	batchResult         *repository.CanonicalBatchResult
	canonicalByPeerSeq  *repository.CanonicalMessage
	editResult          *repository.EditMessageResult
	editErr             error
	editInput           repository.EditCanonicalMessageInput
	history             []repository.HistoryMessage
	historyInput        repository.ListHistoryMessagesInput
	searchInput         repository.SearchHashTagMessagesInput
	searchNoMatch       bool
	historyCursorBounds repository.HistoryCursorBounds
	resolveHistoryInput struct {
		userID   int64
		peerType int32
		peerID   int64
		offsetID int32
		maxID    int32
		minID    int32
	}
	resolvedMessageID            *repository.ResolvedMessageID
	resolveByUserMessageID       map[resolveMessageKey]*repository.ResolvedMessageID
	resolveManyByUserMessageID   map[int64]*repository.ResolvedMessageID
	resolveDeleteByUserMessageID map[int64]*repository.ResolvedMessageID
	resolveInput                 repository.ResolvedMessageID
	resolveDeleteCalls           int
	userMessages                 map[int64]*repository.UserMessageBox
	getUserMessageInput          struct {
		UserID        int64
		UserMessageID int64
	}
	getUserMessageListInput struct {
		UserID int64
		IDs    []int64
	}
	forwardVisible             bool
	resolveForwardInput        repository.ForwardSourceLookup
	revalidateForwardSources   []repository.ForwardSourceIdentity
	revalidateForwardCallCount int
	markCanonicalErr           error
	markSenderErrs             []error
	batchCreateCalls           int
	markCanonicalCalls         int
	markSenderCalls            int
	markReceiverAckedCalls     int
	markCompletedCalls         int
	markRetryableCalls         int
}

type resolveMessageKey struct {
	userID        int64
	peerType      int32
	peerID        int64
	userMessageID int64
}

func (f *fakeMsgRepository) CreateOrLoadSendState(context.Context, repository.CreateSendStateInput) (*repository.SendState, error) {
	return f.sendState, nil
}

func (f *fakeMsgRepository) CreateOrGetByClientRandom(context.Context, repository.CreateCanonicalMessageInput) (*repository.CanonicalMessageResult, error) {
	return f.canonical, nil
}

func (f *fakeMsgRepository) CreateOrGetCanonicalBatchByClientRandom(context.Context, repository.CreateCanonicalBatchInput) (*repository.CanonicalBatchResult, error) {
	f.batchCreateCalls++
	return f.batchResult, nil
}

func (f *fakeMsgRepository) GetCanonicalMessageByPeerSeq(context.Context, int64, int32, int64, int64) (*repository.CanonicalMessage, error) {
	if f.canonicalByPeerSeq == nil {
		return nil, msgpb.ErrSendStateConflict
	}
	return f.canonicalByPeerSeq, nil
}

func (f *fakeMsgRepository) ListHistoryMessages(_ context.Context, in repository.ListHistoryMessagesInput) ([]repository.HistoryMessage, error) {
	f.historyInput = in
	if in.ResolvedCursorBounds.NoMatch {
		return []repository.HistoryMessage{}, nil
	}
	return f.history, nil
}

func (f *fakeMsgRepository) SearchHashTagMessages(_ context.Context, in repository.SearchHashTagMessagesInput) ([]repository.HistoryMessage, error) {
	f.searchInput = in
	if f.searchNoMatch && in.OffsetID > 0 {
		return []repository.HistoryMessage{}, nil
	}
	return f.history, nil
}

func (f *fakeMsgRepository) ResolveMessageID(_ context.Context, userID int64, peerType int32, peerID int64, userMessageID int64) (*repository.ResolvedMessageID, error) {
	f.resolveInput = repository.ResolvedMessageID{
		UserID:        userID,
		PeerType:      peerType,
		PeerID:        peerID,
		UserMessageID: userMessageID,
	}
	if f.resolveByUserMessageID != nil {
		return f.resolveByUserMessageID[resolveMessageKey{
			userID:        userID,
			peerType:      peerType,
			peerID:        peerID,
			userMessageID: userMessageID,
		}], nil
	}
	return f.resolvedMessageID, nil
}

func (f *fakeMsgRepository) ResolveMessageIDs(_ context.Context, userID int64, userMessageIDs []int64) ([]repository.ResolvedMessageID, error) {
	out := make([]repository.ResolvedMessageID, 0, len(userMessageIDs))
	for _, id := range userMessageIDs {
		f.resolveInput = repository.ResolvedMessageID{
			UserID:        userID,
			UserMessageID: id,
		}
		if f.resolveManyByUserMessageID != nil {
			if resolved := f.resolveManyByUserMessageID[id]; resolved != nil {
				out = append(out, *resolved)
			}
			continue
		}
		if f.resolvedMessageID != nil && f.resolvedMessageID.UserMessageID == id {
			out = append(out, *f.resolvedMessageID)
		}
	}
	return out, nil
}

func (f *fakeMsgRepository) ResolveMessageIDsForDelete(_ context.Context, userID int64, userMessageIDs []int64) ([]repository.ResolvedMessageID, error) {
	f.resolveDeleteCalls++
	out := make([]repository.ResolvedMessageID, 0, len(userMessageIDs))
	for _, id := range userMessageIDs {
		f.resolveInput = repository.ResolvedMessageID{
			UserID:        userID,
			UserMessageID: id,
		}
		if f.resolveDeleteByUserMessageID != nil {
			if resolved := f.resolveDeleteByUserMessageID[id]; resolved != nil {
				out = append(out, *resolved)
			}
			continue
		}
		if f.resolveManyByUserMessageID != nil {
			if resolved := f.resolveManyByUserMessageID[id]; resolved != nil {
				out = append(out, *resolved)
			}
			continue
		}
		if f.resolvedMessageID != nil && f.resolvedMessageID.UserMessageID == id {
			out = append(out, *f.resolvedMessageID)
		}
	}
	return out, nil
}

func (f *fakeMsgRepository) ResolveHistoryCursorIDs(_ context.Context, userID int64, peerType int32, peerID int64, offsetID int32, maxID int32, minID int32) (repository.HistoryCursorBounds, error) {
	f.resolveHistoryInput.userID = userID
	f.resolveHistoryInput.peerType = peerType
	f.resolveHistoryInput.peerID = peerID
	f.resolveHistoryInput.offsetID = offsetID
	f.resolveHistoryInput.maxID = maxID
	f.resolveHistoryInput.minID = minID
	return f.historyCursorBounds, nil
}

func (f *fakeMsgRepository) ResolvePeerSeqToUserMessageID(context.Context, int64, int32, int64, int64) (int64, error) {
	return 0, nil
}

func (f *fakeMsgRepository) GetUserMessage(_ context.Context, userID int64, userMessageID int64) (*repository.UserMessageBox, error) {
	f.getUserMessageInput.UserID = userID
	f.getUserMessageInput.UserMessageID = userMessageID
	if f.userMessages == nil {
		return nil, msgpb.ErrMsgIdInvalid
	}
	box := f.userMessages[userMessageID]
	if box == nil || box.UserID != userID {
		return nil, msgpb.ErrMsgIdInvalid
	}
	return box, nil
}

func (f *fakeMsgRepository) GetUserMessageList(_ context.Context, userID int64, ids []int64) ([]repository.UserMessageBox, error) {
	f.getUserMessageListInput.UserID = userID
	f.getUserMessageListInput.IDs = append([]int64(nil), ids...)
	out := make([]repository.UserMessageBox, 0, len(ids))
	for _, id := range ids {
		if id <= 0 {
			return nil, msgpb.ErrMsgIdInvalid
		}
		box := f.userMessages[id]
		if box == nil || box.UserID != userID {
			return nil, msgpb.ErrMsgIdInvalid
		}
		out = append(out, *box)
	}
	return out, nil
}

func (f *fakeMsgRepository) ResolveForwardSourceIdentity(_ context.Context, lookup repository.ForwardSourceLookup) (*repository.ForwardSourceIdentity, error) {
	f.resolveForwardInput = lookup
	if !f.forwardVisible {
		return nil, msgpb.ErrMsgIdInvalid
	}
	return &repository.ForwardSourceIdentity{
		UserID:             lookup.UserID,
		UserMessageID:      lookup.SourceUserMessageID,
		CanonicalMessageID: 7000 + lookup.SourceUserMessageID,
	}, nil
}

func (f *fakeMsgRepository) RevalidateForwardSources(_ context.Context, sources []repository.ForwardSourceIdentity) error {
	f.revalidateForwardCallCount++
	f.revalidateForwardSources = append([]repository.ForwardSourceIdentity(nil), sources...)
	if !f.forwardVisible {
		return msgpb.ErrMsgIdInvalid
	}
	return nil
}

func (f *fakeMsgRepository) EditCanonicalMessage(_ context.Context, in repository.EditCanonicalMessageInput) (*repository.EditMessageResult, error) {
	f.editInput = in
	if f.editErr != nil {
		return nil, f.editErr
	}
	return f.editResult, nil
}

func (f *fakeMsgRepository) MarkCanonicalCreated(context.Context, int64, int64, int64) error {
	f.markCanonicalCalls++
	return f.markCanonicalErr
}

func (f *fakeMsgRepository) MarkSenderCommitted(_ context.Context, _ repository.MarkSenderCommittedInput) error {
	f.markSenderCalls++
	if len(f.markSenderErrs) == 0 {
		return nil
	}
	err := f.markSenderErrs[0]
	f.markSenderErrs = f.markSenderErrs[1:]
	return err
}

func (f *fakeMsgRepository) MarkReceiverOpsAcked(context.Context, int64, int64) error {
	f.markReceiverAckedCalls++
	return nil
}

func (f *fakeMsgRepository) MarkCompleted(context.Context, int64) error {
	f.markCompletedCalls++
	return nil
}

func (f *fakeMsgRepository) MarkRetryableFailure(context.Context, repository.MarkRetryableFailureInput) error {
	f.markRetryableCalls++
	return nil
}

type fakeMsgChatClient struct {
	actions   []*chatpb.TLChatCheckMessageAction
	accesses  []*chatpb.TLChatCheckChatAccess
	memberIDs []int64
	err       error
}

func (f *fakeMsgChatClient) ChatCheckMessageAction(_ context.Context, in *chatpb.TLChatCheckMessageAction) (*chatpb.MessageActionCheckResult, error) {
	f.actions = append(f.actions, in)
	if f.err != nil {
		return nil, f.err
	}
	return chatpb.MakeTLMessageActionCheckResult(&chatpb.TLMessageActionCheckResult{
		SelfId: in.SelfId, ChatId: in.ChatId, Action: in.Action, MediaKind: in.MediaKind,
	}).ToMessageActionCheckResult(), nil
}

func (f *fakeMsgChatClient) ChatCheckChatAccess(_ context.Context, in *chatpb.TLChatCheckChatAccess) (*chatpb.ChatAccessCheckResult, error) {
	f.accesses = append(f.accesses, in)
	if f.err != nil {
		return nil, f.err
	}
	return chatpb.MakeTLChatAccessCheckResult(&chatpb.TLChatAccessCheckResult{
		SelfId: in.SelfId, ChatId: in.ChatId, AccessKind: in.AccessKind,
	}).ToChatAccessCheckResult(), nil
}

func (f *fakeMsgChatClient) ChatGetChatParticipantIdList(context.Context, *chatpb.TLChatGetChatParticipantIdList) (*chatpb.VectorLong, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &chatpb.VectorLong{Datas: append([]int64(nil), f.memberIDs...)}, nil
}

type fakeUserUpdatesClient struct {
	processed                *userupdates.TLUserOperation
	processedList            []*userupdates.TLUserOperation
	processedOperations      []*userupdates.TLUserOperation
	processResult            *userupdates.UserOperationResult
	processResults           []*userupdates.UserOperationResult
	processWithEffects       *userupdates.TLUserupdatesProcessUserOperationWithEffects
	processWithEffectsList   []*userupdates.TLUserupdatesProcessUserOperationWithEffects
	processWithEffectsResult *userupdates.UserOperationResult
	batchApplyCalls          int
	batchApplyLen            int
	processErr               error
	processWithEffectsErr    error
	receiverErr              error
	receiverPublished        []repository.ReceiverOperation
	getResult                *userupdates.UserOperationResult
	getOperationResultCalls  int
}

func (f *fakeUserUpdatesClient) nextProcessResult() *userupdates.UserOperationResult {
	if len(f.processResults) > 0 {
		result := f.processResults[0]
		f.processResults = f.processResults[1:]
		return result
	}
	return f.processResult
}

func (f *fakeUserUpdatesClient) UserupdatesProcessUserOperation(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	f.processed = in.Operation
	f.processedList = append(f.processedList, in.Operation)
	f.processedOperations = append(f.processedOperations, in.Operation)
	if f.processErr != nil {
		return nil, f.processErr
	}
	return upgradeLegacyTestOperationResult(f.nextProcessResult(), in.Operation, false), nil
}

func (f *fakeUserUpdatesClient) UserupdatesProcessUserOperationBatch(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperationBatch) (*userupdates.VectorUserOperationResult, error) {
	f.batchApplyCalls++
	f.batchApplyLen = len(in.Operations)
	f.processedOperations = append(f.processedOperations, in.Operations...)
	if f.processErr != nil {
		return nil, f.processErr
	}
	out := &userupdates.VectorUserOperationResult{Datas: make([]userupdates.UserOperationResultClazz, 0, len(in.Operations))}
	for _, operation := range in.Operations {
		result := f.nextProcessResult()
		if result == nil {
			result = userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{Status: 1})
		}
		out.Datas = append(out.Datas, upgradeLegacyTestOperationResult(result, operation, true))
	}
	return out, nil
}

func (f *fakeUserUpdatesClient) UserupdatesGetOperationResult(_ context.Context, _ *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error) {
	f.getOperationResultCalls++
	return upgradeLegacyTestOperationResult(f.getResult, f.processed, false), nil
}

func upgradeLegacyTestOperationResult(result *userupdates.UserOperationResult, operation *userupdates.TLUserOperation, forceFull bool) *userupdates.UserOperationResult {
	if result == nil || operation == nil || len(result.ResponsePayload) == 0 {
		return result
	}
	var header struct {
		SchemaVersion int `json:"schema_version"`
	}
	if err := json.Unmarshal(result.ResponsePayload, &header); err != nil || header.SchemaVersion != payload.OperationResponseSchemaVersion {
		return result
	}
	var legacy payload.OperationResponseV2
	if err := json.Unmarshal(result.ResponsePayload, &legacy); err != nil {
		return result
	}
	var op payload.MessageOperationV4
	if err := json.Unmarshal(operation.Payload, &op); err != nil || op.SchemaVersion != payload.MessageOperationSchemaVersionV4 {
		return result
	}
	reply := legacyReplyUpdatesForTest(legacy, op, forceFull)
	envelope, err := iface.EncodeObject(reply, 224)
	if err != nil {
		return result
	}
	responsePayload, err := json.Marshal(payload.OperationResponseV3{
		SchemaVersion:       payload.OperationResponseSchemaVersionV3,
		OperationID:         result.OperationId,
		Pts:                 legacy.Pts,
		PtsCount:            legacy.PtsCount,
		EventType:           legacy.EventType,
		UserMessageID:       legacy.UserMessageID,
		ClientRandomID:      op.MessageFact.ClientRandomID,
		ReplyEnvelope:       envelope,
		ReplyEnvelopeCodec:  payload.ReplyEnvelopeCodecTLBinary,
		ReplyEnvelopeSchema: payload.ReplyEnvelopeSchemaV1,
	})
	if err != nil {
		return result
	}
	clone := *result
	clone.ResponsePayload = responsePayload
	clone.ResponsePayloadHash = payload.HashBytes(responsePayload)
	return &clone
}

func mustDecodeMessageOperationV4(t *testing.T, body []byte) payload.MessageOperationV4 {
	t.Helper()
	var op payload.MessageOperationV4
	if err := json.Unmarshal(body, &op); err != nil {
		t.Fatalf("decode message operation v4: %v", err)
	}
	if op.SchemaVersion != payload.MessageOperationSchemaVersionV4 {
		t.Fatalf("message operation schema = %d, want v4; payload=%s", op.SchemaVersion, string(body))
	}
	return op
}

func mustDecodeMessageOperationBatchV1(t *testing.T, body []byte) payload.MessageOperationBatchV1 {
	t.Helper()
	var op payload.MessageOperationBatchV1
	if err := json.Unmarshal(body, &op); err != nil {
		t.Fatalf("decode message operation batch: %v", err)
	}
	if op.SchemaVersion != payload.MessageOperationSchemaVersionBatchV1 {
		t.Fatalf("message operation batch schema = %d, want batch v1; payload=%s", op.SchemaVersion, string(body))
	}
	return op
}

func legacyReplyUpdatesForTest(response payload.OperationResponseV2, op payload.MessageOperationV4, forceFull bool) *tg.Updates {
	messageID := int32(response.UserMessageID)
	pts := int32(response.Pts)
	fact := op.MessageFact
	if !forceFull && fact.ServiceAction == nil && fact.MediaRef == nil && fact.ForwardRef == nil &&
		len(fact.Entities) == 0 && fact.Attrs == nil && fact.ReplyToUserMessageID == 0 {
		return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
			Out:      true,
			Id:       messageID,
			Pts:      pts,
			PtsCount: response.PtsCount,
			Date:     fact.Date,
		}).ToUpdates()
	}
	message := tg.MessageClazz(tg.MakeTLMessage(&tg.TLMessage{
		Out:         true,
		Silent:      fact.Attrs != nil && fact.Attrs.Silent,
		Noforwards:  fact.Attrs != nil && fact.Attrs.Noforwards,
		InvertMedia: fact.Attrs != nil && fact.Attrs.InvertMedia,
		Id:          messageID,
		FromId:      tg.MakePeerUser(fact.SenderUserID),
		PeerId:      sentMessagePeerFromOptional(fact.PeerType, fact.PeerID),
		FwdFrom:     sentMessageForwardHeader(fact.ForwardRef),
		ReplyTo:     sentMessageReplyHeader(fact.ReplyToUserMessageID),
		Date:        fact.Date,
		Message:     fact.MessageText,
		Media:       sentMessageMedia(fact.MediaRef),
		Entities:    sentMessageEntities(fact.Entities),
		GroupedId:   sentMessageGroupedID(fact.Attrs),
		TtlPeriod:   sentMessageTTLPeriod(fact.MediaRef),
	}))
	if fact.ServiceAction != nil {
		action, _ := sentMessageServiceAction(fact.ServiceAction)
		message = tg.MakeTLMessageService(&tg.TLMessageService{
			Out:    true,
			Silent: fact.Attrs != nil && fact.Attrs.Silent,
			Id:     messageID,
			FromId: tg.MakePeerUser(fact.SenderUserID),
			PeerId: sentMessagePeerFromOptional(fact.PeerType, fact.PeerID),
			Date:   fact.Date,
			Action: action,
		})
	}
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: messageID, RandomId: fact.ClientRandomID}),
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{Message: message, Pts: pts, PtsCount: response.PtsCount}),
		},
		Users: []tg.UserClazz{},
		Chats: []tg.ChatClazz{},
		Date:  fact.Date,
	}).ToUpdates()
}

type fakeReceiverPublisher struct {
	calls       int
	published   repository.ReceiverOperation
	publisheds  []repository.ReceiverOperation
	publishErr  error
	publishFunc func(repository.ReceiverOperation) error
}

func (f *fakeReceiverPublisher) Publish(_ context.Context, op repository.ReceiverOperation) (repository.KafkaAck, error) {
	f.calls++
	f.published = op
	f.publisheds = append(f.publisheds, op)
	if f.publishFunc != nil {
		if err := f.publishFunc(op); err != nil {
			return repository.KafkaAck{}, err
		}
	}
	if f.publishErr != nil {
		return repository.KafkaAck{}, f.publishErr
	}
	return repository.KafkaAck{Topic: "memory", Partition: op.PartitionID, Offset: 0}, nil
}

func sendMessageRequest(senderID, peerID, authKeyID int64, text string) *msgpb.TLMsgSendMessage {
	return &msgpb.TLMsgSendMessage{
		UserId:    senderID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerID,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				RandomId: 77,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: text}),
			}),
		},
	}
}

func newFakeSendMessageRepositoryForChat(canonicalID, peerSeq int64) *fakeMsgRepository {
	return &fakeMsgRepository{
		forwardVisible: true,
		sendState:      &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: canonicalID,
			PeerSeq:            peerSeq,
			MessageDate:        1_772_000_000,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
}

func newBatchMsgCoreForTest(t *testing.T) (*MsgCore, *fakeMsgRepository, *fakeUserUpdatesClient) {
	t.Helper()
	senderPayload1 := []byte(`{"schema_version":2,"pts":11,"pts_count":1,"event_type":"new_message","user_message_id":41}`)
	senderPayload2 := []byte(`{"schema_version":2,"pts":12,"pts_count":1,"event_type":"new_message","user_message_id":42}`)
	fakeRepo := &fakeMsgRepository{
		forwardVisible: true,
		batchResult: &repository.CanonicalBatchResult{Items: []repository.CanonicalMessageResult{
			{SendStateID: 501, CanonicalMessageID: 9001, PeerSeq: 1, MessageDate: 1_772_000_000, RequestPayloadHash: payload.HashBytes([]byte("h1")), CreatedNew: true},
			{SendStateID: 502, CanonicalMessageID: 9002, PeerSeq: 2, MessageDate: 1_772_000_001, RequestPayloadHash: payload.HashBytes([]byte("h2")), CreatedNew: true},
		}},
	}
	fakeUpdates := &fakeUserUpdatesClient{processResults: []*userupdates.UserOperationResult{
		userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{Status: 1, Pts: 11, PtsCount: 1, ResponsePayload: senderPayload1, ResponsePayloadHash: payload.HashBytes(senderPayload1)}),
		userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{Status: 1, Pts: 12, PtsCount: 1, ResponsePayload: senderPayload2, ResponsePayloadHash: payload.HashBytes(senderPayload2)}),
	}}
	publisher := &fakeReceiverPublisher{publishFunc: func(op repository.ReceiverOperation) error {
		fakeUpdates.receiverPublished = append(fakeUpdates.receiverPublished, op)
		return fakeUpdates.receiverErr
	}}
	return New(context.Background(), &svc.ServiceContext{Repo: fakeRepo, UserUpdates: fakeUpdates, ReceiverPublisher: publisher}), fakeRepo, fakeUpdates
}

func buildBatchSendRequestForTest(senderID, peerID, authKeyID int64, randomIDs []int64) *msgpb.TLMsgSendMessage {
	out := &msgpb.TLMsgSendMessage{
		UserId:    senderID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerID,
		Message:   make([]msgpb.OutboxMessageClazz, 0, len(randomIDs)),
	}
	for i, randomID := range randomIDs {
		out.Message = append(out.Message, msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			RandomId: randomID,
			Message:  tg.MakeTLMessage(&tg.TLMessage{Message: string(rune('a' + i))}),
		}))
	}
	return out
}

func chatCreateServiceSendRequest(senderID, chatID, authKeyID, randomID int64) *msgpb.TLMsgSendMessage {
	return &msgpb.TLMsgSendMessage{
		UserId:    senderID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeChat,
		PeerId:    chatID,
		Message: []msgpb.OutboxMessageClazz{msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			RandomId: randomID,
			Message: tg.MakeTLMessageService(&tg.TLMessageService{
				Out:    true,
				FromId: tg.MakePeerUser(senderID),
				PeerId: tg.MakePeerChat(chatID),
				Date:   1778648899,
				Action: tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
					Title: "new chat",
					Users: []int64{1002, 1003},
				}),
			}),
		})},
	}
}

func chatParticipantsChangedTLFact(t *testing.T, chatID, actorUserID int64, participantIDs []int64) msgpb.UpdateFactClazz {
	t.Helper()
	participants := make([]payload.ChatParticipantFactV1, 0, len(participantIDs))
	for _, userID := range participantIDs {
		participants = append(participants, payload.ChatParticipantFactV1{UserID: userID, Role: "member", InviterUserID: actorUserID, Date: 1_772_000_000})
	}
	fact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, payload.ChatParticipantsChangedFactV1{
		SchemaVersion: 1,
		ChatID:        chatID,
		ActorUserID:   actorUserID,
		Version:       1,
		Participants:  participants,
	})
	if err != nil {
		t.Fatalf("wrap chat participants fact: %v", err)
	}
	body, err := json.Marshal(fact)
	if err != nil {
		t.Fatalf("marshal update fact envelope: %v", err)
	}
	return msgpb.MakeTLUpdateFact(&msgpb.TLUpdateFact{
		Kind:    fact.Kind,
		Payload: body,
	})
}

func buildMediaSendRequestForTest(senderID, peerID, authKeyID, randomID int64, caption string) *msgpb.TLMsgSendMessage {
	return &msgpb.TLMsgSendMessage{
		UserId:    senderID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerID,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				RandomId: randomID,
				Message: tg.MakeTLMessage(&tg.TLMessage{
					Message: caption,
					Media: tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
						Photo: tg.MakeTLPhoto(&tg.TLPhoto{
							Id:            333,
							AccessHash:    444,
							FileReference: []byte("1234567890123456789012345"),
							Date:          1_772_000_100,
							Sizes: []tg.PhotoSizeClazz{
								tg.MakeTLPhotoStrippedSize(&tg.TLPhotoStrippedSize{Type: "i", Bytes: []byte{0x01, 0x16, 0x28, 0xaa}}),
								tg.MakeTLPhotoSizeProgressive(&tg.TLPhotoSizeProgressive{Type: "y", W: 1280, H: 394, Sizes: []int32{100, 200, 300}}),
							},
							DcId: 2,
						}),
					}),
				}),
			}),
		},
	}
}

func buildForwardSendRequestForTest(senderID, peerID, authKeyID, randomID, sourceUserMessageID int64) *msgpb.TLMsgSendMessage {
	sourceID := int32(sourceUserMessageID)
	return &msgpb.TLMsgSendMessage{
		UserId:    senderID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerID,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				RandomId:        randomID,
				ForwardSourceId: &sourceID,
				Message: tg.MakeTLMessage(&tg.TLMessage{
					Message: "forwarded",
					FwdFrom: tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
						FromId: tg.MakePeerUser(peerID),
						Date:   1_772_000_301,
					}),
				}),
			}),
		},
	}
}

func newMsgReadCoreForTest(t *testing.T) (*MsgCore, *fakeMsgRepository) {
	t.Helper()
	repo := &fakeMsgRepository{forwardVisible: true}
	return New(context.Background(), &svc.ServiceContext{Repo: repo}), repo
}

func newSingleSendMsgCoreForTest(t *testing.T, userID int64, canonicalMessageID int64, userMessageID int64) *MsgCore {
	t.Helper()
	responsePayload := []byte(`{"schema_version":2,"pts":31,"pts_count":1,"event_type":"new_message","user_message_id":61}`)
	var response payload.OperationResponseV2
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Fatalf("decode test response payload: %v", err)
	}
	response.UserMessageID = userMessageID
	responsePayload, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("encode test response payload: %v", err)
	}
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		forwardVisible: true,
		sendState:      &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: canonicalMessageID,
			PeerSeq:            31,
			MessageDate:        1_772_000_200,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              userID,
			OperationId:         payload.SenderOperationID(canonicalMessageID, userID),
			Status:              1,
			Pts:                 31,
			PtsCount:            1,
			CurrentPts:          31,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	return New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: &fakeReceiverPublisher{},
	})
}

func firstSentUpdateMessage(t *testing.T, got *tg.Updates) *tg.TLMessage {
	t.Helper()
	updates, ok := got.Clazz.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Clazz)
	}
	var update *tg.TLUpdateNewMessage
	for _, item := range updates.Updates {
		if u, ok := item.(*tg.TLUpdateNewMessage); ok {
			update = u
			break
		}
	}
	if update == nil {
		t.Fatalf("updates = %#v, want updateNewMessage", updates.Updates)
	}
	message, ok := update.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", update.Message)
	}
	return message
}

func mustMarshalMsgMessageEventV3(t *testing.T, event payload.MessageEventV3) []byte {
	t.Helper()
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal MessageEventV3: %v", err)
	}
	return body
}

func testUserOperationResult(userID int64, operationID string, pts int64, responsePayload []byte) *userupdates.UserOperationResult {
	return userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
		UserId:              userID,
		OperationId:         operationID,
		Status:              1,
		Pts:                 pts,
		PtsCount:            1,
		CurrentPts:          pts,
		ResponsePayload:     responsePayload,
		ResponsePayloadHash: payload.HashBytes(responsePayload),
	})
}

func mustOperationResponseV3Payload(t *testing.T, operationID string, pts int64, ptsCount int32, userMessageID, clientRandomID int64, updates *tg.Updates) []byte {
	t.Helper()
	replyEnvelope, err := iface.EncodeObject(updates, 224)
	if err != nil {
		t.Fatalf("encode reply updates: %v", err)
	}
	body, err := json.Marshal(payload.OperationResponseV3{
		SchemaVersion:       payload.OperationResponseSchemaVersionV3,
		OperationID:         operationID,
		Pts:                 pts,
		PtsCount:            ptsCount,
		EventType:           payload.EventKindNewMessage,
		UserMessageID:       userMessageID,
		ClientRandomID:      clientRandomID,
		ReplyEnvelope:       replyEnvelope,
		ReplyEnvelopeCodec:  payload.ReplyEnvelopeCodecTLBinary,
		ReplyEnvelopeSchema: payload.ReplyEnvelopeSchemaV1,
	})
	if err != nil {
		t.Fatalf("marshal OperationResponseV3: %v", err)
	}
	return body
}

func mustMarshalOperationResponseV3(t *testing.T, response payload.OperationResponseV3) []byte {
	t.Helper()
	body, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("marshal OperationResponseV3: %v", err)
	}
	return body
}

func replyUpdatesForText(t *testing.T, userMessageID int64, randomID int64, pts int64, ptsCount int32) *tg.Updates {
	t.Helper()
	messageID := int64ToTestInt32(t, userMessageID)
	pts32 := int64ToTestInt32(t, pts)
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: messageID, RandomId: randomID}),
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
				Message: tg.MakeTLMessage(&tg.TLMessage{
					Out:     true,
					Id:      messageID,
					FromId:  tg.MakePeerUser(1001),
					PeerId:  tg.MakePeerUser(1002),
					Date:    1_772_000_000,
					Message: "hello",
				}),
				Pts:      pts32,
				PtsCount: ptsCount,
			}),
		},
		Users: []tg.UserClazz{tg.MakeTLUserEmpty(&tg.TLUserEmpty{Id: 1002})},
		Date:  1_772_000_000,
	}).ToUpdates()
}

func replyUpdatesForBatchEnvelope(t *testing.T, userMessageID int64, randomID int64, pts int64, ptsCount int32, text string, userID int64, chatID int64) *tg.Updates {
	t.Helper()
	messageID := int64ToTestInt32(t, userMessageID)
	pts32 := int64ToTestInt32(t, pts)
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: messageID, RandomId: randomID}),
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
				Message: tg.MakeTLMessage(&tg.TLMessage{
					Out:     true,
					Id:      messageID,
					FromId:  tg.MakePeerUser(100),
					PeerId:  tg.MakePeerUser(200),
					Date:    1_772_000_000,
					Message: text,
				}),
				Pts:      pts32,
				PtsCount: ptsCount,
			}),
		},
		Users: []tg.UserClazz{tg.MakeTLUserEmpty(&tg.TLUserEmpty{Id: userID})},
		Chats: []tg.ChatClazz{tg.MakeTLChatEmpty(&tg.TLChatEmpty{Id: chatID})},
		Date:  1_772_000_000,
	}).ToUpdates()
}

func replyShortSentMessageForText(t *testing.T, userMessageID int64, pts int64, ptsCount int32) *tg.Updates {
	t.Helper()
	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       int64ToTestInt32(t, userMessageID),
		Pts:      int64ToTestInt32(t, pts),
		PtsCount: ptsCount,
		Date:     1_772_000_000,
	}).ToUpdates()
}

func replyUpdatesForChatService(t *testing.T, userMessageID int64, randomID int64, pts int64, ptsCount int32) *tg.Updates {
	t.Helper()
	messageID := int64ToTestInt32(t, userMessageID)
	pts32 := int64ToTestInt32(t, pts)
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: messageID, RandomId: randomID}),
			tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
				Message: tg.MakeTLMessageService(&tg.TLMessageService{
					Out:    true,
					Id:     messageID,
					FromId: tg.MakePeerUser(1001),
					PeerId: tg.MakePeerChat(55),
					Date:   1_772_000_000,
					Action: tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
						Title: "new chat",
						Users: []int64{1002, 1003},
					}),
				}),
				Pts:      pts32,
				PtsCount: ptsCount,
			}),
		},
		Users: []tg.UserClazz{tg.MakeTLUserEmpty(&tg.TLUserEmpty{Id: 1001})},
		Chats: []tg.ChatClazz{tg.MakeTLChatEmpty(&tg.TLChatEmpty{Id: 55})},
		Date:  1_772_000_000,
	}).ToUpdates()
}

func int64ToTestInt32(t *testing.T, value int64) int32 {
	t.Helper()
	if value > int64(^uint32(0)>>1) || value < -int64(^uint32(0)>>1)-1 {
		t.Fatalf("value %d does not fit int32", value)
	}
	return int32(value)
}

func makeSequentialRandomIDsForTest(count int) []int64 {
	out := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		out = append(out, int64(i+1))
	}
	return out
}

func editMessageRequest(userID, peerID, authKeyID int64, peerSeq int32, text string) *msgpb.TLMsgEditMessage {
	return &msgpb.TLMsgEditMessage{
		UserId:    userID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerID,
		NewMessage: msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{Message: text}),
		}),
		DstMessage: tg.MakeTLMessageBox(&tg.TLMessageBox{
			MessageId: peerSeq,
		}),
	}
}

func mustHashBytes(t *testing.T, b []byte) []byte {
	t.Helper()
	return payload.HashBytes(b)
}
