package transport

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
)

const (
	abridgedFlag           = 0xef
	abridgedInt32Flag      = 0xefefefef
	intermediateFlag       = 0xeeeeeeee
	paddedIntermediateFlag = 0xdddddddd
	httpHeadFlag           = 0x44414548
	httpPostFlag           = 0x54534f50
	httpGetFlag            = 0x20544547
	httpOptionFlag         = 0x4954504f
	maxFrameSize           = 16 * 1024 * 1024
)

var (
	ErrBadLength            = errors.New("mtproto transport: bad frame length")
	ErrBadCRC               = errors.New("mtproto transport: bad crc32")
	ErrBadSeq               = errors.New("mtproto transport: bad sequence number")
	ErrBadMagic             = errors.New("mtproto transport: bad magic")
	ErrUnsupportedTransport = errors.New("mtproto transport: unsupported transport")
)

type Codec interface {
	ReadFrame(r io.Reader) ([]byte, error)
	WriteFrame(w io.Writer, payload []byte) error
}

type AbridgedCodec struct{}

type IntermediateCodec struct{}

type PaddedIntermediateCodec struct{}

type FullCodec struct {
	recvSeqNo uint32
	sendSeqNo uint32
}

type ObfuscatedCodec struct {
	base Codec
	recv *crypto.AesCTR128Encrypt
	send *crypto.AesCTR128Encrypt
}

func DetectCodec(r *bufio.Reader) (Codec, error) {
	first, err := r.Peek(1)
	if err != nil {
		return nil, err
	}
	if first[0] == abridgedFlag {
		if _, err := r.Discard(1); err != nil {
			return nil, err
		}
		return &AbridgedCodec{}, nil
	}

	header, err := r.Peek(4)
	if err != nil {
		return nil, err
	}
	if binary.LittleEndian.Uint32(header) == intermediateFlag {
		if _, err := r.Discard(4); err != nil {
			return nil, err
		}
		return &IntermediateCodec{}, nil
	}
	if binary.LittleEndian.Uint32(header) == paddedIntermediateFlag {
		if _, err := r.Discard(4); err != nil {
			return nil, err
		}
		return &PaddedIntermediateCodec{}, nil
	}
	switch binary.LittleEndian.Uint32(header) {
	case httpHeadFlag, httpPostFlag, httpGetFlag, httpOptionFlag:
		return nil, fmt.Errorf("%w: http", ErrUnsupportedTransport)
	}

	fullHeader, err := r.Peek(12)
	if err != nil {
		return nil, err
	}
	if binary.LittleEndian.Uint32(fullHeader[4:8]) == 0 {
		return &FullCodec{}, nil
	}

	obfuscatedHeader, err := r.Peek(64)
	if err != nil {
		return nil, err
	}
	codec, err := detectObfuscatedCodec(obfuscatedHeader)
	if err != nil {
		return nil, err
	}
	if _, err := r.Discard(64); err != nil {
		return nil, err
	}
	return codec, nil
}

