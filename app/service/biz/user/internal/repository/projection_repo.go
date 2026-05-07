package repository

import (
	"context"
	"fmt"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetUserProjectionBundle(ctx context.Context, viewerUserIds []int64, targetUserIds []int64, withFacts bool) (*UserProjectionBundle, error) {
	cfg := normalizeProjectionConfig(r.projection)
	req, err := normalizeProjectionRequest(viewerUserIds, targetUserIds, withFacts, cfg)
	if err != nil {
		return nil, err
	}

	facts, err := r.loadProjectionFacts(ctx, req, cfg)
	if err != nil {
		return nil, err
	}
	for _, viewerID := range req.ViewerUserIds {
		if facts.Users[viewerID] == nil {
			return nil, userpb.ErrUserNotFound
		}
	}

	return buildProjectionBundle(req, facts), nil
}

func buildProjectionBundle(req normalizedProjectionRequest, facts projectionFacts) *UserProjectionBundle {
	viewerUsers := make([]ViewerUsers, 0, len(req.ViewerUserIds))
	for _, viewerID := range req.ViewerUserIds {
		users := make([]tg.UserClazz, 0, len(req.TargetUserIds))
		for _, targetID := range req.TargetUserIds {
			user := projectUserForViewer(viewerID, targetID, facts)
			if user != nil {
				users = append(users, user)
			}
		}
		viewerUsers = append(viewerUsers, ViewerUsers{ViewerUserId: viewerID, Users: users})
	}

	missing := make([]int64, 0)
	for _, targetID := range req.TargetUserIds {
		if facts.Users[targetID] == nil {
			missing = append(missing, targetID)
		}
	}

	var immutableFacts []tg.ImmutableUserClazz
	if req.WithFacts {
		immutableFacts = make([]tg.ImmutableUserClazz, 0, len(req.HydrateUserIds))
		for _, id := range req.HydrateUserIds {
			fact := facts.Users[id]
			if fact == nil || fact.User == nil {
				continue
			}
			immutableFacts = append(immutableFacts, tg.MakeTLImmutableUser(&tg.TLImmutableUser{
				User:             fact.User,
				LastSeenAt:       projectionLastSeenAt(facts.Presences[id]),
				Contacts:         projectionContactDataForOwner(id, facts.Contacts),
				ReverseContacts:  projectionReverseContactDataForOwner(id, facts.Contacts),
				KeysPrivacyRules: projectionPrivacyRulesForUser(id, facts.Privacies),
			}).ToImmutableUser())
		}
	}

	return &UserProjectionBundle{
		Facts:          immutableFacts,
		ViewerUsers:    viewerUsers,
		MissingUserIds: missing,
	}
}

func projectionLastSeenAt(presence *projectionPresenceFact) int64 {
	if presence == nil {
		return 0
	}
	return presence.LastSeenAt
}

func projectionContactDataForOwner(ownerID int64, contacts map[contactKey]*projectionContactFact) []tg.ContactDataClazz {
	out := make([]tg.ContactDataClazz, 0)
	for key, contact := range contacts {
		if key.OwnerUserId != ownerID || contact == nil {
			continue
		}
		out = append(out, projectionContactData(key.OwnerUserId, key.ContactUserId, contact))
	}
	return out
}

func projectionReverseContactDataForOwner(ownerID int64, contacts map[contactKey]*projectionContactFact) []tg.ContactDataClazz {
	out := make([]tg.ContactDataClazz, 0)
	for key, contact := range contacts {
		if key.ContactUserId != ownerID || contact == nil {
			continue
		}
		out = append(out, projectionContactData(key.OwnerUserId, key.ContactUserId, contact))
	}
	return out
}

func projectionContactData(ownerUserID, contactUserID int64, contact *projectionContactFact) tg.ContactDataClazz {
	return tg.MakeTLContactData(&tg.TLContactData{
		UserId:        ownerUserID,
		ContactUserId: contactUserID,
		FirstName:     stringPtr(contact.FirstName),
		LastName:      stringPtr(contact.LastName),
		MutualContact: contact.Mutual,
		Phone:         stringPtr(contact.Phone),
		CloseFriend:   contact.CloseFriend,
		StoriesHidden: contact.StoriesHidden,
	}).ToContactData()
}

func projectionPrivacyRulesForUser(userID int64, privacies map[privacyKey][]tg.PrivacyRuleClazz) []tg.PrivacyKeyRulesClazz {
	keys := []int32{tg.STATUS_TIMESTAMP, tg.PROFILE_PHOTO, tg.PHONE_NUMBER}
	out := make([]tg.PrivacyKeyRulesClazz, 0, len(keys))
	for _, keyType := range keys {
		rules, ok := privacies[privacyKey{UserId: userID, KeyType: keyType}]
		if !ok {
			continue
		}
		out = append(out, tg.MakeTLPrivacyKeyRules(&tg.TLPrivacyKeyRules{
			Key:   keyType,
			Rules: rules,
		}).ToPrivacyKeyRules())
	}
	return out
}

func wrapProjectionStorage(op string, err error) error {
	return fmt.Errorf("%w: %s: %w", userpb.ErrUserStorage, op, err)
}
