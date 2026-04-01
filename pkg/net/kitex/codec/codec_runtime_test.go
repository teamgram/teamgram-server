package codec

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
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
