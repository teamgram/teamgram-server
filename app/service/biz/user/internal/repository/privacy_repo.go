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
		if !isNotFound(err) {
			return nil, fmt.Errorf("%w: get privacy %d/%d: %w", userpb.ErrUserStorage, userID, keyType, err)
		}
		privacyDO = nil
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
	rulesData, err := encodePrivacyRules(rules)
	if err != nil {
		return fmt.Errorf("%w: encode privacy %d/%d: %w", userpb.ErrUserStorage, userID, keyType, err)
	}
	if _, _, err := r.model.UserPrivaciesModel.InsertOrUpdate(ctx, &model.UserPrivacies{
		UserId:  userID,
		KeyType: keyType,
		Rules:   rulesData,
	}); err != nil {
		return fmt.Errorf("%w: set privacy %d/%d: %w", userpb.ErrUserStorage, userID, keyType, err)
	}
	if privacyAffectsUserData(keyType) {
		r.invalidateProjectionPrivacyCache(ctx, userID)
		return r.invalidateUserDataCache(ctx, userID, "invalidate privacy user cache")
	}
	return nil
}

func (r *Repository) CheckPrivacy(ctx context.Context, userID int64, keyType int32, peerID int64) (bool, error) {
	rules, err := r.GetPrivacy(ctx, userID, keyType)
	if err != nil {
		if errors.Is(err, userpb.ErrPrivacyKeyInvalid) {
			return false, nil
		}
		return false, err
	}
	eval, err := r.buildPrivacyEvaluationContext(ctx, userID, peerID)
	if err != nil {
		return false, fmt.Errorf("%w: check privacy context %d/%d: %w", userpb.ErrUserStorage, userID, peerID, err)
	}
	return evaluatePrivacyRules(rules.Datas, eval), nil
}

type privacyEvaluationContext struct {
	PeerID        int64
	IsContact     bool
	IsCloseFriend bool
	IsPremium     bool
	IsBot         bool
}

func (r *Repository) buildPrivacyEvaluationContext(ctx context.Context, userID, peerID int64) (privacyEvaluationContext, error) {
	eval := privacyEvaluationContext{PeerID: peerID}
	if userID == 0 || peerID == 0 {
		return eval, nil
	}
	contactDO, err := r.model.UserContactsModel.SelectContact(ctx, userID, peerID)
	if err != nil {
		if !isNotFound(err) {
			return eval, fmt.Errorf("select privacy contact %d/%d: %w", userID, peerID, err)
		}
		contactDO = nil
	}
	if contactDO != nil {
		eval.IsContact = true
		eval.IsCloseFriend = contactDO.CloseFriend
	}
	peerDO, err := r.model.UsersModel.FindOne(ctx, peerID)
	if err != nil {
		if isNotFound(err) {
			return eval, nil
		}
		return eval, fmt.Errorf("select privacy peer %d: %w", peerID, err)
	}
	eval.IsPremium = peerDO.Premium
	eval.IsBot = peerDO.IsBot
	return eval, nil
}

