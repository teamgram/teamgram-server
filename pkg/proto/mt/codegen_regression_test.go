// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mt

import (
	"errors"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func TestTLReqPqDecodePropagatesClazzIDReadError(t *testing.T) {
	var m TLReqPq
	err := m.Decode(bin.NewDecoder(nil))
	if !errors.Is(err, bin.ErrUnexpectedEOF) {
		t.Fatalf("expected unexpected EOF, got %v", err)
	}
}

func TestGeneratedMTServiceUsesSwitchAndNoIgnoredClazzIDError(t *testing.T) {
	data, err := os.ReadFile("schema.tl.mt_service_gen.go")
	if err != nil {
		t.Fatalf("read generated file: %v", err)
	}

	src := string(data)
	for _, bad := range []string{
		"var encodeF = map[uint32]func() error",
		"var decodeF = map[uint32]func() error",
		"m.ClazzID, _ = d.ClazzID()",
	} {
		if strings.Contains(src, bad) {
			t.Fatalf("generated mt service code still contains banned pattern: %q", bad)
		}
	}
}

func TestGeneratedMTCodeDoesNotOverwriteDecodeErrors(t *testing.T) {
	files := []string{
		"schema.tl.mt_gen.go",
		"schema.tl.mt_service_gen.go",
	}
	re := regexp.MustCompile(`err = [^\n]+\n\s*err = `)

	for _, name := range files {
		data, err := os.ReadFile(name)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		if re.Find(data) != nil {
			t.Fatalf("generated file %s still contains consecutive err overwrites", name)
		}
	}
}

func TestValidateRecursesIntoRequiredObjectSlices(t *testing.T) {
	msg := &TLTlsClientHello{
		Blocks: []TlsBlockClazz{
			&TLTlsBlockScope{},
		},
	}

	if err := msg.Validate(223); err == nil {
		t.Fatalf("expected nested validation error, got nil")
	}
}

func TestGeneratedMTVectorRoundTrip(t *testing.T) {
	t.Run("future_salts", func(t *testing.T) {
		want := &TLFutureSalts{
			ReqMsgId: 1,
			Now:      2,
			Salts: []*TLFutureSalt{
				{ValidSince: 3, ValidUntil: 4, Salt: 5},
			},
		}

		data, err := iface.EncodeObject(want, 0)
		if err != nil {
			t.Fatalf("encode error: %v", err)
		}

		gotObj, err := iface.DecodeObject(bin.NewDecoder(data))
		if err != nil {
			t.Fatalf("decode error: %v", err)
		}

		got, ok := gotObj.(*TLFutureSalts)
		if !ok {
			t.Fatalf("decoded object = %T, want *TLFutureSalts", gotObj)
		}
		if got.ReqMsgId != want.ReqMsgId || got.Now != want.Now {
			t.Fatalf("decoded header = (%d,%d), want (%d,%d)", got.ReqMsgId, got.Now, want.ReqMsgId, want.Now)
		}
		if len(got.Salts) != 1 {
			t.Fatalf("decoded salts len = %d, want 1", len(got.Salts))
		}
		if got.Salts[0].ValidSince != want.Salts[0].ValidSince ||
			got.Salts[0].ValidUntil != want.Salts[0].ValidUntil ||
			got.Salts[0].Salt != want.Salts[0].Salt {
			t.Fatalf("decoded salt fields = (%d,%d,%d), want (%d,%d,%d)",
				got.Salts[0].ValidSince, got.Salts[0].ValidUntil, got.Salts[0].Salt,
				want.Salts[0].ValidSince, want.Salts[0].ValidUntil, want.Salts[0].Salt)
		}
	})

	t.Run("tls_client_hello", func(t *testing.T) {
		want := &TLTlsClientHello{
			Blocks: []TlsBlockClazz{
				&TLTlsBlockRandom{Length: 8},
				&TLTlsBlockDomain{},
			},
		}

		data, err := iface.EncodeObject(want, 0)
		if err != nil {
			t.Fatalf("encode error: %v", err)
		}

		gotObj, err := iface.DecodeObject(bin.NewDecoder(data))
		if err != nil {
			t.Fatalf("decode error: %v", err)
		}

		got, ok := gotObj.(*TLTlsClientHello)
		if !ok {
			t.Fatalf("decoded object = %T, want *TLTlsClientHello", gotObj)
		}
		if len(got.Blocks) != len(want.Blocks) {
			t.Fatalf("decoded blocks len = %d, want %d", len(got.Blocks), len(want.Blocks))
		}
		random, ok := got.Blocks[0].(*TLTlsBlockRandom)
		if !ok || random.Length != 8 {
			t.Fatalf("decoded first block = %#v, want *TLTlsBlockRandom{Length:8}", got.Blocks[0])
		}
		if _, ok := got.Blocks[1].(*TLTlsBlockDomain); !ok {
			t.Fatalf("decoded second block = %#v, want *TLTlsBlockDomain", got.Blocks[1])
		}
	})
}

func TestGeneratedMTVectorDecodeRejectsNegativeLength(t *testing.T) {
	x := bin.NewEncoder()
	x.PutInt64(1)
	x.PutInt32(2)
	x.PutInt32(-1)

	msg := &TLFutureSalts{ClazzID: ClazzID_future_salts}
	err := msg.Decode(bin.NewDecoder(x.Bytes()))
	if err == nil {
		t.Fatal("expected decode error")
	}

	var invalidLen *bin.InvalidLengthError
	if !errors.As(err, &invalidLen) {
		t.Fatalf("expected invalid length error, got %v", err)
	}
	if invalidLen.Length != -1 {
		t.Fatalf("invalid length = %d, want -1", invalidLen.Length)
	}
}
