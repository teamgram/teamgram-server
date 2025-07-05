// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/iface/ecode"
	"github.com/teamgram/proto/v2/tg"

	"github.com/bytedance/gopkg/lang/dirtmake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec"
	"github.com/cloudwego/kitex/pkg/remote/codec/perrors"
)

type ZRpcCodec struct {
	printDebugInfo bool
}

func NewZRpcCodec(debug bool) remote.Codec {
	return &ZRpcCodec{printDebugInfo: debug}
}

func (jc *ZRpcCodec) Encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	var validData interface{}
	switch message.MessageType() {
	case remote.Exception:
		switch e := message.Data().(type) {
		case *remote.TransError:
			// validData = &Exception{e.TypeID(), e.Error()}
			validData = tg.NewRpcError(e)
		case error:
			// validData = &Exception{remote.InternalError, e.Error()}
			validData = tg.NewRpcError(e)
		default:
			return errors.New("exception relay must implement error type")
		}
	default:
		validData = message.Data()
	}

	klog.Infof("trans: %v", message.TransInfo().TransStrInfo())

	x := bin.NewEncoder()
	defer x.End()
	_ = validData.(iface.TLObject).Encode(x, 201)
	payload := x.Bytes()
	//payload, err := json.Marshal(validData)
	//if err != nil {
	//	return perrors.NewProtocolError(fmt.Errorf("json encode, marshal payload failed: %w", err))
	//}
	if jc.printDebugInfo {
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
	buf, err := json.Marshal(data)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json encode, marshal data failed: %w", err))
	}
	length := make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(len(buf)))
	_, err = out.WriteBinary(length)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json encode, write length failed: %w", err))
	}
	_, err = out.WriteBinary(buf)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json encode, write data failed: %w", err))
	}
	return nil
}

func (jc *ZRpcCodec) Decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	length := dirtmake.Bytes(8, 8)
	_, err := in.ReadBinary(length)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json decode, read length failed: %w", err))
	}
	l := int(binary.BigEndian.Uint64(length))
	buf := dirtmake.Bytes(l, l)
	_, err = in.ReadBinary(buf)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json decode, read data failed: %w", err))
	}
	data := &Meta{}
	err = json.Unmarshal(buf, data)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json decode, unmarshal Meta data failed: %w", err))
	}
	if err = codec.SetOrCheckSeqID(data.SeqID, message); err != nil {
		return err
	}
	if err = codec.SetOrCheckMethodName(data.MethodName, message); err != nil {
		return err
	}
	if err = codec.NewDataIfNeeded(data.MethodName, message); err != nil {
		return err
	}

	if jc.printDebugInfo {
		klog.Infof("encoded payload: %s\n", hex.EncodeToString(data.Payload))
	}
	if remote.MessageType(data.MsgType) == remote.Exception {
		//var exception tg.TLRpcError
		//d := bin.NewDecoder(data.Payload)
		//err = exception.Decode(d)
		exception2, err2 := tg.DecodeRpcErrorClazz(bin.NewDecoder(data.Payload))
		// var exception Exception
		// err = json.Unmarshal(data.Payload, &exception)
		if err2 != nil {
			return perrors.NewProtocolError(fmt.Errorf("json decode, unmarshal Exception payload failed: %w", err2))
		} else if exception, ok := exception2.(*tg.TLRpcError); !ok {
			return perrors.NewProtocolError(fmt.Errorf("json decode, unmarshal Exception payload failed: %w", exception2))
		} else {
			return ecode.NewCodeError(exception.Code(), exception.ErrorMessage)
		}
	}

	d := bin.NewDecoder(data.Payload)
	// data2, err = iface.DecodeObject(d)
	err = message.Data().(iface.TLObject).Decode(d)

	//err = json.Unmarshal(data.Payload, message.Data())
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json decode, unmarshal payload failed: %w", err))
	}

	klog.Infof("trans: %v", data.Metadata)

	message.TransInfo().PutTransStrInfo(data.Metadata)

	return nil
}

func (jc *ZRpcCodec) Name() string {
	return "ZRPC"
}
