// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package tg

import (
	"fmt"
)

const (
	PEER_EMPTY           = 0
	PEER_SELF            = 1
	PEER_USER            = 2
	PEER_CHAT            = 3
	PEER_CHANNEL         = 4
	PEER_USERS           = 5
	PEER_CHATS           = 6
	PEER_ENCRYPTED_CHAT  = 7
	PEER_BROADCASTS      = 8
	PEER_USER_MESSAGE    = 10
	PEER_CHANNEL_MESSAGE = 11
	PEER_UNKNOWN         = -1
)

var (
// EmptyPeer
)

//type PeerUtil struct {
//	selfId     int64
//	PeerType   int32
//	PeerId     int64
//	AccessHash int64
//}

func (m *PeerUtil) SelfId() int64 {
	return m.Clazz.(*TLPeerUtil).SelfId
}

func (m *PeerUtil) SetSelfId(v int64) {
	m.Clazz.(*TLPeerUtil).SelfId = v
}

func (m *PeerUtil) PeerType() int32 {
	return m.Clazz.(*TLPeerUtil).PeerType
}

func (m *PeerUtil) SetPeerType(v int32) {
	m.Clazz.(*TLPeerUtil).PeerType = v
}

func (m *PeerUtil) PeerId() int64 {
	return m.Clazz.(*TLPeerUtil).PeerId
}

func (m *PeerUtil) SetPeerId(v int64) {
	m.Clazz.(*TLPeerUtil).PeerId = v
}

func (m *PeerUtil) AccessHash() int64 {
	return m.Clazz.(*TLPeerUtil).AccessHash
}

func (m *PeerUtil) SetAccessHash(v int64) {
	m.Clazz.(*TLPeerUtil).AccessHash = v
}

func (m *PeerUtil) ToString() (s string) {
	switch m.PeerType() {
	case PEER_EMPTY:
		return fmt.Sprintf("PEER_EMPTY: {peer_id: %d, access_hash: %d", m.PeerId(), m.AccessHash())
	case PEER_SELF:
		return fmt.Sprintf("PEER_SELF: {peer_id: %d, access_hash: %d", m.PeerId(), m.AccessHash())
	case PEER_USER:
		return fmt.Sprintf("PEER_USER: {peer_id: %d, access_hash: %d", m.PeerId(), m.AccessHash())
	case PEER_CHAT:
		return fmt.Sprintf("PEER_CHAT: {peer_id: %d, access_hash: %d", m.PeerId(), m.AccessHash())
	case PEER_CHANNEL:
		return fmt.Sprintf("PEER_CHANNEL: {peer_id: %d, access_hash: %d", m.PeerId(), m.AccessHash())
	case PEER_USERS:
		return fmt.Sprintf("PEER_USERS: {peer_id: %d, access_hash: %d", m.PeerId(), m.AccessHash())
	case PEER_CHATS:
		return fmt.Sprintf("PEER_CHATS: {peer_id: %d, access_hash: %d", m.PeerId(), m.AccessHash())
	//case PEER_ALL:
	//	return fmt.Sprintf("PEER_ALL: {peer_id: %d, access_hash: %d", m.PeerId, m.AccessHash)
	default:
		return fmt.Sprintf("PEER_UNKNOWN: {peer_id: %d, access_hash: %d", m.PeerId(), m.AccessHash())
	}
	// return
}

func (m *PeerUtil) CanDoSendMessage() bool {
	switch m.PeerType() {
	case PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		return true
	default:
		return false
	}
}

func FromInputUser(selfId int64, user InputUserClazz) PeerUtilClazz {
	p := &TLPeerUtil{
		PeerType: PEER_UNKNOWN,
	}

	switch c := user.(type) {
	case *TLInputUserEmpty:
		p.PeerType = PEER_EMPTY
	case *TLInputUserSelf:
		p.PeerType = PEER_SELF
		p.PeerId = selfId
	case *TLInputUser:
		p.PeerType = PEER_USER
		p.PeerId = c.UserId
		p.AccessHash = c.AccessHash
	case *TLInputUserFromMessage:
		// p.PeerType = PEER_USER
	default:
		p.PeerType = PEER_UNKNOWN
	}

	return p
}

