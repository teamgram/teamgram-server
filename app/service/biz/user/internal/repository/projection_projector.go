package repository

import (
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type projectionFacts struct {
	Users     map[int64]*projectionUserFact
	Contacts  map[contactKey]*projectionContactFact
	Privacies map[privacyKey][]tg.PrivacyRuleClazz
	Presences map[int64]*projectionPresenceFact
}

type projectionUserFact struct {
	User      tg.UserDataClazz
	IsBot     bool
	PhotoId   int64
	Usernames []tg.UsernameClazz
}

type projectionContactFact struct {
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Mutual        bool   `json:"mutual,omitempty"`
	CloseFriend   bool   `json:"close_friend,omitempty"`
	StoriesHidden bool   `json:"stories_hidden,omitempty"`
}

type projectionPresenceFact struct {
	LastSeenAt int64
	Expires    int32
}

const userProfilePhotoDcID int32 = 1

type contactKey struct {
	OwnerUserId   int64
	ContactUserId int64
}

type privacyKey struct {
	UserId  int64
	KeyType int32
}

func projectUserForViewer(viewerUserId, targetUserId int64, facts projectionFacts) tg.UserClazz {
	fact := facts.Users[targetUserId]
	if fact == nil || fact.User == nil {
		return nil
	}

	userData := fact.User.ToUserData()
	if userData.Deleted {
		return tg.MakeTLUser(&tg.TLUser{Id: targetUserId, Deleted: true})
	}

	if viewerUserId == targetUserId {
		return projectSelfUser(userData, fact, facts.Presences[targetUserId])
	}

	contact := facts.Contacts[contactKey{OwnerUserId: viewerUserId, ContactUserId: targetUserId}]
	reverseContact := facts.Contacts[contactKey{OwnerUserId: targetUserId, ContactUserId: viewerUserId}]

	firstName := userData.FirstName
	lastName := userData.LastName
	if contact != nil && (contact.FirstName != "" || contact.LastName != "") {
		firstName = contact.FirstName
		lastName = contact.LastName
	}

	user := projectBaseUser(userData, fact)
	user.FirstName = stringPtr(firstName)
	user.LastName = stringPtr(lastName)
	if contact != nil {
		user.Contact = true
		user.MutualContact = contact.Mutual || reverseContact != nil
		user.CloseFriend = contact.CloseFriend
		user.StoriesHidden = contact.StoriesHidden
		user.Phone = stringPtr(contact.Phone)
	} else if projectionAllowsPrivacy(viewerUserId, targetUserId, tg.PHONE_NUMBER, facts) {
		user.Phone = stringPtr(userData.Phone)
	}
	if projectionAllowsPrivacy(viewerUserId, targetUserId, tg.PROFILE_PHOTO, facts) {
		user.Photo = projectionUserProfilePhoto(fact)
	}
	if projectionAllowsPrivacy(viewerUserId, targetUserId, tg.STATUS_TIMESTAMP, facts) {
		user.Status = projectionUserStatus(facts.Presences[targetUserId])
	}
	return user
}

func projectSelfUser(userData *tg.UserData, fact *projectionUserFact, presence *projectionPresenceFact) tg.UserClazz {
	user := projectBaseUser(userData, fact)
	user.Self = true
	user.Contact = true
	user.MutualContact = true
	user.Phone = stringPtr(userData.Phone)
	user.Photo = projectionUserProfilePhoto(fact)
	user.Status = projectionUserStatus(presence)
	return user
}

func projectBaseUser(userData *tg.UserData, fact *projectionUserFact) *tg.TLUser {
	user := tg.MakeTLUser(&tg.TLUser{
		Id:                userData.Id,
		AccessHash:        &userData.AccessHash,
		FirstName:         stringPtr(userData.FirstName),
		LastName:          stringPtr(userData.LastName),
		Username:          stringPtr(userData.Username),
		Bot:               fact.IsBot || userData.Bot != nil,
		Verified:          userData.Verified,
		Restricted:        userData.Restricted,
		RestrictionReason: userData.RestrictionReason,
		Support:           userData.Support,
		Scam:              userData.Scam,
		Fake:              userData.Fake,
		Premium:           userData.Premium,
		EmojiStatus:       userData.EmojiStatus,
		Usernames:         projectionUsernames(userData, fact),
		Color:             userData.Color,
		ProfileColor:      userData.ProfileColor,
	})
	if userData.Bot != nil {
		bot := userData.Bot.ToBotData()
		user.BotChatHistory = bot.BotChatHistory
		user.BotNochats = bot.BotNochats
		user.BotInlineGeo = bot.BotInlineGeo
		user.BotInfoVersion = &bot.BotInfoVersion
		user.BotInlinePlaceholder = bot.BotInlinePlaceholder
		user.BotAttachMenu = bot.BotAttachMenu
		user.AttachMenuEnabled = bot.AttachMenuEnabled
		user.BotCanEdit = bot.BotCanEdit
		user.BotBusiness = bot.BotBusiness
		user.BotHasMainApp = bot.BotHasMainApp
		user.BotActiveUsers = bot.BotActiveUsers
	}
	return user
}

func projectionUserProfilePhoto(fact *projectionUserFact) tg.UserProfilePhotoClazz {
	if fact == nil || fact.PhotoId == 0 {
		return nil
	}
	return tg.MakeTLUserProfilePhoto(&tg.TLUserProfilePhoto{PhotoId: fact.PhotoId, DcId: userProfilePhotoDcID})
}

func projectionUserStatus(presence *projectionPresenceFact) tg.UserStatusClazz {
	if presence == nil {
		return nil
	}
	now := int32(time.Now().Unix())
	if presence.Expires > now {
		return tg.MakeTLUserStatusOnline(&tg.TLUserStatusOnline{Expires: presence.Expires})
	}
	if presence.LastSeenAt > 0 {
		return tg.MakeTLUserStatusOffline(&tg.TLUserStatusOffline{WasOnline: int32(presence.LastSeenAt)})
	}
	return nil
}

func projectionUsernames(userData *tg.UserData, fact *projectionUserFact) []tg.UsernameClazz {
	if fact != nil && len(fact.Usernames) > 0 {
		return append([]tg.UsernameClazz(nil), fact.Usernames...)
	}
	return tgUsernameList(userData.Username, true)
}

func projectionAllowsPrivacy(viewerUserId, targetUserId int64, keyType int32, facts projectionFacts) bool {
	if viewerUserId == targetUserId {
		return true
	}
	rules, ok := facts.Privacies[privacyKey{UserId: targetUserId, KeyType: keyType}]
	if !ok {
		rules = defaultPrivacyRules(keyType)
	}
	return evaluatePrivacyRules(rules, privacyEvaluationContext{
		PeerID:        viewerUserId,
		IsContact:     facts.Contacts[contactKey{OwnerUserId: targetUserId, ContactUserId: viewerUserId}] != nil,
		IsCloseFriend: contactIsCloseFriend(facts.Contacts[contactKey{OwnerUserId: targetUserId, ContactUserId: viewerUserId}]),
		IsPremium:     projectionUserIsPremium(facts.Users[viewerUserId]),
		IsBot:         projectionUserIsBot(facts.Users[viewerUserId]),
	})
}

func contactIsCloseFriend(contact *projectionContactFact) bool {
	return contact != nil && contact.CloseFriend
}

func projectionUserIsPremium(fact *projectionUserFact) bool {
	return fact != nil && fact.User != nil && fact.User.ToUserData().Premium
}

func projectionUserIsBot(fact *projectionUserFact) bool {
	return fact != nil && fact.IsBot
}
