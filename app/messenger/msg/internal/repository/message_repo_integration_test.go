//go:build integration

package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math"
	"os"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type testIDGenerator struct {
	next int64
	err  error
}

func (g *testIDGenerator) NextID(context.Context) (int64, error) {
	if g.err != nil {
		return 0, g.err
	}
	g.next++
	return g.next, nil
}

func TestMessageRepositoryCreateAndCommitSendState(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 10_000})

	senderID := base + 101
	receiverID := base + 102
	randomID := base + 103
	requestHash := payload.HashBytes([]byte("send request"))

	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      receiverID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState() error = %v", err)
	}
	if state.SendStateID == 0 || state.Status != SendStateStatusInitialized {
		t.Fatalf("unexpected initial state: %+v", state)
	}

	canonical, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             receiverID,
		ClientRandomID:     randomID,
		RequestPayloadHash: requestHash,
		MessageText:        "hello",
		MessageDate:        time.Now().Unix(),
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom() error = %v", err)
	}
	if canonical.SendStateID != state.SendStateID || canonical.CanonicalMessageID == 0 || canonical.PeerSeq != 1 || !canonical.CreatedNew {
		t.Fatalf("unexpected canonical result: %+v", canonical)
	}

	again, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             receiverID,
		ClientRandomID:     randomID,
		RequestPayloadHash: requestHash,
		MessageText:        "hello retry",
		MessageDate:        time.Now().Unix(),
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom() retry error = %v", err)
	}
	if again.CanonicalMessageID != canonical.CanonicalMessageID || again.PeerSeq != canonical.PeerSeq || again.MessageDate != canonical.MessageDate || again.CreatedNew {
		t.Fatalf("idempotent canonical mismatch: first=%+v again=%+v", canonical, again)
	}

	if err := repo.MarkCanonicalCreated(ctx, canonical.SendStateID, canonical.CanonicalMessageID, canonical.PeerSeq); err != nil {
		t.Fatalf("MarkCanonicalCreated() error = %v", err)
	}
	senderUpdate := []byte(`{"schema_version":1,"pts":7,"pts_count":1}`)
	senderUpdateHash := payload.HashBytes(senderUpdate)
	if err := repo.MarkSenderCommitted(ctx, MarkSenderCommittedInput{
		SendStateID:               canonical.SendStateID,
		SenderOperationID:         payload.SenderOperationID(canonical.CanonicalMessageID, senderID),
		SenderPTS:                 7,
		SenderPTSCount:            1,
		SenderUpdateSchemaVersion: payload.OperationResponseSchemaVersion,
		SenderUpdatePayload:       senderUpdate,
		SenderUpdatePayloadHash:   senderUpdateHash,
	}); err != nil {
		t.Fatalf("MarkSenderCommitted() error = %v", err)
	}
	if err := repo.MarkReceiverOpsAcked(ctx, canonical.SendStateID, 0); err != nil {
		t.Fatalf("MarkReceiverOpsAcked() error = %v", err)
	}
	if err := repo.MarkCompleted(ctx, canonical.SendStateID); err != nil {
		t.Fatalf("MarkCompleted() error = %v", err)
	}

	committed, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      receiverID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState() committed load error = %v", err)
	}
	if committed.Status != SendStateStatusCompleted ||
		committed.CanonicalMessageID != canonical.CanonicalMessageID ||
		committed.PeerSeq != canonical.PeerSeq ||
		committed.SenderPTS != 7 ||
		!bytes.Equal(committed.SenderUpdatePayloadHash, senderUpdateHash) {
		t.Fatalf("unexpected committed state: %+v", committed)
	}
}

func TestCanonicalMessageDatesUseUnixSeconds(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 10_000})
	date := int64(math.MaxInt32) + 10_000
	canonical := createCanonicalMessageForTest(t, ctx, repo, base+101, base+102, base+103, "int64 date", date)
	if canonical.MessageDate != date {
		t.Fatalf("MessageDate = %d, want %d", canonical.MessageDate, date)
	}

	got, err := repo.GetCanonicalMessageByPeerSeq(ctx, base+101, payload.PeerTypeUser, base+102, canonical.PeerSeq)
	if err != nil {
		t.Fatalf("GetCanonicalMessageByPeerSeq: %v", err)
	}
	if got.MessageDate != date {
		t.Fatalf("stored MessageDate = %d, want %d", got.MessageDate, date)
	}
}

func TestMessageRepositoryRandomIdConflict(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 20_000})

	in := CreateSendStateInput{
		SenderUserID:                base + 201,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      base + 202,
		ClientRandomID:              base + 203,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          payload.HashBytes([]byte("first")),
	}
	if _, err := repo.CreateOrLoadSendState(ctx, in); err != nil {
		t.Fatalf("CreateOrLoadSendState() error = %v", err)
	}
	in.RequestPayloadHash = payload.HashBytes([]byte("different"))
	_, err := repo.CreateOrLoadSendState(ctx, in)
	if !errors.Is(err, msg.ErrRandomIdConflict) {
		t.Fatalf("CreateOrLoadSendState() error = %v, want ErrRandomIdConflict", err)
	}
}

