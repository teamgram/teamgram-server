package envelope

import (
	"context"
	"fmt"
	"sort"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/dependencies"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func BuildUpdatesWithDependencies(ctx context.Context, projector PeerObjectProjector, viewerUserID int64, in Input) (*tg.Updates, error) {
	deps := dependencies.CollectUpdates(in.Updates)
	if len(deps.ChannelIDs) > 0 {
		return nil, fmt.Errorf("%w: project dependency channels: unsupported ids %v", userupdates.ErrUserupdatesStorage, deps.ChannelIDs)
	}
	if len(deps.UserIDs) > 0 {
		if projector == nil {
			return nil, fmt.Errorf("%w: project dependency users: projector is nil", userupdates.ErrUserupdatesStorage)
		}
		users, err := projector.ProjectUsers(ctx, viewerUserID, deps.UserIDs)
		if err != nil {
			return nil, err
		}
		if len(users) == 0 {
			return nil, fmt.Errorf("%w: project dependency users: empty projection for ids %v", userupdates.ErrUserupdatesStorage, deps.UserIDs)
		}
		in.Users = mergeUsersByID(in.Users, users)
		if missing := missingUserIDs(deps.UserIDs, in.Users); len(missing) > 0 {
			return nil, fmt.Errorf("%w: project dependency users: missing ids %v", userupdates.ErrUserupdatesStorage, missing)
		}
	}
	if len(deps.ChatIDs) > 0 {
		if projector == nil {
			return nil, fmt.Errorf("%w: project dependency chats: projector is nil", userupdates.ErrUserupdatesStorage)
		}
		chats, err := projector.ProjectChats(ctx, viewerUserID, deps.ChatIDs)
		if err != nil {
			return nil, err
		}
		if len(chats) == 0 {
			return nil, fmt.Errorf("%w: project dependency chats: empty projection for ids %v", userupdates.ErrUserupdatesStorage, deps.ChatIDs)
		}
		in.Chats = mergeChatsByID(in.Chats, chats)
		if missing := missingChatIDs(deps.ChatIDs, in.Chats); len(missing) > 0 {
			return nil, fmt.Errorf("%w: project dependency chats: missing ids %v", userupdates.ErrUserupdatesStorage, missing)
		}
	}
	return BuildUpdates(in)
}

func mergeUsersByID(base []tg.UserClazz, projected []tg.UserClazz) []tg.UserClazz {
	return mergeByID(base, projected, userID)
}

func mergeChatsByID(base []tg.ChatClazz, projected []tg.ChatClazz) []tg.ChatClazz {
	return mergeByID(base, projected, chatID)
}

func mergeByID[T any](base []T, projected []T, idFn func(T) int64) []T {
	out := make([]T, 0, len(base)+len(projected))
	seen := make(map[int64]struct{}, len(base)+len(projected))
	for _, item := range base {
		out = append(out, item)
		if id := idFn(item); id > 0 {
			seen[id] = struct{}{}
		}
	}
	for _, item := range projected {
		id := idFn(item)
		if id > 0 {
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
		}
		out = append(out, item)
	}
	return out
}

func missingUserIDs(ids []int64, users []tg.UserClazz) []int64 {
	present := make(map[int64]struct{}, len(users))
	for _, user := range users {
		if id := userID(user); id > 0 {
			present[id] = struct{}{}
		}
	}
	return missingIDs(ids, present)
}

func missingChatIDs(ids []int64, chats []tg.ChatClazz) []int64 {
	present := make(map[int64]struct{}, len(chats))
	for _, chat := range chats {
		if id := chatID(chat); id > 0 {
			present[id] = struct{}{}
		}
	}
	return missingIDs(ids, present)
}

func missingIDs(ids []int64, present map[int64]struct{}) []int64 {
	if len(ids) == 0 {
		return nil
	}
	missingSet := make(map[int64]struct{})
	for _, id := range ids {
		if id <= 0 {
			continue
		}
		if _, ok := present[id]; !ok {
			missingSet[id] = struct{}{}
		}
	}
	if len(missingSet) == 0 {
		return nil
	}
	missing := make([]int64, 0, len(missingSet))
	for id := range missingSet {
		missing = append(missing, id)
	}
	sort.Slice(missing, func(i, j int) bool { return missing[i] < missing[j] })
	return missing
}

func userID(user tg.UserClazz) int64 {
	switch u := user.(type) {
	case *tg.TLUser:
		if u == nil {
			return 0
		}
		return u.Id
	case *tg.TLUserEmpty:
		if u == nil {
			return 0
		}
		return u.Id
	default:
		return 0
	}
}

func chatID(chat tg.ChatClazz) int64 {
	switch c := chat.(type) {
	case *tg.TLChat:
		if c == nil {
			return 0
		}
		return c.Id
	case *tg.TLChatEmpty:
		if c == nil {
			return 0
		}
		return c.Id
	case *tg.TLChatForbidden:
		if c == nil {
			return 0
		}
		return c.Id
	default:
		return 0
	}
}
