package codec

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface/ecode"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type testTLObject struct {
	seenLayer int32
	encodeErr error
}

func (t *testTLObject) Encode(x *bin.Encoder, layer int32) error {
	t.seenLayer = layer
	if t.encodeErr != nil {
		return t.encodeErr
	}

	x.PutInt32(7)
	return nil
}

func (t *testTLObject) Decode(d *bin.Decoder) error {
	_, err := d.Int32()
	return err
}

type testByteBuffer struct {
	readBuf  []byte
	writeBuf []byte
	readOff  int
}

func (b *testByteBuffer) Read(p []byte) (int, error) {
	if b.readOff >= len(b.readBuf) {
		return 0, io.EOF
	}
	n := copy(p, b.readBuf[b.readOff:])
	b.readOff += n
	return n, nil
}

func (b *testByteBuffer) Write(p []byte) (int, error) {
	b.writeBuf = append(b.writeBuf, p...)
	return len(p), nil
}

func (b *testByteBuffer) Next(n int) ([]byte, error) {
	if b.readOff+n > len(b.readBuf) {
		return nil, io.EOF
	}
	buf := b.readBuf[b.readOff : b.readOff+n]
	b.readOff += n
	return buf, nil
}

func (b *testByteBuffer) Peek(n int) ([]byte, error) {
	if b.readOff+n > len(b.readBuf) {
		return nil, io.EOF
	}
	return b.readBuf[b.readOff : b.readOff+n], nil
}

func (b *testByteBuffer) Skip(n int) error {
	if b.readOff+n > len(b.readBuf) {
		return io.EOF
	}
	b.readOff += n
	return nil
}

func (b *testByteBuffer) Release(error) error { return nil }

func (b *testByteBuffer) ReadableLen() int { return len(b.readBuf) - b.readOff }

func (b *testByteBuffer) ReadLen() int { return b.readOff }

func (b *testByteBuffer) ReadString(n int) (string, error) {
	buf, err := b.Next(n)
	return string(buf), err
}

func (b *testByteBuffer) ReadBinary(p []byte) (int, error) {
	if b.readOff+len(p) > len(b.readBuf) {
		return 0, io.EOF
	}
	copy(p, b.readBuf[b.readOff:b.readOff+len(p)])
	b.readOff += len(p)
	return len(p), nil
}

func (b *testByteBuffer) Malloc(n int) ([]byte, error) {
	start := len(b.writeBuf)
	b.writeBuf = append(b.writeBuf, make([]byte, n)...)
	return b.writeBuf[start:], nil
}

func (b *testByteBuffer) WrittenLen() int { return len(b.writeBuf) }

func (b *testByteBuffer) WriteString(s string) (int, error) {
	return b.WriteBinary([]byte(s))
}

func (b *testByteBuffer) WriteBinary(p []byte) (int, error) {
	b.writeBuf = append(b.writeBuf, p...)
	return len(p), nil
}

func (b *testByteBuffer) Flush() error { return nil }

func (b *testByteBuffer) NewBuffer() remote.ByteBuffer { return &testByteBuffer{} }

func (b *testByteBuffer) AppendBuffer(buf remote.ByteBuffer) error {
	data, err := buf.Bytes()
	if err != nil {
		return err
	}
	b.writeBuf = append(b.writeBuf, data...)
	return nil
}

func (b *testByteBuffer) Bytes() ([]byte, error) {
	return b.writeBuf, nil
}

func newTestMessage(data iface.TLObject) remote.Message {
	ri := rpcinfo.NewRPCInfo(nil, nil, rpcinfo.NewInvocation("svc.test", "TestMethod"), nil, nil)
	return remote.NewMessage(data, ri, remote.Call, remote.Client)
}

func newTestExceptionMessage(data interface{}) remote.Message {
	ri := rpcinfo.NewRPCInfo(nil, nil, rpcinfo.NewInvocation("svc.test", "TestMethod"), nil, nil)
	return remote.NewMessage(data, ri, remote.Exception, remote.Client)
}

func TestEncodePropagatesEncodeError(t *testing.T) {
	c := NewZRpcCodec(false)
	obj := &testTLObject{encodeErr: errors.New("boom")}

	err := c.Encode(context.Background(), newTestMessage(obj), &testByteBuffer{})
	if err == nil {
		t.Fatal("expected encode error")
	}
}

func TestEncodeUsesDefaultTLLayer(t *testing.T) {
	c := NewZRpcCodec(false)
	obj := &testTLObject{}

	if err := c.Encode(context.Background(), newTestMessage(obj), &testByteBuffer{}); err != nil {
		t.Fatalf("encode failed: %v", err)
	}
	if obj.seenLayer != defaultTLLayer {
		t.Fatalf("expected layer %d, got %d", defaultTLLayer, obj.seenLayer)
	}
}

