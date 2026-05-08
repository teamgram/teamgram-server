package repository

import (
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestPublicMessageIDsCanIncrementAcrossDialogsWhilePeerSeqResets(t *testing.T) {
	first := payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: 7001,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2001,
		PeerSeq:            1,
		UserMessageID:      1,
		FromUserID:         1001,
		ToUserID:           2001,
		Date:               1_772_000_001,
		Out:                true,
		MessageText:        "first dialog",
	}
	second := first
	second.CanonicalMessageID = 7002
	second.PeerID = 2002
	second.ToUserID = 2002
	second.PeerSeq = 1
	second.UserMessageID = 2
	second.MessageText = "second dialog"

	firstEvent, _, firstResponse, _, err := buildEventAndResponse(ApplyUserOperationInput{OperationID: "first"}, first, 1, 1)
	if err != nil {
		t.Fatalf("buildEventAndResponse(first) error = %v", err)
	}
	secondEvent, _, secondResponse, _, err := buildEventAndResponse(ApplyUserOperationInput{OperationID: "second"}, second, 2, 1)
	if err != nil {
		t.Fatalf("buildEventAndResponse(second) error = %v", err)
	}

	var event1, event2 payload.MessageEventV2
	mustUnmarshalContractJSON(t, firstEvent, &event1)
	mustUnmarshalContractJSON(t, secondEvent, &event2)
	if event1.PeerSeq != 1 || event2.PeerSeq != 1 {
		t.Fatalf("peer seqs = %d,%d want both dialogs at peer_seq 1", event1.PeerSeq, event2.PeerSeq)
	}
	if event1.MessageID != 1 || event2.MessageID != 2 {
		t.Fatalf("event public ids = %d,%d want per-user ids 1,2", event1.MessageID, event2.MessageID)
	}

	var response1, response2 payload.OperationResponseV2
	mustUnmarshalContractJSON(t, firstResponse, &response1)
	mustUnmarshalContractJSON(t, secondResponse, &response2)
	if response1.UserMessageID != 1 || response2.UserMessageID != 2 {
		t.Fatalf("response public ids = %d,%d want per-user ids 1,2", response1.UserMessageID, response2.UserMessageID)
	}
}

func TestSenderAndReceiverEventsCanUseDifferentPublicIDsForSameCanonicalMessage(t *testing.T) {
	sender := payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: 8001,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2002,
		PeerSeq:            3,
		UserMessageID:      12,
		FromUserID:         1001,
		ToUserID:           2002,
		Date:               1_772_000_002,
		Out:                true,
		MessageText:        "same canonical",
	}
	receiver := sender
	receiver.PeerID = 1001
	receiver.UserMessageID = 4
	receiver.Out = false

	senderEventBody, _, senderResponseBody, _, err := buildEventAndResponse(ApplyUserOperationInput{OperationID: "sender"}, sender, 5, 1)
	if err != nil {
		t.Fatalf("buildEventAndResponse(sender) error = %v", err)
	}
	receiverEventBody, _, receiverResponseBody, _, err := buildEventAndResponse(ApplyUserOperationInput{OperationID: "receiver"}, receiver, 6, 1)
	if err != nil {
		t.Fatalf("buildEventAndResponse(receiver) error = %v", err)
	}

	var senderEvent, receiverEvent payload.MessageEventV2
	mustUnmarshalContractJSON(t, senderEventBody, &senderEvent)
	mustUnmarshalContractJSON(t, receiverEventBody, &receiverEvent)
	if senderEvent.CanonicalMessageID != receiverEvent.CanonicalMessageID {
		t.Fatalf("canonical ids differ: sender=%d receiver=%d", senderEvent.CanonicalMessageID, receiverEvent.CanonicalMessageID)
	}
	if senderEvent.MessageID != 12 || receiverEvent.MessageID != 4 {
		t.Fatalf("sender/receiver event ids = %d/%d want 12/4", senderEvent.MessageID, receiverEvent.MessageID)
	}

	var senderResponse, receiverResponse payload.OperationResponseV2
	mustUnmarshalContractJSON(t, senderResponseBody, &senderResponse)
	mustUnmarshalContractJSON(t, receiverResponseBody, &receiverResponse)
	if senderResponse.UserMessageID != 12 || receiverResponse.UserMessageID != 4 {
		t.Fatalf("sender/receiver response ids = %d/%d want 12/4", senderResponse.UserMessageID, receiverResponse.UserMessageID)
	}
}

func TestSavedDialogSideEffectPayloadCarriesPublicTopMessageID(t *testing.T) {
	body, err := json.Marshal(savedDialogSideEffectPayloadV1{
		SchemaVersion:         1,
		SavedPeerType:         payload.PeerTypeUser,
		SavedPeerID:           1001,
		TopPeerSeq:            7,
		TopUserMessageID:      55,
		TopCanonicalMessageID: 7007,
		MessageDate:           1_772_000_003,
		Top:                   true,
	})
	if err != nil {
		t.Fatalf("marshal saved dialog side effect: %v", err)
	}
	var got map[string]interface{}
	mustUnmarshalContractJSON(t, body, &got)
	if got["top_user_message_id"] != float64(55) || got["top_peer_seq"] != float64(7) {
		t.Fatalf("saved dialog payload ids = %v, want public top id 55 and peer_seq 7", got)
	}
}

func mustUnmarshalContractJSON(t *testing.T, body []byte, out interface{}) {
	t.Helper()
	if err := json.Unmarshal(body, out); err != nil {
		t.Fatalf("unmarshal contract json: %v", err)
	}
}
