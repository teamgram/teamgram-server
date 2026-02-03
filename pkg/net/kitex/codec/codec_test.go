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
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeMeta(t *testing.T) {
	codec := &ZRpcCodec{printDebugInfo: false}

	meta := &Meta{
		ServiceName: "test.service",
		MethodName:  "TestMethod",
		SeqID:       12345,
		MsgType:     1,
		Payload:     []byte{0x01, 0x02, 0x03, 0x04},
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Encode
	buf, err := codec.encodeMeta(meta)
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	// Verify magic number is NOT in encoded meta (it's added in Encode, not encodeMeta)
	// Verify the buffer structure manually
	offset := 0

	// ServiceName length
	serviceNameLen := binary.BigEndian.Uint16(buf[offset:])
	assert.Equal(t, uint16(len(meta.ServiceName)), serviceNameLen)
	offset += 2

	// ServiceName
	serviceName := string(buf[offset : offset+int(serviceNameLen)])
	assert.Equal(t, meta.ServiceName, serviceName)
	offset += int(serviceNameLen)

	// MethodName length
	methodNameLen := binary.BigEndian.Uint16(buf[offset:])
	assert.Equal(t, uint16(len(meta.MethodName)), methodNameLen)
	offset += 2

	// MethodName
	methodName := string(buf[offset : offset+int(methodNameLen)])
	assert.Equal(t, meta.MethodName, methodName)
	offset += int(methodNameLen)

	// SeqID
	seqID := int32(binary.BigEndian.Uint32(buf[offset:]))
	assert.Equal(t, meta.SeqID, seqID)
	offset += 4

	// MsgType
	msgType := binary.BigEndian.Uint32(buf[offset:])
	assert.Equal(t, meta.MsgType, msgType)
	offset += 4

	// Payload length
	payloadLen := binary.BigEndian.Uint32(buf[offset:])
	assert.Equal(t, uint32(len(meta.Payload)), payloadLen)
	offset += 4

	// Payload
	payload := buf[offset : offset+int(payloadLen)]
	assert.Equal(t, meta.Payload, payload)
	offset += int(payloadLen)

	// Metadata count
	metadataCount := binary.BigEndian.Uint16(buf[offset:])
	assert.Equal(t, uint16(len(meta.Metadata)), metadataCount)
}

func TestDecodeEncodeMeta(t *testing.T) {
	codec := &ZRpcCodec{printDebugInfo: false}

	original := &Meta{
		ServiceName: "test.service",
		MethodName:  "TestMethod",
		SeqID:       12345,
		MsgType:     1,
		Payload:     []byte{0x01, 0x02, 0x03, 0x04},
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Encode
	buf, err := codec.encodeMeta(original)
	assert.NoError(t, err)

	// Decode
	decoded, err := codec.decodeMeta(buf)
	assert.NoError(t, err)
	assert.NotNil(t, decoded)

	// Verify
	assert.Equal(t, original.ServiceName, decoded.ServiceName)
	assert.Equal(t, original.MethodName, decoded.MethodName)
	assert.Equal(t, original.SeqID, decoded.SeqID)
	assert.Equal(t, original.MsgType, decoded.MsgType)
	assert.Equal(t, original.Payload, decoded.Payload)
	assert.Equal(t, len(original.Metadata), len(decoded.Metadata))
	for k, v := range original.Metadata {
		assert.Equal(t, v, decoded.Metadata[k])
	}
}

func TestMagicNumber(t *testing.T) {
	// Verify magic number is correct
	assert.Equal(t, uint32(0x5A525043), zrpcMagicNumber)

	// Verify it spells "ZRPC" in ASCII
	magicBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(magicBytes, zrpcMagicNumber)
	assert.Equal(t, byte('Z'), magicBytes[0])
	assert.Equal(t, byte('R'), magicBytes[1])
	assert.Equal(t, byte('P'), magicBytes[2])
	assert.Equal(t, byte('C'), magicBytes[3])
}

func BenchmarkEncodeMeta(b *testing.B) {
	codec := &ZRpcCodec{printDebugInfo: false}

	meta := &Meta{
		ServiceName: "test.service",
		MethodName:  "TestMethod",
		SeqID:       12345,
		MsgType:     1,
		Payload:     make([]byte, 1024), // 1KB payload
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = codec.encodeMeta(meta)
	}
}

func BenchmarkDecodeMeta(b *testing.B) {
	codec := &ZRpcCodec{printDebugInfo: false}

	meta := &Meta{
		ServiceName: "test.service",
		MethodName:  "TestMethod",
		SeqID:       12345,
		MsgType:     1,
		Payload:     make([]byte, 1024), // 1KB payload
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	buf, _ := codec.encodeMeta(meta)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = codec.decodeMeta(buf)
	}
}
