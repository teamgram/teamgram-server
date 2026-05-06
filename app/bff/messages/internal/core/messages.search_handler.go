// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

package core

import (
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesSearch
// messages.search#29ee847a flags:# peer:InputPeer q:string from_id:flags.0?InputPeer saved_peer_id:flags.2?InputPeer saved_reaction:flags.3?Vector<Reaction> top_msg_id:flags.1?int filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (c *MessagesCore) MessagesSearch(in *tg.TLMessagesSearch) (*tg.MessagesMessages, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	peerUserID, ok := resolveUserPeerID(in.Peer, c.MD.UserId)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}
	if _, ok := in.Filter.(*tg.TLInputMessagesFilterEmpty); ok && in.Q == "" && in.FromId == nil {
		return nil, tg.ErrSearchQueryEmpty
	}
	if _, ok := in.Filter.(*tg.TLInputMessagesFilterEmpty); ok {
		if tag, ok := normalizeSearchHashTag(in.Q); ok {
			r, err := c.svcCtx.Repo.MsgClient.MsgSearchHashtag(c.ctx, &msg.TLMsgSearchHashtag{
				UserId:    c.MD.UserId,
				AuthKeyId: c.MD.PermAuthKeyId,
				PeerType:  payload.PeerTypeUser,
				PeerId:    peerUserID,
				HashTag:   tag,
				OffsetId:  in.OffsetId,
				Limit:     in.Limit,
			})
			if err != nil {
				c.Logger.Errorf("messages.search hashtag - msg error: self_user_id: %d, peer_id: %d, tag: %s, err: %v",
					c.MD.UserId, peerUserID, tag, err)
				return nil, mapMsgSendError(err)
			}
			return r, nil
		}
	}

	return emptyMessagesMessages(), nil
}

func normalizeSearchHashTag(q string) (string, bool) {
	q = strings.TrimSpace(q)
	if !strings.HasPrefix(q, "#") || len(q) <= 1 {
		return "", false
	}
	tag := strings.TrimPrefix(q, "#")
	if tag == "" || strings.ContainsAny(tag, " \t\r\n") {
		return "", false
	}
	return tag, true
}

func emptyMessagesMessages() *tg.MessagesMessages {
	return tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{},
		Topics:   []tg.ForumTopicClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesMessages()
}
