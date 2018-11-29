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

package username

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/mysql_dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/mtproto"
)

const (
	MIN_USERNAME_LEN = 5
)

const (
	USERNAME_NOT_EXISTED 	= 0
	USERNAME_EXISTED 		= 1
	USERNAME_EXISTED_NOTME 	= 2
	USERNAME_EXISTED_ISME 	= 3
)

// type usernameData *dataobject.UsernameDO
//{
//	*dataobject.UsernameDO
//}

type UsernameModel struct {
	*mysql_dao.UsernameDAO
}

func (m *UsernameModel) InstallModel() {
	m.UsernameDAO = dao.GetUsernameDAO(dao.DB_MASTER)
}

func (m *UsernameModel) RegisterCallback(cb interface{}) {
}

func init() {
	core.RegisterCoreModel(&UsernameModel{})
}

// not found, return 0
func (m *UsernameModel) GetListByUsernameList(names []string) map[string]*dataobject.UsernameDO {
	doList := m.UsernameDAO.SelectList(names)
	m2 := make(map[string]*dataobject.UsernameDO, len(doList))
	for i := 0; i < len(doList); i++ {
		m2[doList[i].Username] = &doList[i]
	}
	return m2
}

// not found, return 0
func (m *UsernameModel) CheckUsername(username string) int {
	// TODO(@benqi): check len(username) >= 5
	usernameDO := m.UsernameDAO.SelectByUsername(username)
	if usernameDO == nil {
		return USERNAME_NOT_EXISTED
	} else {
		return USERNAME_EXISTED
	}
}

func (m *UsernameModel) CheckAccountUsername(userId int32, username string) int {
	// TODO(@benqi): check len(username) >= 5
	usernameDO := m.UsernameDAO.SelectByUsername(username)
	if usernameDO == nil {
		return USERNAME_NOT_EXISTED
	} else {
		if usernameDO.PeerType == int8(base.PEER_USER) && usernameDO.PeerId == userId {
			return USERNAME_EXISTED_ISME
		} else {
			return USERNAME_EXISTED_NOTME
		}
	}
}

func (m *UsernameModel) CheckChannelUsername(channelId int32, username string) int {
	// TODO(@benqi): check len(username) >= 5
	usernameDO := m.UsernameDAO.SelectByUsername(username)
	if usernameDO == nil {
		return USERNAME_NOT_EXISTED
	} else {
		if usernameDO.PeerType == int8(base.PEER_CHANNEL) && usernameDO.PeerId == channelId {
			return USERNAME_EXISTED_ISME
		} else {
			return USERNAME_EXISTED_NOTME
		}
	}
}

func (m *UsernameModel) UpdateUsernameByPeer(peerType, peerId int32, username string) bool {
	// TODO(@benqi): check len(username) >= 5
	if username == "" {
		m.UsernameDAO.UpdateUsername("", int8(2), peerId)
	} else {
		usernameDO := m.UsernameDAO.SelectByPeer(int8(peerType), peerId)
		if usernameDO == nil {
			usernameDO = &dataobject.UsernameDO{
				PeerType: int8(peerType),
				PeerId:   peerId,
				Username: username,
			}
			m.UsernameDAO.Insert(usernameDO)
		} else {
			m.UsernameDAO.UpdateUsername(username, int8(2), peerId)
		}
	}
	return true
}

func (m *UsernameModel) GetAccountUsername(userId int32) (username string) {
	do := m.UsernameDAO.SelectByPeer(int8(base.PEER_USER), userId)
	if do != nil {
		username = do.Username
	}
	return
}

func (m *UsernameModel) GetChannelUsername(channelId int32) (username string) {
	do := m.UsernameDAO.SelectByPeer(int8(base.PEER_CHANNEL), channelId)
	if do != nil {
		username = do.Username
	}
	return
}

func (m *UsernameModel) ResolveUsername(username string) (*base.PeerUtil, error) {
	// TODO(@benqi): check len(username) >= 5
	var (
		peer *base.PeerUtil
		err error
	)

	if len(username) >= 5 {
		usernameDO := m.UsernameDAO.SelectByUsername(username)
		if usernameDO != nil {
			if usernameDO.PeerType == base.PEER_USER || usernameDO.PeerType == base.PEER_CHANNEL {
				peer = &base.PeerUtil{
					PeerType: int32(usernameDO.PeerType),
					PeerId:   usernameDO.PeerId,
				}
			} else {
				err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_NOT_OCCUPIED)
			}
		} else {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_NOT_OCCUPIED)
		}
	} else {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_INVALID)
	}

	return peer, err
}
