package repository

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

func TestStorageErrorPreservesCause(t *testing.T) {
	cause := errors.New("mysql unavailable")

	err := storageError("query", cause)

	if !errors.Is(err, msg.ErrMsgStorage) {
		t.Fatalf("storageError() does not wrap ErrMsgStorage: %v", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("storageError() does not preserve cause: %v", err)
	}
}

func TestMessageIDResolversAllowZeroIDWithoutStorage(t *testing.T) {
	repo := &Repository{}

	resolved, err := repo.ResolveMessageID(context.Background(), 1, payload.PeerTypeUser, 2, 0)
	if err != nil {
		t.Fatalf("ResolveMessageID zero error = %v", err)
	}
	if resolved != nil {
		t.Fatalf("ResolveMessageID zero = %+v, want nil", resolved)
	}

	publicID, err := repo.ResolvePeerSeqToUserMessageID(context.Background(), 1, payload.PeerTypeUser, 2, 0)
	if err != nil {
		t.Fatalf("ResolvePeerSeqToUserMessageID zero error = %v", err)
	}
	if publicID != 0 {
		t.Fatalf("ResolvePeerSeqToUserMessageID zero = %d, want 0", publicID)
	}

	bounds, err := repo.ResolveHistoryCursorIDs(context.Background(), 1, payload.PeerTypeUser, 2, 0, 0, 0)
	if err != nil {
		t.Fatalf("ResolveHistoryCursorIDs zero error = %v", err)
	}
	if bounds != (HistoryCursorBounds{}) {
		t.Fatalf("ResolveHistoryCursorIDs zero = %+v, want zero bounds", bounds)
	}
}

func TestDeleteResolverQueryDoesNotFilterLiveStatus(t *testing.T) {
	body, err := os.ReadFile("model/tables/queries.xml")
	if err != nil {
		t.Fatalf("read queries.xml: %v", err)
	}
	normal := queryXMLBlock(t, string(body), "SelectUserMessageByGlobalID")
	if !strings.Contains(normal, "message_status = :message_status") {
		t.Fatalf("normal resolver query should remain live-status filtered:\n%s", normal)
	}
	deleteQuery := queryXMLBlock(t, string(body), "SelectUserMessageByGlobalIDForDelete")
	if strings.Contains(deleteQuery, "message_status") {
		t.Fatalf("delete resolver query should include deleted views by omitting message_status:\n%s", deleteQuery)
	}
}

func TestGetUserMessageQueriesUseViewerPublicMessageID(t *testing.T) {
	body, err := os.ReadFile("model/tables/queries.xml")
	if err != nil {
		t.Fatalf("read queries.xml: %v", err)
	}
	query := queryXMLBlock(t, string(body), "SelectUserMessageBoxByGlobalID")
	for _, want := range []string{
		"v.user_id = :user_id",
		"v.user_message_id = :user_message_id",
		"v.message_status = :message_status",
		"JOIN",
		"canonical_messages c",
		"c.canonical_message_id = v.canonical_message_id",
	} {
		if !strings.Contains(query, want) {
			t.Fatalf("get user message query missing %q:\n%s", want, query)
		}
	}
	if strings.Contains(query, "v.peer_type = :peer_type") || strings.Contains(query, "v.peer_id = :peer_id") {
		t.Fatalf("get user message query must use user_id + user_message_id public identity only:\n%s", query)
	}
}

func TestGetUserMessageListRejectsInvalidIDBeforeStorage(t *testing.T) {
	repo := &Repository{}

	_, err := repo.GetUserMessageList(context.Background(), 100, []int64{7, 0})
	if !errors.Is(err, msg.ErrMsgIdInvalid) {
		t.Fatalf("GetUserMessageList() error = %v, want ErrMsgIdInvalid", err)
	}
}

func TestGetUserMessageListRejectsMissingID(t *testing.T) {
	repo := &Repository{
		db: &sqlx.DB{},
		models: &model.Models{
			CanonicalQueries: fakeCanonicalQueriesModel{
				boxes: map[int64]*model.UserMessageBoxRow{
					7: {
						UserID:             100,
						UserMessageID:      7,
						CanonicalMessageID: 900,
						PeerType:           payload.PeerTypeUser,
						PeerID:             200,
						PeerSeq:            3,
						FromUserID:         200,
						MessageDate:        1_772_000_300,
					},
				},
			},
		},
	}

	_, err := repo.GetUserMessageList(context.Background(), 100, []int64{7, 99})
	if !errors.Is(err, msg.ErrMsgIdInvalid) {
		t.Fatalf("GetUserMessageList() error = %v, want ErrMsgIdInvalid", err)
	}
}

func TestRevalidateForwardSourcesQueriesLiveViewerPublicMessageID(t *testing.T) {
	body, err := os.ReadFile("model/tables/queries.xml")
	if err != nil {
		t.Fatalf("read queries.xml: %v", err)
	}
	query := queryXMLBlock(t, string(body), "SelectForwardSourceIdentity")
	for _, want := range []string{
		"user_id = :user_id",
		"user_message_id = :user_message_id",
		"message_status = :message_status",
		"canonical_message_id",
	} {
		if !strings.Contains(query, want) {
			t.Fatalf("forward source query missing %q:\n%s", want, query)
		}
	}
}

func TestHistoryMessageRowToMessageUsesPublicIDs(t *testing.T) {
	event := payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.EventKindNewMessage,
		CanonicalMessageID:   1001,
		PeerSeq:              7,
		MessageID:            42,
		PeerType:             payload.PeerTypeUser,
		PeerID:               2002,
		FromUserID:           1002,
		ToUserID:             1001,
		Date:                 1_772_000_000,
		MessageText:          "reply",
		ReplyToUserMessageID: 21,
	}
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}

	msg, err := historyMessageRowToMessage(model.HistoryMessageRow{
		CanonicalMessageID: 1001,
		PeerSeq:            7,
		UserMessageID:      42,
		FromUserID:         1002,
		Outgoing:           false,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2002,
		MessageKind:        MessageKindText,
		MessageText:        "reply",
		MessageDate:        1_772_000_000,
		ViewPayload:        body,
	}, nil)
	if err != nil {
		t.Fatalf("historyMessageRowToMessage error = %v", err)
	}
	if msg.UserMessageID != 42 {
		t.Fatalf("UserMessageID = %d, want 42", msg.UserMessageID)
	}
	if msg.PeerSeq != 7 {
		t.Fatalf("PeerSeq = %d, want internal peer seq 7", msg.PeerSeq)
	}
	if msg.ReplyToUserMessageID != 21 {
		t.Fatalf("ReplyToUserMessageID = %d, want 21", msg.ReplyToUserMessageID)
	}
	if msg.ReplyToPeerSeq != 0 {
		t.Fatalf("ReplyToPeerSeq = %d, want 0 for v2 payload", msg.ReplyToPeerSeq)
	}
}