func FromInputPeer(peer InputPeerClazz) PeerUtilClazz {
	p := &TLPeerUtil{
		PeerType: PEER_UNKNOWN,
	}

	switch c := peer.(type) {
	case *TLInputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case *TLInputPeerSelf:
		p.PeerType = PEER_SELF
		p.PeerId = 0 // selfId is not set here, it should be set by caller
	case *TLInputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = c.UserId
		p.AccessHash = c.AccessHash
	case *TLInputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = c.ChatId
	case *TLInputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = c.ChannelId
		p.AccessHash = c.AccessHash
	default:
		p.PeerType = PEER_UNKNOWN
	}

	return p
}

func FromInputPeer2(selfId int64, peer InputPeerClazz) PeerUtilClazz {
	p := &TLPeerUtil{
		PeerType: PEER_UNKNOWN,
		SelfId:   selfId,
	}

	switch c := peer.(type) {
	case *TLInputPeerEmpty:
		p.PeerType = PEER_EMPTY
		p.PeerId = 0
		p.AccessHash = 0
	case *TLInputPeerSelf:
		p.PeerType = PEER_SELF
		p.PeerId = selfId
		p.AccessHash = 0
	case *TLInputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = c.UserId
		p.AccessHash = c.AccessHash
	case *TLInputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = c.ChatId
		p.AccessHash = 0
	case *TLInputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = c.ChannelId
		p.AccessHash = c.AccessHash
	default:
		p.PeerType = PEER_UNKNOWN
	}

	return p
}

func FromInputEncryptedChat(peer InputEncryptedChatClazz) PeerUtilClazz {
	p := &TLPeerUtil{
		PeerType: PEER_UNKNOWN,
	}

	switch c := peer.(type) {

	case *TLInputEncryptedChat:
		p.PeerType = PEER_ENCRYPTED_CHAT
		p.PeerId = int64(c.ChatId)
		p.AccessHash = c.AccessHash
	default:
		p.PeerType = PEER_UNKNOWN
	}

	return p
}

//
//func (m *PeerUtil) IsSelf() bool {
//	switch m.PeerType {
//	case PEER_SELF:
//		return true
//	case PEER_USER:
//		return m.selfId == m.PeerId
//	}
//	return false
//}

func (m *PeerUtil) ToInputPeer() (peer InputPeerClazz) {
	switch m.PeerType() {
	case PEER_EMPTY:
		peer = &TLInputPeerEmpty{}
	case PEER_SELF:
		peer = &TLInputPeerSelf{}
	case PEER_USER:
		peer = &TLInputPeerUser{
			UserId:     m.PeerId(),
			AccessHash: m.AccessHash(),
		}
	case PEER_CHAT:
		peer = &TLInputPeerChat{
			ChatId: m.PeerId(),
		}
	case PEER_CHANNEL:
		peer = &TLInputPeerChannel{
			ChannelId:  m.PeerId(),
			AccessHash: m.AccessHash(),
		}
	default:
		panic(fmt.Sprintf("ToInputPeer(%v) error!", m))
	}

	return
}

func FromPeer(peer PeerClazz) PeerUtilClazz {
	p := &TLPeerUtil{}

	switch c := peer.(type) {
	case *TLPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = c.UserId
		p.AccessHash = 0
	case *TLPeerChat:
		p.PeerType = PEER_SELF
		p.PeerId = c.ChatId
		p.AccessHash = 0
	case *TLPeerChannel:
		p.PeerType = PEER_USER
		p.PeerId = c.ChannelId
		p.AccessHash = 0
	default:
		p.PeerType = PEER_UNKNOWN
	}

	return p
}

func (m *PeerUtil) ToPeer() (peer PeerClazz) {
	switch m.PeerType() {
	case PEER_SELF:
		if m.PeerId() != 0 {
			peer = &TLPeerUser{
				UserId: m.PeerId(),
			}
		} else if m.SelfId() != 0 {
			peer = &TLPeerUser{
				UserId: m.SelfId(),
			}
		} else {
			panic(fmt.Sprintf("ToPeer(%v) error!", m))
		}
	case PEER_USER:
		peer = &TLPeerUser{
			UserId: m.PeerId(),
		}
	case PEER_CHAT:
		peer = &TLPeerChat{
			ChatId: m.PeerId(),
		}
	case PEER_CHANNEL:
		peer = &TLPeerChannel{
			ChannelId: m.PeerId(),
		}
	default:
		peer = nil
	}

	return
}

