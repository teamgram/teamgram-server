package repository

import (
	"context"

	chatprojection "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chatprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetChatProjectionBundle(ctx context.Context, viewerUserIds []int64, targetChatIds []int64) (*ChatProjectionBundle, error) {
	viewers, targets, uniqueTargets := normalizeChatProjectionRequest(viewerUserIds, targetChatIds)
	if len(viewers) == 0 || len(targets) == 0 {
		return &ChatProjectionBundle{}, nil
	}

	rows, err := r.model.ChatsModel.FindListByIdList(ctx, uniqueTargets...)
	if err != nil {
		return nil, wrapStorage("chats.FindListByIdList", err)
	}

	mutableByID := make(map[int64]tg.MutableChatClazz, len(rows))
	for i := range rows {
		row := &rows[i]
		mutableByID[row.Id] = r.makeMutableChatFromRows(ctx, row, nil)
	}

	return buildChatProjectionBundle(viewers, targets, mutableByID)
}

func buildChatProjectionBundle(viewers []int64, targets []int64, mutableByID map[int64]tg.MutableChatClazz) (*ChatProjectionBundle, error) {
	if len(viewers) == 0 || len(targets) == 0 {
		return &ChatProjectionBundle{}, nil
	}

	orderedChats := make([]tg.MutableChatClazz, 0, len(targets))
	for _, targetID := range targets {
		orderedChats = append(orderedChats, mutableByID[targetID])
	}

	projected, err := chatprojection.ProjectMutableChatForViewers(orderedChats, viewers)
	if err != nil {
		return nil, err
	}

	viewerChats := make([]ViewerChats, 0, len(projected))
	for _, item := range projected {
		if item == nil {
			continue
		}
		viewerChats = append(viewerChats, ViewerChats{
			ViewerUserId: item.ViewerUserId,
			Chats:        item.Chats,
		})
	}

	return &ChatProjectionBundle{
		ViewerChats:    viewerChats,
		MissingChatIds: missingChatIds(targets, mutableByID),
	}, nil
}

func normalizeChatProjectionRequest(viewerUserIds []int64, targetChatIds []int64) ([]int64, []int64, []int64) {
	viewers := uniquePositiveInt64s(viewerUserIds)
	targets := positiveInt64s(targetChatIds)
	uniqueTargets := uniquePositiveInt64s(targets)
	return viewers, targets, uniqueTargets
}

func positiveInt64s(in []int64) []int64 {
	out := make([]int64, 0, len(in))
	for _, id := range in {
		if id > 0 {
			out = append(out, id)
		}
	}
	return out
}

func uniquePositiveInt64s(in []int64) []int64 {
	seen := make(map[int64]struct{}, len(in))
	out := make([]int64, 0, len(in))
	for _, id := range in {
		if id <= 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}
	return out
}

func missingChatIds(targets []int64, mutableByID map[int64]tg.MutableChatClazz) []int64 {
	seen := make(map[int64]struct{}, len(targets))
	out := make([]int64, 0)
	for _, id := range targets {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		if mutableByID[id] == nil || mutableByID[id].Chat == nil {
			out = append(out, id)
		}
	}
	return out
}