func TestMessageRepositoryRetryRecoversLegacyRequestHashAfterCanonicalCreated(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 30_000})

	senderID := base + 301
	peerID := base + 302
	randomID := base + 303
	legacyHash := payload.HashBytes([]byte("legacy request hash with clear_draft_before_date"))
	currentHash := payload.HashBytes([]byte("current request hash without clear_draft_before_date"))
	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      peerID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          legacyHash,
		MessageText:                 "hello",
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState(legacy) error = %v", err)
	}
	canonical, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		ClientRandomID:     randomID,
		RequestPayloadHash: legacyHash,
		MessageText:        "hello",
		MessageDate:        time.Now().Unix(),
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom(legacy) error = %v", err)
	}
	if err := repo.MarkCanonicalCreated(ctx, state.SendStateID, canonical.CanonicalMessageID, canonical.PeerSeq); err != nil {
		t.Fatalf("MarkCanonicalCreated() error = %v", err)
	}

	retryState, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      peerID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          currentHash,
		MessageText:                 "hello",
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState(retry) error = %v", err)
	}
	if retryState.SendStateID != state.SendStateID {
		t.Fatalf("retry send_state_id = %d, want %d", retryState.SendStateID, state.SendStateID)
	}
	retryCanonical, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		ClientRandomID:     randomID,
		RequestPayloadHash: currentHash,
		MessageText:        "hello",
		MessageDate:        time.Now().Unix(),
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom(retry) error = %v", err)
	}
	if retryCanonical.CanonicalMessageID != canonical.CanonicalMessageID || retryCanonical.CreatedNew {
		t.Fatalf("retry canonical = %+v, want existing %+v", retryCanonical, canonical)
	}

	_, err = repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      peerID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          payload.HashBytes([]byte("changed")),
		MessageText:                 "changed",
	})
	if !errors.Is(err, msg.ErrRandomIdConflict) {
		t.Fatalf("CreateOrLoadSendState(changed) error = %v, want ErrRandomIdConflict", err)
	}
}

func TestCreateOrGetCanonicalPersistsAttrsAndForwardRef(t *testing.T) {
	repo, ctx := newIntegrationRepository(t)
	base := time.Now().UnixNano() % 1_000_000_000
	senderID := base + 100
	peerID := base + 200
	randomID := base + 7001
	requestHash := payload.HashBytes([]byte("hash-media"))
	state := createSendStateForTest(t, ctx, repo, senderID, payload.PeerTypeUser, peerID, randomID, requestHash)

	mediaPayload := mustJSON(t, payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 333})
	attrsPayload := mustJSON(t, payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444})
	forwardPayload := mustJSON(t, payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 300, Date: 1700000001})

	got, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:               state.SendStateID,
		SenderUserID:              senderID,
		PeerType:                  payload.PeerTypeUser,
		PeerID:                    peerID,
		ClientRandomID:            randomID,
		RequestPayloadHash:        requestHash,
		MessageText:               "caption",
		MessageDate:               1700000000,
		MediaRefSchemaVersion:     payload.MediaRefSchemaVersionV1,
		MediaRefPayload:           mediaPayload,
		MessageAttrsSchemaVersion: payload.MessageAttrsSchemaVersionV1,
		MessageAttrsPayload:       attrsPayload,
		ForwardRefSchemaVersion:   payload.ForwardRefSchemaVersionV1,
		ForwardRefPayload:         forwardPayload,
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom() error = %v", err)
	}

	again, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		ClientRandomID:     randomID,
		RequestPayloadHash: requestHash,
		MessageText:        "caption",
		MessageDate:        1700000000,
	})
	if err != nil {
		t.Fatalf("retry CreateOrGetByClientRandom() error = %v", err)
	}
	if got.CanonicalMessageID != again.CanonicalMessageID {
		t.Fatalf("canonical id changed: first=%d retry=%d", got.CanonicalMessageID, again.CanonicalMessageID)
	}
	if got.MediaRefSchemaVersion != payload.MediaRefSchemaVersionV1 || !bytes.Equal(got.MediaRefPayload, mediaPayload) || !bytes.Equal(again.MediaRefPayload, mediaPayload) {
		t.Fatalf("media payload not hydrated: first=%+v retry=%+v", got, again)
	}
	if got.MessageAttrsSchemaVersion != payload.MessageAttrsSchemaVersionV1 || !bytes.Equal(got.MessageAttrsPayload, attrsPayload) || !bytes.Equal(again.MessageAttrsPayload, attrsPayload) {
		t.Fatalf("attrs payload not hydrated: first=%+v retry=%+v", got, again)
	}
	if got.ForwardRefSchemaVersion != payload.ForwardRefSchemaVersionV1 || !bytes.Equal(got.ForwardRefPayload, forwardPayload) || !bytes.Equal(again.ForwardRefPayload, forwardPayload) {
		t.Fatalf("forward payload not hydrated: first=%+v retry=%+v", got, again)
	}
}

