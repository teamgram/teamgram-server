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
	if err := r.loadProjectionPrivacyFacts(ctx, req.HydrateUserIds, cfg, facts); err != nil {
		return projectionFacts{}, err
	}
	if err := r.loadProjectionPresenceFacts(ctx, req.HydrateUserIds, cfg, facts); err != nil {
		return projectionFacts{}, err
	}
	if err := r.loadProjectionContactFacts(ctx, req.ViewerUserIds, req.TargetUserIds, cfg, facts); err != nil {
		return projectionFacts{}, err
	}
	return facts, nil
}

func (r *Repository) loadProjectionUserFacts(ctx context.Context, ids []int64, cfg ProjectionConfig, facts projectionFacts) error {
	missIDs := make([]int64, 0, len(ids))
	cacheKeys := make([]string, 0, len(ids))
	for _, id := range ids {
		key := projectionFactsCacheKey(id)
		cacheKeys = append(cacheKeys, key)
	}
	cacheHits := getProjectionComponentCaches[projectionUserCacheDTO](r, ctx, cacheKeys)
	for _, id := range ids {
		key := projectionFactsCacheKey(id)
		dto, ok := cacheHits[key]
		if ok && dto.UserID == id {
			facts.Users[id] = projectionUserFactFromCacheDTO(dto)
			continue
		}
		missIDs = append(missIDs, id)
	}
	if len(missIDs) == 0 {
		return nil
	}

	dbLoadedIDs := make([]int64, 0, len(missIDs))
	for _, chunk := range chunkInt64s(missIDs, cfg.SQLInChunkSize) {
		users, err := r.model.UsersModel.SelectUsersByIdList(ctx, chunk)
		if err != nil {
			return fmt.Errorf("%w: projection load users: %w", userpb.ErrUserStorage, err)
		}
		for i := range users {
			facts.Users[users[i].Id] = projectionUserFactFromModel(&users[i])
			dbLoadedIDs = append(dbLoadedIDs, users[i].Id)
		}
	}
	if err := r.loadProjectionBotFacts(ctx, dbLoadedIDs, cfg, facts); err != nil {
		return err
	}
	if err := r.loadProjectionUsernameFacts(ctx, dbLoadedIDs, cfg, facts); err != nil {
		return err
	}
	for _, id := range dbLoadedIDs {
		fact := facts.Users[id]
		if fact == nil {
			continue
		}
		r.setProjectionComponentCache(ctx, projectionFactsCacheKey(id), projectionUserCacheDTOFromFact(id, fact))
	}
	return nil
}

func (r *Repository) loadProjectionContactFacts(ctx context.Context, viewerIds []int64, targetIds []int64, cfg ProjectionConfig, facts projectionFacts) error {
	if len(viewerIds) == 0 || len(targetIds) == 0 {
		return nil
	}
	ownerIds := unionInt64s(viewerIds, targetIds)
	contactIds := unionInt64s(targetIds, viewerIds)
	contactSet := int64Set(contactIds)
	fallbackOwnerIds := ownerIds
	if cfg.ContactMapCacheEnabled {
		fallbackOwnerIds = make([]int64, 0, len(ownerIds))
		cacheKeys := make([]string, 0, len(ownerIds))
		for _, ownerID := range ownerIds {
			key := projectionContactMapCacheKey(ownerID)
			cacheKeys = append(cacheKeys, key)
		}
		cacheHits := getProjectionComponentCaches[projectionContactMapCacheDTO](r, ctx, cacheKeys)
		for _, ownerID := range ownerIds {
			key := projectionContactMapCacheKey(ownerID)
			dto, ok := cacheHits[key]
			if !ok || dto.OwnerUserID != ownerID || len(dto.Contacts) > cfg.ContactMapMaxEntries {
				fallbackOwnerIds = append(fallbackOwnerIds, ownerID)
				continue
			}
			for contactID, contact := range dto.Contacts {
				if _, ok := contactSet[contactID]; !ok {
					continue
				}
				c := contact
				facts.Contacts[contactKey{OwnerUserId: ownerID, ContactUserId: contactID}] = &c
			}
		}
	}
	return r.loadProjectionContactOwnerMaps(ctx, fallbackOwnerIds, contactSet, cfg, facts)
}

