package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestProcessUserOperationMapsTLToRepository(t *testing.T) {
	operationPayload := []byte(`{"schema_version":1,"operation_kind":"send_message"}`)
	operationHash := payload.HashBytes(operationPayload)
	responsePayload := []byte(`{"schema_version":1,"pts":12,"pts_count":1}`)
	responseHash := payload.HashBytes(responsePayload)

	repo := &fakeUserUpdatesRepository{
		applyResult: &repository.ApplyUserOperationResult{
			UserID:          1001,
			OperationID:     "v1:msg:2001:sender:1001:out",
			Pts:             12,
			PtsCount:        1,
			ResponsePayload: responsePayload,
			ResponseHash:    responseHash,
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})
	authKeyIDExclude := int64(9001)

	got, err := core.UserupdatesProcessUserOperation(&userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               1001,
			BucketId:             77,
			PartitionId:          13,
			OperationId:          "v1:msg:2001:sender:1001:out",
			OpType:               repository.OpTypeSendMessage,
			PeerType:             payload.PeerTypeUser,
			PeerId:               1002,
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         repository.PayloadCodecJSON,
			PayloadHash:          operationHash,
			Payload:              operationPayload,
			AuthKeyIdExclude:     &authKeyIDExclude,
		}),
	})
	if err != nil {
		t.Fatalf("ProcessUserOperation returned error: %v", err)
	}
	if got.Pts != 12 || got.PtsCount != 1 || got.CurrentPts != 12 {
		t.Fatalf("unexpected pts result: pts=%d pts_count=%d current_pts=%d", got.Pts, got.PtsCount, got.CurrentPts)
	}
	if got.ResponseSchemaVersion == nil || *got.ResponseSchemaVersion != payload.OperationResponseSchemaVersion {
		t.Fatalf("unexpected response schema version: %v", got.ResponseSchemaVersion)
	}
	if string(got.ResponsePayload) != string(responsePayload) {
		t.Fatalf("unexpected response payload: %s", string(got.ResponsePayload))
	}
	if !bytes.Equal(got.ResponsePayloadHash, responseHash) {
		t.Fatalf("unexpected response hash: %x", got.ResponsePayloadHash)
	}
	if repo.applyInput.UserID != 1001 ||
		repo.applyInput.OperationID != "v1:msg:2001:sender:1001:out" ||
		!bytes.Equal(repo.applyInput.PayloadHash, operationHash) ||
		repo.applyInput.BucketID != 77 ||
		repo.applyInput.PartitionID != 13 ||
		repo.applyInput.AuthKeyIDExclude == nil ||
		*repo.applyInput.AuthKeyIDExclude != authKeyIDExclude {
		t.Fatalf("unexpected repository input: %+v", repo.applyInput)
	}
}

func TestGetOperationResultRejectsMismatchedPayloadHash(t *testing.T) {
	goodPayload := []byte(`{"good":true}`)
	badPayload := []byte(`{"good":false}`)
	repo := &fakeUserUpdatesRepository{
		operationResult: &repository.OperationResult{
			UserID:      1001,
			OperationID: "v1:msg:2001:receiver:1001:in",
			Status:      repository.OperationResultStatusCompleted,
			Pts:         8,
			PtsCount:    1,
			PayloadHash: payload.HashBytes(goodPayload),
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetOperationResult(&userupdates.TLUserupdatesGetOperationResult{
		UserId:      1001,
		OperationId: "v1:msg:2001:receiver:1001:in",
		PayloadHash: payload.HashBytes(badPayload),
	})
	if !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		t.Fatalf("expected ErrOperationPayloadConflict, got %v", err)
	}
}

func TestGetDifferenceBuildsVisibleMessageFromEventPayload(t *testing.T) {
	eventPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 2001,
		MessageID:          9,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1002,
		ToUserID:           1001,
		Date:               1_772_000_000,
		Out:                false,
		MessageText:        "hello from event payload",
	})
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 18},
			Events: []repository.UserEvent{
				{
					UserID:             1001,
					Pts:                18,
					PtsCount:           1,
					OperationID:        "v1:msg:2001:receiver:1001:in",
					EventType:          repository.EventTypeNewMessage,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1002,
					CanonicalMessageID: 2001,
					PeerSeq:            9,
					ActorUserID:        1002,
					EventSchemaVersion: payload.MessageEventSchemaVersion,
					EventCodec:         repository.PayloadCodecJSON,
					EventPayload:       eventPayload,
					EventPayloadHash:   payload.HashBytes(eventPayload),
				},
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:        1001,
		AuthKeyId:     9001,
		Pts:           17,
		PtsTotalLimit: int32Ptr(10),
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	diff, ok := got.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference, got %s", got.ClazzName())
	}
	if diff.State == nil || diff.State.Pts != 18 {
		t.Fatalf("unexpected state: %#v", diff.State)
	}
	if len(diff.NewMessages) != 1 {
		t.Fatalf("expected one new message, got %d", len(diff.NewMessages))
	}
	message, ok := diff.NewMessages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("expected TLMessage, got %T", diff.NewMessages[0])
	}
	if message.Id != 9 || message.Message != "hello from event payload" || message.Out {
		t.Fatalf("unexpected message projection: %+v", message)
	}
	if len(diff.OtherUpdates) != 1 {
		t.Fatalf("expected one update, got %d", len(diff.OtherUpdates))
	}
	update, ok := diff.OtherUpdates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("expected TLUpdateNewMessage, got %T", diff.OtherUpdates[0])
	}
	if update.Pts != 18 || update.PtsCount != 1 {
		t.Fatalf("unexpected update pts: %+v", update)
	}
}