func TestCreateOrGetByClientRandomRejectsPayloadChangedHashRetry(t *testing.T) {
	repo, ctx := newIntegrationRepository(t)
	base := time.Now().UnixNano() % 1_000_000_000
	senderID := base + 110
	peerID := base + 210
	randomID := base + 7101
	requestHash := payload.HashBytes([]byte("hash-original-payload"))
	state := createSendStateForTest(t, ctx, repo, senderID, payload.PeerTypeUser, peerID, randomID, requestHash)

	entitiesPayload := []byte(`[{"type":"bold","offset":0,"length":4}]`)
	mediaPayload := mustJSON(t, payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 333})
	attrsPayload := mustJSON(t, payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444})
	forwardPayload := mustJSON(t, payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 300, Date: 1700000001})

	if _, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:                  state.SendStateID,
		SenderUserID:                 senderID,
		PeerType:                     payload.PeerTypeUser,
		PeerID:                       peerID,
		ClientRandomID:               randomID,
		RequestPayloadHash:           requestHash,
		MessageText:                  "caption",
		MessageDate:                  1700000000,
		EntitiesPayloadSchemaVersion: 1,
		EntitiesPayload:              entitiesPayload,
		MediaRefSchemaVersion:        payload.MediaRefSchemaVersionV1,
		MediaRefPayload:              mediaPayload,
		MessageAttrsSchemaVersion:    payload.MessageAttrsSchemaVersionV1,
		MessageAttrsPayload:          attrsPayload,
		ForwardRefSchemaVersion:      payload.ForwardRefSchemaVersionV1,
		ForwardRefPayload:            forwardPayload,
	}); err != nil {
		t.Fatalf("CreateOrGetByClientRandom(original) error = %v", err)
	}

	_, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:                  state.SendStateID,
		SenderUserID:                 senderID,
		PeerType:                     payload.PeerTypeUser,
		PeerID:                       peerID,
		ClientRandomID:               randomID,
		RequestPayloadHash:           payload.HashBytes([]byte("hash-changed-payload")),
		MessageText:                  "caption",
		MessageDate:                  1700000000,
		EntitiesPayloadSchemaVersion: 1,
		EntitiesPayload:              []byte(`[{"type":"italic","offset":0,"length":4}]`),
		MediaRefSchemaVersion:        payload.MediaRefSchemaVersionV1,
		MediaRefPayload:              mustJSON(t, payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "document", ID: 333}),
		MessageAttrsSchemaVersion:    payload.MessageAttrsSchemaVersionV1,
		MessageAttrsPayload:          mustJSON(t, payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 555}),
		ForwardRefSchemaVersion:      payload.ForwardRefSchemaVersionV1,
		ForwardRefPayload:            mustJSON(t, payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 301, Date: 1700000001}),
	})
	if !errors.Is(err, msg.ErrRandomIdConflict) {
		t.Fatalf("CreateOrGetByClientRandom(changed payload retry) error = %v, want ErrRandomIdConflict", err)
	}
}

func TestCreateOrGetCanonicalBatchAllocatesOrderedPeerSeqs(t *testing.T) {
	repo, ctx := newIntegrationRepository(t)
	base := time.Now().UnixNano() % 1_000_000_000
	in := CreateCanonicalBatchInput{
		SenderUserID: base + 120,
		PeerType:     payload.PeerTypeUser,
		PeerID:       base + 220,
		Items: []CreateCanonicalBatchItem{
			{ClientRandomID: base + 7201, RequestPayloadHash: payload.HashBytes([]byte("batch-h1")), MessageText: "one", MessageDate: 1700000000},
			{ClientRandomID: base + 7202, RequestPayloadHash: payload.HashBytes([]byte("batch-h2")), MessageText: "two", MessageDate: 1700000001},
		},
	}

	got, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, in)
	if err != nil {
		t.Fatalf("CreateOrGetCanonicalBatchByClientRandom() error = %v", err)
	}
	if len(got.Items) != 2 {
		t.Fatalf("items len = %d, want 2: %+v", len(got.Items), got.Items)
	}
	if got.Items[1].PeerSeq != got.Items[0].PeerSeq+1 {
		t.Fatalf("peer seqs = %+v, want ordered consecutive", got.Items)
	}
	if !got.Items[0].CreatedNew || !got.Items[1].CreatedNew {
		t.Fatalf("created flags = %+v, want both new", got.Items)
	}
}

func TestCreateOrGetCanonicalBatchExactRetryReturnsExistingRows(t *testing.T) {
	repo, ctx := newIntegrationRepository(t)
	base := time.Now().UnixNano() % 1_000_000_000
	in := CreateCanonicalBatchInput{
		SenderUserID: base + 130,
		PeerType:     payload.PeerTypeUser,
		PeerID:       base + 230,
		Items: []CreateCanonicalBatchItem{
			{ClientRandomID: base + 7301, RequestPayloadHash: payload.HashBytes([]byte("batch-retry-h1")), MessageText: "one", MessageDate: 1700000100},
			{ClientRandomID: base + 7302, RequestPayloadHash: payload.HashBytes([]byte("batch-retry-h2")), MessageText: "two", MessageDate: 1700000101},
		},
	}
	first, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, in)
	if err != nil {
		t.Fatalf("first CreateOrGetCanonicalBatchByClientRandom() error = %v", err)
	}

	again, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, in)
	if err != nil {
		t.Fatalf("retry CreateOrGetCanonicalBatchByClientRandom() error = %v", err)
	}
	if len(again.Items) != 2 {
		t.Fatalf("retry items len = %d, want 2: %+v", len(again.Items), again.Items)
	}
	for i := range first.Items {
		if again.Items[i].CanonicalMessageID != first.Items[i].CanonicalMessageID ||
			again.Items[i].PeerSeq != first.Items[i].PeerSeq ||
			again.Items[i].CreatedNew {
			t.Fatalf("retry item %d = %+v, want existing %+v", i, again.Items[i], first.Items[i])
		}
	}
}

