// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package metadata

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
