package repository

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	userSettingSensitiveEnabled          = "sensitive_enabled"
	userSettingContactSignUpNotification = "contactSignUpNotification"
)

func (r *Repository) GetContentSettings(ctx context.Context, userID int64) (*tg.AccountContentSettings, error) {
	enabled, err := r.getBoolUserSetting(ctx, userID, userSettingSensitiveEnabled, false)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLAccountContentSettings(&tg.TLAccountContentSettings{
		SensitiveEnabled:   enabled,
		SensitiveCanChange: true,
	}).ToAccountContentSettings(), nil
}

func (r *Repository) SetContentSettings(ctx context.Context, userID int64, sensitiveEnabled bool) error {
	return r.setBoolUserSetting(ctx, userID, userSettingSensitiveEnabled, sensitiveEnabled)
}

func (r *Repository) GetContactSignUpNotification(ctx context.Context, userID int64) (bool, error) {
	return r.getBoolUserSetting(ctx, userID, userSettingContactSignUpNotification, true)
}

func (r *Repository) SetContactSignUpNotification(ctx context.Context, userID int64, silent bool) error {
	return r.setBoolUserSetting(ctx, userID, userSettingContactSignUpNotification, !silent)
}

func (r *Repository) GetGlobalPrivacySettings(ctx context.Context, userID int64) (*tg.GlobalPrivacySettings, error) {
	settingsDO, err := r.model.UserGlobalPrivacySettingsModel.Select(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: get global privacy settings %d: %w", userpb.ErrUserStorage, userID, err)
	}
	if settingsDO == nil {
		return tg.MakeTLGlobalPrivacySettings(&tg.TLGlobalPrivacySettings{}).ToGlobalPrivacySettings(), nil
	}
	return tg.MakeTLGlobalPrivacySettings(&tg.TLGlobalPrivacySettings{
		ArchiveAndMuteNewNoncontactPeers: settingsDO.ArchiveAndMuteNewNoncontactPeers,
		KeepArchivedUnmuted:              settingsDO.KeepArchivedUnmuted,
		KeepArchivedFolders:              settingsDO.KeepArchivedFolders,
		HideReadMarks:                    settingsDO.HideReadMarks,
		NewNoncontactPeersRequirePremium: settingsDO.NewNoncontactPeersRequirePremium,
	}).ToGlobalPrivacySettings(), nil
}

func (r *Repository) SetGlobalPrivacySettings(ctx context.Context, userID int64, settings tg.GlobalPrivacySettingsClazz) error {
	if settings == nil {
		return userpb.ErrInvalidGlobalPrivacySettings
	}
	_, _, err := r.model.UserGlobalPrivacySettingsModel.InsertOrUpdate(ctx, &model.UserGlobalPrivacySettings{
		UserId:                           userID,
		ArchiveAndMuteNewNoncontactPeers: settings.ArchiveAndMuteNewNoncontactPeers,
		KeepArchivedUnmuted:              settings.KeepArchivedUnmuted,
		KeepArchivedFolders:              settings.KeepArchivedFolders,
		HideReadMarks:                    settings.HideReadMarks,
		NewNoncontactPeersRequirePremium: settings.NewNoncontactPeersRequirePremium,
	})
	if err != nil {
		return fmt.Errorf("%w: set global privacy settings %d: %w", userpb.ErrUserStorage, userID, err)
	}
	return nil
}

func (r *Repository) GetNotifySettings(ctx context.Context, userID int64, peerType int32, peerID int64) (*tg.PeerNotifySettings, error) {
	settingsDO, err := r.model.UserNotifySettingsModel.Select(ctx, userID, peerType, peerID)
	if err != nil {
		return nil, fmt.Errorf("%w: get notify settings %d/%d/%d: %w", userpb.ErrUserStorage, userID, peerType, peerID, err)
	}
	return makePeerNotifySettings(settingsDO), nil
}

func (r *Repository) GetNotifySettingsList(ctx context.Context, userID int64, peers []tg.PeerUtilClazz) (*userpb.VectorPeerPeerNotifySettings, error) {
	datas := make([]userpb.PeerPeerNotifySettingsClazz, 0, len(peers))
	for _, peer := range peers {
		if peer == nil {
			continue
		}
		settings, err := r.GetNotifySettings(ctx, userID, peer.PeerType, peer.PeerId)
		if err != nil {
			return nil, err
		}
		datas = append(datas, userpb.MakeTLPeerPeerNotifySettings(&userpb.TLPeerPeerNotifySettings{
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
			Settings: settings,
		}).ToPeerPeerNotifySettings())
	}
	return &userpb.VectorPeerPeerNotifySettings{Datas: datas}, nil
}