func TestCreateOrGetCanonicalBatchRetryReturnsCommittedSenderResult(t *testing.T) {
	repo, ctx := newIntegrationRepository(t)
	base := time.Now().UnixNano() % 1_000_000_000
	in := CreateCanonicalBatchInput{
		SenderUserID: base + 135,
		PeerType:     payload.PeerTypeUser,
		PeerID:       base + 235,
		Items: []CreateCanonicalBatchItem{
			{ClientRandomID: base + 7351, RequestPayloadHash: payload.HashBytes([]byte("batch-committed-h1")), MessageText: "one", MessageDate: 1700000150},
			{ClientRandomID: base + 7352, RequestPayloadHash: payload.HashBytes([]byte("batch-committed-h2")), MessageText: "two", MessageDate: 1700000151},
		},
	}
	first, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, in)
	if err != nil {
		t.Fatalf("first CreateOrGetCanonicalBatchByClientRandom() error = %v", err)
	}
	senderPayload := []byte(`{"schema_version":2,"pts":33,"pts_count":1,"event_type":"new_message","user_message_id":73}`)
	senderHash := payload.HashBytes(senderPayload)
	operationID := payload.SenderOperationID(first.Items[0].CanonicalMessageID, in.SenderUserID)
	if err := repo.MarkSenderCommitted(ctx, MarkSenderCommittedInput{
		SendStateID:               first.Items[0].SendStateID,
		SenderOperationID:         operationID,
		SenderPTS:                 33,
		SenderPTSCount:            1,
		SenderUpdateSchemaVersion: payload.OperationResponseSchemaVersion,
		SenderUpdatePayload:       senderPayload,
		SenderUpdatePayloadHash:   senderHash,
	}); err != nil {
		t.Fatalf("MarkSenderCommitted() error = %v", err)
	}

	again, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, in)
	if err != nil {
		t.Fatalf("retry CreateOrGetCanonicalBatchByClientRandom() error = %v", err)
	}
	if again.Items[0].SendStateStatus != SendStateStatusSenderCommitted ||
		again.Items[0].SenderOperationID != operationID ||
		again.Items[0].SenderPTS != 33 ||
		again.Items[0].SenderPTSCount != 1 ||
		again.Items[0].SenderUpdateSchemaVersion != payload.OperationResponseSchemaVersion ||
		!bytes.Equal(again.Items[0].SenderUpdatePayload, senderPayload) ||
		!bytes.Equal(again.Items[0].SenderUpdatePayloadHash, senderHash) {
		t.Fatalf("committed sender fields = %+v, want stored sender result", again.Items[0])
	}
	if again.Items[1].SendStateStatus >= SendStateStatusSenderCommitted {
		t.Fatalf("second item status = %d, want below sender committed", again.Items[1].SendStateStatus)
	}
}

func TestCreateOrGetCanonicalBatchCompletesMissingRowsOnRetry(t *testing.T) {
	repo, ctx := newIntegrationRepository(t)
	base := time.Now().UnixNano() % 1_000_000_000
	first := CreateCanonicalBatchInput{
		SenderUserID: base + 140,
		PeerType:     payload.PeerTypeUser,
		PeerID:       base + 240,
		Items: []CreateCanonicalBatchItem{
			{ClientRandomID: base + 7401, RequestPayloadHash: payload.HashBytes([]byte("batch-mixed-h1")), MessageText: "one", MessageDate: 1700000200},
		},
	}
	if _, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, first); err != nil {
		t.Fatalf("pre-create first item: %v", err)
	}
	retry := CreateCanonicalBatchInput{
		SenderUserID: first.SenderUserID,
		PeerType:     first.PeerType,
		PeerID:       first.PeerID,
		Items: []CreateCanonicalBatchItem{
			first.Items[0],
			{ClientRandomID: base + 7402, RequestPayloadHash: payload.HashBytes([]byte("batch-mixed-h2")), MessageText: "two", MessageDate: 1700000201},
		},
	}
	got, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, retry)
	if err != nil {
		t.Fatalf("completion retry error = %v", err)
	}
	if len(got.Items) != 2 || got.Items[1].PeerSeq != got.Items[0].PeerSeq+1 {
		t.Fatalf("completed peer seqs = %+v, want consecutive full batch", got.Items)
	}
	if got.Items[0].CreatedNew || !got.Items[1].CreatedNew {
		t.Fatalf("created flags = %+v, want existing first and new second", got.Items)
	}
}

func TestCreateOrGetCanonicalBatchRejectsExistingHashConflict(t *testing.T) {
	repo, ctx := newIntegrationRepository(t)
	base := time.Now().UnixNano() % 1_000_000_000
	original := CreateCanonicalBatchInput{
		SenderUserID: base + 150,
		PeerType:     payload.PeerTypeUser,
		PeerID:       base + 250,
		Items: []CreateCanonicalBatchItem{
			{ClientRandomID: base + 7501, RequestPayloadHash: payload.HashBytes([]byte("batch-conflict-h1")), MessageText: "one", MessageDate: 1700000300},
		},
	}
	if _, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, original); err != nil {
		t.Fatalf("pre-create original item: %v", err)
	}

	_, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, CreateCanonicalBatchInput{
		SenderUserID: original.SenderUserID,
		PeerType:     original.PeerType,
		PeerID:       original.PeerID,
		Items: []CreateCanonicalBatchItem{
			{ClientRandomID: original.Items[0].ClientRandomID, RequestPayloadHash: payload.HashBytes([]byte("batch-conflict-changed")), MessageText: "changed", MessageDate: 1700000300},
			{ClientRandomID: base + 7502, RequestPayloadHash: payload.HashBytes([]byte("batch-conflict-h2")), MessageText: "two", MessageDate: 1700000301},
		},
	})
	if !errors.Is(err, msg.ErrRandomIdConflict) {
		t.Fatalf("conflict batch error = %v, want ErrRandomIdConflict", err)
	}
}

