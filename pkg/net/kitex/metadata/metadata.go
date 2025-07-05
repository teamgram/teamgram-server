// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package metadata

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/zeromicro/go-zero/core/jsonx"
)

var (
	rpcMetadataKey = "rpc_metadata"
)

type TakeoutMessageRange struct {
	MinId int32 `json:"min_id,omitempty"`
	MaxId int32 `json:"max_id,omitempty"`
}

type Takeout struct {
	Id    int64                `json:"id,omitempty"`
	Range *TakeoutMessageRange `json:"range,omitempty"`
}

type RpcMetadata struct {
	ServerId      string   `json:"server_id,omitempty"`
	ClientAddr    string   `json:"client_addr,omitempty"`
	AuthId        int64    `json:"auth_id,omitempty"`
	SessionId     int64    `json:"session_id,omitempty"`
	ReceiveTime   int64    `json:"receive_time,omitempty"`
	UserId        int64    `json:"user_id,omitempty"`
	ClientMsgId   int64    `json:"client_msg_id,omitempty"`
	IsBot         bool     `json:"is_bot,omitempty"`
	Layer         int32    `json:"layer,omitempty"`
	Client        string   `json:"client,omitempty"`
	IsAdmin       bool     `json:"is_admin,omitempty"`
	Takeout       *Takeout `json:"takeout,omitempty"`
	Langpack      string   `json:"langpack,omitempty"`
	PermAuthKeyId int64    `json:"perm_auth_key_id,omitempty"`
	LangCode      string   `json:"lang_code,omitempty"`
}

func RpcMetadataFromIncoming(ctx context.Context) *RpcMetadata {
	v, ok := metainfo.GetPersistentValue(ctx, rpcMetadataKey)
	if !ok {
		return nil
	}

	md := &RpcMetadata{}
	err := jsonx.UnmarshalFromString(v, md)
	if err != nil {
		// log.Errorf("Unmarshal rpc_metadata error: %v", err)
		return nil
	}

	return md
}

func RpcMetadataToOutgoing(ctx context.Context, md *RpcMetadata) (context.Context, error) {
	v, err := jsonx.MarshalToString(md)
	if err != nil {
		// log.Errorf("Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metainfo.WithPersistentValue(ctx, rpcMetadataKey, v), nil
}

func (m *RpcMetadata) String() (val string) {
	val, _ = jsonx.MarshalToString(m)

	return
}

func (m *RpcMetadata) HasTakeout() bool {
	return m.Takeout != nil
}

func (m *RpcMetadata) GetTakeoutId() int64 {
	if m.Takeout == nil {
		return m.Takeout.Id
	} else {
		return 0
	}
}
