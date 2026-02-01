// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package ws

import (
	"bytes"
	"io"

	"github.com/panjf2000/gnet/v2/pkg/buffer/elastic"
)

type WsConn struct {
	InboundBuffer elastic.RingBuffer // buffer for leftover data from the peer
	Buffer        []byte             // buffer for the latest bytes
	cache         bytes.Buffer       // temporary buffer for scattered bytes
}

func (c *WsConn) Read(p []byte) (n int, err error) {
	if c.InboundBuffer.IsEmpty() {
		n = copy(p, c.Buffer)
		c.Buffer = c.Buffer[n:]
		if n == 0 && len(p) > 0 {
			err = io.ErrShortBuffer
		}
		return
	}
	n, _ = c.InboundBuffer.Read(p)
	if n == len(p) {
		return
	}
	m := copy(p[n:], c.Buffer)
	n += m
	c.Buffer = c.Buffer[m:]
	return
}

func (c *WsConn) Next(n int) (buf []byte, err error) {
	inBufferLen := c.InboundBuffer.Buffered()
	if totalLen := inBufferLen + len(c.Buffer); n > totalLen {
		return nil, io.ErrShortBuffer
	} else if n <= 0 {
		n = totalLen
	}
	if c.InboundBuffer.IsEmpty() {
		buf = c.Buffer[:n]
		c.Buffer = c.Buffer[n:]
		return
	}
	head, tail := c.InboundBuffer.Peek(n)
	defer c.InboundBuffer.Discard(n) //nolint:errcheck
	if len(head) >= n {
		return head[:n], err
	}
	c.cache.Reset()
	c.cache.Write(head)
	c.cache.Write(tail)
	if inBufferLen >= n {
		return c.cache.Bytes(), err
	}

	remaining := n - inBufferLen
	c.cache.Write(c.Buffer[:remaining])
	c.Buffer = c.Buffer[remaining:]
	return c.cache.Bytes(), err
}

func (c *WsConn) Peek(n int) (buf []byte, err error) {
	inBufferLen := c.InboundBuffer.Buffered()
	if totalLen := inBufferLen + len(c.Buffer); n > totalLen {
		return nil, io.ErrShortBuffer
	} else if n <= 0 {
		n = totalLen
	}
	if c.InboundBuffer.IsEmpty() {
		return c.Buffer[:n], err
	}
	head, tail := c.InboundBuffer.Peek(n)
	if len(head) >= n {
		return head[:n], err
	}
	c.cache.Reset()
	c.cache.Write(head)
	c.cache.Write(tail)
	if inBufferLen >= n {
		return c.cache.Bytes(), err
	}

	remaining := n - inBufferLen
	c.cache.Write(c.Buffer[:remaining])
	return c.cache.Bytes(), err
}

func (c *WsConn) Discard(n int) (int, error) {
	inBufferLen := c.InboundBuffer.Buffered()
	tempBufferLen := len(c.Buffer)
	if inBufferLen+tempBufferLen < n || n <= 0 {
		c.resetBuffer()
		return inBufferLen + tempBufferLen, nil
	}
	if c.InboundBuffer.IsEmpty() {
		c.Buffer = c.Buffer[n:]
		return n, nil
	}

	discarded, _ := c.InboundBuffer.Discard(n)
	if discarded < inBufferLen {
		return discarded, nil
	}

	remaining := n - inBufferLen
	c.Buffer = c.Buffer[remaining:]
	return n, nil
}

func (c *WsConn) resetBuffer() {
	c.Buffer = c.Buffer[:0]
	c.InboundBuffer.Reset()
}

func (c *WsConn) InboundBuffered() int {
	return c.InboundBuffer.Buffered() + len(c.Buffer)
}

func (c *WsConn) Release() {
	c.InboundBuffer.Done()
}
