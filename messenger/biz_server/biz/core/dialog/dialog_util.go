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

package dialog

import "github.com/nebula-chat/chatengine/mtproto"

type dialogItems struct {
	MessageIdList       []int32
	ChannelMessageIdMap map[int32]int32
	UserIdList          []int32
	ChatIdList          []int32
	ChannelIdList       []int32
}

func makeDialogItems() *dialogItems {
	return &dialogItems{
		MessageIdList:       make([]int32, 0),
		ChannelMessageIdMap: make(map[int32]int32, 0),
		UserIdList:          make([]int32, 0),
		ChatIdList:          make([]int32, 0),
		ChannelIdList:       make([]int32, 0),
	}
}

func (m *DialogModel) PickAllIDListByDialogs(dialogs []*mtproto.Dialog) (items *dialogItems) {
	items = makeDialogItems()

	for _, d := range dialogs {
		dialog := d.To_Dialog()
		p := dialog.GetPeer()

		// TODO(@benqi): 先假设只有PEER_USER
		switch p.GetConstructor() {
		case mtproto.TLConstructor_CRC32_peerUser:
			items.MessageIdList = append(items.MessageIdList, dialog.GetTopMessage())
			items.UserIdList = append(items.UserIdList, p.GetData2().GetUserId())
		case mtproto.TLConstructor_CRC32_peerChat:
			items.MessageIdList = append(items.MessageIdList, dialog.GetTopMessage())
			items.ChatIdList = append(items.ChatIdList, p.GetData2().GetChatId())
		case mtproto.TLConstructor_CRC32_peerChannel:
			items.ChannelMessageIdMap[p.GetData2().GetChannelId()] = dialog.GetTopMessage()
			items.ChannelIdList = append(items.ChannelIdList, p.GetData2().GetChannelId())
		}
	}
	//items.ChannelMessageIdMap = m.channelCallback.GetTopMessageListByIdList(items.ChannelIdList)
	//for _, d := range dialogs {
	//	dialog := d.To_Dialog()
	//	p := dialog.GetPeer()
	//
	//	// TODO(@benqi): 先假设只有PEER_USER
	//	switch p.GetConstructor() {
	//	case mtproto.TLConstructor_CRC32_peerChannel:
	//		dialog.SetTopMessage(items.ChannelMessageIdMap[p.GetData2().GetChannelId()])
	//	}
	//}

	return
}
