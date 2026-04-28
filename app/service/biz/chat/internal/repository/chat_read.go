package repository

import (
	"context"
	"errors"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type UserChatIDList struct {
	UserID     int64
	ChatIDList []int64
}

func (r *Repository) GetMutableChat(ctx context.Context, chatID int64, participantIDs ...int64) (*tg.MutableChat, error) {
	row, err := r.model.ChatsModel.FindOne(ctx, chatID)
	if err != nil {
		if isNotFound(err) {
			return nil, chatpb.ErrChatNotFound
		}
		return nil, wrapStorage("chats.FindOne", err)
	}

	var participantRows []model.ChatParticipants
	if len(participantIDs) == 0 {
		participantRows, err = r.model.ChatParticipantsModel.SelectList(ctx, chatID)
	} else {
		participantRows, err = r.model.ChatParticipantsModel.SelectListByParticipantIdList(ctx, chatID, participantIDs)
	}
	if err != nil {
		return nil, wrapStorage("chat_participants.SelectList", err)
	}

	return r.makeMutableChatFromRows(ctx, row, participantRows), nil
}

func (r *Repository) GetMutableChatByLink(ctx context.Context, link string) (*tg.MutableChat, error) {
	hash := chatpb.NormalizeInviteHash(link)
	if hash == "" {
		return nil, chatpb.ErrInviteHashInvalid
	}
	invite, err := r.model.ChatInvitesModel.SelectByLink(ctx, hash)
	if err != nil {
		if isNotFound(err) {
			return nil, chatpb.ErrInviteHashInvalid
		}
		return nil, wrapStorage("chat_invites.SelectByLink", err)
	}
	if invite == nil || invite.ChatId == 0 {
		return nil, chatpb.ErrInviteHashInvalid
	}
	mChat, err := r.GetMutableChat(ctx, invite.ChatId)
	if err != nil {
		if errors.Is(err, chatpb.ErrChatNotFound) {
			return nil, chatpb.ErrChatNotFound
		}
		return nil, err
	}
	return mChat, nil
}

func (r *Repository) GetExcludeParticipantsMutableChat(ctx context.Context, chatID int64) (*tg.MutableChat, error) {
	row, err := r.model.ChatsModel.FindOne(ctx, chatID)
	if err != nil {
		if isNotFound(err) {
			return nil, chatpb.ErrChatNotFound
		}
		return nil, wrapStorage("chats.FindOne", err)
	}
	return r.makeMutableChatFromRows(ctx, row, nil), nil
}

func (r *Repository) GetChatBySelfID(ctx context.Context, chatID, selfID int64) (*tg.MutableChat, error) {
	return r.GetMutableChat(ctx, chatID, selfID)
}

func (r *Repository) GetChatListByIDList(ctx context.Context, ids []int64) ([]*tg.MutableChat, error) {
	rows, err := r.model.ChatsModel.FindListByIdList(ctx, ids...)
	if err != nil {
		return nil, wrapStorage("chats.FindListByIdList", err)
	}
	if len(rows) == 0 {
		return []*tg.MutableChat{}, nil
	}

	orderedRows := orderChatRowsByIDs(ids, rows)
	out := make([]*tg.MutableChat, 0, len(orderedRows))
	for _, row := range orderedRows {
		if row == nil {
			continue
		}
		participantRows, err := r.model.ChatParticipantsModel.SelectList(ctx, row.Id)
		if err != nil {
			if isNotFound(err) {
				continue
			}
			return nil, wrapStorage("chat_participants.SelectList", err)
		}
		out = append(out, r.makeMutableChatFromRows(ctx, row, participantRows))
	}
	return out, nil
}

func (r *Repository) GetChatParticipantIDList(ctx context.Context, chatID int64) ([]int64, error) {
	rows, err := r.model.ChatParticipantsModel.SelectList(ctx, chatID)
	if err != nil {
		return nil, wrapStorage("chat_participants.SelectList", err)
	}
	ids := make([]int64, 0, len(rows))
	for i := range rows {
		if isListableChatParticipantState(rows[i].State) {
			ids = append(ids, rows[i].UserId)
		}
	}
	return ids, nil
}

func (r *Repository) GetUsersChatIDList(ctx context.Context, userIDs []int64) ([]UserChatIDList, error) {
	rows, err := r.model.ChatParticipantsModel.SelectUsersChatIdList(ctx, userIDs)
	if err != nil {
		return nil, wrapStorage("chat_participants.SelectUsersChatIdList", err)
	}
	return groupUserChatIDRows(rows), nil
}

func (r *Repository) GetMyChatList(ctx context.Context, userID int64, isCreator bool) ([]*tg.MutableChat, error) {
	var (
		ids []int64
		err error
	)
	if isCreator {
		ids, err = r.model.ChatParticipantsModel.SelectMyAdminList(ctx, userID)
	} else {
		ids, err = r.model.ChatParticipantsModel.SelectMyAllList(ctx, userID)
	}
	if err != nil {
		return nil, wrapStorage("chat_participants.SelectMyList", err)
	}

	out := make([]*tg.MutableChat, 0, len(ids))
	for _, id := range ids {
		mChat, err := r.GetMutableChat(ctx, id, userID)
		if err != nil {
			if errors.Is(err, chatpb.ErrChatNotFound) {
				continue
			}
			return nil, err
		}
		if mChat != nil {
			out = append(out, mChat)
		}
	}
	return out, nil
}

func (r *Repository) Search(ctx context.Context, selfID int64, q string, offset int64, limit int32) ([]*tg.MutableChat, error) {
	ids, err := r.model.ChatsModel.SearchByQueryString(ctx, q, limit)
	if err != nil {
		return nil, wrapStorage("chats.SearchByQueryString", err)
	}

	out := make([]*tg.MutableChat, 0, len(ids))
	for _, id := range ids {
		mChat, err := r.GetExcludeParticipantsMutableChat(ctx, id)
		if err != nil {
			if errors.Is(err, chatpb.ErrChatNotFound) {
				continue
			}
			return nil, err
		}
		if mChat != nil {
			out = append(out, mChat)
		}
	}
	return out, nil
}

func (r *Repository) makeMutableChatFromRows(ctx context.Context, row *model.Chats, participantRows []model.ChatParticipants) *tg.MutableChat {
	if row == nil {
		return makeMutableChat(nil, nil)
	}

	var photo tg.PhotoClazz = tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: row.PhotoId})
	if row.PhotoId != 0 && r.mediaReader != nil {
		// Photo projection is best-effort; media lookup failure falls back to photoEmpty.
		if mediaPhoto, err := r.mediaReader.GetChatPhoto(ctx, row.PhotoId); err == nil && mediaPhoto != nil && mediaPhoto.Clazz != nil {
			photo = mediaPhoto.Clazz
		}
	}

	immutable := makeImmutableChat(row, photo)
	if immutable != nil && r.plugin != nil {
		immutable.CallActive, immutable.CallNotEmpty = r.plugin.GetChatCallActiveAndNotEmpty(ctx, 0, row.Id)
		immutable.Call = r.plugin.GetChatGroupCall(ctx, 0, row.Id)
	}

	participants := make([]*tg.ImmutableChatParticipant, 0, len(participantRows))
	for i := range participantRows {
		participants = append(participants, makeImmutableChatParticipant(&participantRows[i]))
	}
	return makeMutableChat(immutable, participants)
}

func orderChatRowsByIDs(ids []int64, rows []model.Chats) []*model.Chats {
	byID := make(map[int64]*model.Chats, len(rows))
	for i := range rows {
		byID[rows[i].Id] = &rows[i]
	}

	out := make([]*model.Chats, 0, len(rows))
	for _, id := range ids {
		if row := byID[id]; row != nil {
			out = append(out, row)
		}
	}
	return out
}

func groupUserChatIDRows(rows []model.ChatParticipants) []UserChatIDList {
	out := make([]UserChatIDList, 0, len(rows))
	for i := range rows {
		found := -1
		for j := range out {
			if out[j].UserID == rows[i].UserId {
				found = j
				break
			}
		}
		if found < 0 {
			out = append(out, UserChatIDList{
				UserID:     rows[i].UserId,
				ChatIDList: []int64{rows[i].ChatId},
			})
			continue
		}
		out[found].ChatIDList = append(out[found].ChatIDList, rows[i].ChatId)
	}
	return out
}

func isListableChatParticipantState(state int32) bool {
	return state == chatpb.ChatMemberStateNormal || state == chatpb.ChatMemberStateMigrated
}
