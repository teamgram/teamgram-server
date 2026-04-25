package mt

import (
	"bytes"
	"compress/gzip"
	"io"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

func TestGZIP_Encode(t *testing.T) {
	payload := testGzipPayload(t)
	x := bin.NewEncoder()
	defer x.End()

	if err := (GZIP{Data: payload}).Encode(x); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	d := bin.NewDecoder(x.Bytes())
	clazzID, err := d.ClazzID()
	if err != nil {
		t.Fatalf("constructor decode error = %v", err)
	}
	if clazzID != ClazzID_gzip_packed {
		t.Fatalf("constructor = %#x, want %#x", clazzID, ClazzID_gzip_packed)
	}

	compressed, err := d.Bytes()
	if err != nil {
		t.Fatalf("packed_data decode error = %v", err)
	}
	if d.Remaining() != 0 {
		t.Fatalf("encoded object has %d trailing bytes", d.Remaining())
	}

	got := gunzipBytes(t, compressed)
	if !bytes.Equal(got, payload) {
		t.Fatalf("decompressed payload = %x, want %x", got, payload)
	}
}

func TestGZIP_Decode(t *testing.T) {
	payload := testGzipPayload(t)
	x := bin.NewEncoder()
	defer x.End()
	if err := (GZIP{Data: payload}).Encode(x); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	var got GZIP
	if err := got.Decode(bin.NewDecoder(x.Bytes())); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if !bytes.Equal(got.Data, payload) {
		t.Fatalf("Data = %x, want %x", got.Data, payload)
	}
}

func BenchmarkGZIP_Encode(b *testing.B) {
	payload := testGzipPayload(b)
	b.ReportAllocs()
	b.SetBytes(int64(len(payload)))

	for i := 0; i < b.N; i++ {
		x := bin.NewEncoder()
		err := (GZIP{Data: payload}).Encode(x)
		x.End()
		if err != nil {
			b.Fatalf("Encode() error = %v", err)
		}
	}
}

func BenchmarkGZIP_Decode(b *testing.B) {
	payload := testGzipPayload(b)
	x := bin.NewEncoder()
	if err := (GZIP{Data: payload}).Encode(x); err != nil {
		b.Fatalf("Encode() error = %v", err)
	}
	encoded := x.Clone()
	x.End()

	b.ReportAllocs()
	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var got GZIP
		if err := got.Decode(bin.NewDecoder(encoded)); err != nil {
			b.Fatalf("Decode() error = %v", err)
		}
	}
}

func testGzipPayload(tb testing.TB) []byte {
	tb.Helper()

	return []byte(strings.Repeat("teamgram gzip payload ", 64))
}

func gunzipBytes(tb testing.TB, payload []byte) []byte {
	tb.Helper()

	gz, err := gzip.NewReader(bytes.NewReader(payload))
	if err != nil {
		tb.Fatalf("gzip reader: %v", err)
	}
	defer gz.Close()

	data, err := io.ReadAll(gz)
	if err != nil {
		tb.Fatalf("gzip read: %v", err)
	}
	return data
}