/*
func (m *PeerUtil) IsEmpty() bool {
	return m.PeerType() == PEER_EMPTY
}

func (m *PeerUtil) IsSelf() bool {
	return m.PeerType() == PEER_SELF
}

func (m *PeerUtil) IsUser() bool {
	return m.PeerType() == PEER_USER || m.PeerType() == PEER_SELF
}

func (m *PeerUtil) IsChat() bool {
	return m.PeerType() == PEER_CHAT
}

func (m *PeerUtil) IsChatOrChannel() bool {
	return m.IsChat() || m.IsChannel()
}

func (m *PeerUtil) IsUserOrChatOrChannel() bool {
	return m.IsUser() || m.IsChat() || m.IsChannel()
}

func (m *PeerUtil) IsChatOrUser() bool {
	return m.IsUser() || m.IsChat()
}

func (m *PeerUtil) IsChannel() bool {
	return m.PeerType() == PEER_CHANNEL
}

func (m *PeerUtil) IsEncryptedChat() bool {
	return m.PeerType() == PEER_ENCRYPTED_CHAT
}

func (m *PeerUtil) IsSelfUser(id int64) bool {
	return m.PeerType() == PEER_SELF || m.PeerType() == PEER_USER && m.PeerId() == id
}

func FromInputNotifyPeer(selfId int64, peer *InputNotifyPeer) *PeerUtil {
	p := &TLPeerUtil{
		PeerType: PEER_UNKNOWN,
	}

	peer.Match(
		func(c *TLInputNotifyPeer) interface{} {
			p2 := FromInputPeer2(selfId, c.Peer)
			p, _ = p2.ToPeerUtil()

			return nil
		},
		func(c *TLInputNotifyUsers) interface{} {
			p.PeerType = PEER_USERS
			p.PeerId = 0 // Users notification does not have a specific peer ID
			p.AccessHash = 0

			return nil
		},
		func(c *TLInputNotifyChats) interface{} {
			p.PeerType = PEER_CHATS
			p.PeerId = 0 // Chats notification does not have a specific peer ID
			p.AccessHash = 0

			return nil
		},
		func(c *TLInputNotifyBroadcasts) interface{} {
			p.PeerType = PEER_BROADCASTS
			p.PeerId = 0 // Broadcasts notification does not have a specific peer ID
			p.AccessHash = 0

			return nil
		})

	return p.ToPeerUtil()
}

func (m *PeerUtil) ToInputNotifyPeer(peer *InputNotifyPeer) {
	switch m.PeerType() {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = MakeInputNotifyPeer(&TLInputNotifyPeer{
			Peer: m.ToInputPeer(),
		})
	case PEER_USERS:
		peer = MakeInputNotifyPeer(&TLInputNotifyUsers{})
	case PEER_CHATS:
		peer = MakeInputNotifyPeer(&TLInputNotifyChats{})
	default:
		panic(fmt.Sprintf("ToInputNotifyPeer(%v) error!", m))
	}

	return
}

//func FromNotifyPeer(peer *NotifyPeer) (m *PeerUtil) {
//	p = &PeerUtil{}
//	switch peer.GetConstructor() {
//	case CRC32_notifyPeer:
//		p = FromPeer(peer.GetPeer())
//	case CRC32_notifyUsers:
//		m.PeerType = PEER_USERS
//	case CRC32_notifyChats:
//		m.PeerType = PEER_CHATS
//	case CRC32_notifyAll:
//		m.PeerType = PEER_ALL
//	default:
//		panic(fmt.Sprintf("FromNotifyPeer(%v) error!", p))
//	}
//	return
//}

func (m *PeerUtil) ToNotifyPeer() (peer *NotifyPeer) {
	switch m.PeerType() {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = MakeNotifyPeer(&TLNotifyPeer{
			Peer: m.ToPeer(),
		})
	case PEER_USERS:
		peer = MakeNotifyPeer(&TLNotifyUsers{})
	case PEER_CHATS:
		peer = MakeNotifyPeer(&TLNotifyChats{})
	case PEER_BROADCASTS:
		peer = MakeNotifyPeer(&TLNotifyBroadcasts{})
	default:
		panic(fmt.Sprintf("ToNotifyPeer(%v) error!", m))
	}
	return
}
*/