func evaluatePrivacyRules(rules []tg.PrivacyRuleClazz, eval privacyEvaluationContext) bool {
	base := tg.ALLOW_CONTACTS
	for _, rule := range rules {
		switch rule.(type) {
		case *tg.TLPrivacyValueAllowAll:
			base = tg.ALLOW_ALL
		case *tg.TLPrivacyValueAllowContacts:
			base = tg.ALLOW_CONTACTS
		case *tg.TLPrivacyValueDisallowAll:
			base = tg.DISALLOW_ALL
		}
	}

	allowed := false
	switch base {
	case tg.ALLOW_ALL:
		allowed = true
	case tg.ALLOW_CONTACTS:
		allowed = eval.IsContact
	case tg.DISALLOW_ALL:
		allowed = false
	}

	for _, rule := range rules {
		switch r := rule.(type) {
		case *tg.TLPrivacyValueAllowUsers:
			if containsInt64(r.Users, eval.PeerID) {
				allowed = true
			}
		case *tg.TLPrivacyValueDisallowUsers:
			if containsInt64(r.Users, eval.PeerID) {
				allowed = false
			}
		case *tg.TLPrivacyValueAllowContacts:
			if eval.IsContact {
				allowed = true
			}
		case *tg.TLPrivacyValueDisallowContacts:
			if eval.IsContact {
				allowed = false
			}
		case *tg.TLPrivacyValueAllowCloseFriends:
			if eval.IsCloseFriend {
				allowed = true
			}
		case *tg.TLPrivacyValueAllowPremium:
			if eval.IsPremium {
				allowed = true
			}
		case *tg.TLPrivacyValueAllowBots:
			if eval.IsBot {
				allowed = true
			}
		case *tg.TLPrivacyValueDisallowBots:
			if eval.IsBot {
				allowed = false
			}
		case *tg.TLPrivacyValueAllowChatParticipants, *tg.TLPrivacyValueDisallowChatParticipants:
			// User repository currently has no local chat participant reader.
		}
	}
	return allowed
}

type privacyRulesStorageDTO struct {
	SchemaVersion int              `json:"schema_version"`
	Rules         []privacyRuleDTO `json:"rules"`
}

type privacyRuleDTO struct {
	Type       string  `json:"type"`
	LegacyName string  `json:"_name,omitempty"`
	LegacyID   uint32  `json:"_id,omitempty"`
	Users      []int64 `json:"users,omitempty"`
	Chats      []int64 `json:"chats,omitempty"`
}

