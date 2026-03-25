// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package codec

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface/ecode"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/bytedance/gopkg/lang/dirtmake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec"
	"github.com/cloudwego/kitex/pkg/remote/codec/perrors"
)

const (
	// Magic number for ZRPC protocol (ZRPC in ASCII: 0x5A525043)
	zrpcMagicNumber = uint32(0x5A525043)
)

type ZRpcCodec struct {
	printDebugInfo bool
}

func NewZRpcCodec(debug bool) remote.Codec {
	return &ZRpcCodec{printDebugInfo: debug}
}

func (c *ZRpcCodec) Encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	var validData interface{}
	switch message.MessageType() {
	case remote.Exception:
		switch e := message.Data().(type) {
		case *remote.TransError:
			validData = tg.NewRpcError(e)
		case error:
			validData = tg.NewRpcError(e)
		default:
			return errors.New("exception relay must implement error type")
		}
	default:
		validData = message.Data()
	}

	klog.Infof("trans: %v", message.TransInfo().TransStrInfo())

	// Encode payload using TL binary protocol
	x := bin.NewEncoder()
	defer x.End()
	_ = validData.(iface.TLObject).Encode(x, 201)
	payload := x.Bytes()

	if c.printDebugInfo {
		klog.Infof("encoded payload: %s\n", hex.EncodeToString(payload))
	}

	data := &Meta{
		ServiceName: message.RPCInfo().Invocation().ServiceName(),
		MethodName:  message.RPCInfo().Invocation().MethodName(),
		SeqID:       message.RPCInfo().Invocation().SeqID(),
		MsgType:     uint32(message.MessageType()),
		Payload:     payload,
		Metadata:    message.TransInfo().TransStrInfo(),
	}

	// Encode Meta using binary format
	buf, err := c.encodeMeta(data)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary encode, marshal data failed: %w", err))
	}

	// Write magic number (4 bytes)
	magic := make([]byte, 4)
	binary.BigEndian.PutUint32(magic, zrpcMagicNumber)
	_, err = out.WriteBinary(magic)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary encode, write magic failed: %w", err))
	}

	// Write total length (4 bytes)
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(buf)))
	_, err = out.WriteBinary(length)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary encode, write length failed: %w", err))
	}

	// Write data
	_, err = out.WriteBinary(buf)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary encode, write data failed: %w", err))
	}

	return nil
}

func (c *ZRpcCodec) Decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	// Read magic number (4 bytes)
	magic := dirtmake.Bytes(4, 4)
	_, err := in.ReadBinary(magic)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary decode, read magic failed: %w", err))
	}
	magicNum := binary.BigEndian.Uint32(magic)
	if magicNum != zrpcMagicNumber {
		return perrors.NewProtocolError(fmt.Errorf("binary decode, invalid magic number: 0x%08X, expected: 0x%08X", magicNum, zrpcMagicNumber))
	}

	// Read length (4 bytes)
	length := dirtmake.Bytes(4, 4)
	_, err = in.ReadBinary(length)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary decode, read length failed: %w", err))
	}
	l := int(binary.BigEndian.Uint32(length))

	// Read data
	buf := dirtmake.Bytes(l, l)
	_, err = in.ReadBinary(buf)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary decode, read data failed: %w", err))
	}

	// Decode Meta from binary format
	data, err := c.decodeMeta(buf)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary decode, unmarshal Meta data failed: %w", err))
	}

	if err = codec.SetOrCheckSeqID(data.SeqID, message); err != nil {
		return err
	}
	if err = codec.SetOrCheckMethodName(ctx, data.MethodName, message); err != nil {
		return err
	}
	if err = codec.NewDataIfNeeded(data.MethodName, message); err != nil {
		return err
	}

	if c.printDebugInfo {
		klog.Infof("encoded payload: %s\n", hex.EncodeToString(data.Payload))
	}

	if remote.MessageType(data.MsgType) == remote.Exception {
		exception2, err2 := tg.DecodeRpcErrorClazz(bin.NewDecoder(data.Payload))
		if err2 != nil {
			return perrors.NewProtocolError(fmt.Errorf("binary decode, unmarshal Exception payload failed: %w", err2))
		} else if exception, ok := exception2.(*tg.TLRpcError); !ok {
			return perrors.NewProtocolError(fmt.Errorf("binary decode, invalid Exception type: %T", exception2))
		} else {
			return ecode.NewCodeError(exception.Code(), exception.ErrorMessage)
		}
	}

	// Decode payload using TL binary protocol
	d := bin.NewDecoder(data.Payload)
	err = message.Data().(iface.TLObject).Decode(d)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("binary decode, unmarshal payload failed: %w", err))
	}

	klog.Infof("trans: %v", data.Metadata)

	message.TransInfo().PutTransStrInfo(data.Metadata)

	return nil
}

func (c *ZRpcCodec) Name() string {
	return "ZRPC"
}