func TestCreateOrGetCanonicalBatchConflictRollsBackNewRows(t *testing.T) {
	repo, ctx := newIntegrationRepository(t)
	base := time.Now().UnixNano() % 1_000_000_000
	original := CreateCanonicalBatchInput{
		SenderUserID: base + 160,
		PeerType:     payload.PeerTypeUser,
		PeerID:       base + 260,
		Items: []CreateCanonicalBatchItem{
			{ClientRandomID: base + 7601, RequestPayloadHash: payload.HashBytes([]byte("batch-rollback-h1")), MessageText: "one", MessageDate: 1700000400},
		},
	}
	first, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, original)
	if err != nil {
		t.Fatalf("pre-create original item: %v", err)
	}
	conflictedNew := CreateCanonicalBatchItem{ClientRandomID: base + 7602, RequestPayloadHash: payload.HashBytes([]byte("batch-rollback-h2")), MessageText: "two", MessageDate: 1700000401}
	_, err = repo.CreateOrGetCanonicalBatchByClientRandom(ctx, CreateCanonicalBatchInput{
		SenderUserID: original.SenderUserID,
		PeerType:     original.PeerType,
		PeerID:       original.PeerID,
		Items: []CreateCanonicalBatchItem{
			conflictedNew,
			{ClientRandomID: original.Items[0].ClientRandomID, RequestPayloadHash: payload.HashBytes([]byte("batch-rollback-changed")), MessageText: "changed", MessageDate: 1700000400},
		},
	})
	if !errors.Is(err, msg.ErrRandomIdConflict) {
		t.Fatalf("conflict batch error = %v, want ErrRandomIdConflict", err)
	}

	retry, err := repo.CreateOrGetCanonicalBatchByClientRandom(ctx, CreateCanonicalBatchInput{
		SenderUserID: original.SenderUserID,
		PeerType:     original.PeerType,
		PeerID:       original.PeerID,
		Items:        []CreateCanonicalBatchItem{original.Items[0], conflictedNew},
	})
	if err != nil {
		t.Fatalf("retry after rollback error = %v", err)
	}
	if len(retry.Items) != 2 || retry.Items[0].CanonicalMessageID != first.Items[0].CanonicalMessageID || retry.Items[1].PeerSeq != retry.Items[0].PeerSeq+1 {
		t.Fatalf("retry after rollback = %+v, first=%+v", retry.Items, first.Items)
	}
}

func TestMessageRepositoryListHistoryMessages(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 40_000})

	senderID := base + 401
	receiverID := base + 402
	firstDate := time.Now().Unix()
	secondDate := firstDate + 1

	first := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+403, "first", firstDate)
	second := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+404, "second", secondDate)

	history, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:   senderID,
		PeerType: payload.PeerTypeUser,
		PeerID:   receiverID,
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() error = %v", err)
	}
	if len(history) != 2 {
		t.Fatalf("ListHistoryMessages() len = %d, want 2: %+v", len(history), history)
	}
	if history[0].CanonicalMessageID != second.CanonicalMessageID ||
		history[0].PeerSeq != second.PeerSeq ||
		history[0].FromUserID != senderID ||
		history[0].MessageText != "second" ||
		history[0].MessageDate != secondDate {
		t.Fatalf("unexpected newest history row: %+v, canonical: %+v", history[0], second)
	}
	if history[1].CanonicalMessageID != first.CanonicalMessageID ||
		history[1].PeerSeq != first.PeerSeq ||
		history[1].FromUserID != senderID ||
		history[1].MessageText != "first" ||
		history[1].MessageDate != firstDate {
		t.Fatalf("unexpected older history row: %+v, canonical: %+v", history[1], first)
	}

	beforeSecond, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:   senderID,
		PeerType: payload.PeerTypeUser,
		PeerID:   receiverID,
		OffsetID: int32(second.PeerSeq),
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() offset error = %v", err)
	}
	if len(beforeSecond) != 1 || beforeSecond[0].CanonicalMessageID != first.CanonicalMessageID {
		t.Fatalf("ListHistoryMessages() offset = %+v, want first canonical id %d", beforeSecond, first.CanonicalMessageID)
	}
}

func TestMessageRepositoryListHistoryMessagesUsesOffsetIDPositionBeforeFilters(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 45_000})

	senderID := base + 451
	receiverID := base + 452
	now := time.Now().Unix()
	one := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+453, "one", now)
	two := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+454, "two", now+1)
	three := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+455, "three", now+2)

	newerThanOne, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:    senderID,
		PeerType:  payload.PeerTypeUser,
		PeerID:    receiverID,
		OffsetID:  int32(one.PeerSeq),
		AddOffset: -3,
		Limit:     3,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() add_offset error = %v", err)
	}
	if len(newerThanOne) != 3 ||
		newerThanOne[0].CanonicalMessageID != three.CanonicalMessageID ||
		newerThanOne[1].CanonicalMessageID != two.CanonicalMessageID ||
		newerThanOne[2].CanonicalMessageID != one.CanonicalMessageID {
		t.Fatalf("ListHistoryMessages() add_offset = %+v, want three/two/one", newerThanOne)
	}

	filteredAfterSlice, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:   senderID,
		PeerType: payload.PeerTypeUser,
		PeerID:   receiverID,
		Limit:    2,
		MaxID:    int32(three.PeerSeq),
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() max_id error = %v", err)
	}
	if len(filteredAfterSlice) != 1 || filteredAfterSlice[0].CanonicalMessageID != two.CanonicalMessageID {
		t.Fatalf("ListHistoryMessages() max_id after slice = %+v, want only two", filteredAfterSlice)
	}
}

func TestMessageRepositoryListHistoryMessagesMissingOffsetIDStillReturnsVisibleHistory(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 48_000})

	senderID := base + 481
	receiverID := base + 482
	now := time.Now().Unix()
	first := createCanonicalMessageOnlyForTest(t, ctx, repo, senderID, receiverID, base+483, "first", now)
	second := createCanonicalMessageOnlyForTest(t, ctx, repo, senderID, receiverID, base+484, "second", now+1)
	insertUserMessageViewWithUserMessageIDForTest(t, ctx, db, senderID, payload.PeerTypeUser, receiverID, 394, first, senderID, true)
	insertUserMessageViewWithUserMessageIDForTest(t, ctx, db, senderID, payload.PeerTypeUser, receiverID, 395, second, senderID, true)

	history, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:    senderID,
		PeerType:  payload.PeerTypeUser,
		PeerID:    receiverID,
		OffsetID:  1,
		AddOffset: -25,
		Limit:     50,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() missing offset error = %v", err)
	}
	if len(history) != 2 ||
		history[0].CanonicalMessageID != second.CanonicalMessageID ||
		history[1].CanonicalMessageID != first.CanonicalMessageID {
		t.Fatalf("ListHistoryMessages() missing offset = %+v, want second/first", history)
	}
}

