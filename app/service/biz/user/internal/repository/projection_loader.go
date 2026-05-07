package repository

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) loadProjectionFacts(ctx context.Context, req normalizedProjectionRequest, cfg ProjectionConfig) (projectionFacts, error) {
	facts := projectionFacts{
		Users:     make(map[int64]*projectionUserFact, len(req.HydrateUserIds)),
		Contacts:  make(map[contactKey]*projectionContactFact),
		Privacies: make(map[privacyKey][]tg.PrivacyRuleClazz),
		Presences: make(map[int64]*projectionPresenceFact),
	}
	if len(req.HydrateUserIds) == 0 {
		return facts, nil
	}

	if err := r.loadProjectionUserFacts(ctx, req.HydrateUserIds, cfg, facts); err != nil {
		return projectionFacts{}, err
	}
	if err := r.loadProjectionContactFacts(ctx, req.ViewerUserIds, req.TargetUserIds, cfg, facts); err != nil {
		return projectionFacts{}, err
	}
	return facts, nil
}

func (r *Repository) loadProjectionUserFacts(ctx context.Context, ids []int64, cfg ProjectionConfig, facts projectionFacts) error {
	for _, chunk := range chunkInt64s(ids, cfg.SQLInChunkSize) {
		users, err := r.model.UsersModel.SelectUsersByIdList(ctx, chunk)
		if err != nil {
			return fmt.Errorf("%w: projection load users: %w", userpb.ErrUserStorage, err)
		}
		for i := range users {
			facts.Users[users[i].Id] = projectionUserFactFromModel(&users[i])
		}
	}
	return nil
}

func (r *Repository) loadProjectionContactFacts(ctx context.Context, viewerIds []int64, targetIds []int64, cfg ProjectionConfig, facts projectionFacts) error {
	if len(viewerIds) == 0 || len(targetIds) == 0 {
		return nil
	}
	if err := r.loadProjectionContactDirection(ctx, viewerIds, targetIds, cfg, facts); err != nil {
		return err
	}
	return r.loadProjectionContactDirection(ctx, targetIds, viewerIds, cfg, facts)
}

func (r *Repository) loadProjectionContactDirection(ctx context.Context, ownerIds []int64, contactIds []int64, cfg ProjectionConfig, facts projectionFacts) error {
	for _, ownerChunk := range chunkInt64s(ownerIds, cfg.SQLInChunkSize) {
		for _, contactChunk := range chunkInt64s(contactIds, cfg.SQLInChunkSize) {
			contacts, err := r.model.UserContactsModel.SelectListByOwnerListAndContactList(ctx, ownerChunk, contactChunk)
			if err != nil {
				return fmt.Errorf("%w: projection load contacts: %w", userpb.ErrUserStorage, err)
			}
			addProjectionContactFacts(contacts, facts)
		}
	}
	return nil
}

func addProjectionContactFacts(contacts []model.UserContacts, facts projectionFacts) {
	for i := range contacts {
		c := contacts[i]
		facts.Contacts[contactKey{OwnerUserId: c.OwnerUserId, ContactUserId: c.ContactUserId}] = &projectionContactFact{
			FirstName:     c.ContactFirstName,
			LastName:      c.ContactLastName,
			Phone:         c.ContactPhone,
			Mutual:        c.Mutual,
			CloseFriend:   c.CloseFriend,
			StoriesHidden: c.StoriesHidden,
		}
	}
}

func projectionUserFactFromModel(user *model.Users) *projectionUserFact {
	if user == nil {
		return nil
	}
	return &projectionUserFact{
		User:  userDataFromModel(user),
		IsBot: user.IsBot,
	}
}
