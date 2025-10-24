package pp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/pkg/errors"
)

const (
	ipv4AddressLen = 12
	ipv6AddressLen = 36
)

// ReadV2Header assumes the first read will contain the identifier, then reads up until the
// length of the header as specified by the proxy protocol header. If you are using a
// bufio.Reader you can peek at the first 13 bytes to ensure the header identifier exists
// before passing the bufio.Reader to this function.
func ReadV2Header(r io.Reader) (*Header, error) {
	var buf [232]byte
	if _, err := io.ReadFull(r, buf[0:13]); err != nil {
		return nil, errors.Wrap(err, "while reading proxy proto identifier")
	}
	return readV2Header(buf[0:], r)
}

// readV2Header assumes the passed buf contains the first 13 bytes which should look like
// the following. (Where X is the proto proxy version and command)
//
//	"\r\n\r\n\x00\r\nQUIT\nX"
func readV2Header(buf []byte, r io.Reader) (*Header, error) {
	// Read the next 3 bytes which contain the proto, family and the length of the trailing header
	if _, err := io.ReadFull(r, buf[13:16]); err != nil {
		return nil, errors.Wrap(err, "while reading proto, family and length bytes")
	}

	// Ensure the version is 2
	if (buf[12] & 0xF0) != 0x20 {
		return nil, fmt.Errorf("unexpected version number '%X' at pos '13'", buf[12]&0xF0)
	}

	// The length of the remainder of the header including any TLVs in network byte order
	length := binary.BigEndian.Uint16(buf[14:16])

	// The HA proxy implementation does not limit the length of the proxy protocol header plus TLVs as
	// the proxy protocol is supposed to be used between trusted parties. I feel this is an oversight
	// and impose a generous limit of 2k here to account for any future tlv data use.
	if length > 2048 {
		return nil, fmt.Errorf("header lengh of '%d' is greater than the allowed 2048 bytes", length)
	}

	var tr []byte
	if length != 0 {
		// Read the remainder of the header
		tr = make([]byte, length)
		if _, err := io.ReadFull(r, tr); err != nil {
			return nil, errors.Wrap(err, "while reading proto and length bytes")
		}
	}

	var offset int
	h := Header{Version: 2}

	switch buf[12] & 0x0F {
	case 0x00: // LOCAL command
		h.IsLocal = true
		if tr == nil {
			return &h, nil
		}
	case 0x01: // PROXY command
		if tr == nil {
			return nil, errors.New("expected address but got zero length header")
		}

		// Translate the addresses according to the family
		switch buf[13] {
		case 0x11, 0x12: // IPV4 (TCP/UDP)
			if len(tr) < ipv4AddressLen {
				return nil, fmt.Errorf("expected %d bytes for IPV4 address", ipv4AddressLen)
			}

			var src, dest net.TCPAddr
			src.IP = net.IPv4(tr[0], tr[1], tr[2], tr[3])
			dest.IP = net.IPv4(tr[4], tr[5], tr[6], tr[7])
			src.Port = int(binary.BigEndian.Uint16(tr[8:10]))
			dest.Port = int(binary.BigEndian.Uint16(tr[10:12]))

			if (buf[13] & 0x0F) == 0x02 { // UDP
				h.Destination = &net.UDPAddr{IP: dest.IP, Port: dest.Port}
				h.Source = &net.UDPAddr{IP: src.IP, Port: src.Port}
			} else { // TCP
				h.Destination = &dest
				h.Source = &src
			}
			offset = ipv4AddressLen

		case 0x21, 0x22: // IPV6 (TCP/UDP)
			if len(tr) < ipv6AddressLen {
				return nil, fmt.Errorf("expected %d bytes for IPV6 address", ipv6AddressLen)
			}

			var src, dest net.TCPAddr
			src.IP = tr[0:16]
			dest.IP = tr[16:32]
			src.Port = int(binary.BigEndian.Uint16(tr[32:34]))
			dest.Port = int(binary.BigEndian.Uint16(tr[34:36]))
			if (buf[13] & 0x0F) == 0x02 { // UDP
				h.Destination = &net.UDPAddr{IP: dest.IP, Port: dest.Port}
				h.Source = &net.UDPAddr{IP: src.IP, Port: src.Port}
			} else { // TCP
				h.Destination = &dest
				h.Source = &src
			}
			offset = ipv6AddressLen

		case 0x31, 0x32: // UNIX (STREAM/DGRAM)
			// Not implemented by haproxy and I see no need to implement it here, patches welcome!
			return &h, errors.New("Received UNIX socket proxy command, Currently not supported")
		}
	}

	// If there is trailing data, it should be TLVs
	if offset != len(tr) {
		h.RawTLVs = tr[offset:]
	}
	return &h, nil
}