func (r *Repository) GetAllNotifySettings(ctx context.Context, userID int64) (*userpb.VectorPeerPeerNotifySettings, error) {
	settingsList, err := r.model.UserNotifySettingsModel.SelectAll(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: get all notify settings %d: %w", userpb.ErrUserStorage, userID, err)
	}
	datas := make([]userpb.PeerPeerNotifySettingsClazz, 0, len(settingsList))
	for i := range settingsList {
		settingsDO := settingsList[i]
		datas = append(datas, userpb.MakeTLPeerPeerNotifySettings(&userpb.TLPeerPeerNotifySettings{
			PeerType: settingsDO.PeerType,
			PeerId:   settingsDO.PeerId,
			Settings: makePeerNotifySettings(&settingsDO),
		}).ToPeerPeerNotifySettings())
	}
	return &userpb.VectorPeerPeerNotifySettings{Datas: datas}, nil
}

func (r *Repository) SetNotifySettings(ctx context.Context, userID int64, peerType int32, peerID int64, settings tg.PeerNotifySettingsClazz) error {
	settingsDO := makeUserNotifySettingsDO(userID, peerType, peerID, settings)
	_, _, err := r.model.UserNotifySettingsModel.InsertOrUpdate(ctx, settingsDO)
	if err != nil {
		return fmt.Errorf("%w: set notify settings %d/%d/%d: %w", userpb.ErrUserStorage, userID, peerType, peerID, err)
	}
	return nil
}

func (r *Repository) ResetNotifySettings(ctx context.Context, userID int64) error {
	_, err := r.model.UserNotifySettingsModel.DeleteAll(ctx, userID)
	if err != nil {
		return fmt.Errorf("%w: reset notify settings %d: %w", userpb.ErrUserStorage, userID, err)
	}
	return nil
}

func (r *Repository) GetPeerSettings(ctx context.Context, userID int64, peerType int32, peerID int64) (*tg.PeerSettings, error) {
	settingsDO, err := r.model.UserPeerSettingsModel.Select(ctx, userID, peerType, peerID)
	if err != nil {
		return nil, fmt.Errorf("%w: get peer settings %d/%d/%d: %w", userpb.ErrUserStorage, userID, peerType, peerID, err)
	}
	return makePeerSettings(settingsDO), nil
}

func (r *Repository) AddPeerSettings(ctx context.Context, userID int64, peerType int32, peerID int64, settings tg.PeerSettingsClazz) error {
	settingsDO := makeUserPeerSettingsDO(userID, peerType, peerID, settings)
	_, _, err := r.model.UserPeerSettingsModel.InsertOrUpdate(ctx, settingsDO)
	if err != nil {
		return fmt.Errorf("%w: add peer settings %d/%d/%d: %w", userpb.ErrUserStorage, userID, peerType, peerID, err)
	}
	return nil
}

func (r *Repository) DeletePeerSettings(ctx context.Context, userID int64, peerType int32, peerID int64) error {
	_, err := r.model.UserPeerSettingsModel.Delete(ctx, userID, peerType, peerID)
	if err != nil {
		return fmt.Errorf("%w: delete peer settings %d/%d/%d: %w", userpb.ErrUserStorage, userID, peerType, peerID, err)
	}
	return nil
}

func (r *Repository) getBoolUserSetting(ctx context.Context, userID int64, key string, defaultValue bool) (bool, error) {
	settingsDO, err := r.model.UserSettingsModel.SelectByKey(ctx, userID, key)
	if err != nil {
		return false, fmt.Errorf("%w: get user setting %d/%s: %w", userpb.ErrUserStorage, userID, key, err)
	}
	if settingsDO == nil {
		return defaultValue, nil
	}
	return settingsDO.Value == "true", nil
}

func (r *Repository) setBoolUserSetting(ctx context.Context, userID int64, key string, value bool) error {
	valueString := "false"
	if value {
		valueString = "true"
	}
	_, _, err := r.model.UserSettingsModel.InsertOrUpdate(ctx, &model.UserSettings{
		UserId: userID,
		Key2:   key,
		Value:  valueString,
	})
	if err != nil {
		return fmt.Errorf("%w: set user setting %d/%s: %w", userpb.ErrUserStorage, userID, key, err)
	}
	return nil
}