func TestEncodeUsesInternalTLLayerForServiceSchema(t *testing.T) {
	c := NewZRpcCodec(false)
	obj := dialogpb.MakeTLDialogPeer(&dialogpb.TLDialogPeer{PeerType: 1, PeerId: 777000})

	if err := c.Encode(context.Background(), newTestMessage(obj), &testByteBuffer{}); err != nil {
		t.Fatalf("encode failed: %v", err)
	}
}

func TestEncodeUsesDefaultTLLayerForTGSchema(t *testing.T) {
	c := NewZRpcCodec(false)
	obj := tg.MakeTLDialogPeer(&tg.TLDialogPeer{
		Peer: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 777000}),
	})

	if err := c.Encode(context.Background(), newTestMessage(obj), &testByteBuffer{}); err != nil {
		t.Fatalf("encode failed: %v", err)
	}
}

func TestInternalTLLayerIncludesGeneratedDialogPeerServiceArgs(t *testing.T) {
	if !usesInternalTLLayerForType(
		"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog/dialogservice",
		"GetPeerDialogsV2Args",
	) {
		t.Fatal("expected generated dialog peer args package to use internal TL layer")
	}
	if usesInternalTLLayerForType(
		"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/dialogs/dialogsservice",
		"GetPeerDialogsArgs",
	) {
		t.Fatal("expected bff service args package to use default TG TL layer")
	}
	if usesInternalTLLayerForType(
		"github.com/teamgram/teamgram-server/v2/app/service/media/media/mediaservice",
		"UploadPhotoFileArgs",
	) {
		t.Fatal("expected media service args with nested TG schema to use default TG TL layer")
	}
}

func TestRawTLObjectRoundTripPreservesPayload(t *testing.T) {
	c := NewZRpcCodec(false)
	payload := []byte{0xe2, 0x06, 0x05, 0x9f, 0x04, 0x00, 0x00, 0x00}
	buf := &testByteBuffer{}

	if err := c.Encode(context.Background(), newTestMessage(NewRawTLObject(payload)), buf); err != nil {
		t.Fatalf("encode failed: %v", err)
	}

	out := NewRawTLObject(nil)
	if err := c.Decode(context.Background(), newTestMessage(out), &testByteBuffer{readBuf: buf.writeBuf}); err != nil {
		t.Fatalf("decode failed: %v", err)
	}

	if string(out.Payload) != string(payload) {
		t.Fatalf("decoded payload = %x, want %x", out.Payload, payload)
	}
}

func TestDecodeRejectsOversizedFrame(t *testing.T) {
	c := NewZRpcCodec(false)
	frame := make([]byte, 8)
	binary.BigEndian.PutUint32(frame[:4], zrpcMagicNumber)
	binary.BigEndian.PutUint32(frame[4:], uint32(defaultMaxFrameSize+1))

	err := c.Decode(context.Background(), newTestMessage(&testTLObject{}), &testByteBuffer{readBuf: frame})
	if err == nil {
		t.Fatal("expected oversized frame error")
	}
}

func TestExceptionRoundTripPreservesPlainErrorWithoutTGCode(t *testing.T) {
	c := NewZRpcCodec(false)
	buf := &testByteBuffer{}
	src := errors.New("plain internal error")

	if err := c.Encode(context.Background(), newTestExceptionMessage(src), buf); err != nil {
		t.Fatalf("encode failed: %v", err)
	}

	decodeBuf := &testByteBuffer{readBuf: buf.writeBuf}
	err := c.Decode(context.Background(), newTestMessage(&testTLObject{}), decodeBuf)
	if err == nil {
		t.Fatal("expected decode to return exception error")
	}
	var codeErr ecode.CodeError
	if errors.As(err, &codeErr) {
		t.Fatalf("expected plain error, got ecode %d %q", codeErr.Code(), codeErr.Msg())
	}
	if err.Error() != src.Error() {
		t.Fatalf("expected %q, got %q", src.Error(), err.Error())
	}
}

func TestExceptionRoundTripPreservesCodeError(t *testing.T) {
	c := NewZRpcCodec(false)
	buf := &testByteBuffer{}
	src := ecode.NewCodeError(403, "USER_PRIVACY_RESTRICTED")

	if err := c.Encode(context.Background(), newTestExceptionMessage(src), buf); err != nil {
		t.Fatalf("encode failed: %v", err)
	}

	decodeBuf := &testByteBuffer{readBuf: buf.writeBuf}
	err := c.Decode(context.Background(), newTestMessage(&testTLObject{}), decodeBuf)
	if err == nil {
		t.Fatal("expected decode to return exception error")
	}
	var codeErr ecode.CodeError
	if !errors.As(err, &codeErr) {
		t.Fatalf("expected ecode.CodeError, got %T: %v", err, err)
	}
	if codeErr.Code() != src.Code() || codeErr.Msg() != src.Msg() {
		t.Fatalf("expected code=%d msg=%q, got code=%d msg=%q", src.Code(), src.Msg(), codeErr.Code(), codeErr.Msg())
	}
}
