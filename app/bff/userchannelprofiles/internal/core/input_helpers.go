package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func requireSelfID(c *UserChannelProfilesCore) (int64, error) {
	if c == nil || c.MD == nil || c.MD.UserId <= 0 {
		return 0, tg.ErrUserIdInvalid
	}
	return c.MD.UserId, nil
}

func requireUserClient(c *UserChannelProfilesCore) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.UserClient == nil {
		return fmt.Errorf("userchannelprofiles: user client is nil")
	}
	return nil
}

func requireMediaClient(c *UserChannelProfilesCore) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.MediaClient == nil {
		return fmt.Errorf("userchannelprofiles: media client is nil")
	}
	return nil
}

func userIDFromInputUser(selfID int64, in tg.InputUserClazz) (int64, error) {
	switch user := in.(type) {
	case *tg.TLInputUserSelf:
		if selfID <= 0 {
			return 0, tg.ErrUserIdInvalid
		}
		return selfID, nil
	case *tg.TLInputUser:
		if user.UserId <= 0 {
			return 0, tg.ErrUserIdInvalid
		}
		return user.UserId, nil
	default:
		return 0, tg.ErrUserIdInvalid
	}
}

func photoIDFromInputPhoto(in tg.InputPhotoClazz) (int64, error) {
	photo, ok := in.(*tg.TLInputPhoto)
	if !ok || photo.Id <= 0 {
		return 0, tg.ErrInputRequestInvalid
	}
	return photo.Id, nil
}

func documentIDFromInputDocument(in tg.InputDocumentClazz) (int64, error) {
	doc, ok := in.(*tg.TLInputDocument)
	if !ok || doc.Id <= 0 {
		return 0, tg.ErrInputRequestInvalid
	}
	return doc.Id, nil
}

func optionalDocumentID(in tg.InputDocumentClazz) *int64 {
	doc, ok := in.(*tg.TLInputDocument)
	if !ok || doc.Id <= 0 {
		return nil
	}
	return &doc.Id
}

func channelIDFromInputChannel(in tg.InputChannelClazz) (int64, error) {
	channel, ok := in.(*tg.TLInputChannel)
	if !ok || channel.ChannelId <= 0 {
		return 0, tg.ErrInputRequestInvalid
	}
	return channel.ChannelId, nil
}
