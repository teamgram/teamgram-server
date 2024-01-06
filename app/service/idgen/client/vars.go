// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package idgen_client

import (
	"strconv"

	"github.com/teamgram/teamgram-server/app/service/idgen/idgen"
)

// ///////////////////////////////////////////////////////////////////////////////////////////
var (
	// TODO(@benqi): 使用更紧凑的前缀
	messageDataNgenId       = "message_data_ngen_"
	messageBoxUpdatesNgenId = "message_box_ngen_"
	channelMessageBoxNgenId = "channel_message_box_ngen_"
	seqUpdatesNgenId        = "seq_updates_ngen_"
	ptsUpdatesNgenId        = "pts_updates_ngen_"
	qtsUpdatesNgenId        = "qts_updates_ngen_"
	channelPtsUpdatesNgenId = "channel_pts_updates_ngen_"
	scheduledMessageNgenId  = "scheduled_ngen_"
	botUpdatesNgenId        = "bot_updates_ngen_"
	storyNgenId             = "story_ngen_"
	channelStoryNgenId      = "channel_story_ngen_"
)

func genMessageDataNgenIdKey(id int64) string {
	return messageDataNgenId + strconv.FormatInt(id, 10)
}

func genMessageBoxUpdatesNgenIdKey(id int64) string {
	return messageBoxUpdatesNgenId + strconv.FormatInt(id, 10)
}

func genChannelMessageBoxNgenIdKey(id int64) string {
	return channelMessageBoxNgenId + strconv.FormatInt(id, 10)
}

func genSeqUpdatesNgenIdKey(id int64) string {
	return seqUpdatesNgenId + strconv.FormatInt(id, 10)
}

func genPtsUpdatesNgenIdKey(id int64) string {
	return ptsUpdatesNgenId + strconv.FormatInt(id, 10)
}

func genQtsUpdatesNgenIdKey(id int64) string {
	return qtsUpdatesNgenId + strconv.FormatInt(id, 10)
}

func genChannelPtsUpdatesNgenIdKey(id int64) string {
	return channelPtsUpdatesNgenId + strconv.FormatInt(id, 10)
}

func genScheduledMessageNgenIdKey(id int64) string {
	return scheduledMessageNgenId + strconv.FormatInt(id, 10)
}

func genBotUpdatesNgenIdKey(id int64) string {
	return botUpdatesNgenId + strconv.FormatInt(id, 10)
}

func genStoryNgenIdKey(id int64) string {
	return storyNgenId + strconv.FormatInt(id, 10)
}

func genChannelStoryNgenIdKey(id int64) string {
	return channelStoryNgenId + strconv.FormatInt(id, 10)
}

const (
	IDTypeNextId            = 0
	IDTypeMessageData       = 1
	IDTypeMessageBox        = 2
	IDTypeChannelMessageBox = 3
	IDTypeSeq               = 4
	IDTypePts               = 5
	IDTypeQts               = 6
	IDTypeChannelPts        = 7
	IDTypeScheduledMessage  = 8
	IDTypeBot               = 9
	IDTypeStory             = 10
	IDTypeChannelStory      = 11
)

var (
	defaultInputNextId = idgen.MakeTLInputId(nil).To_InputId()
)

type IDTypeNgen struct {
	IDType int
	Key    int64
	N      int
}

func MakeIDTypeNextId() IDTypeNgen {
	return IDTypeNgen{
		IDType: IDTypeNextId,
	}
}

func MakeIDTypeNextIdN(num int) IDTypeNgen {
	return IDTypeNgen{
		IDType: IDTypeNextId,
		N:      num,
	}
}

func MakeIDTypeNgen(idType int, key int64) IDTypeNgen {
	return IDTypeNgen{
		IDType: idType,
		Key:    key,
	}
}

func MakeIDTypeNgenN(idType int, key int64, n int) IDTypeNgen {
	return IDTypeNgen{
		IDType: idType,
		Key:    key,
		N:      n,
	}
}

func (m IDTypeNgen) ToInputId() *idgen.InputId {
	switch m.IDType {
	case IDTypeNextId:
		if m.N > 0 {
			return MakeInputIds(m.N)
		} else {
			return defaultInputNextId
		}
	default:
		switch m.IDType {
		case IDTypeMessageData:
			if m.N == 0 {
				return MakeInputSeqId(genMessageDataNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genMessageDataNgenIdKey(m.Key), m.N)
			}
		case IDTypeMessageBox:
			if m.N == 0 {
				return MakeInputSeqId(genMessageBoxUpdatesNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genMessageBoxUpdatesNgenIdKey(m.Key), m.N)
			}
		case IDTypeChannelMessageBox:
			if m.N == 0 {
				return MakeInputSeqId(genChannelMessageBoxNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genChannelMessageBoxNgenIdKey(m.Key), m.N)
			}
		case IDTypeSeq:
			if m.N == 0 {
				return MakeInputSeqId(genSeqUpdatesNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genSeqUpdatesNgenIdKey(m.Key), m.N)
			}
		case IDTypePts:
			if m.N == 0 {
				return MakeInputSeqId(genPtsUpdatesNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genPtsUpdatesNgenIdKey(m.Key), m.N)
			}
		case IDTypeQts:
			if m.N == 0 {
				return MakeInputSeqId(genQtsUpdatesNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genQtsUpdatesNgenIdKey(m.Key), m.N)
			}
		case IDTypeChannelPts:
			if m.N == 0 {
				return MakeInputSeqId(genChannelPtsUpdatesNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genChannelPtsUpdatesNgenIdKey(m.Key), m.N)
			}
		case IDTypeScheduledMessage:
			if m.N == 0 {
				return MakeInputSeqId(genScheduledMessageNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genScheduledMessageNgenIdKey(m.Key), m.N)
			}
		case IDTypeBot:
			if m.N == 0 {
				return MakeInputSeqId(genBotUpdatesNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genBotUpdatesNgenIdKey(m.Key), m.N)
			}
		case IDTypeStory:
			if m.N == 0 {
				return MakeInputSeqId(genStoryNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genStoryNgenIdKey(m.Key), m.N)
			}
		case IDTypeChannelStory:
			if m.N == 0 {
				return MakeInputSeqId(genChannelStoryNgenIdKey(m.Key))
			} else {
				return MakeInputNSeqId(genChannelStoryNgenIdKey(m.Key), m.N)
			}
		}
	}
	return MakeInputId()
}

type IDValue struct {
	IDType int
	Id     int64
	IdN    []int64
}

func MakeInputId() *idgen.InputId {
	return idgen.MakeTLInputId(nil).To_InputId()
}

func MakeInputIds(num int) *idgen.InputId {
	return idgen.MakeTLInputIds(&idgen.InputId{
		Num: int32(num),
	}).To_InputId()
}

func MakeInputSeqId(key string) *idgen.InputId {
	return idgen.MakeTLInputSeqId(&idgen.InputId{
		Key: key,
	}).To_InputId()
}

func MakeInputNSeqId(key string, n int) *idgen.InputId {
	return idgen.MakeTLInputNSeqId(&idgen.InputId{
		Key: key,
		N:   int32(n),
	}).To_InputId()
}