func makePeerNotifySettings(settingsDO *model.UserNotifySettings) *tg.PeerNotifySettings {
	settings := tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}).ToPeerNotifySettings()
	if settingsDO == nil {
		return settings
	}
	if settingsDO.ShowPreviews != -1 {
		settings.ShowPreviews = tg.ToBoolClazz(settingsDO.ShowPreviews == 1)
	}
	if settingsDO.Silent != -1 {
		settings.Silent = tg.ToBoolClazz(settingsDO.Silent == 1)
	}
	if settingsDO.MuteUntil != -1 {
		muteUntil := settingsDO.MuteUntil
		settings.MuteUntil = &muteUntil
	}
	settings.IosSound = makeNotificationSound(settingsDO.Sound)
	settings.AndroidSound = makeNotificationSound(settingsDO.Sound)
	settings.OtherSound = makeNotificationSound(settingsDO.Sound)
	return settings
}

func makeUserNotifySettingsDO(userID int64, peerType int32, peerID int64, settings tg.PeerNotifySettingsClazz) *model.UserNotifySettings {
	settingsDO := &model.UserNotifySettings{
		UserId:       userID,
		PeerType:     peerType,
		PeerId:       peerID,
		ShowPreviews: -1,
		Silent:       -1,
		MuteUntil:    -1,
		Sound:        "-1",
	}
	if settings == nil {
		return settingsDO
	}
	if settings.ShowPreviews != nil {
		settingsDO.ShowPreviews = boolInt32(tg.FromBoolClazz(settings.ShowPreviews))
	}
	if settings.Silent != nil {
		settingsDO.Silent = boolInt32(tg.FromBoolClazz(settings.Silent))
	}
	if settings.MuteUntil != nil {
		settingsDO.MuteUntil = *settings.MuteUntil
	}
	settingsDO.Sound = notificationSoundString(settings.IosSound)
	return settingsDO
}

func makeNotificationSound(sound string) tg.NotificationSoundClazz {
	if sound == "" || sound == "-1" {
		return nil
	}
	if sound == "default" {
		return tg.NotificationSoundDefaultClazz
	}
	if sound == "none" {
		return tg.NotificationSoundNoneClazz
	}
	return tg.MakeTLNotificationSoundLocal(&tg.TLNotificationSoundLocal{
		Title: sound,
		Data:  sound,
	}).ToNotificationSound().Clazz
}

func notificationSoundString(sound tg.NotificationSoundClazz) string {
	switch v := sound.(type) {
	case *tg.TLNotificationSoundDefault:
		return "default"
	case *tg.TLNotificationSoundNone:
		return "none"
	case *tg.TLNotificationSoundLocal:
		return v.Data
	default:
		return "-1"
	}
}

func makePeerSettings(settingsDO *model.UserPeerSettings) *tg.PeerSettings {
	if settingsDO == nil {
		return tg.MakeTLPeerSettings(&tg.TLPeerSettings{}).ToPeerSettings()
	}
	settings := tg.MakeTLPeerSettings(&tg.TLPeerSettings{
		ReportSpam:            settingsDO.ReportSpam,
		AddContact:            settingsDO.AddContact,
		BlockContact:          settingsDO.BlockContact,
		ShareContact:          settingsDO.ShareContact,
		NeedContactsException: settingsDO.NeedContactsException,
		ReportGeo:             settingsDO.ReportGeo,
		Autoarchived:          settingsDO.Autoarchived,
		InviteMembers:         settingsDO.InviteMembers,
	}).ToPeerSettings()
	if settingsDO.GeoDistance != 0 {
		geoDistance := settingsDO.GeoDistance
		settings.GeoDistance = &geoDistance
	}
	return settings
}

func makeUserPeerSettingsDO(userID int64, peerType int32, peerID int64, settings tg.PeerSettingsClazz) *model.UserPeerSettings {
	settingsDO := &model.UserPeerSettings{
		UserId:   userID,
		PeerType: peerType,
		PeerId:   peerID,
	}
	if settings == nil {
		return settingsDO
	}
	settingsDO.ReportSpam = settings.ReportSpam
	settingsDO.AddContact = settings.AddContact
	settingsDO.BlockContact = settings.BlockContact
	settingsDO.ShareContact = settings.ShareContact
	settingsDO.NeedContactsException = settings.NeedContactsException
	settingsDO.ReportGeo = settings.ReportGeo
	settingsDO.Autoarchived = settings.Autoarchived
	settingsDO.InviteMembers = settings.InviteMembers
	if settings.GeoDistance != nil {
		settingsDO.GeoDistance = *settings.GeoDistance
	}
	return settingsDO
}

func boolInt32(v bool) int32 {
	if v {
		return 1
	}
	return 0
}
