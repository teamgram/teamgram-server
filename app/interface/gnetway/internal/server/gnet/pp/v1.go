package pp

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/pkg/errors"
)

// readV1Header assumes the passed buf contains the first 13 bytes which should look like one of
// the following. (Where XX is the start of the tcp address)
//
//	"PROXY TCP4 XX"
//	"PROXY TCP6 XX"
//	"PROXY UNKNOWN"
func readV1Header(buf []byte, r io.Reader) (*Header, error) {
	// For "UNKNOWN", the rest of the line before the cRLF may be omitted by the
	// sender, and the receiver must ignore anything presented before the cRLF is found.
	if bytes.Equal(buf[6:13], []byte(v1UnKnownProto)) {
		b, err := readUntilCRLF(buf, r, 13)
		if err != nil {
			return nil, errors.Wrap(err, "while looking for cRLF after UNKNOWN proto")
		}
		return &Header{IsLocal: true, Version: 1, Unknown: b}, nil
	}

	var idx int
	if bytes.Equal(buf[6:10], []byte("TCP4")) {
		// Minimum TCP4 line is `PROXY TCP4 1.1.1.1 1.1.1.1 2 3\r\n` which is 32 bytes, minus the 13 we have
		// already read which leaves 18, so we optimistically read them now.
		if _, err := io.ReadFull(r, buf[13:32]); err != nil {
			return nil, errors.Wrap(err, "while reading tcp4 addresses")
		}

		// If the optimistic read ended in cRLF then no more bytes to read
		if bytes.Equal(buf[30:32], []byte(cRLF)) {
			return parseV1Header(buf[0:30])
		}
		idx = 32
	}

	if bytes.Equal(buf[6:10], []byte("TCP6")) {
		// Minimum TCP6 line is `PROXY TCP6 ::1 ::1 2 3\r\n` which is 24 bytes, minus the 13 we have
		// already read which leaves 11, so we optimistically read them now.
		if _, err := io.ReadFull(r, buf[13:24]); err != nil {
			return nil, errors.Wrap(err, "while reading tcp6 addresses")
		}

		//fmt.Printf("cRLF: %X\n", buf[22:24])
		// If the optimistic read ended in cRLF then no more bytes to read
		if bytes.Equal(buf[22:24], []byte(cRLF)) {
			return parseV1Header(buf[0:22])
		}
		idx = 24
	}

	if idx == 0 {
		return nil, errors.Errorf("unrecognized protocol '%s'", buf[6:10])
	}

	// else we have more bytes to read until we find the cRLF
	b, err := readUntilCRLF(buf, r, idx)
	if err != nil {
		return nil, errors.Wrap(err, "while looking for cRLF after proto")
	}
	return parseV1Header(b)
}

// readUntilCRLF reads from the reader placing the bytes into `buf` starting at `idx` until
// it finds the terminating cRLF or we exceed 107 bytes which is the max length of the v1
// proxy proto header.
func readUntilCRLF(buf []byte, r io.Reader, idx int) ([]byte, error) {
	// Read until we find the cRLF or we hit our max possible header length
	for idx < 107 {
		c, err := r.Read(buf[idx : idx+1])
		if c != 1 {
			return nil, errors.New("expected to read more bytes, but got none")
		}
		if err != nil {
			return nil, err
		}
		if bytes.Equal(buf[idx-1:idx+1], []byte(cRLF)) {
			return buf[0 : idx-1], nil
		}
		idx++
	}
	return nil, errors.New("gave up after 107 bytes")
}

// parseV1Header parses the provided v1 proxy protocol header in the form
// "PROXY TCP4 1.1.1.1 1.1.1.1 2 3" into it's individual parts
func parseV1Header(buf []byte) (*Header, error) {
	var src, dest net.TCPAddr
	var done bool

	err := split(buf[11:], func(pos int, buf []byte) error {
		switch pos {
		case 0:
			ip := net.ParseIP(string(buf))
			if ip == nil {
				return errors.Errorf("invalid ip '%s' at pos '%d'", buf, pos)
			}
			src.IP = ip
		case 1:
			ip := net.ParseIP(string(buf))
			if ip == nil {
				return errors.Errorf("invalid ip '%s' at pos '%d'", buf, pos)
			}
			dest.IP = ip
		case 2:
			port, err := strconv.Atoi(string(buf))
			if err != nil {
				return errors.Errorf("invalid port '%s' at pos '%d'", buf, pos)
			}
			src.Port = port
		case 3:
			port, err := strconv.Atoi(string(buf))
			if err != nil {
				return errors.Errorf("invalid port '%s' at pos '%d'", buf, pos)
			}
			dest.Port = port
			done = true
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if !done {
		return nil, fmt.Errorf("address line '%s' corrupted", buf[11:])
	}
	return &Header{IsLocal: false, Version: 1, Source: &src, Destination: &dest}, nil
}

// split takes a given byte buffer and splits on a single space ' ' calling the passed `fn` for each
// section of delimited bytes. This code was favored over standard `bytes.Split()` as this code avoids
// unnecessary allocations.
func split(buf []byte, fn func(pos int, buf []byte) error) error {
	var pos int
	for {
		m := bytes.IndexByte(buf, byte(' '))
		if m < 0 {
			break
		}
		if err := fn(pos, buf[:m]); err != nil {
			return err
		}
		pos++
		buf = buf[m+1:]
	}
	if err := fn(pos, buf); err != nil {
		return err
	}
	return nil
}
