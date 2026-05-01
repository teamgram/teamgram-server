package transport

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
	"testing"
)

func TestCodecAbridgedRoundTrip(t *testing.T) {
	payload := []byte("abcdefghijkl")
	wire := append([]byte{0xef}, mustWriteFrame(t, AbridgedCodec{}, payload)...)

	r := bufio.NewReader(bytes.NewReader(wire))
	codec, err := DetectCodec(r)
	if err != nil {
		t.Fatalf("DetectCodec() error = %v", err)
	}
	if _, ok := codec.(*AbridgedCodec); !ok {
		t.Fatalf("DetectCodec() = %T, want *AbridgedCodec", codec)
	}

	got, err := codec.ReadFrame(r)
	if err != nil {
		t.Fatalf("ReadFrame() error = %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("ReadFrame() = %x, want %x", got, payload)
	}
}

func TestCodecIntermediateRoundTrip(t *testing.T) {
	payload := []byte("abcdefghijklmnop")
	wire := append([]byte{0xee, 0xee, 0xee, 0xee}, mustWriteFrame(t, IntermediateCodec{}, payload)...)

	r := bufio.NewReader(bytes.NewReader(wire))
	codec, err := DetectCodec(r)
	if err != nil {
		t.Fatalf("DetectCodec() error = %v", err)
	}
	if _, ok := codec.(*IntermediateCodec); !ok {
		t.Fatalf("DetectCodec() = %T, want *IntermediateCodec", codec)
	}

	got, err := codec.ReadFrame(r)
	if err != nil {
		t.Fatalf("ReadFrame() error = %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("ReadFrame() = %x, want %x", got, payload)
	}
}

func TestCodecDetectionKeepsIntermediateHeaderIntact(t *testing.T) {
	payload := []byte("abcdefghijkl")
	wire := append([]byte{0xee, 0xee, 0xee, 0xee}, mustWriteFrame(t, IntermediateCodec{}, payload)...)

	r := bufio.NewReader(bytes.NewReader(wire))
	codec, err := DetectCodec(r)
	if err != nil {
		t.Fatalf("DetectCodec() error = %v", err)
	}
	if _, ok := codec.(*IntermediateCodec); !ok {
		t.Fatalf("DetectCodec() = %T, want *IntermediateCodec", codec)
	}
	next, err := r.Peek(4)
	if err != nil {
		t.Fatalf("Peek(4) after DetectCodec() error = %v", err)
	}
	var gotLen uint32
	binary.Read(bytes.NewReader(next), binary.LittleEndian, &gotLen)
	if gotLen != uint32(len(payload)) {
		t.Fatalf("remaining frame length = %d, want %d", gotLen, len(payload))
	}
}

func TestCodecFullRoundTrip(t *testing.T) {
	payload := []byte("abcdefghijklmnop")
	var wire bytes.Buffer
	codec := &FullCodec{}
	if err := codec.WriteFrame(&wire, payload); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}

	r := bufio.NewReader(bytes.NewReader(wire.Bytes()))
	detected, err := DetectCodec(r)
	if err != nil {
		t.Fatalf("DetectCodec() error = %v", err)
	}
	full, ok := detected.(*FullCodec)
	if !ok {
		t.Fatalf("DetectCodec() = %T, want *FullCodec", detected)
	}

	got, err := full.ReadFrame(r)
	if err != nil {
		t.Fatalf("ReadFrame() error = %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("ReadFrame() = %x, want %x", got, payload)
	}
}

func TestCodecFullFrameFields(t *testing.T) {
	payload := []byte("abcdefghijkl")
	var wire bytes.Buffer
	codec := &FullCodec{}
	if err := codec.WriteFrame(&wire, payload); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}
	frame := wire.Bytes()
	totalLen := binary.LittleEndian.Uint32(frame[:4])
	seqNo := binary.LittleEndian.Uint32(frame[4:8])
	wantLen := uint32(4 + 4 + len(payload) + 4)
	if totalLen != wantLen {
		t.Fatalf("length = %d, want %d", totalLen, wantLen)
	}
	if seqNo != 0 {
		t.Fatalf("seqno = %d, want 0", seqNo)
	}
	gotCRC := binary.LittleEndian.Uint32(frame[len(frame)-4:])
	wantCRC := crc32.ChecksumIEEE(frame[:len(frame)-4])
	if gotCRC != wantCRC {
		t.Fatalf("crc = %#x, want %#x", gotCRC, wantCRC)
	}
}

func TestCodecFullInvalidCRC(t *testing.T) {
	payload := []byte("abcdefghijkl")
	var wire bytes.Buffer
	codec := &FullCodec{}
	if err := codec.WriteFrame(&wire, payload); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}
	frame := wire.Bytes()
	frame[len(frame)-1] ^= 0xff

	got, err := (&FullCodec{}).ReadFrame(bufio.NewReader(bytes.NewReader(frame)))
	if err == nil {
		t.Fatal("ReadFrame() error is nil")
	}
	if got != nil {
		t.Fatalf("ReadFrame() payload = %x, want nil", got)
	}
	if !errors.Is(err, ErrBadCRC) {
		t.Fatalf("ReadFrame() error = %v, want ErrBadCRC", err)
	}
}

func mustWriteFrame(t *testing.T, codec Codec, payload []byte) []byte {
	t.Helper()
	var buf bytes.Buffer
	if err := codec.WriteFrame(&buf, payload); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}
	return buf.Bytes()
}

var _ io.Reader = (*bufio.Reader)(nil)
