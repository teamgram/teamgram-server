package core

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestDifferenceSliceAdvancesOnlyDeliveredAuthSeq(t *testing.T) {
	firstPayload := encodeAuthSeqTestUpdate(t, 1001)
	secondPayload := encodeAuthSeqTestUpdate(t, 1002)
	diff := &repository.GetDifferenceResult{
		State: repository.UserState{UserID: 1, Pts: 0, Seq: 5, Date: 50},
		AuthSeqEvents: []repository.AuthSeqEvent{
			{
				UserID:             1,
				PermAuthKeyID:      10,
				Seq:                1,
				Date:               10,
				EventSchemaVersion: repository.AuthSeqLayer,
				EventCodec:         repository.AuthSeqCodecTLBinary,
				EventPayload:       firstPayload,
				EventPayloadHash:   payload.HashBytes(firstPayload),
			},
			{
				UserID:             1,
				PermAuthKeyID:      10,
				Seq:                2,
				Date:               20,
				EventSchemaVersion: repository.AuthSeqLayer,
				EventCodec:         repository.AuthSeqCodecTLBinary,
				EventPayload:       secondPayload,
				EventPayloadHash:   payload.HashBytes(secondPayload),
			},
		},
		HasMore: true,
	}
	got, err := differenceToTL(diff)
	if err != nil {
		t.Fatalf("differenceToTL() error = %v", err)
	}
	slice, ok := got.ToUserDifferenceSlice()
	if !ok {
		t.Fatalf("difference = %s, want slice", got.ClazzName())
	}
	if slice.IntermediateState.Seq != 2 || slice.IntermediateState.Date != 20 {
		t.Fatalf("intermediate seq/date = %d/%d, want 2/20", slice.IntermediateState.Seq, slice.IntermediateState.Date)
	}
}

func encodeAuthSeqTestUpdate(t *testing.T, userID int64) []byte {
	t.Helper()
	update := tg.MakeTLUpdatePeerSettings(&tg.TLUpdatePeerSettings{
		Peer:     tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: userID}),
		Settings: tg.MakeTLPeerSettings(&tg.TLPeerSettings{}),
	})
	body, err := iface.EncodeObject(update, repository.AuthSeqLayer)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	return body
}
