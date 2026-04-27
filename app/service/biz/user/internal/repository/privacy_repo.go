package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetPrivacy(ctx context.Context, userID int64, keyType int32) (*userpb.VectorPrivacyRule, error) {
	if !isValidPrivacyKey(keyType) {
		return nil, userpb.ErrPrivacyKeyInvalid
	}
	privacyDO, err := r.model.UserPrivaciesModel.SelectPrivacy(ctx, userID, keyType)
	if err != nil {
		return nil, fmt.Errorf("%w: get privacy %d/%d: %w", userpb.ErrUserStorage, userID, keyType, err)
	}
	if privacyDO == nil || privacyDO.Rules == "" {
		return &userpb.VectorPrivacyRule{Datas: defaultPrivacyRules(keyType)}, nil
	}
	rules, err := decodePrivacyRules(privacyDO.Rules)
	if err != nil {
		return nil, fmt.Errorf("%w: decode privacy %d/%d: %w", userpb.ErrUserStorage, userID, keyType, err)
	}
	if rules == nil {
		rules = defaultPrivacyRules(keyType)
	}
	return &userpb.VectorPrivacyRule{Datas: rules}, nil
}

func (r *Repository) SetPrivacy(ctx context.Context, userID int64, keyType int32, rules []tg.PrivacyRuleClazz) error {
	if !isValidPrivacyKey(keyType) {
		return userpb.ErrPrivacyKeyInvalid
	}
	rulesData, err := json.Marshal(rules)
	if err != nil {
		return fmt.Errorf("%w: encode privacy %d/%d: %w", userpb.ErrUserStorage, userID, keyType, err)
	}
	if _, _, err := r.model.UserPrivaciesModel.InsertOrUpdate(ctx, &model.UserPrivacies{
		UserId:  userID,
		KeyType: keyType,
		Rules:   string(rulesData),
	}); err != nil {
		return fmt.Errorf("%w: set privacy %d/%d: %w", userpb.ErrUserStorage, userID, keyType, err)
	}
	if privacyAffectsUserData(keyType) {
		return r.invalidateUserDataCache(ctx, userID, "invalidate privacy user cache")
	}
	return nil
}

func (r *Repository) CheckPrivacy(ctx context.Context, userID int64, keyType int32, peerID int64) (bool, error) {
	_, err := r.GetPrivacy(ctx, userID, keyType)
	if err != nil {
		if errors.Is(err, userpb.ErrPrivacyKeyInvalid) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

type privacyRuleDTO struct {
	Name  string  `json:"_name"`
	ID    uint32  `json:"_id"`
	Users []int64 `json:"users"`
	Chats []int64 `json:"chats"`
}

func decodePrivacyRules(rulesData string) ([]tg.PrivacyRuleClazz, error) {
	var rawRules []privacyRuleDTO
	if err := json.Unmarshal([]byte(rulesData), &rawRules); err != nil {
		return nil, err
	}
	rules := make([]tg.PrivacyRuleClazz, 0, len(rawRules))
	for i := range rawRules {
		rule, ok := makePrivacyRuleFromDTO(rawRules[i])
		if ok {
			rules = append(rules, rule)
		}
	}
	return rules, nil
}

func makePrivacyRuleFromDTO(rule privacyRuleDTO) (tg.PrivacyRuleClazz, bool) {
	switch rule.Name {
	case tg.ClazzName_privacyValueAllowContacts:
		return tg.PrivacyValueAllowContactsClazz, true
	case tg.ClazzName_privacyValueAllowAll:
		return tg.PrivacyValueAllowAllClazz, true
	case tg.ClazzName_privacyValueAllowUsers:
		return tg.MakeTLPrivacyValueAllowUsers(&tg.TLPrivacyValueAllowUsers{Users: rule.Users}).ToPrivacyRule().Clazz, true
	case tg.ClazzName_privacyValueDisallowContacts:
		return tg.PrivacyValueDisallowContactsClazz, true
	case tg.ClazzName_privacyValueDisallowAll:
		return tg.PrivacyValueDisallowAllClazz, true
	case tg.ClazzName_privacyValueDisallowUsers:
		return tg.MakeTLPrivacyValueDisallowUsers(&tg.TLPrivacyValueDisallowUsers{Users: rule.Users}).ToPrivacyRule().Clazz, true
	case tg.ClazzName_privacyValueAllowChatParticipants:
		return tg.MakeTLPrivacyValueAllowChatParticipants(&tg.TLPrivacyValueAllowChatParticipants{Chats: rule.Chats}).ToPrivacyRule().Clazz, true
	case tg.ClazzName_privacyValueDisallowChatParticipants:
		return tg.MakeTLPrivacyValueDisallowChatParticipants(&tg.TLPrivacyValueDisallowChatParticipants{Chats: rule.Chats}).ToPrivacyRule().Clazz, true
	case tg.ClazzName_privacyValueAllowCloseFriends:
		return tg.PrivacyValueAllowCloseFriendsClazz, true
	case tg.ClazzName_privacyValueAllowPremium:
		return tg.PrivacyValueAllowPremiumClazz, true
	case tg.ClazzName_privacyValueAllowBots:
		return tg.PrivacyValueAllowBotsClazz, true
	case tg.ClazzName_privacyValueDisallowBots:
		return tg.PrivacyValueDisallowBotsClazz, true
	default:
		return nil, false
	}
}

func defaultPrivacyRules(keyType int32) []tg.PrivacyRuleClazz {
	if keyType == tg.PHONE_NUMBER {
		return []tg.PrivacyRuleClazz{tg.PrivacyValueDisallowAllClazz}
	}
	return []tg.PrivacyRuleClazz{tg.PrivacyValueAllowAllClazz}
}

func isValidPrivacyKey(keyType int32) bool {
	return keyType >= tg.STATUS_TIMESTAMP && keyType <= tg.NO_PAID_MESSAGES
}

func privacyAffectsUserData(keyType int32) bool {
	switch keyType {
	case tg.STATUS_TIMESTAMP, tg.PROFILE_PHOTO, tg.PHONE_NUMBER:
		return true
	default:
		return false
	}
}
