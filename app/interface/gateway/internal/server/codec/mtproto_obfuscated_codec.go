// Copyright (c) 2019-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package codec

import (
	"io"

	"github.com/teamgram/marmota/pkg/net2"
	"github.com/teamgram/proto/mtproto/crypto"

	log "github.com/zeromicro/go-zero/core/logx"
)

type ObfuscatedCodec struct {
	net2.Codec
	dc int16
}

func NewMTProtoObfuscatedCodec(conn io.ReadWriteCloser, d, e *crypto.AesCTR128Encrypt, protocolType uint32, dc int16) *ObfuscatedCodec {
	codec := new(ObfuscatedCodec)
	stream := NewAesCTR128Stream(conn, d, e)
	codec.dc = dc
	switch protocolType {
	case ABRIDGED_INT32_FLAG:
		codec.Codec = NewMTProtoAbridgedCodec(stream)
	case INTERMEDIATE_FLAG:
		codec.Codec = NewMTProtoIntermediateCodec(stream)
	case PADDED_INTERMEDIATE_FLAG:
		codec.Codec = NewMTProtoPaddedIntermediateCodec(stream)
	default:
		log.Errorf("invalid protocolType: %d", protocolType)
	}
	return codec
}

type AesCTR128Stream struct {
	conn    io.ReadWriteCloser
	encrypt *crypto.AesCTR128Encrypt
	decrypt *crypto.AesCTR128Encrypt
}

func NewAesCTR128Stream(conn io.ReadWriteCloser, d, e *crypto.AesCTR128Encrypt) *AesCTR128Stream {
	return &AesCTR128Stream{
		conn:    conn,
		decrypt: d,
		encrypt: e,
	}
}

func (s *AesCTR128Stream) Read(p []byte) (int, error) {
	n, err := s.conn.Read(p)
	if err == nil {
		s.decrypt.Encrypt(p[:n])
		return n, nil
	}
	return n, err
}

func (s *AesCTR128Stream) Write(p []byte) (int, error) {
	s.encrypt.Encrypt(p[:])
	return s.conn.Write(p)
}

func (s *AesCTR128Stream) Close() error {
	return s.conn.Close()
}