func (r *Repository) loadProjectionContactOwnerMaps(ctx context.Context, ownerIds []int64, contactSet map[int64]bool, cfg ProjectionConfig, facts projectionFacts) error {
	if len(ownerIds) == 0 {
		return nil
	}
	contactsByOwner := make(map[int64]map[int64]projectionContactFact, len(ownerIds))
	for _, ownerID := range ownerIds {
		contactsByOwner[ownerID] = make(map[int64]projectionContactFact)
	}
	for _, ownerChunk := range chunkInt64s(ownerIds, cfg.SQLInChunkSize) {
		contacts, err := r.model.UserContactsModel.SelectListByOwnerList(ctx, ownerChunk)
		if err != nil {
			return fmt.Errorf("%w: projection load contact maps: %w", userpb.ErrUserStorage, err)
		}
		for i := range contacts {
			c := projectionContactFactFromModel(&contacts[i])
			contactsByOwner[contacts[i].OwnerUserId][contacts[i].ContactUserId] = c
			if _, ok := contactSet[contacts[i].ContactUserId]; ok {
				c2 := c
				facts.Contacts[contactKey{OwnerUserId: contacts[i].OwnerUserId, ContactUserId: contacts[i].ContactUserId}] = &c2
			}
		}
	}
	if cfg.ContactMapCacheEnabled {
		for _, ownerID := range ownerIds {
			ownerContacts := contactsByOwner[ownerID]
			if len(ownerContacts) > cfg.ContactMapMaxEntries {
				r.deleteProjectionComponentCaches(ctx, projectionContactMapCacheKey(ownerID))
				continue
			}
			r.setProjectionComponentCache(ctx, projectionContactMapCacheKey(ownerID), projectionContactMapCacheDTO{
				OwnerUserID: ownerID,
				Contacts:    ownerContacts,
			})
		}
	}
	return nil
}

func (r *Repository) loadProjectionPrivacyFacts(ctx context.Context, ids []int64, cfg ProjectionConfig, facts projectionFacts) error {
	keyTypes := []int32{tg.STATUS_TIMESTAMP, tg.PROFILE_PHOTO, tg.PHONE_NUMBER}
	missIDs := make([]int64, 0, len(ids))
	cacheKeys := make([]string, 0, len(ids))
	for _, id := range ids {
		cacheKeys = append(cacheKeys, projectionPrivacyCacheKey(id))
	}
	cacheHits := getProjectionComponentCaches[projectionPrivacyCacheDTO](r, ctx, cacheKeys)
	for _, id := range ids {
		key := projectionPrivacyCacheKey(id)
		dto, ok := cacheHits[key]
		if ok && dto.UserID == id {
			addProjectionPrivacyCacheDTO(dto, facts)
			continue
		}
		missIDs = append(missIDs, id)
	}
	if len(missIDs) == 0 {
		return nil
	}

	loaded := make(map[int64]map[int32][]tg.PrivacyRuleClazz, len(missIDs))
	for _, id := range missIDs {
		loaded[id] = make(map[int32][]tg.PrivacyRuleClazz)
	}
	for _, chunk := range chunkInt64s(missIDs, cfg.SQLInChunkSize) {
		rows, err := r.model.UserPrivaciesModel.SelectUsersPrivacyList(ctx, chunk, keyTypes)
		if err != nil {
			return fmt.Errorf("%w: projection load privacies: %w", userpb.ErrUserStorage, err)
		}
		for i := range rows {
			rules, err := decodePrivacyRules(rows[i].Rules)
			if err != nil {
				return fmt.Errorf("%w: projection decode privacy %d/%d: %w", userpb.ErrUserStorage, rows[i].UserId, rows[i].KeyType, err)
			}
			if loaded[rows[i].UserId] == nil {
				loaded[rows[i].UserId] = make(map[int32][]tg.PrivacyRuleClazz)
			}
			loaded[rows[i].UserId][rows[i].KeyType] = rules
			facts.Privacies[privacyKey{UserId: rows[i].UserId, KeyType: rows[i].KeyType}] = rules
		}
	}
	for _, id := range missIDs {
		r.setProjectionComponentCache(ctx, projectionPrivacyCacheKey(id), projectionPrivacyCacheDTOFromRules(id, loaded[id]))
	}
	return nil
}

