package pp_test

import (
	"net"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/pp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseV1Header(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		src     net.Addr
		dest    net.Addr
		unknown []byte
		err     string
	}{
		{
			name:   "TCP4 Minimal",
			header: "PROXY TCP4 1.1.1.5 1.1.1.6 2 3\r\n",
			src:    &net.TCPAddr{IP: net.ParseIP("1.1.1.5"), Port: 2},
			dest:   &net.TCPAddr{IP: net.ParseIP("1.1.1.6"), Port: 3},
		},
		{
			name:   "TCP4 Maximal",
			header: "PROXY TCP4 255.255.255.255 255.255.255.254 65535 65535\r\n",
			src:    &net.TCPAddr{IP: net.ParseIP("255.255.255.255"), Port: 65535},
			dest:   &net.TCPAddr{IP: net.ParseIP("255.255.255.254"), Port: 65535},
		},
		{
			name:   "TCP6 Minimal",
			header: "PROXY TCP6 ::1 ::2 3 4\r\n",
			src:    &net.TCPAddr{IP: net.ParseIP("::1"), Port: 3},
			dest:   &net.TCPAddr{IP: net.ParseIP("::2"), Port: 4},
		},
		{
			name:   "TCP6 Maximal",
			header: "PROXY TCP6 0000:0000:0000:0000:0000:0000:0000:0002 0000:0000:0000:0000:0000:0000:0000:0001 65535 65535\r\n",
			src:    &net.TCPAddr{IP: net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0002"), Port: 65535},
			dest:   &net.TCPAddr{IP: net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0001"), Port: 65535},
		},
		{
			name:    "UNKNOWN Minimal",
			header:  "PROXY UNKNOWN\r\n",
			unknown: []byte("PROXY UNKNOWN"),
		},
		{
			name:    "UNKNOWN Maximal",
			header:  "PROXY UNKNOWN 0000:0000:0000:0000:0000:0000:0000:0002 0000:0000:0000:0000:0000:0000:0000:0001 65535 65535\r\n",
			unknown: []byte("PROXY UNKNOWN 0000:0000:0000:0000:0000:0000:0000:0002 0000:0000:0000:0000:0000:0000:0000:0001 65535 65535"),
		},
		{
			name:   "TCP6 Empty",
			header: "PROXY TCP6\r\n",
			src:    &net.TCPAddr{IP: net.ParseIP("::1"), Port: 3},
			dest:   &net.TCPAddr{IP: net.ParseIP("::2"), Port: 4},
			err:    "while reading proxy proto identifier: unexpected EOF",
		},
		{
			name:   "TCP6 cRLF Not found",
			header: "PROXY TCP6 0000:0000:0000:0000:0000:0000:0000:0001 0000:0000:0000:0000:0000:0000:0000:0001 65535 65535XXXX\r\n",
			err:    "while parsing proxy proto v1 header: while looking for cRLF after proto: gave up after 107 bytes",
		},
		{
			name:   "UNKNOWN cRLF Not found",
			header: "PROXY UNKNOWN 0000:0000:0000:0000:0000:0000:0000:0001 0000:0000:0000:0000:0000:0000:0000:0001 65535 65535X\r\n",
			err:    "while parsing proxy proto v1 header: while looking for cRLF after UNKNOWN proto: gave up after 107 bytes",
		},
		{
			name:   "UNKNOWN No cRLF",
			header: "PROXY UNKNOWN",
			err:    "while parsing proxy proto v1 header: while looking for cRLF after UNKNOWN proto: expected to read more bytes, but got none",
		},
		{
			name:   "Only cRLF Header",
			header: "\r\n",
			err:    "while reading proxy proto identifier: unexpected EOF",
		},
		{
			name:   "Empty Header",
			header: "",
			err:    "while reading proxy proto identifier: EOF",
		},
		{
			name:   "Garbage Header",
			header: "ASDFASDGSAG@#!@#$!WDFGASDGASDFG#TAGASDFASDG@",
			err:    "expected proxy protocol; found '00000000  41 53 44 46 41 53 44 47  53 41 47 40 23 00        |ASDFASDGSAG@#.|\n' instead",
		},
		{
			name:   "TCP4 Incomplete",
			header: "PROXY TCP4 garbage\r\n",
			err:    "while parsing proxy proto v1 header: while reading tcp4 addresses: unexpected EOF",
		},
		{
			name:   "TCP6 Incomplete",
			header: "PROXY TCP6 garbage\r\n",
			err:    "while parsing proxy proto v1 header: while reading tcp6 addresses: unexpected EOF",
		},
		{
			name:   "unrecognized proto",
			header: "PROXY UNIX :1 :1 234 234\r\n",
			err:    "while parsing proxy proto v1 header: unrecognized protocol 'UNIX'",
		},
		{
			name:   "Invalid src ip",
			header: "PROXY TCP4 NOT-AN-IP 192.168.1.1 22 2345\r\n",
			err:    "while parsing proxy proto v1 header: invalid ip 'NOT-AN-IP' at pos '0'",
		},
		{
			name:   "Invalid dest ip",
			header: "PROXY TCP4 192.168.1.1 NOT-AN-IP 22 2345\r\n",
			err:    "while parsing proxy proto v1 header: invalid ip 'NOT-AN-IP' at pos '1'",
		},
		{
			name:   "Invalid src port",
			header: "PROXY TCP4 192.168.1.1 192.168.1.1 NOT-A-PORT 2345\r\n",
			err:    "while parsing proxy proto v1 header: invalid port 'NOT-A-PORT' at pos '2'",
		},
		{
			name:   "Invalid dest port",
			header: "PROXY TCP4 192.168.1.1 192.168.1.1 22 NOT-A-PORT\r\n",
			err:    "while parsing proxy proto v1 header: invalid port 'NOT-A-PORT' at pos '3'",
		},
		{
			name:   "Corrupted address line",
			header: "PROXY TCP4 192.168.1.1 192.168.1.1 2345\r\n",
			err:    "while parsing proxy proto v1 header: address line '192.168.1.1 192.168.1.1 2345' corrupted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.header)
			h, err := pp.ReadHeader(r)
			if err != nil {
				require.Equal(t, err.Error(), tt.err)
				return
			}
			require.NotNil(t, h)
			assert.Equal(t, tt.dest, h.Destination)
			assert.Equal(t, tt.src, h.Source)
			assert.Equal(t, tt.unknown, h.Unknown)
			assert.Equal(t, 1, h.Version)
		})
	}
}

func BenchmarkReadHeaderV1(b *testing.B) {

	tests := []struct {
		name   string
		header string
	}{
		{
			name:   "TCP4-Minimal",
			header: "PROXY TCP4 1.1.1.5 1.1.1.6 2 3\r\n",
		},
		{
			name:   "TCP4-Maximal",
			header: "PROXY TCP4 255.255.255.255 255.255.255.254 65535 65535\r\n",
		},
		{
			name:   "TCP4-Typical",
			header: "PROXY TCP4 67.11.173.63 67.11.173.63 65535 65535\r\n",
		},
		{
			name:   "TCP6-Minimal",
			header: "PROXY TCP6 ::1 ::2 3 4\r\n",
		},
		{
			name:   "TCP6-Maximal",
			header: "PROXY TCP6 0000:0000:0000:0000:0000:0000:0000:0002 0000:0000:0000:0000:0000:0000:0000:0001 65535 65535\r\n",
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				r := strings.NewReader(tt.header)
				_, err := pp.ReadHeader(r)
				if err != nil {
					b.Errorf("ReadHeader err: %s", err)
				}
			}
		})
	}
}