func TestMessageRepositoryListHistoryMessagesUsesViewerScopedViews(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 50_000})

	peerID := base + 501
	viewerID := base + 502
	peerSelf := createCanonicalMessageForTest(t, ctx, repo, peerID, peerID, base+503, "peer self", time.Now().Unix())
	direct := createCanonicalMessageForTest(t, ctx, repo, peerID, viewerID, base+504, "direct to viewer", time.Now().Unix()+1)
	insertUserMessageViewForTest(t, ctx, db, viewerID, payload.PeerTypeUser, peerID, direct, peerID, false)

	history, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:   viewerID,
		PeerType: payload.PeerTypeUser,
		PeerID:   peerID,
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() error = %v", err)
	}
	if len(history) != 1 {
		t.Fatalf("ListHistoryMessages() len = %d, want 1: %+v", len(history), history)
	}
	if history[0].CanonicalMessageID != direct.CanonicalMessageID || history[0].MessageText != "direct to viewer" {
		t.Fatalf("ListHistoryMessages() = %+v, want direct message %+v", history[0], direct)
	}
	if history[0].FromUserID != peerID {
		t.Fatalf("ListHistoryMessages() from_user_id = %d, want canonical sender %d", history[0].FromUserID, peerID)
	}
	if history[0].Outgoing {
		t.Fatalf("ListHistoryMessages() outgoing = true, want false for receiver view: %+v", history[0])
	}
	if history[0].CanonicalMessageID == peerSelf.CanonicalMessageID || history[0].MessageText == "peer self" {
		t.Fatalf("viewer history leaked peer self message: %+v", history[0])
	}
}

func TestMessageRepositoryListHistoryMessagesCarriesReplyToPeerSeqFromViewPayload(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 55_000})

	senderID := base + 551
	receiverID := base + 552
	now := time.Now().Unix()
	target := createCanonicalMessageForTest(t, ctx, repo, receiverID, senderID, base+553, "target", now)
	reply := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+554, "reply", now+1)
	insertUserMessageViewWithPayloadForTest(t, ctx, db, receiverID, payload.PeerTypeUser, senderID, reply, 11, senderID, false, payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: reply.CanonicalMessageID,
		MessageID:          11,
		PeerType:           payload.PeerTypeUser,
		PeerID:             senderID,
		FromUserID:         senderID,
		ToUserID:           receiverID,
		Date:               int32(now + 1),
		Out:                false,
		MessageText:        "reply",
		ReplyToPeerSeq:     target.PeerSeq,
	})

	history, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:   receiverID,
		PeerType: payload.PeerTypeUser,
		PeerID:   senderID,
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() error = %v", err)
	}
	if len(history) == 0 {
		t.Fatal("ListHistoryMessages() returned no messages")
	}
	if history[0].CanonicalMessageID != reply.CanonicalMessageID || history[0].ReplyToPeerSeq != target.PeerSeq {
		t.Fatalf("history[0] = %+v, want reply_to_peer_seq %d", history[0], target.PeerSeq)
	}
}

func TestMessageRepositoryListHistoryMessagesUsesCanonicalSender(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 60_000})

	senderID := base + 601
	receiverID := base + 602
	canonical := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+603, "canonical sender", time.Now().Unix())
	updateUserMessageViewFromForTest(t, ctx, db, senderID, payload.PeerTypeUser, receiverID, canonical.PeerSeq, receiverID)

	history, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:   senderID,
		PeerType: payload.PeerTypeUser,
		PeerID:   receiverID,
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() error = %v", err)
	}
	if len(history) != 1 {
		t.Fatalf("ListHistoryMessages() len = %d, want 1: %+v", len(history), history)
	}
	if history[0].FromUserID != senderID {
		t.Fatalf("ListHistoryMessages() from_user_id = %d, want canonical sender %d", history[0].FromUserID, senderID)
	}
}

func TestMessageRepositoryCreateCanonicalUsesConversationScopedPeerSeq(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 65_000})

	userA := base + 651
	userB := base + 652
	now := time.Now().Unix()
	first := createCanonicalMessageForTest(t, ctx, repo, userA, userB, base+653, "a to b", now)
	second := createCanonicalMessageForTest(t, ctx, repo, userB, userA, base+654, "b to a", now+1)
	insertUserMessageViewForTest(t, ctx, db, userA, payload.PeerTypeUser, userB, second, userB, false)

	if second.PeerSeq <= first.PeerSeq {
		t.Fatalf("reverse message peer_seq = %d, want greater than first direction peer_seq %d", second.PeerSeq, first.PeerSeq)
	}

	history, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:   userA,
		PeerType: payload.PeerTypeUser,
		PeerID:   userB,
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() error = %v", err)
	}
	if len(history) != 2 {
		t.Fatalf("ListHistoryMessages() len = %d, want 2: %+v", len(history), history)
	}
	if history[0].CanonicalMessageID != second.CanonicalMessageID || history[0].PeerSeq != second.PeerSeq || history[0].Outgoing {
		t.Fatalf("newest history row = %+v, want incoming reverse message %+v", history[0], second)
	}
	if history[1].CanonicalMessageID != first.CanonicalMessageID || history[1].PeerSeq != first.PeerSeq || !history[1].Outgoing {
		t.Fatalf("older history row = %+v, want outgoing first message %+v", history[1], first)
	}
}