func (r *Repository) loadProjectionPresenceFacts(ctx context.Context, ids []int64, cfg ProjectionConfig, facts projectionFacts) error {
	missIDs := make([]int64, 0, len(ids))
	cacheKeys := make([]string, 0, len(ids))
	for _, id := range ids {
		cacheKeys = append(cacheKeys, projectionPresenceCacheKey(id))
	}
	cacheHits := getProjectionComponentCaches[projectionPresenceCacheDTO](r, ctx, cacheKeys)
	for _, id := range ids {
		key := projectionPresenceCacheKey(id)
		dto, ok := cacheHits[key]
		if ok && dto.UserID == id {
			if dto.HasPresence {
				facts.Presences[id] = &projectionPresenceFact{LastSeenAt: dto.LastSeenAt, Expires: dto.Expires}
			}
			continue
		}
		missIDs = append(missIDs, id)
	}
	if len(missIDs) == 0 {
		return nil
	}

	loaded := make(map[int64]*projectionPresenceFact, len(missIDs))
	for _, chunk := range chunkInt64s(missIDs, cfg.SQLInChunkSize) {
		rows, err := r.model.UserPresencesModel.SelectList(ctx, chunk)
		if err != nil {
			return fmt.Errorf("%w: projection load presences: %w", userpb.ErrUserStorage, err)
		}
		for i := range rows {
			presence := &projectionPresenceFact{LastSeenAt: rows[i].LastSeenAt, Expires: rows[i].Expires}
			loaded[rows[i].UserId] = presence
			facts.Presences[rows[i].UserId] = presence
		}
	}
	for _, id := range missIDs {
		r.setProjectionComponentCache(ctx, projectionPresenceCacheKey(id), projectionPresenceCacheDTOFromFact(id, loaded[id]))
	}
	return nil
}

func (r *Repository) loadProjectionBotFacts(ctx context.Context, ids []int64, cfg ProjectionConfig, facts projectionFacts) error {
	botIDs := make([]int64, 0)
	for _, id := range ids {
		if projectionUserIsBot(facts.Users[id]) {
			botIDs = append(botIDs, id)
		}
	}
	for _, chunk := range chunkInt64s(botIDs, cfg.SQLInChunkSize) {
		rows, err := r.model.BotsModel.SelectByBotIdList(ctx, chunk)
		if err != nil {
			return fmt.Errorf("%w: projection load bots: %w", userpb.ErrUserStorage, err)
		}
		for i := range rows {
			fact := facts.Users[rows[i].BotId]
			if fact == nil || fact.User == nil {
				continue
			}
			userData := fact.User.ToUserData()
			userData.Bot = botDataFromModel(&rows[i])
		}
	}
	return nil
}

func (r *Repository) loadProjectionUsernameFacts(ctx context.Context, ids []int64, cfg ProjectionConfig, facts projectionFacts) error {
	for _, chunk := range chunkInt64s(ids, cfg.SQLInChunkSize) {
		rows, err := r.model.UsernameModel.SelectListByUserIdList(ctx, chunk)
		if err != nil {
			return fmt.Errorf("%w: projection load usernames: %w", userpb.ErrUserStorage, err)
		}
		for i := range rows {
			fact := facts.Users[rows[i].PeerId]
			if fact == nil {
				continue
			}
			fact.Usernames = append(fact.Usernames, usernameClazzFromModel(&rows[i]))
		}
	}
	return nil
}

