package transport

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
)

const (
	abridgedFlag     = 0xef
	intermediateFlag = 0xeeeeeeee
	maxFrameSize     = 16 * 1024 * 1024
)

var (
	ErrBadLength = errors.New("mtproto transport: bad frame length")
	ErrBadCRC    = errors.New("mtproto transport: bad crc32")
	ErrBadSeq    = errors.New("mtproto transport: bad sequence number")
)

type Codec interface {
	ReadFrame(r io.Reader) ([]byte, error)
	WriteFrame(w io.Writer, payload []byte) error
}

type AbridgedCodec struct{}

type IntermediateCodec struct{}

type FullCodec struct {
	recvSeqNo uint32
	sendSeqNo uint32
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

	return &FullCodec{}, nil
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

func validatePayloadLength(size int) error {
	if size <= 0 || size > maxFrameSize || size%4 != 0 {
		return fmt.Errorf("%w: %d", ErrBadLength, size)
	}
	return nil
}
