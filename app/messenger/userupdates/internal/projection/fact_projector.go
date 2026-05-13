package projection

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/envelope"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type ProjectionMode = envelope.Mode

type ViewerContext struct {
	UserID           int64
	AuthKeyIDExclude *int64
}

type ProjectedUpdate struct {
	Update tg.UpdateClazz
}

func ProjectFacts(facts []payload.UpdateFactV1, viewer ViewerContext, mode ProjectionMode, pts int64, ptsCount int32, userMessageID int64) ([]ProjectedUpdate, error) {
	out := make([]ProjectedUpdate, 0, len(facts))
	for _, fact := range facts {
		decoded, err := payload.DecodeUpdateFact(fact)
		if err != nil {
			return nil, fmt.Errorf("%w: decode update fact kind=%s: %v", userupdates.ErrUserupdatesStorage, fact.Kind, err)
		}
		switch f := decoded.(type) {
		case payload.ChatParticipantsChangedFactV1:
			projected, err := ProjectChatParticipantsChangedFact(f)
			if err != nil {
				return nil, err
			}
			out = append(out, projected)
		case payload.NewMessageFactV1:
			projected, err := ProjectNewMessageFact(f, viewer, mode, pts, ptsCount, userMessageID)
			if err != nil {
				return nil, err
			}
			out = append(out, projected...)
		default:
			return nil, fmt.Errorf("%w: unsupported update fact kind=%s", userupdates.ErrUserupdatesStorage, fact.Kind)
		}
	}
	return out, nil
}

func ProjectNewMessageFact(f payload.NewMessageFactV1, viewer ViewerContext, mode ProjectionMode, pts int64, ptsCount int32, userMessageID int64) ([]ProjectedUpdate, error) {
	message, err := newMessageFactToTLMessage(f, viewer, userMessageID)
	if err != nil {
		return nil, err
	}
	pts32, err := int64ToInt32(pts, "pts")
	if err != nil {
		return nil, err
	}
	return []ProjectedUpdate{{
		Update: tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message:  message,
			Pts:      pts32,
			PtsCount: ptsCount,
		}),
	}}, nil
}

func ProjectChatParticipantsChangedFact(f payload.ChatParticipantsChangedFactV1) (ProjectedUpdate, error) {
	if f.ChatID <= 0 {
		return ProjectedUpdate{}, fmt.Errorf("%w: invalid chat id", userupdates.ErrUserupdatesStorage)
	}
	participants := make([]tg.ChatParticipantClazz, 0, len(f.Participants))
	for _, participant := range f.Participants {
		projected, err := chatParticipantFactToTL(participant)
		if err != nil {
			return ProjectedUpdate{}, err
		}
		participants = append(participants, projected)
	}
	return ProjectedUpdate{
		Update: tg.MakeTLUpdateChatParticipants(&tg.TLUpdateChatParticipants{
			Participants: tg.MakeTLChatParticipants(&tg.TLChatParticipants{
				ChatId:       f.ChatID,
				Participants: participants,
				Version:      f.Version,
			}),
		}),
	}, nil
}

func newMessageFactToTLMessage(f payload.NewMessageFactV1, viewer ViewerContext, userMessageID int64) (tg.MessageClazz, error) {
	messageID, err := messageIDInt32(userMessageID, "message id")
	if err != nil {
		return nil, err
	}
	if _, err := senderPeerFromFact(f.SenderUserID); err != nil {
		return nil, err
	}
	peer, err := peerFromFact(f.PeerType, f.PeerID)
	if err != nil {
		return nil, err
	}
	replyTo, err := replyHeaderFromUserMessageID(f.ReplyToUserMessageID)
	if err != nil {
		return nil, err
	}
	date, err := userupdatesDateInt32FromUnixSeconds(int64(f.Date), "message date")
	if err != nil {
		return nil, err
	}
	out := viewer.UserID == f.SenderUserID
	fromPeer := messageFromPeer(out, f.PeerType, f.SenderUserID)
	if f.ServiceAction != nil {
		action, err := messageServiceAction(f.ServiceAction)
		if err != nil {
			return nil, err
		}
		return tg.MakeTLMessageService(&tg.TLMessageService{
			Out:     out,
			Silent:  messageAttrsSilent(f.Attrs),
			Id:      messageID,
			FromId:  fromPeer,
			PeerId:  peer,
			ReplyTo: replyTo,
			Date:    date,
			Action:  action,
		}), nil
	}
	fwdFrom, err := messageForwardHeader(f.ForwardRef)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:         out,
		Silent:      messageAttrsSilent(f.Attrs),
		Noforwards:  messageAttrsNoforwards(f.Attrs),
		InvertMedia: messageAttrsInvertMedia(f.Attrs),
		Id:          messageID,
		FromId:      fromPeer,
		PeerId:      peer,
		FwdFrom:     fwdFrom,
		ReplyTo:     replyTo,
		Date:        date,
		Message:     f.MessageText,
		Media:       messageMedia(f.MediaRef),
		Entities:    messageEntities(f.Entities),
		GroupedId:   messageGroupedID(f.Attrs),
		TtlPeriod:   messageTTLPeriod(f.MediaRef),
	}), nil
}

func chatParticipantFactToTL(f payload.ChatParticipantFactV1) (tg.ChatParticipantClazz, error) {
	if f.UserID <= 0 {
		return nil, fmt.Errorf("%w: invalid chat participant user id=%d", userupdates.ErrUserupdatesStorage, f.UserID)
	}
	switch f.Role {
	case "creator":
		return tg.MakeTLChatParticipantCreator(&tg.TLChatParticipantCreator{
			UserId: f.UserID,
		}), nil
	case "admin":
		date, err := userupdatesDateInt32FromUnixSeconds(int64(f.Date), "chat participant date")
		if err != nil {
			return nil, err
		}
		return tg.MakeTLChatParticipantAdmin(&tg.TLChatParticipantAdmin{
			UserId:    f.UserID,
			InviterId: f.InviterUserID,
			Date:      date,
		}), nil
	case "member":
		date, err := userupdatesDateInt32FromUnixSeconds(int64(f.Date), "chat participant date")
		if err != nil {
			return nil, err
		}
		return tg.MakeTLChatParticipant(&tg.TLChatParticipant{
			UserId:    f.UserID,
			InviterId: f.InviterUserID,
			Date:      date,
		}), nil
	default:
		return nil, fmt.Errorf("%w: unsupported chat participant role=%s", userupdates.ErrUserupdatesStorage, f.Role)
	}
}

func senderPeerFromFact(senderUserID int64) (tg.PeerClazz, error) {
	if senderUserID <= 0 {
		return nil, fmt.Errorf("%w: invalid new message sender_user_id=%d", userupdates.ErrUserupdatesStorage, senderUserID)
	}
	return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: senderUserID}), nil
}

func peerFromFact(peerType int32, peerID int64) (tg.PeerClazz, error) {
	if peerID <= 0 {
		return nil, fmt.Errorf("%w: invalid new message peer_id=%d", userupdates.ErrUserupdatesStorage, peerID)
	}
	switch peerType {
	case payload.PeerTypeUser:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID}), nil
	case payload.PeerTypeChat:
		return tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: peerID}), nil
	case payload.PeerTypeChannel:
		return tg.MakeTLPeerChannel(&tg.TLPeerChannel{ChannelId: peerID}), nil
	default:
		return nil, fmt.Errorf("%w: unsupported new message peer_type=%d", userupdates.ErrUserupdatesStorage, peerType)
	}
}