func addProjectionContactFacts(contacts []model.UserContacts, facts projectionFacts) {
	for i := range contacts {
		c := contacts[i]
		fact := projectionContactFactFromModel(&c)
		facts.Contacts[contactKey{OwnerUserId: c.OwnerUserId, ContactUserId: c.ContactUserId}] = &fact
	}
}

func projectionContactFactFromModel(c *model.UserContacts) projectionContactFact {
	return projectionContactFact{
		FirstName:     c.ContactFirstName,
		LastName:      c.ContactLastName,
		Phone:         c.ContactPhone,
		Mutual:        c.Mutual,
		CloseFriend:   c.CloseFriend,
		StoriesHidden: c.StoriesHidden,
	}
}

func projectionUserFactFromModel(user *model.Users) *projectionUserFact {
	if user == nil {
		return nil
	}
	return &projectionUserFact{
		User:    userDataFromModel(user),
		IsBot:   user.IsBot,
		PhotoId: user.PhotoId,
	}
}

func projectionUserFactFromCacheDTO(dto projectionUserCacheDTO) *projectionUserFact {
	userData := tg.MakeTLUserData(&tg.TLUserData{
		Id:                dto.UserID,
		AccessHash:        dto.AccessHash,
		UserType:          dto.UserType,
		SceretKeyId:       dto.SecretKeyID,
		FirstName:         dto.FirstName,
		LastName:          dto.LastName,
		Username:          dto.Username,
		Phone:             dto.Phone,
		CountryCode:       dto.CountryCode,
		Verified:          dto.Verified,
		Support:           dto.Support,
		Scam:              dto.Scam,
		Fake:              dto.Fake,
		About:             stringPtr(dto.About),
		Restricted:        dto.Restricted,
		RestrictionReason: restrictionReasonsFromCacheDTO(dto.RestrictionReasons),
		Deleted:           dto.Deleted,
		Premium:           dto.Premium,
		EmojiStatus:       emojiStatusFromCacheDTO(dto.EmojiStatusDocumentID, dto.EmojiStatusUntil),
		Color:             peerColorFromCacheDTO(dto.Color, dto.ColorBackgroundEmojiID),
		ProfileColor:      peerColorFromCacheDTO(dto.ProfileColor, dto.ProfileColorBackgroundEmojiID),
		StoriesMaxId:      dto.StoriesMaxID,
		Birthday:          dto.Birthday,
		PersonalChannelId: dto.PersonalChannelID,
		Bot:               botDataFromCacheDTO(dto.Bot),
	}).ToUserData()
	return &projectionUserFact{
		User:      userData,
		IsBot:     dto.IsBot,
		PhotoId:   dto.PhotoID,
		Usernames: usernameClazzesFromCacheDTO(dto.Usernames),
	}
}

func projectionUserCacheDTOFromFact(id int64, fact *projectionUserFact) projectionUserCacheDTO {
	userData := fact.User.ToUserData()
	dto := projectionUserCacheDTO{
		UserID:            id,
		AccessHash:        userData.AccessHash,
		UserType:          userData.UserType,
		SecretKeyID:       userData.SceretKeyId,
		FirstName:         userData.FirstName,
		LastName:          userData.LastName,
		Username:          userData.Username,
		Phone:             userData.Phone,
		CountryCode:       userData.CountryCode,
		Verified:          userData.Verified,
		Support:           userData.Support,
		Scam:              userData.Scam,
		Fake:              userData.Fake,
		Premium:           userData.Premium,
		Restricted:        userData.Restricted,
		Deleted:           userData.Deleted,
		StoriesMaxID:      userData.StoriesMaxId,
		Birthday:          userData.Birthday,
		PersonalChannelID: userData.PersonalChannelId,
		PhotoID:           fact.PhotoId,
		IsBot:             fact.IsBot,
		Usernames:         usernameCacheDTOsFromClazzes(fact.Usernames),
	}
	if userData.About != nil {
		dto.About = *userData.About
	}
	if emoji, ok := userData.EmojiStatus.(*tg.TLEmojiStatus); ok {
		dto.EmojiStatusDocumentID = emoji.DocumentId
		if emoji.Until != nil {
			dto.EmojiStatusUntil = *emoji.Until
		}
	}
	if color, ok := userData.Color.(*tg.TLPeerColor); ok {
		if color.Color != nil {
			dto.Color = *color.Color
		}
		if color.BackgroundEmojiId != nil {
			dto.ColorBackgroundEmojiID = *color.BackgroundEmojiId
		}
	}
	if color, ok := userData.ProfileColor.(*tg.TLPeerColor); ok {
		if color.Color != nil {
			dto.ProfileColor = *color.Color
		}
		if color.BackgroundEmojiId != nil {
			dto.ProfileColorBackgroundEmojiID = *color.BackgroundEmojiId
		}
	}
	if userData.Bot != nil {
		dto.Bot = botCacheDTOFromData(userData.Bot.ToBotData())
	}
	return dto
}