// encodeMeta encodes Meta struct into binary format
func (c *ZRpcCodec) encodeMeta(meta *Meta) ([]byte, error) {
	// Calculate buffer size
	size := 0
	size += 2 + len(meta.ServiceName) // ServiceName length (2 bytes) + data
	size += 2 + len(meta.MethodName)  // MethodName length (2 bytes) + data
	size += 4                         // SeqID (4 bytes)
	size += 4                         // MsgType (4 bytes)
	size += 4 + len(meta.Payload)     // Payload length (4 bytes) + data
	size += 2                         // Metadata count (2 bytes)
	for k, v := range meta.Metadata {
		size += 2 + len(k) // key length + key
		size += 2 + len(v) // value length + value
	}

	buf := make([]byte, size)
	offset := 0

	// Write ServiceName
	binary.BigEndian.PutUint16(buf[offset:], uint16(len(meta.ServiceName)))
	offset += 2
	copy(buf[offset:], meta.ServiceName)
	offset += len(meta.ServiceName)

	// Write MethodName
	binary.BigEndian.PutUint16(buf[offset:], uint16(len(meta.MethodName)))
	offset += 2
	copy(buf[offset:], meta.MethodName)
	offset += len(meta.MethodName)

	// Write SeqID
	binary.BigEndian.PutUint32(buf[offset:], uint32(meta.SeqID))
	offset += 4

	// Write MsgType
	binary.BigEndian.PutUint32(buf[offset:], meta.MsgType)
	offset += 4

	// Write Payload
	binary.BigEndian.PutUint32(buf[offset:], uint32(len(meta.Payload)))
	offset += 4
	copy(buf[offset:], meta.Payload)
	offset += len(meta.Payload)

	// Write Metadata
	binary.BigEndian.PutUint16(buf[offset:], uint16(len(meta.Metadata)))
	offset += 2
	for k, v := range meta.Metadata {
		// Write key
		binary.BigEndian.PutUint16(buf[offset:], uint16(len(k)))
		offset += 2
		copy(buf[offset:], k)
		offset += len(k)
		// Write value
		binary.BigEndian.PutUint16(buf[offset:], uint16(len(v)))
		offset += 2
		copy(buf[offset:], v)
		offset += len(v)
	}

	return buf, nil
}

// decodeMeta decodes Meta struct from binary format
func (c *ZRpcCodec) decodeMeta(buf []byte) (*Meta, error) {
	if len(buf) < 4 {
		return nil, fmt.Errorf("buffer too small")
	}

	meta := &Meta{}
	offset := 0

	// Read ServiceName
	if offset+2 > len(buf) {
		return nil, fmt.Errorf("invalid ServiceName length")
	}
	serviceNameLen := int(binary.BigEndian.Uint16(buf[offset:]))
	offset += 2
	if offset+serviceNameLen > len(buf) {
		return nil, fmt.Errorf("invalid ServiceName data")
	}
	meta.ServiceName = string(buf[offset : offset+serviceNameLen])
	offset += serviceNameLen

	// Read MethodName
	if offset+2 > len(buf) {
		return nil, fmt.Errorf("invalid MethodName length")
	}
	methodNameLen := int(binary.BigEndian.Uint16(buf[offset:]))
	offset += 2
	if offset+methodNameLen > len(buf) {
		return nil, fmt.Errorf("invalid MethodName data")
	}
	meta.MethodName = string(buf[offset : offset+methodNameLen])
	offset += methodNameLen

	// Read SeqID
	if offset+4 > len(buf) {
		return nil, fmt.Errorf("invalid SeqID")
	}
	meta.SeqID = int32(binary.BigEndian.Uint32(buf[offset:]))
	offset += 4

	// Read MsgType
	if offset+4 > len(buf) {
		return nil, fmt.Errorf("invalid MsgType")
	}
	meta.MsgType = binary.BigEndian.Uint32(buf[offset:])
	offset += 4

	// Read Payload
	if offset+4 > len(buf) {
		return nil, fmt.Errorf("invalid Payload length")
	}
	payloadLen := int(binary.BigEndian.Uint32(buf[offset:]))
	offset += 4
	if offset+payloadLen > len(buf) {
		return nil, fmt.Errorf("invalid Payload data")
	}
	meta.Payload = make([]byte, payloadLen)
	copy(meta.Payload, buf[offset:offset+payloadLen])
	offset += payloadLen

	// Read Metadata
	if offset+2 > len(buf) {
		return nil, fmt.Errorf("invalid Metadata count")
	}
	metadataCount := int(binary.BigEndian.Uint16(buf[offset:]))
	offset += 2
	meta.Metadata = make(map[string]string, metadataCount)
	for i := 0; i < metadataCount; i++ {
		// Read key
		if offset+2 > len(buf) {
			return nil, fmt.Errorf("invalid Metadata key length")
		}
		keyLen := int(binary.BigEndian.Uint16(buf[offset:]))
		offset += 2
		if offset+keyLen > len(buf) {
			return nil, fmt.Errorf("invalid Metadata key data")
		}
		key := string(buf[offset : offset+keyLen])
		offset += keyLen

		// Read value
		if offset+2 > len(buf) {
			return nil, fmt.Errorf("invalid Metadata value length")
		}
		valueLen := int(binary.BigEndian.Uint16(buf[offset:]))
		offset += 2
		if offset+valueLen > len(buf) {
			return nil, fmt.Errorf("invalid Metadata value data")
		}
		value := string(buf[offset : offset+valueLen])
		offset += valueLen

		meta.Metadata[key] = value
	}

	return meta, nil
}
