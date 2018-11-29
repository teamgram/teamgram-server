// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package base

import (
	"fmt"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/golang/glog"
)

const (
	PEER_EMPTY   = 0
	PEER_SELF    = 1
	PEER_USER    = 2
	PEER_CHAT    = 3
	PEER_CHANNEL = 4
	PEER_USERS   = 5
	PEER_CHATS   = 6
	// PEER_ALL     = 7
	PEER_UNKNOWN = -1
)

type PeerUtil struct {
	selfId     int32
	PeerType   int32
	PeerId     int32
	AccessHash int64
}

func (p PeerUtil) String() (s string) {
	switch p.PeerType {
	case PEER_EMPTY:
		return fmt.Sprintf("PEER_EMPTY: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_SELF:
		return fmt.Sprintf("PEER_SELF: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USER:
		return fmt.Sprintf("PEER_USER: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHAT:
		return fmt.Sprintf("PEER_CHAT: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHANNEL:
		return fmt.Sprintf("PEER_CHANNEL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USERS:
		return fmt.Sprintf("PEER_USERS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHATS:
		return fmt.Sprintf("PEER_CHATS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	//case PEER_ALL:
	//	return fmt.Sprintf("PEER_ALL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	default:
		return fmt.Sprintf("PEER_UNKNOWN: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	}
	// return
}

func FromInputPeer(peer *mtproto.InputPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case mtproto.TLConstructor_CRC32_inputPeerSelf:
		p.PeerType = PEER_SELF
	case mtproto.TLConstructor_CRC32_inputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.GetData2().GetUserId()
		p.AccessHash = peer.GetData2().GetAccessHash()
	case mtproto.TLConstructor_CRC32_inputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.GetData2().GetChatId()
	case mtproto.TLConstructor_CRC32_inputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.GetData2().GetChannelId()
		p.AccessHash = peer.GetData2().GetAccessHash()
	default:
		panic(fmt.Sprintf("FromInputPeer(%v) error!", peer))
	}
	return
}

func FromInputPeer2(selfId int32, peer *mtproto.InputPeer) (p *PeerUtil) {
	p = &PeerUtil{
		selfId: selfId,
	}
	switch peer.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case mtproto.TLConstructor_CRC32_inputPeerSelf:
		p.PeerType = PEER_SELF
		p.PeerId = selfId
	case mtproto.TLConstructor_CRC32_inputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.GetData2().GetUserId()
		p.AccessHash = peer.GetData2().GetAccessHash()
	case mtproto.TLConstructor_CRC32_inputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.GetData2().GetChatId()
	case mtproto.TLConstructor_CRC32_inputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.GetData2().GetChannelId()
		p.AccessHash = peer.GetData2().GetAccessHash()
	default:
		panic(fmt.Sprintf("FromInputPeer(%v) error!", peer))
	}
	return

}

func (p *PeerUtil) ToInputPeer() (peer *mtproto.InputPeer) {
	switch p.PeerType {
	case PEER_EMPTY:
		peer = &mtproto.InputPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputPeerEmpty,
			Data2:       &mtproto.InputPeer_Data{},
		}
	case PEER_SELF:
		peer = &mtproto.InputPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputPeerSelf,
			Data2:       &mtproto.InputPeer_Data{},
		}
	case PEER_USER:
		peer = &mtproto.InputPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputPeerUser,
			Data2: &mtproto.InputPeer_Data{
				UserId:     p.PeerId,
				AccessHash: p.AccessHash,
			},
		}
	case PEER_CHAT:
		peer = &mtproto.InputPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputPeerChat,
			Data2: &mtproto.InputPeer_Data{
				ChatId: p.PeerId,
			},
		}
	case PEER_CHANNEL:
		peer = &mtproto.InputPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputPeerChannel,
			Data2: &mtproto.InputPeer_Data{
				ChannelId:  p.PeerId,
				AccessHash: p.AccessHash,
			},
		}
	default:
		panic(fmt.Sprintf("ToInputPeer(%v) error!", p))
	}
	return
}

func FromPeer(peer *mtproto.Peer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.GetConstructor() {
	case mtproto.TLConstructor_CRC32_peerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.GetData2().GetUserId()
	case mtproto.TLConstructor_CRC32_peerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.GetData2().GetChatId()
	case mtproto.TLConstructor_CRC32_peerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.GetData2().GetChannelId()
	default:
		panic(fmt.Sprintf("FromPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToPeer() (peer *mtproto.Peer) {
	switch p.PeerType {
	case PEER_SELF:
		if p.PeerId != 0 {
			peer = &mtproto.Peer{
				Constructor: mtproto.TLConstructor_CRC32_peerUser,
				Data2:       &mtproto.Peer_Data{UserId: p.PeerId},
			}
		} else if p.selfId != 0 {
			peer = &mtproto.Peer{
				Constructor: mtproto.TLConstructor_CRC32_peerUser,
				Data2:       &mtproto.Peer_Data{UserId: p.selfId},
			}
		} else {
			panic(fmt.Sprintf("ToPeer(%v) error!", p))
		}
	case PEER_USER:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerUser,
			Data2:       &mtproto.Peer_Data{UserId: p.PeerId},
		}
	case PEER_CHAT:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerChat,
			Data2:       &mtproto.Peer_Data{ChatId: p.PeerId},
		}
	case PEER_CHANNEL:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerChannel,
			Data2:       &mtproto.Peer_Data{ChannelId: p.PeerId},
		}
	default:
		panic(fmt.Sprintf("ToPeer(%v) error!", p))
	}
	return
}

func FromInputNotifyPeer(selfId int32, peer *mtproto.InputNotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputNotifyPeer:
		p = FromInputPeer2(selfId, peer.GetData2().GetPeer())
	case mtproto.TLConstructor_CRC32_inputNotifyUsers:
		p.PeerType = PEER_USERS
	case mtproto.TLConstructor_CRC32_inputNotifyChats:
		p.PeerType = PEER_CHATS
	//case mtproto.TLConstructor_CRC32_inputNotifyAll:
	//	p.PeerType = PEER_ALL
	default:
		glog.Error("fromInputNotifyPeer: invalid peer - ", peer)
	}
	return
}

func (p *PeerUtil) ToInputNotifyPeer(peer *mtproto.InputNotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = &mtproto.InputNotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputNotifyPeer,
			Data2: &mtproto.InputNotifyPeer_Data{
				Peer: p.ToInputPeer(),
			},
		}
	case PEER_USERS:
		peer = &mtproto.InputNotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputNotifyUsers,
			Data2:       &mtproto.InputNotifyPeer_Data{},
		}
	case PEER_CHATS:
		peer = &mtproto.InputNotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_inputNotifyChats,
			Data2:       &mtproto.InputNotifyPeer_Data{},
		}
	//case PEER_ALL:
	//	peer = &mtproto.InputNotifyPeer{
	//		Constructor: mtproto.TLConstructor_CRC32_inputNotifyAll,
	//		Data2:       &mtproto.InputNotifyPeer_Data{},
	//	}
	default:
		panic(fmt.Sprintf("ToInputNotifyPeer(%v) error!", p))
	}
	return
}