func TestMessageRepositoryListHistoryMessagesOffsetsByViewerTimeline(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 70_000})

	viewerID := base + 701
	peerID := base + 702
	now := time.Now().Unix()
	olderSent := createCanonicalMessageForTest(t, ctx, repo, viewerID, peerID, base+703, "older sent high id", now)
	updateUserMessageViewSeqForTest(t, ctx, db, viewerID, payload.PeerTypeUser, peerID, olderSent.PeerSeq, 100)
	newerIncoming := createCanonicalMessageForTest(t, ctx, repo, peerID, viewerID, base+704, "newer incoming low id", now+10)
	insertUserMessageViewWithSeqForTest(t, ctx, db, viewerID, payload.PeerTypeUser, peerID, newerIncoming, 50, peerID, false)

	history, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		UserID:   viewerID,
		PeerType: payload.PeerTypeUser,
		PeerID:   peerID,
		OffsetID: 50,
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() error = %v", err)
	}
	if len(history) != 1 || history[0].CanonicalMessageID != olderSent.CanonicalMessageID {
		t.Fatalf("ListHistoryMessages() = %+v, want older sent message after newer offset", history)
	}
}

func updateUserMessageViewSeqForTest(
	t *testing.T,
	ctx context.Context,
	db *sqlx.DB,
	userID int64,
	peerType int32,
	peerID int64,
	oldPeerSeq int64,
	newPeerSeq int64,
) {
	t.Helper()
	result, err := db.Exec(ctx, `
UPDATE user_message_views
SET peer_seq = ?
WHERE user_id = ? AND peer_type = ? AND peer_id = ? AND peer_seq = ?`,
		newPeerSeq,
		userID,
		peerType,
		peerID,
		oldPeerSeq,
	)
	if err != nil {
		t.Fatalf("update user_message_views peer_seq: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		t.Fatalf("update user_message_views peer_seq RowsAffected: %v", err)
	}
	if rows != 1 {
		t.Fatalf("update user_message_views peer_seq rows = %d, want 1", rows)
	}
}

func updateUserMessageViewFromForTest(
	t *testing.T,
	ctx context.Context,
	db *sqlx.DB,
	userID int64,
	peerType int32,
	peerID int64,
	peerSeq int64,
	fromUserID int64,
) {
	t.Helper()
	result, err := db.Exec(ctx, `
UPDATE user_message_views
SET from_user_id = ?
WHERE user_id = ? AND peer_type = ? AND peer_id = ? AND peer_seq = ?`,
		fromUserID,
		userID,
		peerType,
		peerID,
		peerSeq,
	)
	if err != nil {
		t.Fatalf("update user_message_views from_user_id: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		t.Fatalf("update user_message_views RowsAffected: %v", err)
	}
	if rows != 1 {
		t.Fatalf("update user_message_views rows = %d, want 1", rows)
	}
}

func insertUserMessageViewWithSeqForTest(
	t *testing.T,
	ctx context.Context,
	db *sqlx.DB,
	userID int64,
	peerType int32,
	peerID int64,
	canonical *CanonicalMessageResult,
	peerSeq int64,
	fromUserID int64,
	outgoing bool,
) {
	t.Helper()
	if db == nil {
		t.Fatal("test db is nil")
	}
	if canonical == nil {
		t.Fatal("canonical is nil")
	}
	_, err := db.Exec(ctx, `
INSERT INTO user_message_views
	(user_id, peer_type, peer_id, user_message_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, date, view_schema_version, view_payload)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userID,
		peerType,
		peerID,
		peerSeq,
		peerSeq,
		canonical.CanonicalMessageID,
		fromUserID,
		outgoing,
		MessageKindText,
		MessageStatusLive,
		0,
		canonical.MessageDate,
		1,
		nil,
	)
	if err != nil {
		t.Fatalf("insert user_message_views user_id=%d peer_id=%d peer_seq=%d canonical=%d: %v", userID, peerID, peerSeq, canonical.CanonicalMessageID, err)
	}
}

func insertUserMessageViewWithPayloadForTest(
	t *testing.T,
	ctx context.Context,
	db *sqlx.DB,
	userID int64,
	peerType int32,
	peerID int64,
	canonical *CanonicalMessageResult,
	peerSeq int64,
	fromUserID int64,
	outgoing bool,
	event payload.MessageEventV1,
) {
	t.Helper()
	viewPayload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal view payload: %v", err)
	}
	_, err = db.Exec(ctx, `
INSERT INTO user_message_views
	(user_id, peer_type, peer_id, user_message_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, date, view_schema_version, view_payload)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userID,
		peerType,
		peerID,
		peerSeq,
		peerSeq,
		canonical.CanonicalMessageID,
		fromUserID,
		outgoing,
		MessageKindText,
		MessageStatusLive,
		0,
		canonical.MessageDate,
		payload.MessageEventSchemaVersion,
		viewPayload,
	)
	if err != nil {
		t.Fatalf("insert user_message_views payload user_id=%d peer_id=%d peer_seq=%d canonical=%d: %v", userID, peerID, peerSeq, canonical.CanonicalMessageID, err)
	}
}

func TestMessageRepositoryMarkSenderCommittedRejectsOutOfRangePTS(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 30_000})

	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                base + 301,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      base + 302,
		ClientRandomID:              base + 303,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          payload.HashBytes([]byte("send request")),
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState() error = %v", err)
	}

	err = repo.MarkSenderCommitted(ctx, MarkSenderCommittedInput{
		SendStateID:               state.SendStateID,
		SenderOperationID:         payload.SenderOperationID(base+304, base+301),
		SenderPTS:                 1 << 40,
		SenderPTSCount:            1,
		SenderUpdateSchemaVersion: payload.OperationResponseSchemaVersion,
		SenderUpdatePayload:       []byte(`{}`),
		SenderUpdatePayloadHash:   payload.HashBytes([]byte(`{}`)),
	})
	if !errors.Is(err, msg.ErrSenderSyncFailed) {
		t.Fatalf("MarkSenderCommitted() error = %v, want ErrSenderSyncFailed", err)
	}
}