func TestGetDifferenceBuildsReadHistoryOutboxUpdate(t *testing.T) {
	eventPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.OperationKindReadHistory,
		MessageID:     66,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1571766987,
		Date:          1_778_029_828,
		Out:           true,
	})
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1571766986, Pts: 157},
			Events: []repository.UserEvent{
				{
					UserID:             1571766986,
					Pts:                157,
					PtsCount:           1,
					OperationID:        "read-outbox",
					EventType:          repository.EventTypeReadHistory,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1571766987,
					PeerSeq:            66,
					ActorUserID:        1571766987,
					EventSchemaVersion: payload.MessageEventSchemaVersion,
					EventCodec:         repository.PayloadCodecJSON,
					EventPayload:       eventPayload,
					EventPayloadHash:   payload.HashBytes(eventPayload),
				},
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:        1571766986,
		AuthKeyId:     9002,
		Pts:           156,
		PtsTotalLimit: int32Ptr(10),
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	diff, ok := got.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference, got %s", got.ClazzName())
	}
	if len(diff.NewMessages) != 0 || len(diff.OtherUpdates) != 1 {
		t.Fatalf("unexpected difference lens: new=%d updates=%d", len(diff.NewMessages), len(diff.OtherUpdates))
	}
	update, ok := diff.OtherUpdates[0].(*tg.TLUpdateReadHistoryOutbox)
	if !ok {
		t.Fatalf("expected TLUpdateReadHistoryOutbox, got %T", diff.OtherUpdates[0])
	}
	peer, ok := update.Peer.(*tg.TLPeerUser)
	if !ok || peer.UserId != 1571766987 {
		t.Fatalf("unexpected peer: %+v ok=%v", update.Peer, ok)
	}
	if update.MaxId != 66 || update.Pts != 157 || update.PtsCount != 1 {
		t.Fatalf("unexpected read outbox update: %+v", update)
	}
}

func TestGetMessageViewsByPeerSeqsBuildsMessagesFromViews(t *testing.T) {
	eventPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 2001,
		MessageID:          7,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1_772_000_000,
		Out:                true,
		MessageText:        "dialog top",
	})
	peer := repository.MessageViewPeerSeq{PeerType: payload.PeerTypeUser, PeerID: 1002, PeerSeq: 7}
	repo := &fakeUserUpdatesRepository{
		messageViews: map[repository.MessageViewPeerSeq]repository.MessageView{
			peer: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				PeerSeq:            7,
				CanonicalMessageID: 2001,
				FromUserID:         1001,
				Outgoing:           true,
				MessageStatus:      repository.MessageStatusLive,
				ViewSchemaVersion:  payload.MessageEventSchemaVersion,
				ViewPayload:        eventPayload,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetMessageViewsByPeerSeqs(&userupdates.TLUserupdatesGetMessageViewsByPeerSeqs{
		UserId: 1001,
		Peers: []userupdates.MessageViewPeerSeqClazz{
			userupdates.MakeTLMessageViewPeerSeq(&userupdates.TLMessageViewPeerSeq{PeerType: payload.PeerTypeUser, PeerId: 1002, PeerSeq: 7}),
		},
	})
	if err != nil {
		t.Fatalf("GetMessageViewsByPeerSeqs returned error: %v", err)
	}
	if repo.messageViewUserID != 1001 || len(repo.messageViewPeers) != 1 || repo.messageViewPeers[0] != peer {
		t.Fatalf("unexpected repository message view request: user_id=%d peers=%+v", repo.messageViewUserID, repo.messageViewPeers)
	}
	if got == nil || len(got.Messages) != 1 {
		t.Fatalf("expected one message, got %+v", got)
	}
	message, ok := got.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("expected TLMessage, got %T", got.Messages[0])
	}
	if message.Id != 7 || message.Message != "dialog top" || !message.Out {
		t.Fatalf("unexpected message view projection: %+v", message)
	}
}

