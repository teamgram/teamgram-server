//go:build integration

package core

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	msgrepo "github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/svc"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	userupdatestestkit "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/testkit"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestSendMessageV2SingleChatAcceptance(t *testing.T) {
	ctx := context.Background()
	db := openAcceptanceDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	idgen := &acceptanceIDGenerator{next: base + 100_000}

	msgRepo := msgrepo.NewForTest(db, idgen)
	updatesKit := userupdatestestkit.New(db, idgen, "local-userupdates")
	senderID := base + 1001
	receiverID := base + 1002
	randomID := base + 1003
	claimUserPartitions(t, ctx, updatesKit, senderID, receiverID)

	publisher := &msgrepo.InMemoryReceiverOperationPublisher{OnPublish: updatesKit.ProcessReceiverOperation}
	msgCore := New(ctx, &svc.ServiceContext{
		Repo:              msgRepo,
		UserUpdates:       updatesKit,
		ReceiverPublisher: publisher,
	})

	first, err := msgCore.MsgSendMessageV2(acceptanceSendRequest(senderID, receiverID, 9001, randomID, "hello acceptance"))
	if err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	firstShort := mustShortSent(t, first)
	if firstShort.Pts != 1 || firstShort.Id != 1 {
		t.Fatalf("unexpected first result: %+v", firstShort)
	}

	assertDifferenceMessage(t, ctx, updatesKit, senderID, true, "hello acceptance")
	assertDifferenceMessage(t, ctx, updatesKit, receiverID, false, "hello acceptance")

	again, err := msgCore.MsgSendMessageV2(acceptanceSendRequest(senderID, receiverID, 9001, randomID, "hello acceptance"))
	if err != nil {
		t.Fatalf("MsgSendMessageV2() retry error = %v", err)
	}
	againShort := mustShortSent(t, again)
	if againShort.Pts != firstShort.Pts || againShort.Id != firstShort.Id {
		t.Fatalf("retry returned different sender state: first=%+v again=%+v", firstShort, againShort)
	}
	assertDifferenceMessage(t, ctx, updatesKit, senderID, true, "hello acceptance")
	assertDifferenceMessage(t, ctx, updatesKit, receiverID, false, "hello acceptance")

	deleteCanonicalMessage(t, ctx, db, msgRepo, senderID, receiverID, randomID, "hello acceptance")
	assertDifferenceMessage(t, ctx, updatesKit, receiverID, false, "hello acceptance")
}

func TestSendMessageV2SingleChatAcceptanceRecoversSenderCommit(t *testing.T) {
	ctx := context.Background()
	db := openAcceptanceDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	idgen := &acceptanceIDGenerator{next: base + 200_000}

	realRepo := msgrepo.NewForTest(db, idgen)
	msgRepo := &acceptanceFailingSenderCommitRepo{Repository: realRepo, failNext: true}
	updatesKit := userupdatestestkit.New(db, idgen, "local-userupdates")
	senderID := base + 2001
	receiverID := base + 2002
	randomID := base + 2003
	claimUserPartitions(t, ctx, updatesKit, senderID, receiverID)

	msgCore := New(ctx, &svc.ServiceContext{
		Repo:              msgRepo,
		UserUpdates:       updatesKit,
		ReceiverPublisher: &msgrepo.InMemoryReceiverOperationPublisher{OnPublish: updatesKit.ProcessReceiverOperation},
	})

	result, err := msgCore.MsgSendMessageV2(acceptanceSendRequest(senderID, receiverID, 9002, randomID, "recover acceptance"))
	if err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	short := mustShortSent(t, result)
	if short.Pts != 1 {
		t.Fatalf("unexpected recovered sender pts: %+v", short)
	}
	if msgRepo.markSenderCalls != 2 {
		t.Fatalf("mark sender calls = %d, want 2", msgRepo.markSenderCalls)
	}
	assertDifferenceMessage(t, ctx, updatesKit, senderID, true, "recover acceptance")
	assertDifferenceMessage(t, ctx, updatesKit, receiverID, false, "recover acceptance")
}