func queryXMLBlock(t *testing.T, body string, name string) string {
	t.Helper()
	startTag := `<query name="` + name + `"`
	start := strings.Index(body, startTag)
	if start < 0 {
		t.Fatalf("query %s not found", name)
	}
	end := strings.Index(body[start:], "</query>")
	if end < 0 {
		t.Fatalf("query %s has no closing tag", name)
	}
	return body[start : start+end]
}

func TestHistoryMessageRowToMessageResolvesLegacyReplyPeerSeq(t *testing.T) {
	event := payload.MessageEventV1{
		SchemaVersion:      1,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 1001,
		MessageID:          42,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2002,
		FromUserID:         1002,
		ToUserID:           1001,
		Date:               1_772_000_000,
		MessageText:        "reply",
		ReplyToPeerSeq:     9,
	}
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}

	msg, err := historyMessageRowToMessage(model.HistoryMessageRow{
		CanonicalMessageID: 1001,
		PeerSeq:            17,
		UserMessageID:      42,
		FromUserID:         1002,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2002,
		MessageKind:        MessageKindText,
		MessageText:        "reply",
		MessageDate:        1_772_000_000,
		ViewPayload:        body,
	}, func(peerSeq int64) (int64, error) {
		if peerSeq != 9 {
			t.Fatalf("resolver peer_seq = %d, want 9", peerSeq)
		}
		return 84, nil
	})
	if err != nil {
		t.Fatalf("historyMessageRowToMessage error = %v", err)
	}
	if msg.ReplyToPeerSeq != 9 {
		t.Fatalf("ReplyToPeerSeq = %d, want internal fallback peer_seq 9", msg.ReplyToPeerSeq)
	}
	if msg.ReplyToUserMessageID != 84 {
		t.Fatalf("ReplyToUserMessageID = %d, want resolved public id 84", msg.ReplyToUserMessageID)
	}
}

func TestHistoryPeerSeqBoundsKeepInt64Values(t *testing.T) {
	peerSeq := int64(math.MaxInt32) + 10
	bounds := HistoryCursorBounds{
		OffsetPeerSeq: int64(math.MaxInt32) + 20,
		MaxPeerSeq:    int64(math.MaxInt32) + 20,
		MinPeerSeq:    int64(math.MaxInt32) + 1,
	}
	if !historyMessageWithinBounds(peerSeq, bounds) {
		t.Fatalf("peer_seq %d should be inside int64 bounds %+v", peerSeq, bounds)
	}
	if historyMessageWithinBounds(bounds.MaxPeerSeq, bounds) {
		t.Fatalf("peer_seq equal max bound should be excluded")
	}
	if historyMessageWithinBounds(bounds.MinPeerSeq, bounds) {
		t.Fatalf("peer_seq equal min bound should be excluded")
	}
	if got := historyOffsetMarker(bounds.OffsetPeerSeq); got.OffsetID <= 0 {
		t.Fatalf("historyOffsetMarker(%d).OffsetID = %d, want positive marker without int32 cast", bounds.OffsetPeerSeq, got.OffsetID)
	}
}