func encodePrivacyRules(rules []tg.PrivacyRuleClazz) (string, error) {
	rawRules := make([]privacyRuleDTO, 0, len(rules))
	for _, rule := range rules {
		rawRule, ok := privacyRuleToDTO(rule)
		if ok {
			rawRules = append(rawRules, rawRule)
		}
	}
	data, err := json.Marshal(privacyRulesStorageDTO{
		SchemaVersion: 1,
		Rules:         rawRules,
	})
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func decodePrivacyRules(rulesData string) ([]tg.PrivacyRuleClazz, error) {
	var payload privacyRulesStorageDTO
	if err := json.Unmarshal([]byte(rulesData), &payload); err == nil && payload.SchemaVersion > 0 {
		return privacyRulesFromDTO(payload.Rules), nil
	}

	var legacyRules []privacyRuleDTO
	if err := json.Unmarshal([]byte(rulesData), &legacyRules); err != nil {
		return nil, err
	}
	return privacyRulesFromDTO(legacyRules), nil
}

func privacyRulesFromDTO(rawRules []privacyRuleDTO) []tg.PrivacyRuleClazz {
	rules := make([]tg.PrivacyRuleClazz, 0, len(rawRules))
	for i := range rawRules {
		rule, ok := makePrivacyRuleFromDTO(rawRules[i])
		if ok {
			rules = append(rules, rule)
		}
	}
	return rules
}

func makePrivacyRuleFromDTO(rule privacyRuleDTO) (tg.PrivacyRuleClazz, bool) {
	ruleType := rule.Type
	if ruleType == "" {
		ruleType = rule.LegacyName
	}
	switch ruleType {
	case "allow_contacts", tg.ClazzName_privacyValueAllowContacts:
		return tg.PrivacyValueAllowContactsClazz, true
	case "allow_all", tg.ClazzName_privacyValueAllowAll:
		return tg.PrivacyValueAllowAllClazz, true
	case "allow_users", tg.ClazzName_privacyValueAllowUsers:
		return tg.MakeTLPrivacyValueAllowUsers(&tg.TLPrivacyValueAllowUsers{Users: rule.Users}).ToPrivacyRule().Clazz, true
	case "disallow_contacts", tg.ClazzName_privacyValueDisallowContacts:
		return tg.PrivacyValueDisallowContactsClazz, true
	case "disallow_all", tg.ClazzName_privacyValueDisallowAll:
		return tg.PrivacyValueDisallowAllClazz, true
	case "disallow_users", tg.ClazzName_privacyValueDisallowUsers:
		return tg.MakeTLPrivacyValueDisallowUsers(&tg.TLPrivacyValueDisallowUsers{Users: rule.Users}).ToPrivacyRule().Clazz, true
	case "allow_chat_participants", tg.ClazzName_privacyValueAllowChatParticipants:
		return tg.MakeTLPrivacyValueAllowChatParticipants(&tg.TLPrivacyValueAllowChatParticipants{Chats: rule.Chats}).ToPrivacyRule().Clazz, true
	case "disallow_chat_participants", tg.ClazzName_privacyValueDisallowChatParticipants:
		return tg.MakeTLPrivacyValueDisallowChatParticipants(&tg.TLPrivacyValueDisallowChatParticipants{Chats: rule.Chats}).ToPrivacyRule().Clazz, true
	case "allow_close_friends", tg.ClazzName_privacyValueAllowCloseFriends:
		return tg.PrivacyValueAllowCloseFriendsClazz, true
	case "allow_premium", tg.ClazzName_privacyValueAllowPremium:
		return tg.PrivacyValueAllowPremiumClazz, true
	case "allow_bots", tg.ClazzName_privacyValueAllowBots:
		return tg.PrivacyValueAllowBotsClazz, true
	case "disallow_bots", tg.ClazzName_privacyValueDisallowBots:
		return tg.PrivacyValueDisallowBotsClazz, true
	default:
		return nil, false
	}
}

func privacyRuleToDTO(rule tg.PrivacyRuleClazz) (privacyRuleDTO, bool) {
	if rule == nil {
		return privacyRuleDTO{}, false
	}
	switch rule.PrivacyRuleClazzName() {
	case tg.ClazzName_privacyValueAllowContacts:
		return privacyRuleDTO{Type: "allow_contacts"}, true
	case tg.ClazzName_privacyValueAllowAll:
		return privacyRuleDTO{Type: "allow_all"}, true
	case tg.ClazzName_privacyValueAllowUsers:
		r, ok := rule.(*tg.TLPrivacyValueAllowUsers)
		if !ok {
			return privacyRuleDTO{}, false
		}
		return privacyRuleDTO{Type: "allow_users", Users: r.Users}, true
	case tg.ClazzName_privacyValueDisallowContacts:
		return privacyRuleDTO{Type: "disallow_contacts"}, true
	case tg.ClazzName_privacyValueDisallowAll:
		return privacyRuleDTO{Type: "disallow_all"}, true
	case tg.ClazzName_privacyValueDisallowUsers:
		r, ok := rule.(*tg.TLPrivacyValueDisallowUsers)
		if !ok {
			return privacyRuleDTO{}, false
		}
		return privacyRuleDTO{Type: "disallow_users", Users: r.Users}, true
	case tg.ClazzName_privacyValueAllowChatParticipants:
		r, ok := rule.(*tg.TLPrivacyValueAllowChatParticipants)
		if !ok {
			return privacyRuleDTO{}, false
		}
		return privacyRuleDTO{Type: "allow_chat_participants", Chats: r.Chats}, true
	case tg.ClazzName_privacyValueDisallowChatParticipants:
		r, ok := rule.(*tg.TLPrivacyValueDisallowChatParticipants)
		if !ok {
			return privacyRuleDTO{}, false
		}
		return privacyRuleDTO{Type: "disallow_chat_participants", Chats: r.Chats}, true
	case tg.ClazzName_privacyValueAllowCloseFriends:
		return privacyRuleDTO{Type: "allow_close_friends"}, true
	case tg.ClazzName_privacyValueAllowPremium:
		return privacyRuleDTO{Type: "allow_premium"}, true
	case tg.ClazzName_privacyValueAllowBots:
		return privacyRuleDTO{Type: "allow_bots"}, true
	case tg.ClazzName_privacyValueDisallowBots:
		return privacyRuleDTO{Type: "disallow_bots"}, true
	default:
		return privacyRuleDTO{}, false
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