func botDataFromModel(bot *model.Bots) tg.BotDataClazz {
	if bot == nil {
		return nil
	}
	return tg.MakeTLBotData(&tg.TLBotData{
		Id:                   bot.BotId,
		BotType:              bot.BotType,
		Creator:              bot.CreatorUserId,
		Token:                bot.Token,
		Description:          bot.Description,
		BotChatHistory:       bot.BotChatHistory,
		BotNochats:           bot.BotNochats,
		BotInlineGeo:         bot.BotInlineGeo,
		BotInfoVersion:       bot.BotInfoVersion,
		BotInlinePlaceholder: stringPtr(bot.BotInlinePlaceholder),
		BotAttachMenu:        bot.BotAttachMenu,
		AttachMenuEnabled:    bot.AttachMenuEnabled,
		BotCanEdit:           bot.BotCanEdit,
		BotBusiness:          bot.BotBusiness,
		BotHasMainApp:        bot.BotHasMainApp,
		BotActiveUsers:       int32Ptr(bot.BotActiveUsers),
	}).ToBotData()
}

func botDataFromCacheDTO(dto *projectionBotCacheDTO) tg.BotDataClazz {
	if dto == nil {
		return nil
	}
	return tg.MakeTLBotData(&tg.TLBotData{
		Id:                   dto.ID,
		BotType:              dto.BotType,
		Creator:              dto.Creator,
		Token:                dto.Token,
		Description:          dto.Description,
		BotChatHistory:       dto.BotChatHistory,
		BotNochats:           dto.BotNochats,
		BotInlineGeo:         dto.BotInlineGeo,
		BotInfoVersion:       dto.BotInfoVersion,
		BotInlinePlaceholder: dto.BotInlinePlaceholder,
		BotAttachMenu:        dto.BotAttachMenu,
		AttachMenuEnabled:    dto.AttachMenuEnabled,
		BotCanEdit:           dto.BotCanEdit,
		BotBusiness:          dto.BotBusiness,
		BotHasMainApp:        dto.BotHasMainApp,
		BotActiveUsers:       dto.BotActiveUsers,
	}).ToBotData()
}

func botCacheDTOFromData(bot *tg.BotData) *projectionBotCacheDTO {
	if bot == nil {
		return nil
	}
	return &projectionBotCacheDTO{
		ID:                   bot.Id,
		BotType:              bot.BotType,
		Creator:              bot.Creator,
		Token:                bot.Token,
		Description:          bot.Description,
		BotChatHistory:       bot.BotChatHistory,
		BotNochats:           bot.BotNochats,
		BotInlineGeo:         bot.BotInlineGeo,
		BotInfoVersion:       bot.BotInfoVersion,
		BotInlinePlaceholder: bot.BotInlinePlaceholder,
		BotAttachMenu:        bot.BotAttachMenu,
		AttachMenuEnabled:    bot.AttachMenuEnabled,
		BotCanEdit:           bot.BotCanEdit,
		BotBusiness:          bot.BotBusiness,
		BotHasMainApp:        bot.BotHasMainApp,
		BotActiveUsers:       bot.BotActiveUsers,
	}
}

