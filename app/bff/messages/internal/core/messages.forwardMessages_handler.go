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
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const maxForwardMessages = 100

// MessagesForwardMessages
// messages.forwardMessages#13704a7c flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true drop_author:flags.11?true drop_media_captions:flags.12?true noforwards:flags.14?true allow_paid_floodskip:flags.19?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer top_msg_id:flags.9?int reply_to:flags.22?InputReplyTo schedule_date:flags.10?int schedule_repeat_period:flags.24?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long video_timestamp:flags.20?int allow_paid_stars:flags.21?long suggested_post:flags.23?SuggestedPost = Updates;
func (c *MessagesCore) MessagesForwardMessages(in *tg.TLMessagesForwardMessages) (*tg.Updates, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	selfUserID := md.UserId
	authKeyID := md.PermAuthKeyId

	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	sourcePeerID, ok := resolveUserPeerID(in.FromPeer, selfUserID)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}
	peerUserID, ok := resolveUserPeerID(in.ToPeer, selfUserID)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}

	if err := checkForwardMessagesUnsupportedFields(in); err != nil {
		return nil, err
	}
	if err := checkForwardMessagesIDs(in.Id, in.RandomId); err != nil {
		return nil, err
	}

	var sendClient sendMessageClient = c.svcCtx.Repo.MsgClient
	sourceList, err := sendClient.MsgGetUserMessageList(c.ctx, &msg.TLMsgGetUserMessageList{
		UserId: selfUserID,
		IdList: in.Id,
	})
	if err != nil {
		c.Logger.Errorf("messages.forwardMessages - get source messages failed: self_user_id: %d, source_peer_id: %d, ids: %v, err: %v",
			selfUserID, sourcePeerID, in.Id, err)
		return nil, mapMsgSendError(err)
	}

	sources, err := orderForwardSources(sourceList, in.Id)
	if err != nil {
		return nil, err
	}

	date := int32(time.Now().Unix())
	sourcePeer := tg.MakePeerUser(sourcePeerID)
	forwardedGroupedIDs := make(map[int64]int64)
	outboxes := make([]msg.OutboxMessageClazz, 0, len(sources))
	for i, source := range sources {
		var groupedID *int64
		if source.GroupedId != nil {
			v, ok := forwardedGroupedIDs[*source.GroupedId]
			if !ok {
				v = newSendMultiMediaGroupedID()
				forwardedGroupedIDs[*source.GroupedId] = v
			}
			groupedID = &v
		}
		outboxes = append(outboxes, buildForwardOutbox(forwardOutboxInput{
			RandomID:        in.RandomId[i],
			Background:      in.Background,
			Silent:          in.Silent,
			Noforwards:      in.Noforwards,
			FromUserID:      selfUserID,
			PeerUserID:      peerUserID,
			SourcePeer:      sourcePeer,
			SourceMessageID: in.Id[i],
			Date:            date,
			Source:          source,
			GroupedID:       groupedID,
		}))
	}

	updates, err := sendClient.MsgSendMessageV2(c.ctx, &msg.TLMsgSendMessageV2{
		UserId:              selfUserID,
		AuthKeyId:           authKeyID,
		SourcePermAuthKeyId: &authKeyID,
		PeerType:            payload.PeerTypeUser,
		PeerId:              peerUserID,
		Message:             outboxes,
	})
	if err != nil {
		c.Logger.Errorf("messages.forwardMessages - msg error: self_user_id: %d, peer_id: %d, ids: %v, err: %v",
			selfUserID, peerUserID, in.Id, err)
		return nil, mapMsgSendError(err)
	}

	return updates, nil
}

func checkForwardMessagesUnsupportedFields(in *tg.TLMessagesForwardMessages) error {
	if in.WithMyScore {
		return tg.ErrInputRequestInvalid
	}
	if in.DropAuthor {
		return tg.ErrInputRequestInvalid
	}
	if in.DropMediaCaptions {
		return tg.ErrInputRequestInvalid
	}
	if in.AllowPaidFloodskip {
		return tg.ErrInputRequestInvalid
	}
	if in.TopMsgId != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.ReplyTo != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.ScheduleDate != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.ScheduleRepeatPeriod != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.SendAs != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.QuickReplyShortcut != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.Effect != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.VideoTimestamp != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.AllowPaidStars != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.SuggestedPost != nil {
		return tg.ErrInputRequestInvalid
	}
	return nil
}