func ToPeerByTypeAndID(peerType int32, peerId int64) (peer PeerClazz) {
	switch peerType {
	case PEER_USER:
		peer = &TLPeerUser{
			UserId: peerId,
		}
	case PEER_CHAT:
		peer = &TLPeerChat{
			ChatId: peerId,
		}
	case PEER_CHANNEL:
		peer = &TLPeerChannel{
			ChannelId: peerId,
		}
	default:
		panic(fmt.Sprintf("ToPeerByTypeAndID(%d, %d) error!", peerType, peerId))
	}

	return
}

func PickAllIdListByPeers(peers []PeerClazz) (idList, chatIdList, channelIdList []int64) {
	for _, p := range peers {
		switch c := p.(type) {
		case *TLPeerUser:
			idList = append(idList, c.UserId)
		case *TLPeerChat:
			chatIdList = append(chatIdList, c.ChatId)
		case *TLPeerChannel:
			channelIdList = append(channelIdList, c.ChannelId)
		default:
			panic(fmt.Sprintf("PickAllIdListByPeers(%v) error!", p))
		}
	}

	if idList == nil {
		idList = []int64{}
	}
	if chatIdList == nil {
		chatIdList = []int64{}
	}
	if channelIdList == nil {
		channelIdList = []int64{}
	}

	return
}

func MakePeerUser(peerId int64) PeerClazz {
	return &TLPeerUser{
		UserId: peerId,
	}
}

func MakePeerChat(peerId int64) PeerClazz {
	return &TLPeerChat{
		ChatId: peerId,
	}
}

func MakePeerChannel(peerId int64) PeerClazz {
	return &TLPeerChannel{
		ChannelId: peerId,
	}
}

func MakePeerHelper(peerType int32, peerId int64) PeerClazz {
	switch peerType {
	case PEER_CHAT:
		return MakePeerChat(peerId)
	case PEER_CHANNEL:
		return MakePeerChannel(peerId)
	default:
		return MakePeerUser(peerId)
	}
}

func MakePeerUtilHelper(peerType int32, peerId int64) PeerUtilClazz {
	return &TLPeerUtil{
		PeerType: peerType,
		PeerId:   peerId,
	}
}

func MakeUserPeerUtil(peerId int64) PeerUtilClazz {
	return &TLPeerUtil{
		PeerType: PEER_USER,
		PeerId:   peerId,
	}
}

func MakeChatPeerUtil(peerId int64) PeerUtilClazz {
	return &TLPeerUtil{
		PeerType: PEER_CHAT,
		PeerId:   peerId,
	}
}

func MakeChannelPeerUtil(peerId int64) PeerUtilClazz {
	return &TLPeerUtil{
		PeerType: PEER_CHANNEL,
		PeerId:   peerId,
	}
}

func MakeInputPeerChat(peerId int64) InputPeerClazz {
	return &TLInputPeerChat{
		ChatId: peerId,
	}
}

func MakeInputPeerChannel(peerId int64) InputPeerClazz {
	return &TLInputPeerChannel{
		ChannelId: peerId,
	}
}

func MakePeerDialogId(peerType int32, peerId int64) int64 {
	var (
		id int64 = 0
	)

	switch peerType {
	case PEER_SELF, PEER_USER:
		id = peerId
	case PEER_CHAT, PEER_CHANNEL:
		id = -peerId
	default:
		// LOGO
	}

	return id
}

func MakePeerUtilDialogId(peer *TLPeerUtil) int64 {
	return MakePeerDialogId(peer.PeerType, peer.PeerId)
}

func GetPeerUtilByPeerDialogId(id int64) (int32, int64) {
	if id > 0 {
		return PEER_USER, id
	} else {
		if ChatIdIsChat(-id) {
			return PEER_CHAT, -id
		} else {
			return PEER_CHANNEL, -id
		}
	}
}

func IsChannelInputPeer(peer InputPeerClazz) bool {
	return peer.InputPeerClazzName() == ClazzName_inputPeerChannel
}

func IsChatInputPeer(peer InputPeerClazz) bool {
	return peer.InputPeerClazzName() == ClazzName_inputPeerChat
}

func IsUserInputPeer(peer InputPeerClazz) bool {
	return peer.InputPeerClazzName() == ClazzName_inputPeerUser ||
		peer.InputPeerClazzName() == ClazzName_inputPeerSelf
}

func PeerIsChannel(peer PeerClazz) bool {
	return peer.PeerClazzName() == ClazzName_peerChannel
}