func createCanonicalMessageForTest(
	t *testing.T,
	ctx context.Context,
	repo *Repository,
	senderID int64,
	receiverID int64,
	randomID int64,
	text string,
	date int64,
) *CanonicalMessageResult {
	t.Helper()
	requestHash := payload.HashBytes([]byte(text))
	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      receiverID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState(%q) error = %v", text, err)
	}
	canonical, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             receiverID,
		ClientRandomID:     randomID,
		RequestPayloadHash: requestHash,
		MessageText:        text,
		MessageDate:        date,
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom(%q) error = %v", text, err)
	}
	insertUserMessageViewForTest(t, ctx, repo.db, senderID, payload.PeerTypeUser, receiverID, canonical, senderID, true)
	return canonical
}

func createCanonicalMessageOnlyForTest(
	t *testing.T,
	ctx context.Context,
	repo *Repository,
	senderID int64,
	receiverID int64,
	randomID int64,
	text string,
	date int64,
) *CanonicalMessageResult {
	t.Helper()
	requestHash := payload.HashBytes([]byte(text))
	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      receiverID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState(%q) error = %v", text, err)
	}
	canonical, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             receiverID,
		ClientRandomID:     randomID,
		RequestPayloadHash: requestHash,
		MessageText:        text,
		MessageDate:        date,
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom(%q) error = %v", text, err)
	}
	return canonical
}

func newIntegrationRepository(t *testing.T) (*Repository, context.Context) {
	t.Helper()
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	return NewForTest(db, &testIDGenerator{next: base + 90_000}), ctx
}

func createSendStateForTest(
	t *testing.T,
	ctx context.Context,
	repo *Repository,
	senderID int64,
	peerType int32,
	peerID int64,
	randomID int64,
	requestHash []byte,
) *SendState {
	t.Helper()
	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    peerType,
		PeerID:                      peerID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState() error = %v", err)
	}
	return state
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal JSON: %v", err)
	}
	return data
}

func insertUserMessageViewForTest(
	t *testing.T,
	ctx context.Context,
	db *sqlx.DB,
	userID int64,
	peerType int32,
	peerID int64,
	canonical *CanonicalMessageResult,
	fromUserID int64,
	outgoing bool,
) {
	t.Helper()
	if db == nil {
		t.Fatal("test db is nil")
	}
	if canonical == nil {
		t.Fatal("canonical is nil")
	}
	_, err := db.Exec(ctx, `
INSERT INTO user_message_views
	(user_id, peer_type, peer_id, user_message_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, date, view_schema_version, view_payload)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userID,
		peerType,
		peerID,
		canonical.PeerSeq,
		canonical.PeerSeq,
		canonical.CanonicalMessageID,
		fromUserID,
		outgoing,
		MessageKindText,
		MessageStatusLive,
		0,
		canonical.MessageDate,
		1,
		nil,
	)
	if err != nil {
		t.Fatalf("insert user_message_views user_id=%d peer_id=%d canonical=%d: %v", userID, peerID, canonical.CanonicalMessageID, err)
	}
}

func insertUserMessageViewWithUserMessageIDForTest(
	t *testing.T,
	ctx context.Context,
	db *sqlx.DB,
	userID int64,
	peerType int32,
	peerID int64,
	userMessageID int64,
	canonical *CanonicalMessageResult,
	fromUserID int64,
	outgoing bool,
) {
	t.Helper()
	if db == nil {
		t.Fatal("test db is nil")
	}
	if canonical == nil {
		t.Fatal("canonical is nil")
	}
	_, err := db.Exec(ctx, `
INSERT INTO user_message_views
	(user_id, peer_type, peer_id, user_message_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, date, view_schema_version, view_payload)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userID,
		peerType,
		peerID,
		userMessageID,
		canonical.PeerSeq,
		canonical.CanonicalMessageID,
		fromUserID,
		outgoing,
		MessageKindText,
		MessageStatusLive,
		0,
		canonical.MessageDate,
		1,
		nil,
	)
	if err != nil {
		t.Fatalf("insert user_message_views user_id=%d peer_id=%d user_message_id=%d canonical=%d: %v", userID, peerID, userMessageID, canonical.CanonicalMessageID, err)
	}
}

func openIntegrationDB(t *testing.T) *sqlx.DB {
	t.Helper()
	if testing.Short() {
		t.Skip("integration test skipped in short mode")
	}
	dsn := os.Getenv("TEAMGRAM_TEST_MYSQL_DSN")
	explicit := dsn != ""
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/teamgooo?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	}
	db, err := sqlx.Open(&sqlx.Config{DSN: dsn})
	if err != nil {
		if explicit {
			t.Fatalf("open mysql: %v", err)
		}
		t.Skipf("mysql unavailable: %v", err)
	}
	if _, err := db.Exec(context.Background(), "SELECT 1"); err != nil {
		if explicit {
			t.Fatalf("ping mysql: %v", err)
		}
		t.Skipf("mysql unavailable: %v", err)
	}
	return db
}
