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
		return userpb.ErrUserStorage
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