func (AbridgedCodec) ReadFrame(r io.Reader) ([]byte, error) {
	var first [1]byte
	if _, err := io.ReadFull(r, first[:]); err != nil {
		return nil, err
	}

	n := int(first[0] & 0x7f)
	if n == 0x7f {
		var ext [3]byte
		if _, err := io.ReadFull(r, ext[:]); err != nil {
			return nil, err
		}
		n = int(ext[0]) | int(ext[1])<<8 | int(ext[2])<<16
	}
	size := n << 2
	if err := validatePayloadLength(size); err != nil {
		return nil, err
	}

	payload := make([]byte, size)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (AbridgedCodec) WriteFrame(w io.Writer, payload []byte) error {
	if err := validatePayloadLength(len(payload)); err != nil {
		return err
	}

	size := len(payload) >> 2
	var header []byte
	if size < 0x7f {
		header = []byte{byte(size)}
	} else {
		header = []byte{0x7f, byte(size), byte(size >> 8), byte(size >> 16)}
	}
	if _, err := w.Write(header); err != nil {
		return err
	}
	_, err := w.Write(payload)
	return err
}

func (IntermediateCodec) ReadFrame(r io.Reader) ([]byte, error) {
	var header [4]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return nil, err
	}
	size := int(binary.LittleEndian.Uint32(header[:]) & 0x7fffffff)
	if err := validatePayloadLength(size); err != nil {
		return nil, err
	}

	payload := make([]byte, size)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (IntermediateCodec) WriteFrame(w io.Writer, payload []byte) error {
	if err := validatePayloadLength(len(payload)); err != nil {
		return err
	}

	var header [4]byte
	binary.LittleEndian.PutUint32(header[:], uint32(len(payload)))
	if _, err := w.Write(header[:]); err != nil {
		return err
	}
	_, err := w.Write(payload)
	return err
}

func (PaddedIntermediateCodec) ReadFrame(r io.Reader) ([]byte, error) {
	var header [4]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return nil, err
	}
	size := int(binary.LittleEndian.Uint32(header[:]) & 0x7fffffff)
	if err := validatePayloadLength(size); err != nil {
		return nil, err
	}

	payload := make([]byte, size)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (PaddedIntermediateCodec) WriteFrame(w io.Writer, payload []byte) error {
	if err := validatePayloadLength(len(payload)); err != nil {
		return err
	}

	var header [4]byte
	binary.LittleEndian.PutUint32(header[:], uint32(len(payload)))
	if _, err := w.Write(header[:]); err != nil {
		return err
	}
	_, err := w.Write(payload)
	return err
}

func (c *FullCodec) ReadFrame(r io.Reader) ([]byte, error) {
	var lengthBuf [4]byte
	if _, err := io.ReadFull(r, lengthBuf[:]); err != nil {
		return nil, err
	}

	totalLen := int(binary.LittleEndian.Uint32(lengthBuf[:]))
	if totalLen < 12 || totalLen > maxFrameSize || totalLen%4 != 0 {
		return nil, fmt.Errorf("%w: %d", ErrBadLength, totalLen)
	}

	rest := make([]byte, totalLen-4)
	if _, err := io.ReadFull(r, rest); err != nil {
		return nil, err
	}

	payloadEnd := len(rest) - 4
	gotCRC := binary.LittleEndian.Uint32(rest[payloadEnd:])
	hash := crc32.NewIEEE()
	_, _ = hash.Write(lengthBuf[:])
	_, _ = hash.Write(rest[:payloadEnd])
	if wantCRC := hash.Sum32(); gotCRC != wantCRC {
		return nil, fmt.Errorf("%w: got %#x want %#x", ErrBadCRC, gotCRC, wantCRC)
	}

	seqNo := binary.LittleEndian.Uint32(rest[:4])
	if seqNo != c.recvSeqNo {
		return nil, fmt.Errorf("%w: got %d want %d", ErrBadSeq, seqNo, c.recvSeqNo)
	}
	c.recvSeqNo++

	payload := make([]byte, payloadEnd-4)
	copy(payload, rest[4:payloadEnd])
	return payload, nil
}

func (c *FullCodec) WriteFrame(w io.Writer, payload []byte) error {
	if err := validatePayloadLength(len(payload)); err != nil {
		return err
	}

	totalLen := 4 + 4 + len(payload) + 4
	if totalLen > maxFrameSize {
		return fmt.Errorf("%w: %d", ErrBadLength, totalLen)
	}

	frame := make([]byte, totalLen)
	binary.LittleEndian.PutUint32(frame[:4], uint32(totalLen))
	binary.LittleEndian.PutUint32(frame[4:8], c.sendSeqNo)
	c.sendSeqNo++
	copy(frame[8:], payload)
	binary.LittleEndian.PutUint32(frame[totalLen-4:], crc32.ChecksumIEEE(frame[:totalLen-4]))

	_, err := w.Write(frame)
	return err
}

func (c *ObfuscatedCodec) ReadFrame(r io.Reader) ([]byte, error) {
	return c.base.ReadFrame(obfuscatedReader{r: r, cipher: c.recv})
}

func (c *ObfuscatedCodec) WriteFrame(w io.Writer, payload []byte) error {
	var frame bytes.Buffer
	if err := c.base.WriteFrame(&frame, payload); err != nil {
		return err
	}
	encrypted := append([]byte(nil), frame.Bytes()...)
	c.send.Encrypt(encrypted)
	_, err := w.Write(encrypted)
	return err
}

func detectObfuscatedCodec(header []byte) (*ObfuscatedCodec, error) {
	obfuscated := append([]byte(nil), header[:64]...)
	var tmp [48]byte
	for i := range tmp {
		tmp[i] = obfuscated[55-i]
	}

	send, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		return nil, fmt.Errorf("mtproto transport: obfuscated send cipher: %w", err)
	}
	recv, err := crypto.NewAesCTR128Encrypt(obfuscated[8:40], obfuscated[40:56])
	if err != nil {
		return nil, fmt.Errorf("mtproto transport: obfuscated receive cipher: %w", err)
	}

	decryptedHeader := append([]byte(nil), obfuscated...)
	recv.Encrypt(decryptedHeader)
	protocol := binary.BigEndian.Uint32(decryptedHeader[56:60])
	dcID := binary.BigEndian.Uint16(decryptedHeader[60:62])
	if dcID == 0 {
		return nil, fmt.Errorf("%w: obfuscated dc id is zero", ErrBadMagic)
	}

	var base Codec
	switch protocol {
	case abridgedInt32Flag:
		base = &AbridgedCodec{}
	case intermediateFlag:
		base = &IntermediateCodec{}
	case paddedIntermediateFlag:
		base = &PaddedIntermediateCodec{}
	default:
		return nil, fmt.Errorf("%w: obfuscated protocol %#x", ErrBadMagic, protocol)
	}

	return &ObfuscatedCodec{base: base, recv: recv, send: send}, nil
}

type obfuscatedReader struct {
	r      io.Reader
	cipher *crypto.AesCTR128Encrypt
}

func (r obfuscatedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if n > 0 {
		r.cipher.Encrypt(p[:n])
	}
	return n, err
}

func validatePayloadLength(size int) error {
	if size <= 0 || size > maxFrameSize || size%4 != 0 {
		return fmt.Errorf("%w: %d", ErrBadLength, size)
	}
	return nil
}