func TestGetStateReturnsRepositoryState(t *testing.T) {
	repo := &fakeUserUpdatesRepository{
		state: &repository.UserState{UserID: 1001, Pts: 55},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetState(&userupdates.TLUserupdatesGetState{UserId: 1001, AuthKeyId: 9001})
	if err != nil {
		t.Fatalf("GetState returned error: %v", err)
	}
	if got.Pts != 55 {
		t.Fatalf("unexpected pts: %d", got.Pts)
	}
}

func TestGetStatePassesPermAuthKeyToRepository(t *testing.T) {
	repo := &fakeUserUpdatesRepository{
		state: &repository.UserState{UserID: 1001, Pts: 55},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetState(&userupdates.TLUserupdatesGetState{UserId: 1001, AuthKeyId: 9001})
	if err != nil {
		t.Fatalf("GetState returned error: %v", err)
	}
	if repo.stateUserID != 1001 || repo.statePermAuthKeyID != 9001 {
		t.Fatalf("unexpected repository state cursor input: user_id=%d perm_auth_key_id=%d", repo.stateUserID, repo.statePermAuthKeyID)
	}
}

func TestGetDifferenceCarriesNilDateAsPtsOnlyRequest(t *testing.T) {
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 18},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:        1001,
		AuthKeyId:     9001,
		Pts:           17,
		PtsTotalLimit: int32Ptr(10),
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	if repo.differenceInput.UserID != 1001 || repo.differenceInput.PermAuthKeyID != 9001 || repo.differenceInput.Pts != 17 || repo.differenceInput.Limit != 10 {
		t.Fatalf("unexpected repository difference input: %+v", repo.differenceInput)
	}
	if repo.differenceInput.Date != nil {
		t.Fatalf("expected nil date, got %v", *repo.differenceInput.Date)
	}
}

func TestGetDifferenceCarriesDateToRepository(t *testing.T) {
	date := int64(1_772_000_000)
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 18},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:    1001,
		AuthKeyId: 9001,
		Pts:       17,
		Date:      &date,
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	if repo.differenceInput.Date == nil || *repo.differenceInput.Date != date {
		t.Fatalf("expected date %d, got %v", date, repo.differenceInput.Date)
	}
}

