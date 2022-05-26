// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/media/media"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

const (
	chatKeyPrefix            = "chat"
	chatParticipantKeyPrefix = "chat_participant2"
)

type ChatCacheData struct {
	ChatData              *chatpb.ImmutableChat `json:"chat_data"`
	ChatParticipantIdList []int64               `json:"chat_participant_id_list"`
}

func genChatCacheKey(chatId int64) string {
	return fmt.Sprintf("%s_%d", chatKeyPrefix, chatId)
}

func (d *Dao) GetChatCacheKey(chatId int64) string {
	return fmt.Sprintf("%s_%d", chatKeyPrefix, chatId)
}

func genChatParticipantCacheKey(chatId, chatParticipantId int64) string {
	return fmt.Sprintf("%s_%d_%d", chatParticipantKeyPrefix, chatId, chatParticipantId)
}

func (d *Dao) GetChatParticipantCacheKey(chatId, chatParticipantId int64) string {
	return fmt.Sprintf("%s_%d_%d", chatParticipantKeyPrefix, chatId, chatParticipantId)
}

func (d *Dao) getChatData(ctx context.Context, chatId int64) (*ChatCacheData, error) {
	var (
		chatData = &ChatCacheData{}
	)

	err := d.CachedConn.QueryRow(
		ctx,
		chatData,
		genChatCacheKey(chatId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			do2, err := d.ChatsDAO.Select(ctx, chatId)
			if err != nil {
				return err
			} else if do2 == nil {
				return sqlc.ErrNotFound
			}
			cacheData := v.(*ChatCacheData)
			cacheData.ChatData = d.MakeImmutableChatByDO(do2)

			if do2.PhotoId != 0 {
				mr.FinishVoid(
					func() {
						cacheData.ChatData.Photo, _ = d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
							PhotoId: do2.PhotoId,
						})
					},
					func() {
						cacheData.ChatParticipantIdList, _ = d.ChatParticipantsDAO.SelectChatParticipantIdList(ctx, chatId)
					})
			} else {
				cacheData.ChatParticipantIdList, _ = d.ChatParticipantsDAO.SelectChatParticipantIdList(ctx, chatId)
			}

			return nil
		})

	if err != nil {
		if err != sqlc.ErrNotFound {
			err = mtproto.ErrChatIdInvalid
		}
		return nil, err
	}

	// TODO: cache
	if d.Plugin != nil {
		chatData.ChatData.CallActive, chatData.ChatData.CallNotEmpty = d.Plugin.GetChatCallActiveAndNotEmpty(ctx, 0, chatId)
		chatData.ChatData.Call = d.Plugin.GetChatGroupCall(ctx, 0, chatId)
	}

	return chatData, nil
}

func (d *Dao) getChatParticipantListByIdList(ctx context.Context, chatId int64, idList []int64) []*chatpb.ImmutableChatParticipant {
	participantList := make([]*chatpb.ImmutableChatParticipant, len(idList))

	mr.ForEach(
		func(source chan<- interface{}) {
			for i, v := range idList {
				source <- idxId{i, v}
			}
		},
		func(item interface{}) {
			idx := item.(idxId)
			var (
				p *chatpb.ImmutableChatParticipant
			)
			err2 := d.CachedConn.QueryRow(
				ctx,
				&p,
				genChatParticipantCacheKey(chatId, idx.id),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					do2, _ := d.ChatParticipantsDAO.SelectByParticipantId(ctx, chatId, idx.id)
					if do2 == nil {
						return sqlc.ErrNotFound
					}
					logx.WithContext(ctx).Infof("do2: %v", do2)
					*v.(**chatpb.ImmutableChatParticipant) = d.MakeImmutableChatParticipant(do2)
					return nil
				})

			logx.WithContext(ctx).Infof("do: %v", p)
			if err2 == nil {
				participantList[idx.idx] = p
			}
		})

	return removeAllNil(participantList)
}

func (d *Dao) GetMutableChat(ctx context.Context, chatId int64, id ...int64) (*chatpb.MutableChat, error) {
	var (
		participants []*chatpb.ImmutableChatParticipant
	)

	cacheData, err := d.getChatData(ctx, chatId)
	if err != nil {
		return nil, err
	}
	if len(id) == 0 {
		participants = d.getChatParticipantListByIdList(ctx, chatId, cacheData.ChatParticipantIdList)
	} else {
		participants = d.getChatParticipantListByIdList(ctx, chatId, id)
	}

	// participants, err = d.getImmutableChatParticipants(ctx, immutableChat, id...)
	// if err != nil {
	// 	return nil, err
	// }

	return chatpb.MakeTLMutableChat(&chatpb.MutableChat{
		Chat:             cacheData.ChatData,
		ChatParticipants: participants,
	}).To_MutableChat(), nil
}

func (d *Dao) PutMutableChat(ctx context.Context, chat *chatpb.MutableChat) error {
	cacheData := &ChatCacheData{
		ChatData:              chat.Chat,
		ChatParticipantIdList: make([]int64, len(chat.ChatParticipants)),
	}

	for i, v := range chat.ChatParticipants {
		cacheData.ChatParticipantIdList[i] = v.UserId
	}

	mr.ForEach(
		func(source chan<- interface{}) {
			source <- &kv{genChatCacheKey(chat.Id()), cacheData}
			for _, v := range chat.ChatParticipants {
				source <- &kv{genChatParticipantCacheKey(v.ChatId, v.UserId), v}
			}
		},
		func(item interface{}) {
			kv2 := item.(*kv)
			d.CachedConn.SetCache(ctx, kv2.k, kv2.v)
		})

	return nil
}
