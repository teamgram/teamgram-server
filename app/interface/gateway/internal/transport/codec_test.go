package transport

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
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

func TestCodecObfuscatedAbridgedRoundTrip(t *testing.T) {
	payload := []byte("abcdefghijkl")
	wire, serverToClient := obfuscatedClientWire(t, 0xefefefef, mustWriteFrame(t, AbridgedCodec{}, payload))

	r := bufio.NewReader(bytes.NewReader(wire))
	codec, err := DetectCodec(r)
	if err != nil {
		t.Fatalf("DetectCodec() error = %v", err)
	}
	got, err := codec.ReadFrame(r)
	if err != nil {
		t.Fatalf("ReadFrame() error = %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("ReadFrame() = %x, want %x", got, payload)
	}

	response := []byte("mnopqrstuvwx")
	var encrypted bytes.Buffer
	if err := codec.WriteFrame(&encrypted, response); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}
	decrypted := append([]byte(nil), encrypted.Bytes()...)
	serverToClient.Encrypt(decrypted)
	gotResponse, err := AbridgedCodec{}.ReadFrame(bytes.NewReader(decrypted))
	if err != nil {
		t.Fatalf("client ReadFrame(response) error = %v", err)
	}
	if !bytes.Equal(gotResponse, response) {
		t.Fatalf("response = %x, want %x", gotResponse, response)
	}
}

func TestCodecObfuscatedAbridgedUsesForwardKeyForClientFrames(t *testing.T) {
	payload := []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		'a', 'b', 'c', 'd',
	}
	wire, _ := obfuscatedClientWireWithMasterDirections(t, abridgedInt32Flag, mustWriteFrame(t, AbridgedCodec{}, payload))

	r := bufio.NewReader(bytes.NewReader(wire))
	codec, err := DetectCodec(r)
	if err != nil {
		t.Fatalf("DetectCodec() error = %v", err)
	}
	got, err := codec.ReadFrame(r)
	if err != nil {
		t.Fatalf("ReadFrame() error = %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("ReadFrame() = %x, want %x", got, payload)
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

func obfuscatedClientWire(t *testing.T, protocol uint32, encryptedPayload []byte) ([]byte, *crypto.AesCTR128Encrypt) {
	t.Helper()
	return obfuscatedClientWireWithMasterDirections(t, protocol, encryptedPayload)
}

func obfuscatedClientWireWithMasterDirections(t *testing.T, protocol uint32, framedPayload []byte) ([]byte, *crypto.AesCTR128Encrypt) {
	t.Helper()
	header := make([]byte, 64)
	for i := range header {
		header[i] = byte(i*5 + 11)
	}
	binary.BigEndian.PutUint32(header[56:60], protocol)
	binary.BigEndian.PutUint16(header[60:62], 1)

	clientToServer, err := crypto.NewAesCTR128Encrypt(header[8:40], header[40:56])
	if err != nil {
		t.Fatalf("client-to-server cipher: %v", err)
	}
	wireHeader := append([]byte(nil), header...)
	clientToServer.Encrypt(wireHeader)
	copy(wireHeader[:56], header[:56])

	var tmp [48]byte
	for i := 0; i < len(tmp); i++ {
		tmp[i] = header[55-i]
	}
	serverToClient, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		t.Fatalf("server-to-client cipher: %v", err)
	}
	wirePayload := append([]byte(nil), framedPayload...)
	clientToServer.Encrypt(wirePayload)

	wire := append(wireHeader, wirePayload...)
	return wire, serverToClient
}

var _ io.Reader = (*bufio.Reader)(nil)