func TestDifferenceMapsDialogAuthSeqEvents(t *testing.T) {
	pinned := false
	folderID := int32(2)
	ttl := int32(86400)
	tests := []struct {
		name      string
		eventKind string
		event     payload.DialogEventV1
		wantType  any
	}{
		{
			name:      "draft saved",
			eventKind: payload.DialogEventDraftSaved,
			wantType:  &tg.TLUpdateDraftMessage{},
		},
		{
			name:      "dialog pin",
			eventKind: payload.DialogEventPinToggled,
			event:     payload.DialogEventV1{Pinned: &pinned, FolderID: &folderID},
			wantType:  &tg.TLUpdateDialogPinned{},
		},
		{
			name:      "pinned order",
			eventKind: payload.DialogEventPinnedDialogsReordered,
			event:     payload.DialogEventV1{FolderID: &folderID},
			wantType:  &tg.TLUpdatePinnedDialogs{},
		},
		{
			name:      "filter updated",
			eventKind: payload.DialogEventFilterUpdated,
			wantType:  &tg.TLUpdateDialogFilter{},
		},
		{
			name:      "filter order",
			eventKind: payload.DialogEventFiltersOrderUpdated,
			wantType:  &tg.TLUpdateDialogFilterOrder{},
		},
		{
			name:      "wallpaper",
			eventKind: payload.DialogEventWallpaperChanged,
			wantType:  &tg.TLUpdatePeerWallpaper{},
		},
		{
			name:      "private ttl",
			eventKind: payload.DialogEventPrivatePeerHistoryTTL,
			event:     payload.DialogEventV1{TTLPeriod: &ttl},
			wantType:  &tg.TLUpdatePeerHistoryTTL{},
		},
		{
			name:      "saved dialog pinned",
			eventKind: payload.DialogEventSavedDialogPinned,
			event:     payload.DialogEventV1{Pinned: &pinned},
			wantType:  &tg.TLUpdateSavedDialogPinned{},
		},
		{
			name:      "pinned saved order",
			eventKind: payload.DialogEventPinnedSavedDialogsChanged,
			wantType:  &tg.TLUpdatePinnedSavedDialogs{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dialogEvent := tt.event
			dialogEvent.SchemaVersion = payload.DialogEventSchemaVersion
			dialogEvent.EventKind = tt.eventKind
			dialogEvent.PublicUpdateType = tt.eventKind
			dialogEvent.PeerType = payload.PeerTypeUser
			dialogEvent.PeerID = 1002
			eventPayload := mustMarshalDialogEvent(t, dialogEvent)
			got, err := differenceToTL(&repository.GetDifferenceResult{
				State: repository.UserState{UserID: 1001, Pts: 18, Seq: 7, Date: 1_772_000_001},
				AuthSeqEvents: []repository.AuthSeqEvent{
					{
						UserID:             1001,
						Seq:                7,
						Date:               1_772_000_001,
						OperationID:        "v1:dialog:auth",
						PublicUpdateType:   tt.eventKind,
						PeerType:           payload.PeerTypeUser,
						PeerID:             1002,
						EventSchemaVersion: payload.DialogEventSchemaVersion,
						EventCodec:         repository.PayloadCodecJSON,
						EventPayload:       eventPayload,
						EventPayloadHash:   payload.HashBytes(eventPayload),
					},
				},
			})
			if err != nil {
				t.Fatalf("differenceToTL returned error: %v", err)
			}
			diff, ok := got.ToUserDifference()
			if !ok {
				t.Fatalf("expected userDifference, got %s", got.ClazzName())
			}
			if diff.State == nil || diff.State.Pts != 18 || diff.State.Seq != 7 || diff.State.Date != 1_772_000_001 {
				t.Fatalf("unexpected state: %#v", diff.State)
			}
			if len(diff.OtherUpdates) != 1 {
				t.Fatalf("expected one auth-seq update, got %d", len(diff.OtherUpdates))
			}
			switch tt.wantType.(type) {
			case *tg.TLUpdateDraftMessage:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateDraftMessage); !ok {
					t.Fatalf("expected TLUpdateDraftMessage, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdateDialogPinned:
				update, ok := diff.OtherUpdates[0].(*tg.TLUpdateDialogPinned)
				if !ok {
					t.Fatalf("expected TLUpdateDialogPinned, got %T", diff.OtherUpdates[0])
				}
				if update.Pinned != pinned || update.FolderId == nil || *update.FolderId != folderID {
					t.Fatalf("unexpected pinned update: %+v", update)
				}
			case *tg.TLUpdatePinnedDialogs:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdatePinnedDialogs); !ok {
					t.Fatalf("expected TLUpdatePinnedDialogs, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdateDialogFilter:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateDialogFilter); !ok {
					t.Fatalf("expected TLUpdateDialogFilter, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdateDialogFilterOrder:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateDialogFilterOrder); !ok {
					t.Fatalf("expected TLUpdateDialogFilterOrder, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdatePeerWallpaper:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdatePeerWallpaper); !ok {
					t.Fatalf("expected TLUpdatePeerWallpaper, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdatePeerHistoryTTL:
				update, ok := diff.OtherUpdates[0].(*tg.TLUpdatePeerHistoryTTL)
				if !ok {
					t.Fatalf("expected TLUpdatePeerHistoryTTL, got %T", diff.OtherUpdates[0])
				}
				if update.TtlPeriod == nil || *update.TtlPeriod != ttl {
					t.Fatalf("unexpected ttl update: %+v", update)
				}
			case *tg.TLUpdateSavedDialogPinned:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateSavedDialogPinned); !ok {
					t.Fatalf("expected TLUpdateSavedDialogPinned, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdatePinnedSavedDialogs:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdatePinnedSavedDialogs); !ok {
					t.Fatalf("expected TLUpdatePinnedSavedDialogs, got %T", diff.OtherUpdates[0])
				}
			default:
				t.Fatalf("unsupported test type %T", tt.wantType)
			}
		})
	}
}

func TestDifferenceMapsFolderPeersAsPTSEvent(t *testing.T) {
	folderID := int32(3)
	eventPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.DialogEventFolderPeersChanged,
	})
	got, err := differenceToTL(&repository.GetDifferenceResult{
		State: repository.UserState{UserID: 1001, Pts: 21},
		Events: []repository.UserEvent{
			{
				UserID:             1001,
				Pts:                21,
				PtsCount:           1,
				OperationID:        "v1:dialog:folder",
				EventType:          repository.EventTypeDialogPublicUpdate,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				EventSchemaVersion: payload.MessageEventSchemaVersion,
				EventCodec:         repository.PayloadCodecJSON,
				EventPayload:       eventPayload,
				EventPayloadHash:   payload.HashBytes(eventPayload),
			},
		},
	})
	if err != nil {
		t.Fatalf("differenceToTL returned error: %v", err)
	}
	diff, ok := got.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference, got %s", got.ClazzName())
	}
	if len(diff.NewMessages) != 0 || len(diff.OtherUpdates) != 1 {
		t.Fatalf("unexpected difference lens: new=%d updates=%d", len(diff.NewMessages), len(diff.OtherUpdates))
	}
	update, ok := diff.OtherUpdates[0].(*tg.TLUpdateFolderPeers)
	if !ok {
		t.Fatalf("expected TLUpdateFolderPeers, got %T", diff.OtherUpdates[0])
	}
	if update.Pts != 21 || update.PtsCount != 1 {
		t.Fatalf("unexpected pts update: %+v", update)
	}
	if len(update.FolderPeers) != 1 {
		t.Fatalf("expected one folder peer, got %d", len(update.FolderPeers))
	}
	fp := update.FolderPeers[0]
	if fp.FolderId != 0 && fp.FolderId != folderID {
		t.Fatalf("unexpected folder peer: %+v", fp)
	}
}

func TestDifferenceRejectsUnknownDialogAuthSeqEvent(t *testing.T) {
	eventPayload := mustMarshalDialogEvent(t, payload.DialogEventV1{
		SchemaVersion:    payload.DialogEventSchemaVersion,
		EventKind:        "dialog.unknown",
		PublicUpdateType: "dialog.unknown",
		PeerType:         payload.PeerTypeUser,
		PeerID:           1002,
	})
	_, err := differenceToTL(&repository.GetDifferenceResult{
		State: repository.UserState{UserID: 1001, Pts: 18, Seq: 7, Date: 1_772_000_001},
		AuthSeqEvents: []repository.AuthSeqEvent{
			{
				UserID:             1001,
				Seq:                7,
				Date:               1_772_000_001,
				OperationID:        "v1:dialog:auth",
				PublicUpdateType:   "dialog.unknown",
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				EventSchemaVersion: payload.DialogEventSchemaVersion,
				EventCodec:         repository.PayloadCodecJSON,
				EventPayload:       eventPayload,
				EventPayloadHash:   payload.HashBytes(eventPayload),
			},
		},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("expected ErrUserupdatesStorage, got %v", err)
	}
}

func TestGetOutboxReadDateRoutesToRepository(t *testing.T) {
	repo := &fakeUserUpdatesRepository{outboxReadDate: 123456}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetOutboxReadDate(&userupdates.TLUserupdatesGetOutboxReadDate{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		MsgId:    9,
	})
	if err != nil {
		t.Fatalf("UserupdatesGetOutboxReadDate error = %v", err)
	}
	if got.Date != 123456 {
		t.Fatalf("date = %d, want 123456", got.Date)
	}
	if repo.outboxReadDateInput.UserID != 1001 || repo.outboxReadDateInput.PeerType != payload.PeerTypeUser ||
		repo.outboxReadDateInput.PeerID != 1002 || repo.outboxReadDateInput.MsgID != 9 {
		t.Fatalf("repository input = %+v", repo.outboxReadDateInput)
	}
}

type fakeUserUpdatesRepository struct {
	applyInput          repository.ApplyUserOperationInput
	applyResult         *repository.ApplyUserOperationResult
	operationResult     *repository.OperationResult
	stateUserID         int64
	statePermAuthKeyID  int64
	state               *repository.UserState
	differenceInput     repository.GetDifferenceInput
	difference          *repository.GetDifferenceResult
	dialogListUserID    int64
	dialogListCursor    repository.DialogProjectionCursor
	dialogListLimit     int32
	dialogProjections   []repository.DialogProjection
	dialogPeerUserID    int64
	dialogPeers         []repository.DialogProjectionPeer
	dialogPeerMap       map[repository.DialogProjectionPeer]repository.DialogProjection
	messageViewUserID   int64
	messageViewPeers    []repository.MessageViewPeerSeq
	messageViews        map[repository.MessageViewPeerSeq]repository.MessageView
	dialogCountUserID   int64
	dialogCount         int32
	outboxReadDate      int32
	outboxReadDateInput repository.OutboxReadDateInput
}

func (f *fakeUserUpdatesRepository) ApplyUserOperation(_ context.Context, in repository.ApplyUserOperationInput) (*repository.ApplyUserOperationResult, error) {
	f.applyInput = in
	return f.applyResult, nil
}

func (f *fakeUserUpdatesRepository) GetOperationResult(_ context.Context, _ int64, _ string) (*repository.OperationResult, error) {
	return f.operationResult, nil
}

func (f *fakeUserUpdatesRepository) GetState(_ context.Context, userID int64, permAuthKeyID int64) (*repository.UserState, error) {
	f.stateUserID = userID
	f.statePermAuthKeyID = permAuthKeyID
	return f.state, nil
}

func (f *fakeUserUpdatesRepository) GetDifference(_ context.Context, in repository.GetDifferenceInput) (*repository.GetDifferenceResult, error) {
	f.differenceInput = in
	return f.difference, nil
}

func (f *fakeUserUpdatesRepository) AppendDialogAuthSeqSideEffect(context.Context, repository.DialogSideEffectAppendInput) (*repository.AuthSeqAppendResult, error) {
	return nil, nil
}

func (f *fakeUserUpdatesRepository) AppendDialogPtsSideEffect(context.Context, repository.DialogSideEffectAppendInput) (*repository.PtsAppendResult, error) {
	return nil, nil
}

func (f *fakeUserUpdatesRepository) ListDialogProjections(_ context.Context, userID int64, cursor repository.DialogProjectionCursor, limit int32) ([]repository.DialogProjection, error) {
	f.dialogListUserID = userID
	f.dialogListCursor = cursor
	f.dialogListLimit = limit
	return f.dialogProjections, nil
}

func (f *fakeUserUpdatesRepository) GetDialogProjectionsByPeers(_ context.Context, userID int64, peers []repository.DialogProjectionPeer) (map[repository.DialogProjectionPeer]repository.DialogProjection, error) {
	f.dialogPeerUserID = userID
	f.dialogPeers = peers
	return f.dialogPeerMap, nil
}

func (f *fakeUserUpdatesRepository) GetMessageViewsByPeerSeqs(_ context.Context, userID int64, peers []repository.MessageViewPeerSeq) (map[repository.MessageViewPeerSeq]repository.MessageView, error) {
	f.messageViewUserID = userID
	f.messageViewPeers = peers
	return f.messageViews, nil
}

func (f *fakeUserUpdatesRepository) CountVisibleDialogs(_ context.Context, userID int64) (int32, error) {
	f.dialogCountUserID = userID
	return f.dialogCount, nil
}

func (f *fakeUserUpdatesRepository) GetOutboxReadDate(_ context.Context, in repository.OutboxReadDateInput) (int32, error) {
	f.outboxReadDateInput = in
	return f.outboxReadDate, nil
}

func mustMarshalMessageEvent(t *testing.T, event payload.MessageEventV1) []byte {
	t.Helper()
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal message event: %v", err)
	}
	return b
}

func mustMarshalDialogEvent(t *testing.T, event payload.DialogEventV1) []byte {
	t.Helper()
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal dialog event: %v", err)
	}
	return b
}

func int32Ptr(v int32) *int32 {
	return &v
}