func FromNotifyPeer(peer *mtproto.NotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.GetConstructor() {
	case mtproto.TLConstructor_CRC32_notifyPeer:
		p = FromPeer(peer.GetData2().GetPeer())
	case mtproto.TLConstructor_CRC32_notifyUsers:
		p.PeerType = PEER_USERS
	case mtproto.TLConstructor_CRC32_notifyChats:
		p.PeerType = PEER_CHATS
	//case mtproto.TLConstructor_CRC32_notifyAll:
	//	p.PeerType = PEER_ALL
	default:
		panic(fmt.Sprintf("FromNotifyPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToNotifyPeer() (peer *mtproto.NotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = &mtproto.NotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_notifyPeer,
			Data2: &mtproto.NotifyPeer_Data{
				Peer: p.ToPeer(),
			},
		}
	case PEER_USERS:
		peer = &mtproto.NotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_notifyUsers,
			Data2:       &mtproto.NotifyPeer_Data{},
		}
	case PEER_CHATS:
		peer = &mtproto.NotifyPeer{
			Constructor: mtproto.TLConstructor_CRC32_notifyChats,
			Data2:       &mtproto.NotifyPeer_Data{},
		}
	//case PEER_ALL:
	//	peer = &mtproto.NotifyPeer{
	//		Constructor: mtproto.TLConstructor_CRC32_notifyAll,
	//		Data2:       &mtproto.NotifyPeer_Data{},
	//	}
	default:
		panic(fmt.Sprintf("ToNotifyPeer(%v) error!", p))
	}
	return
}

func ToPeerByTypeAndID(peerType int8, peerId int32) (peer *mtproto.Peer) {
	switch peerType {
	case PEER_USER:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerUser,
			Data2:       &mtproto.Peer_Data{UserId: peerId},
		}
	case PEER_CHAT:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerChat,
			Data2:       &mtproto.Peer_Data{ChatId: peerId},
		}
	case PEER_CHANNEL:
		peer = &mtproto.Peer{
			Constructor: mtproto.TLConstructor_CRC32_peerChannel,
			Data2:       &mtproto.Peer_Data{ChannelId: peerId},
		}
	default:
		panic(fmt.Sprintf("ToPeerByTypeAndID(%d, %d) error!", peerType, peerId))
	}
	return
}
