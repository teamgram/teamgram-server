package mt

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func TestTLGzipPacked_Encode(t *testing.T) {
	payload := testGzipPayload(t)
	x := bin.NewEncoder()
	defer x.End()

	if err := (&TLGzipPacked{PackedData: payload}).Encode(x, 0); err != nil {
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

func TestTLGzipPacked_Decode(t *testing.T) {
	payload := testGzipPayload(t)
	x := bin.NewEncoder()
	defer x.End()
	if err := (&TLGzipPacked{PackedData: payload}).Encode(x, 0); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	got := decodeGzipPacked(t, x.Bytes())
	if !bytes.Equal(got.PackedData, payload) {
		t.Fatalf("PackedData = %x, want %x", got.PackedData, payload)
	}
	if ping, ok := got.Obj.(*TLPing); !ok || ping.PingId != testGzipPingID {
		t.Fatalf("Obj = %#v, want *TLPing with PingId %d", got.Obj, testGzipPingID)
	}
}

func TestTLGzipPacked_DecodeZlibPayload(t *testing.T) {
	payload := testGzipPayload(t)
	x := bin.NewEncoder()
	defer x.End()
	x.PutClazzID(ClazzID_gzip_packed)
	x.PutBytes(zlibBytes(t, payload))

	got := decodeGzipPacked(t, x.Bytes())
	if !bytes.Equal(got.PackedData, payload) {
		t.Fatalf("PackedData = %x, want %x", got.PackedData, payload)
	}
	if ping, ok := got.Obj.(*TLPing); !ok || ping.PingId != testGzipPingID {
		t.Fatalf("Obj = %#v, want *TLPing with PingId %d", got.Obj, testGzipPingID)
	}
}

func BenchmarkTLGzipPacked_Encode(b *testing.B) {
	payload := testGzipPayload(b)
	b.ReportAllocs()
	b.SetBytes(int64(len(payload)))

	for i := 0; i < b.N; i++ {
		x := bin.NewEncoder()
		err := (&TLGzipPacked{PackedData: payload}).Encode(x, 0)
		x.End()
		if err != nil {
			b.Fatalf("Encode() error = %v", err)
		}
	}
}

func BenchmarkTLGzipPacked_Decode(b *testing.B) {
	payload := testGzipPayload(b)
	x := bin.NewEncoder()
	if err := (&TLGzipPacked{PackedData: payload}).Encode(x, 0); err != nil {
		b.Fatalf("Encode() error = %v", err)
	}
	encoded := x.Clone()
	x.End()

	b.ReportAllocs()
	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := iface.DecodeObject(bin.NewDecoder(encoded)); err != nil {
			b.Fatalf("Decode() error = %v", err)
		}
	}
}

const testGzipPingID = 0x1122334455667788

func testGzipPayload(tb testing.TB) []byte {
	tb.Helper()

	x := bin.NewEncoder()
	defer x.End()
	if err := (&TLPing{PingId: testGzipPingID}).Encode(x, 0); err != nil {
		tb.Fatalf("encode ping payload: %v", err)
	}
	return x.Clone()
}

func decodeGzipPacked(tb testing.TB, data []byte) *TLGzipPacked {
	tb.Helper()

	obj, err := iface.DecodeObject(bin.NewDecoder(data))
	if err != nil {
		tb.Fatalf("DecodeObject() error = %v", err)
	}
	got, ok := obj.(*TLGzipPacked)
	if !ok {
		tb.Fatalf("DecodeObject() = %T, want *TLGzipPacked", obj)
	}
	return got
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

func zlibBytes(tb testing.TB, payload []byte) []byte {
	tb.Helper()

	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	if _, err := zw.Write(payload); err != nil {
		tb.Fatalf("zlib write: %v", err)
	}
	if err := zw.Close(); err != nil {
		tb.Fatalf("zlib close: %v", err)
	}
	return buf.Bytes()
}
