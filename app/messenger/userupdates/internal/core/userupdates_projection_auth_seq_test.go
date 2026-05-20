package core

import (
	"errors"
	"math"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthSeqEventToTLUpdateDecodesTLBinary(t *testing.T) {
	update := tg.MakeTLUpdatePeerSettings(&tg.TLUpdatePeerSettings{
		Peer:     tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 42}),
		Settings: tg.MakeTLPeerSettings(&tg.TLPeerSettings{}),
	})
	body, err := iface.EncodeObject(update, repository.AuthSeqLayer)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	got, err := authSeqEventToTLUpdate(repository.AuthSeqEvent{
		EventCodec:         repository.AuthSeqCodecTLBinary,
		EventSchemaVersion: repository.AuthSeqLayer,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	})
	if err != nil {
		t.Fatalf("authSeqEventToTLUpdate() error = %v", err)
	}
	if _, ok := got.(*tg.TLUpdatePeerSettings); !ok {
		t.Fatalf("decoded update = %T, want *tg.TLUpdatePeerSettings", got)
	}
}

func TestAuthSeqEventToTLUpdateRejectsEmptyPayloadHash(t *testing.T) {
	update := tg.MakeTLUpdatePeerSettings(&tg.TLUpdatePeerSettings{
		Peer:     tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 42}),
		Settings: tg.MakeTLPeerSettings(&tg.TLPeerSettings{}),
	})
	body, err := iface.EncodeObject(update, repository.AuthSeqLayer)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	_, err = authSeqEventToTLUpdate(repository.AuthSeqEvent{
		EventCodec:         repository.AuthSeqCodecTLBinary,
		EventSchemaVersion: repository.AuthSeqLayer,
		EventPayload:       body,
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("authSeqEventToTLUpdate() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestAuthSeqEventToTLUpdateRejectsTrailingBytes(t *testing.T) {
	update := tg.MakeTLUpdatePeerSettings(&tg.TLUpdatePeerSettings{
		Peer:     tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 42}),
		Settings: tg.MakeTLPeerSettings(&tg.TLPeerSettings{}),
	})
	body, err := iface.EncodeObject(update, repository.AuthSeqLayer)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	body = append(body, 1, 2, 3, 4)
	_, err = authSeqEventToTLUpdate(repository.AuthSeqEvent{
		EventCodec:         repository.AuthSeqCodecTLBinary,
		EventSchemaVersion: repository.AuthSeqLayer,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("authSeqEventToTLUpdate() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestDifferenceToTLRejectsStateSeqOverflow(t *testing.T) {
	_, err := differenceToTL(&repository.GetDifferenceResult{
		State: repository.UserState{Seq: int64(math.MaxInt32) + 1},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("differenceToTL() error = %v, want ErrUserupdatesStorage", err)
	}
}