type acceptanceIDGenerator struct {
	next int64
}

func (g *acceptanceIDGenerator) NextID(context.Context) (int64, error) {
	g.next++
	return g.next, nil
}

type acceptanceFailingSenderCommitRepo struct {
	*msgrepo.Repository
	failNext        bool
	markSenderCalls int
}

func (r *acceptanceFailingSenderCommitRepo) MarkSenderCommitted(ctx context.Context, in msgrepo.MarkSenderCommittedInput) error {
	r.markSenderCalls++
	if r.failNext {
		r.failNext = false
		return errors.New("injected mark sender committed failure")
	}
	return r.Repository.MarkSenderCommitted(ctx, in)
}

func openAcceptanceDB(t *testing.T) *sqlx.DB {
	t.Helper()
	if testing.Short() {
		t.Skip("integration test skipped in short mode")
	}
	dsn := os.Getenv("TEAMGRAM_TEST_MYSQL_DSN")
	explicit := dsn != ""
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/teamgram?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
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

func claimUserPartitions(t *testing.T, ctx context.Context, kit *userupdatestestkit.Kit, userIDs ...int64) {
	t.Helper()
	claimed := map[int]bool{}
	for _, userID := range userIDs {
		route := payload.RouteUser(userID)
		if claimed[route.ReceiverPartitionID] {
			continue
		}
		claimed[route.ReceiverPartitionID] = true
		if _, err := kit.ClaimPartitionOwner(ctx, int32(route.ReceiverPartitionID)); err != nil {
			t.Fatalf("ClaimPartitionOwner(%d) error = %v", route.ReceiverPartitionID, err)
		}
	}
}

func acceptanceSendRequest(senderID, receiverID, authKeyID, randomID int64, text string) *msgpb.TLMsgSendMessageV2 {
	return &msgpb.TLMsgSendMessageV2{
		UserId:    senderID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeUser,
		PeerId:    receiverID,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				RandomId: randomID,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: text}),
			}),
		},
	}
}

func mustShortSent(t *testing.T, updates *tg.Updates) *tg.TLUpdateShortSentMessage {
	t.Helper()
	short, ok := updates.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", updates.ClazzName())
	}
	return short
}

func assertDifferenceMessage(t *testing.T, ctx context.Context, kit *userupdatestestkit.Kit, userID int64, out bool, text string) {
	t.Helper()
	diff, err := kit.UserupdatesGetDifference(ctx, &userupdates.TLUserupdatesGetDifference{UserId: userID, AuthKeyId: 1, Pts: 0})
	if err != nil {
		t.Fatalf("UserupdatesGetDifference(%d) error = %v", userID, err)
	}
	full, ok := diff.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference for user %d, got %s", userID, diff.ClazzName())
	}
	if len(full.NewMessages) != 1 {
		t.Fatalf("user %d new_messages len = %d, want 1", userID, len(full.NewMessages))
	}
	message, ok := full.NewMessages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("user %d message type = %T", userID, full.NewMessages[0])
	}
	if message.Out != out || message.Message != text {
		t.Fatalf("user %d message mismatch: %+v", userID, message)
	}
	if len(full.OtherUpdates) != 1 {
		t.Fatalf("user %d updates len = %d, want 1", userID, len(full.OtherUpdates))
	}
}

func deleteCanonicalMessage(t *testing.T, ctx context.Context, db *sqlx.DB, repo *msgrepo.Repository, senderID, receiverID, randomID int64, text string) {
	t.Helper()
	_, requestHash, err := marshalSendRequest(senderID, payload.PeerTypeUser, receiverID, randomID, text, 0, false, 0, 0)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}
	state, err := repo.CreateOrLoadSendState(ctx, msgrepo.CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      receiverID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("load send state: %v", err)
	}
	if state.CanonicalMessageID == 0 {
		t.Fatal("send state missing canonical message id")
	}
	if _, err := db.Exec(ctx, "DELETE FROM canonical_messages WHERE canonical_message_id = ?", state.CanonicalMessageID); err != nil {
		t.Fatalf("delete canonical message: %v", err)
	}
}
