package authsessionservice

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestQueryAuthKeyResultBinaryDecodeReadsInnerConstructor(t *testing.T) {
	src := &QueryAuthKeyResult{
		Success: tg.NewAuthKeyInfo(123, []byte("auth-key"), tg.AuthKeyTypePerm),
	}

	var dst QueryAuthKeyResult
	roundTripResult(t, src, &dst)

	if dst.Success == nil {
		t.Fatal("decoded auth key is nil")
	}
	if dst.Success.AuthKeyId != src.Success.AuthKeyId || string(dst.Success.AuthKey) != string(src.Success.AuthKey) {
		t.Fatalf("decoded auth key = %#v, want %#v", dst.Success, src.Success)
	}
}

func TestGetFutureSaltsResultBinaryDecodeReadsInnerConstructor(t *testing.T) {
	src := &GetFutureSaltsResult{
		Success: tg.MakeTLFutureSalts(&tg.TLFutureSalts{
			ReqMsgId: 11,
			Now:      22,
			Salts: []*tg.TLFutureSalt{{
				ValidSince: 33,
				ValidUntil: 44,
				Salt:       55,
			}},
		}),
	}

	var dst GetFutureSaltsResult
	roundTripResult(t, src, &dst)

	if dst.Success == nil {
		t.Fatal("decoded future salts is nil")
	}
	if dst.Success.ReqMsgId != src.Success.ReqMsgId || len(dst.Success.Salts) != 1 || dst.Success.Salts[0].Salt != 55 {
		t.Fatalf("decoded future salts = %#v, want %#v", dst.Success, src.Success)
	}
}

func TestResetAuthorizationResultBinaryDecodeDoesNotConsumeVectorAsConstructor(t *testing.T) {
	src := &ResetAuthorizationResult{
		Success: &authsession.VectorLong{Datas: []int64{11, 22}},
	}

	var dst ResetAuthorizationResult
	roundTripResult(t, src, &dst)

	if dst.Success == nil {
		t.Fatal("decoded vector is nil")
	}
	if len(dst.Success.Datas) != 2 || dst.Success.Datas[0] != 11 || dst.Success.Datas[1] != 22 {
		t.Fatalf("decoded vector = %#v, want %#v", dst.Success, src.Success)
	}
}

func TestSetAuthKeyResultBinaryDecodeLetsBoolConsumeConstructor(t *testing.T) {
	src := &SetAuthKeyResult{Success: tg.BoolTrue}

	var dst SetAuthKeyResult
	roundTripResult(t, src, &dst)

	if _, ok := dst.Success.ToBoolTrue(); !ok {
		t.Fatalf("decoded bool = %#v, want BoolTrue", dst.Success)
	}
}

type binaryResult interface {
	Encode(*bin.Encoder, int32) error
	Decode(*bin.Decoder) error
}

func roundTripResult(t *testing.T, src, dst binaryResult) {
	t.Helper()
	encoder := bin.NewEncoder()
	defer encoder.End()
	if err := src.Encode(encoder, 223); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	if err := dst.Decode(bin.NewDecoder(encoder.Bytes())); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
}