func checkForwardMessagesIDs(ids []int32, randomIDs []int64) error {
	if len(ids) == 0 {
		return tg.ErrMessageIdInvalid
	}
	if len(ids) > maxForwardMessages {
		return tg.ErrInputRequestInvalid
	}
	if len(ids) != len(randomIDs) {
		return tg.ErrInputRequestInvalid
	}
	seen := make(map[int64]struct{}, len(randomIDs))
	for i := range ids {
		if ids[i] == 0 {
			return tg.ErrMessageIdInvalid
		}
		if randomIDs[i] == 0 {
			return tg.ErrRandomIdEmpty
		}
		if _, ok := seen[randomIDs[i]]; ok {
			return tg.ErrRandomIdDuplicate
		}
		seen[randomIDs[i]] = struct{}{}
	}
	return nil
}

func orderForwardSources(list *msg.VectorMessageBox, ids []int32) ([]*tg.TLMessage, error) {
	if list == nil {
		return nil, tg.ErrMessageIdInvalid
	}
	byID := make(map[int32]*tg.TLMessage, len(list.Datas))
	for _, box := range list.Datas {
		if box == nil {
			return nil, tg.ErrMessageIdInvalid
		}
		source, ok := box.Message.(*tg.TLMessage)
		if !ok || source == nil {
			return nil, tg.ErrMessageIdInvalid
		}
		id := source.Id
		if id == 0 {
			return nil, tg.ErrMessageIdInvalid
		}
		byID[id] = source
	}
	ordered := make([]*tg.TLMessage, 0, len(ids))
	for _, id := range ids {
		source := byID[id]
		if source == nil {
			return nil, tg.ErrMessageIdInvalid
		}
		ordered = append(ordered, source)
	}
	return ordered, nil
}

type forwardOutboxInput struct {
	RandomID        int64
	Background      bool
	Silent          bool
	Noforwards      bool
	FromUserID      int64
	PeerUserID      int64
	SourcePeer      tg.PeerClazz
	SourceMessageID int32
	Date            int32
	Source          *tg.TLMessage
	GroupedID       *int64
}

func buildForwardOutbox(in forwardOutboxInput) msg.OutboxMessageClazz {
	sourceMsgID := in.SourceMessageID
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:        true,
		Silent:     in.Silent,
		Noforwards: in.Noforwards,
		FromId:     tg.MakePeerUser(in.FromUserID),
		PeerId:     tg.MakePeerUser(in.PeerUserID),
		Date:       in.Date,
		Message:    in.Source.Message,
		Media:      in.Source.Media,
		Entities:   in.Source.Entities,
		GroupedId:  in.GroupedID,
		FwdFrom:    forwardHeaderForSource(in.Source, in.SourcePeer, sourceMsgID),
	})
	if in.PeerUserID == in.FromUserID {
		message.SavedPeerId = in.SourcePeer
	}
	return msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
		RandomId:   in.RandomID,
		Background: in.Background,
		Message:    message,
	})
}

func forwardHeaderForSource(source *tg.TLMessage, sourcePeer tg.PeerClazz, sourceMessageID int32) tg.MessageFwdHeaderClazz {
	if source != nil && source.FwdFrom != nil {
		return tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
			Imported:       source.FwdFrom.Imported,
			SavedOut:       source.FwdFrom.SavedOut,
			FromId:         source.FwdFrom.FromId,
			FromName:       source.FwdFrom.FromName,
			Date:           source.FwdFrom.Date,
			ChannelPost:    source.FwdFrom.ChannelPost,
			PostAuthor:     source.FwdFrom.PostAuthor,
			SavedFromPeer:  sourcePeer,
			SavedFromMsgId: &sourceMessageID,
			SavedFromId:    source.FwdFrom.SavedFromId,
			SavedFromName:  source.FwdFrom.SavedFromName,
			SavedDate:      source.FwdFrom.SavedDate,
			PsaType:        source.FwdFrom.PsaType,
		})
	}
	date := int32(0)
	var fromID tg.PeerClazz
	if source != nil {
		date = source.Date
		fromID = source.FromId
	}
	return tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
		FromId:         fromID,
		Date:           date,
		SavedFromPeer:  sourcePeer,
		SavedFromMsgId: &sourceMessageID,
	})
}
