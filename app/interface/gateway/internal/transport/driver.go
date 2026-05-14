package transport

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

var errConnectionClosed = errors.New("gateway transport: connection closed")

type Connection interface {
	ID() string
	RemoteAddr() string
	WriteFrame(ctx context.Context, payload []byte) error
	Close() error
}

type Handler interface {
	OnOpen(ctx context.Context, conn Connection)
	OnFrame(ctx context.Context, conn Connection, frame []byte) error
	OnClose(ctx context.Context, conn Connection, err error)
}

type netDriver struct {
	conn    net.Conn
	handler Handler
}

func newNetDriver(conn net.Conn, handler Handler) *netDriver {
	return &netDriver{conn: conn, handler: handler}
}

func (d *netDriver) Serve(ctx context.Context) error {
	if d == nil || d.conn == nil {
		return nil
	}
	defer d.conn.Close()
	reader := bufio.NewReader(d.conn)
	codec, err := DetectCodec(reader)
	if err != nil {
		return fmt.Errorf("gateway transport detect codec: %w", err)
	}
	conn := newNetConnection(d.conn, codec)
	if d.handler != nil {
		d.handler.OnOpen(ctx, conn)
		defer func() {
			d.handler.OnClose(context.Background(), conn, conn.closeErr())
		}()
	}
	for {
		frame, err := codec.ReadFrame(reader)
		if err != nil {
			if conn.markClosed(err) {
				_ = d.conn.Close()
			}
			return err
		}
		if d.handler == nil {
			continue
		}
		if err := d.handler.OnFrame(ctx, conn, frame); err != nil {
			conn.markClosed(err)
			_ = d.conn.Close()
			return err
		}
		if conn.isClosed() {
			return conn.closeErr()
		}
	}
}

var nextNetConnectionID uint64

type netConnection struct {
	id    string
	conn  net.Conn
	codec Codec

	writeMu sync.Mutex
	closeMu sync.Mutex
	closed  bool
	err     error
}

func newNetConnection(conn net.Conn, codec Codec) *netConnection {
	id := atomic.AddUint64(&nextNetConnectionID, 1)
	return &netConnection{id: fmt.Sprintf("net-%d", id), conn: conn, codec: codec}
}

func (c *netConnection) ID() string {
	if c == nil {
		return ""
	}
	return c.id
}

func (c *netConnection) RemoteAddr() string {
	if c == nil || c.conn == nil || c.conn.RemoteAddr() == nil {
		return ""
	}
	return c.conn.RemoteAddr().String()
}

func (c *netConnection) WriteFrame(ctx context.Context, payload []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if c == nil || c.conn == nil || c.codec == nil {
		return errConnectionClosed
	}
	if c.isClosed() {
		return errConnectionClosed
	}
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	if err := ctx.Err(); err != nil {
		return err
	}
	if c.isClosed() {
		return errConnectionClosed
	}
	if err := c.codec.WriteFrame(c.conn, payload); err != nil {
		c.markClosed(err)
		return err
	}
	return nil
}

func (c *netConnection) Close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	c.markClosed(errConnectionClosed)
	return c.conn.Close()
}

func (c *netConnection) isClosed() bool {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()
	return c.closed
}

func (c *netConnection) closeErr() error {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()
	if c.err == nil {
		return errConnectionClosed
	}
	return c.err
}

func (c *netConnection) markClosed(err error) bool {
	if c == nil {
		return false
	}
	c.closeMu.Lock()
	defer c.closeMu.Unlock()
	if c.closed {
		return false
	}
	if err == nil {
		err = errConnectionClosed
	}
	c.err = err
	c.closed = true
	return true
}
