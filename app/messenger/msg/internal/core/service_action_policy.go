package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type serviceActionPolicy int

const (
	serviceActionDisplayOnly serviceActionPolicy = iota
	serviceActionRequiresOwnerFact
	serviceActionRejectFromGenericSend
)

var serviceActionPolicyByClazzName = map[string]serviceActionPolicy{
	tg.ClazzName_messageActionEmpty:                         serviceActionDisplayOnly,
	tg.ClazzName_messageActionChatCreate:                    serviceActionRequiresOwnerFact,
	tg.ClazzName_messageActionChatEditTitle:                 serviceActionDisplayOnly,
	tg.ClazzName_messageActionChatEditPhoto:                 serviceActionDisplayOnly,
	tg.ClazzName_messageActionChatDeletePhoto:               serviceActionDisplayOnly,
	tg.ClazzName_messageActionChatAddUser:                   serviceActionRequiresOwnerFact,
	tg.ClazzName_messageActionChatDeleteUser:                serviceActionRequiresOwnerFact,
	tg.ClazzName_messageActionChatJoinedByLink:              serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionChannelCreate:                 serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionChatMigrateTo:                 serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionChannelMigrateFrom:            serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPinMessage:                    serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionHistoryClear:                  serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGameScore:                     serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPaymentSentMe:                 serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPaymentSent:                   serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPhoneCall:                     serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionScreenshotTaken:               serviceActionDisplayOnly,
	tg.ClazzName_messageActionCustomAction:                  serviceActionDisplayOnly,
	tg.ClazzName_messageActionBotAllowed:                    serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSecureValuesSentMe:            serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSecureValuesSent:              serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionContactSignUp:                 serviceActionDisplayOnly,
	tg.ClazzName_messageActionGeoProximityReached:           serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGroupCall:                     serviceActionDisplayOnly,
	tg.ClazzName_messageActionInviteToGroupCall:             serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSetMessagesTTL:                serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGroupCallScheduled:            serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSetChatTheme:                  serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionChatJoinedByRequest:           serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionWebViewDataSentMe:             serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionWebViewDataSent:               serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGiftPremium:                   serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionTopicCreate:                   serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionTopicEdit:                     serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSuggestProfilePhoto:           serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionRequestedPeer:                 serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSetChatWallPaper:              serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGiftCode:                      serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGiveawayLaunch:                serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGiveawayResults:               serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionBoostApply:                    serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionRequestedPeerSentMe:           serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPaymentRefunded:               serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGiftStars:                     serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPrizeStars:                    serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionStarGift:                      serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionStarGiftUnique:                serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPaidMessagesRefunded:          serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPaidMessagesPrice:             serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionConferenceCall:                serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionTodoCompletions:               serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionTodoAppendTasks:               serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSuggestedPostApproval:         serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSuggestedPostSuccess:          serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSuggestedPostRefund:           serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionGiftTon:                       serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionSuggestBirthday:               serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionStarGiftPurchaseOffer:         serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionStarGiftPurchaseOfferDeclined: serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionNewCreatorPending:             serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionChangeCreator:                 serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionNoForwardsToggle:              serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionNoForwardsRequest:             serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPollAppendAnswer:              serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionPollDeleteAnswer:              serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionManagedBotCreated:             serviceActionRejectFromGenericSend,
	tg.ClazzName_messageActionBizDataRaw:                    serviceActionRejectFromGenericSend,
}

func validateServiceActionPolicy(action tg.MessageActionClazz, peerType int32, peerID, actorUserID int64, attachFacts []payload.UpdateFactV1) error {
	if action == nil {
		return fmt.Errorf("%w: missing service action", msg.ErrSendStateConflict)
	}
	policy, ok := serviceActionPolicyByClazzName[action.MessageActionClazzName()]
	if !ok {
		return fmt.Errorf("%w: unclassified service action %T", msg.ErrSendStateConflict, action)
	}
	switch policy {
	case serviceActionDisplayOnly:
		return nil
	case serviceActionRequiresOwnerFact:
		if err := validateChatParticipantsOwnerFact(action, peerType, peerID, actorUserID, attachFacts); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("%w: service action %s rejected from generic send", msg.ErrSendStateConflict, action.MessageActionClazzName())
	}
}

func validateChatParticipantsOwnerFact(action tg.MessageActionClazz, peerType int32, peerID, actorUserID int64, attachFacts []payload.UpdateFactV1) error {
	if peerType != payload.PeerTypeChat || peerID <= 0 || actorUserID <= 0 {
		return fmt.Errorf("%w: service action %s requires valid chat peer and actor", msg.ErrSendStateConflict, action.MessageActionClazzName())
	}
	participants, ok, err := chatParticipantsChangedFactFromAttachFacts(attachFacts, peerID, actorUserID)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("%w: service action %s requires chat participants fact", msg.ErrSendStateConflict, action.MessageActionClazzName())
	}
	participantIDs := chatParticipantIDSet(participants.Participants)
	switch a := action.(type) {
	case *tg.TLMessageActionChatCreate:
		if !allUserIDsPresent(participantIDs, a.Users) {
			return fmt.Errorf("%w: chat create service action users are not covered by chat participants fact", msg.ErrSendStateConflict)
		}
	case *tg.TLMessageActionChatAddUser:
		if len(a.Users) == 0 || !allUserIDsPresent(participantIDs, a.Users) {
			return fmt.Errorf("%w: chat add user service action users are not covered by chat participants fact", msg.ErrSendStateConflict)
		}
	case *tg.TLMessageActionChatDeleteUser:
		if a.UserId <= 0 {
			return fmt.Errorf("%w: chat delete user service action has invalid user_id", msg.ErrSendStateConflict)
		}
		if _, ok := participantIDs[a.UserId]; ok {
			return fmt.Errorf("%w: chat delete user service action user_id is still present in chat participants fact", msg.ErrSendStateConflict)
		}
	default:
		return fmt.Errorf("%w: service action %s requires unsupported owner fact contract", msg.ErrSendStateConflict, action.MessageActionClazzName())
	}
	return nil
}

func chatParticipantsChangedFactFromAttachFacts(attachFacts []payload.UpdateFactV1, chatID, actorUserID int64) (payload.ChatParticipantsChangedFactV1, bool, error) {
	for _, fact := range attachFacts {
		if fact.Kind != payload.FactKindChatParticipantsChanged {
			continue
		}
		decoded, err := payload.DecodeUpdateFact(fact)
		if err != nil {
			return payload.ChatParticipantsChangedFactV1{}, false, fmt.Errorf("%w: decode chat participants fact: %v", msg.ErrSendStateConflict, err)
		}
		participants, ok := decoded.(payload.ChatParticipantsChangedFactV1)
		if !ok {
			return payload.ChatParticipantsChangedFactV1{}, false, fmt.Errorf("%w: attach fact %s decoded to %T", msg.ErrSendStateConflict, fact.Kind, decoded)
		}
		if participants.ChatID == chatID && participants.ActorUserID == actorUserID {
			return participants, true, nil
		}
	}
	return payload.ChatParticipantsChangedFactV1{}, false, nil
}

func chatParticipantIDSet(participants []payload.ChatParticipantFactV1) map[int64]struct{} {
	out := make(map[int64]struct{}, len(participants))
	for _, participant := range participants {
		if participant.UserID > 0 {
			out[participant.UserID] = struct{}{}
		}
	}
	return out
}

func allUserIDsPresent(set map[int64]struct{}, userIDs []int64) bool {
	for _, userID := range userIDs {
		if userID <= 0 {
			return false
		}
		if _, ok := set[userID]; !ok {
			return false
		}
	}
	return true
}