func usernameClazzFromModel(username *model.Username) tg.UsernameClazz {
	if username == nil || username.Username == "" {
		return nil
	}
	return tg.MakeTLUsername(&tg.TLUsername{
		Username: username.Username,
		Editable: username.Editable,
		Active:   username.Active,
	})
}

func usernameClazzesFromCacheDTO(in []projectionUsernameCacheDTO) []tg.UsernameClazz {
	out := make([]tg.UsernameClazz, 0, len(in))
	for _, dto := range in {
		if dto.Username == "" {
			continue
		}
		out = append(out, tg.MakeTLUsername(&tg.TLUsername{
			Username: dto.Username,
			Editable: dto.Editable,
			Active:   dto.Active,
		}))
	}
	return out
}

func usernameCacheDTOsFromClazzes(in []tg.UsernameClazz) []projectionUsernameCacheDTO {
	out := make([]projectionUsernameCacheDTO, 0, len(in))
	for _, username := range in {
		if username == nil {
			continue
		}
		out = append(out, projectionUsernameCacheDTO{
			Username: username.Username,
			Editable: username.Editable,
			Active:   username.Active,
		})
	}
	return out
}

func addProjectionPrivacyCacheDTO(dto projectionPrivacyCacheDTO, facts projectionFacts) {
	for keyType, ruleDTOs := range dto.Rules {
		facts.Privacies[privacyKey{UserId: dto.UserID, KeyType: keyType}] = privacyRulesFromDTO(ruleDTOs)
	}
}

func projectionPrivacyCacheDTOFromRules(userID int64, rules map[int32][]tg.PrivacyRuleClazz) projectionPrivacyCacheDTO {
	dto := projectionPrivacyCacheDTO{
		UserID: userID,
		Rules:  make(map[int32][]privacyRuleDTO, len(rules)),
	}
	for keyType, keyRules := range rules {
		dto.Rules[keyType] = privacyRuleDTOsFromRules(keyRules)
	}
	return dto
}

func privacyRuleDTOsFromRules(rules []tg.PrivacyRuleClazz) []privacyRuleDTO {
	out := make([]privacyRuleDTO, 0, len(rules))
	for _, rule := range rules {
		dto, ok := privacyRuleToDTO(rule)
		if ok {
			out = append(out, dto)
		}
	}
	return out
}

func projectionPresenceCacheDTOFromFact(userID int64, presence *projectionPresenceFact) projectionPresenceCacheDTO {
	dto := projectionPresenceCacheDTO{UserID: userID}
	if presence == nil {
		return dto
	}
	dto.HasPresence = true
	dto.LastSeenAt = presence.LastSeenAt
	dto.Expires = presence.Expires
	return dto
}

func restrictionReasonsFromCacheDTO(in []projectionRestrictionCacheDTO) []tg.RestrictionReasonClazz {
	return []tg.RestrictionReasonClazz{}
}

func emojiStatusFromCacheDTO(documentID int64, until int32) tg.EmojiStatusClazz {
	if documentID == 0 {
		return nil
	}
	return tg.MakeTLEmojiStatus(&tg.TLEmojiStatus{
		DocumentId: documentID,
		Until:      int32Ptr(until),
	})
}

func peerColorFromCacheDTO(color int32, backgroundEmojiID int64) tg.PeerColorClazz {
	if color == 0 && backgroundEmojiID == 0 {
		return nil
	}
	return tg.MakeTLPeerColor(&tg.TLPeerColor{
		Color:             int32Ptr(color),
		BackgroundEmojiId: int64Ptr(backgroundEmojiID),
	})
}

func unionInt64s(first []int64, rest ...[]int64) []int64 {
	out := make([]int64, 0, len(first))
	seen := make(map[int64]struct{}, len(first))
	for _, list := range append([][]int64{first}, rest...) {
		for _, id := range list {
			if id <= 0 {
				continue
			}
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			out = append(out, id)
		}
	}
	return out
}

func int32Ptr(v int32) *int32 {
	if v == 0 {
		return nil
	}
	return &v
}

func int64Ptr(v int64) *int64 {
	if v == 0 {
		return nil
	}
	return &v
}
