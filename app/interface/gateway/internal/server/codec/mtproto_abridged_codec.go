// Copyright (c) 2019-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package codec

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/teamgram/proto/mtproto"

	log "github.com/zeromicro/go-zero/core/logx"
)

// https://core.telegram.org/mtproto#tcp-transport
//
// There is an abridged version of the same protocol:
// if the client sends 0xef as the first byte (**important:** only prior to the very first data packet),
// then packet length is encoded by a single byte (0x01..0x7e = data length divided by 4;
// or 0x7f followed by 3 length bytes (little endian) divided by 4) followed
// by the data themselves (sequence number and CRC32 not added).
// In this case, server responses look the same (the server does not send 0xefas the first byte).
//
//type MTProtoAbridgedCodec struct {
//	conn *net2.BufferedConn
//}
//
//func NewMTProtoAbridgedCodec(conn *net2.BufferedConn) *MTProtoAbridgedCodec {
//	return &MTProtoAbridgedCodec{
//		conn: conn,
//	}
//}
//
//func (c *MTProtoAbridgedCodec) Receive() (interface{}, error) {
//	// minus padding
//	//size := len(x.buf) / 4 - 1
//	//
//	//if size < 127 {
//	//	x.buf[3] = byte(size)
//	//	x.buf = x.buf[3:]
//	//} else {
//	//	binary.LittleEndian.PutUint32(x.buf, uint32(size << 8 | 127))
//	//}
//	//_, err := m.conn.Write(x.buf)
//	//if err != nil {
//	//	return err
//	//}
//
//	var size int
//	var n int
//	var err error
//
//	b := make([]byte, 1)
//	n, err = io.ReadFull(c.conn, b)
//	if err != nil {
//		return nil, err
//	}
//
//	// log.Info("first_byte: ", hex.EncodeToString(b[:1]))
//	needAck := bool(b[0]>>7 == 1)
//	_ = needAck
//
//	b[0] = b[0] & 0x7f
//	// log.Info("first_byte2: ", hex.EncodeToString(b[:1]))
//
//	if b[0] < 0x7f {
//		size = int(b[0]) << 2
//		log.Infof("size1: %d", size)
//		if size == 0 {
//			return nil, nil
//		}
//	} else {
//		log.Infof("first_byte2: %s", hex.EncodeToString(b[:1]))
//		b2 := make([]byte, 3)
//		n, err = io.ReadFull(c.conn, b2)
//		if err != nil {
//			return nil, err
//		}
//		size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2
//		log.Infof("size2: %s", size)
//	}
//
//	left := size
//	buf := make([]byte, size)
//	for left > 0 {
//		n, err = io.ReadFull(c.conn, buf[size-left:])
//		if err != nil {
//			log.Errorf("readFull2 error: %v", err)
//			return nil, err
//		}
//		left -= n
//	}
//	if size > 10240 {
//		log.Infof("readFull2: %s", hex.EncodeToString(buf[:256]))
//	}
//
//	// TODO(@benqi): process report ack and quickack
//	// 截断QuickAck消息，客户端有问题
//	if size == 4 {
//		log.Errorf("server response error: ", int32(binary.LittleEndian.Uint32(buf)))
//		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
//		return nil, nil
//	}
//
//	authKeyId := int64(binary.LittleEndian.Uint64(buf))
//	message := NewMTPRawMessage(authKeyId, 0, TRANSPORT_TCP)
//	message.Decode(buf)
//	return message, nil
//}
//
//func (c *MTProtoAbridgedCodec) Send(msg interface{}) error {
//	message, ok := msg.(*MTPRawMessage)
//	if !ok {
//		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
//		log.Error(err.Error())
//		return err
//	}
//
//	b := message.Encode()
//
//	sb := make([]byte, 4)
//	// minus padding
//	size := len(b) / 4
//
//	if size < 127 {
//		sb = []byte{byte(size)}
//	} else {
//		binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
//	}
//
//	b = append(sb, b...)
//	_, err := c.conn.Write(b)
//
//	if err != nil {
//		log.Errorf("Send msg error: %s", err)
//	}
//
//	return err
//}
//
//func (c *MTProtoAbridgedCodec) Close() error {
//	return c.conn.Close()
//}

type AbridgedCodec struct {
	conn io.ReadWriteCloser
}

func NewMTProtoAbridgedCodec(conn io.ReadWriteCloser) *AbridgedCodec {
	return &AbridgedCodec{
		conn: conn,
	}
}

func (c *AbridgedCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 1)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	// log.Info("first_byte: ", hex.EncodeToString(b[:1]))
	needAck := b[0]>>7 == 1
	_ = needAck

	b[0] = b[0] & 0x7f

	if b[0] < 0x7f {
		size = int(b[0]) << 2
		log.Infof("size1: %d", size)
		if size == 0 {
			return nil, nil
		}
	} else {
		log.Infof("first_byte2: %s", hex.EncodeToString(b[:1]))
		b2 := make([]byte, 3)
		n, err = io.ReadFull(c.conn, b2)
		if err != nil {
			return nil, err
		}
		size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2
		log.Infof("size2: %d", size)
	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			log.Errorf("readFull2 error: %v", err)
			return nil, err
		}
		left -= n
	}
	if size > 4096 {
		log.Infof("readFull2: %s", hex.EncodeToString(buf[:256]))
	}

	// TODO(@benqi): process report ack and quickack
	// 截断QuickAck消息，客户端有问题
	if size == 4 {
		log.Errorf("server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := mtproto.NewMTPRawMessage(authKeyId, 0, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *AbridgedCodec) Send(msg interface{}) error {
	message, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	b := message.Encode()

	sb := make([]byte, 4)
	// minus padding
	size := len(b) / 4

	if size < 127 {
		sb = []byte{byte(size)}
	} else {
		binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
	}

	b = append(sb, b...)
	_, err := c.conn.Write(b)

	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *AbridgedCodec) Close() error {
	return c.conn.Close()
}