func TestHistoryPeerSeqBoundsNoMatchExcludesAllRows(t *testing.T) {
	bounds := HistoryCursorBounds{NoMatch: true}
	if historyMessageWithinBounds(1, bounds) {
		t.Fatalf("historyMessageWithinBounds() = true, want false for unresolved positive public cursor")
	}
}

func TestSearchHashTagOffsetUsesResolvedPeerSeq(t *testing.T) {
	publicOffsetID := int32(42)
	resolved := &ResolvedMessageID{
		UserMessageID: int64(publicOffsetID),
		PeerSeq:       9001,
	}
	got, noMatch := searchHashTagOffsetPeerSeq(publicOffsetID, resolved)
	if got != 9001 || noMatch {
		t.Fatalf("searchHashTagOffsetPeerSeq() = %d, want resolved peer_seq 9001", got)
	}
	got, noMatch = searchHashTagOffsetPeerSeq(publicOffsetID, nil)
	if got != 0 || !noMatch {
		t.Fatalf("searchHashTagOffsetPeerSeq(nil) = (%d, %t), want no-match for positive unresolved offset", got, noMatch)
	}
	got, noMatch = searchHashTagOffsetPeerSeq(0, resolved)
	if got != 0 || noMatch {
		t.Fatalf("searchHashTagOffsetPeerSeq(0) = %d, want 0", got)
	}
}

type fakeCanonicalQueriesModel struct {
	boxes map[int64]*model.UserMessageBoxRow
}

func (f fakeCanonicalQueriesModel) SelectUserMessageBoxByGlobalID(_ context.Context, userID int64, userMessageID int64, messageStatus int32) (*model.UserMessageBoxRow, error) {
	row := f.boxes[userMessageID]
	if row == nil || row.UserID != userID || messageStatus != MessageStatusLive {
		return nil, model.ErrNotFound
	}
	return row, nil
}

func (fakeCanonicalQueriesModel) SelectCanonicalByRandom(context.Context, int64, int32, int64, int64) (*model.CanonicalMessageRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) SelectCanonicalByID(context.Context, int64, []byte, int64) (*model.CanonicalMessageRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) SelectUserMessageByID(context.Context, int64, int32, int64, int64, int32) (*model.ResolvedMessageIDRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) SelectUserMessageByGlobalID(context.Context, int64, int64, int32) (*model.ResolvedMessageIDRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) SelectUserMessageByGlobalIDForDelete(context.Context, int64, int64) (*model.ResolvedMessageIDRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) SelectForwardSourceIdentity(context.Context, int64, int64, int32) (*model.ForwardSourceIdentityRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) SelectNearestLiveUserMessageByPeerSeq(context.Context, int64, int32, int64, int64, int32) (*model.ResolvedMessageIDRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) SelectHistoryMessages(context.Context, int64, int32, int64, int32, int64, int64, int32) ([]model.HistoryMessageRow, error) {
	return nil, nil
}

func (fakeCanonicalQueriesModel) SearchHashTagMessages(context.Context, string, int64, int32, int64, int32, int64, int64, string, int32) ([]model.HistoryMessageRow, error) {
	return nil, nil
}

func (fakeCanonicalQueriesModel) SelectCanonicalByUserView(context.Context, int64, int32, int64, int64, int32) (*model.HistoryMessageRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) SelectEditableMessageForUpdate(context.Context, int64, int32, int64, int64, int32) (*model.EditableMessageRow, error) {
	return nil, model.ErrNotFound
}

func (fakeCanonicalQueriesModel) CountHistoryOffset(context.Context, int64, int64, int64, int64, int64, int32, int64, int32) (*model.HistoryOffsetRow, error) {
	return &model.HistoryOffsetRow{}, nil
}

func (fakeCanonicalQueriesModel) SelectHistoryMessagesPage(context.Context, int64, int32, int64, int32, int64, int32) ([]model.HistoryMessageRow, error) {
	return nil, nil
}

func (fakeCanonicalQueriesModel) SelectHistoryMessagesBackwardByUserMessageID(context.Context, int64, int32, int64, int32, int64, int32) ([]model.HistoryMessageRow, error) {
	return nil, nil
}

func (fakeCanonicalQueriesModel) SelectHistoryMessagesForwardByUserMessageID(context.Context, int64, int32, int64, int32, int64, int32) ([]model.HistoryMessageRow, error) {
	return nil, nil
}

func (fakeCanonicalQueriesModel) SelectConversationViewPeerSeqFloor(context.Context, int32, int64, int64, int64, int64) (*model.PeerSeqFloorRow, error) {
	return &model.PeerSeqFloorRow{}, nil
}
